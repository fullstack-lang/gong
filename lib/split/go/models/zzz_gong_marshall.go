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
	var identifiersDecl strings.Builder
	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder

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
		identifiersDecl.WriteString("\n")
	}
	for _, assplit := range assplitOrdered {

		identifiersDecl.WriteString(assplit.GongMarshallIdentifier(stage))

		initializerStatements.WriteString("\n")
		// Insertion point for basic fields value assignment
		initializerStatements.WriteString(assplit.GongMarshallField(stage, "Name"))
		initializerStatements.WriteString(assplit.GongMarshallField(stage, "Direction"))
		pointersInitializesStatements.WriteString(assplit.GongMarshallField(stage, "AsSplitAreas"))
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
		identifiersDecl.WriteString("\n")
	}
	for _, assplitarea := range assplitareaOrdered {

		identifiersDecl.WriteString(assplitarea.GongMarshallIdentifier(stage))

		initializerStatements.WriteString("\n")
		// Insertion point for basic fields value assignment
		initializerStatements.WriteString(assplitarea.GongMarshallField(stage, "Name"))
		initializerStatements.WriteString(assplitarea.GongMarshallField(stage, "ShowNameInHeader"))
		initializerStatements.WriteString(assplitarea.GongMarshallField(stage, "Size"))
		initializerStatements.WriteString(assplitarea.GongMarshallField(stage, "IsAny"))
		pointersInitializesStatements.WriteString(assplitarea.GongMarshallField(stage, "AsSplit"))
		pointersInitializesStatements.WriteString(assplitarea.GongMarshallField(stage, "Button"))
		pointersInitializesStatements.WriteString(assplitarea.GongMarshallField(stage, "Cursor"))
		pointersInitializesStatements.WriteString(assplitarea.GongMarshallField(stage, "Form"))
		pointersInitializesStatements.WriteString(assplitarea.GongMarshallField(stage, "Load"))
		pointersInitializesStatements.WriteString(assplitarea.GongMarshallField(stage, "Markdown"))
		pointersInitializesStatements.WriteString(assplitarea.GongMarshallField(stage, "Slider"))
		pointersInitializesStatements.WriteString(assplitarea.GongMarshallField(stage, "Split"))
		pointersInitializesStatements.WriteString(assplitarea.GongMarshallField(stage, "Svg"))
		pointersInitializesStatements.WriteString(assplitarea.GongMarshallField(stage, "Table"))
		pointersInitializesStatements.WriteString(assplitarea.GongMarshallField(stage, "Tone"))
		pointersInitializesStatements.WriteString(assplitarea.GongMarshallField(stage, "Tree"))
		pointersInitializesStatements.WriteString(assplitarea.GongMarshallField(stage, "Xlsx"))
		initializerStatements.WriteString(assplitarea.GongMarshallField(stage, "HasDiv"))
		initializerStatements.WriteString(assplitarea.GongMarshallField(stage, "DivStyle"))
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
		identifiersDecl.WriteString("\n")
	}
	for _, button := range buttonOrdered {

		identifiersDecl.WriteString(button.GongMarshallIdentifier(stage))

		initializerStatements.WriteString("\n")
		// Insertion point for basic fields value assignment
		initializerStatements.WriteString(button.GongMarshallField(stage, "Name"))
		initializerStatements.WriteString(button.GongMarshallField(stage, "StackName"))
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
		identifiersDecl.WriteString("\n")
	}
	for _, cursor := range cursorOrdered {

		identifiersDecl.WriteString(cursor.GongMarshallIdentifier(stage))

		initializerStatements.WriteString("\n")
		// Insertion point for basic fields value assignment
		initializerStatements.WriteString(cursor.GongMarshallField(stage, "Name"))
		initializerStatements.WriteString(cursor.GongMarshallField(stage, "StackName"))
		initializerStatements.WriteString(cursor.GongMarshallField(stage, "Style"))
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
		identifiersDecl.WriteString("\n")
	}
	for _, favicon := range faviconOrdered {

		identifiersDecl.WriteString(favicon.GongMarshallIdentifier(stage))

		initializerStatements.WriteString("\n")
		// Insertion point for basic fields value assignment
		initializerStatements.WriteString(favicon.GongMarshallField(stage, "Name"))
		initializerStatements.WriteString(favicon.GongMarshallField(stage, "SVG"))
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
		identifiersDecl.WriteString("\n")
	}
	for _, form := range formOrdered {

		identifiersDecl.WriteString(form.GongMarshallIdentifier(stage))

		initializerStatements.WriteString("\n")
		// Insertion point for basic fields value assignment
		initializerStatements.WriteString(form.GongMarshallField(stage, "Name"))
		initializerStatements.WriteString(form.GongMarshallField(stage, "StackName"))
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
		identifiersDecl.WriteString("\n")
	}
	for _, load := range loadOrdered {

		identifiersDecl.WriteString(load.GongMarshallIdentifier(stage))

		initializerStatements.WriteString("\n")
		// Insertion point for basic fields value assignment
		initializerStatements.WriteString(load.GongMarshallField(stage, "Name"))
		initializerStatements.WriteString(load.GongMarshallField(stage, "StackName"))
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
		identifiersDecl.WriteString("\n")
	}
	for _, logoontheleft := range logoontheleftOrdered {

		identifiersDecl.WriteString(logoontheleft.GongMarshallIdentifier(stage))

		initializerStatements.WriteString("\n")
		// Insertion point for basic fields value assignment
		initializerStatements.WriteString(logoontheleft.GongMarshallField(stage, "Name"))
		initializerStatements.WriteString(logoontheleft.GongMarshallField(stage, "Width"))
		initializerStatements.WriteString(logoontheleft.GongMarshallField(stage, "Height"))
		initializerStatements.WriteString(logoontheleft.GongMarshallField(stage, "SVG"))
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
		identifiersDecl.WriteString("\n")
	}
	for _, logoontheright := range logoontherightOrdered {

		identifiersDecl.WriteString(logoontheright.GongMarshallIdentifier(stage))

		initializerStatements.WriteString("\n")
		// Insertion point for basic fields value assignment
		initializerStatements.WriteString(logoontheright.GongMarshallField(stage, "Name"))
		initializerStatements.WriteString(logoontheright.GongMarshallField(stage, "Width"))
		initializerStatements.WriteString(logoontheright.GongMarshallField(stage, "Height"))
		initializerStatements.WriteString(logoontheright.GongMarshallField(stage, "SVG"))
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
		identifiersDecl.WriteString("\n")
	}
	for _, markdown := range markdownOrdered {

		identifiersDecl.WriteString(markdown.GongMarshallIdentifier(stage))

		initializerStatements.WriteString("\n")
		// Insertion point for basic fields value assignment
		initializerStatements.WriteString(markdown.GongMarshallField(stage, "Name"))
		initializerStatements.WriteString(markdown.GongMarshallField(stage, "StackName"))
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
		identifiersDecl.WriteString("\n")
	}
	for _, slider := range sliderOrdered {

		identifiersDecl.WriteString(slider.GongMarshallIdentifier(stage))

		initializerStatements.WriteString("\n")
		// Insertion point for basic fields value assignment
		initializerStatements.WriteString(slider.GongMarshallField(stage, "Name"))
		initializerStatements.WriteString(slider.GongMarshallField(stage, "StackName"))
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
		identifiersDecl.WriteString("\n")
	}
	for _, split := range splitOrdered {

		identifiersDecl.WriteString(split.GongMarshallIdentifier(stage))

		initializerStatements.WriteString("\n")
		// Insertion point for basic fields value assignment
		initializerStatements.WriteString(split.GongMarshallField(stage, "Name"))
		initializerStatements.WriteString(split.GongMarshallField(stage, "StackName"))
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
		identifiersDecl.WriteString("\n")
	}
	for _, svg := range svgOrdered {

		identifiersDecl.WriteString(svg.GongMarshallIdentifier(stage))

		initializerStatements.WriteString("\n")
		// Insertion point for basic fields value assignment
		initializerStatements.WriteString(svg.GongMarshallField(stage, "Name"))
		initializerStatements.WriteString(svg.GongMarshallField(stage, "StackName"))
		initializerStatements.WriteString(svg.GongMarshallField(stage, "Style"))
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
		identifiersDecl.WriteString("\n")
	}
	for _, table := range tableOrdered {

		identifiersDecl.WriteString(table.GongMarshallIdentifier(stage))

		initializerStatements.WriteString("\n")
		// Insertion point for basic fields value assignment
		initializerStatements.WriteString(table.GongMarshallField(stage, "Name"))
		initializerStatements.WriteString(table.GongMarshallField(stage, "StackName"))
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
		identifiersDecl.WriteString("\n")
	}
	for _, title := range titleOrdered {

		identifiersDecl.WriteString(title.GongMarshallIdentifier(stage))

		initializerStatements.WriteString("\n")
		// Insertion point for basic fields value assignment
		initializerStatements.WriteString(title.GongMarshallField(stage, "Name"))
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
		identifiersDecl.WriteString("\n")
	}
	for _, tone := range toneOrdered {

		identifiersDecl.WriteString(tone.GongMarshallIdentifier(stage))

		initializerStatements.WriteString("\n")
		// Insertion point for basic fields value assignment
		initializerStatements.WriteString(tone.GongMarshallField(stage, "Name"))
		initializerStatements.WriteString(tone.GongMarshallField(stage, "StackName"))
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
		identifiersDecl.WriteString("\n")
	}
	for _, tree := range treeOrdered {

		identifiersDecl.WriteString(tree.GongMarshallIdentifier(stage))

		initializerStatements.WriteString("\n")
		// Insertion point for basic fields value assignment
		initializerStatements.WriteString(tree.GongMarshallField(stage, "Name"))
		initializerStatements.WriteString(tree.GongMarshallField(stage, "StackName"))
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
		identifiersDecl.WriteString("\n")
	}
	for _, view := range viewOrdered {

		identifiersDecl.WriteString(view.GongMarshallIdentifier(stage))

		initializerStatements.WriteString("\n")
		// Insertion point for basic fields value assignment
		initializerStatements.WriteString(view.GongMarshallField(stage, "Name"))
		initializerStatements.WriteString(view.GongMarshallField(stage, "ShowViewName"))
		pointersInitializesStatements.WriteString(view.GongMarshallField(stage, "RootAsSplitAreas"))
		initializerStatements.WriteString(view.GongMarshallField(stage, "IsSelectedView"))
		initializerStatements.WriteString(view.GongMarshallField(stage, "Direction"))
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
		identifiersDecl.WriteString("\n")
	}
	for _, xlsx := range xlsxOrdered {

		identifiersDecl.WriteString(xlsx.GongMarshallIdentifier(stage))

		initializerStatements.WriteString("\n")
		// Insertion point for basic fields value assignment
		initializerStatements.WriteString(xlsx.GongMarshallField(stage, "Name"))
		initializerStatements.WriteString(xlsx.GongMarshallField(stage, "StackName"))
	}

	// insertion initialization of objects to stage
	for _, assplit := range assplitOrdered {
		_ = assplit
		var setPointerField string
		_ = setPointerField

		// Insertion point for pointers initialization
	}

	for _, assplitarea := range assplitareaOrdered {
		_ = assplitarea
		var setPointerField string
		_ = setPointerField

		// Insertion point for pointers initialization
	}

	for _, button := range buttonOrdered {
		_ = button
		var setPointerField string
		_ = setPointerField

		// Insertion point for pointers initialization
	}

	for _, cursor := range cursorOrdered {
		_ = cursor
		var setPointerField string
		_ = setPointerField

		// Insertion point for pointers initialization
	}

	for _, favicon := range faviconOrdered {
		_ = favicon
		var setPointerField string
		_ = setPointerField

		// Insertion point for pointers initialization
	}

	for _, form := range formOrdered {
		_ = form
		var setPointerField string
		_ = setPointerField

		// Insertion point for pointers initialization
	}

	for _, load := range loadOrdered {
		_ = load
		var setPointerField string
		_ = setPointerField

		// Insertion point for pointers initialization
	}

	for _, logoontheleft := range logoontheleftOrdered {
		_ = logoontheleft
		var setPointerField string
		_ = setPointerField

		// Insertion point for pointers initialization
	}

	for _, logoontheright := range logoontherightOrdered {
		_ = logoontheright
		var setPointerField string
		_ = setPointerField

		// Insertion point for pointers initialization
	}

	for _, markdown := range markdownOrdered {
		_ = markdown
		var setPointerField string
		_ = setPointerField

		// Insertion point for pointers initialization
	}

	for _, slider := range sliderOrdered {
		_ = slider
		var setPointerField string
		_ = setPointerField

		// Insertion point for pointers initialization
	}

	for _, split := range splitOrdered {
		_ = split
		var setPointerField string
		_ = setPointerField

		// Insertion point for pointers initialization
	}

	for _, svg := range svgOrdered {
		_ = svg
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

	for _, title := range titleOrdered {
		_ = title
		var setPointerField string
		_ = setPointerField

		// Insertion point for pointers initialization
	}

	for _, tone := range toneOrdered {
		_ = tone
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

	for _, view := range viewOrdered {
		_ = view
		var setPointerField string
		_ = setPointerField

		// Insertion point for pointers initialization
	}

	for _, xlsx := range xlsxOrdered {
		_ = xlsx
		var setPointerField string
		_ = setPointerField

		// Insertion point for pointers initialization
	}

	res = strings.ReplaceAll(res, "{{Identifiers}}", identifiersDecl.String())
	res = strings.ReplaceAll(res, "{{ValueInitializers}}", initializerStatements.String())
	res = strings.ReplaceAll(res, "{{PointersInitializers}}", pointersInitializesStatements.String())

	if stage.MetaPackageImportAlias != "" {
		res = strings.ReplaceAll(res, "{{ImportPackageDeclaration}}",
			fmt.Sprintf("\n\t%s %s", stage.MetaPackageImportAlias, stage.MetaPackageImportPath))

		res = strings.ReplaceAll(res, "{{ImportPackageDummyDeclaration}}",
			fmt.Sprintf("\nvar _ %s.Stage",
				stage.MetaPackageImportAlias))

		var entries strings.Builder

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
				entries.WriteString(fmt.Sprintf("\n\n\t\"%s\": %s(0),", value.Ident, value.Ident))
			case GONG__ENUM_CAST_STRING:
				entries.WriteString(fmt.Sprintf("\n\n\t\"%s\": %s(\"\"),", value.Ident, value.Ident))
			case GONG__FIELD_VALUE:
				// substitute the second point with "{})."
				joker := "__substitute_for_first_point__"
				valueIdentifier := strings.Replace(value.Ident, ".", joker, 1)
				valueIdentifier = strings.Replace(valueIdentifier, ".", "{}).", 1)
				valueIdentifier = strings.Replace(valueIdentifier, joker, ".", 1)
				entries.WriteString(fmt.Sprintf("\n\n\t\"%s\": (%s,", value.Ident, valueIdentifier))
			case GONG__IDENTIFIER_CONST:
				entries.WriteString(fmt.Sprintf("\n\n\t\"%s\": %s,", value.Ident, value.Ident))
			case GONG__STRUCT_INSTANCE:
				entries.WriteString(fmt.Sprintf("\n\n\t\"%s\": &(%s{}),", value.Ident, value.Ident))
			}
		}

		// res = strings.ReplaceAll(res, "{{EntriesDocLinkStringDocLinkIdentifier}}", entries.String())
	}
	return
}

// insertion point for marshall field methods
func (assplit *AsSplit) GongMarshallField(stage *Stage, fieldName string) (res string) {

	switch fieldName {
	case "Name":
		res = StringInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", assplit.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "Name")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", string(assplit.Name))
	case "Direction":
		if assplit.Direction.ToCodeString() != "" {
			res = StringEnumInitStatement
			res = strings.ReplaceAll(res, "{{Identifier}}", assplit.GongGetIdentifier(stage))
			res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "Direction")
			res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", "models."+assplit.Direction.ToCodeString())
		} else {
			// in case of empty enum, we need to unstage the previous value
			res = StringEnumInitStatement
			res = strings.ReplaceAll(res, "{{Identifier}}", assplit.GongGetIdentifier(stage))
			res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "Direction")
			res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", "\"\"")
		}

	case "AsSplitAreas":
		var sb strings.Builder
		for _, _assplitarea := range assplit.AsSplitAreas {
			tmp := SliceOfPointersFieldInitStatement
			tmp = strings.ReplaceAll(tmp, "{{Identifier}}", assplit.GongGetIdentifier(stage))
			tmp = strings.ReplaceAll(tmp, "{{GeneratedFieldName}}", "AsSplitAreas")
			tmp = strings.ReplaceAll(tmp, "{{GeneratedFieldNameValue}}", _assplitarea.GongGetIdentifier(stage))
			sb.WriteString(tmp)
		}
		res = sb.String()
	default:
		log.Panicf("Unknown field %s for Gongstruct AsSplit", fieldName)
	}
	return
}

func (assplitarea *AsSplitArea) GongMarshallField(stage *Stage, fieldName string) (res string) {

	switch fieldName {
	case "Name":
		res = StringInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", assplitarea.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "Name")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", string(assplitarea.Name))
	case "ShowNameInHeader":
		res = NumberInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", assplitarea.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "ShowNameInHeader")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%t", assplitarea.ShowNameInHeader))
	case "Size":
		res = NumberInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", assplitarea.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "Size")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%f", assplitarea.Size))
	case "IsAny":
		res = NumberInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", assplitarea.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "IsAny")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%t", assplitarea.IsAny))
	case "HasDiv":
		res = NumberInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", assplitarea.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "HasDiv")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%t", assplitarea.HasDiv))
	case "DivStyle":
		res = StringInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", assplitarea.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "DivStyle")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", string(assplitarea.DivStyle))

	case "AsSplit":
		if assplitarea.AsSplit != nil {
			res = PointerFieldInitStatement
			res = strings.ReplaceAll(res, "{{Identifier}}", assplitarea.GongGetIdentifier(stage))
			res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "AsSplit")
			res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", assplitarea.AsSplit.GongGetIdentifier(stage))
		} else {
			// in case of nil pointer, we need to unstage the previous value
			res = PointerFieldInitStatement
			res = strings.ReplaceAll(res, "{{Identifier}}", assplitarea.GongGetIdentifier(stage))
			res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "AsSplit")
			res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", "nil")
		}
	case "Button":
		if assplitarea.Button != nil {
			res = PointerFieldInitStatement
			res = strings.ReplaceAll(res, "{{Identifier}}", assplitarea.GongGetIdentifier(stage))
			res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "Button")
			res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", assplitarea.Button.GongGetIdentifier(stage))
		} else {
			// in case of nil pointer, we need to unstage the previous value
			res = PointerFieldInitStatement
			res = strings.ReplaceAll(res, "{{Identifier}}", assplitarea.GongGetIdentifier(stage))
			res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "Button")
			res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", "nil")
		}
	case "Cursor":
		if assplitarea.Cursor != nil {
			res = PointerFieldInitStatement
			res = strings.ReplaceAll(res, "{{Identifier}}", assplitarea.GongGetIdentifier(stage))
			res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "Cursor")
			res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", assplitarea.Cursor.GongGetIdentifier(stage))
		} else {
			// in case of nil pointer, we need to unstage the previous value
			res = PointerFieldInitStatement
			res = strings.ReplaceAll(res, "{{Identifier}}", assplitarea.GongGetIdentifier(stage))
			res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "Cursor")
			res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", "nil")
		}
	case "Form":
		if assplitarea.Form != nil {
			res = PointerFieldInitStatement
			res = strings.ReplaceAll(res, "{{Identifier}}", assplitarea.GongGetIdentifier(stage))
			res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "Form")
			res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", assplitarea.Form.GongGetIdentifier(stage))
		} else {
			// in case of nil pointer, we need to unstage the previous value
			res = PointerFieldInitStatement
			res = strings.ReplaceAll(res, "{{Identifier}}", assplitarea.GongGetIdentifier(stage))
			res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "Form")
			res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", "nil")
		}
	case "Load":
		if assplitarea.Load != nil {
			res = PointerFieldInitStatement
			res = strings.ReplaceAll(res, "{{Identifier}}", assplitarea.GongGetIdentifier(stage))
			res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "Load")
			res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", assplitarea.Load.GongGetIdentifier(stage))
		} else {
			// in case of nil pointer, we need to unstage the previous value
			res = PointerFieldInitStatement
			res = strings.ReplaceAll(res, "{{Identifier}}", assplitarea.GongGetIdentifier(stage))
			res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "Load")
			res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", "nil")
		}
	case "Markdown":
		if assplitarea.Markdown != nil {
			res = PointerFieldInitStatement
			res = strings.ReplaceAll(res, "{{Identifier}}", assplitarea.GongGetIdentifier(stage))
			res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "Markdown")
			res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", assplitarea.Markdown.GongGetIdentifier(stage))
		} else {
			// in case of nil pointer, we need to unstage the previous value
			res = PointerFieldInitStatement
			res = strings.ReplaceAll(res, "{{Identifier}}", assplitarea.GongGetIdentifier(stage))
			res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "Markdown")
			res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", "nil")
		}
	case "Slider":
		if assplitarea.Slider != nil {
			res = PointerFieldInitStatement
			res = strings.ReplaceAll(res, "{{Identifier}}", assplitarea.GongGetIdentifier(stage))
			res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "Slider")
			res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", assplitarea.Slider.GongGetIdentifier(stage))
		} else {
			// in case of nil pointer, we need to unstage the previous value
			res = PointerFieldInitStatement
			res = strings.ReplaceAll(res, "{{Identifier}}", assplitarea.GongGetIdentifier(stage))
			res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "Slider")
			res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", "nil")
		}
	case "Split":
		if assplitarea.Split != nil {
			res = PointerFieldInitStatement
			res = strings.ReplaceAll(res, "{{Identifier}}", assplitarea.GongGetIdentifier(stage))
			res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "Split")
			res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", assplitarea.Split.GongGetIdentifier(stage))
		} else {
			// in case of nil pointer, we need to unstage the previous value
			res = PointerFieldInitStatement
			res = strings.ReplaceAll(res, "{{Identifier}}", assplitarea.GongGetIdentifier(stage))
			res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "Split")
			res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", "nil")
		}
	case "Svg":
		if assplitarea.Svg != nil {
			res = PointerFieldInitStatement
			res = strings.ReplaceAll(res, "{{Identifier}}", assplitarea.GongGetIdentifier(stage))
			res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "Svg")
			res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", assplitarea.Svg.GongGetIdentifier(stage))
		} else {
			// in case of nil pointer, we need to unstage the previous value
			res = PointerFieldInitStatement
			res = strings.ReplaceAll(res, "{{Identifier}}", assplitarea.GongGetIdentifier(stage))
			res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "Svg")
			res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", "nil")
		}
	case "Table":
		if assplitarea.Table != nil {
			res = PointerFieldInitStatement
			res = strings.ReplaceAll(res, "{{Identifier}}", assplitarea.GongGetIdentifier(stage))
			res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "Table")
			res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", assplitarea.Table.GongGetIdentifier(stage))
		} else {
			// in case of nil pointer, we need to unstage the previous value
			res = PointerFieldInitStatement
			res = strings.ReplaceAll(res, "{{Identifier}}", assplitarea.GongGetIdentifier(stage))
			res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "Table")
			res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", "nil")
		}
	case "Tone":
		if assplitarea.Tone != nil {
			res = PointerFieldInitStatement
			res = strings.ReplaceAll(res, "{{Identifier}}", assplitarea.GongGetIdentifier(stage))
			res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "Tone")
			res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", assplitarea.Tone.GongGetIdentifier(stage))
		} else {
			// in case of nil pointer, we need to unstage the previous value
			res = PointerFieldInitStatement
			res = strings.ReplaceAll(res, "{{Identifier}}", assplitarea.GongGetIdentifier(stage))
			res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "Tone")
			res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", "nil")
		}
	case "Tree":
		if assplitarea.Tree != nil {
			res = PointerFieldInitStatement
			res = strings.ReplaceAll(res, "{{Identifier}}", assplitarea.GongGetIdentifier(stage))
			res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "Tree")
			res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", assplitarea.Tree.GongGetIdentifier(stage))
		} else {
			// in case of nil pointer, we need to unstage the previous value
			res = PointerFieldInitStatement
			res = strings.ReplaceAll(res, "{{Identifier}}", assplitarea.GongGetIdentifier(stage))
			res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "Tree")
			res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", "nil")
		}
	case "Xlsx":
		if assplitarea.Xlsx != nil {
			res = PointerFieldInitStatement
			res = strings.ReplaceAll(res, "{{Identifier}}", assplitarea.GongGetIdentifier(stage))
			res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "Xlsx")
			res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", assplitarea.Xlsx.GongGetIdentifier(stage))
		} else {
			// in case of nil pointer, we need to unstage the previous value
			res = PointerFieldInitStatement
			res = strings.ReplaceAll(res, "{{Identifier}}", assplitarea.GongGetIdentifier(stage))
			res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "Xlsx")
			res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", "nil")
		}
	default:
		log.Panicf("Unknown field %s for Gongstruct AsSplitArea", fieldName)
	}
	return
}

func (button *Button) GongMarshallField(stage *Stage, fieldName string) (res string) {

	switch fieldName {
	case "Name":
		res = StringInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", button.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "Name")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", string(button.Name))
	case "StackName":
		res = StringInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", button.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "StackName")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", string(button.StackName))

	default:
		log.Panicf("Unknown field %s for Gongstruct Button", fieldName)
	}
	return
}

func (cursor *Cursor) GongMarshallField(stage *Stage, fieldName string) (res string) {

	switch fieldName {
	case "Name":
		res = StringInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", cursor.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "Name")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", string(cursor.Name))
	case "StackName":
		res = StringInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", cursor.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "StackName")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", string(cursor.StackName))
	case "Style":
		res = StringInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", cursor.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "Style")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", string(cursor.Style))

	default:
		log.Panicf("Unknown field %s for Gongstruct Cursor", fieldName)
	}
	return
}

func (favicon *FavIcon) GongMarshallField(stage *Stage, fieldName string) (res string) {

	switch fieldName {
	case "Name":
		res = StringInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", favicon.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "Name")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", string(favicon.Name))
	case "SVG":
		res = StringInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", favicon.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "SVG")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", string(favicon.SVG))

	default:
		log.Panicf("Unknown field %s for Gongstruct FavIcon", fieldName)
	}
	return
}

func (form *Form) GongMarshallField(stage *Stage, fieldName string) (res string) {

	switch fieldName {
	case "Name":
		res = StringInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", form.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "Name")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", string(form.Name))
	case "StackName":
		res = StringInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", form.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "StackName")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", string(form.StackName))

	default:
		log.Panicf("Unknown field %s for Gongstruct Form", fieldName)
	}
	return
}

func (load *Load) GongMarshallField(stage *Stage, fieldName string) (res string) {

	switch fieldName {
	case "Name":
		res = StringInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", load.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "Name")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", string(load.Name))
	case "StackName":
		res = StringInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", load.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "StackName")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", string(load.StackName))

	default:
		log.Panicf("Unknown field %s for Gongstruct Load", fieldName)
	}
	return
}

func (logoontheleft *LogoOnTheLeft) GongMarshallField(stage *Stage, fieldName string) (res string) {

	switch fieldName {
	case "Name":
		res = StringInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", logoontheleft.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "Name")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", string(logoontheleft.Name))
	case "Width":
		res = NumberInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", logoontheleft.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "Width")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%d", logoontheleft.Width))
	case "Height":
		res = NumberInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", logoontheleft.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "Height")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%d", logoontheleft.Height))
	case "SVG":
		res = StringInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", logoontheleft.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "SVG")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", string(logoontheleft.SVG))

	default:
		log.Panicf("Unknown field %s for Gongstruct LogoOnTheLeft", fieldName)
	}
	return
}

func (logoontheright *LogoOnTheRight) GongMarshallField(stage *Stage, fieldName string) (res string) {

	switch fieldName {
	case "Name":
		res = StringInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", logoontheright.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "Name")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", string(logoontheright.Name))
	case "Width":
		res = NumberInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", logoontheright.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "Width")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%d", logoontheright.Width))
	case "Height":
		res = NumberInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", logoontheright.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "Height")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%d", logoontheright.Height))
	case "SVG":
		res = StringInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", logoontheright.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "SVG")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", string(logoontheright.SVG))

	default:
		log.Panicf("Unknown field %s for Gongstruct LogoOnTheRight", fieldName)
	}
	return
}

func (markdown *Markdown) GongMarshallField(stage *Stage, fieldName string) (res string) {

	switch fieldName {
	case "Name":
		res = StringInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", markdown.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "Name")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", string(markdown.Name))
	case "StackName":
		res = StringInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", markdown.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "StackName")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", string(markdown.StackName))

	default:
		log.Panicf("Unknown field %s for Gongstruct Markdown", fieldName)
	}
	return
}

func (slider *Slider) GongMarshallField(stage *Stage, fieldName string) (res string) {

	switch fieldName {
	case "Name":
		res = StringInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", slider.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "Name")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", string(slider.Name))
	case "StackName":
		res = StringInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", slider.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "StackName")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", string(slider.StackName))

	default:
		log.Panicf("Unknown field %s for Gongstruct Slider", fieldName)
	}
	return
}

func (split *Split) GongMarshallField(stage *Stage, fieldName string) (res string) {

	switch fieldName {
	case "Name":
		res = StringInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", split.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "Name")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", string(split.Name))
	case "StackName":
		res = StringInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", split.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "StackName")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", string(split.StackName))

	default:
		log.Panicf("Unknown field %s for Gongstruct Split", fieldName)
	}
	return
}

func (svg *Svg) GongMarshallField(stage *Stage, fieldName string) (res string) {

	switch fieldName {
	case "Name":
		res = StringInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", svg.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "Name")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", string(svg.Name))
	case "StackName":
		res = StringInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", svg.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "StackName")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", string(svg.StackName))
	case "Style":
		res = StringInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", svg.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "Style")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", string(svg.Style))

	default:
		log.Panicf("Unknown field %s for Gongstruct Svg", fieldName)
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
	case "StackName":
		res = StringInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", table.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "StackName")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", string(table.StackName))

	default:
		log.Panicf("Unknown field %s for Gongstruct Table", fieldName)
	}
	return
}

func (title *Title) GongMarshallField(stage *Stage, fieldName string) (res string) {

	switch fieldName {
	case "Name":
		res = StringInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", title.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "Name")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", string(title.Name))

	default:
		log.Panicf("Unknown field %s for Gongstruct Title", fieldName)
	}
	return
}

func (tone *Tone) GongMarshallField(stage *Stage, fieldName string) (res string) {

	switch fieldName {
	case "Name":
		res = StringInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", tone.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "Name")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", string(tone.Name))
	case "StackName":
		res = StringInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", tone.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "StackName")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", string(tone.StackName))

	default:
		log.Panicf("Unknown field %s for Gongstruct Tone", fieldName)
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
	case "StackName":
		res = StringInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", tree.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "StackName")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", string(tree.StackName))

	default:
		log.Panicf("Unknown field %s for Gongstruct Tree", fieldName)
	}
	return
}

func (view *View) GongMarshallField(stage *Stage, fieldName string) (res string) {

	switch fieldName {
	case "Name":
		res = StringInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", view.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "Name")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", string(view.Name))
	case "ShowViewName":
		res = NumberInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", view.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "ShowViewName")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%t", view.ShowViewName))
	case "IsSelectedView":
		res = NumberInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", view.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "IsSelectedView")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%t", view.IsSelectedView))
	case "Direction":
		if view.Direction.ToCodeString() != "" {
			res = StringEnumInitStatement
			res = strings.ReplaceAll(res, "{{Identifier}}", view.GongGetIdentifier(stage))
			res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "Direction")
			res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", "models."+view.Direction.ToCodeString())
		} else {
			// in case of empty enum, we need to unstage the previous value
			res = StringEnumInitStatement
			res = strings.ReplaceAll(res, "{{Identifier}}", view.GongGetIdentifier(stage))
			res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "Direction")
			res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", "\"\"")
		}

	case "RootAsSplitAreas":
		var sb strings.Builder
		for _, _assplitarea := range view.RootAsSplitAreas {
			tmp := SliceOfPointersFieldInitStatement
			tmp = strings.ReplaceAll(tmp, "{{Identifier}}", view.GongGetIdentifier(stage))
			tmp = strings.ReplaceAll(tmp, "{{GeneratedFieldName}}", "RootAsSplitAreas")
			tmp = strings.ReplaceAll(tmp, "{{GeneratedFieldNameValue}}", _assplitarea.GongGetIdentifier(stage))
			sb.WriteString(tmp)
		}
		res = sb.String()
	default:
		log.Panicf("Unknown field %s for Gongstruct View", fieldName)
	}
	return
}

func (xlsx *Xlsx) GongMarshallField(stage *Stage, fieldName string) (res string) {

	switch fieldName {
	case "Name":
		res = StringInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", xlsx.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "Name")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", string(xlsx.Name))
	case "StackName":
		res = StringInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", xlsx.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "StackName")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", string(xlsx.StackName))

	default:
		log.Panicf("Unknown field %s for Gongstruct Xlsx", fieldName)
	}
	return
}

// insertion point for marshall all fields methods
func (assplit *AsSplit) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(assplit.GongMarshallField(stage, "Name"))
		initializerStatements.WriteString(assplit.GongMarshallField(stage, "Direction"))
		pointersInitializesStatements.WriteString(assplit.GongMarshallField(stage, "AsSplitAreas"))
	}
	initRes = initializerStatements.String()
	ptrRes = pointersInitializesStatements.String()
	return
}
func (assplitarea *AsSplitArea) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(assplitarea.GongMarshallField(stage, "Name"))
		initializerStatements.WriteString(assplitarea.GongMarshallField(stage, "ShowNameInHeader"))
		initializerStatements.WriteString(assplitarea.GongMarshallField(stage, "Size"))
		initializerStatements.WriteString(assplitarea.GongMarshallField(stage, "IsAny"))
		pointersInitializesStatements.WriteString(assplitarea.GongMarshallField(stage, "AsSplit"))
		pointersInitializesStatements.WriteString(assplitarea.GongMarshallField(stage, "Button"))
		pointersInitializesStatements.WriteString(assplitarea.GongMarshallField(stage, "Cursor"))
		pointersInitializesStatements.WriteString(assplitarea.GongMarshallField(stage, "Form"))
		pointersInitializesStatements.WriteString(assplitarea.GongMarshallField(stage, "Load"))
		pointersInitializesStatements.WriteString(assplitarea.GongMarshallField(stage, "Markdown"))
		pointersInitializesStatements.WriteString(assplitarea.GongMarshallField(stage, "Slider"))
		pointersInitializesStatements.WriteString(assplitarea.GongMarshallField(stage, "Split"))
		pointersInitializesStatements.WriteString(assplitarea.GongMarshallField(stage, "Svg"))
		pointersInitializesStatements.WriteString(assplitarea.GongMarshallField(stage, "Table"))
		pointersInitializesStatements.WriteString(assplitarea.GongMarshallField(stage, "Tone"))
		pointersInitializesStatements.WriteString(assplitarea.GongMarshallField(stage, "Tree"))
		pointersInitializesStatements.WriteString(assplitarea.GongMarshallField(stage, "Xlsx"))
		initializerStatements.WriteString(assplitarea.GongMarshallField(stage, "HasDiv"))
		initializerStatements.WriteString(assplitarea.GongMarshallField(stage, "DivStyle"))
	}
	initRes = initializerStatements.String()
	ptrRes = pointersInitializesStatements.String()
	return
}
func (button *Button) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(button.GongMarshallField(stage, "Name"))
		initializerStatements.WriteString(button.GongMarshallField(stage, "StackName"))
	}
	initRes = initializerStatements.String()
	ptrRes = pointersInitializesStatements.String()
	return
}
func (cursor *Cursor) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(cursor.GongMarshallField(stage, "Name"))
		initializerStatements.WriteString(cursor.GongMarshallField(stage, "StackName"))
		initializerStatements.WriteString(cursor.GongMarshallField(stage, "Style"))
	}
	initRes = initializerStatements.String()
	ptrRes = pointersInitializesStatements.String()
	return
}
func (favicon *FavIcon) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(favicon.GongMarshallField(stage, "Name"))
		initializerStatements.WriteString(favicon.GongMarshallField(stage, "SVG"))
	}
	initRes = initializerStatements.String()
	ptrRes = pointersInitializesStatements.String()
	return
}
func (form *Form) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(form.GongMarshallField(stage, "Name"))
		initializerStatements.WriteString(form.GongMarshallField(stage, "StackName"))
	}
	initRes = initializerStatements.String()
	ptrRes = pointersInitializesStatements.String()
	return
}
func (load *Load) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(load.GongMarshallField(stage, "Name"))
		initializerStatements.WriteString(load.GongMarshallField(stage, "StackName"))
	}
	initRes = initializerStatements.String()
	ptrRes = pointersInitializesStatements.String()
	return
}
func (logoontheleft *LogoOnTheLeft) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(logoontheleft.GongMarshallField(stage, "Name"))
		initializerStatements.WriteString(logoontheleft.GongMarshallField(stage, "Width"))
		initializerStatements.WriteString(logoontheleft.GongMarshallField(stage, "Height"))
		initializerStatements.WriteString(logoontheleft.GongMarshallField(stage, "SVG"))
	}
	initRes = initializerStatements.String()
	ptrRes = pointersInitializesStatements.String()
	return
}
func (logoontheright *LogoOnTheRight) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(logoontheright.GongMarshallField(stage, "Name"))
		initializerStatements.WriteString(logoontheright.GongMarshallField(stage, "Width"))
		initializerStatements.WriteString(logoontheright.GongMarshallField(stage, "Height"))
		initializerStatements.WriteString(logoontheright.GongMarshallField(stage, "SVG"))
	}
	initRes = initializerStatements.String()
	ptrRes = pointersInitializesStatements.String()
	return
}
func (markdown *Markdown) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(markdown.GongMarshallField(stage, "Name"))
		initializerStatements.WriteString(markdown.GongMarshallField(stage, "StackName"))
	}
	initRes = initializerStatements.String()
	ptrRes = pointersInitializesStatements.String()
	return
}
func (slider *Slider) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(slider.GongMarshallField(stage, "Name"))
		initializerStatements.WriteString(slider.GongMarshallField(stage, "StackName"))
	}
	initRes = initializerStatements.String()
	ptrRes = pointersInitializesStatements.String()
	return
}
func (split *Split) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(split.GongMarshallField(stage, "Name"))
		initializerStatements.WriteString(split.GongMarshallField(stage, "StackName"))
	}
	initRes = initializerStatements.String()
	ptrRes = pointersInitializesStatements.String()
	return
}
func (svg *Svg) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(svg.GongMarshallField(stage, "Name"))
		initializerStatements.WriteString(svg.GongMarshallField(stage, "StackName"))
		initializerStatements.WriteString(svg.GongMarshallField(stage, "Style"))
	}
	initRes = initializerStatements.String()
	ptrRes = pointersInitializesStatements.String()
	return
}
func (table *Table) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(table.GongMarshallField(stage, "Name"))
		initializerStatements.WriteString(table.GongMarshallField(stage, "StackName"))
	}
	initRes = initializerStatements.String()
	ptrRes = pointersInitializesStatements.String()
	return
}
func (title *Title) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(title.GongMarshallField(stage, "Name"))
	}
	initRes = initializerStatements.String()
	ptrRes = pointersInitializesStatements.String()
	return
}
func (tone *Tone) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(tone.GongMarshallField(stage, "Name"))
		initializerStatements.WriteString(tone.GongMarshallField(stage, "StackName"))
	}
	initRes = initializerStatements.String()
	ptrRes = pointersInitializesStatements.String()
	return
}
func (tree *Tree) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(tree.GongMarshallField(stage, "Name"))
		initializerStatements.WriteString(tree.GongMarshallField(stage, "StackName"))
	}
	initRes = initializerStatements.String()
	ptrRes = pointersInitializesStatements.String()
	return
}
func (view *View) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(view.GongMarshallField(stage, "Name"))
		initializerStatements.WriteString(view.GongMarshallField(stage, "ShowViewName"))
		pointersInitializesStatements.WriteString(view.GongMarshallField(stage, "RootAsSplitAreas"))
		initializerStatements.WriteString(view.GongMarshallField(stage, "IsSelectedView"))
		initializerStatements.WriteString(view.GongMarshallField(stage, "Direction"))
	}
	initRes = initializerStatements.String()
	ptrRes = pointersInitializesStatements.String()
	return
}
func (xlsx *Xlsx) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(xlsx.GongMarshallField(stage, "Name"))
		initializerStatements.WriteString(xlsx.GongMarshallField(stage, "StackName"))
	}
	initRes = initializerStatements.String()
	ptrRes = pointersInitializesStatements.String()
	return
}
