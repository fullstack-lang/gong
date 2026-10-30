// generated code - do not edit
package metadataextension

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
	{{Identifier}} := (&metadataextension.{{GeneratedStructName}}{Name: {{GeneratedFieldNameValue}}}).Stage(stage)`

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
			const fnSig = "func _(stage *metadataextension.Stage) {"
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
			strings.HasSuffix(trimmedContentBeforeBrace, "func _(stage *metadataextension.Stage) {") ||
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
	modelsPackageName := "github.com/fullstack-lang/gong/app/xsd/tests/aixm/go/models/metadataextension"
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
		for _, md_datatypecode_propertytype := range __gong__sortStageSetInstances(stage.MD_DatatypeCode_PropertyTypes, stage.MD_DatatypeCode_PropertyType_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			md_datatypecode_propertytypeIdent := md_datatypecode_propertytype.GongGetIdentifier(stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&metadataextension.MD_DatatypeCode_PropertyType{Name: %s}).Stage(stage)", md_datatypecode_propertytypeIdent, __gong__toRawStringLiteral(md_datatypecode_propertytype.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", md_datatypecode_propertytypeIdent, __gong__toRawStringLiteral(md_datatypecode_propertytype.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.NilReason = %s", md_datatypecode_propertytypeIdent, __gong__toRawStringLiteral(md_datatypecode_propertytype.NilReason)))
		}
	}
	if stage != nil {
		for _, md_extendedelementinformation_propertytype := range __gong__sortStageSetInstances(stage.MD_ExtendedElementInformation_PropertyTypes, stage.MD_ExtendedElementInformation_PropertyType_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			md_extendedelementinformation_propertytypeIdent := md_extendedelementinformation_propertytype.GongGetIdentifier(stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&metadataextension.MD_ExtendedElementInformation_PropertyType{Name: %s}).Stage(stage)", md_extendedelementinformation_propertytypeIdent, __gong__toRawStringLiteral(md_extendedelementinformation_propertytype.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", md_extendedelementinformation_propertytypeIdent, __gong__toRawStringLiteral(md_extendedelementinformation_propertytype.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.NilReason = %s", md_extendedelementinformation_propertytypeIdent, __gong__toRawStringLiteral(md_extendedelementinformation_propertytype.NilReason)))
			values.WriteString(fmt.Sprintf("\n\t%s.Uuidref = %s", md_extendedelementinformation_propertytypeIdent, __gong__toRawStringLiteral(md_extendedelementinformation_propertytype.Uuidref)))
			values.WriteString(fmt.Sprintf("\n\t%s.Type_XSD = %s", md_extendedelementinformation_propertytypeIdent, __gong__toRawStringLiteral(md_extendedelementinformation_propertytype.Type_XSD)))
			values.WriteString(fmt.Sprintf("\n\t%s.Href = %s", md_extendedelementinformation_propertytypeIdent, __gong__toRawStringLiteral(md_extendedelementinformation_propertytype.Href)))
			values.WriteString(fmt.Sprintf("\n\t%s.Role = %s", md_extendedelementinformation_propertytypeIdent, __gong__toRawStringLiteral(md_extendedelementinformation_propertytype.Role)))
			values.WriteString(fmt.Sprintf("\n\t%s.Arcrole = %s", md_extendedelementinformation_propertytypeIdent, __gong__toRawStringLiteral(md_extendedelementinformation_propertytype.Arcrole)))
			values.WriteString(fmt.Sprintf("\n\t%s.Title = %s", md_extendedelementinformation_propertytypeIdent, __gong__toRawStringLiteral(md_extendedelementinformation_propertytype.Title)))
			values.WriteString(fmt.Sprintf("\n\t%s.Show = %s", md_extendedelementinformation_propertytypeIdent, __gong__toRawStringLiteral(md_extendedelementinformation_propertytype.Show)))
			values.WriteString(fmt.Sprintf("\n\t%s.Actuate = %s", md_extendedelementinformation_propertytypeIdent, __gong__toRawStringLiteral(md_extendedelementinformation_propertytype.Actuate)))
			if md_extendedelementinformation_propertytype.MD_ExtendedElementInformation != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := md_extendedelementinformation_propertytype.MD_ExtendedElementInformation.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.MD_ExtendedElementInformation = %s", md_extendedelementinformation_propertytypeIdent, targetIdent))
			}
		}
	}
	if stage != nil {
		for _, md_extendedelementinformation_type := range __gong__sortStageSetInstances(stage.MD_ExtendedElementInformation_Types, stage.MD_ExtendedElementInformation_Type_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			md_extendedelementinformation_typeIdent := md_extendedelementinformation_type.GongGetIdentifier(stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&metadataextension.MD_ExtendedElementInformation_Type{Name: %s}).Stage(stage)", md_extendedelementinformation_typeIdent, __gong__toRawStringLiteral(md_extendedelementinformation_type.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", md_extendedelementinformation_typeIdent, __gong__toRawStringLiteral(md_extendedelementinformation_type.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Id = %s", md_extendedelementinformation_typeIdent, __gong__toRawStringLiteral(md_extendedelementinformation_type.Id)))
			values.WriteString(fmt.Sprintf("\n\t%s.Uuid = %s", md_extendedelementinformation_typeIdent, __gong__toRawStringLiteral(md_extendedelementinformation_type.Uuid)))
			if md_extendedelementinformation_type.Obligation != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := md_extendedelementinformation_type.Obligation.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Obligation = %s", md_extendedelementinformation_typeIdent, targetIdent))
			}
			if md_extendedelementinformation_type.DataType != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := md_extendedelementinformation_type.DataType.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.DataType = %s", md_extendedelementinformation_typeIdent, targetIdent))
			}
		}
	}
	if stage != nil {
		for _, md_metadataextensioninformation_propertytype := range __gong__sortStageSetInstances(stage.MD_MetadataExtensionInformation_PropertyTypes, stage.MD_MetadataExtensionInformation_PropertyType_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			md_metadataextensioninformation_propertytypeIdent := md_metadataextensioninformation_propertytype.GongGetIdentifier(stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&metadataextension.MD_MetadataExtensionInformation_PropertyType{Name: %s}).Stage(stage)", md_metadataextensioninformation_propertytypeIdent, __gong__toRawStringLiteral(md_metadataextensioninformation_propertytype.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", md_metadataextensioninformation_propertytypeIdent, __gong__toRawStringLiteral(md_metadataextensioninformation_propertytype.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.NilReason = %s", md_metadataextensioninformation_propertytypeIdent, __gong__toRawStringLiteral(md_metadataextensioninformation_propertytype.NilReason)))
			values.WriteString(fmt.Sprintf("\n\t%s.Uuidref = %s", md_metadataextensioninformation_propertytypeIdent, __gong__toRawStringLiteral(md_metadataextensioninformation_propertytype.Uuidref)))
			values.WriteString(fmt.Sprintf("\n\t%s.Type_XSD = %s", md_metadataextensioninformation_propertytypeIdent, __gong__toRawStringLiteral(md_metadataextensioninformation_propertytype.Type_XSD)))
			values.WriteString(fmt.Sprintf("\n\t%s.Href = %s", md_metadataextensioninformation_propertytypeIdent, __gong__toRawStringLiteral(md_metadataextensioninformation_propertytype.Href)))
			values.WriteString(fmt.Sprintf("\n\t%s.Role = %s", md_metadataextensioninformation_propertytypeIdent, __gong__toRawStringLiteral(md_metadataextensioninformation_propertytype.Role)))
			values.WriteString(fmt.Sprintf("\n\t%s.Arcrole = %s", md_metadataextensioninformation_propertytypeIdent, __gong__toRawStringLiteral(md_metadataextensioninformation_propertytype.Arcrole)))
			values.WriteString(fmt.Sprintf("\n\t%s.Title = %s", md_metadataextensioninformation_propertytypeIdent, __gong__toRawStringLiteral(md_metadataextensioninformation_propertytype.Title)))
			values.WriteString(fmt.Sprintf("\n\t%s.Show = %s", md_metadataextensioninformation_propertytypeIdent, __gong__toRawStringLiteral(md_metadataextensioninformation_propertytype.Show)))
			values.WriteString(fmt.Sprintf("\n\t%s.Actuate = %s", md_metadataextensioninformation_propertytypeIdent, __gong__toRawStringLiteral(md_metadataextensioninformation_propertytype.Actuate)))
			if md_metadataextensioninformation_propertytype.MD_MetadataExtensionInformation != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := md_metadataextensioninformation_propertytype.MD_MetadataExtensionInformation.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.MD_MetadataExtensionInformation = %s", md_metadataextensioninformation_propertytypeIdent, targetIdent))
			}
		}
	}
	if stage != nil {
		for _, md_metadataextensioninformation_type := range __gong__sortStageSetInstances(stage.MD_MetadataExtensionInformation_Types, stage.MD_MetadataExtensionInformation_Type_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			md_metadataextensioninformation_typeIdent := md_metadataextensioninformation_type.GongGetIdentifier(stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&metadataextension.MD_MetadataExtensionInformation_Type{Name: %s}).Stage(stage)", md_metadataextensioninformation_typeIdent, __gong__toRawStringLiteral(md_metadataextensioninformation_type.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", md_metadataextensioninformation_typeIdent, __gong__toRawStringLiteral(md_metadataextensioninformation_type.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Id = %s", md_metadataextensioninformation_typeIdent, __gong__toRawStringLiteral(md_metadataextensioninformation_type.Id)))
			values.WriteString(fmt.Sprintf("\n\t%s.Uuid = %s", md_metadataextensioninformation_typeIdent, __gong__toRawStringLiteral(md_metadataextensioninformation_type.Uuid)))
			for _, elem := range md_metadataextensioninformation_type.ExtendedElementInformation {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := elem.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.ExtendedElementInformation = append(%s.ExtendedElementInformation, %s)", md_metadataextensioninformation_typeIdent, md_metadataextensioninformation_typeIdent, targetIdent))
			}
		}
	}
	if stage != nil {
		for _, md_obligationcode_propertytype := range __gong__sortStageSetInstances(stage.MD_ObligationCode_PropertyTypes, stage.MD_ObligationCode_PropertyType_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			md_obligationcode_propertytypeIdent := md_obligationcode_propertytype.GongGetIdentifier(stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&metadataextension.MD_ObligationCode_PropertyType{Name: %s}).Stage(stage)", md_obligationcode_propertytypeIdent, __gong__toRawStringLiteral(md_obligationcode_propertytype.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", md_obligationcode_propertytypeIdent, __gong__toRawStringLiteral(md_obligationcode_propertytype.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.NilReason = %s", md_obligationcode_propertytypeIdent, __gong__toRawStringLiteral(md_obligationcode_propertytype.NilReason)))
			if md_obligationcode_propertytype.MD_ObligationCode.ToCodeString() != "" {
				values.WriteString(fmt.Sprintf("\n\t%s.MD_ObligationCode = metadataextension.%s", md_obligationcode_propertytypeIdent, md_obligationcode_propertytype.MD_ObligationCode.ToCodeString()))
			} else {
				values.WriteString(fmt.Sprintf("\n\t%s.MD_ObligationCode = \"\"", md_obligationcode_propertytypeIdent))
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
go run github.com/fullstack-lang/gong/app/xsd/tests/aixm/go/models/metadataextension/probe/cmd/probe@%s stage.go
*/
package %s

import (
	"slices"
	"time"

	"github.com/fullstack-lang/gong/app/xsd/tests/aixm/go/models/metadataextension"%s
)

var (
	_ time.Time
	_ = slices.Index[[]int, int]

	_ *metadataextension.Stage
)%s

// function will stage objects across all coordinated stages
func _(stage *metadataextension.Stage) {

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

	if modelsPackageName != "" && modelsPackageName != "github.com/fullstack-lang/gong/app/xsd/tests/aixm/go/models/metadataextension" {
		res = strings.ReplaceAll(res, "\"github.com/fullstack-lang/gong/app/xsd/tests/aixm/go/models/metadataextension\"", fmt.Sprintf("%q", modelsPackageName))
	}

	return res, nil
}

// insertion point for marshall field methods
func (md_datatypecode_propertytype *MD_DatatypeCode_PropertyType) GongMarshallField(stage *Stage, fieldName string) (res string) {
	ident := md_datatypecode_propertytype.GongGetIdentifier(stage)
	_ = ident
	switch fieldName {
	case "Name":
		res = __gong__marshallString(ident, "Name", md_datatypecode_propertytype.Name)
	case "NilReason":
		res = __gong__marshallString(ident, "NilReason", md_datatypecode_propertytype.NilReason)

	default:
		log.Panicf("Unknown field %s for Gongstruct MD_DatatypeCode_PropertyType", fieldName)
	}
	return
}

func (md_extendedelementinformation_propertytype *MD_ExtendedElementInformation_PropertyType) GongMarshallField(stage *Stage, fieldName string) (res string) {
	ident := md_extendedelementinformation_propertytype.GongGetIdentifier(stage)
	_ = ident
	switch fieldName {
	case "Name":
		res = __gong__marshallString(ident, "Name", md_extendedelementinformation_propertytype.Name)
	case "NilReason":
		res = __gong__marshallString(ident, "NilReason", md_extendedelementinformation_propertytype.NilReason)
	case "Uuidref":
		res = __gong__marshallString(ident, "Uuidref", md_extendedelementinformation_propertytype.Uuidref)
	case "Type_XSD":
		res = __gong__marshallString(ident, "Type_XSD", md_extendedelementinformation_propertytype.Type_XSD)
	case "Href":
		res = __gong__marshallString(ident, "Href", md_extendedelementinformation_propertytype.Href)
	case "Role":
		res = __gong__marshallString(ident, "Role", md_extendedelementinformation_propertytype.Role)
	case "Arcrole":
		res = __gong__marshallString(ident, "Arcrole", md_extendedelementinformation_propertytype.Arcrole)
	case "Title":
		res = __gong__marshallString(ident, "Title", md_extendedelementinformation_propertytype.Title)
	case "Show":
		res = __gong__marshallString(ident, "Show", md_extendedelementinformation_propertytype.Show)
	case "Actuate":
		res = __gong__marshallString(ident, "Actuate", md_extendedelementinformation_propertytype.Actuate)

	case "MD_ExtendedElementInformation":
		if md_extendedelementinformation_propertytype.MD_ExtendedElementInformation != nil {
			res = __gong__marshallPointer(ident, "MD_ExtendedElementInformation", md_extendedelementinformation_propertytype.MD_ExtendedElementInformation.GongGetIdentifier(stage))
		} else {
			res = __gong__marshallPointer(ident, "MD_ExtendedElementInformation", "nil")
		}
	default:
		log.Panicf("Unknown field %s for Gongstruct MD_ExtendedElementInformation_PropertyType", fieldName)
	}
	return
}

func (md_extendedelementinformation_type *MD_ExtendedElementInformation_Type) GongMarshallField(stage *Stage, fieldName string) (res string) {
	ident := md_extendedelementinformation_type.GongGetIdentifier(stage)
	_ = ident
	switch fieldName {
	case "Name":
		res = __gong__marshallString(ident, "Name", md_extendedelementinformation_type.Name)
	case "Id":
		res = __gong__marshallString(ident, "Id", md_extendedelementinformation_type.Id)
	case "Uuid":
		res = __gong__marshallString(ident, "Uuid", md_extendedelementinformation_type.Uuid)

	case "Obligation":
		if md_extendedelementinformation_type.Obligation != nil {
			res = __gong__marshallPointer(ident, "Obligation", md_extendedelementinformation_type.Obligation.GongGetIdentifier(stage))
		} else {
			res = __gong__marshallPointer(ident, "Obligation", "nil")
		}
	case "DataType":
		if md_extendedelementinformation_type.DataType != nil {
			res = __gong__marshallPointer(ident, "DataType", md_extendedelementinformation_type.DataType.GongGetIdentifier(stage))
		} else {
			res = __gong__marshallPointer(ident, "DataType", "nil")
		}
	default:
		log.Panicf("Unknown field %s for Gongstruct MD_ExtendedElementInformation_Type", fieldName)
	}
	return
}

func (md_metadataextensioninformation_propertytype *MD_MetadataExtensionInformation_PropertyType) GongMarshallField(stage *Stage, fieldName string) (res string) {
	ident := md_metadataextensioninformation_propertytype.GongGetIdentifier(stage)
	_ = ident
	switch fieldName {
	case "Name":
		res = __gong__marshallString(ident, "Name", md_metadataextensioninformation_propertytype.Name)
	case "NilReason":
		res = __gong__marshallString(ident, "NilReason", md_metadataextensioninformation_propertytype.NilReason)
	case "Uuidref":
		res = __gong__marshallString(ident, "Uuidref", md_metadataextensioninformation_propertytype.Uuidref)
	case "Type_XSD":
		res = __gong__marshallString(ident, "Type_XSD", md_metadataextensioninformation_propertytype.Type_XSD)
	case "Href":
		res = __gong__marshallString(ident, "Href", md_metadataextensioninformation_propertytype.Href)
	case "Role":
		res = __gong__marshallString(ident, "Role", md_metadataextensioninformation_propertytype.Role)
	case "Arcrole":
		res = __gong__marshallString(ident, "Arcrole", md_metadataextensioninformation_propertytype.Arcrole)
	case "Title":
		res = __gong__marshallString(ident, "Title", md_metadataextensioninformation_propertytype.Title)
	case "Show":
		res = __gong__marshallString(ident, "Show", md_metadataextensioninformation_propertytype.Show)
	case "Actuate":
		res = __gong__marshallString(ident, "Actuate", md_metadataextensioninformation_propertytype.Actuate)

	case "MD_MetadataExtensionInformation":
		if md_metadataextensioninformation_propertytype.MD_MetadataExtensionInformation != nil {
			res = __gong__marshallPointer(ident, "MD_MetadataExtensionInformation", md_metadataextensioninformation_propertytype.MD_MetadataExtensionInformation.GongGetIdentifier(stage))
		} else {
			res = __gong__marshallPointer(ident, "MD_MetadataExtensionInformation", "nil")
		}
	default:
		log.Panicf("Unknown field %s for Gongstruct MD_MetadataExtensionInformation_PropertyType", fieldName)
	}
	return
}

func (md_metadataextensioninformation_type *MD_MetadataExtensionInformation_Type) GongMarshallField(stage *Stage, fieldName string) (res string) {
	ident := md_metadataextensioninformation_type.GongGetIdentifier(stage)
	_ = ident
	switch fieldName {
	case "Name":
		res = __gong__marshallString(ident, "Name", md_metadataextensioninformation_type.Name)
	case "Id":
		res = __gong__marshallString(ident, "Id", md_metadataextensioninformation_type.Id)
	case "Uuid":
		res = __gong__marshallString(ident, "Uuid", md_metadataextensioninformation_type.Uuid)

	case "ExtendedElementInformation":
		var sb strings.Builder
		for _, _md_extendedelementinformation_propertytype := range md_metadataextensioninformation_type.ExtendedElementInformation {
			sb.WriteString(__gong__marshallSliceOfPointers(ident, "ExtendedElementInformation", _md_extendedelementinformation_propertytype.GongGetIdentifier(stage)))
		}
		res = sb.String()
	default:
		log.Panicf("Unknown field %s for Gongstruct MD_MetadataExtensionInformation_Type", fieldName)
	}
	return
}

func (md_obligationcode_propertytype *MD_ObligationCode_PropertyType) GongMarshallField(stage *Stage, fieldName string) (res string) {
	ident := md_obligationcode_propertytype.GongGetIdentifier(stage)
	_ = ident
	switch fieldName {
	case "Name":
		res = __gong__marshallString(ident, "Name", md_obligationcode_propertytype.Name)
	case "NilReason":
		res = __gong__marshallString(ident, "NilReason", md_obligationcode_propertytype.NilReason)
	case "MD_ObligationCode":
		res = __gong__marshallEnumString(ident, "MD_ObligationCode", md_obligationcode_propertytype.MD_ObligationCode.ToCodeString())

	default:
		log.Panicf("Unknown field %s for Gongstruct MD_ObligationCode_PropertyType", fieldName)
	}
	return
}

// insertion point for marshall all fields methods
func (md_datatypecode_propertytype *MD_DatatypeCode_PropertyType) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(md_datatypecode_propertytype.GongMarshallField(stage, "Name"))
		initializerStatements.WriteString(md_datatypecode_propertytype.GongMarshallField(stage, "NilReason"))
	}
	initRes = initializerStatements.String()
	ptrRes = pointersInitializesStatements.String()
	return
}
func (md_extendedelementinformation_propertytype *MD_ExtendedElementInformation_PropertyType) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(md_extendedelementinformation_propertytype.GongMarshallField(stage, "Name"))
		initializerStatements.WriteString(md_extendedelementinformation_propertytype.GongMarshallField(stage, "NilReason"))
		initializerStatements.WriteString(md_extendedelementinformation_propertytype.GongMarshallField(stage, "Uuidref"))
		initializerStatements.WriteString(md_extendedelementinformation_propertytype.GongMarshallField(stage, "Type_XSD"))
		initializerStatements.WriteString(md_extendedelementinformation_propertytype.GongMarshallField(stage, "Href"))
		initializerStatements.WriteString(md_extendedelementinformation_propertytype.GongMarshallField(stage, "Role"))
		initializerStatements.WriteString(md_extendedelementinformation_propertytype.GongMarshallField(stage, "Arcrole"))
		initializerStatements.WriteString(md_extendedelementinformation_propertytype.GongMarshallField(stage, "Title"))
		initializerStatements.WriteString(md_extendedelementinformation_propertytype.GongMarshallField(stage, "Show"))
		initializerStatements.WriteString(md_extendedelementinformation_propertytype.GongMarshallField(stage, "Actuate"))
		pointersInitializesStatements.WriteString(md_extendedelementinformation_propertytype.GongMarshallField(stage, "MD_ExtendedElementInformation"))
	}
	initRes = initializerStatements.String()
	ptrRes = pointersInitializesStatements.String()
	return
}
func (md_extendedelementinformation_type *MD_ExtendedElementInformation_Type) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(md_extendedelementinformation_type.GongMarshallField(stage, "Name"))
		initializerStatements.WriteString(md_extendedelementinformation_type.GongMarshallField(stage, "Id"))
		initializerStatements.WriteString(md_extendedelementinformation_type.GongMarshallField(stage, "Uuid"))
		pointersInitializesStatements.WriteString(md_extendedelementinformation_type.GongMarshallField(stage, "Obligation"))
		pointersInitializesStatements.WriteString(md_extendedelementinformation_type.GongMarshallField(stage, "DataType"))
	}
	initRes = initializerStatements.String()
	ptrRes = pointersInitializesStatements.String()
	return
}
func (md_metadataextensioninformation_propertytype *MD_MetadataExtensionInformation_PropertyType) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(md_metadataextensioninformation_propertytype.GongMarshallField(stage, "Name"))
		initializerStatements.WriteString(md_metadataextensioninformation_propertytype.GongMarshallField(stage, "NilReason"))
		initializerStatements.WriteString(md_metadataextensioninformation_propertytype.GongMarshallField(stage, "Uuidref"))
		initializerStatements.WriteString(md_metadataextensioninformation_propertytype.GongMarshallField(stage, "Type_XSD"))
		initializerStatements.WriteString(md_metadataextensioninformation_propertytype.GongMarshallField(stage, "Href"))
		initializerStatements.WriteString(md_metadataextensioninformation_propertytype.GongMarshallField(stage, "Role"))
		initializerStatements.WriteString(md_metadataextensioninformation_propertytype.GongMarshallField(stage, "Arcrole"))
		initializerStatements.WriteString(md_metadataextensioninformation_propertytype.GongMarshallField(stage, "Title"))
		initializerStatements.WriteString(md_metadataextensioninformation_propertytype.GongMarshallField(stage, "Show"))
		initializerStatements.WriteString(md_metadataextensioninformation_propertytype.GongMarshallField(stage, "Actuate"))
		pointersInitializesStatements.WriteString(md_metadataextensioninformation_propertytype.GongMarshallField(stage, "MD_MetadataExtensionInformation"))
	}
	initRes = initializerStatements.String()
	ptrRes = pointersInitializesStatements.String()
	return
}
func (md_metadataextensioninformation_type *MD_MetadataExtensionInformation_Type) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(md_metadataextensioninformation_type.GongMarshallField(stage, "Name"))
		initializerStatements.WriteString(md_metadataextensioninformation_type.GongMarshallField(stage, "Id"))
		initializerStatements.WriteString(md_metadataextensioninformation_type.GongMarshallField(stage, "Uuid"))
		pointersInitializesStatements.WriteString(md_metadataextensioninformation_type.GongMarshallField(stage, "ExtendedElementInformation"))
	}
	initRes = initializerStatements.String()
	ptrRes = pointersInitializesStatements.String()
	return
}
func (md_obligationcode_propertytype *MD_ObligationCode_PropertyType) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(md_obligationcode_propertytype.GongMarshallField(stage, "Name"))
		initializerStatements.WriteString(md_obligationcode_propertytype.GongMarshallField(stage, "NilReason"))
		initializerStatements.WriteString(md_obligationcode_propertytype.GongMarshallField(stage, "MD_ObligationCode"))
	}
	initRes = initializerStatements.String()
	ptrRes = pointersInitializesStatements.String()
	return
}
