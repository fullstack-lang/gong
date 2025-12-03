// generated code - do not edit
package models

import (
	"fmt"
	"log"
	"os"
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
func _(stage *models.Stage) {

	// Declaration of instances to stage{{Identifiers}}

	// Setup of values{{ValueInitializers}}

	// Setup of pointers{{PointersInitializers}}
}
`

const IdentifiersDecls = `
	{{Identifier}} := (&models.{{GeneratedStructName}}{}).Stage(stage)`

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

	id := ""
	_ = id
	decl := ""
	_ = decl
	setValueField := ""
	_ = setValueField

	// insertion initialization of objects to stage
	map_Content_Identifiers := make(map[*Content]string)
	_ = map_Content_Identifiers

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
	for idx, content := range contentOrdered {

		id = generatesIdentifier("Content", idx, content.Name)
		map_Content_Identifiers[content] = id

		decl = IdentifiersDecls
		decl = strings.ReplaceAll(decl, "{{Identifier}}", id)
		decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "Content")
		decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", content.Name)
		identifiersDecl += decl

		initializerStatements += "\n"
		// Initialisation of values
		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Name")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(content.Name))
		initializerStatements += setValueField

		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Content")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(content.Content))
		initializerStatements += setValueField

	}

	map_JpgImage_Identifiers := make(map[*JpgImage]string)
	_ = map_JpgImage_Identifiers

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
	for idx, jpgimage := range jpgimageOrdered {

		id = generatesIdentifier("JpgImage", idx, jpgimage.Name)
		map_JpgImage_Identifiers[jpgimage] = id

		decl = IdentifiersDecls
		decl = strings.ReplaceAll(decl, "{{Identifier}}", id)
		decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "JpgImage")
		decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", jpgimage.Name)
		identifiersDecl += decl

		initializerStatements += "\n"
		// Initialisation of values
		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Name")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(jpgimage.Name))
		initializerStatements += setValueField

		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Base64Content")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(jpgimage.Base64Content))
		initializerStatements += setValueField

	}

	map_PngImage_Identifiers := make(map[*PngImage]string)
	_ = map_PngImage_Identifiers

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
	for idx, pngimage := range pngimageOrdered {

		id = generatesIdentifier("PngImage", idx, pngimage.Name)
		map_PngImage_Identifiers[pngimage] = id

		decl = IdentifiersDecls
		decl = strings.ReplaceAll(decl, "{{Identifier}}", id)
		decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "PngImage")
		decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", pngimage.Name)
		identifiersDecl += decl

		initializerStatements += "\n"
		// Initialisation of values
		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Name")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(pngimage.Name))
		initializerStatements += setValueField

		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Base64Content")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(pngimage.Base64Content))
		initializerStatements += setValueField

	}

	map_SvgImage_Identifiers := make(map[*SvgImage]string)
	_ = map_SvgImage_Identifiers

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
	for idx, svgimage := range svgimageOrdered {

		id = generatesIdentifier("SvgImage", idx, svgimage.Name)
		map_SvgImage_Identifiers[svgimage] = id

		decl = IdentifiersDecls
		decl = strings.ReplaceAll(decl, "{{Identifier}}", id)
		decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "SvgImage")
		decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", svgimage.Name)
		identifiersDecl += decl

		initializerStatements += "\n"
		// Initialisation of values
		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Name")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(svgimage.Name))
		initializerStatements += setValueField

		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Content")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(svgimage.Content))
		initializerStatements += setValueField

	}

	// insertion initialization of objects to stage
	if len(contentOrdered) > 0 {
		pointersInitializesStatements += "\n\t// setup of Content instances pointers"
	}
	for idx, content := range contentOrdered {
		var setPointerField string
		_ = setPointerField

		id = generatesIdentifier("Content", idx, content.Name)
		map_Content_Identifiers[content] = id

		// Initialisation of values
	}

	if len(jpgimageOrdered) > 0 {
		pointersInitializesStatements += "\n\t// setup of JpgImage instances pointers"
	}
	for idx, jpgimage := range jpgimageOrdered {
		var setPointerField string
		_ = setPointerField

		id = generatesIdentifier("JpgImage", idx, jpgimage.Name)
		map_JpgImage_Identifiers[jpgimage] = id

		// Initialisation of values
	}

	if len(pngimageOrdered) > 0 {
		pointersInitializesStatements += "\n\t// setup of PngImage instances pointers"
	}
	for idx, pngimage := range pngimageOrdered {
		var setPointerField string
		_ = setPointerField

		id = generatesIdentifier("PngImage", idx, pngimage.Name)
		map_PngImage_Identifiers[pngimage] = id

		// Initialisation of values
	}

	if len(svgimageOrdered) > 0 {
		pointersInitializesStatements += "\n\t// setup of SvgImage instances pointers"
	}
	for idx, svgimage := range svgimageOrdered {
		var setPointerField string
		_ = setPointerField

		id = generatesIdentifier("SvgImage", idx, svgimage.Name)
		map_SvgImage_Identifiers[svgimage] = id

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
