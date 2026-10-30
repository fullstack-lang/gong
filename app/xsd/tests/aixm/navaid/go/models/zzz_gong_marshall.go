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
	modelsPackageName := "github.com/fullstack-lang/gong/app/xsd/tests/aixm/navaid/go/models"
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
		for _, codeairspacepointpositiontype := range __gong__sortStageSetInstances(stage.AixmdatatypesStage.CodeAirspacePointPositionTypes, stage.AixmdatatypesStage.CodeAirspacePointPositionType_stagedOrder) {
			if lastStageDecl != "AixmdatatypesStage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "AixmdatatypesStage"
			}
			codeairspacepointpositiontypeIdent := codeairspacepointpositiontype.GongGetIdentifier(stage.AixmdatatypesStage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&aixmdatatypes.CodeAirspacePointPositionType{Name: %s}).Stage(stage.AixmdatatypesStage)", codeairspacepointpositiontypeIdent, __gong__toRawStringLiteral(codeairspacepointpositiontype.Name)))
			if lastStageVal != "AixmdatatypesStage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "AixmdatatypesStage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", codeairspacepointpositiontypeIdent, __gong__toRawStringLiteral(codeairspacepointpositiontype.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.NilReason = %s", codeairspacepointpositiontypeIdent, __gong__toRawStringLiteral(codeairspacepointpositiontype.NilReason)))
			values.WriteString(fmt.Sprintf("\n\t%s.EnclosedText = %s", codeairspacepointpositiontypeIdent, __gong__toRawStringLiteral(codeairspacepointpositiontype.EnclosedText)))
		}
	}
	if stage.AixmdatatypesStage != nil {
		for _, codeairspacepointroletype := range __gong__sortStageSetInstances(stage.AixmdatatypesStage.CodeAirspacePointRoleTypes, stage.AixmdatatypesStage.CodeAirspacePointRoleType_stagedOrder) {
			if lastStageDecl != "AixmdatatypesStage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "AixmdatatypesStage"
			}
			codeairspacepointroletypeIdent := codeairspacepointroletype.GongGetIdentifier(stage.AixmdatatypesStage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&aixmdatatypes.CodeAirspacePointRoleType{Name: %s}).Stage(stage.AixmdatatypesStage)", codeairspacepointroletypeIdent, __gong__toRawStringLiteral(codeairspacepointroletype.Name)))
			if lastStageVal != "AixmdatatypesStage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "AixmdatatypesStage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", codeairspacepointroletypeIdent, __gong__toRawStringLiteral(codeairspacepointroletype.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.NilReason = %s", codeairspacepointroletypeIdent, __gong__toRawStringLiteral(codeairspacepointroletype.NilReason)))
			values.WriteString(fmt.Sprintf("\n\t%s.EnclosedText = %s", codeairspacepointroletypeIdent, __gong__toRawStringLiteral(codeairspacepointroletype.EnclosedText)))
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
		for _, codeauralmorsetype := range __gong__sortStageSetInstances(stage.AixmdatatypesStage.CodeAuralMorseTypes, stage.AixmdatatypesStage.CodeAuralMorseType_stagedOrder) {
			if lastStageDecl != "AixmdatatypesStage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "AixmdatatypesStage"
			}
			codeauralmorsetypeIdent := codeauralmorsetype.GongGetIdentifier(stage.AixmdatatypesStage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&aixmdatatypes.CodeAuralMorseType{Name: %s}).Stage(stage.AixmdatatypesStage)", codeauralmorsetypeIdent, __gong__toRawStringLiteral(codeauralmorsetype.Name)))
			if lastStageVal != "AixmdatatypesStage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "AixmdatatypesStage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", codeauralmorsetypeIdent, __gong__toRawStringLiteral(codeauralmorsetype.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.NilReason = %s", codeauralmorsetypeIdent, __gong__toRawStringLiteral(codeauralmorsetype.NilReason)))
			values.WriteString(fmt.Sprintf("\n\t%s.EnclosedText = %s", codeauralmorsetypeIdent, __gong__toRawStringLiteral(codeauralmorsetype.EnclosedText)))
		}
	}
	if stage.AixmdatatypesStage != nil {
		for _, codeauthorityroletype := range __gong__sortStageSetInstances(stage.AixmdatatypesStage.CodeAuthorityRoleTypes, stage.AixmdatatypesStage.CodeAuthorityRoleType_stagedOrder) {
			if lastStageDecl != "AixmdatatypesStage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "AixmdatatypesStage"
			}
			codeauthorityroletypeIdent := codeauthorityroletype.GongGetIdentifier(stage.AixmdatatypesStage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&aixmdatatypes.CodeAuthorityRoleType{Name: %s}).Stage(stage.AixmdatatypesStage)", codeauthorityroletypeIdent, __gong__toRawStringLiteral(codeauthorityroletype.Name)))
			if lastStageVal != "AixmdatatypesStage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "AixmdatatypesStage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", codeauthorityroletypeIdent, __gong__toRawStringLiteral(codeauthorityroletype.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.NilReason = %s", codeauthorityroletypeIdent, __gong__toRawStringLiteral(codeauthorityroletype.NilReason)))
			values.WriteString(fmt.Sprintf("\n\t%s.EnclosedText = %s", codeauthorityroletypeIdent, __gong__toRawStringLiteral(codeauthorityroletype.EnclosedText)))
		}
	}
	if stage.AixmdatatypesStage != nil {
		for _, codebearingtype := range __gong__sortStageSetInstances(stage.AixmdatatypesStage.CodeBearingTypes, stage.AixmdatatypesStage.CodeBearingType_stagedOrder) {
			if lastStageDecl != "AixmdatatypesStage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "AixmdatatypesStage"
			}
			codebearingtypeIdent := codebearingtype.GongGetIdentifier(stage.AixmdatatypesStage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&aixmdatatypes.CodeBearingType{Name: %s}).Stage(stage.AixmdatatypesStage)", codebearingtypeIdent, __gong__toRawStringLiteral(codebearingtype.Name)))
			if lastStageVal != "AixmdatatypesStage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "AixmdatatypesStage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", codebearingtypeIdent, __gong__toRawStringLiteral(codebearingtype.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.NilReason = %s", codebearingtypeIdent, __gong__toRawStringLiteral(codebearingtype.NilReason)))
			values.WriteString(fmt.Sprintf("\n\t%s.EnclosedText = %s", codebearingtypeIdent, __gong__toRawStringLiteral(codebearingtype.EnclosedText)))
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
		for _, codecheckpointcategorytype := range __gong__sortStageSetInstances(stage.AixmdatatypesStage.CodeCheckpointCategoryTypes, stage.AixmdatatypesStage.CodeCheckpointCategoryType_stagedOrder) {
			if lastStageDecl != "AixmdatatypesStage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "AixmdatatypesStage"
			}
			codecheckpointcategorytypeIdent := codecheckpointcategorytype.GongGetIdentifier(stage.AixmdatatypesStage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&aixmdatatypes.CodeCheckpointCategoryType{Name: %s}).Stage(stage.AixmdatatypesStage)", codecheckpointcategorytypeIdent, __gong__toRawStringLiteral(codecheckpointcategorytype.Name)))
			if lastStageVal != "AixmdatatypesStage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "AixmdatatypesStage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", codecheckpointcategorytypeIdent, __gong__toRawStringLiteral(codecheckpointcategorytype.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.NilReason = %s", codecheckpointcategorytypeIdent, __gong__toRawStringLiteral(codecheckpointcategorytype.NilReason)))
			values.WriteString(fmt.Sprintf("\n\t%s.EnclosedText = %s", codecheckpointcategorytypeIdent, __gong__toRawStringLiteral(codecheckpointcategorytype.EnclosedText)))
		}
	}
	if stage.AixmdatatypesStage != nil {
		for _, codecoursequalityilstype := range __gong__sortStageSetInstances(stage.AixmdatatypesStage.CodeCourseQualityILSTypes, stage.AixmdatatypesStage.CodeCourseQualityILSType_stagedOrder) {
			if lastStageDecl != "AixmdatatypesStage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "AixmdatatypesStage"
			}
			codecoursequalityilstypeIdent := codecoursequalityilstype.GongGetIdentifier(stage.AixmdatatypesStage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&aixmdatatypes.CodeCourseQualityILSType{Name: %s}).Stage(stage.AixmdatatypesStage)", codecoursequalityilstypeIdent, __gong__toRawStringLiteral(codecoursequalityilstype.Name)))
			if lastStageVal != "AixmdatatypesStage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "AixmdatatypesStage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", codecoursequalityilstypeIdent, __gong__toRawStringLiteral(codecoursequalityilstype.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.NilReason = %s", codecoursequalityilstypeIdent, __gong__toRawStringLiteral(codecoursequalityilstype.NilReason)))
			values.WriteString(fmt.Sprintf("\n\t%s.EnclosedText = %s", codecoursequalityilstypeIdent, __gong__toRawStringLiteral(codecoursequalityilstype.EnclosedText)))
		}
	}
	if stage.AixmdatatypesStage != nil {
		for _, codedmechanneltype := range __gong__sortStageSetInstances(stage.AixmdatatypesStage.CodeDMEChannelTypes, stage.AixmdatatypesStage.CodeDMEChannelType_stagedOrder) {
			if lastStageDecl != "AixmdatatypesStage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "AixmdatatypesStage"
			}
			codedmechanneltypeIdent := codedmechanneltype.GongGetIdentifier(stage.AixmdatatypesStage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&aixmdatatypes.CodeDMEChannelType{Name: %s}).Stage(stage.AixmdatatypesStage)", codedmechanneltypeIdent, __gong__toRawStringLiteral(codedmechanneltype.Name)))
			if lastStageVal != "AixmdatatypesStage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "AixmdatatypesStage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", codedmechanneltypeIdent, __gong__toRawStringLiteral(codedmechanneltype.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.NilReason = %s", codedmechanneltypeIdent, __gong__toRawStringLiteral(codedmechanneltype.NilReason)))
			values.WriteString(fmt.Sprintf("\n\t%s.EnclosedText = %s", codedmechanneltypeIdent, __gong__toRawStringLiteral(codedmechanneltype.EnclosedText)))
		}
	}
	if stage.AixmdatatypesStage != nil {
		for _, codedmetype := range __gong__sortStageSetInstances(stage.AixmdatatypesStage.CodeDMETypes, stage.AixmdatatypesStage.CodeDMEType_stagedOrder) {
			if lastStageDecl != "AixmdatatypesStage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "AixmdatatypesStage"
			}
			codedmetypeIdent := codedmetype.GongGetIdentifier(stage.AixmdatatypesStage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&aixmdatatypes.CodeDMEType{Name: %s}).Stage(stage.AixmdatatypesStage)", codedmetypeIdent, __gong__toRawStringLiteral(codedmetype.Name)))
			if lastStageVal != "AixmdatatypesStage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "AixmdatatypesStage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", codedmetypeIdent, __gong__toRawStringLiteral(codedmetype.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.NilReason = %s", codedmetypeIdent, __gong__toRawStringLiteral(codedmetype.NilReason)))
			values.WriteString(fmt.Sprintf("\n\t%s.EnclosedText = %s", codedmetypeIdent, __gong__toRawStringLiteral(codedmetype.EnclosedText)))
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
		for _, codedesignatedpointdesignatortype := range __gong__sortStageSetInstances(stage.AixmdatatypesStage.CodeDesignatedPointDesignatorTypes, stage.AixmdatatypesStage.CodeDesignatedPointDesignatorType_stagedOrder) {
			if lastStageDecl != "AixmdatatypesStage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "AixmdatatypesStage"
			}
			codedesignatedpointdesignatortypeIdent := codedesignatedpointdesignatortype.GongGetIdentifier(stage.AixmdatatypesStage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&aixmdatatypes.CodeDesignatedPointDesignatorType{Name: %s}).Stage(stage.AixmdatatypesStage)", codedesignatedpointdesignatortypeIdent, __gong__toRawStringLiteral(codedesignatedpointdesignatortype.Name)))
			if lastStageVal != "AixmdatatypesStage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "AixmdatatypesStage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", codedesignatedpointdesignatortypeIdent, __gong__toRawStringLiteral(codedesignatedpointdesignatortype.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.NilReason = %s", codedesignatedpointdesignatortypeIdent, __gong__toRawStringLiteral(codedesignatedpointdesignatortype.NilReason)))
			values.WriteString(fmt.Sprintf("\n\t%s.EnclosedText = %s", codedesignatedpointdesignatortypeIdent, __gong__toRawStringLiteral(codedesignatedpointdesignatortype.EnclosedText)))
		}
	}
	if stage.AixmdatatypesStage != nil {
		for _, codedesignatedpointtype := range __gong__sortStageSetInstances(stage.AixmdatatypesStage.CodeDesignatedPointTypes, stage.AixmdatatypesStage.CodeDesignatedPointType_stagedOrder) {
			if lastStageDecl != "AixmdatatypesStage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "AixmdatatypesStage"
			}
			codedesignatedpointtypeIdent := codedesignatedpointtype.GongGetIdentifier(stage.AixmdatatypesStage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&aixmdatatypes.CodeDesignatedPointType{Name: %s}).Stage(stage.AixmdatatypesStage)", codedesignatedpointtypeIdent, __gong__toRawStringLiteral(codedesignatedpointtype.Name)))
			if lastStageVal != "AixmdatatypesStage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "AixmdatatypesStage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", codedesignatedpointtypeIdent, __gong__toRawStringLiteral(codedesignatedpointtype.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.NilReason = %s", codedesignatedpointtypeIdent, __gong__toRawStringLiteral(codedesignatedpointtype.NilReason)))
			values.WriteString(fmt.Sprintf("\n\t%s.EnclosedText = %s", codedesignatedpointtypeIdent, __gong__toRawStringLiteral(codedesignatedpointtype.EnclosedText)))
		}
	}
	if stage.AixmdatatypesStage != nil {
		for _, codedirectionreferencetype := range __gong__sortStageSetInstances(stage.AixmdatatypesStage.CodeDirectionReferenceTypes, stage.AixmdatatypesStage.CodeDirectionReferenceType_stagedOrder) {
			if lastStageDecl != "AixmdatatypesStage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "AixmdatatypesStage"
			}
			codedirectionreferencetypeIdent := codedirectionreferencetype.GongGetIdentifier(stage.AixmdatatypesStage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&aixmdatatypes.CodeDirectionReferenceType{Name: %s}).Stage(stage.AixmdatatypesStage)", codedirectionreferencetypeIdent, __gong__toRawStringLiteral(codedirectionreferencetype.Name)))
			if lastStageVal != "AixmdatatypesStage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "AixmdatatypesStage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", codedirectionreferencetypeIdent, __gong__toRawStringLiteral(codedirectionreferencetype.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.NilReason = %s", codedirectionreferencetypeIdent, __gong__toRawStringLiteral(codedirectionreferencetype.NilReason)))
			values.WriteString(fmt.Sprintf("\n\t%s.EnclosedText = %s", codedirectionreferencetypeIdent, __gong__toRawStringLiteral(codedirectionreferencetype.EnclosedText)))
		}
	}
	if stage.AixmdatatypesStage != nil {
		for _, codedistanceindicationtype := range __gong__sortStageSetInstances(stage.AixmdatatypesStage.CodeDistanceIndicationTypes, stage.AixmdatatypesStage.CodeDistanceIndicationType_stagedOrder) {
			if lastStageDecl != "AixmdatatypesStage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "AixmdatatypesStage"
			}
			codedistanceindicationtypeIdent := codedistanceindicationtype.GongGetIdentifier(stage.AixmdatatypesStage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&aixmdatatypes.CodeDistanceIndicationType{Name: %s}).Stage(stage.AixmdatatypesStage)", codedistanceindicationtypeIdent, __gong__toRawStringLiteral(codedistanceindicationtype.Name)))
			if lastStageVal != "AixmdatatypesStage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "AixmdatatypesStage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", codedistanceindicationtypeIdent, __gong__toRawStringLiteral(codedistanceindicationtype.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.NilReason = %s", codedistanceindicationtypeIdent, __gong__toRawStringLiteral(codedistanceindicationtype.NilReason)))
			values.WriteString(fmt.Sprintf("\n\t%s.EnclosedText = %s", codedistanceindicationtypeIdent, __gong__toRawStringLiteral(codedistanceindicationtype.EnclosedText)))
		}
	}
	if stage.AixmdatatypesStage != nil {
		for _, codeemissionbandtype := range __gong__sortStageSetInstances(stage.AixmdatatypesStage.CodeEmissionBandTypes, stage.AixmdatatypesStage.CodeEmissionBandType_stagedOrder) {
			if lastStageDecl != "AixmdatatypesStage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "AixmdatatypesStage"
			}
			codeemissionbandtypeIdent := codeemissionbandtype.GongGetIdentifier(stage.AixmdatatypesStage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&aixmdatatypes.CodeEmissionBandType{Name: %s}).Stage(stage.AixmdatatypesStage)", codeemissionbandtypeIdent, __gong__toRawStringLiteral(codeemissionbandtype.Name)))
			if lastStageVal != "AixmdatatypesStage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "AixmdatatypesStage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", codeemissionbandtypeIdent, __gong__toRawStringLiteral(codeemissionbandtype.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.NilReason = %s", codeemissionbandtypeIdent, __gong__toRawStringLiteral(codeemissionbandtype.NilReason)))
			values.WriteString(fmt.Sprintf("\n\t%s.EnclosedText = %s", codeemissionbandtypeIdent, __gong__toRawStringLiteral(codeemissionbandtype.EnclosedText)))
		}
	}
	if stage.AixmdatatypesStage != nil {
		for _, codeilsbackcoursetype := range __gong__sortStageSetInstances(stage.AixmdatatypesStage.CodeILSBackCourseTypes, stage.AixmdatatypesStage.CodeILSBackCourseType_stagedOrder) {
			if lastStageDecl != "AixmdatatypesStage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "AixmdatatypesStage"
			}
			codeilsbackcoursetypeIdent := codeilsbackcoursetype.GongGetIdentifier(stage.AixmdatatypesStage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&aixmdatatypes.CodeILSBackCourseType{Name: %s}).Stage(stage.AixmdatatypesStage)", codeilsbackcoursetypeIdent, __gong__toRawStringLiteral(codeilsbackcoursetype.Name)))
			if lastStageVal != "AixmdatatypesStage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "AixmdatatypesStage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", codeilsbackcoursetypeIdent, __gong__toRawStringLiteral(codeilsbackcoursetype.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.NilReason = %s", codeilsbackcoursetypeIdent, __gong__toRawStringLiteral(codeilsbackcoursetype.NilReason)))
			values.WriteString(fmt.Sprintf("\n\t%s.EnclosedText = %s", codeilsbackcoursetypeIdent, __gong__toRawStringLiteral(codeilsbackcoursetype.EnclosedText)))
		}
	}
	if stage.AixmdatatypesStage != nil {
		for _, codeintegritylevelilstype := range __gong__sortStageSetInstances(stage.AixmdatatypesStage.CodeIntegrityLevelILSTypes, stage.AixmdatatypesStage.CodeIntegrityLevelILSType_stagedOrder) {
			if lastStageDecl != "AixmdatatypesStage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "AixmdatatypesStage"
			}
			codeintegritylevelilstypeIdent := codeintegritylevelilstype.GongGetIdentifier(stage.AixmdatatypesStage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&aixmdatatypes.CodeIntegrityLevelILSType{Name: %s}).Stage(stage.AixmdatatypesStage)", codeintegritylevelilstypeIdent, __gong__toRawStringLiteral(codeintegritylevelilstype.Name)))
			if lastStageVal != "AixmdatatypesStage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "AixmdatatypesStage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", codeintegritylevelilstypeIdent, __gong__toRawStringLiteral(codeintegritylevelilstype.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.NilReason = %s", codeintegritylevelilstypeIdent, __gong__toRawStringLiteral(codeintegritylevelilstype.NilReason)))
			values.WriteString(fmt.Sprintf("\n\t%s.EnclosedText = %s", codeintegritylevelilstypeIdent, __gong__toRawStringLiteral(codeintegritylevelilstype.EnclosedText)))
		}
	}
	if stage.AixmdatatypesStage != nil {
		for _, codemlsazimuthtype := range __gong__sortStageSetInstances(stage.AixmdatatypesStage.CodeMLSAzimuthTypes, stage.AixmdatatypesStage.CodeMLSAzimuthType_stagedOrder) {
			if lastStageDecl != "AixmdatatypesStage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "AixmdatatypesStage"
			}
			codemlsazimuthtypeIdent := codemlsazimuthtype.GongGetIdentifier(stage.AixmdatatypesStage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&aixmdatatypes.CodeMLSAzimuthType{Name: %s}).Stage(stage.AixmdatatypesStage)", codemlsazimuthtypeIdent, __gong__toRawStringLiteral(codemlsazimuthtype.Name)))
			if lastStageVal != "AixmdatatypesStage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "AixmdatatypesStage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", codemlsazimuthtypeIdent, __gong__toRawStringLiteral(codemlsazimuthtype.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.NilReason = %s", codemlsazimuthtypeIdent, __gong__toRawStringLiteral(codemlsazimuthtype.NilReason)))
			values.WriteString(fmt.Sprintf("\n\t%s.EnclosedText = %s", codemlsazimuthtypeIdent, __gong__toRawStringLiteral(codemlsazimuthtype.EnclosedText)))
		}
	}
	if stage.AixmdatatypesStage != nil {
		for _, codemlschanneltype := range __gong__sortStageSetInstances(stage.AixmdatatypesStage.CodeMLSChannelTypes, stage.AixmdatatypesStage.CodeMLSChannelType_stagedOrder) {
			if lastStageDecl != "AixmdatatypesStage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "AixmdatatypesStage"
			}
			codemlschanneltypeIdent := codemlschanneltype.GongGetIdentifier(stage.AixmdatatypesStage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&aixmdatatypes.CodeMLSChannelType{Name: %s}).Stage(stage.AixmdatatypesStage)", codemlschanneltypeIdent, __gong__toRawStringLiteral(codemlschanneltype.Name)))
			if lastStageVal != "AixmdatatypesStage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "AixmdatatypesStage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", codemlschanneltypeIdent, __gong__toRawStringLiteral(codemlschanneltype.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.NilReason = %s", codemlschanneltypeIdent, __gong__toRawStringLiteral(codemlschanneltype.NilReason)))
			values.WriteString(fmt.Sprintf("\n\t%s.EnclosedText = %s", codemlschanneltypeIdent, __gong__toRawStringLiteral(codemlschanneltype.EnclosedText)))
		}
	}
	if stage.AixmdatatypesStage != nil {
		for _, codemarkerbeaconsignaltype := range __gong__sortStageSetInstances(stage.AixmdatatypesStage.CodeMarkerBeaconSignalTypes, stage.AixmdatatypesStage.CodeMarkerBeaconSignalType_stagedOrder) {
			if lastStageDecl != "AixmdatatypesStage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "AixmdatatypesStage"
			}
			codemarkerbeaconsignaltypeIdent := codemarkerbeaconsignaltype.GongGetIdentifier(stage.AixmdatatypesStage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&aixmdatatypes.CodeMarkerBeaconSignalType{Name: %s}).Stage(stage.AixmdatatypesStage)", codemarkerbeaconsignaltypeIdent, __gong__toRawStringLiteral(codemarkerbeaconsignaltype.Name)))
			if lastStageVal != "AixmdatatypesStage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "AixmdatatypesStage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", codemarkerbeaconsignaltypeIdent, __gong__toRawStringLiteral(codemarkerbeaconsignaltype.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.NilReason = %s", codemarkerbeaconsignaltypeIdent, __gong__toRawStringLiteral(codemarkerbeaconsignaltype.NilReason)))
			values.WriteString(fmt.Sprintf("\n\t%s.EnclosedText = %s", codemarkerbeaconsignaltypeIdent, __gong__toRawStringLiteral(codemarkerbeaconsignaltype.EnclosedText)))
		}
	}
	if stage.AixmdatatypesStage != nil {
		for _, codendbusagetype := range __gong__sortStageSetInstances(stage.AixmdatatypesStage.CodeNDBUsageTypes, stage.AixmdatatypesStage.CodeNDBUsageType_stagedOrder) {
			if lastStageDecl != "AixmdatatypesStage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "AixmdatatypesStage"
			}
			codendbusagetypeIdent := codendbusagetype.GongGetIdentifier(stage.AixmdatatypesStage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&aixmdatatypes.CodeNDBUsageType{Name: %s}).Stage(stage.AixmdatatypesStage)", codendbusagetypeIdent, __gong__toRawStringLiteral(codendbusagetype.Name)))
			if lastStageVal != "AixmdatatypesStage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "AixmdatatypesStage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", codendbusagetypeIdent, __gong__toRawStringLiteral(codendbusagetype.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.NilReason = %s", codendbusagetypeIdent, __gong__toRawStringLiteral(codendbusagetype.NilReason)))
			values.WriteString(fmt.Sprintf("\n\t%s.EnclosedText = %s", codendbusagetypeIdent, __gong__toRawStringLiteral(codendbusagetype.EnclosedText)))
		}
	}
	if stage.AixmdatatypesStage != nil {
		for _, codenavaiddesignatortype := range __gong__sortStageSetInstances(stage.AixmdatatypesStage.CodeNavaidDesignatorTypes, stage.AixmdatatypesStage.CodeNavaidDesignatorType_stagedOrder) {
			if lastStageDecl != "AixmdatatypesStage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "AixmdatatypesStage"
			}
			codenavaiddesignatortypeIdent := codenavaiddesignatortype.GongGetIdentifier(stage.AixmdatatypesStage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&aixmdatatypes.CodeNavaidDesignatorType{Name: %s}).Stage(stage.AixmdatatypesStage)", codenavaiddesignatortypeIdent, __gong__toRawStringLiteral(codenavaiddesignatortype.Name)))
			if lastStageVal != "AixmdatatypesStage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "AixmdatatypesStage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", codenavaiddesignatortypeIdent, __gong__toRawStringLiteral(codenavaiddesignatortype.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.NilReason = %s", codenavaiddesignatortypeIdent, __gong__toRawStringLiteral(codenavaiddesignatortype.NilReason)))
			values.WriteString(fmt.Sprintf("\n\t%s.EnclosedText = %s", codenavaiddesignatortypeIdent, __gong__toRawStringLiteral(codenavaiddesignatortype.EnclosedText)))
		}
	}
	if stage.AixmdatatypesStage != nil {
		for _, codenavaidpurposetype := range __gong__sortStageSetInstances(stage.AixmdatatypesStage.CodeNavaidPurposeTypes, stage.AixmdatatypesStage.CodeNavaidPurposeType_stagedOrder) {
			if lastStageDecl != "AixmdatatypesStage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "AixmdatatypesStage"
			}
			codenavaidpurposetypeIdent := codenavaidpurposetype.GongGetIdentifier(stage.AixmdatatypesStage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&aixmdatatypes.CodeNavaidPurposeType{Name: %s}).Stage(stage.AixmdatatypesStage)", codenavaidpurposetypeIdent, __gong__toRawStringLiteral(codenavaidpurposetype.Name)))
			if lastStageVal != "AixmdatatypesStage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "AixmdatatypesStage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", codenavaidpurposetypeIdent, __gong__toRawStringLiteral(codenavaidpurposetype.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.NilReason = %s", codenavaidpurposetypeIdent, __gong__toRawStringLiteral(codenavaidpurposetype.NilReason)))
			values.WriteString(fmt.Sprintf("\n\t%s.EnclosedText = %s", codenavaidpurposetypeIdent, __gong__toRawStringLiteral(codenavaidpurposetype.EnclosedText)))
		}
	}
	if stage.AixmdatatypesStage != nil {
		for _, codenavaidservicetype := range __gong__sortStageSetInstances(stage.AixmdatatypesStage.CodeNavaidServiceTypes, stage.AixmdatatypesStage.CodeNavaidServiceType_stagedOrder) {
			if lastStageDecl != "AixmdatatypesStage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "AixmdatatypesStage"
			}
			codenavaidservicetypeIdent := codenavaidservicetype.GongGetIdentifier(stage.AixmdatatypesStage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&aixmdatatypes.CodeNavaidServiceType{Name: %s}).Stage(stage.AixmdatatypesStage)", codenavaidservicetypeIdent, __gong__toRawStringLiteral(codenavaidservicetype.Name)))
			if lastStageVal != "AixmdatatypesStage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "AixmdatatypesStage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", codenavaidservicetypeIdent, __gong__toRawStringLiteral(codenavaidservicetype.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.NilReason = %s", codenavaidservicetypeIdent, __gong__toRawStringLiteral(codenavaidservicetype.NilReason)))
			values.WriteString(fmt.Sprintf("\n\t%s.EnclosedText = %s", codenavaidservicetypeIdent, __gong__toRawStringLiteral(codenavaidservicetype.EnclosedText)))
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
		for _, codepartype := range __gong__sortStageSetInstances(stage.AixmdatatypesStage.CodePARTypes, stage.AixmdatatypesStage.CodePARType_stagedOrder) {
			if lastStageDecl != "AixmdatatypesStage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "AixmdatatypesStage"
			}
			codepartypeIdent := codepartype.GongGetIdentifier(stage.AixmdatatypesStage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&aixmdatatypes.CodePARType{Name: %s}).Stage(stage.AixmdatatypesStage)", codepartypeIdent, __gong__toRawStringLiteral(codepartype.Name)))
			if lastStageVal != "AixmdatatypesStage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "AixmdatatypesStage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", codepartypeIdent, __gong__toRawStringLiteral(codepartype.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.NilReason = %s", codepartypeIdent, __gong__toRawStringLiteral(codepartype.NilReason)))
			values.WriteString(fmt.Sprintf("\n\t%s.EnclosedText = %s", codepartypeIdent, __gong__toRawStringLiteral(codepartype.EnclosedText)))
		}
	}
	if stage.AixmdatatypesStage != nil {
		for _, codepositioninilstype := range __gong__sortStageSetInstances(stage.AixmdatatypesStage.CodePositionInILSTypes, stage.AixmdatatypesStage.CodePositionInILSType_stagedOrder) {
			if lastStageDecl != "AixmdatatypesStage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "AixmdatatypesStage"
			}
			codepositioninilstypeIdent := codepositioninilstype.GongGetIdentifier(stage.AixmdatatypesStage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&aixmdatatypes.CodePositionInILSType{Name: %s}).Stage(stage.AixmdatatypesStage)", codepositioninilstypeIdent, __gong__toRawStringLiteral(codepositioninilstype.Name)))
			if lastStageVal != "AixmdatatypesStage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "AixmdatatypesStage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", codepositioninilstypeIdent, __gong__toRawStringLiteral(codepositioninilstype.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.NilReason = %s", codepositioninilstypeIdent, __gong__toRawStringLiteral(codepositioninilstype.NilReason)))
			values.WriteString(fmt.Sprintf("\n\t%s.EnclosedText = %s", codepositioninilstypeIdent, __gong__toRawStringLiteral(codepositioninilstype.EnclosedText)))
		}
	}
	if stage.AixmdatatypesStage != nil {
		for _, codeprimaryradartype := range __gong__sortStageSetInstances(stage.AixmdatatypesStage.CodePrimaryRadarTypes, stage.AixmdatatypesStage.CodePrimaryRadarType_stagedOrder) {
			if lastStageDecl != "AixmdatatypesStage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "AixmdatatypesStage"
			}
			codeprimaryradartypeIdent := codeprimaryradartype.GongGetIdentifier(stage.AixmdatatypesStage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&aixmdatatypes.CodePrimaryRadarType{Name: %s}).Stage(stage.AixmdatatypesStage)", codeprimaryradartypeIdent, __gong__toRawStringLiteral(codeprimaryradartype.Name)))
			if lastStageVal != "AixmdatatypesStage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "AixmdatatypesStage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", codeprimaryradartypeIdent, __gong__toRawStringLiteral(codeprimaryradartype.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.NilReason = %s", codeprimaryradartypeIdent, __gong__toRawStringLiteral(codeprimaryradartype.NilReason)))
			values.WriteString(fmt.Sprintf("\n\t%s.EnclosedText = %s", codeprimaryradartypeIdent, __gong__toRawStringLiteral(codeprimaryradartype.EnclosedText)))
		}
	}
	if stage.AixmdatatypesStage != nil {
		for _, coderadarservicetype := range __gong__sortStageSetInstances(stage.AixmdatatypesStage.CodeRadarServiceTypes, stage.AixmdatatypesStage.CodeRadarServiceType_stagedOrder) {
			if lastStageDecl != "AixmdatatypesStage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "AixmdatatypesStage"
			}
			coderadarservicetypeIdent := coderadarservicetype.GongGetIdentifier(stage.AixmdatatypesStage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&aixmdatatypes.CodeRadarServiceType{Name: %s}).Stage(stage.AixmdatatypesStage)", coderadarservicetypeIdent, __gong__toRawStringLiteral(coderadarservicetype.Name)))
			if lastStageVal != "AixmdatatypesStage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "AixmdatatypesStage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", coderadarservicetypeIdent, __gong__toRawStringLiteral(coderadarservicetype.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.NilReason = %s", coderadarservicetypeIdent, __gong__toRawStringLiteral(coderadarservicetype.NilReason)))
			values.WriteString(fmt.Sprintf("\n\t%s.EnclosedText = %s", coderadarservicetypeIdent, __gong__toRawStringLiteral(coderadarservicetype.EnclosedText)))
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
		for _, coderadiosignaltype := range __gong__sortStageSetInstances(stage.AixmdatatypesStage.CodeRadioSignalTypes, stage.AixmdatatypesStage.CodeRadioSignalType_stagedOrder) {
			if lastStageDecl != "AixmdatatypesStage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "AixmdatatypesStage"
			}
			coderadiosignaltypeIdent := coderadiosignaltype.GongGetIdentifier(stage.AixmdatatypesStage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&aixmdatatypes.CodeRadioSignalType{Name: %s}).Stage(stage.AixmdatatypesStage)", coderadiosignaltypeIdent, __gong__toRawStringLiteral(coderadiosignaltype.Name)))
			if lastStageVal != "AixmdatatypesStage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "AixmdatatypesStage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", coderadiosignaltypeIdent, __gong__toRawStringLiteral(coderadiosignaltype.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.NilReason = %s", coderadiosignaltypeIdent, __gong__toRawStringLiteral(coderadiosignaltype.NilReason)))
			values.WriteString(fmt.Sprintf("\n\t%s.EnclosedText = %s", coderadiosignaltypeIdent, __gong__toRawStringLiteral(coderadiosignaltype.EnclosedText)))
		}
	}
	if stage.AixmdatatypesStage != nil {
		for _, codereflectortype := range __gong__sortStageSetInstances(stage.AixmdatatypesStage.CodeReflectorTypes, stage.AixmdatatypesStage.CodeReflectorType_stagedOrder) {
			if lastStageDecl != "AixmdatatypesStage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "AixmdatatypesStage"
			}
			codereflectortypeIdent := codereflectortype.GongGetIdentifier(stage.AixmdatatypesStage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&aixmdatatypes.CodeReflectorType{Name: %s}).Stage(stage.AixmdatatypesStage)", codereflectortypeIdent, __gong__toRawStringLiteral(codereflectortype.Name)))
			if lastStageVal != "AixmdatatypesStage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "AixmdatatypesStage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", codereflectortypeIdent, __gong__toRawStringLiteral(codereflectortype.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.NilReason = %s", codereflectortypeIdent, __gong__toRawStringLiteral(codereflectortype.NilReason)))
			values.WriteString(fmt.Sprintf("\n\t%s.EnclosedText = %s", codereflectortypeIdent, __gong__toRawStringLiteral(codereflectortype.EnclosedText)))
		}
	}
	if stage.AixmdatatypesStage != nil {
		for _, codesignalperformanceilstype := range __gong__sortStageSetInstances(stage.AixmdatatypesStage.CodeSignalPerformanceILSTypes, stage.AixmdatatypesStage.CodeSignalPerformanceILSType_stagedOrder) {
			if lastStageDecl != "AixmdatatypesStage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "AixmdatatypesStage"
			}
			codesignalperformanceilstypeIdent := codesignalperformanceilstype.GongGetIdentifier(stage.AixmdatatypesStage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&aixmdatatypes.CodeSignalPerformanceILSType{Name: %s}).Stage(stage.AixmdatatypesStage)", codesignalperformanceilstypeIdent, __gong__toRawStringLiteral(codesignalperformanceilstype.Name)))
			if lastStageVal != "AixmdatatypesStage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "AixmdatatypesStage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", codesignalperformanceilstypeIdent, __gong__toRawStringLiteral(codesignalperformanceilstype.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.NilReason = %s", codesignalperformanceilstypeIdent, __gong__toRawStringLiteral(codesignalperformanceilstype.NilReason)))
			values.WriteString(fmt.Sprintf("\n\t%s.EnclosedText = %s", codesignalperformanceilstypeIdent, __gong__toRawStringLiteral(codesignalperformanceilstype.EnclosedText)))
		}
	}
	if stage.AixmdatatypesStage != nil {
		for _, codespecialnavigationchaindesignatortype := range __gong__sortStageSetInstances(stage.AixmdatatypesStage.CodeSpecialNavigationChainDesignatorTypes, stage.AixmdatatypesStage.CodeSpecialNavigationChainDesignatorType_stagedOrder) {
			if lastStageDecl != "AixmdatatypesStage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "AixmdatatypesStage"
			}
			codespecialnavigationchaindesignatortypeIdent := codespecialnavigationchaindesignatortype.GongGetIdentifier(stage.AixmdatatypesStage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&aixmdatatypes.CodeSpecialNavigationChainDesignatorType{Name: %s}).Stage(stage.AixmdatatypesStage)", codespecialnavigationchaindesignatortypeIdent, __gong__toRawStringLiteral(codespecialnavigationchaindesignatortype.Name)))
			if lastStageVal != "AixmdatatypesStage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "AixmdatatypesStage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", codespecialnavigationchaindesignatortypeIdent, __gong__toRawStringLiteral(codespecialnavigationchaindesignatortype.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.NilReason = %s", codespecialnavigationchaindesignatortypeIdent, __gong__toRawStringLiteral(codespecialnavigationchaindesignatortype.NilReason)))
			values.WriteString(fmt.Sprintf("\n\t%s.EnclosedText = %s", codespecialnavigationchaindesignatortypeIdent, __gong__toRawStringLiteral(codespecialnavigationchaindesignatortype.EnclosedText)))
		}
	}
	if stage.AixmdatatypesStage != nil {
		for _, codespecialnavigationstationtype := range __gong__sortStageSetInstances(stage.AixmdatatypesStage.CodeSpecialNavigationStationTypes, stage.AixmdatatypesStage.CodeSpecialNavigationStationType_stagedOrder) {
			if lastStageDecl != "AixmdatatypesStage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "AixmdatatypesStage"
			}
			codespecialnavigationstationtypeIdent := codespecialnavigationstationtype.GongGetIdentifier(stage.AixmdatatypesStage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&aixmdatatypes.CodeSpecialNavigationStationType{Name: %s}).Stage(stage.AixmdatatypesStage)", codespecialnavigationstationtypeIdent, __gong__toRawStringLiteral(codespecialnavigationstationtype.Name)))
			if lastStageVal != "AixmdatatypesStage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "AixmdatatypesStage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", codespecialnavigationstationtypeIdent, __gong__toRawStringLiteral(codespecialnavigationstationtype.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.NilReason = %s", codespecialnavigationstationtypeIdent, __gong__toRawStringLiteral(codespecialnavigationstationtype.NilReason)))
			values.WriteString(fmt.Sprintf("\n\t%s.EnclosedText = %s", codespecialnavigationstationtypeIdent, __gong__toRawStringLiteral(codespecialnavigationstationtype.EnclosedText)))
		}
	}
	if stage.AixmdatatypesStage != nil {
		for _, codespecialnavigationsystemtype := range __gong__sortStageSetInstances(stage.AixmdatatypesStage.CodeSpecialNavigationSystemTypes, stage.AixmdatatypesStage.CodeSpecialNavigationSystemType_stagedOrder) {
			if lastStageDecl != "AixmdatatypesStage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "AixmdatatypesStage"
			}
			codespecialnavigationsystemtypeIdent := codespecialnavigationsystemtype.GongGetIdentifier(stage.AixmdatatypesStage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&aixmdatatypes.CodeSpecialNavigationSystemType{Name: %s}).Stage(stage.AixmdatatypesStage)", codespecialnavigationsystemtypeIdent, __gong__toRawStringLiteral(codespecialnavigationsystemtype.Name)))
			if lastStageVal != "AixmdatatypesStage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "AixmdatatypesStage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", codespecialnavigationsystemtypeIdent, __gong__toRawStringLiteral(codespecialnavigationsystemtype.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.NilReason = %s", codespecialnavigationsystemtypeIdent, __gong__toRawStringLiteral(codespecialnavigationsystemtype.NilReason)))
			values.WriteString(fmt.Sprintf("\n\t%s.EnclosedText = %s", codespecialnavigationsystemtypeIdent, __gong__toRawStringLiteral(codespecialnavigationsystemtype.EnclosedText)))
		}
	}
	if stage.AixmdatatypesStage != nil {
		for _, codestandbypowertype := range __gong__sortStageSetInstances(stage.AixmdatatypesStage.CodeStandbyPowerTypes, stage.AixmdatatypesStage.CodeStandbyPowerType_stagedOrder) {
			if lastStageDecl != "AixmdatatypesStage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "AixmdatatypesStage"
			}
			codestandbypowertypeIdent := codestandbypowertype.GongGetIdentifier(stage.AixmdatatypesStage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&aixmdatatypes.CodeStandbyPowerType{Name: %s}).Stage(stage.AixmdatatypesStage)", codestandbypowertypeIdent, __gong__toRawStringLiteral(codestandbypowertype.Name)))
			if lastStageVal != "AixmdatatypesStage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "AixmdatatypesStage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", codestandbypowertypeIdent, __gong__toRawStringLiteral(codestandbypowertype.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.NilReason = %s", codestandbypowertypeIdent, __gong__toRawStringLiteral(codestandbypowertype.NilReason)))
			values.WriteString(fmt.Sprintf("\n\t%s.EnclosedText = %s", codestandbypowertypeIdent, __gong__toRawStringLiteral(codestandbypowertype.EnclosedText)))
		}
	}
	if stage.AixmdatatypesStage != nil {
		for _, codestatusnavaidtype := range __gong__sortStageSetInstances(stage.AixmdatatypesStage.CodeStatusNavaidTypes, stage.AixmdatatypesStage.CodeStatusNavaidType_stagedOrder) {
			if lastStageDecl != "AixmdatatypesStage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "AixmdatatypesStage"
			}
			codestatusnavaidtypeIdent := codestatusnavaidtype.GongGetIdentifier(stage.AixmdatatypesStage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&aixmdatatypes.CodeStatusNavaidType{Name: %s}).Stage(stage.AixmdatatypesStage)", codestatusnavaidtypeIdent, __gong__toRawStringLiteral(codestatusnavaidtype.Name)))
			if lastStageVal != "AixmdatatypesStage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "AixmdatatypesStage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", codestatusnavaidtypeIdent, __gong__toRawStringLiteral(codestatusnavaidtype.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.NilReason = %s", codestatusnavaidtypeIdent, __gong__toRawStringLiteral(codestatusnavaidtype.NilReason)))
			values.WriteString(fmt.Sprintf("\n\t%s.EnclosedText = %s", codestatusnavaidtypeIdent, __gong__toRawStringLiteral(codestatusnavaidtype.EnclosedText)))
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
		for _, codevortype := range __gong__sortStageSetInstances(stage.AixmdatatypesStage.CodeVORTypes, stage.AixmdatatypesStage.CodeVORType_stagedOrder) {
			if lastStageDecl != "AixmdatatypesStage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "AixmdatatypesStage"
			}
			codevortypeIdent := codevortype.GongGetIdentifier(stage.AixmdatatypesStage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&aixmdatatypes.CodeVORType{Name: %s}).Stage(stage.AixmdatatypesStage)", codevortypeIdent, __gong__toRawStringLiteral(codevortype.Name)))
			if lastStageVal != "AixmdatatypesStage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "AixmdatatypesStage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", codevortypeIdent, __gong__toRawStringLiteral(codevortype.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.NilReason = %s", codevortypeIdent, __gong__toRawStringLiteral(codevortype.NilReason)))
			values.WriteString(fmt.Sprintf("\n\t%s.EnclosedText = %s", codevortypeIdent, __gong__toRawStringLiteral(codevortype.EnclosedText)))
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
	if stage.AixmdatatypesStage != nil {
		for _, valmagneticvariationtype := range __gong__sortStageSetInstances(stage.AixmdatatypesStage.ValMagneticVariationTypes, stage.AixmdatatypesStage.ValMagneticVariationType_stagedOrder) {
			if lastStageDecl != "AixmdatatypesStage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "AixmdatatypesStage"
			}
			valmagneticvariationtypeIdent := valmagneticvariationtype.GongGetIdentifier(stage.AixmdatatypesStage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&aixmdatatypes.ValMagneticVariationType{Name: %s}).Stage(stage.AixmdatatypesStage)", valmagneticvariationtypeIdent, __gong__toRawStringLiteral(valmagneticvariationtype.Name)))
			if lastStageVal != "AixmdatatypesStage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "AixmdatatypesStage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", valmagneticvariationtypeIdent, __gong__toRawStringLiteral(valmagneticvariationtype.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.NilReason = %s", valmagneticvariationtypeIdent, __gong__toRawStringLiteral(valmagneticvariationtype.NilReason)))
			values.WriteString(fmt.Sprintf("\n\t%s.EnclosedText = %s", valmagneticvariationtypeIdent, __gong__toRawStringLiteral(valmagneticvariationtype.EnclosedText)))
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
			for _, elem := range a_extension.AbstractElevatedPointExtension {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := elem.GongGetIdentifier(stage.AixmabstractgmlobjecttypesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.AbstractElevatedPointExtension = append(%s.AbstractElevatedPointExtension, %s)", a_extensionIdent, a_extensionIdent, targetIdent))
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
			for _, elem := range a_extension_1.AbstractAuthorityForNavaidEquipmentExtension {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := elem.GongGetIdentifier(stage.AixmabstractgmlobjecttypesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.AbstractAuthorityForNavaidEquipmentExtension = append(%s.AbstractAuthorityForNavaidEquipmentExtension, %s)", a_extension_1Ident, a_extension_1Ident, targetIdent))
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
			for _, elem := range a_extension_10.AbstractGlidepathExtension {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := elem.GongGetIdentifier(stage.AixmabstractgmlobjecttypesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.AbstractGlidepathExtension = append(%s.AbstractGlidepathExtension, %s)", a_extension_10Ident, a_extension_10Ident, targetIdent))
			}
			for _, elem := range a_extension_10.AbstractNavaidEquipmentExtension {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := elem.GongGetIdentifier(stage.AixmabstractgmlobjecttypesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.AbstractNavaidEquipmentExtension = append(%s.AbstractNavaidEquipmentExtension, %s)", a_extension_10Ident, a_extension_10Ident, targetIdent))
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
			for _, elem := range a_extension_11.AbstractLocalizerExtension {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := elem.GongGetIdentifier(stage.AixmabstractgmlobjecttypesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.AbstractLocalizerExtension = append(%s.AbstractLocalizerExtension, %s)", a_extension_11Ident, a_extension_11Ident, targetIdent))
			}
			for _, elem := range a_extension_11.AbstractNavaidEquipmentExtension {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := elem.GongGetIdentifier(stage.AixmabstractgmlobjecttypesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.AbstractNavaidEquipmentExtension = append(%s.AbstractNavaidEquipmentExtension, %s)", a_extension_11Ident, a_extension_11Ident, targetIdent))
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
			for _, elem := range a_extension_12.AbstractMarkerBeaconExtension {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := elem.GongGetIdentifier(stage.AixmabstractgmlobjecttypesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.AbstractMarkerBeaconExtension = append(%s.AbstractMarkerBeaconExtension, %s)", a_extension_12Ident, a_extension_12Ident, targetIdent))
			}
			for _, elem := range a_extension_12.AbstractNavaidEquipmentExtension {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := elem.GongGetIdentifier(stage.AixmabstractgmlobjecttypesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.AbstractNavaidEquipmentExtension = append(%s.AbstractNavaidEquipmentExtension, %s)", a_extension_12Ident, a_extension_12Ident, targetIdent))
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
			for _, elem := range a_extension_13.AbstractNavaidExtension {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := elem.GongGetIdentifier(stage.AixmabstractgmlobjecttypesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.AbstractNavaidExtension = append(%s.AbstractNavaidExtension, %s)", a_extension_13Ident, a_extension_13Ident, targetIdent))
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
			for _, elem := range a_extension_14.AbstractNavaidComponentExtension {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := elem.GongGetIdentifier(stage.AixmabstractgmlobjecttypesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.AbstractNavaidComponentExtension = append(%s.AbstractNavaidComponentExtension, %s)", a_extension_14Ident, a_extension_14Ident, targetIdent))
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
			for _, elem := range a_extension_15.AbstractNavaidEquipmentMonitoringExtension {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := elem.GongGetIdentifier(stage.AixmabstractgmlobjecttypesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.AbstractNavaidEquipmentMonitoringExtension = append(%s.AbstractNavaidEquipmentMonitoringExtension, %s)", a_extension_15Ident, a_extension_15Ident, targetIdent))
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
			for _, elem := range a_extension_16.AbstractPropertiesWithScheduleExtension {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := elem.GongGetIdentifier(stage.AixmabstractgmlobjecttypesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.AbstractPropertiesWithScheduleExtension = append(%s.AbstractPropertiesWithScheduleExtension, %s)", a_extension_16Ident, a_extension_16Ident, targetIdent))
			}
			for _, elem := range a_extension_16.AbstractNavaidOperationalStatusExtension {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := elem.GongGetIdentifier(stage.AixmabstractgmlobjecttypesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.AbstractNavaidOperationalStatusExtension = append(%s.AbstractNavaidOperationalStatusExtension, %s)", a_extension_16Ident, a_extension_16Ident, targetIdent))
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
			for _, elem := range a_extension_17.AbstractNDBExtension {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := elem.GongGetIdentifier(stage.AixmabstractgmlobjecttypesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.AbstractNDBExtension = append(%s.AbstractNDBExtension, %s)", a_extension_17Ident, a_extension_17Ident, targetIdent))
			}
			for _, elem := range a_extension_17.AbstractNavaidEquipmentExtension {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := elem.GongGetIdentifier(stage.AixmabstractgmlobjecttypesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.AbstractNavaidEquipmentExtension = append(%s.AbstractNavaidEquipmentExtension, %s)", a_extension_17Ident, a_extension_17Ident, targetIdent))
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
			for _, elem := range a_extension_18.AbstractSDFExtension {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := elem.GongGetIdentifier(stage.AixmabstractgmlobjecttypesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.AbstractSDFExtension = append(%s.AbstractSDFExtension, %s)", a_extension_18Ident, a_extension_18Ident, targetIdent))
			}
			for _, elem := range a_extension_18.AbstractNavaidEquipmentExtension {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := elem.GongGetIdentifier(stage.AixmabstractgmlobjecttypesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.AbstractNavaidEquipmentExtension = append(%s.AbstractNavaidEquipmentExtension, %s)", a_extension_18Ident, a_extension_18Ident, targetIdent))
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
			for _, elem := range a_extension_19.AbstractSpecialNavigationStationExtension {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := elem.GongGetIdentifier(stage.AixmabstractgmlobjecttypesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.AbstractSpecialNavigationStationExtension = append(%s.AbstractSpecialNavigationStationExtension, %s)", a_extension_19Ident, a_extension_19Ident, targetIdent))
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
			for _, elem := range a_extension_2.AbstractAuthorityForSpecialNavigationStationExtension {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := elem.GongGetIdentifier(stage.AixmabstractgmlobjecttypesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.AbstractAuthorityForSpecialNavigationStationExtension = append(%s.AbstractAuthorityForSpecialNavigationStationExtension, %s)", a_extension_2Ident, a_extension_2Ident, targetIdent))
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
			for _, elem := range a_extension_20.AbstractPropertiesWithScheduleExtension {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := elem.GongGetIdentifier(stage.AixmabstractgmlobjecttypesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.AbstractPropertiesWithScheduleExtension = append(%s.AbstractPropertiesWithScheduleExtension, %s)", a_extension_20Ident, a_extension_20Ident, targetIdent))
			}
			for _, elem := range a_extension_20.AbstractSpecialNavigationStationStatusExtension {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := elem.GongGetIdentifier(stage.AixmabstractgmlobjecttypesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.AbstractSpecialNavigationStationStatusExtension = append(%s.AbstractSpecialNavigationStationStatusExtension, %s)", a_extension_20Ident, a_extension_20Ident, targetIdent))
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
			for _, elem := range a_extension_21.AbstractSpecialNavigationSystemExtension {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := elem.GongGetIdentifier(stage.AixmabstractgmlobjecttypesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.AbstractSpecialNavigationSystemExtension = append(%s.AbstractSpecialNavigationSystemExtension, %s)", a_extension_21Ident, a_extension_21Ident, targetIdent))
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
			for _, elem := range a_extension_22.AbstractTACANExtension {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := elem.GongGetIdentifier(stage.AixmabstractgmlobjecttypesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.AbstractTACANExtension = append(%s.AbstractTACANExtension, %s)", a_extension_22Ident, a_extension_22Ident, targetIdent))
			}
			for _, elem := range a_extension_22.AbstractNavaidEquipmentExtension {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := elem.GongGetIdentifier(stage.AixmabstractgmlobjecttypesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.AbstractNavaidEquipmentExtension = append(%s.AbstractNavaidEquipmentExtension, %s)", a_extension_22Ident, a_extension_22Ident, targetIdent))
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
			for _, elem := range a_extension_23.AbstractVORExtension {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := elem.GongGetIdentifier(stage.AixmabstractgmlobjecttypesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.AbstractVORExtension = append(%s.AbstractVORExtension, %s)", a_extension_23Ident, a_extension_23Ident, targetIdent))
			}
			for _, elem := range a_extension_23.AbstractNavaidEquipmentExtension {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := elem.GongGetIdentifier(stage.AixmabstractgmlobjecttypesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.AbstractNavaidEquipmentExtension = append(%s.AbstractNavaidEquipmentExtension, %s)", a_extension_23Ident, a_extension_23Ident, targetIdent))
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
			for _, elem := range a_extension_24.AbstractAngleIndicationExtension {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := elem.GongGetIdentifier(stage.AixmabstractgmlobjecttypesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.AbstractAngleIndicationExtension = append(%s.AbstractAngleIndicationExtension, %s)", a_extension_24Ident, a_extension_24Ident, targetIdent))
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
			for _, elem := range a_extension_25.AbstractDistanceIndicationExtension {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := elem.GongGetIdentifier(stage.AixmabstractgmlobjecttypesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.AbstractDistanceIndicationExtension = append(%s.AbstractDistanceIndicationExtension, %s)", a_extension_25Ident, a_extension_25Ident, targetIdent))
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
			for _, elem := range a_extension_26.AbstractDesignatedPointExtension {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := elem.GongGetIdentifier(stage.AixmabstractgmlobjecttypesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.AbstractDesignatedPointExtension = append(%s.AbstractDesignatedPointExtension, %s)", a_extension_26Ident, a_extension_26Ident, targetIdent))
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
			for _, elem := range a_extension_27.AbstractSignificantPointInAirspaceExtension {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := elem.GongGetIdentifier(stage.AixmabstractgmlobjecttypesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.AbstractSignificantPointInAirspaceExtension = append(%s.AbstractSignificantPointInAirspaceExtension, %s)", a_extension_27Ident, a_extension_27Ident, targetIdent))
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
			for _, elem := range a_extension_28.AbstractLinguisticNoteExtension {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := elem.GongGetIdentifier(stage.AixmabstractgmlobjecttypesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.AbstractLinguisticNoteExtension = append(%s.AbstractLinguisticNoteExtension, %s)", a_extension_28Ident, a_extension_28Ident, targetIdent))
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
			for _, elem := range a_extension_29.AbstractNoteExtension {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := elem.GongGetIdentifier(stage.AixmabstractgmlobjecttypesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.AbstractNoteExtension = append(%s.AbstractNoteExtension, %s)", a_extension_29Ident, a_extension_29Ident, targetIdent))
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
			for _, elem := range a_extension_3.AbstractAuthorityForSpecialNavigationSystemExtension {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := elem.GongGetIdentifier(stage.AixmabstractgmlobjecttypesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.AbstractAuthorityForSpecialNavigationSystemExtension = append(%s.AbstractAuthorityForSpecialNavigationSystemExtension, %s)", a_extension_3Ident, a_extension_3Ident, targetIdent))
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
			for _, elem := range a_extension_35.AbstractPrecisionApproachRadarExtension {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := elem.GongGetIdentifier(stage.AixmabstractgmlobjecttypesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.AbstractPrecisionApproachRadarExtension = append(%s.AbstractPrecisionApproachRadarExtension, %s)", a_extension_35Ident, a_extension_35Ident, targetIdent))
			}
			for _, elem := range a_extension_35.AbstractRadarEquipmentExtension {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := elem.GongGetIdentifier(stage.AixmabstractgmlobjecttypesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.AbstractRadarEquipmentExtension = append(%s.AbstractRadarEquipmentExtension, %s)", a_extension_35Ident, a_extension_35Ident, targetIdent))
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
			for _, elem := range a_extension_36.AbstractPrimarySurveillanceRadarExtension {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := elem.GongGetIdentifier(stage.AixmabstractgmlobjecttypesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.AbstractPrimarySurveillanceRadarExtension = append(%s.AbstractPrimarySurveillanceRadarExtension, %s)", a_extension_36Ident, a_extension_36Ident, targetIdent))
			}
			for _, elem := range a_extension_36.AbstractSurveillanceRadarExtension {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := elem.GongGetIdentifier(stage.AixmabstractgmlobjecttypesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.AbstractSurveillanceRadarExtension = append(%s.AbstractSurveillanceRadarExtension, %s)", a_extension_36Ident, a_extension_36Ident, targetIdent))
			}
			for _, elem := range a_extension_36.AbstractRadarEquipmentExtension {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := elem.GongGetIdentifier(stage.AixmabstractgmlobjecttypesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.AbstractRadarEquipmentExtension = append(%s.AbstractRadarEquipmentExtension, %s)", a_extension_36Ident, a_extension_36Ident, targetIdent))
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
			for _, elem := range a_extension_37.AbstractRadarComponentExtension {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := elem.GongGetIdentifier(stage.AixmabstractgmlobjecttypesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.AbstractRadarComponentExtension = append(%s.AbstractRadarComponentExtension, %s)", a_extension_37Ident, a_extension_37Ident, targetIdent))
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
			for _, elem := range a_extension_38.AbstractRadarSystemExtension {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := elem.GongGetIdentifier(stage.AixmabstractgmlobjecttypesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.AbstractRadarSystemExtension = append(%s.AbstractRadarSystemExtension, %s)", a_extension_38Ident, a_extension_38Ident, targetIdent))
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
			for _, elem := range a_extension_39.AbstractReflectorExtension {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := elem.GongGetIdentifier(stage.AixmabstractgmlobjecttypesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.AbstractReflectorExtension = append(%s.AbstractReflectorExtension, %s)", a_extension_39Ident, a_extension_39Ident, targetIdent))
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
			for _, elem := range a_extension_4.AbstractAzimuthExtension {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := elem.GongGetIdentifier(stage.AixmabstractgmlobjecttypesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.AbstractAzimuthExtension = append(%s.AbstractAzimuthExtension, %s)", a_extension_4Ident, a_extension_4Ident, targetIdent))
			}
			for _, elem := range a_extension_4.AbstractNavaidEquipmentExtension {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := elem.GongGetIdentifier(stage.AixmabstractgmlobjecttypesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.AbstractNavaidEquipmentExtension = append(%s.AbstractNavaidEquipmentExtension, %s)", a_extension_4Ident, a_extension_4Ident, targetIdent))
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
			for _, elem := range a_extension_40.AbstractSecondarySurveillanceRadarExtension {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := elem.GongGetIdentifier(stage.AixmabstractgmlobjecttypesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.AbstractSecondarySurveillanceRadarExtension = append(%s.AbstractSecondarySurveillanceRadarExtension, %s)", a_extension_40Ident, a_extension_40Ident, targetIdent))
			}
			for _, elem := range a_extension_40.AbstractSurveillanceRadarExtension {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := elem.GongGetIdentifier(stage.AixmabstractgmlobjecttypesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.AbstractSurveillanceRadarExtension = append(%s.AbstractSurveillanceRadarExtension, %s)", a_extension_40Ident, a_extension_40Ident, targetIdent))
			}
			for _, elem := range a_extension_40.AbstractRadarEquipmentExtension {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := elem.GongGetIdentifier(stage.AixmabstractgmlobjecttypesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.AbstractRadarEquipmentExtension = append(%s.AbstractRadarEquipmentExtension, %s)", a_extension_40Ident, a_extension_40Ident, targetIdent))
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
			for _, elem := range a_extension_41.AbstractSurveillanceGroundStationExtension {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := elem.GongGetIdentifier(stage.AixmabstractgmlobjecttypesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.AbstractSurveillanceGroundStationExtension = append(%s.AbstractSurveillanceGroundStationExtension, %s)", a_extension_41Ident, a_extension_41Ident, targetIdent))
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
			for _, elem := range a_extension_5.AbstractCheckpointINSExtension {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := elem.GongGetIdentifier(stage.AixmabstractgmlobjecttypesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.AbstractCheckpointINSExtension = append(%s.AbstractCheckpointINSExtension, %s)", a_extension_5Ident, a_extension_5Ident, targetIdent))
			}
			for _, elem := range a_extension_5.AbstractNavigationSystemCheckpointExtension {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := elem.GongGetIdentifier(stage.AixmabstractgmlobjecttypesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.AbstractNavigationSystemCheckpointExtension = append(%s.AbstractNavigationSystemCheckpointExtension, %s)", a_extension_5Ident, a_extension_5Ident, targetIdent))
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
			for _, elem := range a_extension_6.AbstractCheckpointVORExtension {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := elem.GongGetIdentifier(stage.AixmabstractgmlobjecttypesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.AbstractCheckpointVORExtension = append(%s.AbstractCheckpointVORExtension, %s)", a_extension_6Ident, a_extension_6Ident, targetIdent))
			}
			for _, elem := range a_extension_6.AbstractNavigationSystemCheckpointExtension {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := elem.GongGetIdentifier(stage.AixmabstractgmlobjecttypesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.AbstractNavigationSystemCheckpointExtension = append(%s.AbstractNavigationSystemCheckpointExtension, %s)", a_extension_6Ident, a_extension_6Ident, targetIdent))
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
			for _, elem := range a_extension_7.AbstractDirectionFinderExtension {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := elem.GongGetIdentifier(stage.AixmabstractgmlobjecttypesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.AbstractDirectionFinderExtension = append(%s.AbstractDirectionFinderExtension, %s)", a_extension_7Ident, a_extension_7Ident, targetIdent))
			}
			for _, elem := range a_extension_7.AbstractNavaidEquipmentExtension {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := elem.GongGetIdentifier(stage.AixmabstractgmlobjecttypesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.AbstractNavaidEquipmentExtension = append(%s.AbstractNavaidEquipmentExtension, %s)", a_extension_7Ident, a_extension_7Ident, targetIdent))
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
			for _, elem := range a_extension_8.AbstractDMEExtension {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := elem.GongGetIdentifier(stage.AixmabstractgmlobjecttypesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.AbstractDMEExtension = append(%s.AbstractDMEExtension, %s)", a_extension_8Ident, a_extension_8Ident, targetIdent))
			}
			for _, elem := range a_extension_8.AbstractNavaidEquipmentExtension {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := elem.GongGetIdentifier(stage.AixmabstractgmlobjecttypesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.AbstractNavaidEquipmentExtension = append(%s.AbstractNavaidEquipmentExtension, %s)", a_extension_8Ident, a_extension_8Ident, targetIdent))
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
			for _, elem := range a_extension_9.AbstractElevationExtension {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := elem.GongGetIdentifier(stage.AixmabstractgmlobjecttypesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.AbstractElevationExtension = append(%s.AbstractElevationExtension, %s)", a_extension_9Ident, a_extension_9Ident, targetIdent))
			}
			for _, elem := range a_extension_9.AbstractNavaidEquipmentExtension {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := elem.GongGetIdentifier(stage.AixmabstractgmlobjecttypesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.AbstractNavaidEquipmentExtension = append(%s.AbstractNavaidEquipmentExtension, %s)", a_extension_9Ident, a_extension_9Ident, targetIdent))
			}
		}
	}
	if stage.AixmfeaturesStage != nil {
		for _, abstractnavaidequipmenttype := range __gong__sortStageSetInstances(stage.AixmfeaturesStage.AbstractNavaidEquipmentTypes, stage.AixmfeaturesStage.AbstractNavaidEquipmentType_stagedOrder) {
			if lastStageDecl != "AixmfeaturesStage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "AixmfeaturesStage"
			}
			abstractnavaidequipmenttypeIdent := abstractnavaidequipmenttype.GongGetIdentifier(stage.AixmfeaturesStage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&aixmfeatures.AbstractNavaidEquipmentType{Name: %s}).Stage(stage.AixmfeaturesStage)", abstractnavaidequipmenttypeIdent, __gong__toRawStringLiteral(abstractnavaidequipmenttype.Name)))
			if lastStageVal != "AixmfeaturesStage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "AixmfeaturesStage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", abstractnavaidequipmenttypeIdent, __gong__toRawStringLiteral(abstractnavaidequipmenttype.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Id = %s", abstractnavaidequipmenttypeIdent, __gong__toRawStringLiteral(abstractnavaidequipmenttype.Id)))
			values.WriteString(fmt.Sprintf("\n\t%s.Description = %s", abstractnavaidequipmenttypeIdent, __gong__toRawStringLiteral(abstractnavaidequipmenttype.Description)))
			if abstractnavaidequipmenttype.Identifier != nil {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := abstractnavaidequipmenttype.Identifier.GongGetIdentifier(stage.GmlStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Identifier = %s", abstractnavaidequipmenttypeIdent, targetIdent))
			}
			for _, elem := range abstractnavaidequipmenttype.NameXSD {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := elem.GongGetIdentifier(stage.GmlStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.NameXSD = append(%s.NameXSD, %s)", abstractnavaidequipmenttypeIdent, abstractnavaidequipmenttypeIdent, targetIdent))
			}
			if abstractnavaidequipmenttype.BoundedBy != nil {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := abstractnavaidequipmenttype.BoundedBy.GongGetIdentifier(stage.GmlStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.BoundedBy = %s", abstractnavaidequipmenttypeIdent, targetIdent))
			}
			if abstractnavaidequipmenttype.FeatureMetadata != nil {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := abstractnavaidequipmenttype.FeatureMetadata.GongGetIdentifier(stage.AixmabstractgmlobjecttypesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.FeatureMetadata = %s", abstractnavaidequipmenttypeIdent, targetIdent))
			}
		}
	}
	if stage.AixmfeaturesStage != nil {
		for _, abstractnavigationsystemcheckpointtype := range __gong__sortStageSetInstances(stage.AixmfeaturesStage.AbstractNavigationSystemCheckpointTypes, stage.AixmfeaturesStage.AbstractNavigationSystemCheckpointType_stagedOrder) {
			if lastStageDecl != "AixmfeaturesStage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "AixmfeaturesStage"
			}
			abstractnavigationsystemcheckpointtypeIdent := abstractnavigationsystemcheckpointtype.GongGetIdentifier(stage.AixmfeaturesStage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&aixmfeatures.AbstractNavigationSystemCheckpointType{Name: %s}).Stage(stage.AixmfeaturesStage)", abstractnavigationsystemcheckpointtypeIdent, __gong__toRawStringLiteral(abstractnavigationsystemcheckpointtype.Name)))
			if lastStageVal != "AixmfeaturesStage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "AixmfeaturesStage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", abstractnavigationsystemcheckpointtypeIdent, __gong__toRawStringLiteral(abstractnavigationsystemcheckpointtype.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Id = %s", abstractnavigationsystemcheckpointtypeIdent, __gong__toRawStringLiteral(abstractnavigationsystemcheckpointtype.Id)))
			values.WriteString(fmt.Sprintf("\n\t%s.Description = %s", abstractnavigationsystemcheckpointtypeIdent, __gong__toRawStringLiteral(abstractnavigationsystemcheckpointtype.Description)))
			if abstractnavigationsystemcheckpointtype.Identifier != nil {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := abstractnavigationsystemcheckpointtype.Identifier.GongGetIdentifier(stage.GmlStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Identifier = %s", abstractnavigationsystemcheckpointtypeIdent, targetIdent))
			}
			for _, elem := range abstractnavigationsystemcheckpointtype.NameXSD {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := elem.GongGetIdentifier(stage.GmlStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.NameXSD = append(%s.NameXSD, %s)", abstractnavigationsystemcheckpointtypeIdent, abstractnavigationsystemcheckpointtypeIdent, targetIdent))
			}
			if abstractnavigationsystemcheckpointtype.BoundedBy != nil {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := abstractnavigationsystemcheckpointtype.BoundedBy.GongGetIdentifier(stage.GmlStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.BoundedBy = %s", abstractnavigationsystemcheckpointtypeIdent, targetIdent))
			}
			if abstractnavigationsystemcheckpointtype.FeatureMetadata != nil {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := abstractnavigationsystemcheckpointtype.FeatureMetadata.GongGetIdentifier(stage.AixmabstractgmlobjecttypesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.FeatureMetadata = %s", abstractnavigationsystemcheckpointtypeIdent, targetIdent))
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
		for _, abstractradarequipmenttype := range __gong__sortStageSetInstances(stage.AixmfeaturesStage.AbstractRadarEquipmentTypes, stage.AixmfeaturesStage.AbstractRadarEquipmentType_stagedOrder) {
			if lastStageDecl != "AixmfeaturesStage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "AixmfeaturesStage"
			}
			abstractradarequipmenttypeIdent := abstractradarequipmenttype.GongGetIdentifier(stage.AixmfeaturesStage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&aixmfeatures.AbstractRadarEquipmentType{Name: %s}).Stage(stage.AixmfeaturesStage)", abstractradarequipmenttypeIdent, __gong__toRawStringLiteral(abstractradarequipmenttype.Name)))
			if lastStageVal != "AixmfeaturesStage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "AixmfeaturesStage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", abstractradarequipmenttypeIdent, __gong__toRawStringLiteral(abstractradarequipmenttype.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Id = %s", abstractradarequipmenttypeIdent, __gong__toRawStringLiteral(abstractradarequipmenttype.Id)))
			values.WriteString(fmt.Sprintf("\n\t%s.Description = %s", abstractradarequipmenttypeIdent, __gong__toRawStringLiteral(abstractradarequipmenttype.Description)))
			if abstractradarequipmenttype.Identifier != nil {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := abstractradarequipmenttype.Identifier.GongGetIdentifier(stage.GmlStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Identifier = %s", abstractradarequipmenttypeIdent, targetIdent))
			}
			for _, elem := range abstractradarequipmenttype.NameXSD {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := elem.GongGetIdentifier(stage.GmlStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.NameXSD = append(%s.NameXSD, %s)", abstractradarequipmenttypeIdent, abstractradarequipmenttypeIdent, targetIdent))
			}
			if abstractradarequipmenttype.BoundedBy != nil {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := abstractradarequipmenttype.BoundedBy.GongGetIdentifier(stage.GmlStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.BoundedBy = %s", abstractradarequipmenttypeIdent, targetIdent))
			}
			if abstractradarequipmenttype.FeatureMetadata != nil {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := abstractradarequipmenttype.FeatureMetadata.GongGetIdentifier(stage.AixmabstractgmlobjecttypesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.FeatureMetadata = %s", abstractradarequipmenttypeIdent, targetIdent))
			}
		}
	}
	if stage.AixmfeaturesStage != nil {
		for _, abstractsurveillanceradartype := range __gong__sortStageSetInstances(stage.AixmfeaturesStage.AbstractSurveillanceRadarTypes, stage.AixmfeaturesStage.AbstractSurveillanceRadarType_stagedOrder) {
			if lastStageDecl != "AixmfeaturesStage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "AixmfeaturesStage"
			}
			abstractsurveillanceradartypeIdent := abstractsurveillanceradartype.GongGetIdentifier(stage.AixmfeaturesStage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&aixmfeatures.AbstractSurveillanceRadarType{Name: %s}).Stage(stage.AixmfeaturesStage)", abstractsurveillanceradartypeIdent, __gong__toRawStringLiteral(abstractsurveillanceradartype.Name)))
			if lastStageVal != "AixmfeaturesStage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "AixmfeaturesStage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", abstractsurveillanceradartypeIdent, __gong__toRawStringLiteral(abstractsurveillanceradartype.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Id = %s", abstractsurveillanceradartypeIdent, __gong__toRawStringLiteral(abstractsurveillanceradartype.Id)))
			values.WriteString(fmt.Sprintf("\n\t%s.Description = %s", abstractsurveillanceradartypeIdent, __gong__toRawStringLiteral(abstractsurveillanceradartype.Description)))
			if abstractsurveillanceradartype.Identifier != nil {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := abstractsurveillanceradartype.Identifier.GongGetIdentifier(stage.GmlStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Identifier = %s", abstractsurveillanceradartypeIdent, targetIdent))
			}
			for _, elem := range abstractsurveillanceradartype.NameXSD {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := elem.GongGetIdentifier(stage.GmlStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.NameXSD = append(%s.NameXSD, %s)", abstractsurveillanceradartypeIdent, abstractsurveillanceradartypeIdent, targetIdent))
			}
			if abstractsurveillanceradartype.BoundedBy != nil {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := abstractsurveillanceradartype.BoundedBy.GongGetIdentifier(stage.GmlStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.BoundedBy = %s", abstractsurveillanceradartypeIdent, targetIdent))
			}
			if abstractsurveillanceradartype.FeatureMetadata != nil {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := abstractsurveillanceradartype.FeatureMetadata.GongGetIdentifier(stage.AixmabstractgmlobjecttypesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.FeatureMetadata = %s", abstractsurveillanceradartypeIdent, targetIdent))
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
		for _, angleindicationtimeslicepropertytype := range __gong__sortStageSetInstances(stage.AixmfeaturesStage.AngleIndicationTimeSlicePropertyTypes, stage.AixmfeaturesStage.AngleIndicationTimeSlicePropertyType_stagedOrder) {
			if lastStageDecl != "AixmfeaturesStage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "AixmfeaturesStage"
			}
			angleindicationtimeslicepropertytypeIdent := angleindicationtimeslicepropertytype.GongGetIdentifier(stage.AixmfeaturesStage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&aixmfeatures.AngleIndicationTimeSlicePropertyType{Name: %s}).Stage(stage.AixmfeaturesStage)", angleindicationtimeslicepropertytypeIdent, __gong__toRawStringLiteral(angleindicationtimeslicepropertytype.Name)))
			if lastStageVal != "AixmfeaturesStage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "AixmfeaturesStage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", angleindicationtimeslicepropertytypeIdent, __gong__toRawStringLiteral(angleindicationtimeslicepropertytype.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Owns = %t", angleindicationtimeslicepropertytypeIdent, angleindicationtimeslicepropertytype.Owns))
			if angleindicationtimeslicepropertytype.AngleIndicationTimeSlice != nil {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := angleindicationtimeslicepropertytype.AngleIndicationTimeSlice.GongGetIdentifier(stage.AixmfeaturesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.AngleIndicationTimeSlice = %s", angleindicationtimeslicepropertytypeIdent, targetIdent))
			}
		}
	}
	if stage.AixmfeaturesStage != nil {
		for _, angleindicationtimeslicetype := range __gong__sortStageSetInstances(stage.AixmfeaturesStage.AngleIndicationTimeSliceTypes, stage.AixmfeaturesStage.AngleIndicationTimeSliceType_stagedOrder) {
			if lastStageDecl != "AixmfeaturesStage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "AixmfeaturesStage"
			}
			angleindicationtimeslicetypeIdent := angleindicationtimeslicetype.GongGetIdentifier(stage.AixmfeaturesStage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&aixmfeatures.AngleIndicationTimeSliceType{Name: %s}).Stage(stage.AixmfeaturesStage)", angleindicationtimeslicetypeIdent, __gong__toRawStringLiteral(angleindicationtimeslicetype.Name)))
			if lastStageVal != "AixmfeaturesStage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "AixmfeaturesStage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", angleindicationtimeslicetypeIdent, __gong__toRawStringLiteral(angleindicationtimeslicetype.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Id = %s", angleindicationtimeslicetypeIdent, __gong__toRawStringLiteral(angleindicationtimeslicetype.Id)))
			values.WriteString(fmt.Sprintf("\n\t%s.Description = %s", angleindicationtimeslicetypeIdent, __gong__toRawStringLiteral(angleindicationtimeslicetype.Description)))
			values.WriteString(fmt.Sprintf("\n\t%s.Interpretation = %s", angleindicationtimeslicetypeIdent, __gong__toRawStringLiteral(angleindicationtimeslicetype.Interpretation)))
			values.WriteString(fmt.Sprintf("\n\t%s.SequenceNumber = %d", angleindicationtimeslicetypeIdent, angleindicationtimeslicetype.SequenceNumber))
			values.WriteString(fmt.Sprintf("\n\t%s.CorrectionNumber = %d", angleindicationtimeslicetypeIdent, angleindicationtimeslicetype.CorrectionNumber))
			if angleindicationtimeslicetype.Identifier != nil {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := angleindicationtimeslicetype.Identifier.GongGetIdentifier(stage.GmlStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Identifier = %s", angleindicationtimeslicetypeIdent, targetIdent))
			}
			for _, elem := range angleindicationtimeslicetype.NameXSD {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := elem.GongGetIdentifier(stage.GmlStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.NameXSD = append(%s.NameXSD, %s)", angleindicationtimeslicetypeIdent, angleindicationtimeslicetypeIdent, targetIdent))
			}
			if angleindicationtimeslicetype.ValidTime != nil {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := angleindicationtimeslicetype.ValidTime.GongGetIdentifier(stage.GmlStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.ValidTime = %s", angleindicationtimeslicetypeIdent, targetIdent))
			}
			if angleindicationtimeslicetype.TimeSliceMetadata != nil {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := angleindicationtimeslicetype.TimeSliceMetadata.GongGetIdentifier(stage.AixmabstractgmlobjecttypesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.TimeSliceMetadata = %s", angleindicationtimeslicetypeIdent, targetIdent))
			}
			if angleindicationtimeslicetype.FeatureLifetime != nil {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := angleindicationtimeslicetype.FeatureLifetime.GongGetIdentifier(stage.GmlStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.FeatureLifetime = %s", angleindicationtimeslicetypeIdent, targetIdent))
			}
			if angleindicationtimeslicetype.Angle != nil {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := angleindicationtimeslicetype.Angle.GongGetIdentifier(stage.AixmdatatypesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Angle = %s", angleindicationtimeslicetypeIdent, targetIdent))
			}
			if angleindicationtimeslicetype.AngleType != nil {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := angleindicationtimeslicetype.AngleType.GongGetIdentifier(stage.AixmdatatypesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.AngleType = %s", angleindicationtimeslicetypeIdent, targetIdent))
			}
			if angleindicationtimeslicetype.IndicationDirection != nil {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := angleindicationtimeslicetype.IndicationDirection.GongGetIdentifier(stage.AixmdatatypesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.IndicationDirection = %s", angleindicationtimeslicetypeIdent, targetIdent))
			}
			if angleindicationtimeslicetype.TrueAngle != nil {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := angleindicationtimeslicetype.TrueAngle.GongGetIdentifier(stage.AixmdatatypesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.TrueAngle = %s", angleindicationtimeslicetypeIdent, targetIdent))
			}
			if angleindicationtimeslicetype.CardinalDirection != nil {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := angleindicationtimeslicetype.CardinalDirection.GongGetIdentifier(stage.AixmdatatypesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.CardinalDirection = %s", angleindicationtimeslicetypeIdent, targetIdent))
			}
			if angleindicationtimeslicetype.MinimumReceptionAltitude != nil {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := angleindicationtimeslicetype.MinimumReceptionAltitude.GongGetIdentifier(stage.AixmdatatypesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.MinimumReceptionAltitude = %s", angleindicationtimeslicetypeIdent, targetIdent))
			}
			if angleindicationtimeslicetype.Fix != nil {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := angleindicationtimeslicetype.Fix.GongGetIdentifier(stage.AixmfeaturesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Fix = %s", angleindicationtimeslicetypeIdent, targetIdent))
			}
			if angleindicationtimeslicetype.PointChoice_fixDesignatedPoint != nil {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := angleindicationtimeslicetype.PointChoice_fixDesignatedPoint.GongGetIdentifier(stage.AixmfeaturesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.PointChoice_fixDesignatedPoint = %s", angleindicationtimeslicetypeIdent, targetIdent))
			}
			if angleindicationtimeslicetype.PointChoice_navaidSystem != nil {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := angleindicationtimeslicetype.PointChoice_navaidSystem.GongGetIdentifier(stage.AixmfeaturesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.PointChoice_navaidSystem = %s", angleindicationtimeslicetypeIdent, targetIdent))
			}
			if angleindicationtimeslicetype.PointChoice_position != nil {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := angleindicationtimeslicetype.PointChoice_position.GongGetIdentifier(stage.AixmfeaturesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.PointChoice_position = %s", angleindicationtimeslicetypeIdent, targetIdent))
			}
			if angleindicationtimeslicetype.PointChoice_runwayPoint != nil {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := angleindicationtimeslicetype.PointChoice_runwayPoint.GongGetIdentifier(stage.AixmfeaturesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.PointChoice_runwayPoint = %s", angleindicationtimeslicetypeIdent, targetIdent))
			}
			if angleindicationtimeslicetype.PointChoice_aimingPoint != nil {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := angleindicationtimeslicetype.PointChoice_aimingPoint.GongGetIdentifier(stage.AixmfeaturesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.PointChoice_aimingPoint = %s", angleindicationtimeslicetypeIdent, targetIdent))
			}
			if angleindicationtimeslicetype.PointChoice_airportReferencePoint != nil {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := angleindicationtimeslicetype.PointChoice_airportReferencePoint.GongGetIdentifier(stage.AixmfeaturesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.PointChoice_airportReferencePoint = %s", angleindicationtimeslicetypeIdent, targetIdent))
			}
			for _, elem := range angleindicationtimeslicetype.Annotation {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := elem.GongGetIdentifier(stage.AixmfeaturesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Annotation = append(%s.Annotation, %s)", angleindicationtimeslicetypeIdent, angleindicationtimeslicetypeIdent, targetIdent))
			}
			for _, elem := range angleindicationtimeslicetype.Extension {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := elem.GongGetIdentifier(stage.AixmfeaturesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Extension = append(%s.Extension, %s)", angleindicationtimeslicetypeIdent, angleindicationtimeslicetypeIdent, targetIdent))
			}
		}
	}
	if stage.AixmfeaturesStage != nil {
		for _, angleindicationtype := range __gong__sortStageSetInstances(stage.AixmfeaturesStage.AngleIndicationTypes, stage.AixmfeaturesStage.AngleIndicationType_stagedOrder) {
			if lastStageDecl != "AixmfeaturesStage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "AixmfeaturesStage"
			}
			angleindicationtypeIdent := angleindicationtype.GongGetIdentifier(stage.AixmfeaturesStage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&aixmfeatures.AngleIndicationType{Name: %s}).Stage(stage.AixmfeaturesStage)", angleindicationtypeIdent, __gong__toRawStringLiteral(angleindicationtype.Name)))
			if lastStageVal != "AixmfeaturesStage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "AixmfeaturesStage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", angleindicationtypeIdent, __gong__toRawStringLiteral(angleindicationtype.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Id = %s", angleindicationtypeIdent, __gong__toRawStringLiteral(angleindicationtype.Id)))
			values.WriteString(fmt.Sprintf("\n\t%s.Description = %s", angleindicationtypeIdent, __gong__toRawStringLiteral(angleindicationtype.Description)))
			if angleindicationtype.Identifier != nil {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := angleindicationtype.Identifier.GongGetIdentifier(stage.GmlStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Identifier = %s", angleindicationtypeIdent, targetIdent))
			}
			for _, elem := range angleindicationtype.NameXSD {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := elem.GongGetIdentifier(stage.GmlStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.NameXSD = append(%s.NameXSD, %s)", angleindicationtypeIdent, angleindicationtypeIdent, targetIdent))
			}
			if angleindicationtype.BoundedBy != nil {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := angleindicationtype.BoundedBy.GongGetIdentifier(stage.GmlStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.BoundedBy = %s", angleindicationtypeIdent, targetIdent))
			}
			if angleindicationtype.FeatureMetadata != nil {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := angleindicationtype.FeatureMetadata.GongGetIdentifier(stage.AixmabstractgmlobjecttypesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.FeatureMetadata = %s", angleindicationtypeIdent, targetIdent))
			}
			for _, elem := range angleindicationtype.TimeSlice {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := elem.GongGetIdentifier(stage.AixmfeaturesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.TimeSlice = append(%s.TimeSlice, %s)", angleindicationtypeIdent, angleindicationtypeIdent, targetIdent))
			}
		}
	}
	if stage.AixmfeaturesStage != nil {
		for _, authorityfornavaidequipmentpropertytype := range __gong__sortStageSetInstances(stage.AixmfeaturesStage.AuthorityForNavaidEquipmentPropertyTypes, stage.AixmfeaturesStage.AuthorityForNavaidEquipmentPropertyType_stagedOrder) {
			if lastStageDecl != "AixmfeaturesStage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "AixmfeaturesStage"
			}
			authorityfornavaidequipmentpropertytypeIdent := authorityfornavaidequipmentpropertytype.GongGetIdentifier(stage.AixmfeaturesStage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&aixmfeatures.AuthorityForNavaidEquipmentPropertyType{Name: %s}).Stage(stage.AixmfeaturesStage)", authorityfornavaidequipmentpropertytypeIdent, __gong__toRawStringLiteral(authorityfornavaidequipmentpropertytype.Name)))
			if lastStageVal != "AixmfeaturesStage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "AixmfeaturesStage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", authorityfornavaidequipmentpropertytypeIdent, __gong__toRawStringLiteral(authorityfornavaidequipmentpropertytype.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.NilReason = %s", authorityfornavaidequipmentpropertytypeIdent, __gong__toRawStringLiteral(authorityfornavaidequipmentpropertytype.NilReason)))
			if authorityfornavaidequipmentpropertytype.AuthorityForNavaidEquipment != nil {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := authorityfornavaidequipmentpropertytype.AuthorityForNavaidEquipment.GongGetIdentifier(stage.AixmfeaturesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.AuthorityForNavaidEquipment = %s", authorityfornavaidequipmentpropertytypeIdent, targetIdent))
			}
		}
	}
	if stage.AixmfeaturesStage != nil {
		for _, authorityfornavaidequipmenttype := range __gong__sortStageSetInstances(stage.AixmfeaturesStage.AuthorityForNavaidEquipmentTypes, stage.AixmfeaturesStage.AuthorityForNavaidEquipmentType_stagedOrder) {
			if lastStageDecl != "AixmfeaturesStage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "AixmfeaturesStage"
			}
			authorityfornavaidequipmenttypeIdent := authorityfornavaidequipmenttype.GongGetIdentifier(stage.AixmfeaturesStage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&aixmfeatures.AuthorityForNavaidEquipmentType{Name: %s}).Stage(stage.AixmfeaturesStage)", authorityfornavaidequipmenttypeIdent, __gong__toRawStringLiteral(authorityfornavaidequipmenttype.Name)))
			if lastStageVal != "AixmfeaturesStage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "AixmfeaturesStage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", authorityfornavaidequipmenttypeIdent, __gong__toRawStringLiteral(authorityfornavaidequipmenttype.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Id = %s", authorityfornavaidequipmenttypeIdent, __gong__toRawStringLiteral(authorityfornavaidequipmenttype.Id)))
			values.WriteString(fmt.Sprintf("\n\t%s.Description = %s", authorityfornavaidequipmenttypeIdent, __gong__toRawStringLiteral(authorityfornavaidequipmenttype.Description)))
			if authorityfornavaidequipmenttype.Identifier != nil {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := authorityfornavaidequipmenttype.Identifier.GongGetIdentifier(stage.GmlStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Identifier = %s", authorityfornavaidequipmenttypeIdent, targetIdent))
			}
			for _, elem := range authorityfornavaidequipmenttype.NameXSD {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := elem.GongGetIdentifier(stage.GmlStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.NameXSD = append(%s.NameXSD, %s)", authorityfornavaidequipmenttypeIdent, authorityfornavaidequipmenttypeIdent, targetIdent))
			}
			if authorityfornavaidequipmenttype.Type_XSD != nil {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := authorityfornavaidequipmenttype.Type_XSD.GongGetIdentifier(stage.AixmdatatypesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Type_XSD = %s", authorityfornavaidequipmenttypeIdent, targetIdent))
			}
			for _, elem := range authorityfornavaidequipmenttype.Annotation {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := elem.GongGetIdentifier(stage.AixmfeaturesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Annotation = append(%s.Annotation, %s)", authorityfornavaidequipmenttypeIdent, authorityfornavaidequipmenttypeIdent, targetIdent))
			}
			if authorityfornavaidequipmenttype.TheOrganisationAuthority != nil {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := authorityfornavaidequipmenttype.TheOrganisationAuthority.GongGetIdentifier(stage.AixmfeaturesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.TheOrganisationAuthority = %s", authorityfornavaidequipmenttypeIdent, targetIdent))
			}
			for _, elem := range authorityfornavaidequipmenttype.Extension {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := elem.GongGetIdentifier(stage.AixmfeaturesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Extension = append(%s.Extension, %s)", authorityfornavaidequipmenttypeIdent, authorityfornavaidequipmenttypeIdent, targetIdent))
			}
		}
	}
	if stage.AixmfeaturesStage != nil {
		for _, authorityforspecialnavigationstationpropertytype := range __gong__sortStageSetInstances(stage.AixmfeaturesStage.AuthorityForSpecialNavigationStationPropertyTypes, stage.AixmfeaturesStage.AuthorityForSpecialNavigationStationPropertyType_stagedOrder) {
			if lastStageDecl != "AixmfeaturesStage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "AixmfeaturesStage"
			}
			authorityforspecialnavigationstationpropertytypeIdent := authorityforspecialnavigationstationpropertytype.GongGetIdentifier(stage.AixmfeaturesStage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&aixmfeatures.AuthorityForSpecialNavigationStationPropertyType{Name: %s}).Stage(stage.AixmfeaturesStage)", authorityforspecialnavigationstationpropertytypeIdent, __gong__toRawStringLiteral(authorityforspecialnavigationstationpropertytype.Name)))
			if lastStageVal != "AixmfeaturesStage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "AixmfeaturesStage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", authorityforspecialnavigationstationpropertytypeIdent, __gong__toRawStringLiteral(authorityforspecialnavigationstationpropertytype.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.NilReason = %s", authorityforspecialnavigationstationpropertytypeIdent, __gong__toRawStringLiteral(authorityforspecialnavigationstationpropertytype.NilReason)))
			if authorityforspecialnavigationstationpropertytype.AuthorityForSpecialNavigationStation != nil {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := authorityforspecialnavigationstationpropertytype.AuthorityForSpecialNavigationStation.GongGetIdentifier(stage.AixmfeaturesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.AuthorityForSpecialNavigationStation = %s", authorityforspecialnavigationstationpropertytypeIdent, targetIdent))
			}
		}
	}
	if stage.AixmfeaturesStage != nil {
		for _, authorityforspecialnavigationstationtype := range __gong__sortStageSetInstances(stage.AixmfeaturesStage.AuthorityForSpecialNavigationStationTypes, stage.AixmfeaturesStage.AuthorityForSpecialNavigationStationType_stagedOrder) {
			if lastStageDecl != "AixmfeaturesStage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "AixmfeaturesStage"
			}
			authorityforspecialnavigationstationtypeIdent := authorityforspecialnavigationstationtype.GongGetIdentifier(stage.AixmfeaturesStage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&aixmfeatures.AuthorityForSpecialNavigationStationType{Name: %s}).Stage(stage.AixmfeaturesStage)", authorityforspecialnavigationstationtypeIdent, __gong__toRawStringLiteral(authorityforspecialnavigationstationtype.Name)))
			if lastStageVal != "AixmfeaturesStage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "AixmfeaturesStage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", authorityforspecialnavigationstationtypeIdent, __gong__toRawStringLiteral(authorityforspecialnavigationstationtype.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Id = %s", authorityforspecialnavigationstationtypeIdent, __gong__toRawStringLiteral(authorityforspecialnavigationstationtype.Id)))
			values.WriteString(fmt.Sprintf("\n\t%s.Description = %s", authorityforspecialnavigationstationtypeIdent, __gong__toRawStringLiteral(authorityforspecialnavigationstationtype.Description)))
			if authorityforspecialnavigationstationtype.Identifier != nil {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := authorityforspecialnavigationstationtype.Identifier.GongGetIdentifier(stage.GmlStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Identifier = %s", authorityforspecialnavigationstationtypeIdent, targetIdent))
			}
			for _, elem := range authorityforspecialnavigationstationtype.NameXSD {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := elem.GongGetIdentifier(stage.GmlStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.NameXSD = append(%s.NameXSD, %s)", authorityforspecialnavigationstationtypeIdent, authorityforspecialnavigationstationtypeIdent, targetIdent))
			}
			if authorityforspecialnavigationstationtype.Type_XSD != nil {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := authorityforspecialnavigationstationtype.Type_XSD.GongGetIdentifier(stage.AixmdatatypesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Type_XSD = %s", authorityforspecialnavigationstationtypeIdent, targetIdent))
			}
			for _, elem := range authorityforspecialnavigationstationtype.Annotation {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := elem.GongGetIdentifier(stage.AixmfeaturesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Annotation = append(%s.Annotation, %s)", authorityforspecialnavigationstationtypeIdent, authorityforspecialnavigationstationtypeIdent, targetIdent))
			}
			if authorityforspecialnavigationstationtype.TheOrganisationAuthority != nil {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := authorityforspecialnavigationstationtype.TheOrganisationAuthority.GongGetIdentifier(stage.AixmfeaturesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.TheOrganisationAuthority = %s", authorityforspecialnavigationstationtypeIdent, targetIdent))
			}
			for _, elem := range authorityforspecialnavigationstationtype.Extension {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := elem.GongGetIdentifier(stage.AixmfeaturesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Extension = append(%s.Extension, %s)", authorityforspecialnavigationstationtypeIdent, authorityforspecialnavigationstationtypeIdent, targetIdent))
			}
		}
	}
	if stage.AixmfeaturesStage != nil {
		for _, authorityforspecialnavigationsystempropertytype := range __gong__sortStageSetInstances(stage.AixmfeaturesStage.AuthorityForSpecialNavigationSystemPropertyTypes, stage.AixmfeaturesStage.AuthorityForSpecialNavigationSystemPropertyType_stagedOrder) {
			if lastStageDecl != "AixmfeaturesStage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "AixmfeaturesStage"
			}
			authorityforspecialnavigationsystempropertytypeIdent := authorityforspecialnavigationsystempropertytype.GongGetIdentifier(stage.AixmfeaturesStage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&aixmfeatures.AuthorityForSpecialNavigationSystemPropertyType{Name: %s}).Stage(stage.AixmfeaturesStage)", authorityforspecialnavigationsystempropertytypeIdent, __gong__toRawStringLiteral(authorityforspecialnavigationsystempropertytype.Name)))
			if lastStageVal != "AixmfeaturesStage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "AixmfeaturesStage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", authorityforspecialnavigationsystempropertytypeIdent, __gong__toRawStringLiteral(authorityforspecialnavigationsystempropertytype.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.NilReason = %s", authorityforspecialnavigationsystempropertytypeIdent, __gong__toRawStringLiteral(authorityforspecialnavigationsystempropertytype.NilReason)))
			if authorityforspecialnavigationsystempropertytype.AuthorityForSpecialNavigationSystem != nil {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := authorityforspecialnavigationsystempropertytype.AuthorityForSpecialNavigationSystem.GongGetIdentifier(stage.AixmfeaturesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.AuthorityForSpecialNavigationSystem = %s", authorityforspecialnavigationsystempropertytypeIdent, targetIdent))
			}
		}
	}
	if stage.AixmfeaturesStage != nil {
		for _, authorityforspecialnavigationsystemtype := range __gong__sortStageSetInstances(stage.AixmfeaturesStage.AuthorityForSpecialNavigationSystemTypes, stage.AixmfeaturesStage.AuthorityForSpecialNavigationSystemType_stagedOrder) {
			if lastStageDecl != "AixmfeaturesStage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "AixmfeaturesStage"
			}
			authorityforspecialnavigationsystemtypeIdent := authorityforspecialnavigationsystemtype.GongGetIdentifier(stage.AixmfeaturesStage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&aixmfeatures.AuthorityForSpecialNavigationSystemType{Name: %s}).Stage(stage.AixmfeaturesStage)", authorityforspecialnavigationsystemtypeIdent, __gong__toRawStringLiteral(authorityforspecialnavigationsystemtype.Name)))
			if lastStageVal != "AixmfeaturesStage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "AixmfeaturesStage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", authorityforspecialnavigationsystemtypeIdent, __gong__toRawStringLiteral(authorityforspecialnavigationsystemtype.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Id = %s", authorityforspecialnavigationsystemtypeIdent, __gong__toRawStringLiteral(authorityforspecialnavigationsystemtype.Id)))
			values.WriteString(fmt.Sprintf("\n\t%s.Description = %s", authorityforspecialnavigationsystemtypeIdent, __gong__toRawStringLiteral(authorityforspecialnavigationsystemtype.Description)))
			if authorityforspecialnavigationsystemtype.Identifier != nil {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := authorityforspecialnavigationsystemtype.Identifier.GongGetIdentifier(stage.GmlStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Identifier = %s", authorityforspecialnavigationsystemtypeIdent, targetIdent))
			}
			for _, elem := range authorityforspecialnavigationsystemtype.NameXSD {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := elem.GongGetIdentifier(stage.GmlStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.NameXSD = append(%s.NameXSD, %s)", authorityforspecialnavigationsystemtypeIdent, authorityforspecialnavigationsystemtypeIdent, targetIdent))
			}
			if authorityforspecialnavigationsystemtype.Type_XSD != nil {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := authorityforspecialnavigationsystemtype.Type_XSD.GongGetIdentifier(stage.AixmdatatypesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Type_XSD = %s", authorityforspecialnavigationsystemtypeIdent, targetIdent))
			}
			for _, elem := range authorityforspecialnavigationsystemtype.Annotation {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := elem.GongGetIdentifier(stage.AixmfeaturesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Annotation = append(%s.Annotation, %s)", authorityforspecialnavigationsystemtypeIdent, authorityforspecialnavigationsystemtypeIdent, targetIdent))
			}
			if authorityforspecialnavigationsystemtype.TheOrganisationAuthority != nil {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := authorityforspecialnavigationsystemtype.TheOrganisationAuthority.GongGetIdentifier(stage.AixmfeaturesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.TheOrganisationAuthority = %s", authorityforspecialnavigationsystemtypeIdent, targetIdent))
			}
			for _, elem := range authorityforspecialnavigationsystemtype.Extension {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := elem.GongGetIdentifier(stage.AixmfeaturesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Extension = append(%s.Extension, %s)", authorityforspecialnavigationsystemtypeIdent, authorityforspecialnavigationsystemtypeIdent, targetIdent))
			}
		}
	}
	if stage.AixmfeaturesStage != nil {
		for _, azimuthtimeslicepropertytype := range __gong__sortStageSetInstances(stage.AixmfeaturesStage.AzimuthTimeSlicePropertyTypes, stage.AixmfeaturesStage.AzimuthTimeSlicePropertyType_stagedOrder) {
			if lastStageDecl != "AixmfeaturesStage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "AixmfeaturesStage"
			}
			azimuthtimeslicepropertytypeIdent := azimuthtimeslicepropertytype.GongGetIdentifier(stage.AixmfeaturesStage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&aixmfeatures.AzimuthTimeSlicePropertyType{Name: %s}).Stage(stage.AixmfeaturesStage)", azimuthtimeslicepropertytypeIdent, __gong__toRawStringLiteral(azimuthtimeslicepropertytype.Name)))
			if lastStageVal != "AixmfeaturesStage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "AixmfeaturesStage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", azimuthtimeslicepropertytypeIdent, __gong__toRawStringLiteral(azimuthtimeslicepropertytype.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Owns = %t", azimuthtimeslicepropertytypeIdent, azimuthtimeslicepropertytype.Owns))
			if azimuthtimeslicepropertytype.AzimuthTimeSlice != nil {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := azimuthtimeslicepropertytype.AzimuthTimeSlice.GongGetIdentifier(stage.AixmfeaturesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.AzimuthTimeSlice = %s", azimuthtimeslicepropertytypeIdent, targetIdent))
			}
		}
	}
	if stage.AixmfeaturesStage != nil {
		for _, azimuthtimeslicetype := range __gong__sortStageSetInstances(stage.AixmfeaturesStage.AzimuthTimeSliceTypes, stage.AixmfeaturesStage.AzimuthTimeSliceType_stagedOrder) {
			if lastStageDecl != "AixmfeaturesStage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "AixmfeaturesStage"
			}
			azimuthtimeslicetypeIdent := azimuthtimeslicetype.GongGetIdentifier(stage.AixmfeaturesStage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&aixmfeatures.AzimuthTimeSliceType{Name: %s}).Stage(stage.AixmfeaturesStage)", azimuthtimeslicetypeIdent, __gong__toRawStringLiteral(azimuthtimeslicetype.Name)))
			if lastStageVal != "AixmfeaturesStage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "AixmfeaturesStage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", azimuthtimeslicetypeIdent, __gong__toRawStringLiteral(azimuthtimeslicetype.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Id = %s", azimuthtimeslicetypeIdent, __gong__toRawStringLiteral(azimuthtimeslicetype.Id)))
			values.WriteString(fmt.Sprintf("\n\t%s.Description = %s", azimuthtimeslicetypeIdent, __gong__toRawStringLiteral(azimuthtimeslicetype.Description)))
			values.WriteString(fmt.Sprintf("\n\t%s.Interpretation = %s", azimuthtimeslicetypeIdent, __gong__toRawStringLiteral(azimuthtimeslicetype.Interpretation)))
			values.WriteString(fmt.Sprintf("\n\t%s.SequenceNumber = %d", azimuthtimeslicetypeIdent, azimuthtimeslicetype.SequenceNumber))
			values.WriteString(fmt.Sprintf("\n\t%s.CorrectionNumber = %d", azimuthtimeslicetypeIdent, azimuthtimeslicetype.CorrectionNumber))
			if azimuthtimeslicetype.Identifier != nil {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := azimuthtimeslicetype.Identifier.GongGetIdentifier(stage.GmlStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Identifier = %s", azimuthtimeslicetypeIdent, targetIdent))
			}
			for _, elem := range azimuthtimeslicetype.NameXSD {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := elem.GongGetIdentifier(stage.GmlStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.NameXSD = append(%s.NameXSD, %s)", azimuthtimeslicetypeIdent, azimuthtimeslicetypeIdent, targetIdent))
			}
			if azimuthtimeslicetype.ValidTime != nil {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := azimuthtimeslicetype.ValidTime.GongGetIdentifier(stage.GmlStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.ValidTime = %s", azimuthtimeslicetypeIdent, targetIdent))
			}
			if azimuthtimeslicetype.TimeSliceMetadata != nil {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := azimuthtimeslicetype.TimeSliceMetadata.GongGetIdentifier(stage.AixmabstractgmlobjecttypesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.TimeSliceMetadata = %s", azimuthtimeslicetypeIdent, targetIdent))
			}
			if azimuthtimeslicetype.FeatureLifetime != nil {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := azimuthtimeslicetype.FeatureLifetime.GongGetIdentifier(stage.GmlStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.FeatureLifetime = %s", azimuthtimeslicetypeIdent, targetIdent))
			}
			if azimuthtimeslicetype.Designator != nil {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := azimuthtimeslicetype.Designator.GongGetIdentifier(stage.AixmdatatypesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Designator = %s", azimuthtimeslicetypeIdent, targetIdent))
			}
			if azimuthtimeslicetype.EmissionClass != nil {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := azimuthtimeslicetype.EmissionClass.GongGetIdentifier(stage.AixmdatatypesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.EmissionClass = %s", azimuthtimeslicetypeIdent, targetIdent))
			}
			if azimuthtimeslicetype.Mobile != nil {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := azimuthtimeslicetype.Mobile.GongGetIdentifier(stage.AixmdatatypesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Mobile = %s", azimuthtimeslicetypeIdent, targetIdent))
			}
			if azimuthtimeslicetype.MagneticVariation != nil {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := azimuthtimeslicetype.MagneticVariation.GongGetIdentifier(stage.AixmdatatypesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.MagneticVariation = %s", azimuthtimeslicetypeIdent, targetIdent))
			}
			if azimuthtimeslicetype.MagneticVariationAccuracy != nil {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := azimuthtimeslicetype.MagneticVariationAccuracy.GongGetIdentifier(stage.AixmdatatypesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.MagneticVariationAccuracy = %s", azimuthtimeslicetypeIdent, targetIdent))
			}
			if azimuthtimeslicetype.DateMagneticVariation != nil {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := azimuthtimeslicetype.DateMagneticVariation.GongGetIdentifier(stage.AixmdatatypesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.DateMagneticVariation = %s", azimuthtimeslicetypeIdent, targetIdent))
			}
			if azimuthtimeslicetype.FlightChecked != nil {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := azimuthtimeslicetype.FlightChecked.GongGetIdentifier(stage.AixmdatatypesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.FlightChecked = %s", azimuthtimeslicetypeIdent, targetIdent))
			}
			if azimuthtimeslicetype.Location != nil {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := azimuthtimeslicetype.Location.GongGetIdentifier(stage.AixmfeaturesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Location = %s", azimuthtimeslicetypeIdent, targetIdent))
			}
			for _, elem := range azimuthtimeslicetype.Authority {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := elem.GongGetIdentifier(stage.AixmfeaturesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Authority = append(%s.Authority, %s)", azimuthtimeslicetypeIdent, azimuthtimeslicetypeIdent, targetIdent))
			}
			for _, elem := range azimuthtimeslicetype.Monitoring {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := elem.GongGetIdentifier(stage.AixmfeaturesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Monitoring = append(%s.Monitoring, %s)", azimuthtimeslicetypeIdent, azimuthtimeslicetypeIdent, targetIdent))
			}
			for _, elem := range azimuthtimeslicetype.Availability {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := elem.GongGetIdentifier(stage.AixmfeaturesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Availability = append(%s.Availability, %s)", azimuthtimeslicetypeIdent, azimuthtimeslicetypeIdent, targetIdent))
			}
			for _, elem := range azimuthtimeslicetype.Annotation {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := elem.GongGetIdentifier(stage.AixmfeaturesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Annotation = append(%s.Annotation, %s)", azimuthtimeslicetypeIdent, azimuthtimeslicetypeIdent, targetIdent))
			}
			if azimuthtimeslicetype.Type_XSD != nil {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := azimuthtimeslicetype.Type_XSD.GongGetIdentifier(stage.AixmdatatypesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Type_XSD = %s", azimuthtimeslicetypeIdent, targetIdent))
			}
			if azimuthtimeslicetype.TrueBearing != nil {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := azimuthtimeslicetype.TrueBearing.GongGetIdentifier(stage.AixmdatatypesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.TrueBearing = %s", azimuthtimeslicetypeIdent, targetIdent))
			}
			if azimuthtimeslicetype.TrueBearingAccuracy != nil {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := azimuthtimeslicetype.TrueBearingAccuracy.GongGetIdentifier(stage.AixmdatatypesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.TrueBearingAccuracy = %s", azimuthtimeslicetypeIdent, targetIdent))
			}
			if azimuthtimeslicetype.MagneticBearing != nil {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := azimuthtimeslicetype.MagneticBearing.GongGetIdentifier(stage.AixmdatatypesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.MagneticBearing = %s", azimuthtimeslicetypeIdent, targetIdent))
			}
			if azimuthtimeslicetype.AngleProportionalLeft != nil {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := azimuthtimeslicetype.AngleProportionalLeft.GongGetIdentifier(stage.AixmdatatypesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.AngleProportionalLeft = %s", azimuthtimeslicetypeIdent, targetIdent))
			}
			if azimuthtimeslicetype.AngleProportionalRight != nil {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := azimuthtimeslicetype.AngleProportionalRight.GongGetIdentifier(stage.AixmdatatypesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.AngleProportionalRight = %s", azimuthtimeslicetypeIdent, targetIdent))
			}
			if azimuthtimeslicetype.AngleCoverLeft != nil {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := azimuthtimeslicetype.AngleCoverLeft.GongGetIdentifier(stage.AixmdatatypesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.AngleCoverLeft = %s", azimuthtimeslicetypeIdent, targetIdent))
			}
			if azimuthtimeslicetype.AngleCoverRight != nil {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := azimuthtimeslicetype.AngleCoverRight.GongGetIdentifier(stage.AixmdatatypesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.AngleCoverRight = %s", azimuthtimeslicetypeIdent, targetIdent))
			}
			if azimuthtimeslicetype.Channel != nil {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := azimuthtimeslicetype.Channel.GongGetIdentifier(stage.AixmdatatypesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Channel = %s", azimuthtimeslicetypeIdent, targetIdent))
			}
			for _, elem := range azimuthtimeslicetype.Extension {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := elem.GongGetIdentifier(stage.AixmfeaturesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Extension = append(%s.Extension, %s)", azimuthtimeslicetypeIdent, azimuthtimeslicetypeIdent, targetIdent))
			}
		}
	}
	if stage.AixmfeaturesStage != nil {
		for _, azimuthtype := range __gong__sortStageSetInstances(stage.AixmfeaturesStage.AzimuthTypes, stage.AixmfeaturesStage.AzimuthType_stagedOrder) {
			if lastStageDecl != "AixmfeaturesStage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "AixmfeaturesStage"
			}
			azimuthtypeIdent := azimuthtype.GongGetIdentifier(stage.AixmfeaturesStage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&aixmfeatures.AzimuthType{Name: %s}).Stage(stage.AixmfeaturesStage)", azimuthtypeIdent, __gong__toRawStringLiteral(azimuthtype.Name)))
			if lastStageVal != "AixmfeaturesStage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "AixmfeaturesStage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", azimuthtypeIdent, __gong__toRawStringLiteral(azimuthtype.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Id = %s", azimuthtypeIdent, __gong__toRawStringLiteral(azimuthtype.Id)))
			values.WriteString(fmt.Sprintf("\n\t%s.Description = %s", azimuthtypeIdent, __gong__toRawStringLiteral(azimuthtype.Description)))
			if azimuthtype.Identifier != nil {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := azimuthtype.Identifier.GongGetIdentifier(stage.GmlStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Identifier = %s", azimuthtypeIdent, targetIdent))
			}
			for _, elem := range azimuthtype.NameXSD {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := elem.GongGetIdentifier(stage.GmlStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.NameXSD = append(%s.NameXSD, %s)", azimuthtypeIdent, azimuthtypeIdent, targetIdent))
			}
			if azimuthtype.BoundedBy != nil {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := azimuthtype.BoundedBy.GongGetIdentifier(stage.GmlStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.BoundedBy = %s", azimuthtypeIdent, targetIdent))
			}
			if azimuthtype.FeatureMetadata != nil {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := azimuthtype.FeatureMetadata.GongGetIdentifier(stage.AixmabstractgmlobjecttypesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.FeatureMetadata = %s", azimuthtypeIdent, targetIdent))
			}
			for _, elem := range azimuthtype.TimeSlice {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := elem.GongGetIdentifier(stage.AixmfeaturesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.TimeSlice = append(%s.TimeSlice, %s)", azimuthtypeIdent, azimuthtypeIdent, targetIdent))
			}
		}
	}
	if stage.AixmfeaturesStage != nil {
		for _, checkpointinstimeslicepropertytype := range __gong__sortStageSetInstances(stage.AixmfeaturesStage.CheckpointINSTimeSlicePropertyTypes, stage.AixmfeaturesStage.CheckpointINSTimeSlicePropertyType_stagedOrder) {
			if lastStageDecl != "AixmfeaturesStage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "AixmfeaturesStage"
			}
			checkpointinstimeslicepropertytypeIdent := checkpointinstimeslicepropertytype.GongGetIdentifier(stage.AixmfeaturesStage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&aixmfeatures.CheckpointINSTimeSlicePropertyType{Name: %s}).Stage(stage.AixmfeaturesStage)", checkpointinstimeslicepropertytypeIdent, __gong__toRawStringLiteral(checkpointinstimeslicepropertytype.Name)))
			if lastStageVal != "AixmfeaturesStage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "AixmfeaturesStage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", checkpointinstimeslicepropertytypeIdent, __gong__toRawStringLiteral(checkpointinstimeslicepropertytype.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Owns = %t", checkpointinstimeslicepropertytypeIdent, checkpointinstimeslicepropertytype.Owns))
			if checkpointinstimeslicepropertytype.CheckpointINSTimeSlice != nil {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := checkpointinstimeslicepropertytype.CheckpointINSTimeSlice.GongGetIdentifier(stage.AixmfeaturesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.CheckpointINSTimeSlice = %s", checkpointinstimeslicepropertytypeIdent, targetIdent))
			}
		}
	}
	if stage.AixmfeaturesStage != nil {
		for _, checkpointinstimeslicetype := range __gong__sortStageSetInstances(stage.AixmfeaturesStage.CheckpointINSTimeSliceTypes, stage.AixmfeaturesStage.CheckpointINSTimeSliceType_stagedOrder) {
			if lastStageDecl != "AixmfeaturesStage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "AixmfeaturesStage"
			}
			checkpointinstimeslicetypeIdent := checkpointinstimeslicetype.GongGetIdentifier(stage.AixmfeaturesStage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&aixmfeatures.CheckpointINSTimeSliceType{Name: %s}).Stage(stage.AixmfeaturesStage)", checkpointinstimeslicetypeIdent, __gong__toRawStringLiteral(checkpointinstimeslicetype.Name)))
			if lastStageVal != "AixmfeaturesStage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "AixmfeaturesStage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", checkpointinstimeslicetypeIdent, __gong__toRawStringLiteral(checkpointinstimeslicetype.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Id = %s", checkpointinstimeslicetypeIdent, __gong__toRawStringLiteral(checkpointinstimeslicetype.Id)))
			values.WriteString(fmt.Sprintf("\n\t%s.Description = %s", checkpointinstimeslicetypeIdent, __gong__toRawStringLiteral(checkpointinstimeslicetype.Description)))
			values.WriteString(fmt.Sprintf("\n\t%s.Interpretation = %s", checkpointinstimeslicetypeIdent, __gong__toRawStringLiteral(checkpointinstimeslicetype.Interpretation)))
			values.WriteString(fmt.Sprintf("\n\t%s.SequenceNumber = %d", checkpointinstimeslicetypeIdent, checkpointinstimeslicetype.SequenceNumber))
			values.WriteString(fmt.Sprintf("\n\t%s.CorrectionNumber = %d", checkpointinstimeslicetypeIdent, checkpointinstimeslicetype.CorrectionNumber))
			if checkpointinstimeslicetype.Identifier != nil {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := checkpointinstimeslicetype.Identifier.GongGetIdentifier(stage.GmlStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Identifier = %s", checkpointinstimeslicetypeIdent, targetIdent))
			}
			for _, elem := range checkpointinstimeslicetype.NameXSD {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := elem.GongGetIdentifier(stage.GmlStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.NameXSD = append(%s.NameXSD, %s)", checkpointinstimeslicetypeIdent, checkpointinstimeslicetypeIdent, targetIdent))
			}
			if checkpointinstimeslicetype.ValidTime != nil {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := checkpointinstimeslicetype.ValidTime.GongGetIdentifier(stage.GmlStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.ValidTime = %s", checkpointinstimeslicetypeIdent, targetIdent))
			}
			if checkpointinstimeslicetype.TimeSliceMetadata != nil {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := checkpointinstimeslicetype.TimeSliceMetadata.GongGetIdentifier(stage.AixmabstractgmlobjecttypesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.TimeSliceMetadata = %s", checkpointinstimeslicetypeIdent, targetIdent))
			}
			if checkpointinstimeslicetype.FeatureLifetime != nil {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := checkpointinstimeslicetype.FeatureLifetime.GongGetIdentifier(stage.GmlStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.FeatureLifetime = %s", checkpointinstimeslicetypeIdent, targetIdent))
			}
			if checkpointinstimeslicetype.Category != nil {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := checkpointinstimeslicetype.Category.GongGetIdentifier(stage.AixmdatatypesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Category = %s", checkpointinstimeslicetypeIdent, targetIdent))
			}
			if checkpointinstimeslicetype.UpperLimit != nil {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := checkpointinstimeslicetype.UpperLimit.GongGetIdentifier(stage.AixmdatatypesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.UpperLimit = %s", checkpointinstimeslicetypeIdent, targetIdent))
			}
			if checkpointinstimeslicetype.UpperLimitReference != nil {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := checkpointinstimeslicetype.UpperLimitReference.GongGetIdentifier(stage.AixmdatatypesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.UpperLimitReference = %s", checkpointinstimeslicetypeIdent, targetIdent))
			}
			if checkpointinstimeslicetype.LowerLimit != nil {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := checkpointinstimeslicetype.LowerLimit.GongGetIdentifier(stage.AixmdatatypesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.LowerLimit = %s", checkpointinstimeslicetypeIdent, targetIdent))
			}
			if checkpointinstimeslicetype.LowerLimitReference != nil {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := checkpointinstimeslicetype.LowerLimitReference.GongGetIdentifier(stage.AixmdatatypesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.LowerLimitReference = %s", checkpointinstimeslicetypeIdent, targetIdent))
			}
			if checkpointinstimeslicetype.AltitudeInterpretation != nil {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := checkpointinstimeslicetype.AltitudeInterpretation.GongGetIdentifier(stage.AixmdatatypesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.AltitudeInterpretation = %s", checkpointinstimeslicetypeIdent, targetIdent))
			}
			if checkpointinstimeslicetype.Distance != nil {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := checkpointinstimeslicetype.Distance.GongGetIdentifier(stage.AixmdatatypesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Distance = %s", checkpointinstimeslicetypeIdent, targetIdent))
			}
			if checkpointinstimeslicetype.Angle != nil {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := checkpointinstimeslicetype.Angle.GongGetIdentifier(stage.AixmdatatypesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Angle = %s", checkpointinstimeslicetypeIdent, targetIdent))
			}
			if checkpointinstimeslicetype.Position != nil {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := checkpointinstimeslicetype.Position.GongGetIdentifier(stage.AixmfeaturesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Position = %s", checkpointinstimeslicetypeIdent, targetIdent))
			}
			if checkpointinstimeslicetype.AirportHeliport != nil {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := checkpointinstimeslicetype.AirportHeliport.GongGetIdentifier(stage.AixmfeaturesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.AirportHeliport = %s", checkpointinstimeslicetypeIdent, targetIdent))
			}
			for _, elem := range checkpointinstimeslicetype.Annotation {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := elem.GongGetIdentifier(stage.AixmfeaturesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Annotation = append(%s.Annotation, %s)", checkpointinstimeslicetypeIdent, checkpointinstimeslicetypeIdent, targetIdent))
			}
			for _, elem := range checkpointinstimeslicetype.Extension {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := elem.GongGetIdentifier(stage.AixmfeaturesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Extension = append(%s.Extension, %s)", checkpointinstimeslicetypeIdent, checkpointinstimeslicetypeIdent, targetIdent))
			}
		}
	}
	if stage.AixmfeaturesStage != nil {
		for _, checkpointinstype := range __gong__sortStageSetInstances(stage.AixmfeaturesStage.CheckpointINSTypes, stage.AixmfeaturesStage.CheckpointINSType_stagedOrder) {
			if lastStageDecl != "AixmfeaturesStage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "AixmfeaturesStage"
			}
			checkpointinstypeIdent := checkpointinstype.GongGetIdentifier(stage.AixmfeaturesStage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&aixmfeatures.CheckpointINSType{Name: %s}).Stage(stage.AixmfeaturesStage)", checkpointinstypeIdent, __gong__toRawStringLiteral(checkpointinstype.Name)))
			if lastStageVal != "AixmfeaturesStage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "AixmfeaturesStage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", checkpointinstypeIdent, __gong__toRawStringLiteral(checkpointinstype.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Id = %s", checkpointinstypeIdent, __gong__toRawStringLiteral(checkpointinstype.Id)))
			values.WriteString(fmt.Sprintf("\n\t%s.Description = %s", checkpointinstypeIdent, __gong__toRawStringLiteral(checkpointinstype.Description)))
			if checkpointinstype.Identifier != nil {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := checkpointinstype.Identifier.GongGetIdentifier(stage.GmlStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Identifier = %s", checkpointinstypeIdent, targetIdent))
			}
			for _, elem := range checkpointinstype.NameXSD {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := elem.GongGetIdentifier(stage.GmlStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.NameXSD = append(%s.NameXSD, %s)", checkpointinstypeIdent, checkpointinstypeIdent, targetIdent))
			}
			if checkpointinstype.BoundedBy != nil {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := checkpointinstype.BoundedBy.GongGetIdentifier(stage.GmlStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.BoundedBy = %s", checkpointinstypeIdent, targetIdent))
			}
			if checkpointinstype.FeatureMetadata != nil {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := checkpointinstype.FeatureMetadata.GongGetIdentifier(stage.AixmabstractgmlobjecttypesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.FeatureMetadata = %s", checkpointinstypeIdent, targetIdent))
			}
			for _, elem := range checkpointinstype.TimeSlice {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := elem.GongGetIdentifier(stage.AixmfeaturesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.TimeSlice = append(%s.TimeSlice, %s)", checkpointinstypeIdent, checkpointinstypeIdent, targetIdent))
			}
		}
	}
	if stage.AixmfeaturesStage != nil {
		for _, checkpointvortimeslicepropertytype := range __gong__sortStageSetInstances(stage.AixmfeaturesStage.CheckpointVORTimeSlicePropertyTypes, stage.AixmfeaturesStage.CheckpointVORTimeSlicePropertyType_stagedOrder) {
			if lastStageDecl != "AixmfeaturesStage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "AixmfeaturesStage"
			}
			checkpointvortimeslicepropertytypeIdent := checkpointvortimeslicepropertytype.GongGetIdentifier(stage.AixmfeaturesStage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&aixmfeatures.CheckpointVORTimeSlicePropertyType{Name: %s}).Stage(stage.AixmfeaturesStage)", checkpointvortimeslicepropertytypeIdent, __gong__toRawStringLiteral(checkpointvortimeslicepropertytype.Name)))
			if lastStageVal != "AixmfeaturesStage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "AixmfeaturesStage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", checkpointvortimeslicepropertytypeIdent, __gong__toRawStringLiteral(checkpointvortimeslicepropertytype.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Owns = %t", checkpointvortimeslicepropertytypeIdent, checkpointvortimeslicepropertytype.Owns))
			if checkpointvortimeslicepropertytype.CheckpointVORTimeSlice != nil {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := checkpointvortimeslicepropertytype.CheckpointVORTimeSlice.GongGetIdentifier(stage.AixmfeaturesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.CheckpointVORTimeSlice = %s", checkpointvortimeslicepropertytypeIdent, targetIdent))
			}
		}
	}
	if stage.AixmfeaturesStage != nil {
		for _, checkpointvortimeslicetype := range __gong__sortStageSetInstances(stage.AixmfeaturesStage.CheckpointVORTimeSliceTypes, stage.AixmfeaturesStage.CheckpointVORTimeSliceType_stagedOrder) {
			if lastStageDecl != "AixmfeaturesStage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "AixmfeaturesStage"
			}
			checkpointvortimeslicetypeIdent := checkpointvortimeslicetype.GongGetIdentifier(stage.AixmfeaturesStage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&aixmfeatures.CheckpointVORTimeSliceType{Name: %s}).Stage(stage.AixmfeaturesStage)", checkpointvortimeslicetypeIdent, __gong__toRawStringLiteral(checkpointvortimeslicetype.Name)))
			if lastStageVal != "AixmfeaturesStage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "AixmfeaturesStage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", checkpointvortimeslicetypeIdent, __gong__toRawStringLiteral(checkpointvortimeslicetype.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Id = %s", checkpointvortimeslicetypeIdent, __gong__toRawStringLiteral(checkpointvortimeslicetype.Id)))
			values.WriteString(fmt.Sprintf("\n\t%s.Description = %s", checkpointvortimeslicetypeIdent, __gong__toRawStringLiteral(checkpointvortimeslicetype.Description)))
			values.WriteString(fmt.Sprintf("\n\t%s.Interpretation = %s", checkpointvortimeslicetypeIdent, __gong__toRawStringLiteral(checkpointvortimeslicetype.Interpretation)))
			values.WriteString(fmt.Sprintf("\n\t%s.SequenceNumber = %d", checkpointvortimeslicetypeIdent, checkpointvortimeslicetype.SequenceNumber))
			values.WriteString(fmt.Sprintf("\n\t%s.CorrectionNumber = %d", checkpointvortimeslicetypeIdent, checkpointvortimeslicetype.CorrectionNumber))
			if checkpointvortimeslicetype.Identifier != nil {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := checkpointvortimeslicetype.Identifier.GongGetIdentifier(stage.GmlStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Identifier = %s", checkpointvortimeslicetypeIdent, targetIdent))
			}
			for _, elem := range checkpointvortimeslicetype.NameXSD {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := elem.GongGetIdentifier(stage.GmlStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.NameXSD = append(%s.NameXSD, %s)", checkpointvortimeslicetypeIdent, checkpointvortimeslicetypeIdent, targetIdent))
			}
			if checkpointvortimeslicetype.ValidTime != nil {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := checkpointvortimeslicetype.ValidTime.GongGetIdentifier(stage.GmlStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.ValidTime = %s", checkpointvortimeslicetypeIdent, targetIdent))
			}
			if checkpointvortimeslicetype.TimeSliceMetadata != nil {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := checkpointvortimeslicetype.TimeSliceMetadata.GongGetIdentifier(stage.AixmabstractgmlobjecttypesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.TimeSliceMetadata = %s", checkpointvortimeslicetypeIdent, targetIdent))
			}
			if checkpointvortimeslicetype.FeatureLifetime != nil {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := checkpointvortimeslicetype.FeatureLifetime.GongGetIdentifier(stage.GmlStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.FeatureLifetime = %s", checkpointvortimeslicetypeIdent, targetIdent))
			}
			if checkpointvortimeslicetype.Category != nil {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := checkpointvortimeslicetype.Category.GongGetIdentifier(stage.AixmdatatypesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Category = %s", checkpointvortimeslicetypeIdent, targetIdent))
			}
			if checkpointvortimeslicetype.UpperLimit != nil {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := checkpointvortimeslicetype.UpperLimit.GongGetIdentifier(stage.AixmdatatypesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.UpperLimit = %s", checkpointvortimeslicetypeIdent, targetIdent))
			}
			if checkpointvortimeslicetype.UpperLimitReference != nil {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := checkpointvortimeslicetype.UpperLimitReference.GongGetIdentifier(stage.AixmdatatypesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.UpperLimitReference = %s", checkpointvortimeslicetypeIdent, targetIdent))
			}
			if checkpointvortimeslicetype.LowerLimit != nil {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := checkpointvortimeslicetype.LowerLimit.GongGetIdentifier(stage.AixmdatatypesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.LowerLimit = %s", checkpointvortimeslicetypeIdent, targetIdent))
			}
			if checkpointvortimeslicetype.LowerLimitReference != nil {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := checkpointvortimeslicetype.LowerLimitReference.GongGetIdentifier(stage.AixmdatatypesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.LowerLimitReference = %s", checkpointvortimeslicetypeIdent, targetIdent))
			}
			if checkpointvortimeslicetype.AltitudeInterpretation != nil {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := checkpointvortimeslicetype.AltitudeInterpretation.GongGetIdentifier(stage.AixmdatatypesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.AltitudeInterpretation = %s", checkpointvortimeslicetypeIdent, targetIdent))
			}
			if checkpointvortimeslicetype.Distance != nil {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := checkpointvortimeslicetype.Distance.GongGetIdentifier(stage.AixmdatatypesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Distance = %s", checkpointvortimeslicetypeIdent, targetIdent))
			}
			if checkpointvortimeslicetype.Angle != nil {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := checkpointvortimeslicetype.Angle.GongGetIdentifier(stage.AixmdatatypesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Angle = %s", checkpointvortimeslicetypeIdent, targetIdent))
			}
			if checkpointvortimeslicetype.Position != nil {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := checkpointvortimeslicetype.Position.GongGetIdentifier(stage.AixmfeaturesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Position = %s", checkpointvortimeslicetypeIdent, targetIdent))
			}
			if checkpointvortimeslicetype.AirportHeliport != nil {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := checkpointvortimeslicetype.AirportHeliport.GongGetIdentifier(stage.AixmfeaturesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.AirportHeliport = %s", checkpointvortimeslicetypeIdent, targetIdent))
			}
			for _, elem := range checkpointvortimeslicetype.Annotation {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := elem.GongGetIdentifier(stage.AixmfeaturesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Annotation = append(%s.Annotation, %s)", checkpointvortimeslicetypeIdent, checkpointvortimeslicetypeIdent, targetIdent))
			}
			if checkpointvortimeslicetype.CheckPointFacility != nil {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := checkpointvortimeslicetype.CheckPointFacility.GongGetIdentifier(stage.AixmfeaturesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.CheckPointFacility = %s", checkpointvortimeslicetypeIdent, targetIdent))
			}
			for _, elem := range checkpointvortimeslicetype.Extension {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := elem.GongGetIdentifier(stage.AixmfeaturesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Extension = append(%s.Extension, %s)", checkpointvortimeslicetypeIdent, checkpointvortimeslicetypeIdent, targetIdent))
			}
		}
	}
	if stage.AixmfeaturesStage != nil {
		for _, checkpointvortype := range __gong__sortStageSetInstances(stage.AixmfeaturesStage.CheckpointVORTypes, stage.AixmfeaturesStage.CheckpointVORType_stagedOrder) {
			if lastStageDecl != "AixmfeaturesStage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "AixmfeaturesStage"
			}
			checkpointvortypeIdent := checkpointvortype.GongGetIdentifier(stage.AixmfeaturesStage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&aixmfeatures.CheckpointVORType{Name: %s}).Stage(stage.AixmfeaturesStage)", checkpointvortypeIdent, __gong__toRawStringLiteral(checkpointvortype.Name)))
			if lastStageVal != "AixmfeaturesStage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "AixmfeaturesStage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", checkpointvortypeIdent, __gong__toRawStringLiteral(checkpointvortype.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Id = %s", checkpointvortypeIdent, __gong__toRawStringLiteral(checkpointvortype.Id)))
			values.WriteString(fmt.Sprintf("\n\t%s.Description = %s", checkpointvortypeIdent, __gong__toRawStringLiteral(checkpointvortype.Description)))
			if checkpointvortype.Identifier != nil {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := checkpointvortype.Identifier.GongGetIdentifier(stage.GmlStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Identifier = %s", checkpointvortypeIdent, targetIdent))
			}
			for _, elem := range checkpointvortype.NameXSD {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := elem.GongGetIdentifier(stage.GmlStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.NameXSD = append(%s.NameXSD, %s)", checkpointvortypeIdent, checkpointvortypeIdent, targetIdent))
			}
			if checkpointvortype.BoundedBy != nil {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := checkpointvortype.BoundedBy.GongGetIdentifier(stage.GmlStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.BoundedBy = %s", checkpointvortypeIdent, targetIdent))
			}
			if checkpointvortype.FeatureMetadata != nil {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := checkpointvortype.FeatureMetadata.GongGetIdentifier(stage.AixmabstractgmlobjecttypesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.FeatureMetadata = %s", checkpointvortypeIdent, targetIdent))
			}
			for _, elem := range checkpointvortype.TimeSlice {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := elem.GongGetIdentifier(stage.AixmfeaturesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.TimeSlice = append(%s.TimeSlice, %s)", checkpointvortypeIdent, checkpointvortypeIdent, targetIdent))
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
		for _, dmetimeslicepropertytype := range __gong__sortStageSetInstances(stage.AixmfeaturesStage.DMETimeSlicePropertyTypes, stage.AixmfeaturesStage.DMETimeSlicePropertyType_stagedOrder) {
			if lastStageDecl != "AixmfeaturesStage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "AixmfeaturesStage"
			}
			dmetimeslicepropertytypeIdent := dmetimeslicepropertytype.GongGetIdentifier(stage.AixmfeaturesStage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&aixmfeatures.DMETimeSlicePropertyType{Name: %s}).Stage(stage.AixmfeaturesStage)", dmetimeslicepropertytypeIdent, __gong__toRawStringLiteral(dmetimeslicepropertytype.Name)))
			if lastStageVal != "AixmfeaturesStage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "AixmfeaturesStage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", dmetimeslicepropertytypeIdent, __gong__toRawStringLiteral(dmetimeslicepropertytype.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Owns = %t", dmetimeslicepropertytypeIdent, dmetimeslicepropertytype.Owns))
			if dmetimeslicepropertytype.DMETimeSlice != nil {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := dmetimeslicepropertytype.DMETimeSlice.GongGetIdentifier(stage.AixmfeaturesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.DMETimeSlice = %s", dmetimeslicepropertytypeIdent, targetIdent))
			}
		}
	}
	if stage.AixmfeaturesStage != nil {
		for _, dmetimeslicetype := range __gong__sortStageSetInstances(stage.AixmfeaturesStage.DMETimeSliceTypes, stage.AixmfeaturesStage.DMETimeSliceType_stagedOrder) {
			if lastStageDecl != "AixmfeaturesStage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "AixmfeaturesStage"
			}
			dmetimeslicetypeIdent := dmetimeslicetype.GongGetIdentifier(stage.AixmfeaturesStage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&aixmfeatures.DMETimeSliceType{Name: %s}).Stage(stage.AixmfeaturesStage)", dmetimeslicetypeIdent, __gong__toRawStringLiteral(dmetimeslicetype.Name)))
			if lastStageVal != "AixmfeaturesStage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "AixmfeaturesStage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", dmetimeslicetypeIdent, __gong__toRawStringLiteral(dmetimeslicetype.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Id = %s", dmetimeslicetypeIdent, __gong__toRawStringLiteral(dmetimeslicetype.Id)))
			values.WriteString(fmt.Sprintf("\n\t%s.Description = %s", dmetimeslicetypeIdent, __gong__toRawStringLiteral(dmetimeslicetype.Description)))
			values.WriteString(fmt.Sprintf("\n\t%s.Interpretation = %s", dmetimeslicetypeIdent, __gong__toRawStringLiteral(dmetimeslicetype.Interpretation)))
			values.WriteString(fmt.Sprintf("\n\t%s.SequenceNumber = %d", dmetimeslicetypeIdent, dmetimeslicetype.SequenceNumber))
			values.WriteString(fmt.Sprintf("\n\t%s.CorrectionNumber = %d", dmetimeslicetypeIdent, dmetimeslicetype.CorrectionNumber))
			if dmetimeslicetype.Identifier != nil {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := dmetimeslicetype.Identifier.GongGetIdentifier(stage.GmlStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Identifier = %s", dmetimeslicetypeIdent, targetIdent))
			}
			for _, elem := range dmetimeslicetype.NameXSD {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := elem.GongGetIdentifier(stage.GmlStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.NameXSD = append(%s.NameXSD, %s)", dmetimeslicetypeIdent, dmetimeslicetypeIdent, targetIdent))
			}
			if dmetimeslicetype.ValidTime != nil {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := dmetimeslicetype.ValidTime.GongGetIdentifier(stage.GmlStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.ValidTime = %s", dmetimeslicetypeIdent, targetIdent))
			}
			if dmetimeslicetype.TimeSliceMetadata != nil {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := dmetimeslicetype.TimeSliceMetadata.GongGetIdentifier(stage.AixmabstractgmlobjecttypesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.TimeSliceMetadata = %s", dmetimeslicetypeIdent, targetIdent))
			}
			if dmetimeslicetype.FeatureLifetime != nil {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := dmetimeslicetype.FeatureLifetime.GongGetIdentifier(stage.GmlStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.FeatureLifetime = %s", dmetimeslicetypeIdent, targetIdent))
			}
			if dmetimeslicetype.Designator != nil {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := dmetimeslicetype.Designator.GongGetIdentifier(stage.AixmdatatypesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Designator = %s", dmetimeslicetypeIdent, targetIdent))
			}
			if dmetimeslicetype.EmissionClass != nil {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := dmetimeslicetype.EmissionClass.GongGetIdentifier(stage.AixmdatatypesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.EmissionClass = %s", dmetimeslicetypeIdent, targetIdent))
			}
			if dmetimeslicetype.Mobile != nil {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := dmetimeslicetype.Mobile.GongGetIdentifier(stage.AixmdatatypesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Mobile = %s", dmetimeslicetypeIdent, targetIdent))
			}
			if dmetimeslicetype.MagneticVariation != nil {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := dmetimeslicetype.MagneticVariation.GongGetIdentifier(stage.AixmdatatypesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.MagneticVariation = %s", dmetimeslicetypeIdent, targetIdent))
			}
			if dmetimeslicetype.MagneticVariationAccuracy != nil {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := dmetimeslicetype.MagneticVariationAccuracy.GongGetIdentifier(stage.AixmdatatypesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.MagneticVariationAccuracy = %s", dmetimeslicetypeIdent, targetIdent))
			}
			if dmetimeslicetype.DateMagneticVariation != nil {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := dmetimeslicetype.DateMagneticVariation.GongGetIdentifier(stage.AixmdatatypesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.DateMagneticVariation = %s", dmetimeslicetypeIdent, targetIdent))
			}
			if dmetimeslicetype.FlightChecked != nil {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := dmetimeslicetype.FlightChecked.GongGetIdentifier(stage.AixmdatatypesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.FlightChecked = %s", dmetimeslicetypeIdent, targetIdent))
			}
			if dmetimeslicetype.Location != nil {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := dmetimeslicetype.Location.GongGetIdentifier(stage.AixmfeaturesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Location = %s", dmetimeslicetypeIdent, targetIdent))
			}
			for _, elem := range dmetimeslicetype.Authority {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := elem.GongGetIdentifier(stage.AixmfeaturesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Authority = append(%s.Authority, %s)", dmetimeslicetypeIdent, dmetimeslicetypeIdent, targetIdent))
			}
			for _, elem := range dmetimeslicetype.Monitoring {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := elem.GongGetIdentifier(stage.AixmfeaturesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Monitoring = append(%s.Monitoring, %s)", dmetimeslicetypeIdent, dmetimeslicetypeIdent, targetIdent))
			}
			for _, elem := range dmetimeslicetype.Availability {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := elem.GongGetIdentifier(stage.AixmfeaturesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Availability = append(%s.Availability, %s)", dmetimeslicetypeIdent, dmetimeslicetypeIdent, targetIdent))
			}
			for _, elem := range dmetimeslicetype.Annotation {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := elem.GongGetIdentifier(stage.AixmfeaturesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Annotation = append(%s.Annotation, %s)", dmetimeslicetypeIdent, dmetimeslicetypeIdent, targetIdent))
			}
			if dmetimeslicetype.Type_XSD != nil {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := dmetimeslicetype.Type_XSD.GongGetIdentifier(stage.AixmdatatypesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Type_XSD = %s", dmetimeslicetypeIdent, targetIdent))
			}
			if dmetimeslicetype.Channel != nil {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := dmetimeslicetype.Channel.GongGetIdentifier(stage.AixmdatatypesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Channel = %s", dmetimeslicetypeIdent, targetIdent))
			}
			if dmetimeslicetype.GhostFrequency != nil {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := dmetimeslicetype.GhostFrequency.GongGetIdentifier(stage.AixmdatatypesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.GhostFrequency = %s", dmetimeslicetypeIdent, targetIdent))
			}
			if dmetimeslicetype.Displace != nil {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := dmetimeslicetype.Displace.GongGetIdentifier(stage.AixmdatatypesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Displace = %s", dmetimeslicetypeIdent, targetIdent))
			}
			for _, elem := range dmetimeslicetype.Extension {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := elem.GongGetIdentifier(stage.AixmfeaturesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Extension = append(%s.Extension, %s)", dmetimeslicetypeIdent, dmetimeslicetypeIdent, targetIdent))
			}
		}
	}
	if stage.AixmfeaturesStage != nil {
		for _, dmetype := range __gong__sortStageSetInstances(stage.AixmfeaturesStage.DMETypes, stage.AixmfeaturesStage.DMEType_stagedOrder) {
			if lastStageDecl != "AixmfeaturesStage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "AixmfeaturesStage"
			}
			dmetypeIdent := dmetype.GongGetIdentifier(stage.AixmfeaturesStage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&aixmfeatures.DMEType{Name: %s}).Stage(stage.AixmfeaturesStage)", dmetypeIdent, __gong__toRawStringLiteral(dmetype.Name)))
			if lastStageVal != "AixmfeaturesStage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "AixmfeaturesStage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", dmetypeIdent, __gong__toRawStringLiteral(dmetype.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Id = %s", dmetypeIdent, __gong__toRawStringLiteral(dmetype.Id)))
			values.WriteString(fmt.Sprintf("\n\t%s.Description = %s", dmetypeIdent, __gong__toRawStringLiteral(dmetype.Description)))
			if dmetype.Identifier != nil {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := dmetype.Identifier.GongGetIdentifier(stage.GmlStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Identifier = %s", dmetypeIdent, targetIdent))
			}
			for _, elem := range dmetype.NameXSD {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := elem.GongGetIdentifier(stage.GmlStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.NameXSD = append(%s.NameXSD, %s)", dmetypeIdent, dmetypeIdent, targetIdent))
			}
			if dmetype.BoundedBy != nil {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := dmetype.BoundedBy.GongGetIdentifier(stage.GmlStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.BoundedBy = %s", dmetypeIdent, targetIdent))
			}
			if dmetype.FeatureMetadata != nil {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := dmetype.FeatureMetadata.GongGetIdentifier(stage.AixmabstractgmlobjecttypesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.FeatureMetadata = %s", dmetypeIdent, targetIdent))
			}
			for _, elem := range dmetype.TimeSlice {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := elem.GongGetIdentifier(stage.AixmfeaturesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.TimeSlice = append(%s.TimeSlice, %s)", dmetypeIdent, dmetypeIdent, targetIdent))
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
		for _, designatedpointtimeslicepropertytype := range __gong__sortStageSetInstances(stage.AixmfeaturesStage.DesignatedPointTimeSlicePropertyTypes, stage.AixmfeaturesStage.DesignatedPointTimeSlicePropertyType_stagedOrder) {
			if lastStageDecl != "AixmfeaturesStage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "AixmfeaturesStage"
			}
			designatedpointtimeslicepropertytypeIdent := designatedpointtimeslicepropertytype.GongGetIdentifier(stage.AixmfeaturesStage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&aixmfeatures.DesignatedPointTimeSlicePropertyType{Name: %s}).Stage(stage.AixmfeaturesStage)", designatedpointtimeslicepropertytypeIdent, __gong__toRawStringLiteral(designatedpointtimeslicepropertytype.Name)))
			if lastStageVal != "AixmfeaturesStage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "AixmfeaturesStage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", designatedpointtimeslicepropertytypeIdent, __gong__toRawStringLiteral(designatedpointtimeslicepropertytype.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Owns = %t", designatedpointtimeslicepropertytypeIdent, designatedpointtimeslicepropertytype.Owns))
			if designatedpointtimeslicepropertytype.DesignatedPointTimeSlice != nil {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := designatedpointtimeslicepropertytype.DesignatedPointTimeSlice.GongGetIdentifier(stage.AixmfeaturesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.DesignatedPointTimeSlice = %s", designatedpointtimeslicepropertytypeIdent, targetIdent))
			}
		}
	}
	if stage.AixmfeaturesStage != nil {
		for _, designatedpointtimeslicetype := range __gong__sortStageSetInstances(stage.AixmfeaturesStage.DesignatedPointTimeSliceTypes, stage.AixmfeaturesStage.DesignatedPointTimeSliceType_stagedOrder) {
			if lastStageDecl != "AixmfeaturesStage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "AixmfeaturesStage"
			}
			designatedpointtimeslicetypeIdent := designatedpointtimeslicetype.GongGetIdentifier(stage.AixmfeaturesStage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&aixmfeatures.DesignatedPointTimeSliceType{Name: %s}).Stage(stage.AixmfeaturesStage)", designatedpointtimeslicetypeIdent, __gong__toRawStringLiteral(designatedpointtimeslicetype.Name)))
			if lastStageVal != "AixmfeaturesStage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "AixmfeaturesStage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", designatedpointtimeslicetypeIdent, __gong__toRawStringLiteral(designatedpointtimeslicetype.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Id = %s", designatedpointtimeslicetypeIdent, __gong__toRawStringLiteral(designatedpointtimeslicetype.Id)))
			values.WriteString(fmt.Sprintf("\n\t%s.Description = %s", designatedpointtimeslicetypeIdent, __gong__toRawStringLiteral(designatedpointtimeslicetype.Description)))
			values.WriteString(fmt.Sprintf("\n\t%s.Interpretation = %s", designatedpointtimeslicetypeIdent, __gong__toRawStringLiteral(designatedpointtimeslicetype.Interpretation)))
			values.WriteString(fmt.Sprintf("\n\t%s.SequenceNumber = %d", designatedpointtimeslicetypeIdent, designatedpointtimeslicetype.SequenceNumber))
			values.WriteString(fmt.Sprintf("\n\t%s.CorrectionNumber = %d", designatedpointtimeslicetypeIdent, designatedpointtimeslicetype.CorrectionNumber))
			if designatedpointtimeslicetype.Identifier != nil {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := designatedpointtimeslicetype.Identifier.GongGetIdentifier(stage.GmlStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Identifier = %s", designatedpointtimeslicetypeIdent, targetIdent))
			}
			for _, elem := range designatedpointtimeslicetype.NameXSD {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := elem.GongGetIdentifier(stage.GmlStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.NameXSD = append(%s.NameXSD, %s)", designatedpointtimeslicetypeIdent, designatedpointtimeslicetypeIdent, targetIdent))
			}
			if designatedpointtimeslicetype.ValidTime != nil {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := designatedpointtimeslicetype.ValidTime.GongGetIdentifier(stage.GmlStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.ValidTime = %s", designatedpointtimeslicetypeIdent, targetIdent))
			}
			if designatedpointtimeslicetype.TimeSliceMetadata != nil {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := designatedpointtimeslicetype.TimeSliceMetadata.GongGetIdentifier(stage.AixmabstractgmlobjecttypesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.TimeSliceMetadata = %s", designatedpointtimeslicetypeIdent, targetIdent))
			}
			if designatedpointtimeslicetype.FeatureLifetime != nil {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := designatedpointtimeslicetype.FeatureLifetime.GongGetIdentifier(stage.GmlStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.FeatureLifetime = %s", designatedpointtimeslicetypeIdent, targetIdent))
			}
			if designatedpointtimeslicetype.Designator != nil {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := designatedpointtimeslicetype.Designator.GongGetIdentifier(stage.AixmdatatypesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Designator = %s", designatedpointtimeslicetypeIdent, targetIdent))
			}
			if designatedpointtimeslicetype.Type_XSD != nil {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := designatedpointtimeslicetype.Type_XSD.GongGetIdentifier(stage.AixmdatatypesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Type_XSD = %s", designatedpointtimeslicetypeIdent, targetIdent))
			}
			if designatedpointtimeslicetype.Location != nil {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := designatedpointtimeslicetype.Location.GongGetIdentifier(stage.AixmfeaturesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Location = %s", designatedpointtimeslicetypeIdent, targetIdent))
			}
			if designatedpointtimeslicetype.AimingPoint != nil {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := designatedpointtimeslicetype.AimingPoint.GongGetIdentifier(stage.AixmfeaturesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.AimingPoint = %s", designatedpointtimeslicetypeIdent, targetIdent))
			}
			if designatedpointtimeslicetype.AirportHeliport != nil {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := designatedpointtimeslicetype.AirportHeliport.GongGetIdentifier(stage.AixmfeaturesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.AirportHeliport = %s", designatedpointtimeslicetypeIdent, targetIdent))
			}
			if designatedpointtimeslicetype.RunwayPoint != nil {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := designatedpointtimeslicetype.RunwayPoint.GongGetIdentifier(stage.AixmfeaturesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.RunwayPoint = %s", designatedpointtimeslicetypeIdent, targetIdent))
			}
			for _, elem := range designatedpointtimeslicetype.Annotation {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := elem.GongGetIdentifier(stage.AixmfeaturesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Annotation = append(%s.Annotation, %s)", designatedpointtimeslicetypeIdent, designatedpointtimeslicetypeIdent, targetIdent))
			}
			for _, elem := range designatedpointtimeslicetype.Extension {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := elem.GongGetIdentifier(stage.AixmfeaturesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Extension = append(%s.Extension, %s)", designatedpointtimeslicetypeIdent, designatedpointtimeslicetypeIdent, targetIdent))
			}
		}
	}
	if stage.AixmfeaturesStage != nil {
		for _, designatedpointtype := range __gong__sortStageSetInstances(stage.AixmfeaturesStage.DesignatedPointTypes, stage.AixmfeaturesStage.DesignatedPointType_stagedOrder) {
			if lastStageDecl != "AixmfeaturesStage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "AixmfeaturesStage"
			}
			designatedpointtypeIdent := designatedpointtype.GongGetIdentifier(stage.AixmfeaturesStage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&aixmfeatures.DesignatedPointType{Name: %s}).Stage(stage.AixmfeaturesStage)", designatedpointtypeIdent, __gong__toRawStringLiteral(designatedpointtype.Name)))
			if lastStageVal != "AixmfeaturesStage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "AixmfeaturesStage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", designatedpointtypeIdent, __gong__toRawStringLiteral(designatedpointtype.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Id = %s", designatedpointtypeIdent, __gong__toRawStringLiteral(designatedpointtype.Id)))
			values.WriteString(fmt.Sprintf("\n\t%s.Description = %s", designatedpointtypeIdent, __gong__toRawStringLiteral(designatedpointtype.Description)))
			if designatedpointtype.Identifier != nil {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := designatedpointtype.Identifier.GongGetIdentifier(stage.GmlStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Identifier = %s", designatedpointtypeIdent, targetIdent))
			}
			for _, elem := range designatedpointtype.NameXSD {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := elem.GongGetIdentifier(stage.GmlStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.NameXSD = append(%s.NameXSD, %s)", designatedpointtypeIdent, designatedpointtypeIdent, targetIdent))
			}
			if designatedpointtype.BoundedBy != nil {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := designatedpointtype.BoundedBy.GongGetIdentifier(stage.GmlStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.BoundedBy = %s", designatedpointtypeIdent, targetIdent))
			}
			if designatedpointtype.FeatureMetadata != nil {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := designatedpointtype.FeatureMetadata.GongGetIdentifier(stage.AixmabstractgmlobjecttypesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.FeatureMetadata = %s", designatedpointtypeIdent, targetIdent))
			}
			for _, elem := range designatedpointtype.TimeSlice {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := elem.GongGetIdentifier(stage.AixmfeaturesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.TimeSlice = append(%s.TimeSlice, %s)", designatedpointtypeIdent, designatedpointtypeIdent, targetIdent))
			}
		}
	}
	if stage.AixmfeaturesStage != nil {
		for _, directionfindertimeslicepropertytype := range __gong__sortStageSetInstances(stage.AixmfeaturesStage.DirectionFinderTimeSlicePropertyTypes, stage.AixmfeaturesStage.DirectionFinderTimeSlicePropertyType_stagedOrder) {
			if lastStageDecl != "AixmfeaturesStage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "AixmfeaturesStage"
			}
			directionfindertimeslicepropertytypeIdent := directionfindertimeslicepropertytype.GongGetIdentifier(stage.AixmfeaturesStage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&aixmfeatures.DirectionFinderTimeSlicePropertyType{Name: %s}).Stage(stage.AixmfeaturesStage)", directionfindertimeslicepropertytypeIdent, __gong__toRawStringLiteral(directionfindertimeslicepropertytype.Name)))
			if lastStageVal != "AixmfeaturesStage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "AixmfeaturesStage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", directionfindertimeslicepropertytypeIdent, __gong__toRawStringLiteral(directionfindertimeslicepropertytype.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Owns = %t", directionfindertimeslicepropertytypeIdent, directionfindertimeslicepropertytype.Owns))
			if directionfindertimeslicepropertytype.DirectionFinderTimeSlice != nil {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := directionfindertimeslicepropertytype.DirectionFinderTimeSlice.GongGetIdentifier(stage.AixmfeaturesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.DirectionFinderTimeSlice = %s", directionfindertimeslicepropertytypeIdent, targetIdent))
			}
		}
	}
	if stage.AixmfeaturesStage != nil {
		for _, directionfindertimeslicetype := range __gong__sortStageSetInstances(stage.AixmfeaturesStage.DirectionFinderTimeSliceTypes, stage.AixmfeaturesStage.DirectionFinderTimeSliceType_stagedOrder) {
			if lastStageDecl != "AixmfeaturesStage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "AixmfeaturesStage"
			}
			directionfindertimeslicetypeIdent := directionfindertimeslicetype.GongGetIdentifier(stage.AixmfeaturesStage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&aixmfeatures.DirectionFinderTimeSliceType{Name: %s}).Stage(stage.AixmfeaturesStage)", directionfindertimeslicetypeIdent, __gong__toRawStringLiteral(directionfindertimeslicetype.Name)))
			if lastStageVal != "AixmfeaturesStage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "AixmfeaturesStage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", directionfindertimeslicetypeIdent, __gong__toRawStringLiteral(directionfindertimeslicetype.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Id = %s", directionfindertimeslicetypeIdent, __gong__toRawStringLiteral(directionfindertimeslicetype.Id)))
			values.WriteString(fmt.Sprintf("\n\t%s.Description = %s", directionfindertimeslicetypeIdent, __gong__toRawStringLiteral(directionfindertimeslicetype.Description)))
			values.WriteString(fmt.Sprintf("\n\t%s.Interpretation = %s", directionfindertimeslicetypeIdent, __gong__toRawStringLiteral(directionfindertimeslicetype.Interpretation)))
			values.WriteString(fmt.Sprintf("\n\t%s.SequenceNumber = %d", directionfindertimeslicetypeIdent, directionfindertimeslicetype.SequenceNumber))
			values.WriteString(fmt.Sprintf("\n\t%s.CorrectionNumber = %d", directionfindertimeslicetypeIdent, directionfindertimeslicetype.CorrectionNumber))
			if directionfindertimeslicetype.Identifier != nil {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := directionfindertimeslicetype.Identifier.GongGetIdentifier(stage.GmlStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Identifier = %s", directionfindertimeslicetypeIdent, targetIdent))
			}
			for _, elem := range directionfindertimeslicetype.NameXSD {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := elem.GongGetIdentifier(stage.GmlStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.NameXSD = append(%s.NameXSD, %s)", directionfindertimeslicetypeIdent, directionfindertimeslicetypeIdent, targetIdent))
			}
			if directionfindertimeslicetype.ValidTime != nil {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := directionfindertimeslicetype.ValidTime.GongGetIdentifier(stage.GmlStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.ValidTime = %s", directionfindertimeslicetypeIdent, targetIdent))
			}
			if directionfindertimeslicetype.TimeSliceMetadata != nil {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := directionfindertimeslicetype.TimeSliceMetadata.GongGetIdentifier(stage.AixmabstractgmlobjecttypesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.TimeSliceMetadata = %s", directionfindertimeslicetypeIdent, targetIdent))
			}
			if directionfindertimeslicetype.FeatureLifetime != nil {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := directionfindertimeslicetype.FeatureLifetime.GongGetIdentifier(stage.GmlStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.FeatureLifetime = %s", directionfindertimeslicetypeIdent, targetIdent))
			}
			if directionfindertimeslicetype.Designator != nil {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := directionfindertimeslicetype.Designator.GongGetIdentifier(stage.AixmdatatypesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Designator = %s", directionfindertimeslicetypeIdent, targetIdent))
			}
			if directionfindertimeslicetype.EmissionClass != nil {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := directionfindertimeslicetype.EmissionClass.GongGetIdentifier(stage.AixmdatatypesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.EmissionClass = %s", directionfindertimeslicetypeIdent, targetIdent))
			}
			if directionfindertimeslicetype.Mobile != nil {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := directionfindertimeslicetype.Mobile.GongGetIdentifier(stage.AixmdatatypesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Mobile = %s", directionfindertimeslicetypeIdent, targetIdent))
			}
			if directionfindertimeslicetype.MagneticVariation != nil {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := directionfindertimeslicetype.MagneticVariation.GongGetIdentifier(stage.AixmdatatypesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.MagneticVariation = %s", directionfindertimeslicetypeIdent, targetIdent))
			}
			if directionfindertimeslicetype.MagneticVariationAccuracy != nil {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := directionfindertimeslicetype.MagneticVariationAccuracy.GongGetIdentifier(stage.AixmdatatypesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.MagneticVariationAccuracy = %s", directionfindertimeslicetypeIdent, targetIdent))
			}
			if directionfindertimeslicetype.DateMagneticVariation != nil {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := directionfindertimeslicetype.DateMagneticVariation.GongGetIdentifier(stage.AixmdatatypesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.DateMagneticVariation = %s", directionfindertimeslicetypeIdent, targetIdent))
			}
			if directionfindertimeslicetype.FlightChecked != nil {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := directionfindertimeslicetype.FlightChecked.GongGetIdentifier(stage.AixmdatatypesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.FlightChecked = %s", directionfindertimeslicetypeIdent, targetIdent))
			}
			if directionfindertimeslicetype.Location != nil {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := directionfindertimeslicetype.Location.GongGetIdentifier(stage.AixmfeaturesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Location = %s", directionfindertimeslicetypeIdent, targetIdent))
			}
			for _, elem := range directionfindertimeslicetype.Authority {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := elem.GongGetIdentifier(stage.AixmfeaturesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Authority = append(%s.Authority, %s)", directionfindertimeslicetypeIdent, directionfindertimeslicetypeIdent, targetIdent))
			}
			for _, elem := range directionfindertimeslicetype.Monitoring {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := elem.GongGetIdentifier(stage.AixmfeaturesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Monitoring = append(%s.Monitoring, %s)", directionfindertimeslicetypeIdent, directionfindertimeslicetypeIdent, targetIdent))
			}
			for _, elem := range directionfindertimeslicetype.Availability {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := elem.GongGetIdentifier(stage.AixmfeaturesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Availability = append(%s.Availability, %s)", directionfindertimeslicetypeIdent, directionfindertimeslicetypeIdent, targetIdent))
			}
			for _, elem := range directionfindertimeslicetype.Annotation {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := elem.GongGetIdentifier(stage.AixmfeaturesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Annotation = append(%s.Annotation, %s)", directionfindertimeslicetypeIdent, directionfindertimeslicetypeIdent, targetIdent))
			}
			if directionfindertimeslicetype.Doppler != nil {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := directionfindertimeslicetype.Doppler.GongGetIdentifier(stage.AixmdatatypesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Doppler = %s", directionfindertimeslicetypeIdent, targetIdent))
			}
			for _, elem := range directionfindertimeslicetype.InformationProvision {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := elem.GongGetIdentifier(stage.AixmfeaturesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.InformationProvision = append(%s.InformationProvision, %s)", directionfindertimeslicetypeIdent, directionfindertimeslicetypeIdent, targetIdent))
			}
			for _, elem := range directionfindertimeslicetype.Extension {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := elem.GongGetIdentifier(stage.AixmfeaturesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Extension = append(%s.Extension, %s)", directionfindertimeslicetypeIdent, directionfindertimeslicetypeIdent, targetIdent))
			}
		}
	}
	if stage.AixmfeaturesStage != nil {
		for _, directionfindertype := range __gong__sortStageSetInstances(stage.AixmfeaturesStage.DirectionFinderTypes, stage.AixmfeaturesStage.DirectionFinderType_stagedOrder) {
			if lastStageDecl != "AixmfeaturesStage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "AixmfeaturesStage"
			}
			directionfindertypeIdent := directionfindertype.GongGetIdentifier(stage.AixmfeaturesStage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&aixmfeatures.DirectionFinderType{Name: %s}).Stage(stage.AixmfeaturesStage)", directionfindertypeIdent, __gong__toRawStringLiteral(directionfindertype.Name)))
			if lastStageVal != "AixmfeaturesStage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "AixmfeaturesStage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", directionfindertypeIdent, __gong__toRawStringLiteral(directionfindertype.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Id = %s", directionfindertypeIdent, __gong__toRawStringLiteral(directionfindertype.Id)))
			values.WriteString(fmt.Sprintf("\n\t%s.Description = %s", directionfindertypeIdent, __gong__toRawStringLiteral(directionfindertype.Description)))
			if directionfindertype.Identifier != nil {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := directionfindertype.Identifier.GongGetIdentifier(stage.GmlStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Identifier = %s", directionfindertypeIdent, targetIdent))
			}
			for _, elem := range directionfindertype.NameXSD {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := elem.GongGetIdentifier(stage.GmlStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.NameXSD = append(%s.NameXSD, %s)", directionfindertypeIdent, directionfindertypeIdent, targetIdent))
			}
			if directionfindertype.BoundedBy != nil {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := directionfindertype.BoundedBy.GongGetIdentifier(stage.GmlStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.BoundedBy = %s", directionfindertypeIdent, targetIdent))
			}
			if directionfindertype.FeatureMetadata != nil {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := directionfindertype.FeatureMetadata.GongGetIdentifier(stage.AixmabstractgmlobjecttypesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.FeatureMetadata = %s", directionfindertypeIdent, targetIdent))
			}
			for _, elem := range directionfindertype.TimeSlice {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := elem.GongGetIdentifier(stage.AixmfeaturesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.TimeSlice = append(%s.TimeSlice, %s)", directionfindertypeIdent, directionfindertypeIdent, targetIdent))
			}
		}
	}
	if stage.AixmfeaturesStage != nil {
		for _, distanceindicationtimeslicepropertytype := range __gong__sortStageSetInstances(stage.AixmfeaturesStage.DistanceIndicationTimeSlicePropertyTypes, stage.AixmfeaturesStage.DistanceIndicationTimeSlicePropertyType_stagedOrder) {
			if lastStageDecl != "AixmfeaturesStage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "AixmfeaturesStage"
			}
			distanceindicationtimeslicepropertytypeIdent := distanceindicationtimeslicepropertytype.GongGetIdentifier(stage.AixmfeaturesStage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&aixmfeatures.DistanceIndicationTimeSlicePropertyType{Name: %s}).Stage(stage.AixmfeaturesStage)", distanceindicationtimeslicepropertytypeIdent, __gong__toRawStringLiteral(distanceindicationtimeslicepropertytype.Name)))
			if lastStageVal != "AixmfeaturesStage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "AixmfeaturesStage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", distanceindicationtimeslicepropertytypeIdent, __gong__toRawStringLiteral(distanceindicationtimeslicepropertytype.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Owns = %t", distanceindicationtimeslicepropertytypeIdent, distanceindicationtimeslicepropertytype.Owns))
			if distanceindicationtimeslicepropertytype.DistanceIndicationTimeSlice != nil {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := distanceindicationtimeslicepropertytype.DistanceIndicationTimeSlice.GongGetIdentifier(stage.AixmfeaturesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.DistanceIndicationTimeSlice = %s", distanceindicationtimeslicepropertytypeIdent, targetIdent))
			}
		}
	}
	if stage.AixmfeaturesStage != nil {
		for _, distanceindicationtimeslicetype := range __gong__sortStageSetInstances(stage.AixmfeaturesStage.DistanceIndicationTimeSliceTypes, stage.AixmfeaturesStage.DistanceIndicationTimeSliceType_stagedOrder) {
			if lastStageDecl != "AixmfeaturesStage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "AixmfeaturesStage"
			}
			distanceindicationtimeslicetypeIdent := distanceindicationtimeslicetype.GongGetIdentifier(stage.AixmfeaturesStage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&aixmfeatures.DistanceIndicationTimeSliceType{Name: %s}).Stage(stage.AixmfeaturesStage)", distanceindicationtimeslicetypeIdent, __gong__toRawStringLiteral(distanceindicationtimeslicetype.Name)))
			if lastStageVal != "AixmfeaturesStage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "AixmfeaturesStage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", distanceindicationtimeslicetypeIdent, __gong__toRawStringLiteral(distanceindicationtimeslicetype.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Id = %s", distanceindicationtimeslicetypeIdent, __gong__toRawStringLiteral(distanceindicationtimeslicetype.Id)))
			values.WriteString(fmt.Sprintf("\n\t%s.Description = %s", distanceindicationtimeslicetypeIdent, __gong__toRawStringLiteral(distanceindicationtimeslicetype.Description)))
			values.WriteString(fmt.Sprintf("\n\t%s.Interpretation = %s", distanceindicationtimeslicetypeIdent, __gong__toRawStringLiteral(distanceindicationtimeslicetype.Interpretation)))
			values.WriteString(fmt.Sprintf("\n\t%s.SequenceNumber = %d", distanceindicationtimeslicetypeIdent, distanceindicationtimeslicetype.SequenceNumber))
			values.WriteString(fmt.Sprintf("\n\t%s.CorrectionNumber = %d", distanceindicationtimeslicetypeIdent, distanceindicationtimeslicetype.CorrectionNumber))
			if distanceindicationtimeslicetype.Identifier != nil {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := distanceindicationtimeslicetype.Identifier.GongGetIdentifier(stage.GmlStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Identifier = %s", distanceindicationtimeslicetypeIdent, targetIdent))
			}
			for _, elem := range distanceindicationtimeslicetype.NameXSD {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := elem.GongGetIdentifier(stage.GmlStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.NameXSD = append(%s.NameXSD, %s)", distanceindicationtimeslicetypeIdent, distanceindicationtimeslicetypeIdent, targetIdent))
			}
			if distanceindicationtimeslicetype.ValidTime != nil {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := distanceindicationtimeslicetype.ValidTime.GongGetIdentifier(stage.GmlStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.ValidTime = %s", distanceindicationtimeslicetypeIdent, targetIdent))
			}
			if distanceindicationtimeslicetype.TimeSliceMetadata != nil {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := distanceindicationtimeslicetype.TimeSliceMetadata.GongGetIdentifier(stage.AixmabstractgmlobjecttypesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.TimeSliceMetadata = %s", distanceindicationtimeslicetypeIdent, targetIdent))
			}
			if distanceindicationtimeslicetype.FeatureLifetime != nil {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := distanceindicationtimeslicetype.FeatureLifetime.GongGetIdentifier(stage.GmlStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.FeatureLifetime = %s", distanceindicationtimeslicetypeIdent, targetIdent))
			}
			if distanceindicationtimeslicetype.Distance != nil {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := distanceindicationtimeslicetype.Distance.GongGetIdentifier(stage.AixmdatatypesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Distance = %s", distanceindicationtimeslicetypeIdent, targetIdent))
			}
			if distanceindicationtimeslicetype.MinimumReceptionAltitude != nil {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := distanceindicationtimeslicetype.MinimumReceptionAltitude.GongGetIdentifier(stage.AixmdatatypesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.MinimumReceptionAltitude = %s", distanceindicationtimeslicetypeIdent, targetIdent))
			}
			if distanceindicationtimeslicetype.Type_XSD != nil {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := distanceindicationtimeslicetype.Type_XSD.GongGetIdentifier(stage.AixmdatatypesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Type_XSD = %s", distanceindicationtimeslicetypeIdent, targetIdent))
			}
			if distanceindicationtimeslicetype.Fix != nil {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := distanceindicationtimeslicetype.Fix.GongGetIdentifier(stage.AixmfeaturesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Fix = %s", distanceindicationtimeslicetypeIdent, targetIdent))
			}
			if distanceindicationtimeslicetype.PointChoice_fixDesignatedPoint != nil {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := distanceindicationtimeslicetype.PointChoice_fixDesignatedPoint.GongGetIdentifier(stage.AixmfeaturesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.PointChoice_fixDesignatedPoint = %s", distanceindicationtimeslicetypeIdent, targetIdent))
			}
			if distanceindicationtimeslicetype.PointChoice_navaidSystem != nil {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := distanceindicationtimeslicetype.PointChoice_navaidSystem.GongGetIdentifier(stage.AixmfeaturesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.PointChoice_navaidSystem = %s", distanceindicationtimeslicetypeIdent, targetIdent))
			}
			if distanceindicationtimeslicetype.PointChoice_position != nil {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := distanceindicationtimeslicetype.PointChoice_position.GongGetIdentifier(stage.AixmfeaturesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.PointChoice_position = %s", distanceindicationtimeslicetypeIdent, targetIdent))
			}
			if distanceindicationtimeslicetype.PointChoice_runwayPoint != nil {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := distanceindicationtimeslicetype.PointChoice_runwayPoint.GongGetIdentifier(stage.AixmfeaturesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.PointChoice_runwayPoint = %s", distanceindicationtimeslicetypeIdent, targetIdent))
			}
			if distanceindicationtimeslicetype.PointChoice_aimingPoint != nil {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := distanceindicationtimeslicetype.PointChoice_aimingPoint.GongGetIdentifier(stage.AixmfeaturesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.PointChoice_aimingPoint = %s", distanceindicationtimeslicetypeIdent, targetIdent))
			}
			if distanceindicationtimeslicetype.PointChoice_airportReferencePoint != nil {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := distanceindicationtimeslicetype.PointChoice_airportReferencePoint.GongGetIdentifier(stage.AixmfeaturesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.PointChoice_airportReferencePoint = %s", distanceindicationtimeslicetypeIdent, targetIdent))
			}
			for _, elem := range distanceindicationtimeslicetype.Annotation {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := elem.GongGetIdentifier(stage.AixmfeaturesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Annotation = append(%s.Annotation, %s)", distanceindicationtimeslicetypeIdent, distanceindicationtimeslicetypeIdent, targetIdent))
			}
			for _, elem := range distanceindicationtimeslicetype.Extension {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := elem.GongGetIdentifier(stage.AixmfeaturesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Extension = append(%s.Extension, %s)", distanceindicationtimeslicetypeIdent, distanceindicationtimeslicetypeIdent, targetIdent))
			}
		}
	}
	if stage.AixmfeaturesStage != nil {
		for _, distanceindicationtype := range __gong__sortStageSetInstances(stage.AixmfeaturesStage.DistanceIndicationTypes, stage.AixmfeaturesStage.DistanceIndicationType_stagedOrder) {
			if lastStageDecl != "AixmfeaturesStage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "AixmfeaturesStage"
			}
			distanceindicationtypeIdent := distanceindicationtype.GongGetIdentifier(stage.AixmfeaturesStage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&aixmfeatures.DistanceIndicationType{Name: %s}).Stage(stage.AixmfeaturesStage)", distanceindicationtypeIdent, __gong__toRawStringLiteral(distanceindicationtype.Name)))
			if lastStageVal != "AixmfeaturesStage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "AixmfeaturesStage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", distanceindicationtypeIdent, __gong__toRawStringLiteral(distanceindicationtype.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Id = %s", distanceindicationtypeIdent, __gong__toRawStringLiteral(distanceindicationtype.Id)))
			values.WriteString(fmt.Sprintf("\n\t%s.Description = %s", distanceindicationtypeIdent, __gong__toRawStringLiteral(distanceindicationtype.Description)))
			if distanceindicationtype.Identifier != nil {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := distanceindicationtype.Identifier.GongGetIdentifier(stage.GmlStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Identifier = %s", distanceindicationtypeIdent, targetIdent))
			}
			for _, elem := range distanceindicationtype.NameXSD {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := elem.GongGetIdentifier(stage.GmlStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.NameXSD = append(%s.NameXSD, %s)", distanceindicationtypeIdent, distanceindicationtypeIdent, targetIdent))
			}
			if distanceindicationtype.BoundedBy != nil {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := distanceindicationtype.BoundedBy.GongGetIdentifier(stage.GmlStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.BoundedBy = %s", distanceindicationtypeIdent, targetIdent))
			}
			if distanceindicationtype.FeatureMetadata != nil {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := distanceindicationtype.FeatureMetadata.GongGetIdentifier(stage.AixmabstractgmlobjecttypesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.FeatureMetadata = %s", distanceindicationtypeIdent, targetIdent))
			}
			for _, elem := range distanceindicationtype.TimeSlice {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := elem.GongGetIdentifier(stage.AixmfeaturesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.TimeSlice = append(%s.TimeSlice, %s)", distanceindicationtypeIdent, distanceindicationtypeIdent, targetIdent))
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
		for _, elevationtimeslicepropertytype := range __gong__sortStageSetInstances(stage.AixmfeaturesStage.ElevationTimeSlicePropertyTypes, stage.AixmfeaturesStage.ElevationTimeSlicePropertyType_stagedOrder) {
			if lastStageDecl != "AixmfeaturesStage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "AixmfeaturesStage"
			}
			elevationtimeslicepropertytypeIdent := elevationtimeslicepropertytype.GongGetIdentifier(stage.AixmfeaturesStage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&aixmfeatures.ElevationTimeSlicePropertyType{Name: %s}).Stage(stage.AixmfeaturesStage)", elevationtimeslicepropertytypeIdent, __gong__toRawStringLiteral(elevationtimeslicepropertytype.Name)))
			if lastStageVal != "AixmfeaturesStage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "AixmfeaturesStage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", elevationtimeslicepropertytypeIdent, __gong__toRawStringLiteral(elevationtimeslicepropertytype.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Owns = %t", elevationtimeslicepropertytypeIdent, elevationtimeslicepropertytype.Owns))
			if elevationtimeslicepropertytype.ElevationTimeSlice != nil {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := elevationtimeslicepropertytype.ElevationTimeSlice.GongGetIdentifier(stage.AixmfeaturesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.ElevationTimeSlice = %s", elevationtimeslicepropertytypeIdent, targetIdent))
			}
		}
	}
	if stage.AixmfeaturesStage != nil {
		for _, elevationtimeslicetype := range __gong__sortStageSetInstances(stage.AixmfeaturesStage.ElevationTimeSliceTypes, stage.AixmfeaturesStage.ElevationTimeSliceType_stagedOrder) {
			if lastStageDecl != "AixmfeaturesStage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "AixmfeaturesStage"
			}
			elevationtimeslicetypeIdent := elevationtimeslicetype.GongGetIdentifier(stage.AixmfeaturesStage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&aixmfeatures.ElevationTimeSliceType{Name: %s}).Stage(stage.AixmfeaturesStage)", elevationtimeslicetypeIdent, __gong__toRawStringLiteral(elevationtimeslicetype.Name)))
			if lastStageVal != "AixmfeaturesStage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "AixmfeaturesStage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", elevationtimeslicetypeIdent, __gong__toRawStringLiteral(elevationtimeslicetype.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Id = %s", elevationtimeslicetypeIdent, __gong__toRawStringLiteral(elevationtimeslicetype.Id)))
			values.WriteString(fmt.Sprintf("\n\t%s.Description = %s", elevationtimeslicetypeIdent, __gong__toRawStringLiteral(elevationtimeslicetype.Description)))
			values.WriteString(fmt.Sprintf("\n\t%s.Interpretation = %s", elevationtimeslicetypeIdent, __gong__toRawStringLiteral(elevationtimeslicetype.Interpretation)))
			values.WriteString(fmt.Sprintf("\n\t%s.SequenceNumber = %d", elevationtimeslicetypeIdent, elevationtimeslicetype.SequenceNumber))
			values.WriteString(fmt.Sprintf("\n\t%s.CorrectionNumber = %d", elevationtimeslicetypeIdent, elevationtimeslicetype.CorrectionNumber))
			if elevationtimeslicetype.Identifier != nil {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := elevationtimeslicetype.Identifier.GongGetIdentifier(stage.GmlStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Identifier = %s", elevationtimeslicetypeIdent, targetIdent))
			}
			for _, elem := range elevationtimeslicetype.NameXSD {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := elem.GongGetIdentifier(stage.GmlStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.NameXSD = append(%s.NameXSD, %s)", elevationtimeslicetypeIdent, elevationtimeslicetypeIdent, targetIdent))
			}
			if elevationtimeslicetype.ValidTime != nil {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := elevationtimeslicetype.ValidTime.GongGetIdentifier(stage.GmlStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.ValidTime = %s", elevationtimeslicetypeIdent, targetIdent))
			}
			if elevationtimeslicetype.TimeSliceMetadata != nil {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := elevationtimeslicetype.TimeSliceMetadata.GongGetIdentifier(stage.AixmabstractgmlobjecttypesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.TimeSliceMetadata = %s", elevationtimeslicetypeIdent, targetIdent))
			}
			if elevationtimeslicetype.FeatureLifetime != nil {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := elevationtimeslicetype.FeatureLifetime.GongGetIdentifier(stage.GmlStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.FeatureLifetime = %s", elevationtimeslicetypeIdent, targetIdent))
			}
			if elevationtimeslicetype.Designator != nil {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := elevationtimeslicetype.Designator.GongGetIdentifier(stage.AixmdatatypesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Designator = %s", elevationtimeslicetypeIdent, targetIdent))
			}
			if elevationtimeslicetype.EmissionClass != nil {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := elevationtimeslicetype.EmissionClass.GongGetIdentifier(stage.AixmdatatypesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.EmissionClass = %s", elevationtimeslicetypeIdent, targetIdent))
			}
			if elevationtimeslicetype.Mobile != nil {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := elevationtimeslicetype.Mobile.GongGetIdentifier(stage.AixmdatatypesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Mobile = %s", elevationtimeslicetypeIdent, targetIdent))
			}
			if elevationtimeslicetype.MagneticVariation != nil {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := elevationtimeslicetype.MagneticVariation.GongGetIdentifier(stage.AixmdatatypesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.MagneticVariation = %s", elevationtimeslicetypeIdent, targetIdent))
			}
			if elevationtimeslicetype.MagneticVariationAccuracy != nil {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := elevationtimeslicetype.MagneticVariationAccuracy.GongGetIdentifier(stage.AixmdatatypesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.MagneticVariationAccuracy = %s", elevationtimeslicetypeIdent, targetIdent))
			}
			if elevationtimeslicetype.DateMagneticVariation != nil {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := elevationtimeslicetype.DateMagneticVariation.GongGetIdentifier(stage.AixmdatatypesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.DateMagneticVariation = %s", elevationtimeslicetypeIdent, targetIdent))
			}
			if elevationtimeslicetype.FlightChecked != nil {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := elevationtimeslicetype.FlightChecked.GongGetIdentifier(stage.AixmdatatypesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.FlightChecked = %s", elevationtimeslicetypeIdent, targetIdent))
			}
			if elevationtimeslicetype.Location != nil {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := elevationtimeslicetype.Location.GongGetIdentifier(stage.AixmfeaturesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Location = %s", elevationtimeslicetypeIdent, targetIdent))
			}
			for _, elem := range elevationtimeslicetype.Authority {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := elem.GongGetIdentifier(stage.AixmfeaturesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Authority = append(%s.Authority, %s)", elevationtimeslicetypeIdent, elevationtimeslicetypeIdent, targetIdent))
			}
			for _, elem := range elevationtimeslicetype.Monitoring {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := elem.GongGetIdentifier(stage.AixmfeaturesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Monitoring = append(%s.Monitoring, %s)", elevationtimeslicetypeIdent, elevationtimeslicetypeIdent, targetIdent))
			}
			for _, elem := range elevationtimeslicetype.Availability {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := elem.GongGetIdentifier(stage.AixmfeaturesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Availability = append(%s.Availability, %s)", elevationtimeslicetypeIdent, elevationtimeslicetypeIdent, targetIdent))
			}
			for _, elem := range elevationtimeslicetype.Annotation {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := elem.GongGetIdentifier(stage.AixmfeaturesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Annotation = append(%s.Annotation, %s)", elevationtimeslicetypeIdent, elevationtimeslicetypeIdent, targetIdent))
			}
			if elevationtimeslicetype.AngleNominal != nil {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := elevationtimeslicetype.AngleNominal.GongGetIdentifier(stage.AixmdatatypesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.AngleNominal = %s", elevationtimeslicetypeIdent, targetIdent))
			}
			if elevationtimeslicetype.AngleMinimum != nil {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := elevationtimeslicetype.AngleMinimum.GongGetIdentifier(stage.AixmdatatypesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.AngleMinimum = %s", elevationtimeslicetypeIdent, targetIdent))
			}
			if elevationtimeslicetype.AngleSpan != nil {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := elevationtimeslicetype.AngleSpan.GongGetIdentifier(stage.AixmdatatypesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.AngleSpan = %s", elevationtimeslicetypeIdent, targetIdent))
			}
			if elevationtimeslicetype.AngleAccuracy != nil {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := elevationtimeslicetype.AngleAccuracy.GongGetIdentifier(stage.AixmdatatypesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.AngleAccuracy = %s", elevationtimeslicetypeIdent, targetIdent))
			}
			for _, elem := range elevationtimeslicetype.Extension {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := elem.GongGetIdentifier(stage.AixmfeaturesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Extension = append(%s.Extension, %s)", elevationtimeslicetypeIdent, elevationtimeslicetypeIdent, targetIdent))
			}
		}
	}
	if stage.AixmfeaturesStage != nil {
		for _, elevationtype := range __gong__sortStageSetInstances(stage.AixmfeaturesStage.ElevationTypes, stage.AixmfeaturesStage.ElevationType_stagedOrder) {
			if lastStageDecl != "AixmfeaturesStage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "AixmfeaturesStage"
			}
			elevationtypeIdent := elevationtype.GongGetIdentifier(stage.AixmfeaturesStage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&aixmfeatures.ElevationType{Name: %s}).Stage(stage.AixmfeaturesStage)", elevationtypeIdent, __gong__toRawStringLiteral(elevationtype.Name)))
			if lastStageVal != "AixmfeaturesStage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "AixmfeaturesStage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", elevationtypeIdent, __gong__toRawStringLiteral(elevationtype.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Id = %s", elevationtypeIdent, __gong__toRawStringLiteral(elevationtype.Id)))
			values.WriteString(fmt.Sprintf("\n\t%s.Description = %s", elevationtypeIdent, __gong__toRawStringLiteral(elevationtype.Description)))
			if elevationtype.Identifier != nil {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := elevationtype.Identifier.GongGetIdentifier(stage.GmlStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Identifier = %s", elevationtypeIdent, targetIdent))
			}
			for _, elem := range elevationtype.NameXSD {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := elem.GongGetIdentifier(stage.GmlStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.NameXSD = append(%s.NameXSD, %s)", elevationtypeIdent, elevationtypeIdent, targetIdent))
			}
			if elevationtype.BoundedBy != nil {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := elevationtype.BoundedBy.GongGetIdentifier(stage.GmlStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.BoundedBy = %s", elevationtypeIdent, targetIdent))
			}
			if elevationtype.FeatureMetadata != nil {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := elevationtype.FeatureMetadata.GongGetIdentifier(stage.AixmabstractgmlobjecttypesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.FeatureMetadata = %s", elevationtypeIdent, targetIdent))
			}
			for _, elem := range elevationtype.TimeSlice {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := elem.GongGetIdentifier(stage.AixmfeaturesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.TimeSlice = append(%s.TimeSlice, %s)", elevationtypeIdent, elevationtypeIdent, targetIdent))
			}
		}
	}
	if stage.AixmfeaturesStage != nil {
		for _, glidepathtimeslicepropertytype := range __gong__sortStageSetInstances(stage.AixmfeaturesStage.GlidepathTimeSlicePropertyTypes, stage.AixmfeaturesStage.GlidepathTimeSlicePropertyType_stagedOrder) {
			if lastStageDecl != "AixmfeaturesStage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "AixmfeaturesStage"
			}
			glidepathtimeslicepropertytypeIdent := glidepathtimeslicepropertytype.GongGetIdentifier(stage.AixmfeaturesStage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&aixmfeatures.GlidepathTimeSlicePropertyType{Name: %s}).Stage(stage.AixmfeaturesStage)", glidepathtimeslicepropertytypeIdent, __gong__toRawStringLiteral(glidepathtimeslicepropertytype.Name)))
			if lastStageVal != "AixmfeaturesStage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "AixmfeaturesStage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", glidepathtimeslicepropertytypeIdent, __gong__toRawStringLiteral(glidepathtimeslicepropertytype.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Owns = %t", glidepathtimeslicepropertytypeIdent, glidepathtimeslicepropertytype.Owns))
			if glidepathtimeslicepropertytype.GlidepathTimeSlice != nil {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := glidepathtimeslicepropertytype.GlidepathTimeSlice.GongGetIdentifier(stage.AixmfeaturesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.GlidepathTimeSlice = %s", glidepathtimeslicepropertytypeIdent, targetIdent))
			}
		}
	}
	if stage.AixmfeaturesStage != nil {
		for _, glidepathtimeslicetype := range __gong__sortStageSetInstances(stage.AixmfeaturesStage.GlidepathTimeSliceTypes, stage.AixmfeaturesStage.GlidepathTimeSliceType_stagedOrder) {
			if lastStageDecl != "AixmfeaturesStage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "AixmfeaturesStage"
			}
			glidepathtimeslicetypeIdent := glidepathtimeslicetype.GongGetIdentifier(stage.AixmfeaturesStage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&aixmfeatures.GlidepathTimeSliceType{Name: %s}).Stage(stage.AixmfeaturesStage)", glidepathtimeslicetypeIdent, __gong__toRawStringLiteral(glidepathtimeslicetype.Name)))
			if lastStageVal != "AixmfeaturesStage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "AixmfeaturesStage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", glidepathtimeslicetypeIdent, __gong__toRawStringLiteral(glidepathtimeslicetype.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Id = %s", glidepathtimeslicetypeIdent, __gong__toRawStringLiteral(glidepathtimeslicetype.Id)))
			values.WriteString(fmt.Sprintf("\n\t%s.Description = %s", glidepathtimeslicetypeIdent, __gong__toRawStringLiteral(glidepathtimeslicetype.Description)))
			values.WriteString(fmt.Sprintf("\n\t%s.Interpretation = %s", glidepathtimeslicetypeIdent, __gong__toRawStringLiteral(glidepathtimeslicetype.Interpretation)))
			values.WriteString(fmt.Sprintf("\n\t%s.SequenceNumber = %d", glidepathtimeslicetypeIdent, glidepathtimeslicetype.SequenceNumber))
			values.WriteString(fmt.Sprintf("\n\t%s.CorrectionNumber = %d", glidepathtimeslicetypeIdent, glidepathtimeslicetype.CorrectionNumber))
			if glidepathtimeslicetype.Identifier != nil {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := glidepathtimeslicetype.Identifier.GongGetIdentifier(stage.GmlStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Identifier = %s", glidepathtimeslicetypeIdent, targetIdent))
			}
			for _, elem := range glidepathtimeslicetype.NameXSD {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := elem.GongGetIdentifier(stage.GmlStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.NameXSD = append(%s.NameXSD, %s)", glidepathtimeslicetypeIdent, glidepathtimeslicetypeIdent, targetIdent))
			}
			if glidepathtimeslicetype.ValidTime != nil {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := glidepathtimeslicetype.ValidTime.GongGetIdentifier(stage.GmlStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.ValidTime = %s", glidepathtimeslicetypeIdent, targetIdent))
			}
			if glidepathtimeslicetype.TimeSliceMetadata != nil {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := glidepathtimeslicetype.TimeSliceMetadata.GongGetIdentifier(stage.AixmabstractgmlobjecttypesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.TimeSliceMetadata = %s", glidepathtimeslicetypeIdent, targetIdent))
			}
			if glidepathtimeslicetype.FeatureLifetime != nil {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := glidepathtimeslicetype.FeatureLifetime.GongGetIdentifier(stage.GmlStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.FeatureLifetime = %s", glidepathtimeslicetypeIdent, targetIdent))
			}
			if glidepathtimeslicetype.Designator != nil {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := glidepathtimeslicetype.Designator.GongGetIdentifier(stage.AixmdatatypesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Designator = %s", glidepathtimeslicetypeIdent, targetIdent))
			}
			if glidepathtimeslicetype.EmissionClass != nil {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := glidepathtimeslicetype.EmissionClass.GongGetIdentifier(stage.AixmdatatypesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.EmissionClass = %s", glidepathtimeslicetypeIdent, targetIdent))
			}
			if glidepathtimeslicetype.Mobile != nil {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := glidepathtimeslicetype.Mobile.GongGetIdentifier(stage.AixmdatatypesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Mobile = %s", glidepathtimeslicetypeIdent, targetIdent))
			}
			if glidepathtimeslicetype.MagneticVariation != nil {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := glidepathtimeslicetype.MagneticVariation.GongGetIdentifier(stage.AixmdatatypesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.MagneticVariation = %s", glidepathtimeslicetypeIdent, targetIdent))
			}
			if glidepathtimeslicetype.MagneticVariationAccuracy != nil {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := glidepathtimeslicetype.MagneticVariationAccuracy.GongGetIdentifier(stage.AixmdatatypesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.MagneticVariationAccuracy = %s", glidepathtimeslicetypeIdent, targetIdent))
			}
			if glidepathtimeslicetype.DateMagneticVariation != nil {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := glidepathtimeslicetype.DateMagneticVariation.GongGetIdentifier(stage.AixmdatatypesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.DateMagneticVariation = %s", glidepathtimeslicetypeIdent, targetIdent))
			}
			if glidepathtimeslicetype.FlightChecked != nil {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := glidepathtimeslicetype.FlightChecked.GongGetIdentifier(stage.AixmdatatypesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.FlightChecked = %s", glidepathtimeslicetypeIdent, targetIdent))
			}
			if glidepathtimeslicetype.Location != nil {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := glidepathtimeslicetype.Location.GongGetIdentifier(stage.AixmfeaturesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Location = %s", glidepathtimeslicetypeIdent, targetIdent))
			}
			for _, elem := range glidepathtimeslicetype.Authority {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := elem.GongGetIdentifier(stage.AixmfeaturesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Authority = append(%s.Authority, %s)", glidepathtimeslicetypeIdent, glidepathtimeslicetypeIdent, targetIdent))
			}
			for _, elem := range glidepathtimeslicetype.Monitoring {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := elem.GongGetIdentifier(stage.AixmfeaturesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Monitoring = append(%s.Monitoring, %s)", glidepathtimeslicetypeIdent, glidepathtimeslicetypeIdent, targetIdent))
			}
			for _, elem := range glidepathtimeslicetype.Availability {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := elem.GongGetIdentifier(stage.AixmfeaturesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Availability = append(%s.Availability, %s)", glidepathtimeslicetypeIdent, glidepathtimeslicetypeIdent, targetIdent))
			}
			for _, elem := range glidepathtimeslicetype.Annotation {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := elem.GongGetIdentifier(stage.AixmfeaturesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Annotation = append(%s.Annotation, %s)", glidepathtimeslicetypeIdent, glidepathtimeslicetypeIdent, targetIdent))
			}
			if glidepathtimeslicetype.Frequency != nil {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := glidepathtimeslicetype.Frequency.GongGetIdentifier(stage.AixmdatatypesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Frequency = %s", glidepathtimeslicetypeIdent, targetIdent))
			}
			if glidepathtimeslicetype.Slope != nil {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := glidepathtimeslicetype.Slope.GongGetIdentifier(stage.AixmdatatypesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Slope = %s", glidepathtimeslicetypeIdent, targetIdent))
			}
			if glidepathtimeslicetype.AngleAccuracy != nil {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := glidepathtimeslicetype.AngleAccuracy.GongGetIdentifier(stage.AixmdatatypesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.AngleAccuracy = %s", glidepathtimeslicetypeIdent, targetIdent))
			}
			if glidepathtimeslicetype.Rdh != nil {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := glidepathtimeslicetype.Rdh.GongGetIdentifier(stage.AixmdatatypesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Rdh = %s", glidepathtimeslicetypeIdent, targetIdent))
			}
			if glidepathtimeslicetype.RdhAccuracy != nil {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := glidepathtimeslicetype.RdhAccuracy.GongGetIdentifier(stage.AixmdatatypesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.RdhAccuracy = %s", glidepathtimeslicetypeIdent, targetIdent))
			}
			for _, elem := range glidepathtimeslicetype.Extension {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := elem.GongGetIdentifier(stage.AixmfeaturesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Extension = append(%s.Extension, %s)", glidepathtimeslicetypeIdent, glidepathtimeslicetypeIdent, targetIdent))
			}
		}
	}
	if stage.AixmfeaturesStage != nil {
		for _, glidepathtype := range __gong__sortStageSetInstances(stage.AixmfeaturesStage.GlidepathTypes, stage.AixmfeaturesStage.GlidepathType_stagedOrder) {
			if lastStageDecl != "AixmfeaturesStage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "AixmfeaturesStage"
			}
			glidepathtypeIdent := glidepathtype.GongGetIdentifier(stage.AixmfeaturesStage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&aixmfeatures.GlidepathType{Name: %s}).Stage(stage.AixmfeaturesStage)", glidepathtypeIdent, __gong__toRawStringLiteral(glidepathtype.Name)))
			if lastStageVal != "AixmfeaturesStage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "AixmfeaturesStage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", glidepathtypeIdent, __gong__toRawStringLiteral(glidepathtype.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Id = %s", glidepathtypeIdent, __gong__toRawStringLiteral(glidepathtype.Id)))
			values.WriteString(fmt.Sprintf("\n\t%s.Description = %s", glidepathtypeIdent, __gong__toRawStringLiteral(glidepathtype.Description)))
			if glidepathtype.Identifier != nil {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := glidepathtype.Identifier.GongGetIdentifier(stage.GmlStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Identifier = %s", glidepathtypeIdent, targetIdent))
			}
			for _, elem := range glidepathtype.NameXSD {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := elem.GongGetIdentifier(stage.GmlStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.NameXSD = append(%s.NameXSD, %s)", glidepathtypeIdent, glidepathtypeIdent, targetIdent))
			}
			if glidepathtype.BoundedBy != nil {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := glidepathtype.BoundedBy.GongGetIdentifier(stage.GmlStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.BoundedBy = %s", glidepathtypeIdent, targetIdent))
			}
			if glidepathtype.FeatureMetadata != nil {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := glidepathtype.FeatureMetadata.GongGetIdentifier(stage.AixmabstractgmlobjecttypesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.FeatureMetadata = %s", glidepathtypeIdent, targetIdent))
			}
			for _, elem := range glidepathtype.TimeSlice {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := elem.GongGetIdentifier(stage.AixmfeaturesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.TimeSlice = append(%s.TimeSlice, %s)", glidepathtypeIdent, glidepathtypeIdent, targetIdent))
			}
		}
	}
	if stage.AixmfeaturesStage != nil {
		for _, informationservicepropertytype := range __gong__sortStageSetInstances(stage.AixmfeaturesStage.InformationServicePropertyTypes, stage.AixmfeaturesStage.InformationServicePropertyType_stagedOrder) {
			if lastStageDecl != "AixmfeaturesStage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "AixmfeaturesStage"
			}
			informationservicepropertytypeIdent := informationservicepropertytype.GongGetIdentifier(stage.AixmfeaturesStage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&aixmfeatures.InformationServicePropertyType{Name: %s}).Stage(stage.AixmfeaturesStage)", informationservicepropertytypeIdent, __gong__toRawStringLiteral(informationservicepropertytype.Name)))
			if lastStageVal != "AixmfeaturesStage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "AixmfeaturesStage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", informationservicepropertytypeIdent, __gong__toRawStringLiteral(informationservicepropertytype.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Owns = %t", informationservicepropertytypeIdent, informationservicepropertytype.Owns))
			values.WriteString(fmt.Sprintf("\n\t%s.NilReason = %s", informationservicepropertytypeIdent, __gong__toRawStringLiteral(informationservicepropertytype.NilReason)))
			values.WriteString(fmt.Sprintf("\n\t%s.Type_XSD = %s", informationservicepropertytypeIdent, __gong__toRawStringLiteral(informationservicepropertytype.Type_XSD)))
			values.WriteString(fmt.Sprintf("\n\t%s.Href = %s", informationservicepropertytypeIdent, __gong__toRawStringLiteral(informationservicepropertytype.Href)))
			values.WriteString(fmt.Sprintf("\n\t%s.Role = %s", informationservicepropertytypeIdent, __gong__toRawStringLiteral(informationservicepropertytype.Role)))
			values.WriteString(fmt.Sprintf("\n\t%s.Arcrole = %s", informationservicepropertytypeIdent, __gong__toRawStringLiteral(informationservicepropertytype.Arcrole)))
			values.WriteString(fmt.Sprintf("\n\t%s.Title = %s", informationservicepropertytypeIdent, __gong__toRawStringLiteral(informationservicepropertytype.Title)))
			values.WriteString(fmt.Sprintf("\n\t%s.Show = %s", informationservicepropertytypeIdent, __gong__toRawStringLiteral(informationservicepropertytype.Show)))
			values.WriteString(fmt.Sprintf("\n\t%s.Actuate = %s", informationservicepropertytypeIdent, __gong__toRawStringLiteral(informationservicepropertytype.Actuate)))
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
		for _, localizertimeslicepropertytype := range __gong__sortStageSetInstances(stage.AixmfeaturesStage.LocalizerTimeSlicePropertyTypes, stage.AixmfeaturesStage.LocalizerTimeSlicePropertyType_stagedOrder) {
			if lastStageDecl != "AixmfeaturesStage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "AixmfeaturesStage"
			}
			localizertimeslicepropertytypeIdent := localizertimeslicepropertytype.GongGetIdentifier(stage.AixmfeaturesStage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&aixmfeatures.LocalizerTimeSlicePropertyType{Name: %s}).Stage(stage.AixmfeaturesStage)", localizertimeslicepropertytypeIdent, __gong__toRawStringLiteral(localizertimeslicepropertytype.Name)))
			if lastStageVal != "AixmfeaturesStage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "AixmfeaturesStage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", localizertimeslicepropertytypeIdent, __gong__toRawStringLiteral(localizertimeslicepropertytype.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Owns = %t", localizertimeslicepropertytypeIdent, localizertimeslicepropertytype.Owns))
			if localizertimeslicepropertytype.LocalizerTimeSlice != nil {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := localizertimeslicepropertytype.LocalizerTimeSlice.GongGetIdentifier(stage.AixmfeaturesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.LocalizerTimeSlice = %s", localizertimeslicepropertytypeIdent, targetIdent))
			}
		}
	}
	if stage.AixmfeaturesStage != nil {
		for _, localizertimeslicetype := range __gong__sortStageSetInstances(stage.AixmfeaturesStage.LocalizerTimeSliceTypes, stage.AixmfeaturesStage.LocalizerTimeSliceType_stagedOrder) {
			if lastStageDecl != "AixmfeaturesStage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "AixmfeaturesStage"
			}
			localizertimeslicetypeIdent := localizertimeslicetype.GongGetIdentifier(stage.AixmfeaturesStage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&aixmfeatures.LocalizerTimeSliceType{Name: %s}).Stage(stage.AixmfeaturesStage)", localizertimeslicetypeIdent, __gong__toRawStringLiteral(localizertimeslicetype.Name)))
			if lastStageVal != "AixmfeaturesStage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "AixmfeaturesStage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", localizertimeslicetypeIdent, __gong__toRawStringLiteral(localizertimeslicetype.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Id = %s", localizertimeslicetypeIdent, __gong__toRawStringLiteral(localizertimeslicetype.Id)))
			values.WriteString(fmt.Sprintf("\n\t%s.Description = %s", localizertimeslicetypeIdent, __gong__toRawStringLiteral(localizertimeslicetype.Description)))
			values.WriteString(fmt.Sprintf("\n\t%s.Interpretation = %s", localizertimeslicetypeIdent, __gong__toRawStringLiteral(localizertimeslicetype.Interpretation)))
			values.WriteString(fmt.Sprintf("\n\t%s.SequenceNumber = %d", localizertimeslicetypeIdent, localizertimeslicetype.SequenceNumber))
			values.WriteString(fmt.Sprintf("\n\t%s.CorrectionNumber = %d", localizertimeslicetypeIdent, localizertimeslicetype.CorrectionNumber))
			if localizertimeslicetype.Identifier != nil {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := localizertimeslicetype.Identifier.GongGetIdentifier(stage.GmlStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Identifier = %s", localizertimeslicetypeIdent, targetIdent))
			}
			for _, elem := range localizertimeslicetype.NameXSD {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := elem.GongGetIdentifier(stage.GmlStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.NameXSD = append(%s.NameXSD, %s)", localizertimeslicetypeIdent, localizertimeslicetypeIdent, targetIdent))
			}
			if localizertimeslicetype.ValidTime != nil {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := localizertimeslicetype.ValidTime.GongGetIdentifier(stage.GmlStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.ValidTime = %s", localizertimeslicetypeIdent, targetIdent))
			}
			if localizertimeslicetype.TimeSliceMetadata != nil {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := localizertimeslicetype.TimeSliceMetadata.GongGetIdentifier(stage.AixmabstractgmlobjecttypesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.TimeSliceMetadata = %s", localizertimeslicetypeIdent, targetIdent))
			}
			if localizertimeslicetype.FeatureLifetime != nil {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := localizertimeslicetype.FeatureLifetime.GongGetIdentifier(stage.GmlStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.FeatureLifetime = %s", localizertimeslicetypeIdent, targetIdent))
			}
			if localizertimeslicetype.Designator != nil {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := localizertimeslicetype.Designator.GongGetIdentifier(stage.AixmdatatypesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Designator = %s", localizertimeslicetypeIdent, targetIdent))
			}
			if localizertimeslicetype.EmissionClass != nil {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := localizertimeslicetype.EmissionClass.GongGetIdentifier(stage.AixmdatatypesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.EmissionClass = %s", localizertimeslicetypeIdent, targetIdent))
			}
			if localizertimeslicetype.Mobile != nil {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := localizertimeslicetype.Mobile.GongGetIdentifier(stage.AixmdatatypesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Mobile = %s", localizertimeslicetypeIdent, targetIdent))
			}
			if localizertimeslicetype.MagneticVariation != nil {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := localizertimeslicetype.MagneticVariation.GongGetIdentifier(stage.AixmdatatypesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.MagneticVariation = %s", localizertimeslicetypeIdent, targetIdent))
			}
			if localizertimeslicetype.MagneticVariationAccuracy != nil {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := localizertimeslicetype.MagneticVariationAccuracy.GongGetIdentifier(stage.AixmdatatypesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.MagneticVariationAccuracy = %s", localizertimeslicetypeIdent, targetIdent))
			}
			if localizertimeslicetype.DateMagneticVariation != nil {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := localizertimeslicetype.DateMagneticVariation.GongGetIdentifier(stage.AixmdatatypesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.DateMagneticVariation = %s", localizertimeslicetypeIdent, targetIdent))
			}
			if localizertimeslicetype.FlightChecked != nil {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := localizertimeslicetype.FlightChecked.GongGetIdentifier(stage.AixmdatatypesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.FlightChecked = %s", localizertimeslicetypeIdent, targetIdent))
			}
			if localizertimeslicetype.Location != nil {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := localizertimeslicetype.Location.GongGetIdentifier(stage.AixmfeaturesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Location = %s", localizertimeslicetypeIdent, targetIdent))
			}
			for _, elem := range localizertimeslicetype.Authority {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := elem.GongGetIdentifier(stage.AixmfeaturesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Authority = append(%s.Authority, %s)", localizertimeslicetypeIdent, localizertimeslicetypeIdent, targetIdent))
			}
			for _, elem := range localizertimeslicetype.Monitoring {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := elem.GongGetIdentifier(stage.AixmfeaturesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Monitoring = append(%s.Monitoring, %s)", localizertimeslicetypeIdent, localizertimeslicetypeIdent, targetIdent))
			}
			for _, elem := range localizertimeslicetype.Availability {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := elem.GongGetIdentifier(stage.AixmfeaturesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Availability = append(%s.Availability, %s)", localizertimeslicetypeIdent, localizertimeslicetypeIdent, targetIdent))
			}
			for _, elem := range localizertimeslicetype.Annotation {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := elem.GongGetIdentifier(stage.AixmfeaturesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Annotation = append(%s.Annotation, %s)", localizertimeslicetypeIdent, localizertimeslicetypeIdent, targetIdent))
			}
			if localizertimeslicetype.Frequency != nil {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := localizertimeslicetype.Frequency.GongGetIdentifier(stage.AixmdatatypesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Frequency = %s", localizertimeslicetypeIdent, targetIdent))
			}
			if localizertimeslicetype.MagneticBearing != nil {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := localizertimeslicetype.MagneticBearing.GongGetIdentifier(stage.AixmdatatypesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.MagneticBearing = %s", localizertimeslicetypeIdent, targetIdent))
			}
			if localizertimeslicetype.MagneticBearingAccuracy != nil {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := localizertimeslicetype.MagneticBearingAccuracy.GongGetIdentifier(stage.AixmdatatypesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.MagneticBearingAccuracy = %s", localizertimeslicetypeIdent, targetIdent))
			}
			if localizertimeslicetype.TrueBearing != nil {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := localizertimeslicetype.TrueBearing.GongGetIdentifier(stage.AixmdatatypesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.TrueBearing = %s", localizertimeslicetypeIdent, targetIdent))
			}
			if localizertimeslicetype.TrueBearingAccuracy != nil {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := localizertimeslicetype.TrueBearingAccuracy.GongGetIdentifier(stage.AixmdatatypesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.TrueBearingAccuracy = %s", localizertimeslicetypeIdent, targetIdent))
			}
			if localizertimeslicetype.Declination != nil {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := localizertimeslicetype.Declination.GongGetIdentifier(stage.AixmdatatypesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Declination = %s", localizertimeslicetypeIdent, targetIdent))
			}
			if localizertimeslicetype.WidthCourse != nil {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := localizertimeslicetype.WidthCourse.GongGetIdentifier(stage.AixmdatatypesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.WidthCourse = %s", localizertimeslicetypeIdent, targetIdent))
			}
			if localizertimeslicetype.WidthCourseAccuracy != nil {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := localizertimeslicetype.WidthCourseAccuracy.GongGetIdentifier(stage.AixmdatatypesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.WidthCourseAccuracy = %s", localizertimeslicetypeIdent, targetIdent))
			}
			if localizertimeslicetype.BackCourseUsable != nil {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := localizertimeslicetype.BackCourseUsable.GongGetIdentifier(stage.AixmdatatypesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.BackCourseUsable = %s", localizertimeslicetypeIdent, targetIdent))
			}
			for _, elem := range localizertimeslicetype.Extension {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := elem.GongGetIdentifier(stage.AixmfeaturesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Extension = append(%s.Extension, %s)", localizertimeslicetypeIdent, localizertimeslicetypeIdent, targetIdent))
			}
		}
	}
	if stage.AixmfeaturesStage != nil {
		for _, localizertype := range __gong__sortStageSetInstances(stage.AixmfeaturesStage.LocalizerTypes, stage.AixmfeaturesStage.LocalizerType_stagedOrder) {
			if lastStageDecl != "AixmfeaturesStage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "AixmfeaturesStage"
			}
			localizertypeIdent := localizertype.GongGetIdentifier(stage.AixmfeaturesStage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&aixmfeatures.LocalizerType{Name: %s}).Stage(stage.AixmfeaturesStage)", localizertypeIdent, __gong__toRawStringLiteral(localizertype.Name)))
			if lastStageVal != "AixmfeaturesStage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "AixmfeaturesStage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", localizertypeIdent, __gong__toRawStringLiteral(localizertype.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Id = %s", localizertypeIdent, __gong__toRawStringLiteral(localizertype.Id)))
			values.WriteString(fmt.Sprintf("\n\t%s.Description = %s", localizertypeIdent, __gong__toRawStringLiteral(localizertype.Description)))
			if localizertype.Identifier != nil {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := localizertype.Identifier.GongGetIdentifier(stage.GmlStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Identifier = %s", localizertypeIdent, targetIdent))
			}
			for _, elem := range localizertype.NameXSD {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := elem.GongGetIdentifier(stage.GmlStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.NameXSD = append(%s.NameXSD, %s)", localizertypeIdent, localizertypeIdent, targetIdent))
			}
			if localizertype.BoundedBy != nil {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := localizertype.BoundedBy.GongGetIdentifier(stage.GmlStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.BoundedBy = %s", localizertypeIdent, targetIdent))
			}
			if localizertype.FeatureMetadata != nil {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := localizertype.FeatureMetadata.GongGetIdentifier(stage.AixmabstractgmlobjecttypesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.FeatureMetadata = %s", localizertypeIdent, targetIdent))
			}
			for _, elem := range localizertype.TimeSlice {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := elem.GongGetIdentifier(stage.AixmfeaturesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.TimeSlice = append(%s.TimeSlice, %s)", localizertypeIdent, localizertypeIdent, targetIdent))
			}
		}
	}
	if stage.AixmfeaturesStage != nil {
		for _, markerbeacontimeslicepropertytype := range __gong__sortStageSetInstances(stage.AixmfeaturesStage.MarkerBeaconTimeSlicePropertyTypes, stage.AixmfeaturesStage.MarkerBeaconTimeSlicePropertyType_stagedOrder) {
			if lastStageDecl != "AixmfeaturesStage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "AixmfeaturesStage"
			}
			markerbeacontimeslicepropertytypeIdent := markerbeacontimeslicepropertytype.GongGetIdentifier(stage.AixmfeaturesStage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&aixmfeatures.MarkerBeaconTimeSlicePropertyType{Name: %s}).Stage(stage.AixmfeaturesStage)", markerbeacontimeslicepropertytypeIdent, __gong__toRawStringLiteral(markerbeacontimeslicepropertytype.Name)))
			if lastStageVal != "AixmfeaturesStage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "AixmfeaturesStage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", markerbeacontimeslicepropertytypeIdent, __gong__toRawStringLiteral(markerbeacontimeslicepropertytype.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Owns = %t", markerbeacontimeslicepropertytypeIdent, markerbeacontimeslicepropertytype.Owns))
			if markerbeacontimeslicepropertytype.MarkerBeaconTimeSlice != nil {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := markerbeacontimeslicepropertytype.MarkerBeaconTimeSlice.GongGetIdentifier(stage.AixmfeaturesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.MarkerBeaconTimeSlice = %s", markerbeacontimeslicepropertytypeIdent, targetIdent))
			}
		}
	}
	if stage.AixmfeaturesStage != nil {
		for _, markerbeacontimeslicetype := range __gong__sortStageSetInstances(stage.AixmfeaturesStage.MarkerBeaconTimeSliceTypes, stage.AixmfeaturesStage.MarkerBeaconTimeSliceType_stagedOrder) {
			if lastStageDecl != "AixmfeaturesStage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "AixmfeaturesStage"
			}
			markerbeacontimeslicetypeIdent := markerbeacontimeslicetype.GongGetIdentifier(stage.AixmfeaturesStage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&aixmfeatures.MarkerBeaconTimeSliceType{Name: %s}).Stage(stage.AixmfeaturesStage)", markerbeacontimeslicetypeIdent, __gong__toRawStringLiteral(markerbeacontimeslicetype.Name)))
			if lastStageVal != "AixmfeaturesStage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "AixmfeaturesStage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", markerbeacontimeslicetypeIdent, __gong__toRawStringLiteral(markerbeacontimeslicetype.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Id = %s", markerbeacontimeslicetypeIdent, __gong__toRawStringLiteral(markerbeacontimeslicetype.Id)))
			values.WriteString(fmt.Sprintf("\n\t%s.Description = %s", markerbeacontimeslicetypeIdent, __gong__toRawStringLiteral(markerbeacontimeslicetype.Description)))
			values.WriteString(fmt.Sprintf("\n\t%s.Interpretation = %s", markerbeacontimeslicetypeIdent, __gong__toRawStringLiteral(markerbeacontimeslicetype.Interpretation)))
			values.WriteString(fmt.Sprintf("\n\t%s.SequenceNumber = %d", markerbeacontimeslicetypeIdent, markerbeacontimeslicetype.SequenceNumber))
			values.WriteString(fmt.Sprintf("\n\t%s.CorrectionNumber = %d", markerbeacontimeslicetypeIdent, markerbeacontimeslicetype.CorrectionNumber))
			if markerbeacontimeslicetype.Identifier != nil {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := markerbeacontimeslicetype.Identifier.GongGetIdentifier(stage.GmlStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Identifier = %s", markerbeacontimeslicetypeIdent, targetIdent))
			}
			for _, elem := range markerbeacontimeslicetype.NameXSD {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := elem.GongGetIdentifier(stage.GmlStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.NameXSD = append(%s.NameXSD, %s)", markerbeacontimeslicetypeIdent, markerbeacontimeslicetypeIdent, targetIdent))
			}
			if markerbeacontimeslicetype.ValidTime != nil {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := markerbeacontimeslicetype.ValidTime.GongGetIdentifier(stage.GmlStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.ValidTime = %s", markerbeacontimeslicetypeIdent, targetIdent))
			}
			if markerbeacontimeslicetype.TimeSliceMetadata != nil {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := markerbeacontimeslicetype.TimeSliceMetadata.GongGetIdentifier(stage.AixmabstractgmlobjecttypesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.TimeSliceMetadata = %s", markerbeacontimeslicetypeIdent, targetIdent))
			}
			if markerbeacontimeslicetype.FeatureLifetime != nil {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := markerbeacontimeslicetype.FeatureLifetime.GongGetIdentifier(stage.GmlStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.FeatureLifetime = %s", markerbeacontimeslicetypeIdent, targetIdent))
			}
			if markerbeacontimeslicetype.Designator != nil {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := markerbeacontimeslicetype.Designator.GongGetIdentifier(stage.AixmdatatypesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Designator = %s", markerbeacontimeslicetypeIdent, targetIdent))
			}
			if markerbeacontimeslicetype.EmissionClass != nil {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := markerbeacontimeslicetype.EmissionClass.GongGetIdentifier(stage.AixmdatatypesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.EmissionClass = %s", markerbeacontimeslicetypeIdent, targetIdent))
			}
			if markerbeacontimeslicetype.Mobile != nil {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := markerbeacontimeslicetype.Mobile.GongGetIdentifier(stage.AixmdatatypesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Mobile = %s", markerbeacontimeslicetypeIdent, targetIdent))
			}
			if markerbeacontimeslicetype.MagneticVariation != nil {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := markerbeacontimeslicetype.MagneticVariation.GongGetIdentifier(stage.AixmdatatypesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.MagneticVariation = %s", markerbeacontimeslicetypeIdent, targetIdent))
			}
			if markerbeacontimeslicetype.MagneticVariationAccuracy != nil {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := markerbeacontimeslicetype.MagneticVariationAccuracy.GongGetIdentifier(stage.AixmdatatypesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.MagneticVariationAccuracy = %s", markerbeacontimeslicetypeIdent, targetIdent))
			}
			if markerbeacontimeslicetype.DateMagneticVariation != nil {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := markerbeacontimeslicetype.DateMagneticVariation.GongGetIdentifier(stage.AixmdatatypesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.DateMagneticVariation = %s", markerbeacontimeslicetypeIdent, targetIdent))
			}
			if markerbeacontimeslicetype.FlightChecked != nil {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := markerbeacontimeslicetype.FlightChecked.GongGetIdentifier(stage.AixmdatatypesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.FlightChecked = %s", markerbeacontimeslicetypeIdent, targetIdent))
			}
			if markerbeacontimeslicetype.Location != nil {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := markerbeacontimeslicetype.Location.GongGetIdentifier(stage.AixmfeaturesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Location = %s", markerbeacontimeslicetypeIdent, targetIdent))
			}
			for _, elem := range markerbeacontimeslicetype.Authority {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := elem.GongGetIdentifier(stage.AixmfeaturesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Authority = append(%s.Authority, %s)", markerbeacontimeslicetypeIdent, markerbeacontimeslicetypeIdent, targetIdent))
			}
			for _, elem := range markerbeacontimeslicetype.Monitoring {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := elem.GongGetIdentifier(stage.AixmfeaturesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Monitoring = append(%s.Monitoring, %s)", markerbeacontimeslicetypeIdent, markerbeacontimeslicetypeIdent, targetIdent))
			}
			for _, elem := range markerbeacontimeslicetype.Availability {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := elem.GongGetIdentifier(stage.AixmfeaturesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Availability = append(%s.Availability, %s)", markerbeacontimeslicetypeIdent, markerbeacontimeslicetypeIdent, targetIdent))
			}
			for _, elem := range markerbeacontimeslicetype.Annotation {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := elem.GongGetIdentifier(stage.AixmfeaturesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Annotation = append(%s.Annotation, %s)", markerbeacontimeslicetypeIdent, markerbeacontimeslicetypeIdent, targetIdent))
			}
			if markerbeacontimeslicetype.Class != nil {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := markerbeacontimeslicetype.Class.GongGetIdentifier(stage.AixmdatatypesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Class = %s", markerbeacontimeslicetypeIdent, targetIdent))
			}
			if markerbeacontimeslicetype.Frequency != nil {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := markerbeacontimeslicetype.Frequency.GongGetIdentifier(stage.AixmdatatypesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Frequency = %s", markerbeacontimeslicetypeIdent, targetIdent))
			}
			if markerbeacontimeslicetype.AxisBearing != nil {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := markerbeacontimeslicetype.AxisBearing.GongGetIdentifier(stage.AixmdatatypesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.AxisBearing = %s", markerbeacontimeslicetypeIdent, targetIdent))
			}
			if markerbeacontimeslicetype.AuralMorseCode != nil {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := markerbeacontimeslicetype.AuralMorseCode.GongGetIdentifier(stage.AixmdatatypesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.AuralMorseCode = %s", markerbeacontimeslicetypeIdent, targetIdent))
			}
			for _, elem := range markerbeacontimeslicetype.Extension {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := elem.GongGetIdentifier(stage.AixmfeaturesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Extension = append(%s.Extension, %s)", markerbeacontimeslicetypeIdent, markerbeacontimeslicetypeIdent, targetIdent))
			}
		}
	}
	if stage.AixmfeaturesStage != nil {
		for _, markerbeacontype := range __gong__sortStageSetInstances(stage.AixmfeaturesStage.MarkerBeaconTypes, stage.AixmfeaturesStage.MarkerBeaconType_stagedOrder) {
			if lastStageDecl != "AixmfeaturesStage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "AixmfeaturesStage"
			}
			markerbeacontypeIdent := markerbeacontype.GongGetIdentifier(stage.AixmfeaturesStage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&aixmfeatures.MarkerBeaconType{Name: %s}).Stage(stage.AixmfeaturesStage)", markerbeacontypeIdent, __gong__toRawStringLiteral(markerbeacontype.Name)))
			if lastStageVal != "AixmfeaturesStage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "AixmfeaturesStage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", markerbeacontypeIdent, __gong__toRawStringLiteral(markerbeacontype.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Id = %s", markerbeacontypeIdent, __gong__toRawStringLiteral(markerbeacontype.Id)))
			values.WriteString(fmt.Sprintf("\n\t%s.Description = %s", markerbeacontypeIdent, __gong__toRawStringLiteral(markerbeacontype.Description)))
			if markerbeacontype.Identifier != nil {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := markerbeacontype.Identifier.GongGetIdentifier(stage.GmlStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Identifier = %s", markerbeacontypeIdent, targetIdent))
			}
			for _, elem := range markerbeacontype.NameXSD {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := elem.GongGetIdentifier(stage.GmlStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.NameXSD = append(%s.NameXSD, %s)", markerbeacontypeIdent, markerbeacontypeIdent, targetIdent))
			}
			if markerbeacontype.BoundedBy != nil {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := markerbeacontype.BoundedBy.GongGetIdentifier(stage.GmlStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.BoundedBy = %s", markerbeacontypeIdent, targetIdent))
			}
			if markerbeacontype.FeatureMetadata != nil {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := markerbeacontype.FeatureMetadata.GongGetIdentifier(stage.AixmabstractgmlobjecttypesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.FeatureMetadata = %s", markerbeacontypeIdent, targetIdent))
			}
			for _, elem := range markerbeacontype.TimeSlice {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := elem.GongGetIdentifier(stage.AixmfeaturesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.TimeSlice = append(%s.TimeSlice, %s)", markerbeacontypeIdent, markerbeacontypeIdent, targetIdent))
			}
		}
	}
	if stage.AixmfeaturesStage != nil {
		for _, ndbtimeslicepropertytype := range __gong__sortStageSetInstances(stage.AixmfeaturesStage.NDBTimeSlicePropertyTypes, stage.AixmfeaturesStage.NDBTimeSlicePropertyType_stagedOrder) {
			if lastStageDecl != "AixmfeaturesStage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "AixmfeaturesStage"
			}
			ndbtimeslicepropertytypeIdent := ndbtimeslicepropertytype.GongGetIdentifier(stage.AixmfeaturesStage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&aixmfeatures.NDBTimeSlicePropertyType{Name: %s}).Stage(stage.AixmfeaturesStage)", ndbtimeslicepropertytypeIdent, __gong__toRawStringLiteral(ndbtimeslicepropertytype.Name)))
			if lastStageVal != "AixmfeaturesStage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "AixmfeaturesStage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", ndbtimeslicepropertytypeIdent, __gong__toRawStringLiteral(ndbtimeslicepropertytype.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Owns = %t", ndbtimeslicepropertytypeIdent, ndbtimeslicepropertytype.Owns))
			if ndbtimeslicepropertytype.NDBTimeSlice != nil {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := ndbtimeslicepropertytype.NDBTimeSlice.GongGetIdentifier(stage.AixmfeaturesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.NDBTimeSlice = %s", ndbtimeslicepropertytypeIdent, targetIdent))
			}
		}
	}
	if stage.AixmfeaturesStage != nil {
		for _, ndbtimeslicetype := range __gong__sortStageSetInstances(stage.AixmfeaturesStage.NDBTimeSliceTypes, stage.AixmfeaturesStage.NDBTimeSliceType_stagedOrder) {
			if lastStageDecl != "AixmfeaturesStage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "AixmfeaturesStage"
			}
			ndbtimeslicetypeIdent := ndbtimeslicetype.GongGetIdentifier(stage.AixmfeaturesStage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&aixmfeatures.NDBTimeSliceType{Name: %s}).Stage(stage.AixmfeaturesStage)", ndbtimeslicetypeIdent, __gong__toRawStringLiteral(ndbtimeslicetype.Name)))
			if lastStageVal != "AixmfeaturesStage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "AixmfeaturesStage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", ndbtimeslicetypeIdent, __gong__toRawStringLiteral(ndbtimeslicetype.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Id = %s", ndbtimeslicetypeIdent, __gong__toRawStringLiteral(ndbtimeslicetype.Id)))
			values.WriteString(fmt.Sprintf("\n\t%s.Description = %s", ndbtimeslicetypeIdent, __gong__toRawStringLiteral(ndbtimeslicetype.Description)))
			values.WriteString(fmt.Sprintf("\n\t%s.Interpretation = %s", ndbtimeslicetypeIdent, __gong__toRawStringLiteral(ndbtimeslicetype.Interpretation)))
			values.WriteString(fmt.Sprintf("\n\t%s.SequenceNumber = %d", ndbtimeslicetypeIdent, ndbtimeslicetype.SequenceNumber))
			values.WriteString(fmt.Sprintf("\n\t%s.CorrectionNumber = %d", ndbtimeslicetypeIdent, ndbtimeslicetype.CorrectionNumber))
			if ndbtimeslicetype.Identifier != nil {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := ndbtimeslicetype.Identifier.GongGetIdentifier(stage.GmlStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Identifier = %s", ndbtimeslicetypeIdent, targetIdent))
			}
			for _, elem := range ndbtimeslicetype.NameXSD {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := elem.GongGetIdentifier(stage.GmlStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.NameXSD = append(%s.NameXSD, %s)", ndbtimeslicetypeIdent, ndbtimeslicetypeIdent, targetIdent))
			}
			if ndbtimeslicetype.ValidTime != nil {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := ndbtimeslicetype.ValidTime.GongGetIdentifier(stage.GmlStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.ValidTime = %s", ndbtimeslicetypeIdent, targetIdent))
			}
			if ndbtimeslicetype.TimeSliceMetadata != nil {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := ndbtimeslicetype.TimeSliceMetadata.GongGetIdentifier(stage.AixmabstractgmlobjecttypesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.TimeSliceMetadata = %s", ndbtimeslicetypeIdent, targetIdent))
			}
			if ndbtimeslicetype.FeatureLifetime != nil {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := ndbtimeslicetype.FeatureLifetime.GongGetIdentifier(stage.GmlStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.FeatureLifetime = %s", ndbtimeslicetypeIdent, targetIdent))
			}
			if ndbtimeslicetype.Designator != nil {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := ndbtimeslicetype.Designator.GongGetIdentifier(stage.AixmdatatypesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Designator = %s", ndbtimeslicetypeIdent, targetIdent))
			}
			if ndbtimeslicetype.EmissionClass != nil {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := ndbtimeslicetype.EmissionClass.GongGetIdentifier(stage.AixmdatatypesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.EmissionClass = %s", ndbtimeslicetypeIdent, targetIdent))
			}
			if ndbtimeslicetype.Mobile != nil {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := ndbtimeslicetype.Mobile.GongGetIdentifier(stage.AixmdatatypesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Mobile = %s", ndbtimeslicetypeIdent, targetIdent))
			}
			if ndbtimeslicetype.MagneticVariation != nil {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := ndbtimeslicetype.MagneticVariation.GongGetIdentifier(stage.AixmdatatypesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.MagneticVariation = %s", ndbtimeslicetypeIdent, targetIdent))
			}
			if ndbtimeslicetype.MagneticVariationAccuracy != nil {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := ndbtimeslicetype.MagneticVariationAccuracy.GongGetIdentifier(stage.AixmdatatypesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.MagneticVariationAccuracy = %s", ndbtimeslicetypeIdent, targetIdent))
			}
			if ndbtimeslicetype.DateMagneticVariation != nil {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := ndbtimeslicetype.DateMagneticVariation.GongGetIdentifier(stage.AixmdatatypesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.DateMagneticVariation = %s", ndbtimeslicetypeIdent, targetIdent))
			}
			if ndbtimeslicetype.FlightChecked != nil {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := ndbtimeslicetype.FlightChecked.GongGetIdentifier(stage.AixmdatatypesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.FlightChecked = %s", ndbtimeslicetypeIdent, targetIdent))
			}
			if ndbtimeslicetype.Location != nil {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := ndbtimeslicetype.Location.GongGetIdentifier(stage.AixmfeaturesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Location = %s", ndbtimeslicetypeIdent, targetIdent))
			}
			for _, elem := range ndbtimeslicetype.Authority {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := elem.GongGetIdentifier(stage.AixmfeaturesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Authority = append(%s.Authority, %s)", ndbtimeslicetypeIdent, ndbtimeslicetypeIdent, targetIdent))
			}
			for _, elem := range ndbtimeslicetype.Monitoring {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := elem.GongGetIdentifier(stage.AixmfeaturesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Monitoring = append(%s.Monitoring, %s)", ndbtimeslicetypeIdent, ndbtimeslicetypeIdent, targetIdent))
			}
			for _, elem := range ndbtimeslicetype.Availability {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := elem.GongGetIdentifier(stage.AixmfeaturesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Availability = append(%s.Availability, %s)", ndbtimeslicetypeIdent, ndbtimeslicetypeIdent, targetIdent))
			}
			for _, elem := range ndbtimeslicetype.Annotation {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := elem.GongGetIdentifier(stage.AixmfeaturesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Annotation = append(%s.Annotation, %s)", ndbtimeslicetypeIdent, ndbtimeslicetypeIdent, targetIdent))
			}
			if ndbtimeslicetype.Frequency != nil {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := ndbtimeslicetype.Frequency.GongGetIdentifier(stage.AixmdatatypesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Frequency = %s", ndbtimeslicetypeIdent, targetIdent))
			}
			if ndbtimeslicetype.Class != nil {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := ndbtimeslicetype.Class.GongGetIdentifier(stage.AixmdatatypesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Class = %s", ndbtimeslicetypeIdent, targetIdent))
			}
			if ndbtimeslicetype.EmissionBand != nil {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := ndbtimeslicetype.EmissionBand.GongGetIdentifier(stage.AixmdatatypesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.EmissionBand = %s", ndbtimeslicetypeIdent, targetIdent))
			}
			for _, elem := range ndbtimeslicetype.Extension {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := elem.GongGetIdentifier(stage.AixmfeaturesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Extension = append(%s.Extension, %s)", ndbtimeslicetypeIdent, ndbtimeslicetypeIdent, targetIdent))
			}
		}
	}
	if stage.AixmfeaturesStage != nil {
		for _, ndbtype := range __gong__sortStageSetInstances(stage.AixmfeaturesStage.NDBTypes, stage.AixmfeaturesStage.NDBType_stagedOrder) {
			if lastStageDecl != "AixmfeaturesStage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "AixmfeaturesStage"
			}
			ndbtypeIdent := ndbtype.GongGetIdentifier(stage.AixmfeaturesStage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&aixmfeatures.NDBType{Name: %s}).Stage(stage.AixmfeaturesStage)", ndbtypeIdent, __gong__toRawStringLiteral(ndbtype.Name)))
			if lastStageVal != "AixmfeaturesStage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "AixmfeaturesStage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", ndbtypeIdent, __gong__toRawStringLiteral(ndbtype.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Id = %s", ndbtypeIdent, __gong__toRawStringLiteral(ndbtype.Id)))
			values.WriteString(fmt.Sprintf("\n\t%s.Description = %s", ndbtypeIdent, __gong__toRawStringLiteral(ndbtype.Description)))
			if ndbtype.Identifier != nil {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := ndbtype.Identifier.GongGetIdentifier(stage.GmlStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Identifier = %s", ndbtypeIdent, targetIdent))
			}
			for _, elem := range ndbtype.NameXSD {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := elem.GongGetIdentifier(stage.GmlStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.NameXSD = append(%s.NameXSD, %s)", ndbtypeIdent, ndbtypeIdent, targetIdent))
			}
			if ndbtype.BoundedBy != nil {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := ndbtype.BoundedBy.GongGetIdentifier(stage.GmlStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.BoundedBy = %s", ndbtypeIdent, targetIdent))
			}
			if ndbtype.FeatureMetadata != nil {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := ndbtype.FeatureMetadata.GongGetIdentifier(stage.AixmabstractgmlobjecttypesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.FeatureMetadata = %s", ndbtypeIdent, targetIdent))
			}
			for _, elem := range ndbtype.TimeSlice {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := elem.GongGetIdentifier(stage.AixmfeaturesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.TimeSlice = append(%s.TimeSlice, %s)", ndbtypeIdent, ndbtypeIdent, targetIdent))
			}
		}
	}
	if stage.AixmfeaturesStage != nil {
		for _, navaidcomponentpropertytype := range __gong__sortStageSetInstances(stage.AixmfeaturesStage.NavaidComponentPropertyTypes, stage.AixmfeaturesStage.NavaidComponentPropertyType_stagedOrder) {
			if lastStageDecl != "AixmfeaturesStage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "AixmfeaturesStage"
			}
			navaidcomponentpropertytypeIdent := navaidcomponentpropertytype.GongGetIdentifier(stage.AixmfeaturesStage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&aixmfeatures.NavaidComponentPropertyType{Name: %s}).Stage(stage.AixmfeaturesStage)", navaidcomponentpropertytypeIdent, __gong__toRawStringLiteral(navaidcomponentpropertytype.Name)))
			if lastStageVal != "AixmfeaturesStage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "AixmfeaturesStage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", navaidcomponentpropertytypeIdent, __gong__toRawStringLiteral(navaidcomponentpropertytype.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.NilReason = %s", navaidcomponentpropertytypeIdent, __gong__toRawStringLiteral(navaidcomponentpropertytype.NilReason)))
			if navaidcomponentpropertytype.NavaidComponent != nil {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := navaidcomponentpropertytype.NavaidComponent.GongGetIdentifier(stage.AixmfeaturesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.NavaidComponent = %s", navaidcomponentpropertytypeIdent, targetIdent))
			}
		}
	}
	if stage.AixmfeaturesStage != nil {
		for _, navaidcomponenttype := range __gong__sortStageSetInstances(stage.AixmfeaturesStage.NavaidComponentTypes, stage.AixmfeaturesStage.NavaidComponentType_stagedOrder) {
			if lastStageDecl != "AixmfeaturesStage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "AixmfeaturesStage"
			}
			navaidcomponenttypeIdent := navaidcomponenttype.GongGetIdentifier(stage.AixmfeaturesStage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&aixmfeatures.NavaidComponentType{Name: %s}).Stage(stage.AixmfeaturesStage)", navaidcomponenttypeIdent, __gong__toRawStringLiteral(navaidcomponenttype.Name)))
			if lastStageVal != "AixmfeaturesStage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "AixmfeaturesStage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", navaidcomponenttypeIdent, __gong__toRawStringLiteral(navaidcomponenttype.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Id = %s", navaidcomponenttypeIdent, __gong__toRawStringLiteral(navaidcomponenttype.Id)))
			values.WriteString(fmt.Sprintf("\n\t%s.Description = %s", navaidcomponenttypeIdent, __gong__toRawStringLiteral(navaidcomponenttype.Description)))
			if navaidcomponenttype.Identifier != nil {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := navaidcomponenttype.Identifier.GongGetIdentifier(stage.GmlStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Identifier = %s", navaidcomponenttypeIdent, targetIdent))
			}
			for _, elem := range navaidcomponenttype.NameXSD {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := elem.GongGetIdentifier(stage.GmlStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.NameXSD = append(%s.NameXSD, %s)", navaidcomponenttypeIdent, navaidcomponenttypeIdent, targetIdent))
			}
			if navaidcomponenttype.CollocationGroup != nil {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := navaidcomponenttype.CollocationGroup.GongGetIdentifier(stage.AixmdatatypesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.CollocationGroup = %s", navaidcomponenttypeIdent, targetIdent))
			}
			if navaidcomponenttype.MarkerPosition != nil {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := navaidcomponenttype.MarkerPosition.GongGetIdentifier(stage.AixmdatatypesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.MarkerPosition = %s", navaidcomponenttypeIdent, targetIdent))
			}
			if navaidcomponenttype.ProvidesNavigableLocation != nil {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := navaidcomponenttype.ProvidesNavigableLocation.GongGetIdentifier(stage.AixmdatatypesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.ProvidesNavigableLocation = %s", navaidcomponenttypeIdent, targetIdent))
			}
			for _, elem := range navaidcomponenttype.Annotation {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := elem.GongGetIdentifier(stage.AixmfeaturesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Annotation = append(%s.Annotation, %s)", navaidcomponenttypeIdent, navaidcomponenttypeIdent, targetIdent))
			}
			if navaidcomponenttype.TheNavaidEquipment != nil {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := navaidcomponenttype.TheNavaidEquipment.GongGetIdentifier(stage.AixmfeaturesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.TheNavaidEquipment = %s", navaidcomponenttypeIdent, targetIdent))
			}
			for _, elem := range navaidcomponenttype.Extension {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := elem.GongGetIdentifier(stage.AixmfeaturesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Extension = append(%s.Extension, %s)", navaidcomponenttypeIdent, navaidcomponenttypeIdent, targetIdent))
			}
		}
	}
	if stage.AixmfeaturesStage != nil {
		for _, navaidequipmentmonitoringpropertytype := range __gong__sortStageSetInstances(stage.AixmfeaturesStage.NavaidEquipmentMonitoringPropertyTypes, stage.AixmfeaturesStage.NavaidEquipmentMonitoringPropertyType_stagedOrder) {
			if lastStageDecl != "AixmfeaturesStage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "AixmfeaturesStage"
			}
			navaidequipmentmonitoringpropertytypeIdent := navaidequipmentmonitoringpropertytype.GongGetIdentifier(stage.AixmfeaturesStage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&aixmfeatures.NavaidEquipmentMonitoringPropertyType{Name: %s}).Stage(stage.AixmfeaturesStage)", navaidequipmentmonitoringpropertytypeIdent, __gong__toRawStringLiteral(navaidequipmentmonitoringpropertytype.Name)))
			if lastStageVal != "AixmfeaturesStage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "AixmfeaturesStage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", navaidequipmentmonitoringpropertytypeIdent, __gong__toRawStringLiteral(navaidequipmentmonitoringpropertytype.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.NilReason = %s", navaidequipmentmonitoringpropertytypeIdent, __gong__toRawStringLiteral(navaidequipmentmonitoringpropertytype.NilReason)))
			if navaidequipmentmonitoringpropertytype.NavaidEquipmentMonitoring != nil {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := navaidequipmentmonitoringpropertytype.NavaidEquipmentMonitoring.GongGetIdentifier(stage.AixmfeaturesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.NavaidEquipmentMonitoring = %s", navaidequipmentmonitoringpropertytypeIdent, targetIdent))
			}
		}
	}
	if stage.AixmfeaturesStage != nil {
		for _, navaidequipmentmonitoringtype := range __gong__sortStageSetInstances(stage.AixmfeaturesStage.NavaidEquipmentMonitoringTypes, stage.AixmfeaturesStage.NavaidEquipmentMonitoringType_stagedOrder) {
			if lastStageDecl != "AixmfeaturesStage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "AixmfeaturesStage"
			}
			navaidequipmentmonitoringtypeIdent := navaidequipmentmonitoringtype.GongGetIdentifier(stage.AixmfeaturesStage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&aixmfeatures.NavaidEquipmentMonitoringType{Name: %s}).Stage(stage.AixmfeaturesStage)", navaidequipmentmonitoringtypeIdent, __gong__toRawStringLiteral(navaidequipmentmonitoringtype.Name)))
			if lastStageVal != "AixmfeaturesStage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "AixmfeaturesStage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", navaidequipmentmonitoringtypeIdent, __gong__toRawStringLiteral(navaidequipmentmonitoringtype.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Id = %s", navaidequipmentmonitoringtypeIdent, __gong__toRawStringLiteral(navaidequipmentmonitoringtype.Id)))
			values.WriteString(fmt.Sprintf("\n\t%s.Description = %s", navaidequipmentmonitoringtypeIdent, __gong__toRawStringLiteral(navaidequipmentmonitoringtype.Description)))
			if navaidequipmentmonitoringtype.Identifier != nil {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := navaidequipmentmonitoringtype.Identifier.GongGetIdentifier(stage.GmlStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Identifier = %s", navaidequipmentmonitoringtypeIdent, targetIdent))
			}
			for _, elem := range navaidequipmentmonitoringtype.NameXSD {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := elem.GongGetIdentifier(stage.GmlStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.NameXSD = append(%s.NameXSD, %s)", navaidequipmentmonitoringtypeIdent, navaidequipmentmonitoringtypeIdent, targetIdent))
			}
			for _, elem := range navaidequipmentmonitoringtype.TimeInterval {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := elem.GongGetIdentifier(stage.AixmfeaturesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.TimeInterval = append(%s.TimeInterval, %s)", navaidequipmentmonitoringtypeIdent, navaidequipmentmonitoringtypeIdent, targetIdent))
			}
			for _, elem := range navaidequipmentmonitoringtype.Annotation {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := elem.GongGetIdentifier(stage.AixmfeaturesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Annotation = append(%s.Annotation, %s)", navaidequipmentmonitoringtypeIdent, navaidequipmentmonitoringtypeIdent, targetIdent))
			}
			for _, elem := range navaidequipmentmonitoringtype.SpecialDateAuthority {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := elem.GongGetIdentifier(stage.AixmfeaturesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.SpecialDateAuthority = append(%s.SpecialDateAuthority, %s)", navaidequipmentmonitoringtypeIdent, navaidequipmentmonitoringtypeIdent, targetIdent))
			}
			if navaidequipmentmonitoringtype.Monitored != nil {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := navaidequipmentmonitoringtype.Monitored.GongGetIdentifier(stage.AixmdatatypesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Monitored = %s", navaidequipmentmonitoringtypeIdent, targetIdent))
			}
			for _, elem := range navaidequipmentmonitoringtype.Extension {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := elem.GongGetIdentifier(stage.AixmfeaturesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Extension = append(%s.Extension, %s)", navaidequipmentmonitoringtypeIdent, navaidequipmentmonitoringtypeIdent, targetIdent))
			}
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
		for _, navaidoperationalstatuspropertytype := range __gong__sortStageSetInstances(stage.AixmfeaturesStage.NavaidOperationalStatusPropertyTypes, stage.AixmfeaturesStage.NavaidOperationalStatusPropertyType_stagedOrder) {
			if lastStageDecl != "AixmfeaturesStage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "AixmfeaturesStage"
			}
			navaidoperationalstatuspropertytypeIdent := navaidoperationalstatuspropertytype.GongGetIdentifier(stage.AixmfeaturesStage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&aixmfeatures.NavaidOperationalStatusPropertyType{Name: %s}).Stage(stage.AixmfeaturesStage)", navaidoperationalstatuspropertytypeIdent, __gong__toRawStringLiteral(navaidoperationalstatuspropertytype.Name)))
			if lastStageVal != "AixmfeaturesStage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "AixmfeaturesStage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", navaidoperationalstatuspropertytypeIdent, __gong__toRawStringLiteral(navaidoperationalstatuspropertytype.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.NilReason = %s", navaidoperationalstatuspropertytypeIdent, __gong__toRawStringLiteral(navaidoperationalstatuspropertytype.NilReason)))
			if navaidoperationalstatuspropertytype.NavaidOperationalStatus != nil {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := navaidoperationalstatuspropertytype.NavaidOperationalStatus.GongGetIdentifier(stage.AixmfeaturesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.NavaidOperationalStatus = %s", navaidoperationalstatuspropertytypeIdent, targetIdent))
			}
		}
	}
	if stage.AixmfeaturesStage != nil {
		for _, navaidoperationalstatustype := range __gong__sortStageSetInstances(stage.AixmfeaturesStage.NavaidOperationalStatusTypes, stage.AixmfeaturesStage.NavaidOperationalStatusType_stagedOrder) {
			if lastStageDecl != "AixmfeaturesStage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "AixmfeaturesStage"
			}
			navaidoperationalstatustypeIdent := navaidoperationalstatustype.GongGetIdentifier(stage.AixmfeaturesStage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&aixmfeatures.NavaidOperationalStatusType{Name: %s}).Stage(stage.AixmfeaturesStage)", navaidoperationalstatustypeIdent, __gong__toRawStringLiteral(navaidoperationalstatustype.Name)))
			if lastStageVal != "AixmfeaturesStage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "AixmfeaturesStage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", navaidoperationalstatustypeIdent, __gong__toRawStringLiteral(navaidoperationalstatustype.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Id = %s", navaidoperationalstatustypeIdent, __gong__toRawStringLiteral(navaidoperationalstatustype.Id)))
			values.WriteString(fmt.Sprintf("\n\t%s.Description = %s", navaidoperationalstatustypeIdent, __gong__toRawStringLiteral(navaidoperationalstatustype.Description)))
			if navaidoperationalstatustype.Identifier != nil {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := navaidoperationalstatustype.Identifier.GongGetIdentifier(stage.GmlStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Identifier = %s", navaidoperationalstatustypeIdent, targetIdent))
			}
			for _, elem := range navaidoperationalstatustype.NameXSD {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := elem.GongGetIdentifier(stage.GmlStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.NameXSD = append(%s.NameXSD, %s)", navaidoperationalstatustypeIdent, navaidoperationalstatustypeIdent, targetIdent))
			}
			for _, elem := range navaidoperationalstatustype.TimeInterval {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := elem.GongGetIdentifier(stage.AixmfeaturesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.TimeInterval = append(%s.TimeInterval, %s)", navaidoperationalstatustypeIdent, navaidoperationalstatustypeIdent, targetIdent))
			}
			for _, elem := range navaidoperationalstatustype.Annotation {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := elem.GongGetIdentifier(stage.AixmfeaturesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Annotation = append(%s.Annotation, %s)", navaidoperationalstatustypeIdent, navaidoperationalstatustypeIdent, targetIdent))
			}
			for _, elem := range navaidoperationalstatustype.SpecialDateAuthority {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := elem.GongGetIdentifier(stage.AixmfeaturesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.SpecialDateAuthority = append(%s.SpecialDateAuthority, %s)", navaidoperationalstatustypeIdent, navaidoperationalstatustypeIdent, targetIdent))
			}
			if navaidoperationalstatustype.OperationalStatus != nil {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := navaidoperationalstatustype.OperationalStatus.GongGetIdentifier(stage.AixmdatatypesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.OperationalStatus = %s", navaidoperationalstatustypeIdent, targetIdent))
			}
			if navaidoperationalstatustype.SignalType != nil {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := navaidoperationalstatustype.SignalType.GongGetIdentifier(stage.AixmdatatypesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.SignalType = %s", navaidoperationalstatustypeIdent, targetIdent))
			}
			for _, elem := range navaidoperationalstatustype.Extension {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := elem.GongGetIdentifier(stage.AixmfeaturesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Extension = append(%s.Extension, %s)", navaidoperationalstatustypeIdent, navaidoperationalstatustypeIdent, targetIdent))
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
		for _, navaidtimeslicepropertytype := range __gong__sortStageSetInstances(stage.AixmfeaturesStage.NavaidTimeSlicePropertyTypes, stage.AixmfeaturesStage.NavaidTimeSlicePropertyType_stagedOrder) {
			if lastStageDecl != "AixmfeaturesStage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "AixmfeaturesStage"
			}
			navaidtimeslicepropertytypeIdent := navaidtimeslicepropertytype.GongGetIdentifier(stage.AixmfeaturesStage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&aixmfeatures.NavaidTimeSlicePropertyType{Name: %s}).Stage(stage.AixmfeaturesStage)", navaidtimeslicepropertytypeIdent, __gong__toRawStringLiteral(navaidtimeslicepropertytype.Name)))
			if lastStageVal != "AixmfeaturesStage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "AixmfeaturesStage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", navaidtimeslicepropertytypeIdent, __gong__toRawStringLiteral(navaidtimeslicepropertytype.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Owns = %t", navaidtimeslicepropertytypeIdent, navaidtimeslicepropertytype.Owns))
			if navaidtimeslicepropertytype.NavaidTimeSlice != nil {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := navaidtimeslicepropertytype.NavaidTimeSlice.GongGetIdentifier(stage.AixmfeaturesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.NavaidTimeSlice = %s", navaidtimeslicepropertytypeIdent, targetIdent))
			}
		}
	}
	if stage.AixmfeaturesStage != nil {
		for _, navaidtimeslicetype := range __gong__sortStageSetInstances(stage.AixmfeaturesStage.NavaidTimeSliceTypes, stage.AixmfeaturesStage.NavaidTimeSliceType_stagedOrder) {
			if lastStageDecl != "AixmfeaturesStage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "AixmfeaturesStage"
			}
			navaidtimeslicetypeIdent := navaidtimeslicetype.GongGetIdentifier(stage.AixmfeaturesStage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&aixmfeatures.NavaidTimeSliceType{Name: %s}).Stage(stage.AixmfeaturesStage)", navaidtimeslicetypeIdent, __gong__toRawStringLiteral(navaidtimeslicetype.Name)))
			if lastStageVal != "AixmfeaturesStage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "AixmfeaturesStage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", navaidtimeslicetypeIdent, __gong__toRawStringLiteral(navaidtimeslicetype.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Id = %s", navaidtimeslicetypeIdent, __gong__toRawStringLiteral(navaidtimeslicetype.Id)))
			values.WriteString(fmt.Sprintf("\n\t%s.Description = %s", navaidtimeslicetypeIdent, __gong__toRawStringLiteral(navaidtimeslicetype.Description)))
			values.WriteString(fmt.Sprintf("\n\t%s.Interpretation = %s", navaidtimeslicetypeIdent, __gong__toRawStringLiteral(navaidtimeslicetype.Interpretation)))
			values.WriteString(fmt.Sprintf("\n\t%s.SequenceNumber = %d", navaidtimeslicetypeIdent, navaidtimeslicetype.SequenceNumber))
			values.WriteString(fmt.Sprintf("\n\t%s.CorrectionNumber = %d", navaidtimeslicetypeIdent, navaidtimeslicetype.CorrectionNumber))
			if navaidtimeslicetype.Identifier != nil {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := navaidtimeslicetype.Identifier.GongGetIdentifier(stage.GmlStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Identifier = %s", navaidtimeslicetypeIdent, targetIdent))
			}
			for _, elem := range navaidtimeslicetype.NameXSD {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := elem.GongGetIdentifier(stage.GmlStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.NameXSD = append(%s.NameXSD, %s)", navaidtimeslicetypeIdent, navaidtimeslicetypeIdent, targetIdent))
			}
			if navaidtimeslicetype.ValidTime != nil {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := navaidtimeslicetype.ValidTime.GongGetIdentifier(stage.GmlStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.ValidTime = %s", navaidtimeslicetypeIdent, targetIdent))
			}
			if navaidtimeslicetype.TimeSliceMetadata != nil {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := navaidtimeslicetype.TimeSliceMetadata.GongGetIdentifier(stage.AixmabstractgmlobjecttypesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.TimeSliceMetadata = %s", navaidtimeslicetypeIdent, targetIdent))
			}
			if navaidtimeslicetype.FeatureLifetime != nil {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := navaidtimeslicetype.FeatureLifetime.GongGetIdentifier(stage.GmlStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.FeatureLifetime = %s", navaidtimeslicetypeIdent, targetIdent))
			}
			if navaidtimeslicetype.Type_XSD != nil {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := navaidtimeslicetype.Type_XSD.GongGetIdentifier(stage.AixmdatatypesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Type_XSD = %s", navaidtimeslicetypeIdent, targetIdent))
			}
			if navaidtimeslicetype.Designator != nil {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := navaidtimeslicetype.Designator.GongGetIdentifier(stage.AixmdatatypesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Designator = %s", navaidtimeslicetypeIdent, targetIdent))
			}
			if navaidtimeslicetype.FlightChecked != nil {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := navaidtimeslicetype.FlightChecked.GongGetIdentifier(stage.AixmdatatypesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.FlightChecked = %s", navaidtimeslicetypeIdent, targetIdent))
			}
			if navaidtimeslicetype.Purpose != nil {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := navaidtimeslicetype.Purpose.GongGetIdentifier(stage.AixmdatatypesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Purpose = %s", navaidtimeslicetypeIdent, targetIdent))
			}
			if navaidtimeslicetype.SignalPerformance != nil {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := navaidtimeslicetype.SignalPerformance.GongGetIdentifier(stage.AixmdatatypesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.SignalPerformance = %s", navaidtimeslicetypeIdent, targetIdent))
			}
			if navaidtimeslicetype.CourseQuality != nil {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := navaidtimeslicetype.CourseQuality.GongGetIdentifier(stage.AixmdatatypesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.CourseQuality = %s", navaidtimeslicetypeIdent, targetIdent))
			}
			if navaidtimeslicetype.IntegrityLevel != nil {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := navaidtimeslicetype.IntegrityLevel.GongGetIdentifier(stage.AixmdatatypesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.IntegrityLevel = %s", navaidtimeslicetypeIdent, targetIdent))
			}
			for _, elem := range navaidtimeslicetype.TouchDownLiftOff {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := elem.GongGetIdentifier(stage.AixmfeaturesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.TouchDownLiftOff = append(%s.TouchDownLiftOff, %s)", navaidtimeslicetypeIdent, navaidtimeslicetypeIdent, targetIdent))
			}
			for _, elem := range navaidtimeslicetype.NavaidEquipment {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := elem.GongGetIdentifier(stage.AixmfeaturesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.NavaidEquipment = append(%s.NavaidEquipment, %s)", navaidtimeslicetypeIdent, navaidtimeslicetypeIdent, targetIdent))
			}
			if navaidtimeslicetype.Location != nil {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := navaidtimeslicetype.Location.GongGetIdentifier(stage.AixmfeaturesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Location = %s", navaidtimeslicetypeIdent, targetIdent))
			}
			for _, elem := range navaidtimeslicetype.RunwayDirection {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := elem.GongGetIdentifier(stage.AixmfeaturesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.RunwayDirection = append(%s.RunwayDirection, %s)", navaidtimeslicetypeIdent, navaidtimeslicetypeIdent, targetIdent))
			}
			for _, elem := range navaidtimeslicetype.ServedAirport {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := elem.GongGetIdentifier(stage.AixmfeaturesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.ServedAirport = append(%s.ServedAirport, %s)", navaidtimeslicetypeIdent, navaidtimeslicetypeIdent, targetIdent))
			}
			for _, elem := range navaidtimeslicetype.Availability {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := elem.GongGetIdentifier(stage.AixmfeaturesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Availability = append(%s.Availability, %s)", navaidtimeslicetypeIdent, navaidtimeslicetypeIdent, targetIdent))
			}
			for _, elem := range navaidtimeslicetype.Annotation {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := elem.GongGetIdentifier(stage.AixmfeaturesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Annotation = append(%s.Annotation, %s)", navaidtimeslicetypeIdent, navaidtimeslicetypeIdent, targetIdent))
			}
			for _, elem := range navaidtimeslicetype.Extension {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := elem.GongGetIdentifier(stage.AixmfeaturesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Extension = append(%s.Extension, %s)", navaidtimeslicetypeIdent, navaidtimeslicetypeIdent, targetIdent))
			}
		}
	}
	if stage.AixmfeaturesStage != nil {
		for _, navaidtype := range __gong__sortStageSetInstances(stage.AixmfeaturesStage.NavaidTypes, stage.AixmfeaturesStage.NavaidType_stagedOrder) {
			if lastStageDecl != "AixmfeaturesStage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "AixmfeaturesStage"
			}
			navaidtypeIdent := navaidtype.GongGetIdentifier(stage.AixmfeaturesStage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&aixmfeatures.NavaidType{Name: %s}).Stage(stage.AixmfeaturesStage)", navaidtypeIdent, __gong__toRawStringLiteral(navaidtype.Name)))
			if lastStageVal != "AixmfeaturesStage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "AixmfeaturesStage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", navaidtypeIdent, __gong__toRawStringLiteral(navaidtype.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Id = %s", navaidtypeIdent, __gong__toRawStringLiteral(navaidtype.Id)))
			values.WriteString(fmt.Sprintf("\n\t%s.Description = %s", navaidtypeIdent, __gong__toRawStringLiteral(navaidtype.Description)))
			if navaidtype.Identifier != nil {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := navaidtype.Identifier.GongGetIdentifier(stage.GmlStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Identifier = %s", navaidtypeIdent, targetIdent))
			}
			for _, elem := range navaidtype.NameXSD {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := elem.GongGetIdentifier(stage.GmlStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.NameXSD = append(%s.NameXSD, %s)", navaidtypeIdent, navaidtypeIdent, targetIdent))
			}
			if navaidtype.BoundedBy != nil {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := navaidtype.BoundedBy.GongGetIdentifier(stage.GmlStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.BoundedBy = %s", navaidtypeIdent, targetIdent))
			}
			if navaidtype.FeatureMetadata != nil {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := navaidtype.FeatureMetadata.GongGetIdentifier(stage.AixmabstractgmlobjecttypesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.FeatureMetadata = %s", navaidtypeIdent, targetIdent))
			}
			for _, elem := range navaidtype.TimeSlice {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := elem.GongGetIdentifier(stage.AixmfeaturesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.TimeSlice = append(%s.TimeSlice, %s)", navaidtypeIdent, navaidtypeIdent, targetIdent))
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
		for _, precisionapproachradartimeslicepropertytype := range __gong__sortStageSetInstances(stage.AixmfeaturesStage.PrecisionApproachRadarTimeSlicePropertyTypes, stage.AixmfeaturesStage.PrecisionApproachRadarTimeSlicePropertyType_stagedOrder) {
			if lastStageDecl != "AixmfeaturesStage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "AixmfeaturesStage"
			}
			precisionapproachradartimeslicepropertytypeIdent := precisionapproachradartimeslicepropertytype.GongGetIdentifier(stage.AixmfeaturesStage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&aixmfeatures.PrecisionApproachRadarTimeSlicePropertyType{Name: %s}).Stage(stage.AixmfeaturesStage)", precisionapproachradartimeslicepropertytypeIdent, __gong__toRawStringLiteral(precisionapproachradartimeslicepropertytype.Name)))
			if lastStageVal != "AixmfeaturesStage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "AixmfeaturesStage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", precisionapproachradartimeslicepropertytypeIdent, __gong__toRawStringLiteral(precisionapproachradartimeslicepropertytype.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Owns = %t", precisionapproachradartimeslicepropertytypeIdent, precisionapproachradartimeslicepropertytype.Owns))
			if precisionapproachradartimeslicepropertytype.PrecisionApproachRadarTimeSlice != nil {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := precisionapproachradartimeslicepropertytype.PrecisionApproachRadarTimeSlice.GongGetIdentifier(stage.AixmfeaturesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.PrecisionApproachRadarTimeSlice = %s", precisionapproachradartimeslicepropertytypeIdent, targetIdent))
			}
		}
	}
	if stage.AixmfeaturesStage != nil {
		for _, precisionapproachradartimeslicetype := range __gong__sortStageSetInstances(stage.AixmfeaturesStage.PrecisionApproachRadarTimeSliceTypes, stage.AixmfeaturesStage.PrecisionApproachRadarTimeSliceType_stagedOrder) {
			if lastStageDecl != "AixmfeaturesStage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "AixmfeaturesStage"
			}
			precisionapproachradartimeslicetypeIdent := precisionapproachradartimeslicetype.GongGetIdentifier(stage.AixmfeaturesStage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&aixmfeatures.PrecisionApproachRadarTimeSliceType{Name: %s}).Stage(stage.AixmfeaturesStage)", precisionapproachradartimeslicetypeIdent, __gong__toRawStringLiteral(precisionapproachradartimeslicetype.Name)))
			if lastStageVal != "AixmfeaturesStage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "AixmfeaturesStage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", precisionapproachradartimeslicetypeIdent, __gong__toRawStringLiteral(precisionapproachradartimeslicetype.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Id = %s", precisionapproachradartimeslicetypeIdent, __gong__toRawStringLiteral(precisionapproachradartimeslicetype.Id)))
			values.WriteString(fmt.Sprintf("\n\t%s.Description = %s", precisionapproachradartimeslicetypeIdent, __gong__toRawStringLiteral(precisionapproachradartimeslicetype.Description)))
			values.WriteString(fmt.Sprintf("\n\t%s.Interpretation = %s", precisionapproachradartimeslicetypeIdent, __gong__toRawStringLiteral(precisionapproachradartimeslicetype.Interpretation)))
			values.WriteString(fmt.Sprintf("\n\t%s.SequenceNumber = %d", precisionapproachradartimeslicetypeIdent, precisionapproachradartimeslicetype.SequenceNumber))
			values.WriteString(fmt.Sprintf("\n\t%s.CorrectionNumber = %d", precisionapproachradartimeslicetypeIdent, precisionapproachradartimeslicetype.CorrectionNumber))
			if precisionapproachradartimeslicetype.Identifier != nil {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := precisionapproachradartimeslicetype.Identifier.GongGetIdentifier(stage.GmlStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Identifier = %s", precisionapproachradartimeslicetypeIdent, targetIdent))
			}
			for _, elem := range precisionapproachradartimeslicetype.NameXSD {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := elem.GongGetIdentifier(stage.GmlStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.NameXSD = append(%s.NameXSD, %s)", precisionapproachradartimeslicetypeIdent, precisionapproachradartimeslicetypeIdent, targetIdent))
			}
			if precisionapproachradartimeslicetype.ValidTime != nil {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := precisionapproachradartimeslicetype.ValidTime.GongGetIdentifier(stage.GmlStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.ValidTime = %s", precisionapproachradartimeslicetypeIdent, targetIdent))
			}
			if precisionapproachradartimeslicetype.TimeSliceMetadata != nil {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := precisionapproachradartimeslicetype.TimeSliceMetadata.GongGetIdentifier(stage.AixmabstractgmlobjecttypesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.TimeSliceMetadata = %s", precisionapproachradartimeslicetypeIdent, targetIdent))
			}
			if precisionapproachradartimeslicetype.FeatureLifetime != nil {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := precisionapproachradartimeslicetype.FeatureLifetime.GongGetIdentifier(stage.GmlStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.FeatureLifetime = %s", precisionapproachradartimeslicetypeIdent, targetIdent))
			}
			if precisionapproachradartimeslicetype.SerialNumber != nil {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := precisionapproachradartimeslicetype.SerialNumber.GongGetIdentifier(stage.AixmdatatypesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.SerialNumber = %s", precisionapproachradartimeslicetypeIdent, targetIdent))
			}
			if precisionapproachradartimeslicetype.Range_XSD != nil {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := precisionapproachradartimeslicetype.Range_XSD.GongGetIdentifier(stage.AixmdatatypesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Range_XSD = %s", precisionapproachradartimeslicetypeIdent, targetIdent))
			}
			if precisionapproachradartimeslicetype.RangeAccuracy != nil {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := precisionapproachradartimeslicetype.RangeAccuracy.GongGetIdentifier(stage.AixmdatatypesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.RangeAccuracy = %s", precisionapproachradartimeslicetypeIdent, targetIdent))
			}
			if precisionapproachradartimeslicetype.DualChannel != nil {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := precisionapproachradartimeslicetype.DualChannel.GongGetIdentifier(stage.AixmdatatypesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.DualChannel = %s", precisionapproachradartimeslicetypeIdent, targetIdent))
			}
			if precisionapproachradartimeslicetype.MovingTargetIndicator != nil {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := precisionapproachradartimeslicetype.MovingTargetIndicator.GongGetIdentifier(stage.AixmdatatypesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.MovingTargetIndicator = %s", precisionapproachradartimeslicetypeIdent, targetIdent))
			}
			if precisionapproachradartimeslicetype.StandbyPower != nil {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := precisionapproachradartimeslicetype.StandbyPower.GongGetIdentifier(stage.AixmdatatypesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.StandbyPower = %s", precisionapproachradartimeslicetypeIdent, targetIdent))
			}
			if precisionapproachradartimeslicetype.Digital != nil {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := precisionapproachradartimeslicetype.Digital.GongGetIdentifier(stage.AixmdatatypesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Digital = %s", precisionapproachradartimeslicetypeIdent, targetIdent))
			}
			if precisionapproachradartimeslicetype.MilitaryUseOnly != nil {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := precisionapproachradartimeslicetype.MilitaryUseOnly.GongGetIdentifier(stage.AixmdatatypesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.MilitaryUseOnly = %s", precisionapproachradartimeslicetypeIdent, targetIdent))
			}
			if precisionapproachradartimeslicetype.SpecialUseOnly != nil {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := precisionapproachradartimeslicetype.SpecialUseOnly.GongGetIdentifier(stage.AixmdatatypesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.SpecialUseOnly = %s", precisionapproachradartimeslicetypeIdent, targetIdent))
			}
			if precisionapproachradartimeslicetype.SpecialAircraftOnly != nil {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := precisionapproachradartimeslicetype.SpecialAircraftOnly.GongGetIdentifier(stage.AixmdatatypesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.SpecialAircraftOnly = %s", precisionapproachradartimeslicetypeIdent, targetIdent))
			}
			if precisionapproachradartimeslicetype.MagneticVariation != nil {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := precisionapproachradartimeslicetype.MagneticVariation.GongGetIdentifier(stage.AixmdatatypesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.MagneticVariation = %s", precisionapproachradartimeslicetypeIdent, targetIdent))
			}
			if precisionapproachradartimeslicetype.MagneticVariationAccuracy != nil {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := precisionapproachradartimeslicetype.MagneticVariationAccuracy.GongGetIdentifier(stage.AixmdatatypesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.MagneticVariationAccuracy = %s", precisionapproachradartimeslicetypeIdent, targetIdent))
			}
			if precisionapproachradartimeslicetype.DateMagneticVariation != nil {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := precisionapproachradartimeslicetype.DateMagneticVariation.GongGetIdentifier(stage.AixmdatatypesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.DateMagneticVariation = %s", precisionapproachradartimeslicetypeIdent, targetIdent))
			}
			for _, elem := range precisionapproachradartimeslicetype.Contact {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := elem.GongGetIdentifier(stage.AixmfeaturesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Contact = append(%s.Contact, %s)", precisionapproachradartimeslicetypeIdent, precisionapproachradartimeslicetypeIdent, targetIdent))
			}
			if precisionapproachradartimeslicetype.Location != nil {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := precisionapproachradartimeslicetype.Location.GongGetIdentifier(stage.AixmfeaturesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Location = %s", precisionapproachradartimeslicetypeIdent, targetIdent))
			}
			for _, elem := range precisionapproachradartimeslicetype.Annotation {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := elem.GongGetIdentifier(stage.AixmfeaturesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Annotation = append(%s.Annotation, %s)", precisionapproachradartimeslicetypeIdent, precisionapproachradartimeslicetypeIdent, targetIdent))
			}
			if precisionapproachradartimeslicetype.PrecisionApproachRadarType != nil {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := precisionapproachradartimeslicetype.PrecisionApproachRadarType.GongGetIdentifier(stage.AixmdatatypesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.PrecisionApproachRadarType = %s", precisionapproachradartimeslicetypeIdent, targetIdent))
			}
			if precisionapproachradartimeslicetype.Slope != nil {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := precisionapproachradartimeslicetype.Slope.GongGetIdentifier(stage.AixmdatatypesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Slope = %s", precisionapproachradartimeslicetypeIdent, targetIdent))
			}
			if precisionapproachradartimeslicetype.SlopeAccuracy != nil {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := precisionapproachradartimeslicetype.SlopeAccuracy.GongGetIdentifier(stage.AixmdatatypesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.SlopeAccuracy = %s", precisionapproachradartimeslicetypeIdent, targetIdent))
			}
			for _, elem := range precisionapproachradartimeslicetype.Reflector {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := elem.GongGetIdentifier(stage.AixmfeaturesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Reflector = append(%s.Reflector, %s)", precisionapproachradartimeslicetypeIdent, precisionapproachradartimeslicetypeIdent, targetIdent))
			}
			for _, elem := range precisionapproachradartimeslicetype.Extension {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := elem.GongGetIdentifier(stage.AixmfeaturesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Extension = append(%s.Extension, %s)", precisionapproachradartimeslicetypeIdent, precisionapproachradartimeslicetypeIdent, targetIdent))
			}
		}
	}
	if stage.AixmfeaturesStage != nil {
		for _, precisionapproachradartype := range __gong__sortStageSetInstances(stage.AixmfeaturesStage.PrecisionApproachRadarTypes, stage.AixmfeaturesStage.PrecisionApproachRadarType_stagedOrder) {
			if lastStageDecl != "AixmfeaturesStage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "AixmfeaturesStage"
			}
			precisionapproachradartypeIdent := precisionapproachradartype.GongGetIdentifier(stage.AixmfeaturesStage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&aixmfeatures.PrecisionApproachRadarType{Name: %s}).Stage(stage.AixmfeaturesStage)", precisionapproachradartypeIdent, __gong__toRawStringLiteral(precisionapproachradartype.Name)))
			if lastStageVal != "AixmfeaturesStage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "AixmfeaturesStage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", precisionapproachradartypeIdent, __gong__toRawStringLiteral(precisionapproachradartype.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Id = %s", precisionapproachradartypeIdent, __gong__toRawStringLiteral(precisionapproachradartype.Id)))
			values.WriteString(fmt.Sprintf("\n\t%s.Description = %s", precisionapproachradartypeIdent, __gong__toRawStringLiteral(precisionapproachradartype.Description)))
			if precisionapproachradartype.Identifier != nil {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := precisionapproachradartype.Identifier.GongGetIdentifier(stage.GmlStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Identifier = %s", precisionapproachradartypeIdent, targetIdent))
			}
			for _, elem := range precisionapproachradartype.NameXSD {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := elem.GongGetIdentifier(stage.GmlStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.NameXSD = append(%s.NameXSD, %s)", precisionapproachradartypeIdent, precisionapproachradartypeIdent, targetIdent))
			}
			if precisionapproachradartype.BoundedBy != nil {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := precisionapproachradartype.BoundedBy.GongGetIdentifier(stage.GmlStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.BoundedBy = %s", precisionapproachradartypeIdent, targetIdent))
			}
			if precisionapproachradartype.FeatureMetadata != nil {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := precisionapproachradartype.FeatureMetadata.GongGetIdentifier(stage.AixmabstractgmlobjecttypesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.FeatureMetadata = %s", precisionapproachradartypeIdent, targetIdent))
			}
			for _, elem := range precisionapproachradartype.TimeSlice {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := elem.GongGetIdentifier(stage.AixmfeaturesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.TimeSlice = append(%s.TimeSlice, %s)", precisionapproachradartypeIdent, precisionapproachradartypeIdent, targetIdent))
			}
		}
	}
	if stage.AixmfeaturesStage != nil {
		for _, primarysurveillanceradartimeslicepropertytype := range __gong__sortStageSetInstances(stage.AixmfeaturesStage.PrimarySurveillanceRadarTimeSlicePropertyTypes, stage.AixmfeaturesStage.PrimarySurveillanceRadarTimeSlicePropertyType_stagedOrder) {
			if lastStageDecl != "AixmfeaturesStage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "AixmfeaturesStage"
			}
			primarysurveillanceradartimeslicepropertytypeIdent := primarysurveillanceradartimeslicepropertytype.GongGetIdentifier(stage.AixmfeaturesStage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&aixmfeatures.PrimarySurveillanceRadarTimeSlicePropertyType{Name: %s}).Stage(stage.AixmfeaturesStage)", primarysurveillanceradartimeslicepropertytypeIdent, __gong__toRawStringLiteral(primarysurveillanceradartimeslicepropertytype.Name)))
			if lastStageVal != "AixmfeaturesStage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "AixmfeaturesStage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", primarysurveillanceradartimeslicepropertytypeIdent, __gong__toRawStringLiteral(primarysurveillanceradartimeslicepropertytype.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Owns = %t", primarysurveillanceradartimeslicepropertytypeIdent, primarysurveillanceradartimeslicepropertytype.Owns))
			if primarysurveillanceradartimeslicepropertytype.PrimarySurveillanceRadarTimeSlice != nil {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := primarysurveillanceradartimeslicepropertytype.PrimarySurveillanceRadarTimeSlice.GongGetIdentifier(stage.AixmfeaturesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.PrimarySurveillanceRadarTimeSlice = %s", primarysurveillanceradartimeslicepropertytypeIdent, targetIdent))
			}
		}
	}
	if stage.AixmfeaturesStage != nil {
		for _, primarysurveillanceradartimeslicetype := range __gong__sortStageSetInstances(stage.AixmfeaturesStage.PrimarySurveillanceRadarTimeSliceTypes, stage.AixmfeaturesStage.PrimarySurveillanceRadarTimeSliceType_stagedOrder) {
			if lastStageDecl != "AixmfeaturesStage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "AixmfeaturesStage"
			}
			primarysurveillanceradartimeslicetypeIdent := primarysurveillanceradartimeslicetype.GongGetIdentifier(stage.AixmfeaturesStage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&aixmfeatures.PrimarySurveillanceRadarTimeSliceType{Name: %s}).Stage(stage.AixmfeaturesStage)", primarysurveillanceradartimeslicetypeIdent, __gong__toRawStringLiteral(primarysurveillanceradartimeslicetype.Name)))
			if lastStageVal != "AixmfeaturesStage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "AixmfeaturesStage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", primarysurveillanceradartimeslicetypeIdent, __gong__toRawStringLiteral(primarysurveillanceradartimeslicetype.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Id = %s", primarysurveillanceradartimeslicetypeIdent, __gong__toRawStringLiteral(primarysurveillanceradartimeslicetype.Id)))
			values.WriteString(fmt.Sprintf("\n\t%s.Description = %s", primarysurveillanceradartimeslicetypeIdent, __gong__toRawStringLiteral(primarysurveillanceradartimeslicetype.Description)))
			values.WriteString(fmt.Sprintf("\n\t%s.Interpretation = %s", primarysurveillanceradartimeslicetypeIdent, __gong__toRawStringLiteral(primarysurveillanceradartimeslicetype.Interpretation)))
			values.WriteString(fmt.Sprintf("\n\t%s.SequenceNumber = %d", primarysurveillanceradartimeslicetypeIdent, primarysurveillanceradartimeslicetype.SequenceNumber))
			values.WriteString(fmt.Sprintf("\n\t%s.CorrectionNumber = %d", primarysurveillanceradartimeslicetypeIdent, primarysurveillanceradartimeslicetype.CorrectionNumber))
			if primarysurveillanceradartimeslicetype.Identifier != nil {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := primarysurveillanceradartimeslicetype.Identifier.GongGetIdentifier(stage.GmlStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Identifier = %s", primarysurveillanceradartimeslicetypeIdent, targetIdent))
			}
			for _, elem := range primarysurveillanceradartimeslicetype.NameXSD {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := elem.GongGetIdentifier(stage.GmlStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.NameXSD = append(%s.NameXSD, %s)", primarysurveillanceradartimeslicetypeIdent, primarysurveillanceradartimeslicetypeIdent, targetIdent))
			}
			if primarysurveillanceradartimeslicetype.ValidTime != nil {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := primarysurveillanceradartimeslicetype.ValidTime.GongGetIdentifier(stage.GmlStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.ValidTime = %s", primarysurveillanceradartimeslicetypeIdent, targetIdent))
			}
			if primarysurveillanceradartimeslicetype.TimeSliceMetadata != nil {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := primarysurveillanceradartimeslicetype.TimeSliceMetadata.GongGetIdentifier(stage.AixmabstractgmlobjecttypesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.TimeSliceMetadata = %s", primarysurveillanceradartimeslicetypeIdent, targetIdent))
			}
			if primarysurveillanceradartimeslicetype.FeatureLifetime != nil {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := primarysurveillanceradartimeslicetype.FeatureLifetime.GongGetIdentifier(stage.GmlStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.FeatureLifetime = %s", primarysurveillanceradartimeslicetypeIdent, targetIdent))
			}
			if primarysurveillanceradartimeslicetype.SerialNumber != nil {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := primarysurveillanceradartimeslicetype.SerialNumber.GongGetIdentifier(stage.AixmdatatypesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.SerialNumber = %s", primarysurveillanceradartimeslicetypeIdent, targetIdent))
			}
			if primarysurveillanceradartimeslicetype.Range_XSD != nil {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := primarysurveillanceradartimeslicetype.Range_XSD.GongGetIdentifier(stage.AixmdatatypesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Range_XSD = %s", primarysurveillanceradartimeslicetypeIdent, targetIdent))
			}
			if primarysurveillanceradartimeslicetype.RangeAccuracy != nil {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := primarysurveillanceradartimeslicetype.RangeAccuracy.GongGetIdentifier(stage.AixmdatatypesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.RangeAccuracy = %s", primarysurveillanceradartimeslicetypeIdent, targetIdent))
			}
			if primarysurveillanceradartimeslicetype.DualChannel != nil {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := primarysurveillanceradartimeslicetype.DualChannel.GongGetIdentifier(stage.AixmdatatypesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.DualChannel = %s", primarysurveillanceradartimeslicetypeIdent, targetIdent))
			}
			if primarysurveillanceradartimeslicetype.MovingTargetIndicator != nil {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := primarysurveillanceradartimeslicetype.MovingTargetIndicator.GongGetIdentifier(stage.AixmdatatypesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.MovingTargetIndicator = %s", primarysurveillanceradartimeslicetypeIdent, targetIdent))
			}
			if primarysurveillanceradartimeslicetype.StandbyPower != nil {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := primarysurveillanceradartimeslicetype.StandbyPower.GongGetIdentifier(stage.AixmdatatypesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.StandbyPower = %s", primarysurveillanceradartimeslicetypeIdent, targetIdent))
			}
			if primarysurveillanceradartimeslicetype.Digital != nil {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := primarysurveillanceradartimeslicetype.Digital.GongGetIdentifier(stage.AixmdatatypesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Digital = %s", primarysurveillanceradartimeslicetypeIdent, targetIdent))
			}
			if primarysurveillanceradartimeslicetype.MilitaryUseOnly != nil {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := primarysurveillanceradartimeslicetype.MilitaryUseOnly.GongGetIdentifier(stage.AixmdatatypesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.MilitaryUseOnly = %s", primarysurveillanceradartimeslicetypeIdent, targetIdent))
			}
			if primarysurveillanceradartimeslicetype.SpecialUseOnly != nil {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := primarysurveillanceradartimeslicetype.SpecialUseOnly.GongGetIdentifier(stage.AixmdatatypesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.SpecialUseOnly = %s", primarysurveillanceradartimeslicetypeIdent, targetIdent))
			}
			if primarysurveillanceradartimeslicetype.SpecialAircraftOnly != nil {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := primarysurveillanceradartimeslicetype.SpecialAircraftOnly.GongGetIdentifier(stage.AixmdatatypesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.SpecialAircraftOnly = %s", primarysurveillanceradartimeslicetypeIdent, targetIdent))
			}
			if primarysurveillanceradartimeslicetype.MagneticVariation != nil {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := primarysurveillanceradartimeslicetype.MagneticVariation.GongGetIdentifier(stage.AixmdatatypesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.MagneticVariation = %s", primarysurveillanceradartimeslicetypeIdent, targetIdent))
			}
			if primarysurveillanceradartimeslicetype.MagneticVariationAccuracy != nil {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := primarysurveillanceradartimeslicetype.MagneticVariationAccuracy.GongGetIdentifier(stage.AixmdatatypesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.MagneticVariationAccuracy = %s", primarysurveillanceradartimeslicetypeIdent, targetIdent))
			}
			if primarysurveillanceradartimeslicetype.DateMagneticVariation != nil {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := primarysurveillanceradartimeslicetype.DateMagneticVariation.GongGetIdentifier(stage.AixmdatatypesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.DateMagneticVariation = %s", primarysurveillanceradartimeslicetypeIdent, targetIdent))
			}
			for _, elem := range primarysurveillanceradartimeslicetype.Contact {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := elem.GongGetIdentifier(stage.AixmfeaturesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Contact = append(%s.Contact, %s)", primarysurveillanceradartimeslicetypeIdent, primarysurveillanceradartimeslicetypeIdent, targetIdent))
			}
			if primarysurveillanceradartimeslicetype.Location != nil {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := primarysurveillanceradartimeslicetype.Location.GongGetIdentifier(stage.AixmfeaturesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Location = %s", primarysurveillanceradartimeslicetypeIdent, targetIdent))
			}
			for _, elem := range primarysurveillanceradartimeslicetype.Annotation {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := elem.GongGetIdentifier(stage.AixmfeaturesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Annotation = append(%s.Annotation, %s)", primarysurveillanceradartimeslicetypeIdent, primarysurveillanceradartimeslicetypeIdent, targetIdent))
			}
			if primarysurveillanceradartimeslicetype.VerticalCoverageAltitude != nil {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := primarysurveillanceradartimeslicetype.VerticalCoverageAltitude.GongGetIdentifier(stage.AixmdatatypesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.VerticalCoverageAltitude = %s", primarysurveillanceradartimeslicetypeIdent, targetIdent))
			}
			if primarysurveillanceradartimeslicetype.VerticalCoverageDistance != nil {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := primarysurveillanceradartimeslicetype.VerticalCoverageDistance.GongGetIdentifier(stage.AixmdatatypesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.VerticalCoverageDistance = %s", primarysurveillanceradartimeslicetypeIdent, targetIdent))
			}
			if primarysurveillanceradartimeslicetype.VerticalCoverageAzimuth != nil {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := primarysurveillanceradartimeslicetype.VerticalCoverageAzimuth.GongGetIdentifier(stage.AixmdatatypesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.VerticalCoverageAzimuth = %s", primarysurveillanceradartimeslicetypeIdent, targetIdent))
			}
			if primarysurveillanceradartimeslicetype.AntennaTiltFixed != nil {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := primarysurveillanceradartimeslicetype.AntennaTiltFixed.GongGetIdentifier(stage.AixmdatatypesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.AntennaTiltFixed = %s", primarysurveillanceradartimeslicetypeIdent, targetIdent))
			}
			if primarysurveillanceradartimeslicetype.TiltAngle != nil {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := primarysurveillanceradartimeslicetype.TiltAngle.GongGetIdentifier(stage.AixmdatatypesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.TiltAngle = %s", primarysurveillanceradartimeslicetypeIdent, targetIdent))
			}
			if primarysurveillanceradartimeslicetype.AutomatedRadarTerminalSystem != nil {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := primarysurveillanceradartimeslicetype.AutomatedRadarTerminalSystem.GongGetIdentifier(stage.AixmdatatypesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.AutomatedRadarTerminalSystem = %s", primarysurveillanceradartimeslicetypeIdent, targetIdent))
			}
			for _, elem := range primarysurveillanceradartimeslicetype.GroundStation {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := elem.GongGetIdentifier(stage.AixmfeaturesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.GroundStation = append(%s.GroundStation, %s)", primarysurveillanceradartimeslicetypeIdent, primarysurveillanceradartimeslicetypeIdent, targetIdent))
			}
			if primarysurveillanceradartimeslicetype.Type_XSD != nil {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := primarysurveillanceradartimeslicetype.Type_XSD.GongGetIdentifier(stage.AixmdatatypesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Type_XSD = %s", primarysurveillanceradartimeslicetypeIdent, targetIdent))
			}
			for _, elem := range primarysurveillanceradartimeslicetype.Extension {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := elem.GongGetIdentifier(stage.AixmfeaturesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Extension = append(%s.Extension, %s)", primarysurveillanceradartimeslicetypeIdent, primarysurveillanceradartimeslicetypeIdent, targetIdent))
			}
		}
	}
	if stage.AixmfeaturesStage != nil {
		for _, primarysurveillanceradartype := range __gong__sortStageSetInstances(stage.AixmfeaturesStage.PrimarySurveillanceRadarTypes, stage.AixmfeaturesStage.PrimarySurveillanceRadarType_stagedOrder) {
			if lastStageDecl != "AixmfeaturesStage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "AixmfeaturesStage"
			}
			primarysurveillanceradartypeIdent := primarysurveillanceradartype.GongGetIdentifier(stage.AixmfeaturesStage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&aixmfeatures.PrimarySurveillanceRadarType{Name: %s}).Stage(stage.AixmfeaturesStage)", primarysurveillanceradartypeIdent, __gong__toRawStringLiteral(primarysurveillanceradartype.Name)))
			if lastStageVal != "AixmfeaturesStage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "AixmfeaturesStage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", primarysurveillanceradartypeIdent, __gong__toRawStringLiteral(primarysurveillanceradartype.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Id = %s", primarysurveillanceradartypeIdent, __gong__toRawStringLiteral(primarysurveillanceradartype.Id)))
			values.WriteString(fmt.Sprintf("\n\t%s.Description = %s", primarysurveillanceradartypeIdent, __gong__toRawStringLiteral(primarysurveillanceradartype.Description)))
			if primarysurveillanceradartype.Identifier != nil {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := primarysurveillanceradartype.Identifier.GongGetIdentifier(stage.GmlStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Identifier = %s", primarysurveillanceradartypeIdent, targetIdent))
			}
			for _, elem := range primarysurveillanceradartype.NameXSD {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := elem.GongGetIdentifier(stage.GmlStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.NameXSD = append(%s.NameXSD, %s)", primarysurveillanceradartypeIdent, primarysurveillanceradartypeIdent, targetIdent))
			}
			if primarysurveillanceradartype.BoundedBy != nil {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := primarysurveillanceradartype.BoundedBy.GongGetIdentifier(stage.GmlStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.BoundedBy = %s", primarysurveillanceradartypeIdent, targetIdent))
			}
			if primarysurveillanceradartype.FeatureMetadata != nil {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := primarysurveillanceradartype.FeatureMetadata.GongGetIdentifier(stage.AixmabstractgmlobjecttypesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.FeatureMetadata = %s", primarysurveillanceradartypeIdent, targetIdent))
			}
			for _, elem := range primarysurveillanceradartype.TimeSlice {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := elem.GongGetIdentifier(stage.AixmfeaturesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.TimeSlice = append(%s.TimeSlice, %s)", primarysurveillanceradartypeIdent, primarysurveillanceradartypeIdent, targetIdent))
			}
		}
	}
	if stage.AixmfeaturesStage != nil {
		for _, radarcomponentpropertytype := range __gong__sortStageSetInstances(stage.AixmfeaturesStage.RadarComponentPropertyTypes, stage.AixmfeaturesStage.RadarComponentPropertyType_stagedOrder) {
			if lastStageDecl != "AixmfeaturesStage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "AixmfeaturesStage"
			}
			radarcomponentpropertytypeIdent := radarcomponentpropertytype.GongGetIdentifier(stage.AixmfeaturesStage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&aixmfeatures.RadarComponentPropertyType{Name: %s}).Stage(stage.AixmfeaturesStage)", radarcomponentpropertytypeIdent, __gong__toRawStringLiteral(radarcomponentpropertytype.Name)))
			if lastStageVal != "AixmfeaturesStage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "AixmfeaturesStage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", radarcomponentpropertytypeIdent, __gong__toRawStringLiteral(radarcomponentpropertytype.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.NilReason = %s", radarcomponentpropertytypeIdent, __gong__toRawStringLiteral(radarcomponentpropertytype.NilReason)))
			if radarcomponentpropertytype.RadarComponent != nil {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := radarcomponentpropertytype.RadarComponent.GongGetIdentifier(stage.AixmfeaturesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.RadarComponent = %s", radarcomponentpropertytypeIdent, targetIdent))
			}
		}
	}
	if stage.AixmfeaturesStage != nil {
		for _, radarcomponenttype := range __gong__sortStageSetInstances(stage.AixmfeaturesStage.RadarComponentTypes, stage.AixmfeaturesStage.RadarComponentType_stagedOrder) {
			if lastStageDecl != "AixmfeaturesStage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "AixmfeaturesStage"
			}
			radarcomponenttypeIdent := radarcomponenttype.GongGetIdentifier(stage.AixmfeaturesStage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&aixmfeatures.RadarComponentType{Name: %s}).Stage(stage.AixmfeaturesStage)", radarcomponenttypeIdent, __gong__toRawStringLiteral(radarcomponenttype.Name)))
			if lastStageVal != "AixmfeaturesStage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "AixmfeaturesStage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", radarcomponenttypeIdent, __gong__toRawStringLiteral(radarcomponenttype.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Id = %s", radarcomponenttypeIdent, __gong__toRawStringLiteral(radarcomponenttype.Id)))
			values.WriteString(fmt.Sprintf("\n\t%s.Description = %s", radarcomponenttypeIdent, __gong__toRawStringLiteral(radarcomponenttype.Description)))
			if radarcomponenttype.Identifier != nil {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := radarcomponenttype.Identifier.GongGetIdentifier(stage.GmlStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Identifier = %s", radarcomponenttypeIdent, targetIdent))
			}
			for _, elem := range radarcomponenttype.NameXSD {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := elem.GongGetIdentifier(stage.GmlStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.NameXSD = append(%s.NameXSD, %s)", radarcomponenttypeIdent, radarcomponenttypeIdent, targetIdent))
			}
			if radarcomponenttype.CollocationGroup != nil {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := radarcomponenttype.CollocationGroup.GongGetIdentifier(stage.AixmdatatypesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.CollocationGroup = %s", radarcomponenttypeIdent, targetIdent))
			}
			for _, elem := range radarcomponenttype.Annotation {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := elem.GongGetIdentifier(stage.AixmfeaturesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Annotation = append(%s.Annotation, %s)", radarcomponenttypeIdent, radarcomponenttypeIdent, targetIdent))
			}
			if radarcomponenttype.TheRadarEquipment != nil {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := radarcomponenttype.TheRadarEquipment.GongGetIdentifier(stage.AixmfeaturesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.TheRadarEquipment = %s", radarcomponenttypeIdent, targetIdent))
			}
			for _, elem := range radarcomponenttype.Extension {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := elem.GongGetIdentifier(stage.AixmfeaturesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Extension = append(%s.Extension, %s)", radarcomponenttypeIdent, radarcomponenttypeIdent, targetIdent))
			}
		}
	}
	if stage.AixmfeaturesStage != nil {
		for _, radarequipmentpropertytype := range __gong__sortStageSetInstances(stage.AixmfeaturesStage.RadarEquipmentPropertyTypes, stage.AixmfeaturesStage.RadarEquipmentPropertyType_stagedOrder) {
			if lastStageDecl != "AixmfeaturesStage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "AixmfeaturesStage"
			}
			radarequipmentpropertytypeIdent := radarequipmentpropertytype.GongGetIdentifier(stage.AixmfeaturesStage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&aixmfeatures.RadarEquipmentPropertyType{Name: %s}).Stage(stage.AixmfeaturesStage)", radarequipmentpropertytypeIdent, __gong__toRawStringLiteral(radarequipmentpropertytype.Name)))
			if lastStageVal != "AixmfeaturesStage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "AixmfeaturesStage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", radarequipmentpropertytypeIdent, __gong__toRawStringLiteral(radarequipmentpropertytype.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Owns = %t", radarequipmentpropertytypeIdent, radarequipmentpropertytype.Owns))
			values.WriteString(fmt.Sprintf("\n\t%s.NilReason = %s", radarequipmentpropertytypeIdent, __gong__toRawStringLiteral(radarequipmentpropertytype.NilReason)))
			values.WriteString(fmt.Sprintf("\n\t%s.Type_XSD = %s", radarequipmentpropertytypeIdent, __gong__toRawStringLiteral(radarequipmentpropertytype.Type_XSD)))
			values.WriteString(fmt.Sprintf("\n\t%s.Href = %s", radarequipmentpropertytypeIdent, __gong__toRawStringLiteral(radarequipmentpropertytype.Href)))
			values.WriteString(fmt.Sprintf("\n\t%s.Role = %s", radarequipmentpropertytypeIdent, __gong__toRawStringLiteral(radarequipmentpropertytype.Role)))
			values.WriteString(fmt.Sprintf("\n\t%s.Arcrole = %s", radarequipmentpropertytypeIdent, __gong__toRawStringLiteral(radarequipmentpropertytype.Arcrole)))
			values.WriteString(fmt.Sprintf("\n\t%s.Title = %s", radarequipmentpropertytypeIdent, __gong__toRawStringLiteral(radarequipmentpropertytype.Title)))
			values.WriteString(fmt.Sprintf("\n\t%s.Show = %s", radarequipmentpropertytypeIdent, __gong__toRawStringLiteral(radarequipmentpropertytype.Show)))
			values.WriteString(fmt.Sprintf("\n\t%s.Actuate = %s", radarequipmentpropertytypeIdent, __gong__toRawStringLiteral(radarequipmentpropertytype.Actuate)))
		}
	}
	if stage.AixmfeaturesStage != nil {
		for _, radarsystemtimeslicepropertytype := range __gong__sortStageSetInstances(stage.AixmfeaturesStage.RadarSystemTimeSlicePropertyTypes, stage.AixmfeaturesStage.RadarSystemTimeSlicePropertyType_stagedOrder) {
			if lastStageDecl != "AixmfeaturesStage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "AixmfeaturesStage"
			}
			radarsystemtimeslicepropertytypeIdent := radarsystemtimeslicepropertytype.GongGetIdentifier(stage.AixmfeaturesStage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&aixmfeatures.RadarSystemTimeSlicePropertyType{Name: %s}).Stage(stage.AixmfeaturesStage)", radarsystemtimeslicepropertytypeIdent, __gong__toRawStringLiteral(radarsystemtimeslicepropertytype.Name)))
			if lastStageVal != "AixmfeaturesStage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "AixmfeaturesStage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", radarsystemtimeslicepropertytypeIdent, __gong__toRawStringLiteral(radarsystemtimeslicepropertytype.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Owns = %t", radarsystemtimeslicepropertytypeIdent, radarsystemtimeslicepropertytype.Owns))
			if radarsystemtimeslicepropertytype.RadarSystemTimeSlice != nil {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := radarsystemtimeslicepropertytype.RadarSystemTimeSlice.GongGetIdentifier(stage.AixmfeaturesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.RadarSystemTimeSlice = %s", radarsystemtimeslicepropertytypeIdent, targetIdent))
			}
		}
	}
	if stage.AixmfeaturesStage != nil {
		for _, radarsystemtimeslicetype := range __gong__sortStageSetInstances(stage.AixmfeaturesStage.RadarSystemTimeSliceTypes, stage.AixmfeaturesStage.RadarSystemTimeSliceType_stagedOrder) {
			if lastStageDecl != "AixmfeaturesStage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "AixmfeaturesStage"
			}
			radarsystemtimeslicetypeIdent := radarsystemtimeslicetype.GongGetIdentifier(stage.AixmfeaturesStage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&aixmfeatures.RadarSystemTimeSliceType{Name: %s}).Stage(stage.AixmfeaturesStage)", radarsystemtimeslicetypeIdent, __gong__toRawStringLiteral(radarsystemtimeslicetype.Name)))
			if lastStageVal != "AixmfeaturesStage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "AixmfeaturesStage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", radarsystemtimeslicetypeIdent, __gong__toRawStringLiteral(radarsystemtimeslicetype.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Id = %s", radarsystemtimeslicetypeIdent, __gong__toRawStringLiteral(radarsystemtimeslicetype.Id)))
			values.WriteString(fmt.Sprintf("\n\t%s.Description = %s", radarsystemtimeslicetypeIdent, __gong__toRawStringLiteral(radarsystemtimeslicetype.Description)))
			values.WriteString(fmt.Sprintf("\n\t%s.Interpretation = %s", radarsystemtimeslicetypeIdent, __gong__toRawStringLiteral(radarsystemtimeslicetype.Interpretation)))
			values.WriteString(fmt.Sprintf("\n\t%s.SequenceNumber = %d", radarsystemtimeslicetypeIdent, radarsystemtimeslicetype.SequenceNumber))
			values.WriteString(fmt.Sprintf("\n\t%s.CorrectionNumber = %d", radarsystemtimeslicetypeIdent, radarsystemtimeslicetype.CorrectionNumber))
			if radarsystemtimeslicetype.Identifier != nil {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := radarsystemtimeslicetype.Identifier.GongGetIdentifier(stage.GmlStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Identifier = %s", radarsystemtimeslicetypeIdent, targetIdent))
			}
			for _, elem := range radarsystemtimeslicetype.NameXSD {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := elem.GongGetIdentifier(stage.GmlStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.NameXSD = append(%s.NameXSD, %s)", radarsystemtimeslicetypeIdent, radarsystemtimeslicetypeIdent, targetIdent))
			}
			if radarsystemtimeslicetype.ValidTime != nil {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := radarsystemtimeslicetype.ValidTime.GongGetIdentifier(stage.GmlStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.ValidTime = %s", radarsystemtimeslicetypeIdent, targetIdent))
			}
			if radarsystemtimeslicetype.TimeSliceMetadata != nil {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := radarsystemtimeslicetype.TimeSliceMetadata.GongGetIdentifier(stage.AixmabstractgmlobjecttypesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.TimeSliceMetadata = %s", radarsystemtimeslicetypeIdent, targetIdent))
			}
			if radarsystemtimeslicetype.FeatureLifetime != nil {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := radarsystemtimeslicetype.FeatureLifetime.GongGetIdentifier(stage.GmlStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.FeatureLifetime = %s", radarsystemtimeslicetypeIdent, targetIdent))
			}
			if radarsystemtimeslicetype.Type_XSD != nil {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := radarsystemtimeslicetype.Type_XSD.GongGetIdentifier(stage.AixmdatatypesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Type_XSD = %s", radarsystemtimeslicetypeIdent, targetIdent))
			}
			if radarsystemtimeslicetype.Model != nil {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := radarsystemtimeslicetype.Model.GongGetIdentifier(stage.AixmdatatypesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Model = %s", radarsystemtimeslicetypeIdent, targetIdent))
			}
			if radarsystemtimeslicetype.GeneralTerrainMonitor != nil {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := radarsystemtimeslicetype.GeneralTerrainMonitor.GongGetIdentifier(stage.AixmdatatypesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.GeneralTerrainMonitor = %s", radarsystemtimeslicetypeIdent, targetIdent))
			}
			if radarsystemtimeslicetype.BroadcastIdentifier != nil {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := radarsystemtimeslicetype.BroadcastIdentifier.GongGetIdentifier(stage.AixmdatatypesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.BroadcastIdentifier = %s", radarsystemtimeslicetypeIdent, targetIdent))
			}
			for _, elem := range radarsystemtimeslicetype.RadarEquipment {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := elem.GongGetIdentifier(stage.AixmfeaturesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.RadarEquipment = append(%s.RadarEquipment, %s)", radarsystemtimeslicetypeIdent, radarsystemtimeslicetypeIdent, targetIdent))
			}
			for _, elem := range radarsystemtimeslicetype.Office {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := elem.GongGetIdentifier(stage.AixmfeaturesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Office = append(%s.Office, %s)", radarsystemtimeslicetypeIdent, radarsystemtimeslicetypeIdent, targetIdent))
			}
			if radarsystemtimeslicetype.AirportHeliport != nil {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := radarsystemtimeslicetype.AirportHeliport.GongGetIdentifier(stage.AixmfeaturesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.AirportHeliport = %s", radarsystemtimeslicetypeIdent, targetIdent))
			}
			for _, elem := range radarsystemtimeslicetype.PARRunway {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := elem.GongGetIdentifier(stage.AixmfeaturesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.PARRunway = append(%s.PARRunway, %s)", radarsystemtimeslicetypeIdent, radarsystemtimeslicetypeIdent, targetIdent))
			}
			if radarsystemtimeslicetype.Location != nil {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := radarsystemtimeslicetype.Location.GongGetIdentifier(stage.AixmfeaturesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Location = %s", radarsystemtimeslicetypeIdent, targetIdent))
			}
			for _, elem := range radarsystemtimeslicetype.Annotation {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := elem.GongGetIdentifier(stage.AixmfeaturesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Annotation = append(%s.Annotation, %s)", radarsystemtimeslicetypeIdent, radarsystemtimeslicetypeIdent, targetIdent))
			}
			for _, elem := range radarsystemtimeslicetype.Extension {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := elem.GongGetIdentifier(stage.AixmfeaturesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Extension = append(%s.Extension, %s)", radarsystemtimeslicetypeIdent, radarsystemtimeslicetypeIdent, targetIdent))
			}
		}
	}
	if stage.AixmfeaturesStage != nil {
		for _, radarsystemtype := range __gong__sortStageSetInstances(stage.AixmfeaturesStage.RadarSystemTypes, stage.AixmfeaturesStage.RadarSystemType_stagedOrder) {
			if lastStageDecl != "AixmfeaturesStage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "AixmfeaturesStage"
			}
			radarsystemtypeIdent := radarsystemtype.GongGetIdentifier(stage.AixmfeaturesStage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&aixmfeatures.RadarSystemType{Name: %s}).Stage(stage.AixmfeaturesStage)", radarsystemtypeIdent, __gong__toRawStringLiteral(radarsystemtype.Name)))
			if lastStageVal != "AixmfeaturesStage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "AixmfeaturesStage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", radarsystemtypeIdent, __gong__toRawStringLiteral(radarsystemtype.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Id = %s", radarsystemtypeIdent, __gong__toRawStringLiteral(radarsystemtype.Id)))
			values.WriteString(fmt.Sprintf("\n\t%s.Description = %s", radarsystemtypeIdent, __gong__toRawStringLiteral(radarsystemtype.Description)))
			if radarsystemtype.Identifier != nil {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := radarsystemtype.Identifier.GongGetIdentifier(stage.GmlStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Identifier = %s", radarsystemtypeIdent, targetIdent))
			}
			for _, elem := range radarsystemtype.NameXSD {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := elem.GongGetIdentifier(stage.GmlStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.NameXSD = append(%s.NameXSD, %s)", radarsystemtypeIdent, radarsystemtypeIdent, targetIdent))
			}
			if radarsystemtype.BoundedBy != nil {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := radarsystemtype.BoundedBy.GongGetIdentifier(stage.GmlStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.BoundedBy = %s", radarsystemtypeIdent, targetIdent))
			}
			if radarsystemtype.FeatureMetadata != nil {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := radarsystemtype.FeatureMetadata.GongGetIdentifier(stage.AixmabstractgmlobjecttypesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.FeatureMetadata = %s", radarsystemtypeIdent, targetIdent))
			}
			for _, elem := range radarsystemtype.TimeSlice {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := elem.GongGetIdentifier(stage.AixmfeaturesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.TimeSlice = append(%s.TimeSlice, %s)", radarsystemtypeIdent, radarsystemtypeIdent, targetIdent))
			}
		}
	}
	if stage.AixmfeaturesStage != nil {
		for _, reflectorpropertytype := range __gong__sortStageSetInstances(stage.AixmfeaturesStage.ReflectorPropertyTypes, stage.AixmfeaturesStage.ReflectorPropertyType_stagedOrder) {
			if lastStageDecl != "AixmfeaturesStage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "AixmfeaturesStage"
			}
			reflectorpropertytypeIdent := reflectorpropertytype.GongGetIdentifier(stage.AixmfeaturesStage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&aixmfeatures.ReflectorPropertyType{Name: %s}).Stage(stage.AixmfeaturesStage)", reflectorpropertytypeIdent, __gong__toRawStringLiteral(reflectorpropertytype.Name)))
			if lastStageVal != "AixmfeaturesStage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "AixmfeaturesStage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", reflectorpropertytypeIdent, __gong__toRawStringLiteral(reflectorpropertytype.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.NilReason = %s", reflectorpropertytypeIdent, __gong__toRawStringLiteral(reflectorpropertytype.NilReason)))
			if reflectorpropertytype.Reflector != nil {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := reflectorpropertytype.Reflector.GongGetIdentifier(stage.AixmfeaturesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Reflector = %s", reflectorpropertytypeIdent, targetIdent))
			}
		}
	}
	if stage.AixmfeaturesStage != nil {
		for _, reflectortype := range __gong__sortStageSetInstances(stage.AixmfeaturesStage.ReflectorTypes, stage.AixmfeaturesStage.ReflectorType_stagedOrder) {
			if lastStageDecl != "AixmfeaturesStage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "AixmfeaturesStage"
			}
			reflectortypeIdent := reflectortype.GongGetIdentifier(stage.AixmfeaturesStage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&aixmfeatures.ReflectorType{Name: %s}).Stage(stage.AixmfeaturesStage)", reflectortypeIdent, __gong__toRawStringLiteral(reflectortype.Name)))
			if lastStageVal != "AixmfeaturesStage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "AixmfeaturesStage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", reflectortypeIdent, __gong__toRawStringLiteral(reflectortype.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Id = %s", reflectortypeIdent, __gong__toRawStringLiteral(reflectortype.Id)))
			values.WriteString(fmt.Sprintf("\n\t%s.Description = %s", reflectortypeIdent, __gong__toRawStringLiteral(reflectortype.Description)))
			if reflectortype.Identifier != nil {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := reflectortype.Identifier.GongGetIdentifier(stage.GmlStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Identifier = %s", reflectortypeIdent, targetIdent))
			}
			for _, elem := range reflectortype.NameXSD {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := elem.GongGetIdentifier(stage.GmlStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.NameXSD = append(%s.NameXSD, %s)", reflectortypeIdent, reflectortypeIdent, targetIdent))
			}
			if reflectortype.Type_XSD != nil {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := reflectortype.Type_XSD.GongGetIdentifier(stage.AixmdatatypesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Type_XSD = %s", reflectortypeIdent, targetIdent))
			}
			if reflectortype.TouchdownReflector != nil {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := reflectortype.TouchdownReflector.GongGetIdentifier(stage.AixmfeaturesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.TouchdownReflector = %s", reflectortypeIdent, targetIdent))
			}
			for _, elem := range reflectortype.Annotation {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := elem.GongGetIdentifier(stage.AixmfeaturesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Annotation = append(%s.Annotation, %s)", reflectortypeIdent, reflectortypeIdent, targetIdent))
			}
			for _, elem := range reflectortype.Extension {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := elem.GongGetIdentifier(stage.AixmfeaturesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Extension = append(%s.Extension, %s)", reflectortypeIdent, reflectortypeIdent, targetIdent))
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
		for _, runwaypropertytype := range __gong__sortStageSetInstances(stage.AixmfeaturesStage.RunwayPropertyTypes, stage.AixmfeaturesStage.RunwayPropertyType_stagedOrder) {
			if lastStageDecl != "AixmfeaturesStage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "AixmfeaturesStage"
			}
			runwaypropertytypeIdent := runwaypropertytype.GongGetIdentifier(stage.AixmfeaturesStage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&aixmfeatures.RunwayPropertyType{Name: %s}).Stage(stage.AixmfeaturesStage)", runwaypropertytypeIdent, __gong__toRawStringLiteral(runwaypropertytype.Name)))
			if lastStageVal != "AixmfeaturesStage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "AixmfeaturesStage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", runwaypropertytypeIdent, __gong__toRawStringLiteral(runwaypropertytype.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Owns = %t", runwaypropertytypeIdent, runwaypropertytype.Owns))
			values.WriteString(fmt.Sprintf("\n\t%s.NilReason = %s", runwaypropertytypeIdent, __gong__toRawStringLiteral(runwaypropertytype.NilReason)))
			values.WriteString(fmt.Sprintf("\n\t%s.Type_XSD = %s", runwaypropertytypeIdent, __gong__toRawStringLiteral(runwaypropertytype.Type_XSD)))
			values.WriteString(fmt.Sprintf("\n\t%s.Href = %s", runwaypropertytypeIdent, __gong__toRawStringLiteral(runwaypropertytype.Href)))
			values.WriteString(fmt.Sprintf("\n\t%s.Role = %s", runwaypropertytypeIdent, __gong__toRawStringLiteral(runwaypropertytype.Role)))
			values.WriteString(fmt.Sprintf("\n\t%s.Arcrole = %s", runwaypropertytypeIdent, __gong__toRawStringLiteral(runwaypropertytype.Arcrole)))
			values.WriteString(fmt.Sprintf("\n\t%s.Title = %s", runwaypropertytypeIdent, __gong__toRawStringLiteral(runwaypropertytype.Title)))
			values.WriteString(fmt.Sprintf("\n\t%s.Show = %s", runwaypropertytypeIdent, __gong__toRawStringLiteral(runwaypropertytype.Show)))
			values.WriteString(fmt.Sprintf("\n\t%s.Actuate = %s", runwaypropertytypeIdent, __gong__toRawStringLiteral(runwaypropertytype.Actuate)))
		}
	}
	if stage.AixmfeaturesStage != nil {
		for _, sdftimeslicepropertytype := range __gong__sortStageSetInstances(stage.AixmfeaturesStage.SDFTimeSlicePropertyTypes, stage.AixmfeaturesStage.SDFTimeSlicePropertyType_stagedOrder) {
			if lastStageDecl != "AixmfeaturesStage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "AixmfeaturesStage"
			}
			sdftimeslicepropertytypeIdent := sdftimeslicepropertytype.GongGetIdentifier(stage.AixmfeaturesStage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&aixmfeatures.SDFTimeSlicePropertyType{Name: %s}).Stage(stage.AixmfeaturesStage)", sdftimeslicepropertytypeIdent, __gong__toRawStringLiteral(sdftimeslicepropertytype.Name)))
			if lastStageVal != "AixmfeaturesStage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "AixmfeaturesStage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", sdftimeslicepropertytypeIdent, __gong__toRawStringLiteral(sdftimeslicepropertytype.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Owns = %t", sdftimeslicepropertytypeIdent, sdftimeslicepropertytype.Owns))
			if sdftimeslicepropertytype.SDFTimeSlice != nil {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := sdftimeslicepropertytype.SDFTimeSlice.GongGetIdentifier(stage.AixmfeaturesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.SDFTimeSlice = %s", sdftimeslicepropertytypeIdent, targetIdent))
			}
		}
	}
	if stage.AixmfeaturesStage != nil {
		for _, sdftimeslicetype := range __gong__sortStageSetInstances(stage.AixmfeaturesStage.SDFTimeSliceTypes, stage.AixmfeaturesStage.SDFTimeSliceType_stagedOrder) {
			if lastStageDecl != "AixmfeaturesStage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "AixmfeaturesStage"
			}
			sdftimeslicetypeIdent := sdftimeslicetype.GongGetIdentifier(stage.AixmfeaturesStage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&aixmfeatures.SDFTimeSliceType{Name: %s}).Stage(stage.AixmfeaturesStage)", sdftimeslicetypeIdent, __gong__toRawStringLiteral(sdftimeslicetype.Name)))
			if lastStageVal != "AixmfeaturesStage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "AixmfeaturesStage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", sdftimeslicetypeIdent, __gong__toRawStringLiteral(sdftimeslicetype.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Id = %s", sdftimeslicetypeIdent, __gong__toRawStringLiteral(sdftimeslicetype.Id)))
			values.WriteString(fmt.Sprintf("\n\t%s.Description = %s", sdftimeslicetypeIdent, __gong__toRawStringLiteral(sdftimeslicetype.Description)))
			values.WriteString(fmt.Sprintf("\n\t%s.Interpretation = %s", sdftimeslicetypeIdent, __gong__toRawStringLiteral(sdftimeslicetype.Interpretation)))
			values.WriteString(fmt.Sprintf("\n\t%s.SequenceNumber = %d", sdftimeslicetypeIdent, sdftimeslicetype.SequenceNumber))
			values.WriteString(fmt.Sprintf("\n\t%s.CorrectionNumber = %d", sdftimeslicetypeIdent, sdftimeslicetype.CorrectionNumber))
			if sdftimeslicetype.Identifier != nil {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := sdftimeslicetype.Identifier.GongGetIdentifier(stage.GmlStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Identifier = %s", sdftimeslicetypeIdent, targetIdent))
			}
			for _, elem := range sdftimeslicetype.NameXSD {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := elem.GongGetIdentifier(stage.GmlStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.NameXSD = append(%s.NameXSD, %s)", sdftimeslicetypeIdent, sdftimeslicetypeIdent, targetIdent))
			}
			if sdftimeslicetype.ValidTime != nil {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := sdftimeslicetype.ValidTime.GongGetIdentifier(stage.GmlStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.ValidTime = %s", sdftimeslicetypeIdent, targetIdent))
			}
			if sdftimeslicetype.TimeSliceMetadata != nil {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := sdftimeslicetype.TimeSliceMetadata.GongGetIdentifier(stage.AixmabstractgmlobjecttypesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.TimeSliceMetadata = %s", sdftimeslicetypeIdent, targetIdent))
			}
			if sdftimeslicetype.FeatureLifetime != nil {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := sdftimeslicetype.FeatureLifetime.GongGetIdentifier(stage.GmlStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.FeatureLifetime = %s", sdftimeslicetypeIdent, targetIdent))
			}
			if sdftimeslicetype.Designator != nil {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := sdftimeslicetype.Designator.GongGetIdentifier(stage.AixmdatatypesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Designator = %s", sdftimeslicetypeIdent, targetIdent))
			}
			if sdftimeslicetype.EmissionClass != nil {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := sdftimeslicetype.EmissionClass.GongGetIdentifier(stage.AixmdatatypesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.EmissionClass = %s", sdftimeslicetypeIdent, targetIdent))
			}
			if sdftimeslicetype.Mobile != nil {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := sdftimeslicetype.Mobile.GongGetIdentifier(stage.AixmdatatypesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Mobile = %s", sdftimeslicetypeIdent, targetIdent))
			}
			if sdftimeslicetype.MagneticVariation != nil {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := sdftimeslicetype.MagneticVariation.GongGetIdentifier(stage.AixmdatatypesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.MagneticVariation = %s", sdftimeslicetypeIdent, targetIdent))
			}
			if sdftimeslicetype.MagneticVariationAccuracy != nil {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := sdftimeslicetype.MagneticVariationAccuracy.GongGetIdentifier(stage.AixmdatatypesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.MagneticVariationAccuracy = %s", sdftimeslicetypeIdent, targetIdent))
			}
			if sdftimeslicetype.DateMagneticVariation != nil {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := sdftimeslicetype.DateMagneticVariation.GongGetIdentifier(stage.AixmdatatypesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.DateMagneticVariation = %s", sdftimeslicetypeIdent, targetIdent))
			}
			if sdftimeslicetype.FlightChecked != nil {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := sdftimeslicetype.FlightChecked.GongGetIdentifier(stage.AixmdatatypesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.FlightChecked = %s", sdftimeslicetypeIdent, targetIdent))
			}
			if sdftimeslicetype.Location != nil {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := sdftimeslicetype.Location.GongGetIdentifier(stage.AixmfeaturesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Location = %s", sdftimeslicetypeIdent, targetIdent))
			}
			for _, elem := range sdftimeslicetype.Authority {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := elem.GongGetIdentifier(stage.AixmfeaturesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Authority = append(%s.Authority, %s)", sdftimeslicetypeIdent, sdftimeslicetypeIdent, targetIdent))
			}
			for _, elem := range sdftimeslicetype.Monitoring {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := elem.GongGetIdentifier(stage.AixmfeaturesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Monitoring = append(%s.Monitoring, %s)", sdftimeslicetypeIdent, sdftimeslicetypeIdent, targetIdent))
			}
			for _, elem := range sdftimeslicetype.Availability {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := elem.GongGetIdentifier(stage.AixmfeaturesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Availability = append(%s.Availability, %s)", sdftimeslicetypeIdent, sdftimeslicetypeIdent, targetIdent))
			}
			for _, elem := range sdftimeslicetype.Annotation {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := elem.GongGetIdentifier(stage.AixmfeaturesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Annotation = append(%s.Annotation, %s)", sdftimeslicetypeIdent, sdftimeslicetypeIdent, targetIdent))
			}
			if sdftimeslicetype.Frequency != nil {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := sdftimeslicetype.Frequency.GongGetIdentifier(stage.AixmdatatypesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Frequency = %s", sdftimeslicetypeIdent, targetIdent))
			}
			if sdftimeslicetype.MagneticBearing != nil {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := sdftimeslicetype.MagneticBearing.GongGetIdentifier(stage.AixmdatatypesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.MagneticBearing = %s", sdftimeslicetypeIdent, targetIdent))
			}
			if sdftimeslicetype.TrueBearing != nil {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := sdftimeslicetype.TrueBearing.GongGetIdentifier(stage.AixmdatatypesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.TrueBearing = %s", sdftimeslicetypeIdent, targetIdent))
			}
			for _, elem := range sdftimeslicetype.Extension {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := elem.GongGetIdentifier(stage.AixmfeaturesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Extension = append(%s.Extension, %s)", sdftimeslicetypeIdent, sdftimeslicetypeIdent, targetIdent))
			}
		}
	}
	if stage.AixmfeaturesStage != nil {
		for _, sdftype := range __gong__sortStageSetInstances(stage.AixmfeaturesStage.SDFTypes, stage.AixmfeaturesStage.SDFType_stagedOrder) {
			if lastStageDecl != "AixmfeaturesStage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "AixmfeaturesStage"
			}
			sdftypeIdent := sdftype.GongGetIdentifier(stage.AixmfeaturesStage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&aixmfeatures.SDFType{Name: %s}).Stage(stage.AixmfeaturesStage)", sdftypeIdent, __gong__toRawStringLiteral(sdftype.Name)))
			if lastStageVal != "AixmfeaturesStage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "AixmfeaturesStage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", sdftypeIdent, __gong__toRawStringLiteral(sdftype.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Id = %s", sdftypeIdent, __gong__toRawStringLiteral(sdftype.Id)))
			values.WriteString(fmt.Sprintf("\n\t%s.Description = %s", sdftypeIdent, __gong__toRawStringLiteral(sdftype.Description)))
			if sdftype.Identifier != nil {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := sdftype.Identifier.GongGetIdentifier(stage.GmlStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Identifier = %s", sdftypeIdent, targetIdent))
			}
			for _, elem := range sdftype.NameXSD {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := elem.GongGetIdentifier(stage.GmlStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.NameXSD = append(%s.NameXSD, %s)", sdftypeIdent, sdftypeIdent, targetIdent))
			}
			if sdftype.BoundedBy != nil {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := sdftype.BoundedBy.GongGetIdentifier(stage.GmlStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.BoundedBy = %s", sdftypeIdent, targetIdent))
			}
			if sdftype.FeatureMetadata != nil {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := sdftype.FeatureMetadata.GongGetIdentifier(stage.AixmabstractgmlobjecttypesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.FeatureMetadata = %s", sdftypeIdent, targetIdent))
			}
			for _, elem := range sdftype.TimeSlice {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := elem.GongGetIdentifier(stage.AixmfeaturesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.TimeSlice = append(%s.TimeSlice, %s)", sdftypeIdent, sdftypeIdent, targetIdent))
			}
		}
	}
	if stage.AixmfeaturesStage != nil {
		for _, secondarysurveillanceradartimeslicepropertytype := range __gong__sortStageSetInstances(stage.AixmfeaturesStage.SecondarySurveillanceRadarTimeSlicePropertyTypes, stage.AixmfeaturesStage.SecondarySurveillanceRadarTimeSlicePropertyType_stagedOrder) {
			if lastStageDecl != "AixmfeaturesStage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "AixmfeaturesStage"
			}
			secondarysurveillanceradartimeslicepropertytypeIdent := secondarysurveillanceradartimeslicepropertytype.GongGetIdentifier(stage.AixmfeaturesStage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&aixmfeatures.SecondarySurveillanceRadarTimeSlicePropertyType{Name: %s}).Stage(stage.AixmfeaturesStage)", secondarysurveillanceradartimeslicepropertytypeIdent, __gong__toRawStringLiteral(secondarysurveillanceradartimeslicepropertytype.Name)))
			if lastStageVal != "AixmfeaturesStage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "AixmfeaturesStage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", secondarysurveillanceradartimeslicepropertytypeIdent, __gong__toRawStringLiteral(secondarysurveillanceradartimeslicepropertytype.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Owns = %t", secondarysurveillanceradartimeslicepropertytypeIdent, secondarysurveillanceradartimeslicepropertytype.Owns))
			if secondarysurveillanceradartimeslicepropertytype.SecondarySurveillanceRadarTimeSlice != nil {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := secondarysurveillanceradartimeslicepropertytype.SecondarySurveillanceRadarTimeSlice.GongGetIdentifier(stage.AixmfeaturesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.SecondarySurveillanceRadarTimeSlice = %s", secondarysurveillanceradartimeslicepropertytypeIdent, targetIdent))
			}
		}
	}
	if stage.AixmfeaturesStage != nil {
		for _, secondarysurveillanceradartimeslicetype := range __gong__sortStageSetInstances(stage.AixmfeaturesStage.SecondarySurveillanceRadarTimeSliceTypes, stage.AixmfeaturesStage.SecondarySurveillanceRadarTimeSliceType_stagedOrder) {
			if lastStageDecl != "AixmfeaturesStage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "AixmfeaturesStage"
			}
			secondarysurveillanceradartimeslicetypeIdent := secondarysurveillanceradartimeslicetype.GongGetIdentifier(stage.AixmfeaturesStage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&aixmfeatures.SecondarySurveillanceRadarTimeSliceType{Name: %s}).Stage(stage.AixmfeaturesStage)", secondarysurveillanceradartimeslicetypeIdent, __gong__toRawStringLiteral(secondarysurveillanceradartimeslicetype.Name)))
			if lastStageVal != "AixmfeaturesStage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "AixmfeaturesStage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", secondarysurveillanceradartimeslicetypeIdent, __gong__toRawStringLiteral(secondarysurveillanceradartimeslicetype.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Id = %s", secondarysurveillanceradartimeslicetypeIdent, __gong__toRawStringLiteral(secondarysurveillanceradartimeslicetype.Id)))
			values.WriteString(fmt.Sprintf("\n\t%s.Description = %s", secondarysurveillanceradartimeslicetypeIdent, __gong__toRawStringLiteral(secondarysurveillanceradartimeslicetype.Description)))
			values.WriteString(fmt.Sprintf("\n\t%s.Interpretation = %s", secondarysurveillanceradartimeslicetypeIdent, __gong__toRawStringLiteral(secondarysurveillanceradartimeslicetype.Interpretation)))
			values.WriteString(fmt.Sprintf("\n\t%s.SequenceNumber = %d", secondarysurveillanceradartimeslicetypeIdent, secondarysurveillanceradartimeslicetype.SequenceNumber))
			values.WriteString(fmt.Sprintf("\n\t%s.CorrectionNumber = %d", secondarysurveillanceradartimeslicetypeIdent, secondarysurveillanceradartimeslicetype.CorrectionNumber))
			if secondarysurveillanceradartimeslicetype.Identifier != nil {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := secondarysurveillanceradartimeslicetype.Identifier.GongGetIdentifier(stage.GmlStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Identifier = %s", secondarysurveillanceradartimeslicetypeIdent, targetIdent))
			}
			for _, elem := range secondarysurveillanceradartimeslicetype.NameXSD {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := elem.GongGetIdentifier(stage.GmlStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.NameXSD = append(%s.NameXSD, %s)", secondarysurveillanceradartimeslicetypeIdent, secondarysurveillanceradartimeslicetypeIdent, targetIdent))
			}
			if secondarysurveillanceradartimeslicetype.ValidTime != nil {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := secondarysurveillanceradartimeslicetype.ValidTime.GongGetIdentifier(stage.GmlStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.ValidTime = %s", secondarysurveillanceradartimeslicetypeIdent, targetIdent))
			}
			if secondarysurveillanceradartimeslicetype.TimeSliceMetadata != nil {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := secondarysurveillanceradartimeslicetype.TimeSliceMetadata.GongGetIdentifier(stage.AixmabstractgmlobjecttypesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.TimeSliceMetadata = %s", secondarysurveillanceradartimeslicetypeIdent, targetIdent))
			}
			if secondarysurveillanceradartimeslicetype.FeatureLifetime != nil {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := secondarysurveillanceradartimeslicetype.FeatureLifetime.GongGetIdentifier(stage.GmlStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.FeatureLifetime = %s", secondarysurveillanceradartimeslicetypeIdent, targetIdent))
			}
			if secondarysurveillanceradartimeslicetype.SerialNumber != nil {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := secondarysurveillanceradartimeslicetype.SerialNumber.GongGetIdentifier(stage.AixmdatatypesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.SerialNumber = %s", secondarysurveillanceradartimeslicetypeIdent, targetIdent))
			}
			if secondarysurveillanceradartimeslicetype.Range_XSD != nil {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := secondarysurveillanceradartimeslicetype.Range_XSD.GongGetIdentifier(stage.AixmdatatypesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Range_XSD = %s", secondarysurveillanceradartimeslicetypeIdent, targetIdent))
			}
			if secondarysurveillanceradartimeslicetype.RangeAccuracy != nil {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := secondarysurveillanceradartimeslicetype.RangeAccuracy.GongGetIdentifier(stage.AixmdatatypesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.RangeAccuracy = %s", secondarysurveillanceradartimeslicetypeIdent, targetIdent))
			}
			if secondarysurveillanceradartimeslicetype.DualChannel != nil {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := secondarysurveillanceradartimeslicetype.DualChannel.GongGetIdentifier(stage.AixmdatatypesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.DualChannel = %s", secondarysurveillanceradartimeslicetypeIdent, targetIdent))
			}
			if secondarysurveillanceradartimeslicetype.MovingTargetIndicator != nil {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := secondarysurveillanceradartimeslicetype.MovingTargetIndicator.GongGetIdentifier(stage.AixmdatatypesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.MovingTargetIndicator = %s", secondarysurveillanceradartimeslicetypeIdent, targetIdent))
			}
			if secondarysurveillanceradartimeslicetype.StandbyPower != nil {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := secondarysurveillanceradartimeslicetype.StandbyPower.GongGetIdentifier(stage.AixmdatatypesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.StandbyPower = %s", secondarysurveillanceradartimeslicetypeIdent, targetIdent))
			}
			if secondarysurveillanceradartimeslicetype.Digital != nil {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := secondarysurveillanceradartimeslicetype.Digital.GongGetIdentifier(stage.AixmdatatypesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Digital = %s", secondarysurveillanceradartimeslicetypeIdent, targetIdent))
			}
			if secondarysurveillanceradartimeslicetype.MilitaryUseOnly != nil {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := secondarysurveillanceradartimeslicetype.MilitaryUseOnly.GongGetIdentifier(stage.AixmdatatypesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.MilitaryUseOnly = %s", secondarysurveillanceradartimeslicetypeIdent, targetIdent))
			}
			if secondarysurveillanceradartimeslicetype.SpecialUseOnly != nil {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := secondarysurveillanceradartimeslicetype.SpecialUseOnly.GongGetIdentifier(stage.AixmdatatypesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.SpecialUseOnly = %s", secondarysurveillanceradartimeslicetypeIdent, targetIdent))
			}
			if secondarysurveillanceradartimeslicetype.SpecialAircraftOnly != nil {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := secondarysurveillanceradartimeslicetype.SpecialAircraftOnly.GongGetIdentifier(stage.AixmdatatypesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.SpecialAircraftOnly = %s", secondarysurveillanceradartimeslicetypeIdent, targetIdent))
			}
			if secondarysurveillanceradartimeslicetype.MagneticVariation != nil {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := secondarysurveillanceradartimeslicetype.MagneticVariation.GongGetIdentifier(stage.AixmdatatypesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.MagneticVariation = %s", secondarysurveillanceradartimeslicetypeIdent, targetIdent))
			}
			if secondarysurveillanceradartimeslicetype.MagneticVariationAccuracy != nil {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := secondarysurveillanceradartimeslicetype.MagneticVariationAccuracy.GongGetIdentifier(stage.AixmdatatypesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.MagneticVariationAccuracy = %s", secondarysurveillanceradartimeslicetypeIdent, targetIdent))
			}
			if secondarysurveillanceradartimeslicetype.DateMagneticVariation != nil {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := secondarysurveillanceradartimeslicetype.DateMagneticVariation.GongGetIdentifier(stage.AixmdatatypesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.DateMagneticVariation = %s", secondarysurveillanceradartimeslicetypeIdent, targetIdent))
			}
			for _, elem := range secondarysurveillanceradartimeslicetype.Contact {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := elem.GongGetIdentifier(stage.AixmfeaturesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Contact = append(%s.Contact, %s)", secondarysurveillanceradartimeslicetypeIdent, secondarysurveillanceradartimeslicetypeIdent, targetIdent))
			}
			if secondarysurveillanceradartimeslicetype.Location != nil {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := secondarysurveillanceradartimeslicetype.Location.GongGetIdentifier(stage.AixmfeaturesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Location = %s", secondarysurveillanceradartimeslicetypeIdent, targetIdent))
			}
			for _, elem := range secondarysurveillanceradartimeslicetype.Annotation {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := elem.GongGetIdentifier(stage.AixmfeaturesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Annotation = append(%s.Annotation, %s)", secondarysurveillanceradartimeslicetypeIdent, secondarysurveillanceradartimeslicetypeIdent, targetIdent))
			}
			if secondarysurveillanceradartimeslicetype.VerticalCoverageAltitude != nil {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := secondarysurveillanceradartimeslicetype.VerticalCoverageAltitude.GongGetIdentifier(stage.AixmdatatypesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.VerticalCoverageAltitude = %s", secondarysurveillanceradartimeslicetypeIdent, targetIdent))
			}
			if secondarysurveillanceradartimeslicetype.VerticalCoverageDistance != nil {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := secondarysurveillanceradartimeslicetype.VerticalCoverageDistance.GongGetIdentifier(stage.AixmdatatypesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.VerticalCoverageDistance = %s", secondarysurveillanceradartimeslicetypeIdent, targetIdent))
			}
			if secondarysurveillanceradartimeslicetype.VerticalCoverageAzimuth != nil {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := secondarysurveillanceradartimeslicetype.VerticalCoverageAzimuth.GongGetIdentifier(stage.AixmdatatypesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.VerticalCoverageAzimuth = %s", secondarysurveillanceradartimeslicetypeIdent, targetIdent))
			}
			if secondarysurveillanceradartimeslicetype.AntennaTiltFixed != nil {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := secondarysurveillanceradartimeslicetype.AntennaTiltFixed.GongGetIdentifier(stage.AixmdatatypesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.AntennaTiltFixed = %s", secondarysurveillanceradartimeslicetypeIdent, targetIdent))
			}
			if secondarysurveillanceradartimeslicetype.TiltAngle != nil {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := secondarysurveillanceradartimeslicetype.TiltAngle.GongGetIdentifier(stage.AixmdatatypesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.TiltAngle = %s", secondarysurveillanceradartimeslicetypeIdent, targetIdent))
			}
			if secondarysurveillanceradartimeslicetype.AutomatedRadarTerminalSystem != nil {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := secondarysurveillanceradartimeslicetype.AutomatedRadarTerminalSystem.GongGetIdentifier(stage.AixmdatatypesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.AutomatedRadarTerminalSystem = %s", secondarysurveillanceradartimeslicetypeIdent, targetIdent))
			}
			for _, elem := range secondarysurveillanceradartimeslicetype.GroundStation {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := elem.GongGetIdentifier(stage.AixmfeaturesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.GroundStation = append(%s.GroundStation, %s)", secondarysurveillanceradartimeslicetypeIdent, secondarysurveillanceradartimeslicetypeIdent, targetIdent))
			}
			if secondarysurveillanceradartimeslicetype.Transponder != nil {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := secondarysurveillanceradartimeslicetype.Transponder.GongGetIdentifier(stage.AixmdatatypesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Transponder = %s", secondarysurveillanceradartimeslicetypeIdent, targetIdent))
			}
			if secondarysurveillanceradartimeslicetype.Autonomous != nil {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := secondarysurveillanceradartimeslicetype.Autonomous.GongGetIdentifier(stage.AixmdatatypesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Autonomous = %s", secondarysurveillanceradartimeslicetypeIdent, targetIdent))
			}
			if secondarysurveillanceradartimeslicetype.Monopulse != nil {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := secondarysurveillanceradartimeslicetype.Monopulse.GongGetIdentifier(stage.AixmdatatypesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Monopulse = %s", secondarysurveillanceradartimeslicetypeIdent, targetIdent))
			}
			for _, elem := range secondarysurveillanceradartimeslicetype.Extension {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := elem.GongGetIdentifier(stage.AixmfeaturesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Extension = append(%s.Extension, %s)", secondarysurveillanceradartimeslicetypeIdent, secondarysurveillanceradartimeslicetypeIdent, targetIdent))
			}
		}
	}
	if stage.AixmfeaturesStage != nil {
		for _, secondarysurveillanceradartype := range __gong__sortStageSetInstances(stage.AixmfeaturesStage.SecondarySurveillanceRadarTypes, stage.AixmfeaturesStage.SecondarySurveillanceRadarType_stagedOrder) {
			if lastStageDecl != "AixmfeaturesStage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "AixmfeaturesStage"
			}
			secondarysurveillanceradartypeIdent := secondarysurveillanceradartype.GongGetIdentifier(stage.AixmfeaturesStage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&aixmfeatures.SecondarySurveillanceRadarType{Name: %s}).Stage(stage.AixmfeaturesStage)", secondarysurveillanceradartypeIdent, __gong__toRawStringLiteral(secondarysurveillanceradartype.Name)))
			if lastStageVal != "AixmfeaturesStage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "AixmfeaturesStage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", secondarysurveillanceradartypeIdent, __gong__toRawStringLiteral(secondarysurveillanceradartype.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Id = %s", secondarysurveillanceradartypeIdent, __gong__toRawStringLiteral(secondarysurveillanceradartype.Id)))
			values.WriteString(fmt.Sprintf("\n\t%s.Description = %s", secondarysurveillanceradartypeIdent, __gong__toRawStringLiteral(secondarysurveillanceradartype.Description)))
			if secondarysurveillanceradartype.Identifier != nil {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := secondarysurveillanceradartype.Identifier.GongGetIdentifier(stage.GmlStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Identifier = %s", secondarysurveillanceradartypeIdent, targetIdent))
			}
			for _, elem := range secondarysurveillanceradartype.NameXSD {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := elem.GongGetIdentifier(stage.GmlStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.NameXSD = append(%s.NameXSD, %s)", secondarysurveillanceradartypeIdent, secondarysurveillanceradartypeIdent, targetIdent))
			}
			if secondarysurveillanceradartype.BoundedBy != nil {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := secondarysurveillanceradartype.BoundedBy.GongGetIdentifier(stage.GmlStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.BoundedBy = %s", secondarysurveillanceradartypeIdent, targetIdent))
			}
			if secondarysurveillanceradartype.FeatureMetadata != nil {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := secondarysurveillanceradartype.FeatureMetadata.GongGetIdentifier(stage.AixmabstractgmlobjecttypesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.FeatureMetadata = %s", secondarysurveillanceradartypeIdent, targetIdent))
			}
			for _, elem := range secondarysurveillanceradartype.TimeSlice {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := elem.GongGetIdentifier(stage.AixmfeaturesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.TimeSlice = append(%s.TimeSlice, %s)", secondarysurveillanceradartypeIdent, secondarysurveillanceradartypeIdent, targetIdent))
			}
		}
	}
	if stage.AixmfeaturesStage != nil {
		for _, significantpointinairspacetimeslicepropertytype := range __gong__sortStageSetInstances(stage.AixmfeaturesStage.SignificantPointInAirspaceTimeSlicePropertyTypes, stage.AixmfeaturesStage.SignificantPointInAirspaceTimeSlicePropertyType_stagedOrder) {
			if lastStageDecl != "AixmfeaturesStage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "AixmfeaturesStage"
			}
			significantpointinairspacetimeslicepropertytypeIdent := significantpointinairspacetimeslicepropertytype.GongGetIdentifier(stage.AixmfeaturesStage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&aixmfeatures.SignificantPointInAirspaceTimeSlicePropertyType{Name: %s}).Stage(stage.AixmfeaturesStage)", significantpointinairspacetimeslicepropertytypeIdent, __gong__toRawStringLiteral(significantpointinairspacetimeslicepropertytype.Name)))
			if lastStageVal != "AixmfeaturesStage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "AixmfeaturesStage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", significantpointinairspacetimeslicepropertytypeIdent, __gong__toRawStringLiteral(significantpointinairspacetimeslicepropertytype.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Owns = %t", significantpointinairspacetimeslicepropertytypeIdent, significantpointinairspacetimeslicepropertytype.Owns))
			if significantpointinairspacetimeslicepropertytype.SignificantPointInAirspaceTimeSlice != nil {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := significantpointinairspacetimeslicepropertytype.SignificantPointInAirspaceTimeSlice.GongGetIdentifier(stage.AixmfeaturesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.SignificantPointInAirspaceTimeSlice = %s", significantpointinairspacetimeslicepropertytypeIdent, targetIdent))
			}
		}
	}
	if stage.AixmfeaturesStage != nil {
		for _, significantpointinairspacetimeslicetype := range __gong__sortStageSetInstances(stage.AixmfeaturesStage.SignificantPointInAirspaceTimeSliceTypes, stage.AixmfeaturesStage.SignificantPointInAirspaceTimeSliceType_stagedOrder) {
			if lastStageDecl != "AixmfeaturesStage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "AixmfeaturesStage"
			}
			significantpointinairspacetimeslicetypeIdent := significantpointinairspacetimeslicetype.GongGetIdentifier(stage.AixmfeaturesStage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&aixmfeatures.SignificantPointInAirspaceTimeSliceType{Name: %s}).Stage(stage.AixmfeaturesStage)", significantpointinairspacetimeslicetypeIdent, __gong__toRawStringLiteral(significantpointinairspacetimeslicetype.Name)))
			if lastStageVal != "AixmfeaturesStage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "AixmfeaturesStage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", significantpointinairspacetimeslicetypeIdent, __gong__toRawStringLiteral(significantpointinairspacetimeslicetype.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Id = %s", significantpointinairspacetimeslicetypeIdent, __gong__toRawStringLiteral(significantpointinairspacetimeslicetype.Id)))
			values.WriteString(fmt.Sprintf("\n\t%s.Description = %s", significantpointinairspacetimeslicetypeIdent, __gong__toRawStringLiteral(significantpointinairspacetimeslicetype.Description)))
			values.WriteString(fmt.Sprintf("\n\t%s.Interpretation = %s", significantpointinairspacetimeslicetypeIdent, __gong__toRawStringLiteral(significantpointinairspacetimeslicetype.Interpretation)))
			values.WriteString(fmt.Sprintf("\n\t%s.SequenceNumber = %d", significantpointinairspacetimeslicetypeIdent, significantpointinairspacetimeslicetype.SequenceNumber))
			values.WriteString(fmt.Sprintf("\n\t%s.CorrectionNumber = %d", significantpointinairspacetimeslicetypeIdent, significantpointinairspacetimeslicetype.CorrectionNumber))
			if significantpointinairspacetimeslicetype.Identifier != nil {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := significantpointinairspacetimeslicetype.Identifier.GongGetIdentifier(stage.GmlStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Identifier = %s", significantpointinairspacetimeslicetypeIdent, targetIdent))
			}
			for _, elem := range significantpointinairspacetimeslicetype.NameXSD {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := elem.GongGetIdentifier(stage.GmlStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.NameXSD = append(%s.NameXSD, %s)", significantpointinairspacetimeslicetypeIdent, significantpointinairspacetimeslicetypeIdent, targetIdent))
			}
			if significantpointinairspacetimeslicetype.ValidTime != nil {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := significantpointinairspacetimeslicetype.ValidTime.GongGetIdentifier(stage.GmlStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.ValidTime = %s", significantpointinairspacetimeslicetypeIdent, targetIdent))
			}
			if significantpointinairspacetimeslicetype.TimeSliceMetadata != nil {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := significantpointinairspacetimeslicetype.TimeSliceMetadata.GongGetIdentifier(stage.AixmabstractgmlobjecttypesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.TimeSliceMetadata = %s", significantpointinairspacetimeslicetypeIdent, targetIdent))
			}
			if significantpointinairspacetimeslicetype.FeatureLifetime != nil {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := significantpointinairspacetimeslicetype.FeatureLifetime.GongGetIdentifier(stage.GmlStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.FeatureLifetime = %s", significantpointinairspacetimeslicetypeIdent, targetIdent))
			}
			if significantpointinairspacetimeslicetype.Type_XSD != nil {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := significantpointinairspacetimeslicetype.Type_XSD.GongGetIdentifier(stage.AixmdatatypesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Type_XSD = %s", significantpointinairspacetimeslicetypeIdent, targetIdent))
			}
			if significantpointinairspacetimeslicetype.RelativeLocation != nil {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := significantpointinairspacetimeslicetype.RelativeLocation.GongGetIdentifier(stage.AixmdatatypesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.RelativeLocation = %s", significantpointinairspacetimeslicetypeIdent, targetIdent))
			}
			if significantpointinairspacetimeslicetype.ContainingAirspace != nil {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := significantpointinairspacetimeslicetype.ContainingAirspace.GongGetIdentifier(stage.AixmfeaturesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.ContainingAirspace = %s", significantpointinairspacetimeslicetypeIdent, targetIdent))
			}
			if significantpointinairspacetimeslicetype.Location_fixDesignatedPoint != nil {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := significantpointinairspacetimeslicetype.Location_fixDesignatedPoint.GongGetIdentifier(stage.AixmfeaturesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Location_fixDesignatedPoint = %s", significantpointinairspacetimeslicetypeIdent, targetIdent))
			}
			if significantpointinairspacetimeslicetype.Location_navaidSystem != nil {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := significantpointinairspacetimeslicetype.Location_navaidSystem.GongGetIdentifier(stage.AixmfeaturesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Location_navaidSystem = %s", significantpointinairspacetimeslicetypeIdent, targetIdent))
			}
			if significantpointinairspacetimeslicetype.Location_position != nil {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := significantpointinairspacetimeslicetype.Location_position.GongGetIdentifier(stage.AixmfeaturesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Location_position = %s", significantpointinairspacetimeslicetypeIdent, targetIdent))
			}
			if significantpointinairspacetimeslicetype.Location_runwayPoint != nil {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := significantpointinairspacetimeslicetype.Location_runwayPoint.GongGetIdentifier(stage.AixmfeaturesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Location_runwayPoint = %s", significantpointinairspacetimeslicetypeIdent, targetIdent))
			}
			if significantpointinairspacetimeslicetype.Location_aimingPoint != nil {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := significantpointinairspacetimeslicetype.Location_aimingPoint.GongGetIdentifier(stage.AixmfeaturesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Location_aimingPoint = %s", significantpointinairspacetimeslicetypeIdent, targetIdent))
			}
			if significantpointinairspacetimeslicetype.Location_airportReferencePoint != nil {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := significantpointinairspacetimeslicetype.Location_airportReferencePoint.GongGetIdentifier(stage.AixmfeaturesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Location_airportReferencePoint = %s", significantpointinairspacetimeslicetypeIdent, targetIdent))
			}
			for _, elem := range significantpointinairspacetimeslicetype.Annotation {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := elem.GongGetIdentifier(stage.AixmfeaturesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Annotation = append(%s.Annotation, %s)", significantpointinairspacetimeslicetypeIdent, significantpointinairspacetimeslicetypeIdent, targetIdent))
			}
			for _, elem := range significantpointinairspacetimeslicetype.Extension {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := elem.GongGetIdentifier(stage.AixmfeaturesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Extension = append(%s.Extension, %s)", significantpointinairspacetimeslicetypeIdent, significantpointinairspacetimeslicetypeIdent, targetIdent))
			}
		}
	}
	if stage.AixmfeaturesStage != nil {
		for _, significantpointinairspacetype := range __gong__sortStageSetInstances(stage.AixmfeaturesStage.SignificantPointInAirspaceTypes, stage.AixmfeaturesStage.SignificantPointInAirspaceType_stagedOrder) {
			if lastStageDecl != "AixmfeaturesStage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "AixmfeaturesStage"
			}
			significantpointinairspacetypeIdent := significantpointinairspacetype.GongGetIdentifier(stage.AixmfeaturesStage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&aixmfeatures.SignificantPointInAirspaceType{Name: %s}).Stage(stage.AixmfeaturesStage)", significantpointinairspacetypeIdent, __gong__toRawStringLiteral(significantpointinairspacetype.Name)))
			if lastStageVal != "AixmfeaturesStage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "AixmfeaturesStage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", significantpointinairspacetypeIdent, __gong__toRawStringLiteral(significantpointinairspacetype.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Id = %s", significantpointinairspacetypeIdent, __gong__toRawStringLiteral(significantpointinairspacetype.Id)))
			values.WriteString(fmt.Sprintf("\n\t%s.Description = %s", significantpointinairspacetypeIdent, __gong__toRawStringLiteral(significantpointinairspacetype.Description)))
			if significantpointinairspacetype.Identifier != nil {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := significantpointinairspacetype.Identifier.GongGetIdentifier(stage.GmlStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Identifier = %s", significantpointinairspacetypeIdent, targetIdent))
			}
			for _, elem := range significantpointinairspacetype.NameXSD {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := elem.GongGetIdentifier(stage.GmlStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.NameXSD = append(%s.NameXSD, %s)", significantpointinairspacetypeIdent, significantpointinairspacetypeIdent, targetIdent))
			}
			if significantpointinairspacetype.BoundedBy != nil {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := significantpointinairspacetype.BoundedBy.GongGetIdentifier(stage.GmlStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.BoundedBy = %s", significantpointinairspacetypeIdent, targetIdent))
			}
			if significantpointinairspacetype.FeatureMetadata != nil {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := significantpointinairspacetype.FeatureMetadata.GongGetIdentifier(stage.AixmabstractgmlobjecttypesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.FeatureMetadata = %s", significantpointinairspacetypeIdent, targetIdent))
			}
			for _, elem := range significantpointinairspacetype.TimeSlice {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := elem.GongGetIdentifier(stage.AixmfeaturesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.TimeSlice = append(%s.TimeSlice, %s)", significantpointinairspacetypeIdent, significantpointinairspacetypeIdent, targetIdent))
			}
		}
	}
	if stage.AixmfeaturesStage != nil {
		for _, specialnavigationstationstatuspropertytype := range __gong__sortStageSetInstances(stage.AixmfeaturesStage.SpecialNavigationStationStatusPropertyTypes, stage.AixmfeaturesStage.SpecialNavigationStationStatusPropertyType_stagedOrder) {
			if lastStageDecl != "AixmfeaturesStage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "AixmfeaturesStage"
			}
			specialnavigationstationstatuspropertytypeIdent := specialnavigationstationstatuspropertytype.GongGetIdentifier(stage.AixmfeaturesStage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&aixmfeatures.SpecialNavigationStationStatusPropertyType{Name: %s}).Stage(stage.AixmfeaturesStage)", specialnavigationstationstatuspropertytypeIdent, __gong__toRawStringLiteral(specialnavigationstationstatuspropertytype.Name)))
			if lastStageVal != "AixmfeaturesStage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "AixmfeaturesStage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", specialnavigationstationstatuspropertytypeIdent, __gong__toRawStringLiteral(specialnavigationstationstatuspropertytype.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.NilReason = %s", specialnavigationstationstatuspropertytypeIdent, __gong__toRawStringLiteral(specialnavigationstationstatuspropertytype.NilReason)))
			if specialnavigationstationstatuspropertytype.SpecialNavigationStationStatus != nil {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := specialnavigationstationstatuspropertytype.SpecialNavigationStationStatus.GongGetIdentifier(stage.AixmfeaturesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.SpecialNavigationStationStatus = %s", specialnavigationstationstatuspropertytypeIdent, targetIdent))
			}
		}
	}
	if stage.AixmfeaturesStage != nil {
		for _, specialnavigationstationstatustype := range __gong__sortStageSetInstances(stage.AixmfeaturesStage.SpecialNavigationStationStatusTypes, stage.AixmfeaturesStage.SpecialNavigationStationStatusType_stagedOrder) {
			if lastStageDecl != "AixmfeaturesStage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "AixmfeaturesStage"
			}
			specialnavigationstationstatustypeIdent := specialnavigationstationstatustype.GongGetIdentifier(stage.AixmfeaturesStage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&aixmfeatures.SpecialNavigationStationStatusType{Name: %s}).Stage(stage.AixmfeaturesStage)", specialnavigationstationstatustypeIdent, __gong__toRawStringLiteral(specialnavigationstationstatustype.Name)))
			if lastStageVal != "AixmfeaturesStage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "AixmfeaturesStage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", specialnavigationstationstatustypeIdent, __gong__toRawStringLiteral(specialnavigationstationstatustype.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Id = %s", specialnavigationstationstatustypeIdent, __gong__toRawStringLiteral(specialnavigationstationstatustype.Id)))
			values.WriteString(fmt.Sprintf("\n\t%s.Description = %s", specialnavigationstationstatustypeIdent, __gong__toRawStringLiteral(specialnavigationstationstatustype.Description)))
			if specialnavigationstationstatustype.Identifier != nil {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := specialnavigationstationstatustype.Identifier.GongGetIdentifier(stage.GmlStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Identifier = %s", specialnavigationstationstatustypeIdent, targetIdent))
			}
			for _, elem := range specialnavigationstationstatustype.NameXSD {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := elem.GongGetIdentifier(stage.GmlStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.NameXSD = append(%s.NameXSD, %s)", specialnavigationstationstatustypeIdent, specialnavigationstationstatustypeIdent, targetIdent))
			}
			for _, elem := range specialnavigationstationstatustype.TimeInterval {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := elem.GongGetIdentifier(stage.AixmfeaturesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.TimeInterval = append(%s.TimeInterval, %s)", specialnavigationstationstatustypeIdent, specialnavigationstationstatustypeIdent, targetIdent))
			}
			for _, elem := range specialnavigationstationstatustype.Annotation {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := elem.GongGetIdentifier(stage.AixmfeaturesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Annotation = append(%s.Annotation, %s)", specialnavigationstationstatustypeIdent, specialnavigationstationstatustypeIdent, targetIdent))
			}
			for _, elem := range specialnavigationstationstatustype.SpecialDateAuthority {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := elem.GongGetIdentifier(stage.AixmfeaturesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.SpecialDateAuthority = append(%s.SpecialDateAuthority, %s)", specialnavigationstationstatustypeIdent, specialnavigationstationstatustypeIdent, targetIdent))
			}
			if specialnavigationstationstatustype.OperationalStatus != nil {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := specialnavigationstationstatustype.OperationalStatus.GongGetIdentifier(stage.AixmdatatypesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.OperationalStatus = %s", specialnavigationstationstatustypeIdent, targetIdent))
			}
			for _, elem := range specialnavigationstationstatustype.Extension {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := elem.GongGetIdentifier(stage.AixmfeaturesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Extension = append(%s.Extension, %s)", specialnavigationstationstatustypeIdent, specialnavigationstationstatustypeIdent, targetIdent))
			}
		}
	}
	if stage.AixmfeaturesStage != nil {
		for _, specialnavigationstationtimeslicepropertytype := range __gong__sortStageSetInstances(stage.AixmfeaturesStage.SpecialNavigationStationTimeSlicePropertyTypes, stage.AixmfeaturesStage.SpecialNavigationStationTimeSlicePropertyType_stagedOrder) {
			if lastStageDecl != "AixmfeaturesStage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "AixmfeaturesStage"
			}
			specialnavigationstationtimeslicepropertytypeIdent := specialnavigationstationtimeslicepropertytype.GongGetIdentifier(stage.AixmfeaturesStage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&aixmfeatures.SpecialNavigationStationTimeSlicePropertyType{Name: %s}).Stage(stage.AixmfeaturesStage)", specialnavigationstationtimeslicepropertytypeIdent, __gong__toRawStringLiteral(specialnavigationstationtimeslicepropertytype.Name)))
			if lastStageVal != "AixmfeaturesStage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "AixmfeaturesStage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", specialnavigationstationtimeslicepropertytypeIdent, __gong__toRawStringLiteral(specialnavigationstationtimeslicepropertytype.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Owns = %t", specialnavigationstationtimeslicepropertytypeIdent, specialnavigationstationtimeslicepropertytype.Owns))
			if specialnavigationstationtimeslicepropertytype.SpecialNavigationStationTimeSlice != nil {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := specialnavigationstationtimeslicepropertytype.SpecialNavigationStationTimeSlice.GongGetIdentifier(stage.AixmfeaturesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.SpecialNavigationStationTimeSlice = %s", specialnavigationstationtimeslicepropertytypeIdent, targetIdent))
			}
		}
	}
	if stage.AixmfeaturesStage != nil {
		for _, specialnavigationstationtimeslicetype := range __gong__sortStageSetInstances(stage.AixmfeaturesStage.SpecialNavigationStationTimeSliceTypes, stage.AixmfeaturesStage.SpecialNavigationStationTimeSliceType_stagedOrder) {
			if lastStageDecl != "AixmfeaturesStage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "AixmfeaturesStage"
			}
			specialnavigationstationtimeslicetypeIdent := specialnavigationstationtimeslicetype.GongGetIdentifier(stage.AixmfeaturesStage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&aixmfeatures.SpecialNavigationStationTimeSliceType{Name: %s}).Stage(stage.AixmfeaturesStage)", specialnavigationstationtimeslicetypeIdent, __gong__toRawStringLiteral(specialnavigationstationtimeslicetype.Name)))
			if lastStageVal != "AixmfeaturesStage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "AixmfeaturesStage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", specialnavigationstationtimeslicetypeIdent, __gong__toRawStringLiteral(specialnavigationstationtimeslicetype.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Id = %s", specialnavigationstationtimeslicetypeIdent, __gong__toRawStringLiteral(specialnavigationstationtimeslicetype.Id)))
			values.WriteString(fmt.Sprintf("\n\t%s.Description = %s", specialnavigationstationtimeslicetypeIdent, __gong__toRawStringLiteral(specialnavigationstationtimeslicetype.Description)))
			values.WriteString(fmt.Sprintf("\n\t%s.Interpretation = %s", specialnavigationstationtimeslicetypeIdent, __gong__toRawStringLiteral(specialnavigationstationtimeslicetype.Interpretation)))
			values.WriteString(fmt.Sprintf("\n\t%s.SequenceNumber = %d", specialnavigationstationtimeslicetypeIdent, specialnavigationstationtimeslicetype.SequenceNumber))
			values.WriteString(fmt.Sprintf("\n\t%s.CorrectionNumber = %d", specialnavigationstationtimeslicetypeIdent, specialnavigationstationtimeslicetype.CorrectionNumber))
			if specialnavigationstationtimeslicetype.Identifier != nil {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := specialnavigationstationtimeslicetype.Identifier.GongGetIdentifier(stage.GmlStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Identifier = %s", specialnavigationstationtimeslicetypeIdent, targetIdent))
			}
			for _, elem := range specialnavigationstationtimeslicetype.NameXSD {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := elem.GongGetIdentifier(stage.GmlStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.NameXSD = append(%s.NameXSD, %s)", specialnavigationstationtimeslicetypeIdent, specialnavigationstationtimeslicetypeIdent, targetIdent))
			}
			if specialnavigationstationtimeslicetype.ValidTime != nil {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := specialnavigationstationtimeslicetype.ValidTime.GongGetIdentifier(stage.GmlStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.ValidTime = %s", specialnavigationstationtimeslicetypeIdent, targetIdent))
			}
			if specialnavigationstationtimeslicetype.TimeSliceMetadata != nil {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := specialnavigationstationtimeslicetype.TimeSliceMetadata.GongGetIdentifier(stage.AixmabstractgmlobjecttypesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.TimeSliceMetadata = %s", specialnavigationstationtimeslicetypeIdent, targetIdent))
			}
			if specialnavigationstationtimeslicetype.FeatureLifetime != nil {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := specialnavigationstationtimeslicetype.FeatureLifetime.GongGetIdentifier(stage.GmlStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.FeatureLifetime = %s", specialnavigationstationtimeslicetypeIdent, targetIdent))
			}
			if specialnavigationstationtimeslicetype.Type_XSD != nil {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := specialnavigationstationtimeslicetype.Type_XSD.GongGetIdentifier(stage.AixmdatatypesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Type_XSD = %s", specialnavigationstationtimeslicetypeIdent, targetIdent))
			}
			if specialnavigationstationtimeslicetype.Frequency != nil {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := specialnavigationstationtimeslicetype.Frequency.GongGetIdentifier(stage.AixmdatatypesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Frequency = %s", specialnavigationstationtimeslicetypeIdent, targetIdent))
			}
			if specialnavigationstationtimeslicetype.Emission != nil {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := specialnavigationstationtimeslicetype.Emission.GongGetIdentifier(stage.AixmdatatypesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Emission = %s", specialnavigationstationtimeslicetypeIdent, targetIdent))
			}
			if specialnavigationstationtimeslicetype.SystemChain != nil {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := specialnavigationstationtimeslicetype.SystemChain.GongGetIdentifier(stage.AixmfeaturesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.SystemChain = %s", specialnavigationstationtimeslicetypeIdent, targetIdent))
			}
			if specialnavigationstationtimeslicetype.ResponsibleOrganisation != nil {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := specialnavigationstationtimeslicetype.ResponsibleOrganisation.GongGetIdentifier(stage.AixmfeaturesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.ResponsibleOrganisation = %s", specialnavigationstationtimeslicetypeIdent, targetIdent))
			}
			if specialnavigationstationtimeslicetype.Position != nil {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := specialnavigationstationtimeslicetype.Position.GongGetIdentifier(stage.AixmfeaturesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Position = %s", specialnavigationstationtimeslicetypeIdent, targetIdent))
			}
			for _, elem := range specialnavigationstationtimeslicetype.Availability {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := elem.GongGetIdentifier(stage.AixmfeaturesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Availability = append(%s.Availability, %s)", specialnavigationstationtimeslicetypeIdent, specialnavigationstationtimeslicetypeIdent, targetIdent))
			}
			for _, elem := range specialnavigationstationtimeslicetype.Annotation {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := elem.GongGetIdentifier(stage.AixmfeaturesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Annotation = append(%s.Annotation, %s)", specialnavigationstationtimeslicetypeIdent, specialnavigationstationtimeslicetypeIdent, targetIdent))
			}
			for _, elem := range specialnavigationstationtimeslicetype.Extension {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := elem.GongGetIdentifier(stage.AixmfeaturesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Extension = append(%s.Extension, %s)", specialnavigationstationtimeslicetypeIdent, specialnavigationstationtimeslicetypeIdent, targetIdent))
			}
		}
	}
	if stage.AixmfeaturesStage != nil {
		for _, specialnavigationstationtype := range __gong__sortStageSetInstances(stage.AixmfeaturesStage.SpecialNavigationStationTypes, stage.AixmfeaturesStage.SpecialNavigationStationType_stagedOrder) {
			if lastStageDecl != "AixmfeaturesStage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "AixmfeaturesStage"
			}
			specialnavigationstationtypeIdent := specialnavigationstationtype.GongGetIdentifier(stage.AixmfeaturesStage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&aixmfeatures.SpecialNavigationStationType{Name: %s}).Stage(stage.AixmfeaturesStage)", specialnavigationstationtypeIdent, __gong__toRawStringLiteral(specialnavigationstationtype.Name)))
			if lastStageVal != "AixmfeaturesStage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "AixmfeaturesStage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", specialnavigationstationtypeIdent, __gong__toRawStringLiteral(specialnavigationstationtype.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Id = %s", specialnavigationstationtypeIdent, __gong__toRawStringLiteral(specialnavigationstationtype.Id)))
			values.WriteString(fmt.Sprintf("\n\t%s.Description = %s", specialnavigationstationtypeIdent, __gong__toRawStringLiteral(specialnavigationstationtype.Description)))
			if specialnavigationstationtype.Identifier != nil {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := specialnavigationstationtype.Identifier.GongGetIdentifier(stage.GmlStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Identifier = %s", specialnavigationstationtypeIdent, targetIdent))
			}
			for _, elem := range specialnavigationstationtype.NameXSD {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := elem.GongGetIdentifier(stage.GmlStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.NameXSD = append(%s.NameXSD, %s)", specialnavigationstationtypeIdent, specialnavigationstationtypeIdent, targetIdent))
			}
			if specialnavigationstationtype.BoundedBy != nil {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := specialnavigationstationtype.BoundedBy.GongGetIdentifier(stage.GmlStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.BoundedBy = %s", specialnavigationstationtypeIdent, targetIdent))
			}
			if specialnavigationstationtype.FeatureMetadata != nil {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := specialnavigationstationtype.FeatureMetadata.GongGetIdentifier(stage.AixmabstractgmlobjecttypesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.FeatureMetadata = %s", specialnavigationstationtypeIdent, targetIdent))
			}
			for _, elem := range specialnavigationstationtype.TimeSlice {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := elem.GongGetIdentifier(stage.AixmfeaturesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.TimeSlice = append(%s.TimeSlice, %s)", specialnavigationstationtypeIdent, specialnavigationstationtypeIdent, targetIdent))
			}
		}
	}
	if stage.AixmfeaturesStage != nil {
		for _, specialnavigationsystempropertytype := range __gong__sortStageSetInstances(stage.AixmfeaturesStage.SpecialNavigationSystemPropertyTypes, stage.AixmfeaturesStage.SpecialNavigationSystemPropertyType_stagedOrder) {
			if lastStageDecl != "AixmfeaturesStage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "AixmfeaturesStage"
			}
			specialnavigationsystempropertytypeIdent := specialnavigationsystempropertytype.GongGetIdentifier(stage.AixmfeaturesStage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&aixmfeatures.SpecialNavigationSystemPropertyType{Name: %s}).Stage(stage.AixmfeaturesStage)", specialnavigationsystempropertytypeIdent, __gong__toRawStringLiteral(specialnavigationsystempropertytype.Name)))
			if lastStageVal != "AixmfeaturesStage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "AixmfeaturesStage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", specialnavigationsystempropertytypeIdent, __gong__toRawStringLiteral(specialnavigationsystempropertytype.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Owns = %t", specialnavigationsystempropertytypeIdent, specialnavigationsystempropertytype.Owns))
			values.WriteString(fmt.Sprintf("\n\t%s.NilReason = %s", specialnavigationsystempropertytypeIdent, __gong__toRawStringLiteral(specialnavigationsystempropertytype.NilReason)))
			values.WriteString(fmt.Sprintf("\n\t%s.Type_XSD = %s", specialnavigationsystempropertytypeIdent, __gong__toRawStringLiteral(specialnavigationsystempropertytype.Type_XSD)))
			values.WriteString(fmt.Sprintf("\n\t%s.Href = %s", specialnavigationsystempropertytypeIdent, __gong__toRawStringLiteral(specialnavigationsystempropertytype.Href)))
			values.WriteString(fmt.Sprintf("\n\t%s.Role = %s", specialnavigationsystempropertytypeIdent, __gong__toRawStringLiteral(specialnavigationsystempropertytype.Role)))
			values.WriteString(fmt.Sprintf("\n\t%s.Arcrole = %s", specialnavigationsystempropertytypeIdent, __gong__toRawStringLiteral(specialnavigationsystempropertytype.Arcrole)))
			values.WriteString(fmt.Sprintf("\n\t%s.Title = %s", specialnavigationsystempropertytypeIdent, __gong__toRawStringLiteral(specialnavigationsystempropertytype.Title)))
			values.WriteString(fmt.Sprintf("\n\t%s.Show = %s", specialnavigationsystempropertytypeIdent, __gong__toRawStringLiteral(specialnavigationsystempropertytype.Show)))
			values.WriteString(fmt.Sprintf("\n\t%s.Actuate = %s", specialnavigationsystempropertytypeIdent, __gong__toRawStringLiteral(specialnavigationsystempropertytype.Actuate)))
		}
	}
	if stage.AixmfeaturesStage != nil {
		for _, specialnavigationsystemtimeslicepropertytype := range __gong__sortStageSetInstances(stage.AixmfeaturesStage.SpecialNavigationSystemTimeSlicePropertyTypes, stage.AixmfeaturesStage.SpecialNavigationSystemTimeSlicePropertyType_stagedOrder) {
			if lastStageDecl != "AixmfeaturesStage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "AixmfeaturesStage"
			}
			specialnavigationsystemtimeslicepropertytypeIdent := specialnavigationsystemtimeslicepropertytype.GongGetIdentifier(stage.AixmfeaturesStage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&aixmfeatures.SpecialNavigationSystemTimeSlicePropertyType{Name: %s}).Stage(stage.AixmfeaturesStage)", specialnavigationsystemtimeslicepropertytypeIdent, __gong__toRawStringLiteral(specialnavigationsystemtimeslicepropertytype.Name)))
			if lastStageVal != "AixmfeaturesStage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "AixmfeaturesStage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", specialnavigationsystemtimeslicepropertytypeIdent, __gong__toRawStringLiteral(specialnavigationsystemtimeslicepropertytype.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Owns = %t", specialnavigationsystemtimeslicepropertytypeIdent, specialnavigationsystemtimeslicepropertytype.Owns))
			if specialnavigationsystemtimeslicepropertytype.SpecialNavigationSystemTimeSlice != nil {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := specialnavigationsystemtimeslicepropertytype.SpecialNavigationSystemTimeSlice.GongGetIdentifier(stage.AixmfeaturesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.SpecialNavigationSystemTimeSlice = %s", specialnavigationsystemtimeslicepropertytypeIdent, targetIdent))
			}
		}
	}
	if stage.AixmfeaturesStage != nil {
		for _, specialnavigationsystemtimeslicetype := range __gong__sortStageSetInstances(stage.AixmfeaturesStage.SpecialNavigationSystemTimeSliceTypes, stage.AixmfeaturesStage.SpecialNavigationSystemTimeSliceType_stagedOrder) {
			if lastStageDecl != "AixmfeaturesStage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "AixmfeaturesStage"
			}
			specialnavigationsystemtimeslicetypeIdent := specialnavigationsystemtimeslicetype.GongGetIdentifier(stage.AixmfeaturesStage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&aixmfeatures.SpecialNavigationSystemTimeSliceType{Name: %s}).Stage(stage.AixmfeaturesStage)", specialnavigationsystemtimeslicetypeIdent, __gong__toRawStringLiteral(specialnavigationsystemtimeslicetype.Name)))
			if lastStageVal != "AixmfeaturesStage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "AixmfeaturesStage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", specialnavigationsystemtimeslicetypeIdent, __gong__toRawStringLiteral(specialnavigationsystemtimeslicetype.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Id = %s", specialnavigationsystemtimeslicetypeIdent, __gong__toRawStringLiteral(specialnavigationsystemtimeslicetype.Id)))
			values.WriteString(fmt.Sprintf("\n\t%s.Description = %s", specialnavigationsystemtimeslicetypeIdent, __gong__toRawStringLiteral(specialnavigationsystemtimeslicetype.Description)))
			values.WriteString(fmt.Sprintf("\n\t%s.Interpretation = %s", specialnavigationsystemtimeslicetypeIdent, __gong__toRawStringLiteral(specialnavigationsystemtimeslicetype.Interpretation)))
			values.WriteString(fmt.Sprintf("\n\t%s.SequenceNumber = %d", specialnavigationsystemtimeslicetypeIdent, specialnavigationsystemtimeslicetype.SequenceNumber))
			values.WriteString(fmt.Sprintf("\n\t%s.CorrectionNumber = %d", specialnavigationsystemtimeslicetypeIdent, specialnavigationsystemtimeslicetype.CorrectionNumber))
			if specialnavigationsystemtimeslicetype.Identifier != nil {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := specialnavigationsystemtimeslicetype.Identifier.GongGetIdentifier(stage.GmlStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Identifier = %s", specialnavigationsystemtimeslicetypeIdent, targetIdent))
			}
			for _, elem := range specialnavigationsystemtimeslicetype.NameXSD {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := elem.GongGetIdentifier(stage.GmlStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.NameXSD = append(%s.NameXSD, %s)", specialnavigationsystemtimeslicetypeIdent, specialnavigationsystemtimeslicetypeIdent, targetIdent))
			}
			if specialnavigationsystemtimeslicetype.ValidTime != nil {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := specialnavigationsystemtimeslicetype.ValidTime.GongGetIdentifier(stage.GmlStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.ValidTime = %s", specialnavigationsystemtimeslicetypeIdent, targetIdent))
			}
			if specialnavigationsystemtimeslicetype.TimeSliceMetadata != nil {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := specialnavigationsystemtimeslicetype.TimeSliceMetadata.GongGetIdentifier(stage.AixmabstractgmlobjecttypesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.TimeSliceMetadata = %s", specialnavigationsystemtimeslicetypeIdent, targetIdent))
			}
			if specialnavigationsystemtimeslicetype.FeatureLifetime != nil {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := specialnavigationsystemtimeslicetype.FeatureLifetime.GongGetIdentifier(stage.GmlStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.FeatureLifetime = %s", specialnavigationsystemtimeslicetypeIdent, targetIdent))
			}
			if specialnavigationsystemtimeslicetype.Type_XSD != nil {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := specialnavigationsystemtimeslicetype.Type_XSD.GongGetIdentifier(stage.AixmdatatypesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Type_XSD = %s", specialnavigationsystemtimeslicetypeIdent, targetIdent))
			}
			if specialnavigationsystemtimeslicetype.Designator != nil {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := specialnavigationsystemtimeslicetype.Designator.GongGetIdentifier(stage.AixmdatatypesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Designator = %s", specialnavigationsystemtimeslicetypeIdent, targetIdent))
			}
			if specialnavigationsystemtimeslicetype.ResponsibleOrganisation != nil {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := specialnavigationsystemtimeslicetype.ResponsibleOrganisation.GongGetIdentifier(stage.AixmfeaturesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.ResponsibleOrganisation = %s", specialnavigationsystemtimeslicetypeIdent, targetIdent))
			}
			for _, elem := range specialnavigationsystemtimeslicetype.Annotation {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := elem.GongGetIdentifier(stage.AixmfeaturesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Annotation = append(%s.Annotation, %s)", specialnavigationsystemtimeslicetypeIdent, specialnavigationsystemtimeslicetypeIdent, targetIdent))
			}
			for _, elem := range specialnavigationsystemtimeslicetype.Extension {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := elem.GongGetIdentifier(stage.AixmfeaturesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Extension = append(%s.Extension, %s)", specialnavigationsystemtimeslicetypeIdent, specialnavigationsystemtimeslicetypeIdent, targetIdent))
			}
		}
	}
	if stage.AixmfeaturesStage != nil {
		for _, specialnavigationsystemtype := range __gong__sortStageSetInstances(stage.AixmfeaturesStage.SpecialNavigationSystemTypes, stage.AixmfeaturesStage.SpecialNavigationSystemType_stagedOrder) {
			if lastStageDecl != "AixmfeaturesStage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "AixmfeaturesStage"
			}
			specialnavigationsystemtypeIdent := specialnavigationsystemtype.GongGetIdentifier(stage.AixmfeaturesStage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&aixmfeatures.SpecialNavigationSystemType{Name: %s}).Stage(stage.AixmfeaturesStage)", specialnavigationsystemtypeIdent, __gong__toRawStringLiteral(specialnavigationsystemtype.Name)))
			if lastStageVal != "AixmfeaturesStage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "AixmfeaturesStage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", specialnavigationsystemtypeIdent, __gong__toRawStringLiteral(specialnavigationsystemtype.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Id = %s", specialnavigationsystemtypeIdent, __gong__toRawStringLiteral(specialnavigationsystemtype.Id)))
			values.WriteString(fmt.Sprintf("\n\t%s.Description = %s", specialnavigationsystemtypeIdent, __gong__toRawStringLiteral(specialnavigationsystemtype.Description)))
			if specialnavigationsystemtype.Identifier != nil {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := specialnavigationsystemtype.Identifier.GongGetIdentifier(stage.GmlStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Identifier = %s", specialnavigationsystemtypeIdent, targetIdent))
			}
			for _, elem := range specialnavigationsystemtype.NameXSD {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := elem.GongGetIdentifier(stage.GmlStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.NameXSD = append(%s.NameXSD, %s)", specialnavigationsystemtypeIdent, specialnavigationsystemtypeIdent, targetIdent))
			}
			if specialnavigationsystemtype.BoundedBy != nil {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := specialnavigationsystemtype.BoundedBy.GongGetIdentifier(stage.GmlStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.BoundedBy = %s", specialnavigationsystemtypeIdent, targetIdent))
			}
			if specialnavigationsystemtype.FeatureMetadata != nil {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := specialnavigationsystemtype.FeatureMetadata.GongGetIdentifier(stage.AixmabstractgmlobjecttypesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.FeatureMetadata = %s", specialnavigationsystemtypeIdent, targetIdent))
			}
			for _, elem := range specialnavigationsystemtype.TimeSlice {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := elem.GongGetIdentifier(stage.AixmfeaturesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.TimeSlice = append(%s.TimeSlice, %s)", specialnavigationsystemtypeIdent, specialnavigationsystemtypeIdent, targetIdent))
			}
		}
	}
	if stage.AixmfeaturesStage != nil {
		for _, surveillancegroundstationpropertytype := range __gong__sortStageSetInstances(stage.AixmfeaturesStage.SurveillanceGroundStationPropertyTypes, stage.AixmfeaturesStage.SurveillanceGroundStationPropertyType_stagedOrder) {
			if lastStageDecl != "AixmfeaturesStage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "AixmfeaturesStage"
			}
			surveillancegroundstationpropertytypeIdent := surveillancegroundstationpropertytype.GongGetIdentifier(stage.AixmfeaturesStage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&aixmfeatures.SurveillanceGroundStationPropertyType{Name: %s}).Stage(stage.AixmfeaturesStage)", surveillancegroundstationpropertytypeIdent, __gong__toRawStringLiteral(surveillancegroundstationpropertytype.Name)))
			if lastStageVal != "AixmfeaturesStage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "AixmfeaturesStage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", surveillancegroundstationpropertytypeIdent, __gong__toRawStringLiteral(surveillancegroundstationpropertytype.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.NilReason = %s", surveillancegroundstationpropertytypeIdent, __gong__toRawStringLiteral(surveillancegroundstationpropertytype.NilReason)))
			if surveillancegroundstationpropertytype.SurveillanceGroundStation != nil {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := surveillancegroundstationpropertytype.SurveillanceGroundStation.GongGetIdentifier(stage.AixmfeaturesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.SurveillanceGroundStation = %s", surveillancegroundstationpropertytypeIdent, targetIdent))
			}
		}
	}
	if stage.AixmfeaturesStage != nil {
		for _, surveillancegroundstationtype := range __gong__sortStageSetInstances(stage.AixmfeaturesStage.SurveillanceGroundStationTypes, stage.AixmfeaturesStage.SurveillanceGroundStationType_stagedOrder) {
			if lastStageDecl != "AixmfeaturesStage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "AixmfeaturesStage"
			}
			surveillancegroundstationtypeIdent := surveillancegroundstationtype.GongGetIdentifier(stage.AixmfeaturesStage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&aixmfeatures.SurveillanceGroundStationType{Name: %s}).Stage(stage.AixmfeaturesStage)", surveillancegroundstationtypeIdent, __gong__toRawStringLiteral(surveillancegroundstationtype.Name)))
			if lastStageVal != "AixmfeaturesStage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "AixmfeaturesStage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", surveillancegroundstationtypeIdent, __gong__toRawStringLiteral(surveillancegroundstationtype.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Id = %s", surveillancegroundstationtypeIdent, __gong__toRawStringLiteral(surveillancegroundstationtype.Id)))
			values.WriteString(fmt.Sprintf("\n\t%s.Description = %s", surveillancegroundstationtypeIdent, __gong__toRawStringLiteral(surveillancegroundstationtype.Description)))
			if surveillancegroundstationtype.Identifier != nil {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := surveillancegroundstationtype.Identifier.GongGetIdentifier(stage.GmlStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Identifier = %s", surveillancegroundstationtypeIdent, targetIdent))
			}
			for _, elem := range surveillancegroundstationtype.NameXSD {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := elem.GongGetIdentifier(stage.GmlStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.NameXSD = append(%s.NameXSD, %s)", surveillancegroundstationtypeIdent, surveillancegroundstationtypeIdent, targetIdent))
			}
			if surveillancegroundstationtype.VideoMap != nil {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := surveillancegroundstationtype.VideoMap.GongGetIdentifier(stage.AixmdatatypesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.VideoMap = %s", surveillancegroundstationtypeIdent, targetIdent))
			}
			for _, elem := range surveillancegroundstationtype.Annotation {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := elem.GongGetIdentifier(stage.AixmfeaturesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Annotation = append(%s.Annotation, %s)", surveillancegroundstationtypeIdent, surveillancegroundstationtypeIdent, targetIdent))
			}
			if surveillancegroundstationtype.TheUnit != nil {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := surveillancegroundstationtype.TheUnit.GongGetIdentifier(stage.AixmfeaturesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.TheUnit = %s", surveillancegroundstationtypeIdent, targetIdent))
			}
			for _, elem := range surveillancegroundstationtype.Extension {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := elem.GongGetIdentifier(stage.AixmfeaturesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Extension = append(%s.Extension, %s)", surveillancegroundstationtypeIdent, surveillancegroundstationtypeIdent, targetIdent))
			}
		}
	}
	if stage.AixmfeaturesStage != nil {
		for _, tacantimeslicepropertytype := range __gong__sortStageSetInstances(stage.AixmfeaturesStage.TACANTimeSlicePropertyTypes, stage.AixmfeaturesStage.TACANTimeSlicePropertyType_stagedOrder) {
			if lastStageDecl != "AixmfeaturesStage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "AixmfeaturesStage"
			}
			tacantimeslicepropertytypeIdent := tacantimeslicepropertytype.GongGetIdentifier(stage.AixmfeaturesStage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&aixmfeatures.TACANTimeSlicePropertyType{Name: %s}).Stage(stage.AixmfeaturesStage)", tacantimeslicepropertytypeIdent, __gong__toRawStringLiteral(tacantimeslicepropertytype.Name)))
			if lastStageVal != "AixmfeaturesStage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "AixmfeaturesStage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", tacantimeslicepropertytypeIdent, __gong__toRawStringLiteral(tacantimeslicepropertytype.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Owns = %t", tacantimeslicepropertytypeIdent, tacantimeslicepropertytype.Owns))
			if tacantimeslicepropertytype.TACANTimeSlice != nil {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := tacantimeslicepropertytype.TACANTimeSlice.GongGetIdentifier(stage.AixmfeaturesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.TACANTimeSlice = %s", tacantimeslicepropertytypeIdent, targetIdent))
			}
		}
	}
	if stage.AixmfeaturesStage != nil {
		for _, tacantimeslicetype := range __gong__sortStageSetInstances(stage.AixmfeaturesStage.TACANTimeSliceTypes, stage.AixmfeaturesStage.TACANTimeSliceType_stagedOrder) {
			if lastStageDecl != "AixmfeaturesStage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "AixmfeaturesStage"
			}
			tacantimeslicetypeIdent := tacantimeslicetype.GongGetIdentifier(stage.AixmfeaturesStage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&aixmfeatures.TACANTimeSliceType{Name: %s}).Stage(stage.AixmfeaturesStage)", tacantimeslicetypeIdent, __gong__toRawStringLiteral(tacantimeslicetype.Name)))
			if lastStageVal != "AixmfeaturesStage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "AixmfeaturesStage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", tacantimeslicetypeIdent, __gong__toRawStringLiteral(tacantimeslicetype.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Id = %s", tacantimeslicetypeIdent, __gong__toRawStringLiteral(tacantimeslicetype.Id)))
			values.WriteString(fmt.Sprintf("\n\t%s.Description = %s", tacantimeslicetypeIdent, __gong__toRawStringLiteral(tacantimeslicetype.Description)))
			values.WriteString(fmt.Sprintf("\n\t%s.Interpretation = %s", tacantimeslicetypeIdent, __gong__toRawStringLiteral(tacantimeslicetype.Interpretation)))
			values.WriteString(fmt.Sprintf("\n\t%s.SequenceNumber = %d", tacantimeslicetypeIdent, tacantimeslicetype.SequenceNumber))
			values.WriteString(fmt.Sprintf("\n\t%s.CorrectionNumber = %d", tacantimeslicetypeIdent, tacantimeslicetype.CorrectionNumber))
			if tacantimeslicetype.Identifier != nil {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := tacantimeslicetype.Identifier.GongGetIdentifier(stage.GmlStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Identifier = %s", tacantimeslicetypeIdent, targetIdent))
			}
			for _, elem := range tacantimeslicetype.NameXSD {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := elem.GongGetIdentifier(stage.GmlStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.NameXSD = append(%s.NameXSD, %s)", tacantimeslicetypeIdent, tacantimeslicetypeIdent, targetIdent))
			}
			if tacantimeslicetype.ValidTime != nil {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := tacantimeslicetype.ValidTime.GongGetIdentifier(stage.GmlStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.ValidTime = %s", tacantimeslicetypeIdent, targetIdent))
			}
			if tacantimeslicetype.TimeSliceMetadata != nil {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := tacantimeslicetype.TimeSliceMetadata.GongGetIdentifier(stage.AixmabstractgmlobjecttypesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.TimeSliceMetadata = %s", tacantimeslicetypeIdent, targetIdent))
			}
			if tacantimeslicetype.FeatureLifetime != nil {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := tacantimeslicetype.FeatureLifetime.GongGetIdentifier(stage.GmlStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.FeatureLifetime = %s", tacantimeslicetypeIdent, targetIdent))
			}
			if tacantimeslicetype.Designator != nil {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := tacantimeslicetype.Designator.GongGetIdentifier(stage.AixmdatatypesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Designator = %s", tacantimeslicetypeIdent, targetIdent))
			}
			if tacantimeslicetype.EmissionClass != nil {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := tacantimeslicetype.EmissionClass.GongGetIdentifier(stage.AixmdatatypesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.EmissionClass = %s", tacantimeslicetypeIdent, targetIdent))
			}
			if tacantimeslicetype.Mobile != nil {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := tacantimeslicetype.Mobile.GongGetIdentifier(stage.AixmdatatypesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Mobile = %s", tacantimeslicetypeIdent, targetIdent))
			}
			if tacantimeslicetype.MagneticVariation != nil {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := tacantimeslicetype.MagneticVariation.GongGetIdentifier(stage.AixmdatatypesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.MagneticVariation = %s", tacantimeslicetypeIdent, targetIdent))
			}
			if tacantimeslicetype.MagneticVariationAccuracy != nil {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := tacantimeslicetype.MagneticVariationAccuracy.GongGetIdentifier(stage.AixmdatatypesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.MagneticVariationAccuracy = %s", tacantimeslicetypeIdent, targetIdent))
			}
			if tacantimeslicetype.DateMagneticVariation != nil {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := tacantimeslicetype.DateMagneticVariation.GongGetIdentifier(stage.AixmdatatypesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.DateMagneticVariation = %s", tacantimeslicetypeIdent, targetIdent))
			}
			if tacantimeslicetype.FlightChecked != nil {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := tacantimeslicetype.FlightChecked.GongGetIdentifier(stage.AixmdatatypesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.FlightChecked = %s", tacantimeslicetypeIdent, targetIdent))
			}
			if tacantimeslicetype.Location != nil {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := tacantimeslicetype.Location.GongGetIdentifier(stage.AixmfeaturesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Location = %s", tacantimeslicetypeIdent, targetIdent))
			}
			for _, elem := range tacantimeslicetype.Authority {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := elem.GongGetIdentifier(stage.AixmfeaturesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Authority = append(%s.Authority, %s)", tacantimeslicetypeIdent, tacantimeslicetypeIdent, targetIdent))
			}
			for _, elem := range tacantimeslicetype.Monitoring {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := elem.GongGetIdentifier(stage.AixmfeaturesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Monitoring = append(%s.Monitoring, %s)", tacantimeslicetypeIdent, tacantimeslicetypeIdent, targetIdent))
			}
			for _, elem := range tacantimeslicetype.Availability {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := elem.GongGetIdentifier(stage.AixmfeaturesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Availability = append(%s.Availability, %s)", tacantimeslicetypeIdent, tacantimeslicetypeIdent, targetIdent))
			}
			for _, elem := range tacantimeslicetype.Annotation {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := elem.GongGetIdentifier(stage.AixmfeaturesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Annotation = append(%s.Annotation, %s)", tacantimeslicetypeIdent, tacantimeslicetypeIdent, targetIdent))
			}
			if tacantimeslicetype.Channel != nil {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := tacantimeslicetype.Channel.GongGetIdentifier(stage.AixmdatatypesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Channel = %s", tacantimeslicetypeIdent, targetIdent))
			}
			if tacantimeslicetype.Declination != nil {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := tacantimeslicetype.Declination.GongGetIdentifier(stage.AixmdatatypesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Declination = %s", tacantimeslicetypeIdent, targetIdent))
			}
			for _, elem := range tacantimeslicetype.Extension {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := elem.GongGetIdentifier(stage.AixmfeaturesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Extension = append(%s.Extension, %s)", tacantimeslicetypeIdent, tacantimeslicetypeIdent, targetIdent))
			}
		}
	}
	if stage.AixmfeaturesStage != nil {
		for _, tacantype := range __gong__sortStageSetInstances(stage.AixmfeaturesStage.TACANTypes, stage.AixmfeaturesStage.TACANType_stagedOrder) {
			if lastStageDecl != "AixmfeaturesStage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "AixmfeaturesStage"
			}
			tacantypeIdent := tacantype.GongGetIdentifier(stage.AixmfeaturesStage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&aixmfeatures.TACANType{Name: %s}).Stage(stage.AixmfeaturesStage)", tacantypeIdent, __gong__toRawStringLiteral(tacantype.Name)))
			if lastStageVal != "AixmfeaturesStage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "AixmfeaturesStage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", tacantypeIdent, __gong__toRawStringLiteral(tacantype.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Id = %s", tacantypeIdent, __gong__toRawStringLiteral(tacantype.Id)))
			values.WriteString(fmt.Sprintf("\n\t%s.Description = %s", tacantypeIdent, __gong__toRawStringLiteral(tacantype.Description)))
			if tacantype.Identifier != nil {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := tacantype.Identifier.GongGetIdentifier(stage.GmlStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Identifier = %s", tacantypeIdent, targetIdent))
			}
			for _, elem := range tacantype.NameXSD {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := elem.GongGetIdentifier(stage.GmlStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.NameXSD = append(%s.NameXSD, %s)", tacantypeIdent, tacantypeIdent, targetIdent))
			}
			if tacantype.BoundedBy != nil {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := tacantype.BoundedBy.GongGetIdentifier(stage.GmlStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.BoundedBy = %s", tacantypeIdent, targetIdent))
			}
			if tacantype.FeatureMetadata != nil {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := tacantype.FeatureMetadata.GongGetIdentifier(stage.AixmabstractgmlobjecttypesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.FeatureMetadata = %s", tacantypeIdent, targetIdent))
			}
			for _, elem := range tacantype.TimeSlice {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := elem.GongGetIdentifier(stage.AixmfeaturesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.TimeSlice = append(%s.TimeSlice, %s)", tacantypeIdent, tacantypeIdent, targetIdent))
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
	if stage.AixmfeaturesStage != nil {
		for _, vortimeslicepropertytype := range __gong__sortStageSetInstances(stage.AixmfeaturesStage.VORTimeSlicePropertyTypes, stage.AixmfeaturesStage.VORTimeSlicePropertyType_stagedOrder) {
			if lastStageDecl != "AixmfeaturesStage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "AixmfeaturesStage"
			}
			vortimeslicepropertytypeIdent := vortimeslicepropertytype.GongGetIdentifier(stage.AixmfeaturesStage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&aixmfeatures.VORTimeSlicePropertyType{Name: %s}).Stage(stage.AixmfeaturesStage)", vortimeslicepropertytypeIdent, __gong__toRawStringLiteral(vortimeslicepropertytype.Name)))
			if lastStageVal != "AixmfeaturesStage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "AixmfeaturesStage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", vortimeslicepropertytypeIdent, __gong__toRawStringLiteral(vortimeslicepropertytype.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Owns = %t", vortimeslicepropertytypeIdent, vortimeslicepropertytype.Owns))
			if vortimeslicepropertytype.VORTimeSlice != nil {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := vortimeslicepropertytype.VORTimeSlice.GongGetIdentifier(stage.AixmfeaturesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.VORTimeSlice = %s", vortimeslicepropertytypeIdent, targetIdent))
			}
		}
	}
	if stage.AixmfeaturesStage != nil {
		for _, vortimeslicetype := range __gong__sortStageSetInstances(stage.AixmfeaturesStage.VORTimeSliceTypes, stage.AixmfeaturesStage.VORTimeSliceType_stagedOrder) {
			if lastStageDecl != "AixmfeaturesStage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "AixmfeaturesStage"
			}
			vortimeslicetypeIdent := vortimeslicetype.GongGetIdentifier(stage.AixmfeaturesStage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&aixmfeatures.VORTimeSliceType{Name: %s}).Stage(stage.AixmfeaturesStage)", vortimeslicetypeIdent, __gong__toRawStringLiteral(vortimeslicetype.Name)))
			if lastStageVal != "AixmfeaturesStage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "AixmfeaturesStage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", vortimeslicetypeIdent, __gong__toRawStringLiteral(vortimeslicetype.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Id = %s", vortimeslicetypeIdent, __gong__toRawStringLiteral(vortimeslicetype.Id)))
			values.WriteString(fmt.Sprintf("\n\t%s.Description = %s", vortimeslicetypeIdent, __gong__toRawStringLiteral(vortimeslicetype.Description)))
			values.WriteString(fmt.Sprintf("\n\t%s.Interpretation = %s", vortimeslicetypeIdent, __gong__toRawStringLiteral(vortimeslicetype.Interpretation)))
			values.WriteString(fmt.Sprintf("\n\t%s.SequenceNumber = %d", vortimeslicetypeIdent, vortimeslicetype.SequenceNumber))
			values.WriteString(fmt.Sprintf("\n\t%s.CorrectionNumber = %d", vortimeslicetypeIdent, vortimeslicetype.CorrectionNumber))
			if vortimeslicetype.Identifier != nil {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := vortimeslicetype.Identifier.GongGetIdentifier(stage.GmlStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Identifier = %s", vortimeslicetypeIdent, targetIdent))
			}
			for _, elem := range vortimeslicetype.NameXSD {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := elem.GongGetIdentifier(stage.GmlStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.NameXSD = append(%s.NameXSD, %s)", vortimeslicetypeIdent, vortimeslicetypeIdent, targetIdent))
			}
			if vortimeslicetype.ValidTime != nil {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := vortimeslicetype.ValidTime.GongGetIdentifier(stage.GmlStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.ValidTime = %s", vortimeslicetypeIdent, targetIdent))
			}
			if vortimeslicetype.TimeSliceMetadata != nil {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := vortimeslicetype.TimeSliceMetadata.GongGetIdentifier(stage.AixmabstractgmlobjecttypesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.TimeSliceMetadata = %s", vortimeslicetypeIdent, targetIdent))
			}
			if vortimeslicetype.FeatureLifetime != nil {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := vortimeslicetype.FeatureLifetime.GongGetIdentifier(stage.GmlStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.FeatureLifetime = %s", vortimeslicetypeIdent, targetIdent))
			}
			if vortimeslicetype.Designator != nil {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := vortimeslicetype.Designator.GongGetIdentifier(stage.AixmdatatypesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Designator = %s", vortimeslicetypeIdent, targetIdent))
			}
			if vortimeslicetype.EmissionClass != nil {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := vortimeslicetype.EmissionClass.GongGetIdentifier(stage.AixmdatatypesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.EmissionClass = %s", vortimeslicetypeIdent, targetIdent))
			}
			if vortimeslicetype.Mobile != nil {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := vortimeslicetype.Mobile.GongGetIdentifier(stage.AixmdatatypesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Mobile = %s", vortimeslicetypeIdent, targetIdent))
			}
			if vortimeslicetype.MagneticVariation != nil {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := vortimeslicetype.MagneticVariation.GongGetIdentifier(stage.AixmdatatypesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.MagneticVariation = %s", vortimeslicetypeIdent, targetIdent))
			}
			if vortimeslicetype.MagneticVariationAccuracy != nil {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := vortimeslicetype.MagneticVariationAccuracy.GongGetIdentifier(stage.AixmdatatypesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.MagneticVariationAccuracy = %s", vortimeslicetypeIdent, targetIdent))
			}
			if vortimeslicetype.DateMagneticVariation != nil {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := vortimeslicetype.DateMagneticVariation.GongGetIdentifier(stage.AixmdatatypesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.DateMagneticVariation = %s", vortimeslicetypeIdent, targetIdent))
			}
			if vortimeslicetype.FlightChecked != nil {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := vortimeslicetype.FlightChecked.GongGetIdentifier(stage.AixmdatatypesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.FlightChecked = %s", vortimeslicetypeIdent, targetIdent))
			}
			if vortimeslicetype.Location != nil {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := vortimeslicetype.Location.GongGetIdentifier(stage.AixmfeaturesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Location = %s", vortimeslicetypeIdent, targetIdent))
			}
			for _, elem := range vortimeslicetype.Authority {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := elem.GongGetIdentifier(stage.AixmfeaturesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Authority = append(%s.Authority, %s)", vortimeslicetypeIdent, vortimeslicetypeIdent, targetIdent))
			}
			for _, elem := range vortimeslicetype.Monitoring {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := elem.GongGetIdentifier(stage.AixmfeaturesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Monitoring = append(%s.Monitoring, %s)", vortimeslicetypeIdent, vortimeslicetypeIdent, targetIdent))
			}
			for _, elem := range vortimeslicetype.Availability {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := elem.GongGetIdentifier(stage.AixmfeaturesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Availability = append(%s.Availability, %s)", vortimeslicetypeIdent, vortimeslicetypeIdent, targetIdent))
			}
			for _, elem := range vortimeslicetype.Annotation {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := elem.GongGetIdentifier(stage.AixmfeaturesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Annotation = append(%s.Annotation, %s)", vortimeslicetypeIdent, vortimeslicetypeIdent, targetIdent))
			}
			if vortimeslicetype.Type_XSD != nil {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := vortimeslicetype.Type_XSD.GongGetIdentifier(stage.AixmdatatypesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Type_XSD = %s", vortimeslicetypeIdent, targetIdent))
			}
			if vortimeslicetype.Frequency != nil {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := vortimeslicetype.Frequency.GongGetIdentifier(stage.AixmdatatypesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Frequency = %s", vortimeslicetypeIdent, targetIdent))
			}
			if vortimeslicetype.ZeroBearingDirection != nil {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := vortimeslicetype.ZeroBearingDirection.GongGetIdentifier(stage.AixmdatatypesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.ZeroBearingDirection = %s", vortimeslicetypeIdent, targetIdent))
			}
			if vortimeslicetype.Declination != nil {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := vortimeslicetype.Declination.GongGetIdentifier(stage.AixmdatatypesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Declination = %s", vortimeslicetypeIdent, targetIdent))
			}
			for _, elem := range vortimeslicetype.Extension {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := elem.GongGetIdentifier(stage.AixmfeaturesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Extension = append(%s.Extension, %s)", vortimeslicetypeIdent, vortimeslicetypeIdent, targetIdent))
			}
		}
	}
	if stage.AixmfeaturesStage != nil {
		for _, vortype := range __gong__sortStageSetInstances(stage.AixmfeaturesStage.VORTypes, stage.AixmfeaturesStage.VORType_stagedOrder) {
			if lastStageDecl != "AixmfeaturesStage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "AixmfeaturesStage"
			}
			vortypeIdent := vortype.GongGetIdentifier(stage.AixmfeaturesStage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&aixmfeatures.VORType{Name: %s}).Stage(stage.AixmfeaturesStage)", vortypeIdent, __gong__toRawStringLiteral(vortype.Name)))
			if lastStageVal != "AixmfeaturesStage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "AixmfeaturesStage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", vortypeIdent, __gong__toRawStringLiteral(vortype.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Id = %s", vortypeIdent, __gong__toRawStringLiteral(vortype.Id)))
			values.WriteString(fmt.Sprintf("\n\t%s.Description = %s", vortypeIdent, __gong__toRawStringLiteral(vortype.Description)))
			if vortype.Identifier != nil {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := vortype.Identifier.GongGetIdentifier(stage.GmlStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Identifier = %s", vortypeIdent, targetIdent))
			}
			for _, elem := range vortype.NameXSD {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := elem.GongGetIdentifier(stage.GmlStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.NameXSD = append(%s.NameXSD, %s)", vortypeIdent, vortypeIdent, targetIdent))
			}
			if vortype.BoundedBy != nil {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := vortype.BoundedBy.GongGetIdentifier(stage.GmlStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.BoundedBy = %s", vortypeIdent, targetIdent))
			}
			if vortype.FeatureMetadata != nil {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := vortype.FeatureMetadata.GongGetIdentifier(stage.AixmabstractgmlobjecttypesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.FeatureMetadata = %s", vortypeIdent, targetIdent))
			}
			for _, elem := range vortype.TimeSlice {
				if lastStagePtr != "AixmfeaturesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AixmfeaturesStage"
				}
				targetIdent := elem.GongGetIdentifier(stage.AixmfeaturesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.TimeSlice = append(%s.TimeSlice, %s)", vortypeIdent, vortypeIdent, targetIdent))
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
			if basicmessagememberaixmpropertytype.Navaid != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := basicmessagememberaixmpropertytype.Navaid.GongGetIdentifier(stage.AixmfeaturesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Navaid = %s", basicmessagememberaixmpropertytypeIdent, targetIdent))
			}
			if basicmessagememberaixmpropertytype.VOR != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := basicmessagememberaixmpropertytype.VOR.GongGetIdentifier(stage.AixmfeaturesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.VOR = %s", basicmessagememberaixmpropertytypeIdent, targetIdent))
			}
			if basicmessagememberaixmpropertytype.DME != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := basicmessagememberaixmpropertytype.DME.GongGetIdentifier(stage.AixmfeaturesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.DME = %s", basicmessagememberaixmpropertytypeIdent, targetIdent))
			}
			if basicmessagememberaixmpropertytype.TACAN != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := basicmessagememberaixmpropertytype.TACAN.GongGetIdentifier(stage.AixmfeaturesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.TACAN = %s", basicmessagememberaixmpropertytypeIdent, targetIdent))
			}
			if basicmessagememberaixmpropertytype.NDB != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := basicmessagememberaixmpropertytype.NDB.GongGetIdentifier(stage.AixmfeaturesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.NDB = %s", basicmessagememberaixmpropertytypeIdent, targetIdent))
			}
			if basicmessagememberaixmpropertytype.Localizer != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := basicmessagememberaixmpropertytype.Localizer.GongGetIdentifier(stage.AixmfeaturesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Localizer = %s", basicmessagememberaixmpropertytypeIdent, targetIdent))
			}
			if basicmessagememberaixmpropertytype.Glidepath != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := basicmessagememberaixmpropertytype.Glidepath.GongGetIdentifier(stage.AixmfeaturesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Glidepath = %s", basicmessagememberaixmpropertytypeIdent, targetIdent))
			}
			if basicmessagememberaixmpropertytype.MarkerBeacon != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := basicmessagememberaixmpropertytype.MarkerBeacon.GongGetIdentifier(stage.AixmfeaturesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.MarkerBeacon = %s", basicmessagememberaixmpropertytypeIdent, targetIdent))
			}
			if basicmessagememberaixmpropertytype.SDF != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := basicmessagememberaixmpropertytype.SDF.GongGetIdentifier(stage.AixmfeaturesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.SDF = %s", basicmessagememberaixmpropertytypeIdent, targetIdent))
			}
			if basicmessagememberaixmpropertytype.DirectionFinder != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := basicmessagememberaixmpropertytype.DirectionFinder.GongGetIdentifier(stage.AixmfeaturesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.DirectionFinder = %s", basicmessagememberaixmpropertytypeIdent, targetIdent))
			}
			if basicmessagememberaixmpropertytype.Azimuth != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := basicmessagememberaixmpropertytype.Azimuth.GongGetIdentifier(stage.AixmfeaturesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Azimuth = %s", basicmessagememberaixmpropertytypeIdent, targetIdent))
			}
			if basicmessagememberaixmpropertytype.Elevation != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := basicmessagememberaixmpropertytype.Elevation.GongGetIdentifier(stage.AixmfeaturesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Elevation = %s", basicmessagememberaixmpropertytypeIdent, targetIdent))
			}
			if basicmessagememberaixmpropertytype.CheckpointVOR != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := basicmessagememberaixmpropertytype.CheckpointVOR.GongGetIdentifier(stage.AixmfeaturesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.CheckpointVOR = %s", basicmessagememberaixmpropertytypeIdent, targetIdent))
			}
			if basicmessagememberaixmpropertytype.CheckpointINS != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := basicmessagememberaixmpropertytype.CheckpointINS.GongGetIdentifier(stage.AixmfeaturesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.CheckpointINS = %s", basicmessagememberaixmpropertytypeIdent, targetIdent))
			}
			if basicmessagememberaixmpropertytype.SpecialNavigationSystem != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := basicmessagememberaixmpropertytype.SpecialNavigationSystem.GongGetIdentifier(stage.AixmfeaturesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.SpecialNavigationSystem = %s", basicmessagememberaixmpropertytypeIdent, targetIdent))
			}
			if basicmessagememberaixmpropertytype.SpecialNavigationStation != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := basicmessagememberaixmpropertytype.SpecialNavigationStation.GongGetIdentifier(stage.AixmfeaturesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.SpecialNavigationStation = %s", basicmessagememberaixmpropertytypeIdent, targetIdent))
			}
			if basicmessagememberaixmpropertytype.DesignatedPoint != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := basicmessagememberaixmpropertytype.DesignatedPoint.GongGetIdentifier(stage.AixmfeaturesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.DesignatedPoint = %s", basicmessagememberaixmpropertytypeIdent, targetIdent))
			}
			if basicmessagememberaixmpropertytype.SignificantPointInAirspace != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := basicmessagememberaixmpropertytype.SignificantPointInAirspace.GongGetIdentifier(stage.AixmfeaturesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.SignificantPointInAirspace = %s", basicmessagememberaixmpropertytypeIdent, targetIdent))
			}
			if basicmessagememberaixmpropertytype.AngleIndication != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := basicmessagememberaixmpropertytype.AngleIndication.GongGetIdentifier(stage.AixmfeaturesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.AngleIndication = %s", basicmessagememberaixmpropertytypeIdent, targetIdent))
			}
			if basicmessagememberaixmpropertytype.DistanceIndication != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := basicmessagememberaixmpropertytype.DistanceIndication.GongGetIdentifier(stage.AixmfeaturesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.DistanceIndication = %s", basicmessagememberaixmpropertytypeIdent, targetIdent))
			}
			if basicmessagememberaixmpropertytype.RadarSystem != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := basicmessagememberaixmpropertytype.RadarSystem.GongGetIdentifier(stage.AixmfeaturesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.RadarSystem = %s", basicmessagememberaixmpropertytypeIdent, targetIdent))
			}
			if basicmessagememberaixmpropertytype.PrimarySurveillanceRadar != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := basicmessagememberaixmpropertytype.PrimarySurveillanceRadar.GongGetIdentifier(stage.AixmfeaturesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.PrimarySurveillanceRadar = %s", basicmessagememberaixmpropertytypeIdent, targetIdent))
			}
			if basicmessagememberaixmpropertytype.SecondarySurveillanceRadar != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := basicmessagememberaixmpropertytype.SecondarySurveillanceRadar.GongGetIdentifier(stage.AixmfeaturesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.SecondarySurveillanceRadar = %s", basicmessagememberaixmpropertytypeIdent, targetIdent))
			}
			if basicmessagememberaixmpropertytype.PrecisionApproachRadar != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := basicmessagememberaixmpropertytype.PrecisionApproachRadar.GongGetIdentifier(stage.AixmfeaturesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.PrecisionApproachRadar = %s", basicmessagememberaixmpropertytypeIdent, targetIdent))
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
go run github.com/fullstack-lang/gong/app/xsd/tests/aixm/navaid/go/models/probe/cmd/probe@%s stage.go
*/
package %s

import (
	"slices"
	"time"

	"github.com/fullstack-lang/gong/app/xsd/tests/aixm/navaid/go/models"
	"github.com/fullstack-lang/gong/app/xsd/tests/aixm/navaid/go/models/aixmfeatures"
	"github.com/fullstack-lang/gong/app/xsd/tests/aixm/navaid/go/models/aixmdatatypes"
	"github.com/fullstack-lang/gong/app/xsd/tests/aixm/navaid/go/models/aixmabstractgmlobjecttypes"
	"github.com/fullstack-lang/gong/app/xsd/tests/aixm/navaid/go/models/gml"
	"github.com/fullstack-lang/gong/app/xsd/tests/aixm/navaid/go/models/xlink"%s
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

	if modelsPackageName != "" && modelsPackageName != "github.com/fullstack-lang/gong/app/xsd/tests/aixm/navaid/go/models" {
		res = strings.ReplaceAll(res, "\"github.com/fullstack-lang/gong/app/xsd/tests/aixm/navaid/go/models\"", fmt.Sprintf("%q", modelsPackageName))
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
