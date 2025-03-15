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
func _(stage *models.StageStruct) {

	// Declaration of instances to stage{{Identifiers}}

	// Setup of values{{ValueInitializers}}

	// Setup of pointers{{PointersInitializers}}
}`

const IdentifiersDecls = `
	{{Identifier}} := (&models.{{GeneratedStructName}}{}).Stage(stage)`

const StringInitStatement = `
	{{Identifier}}.{{GeneratedFieldName}} = ` + "`" + `{{GeneratedFieldNameValue}}` + "`"

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
func (stage *StageStruct) Marshall(file *os.File, modelsPackageName, packageName string) {

	name := file.Name()

	if !strings.HasSuffix(name, ".go") {
		log.Fatalln(name + " is not a go filename")
	}

	log.Println("filename of marshall output is " + name)
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
		asspliti_order, oki := stage.Map_Staged_Order[asspliti]
		assplitj_order, okj := stage.Map_Staged_Order[assplitj]
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
		assplitareai_order, oki := stage.Map_Staged_Order[assplitareai]
		assplitareaj_order, okj := stage.Map_Staged_Order[assplitareaj]
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
		doci_order, oki := stage.Map_Staged_Order[doci]
		docj_order, okj := stage.Map_Staged_Order[docj]
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

	map_Form_Identifiers := make(map[*Form]string)
	_ = map_Form_Identifiers

	formOrdered := []*Form{}
	for form := range stage.Forms {
		formOrdered = append(formOrdered, form)
	}
	sort.Slice(formOrdered[:], func(i, j int) bool {
		formi := formOrdered[i]
		formj := formOrdered[j]
		formi_order, oki := stage.Map_Staged_Order[formi]
		formj_order, okj := stage.Map_Staged_Order[formj]
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

	map_Split_Identifiers := make(map[*Split]string)
	_ = map_Split_Identifiers

	splitOrdered := []*Split{}
	for split := range stage.Splits {
		splitOrdered = append(splitOrdered, split)
	}
	sort.Slice(splitOrdered[:], func(i, j int) bool {
		spliti := splitOrdered[i]
		splitj := splitOrdered[j]
		spliti_order, oki := stage.Map_Staged_Order[spliti]
		splitj_order, okj := stage.Map_Staged_Order[splitj]
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
		svgi_order, oki := stage.Map_Staged_Order[svgi]
		svgj_order, okj := stage.Map_Staged_Order[svgj]
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
		tablei_order, oki := stage.Map_Staged_Order[tablei]
		tablej_order, okj := stage.Map_Staged_Order[tablej]
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

	map_Tree_Identifiers := make(map[*Tree]string)
	_ = map_Tree_Identifiers

	treeOrdered := []*Tree{}
	for tree := range stage.Trees {
		treeOrdered = append(treeOrdered, tree)
	}
	sort.Slice(treeOrdered[:], func(i, j int) bool {
		treei := treeOrdered[i]
		treej := treeOrdered[j]
		treei_order, oki := stage.Map_Staged_Order[treei]
		treej_order, okj := stage.Map_Staged_Order[treej]
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
		viewi_order, oki := stage.Map_Staged_Order[viewi]
		viewj_order, okj := stage.Map_Staged_Order[viewj]
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
		for _, _assplit := range assplitarea.AsSplits {
			setPointerField = SliceOfPointersFieldInitStatement
			setPointerField = strings.ReplaceAll(setPointerField, "{{Identifier}}", id)
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldName}}", "AsSplits")
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldNameValue}}", map_AsSplit_Identifiers[_assplit])
			pointersInitializesStatements += setPointerField
		}

		if assplitarea.Tree != nil {
			setPointerField = PointerFieldInitStatement
			setPointerField = strings.ReplaceAll(setPointerField, "{{Identifier}}", id)
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldName}}", "Tree")
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldNameValue}}", map_Tree_Identifiers[assplitarea.Tree])
			pointersInitializesStatements += setPointerField
		}

		if assplitarea.Table != nil {
			setPointerField = PointerFieldInitStatement
			setPointerField = strings.ReplaceAll(setPointerField, "{{Identifier}}", id)
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldName}}", "Table")
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldNameValue}}", map_Table_Identifiers[assplitarea.Table])
			pointersInitializesStatements += setPointerField
		}

		if assplitarea.Form != nil {
			setPointerField = PointerFieldInitStatement
			setPointerField = strings.ReplaceAll(setPointerField, "{{Identifier}}", id)
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldName}}", "Form")
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldNameValue}}", map_Form_Identifiers[assplitarea.Form])
			pointersInitializesStatements += setPointerField
		}

		if assplitarea.Svg != nil {
			setPointerField = PointerFieldInitStatement
			setPointerField = strings.ReplaceAll(setPointerField, "{{Identifier}}", id)
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldName}}", "Svg")
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldNameValue}}", map_Svg_Identifiers[assplitarea.Svg])
			pointersInitializesStatements += setPointerField
		}

		if assplitarea.Doc != nil {
			setPointerField = PointerFieldInitStatement
			setPointerField = strings.ReplaceAll(setPointerField, "{{Identifier}}", id)
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldName}}", "Doc")
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldNameValue}}", map_Doc_Identifiers[assplitarea.Doc])
			pointersInitializesStatements += setPointerField
		}

		if assplitarea.Split != nil {
			setPointerField = PointerFieldInitStatement
			setPointerField = strings.ReplaceAll(setPointerField, "{{Identifier}}", id)
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldName}}", "Split")
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldNameValue}}", map_Split_Identifiers[assplitarea.Split])
			pointersInitializesStatements += setPointerField
		}

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

	res = strings.ReplaceAll(res, "{{Identifiers}}", identifiersDecl)
	res = strings.ReplaceAll(res, "{{ValueInitializers}}", initializerStatements)
	res = strings.ReplaceAll(res, "{{PointersInitializers}}", pointersInitializesStatements)

	if stage.MetaPackageImportAlias != "" {
		res = strings.ReplaceAll(res, "{{ImportPackageDeclaration}}",
			fmt.Sprintf("\n\t%s %s", stage.MetaPackageImportAlias, stage.MetaPackageImportPath))

		res = strings.ReplaceAll(res, "{{ImportPackageDummyDeclaration}}",
			fmt.Sprintf("\nvar _ %s.StageStruct",
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

	fmt.Fprintln(file, res)
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
