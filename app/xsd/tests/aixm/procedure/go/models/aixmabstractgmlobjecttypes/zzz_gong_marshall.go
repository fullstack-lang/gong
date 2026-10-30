// generated code - do not edit
package aixmabstractgmlobjecttypes

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
	{{Identifier}} := (&aixmabstractgmlobjecttypes.{{GeneratedStructName}}{Name: {{GeneratedFieldNameValue}}}).Stage(stage)`

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
			const fnSig = "func _(stage *aixmabstractgmlobjecttypes.Stage) {"
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
			strings.HasSuffix(trimmedContentBeforeBrace, "func _(stage *aixmabstractgmlobjecttypes.Stage) {") ||
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
	modelsPackageName := "github.com/fullstack-lang/gong/app/xsd/tests/aixm/procedure/go/models/aixmabstractgmlobjecttypes"
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
		for _, abstractaixmfeaturebasetype := range __gong__sortStageSetInstances(stage.AbstractAIXMFeatureBaseTypes, stage.AbstractAIXMFeatureBaseType_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			abstractaixmfeaturebasetypeIdent := abstractaixmfeaturebasetype.GongGetIdentifier(stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&aixmabstractgmlobjecttypes.AbstractAIXMFeatureBaseType{Name: %s}).Stage(stage)", abstractaixmfeaturebasetypeIdent, __gong__toRawStringLiteral(abstractaixmfeaturebasetype.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", abstractaixmfeaturebasetypeIdent, __gong__toRawStringLiteral(abstractaixmfeaturebasetype.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Id = %s", abstractaixmfeaturebasetypeIdent, __gong__toRawStringLiteral(abstractaixmfeaturebasetype.Id)))
			values.WriteString(fmt.Sprintf("\n\t%s.Description = %s", abstractaixmfeaturebasetypeIdent, __gong__toRawStringLiteral(abstractaixmfeaturebasetype.Description)))
		}
	}
	if stage != nil {
		for _, abstractaixmfeaturetype := range __gong__sortStageSetInstances(stage.AbstractAIXMFeatureTypes, stage.AbstractAIXMFeatureType_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			abstractaixmfeaturetypeIdent := abstractaixmfeaturetype.GongGetIdentifier(stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&aixmabstractgmlobjecttypes.AbstractAIXMFeatureType{Name: %s}).Stage(stage)", abstractaixmfeaturetypeIdent, __gong__toRawStringLiteral(abstractaixmfeaturetype.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", abstractaixmfeaturetypeIdent, __gong__toRawStringLiteral(abstractaixmfeaturetype.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Id = %s", abstractaixmfeaturetypeIdent, __gong__toRawStringLiteral(abstractaixmfeaturetype.Id)))
			values.WriteString(fmt.Sprintf("\n\t%s.Description = %s", abstractaixmfeaturetypeIdent, __gong__toRawStringLiteral(abstractaixmfeaturetype.Description)))
			if abstractaixmfeaturetype.FeatureMetadata != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := abstractaixmfeaturetype.FeatureMetadata.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.FeatureMetadata = %s", abstractaixmfeaturetypeIdent, targetIdent))
			}
		}
	}
	if stage != nil {
		for _, abstractaixmmessagebasetype := range __gong__sortStageSetInstances(stage.AbstractAIXMMessageBaseTypes, stage.AbstractAIXMMessageBaseType_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			abstractaixmmessagebasetypeIdent := abstractaixmmessagebasetype.GongGetIdentifier(stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&aixmabstractgmlobjecttypes.AbstractAIXMMessageBaseType{Name: %s}).Stage(stage)", abstractaixmmessagebasetypeIdent, __gong__toRawStringLiteral(abstractaixmmessagebasetype.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", abstractaixmmessagebasetypeIdent, __gong__toRawStringLiteral(abstractaixmmessagebasetype.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Id = %s", abstractaixmmessagebasetypeIdent, __gong__toRawStringLiteral(abstractaixmmessagebasetype.Id)))
			values.WriteString(fmt.Sprintf("\n\t%s.Description = %s", abstractaixmmessagebasetypeIdent, __gong__toRawStringLiteral(abstractaixmmessagebasetype.Description)))
			values.WriteString(fmt.Sprintf("\n\t%s.SequenceNumber = %d", abstractaixmmessagebasetypeIdent, abstractaixmmessagebasetype.SequenceNumber))
			if abstractaixmmessagebasetype.MessageMetadata != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := abstractaixmmessagebasetype.MessageMetadata.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.MessageMetadata = %s", abstractaixmmessagebasetypeIdent, targetIdent))
			}
		}
	}
	if stage != nil {
		for _, abstractaixmmessagetype := range __gong__sortStageSetInstances(stage.AbstractAIXMMessageTypes, stage.AbstractAIXMMessageType_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			abstractaixmmessagetypeIdent := abstractaixmmessagetype.GongGetIdentifier(stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&aixmabstractgmlobjecttypes.AbstractAIXMMessageType{Name: %s}).Stage(stage)", abstractaixmmessagetypeIdent, __gong__toRawStringLiteral(abstractaixmmessagetype.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", abstractaixmmessagetypeIdent, __gong__toRawStringLiteral(abstractaixmmessagetype.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Id = %s", abstractaixmmessagetypeIdent, __gong__toRawStringLiteral(abstractaixmmessagetype.Id)))
			values.WriteString(fmt.Sprintf("\n\t%s.Description = %s", abstractaixmmessagetypeIdent, __gong__toRawStringLiteral(abstractaixmmessagetype.Description)))
			values.WriteString(fmt.Sprintf("\n\t%s.SequenceNumber = %d", abstractaixmmessagetypeIdent, abstractaixmmessagetype.SequenceNumber))
			values.WriteString(fmt.Sprintf("\n\t%s.AggregationType = %s", abstractaixmmessagetypeIdent, __gong__toRawStringLiteral(abstractaixmmessagetype.AggregationType)))
			if abstractaixmmessagetype.MessageMetadata != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := abstractaixmmessagetype.MessageMetadata.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.MessageMetadata = %s", abstractaixmmessagetypeIdent, targetIdent))
			}
		}
	}
	if stage != nil {
		for _, abstractaixmobjecttype := range __gong__sortStageSetInstances(stage.AbstractAIXMObjectTypes, stage.AbstractAIXMObjectType_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			abstractaixmobjecttypeIdent := abstractaixmobjecttype.GongGetIdentifier(stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&aixmabstractgmlobjecttypes.AbstractAIXMObjectType{Name: %s}).Stage(stage)", abstractaixmobjecttypeIdent, __gong__toRawStringLiteral(abstractaixmobjecttype.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", abstractaixmobjecttypeIdent, __gong__toRawStringLiteral(abstractaixmobjecttype.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Id = %s", abstractaixmobjecttypeIdent, __gong__toRawStringLiteral(abstractaixmobjecttype.Id)))
		}
	}
	if stage != nil {
		for _, abstractaixmpropertytype := range __gong__sortStageSetInstances(stage.AbstractAIXMPropertyTypes, stage.AbstractAIXMPropertyType_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			abstractaixmpropertytypeIdent := abstractaixmpropertytype.GongGetIdentifier(stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&aixmabstractgmlobjecttypes.AbstractAIXMPropertyType{Name: %s}).Stage(stage)", abstractaixmpropertytypeIdent, __gong__toRawStringLiteral(abstractaixmpropertytype.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", abstractaixmpropertytypeIdent, __gong__toRawStringLiteral(abstractaixmpropertytype.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.NilReason = %s", abstractaixmpropertytypeIdent, __gong__toRawStringLiteral(abstractaixmpropertytype.NilReason)))
		}
	}
	if stage != nil {
		for _, abstractaixmtimeslicebasetype := range __gong__sortStageSetInstances(stage.AbstractAIXMTimeSliceBaseTypes, stage.AbstractAIXMTimeSliceBaseType_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			abstractaixmtimeslicebasetypeIdent := abstractaixmtimeslicebasetype.GongGetIdentifier(stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&aixmabstractgmlobjecttypes.AbstractAIXMTimeSliceBaseType{Name: %s}).Stage(stage)", abstractaixmtimeslicebasetypeIdent, __gong__toRawStringLiteral(abstractaixmtimeslicebasetype.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", abstractaixmtimeslicebasetypeIdent, __gong__toRawStringLiteral(abstractaixmtimeslicebasetype.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Id = %s", abstractaixmtimeslicebasetypeIdent, __gong__toRawStringLiteral(abstractaixmtimeslicebasetype.Id)))
		}
	}
	if stage != nil {
		for _, abstractaixmtimeslicetype := range __gong__sortStageSetInstances(stage.AbstractAIXMTimeSliceTypes, stage.AbstractAIXMTimeSliceType_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			abstractaixmtimeslicetypeIdent := abstractaixmtimeslicetype.GongGetIdentifier(stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&aixmabstractgmlobjecttypes.AbstractAIXMTimeSliceType{Name: %s}).Stage(stage)", abstractaixmtimeslicetypeIdent, __gong__toRawStringLiteral(abstractaixmtimeslicetype.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
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
			if abstractaixmtimeslicetype.TimeSliceMetadata != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := abstractaixmtimeslicetype.TimeSliceMetadata.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.TimeSliceMetadata = %s", abstractaixmtimeslicetypeIdent, targetIdent))
			}
		}
	}
	if stage != nil {
		for _, abstractextensiontype := range __gong__sortStageSetInstances(stage.AbstractExtensionTypes, stage.AbstractExtensionType_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			abstractextensiontypeIdent := abstractextensiontype.GongGetIdentifier(stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&aixmabstractgmlobjecttypes.AbstractExtensionType{Name: %s}).Stage(stage)", abstractextensiontypeIdent, __gong__toRawStringLiteral(abstractextensiontype.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", abstractextensiontypeIdent, __gong__toRawStringLiteral(abstractextensiontype.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Id = %s", abstractextensiontypeIdent, __gong__toRawStringLiteral(abstractextensiontype.Id)))
			values.WriteString(fmt.Sprintf("\n\t%s.Description = %s", abstractextensiontypeIdent, __gong__toRawStringLiteral(abstractextensiontype.Description)))
		}
	}
	if stage != nil {
		for _, featuremetadatapropertytype := range __gong__sortStageSetInstances(stage.FeatureMetadataPropertyTypes, stage.FeatureMetadataPropertyType_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			featuremetadatapropertytypeIdent := featuremetadatapropertytype.GongGetIdentifier(stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&aixmabstractgmlobjecttypes.FeatureMetadataPropertyType{Name: %s}).Stage(stage)", featuremetadatapropertytypeIdent, __gong__toRawStringLiteral(featuremetadatapropertytype.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", featuremetadatapropertytypeIdent, __gong__toRawStringLiteral(featuremetadatapropertytype.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Owns = %t", featuremetadatapropertytypeIdent, featuremetadatapropertytype.Owns))
		}
	}
	if stage != nil {
		for _, featuretimeslicemetadatapropertytype := range __gong__sortStageSetInstances(stage.FeatureTimeSliceMetadataPropertyTypes, stage.FeatureTimeSliceMetadataPropertyType_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			featuretimeslicemetadatapropertytypeIdent := featuretimeslicemetadatapropertytype.GongGetIdentifier(stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&aixmabstractgmlobjecttypes.FeatureTimeSliceMetadataPropertyType{Name: %s}).Stage(stage)", featuretimeslicemetadatapropertytypeIdent, __gong__toRawStringLiteral(featuretimeslicemetadatapropertytype.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", featuretimeslicemetadatapropertytypeIdent, __gong__toRawStringLiteral(featuretimeslicemetadatapropertytype.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Owns = %t", featuretimeslicemetadatapropertytypeIdent, featuretimeslicemetadatapropertytype.Owns))
		}
	}
	if stage != nil {
		for _, messagemetadatapropertytype := range __gong__sortStageSetInstances(stage.MessageMetadataPropertyTypes, stage.MessageMetadataPropertyType_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			messagemetadatapropertytypeIdent := messagemetadatapropertytype.GongGetIdentifier(stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&aixmabstractgmlobjecttypes.MessageMetadataPropertyType{Name: %s}).Stage(stage)", messagemetadatapropertytypeIdent, __gong__toRawStringLiteral(messagemetadatapropertytype.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", messagemetadatapropertytypeIdent, __gong__toRawStringLiteral(messagemetadatapropertytype.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Owns = %t", messagemetadatapropertytypeIdent, messagemetadatapropertytype.Owns))
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
go run github.com/fullstack-lang/gong/app/xsd/tests/aixm/procedure/go/models/aixmabstractgmlobjecttypes/probe/cmd/probe@%s stage.go
*/
package %s

import (
	"slices"
	"time"

	"github.com/fullstack-lang/gong/app/xsd/tests/aixm/procedure/go/models/aixmabstractgmlobjecttypes"%s
)

var (
	_ time.Time
	_ = slices.Index[[]int, int]

	_ *aixmabstractgmlobjecttypes.Stage
)%s

// function will stage objects across all coordinated stages
func _(stage *aixmabstractgmlobjecttypes.Stage) {

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

	if modelsPackageName != "" && modelsPackageName != "github.com/fullstack-lang/gong/app/xsd/tests/aixm/procedure/go/models/aixmabstractgmlobjecttypes" {
		res = strings.ReplaceAll(res, "\"github.com/fullstack-lang/gong/app/xsd/tests/aixm/procedure/go/models/aixmabstractgmlobjecttypes\"", fmt.Sprintf("%q", modelsPackageName))
	}

	return res, nil
}

// insertion point for marshall field methods
func (abstractaixmfeaturebasetype *AbstractAIXMFeatureBaseType) GongMarshallField(stage *Stage, fieldName string) (res string) {
	ident := abstractaixmfeaturebasetype.GongGetIdentifier(stage)
	_ = ident
	switch fieldName {
	case "Name":
		res = __gong__marshallString(ident, "Name", abstractaixmfeaturebasetype.Name)
	case "Id":
		res = __gong__marshallString(ident, "Id", abstractaixmfeaturebasetype.Id)
	case "Description":
		res = __gong__marshallString(ident, "Description", abstractaixmfeaturebasetype.Description)

	default:
		log.Panicf("Unknown field %s for Gongstruct AbstractAIXMFeatureBaseType", fieldName)
	}
	return
}

func (abstractaixmfeaturetype *AbstractAIXMFeatureType) GongMarshallField(stage *Stage, fieldName string) (res string) {
	ident := abstractaixmfeaturetype.GongGetIdentifier(stage)
	_ = ident
	switch fieldName {
	case "Name":
		res = __gong__marshallString(ident, "Name", abstractaixmfeaturetype.Name)
	case "Id":
		res = __gong__marshallString(ident, "Id", abstractaixmfeaturetype.Id)
	case "Description":
		res = __gong__marshallString(ident, "Description", abstractaixmfeaturetype.Description)

	case "FeatureMetadata":
		if abstractaixmfeaturetype.FeatureMetadata != nil {
			res = __gong__marshallPointer(ident, "FeatureMetadata", abstractaixmfeaturetype.FeatureMetadata.GongGetIdentifier(stage))
		} else {
			res = __gong__marshallPointer(ident, "FeatureMetadata", "nil")
		}
	default:
		log.Panicf("Unknown field %s for Gongstruct AbstractAIXMFeatureType", fieldName)
	}
	return
}

func (abstractaixmmessagebasetype *AbstractAIXMMessageBaseType) GongMarshallField(stage *Stage, fieldName string) (res string) {
	ident := abstractaixmmessagebasetype.GongGetIdentifier(stage)
	_ = ident
	switch fieldName {
	case "Name":
		res = __gong__marshallString(ident, "Name", abstractaixmmessagebasetype.Name)
	case "Id":
		res = __gong__marshallString(ident, "Id", abstractaixmmessagebasetype.Id)
	case "Description":
		res = __gong__marshallString(ident, "Description", abstractaixmmessagebasetype.Description)
	case "SequenceNumber":
		res = __gong__marshallInt(ident, "SequenceNumber", abstractaixmmessagebasetype.SequenceNumber)

	case "MessageMetadata":
		if abstractaixmmessagebasetype.MessageMetadata != nil {
			res = __gong__marshallPointer(ident, "MessageMetadata", abstractaixmmessagebasetype.MessageMetadata.GongGetIdentifier(stage))
		} else {
			res = __gong__marshallPointer(ident, "MessageMetadata", "nil")
		}
	default:
		log.Panicf("Unknown field %s for Gongstruct AbstractAIXMMessageBaseType", fieldName)
	}
	return
}

func (abstractaixmmessagetype *AbstractAIXMMessageType) GongMarshallField(stage *Stage, fieldName string) (res string) {
	ident := abstractaixmmessagetype.GongGetIdentifier(stage)
	_ = ident
	switch fieldName {
	case "Name":
		res = __gong__marshallString(ident, "Name", abstractaixmmessagetype.Name)
	case "Id":
		res = __gong__marshallString(ident, "Id", abstractaixmmessagetype.Id)
	case "Description":
		res = __gong__marshallString(ident, "Description", abstractaixmmessagetype.Description)
	case "SequenceNumber":
		res = __gong__marshallInt(ident, "SequenceNumber", abstractaixmmessagetype.SequenceNumber)
	case "AggregationType":
		res = __gong__marshallString(ident, "AggregationType", abstractaixmmessagetype.AggregationType)

	case "MessageMetadata":
		if abstractaixmmessagetype.MessageMetadata != nil {
			res = __gong__marshallPointer(ident, "MessageMetadata", abstractaixmmessagetype.MessageMetadata.GongGetIdentifier(stage))
		} else {
			res = __gong__marshallPointer(ident, "MessageMetadata", "nil")
		}
	default:
		log.Panicf("Unknown field %s for Gongstruct AbstractAIXMMessageType", fieldName)
	}
	return
}

func (abstractaixmobjecttype *AbstractAIXMObjectType) GongMarshallField(stage *Stage, fieldName string) (res string) {
	ident := abstractaixmobjecttype.GongGetIdentifier(stage)
	_ = ident
	switch fieldName {
	case "Name":
		res = __gong__marshallString(ident, "Name", abstractaixmobjecttype.Name)
	case "Id":
		res = __gong__marshallString(ident, "Id", abstractaixmobjecttype.Id)

	default:
		log.Panicf("Unknown field %s for Gongstruct AbstractAIXMObjectType", fieldName)
	}
	return
}

func (abstractaixmpropertytype *AbstractAIXMPropertyType) GongMarshallField(stage *Stage, fieldName string) (res string) {
	ident := abstractaixmpropertytype.GongGetIdentifier(stage)
	_ = ident
	switch fieldName {
	case "Name":
		res = __gong__marshallString(ident, "Name", abstractaixmpropertytype.Name)
	case "NilReason":
		res = __gong__marshallString(ident, "NilReason", abstractaixmpropertytype.NilReason)

	default:
		log.Panicf("Unknown field %s for Gongstruct AbstractAIXMPropertyType", fieldName)
	}
	return
}

func (abstractaixmtimeslicebasetype *AbstractAIXMTimeSliceBaseType) GongMarshallField(stage *Stage, fieldName string) (res string) {
	ident := abstractaixmtimeslicebasetype.GongGetIdentifier(stage)
	_ = ident
	switch fieldName {
	case "Name":
		res = __gong__marshallString(ident, "Name", abstractaixmtimeslicebasetype.Name)
	case "Id":
		res = __gong__marshallString(ident, "Id", abstractaixmtimeslicebasetype.Id)

	default:
		log.Panicf("Unknown field %s for Gongstruct AbstractAIXMTimeSliceBaseType", fieldName)
	}
	return
}

func (abstractaixmtimeslicetype *AbstractAIXMTimeSliceType) GongMarshallField(stage *Stage, fieldName string) (res string) {
	ident := abstractaixmtimeslicetype.GongGetIdentifier(stage)
	_ = ident
	switch fieldName {
	case "Name":
		res = __gong__marshallString(ident, "Name", abstractaixmtimeslicetype.Name)
	case "Id":
		res = __gong__marshallString(ident, "Id", abstractaixmtimeslicetype.Id)
	case "Description":
		res = __gong__marshallString(ident, "Description", abstractaixmtimeslicetype.Description)
	case "Interpretation":
		res = __gong__marshallEnumString(ident, "Interpretation", abstractaixmtimeslicetype.Interpretation.ToCodeString())
	case "SequenceNumber":
		res = __gong__marshallInt(ident, "SequenceNumber", abstractaixmtimeslicetype.SequenceNumber)
	case "CorrectionNumber":
		res = __gong__marshallInt(ident, "CorrectionNumber", abstractaixmtimeslicetype.CorrectionNumber)

	case "TimeSliceMetadata":
		if abstractaixmtimeslicetype.TimeSliceMetadata != nil {
			res = __gong__marshallPointer(ident, "TimeSliceMetadata", abstractaixmtimeslicetype.TimeSliceMetadata.GongGetIdentifier(stage))
		} else {
			res = __gong__marshallPointer(ident, "TimeSliceMetadata", "nil")
		}
	default:
		log.Panicf("Unknown field %s for Gongstruct AbstractAIXMTimeSliceType", fieldName)
	}
	return
}

func (abstractextensiontype *AbstractExtensionType) GongMarshallField(stage *Stage, fieldName string) (res string) {
	ident := abstractextensiontype.GongGetIdentifier(stage)
	_ = ident
	switch fieldName {
	case "Name":
		res = __gong__marshallString(ident, "Name", abstractextensiontype.Name)
	case "Id":
		res = __gong__marshallString(ident, "Id", abstractextensiontype.Id)
	case "Description":
		res = __gong__marshallString(ident, "Description", abstractextensiontype.Description)

	default:
		log.Panicf("Unknown field %s for Gongstruct AbstractExtensionType", fieldName)
	}
	return
}

func (featuremetadatapropertytype *FeatureMetadataPropertyType) GongMarshallField(stage *Stage, fieldName string) (res string) {
	ident := featuremetadatapropertytype.GongGetIdentifier(stage)
	_ = ident
	switch fieldName {
	case "Name":
		res = __gong__marshallString(ident, "Name", featuremetadatapropertytype.Name)
	case "Owns":
		res = __gong__marshallBool(ident, "Owns", featuremetadatapropertytype.Owns)

	default:
		log.Panicf("Unknown field %s for Gongstruct FeatureMetadataPropertyType", fieldName)
	}
	return
}

func (featuretimeslicemetadatapropertytype *FeatureTimeSliceMetadataPropertyType) GongMarshallField(stage *Stage, fieldName string) (res string) {
	ident := featuretimeslicemetadatapropertytype.GongGetIdentifier(stage)
	_ = ident
	switch fieldName {
	case "Name":
		res = __gong__marshallString(ident, "Name", featuretimeslicemetadatapropertytype.Name)
	case "Owns":
		res = __gong__marshallBool(ident, "Owns", featuretimeslicemetadatapropertytype.Owns)

	default:
		log.Panicf("Unknown field %s for Gongstruct FeatureTimeSliceMetadataPropertyType", fieldName)
	}
	return
}

func (messagemetadatapropertytype *MessageMetadataPropertyType) GongMarshallField(stage *Stage, fieldName string) (res string) {
	ident := messagemetadatapropertytype.GongGetIdentifier(stage)
	_ = ident
	switch fieldName {
	case "Name":
		res = __gong__marshallString(ident, "Name", messagemetadatapropertytype.Name)
	case "Owns":
		res = __gong__marshallBool(ident, "Owns", messagemetadatapropertytype.Owns)

	default:
		log.Panicf("Unknown field %s for Gongstruct MessageMetadataPropertyType", fieldName)
	}
	return
}

// insertion point for marshall all fields methods
func (abstractaixmfeaturebasetype *AbstractAIXMFeatureBaseType) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(abstractaixmfeaturebasetype.GongMarshallField(stage, "Name"))
		initializerStatements.WriteString(abstractaixmfeaturebasetype.GongMarshallField(stage, "Id"))
		initializerStatements.WriteString(abstractaixmfeaturebasetype.GongMarshallField(stage, "Description"))
	}
	initRes = initializerStatements.String()
	ptrRes = pointersInitializesStatements.String()
	return
}
func (abstractaixmfeaturetype *AbstractAIXMFeatureType) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(abstractaixmfeaturetype.GongMarshallField(stage, "Name"))
		initializerStatements.WriteString(abstractaixmfeaturetype.GongMarshallField(stage, "Id"))
		initializerStatements.WriteString(abstractaixmfeaturetype.GongMarshallField(stage, "Description"))
		pointersInitializesStatements.WriteString(abstractaixmfeaturetype.GongMarshallField(stage, "FeatureMetadata"))
	}
	initRes = initializerStatements.String()
	ptrRes = pointersInitializesStatements.String()
	return
}
func (abstractaixmmessagebasetype *AbstractAIXMMessageBaseType) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(abstractaixmmessagebasetype.GongMarshallField(stage, "Name"))
		initializerStatements.WriteString(abstractaixmmessagebasetype.GongMarshallField(stage, "Id"))
		initializerStatements.WriteString(abstractaixmmessagebasetype.GongMarshallField(stage, "Description"))
		initializerStatements.WriteString(abstractaixmmessagebasetype.GongMarshallField(stage, "SequenceNumber"))
		pointersInitializesStatements.WriteString(abstractaixmmessagebasetype.GongMarshallField(stage, "MessageMetadata"))
	}
	initRes = initializerStatements.String()
	ptrRes = pointersInitializesStatements.String()
	return
}
func (abstractaixmmessagetype *AbstractAIXMMessageType) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(abstractaixmmessagetype.GongMarshallField(stage, "Name"))
		initializerStatements.WriteString(abstractaixmmessagetype.GongMarshallField(stage, "Id"))
		initializerStatements.WriteString(abstractaixmmessagetype.GongMarshallField(stage, "Description"))
		initializerStatements.WriteString(abstractaixmmessagetype.GongMarshallField(stage, "SequenceNumber"))
		pointersInitializesStatements.WriteString(abstractaixmmessagetype.GongMarshallField(stage, "MessageMetadata"))
		initializerStatements.WriteString(abstractaixmmessagetype.GongMarshallField(stage, "AggregationType"))
	}
	initRes = initializerStatements.String()
	ptrRes = pointersInitializesStatements.String()
	return
}
func (abstractaixmobjecttype *AbstractAIXMObjectType) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(abstractaixmobjecttype.GongMarshallField(stage, "Name"))
		initializerStatements.WriteString(abstractaixmobjecttype.GongMarshallField(stage, "Id"))
	}
	initRes = initializerStatements.String()
	ptrRes = pointersInitializesStatements.String()
	return
}
func (abstractaixmpropertytype *AbstractAIXMPropertyType) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(abstractaixmpropertytype.GongMarshallField(stage, "Name"))
		initializerStatements.WriteString(abstractaixmpropertytype.GongMarshallField(stage, "NilReason"))
	}
	initRes = initializerStatements.String()
	ptrRes = pointersInitializesStatements.String()
	return
}
func (abstractaixmtimeslicebasetype *AbstractAIXMTimeSliceBaseType) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(abstractaixmtimeslicebasetype.GongMarshallField(stage, "Name"))
		initializerStatements.WriteString(abstractaixmtimeslicebasetype.GongMarshallField(stage, "Id"))
	}
	initRes = initializerStatements.String()
	ptrRes = pointersInitializesStatements.String()
	return
}
func (abstractaixmtimeslicetype *AbstractAIXMTimeSliceType) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(abstractaixmtimeslicetype.GongMarshallField(stage, "Name"))
		initializerStatements.WriteString(abstractaixmtimeslicetype.GongMarshallField(stage, "Id"))
		initializerStatements.WriteString(abstractaixmtimeslicetype.GongMarshallField(stage, "Description"))
		initializerStatements.WriteString(abstractaixmtimeslicetype.GongMarshallField(stage, "Interpretation"))
		initializerStatements.WriteString(abstractaixmtimeslicetype.GongMarshallField(stage, "SequenceNumber"))
		initializerStatements.WriteString(abstractaixmtimeslicetype.GongMarshallField(stage, "CorrectionNumber"))
		pointersInitializesStatements.WriteString(abstractaixmtimeslicetype.GongMarshallField(stage, "TimeSliceMetadata"))
	}
	initRes = initializerStatements.String()
	ptrRes = pointersInitializesStatements.String()
	return
}
func (abstractextensiontype *AbstractExtensionType) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(abstractextensiontype.GongMarshallField(stage, "Name"))
		initializerStatements.WriteString(abstractextensiontype.GongMarshallField(stage, "Id"))
		initializerStatements.WriteString(abstractextensiontype.GongMarshallField(stage, "Description"))
	}
	initRes = initializerStatements.String()
	ptrRes = pointersInitializesStatements.String()
	return
}
func (featuremetadatapropertytype *FeatureMetadataPropertyType) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(featuremetadatapropertytype.GongMarshallField(stage, "Name"))
		initializerStatements.WriteString(featuremetadatapropertytype.GongMarshallField(stage, "Owns"))
	}
	initRes = initializerStatements.String()
	ptrRes = pointersInitializesStatements.String()
	return
}
func (featuretimeslicemetadatapropertytype *FeatureTimeSliceMetadataPropertyType) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(featuretimeslicemetadatapropertytype.GongMarshallField(stage, "Name"))
		initializerStatements.WriteString(featuretimeslicemetadatapropertytype.GongMarshallField(stage, "Owns"))
	}
	initRes = initializerStatements.String()
	ptrRes = pointersInitializesStatements.String()
	return
}
func (messagemetadatapropertytype *MessageMetadataPropertyType) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(messagemetadatapropertytype.GongMarshallField(stage, "Name"))
		initializerStatements.WriteString(messagemetadatapropertytype.GongMarshallField(stage, "Owns"))
	}
	initRes = initializerStatements.String()
	ptrRes = pointersInitializesStatements.String()
	return
}
