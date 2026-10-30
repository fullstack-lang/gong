// generated code - do not edit
package gmd

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
	{{Identifier}} := (&gmd.{{GeneratedStructName}}{Name: {{GeneratedFieldNameValue}}}).Stage(stage)`

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
			const fnSig = "func _(stage *gmd.Stage) {"
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
			strings.HasSuffix(trimmedContentBeforeBrace, "func _(stage *gmd.Stage) {") ||
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
	modelsPackageName := "github.com/fullstack-lang/gong/app/xsd/tests/aixm/go/models/gmd"
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
		for _, abstractds_aggregate_type := range __gong__sortStageSetInstances(stage.AbstractDS_Aggregate_Types, stage.AbstractDS_Aggregate_Type_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			abstractds_aggregate_typeIdent := abstractds_aggregate_type.GongGetIdentifier(stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&gmd.AbstractDS_Aggregate_Type{Name: %s}).Stage(stage)", abstractds_aggregate_typeIdent, __gong__toRawStringLiteral(abstractds_aggregate_type.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", abstractds_aggregate_typeIdent, __gong__toRawStringLiteral(abstractds_aggregate_type.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Id = %s", abstractds_aggregate_typeIdent, __gong__toRawStringLiteral(abstractds_aggregate_type.Id)))
			values.WriteString(fmt.Sprintf("\n\t%s.Uuid = %s", abstractds_aggregate_typeIdent, __gong__toRawStringLiteral(abstractds_aggregate_type.Uuid)))
			for _, elem := range abstractds_aggregate_type.ComposedOf {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := elem.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.ComposedOf = append(%s.ComposedOf, %s)", abstractds_aggregate_typeIdent, abstractds_aggregate_typeIdent, targetIdent))
			}
			for _, elem := range abstractds_aggregate_type.SeriesMetadata {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := elem.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.SeriesMetadata = append(%s.SeriesMetadata, %s)", abstractds_aggregate_typeIdent, abstractds_aggregate_typeIdent, targetIdent))
			}
			for _, elem := range abstractds_aggregate_type.Subset {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := elem.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Subset = append(%s.Subset, %s)", abstractds_aggregate_typeIdent, abstractds_aggregate_typeIdent, targetIdent))
			}
			for _, elem := range abstractds_aggregate_type.Superset {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := elem.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Superset = append(%s.Superset, %s)", abstractds_aggregate_typeIdent, abstractds_aggregate_typeIdent, targetIdent))
			}
		}
	}
	if stage != nil {
		for _, ds_aggregate_propertytype := range __gong__sortStageSetInstances(stage.DS_Aggregate_PropertyTypes, stage.DS_Aggregate_PropertyType_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			ds_aggregate_propertytypeIdent := ds_aggregate_propertytype.GongGetIdentifier(stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&gmd.DS_Aggregate_PropertyType{Name: %s}).Stage(stage)", ds_aggregate_propertytypeIdent, __gong__toRawStringLiteral(ds_aggregate_propertytype.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", ds_aggregate_propertytypeIdent, __gong__toRawStringLiteral(ds_aggregate_propertytype.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.NilReason = %s", ds_aggregate_propertytypeIdent, __gong__toRawStringLiteral(ds_aggregate_propertytype.NilReason)))
			values.WriteString(fmt.Sprintf("\n\t%s.Uuidref = %s", ds_aggregate_propertytypeIdent, __gong__toRawStringLiteral(ds_aggregate_propertytype.Uuidref)))
			values.WriteString(fmt.Sprintf("\n\t%s.Type_XSD = %s", ds_aggregate_propertytypeIdent, __gong__toRawStringLiteral(ds_aggregate_propertytype.Type_XSD)))
			values.WriteString(fmt.Sprintf("\n\t%s.Href = %s", ds_aggregate_propertytypeIdent, __gong__toRawStringLiteral(ds_aggregate_propertytype.Href)))
			values.WriteString(fmt.Sprintf("\n\t%s.Role = %s", ds_aggregate_propertytypeIdent, __gong__toRawStringLiteral(ds_aggregate_propertytype.Role)))
			values.WriteString(fmt.Sprintf("\n\t%s.Arcrole = %s", ds_aggregate_propertytypeIdent, __gong__toRawStringLiteral(ds_aggregate_propertytype.Arcrole)))
			values.WriteString(fmt.Sprintf("\n\t%s.Title = %s", ds_aggregate_propertytypeIdent, __gong__toRawStringLiteral(ds_aggregate_propertytype.Title)))
			values.WriteString(fmt.Sprintf("\n\t%s.Show = %s", ds_aggregate_propertytypeIdent, __gong__toRawStringLiteral(ds_aggregate_propertytype.Show)))
			values.WriteString(fmt.Sprintf("\n\t%s.Actuate = %s", ds_aggregate_propertytypeIdent, __gong__toRawStringLiteral(ds_aggregate_propertytype.Actuate)))
			if ds_aggregate_propertytype.DS_OtherAggregate != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := ds_aggregate_propertytype.DS_OtherAggregate.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.DS_OtherAggregate = %s", ds_aggregate_propertytypeIdent, targetIdent))
			}
			if ds_aggregate_propertytype.DS_StereoMate != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := ds_aggregate_propertytype.DS_StereoMate.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.DS_StereoMate = %s", ds_aggregate_propertytypeIdent, targetIdent))
			}
			if ds_aggregate_propertytype.DS_Series != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := ds_aggregate_propertytype.DS_Series.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.DS_Series = %s", ds_aggregate_propertytypeIdent, targetIdent))
			}
			if ds_aggregate_propertytype.DS_Platform != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := ds_aggregate_propertytype.DS_Platform.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.DS_Platform = %s", ds_aggregate_propertytypeIdent, targetIdent))
			}
			if ds_aggregate_propertytype.DS_Sensor != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := ds_aggregate_propertytype.DS_Sensor.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.DS_Sensor = %s", ds_aggregate_propertytypeIdent, targetIdent))
			}
			if ds_aggregate_propertytype.DS_ProductionSeries != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := ds_aggregate_propertytype.DS_ProductionSeries.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.DS_ProductionSeries = %s", ds_aggregate_propertytypeIdent, targetIdent))
			}
			if ds_aggregate_propertytype.DS_Initiative != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := ds_aggregate_propertytype.DS_Initiative.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.DS_Initiative = %s", ds_aggregate_propertytypeIdent, targetIdent))
			}
		}
	}
	if stage != nil {
		for _, ds_dataset_propertytype := range __gong__sortStageSetInstances(stage.DS_DataSet_PropertyTypes, stage.DS_DataSet_PropertyType_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			ds_dataset_propertytypeIdent := ds_dataset_propertytype.GongGetIdentifier(stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&gmd.DS_DataSet_PropertyType{Name: %s}).Stage(stage)", ds_dataset_propertytypeIdent, __gong__toRawStringLiteral(ds_dataset_propertytype.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", ds_dataset_propertytypeIdent, __gong__toRawStringLiteral(ds_dataset_propertytype.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.NilReason = %s", ds_dataset_propertytypeIdent, __gong__toRawStringLiteral(ds_dataset_propertytype.NilReason)))
			values.WriteString(fmt.Sprintf("\n\t%s.Uuidref = %s", ds_dataset_propertytypeIdent, __gong__toRawStringLiteral(ds_dataset_propertytype.Uuidref)))
			values.WriteString(fmt.Sprintf("\n\t%s.Type_XSD = %s", ds_dataset_propertytypeIdent, __gong__toRawStringLiteral(ds_dataset_propertytype.Type_XSD)))
			values.WriteString(fmt.Sprintf("\n\t%s.Href = %s", ds_dataset_propertytypeIdent, __gong__toRawStringLiteral(ds_dataset_propertytype.Href)))
			values.WriteString(fmt.Sprintf("\n\t%s.Role = %s", ds_dataset_propertytypeIdent, __gong__toRawStringLiteral(ds_dataset_propertytype.Role)))
			values.WriteString(fmt.Sprintf("\n\t%s.Arcrole = %s", ds_dataset_propertytypeIdent, __gong__toRawStringLiteral(ds_dataset_propertytype.Arcrole)))
			values.WriteString(fmt.Sprintf("\n\t%s.Title = %s", ds_dataset_propertytypeIdent, __gong__toRawStringLiteral(ds_dataset_propertytype.Title)))
			values.WriteString(fmt.Sprintf("\n\t%s.Show = %s", ds_dataset_propertytypeIdent, __gong__toRawStringLiteral(ds_dataset_propertytype.Show)))
			values.WriteString(fmt.Sprintf("\n\t%s.Actuate = %s", ds_dataset_propertytypeIdent, __gong__toRawStringLiteral(ds_dataset_propertytype.Actuate)))
			if ds_dataset_propertytype.DS_DataSet != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := ds_dataset_propertytype.DS_DataSet.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.DS_DataSet = %s", ds_dataset_propertytypeIdent, targetIdent))
			}
		}
	}
	if stage != nil {
		for _, ds_dataset_type := range __gong__sortStageSetInstances(stage.DS_DataSet_Types, stage.DS_DataSet_Type_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			ds_dataset_typeIdent := ds_dataset_type.GongGetIdentifier(stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&gmd.DS_DataSet_Type{Name: %s}).Stage(stage)", ds_dataset_typeIdent, __gong__toRawStringLiteral(ds_dataset_type.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", ds_dataset_typeIdent, __gong__toRawStringLiteral(ds_dataset_type.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Id = %s", ds_dataset_typeIdent, __gong__toRawStringLiteral(ds_dataset_type.Id)))
			values.WriteString(fmt.Sprintf("\n\t%s.Uuid = %s", ds_dataset_typeIdent, __gong__toRawStringLiteral(ds_dataset_type.Uuid)))
			for _, elem := range ds_dataset_type.Has {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := elem.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Has = append(%s.Has, %s)", ds_dataset_typeIdent, ds_dataset_typeIdent, targetIdent))
			}
			for _, elem := range ds_dataset_type.PartOf {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := elem.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.PartOf = append(%s.PartOf, %s)", ds_dataset_typeIdent, ds_dataset_typeIdent, targetIdent))
			}
		}
	}
	if stage != nil {
		for _, ds_initiative_propertytype := range __gong__sortStageSetInstances(stage.DS_Initiative_PropertyTypes, stage.DS_Initiative_PropertyType_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			ds_initiative_propertytypeIdent := ds_initiative_propertytype.GongGetIdentifier(stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&gmd.DS_Initiative_PropertyType{Name: %s}).Stage(stage)", ds_initiative_propertytypeIdent, __gong__toRawStringLiteral(ds_initiative_propertytype.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", ds_initiative_propertytypeIdent, __gong__toRawStringLiteral(ds_initiative_propertytype.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.NilReason = %s", ds_initiative_propertytypeIdent, __gong__toRawStringLiteral(ds_initiative_propertytype.NilReason)))
			values.WriteString(fmt.Sprintf("\n\t%s.Uuidref = %s", ds_initiative_propertytypeIdent, __gong__toRawStringLiteral(ds_initiative_propertytype.Uuidref)))
			values.WriteString(fmt.Sprintf("\n\t%s.Type_XSD = %s", ds_initiative_propertytypeIdent, __gong__toRawStringLiteral(ds_initiative_propertytype.Type_XSD)))
			values.WriteString(fmt.Sprintf("\n\t%s.Href = %s", ds_initiative_propertytypeIdent, __gong__toRawStringLiteral(ds_initiative_propertytype.Href)))
			values.WriteString(fmt.Sprintf("\n\t%s.Role = %s", ds_initiative_propertytypeIdent, __gong__toRawStringLiteral(ds_initiative_propertytype.Role)))
			values.WriteString(fmt.Sprintf("\n\t%s.Arcrole = %s", ds_initiative_propertytypeIdent, __gong__toRawStringLiteral(ds_initiative_propertytype.Arcrole)))
			values.WriteString(fmt.Sprintf("\n\t%s.Title = %s", ds_initiative_propertytypeIdent, __gong__toRawStringLiteral(ds_initiative_propertytype.Title)))
			values.WriteString(fmt.Sprintf("\n\t%s.Show = %s", ds_initiative_propertytypeIdent, __gong__toRawStringLiteral(ds_initiative_propertytype.Show)))
			values.WriteString(fmt.Sprintf("\n\t%s.Actuate = %s", ds_initiative_propertytypeIdent, __gong__toRawStringLiteral(ds_initiative_propertytype.Actuate)))
			if ds_initiative_propertytype.DS_Initiative != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := ds_initiative_propertytype.DS_Initiative.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.DS_Initiative = %s", ds_initiative_propertytypeIdent, targetIdent))
			}
		}
	}
	if stage != nil {
		for _, ds_initiative_type := range __gong__sortStageSetInstances(stage.DS_Initiative_Types, stage.DS_Initiative_Type_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			ds_initiative_typeIdent := ds_initiative_type.GongGetIdentifier(stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&gmd.DS_Initiative_Type{Name: %s}).Stage(stage)", ds_initiative_typeIdent, __gong__toRawStringLiteral(ds_initiative_type.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", ds_initiative_typeIdent, __gong__toRawStringLiteral(ds_initiative_type.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Id = %s", ds_initiative_typeIdent, __gong__toRawStringLiteral(ds_initiative_type.Id)))
			values.WriteString(fmt.Sprintf("\n\t%s.Uuid = %s", ds_initiative_typeIdent, __gong__toRawStringLiteral(ds_initiative_type.Uuid)))
			for _, elem := range ds_initiative_type.ComposedOf {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := elem.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.ComposedOf = append(%s.ComposedOf, %s)", ds_initiative_typeIdent, ds_initiative_typeIdent, targetIdent))
			}
			for _, elem := range ds_initiative_type.SeriesMetadata {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := elem.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.SeriesMetadata = append(%s.SeriesMetadata, %s)", ds_initiative_typeIdent, ds_initiative_typeIdent, targetIdent))
			}
			for _, elem := range ds_initiative_type.Subset {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := elem.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Subset = append(%s.Subset, %s)", ds_initiative_typeIdent, ds_initiative_typeIdent, targetIdent))
			}
			for _, elem := range ds_initiative_type.Superset {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := elem.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Superset = append(%s.Superset, %s)", ds_initiative_typeIdent, ds_initiative_typeIdent, targetIdent))
			}
		}
	}
	if stage != nil {
		for _, ds_otheraggregate_propertytype := range __gong__sortStageSetInstances(stage.DS_OtherAggregate_PropertyTypes, stage.DS_OtherAggregate_PropertyType_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			ds_otheraggregate_propertytypeIdent := ds_otheraggregate_propertytype.GongGetIdentifier(stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&gmd.DS_OtherAggregate_PropertyType{Name: %s}).Stage(stage)", ds_otheraggregate_propertytypeIdent, __gong__toRawStringLiteral(ds_otheraggregate_propertytype.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", ds_otheraggregate_propertytypeIdent, __gong__toRawStringLiteral(ds_otheraggregate_propertytype.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.NilReason = %s", ds_otheraggregate_propertytypeIdent, __gong__toRawStringLiteral(ds_otheraggregate_propertytype.NilReason)))
			values.WriteString(fmt.Sprintf("\n\t%s.Uuidref = %s", ds_otheraggregate_propertytypeIdent, __gong__toRawStringLiteral(ds_otheraggregate_propertytype.Uuidref)))
			values.WriteString(fmt.Sprintf("\n\t%s.Type_XSD = %s", ds_otheraggregate_propertytypeIdent, __gong__toRawStringLiteral(ds_otheraggregate_propertytype.Type_XSD)))
			values.WriteString(fmt.Sprintf("\n\t%s.Href = %s", ds_otheraggregate_propertytypeIdent, __gong__toRawStringLiteral(ds_otheraggregate_propertytype.Href)))
			values.WriteString(fmt.Sprintf("\n\t%s.Role = %s", ds_otheraggregate_propertytypeIdent, __gong__toRawStringLiteral(ds_otheraggregate_propertytype.Role)))
			values.WriteString(fmt.Sprintf("\n\t%s.Arcrole = %s", ds_otheraggregate_propertytypeIdent, __gong__toRawStringLiteral(ds_otheraggregate_propertytype.Arcrole)))
			values.WriteString(fmt.Sprintf("\n\t%s.Title = %s", ds_otheraggregate_propertytypeIdent, __gong__toRawStringLiteral(ds_otheraggregate_propertytype.Title)))
			values.WriteString(fmt.Sprintf("\n\t%s.Show = %s", ds_otheraggregate_propertytypeIdent, __gong__toRawStringLiteral(ds_otheraggregate_propertytype.Show)))
			values.WriteString(fmt.Sprintf("\n\t%s.Actuate = %s", ds_otheraggregate_propertytypeIdent, __gong__toRawStringLiteral(ds_otheraggregate_propertytype.Actuate)))
			if ds_otheraggregate_propertytype.DS_OtherAggregate != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := ds_otheraggregate_propertytype.DS_OtherAggregate.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.DS_OtherAggregate = %s", ds_otheraggregate_propertytypeIdent, targetIdent))
			}
			if ds_otheraggregate_propertytype.DS_StereoMate != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := ds_otheraggregate_propertytype.DS_StereoMate.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.DS_StereoMate = %s", ds_otheraggregate_propertytypeIdent, targetIdent))
			}
		}
	}
	if stage != nil {
		for _, ds_otheraggregate_type := range __gong__sortStageSetInstances(stage.DS_OtherAggregate_Types, stage.DS_OtherAggregate_Type_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			ds_otheraggregate_typeIdent := ds_otheraggregate_type.GongGetIdentifier(stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&gmd.DS_OtherAggregate_Type{Name: %s}).Stage(stage)", ds_otheraggregate_typeIdent, __gong__toRawStringLiteral(ds_otheraggregate_type.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", ds_otheraggregate_typeIdent, __gong__toRawStringLiteral(ds_otheraggregate_type.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Id = %s", ds_otheraggregate_typeIdent, __gong__toRawStringLiteral(ds_otheraggregate_type.Id)))
			values.WriteString(fmt.Sprintf("\n\t%s.Uuid = %s", ds_otheraggregate_typeIdent, __gong__toRawStringLiteral(ds_otheraggregate_type.Uuid)))
			for _, elem := range ds_otheraggregate_type.ComposedOf {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := elem.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.ComposedOf = append(%s.ComposedOf, %s)", ds_otheraggregate_typeIdent, ds_otheraggregate_typeIdent, targetIdent))
			}
			for _, elem := range ds_otheraggregate_type.SeriesMetadata {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := elem.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.SeriesMetadata = append(%s.SeriesMetadata, %s)", ds_otheraggregate_typeIdent, ds_otheraggregate_typeIdent, targetIdent))
			}
			for _, elem := range ds_otheraggregate_type.Subset {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := elem.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Subset = append(%s.Subset, %s)", ds_otheraggregate_typeIdent, ds_otheraggregate_typeIdent, targetIdent))
			}
			for _, elem := range ds_otheraggregate_type.Superset {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := elem.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Superset = append(%s.Superset, %s)", ds_otheraggregate_typeIdent, ds_otheraggregate_typeIdent, targetIdent))
			}
		}
	}
	if stage != nil {
		for _, ds_platform_propertytype := range __gong__sortStageSetInstances(stage.DS_Platform_PropertyTypes, stage.DS_Platform_PropertyType_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			ds_platform_propertytypeIdent := ds_platform_propertytype.GongGetIdentifier(stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&gmd.DS_Platform_PropertyType{Name: %s}).Stage(stage)", ds_platform_propertytypeIdent, __gong__toRawStringLiteral(ds_platform_propertytype.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", ds_platform_propertytypeIdent, __gong__toRawStringLiteral(ds_platform_propertytype.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.NilReason = %s", ds_platform_propertytypeIdent, __gong__toRawStringLiteral(ds_platform_propertytype.NilReason)))
			values.WriteString(fmt.Sprintf("\n\t%s.Uuidref = %s", ds_platform_propertytypeIdent, __gong__toRawStringLiteral(ds_platform_propertytype.Uuidref)))
			values.WriteString(fmt.Sprintf("\n\t%s.Type_XSD = %s", ds_platform_propertytypeIdent, __gong__toRawStringLiteral(ds_platform_propertytype.Type_XSD)))
			values.WriteString(fmt.Sprintf("\n\t%s.Href = %s", ds_platform_propertytypeIdent, __gong__toRawStringLiteral(ds_platform_propertytype.Href)))
			values.WriteString(fmt.Sprintf("\n\t%s.Role = %s", ds_platform_propertytypeIdent, __gong__toRawStringLiteral(ds_platform_propertytype.Role)))
			values.WriteString(fmt.Sprintf("\n\t%s.Arcrole = %s", ds_platform_propertytypeIdent, __gong__toRawStringLiteral(ds_platform_propertytype.Arcrole)))
			values.WriteString(fmt.Sprintf("\n\t%s.Title = %s", ds_platform_propertytypeIdent, __gong__toRawStringLiteral(ds_platform_propertytype.Title)))
			values.WriteString(fmt.Sprintf("\n\t%s.Show = %s", ds_platform_propertytypeIdent, __gong__toRawStringLiteral(ds_platform_propertytype.Show)))
			values.WriteString(fmt.Sprintf("\n\t%s.Actuate = %s", ds_platform_propertytypeIdent, __gong__toRawStringLiteral(ds_platform_propertytype.Actuate)))
			if ds_platform_propertytype.DS_Platform != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := ds_platform_propertytype.DS_Platform.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.DS_Platform = %s", ds_platform_propertytypeIdent, targetIdent))
			}
		}
	}
	if stage != nil {
		for _, ds_platform_type := range __gong__sortStageSetInstances(stage.DS_Platform_Types, stage.DS_Platform_Type_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			ds_platform_typeIdent := ds_platform_type.GongGetIdentifier(stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&gmd.DS_Platform_Type{Name: %s}).Stage(stage)", ds_platform_typeIdent, __gong__toRawStringLiteral(ds_platform_type.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", ds_platform_typeIdent, __gong__toRawStringLiteral(ds_platform_type.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Id = %s", ds_platform_typeIdent, __gong__toRawStringLiteral(ds_platform_type.Id)))
			values.WriteString(fmt.Sprintf("\n\t%s.Uuid = %s", ds_platform_typeIdent, __gong__toRawStringLiteral(ds_platform_type.Uuid)))
			for _, elem := range ds_platform_type.ComposedOf {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := elem.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.ComposedOf = append(%s.ComposedOf, %s)", ds_platform_typeIdent, ds_platform_typeIdent, targetIdent))
			}
			for _, elem := range ds_platform_type.SeriesMetadata {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := elem.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.SeriesMetadata = append(%s.SeriesMetadata, %s)", ds_platform_typeIdent, ds_platform_typeIdent, targetIdent))
			}
			for _, elem := range ds_platform_type.Subset {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := elem.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Subset = append(%s.Subset, %s)", ds_platform_typeIdent, ds_platform_typeIdent, targetIdent))
			}
			for _, elem := range ds_platform_type.Superset {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := elem.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Superset = append(%s.Superset, %s)", ds_platform_typeIdent, ds_platform_typeIdent, targetIdent))
			}
		}
	}
	if stage != nil {
		for _, ds_productionseries_propertytype := range __gong__sortStageSetInstances(stage.DS_ProductionSeries_PropertyTypes, stage.DS_ProductionSeries_PropertyType_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			ds_productionseries_propertytypeIdent := ds_productionseries_propertytype.GongGetIdentifier(stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&gmd.DS_ProductionSeries_PropertyType{Name: %s}).Stage(stage)", ds_productionseries_propertytypeIdent, __gong__toRawStringLiteral(ds_productionseries_propertytype.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", ds_productionseries_propertytypeIdent, __gong__toRawStringLiteral(ds_productionseries_propertytype.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.NilReason = %s", ds_productionseries_propertytypeIdent, __gong__toRawStringLiteral(ds_productionseries_propertytype.NilReason)))
			values.WriteString(fmt.Sprintf("\n\t%s.Uuidref = %s", ds_productionseries_propertytypeIdent, __gong__toRawStringLiteral(ds_productionseries_propertytype.Uuidref)))
			values.WriteString(fmt.Sprintf("\n\t%s.Type_XSD = %s", ds_productionseries_propertytypeIdent, __gong__toRawStringLiteral(ds_productionseries_propertytype.Type_XSD)))
			values.WriteString(fmt.Sprintf("\n\t%s.Href = %s", ds_productionseries_propertytypeIdent, __gong__toRawStringLiteral(ds_productionseries_propertytype.Href)))
			values.WriteString(fmt.Sprintf("\n\t%s.Role = %s", ds_productionseries_propertytypeIdent, __gong__toRawStringLiteral(ds_productionseries_propertytype.Role)))
			values.WriteString(fmt.Sprintf("\n\t%s.Arcrole = %s", ds_productionseries_propertytypeIdent, __gong__toRawStringLiteral(ds_productionseries_propertytype.Arcrole)))
			values.WriteString(fmt.Sprintf("\n\t%s.Title = %s", ds_productionseries_propertytypeIdent, __gong__toRawStringLiteral(ds_productionseries_propertytype.Title)))
			values.WriteString(fmt.Sprintf("\n\t%s.Show = %s", ds_productionseries_propertytypeIdent, __gong__toRawStringLiteral(ds_productionseries_propertytype.Show)))
			values.WriteString(fmt.Sprintf("\n\t%s.Actuate = %s", ds_productionseries_propertytypeIdent, __gong__toRawStringLiteral(ds_productionseries_propertytype.Actuate)))
			if ds_productionseries_propertytype.DS_ProductionSeries != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := ds_productionseries_propertytype.DS_ProductionSeries.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.DS_ProductionSeries = %s", ds_productionseries_propertytypeIdent, targetIdent))
			}
		}
	}
	if stage != nil {
		for _, ds_productionseries_type := range __gong__sortStageSetInstances(stage.DS_ProductionSeries_Types, stage.DS_ProductionSeries_Type_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			ds_productionseries_typeIdent := ds_productionseries_type.GongGetIdentifier(stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&gmd.DS_ProductionSeries_Type{Name: %s}).Stage(stage)", ds_productionseries_typeIdent, __gong__toRawStringLiteral(ds_productionseries_type.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", ds_productionseries_typeIdent, __gong__toRawStringLiteral(ds_productionseries_type.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Id = %s", ds_productionseries_typeIdent, __gong__toRawStringLiteral(ds_productionseries_type.Id)))
			values.WriteString(fmt.Sprintf("\n\t%s.Uuid = %s", ds_productionseries_typeIdent, __gong__toRawStringLiteral(ds_productionseries_type.Uuid)))
			for _, elem := range ds_productionseries_type.ComposedOf {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := elem.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.ComposedOf = append(%s.ComposedOf, %s)", ds_productionseries_typeIdent, ds_productionseries_typeIdent, targetIdent))
			}
			for _, elem := range ds_productionseries_type.SeriesMetadata {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := elem.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.SeriesMetadata = append(%s.SeriesMetadata, %s)", ds_productionseries_typeIdent, ds_productionseries_typeIdent, targetIdent))
			}
			for _, elem := range ds_productionseries_type.Subset {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := elem.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Subset = append(%s.Subset, %s)", ds_productionseries_typeIdent, ds_productionseries_typeIdent, targetIdent))
			}
			for _, elem := range ds_productionseries_type.Superset {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := elem.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Superset = append(%s.Superset, %s)", ds_productionseries_typeIdent, ds_productionseries_typeIdent, targetIdent))
			}
		}
	}
	if stage != nil {
		for _, ds_sensor_propertytype := range __gong__sortStageSetInstances(stage.DS_Sensor_PropertyTypes, stage.DS_Sensor_PropertyType_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			ds_sensor_propertytypeIdent := ds_sensor_propertytype.GongGetIdentifier(stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&gmd.DS_Sensor_PropertyType{Name: %s}).Stage(stage)", ds_sensor_propertytypeIdent, __gong__toRawStringLiteral(ds_sensor_propertytype.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", ds_sensor_propertytypeIdent, __gong__toRawStringLiteral(ds_sensor_propertytype.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.NilReason = %s", ds_sensor_propertytypeIdent, __gong__toRawStringLiteral(ds_sensor_propertytype.NilReason)))
			values.WriteString(fmt.Sprintf("\n\t%s.Uuidref = %s", ds_sensor_propertytypeIdent, __gong__toRawStringLiteral(ds_sensor_propertytype.Uuidref)))
			values.WriteString(fmt.Sprintf("\n\t%s.Type_XSD = %s", ds_sensor_propertytypeIdent, __gong__toRawStringLiteral(ds_sensor_propertytype.Type_XSD)))
			values.WriteString(fmt.Sprintf("\n\t%s.Href = %s", ds_sensor_propertytypeIdent, __gong__toRawStringLiteral(ds_sensor_propertytype.Href)))
			values.WriteString(fmt.Sprintf("\n\t%s.Role = %s", ds_sensor_propertytypeIdent, __gong__toRawStringLiteral(ds_sensor_propertytype.Role)))
			values.WriteString(fmt.Sprintf("\n\t%s.Arcrole = %s", ds_sensor_propertytypeIdent, __gong__toRawStringLiteral(ds_sensor_propertytype.Arcrole)))
			values.WriteString(fmt.Sprintf("\n\t%s.Title = %s", ds_sensor_propertytypeIdent, __gong__toRawStringLiteral(ds_sensor_propertytype.Title)))
			values.WriteString(fmt.Sprintf("\n\t%s.Show = %s", ds_sensor_propertytypeIdent, __gong__toRawStringLiteral(ds_sensor_propertytype.Show)))
			values.WriteString(fmt.Sprintf("\n\t%s.Actuate = %s", ds_sensor_propertytypeIdent, __gong__toRawStringLiteral(ds_sensor_propertytype.Actuate)))
			if ds_sensor_propertytype.DS_Sensor != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := ds_sensor_propertytype.DS_Sensor.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.DS_Sensor = %s", ds_sensor_propertytypeIdent, targetIdent))
			}
		}
	}
	if stage != nil {
		for _, ds_sensor_type := range __gong__sortStageSetInstances(stage.DS_Sensor_Types, stage.DS_Sensor_Type_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			ds_sensor_typeIdent := ds_sensor_type.GongGetIdentifier(stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&gmd.DS_Sensor_Type{Name: %s}).Stage(stage)", ds_sensor_typeIdent, __gong__toRawStringLiteral(ds_sensor_type.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", ds_sensor_typeIdent, __gong__toRawStringLiteral(ds_sensor_type.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Id = %s", ds_sensor_typeIdent, __gong__toRawStringLiteral(ds_sensor_type.Id)))
			values.WriteString(fmt.Sprintf("\n\t%s.Uuid = %s", ds_sensor_typeIdent, __gong__toRawStringLiteral(ds_sensor_type.Uuid)))
			for _, elem := range ds_sensor_type.ComposedOf {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := elem.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.ComposedOf = append(%s.ComposedOf, %s)", ds_sensor_typeIdent, ds_sensor_typeIdent, targetIdent))
			}
			for _, elem := range ds_sensor_type.SeriesMetadata {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := elem.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.SeriesMetadata = append(%s.SeriesMetadata, %s)", ds_sensor_typeIdent, ds_sensor_typeIdent, targetIdent))
			}
			for _, elem := range ds_sensor_type.Subset {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := elem.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Subset = append(%s.Subset, %s)", ds_sensor_typeIdent, ds_sensor_typeIdent, targetIdent))
			}
			for _, elem := range ds_sensor_type.Superset {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := elem.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Superset = append(%s.Superset, %s)", ds_sensor_typeIdent, ds_sensor_typeIdent, targetIdent))
			}
		}
	}
	if stage != nil {
		for _, ds_series_propertytype := range __gong__sortStageSetInstances(stage.DS_Series_PropertyTypes, stage.DS_Series_PropertyType_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			ds_series_propertytypeIdent := ds_series_propertytype.GongGetIdentifier(stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&gmd.DS_Series_PropertyType{Name: %s}).Stage(stage)", ds_series_propertytypeIdent, __gong__toRawStringLiteral(ds_series_propertytype.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", ds_series_propertytypeIdent, __gong__toRawStringLiteral(ds_series_propertytype.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.NilReason = %s", ds_series_propertytypeIdent, __gong__toRawStringLiteral(ds_series_propertytype.NilReason)))
			values.WriteString(fmt.Sprintf("\n\t%s.Uuidref = %s", ds_series_propertytypeIdent, __gong__toRawStringLiteral(ds_series_propertytype.Uuidref)))
			values.WriteString(fmt.Sprintf("\n\t%s.Type_XSD = %s", ds_series_propertytypeIdent, __gong__toRawStringLiteral(ds_series_propertytype.Type_XSD)))
			values.WriteString(fmt.Sprintf("\n\t%s.Href = %s", ds_series_propertytypeIdent, __gong__toRawStringLiteral(ds_series_propertytype.Href)))
			values.WriteString(fmt.Sprintf("\n\t%s.Role = %s", ds_series_propertytypeIdent, __gong__toRawStringLiteral(ds_series_propertytype.Role)))
			values.WriteString(fmt.Sprintf("\n\t%s.Arcrole = %s", ds_series_propertytypeIdent, __gong__toRawStringLiteral(ds_series_propertytype.Arcrole)))
			values.WriteString(fmt.Sprintf("\n\t%s.Title = %s", ds_series_propertytypeIdent, __gong__toRawStringLiteral(ds_series_propertytype.Title)))
			values.WriteString(fmt.Sprintf("\n\t%s.Show = %s", ds_series_propertytypeIdent, __gong__toRawStringLiteral(ds_series_propertytype.Show)))
			values.WriteString(fmt.Sprintf("\n\t%s.Actuate = %s", ds_series_propertytypeIdent, __gong__toRawStringLiteral(ds_series_propertytype.Actuate)))
			if ds_series_propertytype.DS_Series != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := ds_series_propertytype.DS_Series.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.DS_Series = %s", ds_series_propertytypeIdent, targetIdent))
			}
			if ds_series_propertytype.DS_Platform != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := ds_series_propertytype.DS_Platform.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.DS_Platform = %s", ds_series_propertytypeIdent, targetIdent))
			}
			if ds_series_propertytype.DS_Sensor != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := ds_series_propertytype.DS_Sensor.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.DS_Sensor = %s", ds_series_propertytypeIdent, targetIdent))
			}
			if ds_series_propertytype.DS_ProductionSeries != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := ds_series_propertytype.DS_ProductionSeries.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.DS_ProductionSeries = %s", ds_series_propertytypeIdent, targetIdent))
			}
		}
	}
	if stage != nil {
		for _, ds_series_type := range __gong__sortStageSetInstances(stage.DS_Series_Types, stage.DS_Series_Type_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			ds_series_typeIdent := ds_series_type.GongGetIdentifier(stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&gmd.DS_Series_Type{Name: %s}).Stage(stage)", ds_series_typeIdent, __gong__toRawStringLiteral(ds_series_type.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", ds_series_typeIdent, __gong__toRawStringLiteral(ds_series_type.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Id = %s", ds_series_typeIdent, __gong__toRawStringLiteral(ds_series_type.Id)))
			values.WriteString(fmt.Sprintf("\n\t%s.Uuid = %s", ds_series_typeIdent, __gong__toRawStringLiteral(ds_series_type.Uuid)))
			for _, elem := range ds_series_type.ComposedOf {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := elem.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.ComposedOf = append(%s.ComposedOf, %s)", ds_series_typeIdent, ds_series_typeIdent, targetIdent))
			}
			for _, elem := range ds_series_type.SeriesMetadata {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := elem.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.SeriesMetadata = append(%s.SeriesMetadata, %s)", ds_series_typeIdent, ds_series_typeIdent, targetIdent))
			}
			for _, elem := range ds_series_type.Subset {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := elem.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Subset = append(%s.Subset, %s)", ds_series_typeIdent, ds_series_typeIdent, targetIdent))
			}
			for _, elem := range ds_series_type.Superset {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := elem.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Superset = append(%s.Superset, %s)", ds_series_typeIdent, ds_series_typeIdent, targetIdent))
			}
		}
	}
	if stage != nil {
		for _, ds_stereomate_propertytype := range __gong__sortStageSetInstances(stage.DS_StereoMate_PropertyTypes, stage.DS_StereoMate_PropertyType_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			ds_stereomate_propertytypeIdent := ds_stereomate_propertytype.GongGetIdentifier(stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&gmd.DS_StereoMate_PropertyType{Name: %s}).Stage(stage)", ds_stereomate_propertytypeIdent, __gong__toRawStringLiteral(ds_stereomate_propertytype.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", ds_stereomate_propertytypeIdent, __gong__toRawStringLiteral(ds_stereomate_propertytype.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.NilReason = %s", ds_stereomate_propertytypeIdent, __gong__toRawStringLiteral(ds_stereomate_propertytype.NilReason)))
			values.WriteString(fmt.Sprintf("\n\t%s.Uuidref = %s", ds_stereomate_propertytypeIdent, __gong__toRawStringLiteral(ds_stereomate_propertytype.Uuidref)))
			values.WriteString(fmt.Sprintf("\n\t%s.Type_XSD = %s", ds_stereomate_propertytypeIdent, __gong__toRawStringLiteral(ds_stereomate_propertytype.Type_XSD)))
			values.WriteString(fmt.Sprintf("\n\t%s.Href = %s", ds_stereomate_propertytypeIdent, __gong__toRawStringLiteral(ds_stereomate_propertytype.Href)))
			values.WriteString(fmt.Sprintf("\n\t%s.Role = %s", ds_stereomate_propertytypeIdent, __gong__toRawStringLiteral(ds_stereomate_propertytype.Role)))
			values.WriteString(fmt.Sprintf("\n\t%s.Arcrole = %s", ds_stereomate_propertytypeIdent, __gong__toRawStringLiteral(ds_stereomate_propertytype.Arcrole)))
			values.WriteString(fmt.Sprintf("\n\t%s.Title = %s", ds_stereomate_propertytypeIdent, __gong__toRawStringLiteral(ds_stereomate_propertytype.Title)))
			values.WriteString(fmt.Sprintf("\n\t%s.Show = %s", ds_stereomate_propertytypeIdent, __gong__toRawStringLiteral(ds_stereomate_propertytype.Show)))
			values.WriteString(fmt.Sprintf("\n\t%s.Actuate = %s", ds_stereomate_propertytypeIdent, __gong__toRawStringLiteral(ds_stereomate_propertytype.Actuate)))
			if ds_stereomate_propertytype.DS_StereoMate != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := ds_stereomate_propertytype.DS_StereoMate.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.DS_StereoMate = %s", ds_stereomate_propertytypeIdent, targetIdent))
			}
		}
	}
	if stage != nil {
		for _, ds_stereomate_type := range __gong__sortStageSetInstances(stage.DS_StereoMate_Types, stage.DS_StereoMate_Type_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			ds_stereomate_typeIdent := ds_stereomate_type.GongGetIdentifier(stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&gmd.DS_StereoMate_Type{Name: %s}).Stage(stage)", ds_stereomate_typeIdent, __gong__toRawStringLiteral(ds_stereomate_type.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", ds_stereomate_typeIdent, __gong__toRawStringLiteral(ds_stereomate_type.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Id = %s", ds_stereomate_typeIdent, __gong__toRawStringLiteral(ds_stereomate_type.Id)))
			values.WriteString(fmt.Sprintf("\n\t%s.Uuid = %s", ds_stereomate_typeIdent, __gong__toRawStringLiteral(ds_stereomate_type.Uuid)))
			for _, elem := range ds_stereomate_type.ComposedOf {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := elem.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.ComposedOf = append(%s.ComposedOf, %s)", ds_stereomate_typeIdent, ds_stereomate_typeIdent, targetIdent))
			}
			for _, elem := range ds_stereomate_type.SeriesMetadata {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := elem.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.SeriesMetadata = append(%s.SeriesMetadata, %s)", ds_stereomate_typeIdent, ds_stereomate_typeIdent, targetIdent))
			}
			for _, elem := range ds_stereomate_type.Subset {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := elem.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Subset = append(%s.Subset, %s)", ds_stereomate_typeIdent, ds_stereomate_typeIdent, targetIdent))
			}
			for _, elem := range ds_stereomate_type.Superset {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := elem.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Superset = append(%s.Superset, %s)", ds_stereomate_typeIdent, ds_stereomate_typeIdent, targetIdent))
			}
		}
	}
	if stage != nil {
		for _, md_metadata_propertytype := range __gong__sortStageSetInstances(stage.MD_Metadata_PropertyTypes, stage.MD_Metadata_PropertyType_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			md_metadata_propertytypeIdent := md_metadata_propertytype.GongGetIdentifier(stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&gmd.MD_Metadata_PropertyType{Name: %s}).Stage(stage)", md_metadata_propertytypeIdent, __gong__toRawStringLiteral(md_metadata_propertytype.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", md_metadata_propertytypeIdent, __gong__toRawStringLiteral(md_metadata_propertytype.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.NilReason = %s", md_metadata_propertytypeIdent, __gong__toRawStringLiteral(md_metadata_propertytype.NilReason)))
			values.WriteString(fmt.Sprintf("\n\t%s.Uuidref = %s", md_metadata_propertytypeIdent, __gong__toRawStringLiteral(md_metadata_propertytype.Uuidref)))
			values.WriteString(fmt.Sprintf("\n\t%s.Type_XSD = %s", md_metadata_propertytypeIdent, __gong__toRawStringLiteral(md_metadata_propertytype.Type_XSD)))
			values.WriteString(fmt.Sprintf("\n\t%s.Href = %s", md_metadata_propertytypeIdent, __gong__toRawStringLiteral(md_metadata_propertytype.Href)))
			values.WriteString(fmt.Sprintf("\n\t%s.Role = %s", md_metadata_propertytypeIdent, __gong__toRawStringLiteral(md_metadata_propertytype.Role)))
			values.WriteString(fmt.Sprintf("\n\t%s.Arcrole = %s", md_metadata_propertytypeIdent, __gong__toRawStringLiteral(md_metadata_propertytype.Arcrole)))
			values.WriteString(fmt.Sprintf("\n\t%s.Title = %s", md_metadata_propertytypeIdent, __gong__toRawStringLiteral(md_metadata_propertytype.Title)))
			values.WriteString(fmt.Sprintf("\n\t%s.Show = %s", md_metadata_propertytypeIdent, __gong__toRawStringLiteral(md_metadata_propertytype.Show)))
			values.WriteString(fmt.Sprintf("\n\t%s.Actuate = %s", md_metadata_propertytypeIdent, __gong__toRawStringLiteral(md_metadata_propertytype.Actuate)))
			if md_metadata_propertytype.MD_Metadata != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := md_metadata_propertytype.MD_Metadata.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.MD_Metadata = %s", md_metadata_propertytypeIdent, targetIdent))
			}
		}
	}
	if stage != nil {
		for _, md_metadata_type := range __gong__sortStageSetInstances(stage.MD_Metadata_Types, stage.MD_Metadata_Type_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			md_metadata_typeIdent := md_metadata_type.GongGetIdentifier(stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&gmd.MD_Metadata_Type{Name: %s}).Stage(stage)", md_metadata_typeIdent, __gong__toRawStringLiteral(md_metadata_type.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", md_metadata_typeIdent, __gong__toRawStringLiteral(md_metadata_type.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Id = %s", md_metadata_typeIdent, __gong__toRawStringLiteral(md_metadata_type.Id)))
			values.WriteString(fmt.Sprintf("\n\t%s.Uuid = %s", md_metadata_typeIdent, __gong__toRawStringLiteral(md_metadata_type.Uuid)))
			for _, elem := range md_metadata_type.Series {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := elem.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Series = append(%s.Series, %s)", md_metadata_typeIdent, md_metadata_typeIdent, targetIdent))
			}
			for _, elem := range md_metadata_type.Describes {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := elem.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Describes = append(%s.Describes, %s)", md_metadata_typeIdent, md_metadata_typeIdent, targetIdent))
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
go run github.com/fullstack-lang/gong/app/xsd/tests/aixm/go/models/gmd/probe/cmd/probe@%s stage.go
*/
package %s

import (
	"slices"
	"time"

	"github.com/fullstack-lang/gong/app/xsd/tests/aixm/go/models/gmd"%s
)

var (
	_ time.Time
	_ = slices.Index[[]int, int]

	_ *gmd.Stage
)%s

// function will stage objects across all coordinated stages
func _(stage *gmd.Stage) {

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

	if modelsPackageName != "" && modelsPackageName != "github.com/fullstack-lang/gong/app/xsd/tests/aixm/go/models/gmd" {
		res = strings.ReplaceAll(res, "\"github.com/fullstack-lang/gong/app/xsd/tests/aixm/go/models/gmd\"", fmt.Sprintf("%q", modelsPackageName))
	}

	return res, nil
}

// insertion point for marshall field methods
func (abstractds_aggregate_type *AbstractDS_Aggregate_Type) GongMarshallField(stage *Stage, fieldName string) (res string) {
	ident := abstractds_aggregate_type.GongGetIdentifier(stage)
	_ = ident
	switch fieldName {
	case "Name":
		res = __gong__marshallString(ident, "Name", abstractds_aggregate_type.Name)
	case "Id":
		res = __gong__marshallString(ident, "Id", abstractds_aggregate_type.Id)
	case "Uuid":
		res = __gong__marshallString(ident, "Uuid", abstractds_aggregate_type.Uuid)

	case "ComposedOf":
		var sb strings.Builder
		for _, _ds_dataset_propertytype := range abstractds_aggregate_type.ComposedOf {
			sb.WriteString(__gong__marshallSliceOfPointers(ident, "ComposedOf", _ds_dataset_propertytype.GongGetIdentifier(stage)))
		}
		res = sb.String()
	case "SeriesMetadata":
		var sb strings.Builder
		for _, _md_metadata_propertytype := range abstractds_aggregate_type.SeriesMetadata {
			sb.WriteString(__gong__marshallSliceOfPointers(ident, "SeriesMetadata", _md_metadata_propertytype.GongGetIdentifier(stage)))
		}
		res = sb.String()
	case "Subset":
		var sb strings.Builder
		for _, _ds_aggregate_propertytype := range abstractds_aggregate_type.Subset {
			sb.WriteString(__gong__marshallSliceOfPointers(ident, "Subset", _ds_aggregate_propertytype.GongGetIdentifier(stage)))
		}
		res = sb.String()
	case "Superset":
		var sb strings.Builder
		for _, _ds_aggregate_propertytype := range abstractds_aggregate_type.Superset {
			sb.WriteString(__gong__marshallSliceOfPointers(ident, "Superset", _ds_aggregate_propertytype.GongGetIdentifier(stage)))
		}
		res = sb.String()
	default:
		log.Panicf("Unknown field %s for Gongstruct AbstractDS_Aggregate_Type", fieldName)
	}
	return
}

func (ds_aggregate_propertytype *DS_Aggregate_PropertyType) GongMarshallField(stage *Stage, fieldName string) (res string) {
	ident := ds_aggregate_propertytype.GongGetIdentifier(stage)
	_ = ident
	switch fieldName {
	case "Name":
		res = __gong__marshallString(ident, "Name", ds_aggregate_propertytype.Name)
	case "NilReason":
		res = __gong__marshallString(ident, "NilReason", ds_aggregate_propertytype.NilReason)
	case "Uuidref":
		res = __gong__marshallString(ident, "Uuidref", ds_aggregate_propertytype.Uuidref)
	case "Type_XSD":
		res = __gong__marshallString(ident, "Type_XSD", ds_aggregate_propertytype.Type_XSD)
	case "Href":
		res = __gong__marshallString(ident, "Href", ds_aggregate_propertytype.Href)
	case "Role":
		res = __gong__marshallString(ident, "Role", ds_aggregate_propertytype.Role)
	case "Arcrole":
		res = __gong__marshallString(ident, "Arcrole", ds_aggregate_propertytype.Arcrole)
	case "Title":
		res = __gong__marshallString(ident, "Title", ds_aggregate_propertytype.Title)
	case "Show":
		res = __gong__marshallString(ident, "Show", ds_aggregate_propertytype.Show)
	case "Actuate":
		res = __gong__marshallString(ident, "Actuate", ds_aggregate_propertytype.Actuate)

	case "DS_OtherAggregate":
		if ds_aggregate_propertytype.DS_OtherAggregate != nil {
			res = __gong__marshallPointer(ident, "DS_OtherAggregate", ds_aggregate_propertytype.DS_OtherAggregate.GongGetIdentifier(stage))
		} else {
			res = __gong__marshallPointer(ident, "DS_OtherAggregate", "nil")
		}
	case "DS_StereoMate":
		if ds_aggregate_propertytype.DS_StereoMate != nil {
			res = __gong__marshallPointer(ident, "DS_StereoMate", ds_aggregate_propertytype.DS_StereoMate.GongGetIdentifier(stage))
		} else {
			res = __gong__marshallPointer(ident, "DS_StereoMate", "nil")
		}
	case "DS_Series":
		if ds_aggregate_propertytype.DS_Series != nil {
			res = __gong__marshallPointer(ident, "DS_Series", ds_aggregate_propertytype.DS_Series.GongGetIdentifier(stage))
		} else {
			res = __gong__marshallPointer(ident, "DS_Series", "nil")
		}
	case "DS_Platform":
		if ds_aggregate_propertytype.DS_Platform != nil {
			res = __gong__marshallPointer(ident, "DS_Platform", ds_aggregate_propertytype.DS_Platform.GongGetIdentifier(stage))
		} else {
			res = __gong__marshallPointer(ident, "DS_Platform", "nil")
		}
	case "DS_Sensor":
		if ds_aggregate_propertytype.DS_Sensor != nil {
			res = __gong__marshallPointer(ident, "DS_Sensor", ds_aggregate_propertytype.DS_Sensor.GongGetIdentifier(stage))
		} else {
			res = __gong__marshallPointer(ident, "DS_Sensor", "nil")
		}
	case "DS_ProductionSeries":
		if ds_aggregate_propertytype.DS_ProductionSeries != nil {
			res = __gong__marshallPointer(ident, "DS_ProductionSeries", ds_aggregate_propertytype.DS_ProductionSeries.GongGetIdentifier(stage))
		} else {
			res = __gong__marshallPointer(ident, "DS_ProductionSeries", "nil")
		}
	case "DS_Initiative":
		if ds_aggregate_propertytype.DS_Initiative != nil {
			res = __gong__marshallPointer(ident, "DS_Initiative", ds_aggregate_propertytype.DS_Initiative.GongGetIdentifier(stage))
		} else {
			res = __gong__marshallPointer(ident, "DS_Initiative", "nil")
		}
	default:
		log.Panicf("Unknown field %s for Gongstruct DS_Aggregate_PropertyType", fieldName)
	}
	return
}

func (ds_dataset_propertytype *DS_DataSet_PropertyType) GongMarshallField(stage *Stage, fieldName string) (res string) {
	ident := ds_dataset_propertytype.GongGetIdentifier(stage)
	_ = ident
	switch fieldName {
	case "Name":
		res = __gong__marshallString(ident, "Name", ds_dataset_propertytype.Name)
	case "NilReason":
		res = __gong__marshallString(ident, "NilReason", ds_dataset_propertytype.NilReason)
	case "Uuidref":
		res = __gong__marshallString(ident, "Uuidref", ds_dataset_propertytype.Uuidref)
	case "Type_XSD":
		res = __gong__marshallString(ident, "Type_XSD", ds_dataset_propertytype.Type_XSD)
	case "Href":
		res = __gong__marshallString(ident, "Href", ds_dataset_propertytype.Href)
	case "Role":
		res = __gong__marshallString(ident, "Role", ds_dataset_propertytype.Role)
	case "Arcrole":
		res = __gong__marshallString(ident, "Arcrole", ds_dataset_propertytype.Arcrole)
	case "Title":
		res = __gong__marshallString(ident, "Title", ds_dataset_propertytype.Title)
	case "Show":
		res = __gong__marshallString(ident, "Show", ds_dataset_propertytype.Show)
	case "Actuate":
		res = __gong__marshallString(ident, "Actuate", ds_dataset_propertytype.Actuate)

	case "DS_DataSet":
		if ds_dataset_propertytype.DS_DataSet != nil {
			res = __gong__marshallPointer(ident, "DS_DataSet", ds_dataset_propertytype.DS_DataSet.GongGetIdentifier(stage))
		} else {
			res = __gong__marshallPointer(ident, "DS_DataSet", "nil")
		}
	default:
		log.Panicf("Unknown field %s for Gongstruct DS_DataSet_PropertyType", fieldName)
	}
	return
}

func (ds_dataset_type *DS_DataSet_Type) GongMarshallField(stage *Stage, fieldName string) (res string) {
	ident := ds_dataset_type.GongGetIdentifier(stage)
	_ = ident
	switch fieldName {
	case "Name":
		res = __gong__marshallString(ident, "Name", ds_dataset_type.Name)
	case "Id":
		res = __gong__marshallString(ident, "Id", ds_dataset_type.Id)
	case "Uuid":
		res = __gong__marshallString(ident, "Uuid", ds_dataset_type.Uuid)

	case "Has":
		var sb strings.Builder
		for _, _md_metadata_propertytype := range ds_dataset_type.Has {
			sb.WriteString(__gong__marshallSliceOfPointers(ident, "Has", _md_metadata_propertytype.GongGetIdentifier(stage)))
		}
		res = sb.String()
	case "PartOf":
		var sb strings.Builder
		for _, _ds_aggregate_propertytype := range ds_dataset_type.PartOf {
			sb.WriteString(__gong__marshallSliceOfPointers(ident, "PartOf", _ds_aggregate_propertytype.GongGetIdentifier(stage)))
		}
		res = sb.String()
	default:
		log.Panicf("Unknown field %s for Gongstruct DS_DataSet_Type", fieldName)
	}
	return
}

func (ds_initiative_propertytype *DS_Initiative_PropertyType) GongMarshallField(stage *Stage, fieldName string) (res string) {
	ident := ds_initiative_propertytype.GongGetIdentifier(stage)
	_ = ident
	switch fieldName {
	case "Name":
		res = __gong__marshallString(ident, "Name", ds_initiative_propertytype.Name)
	case "NilReason":
		res = __gong__marshallString(ident, "NilReason", ds_initiative_propertytype.NilReason)
	case "Uuidref":
		res = __gong__marshallString(ident, "Uuidref", ds_initiative_propertytype.Uuidref)
	case "Type_XSD":
		res = __gong__marshallString(ident, "Type_XSD", ds_initiative_propertytype.Type_XSD)
	case "Href":
		res = __gong__marshallString(ident, "Href", ds_initiative_propertytype.Href)
	case "Role":
		res = __gong__marshallString(ident, "Role", ds_initiative_propertytype.Role)
	case "Arcrole":
		res = __gong__marshallString(ident, "Arcrole", ds_initiative_propertytype.Arcrole)
	case "Title":
		res = __gong__marshallString(ident, "Title", ds_initiative_propertytype.Title)
	case "Show":
		res = __gong__marshallString(ident, "Show", ds_initiative_propertytype.Show)
	case "Actuate":
		res = __gong__marshallString(ident, "Actuate", ds_initiative_propertytype.Actuate)

	case "DS_Initiative":
		if ds_initiative_propertytype.DS_Initiative != nil {
			res = __gong__marshallPointer(ident, "DS_Initiative", ds_initiative_propertytype.DS_Initiative.GongGetIdentifier(stage))
		} else {
			res = __gong__marshallPointer(ident, "DS_Initiative", "nil")
		}
	default:
		log.Panicf("Unknown field %s for Gongstruct DS_Initiative_PropertyType", fieldName)
	}
	return
}

func (ds_initiative_type *DS_Initiative_Type) GongMarshallField(stage *Stage, fieldName string) (res string) {
	ident := ds_initiative_type.GongGetIdentifier(stage)
	_ = ident
	switch fieldName {
	case "Name":
		res = __gong__marshallString(ident, "Name", ds_initiative_type.Name)
	case "Id":
		res = __gong__marshallString(ident, "Id", ds_initiative_type.Id)
	case "Uuid":
		res = __gong__marshallString(ident, "Uuid", ds_initiative_type.Uuid)

	case "ComposedOf":
		var sb strings.Builder
		for _, _ds_dataset_propertytype := range ds_initiative_type.ComposedOf {
			sb.WriteString(__gong__marshallSliceOfPointers(ident, "ComposedOf", _ds_dataset_propertytype.GongGetIdentifier(stage)))
		}
		res = sb.String()
	case "SeriesMetadata":
		var sb strings.Builder
		for _, _md_metadata_propertytype := range ds_initiative_type.SeriesMetadata {
			sb.WriteString(__gong__marshallSliceOfPointers(ident, "SeriesMetadata", _md_metadata_propertytype.GongGetIdentifier(stage)))
		}
		res = sb.String()
	case "Subset":
		var sb strings.Builder
		for _, _ds_aggregate_propertytype := range ds_initiative_type.Subset {
			sb.WriteString(__gong__marshallSliceOfPointers(ident, "Subset", _ds_aggregate_propertytype.GongGetIdentifier(stage)))
		}
		res = sb.String()
	case "Superset":
		var sb strings.Builder
		for _, _ds_aggregate_propertytype := range ds_initiative_type.Superset {
			sb.WriteString(__gong__marshallSliceOfPointers(ident, "Superset", _ds_aggregate_propertytype.GongGetIdentifier(stage)))
		}
		res = sb.String()
	default:
		log.Panicf("Unknown field %s for Gongstruct DS_Initiative_Type", fieldName)
	}
	return
}

func (ds_otheraggregate_propertytype *DS_OtherAggregate_PropertyType) GongMarshallField(stage *Stage, fieldName string) (res string) {
	ident := ds_otheraggregate_propertytype.GongGetIdentifier(stage)
	_ = ident
	switch fieldName {
	case "Name":
		res = __gong__marshallString(ident, "Name", ds_otheraggregate_propertytype.Name)
	case "NilReason":
		res = __gong__marshallString(ident, "NilReason", ds_otheraggregate_propertytype.NilReason)
	case "Uuidref":
		res = __gong__marshallString(ident, "Uuidref", ds_otheraggregate_propertytype.Uuidref)
	case "Type_XSD":
		res = __gong__marshallString(ident, "Type_XSD", ds_otheraggregate_propertytype.Type_XSD)
	case "Href":
		res = __gong__marshallString(ident, "Href", ds_otheraggregate_propertytype.Href)
	case "Role":
		res = __gong__marshallString(ident, "Role", ds_otheraggregate_propertytype.Role)
	case "Arcrole":
		res = __gong__marshallString(ident, "Arcrole", ds_otheraggregate_propertytype.Arcrole)
	case "Title":
		res = __gong__marshallString(ident, "Title", ds_otheraggregate_propertytype.Title)
	case "Show":
		res = __gong__marshallString(ident, "Show", ds_otheraggregate_propertytype.Show)
	case "Actuate":
		res = __gong__marshallString(ident, "Actuate", ds_otheraggregate_propertytype.Actuate)

	case "DS_OtherAggregate":
		if ds_otheraggregate_propertytype.DS_OtherAggregate != nil {
			res = __gong__marshallPointer(ident, "DS_OtherAggregate", ds_otheraggregate_propertytype.DS_OtherAggregate.GongGetIdentifier(stage))
		} else {
			res = __gong__marshallPointer(ident, "DS_OtherAggregate", "nil")
		}
	case "DS_StereoMate":
		if ds_otheraggregate_propertytype.DS_StereoMate != nil {
			res = __gong__marshallPointer(ident, "DS_StereoMate", ds_otheraggregate_propertytype.DS_StereoMate.GongGetIdentifier(stage))
		} else {
			res = __gong__marshallPointer(ident, "DS_StereoMate", "nil")
		}
	default:
		log.Panicf("Unknown field %s for Gongstruct DS_OtherAggregate_PropertyType", fieldName)
	}
	return
}

func (ds_otheraggregate_type *DS_OtherAggregate_Type) GongMarshallField(stage *Stage, fieldName string) (res string) {
	ident := ds_otheraggregate_type.GongGetIdentifier(stage)
	_ = ident
	switch fieldName {
	case "Name":
		res = __gong__marshallString(ident, "Name", ds_otheraggregate_type.Name)
	case "Id":
		res = __gong__marshallString(ident, "Id", ds_otheraggregate_type.Id)
	case "Uuid":
		res = __gong__marshallString(ident, "Uuid", ds_otheraggregate_type.Uuid)

	case "ComposedOf":
		var sb strings.Builder
		for _, _ds_dataset_propertytype := range ds_otheraggregate_type.ComposedOf {
			sb.WriteString(__gong__marshallSliceOfPointers(ident, "ComposedOf", _ds_dataset_propertytype.GongGetIdentifier(stage)))
		}
		res = sb.String()
	case "SeriesMetadata":
		var sb strings.Builder
		for _, _md_metadata_propertytype := range ds_otheraggregate_type.SeriesMetadata {
			sb.WriteString(__gong__marshallSliceOfPointers(ident, "SeriesMetadata", _md_metadata_propertytype.GongGetIdentifier(stage)))
		}
		res = sb.String()
	case "Subset":
		var sb strings.Builder
		for _, _ds_aggregate_propertytype := range ds_otheraggregate_type.Subset {
			sb.WriteString(__gong__marshallSliceOfPointers(ident, "Subset", _ds_aggregate_propertytype.GongGetIdentifier(stage)))
		}
		res = sb.String()
	case "Superset":
		var sb strings.Builder
		for _, _ds_aggregate_propertytype := range ds_otheraggregate_type.Superset {
			sb.WriteString(__gong__marshallSliceOfPointers(ident, "Superset", _ds_aggregate_propertytype.GongGetIdentifier(stage)))
		}
		res = sb.String()
	default:
		log.Panicf("Unknown field %s for Gongstruct DS_OtherAggregate_Type", fieldName)
	}
	return
}

func (ds_platform_propertytype *DS_Platform_PropertyType) GongMarshallField(stage *Stage, fieldName string) (res string) {
	ident := ds_platform_propertytype.GongGetIdentifier(stage)
	_ = ident
	switch fieldName {
	case "Name":
		res = __gong__marshallString(ident, "Name", ds_platform_propertytype.Name)
	case "NilReason":
		res = __gong__marshallString(ident, "NilReason", ds_platform_propertytype.NilReason)
	case "Uuidref":
		res = __gong__marshallString(ident, "Uuidref", ds_platform_propertytype.Uuidref)
	case "Type_XSD":
		res = __gong__marshallString(ident, "Type_XSD", ds_platform_propertytype.Type_XSD)
	case "Href":
		res = __gong__marshallString(ident, "Href", ds_platform_propertytype.Href)
	case "Role":
		res = __gong__marshallString(ident, "Role", ds_platform_propertytype.Role)
	case "Arcrole":
		res = __gong__marshallString(ident, "Arcrole", ds_platform_propertytype.Arcrole)
	case "Title":
		res = __gong__marshallString(ident, "Title", ds_platform_propertytype.Title)
	case "Show":
		res = __gong__marshallString(ident, "Show", ds_platform_propertytype.Show)
	case "Actuate":
		res = __gong__marshallString(ident, "Actuate", ds_platform_propertytype.Actuate)

	case "DS_Platform":
		if ds_platform_propertytype.DS_Platform != nil {
			res = __gong__marshallPointer(ident, "DS_Platform", ds_platform_propertytype.DS_Platform.GongGetIdentifier(stage))
		} else {
			res = __gong__marshallPointer(ident, "DS_Platform", "nil")
		}
	default:
		log.Panicf("Unknown field %s for Gongstruct DS_Platform_PropertyType", fieldName)
	}
	return
}

func (ds_platform_type *DS_Platform_Type) GongMarshallField(stage *Stage, fieldName string) (res string) {
	ident := ds_platform_type.GongGetIdentifier(stage)
	_ = ident
	switch fieldName {
	case "Name":
		res = __gong__marshallString(ident, "Name", ds_platform_type.Name)
	case "Id":
		res = __gong__marshallString(ident, "Id", ds_platform_type.Id)
	case "Uuid":
		res = __gong__marshallString(ident, "Uuid", ds_platform_type.Uuid)

	case "ComposedOf":
		var sb strings.Builder
		for _, _ds_dataset_propertytype := range ds_platform_type.ComposedOf {
			sb.WriteString(__gong__marshallSliceOfPointers(ident, "ComposedOf", _ds_dataset_propertytype.GongGetIdentifier(stage)))
		}
		res = sb.String()
	case "SeriesMetadata":
		var sb strings.Builder
		for _, _md_metadata_propertytype := range ds_platform_type.SeriesMetadata {
			sb.WriteString(__gong__marshallSliceOfPointers(ident, "SeriesMetadata", _md_metadata_propertytype.GongGetIdentifier(stage)))
		}
		res = sb.String()
	case "Subset":
		var sb strings.Builder
		for _, _ds_aggregate_propertytype := range ds_platform_type.Subset {
			sb.WriteString(__gong__marshallSliceOfPointers(ident, "Subset", _ds_aggregate_propertytype.GongGetIdentifier(stage)))
		}
		res = sb.String()
	case "Superset":
		var sb strings.Builder
		for _, _ds_aggregate_propertytype := range ds_platform_type.Superset {
			sb.WriteString(__gong__marshallSliceOfPointers(ident, "Superset", _ds_aggregate_propertytype.GongGetIdentifier(stage)))
		}
		res = sb.String()
	default:
		log.Panicf("Unknown field %s for Gongstruct DS_Platform_Type", fieldName)
	}
	return
}

func (ds_productionseries_propertytype *DS_ProductionSeries_PropertyType) GongMarshallField(stage *Stage, fieldName string) (res string) {
	ident := ds_productionseries_propertytype.GongGetIdentifier(stage)
	_ = ident
	switch fieldName {
	case "Name":
		res = __gong__marshallString(ident, "Name", ds_productionseries_propertytype.Name)
	case "NilReason":
		res = __gong__marshallString(ident, "NilReason", ds_productionseries_propertytype.NilReason)
	case "Uuidref":
		res = __gong__marshallString(ident, "Uuidref", ds_productionseries_propertytype.Uuidref)
	case "Type_XSD":
		res = __gong__marshallString(ident, "Type_XSD", ds_productionseries_propertytype.Type_XSD)
	case "Href":
		res = __gong__marshallString(ident, "Href", ds_productionseries_propertytype.Href)
	case "Role":
		res = __gong__marshallString(ident, "Role", ds_productionseries_propertytype.Role)
	case "Arcrole":
		res = __gong__marshallString(ident, "Arcrole", ds_productionseries_propertytype.Arcrole)
	case "Title":
		res = __gong__marshallString(ident, "Title", ds_productionseries_propertytype.Title)
	case "Show":
		res = __gong__marshallString(ident, "Show", ds_productionseries_propertytype.Show)
	case "Actuate":
		res = __gong__marshallString(ident, "Actuate", ds_productionseries_propertytype.Actuate)

	case "DS_ProductionSeries":
		if ds_productionseries_propertytype.DS_ProductionSeries != nil {
			res = __gong__marshallPointer(ident, "DS_ProductionSeries", ds_productionseries_propertytype.DS_ProductionSeries.GongGetIdentifier(stage))
		} else {
			res = __gong__marshallPointer(ident, "DS_ProductionSeries", "nil")
		}
	default:
		log.Panicf("Unknown field %s for Gongstruct DS_ProductionSeries_PropertyType", fieldName)
	}
	return
}

func (ds_productionseries_type *DS_ProductionSeries_Type) GongMarshallField(stage *Stage, fieldName string) (res string) {
	ident := ds_productionseries_type.GongGetIdentifier(stage)
	_ = ident
	switch fieldName {
	case "Name":
		res = __gong__marshallString(ident, "Name", ds_productionseries_type.Name)
	case "Id":
		res = __gong__marshallString(ident, "Id", ds_productionseries_type.Id)
	case "Uuid":
		res = __gong__marshallString(ident, "Uuid", ds_productionseries_type.Uuid)

	case "ComposedOf":
		var sb strings.Builder
		for _, _ds_dataset_propertytype := range ds_productionseries_type.ComposedOf {
			sb.WriteString(__gong__marshallSliceOfPointers(ident, "ComposedOf", _ds_dataset_propertytype.GongGetIdentifier(stage)))
		}
		res = sb.String()
	case "SeriesMetadata":
		var sb strings.Builder
		for _, _md_metadata_propertytype := range ds_productionseries_type.SeriesMetadata {
			sb.WriteString(__gong__marshallSliceOfPointers(ident, "SeriesMetadata", _md_metadata_propertytype.GongGetIdentifier(stage)))
		}
		res = sb.String()
	case "Subset":
		var sb strings.Builder
		for _, _ds_aggregate_propertytype := range ds_productionseries_type.Subset {
			sb.WriteString(__gong__marshallSliceOfPointers(ident, "Subset", _ds_aggregate_propertytype.GongGetIdentifier(stage)))
		}
		res = sb.String()
	case "Superset":
		var sb strings.Builder
		for _, _ds_aggregate_propertytype := range ds_productionseries_type.Superset {
			sb.WriteString(__gong__marshallSliceOfPointers(ident, "Superset", _ds_aggregate_propertytype.GongGetIdentifier(stage)))
		}
		res = sb.String()
	default:
		log.Panicf("Unknown field %s for Gongstruct DS_ProductionSeries_Type", fieldName)
	}
	return
}

func (ds_sensor_propertytype *DS_Sensor_PropertyType) GongMarshallField(stage *Stage, fieldName string) (res string) {
	ident := ds_sensor_propertytype.GongGetIdentifier(stage)
	_ = ident
	switch fieldName {
	case "Name":
		res = __gong__marshallString(ident, "Name", ds_sensor_propertytype.Name)
	case "NilReason":
		res = __gong__marshallString(ident, "NilReason", ds_sensor_propertytype.NilReason)
	case "Uuidref":
		res = __gong__marshallString(ident, "Uuidref", ds_sensor_propertytype.Uuidref)
	case "Type_XSD":
		res = __gong__marshallString(ident, "Type_XSD", ds_sensor_propertytype.Type_XSD)
	case "Href":
		res = __gong__marshallString(ident, "Href", ds_sensor_propertytype.Href)
	case "Role":
		res = __gong__marshallString(ident, "Role", ds_sensor_propertytype.Role)
	case "Arcrole":
		res = __gong__marshallString(ident, "Arcrole", ds_sensor_propertytype.Arcrole)
	case "Title":
		res = __gong__marshallString(ident, "Title", ds_sensor_propertytype.Title)
	case "Show":
		res = __gong__marshallString(ident, "Show", ds_sensor_propertytype.Show)
	case "Actuate":
		res = __gong__marshallString(ident, "Actuate", ds_sensor_propertytype.Actuate)

	case "DS_Sensor":
		if ds_sensor_propertytype.DS_Sensor != nil {
			res = __gong__marshallPointer(ident, "DS_Sensor", ds_sensor_propertytype.DS_Sensor.GongGetIdentifier(stage))
		} else {
			res = __gong__marshallPointer(ident, "DS_Sensor", "nil")
		}
	default:
		log.Panicf("Unknown field %s for Gongstruct DS_Sensor_PropertyType", fieldName)
	}
	return
}

func (ds_sensor_type *DS_Sensor_Type) GongMarshallField(stage *Stage, fieldName string) (res string) {
	ident := ds_sensor_type.GongGetIdentifier(stage)
	_ = ident
	switch fieldName {
	case "Name":
		res = __gong__marshallString(ident, "Name", ds_sensor_type.Name)
	case "Id":
		res = __gong__marshallString(ident, "Id", ds_sensor_type.Id)
	case "Uuid":
		res = __gong__marshallString(ident, "Uuid", ds_sensor_type.Uuid)

	case "ComposedOf":
		var sb strings.Builder
		for _, _ds_dataset_propertytype := range ds_sensor_type.ComposedOf {
			sb.WriteString(__gong__marshallSliceOfPointers(ident, "ComposedOf", _ds_dataset_propertytype.GongGetIdentifier(stage)))
		}
		res = sb.String()
	case "SeriesMetadata":
		var sb strings.Builder
		for _, _md_metadata_propertytype := range ds_sensor_type.SeriesMetadata {
			sb.WriteString(__gong__marshallSliceOfPointers(ident, "SeriesMetadata", _md_metadata_propertytype.GongGetIdentifier(stage)))
		}
		res = sb.String()
	case "Subset":
		var sb strings.Builder
		for _, _ds_aggregate_propertytype := range ds_sensor_type.Subset {
			sb.WriteString(__gong__marshallSliceOfPointers(ident, "Subset", _ds_aggregate_propertytype.GongGetIdentifier(stage)))
		}
		res = sb.String()
	case "Superset":
		var sb strings.Builder
		for _, _ds_aggregate_propertytype := range ds_sensor_type.Superset {
			sb.WriteString(__gong__marshallSliceOfPointers(ident, "Superset", _ds_aggregate_propertytype.GongGetIdentifier(stage)))
		}
		res = sb.String()
	default:
		log.Panicf("Unknown field %s for Gongstruct DS_Sensor_Type", fieldName)
	}
	return
}

func (ds_series_propertytype *DS_Series_PropertyType) GongMarshallField(stage *Stage, fieldName string) (res string) {
	ident := ds_series_propertytype.GongGetIdentifier(stage)
	_ = ident
	switch fieldName {
	case "Name":
		res = __gong__marshallString(ident, "Name", ds_series_propertytype.Name)
	case "NilReason":
		res = __gong__marshallString(ident, "NilReason", ds_series_propertytype.NilReason)
	case "Uuidref":
		res = __gong__marshallString(ident, "Uuidref", ds_series_propertytype.Uuidref)
	case "Type_XSD":
		res = __gong__marshallString(ident, "Type_XSD", ds_series_propertytype.Type_XSD)
	case "Href":
		res = __gong__marshallString(ident, "Href", ds_series_propertytype.Href)
	case "Role":
		res = __gong__marshallString(ident, "Role", ds_series_propertytype.Role)
	case "Arcrole":
		res = __gong__marshallString(ident, "Arcrole", ds_series_propertytype.Arcrole)
	case "Title":
		res = __gong__marshallString(ident, "Title", ds_series_propertytype.Title)
	case "Show":
		res = __gong__marshallString(ident, "Show", ds_series_propertytype.Show)
	case "Actuate":
		res = __gong__marshallString(ident, "Actuate", ds_series_propertytype.Actuate)

	case "DS_Series":
		if ds_series_propertytype.DS_Series != nil {
			res = __gong__marshallPointer(ident, "DS_Series", ds_series_propertytype.DS_Series.GongGetIdentifier(stage))
		} else {
			res = __gong__marshallPointer(ident, "DS_Series", "nil")
		}
	case "DS_Platform":
		if ds_series_propertytype.DS_Platform != nil {
			res = __gong__marshallPointer(ident, "DS_Platform", ds_series_propertytype.DS_Platform.GongGetIdentifier(stage))
		} else {
			res = __gong__marshallPointer(ident, "DS_Platform", "nil")
		}
	case "DS_Sensor":
		if ds_series_propertytype.DS_Sensor != nil {
			res = __gong__marshallPointer(ident, "DS_Sensor", ds_series_propertytype.DS_Sensor.GongGetIdentifier(stage))
		} else {
			res = __gong__marshallPointer(ident, "DS_Sensor", "nil")
		}
	case "DS_ProductionSeries":
		if ds_series_propertytype.DS_ProductionSeries != nil {
			res = __gong__marshallPointer(ident, "DS_ProductionSeries", ds_series_propertytype.DS_ProductionSeries.GongGetIdentifier(stage))
		} else {
			res = __gong__marshallPointer(ident, "DS_ProductionSeries", "nil")
		}
	default:
		log.Panicf("Unknown field %s for Gongstruct DS_Series_PropertyType", fieldName)
	}
	return
}

func (ds_series_type *DS_Series_Type) GongMarshallField(stage *Stage, fieldName string) (res string) {
	ident := ds_series_type.GongGetIdentifier(stage)
	_ = ident
	switch fieldName {
	case "Name":
		res = __gong__marshallString(ident, "Name", ds_series_type.Name)
	case "Id":
		res = __gong__marshallString(ident, "Id", ds_series_type.Id)
	case "Uuid":
		res = __gong__marshallString(ident, "Uuid", ds_series_type.Uuid)

	case "ComposedOf":
		var sb strings.Builder
		for _, _ds_dataset_propertytype := range ds_series_type.ComposedOf {
			sb.WriteString(__gong__marshallSliceOfPointers(ident, "ComposedOf", _ds_dataset_propertytype.GongGetIdentifier(stage)))
		}
		res = sb.String()
	case "SeriesMetadata":
		var sb strings.Builder
		for _, _md_metadata_propertytype := range ds_series_type.SeriesMetadata {
			sb.WriteString(__gong__marshallSliceOfPointers(ident, "SeriesMetadata", _md_metadata_propertytype.GongGetIdentifier(stage)))
		}
		res = sb.String()
	case "Subset":
		var sb strings.Builder
		for _, _ds_aggregate_propertytype := range ds_series_type.Subset {
			sb.WriteString(__gong__marshallSliceOfPointers(ident, "Subset", _ds_aggregate_propertytype.GongGetIdentifier(stage)))
		}
		res = sb.String()
	case "Superset":
		var sb strings.Builder
		for _, _ds_aggregate_propertytype := range ds_series_type.Superset {
			sb.WriteString(__gong__marshallSliceOfPointers(ident, "Superset", _ds_aggregate_propertytype.GongGetIdentifier(stage)))
		}
		res = sb.String()
	default:
		log.Panicf("Unknown field %s for Gongstruct DS_Series_Type", fieldName)
	}
	return
}

func (ds_stereomate_propertytype *DS_StereoMate_PropertyType) GongMarshallField(stage *Stage, fieldName string) (res string) {
	ident := ds_stereomate_propertytype.GongGetIdentifier(stage)
	_ = ident
	switch fieldName {
	case "Name":
		res = __gong__marshallString(ident, "Name", ds_stereomate_propertytype.Name)
	case "NilReason":
		res = __gong__marshallString(ident, "NilReason", ds_stereomate_propertytype.NilReason)
	case "Uuidref":
		res = __gong__marshallString(ident, "Uuidref", ds_stereomate_propertytype.Uuidref)
	case "Type_XSD":
		res = __gong__marshallString(ident, "Type_XSD", ds_stereomate_propertytype.Type_XSD)
	case "Href":
		res = __gong__marshallString(ident, "Href", ds_stereomate_propertytype.Href)
	case "Role":
		res = __gong__marshallString(ident, "Role", ds_stereomate_propertytype.Role)
	case "Arcrole":
		res = __gong__marshallString(ident, "Arcrole", ds_stereomate_propertytype.Arcrole)
	case "Title":
		res = __gong__marshallString(ident, "Title", ds_stereomate_propertytype.Title)
	case "Show":
		res = __gong__marshallString(ident, "Show", ds_stereomate_propertytype.Show)
	case "Actuate":
		res = __gong__marshallString(ident, "Actuate", ds_stereomate_propertytype.Actuate)

	case "DS_StereoMate":
		if ds_stereomate_propertytype.DS_StereoMate != nil {
			res = __gong__marshallPointer(ident, "DS_StereoMate", ds_stereomate_propertytype.DS_StereoMate.GongGetIdentifier(stage))
		} else {
			res = __gong__marshallPointer(ident, "DS_StereoMate", "nil")
		}
	default:
		log.Panicf("Unknown field %s for Gongstruct DS_StereoMate_PropertyType", fieldName)
	}
	return
}

func (ds_stereomate_type *DS_StereoMate_Type) GongMarshallField(stage *Stage, fieldName string) (res string) {
	ident := ds_stereomate_type.GongGetIdentifier(stage)
	_ = ident
	switch fieldName {
	case "Name":
		res = __gong__marshallString(ident, "Name", ds_stereomate_type.Name)
	case "Id":
		res = __gong__marshallString(ident, "Id", ds_stereomate_type.Id)
	case "Uuid":
		res = __gong__marshallString(ident, "Uuid", ds_stereomate_type.Uuid)

	case "ComposedOf":
		var sb strings.Builder
		for _, _ds_dataset_propertytype := range ds_stereomate_type.ComposedOf {
			sb.WriteString(__gong__marshallSliceOfPointers(ident, "ComposedOf", _ds_dataset_propertytype.GongGetIdentifier(stage)))
		}
		res = sb.String()
	case "SeriesMetadata":
		var sb strings.Builder
		for _, _md_metadata_propertytype := range ds_stereomate_type.SeriesMetadata {
			sb.WriteString(__gong__marshallSliceOfPointers(ident, "SeriesMetadata", _md_metadata_propertytype.GongGetIdentifier(stage)))
		}
		res = sb.String()
	case "Subset":
		var sb strings.Builder
		for _, _ds_aggregate_propertytype := range ds_stereomate_type.Subset {
			sb.WriteString(__gong__marshallSliceOfPointers(ident, "Subset", _ds_aggregate_propertytype.GongGetIdentifier(stage)))
		}
		res = sb.String()
	case "Superset":
		var sb strings.Builder
		for _, _ds_aggregate_propertytype := range ds_stereomate_type.Superset {
			sb.WriteString(__gong__marshallSliceOfPointers(ident, "Superset", _ds_aggregate_propertytype.GongGetIdentifier(stage)))
		}
		res = sb.String()
	default:
		log.Panicf("Unknown field %s for Gongstruct DS_StereoMate_Type", fieldName)
	}
	return
}

func (md_metadata_propertytype *MD_Metadata_PropertyType) GongMarshallField(stage *Stage, fieldName string) (res string) {
	ident := md_metadata_propertytype.GongGetIdentifier(stage)
	_ = ident
	switch fieldName {
	case "Name":
		res = __gong__marshallString(ident, "Name", md_metadata_propertytype.Name)
	case "NilReason":
		res = __gong__marshallString(ident, "NilReason", md_metadata_propertytype.NilReason)
	case "Uuidref":
		res = __gong__marshallString(ident, "Uuidref", md_metadata_propertytype.Uuidref)
	case "Type_XSD":
		res = __gong__marshallString(ident, "Type_XSD", md_metadata_propertytype.Type_XSD)
	case "Href":
		res = __gong__marshallString(ident, "Href", md_metadata_propertytype.Href)
	case "Role":
		res = __gong__marshallString(ident, "Role", md_metadata_propertytype.Role)
	case "Arcrole":
		res = __gong__marshallString(ident, "Arcrole", md_metadata_propertytype.Arcrole)
	case "Title":
		res = __gong__marshallString(ident, "Title", md_metadata_propertytype.Title)
	case "Show":
		res = __gong__marshallString(ident, "Show", md_metadata_propertytype.Show)
	case "Actuate":
		res = __gong__marshallString(ident, "Actuate", md_metadata_propertytype.Actuate)

	case "MD_Metadata":
		if md_metadata_propertytype.MD_Metadata != nil {
			res = __gong__marshallPointer(ident, "MD_Metadata", md_metadata_propertytype.MD_Metadata.GongGetIdentifier(stage))
		} else {
			res = __gong__marshallPointer(ident, "MD_Metadata", "nil")
		}
	default:
		log.Panicf("Unknown field %s for Gongstruct MD_Metadata_PropertyType", fieldName)
	}
	return
}

func (md_metadata_type *MD_Metadata_Type) GongMarshallField(stage *Stage, fieldName string) (res string) {
	ident := md_metadata_type.GongGetIdentifier(stage)
	_ = ident
	switch fieldName {
	case "Name":
		res = __gong__marshallString(ident, "Name", md_metadata_type.Name)
	case "Id":
		res = __gong__marshallString(ident, "Id", md_metadata_type.Id)
	case "Uuid":
		res = __gong__marshallString(ident, "Uuid", md_metadata_type.Uuid)

	case "Series":
		var sb strings.Builder
		for _, _ds_aggregate_propertytype := range md_metadata_type.Series {
			sb.WriteString(__gong__marshallSliceOfPointers(ident, "Series", _ds_aggregate_propertytype.GongGetIdentifier(stage)))
		}
		res = sb.String()
	case "Describes":
		var sb strings.Builder
		for _, _ds_dataset_propertytype := range md_metadata_type.Describes {
			sb.WriteString(__gong__marshallSliceOfPointers(ident, "Describes", _ds_dataset_propertytype.GongGetIdentifier(stage)))
		}
		res = sb.String()
	default:
		log.Panicf("Unknown field %s for Gongstruct MD_Metadata_Type", fieldName)
	}
	return
}

// insertion point for marshall all fields methods
func (abstractds_aggregate_type *AbstractDS_Aggregate_Type) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(abstractds_aggregate_type.GongMarshallField(stage, "Name"))
		initializerStatements.WriteString(abstractds_aggregate_type.GongMarshallField(stage, "Id"))
		initializerStatements.WriteString(abstractds_aggregate_type.GongMarshallField(stage, "Uuid"))
		pointersInitializesStatements.WriteString(abstractds_aggregate_type.GongMarshallField(stage, "ComposedOf"))
		pointersInitializesStatements.WriteString(abstractds_aggregate_type.GongMarshallField(stage, "SeriesMetadata"))
		pointersInitializesStatements.WriteString(abstractds_aggregate_type.GongMarshallField(stage, "Subset"))
		pointersInitializesStatements.WriteString(abstractds_aggregate_type.GongMarshallField(stage, "Superset"))
	}
	initRes = initializerStatements.String()
	ptrRes = pointersInitializesStatements.String()
	return
}
func (ds_aggregate_propertytype *DS_Aggregate_PropertyType) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(ds_aggregate_propertytype.GongMarshallField(stage, "Name"))
		initializerStatements.WriteString(ds_aggregate_propertytype.GongMarshallField(stage, "NilReason"))
		initializerStatements.WriteString(ds_aggregate_propertytype.GongMarshallField(stage, "Uuidref"))
		initializerStatements.WriteString(ds_aggregate_propertytype.GongMarshallField(stage, "Type_XSD"))
		initializerStatements.WriteString(ds_aggregate_propertytype.GongMarshallField(stage, "Href"))
		initializerStatements.WriteString(ds_aggregate_propertytype.GongMarshallField(stage, "Role"))
		initializerStatements.WriteString(ds_aggregate_propertytype.GongMarshallField(stage, "Arcrole"))
		initializerStatements.WriteString(ds_aggregate_propertytype.GongMarshallField(stage, "Title"))
		initializerStatements.WriteString(ds_aggregate_propertytype.GongMarshallField(stage, "Show"))
		initializerStatements.WriteString(ds_aggregate_propertytype.GongMarshallField(stage, "Actuate"))
		pointersInitializesStatements.WriteString(ds_aggregate_propertytype.GongMarshallField(stage, "DS_OtherAggregate"))
		pointersInitializesStatements.WriteString(ds_aggregate_propertytype.GongMarshallField(stage, "DS_StereoMate"))
		pointersInitializesStatements.WriteString(ds_aggregate_propertytype.GongMarshallField(stage, "DS_Series"))
		pointersInitializesStatements.WriteString(ds_aggregate_propertytype.GongMarshallField(stage, "DS_Platform"))
		pointersInitializesStatements.WriteString(ds_aggregate_propertytype.GongMarshallField(stage, "DS_Sensor"))
		pointersInitializesStatements.WriteString(ds_aggregate_propertytype.GongMarshallField(stage, "DS_ProductionSeries"))
		pointersInitializesStatements.WriteString(ds_aggregate_propertytype.GongMarshallField(stage, "DS_Initiative"))
	}
	initRes = initializerStatements.String()
	ptrRes = pointersInitializesStatements.String()
	return
}
func (ds_dataset_propertytype *DS_DataSet_PropertyType) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(ds_dataset_propertytype.GongMarshallField(stage, "Name"))
		initializerStatements.WriteString(ds_dataset_propertytype.GongMarshallField(stage, "NilReason"))
		initializerStatements.WriteString(ds_dataset_propertytype.GongMarshallField(stage, "Uuidref"))
		initializerStatements.WriteString(ds_dataset_propertytype.GongMarshallField(stage, "Type_XSD"))
		initializerStatements.WriteString(ds_dataset_propertytype.GongMarshallField(stage, "Href"))
		initializerStatements.WriteString(ds_dataset_propertytype.GongMarshallField(stage, "Role"))
		initializerStatements.WriteString(ds_dataset_propertytype.GongMarshallField(stage, "Arcrole"))
		initializerStatements.WriteString(ds_dataset_propertytype.GongMarshallField(stage, "Title"))
		initializerStatements.WriteString(ds_dataset_propertytype.GongMarshallField(stage, "Show"))
		initializerStatements.WriteString(ds_dataset_propertytype.GongMarshallField(stage, "Actuate"))
		pointersInitializesStatements.WriteString(ds_dataset_propertytype.GongMarshallField(stage, "DS_DataSet"))
	}
	initRes = initializerStatements.String()
	ptrRes = pointersInitializesStatements.String()
	return
}
func (ds_dataset_type *DS_DataSet_Type) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(ds_dataset_type.GongMarshallField(stage, "Name"))
		initializerStatements.WriteString(ds_dataset_type.GongMarshallField(stage, "Id"))
		initializerStatements.WriteString(ds_dataset_type.GongMarshallField(stage, "Uuid"))
		pointersInitializesStatements.WriteString(ds_dataset_type.GongMarshallField(stage, "Has"))
		pointersInitializesStatements.WriteString(ds_dataset_type.GongMarshallField(stage, "PartOf"))
	}
	initRes = initializerStatements.String()
	ptrRes = pointersInitializesStatements.String()
	return
}
func (ds_initiative_propertytype *DS_Initiative_PropertyType) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(ds_initiative_propertytype.GongMarshallField(stage, "Name"))
		initializerStatements.WriteString(ds_initiative_propertytype.GongMarshallField(stage, "NilReason"))
		initializerStatements.WriteString(ds_initiative_propertytype.GongMarshallField(stage, "Uuidref"))
		initializerStatements.WriteString(ds_initiative_propertytype.GongMarshallField(stage, "Type_XSD"))
		initializerStatements.WriteString(ds_initiative_propertytype.GongMarshallField(stage, "Href"))
		initializerStatements.WriteString(ds_initiative_propertytype.GongMarshallField(stage, "Role"))
		initializerStatements.WriteString(ds_initiative_propertytype.GongMarshallField(stage, "Arcrole"))
		initializerStatements.WriteString(ds_initiative_propertytype.GongMarshallField(stage, "Title"))
		initializerStatements.WriteString(ds_initiative_propertytype.GongMarshallField(stage, "Show"))
		initializerStatements.WriteString(ds_initiative_propertytype.GongMarshallField(stage, "Actuate"))
		pointersInitializesStatements.WriteString(ds_initiative_propertytype.GongMarshallField(stage, "DS_Initiative"))
	}
	initRes = initializerStatements.String()
	ptrRes = pointersInitializesStatements.String()
	return
}
func (ds_initiative_type *DS_Initiative_Type) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(ds_initiative_type.GongMarshallField(stage, "Name"))
		initializerStatements.WriteString(ds_initiative_type.GongMarshallField(stage, "Id"))
		initializerStatements.WriteString(ds_initiative_type.GongMarshallField(stage, "Uuid"))
		pointersInitializesStatements.WriteString(ds_initiative_type.GongMarshallField(stage, "ComposedOf"))
		pointersInitializesStatements.WriteString(ds_initiative_type.GongMarshallField(stage, "SeriesMetadata"))
		pointersInitializesStatements.WriteString(ds_initiative_type.GongMarshallField(stage, "Subset"))
		pointersInitializesStatements.WriteString(ds_initiative_type.GongMarshallField(stage, "Superset"))
	}
	initRes = initializerStatements.String()
	ptrRes = pointersInitializesStatements.String()
	return
}
func (ds_otheraggregate_propertytype *DS_OtherAggregate_PropertyType) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(ds_otheraggregate_propertytype.GongMarshallField(stage, "Name"))
		initializerStatements.WriteString(ds_otheraggregate_propertytype.GongMarshallField(stage, "NilReason"))
		initializerStatements.WriteString(ds_otheraggregate_propertytype.GongMarshallField(stage, "Uuidref"))
		initializerStatements.WriteString(ds_otheraggregate_propertytype.GongMarshallField(stage, "Type_XSD"))
		initializerStatements.WriteString(ds_otheraggregate_propertytype.GongMarshallField(stage, "Href"))
		initializerStatements.WriteString(ds_otheraggregate_propertytype.GongMarshallField(stage, "Role"))
		initializerStatements.WriteString(ds_otheraggregate_propertytype.GongMarshallField(stage, "Arcrole"))
		initializerStatements.WriteString(ds_otheraggregate_propertytype.GongMarshallField(stage, "Title"))
		initializerStatements.WriteString(ds_otheraggregate_propertytype.GongMarshallField(stage, "Show"))
		initializerStatements.WriteString(ds_otheraggregate_propertytype.GongMarshallField(stage, "Actuate"))
		pointersInitializesStatements.WriteString(ds_otheraggregate_propertytype.GongMarshallField(stage, "DS_OtherAggregate"))
		pointersInitializesStatements.WriteString(ds_otheraggregate_propertytype.GongMarshallField(stage, "DS_StereoMate"))
	}
	initRes = initializerStatements.String()
	ptrRes = pointersInitializesStatements.String()
	return
}
func (ds_otheraggregate_type *DS_OtherAggregate_Type) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(ds_otheraggregate_type.GongMarshallField(stage, "Name"))
		initializerStatements.WriteString(ds_otheraggregate_type.GongMarshallField(stage, "Id"))
		initializerStatements.WriteString(ds_otheraggregate_type.GongMarshallField(stage, "Uuid"))
		pointersInitializesStatements.WriteString(ds_otheraggregate_type.GongMarshallField(stage, "ComposedOf"))
		pointersInitializesStatements.WriteString(ds_otheraggregate_type.GongMarshallField(stage, "SeriesMetadata"))
		pointersInitializesStatements.WriteString(ds_otheraggregate_type.GongMarshallField(stage, "Subset"))
		pointersInitializesStatements.WriteString(ds_otheraggregate_type.GongMarshallField(stage, "Superset"))
	}
	initRes = initializerStatements.String()
	ptrRes = pointersInitializesStatements.String()
	return
}
func (ds_platform_propertytype *DS_Platform_PropertyType) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(ds_platform_propertytype.GongMarshallField(stage, "Name"))
		initializerStatements.WriteString(ds_platform_propertytype.GongMarshallField(stage, "NilReason"))
		initializerStatements.WriteString(ds_platform_propertytype.GongMarshallField(stage, "Uuidref"))
		initializerStatements.WriteString(ds_platform_propertytype.GongMarshallField(stage, "Type_XSD"))
		initializerStatements.WriteString(ds_platform_propertytype.GongMarshallField(stage, "Href"))
		initializerStatements.WriteString(ds_platform_propertytype.GongMarshallField(stage, "Role"))
		initializerStatements.WriteString(ds_platform_propertytype.GongMarshallField(stage, "Arcrole"))
		initializerStatements.WriteString(ds_platform_propertytype.GongMarshallField(stage, "Title"))
		initializerStatements.WriteString(ds_platform_propertytype.GongMarshallField(stage, "Show"))
		initializerStatements.WriteString(ds_platform_propertytype.GongMarshallField(stage, "Actuate"))
		pointersInitializesStatements.WriteString(ds_platform_propertytype.GongMarshallField(stage, "DS_Platform"))
	}
	initRes = initializerStatements.String()
	ptrRes = pointersInitializesStatements.String()
	return
}
func (ds_platform_type *DS_Platform_Type) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(ds_platform_type.GongMarshallField(stage, "Name"))
		initializerStatements.WriteString(ds_platform_type.GongMarshallField(stage, "Id"))
		initializerStatements.WriteString(ds_platform_type.GongMarshallField(stage, "Uuid"))
		pointersInitializesStatements.WriteString(ds_platform_type.GongMarshallField(stage, "ComposedOf"))
		pointersInitializesStatements.WriteString(ds_platform_type.GongMarshallField(stage, "SeriesMetadata"))
		pointersInitializesStatements.WriteString(ds_platform_type.GongMarshallField(stage, "Subset"))
		pointersInitializesStatements.WriteString(ds_platform_type.GongMarshallField(stage, "Superset"))
	}
	initRes = initializerStatements.String()
	ptrRes = pointersInitializesStatements.String()
	return
}
func (ds_productionseries_propertytype *DS_ProductionSeries_PropertyType) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(ds_productionseries_propertytype.GongMarshallField(stage, "Name"))
		initializerStatements.WriteString(ds_productionseries_propertytype.GongMarshallField(stage, "NilReason"))
		initializerStatements.WriteString(ds_productionseries_propertytype.GongMarshallField(stage, "Uuidref"))
		initializerStatements.WriteString(ds_productionseries_propertytype.GongMarshallField(stage, "Type_XSD"))
		initializerStatements.WriteString(ds_productionseries_propertytype.GongMarshallField(stage, "Href"))
		initializerStatements.WriteString(ds_productionseries_propertytype.GongMarshallField(stage, "Role"))
		initializerStatements.WriteString(ds_productionseries_propertytype.GongMarshallField(stage, "Arcrole"))
		initializerStatements.WriteString(ds_productionseries_propertytype.GongMarshallField(stage, "Title"))
		initializerStatements.WriteString(ds_productionseries_propertytype.GongMarshallField(stage, "Show"))
		initializerStatements.WriteString(ds_productionseries_propertytype.GongMarshallField(stage, "Actuate"))
		pointersInitializesStatements.WriteString(ds_productionseries_propertytype.GongMarshallField(stage, "DS_ProductionSeries"))
	}
	initRes = initializerStatements.String()
	ptrRes = pointersInitializesStatements.String()
	return
}
func (ds_productionseries_type *DS_ProductionSeries_Type) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(ds_productionseries_type.GongMarshallField(stage, "Name"))
		initializerStatements.WriteString(ds_productionseries_type.GongMarshallField(stage, "Id"))
		initializerStatements.WriteString(ds_productionseries_type.GongMarshallField(stage, "Uuid"))
		pointersInitializesStatements.WriteString(ds_productionseries_type.GongMarshallField(stage, "ComposedOf"))
		pointersInitializesStatements.WriteString(ds_productionseries_type.GongMarshallField(stage, "SeriesMetadata"))
		pointersInitializesStatements.WriteString(ds_productionseries_type.GongMarshallField(stage, "Subset"))
		pointersInitializesStatements.WriteString(ds_productionseries_type.GongMarshallField(stage, "Superset"))
	}
	initRes = initializerStatements.String()
	ptrRes = pointersInitializesStatements.String()
	return
}
func (ds_sensor_propertytype *DS_Sensor_PropertyType) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(ds_sensor_propertytype.GongMarshallField(stage, "Name"))
		initializerStatements.WriteString(ds_sensor_propertytype.GongMarshallField(stage, "NilReason"))
		initializerStatements.WriteString(ds_sensor_propertytype.GongMarshallField(stage, "Uuidref"))
		initializerStatements.WriteString(ds_sensor_propertytype.GongMarshallField(stage, "Type_XSD"))
		initializerStatements.WriteString(ds_sensor_propertytype.GongMarshallField(stage, "Href"))
		initializerStatements.WriteString(ds_sensor_propertytype.GongMarshallField(stage, "Role"))
		initializerStatements.WriteString(ds_sensor_propertytype.GongMarshallField(stage, "Arcrole"))
		initializerStatements.WriteString(ds_sensor_propertytype.GongMarshallField(stage, "Title"))
		initializerStatements.WriteString(ds_sensor_propertytype.GongMarshallField(stage, "Show"))
		initializerStatements.WriteString(ds_sensor_propertytype.GongMarshallField(stage, "Actuate"))
		pointersInitializesStatements.WriteString(ds_sensor_propertytype.GongMarshallField(stage, "DS_Sensor"))
	}
	initRes = initializerStatements.String()
	ptrRes = pointersInitializesStatements.String()
	return
}
func (ds_sensor_type *DS_Sensor_Type) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(ds_sensor_type.GongMarshallField(stage, "Name"))
		initializerStatements.WriteString(ds_sensor_type.GongMarshallField(stage, "Id"))
		initializerStatements.WriteString(ds_sensor_type.GongMarshallField(stage, "Uuid"))
		pointersInitializesStatements.WriteString(ds_sensor_type.GongMarshallField(stage, "ComposedOf"))
		pointersInitializesStatements.WriteString(ds_sensor_type.GongMarshallField(stage, "SeriesMetadata"))
		pointersInitializesStatements.WriteString(ds_sensor_type.GongMarshallField(stage, "Subset"))
		pointersInitializesStatements.WriteString(ds_sensor_type.GongMarshallField(stage, "Superset"))
	}
	initRes = initializerStatements.String()
	ptrRes = pointersInitializesStatements.String()
	return
}
func (ds_series_propertytype *DS_Series_PropertyType) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(ds_series_propertytype.GongMarshallField(stage, "Name"))
		initializerStatements.WriteString(ds_series_propertytype.GongMarshallField(stage, "NilReason"))
		initializerStatements.WriteString(ds_series_propertytype.GongMarshallField(stage, "Uuidref"))
		initializerStatements.WriteString(ds_series_propertytype.GongMarshallField(stage, "Type_XSD"))
		initializerStatements.WriteString(ds_series_propertytype.GongMarshallField(stage, "Href"))
		initializerStatements.WriteString(ds_series_propertytype.GongMarshallField(stage, "Role"))
		initializerStatements.WriteString(ds_series_propertytype.GongMarshallField(stage, "Arcrole"))
		initializerStatements.WriteString(ds_series_propertytype.GongMarshallField(stage, "Title"))
		initializerStatements.WriteString(ds_series_propertytype.GongMarshallField(stage, "Show"))
		initializerStatements.WriteString(ds_series_propertytype.GongMarshallField(stage, "Actuate"))
		pointersInitializesStatements.WriteString(ds_series_propertytype.GongMarshallField(stage, "DS_Series"))
		pointersInitializesStatements.WriteString(ds_series_propertytype.GongMarshallField(stage, "DS_Platform"))
		pointersInitializesStatements.WriteString(ds_series_propertytype.GongMarshallField(stage, "DS_Sensor"))
		pointersInitializesStatements.WriteString(ds_series_propertytype.GongMarshallField(stage, "DS_ProductionSeries"))
	}
	initRes = initializerStatements.String()
	ptrRes = pointersInitializesStatements.String()
	return
}
func (ds_series_type *DS_Series_Type) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(ds_series_type.GongMarshallField(stage, "Name"))
		initializerStatements.WriteString(ds_series_type.GongMarshallField(stage, "Id"))
		initializerStatements.WriteString(ds_series_type.GongMarshallField(stage, "Uuid"))
		pointersInitializesStatements.WriteString(ds_series_type.GongMarshallField(stage, "ComposedOf"))
		pointersInitializesStatements.WriteString(ds_series_type.GongMarshallField(stage, "SeriesMetadata"))
		pointersInitializesStatements.WriteString(ds_series_type.GongMarshallField(stage, "Subset"))
		pointersInitializesStatements.WriteString(ds_series_type.GongMarshallField(stage, "Superset"))
	}
	initRes = initializerStatements.String()
	ptrRes = pointersInitializesStatements.String()
	return
}
func (ds_stereomate_propertytype *DS_StereoMate_PropertyType) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(ds_stereomate_propertytype.GongMarshallField(stage, "Name"))
		initializerStatements.WriteString(ds_stereomate_propertytype.GongMarshallField(stage, "NilReason"))
		initializerStatements.WriteString(ds_stereomate_propertytype.GongMarshallField(stage, "Uuidref"))
		initializerStatements.WriteString(ds_stereomate_propertytype.GongMarshallField(stage, "Type_XSD"))
		initializerStatements.WriteString(ds_stereomate_propertytype.GongMarshallField(stage, "Href"))
		initializerStatements.WriteString(ds_stereomate_propertytype.GongMarshallField(stage, "Role"))
		initializerStatements.WriteString(ds_stereomate_propertytype.GongMarshallField(stage, "Arcrole"))
		initializerStatements.WriteString(ds_stereomate_propertytype.GongMarshallField(stage, "Title"))
		initializerStatements.WriteString(ds_stereomate_propertytype.GongMarshallField(stage, "Show"))
		initializerStatements.WriteString(ds_stereomate_propertytype.GongMarshallField(stage, "Actuate"))
		pointersInitializesStatements.WriteString(ds_stereomate_propertytype.GongMarshallField(stage, "DS_StereoMate"))
	}
	initRes = initializerStatements.String()
	ptrRes = pointersInitializesStatements.String()
	return
}
func (ds_stereomate_type *DS_StereoMate_Type) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(ds_stereomate_type.GongMarshallField(stage, "Name"))
		initializerStatements.WriteString(ds_stereomate_type.GongMarshallField(stage, "Id"))
		initializerStatements.WriteString(ds_stereomate_type.GongMarshallField(stage, "Uuid"))
		pointersInitializesStatements.WriteString(ds_stereomate_type.GongMarshallField(stage, "ComposedOf"))
		pointersInitializesStatements.WriteString(ds_stereomate_type.GongMarshallField(stage, "SeriesMetadata"))
		pointersInitializesStatements.WriteString(ds_stereomate_type.GongMarshallField(stage, "Subset"))
		pointersInitializesStatements.WriteString(ds_stereomate_type.GongMarshallField(stage, "Superset"))
	}
	initRes = initializerStatements.String()
	ptrRes = pointersInitializesStatements.String()
	return
}
func (md_metadata_propertytype *MD_Metadata_PropertyType) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(md_metadata_propertytype.GongMarshallField(stage, "Name"))
		initializerStatements.WriteString(md_metadata_propertytype.GongMarshallField(stage, "NilReason"))
		initializerStatements.WriteString(md_metadata_propertytype.GongMarshallField(stage, "Uuidref"))
		initializerStatements.WriteString(md_metadata_propertytype.GongMarshallField(stage, "Type_XSD"))
		initializerStatements.WriteString(md_metadata_propertytype.GongMarshallField(stage, "Href"))
		initializerStatements.WriteString(md_metadata_propertytype.GongMarshallField(stage, "Role"))
		initializerStatements.WriteString(md_metadata_propertytype.GongMarshallField(stage, "Arcrole"))
		initializerStatements.WriteString(md_metadata_propertytype.GongMarshallField(stage, "Title"))
		initializerStatements.WriteString(md_metadata_propertytype.GongMarshallField(stage, "Show"))
		initializerStatements.WriteString(md_metadata_propertytype.GongMarshallField(stage, "Actuate"))
		pointersInitializesStatements.WriteString(md_metadata_propertytype.GongMarshallField(stage, "MD_Metadata"))
	}
	initRes = initializerStatements.String()
	ptrRes = pointersInitializesStatements.String()
	return
}
func (md_metadata_type *MD_Metadata_Type) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(md_metadata_type.GongMarshallField(stage, "Name"))
		initializerStatements.WriteString(md_metadata_type.GongMarshallField(stage, "Id"))
		initializerStatements.WriteString(md_metadata_type.GongMarshallField(stage, "Uuid"))
		pointersInitializesStatements.WriteString(md_metadata_type.GongMarshallField(stage, "Series"))
		pointersInitializesStatements.WriteString(md_metadata_type.GongMarshallField(stage, "Describes"))
	}
	initRes = initializerStatements.String()
	ptrRes = pointersInitializesStatements.String()
	return
}
