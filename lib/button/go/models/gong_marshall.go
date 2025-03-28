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
	map_Button_Identifiers := make(map[*Button]string)
	_ = map_Button_Identifiers

	buttonOrdered := []*Button{}
	for button := range stage.Buttons {
		buttonOrdered = append(buttonOrdered, button)
	}
	sort.Slice(buttonOrdered[:], func(i, j int) bool {
		buttoni := buttonOrdered[i]
		buttonj := buttonOrdered[j]
		buttoni_order, oki := stage.ButtonMap_Staged_Order[buttoni]
		buttonj_order, okj := stage.ButtonMap_Staged_Order[buttonj]
		if !oki || !okj {
			log.Fatalln("unknown pointers")
		}
		return buttoni_order < buttonj_order
	})
	if len(buttonOrdered) > 0 {
		identifiersDecl += "\n"
	}
	for idx, button := range buttonOrdered {

		id = generatesIdentifier("Button", idx, button.Name)
		map_Button_Identifiers[button] = id

		decl = IdentifiersDecls
		decl = strings.ReplaceAll(decl, "{{Identifier}}", id)
		decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "Button")
		decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", button.Name)
		identifiersDecl += decl

		initializerStatements += "\n"
		// Initialisation of values
		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Name")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(button.Name))
		initializerStatements += setValueField

		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Label")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(button.Label))
		initializerStatements += setValueField

		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Icon")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(button.Icon))
		initializerStatements += setValueField

	}

	map_Group_Identifiers := make(map[*Group]string)
	_ = map_Group_Identifiers

	groupOrdered := []*Group{}
	for group := range stage.Groups {
		groupOrdered = append(groupOrdered, group)
	}
	sort.Slice(groupOrdered[:], func(i, j int) bool {
		groupi := groupOrdered[i]
		groupj := groupOrdered[j]
		groupi_order, oki := stage.GroupMap_Staged_Order[groupi]
		groupj_order, okj := stage.GroupMap_Staged_Order[groupj]
		if !oki || !okj {
			log.Fatalln("unknown pointers")
		}
		return groupi_order < groupj_order
	})
	if len(groupOrdered) > 0 {
		identifiersDecl += "\n"
	}
	for idx, group := range groupOrdered {

		id = generatesIdentifier("Group", idx, group.Name)
		map_Group_Identifiers[group] = id

		decl = IdentifiersDecls
		decl = strings.ReplaceAll(decl, "{{Identifier}}", id)
		decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "Group")
		decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", group.Name)
		identifiersDecl += decl

		initializerStatements += "\n"
		// Initialisation of values
		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Name")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(group.Name))
		initializerStatements += setValueField

		setValueField = NumberInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Percentage")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%f", group.Percentage))
		initializerStatements += setValueField

	}

	map_Layout_Identifiers := make(map[*Layout]string)
	_ = map_Layout_Identifiers

	layoutOrdered := []*Layout{}
	for layout := range stage.Layouts {
		layoutOrdered = append(layoutOrdered, layout)
	}
	sort.Slice(layoutOrdered[:], func(i, j int) bool {
		layouti := layoutOrdered[i]
		layoutj := layoutOrdered[j]
		layouti_order, oki := stage.LayoutMap_Staged_Order[layouti]
		layoutj_order, okj := stage.LayoutMap_Staged_Order[layoutj]
		if !oki || !okj {
			log.Fatalln("unknown pointers")
		}
		return layouti_order < layoutj_order
	})
	if len(layoutOrdered) > 0 {
		identifiersDecl += "\n"
	}
	for idx, layout := range layoutOrdered {

		id = generatesIdentifier("Layout", idx, layout.Name)
		map_Layout_Identifiers[layout] = id

		decl = IdentifiersDecls
		decl = strings.ReplaceAll(decl, "{{Identifier}}", id)
		decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "Layout")
		decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", layout.Name)
		identifiersDecl += decl

		initializerStatements += "\n"
		// Initialisation of values
		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Name")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(layout.Name))
		initializerStatements += setValueField

	}

	// insertion initialization of objects to stage
	if len(buttonOrdered) > 0 {
		pointersInitializesStatements += "\n\t// setup of Button instances pointers"
	}
	for idx, button := range buttonOrdered {
		var setPointerField string
		_ = setPointerField

		id = generatesIdentifier("Button", idx, button.Name)
		map_Button_Identifiers[button] = id

		// Initialisation of values
	}

	if len(groupOrdered) > 0 {
		pointersInitializesStatements += "\n\t// setup of Group instances pointers"
	}
	for idx, group := range groupOrdered {
		var setPointerField string
		_ = setPointerField

		id = generatesIdentifier("Group", idx, group.Name)
		map_Group_Identifiers[group] = id

		// Initialisation of values
		for _, _button := range group.Buttons {
			setPointerField = SliceOfPointersFieldInitStatement
			setPointerField = strings.ReplaceAll(setPointerField, "{{Identifier}}", id)
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldName}}", "Buttons")
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldNameValue}}", map_Button_Identifiers[_button])
			pointersInitializesStatements += setPointerField
		}

	}

	if len(layoutOrdered) > 0 {
		pointersInitializesStatements += "\n\t// setup of Layout instances pointers"
	}
	for idx, layout := range layoutOrdered {
		var setPointerField string
		_ = setPointerField

		id = generatesIdentifier("Layout", idx, layout.Name)
		map_Layout_Identifiers[layout] = id

		// Initialisation of values
		for _, _group := range layout.Groups {
			setPointerField = SliceOfPointersFieldInitStatement
			setPointerField = strings.ReplaceAll(setPointerField, "{{Identifier}}", id)
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldName}}", "Groups")
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldNameValue}}", map_Group_Identifiers[_group])
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
