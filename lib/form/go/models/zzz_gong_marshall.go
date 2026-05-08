// generated code - do not edit
package models

import (
	"fmt"
	"log"
	"os"
	"sort"
	"strings"
)

const marshallRes = `package {{PackageName}}

import (
	"slices"
	"time"

	"{{ModelsPackageName}}"
	// injection point for ident package import declaration{{ImportPackageDeclaration}}
)

// generated in order to avoid error in the package import
// if there are no elements in the stage to marshall
var (
	_ time.Time
	_ = slices.Index[[]int, int]
)

// _ point for meta package dummy declaration{{ImportPackageDummyDeclaration}}

// When parsed, those maps will help with the renaming process
var _ map[string]any = map[string]any{
	// injection point for docLink to identifiers{{EntriesDocLinkStringDocLinkIdentifier}}
}

// function will stage objects
func _(stage *models.Stage) {

	// insertion point for declaration of instances to stage{{Identifiers}}

	// insertion point for initialization of values{{ValueInitializers}}

	// insertion point for setup of pointers{{PointersInitializers}}
}`

const GongIdentifiersDecls = `
	{{Identifier}} := (&models.{{GeneratedStructName}}{Name: {{GeneratedFieldNameValue}}}).Stage(stage)`

const GongUnstageStmt = `
	{{Identifier}}.Unstage(stage)`

// previous version does not hanldle embedded structs (https://github.com/golang/go/issues/9859)
// simpler version but the name of the instance cannot be human read before the fields initialization
const IdentifiersDeclsWithoutNameInit = `
	{{Identifier}} := (&models.{{GeneratedStructName}}{}).Stage(stage)` /* */

const StringInitStatement = `
	{{Identifier}}.{{GeneratedFieldName}} = {{GeneratedFieldNameValue}}`

const MetaFieldStructInitStatement = `
	{{Identifier}}.{{GeneratedFieldName}} = ` + `{{GeneratedFieldNameValue}}`

const StringEnumInitStatement = `
	{{Identifier}}.{{GeneratedFieldName}} = {{GeneratedFieldNameValue}}`

const NumberInitStatement = `
	{{Identifier}}.{{GeneratedFieldName}} = {{GeneratedFieldNameValue}}`

const PointerFieldInitStatement = `
	{{Identifier}}.{{GeneratedFieldName}} = {{GeneratedFieldNameValue}}`

const SliceOfPointersFieldInitStatement = `
	{{Identifier}}.{{GeneratedFieldName}} = append({{Identifier}}.{{GeneratedFieldName}}, {{GeneratedFieldNameValue}})`

const TimeInitStatement = `
	{{Identifier}}.{{GeneratedFieldName}}, _ = time.Parse("2006-01-02 15:04:05.999999999 -0700 MST", "{{GeneratedFieldNameValue}}")`

// ToRawStringLiteral formats a string into safe Go source code,
// using backticks to preserve newlines and readability.
func ToRawStringLiteral(s string) string {
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

// MarshallFile marshall the stage content into a file as an instanciation into a stage
// according to the marshalling policy of the stage.
//
// In GongMarshallingAppendCommit mode, it will append the last commit to the file.
// In other modes, it will rewrite the entire file.
func (stage *Stage) MarshallFile(filename, modelsPackageName, packageName string) {

	if stage.GetGongMarshallingMode() == GongMarshallingAppendCommit {
		contentBytes, err := os.ReadFile(filename)

		// if the file does not exist, marshall the full stage
		if os.IsNotExist(err) {
			file, createErr := os.Create(filename)
			if createErr != nil {
				log.Fatal(createErr.Error())
			}
			defer file.Close()
			stage.Marshall(file, modelsPackageName, packageName)
			return
		}
		if err != nil {
			log.Fatal(err.Error())
		}

		content := string(contentBytes)

		if stage.isSquashing {
			// we squash: we want to clear the current function body
			// and let the append logic write the squashed commit
			firstBrace := strings.Index(content, "func _(stage *models.Stage) {")
			if firstBrace != -1 {
				firstBrace += len("func _(stage *models.Stage) {")
				content = content[:firstBrace] + "\n}\n"
			}
		}

		if stage.isApplyingBackwardCommit {
			// we are going backward, we need to remove the last forward commit from the file

			// because commitsBehind has been incremented before the call to this function
			// the index of the commit to remove is len(forwardCommits) - commitsBehind
			commitIndexToRemove := len(stage.forwardCommits) - stage.GetCommitsBehind()

			if commitIndexToRemove < 0 || commitIndexToRemove >= len(stage.forwardCommits) {
				return // Should not happen if history is consistent
			}

			commitToRemove := stage.forwardCommits[commitIndexToRemove]

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
					stage.Marshall(file, modelsPackageName, packageName)
				}
			}
			return // we are done for the backward case
		}

		if stage.isApplyingForwardCommit {
			// bypass the modified check
		} else if !stage.modified {
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
			stage.Marshall(file, modelsPackageName, packageName)
			return
		}

		contentBeforeBrace := content[:lastBrace]
		trimmedContentBeforeBrace := strings.TrimSpace(contentBeforeBrace)
		emptyBody := stage.isSquashing ||
			strings.HasSuffix(trimmedContentBeforeBrace, "func _(stage *models.Stage) {") ||
			strings.HasSuffix(trimmedContentBeforeBrace, "// insertion point for setup of pointers")

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

		stage.Marshall(file, modelsPackageName, packageName)
	}
}

// Marshall marshall the stage content into the file as an instanciation into a stage
func (stage *Stage) Marshall(file *os.File, modelsPackageName, packageName string) {

	name := file.Name()

	if !strings.HasSuffix(name, ".go") {
		log.Println(name + " is not a go filename")
	}

	log.Printf("Marshalling %s", name)

	res, err := stage.MarshallToString(modelsPackageName, packageName)
	if err != nil {
		log.Fatalln("Error marshalling to string:", err)
	}

	fmt.Fprintln(file, res)
}

// MarshallToString marshall the stage content into a string
func (stage *Stage) MarshallToString(modelsPackageName, packageName string) (res string, err error) {

	res = marshallRes
	res = strings.ReplaceAll(res, "{{PackageName}}", packageName)
	res = strings.ReplaceAll(res, "{{ModelsPackageName}}", modelsPackageName)

	// map of identifiers
	// var StageMapDstructIds map[*Dstruct]string
	var identifiersDecl strings.Builder
	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder

	decl := ""
	_ = decl
	setValueField := ""
	_ = setValueField

	// insertion initialization of objects to stage
	checkboxOrdered := []*CheckBox{}
	for checkbox := range stage.CheckBoxs {
		checkboxOrdered = append(checkboxOrdered, checkbox)
	}
	sort.Slice(checkboxOrdered[:], func(i, j int) bool {
		checkboxi := checkboxOrdered[i]
		checkboxj := checkboxOrdered[j]
		checkboxi_order, oki := stage.CheckBox_stagedOrder[checkboxi]
		checkboxj_order, okj := stage.CheckBox_stagedOrder[checkboxj]
		if !oki || !okj {
			log.Fatalln("unknown pointers")
		}
		return checkboxi_order < checkboxj_order
	})
	if len(checkboxOrdered) > 0 {
		identifiersDecl.WriteString("\n")
	}
	for _, checkbox := range checkboxOrdered {

		identifiersDecl.WriteString(checkbox.GongMarshallIdentifier(stage))

		initializerStatements.WriteString("\n")
		// Insertion point for basic fields value assignment
		initializerStatements.WriteString(checkbox.GongMarshallField(stage, "Name"))
		initializerStatements.WriteString(checkbox.GongMarshallField(stage, "Value"))
	}

	formdivOrdered := []*FormDiv{}
	for formdiv := range stage.FormDivs {
		formdivOrdered = append(formdivOrdered, formdiv)
	}
	sort.Slice(formdivOrdered[:], func(i, j int) bool {
		formdivi := formdivOrdered[i]
		formdivj := formdivOrdered[j]
		formdivi_order, oki := stage.FormDiv_stagedOrder[formdivi]
		formdivj_order, okj := stage.FormDiv_stagedOrder[formdivj]
		if !oki || !okj {
			log.Fatalln("unknown pointers")
		}
		return formdivi_order < formdivj_order
	})
	if len(formdivOrdered) > 0 {
		identifiersDecl.WriteString("\n")
	}
	for _, formdiv := range formdivOrdered {

		identifiersDecl.WriteString(formdiv.GongMarshallIdentifier(stage))

		initializerStatements.WriteString("\n")
		// Insertion point for basic fields value assignment
		initializerStatements.WriteString(formdiv.GongMarshallField(stage, "Name"))
		pointersInitializesStatements.WriteString(formdiv.GongMarshallField(stage, "FormFields"))
		pointersInitializesStatements.WriteString(formdiv.GongMarshallField(stage, "CheckBoxs"))
		pointersInitializesStatements.WriteString(formdiv.GongMarshallField(stage, "FormEditAssocButton"))
		pointersInitializesStatements.WriteString(formdiv.GongMarshallField(stage, "FormSortAssocButton"))
	}

	formeditassocbuttonOrdered := []*FormEditAssocButton{}
	for formeditassocbutton := range stage.FormEditAssocButtons {
		formeditassocbuttonOrdered = append(formeditassocbuttonOrdered, formeditassocbutton)
	}
	sort.Slice(formeditassocbuttonOrdered[:], func(i, j int) bool {
		formeditassocbuttoni := formeditassocbuttonOrdered[i]
		formeditassocbuttonj := formeditassocbuttonOrdered[j]
		formeditassocbuttoni_order, oki := stage.FormEditAssocButton_stagedOrder[formeditassocbuttoni]
		formeditassocbuttonj_order, okj := stage.FormEditAssocButton_stagedOrder[formeditassocbuttonj]
		if !oki || !okj {
			log.Fatalln("unknown pointers")
		}
		return formeditassocbuttoni_order < formeditassocbuttonj_order
	})
	if len(formeditassocbuttonOrdered) > 0 {
		identifiersDecl.WriteString("\n")
	}
	for _, formeditassocbutton := range formeditassocbuttonOrdered {

		identifiersDecl.WriteString(formeditassocbutton.GongMarshallIdentifier(stage))

		initializerStatements.WriteString("\n")
		// Insertion point for basic fields value assignment
		initializerStatements.WriteString(formeditassocbutton.GongMarshallField(stage, "Name"))
		initializerStatements.WriteString(formeditassocbutton.GongMarshallField(stage, "Label"))
		initializerStatements.WriteString(formeditassocbutton.GongMarshallField(stage, "AssociationStorage"))
		initializerStatements.WriteString(formeditassocbutton.GongMarshallField(stage, "HasChanged"))
		initializerStatements.WriteString(formeditassocbutton.GongMarshallField(stage, "IsForSavePurpose"))
		initializerStatements.WriteString(formeditassocbutton.GongMarshallField(stage, "HasToolTip"))
		initializerStatements.WriteString(formeditassocbutton.GongMarshallField(stage, "ToolTipText"))
		initializerStatements.WriteString(formeditassocbutton.GongMarshallField(stage, "MatTooltipShowDelay"))
	}

	formfieldOrdered := []*FormField{}
	for formfield := range stage.FormFields {
		formfieldOrdered = append(formfieldOrdered, formfield)
	}
	sort.Slice(formfieldOrdered[:], func(i, j int) bool {
		formfieldi := formfieldOrdered[i]
		formfieldj := formfieldOrdered[j]
		formfieldi_order, oki := stage.FormField_stagedOrder[formfieldi]
		formfieldj_order, okj := stage.FormField_stagedOrder[formfieldj]
		if !oki || !okj {
			log.Fatalln("unknown pointers")
		}
		return formfieldi_order < formfieldj_order
	})
	if len(formfieldOrdered) > 0 {
		identifiersDecl.WriteString("\n")
	}
	for _, formfield := range formfieldOrdered {

		identifiersDecl.WriteString(formfield.GongMarshallIdentifier(stage))

		initializerStatements.WriteString("\n")
		// Insertion point for basic fields value assignment
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

	formfielddateOrdered := []*FormFieldDate{}
	for formfielddate := range stage.FormFieldDates {
		formfielddateOrdered = append(formfielddateOrdered, formfielddate)
	}
	sort.Slice(formfielddateOrdered[:], func(i, j int) bool {
		formfielddatei := formfielddateOrdered[i]
		formfielddatej := formfielddateOrdered[j]
		formfielddatei_order, oki := stage.FormFieldDate_stagedOrder[formfielddatei]
		formfielddatej_order, okj := stage.FormFieldDate_stagedOrder[formfielddatej]
		if !oki || !okj {
			log.Fatalln("unknown pointers")
		}
		return formfielddatei_order < formfielddatej_order
	})
	if len(formfielddateOrdered) > 0 {
		identifiersDecl.WriteString("\n")
	}
	for _, formfielddate := range formfielddateOrdered {

		identifiersDecl.WriteString(formfielddate.GongMarshallIdentifier(stage))

		initializerStatements.WriteString("\n")
		// Insertion point for basic fields value assignment
		initializerStatements.WriteString(formfielddate.GongMarshallField(stage, "Name"))
		initializerStatements.WriteString(formfielddate.GongMarshallField(stage, "Value"))
	}

	formfielddatetimeOrdered := []*FormFieldDateTime{}
	for formfielddatetime := range stage.FormFieldDateTimes {
		formfielddatetimeOrdered = append(formfielddatetimeOrdered, formfielddatetime)
	}
	sort.Slice(formfielddatetimeOrdered[:], func(i, j int) bool {
		formfielddatetimei := formfielddatetimeOrdered[i]
		formfielddatetimej := formfielddatetimeOrdered[j]
		formfielddatetimei_order, oki := stage.FormFieldDateTime_stagedOrder[formfielddatetimei]
		formfielddatetimej_order, okj := stage.FormFieldDateTime_stagedOrder[formfielddatetimej]
		if !oki || !okj {
			log.Fatalln("unknown pointers")
		}
		return formfielddatetimei_order < formfielddatetimej_order
	})
	if len(formfielddatetimeOrdered) > 0 {
		identifiersDecl.WriteString("\n")
	}
	for _, formfielddatetime := range formfielddatetimeOrdered {

		identifiersDecl.WriteString(formfielddatetime.GongMarshallIdentifier(stage))

		initializerStatements.WriteString("\n")
		// Insertion point for basic fields value assignment
		initializerStatements.WriteString(formfielddatetime.GongMarshallField(stage, "Name"))
		initializerStatements.WriteString(formfielddatetime.GongMarshallField(stage, "Value"))
	}

	formfieldfloat64Ordered := []*FormFieldFloat64{}
	for formfieldfloat64 := range stage.FormFieldFloat64s {
		formfieldfloat64Ordered = append(formfieldfloat64Ordered, formfieldfloat64)
	}
	sort.Slice(formfieldfloat64Ordered[:], func(i, j int) bool {
		formfieldfloat64i := formfieldfloat64Ordered[i]
		formfieldfloat64j := formfieldfloat64Ordered[j]
		formfieldfloat64i_order, oki := stage.FormFieldFloat64_stagedOrder[formfieldfloat64i]
		formfieldfloat64j_order, okj := stage.FormFieldFloat64_stagedOrder[formfieldfloat64j]
		if !oki || !okj {
			log.Fatalln("unknown pointers")
		}
		return formfieldfloat64i_order < formfieldfloat64j_order
	})
	if len(formfieldfloat64Ordered) > 0 {
		identifiersDecl.WriteString("\n")
	}
	for _, formfieldfloat64 := range formfieldfloat64Ordered {

		identifiersDecl.WriteString(formfieldfloat64.GongMarshallIdentifier(stage))

		initializerStatements.WriteString("\n")
		// Insertion point for basic fields value assignment
		initializerStatements.WriteString(formfieldfloat64.GongMarshallField(stage, "Name"))
		initializerStatements.WriteString(formfieldfloat64.GongMarshallField(stage, "Value"))
		initializerStatements.WriteString(formfieldfloat64.GongMarshallField(stage, "HasMinValidator"))
		initializerStatements.WriteString(formfieldfloat64.GongMarshallField(stage, "MinValue"))
		initializerStatements.WriteString(formfieldfloat64.GongMarshallField(stage, "HasMaxValidator"))
		initializerStatements.WriteString(formfieldfloat64.GongMarshallField(stage, "MaxValue"))
	}

	formfieldintOrdered := []*FormFieldInt{}
	for formfieldint := range stage.FormFieldInts {
		formfieldintOrdered = append(formfieldintOrdered, formfieldint)
	}
	sort.Slice(formfieldintOrdered[:], func(i, j int) bool {
		formfieldinti := formfieldintOrdered[i]
		formfieldintj := formfieldintOrdered[j]
		formfieldinti_order, oki := stage.FormFieldInt_stagedOrder[formfieldinti]
		formfieldintj_order, okj := stage.FormFieldInt_stagedOrder[formfieldintj]
		if !oki || !okj {
			log.Fatalln("unknown pointers")
		}
		return formfieldinti_order < formfieldintj_order
	})
	if len(formfieldintOrdered) > 0 {
		identifiersDecl.WriteString("\n")
	}
	for _, formfieldint := range formfieldintOrdered {

		identifiersDecl.WriteString(formfieldint.GongMarshallIdentifier(stage))

		initializerStatements.WriteString("\n")
		// Insertion point for basic fields value assignment
		initializerStatements.WriteString(formfieldint.GongMarshallField(stage, "Name"))
		initializerStatements.WriteString(formfieldint.GongMarshallField(stage, "Value"))
		initializerStatements.WriteString(formfieldint.GongMarshallField(stage, "HasMinValidator"))
		initializerStatements.WriteString(formfieldint.GongMarshallField(stage, "MinValue"))
		initializerStatements.WriteString(formfieldint.GongMarshallField(stage, "HasMaxValidator"))
		initializerStatements.WriteString(formfieldint.GongMarshallField(stage, "MaxValue"))
	}

	formfieldselectOrdered := []*FormFieldSelect{}
	for formfieldselect := range stage.FormFieldSelects {
		formfieldselectOrdered = append(formfieldselectOrdered, formfieldselect)
	}
	sort.Slice(formfieldselectOrdered[:], func(i, j int) bool {
		formfieldselecti := formfieldselectOrdered[i]
		formfieldselectj := formfieldselectOrdered[j]
		formfieldselecti_order, oki := stage.FormFieldSelect_stagedOrder[formfieldselecti]
		formfieldselectj_order, okj := stage.FormFieldSelect_stagedOrder[formfieldselectj]
		if !oki || !okj {
			log.Fatalln("unknown pointers")
		}
		return formfieldselecti_order < formfieldselectj_order
	})
	if len(formfieldselectOrdered) > 0 {
		identifiersDecl.WriteString("\n")
	}
	for _, formfieldselect := range formfieldselectOrdered {

		identifiersDecl.WriteString(formfieldselect.GongMarshallIdentifier(stage))

		initializerStatements.WriteString("\n")
		// Insertion point for basic fields value assignment
		initializerStatements.WriteString(formfieldselect.GongMarshallField(stage, "Name"))
		pointersInitializesStatements.WriteString(formfieldselect.GongMarshallField(stage, "Value"))
		pointersInitializesStatements.WriteString(formfieldselect.GongMarshallField(stage, "Options"))
		initializerStatements.WriteString(formfieldselect.GongMarshallField(stage, "CanBeEmpty"))
		initializerStatements.WriteString(formfieldselect.GongMarshallField(stage, "PreserveInitialOrder"))
	}

	formfieldstringOrdered := []*FormFieldString{}
	for formfieldstring := range stage.FormFieldStrings {
		formfieldstringOrdered = append(formfieldstringOrdered, formfieldstring)
	}
	sort.Slice(formfieldstringOrdered[:], func(i, j int) bool {
		formfieldstringi := formfieldstringOrdered[i]
		formfieldstringj := formfieldstringOrdered[j]
		formfieldstringi_order, oki := stage.FormFieldString_stagedOrder[formfieldstringi]
		formfieldstringj_order, okj := stage.FormFieldString_stagedOrder[formfieldstringj]
		if !oki || !okj {
			log.Fatalln("unknown pointers")
		}
		return formfieldstringi_order < formfieldstringj_order
	})
	if len(formfieldstringOrdered) > 0 {
		identifiersDecl.WriteString("\n")
	}
	for _, formfieldstring := range formfieldstringOrdered {

		identifiersDecl.WriteString(formfieldstring.GongMarshallIdentifier(stage))

		initializerStatements.WriteString("\n")
		// Insertion point for basic fields value assignment
		initializerStatements.WriteString(formfieldstring.GongMarshallField(stage, "Name"))
		initializerStatements.WriteString(formfieldstring.GongMarshallField(stage, "Value"))
		initializerStatements.WriteString(formfieldstring.GongMarshallField(stage, "IsTextArea"))
	}

	formfieldtimeOrdered := []*FormFieldTime{}
	for formfieldtime := range stage.FormFieldTimes {
		formfieldtimeOrdered = append(formfieldtimeOrdered, formfieldtime)
	}
	sort.Slice(formfieldtimeOrdered[:], func(i, j int) bool {
		formfieldtimei := formfieldtimeOrdered[i]
		formfieldtimej := formfieldtimeOrdered[j]
		formfieldtimei_order, oki := stage.FormFieldTime_stagedOrder[formfieldtimei]
		formfieldtimej_order, okj := stage.FormFieldTime_stagedOrder[formfieldtimej]
		if !oki || !okj {
			log.Fatalln("unknown pointers")
		}
		return formfieldtimei_order < formfieldtimej_order
	})
	if len(formfieldtimeOrdered) > 0 {
		identifiersDecl.WriteString("\n")
	}
	for _, formfieldtime := range formfieldtimeOrdered {

		identifiersDecl.WriteString(formfieldtime.GongMarshallIdentifier(stage))

		initializerStatements.WriteString("\n")
		// Insertion point for basic fields value assignment
		initializerStatements.WriteString(formfieldtime.GongMarshallField(stage, "Name"))
		initializerStatements.WriteString(formfieldtime.GongMarshallField(stage, "Value"))
		initializerStatements.WriteString(formfieldtime.GongMarshallField(stage, "Step"))
	}

	formgroupOrdered := []*FormGroup{}
	for formgroup := range stage.FormGroups {
		formgroupOrdered = append(formgroupOrdered, formgroup)
	}
	sort.Slice(formgroupOrdered[:], func(i, j int) bool {
		formgroupi := formgroupOrdered[i]
		formgroupj := formgroupOrdered[j]
		formgroupi_order, oki := stage.FormGroup_stagedOrder[formgroupi]
		formgroupj_order, okj := stage.FormGroup_stagedOrder[formgroupj]
		if !oki || !okj {
			log.Fatalln("unknown pointers")
		}
		return formgroupi_order < formgroupj_order
	})
	if len(formgroupOrdered) > 0 {
		identifiersDecl.WriteString("\n")
	}
	for _, formgroup := range formgroupOrdered {

		identifiersDecl.WriteString(formgroup.GongMarshallIdentifier(stage))

		initializerStatements.WriteString("\n")
		// Insertion point for basic fields value assignment
		initializerStatements.WriteString(formgroup.GongMarshallField(stage, "Name"))
		initializerStatements.WriteString(formgroup.GongMarshallField(stage, "Label"))
		pointersInitializesStatements.WriteString(formgroup.GongMarshallField(stage, "FormDivs"))
		initializerStatements.WriteString(formgroup.GongMarshallField(stage, "HasSuppressButton"))
		initializerStatements.WriteString(formgroup.GongMarshallField(stage, "HasSuppressButtonBeenPressed"))
	}

	formsortassocbuttonOrdered := []*FormSortAssocButton{}
	for formsortassocbutton := range stage.FormSortAssocButtons {
		formsortassocbuttonOrdered = append(formsortassocbuttonOrdered, formsortassocbutton)
	}
	sort.Slice(formsortassocbuttonOrdered[:], func(i, j int) bool {
		formsortassocbuttoni := formsortassocbuttonOrdered[i]
		formsortassocbuttonj := formsortassocbuttonOrdered[j]
		formsortassocbuttoni_order, oki := stage.FormSortAssocButton_stagedOrder[formsortassocbuttoni]
		formsortassocbuttonj_order, okj := stage.FormSortAssocButton_stagedOrder[formsortassocbuttonj]
		if !oki || !okj {
			log.Fatalln("unknown pointers")
		}
		return formsortassocbuttoni_order < formsortassocbuttonj_order
	})
	if len(formsortassocbuttonOrdered) > 0 {
		identifiersDecl.WriteString("\n")
	}
	for _, formsortassocbutton := range formsortassocbuttonOrdered {

		identifiersDecl.WriteString(formsortassocbutton.GongMarshallIdentifier(stage))

		initializerStatements.WriteString("\n")
		// Insertion point for basic fields value assignment
		initializerStatements.WriteString(formsortassocbutton.GongMarshallField(stage, "Name"))
		initializerStatements.WriteString(formsortassocbutton.GongMarshallField(stage, "Label"))
		initializerStatements.WriteString(formsortassocbutton.GongMarshallField(stage, "HasToolTip"))
		initializerStatements.WriteString(formsortassocbutton.GongMarshallField(stage, "ToolTipText"))
		initializerStatements.WriteString(formsortassocbutton.GongMarshallField(stage, "MatTooltipShowDelay"))
		pointersInitializesStatements.WriteString(formsortassocbutton.GongMarshallField(stage, "FormEditAssocButton"))
	}

	optionOrdered := []*Option{}
	for option := range stage.Options {
		optionOrdered = append(optionOrdered, option)
	}
	sort.Slice(optionOrdered[:], func(i, j int) bool {
		optioni := optionOrdered[i]
		optionj := optionOrdered[j]
		optioni_order, oki := stage.Option_stagedOrder[optioni]
		optionj_order, okj := stage.Option_stagedOrder[optionj]
		if !oki || !okj {
			log.Fatalln("unknown pointers")
		}
		return optioni_order < optionj_order
	})
	if len(optionOrdered) > 0 {
		identifiersDecl.WriteString("\n")
	}
	for _, option := range optionOrdered {

		identifiersDecl.WriteString(option.GongMarshallIdentifier(stage))

		initializerStatements.WriteString("\n")
		// Insertion point for basic fields value assignment
		initializerStatements.WriteString(option.GongMarshallField(stage, "Name"))
	}

	// insertion initialization of objects to stage
	for _, checkbox := range checkboxOrdered {
		_ = checkbox
		var setPointerField string
		_ = setPointerField

		// Insertion point for pointers initialization
	}

	for _, formdiv := range formdivOrdered {
		_ = formdiv
		var setPointerField string
		_ = setPointerField

		// Insertion point for pointers initialization
	}

	for _, formeditassocbutton := range formeditassocbuttonOrdered {
		_ = formeditassocbutton
		var setPointerField string
		_ = setPointerField

		// Insertion point for pointers initialization
	}

	for _, formfield := range formfieldOrdered {
		_ = formfield
		var setPointerField string
		_ = setPointerField

		// Insertion point for pointers initialization
	}

	for _, formfielddate := range formfielddateOrdered {
		_ = formfielddate
		var setPointerField string
		_ = setPointerField

		// Insertion point for pointers initialization
	}

	for _, formfielddatetime := range formfielddatetimeOrdered {
		_ = formfielddatetime
		var setPointerField string
		_ = setPointerField

		// Insertion point for pointers initialization
	}

	for _, formfieldfloat64 := range formfieldfloat64Ordered {
		_ = formfieldfloat64
		var setPointerField string
		_ = setPointerField

		// Insertion point for pointers initialization
	}

	for _, formfieldint := range formfieldintOrdered {
		_ = formfieldint
		var setPointerField string
		_ = setPointerField

		// Insertion point for pointers initialization
	}

	for _, formfieldselect := range formfieldselectOrdered {
		_ = formfieldselect
		var setPointerField string
		_ = setPointerField

		// Insertion point for pointers initialization
	}

	for _, formfieldstring := range formfieldstringOrdered {
		_ = formfieldstring
		var setPointerField string
		_ = setPointerField

		// Insertion point for pointers initialization
	}

	for _, formfieldtime := range formfieldtimeOrdered {
		_ = formfieldtime
		var setPointerField string
		_ = setPointerField

		// Insertion point for pointers initialization
	}

	for _, formgroup := range formgroupOrdered {
		_ = formgroup
		var setPointerField string
		_ = setPointerField

		// Insertion point for pointers initialization
	}

	for _, formsortassocbutton := range formsortassocbuttonOrdered {
		_ = formsortassocbutton
		var setPointerField string
		_ = setPointerField

		// Insertion point for pointers initialization
	}

	for _, option := range optionOrdered {
		_ = option
		var setPointerField string
		_ = setPointerField

		// Insertion point for pointers initialization
	}

	res = strings.ReplaceAll(res, "{{Identifiers}}", identifiersDecl.String())
	res = strings.ReplaceAll(res, "{{ValueInitializers}}", initializerStatements.String())
	res = strings.ReplaceAll(res, "{{PointersInitializers}}", pointersInitializesStatements.String())

	if stage.MetaPackageImportAlias != "" {
		res = strings.ReplaceAll(res, "{{ImportPackageDeclaration}}",
			fmt.Sprintf("\n\t%s %s", stage.MetaPackageImportAlias, stage.MetaPackageImportPath))

		res = strings.ReplaceAll(res, "{{ImportPackageDummyDeclaration}}",
			fmt.Sprintf("\nvar _ %s.Stage",
				stage.MetaPackageImportAlias))

		var entries strings.Builder

		// regenerate the map of doc link renaming
		// the key and value are set to the value because
		// if it has been renamed, this is the new value that matters
		valuesOrdered := make([]GONG__Identifier, 0)
		for _, value := range stage.Map_DocLink_Renaming {
			valuesOrdered = append(valuesOrdered, value)
		}
		sort.Slice(valuesOrdered[:], func(i, j int) bool {
			return valuesOrdered[i].Ident < valuesOrdered[j].Ident
		})
		for _, value := range valuesOrdered {

			// get the number of points in the value to find if it is a field
			// or a struct

			switch value.Type {
			case GONG__ENUM_CAST_INT:
				entries.WriteString(fmt.Sprintf("\n\n\t\"%s\": %s(0),", value.Ident, value.Ident))
			case GONG__ENUM_CAST_STRING:
				entries.WriteString(fmt.Sprintf("\n\n\t\"%s\": %s(\"\"),", value.Ident, value.Ident))
			case GONG__FIELD_VALUE:
				// substitute the second point with "{})."
				joker := "__substitute_for_first_point__"
				valueIdentifier := strings.Replace(value.Ident, ".", joker, 1)
				valueIdentifier = strings.Replace(valueIdentifier, ".", "{}).", 1)
				valueIdentifier = strings.Replace(valueIdentifier, joker, ".", 1)
				entries.WriteString(fmt.Sprintf("\n\n\t\"%s\": (%s,", value.Ident, valueIdentifier))
			case GONG__IDENTIFIER_CONST:
				entries.WriteString(fmt.Sprintf("\n\n\t\"%s\": %s,", value.Ident, value.Ident))
			case GONG__STRUCT_INSTANCE:
				entries.WriteString(fmt.Sprintf("\n\n\t\"%s\": &(%s{}),", value.Ident, value.Ident))
			}
		}

		// res = strings.ReplaceAll(res, "{{EntriesDocLinkStringDocLinkIdentifier}}", entries.String())
	}
	return
}

// insertion point for marshall field methods
func (checkbox *CheckBox) GongMarshallField(stage *Stage, fieldName string) (res string) {

	switch fieldName {
	case "Name":
		res = StringInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", checkbox.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "Name")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", ToRawStringLiteral(checkbox.Name))
	case "Value":
		res = NumberInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", checkbox.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "Value")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%t", checkbox.Value))

	default:
		log.Panicf("Unknown field %s for Gongstruct CheckBox", fieldName)
	}
	return
}

func (formdiv *FormDiv) GongMarshallField(stage *Stage, fieldName string) (res string) {

	switch fieldName {
	case "Name":
		res = StringInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", formdiv.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "Name")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", ToRawStringLiteral(formdiv.Name))

	case "FormFields":
		var sb strings.Builder
		for _, _formfield := range formdiv.FormFields {
			tmp := SliceOfPointersFieldInitStatement
			tmp = strings.ReplaceAll(tmp, "{{Identifier}}", formdiv.GongGetIdentifier(stage))
			tmp = strings.ReplaceAll(tmp, "{{GeneratedFieldName}}", "FormFields")
			tmp = strings.ReplaceAll(tmp, "{{GeneratedFieldNameValue}}", _formfield.GongGetIdentifier(stage))
			sb.WriteString(tmp)
		}
		res = sb.String()
	case "CheckBoxs":
		var sb strings.Builder
		for _, _checkbox := range formdiv.CheckBoxs {
			tmp := SliceOfPointersFieldInitStatement
			tmp = strings.ReplaceAll(tmp, "{{Identifier}}", formdiv.GongGetIdentifier(stage))
			tmp = strings.ReplaceAll(tmp, "{{GeneratedFieldName}}", "CheckBoxs")
			tmp = strings.ReplaceAll(tmp, "{{GeneratedFieldNameValue}}", _checkbox.GongGetIdentifier(stage))
			sb.WriteString(tmp)
		}
		res = sb.String()
	case "FormEditAssocButton":
		if formdiv.FormEditAssocButton != nil {
			res = PointerFieldInitStatement
			res = strings.ReplaceAll(res, "{{Identifier}}", formdiv.GongGetIdentifier(stage))
			res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "FormEditAssocButton")
			res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", formdiv.FormEditAssocButton.GongGetIdentifier(stage))
		} else {
			// in case of nil pointer, we need to unstage the previous value
			res = PointerFieldInitStatement
			res = strings.ReplaceAll(res, "{{Identifier}}", formdiv.GongGetIdentifier(stage))
			res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "FormEditAssocButton")
			res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", "nil")
		}
	case "FormSortAssocButton":
		if formdiv.FormSortAssocButton != nil {
			res = PointerFieldInitStatement
			res = strings.ReplaceAll(res, "{{Identifier}}", formdiv.GongGetIdentifier(stage))
			res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "FormSortAssocButton")
			res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", formdiv.FormSortAssocButton.GongGetIdentifier(stage))
		} else {
			// in case of nil pointer, we need to unstage the previous value
			res = PointerFieldInitStatement
			res = strings.ReplaceAll(res, "{{Identifier}}", formdiv.GongGetIdentifier(stage))
			res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "FormSortAssocButton")
			res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", "nil")
		}
	default:
		log.Panicf("Unknown field %s for Gongstruct FormDiv", fieldName)
	}
	return
}

func (formeditassocbutton *FormEditAssocButton) GongMarshallField(stage *Stage, fieldName string) (res string) {

	switch fieldName {
	case "Name":
		res = StringInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", formeditassocbutton.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "Name")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", ToRawStringLiteral(formeditassocbutton.Name))
	case "Label":
		res = StringInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", formeditassocbutton.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "Label")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", ToRawStringLiteral(formeditassocbutton.Label))
	case "AssociationStorage":
		res = StringInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", formeditassocbutton.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "AssociationStorage")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", ToRawStringLiteral(formeditassocbutton.AssociationStorage))
	case "HasChanged":
		res = NumberInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", formeditassocbutton.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "HasChanged")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%t", formeditassocbutton.HasChanged))
	case "IsForSavePurpose":
		res = NumberInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", formeditassocbutton.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "IsForSavePurpose")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%t", formeditassocbutton.IsForSavePurpose))
	case "HasToolTip":
		res = NumberInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", formeditassocbutton.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "HasToolTip")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%t", formeditassocbutton.HasToolTip))
	case "ToolTipText":
		res = StringInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", formeditassocbutton.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "ToolTipText")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", ToRawStringLiteral(formeditassocbutton.ToolTipText))
	case "MatTooltipShowDelay":
		res = StringInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", formeditassocbutton.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "MatTooltipShowDelay")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", ToRawStringLiteral(formeditassocbutton.MatTooltipShowDelay))

	default:
		log.Panicf("Unknown field %s for Gongstruct FormEditAssocButton", fieldName)
	}
	return
}

func (formfield *FormField) GongMarshallField(stage *Stage, fieldName string) (res string) {

	switch fieldName {
	case "Name":
		res = StringInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", formfield.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "Name")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", ToRawStringLiteral(formfield.Name))
	case "InputTypeEnum":
		if formfield.InputTypeEnum.ToCodeString() != "" {
			res = StringEnumInitStatement
			res = strings.ReplaceAll(res, "{{Identifier}}", formfield.GongGetIdentifier(stage))
			res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "InputTypeEnum")
			res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", "models."+formfield.InputTypeEnum.ToCodeString())
		} else {
			// in case of empty enum, we need to unstage the previous value
			res = StringEnumInitStatement
			res = strings.ReplaceAll(res, "{{Identifier}}", formfield.GongGetIdentifier(stage))
			res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "InputTypeEnum")
			res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", "\"\"")
		}
	case "Label":
		res = StringInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", formfield.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "Label")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", ToRawStringLiteral(formfield.Label))
	case "Placeholder":
		res = StringInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", formfield.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "Placeholder")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", ToRawStringLiteral(formfield.Placeholder))
	case "HasBespokeWidth":
		res = NumberInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", formfield.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "HasBespokeWidth")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%t", formfield.HasBespokeWidth))
	case "BespokeWidthPx":
		res = NumberInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", formfield.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "BespokeWidthPx")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%d", formfield.BespokeWidthPx))
	case "HasBespokeHeight":
		res = NumberInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", formfield.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "HasBespokeHeight")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%t", formfield.HasBespokeHeight))
	case "BespokeHeightPx":
		res = NumberInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", formfield.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "BespokeHeightPx")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%d", formfield.BespokeHeightPx))

	case "FormFieldString":
		if formfield.FormFieldString != nil {
			res = PointerFieldInitStatement
			res = strings.ReplaceAll(res, "{{Identifier}}", formfield.GongGetIdentifier(stage))
			res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "FormFieldString")
			res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", formfield.FormFieldString.GongGetIdentifier(stage))
		} else {
			// in case of nil pointer, we need to unstage the previous value
			res = PointerFieldInitStatement
			res = strings.ReplaceAll(res, "{{Identifier}}", formfield.GongGetIdentifier(stage))
			res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "FormFieldString")
			res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", "nil")
		}
	case "FormFieldFloat64":
		if formfield.FormFieldFloat64 != nil {
			res = PointerFieldInitStatement
			res = strings.ReplaceAll(res, "{{Identifier}}", formfield.GongGetIdentifier(stage))
			res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "FormFieldFloat64")
			res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", formfield.FormFieldFloat64.GongGetIdentifier(stage))
		} else {
			// in case of nil pointer, we need to unstage the previous value
			res = PointerFieldInitStatement
			res = strings.ReplaceAll(res, "{{Identifier}}", formfield.GongGetIdentifier(stage))
			res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "FormFieldFloat64")
			res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", "nil")
		}
	case "FormFieldInt":
		if formfield.FormFieldInt != nil {
			res = PointerFieldInitStatement
			res = strings.ReplaceAll(res, "{{Identifier}}", formfield.GongGetIdentifier(stage))
			res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "FormFieldInt")
			res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", formfield.FormFieldInt.GongGetIdentifier(stage))
		} else {
			// in case of nil pointer, we need to unstage the previous value
			res = PointerFieldInitStatement
			res = strings.ReplaceAll(res, "{{Identifier}}", formfield.GongGetIdentifier(stage))
			res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "FormFieldInt")
			res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", "nil")
		}
	case "FormFieldDate":
		if formfield.FormFieldDate != nil {
			res = PointerFieldInitStatement
			res = strings.ReplaceAll(res, "{{Identifier}}", formfield.GongGetIdentifier(stage))
			res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "FormFieldDate")
			res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", formfield.FormFieldDate.GongGetIdentifier(stage))
		} else {
			// in case of nil pointer, we need to unstage the previous value
			res = PointerFieldInitStatement
			res = strings.ReplaceAll(res, "{{Identifier}}", formfield.GongGetIdentifier(stage))
			res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "FormFieldDate")
			res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", "nil")
		}
	case "FormFieldTime":
		if formfield.FormFieldTime != nil {
			res = PointerFieldInitStatement
			res = strings.ReplaceAll(res, "{{Identifier}}", formfield.GongGetIdentifier(stage))
			res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "FormFieldTime")
			res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", formfield.FormFieldTime.GongGetIdentifier(stage))
		} else {
			// in case of nil pointer, we need to unstage the previous value
			res = PointerFieldInitStatement
			res = strings.ReplaceAll(res, "{{Identifier}}", formfield.GongGetIdentifier(stage))
			res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "FormFieldTime")
			res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", "nil")
		}
	case "FormFieldDateTime":
		if formfield.FormFieldDateTime != nil {
			res = PointerFieldInitStatement
			res = strings.ReplaceAll(res, "{{Identifier}}", formfield.GongGetIdentifier(stage))
			res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "FormFieldDateTime")
			res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", formfield.FormFieldDateTime.GongGetIdentifier(stage))
		} else {
			// in case of nil pointer, we need to unstage the previous value
			res = PointerFieldInitStatement
			res = strings.ReplaceAll(res, "{{Identifier}}", formfield.GongGetIdentifier(stage))
			res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "FormFieldDateTime")
			res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", "nil")
		}
	case "FormFieldSelect":
		if formfield.FormFieldSelect != nil {
			res = PointerFieldInitStatement
			res = strings.ReplaceAll(res, "{{Identifier}}", formfield.GongGetIdentifier(stage))
			res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "FormFieldSelect")
			res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", formfield.FormFieldSelect.GongGetIdentifier(stage))
		} else {
			// in case of nil pointer, we need to unstage the previous value
			res = PointerFieldInitStatement
			res = strings.ReplaceAll(res, "{{Identifier}}", formfield.GongGetIdentifier(stage))
			res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "FormFieldSelect")
			res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", "nil")
		}
	default:
		log.Panicf("Unknown field %s for Gongstruct FormField", fieldName)
	}
	return
}

func (formfielddate *FormFieldDate) GongMarshallField(stage *Stage, fieldName string) (res string) {

	switch fieldName {
	case "Name":
		res = StringInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", formfielddate.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "Name")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", ToRawStringLiteral(formfielddate.Name))
	case "Value":
		res = TimeInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", formfielddate.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "Value")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", formfielddate.Value.String())

	default:
		log.Panicf("Unknown field %s for Gongstruct FormFieldDate", fieldName)
	}
	return
}

func (formfielddatetime *FormFieldDateTime) GongMarshallField(stage *Stage, fieldName string) (res string) {

	switch fieldName {
	case "Name":
		res = StringInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", formfielddatetime.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "Name")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", ToRawStringLiteral(formfielddatetime.Name))
	case "Value":
		res = TimeInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", formfielddatetime.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "Value")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", formfielddatetime.Value.String())

	default:
		log.Panicf("Unknown field %s for Gongstruct FormFieldDateTime", fieldName)
	}
	return
}

func (formfieldfloat64 *FormFieldFloat64) GongMarshallField(stage *Stage, fieldName string) (res string) {

	switch fieldName {
	case "Name":
		res = StringInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", formfieldfloat64.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "Name")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", ToRawStringLiteral(formfieldfloat64.Name))
	case "Value":
		res = NumberInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", formfieldfloat64.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "Value")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%f", formfieldfloat64.Value))
	case "HasMinValidator":
		res = NumberInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", formfieldfloat64.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "HasMinValidator")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%t", formfieldfloat64.HasMinValidator))
	case "MinValue":
		res = NumberInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", formfieldfloat64.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "MinValue")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%f", formfieldfloat64.MinValue))
	case "HasMaxValidator":
		res = NumberInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", formfieldfloat64.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "HasMaxValidator")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%t", formfieldfloat64.HasMaxValidator))
	case "MaxValue":
		res = NumberInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", formfieldfloat64.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "MaxValue")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%f", formfieldfloat64.MaxValue))

	default:
		log.Panicf("Unknown field %s for Gongstruct FormFieldFloat64", fieldName)
	}
	return
}

func (formfieldint *FormFieldInt) GongMarshallField(stage *Stage, fieldName string) (res string) {

	switch fieldName {
	case "Name":
		res = StringInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", formfieldint.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "Name")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", ToRawStringLiteral(formfieldint.Name))
	case "Value":
		res = NumberInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", formfieldint.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "Value")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%d", formfieldint.Value))
	case "HasMinValidator":
		res = NumberInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", formfieldint.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "HasMinValidator")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%t", formfieldint.HasMinValidator))
	case "MinValue":
		res = NumberInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", formfieldint.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "MinValue")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%d", formfieldint.MinValue))
	case "HasMaxValidator":
		res = NumberInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", formfieldint.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "HasMaxValidator")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%t", formfieldint.HasMaxValidator))
	case "MaxValue":
		res = NumberInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", formfieldint.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "MaxValue")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%d", formfieldint.MaxValue))

	default:
		log.Panicf("Unknown field %s for Gongstruct FormFieldInt", fieldName)
	}
	return
}

func (formfieldselect *FormFieldSelect) GongMarshallField(stage *Stage, fieldName string) (res string) {

	switch fieldName {
	case "Name":
		res = StringInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", formfieldselect.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "Name")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", ToRawStringLiteral(formfieldselect.Name))
	case "CanBeEmpty":
		res = NumberInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", formfieldselect.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "CanBeEmpty")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%t", formfieldselect.CanBeEmpty))
	case "PreserveInitialOrder":
		res = NumberInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", formfieldselect.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "PreserveInitialOrder")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%t", formfieldselect.PreserveInitialOrder))

	case "Value":
		if formfieldselect.Value != nil {
			res = PointerFieldInitStatement
			res = strings.ReplaceAll(res, "{{Identifier}}", formfieldselect.GongGetIdentifier(stage))
			res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "Value")
			res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", formfieldselect.Value.GongGetIdentifier(stage))
		} else {
			// in case of nil pointer, we need to unstage the previous value
			res = PointerFieldInitStatement
			res = strings.ReplaceAll(res, "{{Identifier}}", formfieldselect.GongGetIdentifier(stage))
			res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "Value")
			res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", "nil")
		}
	case "Options":
		var sb strings.Builder
		for _, _option := range formfieldselect.Options {
			tmp := SliceOfPointersFieldInitStatement
			tmp = strings.ReplaceAll(tmp, "{{Identifier}}", formfieldselect.GongGetIdentifier(stage))
			tmp = strings.ReplaceAll(tmp, "{{GeneratedFieldName}}", "Options")
			tmp = strings.ReplaceAll(tmp, "{{GeneratedFieldNameValue}}", _option.GongGetIdentifier(stage))
			sb.WriteString(tmp)
		}
		res = sb.String()
	default:
		log.Panicf("Unknown field %s for Gongstruct FormFieldSelect", fieldName)
	}
	return
}

func (formfieldstring *FormFieldString) GongMarshallField(stage *Stage, fieldName string) (res string) {

	switch fieldName {
	case "Name":
		res = StringInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", formfieldstring.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "Name")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", ToRawStringLiteral(formfieldstring.Name))
	case "Value":
		res = StringInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", formfieldstring.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "Value")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", ToRawStringLiteral(formfieldstring.Value))
	case "IsTextArea":
		res = NumberInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", formfieldstring.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "IsTextArea")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%t", formfieldstring.IsTextArea))

	default:
		log.Panicf("Unknown field %s for Gongstruct FormFieldString", fieldName)
	}
	return
}

func (formfieldtime *FormFieldTime) GongMarshallField(stage *Stage, fieldName string) (res string) {

	switch fieldName {
	case "Name":
		res = StringInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", formfieldtime.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "Name")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", ToRawStringLiteral(formfieldtime.Name))
	case "Value":
		res = TimeInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", formfieldtime.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "Value")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", formfieldtime.Value.String())
	case "Step":
		res = NumberInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", formfieldtime.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "Step")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%f", formfieldtime.Step))

	default:
		log.Panicf("Unknown field %s for Gongstruct FormFieldTime", fieldName)
	}
	return
}

func (formgroup *FormGroup) GongMarshallField(stage *Stage, fieldName string) (res string) {

	switch fieldName {
	case "Name":
		res = StringInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", formgroup.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "Name")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", ToRawStringLiteral(formgroup.Name))
	case "Label":
		res = StringInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", formgroup.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "Label")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", ToRawStringLiteral(formgroup.Label))
	case "HasSuppressButton":
		res = NumberInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", formgroup.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "HasSuppressButton")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%t", formgroup.HasSuppressButton))
	case "HasSuppressButtonBeenPressed":
		res = NumberInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", formgroup.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "HasSuppressButtonBeenPressed")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%t", formgroup.HasSuppressButtonBeenPressed))

	case "FormDivs":
		var sb strings.Builder
		for _, _formdiv := range formgroup.FormDivs {
			tmp := SliceOfPointersFieldInitStatement
			tmp = strings.ReplaceAll(tmp, "{{Identifier}}", formgroup.GongGetIdentifier(stage))
			tmp = strings.ReplaceAll(tmp, "{{GeneratedFieldName}}", "FormDivs")
			tmp = strings.ReplaceAll(tmp, "{{GeneratedFieldNameValue}}", _formdiv.GongGetIdentifier(stage))
			sb.WriteString(tmp)
		}
		res = sb.String()
	default:
		log.Panicf("Unknown field %s for Gongstruct FormGroup", fieldName)
	}
	return
}

func (formsortassocbutton *FormSortAssocButton) GongMarshallField(stage *Stage, fieldName string) (res string) {

	switch fieldName {
	case "Name":
		res = StringInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", formsortassocbutton.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "Name")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", ToRawStringLiteral(formsortassocbutton.Name))
	case "Label":
		res = StringInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", formsortassocbutton.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "Label")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", ToRawStringLiteral(formsortassocbutton.Label))
	case "HasToolTip":
		res = NumberInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", formsortassocbutton.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "HasToolTip")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%t", formsortassocbutton.HasToolTip))
	case "ToolTipText":
		res = StringInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", formsortassocbutton.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "ToolTipText")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", ToRawStringLiteral(formsortassocbutton.ToolTipText))
	case "MatTooltipShowDelay":
		res = StringInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", formsortassocbutton.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "MatTooltipShowDelay")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", ToRawStringLiteral(formsortassocbutton.MatTooltipShowDelay))

	case "FormEditAssocButton":
		if formsortassocbutton.FormEditAssocButton != nil {
			res = PointerFieldInitStatement
			res = strings.ReplaceAll(res, "{{Identifier}}", formsortassocbutton.GongGetIdentifier(stage))
			res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "FormEditAssocButton")
			res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", formsortassocbutton.FormEditAssocButton.GongGetIdentifier(stage))
		} else {
			// in case of nil pointer, we need to unstage the previous value
			res = PointerFieldInitStatement
			res = strings.ReplaceAll(res, "{{Identifier}}", formsortassocbutton.GongGetIdentifier(stage))
			res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "FormEditAssocButton")
			res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", "nil")
		}
	default:
		log.Panicf("Unknown field %s for Gongstruct FormSortAssocButton", fieldName)
	}
	return
}

func (option *Option) GongMarshallField(stage *Stage, fieldName string) (res string) {

	switch fieldName {
	case "Name":
		res = StringInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", option.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "Name")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", ToRawStringLiteral(option.Name))

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
