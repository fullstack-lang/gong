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

const GongIdentifiersDecls = `
	{{Identifier}} := (&models.{{GeneratedStructName}}{Name: ` + "`" + `{{GeneratedFieldNameValue}}` + "`" + `}).Stage(stage)`

const GongUnstageStmt = `
	{{Identifier}}.Unstage(stage)`

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
	for _, button := range buttonOrdered {

		identifiersDecl += button.GongMarshallIdentifier(stage)

		initializerStatements += "\n"
		// Insertion point for basic fields value assignment
		initializerStatements += button.GongMarshallField(stage, "Name")
		initializerStatements += button.GongMarshallField(stage, "Label")
		initializerStatements += button.GongMarshallField(stage, "Icon")
		initializerStatements += button.GongMarshallField(stage, "IsDisabled")
		initializerStatements += button.GongMarshallField(stage, "Color")
		initializerStatements += button.GongMarshallField(stage, "MatButtonType")
		initializerStatements += button.GongMarshallField(stage, "MatButtonAppearance")
	}

	buttontoggleOrdered := []*ButtonToggle{}
	for buttontoggle := range stage.ButtonToggles {
		buttontoggleOrdered = append(buttontoggleOrdered, buttontoggle)
	}
	sort.Slice(buttontoggleOrdered[:], func(i, j int) bool {
		buttontogglei := buttontoggleOrdered[i]
		buttontogglej := buttontoggleOrdered[j]
		buttontogglei_order, oki := stage.ButtonToggleMap_Staged_Order[buttontogglei]
		buttontogglej_order, okj := stage.ButtonToggleMap_Staged_Order[buttontogglej]
		if !oki || !okj {
			log.Fatalln("unknown pointers")
		}
		return buttontogglei_order < buttontogglej_order
	})
	if len(buttontoggleOrdered) > 0 {
		identifiersDecl += "\n"
	}
	for _, buttontoggle := range buttontoggleOrdered {

		identifiersDecl += buttontoggle.GongMarshallIdentifier(stage)

		initializerStatements += "\n"
		// Insertion point for basic fields value assignment
		initializerStatements += buttontoggle.GongMarshallField(stage, "Name")
		initializerStatements += buttontoggle.GongMarshallField(stage, "Label")
		initializerStatements += buttontoggle.GongMarshallField(stage, "Icon")
		initializerStatements += buttontoggle.GongMarshallField(stage, "IsDisabled")
		initializerStatements += buttontoggle.GongMarshallField(stage, "IsChecked")
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
		pointersInitializesStatements += group.GongMarshallField(stage, "Buttons")
		initializerStatements += group.GongMarshallField(stage, "NbColumns")
	}

	grouptoogleOrdered := []*GroupToogle{}
	for grouptoogle := range stage.GroupToogles {
		grouptoogleOrdered = append(grouptoogleOrdered, grouptoogle)
	}
	sort.Slice(grouptoogleOrdered[:], func(i, j int) bool {
		grouptooglei := grouptoogleOrdered[i]
		grouptooglej := grouptoogleOrdered[j]
		grouptooglei_order, oki := stage.GroupToogleMap_Staged_Order[grouptooglei]
		grouptooglej_order, okj := stage.GroupToogleMap_Staged_Order[grouptooglej]
		if !oki || !okj {
			log.Fatalln("unknown pointers")
		}
		return grouptooglei_order < grouptooglej_order
	})
	if len(grouptoogleOrdered) > 0 {
		identifiersDecl += "\n"
	}
	for _, grouptoogle := range grouptoogleOrdered {

		identifiersDecl += grouptoogle.GongMarshallIdentifier(stage)

		initializerStatements += "\n"
		// Insertion point for basic fields value assignment
		initializerStatements += grouptoogle.GongMarshallField(stage, "Name")
		initializerStatements += grouptoogle.GongMarshallField(stage, "Percentage")
		pointersInitializesStatements += grouptoogle.GongMarshallField(stage, "ButtonToggles")
		initializerStatements += grouptoogle.GongMarshallField(stage, "IsSingleSelector")
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
		pointersInitializesStatements += layout.GongMarshallField(stage, "GroupToogles")
	}

	// insertion initialization of objects to stage
	for _, button := range buttonOrdered {
		_ = button
		var setPointerField string
		_ = setPointerField

		// Insertion point for pointers initialization
	}

	for _, buttontoggle := range buttontoggleOrdered {
		_ = buttontoggle
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

	for _, grouptoogle := range grouptoogleOrdered {
		_ = grouptoogle
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
func (button *Button) GongMarshallField(stage *Stage, fieldName string) (res string) {

	switch fieldName {
	case "Name":
		res = StringInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", button.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "Name")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", string(button.Name))
	case "Label":
		res = StringInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", button.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "Label")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", string(button.Label))
	case "Icon":
		res = StringInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", button.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "Icon")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", string(button.Icon))
	case "IsDisabled":
		res = NumberInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", button.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "IsDisabled")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%t", button.IsDisabled))
	case "Color":
		if button.Color != "" {
			res = StringEnumInitStatement
			res = strings.ReplaceAll(res, "{{Identifier}}", button.GongGetIdentifier(stage))
			res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "Color")
			res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", "models."+button.Color.ToCodeString())
		}
	case "MatButtonType":
		if button.MatButtonType != "" {
			res = StringEnumInitStatement
			res = strings.ReplaceAll(res, "{{Identifier}}", button.GongGetIdentifier(stage))
			res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "MatButtonType")
			res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", "models."+button.MatButtonType.ToCodeString())
		}
	case "MatButtonAppearance":
		if button.MatButtonAppearance != "" {
			res = StringEnumInitStatement
			res = strings.ReplaceAll(res, "{{Identifier}}", button.GongGetIdentifier(stage))
			res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "MatButtonAppearance")
			res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", "models."+button.MatButtonAppearance.ToCodeString())
		}

	default:
		log.Panicf("Unknown field %s for Gongstruct Button", fieldName)
	}
	return
}

func (buttontoggle *ButtonToggle) GongMarshallField(stage *Stage, fieldName string) (res string) {

	switch fieldName {
	case "Name":
		res = StringInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", buttontoggle.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "Name")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", string(buttontoggle.Name))
	case "Label":
		res = StringInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", buttontoggle.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "Label")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", string(buttontoggle.Label))
	case "Icon":
		res = StringInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", buttontoggle.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "Icon")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", string(buttontoggle.Icon))
	case "IsDisabled":
		res = NumberInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", buttontoggle.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "IsDisabled")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%t", buttontoggle.IsDisabled))
	case "IsChecked":
		res = NumberInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", buttontoggle.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "IsChecked")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%t", buttontoggle.IsChecked))

	default:
		log.Panicf("Unknown field %s for Gongstruct ButtonToggle", fieldName)
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
	case "NbColumns":
		res = NumberInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", group.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "NbColumns")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%d", group.NbColumns))

	case "Buttons":
		for _, _button := range group.Buttons {
			tmp := SliceOfPointersFieldInitStatement
			tmp = strings.ReplaceAll(tmp, "{{Identifier}}", group.GongGetIdentifier(stage))
			tmp = strings.ReplaceAll(tmp, "{{GeneratedFieldName}}", "Buttons")
			tmp = strings.ReplaceAll(tmp, "{{GeneratedFieldNameValue}}", _button.GongGetIdentifier(stage))
			res += tmp
		}
	default:
		log.Panicf("Unknown field %s for Gongstruct Group", fieldName)
	}
	return
}

func (grouptoogle *GroupToogle) GongMarshallField(stage *Stage, fieldName string) (res string) {

	switch fieldName {
	case "Name":
		res = StringInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", grouptoogle.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "Name")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", string(grouptoogle.Name))
	case "Percentage":
		res = NumberInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", grouptoogle.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "Percentage")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%f", grouptoogle.Percentage))
	case "IsSingleSelector":
		res = NumberInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", grouptoogle.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "IsSingleSelector")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%t", grouptoogle.IsSingleSelector))

	case "ButtonToggles":
		for _, _buttontoggle := range grouptoogle.ButtonToggles {
			tmp := SliceOfPointersFieldInitStatement
			tmp = strings.ReplaceAll(tmp, "{{Identifier}}", grouptoogle.GongGetIdentifier(stage))
			tmp = strings.ReplaceAll(tmp, "{{GeneratedFieldName}}", "ButtonToggles")
			tmp = strings.ReplaceAll(tmp, "{{GeneratedFieldNameValue}}", _buttontoggle.GongGetIdentifier(stage))
			res += tmp
		}
	default:
		log.Panicf("Unknown field %s for Gongstruct GroupToogle", fieldName)
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
	case "GroupToogles":
		for _, _grouptoogle := range layout.GroupToogles {
			tmp := SliceOfPointersFieldInitStatement
			tmp = strings.ReplaceAll(tmp, "{{Identifier}}", layout.GongGetIdentifier(stage))
			tmp = strings.ReplaceAll(tmp, "{{GeneratedFieldName}}", "GroupToogles")
			tmp = strings.ReplaceAll(tmp, "{{GeneratedFieldNameValue}}", _grouptoogle.GongGetIdentifier(stage))
			res += tmp
		}
	default:
		log.Panicf("Unknown field %s for Gongstruct Layout", fieldName)
	}
	return
}

// insertion point for marshall all fields methods
func (button *Button) GongMarshallAllFields(stage *Stage) (initializerStatements string, pointersInitializesStatements string) {

	initializerStatements += "\n"
	pointersInitializesStatements += "\n"
	{ // Insertion point for basic fields value assignment
		initializerStatements += button.GongMarshallField(stage, "Name")
		initializerStatements += button.GongMarshallField(stage, "Label")
		initializerStatements += button.GongMarshallField(stage, "Icon")
		initializerStatements += button.GongMarshallField(stage, "IsDisabled")
		initializerStatements += button.GongMarshallField(stage, "Color")
		initializerStatements += button.GongMarshallField(stage, "MatButtonType")
		initializerStatements += button.GongMarshallField(stage, "MatButtonAppearance")
	}
	return
}
func (buttontoggle *ButtonToggle) GongMarshallAllFields(stage *Stage) (initializerStatements string, pointersInitializesStatements string) {

	initializerStatements += "\n"
	pointersInitializesStatements += "\n"
	{ // Insertion point for basic fields value assignment
		initializerStatements += buttontoggle.GongMarshallField(stage, "Name")
		initializerStatements += buttontoggle.GongMarshallField(stage, "Label")
		initializerStatements += buttontoggle.GongMarshallField(stage, "Icon")
		initializerStatements += buttontoggle.GongMarshallField(stage, "IsDisabled")
		initializerStatements += buttontoggle.GongMarshallField(stage, "IsChecked")
	}
	return
}
func (group *Group) GongMarshallAllFields(stage *Stage) (initializerStatements string, pointersInitializesStatements string) {

	initializerStatements += "\n"
	pointersInitializesStatements += "\n"
	{ // Insertion point for basic fields value assignment
		initializerStatements += group.GongMarshallField(stage, "Name")
		initializerStatements += group.GongMarshallField(stage, "Percentage")
		pointersInitializesStatements += group.GongMarshallField(stage, "Buttons")
		initializerStatements += group.GongMarshallField(stage, "NbColumns")
	}
	return
}
func (grouptoogle *GroupToogle) GongMarshallAllFields(stage *Stage) (initializerStatements string, pointersInitializesStatements string) {

	initializerStatements += "\n"
	pointersInitializesStatements += "\n"
	{ // Insertion point for basic fields value assignment
		initializerStatements += grouptoogle.GongMarshallField(stage, "Name")
		initializerStatements += grouptoogle.GongMarshallField(stage, "Percentage")
		pointersInitializesStatements += grouptoogle.GongMarshallField(stage, "ButtonToggles")
		initializerStatements += grouptoogle.GongMarshallField(stage, "IsSingleSelector")
	}
	return
}
func (layout *Layout) GongMarshallAllFields(stage *Stage) (initializerStatements string, pointersInitializesStatements string) {

	initializerStatements += "\n"
	pointersInitializesStatements += "\n"
	{ // Insertion point for basic fields value assignment
		initializerStatements += layout.GongMarshallField(stage, "Name")
		pointersInitializesStatements += layout.GongMarshallField(stage, "Groups")
		pointersInitializesStatements += layout.GongMarshallField(stage, "GroupToogles")
	}
	return
}