// generated code - do not edit
package xlink

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
	{{Identifier}} := (&xlink.{{GeneratedStructName}}{Name: {{GeneratedFieldNameValue}}}).Stage(stage)`

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
			const fnSig = "func _(stage *xlink.Stage) {"
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
			strings.HasSuffix(trimmedContentBeforeBrace, "func _(stage *xlink.Stage) {") ||
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
	modelsPackageName := "github.com/fullstack-lang/gong/app/xsd/tests/aixm/go/models/xlink"
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
		for _, arctype := range __gong__sortStageSetInstances(stage.ArcTypes, stage.ArcType_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			arctypeIdent := arctype.GongGetIdentifier(stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&xlink.ArcType{Name: %s}).Stage(stage)", arctypeIdent, __gong__toRawStringLiteral(arctype.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", arctypeIdent, __gong__toRawStringLiteral(arctype.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Type_XSD = %s", arctypeIdent, __gong__toRawStringLiteral(arctype.Type_XSD)))
			values.WriteString(fmt.Sprintf("\n\t%s.Arcrole = %s", arctypeIdent, __gong__toRawStringLiteral(arctype.Arcrole)))
			values.WriteString(fmt.Sprintf("\n\t%s.Title = %s", arctypeIdent, __gong__toRawStringLiteral(arctype.Title)))
			values.WriteString(fmt.Sprintf("\n\t%s.Show = %s", arctypeIdent, __gong__toRawStringLiteral(arctype.Show)))
			values.WriteString(fmt.Sprintf("\n\t%s.Actuate = %s", arctypeIdent, __gong__toRawStringLiteral(arctype.Actuate)))
			values.WriteString(fmt.Sprintf("\n\t%s.From = %s", arctypeIdent, __gong__toRawStringLiteral(arctype.From)))
			values.WriteString(fmt.Sprintf("\n\t%s.To = %s", arctypeIdent, __gong__toRawStringLiteral(arctype.To)))
			for _, elem := range arctype.Title_1 {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := elem.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Title_1 = append(%s.Title_1, %s)", arctypeIdent, arctypeIdent, targetIdent))
			}
		}
	}
	if stage != nil {
		for _, extended := range __gong__sortStageSetInstances(stage.Extendeds, stage.Extended_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			extendedIdent := extended.GongGetIdentifier(stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&xlink.Extended{Name: %s}).Stage(stage)", extendedIdent, __gong__toRawStringLiteral(extended.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", extendedIdent, __gong__toRawStringLiteral(extended.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Type_XSD = %s", extendedIdent, __gong__toRawStringLiteral(extended.Type_XSD)))
			values.WriteString(fmt.Sprintf("\n\t%s.Role = %s", extendedIdent, __gong__toRawStringLiteral(extended.Role)))
			values.WriteString(fmt.Sprintf("\n\t%s.Title = %s", extendedIdent, __gong__toRawStringLiteral(extended.Title)))
			if extended.Title_1 != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := extended.Title_1.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Title_1 = %s", extendedIdent, targetIdent))
			}
			if extended.Resource != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := extended.Resource.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Resource = %s", extendedIdent, targetIdent))
			}
			if extended.Locator != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := extended.Locator.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Locator = %s", extendedIdent, targetIdent))
			}
			if extended.Arc != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := extended.Arc.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Arc = %s", extendedIdent, targetIdent))
			}
		}
	}
	if stage != nil {
		for _, locatortype := range __gong__sortStageSetInstances(stage.LocatorTypes, stage.LocatorType_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			locatortypeIdent := locatortype.GongGetIdentifier(stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&xlink.LocatorType{Name: %s}).Stage(stage)", locatortypeIdent, __gong__toRawStringLiteral(locatortype.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", locatortypeIdent, __gong__toRawStringLiteral(locatortype.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Type_XSD = %s", locatortypeIdent, __gong__toRawStringLiteral(locatortype.Type_XSD)))
			values.WriteString(fmt.Sprintf("\n\t%s.Href = %s", locatortypeIdent, __gong__toRawStringLiteral(locatortype.Href)))
			values.WriteString(fmt.Sprintf("\n\t%s.Role = %s", locatortypeIdent, __gong__toRawStringLiteral(locatortype.Role)))
			values.WriteString(fmt.Sprintf("\n\t%s.Title = %s", locatortypeIdent, __gong__toRawStringLiteral(locatortype.Title)))
			values.WriteString(fmt.Sprintf("\n\t%s.Label = %s", locatortypeIdent, __gong__toRawStringLiteral(locatortype.Label)))
			for _, elem := range locatortype.Title_1 {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := elem.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Title_1 = append(%s.Title_1, %s)", locatortypeIdent, locatortypeIdent, targetIdent))
			}
		}
	}
	if stage != nil {
		for _, resourcetype := range __gong__sortStageSetInstances(stage.ResourceTypes, stage.ResourceType_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			resourcetypeIdent := resourcetype.GongGetIdentifier(stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&xlink.ResourceType{Name: %s}).Stage(stage)", resourcetypeIdent, __gong__toRawStringLiteral(resourcetype.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", resourcetypeIdent, __gong__toRawStringLiteral(resourcetype.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Type_XSD = %s", resourcetypeIdent, __gong__toRawStringLiteral(resourcetype.Type_XSD)))
			values.WriteString(fmt.Sprintf("\n\t%s.Role = %s", resourcetypeIdent, __gong__toRawStringLiteral(resourcetype.Role)))
			values.WriteString(fmt.Sprintf("\n\t%s.Title = %s", resourcetypeIdent, __gong__toRawStringLiteral(resourcetype.Title)))
			values.WriteString(fmt.Sprintf("\n\t%s.Label = %s", resourcetypeIdent, __gong__toRawStringLiteral(resourcetype.Label)))
		}
	}
	if stage != nil {
		for _, simple := range __gong__sortStageSetInstances(stage.Simples, stage.Simple_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			simpleIdent := simple.GongGetIdentifier(stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&xlink.Simple{Name: %s}).Stage(stage)", simpleIdent, __gong__toRawStringLiteral(simple.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", simpleIdent, __gong__toRawStringLiteral(simple.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Type_XSD = %s", simpleIdent, __gong__toRawStringLiteral(simple.Type_XSD)))
			values.WriteString(fmt.Sprintf("\n\t%s.Href = %s", simpleIdent, __gong__toRawStringLiteral(simple.Href)))
			values.WriteString(fmt.Sprintf("\n\t%s.Role = %s", simpleIdent, __gong__toRawStringLiteral(simple.Role)))
			values.WriteString(fmt.Sprintf("\n\t%s.Arcrole = %s", simpleIdent, __gong__toRawStringLiteral(simple.Arcrole)))
			values.WriteString(fmt.Sprintf("\n\t%s.Title = %s", simpleIdent, __gong__toRawStringLiteral(simple.Title)))
			values.WriteString(fmt.Sprintf("\n\t%s.Show = %s", simpleIdent, __gong__toRawStringLiteral(simple.Show)))
			values.WriteString(fmt.Sprintf("\n\t%s.Actuate = %s", simpleIdent, __gong__toRawStringLiteral(simple.Actuate)))
		}
	}
	if stage != nil {
		for _, titleelttype := range __gong__sortStageSetInstances(stage.TitleEltTypes, stage.TitleEltType_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			titleelttypeIdent := titleelttype.GongGetIdentifier(stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&xlink.TitleEltType{Name: %s}).Stage(stage)", titleelttypeIdent, __gong__toRawStringLiteral(titleelttype.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", titleelttypeIdent, __gong__toRawStringLiteral(titleelttype.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Type_XSD = %s", titleelttypeIdent, __gong__toRawStringLiteral(titleelttype.Type_XSD)))
			values.WriteString(fmt.Sprintf("\n\t%s.Lang = %s", titleelttypeIdent, __gong__toRawStringLiteral(titleelttype.Lang)))
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
go run github.com/fullstack-lang/gong/app/xsd/tests/aixm/go/models/xlink/probe/cmd/probe@%s stage.go
*/
package %s

import (
	"slices"
	"time"

	"github.com/fullstack-lang/gong/app/xsd/tests/aixm/go/models/xlink"%s
)

var (
	_ time.Time
	_ = slices.Index[[]int, int]

	_ *xlink.Stage
)%s

// function will stage objects across all coordinated stages
func _(stage *xlink.Stage) {

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

	if modelsPackageName != "" && modelsPackageName != "github.com/fullstack-lang/gong/app/xsd/tests/aixm/go/models/xlink" {
		res = strings.ReplaceAll(res, "\"github.com/fullstack-lang/gong/app/xsd/tests/aixm/go/models/xlink\"", fmt.Sprintf("%q", modelsPackageName))
	}

	return res, nil
}

// insertion point for marshall field methods
func (arctype *ArcType) GongMarshallField(stage *Stage, fieldName string) (res string) {
	ident := arctype.GongGetIdentifier(stage)
	_ = ident
	switch fieldName {
	case "Name":
		res = __gong__marshallString(ident, "Name", arctype.Name)
	case "Type_XSD":
		res = __gong__marshallString(ident, "Type_XSD", arctype.Type_XSD)
	case "Arcrole":
		res = __gong__marshallString(ident, "Arcrole", arctype.Arcrole)
	case "Title":
		res = __gong__marshallString(ident, "Title", arctype.Title)
	case "Show":
		res = __gong__marshallString(ident, "Show", arctype.Show)
	case "Actuate":
		res = __gong__marshallString(ident, "Actuate", arctype.Actuate)
	case "From":
		res = __gong__marshallString(ident, "From", arctype.From)
	case "To":
		res = __gong__marshallString(ident, "To", arctype.To)

	case "Title_1":
		var sb strings.Builder
		for _, _titleelttype := range arctype.Title_1 {
			sb.WriteString(__gong__marshallSliceOfPointers(ident, "Title_1", _titleelttype.GongGetIdentifier(stage)))
		}
		res = sb.String()
	default:
		log.Panicf("Unknown field %s for Gongstruct ArcType", fieldName)
	}
	return
}

func (extended *Extended) GongMarshallField(stage *Stage, fieldName string) (res string) {
	ident := extended.GongGetIdentifier(stage)
	_ = ident
	switch fieldName {
	case "Name":
		res = __gong__marshallString(ident, "Name", extended.Name)
	case "Type_XSD":
		res = __gong__marshallString(ident, "Type_XSD", extended.Type_XSD)
	case "Role":
		res = __gong__marshallString(ident, "Role", extended.Role)
	case "Title":
		res = __gong__marshallString(ident, "Title", extended.Title)

	case "Title_1":
		if extended.Title_1 != nil {
			res = __gong__marshallPointer(ident, "Title_1", extended.Title_1.GongGetIdentifier(stage))
		} else {
			res = __gong__marshallPointer(ident, "Title_1", "nil")
		}
	case "Resource":
		if extended.Resource != nil {
			res = __gong__marshallPointer(ident, "Resource", extended.Resource.GongGetIdentifier(stage))
		} else {
			res = __gong__marshallPointer(ident, "Resource", "nil")
		}
	case "Locator":
		if extended.Locator != nil {
			res = __gong__marshallPointer(ident, "Locator", extended.Locator.GongGetIdentifier(stage))
		} else {
			res = __gong__marshallPointer(ident, "Locator", "nil")
		}
	case "Arc":
		if extended.Arc != nil {
			res = __gong__marshallPointer(ident, "Arc", extended.Arc.GongGetIdentifier(stage))
		} else {
			res = __gong__marshallPointer(ident, "Arc", "nil")
		}
	default:
		log.Panicf("Unknown field %s for Gongstruct Extended", fieldName)
	}
	return
}

func (locatortype *LocatorType) GongMarshallField(stage *Stage, fieldName string) (res string) {
	ident := locatortype.GongGetIdentifier(stage)
	_ = ident
	switch fieldName {
	case "Name":
		res = __gong__marshallString(ident, "Name", locatortype.Name)
	case "Type_XSD":
		res = __gong__marshallString(ident, "Type_XSD", locatortype.Type_XSD)
	case "Href":
		res = __gong__marshallString(ident, "Href", locatortype.Href)
	case "Role":
		res = __gong__marshallString(ident, "Role", locatortype.Role)
	case "Title":
		res = __gong__marshallString(ident, "Title", locatortype.Title)
	case "Label":
		res = __gong__marshallString(ident, "Label", locatortype.Label)

	case "Title_1":
		var sb strings.Builder
		for _, _titleelttype := range locatortype.Title_1 {
			sb.WriteString(__gong__marshallSliceOfPointers(ident, "Title_1", _titleelttype.GongGetIdentifier(stage)))
		}
		res = sb.String()
	default:
		log.Panicf("Unknown field %s for Gongstruct LocatorType", fieldName)
	}
	return
}

func (resourcetype *ResourceType) GongMarshallField(stage *Stage, fieldName string) (res string) {
	ident := resourcetype.GongGetIdentifier(stage)
	_ = ident
	switch fieldName {
	case "Name":
		res = __gong__marshallString(ident, "Name", resourcetype.Name)
	case "Type_XSD":
		res = __gong__marshallString(ident, "Type_XSD", resourcetype.Type_XSD)
	case "Role":
		res = __gong__marshallString(ident, "Role", resourcetype.Role)
	case "Title":
		res = __gong__marshallString(ident, "Title", resourcetype.Title)
	case "Label":
		res = __gong__marshallString(ident, "Label", resourcetype.Label)

	default:
		log.Panicf("Unknown field %s for Gongstruct ResourceType", fieldName)
	}
	return
}

func (simple *Simple) GongMarshallField(stage *Stage, fieldName string) (res string) {
	ident := simple.GongGetIdentifier(stage)
	_ = ident
	switch fieldName {
	case "Name":
		res = __gong__marshallString(ident, "Name", simple.Name)
	case "Type_XSD":
		res = __gong__marshallString(ident, "Type_XSD", simple.Type_XSD)
	case "Href":
		res = __gong__marshallString(ident, "Href", simple.Href)
	case "Role":
		res = __gong__marshallString(ident, "Role", simple.Role)
	case "Arcrole":
		res = __gong__marshallString(ident, "Arcrole", simple.Arcrole)
	case "Title":
		res = __gong__marshallString(ident, "Title", simple.Title)
	case "Show":
		res = __gong__marshallString(ident, "Show", simple.Show)
	case "Actuate":
		res = __gong__marshallString(ident, "Actuate", simple.Actuate)

	default:
		log.Panicf("Unknown field %s for Gongstruct Simple", fieldName)
	}
	return
}

func (titleelttype *TitleEltType) GongMarshallField(stage *Stage, fieldName string) (res string) {
	ident := titleelttype.GongGetIdentifier(stage)
	_ = ident
	switch fieldName {
	case "Name":
		res = __gong__marshallString(ident, "Name", titleelttype.Name)
	case "Type_XSD":
		res = __gong__marshallString(ident, "Type_XSD", titleelttype.Type_XSD)
	case "Lang":
		res = __gong__marshallString(ident, "Lang", titleelttype.Lang)

	default:
		log.Panicf("Unknown field %s for Gongstruct TitleEltType", fieldName)
	}
	return
}

// insertion point for marshall all fields methods
func (arctype *ArcType) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(arctype.GongMarshallField(stage, "Name"))
		initializerStatements.WriteString(arctype.GongMarshallField(stage, "Type_XSD"))
		initializerStatements.WriteString(arctype.GongMarshallField(stage, "Arcrole"))
		initializerStatements.WriteString(arctype.GongMarshallField(stage, "Title"))
		initializerStatements.WriteString(arctype.GongMarshallField(stage, "Show"))
		initializerStatements.WriteString(arctype.GongMarshallField(stage, "Actuate"))
		initializerStatements.WriteString(arctype.GongMarshallField(stage, "From"))
		initializerStatements.WriteString(arctype.GongMarshallField(stage, "To"))
		pointersInitializesStatements.WriteString(arctype.GongMarshallField(stage, "Title_1"))
	}
	initRes = initializerStatements.String()
	ptrRes = pointersInitializesStatements.String()
	return
}
func (extended *Extended) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(extended.GongMarshallField(stage, "Name"))
		initializerStatements.WriteString(extended.GongMarshallField(stage, "Type_XSD"))
		initializerStatements.WriteString(extended.GongMarshallField(stage, "Role"))
		initializerStatements.WriteString(extended.GongMarshallField(stage, "Title"))
		pointersInitializesStatements.WriteString(extended.GongMarshallField(stage, "Title_1"))
		pointersInitializesStatements.WriteString(extended.GongMarshallField(stage, "Resource"))
		pointersInitializesStatements.WriteString(extended.GongMarshallField(stage, "Locator"))
		pointersInitializesStatements.WriteString(extended.GongMarshallField(stage, "Arc"))
	}
	initRes = initializerStatements.String()
	ptrRes = pointersInitializesStatements.String()
	return
}
func (locatortype *LocatorType) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(locatortype.GongMarshallField(stage, "Name"))
		initializerStatements.WriteString(locatortype.GongMarshallField(stage, "Type_XSD"))
		initializerStatements.WriteString(locatortype.GongMarshallField(stage, "Href"))
		initializerStatements.WriteString(locatortype.GongMarshallField(stage, "Role"))
		initializerStatements.WriteString(locatortype.GongMarshallField(stage, "Title"))
		initializerStatements.WriteString(locatortype.GongMarshallField(stage, "Label"))
		pointersInitializesStatements.WriteString(locatortype.GongMarshallField(stage, "Title_1"))
	}
	initRes = initializerStatements.String()
	ptrRes = pointersInitializesStatements.String()
	return
}
func (resourcetype *ResourceType) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(resourcetype.GongMarshallField(stage, "Name"))
		initializerStatements.WriteString(resourcetype.GongMarshallField(stage, "Type_XSD"))
		initializerStatements.WriteString(resourcetype.GongMarshallField(stage, "Role"))
		initializerStatements.WriteString(resourcetype.GongMarshallField(stage, "Title"))
		initializerStatements.WriteString(resourcetype.GongMarshallField(stage, "Label"))
	}
	initRes = initializerStatements.String()
	ptrRes = pointersInitializesStatements.String()
	return
}
func (simple *Simple) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(simple.GongMarshallField(stage, "Name"))
		initializerStatements.WriteString(simple.GongMarshallField(stage, "Type_XSD"))
		initializerStatements.WriteString(simple.GongMarshallField(stage, "Href"))
		initializerStatements.WriteString(simple.GongMarshallField(stage, "Role"))
		initializerStatements.WriteString(simple.GongMarshallField(stage, "Arcrole"))
		initializerStatements.WriteString(simple.GongMarshallField(stage, "Title"))
		initializerStatements.WriteString(simple.GongMarshallField(stage, "Show"))
		initializerStatements.WriteString(simple.GongMarshallField(stage, "Actuate"))
	}
	initRes = initializerStatements.String()
	ptrRes = pointersInitializesStatements.String()
	return
}
func (titleelttype *TitleEltType) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(titleelttype.GongMarshallField(stage, "Name"))
		initializerStatements.WriteString(titleelttype.GongMarshallField(stage, "Type_XSD"))
		initializerStatements.WriteString(titleelttype.GongMarshallField(stage, "Lang"))
	}
	initRes = initializerStatements.String()
	ptrRes = pointersInitializesStatements.String()
	return
}
