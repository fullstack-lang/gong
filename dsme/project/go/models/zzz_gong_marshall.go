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
	map_CompositionShape_Identifiers := make(map[*CompositionShape]string)
	_ = map_CompositionShape_Identifiers

	compositionshapeOrdered := []*CompositionShape{}
	for compositionshape := range stage.CompositionShapes {
		compositionshapeOrdered = append(compositionshapeOrdered, compositionshape)
	}
	sort.Slice(compositionshapeOrdered[:], func(i, j int) bool {
		compositionshapei := compositionshapeOrdered[i]
		compositionshapej := compositionshapeOrdered[j]
		compositionshapei_order, oki := stage.CompositionShapeMap_Staged_Order[compositionshapei]
		compositionshapej_order, okj := stage.CompositionShapeMap_Staged_Order[compositionshapej]
		if !oki || !okj {
			log.Fatalln("unknown pointers")
		}
		return compositionshapei_order < compositionshapej_order
	})
	if len(compositionshapeOrdered) > 0 {
		identifiersDecl += "\n"
	}
	for _, compositionshape := range compositionshapeOrdered {

		id = generatesIdentifier("CompositionShape", int(stage.CompositionShapeMap_Staged_Order[compositionshape]), compositionshape.Name)
		map_CompositionShape_Identifiers[compositionshape] = id

		decl = IdentifiersDecls
		decl = strings.ReplaceAll(decl, "{{Identifier}}", id)
		decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "CompositionShape")
		decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", compositionshape.Name)
		identifiersDecl += decl

		initializerStatements += "\n"
		// Initialisation of values
		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Name")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(compositionshape.Name))
		initializerStatements += setValueField

		setValueField = NumberInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "StartRatio")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%f", compositionshape.StartRatio))
		initializerStatements += setValueField

		setValueField = NumberInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "EndRatio")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%f", compositionshape.EndRatio))
		initializerStatements += setValueField

		if compositionshape.StartOrientation != "" {
			setValueField = StringEnumInitStatement
			setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
			setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "StartOrientation")
			setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", "models."+compositionshape.StartOrientation.ToCodeString())
			initializerStatements += setValueField
		}

		if compositionshape.EndOrientation != "" {
			setValueField = StringEnumInitStatement
			setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
			setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "EndOrientation")
			setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", "models."+compositionshape.EndOrientation.ToCodeString())
			initializerStatements += setValueField
		}

		setValueField = NumberInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "CornerOffsetRatio")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%f", compositionshape.CornerOffsetRatio))
		initializerStatements += setValueField

	}

	map_Diagram_Identifiers := make(map[*Diagram]string)
	_ = map_Diagram_Identifiers

	diagramOrdered := []*Diagram{}
	for diagram := range stage.Diagrams {
		diagramOrdered = append(diagramOrdered, diagram)
	}
	sort.Slice(diagramOrdered[:], func(i, j int) bool {
		diagrami := diagramOrdered[i]
		diagramj := diagramOrdered[j]
		diagrami_order, oki := stage.DiagramMap_Staged_Order[diagrami]
		diagramj_order, okj := stage.DiagramMap_Staged_Order[diagramj]
		if !oki || !okj {
			log.Fatalln("unknown pointers")
		}
		return diagrami_order < diagramj_order
	})
	if len(diagramOrdered) > 0 {
		identifiersDecl += "\n"
	}
	for _, diagram := range diagramOrdered {

		id = generatesIdentifier("Diagram", int(stage.DiagramMap_Staged_Order[diagram]), diagram.Name)
		map_Diagram_Identifiers[diagram] = id

		decl = IdentifiersDecls
		decl = strings.ReplaceAll(decl, "{{Identifier}}", id)
		decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "Diagram")
		decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", diagram.Name)
		identifiersDecl += decl

		initializerStatements += "\n"
		// Initialisation of values
		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Name")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(diagram.Name))
		initializerStatements += setValueField

		setValueField = NumberInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "IsChecked")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%t", diagram.IsChecked))
		initializerStatements += setValueField

		setValueField = NumberInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "IsEditable_")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%t", diagram.IsEditable_))
		initializerStatements += setValueField

		setValueField = NumberInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "IsInRenameMode")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%t", diagram.IsInRenameMode))
		initializerStatements += setValueField

		setValueField = NumberInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "IsExpanded")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%t", diagram.IsExpanded))
		initializerStatements += setValueField

		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "ComputedPrefix")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(diagram.ComputedPrefix))
		initializerStatements += setValueField

		setValueField = NumberInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "IsPBSNodeExpanded")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%t", diagram.IsPBSNodeExpanded))
		initializerStatements += setValueField

		setValueField = NumberInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "IsWBSNodeExpanded")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%t", diagram.IsWBSNodeExpanded))
		initializerStatements += setValueField

	}

	map_Product_Identifiers := make(map[*Product]string)
	_ = map_Product_Identifiers

	productOrdered := []*Product{}
	for product := range stage.Products {
		productOrdered = append(productOrdered, product)
	}
	sort.Slice(productOrdered[:], func(i, j int) bool {
		producti := productOrdered[i]
		productj := productOrdered[j]
		producti_order, oki := stage.ProductMap_Staged_Order[producti]
		productj_order, okj := stage.ProductMap_Staged_Order[productj]
		if !oki || !okj {
			log.Fatalln("unknown pointers")
		}
		return producti_order < productj_order
	})
	if len(productOrdered) > 0 {
		identifiersDecl += "\n"
	}
	for _, product := range productOrdered {

		id = generatesIdentifier("Product", int(stage.ProductMap_Staged_Order[product]), product.Name)
		map_Product_Identifiers[product] = id

		decl = IdentifiersDecls
		decl = strings.ReplaceAll(decl, "{{Identifier}}", id)
		decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "Product")
		decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", product.Name)
		identifiersDecl += decl

		initializerStatements += "\n"
		// Initialisation of values
		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Name")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(product.Name))
		initializerStatements += setValueField

		setValueField = NumberInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "IsExpanded")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%t", product.IsExpanded))
		initializerStatements += setValueField

		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "ComputedPrefix")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(product.ComputedPrefix))
		initializerStatements += setValueField

		setValueField = NumberInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "IsProducersNodeExpanded")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%t", product.IsProducersNodeExpanded))
		initializerStatements += setValueField

		setValueField = NumberInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "IsConsumersNodeExpanded")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%t", product.IsConsumersNodeExpanded))
		initializerStatements += setValueField

	}

	map_ProductShape_Identifiers := make(map[*ProductShape]string)
	_ = map_ProductShape_Identifiers

	productshapeOrdered := []*ProductShape{}
	for productshape := range stage.ProductShapes {
		productshapeOrdered = append(productshapeOrdered, productshape)
	}
	sort.Slice(productshapeOrdered[:], func(i, j int) bool {
		productshapei := productshapeOrdered[i]
		productshapej := productshapeOrdered[j]
		productshapei_order, oki := stage.ProductShapeMap_Staged_Order[productshapei]
		productshapej_order, okj := stage.ProductShapeMap_Staged_Order[productshapej]
		if !oki || !okj {
			log.Fatalln("unknown pointers")
		}
		return productshapei_order < productshapej_order
	})
	if len(productshapeOrdered) > 0 {
		identifiersDecl += "\n"
	}
	for _, productshape := range productshapeOrdered {

		id = generatesIdentifier("ProductShape", int(stage.ProductShapeMap_Staged_Order[productshape]), productshape.Name)
		map_ProductShape_Identifiers[productshape] = id

		decl = IdentifiersDecls
		decl = strings.ReplaceAll(decl, "{{Identifier}}", id)
		decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "ProductShape")
		decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", productshape.Name)
		identifiersDecl += decl

		initializerStatements += "\n"
		// Initialisation of values
		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Name")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(productshape.Name))
		initializerStatements += setValueField

		setValueField = NumberInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "IsExpanded")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%t", productshape.IsExpanded))
		initializerStatements += setValueField

		setValueField = NumberInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "X")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%f", productshape.X))
		initializerStatements += setValueField

		setValueField = NumberInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Y")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%f", productshape.Y))
		initializerStatements += setValueField

		setValueField = NumberInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Width")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%f", productshape.Width))
		initializerStatements += setValueField

		setValueField = NumberInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Height")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%f", productshape.Height))
		initializerStatements += setValueField

	}

	map_Project_Identifiers := make(map[*Project]string)
	_ = map_Project_Identifiers

	projectOrdered := []*Project{}
	for project := range stage.Projects {
		projectOrdered = append(projectOrdered, project)
	}
	sort.Slice(projectOrdered[:], func(i, j int) bool {
		projecti := projectOrdered[i]
		projectj := projectOrdered[j]
		projecti_order, oki := stage.ProjectMap_Staged_Order[projecti]
		projectj_order, okj := stage.ProjectMap_Staged_Order[projectj]
		if !oki || !okj {
			log.Fatalln("unknown pointers")
		}
		return projecti_order < projectj_order
	})
	if len(projectOrdered) > 0 {
		identifiersDecl += "\n"
	}
	for _, project := range projectOrdered {

		id = generatesIdentifier("Project", int(stage.ProjectMap_Staged_Order[project]), project.Name)
		map_Project_Identifiers[project] = id

		decl = IdentifiersDecls
		decl = strings.ReplaceAll(decl, "{{Identifier}}", id)
		decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "Project")
		decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", project.Name)
		identifiersDecl += decl

		initializerStatements += "\n"
		// Initialisation of values
		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Name")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(project.Name))
		initializerStatements += setValueField

		setValueField = NumberInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "IsPBSNodeExpanded")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%t", project.IsPBSNodeExpanded))
		initializerStatements += setValueField

		setValueField = NumberInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "IsWBSNodeExpanded")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%t", project.IsWBSNodeExpanded))
		initializerStatements += setValueField

		setValueField = NumberInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "IsDiagramsNodeExpanded")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%t", project.IsDiagramsNodeExpanded))
		initializerStatements += setValueField

		setValueField = NumberInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "IsExpanded")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%t", project.IsExpanded))
		initializerStatements += setValueField

		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "ComputedPrefix")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(project.ComputedPrefix))
		initializerStatements += setValueField

	}

	map_Root_Identifiers := make(map[*Root]string)
	_ = map_Root_Identifiers

	rootOrdered := []*Root{}
	for root := range stage.Roots {
		rootOrdered = append(rootOrdered, root)
	}
	sort.Slice(rootOrdered[:], func(i, j int) bool {
		rooti := rootOrdered[i]
		rootj := rootOrdered[j]
		rooti_order, oki := stage.RootMap_Staged_Order[rooti]
		rootj_order, okj := stage.RootMap_Staged_Order[rootj]
		if !oki || !okj {
			log.Fatalln("unknown pointers")
		}
		return rooti_order < rootj_order
	})
	if len(rootOrdered) > 0 {
		identifiersDecl += "\n"
	}
	for _, root := range rootOrdered {

		id = generatesIdentifier("Root", int(stage.RootMap_Staged_Order[root]), root.Name)
		map_Root_Identifiers[root] = id

		decl = IdentifiersDecls
		decl = strings.ReplaceAll(decl, "{{Identifier}}", id)
		decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "Root")
		decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", root.Name)
		identifiersDecl += decl

		initializerStatements += "\n"
		// Initialisation of values
		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Name")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(root.Name))
		initializerStatements += setValueField

		setValueField = NumberInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "NbPixPerCharacter")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%f", root.NbPixPerCharacter))
		initializerStatements += setValueField

	}

	map_Task_Identifiers := make(map[*Task]string)
	_ = map_Task_Identifiers

	taskOrdered := []*Task{}
	for task := range stage.Tasks {
		taskOrdered = append(taskOrdered, task)
	}
	sort.Slice(taskOrdered[:], func(i, j int) bool {
		taski := taskOrdered[i]
		taskj := taskOrdered[j]
		taski_order, oki := stage.TaskMap_Staged_Order[taski]
		taskj_order, okj := stage.TaskMap_Staged_Order[taskj]
		if !oki || !okj {
			log.Fatalln("unknown pointers")
		}
		return taski_order < taskj_order
	})
	if len(taskOrdered) > 0 {
		identifiersDecl += "\n"
	}
	for _, task := range taskOrdered {

		id = generatesIdentifier("Task", int(stage.TaskMap_Staged_Order[task]), task.Name)
		map_Task_Identifiers[task] = id

		decl = IdentifiersDecls
		decl = strings.ReplaceAll(decl, "{{Identifier}}", id)
		decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "Task")
		decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", task.Name)
		identifiersDecl += decl

		initializerStatements += "\n"
		// Initialisation of values
		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Name")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(task.Name))
		initializerStatements += setValueField

		setValueField = NumberInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "IsExpanded")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%t", task.IsExpanded))
		initializerStatements += setValueField

		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "ComputedPrefix")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(task.ComputedPrefix))
		initializerStatements += setValueField

		setValueField = NumberInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "IsInputsNodeExpanded")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%t", task.IsInputsNodeExpanded))
		initializerStatements += setValueField

		setValueField = NumberInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "IsOutputsNodeExpanded")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%t", task.IsOutputsNodeExpanded))
		initializerStatements += setValueField

	}

	map_TaskShape_Identifiers := make(map[*TaskShape]string)
	_ = map_TaskShape_Identifiers

	taskshapeOrdered := []*TaskShape{}
	for taskshape := range stage.TaskShapes {
		taskshapeOrdered = append(taskshapeOrdered, taskshape)
	}
	sort.Slice(taskshapeOrdered[:], func(i, j int) bool {
		taskshapei := taskshapeOrdered[i]
		taskshapej := taskshapeOrdered[j]
		taskshapei_order, oki := stage.TaskShapeMap_Staged_Order[taskshapei]
		taskshapej_order, okj := stage.TaskShapeMap_Staged_Order[taskshapej]
		if !oki || !okj {
			log.Fatalln("unknown pointers")
		}
		return taskshapei_order < taskshapej_order
	})
	if len(taskshapeOrdered) > 0 {
		identifiersDecl += "\n"
	}
	for _, taskshape := range taskshapeOrdered {

		id = generatesIdentifier("TaskShape", int(stage.TaskShapeMap_Staged_Order[taskshape]), taskshape.Name)
		map_TaskShape_Identifiers[taskshape] = id

		decl = IdentifiersDecls
		decl = strings.ReplaceAll(decl, "{{Identifier}}", id)
		decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "TaskShape")
		decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", taskshape.Name)
		identifiersDecl += decl

		initializerStatements += "\n"
		// Initialisation of values
		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Name")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(taskshape.Name))
		initializerStatements += setValueField

		setValueField = NumberInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "IsExpanded")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%t", taskshape.IsExpanded))
		initializerStatements += setValueField

		setValueField = NumberInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "X")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%f", taskshape.X))
		initializerStatements += setValueField

		setValueField = NumberInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Y")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%f", taskshape.Y))
		initializerStatements += setValueField

		setValueField = NumberInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Width")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%f", taskshape.Width))
		initializerStatements += setValueField

		setValueField = NumberInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Height")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%f", taskshape.Height))
		initializerStatements += setValueField

	}

	// insertion initialization of objects to stage
	if len(compositionshapeOrdered) > 0 {
		pointersInitializesStatements += "\n\t// setup of CompositionShape instances pointers"
	}
	for _, compositionshape := range compositionshapeOrdered {
		var setPointerField string
		_ = setPointerField

		id = generatesIdentifier("CompositionShape", int(stage.CompositionShapeMap_Staged_Order[compositionshape]), compositionshape.Name)
		map_CompositionShape_Identifiers[compositionshape] = id

		// Initialisation of values
		if compositionshape.Product != nil {
			setPointerField = PointerFieldInitStatement
			setPointerField = strings.ReplaceAll(setPointerField, "{{Identifier}}", id)
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldName}}", "Product")
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldNameValue}}", map_Product_Identifiers[compositionshape.Product])
			pointersInitializesStatements += setPointerField
		}

	}

	if len(diagramOrdered) > 0 {
		pointersInitializesStatements += "\n\t// setup of Diagram instances pointers"
	}
	for _, diagram := range diagramOrdered {
		var setPointerField string
		_ = setPointerField

		id = generatesIdentifier("Diagram", int(stage.DiagramMap_Staged_Order[diagram]), diagram.Name)
		map_Diagram_Identifiers[diagram] = id

		// Initialisation of values
		for _, _productshape := range diagram.Product_Shapes {
			setPointerField = SliceOfPointersFieldInitStatement
			setPointerField = strings.ReplaceAll(setPointerField, "{{Identifier}}", id)
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldName}}", "Product_Shapes")
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldNameValue}}", map_ProductShape_Identifiers[_productshape])
			pointersInitializesStatements += setPointerField
		}

		for _, _product := range diagram.ProductsWhoseNodeIsExpanded {
			setPointerField = SliceOfPointersFieldInitStatement
			setPointerField = strings.ReplaceAll(setPointerField, "{{Identifier}}", id)
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldName}}", "ProductsWhoseNodeIsExpanded")
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldNameValue}}", map_Product_Identifiers[_product])
			pointersInitializesStatements += setPointerField
		}

		for _, _compositionshape := range diagram.Composition_Shapes {
			setPointerField = SliceOfPointersFieldInitStatement
			setPointerField = strings.ReplaceAll(setPointerField, "{{Identifier}}", id)
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldName}}", "Composition_Shapes")
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldNameValue}}", map_CompositionShape_Identifiers[_compositionshape])
			pointersInitializesStatements += setPointerField
		}

		for _, _taskshape := range diagram.Task_Shapes {
			setPointerField = SliceOfPointersFieldInitStatement
			setPointerField = strings.ReplaceAll(setPointerField, "{{Identifier}}", id)
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldName}}", "Task_Shapes")
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldNameValue}}", map_TaskShape_Identifiers[_taskshape])
			pointersInitializesStatements += setPointerField
		}

		for _, _task := range diagram.TasksWhoseNodeIsExpanded {
			setPointerField = SliceOfPointersFieldInitStatement
			setPointerField = strings.ReplaceAll(setPointerField, "{{Identifier}}", id)
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldName}}", "TasksWhoseNodeIsExpanded")
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldNameValue}}", map_Task_Identifiers[_task])
			pointersInitializesStatements += setPointerField
		}

	}

	if len(productOrdered) > 0 {
		pointersInitializesStatements += "\n\t// setup of Product instances pointers"
	}
	for _, product := range productOrdered {
		var setPointerField string
		_ = setPointerField

		id = generatesIdentifier("Product", int(stage.ProductMap_Staged_Order[product]), product.Name)
		map_Product_Identifiers[product] = id

		// Initialisation of values
		for _, _product := range product.SubProducts {
			setPointerField = SliceOfPointersFieldInitStatement
			setPointerField = strings.ReplaceAll(setPointerField, "{{Identifier}}", id)
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldName}}", "SubProducts")
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldNameValue}}", map_Product_Identifiers[_product])
			pointersInitializesStatements += setPointerField
		}

	}

	if len(productshapeOrdered) > 0 {
		pointersInitializesStatements += "\n\t// setup of ProductShape instances pointers"
	}
	for _, productshape := range productshapeOrdered {
		var setPointerField string
		_ = setPointerField

		id = generatesIdentifier("ProductShape", int(stage.ProductShapeMap_Staged_Order[productshape]), productshape.Name)
		map_ProductShape_Identifiers[productshape] = id

		// Initialisation of values
		if productshape.Product != nil {
			setPointerField = PointerFieldInitStatement
			setPointerField = strings.ReplaceAll(setPointerField, "{{Identifier}}", id)
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldName}}", "Product")
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldNameValue}}", map_Product_Identifiers[productshape.Product])
			pointersInitializesStatements += setPointerField
		}

	}

	if len(projectOrdered) > 0 {
		pointersInitializesStatements += "\n\t// setup of Project instances pointers"
	}
	for _, project := range projectOrdered {
		var setPointerField string
		_ = setPointerField

		id = generatesIdentifier("Project", int(stage.ProjectMap_Staged_Order[project]), project.Name)
		map_Project_Identifiers[project] = id

		// Initialisation of values
		for _, _product := range project.RootProducts {
			setPointerField = SliceOfPointersFieldInitStatement
			setPointerField = strings.ReplaceAll(setPointerField, "{{Identifier}}", id)
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldName}}", "RootProducts")
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldNameValue}}", map_Product_Identifiers[_product])
			pointersInitializesStatements += setPointerField
		}

		for _, _task := range project.RootTasks {
			setPointerField = SliceOfPointersFieldInitStatement
			setPointerField = strings.ReplaceAll(setPointerField, "{{Identifier}}", id)
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldName}}", "RootTasks")
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldNameValue}}", map_Task_Identifiers[_task])
			pointersInitializesStatements += setPointerField
		}

		for _, _diagram := range project.Diagrams {
			setPointerField = SliceOfPointersFieldInitStatement
			setPointerField = strings.ReplaceAll(setPointerField, "{{Identifier}}", id)
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldName}}", "Diagrams")
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldNameValue}}", map_Diagram_Identifiers[_diagram])
			pointersInitializesStatements += setPointerField
		}

	}

	if len(rootOrdered) > 0 {
		pointersInitializesStatements += "\n\t// setup of Root instances pointers"
	}
	for _, root := range rootOrdered {
		var setPointerField string
		_ = setPointerField

		id = generatesIdentifier("Root", int(stage.RootMap_Staged_Order[root]), root.Name)
		map_Root_Identifiers[root] = id

		// Initialisation of values
		for _, _project := range root.Projects {
			setPointerField = SliceOfPointersFieldInitStatement
			setPointerField = strings.ReplaceAll(setPointerField, "{{Identifier}}", id)
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldName}}", "Projects")
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldNameValue}}", map_Project_Identifiers[_project])
			pointersInitializesStatements += setPointerField
		}

		for _, _product := range root.OrphanedProducts {
			setPointerField = SliceOfPointersFieldInitStatement
			setPointerField = strings.ReplaceAll(setPointerField, "{{Identifier}}", id)
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldName}}", "OrphanedProducts")
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldNameValue}}", map_Product_Identifiers[_product])
			pointersInitializesStatements += setPointerField
		}

		for _, _task := range root.OrphanedTasks {
			setPointerField = SliceOfPointersFieldInitStatement
			setPointerField = strings.ReplaceAll(setPointerField, "{{Identifier}}", id)
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldName}}", "OrphanedTasks")
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldNameValue}}", map_Task_Identifiers[_task])
			pointersInitializesStatements += setPointerField
		}

	}

	if len(taskOrdered) > 0 {
		pointersInitializesStatements += "\n\t// setup of Task instances pointers"
	}
	for _, task := range taskOrdered {
		var setPointerField string
		_ = setPointerField

		id = generatesIdentifier("Task", int(stage.TaskMap_Staged_Order[task]), task.Name)
		map_Task_Identifiers[task] = id

		// Initialisation of values
		for _, _task := range task.SubTasks {
			setPointerField = SliceOfPointersFieldInitStatement
			setPointerField = strings.ReplaceAll(setPointerField, "{{Identifier}}", id)
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldName}}", "SubTasks")
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldNameValue}}", map_Task_Identifiers[_task])
			pointersInitializesStatements += setPointerField
		}

		for _, _product := range task.Inputs {
			setPointerField = SliceOfPointersFieldInitStatement
			setPointerField = strings.ReplaceAll(setPointerField, "{{Identifier}}", id)
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldName}}", "Inputs")
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldNameValue}}", map_Product_Identifiers[_product])
			pointersInitializesStatements += setPointerField
		}

		for _, _product := range task.Outputs {
			setPointerField = SliceOfPointersFieldInitStatement
			setPointerField = strings.ReplaceAll(setPointerField, "{{Identifier}}", id)
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldName}}", "Outputs")
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldNameValue}}", map_Product_Identifiers[_product])
			pointersInitializesStatements += setPointerField
		}

	}

	if len(taskshapeOrdered) > 0 {
		pointersInitializesStatements += "\n\t// setup of TaskShape instances pointers"
	}
	for _, taskshape := range taskshapeOrdered {
		var setPointerField string
		_ = setPointerField

		id = generatesIdentifier("TaskShape", int(stage.TaskShapeMap_Staged_Order[taskshape]), taskshape.Name)
		map_TaskShape_Identifiers[taskshape] = id

		// Initialisation of values
		if taskshape.Task != nil {
			setPointerField = PointerFieldInitStatement
			setPointerField = strings.ReplaceAll(setPointerField, "{{Identifier}}", id)
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldName}}", "Task")
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldNameValue}}", map_Task_Identifiers[taskshape.Task])
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

	identifier = fmt.Sprintf("__%s__%08d_%s", gongStructName, idx, processedString)

	return
}
