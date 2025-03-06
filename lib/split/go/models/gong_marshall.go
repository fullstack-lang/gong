// generated code - do not edit
package models

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"regexp"
	"sort"
	"strings"
)

const marshallRes = `package {{PackageName}}

import (
	"time"

	"{{ModelsPackageName}}"

	// injection point for ident package import declaration{{ImportPackageDeclaration}}
)

// generated in order to avoid error in the package import
// if there are no elements in the stage to marshall
var _ time.Time

// _ point for meta package dummy declaration{{ImportPackageDummyDeclaration}}

// When parsed, those maps will help with the renaming process
var _ map[string]any = map[string]any{
	// injection point for docLink to identifiers{{EntriesDocLinkStringDocLinkIdentifier}}
}

// function will stage objects
func _(stage *models.StageStruct) {

	// Declaration of instances to stage{{Identifiers}}

	// Setup of values{{ValueInitializers}}

	// Setup of pointers{{PointersInitializers}}
}`

const IdentifiersDecls = `
	{{Identifier}} := (&models.{{GeneratedStructName}}{}).Stage(stage)`

const StringInitStatement = `
	{{Identifier}}.{{GeneratedFieldName}} = ` + "`" + `{{GeneratedFieldNameValue}}` + "`"

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

// Marshall marshall the stage content into the file as an instanciation into a stage
func (stage *StageStruct) Marshall(file *os.File, modelsPackageName, packageName string) {

	name := file.Name()

	if !strings.HasSuffix(name, ".go") {
		log.Fatalln(name + " is not a go filename")
	}

	log.Println("filename of marshall output is " + name)
	newBase := filepath.Base(file.Name())

	res := marshallRes
	res = strings.ReplaceAll(res, "{{databaseName}}", strings.ReplaceAll(newBase, ".go", ""))
	res = strings.ReplaceAll(res, "{{PackageName}}", packageName)
	res = strings.ReplaceAll(res, "{{ModelsPackageName}}", modelsPackageName)

	// map of identifiers
	// var StageMapDstructIds map[*Dstruct]string
	identifiersDecl := ""
	initializerStatements := ""
	pointersInitializesStatements := ""

	id := ""
	_ = id
	decl := ""
	_ = decl
	setValueField := ""
	_ = setValueField

	// insertion initialization of objects to stage
	map_AsSplit_Identifiers := make(map[*AsSplit]string)
	_ = map_AsSplit_Identifiers

	assplitOrdered := []*AsSplit{}
	for assplit := range stage.AsSplits {
		assplitOrdered = append(assplitOrdered, assplit)
	}
	sort.Slice(assplitOrdered[:], func(i, j int) bool {
		asspliti := assplitOrdered[i]
		assplitj := assplitOrdered[j]
		asspliti_order, oki := stage.Map_Staged_Order[asspliti]
		assplitj_order, okj := stage.Map_Staged_Order[assplitj]
		if !oki || !okj {
			log.Fatalln("unknown pointers")
		}
		return asspliti_order < assplitj_order
	})
	if len(assplitOrdered) > 0 {
		identifiersDecl += "\n"
	}
	for idx, assplit := range assplitOrdered {

		id = generatesIdentifier("AsSplit", idx, assplit.Name)
		map_AsSplit_Identifiers[assplit] = id

		decl = IdentifiersDecls
		decl = strings.ReplaceAll(decl, "{{Identifier}}", id)
		decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "AsSplit")
		decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", assplit.Name)
		identifiersDecl += decl

		initializerStatements += "\n"
		// Initialisation of values
		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Name")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(assplit.Name))
		initializerStatements += setValueField

		if assplit.Direction != "" {
			setValueField = StringEnumInitStatement
			setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
			setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Direction")
			setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", "models."+assplit.Direction.ToCodeString())
			initializerStatements += setValueField
		}

	}

	map_AsSplitArea_Identifiers := make(map[*AsSplitArea]string)
	_ = map_AsSplitArea_Identifiers

	assplitareaOrdered := []*AsSplitArea{}
	for assplitarea := range stage.AsSplitAreas {
		assplitareaOrdered = append(assplitareaOrdered, assplitarea)
	}
	sort.Slice(assplitareaOrdered[:], func(i, j int) bool {
		assplitareai := assplitareaOrdered[i]
		assplitareaj := assplitareaOrdered[j]
		assplitareai_order, oki := stage.Map_Staged_Order[assplitareai]
		assplitareaj_order, okj := stage.Map_Staged_Order[assplitareaj]
		if !oki || !okj {
			log.Fatalln("unknown pointers")
		}
		return assplitareai_order < assplitareaj_order
	})
	if len(assplitareaOrdered) > 0 {
		identifiersDecl += "\n"
	}
	for idx, assplitarea := range assplitareaOrdered {

		id = generatesIdentifier("AsSplitArea", idx, assplitarea.Name)
		map_AsSplitArea_Identifiers[assplitarea] = id

		decl = IdentifiersDecls
		decl = strings.ReplaceAll(decl, "{{Identifier}}", id)
		decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "AsSplitArea")
		decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", assplitarea.Name)
		identifiersDecl += decl

		initializerStatements += "\n"
		// Initialisation of values
		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Name")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(assplitarea.Name))
		initializerStatements += setValueField

		setValueField = NumberInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Size")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%f", assplitarea.Size))
		initializerStatements += setValueField

		setValueField = NumberInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "IsAny")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%t", assplitarea.IsAny))
		initializerStatements += setValueField

	}

	map_View_Identifiers := make(map[*View]string)
	_ = map_View_Identifiers

	viewOrdered := []*View{}
	for view := range stage.Views {
		viewOrdered = append(viewOrdered, view)
	}
	sort.Slice(viewOrdered[:], func(i, j int) bool {
		viewi := viewOrdered[i]
		viewj := viewOrdered[j]
		viewi_order, oki := stage.Map_Staged_Order[viewi]
		viewj_order, okj := stage.Map_Staged_Order[viewj]
		if !oki || !okj {
			log.Fatalln("unknown pointers")
		}
		return viewi_order < viewj_order
	})
	if len(viewOrdered) > 0 {
		identifiersDecl += "\n"
	}
	for idx, view := range viewOrdered {

		id = generatesIdentifier("View", idx, view.Name)
		map_View_Identifiers[view] = id

		decl = IdentifiersDecls
		decl = strings.ReplaceAll(decl, "{{Identifier}}", id)
		decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "View")
		decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", view.Name)
		identifiersDecl += decl

		initializerStatements += "\n"
		// Initialisation of values
		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Name")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(view.Name))
		initializerStatements += setValueField

	}

	// insertion initialization of objects to stage
	if len(assplitOrdered) > 0 {
		pointersInitializesStatements += "\n\t// setup of AsSplit instances pointers"
	}
	for idx, assplit := range assplitOrdered {
		var setPointerField string
		_ = setPointerField

		id = generatesIdentifier("AsSplit", idx, assplit.Name)
		map_AsSplit_Identifiers[assplit] = id

		// Initialisation of values
		for _, _assplitarea := range assplit.AsSplitAreas {
			setPointerField = SliceOfPointersFieldInitStatement
			setPointerField = strings.ReplaceAll(setPointerField, "{{Identifier}}", id)
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldName}}", "AsSplitAreas")
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldNameValue}}", map_AsSplitArea_Identifiers[_assplitarea])
			pointersInitializesStatements += setPointerField
		}

	}

	if len(assplitareaOrdered) > 0 {
		pointersInitializesStatements += "\n\t// setup of AsSplitArea instances pointers"
	}
	for idx, assplitarea := range assplitareaOrdered {
		var setPointerField string
		_ = setPointerField

		id = generatesIdentifier("AsSplitArea", idx, assplitarea.Name)
		map_AsSplitArea_Identifiers[assplitarea] = id

		// Initialisation of values
		for _, _assplit := range assplitarea.AsSplits {
			setPointerField = SliceOfPointersFieldInitStatement
			setPointerField = strings.ReplaceAll(setPointerField, "{{Identifier}}", id)
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldName}}", "AsSplits")
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldNameValue}}", map_AsSplit_Identifiers[_assplit])
			pointersInitializesStatements += setPointerField
		}

	}

	if len(viewOrdered) > 0 {
		pointersInitializesStatements += "\n\t// setup of View instances pointers"
	}
	for idx, view := range viewOrdered {
		var setPointerField string
		_ = setPointerField

		id = generatesIdentifier("View", idx, view.Name)
		map_View_Identifiers[view] = id

		// Initialisation of values
		for _, _assplitarea := range view.RootAsSplitAreas {
			setPointerField = SliceOfPointersFieldInitStatement
			setPointerField = strings.ReplaceAll(setPointerField, "{{Identifier}}", id)
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldName}}", "RootAsSplitAreas")
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldNameValue}}", map_AsSplitArea_Identifiers[_assplitarea])
			pointersInitializesStatements += setPointerField
		}

	}

	res = strings.ReplaceAll(res, "{{Identifiers}}", identifiersDecl)
	res = strings.ReplaceAll(res, "{{ValueInitializers}}", initializerStatements)
	res = strings.ReplaceAll(res, "{{PointersInitializers}}", pointersInitializesStatements)

	if stage.MetaPackageImportAlias != "" {
		res = strings.ReplaceAll(res, "{{ImportPackageDeclaration}}",
			fmt.Sprintf("\n\t%s %s", stage.MetaPackageImportAlias, stage.MetaPackageImportPath))

		res = strings.ReplaceAll(res, "{{ImportPackageDummyDeclaration}}",
			fmt.Sprintf("\nvar _ %s.StageStruct",
				stage.MetaPackageImportAlias))

		var entries string

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
				entries += fmt.Sprintf("\n\n\t\"%s\": %s(0),", value.Ident, value.Ident)
			case GONG__ENUM_CAST_STRING:
				entries += fmt.Sprintf("\n\n\t\"%s\": %s(\"\"),", value.Ident, value.Ident)
			case GONG__FIELD_VALUE:
				// substitute the second point with "{})."
				joker := "__substitute_for_first_point__"
				valueIdentifier := strings.Replace(value.Ident, ".", joker, 1)
				valueIdentifier = strings.Replace(valueIdentifier, ".", "{}).", 1)
				valueIdentifier = strings.Replace(valueIdentifier, joker, ".", 1)
				entries += fmt.Sprintf("\n\n\t\"%s\": (%s,", value.Ident, valueIdentifier)
			case GONG__IDENTIFIER_CONST:
				entries += fmt.Sprintf("\n\n\t\"%s\": %s,", value.Ident, value.Ident)
			case GONG__STRUCT_INSTANCE:
				entries += fmt.Sprintf("\n\n\t\"%s\": &(%s{}),", value.Ident, value.Ident)
			}
		}

		// res = strings.ReplaceAll(res, "{{EntriesDocLinkStringDocLinkIdentifier}}", entries)
	}

	fmt.Fprintln(file, res)
}

// unique identifier per struct
func generatesIdentifier(gongStructName string, idx int, instanceName string) (identifier string) {

	identifier = instanceName
	// Make a Regex to say we only want letters and numbers
	reg, err := regexp.Compile("[^a-zA-Z0-9]+")
	if err != nil {
		log.Fatal(err)
	}
	processedString := reg.ReplaceAllString(instanceName, "_")

	identifier = fmt.Sprintf("__%s__%06d_%s", gongStructName, idx, processedString)

	return
}
