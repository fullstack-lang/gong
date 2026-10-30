// generated code - do not edit
package gml

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
	{{Identifier}} := (&gml.{{GeneratedStructName}}{Name: {{GeneratedFieldNameValue}}}).Stage(stage)`

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
			const fnSig = "func _(stage *gml.Stage) {"
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
			strings.HasSuffix(trimmedContentBeforeBrace, "func _(stage *gml.Stage) {") ||
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
	modelsPackageName := "github.com/fullstack-lang/gong/app/xsd/tests/aixm/service/go/models/gml"
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
		for _, abstractfeaturemembertype := range __gong__sortStageSetInstances(stage.AbstractFeatureMemberTypes, stage.AbstractFeatureMemberType_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			abstractfeaturemembertypeIdent := abstractfeaturemembertype.GongGetIdentifier(stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&gml.AbstractFeatureMemberType{Name: %s}).Stage(stage)", abstractfeaturemembertypeIdent, __gong__toRawStringLiteral(abstractfeaturemembertype.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", abstractfeaturemembertypeIdent, __gong__toRawStringLiteral(abstractfeaturemembertype.Name)))
		}
	}
	if stage != nil {
		for _, abstractfeaturetype := range __gong__sortStageSetInstances(stage.AbstractFeatureTypes, stage.AbstractFeatureType_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			abstractfeaturetypeIdent := abstractfeaturetype.GongGetIdentifier(stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&gml.AbstractFeatureType{Name: %s}).Stage(stage)", abstractfeaturetypeIdent, __gong__toRawStringLiteral(abstractfeaturetype.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", abstractfeaturetypeIdent, __gong__toRawStringLiteral(abstractfeaturetype.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Id = %s", abstractfeaturetypeIdent, __gong__toRawStringLiteral(abstractfeaturetype.Id)))
			values.WriteString(fmt.Sprintf("\n\t%s.Description = %s", abstractfeaturetypeIdent, __gong__toRawStringLiteral(abstractfeaturetype.Description)))
			if abstractfeaturetype.Identifier != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := abstractfeaturetype.Identifier.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Identifier = %s", abstractfeaturetypeIdent, targetIdent))
			}
			for _, elem := range abstractfeaturetype.NameXSD {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := elem.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.NameXSD = append(%s.NameXSD, %s)", abstractfeaturetypeIdent, abstractfeaturetypeIdent, targetIdent))
			}
			if abstractfeaturetype.BoundedBy != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := abstractfeaturetype.BoundedBy.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.BoundedBy = %s", abstractfeaturetypeIdent, targetIdent))
			}
		}
	}
	if stage != nil {
		for _, abstractgmltype := range __gong__sortStageSetInstances(stage.AbstractGMLTypes, stage.AbstractGMLType_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			abstractgmltypeIdent := abstractgmltype.GongGetIdentifier(stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&gml.AbstractGMLType{Name: %s}).Stage(stage)", abstractgmltypeIdent, __gong__toRawStringLiteral(abstractgmltype.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", abstractgmltypeIdent, __gong__toRawStringLiteral(abstractgmltype.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Id = %s", abstractgmltypeIdent, __gong__toRawStringLiteral(abstractgmltype.Id)))
			values.WriteString(fmt.Sprintf("\n\t%s.Description = %s", abstractgmltypeIdent, __gong__toRawStringLiteral(abstractgmltype.Description)))
			if abstractgmltype.Identifier != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := abstractgmltype.Identifier.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Identifier = %s", abstractgmltypeIdent, targetIdent))
			}
			for _, elem := range abstractgmltype.NameXSD {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := elem.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.NameXSD = append(%s.NameXSD, %s)", abstractgmltypeIdent, abstractgmltypeIdent, targetIdent))
			}
		}
	}
	if stage != nil {
		for _, abstractgeometricprimitivetype := range __gong__sortStageSetInstances(stage.AbstractGeometricPrimitiveTypes, stage.AbstractGeometricPrimitiveType_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			abstractgeometricprimitivetypeIdent := abstractgeometricprimitivetype.GongGetIdentifier(stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&gml.AbstractGeometricPrimitiveType{Name: %s}).Stage(stage)", abstractgeometricprimitivetypeIdent, __gong__toRawStringLiteral(abstractgeometricprimitivetype.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", abstractgeometricprimitivetypeIdent, __gong__toRawStringLiteral(abstractgeometricprimitivetype.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Id = %s", abstractgeometricprimitivetypeIdent, __gong__toRawStringLiteral(abstractgeometricprimitivetype.Id)))
			values.WriteString(fmt.Sprintf("\n\t%s.Description = %s", abstractgeometricprimitivetypeIdent, __gong__toRawStringLiteral(abstractgeometricprimitivetype.Description)))
			values.WriteString(fmt.Sprintf("\n\t%s.SrsName = %s", abstractgeometricprimitivetypeIdent, __gong__toRawStringLiteral(abstractgeometricprimitivetype.SrsName)))
			if abstractgeometricprimitivetype.Identifier != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := abstractgeometricprimitivetype.Identifier.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Identifier = %s", abstractgeometricprimitivetypeIdent, targetIdent))
			}
			for _, elem := range abstractgeometricprimitivetype.NameXSD {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := elem.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.NameXSD = append(%s.NameXSD, %s)", abstractgeometricprimitivetypeIdent, abstractgeometricprimitivetypeIdent, targetIdent))
			}
		}
	}
	if stage != nil {
		for _, abstractgeometrytype := range __gong__sortStageSetInstances(stage.AbstractGeometryTypes, stage.AbstractGeometryType_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			abstractgeometrytypeIdent := abstractgeometrytype.GongGetIdentifier(stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&gml.AbstractGeometryType{Name: %s}).Stage(stage)", abstractgeometrytypeIdent, __gong__toRawStringLiteral(abstractgeometrytype.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", abstractgeometrytypeIdent, __gong__toRawStringLiteral(abstractgeometrytype.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Id = %s", abstractgeometrytypeIdent, __gong__toRawStringLiteral(abstractgeometrytype.Id)))
			values.WriteString(fmt.Sprintf("\n\t%s.Description = %s", abstractgeometrytypeIdent, __gong__toRawStringLiteral(abstractgeometrytype.Description)))
			values.WriteString(fmt.Sprintf("\n\t%s.SrsName = %s", abstractgeometrytypeIdent, __gong__toRawStringLiteral(abstractgeometrytype.SrsName)))
			if abstractgeometrytype.Identifier != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := abstractgeometrytype.Identifier.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Identifier = %s", abstractgeometrytypeIdent, targetIdent))
			}
			for _, elem := range abstractgeometrytype.NameXSD {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := elem.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.NameXSD = append(%s.NameXSD, %s)", abstractgeometrytypeIdent, abstractgeometrytypeIdent, targetIdent))
			}
		}
	}
	if stage != nil {
		for _, abstractmetadatapropertytype := range __gong__sortStageSetInstances(stage.AbstractMetadataPropertyTypes, stage.AbstractMetadataPropertyType_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			abstractmetadatapropertytypeIdent := abstractmetadatapropertytype.GongGetIdentifier(stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&gml.AbstractMetadataPropertyType{Name: %s}).Stage(stage)", abstractmetadatapropertytypeIdent, __gong__toRawStringLiteral(abstractmetadatapropertytype.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", abstractmetadatapropertytypeIdent, __gong__toRawStringLiteral(abstractmetadatapropertytype.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Owns = %t", abstractmetadatapropertytypeIdent, abstractmetadatapropertytype.Owns))
		}
	}
	if stage != nil {
		for _, abstracttimegeometricprimitivetype := range __gong__sortStageSetInstances(stage.AbstractTimeGeometricPrimitiveTypes, stage.AbstractTimeGeometricPrimitiveType_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			abstracttimegeometricprimitivetypeIdent := abstracttimegeometricprimitivetype.GongGetIdentifier(stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&gml.AbstractTimeGeometricPrimitiveType{Name: %s}).Stage(stage)", abstracttimegeometricprimitivetypeIdent, __gong__toRawStringLiteral(abstracttimegeometricprimitivetype.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", abstracttimegeometricprimitivetypeIdent, __gong__toRawStringLiteral(abstracttimegeometricprimitivetype.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Id = %s", abstracttimegeometricprimitivetypeIdent, __gong__toRawStringLiteral(abstracttimegeometricprimitivetype.Id)))
			values.WriteString(fmt.Sprintf("\n\t%s.Description = %s", abstracttimegeometricprimitivetypeIdent, __gong__toRawStringLiteral(abstracttimegeometricprimitivetype.Description)))
			values.WriteString(fmt.Sprintf("\n\t%s.Frame = %s", abstracttimegeometricprimitivetypeIdent, __gong__toRawStringLiteral(abstracttimegeometricprimitivetype.Frame)))
			if abstracttimegeometricprimitivetype.Identifier != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := abstracttimegeometricprimitivetype.Identifier.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Identifier = %s", abstracttimegeometricprimitivetypeIdent, targetIdent))
			}
			for _, elem := range abstracttimegeometricprimitivetype.NameXSD {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := elem.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.NameXSD = append(%s.NameXSD, %s)", abstracttimegeometricprimitivetypeIdent, abstracttimegeometricprimitivetypeIdent, targetIdent))
			}
		}
	}
	if stage != nil {
		for _, abstracttimeobjecttype := range __gong__sortStageSetInstances(stage.AbstractTimeObjectTypes, stage.AbstractTimeObjectType_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			abstracttimeobjecttypeIdent := abstracttimeobjecttype.GongGetIdentifier(stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&gml.AbstractTimeObjectType{Name: %s}).Stage(stage)", abstracttimeobjecttypeIdent, __gong__toRawStringLiteral(abstracttimeobjecttype.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", abstracttimeobjecttypeIdent, __gong__toRawStringLiteral(abstracttimeobjecttype.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Id = %s", abstracttimeobjecttypeIdent, __gong__toRawStringLiteral(abstracttimeobjecttype.Id)))
			values.WriteString(fmt.Sprintf("\n\t%s.Description = %s", abstracttimeobjecttypeIdent, __gong__toRawStringLiteral(abstracttimeobjecttype.Description)))
			if abstracttimeobjecttype.Identifier != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := abstracttimeobjecttype.Identifier.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Identifier = %s", abstracttimeobjecttypeIdent, targetIdent))
			}
			for _, elem := range abstracttimeobjecttype.NameXSD {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := elem.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.NameXSD = append(%s.NameXSD, %s)", abstracttimeobjecttypeIdent, abstracttimeobjecttypeIdent, targetIdent))
			}
		}
	}
	if stage != nil {
		for _, abstracttimeprimitivetype := range __gong__sortStageSetInstances(stage.AbstractTimePrimitiveTypes, stage.AbstractTimePrimitiveType_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			abstracttimeprimitivetypeIdent := abstracttimeprimitivetype.GongGetIdentifier(stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&gml.AbstractTimePrimitiveType{Name: %s}).Stage(stage)", abstracttimeprimitivetypeIdent, __gong__toRawStringLiteral(abstracttimeprimitivetype.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", abstracttimeprimitivetypeIdent, __gong__toRawStringLiteral(abstracttimeprimitivetype.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Id = %s", abstracttimeprimitivetypeIdent, __gong__toRawStringLiteral(abstracttimeprimitivetype.Id)))
			values.WriteString(fmt.Sprintf("\n\t%s.Description = %s", abstracttimeprimitivetypeIdent, __gong__toRawStringLiteral(abstracttimeprimitivetype.Description)))
			if abstracttimeprimitivetype.Identifier != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := abstracttimeprimitivetype.Identifier.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Identifier = %s", abstracttimeprimitivetypeIdent, targetIdent))
			}
			for _, elem := range abstracttimeprimitivetype.NameXSD {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := elem.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.NameXSD = append(%s.NameXSD, %s)", abstracttimeprimitivetypeIdent, abstracttimeprimitivetypeIdent, targetIdent))
			}
		}
	}
	if stage != nil {
		for _, abstracttimeslicetype := range __gong__sortStageSetInstances(stage.AbstractTimeSliceTypes, stage.AbstractTimeSliceType_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			abstracttimeslicetypeIdent := abstracttimeslicetype.GongGetIdentifier(stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&gml.AbstractTimeSliceType{Name: %s}).Stage(stage)", abstracttimeslicetypeIdent, __gong__toRawStringLiteral(abstracttimeslicetype.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", abstracttimeslicetypeIdent, __gong__toRawStringLiteral(abstracttimeslicetype.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Id = %s", abstracttimeslicetypeIdent, __gong__toRawStringLiteral(abstracttimeslicetype.Id)))
			values.WriteString(fmt.Sprintf("\n\t%s.Description = %s", abstracttimeslicetypeIdent, __gong__toRawStringLiteral(abstracttimeslicetype.Description)))
			if abstracttimeslicetype.Identifier != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := abstracttimeslicetype.Identifier.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Identifier = %s", abstracttimeslicetypeIdent, targetIdent))
			}
			for _, elem := range abstracttimeslicetype.NameXSD {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := elem.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.NameXSD = append(%s.NameXSD, %s)", abstracttimeslicetypeIdent, abstracttimeslicetypeIdent, targetIdent))
			}
			if abstracttimeslicetype.ValidTime != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := abstracttimeslicetype.ValidTime.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.ValidTime = %s", abstracttimeslicetypeIdent, targetIdent))
			}
		}
	}
	if stage != nil {
		for _, boundingshapetype := range __gong__sortStageSetInstances(stage.BoundingShapeTypes, stage.BoundingShapeType_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			boundingshapetypeIdent := boundingshapetype.GongGetIdentifier(stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&gml.BoundingShapeType{Name: %s}).Stage(stage)", boundingshapetypeIdent, __gong__toRawStringLiteral(boundingshapetype.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", boundingshapetypeIdent, __gong__toRawStringLiteral(boundingshapetype.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.NilReason = %s", boundingshapetypeIdent, __gong__toRawStringLiteral(boundingshapetype.NilReason)))
			if boundingshapetype.Envelope != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := boundingshapetype.Envelope.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Envelope = %s", boundingshapetypeIdent, targetIdent))
			}
		}
	}
	if stage != nil {
		for _, codetype := range __gong__sortStageSetInstances(stage.CodeTypes, stage.CodeType_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			codetypeIdent := codetype.GongGetIdentifier(stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&gml.CodeType{Name: %s}).Stage(stage)", codetypeIdent, __gong__toRawStringLiteral(codetype.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", codetypeIdent, __gong__toRawStringLiteral(codetype.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.CodeSpace = %s", codetypeIdent, __gong__toRawStringLiteral(codetype.CodeSpace)))
			values.WriteString(fmt.Sprintf("\n\t%s.EnclosedText = %s", codetypeIdent, __gong__toRawStringLiteral(codetype.EnclosedText)))
		}
	}
	if stage != nil {
		for _, codewithauthoritytype := range __gong__sortStageSetInstances(stage.CodeWithAuthorityTypes, stage.CodeWithAuthorityType_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			codewithauthoritytypeIdent := codewithauthoritytype.GongGetIdentifier(stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&gml.CodeWithAuthorityType{Name: %s}).Stage(stage)", codewithauthoritytypeIdent, __gong__toRawStringLiteral(codewithauthoritytype.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", codewithauthoritytypeIdent, __gong__toRawStringLiteral(codewithauthoritytype.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.CodeSpace = %s", codewithauthoritytypeIdent, __gong__toRawStringLiteral(codewithauthoritytype.CodeSpace)))
			values.WriteString(fmt.Sprintf("\n\t%s.EnclosedText = %s", codewithauthoritytypeIdent, __gong__toRawStringLiteral(codewithauthoritytype.EnclosedText)))
		}
	}
	if stage != nil {
		for _, directpositiontype := range __gong__sortStageSetInstances(stage.DirectPositionTypes, stage.DirectPositionType_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			directpositiontypeIdent := directpositiontype.GongGetIdentifier(stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&gml.DirectPositionType{Name: %s}).Stage(stage)", directpositiontypeIdent, __gong__toRawStringLiteral(directpositiontype.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", directpositiontypeIdent, __gong__toRawStringLiteral(directpositiontype.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.SrsName = %s", directpositiontypeIdent, __gong__toRawStringLiteral(directpositiontype.SrsName)))
			values.WriteString(fmt.Sprintf("\n\t%s.EnclosedText = %s", directpositiontypeIdent, __gong__toRawStringLiteral(directpositiontype.EnclosedText)))
		}
	}
	if stage != nil {
		for _, dynamicfeaturetype := range __gong__sortStageSetInstances(stage.DynamicFeatureTypes, stage.DynamicFeatureType_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			dynamicfeaturetypeIdent := dynamicfeaturetype.GongGetIdentifier(stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&gml.DynamicFeatureType{Name: %s}).Stage(stage)", dynamicfeaturetypeIdent, __gong__toRawStringLiteral(dynamicfeaturetype.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", dynamicfeaturetypeIdent, __gong__toRawStringLiteral(dynamicfeaturetype.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Id = %s", dynamicfeaturetypeIdent, __gong__toRawStringLiteral(dynamicfeaturetype.Id)))
			values.WriteString(fmt.Sprintf("\n\t%s.Description = %s", dynamicfeaturetypeIdent, __gong__toRawStringLiteral(dynamicfeaturetype.Description)))
			if dynamicfeaturetype.Identifier != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := dynamicfeaturetype.Identifier.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Identifier = %s", dynamicfeaturetypeIdent, targetIdent))
			}
			for _, elem := range dynamicfeaturetype.NameXSD {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := elem.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.NameXSD = append(%s.NameXSD, %s)", dynamicfeaturetypeIdent, dynamicfeaturetypeIdent, targetIdent))
			}
			if dynamicfeaturetype.BoundedBy != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := dynamicfeaturetype.BoundedBy.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.BoundedBy = %s", dynamicfeaturetypeIdent, targetIdent))
			}
		}
	}
	if stage != nil {
		for _, envelopetype := range __gong__sortStageSetInstances(stage.EnvelopeTypes, stage.EnvelopeType_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			envelopetypeIdent := envelopetype.GongGetIdentifier(stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&gml.EnvelopeType{Name: %s}).Stage(stage)", envelopetypeIdent, __gong__toRawStringLiteral(envelopetype.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", envelopetypeIdent, __gong__toRawStringLiteral(envelopetype.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.SrsName = %s", envelopetypeIdent, __gong__toRawStringLiteral(envelopetype.SrsName)))
			if envelopetype.LowerCorner != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := envelopetype.LowerCorner.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.LowerCorner = %s", envelopetypeIdent, targetIdent))
			}
			if envelopetype.UpperCorner != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := envelopetype.UpperCorner.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.UpperCorner = %s", envelopetypeIdent, targetIdent))
			}
		}
	}
	if stage != nil {
		for _, pointtype := range __gong__sortStageSetInstances(stage.PointTypes, stage.PointType_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			pointtypeIdent := pointtype.GongGetIdentifier(stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&gml.PointType{Name: %s}).Stage(stage)", pointtypeIdent, __gong__toRawStringLiteral(pointtype.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", pointtypeIdent, __gong__toRawStringLiteral(pointtype.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Id = %s", pointtypeIdent, __gong__toRawStringLiteral(pointtype.Id)))
			values.WriteString(fmt.Sprintf("\n\t%s.Description = %s", pointtypeIdent, __gong__toRawStringLiteral(pointtype.Description)))
			values.WriteString(fmt.Sprintf("\n\t%s.SrsName = %s", pointtypeIdent, __gong__toRawStringLiteral(pointtype.SrsName)))
			if pointtype.Identifier != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := pointtype.Identifier.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Identifier = %s", pointtypeIdent, targetIdent))
			}
			for _, elem := range pointtype.NameXSD {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := elem.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.NameXSD = append(%s.NameXSD, %s)", pointtypeIdent, pointtypeIdent, targetIdent))
			}
			if pointtype.Pos != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := pointtype.Pos.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Pos = %s", pointtypeIdent, targetIdent))
			}
		}
	}
	if stage != nil {
		for _, timeinstanttype := range __gong__sortStageSetInstances(stage.TimeInstantTypes, stage.TimeInstantType_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			timeinstanttypeIdent := timeinstanttype.GongGetIdentifier(stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&gml.TimeInstantType{Name: %s}).Stage(stage)", timeinstanttypeIdent, __gong__toRawStringLiteral(timeinstanttype.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", timeinstanttypeIdent, __gong__toRawStringLiteral(timeinstanttype.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Id = %s", timeinstanttypeIdent, __gong__toRawStringLiteral(timeinstanttype.Id)))
			values.WriteString(fmt.Sprintf("\n\t%s.Description = %s", timeinstanttypeIdent, __gong__toRawStringLiteral(timeinstanttype.Description)))
			values.WriteString(fmt.Sprintf("\n\t%s.Frame = %s", timeinstanttypeIdent, __gong__toRawStringLiteral(timeinstanttype.Frame)))
			if timeinstanttype.Identifier != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := timeinstanttype.Identifier.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Identifier = %s", timeinstanttypeIdent, targetIdent))
			}
			for _, elem := range timeinstanttype.NameXSD {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := elem.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.NameXSD = append(%s.NameXSD, %s)", timeinstanttypeIdent, timeinstanttypeIdent, targetIdent))
			}
			if timeinstanttype.TimePosition != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := timeinstanttype.TimePosition.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.TimePosition = %s", timeinstanttypeIdent, targetIdent))
			}
		}
	}
	if stage != nil {
		for _, timeperiodtype := range __gong__sortStageSetInstances(stage.TimePeriodTypes, stage.TimePeriodType_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			timeperiodtypeIdent := timeperiodtype.GongGetIdentifier(stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&gml.TimePeriodType{Name: %s}).Stage(stage)", timeperiodtypeIdent, __gong__toRawStringLiteral(timeperiodtype.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", timeperiodtypeIdent, __gong__toRawStringLiteral(timeperiodtype.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Id = %s", timeperiodtypeIdent, __gong__toRawStringLiteral(timeperiodtype.Id)))
			values.WriteString(fmt.Sprintf("\n\t%s.Description = %s", timeperiodtypeIdent, __gong__toRawStringLiteral(timeperiodtype.Description)))
			values.WriteString(fmt.Sprintf("\n\t%s.Frame = %s", timeperiodtypeIdent, __gong__toRawStringLiteral(timeperiodtype.Frame)))
			if timeperiodtype.Identifier != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := timeperiodtype.Identifier.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Identifier = %s", timeperiodtypeIdent, targetIdent))
			}
			for _, elem := range timeperiodtype.NameXSD {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := elem.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.NameXSD = append(%s.NameXSD, %s)", timeperiodtypeIdent, timeperiodtypeIdent, targetIdent))
			}
			if timeperiodtype.BeginPosition != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := timeperiodtype.BeginPosition.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.BeginPosition = %s", timeperiodtypeIdent, targetIdent))
			}
			if timeperiodtype.EndPosition != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := timeperiodtype.EndPosition.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.EndPosition = %s", timeperiodtypeIdent, targetIdent))
			}
		}
	}
	if stage != nil {
		for _, timepositiontype := range __gong__sortStageSetInstances(stage.TimePositionTypes, stage.TimePositionType_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			timepositiontypeIdent := timepositiontype.GongGetIdentifier(stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&gml.TimePositionType{Name: %s}).Stage(stage)", timepositiontypeIdent, __gong__toRawStringLiteral(timepositiontype.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
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
	if stage != nil {
		for _, timeprimitivepropertytype := range __gong__sortStageSetInstances(stage.TimePrimitivePropertyTypes, stage.TimePrimitivePropertyType_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			timeprimitivepropertytypeIdent := timeprimitivepropertytype.GongGetIdentifier(stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&gml.TimePrimitivePropertyType{Name: %s}).Stage(stage)", timeprimitivepropertytypeIdent, __gong__toRawStringLiteral(timeprimitivepropertytype.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
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
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := timeprimitivepropertytype.TimeInstant.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.TimeInstant = %s", timeprimitivepropertytypeIdent, targetIdent))
			}
			if timeprimitivepropertytype.TimePeriod != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := timeprimitivepropertytype.TimePeriod.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.TimePeriod = %s", timeprimitivepropertytypeIdent, targetIdent))
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
go run github.com/fullstack-lang/gong/app/xsd/tests/aixm/service/go/models/gml/probe/cmd/probe@%s stage.go
*/
package %s

import (
	"slices"
	"time"

	"github.com/fullstack-lang/gong/app/xsd/tests/aixm/service/go/models/gml"%s
)

var (
	_ time.Time
	_ = slices.Index[[]int, int]

	_ *gml.Stage
)%s

// function will stage objects across all coordinated stages
func _(stage *gml.Stage) {

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

	if modelsPackageName != "" && modelsPackageName != "github.com/fullstack-lang/gong/app/xsd/tests/aixm/service/go/models/gml" {
		res = strings.ReplaceAll(res, "\"github.com/fullstack-lang/gong/app/xsd/tests/aixm/service/go/models/gml\"", fmt.Sprintf("%q", modelsPackageName))
	}

	return res, nil
}

// insertion point for marshall field methods
func (abstractfeaturemembertype *AbstractFeatureMemberType) GongMarshallField(stage *Stage, fieldName string) (res string) {
	ident := abstractfeaturemembertype.GongGetIdentifier(stage)
	_ = ident
	switch fieldName {
	case "Name":
		res = __gong__marshallString(ident, "Name", abstractfeaturemembertype.Name)

	default:
		log.Panicf("Unknown field %s for Gongstruct AbstractFeatureMemberType", fieldName)
	}
	return
}

func (abstractfeaturetype *AbstractFeatureType) GongMarshallField(stage *Stage, fieldName string) (res string) {
	ident := abstractfeaturetype.GongGetIdentifier(stage)
	_ = ident
	switch fieldName {
	case "Name":
		res = __gong__marshallString(ident, "Name", abstractfeaturetype.Name)
	case "Id":
		res = __gong__marshallString(ident, "Id", abstractfeaturetype.Id)
	case "Description":
		res = __gong__marshallString(ident, "Description", abstractfeaturetype.Description)

	case "Identifier":
		if abstractfeaturetype.Identifier != nil {
			res = __gong__marshallPointer(ident, "Identifier", abstractfeaturetype.Identifier.GongGetIdentifier(stage))
		} else {
			res = __gong__marshallPointer(ident, "Identifier", "nil")
		}
	case "NameXSD":
		var sb strings.Builder
		for _, _codetype := range abstractfeaturetype.NameXSD {
			sb.WriteString(__gong__marshallSliceOfPointers(ident, "NameXSD", _codetype.GongGetIdentifier(stage)))
		}
		res = sb.String()
	case "BoundedBy":
		if abstractfeaturetype.BoundedBy != nil {
			res = __gong__marshallPointer(ident, "BoundedBy", abstractfeaturetype.BoundedBy.GongGetIdentifier(stage))
		} else {
			res = __gong__marshallPointer(ident, "BoundedBy", "nil")
		}
	default:
		log.Panicf("Unknown field %s for Gongstruct AbstractFeatureType", fieldName)
	}
	return
}

func (abstractgmltype *AbstractGMLType) GongMarshallField(stage *Stage, fieldName string) (res string) {
	ident := abstractgmltype.GongGetIdentifier(stage)
	_ = ident
	switch fieldName {
	case "Name":
		res = __gong__marshallString(ident, "Name", abstractgmltype.Name)
	case "Id":
		res = __gong__marshallString(ident, "Id", abstractgmltype.Id)
	case "Description":
		res = __gong__marshallString(ident, "Description", abstractgmltype.Description)

	case "Identifier":
		if abstractgmltype.Identifier != nil {
			res = __gong__marshallPointer(ident, "Identifier", abstractgmltype.Identifier.GongGetIdentifier(stage))
		} else {
			res = __gong__marshallPointer(ident, "Identifier", "nil")
		}
	case "NameXSD":
		var sb strings.Builder
		for _, _codetype := range abstractgmltype.NameXSD {
			sb.WriteString(__gong__marshallSliceOfPointers(ident, "NameXSD", _codetype.GongGetIdentifier(stage)))
		}
		res = sb.String()
	default:
		log.Panicf("Unknown field %s for Gongstruct AbstractGMLType", fieldName)
	}
	return
}

func (abstractgeometricprimitivetype *AbstractGeometricPrimitiveType) GongMarshallField(stage *Stage, fieldName string) (res string) {
	ident := abstractgeometricprimitivetype.GongGetIdentifier(stage)
	_ = ident
	switch fieldName {
	case "Name":
		res = __gong__marshallString(ident, "Name", abstractgeometricprimitivetype.Name)
	case "Id":
		res = __gong__marshallString(ident, "Id", abstractgeometricprimitivetype.Id)
	case "Description":
		res = __gong__marshallString(ident, "Description", abstractgeometricprimitivetype.Description)
	case "SrsName":
		res = __gong__marshallString(ident, "SrsName", abstractgeometricprimitivetype.SrsName)

	case "Identifier":
		if abstractgeometricprimitivetype.Identifier != nil {
			res = __gong__marshallPointer(ident, "Identifier", abstractgeometricprimitivetype.Identifier.GongGetIdentifier(stage))
		} else {
			res = __gong__marshallPointer(ident, "Identifier", "nil")
		}
	case "NameXSD":
		var sb strings.Builder
		for _, _codetype := range abstractgeometricprimitivetype.NameXSD {
			sb.WriteString(__gong__marshallSliceOfPointers(ident, "NameXSD", _codetype.GongGetIdentifier(stage)))
		}
		res = sb.String()
	default:
		log.Panicf("Unknown field %s for Gongstruct AbstractGeometricPrimitiveType", fieldName)
	}
	return
}

func (abstractgeometrytype *AbstractGeometryType) GongMarshallField(stage *Stage, fieldName string) (res string) {
	ident := abstractgeometrytype.GongGetIdentifier(stage)
	_ = ident
	switch fieldName {
	case "Name":
		res = __gong__marshallString(ident, "Name", abstractgeometrytype.Name)
	case "Id":
		res = __gong__marshallString(ident, "Id", abstractgeometrytype.Id)
	case "Description":
		res = __gong__marshallString(ident, "Description", abstractgeometrytype.Description)
	case "SrsName":
		res = __gong__marshallString(ident, "SrsName", abstractgeometrytype.SrsName)

	case "Identifier":
		if abstractgeometrytype.Identifier != nil {
			res = __gong__marshallPointer(ident, "Identifier", abstractgeometrytype.Identifier.GongGetIdentifier(stage))
		} else {
			res = __gong__marshallPointer(ident, "Identifier", "nil")
		}
	case "NameXSD":
		var sb strings.Builder
		for _, _codetype := range abstractgeometrytype.NameXSD {
			sb.WriteString(__gong__marshallSliceOfPointers(ident, "NameXSD", _codetype.GongGetIdentifier(stage)))
		}
		res = sb.String()
	default:
		log.Panicf("Unknown field %s for Gongstruct AbstractGeometryType", fieldName)
	}
	return
}

func (abstractmetadatapropertytype *AbstractMetadataPropertyType) GongMarshallField(stage *Stage, fieldName string) (res string) {
	ident := abstractmetadatapropertytype.GongGetIdentifier(stage)
	_ = ident
	switch fieldName {
	case "Name":
		res = __gong__marshallString(ident, "Name", abstractmetadatapropertytype.Name)
	case "Owns":
		res = __gong__marshallBool(ident, "Owns", abstractmetadatapropertytype.Owns)

	default:
		log.Panicf("Unknown field %s for Gongstruct AbstractMetadataPropertyType", fieldName)
	}
	return
}

func (abstracttimegeometricprimitivetype *AbstractTimeGeometricPrimitiveType) GongMarshallField(stage *Stage, fieldName string) (res string) {
	ident := abstracttimegeometricprimitivetype.GongGetIdentifier(stage)
	_ = ident
	switch fieldName {
	case "Name":
		res = __gong__marshallString(ident, "Name", abstracttimegeometricprimitivetype.Name)
	case "Id":
		res = __gong__marshallString(ident, "Id", abstracttimegeometricprimitivetype.Id)
	case "Description":
		res = __gong__marshallString(ident, "Description", abstracttimegeometricprimitivetype.Description)
	case "Frame":
		res = __gong__marshallString(ident, "Frame", abstracttimegeometricprimitivetype.Frame)

	case "Identifier":
		if abstracttimegeometricprimitivetype.Identifier != nil {
			res = __gong__marshallPointer(ident, "Identifier", abstracttimegeometricprimitivetype.Identifier.GongGetIdentifier(stage))
		} else {
			res = __gong__marshallPointer(ident, "Identifier", "nil")
		}
	case "NameXSD":
		var sb strings.Builder
		for _, _codetype := range abstracttimegeometricprimitivetype.NameXSD {
			sb.WriteString(__gong__marshallSliceOfPointers(ident, "NameXSD", _codetype.GongGetIdentifier(stage)))
		}
		res = sb.String()
	default:
		log.Panicf("Unknown field %s for Gongstruct AbstractTimeGeometricPrimitiveType", fieldName)
	}
	return
}

func (abstracttimeobjecttype *AbstractTimeObjectType) GongMarshallField(stage *Stage, fieldName string) (res string) {
	ident := abstracttimeobjecttype.GongGetIdentifier(stage)
	_ = ident
	switch fieldName {
	case "Name":
		res = __gong__marshallString(ident, "Name", abstracttimeobjecttype.Name)
	case "Id":
		res = __gong__marshallString(ident, "Id", abstracttimeobjecttype.Id)
	case "Description":
		res = __gong__marshallString(ident, "Description", abstracttimeobjecttype.Description)

	case "Identifier":
		if abstracttimeobjecttype.Identifier != nil {
			res = __gong__marshallPointer(ident, "Identifier", abstracttimeobjecttype.Identifier.GongGetIdentifier(stage))
		} else {
			res = __gong__marshallPointer(ident, "Identifier", "nil")
		}
	case "NameXSD":
		var sb strings.Builder
		for _, _codetype := range abstracttimeobjecttype.NameXSD {
			sb.WriteString(__gong__marshallSliceOfPointers(ident, "NameXSD", _codetype.GongGetIdentifier(stage)))
		}
		res = sb.String()
	default:
		log.Panicf("Unknown field %s for Gongstruct AbstractTimeObjectType", fieldName)
	}
	return
}

func (abstracttimeprimitivetype *AbstractTimePrimitiveType) GongMarshallField(stage *Stage, fieldName string) (res string) {
	ident := abstracttimeprimitivetype.GongGetIdentifier(stage)
	_ = ident
	switch fieldName {
	case "Name":
		res = __gong__marshallString(ident, "Name", abstracttimeprimitivetype.Name)
	case "Id":
		res = __gong__marshallString(ident, "Id", abstracttimeprimitivetype.Id)
	case "Description":
		res = __gong__marshallString(ident, "Description", abstracttimeprimitivetype.Description)

	case "Identifier":
		if abstracttimeprimitivetype.Identifier != nil {
			res = __gong__marshallPointer(ident, "Identifier", abstracttimeprimitivetype.Identifier.GongGetIdentifier(stage))
		} else {
			res = __gong__marshallPointer(ident, "Identifier", "nil")
		}
	case "NameXSD":
		var sb strings.Builder
		for _, _codetype := range abstracttimeprimitivetype.NameXSD {
			sb.WriteString(__gong__marshallSliceOfPointers(ident, "NameXSD", _codetype.GongGetIdentifier(stage)))
		}
		res = sb.String()
	default:
		log.Panicf("Unknown field %s for Gongstruct AbstractTimePrimitiveType", fieldName)
	}
	return
}

func (abstracttimeslicetype *AbstractTimeSliceType) GongMarshallField(stage *Stage, fieldName string) (res string) {
	ident := abstracttimeslicetype.GongGetIdentifier(stage)
	_ = ident
	switch fieldName {
	case "Name":
		res = __gong__marshallString(ident, "Name", abstracttimeslicetype.Name)
	case "Id":
		res = __gong__marshallString(ident, "Id", abstracttimeslicetype.Id)
	case "Description":
		res = __gong__marshallString(ident, "Description", abstracttimeslicetype.Description)

	case "Identifier":
		if abstracttimeslicetype.Identifier != nil {
			res = __gong__marshallPointer(ident, "Identifier", abstracttimeslicetype.Identifier.GongGetIdentifier(stage))
		} else {
			res = __gong__marshallPointer(ident, "Identifier", "nil")
		}
	case "NameXSD":
		var sb strings.Builder
		for _, _codetype := range abstracttimeslicetype.NameXSD {
			sb.WriteString(__gong__marshallSliceOfPointers(ident, "NameXSD", _codetype.GongGetIdentifier(stage)))
		}
		res = sb.String()
	case "ValidTime":
		if abstracttimeslicetype.ValidTime != nil {
			res = __gong__marshallPointer(ident, "ValidTime", abstracttimeslicetype.ValidTime.GongGetIdentifier(stage))
		} else {
			res = __gong__marshallPointer(ident, "ValidTime", "nil")
		}
	default:
		log.Panicf("Unknown field %s for Gongstruct AbstractTimeSliceType", fieldName)
	}
	return
}

func (boundingshapetype *BoundingShapeType) GongMarshallField(stage *Stage, fieldName string) (res string) {
	ident := boundingshapetype.GongGetIdentifier(stage)
	_ = ident
	switch fieldName {
	case "Name":
		res = __gong__marshallString(ident, "Name", boundingshapetype.Name)
	case "NilReason":
		res = __gong__marshallString(ident, "NilReason", boundingshapetype.NilReason)

	case "Envelope":
		if boundingshapetype.Envelope != nil {
			res = __gong__marshallPointer(ident, "Envelope", boundingshapetype.Envelope.GongGetIdentifier(stage))
		} else {
			res = __gong__marshallPointer(ident, "Envelope", "nil")
		}
	default:
		log.Panicf("Unknown field %s for Gongstruct BoundingShapeType", fieldName)
	}
	return
}

func (codetype *CodeType) GongMarshallField(stage *Stage, fieldName string) (res string) {
	ident := codetype.GongGetIdentifier(stage)
	_ = ident
	switch fieldName {
	case "Name":
		res = __gong__marshallString(ident, "Name", codetype.Name)
	case "CodeSpace":
		res = __gong__marshallString(ident, "CodeSpace", codetype.CodeSpace)
	case "EnclosedText":
		res = __gong__marshallString(ident, "EnclosedText", codetype.EnclosedText)

	default:
		log.Panicf("Unknown field %s for Gongstruct CodeType", fieldName)
	}
	return
}

func (codewithauthoritytype *CodeWithAuthorityType) GongMarshallField(stage *Stage, fieldName string) (res string) {
	ident := codewithauthoritytype.GongGetIdentifier(stage)
	_ = ident
	switch fieldName {
	case "Name":
		res = __gong__marshallString(ident, "Name", codewithauthoritytype.Name)
	case "CodeSpace":
		res = __gong__marshallString(ident, "CodeSpace", codewithauthoritytype.CodeSpace)
	case "EnclosedText":
		res = __gong__marshallString(ident, "EnclosedText", codewithauthoritytype.EnclosedText)

	default:
		log.Panicf("Unknown field %s for Gongstruct CodeWithAuthorityType", fieldName)
	}
	return
}

func (directpositiontype *DirectPositionType) GongMarshallField(stage *Stage, fieldName string) (res string) {
	ident := directpositiontype.GongGetIdentifier(stage)
	_ = ident
	switch fieldName {
	case "Name":
		res = __gong__marshallString(ident, "Name", directpositiontype.Name)
	case "SrsName":
		res = __gong__marshallString(ident, "SrsName", directpositiontype.SrsName)
	case "EnclosedText":
		res = __gong__marshallString(ident, "EnclosedText", directpositiontype.EnclosedText)

	default:
		log.Panicf("Unknown field %s for Gongstruct DirectPositionType", fieldName)
	}
	return
}

func (dynamicfeaturetype *DynamicFeatureType) GongMarshallField(stage *Stage, fieldName string) (res string) {
	ident := dynamicfeaturetype.GongGetIdentifier(stage)
	_ = ident
	switch fieldName {
	case "Name":
		res = __gong__marshallString(ident, "Name", dynamicfeaturetype.Name)
	case "Id":
		res = __gong__marshallString(ident, "Id", dynamicfeaturetype.Id)
	case "Description":
		res = __gong__marshallString(ident, "Description", dynamicfeaturetype.Description)

	case "Identifier":
		if dynamicfeaturetype.Identifier != nil {
			res = __gong__marshallPointer(ident, "Identifier", dynamicfeaturetype.Identifier.GongGetIdentifier(stage))
		} else {
			res = __gong__marshallPointer(ident, "Identifier", "nil")
		}
	case "NameXSD":
		var sb strings.Builder
		for _, _codetype := range dynamicfeaturetype.NameXSD {
			sb.WriteString(__gong__marshallSliceOfPointers(ident, "NameXSD", _codetype.GongGetIdentifier(stage)))
		}
		res = sb.String()
	case "BoundedBy":
		if dynamicfeaturetype.BoundedBy != nil {
			res = __gong__marshallPointer(ident, "BoundedBy", dynamicfeaturetype.BoundedBy.GongGetIdentifier(stage))
		} else {
			res = __gong__marshallPointer(ident, "BoundedBy", "nil")
		}
	default:
		log.Panicf("Unknown field %s for Gongstruct DynamicFeatureType", fieldName)
	}
	return
}

func (envelopetype *EnvelopeType) GongMarshallField(stage *Stage, fieldName string) (res string) {
	ident := envelopetype.GongGetIdentifier(stage)
	_ = ident
	switch fieldName {
	case "Name":
		res = __gong__marshallString(ident, "Name", envelopetype.Name)
	case "SrsName":
		res = __gong__marshallString(ident, "SrsName", envelopetype.SrsName)

	case "LowerCorner":
		if envelopetype.LowerCorner != nil {
			res = __gong__marshallPointer(ident, "LowerCorner", envelopetype.LowerCorner.GongGetIdentifier(stage))
		} else {
			res = __gong__marshallPointer(ident, "LowerCorner", "nil")
		}
	case "UpperCorner":
		if envelopetype.UpperCorner != nil {
			res = __gong__marshallPointer(ident, "UpperCorner", envelopetype.UpperCorner.GongGetIdentifier(stage))
		} else {
			res = __gong__marshallPointer(ident, "UpperCorner", "nil")
		}
	default:
		log.Panicf("Unknown field %s for Gongstruct EnvelopeType", fieldName)
	}
	return
}

func (pointtype *PointType) GongMarshallField(stage *Stage, fieldName string) (res string) {
	ident := pointtype.GongGetIdentifier(stage)
	_ = ident
	switch fieldName {
	case "Name":
		res = __gong__marshallString(ident, "Name", pointtype.Name)
	case "Id":
		res = __gong__marshallString(ident, "Id", pointtype.Id)
	case "Description":
		res = __gong__marshallString(ident, "Description", pointtype.Description)
	case "SrsName":
		res = __gong__marshallString(ident, "SrsName", pointtype.SrsName)

	case "Identifier":
		if pointtype.Identifier != nil {
			res = __gong__marshallPointer(ident, "Identifier", pointtype.Identifier.GongGetIdentifier(stage))
		} else {
			res = __gong__marshallPointer(ident, "Identifier", "nil")
		}
	case "NameXSD":
		var sb strings.Builder
		for _, _codetype := range pointtype.NameXSD {
			sb.WriteString(__gong__marshallSliceOfPointers(ident, "NameXSD", _codetype.GongGetIdentifier(stage)))
		}
		res = sb.String()
	case "Pos":
		if pointtype.Pos != nil {
			res = __gong__marshallPointer(ident, "Pos", pointtype.Pos.GongGetIdentifier(stage))
		} else {
			res = __gong__marshallPointer(ident, "Pos", "nil")
		}
	default:
		log.Panicf("Unknown field %s for Gongstruct PointType", fieldName)
	}
	return
}

func (timeinstanttype *TimeInstantType) GongMarshallField(stage *Stage, fieldName string) (res string) {
	ident := timeinstanttype.GongGetIdentifier(stage)
	_ = ident
	switch fieldName {
	case "Name":
		res = __gong__marshallString(ident, "Name", timeinstanttype.Name)
	case "Id":
		res = __gong__marshallString(ident, "Id", timeinstanttype.Id)
	case "Description":
		res = __gong__marshallString(ident, "Description", timeinstanttype.Description)
	case "Frame":
		res = __gong__marshallString(ident, "Frame", timeinstanttype.Frame)

	case "Identifier":
		if timeinstanttype.Identifier != nil {
			res = __gong__marshallPointer(ident, "Identifier", timeinstanttype.Identifier.GongGetIdentifier(stage))
		} else {
			res = __gong__marshallPointer(ident, "Identifier", "nil")
		}
	case "NameXSD":
		var sb strings.Builder
		for _, _codetype := range timeinstanttype.NameXSD {
			sb.WriteString(__gong__marshallSliceOfPointers(ident, "NameXSD", _codetype.GongGetIdentifier(stage)))
		}
		res = sb.String()
	case "TimePosition":
		if timeinstanttype.TimePosition != nil {
			res = __gong__marshallPointer(ident, "TimePosition", timeinstanttype.TimePosition.GongGetIdentifier(stage))
		} else {
			res = __gong__marshallPointer(ident, "TimePosition", "nil")
		}
	default:
		log.Panicf("Unknown field %s for Gongstruct TimeInstantType", fieldName)
	}
	return
}

func (timeperiodtype *TimePeriodType) GongMarshallField(stage *Stage, fieldName string) (res string) {
	ident := timeperiodtype.GongGetIdentifier(stage)
	_ = ident
	switch fieldName {
	case "Name":
		res = __gong__marshallString(ident, "Name", timeperiodtype.Name)
	case "Id":
		res = __gong__marshallString(ident, "Id", timeperiodtype.Id)
	case "Description":
		res = __gong__marshallString(ident, "Description", timeperiodtype.Description)
	case "Frame":
		res = __gong__marshallString(ident, "Frame", timeperiodtype.Frame)

	case "Identifier":
		if timeperiodtype.Identifier != nil {
			res = __gong__marshallPointer(ident, "Identifier", timeperiodtype.Identifier.GongGetIdentifier(stage))
		} else {
			res = __gong__marshallPointer(ident, "Identifier", "nil")
		}
	case "NameXSD":
		var sb strings.Builder
		for _, _codetype := range timeperiodtype.NameXSD {
			sb.WriteString(__gong__marshallSliceOfPointers(ident, "NameXSD", _codetype.GongGetIdentifier(stage)))
		}
		res = sb.String()
	case "BeginPosition":
		if timeperiodtype.BeginPosition != nil {
			res = __gong__marshallPointer(ident, "BeginPosition", timeperiodtype.BeginPosition.GongGetIdentifier(stage))
		} else {
			res = __gong__marshallPointer(ident, "BeginPosition", "nil")
		}
	case "EndPosition":
		if timeperiodtype.EndPosition != nil {
			res = __gong__marshallPointer(ident, "EndPosition", timeperiodtype.EndPosition.GongGetIdentifier(stage))
		} else {
			res = __gong__marshallPointer(ident, "EndPosition", "nil")
		}
	default:
		log.Panicf("Unknown field %s for Gongstruct TimePeriodType", fieldName)
	}
	return
}

func (timepositiontype *TimePositionType) GongMarshallField(stage *Stage, fieldName string) (res string) {
	ident := timepositiontype.GongGetIdentifier(stage)
	_ = ident
	switch fieldName {
	case "Name":
		res = __gong__marshallString(ident, "Name", timepositiontype.Name)
	case "Frame":
		res = __gong__marshallString(ident, "Frame", timepositiontype.Frame)
	case "IndeterminatePosition":
		res = __gong__marshallEnumString(ident, "IndeterminatePosition", timepositiontype.IndeterminatePosition.ToCodeString())
	case "EnclosedText":
		res = __gong__marshallString(ident, "EnclosedText", timepositiontype.EnclosedText)

	default:
		log.Panicf("Unknown field %s for Gongstruct TimePositionType", fieldName)
	}
	return
}

func (timeprimitivepropertytype *TimePrimitivePropertyType) GongMarshallField(stage *Stage, fieldName string) (res string) {
	ident := timeprimitivepropertytype.GongGetIdentifier(stage)
	_ = ident
	switch fieldName {
	case "Name":
		res = __gong__marshallString(ident, "Name", timeprimitivepropertytype.Name)
	case "NilReason":
		res = __gong__marshallString(ident, "NilReason", timeprimitivepropertytype.NilReason)
	case "Type_XSD":
		res = __gong__marshallString(ident, "Type_XSD", timeprimitivepropertytype.Type_XSD)
	case "Href":
		res = __gong__marshallString(ident, "Href", timeprimitivepropertytype.Href)
	case "Role":
		res = __gong__marshallString(ident, "Role", timeprimitivepropertytype.Role)
	case "Arcrole":
		res = __gong__marshallString(ident, "Arcrole", timeprimitivepropertytype.Arcrole)
	case "Title":
		res = __gong__marshallString(ident, "Title", timeprimitivepropertytype.Title)
	case "Show":
		res = __gong__marshallString(ident, "Show", timeprimitivepropertytype.Show)
	case "Actuate":
		res = __gong__marshallString(ident, "Actuate", timeprimitivepropertytype.Actuate)

	case "TimeInstant":
		if timeprimitivepropertytype.TimeInstant != nil {
			res = __gong__marshallPointer(ident, "TimeInstant", timeprimitivepropertytype.TimeInstant.GongGetIdentifier(stage))
		} else {
			res = __gong__marshallPointer(ident, "TimeInstant", "nil")
		}
	case "TimePeriod":
		if timeprimitivepropertytype.TimePeriod != nil {
			res = __gong__marshallPointer(ident, "TimePeriod", timeprimitivepropertytype.TimePeriod.GongGetIdentifier(stage))
		} else {
			res = __gong__marshallPointer(ident, "TimePeriod", "nil")
		}
	default:
		log.Panicf("Unknown field %s for Gongstruct TimePrimitivePropertyType", fieldName)
	}
	return
}

// insertion point for marshall all fields methods
func (abstractfeaturemembertype *AbstractFeatureMemberType) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(abstractfeaturemembertype.GongMarshallField(stage, "Name"))
	}
	initRes = initializerStatements.String()
	ptrRes = pointersInitializesStatements.String()
	return
}
func (abstractfeaturetype *AbstractFeatureType) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(abstractfeaturetype.GongMarshallField(stage, "Name"))
		initializerStatements.WriteString(abstractfeaturetype.GongMarshallField(stage, "Id"))
		initializerStatements.WriteString(abstractfeaturetype.GongMarshallField(stage, "Description"))
		pointersInitializesStatements.WriteString(abstractfeaturetype.GongMarshallField(stage, "Identifier"))
		pointersInitializesStatements.WriteString(abstractfeaturetype.GongMarshallField(stage, "NameXSD"))
		pointersInitializesStatements.WriteString(abstractfeaturetype.GongMarshallField(stage, "BoundedBy"))
	}
	initRes = initializerStatements.String()
	ptrRes = pointersInitializesStatements.String()
	return
}
func (abstractgmltype *AbstractGMLType) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(abstractgmltype.GongMarshallField(stage, "Name"))
		initializerStatements.WriteString(abstractgmltype.GongMarshallField(stage, "Id"))
		initializerStatements.WriteString(abstractgmltype.GongMarshallField(stage, "Description"))
		pointersInitializesStatements.WriteString(abstractgmltype.GongMarshallField(stage, "Identifier"))
		pointersInitializesStatements.WriteString(abstractgmltype.GongMarshallField(stage, "NameXSD"))
	}
	initRes = initializerStatements.String()
	ptrRes = pointersInitializesStatements.String()
	return
}
func (abstractgeometricprimitivetype *AbstractGeometricPrimitiveType) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(abstractgeometricprimitivetype.GongMarshallField(stage, "Name"))
		initializerStatements.WriteString(abstractgeometricprimitivetype.GongMarshallField(stage, "Id"))
		initializerStatements.WriteString(abstractgeometricprimitivetype.GongMarshallField(stage, "Description"))
		pointersInitializesStatements.WriteString(abstractgeometricprimitivetype.GongMarshallField(stage, "Identifier"))
		pointersInitializesStatements.WriteString(abstractgeometricprimitivetype.GongMarshallField(stage, "NameXSD"))
		initializerStatements.WriteString(abstractgeometricprimitivetype.GongMarshallField(stage, "SrsName"))
	}
	initRes = initializerStatements.String()
	ptrRes = pointersInitializesStatements.String()
	return
}
func (abstractgeometrytype *AbstractGeometryType) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(abstractgeometrytype.GongMarshallField(stage, "Name"))
		initializerStatements.WriteString(abstractgeometrytype.GongMarshallField(stage, "Id"))
		initializerStatements.WriteString(abstractgeometrytype.GongMarshallField(stage, "Description"))
		pointersInitializesStatements.WriteString(abstractgeometrytype.GongMarshallField(stage, "Identifier"))
		pointersInitializesStatements.WriteString(abstractgeometrytype.GongMarshallField(stage, "NameXSD"))
		initializerStatements.WriteString(abstractgeometrytype.GongMarshallField(stage, "SrsName"))
	}
	initRes = initializerStatements.String()
	ptrRes = pointersInitializesStatements.String()
	return
}
func (abstractmetadatapropertytype *AbstractMetadataPropertyType) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(abstractmetadatapropertytype.GongMarshallField(stage, "Name"))
		initializerStatements.WriteString(abstractmetadatapropertytype.GongMarshallField(stage, "Owns"))
	}
	initRes = initializerStatements.String()
	ptrRes = pointersInitializesStatements.String()
	return
}
func (abstracttimegeometricprimitivetype *AbstractTimeGeometricPrimitiveType) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(abstracttimegeometricprimitivetype.GongMarshallField(stage, "Name"))
		initializerStatements.WriteString(abstracttimegeometricprimitivetype.GongMarshallField(stage, "Id"))
		initializerStatements.WriteString(abstracttimegeometricprimitivetype.GongMarshallField(stage, "Description"))
		pointersInitializesStatements.WriteString(abstracttimegeometricprimitivetype.GongMarshallField(stage, "Identifier"))
		pointersInitializesStatements.WriteString(abstracttimegeometricprimitivetype.GongMarshallField(stage, "NameXSD"))
		initializerStatements.WriteString(abstracttimegeometricprimitivetype.GongMarshallField(stage, "Frame"))
	}
	initRes = initializerStatements.String()
	ptrRes = pointersInitializesStatements.String()
	return
}
func (abstracttimeobjecttype *AbstractTimeObjectType) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(abstracttimeobjecttype.GongMarshallField(stage, "Name"))
		initializerStatements.WriteString(abstracttimeobjecttype.GongMarshallField(stage, "Id"))
		initializerStatements.WriteString(abstracttimeobjecttype.GongMarshallField(stage, "Description"))
		pointersInitializesStatements.WriteString(abstracttimeobjecttype.GongMarshallField(stage, "Identifier"))
		pointersInitializesStatements.WriteString(abstracttimeobjecttype.GongMarshallField(stage, "NameXSD"))
	}
	initRes = initializerStatements.String()
	ptrRes = pointersInitializesStatements.String()
	return
}
func (abstracttimeprimitivetype *AbstractTimePrimitiveType) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(abstracttimeprimitivetype.GongMarshallField(stage, "Name"))
		initializerStatements.WriteString(abstracttimeprimitivetype.GongMarshallField(stage, "Id"))
		initializerStatements.WriteString(abstracttimeprimitivetype.GongMarshallField(stage, "Description"))
		pointersInitializesStatements.WriteString(abstracttimeprimitivetype.GongMarshallField(stage, "Identifier"))
		pointersInitializesStatements.WriteString(abstracttimeprimitivetype.GongMarshallField(stage, "NameXSD"))
	}
	initRes = initializerStatements.String()
	ptrRes = pointersInitializesStatements.String()
	return
}
func (abstracttimeslicetype *AbstractTimeSliceType) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(abstracttimeslicetype.GongMarshallField(stage, "Name"))
		initializerStatements.WriteString(abstracttimeslicetype.GongMarshallField(stage, "Id"))
		initializerStatements.WriteString(abstracttimeslicetype.GongMarshallField(stage, "Description"))
		pointersInitializesStatements.WriteString(abstracttimeslicetype.GongMarshallField(stage, "Identifier"))
		pointersInitializesStatements.WriteString(abstracttimeslicetype.GongMarshallField(stage, "NameXSD"))
		pointersInitializesStatements.WriteString(abstracttimeslicetype.GongMarshallField(stage, "ValidTime"))
	}
	initRes = initializerStatements.String()
	ptrRes = pointersInitializesStatements.String()
	return
}
func (boundingshapetype *BoundingShapeType) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(boundingshapetype.GongMarshallField(stage, "Name"))
		initializerStatements.WriteString(boundingshapetype.GongMarshallField(stage, "NilReason"))
		pointersInitializesStatements.WriteString(boundingshapetype.GongMarshallField(stage, "Envelope"))
	}
	initRes = initializerStatements.String()
	ptrRes = pointersInitializesStatements.String()
	return
}
func (codetype *CodeType) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(codetype.GongMarshallField(stage, "Name"))
		initializerStatements.WriteString(codetype.GongMarshallField(stage, "CodeSpace"))
		initializerStatements.WriteString(codetype.GongMarshallField(stage, "EnclosedText"))
	}
	initRes = initializerStatements.String()
	ptrRes = pointersInitializesStatements.String()
	return
}
func (codewithauthoritytype *CodeWithAuthorityType) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(codewithauthoritytype.GongMarshallField(stage, "Name"))
		initializerStatements.WriteString(codewithauthoritytype.GongMarshallField(stage, "CodeSpace"))
		initializerStatements.WriteString(codewithauthoritytype.GongMarshallField(stage, "EnclosedText"))
	}
	initRes = initializerStatements.String()
	ptrRes = pointersInitializesStatements.String()
	return
}
func (directpositiontype *DirectPositionType) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(directpositiontype.GongMarshallField(stage, "Name"))
		initializerStatements.WriteString(directpositiontype.GongMarshallField(stage, "SrsName"))
		initializerStatements.WriteString(directpositiontype.GongMarshallField(stage, "EnclosedText"))
	}
	initRes = initializerStatements.String()
	ptrRes = pointersInitializesStatements.String()
	return
}
func (dynamicfeaturetype *DynamicFeatureType) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(dynamicfeaturetype.GongMarshallField(stage, "Name"))
		initializerStatements.WriteString(dynamicfeaturetype.GongMarshallField(stage, "Id"))
		initializerStatements.WriteString(dynamicfeaturetype.GongMarshallField(stage, "Description"))
		pointersInitializesStatements.WriteString(dynamicfeaturetype.GongMarshallField(stage, "Identifier"))
		pointersInitializesStatements.WriteString(dynamicfeaturetype.GongMarshallField(stage, "NameXSD"))
		pointersInitializesStatements.WriteString(dynamicfeaturetype.GongMarshallField(stage, "BoundedBy"))
	}
	initRes = initializerStatements.String()
	ptrRes = pointersInitializesStatements.String()
	return
}
func (envelopetype *EnvelopeType) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(envelopetype.GongMarshallField(stage, "Name"))
		initializerStatements.WriteString(envelopetype.GongMarshallField(stage, "SrsName"))
		pointersInitializesStatements.WriteString(envelopetype.GongMarshallField(stage, "LowerCorner"))
		pointersInitializesStatements.WriteString(envelopetype.GongMarshallField(stage, "UpperCorner"))
	}
	initRes = initializerStatements.String()
	ptrRes = pointersInitializesStatements.String()
	return
}
func (pointtype *PointType) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(pointtype.GongMarshallField(stage, "Name"))
		initializerStatements.WriteString(pointtype.GongMarshallField(stage, "Id"))
		initializerStatements.WriteString(pointtype.GongMarshallField(stage, "Description"))
		pointersInitializesStatements.WriteString(pointtype.GongMarshallField(stage, "Identifier"))
		pointersInitializesStatements.WriteString(pointtype.GongMarshallField(stage, "NameXSD"))
		initializerStatements.WriteString(pointtype.GongMarshallField(stage, "SrsName"))
		pointersInitializesStatements.WriteString(pointtype.GongMarshallField(stage, "Pos"))
	}
	initRes = initializerStatements.String()
	ptrRes = pointersInitializesStatements.String()
	return
}
func (timeinstanttype *TimeInstantType) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(timeinstanttype.GongMarshallField(stage, "Name"))
		initializerStatements.WriteString(timeinstanttype.GongMarshallField(stage, "Id"))
		initializerStatements.WriteString(timeinstanttype.GongMarshallField(stage, "Description"))
		pointersInitializesStatements.WriteString(timeinstanttype.GongMarshallField(stage, "Identifier"))
		pointersInitializesStatements.WriteString(timeinstanttype.GongMarshallField(stage, "NameXSD"))
		initializerStatements.WriteString(timeinstanttype.GongMarshallField(stage, "Frame"))
		pointersInitializesStatements.WriteString(timeinstanttype.GongMarshallField(stage, "TimePosition"))
	}
	initRes = initializerStatements.String()
	ptrRes = pointersInitializesStatements.String()
	return
}
func (timeperiodtype *TimePeriodType) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(timeperiodtype.GongMarshallField(stage, "Name"))
		initializerStatements.WriteString(timeperiodtype.GongMarshallField(stage, "Id"))
		initializerStatements.WriteString(timeperiodtype.GongMarshallField(stage, "Description"))
		pointersInitializesStatements.WriteString(timeperiodtype.GongMarshallField(stage, "Identifier"))
		pointersInitializesStatements.WriteString(timeperiodtype.GongMarshallField(stage, "NameXSD"))
		initializerStatements.WriteString(timeperiodtype.GongMarshallField(stage, "Frame"))
		pointersInitializesStatements.WriteString(timeperiodtype.GongMarshallField(stage, "BeginPosition"))
		pointersInitializesStatements.WriteString(timeperiodtype.GongMarshallField(stage, "EndPosition"))
	}
	initRes = initializerStatements.String()
	ptrRes = pointersInitializesStatements.String()
	return
}
func (timepositiontype *TimePositionType) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(timepositiontype.GongMarshallField(stage, "Name"))
		initializerStatements.WriteString(timepositiontype.GongMarshallField(stage, "Frame"))
		initializerStatements.WriteString(timepositiontype.GongMarshallField(stage, "IndeterminatePosition"))
		initializerStatements.WriteString(timepositiontype.GongMarshallField(stage, "EnclosedText"))
	}
	initRes = initializerStatements.String()
	ptrRes = pointersInitializesStatements.String()
	return
}
func (timeprimitivepropertytype *TimePrimitivePropertyType) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(timeprimitivepropertytype.GongMarshallField(stage, "Name"))
		initializerStatements.WriteString(timeprimitivepropertytype.GongMarshallField(stage, "NilReason"))
		initializerStatements.WriteString(timeprimitivepropertytype.GongMarshallField(stage, "Type_XSD"))
		initializerStatements.WriteString(timeprimitivepropertytype.GongMarshallField(stage, "Href"))
		initializerStatements.WriteString(timeprimitivepropertytype.GongMarshallField(stage, "Role"))
		initializerStatements.WriteString(timeprimitivepropertytype.GongMarshallField(stage, "Arcrole"))
		initializerStatements.WriteString(timeprimitivepropertytype.GongMarshallField(stage, "Title"))
		initializerStatements.WriteString(timeprimitivepropertytype.GongMarshallField(stage, "Show"))
		initializerStatements.WriteString(timeprimitivepropertytype.GongMarshallField(stage, "Actuate"))
		pointersInitializesStatements.WriteString(timeprimitivepropertytype.GongMarshallField(stage, "TimeInstant"))
		pointersInitializesStatements.WriteString(timeprimitivepropertytype.GongMarshallField(stage, "TimePeriod"))
	}
	initRes = initializerStatements.String()
	ptrRes = pointersInitializesStatements.String()
	return
}
