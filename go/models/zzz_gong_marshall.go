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
	modelsPackageName := "github.com/fullstack-lang/gong/go/models"
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
		for _, gongbasicfield := range __gong__sortStageSetInstances(stage.GongBasicFields, stage.GongBasicField_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			gongbasicfieldIdent := gongbasicfield.GongGetIdentifier(stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&models.GongBasicField{Name: %s}).Stage(stage)", gongbasicfieldIdent, __gong__toRawStringLiteral(gongbasicfield.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", gongbasicfieldIdent, __gong__toRawStringLiteral(gongbasicfield.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.BasicKindName = %s", gongbasicfieldIdent, __gong__toRawStringLiteral(gongbasicfield.BasicKindName)))
			values.WriteString(fmt.Sprintf("\n\t%s.DeclaredType = %s", gongbasicfieldIdent, __gong__toRawStringLiteral(gongbasicfield.DeclaredType)))
			values.WriteString(fmt.Sprintf("\n\t%s.CompositeStructName = %s", gongbasicfieldIdent, __gong__toRawStringLiteral(gongbasicfield.CompositeStructName)))
			values.WriteString(fmt.Sprintf("\n\t%s.IsAccordionStart = %t", gongbasicfieldIdent, gongbasicfield.IsAccordionStart))
			values.WriteString(fmt.Sprintf("\n\t%s.AccordionName = %s", gongbasicfieldIdent, __gong__toRawStringLiteral(gongbasicfield.AccordionName)))
			values.WriteString(fmt.Sprintf("\n\t%s.IsAccordionEnd = %t", gongbasicfieldIdent, gongbasicfield.IsAccordionEnd))
			values.WriteString(fmt.Sprintf("\n\t%s.Index = %d", gongbasicfieldIdent, gongbasicfield.Index))
			values.WriteString(fmt.Sprintf("\n\t%s.IsTextArea = %t", gongbasicfieldIdent, gongbasicfield.IsTextArea))
			values.WriteString(fmt.Sprintf("\n\t%s.IsBespokeWidth = %t", gongbasicfieldIdent, gongbasicfield.IsBespokeWidth))
			values.WriteString(fmt.Sprintf("\n\t%s.BespokeWidth = %d", gongbasicfieldIdent, gongbasicfield.BespokeWidth))
			values.WriteString(fmt.Sprintf("\n\t%s.IsBespokeHeight = %t", gongbasicfieldIdent, gongbasicfield.IsBespokeHeight))
			values.WriteString(fmt.Sprintf("\n\t%s.BespokeHeight = %d", gongbasicfieldIdent, gongbasicfield.BespokeHeight))
			if gongbasicfield.GongEnum != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := gongbasicfield.GongEnum.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.GongEnum = %s", gongbasicfieldIdent, targetIdent))
			}
		}
	}
	if stage != nil {
		for _, gongenum := range __gong__sortStageSetInstances(stage.GongEnums, stage.GongEnum_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			gongenumIdent := gongenum.GongGetIdentifier(stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&models.GongEnum{Name: %s}).Stage(stage)", gongenumIdent, __gong__toRawStringLiteral(gongenum.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", gongenumIdent, __gong__toRawStringLiteral(gongenum.Name)))
			if gongenum.Type.ToCodeString() != "" {
				values.WriteString(fmt.Sprintf("\n\t%s.Type = models.%s", gongenumIdent, gongenum.Type.ToCodeString()))
			} else {
				values.WriteString(fmt.Sprintf("\n\t%s.Type = 0", gongenumIdent))
			}
			for _, elem := range gongenum.GongEnumValues {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := elem.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.GongEnumValues = append(%s.GongEnumValues, %s)", gongenumIdent, gongenumIdent, targetIdent))
			}
		}
	}
	if stage != nil {
		for _, gongenumvalue := range __gong__sortStageSetInstances(stage.GongEnumValues, stage.GongEnumValue_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			gongenumvalueIdent := gongenumvalue.GongGetIdentifier(stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&models.GongEnumValue{Name: %s}).Stage(stage)", gongenumvalueIdent, __gong__toRawStringLiteral(gongenumvalue.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", gongenumvalueIdent, __gong__toRawStringLiteral(gongenumvalue.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Value = %s", gongenumvalueIdent, __gong__toRawStringLiteral(gongenumvalue.Value)))
		}
	}
	if stage != nil {
		for _, gonglink := range __gong__sortStageSetInstances(stage.GongLinks, stage.GongLink_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			gonglinkIdent := gonglink.GongGetIdentifier(stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&models.GongLink{Name: %s}).Stage(stage)", gonglinkIdent, __gong__toRawStringLiteral(gonglink.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", gonglinkIdent, __gong__toRawStringLiteral(gonglink.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Recv = %s", gonglinkIdent, __gong__toRawStringLiteral(gonglink.Recv)))
			values.WriteString(fmt.Sprintf("\n\t%s.ImportPath = %s", gonglinkIdent, __gong__toRawStringLiteral(gonglink.ImportPath)))
		}
	}
	if stage != nil {
		for _, gongnote := range __gong__sortStageSetInstances(stage.GongNotes, stage.GongNote_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			gongnoteIdent := gongnote.GongGetIdentifier(stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&models.GongNote{Name: %s}).Stage(stage)", gongnoteIdent, __gong__toRawStringLiteral(gongnote.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", gongnoteIdent, __gong__toRawStringLiteral(gongnote.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Body = %s", gongnoteIdent, __gong__toRawStringLiteral(gongnote.Body)))
			values.WriteString(fmt.Sprintf("\n\t%s.BodyHTML = %s", gongnoteIdent, __gong__toRawStringLiteral(gongnote.BodyHTML)))
			for _, elem := range gongnote.Links {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := elem.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Links = append(%s.Links, %s)", gongnoteIdent, gongnoteIdent, targetIdent))
			}
		}
	}
	if stage != nil {
		for _, gongstruct := range __gong__sortStageSetInstances(stage.GongStructs, stage.GongStruct_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			gongstructIdent := gongstruct.GongGetIdentifier(stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&models.GongStruct{Name: %s}).Stage(stage)", gongstructIdent, __gong__toRawStringLiteral(gongstruct.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", gongstructIdent, __gong__toRawStringLiteral(gongstruct.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.HasOnAfterUpdateSignature = %t", gongstructIdent, gongstruct.HasOnAfterUpdateSignature))
			values.WriteString(fmt.Sprintf("\n\t%s.IsIgnoredForFront = %t", gongstructIdent, gongstruct.IsIgnoredForFront))
			values.WriteString(fmt.Sprintf("\n\t%s.IsOmittedForMarshalling = %t", gongstructIdent, gongstruct.IsOmittedForMarshalling))
			for _, elem := range gongstruct.GongBasicFields {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := elem.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.GongBasicFields = append(%s.GongBasicFields, %s)", gongstructIdent, gongstructIdent, targetIdent))
			}
			for _, elem := range gongstruct.GongTimeFields {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := elem.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.GongTimeFields = append(%s.GongTimeFields, %s)", gongstructIdent, gongstructIdent, targetIdent))
			}
			for _, elem := range gongstruct.PointerToGongStructFields {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := elem.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.PointerToGongStructFields = append(%s.PointerToGongStructFields, %s)", gongstructIdent, gongstructIdent, targetIdent))
			}
			for _, elem := range gongstruct.SliceOfPointerToGongStructFields {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := elem.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.SliceOfPointerToGongStructFields = append(%s.SliceOfPointerToGongStructFields, %s)", gongstructIdent, gongstructIdent, targetIdent))
			}
			if gongstruct.ModelPkg != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := gongstruct.ModelPkg.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.ModelPkg = %s", gongstructIdent, targetIdent))
			}
		}
	}
	if stage != nil {
		for _, gongtimefield := range __gong__sortStageSetInstances(stage.GongTimeFields, stage.GongTimeField_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			gongtimefieldIdent := gongtimefield.GongGetIdentifier(stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&models.GongTimeField{Name: %s}).Stage(stage)", gongtimefieldIdent, __gong__toRawStringLiteral(gongtimefield.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", gongtimefieldIdent, __gong__toRawStringLiteral(gongtimefield.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Index = %d", gongtimefieldIdent, gongtimefield.Index))
			values.WriteString(fmt.Sprintf("\n\t%s.CompositeStructName = %s", gongtimefieldIdent, __gong__toRawStringLiteral(gongtimefield.CompositeStructName)))
			values.WriteString(fmt.Sprintf("\n\t%s.IsAccordionStart = %t", gongtimefieldIdent, gongtimefield.IsAccordionStart))
			values.WriteString(fmt.Sprintf("\n\t%s.AccordionName = %s", gongtimefieldIdent, __gong__toRawStringLiteral(gongtimefield.AccordionName)))
			values.WriteString(fmt.Sprintf("\n\t%s.IsAccordionEnd = %t", gongtimefieldIdent, gongtimefield.IsAccordionEnd))
			values.WriteString(fmt.Sprintf("\n\t%s.BespokeTimeFormat = %s", gongtimefieldIdent, __gong__toRawStringLiteral(gongtimefield.BespokeTimeFormat)))
			values.WriteString(fmt.Sprintf("\n\t%s.TimeFormOnly = %t", gongtimefieldIdent, gongtimefield.TimeFormOnly))
		}
	}
	if stage != nil {
		for _, metareference := range __gong__sortStageSetInstances(stage.MetaReferences, stage.MetaReference_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			metareferenceIdent := metareference.GongGetIdentifier(stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&models.MetaReference{Name: %s}).Stage(stage)", metareferenceIdent, __gong__toRawStringLiteral(metareference.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", metareferenceIdent, __gong__toRawStringLiteral(metareference.Name)))
		}
	}
	if stage != nil {
		for _, modelpkg := range __gong__sortStageSetInstances(stage.ModelPkgs, stage.ModelPkg_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			modelpkgIdent := modelpkg.GongGetIdentifier(stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&models.ModelPkg{Name: %s}).Stage(stage)", modelpkgIdent, __gong__toRawStringLiteral(modelpkg.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", modelpkgIdent, __gong__toRawStringLiteral(modelpkg.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.PkgGoName = %s", modelpkgIdent, __gong__toRawStringLiteral(modelpkg.PkgGoName)))
			values.WriteString(fmt.Sprintf("\n\t%s.PkgPath = %s", modelpkgIdent, __gong__toRawStringLiteral(modelpkg.PkgPath)))
			values.WriteString(fmt.Sprintf("\n\t%s.PathToGoSubDirectory = %s", modelpkgIdent, __gong__toRawStringLiteral(modelpkg.PathToGoSubDirectory)))
			values.WriteString(fmt.Sprintf("\n\t%s.OrmPkgGenPath = %s", modelpkgIdent, __gong__toRawStringLiteral(modelpkg.OrmPkgGenPath)))
			values.WriteString(fmt.Sprintf("\n\t%s.DbOrmPkgGenPath = %s", modelpkgIdent, __gong__toRawStringLiteral(modelpkg.DbOrmPkgGenPath)))
			values.WriteString(fmt.Sprintf("\n\t%s.DbLiteOrmPkgGenPath = %s", modelpkgIdent, __gong__toRawStringLiteral(modelpkg.DbLiteOrmPkgGenPath)))
			values.WriteString(fmt.Sprintf("\n\t%s.DbPkgGenPath = %s", modelpkgIdent, __gong__toRawStringLiteral(modelpkg.DbPkgGenPath)))
			values.WriteString(fmt.Sprintf("\n\t%s.ControllersPkgGenPath = %s", modelpkgIdent, __gong__toRawStringLiteral(modelpkg.ControllersPkgGenPath)))
			values.WriteString(fmt.Sprintf("\n\t%s.FullstackPkgGenPath = %s", modelpkgIdent, __gong__toRawStringLiteral(modelpkg.FullstackPkgGenPath)))
			values.WriteString(fmt.Sprintf("\n\t%s.StackPkgGenPath = %s", modelpkgIdent, __gong__toRawStringLiteral(modelpkg.StackPkgGenPath)))
			values.WriteString(fmt.Sprintf("\n\t%s.Level1StackPkgGenPath = %s", modelpkgIdent, __gong__toRawStringLiteral(modelpkg.Level1StackPkgGenPath)))
			values.WriteString(fmt.Sprintf("\n\t%s.StaticPkgGenPath = %s", modelpkgIdent, __gong__toRawStringLiteral(modelpkg.StaticPkgGenPath)))
			values.WriteString(fmt.Sprintf("\n\t%s.ProbePkgGenPath = %s", modelpkgIdent, __gong__toRawStringLiteral(modelpkg.ProbePkgGenPath)))
			values.WriteString(fmt.Sprintf("\n\t%s.NgWorkspacePath = %s", modelpkgIdent, __gong__toRawStringLiteral(modelpkg.NgWorkspacePath)))
			values.WriteString(fmt.Sprintf("\n\t%s.NgWorkspaceName = %s", modelpkgIdent, __gong__toRawStringLiteral(modelpkg.NgWorkspaceName)))
			values.WriteString(fmt.Sprintf("\n\t%s.NgDataLibrarySourceCodeDirectory = %s", modelpkgIdent, __gong__toRawStringLiteral(modelpkg.NgDataLibrarySourceCodeDirectory)))
			values.WriteString(fmt.Sprintf("\n\t%s.NgSpecificLibrarySourceCodeDirectory = %s", modelpkgIdent, __gong__toRawStringLiteral(modelpkg.NgSpecificLibrarySourceCodeDirectory)))
			values.WriteString(fmt.Sprintf("\n\t%s.MaterialLibDatamodelTargetPath = %s", modelpkgIdent, __gong__toRawStringLiteral(modelpkg.MaterialLibDatamodelTargetPath)))
			if modelpkg.StageSet != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := modelpkg.StageSet.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.StageSet = %s", modelpkgIdent, targetIdent))
			}
		}
	}
	if stage != nil {
		for _, pointertogongstructfield := range __gong__sortStageSetInstances(stage.PointerToGongStructFields, stage.PointerToGongStructField_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			pointertogongstructfieldIdent := pointertogongstructfield.GongGetIdentifier(stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&models.PointerToGongStructField{Name: %s}).Stage(stage)", pointertogongstructfieldIdent, __gong__toRawStringLiteral(pointertogongstructfield.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", pointertogongstructfieldIdent, __gong__toRawStringLiteral(pointertogongstructfield.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Index = %d", pointertogongstructfieldIdent, pointertogongstructfield.Index))
			values.WriteString(fmt.Sprintf("\n\t%s.CompositeStructName = %s", pointertogongstructfieldIdent, __gong__toRawStringLiteral(pointertogongstructfield.CompositeStructName)))
			values.WriteString(fmt.Sprintf("\n\t%s.IsAccordionStart = %t", pointertogongstructfieldIdent, pointertogongstructfield.IsAccordionStart))
			values.WriteString(fmt.Sprintf("\n\t%s.AccordionName = %s", pointertogongstructfieldIdent, __gong__toRawStringLiteral(pointertogongstructfield.AccordionName)))
			values.WriteString(fmt.Sprintf("\n\t%s.IsAccordionEnd = %t", pointertogongstructfieldIdent, pointertogongstructfield.IsAccordionEnd))
			values.WriteString(fmt.Sprintf("\n\t%s.IsType = %t", pointertogongstructfieldIdent, pointertogongstructfield.IsType))
			if pointertogongstructfield.GongStruct != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := pointertogongstructfield.GongStruct.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.GongStruct = %s", pointertogongstructfieldIdent, targetIdent))
			}
		}
	}
	if stage != nil {
		for _, sliceofpointertogongstructfield := range __gong__sortStageSetInstances(stage.SliceOfPointerToGongStructFields, stage.SliceOfPointerToGongStructField_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			sliceofpointertogongstructfieldIdent := sliceofpointertogongstructfield.GongGetIdentifier(stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&models.SliceOfPointerToGongStructField{Name: %s}).Stage(stage)", sliceofpointertogongstructfieldIdent, __gong__toRawStringLiteral(sliceofpointertogongstructfield.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", sliceofpointertogongstructfieldIdent, __gong__toRawStringLiteral(sliceofpointertogongstructfield.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Index = %d", sliceofpointertogongstructfieldIdent, sliceofpointertogongstructfield.Index))
			values.WriteString(fmt.Sprintf("\n\t%s.CompositeStructName = %s", sliceofpointertogongstructfieldIdent, __gong__toRawStringLiteral(sliceofpointertogongstructfield.CompositeStructName)))
			values.WriteString(fmt.Sprintf("\n\t%s.IsAccordionStart = %t", sliceofpointertogongstructfieldIdent, sliceofpointertogongstructfield.IsAccordionStart))
			values.WriteString(fmt.Sprintf("\n\t%s.AccordionName = %s", sliceofpointertogongstructfieldIdent, __gong__toRawStringLiteral(sliceofpointertogongstructfield.AccordionName)))
			values.WriteString(fmt.Sprintf("\n\t%s.IsAccordionEnd = %t", sliceofpointertogongstructfieldIdent, sliceofpointertogongstructfield.IsAccordionEnd))
			if sliceofpointertogongstructfield.GongStruct != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := sliceofpointertogongstructfield.GongStruct.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.GongStruct = %s", sliceofpointertogongstructfieldIdent, targetIdent))
			}
		}
	}
	if stage != nil {
		for _, stagesetfield := range __gong__sortStageSetInstances(stage.StageSetFields, stage.StageSetField_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			stagesetfieldIdent := stagesetfield.GongGetIdentifier(stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&models.StageSetField{Name: %s}).Stage(stage)", stagesetfieldIdent, __gong__toRawStringLiteral(stagesetfield.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", stagesetfieldIdent, __gong__toRawStringLiteral(stagesetfield.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.PackageName = %s", stagesetfieldIdent, __gong__toRawStringLiteral(stagesetfield.PackageName)))
			values.WriteString(fmt.Sprintf("\n\t%s.PackagePath = %s", stagesetfieldIdent, __gong__toRawStringLiteral(stagesetfield.PackagePath)))
			values.WriteString(fmt.Sprintf("\n\t%s.IsLocal = %t", stagesetfieldIdent, stagesetfield.IsLocal))
			values.WriteString(fmt.Sprintf("\n\t%s.ImportAlias = %s", stagesetfieldIdent, __gong__toRawStringLiteral(stagesetfield.ImportAlias)))
		}
	}
	if stage != nil {
		for _, stagesetmodel := range __gong__sortStageSetInstances(stage.StageSetModels, stage.StageSetModel_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			stagesetmodelIdent := stagesetmodel.GongGetIdentifier(stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&models.StageSetModel{Name: %s}).Stage(stage)", stagesetmodelIdent, __gong__toRawStringLiteral(stagesetmodel.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", stagesetmodelIdent, __gong__toRawStringLiteral(stagesetmodel.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.IsManual = %t", stagesetmodelIdent, stagesetmodel.IsManual))
			for _, elem := range stagesetmodel.Fields {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := elem.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Fields = append(%s.Fields, %s)", stagesetmodelIdent, stagesetmodelIdent, targetIdent))
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
go run github.com/fullstack-lang/gong/go/models/probe/cmd/probe@%s stage.go
*/
package %s

import (
	"slices"
	"time"

	"github.com/fullstack-lang/gong/go/models"%s
)

var (
	_ time.Time
	_ = slices.Index[[]int, int]

	_ *models.Stage
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

	if modelsPackageName != "" && modelsPackageName != "github.com/fullstack-lang/gong/go/models" {
		res = strings.ReplaceAll(res, "\"github.com/fullstack-lang/gong/go/models\"", fmt.Sprintf("%q", modelsPackageName))
	}

	return res, nil
}

// insertion point for marshall field methods
func (gongbasicfield *GongBasicField) GongMarshallField(stage *Stage, fieldName string) (res string) {
	ident := gongbasicfield.GongGetIdentifier(stage)
	_ = ident
	switch fieldName {
	case "Name":
		res = __gong__marshallString(ident, "Name", gongbasicfield.Name)
	case "BasicKindName":
		res = __gong__marshallString(ident, "BasicKindName", gongbasicfield.BasicKindName)
	case "DeclaredType":
		res = __gong__marshallString(ident, "DeclaredType", gongbasicfield.DeclaredType)
	case "CompositeStructName":
		res = __gong__marshallString(ident, "CompositeStructName", gongbasicfield.CompositeStructName)
	case "IsAccordionStart":
		res = __gong__marshallBool(ident, "IsAccordionStart", gongbasicfield.IsAccordionStart)
	case "AccordionName":
		res = __gong__marshallString(ident, "AccordionName", gongbasicfield.AccordionName)
	case "IsAccordionEnd":
		res = __gong__marshallBool(ident, "IsAccordionEnd", gongbasicfield.IsAccordionEnd)
	case "Index":
		res = __gong__marshallInt(ident, "Index", gongbasicfield.Index)
	case "IsTextArea":
		res = __gong__marshallBool(ident, "IsTextArea", gongbasicfield.IsTextArea)
	case "IsBespokeWidth":
		res = __gong__marshallBool(ident, "IsBespokeWidth", gongbasicfield.IsBespokeWidth)
	case "BespokeWidth":
		res = __gong__marshallInt(ident, "BespokeWidth", gongbasicfield.BespokeWidth)
	case "IsBespokeHeight":
		res = __gong__marshallBool(ident, "IsBespokeHeight", gongbasicfield.IsBespokeHeight)
	case "BespokeHeight":
		res = __gong__marshallInt(ident, "BespokeHeight", gongbasicfield.BespokeHeight)

	case "GongEnum":
		if gongbasicfield.GongEnum != nil {
			res = __gong__marshallPointer(ident, "GongEnum", gongbasicfield.GongEnum.GongGetIdentifier(stage))
		} else {
			res = __gong__marshallPointer(ident, "GongEnum", "nil")
		}
	default:
		log.Panicf("Unknown field %s for Gongstruct GongBasicField", fieldName)
	}
	return
}

func (gongenum *GongEnum) GongMarshallField(stage *Stage, fieldName string) (res string) {
	ident := gongenum.GongGetIdentifier(stage)
	_ = ident
	switch fieldName {
	case "Name":
		res = __gong__marshallString(ident, "Name", gongenum.Name)
	case "Type":
		res = __gong__marshallEnumInt(ident, "Type", gongenum.Type.ToCodeString())

	case "GongEnumValues":
		var sb strings.Builder
		for _, _gongenumvalue := range gongenum.GongEnumValues {
			sb.WriteString(__gong__marshallSliceOfPointers(ident, "GongEnumValues", _gongenumvalue.GongGetIdentifier(stage)))
		}
		res = sb.String()
	default:
		log.Panicf("Unknown field %s for Gongstruct GongEnum", fieldName)
	}
	return
}

func (gongenumvalue *GongEnumValue) GongMarshallField(stage *Stage, fieldName string) (res string) {
	ident := gongenumvalue.GongGetIdentifier(stage)
	_ = ident
	switch fieldName {
	case "Name":
		res = __gong__marshallString(ident, "Name", gongenumvalue.Name)
	case "Value":
		res = __gong__marshallString(ident, "Value", gongenumvalue.Value)

	default:
		log.Panicf("Unknown field %s for Gongstruct GongEnumValue", fieldName)
	}
	return
}

func (gonglink *GongLink) GongMarshallField(stage *Stage, fieldName string) (res string) {
	ident := gonglink.GongGetIdentifier(stage)
	_ = ident
	switch fieldName {
	case "Name":
		res = __gong__marshallString(ident, "Name", gonglink.Name)
	case "Recv":
		res = __gong__marshallString(ident, "Recv", gonglink.Recv)
	case "ImportPath":
		res = __gong__marshallString(ident, "ImportPath", gonglink.ImportPath)

	default:
		log.Panicf("Unknown field %s for Gongstruct GongLink", fieldName)
	}
	return
}

func (gongnote *GongNote) GongMarshallField(stage *Stage, fieldName string) (res string) {
	ident := gongnote.GongGetIdentifier(stage)
	_ = ident
	switch fieldName {
	case "Name":
		res = __gong__marshallString(ident, "Name", gongnote.Name)
	case "Body":
		res = __gong__marshallString(ident, "Body", gongnote.Body)
	case "BodyHTML":
		res = __gong__marshallString(ident, "BodyHTML", gongnote.BodyHTML)

	case "Links":
		var sb strings.Builder
		for _, _gonglink := range gongnote.Links {
			sb.WriteString(__gong__marshallSliceOfPointers(ident, "Links", _gonglink.GongGetIdentifier(stage)))
		}
		res = sb.String()
	default:
		log.Panicf("Unknown field %s for Gongstruct GongNote", fieldName)
	}
	return
}

func (gongstruct *GongStruct) GongMarshallField(stage *Stage, fieldName string) (res string) {
	ident := gongstruct.GongGetIdentifier(stage)
	_ = ident
	switch fieldName {
	case "Name":
		res = __gong__marshallString(ident, "Name", gongstruct.Name)
	case "HasOnAfterUpdateSignature":
		res = __gong__marshallBool(ident, "HasOnAfterUpdateSignature", gongstruct.HasOnAfterUpdateSignature)
	case "IsIgnoredForFront":
		res = __gong__marshallBool(ident, "IsIgnoredForFront", gongstruct.IsIgnoredForFront)
	case "IsOmittedForMarshalling":
		res = __gong__marshallBool(ident, "IsOmittedForMarshalling", gongstruct.IsOmittedForMarshalling)

	case "GongBasicFields":
		var sb strings.Builder
		for _, _gongbasicfield := range gongstruct.GongBasicFields {
			sb.WriteString(__gong__marshallSliceOfPointers(ident, "GongBasicFields", _gongbasicfield.GongGetIdentifier(stage)))
		}
		res = sb.String()
	case "GongTimeFields":
		var sb strings.Builder
		for _, _gongtimefield := range gongstruct.GongTimeFields {
			sb.WriteString(__gong__marshallSliceOfPointers(ident, "GongTimeFields", _gongtimefield.GongGetIdentifier(stage)))
		}
		res = sb.String()
	case "PointerToGongStructFields":
		var sb strings.Builder
		for _, _pointertogongstructfield := range gongstruct.PointerToGongStructFields {
			sb.WriteString(__gong__marshallSliceOfPointers(ident, "PointerToGongStructFields", _pointertogongstructfield.GongGetIdentifier(stage)))
		}
		res = sb.String()
	case "SliceOfPointerToGongStructFields":
		var sb strings.Builder
		for _, _sliceofpointertogongstructfield := range gongstruct.SliceOfPointerToGongStructFields {
			sb.WriteString(__gong__marshallSliceOfPointers(ident, "SliceOfPointerToGongStructFields", _sliceofpointertogongstructfield.GongGetIdentifier(stage)))
		}
		res = sb.String()
	default:
		log.Panicf("Unknown field %s for Gongstruct GongStruct", fieldName)
	}
	return
}

func (gongtimefield *GongTimeField) GongMarshallField(stage *Stage, fieldName string) (res string) {
	ident := gongtimefield.GongGetIdentifier(stage)
	_ = ident
	switch fieldName {
	case "Name":
		res = __gong__marshallString(ident, "Name", gongtimefield.Name)
	case "Index":
		res = __gong__marshallInt(ident, "Index", gongtimefield.Index)
	case "CompositeStructName":
		res = __gong__marshallString(ident, "CompositeStructName", gongtimefield.CompositeStructName)
	case "IsAccordionStart":
		res = __gong__marshallBool(ident, "IsAccordionStart", gongtimefield.IsAccordionStart)
	case "AccordionName":
		res = __gong__marshallString(ident, "AccordionName", gongtimefield.AccordionName)
	case "IsAccordionEnd":
		res = __gong__marshallBool(ident, "IsAccordionEnd", gongtimefield.IsAccordionEnd)
	case "BespokeTimeFormat":
		res = __gong__marshallString(ident, "BespokeTimeFormat", gongtimefield.BespokeTimeFormat)
	case "TimeFormOnly":
		res = __gong__marshallBool(ident, "TimeFormOnly", gongtimefield.TimeFormOnly)

	default:
		log.Panicf("Unknown field %s for Gongstruct GongTimeField", fieldName)
	}
	return
}

func (metareference *MetaReference) GongMarshallField(stage *Stage, fieldName string) (res string) {
	ident := metareference.GongGetIdentifier(stage)
	_ = ident
	switch fieldName {
	case "Name":
		res = __gong__marshallString(ident, "Name", metareference.Name)

	default:
		log.Panicf("Unknown field %s for Gongstruct MetaReference", fieldName)
	}
	return
}

func (modelpkg *ModelPkg) GongMarshallField(stage *Stage, fieldName string) (res string) {
	ident := modelpkg.GongGetIdentifier(stage)
	_ = ident
	switch fieldName {
	case "Name":
		res = __gong__marshallString(ident, "Name", modelpkg.Name)
	case "PkgGoName":
		res = __gong__marshallString(ident, "PkgGoName", modelpkg.PkgGoName)
	case "PkgPath":
		res = __gong__marshallString(ident, "PkgPath", modelpkg.PkgPath)
	case "PathToGoSubDirectory":
		res = __gong__marshallString(ident, "PathToGoSubDirectory", modelpkg.PathToGoSubDirectory)
	case "OrmPkgGenPath":
		res = __gong__marshallString(ident, "OrmPkgGenPath", modelpkg.OrmPkgGenPath)
	case "DbOrmPkgGenPath":
		res = __gong__marshallString(ident, "DbOrmPkgGenPath", modelpkg.DbOrmPkgGenPath)
	case "DbLiteOrmPkgGenPath":
		res = __gong__marshallString(ident, "DbLiteOrmPkgGenPath", modelpkg.DbLiteOrmPkgGenPath)
	case "DbPkgGenPath":
		res = __gong__marshallString(ident, "DbPkgGenPath", modelpkg.DbPkgGenPath)
	case "ControllersPkgGenPath":
		res = __gong__marshallString(ident, "ControllersPkgGenPath", modelpkg.ControllersPkgGenPath)
	case "FullstackPkgGenPath":
		res = __gong__marshallString(ident, "FullstackPkgGenPath", modelpkg.FullstackPkgGenPath)
	case "StackPkgGenPath":
		res = __gong__marshallString(ident, "StackPkgGenPath", modelpkg.StackPkgGenPath)
	case "Level1StackPkgGenPath":
		res = __gong__marshallString(ident, "Level1StackPkgGenPath", modelpkg.Level1StackPkgGenPath)
	case "StaticPkgGenPath":
		res = __gong__marshallString(ident, "StaticPkgGenPath", modelpkg.StaticPkgGenPath)
	case "ProbePkgGenPath":
		res = __gong__marshallString(ident, "ProbePkgGenPath", modelpkg.ProbePkgGenPath)
	case "NgWorkspacePath":
		res = __gong__marshallString(ident, "NgWorkspacePath", modelpkg.NgWorkspacePath)
	case "NgWorkspaceName":
		res = __gong__marshallString(ident, "NgWorkspaceName", modelpkg.NgWorkspaceName)
	case "NgDataLibrarySourceCodeDirectory":
		res = __gong__marshallString(ident, "NgDataLibrarySourceCodeDirectory", modelpkg.NgDataLibrarySourceCodeDirectory)
	case "NgSpecificLibrarySourceCodeDirectory":
		res = __gong__marshallString(ident, "NgSpecificLibrarySourceCodeDirectory", modelpkg.NgSpecificLibrarySourceCodeDirectory)
	case "MaterialLibDatamodelTargetPath":
		res = __gong__marshallString(ident, "MaterialLibDatamodelTargetPath", modelpkg.MaterialLibDatamodelTargetPath)

	default:
		log.Panicf("Unknown field %s for Gongstruct ModelPkg", fieldName)
	}
	return
}

func (pointertogongstructfield *PointerToGongStructField) GongMarshallField(stage *Stage, fieldName string) (res string) {
	ident := pointertogongstructfield.GongGetIdentifier(stage)
	_ = ident
	switch fieldName {
	case "Name":
		res = __gong__marshallString(ident, "Name", pointertogongstructfield.Name)
	case "Index":
		res = __gong__marshallInt(ident, "Index", pointertogongstructfield.Index)
	case "CompositeStructName":
		res = __gong__marshallString(ident, "CompositeStructName", pointertogongstructfield.CompositeStructName)
	case "IsAccordionStart":
		res = __gong__marshallBool(ident, "IsAccordionStart", pointertogongstructfield.IsAccordionStart)
	case "AccordionName":
		res = __gong__marshallString(ident, "AccordionName", pointertogongstructfield.AccordionName)
	case "IsAccordionEnd":
		res = __gong__marshallBool(ident, "IsAccordionEnd", pointertogongstructfield.IsAccordionEnd)
	case "IsType":
		res = __gong__marshallBool(ident, "IsType", pointertogongstructfield.IsType)

	case "GongStruct":
		if pointertogongstructfield.GongStruct != nil {
			res = __gong__marshallPointer(ident, "GongStruct", pointertogongstructfield.GongStruct.GongGetIdentifier(stage))
		} else {
			res = __gong__marshallPointer(ident, "GongStruct", "nil")
		}
	default:
		log.Panicf("Unknown field %s for Gongstruct PointerToGongStructField", fieldName)
	}
	return
}

func (sliceofpointertogongstructfield *SliceOfPointerToGongStructField) GongMarshallField(stage *Stage, fieldName string) (res string) {
	ident := sliceofpointertogongstructfield.GongGetIdentifier(stage)
	_ = ident
	switch fieldName {
	case "Name":
		res = __gong__marshallString(ident, "Name", sliceofpointertogongstructfield.Name)
	case "Index":
		res = __gong__marshallInt(ident, "Index", sliceofpointertogongstructfield.Index)
	case "CompositeStructName":
		res = __gong__marshallString(ident, "CompositeStructName", sliceofpointertogongstructfield.CompositeStructName)
	case "IsAccordionStart":
		res = __gong__marshallBool(ident, "IsAccordionStart", sliceofpointertogongstructfield.IsAccordionStart)
	case "AccordionName":
		res = __gong__marshallString(ident, "AccordionName", sliceofpointertogongstructfield.AccordionName)
	case "IsAccordionEnd":
		res = __gong__marshallBool(ident, "IsAccordionEnd", sliceofpointertogongstructfield.IsAccordionEnd)

	case "GongStruct":
		if sliceofpointertogongstructfield.GongStruct != nil {
			res = __gong__marshallPointer(ident, "GongStruct", sliceofpointertogongstructfield.GongStruct.GongGetIdentifier(stage))
		} else {
			res = __gong__marshallPointer(ident, "GongStruct", "nil")
		}
	default:
		log.Panicf("Unknown field %s for Gongstruct SliceOfPointerToGongStructField", fieldName)
	}
	return
}

func (stagesetfield *StageSetField) GongMarshallField(stage *Stage, fieldName string) (res string) {
	ident := stagesetfield.GongGetIdentifier(stage)
	_ = ident
	switch fieldName {
	case "Name":
		res = __gong__marshallString(ident, "Name", stagesetfield.Name)
	case "PackageName":
		res = __gong__marshallString(ident, "PackageName", stagesetfield.PackageName)
	case "PackagePath":
		res = __gong__marshallString(ident, "PackagePath", stagesetfield.PackagePath)
	case "IsLocal":
		res = __gong__marshallBool(ident, "IsLocal", stagesetfield.IsLocal)
	case "ImportAlias":
		res = __gong__marshallString(ident, "ImportAlias", stagesetfield.ImportAlias)

	default:
		log.Panicf("Unknown field %s for Gongstruct StageSetField", fieldName)
	}
	return
}

func (stagesetmodel *StageSetModel) GongMarshallField(stage *Stage, fieldName string) (res string) {
	ident := stagesetmodel.GongGetIdentifier(stage)
	_ = ident
	switch fieldName {
	case "Name":
		res = __gong__marshallString(ident, "Name", stagesetmodel.Name)
	case "IsManual":
		res = __gong__marshallBool(ident, "IsManual", stagesetmodel.IsManual)

	case "Fields":
		var sb strings.Builder
		for _, _stagesetfield := range stagesetmodel.Fields {
			sb.WriteString(__gong__marshallSliceOfPointers(ident, "Fields", _stagesetfield.GongGetIdentifier(stage)))
		}
		res = sb.String()
	default:
		log.Panicf("Unknown field %s for Gongstruct StageSetModel", fieldName)
	}
	return
}

// insertion point for marshall all fields methods
func (gongbasicfield *GongBasicField) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(gongbasicfield.GongMarshallField(stage, "Name"))
		initializerStatements.WriteString(gongbasicfield.GongMarshallField(stage, "BasicKindName"))
		pointersInitializesStatements.WriteString(gongbasicfield.GongMarshallField(stage, "GongEnum"))
		initializerStatements.WriteString(gongbasicfield.GongMarshallField(stage, "DeclaredType"))
		initializerStatements.WriteString(gongbasicfield.GongMarshallField(stage, "CompositeStructName"))
		initializerStatements.WriteString(gongbasicfield.GongMarshallField(stage, "IsAccordionStart"))
		initializerStatements.WriteString(gongbasicfield.GongMarshallField(stage, "AccordionName"))
		initializerStatements.WriteString(gongbasicfield.GongMarshallField(stage, "IsAccordionEnd"))
		initializerStatements.WriteString(gongbasicfield.GongMarshallField(stage, "Index"))
		initializerStatements.WriteString(gongbasicfield.GongMarshallField(stage, "IsTextArea"))
		initializerStatements.WriteString(gongbasicfield.GongMarshallField(stage, "IsBespokeWidth"))
		initializerStatements.WriteString(gongbasicfield.GongMarshallField(stage, "BespokeWidth"))
		initializerStatements.WriteString(gongbasicfield.GongMarshallField(stage, "IsBespokeHeight"))
		initializerStatements.WriteString(gongbasicfield.GongMarshallField(stage, "BespokeHeight"))
	}
	initRes = initializerStatements.String()
	ptrRes = pointersInitializesStatements.String()
	return
}
func (gongenum *GongEnum) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(gongenum.GongMarshallField(stage, "Name"))
		initializerStatements.WriteString(gongenum.GongMarshallField(stage, "Type"))
		pointersInitializesStatements.WriteString(gongenum.GongMarshallField(stage, "GongEnumValues"))
	}
	initRes = initializerStatements.String()
	ptrRes = pointersInitializesStatements.String()
	return
}
func (gongenumvalue *GongEnumValue) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(gongenumvalue.GongMarshallField(stage, "Name"))
		initializerStatements.WriteString(gongenumvalue.GongMarshallField(stage, "Value"))
	}
	initRes = initializerStatements.String()
	ptrRes = pointersInitializesStatements.String()
	return
}
func (gonglink *GongLink) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(gonglink.GongMarshallField(stage, "Name"))
		initializerStatements.WriteString(gonglink.GongMarshallField(stage, "Recv"))
		initializerStatements.WriteString(gonglink.GongMarshallField(stage, "ImportPath"))
	}
	initRes = initializerStatements.String()
	ptrRes = pointersInitializesStatements.String()
	return
}
func (gongnote *GongNote) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(gongnote.GongMarshallField(stage, "Name"))
		initializerStatements.WriteString(gongnote.GongMarshallField(stage, "Body"))
		initializerStatements.WriteString(gongnote.GongMarshallField(stage, "BodyHTML"))
		pointersInitializesStatements.WriteString(gongnote.GongMarshallField(stage, "Links"))
	}
	initRes = initializerStatements.String()
	ptrRes = pointersInitializesStatements.String()
	return
}
func (gongstruct *GongStruct) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(gongstruct.GongMarshallField(stage, "Name"))
		pointersInitializesStatements.WriteString(gongstruct.GongMarshallField(stage, "GongBasicFields"))
		pointersInitializesStatements.WriteString(gongstruct.GongMarshallField(stage, "GongTimeFields"))
		pointersInitializesStatements.WriteString(gongstruct.GongMarshallField(stage, "PointerToGongStructFields"))
		pointersInitializesStatements.WriteString(gongstruct.GongMarshallField(stage, "SliceOfPointerToGongStructFields"))
		initializerStatements.WriteString(gongstruct.GongMarshallField(stage, "HasOnAfterUpdateSignature"))
		initializerStatements.WriteString(gongstruct.GongMarshallField(stage, "IsIgnoredForFront"))
		initializerStatements.WriteString(gongstruct.GongMarshallField(stage, "IsOmittedForMarshalling"))
	}
	initRes = initializerStatements.String()
	ptrRes = pointersInitializesStatements.String()
	return
}
func (gongtimefield *GongTimeField) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(gongtimefield.GongMarshallField(stage, "Name"))
		initializerStatements.WriteString(gongtimefield.GongMarshallField(stage, "Index"))
		initializerStatements.WriteString(gongtimefield.GongMarshallField(stage, "CompositeStructName"))
		initializerStatements.WriteString(gongtimefield.GongMarshallField(stage, "IsAccordionStart"))
		initializerStatements.WriteString(gongtimefield.GongMarshallField(stage, "AccordionName"))
		initializerStatements.WriteString(gongtimefield.GongMarshallField(stage, "IsAccordionEnd"))
		initializerStatements.WriteString(gongtimefield.GongMarshallField(stage, "BespokeTimeFormat"))
		initializerStatements.WriteString(gongtimefield.GongMarshallField(stage, "TimeFormOnly"))
	}
	initRes = initializerStatements.String()
	ptrRes = pointersInitializesStatements.String()
	return
}
func (metareference *MetaReference) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(metareference.GongMarshallField(stage, "Name"))
	}
	initRes = initializerStatements.String()
	ptrRes = pointersInitializesStatements.String()
	return
}
func (modelpkg *ModelPkg) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(modelpkg.GongMarshallField(stage, "Name"))
		initializerStatements.WriteString(modelpkg.GongMarshallField(stage, "PkgGoName"))
		initializerStatements.WriteString(modelpkg.GongMarshallField(stage, "PkgPath"))
		initializerStatements.WriteString(modelpkg.GongMarshallField(stage, "PathToGoSubDirectory"))
		initializerStatements.WriteString(modelpkg.GongMarshallField(stage, "OrmPkgGenPath"))
		initializerStatements.WriteString(modelpkg.GongMarshallField(stage, "DbOrmPkgGenPath"))
		initializerStatements.WriteString(modelpkg.GongMarshallField(stage, "DbLiteOrmPkgGenPath"))
		initializerStatements.WriteString(modelpkg.GongMarshallField(stage, "DbPkgGenPath"))
		initializerStatements.WriteString(modelpkg.GongMarshallField(stage, "ControllersPkgGenPath"))
		initializerStatements.WriteString(modelpkg.GongMarshallField(stage, "FullstackPkgGenPath"))
		initializerStatements.WriteString(modelpkg.GongMarshallField(stage, "StackPkgGenPath"))
		initializerStatements.WriteString(modelpkg.GongMarshallField(stage, "Level1StackPkgGenPath"))
		initializerStatements.WriteString(modelpkg.GongMarshallField(stage, "StaticPkgGenPath"))
		initializerStatements.WriteString(modelpkg.GongMarshallField(stage, "ProbePkgGenPath"))
		initializerStatements.WriteString(modelpkg.GongMarshallField(stage, "NgWorkspacePath"))
		initializerStatements.WriteString(modelpkg.GongMarshallField(stage, "NgWorkspaceName"))
		initializerStatements.WriteString(modelpkg.GongMarshallField(stage, "NgDataLibrarySourceCodeDirectory"))
		initializerStatements.WriteString(modelpkg.GongMarshallField(stage, "NgSpecificLibrarySourceCodeDirectory"))
		initializerStatements.WriteString(modelpkg.GongMarshallField(stage, "MaterialLibDatamodelTargetPath"))
	}
	initRes = initializerStatements.String()
	ptrRes = pointersInitializesStatements.String()
	return
}
func (pointertogongstructfield *PointerToGongStructField) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(pointertogongstructfield.GongMarshallField(stage, "Name"))
		pointersInitializesStatements.WriteString(pointertogongstructfield.GongMarshallField(stage, "GongStruct"))
		initializerStatements.WriteString(pointertogongstructfield.GongMarshallField(stage, "Index"))
		initializerStatements.WriteString(pointertogongstructfield.GongMarshallField(stage, "CompositeStructName"))
		initializerStatements.WriteString(pointertogongstructfield.GongMarshallField(stage, "IsAccordionStart"))
		initializerStatements.WriteString(pointertogongstructfield.GongMarshallField(stage, "AccordionName"))
		initializerStatements.WriteString(pointertogongstructfield.GongMarshallField(stage, "IsAccordionEnd"))
		initializerStatements.WriteString(pointertogongstructfield.GongMarshallField(stage, "IsType"))
	}
	initRes = initializerStatements.String()
	ptrRes = pointersInitializesStatements.String()
	return
}
func (sliceofpointertogongstructfield *SliceOfPointerToGongStructField) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(sliceofpointertogongstructfield.GongMarshallField(stage, "Name"))
		pointersInitializesStatements.WriteString(sliceofpointertogongstructfield.GongMarshallField(stage, "GongStruct"))
		initializerStatements.WriteString(sliceofpointertogongstructfield.GongMarshallField(stage, "Index"))
		initializerStatements.WriteString(sliceofpointertogongstructfield.GongMarshallField(stage, "CompositeStructName"))
		initializerStatements.WriteString(sliceofpointertogongstructfield.GongMarshallField(stage, "IsAccordionStart"))
		initializerStatements.WriteString(sliceofpointertogongstructfield.GongMarshallField(stage, "AccordionName"))
		initializerStatements.WriteString(sliceofpointertogongstructfield.GongMarshallField(stage, "IsAccordionEnd"))
	}
	initRes = initializerStatements.String()
	ptrRes = pointersInitializesStatements.String()
	return
}
func (stagesetfield *StageSetField) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(stagesetfield.GongMarshallField(stage, "Name"))
		initializerStatements.WriteString(stagesetfield.GongMarshallField(stage, "PackageName"))
		initializerStatements.WriteString(stagesetfield.GongMarshallField(stage, "PackagePath"))
		initializerStatements.WriteString(stagesetfield.GongMarshallField(stage, "IsLocal"))
		initializerStatements.WriteString(stagesetfield.GongMarshallField(stage, "ImportAlias"))
	}
	initRes = initializerStatements.String()
	ptrRes = pointersInitializesStatements.String()
	return
}
func (stagesetmodel *StageSetModel) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(stagesetmodel.GongMarshallField(stage, "Name"))
		pointersInitializesStatements.WriteString(stagesetmodel.GongMarshallField(stage, "Fields"))
		initializerStatements.WriteString(stagesetmodel.GongMarshallField(stage, "IsManual"))
	}
	initRes = initializerStatements.String()
	ptrRes = pointersInitializesStatements.String()
	return
}
