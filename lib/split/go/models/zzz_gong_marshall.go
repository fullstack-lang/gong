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

	"github.com/sergi/go-diff/diffmatchpatch"
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

	const __write__local_time = "{{LocalTimeStamp}}"
	const __write__utc_time__ = "{{UTCTimeStamp}}"

	const __commitId__ = "{{CommitId}}"

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
		log.Fatalln(name + " is not a go filename")
	}

	log.Printf("Marshalling %s", name)
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
	map_AsSplit_Identifiers := make(map[*AsSplit]string)
	_ = map_AsSplit_Identifiers

	assplitOrdered := []*AsSplit{}
	for assplit := range stage.AsSplits {
		assplitOrdered = append(assplitOrdered, assplit)
	}
	sort.Slice(assplitOrdered[:], func(i, j int) bool {
		asspliti := assplitOrdered[i]
		assplitj := assplitOrdered[j]
		asspliti_order, oki := stage.AsSplitMap_Staged_Order[asspliti]
		assplitj_order, okj := stage.AsSplitMap_Staged_Order[assplitj]
		if !oki || !okj {
			log.Fatalln("unknown pointers")
		}
		return asspliti_order < assplitj_order
	})
	if len(assplitOrdered) > 0 {
		identifiersDecl += "\n"
	}
	for idx, assplit := range assplitOrdered {

		id = generatesIdentifier("AsSplit", idx, assplit.Name)
		map_AsSplit_Identifiers[assplit] = id

		decl = IdentifiersDecls
		decl = strings.ReplaceAll(decl, "{{Identifier}}", id)
		decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "AsSplit")
		decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", assplit.Name)
		identifiersDecl += decl

		initializerStatements += "\n"
		// Initialisation of values
		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Name")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(assplit.Name))
		initializerStatements += setValueField

		if assplit.Direction != "" {
			setValueField = StringEnumInitStatement
			setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
			setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Direction")
			setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", "models."+assplit.Direction.ToCodeString())
			initializerStatements += setValueField
		}

	}

	map_AsSplitArea_Identifiers := make(map[*AsSplitArea]string)
	_ = map_AsSplitArea_Identifiers

	assplitareaOrdered := []*AsSplitArea{}
	for assplitarea := range stage.AsSplitAreas {
		assplitareaOrdered = append(assplitareaOrdered, assplitarea)
	}
	sort.Slice(assplitareaOrdered[:], func(i, j int) bool {
		assplitareai := assplitareaOrdered[i]
		assplitareaj := assplitareaOrdered[j]
		assplitareai_order, oki := stage.AsSplitAreaMap_Staged_Order[assplitareai]
		assplitareaj_order, okj := stage.AsSplitAreaMap_Staged_Order[assplitareaj]
		if !oki || !okj {
			log.Fatalln("unknown pointers")
		}
		return assplitareai_order < assplitareaj_order
	})
	if len(assplitareaOrdered) > 0 {
		identifiersDecl += "\n"
	}
	for idx, assplitarea := range assplitareaOrdered {

		id = generatesIdentifier("AsSplitArea", idx, assplitarea.Name)
		map_AsSplitArea_Identifiers[assplitarea] = id

		decl = IdentifiersDecls
		decl = strings.ReplaceAll(decl, "{{Identifier}}", id)
		decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "AsSplitArea")
		decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", assplitarea.Name)
		identifiersDecl += decl

		initializerStatements += "\n"
		// Initialisation of values
		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Name")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(assplitarea.Name))
		initializerStatements += setValueField

		setValueField = NumberInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "ShowNameInHeader")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%t", assplitarea.ShowNameInHeader))
		initializerStatements += setValueField

		setValueField = NumberInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Size")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%f", assplitarea.Size))
		initializerStatements += setValueField

		setValueField = NumberInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "IsAny")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%t", assplitarea.IsAny))
		initializerStatements += setValueField

		setValueField = NumberInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "HasDiv")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%t", assplitarea.HasDiv))
		initializerStatements += setValueField

		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "DivStyle")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(assplitarea.DivStyle))
		initializerStatements += setValueField

	}

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
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "StackName")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(button.StackName))
		initializerStatements += setValueField

	}

	map_Cursor_Identifiers := make(map[*Cursor]string)
	_ = map_Cursor_Identifiers

	cursorOrdered := []*Cursor{}
	for cursor := range stage.Cursors {
		cursorOrdered = append(cursorOrdered, cursor)
	}
	sort.Slice(cursorOrdered[:], func(i, j int) bool {
		cursori := cursorOrdered[i]
		cursorj := cursorOrdered[j]
		cursori_order, oki := stage.CursorMap_Staged_Order[cursori]
		cursorj_order, okj := stage.CursorMap_Staged_Order[cursorj]
		if !oki || !okj {
			log.Fatalln("unknown pointers")
		}
		return cursori_order < cursorj_order
	})
	if len(cursorOrdered) > 0 {
		identifiersDecl += "\n"
	}
	for idx, cursor := range cursorOrdered {

		id = generatesIdentifier("Cursor", idx, cursor.Name)
		map_Cursor_Identifiers[cursor] = id

		decl = IdentifiersDecls
		decl = strings.ReplaceAll(decl, "{{Identifier}}", id)
		decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "Cursor")
		decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", cursor.Name)
		identifiersDecl += decl

		initializerStatements += "\n"
		// Initialisation of values
		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Name")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(cursor.Name))
		initializerStatements += setValueField

		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "StackName")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(cursor.StackName))
		initializerStatements += setValueField

		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Style")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(cursor.Style))
		initializerStatements += setValueField

	}

	map_Doc_Identifiers := make(map[*Doc]string)
	_ = map_Doc_Identifiers

	docOrdered := []*Doc{}
	for doc := range stage.Docs {
		docOrdered = append(docOrdered, doc)
	}
	sort.Slice(docOrdered[:], func(i, j int) bool {
		doci := docOrdered[i]
		docj := docOrdered[j]
		doci_order, oki := stage.DocMap_Staged_Order[doci]
		docj_order, okj := stage.DocMap_Staged_Order[docj]
		if !oki || !okj {
			log.Fatalln("unknown pointers")
		}
		return doci_order < docj_order
	})
	if len(docOrdered) > 0 {
		identifiersDecl += "\n"
	}
	for idx, doc := range docOrdered {

		id = generatesIdentifier("Doc", idx, doc.Name)
		map_Doc_Identifiers[doc] = id

		decl = IdentifiersDecls
		decl = strings.ReplaceAll(decl, "{{Identifier}}", id)
		decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "Doc")
		decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", doc.Name)
		identifiersDecl += decl

		initializerStatements += "\n"
		// Initialisation of values
		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Name")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(doc.Name))
		initializerStatements += setValueField

		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "StackName")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(doc.StackName))
		initializerStatements += setValueField

	}

	map_FavIcon_Identifiers := make(map[*FavIcon]string)
	_ = map_FavIcon_Identifiers

	faviconOrdered := []*FavIcon{}
	for favicon := range stage.FavIcons {
		faviconOrdered = append(faviconOrdered, favicon)
	}
	sort.Slice(faviconOrdered[:], func(i, j int) bool {
		faviconi := faviconOrdered[i]
		faviconj := faviconOrdered[j]
		faviconi_order, oki := stage.FavIconMap_Staged_Order[faviconi]
		faviconj_order, okj := stage.FavIconMap_Staged_Order[faviconj]
		if !oki || !okj {
			log.Fatalln("unknown pointers")
		}
		return faviconi_order < faviconj_order
	})
	if len(faviconOrdered) > 0 {
		identifiersDecl += "\n"
	}
	for idx, favicon := range faviconOrdered {

		id = generatesIdentifier("FavIcon", idx, favicon.Name)
		map_FavIcon_Identifiers[favicon] = id

		decl = IdentifiersDecls
		decl = strings.ReplaceAll(decl, "{{Identifier}}", id)
		decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "FavIcon")
		decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", favicon.Name)
		identifiersDecl += decl

		initializerStatements += "\n"
		// Initialisation of values
		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Name")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(favicon.Name))
		initializerStatements += setValueField

		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "SVG")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(favicon.SVG))
		initializerStatements += setValueField

	}

	map_Form_Identifiers := make(map[*Form]string)
	_ = map_Form_Identifiers

	formOrdered := []*Form{}
	for form := range stage.Forms {
		formOrdered = append(formOrdered, form)
	}
	sort.Slice(formOrdered[:], func(i, j int) bool {
		formi := formOrdered[i]
		formj := formOrdered[j]
		formi_order, oki := stage.FormMap_Staged_Order[formi]
		formj_order, okj := stage.FormMap_Staged_Order[formj]
		if !oki || !okj {
			log.Fatalln("unknown pointers")
		}
		return formi_order < formj_order
	})
	if len(formOrdered) > 0 {
		identifiersDecl += "\n"
	}
	for idx, form := range formOrdered {

		id = generatesIdentifier("Form", idx, form.Name)
		map_Form_Identifiers[form] = id

		decl = IdentifiersDecls
		decl = strings.ReplaceAll(decl, "{{Identifier}}", id)
		decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "Form")
		decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", form.Name)
		identifiersDecl += decl

		initializerStatements += "\n"
		// Initialisation of values
		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Name")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(form.Name))
		initializerStatements += setValueField

		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "StackName")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(form.StackName))
		initializerStatements += setValueField

		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "FormName")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(form.FormName))
		initializerStatements += setValueField

	}

	map_Load_Identifiers := make(map[*Load]string)
	_ = map_Load_Identifiers

	loadOrdered := []*Load{}
	for load := range stage.Loads {
		loadOrdered = append(loadOrdered, load)
	}
	sort.Slice(loadOrdered[:], func(i, j int) bool {
		loadi := loadOrdered[i]
		loadj := loadOrdered[j]
		loadi_order, oki := stage.LoadMap_Staged_Order[loadi]
		loadj_order, okj := stage.LoadMap_Staged_Order[loadj]
		if !oki || !okj {
			log.Fatalln("unknown pointers")
		}
		return loadi_order < loadj_order
	})
	if len(loadOrdered) > 0 {
		identifiersDecl += "\n"
	}
	for idx, load := range loadOrdered {

		id = generatesIdentifier("Load", idx, load.Name)
		map_Load_Identifiers[load] = id

		decl = IdentifiersDecls
		decl = strings.ReplaceAll(decl, "{{Identifier}}", id)
		decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "Load")
		decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", load.Name)
		identifiersDecl += decl

		initializerStatements += "\n"
		// Initialisation of values
		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Name")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(load.Name))
		initializerStatements += setValueField

		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "StackName")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(load.StackName))
		initializerStatements += setValueField

	}

	map_LogoOnTheLeft_Identifiers := make(map[*LogoOnTheLeft]string)
	_ = map_LogoOnTheLeft_Identifiers

	logoontheleftOrdered := []*LogoOnTheLeft{}
	for logoontheleft := range stage.LogoOnTheLefts {
		logoontheleftOrdered = append(logoontheleftOrdered, logoontheleft)
	}
	sort.Slice(logoontheleftOrdered[:], func(i, j int) bool {
		logoonthelefti := logoontheleftOrdered[i]
		logoontheleftj := logoontheleftOrdered[j]
		logoonthelefti_order, oki := stage.LogoOnTheLeftMap_Staged_Order[logoonthelefti]
		logoontheleftj_order, okj := stage.LogoOnTheLeftMap_Staged_Order[logoontheleftj]
		if !oki || !okj {
			log.Fatalln("unknown pointers")
		}
		return logoonthelefti_order < logoontheleftj_order
	})
	if len(logoontheleftOrdered) > 0 {
		identifiersDecl += "\n"
	}
	for idx, logoontheleft := range logoontheleftOrdered {

		id = generatesIdentifier("LogoOnTheLeft", idx, logoontheleft.Name)
		map_LogoOnTheLeft_Identifiers[logoontheleft] = id

		decl = IdentifiersDecls
		decl = strings.ReplaceAll(decl, "{{Identifier}}", id)
		decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "LogoOnTheLeft")
		decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", logoontheleft.Name)
		identifiersDecl += decl

		initializerStatements += "\n"
		// Initialisation of values
		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Name")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(logoontheleft.Name))
		initializerStatements += setValueField

		setValueField = NumberInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Width")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%d", logoontheleft.Width))
		initializerStatements += setValueField

		setValueField = NumberInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Height")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%d", logoontheleft.Height))
		initializerStatements += setValueField

		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "SVG")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(logoontheleft.SVG))
		initializerStatements += setValueField

	}

	map_LogoOnTheRight_Identifiers := make(map[*LogoOnTheRight]string)
	_ = map_LogoOnTheRight_Identifiers

	logoontherightOrdered := []*LogoOnTheRight{}
	for logoontheright := range stage.LogoOnTheRights {
		logoontherightOrdered = append(logoontherightOrdered, logoontheright)
	}
	sort.Slice(logoontherightOrdered[:], func(i, j int) bool {
		logoontherighti := logoontherightOrdered[i]
		logoontherightj := logoontherightOrdered[j]
		logoontherighti_order, oki := stage.LogoOnTheRightMap_Staged_Order[logoontherighti]
		logoontherightj_order, okj := stage.LogoOnTheRightMap_Staged_Order[logoontherightj]
		if !oki || !okj {
			log.Fatalln("unknown pointers")
		}
		return logoontherighti_order < logoontherightj_order
	})
	if len(logoontherightOrdered) > 0 {
		identifiersDecl += "\n"
	}
	for idx, logoontheright := range logoontherightOrdered {

		id = generatesIdentifier("LogoOnTheRight", idx, logoontheright.Name)
		map_LogoOnTheRight_Identifiers[logoontheright] = id

		decl = IdentifiersDecls
		decl = strings.ReplaceAll(decl, "{{Identifier}}", id)
		decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "LogoOnTheRight")
		decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", logoontheright.Name)
		identifiersDecl += decl

		initializerStatements += "\n"
		// Initialisation of values
		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Name")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(logoontheright.Name))
		initializerStatements += setValueField

		setValueField = NumberInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Width")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%d", logoontheright.Width))
		initializerStatements += setValueField

		setValueField = NumberInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Height")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%d", logoontheright.Height))
		initializerStatements += setValueField

		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "SVG")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(logoontheright.SVG))
		initializerStatements += setValueField

	}

	map_Slider_Identifiers := make(map[*Slider]string)
	_ = map_Slider_Identifiers

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
	for idx, slider := range sliderOrdered {

		id = generatesIdentifier("Slider", idx, slider.Name)
		map_Slider_Identifiers[slider] = id

		decl = IdentifiersDecls
		decl = strings.ReplaceAll(decl, "{{Identifier}}", id)
		decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "Slider")
		decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", slider.Name)
		identifiersDecl += decl

		initializerStatements += "\n"
		// Initialisation of values
		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Name")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(slider.Name))
		initializerStatements += setValueField

		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "StackName")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(slider.StackName))
		initializerStatements += setValueField

	}

	map_Split_Identifiers := make(map[*Split]string)
	_ = map_Split_Identifiers

	splitOrdered := []*Split{}
	for split := range stage.Splits {
		splitOrdered = append(splitOrdered, split)
	}
	sort.Slice(splitOrdered[:], func(i, j int) bool {
		spliti := splitOrdered[i]
		splitj := splitOrdered[j]
		spliti_order, oki := stage.SplitMap_Staged_Order[spliti]
		splitj_order, okj := stage.SplitMap_Staged_Order[splitj]
		if !oki || !okj {
			log.Fatalln("unknown pointers")
		}
		return spliti_order < splitj_order
	})
	if len(splitOrdered) > 0 {
		identifiersDecl += "\n"
	}
	for idx, split := range splitOrdered {

		id = generatesIdentifier("Split", idx, split.Name)
		map_Split_Identifiers[split] = id

		decl = IdentifiersDecls
		decl = strings.ReplaceAll(decl, "{{Identifier}}", id)
		decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "Split")
		decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", split.Name)
		identifiersDecl += decl

		initializerStatements += "\n"
		// Initialisation of values
		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Name")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(split.Name))
		initializerStatements += setValueField

		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "StackName")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(split.StackName))
		initializerStatements += setValueField

	}

	map_Svg_Identifiers := make(map[*Svg]string)
	_ = map_Svg_Identifiers

	svgOrdered := []*Svg{}
	for svg := range stage.Svgs {
		svgOrdered = append(svgOrdered, svg)
	}
	sort.Slice(svgOrdered[:], func(i, j int) bool {
		svgi := svgOrdered[i]
		svgj := svgOrdered[j]
		svgi_order, oki := stage.SvgMap_Staged_Order[svgi]
		svgj_order, okj := stage.SvgMap_Staged_Order[svgj]
		if !oki || !okj {
			log.Fatalln("unknown pointers")
		}
		return svgi_order < svgj_order
	})
	if len(svgOrdered) > 0 {
		identifiersDecl += "\n"
	}
	for idx, svg := range svgOrdered {

		id = generatesIdentifier("Svg", idx, svg.Name)
		map_Svg_Identifiers[svg] = id

		decl = IdentifiersDecls
		decl = strings.ReplaceAll(decl, "{{Identifier}}", id)
		decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "Svg")
		decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", svg.Name)
		identifiersDecl += decl

		initializerStatements += "\n"
		// Initialisation of values
		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Name")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(svg.Name))
		initializerStatements += setValueField

		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "StackName")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(svg.StackName))
		initializerStatements += setValueField

		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Style")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(svg.Style))
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
	for idx, table := range tableOrdered {

		id = generatesIdentifier("Table", idx, table.Name)
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
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "StackName")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(table.StackName))
		initializerStatements += setValueField

		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "TableName")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(table.TableName))
		initializerStatements += setValueField

	}

	map_Title_Identifiers := make(map[*Title]string)
	_ = map_Title_Identifiers

	titleOrdered := []*Title{}
	for title := range stage.Titles {
		titleOrdered = append(titleOrdered, title)
	}
	sort.Slice(titleOrdered[:], func(i, j int) bool {
		titlei := titleOrdered[i]
		titlej := titleOrdered[j]
		titlei_order, oki := stage.TitleMap_Staged_Order[titlei]
		titlej_order, okj := stage.TitleMap_Staged_Order[titlej]
		if !oki || !okj {
			log.Fatalln("unknown pointers")
		}
		return titlei_order < titlej_order
	})
	if len(titleOrdered) > 0 {
		identifiersDecl += "\n"
	}
	for idx, title := range titleOrdered {

		id = generatesIdentifier("Title", idx, title.Name)
		map_Title_Identifiers[title] = id

		decl = IdentifiersDecls
		decl = strings.ReplaceAll(decl, "{{Identifier}}", id)
		decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "Title")
		decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", title.Name)
		identifiersDecl += decl

		initializerStatements += "\n"
		// Initialisation of values
		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Name")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(title.Name))
		initializerStatements += setValueField

	}

	map_Tone_Identifiers := make(map[*Tone]string)
	_ = map_Tone_Identifiers

	toneOrdered := []*Tone{}
	for tone := range stage.Tones {
		toneOrdered = append(toneOrdered, tone)
	}
	sort.Slice(toneOrdered[:], func(i, j int) bool {
		tonei := toneOrdered[i]
		tonej := toneOrdered[j]
		tonei_order, oki := stage.ToneMap_Staged_Order[tonei]
		tonej_order, okj := stage.ToneMap_Staged_Order[tonej]
		if !oki || !okj {
			log.Fatalln("unknown pointers")
		}
		return tonei_order < tonej_order
	})
	if len(toneOrdered) > 0 {
		identifiersDecl += "\n"
	}
	for idx, tone := range toneOrdered {

		id = generatesIdentifier("Tone", idx, tone.Name)
		map_Tone_Identifiers[tone] = id

		decl = IdentifiersDecls
		decl = strings.ReplaceAll(decl, "{{Identifier}}", id)
		decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "Tone")
		decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", tone.Name)
		identifiersDecl += decl

		initializerStatements += "\n"
		// Initialisation of values
		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Name")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(tone.Name))
		initializerStatements += setValueField

		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "StackName")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(tone.StackName))
		initializerStatements += setValueField

	}

	map_Tree_Identifiers := make(map[*Tree]string)
	_ = map_Tree_Identifiers

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
	for idx, tree := range treeOrdered {

		id = generatesIdentifier("Tree", idx, tree.Name)
		map_Tree_Identifiers[tree] = id

		decl = IdentifiersDecls
		decl = strings.ReplaceAll(decl, "{{Identifier}}", id)
		decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "Tree")
		decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", tree.Name)
		identifiersDecl += decl

		initializerStatements += "\n"
		// Initialisation of values
		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Name")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(tree.Name))
		initializerStatements += setValueField

		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "StackName")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(tree.StackName))
		initializerStatements += setValueField

		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "TreeName")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(tree.TreeName))
		initializerStatements += setValueField

	}

	map_View_Identifiers := make(map[*View]string)
	_ = map_View_Identifiers

	viewOrdered := []*View{}
	for view := range stage.Views {
		viewOrdered = append(viewOrdered, view)
	}
	sort.Slice(viewOrdered[:], func(i, j int) bool {
		viewi := viewOrdered[i]
		viewj := viewOrdered[j]
		viewi_order, oki := stage.ViewMap_Staged_Order[viewi]
		viewj_order, okj := stage.ViewMap_Staged_Order[viewj]
		if !oki || !okj {
			log.Fatalln("unknown pointers")
		}
		return viewi_order < viewj_order
	})
	if len(viewOrdered) > 0 {
		identifiersDecl += "\n"
	}
	for idx, view := range viewOrdered {

		id = generatesIdentifier("View", idx, view.Name)
		map_View_Identifiers[view] = id

		decl = IdentifiersDecls
		decl = strings.ReplaceAll(decl, "{{Identifier}}", id)
		decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "View")
		decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", view.Name)
		identifiersDecl += decl

		initializerStatements += "\n"
		// Initialisation of values
		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Name")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(view.Name))
		initializerStatements += setValueField

		setValueField = NumberInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "ShowViewName")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%t", view.ShowViewName))
		initializerStatements += setValueField

		setValueField = NumberInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "IsSelectedView")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%t", view.IsSelectedView))
		initializerStatements += setValueField

	}

	map_Xlsx_Identifiers := make(map[*Xlsx]string)
	_ = map_Xlsx_Identifiers

	xlsxOrdered := []*Xlsx{}
	for xlsx := range stage.Xlsxs {
		xlsxOrdered = append(xlsxOrdered, xlsx)
	}
	sort.Slice(xlsxOrdered[:], func(i, j int) bool {
		xlsxi := xlsxOrdered[i]
		xlsxj := xlsxOrdered[j]
		xlsxi_order, oki := stage.XlsxMap_Staged_Order[xlsxi]
		xlsxj_order, okj := stage.XlsxMap_Staged_Order[xlsxj]
		if !oki || !okj {
			log.Fatalln("unknown pointers")
		}
		return xlsxi_order < xlsxj_order
	})
	if len(xlsxOrdered) > 0 {
		identifiersDecl += "\n"
	}
	for idx, xlsx := range xlsxOrdered {

		id = generatesIdentifier("Xlsx", idx, xlsx.Name)
		map_Xlsx_Identifiers[xlsx] = id

		decl = IdentifiersDecls
		decl = strings.ReplaceAll(decl, "{{Identifier}}", id)
		decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "Xlsx")
		decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", xlsx.Name)
		identifiersDecl += decl

		initializerStatements += "\n"
		// Initialisation of values
		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Name")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(xlsx.Name))
		initializerStatements += setValueField

		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "StackName")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(xlsx.StackName))
		initializerStatements += setValueField

	}

	// insertion initialization of objects to stage
	if len(assplitOrdered) > 0 {
		pointersInitializesStatements += "\n\t// setup of AsSplit instances pointers"
	}
	for idx, assplit := range assplitOrdered {
		var setPointerField string
		_ = setPointerField

		id = generatesIdentifier("AsSplit", idx, assplit.Name)
		map_AsSplit_Identifiers[assplit] = id

		// Initialisation of values
		for _, _assplitarea := range assplit.AsSplitAreas {
			setPointerField = SliceOfPointersFieldInitStatement
			setPointerField = strings.ReplaceAll(setPointerField, "{{Identifier}}", id)
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldName}}", "AsSplitAreas")
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldNameValue}}", map_AsSplitArea_Identifiers[_assplitarea])
			pointersInitializesStatements += setPointerField
		}

	}

	if len(assplitareaOrdered) > 0 {
		pointersInitializesStatements += "\n\t// setup of AsSplitArea instances pointers"
	}
	for idx, assplitarea := range assplitareaOrdered {
		var setPointerField string
		_ = setPointerField

		id = generatesIdentifier("AsSplitArea", idx, assplitarea.Name)
		map_AsSplitArea_Identifiers[assplitarea] = id

		// Initialisation of values
		if assplitarea.AsSplit != nil {
			setPointerField = PointerFieldInitStatement
			setPointerField = strings.ReplaceAll(setPointerField, "{{Identifier}}", id)
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldName}}", "AsSplit")
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldNameValue}}", map_AsSplit_Identifiers[assplitarea.AsSplit])
			pointersInitializesStatements += setPointerField
		}

		if assplitarea.Button != nil {
			setPointerField = PointerFieldInitStatement
			setPointerField = strings.ReplaceAll(setPointerField, "{{Identifier}}", id)
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldName}}", "Button")
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldNameValue}}", map_Button_Identifiers[assplitarea.Button])
			pointersInitializesStatements += setPointerField
		}

		if assplitarea.Cursor != nil {
			setPointerField = PointerFieldInitStatement
			setPointerField = strings.ReplaceAll(setPointerField, "{{Identifier}}", id)
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldName}}", "Cursor")
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldNameValue}}", map_Cursor_Identifiers[assplitarea.Cursor])
			pointersInitializesStatements += setPointerField
		}

		if assplitarea.Doc != nil {
			setPointerField = PointerFieldInitStatement
			setPointerField = strings.ReplaceAll(setPointerField, "{{Identifier}}", id)
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldName}}", "Doc")
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldNameValue}}", map_Doc_Identifiers[assplitarea.Doc])
			pointersInitializesStatements += setPointerField
		}

		if assplitarea.Form != nil {
			setPointerField = PointerFieldInitStatement
			setPointerField = strings.ReplaceAll(setPointerField, "{{Identifier}}", id)
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldName}}", "Form")
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldNameValue}}", map_Form_Identifiers[assplitarea.Form])
			pointersInitializesStatements += setPointerField
		}

		if assplitarea.Load != nil {
			setPointerField = PointerFieldInitStatement
			setPointerField = strings.ReplaceAll(setPointerField, "{{Identifier}}", id)
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldName}}", "Load")
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldNameValue}}", map_Load_Identifiers[assplitarea.Load])
			pointersInitializesStatements += setPointerField
		}

		if assplitarea.Slider != nil {
			setPointerField = PointerFieldInitStatement
			setPointerField = strings.ReplaceAll(setPointerField, "{{Identifier}}", id)
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldName}}", "Slider")
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldNameValue}}", map_Slider_Identifiers[assplitarea.Slider])
			pointersInitializesStatements += setPointerField
		}

		if assplitarea.Split != nil {
			setPointerField = PointerFieldInitStatement
			setPointerField = strings.ReplaceAll(setPointerField, "{{Identifier}}", id)
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldName}}", "Split")
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldNameValue}}", map_Split_Identifiers[assplitarea.Split])
			pointersInitializesStatements += setPointerField
		}

		if assplitarea.Svg != nil {
			setPointerField = PointerFieldInitStatement
			setPointerField = strings.ReplaceAll(setPointerField, "{{Identifier}}", id)
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldName}}", "Svg")
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldNameValue}}", map_Svg_Identifiers[assplitarea.Svg])
			pointersInitializesStatements += setPointerField
		}

		if assplitarea.Table != nil {
			setPointerField = PointerFieldInitStatement
			setPointerField = strings.ReplaceAll(setPointerField, "{{Identifier}}", id)
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldName}}", "Table")
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldNameValue}}", map_Table_Identifiers[assplitarea.Table])
			pointersInitializesStatements += setPointerField
		}

		if assplitarea.Tone != nil {
			setPointerField = PointerFieldInitStatement
			setPointerField = strings.ReplaceAll(setPointerField, "{{Identifier}}", id)
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldName}}", "Tone")
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldNameValue}}", map_Tone_Identifiers[assplitarea.Tone])
			pointersInitializesStatements += setPointerField
		}

		if assplitarea.Tree != nil {
			setPointerField = PointerFieldInitStatement
			setPointerField = strings.ReplaceAll(setPointerField, "{{Identifier}}", id)
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldName}}", "Tree")
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldNameValue}}", map_Tree_Identifiers[assplitarea.Tree])
			pointersInitializesStatements += setPointerField
		}

		if assplitarea.Xlsx != nil {
			setPointerField = PointerFieldInitStatement
			setPointerField = strings.ReplaceAll(setPointerField, "{{Identifier}}", id)
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldName}}", "Xlsx")
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldNameValue}}", map_Xlsx_Identifiers[assplitarea.Xlsx])
			pointersInitializesStatements += setPointerField
		}

	}

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

	if len(cursorOrdered) > 0 {
		pointersInitializesStatements += "\n\t// setup of Cursor instances pointers"
	}
	for idx, cursor := range cursorOrdered {
		var setPointerField string
		_ = setPointerField

		id = generatesIdentifier("Cursor", idx, cursor.Name)
		map_Cursor_Identifiers[cursor] = id

		// Initialisation of values
	}

	if len(docOrdered) > 0 {
		pointersInitializesStatements += "\n\t// setup of Doc instances pointers"
	}
	for idx, doc := range docOrdered {
		var setPointerField string
		_ = setPointerField

		id = generatesIdentifier("Doc", idx, doc.Name)
		map_Doc_Identifiers[doc] = id

		// Initialisation of values
	}

	if len(faviconOrdered) > 0 {
		pointersInitializesStatements += "\n\t// setup of FavIcon instances pointers"
	}
	for idx, favicon := range faviconOrdered {
		var setPointerField string
		_ = setPointerField

		id = generatesIdentifier("FavIcon", idx, favicon.Name)
		map_FavIcon_Identifiers[favicon] = id

		// Initialisation of values
	}

	if len(formOrdered) > 0 {
		pointersInitializesStatements += "\n\t// setup of Form instances pointers"
	}
	for idx, form := range formOrdered {
		var setPointerField string
		_ = setPointerField

		id = generatesIdentifier("Form", idx, form.Name)
		map_Form_Identifiers[form] = id

		// Initialisation of values
	}

	if len(loadOrdered) > 0 {
		pointersInitializesStatements += "\n\t// setup of Load instances pointers"
	}
	for idx, load := range loadOrdered {
		var setPointerField string
		_ = setPointerField

		id = generatesIdentifier("Load", idx, load.Name)
		map_Load_Identifiers[load] = id

		// Initialisation of values
	}

	if len(logoontheleftOrdered) > 0 {
		pointersInitializesStatements += "\n\t// setup of LogoOnTheLeft instances pointers"
	}
	for idx, logoontheleft := range logoontheleftOrdered {
		var setPointerField string
		_ = setPointerField

		id = generatesIdentifier("LogoOnTheLeft", idx, logoontheleft.Name)
		map_LogoOnTheLeft_Identifiers[logoontheleft] = id

		// Initialisation of values
	}

	if len(logoontherightOrdered) > 0 {
		pointersInitializesStatements += "\n\t// setup of LogoOnTheRight instances pointers"
	}
	for idx, logoontheright := range logoontherightOrdered {
		var setPointerField string
		_ = setPointerField

		id = generatesIdentifier("LogoOnTheRight", idx, logoontheright.Name)
		map_LogoOnTheRight_Identifiers[logoontheright] = id

		// Initialisation of values
	}

	if len(sliderOrdered) > 0 {
		pointersInitializesStatements += "\n\t// setup of Slider instances pointers"
	}
	for idx, slider := range sliderOrdered {
		var setPointerField string
		_ = setPointerField

		id = generatesIdentifier("Slider", idx, slider.Name)
		map_Slider_Identifiers[slider] = id

		// Initialisation of values
	}

	if len(splitOrdered) > 0 {
		pointersInitializesStatements += "\n\t// setup of Split instances pointers"
	}
	for idx, split := range splitOrdered {
		var setPointerField string
		_ = setPointerField

		id = generatesIdentifier("Split", idx, split.Name)
		map_Split_Identifiers[split] = id

		// Initialisation of values
	}

	if len(svgOrdered) > 0 {
		pointersInitializesStatements += "\n\t// setup of Svg instances pointers"
	}
	for idx, svg := range svgOrdered {
		var setPointerField string
		_ = setPointerField

		id = generatesIdentifier("Svg", idx, svg.Name)
		map_Svg_Identifiers[svg] = id

		// Initialisation of values
	}

	if len(tableOrdered) > 0 {
		pointersInitializesStatements += "\n\t// setup of Table instances pointers"
	}
	for idx, table := range tableOrdered {
		var setPointerField string
		_ = setPointerField

		id = generatesIdentifier("Table", idx, table.Name)
		map_Table_Identifiers[table] = id

		// Initialisation of values
	}

	if len(titleOrdered) > 0 {
		pointersInitializesStatements += "\n\t// setup of Title instances pointers"
	}
	for idx, title := range titleOrdered {
		var setPointerField string
		_ = setPointerField

		id = generatesIdentifier("Title", idx, title.Name)
		map_Title_Identifiers[title] = id

		// Initialisation of values
	}

	if len(toneOrdered) > 0 {
		pointersInitializesStatements += "\n\t// setup of Tone instances pointers"
	}
	for idx, tone := range toneOrdered {
		var setPointerField string
		_ = setPointerField

		id = generatesIdentifier("Tone", idx, tone.Name)
		map_Tone_Identifiers[tone] = id

		// Initialisation of values
	}

	if len(treeOrdered) > 0 {
		pointersInitializesStatements += "\n\t// setup of Tree instances pointers"
	}
	for idx, tree := range treeOrdered {
		var setPointerField string
		_ = setPointerField

		id = generatesIdentifier("Tree", idx, tree.Name)
		map_Tree_Identifiers[tree] = id

		// Initialisation of values
	}

	if len(viewOrdered) > 0 {
		pointersInitializesStatements += "\n\t// setup of View instances pointers"
	}
	for idx, view := range viewOrdered {
		var setPointerField string
		_ = setPointerField

		id = generatesIdentifier("View", idx, view.Name)
		map_View_Identifiers[view] = id

		// Initialisation of values
		for _, _assplitarea := range view.RootAsSplitAreas {
			setPointerField = SliceOfPointersFieldInitStatement
			setPointerField = strings.ReplaceAll(setPointerField, "{{Identifier}}", id)
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldName}}", "RootAsSplitAreas")
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldNameValue}}", map_AsSplitArea_Identifiers[_assplitarea])
			pointersInitializesStatements += setPointerField
		}

	}

	if len(xlsxOrdered) > 0 {
		pointersInitializesStatements += "\n\t// setup of Xlsx instances pointers"
	}
	for idx, xlsx := range xlsxOrdered {
		var setPointerField string
		_ = setPointerField

		id = generatesIdentifier("Xlsx", idx, xlsx.Name)
		map_Xlsx_Identifiers[xlsx] = id

		// Initialisation of values
	}

	res = strings.ReplaceAll(res, "{{Identifiers}}", identifiersDecl)
	res = strings.ReplaceAll(res, "{{ValueInitializers}}", initializerStatements)
	res = strings.ReplaceAll(res, "{{PointersInitializers}}", pointersInitializesStatements)

	// Local time with timezone
	localTimestamp := stage.commitTimeStamp.Format("2006-01-02 15:04:05.000000 MST")

	// UTC time
	utcTimestamp := stage.commitTimeStamp.UTC().Format("2006-01-02 15:04:05.000000 UTC")
	res = strings.ReplaceAll(res, "{{LocalTimeStamp}}", localTimestamp)
	res = strings.ReplaceAll(res, "{{UTCTimeStamp}}", utcTimestamp)
	res = strings.ReplaceAll(res, "{{CommitId}}", fmt.Sprintf("%.10d", stage.commitId))

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

	if stage.generatesDiff {
		diff := computeDiff(stage.contentWhenParsed, res)
		os.WriteFile(fmt.Sprintf("%s-%.10d-%.10d.delta", name, stage.commitIdWhenParsed, stage.commitId), []byte(diff), os.FileMode(0666))
		diff = ComputeDiff(stage.contentWhenParsed, res)
		os.WriteFile(fmt.Sprintf("%s-%.10d-%.10d.diff", name, stage.commitIdWhenParsed, stage.commitId), []byte(diff), os.FileMode(0666))
	}
	stage.contentWhenParsed = res
	stage.commitIdWhenParsed = stage.commitId

	fmt.Fprintln(file, res)
}

// computeDiff calculates the git-style unified diff between two strings.
func computeDiff(a, b string) string {
	dmp := diffmatchpatch.New()
	diffs := dmp.DiffMain(a, b, false)
	return dmp.DiffToDelta(diffs)
}

// computePrettyDiff calculates the git-style unified diff between two strings.
func computePrettyDiff(a, b string) string {
	dmp := diffmatchpatch.New()
	diffs := dmp.DiffMain(a, b, false)
	return dmp.DiffPrettyHtml(diffs)
}

// applyDiff reconstructs the original string 'a' from the new string 'b' and the diff string 'c'.
func applyDiff(b, c string) (string, error) {
	dmp := diffmatchpatch.New()
	diffs, err := dmp.DiffFromDelta(b, c)
	if err != nil {
		return "", err
	}
	patches := dmp.PatchMake(b, diffs)
	// We are applying the patch in reverse to get from 'b' to 'a'.
	// The library's PatchApply function returns the new string and a slice of booleans indicating the success of each patch application.
	result, _ := dmp.PatchApply(patches, b)
	return result, nil
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
