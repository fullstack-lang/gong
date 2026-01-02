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
func _(stage *models.Stage) {

	// insertion point for declaration of instances to stage{{Identifiers}}

	// insertion point for initialization of values{{ValueInitializers}}

	// insertion point for setup of pointers{{PointersInitializers}}
}`

const IdentifiersDecls = `
	{{Identifier}} := (&models.{{GeneratedStructName}}{Name: ` + "`" + `{{GeneratedFieldNameValue}}` + "`" + `}).Stage(stage)`

// previous version does not hanldle embedded structs (https://github.com/golang/go/issues/9859)
// simpler version but the name of the instance cannot be human read before the fields initialization
const IdentifiersDeclsWithoutNameInit = `
	{{Identifier}} := (&models.{{GeneratedStructName}}{}).Stage(stage)` /* */

const StringInitStatement = `
	{{Identifier}}.{{GeneratedFieldName}} = ` + "`" + `{{GeneratedFieldNameValue}}` + "`"

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
	identifiersDecl := ""
	initializerStatements := ""
	pointersInitializesStatements := ""

	decl := ""
	_ = decl
	setValueField := ""
	_ = setValueField

	// insertion initialization of objects to stage
	checkboxOrdered := []*Checkbox{}
	for checkbox := range stage.Checkboxs {
		checkboxOrdered = append(checkboxOrdered, checkbox)
	}
	sort.Slice(checkboxOrdered[:], func(i, j int) bool {
		checkboxi := checkboxOrdered[i]
		checkboxj := checkboxOrdered[j]
		checkboxi_order, oki := stage.CheckboxMap_Staged_Order[checkboxi]
		checkboxj_order, okj := stage.CheckboxMap_Staged_Order[checkboxj]
		if !oki || !okj {
			log.Fatalln("unknown pointers")
		}
		return checkboxi_order < checkboxj_order
	})
	if len(checkboxOrdered) > 0 {
		identifiersDecl += "\n"
	}
	for _, checkbox := range checkboxOrdered {

		identifiersDecl += checkbox.GongMarshallIdentifier(stage)

		initializerStatements += "\n"
		// Insertion point for basic fields value assignment
		initializerStatements += checkbox.GongMarshallField(stage, "Name")
		initializerStatements += checkbox.GongMarshallField(stage, "ValueBool")
		initializerStatements += checkbox.GongMarshallField(stage, "LabelForTrue")
		initializerStatements += checkbox.GongMarshallField(stage, "LabelForFalse")
	}

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
	for _, group := range groupOrdered {

		identifiersDecl += group.GongMarshallIdentifier(stage)

		initializerStatements += "\n"
		// Insertion point for basic fields value assignment
		initializerStatements += group.GongMarshallField(stage, "Name")
		initializerStatements += group.GongMarshallField(stage, "Percentage")
		pointersInitializesStatements += group.GongMarshallField(stage, "Sliders")
		pointersInitializesStatements += group.GongMarshallField(stage, "Checkboxes")
	}

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
	for _, layout := range layoutOrdered {

		identifiersDecl += layout.GongMarshallIdentifier(stage)

		initializerStatements += "\n"
		// Insertion point for basic fields value assignment
		initializerStatements += layout.GongMarshallField(stage, "Name")
		pointersInitializesStatements += layout.GongMarshallField(stage, "Groups")
	}

	sliderOrdered := []*Slider{}
	for slider := range stage.Sliders {
		sliderOrdered = append(sliderOrdered, slider)
	}
	sort.Slice(sliderOrdered[:], func(i, j int) bool {
		slideri := sliderOrdered[i]
		sliderj := sliderOrdered[j]
		slideri_order, oki := stage.SliderMap_Staged_Order[slideri]
		sliderj_order, okj := stage.SliderMap_Staged_Order[sliderj]
		if !oki || !okj {
			log.Fatalln("unknown pointers")
		}
		return slideri_order < sliderj_order
	})
	if len(sliderOrdered) > 0 {
		identifiersDecl += "\n"
	}
	for _, slider := range sliderOrdered {

		identifiersDecl += slider.GongMarshallIdentifier(stage)

		initializerStatements += "\n"
		// Insertion point for basic fields value assignment
		initializerStatements += slider.GongMarshallField(stage, "Name")
		initializerStatements += slider.GongMarshallField(stage, "IsFloat64")
		initializerStatements += slider.GongMarshallField(stage, "IsInt")
		initializerStatements += slider.GongMarshallField(stage, "MinInt")
		initializerStatements += slider.GongMarshallField(stage, "MaxInt")
		initializerStatements += slider.GongMarshallField(stage, "StepInt")
		initializerStatements += slider.GongMarshallField(stage, "ValueInt")
		initializerStatements += slider.GongMarshallField(stage, "MinFloat64")
		initializerStatements += slider.GongMarshallField(stage, "MaxFloat64")
		initializerStatements += slider.GongMarshallField(stage, "StepFloat64")
		initializerStatements += slider.GongMarshallField(stage, "ValueFloat64")
	}

	// insertion initialization of objects to stage
	for _, checkbox := range checkboxOrdered {
		_ = checkbox
		var setPointerField string
		_ = setPointerField

		// Insertion point for pointers initialization
	}

	for _, group := range groupOrdered {
		_ = group
		var setPointerField string
		_ = setPointerField

		// Insertion point for pointers initialization
	}

	for _, layout := range layoutOrdered {
		_ = layout
		var setPointerField string
		_ = setPointerField

		// Insertion point for pointers initialization
	}

	for _, slider := range sliderOrdered {
		_ = slider
		var setPointerField string
		_ = setPointerField

		// Insertion point for pointers initialization
	}

	res = strings.ReplaceAll(res, "{{Identifiers}}", identifiersDecl)
	res = strings.ReplaceAll(res, "{{ValueInitializers}}", initializerStatements)
	res = strings.ReplaceAll(res, "{{PointersInitializers}}", pointersInitializesStatements)

	if stage.MetaPackageImportAlias != "" {
		res = strings.ReplaceAll(res, "{{ImportPackageDeclaration}}",
			fmt.Sprintf("\n\t%s %s", stage.MetaPackageImportAlias, stage.MetaPackageImportPath))

		res = strings.ReplaceAll(res, "{{ImportPackageDummyDeclaration}}",
			fmt.Sprintf("\nvar _ %s.Stage",
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
	return
}

// insertion point for marshall field methods
func (checkbox *Checkbox) GongMarshallField(stage *Stage, fieldName string) (res string) {

	switch fieldName {
	case "Name":
		res = StringInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", checkbox.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "Name")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", string(checkbox.Name))
	case "ValueBool":
		res = NumberInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", checkbox.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "ValueBool")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%t", checkbox.ValueBool))
	case "LabelForTrue":
		res = StringInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", checkbox.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "LabelForTrue")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", string(checkbox.LabelForTrue))
	case "LabelForFalse":
		res = StringInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", checkbox.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "LabelForFalse")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", string(checkbox.LabelForFalse))

	default:
		log.Panicf("Unknown field %s for Gongstruct Checkbox", fieldName)
	}
	return
}

func (group *Group) GongMarshallField(stage *Stage, fieldName string) (res string) {

	switch fieldName {
	case "Name":
		res = StringInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", group.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "Name")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", string(group.Name))
	case "Percentage":
		res = NumberInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", group.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "Percentage")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%f", group.Percentage))

	case "Sliders":
		for _, _slider := range group.Sliders {
			tmp := SliceOfPointersFieldInitStatement
			tmp = strings.ReplaceAll(tmp, "{{Identifier}}", group.GongGetIdentifier(stage))
			tmp = strings.ReplaceAll(tmp, "{{GeneratedFieldName}}", "Sliders")
			tmp = strings.ReplaceAll(tmp, "{{GeneratedFieldNameValue}}", _slider.GongGetIdentifier(stage))
			res += tmp
		}
	case "Checkboxes":
		for _, _checkbox := range group.Checkboxes {
			tmp := SliceOfPointersFieldInitStatement
			tmp = strings.ReplaceAll(tmp, "{{Identifier}}", group.GongGetIdentifier(stage))
			tmp = strings.ReplaceAll(tmp, "{{GeneratedFieldName}}", "Checkboxes")
			tmp = strings.ReplaceAll(tmp, "{{GeneratedFieldNameValue}}", _checkbox.GongGetIdentifier(stage))
			res += tmp
		}
	default:
		log.Panicf("Unknown field %s for Gongstruct Group", fieldName)
	}
	return
}

func (layout *Layout) GongMarshallField(stage *Stage, fieldName string) (res string) {

	switch fieldName {
	case "Name":
		res = StringInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", layout.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "Name")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", string(layout.Name))

	case "Groups":
		for _, _group := range layout.Groups {
			tmp := SliceOfPointersFieldInitStatement
			tmp = strings.ReplaceAll(tmp, "{{Identifier}}", layout.GongGetIdentifier(stage))
			tmp = strings.ReplaceAll(tmp, "{{GeneratedFieldName}}", "Groups")
			tmp = strings.ReplaceAll(tmp, "{{GeneratedFieldNameValue}}", _group.GongGetIdentifier(stage))
			res += tmp
		}
	default:
		log.Panicf("Unknown field %s for Gongstruct Layout", fieldName)
	}
	return
}

func (slider *Slider) GongMarshallField(stage *Stage, fieldName string) (res string) {

	switch fieldName {
	case "Name":
		res = StringInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", slider.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "Name")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", string(slider.Name))
	case "IsFloat64":
		res = NumberInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", slider.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "IsFloat64")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%t", slider.IsFloat64))
	case "IsInt":
		res = NumberInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", slider.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "IsInt")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%t", slider.IsInt))
	case "MinInt":
		res = NumberInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", slider.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "MinInt")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%d", slider.MinInt))
	case "MaxInt":
		res = NumberInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", slider.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "MaxInt")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%d", slider.MaxInt))
	case "StepInt":
		res = NumberInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", slider.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "StepInt")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%d", slider.StepInt))
	case "ValueInt":
		res = NumberInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", slider.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "ValueInt")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%d", slider.ValueInt))
	case "MinFloat64":
		res = NumberInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", slider.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "MinFloat64")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%f", slider.MinFloat64))
	case "MaxFloat64":
		res = NumberInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", slider.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "MaxFloat64")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%f", slider.MaxFloat64))
	case "StepFloat64":
		res = NumberInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", slider.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "StepFloat64")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%f", slider.StepFloat64))
	case "ValueFloat64":
		res = NumberInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", slider.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "ValueFloat64")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%f", slider.ValueFloat64))

	default:
		log.Panicf("Unknown field %s for Gongstruct Slider", fieldName)
	}
	return
}

// insertion point for marshall all fields methods
func (checkbox *Checkbox) GongMarshallAllFields(stage *Stage) (initializerStatements string, pointersInitializesStatements string) {

	initializerStatements += "\n"
	pointersInitializesStatements += "\n"
	{ // Insertion point for basic fields value assignment
		initializerStatements += checkbox.GongMarshallField(stage, "Name")
		initializerStatements += checkbox.GongMarshallField(stage, "ValueBool")
		initializerStatements += checkbox.GongMarshallField(stage, "LabelForTrue")
		initializerStatements += checkbox.GongMarshallField(stage, "LabelForFalse")
	}
	return
}
func (group *Group) GongMarshallAllFields(stage *Stage) (initializerStatements string, pointersInitializesStatements string) {

	initializerStatements += "\n"
	pointersInitializesStatements += "\n"
	{ // Insertion point for basic fields value assignment
		initializerStatements += group.GongMarshallField(stage, "Name")
		initializerStatements += group.GongMarshallField(stage, "Percentage")
		pointersInitializesStatements += group.GongMarshallField(stage, "Sliders")
		pointersInitializesStatements += group.GongMarshallField(stage, "Checkboxes")
	}
	return
}
func (layout *Layout) GongMarshallAllFields(stage *Stage) (initializerStatements string, pointersInitializesStatements string) {

	initializerStatements += "\n"
	pointersInitializesStatements += "\n"
	{ // Insertion point for basic fields value assignment
		initializerStatements += layout.GongMarshallField(stage, "Name")
		pointersInitializesStatements += layout.GongMarshallField(stage, "Groups")
	}
	return
}
func (slider *Slider) GongMarshallAllFields(stage *Stage) (initializerStatements string, pointersInitializesStatements string) {

	initializerStatements += "\n"
	pointersInitializesStatements += "\n"
	{ // Insertion point for basic fields value assignment
		initializerStatements += slider.GongMarshallField(stage, "Name")
		initializerStatements += slider.GongMarshallField(stage, "IsFloat64")
		initializerStatements += slider.GongMarshallField(stage, "IsInt")
		initializerStatements += slider.GongMarshallField(stage, "MinInt")
		initializerStatements += slider.GongMarshallField(stage, "MaxInt")
		initializerStatements += slider.GongMarshallField(stage, "StepInt")
		initializerStatements += slider.GongMarshallField(stage, "ValueInt")
		initializerStatements += slider.GongMarshallField(stage, "MinFloat64")
		initializerStatements += slider.GongMarshallField(stage, "MaxFloat64")
		initializerStatements += slider.GongMarshallField(stage, "StepFloat64")
		initializerStatements += slider.GongMarshallField(stage, "ValueFloat64")
	}
	return
}