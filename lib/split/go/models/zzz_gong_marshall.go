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

// previous version does not hanldle embedded structs (https://github.com/golang/go/issues/9859)
// simpler version but the name of the instance cannot be human read before the fields initialization
const IdentifiersDeclsWithoutNameInit = `
	{{Identifier}} := (&models.{{GeneratedStructName}}{}).Stage(stage)`/* */

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
	for _, assplit := range assplitOrdered {
	
		identifiersDecl += assplit.GongMarshallIdentifier(stage)

		initializerStatements += "\n"
		// Initialisation of values
		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", assplit.GongGetIdentifier(stage))
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Name")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(assplit.Name))
		initializerStatements += setValueField

		if assplit.Direction != "" {
			setValueField = StringEnumInitStatement
			setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", assplit.GongGetIdentifier(stage))
			setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Direction")
			setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", "models."+assplit.Direction.ToCodeString())
			initializerStatements += setValueField
		}

	}

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
	for _, assplitarea := range assplitareaOrdered {
	
		identifiersDecl += assplitarea.GongMarshallIdentifier(stage)

		initializerStatements += "\n"
		// Initialisation of values
		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", assplitarea.GongGetIdentifier(stage))
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Name")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(assplitarea.Name))
		initializerStatements += setValueField

		setValueField = NumberInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", assplitarea.GongGetIdentifier(stage))
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "ShowNameInHeader")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%t", assplitarea.ShowNameInHeader))
		initializerStatements += setValueField

		setValueField = NumberInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", assplitarea.GongGetIdentifier(stage))
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Size")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%f", assplitarea.Size))
		initializerStatements += setValueField

		setValueField = NumberInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", assplitarea.GongGetIdentifier(stage))
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "IsAny")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%t", assplitarea.IsAny))
		initializerStatements += setValueField

		setValueField = NumberInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", assplitarea.GongGetIdentifier(stage))
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "HasDiv")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%t", assplitarea.HasDiv))
		initializerStatements += setValueField

		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", assplitarea.GongGetIdentifier(stage))
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "DivStyle")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(assplitarea.DivStyle))
		initializerStatements += setValueField

	}

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
		// Initialisation of values
		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", button.GongGetIdentifier(stage))
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Name")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(button.Name))
		initializerStatements += setValueField

		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", button.GongGetIdentifier(stage))
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "StackName")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(button.StackName))
		initializerStatements += setValueField

	}

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
	for _, cursor := range cursorOrdered {
	
		identifiersDecl += cursor.GongMarshallIdentifier(stage)

		initializerStatements += "\n"
		// Initialisation of values
		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", cursor.GongGetIdentifier(stage))
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Name")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(cursor.Name))
		initializerStatements += setValueField

		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", cursor.GongGetIdentifier(stage))
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "StackName")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(cursor.StackName))
		initializerStatements += setValueField

		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", cursor.GongGetIdentifier(stage))
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Style")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(cursor.Style))
		initializerStatements += setValueField

	}

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
	for _, favicon := range faviconOrdered {
	
		identifiersDecl += favicon.GongMarshallIdentifier(stage)

		initializerStatements += "\n"
		// Initialisation of values
		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", favicon.GongGetIdentifier(stage))
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Name")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(favicon.Name))
		initializerStatements += setValueField

		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", favicon.GongGetIdentifier(stage))
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "SVG")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(favicon.SVG))
		initializerStatements += setValueField

	}

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
	for _, form := range formOrdered {
	
		identifiersDecl += form.GongMarshallIdentifier(stage)

		initializerStatements += "\n"
		// Initialisation of values
		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", form.GongGetIdentifier(stage))
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Name")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(form.Name))
		initializerStatements += setValueField

		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", form.GongGetIdentifier(stage))
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "StackName")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(form.StackName))
		initializerStatements += setValueField

	}

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
	for _, load := range loadOrdered {
	
		identifiersDecl += load.GongMarshallIdentifier(stage)

		initializerStatements += "\n"
		// Initialisation of values
		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", load.GongGetIdentifier(stage))
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Name")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(load.Name))
		initializerStatements += setValueField

		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", load.GongGetIdentifier(stage))
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "StackName")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(load.StackName))
		initializerStatements += setValueField

	}

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
	for _, logoontheleft := range logoontheleftOrdered {
	
		identifiersDecl += logoontheleft.GongMarshallIdentifier(stage)

		initializerStatements += "\n"
		// Initialisation of values
		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", logoontheleft.GongGetIdentifier(stage))
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Name")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(logoontheleft.Name))
		initializerStatements += setValueField

		setValueField = NumberInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", logoontheleft.GongGetIdentifier(stage))
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Width")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%d", logoontheleft.Width))
		initializerStatements += setValueField

		setValueField = NumberInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", logoontheleft.GongGetIdentifier(stage))
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Height")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%d", logoontheleft.Height))
		initializerStatements += setValueField

		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", logoontheleft.GongGetIdentifier(stage))
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "SVG")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(logoontheleft.SVG))
		initializerStatements += setValueField

	}

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
	for _, logoontheright := range logoontherightOrdered {
	
		identifiersDecl += logoontheright.GongMarshallIdentifier(stage)

		initializerStatements += "\n"
		// Initialisation of values
		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", logoontheright.GongGetIdentifier(stage))
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Name")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(logoontheright.Name))
		initializerStatements += setValueField

		setValueField = NumberInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", logoontheright.GongGetIdentifier(stage))
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Width")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%d", logoontheright.Width))
		initializerStatements += setValueField

		setValueField = NumberInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", logoontheright.GongGetIdentifier(stage))
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Height")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%d", logoontheright.Height))
		initializerStatements += setValueField

		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", logoontheright.GongGetIdentifier(stage))
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "SVG")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(logoontheright.SVG))
		initializerStatements += setValueField

	}

	markdownOrdered := []*Markdown{}
	for markdown := range stage.Markdowns {
		markdownOrdered = append(markdownOrdered, markdown)
	}
	sort.Slice(markdownOrdered[:], func(i, j int) bool {
		markdowni := markdownOrdered[i]
		markdownj := markdownOrdered[j]
		markdowni_order, oki := stage.MarkdownMap_Staged_Order[markdowni]
		markdownj_order, okj := stage.MarkdownMap_Staged_Order[markdownj]
		if !oki || !okj {
			log.Fatalln("unknown pointers")
		}
		return markdowni_order < markdownj_order
	})
	if len(markdownOrdered) > 0 {
		identifiersDecl += "\n"
	}
	for _, markdown := range markdownOrdered {
	
		identifiersDecl += markdown.GongMarshallIdentifier(stage)

		initializerStatements += "\n"
		// Initialisation of values
		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", markdown.GongGetIdentifier(stage))
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Name")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(markdown.Name))
		initializerStatements += setValueField

		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", markdown.GongGetIdentifier(stage))
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "StackName")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(markdown.StackName))
		initializerStatements += setValueField

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
		// Initialisation of values
		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", slider.GongGetIdentifier(stage))
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Name")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(slider.Name))
		initializerStatements += setValueField

		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", slider.GongGetIdentifier(stage))
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "StackName")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(slider.StackName))
		initializerStatements += setValueField

	}

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
	for _, split := range splitOrdered {
	
		identifiersDecl += split.GongMarshallIdentifier(stage)

		initializerStatements += "\n"
		// Initialisation of values
		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", split.GongGetIdentifier(stage))
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Name")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(split.Name))
		initializerStatements += setValueField

		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", split.GongGetIdentifier(stage))
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "StackName")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(split.StackName))
		initializerStatements += setValueField

	}

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
	for _, svg := range svgOrdered {
	
		identifiersDecl += svg.GongMarshallIdentifier(stage)

		initializerStatements += "\n"
		// Initialisation of values
		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", svg.GongGetIdentifier(stage))
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Name")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(svg.Name))
		initializerStatements += setValueField

		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", svg.GongGetIdentifier(stage))
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "StackName")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(svg.StackName))
		initializerStatements += setValueField

		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", svg.GongGetIdentifier(stage))
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Style")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(svg.Style))
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
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "StackName")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(table.StackName))
		initializerStatements += setValueField

	}

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
	for _, title := range titleOrdered {
	
		identifiersDecl += title.GongMarshallIdentifier(stage)

		initializerStatements += "\n"
		// Initialisation of values
		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", title.GongGetIdentifier(stage))
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Name")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(title.Name))
		initializerStatements += setValueField

	}

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
	for _, tone := range toneOrdered {
	
		identifiersDecl += tone.GongMarshallIdentifier(stage)

		initializerStatements += "\n"
		// Initialisation of values
		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", tone.GongGetIdentifier(stage))
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Name")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(tone.Name))
		initializerStatements += setValueField

		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", tone.GongGetIdentifier(stage))
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "StackName")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(tone.StackName))
		initializerStatements += setValueField

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
		// Initialisation of values
		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", tree.GongGetIdentifier(stage))
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Name")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(tree.Name))
		initializerStatements += setValueField

		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", tree.GongGetIdentifier(stage))
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "StackName")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(tree.StackName))
		initializerStatements += setValueField

	}

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
	for _, view := range viewOrdered {
	
		identifiersDecl += view.GongMarshallIdentifier(stage)

		initializerStatements += "\n"
		// Initialisation of values
		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", view.GongGetIdentifier(stage))
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Name")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(view.Name))
		initializerStatements += setValueField

		setValueField = NumberInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", view.GongGetIdentifier(stage))
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "ShowViewName")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%t", view.ShowViewName))
		initializerStatements += setValueField

		setValueField = NumberInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", view.GongGetIdentifier(stage))
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "IsSelectedView")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%t", view.IsSelectedView))
		initializerStatements += setValueField

		if view.Direction != "" {
			setValueField = StringEnumInitStatement
			setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", view.GongGetIdentifier(stage))
			setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Direction")
			setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", "models."+view.Direction.ToCodeString())
			initializerStatements += setValueField
		}

	}

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
	for _, xlsx := range xlsxOrdered {
	
		identifiersDecl += xlsx.GongMarshallIdentifier(stage)

		initializerStatements += "\n"
		// Initialisation of values
		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", xlsx.GongGetIdentifier(stage))
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Name")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(xlsx.Name))
		initializerStatements += setValueField

		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", xlsx.GongGetIdentifier(stage))
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "StackName")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(xlsx.StackName))
		initializerStatements += setValueField

	}

	// insertion initialization of objects to stage
	if len(assplitOrdered) > 0 {
		pointersInitializesStatements += "\n\t// setup of AsSplit instances pointers"
	}
	for _, assplit := range assplitOrdered {
		_ = assplit
		var setPointerField string
		_ = setPointerField

		// Initialisation of values
		for _, _assplitarea := range assplit.AsSplitAreas {
			setPointerField = SliceOfPointersFieldInitStatement
			setPointerField = strings.ReplaceAll(setPointerField, "{{Identifier}}", assplit.GongGetIdentifier(stage))
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldName}}", "AsSplitAreas")
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldNameValue}}", _assplitarea.GongGetIdentifier(stage))
			pointersInitializesStatements += setPointerField
		}

	}

	if len(assplitareaOrdered) > 0 {
		pointersInitializesStatements += "\n\t// setup of AsSplitArea instances pointers"
	}
	for _, assplitarea := range assplitareaOrdered {
		_ = assplitarea
		var setPointerField string
		_ = setPointerField

		// Initialisation of values
		if assplitarea.AsSplit != nil {
			setPointerField = PointerFieldInitStatement
			setPointerField = strings.ReplaceAll(setPointerField, "{{Identifier}}", assplitarea.GongGetIdentifier(stage))
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldName}}", "AsSplit")
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldNameValue}}", assplitarea.AsSplit.GongGetIdentifier(stage))
			pointersInitializesStatements += setPointerField
		}

		if assplitarea.Button != nil {
			setPointerField = PointerFieldInitStatement
			setPointerField = strings.ReplaceAll(setPointerField, "{{Identifier}}", assplitarea.GongGetIdentifier(stage))
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldName}}", "Button")
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldNameValue}}", assplitarea.Button.GongGetIdentifier(stage))
			pointersInitializesStatements += setPointerField
		}

		if assplitarea.Cursor != nil {
			setPointerField = PointerFieldInitStatement
			setPointerField = strings.ReplaceAll(setPointerField, "{{Identifier}}", assplitarea.GongGetIdentifier(stage))
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldName}}", "Cursor")
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldNameValue}}", assplitarea.Cursor.GongGetIdentifier(stage))
			pointersInitializesStatements += setPointerField
		}

		if assplitarea.Form != nil {
			setPointerField = PointerFieldInitStatement
			setPointerField = strings.ReplaceAll(setPointerField, "{{Identifier}}", assplitarea.GongGetIdentifier(stage))
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldName}}", "Form")
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldNameValue}}", assplitarea.Form.GongGetIdentifier(stage))
			pointersInitializesStatements += setPointerField
		}

		if assplitarea.Load != nil {
			setPointerField = PointerFieldInitStatement
			setPointerField = strings.ReplaceAll(setPointerField, "{{Identifier}}", assplitarea.GongGetIdentifier(stage))
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldName}}", "Load")
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldNameValue}}", assplitarea.Load.GongGetIdentifier(stage))
			pointersInitializesStatements += setPointerField
		}

		if assplitarea.Markdown != nil {
			setPointerField = PointerFieldInitStatement
			setPointerField = strings.ReplaceAll(setPointerField, "{{Identifier}}", assplitarea.GongGetIdentifier(stage))
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldName}}", "Markdown")
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldNameValue}}", assplitarea.Markdown.GongGetIdentifier(stage))
			pointersInitializesStatements += setPointerField
		}

		if assplitarea.Slider != nil {
			setPointerField = PointerFieldInitStatement
			setPointerField = strings.ReplaceAll(setPointerField, "{{Identifier}}", assplitarea.GongGetIdentifier(stage))
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldName}}", "Slider")
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldNameValue}}", assplitarea.Slider.GongGetIdentifier(stage))
			pointersInitializesStatements += setPointerField
		}

		if assplitarea.Split != nil {
			setPointerField = PointerFieldInitStatement
			setPointerField = strings.ReplaceAll(setPointerField, "{{Identifier}}", assplitarea.GongGetIdentifier(stage))
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldName}}", "Split")
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldNameValue}}", assplitarea.Split.GongGetIdentifier(stage))
			pointersInitializesStatements += setPointerField
		}

		if assplitarea.Svg != nil {
			setPointerField = PointerFieldInitStatement
			setPointerField = strings.ReplaceAll(setPointerField, "{{Identifier}}", assplitarea.GongGetIdentifier(stage))
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldName}}", "Svg")
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldNameValue}}", assplitarea.Svg.GongGetIdentifier(stage))
			pointersInitializesStatements += setPointerField
		}

		if assplitarea.Table != nil {
			setPointerField = PointerFieldInitStatement
			setPointerField = strings.ReplaceAll(setPointerField, "{{Identifier}}", assplitarea.GongGetIdentifier(stage))
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldName}}", "Table")
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldNameValue}}", assplitarea.Table.GongGetIdentifier(stage))
			pointersInitializesStatements += setPointerField
		}

		if assplitarea.Tone != nil {
			setPointerField = PointerFieldInitStatement
			setPointerField = strings.ReplaceAll(setPointerField, "{{Identifier}}", assplitarea.GongGetIdentifier(stage))
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldName}}", "Tone")
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldNameValue}}", assplitarea.Tone.GongGetIdentifier(stage))
			pointersInitializesStatements += setPointerField
		}

		if assplitarea.Tree != nil {
			setPointerField = PointerFieldInitStatement
			setPointerField = strings.ReplaceAll(setPointerField, "{{Identifier}}", assplitarea.GongGetIdentifier(stage))
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldName}}", "Tree")
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldNameValue}}", assplitarea.Tree.GongGetIdentifier(stage))
			pointersInitializesStatements += setPointerField
		}

		if assplitarea.Xlsx != nil {
			setPointerField = PointerFieldInitStatement
			setPointerField = strings.ReplaceAll(setPointerField, "{{Identifier}}", assplitarea.GongGetIdentifier(stage))
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldName}}", "Xlsx")
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldNameValue}}", assplitarea.Xlsx.GongGetIdentifier(stage))
			pointersInitializesStatements += setPointerField
		}

	}

	if len(buttonOrdered) > 0 {
		pointersInitializesStatements += "\n\t// setup of Button instances pointers"
	}
	for _, button := range buttonOrdered {
		_ = button
		var setPointerField string
		_ = setPointerField

		// Initialisation of values
	}

	if len(cursorOrdered) > 0 {
		pointersInitializesStatements += "\n\t// setup of Cursor instances pointers"
	}
	for _, cursor := range cursorOrdered {
		_ = cursor
		var setPointerField string
		_ = setPointerField

		// Initialisation of values
	}

	if len(faviconOrdered) > 0 {
		pointersInitializesStatements += "\n\t// setup of FavIcon instances pointers"
	}
	for _, favicon := range faviconOrdered {
		_ = favicon
		var setPointerField string
		_ = setPointerField

		// Initialisation of values
	}

	if len(formOrdered) > 0 {
		pointersInitializesStatements += "\n\t// setup of Form instances pointers"
	}
	for _, form := range formOrdered {
		_ = form
		var setPointerField string
		_ = setPointerField

		// Initialisation of values
	}

	if len(loadOrdered) > 0 {
		pointersInitializesStatements += "\n\t// setup of Load instances pointers"
	}
	for _, load := range loadOrdered {
		_ = load
		var setPointerField string
		_ = setPointerField

		// Initialisation of values
	}

	if len(logoontheleftOrdered) > 0 {
		pointersInitializesStatements += "\n\t// setup of LogoOnTheLeft instances pointers"
	}
	for _, logoontheleft := range logoontheleftOrdered {
		_ = logoontheleft
		var setPointerField string
		_ = setPointerField

		// Initialisation of values
	}

	if len(logoontherightOrdered) > 0 {
		pointersInitializesStatements += "\n\t// setup of LogoOnTheRight instances pointers"
	}
	for _, logoontheright := range logoontherightOrdered {
		_ = logoontheright
		var setPointerField string
		_ = setPointerField

		// Initialisation of values
	}

	if len(markdownOrdered) > 0 {
		pointersInitializesStatements += "\n\t// setup of Markdown instances pointers"
	}
	for _, markdown := range markdownOrdered {
		_ = markdown
		var setPointerField string
		_ = setPointerField

		// Initialisation of values
	}

	if len(sliderOrdered) > 0 {
		pointersInitializesStatements += "\n\t// setup of Slider instances pointers"
	}
	for _, slider := range sliderOrdered {
		_ = slider
		var setPointerField string
		_ = setPointerField

		// Initialisation of values
	}

	if len(splitOrdered) > 0 {
		pointersInitializesStatements += "\n\t// setup of Split instances pointers"
	}
	for _, split := range splitOrdered {
		_ = split
		var setPointerField string
		_ = setPointerField

		// Initialisation of values
	}

	if len(svgOrdered) > 0 {
		pointersInitializesStatements += "\n\t// setup of Svg instances pointers"
	}
	for _, svg := range svgOrdered {
		_ = svg
		var setPointerField string
		_ = setPointerField

		// Initialisation of values
	}

	if len(tableOrdered) > 0 {
		pointersInitializesStatements += "\n\t// setup of Table instances pointers"
	}
	for _, table := range tableOrdered {
		_ = table
		var setPointerField string
		_ = setPointerField

		// Initialisation of values
	}

	if len(titleOrdered) > 0 {
		pointersInitializesStatements += "\n\t// setup of Title instances pointers"
	}
	for _, title := range titleOrdered {
		_ = title
		var setPointerField string
		_ = setPointerField

		// Initialisation of values
	}

	if len(toneOrdered) > 0 {
		pointersInitializesStatements += "\n\t// setup of Tone instances pointers"
	}
	for _, tone := range toneOrdered {
		_ = tone
		var setPointerField string
		_ = setPointerField

		// Initialisation of values
	}

	if len(treeOrdered) > 0 {
		pointersInitializesStatements += "\n\t// setup of Tree instances pointers"
	}
	for _, tree := range treeOrdered {
		_ = tree
		var setPointerField string
		_ = setPointerField

		// Initialisation of values
	}

	if len(viewOrdered) > 0 {
		pointersInitializesStatements += "\n\t// setup of View instances pointers"
	}
	for _, view := range viewOrdered {
		_ = view
		var setPointerField string
		_ = setPointerField

		// Initialisation of values
		for _, _assplitarea := range view.RootAsSplitAreas {
			setPointerField = SliceOfPointersFieldInitStatement
			setPointerField = strings.ReplaceAll(setPointerField, "{{Identifier}}", view.GongGetIdentifier(stage))
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldName}}", "RootAsSplitAreas")
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldNameValue}}", _assplitarea.GongGetIdentifier(stage))
			pointersInitializesStatements += setPointerField
		}

	}

	if len(xlsxOrdered) > 0 {
		pointersInitializesStatements += "\n\t// setup of Xlsx instances pointers"
	}
	for _, xlsx := range xlsxOrdered {
		_ = xlsx
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

// insertion initialization of objects to stage
func (assplit *AsSplit) GongMarshallField(stage *Stage, fieldName string) (setValueField, setPointerField string) {
	initializerStatements := ""
	_ = initializerStatements
	pointersInitializesStatements := ""
	_ = pointersInitializesStatements

	switch fieldName {
	case "Name":
		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", assplit.GongGetIdentifier(stage))
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Name")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(assplit.Name))
		initializerStatements += setValueField
	case "Direction":
		if assplit.Direction != "" {
			setValueField = StringEnumInitStatement
			setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", assplit.GongGetIdentifier(stage))
			setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Direction")
			setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", "models."+assplit.Direction.ToCodeString())
			initializerStatements += setValueField
		}

	case "AsSplitAreas":
		for _, _assplitarea := range assplit.AsSplitAreas {
			setPointerField = SliceOfPointersFieldInitStatement
			setPointerField = strings.ReplaceAll(setPointerField, "{{Identifier}}", assplit.GongGetIdentifier(stage))
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldName}}", "AsSplitAreas")
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldNameValue}}", _assplitarea.GongGetIdentifier(stage))
			pointersInitializesStatements += setPointerField
		}

	default:
		log.Panicf("Unknown field %s for Gongstruct AsSplit", fieldName)
	}
	return
}

func (assplitarea *AsSplitArea) GongMarshallField(stage *Stage, fieldName string) (setValueField, setPointerField string) {
	initializerStatements := ""
	_ = initializerStatements
	pointersInitializesStatements := ""
	_ = pointersInitializesStatements

	switch fieldName {
	case "Name":
		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", assplitarea.GongGetIdentifier(stage))
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Name")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(assplitarea.Name))
		initializerStatements += setValueField
	case "ShowNameInHeader":
		setValueField = NumberInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", assplitarea.GongGetIdentifier(stage))
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "ShowNameInHeader")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%t", assplitarea.ShowNameInHeader))
		initializerStatements += setValueField
	case "Size":
		setValueField = NumberInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", assplitarea.GongGetIdentifier(stage))
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Size")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%f", assplitarea.Size))
		initializerStatements += setValueField
	case "IsAny":
		setValueField = NumberInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", assplitarea.GongGetIdentifier(stage))
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "IsAny")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%t", assplitarea.IsAny))
		initializerStatements += setValueField
	case "HasDiv":
		setValueField = NumberInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", assplitarea.GongGetIdentifier(stage))
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "HasDiv")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%t", assplitarea.HasDiv))
		initializerStatements += setValueField
	case "DivStyle":
		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", assplitarea.GongGetIdentifier(stage))
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "DivStyle")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(assplitarea.DivStyle))
		initializerStatements += setValueField

	case "AsSplit":
		if assplitarea.AsSplit != nil {
			setPointerField = PointerFieldInitStatement
			setPointerField = strings.ReplaceAll(setPointerField, "{{Identifier}}", assplitarea.GongGetIdentifier(stage))
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldName}}", "AsSplit")
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldNameValue}}", assplitarea.AsSplit.GongGetIdentifier(stage))
			pointersInitializesStatements += setPointerField
		}
	case "Button":
		if assplitarea.Button != nil {
			setPointerField = PointerFieldInitStatement
			setPointerField = strings.ReplaceAll(setPointerField, "{{Identifier}}", assplitarea.GongGetIdentifier(stage))
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldName}}", "Button")
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldNameValue}}", assplitarea.Button.GongGetIdentifier(stage))
			pointersInitializesStatements += setPointerField
		}
	case "Cursor":
		if assplitarea.Cursor != nil {
			setPointerField = PointerFieldInitStatement
			setPointerField = strings.ReplaceAll(setPointerField, "{{Identifier}}", assplitarea.GongGetIdentifier(stage))
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldName}}", "Cursor")
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldNameValue}}", assplitarea.Cursor.GongGetIdentifier(stage))
			pointersInitializesStatements += setPointerField
		}
	case "Form":
		if assplitarea.Form != nil {
			setPointerField = PointerFieldInitStatement
			setPointerField = strings.ReplaceAll(setPointerField, "{{Identifier}}", assplitarea.GongGetIdentifier(stage))
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldName}}", "Form")
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldNameValue}}", assplitarea.Form.GongGetIdentifier(stage))
			pointersInitializesStatements += setPointerField
		}
	case "Load":
		if assplitarea.Load != nil {
			setPointerField = PointerFieldInitStatement
			setPointerField = strings.ReplaceAll(setPointerField, "{{Identifier}}", assplitarea.GongGetIdentifier(stage))
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldName}}", "Load")
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldNameValue}}", assplitarea.Load.GongGetIdentifier(stage))
			pointersInitializesStatements += setPointerField
		}
	case "Markdown":
		if assplitarea.Markdown != nil {
			setPointerField = PointerFieldInitStatement
			setPointerField = strings.ReplaceAll(setPointerField, "{{Identifier}}", assplitarea.GongGetIdentifier(stage))
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldName}}", "Markdown")
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldNameValue}}", assplitarea.Markdown.GongGetIdentifier(stage))
			pointersInitializesStatements += setPointerField
		}
	case "Slider":
		if assplitarea.Slider != nil {
			setPointerField = PointerFieldInitStatement
			setPointerField = strings.ReplaceAll(setPointerField, "{{Identifier}}", assplitarea.GongGetIdentifier(stage))
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldName}}", "Slider")
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldNameValue}}", assplitarea.Slider.GongGetIdentifier(stage))
			pointersInitializesStatements += setPointerField
		}
	case "Split":
		if assplitarea.Split != nil {
			setPointerField = PointerFieldInitStatement
			setPointerField = strings.ReplaceAll(setPointerField, "{{Identifier}}", assplitarea.GongGetIdentifier(stage))
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldName}}", "Split")
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldNameValue}}", assplitarea.Split.GongGetIdentifier(stage))
			pointersInitializesStatements += setPointerField
		}
	case "Svg":
		if assplitarea.Svg != nil {
			setPointerField = PointerFieldInitStatement
			setPointerField = strings.ReplaceAll(setPointerField, "{{Identifier}}", assplitarea.GongGetIdentifier(stage))
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldName}}", "Svg")
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldNameValue}}", assplitarea.Svg.GongGetIdentifier(stage))
			pointersInitializesStatements += setPointerField
		}
	case "Table":
		if assplitarea.Table != nil {
			setPointerField = PointerFieldInitStatement
			setPointerField = strings.ReplaceAll(setPointerField, "{{Identifier}}", assplitarea.GongGetIdentifier(stage))
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldName}}", "Table")
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldNameValue}}", assplitarea.Table.GongGetIdentifier(stage))
			pointersInitializesStatements += setPointerField
		}
	case "Tone":
		if assplitarea.Tone != nil {
			setPointerField = PointerFieldInitStatement
			setPointerField = strings.ReplaceAll(setPointerField, "{{Identifier}}", assplitarea.GongGetIdentifier(stage))
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldName}}", "Tone")
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldNameValue}}", assplitarea.Tone.GongGetIdentifier(stage))
			pointersInitializesStatements += setPointerField
		}
	case "Tree":
		if assplitarea.Tree != nil {
			setPointerField = PointerFieldInitStatement
			setPointerField = strings.ReplaceAll(setPointerField, "{{Identifier}}", assplitarea.GongGetIdentifier(stage))
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldName}}", "Tree")
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldNameValue}}", assplitarea.Tree.GongGetIdentifier(stage))
			pointersInitializesStatements += setPointerField
		}
	case "Xlsx":
		if assplitarea.Xlsx != nil {
			setPointerField = PointerFieldInitStatement
			setPointerField = strings.ReplaceAll(setPointerField, "{{Identifier}}", assplitarea.GongGetIdentifier(stage))
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldName}}", "Xlsx")
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldNameValue}}", assplitarea.Xlsx.GongGetIdentifier(stage))
			pointersInitializesStatements += setPointerField
		}

	default:
		log.Panicf("Unknown field %s for Gongstruct AsSplitArea", fieldName)
	}
	return
}

func (button *Button) GongMarshallField(stage *Stage, fieldName string) (setValueField, setPointerField string) {
	initializerStatements := ""
	_ = initializerStatements
	pointersInitializesStatements := ""
	_ = pointersInitializesStatements

	switch fieldName {
	case "Name":
		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", button.GongGetIdentifier(stage))
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Name")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(button.Name))
		initializerStatements += setValueField
	case "StackName":
		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", button.GongGetIdentifier(stage))
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "StackName")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(button.StackName))
		initializerStatements += setValueField


	default:
		log.Panicf("Unknown field %s for Gongstruct Button", fieldName)
	}
	return
}

func (cursor *Cursor) GongMarshallField(stage *Stage, fieldName string) (setValueField, setPointerField string) {
	initializerStatements := ""
	_ = initializerStatements
	pointersInitializesStatements := ""
	_ = pointersInitializesStatements

	switch fieldName {
	case "Name":
		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", cursor.GongGetIdentifier(stage))
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Name")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(cursor.Name))
		initializerStatements += setValueField
	case "StackName":
		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", cursor.GongGetIdentifier(stage))
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "StackName")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(cursor.StackName))
		initializerStatements += setValueField
	case "Style":
		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", cursor.GongGetIdentifier(stage))
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Style")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(cursor.Style))
		initializerStatements += setValueField


	default:
		log.Panicf("Unknown field %s for Gongstruct Cursor", fieldName)
	}
	return
}

func (favicon *FavIcon) GongMarshallField(stage *Stage, fieldName string) (setValueField, setPointerField string) {
	initializerStatements := ""
	_ = initializerStatements
	pointersInitializesStatements := ""
	_ = pointersInitializesStatements

	switch fieldName {
	case "Name":
		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", favicon.GongGetIdentifier(stage))
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Name")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(favicon.Name))
		initializerStatements += setValueField
	case "SVG":
		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", favicon.GongGetIdentifier(stage))
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "SVG")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(favicon.SVG))
		initializerStatements += setValueField


	default:
		log.Panicf("Unknown field %s for Gongstruct FavIcon", fieldName)
	}
	return
}

func (form *Form) GongMarshallField(stage *Stage, fieldName string) (setValueField, setPointerField string) {
	initializerStatements := ""
	_ = initializerStatements
	pointersInitializesStatements := ""
	_ = pointersInitializesStatements

	switch fieldName {
	case "Name":
		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", form.GongGetIdentifier(stage))
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Name")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(form.Name))
		initializerStatements += setValueField
	case "StackName":
		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", form.GongGetIdentifier(stage))
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "StackName")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(form.StackName))
		initializerStatements += setValueField


	default:
		log.Panicf("Unknown field %s for Gongstruct Form", fieldName)
	}
	return
}

func (load *Load) GongMarshallField(stage *Stage, fieldName string) (setValueField, setPointerField string) {
	initializerStatements := ""
	_ = initializerStatements
	pointersInitializesStatements := ""
	_ = pointersInitializesStatements

	switch fieldName {
	case "Name":
		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", load.GongGetIdentifier(stage))
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Name")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(load.Name))
		initializerStatements += setValueField
	case "StackName":
		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", load.GongGetIdentifier(stage))
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "StackName")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(load.StackName))
		initializerStatements += setValueField


	default:
		log.Panicf("Unknown field %s for Gongstruct Load", fieldName)
	}
	return
}

func (logoontheleft *LogoOnTheLeft) GongMarshallField(stage *Stage, fieldName string) (setValueField, setPointerField string) {
	initializerStatements := ""
	_ = initializerStatements
	pointersInitializesStatements := ""
	_ = pointersInitializesStatements

	switch fieldName {
	case "Name":
		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", logoontheleft.GongGetIdentifier(stage))
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Name")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(logoontheleft.Name))
		initializerStatements += setValueField
	case "Width":
		setValueField = NumberInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", logoontheleft.GongGetIdentifier(stage))
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Width")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%d", logoontheleft.Width))
		initializerStatements += setValueField
	case "Height":
		setValueField = NumberInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", logoontheleft.GongGetIdentifier(stage))
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Height")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%d", logoontheleft.Height))
		initializerStatements += setValueField
	case "SVG":
		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", logoontheleft.GongGetIdentifier(stage))
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "SVG")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(logoontheleft.SVG))
		initializerStatements += setValueField


	default:
		log.Panicf("Unknown field %s for Gongstruct LogoOnTheLeft", fieldName)
	}
	return
}

func (logoontheright *LogoOnTheRight) GongMarshallField(stage *Stage, fieldName string) (setValueField, setPointerField string) {
	initializerStatements := ""
	_ = initializerStatements
	pointersInitializesStatements := ""
	_ = pointersInitializesStatements

	switch fieldName {
	case "Name":
		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", logoontheright.GongGetIdentifier(stage))
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Name")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(logoontheright.Name))
		initializerStatements += setValueField
	case "Width":
		setValueField = NumberInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", logoontheright.GongGetIdentifier(stage))
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Width")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%d", logoontheright.Width))
		initializerStatements += setValueField
	case "Height":
		setValueField = NumberInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", logoontheright.GongGetIdentifier(stage))
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Height")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%d", logoontheright.Height))
		initializerStatements += setValueField
	case "SVG":
		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", logoontheright.GongGetIdentifier(stage))
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "SVG")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(logoontheright.SVG))
		initializerStatements += setValueField


	default:
		log.Panicf("Unknown field %s for Gongstruct LogoOnTheRight", fieldName)
	}
	return
}

func (markdown *Markdown) GongMarshallField(stage *Stage, fieldName string) (setValueField, setPointerField string) {
	initializerStatements := ""
	_ = initializerStatements
	pointersInitializesStatements := ""
	_ = pointersInitializesStatements

	switch fieldName {
	case "Name":
		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", markdown.GongGetIdentifier(stage))
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Name")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(markdown.Name))
		initializerStatements += setValueField
	case "StackName":
		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", markdown.GongGetIdentifier(stage))
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "StackName")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(markdown.StackName))
		initializerStatements += setValueField


	default:
		log.Panicf("Unknown field %s for Gongstruct Markdown", fieldName)
	}
	return
}

func (slider *Slider) GongMarshallField(stage *Stage, fieldName string) (setValueField, setPointerField string) {
	initializerStatements := ""
	_ = initializerStatements
	pointersInitializesStatements := ""
	_ = pointersInitializesStatements

	switch fieldName {
	case "Name":
		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", slider.GongGetIdentifier(stage))
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Name")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(slider.Name))
		initializerStatements += setValueField
	case "StackName":
		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", slider.GongGetIdentifier(stage))
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "StackName")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(slider.StackName))
		initializerStatements += setValueField


	default:
		log.Panicf("Unknown field %s for Gongstruct Slider", fieldName)
	}
	return
}

func (split *Split) GongMarshallField(stage *Stage, fieldName string) (setValueField, setPointerField string) {
	initializerStatements := ""
	_ = initializerStatements
	pointersInitializesStatements := ""
	_ = pointersInitializesStatements

	switch fieldName {
	case "Name":
		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", split.GongGetIdentifier(stage))
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Name")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(split.Name))
		initializerStatements += setValueField
	case "StackName":
		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", split.GongGetIdentifier(stage))
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "StackName")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(split.StackName))
		initializerStatements += setValueField


	default:
		log.Panicf("Unknown field %s for Gongstruct Split", fieldName)
	}
	return
}

func (svg *Svg) GongMarshallField(stage *Stage, fieldName string) (setValueField, setPointerField string) {
	initializerStatements := ""
	_ = initializerStatements
	pointersInitializesStatements := ""
	_ = pointersInitializesStatements

	switch fieldName {
	case "Name":
		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", svg.GongGetIdentifier(stage))
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Name")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(svg.Name))
		initializerStatements += setValueField
	case "StackName":
		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", svg.GongGetIdentifier(stage))
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "StackName")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(svg.StackName))
		initializerStatements += setValueField
	case "Style":
		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", svg.GongGetIdentifier(stage))
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Style")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(svg.Style))
		initializerStatements += setValueField


	default:
		log.Panicf("Unknown field %s for Gongstruct Svg", fieldName)
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
	case "StackName":
		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", table.GongGetIdentifier(stage))
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "StackName")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(table.StackName))
		initializerStatements += setValueField


	default:
		log.Panicf("Unknown field %s for Gongstruct Table", fieldName)
	}
	return
}

func (title *Title) GongMarshallField(stage *Stage, fieldName string) (setValueField, setPointerField string) {
	initializerStatements := ""
	_ = initializerStatements
	pointersInitializesStatements := ""
	_ = pointersInitializesStatements

	switch fieldName {
	case "Name":
		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", title.GongGetIdentifier(stage))
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Name")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(title.Name))
		initializerStatements += setValueField


	default:
		log.Panicf("Unknown field %s for Gongstruct Title", fieldName)
	}
	return
}

func (tone *Tone) GongMarshallField(stage *Stage, fieldName string) (setValueField, setPointerField string) {
	initializerStatements := ""
	_ = initializerStatements
	pointersInitializesStatements := ""
	_ = pointersInitializesStatements

	switch fieldName {
	case "Name":
		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", tone.GongGetIdentifier(stage))
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Name")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(tone.Name))
		initializerStatements += setValueField
	case "StackName":
		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", tone.GongGetIdentifier(stage))
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "StackName")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(tone.StackName))
		initializerStatements += setValueField


	default:
		log.Panicf("Unknown field %s for Gongstruct Tone", fieldName)
	}
	return
}

func (tree *Tree) GongMarshallField(stage *Stage, fieldName string) (setValueField, setPointerField string) {
	initializerStatements := ""
	_ = initializerStatements
	pointersInitializesStatements := ""
	_ = pointersInitializesStatements

	switch fieldName {
	case "Name":
		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", tree.GongGetIdentifier(stage))
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Name")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(tree.Name))
		initializerStatements += setValueField
	case "StackName":
		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", tree.GongGetIdentifier(stage))
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "StackName")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(tree.StackName))
		initializerStatements += setValueField


	default:
		log.Panicf("Unknown field %s for Gongstruct Tree", fieldName)
	}
	return
}

func (view *View) GongMarshallField(stage *Stage, fieldName string) (setValueField, setPointerField string) {
	initializerStatements := ""
	_ = initializerStatements
	pointersInitializesStatements := ""
	_ = pointersInitializesStatements

	switch fieldName {
	case "Name":
		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", view.GongGetIdentifier(stage))
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Name")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(view.Name))
		initializerStatements += setValueField
	case "ShowViewName":
		setValueField = NumberInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", view.GongGetIdentifier(stage))
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "ShowViewName")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%t", view.ShowViewName))
		initializerStatements += setValueField
	case "IsSelectedView":
		setValueField = NumberInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", view.GongGetIdentifier(stage))
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "IsSelectedView")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%t", view.IsSelectedView))
		initializerStatements += setValueField
	case "Direction":
		if view.Direction != "" {
			setValueField = StringEnumInitStatement
			setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", view.GongGetIdentifier(stage))
			setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Direction")
			setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", "models."+view.Direction.ToCodeString())
			initializerStatements += setValueField
		}

	case "RootAsSplitAreas":
		for _, _assplitarea := range view.RootAsSplitAreas {
			setPointerField = SliceOfPointersFieldInitStatement
			setPointerField = strings.ReplaceAll(setPointerField, "{{Identifier}}", view.GongGetIdentifier(stage))
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldName}}", "RootAsSplitAreas")
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldNameValue}}", _assplitarea.GongGetIdentifier(stage))
			pointersInitializesStatements += setPointerField
		}

	default:
		log.Panicf("Unknown field %s for Gongstruct View", fieldName)
	}
	return
}

func (xlsx *Xlsx) GongMarshallField(stage *Stage, fieldName string) (setValueField, setPointerField string) {
	initializerStatements := ""
	_ = initializerStatements
	pointersInitializesStatements := ""
	_ = pointersInitializesStatements

	switch fieldName {
	case "Name":
		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", xlsx.GongGetIdentifier(stage))
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Name")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(xlsx.Name))
		initializerStatements += setValueField
	case "StackName":
		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", xlsx.GongGetIdentifier(stage))
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "StackName")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(xlsx.StackName))
		initializerStatements += setValueField


	default:
		log.Panicf("Unknown field %s for Gongstruct Xlsx", fieldName)
	}
	return
}
