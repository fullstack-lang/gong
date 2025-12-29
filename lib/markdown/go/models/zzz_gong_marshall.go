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

	// Declaration of instances to stage{{Identifiers}}

	// Setup of values{{ValueInitializers}}

	// Setup of pointers{{PointersInitializers}}
}
`

const IdentifiersDecls = `
	{{Identifier}} := (&models.{{GeneratedStructName}}{Name: ` + "`" + `{{GeneratedFieldNameValue}}` + "`" + `}).Stage(stage)`

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
	contentOrdered := []*Content{}
	for content := range stage.Contents {
		contentOrdered = append(contentOrdered, content)
	}
	sort.Slice(contentOrdered[:], func(i, j int) bool {
		contenti := contentOrdered[i]
		contentj := contentOrdered[j]
		contenti_order, oki := stage.ContentMap_Staged_Order[contenti]
		contentj_order, okj := stage.ContentMap_Staged_Order[contentj]
		if !oki || !okj {
			log.Fatalln("unknown pointers")
		}
		return contenti_order < contentj_order
	})
	if len(contentOrdered) > 0 {
		identifiersDecl += "\n"
	}
	for _, content := range contentOrdered {

		decl = IdentifiersDecls
		decl = strings.ReplaceAll(decl, "{{Identifier}}", content.GongGetIdentifier(stage))
		decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "Content")
		decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", content.Name)
		identifiersDecl += decl

		initializerStatements += "\n"
		// Initialisation of values
		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", content.GongGetIdentifier(stage))
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Name")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(content.Name))
		initializerStatements += setValueField

		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", content.GongGetIdentifier(stage))
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Content")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(content.Content))
		initializerStatements += setValueField

	}

	jpgimageOrdered := []*JpgImage{}
	for jpgimage := range stage.JpgImages {
		jpgimageOrdered = append(jpgimageOrdered, jpgimage)
	}
	sort.Slice(jpgimageOrdered[:], func(i, j int) bool {
		jpgimagei := jpgimageOrdered[i]
		jpgimagej := jpgimageOrdered[j]
		jpgimagei_order, oki := stage.JpgImageMap_Staged_Order[jpgimagei]
		jpgimagej_order, okj := stage.JpgImageMap_Staged_Order[jpgimagej]
		if !oki || !okj {
			log.Fatalln("unknown pointers")
		}
		return jpgimagei_order < jpgimagej_order
	})
	if len(jpgimageOrdered) > 0 {
		identifiersDecl += "\n"
	}
	for _, jpgimage := range jpgimageOrdered {

		decl = IdentifiersDecls
		decl = strings.ReplaceAll(decl, "{{Identifier}}", jpgimage.GongGetIdentifier(stage))
		decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "JpgImage")
		decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", jpgimage.Name)
		identifiersDecl += decl

		initializerStatements += "\n"
		// Initialisation of values
		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", jpgimage.GongGetIdentifier(stage))
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Name")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(jpgimage.Name))
		initializerStatements += setValueField

		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", jpgimage.GongGetIdentifier(stage))
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Base64Content")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(jpgimage.Base64Content))
		initializerStatements += setValueField

	}

	pngimageOrdered := []*PngImage{}
	for pngimage := range stage.PngImages {
		pngimageOrdered = append(pngimageOrdered, pngimage)
	}
	sort.Slice(pngimageOrdered[:], func(i, j int) bool {
		pngimagei := pngimageOrdered[i]
		pngimagej := pngimageOrdered[j]
		pngimagei_order, oki := stage.PngImageMap_Staged_Order[pngimagei]
		pngimagej_order, okj := stage.PngImageMap_Staged_Order[pngimagej]
		if !oki || !okj {
			log.Fatalln("unknown pointers")
		}
		return pngimagei_order < pngimagej_order
	})
	if len(pngimageOrdered) > 0 {
		identifiersDecl += "\n"
	}
	for _, pngimage := range pngimageOrdered {

		decl = IdentifiersDecls
		decl = strings.ReplaceAll(decl, "{{Identifier}}", pngimage.GongGetIdentifier(stage))
		decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "PngImage")
		decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", pngimage.Name)
		identifiersDecl += decl

		initializerStatements += "\n"
		// Initialisation of values
		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", pngimage.GongGetIdentifier(stage))
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Name")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(pngimage.Name))
		initializerStatements += setValueField

		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", pngimage.GongGetIdentifier(stage))
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Base64Content")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(pngimage.Base64Content))
		initializerStatements += setValueField

	}

	svgimageOrdered := []*SvgImage{}
	for svgimage := range stage.SvgImages {
		svgimageOrdered = append(svgimageOrdered, svgimage)
	}
	sort.Slice(svgimageOrdered[:], func(i, j int) bool {
		svgimagei := svgimageOrdered[i]
		svgimagej := svgimageOrdered[j]
		svgimagei_order, oki := stage.SvgImageMap_Staged_Order[svgimagei]
		svgimagej_order, okj := stage.SvgImageMap_Staged_Order[svgimagej]
		if !oki || !okj {
			log.Fatalln("unknown pointers")
		}
		return svgimagei_order < svgimagej_order
	})
	if len(svgimageOrdered) > 0 {
		identifiersDecl += "\n"
	}
	for _, svgimage := range svgimageOrdered {

		decl = IdentifiersDecls
		decl = strings.ReplaceAll(decl, "{{Identifier}}", svgimage.GongGetIdentifier(stage))
		decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "SvgImage")
		decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", svgimage.Name)
		identifiersDecl += decl

		initializerStatements += "\n"
		// Initialisation of values
		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", svgimage.GongGetIdentifier(stage))
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Name")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(svgimage.Name))
		initializerStatements += setValueField

		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", svgimage.GongGetIdentifier(stage))
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Content")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(svgimage.Content))
		initializerStatements += setValueField

	}

	// insertion initialization of objects to stage
	if len(contentOrdered) > 0 {
		pointersInitializesStatements += "\n\t// setup of Content instances pointers"
	}
	for _, content := range contentOrdered {
		_ = content
		var setPointerField string
		_ = setPointerField

		// Initialisation of values
	}

	if len(jpgimageOrdered) > 0 {
		pointersInitializesStatements += "\n\t// setup of JpgImage instances pointers"
	}
	for _, jpgimage := range jpgimageOrdered {
		_ = jpgimage
		var setPointerField string
		_ = setPointerField

		// Initialisation of values
	}

	if len(pngimageOrdered) > 0 {
		pointersInitializesStatements += "\n\t// setup of PngImage instances pointers"
	}
	for _, pngimage := range pngimageOrdered {
		_ = pngimage
		var setPointerField string
		_ = setPointerField

		// Initialisation of values
	}

	if len(svgimageOrdered) > 0 {
		pointersInitializesStatements += "\n\t// setup of SvgImage instances pointers"
	}
	for _, svgimage := range svgimageOrdered {
		_ = svgimage
		var setPointerField string
		_ = setPointerField

		// Initialisation of values
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
