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
	allocatedresourceshapeOrdered := []*AllocatedResourceShape{}
	for allocatedresourceshape := range stage.AllocatedResourceShapes {
		allocatedresourceshapeOrdered = append(allocatedresourceshapeOrdered, allocatedresourceshape)
	}
	sort.Slice(allocatedresourceshapeOrdered[:], func(i, j int) bool {
		allocatedresourceshapei := allocatedresourceshapeOrdered[i]
		allocatedresourceshapej := allocatedresourceshapeOrdered[j]
		allocatedresourceshapei_order, oki := stage.AllocatedResourceShape_stagedOrder[allocatedresourceshapei]
		allocatedresourceshapej_order, okj := stage.AllocatedResourceShape_stagedOrder[allocatedresourceshapej]
		if !oki || !okj {
			log.Fatalln("unknown pointers")
		}
		return allocatedresourceshapei_order < allocatedresourceshapej_order
	})
	if len(allocatedresourceshapeOrdered) > 0 {
		identifiersDecl.WriteString("\n")
	}
	for _, allocatedresourceshape := range allocatedresourceshapeOrdered {

		identifiersDecl.WriteString(allocatedresourceshape.GongMarshallIdentifier(stage))

		initializerStatements.WriteString("\n")
		// Insertion point for basic fields value assignment
		initializerStatements.WriteString(allocatedresourceshape.GongMarshallField(stage, "Name"))
		pointersInitializesStatements.WriteString(allocatedresourceshape.GongMarshallField(stage, "Part"))
		pointersInitializesStatements.WriteString(allocatedresourceshape.GongMarshallField(stage, "Resource"))
	}

	allocatedsystemshapeOrdered := []*AllocatedSystemShape{}
	for allocatedsystemshape := range stage.AllocatedSystemShapes {
		allocatedsystemshapeOrdered = append(allocatedsystemshapeOrdered, allocatedsystemshape)
	}
	sort.Slice(allocatedsystemshapeOrdered[:], func(i, j int) bool {
		allocatedsystemshapei := allocatedsystemshapeOrdered[i]
		allocatedsystemshapej := allocatedsystemshapeOrdered[j]
		allocatedsystemshapei_order, oki := stage.AllocatedSystemShape_stagedOrder[allocatedsystemshapei]
		allocatedsystemshapej_order, okj := stage.AllocatedSystemShape_stagedOrder[allocatedsystemshapej]
		if !oki || !okj {
			log.Fatalln("unknown pointers")
		}
		return allocatedsystemshapei_order < allocatedsystemshapej_order
	})
	if len(allocatedsystemshapeOrdered) > 0 {
		identifiersDecl.WriteString("\n")
	}
	for _, allocatedsystemshape := range allocatedsystemshapeOrdered {

		identifiersDecl.WriteString(allocatedsystemshape.GongMarshallIdentifier(stage))

		initializerStatements.WriteString("\n")
		// Insertion point for basic fields value assignment
		initializerStatements.WriteString(allocatedsystemshape.GongMarshallField(stage, "Name"))
		pointersInitializesStatements.WriteString(allocatedsystemshape.GongMarshallField(stage, "Part"))
		pointersInitializesStatements.WriteString(allocatedsystemshape.GongMarshallField(stage, "System"))
	}

	controlflowOrdered := []*ControlFlow{}
	for controlflow := range stage.ControlFlows {
		controlflowOrdered = append(controlflowOrdered, controlflow)
	}
	sort.Slice(controlflowOrdered[:], func(i, j int) bool {
		controlflowi := controlflowOrdered[i]
		controlflowj := controlflowOrdered[j]
		controlflowi_order, oki := stage.ControlFlow_stagedOrder[controlflowi]
		controlflowj_order, okj := stage.ControlFlow_stagedOrder[controlflowj]
		if !oki || !okj {
			log.Fatalln("unknown pointers")
		}
		return controlflowi_order < controlflowj_order
	})
	if len(controlflowOrdered) > 0 {
		identifiersDecl.WriteString("\n")
	}
	for _, controlflow := range controlflowOrdered {

		identifiersDecl.WriteString(controlflow.GongMarshallIdentifier(stage))

		initializerStatements.WriteString("\n")
		// Insertion point for basic fields value assignment
		initializerStatements.WriteString(controlflow.GongMarshallField(stage, "Name"))
		initializerStatements.WriteString(controlflow.GongMarshallField(stage, "Description"))
		initializerStatements.WriteString(controlflow.GongMarshallField(stage, "ComputedPrefix"))
		initializerStatements.WriteString(controlflow.GongMarshallField(stage, "IsExpanded"))
		initializerStatements.WriteString(controlflow.GongMarshallField(stage, "LayoutDirection"))
		pointersInitializesStatements.WriteString(controlflow.GongMarshallField(stage, "Start"))
		pointersInitializesStatements.WriteString(controlflow.GongMarshallField(stage, "End"))
	}

	controlflowshapeOrdered := []*ControlFlowShape{}
	for controlflowshape := range stage.ControlFlowShapes {
		controlflowshapeOrdered = append(controlflowshapeOrdered, controlflowshape)
	}
	sort.Slice(controlflowshapeOrdered[:], func(i, j int) bool {
		controlflowshapei := controlflowshapeOrdered[i]
		controlflowshapej := controlflowshapeOrdered[j]
		controlflowshapei_order, oki := stage.ControlFlowShape_stagedOrder[controlflowshapei]
		controlflowshapej_order, okj := stage.ControlFlowShape_stagedOrder[controlflowshapej]
		if !oki || !okj {
			log.Fatalln("unknown pointers")
		}
		return controlflowshapei_order < controlflowshapej_order
	})
	if len(controlflowshapeOrdered) > 0 {
		identifiersDecl.WriteString("\n")
	}
	for _, controlflowshape := range controlflowshapeOrdered {

		identifiersDecl.WriteString(controlflowshape.GongMarshallIdentifier(stage))

		initializerStatements.WriteString("\n")
		// Insertion point for basic fields value assignment
		initializerStatements.WriteString(controlflowshape.GongMarshallField(stage, "Name"))
		pointersInitializesStatements.WriteString(controlflowshape.GongMarshallField(stage, "ControlFlow"))
		initializerStatements.WriteString(controlflowshape.GongMarshallField(stage, "StartRatio"))
		initializerStatements.WriteString(controlflowshape.GongMarshallField(stage, "EndRatio"))
		initializerStatements.WriteString(controlflowshape.GongMarshallField(stage, "StartOrientation"))
		initializerStatements.WriteString(controlflowshape.GongMarshallField(stage, "EndOrientation"))
		initializerStatements.WriteString(controlflowshape.GongMarshallField(stage, "CornerOffsetRatio"))
		initializerStatements.WriteString(controlflowshape.GongMarshallField(stage, "IsHidden"))
	}

	dataOrdered := []*Data{}
	for data := range stage.Datas {
		dataOrdered = append(dataOrdered, data)
	}
	sort.Slice(dataOrdered[:], func(i, j int) bool {
		datai := dataOrdered[i]
		dataj := dataOrdered[j]
		datai_order, oki := stage.Data_stagedOrder[datai]
		dataj_order, okj := stage.Data_stagedOrder[dataj]
		if !oki || !okj {
			log.Fatalln("unknown pointers")
		}
		return datai_order < dataj_order
	})
	if len(dataOrdered) > 0 {
		identifiersDecl.WriteString("\n")
	}
	for _, data := range dataOrdered {

		identifiersDecl.WriteString(data.GongMarshallIdentifier(stage))

		initializerStatements.WriteString("\n")
		// Insertion point for basic fields value assignment
		initializerStatements.WriteString(data.GongMarshallField(stage, "Name"))
		initializerStatements.WriteString(data.GongMarshallField(stage, "Acronym"))
		initializerStatements.WriteString(data.GongMarshallField(stage, "Description"))
		initializerStatements.WriteString(data.GongMarshallField(stage, "ComputedPrefix"))
		initializerStatements.WriteString(data.GongMarshallField(stage, "IsExpanded"))
		initializerStatements.WriteString(data.GongMarshallField(stage, "LayoutDirection"))
		initializerStatements.WriteString(data.GongMarshallField(stage, "SVG_Path"))
		initializerStatements.WriteString(data.GongMarshallField(stage, "InverseAppliedScaling"))
	}

	dataflowOrdered := []*DataFlow{}
	for dataflow := range stage.DataFlows {
		dataflowOrdered = append(dataflowOrdered, dataflow)
	}
	sort.Slice(dataflowOrdered[:], func(i, j int) bool {
		dataflowi := dataflowOrdered[i]
		dataflowj := dataflowOrdered[j]
		dataflowi_order, oki := stage.DataFlow_stagedOrder[dataflowi]
		dataflowj_order, okj := stage.DataFlow_stagedOrder[dataflowj]
		if !oki || !okj {
			log.Fatalln("unknown pointers")
		}
		return dataflowi_order < dataflowj_order
	})
	if len(dataflowOrdered) > 0 {
		identifiersDecl.WriteString("\n")
	}
	for _, dataflow := range dataflowOrdered {

		identifiersDecl.WriteString(dataflow.GongMarshallIdentifier(stage))

		initializerStatements.WriteString("\n")
		// Insertion point for basic fields value assignment
		initializerStatements.WriteString(dataflow.GongMarshallField(stage, "Name"))
		pointersInitializesStatements.WriteString(dataflow.GongMarshallField(stage, "Datas"))
		initializerStatements.WriteString(dataflow.GongMarshallField(stage, "Description"))
		initializerStatements.WriteString(dataflow.GongMarshallField(stage, "ComputedPrefix"))
		initializerStatements.WriteString(dataflow.GongMarshallField(stage, "IsExpanded"))
		initializerStatements.WriteString(dataflow.GongMarshallField(stage, "LayoutDirection"))
		initializerStatements.WriteString(dataflow.GongMarshallField(stage, "Type"))
		pointersInitializesStatements.WriteString(dataflow.GongMarshallField(stage, "StartPort"))
		pointersInitializesStatements.WriteString(dataflow.GongMarshallField(stage, "EndPort"))
		pointersInitializesStatements.WriteString(dataflow.GongMarshallField(stage, "StartExternalPart"))
		pointersInitializesStatements.WriteString(dataflow.GongMarshallField(stage, "EndExternalPart"))
		initializerStatements.WriteString(dataflow.GongMarshallField(stage, "IsDatasNodeExpanded"))
	}

	dataflowshapeOrdered := []*DataFlowShape{}
	for dataflowshape := range stage.DataFlowShapes {
		dataflowshapeOrdered = append(dataflowshapeOrdered, dataflowshape)
	}
	sort.Slice(dataflowshapeOrdered[:], func(i, j int) bool {
		dataflowshapei := dataflowshapeOrdered[i]
		dataflowshapej := dataflowshapeOrdered[j]
		dataflowshapei_order, oki := stage.DataFlowShape_stagedOrder[dataflowshapei]
		dataflowshapej_order, okj := stage.DataFlowShape_stagedOrder[dataflowshapej]
		if !oki || !okj {
			log.Fatalln("unknown pointers")
		}
		return dataflowshapei_order < dataflowshapej_order
	})
	if len(dataflowshapeOrdered) > 0 {
		identifiersDecl.WriteString("\n")
	}
	for _, dataflowshape := range dataflowshapeOrdered {

		identifiersDecl.WriteString(dataflowshape.GongMarshallIdentifier(stage))

		initializerStatements.WriteString("\n")
		// Insertion point for basic fields value assignment
		initializerStatements.WriteString(dataflowshape.GongMarshallField(stage, "Name"))
		pointersInitializesStatements.WriteString(dataflowshape.GongMarshallField(stage, "DataFlow"))
		initializerStatements.WriteString(dataflowshape.GongMarshallField(stage, "StartRatio"))
		initializerStatements.WriteString(dataflowshape.GongMarshallField(stage, "EndRatio"))
		initializerStatements.WriteString(dataflowshape.GongMarshallField(stage, "StartOrientation"))
		initializerStatements.WriteString(dataflowshape.GongMarshallField(stage, "EndOrientation"))
		initializerStatements.WriteString(dataflowshape.GongMarshallField(stage, "CornerOffsetRatio"))
		initializerStatements.WriteString(dataflowshape.GongMarshallField(stage, "IsHidden"))
	}

	datashapeOrdered := []*DataShape{}
	for datashape := range stage.DataShapes {
		datashapeOrdered = append(datashapeOrdered, datashape)
	}
	sort.Slice(datashapeOrdered[:], func(i, j int) bool {
		datashapei := datashapeOrdered[i]
		datashapej := datashapeOrdered[j]
		datashapei_order, oki := stage.DataShape_stagedOrder[datashapei]
		datashapej_order, okj := stage.DataShape_stagedOrder[datashapej]
		if !oki || !okj {
			log.Fatalln("unknown pointers")
		}
		return datashapei_order < datashapej_order
	})
	if len(datashapeOrdered) > 0 {
		identifiersDecl.WriteString("\n")
	}
	for _, datashape := range datashapeOrdered {

		identifiersDecl.WriteString(datashape.GongMarshallIdentifier(stage))

		initializerStatements.WriteString("\n")
		// Insertion point for basic fields value assignment
		initializerStatements.WriteString(datashape.GongMarshallField(stage, "Name"))
		pointersInitializesStatements.WriteString(datashape.GongMarshallField(stage, "Data"))
		pointersInitializesStatements.WriteString(datashape.GongMarshallField(stage, "DataFlow"))
	}

	diagramstructureOrdered := []*DiagramStructure{}
	for diagramstructure := range stage.DiagramStructures {
		diagramstructureOrdered = append(diagramstructureOrdered, diagramstructure)
	}
	sort.Slice(diagramstructureOrdered[:], func(i, j int) bool {
		diagramstructurei := diagramstructureOrdered[i]
		diagramstructurej := diagramstructureOrdered[j]
		diagramstructurei_order, oki := stage.DiagramStructure_stagedOrder[diagramstructurei]
		diagramstructurej_order, okj := stage.DiagramStructure_stagedOrder[diagramstructurej]
		if !oki || !okj {
			log.Fatalln("unknown pointers")
		}
		return diagramstructurei_order < diagramstructurej_order
	})
	if len(diagramstructureOrdered) > 0 {
		identifiersDecl.WriteString("\n")
	}
	for _, diagramstructure := range diagramstructureOrdered {

		identifiersDecl.WriteString(diagramstructure.GongMarshallIdentifier(stage))

		initializerStatements.WriteString("\n")
		// Insertion point for basic fields value assignment
		initializerStatements.WriteString(diagramstructure.GongMarshallField(stage, "Name"))
		initializerStatements.WriteString(diagramstructure.GongMarshallField(stage, "Description"))
		initializerStatements.WriteString(diagramstructure.GongMarshallField(stage, "ComputedPrefix"))
		initializerStatements.WriteString(diagramstructure.GongMarshallField(stage, "IsExpanded"))
		initializerStatements.WriteString(diagramstructure.GongMarshallField(stage, "LayoutDirection"))
		initializerStatements.WriteString(diagramstructure.GongMarshallField(stage, "IsChecked"))
		initializerStatements.WriteString(diagramstructure.GongMarshallField(stage, "IsEditable_"))
		initializerStatements.WriteString(diagramstructure.GongMarshallField(stage, "IsShowPrefix"))
		initializerStatements.WriteString(diagramstructure.GongMarshallField(stage, "DefaultBoxWidth"))
		initializerStatements.WriteString(diagramstructure.GongMarshallField(stage, "DefaultBoxHeigth"))
		initializerStatements.WriteString(diagramstructure.GongMarshallField(stage, "Width"))
		initializerStatements.WriteString(diagramstructure.GongMarshallField(stage, "Height"))
		pointersInitializesStatements.WriteString(diagramstructure.GongMarshallField(stage, "System_Shapes"))
		initializerStatements.WriteString(diagramstructure.GongMarshallField(stage, "IsSystemsNodeExpanded"))
		pointersInitializesStatements.WriteString(diagramstructure.GongMarshallField(stage, "SystemsWhoseNodeIsExpanded"))
		pointersInitializesStatements.WriteString(diagramstructure.GongMarshallField(stage, "Part_Shapes"))
		initializerStatements.WriteString(diagramstructure.GongMarshallField(stage, "IsPartsNodeExpanded"))
		pointersInitializesStatements.WriteString(diagramstructure.GongMarshallField(stage, "PartWhoseNodeIsExpanded"))
		pointersInitializesStatements.WriteString(diagramstructure.GongMarshallField(stage, "ExternalPart_Shapes"))
		initializerStatements.WriteString(diagramstructure.GongMarshallField(stage, "IsExternalPartsNodeExpanded"))
		pointersInitializesStatements.WriteString(diagramstructure.GongMarshallField(stage, "ExternalPartWhoseNodeIsExpanded"))
		pointersInitializesStatements.WriteString(diagramstructure.GongMarshallField(stage, "ExternalPartsWhoseOutDataFlowsNodeIsExpanded"))
		pointersInitializesStatements.WriteString(diagramstructure.GongMarshallField(stage, "ExternalPartsWhoseInDataFlowsNodeIsExpanded"))
		pointersInitializesStatements.WriteString(diagramstructure.GongMarshallField(stage, "PortsWhoseNodeIsExpanded"))
		pointersInitializesStatements.WriteString(diagramstructure.GongMarshallField(stage, "Port_Shapes"))
		pointersInitializesStatements.WriteString(diagramstructure.GongMarshallField(stage, "ControlFlowsWhoseNodeIsExpanded"))
		pointersInitializesStatements.WriteString(diagramstructure.GongMarshallField(stage, "ControlFlow_Shapes"))
		pointersInitializesStatements.WriteString(diagramstructure.GongMarshallField(stage, "DataFlowsWhoseNodeIsExpanded"))
		pointersInitializesStatements.WriteString(diagramstructure.GongMarshallField(stage, "DataFlow_Shapes"))
		pointersInitializesStatements.WriteString(diagramstructure.GongMarshallField(stage, "DatasWhoseNodeIsExpanded"))
		pointersInitializesStatements.WriteString(diagramstructure.GongMarshallField(stage, "Data_Shapes"))
		pointersInitializesStatements.WriteString(diagramstructure.GongMarshallField(stage, "DataFlowsWhoseDataNodeIsExpanded"))
		pointersInitializesStatements.WriteString(diagramstructure.GongMarshallField(stage, "AllocatedResourcesWhoseNodeIsExpanded"))
		pointersInitializesStatements.WriteString(diagramstructure.GongMarshallField(stage, "AllocatedResourceShapes"))
		pointersInitializesStatements.WriteString(diagramstructure.GongMarshallField(stage, "AllocatedSystemesWhoseNodeIsExpanded"))
		pointersInitializesStatements.WriteString(diagramstructure.GongMarshallField(stage, "AllocatedSystemShapes"))
		pointersInitializesStatements.WriteString(diagramstructure.GongMarshallField(stage, "Note_Shapes"))
		pointersInitializesStatements.WriteString(diagramstructure.GongMarshallField(stage, "NotesWhoseNodeIsExpanded"))
		initializerStatements.WriteString(diagramstructure.GongMarshallField(stage, "IsNotesNodeExpanded"))
		pointersInitializesStatements.WriteString(diagramstructure.GongMarshallField(stage, "NotePortShapes"))
	}

	externalpartshapeOrdered := []*ExternalPartShape{}
	for externalpartshape := range stage.ExternalPartShapes {
		externalpartshapeOrdered = append(externalpartshapeOrdered, externalpartshape)
	}
	sort.Slice(externalpartshapeOrdered[:], func(i, j int) bool {
		externalpartshapei := externalpartshapeOrdered[i]
		externalpartshapej := externalpartshapeOrdered[j]
		externalpartshapei_order, oki := stage.ExternalPartShape_stagedOrder[externalpartshapei]
		externalpartshapej_order, okj := stage.ExternalPartShape_stagedOrder[externalpartshapej]
		if !oki || !okj {
			log.Fatalln("unknown pointers")
		}
		return externalpartshapei_order < externalpartshapej_order
	})
	if len(externalpartshapeOrdered) > 0 {
		identifiersDecl.WriteString("\n")
	}
	for _, externalpartshape := range externalpartshapeOrdered {

		identifiersDecl.WriteString(externalpartshape.GongMarshallIdentifier(stage))

		initializerStatements.WriteString("\n")
		// Insertion point for basic fields value assignment
		initializerStatements.WriteString(externalpartshape.GongMarshallField(stage, "Name"))
		pointersInitializesStatements.WriteString(externalpartshape.GongMarshallField(stage, "Part"))
		initializerStatements.WriteString(externalpartshape.GongMarshallField(stage, "IsExpanded"))
		initializerStatements.WriteString(externalpartshape.GongMarshallField(stage, "X"))
		initializerStatements.WriteString(externalpartshape.GongMarshallField(stage, "Y"))
		initializerStatements.WriteString(externalpartshape.GongMarshallField(stage, "Width"))
		initializerStatements.WriteString(externalpartshape.GongMarshallField(stage, "Height"))
		initializerStatements.WriteString(externalpartshape.GongMarshallField(stage, "IsHidden"))
		initializerStatements.WriteString(externalpartshape.GongMarshallField(stage, "TailHeigth"))
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
		initializerStatements.WriteString(library.GongMarshallField(stage, "Description"))
		initializerStatements.WriteString(library.GongMarshallField(stage, "ComputedPrefix"))
		initializerStatements.WriteString(library.GongMarshallField(stage, "IsExpanded"))
		initializerStatements.WriteString(library.GongMarshallField(stage, "LayoutDirection"))
		initializerStatements.WriteString(library.GongMarshallField(stage, "IsRootLibrary"))
		pointersInitializesStatements.WriteString(library.GongMarshallField(stage, "SubLibraries"))
		initializerStatements.WriteString(library.GongMarshallField(stage, "IsSubLibrariesNodeExpanded"))
		pointersInitializesStatements.WriteString(library.GongMarshallField(stage, "SubLibrariesWhoseNodeIsExpanded"))
		initializerStatements.WriteString(library.GongMarshallField(stage, "NbPixPerCharacter"))
		initializerStatements.WriteString(library.GongMarshallField(stage, "LogoSVGFile"))
		pointersInitializesStatements.WriteString(library.GongMarshallField(stage, "RootSystemes"))
		initializerStatements.WriteString(library.GongMarshallField(stage, "IsSystemesNodeExpanded"))
		pointersInitializesStatements.WriteString(library.GongMarshallField(stage, "SystemsWhoseNodeIsExpanded"))
		pointersInitializesStatements.WriteString(library.GongMarshallField(stage, "RootDataFlows"))
		initializerStatements.WriteString(library.GongMarshallField(stage, "IsDataFlowsNodeExpanded"))
		pointersInitializesStatements.WriteString(library.GongMarshallField(stage, "DataFlowsWhoseNodeIsExpanded"))
		pointersInitializesStatements.WriteString(library.GongMarshallField(stage, "RootDatas"))
		initializerStatements.WriteString(library.GongMarshallField(stage, "IsDatasNodeExpanded"))
		pointersInitializesStatements.WriteString(library.GongMarshallField(stage, "DatasWhoseNodeIsExpanded"))
		pointersInitializesStatements.WriteString(library.GongMarshallField(stage, "RootResources"))
		initializerStatements.WriteString(library.GongMarshallField(stage, "IsResourcesNodeExpanded"))
		pointersInitializesStatements.WriteString(library.GongMarshallField(stage, "ResourcesWhoseNodeIsExpanded"))
		pointersInitializesStatements.WriteString(library.GongMarshallField(stage, "PartsWhoseNodeIsExpanded"))
		pointersInitializesStatements.WriteString(library.GongMarshallField(stage, "RootNotes"))
		initializerStatements.WriteString(library.GongMarshallField(stage, "IsNotesNodeExpanded"))
		pointersInitializesStatements.WriteString(library.GongMarshallField(stage, "NotesWhoseNodeIsExpanded"))
		initializerStatements.WriteString(library.GongMarshallField(stage, "IsExpandedTmp"))
	}

	noteOrdered := []*Note{}
	for note := range stage.Notes {
		noteOrdered = append(noteOrdered, note)
	}
	sort.Slice(noteOrdered[:], func(i, j int) bool {
		notei := noteOrdered[i]
		notej := noteOrdered[j]
		notei_order, oki := stage.Note_stagedOrder[notei]
		notej_order, okj := stage.Note_stagedOrder[notej]
		if !oki || !okj {
			log.Fatalln("unknown pointers")
		}
		return notei_order < notej_order
	})
	if len(noteOrdered) > 0 {
		identifiersDecl.WriteString("\n")
	}
	for _, note := range noteOrdered {

		identifiersDecl.WriteString(note.GongMarshallIdentifier(stage))

		initializerStatements.WriteString("\n")
		// Insertion point for basic fields value assignment
		initializerStatements.WriteString(note.GongMarshallField(stage, "Name"))
		initializerStatements.WriteString(note.GongMarshallField(stage, "Description"))
		initializerStatements.WriteString(note.GongMarshallField(stage, "ComputedPrefix"))
		initializerStatements.WriteString(note.GongMarshallField(stage, "IsExpanded"))
		initializerStatements.WriteString(note.GongMarshallField(stage, "LayoutDirection"))
		initializerStatements.WriteString(note.GongMarshallField(stage, "IsPortsNodeExpanded"))
		pointersInitializesStatements.WriteString(note.GongMarshallField(stage, "Ports"))
	}

	noteportshapeOrdered := []*NotePortShape{}
	for noteportshape := range stage.NotePortShapes {
		noteportshapeOrdered = append(noteportshapeOrdered, noteportshape)
	}
	sort.Slice(noteportshapeOrdered[:], func(i, j int) bool {
		noteportshapei := noteportshapeOrdered[i]
		noteportshapej := noteportshapeOrdered[j]
		noteportshapei_order, oki := stage.NotePortShape_stagedOrder[noteportshapei]
		noteportshapej_order, okj := stage.NotePortShape_stagedOrder[noteportshapej]
		if !oki || !okj {
			log.Fatalln("unknown pointers")
		}
		return noteportshapei_order < noteportshapej_order
	})
	if len(noteportshapeOrdered) > 0 {
		identifiersDecl.WriteString("\n")
	}
	for _, noteportshape := range noteportshapeOrdered {

		identifiersDecl.WriteString(noteportshape.GongMarshallIdentifier(stage))

		initializerStatements.WriteString("\n")
		// Insertion point for basic fields value assignment
		initializerStatements.WriteString(noteportshape.GongMarshallField(stage, "Name"))
		pointersInitializesStatements.WriteString(noteportshape.GongMarshallField(stage, "Note"))
		pointersInitializesStatements.WriteString(noteportshape.GongMarshallField(stage, "Port"))
		initializerStatements.WriteString(noteportshape.GongMarshallField(stage, "StartRatio"))
		initializerStatements.WriteString(noteportshape.GongMarshallField(stage, "EndRatio"))
		initializerStatements.WriteString(noteportshape.GongMarshallField(stage, "StartOrientation"))
		initializerStatements.WriteString(noteportshape.GongMarshallField(stage, "EndOrientation"))
		initializerStatements.WriteString(noteportshape.GongMarshallField(stage, "CornerOffsetRatio"))
		initializerStatements.WriteString(noteportshape.GongMarshallField(stage, "IsHidden"))
	}

	noteshapeOrdered := []*NoteShape{}
	for noteshape := range stage.NoteShapes {
		noteshapeOrdered = append(noteshapeOrdered, noteshape)
	}
	sort.Slice(noteshapeOrdered[:], func(i, j int) bool {
		noteshapei := noteshapeOrdered[i]
		noteshapej := noteshapeOrdered[j]
		noteshapei_order, oki := stage.NoteShape_stagedOrder[noteshapei]
		noteshapej_order, okj := stage.NoteShape_stagedOrder[noteshapej]
		if !oki || !okj {
			log.Fatalln("unknown pointers")
		}
		return noteshapei_order < noteshapej_order
	})
	if len(noteshapeOrdered) > 0 {
		identifiersDecl.WriteString("\n")
	}
	for _, noteshape := range noteshapeOrdered {

		identifiersDecl.WriteString(noteshape.GongMarshallIdentifier(stage))

		initializerStatements.WriteString("\n")
		// Insertion point for basic fields value assignment
		initializerStatements.WriteString(noteshape.GongMarshallField(stage, "Name"))
		pointersInitializesStatements.WriteString(noteshape.GongMarshallField(stage, "Note"))
		initializerStatements.WriteString(noteshape.GongMarshallField(stage, "X"))
		initializerStatements.WriteString(noteshape.GongMarshallField(stage, "Y"))
		initializerStatements.WriteString(noteshape.GongMarshallField(stage, "Width"))
		initializerStatements.WriteString(noteshape.GongMarshallField(stage, "Height"))
		initializerStatements.WriteString(noteshape.GongMarshallField(stage, "IsHidden"))
	}

	partOrdered := []*Part{}
	for part := range stage.Parts {
		partOrdered = append(partOrdered, part)
	}
	sort.Slice(partOrdered[:], func(i, j int) bool {
		parti := partOrdered[i]
		partj := partOrdered[j]
		parti_order, oki := stage.Part_stagedOrder[parti]
		partj_order, okj := stage.Part_stagedOrder[partj]
		if !oki || !okj {
			log.Fatalln("unknown pointers")
		}
		return parti_order < partj_order
	})
	if len(partOrdered) > 0 {
		identifiersDecl.WriteString("\n")
	}
	for _, part := range partOrdered {

		identifiersDecl.WriteString(part.GongMarshallIdentifier(stage))

		initializerStatements.WriteString("\n")
		// Insertion point for basic fields value assignment
		initializerStatements.WriteString(part.GongMarshallField(stage, "Name"))
		initializerStatements.WriteString(part.GongMarshallField(stage, "IsSystemResource"))
		initializerStatements.WriteString(part.GongMarshallField(stage, "Description"))
		pointersInitializesStatements.WriteString(part.GongMarshallField(stage, "Resources"))
		initializerStatements.WriteString(part.GongMarshallField(stage, "IsResourcesNodeExpanded"))
		pointersInitializesStatements.WriteString(part.GongMarshallField(stage, "Systemes"))
		initializerStatements.WriteString(part.GongMarshallField(stage, "IsSystemesNodeExpanded"))
		initializerStatements.WriteString(part.GongMarshallField(stage, "ComputedPrefix"))
		initializerStatements.WriteString(part.GongMarshallField(stage, "IsExpanded"))
		initializerStatements.WriteString(part.GongMarshallField(stage, "LayoutDirection"))
		initializerStatements.WriteString(part.GongMarshallField(stage, "IsPortsNodeExpanded"))
		pointersInitializesStatements.WriteString(part.GongMarshallField(stage, "Ports"))
		initializerStatements.WriteString(part.GongMarshallField(stage, "IsControlFlowsNodeExpanded"))
		pointersInitializesStatements.WriteString(part.GongMarshallField(stage, "ControlFlows"))
		pointersInitializesStatements.WriteString(part.GongMarshallField(stage, "PortWhoseOutControlFlowsNodeIsExpanded"))
		pointersInitializesStatements.WriteString(part.GongMarshallField(stage, "PortWhoseInControlFlowsNodeIsExpanded"))
		initializerStatements.WriteString(part.GongMarshallField(stage, "IsDataFlowsNodeExpanded"))
		pointersInitializesStatements.WriteString(part.GongMarshallField(stage, "PortWhoseOutDataFlowsNodeIsExpanded"))
		pointersInitializesStatements.WriteString(part.GongMarshallField(stage, "PortWhoseInDataFlowsNodeIsExpanded"))
	}

	partshapeOrdered := []*PartShape{}
	for partshape := range stage.PartShapes {
		partshapeOrdered = append(partshapeOrdered, partshape)
	}
	sort.Slice(partshapeOrdered[:], func(i, j int) bool {
		partshapei := partshapeOrdered[i]
		partshapej := partshapeOrdered[j]
		partshapei_order, oki := stage.PartShape_stagedOrder[partshapei]
		partshapej_order, okj := stage.PartShape_stagedOrder[partshapej]
		if !oki || !okj {
			log.Fatalln("unknown pointers")
		}
		return partshapei_order < partshapej_order
	})
	if len(partshapeOrdered) > 0 {
		identifiersDecl.WriteString("\n")
	}
	for _, partshape := range partshapeOrdered {

		identifiersDecl.WriteString(partshape.GongMarshallIdentifier(stage))

		initializerStatements.WriteString("\n")
		// Insertion point for basic fields value assignment
		initializerStatements.WriteString(partshape.GongMarshallField(stage, "Name"))
		pointersInitializesStatements.WriteString(partshape.GongMarshallField(stage, "Part"))
		initializerStatements.WriteString(partshape.GongMarshallField(stage, "IsExpanded"))
		initializerStatements.WriteString(partshape.GongMarshallField(stage, "X"))
		initializerStatements.WriteString(partshape.GongMarshallField(stage, "Y"))
		initializerStatements.WriteString(partshape.GongMarshallField(stage, "Width"))
		initializerStatements.WriteString(partshape.GongMarshallField(stage, "Height"))
		initializerStatements.WriteString(partshape.GongMarshallField(stage, "IsHidden"))
		initializerStatements.WriteString(partshape.GongMarshallField(stage, "WidthWeight"))
	}

	portOrdered := []*Port{}
	for port := range stage.Ports {
		portOrdered = append(portOrdered, port)
	}
	sort.Slice(portOrdered[:], func(i, j int) bool {
		porti := portOrdered[i]
		portj := portOrdered[j]
		porti_order, oki := stage.Port_stagedOrder[porti]
		portj_order, okj := stage.Port_stagedOrder[portj]
		if !oki || !okj {
			log.Fatalln("unknown pointers")
		}
		return porti_order < portj_order
	})
	if len(portOrdered) > 0 {
		identifiersDecl.WriteString("\n")
	}
	for _, port := range portOrdered {

		identifiersDecl.WriteString(port.GongMarshallIdentifier(stage))

		initializerStatements.WriteString("\n")
		// Insertion point for basic fields value assignment
		initializerStatements.WriteString(port.GongMarshallField(stage, "Name"))
		initializerStatements.WriteString(port.GongMarshallField(stage, "Description"))
		initializerStatements.WriteString(port.GongMarshallField(stage, "ComputedPrefix"))
		initializerStatements.WriteString(port.GongMarshallField(stage, "IsExpanded"))
		initializerStatements.WriteString(port.GongMarshallField(stage, "LayoutDirection"))
		initializerStatements.WriteString(port.GongMarshallField(stage, "IsStartPort"))
		initializerStatements.WriteString(port.GongMarshallField(stage, "IsEndPort"))
		pointersInitializesStatements.WriteString(port.GongMarshallField(stage, "Type"))
		initializerStatements.WriteString(port.GongMarshallField(stage, "IsPortNameNotSystemName"))
	}

	portshapeOrdered := []*PortShape{}
	for portshape := range stage.PortShapes {
		portshapeOrdered = append(portshapeOrdered, portshape)
	}
	sort.Slice(portshapeOrdered[:], func(i, j int) bool {
		portshapei := portshapeOrdered[i]
		portshapej := portshapeOrdered[j]
		portshapei_order, oki := stage.PortShape_stagedOrder[portshapei]
		portshapej_order, okj := stage.PortShape_stagedOrder[portshapej]
		if !oki || !okj {
			log.Fatalln("unknown pointers")
		}
		return portshapei_order < portshapej_order
	})
	if len(portshapeOrdered) > 0 {
		identifiersDecl.WriteString("\n")
	}
	for _, portshape := range portshapeOrdered {

		identifiersDecl.WriteString(portshape.GongMarshallIdentifier(stage))

		initializerStatements.WriteString("\n")
		// Insertion point for basic fields value assignment
		initializerStatements.WriteString(portshape.GongMarshallField(stage, "Name"))
		pointersInitializesStatements.WriteString(portshape.GongMarshallField(stage, "Port"))
		initializerStatements.WriteString(portshape.GongMarshallField(stage, "IsExpanded"))
		initializerStatements.WriteString(portshape.GongMarshallField(stage, "X"))
		initializerStatements.WriteString(portshape.GongMarshallField(stage, "Y"))
		initializerStatements.WriteString(portshape.GongMarshallField(stage, "Width"))
		initializerStatements.WriteString(portshape.GongMarshallField(stage, "Height"))
		initializerStatements.WriteString(portshape.GongMarshallField(stage, "IsHidden"))
	}

	resourceOrdered := []*Resource{}
	for resource := range stage.Resources {
		resourceOrdered = append(resourceOrdered, resource)
	}
	sort.Slice(resourceOrdered[:], func(i, j int) bool {
		resourcei := resourceOrdered[i]
		resourcej := resourceOrdered[j]
		resourcei_order, oki := stage.Resource_stagedOrder[resourcei]
		resourcej_order, okj := stage.Resource_stagedOrder[resourcej]
		if !oki || !okj {
			log.Fatalln("unknown pointers")
		}
		return resourcei_order < resourcej_order
	})
	if len(resourceOrdered) > 0 {
		identifiersDecl.WriteString("\n")
	}
	for _, resource := range resourceOrdered {

		identifiersDecl.WriteString(resource.GongMarshallIdentifier(stage))

		initializerStatements.WriteString("\n")
		// Insertion point for basic fields value assignment
		initializerStatements.WriteString(resource.GongMarshallField(stage, "Name"))
		initializerStatements.WriteString(resource.GongMarshallField(stage, "Acronym"))
		initializerStatements.WriteString(resource.GongMarshallField(stage, "Description"))
		initializerStatements.WriteString(resource.GongMarshallField(stage, "ComputedPrefix"))
		initializerStatements.WriteString(resource.GongMarshallField(stage, "IsExpanded"))
		initializerStatements.WriteString(resource.GongMarshallField(stage, "LayoutDirection"))
		initializerStatements.WriteString(resource.GongMarshallField(stage, "SVG_Path"))
		initializerStatements.WriteString(resource.GongMarshallField(stage, "InverseAppliedScaling"))
	}

	systemOrdered := []*System{}
	for system := range stage.Systems {
		systemOrdered = append(systemOrdered, system)
	}
	sort.Slice(systemOrdered[:], func(i, j int) bool {
		systemi := systemOrdered[i]
		systemj := systemOrdered[j]
		systemi_order, oki := stage.System_stagedOrder[systemi]
		systemj_order, okj := stage.System_stagedOrder[systemj]
		if !oki || !okj {
			log.Fatalln("unknown pointers")
		}
		return systemi_order < systemj_order
	})
	if len(systemOrdered) > 0 {
		identifiersDecl.WriteString("\n")
	}
	for _, system := range systemOrdered {

		identifiersDecl.WriteString(system.GongMarshallIdentifier(stage))

		initializerStatements.WriteString("\n")
		// Insertion point for basic fields value assignment
		initializerStatements.WriteString(system.GongMarshallField(stage, "Name"))
		initializerStatements.WriteString(system.GongMarshallField(stage, "Description"))
		initializerStatements.WriteString(system.GongMarshallField(stage, "ComputedPrefix"))
		initializerStatements.WriteString(system.GongMarshallField(stage, "IsExpanded"))
		initializerStatements.WriteString(system.GongMarshallField(stage, "LayoutDirection"))
		initializerStatements.WriteString(system.GongMarshallField(stage, "SVG_Path"))
		initializerStatements.WriteString(system.GongMarshallField(stage, "InverseAppliedScaling"))
		pointersInitializesStatements.WriteString(system.GongMarshallField(stage, "DiagramStructures"))
		pointersInitializesStatements.WriteString(system.GongMarshallField(stage, "DiagramStructureWhoseNodeIsExpanded"))
		initializerStatements.WriteString(system.GongMarshallField(stage, "IsSubSystemNodeExpanded"))
		pointersInitializesStatements.WriteString(system.GongMarshallField(stage, "SubSystemes"))
		pointersInitializesStatements.WriteString(system.GongMarshallField(stage, "Parts"))
		pointersInitializesStatements.WriteString(system.GongMarshallField(stage, "PartWhoseNodeIsExpanded"))
		pointersInitializesStatements.WriteString(system.GongMarshallField(stage, "DataFlows"))
		initializerStatements.WriteString(system.GongMarshallField(stage, "IsDataFlowsNodeExpanded"))
		pointersInitializesStatements.WriteString(system.GongMarshallField(stage, "ExternalParts"))
		pointersInitializesStatements.WriteString(system.GongMarshallField(stage, "ExternalPartWhoseNodeIsExpanded"))
	}

	systemshapeOrdered := []*SystemShape{}
	for systemshape := range stage.SystemShapes {
		systemshapeOrdered = append(systemshapeOrdered, systemshape)
	}
	sort.Slice(systemshapeOrdered[:], func(i, j int) bool {
		systemshapei := systemshapeOrdered[i]
		systemshapej := systemshapeOrdered[j]
		systemshapei_order, oki := stage.SystemShape_stagedOrder[systemshapei]
		systemshapej_order, okj := stage.SystemShape_stagedOrder[systemshapej]
		if !oki || !okj {
			log.Fatalln("unknown pointers")
		}
		return systemshapei_order < systemshapej_order
	})
	if len(systemshapeOrdered) > 0 {
		identifiersDecl.WriteString("\n")
	}
	for _, systemshape := range systemshapeOrdered {

		identifiersDecl.WriteString(systemshape.GongMarshallIdentifier(stage))

		initializerStatements.WriteString("\n")
		// Insertion point for basic fields value assignment
		initializerStatements.WriteString(systemshape.GongMarshallField(stage, "Name"))
		pointersInitializesStatements.WriteString(systemshape.GongMarshallField(stage, "System"))
		initializerStatements.WriteString(systemshape.GongMarshallField(stage, "IsExpanded"))
		initializerStatements.WriteString(systemshape.GongMarshallField(stage, "X"))
		initializerStatements.WriteString(systemshape.GongMarshallField(stage, "Y"))
		initializerStatements.WriteString(systemshape.GongMarshallField(stage, "Width"))
		initializerStatements.WriteString(systemshape.GongMarshallField(stage, "Height"))
		initializerStatements.WriteString(systemshape.GongMarshallField(stage, "IsHidden"))
	}

	// insertion initialization of objects to stage
	for _, allocatedresourceshape := range allocatedresourceshapeOrdered {
		_ = allocatedresourceshape
		var setPointerField string
		_ = setPointerField

		// Insertion point for pointers initialization
	}

	for _, allocatedsystemshape := range allocatedsystemshapeOrdered {
		_ = allocatedsystemshape
		var setPointerField string
		_ = setPointerField

		// Insertion point for pointers initialization
	}

	for _, controlflow := range controlflowOrdered {
		_ = controlflow
		var setPointerField string
		_ = setPointerField

		// Insertion point for pointers initialization
	}

	for _, controlflowshape := range controlflowshapeOrdered {
		_ = controlflowshape
		var setPointerField string
		_ = setPointerField

		// Insertion point for pointers initialization
	}

	for _, data := range dataOrdered {
		_ = data
		var setPointerField string
		_ = setPointerField

		// Insertion point for pointers initialization
	}

	for _, dataflow := range dataflowOrdered {
		_ = dataflow
		var setPointerField string
		_ = setPointerField

		// Insertion point for pointers initialization
	}

	for _, dataflowshape := range dataflowshapeOrdered {
		_ = dataflowshape
		var setPointerField string
		_ = setPointerField

		// Insertion point for pointers initialization
	}

	for _, datashape := range datashapeOrdered {
		_ = datashape
		var setPointerField string
		_ = setPointerField

		// Insertion point for pointers initialization
	}

	for _, diagramstructure := range diagramstructureOrdered {
		_ = diagramstructure
		var setPointerField string
		_ = setPointerField

		// Insertion point for pointers initialization
	}

	for _, externalpartshape := range externalpartshapeOrdered {
		_ = externalpartshape
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

	for _, note := range noteOrdered {
		_ = note
		var setPointerField string
		_ = setPointerField

		// Insertion point for pointers initialization
	}

	for _, noteportshape := range noteportshapeOrdered {
		_ = noteportshape
		var setPointerField string
		_ = setPointerField

		// Insertion point for pointers initialization
	}

	for _, noteshape := range noteshapeOrdered {
		_ = noteshape
		var setPointerField string
		_ = setPointerField

		// Insertion point for pointers initialization
	}

	for _, part := range partOrdered {
		_ = part
		var setPointerField string
		_ = setPointerField

		// Insertion point for pointers initialization
	}

	for _, partshape := range partshapeOrdered {
		_ = partshape
		var setPointerField string
		_ = setPointerField

		// Insertion point for pointers initialization
	}

	for _, port := range portOrdered {
		_ = port
		var setPointerField string
		_ = setPointerField

		// Insertion point for pointers initialization
	}

	for _, portshape := range portshapeOrdered {
		_ = portshape
		var setPointerField string
		_ = setPointerField

		// Insertion point for pointers initialization
	}

	for _, resource := range resourceOrdered {
		_ = resource
		var setPointerField string
		_ = setPointerField

		// Insertion point for pointers initialization
	}

	for _, system := range systemOrdered {
		_ = system
		var setPointerField string
		_ = setPointerField

		// Insertion point for pointers initialization
	}

	for _, systemshape := range systemshapeOrdered {
		_ = systemshape
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
func (allocatedresourceshape *AllocatedResourceShape) GongMarshallField(stage *Stage, fieldName string) (res string) {

	switch fieldName {
	case "Name":
		res = StringInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", allocatedresourceshape.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "Name")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", ToRawStringLiteral(allocatedresourceshape.Name))

	case "Part":
		if allocatedresourceshape.Part != nil {
			res = PointerFieldInitStatement
			res = strings.ReplaceAll(res, "{{Identifier}}", allocatedresourceshape.GongGetIdentifier(stage))
			res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "Part")
			res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", allocatedresourceshape.Part.GongGetIdentifier(stage))
		} else {
			// in case of nil pointer, we need to unstage the previous value
			res = PointerFieldInitStatement
			res = strings.ReplaceAll(res, "{{Identifier}}", allocatedresourceshape.GongGetIdentifier(stage))
			res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "Part")
			res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", "nil")
		}
	case "Resource":
		if allocatedresourceshape.Resource != nil {
			res = PointerFieldInitStatement
			res = strings.ReplaceAll(res, "{{Identifier}}", allocatedresourceshape.GongGetIdentifier(stage))
			res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "Resource")
			res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", allocatedresourceshape.Resource.GongGetIdentifier(stage))
		} else {
			// in case of nil pointer, we need to unstage the previous value
			res = PointerFieldInitStatement
			res = strings.ReplaceAll(res, "{{Identifier}}", allocatedresourceshape.GongGetIdentifier(stage))
			res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "Resource")
			res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", "nil")
		}
	default:
		log.Panicf("Unknown field %s for Gongstruct AllocatedResourceShape", fieldName)
	}
	return
}

func (allocatedsystemshape *AllocatedSystemShape) GongMarshallField(stage *Stage, fieldName string) (res string) {

	switch fieldName {
	case "Name":
		res = StringInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", allocatedsystemshape.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "Name")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", ToRawStringLiteral(allocatedsystemshape.Name))

	case "Part":
		if allocatedsystemshape.Part != nil {
			res = PointerFieldInitStatement
			res = strings.ReplaceAll(res, "{{Identifier}}", allocatedsystemshape.GongGetIdentifier(stage))
			res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "Part")
			res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", allocatedsystemshape.Part.GongGetIdentifier(stage))
		} else {
			// in case of nil pointer, we need to unstage the previous value
			res = PointerFieldInitStatement
			res = strings.ReplaceAll(res, "{{Identifier}}", allocatedsystemshape.GongGetIdentifier(stage))
			res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "Part")
			res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", "nil")
		}
	case "System":
		if allocatedsystemshape.System != nil {
			res = PointerFieldInitStatement
			res = strings.ReplaceAll(res, "{{Identifier}}", allocatedsystemshape.GongGetIdentifier(stage))
			res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "System")
			res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", allocatedsystemshape.System.GongGetIdentifier(stage))
		} else {
			// in case of nil pointer, we need to unstage the previous value
			res = PointerFieldInitStatement
			res = strings.ReplaceAll(res, "{{Identifier}}", allocatedsystemshape.GongGetIdentifier(stage))
			res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "System")
			res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", "nil")
		}
	default:
		log.Panicf("Unknown field %s for Gongstruct AllocatedSystemShape", fieldName)
	}
	return
}

func (controlflow *ControlFlow) GongMarshallField(stage *Stage, fieldName string) (res string) {

	switch fieldName {
	case "Name":
		res = StringInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", controlflow.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "Name")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", ToRawStringLiteral(controlflow.Name))
	case "Description":
		res = StringInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", controlflow.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "Description")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", ToRawStringLiteral(controlflow.Description))
	case "ComputedPrefix":
		res = StringInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", controlflow.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "ComputedPrefix")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", ToRawStringLiteral(controlflow.ComputedPrefix))
	case "IsExpanded":
		res = NumberInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", controlflow.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "IsExpanded")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%t", controlflow.IsExpanded))
	case "LayoutDirection":
		if controlflow.LayoutDirection.ToCodeString() != "" {
			res = NumberInitStatement
			res = strings.ReplaceAll(res, "{{Identifier}}", controlflow.GongGetIdentifier(stage))
			res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "LayoutDirection")
			res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", "models."+controlflow.LayoutDirection.ToCodeString())
		} else {
			// in case of empty enum, we need to unstage the previous value
			res = NumberInitStatement
			res = strings.ReplaceAll(res, "{{Identifier}}", controlflow.GongGetIdentifier(stage))
			res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "LayoutDirection")
			res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", "0")
		}

	case "Start":
		if controlflow.Start != nil {
			res = PointerFieldInitStatement
			res = strings.ReplaceAll(res, "{{Identifier}}", controlflow.GongGetIdentifier(stage))
			res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "Start")
			res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", controlflow.Start.GongGetIdentifier(stage))
		} else {
			// in case of nil pointer, we need to unstage the previous value
			res = PointerFieldInitStatement
			res = strings.ReplaceAll(res, "{{Identifier}}", controlflow.GongGetIdentifier(stage))
			res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "Start")
			res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", "nil")
		}
	case "End":
		if controlflow.End != nil {
			res = PointerFieldInitStatement
			res = strings.ReplaceAll(res, "{{Identifier}}", controlflow.GongGetIdentifier(stage))
			res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "End")
			res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", controlflow.End.GongGetIdentifier(stage))
		} else {
			// in case of nil pointer, we need to unstage the previous value
			res = PointerFieldInitStatement
			res = strings.ReplaceAll(res, "{{Identifier}}", controlflow.GongGetIdentifier(stage))
			res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "End")
			res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", "nil")
		}
	default:
		log.Panicf("Unknown field %s for Gongstruct ControlFlow", fieldName)
	}
	return
}

func (controlflowshape *ControlFlowShape) GongMarshallField(stage *Stage, fieldName string) (res string) {

	switch fieldName {
	case "Name":
		res = StringInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", controlflowshape.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "Name")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", ToRawStringLiteral(controlflowshape.Name))
	case "StartRatio":
		res = NumberInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", controlflowshape.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "StartRatio")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%f", controlflowshape.StartRatio))
	case "EndRatio":
		res = NumberInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", controlflowshape.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "EndRatio")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%f", controlflowshape.EndRatio))
	case "StartOrientation":
		if controlflowshape.StartOrientation.ToCodeString() != "" {
			res = StringEnumInitStatement
			res = strings.ReplaceAll(res, "{{Identifier}}", controlflowshape.GongGetIdentifier(stage))
			res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "StartOrientation")
			res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", "models."+controlflowshape.StartOrientation.ToCodeString())
		} else {
			// in case of empty enum, we need to unstage the previous value
			res = StringEnumInitStatement
			res = strings.ReplaceAll(res, "{{Identifier}}", controlflowshape.GongGetIdentifier(stage))
			res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "StartOrientation")
			res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", "\"\"")
		}
	case "EndOrientation":
		if controlflowshape.EndOrientation.ToCodeString() != "" {
			res = StringEnumInitStatement
			res = strings.ReplaceAll(res, "{{Identifier}}", controlflowshape.GongGetIdentifier(stage))
			res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "EndOrientation")
			res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", "models."+controlflowshape.EndOrientation.ToCodeString())
		} else {
			// in case of empty enum, we need to unstage the previous value
			res = StringEnumInitStatement
			res = strings.ReplaceAll(res, "{{Identifier}}", controlflowshape.GongGetIdentifier(stage))
			res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "EndOrientation")
			res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", "\"\"")
		}
	case "CornerOffsetRatio":
		res = NumberInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", controlflowshape.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "CornerOffsetRatio")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%f", controlflowshape.CornerOffsetRatio))
	case "IsHidden":
		res = NumberInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", controlflowshape.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "IsHidden")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%t", controlflowshape.IsHidden))

	case "ControlFlow":
		if controlflowshape.ControlFlow != nil {
			res = PointerFieldInitStatement
			res = strings.ReplaceAll(res, "{{Identifier}}", controlflowshape.GongGetIdentifier(stage))
			res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "ControlFlow")
			res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", controlflowshape.ControlFlow.GongGetIdentifier(stage))
		} else {
			// in case of nil pointer, we need to unstage the previous value
			res = PointerFieldInitStatement
			res = strings.ReplaceAll(res, "{{Identifier}}", controlflowshape.GongGetIdentifier(stage))
			res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "ControlFlow")
			res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", "nil")
		}
	default:
		log.Panicf("Unknown field %s for Gongstruct ControlFlowShape", fieldName)
	}
	return
}

func (data *Data) GongMarshallField(stage *Stage, fieldName string) (res string) {

	switch fieldName {
	case "Name":
		res = StringInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", data.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "Name")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", ToRawStringLiteral(data.Name))
	case "Acronym":
		res = StringInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", data.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "Acronym")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", ToRawStringLiteral(data.Acronym))
	case "Description":
		res = StringInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", data.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "Description")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", ToRawStringLiteral(data.Description))
	case "ComputedPrefix":
		res = StringInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", data.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "ComputedPrefix")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", ToRawStringLiteral(data.ComputedPrefix))
	case "IsExpanded":
		res = NumberInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", data.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "IsExpanded")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%t", data.IsExpanded))
	case "LayoutDirection":
		if data.LayoutDirection.ToCodeString() != "" {
			res = NumberInitStatement
			res = strings.ReplaceAll(res, "{{Identifier}}", data.GongGetIdentifier(stage))
			res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "LayoutDirection")
			res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", "models."+data.LayoutDirection.ToCodeString())
		} else {
			// in case of empty enum, we need to unstage the previous value
			res = NumberInitStatement
			res = strings.ReplaceAll(res, "{{Identifier}}", data.GongGetIdentifier(stage))
			res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "LayoutDirection")
			res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", "0")
		}
	case "SVG_Path":
		res = StringInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", data.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "SVG_Path")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", ToRawStringLiteral(data.SVG_Path))
	case "InverseAppliedScaling":
		res = NumberInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", data.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "InverseAppliedScaling")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%f", data.InverseAppliedScaling))

	default:
		log.Panicf("Unknown field %s for Gongstruct Data", fieldName)
	}
	return
}

func (dataflow *DataFlow) GongMarshallField(stage *Stage, fieldName string) (res string) {

	switch fieldName {
	case "Name":
		res = StringInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", dataflow.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "Name")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", ToRawStringLiteral(dataflow.Name))
	case "Description":
		res = StringInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", dataflow.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "Description")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", ToRawStringLiteral(dataflow.Description))
	case "ComputedPrefix":
		res = StringInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", dataflow.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "ComputedPrefix")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", ToRawStringLiteral(dataflow.ComputedPrefix))
	case "IsExpanded":
		res = NumberInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", dataflow.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "IsExpanded")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%t", dataflow.IsExpanded))
	case "LayoutDirection":
		if dataflow.LayoutDirection.ToCodeString() != "" {
			res = NumberInitStatement
			res = strings.ReplaceAll(res, "{{Identifier}}", dataflow.GongGetIdentifier(stage))
			res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "LayoutDirection")
			res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", "models."+dataflow.LayoutDirection.ToCodeString())
		} else {
			// in case of empty enum, we need to unstage the previous value
			res = NumberInitStatement
			res = strings.ReplaceAll(res, "{{Identifier}}", dataflow.GongGetIdentifier(stage))
			res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "LayoutDirection")
			res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", "0")
		}
	case "Type":
		if dataflow.Type.ToCodeString() != "" {
			res = StringEnumInitStatement
			res = strings.ReplaceAll(res, "{{Identifier}}", dataflow.GongGetIdentifier(stage))
			res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "Type")
			res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", "models."+dataflow.Type.ToCodeString())
		} else {
			// in case of empty enum, we need to unstage the previous value
			res = StringEnumInitStatement
			res = strings.ReplaceAll(res, "{{Identifier}}", dataflow.GongGetIdentifier(stage))
			res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "Type")
			res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", "\"\"")
		}
	case "IsDatasNodeExpanded":
		res = NumberInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", dataflow.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "IsDatasNodeExpanded")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%t", dataflow.IsDatasNodeExpanded))

	case "Datas":
		var sb strings.Builder
		for _, _data := range dataflow.Datas {
			tmp := SliceOfPointersFieldInitStatement
			tmp = strings.ReplaceAll(tmp, "{{Identifier}}", dataflow.GongGetIdentifier(stage))
			tmp = strings.ReplaceAll(tmp, "{{GeneratedFieldName}}", "Datas")
			tmp = strings.ReplaceAll(tmp, "{{GeneratedFieldNameValue}}", _data.GongGetIdentifier(stage))
			sb.WriteString(tmp)
		}
		res = sb.String()
	case "StartPort":
		if dataflow.StartPort != nil {
			res = PointerFieldInitStatement
			res = strings.ReplaceAll(res, "{{Identifier}}", dataflow.GongGetIdentifier(stage))
			res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "StartPort")
			res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", dataflow.StartPort.GongGetIdentifier(stage))
		} else {
			// in case of nil pointer, we need to unstage the previous value
			res = PointerFieldInitStatement
			res = strings.ReplaceAll(res, "{{Identifier}}", dataflow.GongGetIdentifier(stage))
			res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "StartPort")
			res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", "nil")
		}
	case "EndPort":
		if dataflow.EndPort != nil {
			res = PointerFieldInitStatement
			res = strings.ReplaceAll(res, "{{Identifier}}", dataflow.GongGetIdentifier(stage))
			res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "EndPort")
			res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", dataflow.EndPort.GongGetIdentifier(stage))
		} else {
			// in case of nil pointer, we need to unstage the previous value
			res = PointerFieldInitStatement
			res = strings.ReplaceAll(res, "{{Identifier}}", dataflow.GongGetIdentifier(stage))
			res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "EndPort")
			res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", "nil")
		}
	case "StartExternalPart":
		if dataflow.StartExternalPart != nil {
			res = PointerFieldInitStatement
			res = strings.ReplaceAll(res, "{{Identifier}}", dataflow.GongGetIdentifier(stage))
			res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "StartExternalPart")
			res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", dataflow.StartExternalPart.GongGetIdentifier(stage))
		} else {
			// in case of nil pointer, we need to unstage the previous value
			res = PointerFieldInitStatement
			res = strings.ReplaceAll(res, "{{Identifier}}", dataflow.GongGetIdentifier(stage))
			res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "StartExternalPart")
			res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", "nil")
		}
	case "EndExternalPart":
		if dataflow.EndExternalPart != nil {
			res = PointerFieldInitStatement
			res = strings.ReplaceAll(res, "{{Identifier}}", dataflow.GongGetIdentifier(stage))
			res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "EndExternalPart")
			res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", dataflow.EndExternalPart.GongGetIdentifier(stage))
		} else {
			// in case of nil pointer, we need to unstage the previous value
			res = PointerFieldInitStatement
			res = strings.ReplaceAll(res, "{{Identifier}}", dataflow.GongGetIdentifier(stage))
			res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "EndExternalPart")
			res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", "nil")
		}
	default:
		log.Panicf("Unknown field %s for Gongstruct DataFlow", fieldName)
	}
	return
}

func (dataflowshape *DataFlowShape) GongMarshallField(stage *Stage, fieldName string) (res string) {

	switch fieldName {
	case "Name":
		res = StringInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", dataflowshape.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "Name")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", ToRawStringLiteral(dataflowshape.Name))
	case "StartRatio":
		res = NumberInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", dataflowshape.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "StartRatio")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%f", dataflowshape.StartRatio))
	case "EndRatio":
		res = NumberInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", dataflowshape.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "EndRatio")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%f", dataflowshape.EndRatio))
	case "StartOrientation":
		if dataflowshape.StartOrientation.ToCodeString() != "" {
			res = StringEnumInitStatement
			res = strings.ReplaceAll(res, "{{Identifier}}", dataflowshape.GongGetIdentifier(stage))
			res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "StartOrientation")
			res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", "models."+dataflowshape.StartOrientation.ToCodeString())
		} else {
			// in case of empty enum, we need to unstage the previous value
			res = StringEnumInitStatement
			res = strings.ReplaceAll(res, "{{Identifier}}", dataflowshape.GongGetIdentifier(stage))
			res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "StartOrientation")
			res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", "\"\"")
		}
	case "EndOrientation":
		if dataflowshape.EndOrientation.ToCodeString() != "" {
			res = StringEnumInitStatement
			res = strings.ReplaceAll(res, "{{Identifier}}", dataflowshape.GongGetIdentifier(stage))
			res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "EndOrientation")
			res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", "models."+dataflowshape.EndOrientation.ToCodeString())
		} else {
			// in case of empty enum, we need to unstage the previous value
			res = StringEnumInitStatement
			res = strings.ReplaceAll(res, "{{Identifier}}", dataflowshape.GongGetIdentifier(stage))
			res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "EndOrientation")
			res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", "\"\"")
		}
	case "CornerOffsetRatio":
		res = NumberInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", dataflowshape.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "CornerOffsetRatio")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%f", dataflowshape.CornerOffsetRatio))
	case "IsHidden":
		res = NumberInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", dataflowshape.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "IsHidden")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%t", dataflowshape.IsHidden))

	case "DataFlow":
		if dataflowshape.DataFlow != nil {
			res = PointerFieldInitStatement
			res = strings.ReplaceAll(res, "{{Identifier}}", dataflowshape.GongGetIdentifier(stage))
			res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "DataFlow")
			res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", dataflowshape.DataFlow.GongGetIdentifier(stage))
		} else {
			// in case of nil pointer, we need to unstage the previous value
			res = PointerFieldInitStatement
			res = strings.ReplaceAll(res, "{{Identifier}}", dataflowshape.GongGetIdentifier(stage))
			res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "DataFlow")
			res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", "nil")
		}
	default:
		log.Panicf("Unknown field %s for Gongstruct DataFlowShape", fieldName)
	}
	return
}

func (datashape *DataShape) GongMarshallField(stage *Stage, fieldName string) (res string) {

	switch fieldName {
	case "Name":
		res = StringInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", datashape.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "Name")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", ToRawStringLiteral(datashape.Name))

	case "Data":
		if datashape.Data != nil {
			res = PointerFieldInitStatement
			res = strings.ReplaceAll(res, "{{Identifier}}", datashape.GongGetIdentifier(stage))
			res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "Data")
			res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", datashape.Data.GongGetIdentifier(stage))
		} else {
			// in case of nil pointer, we need to unstage the previous value
			res = PointerFieldInitStatement
			res = strings.ReplaceAll(res, "{{Identifier}}", datashape.GongGetIdentifier(stage))
			res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "Data")
			res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", "nil")
		}
	case "DataFlow":
		if datashape.DataFlow != nil {
			res = PointerFieldInitStatement
			res = strings.ReplaceAll(res, "{{Identifier}}", datashape.GongGetIdentifier(stage))
			res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "DataFlow")
			res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", datashape.DataFlow.GongGetIdentifier(stage))
		} else {
			// in case of nil pointer, we need to unstage the previous value
			res = PointerFieldInitStatement
			res = strings.ReplaceAll(res, "{{Identifier}}", datashape.GongGetIdentifier(stage))
			res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "DataFlow")
			res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", "nil")
		}
	default:
		log.Panicf("Unknown field %s for Gongstruct DataShape", fieldName)
	}
	return
}

func (diagramstructure *DiagramStructure) GongMarshallField(stage *Stage, fieldName string) (res string) {

	switch fieldName {
	case "Name":
		res = StringInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", diagramstructure.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "Name")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", ToRawStringLiteral(diagramstructure.Name))
	case "Description":
		res = StringInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", diagramstructure.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "Description")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", ToRawStringLiteral(diagramstructure.Description))
	case "ComputedPrefix":
		res = StringInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", diagramstructure.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "ComputedPrefix")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", ToRawStringLiteral(diagramstructure.ComputedPrefix))
	case "IsExpanded":
		res = NumberInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", diagramstructure.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "IsExpanded")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%t", diagramstructure.IsExpanded))
	case "LayoutDirection":
		if diagramstructure.LayoutDirection.ToCodeString() != "" {
			res = NumberInitStatement
			res = strings.ReplaceAll(res, "{{Identifier}}", diagramstructure.GongGetIdentifier(stage))
			res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "LayoutDirection")
			res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", "models."+diagramstructure.LayoutDirection.ToCodeString())
		} else {
			// in case of empty enum, we need to unstage the previous value
			res = NumberInitStatement
			res = strings.ReplaceAll(res, "{{Identifier}}", diagramstructure.GongGetIdentifier(stage))
			res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "LayoutDirection")
			res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", "0")
		}
	case "IsChecked":
		res = NumberInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", diagramstructure.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "IsChecked")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%t", diagramstructure.IsChecked))
	case "IsEditable_":
		res = NumberInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", diagramstructure.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "IsEditable_")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%t", diagramstructure.IsEditable_))
	case "IsShowPrefix":
		res = NumberInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", diagramstructure.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "IsShowPrefix")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%t", diagramstructure.IsShowPrefix))
	case "DefaultBoxWidth":
		res = NumberInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", diagramstructure.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "DefaultBoxWidth")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%f", diagramstructure.DefaultBoxWidth))
	case "DefaultBoxHeigth":
		res = NumberInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", diagramstructure.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "DefaultBoxHeigth")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%f", diagramstructure.DefaultBoxHeigth))
	case "Width":
		res = NumberInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", diagramstructure.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "Width")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%f", diagramstructure.Width))
	case "Height":
		res = NumberInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", diagramstructure.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "Height")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%f", diagramstructure.Height))
	case "IsSystemsNodeExpanded":
		res = NumberInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", diagramstructure.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "IsSystemsNodeExpanded")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%t", diagramstructure.IsSystemsNodeExpanded))
	case "IsPartsNodeExpanded":
		res = NumberInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", diagramstructure.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "IsPartsNodeExpanded")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%t", diagramstructure.IsPartsNodeExpanded))
	case "IsExternalPartsNodeExpanded":
		res = NumberInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", diagramstructure.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "IsExternalPartsNodeExpanded")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%t", diagramstructure.IsExternalPartsNodeExpanded))
	case "IsNotesNodeExpanded":
		res = NumberInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", diagramstructure.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "IsNotesNodeExpanded")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%t", diagramstructure.IsNotesNodeExpanded))

	case "System_Shapes":
		var sb strings.Builder
		for _, _systemshape := range diagramstructure.System_Shapes {
			tmp := SliceOfPointersFieldInitStatement
			tmp = strings.ReplaceAll(tmp, "{{Identifier}}", diagramstructure.GongGetIdentifier(stage))
			tmp = strings.ReplaceAll(tmp, "{{GeneratedFieldName}}", "System_Shapes")
			tmp = strings.ReplaceAll(tmp, "{{GeneratedFieldNameValue}}", _systemshape.GongGetIdentifier(stage))
			sb.WriteString(tmp)
		}
		res = sb.String()
	case "SystemsWhoseNodeIsExpanded":
		var sb strings.Builder
		for _, _system := range diagramstructure.SystemsWhoseNodeIsExpanded {
			tmp := SliceOfPointersFieldInitStatement
			tmp = strings.ReplaceAll(tmp, "{{Identifier}}", diagramstructure.GongGetIdentifier(stage))
			tmp = strings.ReplaceAll(tmp, "{{GeneratedFieldName}}", "SystemsWhoseNodeIsExpanded")
			tmp = strings.ReplaceAll(tmp, "{{GeneratedFieldNameValue}}", _system.GongGetIdentifier(stage))
			sb.WriteString(tmp)
		}
		res = sb.String()
	case "Part_Shapes":
		var sb strings.Builder
		for _, _partshape := range diagramstructure.Part_Shapes {
			tmp := SliceOfPointersFieldInitStatement
			tmp = strings.ReplaceAll(tmp, "{{Identifier}}", diagramstructure.GongGetIdentifier(stage))
			tmp = strings.ReplaceAll(tmp, "{{GeneratedFieldName}}", "Part_Shapes")
			tmp = strings.ReplaceAll(tmp, "{{GeneratedFieldNameValue}}", _partshape.GongGetIdentifier(stage))
			sb.WriteString(tmp)
		}
		res = sb.String()
	case "PartWhoseNodeIsExpanded":
		var sb strings.Builder
		for _, _part := range diagramstructure.PartWhoseNodeIsExpanded {
			tmp := SliceOfPointersFieldInitStatement
			tmp = strings.ReplaceAll(tmp, "{{Identifier}}", diagramstructure.GongGetIdentifier(stage))
			tmp = strings.ReplaceAll(tmp, "{{GeneratedFieldName}}", "PartWhoseNodeIsExpanded")
			tmp = strings.ReplaceAll(tmp, "{{GeneratedFieldNameValue}}", _part.GongGetIdentifier(stage))
			sb.WriteString(tmp)
		}
		res = sb.String()
	case "ExternalPart_Shapes":
		var sb strings.Builder
		for _, _externalpartshape := range diagramstructure.ExternalPart_Shapes {
			tmp := SliceOfPointersFieldInitStatement
			tmp = strings.ReplaceAll(tmp, "{{Identifier}}", diagramstructure.GongGetIdentifier(stage))
			tmp = strings.ReplaceAll(tmp, "{{GeneratedFieldName}}", "ExternalPart_Shapes")
			tmp = strings.ReplaceAll(tmp, "{{GeneratedFieldNameValue}}", _externalpartshape.GongGetIdentifier(stage))
			sb.WriteString(tmp)
		}
		res = sb.String()
	case "ExternalPartWhoseNodeIsExpanded":
		var sb strings.Builder
		for _, _part := range diagramstructure.ExternalPartWhoseNodeIsExpanded {
			tmp := SliceOfPointersFieldInitStatement
			tmp = strings.ReplaceAll(tmp, "{{Identifier}}", diagramstructure.GongGetIdentifier(stage))
			tmp = strings.ReplaceAll(tmp, "{{GeneratedFieldName}}", "ExternalPartWhoseNodeIsExpanded")
			tmp = strings.ReplaceAll(tmp, "{{GeneratedFieldNameValue}}", _part.GongGetIdentifier(stage))
			sb.WriteString(tmp)
		}
		res = sb.String()
	case "ExternalPartsWhoseOutDataFlowsNodeIsExpanded":
		var sb strings.Builder
		for _, _part := range diagramstructure.ExternalPartsWhoseOutDataFlowsNodeIsExpanded {
			tmp := SliceOfPointersFieldInitStatement
			tmp = strings.ReplaceAll(tmp, "{{Identifier}}", diagramstructure.GongGetIdentifier(stage))
			tmp = strings.ReplaceAll(tmp, "{{GeneratedFieldName}}", "ExternalPartsWhoseOutDataFlowsNodeIsExpanded")
			tmp = strings.ReplaceAll(tmp, "{{GeneratedFieldNameValue}}", _part.GongGetIdentifier(stage))
			sb.WriteString(tmp)
		}
		res = sb.String()
	case "ExternalPartsWhoseInDataFlowsNodeIsExpanded":
		var sb strings.Builder
		for _, _part := range diagramstructure.ExternalPartsWhoseInDataFlowsNodeIsExpanded {
			tmp := SliceOfPointersFieldInitStatement
			tmp = strings.ReplaceAll(tmp, "{{Identifier}}", diagramstructure.GongGetIdentifier(stage))
			tmp = strings.ReplaceAll(tmp, "{{GeneratedFieldName}}", "ExternalPartsWhoseInDataFlowsNodeIsExpanded")
			tmp = strings.ReplaceAll(tmp, "{{GeneratedFieldNameValue}}", _part.GongGetIdentifier(stage))
			sb.WriteString(tmp)
		}
		res = sb.String()
	case "PortsWhoseNodeIsExpanded":
		var sb strings.Builder
		for _, _port := range diagramstructure.PortsWhoseNodeIsExpanded {
			tmp := SliceOfPointersFieldInitStatement
			tmp = strings.ReplaceAll(tmp, "{{Identifier}}", diagramstructure.GongGetIdentifier(stage))
			tmp = strings.ReplaceAll(tmp, "{{GeneratedFieldName}}", "PortsWhoseNodeIsExpanded")
			tmp = strings.ReplaceAll(tmp, "{{GeneratedFieldNameValue}}", _port.GongGetIdentifier(stage))
			sb.WriteString(tmp)
		}
		res = sb.String()
	case "Port_Shapes":
		var sb strings.Builder
		for _, _portshape := range diagramstructure.Port_Shapes {
			tmp := SliceOfPointersFieldInitStatement
			tmp = strings.ReplaceAll(tmp, "{{Identifier}}", diagramstructure.GongGetIdentifier(stage))
			tmp = strings.ReplaceAll(tmp, "{{GeneratedFieldName}}", "Port_Shapes")
			tmp = strings.ReplaceAll(tmp, "{{GeneratedFieldNameValue}}", _portshape.GongGetIdentifier(stage))
			sb.WriteString(tmp)
		}
		res = sb.String()
	case "ControlFlowsWhoseNodeIsExpanded":
		var sb strings.Builder
		for _, _controlflow := range diagramstructure.ControlFlowsWhoseNodeIsExpanded {
			tmp := SliceOfPointersFieldInitStatement
			tmp = strings.ReplaceAll(tmp, "{{Identifier}}", diagramstructure.GongGetIdentifier(stage))
			tmp = strings.ReplaceAll(tmp, "{{GeneratedFieldName}}", "ControlFlowsWhoseNodeIsExpanded")
			tmp = strings.ReplaceAll(tmp, "{{GeneratedFieldNameValue}}", _controlflow.GongGetIdentifier(stage))
			sb.WriteString(tmp)
		}
		res = sb.String()
	case "ControlFlow_Shapes":
		var sb strings.Builder
		for _, _controlflowshape := range diagramstructure.ControlFlow_Shapes {
			tmp := SliceOfPointersFieldInitStatement
			tmp = strings.ReplaceAll(tmp, "{{Identifier}}", diagramstructure.GongGetIdentifier(stage))
			tmp = strings.ReplaceAll(tmp, "{{GeneratedFieldName}}", "ControlFlow_Shapes")
			tmp = strings.ReplaceAll(tmp, "{{GeneratedFieldNameValue}}", _controlflowshape.GongGetIdentifier(stage))
			sb.WriteString(tmp)
		}
		res = sb.String()
	case "DataFlowsWhoseNodeIsExpanded":
		var sb strings.Builder
		for _, _dataflow := range diagramstructure.DataFlowsWhoseNodeIsExpanded {
			tmp := SliceOfPointersFieldInitStatement
			tmp = strings.ReplaceAll(tmp, "{{Identifier}}", diagramstructure.GongGetIdentifier(stage))
			tmp = strings.ReplaceAll(tmp, "{{GeneratedFieldName}}", "DataFlowsWhoseNodeIsExpanded")
			tmp = strings.ReplaceAll(tmp, "{{GeneratedFieldNameValue}}", _dataflow.GongGetIdentifier(stage))
			sb.WriteString(tmp)
		}
		res = sb.String()
	case "DataFlow_Shapes":
		var sb strings.Builder
		for _, _dataflowshape := range diagramstructure.DataFlow_Shapes {
			tmp := SliceOfPointersFieldInitStatement
			tmp = strings.ReplaceAll(tmp, "{{Identifier}}", diagramstructure.GongGetIdentifier(stage))
			tmp = strings.ReplaceAll(tmp, "{{GeneratedFieldName}}", "DataFlow_Shapes")
			tmp = strings.ReplaceAll(tmp, "{{GeneratedFieldNameValue}}", _dataflowshape.GongGetIdentifier(stage))
			sb.WriteString(tmp)
		}
		res = sb.String()
	case "DatasWhoseNodeIsExpanded":
		var sb strings.Builder
		for _, _data := range diagramstructure.DatasWhoseNodeIsExpanded {
			tmp := SliceOfPointersFieldInitStatement
			tmp = strings.ReplaceAll(tmp, "{{Identifier}}", diagramstructure.GongGetIdentifier(stage))
			tmp = strings.ReplaceAll(tmp, "{{GeneratedFieldName}}", "DatasWhoseNodeIsExpanded")
			tmp = strings.ReplaceAll(tmp, "{{GeneratedFieldNameValue}}", _data.GongGetIdentifier(stage))
			sb.WriteString(tmp)
		}
		res = sb.String()
	case "Data_Shapes":
		var sb strings.Builder
		for _, _datashape := range diagramstructure.Data_Shapes {
			tmp := SliceOfPointersFieldInitStatement
			tmp = strings.ReplaceAll(tmp, "{{Identifier}}", diagramstructure.GongGetIdentifier(stage))
			tmp = strings.ReplaceAll(tmp, "{{GeneratedFieldName}}", "Data_Shapes")
			tmp = strings.ReplaceAll(tmp, "{{GeneratedFieldNameValue}}", _datashape.GongGetIdentifier(stage))
			sb.WriteString(tmp)
		}
		res = sb.String()
	case "DataFlowsWhoseDataNodeIsExpanded":
		var sb strings.Builder
		for _, _dataflow := range diagramstructure.DataFlowsWhoseDataNodeIsExpanded {
			tmp := SliceOfPointersFieldInitStatement
			tmp = strings.ReplaceAll(tmp, "{{Identifier}}", diagramstructure.GongGetIdentifier(stage))
			tmp = strings.ReplaceAll(tmp, "{{GeneratedFieldName}}", "DataFlowsWhoseDataNodeIsExpanded")
			tmp = strings.ReplaceAll(tmp, "{{GeneratedFieldNameValue}}", _dataflow.GongGetIdentifier(stage))
			sb.WriteString(tmp)
		}
		res = sb.String()
	case "AllocatedResourcesWhoseNodeIsExpanded":
		var sb strings.Builder
		for _, _resource := range diagramstructure.AllocatedResourcesWhoseNodeIsExpanded {
			tmp := SliceOfPointersFieldInitStatement
			tmp = strings.ReplaceAll(tmp, "{{Identifier}}", diagramstructure.GongGetIdentifier(stage))
			tmp = strings.ReplaceAll(tmp, "{{GeneratedFieldName}}", "AllocatedResourcesWhoseNodeIsExpanded")
			tmp = strings.ReplaceAll(tmp, "{{GeneratedFieldNameValue}}", _resource.GongGetIdentifier(stage))
			sb.WriteString(tmp)
		}
		res = sb.String()
	case "AllocatedResourceShapes":
		var sb strings.Builder
		for _, _allocatedresourceshape := range diagramstructure.AllocatedResourceShapes {
			tmp := SliceOfPointersFieldInitStatement
			tmp = strings.ReplaceAll(tmp, "{{Identifier}}", diagramstructure.GongGetIdentifier(stage))
			tmp = strings.ReplaceAll(tmp, "{{GeneratedFieldName}}", "AllocatedResourceShapes")
			tmp = strings.ReplaceAll(tmp, "{{GeneratedFieldNameValue}}", _allocatedresourceshape.GongGetIdentifier(stage))
			sb.WriteString(tmp)
		}
		res = sb.String()
	case "AllocatedSystemesWhoseNodeIsExpanded":
		var sb strings.Builder
		for _, _system := range diagramstructure.AllocatedSystemesWhoseNodeIsExpanded {
			tmp := SliceOfPointersFieldInitStatement
			tmp = strings.ReplaceAll(tmp, "{{Identifier}}", diagramstructure.GongGetIdentifier(stage))
			tmp = strings.ReplaceAll(tmp, "{{GeneratedFieldName}}", "AllocatedSystemesWhoseNodeIsExpanded")
			tmp = strings.ReplaceAll(tmp, "{{GeneratedFieldNameValue}}", _system.GongGetIdentifier(stage))
			sb.WriteString(tmp)
		}
		res = sb.String()
	case "AllocatedSystemShapes":
		var sb strings.Builder
		for _, _allocatedsystemshape := range diagramstructure.AllocatedSystemShapes {
			tmp := SliceOfPointersFieldInitStatement
			tmp = strings.ReplaceAll(tmp, "{{Identifier}}", diagramstructure.GongGetIdentifier(stage))
			tmp = strings.ReplaceAll(tmp, "{{GeneratedFieldName}}", "AllocatedSystemShapes")
			tmp = strings.ReplaceAll(tmp, "{{GeneratedFieldNameValue}}", _allocatedsystemshape.GongGetIdentifier(stage))
			sb.WriteString(tmp)
		}
		res = sb.String()
	case "Note_Shapes":
		var sb strings.Builder
		for _, _noteshape := range diagramstructure.Note_Shapes {
			tmp := SliceOfPointersFieldInitStatement
			tmp = strings.ReplaceAll(tmp, "{{Identifier}}", diagramstructure.GongGetIdentifier(stage))
			tmp = strings.ReplaceAll(tmp, "{{GeneratedFieldName}}", "Note_Shapes")
			tmp = strings.ReplaceAll(tmp, "{{GeneratedFieldNameValue}}", _noteshape.GongGetIdentifier(stage))
			sb.WriteString(tmp)
		}
		res = sb.String()
	case "NotesWhoseNodeIsExpanded":
		var sb strings.Builder
		for _, _note := range diagramstructure.NotesWhoseNodeIsExpanded {
			tmp := SliceOfPointersFieldInitStatement
			tmp = strings.ReplaceAll(tmp, "{{Identifier}}", diagramstructure.GongGetIdentifier(stage))
			tmp = strings.ReplaceAll(tmp, "{{GeneratedFieldName}}", "NotesWhoseNodeIsExpanded")
			tmp = strings.ReplaceAll(tmp, "{{GeneratedFieldNameValue}}", _note.GongGetIdentifier(stage))
			sb.WriteString(tmp)
		}
		res = sb.String()
	case "NotePortShapes":
		var sb strings.Builder
		for _, _noteportshape := range diagramstructure.NotePortShapes {
			tmp := SliceOfPointersFieldInitStatement
			tmp = strings.ReplaceAll(tmp, "{{Identifier}}", diagramstructure.GongGetIdentifier(stage))
			tmp = strings.ReplaceAll(tmp, "{{GeneratedFieldName}}", "NotePortShapes")
			tmp = strings.ReplaceAll(tmp, "{{GeneratedFieldNameValue}}", _noteportshape.GongGetIdentifier(stage))
			sb.WriteString(tmp)
		}
		res = sb.String()
	default:
		log.Panicf("Unknown field %s for Gongstruct DiagramStructure", fieldName)
	}
	return
}

func (externalpartshape *ExternalPartShape) GongMarshallField(stage *Stage, fieldName string) (res string) {

	switch fieldName {
	case "Name":
		res = StringInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", externalpartshape.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "Name")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", ToRawStringLiteral(externalpartshape.Name))
	case "IsExpanded":
		res = NumberInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", externalpartshape.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "IsExpanded")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%t", externalpartshape.IsExpanded))
	case "X":
		res = NumberInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", externalpartshape.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "X")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%f", externalpartshape.X))
	case "Y":
		res = NumberInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", externalpartshape.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "Y")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%f", externalpartshape.Y))
	case "Width":
		res = NumberInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", externalpartshape.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "Width")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%f", externalpartshape.Width))
	case "Height":
		res = NumberInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", externalpartshape.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "Height")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%f", externalpartshape.Height))
	case "IsHidden":
		res = NumberInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", externalpartshape.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "IsHidden")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%t", externalpartshape.IsHidden))
	case "TailHeigth":
		res = NumberInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", externalpartshape.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "TailHeigth")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%f", externalpartshape.TailHeigth))

	case "Part":
		if externalpartshape.Part != nil {
			res = PointerFieldInitStatement
			res = strings.ReplaceAll(res, "{{Identifier}}", externalpartshape.GongGetIdentifier(stage))
			res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "Part")
			res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", externalpartshape.Part.GongGetIdentifier(stage))
		} else {
			// in case of nil pointer, we need to unstage the previous value
			res = PointerFieldInitStatement
			res = strings.ReplaceAll(res, "{{Identifier}}", externalpartshape.GongGetIdentifier(stage))
			res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "Part")
			res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", "nil")
		}
	default:
		log.Panicf("Unknown field %s for Gongstruct ExternalPartShape", fieldName)
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
	case "Description":
		res = StringInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", library.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "Description")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", ToRawStringLiteral(library.Description))
	case "ComputedPrefix":
		res = StringInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", library.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "ComputedPrefix")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", ToRawStringLiteral(library.ComputedPrefix))
	case "IsExpanded":
		res = NumberInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", library.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "IsExpanded")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%t", library.IsExpanded))
	case "LayoutDirection":
		if library.LayoutDirection.ToCodeString() != "" {
			res = NumberInitStatement
			res = strings.ReplaceAll(res, "{{Identifier}}", library.GongGetIdentifier(stage))
			res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "LayoutDirection")
			res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", "models."+library.LayoutDirection.ToCodeString())
		} else {
			// in case of empty enum, we need to unstage the previous value
			res = NumberInitStatement
			res = strings.ReplaceAll(res, "{{Identifier}}", library.GongGetIdentifier(stage))
			res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "LayoutDirection")
			res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", "0")
		}
	case "IsRootLibrary":
		res = NumberInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", library.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "IsRootLibrary")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%t", library.IsRootLibrary))
	case "IsSubLibrariesNodeExpanded":
		res = NumberInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", library.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "IsSubLibrariesNodeExpanded")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%t", library.IsSubLibrariesNodeExpanded))
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
	case "IsSystemesNodeExpanded":
		res = NumberInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", library.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "IsSystemesNodeExpanded")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%t", library.IsSystemesNodeExpanded))
	case "IsDataFlowsNodeExpanded":
		res = NumberInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", library.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "IsDataFlowsNodeExpanded")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%t", library.IsDataFlowsNodeExpanded))
	case "IsDatasNodeExpanded":
		res = NumberInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", library.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "IsDatasNodeExpanded")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%t", library.IsDatasNodeExpanded))
	case "IsResourcesNodeExpanded":
		res = NumberInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", library.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "IsResourcesNodeExpanded")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%t", library.IsResourcesNodeExpanded))
	case "IsNotesNodeExpanded":
		res = NumberInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", library.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "IsNotesNodeExpanded")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%t", library.IsNotesNodeExpanded))
	case "IsExpandedTmp":
		res = NumberInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", library.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "IsExpandedTmp")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%t", library.IsExpandedTmp))

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
	case "SubLibrariesWhoseNodeIsExpanded":
		var sb strings.Builder
		for _, _library := range library.SubLibrariesWhoseNodeIsExpanded {
			tmp := SliceOfPointersFieldInitStatement
			tmp = strings.ReplaceAll(tmp, "{{Identifier}}", library.GongGetIdentifier(stage))
			tmp = strings.ReplaceAll(tmp, "{{GeneratedFieldName}}", "SubLibrariesWhoseNodeIsExpanded")
			tmp = strings.ReplaceAll(tmp, "{{GeneratedFieldNameValue}}", _library.GongGetIdentifier(stage))
			sb.WriteString(tmp)
		}
		res = sb.String()
	case "RootSystemes":
		var sb strings.Builder
		for _, _system := range library.RootSystemes {
			tmp := SliceOfPointersFieldInitStatement
			tmp = strings.ReplaceAll(tmp, "{{Identifier}}", library.GongGetIdentifier(stage))
			tmp = strings.ReplaceAll(tmp, "{{GeneratedFieldName}}", "RootSystemes")
			tmp = strings.ReplaceAll(tmp, "{{GeneratedFieldNameValue}}", _system.GongGetIdentifier(stage))
			sb.WriteString(tmp)
		}
		res = sb.String()
	case "SystemsWhoseNodeIsExpanded":
		var sb strings.Builder
		for _, _system := range library.SystemsWhoseNodeIsExpanded {
			tmp := SliceOfPointersFieldInitStatement
			tmp = strings.ReplaceAll(tmp, "{{Identifier}}", library.GongGetIdentifier(stage))
			tmp = strings.ReplaceAll(tmp, "{{GeneratedFieldName}}", "SystemsWhoseNodeIsExpanded")
			tmp = strings.ReplaceAll(tmp, "{{GeneratedFieldNameValue}}", _system.GongGetIdentifier(stage))
			sb.WriteString(tmp)
		}
		res = sb.String()
	case "RootDataFlows":
		var sb strings.Builder
		for _, _dataflow := range library.RootDataFlows {
			tmp := SliceOfPointersFieldInitStatement
			tmp = strings.ReplaceAll(tmp, "{{Identifier}}", library.GongGetIdentifier(stage))
			tmp = strings.ReplaceAll(tmp, "{{GeneratedFieldName}}", "RootDataFlows")
			tmp = strings.ReplaceAll(tmp, "{{GeneratedFieldNameValue}}", _dataflow.GongGetIdentifier(stage))
			sb.WriteString(tmp)
		}
		res = sb.String()
	case "DataFlowsWhoseNodeIsExpanded":
		var sb strings.Builder
		for _, _dataflow := range library.DataFlowsWhoseNodeIsExpanded {
			tmp := SliceOfPointersFieldInitStatement
			tmp = strings.ReplaceAll(tmp, "{{Identifier}}", library.GongGetIdentifier(stage))
			tmp = strings.ReplaceAll(tmp, "{{GeneratedFieldName}}", "DataFlowsWhoseNodeIsExpanded")
			tmp = strings.ReplaceAll(tmp, "{{GeneratedFieldNameValue}}", _dataflow.GongGetIdentifier(stage))
			sb.WriteString(tmp)
		}
		res = sb.String()
	case "RootDatas":
		var sb strings.Builder
		for _, _data := range library.RootDatas {
			tmp := SliceOfPointersFieldInitStatement
			tmp = strings.ReplaceAll(tmp, "{{Identifier}}", library.GongGetIdentifier(stage))
			tmp = strings.ReplaceAll(tmp, "{{GeneratedFieldName}}", "RootDatas")
			tmp = strings.ReplaceAll(tmp, "{{GeneratedFieldNameValue}}", _data.GongGetIdentifier(stage))
			sb.WriteString(tmp)
		}
		res = sb.String()
	case "DatasWhoseNodeIsExpanded":
		var sb strings.Builder
		for _, _data := range library.DatasWhoseNodeIsExpanded {
			tmp := SliceOfPointersFieldInitStatement
			tmp = strings.ReplaceAll(tmp, "{{Identifier}}", library.GongGetIdentifier(stage))
			tmp = strings.ReplaceAll(tmp, "{{GeneratedFieldName}}", "DatasWhoseNodeIsExpanded")
			tmp = strings.ReplaceAll(tmp, "{{GeneratedFieldNameValue}}", _data.GongGetIdentifier(stage))
			sb.WriteString(tmp)
		}
		res = sb.String()
	case "RootResources":
		var sb strings.Builder
		for _, _resource := range library.RootResources {
			tmp := SliceOfPointersFieldInitStatement
			tmp = strings.ReplaceAll(tmp, "{{Identifier}}", library.GongGetIdentifier(stage))
			tmp = strings.ReplaceAll(tmp, "{{GeneratedFieldName}}", "RootResources")
			tmp = strings.ReplaceAll(tmp, "{{GeneratedFieldNameValue}}", _resource.GongGetIdentifier(stage))
			sb.WriteString(tmp)
		}
		res = sb.String()
	case "ResourcesWhoseNodeIsExpanded":
		var sb strings.Builder
		for _, _resource := range library.ResourcesWhoseNodeIsExpanded {
			tmp := SliceOfPointersFieldInitStatement
			tmp = strings.ReplaceAll(tmp, "{{Identifier}}", library.GongGetIdentifier(stage))
			tmp = strings.ReplaceAll(tmp, "{{GeneratedFieldName}}", "ResourcesWhoseNodeIsExpanded")
			tmp = strings.ReplaceAll(tmp, "{{GeneratedFieldNameValue}}", _resource.GongGetIdentifier(stage))
			sb.WriteString(tmp)
		}
		res = sb.String()
	case "PartsWhoseNodeIsExpanded":
		var sb strings.Builder
		for _, _part := range library.PartsWhoseNodeIsExpanded {
			tmp := SliceOfPointersFieldInitStatement
			tmp = strings.ReplaceAll(tmp, "{{Identifier}}", library.GongGetIdentifier(stage))
			tmp = strings.ReplaceAll(tmp, "{{GeneratedFieldName}}", "PartsWhoseNodeIsExpanded")
			tmp = strings.ReplaceAll(tmp, "{{GeneratedFieldNameValue}}", _part.GongGetIdentifier(stage))
			sb.WriteString(tmp)
		}
		res = sb.String()
	case "RootNotes":
		var sb strings.Builder
		for _, _note := range library.RootNotes {
			tmp := SliceOfPointersFieldInitStatement
			tmp = strings.ReplaceAll(tmp, "{{Identifier}}", library.GongGetIdentifier(stage))
			tmp = strings.ReplaceAll(tmp, "{{GeneratedFieldName}}", "RootNotes")
			tmp = strings.ReplaceAll(tmp, "{{GeneratedFieldNameValue}}", _note.GongGetIdentifier(stage))
			sb.WriteString(tmp)
		}
		res = sb.String()
	case "NotesWhoseNodeIsExpanded":
		var sb strings.Builder
		for _, _note := range library.NotesWhoseNodeIsExpanded {
			tmp := SliceOfPointersFieldInitStatement
			tmp = strings.ReplaceAll(tmp, "{{Identifier}}", library.GongGetIdentifier(stage))
			tmp = strings.ReplaceAll(tmp, "{{GeneratedFieldName}}", "NotesWhoseNodeIsExpanded")
			tmp = strings.ReplaceAll(tmp, "{{GeneratedFieldNameValue}}", _note.GongGetIdentifier(stage))
			sb.WriteString(tmp)
		}
		res = sb.String()
	default:
		log.Panicf("Unknown field %s for Gongstruct Library", fieldName)
	}
	return
}

func (note *Note) GongMarshallField(stage *Stage, fieldName string) (res string) {

	switch fieldName {
	case "Name":
		res = StringInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", note.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "Name")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", ToRawStringLiteral(note.Name))
	case "Description":
		res = StringInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", note.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "Description")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", ToRawStringLiteral(note.Description))
	case "ComputedPrefix":
		res = StringInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", note.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "ComputedPrefix")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", ToRawStringLiteral(note.ComputedPrefix))
	case "IsExpanded":
		res = NumberInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", note.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "IsExpanded")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%t", note.IsExpanded))
	case "LayoutDirection":
		if note.LayoutDirection.ToCodeString() != "" {
			res = NumberInitStatement
			res = strings.ReplaceAll(res, "{{Identifier}}", note.GongGetIdentifier(stage))
			res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "LayoutDirection")
			res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", "models."+note.LayoutDirection.ToCodeString())
		} else {
			// in case of empty enum, we need to unstage the previous value
			res = NumberInitStatement
			res = strings.ReplaceAll(res, "{{Identifier}}", note.GongGetIdentifier(stage))
			res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "LayoutDirection")
			res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", "0")
		}
	case "IsPortsNodeExpanded":
		res = NumberInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", note.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "IsPortsNodeExpanded")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%t", note.IsPortsNodeExpanded))

	case "Ports":
		var sb strings.Builder
		for _, _port := range note.Ports {
			tmp := SliceOfPointersFieldInitStatement
			tmp = strings.ReplaceAll(tmp, "{{Identifier}}", note.GongGetIdentifier(stage))
			tmp = strings.ReplaceAll(tmp, "{{GeneratedFieldName}}", "Ports")
			tmp = strings.ReplaceAll(tmp, "{{GeneratedFieldNameValue}}", _port.GongGetIdentifier(stage))
			sb.WriteString(tmp)
		}
		res = sb.String()
	default:
		log.Panicf("Unknown field %s for Gongstruct Note", fieldName)
	}
	return
}

func (noteportshape *NotePortShape) GongMarshallField(stage *Stage, fieldName string) (res string) {

	switch fieldName {
	case "Name":
		res = StringInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", noteportshape.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "Name")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", ToRawStringLiteral(noteportshape.Name))
	case "StartRatio":
		res = NumberInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", noteportshape.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "StartRatio")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%f", noteportshape.StartRatio))
	case "EndRatio":
		res = NumberInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", noteportshape.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "EndRatio")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%f", noteportshape.EndRatio))
	case "StartOrientation":
		if noteportshape.StartOrientation.ToCodeString() != "" {
			res = StringEnumInitStatement
			res = strings.ReplaceAll(res, "{{Identifier}}", noteportshape.GongGetIdentifier(stage))
			res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "StartOrientation")
			res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", "models."+noteportshape.StartOrientation.ToCodeString())
		} else {
			// in case of empty enum, we need to unstage the previous value
			res = StringEnumInitStatement
			res = strings.ReplaceAll(res, "{{Identifier}}", noteportshape.GongGetIdentifier(stage))
			res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "StartOrientation")
			res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", "\"\"")
		}
	case "EndOrientation":
		if noteportshape.EndOrientation.ToCodeString() != "" {
			res = StringEnumInitStatement
			res = strings.ReplaceAll(res, "{{Identifier}}", noteportshape.GongGetIdentifier(stage))
			res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "EndOrientation")
			res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", "models."+noteportshape.EndOrientation.ToCodeString())
		} else {
			// in case of empty enum, we need to unstage the previous value
			res = StringEnumInitStatement
			res = strings.ReplaceAll(res, "{{Identifier}}", noteportshape.GongGetIdentifier(stage))
			res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "EndOrientation")
			res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", "\"\"")
		}
	case "CornerOffsetRatio":
		res = NumberInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", noteportshape.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "CornerOffsetRatio")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%f", noteportshape.CornerOffsetRatio))
	case "IsHidden":
		res = NumberInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", noteportshape.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "IsHidden")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%t", noteportshape.IsHidden))

	case "Note":
		if noteportshape.Note != nil {
			res = PointerFieldInitStatement
			res = strings.ReplaceAll(res, "{{Identifier}}", noteportshape.GongGetIdentifier(stage))
			res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "Note")
			res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", noteportshape.Note.GongGetIdentifier(stage))
		} else {
			// in case of nil pointer, we need to unstage the previous value
			res = PointerFieldInitStatement
			res = strings.ReplaceAll(res, "{{Identifier}}", noteportshape.GongGetIdentifier(stage))
			res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "Note")
			res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", "nil")
		}
	case "Port":
		if noteportshape.Port != nil {
			res = PointerFieldInitStatement
			res = strings.ReplaceAll(res, "{{Identifier}}", noteportshape.GongGetIdentifier(stage))
			res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "Port")
			res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", noteportshape.Port.GongGetIdentifier(stage))
		} else {
			// in case of nil pointer, we need to unstage the previous value
			res = PointerFieldInitStatement
			res = strings.ReplaceAll(res, "{{Identifier}}", noteportshape.GongGetIdentifier(stage))
			res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "Port")
			res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", "nil")
		}
	default:
		log.Panicf("Unknown field %s for Gongstruct NotePortShape", fieldName)
	}
	return
}

func (noteshape *NoteShape) GongMarshallField(stage *Stage, fieldName string) (res string) {

	switch fieldName {
	case "Name":
		res = StringInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", noteshape.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "Name")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", ToRawStringLiteral(noteshape.Name))
	case "X":
		res = NumberInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", noteshape.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "X")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%f", noteshape.X))
	case "Y":
		res = NumberInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", noteshape.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "Y")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%f", noteshape.Y))
	case "Width":
		res = NumberInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", noteshape.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "Width")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%f", noteshape.Width))
	case "Height":
		res = NumberInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", noteshape.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "Height")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%f", noteshape.Height))
	case "IsHidden":
		res = NumberInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", noteshape.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "IsHidden")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%t", noteshape.IsHidden))

	case "Note":
		if noteshape.Note != nil {
			res = PointerFieldInitStatement
			res = strings.ReplaceAll(res, "{{Identifier}}", noteshape.GongGetIdentifier(stage))
			res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "Note")
			res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", noteshape.Note.GongGetIdentifier(stage))
		} else {
			// in case of nil pointer, we need to unstage the previous value
			res = PointerFieldInitStatement
			res = strings.ReplaceAll(res, "{{Identifier}}", noteshape.GongGetIdentifier(stage))
			res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "Note")
			res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", "nil")
		}
	default:
		log.Panicf("Unknown field %s for Gongstruct NoteShape", fieldName)
	}
	return
}

func (part *Part) GongMarshallField(stage *Stage, fieldName string) (res string) {

	switch fieldName {
	case "Name":
		res = StringInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", part.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "Name")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", ToRawStringLiteral(part.Name))
	case "IsSystemResource":
		res = NumberInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", part.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "IsSystemResource")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%t", part.IsSystemResource))
	case "Description":
		res = StringInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", part.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "Description")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", ToRawStringLiteral(part.Description))
	case "IsResourcesNodeExpanded":
		res = NumberInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", part.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "IsResourcesNodeExpanded")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%t", part.IsResourcesNodeExpanded))
	case "IsSystemesNodeExpanded":
		res = NumberInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", part.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "IsSystemesNodeExpanded")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%t", part.IsSystemesNodeExpanded))
	case "ComputedPrefix":
		res = StringInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", part.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "ComputedPrefix")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", ToRawStringLiteral(part.ComputedPrefix))
	case "IsExpanded":
		res = NumberInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", part.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "IsExpanded")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%t", part.IsExpanded))
	case "LayoutDirection":
		if part.LayoutDirection.ToCodeString() != "" {
			res = NumberInitStatement
			res = strings.ReplaceAll(res, "{{Identifier}}", part.GongGetIdentifier(stage))
			res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "LayoutDirection")
			res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", "models."+part.LayoutDirection.ToCodeString())
		} else {
			// in case of empty enum, we need to unstage the previous value
			res = NumberInitStatement
			res = strings.ReplaceAll(res, "{{Identifier}}", part.GongGetIdentifier(stage))
			res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "LayoutDirection")
			res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", "0")
		}
	case "IsPortsNodeExpanded":
		res = NumberInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", part.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "IsPortsNodeExpanded")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%t", part.IsPortsNodeExpanded))
	case "IsControlFlowsNodeExpanded":
		res = NumberInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", part.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "IsControlFlowsNodeExpanded")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%t", part.IsControlFlowsNodeExpanded))
	case "IsDataFlowsNodeExpanded":
		res = NumberInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", part.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "IsDataFlowsNodeExpanded")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%t", part.IsDataFlowsNodeExpanded))

	case "Resources":
		var sb strings.Builder
		for _, _resource := range part.Resources {
			tmp := SliceOfPointersFieldInitStatement
			tmp = strings.ReplaceAll(tmp, "{{Identifier}}", part.GongGetIdentifier(stage))
			tmp = strings.ReplaceAll(tmp, "{{GeneratedFieldName}}", "Resources")
			tmp = strings.ReplaceAll(tmp, "{{GeneratedFieldNameValue}}", _resource.GongGetIdentifier(stage))
			sb.WriteString(tmp)
		}
		res = sb.String()
	case "Systemes":
		var sb strings.Builder
		for _, _system := range part.Systemes {
			tmp := SliceOfPointersFieldInitStatement
			tmp = strings.ReplaceAll(tmp, "{{Identifier}}", part.GongGetIdentifier(stage))
			tmp = strings.ReplaceAll(tmp, "{{GeneratedFieldName}}", "Systemes")
			tmp = strings.ReplaceAll(tmp, "{{GeneratedFieldNameValue}}", _system.GongGetIdentifier(stage))
			sb.WriteString(tmp)
		}
		res = sb.String()
	case "Ports":
		var sb strings.Builder
		for _, _port := range part.Ports {
			tmp := SliceOfPointersFieldInitStatement
			tmp = strings.ReplaceAll(tmp, "{{Identifier}}", part.GongGetIdentifier(stage))
			tmp = strings.ReplaceAll(tmp, "{{GeneratedFieldName}}", "Ports")
			tmp = strings.ReplaceAll(tmp, "{{GeneratedFieldNameValue}}", _port.GongGetIdentifier(stage))
			sb.WriteString(tmp)
		}
		res = sb.String()
	case "ControlFlows":
		var sb strings.Builder
		for _, _controlflow := range part.ControlFlows {
			tmp := SliceOfPointersFieldInitStatement
			tmp = strings.ReplaceAll(tmp, "{{Identifier}}", part.GongGetIdentifier(stage))
			tmp = strings.ReplaceAll(tmp, "{{GeneratedFieldName}}", "ControlFlows")
			tmp = strings.ReplaceAll(tmp, "{{GeneratedFieldNameValue}}", _controlflow.GongGetIdentifier(stage))
			sb.WriteString(tmp)
		}
		res = sb.String()
	case "PortWhoseOutControlFlowsNodeIsExpanded":
		var sb strings.Builder
		for _, _port := range part.PortWhoseOutControlFlowsNodeIsExpanded {
			tmp := SliceOfPointersFieldInitStatement
			tmp = strings.ReplaceAll(tmp, "{{Identifier}}", part.GongGetIdentifier(stage))
			tmp = strings.ReplaceAll(tmp, "{{GeneratedFieldName}}", "PortWhoseOutControlFlowsNodeIsExpanded")
			tmp = strings.ReplaceAll(tmp, "{{GeneratedFieldNameValue}}", _port.GongGetIdentifier(stage))
			sb.WriteString(tmp)
		}
		res = sb.String()
	case "PortWhoseInControlFlowsNodeIsExpanded":
		var sb strings.Builder
		for _, _port := range part.PortWhoseInControlFlowsNodeIsExpanded {
			tmp := SliceOfPointersFieldInitStatement
			tmp = strings.ReplaceAll(tmp, "{{Identifier}}", part.GongGetIdentifier(stage))
			tmp = strings.ReplaceAll(tmp, "{{GeneratedFieldName}}", "PortWhoseInControlFlowsNodeIsExpanded")
			tmp = strings.ReplaceAll(tmp, "{{GeneratedFieldNameValue}}", _port.GongGetIdentifier(stage))
			sb.WriteString(tmp)
		}
		res = sb.String()
	case "PortWhoseOutDataFlowsNodeIsExpanded":
		var sb strings.Builder
		for _, _port := range part.PortWhoseOutDataFlowsNodeIsExpanded {
			tmp := SliceOfPointersFieldInitStatement
			tmp = strings.ReplaceAll(tmp, "{{Identifier}}", part.GongGetIdentifier(stage))
			tmp = strings.ReplaceAll(tmp, "{{GeneratedFieldName}}", "PortWhoseOutDataFlowsNodeIsExpanded")
			tmp = strings.ReplaceAll(tmp, "{{GeneratedFieldNameValue}}", _port.GongGetIdentifier(stage))
			sb.WriteString(tmp)
		}
		res = sb.String()
	case "PortWhoseInDataFlowsNodeIsExpanded":
		var sb strings.Builder
		for _, _port := range part.PortWhoseInDataFlowsNodeIsExpanded {
			tmp := SliceOfPointersFieldInitStatement
			tmp = strings.ReplaceAll(tmp, "{{Identifier}}", part.GongGetIdentifier(stage))
			tmp = strings.ReplaceAll(tmp, "{{GeneratedFieldName}}", "PortWhoseInDataFlowsNodeIsExpanded")
			tmp = strings.ReplaceAll(tmp, "{{GeneratedFieldNameValue}}", _port.GongGetIdentifier(stage))
			sb.WriteString(tmp)
		}
		res = sb.String()
	default:
		log.Panicf("Unknown field %s for Gongstruct Part", fieldName)
	}
	return
}

func (partshape *PartShape) GongMarshallField(stage *Stage, fieldName string) (res string) {

	switch fieldName {
	case "Name":
		res = StringInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", partshape.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "Name")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", ToRawStringLiteral(partshape.Name))
	case "IsExpanded":
		res = NumberInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", partshape.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "IsExpanded")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%t", partshape.IsExpanded))
	case "X":
		res = NumberInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", partshape.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "X")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%f", partshape.X))
	case "Y":
		res = NumberInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", partshape.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "Y")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%f", partshape.Y))
	case "Width":
		res = NumberInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", partshape.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "Width")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%f", partshape.Width))
	case "Height":
		res = NumberInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", partshape.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "Height")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%f", partshape.Height))
	case "IsHidden":
		res = NumberInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", partshape.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "IsHidden")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%t", partshape.IsHidden))
	case "WidthWeight":
		res = NumberInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", partshape.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "WidthWeight")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%f", partshape.WidthWeight))

	case "Part":
		if partshape.Part != nil {
			res = PointerFieldInitStatement
			res = strings.ReplaceAll(res, "{{Identifier}}", partshape.GongGetIdentifier(stage))
			res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "Part")
			res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", partshape.Part.GongGetIdentifier(stage))
		} else {
			// in case of nil pointer, we need to unstage the previous value
			res = PointerFieldInitStatement
			res = strings.ReplaceAll(res, "{{Identifier}}", partshape.GongGetIdentifier(stage))
			res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "Part")
			res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", "nil")
		}
	default:
		log.Panicf("Unknown field %s for Gongstruct PartShape", fieldName)
	}
	return
}

func (port *Port) GongMarshallField(stage *Stage, fieldName string) (res string) {

	switch fieldName {
	case "Name":
		res = StringInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", port.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "Name")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", ToRawStringLiteral(port.Name))
	case "Description":
		res = StringInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", port.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "Description")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", ToRawStringLiteral(port.Description))
	case "ComputedPrefix":
		res = StringInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", port.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "ComputedPrefix")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", ToRawStringLiteral(port.ComputedPrefix))
	case "IsExpanded":
		res = NumberInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", port.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "IsExpanded")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%t", port.IsExpanded))
	case "LayoutDirection":
		if port.LayoutDirection.ToCodeString() != "" {
			res = NumberInitStatement
			res = strings.ReplaceAll(res, "{{Identifier}}", port.GongGetIdentifier(stage))
			res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "LayoutDirection")
			res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", "models."+port.LayoutDirection.ToCodeString())
		} else {
			// in case of empty enum, we need to unstage the previous value
			res = NumberInitStatement
			res = strings.ReplaceAll(res, "{{Identifier}}", port.GongGetIdentifier(stage))
			res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "LayoutDirection")
			res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", "0")
		}
	case "IsStartPort":
		res = NumberInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", port.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "IsStartPort")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%t", port.IsStartPort))
	case "IsEndPort":
		res = NumberInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", port.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "IsEndPort")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%t", port.IsEndPort))
	case "IsPortNameNotSystemName":
		res = NumberInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", port.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "IsPortNameNotSystemName")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%t", port.IsPortNameNotSystemName))

	case "Type":
		if port.Type != nil {
			res = PointerFieldInitStatement
			res = strings.ReplaceAll(res, "{{Identifier}}", port.GongGetIdentifier(stage))
			res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "Type")
			res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", port.Type.GongGetIdentifier(stage))
		} else {
			// in case of nil pointer, we need to unstage the previous value
			res = PointerFieldInitStatement
			res = strings.ReplaceAll(res, "{{Identifier}}", port.GongGetIdentifier(stage))
			res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "Type")
			res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", "nil")
		}
	default:
		log.Panicf("Unknown field %s for Gongstruct Port", fieldName)
	}
	return
}

func (portshape *PortShape) GongMarshallField(stage *Stage, fieldName string) (res string) {

	switch fieldName {
	case "Name":
		res = StringInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", portshape.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "Name")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", ToRawStringLiteral(portshape.Name))
	case "IsExpanded":
		res = NumberInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", portshape.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "IsExpanded")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%t", portshape.IsExpanded))
	case "X":
		res = NumberInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", portshape.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "X")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%f", portshape.X))
	case "Y":
		res = NumberInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", portshape.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "Y")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%f", portshape.Y))
	case "Width":
		res = NumberInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", portshape.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "Width")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%f", portshape.Width))
	case "Height":
		res = NumberInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", portshape.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "Height")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%f", portshape.Height))
	case "IsHidden":
		res = NumberInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", portshape.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "IsHidden")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%t", portshape.IsHidden))

	case "Port":
		if portshape.Port != nil {
			res = PointerFieldInitStatement
			res = strings.ReplaceAll(res, "{{Identifier}}", portshape.GongGetIdentifier(stage))
			res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "Port")
			res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", portshape.Port.GongGetIdentifier(stage))
		} else {
			// in case of nil pointer, we need to unstage the previous value
			res = PointerFieldInitStatement
			res = strings.ReplaceAll(res, "{{Identifier}}", portshape.GongGetIdentifier(stage))
			res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "Port")
			res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", "nil")
		}
	default:
		log.Panicf("Unknown field %s for Gongstruct PortShape", fieldName)
	}
	return
}

func (resource *Resource) GongMarshallField(stage *Stage, fieldName string) (res string) {

	switch fieldName {
	case "Name":
		res = StringInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", resource.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "Name")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", ToRawStringLiteral(resource.Name))
	case "Acronym":
		res = StringInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", resource.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "Acronym")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", ToRawStringLiteral(resource.Acronym))
	case "Description":
		res = StringInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", resource.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "Description")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", ToRawStringLiteral(resource.Description))
	case "ComputedPrefix":
		res = StringInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", resource.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "ComputedPrefix")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", ToRawStringLiteral(resource.ComputedPrefix))
	case "IsExpanded":
		res = NumberInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", resource.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "IsExpanded")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%t", resource.IsExpanded))
	case "LayoutDirection":
		if resource.LayoutDirection.ToCodeString() != "" {
			res = NumberInitStatement
			res = strings.ReplaceAll(res, "{{Identifier}}", resource.GongGetIdentifier(stage))
			res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "LayoutDirection")
			res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", "models."+resource.LayoutDirection.ToCodeString())
		} else {
			// in case of empty enum, we need to unstage the previous value
			res = NumberInitStatement
			res = strings.ReplaceAll(res, "{{Identifier}}", resource.GongGetIdentifier(stage))
			res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "LayoutDirection")
			res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", "0")
		}
	case "SVG_Path":
		res = StringInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", resource.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "SVG_Path")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", ToRawStringLiteral(resource.SVG_Path))
	case "InverseAppliedScaling":
		res = NumberInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", resource.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "InverseAppliedScaling")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%f", resource.InverseAppliedScaling))

	default:
		log.Panicf("Unknown field %s for Gongstruct Resource", fieldName)
	}
	return
}

func (system *System) GongMarshallField(stage *Stage, fieldName string) (res string) {

	switch fieldName {
	case "Name":
		res = StringInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", system.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "Name")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", ToRawStringLiteral(system.Name))
	case "Description":
		res = StringInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", system.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "Description")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", ToRawStringLiteral(system.Description))
	case "ComputedPrefix":
		res = StringInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", system.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "ComputedPrefix")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", ToRawStringLiteral(system.ComputedPrefix))
	case "IsExpanded":
		res = NumberInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", system.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "IsExpanded")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%t", system.IsExpanded))
	case "LayoutDirection":
		if system.LayoutDirection.ToCodeString() != "" {
			res = NumberInitStatement
			res = strings.ReplaceAll(res, "{{Identifier}}", system.GongGetIdentifier(stage))
			res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "LayoutDirection")
			res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", "models."+system.LayoutDirection.ToCodeString())
		} else {
			// in case of empty enum, we need to unstage the previous value
			res = NumberInitStatement
			res = strings.ReplaceAll(res, "{{Identifier}}", system.GongGetIdentifier(stage))
			res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "LayoutDirection")
			res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", "0")
		}
	case "SVG_Path":
		res = StringInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", system.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "SVG_Path")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", ToRawStringLiteral(system.SVG_Path))
	case "InverseAppliedScaling":
		res = NumberInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", system.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "InverseAppliedScaling")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%f", system.InverseAppliedScaling))
	case "IsSubSystemNodeExpanded":
		res = NumberInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", system.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "IsSubSystemNodeExpanded")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%t", system.IsSubSystemNodeExpanded))
	case "IsDataFlowsNodeExpanded":
		res = NumberInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", system.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "IsDataFlowsNodeExpanded")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%t", system.IsDataFlowsNodeExpanded))

	case "DiagramStructures":
		var sb strings.Builder
		for _, _diagramstructure := range system.DiagramStructures {
			tmp := SliceOfPointersFieldInitStatement
			tmp = strings.ReplaceAll(tmp, "{{Identifier}}", system.GongGetIdentifier(stage))
			tmp = strings.ReplaceAll(tmp, "{{GeneratedFieldName}}", "DiagramStructures")
			tmp = strings.ReplaceAll(tmp, "{{GeneratedFieldNameValue}}", _diagramstructure.GongGetIdentifier(stage))
			sb.WriteString(tmp)
		}
		res = sb.String()
	case "DiagramStructureWhoseNodeIsExpanded":
		var sb strings.Builder
		for _, _diagramstructure := range system.DiagramStructureWhoseNodeIsExpanded {
			tmp := SliceOfPointersFieldInitStatement
			tmp = strings.ReplaceAll(tmp, "{{Identifier}}", system.GongGetIdentifier(stage))
			tmp = strings.ReplaceAll(tmp, "{{GeneratedFieldName}}", "DiagramStructureWhoseNodeIsExpanded")
			tmp = strings.ReplaceAll(tmp, "{{GeneratedFieldNameValue}}", _diagramstructure.GongGetIdentifier(stage))
			sb.WriteString(tmp)
		}
		res = sb.String()
	case "SubSystemes":
		var sb strings.Builder
		for _, _system := range system.SubSystemes {
			tmp := SliceOfPointersFieldInitStatement
			tmp = strings.ReplaceAll(tmp, "{{Identifier}}", system.GongGetIdentifier(stage))
			tmp = strings.ReplaceAll(tmp, "{{GeneratedFieldName}}", "SubSystemes")
			tmp = strings.ReplaceAll(tmp, "{{GeneratedFieldNameValue}}", _system.GongGetIdentifier(stage))
			sb.WriteString(tmp)
		}
		res = sb.String()
	case "Parts":
		var sb strings.Builder
		for _, _part := range system.Parts {
			tmp := SliceOfPointersFieldInitStatement
			tmp = strings.ReplaceAll(tmp, "{{Identifier}}", system.GongGetIdentifier(stage))
			tmp = strings.ReplaceAll(tmp, "{{GeneratedFieldName}}", "Parts")
			tmp = strings.ReplaceAll(tmp, "{{GeneratedFieldNameValue}}", _part.GongGetIdentifier(stage))
			sb.WriteString(tmp)
		}
		res = sb.String()
	case "PartWhoseNodeIsExpanded":
		var sb strings.Builder
		for _, _part := range system.PartWhoseNodeIsExpanded {
			tmp := SliceOfPointersFieldInitStatement
			tmp = strings.ReplaceAll(tmp, "{{Identifier}}", system.GongGetIdentifier(stage))
			tmp = strings.ReplaceAll(tmp, "{{GeneratedFieldName}}", "PartWhoseNodeIsExpanded")
			tmp = strings.ReplaceAll(tmp, "{{GeneratedFieldNameValue}}", _part.GongGetIdentifier(stage))
			sb.WriteString(tmp)
		}
		res = sb.String()
	case "DataFlows":
		var sb strings.Builder
		for _, _dataflow := range system.DataFlows {
			tmp := SliceOfPointersFieldInitStatement
			tmp = strings.ReplaceAll(tmp, "{{Identifier}}", system.GongGetIdentifier(stage))
			tmp = strings.ReplaceAll(tmp, "{{GeneratedFieldName}}", "DataFlows")
			tmp = strings.ReplaceAll(tmp, "{{GeneratedFieldNameValue}}", _dataflow.GongGetIdentifier(stage))
			sb.WriteString(tmp)
		}
		res = sb.String()
	case "ExternalParts":
		var sb strings.Builder
		for _, _part := range system.ExternalParts {
			tmp := SliceOfPointersFieldInitStatement
			tmp = strings.ReplaceAll(tmp, "{{Identifier}}", system.GongGetIdentifier(stage))
			tmp = strings.ReplaceAll(tmp, "{{GeneratedFieldName}}", "ExternalParts")
			tmp = strings.ReplaceAll(tmp, "{{GeneratedFieldNameValue}}", _part.GongGetIdentifier(stage))
			sb.WriteString(tmp)
		}
		res = sb.String()
	case "ExternalPartWhoseNodeIsExpanded":
		var sb strings.Builder
		for _, _part := range system.ExternalPartWhoseNodeIsExpanded {
			tmp := SliceOfPointersFieldInitStatement
			tmp = strings.ReplaceAll(tmp, "{{Identifier}}", system.GongGetIdentifier(stage))
			tmp = strings.ReplaceAll(tmp, "{{GeneratedFieldName}}", "ExternalPartWhoseNodeIsExpanded")
			tmp = strings.ReplaceAll(tmp, "{{GeneratedFieldNameValue}}", _part.GongGetIdentifier(stage))
			sb.WriteString(tmp)
		}
		res = sb.String()
	default:
		log.Panicf("Unknown field %s for Gongstruct System", fieldName)
	}
	return
}

func (systemshape *SystemShape) GongMarshallField(stage *Stage, fieldName string) (res string) {

	switch fieldName {
	case "Name":
		res = StringInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", systemshape.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "Name")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", ToRawStringLiteral(systemshape.Name))
	case "IsExpanded":
		res = NumberInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", systemshape.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "IsExpanded")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%t", systemshape.IsExpanded))
	case "X":
		res = NumberInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", systemshape.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "X")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%f", systemshape.X))
	case "Y":
		res = NumberInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", systemshape.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "Y")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%f", systemshape.Y))
	case "Width":
		res = NumberInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", systemshape.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "Width")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%f", systemshape.Width))
	case "Height":
		res = NumberInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", systemshape.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "Height")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%f", systemshape.Height))
	case "IsHidden":
		res = NumberInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", systemshape.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "IsHidden")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%t", systemshape.IsHidden))

	case "System":
		if systemshape.System != nil {
			res = PointerFieldInitStatement
			res = strings.ReplaceAll(res, "{{Identifier}}", systemshape.GongGetIdentifier(stage))
			res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "System")
			res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", systemshape.System.GongGetIdentifier(stage))
		} else {
			// in case of nil pointer, we need to unstage the previous value
			res = PointerFieldInitStatement
			res = strings.ReplaceAll(res, "{{Identifier}}", systemshape.GongGetIdentifier(stage))
			res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "System")
			res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", "nil")
		}
	default:
		log.Panicf("Unknown field %s for Gongstruct SystemShape", fieldName)
	}
	return
}

// insertion point for marshall all fields methods
func (allocatedresourceshape *AllocatedResourceShape) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(allocatedresourceshape.GongMarshallField(stage, "Name"))
		pointersInitializesStatements.WriteString(allocatedresourceshape.GongMarshallField(stage, "Part"))
		pointersInitializesStatements.WriteString(allocatedresourceshape.GongMarshallField(stage, "Resource"))
	}
	initRes = initializerStatements.String()
	ptrRes = pointersInitializesStatements.String()
	return
}
func (allocatedsystemshape *AllocatedSystemShape) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(allocatedsystemshape.GongMarshallField(stage, "Name"))
		pointersInitializesStatements.WriteString(allocatedsystemshape.GongMarshallField(stage, "Part"))
		pointersInitializesStatements.WriteString(allocatedsystemshape.GongMarshallField(stage, "System"))
	}
	initRes = initializerStatements.String()
	ptrRes = pointersInitializesStatements.String()
	return
}
func (controlflow *ControlFlow) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(controlflow.GongMarshallField(stage, "Name"))
		initializerStatements.WriteString(controlflow.GongMarshallField(stage, "Description"))
		initializerStatements.WriteString(controlflow.GongMarshallField(stage, "ComputedPrefix"))
		initializerStatements.WriteString(controlflow.GongMarshallField(stage, "IsExpanded"))
		initializerStatements.WriteString(controlflow.GongMarshallField(stage, "LayoutDirection"))
		pointersInitializesStatements.WriteString(controlflow.GongMarshallField(stage, "Start"))
		pointersInitializesStatements.WriteString(controlflow.GongMarshallField(stage, "End"))
	}
	initRes = initializerStatements.String()
	ptrRes = pointersInitializesStatements.String()
	return
}
func (controlflowshape *ControlFlowShape) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(controlflowshape.GongMarshallField(stage, "Name"))
		pointersInitializesStatements.WriteString(controlflowshape.GongMarshallField(stage, "ControlFlow"))
		initializerStatements.WriteString(controlflowshape.GongMarshallField(stage, "StartRatio"))
		initializerStatements.WriteString(controlflowshape.GongMarshallField(stage, "EndRatio"))
		initializerStatements.WriteString(controlflowshape.GongMarshallField(stage, "StartOrientation"))
		initializerStatements.WriteString(controlflowshape.GongMarshallField(stage, "EndOrientation"))
		initializerStatements.WriteString(controlflowshape.GongMarshallField(stage, "CornerOffsetRatio"))
		initializerStatements.WriteString(controlflowshape.GongMarshallField(stage, "IsHidden"))
	}
	initRes = initializerStatements.String()
	ptrRes = pointersInitializesStatements.String()
	return
}
func (data *Data) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(data.GongMarshallField(stage, "Name"))
		initializerStatements.WriteString(data.GongMarshallField(stage, "Acronym"))
		initializerStatements.WriteString(data.GongMarshallField(stage, "Description"))
		initializerStatements.WriteString(data.GongMarshallField(stage, "ComputedPrefix"))
		initializerStatements.WriteString(data.GongMarshallField(stage, "IsExpanded"))
		initializerStatements.WriteString(data.GongMarshallField(stage, "LayoutDirection"))
		initializerStatements.WriteString(data.GongMarshallField(stage, "SVG_Path"))
		initializerStatements.WriteString(data.GongMarshallField(stage, "InverseAppliedScaling"))
	}
	initRes = initializerStatements.String()
	ptrRes = pointersInitializesStatements.String()
	return
}
func (dataflow *DataFlow) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(dataflow.GongMarshallField(stage, "Name"))
		pointersInitializesStatements.WriteString(dataflow.GongMarshallField(stage, "Datas"))
		initializerStatements.WriteString(dataflow.GongMarshallField(stage, "Description"))
		initializerStatements.WriteString(dataflow.GongMarshallField(stage, "ComputedPrefix"))
		initializerStatements.WriteString(dataflow.GongMarshallField(stage, "IsExpanded"))
		initializerStatements.WriteString(dataflow.GongMarshallField(stage, "LayoutDirection"))
		initializerStatements.WriteString(dataflow.GongMarshallField(stage, "Type"))
		pointersInitializesStatements.WriteString(dataflow.GongMarshallField(stage, "StartPort"))
		pointersInitializesStatements.WriteString(dataflow.GongMarshallField(stage, "EndPort"))
		pointersInitializesStatements.WriteString(dataflow.GongMarshallField(stage, "StartExternalPart"))
		pointersInitializesStatements.WriteString(dataflow.GongMarshallField(stage, "EndExternalPart"))
		initializerStatements.WriteString(dataflow.GongMarshallField(stage, "IsDatasNodeExpanded"))
	}
	initRes = initializerStatements.String()
	ptrRes = pointersInitializesStatements.String()
	return
}
func (dataflowshape *DataFlowShape) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(dataflowshape.GongMarshallField(stage, "Name"))
		pointersInitializesStatements.WriteString(dataflowshape.GongMarshallField(stage, "DataFlow"))
		initializerStatements.WriteString(dataflowshape.GongMarshallField(stage, "StartRatio"))
		initializerStatements.WriteString(dataflowshape.GongMarshallField(stage, "EndRatio"))
		initializerStatements.WriteString(dataflowshape.GongMarshallField(stage, "StartOrientation"))
		initializerStatements.WriteString(dataflowshape.GongMarshallField(stage, "EndOrientation"))
		initializerStatements.WriteString(dataflowshape.GongMarshallField(stage, "CornerOffsetRatio"))
		initializerStatements.WriteString(dataflowshape.GongMarshallField(stage, "IsHidden"))
	}
	initRes = initializerStatements.String()
	ptrRes = pointersInitializesStatements.String()
	return
}
func (datashape *DataShape) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(datashape.GongMarshallField(stage, "Name"))
		pointersInitializesStatements.WriteString(datashape.GongMarshallField(stage, "Data"))
		pointersInitializesStatements.WriteString(datashape.GongMarshallField(stage, "DataFlow"))
	}
	initRes = initializerStatements.String()
	ptrRes = pointersInitializesStatements.String()
	return
}
func (diagramstructure *DiagramStructure) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(diagramstructure.GongMarshallField(stage, "Name"))
		initializerStatements.WriteString(diagramstructure.GongMarshallField(stage, "Description"))
		initializerStatements.WriteString(diagramstructure.GongMarshallField(stage, "ComputedPrefix"))
		initializerStatements.WriteString(diagramstructure.GongMarshallField(stage, "IsExpanded"))
		initializerStatements.WriteString(diagramstructure.GongMarshallField(stage, "LayoutDirection"))
		initializerStatements.WriteString(diagramstructure.GongMarshallField(stage, "IsChecked"))
		initializerStatements.WriteString(diagramstructure.GongMarshallField(stage, "IsEditable_"))
		initializerStatements.WriteString(diagramstructure.GongMarshallField(stage, "IsShowPrefix"))
		initializerStatements.WriteString(diagramstructure.GongMarshallField(stage, "DefaultBoxWidth"))
		initializerStatements.WriteString(diagramstructure.GongMarshallField(stage, "DefaultBoxHeigth"))
		initializerStatements.WriteString(diagramstructure.GongMarshallField(stage, "Width"))
		initializerStatements.WriteString(diagramstructure.GongMarshallField(stage, "Height"))
		pointersInitializesStatements.WriteString(diagramstructure.GongMarshallField(stage, "System_Shapes"))
		initializerStatements.WriteString(diagramstructure.GongMarshallField(stage, "IsSystemsNodeExpanded"))
		pointersInitializesStatements.WriteString(diagramstructure.GongMarshallField(stage, "SystemsWhoseNodeIsExpanded"))
		pointersInitializesStatements.WriteString(diagramstructure.GongMarshallField(stage, "Part_Shapes"))
		initializerStatements.WriteString(diagramstructure.GongMarshallField(stage, "IsPartsNodeExpanded"))
		pointersInitializesStatements.WriteString(diagramstructure.GongMarshallField(stage, "PartWhoseNodeIsExpanded"))
		pointersInitializesStatements.WriteString(diagramstructure.GongMarshallField(stage, "ExternalPart_Shapes"))
		initializerStatements.WriteString(diagramstructure.GongMarshallField(stage, "IsExternalPartsNodeExpanded"))
		pointersInitializesStatements.WriteString(diagramstructure.GongMarshallField(stage, "ExternalPartWhoseNodeIsExpanded"))
		pointersInitializesStatements.WriteString(diagramstructure.GongMarshallField(stage, "ExternalPartsWhoseOutDataFlowsNodeIsExpanded"))
		pointersInitializesStatements.WriteString(diagramstructure.GongMarshallField(stage, "ExternalPartsWhoseInDataFlowsNodeIsExpanded"))
		pointersInitializesStatements.WriteString(diagramstructure.GongMarshallField(stage, "PortsWhoseNodeIsExpanded"))
		pointersInitializesStatements.WriteString(diagramstructure.GongMarshallField(stage, "Port_Shapes"))
		pointersInitializesStatements.WriteString(diagramstructure.GongMarshallField(stage, "ControlFlowsWhoseNodeIsExpanded"))
		pointersInitializesStatements.WriteString(diagramstructure.GongMarshallField(stage, "ControlFlow_Shapes"))
		pointersInitializesStatements.WriteString(diagramstructure.GongMarshallField(stage, "DataFlowsWhoseNodeIsExpanded"))
		pointersInitializesStatements.WriteString(diagramstructure.GongMarshallField(stage, "DataFlow_Shapes"))
		pointersInitializesStatements.WriteString(diagramstructure.GongMarshallField(stage, "DatasWhoseNodeIsExpanded"))
		pointersInitializesStatements.WriteString(diagramstructure.GongMarshallField(stage, "Data_Shapes"))
		pointersInitializesStatements.WriteString(diagramstructure.GongMarshallField(stage, "DataFlowsWhoseDataNodeIsExpanded"))
		pointersInitializesStatements.WriteString(diagramstructure.GongMarshallField(stage, "AllocatedResourcesWhoseNodeIsExpanded"))
		pointersInitializesStatements.WriteString(diagramstructure.GongMarshallField(stage, "AllocatedResourceShapes"))
		pointersInitializesStatements.WriteString(diagramstructure.GongMarshallField(stage, "AllocatedSystemesWhoseNodeIsExpanded"))
		pointersInitializesStatements.WriteString(diagramstructure.GongMarshallField(stage, "AllocatedSystemShapes"))
		pointersInitializesStatements.WriteString(diagramstructure.GongMarshallField(stage, "Note_Shapes"))
		pointersInitializesStatements.WriteString(diagramstructure.GongMarshallField(stage, "NotesWhoseNodeIsExpanded"))
		initializerStatements.WriteString(diagramstructure.GongMarshallField(stage, "IsNotesNodeExpanded"))
		pointersInitializesStatements.WriteString(diagramstructure.GongMarshallField(stage, "NotePortShapes"))
	}
	initRes = initializerStatements.String()
	ptrRes = pointersInitializesStatements.String()
	return
}
func (externalpartshape *ExternalPartShape) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(externalpartshape.GongMarshallField(stage, "Name"))
		pointersInitializesStatements.WriteString(externalpartshape.GongMarshallField(stage, "Part"))
		initializerStatements.WriteString(externalpartshape.GongMarshallField(stage, "IsExpanded"))
		initializerStatements.WriteString(externalpartshape.GongMarshallField(stage, "X"))
		initializerStatements.WriteString(externalpartshape.GongMarshallField(stage, "Y"))
		initializerStatements.WriteString(externalpartshape.GongMarshallField(stage, "Width"))
		initializerStatements.WriteString(externalpartshape.GongMarshallField(stage, "Height"))
		initializerStatements.WriteString(externalpartshape.GongMarshallField(stage, "IsHidden"))
		initializerStatements.WriteString(externalpartshape.GongMarshallField(stage, "TailHeigth"))
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
		initializerStatements.WriteString(library.GongMarshallField(stage, "Description"))
		initializerStatements.WriteString(library.GongMarshallField(stage, "ComputedPrefix"))
		initializerStatements.WriteString(library.GongMarshallField(stage, "IsExpanded"))
		initializerStatements.WriteString(library.GongMarshallField(stage, "LayoutDirection"))
		initializerStatements.WriteString(library.GongMarshallField(stage, "IsRootLibrary"))
		pointersInitializesStatements.WriteString(library.GongMarshallField(stage, "SubLibraries"))
		initializerStatements.WriteString(library.GongMarshallField(stage, "IsSubLibrariesNodeExpanded"))
		pointersInitializesStatements.WriteString(library.GongMarshallField(stage, "SubLibrariesWhoseNodeIsExpanded"))
		initializerStatements.WriteString(library.GongMarshallField(stage, "NbPixPerCharacter"))
		initializerStatements.WriteString(library.GongMarshallField(stage, "LogoSVGFile"))
		pointersInitializesStatements.WriteString(library.GongMarshallField(stage, "RootSystemes"))
		initializerStatements.WriteString(library.GongMarshallField(stage, "IsSystemesNodeExpanded"))
		pointersInitializesStatements.WriteString(library.GongMarshallField(stage, "SystemsWhoseNodeIsExpanded"))
		pointersInitializesStatements.WriteString(library.GongMarshallField(stage, "RootDataFlows"))
		initializerStatements.WriteString(library.GongMarshallField(stage, "IsDataFlowsNodeExpanded"))
		pointersInitializesStatements.WriteString(library.GongMarshallField(stage, "DataFlowsWhoseNodeIsExpanded"))
		pointersInitializesStatements.WriteString(library.GongMarshallField(stage, "RootDatas"))
		initializerStatements.WriteString(library.GongMarshallField(stage, "IsDatasNodeExpanded"))
		pointersInitializesStatements.WriteString(library.GongMarshallField(stage, "DatasWhoseNodeIsExpanded"))
		pointersInitializesStatements.WriteString(library.GongMarshallField(stage, "RootResources"))
		initializerStatements.WriteString(library.GongMarshallField(stage, "IsResourcesNodeExpanded"))
		pointersInitializesStatements.WriteString(library.GongMarshallField(stage, "ResourcesWhoseNodeIsExpanded"))
		pointersInitializesStatements.WriteString(library.GongMarshallField(stage, "PartsWhoseNodeIsExpanded"))
		pointersInitializesStatements.WriteString(library.GongMarshallField(stage, "RootNotes"))
		initializerStatements.WriteString(library.GongMarshallField(stage, "IsNotesNodeExpanded"))
		pointersInitializesStatements.WriteString(library.GongMarshallField(stage, "NotesWhoseNodeIsExpanded"))
		initializerStatements.WriteString(library.GongMarshallField(stage, "IsExpandedTmp"))
	}
	initRes = initializerStatements.String()
	ptrRes = pointersInitializesStatements.String()
	return
}
func (note *Note) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(note.GongMarshallField(stage, "Name"))
		initializerStatements.WriteString(note.GongMarshallField(stage, "Description"))
		initializerStatements.WriteString(note.GongMarshallField(stage, "ComputedPrefix"))
		initializerStatements.WriteString(note.GongMarshallField(stage, "IsExpanded"))
		initializerStatements.WriteString(note.GongMarshallField(stage, "LayoutDirection"))
		initializerStatements.WriteString(note.GongMarshallField(stage, "IsPortsNodeExpanded"))
		pointersInitializesStatements.WriteString(note.GongMarshallField(stage, "Ports"))
	}
	initRes = initializerStatements.String()
	ptrRes = pointersInitializesStatements.String()
	return
}
func (noteportshape *NotePortShape) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(noteportshape.GongMarshallField(stage, "Name"))
		pointersInitializesStatements.WriteString(noteportshape.GongMarshallField(stage, "Note"))
		pointersInitializesStatements.WriteString(noteportshape.GongMarshallField(stage, "Port"))
		initializerStatements.WriteString(noteportshape.GongMarshallField(stage, "StartRatio"))
		initializerStatements.WriteString(noteportshape.GongMarshallField(stage, "EndRatio"))
		initializerStatements.WriteString(noteportshape.GongMarshallField(stage, "StartOrientation"))
		initializerStatements.WriteString(noteportshape.GongMarshallField(stage, "EndOrientation"))
		initializerStatements.WriteString(noteportshape.GongMarshallField(stage, "CornerOffsetRatio"))
		initializerStatements.WriteString(noteportshape.GongMarshallField(stage, "IsHidden"))
	}
	initRes = initializerStatements.String()
	ptrRes = pointersInitializesStatements.String()
	return
}
func (noteshape *NoteShape) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(noteshape.GongMarshallField(stage, "Name"))
		pointersInitializesStatements.WriteString(noteshape.GongMarshallField(stage, "Note"))
		initializerStatements.WriteString(noteshape.GongMarshallField(stage, "X"))
		initializerStatements.WriteString(noteshape.GongMarshallField(stage, "Y"))
		initializerStatements.WriteString(noteshape.GongMarshallField(stage, "Width"))
		initializerStatements.WriteString(noteshape.GongMarshallField(stage, "Height"))
		initializerStatements.WriteString(noteshape.GongMarshallField(stage, "IsHidden"))
	}
	initRes = initializerStatements.String()
	ptrRes = pointersInitializesStatements.String()
	return
}
func (part *Part) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(part.GongMarshallField(stage, "Name"))
		initializerStatements.WriteString(part.GongMarshallField(stage, "IsSystemResource"))
		initializerStatements.WriteString(part.GongMarshallField(stage, "Description"))
		pointersInitializesStatements.WriteString(part.GongMarshallField(stage, "Resources"))
		initializerStatements.WriteString(part.GongMarshallField(stage, "IsResourcesNodeExpanded"))
		pointersInitializesStatements.WriteString(part.GongMarshallField(stage, "Systemes"))
		initializerStatements.WriteString(part.GongMarshallField(stage, "IsSystemesNodeExpanded"))
		initializerStatements.WriteString(part.GongMarshallField(stage, "ComputedPrefix"))
		initializerStatements.WriteString(part.GongMarshallField(stage, "IsExpanded"))
		initializerStatements.WriteString(part.GongMarshallField(stage, "LayoutDirection"))
		initializerStatements.WriteString(part.GongMarshallField(stage, "IsPortsNodeExpanded"))
		pointersInitializesStatements.WriteString(part.GongMarshallField(stage, "Ports"))
		initializerStatements.WriteString(part.GongMarshallField(stage, "IsControlFlowsNodeExpanded"))
		pointersInitializesStatements.WriteString(part.GongMarshallField(stage, "ControlFlows"))
		pointersInitializesStatements.WriteString(part.GongMarshallField(stage, "PortWhoseOutControlFlowsNodeIsExpanded"))
		pointersInitializesStatements.WriteString(part.GongMarshallField(stage, "PortWhoseInControlFlowsNodeIsExpanded"))
		initializerStatements.WriteString(part.GongMarshallField(stage, "IsDataFlowsNodeExpanded"))
		pointersInitializesStatements.WriteString(part.GongMarshallField(stage, "PortWhoseOutDataFlowsNodeIsExpanded"))
		pointersInitializesStatements.WriteString(part.GongMarshallField(stage, "PortWhoseInDataFlowsNodeIsExpanded"))
	}
	initRes = initializerStatements.String()
	ptrRes = pointersInitializesStatements.String()
	return
}
func (partshape *PartShape) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(partshape.GongMarshallField(stage, "Name"))
		pointersInitializesStatements.WriteString(partshape.GongMarshallField(stage, "Part"))
		initializerStatements.WriteString(partshape.GongMarshallField(stage, "IsExpanded"))
		initializerStatements.WriteString(partshape.GongMarshallField(stage, "X"))
		initializerStatements.WriteString(partshape.GongMarshallField(stage, "Y"))
		initializerStatements.WriteString(partshape.GongMarshallField(stage, "Width"))
		initializerStatements.WriteString(partshape.GongMarshallField(stage, "Height"))
		initializerStatements.WriteString(partshape.GongMarshallField(stage, "IsHidden"))
		initializerStatements.WriteString(partshape.GongMarshallField(stage, "WidthWeight"))
	}
	initRes = initializerStatements.String()
	ptrRes = pointersInitializesStatements.String()
	return
}
func (port *Port) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(port.GongMarshallField(stage, "Name"))
		initializerStatements.WriteString(port.GongMarshallField(stage, "Description"))
		initializerStatements.WriteString(port.GongMarshallField(stage, "ComputedPrefix"))
		initializerStatements.WriteString(port.GongMarshallField(stage, "IsExpanded"))
		initializerStatements.WriteString(port.GongMarshallField(stage, "LayoutDirection"))
		initializerStatements.WriteString(port.GongMarshallField(stage, "IsStartPort"))
		initializerStatements.WriteString(port.GongMarshallField(stage, "IsEndPort"))
		pointersInitializesStatements.WriteString(port.GongMarshallField(stage, "Type"))
		initializerStatements.WriteString(port.GongMarshallField(stage, "IsPortNameNotSystemName"))
	}
	initRes = initializerStatements.String()
	ptrRes = pointersInitializesStatements.String()
	return
}
func (portshape *PortShape) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(portshape.GongMarshallField(stage, "Name"))
		pointersInitializesStatements.WriteString(portshape.GongMarshallField(stage, "Port"))
		initializerStatements.WriteString(portshape.GongMarshallField(stage, "IsExpanded"))
		initializerStatements.WriteString(portshape.GongMarshallField(stage, "X"))
		initializerStatements.WriteString(portshape.GongMarshallField(stage, "Y"))
		initializerStatements.WriteString(portshape.GongMarshallField(stage, "Width"))
		initializerStatements.WriteString(portshape.GongMarshallField(stage, "Height"))
		initializerStatements.WriteString(portshape.GongMarshallField(stage, "IsHidden"))
	}
	initRes = initializerStatements.String()
	ptrRes = pointersInitializesStatements.String()
	return
}
func (resource *Resource) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(resource.GongMarshallField(stage, "Name"))
		initializerStatements.WriteString(resource.GongMarshallField(stage, "Acronym"))
		initializerStatements.WriteString(resource.GongMarshallField(stage, "Description"))
		initializerStatements.WriteString(resource.GongMarshallField(stage, "ComputedPrefix"))
		initializerStatements.WriteString(resource.GongMarshallField(stage, "IsExpanded"))
		initializerStatements.WriteString(resource.GongMarshallField(stage, "LayoutDirection"))
		initializerStatements.WriteString(resource.GongMarshallField(stage, "SVG_Path"))
		initializerStatements.WriteString(resource.GongMarshallField(stage, "InverseAppliedScaling"))
	}
	initRes = initializerStatements.String()
	ptrRes = pointersInitializesStatements.String()
	return
}
func (system *System) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(system.GongMarshallField(stage, "Name"))
		initializerStatements.WriteString(system.GongMarshallField(stage, "Description"))
		initializerStatements.WriteString(system.GongMarshallField(stage, "ComputedPrefix"))
		initializerStatements.WriteString(system.GongMarshallField(stage, "IsExpanded"))
		initializerStatements.WriteString(system.GongMarshallField(stage, "LayoutDirection"))
		initializerStatements.WriteString(system.GongMarshallField(stage, "SVG_Path"))
		initializerStatements.WriteString(system.GongMarshallField(stage, "InverseAppliedScaling"))
		pointersInitializesStatements.WriteString(system.GongMarshallField(stage, "DiagramStructures"))
		pointersInitializesStatements.WriteString(system.GongMarshallField(stage, "DiagramStructureWhoseNodeIsExpanded"))
		initializerStatements.WriteString(system.GongMarshallField(stage, "IsSubSystemNodeExpanded"))
		pointersInitializesStatements.WriteString(system.GongMarshallField(stage, "SubSystemes"))
		pointersInitializesStatements.WriteString(system.GongMarshallField(stage, "Parts"))
		pointersInitializesStatements.WriteString(system.GongMarshallField(stage, "PartWhoseNodeIsExpanded"))
		pointersInitializesStatements.WriteString(system.GongMarshallField(stage, "DataFlows"))
		initializerStatements.WriteString(system.GongMarshallField(stage, "IsDataFlowsNodeExpanded"))
		pointersInitializesStatements.WriteString(system.GongMarshallField(stage, "ExternalParts"))
		pointersInitializesStatements.WriteString(system.GongMarshallField(stage, "ExternalPartWhoseNodeIsExpanded"))
	}
	initRes = initializerStatements.String()
	ptrRes = pointersInitializesStatements.String()
	return
}
func (systemshape *SystemShape) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(systemshape.GongMarshallField(stage, "Name"))
		pointersInitializesStatements.WriteString(systemshape.GongMarshallField(stage, "System"))
		initializerStatements.WriteString(systemshape.GongMarshallField(stage, "IsExpanded"))
		initializerStatements.WriteString(systemshape.GongMarshallField(stage, "X"))
		initializerStatements.WriteString(systemshape.GongMarshallField(stage, "Y"))
		initializerStatements.WriteString(systemshape.GongMarshallField(stage, "Width"))
		initializerStatements.WriteString(systemshape.GongMarshallField(stage, "Height"))
		initializerStatements.WriteString(systemshape.GongMarshallField(stage, "IsHidden"))
	}
	initRes = initializerStatements.String()
	ptrRes = pointersInitializesStatements.String()
	return
}
