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
	
		identifiersDecl += body.GongMarshallIdentifier(stage)

		initializerStatements += "\n"
		// Initialisation of values
		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", body.GongGetIdentifier(stage))
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Name")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(body.Name))
		initializerStatements += setValueField

	}

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
	
		identifiersDecl += document.GongMarshallIdentifier(stage)

		initializerStatements += "\n"
		// Initialisation of values
		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", document.GongGetIdentifier(stage))
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Name")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(document.Name))
		initializerStatements += setValueField

	}

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
	
		identifiersDecl += docx.GongMarshallIdentifier(stage)

		initializerStatements += "\n"
		// Initialisation of values
		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", docx.GongGetIdentifier(stage))
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Name")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(docx.Name))
		initializerStatements += setValueField

	}

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
	
		identifiersDecl += file.GongMarshallIdentifier(stage)

		initializerStatements += "\n"
		// Initialisation of values
		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", file.GongGetIdentifier(stage))
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Name")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(file.Name))
		initializerStatements += setValueField

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
		// Initialisation of values
		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", node.GongGetIdentifier(stage))
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Name")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(node.Name))
		initializerStatements += setValueField

	}

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
	
		identifiersDecl += paragraph.GongMarshallIdentifier(stage)

		initializerStatements += "\n"
		// Initialisation of values
		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", paragraph.GongGetIdentifier(stage))
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Name")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(paragraph.Name))
		initializerStatements += setValueField

		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", paragraph.GongGetIdentifier(stage))
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Content")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(paragraph.Content))
		initializerStatements += setValueField

		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", paragraph.GongGetIdentifier(stage))
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Text")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(paragraph.Text))
		initializerStatements += setValueField

	}

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
	
		identifiersDecl += paragraphproperties.GongMarshallIdentifier(stage)

		initializerStatements += "\n"
		// Initialisation of values
		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", paragraphproperties.GongGetIdentifier(stage))
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Name")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(paragraphproperties.Name))
		initializerStatements += setValueField

		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", paragraphproperties.GongGetIdentifier(stage))
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Content")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(paragraphproperties.Content))
		initializerStatements += setValueField

	}

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
	
		identifiersDecl += paragraphstyle.GongMarshallIdentifier(stage)

		initializerStatements += "\n"
		// Initialisation of values
		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", paragraphstyle.GongGetIdentifier(stage))
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Name")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(paragraphstyle.Name))
		initializerStatements += setValueField

		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", paragraphstyle.GongGetIdentifier(stage))
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Content")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(paragraphstyle.Content))
		initializerStatements += setValueField

		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", paragraphstyle.GongGetIdentifier(stage))
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "ValAttr")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(paragraphstyle.ValAttr))
		initializerStatements += setValueField

	}

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
	
		identifiersDecl += rune.GongMarshallIdentifier(stage)

		initializerStatements += "\n"
		// Initialisation of values
		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", rune.GongGetIdentifier(stage))
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Name")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(rune.Name))
		initializerStatements += setValueField

		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", rune.GongGetIdentifier(stage))
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Content")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(rune.Content))
		initializerStatements += setValueField

	}

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
	
		identifiersDecl += runeproperties.GongMarshallIdentifier(stage)

		initializerStatements += "\n"
		// Initialisation of values
		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", runeproperties.GongGetIdentifier(stage))
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Name")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(runeproperties.Name))
		initializerStatements += setValueField

		setValueField = NumberInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", runeproperties.GongGetIdentifier(stage))
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "IsBold")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%t", runeproperties.IsBold))
		initializerStatements += setValueField

		setValueField = NumberInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", runeproperties.GongGetIdentifier(stage))
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "IsStrike")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%t", runeproperties.IsStrike))
		initializerStatements += setValueField

		setValueField = NumberInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", runeproperties.GongGetIdentifier(stage))
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "IsItalic")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%t", runeproperties.IsItalic))
		initializerStatements += setValueField

		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", runeproperties.GongGetIdentifier(stage))
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Content")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(runeproperties.Content))
		initializerStatements += setValueField

	}

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
	
		identifiersDecl += table.GongMarshallIdentifier(stage)

		initializerStatements += "\n"
		// Initialisation of values
		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", table.GongGetIdentifier(stage))
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Name")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(table.Name))
		initializerStatements += setValueField

		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", table.GongGetIdentifier(stage))
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Content")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(table.Content))
		initializerStatements += setValueField

	}

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
	
		identifiersDecl += tablecolumn.GongMarshallIdentifier(stage)

		initializerStatements += "\n"
		// Initialisation of values
		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", tablecolumn.GongGetIdentifier(stage))
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Name")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(tablecolumn.Name))
		initializerStatements += setValueField

		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", tablecolumn.GongGetIdentifier(stage))
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Content")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(tablecolumn.Content))
		initializerStatements += setValueField

	}

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
	
		identifiersDecl += tableproperties.GongMarshallIdentifier(stage)

		initializerStatements += "\n"
		// Initialisation of values
		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", tableproperties.GongGetIdentifier(stage))
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Name")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(tableproperties.Name))
		initializerStatements += setValueField

		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", tableproperties.GongGetIdentifier(stage))
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Content")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(tableproperties.Content))
		initializerStatements += setValueField

	}

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
	
		identifiersDecl += tablerow.GongMarshallIdentifier(stage)

		initializerStatements += "\n"
		// Initialisation of values
		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", tablerow.GongGetIdentifier(stage))
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Name")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(tablerow.Name))
		initializerStatements += setValueField

		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", tablerow.GongGetIdentifier(stage))
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Content")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(tablerow.Content))
		initializerStatements += setValueField

	}

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
	
		identifiersDecl += tablestyle.GongMarshallIdentifier(stage)

		initializerStatements += "\n"
		// Initialisation of values
		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", tablestyle.GongGetIdentifier(stage))
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Name")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(tablestyle.Name))
		initializerStatements += setValueField

		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", tablestyle.GongGetIdentifier(stage))
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Content")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(tablestyle.Content))
		initializerStatements += setValueField

		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", tablestyle.GongGetIdentifier(stage))
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Val")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(tablestyle.Val))
		initializerStatements += setValueField

	}

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
	
		identifiersDecl += text.GongMarshallIdentifier(stage)

		initializerStatements += "\n"
		// Initialisation of values
		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", text.GongGetIdentifier(stage))
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Name")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(text.Name))
		initializerStatements += setValueField

		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", text.GongGetIdentifier(stage))
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Content")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(text.Content))
		initializerStatements += setValueField

		setValueField = NumberInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", text.GongGetIdentifier(stage))
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "PreserveWhiteSpace")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%t", text.PreserveWhiteSpace))
		initializerStatements += setValueField

	}

	// insertion initialization of objects to stage
	if len(bodyOrdered) > 0 {
		pointersInitializesStatements += "\n\t// setup of Body instances pointers"
	}
	for _, body := range bodyOrdered {
		_ = body
		var setPointerField string
		_ = setPointerField

		// Initialisation of values
		for _, _paragraph := range body.Paragraphs {
			setPointerField = SliceOfPointersFieldInitStatement
			setPointerField = strings.ReplaceAll(setPointerField, "{{Identifier}}", body.GongGetIdentifier(stage))
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldName}}", "Paragraphs")
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldNameValue}}", _paragraph.GongGetIdentifier(stage))
			pointersInitializesStatements += setPointerField
		}

		for _, _table := range body.Tables {
			setPointerField = SliceOfPointersFieldInitStatement
			setPointerField = strings.ReplaceAll(setPointerField, "{{Identifier}}", body.GongGetIdentifier(stage))
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldName}}", "Tables")
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldNameValue}}", _table.GongGetIdentifier(stage))
			pointersInitializesStatements += setPointerField
		}

		if body.LastParagraph != nil {
			setPointerField = PointerFieldInitStatement
			setPointerField = strings.ReplaceAll(setPointerField, "{{Identifier}}", body.GongGetIdentifier(stage))
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldName}}", "LastParagraph")
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldNameValue}}", body.LastParagraph.GongGetIdentifier(stage))
			pointersInitializesStatements += setPointerField
		}

	}

	if len(documentOrdered) > 0 {
		pointersInitializesStatements += "\n\t// setup of Document instances pointers"
	}
	for _, document := range documentOrdered {
		_ = document
		var setPointerField string
		_ = setPointerField

		// Initialisation of values
		if document.File != nil {
			setPointerField = PointerFieldInitStatement
			setPointerField = strings.ReplaceAll(setPointerField, "{{Identifier}}", document.GongGetIdentifier(stage))
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldName}}", "File")
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldNameValue}}", document.File.GongGetIdentifier(stage))
			pointersInitializesStatements += setPointerField
		}

		if document.Root != nil {
			setPointerField = PointerFieldInitStatement
			setPointerField = strings.ReplaceAll(setPointerField, "{{Identifier}}", document.GongGetIdentifier(stage))
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldName}}", "Root")
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldNameValue}}", document.Root.GongGetIdentifier(stage))
			pointersInitializesStatements += setPointerField
		}

		if document.Body != nil {
			setPointerField = PointerFieldInitStatement
			setPointerField = strings.ReplaceAll(setPointerField, "{{Identifier}}", document.GongGetIdentifier(stage))
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldName}}", "Body")
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldNameValue}}", document.Body.GongGetIdentifier(stage))
			pointersInitializesStatements += setPointerField
		}

	}

	if len(docxOrdered) > 0 {
		pointersInitializesStatements += "\n\t// setup of Docx instances pointers"
	}
	for _, docx := range docxOrdered {
		_ = docx
		var setPointerField string
		_ = setPointerField

		// Initialisation of values
		for _, _file := range docx.Files {
			setPointerField = SliceOfPointersFieldInitStatement
			setPointerField = strings.ReplaceAll(setPointerField, "{{Identifier}}", docx.GongGetIdentifier(stage))
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldName}}", "Files")
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldNameValue}}", _file.GongGetIdentifier(stage))
			pointersInitializesStatements += setPointerField
		}

		if docx.Document != nil {
			setPointerField = PointerFieldInitStatement
			setPointerField = strings.ReplaceAll(setPointerField, "{{Identifier}}", docx.GongGetIdentifier(stage))
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldName}}", "Document")
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldNameValue}}", docx.Document.GongGetIdentifier(stage))
			pointersInitializesStatements += setPointerField
		}

	}

	if len(fileOrdered) > 0 {
		pointersInitializesStatements += "\n\t// setup of File instances pointers"
	}
	for _, file := range fileOrdered {
		_ = file
		var setPointerField string
		_ = setPointerField

		// Initialisation of values
	}

	if len(nodeOrdered) > 0 {
		pointersInitializesStatements += "\n\t// setup of Node instances pointers"
	}
	for _, node := range nodeOrdered {
		_ = node
		var setPointerField string
		_ = setPointerField

		// Initialisation of values
		for _, _node := range node.Nodes {
			setPointerField = SliceOfPointersFieldInitStatement
			setPointerField = strings.ReplaceAll(setPointerField, "{{Identifier}}", node.GongGetIdentifier(stage))
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldName}}", "Nodes")
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldNameValue}}", _node.GongGetIdentifier(stage))
			pointersInitializesStatements += setPointerField
		}

	}

	if len(paragraphOrdered) > 0 {
		pointersInitializesStatements += "\n\t// setup of Paragraph instances pointers"
	}
	for _, paragraph := range paragraphOrdered {
		_ = paragraph
		var setPointerField string
		_ = setPointerField

		// Initialisation of values
		if paragraph.Node != nil {
			setPointerField = PointerFieldInitStatement
			setPointerField = strings.ReplaceAll(setPointerField, "{{Identifier}}", paragraph.GongGetIdentifier(stage))
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldName}}", "Node")
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldNameValue}}", paragraph.Node.GongGetIdentifier(stage))
			pointersInitializesStatements += setPointerField
		}

		if paragraph.ParagraphProperties != nil {
			setPointerField = PointerFieldInitStatement
			setPointerField = strings.ReplaceAll(setPointerField, "{{Identifier}}", paragraph.GongGetIdentifier(stage))
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldName}}", "ParagraphProperties")
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldNameValue}}", paragraph.ParagraphProperties.GongGetIdentifier(stage))
			pointersInitializesStatements += setPointerField
		}

		for _, _rune := range paragraph.Runes {
			setPointerField = SliceOfPointersFieldInitStatement
			setPointerField = strings.ReplaceAll(setPointerField, "{{Identifier}}", paragraph.GongGetIdentifier(stage))
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldName}}", "Runes")
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldNameValue}}", _rune.GongGetIdentifier(stage))
			pointersInitializesStatements += setPointerField
		}

		if paragraph.Next != nil {
			setPointerField = PointerFieldInitStatement
			setPointerField = strings.ReplaceAll(setPointerField, "{{Identifier}}", paragraph.GongGetIdentifier(stage))
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldName}}", "Next")
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldNameValue}}", paragraph.Next.GongGetIdentifier(stage))
			pointersInitializesStatements += setPointerField
		}

		if paragraph.Previous != nil {
			setPointerField = PointerFieldInitStatement
			setPointerField = strings.ReplaceAll(setPointerField, "{{Identifier}}", paragraph.GongGetIdentifier(stage))
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldName}}", "Previous")
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldNameValue}}", paragraph.Previous.GongGetIdentifier(stage))
			pointersInitializesStatements += setPointerField
		}

		if paragraph.EnclosingBody != nil {
			setPointerField = PointerFieldInitStatement
			setPointerField = strings.ReplaceAll(setPointerField, "{{Identifier}}", paragraph.GongGetIdentifier(stage))
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldName}}", "EnclosingBody")
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldNameValue}}", paragraph.EnclosingBody.GongGetIdentifier(stage))
			pointersInitializesStatements += setPointerField
		}

		if paragraph.EnclosingTableColumn != nil {
			setPointerField = PointerFieldInitStatement
			setPointerField = strings.ReplaceAll(setPointerField, "{{Identifier}}", paragraph.GongGetIdentifier(stage))
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldName}}", "EnclosingTableColumn")
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldNameValue}}", paragraph.EnclosingTableColumn.GongGetIdentifier(stage))
			pointersInitializesStatements += setPointerField
		}

	}

	if len(paragraphpropertiesOrdered) > 0 {
		pointersInitializesStatements += "\n\t// setup of ParagraphProperties instances pointers"
	}
	for _, paragraphproperties := range paragraphpropertiesOrdered {
		_ = paragraphproperties
		var setPointerField string
		_ = setPointerField

		// Initialisation of values
		if paragraphproperties.ParagraphStyle != nil {
			setPointerField = PointerFieldInitStatement
			setPointerField = strings.ReplaceAll(setPointerField, "{{Identifier}}", paragraphproperties.GongGetIdentifier(stage))
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldName}}", "ParagraphStyle")
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldNameValue}}", paragraphproperties.ParagraphStyle.GongGetIdentifier(stage))
			pointersInitializesStatements += setPointerField
		}

		if paragraphproperties.Node != nil {
			setPointerField = PointerFieldInitStatement
			setPointerField = strings.ReplaceAll(setPointerField, "{{Identifier}}", paragraphproperties.GongGetIdentifier(stage))
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldName}}", "Node")
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldNameValue}}", paragraphproperties.Node.GongGetIdentifier(stage))
			pointersInitializesStatements += setPointerField
		}

	}

	if len(paragraphstyleOrdered) > 0 {
		pointersInitializesStatements += "\n\t// setup of ParagraphStyle instances pointers"
	}
	for _, paragraphstyle := range paragraphstyleOrdered {
		_ = paragraphstyle
		var setPointerField string
		_ = setPointerField

		// Initialisation of values
		if paragraphstyle.Node != nil {
			setPointerField = PointerFieldInitStatement
			setPointerField = strings.ReplaceAll(setPointerField, "{{Identifier}}", paragraphstyle.GongGetIdentifier(stage))
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldName}}", "Node")
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldNameValue}}", paragraphstyle.Node.GongGetIdentifier(stage))
			pointersInitializesStatements += setPointerField
		}

	}

	if len(runeOrdered) > 0 {
		pointersInitializesStatements += "\n\t// setup of Rune instances pointers"
	}
	for _, rune := range runeOrdered {
		_ = rune
		var setPointerField string
		_ = setPointerField

		// Initialisation of values
		if rune.Node != nil {
			setPointerField = PointerFieldInitStatement
			setPointerField = strings.ReplaceAll(setPointerField, "{{Identifier}}", rune.GongGetIdentifier(stage))
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldName}}", "Node")
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldNameValue}}", rune.Node.GongGetIdentifier(stage))
			pointersInitializesStatements += setPointerField
		}

		if rune.Text != nil {
			setPointerField = PointerFieldInitStatement
			setPointerField = strings.ReplaceAll(setPointerField, "{{Identifier}}", rune.GongGetIdentifier(stage))
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldName}}", "Text")
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldNameValue}}", rune.Text.GongGetIdentifier(stage))
			pointersInitializesStatements += setPointerField
		}

		if rune.RuneProperties != nil {
			setPointerField = PointerFieldInitStatement
			setPointerField = strings.ReplaceAll(setPointerField, "{{Identifier}}", rune.GongGetIdentifier(stage))
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldName}}", "RuneProperties")
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldNameValue}}", rune.RuneProperties.GongGetIdentifier(stage))
			pointersInitializesStatements += setPointerField
		}

		if rune.EnclosingParagraph != nil {
			setPointerField = PointerFieldInitStatement
			setPointerField = strings.ReplaceAll(setPointerField, "{{Identifier}}", rune.GongGetIdentifier(stage))
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldName}}", "EnclosingParagraph")
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldNameValue}}", rune.EnclosingParagraph.GongGetIdentifier(stage))
			pointersInitializesStatements += setPointerField
		}

	}

	if len(runepropertiesOrdered) > 0 {
		pointersInitializesStatements += "\n\t// setup of RuneProperties instances pointers"
	}
	for _, runeproperties := range runepropertiesOrdered {
		_ = runeproperties
		var setPointerField string
		_ = setPointerField

		// Initialisation of values
		if runeproperties.Node != nil {
			setPointerField = PointerFieldInitStatement
			setPointerField = strings.ReplaceAll(setPointerField, "{{Identifier}}", runeproperties.GongGetIdentifier(stage))
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldName}}", "Node")
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldNameValue}}", runeproperties.Node.GongGetIdentifier(stage))
			pointersInitializesStatements += setPointerField
		}

	}

	if len(tableOrdered) > 0 {
		pointersInitializesStatements += "\n\t// setup of Table instances pointers"
	}
	for _, table := range tableOrdered {
		_ = table
		var setPointerField string
		_ = setPointerField

		// Initialisation of values
		if table.Node != nil {
			setPointerField = PointerFieldInitStatement
			setPointerField = strings.ReplaceAll(setPointerField, "{{Identifier}}", table.GongGetIdentifier(stage))
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldName}}", "Node")
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldNameValue}}", table.Node.GongGetIdentifier(stage))
			pointersInitializesStatements += setPointerField
		}

		if table.TableProperties != nil {
			setPointerField = PointerFieldInitStatement
			setPointerField = strings.ReplaceAll(setPointerField, "{{Identifier}}", table.GongGetIdentifier(stage))
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldName}}", "TableProperties")
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldNameValue}}", table.TableProperties.GongGetIdentifier(stage))
			pointersInitializesStatements += setPointerField
		}

		for _, _tablerow := range table.TableRows {
			setPointerField = SliceOfPointersFieldInitStatement
			setPointerField = strings.ReplaceAll(setPointerField, "{{Identifier}}", table.GongGetIdentifier(stage))
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldName}}", "TableRows")
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldNameValue}}", _tablerow.GongGetIdentifier(stage))
			pointersInitializesStatements += setPointerField
		}

	}

	if len(tablecolumnOrdered) > 0 {
		pointersInitializesStatements += "\n\t// setup of TableColumn instances pointers"
	}
	for _, tablecolumn := range tablecolumnOrdered {
		_ = tablecolumn
		var setPointerField string
		_ = setPointerField

		// Initialisation of values
		if tablecolumn.Node != nil {
			setPointerField = PointerFieldInitStatement
			setPointerField = strings.ReplaceAll(setPointerField, "{{Identifier}}", tablecolumn.GongGetIdentifier(stage))
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldName}}", "Node")
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldNameValue}}", tablecolumn.Node.GongGetIdentifier(stage))
			pointersInitializesStatements += setPointerField
		}

		for _, _paragraph := range tablecolumn.Paragraphs {
			setPointerField = SliceOfPointersFieldInitStatement
			setPointerField = strings.ReplaceAll(setPointerField, "{{Identifier}}", tablecolumn.GongGetIdentifier(stage))
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldName}}", "Paragraphs")
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldNameValue}}", _paragraph.GongGetIdentifier(stage))
			pointersInitializesStatements += setPointerField
		}

	}

	if len(tablepropertiesOrdered) > 0 {
		pointersInitializesStatements += "\n\t// setup of TableProperties instances pointers"
	}
	for _, tableproperties := range tablepropertiesOrdered {
		_ = tableproperties
		var setPointerField string
		_ = setPointerField

		// Initialisation of values
		if tableproperties.Node != nil {
			setPointerField = PointerFieldInitStatement
			setPointerField = strings.ReplaceAll(setPointerField, "{{Identifier}}", tableproperties.GongGetIdentifier(stage))
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldName}}", "Node")
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldNameValue}}", tableproperties.Node.GongGetIdentifier(stage))
			pointersInitializesStatements += setPointerField
		}

		if tableproperties.TableStyle != nil {
			setPointerField = PointerFieldInitStatement
			setPointerField = strings.ReplaceAll(setPointerField, "{{Identifier}}", tableproperties.GongGetIdentifier(stage))
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldName}}", "TableStyle")
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldNameValue}}", tableproperties.TableStyle.GongGetIdentifier(stage))
			pointersInitializesStatements += setPointerField
		}

	}

	if len(tablerowOrdered) > 0 {
		pointersInitializesStatements += "\n\t// setup of TableRow instances pointers"
	}
	for _, tablerow := range tablerowOrdered {
		_ = tablerow
		var setPointerField string
		_ = setPointerField

		// Initialisation of values
		if tablerow.Node != nil {
			setPointerField = PointerFieldInitStatement
			setPointerField = strings.ReplaceAll(setPointerField, "{{Identifier}}", tablerow.GongGetIdentifier(stage))
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldName}}", "Node")
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldNameValue}}", tablerow.Node.GongGetIdentifier(stage))
			pointersInitializesStatements += setPointerField
		}

		for _, _tablecolumn := range tablerow.TableColumns {
			setPointerField = SliceOfPointersFieldInitStatement
			setPointerField = strings.ReplaceAll(setPointerField, "{{Identifier}}", tablerow.GongGetIdentifier(stage))
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldName}}", "TableColumns")
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldNameValue}}", _tablecolumn.GongGetIdentifier(stage))
			pointersInitializesStatements += setPointerField
		}

	}

	if len(tablestyleOrdered) > 0 {
		pointersInitializesStatements += "\n\t// setup of TableStyle instances pointers"
	}
	for _, tablestyle := range tablestyleOrdered {
		_ = tablestyle
		var setPointerField string
		_ = setPointerField

		// Initialisation of values
		if tablestyle.Node != nil {
			setPointerField = PointerFieldInitStatement
			setPointerField = strings.ReplaceAll(setPointerField, "{{Identifier}}", tablestyle.GongGetIdentifier(stage))
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldName}}", "Node")
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldNameValue}}", tablestyle.Node.GongGetIdentifier(stage))
			pointersInitializesStatements += setPointerField
		}

	}

	if len(textOrdered) > 0 {
		pointersInitializesStatements += "\n\t// setup of Text instances pointers"
	}
	for _, text := range textOrdered {
		_ = text
		var setPointerField string
		_ = setPointerField

		// Initialisation of values
		if text.Node != nil {
			setPointerField = PointerFieldInitStatement
			setPointerField = strings.ReplaceAll(setPointerField, "{{Identifier}}", text.GongGetIdentifier(stage))
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldName}}", "Node")
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldNameValue}}", text.Node.GongGetIdentifier(stage))
			pointersInitializesStatements += setPointerField
		}

		if text.EnclosingRune != nil {
			setPointerField = PointerFieldInitStatement
			setPointerField = strings.ReplaceAll(setPointerField, "{{Identifier}}", text.GongGetIdentifier(stage))
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldName}}", "EnclosingRune")
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldNameValue}}", text.EnclosingRune.GongGetIdentifier(stage))
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

// insertion initialization of objects to stage
func (body *Body) GongMarshallField(stage *Stage, fieldName string) (setValueField, setPointerField string) {
	initializerStatements := ""
	_ = initializerStatements
	pointersInitializesStatements := ""
	_ = pointersInitializesStatements

	switch fieldName {
	case "Name":
		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", body.GongGetIdentifier(stage))
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Name")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(body.Name))
		initializerStatements += setValueField

	case "Paragraphs":
		for _, _paragraph := range body.Paragraphs {
			setPointerField = SliceOfPointersFieldInitStatement
			setPointerField = strings.ReplaceAll(setPointerField, "{{Identifier}}", body.GongGetIdentifier(stage))
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldName}}", "Paragraphs")
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldNameValue}}", _paragraph.GongGetIdentifier(stage))
			pointersInitializesStatements += setPointerField
		}
	case "Tables":
		for _, _table := range body.Tables {
			setPointerField = SliceOfPointersFieldInitStatement
			setPointerField = strings.ReplaceAll(setPointerField, "{{Identifier}}", body.GongGetIdentifier(stage))
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldName}}", "Tables")
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldNameValue}}", _table.GongGetIdentifier(stage))
			pointersInitializesStatements += setPointerField
		}
	case "LastParagraph":
		if body.LastParagraph != nil {
			setPointerField = PointerFieldInitStatement
			setPointerField = strings.ReplaceAll(setPointerField, "{{Identifier}}", body.GongGetIdentifier(stage))
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldName}}", "LastParagraph")
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldNameValue}}", body.LastParagraph.GongGetIdentifier(stage))
			pointersInitializesStatements += setPointerField
		}

	default:
		log.Panicf("Unknown field %s for Gongstruct Body", fieldName)
	}
	return
}

func (document *Document) GongMarshallField(stage *Stage, fieldName string) (setValueField, setPointerField string) {
	initializerStatements := ""
	_ = initializerStatements
	pointersInitializesStatements := ""
	_ = pointersInitializesStatements

	switch fieldName {
	case "Name":
		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", document.GongGetIdentifier(stage))
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Name")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(document.Name))
		initializerStatements += setValueField

	case "File":
		if document.File != nil {
			setPointerField = PointerFieldInitStatement
			setPointerField = strings.ReplaceAll(setPointerField, "{{Identifier}}", document.GongGetIdentifier(stage))
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldName}}", "File")
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldNameValue}}", document.File.GongGetIdentifier(stage))
			pointersInitializesStatements += setPointerField
		}
	case "Root":
		if document.Root != nil {
			setPointerField = PointerFieldInitStatement
			setPointerField = strings.ReplaceAll(setPointerField, "{{Identifier}}", document.GongGetIdentifier(stage))
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldName}}", "Root")
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldNameValue}}", document.Root.GongGetIdentifier(stage))
			pointersInitializesStatements += setPointerField
		}
	case "Body":
		if document.Body != nil {
			setPointerField = PointerFieldInitStatement
			setPointerField = strings.ReplaceAll(setPointerField, "{{Identifier}}", document.GongGetIdentifier(stage))
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldName}}", "Body")
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldNameValue}}", document.Body.GongGetIdentifier(stage))
			pointersInitializesStatements += setPointerField
		}

	default:
		log.Panicf("Unknown field %s for Gongstruct Document", fieldName)
	}
	return
}

func (docx *Docx) GongMarshallField(stage *Stage, fieldName string) (setValueField, setPointerField string) {
	initializerStatements := ""
	_ = initializerStatements
	pointersInitializesStatements := ""
	_ = pointersInitializesStatements

	switch fieldName {
	case "Name":
		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", docx.GongGetIdentifier(stage))
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Name")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(docx.Name))
		initializerStatements += setValueField

	case "Files":
		for _, _file := range docx.Files {
			setPointerField = SliceOfPointersFieldInitStatement
			setPointerField = strings.ReplaceAll(setPointerField, "{{Identifier}}", docx.GongGetIdentifier(stage))
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldName}}", "Files")
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldNameValue}}", _file.GongGetIdentifier(stage))
			pointersInitializesStatements += setPointerField
		}
	case "Document":
		if docx.Document != nil {
			setPointerField = PointerFieldInitStatement
			setPointerField = strings.ReplaceAll(setPointerField, "{{Identifier}}", docx.GongGetIdentifier(stage))
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldName}}", "Document")
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldNameValue}}", docx.Document.GongGetIdentifier(stage))
			pointersInitializesStatements += setPointerField
		}

	default:
		log.Panicf("Unknown field %s for Gongstruct Docx", fieldName)
	}
	return
}

func (file *File) GongMarshallField(stage *Stage, fieldName string) (setValueField, setPointerField string) {
	initializerStatements := ""
	_ = initializerStatements
	pointersInitializesStatements := ""
	_ = pointersInitializesStatements

	switch fieldName {
	case "Name":
		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", file.GongGetIdentifier(stage))
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Name")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(file.Name))
		initializerStatements += setValueField


	default:
		log.Panicf("Unknown field %s for Gongstruct File", fieldName)
	}
	return
}

func (node *Node) GongMarshallField(stage *Stage, fieldName string) (setValueField, setPointerField string) {
	initializerStatements := ""
	_ = initializerStatements
	pointersInitializesStatements := ""
	_ = pointersInitializesStatements

	switch fieldName {
	case "Name":
		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", node.GongGetIdentifier(stage))
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Name")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(node.Name))
		initializerStatements += setValueField

	case "Nodes":
		for _, _node := range node.Nodes {
			setPointerField = SliceOfPointersFieldInitStatement
			setPointerField = strings.ReplaceAll(setPointerField, "{{Identifier}}", node.GongGetIdentifier(stage))
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldName}}", "Nodes")
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldNameValue}}", _node.GongGetIdentifier(stage))
			pointersInitializesStatements += setPointerField
		}

	default:
		log.Panicf("Unknown field %s for Gongstruct Node", fieldName)
	}
	return
}

func (paragraph *Paragraph) GongMarshallField(stage *Stage, fieldName string) (setValueField, setPointerField string) {
	initializerStatements := ""
	_ = initializerStatements
	pointersInitializesStatements := ""
	_ = pointersInitializesStatements

	switch fieldName {
	case "Name":
		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", paragraph.GongGetIdentifier(stage))
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Name")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(paragraph.Name))
		initializerStatements += setValueField
	case "Content":
		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", paragraph.GongGetIdentifier(stage))
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Content")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(paragraph.Content))
		initializerStatements += setValueField
	case "Text":
		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", paragraph.GongGetIdentifier(stage))
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Text")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(paragraph.Text))
		initializerStatements += setValueField

	case "Node":
		if paragraph.Node != nil {
			setPointerField = PointerFieldInitStatement
			setPointerField = strings.ReplaceAll(setPointerField, "{{Identifier}}", paragraph.GongGetIdentifier(stage))
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldName}}", "Node")
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldNameValue}}", paragraph.Node.GongGetIdentifier(stage))
			pointersInitializesStatements += setPointerField
		}
	case "ParagraphProperties":
		if paragraph.ParagraphProperties != nil {
			setPointerField = PointerFieldInitStatement
			setPointerField = strings.ReplaceAll(setPointerField, "{{Identifier}}", paragraph.GongGetIdentifier(stage))
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldName}}", "ParagraphProperties")
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldNameValue}}", paragraph.ParagraphProperties.GongGetIdentifier(stage))
			pointersInitializesStatements += setPointerField
		}
	case "Runes":
		for _, _rune := range paragraph.Runes {
			setPointerField = SliceOfPointersFieldInitStatement
			setPointerField = strings.ReplaceAll(setPointerField, "{{Identifier}}", paragraph.GongGetIdentifier(stage))
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldName}}", "Runes")
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldNameValue}}", _rune.GongGetIdentifier(stage))
			pointersInitializesStatements += setPointerField
		}
	case "Next":
		if paragraph.Next != nil {
			setPointerField = PointerFieldInitStatement
			setPointerField = strings.ReplaceAll(setPointerField, "{{Identifier}}", paragraph.GongGetIdentifier(stage))
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldName}}", "Next")
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldNameValue}}", paragraph.Next.GongGetIdentifier(stage))
			pointersInitializesStatements += setPointerField
		}
	case "Previous":
		if paragraph.Previous != nil {
			setPointerField = PointerFieldInitStatement
			setPointerField = strings.ReplaceAll(setPointerField, "{{Identifier}}", paragraph.GongGetIdentifier(stage))
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldName}}", "Previous")
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldNameValue}}", paragraph.Previous.GongGetIdentifier(stage))
			pointersInitializesStatements += setPointerField
		}
	case "EnclosingBody":
		if paragraph.EnclosingBody != nil {
			setPointerField = PointerFieldInitStatement
			setPointerField = strings.ReplaceAll(setPointerField, "{{Identifier}}", paragraph.GongGetIdentifier(stage))
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldName}}", "EnclosingBody")
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldNameValue}}", paragraph.EnclosingBody.GongGetIdentifier(stage))
			pointersInitializesStatements += setPointerField
		}
	case "EnclosingTableColumn":
		if paragraph.EnclosingTableColumn != nil {
			setPointerField = PointerFieldInitStatement
			setPointerField = strings.ReplaceAll(setPointerField, "{{Identifier}}", paragraph.GongGetIdentifier(stage))
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldName}}", "EnclosingTableColumn")
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldNameValue}}", paragraph.EnclosingTableColumn.GongGetIdentifier(stage))
			pointersInitializesStatements += setPointerField
		}

	default:
		log.Panicf("Unknown field %s for Gongstruct Paragraph", fieldName)
	}
	return
}

func (paragraphproperties *ParagraphProperties) GongMarshallField(stage *Stage, fieldName string) (setValueField, setPointerField string) {
	initializerStatements := ""
	_ = initializerStatements
	pointersInitializesStatements := ""
	_ = pointersInitializesStatements

	switch fieldName {
	case "Name":
		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", paragraphproperties.GongGetIdentifier(stage))
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Name")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(paragraphproperties.Name))
		initializerStatements += setValueField
	case "Content":
		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", paragraphproperties.GongGetIdentifier(stage))
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Content")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(paragraphproperties.Content))
		initializerStatements += setValueField

	case "ParagraphStyle":
		if paragraphproperties.ParagraphStyle != nil {
			setPointerField = PointerFieldInitStatement
			setPointerField = strings.ReplaceAll(setPointerField, "{{Identifier}}", paragraphproperties.GongGetIdentifier(stage))
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldName}}", "ParagraphStyle")
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldNameValue}}", paragraphproperties.ParagraphStyle.GongGetIdentifier(stage))
			pointersInitializesStatements += setPointerField
		}
	case "Node":
		if paragraphproperties.Node != nil {
			setPointerField = PointerFieldInitStatement
			setPointerField = strings.ReplaceAll(setPointerField, "{{Identifier}}", paragraphproperties.GongGetIdentifier(stage))
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldName}}", "Node")
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldNameValue}}", paragraphproperties.Node.GongGetIdentifier(stage))
			pointersInitializesStatements += setPointerField
		}

	default:
		log.Panicf("Unknown field %s for Gongstruct ParagraphProperties", fieldName)
	}
	return
}

func (paragraphstyle *ParagraphStyle) GongMarshallField(stage *Stage, fieldName string) (setValueField, setPointerField string) {
	initializerStatements := ""
	_ = initializerStatements
	pointersInitializesStatements := ""
	_ = pointersInitializesStatements

	switch fieldName {
	case "Name":
		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", paragraphstyle.GongGetIdentifier(stage))
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Name")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(paragraphstyle.Name))
		initializerStatements += setValueField
	case "Content":
		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", paragraphstyle.GongGetIdentifier(stage))
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Content")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(paragraphstyle.Content))
		initializerStatements += setValueField
	case "ValAttr":
		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", paragraphstyle.GongGetIdentifier(stage))
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "ValAttr")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(paragraphstyle.ValAttr))
		initializerStatements += setValueField

	case "Node":
		if paragraphstyle.Node != nil {
			setPointerField = PointerFieldInitStatement
			setPointerField = strings.ReplaceAll(setPointerField, "{{Identifier}}", paragraphstyle.GongGetIdentifier(stage))
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldName}}", "Node")
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldNameValue}}", paragraphstyle.Node.GongGetIdentifier(stage))
			pointersInitializesStatements += setPointerField
		}

	default:
		log.Panicf("Unknown field %s for Gongstruct ParagraphStyle", fieldName)
	}
	return
}

func (rune *Rune) GongMarshallField(stage *Stage, fieldName string) (setValueField, setPointerField string) {
	initializerStatements := ""
	_ = initializerStatements
	pointersInitializesStatements := ""
	_ = pointersInitializesStatements

	switch fieldName {
	case "Name":
		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", rune.GongGetIdentifier(stage))
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Name")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(rune.Name))
		initializerStatements += setValueField
	case "Content":
		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", rune.GongGetIdentifier(stage))
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Content")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(rune.Content))
		initializerStatements += setValueField

	case "Node":
		if rune.Node != nil {
			setPointerField = PointerFieldInitStatement
			setPointerField = strings.ReplaceAll(setPointerField, "{{Identifier}}", rune.GongGetIdentifier(stage))
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldName}}", "Node")
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldNameValue}}", rune.Node.GongGetIdentifier(stage))
			pointersInitializesStatements += setPointerField
		}
	case "Text":
		if rune.Text != nil {
			setPointerField = PointerFieldInitStatement
			setPointerField = strings.ReplaceAll(setPointerField, "{{Identifier}}", rune.GongGetIdentifier(stage))
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldName}}", "Text")
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldNameValue}}", rune.Text.GongGetIdentifier(stage))
			pointersInitializesStatements += setPointerField
		}
	case "RuneProperties":
		if rune.RuneProperties != nil {
			setPointerField = PointerFieldInitStatement
			setPointerField = strings.ReplaceAll(setPointerField, "{{Identifier}}", rune.GongGetIdentifier(stage))
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldName}}", "RuneProperties")
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldNameValue}}", rune.RuneProperties.GongGetIdentifier(stage))
			pointersInitializesStatements += setPointerField
		}
	case "EnclosingParagraph":
		if rune.EnclosingParagraph != nil {
			setPointerField = PointerFieldInitStatement
			setPointerField = strings.ReplaceAll(setPointerField, "{{Identifier}}", rune.GongGetIdentifier(stage))
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldName}}", "EnclosingParagraph")
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldNameValue}}", rune.EnclosingParagraph.GongGetIdentifier(stage))
			pointersInitializesStatements += setPointerField
		}

	default:
		log.Panicf("Unknown field %s for Gongstruct Rune", fieldName)
	}
	return
}

func (runeproperties *RuneProperties) GongMarshallField(stage *Stage, fieldName string) (setValueField, setPointerField string) {
	initializerStatements := ""
	_ = initializerStatements
	pointersInitializesStatements := ""
	_ = pointersInitializesStatements

	switch fieldName {
	case "Name":
		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", runeproperties.GongGetIdentifier(stage))
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Name")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(runeproperties.Name))
		initializerStatements += setValueField
	case "IsBold":
		setValueField = NumberInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", runeproperties.GongGetIdentifier(stage))
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "IsBold")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%t", runeproperties.IsBold))
		initializerStatements += setValueField
	case "IsStrike":
		setValueField = NumberInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", runeproperties.GongGetIdentifier(stage))
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "IsStrike")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%t", runeproperties.IsStrike))
		initializerStatements += setValueField
	case "IsItalic":
		setValueField = NumberInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", runeproperties.GongGetIdentifier(stage))
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "IsItalic")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%t", runeproperties.IsItalic))
		initializerStatements += setValueField
	case "Content":
		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", runeproperties.GongGetIdentifier(stage))
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Content")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(runeproperties.Content))
		initializerStatements += setValueField

	case "Node":
		if runeproperties.Node != nil {
			setPointerField = PointerFieldInitStatement
			setPointerField = strings.ReplaceAll(setPointerField, "{{Identifier}}", runeproperties.GongGetIdentifier(stage))
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldName}}", "Node")
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldNameValue}}", runeproperties.Node.GongGetIdentifier(stage))
			pointersInitializesStatements += setPointerField
		}

	default:
		log.Panicf("Unknown field %s for Gongstruct RuneProperties", fieldName)
	}
	return
}

func (table *Table) GongMarshallField(stage *Stage, fieldName string) (setValueField, setPointerField string) {
	initializerStatements := ""
	_ = initializerStatements
	pointersInitializesStatements := ""
	_ = pointersInitializesStatements

	switch fieldName {
	case "Name":
		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", table.GongGetIdentifier(stage))
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Name")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(table.Name))
		initializerStatements += setValueField
	case "Content":
		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", table.GongGetIdentifier(stage))
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Content")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(table.Content))
		initializerStatements += setValueField

	case "Node":
		if table.Node != nil {
			setPointerField = PointerFieldInitStatement
			setPointerField = strings.ReplaceAll(setPointerField, "{{Identifier}}", table.GongGetIdentifier(stage))
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldName}}", "Node")
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldNameValue}}", table.Node.GongGetIdentifier(stage))
			pointersInitializesStatements += setPointerField
		}
	case "TableProperties":
		if table.TableProperties != nil {
			setPointerField = PointerFieldInitStatement
			setPointerField = strings.ReplaceAll(setPointerField, "{{Identifier}}", table.GongGetIdentifier(stage))
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldName}}", "TableProperties")
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldNameValue}}", table.TableProperties.GongGetIdentifier(stage))
			pointersInitializesStatements += setPointerField
		}
	case "TableRows":
		for _, _tablerow := range table.TableRows {
			setPointerField = SliceOfPointersFieldInitStatement
			setPointerField = strings.ReplaceAll(setPointerField, "{{Identifier}}", table.GongGetIdentifier(stage))
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldName}}", "TableRows")
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldNameValue}}", _tablerow.GongGetIdentifier(stage))
			pointersInitializesStatements += setPointerField
		}

	default:
		log.Panicf("Unknown field %s for Gongstruct Table", fieldName)
	}
	return
}

func (tablecolumn *TableColumn) GongMarshallField(stage *Stage, fieldName string) (setValueField, setPointerField string) {
	initializerStatements := ""
	_ = initializerStatements
	pointersInitializesStatements := ""
	_ = pointersInitializesStatements

	switch fieldName {
	case "Name":
		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", tablecolumn.GongGetIdentifier(stage))
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Name")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(tablecolumn.Name))
		initializerStatements += setValueField
	case "Content":
		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", tablecolumn.GongGetIdentifier(stage))
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Content")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(tablecolumn.Content))
		initializerStatements += setValueField

	case "Node":
		if tablecolumn.Node != nil {
			setPointerField = PointerFieldInitStatement
			setPointerField = strings.ReplaceAll(setPointerField, "{{Identifier}}", tablecolumn.GongGetIdentifier(stage))
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldName}}", "Node")
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldNameValue}}", tablecolumn.Node.GongGetIdentifier(stage))
			pointersInitializesStatements += setPointerField
		}
	case "Paragraphs":
		for _, _paragraph := range tablecolumn.Paragraphs {
			setPointerField = SliceOfPointersFieldInitStatement
			setPointerField = strings.ReplaceAll(setPointerField, "{{Identifier}}", tablecolumn.GongGetIdentifier(stage))
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldName}}", "Paragraphs")
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldNameValue}}", _paragraph.GongGetIdentifier(stage))
			pointersInitializesStatements += setPointerField
		}

	default:
		log.Panicf("Unknown field %s for Gongstruct TableColumn", fieldName)
	}
	return
}

func (tableproperties *TableProperties) GongMarshallField(stage *Stage, fieldName string) (setValueField, setPointerField string) {
	initializerStatements := ""
	_ = initializerStatements
	pointersInitializesStatements := ""
	_ = pointersInitializesStatements

	switch fieldName {
	case "Name":
		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", tableproperties.GongGetIdentifier(stage))
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Name")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(tableproperties.Name))
		initializerStatements += setValueField
	case "Content":
		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", tableproperties.GongGetIdentifier(stage))
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Content")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(tableproperties.Content))
		initializerStatements += setValueField

	case "Node":
		if tableproperties.Node != nil {
			setPointerField = PointerFieldInitStatement
			setPointerField = strings.ReplaceAll(setPointerField, "{{Identifier}}", tableproperties.GongGetIdentifier(stage))
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldName}}", "Node")
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldNameValue}}", tableproperties.Node.GongGetIdentifier(stage))
			pointersInitializesStatements += setPointerField
		}
	case "TableStyle":
		if tableproperties.TableStyle != nil {
			setPointerField = PointerFieldInitStatement
			setPointerField = strings.ReplaceAll(setPointerField, "{{Identifier}}", tableproperties.GongGetIdentifier(stage))
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldName}}", "TableStyle")
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldNameValue}}", tableproperties.TableStyle.GongGetIdentifier(stage))
			pointersInitializesStatements += setPointerField
		}

	default:
		log.Panicf("Unknown field %s for Gongstruct TableProperties", fieldName)
	}
	return
}

func (tablerow *TableRow) GongMarshallField(stage *Stage, fieldName string) (setValueField, setPointerField string) {
	initializerStatements := ""
	_ = initializerStatements
	pointersInitializesStatements := ""
	_ = pointersInitializesStatements

	switch fieldName {
	case "Name":
		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", tablerow.GongGetIdentifier(stage))
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Name")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(tablerow.Name))
		initializerStatements += setValueField
	case "Content":
		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", tablerow.GongGetIdentifier(stage))
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Content")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(tablerow.Content))
		initializerStatements += setValueField

	case "Node":
		if tablerow.Node != nil {
			setPointerField = PointerFieldInitStatement
			setPointerField = strings.ReplaceAll(setPointerField, "{{Identifier}}", tablerow.GongGetIdentifier(stage))
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldName}}", "Node")
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldNameValue}}", tablerow.Node.GongGetIdentifier(stage))
			pointersInitializesStatements += setPointerField
		}
	case "TableColumns":
		for _, _tablecolumn := range tablerow.TableColumns {
			setPointerField = SliceOfPointersFieldInitStatement
			setPointerField = strings.ReplaceAll(setPointerField, "{{Identifier}}", tablerow.GongGetIdentifier(stage))
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldName}}", "TableColumns")
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldNameValue}}", _tablecolumn.GongGetIdentifier(stage))
			pointersInitializesStatements += setPointerField
		}

	default:
		log.Panicf("Unknown field %s for Gongstruct TableRow", fieldName)
	}
	return
}

func (tablestyle *TableStyle) GongMarshallField(stage *Stage, fieldName string) (setValueField, setPointerField string) {
	initializerStatements := ""
	_ = initializerStatements
	pointersInitializesStatements := ""
	_ = pointersInitializesStatements

	switch fieldName {
	case "Name":
		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", tablestyle.GongGetIdentifier(stage))
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Name")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(tablestyle.Name))
		initializerStatements += setValueField
	case "Content":
		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", tablestyle.GongGetIdentifier(stage))
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Content")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(tablestyle.Content))
		initializerStatements += setValueField
	case "Val":
		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", tablestyle.GongGetIdentifier(stage))
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Val")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(tablestyle.Val))
		initializerStatements += setValueField

	case "Node":
		if tablestyle.Node != nil {
			setPointerField = PointerFieldInitStatement
			setPointerField = strings.ReplaceAll(setPointerField, "{{Identifier}}", tablestyle.GongGetIdentifier(stage))
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldName}}", "Node")
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldNameValue}}", tablestyle.Node.GongGetIdentifier(stage))
			pointersInitializesStatements += setPointerField
		}

	default:
		log.Panicf("Unknown field %s for Gongstruct TableStyle", fieldName)
	}
	return
}

func (text *Text) GongMarshallField(stage *Stage, fieldName string) (setValueField, setPointerField string) {
	initializerStatements := ""
	_ = initializerStatements
	pointersInitializesStatements := ""
	_ = pointersInitializesStatements

	switch fieldName {
	case "Name":
		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", text.GongGetIdentifier(stage))
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Name")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(text.Name))
		initializerStatements += setValueField
	case "Content":
		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", text.GongGetIdentifier(stage))
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Content")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(text.Content))
		initializerStatements += setValueField
	case "PreserveWhiteSpace":
		setValueField = NumberInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", text.GongGetIdentifier(stage))
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "PreserveWhiteSpace")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%t", text.PreserveWhiteSpace))
		initializerStatements += setValueField

	case "Node":
		if text.Node != nil {
			setPointerField = PointerFieldInitStatement
			setPointerField = strings.ReplaceAll(setPointerField, "{{Identifier}}", text.GongGetIdentifier(stage))
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldName}}", "Node")
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldNameValue}}", text.Node.GongGetIdentifier(stage))
			pointersInitializesStatements += setPointerField
		}
	case "EnclosingRune":
		if text.EnclosingRune != nil {
			setPointerField = PointerFieldInitStatement
			setPointerField = strings.ReplaceAll(setPointerField, "{{Identifier}}", text.GongGetIdentifier(stage))
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldName}}", "EnclosingRune")
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldNameValue}}", text.EnclosingRune.GongGetIdentifier(stage))
			pointersInitializesStatements += setPointerField
		}

	default:
		log.Panicf("Unknown field %s for Gongstruct Text", fieldName)
	}
	return
}
