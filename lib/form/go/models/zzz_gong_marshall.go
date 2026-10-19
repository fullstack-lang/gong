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
	modelsPackageName := "github.com/fullstack-lang/gong/lib/form/go/models"
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
	localAliasPrefix := "__models"
	if stage.IsInDeltaMode() || stage.GetGongMarshallingMode() == GongMarshallingAppendCommit {
		localAliasPrefix = ""
	}
	_ = localAliasPrefix

	if stage != nil {
		for _, checkbox := range __gong__sortStageSetInstances(stage.CheckBoxs, stage.CheckBox_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			checkboxIdent := localAliasPrefix + checkbox.GongGetIdentifier(stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&models.CheckBox{Name: %s}).Stage(stage)", checkboxIdent, __gong__toRawStringLiteral(checkbox.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", checkboxIdent, __gong__toRawStringLiteral(checkbox.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Value = %t", checkboxIdent, checkbox.Value))
		}
	}
	if stage != nil {
		for _, formdiv := range __gong__sortStageSetInstances(stage.FormDivs, stage.FormDiv_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			formdivIdent := localAliasPrefix + formdiv.GongGetIdentifier(stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&models.FormDiv{Name: %s}).Stage(stage)", formdivIdent, __gong__toRawStringLiteral(formdiv.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", formdivIdent, __gong__toRawStringLiteral(formdiv.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.IsADivider = %t", formdivIdent, formdiv.IsADivider))
			values.WriteString(fmt.Sprintf("\n\t%s.IsAStartAccordionGroup = %t", formdivIdent, formdiv.IsAStartAccordionGroup))
			values.WriteString(fmt.Sprintf("\n\t%s.AccordionGroupName = %s", formdivIdent, __gong__toRawStringLiteral(formdiv.AccordionGroupName)))
			values.WriteString(fmt.Sprintf("\n\t%s.IsAEndAccordionGroup = %t", formdivIdent, formdiv.IsAEndAccordionGroup))
			for _, elem := range formdiv.FormFields {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := localAliasPrefix + elem.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.FormFields = append(%s.FormFields, %s)", formdivIdent, formdivIdent, targetIdent))
			}
			for _, elem := range formdiv.CheckBoxs {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := localAliasPrefix + elem.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.CheckBoxs = append(%s.CheckBoxs, %s)", formdivIdent, formdivIdent, targetIdent))
			}
			if formdiv.FormEditAssocButton != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := localAliasPrefix + formdiv.FormEditAssocButton.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.FormEditAssocButton = %s", formdivIdent, targetIdent))
			}
			if formdiv.FormSortAssocButton != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := localAliasPrefix + formdiv.FormSortAssocButton.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.FormSortAssocButton = %s", formdivIdent, targetIdent))
			}
		}
	}
	if stage != nil {
		for _, formeditassocbutton := range __gong__sortStageSetInstances(stage.FormEditAssocButtons, stage.FormEditAssocButton_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			formeditassocbuttonIdent := localAliasPrefix + formeditassocbutton.GongGetIdentifier(stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&models.FormEditAssocButton{Name: %s}).Stage(stage)", formeditassocbuttonIdent, __gong__toRawStringLiteral(formeditassocbutton.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", formeditassocbuttonIdent, __gong__toRawStringLiteral(formeditassocbutton.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Label = %s", formeditassocbuttonIdent, __gong__toRawStringLiteral(formeditassocbutton.Label)))
			values.WriteString(fmt.Sprintf("\n\t%s.AssociationStorage = %s", formeditassocbuttonIdent, __gong__toRawStringLiteral(formeditassocbutton.AssociationStorage)))
			values.WriteString(fmt.Sprintf("\n\t%s.HasChanged = %t", formeditassocbuttonIdent, formeditassocbutton.HasChanged))
			values.WriteString(fmt.Sprintf("\n\t%s.IsForSavePurpose = %t", formeditassocbuttonIdent, formeditassocbutton.IsForSavePurpose))
			values.WriteString(fmt.Sprintf("\n\t%s.HasToolTip = %t", formeditassocbuttonIdent, formeditassocbutton.HasToolTip))
			values.WriteString(fmt.Sprintf("\n\t%s.ToolTipText = %s", formeditassocbuttonIdent, __gong__toRawStringLiteral(formeditassocbutton.ToolTipText)))
			values.WriteString(fmt.Sprintf("\n\t%s.MatTooltipShowDelay = %s", formeditassocbuttonIdent, __gong__toRawStringLiteral(formeditassocbutton.MatTooltipShowDelay)))
		}
	}
	if stage != nil {
		for _, formfield := range __gong__sortStageSetInstances(stage.FormFields, stage.FormField_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			formfieldIdent := localAliasPrefix + formfield.GongGetIdentifier(stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&models.FormField{Name: %s}).Stage(stage)", formfieldIdent, __gong__toRawStringLiteral(formfield.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", formfieldIdent, __gong__toRawStringLiteral(formfield.Name)))
			if formfield.InputTypeEnum.ToCodeString() != "" {
				values.WriteString(fmt.Sprintf("\n\t%s.InputTypeEnum = models.%s", formfieldIdent, formfield.InputTypeEnum.ToCodeString()))
			} else {
				values.WriteString(fmt.Sprintf("\n\t%s.InputTypeEnum = \"\"", formfieldIdent))
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Label = %s", formfieldIdent, __gong__toRawStringLiteral(formfield.Label)))
			values.WriteString(fmt.Sprintf("\n\t%s.Placeholder = %s", formfieldIdent, __gong__toRawStringLiteral(formfield.Placeholder)))
			values.WriteString(fmt.Sprintf("\n\t%s.HasBespokeWidth = %t", formfieldIdent, formfield.HasBespokeWidth))
			values.WriteString(fmt.Sprintf("\n\t%s.BespokeWidthPx = %d", formfieldIdent, formfield.BespokeWidthPx))
			values.WriteString(fmt.Sprintf("\n\t%s.HasBespokeHeight = %t", formfieldIdent, formfield.HasBespokeHeight))
			values.WriteString(fmt.Sprintf("\n\t%s.BespokeHeightPx = %d", formfieldIdent, formfield.BespokeHeightPx))
			if formfield.FormFieldString != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := localAliasPrefix + formfield.FormFieldString.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.FormFieldString = %s", formfieldIdent, targetIdent))
			}
			if formfield.FormFieldFloat64 != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := localAliasPrefix + formfield.FormFieldFloat64.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.FormFieldFloat64 = %s", formfieldIdent, targetIdent))
			}
			if formfield.FormFieldInt != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := localAliasPrefix + formfield.FormFieldInt.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.FormFieldInt = %s", formfieldIdent, targetIdent))
			}
			if formfield.FormFieldDate != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := localAliasPrefix + formfield.FormFieldDate.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.FormFieldDate = %s", formfieldIdent, targetIdent))
			}
			if formfield.FormFieldTime != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := localAliasPrefix + formfield.FormFieldTime.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.FormFieldTime = %s", formfieldIdent, targetIdent))
			}
			if formfield.FormFieldDateTime != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := localAliasPrefix + formfield.FormFieldDateTime.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.FormFieldDateTime = %s", formfieldIdent, targetIdent))
			}
			if formfield.FormFieldSelect != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := localAliasPrefix + formfield.FormFieldSelect.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.FormFieldSelect = %s", formfieldIdent, targetIdent))
			}
		}
	}
	if stage != nil {
		for _, formfielddate := range __gong__sortStageSetInstances(stage.FormFieldDates, stage.FormFieldDate_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			formfielddateIdent := localAliasPrefix + formfielddate.GongGetIdentifier(stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&models.FormFieldDate{Name: %s}).Stage(stage)", formfielddateIdent, __gong__toRawStringLiteral(formfielddate.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", formfielddateIdent, __gong__toRawStringLiteral(formfielddate.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Value, _ = time.Parse(\"2006-01-02 15:04:05.999999999 -0700 MST\", \"%s\")", formfielddateIdent, formfielddate.Value.String()))
		}
	}
	if stage != nil {
		for _, formfielddatetime := range __gong__sortStageSetInstances(stage.FormFieldDateTimes, stage.FormFieldDateTime_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			formfielddatetimeIdent := localAliasPrefix + formfielddatetime.GongGetIdentifier(stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&models.FormFieldDateTime{Name: %s}).Stage(stage)", formfielddatetimeIdent, __gong__toRawStringLiteral(formfielddatetime.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", formfielddatetimeIdent, __gong__toRawStringLiteral(formfielddatetime.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Value, _ = time.Parse(\"2006-01-02 15:04:05.999999999 -0700 MST\", \"%s\")", formfielddatetimeIdent, formfielddatetime.Value.String()))
		}
	}
	if stage != nil {
		for _, formfieldfloat64 := range __gong__sortStageSetInstances(stage.FormFieldFloat64s, stage.FormFieldFloat64_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			formfieldfloat64Ident := localAliasPrefix + formfieldfloat64.GongGetIdentifier(stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&models.FormFieldFloat64{Name: %s}).Stage(stage)", formfieldfloat64Ident, __gong__toRawStringLiteral(formfieldfloat64.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", formfieldfloat64Ident, __gong__toRawStringLiteral(formfieldfloat64.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Value = %f", formfieldfloat64Ident, formfieldfloat64.Value))
			values.WriteString(fmt.Sprintf("\n\t%s.HasMinValidator = %t", formfieldfloat64Ident, formfieldfloat64.HasMinValidator))
			values.WriteString(fmt.Sprintf("\n\t%s.MinValue = %f", formfieldfloat64Ident, formfieldfloat64.MinValue))
			values.WriteString(fmt.Sprintf("\n\t%s.HasMaxValidator = %t", formfieldfloat64Ident, formfieldfloat64.HasMaxValidator))
			values.WriteString(fmt.Sprintf("\n\t%s.MaxValue = %f", formfieldfloat64Ident, formfieldfloat64.MaxValue))
		}
	}
	if stage != nil {
		for _, formfieldint := range __gong__sortStageSetInstances(stage.FormFieldInts, stage.FormFieldInt_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			formfieldintIdent := localAliasPrefix + formfieldint.GongGetIdentifier(stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&models.FormFieldInt{Name: %s}).Stage(stage)", formfieldintIdent, __gong__toRawStringLiteral(formfieldint.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", formfieldintIdent, __gong__toRawStringLiteral(formfieldint.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Value = %d", formfieldintIdent, formfieldint.Value))
			values.WriteString(fmt.Sprintf("\n\t%s.HasMinValidator = %t", formfieldintIdent, formfieldint.HasMinValidator))
			values.WriteString(fmt.Sprintf("\n\t%s.MinValue = %d", formfieldintIdent, formfieldint.MinValue))
			values.WriteString(fmt.Sprintf("\n\t%s.HasMaxValidator = %t", formfieldintIdent, formfieldint.HasMaxValidator))
			values.WriteString(fmt.Sprintf("\n\t%s.MaxValue = %d", formfieldintIdent, formfieldint.MaxValue))
		}
	}
	if stage != nil {
		for _, formfieldselect := range __gong__sortStageSetInstances(stage.FormFieldSelects, stage.FormFieldSelect_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			formfieldselectIdent := localAliasPrefix + formfieldselect.GongGetIdentifier(stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&models.FormFieldSelect{Name: %s}).Stage(stage)", formfieldselectIdent, __gong__toRawStringLiteral(formfieldselect.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", formfieldselectIdent, __gong__toRawStringLiteral(formfieldselect.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.CanBeEmpty = %t", formfieldselectIdent, formfieldselect.CanBeEmpty))
			values.WriteString(fmt.Sprintf("\n\t%s.PreserveInitialOrder = %t", formfieldselectIdent, formfieldselect.PreserveInitialOrder))
			if formfieldselect.Value != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := localAliasPrefix + formfieldselect.Value.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Value = %s", formfieldselectIdent, targetIdent))
			}
			for _, elem := range formfieldselect.Options {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := localAliasPrefix + elem.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Options = append(%s.Options, %s)", formfieldselectIdent, formfieldselectIdent, targetIdent))
			}
		}
	}
	if stage != nil {
		for _, formfieldstring := range __gong__sortStageSetInstances(stage.FormFieldStrings, stage.FormFieldString_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			formfieldstringIdent := localAliasPrefix + formfieldstring.GongGetIdentifier(stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&models.FormFieldString{Name: %s}).Stage(stage)", formfieldstringIdent, __gong__toRawStringLiteral(formfieldstring.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", formfieldstringIdent, __gong__toRawStringLiteral(formfieldstring.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Value = %s", formfieldstringIdent, __gong__toRawStringLiteral(formfieldstring.Value)))
			values.WriteString(fmt.Sprintf("\n\t%s.IsTextArea = %t", formfieldstringIdent, formfieldstring.IsTextArea))
		}
	}
	if stage != nil {
		for _, formfieldtime := range __gong__sortStageSetInstances(stage.FormFieldTimes, stage.FormFieldTime_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			formfieldtimeIdent := localAliasPrefix + formfieldtime.GongGetIdentifier(stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&models.FormFieldTime{Name: %s}).Stage(stage)", formfieldtimeIdent, __gong__toRawStringLiteral(formfieldtime.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", formfieldtimeIdent, __gong__toRawStringLiteral(formfieldtime.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Value, _ = time.Parse(\"2006-01-02 15:04:05.999999999 -0700 MST\", \"%s\")", formfieldtimeIdent, formfieldtime.Value.String()))
			values.WriteString(fmt.Sprintf("\n\t%s.Step = %f", formfieldtimeIdent, formfieldtime.Step))
		}
	}
	if stage != nil {
		for _, formgroup := range __gong__sortStageSetInstances(stage.FormGroups, stage.FormGroup_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			formgroupIdent := localAliasPrefix + formgroup.GongGetIdentifier(stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&models.FormGroup{Name: %s}).Stage(stage)", formgroupIdent, __gong__toRawStringLiteral(formgroup.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", formgroupIdent, __gong__toRawStringLiteral(formgroup.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Label = %s", formgroupIdent, __gong__toRawStringLiteral(formgroup.Label)))
			values.WriteString(fmt.Sprintf("\n\t%s.TypeLabel = %s", formgroupIdent, __gong__toRawStringLiteral(formgroup.TypeLabel)))
			values.WriteString(fmt.Sprintf("\n\t%s.HasSuppressButton = %t", formgroupIdent, formgroup.HasSuppressButton))
			values.WriteString(fmt.Sprintf("\n\t%s.HasSuppressButtonBeenPressed = %t", formgroupIdent, formgroup.HasSuppressButtonBeenPressed))
			for _, elem := range formgroup.FormDivs {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := localAliasPrefix + elem.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.FormDivs = append(%s.FormDivs, %s)", formgroupIdent, formgroupIdent, targetIdent))
			}
		}
	}
	if stage != nil {
		for _, formsortassocbutton := range __gong__sortStageSetInstances(stage.FormSortAssocButtons, stage.FormSortAssocButton_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			formsortassocbuttonIdent := localAliasPrefix + formsortassocbutton.GongGetIdentifier(stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&models.FormSortAssocButton{Name: %s}).Stage(stage)", formsortassocbuttonIdent, __gong__toRawStringLiteral(formsortassocbutton.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", formsortassocbuttonIdent, __gong__toRawStringLiteral(formsortassocbutton.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Label = %s", formsortassocbuttonIdent, __gong__toRawStringLiteral(formsortassocbutton.Label)))
			values.WriteString(fmt.Sprintf("\n\t%s.HasToolTip = %t", formsortassocbuttonIdent, formsortassocbutton.HasToolTip))
			values.WriteString(fmt.Sprintf("\n\t%s.ToolTipText = %s", formsortassocbuttonIdent, __gong__toRawStringLiteral(formsortassocbutton.ToolTipText)))
			values.WriteString(fmt.Sprintf("\n\t%s.MatTooltipShowDelay = %s", formsortassocbuttonIdent, __gong__toRawStringLiteral(formsortassocbutton.MatTooltipShowDelay)))
			if formsortassocbutton.FormEditAssocButton != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := localAliasPrefix + formsortassocbutton.FormEditAssocButton.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.FormEditAssocButton = %s", formsortassocbuttonIdent, targetIdent))
			}
		}
	}
	if stage != nil {
		for _, option := range __gong__sortStageSetInstances(stage.Options, stage.Option_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			optionIdent := localAliasPrefix + option.GongGetIdentifier(stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&models.Option{Name: %s}).Stage(stage)", optionIdent, __gong__toRawStringLiteral(option.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", optionIdent, __gong__toRawStringLiteral(option.Name)))
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
go run github.com/fullstack-lang/gong/lib/form/go/models/probe/cmd/probe@%s stage.go
*/
package %s

import (
	"slices"
	"time"

	"github.com/fullstack-lang/gong/lib/form/go/models"%s
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

	if modelsPackageName != "" && modelsPackageName != "github.com/fullstack-lang/gong/lib/form/go/models" {
		res = strings.ReplaceAll(res, "\"github.com/fullstack-lang/gong/lib/form/go/models\"", fmt.Sprintf("%q", modelsPackageName))
	}

	return res, nil
}

// insertion point for marshall field methods
func (checkbox *CheckBox) GongMarshallField(stage *Stage, fieldName string) (res string) {
	ident := checkbox.GongGetIdentifier(stage)
	_ = ident
	switch fieldName {
	case "Name":
		res = __gong__marshallString(ident, "Name", checkbox.Name)
	case "Value":
		res = __gong__marshallBool(ident, "Value", checkbox.Value)

	default:
		log.Panicf("Unknown field %s for Gongstruct CheckBox", fieldName)
	}
	return
}

func (formdiv *FormDiv) GongMarshallField(stage *Stage, fieldName string) (res string) {
	ident := formdiv.GongGetIdentifier(stage)
	_ = ident
	switch fieldName {
	case "Name":
		res = __gong__marshallString(ident, "Name", formdiv.Name)
	case "IsADivider":
		res = __gong__marshallBool(ident, "IsADivider", formdiv.IsADivider)
	case "IsAStartAccordionGroup":
		res = __gong__marshallBool(ident, "IsAStartAccordionGroup", formdiv.IsAStartAccordionGroup)
	case "AccordionGroupName":
		res = __gong__marshallString(ident, "AccordionGroupName", formdiv.AccordionGroupName)
	case "IsAEndAccordionGroup":
		res = __gong__marshallBool(ident, "IsAEndAccordionGroup", formdiv.IsAEndAccordionGroup)

	case "FormFields":
		var sb strings.Builder
		for _, _formfield := range formdiv.FormFields {
			sb.WriteString(__gong__marshallSliceOfPointers(ident, "FormFields", _formfield.GongGetIdentifier(stage)))
		}
		res = sb.String()
	case "CheckBoxs":
		var sb strings.Builder
		for _, _checkbox := range formdiv.CheckBoxs {
			sb.WriteString(__gong__marshallSliceOfPointers(ident, "CheckBoxs", _checkbox.GongGetIdentifier(stage)))
		}
		res = sb.String()
	case "FormEditAssocButton":
		if formdiv.FormEditAssocButton != nil {
			res = __gong__marshallPointer(ident, "FormEditAssocButton", formdiv.FormEditAssocButton.GongGetIdentifier(stage))
		} else {
			res = __gong__marshallPointer(ident, "FormEditAssocButton", "nil")
		}
	case "FormSortAssocButton":
		if formdiv.FormSortAssocButton != nil {
			res = __gong__marshallPointer(ident, "FormSortAssocButton", formdiv.FormSortAssocButton.GongGetIdentifier(stage))
		} else {
			res = __gong__marshallPointer(ident, "FormSortAssocButton", "nil")
		}
	default:
		log.Panicf("Unknown field %s for Gongstruct FormDiv", fieldName)
	}
	return
}

func (formeditassocbutton *FormEditAssocButton) GongMarshallField(stage *Stage, fieldName string) (res string) {
	ident := formeditassocbutton.GongGetIdentifier(stage)
	_ = ident
	switch fieldName {
	case "Name":
		res = __gong__marshallString(ident, "Name", formeditassocbutton.Name)
	case "Label":
		res = __gong__marshallString(ident, "Label", formeditassocbutton.Label)
	case "AssociationStorage":
		res = __gong__marshallString(ident, "AssociationStorage", formeditassocbutton.AssociationStorage)
	case "HasChanged":
		res = __gong__marshallBool(ident, "HasChanged", formeditassocbutton.HasChanged)
	case "IsForSavePurpose":
		res = __gong__marshallBool(ident, "IsForSavePurpose", formeditassocbutton.IsForSavePurpose)
	case "HasToolTip":
		res = __gong__marshallBool(ident, "HasToolTip", formeditassocbutton.HasToolTip)
	case "ToolTipText":
		res = __gong__marshallString(ident, "ToolTipText", formeditassocbutton.ToolTipText)
	case "MatTooltipShowDelay":
		res = __gong__marshallString(ident, "MatTooltipShowDelay", formeditassocbutton.MatTooltipShowDelay)

	default:
		log.Panicf("Unknown field %s for Gongstruct FormEditAssocButton", fieldName)
	}
	return
}

func (formfield *FormField) GongMarshallField(stage *Stage, fieldName string) (res string) {
	ident := formfield.GongGetIdentifier(stage)
	_ = ident
	switch fieldName {
	case "Name":
		res = __gong__marshallString(ident, "Name", formfield.Name)
	case "InputTypeEnum":
		res = __gong__marshallEnumString(ident, "InputTypeEnum", formfield.InputTypeEnum.ToCodeString())
	case "Label":
		res = __gong__marshallString(ident, "Label", formfield.Label)
	case "Placeholder":
		res = __gong__marshallString(ident, "Placeholder", formfield.Placeholder)
	case "HasBespokeWidth":
		res = __gong__marshallBool(ident, "HasBespokeWidth", formfield.HasBespokeWidth)
	case "BespokeWidthPx":
		res = __gong__marshallInt(ident, "BespokeWidthPx", formfield.BespokeWidthPx)
	case "HasBespokeHeight":
		res = __gong__marshallBool(ident, "HasBespokeHeight", formfield.HasBespokeHeight)
	case "BespokeHeightPx":
		res = __gong__marshallInt(ident, "BespokeHeightPx", formfield.BespokeHeightPx)

	case "FormFieldString":
		if formfield.FormFieldString != nil {
			res = __gong__marshallPointer(ident, "FormFieldString", formfield.FormFieldString.GongGetIdentifier(stage))
		} else {
			res = __gong__marshallPointer(ident, "FormFieldString", "nil")
		}
	case "FormFieldFloat64":
		if formfield.FormFieldFloat64 != nil {
			res = __gong__marshallPointer(ident, "FormFieldFloat64", formfield.FormFieldFloat64.GongGetIdentifier(stage))
		} else {
			res = __gong__marshallPointer(ident, "FormFieldFloat64", "nil")
		}
	case "FormFieldInt":
		if formfield.FormFieldInt != nil {
			res = __gong__marshallPointer(ident, "FormFieldInt", formfield.FormFieldInt.GongGetIdentifier(stage))
		} else {
			res = __gong__marshallPointer(ident, "FormFieldInt", "nil")
		}
	case "FormFieldDate":
		if formfield.FormFieldDate != nil {
			res = __gong__marshallPointer(ident, "FormFieldDate", formfield.FormFieldDate.GongGetIdentifier(stage))
		} else {
			res = __gong__marshallPointer(ident, "FormFieldDate", "nil")
		}
	case "FormFieldTime":
		if formfield.FormFieldTime != nil {
			res = __gong__marshallPointer(ident, "FormFieldTime", formfield.FormFieldTime.GongGetIdentifier(stage))
		} else {
			res = __gong__marshallPointer(ident, "FormFieldTime", "nil")
		}
	case "FormFieldDateTime":
		if formfield.FormFieldDateTime != nil {
			res = __gong__marshallPointer(ident, "FormFieldDateTime", formfield.FormFieldDateTime.GongGetIdentifier(stage))
		} else {
			res = __gong__marshallPointer(ident, "FormFieldDateTime", "nil")
		}
	case "FormFieldSelect":
		if formfield.FormFieldSelect != nil {
			res = __gong__marshallPointer(ident, "FormFieldSelect", formfield.FormFieldSelect.GongGetIdentifier(stage))
		} else {
			res = __gong__marshallPointer(ident, "FormFieldSelect", "nil")
		}
	default:
		log.Panicf("Unknown field %s for Gongstruct FormField", fieldName)
	}
	return
}

func (formfielddate *FormFieldDate) GongMarshallField(stage *Stage, fieldName string) (res string) {
	ident := formfielddate.GongGetIdentifier(stage)
	_ = ident
	switch fieldName {
	case "Name":
		res = __gong__marshallString(ident, "Name", formfielddate.Name)
	case "Value":
		res = __gong__marshallTime(ident, "Value", formfielddate.Value.String())

	default:
		log.Panicf("Unknown field %s for Gongstruct FormFieldDate", fieldName)
	}
	return
}

func (formfielddatetime *FormFieldDateTime) GongMarshallField(stage *Stage, fieldName string) (res string) {
	ident := formfielddatetime.GongGetIdentifier(stage)
	_ = ident
	switch fieldName {
	case "Name":
		res = __gong__marshallString(ident, "Name", formfielddatetime.Name)
	case "Value":
		res = __gong__marshallTime(ident, "Value", formfielddatetime.Value.String())

	default:
		log.Panicf("Unknown field %s for Gongstruct FormFieldDateTime", fieldName)
	}
	return
}

func (formfieldfloat64 *FormFieldFloat64) GongMarshallField(stage *Stage, fieldName string) (res string) {
	ident := formfieldfloat64.GongGetIdentifier(stage)
	_ = ident
	switch fieldName {
	case "Name":
		res = __gong__marshallString(ident, "Name", formfieldfloat64.Name)
	case "Value":
		res = __gong__marshallFloat(ident, "Value", formfieldfloat64.Value)
	case "HasMinValidator":
		res = __gong__marshallBool(ident, "HasMinValidator", formfieldfloat64.HasMinValidator)
	case "MinValue":
		res = __gong__marshallFloat(ident, "MinValue", formfieldfloat64.MinValue)
	case "HasMaxValidator":
		res = __gong__marshallBool(ident, "HasMaxValidator", formfieldfloat64.HasMaxValidator)
	case "MaxValue":
		res = __gong__marshallFloat(ident, "MaxValue", formfieldfloat64.MaxValue)

	default:
		log.Panicf("Unknown field %s for Gongstruct FormFieldFloat64", fieldName)
	}
	return
}

func (formfieldint *FormFieldInt) GongMarshallField(stage *Stage, fieldName string) (res string) {
	ident := formfieldint.GongGetIdentifier(stage)
	_ = ident
	switch fieldName {
	case "Name":
		res = __gong__marshallString(ident, "Name", formfieldint.Name)
	case "Value":
		res = __gong__marshallInt(ident, "Value", formfieldint.Value)
	case "HasMinValidator":
		res = __gong__marshallBool(ident, "HasMinValidator", formfieldint.HasMinValidator)
	case "MinValue":
		res = __gong__marshallInt(ident, "MinValue", formfieldint.MinValue)
	case "HasMaxValidator":
		res = __gong__marshallBool(ident, "HasMaxValidator", formfieldint.HasMaxValidator)
	case "MaxValue":
		res = __gong__marshallInt(ident, "MaxValue", formfieldint.MaxValue)

	default:
		log.Panicf("Unknown field %s for Gongstruct FormFieldInt", fieldName)
	}
	return
}

func (formfieldselect *FormFieldSelect) GongMarshallField(stage *Stage, fieldName string) (res string) {
	ident := formfieldselect.GongGetIdentifier(stage)
	_ = ident
	switch fieldName {
	case "Name":
		res = __gong__marshallString(ident, "Name", formfieldselect.Name)
	case "CanBeEmpty":
		res = __gong__marshallBool(ident, "CanBeEmpty", formfieldselect.CanBeEmpty)
	case "PreserveInitialOrder":
		res = __gong__marshallBool(ident, "PreserveInitialOrder", formfieldselect.PreserveInitialOrder)

	case "Value":
		if formfieldselect.Value != nil {
			res = __gong__marshallPointer(ident, "Value", formfieldselect.Value.GongGetIdentifier(stage))
		} else {
			res = __gong__marshallPointer(ident, "Value", "nil")
		}
	case "Options":
		var sb strings.Builder
		for _, _option := range formfieldselect.Options {
			sb.WriteString(__gong__marshallSliceOfPointers(ident, "Options", _option.GongGetIdentifier(stage)))
		}
		res = sb.String()
	default:
		log.Panicf("Unknown field %s for Gongstruct FormFieldSelect", fieldName)
	}
	return
}

func (formfieldstring *FormFieldString) GongMarshallField(stage *Stage, fieldName string) (res string) {
	ident := formfieldstring.GongGetIdentifier(stage)
	_ = ident
	switch fieldName {
	case "Name":
		res = __gong__marshallString(ident, "Name", formfieldstring.Name)
	case "Value":
		res = __gong__marshallString(ident, "Value", formfieldstring.Value)
	case "IsTextArea":
		res = __gong__marshallBool(ident, "IsTextArea", formfieldstring.IsTextArea)

	default:
		log.Panicf("Unknown field %s for Gongstruct FormFieldString", fieldName)
	}
	return
}

func (formfieldtime *FormFieldTime) GongMarshallField(stage *Stage, fieldName string) (res string) {
	ident := formfieldtime.GongGetIdentifier(stage)
	_ = ident
	switch fieldName {
	case "Name":
		res = __gong__marshallString(ident, "Name", formfieldtime.Name)
	case "Value":
		res = __gong__marshallTime(ident, "Value", formfieldtime.Value.String())
	case "Step":
		res = __gong__marshallFloat(ident, "Step", formfieldtime.Step)

	default:
		log.Panicf("Unknown field %s for Gongstruct FormFieldTime", fieldName)
	}
	return
}

func (formgroup *FormGroup) GongMarshallField(stage *Stage, fieldName string) (res string) {
	ident := formgroup.GongGetIdentifier(stage)
	_ = ident
	switch fieldName {
	case "Name":
		res = __gong__marshallString(ident, "Name", formgroup.Name)
	case "Label":
		res = __gong__marshallString(ident, "Label", formgroup.Label)
	case "TypeLabel":
		res = __gong__marshallString(ident, "TypeLabel", formgroup.TypeLabel)
	case "HasSuppressButton":
		res = __gong__marshallBool(ident, "HasSuppressButton", formgroup.HasSuppressButton)
	case "HasSuppressButtonBeenPressed":
		res = __gong__marshallBool(ident, "HasSuppressButtonBeenPressed", formgroup.HasSuppressButtonBeenPressed)

	case "FormDivs":
		var sb strings.Builder
		for _, _formdiv := range formgroup.FormDivs {
			sb.WriteString(__gong__marshallSliceOfPointers(ident, "FormDivs", _formdiv.GongGetIdentifier(stage)))
		}
		res = sb.String()
	default:
		log.Panicf("Unknown field %s for Gongstruct FormGroup", fieldName)
	}
	return
}

func (formsortassocbutton *FormSortAssocButton) GongMarshallField(stage *Stage, fieldName string) (res string) {
	ident := formsortassocbutton.GongGetIdentifier(stage)
	_ = ident
	switch fieldName {
	case "Name":
		res = __gong__marshallString(ident, "Name", formsortassocbutton.Name)
	case "Label":
		res = __gong__marshallString(ident, "Label", formsortassocbutton.Label)
	case "HasToolTip":
		res = __gong__marshallBool(ident, "HasToolTip", formsortassocbutton.HasToolTip)
	case "ToolTipText":
		res = __gong__marshallString(ident, "ToolTipText", formsortassocbutton.ToolTipText)
	case "MatTooltipShowDelay":
		res = __gong__marshallString(ident, "MatTooltipShowDelay", formsortassocbutton.MatTooltipShowDelay)

	case "FormEditAssocButton":
		if formsortassocbutton.FormEditAssocButton != nil {
			res = __gong__marshallPointer(ident, "FormEditAssocButton", formsortassocbutton.FormEditAssocButton.GongGetIdentifier(stage))
		} else {
			res = __gong__marshallPointer(ident, "FormEditAssocButton", "nil")
		}
	default:
		log.Panicf("Unknown field %s for Gongstruct FormSortAssocButton", fieldName)
	}
	return
}

func (option *Option) GongMarshallField(stage *Stage, fieldName string) (res string) {
	ident := option.GongGetIdentifier(stage)
	_ = ident
	switch fieldName {
	case "Name":
		res = __gong__marshallString(ident, "Name", option.Name)

	default:
		log.Panicf("Unknown field %s for Gongstruct Option", fieldName)
	}
	return
}

// insertion point for marshall all fields methods
func (checkbox *CheckBox) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(checkbox.GongMarshallField(stage, "Name"))
		initializerStatements.WriteString(checkbox.GongMarshallField(stage, "Value"))
	}
	initRes = initializerStatements.String()
	ptrRes = pointersInitializesStatements.String()
	return
}
func (formdiv *FormDiv) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(formdiv.GongMarshallField(stage, "Name"))
		pointersInitializesStatements.WriteString(formdiv.GongMarshallField(stage, "FormFields"))
		pointersInitializesStatements.WriteString(formdiv.GongMarshallField(stage, "CheckBoxs"))
		pointersInitializesStatements.WriteString(formdiv.GongMarshallField(stage, "FormEditAssocButton"))
		pointersInitializesStatements.WriteString(formdiv.GongMarshallField(stage, "FormSortAssocButton"))
		initializerStatements.WriteString(formdiv.GongMarshallField(stage, "IsADivider"))
		initializerStatements.WriteString(formdiv.GongMarshallField(stage, "IsAStartAccordionGroup"))
		initializerStatements.WriteString(formdiv.GongMarshallField(stage, "AccordionGroupName"))
		initializerStatements.WriteString(formdiv.GongMarshallField(stage, "IsAEndAccordionGroup"))
	}
	initRes = initializerStatements.String()
	ptrRes = pointersInitializesStatements.String()
	return
}
func (formeditassocbutton *FormEditAssocButton) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(formeditassocbutton.GongMarshallField(stage, "Name"))
		initializerStatements.WriteString(formeditassocbutton.GongMarshallField(stage, "Label"))
		initializerStatements.WriteString(formeditassocbutton.GongMarshallField(stage, "AssociationStorage"))
		initializerStatements.WriteString(formeditassocbutton.GongMarshallField(stage, "HasChanged"))
		initializerStatements.WriteString(formeditassocbutton.GongMarshallField(stage, "IsForSavePurpose"))
		initializerStatements.WriteString(formeditassocbutton.GongMarshallField(stage, "HasToolTip"))
		initializerStatements.WriteString(formeditassocbutton.GongMarshallField(stage, "ToolTipText"))
		initializerStatements.WriteString(formeditassocbutton.GongMarshallField(stage, "MatTooltipShowDelay"))
	}
	initRes = initializerStatements.String()
	ptrRes = pointersInitializesStatements.String()
	return
}
func (formfield *FormField) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(formfield.GongMarshallField(stage, "Name"))
		initializerStatements.WriteString(formfield.GongMarshallField(stage, "InputTypeEnum"))
		initializerStatements.WriteString(formfield.GongMarshallField(stage, "Label"))
		initializerStatements.WriteString(formfield.GongMarshallField(stage, "Placeholder"))
		pointersInitializesStatements.WriteString(formfield.GongMarshallField(stage, "FormFieldString"))
		pointersInitializesStatements.WriteString(formfield.GongMarshallField(stage, "FormFieldFloat64"))
		pointersInitializesStatements.WriteString(formfield.GongMarshallField(stage, "FormFieldInt"))
		pointersInitializesStatements.WriteString(formfield.GongMarshallField(stage, "FormFieldDate"))
		pointersInitializesStatements.WriteString(formfield.GongMarshallField(stage, "FormFieldTime"))
		pointersInitializesStatements.WriteString(formfield.GongMarshallField(stage, "FormFieldDateTime"))
		pointersInitializesStatements.WriteString(formfield.GongMarshallField(stage, "FormFieldSelect"))
		initializerStatements.WriteString(formfield.GongMarshallField(stage, "HasBespokeWidth"))
		initializerStatements.WriteString(formfield.GongMarshallField(stage, "BespokeWidthPx"))
		initializerStatements.WriteString(formfield.GongMarshallField(stage, "HasBespokeHeight"))
		initializerStatements.WriteString(formfield.GongMarshallField(stage, "BespokeHeightPx"))
	}
	initRes = initializerStatements.String()
	ptrRes = pointersInitializesStatements.String()
	return
}
func (formfielddate *FormFieldDate) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(formfielddate.GongMarshallField(stage, "Name"))
		initializerStatements.WriteString(formfielddate.GongMarshallField(stage, "Value"))
	}
	initRes = initializerStatements.String()
	ptrRes = pointersInitializesStatements.String()
	return
}
func (formfielddatetime *FormFieldDateTime) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(formfielddatetime.GongMarshallField(stage, "Name"))
		initializerStatements.WriteString(formfielddatetime.GongMarshallField(stage, "Value"))
	}
	initRes = initializerStatements.String()
	ptrRes = pointersInitializesStatements.String()
	return
}
func (formfieldfloat64 *FormFieldFloat64) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(formfieldfloat64.GongMarshallField(stage, "Name"))
		initializerStatements.WriteString(formfieldfloat64.GongMarshallField(stage, "Value"))
		initializerStatements.WriteString(formfieldfloat64.GongMarshallField(stage, "HasMinValidator"))
		initializerStatements.WriteString(formfieldfloat64.GongMarshallField(stage, "MinValue"))
		initializerStatements.WriteString(formfieldfloat64.GongMarshallField(stage, "HasMaxValidator"))
		initializerStatements.WriteString(formfieldfloat64.GongMarshallField(stage, "MaxValue"))
	}
	initRes = initializerStatements.String()
	ptrRes = pointersInitializesStatements.String()
	return
}
func (formfieldint *FormFieldInt) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(formfieldint.GongMarshallField(stage, "Name"))
		initializerStatements.WriteString(formfieldint.GongMarshallField(stage, "Value"))
		initializerStatements.WriteString(formfieldint.GongMarshallField(stage, "HasMinValidator"))
		initializerStatements.WriteString(formfieldint.GongMarshallField(stage, "MinValue"))
		initializerStatements.WriteString(formfieldint.GongMarshallField(stage, "HasMaxValidator"))
		initializerStatements.WriteString(formfieldint.GongMarshallField(stage, "MaxValue"))
	}
	initRes = initializerStatements.String()
	ptrRes = pointersInitializesStatements.String()
	return
}
func (formfieldselect *FormFieldSelect) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(formfieldselect.GongMarshallField(stage, "Name"))
		pointersInitializesStatements.WriteString(formfieldselect.GongMarshallField(stage, "Value"))
		pointersInitializesStatements.WriteString(formfieldselect.GongMarshallField(stage, "Options"))
		initializerStatements.WriteString(formfieldselect.GongMarshallField(stage, "CanBeEmpty"))
		initializerStatements.WriteString(formfieldselect.GongMarshallField(stage, "PreserveInitialOrder"))
	}
	initRes = initializerStatements.String()
	ptrRes = pointersInitializesStatements.String()
	return
}
func (formfieldstring *FormFieldString) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(formfieldstring.GongMarshallField(stage, "Name"))
		initializerStatements.WriteString(formfieldstring.GongMarshallField(stage, "Value"))
		initializerStatements.WriteString(formfieldstring.GongMarshallField(stage, "IsTextArea"))
	}
	initRes = initializerStatements.String()
	ptrRes = pointersInitializesStatements.String()
	return
}
func (formfieldtime *FormFieldTime) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(formfieldtime.GongMarshallField(stage, "Name"))
		initializerStatements.WriteString(formfieldtime.GongMarshallField(stage, "Value"))
		initializerStatements.WriteString(formfieldtime.GongMarshallField(stage, "Step"))
	}
	initRes = initializerStatements.String()
	ptrRes = pointersInitializesStatements.String()
	return
}
func (formgroup *FormGroup) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(formgroup.GongMarshallField(stage, "Name"))
		initializerStatements.WriteString(formgroup.GongMarshallField(stage, "Label"))
		initializerStatements.WriteString(formgroup.GongMarshallField(stage, "TypeLabel"))
		pointersInitializesStatements.WriteString(formgroup.GongMarshallField(stage, "FormDivs"))
		initializerStatements.WriteString(formgroup.GongMarshallField(stage, "HasSuppressButton"))
		initializerStatements.WriteString(formgroup.GongMarshallField(stage, "HasSuppressButtonBeenPressed"))
	}
	initRes = initializerStatements.String()
	ptrRes = pointersInitializesStatements.String()
	return
}
func (formsortassocbutton *FormSortAssocButton) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(formsortassocbutton.GongMarshallField(stage, "Name"))
		initializerStatements.WriteString(formsortassocbutton.GongMarshallField(stage, "Label"))
		initializerStatements.WriteString(formsortassocbutton.GongMarshallField(stage, "HasToolTip"))
		initializerStatements.WriteString(formsortassocbutton.GongMarshallField(stage, "ToolTipText"))
		initializerStatements.WriteString(formsortassocbutton.GongMarshallField(stage, "MatTooltipShowDelay"))
		pointersInitializesStatements.WriteString(formsortassocbutton.GongMarshallField(stage, "FormEditAssocButton"))
	}
	initRes = initializerStatements.String()
	ptrRes = pointersInitializesStatements.String()
	return
}
func (option *Option) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(option.GongMarshallField(stage, "Name"))
	}
	initRes = initializerStatements.String()
	ptrRes = pointersInitializesStatements.String()
	return
}
