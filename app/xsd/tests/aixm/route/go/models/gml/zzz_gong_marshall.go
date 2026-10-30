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
	modelsPackageName := "github.com/fullstack-lang/gong/app/xsd/tests/aixm/route/go/models/gml"
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
		for _, abstractcurvesegmenttype := range __gong__sortStageSetInstances(stage.AbstractCurveSegmentTypes, stage.AbstractCurveSegmentType_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			abstractcurvesegmenttypeIdent := abstractcurvesegmenttype.GongGetIdentifier(stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&gml.AbstractCurveSegmentType{Name: %s}).Stage(stage)", abstractcurvesegmenttypeIdent, __gong__toRawStringLiteral(abstractcurvesegmenttype.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", abstractcurvesegmenttypeIdent, __gong__toRawStringLiteral(abstractcurvesegmenttype.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.NumDerivativesAtStart = %d", abstractcurvesegmenttypeIdent, abstractcurvesegmenttype.NumDerivativesAtStart))
			values.WriteString(fmt.Sprintf("\n\t%s.NumDerivativesAtEnd = %d", abstractcurvesegmenttypeIdent, abstractcurvesegmenttype.NumDerivativesAtEnd))
			values.WriteString(fmt.Sprintf("\n\t%s.NumDerivativeInterior = %d", abstractcurvesegmenttypeIdent, abstractcurvesegmenttype.NumDerivativeInterior))
		}
	}
	if stage != nil {
		for _, abstractcurvetype := range __gong__sortStageSetInstances(stage.AbstractCurveTypes, stage.AbstractCurveType_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			abstractcurvetypeIdent := abstractcurvetype.GongGetIdentifier(stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&gml.AbstractCurveType{Name: %s}).Stage(stage)", abstractcurvetypeIdent, __gong__toRawStringLiteral(abstractcurvetype.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", abstractcurvetypeIdent, __gong__toRawStringLiteral(abstractcurvetype.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Id = %s", abstractcurvetypeIdent, __gong__toRawStringLiteral(abstractcurvetype.Id)))
			values.WriteString(fmt.Sprintf("\n\t%s.Description = %s", abstractcurvetypeIdent, __gong__toRawStringLiteral(abstractcurvetype.Description)))
			values.WriteString(fmt.Sprintf("\n\t%s.SrsName = %s", abstractcurvetypeIdent, __gong__toRawStringLiteral(abstractcurvetype.SrsName)))
			if abstractcurvetype.Identifier != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := abstractcurvetype.Identifier.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Identifier = %s", abstractcurvetypeIdent, targetIdent))
			}
			for _, elem := range abstractcurvetype.NameXSD {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := elem.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.NameXSD = append(%s.NameXSD, %s)", abstractcurvetypeIdent, abstractcurvetypeIdent, targetIdent))
			}
		}
	}
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
		for _, abstractringpropertytype := range __gong__sortStageSetInstances(stage.AbstractRingPropertyTypes, stage.AbstractRingPropertyType_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			abstractringpropertytypeIdent := abstractringpropertytype.GongGetIdentifier(stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&gml.AbstractRingPropertyType{Name: %s}).Stage(stage)", abstractringpropertytypeIdent, __gong__toRawStringLiteral(abstractringpropertytype.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", abstractringpropertytypeIdent, __gong__toRawStringLiteral(abstractringpropertytype.Name)))
			if abstractringpropertytype.Ring != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := abstractringpropertytype.Ring.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Ring = %s", abstractringpropertytypeIdent, targetIdent))
			}
		}
	}
	if stage != nil {
		for _, abstractringtype := range __gong__sortStageSetInstances(stage.AbstractRingTypes, stage.AbstractRingType_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			abstractringtypeIdent := abstractringtype.GongGetIdentifier(stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&gml.AbstractRingType{Name: %s}).Stage(stage)", abstractringtypeIdent, __gong__toRawStringLiteral(abstractringtype.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", abstractringtypeIdent, __gong__toRawStringLiteral(abstractringtype.Name)))
		}
	}
	if stage != nil {
		for _, abstractsurfacepatchtype := range __gong__sortStageSetInstances(stage.AbstractSurfacePatchTypes, stage.AbstractSurfacePatchType_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			abstractsurfacepatchtypeIdent := abstractsurfacepatchtype.GongGetIdentifier(stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&gml.AbstractSurfacePatchType{Name: %s}).Stage(stage)", abstractsurfacepatchtypeIdent, __gong__toRawStringLiteral(abstractsurfacepatchtype.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", abstractsurfacepatchtypeIdent, __gong__toRawStringLiteral(abstractsurfacepatchtype.Name)))
		}
	}
	if stage != nil {
		for _, abstractsurfacetype := range __gong__sortStageSetInstances(stage.AbstractSurfaceTypes, stage.AbstractSurfaceType_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			abstractsurfacetypeIdent := abstractsurfacetype.GongGetIdentifier(stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&gml.AbstractSurfaceType{Name: %s}).Stage(stage)", abstractsurfacetypeIdent, __gong__toRawStringLiteral(abstractsurfacetype.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", abstractsurfacetypeIdent, __gong__toRawStringLiteral(abstractsurfacetype.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Id = %s", abstractsurfacetypeIdent, __gong__toRawStringLiteral(abstractsurfacetype.Id)))
			values.WriteString(fmt.Sprintf("\n\t%s.Description = %s", abstractsurfacetypeIdent, __gong__toRawStringLiteral(abstractsurfacetype.Description)))
			values.WriteString(fmt.Sprintf("\n\t%s.SrsName = %s", abstractsurfacetypeIdent, __gong__toRawStringLiteral(abstractsurfacetype.SrsName)))
			if abstractsurfacetype.Identifier != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := abstractsurfacetype.Identifier.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Identifier = %s", abstractsurfacetypeIdent, targetIdent))
			}
			for _, elem := range abstractsurfacetype.NameXSD {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := elem.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.NameXSD = append(%s.NameXSD, %s)", abstractsurfacetypeIdent, abstractsurfacetypeIdent, targetIdent))
			}
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
		for _, angletype := range __gong__sortStageSetInstances(stage.AngleTypes, stage.AngleType_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			angletypeIdent := angletype.GongGetIdentifier(stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&gml.AngleType{Name: %s}).Stage(stage)", angletypeIdent, __gong__toRawStringLiteral(angletype.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", angletypeIdent, __gong__toRawStringLiteral(angletype.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.EnclosedText = %s", angletypeIdent, __gong__toRawStringLiteral(angletype.EnclosedText)))
		}
	}
	if stage != nil {
		for _, arcbycenterpointtype := range __gong__sortStageSetInstances(stage.ArcByCenterPointTypes, stage.ArcByCenterPointType_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			arcbycenterpointtypeIdent := arcbycenterpointtype.GongGetIdentifier(stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&gml.ArcByCenterPointType{Name: %s}).Stage(stage)", arcbycenterpointtypeIdent, __gong__toRawStringLiteral(arcbycenterpointtype.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
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
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := arcbycenterpointtype.Pos.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Pos = %s", arcbycenterpointtypeIdent, targetIdent))
			}
			if arcbycenterpointtype.PointProperty != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := arcbycenterpointtype.PointProperty.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.PointProperty = %s", arcbycenterpointtypeIdent, targetIdent))
			}
			if arcbycenterpointtype.PosList != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := arcbycenterpointtype.PosList.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.PosList = %s", arcbycenterpointtypeIdent, targetIdent))
			}
			if arcbycenterpointtype.Radius != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := arcbycenterpointtype.Radius.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Radius = %s", arcbycenterpointtypeIdent, targetIdent))
			}
			if arcbycenterpointtype.StartAngle != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := arcbycenterpointtype.StartAngle.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.StartAngle = %s", arcbycenterpointtypeIdent, targetIdent))
			}
			if arcbycenterpointtype.EndAngle != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := arcbycenterpointtype.EndAngle.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.EndAngle = %s", arcbycenterpointtypeIdent, targetIdent))
			}
		}
	}
	if stage != nil {
		for _, arcstringtype := range __gong__sortStageSetInstances(stage.ArcStringTypes, stage.ArcStringType_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			arcstringtypeIdent := arcstringtype.GongGetIdentifier(stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&gml.ArcStringType{Name: %s}).Stage(stage)", arcstringtypeIdent, __gong__toRawStringLiteral(arcstringtype.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
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
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := elem.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Pos = append(%s.Pos, %s)", arcstringtypeIdent, arcstringtypeIdent, targetIdent))
			}
			for _, elem := range arcstringtype.PointProperty {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := elem.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.PointProperty = append(%s.PointProperty, %s)", arcstringtypeIdent, arcstringtypeIdent, targetIdent))
			}
			if arcstringtype.PosList != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := arcstringtype.PosList.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.PosList = %s", arcstringtypeIdent, targetIdent))
			}
		}
	}
	if stage != nil {
		for _, arctype := range __gong__sortStageSetInstances(stage.ArcTypes, stage.ArcType_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			arctypeIdent := arctype.GongGetIdentifier(stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&gml.ArcType{Name: %s}).Stage(stage)", arctypeIdent, __gong__toRawStringLiteral(arctype.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
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
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := arctype.Pos.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Pos = %s", arctypeIdent, targetIdent))
			}
			if arctype.PointProperty != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := arctype.PointProperty.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.PointProperty = %s", arctypeIdent, targetIdent))
			}
			if arctype.PosList != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := arctype.PosList.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.PosList = %s", arctypeIdent, targetIdent))
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
		for _, circlebycenterpointtype := range __gong__sortStageSetInstances(stage.CircleByCenterPointTypes, stage.CircleByCenterPointType_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			circlebycenterpointtypeIdent := circlebycenterpointtype.GongGetIdentifier(stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&gml.CircleByCenterPointType{Name: %s}).Stage(stage)", circlebycenterpointtypeIdent, __gong__toRawStringLiteral(circlebycenterpointtype.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
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
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := circlebycenterpointtype.Pos.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Pos = %s", circlebycenterpointtypeIdent, targetIdent))
			}
			if circlebycenterpointtype.PointProperty != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := circlebycenterpointtype.PointProperty.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.PointProperty = %s", circlebycenterpointtypeIdent, targetIdent))
			}
			if circlebycenterpointtype.PosList != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := circlebycenterpointtype.PosList.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.PosList = %s", circlebycenterpointtypeIdent, targetIdent))
			}
			if circlebycenterpointtype.Radius != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := circlebycenterpointtype.Radius.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Radius = %s", circlebycenterpointtypeIdent, targetIdent))
			}
		}
	}
	if stage != nil {
		for _, circletype := range __gong__sortStageSetInstances(stage.CircleTypes, stage.CircleType_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			circletypeIdent := circletype.GongGetIdentifier(stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&gml.CircleType{Name: %s}).Stage(stage)", circletypeIdent, __gong__toRawStringLiteral(circletype.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
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
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := elem.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Pos = append(%s.Pos, %s)", circletypeIdent, circletypeIdent, targetIdent))
			}
			for _, elem := range circletype.PointProperty {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := elem.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.PointProperty = append(%s.PointProperty, %s)", circletypeIdent, circletypeIdent, targetIdent))
			}
			if circletype.PosList != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := circletype.PosList.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.PosList = %s", circletypeIdent, targetIdent))
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
		for _, compositecurvetype := range __gong__sortStageSetInstances(stage.CompositeCurveTypes, stage.CompositeCurveType_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			compositecurvetypeIdent := compositecurvetype.GongGetIdentifier(stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&gml.CompositeCurveType{Name: %s}).Stage(stage)", compositecurvetypeIdent, __gong__toRawStringLiteral(compositecurvetype.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", compositecurvetypeIdent, __gong__toRawStringLiteral(compositecurvetype.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Id = %s", compositecurvetypeIdent, __gong__toRawStringLiteral(compositecurvetype.Id)))
			values.WriteString(fmt.Sprintf("\n\t%s.Description = %s", compositecurvetypeIdent, __gong__toRawStringLiteral(compositecurvetype.Description)))
			values.WriteString(fmt.Sprintf("\n\t%s.SrsName = %s", compositecurvetypeIdent, __gong__toRawStringLiteral(compositecurvetype.SrsName)))
			values.WriteString(fmt.Sprintf("\n\t%s.AggregationType = %s", compositecurvetypeIdent, __gong__toRawStringLiteral(compositecurvetype.AggregationType)))
			if compositecurvetype.Identifier != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := compositecurvetype.Identifier.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Identifier = %s", compositecurvetypeIdent, targetIdent))
			}
			for _, elem := range compositecurvetype.NameXSD {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := elem.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.NameXSD = append(%s.NameXSD, %s)", compositecurvetypeIdent, compositecurvetypeIdent, targetIdent))
			}
			for _, elem := range compositecurvetype.CurveMember {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := elem.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.CurveMember = append(%s.CurveMember, %s)", compositecurvetypeIdent, compositecurvetypeIdent, targetIdent))
			}
		}
	}
	if stage != nil {
		for _, curvepropertytype := range __gong__sortStageSetInstances(stage.CurvePropertyTypes, stage.CurvePropertyType_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			curvepropertytypeIdent := curvepropertytype.GongGetIdentifier(stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&gml.CurvePropertyType{Name: %s}).Stage(stage)", curvepropertytypeIdent, __gong__toRawStringLiteral(curvepropertytype.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
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
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := curvepropertytype.Curve.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Curve = %s", curvepropertytypeIdent, targetIdent))
			}
			if curvepropertytype.OrientableCurve != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := curvepropertytype.OrientableCurve.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.OrientableCurve = %s", curvepropertytypeIdent, targetIdent))
			}
			if curvepropertytype.CompositeCurve != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := curvepropertytype.CompositeCurve.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.CompositeCurve = %s", curvepropertytypeIdent, targetIdent))
			}
		}
	}
	if stage != nil {
		for _, curvesegmentarraypropertytype := range __gong__sortStageSetInstances(stage.CurveSegmentArrayPropertyTypes, stage.CurveSegmentArrayPropertyType_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			curvesegmentarraypropertytypeIdent := curvesegmentarraypropertytype.GongGetIdentifier(stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&gml.CurveSegmentArrayPropertyType{Name: %s}).Stage(stage)", curvesegmentarraypropertytypeIdent, __gong__toRawStringLiteral(curvesegmentarraypropertytype.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", curvesegmentarraypropertytypeIdent, __gong__toRawStringLiteral(curvesegmentarraypropertytype.Name)))
			for _, elem := range curvesegmentarraypropertytype.LineStringSegment {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := elem.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.LineStringSegment = append(%s.LineStringSegment, %s)", curvesegmentarraypropertytypeIdent, curvesegmentarraypropertytypeIdent, targetIdent))
			}
			for _, elem := range curvesegmentarraypropertytype.ArcString {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := elem.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.ArcString = append(%s.ArcString, %s)", curvesegmentarraypropertytypeIdent, curvesegmentarraypropertytypeIdent, targetIdent))
			}
			for _, elem := range curvesegmentarraypropertytype.Arc {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := elem.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Arc = append(%s.Arc, %s)", curvesegmentarraypropertytypeIdent, curvesegmentarraypropertytypeIdent, targetIdent))
			}
			for _, elem := range curvesegmentarraypropertytype.Circle {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := elem.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Circle = append(%s.Circle, %s)", curvesegmentarraypropertytypeIdent, curvesegmentarraypropertytypeIdent, targetIdent))
			}
			for _, elem := range curvesegmentarraypropertytype.ArcByCenterPoint {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := elem.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.ArcByCenterPoint = append(%s.ArcByCenterPoint, %s)", curvesegmentarraypropertytypeIdent, curvesegmentarraypropertytypeIdent, targetIdent))
			}
			for _, elem := range curvesegmentarraypropertytype.CircleByCenterPoint {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := elem.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.CircleByCenterPoint = append(%s.CircleByCenterPoint, %s)", curvesegmentarraypropertytypeIdent, curvesegmentarraypropertytypeIdent, targetIdent))
			}
			for _, elem := range curvesegmentarraypropertytype.GeodesicString {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := elem.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.GeodesicString = append(%s.GeodesicString, %s)", curvesegmentarraypropertytypeIdent, curvesegmentarraypropertytypeIdent, targetIdent))
			}
			for _, elem := range curvesegmentarraypropertytype.Geodesic {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := elem.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Geodesic = append(%s.Geodesic, %s)", curvesegmentarraypropertytypeIdent, curvesegmentarraypropertytypeIdent, targetIdent))
			}
		}
	}
	if stage != nil {
		for _, curvetype := range __gong__sortStageSetInstances(stage.CurveTypes, stage.CurveType_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			curvetypeIdent := curvetype.GongGetIdentifier(stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&gml.CurveType{Name: %s}).Stage(stage)", curvetypeIdent, __gong__toRawStringLiteral(curvetype.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", curvetypeIdent, __gong__toRawStringLiteral(curvetype.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Id = %s", curvetypeIdent, __gong__toRawStringLiteral(curvetype.Id)))
			values.WriteString(fmt.Sprintf("\n\t%s.Description = %s", curvetypeIdent, __gong__toRawStringLiteral(curvetype.Description)))
			values.WriteString(fmt.Sprintf("\n\t%s.SrsName = %s", curvetypeIdent, __gong__toRawStringLiteral(curvetype.SrsName)))
			if curvetype.Identifier != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := curvetype.Identifier.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Identifier = %s", curvetypeIdent, targetIdent))
			}
			for _, elem := range curvetype.NameXSD {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := elem.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.NameXSD = append(%s.NameXSD, %s)", curvetypeIdent, curvetypeIdent, targetIdent))
			}
			if curvetype.Segments != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := curvetype.Segments.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Segments = %s", curvetypeIdent, targetIdent))
			}
		}
	}
	if stage != nil {
		for _, directpositionlisttype := range __gong__sortStageSetInstances(stage.DirectPositionListTypes, stage.DirectPositionListType_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			directpositionlisttypeIdent := directpositionlisttype.GongGetIdentifier(stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&gml.DirectPositionListType{Name: %s}).Stage(stage)", directpositionlisttypeIdent, __gong__toRawStringLiteral(directpositionlisttype.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", directpositionlisttypeIdent, __gong__toRawStringLiteral(directpositionlisttype.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.SrsName = %s", directpositionlisttypeIdent, __gong__toRawStringLiteral(directpositionlisttype.SrsName)))
			values.WriteString(fmt.Sprintf("\n\t%s.EnclosedText = %s", directpositionlisttypeIdent, __gong__toRawStringLiteral(directpositionlisttype.EnclosedText)))
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
		for _, geodesicstringtype := range __gong__sortStageSetInstances(stage.GeodesicStringTypes, stage.GeodesicStringType_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			geodesicstringtypeIdent := geodesicstringtype.GongGetIdentifier(stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&gml.GeodesicStringType{Name: %s}).Stage(stage)", geodesicstringtypeIdent, __gong__toRawStringLiteral(geodesicstringtype.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
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
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := geodesicstringtype.PosList.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.PosList = %s", geodesicstringtypeIdent, targetIdent))
			}
			if geodesicstringtype.Pos != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := geodesicstringtype.Pos.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Pos = %s", geodesicstringtypeIdent, targetIdent))
			}
			if geodesicstringtype.PointProperty != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := geodesicstringtype.PointProperty.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.PointProperty = %s", geodesicstringtypeIdent, targetIdent))
			}
		}
	}
	if stage != nil {
		for _, geodesictype := range __gong__sortStageSetInstances(stage.GeodesicTypes, stage.GeodesicType_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			geodesictypeIdent := geodesictype.GongGetIdentifier(stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&gml.GeodesicType{Name: %s}).Stage(stage)", geodesictypeIdent, __gong__toRawStringLiteral(geodesictype.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
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
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := geodesictype.PosList.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.PosList = %s", geodesictypeIdent, targetIdent))
			}
			if geodesictype.Pos != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := geodesictype.Pos.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Pos = %s", geodesictypeIdent, targetIdent))
			}
			if geodesictype.PointProperty != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := geodesictype.PointProperty.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.PointProperty = %s", geodesictypeIdent, targetIdent))
			}
		}
	}
	if stage != nil {
		for _, lengthtype := range __gong__sortStageSetInstances(stage.LengthTypes, stage.LengthType_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			lengthtypeIdent := lengthtype.GongGetIdentifier(stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&gml.LengthType{Name: %s}).Stage(stage)", lengthtypeIdent, __gong__toRawStringLiteral(lengthtype.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", lengthtypeIdent, __gong__toRawStringLiteral(lengthtype.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.EnclosedText = %s", lengthtypeIdent, __gong__toRawStringLiteral(lengthtype.EnclosedText)))
		}
	}
	if stage != nil {
		for _, linestringsegmenttype := range __gong__sortStageSetInstances(stage.LineStringSegmentTypes, stage.LineStringSegmentType_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			linestringsegmenttypeIdent := linestringsegmenttype.GongGetIdentifier(stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&gml.LineStringSegmentType{Name: %s}).Stage(stage)", linestringsegmenttypeIdent, __gong__toRawStringLiteral(linestringsegmenttype.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
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
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := elem.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Pos = append(%s.Pos, %s)", linestringsegmenttypeIdent, linestringsegmenttypeIdent, targetIdent))
			}
			for _, elem := range linestringsegmenttype.PointProperty {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := elem.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.PointProperty = append(%s.PointProperty, %s)", linestringsegmenttypeIdent, linestringsegmenttypeIdent, targetIdent))
			}
			if linestringsegmenttype.PosList != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := linestringsegmenttype.PosList.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.PosList = %s", linestringsegmenttypeIdent, targetIdent))
			}
		}
	}
	if stage != nil {
		for _, measuretype := range __gong__sortStageSetInstances(stage.MeasureTypes, stage.MeasureType_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			measuretypeIdent := measuretype.GongGetIdentifier(stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&gml.MeasureType{Name: %s}).Stage(stage)", measuretypeIdent, __gong__toRawStringLiteral(measuretype.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", measuretypeIdent, __gong__toRawStringLiteral(measuretype.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Uom = %s", measuretypeIdent, __gong__toRawStringLiteral(measuretype.Uom)))
			values.WriteString(fmt.Sprintf("\n\t%s.EnclosedText = %f", measuretypeIdent, measuretype.EnclosedText))
		}
	}
	if stage != nil {
		for _, orientablecurvetype := range __gong__sortStageSetInstances(stage.OrientableCurveTypes, stage.OrientableCurveType_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			orientablecurvetypeIdent := orientablecurvetype.GongGetIdentifier(stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&gml.OrientableCurveType{Name: %s}).Stage(stage)", orientablecurvetypeIdent, __gong__toRawStringLiteral(orientablecurvetype.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
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
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := orientablecurvetype.Identifier.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Identifier = %s", orientablecurvetypeIdent, targetIdent))
			}
			for _, elem := range orientablecurvetype.NameXSD {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := elem.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.NameXSD = append(%s.NameXSD, %s)", orientablecurvetypeIdent, orientablecurvetypeIdent, targetIdent))
			}
			if orientablecurvetype.BaseCurve != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := orientablecurvetype.BaseCurve.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.BaseCurve = %s", orientablecurvetypeIdent, targetIdent))
			}
		}
	}
	if stage != nil {
		for _, pointpropertytype := range __gong__sortStageSetInstances(stage.PointPropertyTypes, stage.PointPropertyType_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			pointpropertytypeIdent := pointpropertytype.GongGetIdentifier(stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&gml.PointPropertyType{Name: %s}).Stage(stage)", pointpropertytypeIdent, __gong__toRawStringLiteral(pointpropertytype.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
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
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := pointpropertytype.Point.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Point = %s", pointpropertytypeIdent, targetIdent))
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
		for _, polygonpatchtype := range __gong__sortStageSetInstances(stage.PolygonPatchTypes, stage.PolygonPatchType_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			polygonpatchtypeIdent := polygonpatchtype.GongGetIdentifier(stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&gml.PolygonPatchType{Name: %s}).Stage(stage)", polygonpatchtypeIdent, __gong__toRawStringLiteral(polygonpatchtype.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", polygonpatchtypeIdent, __gong__toRawStringLiteral(polygonpatchtype.Name)))
			if polygonpatchtype.Interpolation.ToCodeString() != "" {
				values.WriteString(fmt.Sprintf("\n\t%s.Interpolation = gml.%s", polygonpatchtypeIdent, polygonpatchtype.Interpolation.ToCodeString()))
			} else {
				values.WriteString(fmt.Sprintf("\n\t%s.Interpolation = \"\"", polygonpatchtypeIdent))
			}
			if polygonpatchtype.Exterior != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := polygonpatchtype.Exterior.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Exterior = %s", polygonpatchtypeIdent, targetIdent))
			}
		}
	}
	if stage != nil {
		for _, ringtype := range __gong__sortStageSetInstances(stage.RingTypes, stage.RingType_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			ringtypeIdent := ringtype.GongGetIdentifier(stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&gml.RingType{Name: %s}).Stage(stage)", ringtypeIdent, __gong__toRawStringLiteral(ringtype.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", ringtypeIdent, __gong__toRawStringLiteral(ringtype.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.AggregationType = %s", ringtypeIdent, __gong__toRawStringLiteral(ringtype.AggregationType)))
			for _, elem := range ringtype.CurveMember {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := elem.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.CurveMember = append(%s.CurveMember, %s)", ringtypeIdent, ringtypeIdent, targetIdent))
			}
		}
	}
	if stage != nil {
		for _, surfacepatcharraypropertytype := range __gong__sortStageSetInstances(stage.SurfacePatchArrayPropertyTypes, stage.SurfacePatchArrayPropertyType_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			surfacepatcharraypropertytypeIdent := surfacepatcharraypropertytype.GongGetIdentifier(stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&gml.SurfacePatchArrayPropertyType{Name: %s}).Stage(stage)", surfacepatcharraypropertytypeIdent, __gong__toRawStringLiteral(surfacepatcharraypropertytype.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", surfacepatcharraypropertytypeIdent, __gong__toRawStringLiteral(surfacepatcharraypropertytype.Name)))
			for _, elem := range surfacepatcharraypropertytype.PolygonPatch {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := elem.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.PolygonPatch = append(%s.PolygonPatch, %s)", surfacepatcharraypropertytypeIdent, surfacepatcharraypropertytypeIdent, targetIdent))
			}
		}
	}
	if stage != nil {
		for _, surfacetype := range __gong__sortStageSetInstances(stage.SurfaceTypes, stage.SurfaceType_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			surfacetypeIdent := surfacetype.GongGetIdentifier(stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&gml.SurfaceType{Name: %s}).Stage(stage)", surfacetypeIdent, __gong__toRawStringLiteral(surfacetype.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", surfacetypeIdent, __gong__toRawStringLiteral(surfacetype.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Id = %s", surfacetypeIdent, __gong__toRawStringLiteral(surfacetype.Id)))
			values.WriteString(fmt.Sprintf("\n\t%s.Description = %s", surfacetypeIdent, __gong__toRawStringLiteral(surfacetype.Description)))
			values.WriteString(fmt.Sprintf("\n\t%s.SrsName = %s", surfacetypeIdent, __gong__toRawStringLiteral(surfacetype.SrsName)))
			if surfacetype.Identifier != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := surfacetype.Identifier.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Identifier = %s", surfacetypeIdent, targetIdent))
			}
			for _, elem := range surfacetype.NameXSD {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := elem.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.NameXSD = append(%s.NameXSD, %s)", surfacetypeIdent, surfacetypeIdent, targetIdent))
			}
			if surfacetype.Patches != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := surfacetype.Patches.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Patches = %s", surfacetypeIdent, targetIdent))
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
go run github.com/fullstack-lang/gong/app/xsd/tests/aixm/route/go/models/gml/probe/cmd/probe@%s stage.go
*/
package %s

import (
	"slices"
	"time"

	"github.com/fullstack-lang/gong/app/xsd/tests/aixm/route/go/models/gml"%s
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

	if modelsPackageName != "" && modelsPackageName != "github.com/fullstack-lang/gong/app/xsd/tests/aixm/route/go/models/gml" {
		res = strings.ReplaceAll(res, "\"github.com/fullstack-lang/gong/app/xsd/tests/aixm/route/go/models/gml\"", fmt.Sprintf("%q", modelsPackageName))
	}

	return res, nil
}

// insertion point for marshall field methods
func (abstractcurvesegmenttype *AbstractCurveSegmentType) GongMarshallField(stage *Stage, fieldName string) (res string) {
	ident := abstractcurvesegmenttype.GongGetIdentifier(stage)
	_ = ident
	switch fieldName {
	case "Name":
		res = __gong__marshallString(ident, "Name", abstractcurvesegmenttype.Name)
	case "NumDerivativesAtStart":
		res = __gong__marshallInt(ident, "NumDerivativesAtStart", abstractcurvesegmenttype.NumDerivativesAtStart)
	case "NumDerivativesAtEnd":
		res = __gong__marshallInt(ident, "NumDerivativesAtEnd", abstractcurvesegmenttype.NumDerivativesAtEnd)
	case "NumDerivativeInterior":
		res = __gong__marshallInt(ident, "NumDerivativeInterior", abstractcurvesegmenttype.NumDerivativeInterior)

	default:
		log.Panicf("Unknown field %s for Gongstruct AbstractCurveSegmentType", fieldName)
	}
	return
}

func (abstractcurvetype *AbstractCurveType) GongMarshallField(stage *Stage, fieldName string) (res string) {
	ident := abstractcurvetype.GongGetIdentifier(stage)
	_ = ident
	switch fieldName {
	case "Name":
		res = __gong__marshallString(ident, "Name", abstractcurvetype.Name)
	case "Id":
		res = __gong__marshallString(ident, "Id", abstractcurvetype.Id)
	case "Description":
		res = __gong__marshallString(ident, "Description", abstractcurvetype.Description)
	case "SrsName":
		res = __gong__marshallString(ident, "SrsName", abstractcurvetype.SrsName)

	case "Identifier":
		if abstractcurvetype.Identifier != nil {
			res = __gong__marshallPointer(ident, "Identifier", abstractcurvetype.Identifier.GongGetIdentifier(stage))
		} else {
			res = __gong__marshallPointer(ident, "Identifier", "nil")
		}
	case "NameXSD":
		var sb strings.Builder
		for _, _codetype := range abstractcurvetype.NameXSD {
			sb.WriteString(__gong__marshallSliceOfPointers(ident, "NameXSD", _codetype.GongGetIdentifier(stage)))
		}
		res = sb.String()
	default:
		log.Panicf("Unknown field %s for Gongstruct AbstractCurveType", fieldName)
	}
	return
}

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

func (abstractringpropertytype *AbstractRingPropertyType) GongMarshallField(stage *Stage, fieldName string) (res string) {
	ident := abstractringpropertytype.GongGetIdentifier(stage)
	_ = ident
	switch fieldName {
	case "Name":
		res = __gong__marshallString(ident, "Name", abstractringpropertytype.Name)

	case "Ring":
		if abstractringpropertytype.Ring != nil {
			res = __gong__marshallPointer(ident, "Ring", abstractringpropertytype.Ring.GongGetIdentifier(stage))
		} else {
			res = __gong__marshallPointer(ident, "Ring", "nil")
		}
	default:
		log.Panicf("Unknown field %s for Gongstruct AbstractRingPropertyType", fieldName)
	}
	return
}

func (abstractringtype *AbstractRingType) GongMarshallField(stage *Stage, fieldName string) (res string) {
	ident := abstractringtype.GongGetIdentifier(stage)
	_ = ident
	switch fieldName {
	case "Name":
		res = __gong__marshallString(ident, "Name", abstractringtype.Name)

	default:
		log.Panicf("Unknown field %s for Gongstruct AbstractRingType", fieldName)
	}
	return
}

func (abstractsurfacepatchtype *AbstractSurfacePatchType) GongMarshallField(stage *Stage, fieldName string) (res string) {
	ident := abstractsurfacepatchtype.GongGetIdentifier(stage)
	_ = ident
	switch fieldName {
	case "Name":
		res = __gong__marshallString(ident, "Name", abstractsurfacepatchtype.Name)

	default:
		log.Panicf("Unknown field %s for Gongstruct AbstractSurfacePatchType", fieldName)
	}
	return
}

func (abstractsurfacetype *AbstractSurfaceType) GongMarshallField(stage *Stage, fieldName string) (res string) {
	ident := abstractsurfacetype.GongGetIdentifier(stage)
	_ = ident
	switch fieldName {
	case "Name":
		res = __gong__marshallString(ident, "Name", abstractsurfacetype.Name)
	case "Id":
		res = __gong__marshallString(ident, "Id", abstractsurfacetype.Id)
	case "Description":
		res = __gong__marshallString(ident, "Description", abstractsurfacetype.Description)
	case "SrsName":
		res = __gong__marshallString(ident, "SrsName", abstractsurfacetype.SrsName)

	case "Identifier":
		if abstractsurfacetype.Identifier != nil {
			res = __gong__marshallPointer(ident, "Identifier", abstractsurfacetype.Identifier.GongGetIdentifier(stage))
		} else {
			res = __gong__marshallPointer(ident, "Identifier", "nil")
		}
	case "NameXSD":
		var sb strings.Builder
		for _, _codetype := range abstractsurfacetype.NameXSD {
			sb.WriteString(__gong__marshallSliceOfPointers(ident, "NameXSD", _codetype.GongGetIdentifier(stage)))
		}
		res = sb.String()
	default:
		log.Panicf("Unknown field %s for Gongstruct AbstractSurfaceType", fieldName)
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

func (angletype *AngleType) GongMarshallField(stage *Stage, fieldName string) (res string) {
	ident := angletype.GongGetIdentifier(stage)
	_ = ident
	switch fieldName {
	case "Name":
		res = __gong__marshallString(ident, "Name", angletype.Name)
	case "EnclosedText":
		res = __gong__marshallString(ident, "EnclosedText", angletype.EnclosedText)

	default:
		log.Panicf("Unknown field %s for Gongstruct AngleType", fieldName)
	}
	return
}

func (arcbycenterpointtype *ArcByCenterPointType) GongMarshallField(stage *Stage, fieldName string) (res string) {
	ident := arcbycenterpointtype.GongGetIdentifier(stage)
	_ = ident
	switch fieldName {
	case "Name":
		res = __gong__marshallString(ident, "Name", arcbycenterpointtype.Name)
	case "NumDerivativesAtStart":
		res = __gong__marshallInt(ident, "NumDerivativesAtStart", arcbycenterpointtype.NumDerivativesAtStart)
	case "NumDerivativesAtEnd":
		res = __gong__marshallInt(ident, "NumDerivativesAtEnd", arcbycenterpointtype.NumDerivativesAtEnd)
	case "NumDerivativeInterior":
		res = __gong__marshallInt(ident, "NumDerivativeInterior", arcbycenterpointtype.NumDerivativeInterior)
	case "Interpolation":
		res = __gong__marshallEnumString(ident, "Interpolation", arcbycenterpointtype.Interpolation.ToCodeString())
	case "NumArc":
		res = __gong__marshallInt(ident, "NumArc", arcbycenterpointtype.NumArc)

	case "Pos":
		if arcbycenterpointtype.Pos != nil {
			res = __gong__marshallPointer(ident, "Pos", arcbycenterpointtype.Pos.GongGetIdentifier(stage))
		} else {
			res = __gong__marshallPointer(ident, "Pos", "nil")
		}
	case "PointProperty":
		if arcbycenterpointtype.PointProperty != nil {
			res = __gong__marshallPointer(ident, "PointProperty", arcbycenterpointtype.PointProperty.GongGetIdentifier(stage))
		} else {
			res = __gong__marshallPointer(ident, "PointProperty", "nil")
		}
	case "PosList":
		if arcbycenterpointtype.PosList != nil {
			res = __gong__marshallPointer(ident, "PosList", arcbycenterpointtype.PosList.GongGetIdentifier(stage))
		} else {
			res = __gong__marshallPointer(ident, "PosList", "nil")
		}
	case "Radius":
		if arcbycenterpointtype.Radius != nil {
			res = __gong__marshallPointer(ident, "Radius", arcbycenterpointtype.Radius.GongGetIdentifier(stage))
		} else {
			res = __gong__marshallPointer(ident, "Radius", "nil")
		}
	case "StartAngle":
		if arcbycenterpointtype.StartAngle != nil {
			res = __gong__marshallPointer(ident, "StartAngle", arcbycenterpointtype.StartAngle.GongGetIdentifier(stage))
		} else {
			res = __gong__marshallPointer(ident, "StartAngle", "nil")
		}
	case "EndAngle":
		if arcbycenterpointtype.EndAngle != nil {
			res = __gong__marshallPointer(ident, "EndAngle", arcbycenterpointtype.EndAngle.GongGetIdentifier(stage))
		} else {
			res = __gong__marshallPointer(ident, "EndAngle", "nil")
		}
	default:
		log.Panicf("Unknown field %s for Gongstruct ArcByCenterPointType", fieldName)
	}
	return
}

func (arcstringtype *ArcStringType) GongMarshallField(stage *Stage, fieldName string) (res string) {
	ident := arcstringtype.GongGetIdentifier(stage)
	_ = ident
	switch fieldName {
	case "Name":
		res = __gong__marshallString(ident, "Name", arcstringtype.Name)
	case "NumDerivativesAtStart":
		res = __gong__marshallInt(ident, "NumDerivativesAtStart", arcstringtype.NumDerivativesAtStart)
	case "NumDerivativesAtEnd":
		res = __gong__marshallInt(ident, "NumDerivativesAtEnd", arcstringtype.NumDerivativesAtEnd)
	case "NumDerivativeInterior":
		res = __gong__marshallInt(ident, "NumDerivativeInterior", arcstringtype.NumDerivativeInterior)
	case "Interpolation":
		res = __gong__marshallEnumString(ident, "Interpolation", arcstringtype.Interpolation.ToCodeString())
	case "NumArc":
		res = __gong__marshallInt(ident, "NumArc", arcstringtype.NumArc)

	case "Pos":
		var sb strings.Builder
		for _, _directpositiontype := range arcstringtype.Pos {
			sb.WriteString(__gong__marshallSliceOfPointers(ident, "Pos", _directpositiontype.GongGetIdentifier(stage)))
		}
		res = sb.String()
	case "PointProperty":
		var sb strings.Builder
		for _, _pointpropertytype := range arcstringtype.PointProperty {
			sb.WriteString(__gong__marshallSliceOfPointers(ident, "PointProperty", _pointpropertytype.GongGetIdentifier(stage)))
		}
		res = sb.String()
	case "PosList":
		if arcstringtype.PosList != nil {
			res = __gong__marshallPointer(ident, "PosList", arcstringtype.PosList.GongGetIdentifier(stage))
		} else {
			res = __gong__marshallPointer(ident, "PosList", "nil")
		}
	default:
		log.Panicf("Unknown field %s for Gongstruct ArcStringType", fieldName)
	}
	return
}

func (arctype *ArcType) GongMarshallField(stage *Stage, fieldName string) (res string) {
	ident := arctype.GongGetIdentifier(stage)
	_ = ident
	switch fieldName {
	case "Name":
		res = __gong__marshallString(ident, "Name", arctype.Name)
	case "NumDerivativesAtStart":
		res = __gong__marshallInt(ident, "NumDerivativesAtStart", arctype.NumDerivativesAtStart)
	case "NumDerivativesAtEnd":
		res = __gong__marshallInt(ident, "NumDerivativesAtEnd", arctype.NumDerivativesAtEnd)
	case "NumDerivativeInterior":
		res = __gong__marshallInt(ident, "NumDerivativeInterior", arctype.NumDerivativeInterior)
	case "Interpolation":
		res = __gong__marshallEnumString(ident, "Interpolation", arctype.Interpolation.ToCodeString())
	case "NumArc":
		res = __gong__marshallInt(ident, "NumArc", arctype.NumArc)

	case "Pos":
		if arctype.Pos != nil {
			res = __gong__marshallPointer(ident, "Pos", arctype.Pos.GongGetIdentifier(stage))
		} else {
			res = __gong__marshallPointer(ident, "Pos", "nil")
		}
	case "PointProperty":
		if arctype.PointProperty != nil {
			res = __gong__marshallPointer(ident, "PointProperty", arctype.PointProperty.GongGetIdentifier(stage))
		} else {
			res = __gong__marshallPointer(ident, "PointProperty", "nil")
		}
	case "PosList":
		if arctype.PosList != nil {
			res = __gong__marshallPointer(ident, "PosList", arctype.PosList.GongGetIdentifier(stage))
		} else {
			res = __gong__marshallPointer(ident, "PosList", "nil")
		}
	default:
		log.Panicf("Unknown field %s for Gongstruct ArcType", fieldName)
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

func (circlebycenterpointtype *CircleByCenterPointType) GongMarshallField(stage *Stage, fieldName string) (res string) {
	ident := circlebycenterpointtype.GongGetIdentifier(stage)
	_ = ident
	switch fieldName {
	case "Name":
		res = __gong__marshallString(ident, "Name", circlebycenterpointtype.Name)
	case "NumDerivativesAtStart":
		res = __gong__marshallInt(ident, "NumDerivativesAtStart", circlebycenterpointtype.NumDerivativesAtStart)
	case "NumDerivativesAtEnd":
		res = __gong__marshallInt(ident, "NumDerivativesAtEnd", circlebycenterpointtype.NumDerivativesAtEnd)
	case "NumDerivativeInterior":
		res = __gong__marshallInt(ident, "NumDerivativeInterior", circlebycenterpointtype.NumDerivativeInterior)
	case "Interpolation":
		res = __gong__marshallEnumString(ident, "Interpolation", circlebycenterpointtype.Interpolation.ToCodeString())
	case "NumArc":
		res = __gong__marshallInt(ident, "NumArc", circlebycenterpointtype.NumArc)

	case "Pos":
		if circlebycenterpointtype.Pos != nil {
			res = __gong__marshallPointer(ident, "Pos", circlebycenterpointtype.Pos.GongGetIdentifier(stage))
		} else {
			res = __gong__marshallPointer(ident, "Pos", "nil")
		}
	case "PointProperty":
		if circlebycenterpointtype.PointProperty != nil {
			res = __gong__marshallPointer(ident, "PointProperty", circlebycenterpointtype.PointProperty.GongGetIdentifier(stage))
		} else {
			res = __gong__marshallPointer(ident, "PointProperty", "nil")
		}
	case "PosList":
		if circlebycenterpointtype.PosList != nil {
			res = __gong__marshallPointer(ident, "PosList", circlebycenterpointtype.PosList.GongGetIdentifier(stage))
		} else {
			res = __gong__marshallPointer(ident, "PosList", "nil")
		}
	case "Radius":
		if circlebycenterpointtype.Radius != nil {
			res = __gong__marshallPointer(ident, "Radius", circlebycenterpointtype.Radius.GongGetIdentifier(stage))
		} else {
			res = __gong__marshallPointer(ident, "Radius", "nil")
		}
	default:
		log.Panicf("Unknown field %s for Gongstruct CircleByCenterPointType", fieldName)
	}
	return
}

func (circletype *CircleType) GongMarshallField(stage *Stage, fieldName string) (res string) {
	ident := circletype.GongGetIdentifier(stage)
	_ = ident
	switch fieldName {
	case "Name":
		res = __gong__marshallString(ident, "Name", circletype.Name)
	case "NumDerivativesAtStart":
		res = __gong__marshallInt(ident, "NumDerivativesAtStart", circletype.NumDerivativesAtStart)
	case "NumDerivativesAtEnd":
		res = __gong__marshallInt(ident, "NumDerivativesAtEnd", circletype.NumDerivativesAtEnd)
	case "NumDerivativeInterior":
		res = __gong__marshallInt(ident, "NumDerivativeInterior", circletype.NumDerivativeInterior)
	case "Interpolation":
		res = __gong__marshallEnumString(ident, "Interpolation", circletype.Interpolation.ToCodeString())
	case "NumArc":
		res = __gong__marshallInt(ident, "NumArc", circletype.NumArc)

	case "Pos":
		var sb strings.Builder
		for _, _directpositiontype := range circletype.Pos {
			sb.WriteString(__gong__marshallSliceOfPointers(ident, "Pos", _directpositiontype.GongGetIdentifier(stage)))
		}
		res = sb.String()
	case "PointProperty":
		var sb strings.Builder
		for _, _pointpropertytype := range circletype.PointProperty {
			sb.WriteString(__gong__marshallSliceOfPointers(ident, "PointProperty", _pointpropertytype.GongGetIdentifier(stage)))
		}
		res = sb.String()
	case "PosList":
		if circletype.PosList != nil {
			res = __gong__marshallPointer(ident, "PosList", circletype.PosList.GongGetIdentifier(stage))
		} else {
			res = __gong__marshallPointer(ident, "PosList", "nil")
		}
	default:
		log.Panicf("Unknown field %s for Gongstruct CircleType", fieldName)
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

func (compositecurvetype *CompositeCurveType) GongMarshallField(stage *Stage, fieldName string) (res string) {
	ident := compositecurvetype.GongGetIdentifier(stage)
	_ = ident
	switch fieldName {
	case "Name":
		res = __gong__marshallString(ident, "Name", compositecurvetype.Name)
	case "Id":
		res = __gong__marshallString(ident, "Id", compositecurvetype.Id)
	case "Description":
		res = __gong__marshallString(ident, "Description", compositecurvetype.Description)
	case "SrsName":
		res = __gong__marshallString(ident, "SrsName", compositecurvetype.SrsName)
	case "AggregationType":
		res = __gong__marshallString(ident, "AggregationType", compositecurvetype.AggregationType)

	case "Identifier":
		if compositecurvetype.Identifier != nil {
			res = __gong__marshallPointer(ident, "Identifier", compositecurvetype.Identifier.GongGetIdentifier(stage))
		} else {
			res = __gong__marshallPointer(ident, "Identifier", "nil")
		}
	case "NameXSD":
		var sb strings.Builder
		for _, _codetype := range compositecurvetype.NameXSD {
			sb.WriteString(__gong__marshallSliceOfPointers(ident, "NameXSD", _codetype.GongGetIdentifier(stage)))
		}
		res = sb.String()
	case "CurveMember":
		var sb strings.Builder
		for _, _curvepropertytype := range compositecurvetype.CurveMember {
			sb.WriteString(__gong__marshallSliceOfPointers(ident, "CurveMember", _curvepropertytype.GongGetIdentifier(stage)))
		}
		res = sb.String()
	default:
		log.Panicf("Unknown field %s for Gongstruct CompositeCurveType", fieldName)
	}
	return
}

func (curvepropertytype *CurvePropertyType) GongMarshallField(stage *Stage, fieldName string) (res string) {
	ident := curvepropertytype.GongGetIdentifier(stage)
	_ = ident
	switch fieldName {
	case "Name":
		res = __gong__marshallString(ident, "Name", curvepropertytype.Name)
	case "NilReason":
		res = __gong__marshallString(ident, "NilReason", curvepropertytype.NilReason)
	case "Type_XSD":
		res = __gong__marshallString(ident, "Type_XSD", curvepropertytype.Type_XSD)
	case "Href":
		res = __gong__marshallString(ident, "Href", curvepropertytype.Href)
	case "Role":
		res = __gong__marshallString(ident, "Role", curvepropertytype.Role)
	case "Arcrole":
		res = __gong__marshallString(ident, "Arcrole", curvepropertytype.Arcrole)
	case "Title":
		res = __gong__marshallString(ident, "Title", curvepropertytype.Title)
	case "Show":
		res = __gong__marshallString(ident, "Show", curvepropertytype.Show)
	case "Actuate":
		res = __gong__marshallString(ident, "Actuate", curvepropertytype.Actuate)

	case "Curve":
		if curvepropertytype.Curve != nil {
			res = __gong__marshallPointer(ident, "Curve", curvepropertytype.Curve.GongGetIdentifier(stage))
		} else {
			res = __gong__marshallPointer(ident, "Curve", "nil")
		}
	case "OrientableCurve":
		if curvepropertytype.OrientableCurve != nil {
			res = __gong__marshallPointer(ident, "OrientableCurve", curvepropertytype.OrientableCurve.GongGetIdentifier(stage))
		} else {
			res = __gong__marshallPointer(ident, "OrientableCurve", "nil")
		}
	case "CompositeCurve":
		if curvepropertytype.CompositeCurve != nil {
			res = __gong__marshallPointer(ident, "CompositeCurve", curvepropertytype.CompositeCurve.GongGetIdentifier(stage))
		} else {
			res = __gong__marshallPointer(ident, "CompositeCurve", "nil")
		}
	default:
		log.Panicf("Unknown field %s for Gongstruct CurvePropertyType", fieldName)
	}
	return
}

func (curvesegmentarraypropertytype *CurveSegmentArrayPropertyType) GongMarshallField(stage *Stage, fieldName string) (res string) {
	ident := curvesegmentarraypropertytype.GongGetIdentifier(stage)
	_ = ident
	switch fieldName {
	case "Name":
		res = __gong__marshallString(ident, "Name", curvesegmentarraypropertytype.Name)

	case "LineStringSegment":
		var sb strings.Builder
		for _, _linestringsegmenttype := range curvesegmentarraypropertytype.LineStringSegment {
			sb.WriteString(__gong__marshallSliceOfPointers(ident, "LineStringSegment", _linestringsegmenttype.GongGetIdentifier(stage)))
		}
		res = sb.String()
	case "ArcString":
		var sb strings.Builder
		for _, _arcstringtype := range curvesegmentarraypropertytype.ArcString {
			sb.WriteString(__gong__marshallSliceOfPointers(ident, "ArcString", _arcstringtype.GongGetIdentifier(stage)))
		}
		res = sb.String()
	case "Arc":
		var sb strings.Builder
		for _, _arctype := range curvesegmentarraypropertytype.Arc {
			sb.WriteString(__gong__marshallSliceOfPointers(ident, "Arc", _arctype.GongGetIdentifier(stage)))
		}
		res = sb.String()
	case "Circle":
		var sb strings.Builder
		for _, _circletype := range curvesegmentarraypropertytype.Circle {
			sb.WriteString(__gong__marshallSliceOfPointers(ident, "Circle", _circletype.GongGetIdentifier(stage)))
		}
		res = sb.String()
	case "ArcByCenterPoint":
		var sb strings.Builder
		for _, _arcbycenterpointtype := range curvesegmentarraypropertytype.ArcByCenterPoint {
			sb.WriteString(__gong__marshallSliceOfPointers(ident, "ArcByCenterPoint", _arcbycenterpointtype.GongGetIdentifier(stage)))
		}
		res = sb.String()
	case "CircleByCenterPoint":
		var sb strings.Builder
		for _, _circlebycenterpointtype := range curvesegmentarraypropertytype.CircleByCenterPoint {
			sb.WriteString(__gong__marshallSliceOfPointers(ident, "CircleByCenterPoint", _circlebycenterpointtype.GongGetIdentifier(stage)))
		}
		res = sb.String()
	case "GeodesicString":
		var sb strings.Builder
		for _, _geodesicstringtype := range curvesegmentarraypropertytype.GeodesicString {
			sb.WriteString(__gong__marshallSliceOfPointers(ident, "GeodesicString", _geodesicstringtype.GongGetIdentifier(stage)))
		}
		res = sb.String()
	case "Geodesic":
		var sb strings.Builder
		for _, _geodesictype := range curvesegmentarraypropertytype.Geodesic {
			sb.WriteString(__gong__marshallSliceOfPointers(ident, "Geodesic", _geodesictype.GongGetIdentifier(stage)))
		}
		res = sb.String()
	default:
		log.Panicf("Unknown field %s for Gongstruct CurveSegmentArrayPropertyType", fieldName)
	}
	return
}

func (curvetype *CurveType) GongMarshallField(stage *Stage, fieldName string) (res string) {
	ident := curvetype.GongGetIdentifier(stage)
	_ = ident
	switch fieldName {
	case "Name":
		res = __gong__marshallString(ident, "Name", curvetype.Name)
	case "Id":
		res = __gong__marshallString(ident, "Id", curvetype.Id)
	case "Description":
		res = __gong__marshallString(ident, "Description", curvetype.Description)
	case "SrsName":
		res = __gong__marshallString(ident, "SrsName", curvetype.SrsName)

	case "Identifier":
		if curvetype.Identifier != nil {
			res = __gong__marshallPointer(ident, "Identifier", curvetype.Identifier.GongGetIdentifier(stage))
		} else {
			res = __gong__marshallPointer(ident, "Identifier", "nil")
		}
	case "NameXSD":
		var sb strings.Builder
		for _, _codetype := range curvetype.NameXSD {
			sb.WriteString(__gong__marshallSliceOfPointers(ident, "NameXSD", _codetype.GongGetIdentifier(stage)))
		}
		res = sb.String()
	case "Segments":
		if curvetype.Segments != nil {
			res = __gong__marshallPointer(ident, "Segments", curvetype.Segments.GongGetIdentifier(stage))
		} else {
			res = __gong__marshallPointer(ident, "Segments", "nil")
		}
	default:
		log.Panicf("Unknown field %s for Gongstruct CurveType", fieldName)
	}
	return
}

func (directpositionlisttype *DirectPositionListType) GongMarshallField(stage *Stage, fieldName string) (res string) {
	ident := directpositionlisttype.GongGetIdentifier(stage)
	_ = ident
	switch fieldName {
	case "Name":
		res = __gong__marshallString(ident, "Name", directpositionlisttype.Name)
	case "SrsName":
		res = __gong__marshallString(ident, "SrsName", directpositionlisttype.SrsName)
	case "EnclosedText":
		res = __gong__marshallString(ident, "EnclosedText", directpositionlisttype.EnclosedText)

	default:
		log.Panicf("Unknown field %s for Gongstruct DirectPositionListType", fieldName)
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

func (geodesicstringtype *GeodesicStringType) GongMarshallField(stage *Stage, fieldName string) (res string) {
	ident := geodesicstringtype.GongGetIdentifier(stage)
	_ = ident
	switch fieldName {
	case "Name":
		res = __gong__marshallString(ident, "Name", geodesicstringtype.Name)
	case "NumDerivativesAtStart":
		res = __gong__marshallInt(ident, "NumDerivativesAtStart", geodesicstringtype.NumDerivativesAtStart)
	case "NumDerivativesAtEnd":
		res = __gong__marshallInt(ident, "NumDerivativesAtEnd", geodesicstringtype.NumDerivativesAtEnd)
	case "NumDerivativeInterior":
		res = __gong__marshallInt(ident, "NumDerivativeInterior", geodesicstringtype.NumDerivativeInterior)
	case "Interpolation":
		res = __gong__marshallEnumString(ident, "Interpolation", geodesicstringtype.Interpolation.ToCodeString())

	case "PosList":
		if geodesicstringtype.PosList != nil {
			res = __gong__marshallPointer(ident, "PosList", geodesicstringtype.PosList.GongGetIdentifier(stage))
		} else {
			res = __gong__marshallPointer(ident, "PosList", "nil")
		}
	case "Pos":
		if geodesicstringtype.Pos != nil {
			res = __gong__marshallPointer(ident, "Pos", geodesicstringtype.Pos.GongGetIdentifier(stage))
		} else {
			res = __gong__marshallPointer(ident, "Pos", "nil")
		}
	case "PointProperty":
		if geodesicstringtype.PointProperty != nil {
			res = __gong__marshallPointer(ident, "PointProperty", geodesicstringtype.PointProperty.GongGetIdentifier(stage))
		} else {
			res = __gong__marshallPointer(ident, "PointProperty", "nil")
		}
	default:
		log.Panicf("Unknown field %s for Gongstruct GeodesicStringType", fieldName)
	}
	return
}

func (geodesictype *GeodesicType) GongMarshallField(stage *Stage, fieldName string) (res string) {
	ident := geodesictype.GongGetIdentifier(stage)
	_ = ident
	switch fieldName {
	case "Name":
		res = __gong__marshallString(ident, "Name", geodesictype.Name)
	case "NumDerivativesAtStart":
		res = __gong__marshallInt(ident, "NumDerivativesAtStart", geodesictype.NumDerivativesAtStart)
	case "NumDerivativesAtEnd":
		res = __gong__marshallInt(ident, "NumDerivativesAtEnd", geodesictype.NumDerivativesAtEnd)
	case "NumDerivativeInterior":
		res = __gong__marshallInt(ident, "NumDerivativeInterior", geodesictype.NumDerivativeInterior)
	case "Interpolation":
		res = __gong__marshallEnumString(ident, "Interpolation", geodesictype.Interpolation.ToCodeString())

	case "PosList":
		if geodesictype.PosList != nil {
			res = __gong__marshallPointer(ident, "PosList", geodesictype.PosList.GongGetIdentifier(stage))
		} else {
			res = __gong__marshallPointer(ident, "PosList", "nil")
		}
	case "Pos":
		if geodesictype.Pos != nil {
			res = __gong__marshallPointer(ident, "Pos", geodesictype.Pos.GongGetIdentifier(stage))
		} else {
			res = __gong__marshallPointer(ident, "Pos", "nil")
		}
	case "PointProperty":
		if geodesictype.PointProperty != nil {
			res = __gong__marshallPointer(ident, "PointProperty", geodesictype.PointProperty.GongGetIdentifier(stage))
		} else {
			res = __gong__marshallPointer(ident, "PointProperty", "nil")
		}
	default:
		log.Panicf("Unknown field %s for Gongstruct GeodesicType", fieldName)
	}
	return
}

func (lengthtype *LengthType) GongMarshallField(stage *Stage, fieldName string) (res string) {
	ident := lengthtype.GongGetIdentifier(stage)
	_ = ident
	switch fieldName {
	case "Name":
		res = __gong__marshallString(ident, "Name", lengthtype.Name)
	case "EnclosedText":
		res = __gong__marshallString(ident, "EnclosedText", lengthtype.EnclosedText)

	default:
		log.Panicf("Unknown field %s for Gongstruct LengthType", fieldName)
	}
	return
}

func (linestringsegmenttype *LineStringSegmentType) GongMarshallField(stage *Stage, fieldName string) (res string) {
	ident := linestringsegmenttype.GongGetIdentifier(stage)
	_ = ident
	switch fieldName {
	case "Name":
		res = __gong__marshallString(ident, "Name", linestringsegmenttype.Name)
	case "NumDerivativesAtStart":
		res = __gong__marshallInt(ident, "NumDerivativesAtStart", linestringsegmenttype.NumDerivativesAtStart)
	case "NumDerivativesAtEnd":
		res = __gong__marshallInt(ident, "NumDerivativesAtEnd", linestringsegmenttype.NumDerivativesAtEnd)
	case "NumDerivativeInterior":
		res = __gong__marshallInt(ident, "NumDerivativeInterior", linestringsegmenttype.NumDerivativeInterior)
	case "Interpolation":
		res = __gong__marshallEnumString(ident, "Interpolation", linestringsegmenttype.Interpolation.ToCodeString())

	case "Pos":
		var sb strings.Builder
		for _, _directpositiontype := range linestringsegmenttype.Pos {
			sb.WriteString(__gong__marshallSliceOfPointers(ident, "Pos", _directpositiontype.GongGetIdentifier(stage)))
		}
		res = sb.String()
	case "PointProperty":
		var sb strings.Builder
		for _, _pointpropertytype := range linestringsegmenttype.PointProperty {
			sb.WriteString(__gong__marshallSliceOfPointers(ident, "PointProperty", _pointpropertytype.GongGetIdentifier(stage)))
		}
		res = sb.String()
	case "PosList":
		if linestringsegmenttype.PosList != nil {
			res = __gong__marshallPointer(ident, "PosList", linestringsegmenttype.PosList.GongGetIdentifier(stage))
		} else {
			res = __gong__marshallPointer(ident, "PosList", "nil")
		}
	default:
		log.Panicf("Unknown field %s for Gongstruct LineStringSegmentType", fieldName)
	}
	return
}

func (measuretype *MeasureType) GongMarshallField(stage *Stage, fieldName string) (res string) {
	ident := measuretype.GongGetIdentifier(stage)
	_ = ident
	switch fieldName {
	case "Name":
		res = __gong__marshallString(ident, "Name", measuretype.Name)
	case "Uom":
		res = __gong__marshallString(ident, "Uom", measuretype.Uom)
	case "EnclosedText":
		res = __gong__marshallFloat(ident, "EnclosedText", measuretype.EnclosedText)

	default:
		log.Panicf("Unknown field %s for Gongstruct MeasureType", fieldName)
	}
	return
}

func (orientablecurvetype *OrientableCurveType) GongMarshallField(stage *Stage, fieldName string) (res string) {
	ident := orientablecurvetype.GongGetIdentifier(stage)
	_ = ident
	switch fieldName {
	case "Name":
		res = __gong__marshallString(ident, "Name", orientablecurvetype.Name)
	case "Id":
		res = __gong__marshallString(ident, "Id", orientablecurvetype.Id)
	case "Description":
		res = __gong__marshallString(ident, "Description", orientablecurvetype.Description)
	case "SrsName":
		res = __gong__marshallString(ident, "SrsName", orientablecurvetype.SrsName)
	case "Orientation":
		res = __gong__marshallEnumString(ident, "Orientation", orientablecurvetype.Orientation.ToCodeString())

	case "Identifier":
		if orientablecurvetype.Identifier != nil {
			res = __gong__marshallPointer(ident, "Identifier", orientablecurvetype.Identifier.GongGetIdentifier(stage))
		} else {
			res = __gong__marshallPointer(ident, "Identifier", "nil")
		}
	case "NameXSD":
		var sb strings.Builder
		for _, _codetype := range orientablecurvetype.NameXSD {
			sb.WriteString(__gong__marshallSliceOfPointers(ident, "NameXSD", _codetype.GongGetIdentifier(stage)))
		}
		res = sb.String()
	case "BaseCurve":
		if orientablecurvetype.BaseCurve != nil {
			res = __gong__marshallPointer(ident, "BaseCurve", orientablecurvetype.BaseCurve.GongGetIdentifier(stage))
		} else {
			res = __gong__marshallPointer(ident, "BaseCurve", "nil")
		}
	default:
		log.Panicf("Unknown field %s for Gongstruct OrientableCurveType", fieldName)
	}
	return
}

func (pointpropertytype *PointPropertyType) GongMarshallField(stage *Stage, fieldName string) (res string) {
	ident := pointpropertytype.GongGetIdentifier(stage)
	_ = ident
	switch fieldName {
	case "Name":
		res = __gong__marshallString(ident, "Name", pointpropertytype.Name)
	case "NilReason":
		res = __gong__marshallString(ident, "NilReason", pointpropertytype.NilReason)
	case "Type_XSD":
		res = __gong__marshallString(ident, "Type_XSD", pointpropertytype.Type_XSD)
	case "Href":
		res = __gong__marshallString(ident, "Href", pointpropertytype.Href)
	case "Role":
		res = __gong__marshallString(ident, "Role", pointpropertytype.Role)
	case "Arcrole":
		res = __gong__marshallString(ident, "Arcrole", pointpropertytype.Arcrole)
	case "Title":
		res = __gong__marshallString(ident, "Title", pointpropertytype.Title)
	case "Show":
		res = __gong__marshallString(ident, "Show", pointpropertytype.Show)
	case "Actuate":
		res = __gong__marshallString(ident, "Actuate", pointpropertytype.Actuate)

	case "Point":
		if pointpropertytype.Point != nil {
			res = __gong__marshallPointer(ident, "Point", pointpropertytype.Point.GongGetIdentifier(stage))
		} else {
			res = __gong__marshallPointer(ident, "Point", "nil")
		}
	default:
		log.Panicf("Unknown field %s for Gongstruct PointPropertyType", fieldName)
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

func (polygonpatchtype *PolygonPatchType) GongMarshallField(stage *Stage, fieldName string) (res string) {
	ident := polygonpatchtype.GongGetIdentifier(stage)
	_ = ident
	switch fieldName {
	case "Name":
		res = __gong__marshallString(ident, "Name", polygonpatchtype.Name)
	case "Interpolation":
		res = __gong__marshallEnumString(ident, "Interpolation", polygonpatchtype.Interpolation.ToCodeString())

	case "Exterior":
		if polygonpatchtype.Exterior != nil {
			res = __gong__marshallPointer(ident, "Exterior", polygonpatchtype.Exterior.GongGetIdentifier(stage))
		} else {
			res = __gong__marshallPointer(ident, "Exterior", "nil")
		}
	default:
		log.Panicf("Unknown field %s for Gongstruct PolygonPatchType", fieldName)
	}
	return
}

func (ringtype *RingType) GongMarshallField(stage *Stage, fieldName string) (res string) {
	ident := ringtype.GongGetIdentifier(stage)
	_ = ident
	switch fieldName {
	case "Name":
		res = __gong__marshallString(ident, "Name", ringtype.Name)
	case "AggregationType":
		res = __gong__marshallString(ident, "AggregationType", ringtype.AggregationType)

	case "CurveMember":
		var sb strings.Builder
		for _, _curvepropertytype := range ringtype.CurveMember {
			sb.WriteString(__gong__marshallSliceOfPointers(ident, "CurveMember", _curvepropertytype.GongGetIdentifier(stage)))
		}
		res = sb.String()
	default:
		log.Panicf("Unknown field %s for Gongstruct RingType", fieldName)
	}
	return
}

func (surfacepatcharraypropertytype *SurfacePatchArrayPropertyType) GongMarshallField(stage *Stage, fieldName string) (res string) {
	ident := surfacepatcharraypropertytype.GongGetIdentifier(stage)
	_ = ident
	switch fieldName {
	case "Name":
		res = __gong__marshallString(ident, "Name", surfacepatcharraypropertytype.Name)

	case "PolygonPatch":
		var sb strings.Builder
		for _, _polygonpatchtype := range surfacepatcharraypropertytype.PolygonPatch {
			sb.WriteString(__gong__marshallSliceOfPointers(ident, "PolygonPatch", _polygonpatchtype.GongGetIdentifier(stage)))
		}
		res = sb.String()
	default:
		log.Panicf("Unknown field %s for Gongstruct SurfacePatchArrayPropertyType", fieldName)
	}
	return
}

func (surfacetype *SurfaceType) GongMarshallField(stage *Stage, fieldName string) (res string) {
	ident := surfacetype.GongGetIdentifier(stage)
	_ = ident
	switch fieldName {
	case "Name":
		res = __gong__marshallString(ident, "Name", surfacetype.Name)
	case "Id":
		res = __gong__marshallString(ident, "Id", surfacetype.Id)
	case "Description":
		res = __gong__marshallString(ident, "Description", surfacetype.Description)
	case "SrsName":
		res = __gong__marshallString(ident, "SrsName", surfacetype.SrsName)

	case "Identifier":
		if surfacetype.Identifier != nil {
			res = __gong__marshallPointer(ident, "Identifier", surfacetype.Identifier.GongGetIdentifier(stage))
		} else {
			res = __gong__marshallPointer(ident, "Identifier", "nil")
		}
	case "NameXSD":
		var sb strings.Builder
		for _, _codetype := range surfacetype.NameXSD {
			sb.WriteString(__gong__marshallSliceOfPointers(ident, "NameXSD", _codetype.GongGetIdentifier(stage)))
		}
		res = sb.String()
	case "Patches":
		if surfacetype.Patches != nil {
			res = __gong__marshallPointer(ident, "Patches", surfacetype.Patches.GongGetIdentifier(stage))
		} else {
			res = __gong__marshallPointer(ident, "Patches", "nil")
		}
	default:
		log.Panicf("Unknown field %s for Gongstruct SurfaceType", fieldName)
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
func (abstractcurvesegmenttype *AbstractCurveSegmentType) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(abstractcurvesegmenttype.GongMarshallField(stage, "Name"))
		initializerStatements.WriteString(abstractcurvesegmenttype.GongMarshallField(stage, "NumDerivativesAtStart"))
		initializerStatements.WriteString(abstractcurvesegmenttype.GongMarshallField(stage, "NumDerivativesAtEnd"))
		initializerStatements.WriteString(abstractcurvesegmenttype.GongMarshallField(stage, "NumDerivativeInterior"))
	}
	initRes = initializerStatements.String()
	ptrRes = pointersInitializesStatements.String()
	return
}
func (abstractcurvetype *AbstractCurveType) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(abstractcurvetype.GongMarshallField(stage, "Name"))
		initializerStatements.WriteString(abstractcurvetype.GongMarshallField(stage, "Id"))
		initializerStatements.WriteString(abstractcurvetype.GongMarshallField(stage, "Description"))
		pointersInitializesStatements.WriteString(abstractcurvetype.GongMarshallField(stage, "Identifier"))
		pointersInitializesStatements.WriteString(abstractcurvetype.GongMarshallField(stage, "NameXSD"))
		initializerStatements.WriteString(abstractcurvetype.GongMarshallField(stage, "SrsName"))
	}
	initRes = initializerStatements.String()
	ptrRes = pointersInitializesStatements.String()
	return
}
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
func (abstractringpropertytype *AbstractRingPropertyType) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(abstractringpropertytype.GongMarshallField(stage, "Name"))
		pointersInitializesStatements.WriteString(abstractringpropertytype.GongMarshallField(stage, "Ring"))
	}
	initRes = initializerStatements.String()
	ptrRes = pointersInitializesStatements.String()
	return
}
func (abstractringtype *AbstractRingType) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(abstractringtype.GongMarshallField(stage, "Name"))
	}
	initRes = initializerStatements.String()
	ptrRes = pointersInitializesStatements.String()
	return
}
func (abstractsurfacepatchtype *AbstractSurfacePatchType) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(abstractsurfacepatchtype.GongMarshallField(stage, "Name"))
	}
	initRes = initializerStatements.String()
	ptrRes = pointersInitializesStatements.String()
	return
}
func (abstractsurfacetype *AbstractSurfaceType) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(abstractsurfacetype.GongMarshallField(stage, "Name"))
		initializerStatements.WriteString(abstractsurfacetype.GongMarshallField(stage, "Id"))
		initializerStatements.WriteString(abstractsurfacetype.GongMarshallField(stage, "Description"))
		pointersInitializesStatements.WriteString(abstractsurfacetype.GongMarshallField(stage, "Identifier"))
		pointersInitializesStatements.WriteString(abstractsurfacetype.GongMarshallField(stage, "NameXSD"))
		initializerStatements.WriteString(abstractsurfacetype.GongMarshallField(stage, "SrsName"))
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
func (angletype *AngleType) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(angletype.GongMarshallField(stage, "Name"))
		initializerStatements.WriteString(angletype.GongMarshallField(stage, "EnclosedText"))
	}
	initRes = initializerStatements.String()
	ptrRes = pointersInitializesStatements.String()
	return
}
func (arcbycenterpointtype *ArcByCenterPointType) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(arcbycenterpointtype.GongMarshallField(stage, "Name"))
		initializerStatements.WriteString(arcbycenterpointtype.GongMarshallField(stage, "NumDerivativesAtStart"))
		initializerStatements.WriteString(arcbycenterpointtype.GongMarshallField(stage, "NumDerivativesAtEnd"))
		initializerStatements.WriteString(arcbycenterpointtype.GongMarshallField(stage, "NumDerivativeInterior"))
		initializerStatements.WriteString(arcbycenterpointtype.GongMarshallField(stage, "Interpolation"))
		initializerStatements.WriteString(arcbycenterpointtype.GongMarshallField(stage, "NumArc"))
		pointersInitializesStatements.WriteString(arcbycenterpointtype.GongMarshallField(stage, "Pos"))
		pointersInitializesStatements.WriteString(arcbycenterpointtype.GongMarshallField(stage, "PointProperty"))
		pointersInitializesStatements.WriteString(arcbycenterpointtype.GongMarshallField(stage, "PosList"))
		pointersInitializesStatements.WriteString(arcbycenterpointtype.GongMarshallField(stage, "Radius"))
		pointersInitializesStatements.WriteString(arcbycenterpointtype.GongMarshallField(stage, "StartAngle"))
		pointersInitializesStatements.WriteString(arcbycenterpointtype.GongMarshallField(stage, "EndAngle"))
	}
	initRes = initializerStatements.String()
	ptrRes = pointersInitializesStatements.String()
	return
}
func (arcstringtype *ArcStringType) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(arcstringtype.GongMarshallField(stage, "Name"))
		initializerStatements.WriteString(arcstringtype.GongMarshallField(stage, "NumDerivativesAtStart"))
		initializerStatements.WriteString(arcstringtype.GongMarshallField(stage, "NumDerivativesAtEnd"))
		initializerStatements.WriteString(arcstringtype.GongMarshallField(stage, "NumDerivativeInterior"))
		initializerStatements.WriteString(arcstringtype.GongMarshallField(stage, "Interpolation"))
		initializerStatements.WriteString(arcstringtype.GongMarshallField(stage, "NumArc"))
		pointersInitializesStatements.WriteString(arcstringtype.GongMarshallField(stage, "Pos"))
		pointersInitializesStatements.WriteString(arcstringtype.GongMarshallField(stage, "PointProperty"))
		pointersInitializesStatements.WriteString(arcstringtype.GongMarshallField(stage, "PosList"))
	}
	initRes = initializerStatements.String()
	ptrRes = pointersInitializesStatements.String()
	return
}
func (arctype *ArcType) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(arctype.GongMarshallField(stage, "Name"))
		initializerStatements.WriteString(arctype.GongMarshallField(stage, "NumDerivativesAtStart"))
		initializerStatements.WriteString(arctype.GongMarshallField(stage, "NumDerivativesAtEnd"))
		initializerStatements.WriteString(arctype.GongMarshallField(stage, "NumDerivativeInterior"))
		initializerStatements.WriteString(arctype.GongMarshallField(stage, "Interpolation"))
		initializerStatements.WriteString(arctype.GongMarshallField(stage, "NumArc"))
		pointersInitializesStatements.WriteString(arctype.GongMarshallField(stage, "Pos"))
		pointersInitializesStatements.WriteString(arctype.GongMarshallField(stage, "PointProperty"))
		pointersInitializesStatements.WriteString(arctype.GongMarshallField(stage, "PosList"))
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
func (circlebycenterpointtype *CircleByCenterPointType) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(circlebycenterpointtype.GongMarshallField(stage, "Name"))
		initializerStatements.WriteString(circlebycenterpointtype.GongMarshallField(stage, "NumDerivativesAtStart"))
		initializerStatements.WriteString(circlebycenterpointtype.GongMarshallField(stage, "NumDerivativesAtEnd"))
		initializerStatements.WriteString(circlebycenterpointtype.GongMarshallField(stage, "NumDerivativeInterior"))
		initializerStatements.WriteString(circlebycenterpointtype.GongMarshallField(stage, "Interpolation"))
		initializerStatements.WriteString(circlebycenterpointtype.GongMarshallField(stage, "NumArc"))
		pointersInitializesStatements.WriteString(circlebycenterpointtype.GongMarshallField(stage, "Pos"))
		pointersInitializesStatements.WriteString(circlebycenterpointtype.GongMarshallField(stage, "PointProperty"))
		pointersInitializesStatements.WriteString(circlebycenterpointtype.GongMarshallField(stage, "PosList"))
		pointersInitializesStatements.WriteString(circlebycenterpointtype.GongMarshallField(stage, "Radius"))
	}
	initRes = initializerStatements.String()
	ptrRes = pointersInitializesStatements.String()
	return
}
func (circletype *CircleType) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(circletype.GongMarshallField(stage, "Name"))
		initializerStatements.WriteString(circletype.GongMarshallField(stage, "NumDerivativesAtStart"))
		initializerStatements.WriteString(circletype.GongMarshallField(stage, "NumDerivativesAtEnd"))
		initializerStatements.WriteString(circletype.GongMarshallField(stage, "NumDerivativeInterior"))
		initializerStatements.WriteString(circletype.GongMarshallField(stage, "Interpolation"))
		initializerStatements.WriteString(circletype.GongMarshallField(stage, "NumArc"))
		pointersInitializesStatements.WriteString(circletype.GongMarshallField(stage, "Pos"))
		pointersInitializesStatements.WriteString(circletype.GongMarshallField(stage, "PointProperty"))
		pointersInitializesStatements.WriteString(circletype.GongMarshallField(stage, "PosList"))
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
func (compositecurvetype *CompositeCurveType) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(compositecurvetype.GongMarshallField(stage, "Name"))
		initializerStatements.WriteString(compositecurvetype.GongMarshallField(stage, "Id"))
		initializerStatements.WriteString(compositecurvetype.GongMarshallField(stage, "Description"))
		pointersInitializesStatements.WriteString(compositecurvetype.GongMarshallField(stage, "Identifier"))
		pointersInitializesStatements.WriteString(compositecurvetype.GongMarshallField(stage, "NameXSD"))
		initializerStatements.WriteString(compositecurvetype.GongMarshallField(stage, "SrsName"))
		initializerStatements.WriteString(compositecurvetype.GongMarshallField(stage, "AggregationType"))
		pointersInitializesStatements.WriteString(compositecurvetype.GongMarshallField(stage, "CurveMember"))
	}
	initRes = initializerStatements.String()
	ptrRes = pointersInitializesStatements.String()
	return
}
func (curvepropertytype *CurvePropertyType) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(curvepropertytype.GongMarshallField(stage, "Name"))
		initializerStatements.WriteString(curvepropertytype.GongMarshallField(stage, "NilReason"))
		initializerStatements.WriteString(curvepropertytype.GongMarshallField(stage, "Type_XSD"))
		initializerStatements.WriteString(curvepropertytype.GongMarshallField(stage, "Href"))
		initializerStatements.WriteString(curvepropertytype.GongMarshallField(stage, "Role"))
		initializerStatements.WriteString(curvepropertytype.GongMarshallField(stage, "Arcrole"))
		initializerStatements.WriteString(curvepropertytype.GongMarshallField(stage, "Title"))
		initializerStatements.WriteString(curvepropertytype.GongMarshallField(stage, "Show"))
		initializerStatements.WriteString(curvepropertytype.GongMarshallField(stage, "Actuate"))
		pointersInitializesStatements.WriteString(curvepropertytype.GongMarshallField(stage, "Curve"))
		pointersInitializesStatements.WriteString(curvepropertytype.GongMarshallField(stage, "OrientableCurve"))
		pointersInitializesStatements.WriteString(curvepropertytype.GongMarshallField(stage, "CompositeCurve"))
	}
	initRes = initializerStatements.String()
	ptrRes = pointersInitializesStatements.String()
	return
}
func (curvesegmentarraypropertytype *CurveSegmentArrayPropertyType) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(curvesegmentarraypropertytype.GongMarshallField(stage, "Name"))
		pointersInitializesStatements.WriteString(curvesegmentarraypropertytype.GongMarshallField(stage, "LineStringSegment"))
		pointersInitializesStatements.WriteString(curvesegmentarraypropertytype.GongMarshallField(stage, "ArcString"))
		pointersInitializesStatements.WriteString(curvesegmentarraypropertytype.GongMarshallField(stage, "Arc"))
		pointersInitializesStatements.WriteString(curvesegmentarraypropertytype.GongMarshallField(stage, "Circle"))
		pointersInitializesStatements.WriteString(curvesegmentarraypropertytype.GongMarshallField(stage, "ArcByCenterPoint"))
		pointersInitializesStatements.WriteString(curvesegmentarraypropertytype.GongMarshallField(stage, "CircleByCenterPoint"))
		pointersInitializesStatements.WriteString(curvesegmentarraypropertytype.GongMarshallField(stage, "GeodesicString"))
		pointersInitializesStatements.WriteString(curvesegmentarraypropertytype.GongMarshallField(stage, "Geodesic"))
	}
	initRes = initializerStatements.String()
	ptrRes = pointersInitializesStatements.String()
	return
}
func (curvetype *CurveType) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(curvetype.GongMarshallField(stage, "Name"))
		initializerStatements.WriteString(curvetype.GongMarshallField(stage, "Id"))
		initializerStatements.WriteString(curvetype.GongMarshallField(stage, "Description"))
		pointersInitializesStatements.WriteString(curvetype.GongMarshallField(stage, "Identifier"))
		pointersInitializesStatements.WriteString(curvetype.GongMarshallField(stage, "NameXSD"))
		initializerStatements.WriteString(curvetype.GongMarshallField(stage, "SrsName"))
		pointersInitializesStatements.WriteString(curvetype.GongMarshallField(stage, "Segments"))
	}
	initRes = initializerStatements.String()
	ptrRes = pointersInitializesStatements.String()
	return
}
func (directpositionlisttype *DirectPositionListType) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(directpositionlisttype.GongMarshallField(stage, "Name"))
		initializerStatements.WriteString(directpositionlisttype.GongMarshallField(stage, "SrsName"))
		initializerStatements.WriteString(directpositionlisttype.GongMarshallField(stage, "EnclosedText"))
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
func (geodesicstringtype *GeodesicStringType) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(geodesicstringtype.GongMarshallField(stage, "Name"))
		initializerStatements.WriteString(geodesicstringtype.GongMarshallField(stage, "NumDerivativesAtStart"))
		initializerStatements.WriteString(geodesicstringtype.GongMarshallField(stage, "NumDerivativesAtEnd"))
		initializerStatements.WriteString(geodesicstringtype.GongMarshallField(stage, "NumDerivativeInterior"))
		initializerStatements.WriteString(geodesicstringtype.GongMarshallField(stage, "Interpolation"))
		pointersInitializesStatements.WriteString(geodesicstringtype.GongMarshallField(stage, "PosList"))
		pointersInitializesStatements.WriteString(geodesicstringtype.GongMarshallField(stage, "Pos"))
		pointersInitializesStatements.WriteString(geodesicstringtype.GongMarshallField(stage, "PointProperty"))
	}
	initRes = initializerStatements.String()
	ptrRes = pointersInitializesStatements.String()
	return
}
func (geodesictype *GeodesicType) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(geodesictype.GongMarshallField(stage, "Name"))
		initializerStatements.WriteString(geodesictype.GongMarshallField(stage, "NumDerivativesAtStart"))
		initializerStatements.WriteString(geodesictype.GongMarshallField(stage, "NumDerivativesAtEnd"))
		initializerStatements.WriteString(geodesictype.GongMarshallField(stage, "NumDerivativeInterior"))
		initializerStatements.WriteString(geodesictype.GongMarshallField(stage, "Interpolation"))
		pointersInitializesStatements.WriteString(geodesictype.GongMarshallField(stage, "PosList"))
		pointersInitializesStatements.WriteString(geodesictype.GongMarshallField(stage, "Pos"))
		pointersInitializesStatements.WriteString(geodesictype.GongMarshallField(stage, "PointProperty"))
	}
	initRes = initializerStatements.String()
	ptrRes = pointersInitializesStatements.String()
	return
}
func (lengthtype *LengthType) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(lengthtype.GongMarshallField(stage, "Name"))
		initializerStatements.WriteString(lengthtype.GongMarshallField(stage, "EnclosedText"))
	}
	initRes = initializerStatements.String()
	ptrRes = pointersInitializesStatements.String()
	return
}
func (linestringsegmenttype *LineStringSegmentType) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(linestringsegmenttype.GongMarshallField(stage, "Name"))
		initializerStatements.WriteString(linestringsegmenttype.GongMarshallField(stage, "NumDerivativesAtStart"))
		initializerStatements.WriteString(linestringsegmenttype.GongMarshallField(stage, "NumDerivativesAtEnd"))
		initializerStatements.WriteString(linestringsegmenttype.GongMarshallField(stage, "NumDerivativeInterior"))
		initializerStatements.WriteString(linestringsegmenttype.GongMarshallField(stage, "Interpolation"))
		pointersInitializesStatements.WriteString(linestringsegmenttype.GongMarshallField(stage, "Pos"))
		pointersInitializesStatements.WriteString(linestringsegmenttype.GongMarshallField(stage, "PointProperty"))
		pointersInitializesStatements.WriteString(linestringsegmenttype.GongMarshallField(stage, "PosList"))
	}
	initRes = initializerStatements.String()
	ptrRes = pointersInitializesStatements.String()
	return
}
func (measuretype *MeasureType) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(measuretype.GongMarshallField(stage, "Name"))
		initializerStatements.WriteString(measuretype.GongMarshallField(stage, "Uom"))
		initializerStatements.WriteString(measuretype.GongMarshallField(stage, "EnclosedText"))
	}
	initRes = initializerStatements.String()
	ptrRes = pointersInitializesStatements.String()
	return
}
func (orientablecurvetype *OrientableCurveType) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(orientablecurvetype.GongMarshallField(stage, "Name"))
		initializerStatements.WriteString(orientablecurvetype.GongMarshallField(stage, "Id"))
		initializerStatements.WriteString(orientablecurvetype.GongMarshallField(stage, "Description"))
		pointersInitializesStatements.WriteString(orientablecurvetype.GongMarshallField(stage, "Identifier"))
		pointersInitializesStatements.WriteString(orientablecurvetype.GongMarshallField(stage, "NameXSD"))
		initializerStatements.WriteString(orientablecurvetype.GongMarshallField(stage, "SrsName"))
		initializerStatements.WriteString(orientablecurvetype.GongMarshallField(stage, "Orientation"))
		pointersInitializesStatements.WriteString(orientablecurvetype.GongMarshallField(stage, "BaseCurve"))
	}
	initRes = initializerStatements.String()
	ptrRes = pointersInitializesStatements.String()
	return
}
func (pointpropertytype *PointPropertyType) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(pointpropertytype.GongMarshallField(stage, "Name"))
		initializerStatements.WriteString(pointpropertytype.GongMarshallField(stage, "NilReason"))
		initializerStatements.WriteString(pointpropertytype.GongMarshallField(stage, "Type_XSD"))
		initializerStatements.WriteString(pointpropertytype.GongMarshallField(stage, "Href"))
		initializerStatements.WriteString(pointpropertytype.GongMarshallField(stage, "Role"))
		initializerStatements.WriteString(pointpropertytype.GongMarshallField(stage, "Arcrole"))
		initializerStatements.WriteString(pointpropertytype.GongMarshallField(stage, "Title"))
		initializerStatements.WriteString(pointpropertytype.GongMarshallField(stage, "Show"))
		initializerStatements.WriteString(pointpropertytype.GongMarshallField(stage, "Actuate"))
		pointersInitializesStatements.WriteString(pointpropertytype.GongMarshallField(stage, "Point"))
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
func (polygonpatchtype *PolygonPatchType) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(polygonpatchtype.GongMarshallField(stage, "Name"))
		initializerStatements.WriteString(polygonpatchtype.GongMarshallField(stage, "Interpolation"))
		pointersInitializesStatements.WriteString(polygonpatchtype.GongMarshallField(stage, "Exterior"))
	}
	initRes = initializerStatements.String()
	ptrRes = pointersInitializesStatements.String()
	return
}
func (ringtype *RingType) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(ringtype.GongMarshallField(stage, "Name"))
		initializerStatements.WriteString(ringtype.GongMarshallField(stage, "AggregationType"))
		pointersInitializesStatements.WriteString(ringtype.GongMarshallField(stage, "CurveMember"))
	}
	initRes = initializerStatements.String()
	ptrRes = pointersInitializesStatements.String()
	return
}
func (surfacepatcharraypropertytype *SurfacePatchArrayPropertyType) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(surfacepatcharraypropertytype.GongMarshallField(stage, "Name"))
		pointersInitializesStatements.WriteString(surfacepatcharraypropertytype.GongMarshallField(stage, "PolygonPatch"))
	}
	initRes = initializerStatements.String()
	ptrRes = pointersInitializesStatements.String()
	return
}
func (surfacetype *SurfaceType) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(surfacetype.GongMarshallField(stage, "Name"))
		initializerStatements.WriteString(surfacetype.GongMarshallField(stage, "Id"))
		initializerStatements.WriteString(surfacetype.GongMarshallField(stage, "Description"))
		pointersInitializesStatements.WriteString(surfacetype.GongMarshallField(stage, "Identifier"))
		pointersInitializesStatements.WriteString(surfacetype.GongMarshallField(stage, "NameXSD"))
		initializerStatements.WriteString(surfacetype.GongMarshallField(stage, "SrsName"))
		pointersInitializesStatements.WriteString(surfacetype.GongMarshallField(stage, "Patches"))
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
