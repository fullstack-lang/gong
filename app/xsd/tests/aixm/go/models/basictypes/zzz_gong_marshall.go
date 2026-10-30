// generated code - do not edit
package basictypes

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
	{{Identifier}} := (&basictypes.{{GeneratedStructName}}{Name: {{GeneratedFieldNameValue}}}).Stage(stage)`

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
			const fnSig = "func _(stage *basictypes.Stage) {"
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
			strings.HasSuffix(trimmedContentBeforeBrace, "func _(stage *basictypes.Stage) {") ||
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
	modelsPackageName := "github.com/fullstack-lang/gong/app/xsd/tests/aixm/go/models/basictypes"
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
		for _, codelisttype := range __gong__sortStageSetInstances(stage.CodeListTypes, stage.CodeListType_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			codelisttypeIdent := codelisttype.GongGetIdentifier(stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&basictypes.CodeListType{Name: %s}).Stage(stage)", codelisttypeIdent, __gong__toRawStringLiteral(codelisttype.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", codelisttypeIdent, __gong__toRawStringLiteral(codelisttype.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.CodeSpace = %s", codelisttypeIdent, __gong__toRawStringLiteral(codelisttype.CodeSpace)))
			values.WriteString(fmt.Sprintf("\n\t%s.EnclosedText = %s", codelisttypeIdent, __gong__toRawStringLiteral(codelisttype.EnclosedText)))
		}
	}
	if stage != nil {
		for _, codeornilreasonlisttype := range __gong__sortStageSetInstances(stage.CodeOrNilReasonListTypes, stage.CodeOrNilReasonListType_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			codeornilreasonlisttypeIdent := codeornilreasonlisttype.GongGetIdentifier(stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&basictypes.CodeOrNilReasonListType{Name: %s}).Stage(stage)", codeornilreasonlisttypeIdent, __gong__toRawStringLiteral(codeornilreasonlisttype.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", codeornilreasonlisttypeIdent, __gong__toRawStringLiteral(codeornilreasonlisttype.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.CodeSpace = %s", codeornilreasonlisttypeIdent, __gong__toRawStringLiteral(codeornilreasonlisttype.CodeSpace)))
			values.WriteString(fmt.Sprintf("\n\t%s.EnclosedText = %s", codeornilreasonlisttypeIdent, __gong__toRawStringLiteral(codeornilreasonlisttype.EnclosedText)))
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
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&basictypes.CodeType{Name: %s}).Stage(stage)", codetypeIdent, __gong__toRawStringLiteral(codetype.Name)))
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
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&basictypes.CodeWithAuthorityType{Name: %s}).Stage(stage)", codewithauthoritytypeIdent, __gong__toRawStringLiteral(codewithauthoritytype.Name)))
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
		for _, coordinatestype := range __gong__sortStageSetInstances(stage.CoordinatesTypes, stage.CoordinatesType_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			coordinatestypeIdent := coordinatestype.GongGetIdentifier(stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&basictypes.CoordinatesType{Name: %s}).Stage(stage)", coordinatestypeIdent, __gong__toRawStringLiteral(coordinatestype.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", coordinatestypeIdent, __gong__toRawStringLiteral(coordinatestype.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Decimal = %s", coordinatestypeIdent, __gong__toRawStringLiteral(coordinatestype.Decimal)))
			values.WriteString(fmt.Sprintf("\n\t%s.Cs = %s", coordinatestypeIdent, __gong__toRawStringLiteral(coordinatestype.Cs)))
			values.WriteString(fmt.Sprintf("\n\t%s.Ts = %s", coordinatestypeIdent, __gong__toRawStringLiteral(coordinatestype.Ts)))
			values.WriteString(fmt.Sprintf("\n\t%s.EnclosedText = %s", coordinatestypeIdent, __gong__toRawStringLiteral(coordinatestype.EnclosedText)))
		}
	}
	if stage != nil {
		for _, measurelisttype := range __gong__sortStageSetInstances(stage.MeasureListTypes, stage.MeasureListType_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			measurelisttypeIdent := measurelisttype.GongGetIdentifier(stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&basictypes.MeasureListType{Name: %s}).Stage(stage)", measurelisttypeIdent, __gong__toRawStringLiteral(measurelisttype.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", measurelisttypeIdent, __gong__toRawStringLiteral(measurelisttype.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Uom = %s", measurelisttypeIdent, __gong__toRawStringLiteral(measurelisttype.Uom)))
			values.WriteString(fmt.Sprintf("\n\t%s.EnclosedText = %s", measurelisttypeIdent, __gong__toRawStringLiteral(measurelisttype.EnclosedText)))
		}
	}
	if stage != nil {
		for _, measureornilreasonlisttype := range __gong__sortStageSetInstances(stage.MeasureOrNilReasonListTypes, stage.MeasureOrNilReasonListType_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			measureornilreasonlisttypeIdent := measureornilreasonlisttype.GongGetIdentifier(stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&basictypes.MeasureOrNilReasonListType{Name: %s}).Stage(stage)", measureornilreasonlisttypeIdent, __gong__toRawStringLiteral(measureornilreasonlisttype.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", measureornilreasonlisttypeIdent, __gong__toRawStringLiteral(measureornilreasonlisttype.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Uom = %s", measureornilreasonlisttypeIdent, __gong__toRawStringLiteral(measureornilreasonlisttype.Uom)))
			values.WriteString(fmt.Sprintf("\n\t%s.EnclosedText = %s", measureornilreasonlisttypeIdent, __gong__toRawStringLiteral(measureornilreasonlisttype.EnclosedText)))
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
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&basictypes.MeasureType{Name: %s}).Stage(stage)", measuretypeIdent, __gong__toRawStringLiteral(measuretype.Name)))
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
go run github.com/fullstack-lang/gong/app/xsd/tests/aixm/go/models/basictypes/probe/cmd/probe@%s stage.go
*/
package %s

import (
	"slices"
	"time"

	"github.com/fullstack-lang/gong/app/xsd/tests/aixm/go/models/basictypes"%s
)

var (
	_ time.Time
	_ = slices.Index[[]int, int]

	_ *basictypes.Stage
)%s

// function will stage objects across all coordinated stages
func _(stage *basictypes.Stage) {

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

	if modelsPackageName != "" && modelsPackageName != "github.com/fullstack-lang/gong/app/xsd/tests/aixm/go/models/basictypes" {
		res = strings.ReplaceAll(res, "\"github.com/fullstack-lang/gong/app/xsd/tests/aixm/go/models/basictypes\"", fmt.Sprintf("%q", modelsPackageName))
	}

	return res, nil
}

// insertion point for marshall field methods
func (codelisttype *CodeListType) GongMarshallField(stage *Stage, fieldName string) (res string) {
	ident := codelisttype.GongGetIdentifier(stage)
	_ = ident
	switch fieldName {
	case "Name":
		res = __gong__marshallString(ident, "Name", codelisttype.Name)
	case "CodeSpace":
		res = __gong__marshallString(ident, "CodeSpace", codelisttype.CodeSpace)
	case "EnclosedText":
		res = __gong__marshallString(ident, "EnclosedText", codelisttype.EnclosedText)

	default:
		log.Panicf("Unknown field %s for Gongstruct CodeListType", fieldName)
	}
	return
}

func (codeornilreasonlisttype *CodeOrNilReasonListType) GongMarshallField(stage *Stage, fieldName string) (res string) {
	ident := codeornilreasonlisttype.GongGetIdentifier(stage)
	_ = ident
	switch fieldName {
	case "Name":
		res = __gong__marshallString(ident, "Name", codeornilreasonlisttype.Name)
	case "CodeSpace":
		res = __gong__marshallString(ident, "CodeSpace", codeornilreasonlisttype.CodeSpace)
	case "EnclosedText":
		res = __gong__marshallString(ident, "EnclosedText", codeornilreasonlisttype.EnclosedText)

	default:
		log.Panicf("Unknown field %s for Gongstruct CodeOrNilReasonListType", fieldName)
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

func (coordinatestype *CoordinatesType) GongMarshallField(stage *Stage, fieldName string) (res string) {
	ident := coordinatestype.GongGetIdentifier(stage)
	_ = ident
	switch fieldName {
	case "Name":
		res = __gong__marshallString(ident, "Name", coordinatestype.Name)
	case "Decimal":
		res = __gong__marshallString(ident, "Decimal", coordinatestype.Decimal)
	case "Cs":
		res = __gong__marshallString(ident, "Cs", coordinatestype.Cs)
	case "Ts":
		res = __gong__marshallString(ident, "Ts", coordinatestype.Ts)
	case "EnclosedText":
		res = __gong__marshallString(ident, "EnclosedText", coordinatestype.EnclosedText)

	default:
		log.Panicf("Unknown field %s for Gongstruct CoordinatesType", fieldName)
	}
	return
}

func (measurelisttype *MeasureListType) GongMarshallField(stage *Stage, fieldName string) (res string) {
	ident := measurelisttype.GongGetIdentifier(stage)
	_ = ident
	switch fieldName {
	case "Name":
		res = __gong__marshallString(ident, "Name", measurelisttype.Name)
	case "Uom":
		res = __gong__marshallString(ident, "Uom", measurelisttype.Uom)
	case "EnclosedText":
		res = __gong__marshallString(ident, "EnclosedText", measurelisttype.EnclosedText)

	default:
		log.Panicf("Unknown field %s for Gongstruct MeasureListType", fieldName)
	}
	return
}

func (measureornilreasonlisttype *MeasureOrNilReasonListType) GongMarshallField(stage *Stage, fieldName string) (res string) {
	ident := measureornilreasonlisttype.GongGetIdentifier(stage)
	_ = ident
	switch fieldName {
	case "Name":
		res = __gong__marshallString(ident, "Name", measureornilreasonlisttype.Name)
	case "Uom":
		res = __gong__marshallString(ident, "Uom", measureornilreasonlisttype.Uom)
	case "EnclosedText":
		res = __gong__marshallString(ident, "EnclosedText", measureornilreasonlisttype.EnclosedText)

	default:
		log.Panicf("Unknown field %s for Gongstruct MeasureOrNilReasonListType", fieldName)
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

// insertion point for marshall all fields methods
func (codelisttype *CodeListType) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(codelisttype.GongMarshallField(stage, "Name"))
		initializerStatements.WriteString(codelisttype.GongMarshallField(stage, "CodeSpace"))
		initializerStatements.WriteString(codelisttype.GongMarshallField(stage, "EnclosedText"))
	}
	initRes = initializerStatements.String()
	ptrRes = pointersInitializesStatements.String()
	return
}
func (codeornilreasonlisttype *CodeOrNilReasonListType) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(codeornilreasonlisttype.GongMarshallField(stage, "Name"))
		initializerStatements.WriteString(codeornilreasonlisttype.GongMarshallField(stage, "CodeSpace"))
		initializerStatements.WriteString(codeornilreasonlisttype.GongMarshallField(stage, "EnclosedText"))
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
func (coordinatestype *CoordinatesType) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(coordinatestype.GongMarshallField(stage, "Name"))
		initializerStatements.WriteString(coordinatestype.GongMarshallField(stage, "Decimal"))
		initializerStatements.WriteString(coordinatestype.GongMarshallField(stage, "Cs"))
		initializerStatements.WriteString(coordinatestype.GongMarshallField(stage, "Ts"))
		initializerStatements.WriteString(coordinatestype.GongMarshallField(stage, "EnclosedText"))
	}
	initRes = initializerStatements.String()
	ptrRes = pointersInitializesStatements.String()
	return
}
func (measurelisttype *MeasureListType) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(measurelisttype.GongMarshallField(stage, "Name"))
		initializerStatements.WriteString(measurelisttype.GongMarshallField(stage, "Uom"))
		initializerStatements.WriteString(measurelisttype.GongMarshallField(stage, "EnclosedText"))
	}
	initRes = initializerStatements.String()
	ptrRes = pointersInitializesStatements.String()
	return
}
func (measureornilreasonlisttype *MeasureOrNilReasonListType) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(measureornilreasonlisttype.GongMarshallField(stage, "Name"))
		initializerStatements.WriteString(measureornilreasonlisttype.GongMarshallField(stage, "Uom"))
		initializerStatements.WriteString(measureornilreasonlisttype.GongMarshallField(stage, "EnclosedText"))
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
