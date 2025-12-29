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

	id := ""
	_ = id
	decl := ""
	_ = decl
	setValueField := ""
	_ = setValueField

	// insertion initialization of objects to stage
	map_Body_Identifiers := make(map[*Body]string)
	_ = map_Body_Identifiers

	bodyOrdered := []*Body{}
	for body := range stage.Bodys {
		bodyOrdered = append(bodyOrdered, body)
	}
	sort.Slice(bodyOrdered[:], func(i, j int) bool {
		bodyi := bodyOrdered[i]
		bodyj := bodyOrdered[j]
		bodyi_order, oki := stage.BodyMap_Staged_Order[bodyi]
		bodyj_order, okj := stage.BodyMap_Staged_Order[bodyj]
		if !oki || !okj {
			log.Fatalln("unknown pointers")
		}
		return bodyi_order < bodyj_order
	})
	if len(bodyOrdered) > 0 {
		identifiersDecl += "\n"
	}
	for _, body := range bodyOrdered {

		id = generatesIdentifier("Body", int(stage.BodyMap_Staged_Order[body]), body.Name)
		map_Body_Identifiers[body] = id

		decl = IdentifiersDecls
		decl = strings.ReplaceAll(decl, "{{Identifier}}", id)
		decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "Body")
		decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", body.Name)
		identifiersDecl += decl

		initializerStatements += "\n"
		// Initialisation of values
		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Name")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(body.Name))
		initializerStatements += setValueField

	}

	map_Document_Identifiers := make(map[*Document]string)
	_ = map_Document_Identifiers

	documentOrdered := []*Document{}
	for document := range stage.Documents {
		documentOrdered = append(documentOrdered, document)
	}
	sort.Slice(documentOrdered[:], func(i, j int) bool {
		documenti := documentOrdered[i]
		documentj := documentOrdered[j]
		documenti_order, oki := stage.DocumentMap_Staged_Order[documenti]
		documentj_order, okj := stage.DocumentMap_Staged_Order[documentj]
		if !oki || !okj {
			log.Fatalln("unknown pointers")
		}
		return documenti_order < documentj_order
	})
	if len(documentOrdered) > 0 {
		identifiersDecl += "\n"
	}
	for _, document := range documentOrdered {

		id = generatesIdentifier("Document", int(stage.DocumentMap_Staged_Order[document]), document.Name)
		map_Document_Identifiers[document] = id

		decl = IdentifiersDecls
		decl = strings.ReplaceAll(decl, "{{Identifier}}", id)
		decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "Document")
		decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", document.Name)
		identifiersDecl += decl

		initializerStatements += "\n"
		// Initialisation of values
		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Name")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(document.Name))
		initializerStatements += setValueField

	}

	map_Docx_Identifiers := make(map[*Docx]string)
	_ = map_Docx_Identifiers

	docxOrdered := []*Docx{}
	for docx := range stage.Docxs {
		docxOrdered = append(docxOrdered, docx)
	}
	sort.Slice(docxOrdered[:], func(i, j int) bool {
		docxi := docxOrdered[i]
		docxj := docxOrdered[j]
		docxi_order, oki := stage.DocxMap_Staged_Order[docxi]
		docxj_order, okj := stage.DocxMap_Staged_Order[docxj]
		if !oki || !okj {
			log.Fatalln("unknown pointers")
		}
		return docxi_order < docxj_order
	})
	if len(docxOrdered) > 0 {
		identifiersDecl += "\n"
	}
	for _, docx := range docxOrdered {

		id = generatesIdentifier("Docx", int(stage.DocxMap_Staged_Order[docx]), docx.Name)
		map_Docx_Identifiers[docx] = id

		decl = IdentifiersDecls
		decl = strings.ReplaceAll(decl, "{{Identifier}}", id)
		decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "Docx")
		decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", docx.Name)
		identifiersDecl += decl

		initializerStatements += "\n"
		// Initialisation of values
		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Name")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(docx.Name))
		initializerStatements += setValueField

	}

	map_File_Identifiers := make(map[*File]string)
	_ = map_File_Identifiers

	fileOrdered := []*File{}
	for file := range stage.Files {
		fileOrdered = append(fileOrdered, file)
	}
	sort.Slice(fileOrdered[:], func(i, j int) bool {
		filei := fileOrdered[i]
		filej := fileOrdered[j]
		filei_order, oki := stage.FileMap_Staged_Order[filei]
		filej_order, okj := stage.FileMap_Staged_Order[filej]
		if !oki || !okj {
			log.Fatalln("unknown pointers")
		}
		return filei_order < filej_order
	})
	if len(fileOrdered) > 0 {
		identifiersDecl += "\n"
	}
	for _, file := range fileOrdered {

		id = generatesIdentifier("File", int(stage.FileMap_Staged_Order[file]), file.Name)
		map_File_Identifiers[file] = id

		decl = IdentifiersDecls
		decl = strings.ReplaceAll(decl, "{{Identifier}}", id)
		decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "File")
		decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", file.Name)
		identifiersDecl += decl

		initializerStatements += "\n"
		// Initialisation of values
		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Name")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(file.Name))
		initializerStatements += setValueField

	}

	map_Node_Identifiers := make(map[*Node]string)
	_ = map_Node_Identifiers

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

		id = generatesIdentifier("Node", int(stage.NodeMap_Staged_Order[node]), node.Name)
		map_Node_Identifiers[node] = id

		decl = IdentifiersDecls
		decl = strings.ReplaceAll(decl, "{{Identifier}}", id)
		decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "Node")
		decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", node.Name)
		identifiersDecl += decl

		initializerStatements += "\n"
		// Initialisation of values
		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Name")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(node.Name))
		initializerStatements += setValueField

	}

	map_Paragraph_Identifiers := make(map[*Paragraph]string)
	_ = map_Paragraph_Identifiers

	paragraphOrdered := []*Paragraph{}
	for paragraph := range stage.Paragraphs {
		paragraphOrdered = append(paragraphOrdered, paragraph)
	}
	sort.Slice(paragraphOrdered[:], func(i, j int) bool {
		paragraphi := paragraphOrdered[i]
		paragraphj := paragraphOrdered[j]
		paragraphi_order, oki := stage.ParagraphMap_Staged_Order[paragraphi]
		paragraphj_order, okj := stage.ParagraphMap_Staged_Order[paragraphj]
		if !oki || !okj {
			log.Fatalln("unknown pointers")
		}
		return paragraphi_order < paragraphj_order
	})
	if len(paragraphOrdered) > 0 {
		identifiersDecl += "\n"
	}
	for _, paragraph := range paragraphOrdered {

		id = generatesIdentifier("Paragraph", int(stage.ParagraphMap_Staged_Order[paragraph]), paragraph.Name)
		map_Paragraph_Identifiers[paragraph] = id

		decl = IdentifiersDecls
		decl = strings.ReplaceAll(decl, "{{Identifier}}", id)
		decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "Paragraph")
		decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", paragraph.Name)
		identifiersDecl += decl

		initializerStatements += "\n"
		// Initialisation of values
		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Name")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(paragraph.Name))
		initializerStatements += setValueField

		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Content")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(paragraph.Content))
		initializerStatements += setValueField

		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Text")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(paragraph.Text))
		initializerStatements += setValueField

	}

	map_ParagraphProperties_Identifiers := make(map[*ParagraphProperties]string)
	_ = map_ParagraphProperties_Identifiers

	paragraphpropertiesOrdered := []*ParagraphProperties{}
	for paragraphproperties := range stage.ParagraphPropertiess {
		paragraphpropertiesOrdered = append(paragraphpropertiesOrdered, paragraphproperties)
	}
	sort.Slice(paragraphpropertiesOrdered[:], func(i, j int) bool {
		paragraphpropertiesi := paragraphpropertiesOrdered[i]
		paragraphpropertiesj := paragraphpropertiesOrdered[j]
		paragraphpropertiesi_order, oki := stage.ParagraphPropertiesMap_Staged_Order[paragraphpropertiesi]
		paragraphpropertiesj_order, okj := stage.ParagraphPropertiesMap_Staged_Order[paragraphpropertiesj]
		if !oki || !okj {
			log.Fatalln("unknown pointers")
		}
		return paragraphpropertiesi_order < paragraphpropertiesj_order
	})
	if len(paragraphpropertiesOrdered) > 0 {
		identifiersDecl += "\n"
	}
	for _, paragraphproperties := range paragraphpropertiesOrdered {

		id = generatesIdentifier("ParagraphProperties", int(stage.ParagraphPropertiesMap_Staged_Order[paragraphproperties]), paragraphproperties.Name)
		map_ParagraphProperties_Identifiers[paragraphproperties] = id

		decl = IdentifiersDecls
		decl = strings.ReplaceAll(decl, "{{Identifier}}", id)
		decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "ParagraphProperties")
		decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", paragraphproperties.Name)
		identifiersDecl += decl

		initializerStatements += "\n"
		// Initialisation of values
		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Name")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(paragraphproperties.Name))
		initializerStatements += setValueField

		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Content")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(paragraphproperties.Content))
		initializerStatements += setValueField

	}

	map_ParagraphStyle_Identifiers := make(map[*ParagraphStyle]string)
	_ = map_ParagraphStyle_Identifiers

	paragraphstyleOrdered := []*ParagraphStyle{}
	for paragraphstyle := range stage.ParagraphStyles {
		paragraphstyleOrdered = append(paragraphstyleOrdered, paragraphstyle)
	}
	sort.Slice(paragraphstyleOrdered[:], func(i, j int) bool {
		paragraphstylei := paragraphstyleOrdered[i]
		paragraphstylej := paragraphstyleOrdered[j]
		paragraphstylei_order, oki := stage.ParagraphStyleMap_Staged_Order[paragraphstylei]
		paragraphstylej_order, okj := stage.ParagraphStyleMap_Staged_Order[paragraphstylej]
		if !oki || !okj {
			log.Fatalln("unknown pointers")
		}
		return paragraphstylei_order < paragraphstylej_order
	})
	if len(paragraphstyleOrdered) > 0 {
		identifiersDecl += "\n"
	}
	for _, paragraphstyle := range paragraphstyleOrdered {

		id = generatesIdentifier("ParagraphStyle", int(stage.ParagraphStyleMap_Staged_Order[paragraphstyle]), paragraphstyle.Name)
		map_ParagraphStyle_Identifiers[paragraphstyle] = id

		decl = IdentifiersDecls
		decl = strings.ReplaceAll(decl, "{{Identifier}}", id)
		decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "ParagraphStyle")
		decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", paragraphstyle.Name)
		identifiersDecl += decl

		initializerStatements += "\n"
		// Initialisation of values
		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Name")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(paragraphstyle.Name))
		initializerStatements += setValueField

		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Content")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(paragraphstyle.Content))
		initializerStatements += setValueField

		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "ValAttr")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(paragraphstyle.ValAttr))
		initializerStatements += setValueField

	}

	map_Rune_Identifiers := make(map[*Rune]string)
	_ = map_Rune_Identifiers

	runeOrdered := []*Rune{}
	for rune := range stage.Runes {
		runeOrdered = append(runeOrdered, rune)
	}
	sort.Slice(runeOrdered[:], func(i, j int) bool {
		runei := runeOrdered[i]
		runej := runeOrdered[j]
		runei_order, oki := stage.RuneMap_Staged_Order[runei]
		runej_order, okj := stage.RuneMap_Staged_Order[runej]
		if !oki || !okj {
			log.Fatalln("unknown pointers")
		}
		return runei_order < runej_order
	})
	if len(runeOrdered) > 0 {
		identifiersDecl += "\n"
	}
	for _, rune := range runeOrdered {

		id = generatesIdentifier("Rune", int(stage.RuneMap_Staged_Order[rune]), rune.Name)
		map_Rune_Identifiers[rune] = id

		decl = IdentifiersDecls
		decl = strings.ReplaceAll(decl, "{{Identifier}}", id)
		decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "Rune")
		decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", rune.Name)
		identifiersDecl += decl

		initializerStatements += "\n"
		// Initialisation of values
		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Name")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(rune.Name))
		initializerStatements += setValueField

		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Content")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(rune.Content))
		initializerStatements += setValueField

	}

	map_RuneProperties_Identifiers := make(map[*RuneProperties]string)
	_ = map_RuneProperties_Identifiers

	runepropertiesOrdered := []*RuneProperties{}
	for runeproperties := range stage.RunePropertiess {
		runepropertiesOrdered = append(runepropertiesOrdered, runeproperties)
	}
	sort.Slice(runepropertiesOrdered[:], func(i, j int) bool {
		runepropertiesi := runepropertiesOrdered[i]
		runepropertiesj := runepropertiesOrdered[j]
		runepropertiesi_order, oki := stage.RunePropertiesMap_Staged_Order[runepropertiesi]
		runepropertiesj_order, okj := stage.RunePropertiesMap_Staged_Order[runepropertiesj]
		if !oki || !okj {
			log.Fatalln("unknown pointers")
		}
		return runepropertiesi_order < runepropertiesj_order
	})
	if len(runepropertiesOrdered) > 0 {
		identifiersDecl += "\n"
	}
	for _, runeproperties := range runepropertiesOrdered {

		id = generatesIdentifier("RuneProperties", int(stage.RunePropertiesMap_Staged_Order[runeproperties]), runeproperties.Name)
		map_RuneProperties_Identifiers[runeproperties] = id

		decl = IdentifiersDecls
		decl = strings.ReplaceAll(decl, "{{Identifier}}", id)
		decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "RuneProperties")
		decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", runeproperties.Name)
		identifiersDecl += decl

		initializerStatements += "\n"
		// Initialisation of values
		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Name")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(runeproperties.Name))
		initializerStatements += setValueField

		setValueField = NumberInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "IsBold")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%t", runeproperties.IsBold))
		initializerStatements += setValueField

		setValueField = NumberInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "IsStrike")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%t", runeproperties.IsStrike))
		initializerStatements += setValueField

		setValueField = NumberInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "IsItalic")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%t", runeproperties.IsItalic))
		initializerStatements += setValueField

		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Content")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(runeproperties.Content))
		initializerStatements += setValueField

	}

	map_Table_Identifiers := make(map[*Table]string)
	_ = map_Table_Identifiers

	tableOrdered := []*Table{}
	for table := range stage.Tables {
		tableOrdered = append(tableOrdered, table)
	}
	sort.Slice(tableOrdered[:], func(i, j int) bool {
		tablei := tableOrdered[i]
		tablej := tableOrdered[j]
		tablei_order, oki := stage.TableMap_Staged_Order[tablei]
		tablej_order, okj := stage.TableMap_Staged_Order[tablej]
		if !oki || !okj {
			log.Fatalln("unknown pointers")
		}
		return tablei_order < tablej_order
	})
	if len(tableOrdered) > 0 {
		identifiersDecl += "\n"
	}
	for _, table := range tableOrdered {

		id = generatesIdentifier("Table", int(stage.TableMap_Staged_Order[table]), table.Name)
		map_Table_Identifiers[table] = id

		decl = IdentifiersDecls
		decl = strings.ReplaceAll(decl, "{{Identifier}}", id)
		decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "Table")
		decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", table.Name)
		identifiersDecl += decl

		initializerStatements += "\n"
		// Initialisation of values
		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Name")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(table.Name))
		initializerStatements += setValueField

		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Content")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(table.Content))
		initializerStatements += setValueField

	}

	map_TableColumn_Identifiers := make(map[*TableColumn]string)
	_ = map_TableColumn_Identifiers

	tablecolumnOrdered := []*TableColumn{}
	for tablecolumn := range stage.TableColumns {
		tablecolumnOrdered = append(tablecolumnOrdered, tablecolumn)
	}
	sort.Slice(tablecolumnOrdered[:], func(i, j int) bool {
		tablecolumni := tablecolumnOrdered[i]
		tablecolumnj := tablecolumnOrdered[j]
		tablecolumni_order, oki := stage.TableColumnMap_Staged_Order[tablecolumni]
		tablecolumnj_order, okj := stage.TableColumnMap_Staged_Order[tablecolumnj]
		if !oki || !okj {
			log.Fatalln("unknown pointers")
		}
		return tablecolumni_order < tablecolumnj_order
	})
	if len(tablecolumnOrdered) > 0 {
		identifiersDecl += "\n"
	}
	for _, tablecolumn := range tablecolumnOrdered {

		id = generatesIdentifier("TableColumn", int(stage.TableColumnMap_Staged_Order[tablecolumn]), tablecolumn.Name)
		map_TableColumn_Identifiers[tablecolumn] = id

		decl = IdentifiersDecls
		decl = strings.ReplaceAll(decl, "{{Identifier}}", id)
		decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "TableColumn")
		decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", tablecolumn.Name)
		identifiersDecl += decl

		initializerStatements += "\n"
		// Initialisation of values
		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Name")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(tablecolumn.Name))
		initializerStatements += setValueField

		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Content")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(tablecolumn.Content))
		initializerStatements += setValueField

	}

	map_TableProperties_Identifiers := make(map[*TableProperties]string)
	_ = map_TableProperties_Identifiers

	tablepropertiesOrdered := []*TableProperties{}
	for tableproperties := range stage.TablePropertiess {
		tablepropertiesOrdered = append(tablepropertiesOrdered, tableproperties)
	}
	sort.Slice(tablepropertiesOrdered[:], func(i, j int) bool {
		tablepropertiesi := tablepropertiesOrdered[i]
		tablepropertiesj := tablepropertiesOrdered[j]
		tablepropertiesi_order, oki := stage.TablePropertiesMap_Staged_Order[tablepropertiesi]
		tablepropertiesj_order, okj := stage.TablePropertiesMap_Staged_Order[tablepropertiesj]
		if !oki || !okj {
			log.Fatalln("unknown pointers")
		}
		return tablepropertiesi_order < tablepropertiesj_order
	})
	if len(tablepropertiesOrdered) > 0 {
		identifiersDecl += "\n"
	}
	for _, tableproperties := range tablepropertiesOrdered {

		id = generatesIdentifier("TableProperties", int(stage.TablePropertiesMap_Staged_Order[tableproperties]), tableproperties.Name)
		map_TableProperties_Identifiers[tableproperties] = id

		decl = IdentifiersDecls
		decl = strings.ReplaceAll(decl, "{{Identifier}}", id)
		decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "TableProperties")
		decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", tableproperties.Name)
		identifiersDecl += decl

		initializerStatements += "\n"
		// Initialisation of values
		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Name")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(tableproperties.Name))
		initializerStatements += setValueField

		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Content")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(tableproperties.Content))
		initializerStatements += setValueField

	}

	map_TableRow_Identifiers := make(map[*TableRow]string)
	_ = map_TableRow_Identifiers

	tablerowOrdered := []*TableRow{}
	for tablerow := range stage.TableRows {
		tablerowOrdered = append(tablerowOrdered, tablerow)
	}
	sort.Slice(tablerowOrdered[:], func(i, j int) bool {
		tablerowi := tablerowOrdered[i]
		tablerowj := tablerowOrdered[j]
		tablerowi_order, oki := stage.TableRowMap_Staged_Order[tablerowi]
		tablerowj_order, okj := stage.TableRowMap_Staged_Order[tablerowj]
		if !oki || !okj {
			log.Fatalln("unknown pointers")
		}
		return tablerowi_order < tablerowj_order
	})
	if len(tablerowOrdered) > 0 {
		identifiersDecl += "\n"
	}
	for _, tablerow := range tablerowOrdered {

		id = generatesIdentifier("TableRow", int(stage.TableRowMap_Staged_Order[tablerow]), tablerow.Name)
		map_TableRow_Identifiers[tablerow] = id

		decl = IdentifiersDecls
		decl = strings.ReplaceAll(decl, "{{Identifier}}", id)
		decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "TableRow")
		decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", tablerow.Name)
		identifiersDecl += decl

		initializerStatements += "\n"
		// Initialisation of values
		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Name")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(tablerow.Name))
		initializerStatements += setValueField

		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Content")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(tablerow.Content))
		initializerStatements += setValueField

	}

	map_TableStyle_Identifiers := make(map[*TableStyle]string)
	_ = map_TableStyle_Identifiers

	tablestyleOrdered := []*TableStyle{}
	for tablestyle := range stage.TableStyles {
		tablestyleOrdered = append(tablestyleOrdered, tablestyle)
	}
	sort.Slice(tablestyleOrdered[:], func(i, j int) bool {
		tablestylei := tablestyleOrdered[i]
		tablestylej := tablestyleOrdered[j]
		tablestylei_order, oki := stage.TableStyleMap_Staged_Order[tablestylei]
		tablestylej_order, okj := stage.TableStyleMap_Staged_Order[tablestylej]
		if !oki || !okj {
			log.Fatalln("unknown pointers")
		}
		return tablestylei_order < tablestylej_order
	})
	if len(tablestyleOrdered) > 0 {
		identifiersDecl += "\n"
	}
	for _, tablestyle := range tablestyleOrdered {

		id = generatesIdentifier("TableStyle", int(stage.TableStyleMap_Staged_Order[tablestyle]), tablestyle.Name)
		map_TableStyle_Identifiers[tablestyle] = id

		decl = IdentifiersDecls
		decl = strings.ReplaceAll(decl, "{{Identifier}}", id)
		decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "TableStyle")
		decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", tablestyle.Name)
		identifiersDecl += decl

		initializerStatements += "\n"
		// Initialisation of values
		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Name")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(tablestyle.Name))
		initializerStatements += setValueField

		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Content")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(tablestyle.Content))
		initializerStatements += setValueField

		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Val")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(tablestyle.Val))
		initializerStatements += setValueField

	}

	map_Text_Identifiers := make(map[*Text]string)
	_ = map_Text_Identifiers

	textOrdered := []*Text{}
	for text := range stage.Texts {
		textOrdered = append(textOrdered, text)
	}
	sort.Slice(textOrdered[:], func(i, j int) bool {
		texti := textOrdered[i]
		textj := textOrdered[j]
		texti_order, oki := stage.TextMap_Staged_Order[texti]
		textj_order, okj := stage.TextMap_Staged_Order[textj]
		if !oki || !okj {
			log.Fatalln("unknown pointers")
		}
		return texti_order < textj_order
	})
	if len(textOrdered) > 0 {
		identifiersDecl += "\n"
	}
	for _, text := range textOrdered {

		id = generatesIdentifier("Text", int(stage.TextMap_Staged_Order[text]), text.Name)
		map_Text_Identifiers[text] = id

		decl = IdentifiersDecls
		decl = strings.ReplaceAll(decl, "{{Identifier}}", id)
		decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "Text")
		decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", text.Name)
		identifiersDecl += decl

		initializerStatements += "\n"
		// Initialisation of values
		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Name")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(text.Name))
		initializerStatements += setValueField

		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Content")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(text.Content))
		initializerStatements += setValueField

		setValueField = NumberInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "PreserveWhiteSpace")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%t", text.PreserveWhiteSpace))
		initializerStatements += setValueField

	}

	// insertion initialization of objects to stage
	if len(bodyOrdered) > 0 {
		pointersInitializesStatements += "\n\t// setup of Body instances pointers"
	}
	for _, body := range bodyOrdered {
		var setPointerField string
		_ = setPointerField

		id = generatesIdentifier("Body", int(stage.BodyMap_Staged_Order[body]), body.Name)
		map_Body_Identifiers[body] = id

		// Initialisation of values
		for _, _paragraph := range body.Paragraphs {
			setPointerField = SliceOfPointersFieldInitStatement
			setPointerField = strings.ReplaceAll(setPointerField, "{{Identifier}}", id)
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldName}}", "Paragraphs")
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldNameValue}}", map_Paragraph_Identifiers[_paragraph])
			pointersInitializesStatements += setPointerField
		}

		for _, _table := range body.Tables {
			setPointerField = SliceOfPointersFieldInitStatement
			setPointerField = strings.ReplaceAll(setPointerField, "{{Identifier}}", id)
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldName}}", "Tables")
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldNameValue}}", map_Table_Identifiers[_table])
			pointersInitializesStatements += setPointerField
		}

		if body.LastParagraph != nil {
			setPointerField = PointerFieldInitStatement
			setPointerField = strings.ReplaceAll(setPointerField, "{{Identifier}}", id)
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldName}}", "LastParagraph")
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldNameValue}}", map_Paragraph_Identifiers[body.LastParagraph])
			pointersInitializesStatements += setPointerField
		}

	}

	if len(documentOrdered) > 0 {
		pointersInitializesStatements += "\n\t// setup of Document instances pointers"
	}
	for _, document := range documentOrdered {
		var setPointerField string
		_ = setPointerField

		id = generatesIdentifier("Document", int(stage.DocumentMap_Staged_Order[document]), document.Name)
		map_Document_Identifiers[document] = id

		// Initialisation of values
		if document.File != nil {
			setPointerField = PointerFieldInitStatement
			setPointerField = strings.ReplaceAll(setPointerField, "{{Identifier}}", id)
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldName}}", "File")
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldNameValue}}", map_File_Identifiers[document.File])
			pointersInitializesStatements += setPointerField
		}

		if document.Root != nil {
			setPointerField = PointerFieldInitStatement
			setPointerField = strings.ReplaceAll(setPointerField, "{{Identifier}}", id)
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldName}}", "Root")
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldNameValue}}", map_Node_Identifiers[document.Root])
			pointersInitializesStatements += setPointerField
		}

		if document.Body != nil {
			setPointerField = PointerFieldInitStatement
			setPointerField = strings.ReplaceAll(setPointerField, "{{Identifier}}", id)
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldName}}", "Body")
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldNameValue}}", map_Body_Identifiers[document.Body])
			pointersInitializesStatements += setPointerField
		}

	}

	if len(docxOrdered) > 0 {
		pointersInitializesStatements += "\n\t// setup of Docx instances pointers"
	}
	for _, docx := range docxOrdered {
		var setPointerField string
		_ = setPointerField

		id = generatesIdentifier("Docx", int(stage.DocxMap_Staged_Order[docx]), docx.Name)
		map_Docx_Identifiers[docx] = id

		// Initialisation of values
		for _, _file := range docx.Files {
			setPointerField = SliceOfPointersFieldInitStatement
			setPointerField = strings.ReplaceAll(setPointerField, "{{Identifier}}", id)
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldName}}", "Files")
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldNameValue}}", map_File_Identifiers[_file])
			pointersInitializesStatements += setPointerField
		}

		if docx.Document != nil {
			setPointerField = PointerFieldInitStatement
			setPointerField = strings.ReplaceAll(setPointerField, "{{Identifier}}", id)
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldName}}", "Document")
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldNameValue}}", map_Document_Identifiers[docx.Document])
			pointersInitializesStatements += setPointerField
		}

	}

	if len(fileOrdered) > 0 {
		pointersInitializesStatements += "\n\t// setup of File instances pointers"
	}
	for _, file := range fileOrdered {
		var setPointerField string
		_ = setPointerField

		id = generatesIdentifier("File", int(stage.FileMap_Staged_Order[file]), file.Name)
		map_File_Identifiers[file] = id

		// Initialisation of values
	}

	if len(nodeOrdered) > 0 {
		pointersInitializesStatements += "\n\t// setup of Node instances pointers"
	}
	for _, node := range nodeOrdered {
		var setPointerField string
		_ = setPointerField

		id = generatesIdentifier("Node", int(stage.NodeMap_Staged_Order[node]), node.Name)
		map_Node_Identifiers[node] = id

		// Initialisation of values
		for _, _node := range node.Nodes {
			setPointerField = SliceOfPointersFieldInitStatement
			setPointerField = strings.ReplaceAll(setPointerField, "{{Identifier}}", id)
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldName}}", "Nodes")
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldNameValue}}", map_Node_Identifiers[_node])
			pointersInitializesStatements += setPointerField
		}

	}

	if len(paragraphOrdered) > 0 {
		pointersInitializesStatements += "\n\t// setup of Paragraph instances pointers"
	}
	for _, paragraph := range paragraphOrdered {
		var setPointerField string
		_ = setPointerField

		id = generatesIdentifier("Paragraph", int(stage.ParagraphMap_Staged_Order[paragraph]), paragraph.Name)
		map_Paragraph_Identifiers[paragraph] = id

		// Initialisation of values
		if paragraph.Node != nil {
			setPointerField = PointerFieldInitStatement
			setPointerField = strings.ReplaceAll(setPointerField, "{{Identifier}}", id)
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldName}}", "Node")
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldNameValue}}", map_Node_Identifiers[paragraph.Node])
			pointersInitializesStatements += setPointerField
		}

		if paragraph.ParagraphProperties != nil {
			setPointerField = PointerFieldInitStatement
			setPointerField = strings.ReplaceAll(setPointerField, "{{Identifier}}", id)
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldName}}", "ParagraphProperties")
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldNameValue}}", map_ParagraphProperties_Identifiers[paragraph.ParagraphProperties])
			pointersInitializesStatements += setPointerField
		}

		for _, _rune := range paragraph.Runes {
			setPointerField = SliceOfPointersFieldInitStatement
			setPointerField = strings.ReplaceAll(setPointerField, "{{Identifier}}", id)
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldName}}", "Runes")
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldNameValue}}", map_Rune_Identifiers[_rune])
			pointersInitializesStatements += setPointerField
		}

		if paragraph.Next != nil {
			setPointerField = PointerFieldInitStatement
			setPointerField = strings.ReplaceAll(setPointerField, "{{Identifier}}", id)
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldName}}", "Next")
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldNameValue}}", map_Paragraph_Identifiers[paragraph.Next])
			pointersInitializesStatements += setPointerField
		}

		if paragraph.Previous != nil {
			setPointerField = PointerFieldInitStatement
			setPointerField = strings.ReplaceAll(setPointerField, "{{Identifier}}", id)
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldName}}", "Previous")
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldNameValue}}", map_Paragraph_Identifiers[paragraph.Previous])
			pointersInitializesStatements += setPointerField
		}

		if paragraph.EnclosingBody != nil {
			setPointerField = PointerFieldInitStatement
			setPointerField = strings.ReplaceAll(setPointerField, "{{Identifier}}", id)
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldName}}", "EnclosingBody")
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldNameValue}}", map_Body_Identifiers[paragraph.EnclosingBody])
			pointersInitializesStatements += setPointerField
		}

		if paragraph.EnclosingTableColumn != nil {
			setPointerField = PointerFieldInitStatement
			setPointerField = strings.ReplaceAll(setPointerField, "{{Identifier}}", id)
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldName}}", "EnclosingTableColumn")
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldNameValue}}", map_TableColumn_Identifiers[paragraph.EnclosingTableColumn])
			pointersInitializesStatements += setPointerField
		}

	}

	if len(paragraphpropertiesOrdered) > 0 {
		pointersInitializesStatements += "\n\t// setup of ParagraphProperties instances pointers"
	}
	for _, paragraphproperties := range paragraphpropertiesOrdered {
		var setPointerField string
		_ = setPointerField

		id = generatesIdentifier("ParagraphProperties", int(stage.ParagraphPropertiesMap_Staged_Order[paragraphproperties]), paragraphproperties.Name)
		map_ParagraphProperties_Identifiers[paragraphproperties] = id

		// Initialisation of values
		if paragraphproperties.ParagraphStyle != nil {
			setPointerField = PointerFieldInitStatement
			setPointerField = strings.ReplaceAll(setPointerField, "{{Identifier}}", id)
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldName}}", "ParagraphStyle")
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldNameValue}}", map_ParagraphStyle_Identifiers[paragraphproperties.ParagraphStyle])
			pointersInitializesStatements += setPointerField
		}

		if paragraphproperties.Node != nil {
			setPointerField = PointerFieldInitStatement
			setPointerField = strings.ReplaceAll(setPointerField, "{{Identifier}}", id)
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldName}}", "Node")
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldNameValue}}", map_Node_Identifiers[paragraphproperties.Node])
			pointersInitializesStatements += setPointerField
		}

	}

	if len(paragraphstyleOrdered) > 0 {
		pointersInitializesStatements += "\n\t// setup of ParagraphStyle instances pointers"
	}
	for _, paragraphstyle := range paragraphstyleOrdered {
		var setPointerField string
		_ = setPointerField

		id = generatesIdentifier("ParagraphStyle", int(stage.ParagraphStyleMap_Staged_Order[paragraphstyle]), paragraphstyle.Name)
		map_ParagraphStyle_Identifiers[paragraphstyle] = id

		// Initialisation of values
		if paragraphstyle.Node != nil {
			setPointerField = PointerFieldInitStatement
			setPointerField = strings.ReplaceAll(setPointerField, "{{Identifier}}", id)
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldName}}", "Node")
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldNameValue}}", map_Node_Identifiers[paragraphstyle.Node])
			pointersInitializesStatements += setPointerField
		}

	}

	if len(runeOrdered) > 0 {
		pointersInitializesStatements += "\n\t// setup of Rune instances pointers"
	}
	for _, rune := range runeOrdered {
		var setPointerField string
		_ = setPointerField

		id = generatesIdentifier("Rune", int(stage.RuneMap_Staged_Order[rune]), rune.Name)
		map_Rune_Identifiers[rune] = id

		// Initialisation of values
		if rune.Node != nil {
			setPointerField = PointerFieldInitStatement
			setPointerField = strings.ReplaceAll(setPointerField, "{{Identifier}}", id)
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldName}}", "Node")
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldNameValue}}", map_Node_Identifiers[rune.Node])
			pointersInitializesStatements += setPointerField
		}

		if rune.Text != nil {
			setPointerField = PointerFieldInitStatement
			setPointerField = strings.ReplaceAll(setPointerField, "{{Identifier}}", id)
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldName}}", "Text")
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldNameValue}}", map_Text_Identifiers[rune.Text])
			pointersInitializesStatements += setPointerField
		}

		if rune.RuneProperties != nil {
			setPointerField = PointerFieldInitStatement
			setPointerField = strings.ReplaceAll(setPointerField, "{{Identifier}}", id)
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldName}}", "RuneProperties")
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldNameValue}}", map_RuneProperties_Identifiers[rune.RuneProperties])
			pointersInitializesStatements += setPointerField
		}

		if rune.EnclosingParagraph != nil {
			setPointerField = PointerFieldInitStatement
			setPointerField = strings.ReplaceAll(setPointerField, "{{Identifier}}", id)
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldName}}", "EnclosingParagraph")
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldNameValue}}", map_Paragraph_Identifiers[rune.EnclosingParagraph])
			pointersInitializesStatements += setPointerField
		}

	}

	if len(runepropertiesOrdered) > 0 {
		pointersInitializesStatements += "\n\t// setup of RuneProperties instances pointers"
	}
	for _, runeproperties := range runepropertiesOrdered {
		var setPointerField string
		_ = setPointerField

		id = generatesIdentifier("RuneProperties", int(stage.RunePropertiesMap_Staged_Order[runeproperties]), runeproperties.Name)
		map_RuneProperties_Identifiers[runeproperties] = id

		// Initialisation of values
		if runeproperties.Node != nil {
			setPointerField = PointerFieldInitStatement
			setPointerField = strings.ReplaceAll(setPointerField, "{{Identifier}}", id)
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldName}}", "Node")
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldNameValue}}", map_Node_Identifiers[runeproperties.Node])
			pointersInitializesStatements += setPointerField
		}

	}

	if len(tableOrdered) > 0 {
		pointersInitializesStatements += "\n\t// setup of Table instances pointers"
	}
	for _, table := range tableOrdered {
		var setPointerField string
		_ = setPointerField

		id = generatesIdentifier("Table", int(stage.TableMap_Staged_Order[table]), table.Name)
		map_Table_Identifiers[table] = id

		// Initialisation of values
		if table.Node != nil {
			setPointerField = PointerFieldInitStatement
			setPointerField = strings.ReplaceAll(setPointerField, "{{Identifier}}", id)
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldName}}", "Node")
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldNameValue}}", map_Node_Identifiers[table.Node])
			pointersInitializesStatements += setPointerField
		}

		if table.TableProperties != nil {
			setPointerField = PointerFieldInitStatement
			setPointerField = strings.ReplaceAll(setPointerField, "{{Identifier}}", id)
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldName}}", "TableProperties")
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldNameValue}}", map_TableProperties_Identifiers[table.TableProperties])
			pointersInitializesStatements += setPointerField
		}

		for _, _tablerow := range table.TableRows {
			setPointerField = SliceOfPointersFieldInitStatement
			setPointerField = strings.ReplaceAll(setPointerField, "{{Identifier}}", id)
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldName}}", "TableRows")
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldNameValue}}", map_TableRow_Identifiers[_tablerow])
			pointersInitializesStatements += setPointerField
		}

	}

	if len(tablecolumnOrdered) > 0 {
		pointersInitializesStatements += "\n\t// setup of TableColumn instances pointers"
	}
	for _, tablecolumn := range tablecolumnOrdered {
		var setPointerField string
		_ = setPointerField

		id = generatesIdentifier("TableColumn", int(stage.TableColumnMap_Staged_Order[tablecolumn]), tablecolumn.Name)
		map_TableColumn_Identifiers[tablecolumn] = id

		// Initialisation of values
		if tablecolumn.Node != nil {
			setPointerField = PointerFieldInitStatement
			setPointerField = strings.ReplaceAll(setPointerField, "{{Identifier}}", id)
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldName}}", "Node")
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldNameValue}}", map_Node_Identifiers[tablecolumn.Node])
			pointersInitializesStatements += setPointerField
		}

		for _, _paragraph := range tablecolumn.Paragraphs {
			setPointerField = SliceOfPointersFieldInitStatement
			setPointerField = strings.ReplaceAll(setPointerField, "{{Identifier}}", id)
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldName}}", "Paragraphs")
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldNameValue}}", map_Paragraph_Identifiers[_paragraph])
			pointersInitializesStatements += setPointerField
		}

	}

	if len(tablepropertiesOrdered) > 0 {
		pointersInitializesStatements += "\n\t// setup of TableProperties instances pointers"
	}
	for _, tableproperties := range tablepropertiesOrdered {
		var setPointerField string
		_ = setPointerField

		id = generatesIdentifier("TableProperties", int(stage.TablePropertiesMap_Staged_Order[tableproperties]), tableproperties.Name)
		map_TableProperties_Identifiers[tableproperties] = id

		// Initialisation of values
		if tableproperties.Node != nil {
			setPointerField = PointerFieldInitStatement
			setPointerField = strings.ReplaceAll(setPointerField, "{{Identifier}}", id)
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldName}}", "Node")
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldNameValue}}", map_Node_Identifiers[tableproperties.Node])
			pointersInitializesStatements += setPointerField
		}

		if tableproperties.TableStyle != nil {
			setPointerField = PointerFieldInitStatement
			setPointerField = strings.ReplaceAll(setPointerField, "{{Identifier}}", id)
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldName}}", "TableStyle")
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldNameValue}}", map_TableStyle_Identifiers[tableproperties.TableStyle])
			pointersInitializesStatements += setPointerField
		}

	}

	if len(tablerowOrdered) > 0 {
		pointersInitializesStatements += "\n\t// setup of TableRow instances pointers"
	}
	for _, tablerow := range tablerowOrdered {
		var setPointerField string
		_ = setPointerField

		id = generatesIdentifier("TableRow", int(stage.TableRowMap_Staged_Order[tablerow]), tablerow.Name)
		map_TableRow_Identifiers[tablerow] = id

		// Initialisation of values
		if tablerow.Node != nil {
			setPointerField = PointerFieldInitStatement
			setPointerField = strings.ReplaceAll(setPointerField, "{{Identifier}}", id)
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldName}}", "Node")
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldNameValue}}", map_Node_Identifiers[tablerow.Node])
			pointersInitializesStatements += setPointerField
		}

		for _, _tablecolumn := range tablerow.TableColumns {
			setPointerField = SliceOfPointersFieldInitStatement
			setPointerField = strings.ReplaceAll(setPointerField, "{{Identifier}}", id)
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldName}}", "TableColumns")
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldNameValue}}", map_TableColumn_Identifiers[_tablecolumn])
			pointersInitializesStatements += setPointerField
		}

	}

	if len(tablestyleOrdered) > 0 {
		pointersInitializesStatements += "\n\t// setup of TableStyle instances pointers"
	}
	for _, tablestyle := range tablestyleOrdered {
		var setPointerField string
		_ = setPointerField

		id = generatesIdentifier("TableStyle", int(stage.TableStyleMap_Staged_Order[tablestyle]), tablestyle.Name)
		map_TableStyle_Identifiers[tablestyle] = id

		// Initialisation of values
		if tablestyle.Node != nil {
			setPointerField = PointerFieldInitStatement
			setPointerField = strings.ReplaceAll(setPointerField, "{{Identifier}}", id)
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldName}}", "Node")
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldNameValue}}", map_Node_Identifiers[tablestyle.Node])
			pointersInitializesStatements += setPointerField
		}

	}

	if len(textOrdered) > 0 {
		pointersInitializesStatements += "\n\t// setup of Text instances pointers"
	}
	for _, text := range textOrdered {
		var setPointerField string
		_ = setPointerField

		id = generatesIdentifier("Text", int(stage.TextMap_Staged_Order[text]), text.Name)
		map_Text_Identifiers[text] = id

		// Initialisation of values
		if text.Node != nil {
			setPointerField = PointerFieldInitStatement
			setPointerField = strings.ReplaceAll(setPointerField, "{{Identifier}}", id)
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldName}}", "Node")
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldNameValue}}", map_Node_Identifiers[text.Node])
			pointersInitializesStatements += setPointerField
		}

		if text.EnclosingRune != nil {
			setPointerField = PointerFieldInitStatement
			setPointerField = strings.ReplaceAll(setPointerField, "{{Identifier}}", id)
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldName}}", "EnclosingRune")
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldNameValue}}", map_Rune_Identifiers[text.EnclosingRune])
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
	_ = processedString

	//#1030
	identifier = fmt.Sprintf("__%s__%08d_", gongStructName, idx)

	return
}
