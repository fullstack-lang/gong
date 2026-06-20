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
		initializerStatements.WriteString(diagramstructure.GongMarshallField(stage, "ComputedPrefix"))
		initializerStatements.WriteString(diagramstructure.GongMarshallField(stage, "IsExpanded"))
		initializerStatements.WriteString(diagramstructure.GongMarshallField(stage, "LayoutDirection"))
		initializerStatements.WriteString(diagramstructure.GongMarshallField(stage, "IsChecked"))
		initializerStatements.WriteString(diagramstructure.GongMarshallField(stage, "IsEditable_"))
		initializerStatements.WriteString(diagramstructure.GongMarshallField(stage, "IsShowPrefix"))
		initializerStatements.WriteString(diagramstructure.GongMarshallField(stage, "Width"))
		initializerStatements.WriteString(diagramstructure.GongMarshallField(stage, "Height"))
		initializerStatements.WriteString(diagramstructure.GongMarshallField(stage, "DefaultBoxWidth"))
		initializerStatements.WriteString(diagramstructure.GongMarshallField(stage, "DefaultBoxHeigth"))
		pointersInitializesStatements.WriteString(diagramstructure.GongMarshallField(stage, "Part_Shapes"))
		initializerStatements.WriteString(diagramstructure.GongMarshallField(stage, "IsPartsNodeExpanded"))
		pointersInitializesStatements.WriteString(diagramstructure.GongMarshallField(stage, "PartsWhoseNodeIsExpanded"))
		pointersInitializesStatements.WriteString(diagramstructure.GongMarshallField(stage, "Link_Shapes"))
		initializerStatements.WriteString(diagramstructure.GongMarshallField(stage, "IsLinksNodeExpanded"))
		pointersInitializesStatements.WriteString(diagramstructure.GongMarshallField(stage, "LinksWhoseNodeIsExpanded"))
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
		pointersInitializesStatements.WriteString(library.GongMarshallField(stage, "SubLibraries"))
		initializerStatements.WriteString(library.GongMarshallField(stage, "IsSubLibrariesNodeExpanded"))
		pointersInitializesStatements.WriteString(library.GongMarshallField(stage, "SubLibrariesWhoseNodeIsExpanded"))
		initializerStatements.WriteString(library.GongMarshallField(stage, "NbPixPerCharacter"))
		initializerStatements.WriteString(library.GongMarshallField(stage, "LogoSVGFile"))
		initializerStatements.WriteString(library.GongMarshallField(stage, "ComputedPrefix"))
		initializerStatements.WriteString(library.GongMarshallField(stage, "IsExpanded"))
		initializerStatements.WriteString(library.GongMarshallField(stage, "LayoutDirection"))
		initializerStatements.WriteString(library.GongMarshallField(stage, "IsRootLibrary"))
		pointersInitializesStatements.WriteString(library.GongMarshallField(stage, "RootSystems"))
		initializerStatements.WriteString(library.GongMarshallField(stage, "IsSystemsNodeExpanded"))
		pointersInitializesStatements.WriteString(library.GongMarshallField(stage, "SystemsWhoseNodeIsExpanded"))
	}

	linkOrdered := []*Link{}
	for link := range stage.Links {
		linkOrdered = append(linkOrdered, link)
	}
	sort.Slice(linkOrdered[:], func(i, j int) bool {
		linki := linkOrdered[i]
		linkj := linkOrdered[j]
		linki_order, oki := stage.Link_stagedOrder[linki]
		linkj_order, okj := stage.Link_stagedOrder[linkj]
		if !oki || !okj {
			log.Fatalln("unknown pointers")
		}
		return linki_order < linkj_order
	})
	if len(linkOrdered) > 0 {
		identifiersDecl.WriteString("\n")
	}
	for _, link := range linkOrdered {

		identifiersDecl.WriteString(link.GongMarshallIdentifier(stage))

		initializerStatements.WriteString("\n")
		// Insertion point for basic fields value assignment
		initializerStatements.WriteString(link.GongMarshallField(stage, "Name"))
		initializerStatements.WriteString(link.GongMarshallField(stage, "ComputedPrefix"))
		initializerStatements.WriteString(link.GongMarshallField(stage, "IsExpanded"))
		initializerStatements.WriteString(link.GongMarshallField(stage, "LayoutDirection"))
		pointersInitializesStatements.WriteString(link.GongMarshallField(stage, "Source"))
		pointersInitializesStatements.WriteString(link.GongMarshallField(stage, "Target"))
	}

	linkassociationshapeOrdered := []*LinkAssociationShape{}
	for linkassociationshape := range stage.LinkAssociationShapes {
		linkassociationshapeOrdered = append(linkassociationshapeOrdered, linkassociationshape)
	}
	sort.Slice(linkassociationshapeOrdered[:], func(i, j int) bool {
		linkassociationshapei := linkassociationshapeOrdered[i]
		linkassociationshapej := linkassociationshapeOrdered[j]
		linkassociationshapei_order, oki := stage.LinkAssociationShape_stagedOrder[linkassociationshapei]
		linkassociationshapej_order, okj := stage.LinkAssociationShape_stagedOrder[linkassociationshapej]
		if !oki || !okj {
			log.Fatalln("unknown pointers")
		}
		return linkassociationshapei_order < linkassociationshapej_order
	})
	if len(linkassociationshapeOrdered) > 0 {
		identifiersDecl.WriteString("\n")
	}
	for _, linkassociationshape := range linkassociationshapeOrdered {

		identifiersDecl.WriteString(linkassociationshape.GongMarshallIdentifier(stage))

		initializerStatements.WriteString("\n")
		// Insertion point for basic fields value assignment
		initializerStatements.WriteString(linkassociationshape.GongMarshallField(stage, "Name"))
		pointersInitializesStatements.WriteString(linkassociationshape.GongMarshallField(stage, "Link"))
		initializerStatements.WriteString(linkassociationshape.GongMarshallField(stage, "StartRatio"))
		initializerStatements.WriteString(linkassociationshape.GongMarshallField(stage, "EndRatio"))
		initializerStatements.WriteString(linkassociationshape.GongMarshallField(stage, "StartOrientation"))
		initializerStatements.WriteString(linkassociationshape.GongMarshallField(stage, "EndOrientation"))
		initializerStatements.WriteString(linkassociationshape.GongMarshallField(stage, "CornerOffsetRatio"))
		initializerStatements.WriteString(linkassociationshape.GongMarshallField(stage, "IsHidden"))
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
		initializerStatements.WriteString(part.GongMarshallField(stage, "ComputedPrefix"))
		initializerStatements.WriteString(part.GongMarshallField(stage, "IsExpanded"))
		initializerStatements.WriteString(part.GongMarshallField(stage, "LayoutDirection"))
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
		initializerStatements.WriteString(partshape.GongMarshallField(stage, "OverideLayoutDirection"))
		initializerStatements.WriteString(partshape.GongMarshallField(stage, "LayoutDirection"))
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
		initializerStatements.WriteString(system.GongMarshallField(stage, "ComputedPrefix"))
		initializerStatements.WriteString(system.GongMarshallField(stage, "IsExpanded"))
		initializerStatements.WriteString(system.GongMarshallField(stage, "LayoutDirection"))
		pointersInitializesStatements.WriteString(system.GongMarshallField(stage, "Parts"))
		initializerStatements.WriteString(system.GongMarshallField(stage, "IsPartsNodeExpanded"))
		pointersInitializesStatements.WriteString(system.GongMarshallField(stage, "PartsWhoseNodeIsExpanded"))
		pointersInitializesStatements.WriteString(system.GongMarshallField(stage, "SubSystems"))
		initializerStatements.WriteString(system.GongMarshallField(stage, "IsSubSystemsNodeExpanded"))
		pointersInitializesStatements.WriteString(system.GongMarshallField(stage, "SubSystemsWhoseNodeIsExpanded"))
		pointersInitializesStatements.WriteString(system.GongMarshallField(stage, "Links"))
		initializerStatements.WriteString(system.GongMarshallField(stage, "IsLinksNodeExpanded"))
		pointersInitializesStatements.WriteString(system.GongMarshallField(stage, "LinksWhoseNodeIsExpanded"))
		pointersInitializesStatements.WriteString(system.GongMarshallField(stage, "DiagramStructures"))
		initializerStatements.WriteString(system.GongMarshallField(stage, "IsDiagramStructuresNodeExpanded"))
		pointersInitializesStatements.WriteString(system.GongMarshallField(stage, "DiagramStructuresWhoseNodeIsExpanded"))
	}

	// insertion initialization of objects to stage
	for _, diagramstructure := range diagramstructureOrdered {
		_ = diagramstructure
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

	for _, link := range linkOrdered {
		_ = link
		var setPointerField string
		_ = setPointerField

		// Insertion point for pointers initialization
	}

	for _, linkassociationshape := range linkassociationshapeOrdered {
		_ = linkassociationshape
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

	for _, system := range systemOrdered {
		_ = system
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
func (diagramstructure *DiagramStructure) GongMarshallField(stage *Stage, fieldName string) (res string) {

	switch fieldName {
	case "Name":
		res = StringInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", diagramstructure.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "Name")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", ToRawStringLiteral(diagramstructure.Name))
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
	case "IsPartsNodeExpanded":
		res = NumberInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", diagramstructure.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "IsPartsNodeExpanded")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%t", diagramstructure.IsPartsNodeExpanded))
	case "IsLinksNodeExpanded":
		res = NumberInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", diagramstructure.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "IsLinksNodeExpanded")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%t", diagramstructure.IsLinksNodeExpanded))

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
	case "PartsWhoseNodeIsExpanded":
		var sb strings.Builder
		for _, _part := range diagramstructure.PartsWhoseNodeIsExpanded {
			tmp := SliceOfPointersFieldInitStatement
			tmp = strings.ReplaceAll(tmp, "{{Identifier}}", diagramstructure.GongGetIdentifier(stage))
			tmp = strings.ReplaceAll(tmp, "{{GeneratedFieldName}}", "PartsWhoseNodeIsExpanded")
			tmp = strings.ReplaceAll(tmp, "{{GeneratedFieldNameValue}}", _part.GongGetIdentifier(stage))
			sb.WriteString(tmp)
		}
		res = sb.String()
	case "Link_Shapes":
		var sb strings.Builder
		for _, _linkassociationshape := range diagramstructure.Link_Shapes {
			tmp := SliceOfPointersFieldInitStatement
			tmp = strings.ReplaceAll(tmp, "{{Identifier}}", diagramstructure.GongGetIdentifier(stage))
			tmp = strings.ReplaceAll(tmp, "{{GeneratedFieldName}}", "Link_Shapes")
			tmp = strings.ReplaceAll(tmp, "{{GeneratedFieldNameValue}}", _linkassociationshape.GongGetIdentifier(stage))
			sb.WriteString(tmp)
		}
		res = sb.String()
	case "LinksWhoseNodeIsExpanded":
		var sb strings.Builder
		for _, _link := range diagramstructure.LinksWhoseNodeIsExpanded {
			tmp := SliceOfPointersFieldInitStatement
			tmp = strings.ReplaceAll(tmp, "{{Identifier}}", diagramstructure.GongGetIdentifier(stage))
			tmp = strings.ReplaceAll(tmp, "{{GeneratedFieldName}}", "LinksWhoseNodeIsExpanded")
			tmp = strings.ReplaceAll(tmp, "{{GeneratedFieldNameValue}}", _link.GongGetIdentifier(stage))
			sb.WriteString(tmp)
		}
		res = sb.String()
	default:
		log.Panicf("Unknown field %s for Gongstruct DiagramStructure", fieldName)
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
	case "IsSystemsNodeExpanded":
		res = NumberInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", library.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "IsSystemsNodeExpanded")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%t", library.IsSystemsNodeExpanded))

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
	case "RootSystems":
		var sb strings.Builder
		for _, _system := range library.RootSystems {
			tmp := SliceOfPointersFieldInitStatement
			tmp = strings.ReplaceAll(tmp, "{{Identifier}}", library.GongGetIdentifier(stage))
			tmp = strings.ReplaceAll(tmp, "{{GeneratedFieldName}}", "RootSystems")
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
	default:
		log.Panicf("Unknown field %s for Gongstruct Library", fieldName)
	}
	return
}

func (link *Link) GongMarshallField(stage *Stage, fieldName string) (res string) {

	switch fieldName {
	case "Name":
		res = StringInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", link.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "Name")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", ToRawStringLiteral(link.Name))
	case "ComputedPrefix":
		res = StringInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", link.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "ComputedPrefix")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", ToRawStringLiteral(link.ComputedPrefix))
	case "IsExpanded":
		res = NumberInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", link.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "IsExpanded")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%t", link.IsExpanded))
	case "LayoutDirection":
		if link.LayoutDirection.ToCodeString() != "" {
			res = NumberInitStatement
			res = strings.ReplaceAll(res, "{{Identifier}}", link.GongGetIdentifier(stage))
			res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "LayoutDirection")
			res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", "models."+link.LayoutDirection.ToCodeString())
		} else {
			// in case of empty enum, we need to unstage the previous value
			res = NumberInitStatement
			res = strings.ReplaceAll(res, "{{Identifier}}", link.GongGetIdentifier(stage))
			res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "LayoutDirection")
			res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", "0")
		}

	case "Source":
		if link.Source != nil {
			res = PointerFieldInitStatement
			res = strings.ReplaceAll(res, "{{Identifier}}", link.GongGetIdentifier(stage))
			res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "Source")
			res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", link.Source.GongGetIdentifier(stage))
		} else {
			// in case of nil pointer, we need to unstage the previous value
			res = PointerFieldInitStatement
			res = strings.ReplaceAll(res, "{{Identifier}}", link.GongGetIdentifier(stage))
			res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "Source")
			res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", "nil")
		}
	case "Target":
		if link.Target != nil {
			res = PointerFieldInitStatement
			res = strings.ReplaceAll(res, "{{Identifier}}", link.GongGetIdentifier(stage))
			res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "Target")
			res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", link.Target.GongGetIdentifier(stage))
		} else {
			// in case of nil pointer, we need to unstage the previous value
			res = PointerFieldInitStatement
			res = strings.ReplaceAll(res, "{{Identifier}}", link.GongGetIdentifier(stage))
			res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "Target")
			res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", "nil")
		}
	default:
		log.Panicf("Unknown field %s for Gongstruct Link", fieldName)
	}
	return
}

func (linkassociationshape *LinkAssociationShape) GongMarshallField(stage *Stage, fieldName string) (res string) {

	switch fieldName {
	case "Name":
		res = StringInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", linkassociationshape.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "Name")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", ToRawStringLiteral(linkassociationshape.Name))
	case "StartRatio":
		res = NumberInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", linkassociationshape.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "StartRatio")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%f", linkassociationshape.StartRatio))
	case "EndRatio":
		res = NumberInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", linkassociationshape.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "EndRatio")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%f", linkassociationshape.EndRatio))
	case "StartOrientation":
		if linkassociationshape.StartOrientation.ToCodeString() != "" {
			res = StringEnumInitStatement
			res = strings.ReplaceAll(res, "{{Identifier}}", linkassociationshape.GongGetIdentifier(stage))
			res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "StartOrientation")
			res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", "models."+linkassociationshape.StartOrientation.ToCodeString())
		} else {
			// in case of empty enum, we need to unstage the previous value
			res = StringEnumInitStatement
			res = strings.ReplaceAll(res, "{{Identifier}}", linkassociationshape.GongGetIdentifier(stage))
			res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "StartOrientation")
			res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", "\"\"")
		}
	case "EndOrientation":
		if linkassociationshape.EndOrientation.ToCodeString() != "" {
			res = StringEnumInitStatement
			res = strings.ReplaceAll(res, "{{Identifier}}", linkassociationshape.GongGetIdentifier(stage))
			res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "EndOrientation")
			res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", "models."+linkassociationshape.EndOrientation.ToCodeString())
		} else {
			// in case of empty enum, we need to unstage the previous value
			res = StringEnumInitStatement
			res = strings.ReplaceAll(res, "{{Identifier}}", linkassociationshape.GongGetIdentifier(stage))
			res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "EndOrientation")
			res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", "\"\"")
		}
	case "CornerOffsetRatio":
		res = NumberInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", linkassociationshape.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "CornerOffsetRatio")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%f", linkassociationshape.CornerOffsetRatio))
	case "IsHidden":
		res = NumberInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", linkassociationshape.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "IsHidden")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%t", linkassociationshape.IsHidden))

	case "Link":
		if linkassociationshape.Link != nil {
			res = PointerFieldInitStatement
			res = strings.ReplaceAll(res, "{{Identifier}}", linkassociationshape.GongGetIdentifier(stage))
			res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "Link")
			res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", linkassociationshape.Link.GongGetIdentifier(stage))
		} else {
			// in case of nil pointer, we need to unstage the previous value
			res = PointerFieldInitStatement
			res = strings.ReplaceAll(res, "{{Identifier}}", linkassociationshape.GongGetIdentifier(stage))
			res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "Link")
			res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", "nil")
		}
	default:
		log.Panicf("Unknown field %s for Gongstruct LinkAssociationShape", fieldName)
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
	case "OverideLayoutDirection":
		res = NumberInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", partshape.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "OverideLayoutDirection")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%t", partshape.OverideLayoutDirection))
	case "LayoutDirection":
		if partshape.LayoutDirection.ToCodeString() != "" {
			res = NumberInitStatement
			res = strings.ReplaceAll(res, "{{Identifier}}", partshape.GongGetIdentifier(stage))
			res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "LayoutDirection")
			res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", "models."+partshape.LayoutDirection.ToCodeString())
		} else {
			// in case of empty enum, we need to unstage the previous value
			res = NumberInitStatement
			res = strings.ReplaceAll(res, "{{Identifier}}", partshape.GongGetIdentifier(stage))
			res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "LayoutDirection")
			res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", "0")
		}

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

func (system *System) GongMarshallField(stage *Stage, fieldName string) (res string) {

	switch fieldName {
	case "Name":
		res = StringInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", system.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "Name")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", ToRawStringLiteral(system.Name))
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
	case "IsPartsNodeExpanded":
		res = NumberInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", system.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "IsPartsNodeExpanded")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%t", system.IsPartsNodeExpanded))
	case "IsSubSystemsNodeExpanded":
		res = NumberInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", system.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "IsSubSystemsNodeExpanded")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%t", system.IsSubSystemsNodeExpanded))
	case "IsLinksNodeExpanded":
		res = NumberInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", system.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "IsLinksNodeExpanded")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%t", system.IsLinksNodeExpanded))
	case "IsDiagramStructuresNodeExpanded":
		res = NumberInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", system.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "IsDiagramStructuresNodeExpanded")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%t", system.IsDiagramStructuresNodeExpanded))

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
	case "PartsWhoseNodeIsExpanded":
		var sb strings.Builder
		for _, _part := range system.PartsWhoseNodeIsExpanded {
			tmp := SliceOfPointersFieldInitStatement
			tmp = strings.ReplaceAll(tmp, "{{Identifier}}", system.GongGetIdentifier(stage))
			tmp = strings.ReplaceAll(tmp, "{{GeneratedFieldName}}", "PartsWhoseNodeIsExpanded")
			tmp = strings.ReplaceAll(tmp, "{{GeneratedFieldNameValue}}", _part.GongGetIdentifier(stage))
			sb.WriteString(tmp)
		}
		res = sb.String()
	case "SubSystems":
		var sb strings.Builder
		for _, _system := range system.SubSystems {
			tmp := SliceOfPointersFieldInitStatement
			tmp = strings.ReplaceAll(tmp, "{{Identifier}}", system.GongGetIdentifier(stage))
			tmp = strings.ReplaceAll(tmp, "{{GeneratedFieldName}}", "SubSystems")
			tmp = strings.ReplaceAll(tmp, "{{GeneratedFieldNameValue}}", _system.GongGetIdentifier(stage))
			sb.WriteString(tmp)
		}
		res = sb.String()
	case "SubSystemsWhoseNodeIsExpanded":
		var sb strings.Builder
		for _, _system := range system.SubSystemsWhoseNodeIsExpanded {
			tmp := SliceOfPointersFieldInitStatement
			tmp = strings.ReplaceAll(tmp, "{{Identifier}}", system.GongGetIdentifier(stage))
			tmp = strings.ReplaceAll(tmp, "{{GeneratedFieldName}}", "SubSystemsWhoseNodeIsExpanded")
			tmp = strings.ReplaceAll(tmp, "{{GeneratedFieldNameValue}}", _system.GongGetIdentifier(stage))
			sb.WriteString(tmp)
		}
		res = sb.String()
	case "Links":
		var sb strings.Builder
		for _, _link := range system.Links {
			tmp := SliceOfPointersFieldInitStatement
			tmp = strings.ReplaceAll(tmp, "{{Identifier}}", system.GongGetIdentifier(stage))
			tmp = strings.ReplaceAll(tmp, "{{GeneratedFieldName}}", "Links")
			tmp = strings.ReplaceAll(tmp, "{{GeneratedFieldNameValue}}", _link.GongGetIdentifier(stage))
			sb.WriteString(tmp)
		}
		res = sb.String()
	case "LinksWhoseNodeIsExpanded":
		var sb strings.Builder
		for _, _link := range system.LinksWhoseNodeIsExpanded {
			tmp := SliceOfPointersFieldInitStatement
			tmp = strings.ReplaceAll(tmp, "{{Identifier}}", system.GongGetIdentifier(stage))
			tmp = strings.ReplaceAll(tmp, "{{GeneratedFieldName}}", "LinksWhoseNodeIsExpanded")
			tmp = strings.ReplaceAll(tmp, "{{GeneratedFieldNameValue}}", _link.GongGetIdentifier(stage))
			sb.WriteString(tmp)
		}
		res = sb.String()
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
	case "DiagramStructuresWhoseNodeIsExpanded":
		var sb strings.Builder
		for _, _diagramstructure := range system.DiagramStructuresWhoseNodeIsExpanded {
			tmp := SliceOfPointersFieldInitStatement
			tmp = strings.ReplaceAll(tmp, "{{Identifier}}", system.GongGetIdentifier(stage))
			tmp = strings.ReplaceAll(tmp, "{{GeneratedFieldName}}", "DiagramStructuresWhoseNodeIsExpanded")
			tmp = strings.ReplaceAll(tmp, "{{GeneratedFieldNameValue}}", _diagramstructure.GongGetIdentifier(stage))
			sb.WriteString(tmp)
		}
		res = sb.String()
	default:
		log.Panicf("Unknown field %s for Gongstruct System", fieldName)
	}
	return
}

// insertion point for marshall all fields methods
func (diagramstructure *DiagramStructure) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(diagramstructure.GongMarshallField(stage, "Name"))
		initializerStatements.WriteString(diagramstructure.GongMarshallField(stage, "ComputedPrefix"))
		initializerStatements.WriteString(diagramstructure.GongMarshallField(stage, "IsExpanded"))
		initializerStatements.WriteString(diagramstructure.GongMarshallField(stage, "LayoutDirection"))
		initializerStatements.WriteString(diagramstructure.GongMarshallField(stage, "IsChecked"))
		initializerStatements.WriteString(diagramstructure.GongMarshallField(stage, "IsEditable_"))
		initializerStatements.WriteString(diagramstructure.GongMarshallField(stage, "IsShowPrefix"))
		initializerStatements.WriteString(diagramstructure.GongMarshallField(stage, "Width"))
		initializerStatements.WriteString(diagramstructure.GongMarshallField(stage, "Height"))
		initializerStatements.WriteString(diagramstructure.GongMarshallField(stage, "DefaultBoxWidth"))
		initializerStatements.WriteString(diagramstructure.GongMarshallField(stage, "DefaultBoxHeigth"))
		pointersInitializesStatements.WriteString(diagramstructure.GongMarshallField(stage, "Part_Shapes"))
		initializerStatements.WriteString(diagramstructure.GongMarshallField(stage, "IsPartsNodeExpanded"))
		pointersInitializesStatements.WriteString(diagramstructure.GongMarshallField(stage, "PartsWhoseNodeIsExpanded"))
		pointersInitializesStatements.WriteString(diagramstructure.GongMarshallField(stage, "Link_Shapes"))
		initializerStatements.WriteString(diagramstructure.GongMarshallField(stage, "IsLinksNodeExpanded"))
		pointersInitializesStatements.WriteString(diagramstructure.GongMarshallField(stage, "LinksWhoseNodeIsExpanded"))
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
		pointersInitializesStatements.WriteString(library.GongMarshallField(stage, "SubLibraries"))
		initializerStatements.WriteString(library.GongMarshallField(stage, "IsSubLibrariesNodeExpanded"))
		pointersInitializesStatements.WriteString(library.GongMarshallField(stage, "SubLibrariesWhoseNodeIsExpanded"))
		initializerStatements.WriteString(library.GongMarshallField(stage, "NbPixPerCharacter"))
		initializerStatements.WriteString(library.GongMarshallField(stage, "LogoSVGFile"))
		initializerStatements.WriteString(library.GongMarshallField(stage, "ComputedPrefix"))
		initializerStatements.WriteString(library.GongMarshallField(stage, "IsExpanded"))
		initializerStatements.WriteString(library.GongMarshallField(stage, "LayoutDirection"))
		initializerStatements.WriteString(library.GongMarshallField(stage, "IsRootLibrary"))
		pointersInitializesStatements.WriteString(library.GongMarshallField(stage, "RootSystems"))
		initializerStatements.WriteString(library.GongMarshallField(stage, "IsSystemsNodeExpanded"))
		pointersInitializesStatements.WriteString(library.GongMarshallField(stage, "SystemsWhoseNodeIsExpanded"))
	}
	initRes = initializerStatements.String()
	ptrRes = pointersInitializesStatements.String()
	return
}
func (link *Link) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(link.GongMarshallField(stage, "Name"))
		initializerStatements.WriteString(link.GongMarshallField(stage, "ComputedPrefix"))
		initializerStatements.WriteString(link.GongMarshallField(stage, "IsExpanded"))
		initializerStatements.WriteString(link.GongMarshallField(stage, "LayoutDirection"))
		pointersInitializesStatements.WriteString(link.GongMarshallField(stage, "Source"))
		pointersInitializesStatements.WriteString(link.GongMarshallField(stage, "Target"))
	}
	initRes = initializerStatements.String()
	ptrRes = pointersInitializesStatements.String()
	return
}
func (linkassociationshape *LinkAssociationShape) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(linkassociationshape.GongMarshallField(stage, "Name"))
		pointersInitializesStatements.WriteString(linkassociationshape.GongMarshallField(stage, "Link"))
		initializerStatements.WriteString(linkassociationshape.GongMarshallField(stage, "StartRatio"))
		initializerStatements.WriteString(linkassociationshape.GongMarshallField(stage, "EndRatio"))
		initializerStatements.WriteString(linkassociationshape.GongMarshallField(stage, "StartOrientation"))
		initializerStatements.WriteString(linkassociationshape.GongMarshallField(stage, "EndOrientation"))
		initializerStatements.WriteString(linkassociationshape.GongMarshallField(stage, "CornerOffsetRatio"))
		initializerStatements.WriteString(linkassociationshape.GongMarshallField(stage, "IsHidden"))
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
		initializerStatements.WriteString(part.GongMarshallField(stage, "ComputedPrefix"))
		initializerStatements.WriteString(part.GongMarshallField(stage, "IsExpanded"))
		initializerStatements.WriteString(part.GongMarshallField(stage, "LayoutDirection"))
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
		initializerStatements.WriteString(partshape.GongMarshallField(stage, "OverideLayoutDirection"))
		initializerStatements.WriteString(partshape.GongMarshallField(stage, "LayoutDirection"))
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
		initializerStatements.WriteString(system.GongMarshallField(stage, "ComputedPrefix"))
		initializerStatements.WriteString(system.GongMarshallField(stage, "IsExpanded"))
		initializerStatements.WriteString(system.GongMarshallField(stage, "LayoutDirection"))
		pointersInitializesStatements.WriteString(system.GongMarshallField(stage, "Parts"))
		initializerStatements.WriteString(system.GongMarshallField(stage, "IsPartsNodeExpanded"))
		pointersInitializesStatements.WriteString(system.GongMarshallField(stage, "PartsWhoseNodeIsExpanded"))
		pointersInitializesStatements.WriteString(system.GongMarshallField(stage, "SubSystems"))
		initializerStatements.WriteString(system.GongMarshallField(stage, "IsSubSystemsNodeExpanded"))
		pointersInitializesStatements.WriteString(system.GongMarshallField(stage, "SubSystemsWhoseNodeIsExpanded"))
		pointersInitializesStatements.WriteString(system.GongMarshallField(stage, "Links"))
		initializerStatements.WriteString(system.GongMarshallField(stage, "IsLinksNodeExpanded"))
		pointersInitializesStatements.WriteString(system.GongMarshallField(stage, "LinksWhoseNodeIsExpanded"))
		pointersInitializesStatements.WriteString(system.GongMarshallField(stage, "DiagramStructures"))
		initializerStatements.WriteString(system.GongMarshallField(stage, "IsDiagramStructuresNodeExpanded"))
		pointersInitializesStatements.WriteString(system.GongMarshallField(stage, "DiagramStructuresWhoseNodeIsExpanded"))
	}
	initRes = initializerStatements.String()
	ptrRes = pointersInitializesStatements.String()
	return
}
