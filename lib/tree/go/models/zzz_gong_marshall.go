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
		initializerStatements += button.GongMarshallField(stage, "Icon")
		pointersInitializesStatements += button.GongMarshallField(stage, "SVGIcon")
		initializerStatements += button.GongMarshallField(stage, "IsDisabled")
		initializerStatements += button.GongMarshallField(stage, "HasToolTip")
		initializerStatements += button.GongMarshallField(stage, "ToolTipText")
		initializerStatements += button.GongMarshallField(stage, "ToolTipPosition")
	}

	nodeOrdered := []*Node{}
	for node := range stage.Nodes {
		nodeOrdered = append(nodeOrdered, node)
	}
	sort.Slice(nodeOrdered[:], func(i, j int) bool {
		nodei := nodeOrdered[i]
		nodej := nodeOrdered[j]
		nodei_order, oki := stage.NodeMap_Staged_Order[nodei]
		nodej_order, okj := stage.NodeMap_Staged_Order[nodej]
		if !oki || !okj {
			log.Fatalln("unknown pointers")
		}
		return nodei_order < nodej_order
	})
	if len(nodeOrdered) > 0 {
		identifiersDecl += "\n"
	}
	for _, node := range nodeOrdered {

		identifiersDecl += node.GongMarshallIdentifier(stage)

		initializerStatements += "\n"
		// Insertion point for basic fields value assignment
		initializerStatements += node.GongMarshallField(stage, "Name")
		initializerStatements += node.GongMarshallField(stage, "FontStyle")
		initializerStatements += node.GongMarshallField(stage, "BackgroundColor")
		initializerStatements += node.GongMarshallField(stage, "IsExpanded")
		initializerStatements += node.GongMarshallField(stage, "HasCheckboxButton")
		initializerStatements += node.GongMarshallField(stage, "IsChecked")
		initializerStatements += node.GongMarshallField(stage, "IsCheckboxDisabled")
		initializerStatements += node.GongMarshallField(stage, "CheckboxHasToolTip")
		initializerStatements += node.GongMarshallField(stage, "CheckboxToolTipText")
		initializerStatements += node.GongMarshallField(stage, "CheckboxToolTipPosition")
		initializerStatements += node.GongMarshallField(stage, "HasSecondCheckboxButton")
		initializerStatements += node.GongMarshallField(stage, "IsSecondCheckboxChecked")
		initializerStatements += node.GongMarshallField(stage, "IsSecondCheckboxDisabled")
		initializerStatements += node.GongMarshallField(stage, "SecondCheckboxHasToolTip")
		initializerStatements += node.GongMarshallField(stage, "SecondCheckboxToolTipText")
		initializerStatements += node.GongMarshallField(stage, "SecondCheckboxToolTipPosition")
		initializerStatements += node.GongMarshallField(stage, "TextAfterSecondCheckbox")
		initializerStatements += node.GongMarshallField(stage, "HasToolTip")
		initializerStatements += node.GongMarshallField(stage, "ToolTipText")
		initializerStatements += node.GongMarshallField(stage, "ToolTipPosition")
		initializerStatements += node.GongMarshallField(stage, "IsInEditMode")
		initializerStatements += node.GongMarshallField(stage, "IsNodeClickable")
		initializerStatements += node.GongMarshallField(stage, "IsWithPreceedingIcon")
		initializerStatements += node.GongMarshallField(stage, "PreceedingIcon")
		pointersInitializesStatements += node.GongMarshallField(stage, "PreceedingSVGIcon")
		pointersInitializesStatements += node.GongMarshallField(stage, "Children")
		pointersInitializesStatements += node.GongMarshallField(stage, "Buttons")
	}

	svgiconOrdered := []*SVGIcon{}
	for svgicon := range stage.SVGIcons {
		svgiconOrdered = append(svgiconOrdered, svgicon)
	}
	sort.Slice(svgiconOrdered[:], func(i, j int) bool {
		svgiconi := svgiconOrdered[i]
		svgiconj := svgiconOrdered[j]
		svgiconi_order, oki := stage.SVGIconMap_Staged_Order[svgiconi]
		svgiconj_order, okj := stage.SVGIconMap_Staged_Order[svgiconj]
		if !oki || !okj {
			log.Fatalln("unknown pointers")
		}
		return svgiconi_order < svgiconj_order
	})
	if len(svgiconOrdered) > 0 {
		identifiersDecl += "\n"
	}
	for _, svgicon := range svgiconOrdered {

		identifiersDecl += svgicon.GongMarshallIdentifier(stage)

		initializerStatements += "\n"
		// Insertion point for basic fields value assignment
		initializerStatements += svgicon.GongMarshallField(stage, "Name")
		initializerStatements += svgicon.GongMarshallField(stage, "SVG")
	}

	treeOrdered := []*Tree{}
	for tree := range stage.Trees {
		treeOrdered = append(treeOrdered, tree)
	}
	sort.Slice(treeOrdered[:], func(i, j int) bool {
		treei := treeOrdered[i]
		treej := treeOrdered[j]
		treei_order, oki := stage.TreeMap_Staged_Order[treei]
		treej_order, okj := stage.TreeMap_Staged_Order[treej]
		if !oki || !okj {
			log.Fatalln("unknown pointers")
		}
		return treei_order < treej_order
	})
	if len(treeOrdered) > 0 {
		identifiersDecl += "\n"
	}
	for _, tree := range treeOrdered {

		identifiersDecl += tree.GongMarshallIdentifier(stage)

		initializerStatements += "\n"
		// Insertion point for basic fields value assignment
		initializerStatements += tree.GongMarshallField(stage, "Name")
		pointersInitializesStatements += tree.GongMarshallField(stage, "RootNodes")
	}

	// insertion initialization of objects to stage
	for _, button := range buttonOrdered {
		_ = button
		var setPointerField string
		_ = setPointerField

		// Insertion point for pointers initialization
	}

	for _, node := range nodeOrdered {
		_ = node
		var setPointerField string
		_ = setPointerField

		// Insertion point for pointers initialization
	}

	for _, svgicon := range svgiconOrdered {
		_ = svgicon
		var setPointerField string
		_ = setPointerField

		// Insertion point for pointers initialization
	}

	for _, tree := range treeOrdered {
		_ = tree
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
	case "HasToolTip":
		res = NumberInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", button.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "HasToolTip")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%t", button.HasToolTip))
	case "ToolTipText":
		res = StringInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", button.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "ToolTipText")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", string(button.ToolTipText))
	case "ToolTipPosition":
		if button.ToolTipPosition != "" {
			res = StringEnumInitStatement
			res = strings.ReplaceAll(res, "{{Identifier}}", button.GongGetIdentifier(stage))
			res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "ToolTipPosition")
			res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", "models."+button.ToolTipPosition.ToCodeString())
		} else {
			// in case of empty enum, we need to unstage the previous value
			res = StringEnumInitStatement
			res = strings.ReplaceAll(res, "{{Identifier}}", button.GongGetIdentifier(stage))
			res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "ToolTipPosition")
			res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", "\"\"")
		}

	case "SVGIcon":
		if button.SVGIcon != nil {
			res = PointerFieldInitStatement
			res = strings.ReplaceAll(res, "{{Identifier}}", button.GongGetIdentifier(stage))
			res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "SVGIcon")
			res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", button.SVGIcon.GongGetIdentifier(stage))
		} else {
			// in case of nil pointer, we need to unstage the previous value
			res = PointerFieldInitStatement
			res = strings.ReplaceAll(res, "{{Identifier}}", button.GongGetIdentifier(stage))
			res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "SVGIcon")
			res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", "nil")
		}
	default:
		log.Panicf("Unknown field %s for Gongstruct Button", fieldName)
	}
	return
}

func (node *Node) GongMarshallField(stage *Stage, fieldName string) (res string) {

	switch fieldName {
	case "Name":
		res = StringInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", node.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "Name")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", string(node.Name))
	case "FontStyle":
		if node.FontStyle != "" {
			res = StringEnumInitStatement
			res = strings.ReplaceAll(res, "{{Identifier}}", node.GongGetIdentifier(stage))
			res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "FontStyle")
			res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", "models."+node.FontStyle.ToCodeString())
		} else {
			// in case of empty enum, we need to unstage the previous value
			res = StringEnumInitStatement
			res = strings.ReplaceAll(res, "{{Identifier}}", node.GongGetIdentifier(stage))
			res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "FontStyle")
			res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", "\"\"")
		}
	case "BackgroundColor":
		res = StringInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", node.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "BackgroundColor")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", string(node.BackgroundColor))
	case "IsExpanded":
		res = NumberInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", node.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "IsExpanded")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%t", node.IsExpanded))
	case "HasCheckboxButton":
		res = NumberInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", node.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "HasCheckboxButton")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%t", node.HasCheckboxButton))
	case "IsChecked":
		res = NumberInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", node.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "IsChecked")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%t", node.IsChecked))
	case "IsCheckboxDisabled":
		res = NumberInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", node.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "IsCheckboxDisabled")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%t", node.IsCheckboxDisabled))
	case "CheckboxHasToolTip":
		res = NumberInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", node.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "CheckboxHasToolTip")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%t", node.CheckboxHasToolTip))
	case "CheckboxToolTipText":
		res = StringInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", node.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "CheckboxToolTipText")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", string(node.CheckboxToolTipText))
	case "CheckboxToolTipPosition":
		if node.CheckboxToolTipPosition != "" {
			res = StringEnumInitStatement
			res = strings.ReplaceAll(res, "{{Identifier}}", node.GongGetIdentifier(stage))
			res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "CheckboxToolTipPosition")
			res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", "models."+node.CheckboxToolTipPosition.ToCodeString())
		} else {
			// in case of empty enum, we need to unstage the previous value
			res = StringEnumInitStatement
			res = strings.ReplaceAll(res, "{{Identifier}}", node.GongGetIdentifier(stage))
			res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "CheckboxToolTipPosition")
			res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", "\"\"")
		}
	case "HasSecondCheckboxButton":
		res = NumberInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", node.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "HasSecondCheckboxButton")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%t", node.HasSecondCheckboxButton))
	case "IsSecondCheckboxChecked":
		res = NumberInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", node.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "IsSecondCheckboxChecked")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%t", node.IsSecondCheckboxChecked))
	case "IsSecondCheckboxDisabled":
		res = NumberInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", node.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "IsSecondCheckboxDisabled")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%t", node.IsSecondCheckboxDisabled))
	case "SecondCheckboxHasToolTip":
		res = NumberInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", node.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "SecondCheckboxHasToolTip")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%t", node.SecondCheckboxHasToolTip))
	case "SecondCheckboxToolTipText":
		res = StringInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", node.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "SecondCheckboxToolTipText")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", string(node.SecondCheckboxToolTipText))
	case "SecondCheckboxToolTipPosition":
		if node.SecondCheckboxToolTipPosition != "" {
			res = StringEnumInitStatement
			res = strings.ReplaceAll(res, "{{Identifier}}", node.GongGetIdentifier(stage))
			res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "SecondCheckboxToolTipPosition")
			res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", "models."+node.SecondCheckboxToolTipPosition.ToCodeString())
		} else {
			// in case of empty enum, we need to unstage the previous value
			res = StringEnumInitStatement
			res = strings.ReplaceAll(res, "{{Identifier}}", node.GongGetIdentifier(stage))
			res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "SecondCheckboxToolTipPosition")
			res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", "\"\"")
		}
	case "TextAfterSecondCheckbox":
		res = StringInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", node.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "TextAfterSecondCheckbox")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", string(node.TextAfterSecondCheckbox))
	case "HasToolTip":
		res = NumberInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", node.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "HasToolTip")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%t", node.HasToolTip))
	case "ToolTipText":
		res = StringInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", node.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "ToolTipText")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", string(node.ToolTipText))
	case "ToolTipPosition":
		if node.ToolTipPosition != "" {
			res = StringEnumInitStatement
			res = strings.ReplaceAll(res, "{{Identifier}}", node.GongGetIdentifier(stage))
			res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "ToolTipPosition")
			res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", "models."+node.ToolTipPosition.ToCodeString())
		} else {
			// in case of empty enum, we need to unstage the previous value
			res = StringEnumInitStatement
			res = strings.ReplaceAll(res, "{{Identifier}}", node.GongGetIdentifier(stage))
			res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "ToolTipPosition")
			res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", "\"\"")
		}
	case "IsInEditMode":
		res = NumberInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", node.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "IsInEditMode")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%t", node.IsInEditMode))
	case "IsNodeClickable":
		res = NumberInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", node.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "IsNodeClickable")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%t", node.IsNodeClickable))
	case "IsWithPreceedingIcon":
		res = NumberInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", node.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "IsWithPreceedingIcon")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%t", node.IsWithPreceedingIcon))
	case "PreceedingIcon":
		res = StringInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", node.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "PreceedingIcon")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", string(node.PreceedingIcon))

	case "PreceedingSVGIcon":
		if node.PreceedingSVGIcon != nil {
			res = PointerFieldInitStatement
			res = strings.ReplaceAll(res, "{{Identifier}}", node.GongGetIdentifier(stage))
			res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "PreceedingSVGIcon")
			res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", node.PreceedingSVGIcon.GongGetIdentifier(stage))
		} else {
			// in case of nil pointer, we need to unstage the previous value
			res = PointerFieldInitStatement
			res = strings.ReplaceAll(res, "{{Identifier}}", node.GongGetIdentifier(stage))
			res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "PreceedingSVGIcon")
			res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", "nil")
		}
	case "Children":
		for _, _node := range node.Children {
			tmp := SliceOfPointersFieldInitStatement
			tmp = strings.ReplaceAll(tmp, "{{Identifier}}", node.GongGetIdentifier(stage))
			tmp = strings.ReplaceAll(tmp, "{{GeneratedFieldName}}", "Children")
			tmp = strings.ReplaceAll(tmp, "{{GeneratedFieldNameValue}}", _node.GongGetIdentifier(stage))
			res += tmp
		}
	case "Buttons":
		for _, _button := range node.Buttons {
			tmp := SliceOfPointersFieldInitStatement
			tmp = strings.ReplaceAll(tmp, "{{Identifier}}", node.GongGetIdentifier(stage))
			tmp = strings.ReplaceAll(tmp, "{{GeneratedFieldName}}", "Buttons")
			tmp = strings.ReplaceAll(tmp, "{{GeneratedFieldNameValue}}", _button.GongGetIdentifier(stage))
			res += tmp
		}
	default:
		log.Panicf("Unknown field %s for Gongstruct Node", fieldName)
	}
	return
}

func (svgicon *SVGIcon) GongMarshallField(stage *Stage, fieldName string) (res string) {

	switch fieldName {
	case "Name":
		res = StringInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", svgicon.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "Name")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", string(svgicon.Name))
	case "SVG":
		res = StringInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", svgicon.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "SVG")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", string(svgicon.SVG))

	default:
		log.Panicf("Unknown field %s for Gongstruct SVGIcon", fieldName)
	}
	return
}

func (tree *Tree) GongMarshallField(stage *Stage, fieldName string) (res string) {

	switch fieldName {
	case "Name":
		res = StringInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", tree.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "Name")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", string(tree.Name))

	case "RootNodes":
		for _, _node := range tree.RootNodes {
			tmp := SliceOfPointersFieldInitStatement
			tmp = strings.ReplaceAll(tmp, "{{Identifier}}", tree.GongGetIdentifier(stage))
			tmp = strings.ReplaceAll(tmp, "{{GeneratedFieldName}}", "RootNodes")
			tmp = strings.ReplaceAll(tmp, "{{GeneratedFieldNameValue}}", _node.GongGetIdentifier(stage))
			res += tmp
		}
	default:
		log.Panicf("Unknown field %s for Gongstruct Tree", fieldName)
	}
	return
}

// insertion point for marshall all fields methods
func (button *Button) GongMarshallAllFields(stage *Stage) (initializerStatements string, pointersInitializesStatements string) {

	{ // Insertion point for basic fields value assignment
		initializerStatements += button.GongMarshallField(stage, "Name")
		initializerStatements += button.GongMarshallField(stage, "Icon")
		pointersInitializesStatements += button.GongMarshallField(stage, "SVGIcon")
		initializerStatements += button.GongMarshallField(stage, "IsDisabled")
		initializerStatements += button.GongMarshallField(stage, "HasToolTip")
		initializerStatements += button.GongMarshallField(stage, "ToolTipText")
		initializerStatements += button.GongMarshallField(stage, "ToolTipPosition")
	}
	return
}
func (node *Node) GongMarshallAllFields(stage *Stage) (initializerStatements string, pointersInitializesStatements string) {

	{ // Insertion point for basic fields value assignment
		initializerStatements += node.GongMarshallField(stage, "Name")
		initializerStatements += node.GongMarshallField(stage, "FontStyle")
		initializerStatements += node.GongMarshallField(stage, "BackgroundColor")
		initializerStatements += node.GongMarshallField(stage, "IsExpanded")
		initializerStatements += node.GongMarshallField(stage, "HasCheckboxButton")
		initializerStatements += node.GongMarshallField(stage, "IsChecked")
		initializerStatements += node.GongMarshallField(stage, "IsCheckboxDisabled")
		initializerStatements += node.GongMarshallField(stage, "CheckboxHasToolTip")
		initializerStatements += node.GongMarshallField(stage, "CheckboxToolTipText")
		initializerStatements += node.GongMarshallField(stage, "CheckboxToolTipPosition")
		initializerStatements += node.GongMarshallField(stage, "HasSecondCheckboxButton")
		initializerStatements += node.GongMarshallField(stage, "IsSecondCheckboxChecked")
		initializerStatements += node.GongMarshallField(stage, "IsSecondCheckboxDisabled")
		initializerStatements += node.GongMarshallField(stage, "SecondCheckboxHasToolTip")
		initializerStatements += node.GongMarshallField(stage, "SecondCheckboxToolTipText")
		initializerStatements += node.GongMarshallField(stage, "SecondCheckboxToolTipPosition")
		initializerStatements += node.GongMarshallField(stage, "TextAfterSecondCheckbox")
		initializerStatements += node.GongMarshallField(stage, "HasToolTip")
		initializerStatements += node.GongMarshallField(stage, "ToolTipText")
		initializerStatements += node.GongMarshallField(stage, "ToolTipPosition")
		initializerStatements += node.GongMarshallField(stage, "IsInEditMode")
		initializerStatements += node.GongMarshallField(stage, "IsNodeClickable")
		initializerStatements += node.GongMarshallField(stage, "IsWithPreceedingIcon")
		initializerStatements += node.GongMarshallField(stage, "PreceedingIcon")
		pointersInitializesStatements += node.GongMarshallField(stage, "PreceedingSVGIcon")
		pointersInitializesStatements += node.GongMarshallField(stage, "Children")
		pointersInitializesStatements += node.GongMarshallField(stage, "Buttons")
	}
	return
}
func (svgicon *SVGIcon) GongMarshallAllFields(stage *Stage) (initializerStatements string, pointersInitializesStatements string) {

	{ // Insertion point for basic fields value assignment
		initializerStatements += svgicon.GongMarshallField(stage, "Name")
		initializerStatements += svgicon.GongMarshallField(stage, "SVG")
	}
	return
}
func (tree *Tree) GongMarshallAllFields(stage *Stage) (initializerStatements string, pointersInitializesStatements string) {

	{ // Insertion point for basic fields value assignment
		initializerStatements += tree.GongMarshallField(stage, "Name")
		pointersInitializesStatements += tree.GongMarshallField(stage, "RootNodes")
	}
	return
}
