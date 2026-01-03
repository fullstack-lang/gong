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
		// Insertion point for basic fields value assignment
		initializerStatements += body.GongMarshallField(stage, "Name")
		pointersInitializesStatements += body.GongMarshallField(stage, "Paragraphs")
		pointersInitializesStatements += body.GongMarshallField(stage, "Tables")
		pointersInitializesStatements += body.GongMarshallField(stage, "LastParagraph")
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
		// Insertion point for basic fields value assignment
		initializerStatements += document.GongMarshallField(stage, "Name")
		pointersInitializesStatements += document.GongMarshallField(stage, "File")
		pointersInitializesStatements += document.GongMarshallField(stage, "Root")
		pointersInitializesStatements += document.GongMarshallField(stage, "Body")
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
		// Insertion point for basic fields value assignment
		initializerStatements += docx.GongMarshallField(stage, "Name")
		pointersInitializesStatements += docx.GongMarshallField(stage, "Files")
		pointersInitializesStatements += docx.GongMarshallField(stage, "Document")
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
		// Insertion point for basic fields value assignment
		initializerStatements += file.GongMarshallField(stage, "Name")
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
		pointersInitializesStatements += node.GongMarshallField(stage, "Nodes")
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
		// Insertion point for basic fields value assignment
		initializerStatements += paragraph.GongMarshallField(stage, "Name")
		initializerStatements += paragraph.GongMarshallField(stage, "Content")
		pointersInitializesStatements += paragraph.GongMarshallField(stage, "Node")
		pointersInitializesStatements += paragraph.GongMarshallField(stage, "ParagraphProperties")
		pointersInitializesStatements += paragraph.GongMarshallField(stage, "Runes")
		initializerStatements += paragraph.GongMarshallField(stage, "Text")
		pointersInitializesStatements += paragraph.GongMarshallField(stage, "Next")
		pointersInitializesStatements += paragraph.GongMarshallField(stage, "Previous")
		pointersInitializesStatements += paragraph.GongMarshallField(stage, "EnclosingBody")
		pointersInitializesStatements += paragraph.GongMarshallField(stage, "EnclosingTableColumn")
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
		// Insertion point for basic fields value assignment
		initializerStatements += paragraphproperties.GongMarshallField(stage, "Name")
		initializerStatements += paragraphproperties.GongMarshallField(stage, "Content")
		pointersInitializesStatements += paragraphproperties.GongMarshallField(stage, "ParagraphStyle")
		pointersInitializesStatements += paragraphproperties.GongMarshallField(stage, "Node")
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
		// Insertion point for basic fields value assignment
		initializerStatements += paragraphstyle.GongMarshallField(stage, "Name")
		pointersInitializesStatements += paragraphstyle.GongMarshallField(stage, "Node")
		initializerStatements += paragraphstyle.GongMarshallField(stage, "Content")
		initializerStatements += paragraphstyle.GongMarshallField(stage, "ValAttr")
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
		// Insertion point for basic fields value assignment
		initializerStatements += rune.GongMarshallField(stage, "Name")
		initializerStatements += rune.GongMarshallField(stage, "Content")
		pointersInitializesStatements += rune.GongMarshallField(stage, "Node")
		pointersInitializesStatements += rune.GongMarshallField(stage, "Text")
		pointersInitializesStatements += rune.GongMarshallField(stage, "RuneProperties")
		pointersInitializesStatements += rune.GongMarshallField(stage, "EnclosingParagraph")
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
		// Insertion point for basic fields value assignment
		initializerStatements += runeproperties.GongMarshallField(stage, "Name")
		pointersInitializesStatements += runeproperties.GongMarshallField(stage, "Node")
		initializerStatements += runeproperties.GongMarshallField(stage, "IsBold")
		initializerStatements += runeproperties.GongMarshallField(stage, "IsStrike")
		initializerStatements += runeproperties.GongMarshallField(stage, "IsItalic")
		initializerStatements += runeproperties.GongMarshallField(stage, "Content")
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
		// Insertion point for basic fields value assignment
		initializerStatements += table.GongMarshallField(stage, "Name")
		pointersInitializesStatements += table.GongMarshallField(stage, "Node")
		initializerStatements += table.GongMarshallField(stage, "Content")
		pointersInitializesStatements += table.GongMarshallField(stage, "TableProperties")
		pointersInitializesStatements += table.GongMarshallField(stage, "TableRows")
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
		// Insertion point for basic fields value assignment
		initializerStatements += tablecolumn.GongMarshallField(stage, "Name")
		initializerStatements += tablecolumn.GongMarshallField(stage, "Content")
		pointersInitializesStatements += tablecolumn.GongMarshallField(stage, "Node")
		pointersInitializesStatements += tablecolumn.GongMarshallField(stage, "Paragraphs")
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
		// Insertion point for basic fields value assignment
		initializerStatements += tableproperties.GongMarshallField(stage, "Name")
		pointersInitializesStatements += tableproperties.GongMarshallField(stage, "Node")
		initializerStatements += tableproperties.GongMarshallField(stage, "Content")
		pointersInitializesStatements += tableproperties.GongMarshallField(stage, "TableStyle")
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
		// Insertion point for basic fields value assignment
		initializerStatements += tablerow.GongMarshallField(stage, "Name")
		initializerStatements += tablerow.GongMarshallField(stage, "Content")
		pointersInitializesStatements += tablerow.GongMarshallField(stage, "Node")
		pointersInitializesStatements += tablerow.GongMarshallField(stage, "TableColumns")
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
		// Insertion point for basic fields value assignment
		initializerStatements += tablestyle.GongMarshallField(stage, "Name")
		pointersInitializesStatements += tablestyle.GongMarshallField(stage, "Node")
		initializerStatements += tablestyle.GongMarshallField(stage, "Content")
		initializerStatements += tablestyle.GongMarshallField(stage, "Val")
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
		// Insertion point for basic fields value assignment
		initializerStatements += text.GongMarshallField(stage, "Name")
		initializerStatements += text.GongMarshallField(stage, "Content")
		pointersInitializesStatements += text.GongMarshallField(stage, "Node")
		initializerStatements += text.GongMarshallField(stage, "PreserveWhiteSpace")
		pointersInitializesStatements += text.GongMarshallField(stage, "EnclosingRune")
	}

	// insertion initialization of objects to stage
	for _, body := range bodyOrdered {
		_ = body
		var setPointerField string
		_ = setPointerField

		// Insertion point for pointers initialization
	}

	for _, document := range documentOrdered {
		_ = document
		var setPointerField string
		_ = setPointerField

		// Insertion point for pointers initialization
	}

	for _, docx := range docxOrdered {
		_ = docx
		var setPointerField string
		_ = setPointerField

		// Insertion point for pointers initialization
	}

	for _, file := range fileOrdered {
		_ = file
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

	for _, paragraph := range paragraphOrdered {
		_ = paragraph
		var setPointerField string
		_ = setPointerField

		// Insertion point for pointers initialization
	}

	for _, paragraphproperties := range paragraphpropertiesOrdered {
		_ = paragraphproperties
		var setPointerField string
		_ = setPointerField

		// Insertion point for pointers initialization
	}

	for _, paragraphstyle := range paragraphstyleOrdered {
		_ = paragraphstyle
		var setPointerField string
		_ = setPointerField

		// Insertion point for pointers initialization
	}

	for _, rune := range runeOrdered {
		_ = rune
		var setPointerField string
		_ = setPointerField

		// Insertion point for pointers initialization
	}

	for _, runeproperties := range runepropertiesOrdered {
		_ = runeproperties
		var setPointerField string
		_ = setPointerField

		// Insertion point for pointers initialization
	}

	for _, table := range tableOrdered {
		_ = table
		var setPointerField string
		_ = setPointerField

		// Insertion point for pointers initialization
	}

	for _, tablecolumn := range tablecolumnOrdered {
		_ = tablecolumn
		var setPointerField string
		_ = setPointerField

		// Insertion point for pointers initialization
	}

	for _, tableproperties := range tablepropertiesOrdered {
		_ = tableproperties
		var setPointerField string
		_ = setPointerField

		// Insertion point for pointers initialization
	}

	for _, tablerow := range tablerowOrdered {
		_ = tablerow
		var setPointerField string
		_ = setPointerField

		// Insertion point for pointers initialization
	}

	for _, tablestyle := range tablestyleOrdered {
		_ = tablestyle
		var setPointerField string
		_ = setPointerField

		// Insertion point for pointers initialization
	}

	for _, text := range textOrdered {
		_ = text
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
func (body *Body) GongMarshallField(stage *Stage, fieldName string) (res string) {

	switch fieldName {
	case "Name":
		res = StringInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", body.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "Name")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", string(body.Name))

	case "Paragraphs":
		for _, _paragraph := range body.Paragraphs {
			tmp := SliceOfPointersFieldInitStatement
			tmp = strings.ReplaceAll(tmp, "{{Identifier}}", body.GongGetIdentifier(stage))
			tmp = strings.ReplaceAll(tmp, "{{GeneratedFieldName}}", "Paragraphs")
			tmp = strings.ReplaceAll(tmp, "{{GeneratedFieldNameValue}}", _paragraph.GongGetIdentifier(stage))
			res += tmp
		}
	case "Tables":
		for _, _table := range body.Tables {
			tmp := SliceOfPointersFieldInitStatement
			tmp = strings.ReplaceAll(tmp, "{{Identifier}}", body.GongGetIdentifier(stage))
			tmp = strings.ReplaceAll(tmp, "{{GeneratedFieldName}}", "Tables")
			tmp = strings.ReplaceAll(tmp, "{{GeneratedFieldNameValue}}", _table.GongGetIdentifier(stage))
			res += tmp
		}
	case "LastParagraph":
		if body.LastParagraph != nil {
			res = PointerFieldInitStatement
			res = strings.ReplaceAll(res, "{{Identifier}}", body.GongGetIdentifier(stage))
			res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "LastParagraph")
			res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", body.LastParagraph.GongGetIdentifier(stage))
		}
	default:
		log.Panicf("Unknown field %s for Gongstruct Body", fieldName)
	}
	return
}

func (document *Document) GongMarshallField(stage *Stage, fieldName string) (res string) {

	switch fieldName {
	case "Name":
		res = StringInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", document.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "Name")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", string(document.Name))

	case "File":
		if document.File != nil {
			res = PointerFieldInitStatement
			res = strings.ReplaceAll(res, "{{Identifier}}", document.GongGetIdentifier(stage))
			res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "File")
			res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", document.File.GongGetIdentifier(stage))
		}
	case "Root":
		if document.Root != nil {
			res = PointerFieldInitStatement
			res = strings.ReplaceAll(res, "{{Identifier}}", document.GongGetIdentifier(stage))
			res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "Root")
			res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", document.Root.GongGetIdentifier(stage))
		}
	case "Body":
		if document.Body != nil {
			res = PointerFieldInitStatement
			res = strings.ReplaceAll(res, "{{Identifier}}", document.GongGetIdentifier(stage))
			res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "Body")
			res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", document.Body.GongGetIdentifier(stage))
		}
	default:
		log.Panicf("Unknown field %s for Gongstruct Document", fieldName)
	}
	return
}

func (docx *Docx) GongMarshallField(stage *Stage, fieldName string) (res string) {

	switch fieldName {
	case "Name":
		res = StringInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", docx.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "Name")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", string(docx.Name))

	case "Files":
		for _, _file := range docx.Files {
			tmp := SliceOfPointersFieldInitStatement
			tmp = strings.ReplaceAll(tmp, "{{Identifier}}", docx.GongGetIdentifier(stage))
			tmp = strings.ReplaceAll(tmp, "{{GeneratedFieldName}}", "Files")
			tmp = strings.ReplaceAll(tmp, "{{GeneratedFieldNameValue}}", _file.GongGetIdentifier(stage))
			res += tmp
		}
	case "Document":
		if docx.Document != nil {
			res = PointerFieldInitStatement
			res = strings.ReplaceAll(res, "{{Identifier}}", docx.GongGetIdentifier(stage))
			res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "Document")
			res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", docx.Document.GongGetIdentifier(stage))
		}
	default:
		log.Panicf("Unknown field %s for Gongstruct Docx", fieldName)
	}
	return
}

func (file *File) GongMarshallField(stage *Stage, fieldName string) (res string) {

	switch fieldName {
	case "Name":
		res = StringInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", file.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "Name")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", string(file.Name))

	default:
		log.Panicf("Unknown field %s for Gongstruct File", fieldName)
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

	case "Nodes":
		for _, _node := range node.Nodes {
			tmp := SliceOfPointersFieldInitStatement
			tmp = strings.ReplaceAll(tmp, "{{Identifier}}", node.GongGetIdentifier(stage))
			tmp = strings.ReplaceAll(tmp, "{{GeneratedFieldName}}", "Nodes")
			tmp = strings.ReplaceAll(tmp, "{{GeneratedFieldNameValue}}", _node.GongGetIdentifier(stage))
			res += tmp
		}
	default:
		log.Panicf("Unknown field %s for Gongstruct Node", fieldName)
	}
	return
}

func (paragraph *Paragraph) GongMarshallField(stage *Stage, fieldName string) (res string) {

	switch fieldName {
	case "Name":
		res = StringInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", paragraph.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "Name")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", string(paragraph.Name))
	case "Content":
		res = StringInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", paragraph.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "Content")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", string(paragraph.Content))
	case "Text":
		res = StringInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", paragraph.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "Text")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", string(paragraph.Text))

	case "Node":
		if paragraph.Node != nil {
			res = PointerFieldInitStatement
			res = strings.ReplaceAll(res, "{{Identifier}}", paragraph.GongGetIdentifier(stage))
			res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "Node")
			res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", paragraph.Node.GongGetIdentifier(stage))
		}
	case "ParagraphProperties":
		if paragraph.ParagraphProperties != nil {
			res = PointerFieldInitStatement
			res = strings.ReplaceAll(res, "{{Identifier}}", paragraph.GongGetIdentifier(stage))
			res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "ParagraphProperties")
			res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", paragraph.ParagraphProperties.GongGetIdentifier(stage))
		}
	case "Runes":
		for _, _rune := range paragraph.Runes {
			tmp := SliceOfPointersFieldInitStatement
			tmp = strings.ReplaceAll(tmp, "{{Identifier}}", paragraph.GongGetIdentifier(stage))
			tmp = strings.ReplaceAll(tmp, "{{GeneratedFieldName}}", "Runes")
			tmp = strings.ReplaceAll(tmp, "{{GeneratedFieldNameValue}}", _rune.GongGetIdentifier(stage))
			res += tmp
		}
	case "Next":
		if paragraph.Next != nil {
			res = PointerFieldInitStatement
			res = strings.ReplaceAll(res, "{{Identifier}}", paragraph.GongGetIdentifier(stage))
			res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "Next")
			res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", paragraph.Next.GongGetIdentifier(stage))
		}
	case "Previous":
		if paragraph.Previous != nil {
			res = PointerFieldInitStatement
			res = strings.ReplaceAll(res, "{{Identifier}}", paragraph.GongGetIdentifier(stage))
			res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "Previous")
			res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", paragraph.Previous.GongGetIdentifier(stage))
		}
	case "EnclosingBody":
		if paragraph.EnclosingBody != nil {
			res = PointerFieldInitStatement
			res = strings.ReplaceAll(res, "{{Identifier}}", paragraph.GongGetIdentifier(stage))
			res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "EnclosingBody")
			res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", paragraph.EnclosingBody.GongGetIdentifier(stage))
		}
	case "EnclosingTableColumn":
		if paragraph.EnclosingTableColumn != nil {
			res = PointerFieldInitStatement
			res = strings.ReplaceAll(res, "{{Identifier}}", paragraph.GongGetIdentifier(stage))
			res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "EnclosingTableColumn")
			res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", paragraph.EnclosingTableColumn.GongGetIdentifier(stage))
		}
	default:
		log.Panicf("Unknown field %s for Gongstruct Paragraph", fieldName)
	}
	return
}

func (paragraphproperties *ParagraphProperties) GongMarshallField(stage *Stage, fieldName string) (res string) {

	switch fieldName {
	case "Name":
		res = StringInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", paragraphproperties.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "Name")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", string(paragraphproperties.Name))
	case "Content":
		res = StringInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", paragraphproperties.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "Content")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", string(paragraphproperties.Content))

	case "ParagraphStyle":
		if paragraphproperties.ParagraphStyle != nil {
			res = PointerFieldInitStatement
			res = strings.ReplaceAll(res, "{{Identifier}}", paragraphproperties.GongGetIdentifier(stage))
			res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "ParagraphStyle")
			res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", paragraphproperties.ParagraphStyle.GongGetIdentifier(stage))
		}
	case "Node":
		if paragraphproperties.Node != nil {
			res = PointerFieldInitStatement
			res = strings.ReplaceAll(res, "{{Identifier}}", paragraphproperties.GongGetIdentifier(stage))
			res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "Node")
			res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", paragraphproperties.Node.GongGetIdentifier(stage))
		}
	default:
		log.Panicf("Unknown field %s for Gongstruct ParagraphProperties", fieldName)
	}
	return
}

func (paragraphstyle *ParagraphStyle) GongMarshallField(stage *Stage, fieldName string) (res string) {

	switch fieldName {
	case "Name":
		res = StringInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", paragraphstyle.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "Name")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", string(paragraphstyle.Name))
	case "Content":
		res = StringInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", paragraphstyle.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "Content")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", string(paragraphstyle.Content))
	case "ValAttr":
		res = StringInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", paragraphstyle.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "ValAttr")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", string(paragraphstyle.ValAttr))

	case "Node":
		if paragraphstyle.Node != nil {
			res = PointerFieldInitStatement
			res = strings.ReplaceAll(res, "{{Identifier}}", paragraphstyle.GongGetIdentifier(stage))
			res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "Node")
			res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", paragraphstyle.Node.GongGetIdentifier(stage))
		}
	default:
		log.Panicf("Unknown field %s for Gongstruct ParagraphStyle", fieldName)
	}
	return
}

func (rune *Rune) GongMarshallField(stage *Stage, fieldName string) (res string) {

	switch fieldName {
	case "Name":
		res = StringInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", rune.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "Name")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", string(rune.Name))
	case "Content":
		res = StringInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", rune.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "Content")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", string(rune.Content))

	case "Node":
		if rune.Node != nil {
			res = PointerFieldInitStatement
			res = strings.ReplaceAll(res, "{{Identifier}}", rune.GongGetIdentifier(stage))
			res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "Node")
			res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", rune.Node.GongGetIdentifier(stage))
		}
	case "Text":
		if rune.Text != nil {
			res = PointerFieldInitStatement
			res = strings.ReplaceAll(res, "{{Identifier}}", rune.GongGetIdentifier(stage))
			res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "Text")
			res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", rune.Text.GongGetIdentifier(stage))
		}
	case "RuneProperties":
		if rune.RuneProperties != nil {
			res = PointerFieldInitStatement
			res = strings.ReplaceAll(res, "{{Identifier}}", rune.GongGetIdentifier(stage))
			res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "RuneProperties")
			res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", rune.RuneProperties.GongGetIdentifier(stage))
		}
	case "EnclosingParagraph":
		if rune.EnclosingParagraph != nil {
			res = PointerFieldInitStatement
			res = strings.ReplaceAll(res, "{{Identifier}}", rune.GongGetIdentifier(stage))
			res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "EnclosingParagraph")
			res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", rune.EnclosingParagraph.GongGetIdentifier(stage))
		}
	default:
		log.Panicf("Unknown field %s for Gongstruct Rune", fieldName)
	}
	return
}

func (runeproperties *RuneProperties) GongMarshallField(stage *Stage, fieldName string) (res string) {

	switch fieldName {
	case "Name":
		res = StringInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", runeproperties.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "Name")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", string(runeproperties.Name))
	case "IsBold":
		res = NumberInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", runeproperties.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "IsBold")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%t", runeproperties.IsBold))
	case "IsStrike":
		res = NumberInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", runeproperties.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "IsStrike")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%t", runeproperties.IsStrike))
	case "IsItalic":
		res = NumberInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", runeproperties.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "IsItalic")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%t", runeproperties.IsItalic))
	case "Content":
		res = StringInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", runeproperties.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "Content")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", string(runeproperties.Content))

	case "Node":
		if runeproperties.Node != nil {
			res = PointerFieldInitStatement
			res = strings.ReplaceAll(res, "{{Identifier}}", runeproperties.GongGetIdentifier(stage))
			res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "Node")
			res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", runeproperties.Node.GongGetIdentifier(stage))
		}
	default:
		log.Panicf("Unknown field %s for Gongstruct RuneProperties", fieldName)
	}
	return
}

func (table *Table) GongMarshallField(stage *Stage, fieldName string) (res string) {

	switch fieldName {
	case "Name":
		res = StringInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", table.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "Name")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", string(table.Name))
	case "Content":
		res = StringInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", table.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "Content")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", string(table.Content))

	case "Node":
		if table.Node != nil {
			res = PointerFieldInitStatement
			res = strings.ReplaceAll(res, "{{Identifier}}", table.GongGetIdentifier(stage))
			res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "Node")
			res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", table.Node.GongGetIdentifier(stage))
		}
	case "TableProperties":
		if table.TableProperties != nil {
			res = PointerFieldInitStatement
			res = strings.ReplaceAll(res, "{{Identifier}}", table.GongGetIdentifier(stage))
			res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "TableProperties")
			res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", table.TableProperties.GongGetIdentifier(stage))
		}
	case "TableRows":
		for _, _tablerow := range table.TableRows {
			tmp := SliceOfPointersFieldInitStatement
			tmp = strings.ReplaceAll(tmp, "{{Identifier}}", table.GongGetIdentifier(stage))
			tmp = strings.ReplaceAll(tmp, "{{GeneratedFieldName}}", "TableRows")
			tmp = strings.ReplaceAll(tmp, "{{GeneratedFieldNameValue}}", _tablerow.GongGetIdentifier(stage))
			res += tmp
		}
	default:
		log.Panicf("Unknown field %s for Gongstruct Table", fieldName)
	}
	return
}

func (tablecolumn *TableColumn) GongMarshallField(stage *Stage, fieldName string) (res string) {

	switch fieldName {
	case "Name":
		res = StringInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", tablecolumn.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "Name")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", string(tablecolumn.Name))
	case "Content":
		res = StringInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", tablecolumn.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "Content")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", string(tablecolumn.Content))

	case "Node":
		if tablecolumn.Node != nil {
			res = PointerFieldInitStatement
			res = strings.ReplaceAll(res, "{{Identifier}}", tablecolumn.GongGetIdentifier(stage))
			res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "Node")
			res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", tablecolumn.Node.GongGetIdentifier(stage))
		}
	case "Paragraphs":
		for _, _paragraph := range tablecolumn.Paragraphs {
			tmp := SliceOfPointersFieldInitStatement
			tmp = strings.ReplaceAll(tmp, "{{Identifier}}", tablecolumn.GongGetIdentifier(stage))
			tmp = strings.ReplaceAll(tmp, "{{GeneratedFieldName}}", "Paragraphs")
			tmp = strings.ReplaceAll(tmp, "{{GeneratedFieldNameValue}}", _paragraph.GongGetIdentifier(stage))
			res += tmp
		}
	default:
		log.Panicf("Unknown field %s for Gongstruct TableColumn", fieldName)
	}
	return
}

func (tableproperties *TableProperties) GongMarshallField(stage *Stage, fieldName string) (res string) {

	switch fieldName {
	case "Name":
		res = StringInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", tableproperties.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "Name")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", string(tableproperties.Name))
	case "Content":
		res = StringInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", tableproperties.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "Content")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", string(tableproperties.Content))

	case "Node":
		if tableproperties.Node != nil {
			res = PointerFieldInitStatement
			res = strings.ReplaceAll(res, "{{Identifier}}", tableproperties.GongGetIdentifier(stage))
			res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "Node")
			res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", tableproperties.Node.GongGetIdentifier(stage))
		}
	case "TableStyle":
		if tableproperties.TableStyle != nil {
			res = PointerFieldInitStatement
			res = strings.ReplaceAll(res, "{{Identifier}}", tableproperties.GongGetIdentifier(stage))
			res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "TableStyle")
			res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", tableproperties.TableStyle.GongGetIdentifier(stage))
		}
	default:
		log.Panicf("Unknown field %s for Gongstruct TableProperties", fieldName)
	}
	return
}

func (tablerow *TableRow) GongMarshallField(stage *Stage, fieldName string) (res string) {

	switch fieldName {
	case "Name":
		res = StringInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", tablerow.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "Name")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", string(tablerow.Name))
	case "Content":
		res = StringInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", tablerow.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "Content")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", string(tablerow.Content))

	case "Node":
		if tablerow.Node != nil {
			res = PointerFieldInitStatement
			res = strings.ReplaceAll(res, "{{Identifier}}", tablerow.GongGetIdentifier(stage))
			res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "Node")
			res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", tablerow.Node.GongGetIdentifier(stage))
		}
	case "TableColumns":
		for _, _tablecolumn := range tablerow.TableColumns {
			tmp := SliceOfPointersFieldInitStatement
			tmp = strings.ReplaceAll(tmp, "{{Identifier}}", tablerow.GongGetIdentifier(stage))
			tmp = strings.ReplaceAll(tmp, "{{GeneratedFieldName}}", "TableColumns")
			tmp = strings.ReplaceAll(tmp, "{{GeneratedFieldNameValue}}", _tablecolumn.GongGetIdentifier(stage))
			res += tmp
		}
	default:
		log.Panicf("Unknown field %s for Gongstruct TableRow", fieldName)
	}
	return
}

func (tablestyle *TableStyle) GongMarshallField(stage *Stage, fieldName string) (res string) {

	switch fieldName {
	case "Name":
		res = StringInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", tablestyle.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "Name")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", string(tablestyle.Name))
	case "Content":
		res = StringInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", tablestyle.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "Content")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", string(tablestyle.Content))
	case "Val":
		res = StringInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", tablestyle.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "Val")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", string(tablestyle.Val))

	case "Node":
		if tablestyle.Node != nil {
			res = PointerFieldInitStatement
			res = strings.ReplaceAll(res, "{{Identifier}}", tablestyle.GongGetIdentifier(stage))
			res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "Node")
			res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", tablestyle.Node.GongGetIdentifier(stage))
		}
	default:
		log.Panicf("Unknown field %s for Gongstruct TableStyle", fieldName)
	}
	return
}

func (text *Text) GongMarshallField(stage *Stage, fieldName string) (res string) {

	switch fieldName {
	case "Name":
		res = StringInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", text.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "Name")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", string(text.Name))
	case "Content":
		res = StringInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", text.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "Content")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", string(text.Content))
	case "PreserveWhiteSpace":
		res = NumberInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", text.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "PreserveWhiteSpace")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%t", text.PreserveWhiteSpace))

	case "Node":
		if text.Node != nil {
			res = PointerFieldInitStatement
			res = strings.ReplaceAll(res, "{{Identifier}}", text.GongGetIdentifier(stage))
			res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "Node")
			res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", text.Node.GongGetIdentifier(stage))
		}
	case "EnclosingRune":
		if text.EnclosingRune != nil {
			res = PointerFieldInitStatement
			res = strings.ReplaceAll(res, "{{Identifier}}", text.GongGetIdentifier(stage))
			res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "EnclosingRune")
			res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", text.EnclosingRune.GongGetIdentifier(stage))
		}
	default:
		log.Panicf("Unknown field %s for Gongstruct Text", fieldName)
	}
	return
}

// insertion point for marshall all fields methods
func (body *Body) GongMarshallAllFields(stage *Stage) (initializerStatements string, pointersInitializesStatements string) {

	initializerStatements += "\n"
	pointersInitializesStatements += "\n"
	{ // Insertion point for basic fields value assignment
		initializerStatements += body.GongMarshallField(stage, "Name")
		pointersInitializesStatements += body.GongMarshallField(stage, "Paragraphs")
		pointersInitializesStatements += body.GongMarshallField(stage, "Tables")
		pointersInitializesStatements += body.GongMarshallField(stage, "LastParagraph")
	}
	return
}
func (document *Document) GongMarshallAllFields(stage *Stage) (initializerStatements string, pointersInitializesStatements string) {

	initializerStatements += "\n"
	pointersInitializesStatements += "\n"
	{ // Insertion point for basic fields value assignment
		initializerStatements += document.GongMarshallField(stage, "Name")
		pointersInitializesStatements += document.GongMarshallField(stage, "File")
		pointersInitializesStatements += document.GongMarshallField(stage, "Root")
		pointersInitializesStatements += document.GongMarshallField(stage, "Body")
	}
	return
}
func (docx *Docx) GongMarshallAllFields(stage *Stage) (initializerStatements string, pointersInitializesStatements string) {

	initializerStatements += "\n"
	pointersInitializesStatements += "\n"
	{ // Insertion point for basic fields value assignment
		initializerStatements += docx.GongMarshallField(stage, "Name")
		pointersInitializesStatements += docx.GongMarshallField(stage, "Files")
		pointersInitializesStatements += docx.GongMarshallField(stage, "Document")
	}
	return
}
func (file *File) GongMarshallAllFields(stage *Stage) (initializerStatements string, pointersInitializesStatements string) {

	initializerStatements += "\n"
	pointersInitializesStatements += "\n"
	{ // Insertion point for basic fields value assignment
		initializerStatements += file.GongMarshallField(stage, "Name")
	}
	return
}
func (node *Node) GongMarshallAllFields(stage *Stage) (initializerStatements string, pointersInitializesStatements string) {

	initializerStatements += "\n"
	pointersInitializesStatements += "\n"
	{ // Insertion point for basic fields value assignment
		initializerStatements += node.GongMarshallField(stage, "Name")
		pointersInitializesStatements += node.GongMarshallField(stage, "Nodes")
	}
	return
}
func (paragraph *Paragraph) GongMarshallAllFields(stage *Stage) (initializerStatements string, pointersInitializesStatements string) {

	initializerStatements += "\n"
	pointersInitializesStatements += "\n"
	{ // Insertion point for basic fields value assignment
		initializerStatements += paragraph.GongMarshallField(stage, "Name")
		initializerStatements += paragraph.GongMarshallField(stage, "Content")
		pointersInitializesStatements += paragraph.GongMarshallField(stage, "Node")
		pointersInitializesStatements += paragraph.GongMarshallField(stage, "ParagraphProperties")
		pointersInitializesStatements += paragraph.GongMarshallField(stage, "Runes")
		initializerStatements += paragraph.GongMarshallField(stage, "Text")
		pointersInitializesStatements += paragraph.GongMarshallField(stage, "Next")
		pointersInitializesStatements += paragraph.GongMarshallField(stage, "Previous")
		pointersInitializesStatements += paragraph.GongMarshallField(stage, "EnclosingBody")
		pointersInitializesStatements += paragraph.GongMarshallField(stage, "EnclosingTableColumn")
	}
	return
}
func (paragraphproperties *ParagraphProperties) GongMarshallAllFields(stage *Stage) (initializerStatements string, pointersInitializesStatements string) {

	initializerStatements += "\n"
	pointersInitializesStatements += "\n"
	{ // Insertion point for basic fields value assignment
		initializerStatements += paragraphproperties.GongMarshallField(stage, "Name")
		initializerStatements += paragraphproperties.GongMarshallField(stage, "Content")
		pointersInitializesStatements += paragraphproperties.GongMarshallField(stage, "ParagraphStyle")
		pointersInitializesStatements += paragraphproperties.GongMarshallField(stage, "Node")
	}
	return
}
func (paragraphstyle *ParagraphStyle) GongMarshallAllFields(stage *Stage) (initializerStatements string, pointersInitializesStatements string) {

	initializerStatements += "\n"
	pointersInitializesStatements += "\n"
	{ // Insertion point for basic fields value assignment
		initializerStatements += paragraphstyle.GongMarshallField(stage, "Name")
		pointersInitializesStatements += paragraphstyle.GongMarshallField(stage, "Node")
		initializerStatements += paragraphstyle.GongMarshallField(stage, "Content")
		initializerStatements += paragraphstyle.GongMarshallField(stage, "ValAttr")
	}
	return
}
func (rune *Rune) GongMarshallAllFields(stage *Stage) (initializerStatements string, pointersInitializesStatements string) {

	initializerStatements += "\n"
	pointersInitializesStatements += "\n"
	{ // Insertion point for basic fields value assignment
		initializerStatements += rune.GongMarshallField(stage, "Name")
		initializerStatements += rune.GongMarshallField(stage, "Content")
		pointersInitializesStatements += rune.GongMarshallField(stage, "Node")
		pointersInitializesStatements += rune.GongMarshallField(stage, "Text")
		pointersInitializesStatements += rune.GongMarshallField(stage, "RuneProperties")
		pointersInitializesStatements += rune.GongMarshallField(stage, "EnclosingParagraph")
	}
	return
}
func (runeproperties *RuneProperties) GongMarshallAllFields(stage *Stage) (initializerStatements string, pointersInitializesStatements string) {

	initializerStatements += "\n"
	pointersInitializesStatements += "\n"
	{ // Insertion point for basic fields value assignment
		initializerStatements += runeproperties.GongMarshallField(stage, "Name")
		pointersInitializesStatements += runeproperties.GongMarshallField(stage, "Node")
		initializerStatements += runeproperties.GongMarshallField(stage, "IsBold")
		initializerStatements += runeproperties.GongMarshallField(stage, "IsStrike")
		initializerStatements += runeproperties.GongMarshallField(stage, "IsItalic")
		initializerStatements += runeproperties.GongMarshallField(stage, "Content")
	}
	return
}
func (table *Table) GongMarshallAllFields(stage *Stage) (initializerStatements string, pointersInitializesStatements string) {

	initializerStatements += "\n"
	pointersInitializesStatements += "\n"
	{ // Insertion point for basic fields value assignment
		initializerStatements += table.GongMarshallField(stage, "Name")
		pointersInitializesStatements += table.GongMarshallField(stage, "Node")
		initializerStatements += table.GongMarshallField(stage, "Content")
		pointersInitializesStatements += table.GongMarshallField(stage, "TableProperties")
		pointersInitializesStatements += table.GongMarshallField(stage, "TableRows")
	}
	return
}
func (tablecolumn *TableColumn) GongMarshallAllFields(stage *Stage) (initializerStatements string, pointersInitializesStatements string) {

	initializerStatements += "\n"
	pointersInitializesStatements += "\n"
	{ // Insertion point for basic fields value assignment
		initializerStatements += tablecolumn.GongMarshallField(stage, "Name")
		initializerStatements += tablecolumn.GongMarshallField(stage, "Content")
		pointersInitializesStatements += tablecolumn.GongMarshallField(stage, "Node")
		pointersInitializesStatements += tablecolumn.GongMarshallField(stage, "Paragraphs")
	}
	return
}
func (tableproperties *TableProperties) GongMarshallAllFields(stage *Stage) (initializerStatements string, pointersInitializesStatements string) {

	initializerStatements += "\n"
	pointersInitializesStatements += "\n"
	{ // Insertion point for basic fields value assignment
		initializerStatements += tableproperties.GongMarshallField(stage, "Name")
		pointersInitializesStatements += tableproperties.GongMarshallField(stage, "Node")
		initializerStatements += tableproperties.GongMarshallField(stage, "Content")
		pointersInitializesStatements += tableproperties.GongMarshallField(stage, "TableStyle")
	}
	return
}
func (tablerow *TableRow) GongMarshallAllFields(stage *Stage) (initializerStatements string, pointersInitializesStatements string) {

	initializerStatements += "\n"
	pointersInitializesStatements += "\n"
	{ // Insertion point for basic fields value assignment
		initializerStatements += tablerow.GongMarshallField(stage, "Name")
		initializerStatements += tablerow.GongMarshallField(stage, "Content")
		pointersInitializesStatements += tablerow.GongMarshallField(stage, "Node")
		pointersInitializesStatements += tablerow.GongMarshallField(stage, "TableColumns")
	}
	return
}
func (tablestyle *TableStyle) GongMarshallAllFields(stage *Stage) (initializerStatements string, pointersInitializesStatements string) {

	initializerStatements += "\n"
	pointersInitializesStatements += "\n"
	{ // Insertion point for basic fields value assignment
		initializerStatements += tablestyle.GongMarshallField(stage, "Name")
		pointersInitializesStatements += tablestyle.GongMarshallField(stage, "Node")
		initializerStatements += tablestyle.GongMarshallField(stage, "Content")
		initializerStatements += tablestyle.GongMarshallField(stage, "Val")
	}
	return
}
func (text *Text) GongMarshallAllFields(stage *Stage) (initializerStatements string, pointersInitializesStatements string) {

	initializerStatements += "\n"
	pointersInitializesStatements += "\n"
	{ // Insertion point for basic fields value assignment
		initializerStatements += text.GongMarshallField(stage, "Name")
		initializerStatements += text.GongMarshallField(stage, "Content")
		pointersInitializesStatements += text.GongMarshallField(stage, "Node")
		initializerStatements += text.GongMarshallField(stage, "PreserveWhiteSpace")
		pointersInitializesStatements += text.GongMarshallField(stage, "EnclosingRune")
	}
	return
}