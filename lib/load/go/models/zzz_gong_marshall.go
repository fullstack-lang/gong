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
	filetodownloadOrdered := []*FileToDownload{}
	for filetodownload := range stage.FileToDownloads {
		filetodownloadOrdered = append(filetodownloadOrdered, filetodownload)
	}
	sort.Slice(filetodownloadOrdered[:], func(i, j int) bool {
		filetodownloadi := filetodownloadOrdered[i]
		filetodownloadj := filetodownloadOrdered[j]
		filetodownloadi_order, oki := stage.FileToDownloadMap_Staged_Order[filetodownloadi]
		filetodownloadj_order, okj := stage.FileToDownloadMap_Staged_Order[filetodownloadj]
		if !oki || !okj {
			log.Fatalln("unknown pointers")
		}
		return filetodownloadi_order < filetodownloadj_order
	})
	if len(filetodownloadOrdered) > 0 {
		identifiersDecl += "\n"
	}
	for _, filetodownload := range filetodownloadOrdered {

		identifiersDecl += filetodownload.GongMarshallIdentifier(stage)

		initializerStatements += "\n"
		// Insertion point for basic fields value assignment
		initializerStatements += filetodownload.GongMarshallField(stage, "Name")
		initializerStatements += filetodownload.GongMarshallField(stage, "Base64EncodedContent")
	}

	filetouploadOrdered := []*FileToUpload{}
	for filetoupload := range stage.FileToUploads {
		filetouploadOrdered = append(filetouploadOrdered, filetoupload)
	}
	sort.Slice(filetouploadOrdered[:], func(i, j int) bool {
		filetouploadi := filetouploadOrdered[i]
		filetouploadj := filetouploadOrdered[j]
		filetouploadi_order, oki := stage.FileToUploadMap_Staged_Order[filetouploadi]
		filetouploadj_order, okj := stage.FileToUploadMap_Staged_Order[filetouploadj]
		if !oki || !okj {
			log.Fatalln("unknown pointers")
		}
		return filetouploadi_order < filetouploadj_order
	})
	if len(filetouploadOrdered) > 0 {
		identifiersDecl += "\n"
	}
	for _, filetoupload := range filetouploadOrdered {

		identifiersDecl += filetoupload.GongMarshallIdentifier(stage)

		initializerStatements += "\n"
		// Insertion point for basic fields value assignment
		initializerStatements += filetoupload.GongMarshallField(stage, "Name")
		initializerStatements += filetoupload.GongMarshallField(stage, "Base64EncodedContent")
	}

	messageOrdered := []*Message{}
	for message := range stage.Messages {
		messageOrdered = append(messageOrdered, message)
	}
	sort.Slice(messageOrdered[:], func(i, j int) bool {
		messagei := messageOrdered[i]
		messagej := messageOrdered[j]
		messagei_order, oki := stage.MessageMap_Staged_Order[messagei]
		messagej_order, okj := stage.MessageMap_Staged_Order[messagej]
		if !oki || !okj {
			log.Fatalln("unknown pointers")
		}
		return messagei_order < messagej_order
	})
	if len(messageOrdered) > 0 {
		identifiersDecl += "\n"
	}
	for _, message := range messageOrdered {

		identifiersDecl += message.GongMarshallIdentifier(stage)

		initializerStatements += "\n"
		// Insertion point for basic fields value assignment
		initializerStatements += message.GongMarshallField(stage, "Name")
	}

	// insertion initialization of objects to stage
	for _, filetodownload := range filetodownloadOrdered {
		_ = filetodownload
		var setPointerField string
		_ = setPointerField

		// Insertion point for pointers initialization
	}

	for _, filetoupload := range filetouploadOrdered {
		_ = filetoupload
		var setPointerField string
		_ = setPointerField

		// Insertion point for pointers initialization
	}

	for _, message := range messageOrdered {
		_ = message
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
func (filetodownload *FileToDownload) GongMarshallField(stage *Stage, fieldName string) (res string) {

	switch fieldName {
	case "Name":
		res = StringInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", filetodownload.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "Name")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", string(filetodownload.Name))
	case "Base64EncodedContent":
		res = StringInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", filetodownload.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "Base64EncodedContent")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", string(filetodownload.Base64EncodedContent))

	default:
		log.Panicf("Unknown field %s for Gongstruct FileToDownload", fieldName)
	}
	return
}

func (filetoupload *FileToUpload) GongMarshallField(stage *Stage, fieldName string) (res string) {

	switch fieldName {
	case "Name":
		res = StringInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", filetoupload.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "Name")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", string(filetoupload.Name))
	case "Base64EncodedContent":
		res = StringInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", filetoupload.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "Base64EncodedContent")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", string(filetoupload.Base64EncodedContent))

	default:
		log.Panicf("Unknown field %s for Gongstruct FileToUpload", fieldName)
	}
	return
}

func (message *Message) GongMarshallField(stage *Stage, fieldName string) (res string) {

	switch fieldName {
	case "Name":
		res = StringInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", message.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "Name")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", string(message.Name))

	default:
		log.Panicf("Unknown field %s for Gongstruct Message", fieldName)
	}
	return
}

// insertion point for marshall all fields methods
func (filetodownload *FileToDownload) GongMarshallAllFields(stage *Stage) (initializerStatements string, pointersInitializesStatements string) {

	initializerStatements += "\n"
	pointersInitializesStatements += "\n"
	{ // Insertion point for basic fields value assignment
		initializerStatements += filetodownload.GongMarshallField(stage, "Name")
		initializerStatements += filetodownload.GongMarshallField(stage, "Base64EncodedContent")
	}
	return
}
func (filetoupload *FileToUpload) GongMarshallAllFields(stage *Stage) (initializerStatements string, pointersInitializesStatements string) {

	initializerStatements += "\n"
	pointersInitializesStatements += "\n"
	{ // Insertion point for basic fields value assignment
		initializerStatements += filetoupload.GongMarshallField(stage, "Name")
		initializerStatements += filetoupload.GongMarshallField(stage, "Base64EncodedContent")
	}
	return
}
func (message *Message) GongMarshallAllFields(stage *Stage) (initializerStatements string, pointersInitializesStatements string) {

	initializerStatements += "\n"
	pointersInitializesStatements += "\n"
	{ // Insertion point for basic fields value assignment
		initializerStatements += message.GongMarshallField(stage, "Name")
	}
	return
}