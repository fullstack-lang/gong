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
				content = content[:firstBrace] + "\n}"
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
	diagramprocessOrdered := []*DiagramProcess{}
	for diagramprocess := range stage.DiagramProcesss {
		diagramprocessOrdered = append(diagramprocessOrdered, diagramprocess)
	}
	sort.Slice(diagramprocessOrdered[:], func(i, j int) bool {
		diagramprocessi := diagramprocessOrdered[i]
		diagramprocessj := diagramprocessOrdered[j]
		diagramprocessi_order, oki := stage.DiagramProcess_stagedOrder[diagramprocessi]
		diagramprocessj_order, okj := stage.DiagramProcess_stagedOrder[diagramprocessj]
		if !oki || !okj {
			log.Fatalln("unknown pointers")
		}
		return diagramprocessi_order < diagramprocessj_order
	})
	if len(diagramprocessOrdered) > 0 {
		identifiersDecl.WriteString("\n")
	}
	for _, diagramprocess := range diagramprocessOrdered {

		identifiersDecl.WriteString(diagramprocess.GongMarshallIdentifier(stage))

		initializerStatements.WriteString("\n")
		// Insertion point for basic fields value assignment
		initializerStatements.WriteString(diagramprocess.GongMarshallField(stage, "Name"))
		initializerStatements.WriteString(diagramprocess.GongMarshallField(stage, "ComputedPrefix"))
		initializerStatements.WriteString(diagramprocess.GongMarshallField(stage, "IsInRenameMode"))
		initializerStatements.WriteString(diagramprocess.GongMarshallField(stage, "IsExpanded"))
		initializerStatements.WriteString(diagramprocess.GongMarshallField(stage, "IsChecked"))
		initializerStatements.WriteString(diagramprocess.GongMarshallField(stage, "IsEditable_"))
		initializerStatements.WriteString(diagramprocess.GongMarshallField(stage, "IsShowPrefix"))
		initializerStatements.WriteString(diagramprocess.GongMarshallField(stage, "DefaultBoxWidth"))
		initializerStatements.WriteString(diagramprocess.GongMarshallField(stage, "DefaultBoxHeigth"))
		initializerStatements.WriteString(diagramprocess.GongMarshallField(stage, "Width"))
		initializerStatements.WriteString(diagramprocess.GongMarshallField(stage, "Height"))
		pointersInitializesStatements.WriteString(diagramprocess.GongMarshallField(stage, "Process_Shapes"))
		pointersInitializesStatements.WriteString(diagramprocess.GongMarshallField(stage, "ProcesssWhoseNodeIsExpanded"))
		initializerStatements.WriteString(diagramprocess.GongMarshallField(stage, "IsProcesssNodeExpanded"))
		pointersInitializesStatements.WriteString(diagramprocess.GongMarshallField(stage, "ProcessComposition_Shapes"))
	}

	libraryOrdered := []*Library{}
	for library := range stage.Librarys {
		libraryOrdered = append(libraryOrdered, library)
	}
	sort.Slice(libraryOrdered[:], func(i, j int) bool {
		libraryi := libraryOrdered[i]
		libraryj := libraryOrdered[j]
		libraryi_order, oki := stage.Library_stagedOrder[libraryi]
		libraryj_order, okj := stage.Library_stagedOrder[libraryj]
		if !oki || !okj {
			log.Fatalln("unknown pointers")
		}
		return libraryi_order < libraryj_order
	})
	if len(libraryOrdered) > 0 {
		identifiersDecl.WriteString("\n")
	}
	for _, library := range libraryOrdered {

		identifiersDecl.WriteString(library.GongMarshallIdentifier(stage))

		initializerStatements.WriteString("\n")
		// Insertion point for basic fields value assignment
		initializerStatements.WriteString(library.GongMarshallField(stage, "Name"))
		initializerStatements.WriteString(library.GongMarshallField(stage, "ComputedPrefix"))
		initializerStatements.WriteString(library.GongMarshallField(stage, "IsInRenameMode"))
		initializerStatements.WriteString(library.GongMarshallField(stage, "IsExpanded"))
		pointersInitializesStatements.WriteString(library.GongMarshallField(stage, "Diagrams"))
		pointersInitializesStatements.WriteString(library.GongMarshallField(stage, "SubLibraries"))
		initializerStatements.WriteString(library.GongMarshallField(stage, "NbPixPerCharacter"))
		initializerStatements.WriteString(library.GongMarshallField(stage, "LogoSVGFile"))
	}

	processOrdered := []*Process{}
	for process := range stage.Processs {
		processOrdered = append(processOrdered, process)
	}
	sort.Slice(processOrdered[:], func(i, j int) bool {
		processi := processOrdered[i]
		processj := processOrdered[j]
		processi_order, oki := stage.Process_stagedOrder[processi]
		processj_order, okj := stage.Process_stagedOrder[processj]
		if !oki || !okj {
			log.Fatalln("unknown pointers")
		}
		return processi_order < processj_order
	})
	if len(processOrdered) > 0 {
		identifiersDecl.WriteString("\n")
	}
	for _, process := range processOrdered {

		identifiersDecl.WriteString(process.GongMarshallIdentifier(stage))

		initializerStatements.WriteString("\n")
		// Insertion point for basic fields value assignment
		initializerStatements.WriteString(process.GongMarshallField(stage, "Name"))
		initializerStatements.WriteString(process.GongMarshallField(stage, "ComputedPrefix"))
		initializerStatements.WriteString(process.GongMarshallField(stage, "IsInRenameMode"))
		initializerStatements.WriteString(process.GongMarshallField(stage, "IsExpanded"))
		pointersInitializesStatements.WriteString(process.GongMarshallField(stage, "SubProcesses"))
	}

	processcompositionshapeOrdered := []*ProcessCompositionShape{}
	for processcompositionshape := range stage.ProcessCompositionShapes {
		processcompositionshapeOrdered = append(processcompositionshapeOrdered, processcompositionshape)
	}
	sort.Slice(processcompositionshapeOrdered[:], func(i, j int) bool {
		processcompositionshapei := processcompositionshapeOrdered[i]
		processcompositionshapej := processcompositionshapeOrdered[j]
		processcompositionshapei_order, oki := stage.ProcessCompositionShape_stagedOrder[processcompositionshapei]
		processcompositionshapej_order, okj := stage.ProcessCompositionShape_stagedOrder[processcompositionshapej]
		if !oki || !okj {
			log.Fatalln("unknown pointers")
		}
		return processcompositionshapei_order < processcompositionshapej_order
	})
	if len(processcompositionshapeOrdered) > 0 {
		identifiersDecl.WriteString("\n")
	}
	for _, processcompositionshape := range processcompositionshapeOrdered {

		identifiersDecl.WriteString(processcompositionshape.GongMarshallIdentifier(stage))

		initializerStatements.WriteString("\n")
		// Insertion point for basic fields value assignment
		initializerStatements.WriteString(processcompositionshape.GongMarshallField(stage, "Name"))
		pointersInitializesStatements.WriteString(processcompositionshape.GongMarshallField(stage, "Process"))
		initializerStatements.WriteString(processcompositionshape.GongMarshallField(stage, "StartRatio"))
		initializerStatements.WriteString(processcompositionshape.GongMarshallField(stage, "EndRatio"))
		initializerStatements.WriteString(processcompositionshape.GongMarshallField(stage, "StartOrientation"))
		initializerStatements.WriteString(processcompositionshape.GongMarshallField(stage, "EndOrientation"))
		initializerStatements.WriteString(processcompositionshape.GongMarshallField(stage, "CornerOffsetRatio"))
		initializerStatements.WriteString(processcompositionshape.GongMarshallField(stage, "IsHidden"))
	}

	processshapeOrdered := []*ProcessShape{}
	for processshape := range stage.ProcessShapes {
		processshapeOrdered = append(processshapeOrdered, processshape)
	}
	sort.Slice(processshapeOrdered[:], func(i, j int) bool {
		processshapei := processshapeOrdered[i]
		processshapej := processshapeOrdered[j]
		processshapei_order, oki := stage.ProcessShape_stagedOrder[processshapei]
		processshapej_order, okj := stage.ProcessShape_stagedOrder[processshapej]
		if !oki || !okj {
			log.Fatalln("unknown pointers")
		}
		return processshapei_order < processshapej_order
	})
	if len(processshapeOrdered) > 0 {
		identifiersDecl.WriteString("\n")
	}
	for _, processshape := range processshapeOrdered {

		identifiersDecl.WriteString(processshape.GongMarshallIdentifier(stage))

		initializerStatements.WriteString("\n")
		// Insertion point for basic fields value assignment
		initializerStatements.WriteString(processshape.GongMarshallField(stage, "Name"))
		pointersInitializesStatements.WriteString(processshape.GongMarshallField(stage, "Process"))
		initializerStatements.WriteString(processshape.GongMarshallField(stage, "IsExpanded"))
		initializerStatements.WriteString(processshape.GongMarshallField(stage, "X"))
		initializerStatements.WriteString(processshape.GongMarshallField(stage, "Y"))
		initializerStatements.WriteString(processshape.GongMarshallField(stage, "Width"))
		initializerStatements.WriteString(processshape.GongMarshallField(stage, "Height"))
		initializerStatements.WriteString(processshape.GongMarshallField(stage, "IsHidden"))
	}

	// insertion initialization of objects to stage
	for _, diagramprocess := range diagramprocessOrdered {
		_ = diagramprocess
		var setPointerField string
		_ = setPointerField

		// Insertion point for pointers initialization
	}

	for _, library := range libraryOrdered {
		_ = library
		var setPointerField string
		_ = setPointerField

		// Insertion point for pointers initialization
	}

	for _, process := range processOrdered {
		_ = process
		var setPointerField string
		_ = setPointerField

		// Insertion point for pointers initialization
	}

	for _, processcompositionshape := range processcompositionshapeOrdered {
		_ = processcompositionshape
		var setPointerField string
		_ = setPointerField

		// Insertion point for pointers initialization
	}

	for _, processshape := range processshapeOrdered {
		_ = processshape
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
func (diagramprocess *DiagramProcess) GongMarshallField(stage *Stage, fieldName string) (res string) {

	switch fieldName {
	case "Name":
		res = StringInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", diagramprocess.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "Name")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", ToRawStringLiteral(diagramprocess.Name))
	case "ComputedPrefix":
		res = StringInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", diagramprocess.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "ComputedPrefix")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", ToRawStringLiteral(diagramprocess.ComputedPrefix))
	case "IsInRenameMode":
		res = NumberInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", diagramprocess.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "IsInRenameMode")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%t", diagramprocess.IsInRenameMode))
	case "IsExpanded":
		res = NumberInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", diagramprocess.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "IsExpanded")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%t", diagramprocess.IsExpanded))
	case "IsChecked":
		res = NumberInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", diagramprocess.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "IsChecked")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%t", diagramprocess.IsChecked))
	case "IsEditable_":
		res = NumberInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", diagramprocess.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "IsEditable_")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%t", diagramprocess.IsEditable_))
	case "IsShowPrefix":
		res = NumberInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", diagramprocess.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "IsShowPrefix")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%t", diagramprocess.IsShowPrefix))
	case "DefaultBoxWidth":
		res = NumberInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", diagramprocess.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "DefaultBoxWidth")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%f", diagramprocess.DefaultBoxWidth))
	case "DefaultBoxHeigth":
		res = NumberInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", diagramprocess.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "DefaultBoxHeigth")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%f", diagramprocess.DefaultBoxHeigth))
	case "Width":
		res = NumberInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", diagramprocess.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "Width")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%f", diagramprocess.Width))
	case "Height":
		res = NumberInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", diagramprocess.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "Height")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%f", diagramprocess.Height))
	case "IsProcesssNodeExpanded":
		res = NumberInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", diagramprocess.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "IsProcesssNodeExpanded")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%t", diagramprocess.IsProcesssNodeExpanded))

	case "Process_Shapes":
		var sb strings.Builder
		for _, _processshape := range diagramprocess.Process_Shapes {
			tmp := SliceOfPointersFieldInitStatement
			tmp = strings.ReplaceAll(tmp, "{{Identifier}}", diagramprocess.GongGetIdentifier(stage))
			tmp = strings.ReplaceAll(tmp, "{{GeneratedFieldName}}", "Process_Shapes")
			tmp = strings.ReplaceAll(tmp, "{{GeneratedFieldNameValue}}", _processshape.GongGetIdentifier(stage))
			sb.WriteString(tmp)
		}
		res = sb.String()
	case "ProcesssWhoseNodeIsExpanded":
		var sb strings.Builder
		for _, _process := range diagramprocess.ProcesssWhoseNodeIsExpanded {
			tmp := SliceOfPointersFieldInitStatement
			tmp = strings.ReplaceAll(tmp, "{{Identifier}}", diagramprocess.GongGetIdentifier(stage))
			tmp = strings.ReplaceAll(tmp, "{{GeneratedFieldName}}", "ProcesssWhoseNodeIsExpanded")
			tmp = strings.ReplaceAll(tmp, "{{GeneratedFieldNameValue}}", _process.GongGetIdentifier(stage))
			sb.WriteString(tmp)
		}
		res = sb.String()
	case "ProcessComposition_Shapes":
		var sb strings.Builder
		for _, _processcompositionshape := range diagramprocess.ProcessComposition_Shapes {
			tmp := SliceOfPointersFieldInitStatement
			tmp = strings.ReplaceAll(tmp, "{{Identifier}}", diagramprocess.GongGetIdentifier(stage))
			tmp = strings.ReplaceAll(tmp, "{{GeneratedFieldName}}", "ProcessComposition_Shapes")
			tmp = strings.ReplaceAll(tmp, "{{GeneratedFieldNameValue}}", _processcompositionshape.GongGetIdentifier(stage))
			sb.WriteString(tmp)
		}
		res = sb.String()
	default:
		log.Panicf("Unknown field %s for Gongstruct DiagramProcess", fieldName)
	}
	return
}

func (library *Library) GongMarshallField(stage *Stage, fieldName string) (res string) {

	switch fieldName {
	case "Name":
		res = StringInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", library.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "Name")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", ToRawStringLiteral(library.Name))
	case "ComputedPrefix":
		res = StringInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", library.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "ComputedPrefix")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", ToRawStringLiteral(library.ComputedPrefix))
	case "IsInRenameMode":
		res = NumberInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", library.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "IsInRenameMode")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%t", library.IsInRenameMode))
	case "IsExpanded":
		res = NumberInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", library.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "IsExpanded")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%t", library.IsExpanded))
	case "NbPixPerCharacter":
		res = NumberInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", library.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "NbPixPerCharacter")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%f", library.NbPixPerCharacter))
	case "LogoSVGFile":
		res = StringInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", library.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "LogoSVGFile")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", ToRawStringLiteral(library.LogoSVGFile))

	case "Diagrams":
		var sb strings.Builder
		for _, _diagramprocess := range library.Diagrams {
			tmp := SliceOfPointersFieldInitStatement
			tmp = strings.ReplaceAll(tmp, "{{Identifier}}", library.GongGetIdentifier(stage))
			tmp = strings.ReplaceAll(tmp, "{{GeneratedFieldName}}", "Diagrams")
			tmp = strings.ReplaceAll(tmp, "{{GeneratedFieldNameValue}}", _diagramprocess.GongGetIdentifier(stage))
			sb.WriteString(tmp)
		}
		res = sb.String()
	case "SubLibraries":
		var sb strings.Builder
		for _, _library := range library.SubLibraries {
			tmp := SliceOfPointersFieldInitStatement
			tmp = strings.ReplaceAll(tmp, "{{Identifier}}", library.GongGetIdentifier(stage))
			tmp = strings.ReplaceAll(tmp, "{{GeneratedFieldName}}", "SubLibraries")
			tmp = strings.ReplaceAll(tmp, "{{GeneratedFieldNameValue}}", _library.GongGetIdentifier(stage))
			sb.WriteString(tmp)
		}
		res = sb.String()
	default:
		log.Panicf("Unknown field %s for Gongstruct Library", fieldName)
	}
	return
}

func (process *Process) GongMarshallField(stage *Stage, fieldName string) (res string) {

	switch fieldName {
	case "Name":
		res = StringInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", process.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "Name")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", ToRawStringLiteral(process.Name))
	case "ComputedPrefix":
		res = StringInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", process.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "ComputedPrefix")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", ToRawStringLiteral(process.ComputedPrefix))
	case "IsInRenameMode":
		res = NumberInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", process.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "IsInRenameMode")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%t", process.IsInRenameMode))
	case "IsExpanded":
		res = NumberInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", process.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "IsExpanded")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%t", process.IsExpanded))

	case "SubProcesses":
		var sb strings.Builder
		for _, _process := range process.SubProcesses {
			tmp := SliceOfPointersFieldInitStatement
			tmp = strings.ReplaceAll(tmp, "{{Identifier}}", process.GongGetIdentifier(stage))
			tmp = strings.ReplaceAll(tmp, "{{GeneratedFieldName}}", "SubProcesses")
			tmp = strings.ReplaceAll(tmp, "{{GeneratedFieldNameValue}}", _process.GongGetIdentifier(stage))
			sb.WriteString(tmp)
		}
		res = sb.String()
	default:
		log.Panicf("Unknown field %s for Gongstruct Process", fieldName)
	}
	return
}

func (processcompositionshape *ProcessCompositionShape) GongMarshallField(stage *Stage, fieldName string) (res string) {

	switch fieldName {
	case "Name":
		res = StringInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", processcompositionshape.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "Name")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", ToRawStringLiteral(processcompositionshape.Name))
	case "StartRatio":
		res = NumberInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", processcompositionshape.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "StartRatio")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%f", processcompositionshape.StartRatio))
	case "EndRatio":
		res = NumberInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", processcompositionshape.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "EndRatio")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%f", processcompositionshape.EndRatio))
	case "StartOrientation":
		if processcompositionshape.StartOrientation.ToCodeString() != "" {
			res = StringEnumInitStatement
			res = strings.ReplaceAll(res, "{{Identifier}}", processcompositionshape.GongGetIdentifier(stage))
			res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "StartOrientation")
			res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", "models."+processcompositionshape.StartOrientation.ToCodeString())
		} else {
			// in case of empty enum, we need to unstage the previous value
			res = StringEnumInitStatement
			res = strings.ReplaceAll(res, "{{Identifier}}", processcompositionshape.GongGetIdentifier(stage))
			res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "StartOrientation")
			res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", "\"\"")
		}
	case "EndOrientation":
		if processcompositionshape.EndOrientation.ToCodeString() != "" {
			res = StringEnumInitStatement
			res = strings.ReplaceAll(res, "{{Identifier}}", processcompositionshape.GongGetIdentifier(stage))
			res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "EndOrientation")
			res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", "models."+processcompositionshape.EndOrientation.ToCodeString())
		} else {
			// in case of empty enum, we need to unstage the previous value
			res = StringEnumInitStatement
			res = strings.ReplaceAll(res, "{{Identifier}}", processcompositionshape.GongGetIdentifier(stage))
			res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "EndOrientation")
			res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", "\"\"")
		}
	case "CornerOffsetRatio":
		res = NumberInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", processcompositionshape.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "CornerOffsetRatio")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%f", processcompositionshape.CornerOffsetRatio))
	case "IsHidden":
		res = NumberInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", processcompositionshape.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "IsHidden")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%t", processcompositionshape.IsHidden))

	case "Process":
		if processcompositionshape.Process != nil {
			res = PointerFieldInitStatement
			res = strings.ReplaceAll(res, "{{Identifier}}", processcompositionshape.GongGetIdentifier(stage))
			res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "Process")
			res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", processcompositionshape.Process.GongGetIdentifier(stage))
		} else {
			// in case of nil pointer, we need to unstage the previous value
			res = PointerFieldInitStatement
			res = strings.ReplaceAll(res, "{{Identifier}}", processcompositionshape.GongGetIdentifier(stage))
			res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "Process")
			res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", "nil")
		}
	default:
		log.Panicf("Unknown field %s for Gongstruct ProcessCompositionShape", fieldName)
	}
	return
}

func (processshape *ProcessShape) GongMarshallField(stage *Stage, fieldName string) (res string) {

	switch fieldName {
	case "Name":
		res = StringInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", processshape.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "Name")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", ToRawStringLiteral(processshape.Name))
	case "IsExpanded":
		res = NumberInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", processshape.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "IsExpanded")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%t", processshape.IsExpanded))
	case "X":
		res = NumberInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", processshape.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "X")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%f", processshape.X))
	case "Y":
		res = NumberInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", processshape.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "Y")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%f", processshape.Y))
	case "Width":
		res = NumberInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", processshape.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "Width")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%f", processshape.Width))
	case "Height":
		res = NumberInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", processshape.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "Height")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%f", processshape.Height))
	case "IsHidden":
		res = NumberInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", processshape.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "IsHidden")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%t", processshape.IsHidden))

	case "Process":
		if processshape.Process != nil {
			res = PointerFieldInitStatement
			res = strings.ReplaceAll(res, "{{Identifier}}", processshape.GongGetIdentifier(stage))
			res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "Process")
			res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", processshape.Process.GongGetIdentifier(stage))
		} else {
			// in case of nil pointer, we need to unstage the previous value
			res = PointerFieldInitStatement
			res = strings.ReplaceAll(res, "{{Identifier}}", processshape.GongGetIdentifier(stage))
			res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "Process")
			res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", "nil")
		}
	default:
		log.Panicf("Unknown field %s for Gongstruct ProcessShape", fieldName)
	}
	return
}

// insertion point for marshall all fields methods
func (diagramprocess *DiagramProcess) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(diagramprocess.GongMarshallField(stage, "Name"))
		initializerStatements.WriteString(diagramprocess.GongMarshallField(stage, "ComputedPrefix"))
		initializerStatements.WriteString(diagramprocess.GongMarshallField(stage, "IsInRenameMode"))
		initializerStatements.WriteString(diagramprocess.GongMarshallField(stage, "IsExpanded"))
		initializerStatements.WriteString(diagramprocess.GongMarshallField(stage, "IsChecked"))
		initializerStatements.WriteString(diagramprocess.GongMarshallField(stage, "IsEditable_"))
		initializerStatements.WriteString(diagramprocess.GongMarshallField(stage, "IsShowPrefix"))
		initializerStatements.WriteString(diagramprocess.GongMarshallField(stage, "DefaultBoxWidth"))
		initializerStatements.WriteString(diagramprocess.GongMarshallField(stage, "DefaultBoxHeigth"))
		initializerStatements.WriteString(diagramprocess.GongMarshallField(stage, "Width"))
		initializerStatements.WriteString(diagramprocess.GongMarshallField(stage, "Height"))
		pointersInitializesStatements.WriteString(diagramprocess.GongMarshallField(stage, "Process_Shapes"))
		pointersInitializesStatements.WriteString(diagramprocess.GongMarshallField(stage, "ProcesssWhoseNodeIsExpanded"))
		initializerStatements.WriteString(diagramprocess.GongMarshallField(stage, "IsProcesssNodeExpanded"))
		pointersInitializesStatements.WriteString(diagramprocess.GongMarshallField(stage, "ProcessComposition_Shapes"))
	}
	initRes = initializerStatements.String()
	ptrRes = pointersInitializesStatements.String()
	return
}
func (library *Library) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(library.GongMarshallField(stage, "Name"))
		initializerStatements.WriteString(library.GongMarshallField(stage, "ComputedPrefix"))
		initializerStatements.WriteString(library.GongMarshallField(stage, "IsInRenameMode"))
		initializerStatements.WriteString(library.GongMarshallField(stage, "IsExpanded"))
		pointersInitializesStatements.WriteString(library.GongMarshallField(stage, "Diagrams"))
		pointersInitializesStatements.WriteString(library.GongMarshallField(stage, "SubLibraries"))
		initializerStatements.WriteString(library.GongMarshallField(stage, "NbPixPerCharacter"))
		initializerStatements.WriteString(library.GongMarshallField(stage, "LogoSVGFile"))
	}
	initRes = initializerStatements.String()
	ptrRes = pointersInitializesStatements.String()
	return
}
func (process *Process) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(process.GongMarshallField(stage, "Name"))
		initializerStatements.WriteString(process.GongMarshallField(stage, "ComputedPrefix"))
		initializerStatements.WriteString(process.GongMarshallField(stage, "IsInRenameMode"))
		initializerStatements.WriteString(process.GongMarshallField(stage, "IsExpanded"))
		pointersInitializesStatements.WriteString(process.GongMarshallField(stage, "SubProcesses"))
	}
	initRes = initializerStatements.String()
	ptrRes = pointersInitializesStatements.String()
	return
}
func (processcompositionshape *ProcessCompositionShape) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(processcompositionshape.GongMarshallField(stage, "Name"))
		pointersInitializesStatements.WriteString(processcompositionshape.GongMarshallField(stage, "Process"))
		initializerStatements.WriteString(processcompositionshape.GongMarshallField(stage, "StartRatio"))
		initializerStatements.WriteString(processcompositionshape.GongMarshallField(stage, "EndRatio"))
		initializerStatements.WriteString(processcompositionshape.GongMarshallField(stage, "StartOrientation"))
		initializerStatements.WriteString(processcompositionshape.GongMarshallField(stage, "EndOrientation"))
		initializerStatements.WriteString(processcompositionshape.GongMarshallField(stage, "CornerOffsetRatio"))
		initializerStatements.WriteString(processcompositionshape.GongMarshallField(stage, "IsHidden"))
	}
	initRes = initializerStatements.String()
	ptrRes = pointersInitializesStatements.String()
	return
}
func (processshape *ProcessShape) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(processshape.GongMarshallField(stage, "Name"))
		pointersInitializesStatements.WriteString(processshape.GongMarshallField(stage, "Process"))
		initializerStatements.WriteString(processshape.GongMarshallField(stage, "IsExpanded"))
		initializerStatements.WriteString(processshape.GongMarshallField(stage, "X"))
		initializerStatements.WriteString(processshape.GongMarshallField(stage, "Y"))
		initializerStatements.WriteString(processshape.GongMarshallField(stage, "Width"))
		initializerStatements.WriteString(processshape.GongMarshallField(stage, "Height"))
		initializerStatements.WriteString(processshape.GongMarshallField(stage, "IsHidden"))
	}
	initRes = initializerStatements.String()
	ptrRes = pointersInitializesStatements.String()
	return
}
