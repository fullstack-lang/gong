// generated code - do not edit
package spatialrepresentation

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
	{{Identifier}} := (&spatialrepresentation.{{GeneratedStructName}}{Name: {{GeneratedFieldNameValue}}}).Stage(stage)`

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
			const fnSig = "func _(stage *spatialrepresentation.Stage) {"
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
			strings.HasSuffix(trimmedContentBeforeBrace, "func _(stage *spatialrepresentation.Stage) {") ||
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
	modelsPackageName := "github.com/fullstack-lang/gong/app/xsd/tests/aixm/go/models/spatialrepresentation"
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
		for _, abstractmd_spatialrepresentation_type := range __gong__sortStageSetInstances(stage.AbstractMD_SpatialRepresentation_Types, stage.AbstractMD_SpatialRepresentation_Type_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			abstractmd_spatialrepresentation_typeIdent := abstractmd_spatialrepresentation_type.GongGetIdentifier(stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&spatialrepresentation.AbstractMD_SpatialRepresentation_Type{Name: %s}).Stage(stage)", abstractmd_spatialrepresentation_typeIdent, __gong__toRawStringLiteral(abstractmd_spatialrepresentation_type.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", abstractmd_spatialrepresentation_typeIdent, __gong__toRawStringLiteral(abstractmd_spatialrepresentation_type.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Id = %s", abstractmd_spatialrepresentation_typeIdent, __gong__toRawStringLiteral(abstractmd_spatialrepresentation_type.Id)))
			values.WriteString(fmt.Sprintf("\n\t%s.Uuid = %s", abstractmd_spatialrepresentation_typeIdent, __gong__toRawStringLiteral(abstractmd_spatialrepresentation_type.Uuid)))
		}
	}
	if stage != nil {
		for _, md_cellgeometrycode_propertytype := range __gong__sortStageSetInstances(stage.MD_CellGeometryCode_PropertyTypes, stage.MD_CellGeometryCode_PropertyType_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			md_cellgeometrycode_propertytypeIdent := md_cellgeometrycode_propertytype.GongGetIdentifier(stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&spatialrepresentation.MD_CellGeometryCode_PropertyType{Name: %s}).Stage(stage)", md_cellgeometrycode_propertytypeIdent, __gong__toRawStringLiteral(md_cellgeometrycode_propertytype.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", md_cellgeometrycode_propertytypeIdent, __gong__toRawStringLiteral(md_cellgeometrycode_propertytype.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.NilReason = %s", md_cellgeometrycode_propertytypeIdent, __gong__toRawStringLiteral(md_cellgeometrycode_propertytype.NilReason)))
		}
	}
	if stage != nil {
		for _, md_dimensionnametypecode_propertytype := range __gong__sortStageSetInstances(stage.MD_DimensionNameTypeCode_PropertyTypes, stage.MD_DimensionNameTypeCode_PropertyType_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			md_dimensionnametypecode_propertytypeIdent := md_dimensionnametypecode_propertytype.GongGetIdentifier(stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&spatialrepresentation.MD_DimensionNameTypeCode_PropertyType{Name: %s}).Stage(stage)", md_dimensionnametypecode_propertytypeIdent, __gong__toRawStringLiteral(md_dimensionnametypecode_propertytype.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", md_dimensionnametypecode_propertytypeIdent, __gong__toRawStringLiteral(md_dimensionnametypecode_propertytype.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.NilReason = %s", md_dimensionnametypecode_propertytypeIdent, __gong__toRawStringLiteral(md_dimensionnametypecode_propertytype.NilReason)))
		}
	}
	if stage != nil {
		for _, md_dimension_propertytype := range __gong__sortStageSetInstances(stage.MD_Dimension_PropertyTypes, stage.MD_Dimension_PropertyType_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			md_dimension_propertytypeIdent := md_dimension_propertytype.GongGetIdentifier(stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&spatialrepresentation.MD_Dimension_PropertyType{Name: %s}).Stage(stage)", md_dimension_propertytypeIdent, __gong__toRawStringLiteral(md_dimension_propertytype.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", md_dimension_propertytypeIdent, __gong__toRawStringLiteral(md_dimension_propertytype.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.NilReason = %s", md_dimension_propertytypeIdent, __gong__toRawStringLiteral(md_dimension_propertytype.NilReason)))
			values.WriteString(fmt.Sprintf("\n\t%s.Uuidref = %s", md_dimension_propertytypeIdent, __gong__toRawStringLiteral(md_dimension_propertytype.Uuidref)))
			values.WriteString(fmt.Sprintf("\n\t%s.Type_XSD = %s", md_dimension_propertytypeIdent, __gong__toRawStringLiteral(md_dimension_propertytype.Type_XSD)))
			values.WriteString(fmt.Sprintf("\n\t%s.Href = %s", md_dimension_propertytypeIdent, __gong__toRawStringLiteral(md_dimension_propertytype.Href)))
			values.WriteString(fmt.Sprintf("\n\t%s.Role = %s", md_dimension_propertytypeIdent, __gong__toRawStringLiteral(md_dimension_propertytype.Role)))
			values.WriteString(fmt.Sprintf("\n\t%s.Arcrole = %s", md_dimension_propertytypeIdent, __gong__toRawStringLiteral(md_dimension_propertytype.Arcrole)))
			values.WriteString(fmt.Sprintf("\n\t%s.Title = %s", md_dimension_propertytypeIdent, __gong__toRawStringLiteral(md_dimension_propertytype.Title)))
			values.WriteString(fmt.Sprintf("\n\t%s.Show = %s", md_dimension_propertytypeIdent, __gong__toRawStringLiteral(md_dimension_propertytype.Show)))
			values.WriteString(fmt.Sprintf("\n\t%s.Actuate = %s", md_dimension_propertytypeIdent, __gong__toRawStringLiteral(md_dimension_propertytype.Actuate)))
			if md_dimension_propertytype.MD_Dimension != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := md_dimension_propertytype.MD_Dimension.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.MD_Dimension = %s", md_dimension_propertytypeIdent, targetIdent))
			}
		}
	}
	if stage != nil {
		for _, md_dimension_type := range __gong__sortStageSetInstances(stage.MD_Dimension_Types, stage.MD_Dimension_Type_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			md_dimension_typeIdent := md_dimension_type.GongGetIdentifier(stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&spatialrepresentation.MD_Dimension_Type{Name: %s}).Stage(stage)", md_dimension_typeIdent, __gong__toRawStringLiteral(md_dimension_type.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", md_dimension_typeIdent, __gong__toRawStringLiteral(md_dimension_type.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Id = %s", md_dimension_typeIdent, __gong__toRawStringLiteral(md_dimension_type.Id)))
			values.WriteString(fmt.Sprintf("\n\t%s.Uuid = %s", md_dimension_typeIdent, __gong__toRawStringLiteral(md_dimension_type.Uuid)))
			if md_dimension_type.DimensionName != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := md_dimension_type.DimensionName.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.DimensionName = %s", md_dimension_typeIdent, targetIdent))
			}
		}
	}
	if stage != nil {
		for _, md_geometricobjecttypecode_propertytype := range __gong__sortStageSetInstances(stage.MD_GeometricObjectTypeCode_PropertyTypes, stage.MD_GeometricObjectTypeCode_PropertyType_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			md_geometricobjecttypecode_propertytypeIdent := md_geometricobjecttypecode_propertytype.GongGetIdentifier(stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&spatialrepresentation.MD_GeometricObjectTypeCode_PropertyType{Name: %s}).Stage(stage)", md_geometricobjecttypecode_propertytypeIdent, __gong__toRawStringLiteral(md_geometricobjecttypecode_propertytype.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", md_geometricobjecttypecode_propertytypeIdent, __gong__toRawStringLiteral(md_geometricobjecttypecode_propertytype.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.NilReason = %s", md_geometricobjecttypecode_propertytypeIdent, __gong__toRawStringLiteral(md_geometricobjecttypecode_propertytype.NilReason)))
		}
	}
	if stage != nil {
		for _, md_geometricobjects_propertytype := range __gong__sortStageSetInstances(stage.MD_GeometricObjects_PropertyTypes, stage.MD_GeometricObjects_PropertyType_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			md_geometricobjects_propertytypeIdent := md_geometricobjects_propertytype.GongGetIdentifier(stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&spatialrepresentation.MD_GeometricObjects_PropertyType{Name: %s}).Stage(stage)", md_geometricobjects_propertytypeIdent, __gong__toRawStringLiteral(md_geometricobjects_propertytype.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", md_geometricobjects_propertytypeIdent, __gong__toRawStringLiteral(md_geometricobjects_propertytype.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.NilReason = %s", md_geometricobjects_propertytypeIdent, __gong__toRawStringLiteral(md_geometricobjects_propertytype.NilReason)))
			values.WriteString(fmt.Sprintf("\n\t%s.Uuidref = %s", md_geometricobjects_propertytypeIdent, __gong__toRawStringLiteral(md_geometricobjects_propertytype.Uuidref)))
			values.WriteString(fmt.Sprintf("\n\t%s.Type_XSD = %s", md_geometricobjects_propertytypeIdent, __gong__toRawStringLiteral(md_geometricobjects_propertytype.Type_XSD)))
			values.WriteString(fmt.Sprintf("\n\t%s.Href = %s", md_geometricobjects_propertytypeIdent, __gong__toRawStringLiteral(md_geometricobjects_propertytype.Href)))
			values.WriteString(fmt.Sprintf("\n\t%s.Role = %s", md_geometricobjects_propertytypeIdent, __gong__toRawStringLiteral(md_geometricobjects_propertytype.Role)))
			values.WriteString(fmt.Sprintf("\n\t%s.Arcrole = %s", md_geometricobjects_propertytypeIdent, __gong__toRawStringLiteral(md_geometricobjects_propertytype.Arcrole)))
			values.WriteString(fmt.Sprintf("\n\t%s.Title = %s", md_geometricobjects_propertytypeIdent, __gong__toRawStringLiteral(md_geometricobjects_propertytype.Title)))
			values.WriteString(fmt.Sprintf("\n\t%s.Show = %s", md_geometricobjects_propertytypeIdent, __gong__toRawStringLiteral(md_geometricobjects_propertytype.Show)))
			values.WriteString(fmt.Sprintf("\n\t%s.Actuate = %s", md_geometricobjects_propertytypeIdent, __gong__toRawStringLiteral(md_geometricobjects_propertytype.Actuate)))
			if md_geometricobjects_propertytype.MD_GeometricObjects != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := md_geometricobjects_propertytype.MD_GeometricObjects.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.MD_GeometricObjects = %s", md_geometricobjects_propertytypeIdent, targetIdent))
			}
		}
	}
	if stage != nil {
		for _, md_geometricobjects_type := range __gong__sortStageSetInstances(stage.MD_GeometricObjects_Types, stage.MD_GeometricObjects_Type_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			md_geometricobjects_typeIdent := md_geometricobjects_type.GongGetIdentifier(stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&spatialrepresentation.MD_GeometricObjects_Type{Name: %s}).Stage(stage)", md_geometricobjects_typeIdent, __gong__toRawStringLiteral(md_geometricobjects_type.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", md_geometricobjects_typeIdent, __gong__toRawStringLiteral(md_geometricobjects_type.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Id = %s", md_geometricobjects_typeIdent, __gong__toRawStringLiteral(md_geometricobjects_type.Id)))
			values.WriteString(fmt.Sprintf("\n\t%s.Uuid = %s", md_geometricobjects_typeIdent, __gong__toRawStringLiteral(md_geometricobjects_type.Uuid)))
			if md_geometricobjects_type.GeometricObjectType != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := md_geometricobjects_type.GeometricObjectType.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.GeometricObjectType = %s", md_geometricobjects_typeIdent, targetIdent))
			}
		}
	}
	if stage != nil {
		for _, md_georectified_propertytype := range __gong__sortStageSetInstances(stage.MD_Georectified_PropertyTypes, stage.MD_Georectified_PropertyType_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			md_georectified_propertytypeIdent := md_georectified_propertytype.GongGetIdentifier(stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&spatialrepresentation.MD_Georectified_PropertyType{Name: %s}).Stage(stage)", md_georectified_propertytypeIdent, __gong__toRawStringLiteral(md_georectified_propertytype.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", md_georectified_propertytypeIdent, __gong__toRawStringLiteral(md_georectified_propertytype.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.NilReason = %s", md_georectified_propertytypeIdent, __gong__toRawStringLiteral(md_georectified_propertytype.NilReason)))
			values.WriteString(fmt.Sprintf("\n\t%s.Uuidref = %s", md_georectified_propertytypeIdent, __gong__toRawStringLiteral(md_georectified_propertytype.Uuidref)))
			values.WriteString(fmt.Sprintf("\n\t%s.Type_XSD = %s", md_georectified_propertytypeIdent, __gong__toRawStringLiteral(md_georectified_propertytype.Type_XSD)))
			values.WriteString(fmt.Sprintf("\n\t%s.Href = %s", md_georectified_propertytypeIdent, __gong__toRawStringLiteral(md_georectified_propertytype.Href)))
			values.WriteString(fmt.Sprintf("\n\t%s.Role = %s", md_georectified_propertytypeIdent, __gong__toRawStringLiteral(md_georectified_propertytype.Role)))
			values.WriteString(fmt.Sprintf("\n\t%s.Arcrole = %s", md_georectified_propertytypeIdent, __gong__toRawStringLiteral(md_georectified_propertytype.Arcrole)))
			values.WriteString(fmt.Sprintf("\n\t%s.Title = %s", md_georectified_propertytypeIdent, __gong__toRawStringLiteral(md_georectified_propertytype.Title)))
			values.WriteString(fmt.Sprintf("\n\t%s.Show = %s", md_georectified_propertytypeIdent, __gong__toRawStringLiteral(md_georectified_propertytype.Show)))
			values.WriteString(fmt.Sprintf("\n\t%s.Actuate = %s", md_georectified_propertytypeIdent, __gong__toRawStringLiteral(md_georectified_propertytype.Actuate)))
			if md_georectified_propertytype.MD_Georectified != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := md_georectified_propertytype.MD_Georectified.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.MD_Georectified = %s", md_georectified_propertytypeIdent, targetIdent))
			}
		}
	}
	if stage != nil {
		for _, md_georectified_type := range __gong__sortStageSetInstances(stage.MD_Georectified_Types, stage.MD_Georectified_Type_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			md_georectified_typeIdent := md_georectified_type.GongGetIdentifier(stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&spatialrepresentation.MD_Georectified_Type{Name: %s}).Stage(stage)", md_georectified_typeIdent, __gong__toRawStringLiteral(md_georectified_type.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", md_georectified_typeIdent, __gong__toRawStringLiteral(md_georectified_type.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Id = %s", md_georectified_typeIdent, __gong__toRawStringLiteral(md_georectified_type.Id)))
			values.WriteString(fmt.Sprintf("\n\t%s.Uuid = %s", md_georectified_typeIdent, __gong__toRawStringLiteral(md_georectified_type.Uuid)))
			for _, elem := range md_georectified_type.AxisDimensionProperties {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := elem.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.AxisDimensionProperties = append(%s.AxisDimensionProperties, %s)", md_georectified_typeIdent, md_georectified_typeIdent, targetIdent))
			}
			if md_georectified_type.CellGeometry != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := md_georectified_type.CellGeometry.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.CellGeometry = %s", md_georectified_typeIdent, targetIdent))
			}
			if md_georectified_type.PointInPixel != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := md_georectified_type.PointInPixel.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.PointInPixel = %s", md_georectified_typeIdent, targetIdent))
			}
		}
	}
	if stage != nil {
		for _, md_georeferenceable_propertytype := range __gong__sortStageSetInstances(stage.MD_Georeferenceable_PropertyTypes, stage.MD_Georeferenceable_PropertyType_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			md_georeferenceable_propertytypeIdent := md_georeferenceable_propertytype.GongGetIdentifier(stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&spatialrepresentation.MD_Georeferenceable_PropertyType{Name: %s}).Stage(stage)", md_georeferenceable_propertytypeIdent, __gong__toRawStringLiteral(md_georeferenceable_propertytype.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", md_georeferenceable_propertytypeIdent, __gong__toRawStringLiteral(md_georeferenceable_propertytype.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.NilReason = %s", md_georeferenceable_propertytypeIdent, __gong__toRawStringLiteral(md_georeferenceable_propertytype.NilReason)))
			values.WriteString(fmt.Sprintf("\n\t%s.Uuidref = %s", md_georeferenceable_propertytypeIdent, __gong__toRawStringLiteral(md_georeferenceable_propertytype.Uuidref)))
			values.WriteString(fmt.Sprintf("\n\t%s.Type_XSD = %s", md_georeferenceable_propertytypeIdent, __gong__toRawStringLiteral(md_georeferenceable_propertytype.Type_XSD)))
			values.WriteString(fmt.Sprintf("\n\t%s.Href = %s", md_georeferenceable_propertytypeIdent, __gong__toRawStringLiteral(md_georeferenceable_propertytype.Href)))
			values.WriteString(fmt.Sprintf("\n\t%s.Role = %s", md_georeferenceable_propertytypeIdent, __gong__toRawStringLiteral(md_georeferenceable_propertytype.Role)))
			values.WriteString(fmt.Sprintf("\n\t%s.Arcrole = %s", md_georeferenceable_propertytypeIdent, __gong__toRawStringLiteral(md_georeferenceable_propertytype.Arcrole)))
			values.WriteString(fmt.Sprintf("\n\t%s.Title = %s", md_georeferenceable_propertytypeIdent, __gong__toRawStringLiteral(md_georeferenceable_propertytype.Title)))
			values.WriteString(fmt.Sprintf("\n\t%s.Show = %s", md_georeferenceable_propertytypeIdent, __gong__toRawStringLiteral(md_georeferenceable_propertytype.Show)))
			values.WriteString(fmt.Sprintf("\n\t%s.Actuate = %s", md_georeferenceable_propertytypeIdent, __gong__toRawStringLiteral(md_georeferenceable_propertytype.Actuate)))
			if md_georeferenceable_propertytype.MD_Georeferenceable != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := md_georeferenceable_propertytype.MD_Georeferenceable.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.MD_Georeferenceable = %s", md_georeferenceable_propertytypeIdent, targetIdent))
			}
		}
	}
	if stage != nil {
		for _, md_georeferenceable_type := range __gong__sortStageSetInstances(stage.MD_Georeferenceable_Types, stage.MD_Georeferenceable_Type_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			md_georeferenceable_typeIdent := md_georeferenceable_type.GongGetIdentifier(stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&spatialrepresentation.MD_Georeferenceable_Type{Name: %s}).Stage(stage)", md_georeferenceable_typeIdent, __gong__toRawStringLiteral(md_georeferenceable_type.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", md_georeferenceable_typeIdent, __gong__toRawStringLiteral(md_georeferenceable_type.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Id = %s", md_georeferenceable_typeIdent, __gong__toRawStringLiteral(md_georeferenceable_type.Id)))
			values.WriteString(fmt.Sprintf("\n\t%s.Uuid = %s", md_georeferenceable_typeIdent, __gong__toRawStringLiteral(md_georeferenceable_type.Uuid)))
			for _, elem := range md_georeferenceable_type.AxisDimensionProperties {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := elem.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.AxisDimensionProperties = append(%s.AxisDimensionProperties, %s)", md_georeferenceable_typeIdent, md_georeferenceable_typeIdent, targetIdent))
			}
			if md_georeferenceable_type.CellGeometry != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := md_georeferenceable_type.CellGeometry.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.CellGeometry = %s", md_georeferenceable_typeIdent, targetIdent))
			}
		}
	}
	if stage != nil {
		for _, md_gridspatialrepresentation_propertytype := range __gong__sortStageSetInstances(stage.MD_GridSpatialRepresentation_PropertyTypes, stage.MD_GridSpatialRepresentation_PropertyType_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			md_gridspatialrepresentation_propertytypeIdent := md_gridspatialrepresentation_propertytype.GongGetIdentifier(stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&spatialrepresentation.MD_GridSpatialRepresentation_PropertyType{Name: %s}).Stage(stage)", md_gridspatialrepresentation_propertytypeIdent, __gong__toRawStringLiteral(md_gridspatialrepresentation_propertytype.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", md_gridspatialrepresentation_propertytypeIdent, __gong__toRawStringLiteral(md_gridspatialrepresentation_propertytype.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.NilReason = %s", md_gridspatialrepresentation_propertytypeIdent, __gong__toRawStringLiteral(md_gridspatialrepresentation_propertytype.NilReason)))
			values.WriteString(fmt.Sprintf("\n\t%s.Uuidref = %s", md_gridspatialrepresentation_propertytypeIdent, __gong__toRawStringLiteral(md_gridspatialrepresentation_propertytype.Uuidref)))
			values.WriteString(fmt.Sprintf("\n\t%s.Type_XSD = %s", md_gridspatialrepresentation_propertytypeIdent, __gong__toRawStringLiteral(md_gridspatialrepresentation_propertytype.Type_XSD)))
			values.WriteString(fmt.Sprintf("\n\t%s.Href = %s", md_gridspatialrepresentation_propertytypeIdent, __gong__toRawStringLiteral(md_gridspatialrepresentation_propertytype.Href)))
			values.WriteString(fmt.Sprintf("\n\t%s.Role = %s", md_gridspatialrepresentation_propertytypeIdent, __gong__toRawStringLiteral(md_gridspatialrepresentation_propertytype.Role)))
			values.WriteString(fmt.Sprintf("\n\t%s.Arcrole = %s", md_gridspatialrepresentation_propertytypeIdent, __gong__toRawStringLiteral(md_gridspatialrepresentation_propertytype.Arcrole)))
			values.WriteString(fmt.Sprintf("\n\t%s.Title = %s", md_gridspatialrepresentation_propertytypeIdent, __gong__toRawStringLiteral(md_gridspatialrepresentation_propertytype.Title)))
			values.WriteString(fmt.Sprintf("\n\t%s.Show = %s", md_gridspatialrepresentation_propertytypeIdent, __gong__toRawStringLiteral(md_gridspatialrepresentation_propertytype.Show)))
			values.WriteString(fmt.Sprintf("\n\t%s.Actuate = %s", md_gridspatialrepresentation_propertytypeIdent, __gong__toRawStringLiteral(md_gridspatialrepresentation_propertytype.Actuate)))
			if md_gridspatialrepresentation_propertytype.MD_GridSpatialRepresentation != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := md_gridspatialrepresentation_propertytype.MD_GridSpatialRepresentation.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.MD_GridSpatialRepresentation = %s", md_gridspatialrepresentation_propertytypeIdent, targetIdent))
			}
			if md_gridspatialrepresentation_propertytype.MD_Georeferenceable != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := md_gridspatialrepresentation_propertytype.MD_Georeferenceable.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.MD_Georeferenceable = %s", md_gridspatialrepresentation_propertytypeIdent, targetIdent))
			}
			if md_gridspatialrepresentation_propertytype.MD_Georectified != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := md_gridspatialrepresentation_propertytype.MD_Georectified.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.MD_Georectified = %s", md_gridspatialrepresentation_propertytypeIdent, targetIdent))
			}
		}
	}
	if stage != nil {
		for _, md_gridspatialrepresentation_type := range __gong__sortStageSetInstances(stage.MD_GridSpatialRepresentation_Types, stage.MD_GridSpatialRepresentation_Type_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			md_gridspatialrepresentation_typeIdent := md_gridspatialrepresentation_type.GongGetIdentifier(stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&spatialrepresentation.MD_GridSpatialRepresentation_Type{Name: %s}).Stage(stage)", md_gridspatialrepresentation_typeIdent, __gong__toRawStringLiteral(md_gridspatialrepresentation_type.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", md_gridspatialrepresentation_typeIdent, __gong__toRawStringLiteral(md_gridspatialrepresentation_type.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Id = %s", md_gridspatialrepresentation_typeIdent, __gong__toRawStringLiteral(md_gridspatialrepresentation_type.Id)))
			values.WriteString(fmt.Sprintf("\n\t%s.Uuid = %s", md_gridspatialrepresentation_typeIdent, __gong__toRawStringLiteral(md_gridspatialrepresentation_type.Uuid)))
			for _, elem := range md_gridspatialrepresentation_type.AxisDimensionProperties {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := elem.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.AxisDimensionProperties = append(%s.AxisDimensionProperties, %s)", md_gridspatialrepresentation_typeIdent, md_gridspatialrepresentation_typeIdent, targetIdent))
			}
			if md_gridspatialrepresentation_type.CellGeometry != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := md_gridspatialrepresentation_type.CellGeometry.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.CellGeometry = %s", md_gridspatialrepresentation_typeIdent, targetIdent))
			}
		}
	}
	if stage != nil {
		for _, md_pixelorientationcode_propertytype := range __gong__sortStageSetInstances(stage.MD_PixelOrientationCode_PropertyTypes, stage.MD_PixelOrientationCode_PropertyType_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			md_pixelorientationcode_propertytypeIdent := md_pixelorientationcode_propertytype.GongGetIdentifier(stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&spatialrepresentation.MD_PixelOrientationCode_PropertyType{Name: %s}).Stage(stage)", md_pixelorientationcode_propertytypeIdent, __gong__toRawStringLiteral(md_pixelorientationcode_propertytype.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", md_pixelorientationcode_propertytypeIdent, __gong__toRawStringLiteral(md_pixelorientationcode_propertytype.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.NilReason = %s", md_pixelorientationcode_propertytypeIdent, __gong__toRawStringLiteral(md_pixelorientationcode_propertytype.NilReason)))
			if md_pixelorientationcode_propertytype.MD_PixelOrientationCode.ToCodeString() != "" {
				values.WriteString(fmt.Sprintf("\n\t%s.MD_PixelOrientationCode = spatialrepresentation.%s", md_pixelorientationcode_propertytypeIdent, md_pixelorientationcode_propertytype.MD_PixelOrientationCode.ToCodeString()))
			} else {
				values.WriteString(fmt.Sprintf("\n\t%s.MD_PixelOrientationCode = \"\"", md_pixelorientationcode_propertytypeIdent))
			}
		}
	}
	if stage != nil {
		for _, md_spatialrepresentation_propertytype := range __gong__sortStageSetInstances(stage.MD_SpatialRepresentation_PropertyTypes, stage.MD_SpatialRepresentation_PropertyType_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			md_spatialrepresentation_propertytypeIdent := md_spatialrepresentation_propertytype.GongGetIdentifier(stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&spatialrepresentation.MD_SpatialRepresentation_PropertyType{Name: %s}).Stage(stage)", md_spatialrepresentation_propertytypeIdent, __gong__toRawStringLiteral(md_spatialrepresentation_propertytype.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", md_spatialrepresentation_propertytypeIdent, __gong__toRawStringLiteral(md_spatialrepresentation_propertytype.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.NilReason = %s", md_spatialrepresentation_propertytypeIdent, __gong__toRawStringLiteral(md_spatialrepresentation_propertytype.NilReason)))
			values.WriteString(fmt.Sprintf("\n\t%s.Uuidref = %s", md_spatialrepresentation_propertytypeIdent, __gong__toRawStringLiteral(md_spatialrepresentation_propertytype.Uuidref)))
			values.WriteString(fmt.Sprintf("\n\t%s.Type_XSD = %s", md_spatialrepresentation_propertytypeIdent, __gong__toRawStringLiteral(md_spatialrepresentation_propertytype.Type_XSD)))
			values.WriteString(fmt.Sprintf("\n\t%s.Href = %s", md_spatialrepresentation_propertytypeIdent, __gong__toRawStringLiteral(md_spatialrepresentation_propertytype.Href)))
			values.WriteString(fmt.Sprintf("\n\t%s.Role = %s", md_spatialrepresentation_propertytypeIdent, __gong__toRawStringLiteral(md_spatialrepresentation_propertytype.Role)))
			values.WriteString(fmt.Sprintf("\n\t%s.Arcrole = %s", md_spatialrepresentation_propertytypeIdent, __gong__toRawStringLiteral(md_spatialrepresentation_propertytype.Arcrole)))
			values.WriteString(fmt.Sprintf("\n\t%s.Title = %s", md_spatialrepresentation_propertytypeIdent, __gong__toRawStringLiteral(md_spatialrepresentation_propertytype.Title)))
			values.WriteString(fmt.Sprintf("\n\t%s.Show = %s", md_spatialrepresentation_propertytypeIdent, __gong__toRawStringLiteral(md_spatialrepresentation_propertytype.Show)))
			values.WriteString(fmt.Sprintf("\n\t%s.Actuate = %s", md_spatialrepresentation_propertytypeIdent, __gong__toRawStringLiteral(md_spatialrepresentation_propertytype.Actuate)))
			if md_spatialrepresentation_propertytype.MD_GridSpatialRepresentation != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := md_spatialrepresentation_propertytype.MD_GridSpatialRepresentation.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.MD_GridSpatialRepresentation = %s", md_spatialrepresentation_propertytypeIdent, targetIdent))
			}
			if md_spatialrepresentation_propertytype.MD_Georeferenceable != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := md_spatialrepresentation_propertytype.MD_Georeferenceable.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.MD_Georeferenceable = %s", md_spatialrepresentation_propertytypeIdent, targetIdent))
			}
			if md_spatialrepresentation_propertytype.MD_Georectified != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := md_spatialrepresentation_propertytype.MD_Georectified.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.MD_Georectified = %s", md_spatialrepresentation_propertytypeIdent, targetIdent))
			}
			if md_spatialrepresentation_propertytype.MD_VectorSpatialRepresentation != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := md_spatialrepresentation_propertytype.MD_VectorSpatialRepresentation.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.MD_VectorSpatialRepresentation = %s", md_spatialrepresentation_propertytypeIdent, targetIdent))
			}
		}
	}
	if stage != nil {
		for _, md_topologylevelcode_propertytype := range __gong__sortStageSetInstances(stage.MD_TopologyLevelCode_PropertyTypes, stage.MD_TopologyLevelCode_PropertyType_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			md_topologylevelcode_propertytypeIdent := md_topologylevelcode_propertytype.GongGetIdentifier(stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&spatialrepresentation.MD_TopologyLevelCode_PropertyType{Name: %s}).Stage(stage)", md_topologylevelcode_propertytypeIdent, __gong__toRawStringLiteral(md_topologylevelcode_propertytype.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", md_topologylevelcode_propertytypeIdent, __gong__toRawStringLiteral(md_topologylevelcode_propertytype.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.NilReason = %s", md_topologylevelcode_propertytypeIdent, __gong__toRawStringLiteral(md_topologylevelcode_propertytype.NilReason)))
		}
	}
	if stage != nil {
		for _, md_vectorspatialrepresentation_propertytype := range __gong__sortStageSetInstances(stage.MD_VectorSpatialRepresentation_PropertyTypes, stage.MD_VectorSpatialRepresentation_PropertyType_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			md_vectorspatialrepresentation_propertytypeIdent := md_vectorspatialrepresentation_propertytype.GongGetIdentifier(stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&spatialrepresentation.MD_VectorSpatialRepresentation_PropertyType{Name: %s}).Stage(stage)", md_vectorspatialrepresentation_propertytypeIdent, __gong__toRawStringLiteral(md_vectorspatialrepresentation_propertytype.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", md_vectorspatialrepresentation_propertytypeIdent, __gong__toRawStringLiteral(md_vectorspatialrepresentation_propertytype.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.NilReason = %s", md_vectorspatialrepresentation_propertytypeIdent, __gong__toRawStringLiteral(md_vectorspatialrepresentation_propertytype.NilReason)))
			values.WriteString(fmt.Sprintf("\n\t%s.Uuidref = %s", md_vectorspatialrepresentation_propertytypeIdent, __gong__toRawStringLiteral(md_vectorspatialrepresentation_propertytype.Uuidref)))
			values.WriteString(fmt.Sprintf("\n\t%s.Type_XSD = %s", md_vectorspatialrepresentation_propertytypeIdent, __gong__toRawStringLiteral(md_vectorspatialrepresentation_propertytype.Type_XSD)))
			values.WriteString(fmt.Sprintf("\n\t%s.Href = %s", md_vectorspatialrepresentation_propertytypeIdent, __gong__toRawStringLiteral(md_vectorspatialrepresentation_propertytype.Href)))
			values.WriteString(fmt.Sprintf("\n\t%s.Role = %s", md_vectorspatialrepresentation_propertytypeIdent, __gong__toRawStringLiteral(md_vectorspatialrepresentation_propertytype.Role)))
			values.WriteString(fmt.Sprintf("\n\t%s.Arcrole = %s", md_vectorspatialrepresentation_propertytypeIdent, __gong__toRawStringLiteral(md_vectorspatialrepresentation_propertytype.Arcrole)))
			values.WriteString(fmt.Sprintf("\n\t%s.Title = %s", md_vectorspatialrepresentation_propertytypeIdent, __gong__toRawStringLiteral(md_vectorspatialrepresentation_propertytype.Title)))
			values.WriteString(fmt.Sprintf("\n\t%s.Show = %s", md_vectorspatialrepresentation_propertytypeIdent, __gong__toRawStringLiteral(md_vectorspatialrepresentation_propertytype.Show)))
			values.WriteString(fmt.Sprintf("\n\t%s.Actuate = %s", md_vectorspatialrepresentation_propertytypeIdent, __gong__toRawStringLiteral(md_vectorspatialrepresentation_propertytype.Actuate)))
			if md_vectorspatialrepresentation_propertytype.MD_VectorSpatialRepresentation != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := md_vectorspatialrepresentation_propertytype.MD_VectorSpatialRepresentation.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.MD_VectorSpatialRepresentation = %s", md_vectorspatialrepresentation_propertytypeIdent, targetIdent))
			}
		}
	}
	if stage != nil {
		for _, md_vectorspatialrepresentation_type := range __gong__sortStageSetInstances(stage.MD_VectorSpatialRepresentation_Types, stage.MD_VectorSpatialRepresentation_Type_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			md_vectorspatialrepresentation_typeIdent := md_vectorspatialrepresentation_type.GongGetIdentifier(stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&spatialrepresentation.MD_VectorSpatialRepresentation_Type{Name: %s}).Stage(stage)", md_vectorspatialrepresentation_typeIdent, __gong__toRawStringLiteral(md_vectorspatialrepresentation_type.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", md_vectorspatialrepresentation_typeIdent, __gong__toRawStringLiteral(md_vectorspatialrepresentation_type.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Id = %s", md_vectorspatialrepresentation_typeIdent, __gong__toRawStringLiteral(md_vectorspatialrepresentation_type.Id)))
			values.WriteString(fmt.Sprintf("\n\t%s.Uuid = %s", md_vectorspatialrepresentation_typeIdent, __gong__toRawStringLiteral(md_vectorspatialrepresentation_type.Uuid)))
			if md_vectorspatialrepresentation_type.TopologyLevel != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := md_vectorspatialrepresentation_type.TopologyLevel.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.TopologyLevel = %s", md_vectorspatialrepresentation_typeIdent, targetIdent))
			}
			for _, elem := range md_vectorspatialrepresentation_type.GeometricObjects {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := elem.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.GeometricObjects = append(%s.GeometricObjects, %s)", md_vectorspatialrepresentation_typeIdent, md_vectorspatialrepresentation_typeIdent, targetIdent))
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
go run github.com/fullstack-lang/gong/app/xsd/tests/aixm/go/models/spatialrepresentation/probe/cmd/probe@%s stage.go
*/
package %s

import (
	"slices"
	"time"

	"github.com/fullstack-lang/gong/app/xsd/tests/aixm/go/models/spatialrepresentation"%s
)

var (
	_ time.Time
	_ = slices.Index[[]int, int]

	_ *spatialrepresentation.Stage
)%s

// function will stage objects across all coordinated stages
func _(stage *spatialrepresentation.Stage) {

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

	if modelsPackageName != "" && modelsPackageName != "github.com/fullstack-lang/gong/app/xsd/tests/aixm/go/models/spatialrepresentation" {
		res = strings.ReplaceAll(res, "\"github.com/fullstack-lang/gong/app/xsd/tests/aixm/go/models/spatialrepresentation\"", fmt.Sprintf("%q", modelsPackageName))
	}

	return res, nil
}

// insertion point for marshall field methods
func (abstractmd_spatialrepresentation_type *AbstractMD_SpatialRepresentation_Type) GongMarshallField(stage *Stage, fieldName string) (res string) {
	ident := abstractmd_spatialrepresentation_type.GongGetIdentifier(stage)
	_ = ident
	switch fieldName {
	case "Name":
		res = __gong__marshallString(ident, "Name", abstractmd_spatialrepresentation_type.Name)
	case "Id":
		res = __gong__marshallString(ident, "Id", abstractmd_spatialrepresentation_type.Id)
	case "Uuid":
		res = __gong__marshallString(ident, "Uuid", abstractmd_spatialrepresentation_type.Uuid)

	default:
		log.Panicf("Unknown field %s for Gongstruct AbstractMD_SpatialRepresentation_Type", fieldName)
	}
	return
}

func (md_cellgeometrycode_propertytype *MD_CellGeometryCode_PropertyType) GongMarshallField(stage *Stage, fieldName string) (res string) {
	ident := md_cellgeometrycode_propertytype.GongGetIdentifier(stage)
	_ = ident
	switch fieldName {
	case "Name":
		res = __gong__marshallString(ident, "Name", md_cellgeometrycode_propertytype.Name)
	case "NilReason":
		res = __gong__marshallString(ident, "NilReason", md_cellgeometrycode_propertytype.NilReason)

	default:
		log.Panicf("Unknown field %s for Gongstruct MD_CellGeometryCode_PropertyType", fieldName)
	}
	return
}

func (md_dimensionnametypecode_propertytype *MD_DimensionNameTypeCode_PropertyType) GongMarshallField(stage *Stage, fieldName string) (res string) {
	ident := md_dimensionnametypecode_propertytype.GongGetIdentifier(stage)
	_ = ident
	switch fieldName {
	case "Name":
		res = __gong__marshallString(ident, "Name", md_dimensionnametypecode_propertytype.Name)
	case "NilReason":
		res = __gong__marshallString(ident, "NilReason", md_dimensionnametypecode_propertytype.NilReason)

	default:
		log.Panicf("Unknown field %s for Gongstruct MD_DimensionNameTypeCode_PropertyType", fieldName)
	}
	return
}

func (md_dimension_propertytype *MD_Dimension_PropertyType) GongMarshallField(stage *Stage, fieldName string) (res string) {
	ident := md_dimension_propertytype.GongGetIdentifier(stage)
	_ = ident
	switch fieldName {
	case "Name":
		res = __gong__marshallString(ident, "Name", md_dimension_propertytype.Name)
	case "NilReason":
		res = __gong__marshallString(ident, "NilReason", md_dimension_propertytype.NilReason)
	case "Uuidref":
		res = __gong__marshallString(ident, "Uuidref", md_dimension_propertytype.Uuidref)
	case "Type_XSD":
		res = __gong__marshallString(ident, "Type_XSD", md_dimension_propertytype.Type_XSD)
	case "Href":
		res = __gong__marshallString(ident, "Href", md_dimension_propertytype.Href)
	case "Role":
		res = __gong__marshallString(ident, "Role", md_dimension_propertytype.Role)
	case "Arcrole":
		res = __gong__marshallString(ident, "Arcrole", md_dimension_propertytype.Arcrole)
	case "Title":
		res = __gong__marshallString(ident, "Title", md_dimension_propertytype.Title)
	case "Show":
		res = __gong__marshallString(ident, "Show", md_dimension_propertytype.Show)
	case "Actuate":
		res = __gong__marshallString(ident, "Actuate", md_dimension_propertytype.Actuate)

	case "MD_Dimension":
		if md_dimension_propertytype.MD_Dimension != nil {
			res = __gong__marshallPointer(ident, "MD_Dimension", md_dimension_propertytype.MD_Dimension.GongGetIdentifier(stage))
		} else {
			res = __gong__marshallPointer(ident, "MD_Dimension", "nil")
		}
	default:
		log.Panicf("Unknown field %s for Gongstruct MD_Dimension_PropertyType", fieldName)
	}
	return
}

func (md_dimension_type *MD_Dimension_Type) GongMarshallField(stage *Stage, fieldName string) (res string) {
	ident := md_dimension_type.GongGetIdentifier(stage)
	_ = ident
	switch fieldName {
	case "Name":
		res = __gong__marshallString(ident, "Name", md_dimension_type.Name)
	case "Id":
		res = __gong__marshallString(ident, "Id", md_dimension_type.Id)
	case "Uuid":
		res = __gong__marshallString(ident, "Uuid", md_dimension_type.Uuid)

	case "DimensionName":
		if md_dimension_type.DimensionName != nil {
			res = __gong__marshallPointer(ident, "DimensionName", md_dimension_type.DimensionName.GongGetIdentifier(stage))
		} else {
			res = __gong__marshallPointer(ident, "DimensionName", "nil")
		}
	default:
		log.Panicf("Unknown field %s for Gongstruct MD_Dimension_Type", fieldName)
	}
	return
}

func (md_geometricobjecttypecode_propertytype *MD_GeometricObjectTypeCode_PropertyType) GongMarshallField(stage *Stage, fieldName string) (res string) {
	ident := md_geometricobjecttypecode_propertytype.GongGetIdentifier(stage)
	_ = ident
	switch fieldName {
	case "Name":
		res = __gong__marshallString(ident, "Name", md_geometricobjecttypecode_propertytype.Name)
	case "NilReason":
		res = __gong__marshallString(ident, "NilReason", md_geometricobjecttypecode_propertytype.NilReason)

	default:
		log.Panicf("Unknown field %s for Gongstruct MD_GeometricObjectTypeCode_PropertyType", fieldName)
	}
	return
}

func (md_geometricobjects_propertytype *MD_GeometricObjects_PropertyType) GongMarshallField(stage *Stage, fieldName string) (res string) {
	ident := md_geometricobjects_propertytype.GongGetIdentifier(stage)
	_ = ident
	switch fieldName {
	case "Name":
		res = __gong__marshallString(ident, "Name", md_geometricobjects_propertytype.Name)
	case "NilReason":
		res = __gong__marshallString(ident, "NilReason", md_geometricobjects_propertytype.NilReason)
	case "Uuidref":
		res = __gong__marshallString(ident, "Uuidref", md_geometricobjects_propertytype.Uuidref)
	case "Type_XSD":
		res = __gong__marshallString(ident, "Type_XSD", md_geometricobjects_propertytype.Type_XSD)
	case "Href":
		res = __gong__marshallString(ident, "Href", md_geometricobjects_propertytype.Href)
	case "Role":
		res = __gong__marshallString(ident, "Role", md_geometricobjects_propertytype.Role)
	case "Arcrole":
		res = __gong__marshallString(ident, "Arcrole", md_geometricobjects_propertytype.Arcrole)
	case "Title":
		res = __gong__marshallString(ident, "Title", md_geometricobjects_propertytype.Title)
	case "Show":
		res = __gong__marshallString(ident, "Show", md_geometricobjects_propertytype.Show)
	case "Actuate":
		res = __gong__marshallString(ident, "Actuate", md_geometricobjects_propertytype.Actuate)

	case "MD_GeometricObjects":
		if md_geometricobjects_propertytype.MD_GeometricObjects != nil {
			res = __gong__marshallPointer(ident, "MD_GeometricObjects", md_geometricobjects_propertytype.MD_GeometricObjects.GongGetIdentifier(stage))
		} else {
			res = __gong__marshallPointer(ident, "MD_GeometricObjects", "nil")
		}
	default:
		log.Panicf("Unknown field %s for Gongstruct MD_GeometricObjects_PropertyType", fieldName)
	}
	return
}

func (md_geometricobjects_type *MD_GeometricObjects_Type) GongMarshallField(stage *Stage, fieldName string) (res string) {
	ident := md_geometricobjects_type.GongGetIdentifier(stage)
	_ = ident
	switch fieldName {
	case "Name":
		res = __gong__marshallString(ident, "Name", md_geometricobjects_type.Name)
	case "Id":
		res = __gong__marshallString(ident, "Id", md_geometricobjects_type.Id)
	case "Uuid":
		res = __gong__marshallString(ident, "Uuid", md_geometricobjects_type.Uuid)

	case "GeometricObjectType":
		if md_geometricobjects_type.GeometricObjectType != nil {
			res = __gong__marshallPointer(ident, "GeometricObjectType", md_geometricobjects_type.GeometricObjectType.GongGetIdentifier(stage))
		} else {
			res = __gong__marshallPointer(ident, "GeometricObjectType", "nil")
		}
	default:
		log.Panicf("Unknown field %s for Gongstruct MD_GeometricObjects_Type", fieldName)
	}
	return
}

func (md_georectified_propertytype *MD_Georectified_PropertyType) GongMarshallField(stage *Stage, fieldName string) (res string) {
	ident := md_georectified_propertytype.GongGetIdentifier(stage)
	_ = ident
	switch fieldName {
	case "Name":
		res = __gong__marshallString(ident, "Name", md_georectified_propertytype.Name)
	case "NilReason":
		res = __gong__marshallString(ident, "NilReason", md_georectified_propertytype.NilReason)
	case "Uuidref":
		res = __gong__marshallString(ident, "Uuidref", md_georectified_propertytype.Uuidref)
	case "Type_XSD":
		res = __gong__marshallString(ident, "Type_XSD", md_georectified_propertytype.Type_XSD)
	case "Href":
		res = __gong__marshallString(ident, "Href", md_georectified_propertytype.Href)
	case "Role":
		res = __gong__marshallString(ident, "Role", md_georectified_propertytype.Role)
	case "Arcrole":
		res = __gong__marshallString(ident, "Arcrole", md_georectified_propertytype.Arcrole)
	case "Title":
		res = __gong__marshallString(ident, "Title", md_georectified_propertytype.Title)
	case "Show":
		res = __gong__marshallString(ident, "Show", md_georectified_propertytype.Show)
	case "Actuate":
		res = __gong__marshallString(ident, "Actuate", md_georectified_propertytype.Actuate)

	case "MD_Georectified":
		if md_georectified_propertytype.MD_Georectified != nil {
			res = __gong__marshallPointer(ident, "MD_Georectified", md_georectified_propertytype.MD_Georectified.GongGetIdentifier(stage))
		} else {
			res = __gong__marshallPointer(ident, "MD_Georectified", "nil")
		}
	default:
		log.Panicf("Unknown field %s for Gongstruct MD_Georectified_PropertyType", fieldName)
	}
	return
}

func (md_georectified_type *MD_Georectified_Type) GongMarshallField(stage *Stage, fieldName string) (res string) {
	ident := md_georectified_type.GongGetIdentifier(stage)
	_ = ident
	switch fieldName {
	case "Name":
		res = __gong__marshallString(ident, "Name", md_georectified_type.Name)
	case "Id":
		res = __gong__marshallString(ident, "Id", md_georectified_type.Id)
	case "Uuid":
		res = __gong__marshallString(ident, "Uuid", md_georectified_type.Uuid)

	case "AxisDimensionProperties":
		var sb strings.Builder
		for _, _md_dimension_propertytype := range md_georectified_type.AxisDimensionProperties {
			sb.WriteString(__gong__marshallSliceOfPointers(ident, "AxisDimensionProperties", _md_dimension_propertytype.GongGetIdentifier(stage)))
		}
		res = sb.String()
	case "CellGeometry":
		if md_georectified_type.CellGeometry != nil {
			res = __gong__marshallPointer(ident, "CellGeometry", md_georectified_type.CellGeometry.GongGetIdentifier(stage))
		} else {
			res = __gong__marshallPointer(ident, "CellGeometry", "nil")
		}
	case "PointInPixel":
		if md_georectified_type.PointInPixel != nil {
			res = __gong__marshallPointer(ident, "PointInPixel", md_georectified_type.PointInPixel.GongGetIdentifier(stage))
		} else {
			res = __gong__marshallPointer(ident, "PointInPixel", "nil")
		}
	default:
		log.Panicf("Unknown field %s for Gongstruct MD_Georectified_Type", fieldName)
	}
	return
}

func (md_georeferenceable_propertytype *MD_Georeferenceable_PropertyType) GongMarshallField(stage *Stage, fieldName string) (res string) {
	ident := md_georeferenceable_propertytype.GongGetIdentifier(stage)
	_ = ident
	switch fieldName {
	case "Name":
		res = __gong__marshallString(ident, "Name", md_georeferenceable_propertytype.Name)
	case "NilReason":
		res = __gong__marshallString(ident, "NilReason", md_georeferenceable_propertytype.NilReason)
	case "Uuidref":
		res = __gong__marshallString(ident, "Uuidref", md_georeferenceable_propertytype.Uuidref)
	case "Type_XSD":
		res = __gong__marshallString(ident, "Type_XSD", md_georeferenceable_propertytype.Type_XSD)
	case "Href":
		res = __gong__marshallString(ident, "Href", md_georeferenceable_propertytype.Href)
	case "Role":
		res = __gong__marshallString(ident, "Role", md_georeferenceable_propertytype.Role)
	case "Arcrole":
		res = __gong__marshallString(ident, "Arcrole", md_georeferenceable_propertytype.Arcrole)
	case "Title":
		res = __gong__marshallString(ident, "Title", md_georeferenceable_propertytype.Title)
	case "Show":
		res = __gong__marshallString(ident, "Show", md_georeferenceable_propertytype.Show)
	case "Actuate":
		res = __gong__marshallString(ident, "Actuate", md_georeferenceable_propertytype.Actuate)

	case "MD_Georeferenceable":
		if md_georeferenceable_propertytype.MD_Georeferenceable != nil {
			res = __gong__marshallPointer(ident, "MD_Georeferenceable", md_georeferenceable_propertytype.MD_Georeferenceable.GongGetIdentifier(stage))
		} else {
			res = __gong__marshallPointer(ident, "MD_Georeferenceable", "nil")
		}
	default:
		log.Panicf("Unknown field %s for Gongstruct MD_Georeferenceable_PropertyType", fieldName)
	}
	return
}

func (md_georeferenceable_type *MD_Georeferenceable_Type) GongMarshallField(stage *Stage, fieldName string) (res string) {
	ident := md_georeferenceable_type.GongGetIdentifier(stage)
	_ = ident
	switch fieldName {
	case "Name":
		res = __gong__marshallString(ident, "Name", md_georeferenceable_type.Name)
	case "Id":
		res = __gong__marshallString(ident, "Id", md_georeferenceable_type.Id)
	case "Uuid":
		res = __gong__marshallString(ident, "Uuid", md_georeferenceable_type.Uuid)

	case "AxisDimensionProperties":
		var sb strings.Builder
		for _, _md_dimension_propertytype := range md_georeferenceable_type.AxisDimensionProperties {
			sb.WriteString(__gong__marshallSliceOfPointers(ident, "AxisDimensionProperties", _md_dimension_propertytype.GongGetIdentifier(stage)))
		}
		res = sb.String()
	case "CellGeometry":
		if md_georeferenceable_type.CellGeometry != nil {
			res = __gong__marshallPointer(ident, "CellGeometry", md_georeferenceable_type.CellGeometry.GongGetIdentifier(stage))
		} else {
			res = __gong__marshallPointer(ident, "CellGeometry", "nil")
		}
	default:
		log.Panicf("Unknown field %s for Gongstruct MD_Georeferenceable_Type", fieldName)
	}
	return
}

func (md_gridspatialrepresentation_propertytype *MD_GridSpatialRepresentation_PropertyType) GongMarshallField(stage *Stage, fieldName string) (res string) {
	ident := md_gridspatialrepresentation_propertytype.GongGetIdentifier(stage)
	_ = ident
	switch fieldName {
	case "Name":
		res = __gong__marshallString(ident, "Name", md_gridspatialrepresentation_propertytype.Name)
	case "NilReason":
		res = __gong__marshallString(ident, "NilReason", md_gridspatialrepresentation_propertytype.NilReason)
	case "Uuidref":
		res = __gong__marshallString(ident, "Uuidref", md_gridspatialrepresentation_propertytype.Uuidref)
	case "Type_XSD":
		res = __gong__marshallString(ident, "Type_XSD", md_gridspatialrepresentation_propertytype.Type_XSD)
	case "Href":
		res = __gong__marshallString(ident, "Href", md_gridspatialrepresentation_propertytype.Href)
	case "Role":
		res = __gong__marshallString(ident, "Role", md_gridspatialrepresentation_propertytype.Role)
	case "Arcrole":
		res = __gong__marshallString(ident, "Arcrole", md_gridspatialrepresentation_propertytype.Arcrole)
	case "Title":
		res = __gong__marshallString(ident, "Title", md_gridspatialrepresentation_propertytype.Title)
	case "Show":
		res = __gong__marshallString(ident, "Show", md_gridspatialrepresentation_propertytype.Show)
	case "Actuate":
		res = __gong__marshallString(ident, "Actuate", md_gridspatialrepresentation_propertytype.Actuate)

	case "MD_GridSpatialRepresentation":
		if md_gridspatialrepresentation_propertytype.MD_GridSpatialRepresentation != nil {
			res = __gong__marshallPointer(ident, "MD_GridSpatialRepresentation", md_gridspatialrepresentation_propertytype.MD_GridSpatialRepresentation.GongGetIdentifier(stage))
		} else {
			res = __gong__marshallPointer(ident, "MD_GridSpatialRepresentation", "nil")
		}
	case "MD_Georeferenceable":
		if md_gridspatialrepresentation_propertytype.MD_Georeferenceable != nil {
			res = __gong__marshallPointer(ident, "MD_Georeferenceable", md_gridspatialrepresentation_propertytype.MD_Georeferenceable.GongGetIdentifier(stage))
		} else {
			res = __gong__marshallPointer(ident, "MD_Georeferenceable", "nil")
		}
	case "MD_Georectified":
		if md_gridspatialrepresentation_propertytype.MD_Georectified != nil {
			res = __gong__marshallPointer(ident, "MD_Georectified", md_gridspatialrepresentation_propertytype.MD_Georectified.GongGetIdentifier(stage))
		} else {
			res = __gong__marshallPointer(ident, "MD_Georectified", "nil")
		}
	default:
		log.Panicf("Unknown field %s for Gongstruct MD_GridSpatialRepresentation_PropertyType", fieldName)
	}
	return
}

func (md_gridspatialrepresentation_type *MD_GridSpatialRepresentation_Type) GongMarshallField(stage *Stage, fieldName string) (res string) {
	ident := md_gridspatialrepresentation_type.GongGetIdentifier(stage)
	_ = ident
	switch fieldName {
	case "Name":
		res = __gong__marshallString(ident, "Name", md_gridspatialrepresentation_type.Name)
	case "Id":
		res = __gong__marshallString(ident, "Id", md_gridspatialrepresentation_type.Id)
	case "Uuid":
		res = __gong__marshallString(ident, "Uuid", md_gridspatialrepresentation_type.Uuid)

	case "AxisDimensionProperties":
		var sb strings.Builder
		for _, _md_dimension_propertytype := range md_gridspatialrepresentation_type.AxisDimensionProperties {
			sb.WriteString(__gong__marshallSliceOfPointers(ident, "AxisDimensionProperties", _md_dimension_propertytype.GongGetIdentifier(stage)))
		}
		res = sb.String()
	case "CellGeometry":
		if md_gridspatialrepresentation_type.CellGeometry != nil {
			res = __gong__marshallPointer(ident, "CellGeometry", md_gridspatialrepresentation_type.CellGeometry.GongGetIdentifier(stage))
		} else {
			res = __gong__marshallPointer(ident, "CellGeometry", "nil")
		}
	default:
		log.Panicf("Unknown field %s for Gongstruct MD_GridSpatialRepresentation_Type", fieldName)
	}
	return
}

func (md_pixelorientationcode_propertytype *MD_PixelOrientationCode_PropertyType) GongMarshallField(stage *Stage, fieldName string) (res string) {
	ident := md_pixelorientationcode_propertytype.GongGetIdentifier(stage)
	_ = ident
	switch fieldName {
	case "Name":
		res = __gong__marshallString(ident, "Name", md_pixelorientationcode_propertytype.Name)
	case "NilReason":
		res = __gong__marshallString(ident, "NilReason", md_pixelorientationcode_propertytype.NilReason)
	case "MD_PixelOrientationCode":
		res = __gong__marshallEnumString(ident, "MD_PixelOrientationCode", md_pixelorientationcode_propertytype.MD_PixelOrientationCode.ToCodeString())

	default:
		log.Panicf("Unknown field %s for Gongstruct MD_PixelOrientationCode_PropertyType", fieldName)
	}
	return
}

func (md_spatialrepresentation_propertytype *MD_SpatialRepresentation_PropertyType) GongMarshallField(stage *Stage, fieldName string) (res string) {
	ident := md_spatialrepresentation_propertytype.GongGetIdentifier(stage)
	_ = ident
	switch fieldName {
	case "Name":
		res = __gong__marshallString(ident, "Name", md_spatialrepresentation_propertytype.Name)
	case "NilReason":
		res = __gong__marshallString(ident, "NilReason", md_spatialrepresentation_propertytype.NilReason)
	case "Uuidref":
		res = __gong__marshallString(ident, "Uuidref", md_spatialrepresentation_propertytype.Uuidref)
	case "Type_XSD":
		res = __gong__marshallString(ident, "Type_XSD", md_spatialrepresentation_propertytype.Type_XSD)
	case "Href":
		res = __gong__marshallString(ident, "Href", md_spatialrepresentation_propertytype.Href)
	case "Role":
		res = __gong__marshallString(ident, "Role", md_spatialrepresentation_propertytype.Role)
	case "Arcrole":
		res = __gong__marshallString(ident, "Arcrole", md_spatialrepresentation_propertytype.Arcrole)
	case "Title":
		res = __gong__marshallString(ident, "Title", md_spatialrepresentation_propertytype.Title)
	case "Show":
		res = __gong__marshallString(ident, "Show", md_spatialrepresentation_propertytype.Show)
	case "Actuate":
		res = __gong__marshallString(ident, "Actuate", md_spatialrepresentation_propertytype.Actuate)

	case "MD_GridSpatialRepresentation":
		if md_spatialrepresentation_propertytype.MD_GridSpatialRepresentation != nil {
			res = __gong__marshallPointer(ident, "MD_GridSpatialRepresentation", md_spatialrepresentation_propertytype.MD_GridSpatialRepresentation.GongGetIdentifier(stage))
		} else {
			res = __gong__marshallPointer(ident, "MD_GridSpatialRepresentation", "nil")
		}
	case "MD_Georeferenceable":
		if md_spatialrepresentation_propertytype.MD_Georeferenceable != nil {
			res = __gong__marshallPointer(ident, "MD_Georeferenceable", md_spatialrepresentation_propertytype.MD_Georeferenceable.GongGetIdentifier(stage))
		} else {
			res = __gong__marshallPointer(ident, "MD_Georeferenceable", "nil")
		}
	case "MD_Georectified":
		if md_spatialrepresentation_propertytype.MD_Georectified != nil {
			res = __gong__marshallPointer(ident, "MD_Georectified", md_spatialrepresentation_propertytype.MD_Georectified.GongGetIdentifier(stage))
		} else {
			res = __gong__marshallPointer(ident, "MD_Georectified", "nil")
		}
	case "MD_VectorSpatialRepresentation":
		if md_spatialrepresentation_propertytype.MD_VectorSpatialRepresentation != nil {
			res = __gong__marshallPointer(ident, "MD_VectorSpatialRepresentation", md_spatialrepresentation_propertytype.MD_VectorSpatialRepresentation.GongGetIdentifier(stage))
		} else {
			res = __gong__marshallPointer(ident, "MD_VectorSpatialRepresentation", "nil")
		}
	default:
		log.Panicf("Unknown field %s for Gongstruct MD_SpatialRepresentation_PropertyType", fieldName)
	}
	return
}

func (md_topologylevelcode_propertytype *MD_TopologyLevelCode_PropertyType) GongMarshallField(stage *Stage, fieldName string) (res string) {
	ident := md_topologylevelcode_propertytype.GongGetIdentifier(stage)
	_ = ident
	switch fieldName {
	case "Name":
		res = __gong__marshallString(ident, "Name", md_topologylevelcode_propertytype.Name)
	case "NilReason":
		res = __gong__marshallString(ident, "NilReason", md_topologylevelcode_propertytype.NilReason)

	default:
		log.Panicf("Unknown field %s for Gongstruct MD_TopologyLevelCode_PropertyType", fieldName)
	}
	return
}

func (md_vectorspatialrepresentation_propertytype *MD_VectorSpatialRepresentation_PropertyType) GongMarshallField(stage *Stage, fieldName string) (res string) {
	ident := md_vectorspatialrepresentation_propertytype.GongGetIdentifier(stage)
	_ = ident
	switch fieldName {
	case "Name":
		res = __gong__marshallString(ident, "Name", md_vectorspatialrepresentation_propertytype.Name)
	case "NilReason":
		res = __gong__marshallString(ident, "NilReason", md_vectorspatialrepresentation_propertytype.NilReason)
	case "Uuidref":
		res = __gong__marshallString(ident, "Uuidref", md_vectorspatialrepresentation_propertytype.Uuidref)
	case "Type_XSD":
		res = __gong__marshallString(ident, "Type_XSD", md_vectorspatialrepresentation_propertytype.Type_XSD)
	case "Href":
		res = __gong__marshallString(ident, "Href", md_vectorspatialrepresentation_propertytype.Href)
	case "Role":
		res = __gong__marshallString(ident, "Role", md_vectorspatialrepresentation_propertytype.Role)
	case "Arcrole":
		res = __gong__marshallString(ident, "Arcrole", md_vectorspatialrepresentation_propertytype.Arcrole)
	case "Title":
		res = __gong__marshallString(ident, "Title", md_vectorspatialrepresentation_propertytype.Title)
	case "Show":
		res = __gong__marshallString(ident, "Show", md_vectorspatialrepresentation_propertytype.Show)
	case "Actuate":
		res = __gong__marshallString(ident, "Actuate", md_vectorspatialrepresentation_propertytype.Actuate)

	case "MD_VectorSpatialRepresentation":
		if md_vectorspatialrepresentation_propertytype.MD_VectorSpatialRepresentation != nil {
			res = __gong__marshallPointer(ident, "MD_VectorSpatialRepresentation", md_vectorspatialrepresentation_propertytype.MD_VectorSpatialRepresentation.GongGetIdentifier(stage))
		} else {
			res = __gong__marshallPointer(ident, "MD_VectorSpatialRepresentation", "nil")
		}
	default:
		log.Panicf("Unknown field %s for Gongstruct MD_VectorSpatialRepresentation_PropertyType", fieldName)
	}
	return
}

func (md_vectorspatialrepresentation_type *MD_VectorSpatialRepresentation_Type) GongMarshallField(stage *Stage, fieldName string) (res string) {
	ident := md_vectorspatialrepresentation_type.GongGetIdentifier(stage)
	_ = ident
	switch fieldName {
	case "Name":
		res = __gong__marshallString(ident, "Name", md_vectorspatialrepresentation_type.Name)
	case "Id":
		res = __gong__marshallString(ident, "Id", md_vectorspatialrepresentation_type.Id)
	case "Uuid":
		res = __gong__marshallString(ident, "Uuid", md_vectorspatialrepresentation_type.Uuid)

	case "TopologyLevel":
		if md_vectorspatialrepresentation_type.TopologyLevel != nil {
			res = __gong__marshallPointer(ident, "TopologyLevel", md_vectorspatialrepresentation_type.TopologyLevel.GongGetIdentifier(stage))
		} else {
			res = __gong__marshallPointer(ident, "TopologyLevel", "nil")
		}
	case "GeometricObjects":
		var sb strings.Builder
		for _, _md_geometricobjects_propertytype := range md_vectorspatialrepresentation_type.GeometricObjects {
			sb.WriteString(__gong__marshallSliceOfPointers(ident, "GeometricObjects", _md_geometricobjects_propertytype.GongGetIdentifier(stage)))
		}
		res = sb.String()
	default:
		log.Panicf("Unknown field %s for Gongstruct MD_VectorSpatialRepresentation_Type", fieldName)
	}
	return
}

// insertion point for marshall all fields methods
func (abstractmd_spatialrepresentation_type *AbstractMD_SpatialRepresentation_Type) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(abstractmd_spatialrepresentation_type.GongMarshallField(stage, "Name"))
		initializerStatements.WriteString(abstractmd_spatialrepresentation_type.GongMarshallField(stage, "Id"))
		initializerStatements.WriteString(abstractmd_spatialrepresentation_type.GongMarshallField(stage, "Uuid"))
	}
	initRes = initializerStatements.String()
	ptrRes = pointersInitializesStatements.String()
	return
}
func (md_cellgeometrycode_propertytype *MD_CellGeometryCode_PropertyType) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(md_cellgeometrycode_propertytype.GongMarshallField(stage, "Name"))
		initializerStatements.WriteString(md_cellgeometrycode_propertytype.GongMarshallField(stage, "NilReason"))
	}
	initRes = initializerStatements.String()
	ptrRes = pointersInitializesStatements.String()
	return
}
func (md_dimensionnametypecode_propertytype *MD_DimensionNameTypeCode_PropertyType) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(md_dimensionnametypecode_propertytype.GongMarshallField(stage, "Name"))
		initializerStatements.WriteString(md_dimensionnametypecode_propertytype.GongMarshallField(stage, "NilReason"))
	}
	initRes = initializerStatements.String()
	ptrRes = pointersInitializesStatements.String()
	return
}
func (md_dimension_propertytype *MD_Dimension_PropertyType) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(md_dimension_propertytype.GongMarshallField(stage, "Name"))
		initializerStatements.WriteString(md_dimension_propertytype.GongMarshallField(stage, "NilReason"))
		initializerStatements.WriteString(md_dimension_propertytype.GongMarshallField(stage, "Uuidref"))
		initializerStatements.WriteString(md_dimension_propertytype.GongMarshallField(stage, "Type_XSD"))
		initializerStatements.WriteString(md_dimension_propertytype.GongMarshallField(stage, "Href"))
		initializerStatements.WriteString(md_dimension_propertytype.GongMarshallField(stage, "Role"))
		initializerStatements.WriteString(md_dimension_propertytype.GongMarshallField(stage, "Arcrole"))
		initializerStatements.WriteString(md_dimension_propertytype.GongMarshallField(stage, "Title"))
		initializerStatements.WriteString(md_dimension_propertytype.GongMarshallField(stage, "Show"))
		initializerStatements.WriteString(md_dimension_propertytype.GongMarshallField(stage, "Actuate"))
		pointersInitializesStatements.WriteString(md_dimension_propertytype.GongMarshallField(stage, "MD_Dimension"))
	}
	initRes = initializerStatements.String()
	ptrRes = pointersInitializesStatements.String()
	return
}
func (md_dimension_type *MD_Dimension_Type) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(md_dimension_type.GongMarshallField(stage, "Name"))
		initializerStatements.WriteString(md_dimension_type.GongMarshallField(stage, "Id"))
		initializerStatements.WriteString(md_dimension_type.GongMarshallField(stage, "Uuid"))
		pointersInitializesStatements.WriteString(md_dimension_type.GongMarshallField(stage, "DimensionName"))
	}
	initRes = initializerStatements.String()
	ptrRes = pointersInitializesStatements.String()
	return
}
func (md_geometricobjecttypecode_propertytype *MD_GeometricObjectTypeCode_PropertyType) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(md_geometricobjecttypecode_propertytype.GongMarshallField(stage, "Name"))
		initializerStatements.WriteString(md_geometricobjecttypecode_propertytype.GongMarshallField(stage, "NilReason"))
	}
	initRes = initializerStatements.String()
	ptrRes = pointersInitializesStatements.String()
	return
}
func (md_geometricobjects_propertytype *MD_GeometricObjects_PropertyType) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(md_geometricobjects_propertytype.GongMarshallField(stage, "Name"))
		initializerStatements.WriteString(md_geometricobjects_propertytype.GongMarshallField(stage, "NilReason"))
		initializerStatements.WriteString(md_geometricobjects_propertytype.GongMarshallField(stage, "Uuidref"))
		initializerStatements.WriteString(md_geometricobjects_propertytype.GongMarshallField(stage, "Type_XSD"))
		initializerStatements.WriteString(md_geometricobjects_propertytype.GongMarshallField(stage, "Href"))
		initializerStatements.WriteString(md_geometricobjects_propertytype.GongMarshallField(stage, "Role"))
		initializerStatements.WriteString(md_geometricobjects_propertytype.GongMarshallField(stage, "Arcrole"))
		initializerStatements.WriteString(md_geometricobjects_propertytype.GongMarshallField(stage, "Title"))
		initializerStatements.WriteString(md_geometricobjects_propertytype.GongMarshallField(stage, "Show"))
		initializerStatements.WriteString(md_geometricobjects_propertytype.GongMarshallField(stage, "Actuate"))
		pointersInitializesStatements.WriteString(md_geometricobjects_propertytype.GongMarshallField(stage, "MD_GeometricObjects"))
	}
	initRes = initializerStatements.String()
	ptrRes = pointersInitializesStatements.String()
	return
}
func (md_geometricobjects_type *MD_GeometricObjects_Type) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(md_geometricobjects_type.GongMarshallField(stage, "Name"))
		initializerStatements.WriteString(md_geometricobjects_type.GongMarshallField(stage, "Id"))
		initializerStatements.WriteString(md_geometricobjects_type.GongMarshallField(stage, "Uuid"))
		pointersInitializesStatements.WriteString(md_geometricobjects_type.GongMarshallField(stage, "GeometricObjectType"))
	}
	initRes = initializerStatements.String()
	ptrRes = pointersInitializesStatements.String()
	return
}
func (md_georectified_propertytype *MD_Georectified_PropertyType) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(md_georectified_propertytype.GongMarshallField(stage, "Name"))
		initializerStatements.WriteString(md_georectified_propertytype.GongMarshallField(stage, "NilReason"))
		initializerStatements.WriteString(md_georectified_propertytype.GongMarshallField(stage, "Uuidref"))
		initializerStatements.WriteString(md_georectified_propertytype.GongMarshallField(stage, "Type_XSD"))
		initializerStatements.WriteString(md_georectified_propertytype.GongMarshallField(stage, "Href"))
		initializerStatements.WriteString(md_georectified_propertytype.GongMarshallField(stage, "Role"))
		initializerStatements.WriteString(md_georectified_propertytype.GongMarshallField(stage, "Arcrole"))
		initializerStatements.WriteString(md_georectified_propertytype.GongMarshallField(stage, "Title"))
		initializerStatements.WriteString(md_georectified_propertytype.GongMarshallField(stage, "Show"))
		initializerStatements.WriteString(md_georectified_propertytype.GongMarshallField(stage, "Actuate"))
		pointersInitializesStatements.WriteString(md_georectified_propertytype.GongMarshallField(stage, "MD_Georectified"))
	}
	initRes = initializerStatements.String()
	ptrRes = pointersInitializesStatements.String()
	return
}
func (md_georectified_type *MD_Georectified_Type) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(md_georectified_type.GongMarshallField(stage, "Name"))
		initializerStatements.WriteString(md_georectified_type.GongMarshallField(stage, "Id"))
		initializerStatements.WriteString(md_georectified_type.GongMarshallField(stage, "Uuid"))
		pointersInitializesStatements.WriteString(md_georectified_type.GongMarshallField(stage, "AxisDimensionProperties"))
		pointersInitializesStatements.WriteString(md_georectified_type.GongMarshallField(stage, "CellGeometry"))
		pointersInitializesStatements.WriteString(md_georectified_type.GongMarshallField(stage, "PointInPixel"))
	}
	initRes = initializerStatements.String()
	ptrRes = pointersInitializesStatements.String()
	return
}
func (md_georeferenceable_propertytype *MD_Georeferenceable_PropertyType) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(md_georeferenceable_propertytype.GongMarshallField(stage, "Name"))
		initializerStatements.WriteString(md_georeferenceable_propertytype.GongMarshallField(stage, "NilReason"))
		initializerStatements.WriteString(md_georeferenceable_propertytype.GongMarshallField(stage, "Uuidref"))
		initializerStatements.WriteString(md_georeferenceable_propertytype.GongMarshallField(stage, "Type_XSD"))
		initializerStatements.WriteString(md_georeferenceable_propertytype.GongMarshallField(stage, "Href"))
		initializerStatements.WriteString(md_georeferenceable_propertytype.GongMarshallField(stage, "Role"))
		initializerStatements.WriteString(md_georeferenceable_propertytype.GongMarshallField(stage, "Arcrole"))
		initializerStatements.WriteString(md_georeferenceable_propertytype.GongMarshallField(stage, "Title"))
		initializerStatements.WriteString(md_georeferenceable_propertytype.GongMarshallField(stage, "Show"))
		initializerStatements.WriteString(md_georeferenceable_propertytype.GongMarshallField(stage, "Actuate"))
		pointersInitializesStatements.WriteString(md_georeferenceable_propertytype.GongMarshallField(stage, "MD_Georeferenceable"))
	}
	initRes = initializerStatements.String()
	ptrRes = pointersInitializesStatements.String()
	return
}
func (md_georeferenceable_type *MD_Georeferenceable_Type) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(md_georeferenceable_type.GongMarshallField(stage, "Name"))
		initializerStatements.WriteString(md_georeferenceable_type.GongMarshallField(stage, "Id"))
		initializerStatements.WriteString(md_georeferenceable_type.GongMarshallField(stage, "Uuid"))
		pointersInitializesStatements.WriteString(md_georeferenceable_type.GongMarshallField(stage, "AxisDimensionProperties"))
		pointersInitializesStatements.WriteString(md_georeferenceable_type.GongMarshallField(stage, "CellGeometry"))
	}
	initRes = initializerStatements.String()
	ptrRes = pointersInitializesStatements.String()
	return
}
func (md_gridspatialrepresentation_propertytype *MD_GridSpatialRepresentation_PropertyType) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(md_gridspatialrepresentation_propertytype.GongMarshallField(stage, "Name"))
		initializerStatements.WriteString(md_gridspatialrepresentation_propertytype.GongMarshallField(stage, "NilReason"))
		initializerStatements.WriteString(md_gridspatialrepresentation_propertytype.GongMarshallField(stage, "Uuidref"))
		initializerStatements.WriteString(md_gridspatialrepresentation_propertytype.GongMarshallField(stage, "Type_XSD"))
		initializerStatements.WriteString(md_gridspatialrepresentation_propertytype.GongMarshallField(stage, "Href"))
		initializerStatements.WriteString(md_gridspatialrepresentation_propertytype.GongMarshallField(stage, "Role"))
		initializerStatements.WriteString(md_gridspatialrepresentation_propertytype.GongMarshallField(stage, "Arcrole"))
		initializerStatements.WriteString(md_gridspatialrepresentation_propertytype.GongMarshallField(stage, "Title"))
		initializerStatements.WriteString(md_gridspatialrepresentation_propertytype.GongMarshallField(stage, "Show"))
		initializerStatements.WriteString(md_gridspatialrepresentation_propertytype.GongMarshallField(stage, "Actuate"))
		pointersInitializesStatements.WriteString(md_gridspatialrepresentation_propertytype.GongMarshallField(stage, "MD_GridSpatialRepresentation"))
		pointersInitializesStatements.WriteString(md_gridspatialrepresentation_propertytype.GongMarshallField(stage, "MD_Georeferenceable"))
		pointersInitializesStatements.WriteString(md_gridspatialrepresentation_propertytype.GongMarshallField(stage, "MD_Georectified"))
	}
	initRes = initializerStatements.String()
	ptrRes = pointersInitializesStatements.String()
	return
}
func (md_gridspatialrepresentation_type *MD_GridSpatialRepresentation_Type) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(md_gridspatialrepresentation_type.GongMarshallField(stage, "Name"))
		initializerStatements.WriteString(md_gridspatialrepresentation_type.GongMarshallField(stage, "Id"))
		initializerStatements.WriteString(md_gridspatialrepresentation_type.GongMarshallField(stage, "Uuid"))
		pointersInitializesStatements.WriteString(md_gridspatialrepresentation_type.GongMarshallField(stage, "AxisDimensionProperties"))
		pointersInitializesStatements.WriteString(md_gridspatialrepresentation_type.GongMarshallField(stage, "CellGeometry"))
	}
	initRes = initializerStatements.String()
	ptrRes = pointersInitializesStatements.String()
	return
}
func (md_pixelorientationcode_propertytype *MD_PixelOrientationCode_PropertyType) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(md_pixelorientationcode_propertytype.GongMarshallField(stage, "Name"))
		initializerStatements.WriteString(md_pixelorientationcode_propertytype.GongMarshallField(stage, "NilReason"))
		initializerStatements.WriteString(md_pixelorientationcode_propertytype.GongMarshallField(stage, "MD_PixelOrientationCode"))
	}
	initRes = initializerStatements.String()
	ptrRes = pointersInitializesStatements.String()
	return
}
func (md_spatialrepresentation_propertytype *MD_SpatialRepresentation_PropertyType) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(md_spatialrepresentation_propertytype.GongMarshallField(stage, "Name"))
		initializerStatements.WriteString(md_spatialrepresentation_propertytype.GongMarshallField(stage, "NilReason"))
		initializerStatements.WriteString(md_spatialrepresentation_propertytype.GongMarshallField(stage, "Uuidref"))
		initializerStatements.WriteString(md_spatialrepresentation_propertytype.GongMarshallField(stage, "Type_XSD"))
		initializerStatements.WriteString(md_spatialrepresentation_propertytype.GongMarshallField(stage, "Href"))
		initializerStatements.WriteString(md_spatialrepresentation_propertytype.GongMarshallField(stage, "Role"))
		initializerStatements.WriteString(md_spatialrepresentation_propertytype.GongMarshallField(stage, "Arcrole"))
		initializerStatements.WriteString(md_spatialrepresentation_propertytype.GongMarshallField(stage, "Title"))
		initializerStatements.WriteString(md_spatialrepresentation_propertytype.GongMarshallField(stage, "Show"))
		initializerStatements.WriteString(md_spatialrepresentation_propertytype.GongMarshallField(stage, "Actuate"))
		pointersInitializesStatements.WriteString(md_spatialrepresentation_propertytype.GongMarshallField(stage, "MD_GridSpatialRepresentation"))
		pointersInitializesStatements.WriteString(md_spatialrepresentation_propertytype.GongMarshallField(stage, "MD_Georeferenceable"))
		pointersInitializesStatements.WriteString(md_spatialrepresentation_propertytype.GongMarshallField(stage, "MD_Georectified"))
		pointersInitializesStatements.WriteString(md_spatialrepresentation_propertytype.GongMarshallField(stage, "MD_VectorSpatialRepresentation"))
	}
	initRes = initializerStatements.String()
	ptrRes = pointersInitializesStatements.String()
	return
}
func (md_topologylevelcode_propertytype *MD_TopologyLevelCode_PropertyType) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(md_topologylevelcode_propertytype.GongMarshallField(stage, "Name"))
		initializerStatements.WriteString(md_topologylevelcode_propertytype.GongMarshallField(stage, "NilReason"))
	}
	initRes = initializerStatements.String()
	ptrRes = pointersInitializesStatements.String()
	return
}
func (md_vectorspatialrepresentation_propertytype *MD_VectorSpatialRepresentation_PropertyType) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(md_vectorspatialrepresentation_propertytype.GongMarshallField(stage, "Name"))
		initializerStatements.WriteString(md_vectorspatialrepresentation_propertytype.GongMarshallField(stage, "NilReason"))
		initializerStatements.WriteString(md_vectorspatialrepresentation_propertytype.GongMarshallField(stage, "Uuidref"))
		initializerStatements.WriteString(md_vectorspatialrepresentation_propertytype.GongMarshallField(stage, "Type_XSD"))
		initializerStatements.WriteString(md_vectorspatialrepresentation_propertytype.GongMarshallField(stage, "Href"))
		initializerStatements.WriteString(md_vectorspatialrepresentation_propertytype.GongMarshallField(stage, "Role"))
		initializerStatements.WriteString(md_vectorspatialrepresentation_propertytype.GongMarshallField(stage, "Arcrole"))
		initializerStatements.WriteString(md_vectorspatialrepresentation_propertytype.GongMarshallField(stage, "Title"))
		initializerStatements.WriteString(md_vectorspatialrepresentation_propertytype.GongMarshallField(stage, "Show"))
		initializerStatements.WriteString(md_vectorspatialrepresentation_propertytype.GongMarshallField(stage, "Actuate"))
		pointersInitializesStatements.WriteString(md_vectorspatialrepresentation_propertytype.GongMarshallField(stage, "MD_VectorSpatialRepresentation"))
	}
	initRes = initializerStatements.String()
	ptrRes = pointersInitializesStatements.String()
	return
}
func (md_vectorspatialrepresentation_type *MD_VectorSpatialRepresentation_Type) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(md_vectorspatialrepresentation_type.GongMarshallField(stage, "Name"))
		initializerStatements.WriteString(md_vectorspatialrepresentation_type.GongMarshallField(stage, "Id"))
		initializerStatements.WriteString(md_vectorspatialrepresentation_type.GongMarshallField(stage, "Uuid"))
		pointersInitializesStatements.WriteString(md_vectorspatialrepresentation_type.GongMarshallField(stage, "TopologyLevel"))
		pointersInitializesStatements.WriteString(md_vectorspatialrepresentation_type.GongMarshallField(stage, "GeometricObjects"))
	}
	initRes = initializerStatements.String()
	ptrRes = pointersInitializesStatements.String()
	return
}
