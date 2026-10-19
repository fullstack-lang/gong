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
	modelsPackageName := "github.com/fullstack-lang/gong/test/test2/go/models"
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

	if stage.ModelStage != nil {
		for _, submodel := range __gong__sortStageSetInstances(stage.ModelStage.SubModels, stage.ModelStage.SubModel_stagedOrder) {
			if lastStageDecl != "ModelStage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "ModelStage"
			}
			submodelIdent := "__model" + submodel.GongGetIdentifier(stage.ModelStage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&model.SubModel{Name: %s}).Stage(stage.ModelStage)", submodelIdent, __gong__toRawStringLiteral(submodel.Name)))
			if lastStageVal != "ModelStage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "ModelStage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", submodelIdent, __gong__toRawStringLiteral(submodel.Name)))
		}
	}
	if stage.YStage != nil {
		for _, y := range __gong__sortStageSetInstances(stage.YStage.Ys, stage.YStage.Y_stagedOrder) {
			if lastStageDecl != "YStage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "YStage"
			}
			yIdent := "__y" + y.GongGetIdentifier(stage.YStage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&y.Y{Name: %s}).Stage(stage.YStage)", yIdent, __gong__toRawStringLiteral(y.Name)))
			if lastStageVal != "YStage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "YStage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", yIdent, __gong__toRawStringLiteral(y.Name)))
			if y.SubModel != nil {
				if lastStagePtr != "YStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "YStage"
				}
				targetIdent := "__model" + y.SubModel.GongGetIdentifier(stage.ModelStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.SubModel = %s", yIdent, targetIdent))
			}
		}
	}
	if stage.XStage != nil {
		for _, x := range __gong__sortStageSetInstances(stage.XStage.Xs, stage.XStage.X_stagedOrder) {
			if lastStageDecl != "XStage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "XStage"
			}
			xIdent := "__x" + x.GongGetIdentifier(stage.XStage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&x.X{Name: %s}).Stage(stage.XStage)", xIdent, __gong__toRawStringLiteral(x.Name)))
			if lastStageVal != "XStage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "XStage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", xIdent, __gong__toRawStringLiteral(x.Name)))
			if x.Y != nil {
				if lastStagePtr != "XStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "XStage"
				}
				targetIdent := "__y" + x.Y.GongGetIdentifier(stage.YStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Y = %s", xIdent, targetIdent))
			}
		}
	}
	if stage != nil {
		for _, a := range __gong__sortStageSetInstances(stage.As, stage.A_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			aIdent := localAliasPrefix + a.GongGetIdentifier(stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&models.A{Name: %s}).Stage(stage)", aIdent, __gong__toRawStringLiteral(a.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", aIdent, __gong__toRawStringLiteral(a.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.NumberField = %d", aIdent, a.NumberField))
			values.WriteString(fmt.Sprintf("\n\t%s.Foo = %d", aIdent, a.Foo))
			values.WriteString(fmt.Sprintf("\n\t%s.Bar = %f", aIdent, a.Bar))
			values.WriteString(fmt.Sprintf("\n\t%s.Zorgh = %s", aIdent, __gong__toRawStringLiteral(a.Zorgh)))
			if a.B != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := localAliasPrefix + a.B.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.B = %s", aIdent, targetIdent))
			}
			for _, elem := range a.Bs {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := localAliasPrefix + elem.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Bs = append(%s.Bs, %s)", aIdent, aIdent, targetIdent))
			}
			if a.X != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := "__x" + a.X.GongGetIdentifier(stage.XStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.X = %s", aIdent, targetIdent))
			}
		}
	}
	if stage != nil {
		for _, b := range __gong__sortStageSetInstances(stage.Bs, stage.B_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			bIdent := localAliasPrefix + b.GongGetIdentifier(stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&models.B{Name: %s}).Stage(stage)", bIdent, __gong__toRawStringLiteral(b.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", bIdent, __gong__toRawStringLiteral(b.Name)))
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
go run github.com/fullstack-lang/gong/test/test2/go/models/probe/cmd/probe@%s stage.go
*/
package %s

import (
	"slices"
	"time"

	"github.com/fullstack-lang/gong/test/test2/go/models"
	"github.com/fullstack-lang/gong/test/test2/go/models/x"
	"github.com/fullstack-lang/gong/test/test2/go/models/y"
	model "github.com/fullstack-lang/gong/test/test2/go/models/x/models"%s
)

var (
	_ time.Time
	_ = slices.Index[[]int, int]

	_ *models.Stage
	_ *x.Stage
	_ *y.Stage
	_ *model.Stage
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

	if modelsPackageName != "" && modelsPackageName != "github.com/fullstack-lang/gong/test/test2/go/models" {
		res = strings.ReplaceAll(res, "\"github.com/fullstack-lang/gong/test/test2/go/models\"", fmt.Sprintf("%q", modelsPackageName))
	}

	return res, nil
}

// insertion point for marshall field methods
func (a *A) GongMarshallField(stage *Stage, fieldName string) (res string) {
	ident := a.GongGetIdentifier(stage)
	_ = ident
	switch fieldName {
	case "Name":
		res = __gong__marshallString(ident, "Name", a.Name)
	case "NumberField":
		res = __gong__marshallInt(ident, "NumberField", a.NumberField)
	case "Foo":
		res = __gong__marshallInt(ident, "Foo", a.Foo)
	case "Bar":
		res = __gong__marshallFloat(ident, "Bar", a.Bar)
	case "Zorgh":
		res = __gong__marshallString(ident, "Zorgh", a.Zorgh)

	case "B":
		if a.B != nil {
			res = __gong__marshallPointer(ident, "B", a.B.GongGetIdentifier(stage))
		} else {
			res = __gong__marshallPointer(ident, "B", "nil")
		}
	case "Bs":
		var sb strings.Builder
		for _, _b := range a.Bs {
			sb.WriteString(__gong__marshallSliceOfPointers(ident, "Bs", _b.GongGetIdentifier(stage)))
		}
		res = sb.String()
	default:
		log.Panicf("Unknown field %s for Gongstruct A", fieldName)
	}
	return
}

func (b *B) GongMarshallField(stage *Stage, fieldName string) (res string) {
	ident := b.GongGetIdentifier(stage)
	_ = ident
	switch fieldName {
	case "Name":
		res = __gong__marshallString(ident, "Name", b.Name)

	default:
		log.Panicf("Unknown field %s for Gongstruct B", fieldName)
	}
	return
}

// insertion point for marshall all fields methods
func (a *A) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(a.GongMarshallField(stage, "Name"))
		initializerStatements.WriteString(a.GongMarshallField(stage, "NumberField"))
		pointersInitializesStatements.WriteString(a.GongMarshallField(stage, "B"))
		pointersInitializesStatements.WriteString(a.GongMarshallField(stage, "Bs"))
		initializerStatements.WriteString(a.GongMarshallField(stage, "Foo"))
		initializerStatements.WriteString(a.GongMarshallField(stage, "Bar"))
		initializerStatements.WriteString(a.GongMarshallField(stage, "Zorgh"))
	}
	initRes = initializerStatements.String()
	ptrRes = pointersInitializesStatements.String()
	return
}
func (b *B) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(b.GongMarshallField(stage, "Name"))
	}
	initRes = initializerStatements.String()
	ptrRes = pointersInitializesStatements.String()
	return
}
