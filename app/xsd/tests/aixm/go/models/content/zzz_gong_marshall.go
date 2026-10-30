// generated code - do not edit
package content

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
	{{Identifier}} := (&content.{{GeneratedStructName}}{Name: {{GeneratedFieldNameValue}}}).Stage(stage)`

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
			const fnSig = "func _(stage *content.Stage) {"
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
			strings.HasSuffix(trimmedContentBeforeBrace, "func _(stage *content.Stage) {") ||
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
	modelsPackageName := "github.com/fullstack-lang/gong/app/xsd/tests/aixm/go/models/content"
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
		for _, abstractmd_contentinformation_type := range __gong__sortStageSetInstances(stage.AbstractMD_ContentInformation_Types, stage.AbstractMD_ContentInformation_Type_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			abstractmd_contentinformation_typeIdent := abstractmd_contentinformation_type.GongGetIdentifier(stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&content.AbstractMD_ContentInformation_Type{Name: %s}).Stage(stage)", abstractmd_contentinformation_typeIdent, __gong__toRawStringLiteral(abstractmd_contentinformation_type.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", abstractmd_contentinformation_typeIdent, __gong__toRawStringLiteral(abstractmd_contentinformation_type.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Id = %s", abstractmd_contentinformation_typeIdent, __gong__toRawStringLiteral(abstractmd_contentinformation_type.Id)))
			values.WriteString(fmt.Sprintf("\n\t%s.Uuid = %s", abstractmd_contentinformation_typeIdent, __gong__toRawStringLiteral(abstractmd_contentinformation_type.Uuid)))
		}
	}
	if stage != nil {
		for _, md_band_propertytype := range __gong__sortStageSetInstances(stage.MD_Band_PropertyTypes, stage.MD_Band_PropertyType_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			md_band_propertytypeIdent := md_band_propertytype.GongGetIdentifier(stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&content.MD_Band_PropertyType{Name: %s}).Stage(stage)", md_band_propertytypeIdent, __gong__toRawStringLiteral(md_band_propertytype.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", md_band_propertytypeIdent, __gong__toRawStringLiteral(md_band_propertytype.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.NilReason = %s", md_band_propertytypeIdent, __gong__toRawStringLiteral(md_band_propertytype.NilReason)))
			values.WriteString(fmt.Sprintf("\n\t%s.Uuidref = %s", md_band_propertytypeIdent, __gong__toRawStringLiteral(md_band_propertytype.Uuidref)))
			values.WriteString(fmt.Sprintf("\n\t%s.Type_XSD = %s", md_band_propertytypeIdent, __gong__toRawStringLiteral(md_band_propertytype.Type_XSD)))
			values.WriteString(fmt.Sprintf("\n\t%s.Href = %s", md_band_propertytypeIdent, __gong__toRawStringLiteral(md_band_propertytype.Href)))
			values.WriteString(fmt.Sprintf("\n\t%s.Role = %s", md_band_propertytypeIdent, __gong__toRawStringLiteral(md_band_propertytype.Role)))
			values.WriteString(fmt.Sprintf("\n\t%s.Arcrole = %s", md_band_propertytypeIdent, __gong__toRawStringLiteral(md_band_propertytype.Arcrole)))
			values.WriteString(fmt.Sprintf("\n\t%s.Title = %s", md_band_propertytypeIdent, __gong__toRawStringLiteral(md_band_propertytype.Title)))
			values.WriteString(fmt.Sprintf("\n\t%s.Show = %s", md_band_propertytypeIdent, __gong__toRawStringLiteral(md_band_propertytype.Show)))
			values.WriteString(fmt.Sprintf("\n\t%s.Actuate = %s", md_band_propertytypeIdent, __gong__toRawStringLiteral(md_band_propertytype.Actuate)))
			if md_band_propertytype.MD_Band != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := md_band_propertytype.MD_Band.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.MD_Band = %s", md_band_propertytypeIdent, targetIdent))
			}
		}
	}
	if stage != nil {
		for _, md_band_type := range __gong__sortStageSetInstances(stage.MD_Band_Types, stage.MD_Band_Type_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			md_band_typeIdent := md_band_type.GongGetIdentifier(stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&content.MD_Band_Type{Name: %s}).Stage(stage)", md_band_typeIdent, __gong__toRawStringLiteral(md_band_type.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", md_band_typeIdent, __gong__toRawStringLiteral(md_band_type.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Id = %s", md_band_typeIdent, __gong__toRawStringLiteral(md_band_type.Id)))
			values.WriteString(fmt.Sprintf("\n\t%s.Uuid = %s", md_band_typeIdent, __gong__toRawStringLiteral(md_band_type.Uuid)))
		}
	}
	if stage != nil {
		for _, md_contentinformation_propertytype := range __gong__sortStageSetInstances(stage.MD_ContentInformation_PropertyTypes, stage.MD_ContentInformation_PropertyType_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			md_contentinformation_propertytypeIdent := md_contentinformation_propertytype.GongGetIdentifier(stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&content.MD_ContentInformation_PropertyType{Name: %s}).Stage(stage)", md_contentinformation_propertytypeIdent, __gong__toRawStringLiteral(md_contentinformation_propertytype.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", md_contentinformation_propertytypeIdent, __gong__toRawStringLiteral(md_contentinformation_propertytype.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.NilReason = %s", md_contentinformation_propertytypeIdent, __gong__toRawStringLiteral(md_contentinformation_propertytype.NilReason)))
			values.WriteString(fmt.Sprintf("\n\t%s.Uuidref = %s", md_contentinformation_propertytypeIdent, __gong__toRawStringLiteral(md_contentinformation_propertytype.Uuidref)))
			values.WriteString(fmt.Sprintf("\n\t%s.Type_XSD = %s", md_contentinformation_propertytypeIdent, __gong__toRawStringLiteral(md_contentinformation_propertytype.Type_XSD)))
			values.WriteString(fmt.Sprintf("\n\t%s.Href = %s", md_contentinformation_propertytypeIdent, __gong__toRawStringLiteral(md_contentinformation_propertytype.Href)))
			values.WriteString(fmt.Sprintf("\n\t%s.Role = %s", md_contentinformation_propertytypeIdent, __gong__toRawStringLiteral(md_contentinformation_propertytype.Role)))
			values.WriteString(fmt.Sprintf("\n\t%s.Arcrole = %s", md_contentinformation_propertytypeIdent, __gong__toRawStringLiteral(md_contentinformation_propertytype.Arcrole)))
			values.WriteString(fmt.Sprintf("\n\t%s.Title = %s", md_contentinformation_propertytypeIdent, __gong__toRawStringLiteral(md_contentinformation_propertytype.Title)))
			values.WriteString(fmt.Sprintf("\n\t%s.Show = %s", md_contentinformation_propertytypeIdent, __gong__toRawStringLiteral(md_contentinformation_propertytype.Show)))
			values.WriteString(fmt.Sprintf("\n\t%s.Actuate = %s", md_contentinformation_propertytypeIdent, __gong__toRawStringLiteral(md_contentinformation_propertytype.Actuate)))
			if md_contentinformation_propertytype.MD_FeatureCatalogueDescription != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := md_contentinformation_propertytype.MD_FeatureCatalogueDescription.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.MD_FeatureCatalogueDescription = %s", md_contentinformation_propertytypeIdent, targetIdent))
			}
			if md_contentinformation_propertytype.MD_CoverageDescription != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := md_contentinformation_propertytype.MD_CoverageDescription.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.MD_CoverageDescription = %s", md_contentinformation_propertytypeIdent, targetIdent))
			}
			if md_contentinformation_propertytype.MD_ImageDescription != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := md_contentinformation_propertytype.MD_ImageDescription.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.MD_ImageDescription = %s", md_contentinformation_propertytypeIdent, targetIdent))
			}
		}
	}
	if stage != nil {
		for _, md_coveragecontenttypecode_propertytype := range __gong__sortStageSetInstances(stage.MD_CoverageContentTypeCode_PropertyTypes, stage.MD_CoverageContentTypeCode_PropertyType_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			md_coveragecontenttypecode_propertytypeIdent := md_coveragecontenttypecode_propertytype.GongGetIdentifier(stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&content.MD_CoverageContentTypeCode_PropertyType{Name: %s}).Stage(stage)", md_coveragecontenttypecode_propertytypeIdent, __gong__toRawStringLiteral(md_coveragecontenttypecode_propertytype.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", md_coveragecontenttypecode_propertytypeIdent, __gong__toRawStringLiteral(md_coveragecontenttypecode_propertytype.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.NilReason = %s", md_coveragecontenttypecode_propertytypeIdent, __gong__toRawStringLiteral(md_coveragecontenttypecode_propertytype.NilReason)))
		}
	}
	if stage != nil {
		for _, md_coveragedescription_propertytype := range __gong__sortStageSetInstances(stage.MD_CoverageDescription_PropertyTypes, stage.MD_CoverageDescription_PropertyType_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			md_coveragedescription_propertytypeIdent := md_coveragedescription_propertytype.GongGetIdentifier(stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&content.MD_CoverageDescription_PropertyType{Name: %s}).Stage(stage)", md_coveragedescription_propertytypeIdent, __gong__toRawStringLiteral(md_coveragedescription_propertytype.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", md_coveragedescription_propertytypeIdent, __gong__toRawStringLiteral(md_coveragedescription_propertytype.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.NilReason = %s", md_coveragedescription_propertytypeIdent, __gong__toRawStringLiteral(md_coveragedescription_propertytype.NilReason)))
			values.WriteString(fmt.Sprintf("\n\t%s.Uuidref = %s", md_coveragedescription_propertytypeIdent, __gong__toRawStringLiteral(md_coveragedescription_propertytype.Uuidref)))
			values.WriteString(fmt.Sprintf("\n\t%s.Type_XSD = %s", md_coveragedescription_propertytypeIdent, __gong__toRawStringLiteral(md_coveragedescription_propertytype.Type_XSD)))
			values.WriteString(fmt.Sprintf("\n\t%s.Href = %s", md_coveragedescription_propertytypeIdent, __gong__toRawStringLiteral(md_coveragedescription_propertytype.Href)))
			values.WriteString(fmt.Sprintf("\n\t%s.Role = %s", md_coveragedescription_propertytypeIdent, __gong__toRawStringLiteral(md_coveragedescription_propertytype.Role)))
			values.WriteString(fmt.Sprintf("\n\t%s.Arcrole = %s", md_coveragedescription_propertytypeIdent, __gong__toRawStringLiteral(md_coveragedescription_propertytype.Arcrole)))
			values.WriteString(fmt.Sprintf("\n\t%s.Title = %s", md_coveragedescription_propertytypeIdent, __gong__toRawStringLiteral(md_coveragedescription_propertytype.Title)))
			values.WriteString(fmt.Sprintf("\n\t%s.Show = %s", md_coveragedescription_propertytypeIdent, __gong__toRawStringLiteral(md_coveragedescription_propertytype.Show)))
			values.WriteString(fmt.Sprintf("\n\t%s.Actuate = %s", md_coveragedescription_propertytypeIdent, __gong__toRawStringLiteral(md_coveragedescription_propertytype.Actuate)))
			if md_coveragedescription_propertytype.MD_CoverageDescription != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := md_coveragedescription_propertytype.MD_CoverageDescription.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.MD_CoverageDescription = %s", md_coveragedescription_propertytypeIdent, targetIdent))
			}
			if md_coveragedescription_propertytype.MD_ImageDescription != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := md_coveragedescription_propertytype.MD_ImageDescription.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.MD_ImageDescription = %s", md_coveragedescription_propertytypeIdent, targetIdent))
			}
		}
	}
	if stage != nil {
		for _, md_coveragedescription_type := range __gong__sortStageSetInstances(stage.MD_CoverageDescription_Types, stage.MD_CoverageDescription_Type_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			md_coveragedescription_typeIdent := md_coveragedescription_type.GongGetIdentifier(stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&content.MD_CoverageDescription_Type{Name: %s}).Stage(stage)", md_coveragedescription_typeIdent, __gong__toRawStringLiteral(md_coveragedescription_type.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", md_coveragedescription_typeIdent, __gong__toRawStringLiteral(md_coveragedescription_type.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Id = %s", md_coveragedescription_typeIdent, __gong__toRawStringLiteral(md_coveragedescription_type.Id)))
			values.WriteString(fmt.Sprintf("\n\t%s.Uuid = %s", md_coveragedescription_typeIdent, __gong__toRawStringLiteral(md_coveragedescription_type.Uuid)))
			if md_coveragedescription_type.ContentType != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := md_coveragedescription_type.ContentType.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.ContentType = %s", md_coveragedescription_typeIdent, targetIdent))
			}
			for _, elem := range md_coveragedescription_type.Dimension {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := elem.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Dimension = append(%s.Dimension, %s)", md_coveragedescription_typeIdent, md_coveragedescription_typeIdent, targetIdent))
			}
		}
	}
	if stage != nil {
		for _, md_featurecataloguedescription_propertytype := range __gong__sortStageSetInstances(stage.MD_FeatureCatalogueDescription_PropertyTypes, stage.MD_FeatureCatalogueDescription_PropertyType_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			md_featurecataloguedescription_propertytypeIdent := md_featurecataloguedescription_propertytype.GongGetIdentifier(stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&content.MD_FeatureCatalogueDescription_PropertyType{Name: %s}).Stage(stage)", md_featurecataloguedescription_propertytypeIdent, __gong__toRawStringLiteral(md_featurecataloguedescription_propertytype.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", md_featurecataloguedescription_propertytypeIdent, __gong__toRawStringLiteral(md_featurecataloguedescription_propertytype.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.NilReason = %s", md_featurecataloguedescription_propertytypeIdent, __gong__toRawStringLiteral(md_featurecataloguedescription_propertytype.NilReason)))
			values.WriteString(fmt.Sprintf("\n\t%s.Uuidref = %s", md_featurecataloguedescription_propertytypeIdent, __gong__toRawStringLiteral(md_featurecataloguedescription_propertytype.Uuidref)))
			values.WriteString(fmt.Sprintf("\n\t%s.Type_XSD = %s", md_featurecataloguedescription_propertytypeIdent, __gong__toRawStringLiteral(md_featurecataloguedescription_propertytype.Type_XSD)))
			values.WriteString(fmt.Sprintf("\n\t%s.Href = %s", md_featurecataloguedescription_propertytypeIdent, __gong__toRawStringLiteral(md_featurecataloguedescription_propertytype.Href)))
			values.WriteString(fmt.Sprintf("\n\t%s.Role = %s", md_featurecataloguedescription_propertytypeIdent, __gong__toRawStringLiteral(md_featurecataloguedescription_propertytype.Role)))
			values.WriteString(fmt.Sprintf("\n\t%s.Arcrole = %s", md_featurecataloguedescription_propertytypeIdent, __gong__toRawStringLiteral(md_featurecataloguedescription_propertytype.Arcrole)))
			values.WriteString(fmt.Sprintf("\n\t%s.Title = %s", md_featurecataloguedescription_propertytypeIdent, __gong__toRawStringLiteral(md_featurecataloguedescription_propertytype.Title)))
			values.WriteString(fmt.Sprintf("\n\t%s.Show = %s", md_featurecataloguedescription_propertytypeIdent, __gong__toRawStringLiteral(md_featurecataloguedescription_propertytype.Show)))
			values.WriteString(fmt.Sprintf("\n\t%s.Actuate = %s", md_featurecataloguedescription_propertytypeIdent, __gong__toRawStringLiteral(md_featurecataloguedescription_propertytype.Actuate)))
			if md_featurecataloguedescription_propertytype.MD_FeatureCatalogueDescription != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := md_featurecataloguedescription_propertytype.MD_FeatureCatalogueDescription.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.MD_FeatureCatalogueDescription = %s", md_featurecataloguedescription_propertytypeIdent, targetIdent))
			}
		}
	}
	if stage != nil {
		for _, md_featurecataloguedescription_type := range __gong__sortStageSetInstances(stage.MD_FeatureCatalogueDescription_Types, stage.MD_FeatureCatalogueDescription_Type_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			md_featurecataloguedescription_typeIdent := md_featurecataloguedescription_type.GongGetIdentifier(stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&content.MD_FeatureCatalogueDescription_Type{Name: %s}).Stage(stage)", md_featurecataloguedescription_typeIdent, __gong__toRawStringLiteral(md_featurecataloguedescription_type.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", md_featurecataloguedescription_typeIdent, __gong__toRawStringLiteral(md_featurecataloguedescription_type.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Id = %s", md_featurecataloguedescription_typeIdent, __gong__toRawStringLiteral(md_featurecataloguedescription_type.Id)))
			values.WriteString(fmt.Sprintf("\n\t%s.Uuid = %s", md_featurecataloguedescription_typeIdent, __gong__toRawStringLiteral(md_featurecataloguedescription_type.Uuid)))
		}
	}
	if stage != nil {
		for _, md_imagedescription_propertytype := range __gong__sortStageSetInstances(stage.MD_ImageDescription_PropertyTypes, stage.MD_ImageDescription_PropertyType_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			md_imagedescription_propertytypeIdent := md_imagedescription_propertytype.GongGetIdentifier(stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&content.MD_ImageDescription_PropertyType{Name: %s}).Stage(stage)", md_imagedescription_propertytypeIdent, __gong__toRawStringLiteral(md_imagedescription_propertytype.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", md_imagedescription_propertytypeIdent, __gong__toRawStringLiteral(md_imagedescription_propertytype.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.NilReason = %s", md_imagedescription_propertytypeIdent, __gong__toRawStringLiteral(md_imagedescription_propertytype.NilReason)))
			values.WriteString(fmt.Sprintf("\n\t%s.Uuidref = %s", md_imagedescription_propertytypeIdent, __gong__toRawStringLiteral(md_imagedescription_propertytype.Uuidref)))
			values.WriteString(fmt.Sprintf("\n\t%s.Type_XSD = %s", md_imagedescription_propertytypeIdent, __gong__toRawStringLiteral(md_imagedescription_propertytype.Type_XSD)))
			values.WriteString(fmt.Sprintf("\n\t%s.Href = %s", md_imagedescription_propertytypeIdent, __gong__toRawStringLiteral(md_imagedescription_propertytype.Href)))
			values.WriteString(fmt.Sprintf("\n\t%s.Role = %s", md_imagedescription_propertytypeIdent, __gong__toRawStringLiteral(md_imagedescription_propertytype.Role)))
			values.WriteString(fmt.Sprintf("\n\t%s.Arcrole = %s", md_imagedescription_propertytypeIdent, __gong__toRawStringLiteral(md_imagedescription_propertytype.Arcrole)))
			values.WriteString(fmt.Sprintf("\n\t%s.Title = %s", md_imagedescription_propertytypeIdent, __gong__toRawStringLiteral(md_imagedescription_propertytype.Title)))
			values.WriteString(fmt.Sprintf("\n\t%s.Show = %s", md_imagedescription_propertytypeIdent, __gong__toRawStringLiteral(md_imagedescription_propertytype.Show)))
			values.WriteString(fmt.Sprintf("\n\t%s.Actuate = %s", md_imagedescription_propertytypeIdent, __gong__toRawStringLiteral(md_imagedescription_propertytype.Actuate)))
			if md_imagedescription_propertytype.MD_ImageDescription != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := md_imagedescription_propertytype.MD_ImageDescription.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.MD_ImageDescription = %s", md_imagedescription_propertytypeIdent, targetIdent))
			}
		}
	}
	if stage != nil {
		for _, md_imagedescription_type := range __gong__sortStageSetInstances(stage.MD_ImageDescription_Types, stage.MD_ImageDescription_Type_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			md_imagedescription_typeIdent := md_imagedescription_type.GongGetIdentifier(stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&content.MD_ImageDescription_Type{Name: %s}).Stage(stage)", md_imagedescription_typeIdent, __gong__toRawStringLiteral(md_imagedescription_type.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", md_imagedescription_typeIdent, __gong__toRawStringLiteral(md_imagedescription_type.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Id = %s", md_imagedescription_typeIdent, __gong__toRawStringLiteral(md_imagedescription_type.Id)))
			values.WriteString(fmt.Sprintf("\n\t%s.Uuid = %s", md_imagedescription_typeIdent, __gong__toRawStringLiteral(md_imagedescription_type.Uuid)))
			if md_imagedescription_type.ContentType != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := md_imagedescription_type.ContentType.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.ContentType = %s", md_imagedescription_typeIdent, targetIdent))
			}
			for _, elem := range md_imagedescription_type.Dimension {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := elem.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Dimension = append(%s.Dimension, %s)", md_imagedescription_typeIdent, md_imagedescription_typeIdent, targetIdent))
			}
			if md_imagedescription_type.ImagingCondition != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := md_imagedescription_type.ImagingCondition.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.ImagingCondition = %s", md_imagedescription_typeIdent, targetIdent))
			}
		}
	}
	if stage != nil {
		for _, md_imagingconditioncode_propertytype := range __gong__sortStageSetInstances(stage.MD_ImagingConditionCode_PropertyTypes, stage.MD_ImagingConditionCode_PropertyType_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			md_imagingconditioncode_propertytypeIdent := md_imagingconditioncode_propertytype.GongGetIdentifier(stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&content.MD_ImagingConditionCode_PropertyType{Name: %s}).Stage(stage)", md_imagingconditioncode_propertytypeIdent, __gong__toRawStringLiteral(md_imagingconditioncode_propertytype.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", md_imagingconditioncode_propertytypeIdent, __gong__toRawStringLiteral(md_imagingconditioncode_propertytype.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.NilReason = %s", md_imagingconditioncode_propertytypeIdent, __gong__toRawStringLiteral(md_imagingconditioncode_propertytype.NilReason)))
		}
	}
	if stage != nil {
		for _, md_rangedimension_propertytype := range __gong__sortStageSetInstances(stage.MD_RangeDimension_PropertyTypes, stage.MD_RangeDimension_PropertyType_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			md_rangedimension_propertytypeIdent := md_rangedimension_propertytype.GongGetIdentifier(stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&content.MD_RangeDimension_PropertyType{Name: %s}).Stage(stage)", md_rangedimension_propertytypeIdent, __gong__toRawStringLiteral(md_rangedimension_propertytype.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", md_rangedimension_propertytypeIdent, __gong__toRawStringLiteral(md_rangedimension_propertytype.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.NilReason = %s", md_rangedimension_propertytypeIdent, __gong__toRawStringLiteral(md_rangedimension_propertytype.NilReason)))
			values.WriteString(fmt.Sprintf("\n\t%s.Uuidref = %s", md_rangedimension_propertytypeIdent, __gong__toRawStringLiteral(md_rangedimension_propertytype.Uuidref)))
			values.WriteString(fmt.Sprintf("\n\t%s.Type_XSD = %s", md_rangedimension_propertytypeIdent, __gong__toRawStringLiteral(md_rangedimension_propertytype.Type_XSD)))
			values.WriteString(fmt.Sprintf("\n\t%s.Href = %s", md_rangedimension_propertytypeIdent, __gong__toRawStringLiteral(md_rangedimension_propertytype.Href)))
			values.WriteString(fmt.Sprintf("\n\t%s.Role = %s", md_rangedimension_propertytypeIdent, __gong__toRawStringLiteral(md_rangedimension_propertytype.Role)))
			values.WriteString(fmt.Sprintf("\n\t%s.Arcrole = %s", md_rangedimension_propertytypeIdent, __gong__toRawStringLiteral(md_rangedimension_propertytype.Arcrole)))
			values.WriteString(fmt.Sprintf("\n\t%s.Title = %s", md_rangedimension_propertytypeIdent, __gong__toRawStringLiteral(md_rangedimension_propertytype.Title)))
			values.WriteString(fmt.Sprintf("\n\t%s.Show = %s", md_rangedimension_propertytypeIdent, __gong__toRawStringLiteral(md_rangedimension_propertytype.Show)))
			values.WriteString(fmt.Sprintf("\n\t%s.Actuate = %s", md_rangedimension_propertytypeIdent, __gong__toRawStringLiteral(md_rangedimension_propertytype.Actuate)))
			if md_rangedimension_propertytype.MD_RangeDimension != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := md_rangedimension_propertytype.MD_RangeDimension.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.MD_RangeDimension = %s", md_rangedimension_propertytypeIdent, targetIdent))
			}
			if md_rangedimension_propertytype.MD_Band != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := md_rangedimension_propertytype.MD_Band.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.MD_Band = %s", md_rangedimension_propertytypeIdent, targetIdent))
			}
		}
	}
	if stage != nil {
		for _, md_rangedimension_type := range __gong__sortStageSetInstances(stage.MD_RangeDimension_Types, stage.MD_RangeDimension_Type_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			md_rangedimension_typeIdent := md_rangedimension_type.GongGetIdentifier(stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&content.MD_RangeDimension_Type{Name: %s}).Stage(stage)", md_rangedimension_typeIdent, __gong__toRawStringLiteral(md_rangedimension_type.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", md_rangedimension_typeIdent, __gong__toRawStringLiteral(md_rangedimension_type.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Id = %s", md_rangedimension_typeIdent, __gong__toRawStringLiteral(md_rangedimension_type.Id)))
			values.WriteString(fmt.Sprintf("\n\t%s.Uuid = %s", md_rangedimension_typeIdent, __gong__toRawStringLiteral(md_rangedimension_type.Uuid)))
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
go run github.com/fullstack-lang/gong/app/xsd/tests/aixm/go/models/content/probe/cmd/probe@%s stage.go
*/
package %s

import (
	"slices"
	"time"

	"github.com/fullstack-lang/gong/app/xsd/tests/aixm/go/models/content"%s
)

var (
	_ time.Time
	_ = slices.Index[[]int, int]

	_ *content.Stage
)%s

// function will stage objects across all coordinated stages
func _(stage *content.Stage) {

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

	if modelsPackageName != "" && modelsPackageName != "github.com/fullstack-lang/gong/app/xsd/tests/aixm/go/models/content" {
		res = strings.ReplaceAll(res, "\"github.com/fullstack-lang/gong/app/xsd/tests/aixm/go/models/content\"", fmt.Sprintf("%q", modelsPackageName))
	}

	return res, nil
}

// insertion point for marshall field methods
func (abstractmd_contentinformation_type *AbstractMD_ContentInformation_Type) GongMarshallField(stage *Stage, fieldName string) (res string) {
	ident := abstractmd_contentinformation_type.GongGetIdentifier(stage)
	_ = ident
	switch fieldName {
	case "Name":
		res = __gong__marshallString(ident, "Name", abstractmd_contentinformation_type.Name)
	case "Id":
		res = __gong__marshallString(ident, "Id", abstractmd_contentinformation_type.Id)
	case "Uuid":
		res = __gong__marshallString(ident, "Uuid", abstractmd_contentinformation_type.Uuid)

	default:
		log.Panicf("Unknown field %s for Gongstruct AbstractMD_ContentInformation_Type", fieldName)
	}
	return
}

func (md_band_propertytype *MD_Band_PropertyType) GongMarshallField(stage *Stage, fieldName string) (res string) {
	ident := md_band_propertytype.GongGetIdentifier(stage)
	_ = ident
	switch fieldName {
	case "Name":
		res = __gong__marshallString(ident, "Name", md_band_propertytype.Name)
	case "NilReason":
		res = __gong__marshallString(ident, "NilReason", md_band_propertytype.NilReason)
	case "Uuidref":
		res = __gong__marshallString(ident, "Uuidref", md_band_propertytype.Uuidref)
	case "Type_XSD":
		res = __gong__marshallString(ident, "Type_XSD", md_band_propertytype.Type_XSD)
	case "Href":
		res = __gong__marshallString(ident, "Href", md_band_propertytype.Href)
	case "Role":
		res = __gong__marshallString(ident, "Role", md_band_propertytype.Role)
	case "Arcrole":
		res = __gong__marshallString(ident, "Arcrole", md_band_propertytype.Arcrole)
	case "Title":
		res = __gong__marshallString(ident, "Title", md_band_propertytype.Title)
	case "Show":
		res = __gong__marshallString(ident, "Show", md_band_propertytype.Show)
	case "Actuate":
		res = __gong__marshallString(ident, "Actuate", md_band_propertytype.Actuate)

	case "MD_Band":
		if md_band_propertytype.MD_Band != nil {
			res = __gong__marshallPointer(ident, "MD_Band", md_band_propertytype.MD_Band.GongGetIdentifier(stage))
		} else {
			res = __gong__marshallPointer(ident, "MD_Band", "nil")
		}
	default:
		log.Panicf("Unknown field %s for Gongstruct MD_Band_PropertyType", fieldName)
	}
	return
}

func (md_band_type *MD_Band_Type) GongMarshallField(stage *Stage, fieldName string) (res string) {
	ident := md_band_type.GongGetIdentifier(stage)
	_ = ident
	switch fieldName {
	case "Name":
		res = __gong__marshallString(ident, "Name", md_band_type.Name)
	case "Id":
		res = __gong__marshallString(ident, "Id", md_band_type.Id)
	case "Uuid":
		res = __gong__marshallString(ident, "Uuid", md_band_type.Uuid)

	default:
		log.Panicf("Unknown field %s for Gongstruct MD_Band_Type", fieldName)
	}
	return
}

func (md_contentinformation_propertytype *MD_ContentInformation_PropertyType) GongMarshallField(stage *Stage, fieldName string) (res string) {
	ident := md_contentinformation_propertytype.GongGetIdentifier(stage)
	_ = ident
	switch fieldName {
	case "Name":
		res = __gong__marshallString(ident, "Name", md_contentinformation_propertytype.Name)
	case "NilReason":
		res = __gong__marshallString(ident, "NilReason", md_contentinformation_propertytype.NilReason)
	case "Uuidref":
		res = __gong__marshallString(ident, "Uuidref", md_contentinformation_propertytype.Uuidref)
	case "Type_XSD":
		res = __gong__marshallString(ident, "Type_XSD", md_contentinformation_propertytype.Type_XSD)
	case "Href":
		res = __gong__marshallString(ident, "Href", md_contentinformation_propertytype.Href)
	case "Role":
		res = __gong__marshallString(ident, "Role", md_contentinformation_propertytype.Role)
	case "Arcrole":
		res = __gong__marshallString(ident, "Arcrole", md_contentinformation_propertytype.Arcrole)
	case "Title":
		res = __gong__marshallString(ident, "Title", md_contentinformation_propertytype.Title)
	case "Show":
		res = __gong__marshallString(ident, "Show", md_contentinformation_propertytype.Show)
	case "Actuate":
		res = __gong__marshallString(ident, "Actuate", md_contentinformation_propertytype.Actuate)

	case "MD_FeatureCatalogueDescription":
		if md_contentinformation_propertytype.MD_FeatureCatalogueDescription != nil {
			res = __gong__marshallPointer(ident, "MD_FeatureCatalogueDescription", md_contentinformation_propertytype.MD_FeatureCatalogueDescription.GongGetIdentifier(stage))
		} else {
			res = __gong__marshallPointer(ident, "MD_FeatureCatalogueDescription", "nil")
		}
	case "MD_CoverageDescription":
		if md_contentinformation_propertytype.MD_CoverageDescription != nil {
			res = __gong__marshallPointer(ident, "MD_CoverageDescription", md_contentinformation_propertytype.MD_CoverageDescription.GongGetIdentifier(stage))
		} else {
			res = __gong__marshallPointer(ident, "MD_CoverageDescription", "nil")
		}
	case "MD_ImageDescription":
		if md_contentinformation_propertytype.MD_ImageDescription != nil {
			res = __gong__marshallPointer(ident, "MD_ImageDescription", md_contentinformation_propertytype.MD_ImageDescription.GongGetIdentifier(stage))
		} else {
			res = __gong__marshallPointer(ident, "MD_ImageDescription", "nil")
		}
	default:
		log.Panicf("Unknown field %s for Gongstruct MD_ContentInformation_PropertyType", fieldName)
	}
	return
}

func (md_coveragecontenttypecode_propertytype *MD_CoverageContentTypeCode_PropertyType) GongMarshallField(stage *Stage, fieldName string) (res string) {
	ident := md_coveragecontenttypecode_propertytype.GongGetIdentifier(stage)
	_ = ident
	switch fieldName {
	case "Name":
		res = __gong__marshallString(ident, "Name", md_coveragecontenttypecode_propertytype.Name)
	case "NilReason":
		res = __gong__marshallString(ident, "NilReason", md_coveragecontenttypecode_propertytype.NilReason)

	default:
		log.Panicf("Unknown field %s for Gongstruct MD_CoverageContentTypeCode_PropertyType", fieldName)
	}
	return
}

func (md_coveragedescription_propertytype *MD_CoverageDescription_PropertyType) GongMarshallField(stage *Stage, fieldName string) (res string) {
	ident := md_coveragedescription_propertytype.GongGetIdentifier(stage)
	_ = ident
	switch fieldName {
	case "Name":
		res = __gong__marshallString(ident, "Name", md_coveragedescription_propertytype.Name)
	case "NilReason":
		res = __gong__marshallString(ident, "NilReason", md_coveragedescription_propertytype.NilReason)
	case "Uuidref":
		res = __gong__marshallString(ident, "Uuidref", md_coveragedescription_propertytype.Uuidref)
	case "Type_XSD":
		res = __gong__marshallString(ident, "Type_XSD", md_coveragedescription_propertytype.Type_XSD)
	case "Href":
		res = __gong__marshallString(ident, "Href", md_coveragedescription_propertytype.Href)
	case "Role":
		res = __gong__marshallString(ident, "Role", md_coveragedescription_propertytype.Role)
	case "Arcrole":
		res = __gong__marshallString(ident, "Arcrole", md_coveragedescription_propertytype.Arcrole)
	case "Title":
		res = __gong__marshallString(ident, "Title", md_coveragedescription_propertytype.Title)
	case "Show":
		res = __gong__marshallString(ident, "Show", md_coveragedescription_propertytype.Show)
	case "Actuate":
		res = __gong__marshallString(ident, "Actuate", md_coveragedescription_propertytype.Actuate)

	case "MD_CoverageDescription":
		if md_coveragedescription_propertytype.MD_CoverageDescription != nil {
			res = __gong__marshallPointer(ident, "MD_CoverageDescription", md_coveragedescription_propertytype.MD_CoverageDescription.GongGetIdentifier(stage))
		} else {
			res = __gong__marshallPointer(ident, "MD_CoverageDescription", "nil")
		}
	case "MD_ImageDescription":
		if md_coveragedescription_propertytype.MD_ImageDescription != nil {
			res = __gong__marshallPointer(ident, "MD_ImageDescription", md_coveragedescription_propertytype.MD_ImageDescription.GongGetIdentifier(stage))
		} else {
			res = __gong__marshallPointer(ident, "MD_ImageDescription", "nil")
		}
	default:
		log.Panicf("Unknown field %s for Gongstruct MD_CoverageDescription_PropertyType", fieldName)
	}
	return
}

func (md_coveragedescription_type *MD_CoverageDescription_Type) GongMarshallField(stage *Stage, fieldName string) (res string) {
	ident := md_coveragedescription_type.GongGetIdentifier(stage)
	_ = ident
	switch fieldName {
	case "Name":
		res = __gong__marshallString(ident, "Name", md_coveragedescription_type.Name)
	case "Id":
		res = __gong__marshallString(ident, "Id", md_coveragedescription_type.Id)
	case "Uuid":
		res = __gong__marshallString(ident, "Uuid", md_coveragedescription_type.Uuid)

	case "ContentType":
		if md_coveragedescription_type.ContentType != nil {
			res = __gong__marshallPointer(ident, "ContentType", md_coveragedescription_type.ContentType.GongGetIdentifier(stage))
		} else {
			res = __gong__marshallPointer(ident, "ContentType", "nil")
		}
	case "Dimension":
		var sb strings.Builder
		for _, _md_rangedimension_propertytype := range md_coveragedescription_type.Dimension {
			sb.WriteString(__gong__marshallSliceOfPointers(ident, "Dimension", _md_rangedimension_propertytype.GongGetIdentifier(stage)))
		}
		res = sb.String()
	default:
		log.Panicf("Unknown field %s for Gongstruct MD_CoverageDescription_Type", fieldName)
	}
	return
}

func (md_featurecataloguedescription_propertytype *MD_FeatureCatalogueDescription_PropertyType) GongMarshallField(stage *Stage, fieldName string) (res string) {
	ident := md_featurecataloguedescription_propertytype.GongGetIdentifier(stage)
	_ = ident
	switch fieldName {
	case "Name":
		res = __gong__marshallString(ident, "Name", md_featurecataloguedescription_propertytype.Name)
	case "NilReason":
		res = __gong__marshallString(ident, "NilReason", md_featurecataloguedescription_propertytype.NilReason)
	case "Uuidref":
		res = __gong__marshallString(ident, "Uuidref", md_featurecataloguedescription_propertytype.Uuidref)
	case "Type_XSD":
		res = __gong__marshallString(ident, "Type_XSD", md_featurecataloguedescription_propertytype.Type_XSD)
	case "Href":
		res = __gong__marshallString(ident, "Href", md_featurecataloguedescription_propertytype.Href)
	case "Role":
		res = __gong__marshallString(ident, "Role", md_featurecataloguedescription_propertytype.Role)
	case "Arcrole":
		res = __gong__marshallString(ident, "Arcrole", md_featurecataloguedescription_propertytype.Arcrole)
	case "Title":
		res = __gong__marshallString(ident, "Title", md_featurecataloguedescription_propertytype.Title)
	case "Show":
		res = __gong__marshallString(ident, "Show", md_featurecataloguedescription_propertytype.Show)
	case "Actuate":
		res = __gong__marshallString(ident, "Actuate", md_featurecataloguedescription_propertytype.Actuate)

	case "MD_FeatureCatalogueDescription":
		if md_featurecataloguedescription_propertytype.MD_FeatureCatalogueDescription != nil {
			res = __gong__marshallPointer(ident, "MD_FeatureCatalogueDescription", md_featurecataloguedescription_propertytype.MD_FeatureCatalogueDescription.GongGetIdentifier(stage))
		} else {
			res = __gong__marshallPointer(ident, "MD_FeatureCatalogueDescription", "nil")
		}
	default:
		log.Panicf("Unknown field %s for Gongstruct MD_FeatureCatalogueDescription_PropertyType", fieldName)
	}
	return
}

func (md_featurecataloguedescription_type *MD_FeatureCatalogueDescription_Type) GongMarshallField(stage *Stage, fieldName string) (res string) {
	ident := md_featurecataloguedescription_type.GongGetIdentifier(stage)
	_ = ident
	switch fieldName {
	case "Name":
		res = __gong__marshallString(ident, "Name", md_featurecataloguedescription_type.Name)
	case "Id":
		res = __gong__marshallString(ident, "Id", md_featurecataloguedescription_type.Id)
	case "Uuid":
		res = __gong__marshallString(ident, "Uuid", md_featurecataloguedescription_type.Uuid)

	default:
		log.Panicf("Unknown field %s for Gongstruct MD_FeatureCatalogueDescription_Type", fieldName)
	}
	return
}

func (md_imagedescription_propertytype *MD_ImageDescription_PropertyType) GongMarshallField(stage *Stage, fieldName string) (res string) {
	ident := md_imagedescription_propertytype.GongGetIdentifier(stage)
	_ = ident
	switch fieldName {
	case "Name":
		res = __gong__marshallString(ident, "Name", md_imagedescription_propertytype.Name)
	case "NilReason":
		res = __gong__marshallString(ident, "NilReason", md_imagedescription_propertytype.NilReason)
	case "Uuidref":
		res = __gong__marshallString(ident, "Uuidref", md_imagedescription_propertytype.Uuidref)
	case "Type_XSD":
		res = __gong__marshallString(ident, "Type_XSD", md_imagedescription_propertytype.Type_XSD)
	case "Href":
		res = __gong__marshallString(ident, "Href", md_imagedescription_propertytype.Href)
	case "Role":
		res = __gong__marshallString(ident, "Role", md_imagedescription_propertytype.Role)
	case "Arcrole":
		res = __gong__marshallString(ident, "Arcrole", md_imagedescription_propertytype.Arcrole)
	case "Title":
		res = __gong__marshallString(ident, "Title", md_imagedescription_propertytype.Title)
	case "Show":
		res = __gong__marshallString(ident, "Show", md_imagedescription_propertytype.Show)
	case "Actuate":
		res = __gong__marshallString(ident, "Actuate", md_imagedescription_propertytype.Actuate)

	case "MD_ImageDescription":
		if md_imagedescription_propertytype.MD_ImageDescription != nil {
			res = __gong__marshallPointer(ident, "MD_ImageDescription", md_imagedescription_propertytype.MD_ImageDescription.GongGetIdentifier(stage))
		} else {
			res = __gong__marshallPointer(ident, "MD_ImageDescription", "nil")
		}
	default:
		log.Panicf("Unknown field %s for Gongstruct MD_ImageDescription_PropertyType", fieldName)
	}
	return
}

func (md_imagedescription_type *MD_ImageDescription_Type) GongMarshallField(stage *Stage, fieldName string) (res string) {
	ident := md_imagedescription_type.GongGetIdentifier(stage)
	_ = ident
	switch fieldName {
	case "Name":
		res = __gong__marshallString(ident, "Name", md_imagedescription_type.Name)
	case "Id":
		res = __gong__marshallString(ident, "Id", md_imagedescription_type.Id)
	case "Uuid":
		res = __gong__marshallString(ident, "Uuid", md_imagedescription_type.Uuid)

	case "ContentType":
		if md_imagedescription_type.ContentType != nil {
			res = __gong__marshallPointer(ident, "ContentType", md_imagedescription_type.ContentType.GongGetIdentifier(stage))
		} else {
			res = __gong__marshallPointer(ident, "ContentType", "nil")
		}
	case "Dimension":
		var sb strings.Builder
		for _, _md_rangedimension_propertytype := range md_imagedescription_type.Dimension {
			sb.WriteString(__gong__marshallSliceOfPointers(ident, "Dimension", _md_rangedimension_propertytype.GongGetIdentifier(stage)))
		}
		res = sb.String()
	case "ImagingCondition":
		if md_imagedescription_type.ImagingCondition != nil {
			res = __gong__marshallPointer(ident, "ImagingCondition", md_imagedescription_type.ImagingCondition.GongGetIdentifier(stage))
		} else {
			res = __gong__marshallPointer(ident, "ImagingCondition", "nil")
		}
	default:
		log.Panicf("Unknown field %s for Gongstruct MD_ImageDescription_Type", fieldName)
	}
	return
}

func (md_imagingconditioncode_propertytype *MD_ImagingConditionCode_PropertyType) GongMarshallField(stage *Stage, fieldName string) (res string) {
	ident := md_imagingconditioncode_propertytype.GongGetIdentifier(stage)
	_ = ident
	switch fieldName {
	case "Name":
		res = __gong__marshallString(ident, "Name", md_imagingconditioncode_propertytype.Name)
	case "NilReason":
		res = __gong__marshallString(ident, "NilReason", md_imagingconditioncode_propertytype.NilReason)

	default:
		log.Panicf("Unknown field %s for Gongstruct MD_ImagingConditionCode_PropertyType", fieldName)
	}
	return
}

func (md_rangedimension_propertytype *MD_RangeDimension_PropertyType) GongMarshallField(stage *Stage, fieldName string) (res string) {
	ident := md_rangedimension_propertytype.GongGetIdentifier(stage)
	_ = ident
	switch fieldName {
	case "Name":
		res = __gong__marshallString(ident, "Name", md_rangedimension_propertytype.Name)
	case "NilReason":
		res = __gong__marshallString(ident, "NilReason", md_rangedimension_propertytype.NilReason)
	case "Uuidref":
		res = __gong__marshallString(ident, "Uuidref", md_rangedimension_propertytype.Uuidref)
	case "Type_XSD":
		res = __gong__marshallString(ident, "Type_XSD", md_rangedimension_propertytype.Type_XSD)
	case "Href":
		res = __gong__marshallString(ident, "Href", md_rangedimension_propertytype.Href)
	case "Role":
		res = __gong__marshallString(ident, "Role", md_rangedimension_propertytype.Role)
	case "Arcrole":
		res = __gong__marshallString(ident, "Arcrole", md_rangedimension_propertytype.Arcrole)
	case "Title":
		res = __gong__marshallString(ident, "Title", md_rangedimension_propertytype.Title)
	case "Show":
		res = __gong__marshallString(ident, "Show", md_rangedimension_propertytype.Show)
	case "Actuate":
		res = __gong__marshallString(ident, "Actuate", md_rangedimension_propertytype.Actuate)

	case "MD_RangeDimension":
		if md_rangedimension_propertytype.MD_RangeDimension != nil {
			res = __gong__marshallPointer(ident, "MD_RangeDimension", md_rangedimension_propertytype.MD_RangeDimension.GongGetIdentifier(stage))
		} else {
			res = __gong__marshallPointer(ident, "MD_RangeDimension", "nil")
		}
	case "MD_Band":
		if md_rangedimension_propertytype.MD_Band != nil {
			res = __gong__marshallPointer(ident, "MD_Band", md_rangedimension_propertytype.MD_Band.GongGetIdentifier(stage))
		} else {
			res = __gong__marshallPointer(ident, "MD_Band", "nil")
		}
	default:
		log.Panicf("Unknown field %s for Gongstruct MD_RangeDimension_PropertyType", fieldName)
	}
	return
}

func (md_rangedimension_type *MD_RangeDimension_Type) GongMarshallField(stage *Stage, fieldName string) (res string) {
	ident := md_rangedimension_type.GongGetIdentifier(stage)
	_ = ident
	switch fieldName {
	case "Name":
		res = __gong__marshallString(ident, "Name", md_rangedimension_type.Name)
	case "Id":
		res = __gong__marshallString(ident, "Id", md_rangedimension_type.Id)
	case "Uuid":
		res = __gong__marshallString(ident, "Uuid", md_rangedimension_type.Uuid)

	default:
		log.Panicf("Unknown field %s for Gongstruct MD_RangeDimension_Type", fieldName)
	}
	return
}

// insertion point for marshall all fields methods
func (abstractmd_contentinformation_type *AbstractMD_ContentInformation_Type) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(abstractmd_contentinformation_type.GongMarshallField(stage, "Name"))
		initializerStatements.WriteString(abstractmd_contentinformation_type.GongMarshallField(stage, "Id"))
		initializerStatements.WriteString(abstractmd_contentinformation_type.GongMarshallField(stage, "Uuid"))
	}
	initRes = initializerStatements.String()
	ptrRes = pointersInitializesStatements.String()
	return
}
func (md_band_propertytype *MD_Band_PropertyType) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(md_band_propertytype.GongMarshallField(stage, "Name"))
		initializerStatements.WriteString(md_band_propertytype.GongMarshallField(stage, "NilReason"))
		initializerStatements.WriteString(md_band_propertytype.GongMarshallField(stage, "Uuidref"))
		initializerStatements.WriteString(md_band_propertytype.GongMarshallField(stage, "Type_XSD"))
		initializerStatements.WriteString(md_band_propertytype.GongMarshallField(stage, "Href"))
		initializerStatements.WriteString(md_band_propertytype.GongMarshallField(stage, "Role"))
		initializerStatements.WriteString(md_band_propertytype.GongMarshallField(stage, "Arcrole"))
		initializerStatements.WriteString(md_band_propertytype.GongMarshallField(stage, "Title"))
		initializerStatements.WriteString(md_band_propertytype.GongMarshallField(stage, "Show"))
		initializerStatements.WriteString(md_band_propertytype.GongMarshallField(stage, "Actuate"))
		pointersInitializesStatements.WriteString(md_band_propertytype.GongMarshallField(stage, "MD_Band"))
	}
	initRes = initializerStatements.String()
	ptrRes = pointersInitializesStatements.String()
	return
}
func (md_band_type *MD_Band_Type) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(md_band_type.GongMarshallField(stage, "Name"))
		initializerStatements.WriteString(md_band_type.GongMarshallField(stage, "Id"))
		initializerStatements.WriteString(md_band_type.GongMarshallField(stage, "Uuid"))
	}
	initRes = initializerStatements.String()
	ptrRes = pointersInitializesStatements.String()
	return
}
func (md_contentinformation_propertytype *MD_ContentInformation_PropertyType) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(md_contentinformation_propertytype.GongMarshallField(stage, "Name"))
		initializerStatements.WriteString(md_contentinformation_propertytype.GongMarshallField(stage, "NilReason"))
		initializerStatements.WriteString(md_contentinformation_propertytype.GongMarshallField(stage, "Uuidref"))
		initializerStatements.WriteString(md_contentinformation_propertytype.GongMarshallField(stage, "Type_XSD"))
		initializerStatements.WriteString(md_contentinformation_propertytype.GongMarshallField(stage, "Href"))
		initializerStatements.WriteString(md_contentinformation_propertytype.GongMarshallField(stage, "Role"))
		initializerStatements.WriteString(md_contentinformation_propertytype.GongMarshallField(stage, "Arcrole"))
		initializerStatements.WriteString(md_contentinformation_propertytype.GongMarshallField(stage, "Title"))
		initializerStatements.WriteString(md_contentinformation_propertytype.GongMarshallField(stage, "Show"))
		initializerStatements.WriteString(md_contentinformation_propertytype.GongMarshallField(stage, "Actuate"))
		pointersInitializesStatements.WriteString(md_contentinformation_propertytype.GongMarshallField(stage, "MD_FeatureCatalogueDescription"))
		pointersInitializesStatements.WriteString(md_contentinformation_propertytype.GongMarshallField(stage, "MD_CoverageDescription"))
		pointersInitializesStatements.WriteString(md_contentinformation_propertytype.GongMarshallField(stage, "MD_ImageDescription"))
	}
	initRes = initializerStatements.String()
	ptrRes = pointersInitializesStatements.String()
	return
}
func (md_coveragecontenttypecode_propertytype *MD_CoverageContentTypeCode_PropertyType) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(md_coveragecontenttypecode_propertytype.GongMarshallField(stage, "Name"))
		initializerStatements.WriteString(md_coveragecontenttypecode_propertytype.GongMarshallField(stage, "NilReason"))
	}
	initRes = initializerStatements.String()
	ptrRes = pointersInitializesStatements.String()
	return
}
func (md_coveragedescription_propertytype *MD_CoverageDescription_PropertyType) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(md_coveragedescription_propertytype.GongMarshallField(stage, "Name"))
		initializerStatements.WriteString(md_coveragedescription_propertytype.GongMarshallField(stage, "NilReason"))
		initializerStatements.WriteString(md_coveragedescription_propertytype.GongMarshallField(stage, "Uuidref"))
		initializerStatements.WriteString(md_coveragedescription_propertytype.GongMarshallField(stage, "Type_XSD"))
		initializerStatements.WriteString(md_coveragedescription_propertytype.GongMarshallField(stage, "Href"))
		initializerStatements.WriteString(md_coveragedescription_propertytype.GongMarshallField(stage, "Role"))
		initializerStatements.WriteString(md_coveragedescription_propertytype.GongMarshallField(stage, "Arcrole"))
		initializerStatements.WriteString(md_coveragedescription_propertytype.GongMarshallField(stage, "Title"))
		initializerStatements.WriteString(md_coveragedescription_propertytype.GongMarshallField(stage, "Show"))
		initializerStatements.WriteString(md_coveragedescription_propertytype.GongMarshallField(stage, "Actuate"))
		pointersInitializesStatements.WriteString(md_coveragedescription_propertytype.GongMarshallField(stage, "MD_CoverageDescription"))
		pointersInitializesStatements.WriteString(md_coveragedescription_propertytype.GongMarshallField(stage, "MD_ImageDescription"))
	}
	initRes = initializerStatements.String()
	ptrRes = pointersInitializesStatements.String()
	return
}
func (md_coveragedescription_type *MD_CoverageDescription_Type) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(md_coveragedescription_type.GongMarshallField(stage, "Name"))
		initializerStatements.WriteString(md_coveragedescription_type.GongMarshallField(stage, "Id"))
		initializerStatements.WriteString(md_coveragedescription_type.GongMarshallField(stage, "Uuid"))
		pointersInitializesStatements.WriteString(md_coveragedescription_type.GongMarshallField(stage, "ContentType"))
		pointersInitializesStatements.WriteString(md_coveragedescription_type.GongMarshallField(stage, "Dimension"))
	}
	initRes = initializerStatements.String()
	ptrRes = pointersInitializesStatements.String()
	return
}
func (md_featurecataloguedescription_propertytype *MD_FeatureCatalogueDescription_PropertyType) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(md_featurecataloguedescription_propertytype.GongMarshallField(stage, "Name"))
		initializerStatements.WriteString(md_featurecataloguedescription_propertytype.GongMarshallField(stage, "NilReason"))
		initializerStatements.WriteString(md_featurecataloguedescription_propertytype.GongMarshallField(stage, "Uuidref"))
		initializerStatements.WriteString(md_featurecataloguedescription_propertytype.GongMarshallField(stage, "Type_XSD"))
		initializerStatements.WriteString(md_featurecataloguedescription_propertytype.GongMarshallField(stage, "Href"))
		initializerStatements.WriteString(md_featurecataloguedescription_propertytype.GongMarshallField(stage, "Role"))
		initializerStatements.WriteString(md_featurecataloguedescription_propertytype.GongMarshallField(stage, "Arcrole"))
		initializerStatements.WriteString(md_featurecataloguedescription_propertytype.GongMarshallField(stage, "Title"))
		initializerStatements.WriteString(md_featurecataloguedescription_propertytype.GongMarshallField(stage, "Show"))
		initializerStatements.WriteString(md_featurecataloguedescription_propertytype.GongMarshallField(stage, "Actuate"))
		pointersInitializesStatements.WriteString(md_featurecataloguedescription_propertytype.GongMarshallField(stage, "MD_FeatureCatalogueDescription"))
	}
	initRes = initializerStatements.String()
	ptrRes = pointersInitializesStatements.String()
	return
}
func (md_featurecataloguedescription_type *MD_FeatureCatalogueDescription_Type) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(md_featurecataloguedescription_type.GongMarshallField(stage, "Name"))
		initializerStatements.WriteString(md_featurecataloguedescription_type.GongMarshallField(stage, "Id"))
		initializerStatements.WriteString(md_featurecataloguedescription_type.GongMarshallField(stage, "Uuid"))
	}
	initRes = initializerStatements.String()
	ptrRes = pointersInitializesStatements.String()
	return
}
func (md_imagedescription_propertytype *MD_ImageDescription_PropertyType) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(md_imagedescription_propertytype.GongMarshallField(stage, "Name"))
		initializerStatements.WriteString(md_imagedescription_propertytype.GongMarshallField(stage, "NilReason"))
		initializerStatements.WriteString(md_imagedescription_propertytype.GongMarshallField(stage, "Uuidref"))
		initializerStatements.WriteString(md_imagedescription_propertytype.GongMarshallField(stage, "Type_XSD"))
		initializerStatements.WriteString(md_imagedescription_propertytype.GongMarshallField(stage, "Href"))
		initializerStatements.WriteString(md_imagedescription_propertytype.GongMarshallField(stage, "Role"))
		initializerStatements.WriteString(md_imagedescription_propertytype.GongMarshallField(stage, "Arcrole"))
		initializerStatements.WriteString(md_imagedescription_propertytype.GongMarshallField(stage, "Title"))
		initializerStatements.WriteString(md_imagedescription_propertytype.GongMarshallField(stage, "Show"))
		initializerStatements.WriteString(md_imagedescription_propertytype.GongMarshallField(stage, "Actuate"))
		pointersInitializesStatements.WriteString(md_imagedescription_propertytype.GongMarshallField(stage, "MD_ImageDescription"))
	}
	initRes = initializerStatements.String()
	ptrRes = pointersInitializesStatements.String()
	return
}
func (md_imagedescription_type *MD_ImageDescription_Type) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(md_imagedescription_type.GongMarshallField(stage, "Name"))
		initializerStatements.WriteString(md_imagedescription_type.GongMarshallField(stage, "Id"))
		initializerStatements.WriteString(md_imagedescription_type.GongMarshallField(stage, "Uuid"))
		pointersInitializesStatements.WriteString(md_imagedescription_type.GongMarshallField(stage, "ContentType"))
		pointersInitializesStatements.WriteString(md_imagedescription_type.GongMarshallField(stage, "Dimension"))
		pointersInitializesStatements.WriteString(md_imagedescription_type.GongMarshallField(stage, "ImagingCondition"))
	}
	initRes = initializerStatements.String()
	ptrRes = pointersInitializesStatements.String()
	return
}
func (md_imagingconditioncode_propertytype *MD_ImagingConditionCode_PropertyType) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(md_imagingconditioncode_propertytype.GongMarshallField(stage, "Name"))
		initializerStatements.WriteString(md_imagingconditioncode_propertytype.GongMarshallField(stage, "NilReason"))
	}
	initRes = initializerStatements.String()
	ptrRes = pointersInitializesStatements.String()
	return
}
func (md_rangedimension_propertytype *MD_RangeDimension_PropertyType) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(md_rangedimension_propertytype.GongMarshallField(stage, "Name"))
		initializerStatements.WriteString(md_rangedimension_propertytype.GongMarshallField(stage, "NilReason"))
		initializerStatements.WriteString(md_rangedimension_propertytype.GongMarshallField(stage, "Uuidref"))
		initializerStatements.WriteString(md_rangedimension_propertytype.GongMarshallField(stage, "Type_XSD"))
		initializerStatements.WriteString(md_rangedimension_propertytype.GongMarshallField(stage, "Href"))
		initializerStatements.WriteString(md_rangedimension_propertytype.GongMarshallField(stage, "Role"))
		initializerStatements.WriteString(md_rangedimension_propertytype.GongMarshallField(stage, "Arcrole"))
		initializerStatements.WriteString(md_rangedimension_propertytype.GongMarshallField(stage, "Title"))
		initializerStatements.WriteString(md_rangedimension_propertytype.GongMarshallField(stage, "Show"))
		initializerStatements.WriteString(md_rangedimension_propertytype.GongMarshallField(stage, "Actuate"))
		pointersInitializesStatements.WriteString(md_rangedimension_propertytype.GongMarshallField(stage, "MD_RangeDimension"))
		pointersInitializesStatements.WriteString(md_rangedimension_propertytype.GongMarshallField(stage, "MD_Band"))
	}
	initRes = initializerStatements.String()
	ptrRes = pointersInitializesStatements.String()
	return
}
func (md_rangedimension_type *MD_RangeDimension_Type) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(md_rangedimension_type.GongMarshallField(stage, "Name"))
		initializerStatements.WriteString(md_rangedimension_type.GongMarshallField(stage, "Id"))
		initializerStatements.WriteString(md_rangedimension_type.GongMarshallField(stage, "Uuid"))
	}
	initRes = initializerStatements.String()
	ptrRes = pointersInitializesStatements.String()
	return
}
