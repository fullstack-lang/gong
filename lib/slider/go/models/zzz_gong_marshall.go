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
	modelsPackageName := "github.com/fullstack-lang/gong/lib/slider/go/models"
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
		for _, checkbox := range __gong__sortStageSetInstances(stage.Checkboxs, stage.Checkbox_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			checkboxIdent := localAliasPrefix + checkbox.GongGetIdentifier(stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&models.Checkbox{Name: %s}).Stage(stage)", checkboxIdent, __gong__toRawStringLiteral(checkbox.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", checkboxIdent, __gong__toRawStringLiteral(checkbox.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.ValueBool = %t", checkboxIdent, checkbox.ValueBool))
			values.WriteString(fmt.Sprintf("\n\t%s.LabelForTrue = %s", checkboxIdent, __gong__toRawStringLiteral(checkbox.LabelForTrue)))
			values.WriteString(fmt.Sprintf("\n\t%s.LabelForFalse = %s", checkboxIdent, __gong__toRawStringLiteral(checkbox.LabelForFalse)))
		}
	}
	if stage != nil {
		for _, group := range __gong__sortStageSetInstances(stage.Groups, stage.Group_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			groupIdent := localAliasPrefix + group.GongGetIdentifier(stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&models.Group{Name: %s}).Stage(stage)", groupIdent, __gong__toRawStringLiteral(group.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", groupIdent, __gong__toRawStringLiteral(group.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Percentage = %f", groupIdent, group.Percentage))
			for _, elem := range group.Sliders {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := localAliasPrefix + elem.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Sliders = append(%s.Sliders, %s)", groupIdent, groupIdent, targetIdent))
			}
			for _, elem := range group.Checkboxes {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := localAliasPrefix + elem.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Checkboxes = append(%s.Checkboxes, %s)", groupIdent, groupIdent, targetIdent))
			}
		}
	}
	if stage != nil {
		for _, layout := range __gong__sortStageSetInstances(stage.Layouts, stage.Layout_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			layoutIdent := localAliasPrefix + layout.GongGetIdentifier(stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&models.Layout{Name: %s}).Stage(stage)", layoutIdent, __gong__toRawStringLiteral(layout.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", layoutIdent, __gong__toRawStringLiteral(layout.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.IsWithCustomGutterSize = %t", layoutIdent, layout.IsWithCustomGutterSize))
			values.WriteString(fmt.Sprintf("\n\t%s.GutterSize = %f", layoutIdent, layout.GutterSize))
			for _, elem := range layout.Groups {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := localAliasPrefix + elem.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Groups = append(%s.Groups, %s)", layoutIdent, layoutIdent, targetIdent))
			}
		}
	}
	if stage != nil {
		for _, slider := range __gong__sortStageSetInstances(stage.Sliders, stage.Slider_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			sliderIdent := localAliasPrefix + slider.GongGetIdentifier(stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&models.Slider{Name: %s}).Stage(stage)", sliderIdent, __gong__toRawStringLiteral(slider.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", sliderIdent, __gong__toRawStringLiteral(slider.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.IsFloat64 = %t", sliderIdent, slider.IsFloat64))
			values.WriteString(fmt.Sprintf("\n\t%s.IsInt = %t", sliderIdent, slider.IsInt))
			values.WriteString(fmt.Sprintf("\n\t%s.MinInt = %d", sliderIdent, slider.MinInt))
			values.WriteString(fmt.Sprintf("\n\t%s.MaxInt = %d", sliderIdent, slider.MaxInt))
			values.WriteString(fmt.Sprintf("\n\t%s.StepInt = %d", sliderIdent, slider.StepInt))
			values.WriteString(fmt.Sprintf("\n\t%s.ValueInt = %d", sliderIdent, slider.ValueInt))
			values.WriteString(fmt.Sprintf("\n\t%s.MinFloat64 = %f", sliderIdent, slider.MinFloat64))
			values.WriteString(fmt.Sprintf("\n\t%s.MaxFloat64 = %f", sliderIdent, slider.MaxFloat64))
			values.WriteString(fmt.Sprintf("\n\t%s.StepFloat64 = %f", sliderIdent, slider.StepFloat64))
			values.WriteString(fmt.Sprintf("\n\t%s.ValueFloat64 = %f", sliderIdent, slider.ValueFloat64))
			values.WriteString(fmt.Sprintf("\n\t%s.IsDisabled = %t", sliderIdent, slider.IsDisabled))
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
go run github.com/fullstack-lang/gong/lib/slider/go/models/probe/cmd/probe@%s stage.go
*/
package %s

import (
	"slices"
	"time"

	"github.com/fullstack-lang/gong/lib/slider/go/models"%s
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

	if modelsPackageName != "" && modelsPackageName != "github.com/fullstack-lang/gong/lib/slider/go/models" {
		res = strings.ReplaceAll(res, "\"github.com/fullstack-lang/gong/lib/slider/go/models\"", fmt.Sprintf("%q", modelsPackageName))
	}

	return res, nil
}

// insertion point for marshall field methods
func (checkbox *Checkbox) GongMarshallField(stage *Stage, fieldName string) (res string) {
	ident := checkbox.GongGetIdentifier(stage)
	_ = ident
	switch fieldName {
	case "Name":
		res = __gong__marshallString(ident, "Name", checkbox.Name)
	case "ValueBool":
		res = __gong__marshallBool(ident, "ValueBool", checkbox.ValueBool)
	case "LabelForTrue":
		res = __gong__marshallString(ident, "LabelForTrue", checkbox.LabelForTrue)
	case "LabelForFalse":
		res = __gong__marshallString(ident, "LabelForFalse", checkbox.LabelForFalse)

	default:
		log.Panicf("Unknown field %s for Gongstruct Checkbox", fieldName)
	}
	return
}

func (group *Group) GongMarshallField(stage *Stage, fieldName string) (res string) {
	ident := group.GongGetIdentifier(stage)
	_ = ident
	switch fieldName {
	case "Name":
		res = __gong__marshallString(ident, "Name", group.Name)
	case "Percentage":
		res = __gong__marshallFloat(ident, "Percentage", group.Percentage)

	case "Sliders":
		var sb strings.Builder
		for _, _slider := range group.Sliders {
			sb.WriteString(__gong__marshallSliceOfPointers(ident, "Sliders", _slider.GongGetIdentifier(stage)))
		}
		res = sb.String()
	case "Checkboxes":
		var sb strings.Builder
		for _, _checkbox := range group.Checkboxes {
			sb.WriteString(__gong__marshallSliceOfPointers(ident, "Checkboxes", _checkbox.GongGetIdentifier(stage)))
		}
		res = sb.String()
	default:
		log.Panicf("Unknown field %s for Gongstruct Group", fieldName)
	}
	return
}

func (layout *Layout) GongMarshallField(stage *Stage, fieldName string) (res string) {
	ident := layout.GongGetIdentifier(stage)
	_ = ident
	switch fieldName {
	case "Name":
		res = __gong__marshallString(ident, "Name", layout.Name)
	case "IsWithCustomGutterSize":
		res = __gong__marshallBool(ident, "IsWithCustomGutterSize", layout.IsWithCustomGutterSize)
	case "GutterSize":
		res = __gong__marshallFloat(ident, "GutterSize", layout.GutterSize)

	case "Groups":
		var sb strings.Builder
		for _, _group := range layout.Groups {
			sb.WriteString(__gong__marshallSliceOfPointers(ident, "Groups", _group.GongGetIdentifier(stage)))
		}
		res = sb.String()
	default:
		log.Panicf("Unknown field %s for Gongstruct Layout", fieldName)
	}
	return
}

func (slider *Slider) GongMarshallField(stage *Stage, fieldName string) (res string) {
	ident := slider.GongGetIdentifier(stage)
	_ = ident
	switch fieldName {
	case "Name":
		res = __gong__marshallString(ident, "Name", slider.Name)
	case "IsFloat64":
		res = __gong__marshallBool(ident, "IsFloat64", slider.IsFloat64)
	case "IsInt":
		res = __gong__marshallBool(ident, "IsInt", slider.IsInt)
	case "MinInt":
		res = __gong__marshallInt(ident, "MinInt", slider.MinInt)
	case "MaxInt":
		res = __gong__marshallInt(ident, "MaxInt", slider.MaxInt)
	case "StepInt":
		res = __gong__marshallInt(ident, "StepInt", slider.StepInt)
	case "ValueInt":
		res = __gong__marshallInt(ident, "ValueInt", slider.ValueInt)
	case "MinFloat64":
		res = __gong__marshallFloat(ident, "MinFloat64", slider.MinFloat64)
	case "MaxFloat64":
		res = __gong__marshallFloat(ident, "MaxFloat64", slider.MaxFloat64)
	case "StepFloat64":
		res = __gong__marshallFloat(ident, "StepFloat64", slider.StepFloat64)
	case "ValueFloat64":
		res = __gong__marshallFloat(ident, "ValueFloat64", slider.ValueFloat64)
	case "IsDisabled":
		res = __gong__marshallBool(ident, "IsDisabled", slider.IsDisabled)

	default:
		log.Panicf("Unknown field %s for Gongstruct Slider", fieldName)
	}
	return
}

// insertion point for marshall all fields methods
func (checkbox *Checkbox) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(checkbox.GongMarshallField(stage, "Name"))
		initializerStatements.WriteString(checkbox.GongMarshallField(stage, "ValueBool"))
		initializerStatements.WriteString(checkbox.GongMarshallField(stage, "LabelForTrue"))
		initializerStatements.WriteString(checkbox.GongMarshallField(stage, "LabelForFalse"))
	}
	initRes = initializerStatements.String()
	ptrRes = pointersInitializesStatements.String()
	return
}
func (group *Group) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(group.GongMarshallField(stage, "Name"))
		initializerStatements.WriteString(group.GongMarshallField(stage, "Percentage"))
		pointersInitializesStatements.WriteString(group.GongMarshallField(stage, "Sliders"))
		pointersInitializesStatements.WriteString(group.GongMarshallField(stage, "Checkboxes"))
	}
	initRes = initializerStatements.String()
	ptrRes = pointersInitializesStatements.String()
	return
}
func (layout *Layout) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(layout.GongMarshallField(stage, "Name"))
		pointersInitializesStatements.WriteString(layout.GongMarshallField(stage, "Groups"))
		initializerStatements.WriteString(layout.GongMarshallField(stage, "IsWithCustomGutterSize"))
		initializerStatements.WriteString(layout.GongMarshallField(stage, "GutterSize"))
	}
	initRes = initializerStatements.String()
	ptrRes = pointersInitializesStatements.String()
	return
}
func (slider *Slider) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(slider.GongMarshallField(stage, "Name"))
		initializerStatements.WriteString(slider.GongMarshallField(stage, "IsFloat64"))
		initializerStatements.WriteString(slider.GongMarshallField(stage, "IsInt"))
		initializerStatements.WriteString(slider.GongMarshallField(stage, "MinInt"))
		initializerStatements.WriteString(slider.GongMarshallField(stage, "MaxInt"))
		initializerStatements.WriteString(slider.GongMarshallField(stage, "StepInt"))
		initializerStatements.WriteString(slider.GongMarshallField(stage, "ValueInt"))
		initializerStatements.WriteString(slider.GongMarshallField(stage, "MinFloat64"))
		initializerStatements.WriteString(slider.GongMarshallField(stage, "MaxFloat64"))
		initializerStatements.WriteString(slider.GongMarshallField(stage, "StepFloat64"))
		initializerStatements.WriteString(slider.GongMarshallField(stage, "ValueFloat64"))
		initializerStatements.WriteString(slider.GongMarshallField(stage, "IsDisabled"))
	}
	initRes = initializerStatements.String()
	ptrRes = pointersInitializesStatements.String()
	return
}
