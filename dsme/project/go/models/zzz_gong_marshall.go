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

		identifiersDecl += diagram.GongMarshallIdentifier(stage)

		initializerStatements += "\n"
		// Insertion point for basic fields value assignment
		initializerStatements += diagram.GongMarshallField(stage, "Name")
		initializerStatements += diagram.GongMarshallField(stage, "IsChecked")
		initializerStatements += diagram.GongMarshallField(stage, "IsEditable_")
		initializerStatements += diagram.GongMarshallField(stage, "IsInRenameMode")
		initializerStatements += diagram.GongMarshallField(stage, "ShowPrefix")
		initializerStatements += diagram.GongMarshallField(stage, "DefaultBoxWidth")
		initializerStatements += diagram.GongMarshallField(stage, "DefaultBoxHeigth")
		initializerStatements += diagram.GongMarshallField(stage, "IsExpanded")
		initializerStatements += diagram.GongMarshallField(stage, "ComputedPrefix")
		pointersInitializesStatements += diagram.GongMarshallField(stage, "Product_Shapes")
		pointersInitializesStatements += diagram.GongMarshallField(stage, "ProductsWhoseNodeIsExpanded")
		initializerStatements += diagram.GongMarshallField(stage, "IsPBSNodeExpanded")
		pointersInitializesStatements += diagram.GongMarshallField(stage, "ProductComposition_Shapes")
		initializerStatements += diagram.GongMarshallField(stage, "IsWBSNodeExpanded")
		pointersInitializesStatements += diagram.GongMarshallField(stage, "Task_Shapes")
		pointersInitializesStatements += diagram.GongMarshallField(stage, "TasksWhoseNodeIsExpanded")
		pointersInitializesStatements += diagram.GongMarshallField(stage, "TasksWhoseInputNodeIsExpanded")
		pointersInitializesStatements += diagram.GongMarshallField(stage, "TasksWhoseOutputNodeIsExpanded")
		pointersInitializesStatements += diagram.GongMarshallField(stage, "TaskComposition_Shapes")
		pointersInitializesStatements += diagram.GongMarshallField(stage, "TaskInputShapes")
		pointersInitializesStatements += diagram.GongMarshallField(stage, "TaskOutputShapes")
		pointersInitializesStatements += diagram.GongMarshallField(stage, "Note_Shapes")
		pointersInitializesStatements += diagram.GongMarshallField(stage, "NotesWhoseNodeIsExpanded")
		initializerStatements += diagram.GongMarshallField(stage, "IsNotesNodeExpanded")
		pointersInitializesStatements += diagram.GongMarshallField(stage, "NoteProductShapes")
		pointersInitializesStatements += diagram.GongMarshallField(stage, "NoteTaskShapes")
	}

	noteOrdered := []*Note{}
	for note := range stage.Notes {
		noteOrdered = append(noteOrdered, note)
	}
	sort.Slice(noteOrdered[:], func(i, j int) bool {
		notei := noteOrdered[i]
		notej := noteOrdered[j]
		notei_order, oki := stage.NoteMap_Staged_Order[notei]
		notej_order, okj := stage.NoteMap_Staged_Order[notej]
		if !oki || !okj {
			log.Fatalln("unknown pointers")
		}
		return notei_order < notej_order
	})
	if len(noteOrdered) > 0 {
		identifiersDecl += "\n"
	}
	for _, note := range noteOrdered {

		identifiersDecl += note.GongMarshallIdentifier(stage)

		initializerStatements += "\n"
		// Insertion point for basic fields value assignment
		initializerStatements += note.GongMarshallField(stage, "Name")
		pointersInitializesStatements += note.GongMarshallField(stage, "Products")
		pointersInitializesStatements += note.GongMarshallField(stage, "Tasks")
		initializerStatements += note.GongMarshallField(stage, "IsExpanded")
	}

	noteproductshapeOrdered := []*NoteProductShape{}
	for noteproductshape := range stage.NoteProductShapes {
		noteproductshapeOrdered = append(noteproductshapeOrdered, noteproductshape)
	}
	sort.Slice(noteproductshapeOrdered[:], func(i, j int) bool {
		noteproductshapei := noteproductshapeOrdered[i]
		noteproductshapej := noteproductshapeOrdered[j]
		noteproductshapei_order, oki := stage.NoteProductShapeMap_Staged_Order[noteproductshapei]
		noteproductshapej_order, okj := stage.NoteProductShapeMap_Staged_Order[noteproductshapej]
		if !oki || !okj {
			log.Fatalln("unknown pointers")
		}
		return noteproductshapei_order < noteproductshapej_order
	})
	if len(noteproductshapeOrdered) > 0 {
		identifiersDecl += "\n"
	}
	for _, noteproductshape := range noteproductshapeOrdered {

		identifiersDecl += noteproductshape.GongMarshallIdentifier(stage)

		initializerStatements += "\n"
		// Insertion point for basic fields value assignment
		initializerStatements += noteproductshape.GongMarshallField(stage, "Name")
		pointersInitializesStatements += noteproductshape.GongMarshallField(stage, "Note")
		pointersInitializesStatements += noteproductshape.GongMarshallField(stage, "Product")
		initializerStatements += noteproductshape.GongMarshallField(stage, "StartRatio")
		initializerStatements += noteproductshape.GongMarshallField(stage, "EndRatio")
		initializerStatements += noteproductshape.GongMarshallField(stage, "StartOrientation")
		initializerStatements += noteproductshape.GongMarshallField(stage, "EndOrientation")
		initializerStatements += noteproductshape.GongMarshallField(stage, "CornerOffsetRatio")
	}

	noteshapeOrdered := []*NoteShape{}
	for noteshape := range stage.NoteShapes {
		noteshapeOrdered = append(noteshapeOrdered, noteshape)
	}
	sort.Slice(noteshapeOrdered[:], func(i, j int) bool {
		noteshapei := noteshapeOrdered[i]
		noteshapej := noteshapeOrdered[j]
		noteshapei_order, oki := stage.NoteShapeMap_Staged_Order[noteshapei]
		noteshapej_order, okj := stage.NoteShapeMap_Staged_Order[noteshapej]
		if !oki || !okj {
			log.Fatalln("unknown pointers")
		}
		return noteshapei_order < noteshapej_order
	})
	if len(noteshapeOrdered) > 0 {
		identifiersDecl += "\n"
	}
	for _, noteshape := range noteshapeOrdered {

		identifiersDecl += noteshape.GongMarshallIdentifier(stage)

		initializerStatements += "\n"
		// Insertion point for basic fields value assignment
		initializerStatements += noteshape.GongMarshallField(stage, "Name")
		pointersInitializesStatements += noteshape.GongMarshallField(stage, "Note")
		initializerStatements += noteshape.GongMarshallField(stage, "IsExpanded")
		initializerStatements += noteshape.GongMarshallField(stage, "X")
		initializerStatements += noteshape.GongMarshallField(stage, "Y")
		initializerStatements += noteshape.GongMarshallField(stage, "Width")
		initializerStatements += noteshape.GongMarshallField(stage, "Height")
	}

	notetaskshapeOrdered := []*NoteTaskShape{}
	for notetaskshape := range stage.NoteTaskShapes {
		notetaskshapeOrdered = append(notetaskshapeOrdered, notetaskshape)
	}
	sort.Slice(notetaskshapeOrdered[:], func(i, j int) bool {
		notetaskshapei := notetaskshapeOrdered[i]
		notetaskshapej := notetaskshapeOrdered[j]
		notetaskshapei_order, oki := stage.NoteTaskShapeMap_Staged_Order[notetaskshapei]
		notetaskshapej_order, okj := stage.NoteTaskShapeMap_Staged_Order[notetaskshapej]
		if !oki || !okj {
			log.Fatalln("unknown pointers")
		}
		return notetaskshapei_order < notetaskshapej_order
	})
	if len(notetaskshapeOrdered) > 0 {
		identifiersDecl += "\n"
	}
	for _, notetaskshape := range notetaskshapeOrdered {

		identifiersDecl += notetaskshape.GongMarshallIdentifier(stage)

		initializerStatements += "\n"
		// Insertion point for basic fields value assignment
		initializerStatements += notetaskshape.GongMarshallField(stage, "Name")
		pointersInitializesStatements += notetaskshape.GongMarshallField(stage, "Note")
		pointersInitializesStatements += notetaskshape.GongMarshallField(stage, "Task")
		initializerStatements += notetaskshape.GongMarshallField(stage, "StartRatio")
		initializerStatements += notetaskshape.GongMarshallField(stage, "EndRatio")
		initializerStatements += notetaskshape.GongMarshallField(stage, "StartOrientation")
		initializerStatements += notetaskshape.GongMarshallField(stage, "EndOrientation")
		initializerStatements += notetaskshape.GongMarshallField(stage, "CornerOffsetRatio")
	}

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

		identifiersDecl += product.GongMarshallIdentifier(stage)

		initializerStatements += "\n"
		// Insertion point for basic fields value assignment
		initializerStatements += product.GongMarshallField(stage, "Name")
		initializerStatements += product.GongMarshallField(stage, "Description")
		pointersInitializesStatements += product.GongMarshallField(stage, "SubProducts")
		initializerStatements += product.GongMarshallField(stage, "IsExpanded")
		initializerStatements += product.GongMarshallField(stage, "ComputedPrefix")
		initializerStatements += product.GongMarshallField(stage, "IsProducersNodeExpanded")
		initializerStatements += product.GongMarshallField(stage, "IsConsumersNodeExpanded")
	}

	productcompositionshapeOrdered := []*ProductCompositionShape{}
	for productcompositionshape := range stage.ProductCompositionShapes {
		productcompositionshapeOrdered = append(productcompositionshapeOrdered, productcompositionshape)
	}
	sort.Slice(productcompositionshapeOrdered[:], func(i, j int) bool {
		productcompositionshapei := productcompositionshapeOrdered[i]
		productcompositionshapej := productcompositionshapeOrdered[j]
		productcompositionshapei_order, oki := stage.ProductCompositionShapeMap_Staged_Order[productcompositionshapei]
		productcompositionshapej_order, okj := stage.ProductCompositionShapeMap_Staged_Order[productcompositionshapej]
		if !oki || !okj {
			log.Fatalln("unknown pointers")
		}
		return productcompositionshapei_order < productcompositionshapej_order
	})
	if len(productcompositionshapeOrdered) > 0 {
		identifiersDecl += "\n"
	}
	for _, productcompositionshape := range productcompositionshapeOrdered {

		identifiersDecl += productcompositionshape.GongMarshallIdentifier(stage)

		initializerStatements += "\n"
		// Insertion point for basic fields value assignment
		initializerStatements += productcompositionshape.GongMarshallField(stage, "Name")
		pointersInitializesStatements += productcompositionshape.GongMarshallField(stage, "Product")
		initializerStatements += productcompositionshape.GongMarshallField(stage, "StartRatio")
		initializerStatements += productcompositionshape.GongMarshallField(stage, "EndRatio")
		initializerStatements += productcompositionshape.GongMarshallField(stage, "StartOrientation")
		initializerStatements += productcompositionshape.GongMarshallField(stage, "EndOrientation")
		initializerStatements += productcompositionshape.GongMarshallField(stage, "CornerOffsetRatio")
	}

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

		identifiersDecl += productshape.GongMarshallIdentifier(stage)

		initializerStatements += "\n"
		// Insertion point for basic fields value assignment
		initializerStatements += productshape.GongMarshallField(stage, "Name")
		pointersInitializesStatements += productshape.GongMarshallField(stage, "Product")
		initializerStatements += productshape.GongMarshallField(stage, "IsExpanded")
		initializerStatements += productshape.GongMarshallField(stage, "X")
		initializerStatements += productshape.GongMarshallField(stage, "Y")
		initializerStatements += productshape.GongMarshallField(stage, "Width")
		initializerStatements += productshape.GongMarshallField(stage, "Height")
	}

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

		identifiersDecl += project.GongMarshallIdentifier(stage)

		initializerStatements += "\n"
		// Insertion point for basic fields value assignment
		initializerStatements += project.GongMarshallField(stage, "Name")
		pointersInitializesStatements += project.GongMarshallField(stage, "RootProducts")
		initializerStatements += project.GongMarshallField(stage, "IsPBSNodeExpanded")
		pointersInitializesStatements += project.GongMarshallField(stage, "RootTasks")
		initializerStatements += project.GongMarshallField(stage, "IsWBSNodeExpanded")
		pointersInitializesStatements += project.GongMarshallField(stage, "Diagrams")
		initializerStatements += project.GongMarshallField(stage, "IsDiagramsNodeExpanded")
		pointersInitializesStatements += project.GongMarshallField(stage, "Notes")
		initializerStatements += project.GongMarshallField(stage, "IsNotesNodeExpanded")
		initializerStatements += project.GongMarshallField(stage, "IsExpanded")
		initializerStatements += project.GongMarshallField(stage, "ComputedPrefix")
	}

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

		identifiersDecl += root.GongMarshallIdentifier(stage)

		initializerStatements += "\n"
		// Insertion point for basic fields value assignment
		initializerStatements += root.GongMarshallField(stage, "Name")
		pointersInitializesStatements += root.GongMarshallField(stage, "Projects")
		pointersInitializesStatements += root.GongMarshallField(stage, "OrphanedProducts")
		pointersInitializesStatements += root.GongMarshallField(stage, "OrphanedTasks")
		initializerStatements += root.GongMarshallField(stage, "NbPixPerCharacter")
	}

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

		identifiersDecl += task.GongMarshallIdentifier(stage)

		initializerStatements += "\n"
		// Insertion point for basic fields value assignment
		initializerStatements += task.GongMarshallField(stage, "Name")
		initializerStatements += task.GongMarshallField(stage, "Description")
		pointersInitializesStatements += task.GongMarshallField(stage, "SubTasks")
		initializerStatements += task.GongMarshallField(stage, "IsExpanded")
		initializerStatements += task.GongMarshallField(stage, "ComputedPrefix")
		pointersInitializesStatements += task.GongMarshallField(stage, "Inputs")
		initializerStatements += task.GongMarshallField(stage, "IsInputsNodeExpanded")
		pointersInitializesStatements += task.GongMarshallField(stage, "Outputs")
		initializerStatements += task.GongMarshallField(stage, "IsOutputsNodeExpanded")
	}

	taskcompositionshapeOrdered := []*TaskCompositionShape{}
	for taskcompositionshape := range stage.TaskCompositionShapes {
		taskcompositionshapeOrdered = append(taskcompositionshapeOrdered, taskcompositionshape)
	}
	sort.Slice(taskcompositionshapeOrdered[:], func(i, j int) bool {
		taskcompositionshapei := taskcompositionshapeOrdered[i]
		taskcompositionshapej := taskcompositionshapeOrdered[j]
		taskcompositionshapei_order, oki := stage.TaskCompositionShapeMap_Staged_Order[taskcompositionshapei]
		taskcompositionshapej_order, okj := stage.TaskCompositionShapeMap_Staged_Order[taskcompositionshapej]
		if !oki || !okj {
			log.Fatalln("unknown pointers")
		}
		return taskcompositionshapei_order < taskcompositionshapej_order
	})
	if len(taskcompositionshapeOrdered) > 0 {
		identifiersDecl += "\n"
	}
	for _, taskcompositionshape := range taskcompositionshapeOrdered {

		identifiersDecl += taskcompositionshape.GongMarshallIdentifier(stage)

		initializerStatements += "\n"
		// Insertion point for basic fields value assignment
		initializerStatements += taskcompositionshape.GongMarshallField(stage, "Name")
		pointersInitializesStatements += taskcompositionshape.GongMarshallField(stage, "Task")
		initializerStatements += taskcompositionshape.GongMarshallField(stage, "StartRatio")
		initializerStatements += taskcompositionshape.GongMarshallField(stage, "EndRatio")
		initializerStatements += taskcompositionshape.GongMarshallField(stage, "StartOrientation")
		initializerStatements += taskcompositionshape.GongMarshallField(stage, "EndOrientation")
		initializerStatements += taskcompositionshape.GongMarshallField(stage, "CornerOffsetRatio")
	}

	taskinputshapeOrdered := []*TaskInputShape{}
	for taskinputshape := range stage.TaskInputShapes {
		taskinputshapeOrdered = append(taskinputshapeOrdered, taskinputshape)
	}
	sort.Slice(taskinputshapeOrdered[:], func(i, j int) bool {
		taskinputshapei := taskinputshapeOrdered[i]
		taskinputshapej := taskinputshapeOrdered[j]
		taskinputshapei_order, oki := stage.TaskInputShapeMap_Staged_Order[taskinputshapei]
		taskinputshapej_order, okj := stage.TaskInputShapeMap_Staged_Order[taskinputshapej]
		if !oki || !okj {
			log.Fatalln("unknown pointers")
		}
		return taskinputshapei_order < taskinputshapej_order
	})
	if len(taskinputshapeOrdered) > 0 {
		identifiersDecl += "\n"
	}
	for _, taskinputshape := range taskinputshapeOrdered {

		identifiersDecl += taskinputshape.GongMarshallIdentifier(stage)

		initializerStatements += "\n"
		// Insertion point for basic fields value assignment
		initializerStatements += taskinputshape.GongMarshallField(stage, "Name")
		pointersInitializesStatements += taskinputshape.GongMarshallField(stage, "Task")
		pointersInitializesStatements += taskinputshape.GongMarshallField(stage, "Product")
		initializerStatements += taskinputshape.GongMarshallField(stage, "StartRatio")
		initializerStatements += taskinputshape.GongMarshallField(stage, "EndRatio")
		initializerStatements += taskinputshape.GongMarshallField(stage, "StartOrientation")
		initializerStatements += taskinputshape.GongMarshallField(stage, "EndOrientation")
		initializerStatements += taskinputshape.GongMarshallField(stage, "CornerOffsetRatio")
	}

	taskoutputshapeOrdered := []*TaskOutputShape{}
	for taskoutputshape := range stage.TaskOutputShapes {
		taskoutputshapeOrdered = append(taskoutputshapeOrdered, taskoutputshape)
	}
	sort.Slice(taskoutputshapeOrdered[:], func(i, j int) bool {
		taskoutputshapei := taskoutputshapeOrdered[i]
		taskoutputshapej := taskoutputshapeOrdered[j]
		taskoutputshapei_order, oki := stage.TaskOutputShapeMap_Staged_Order[taskoutputshapei]
		taskoutputshapej_order, okj := stage.TaskOutputShapeMap_Staged_Order[taskoutputshapej]
		if !oki || !okj {
			log.Fatalln("unknown pointers")
		}
		return taskoutputshapei_order < taskoutputshapej_order
	})
	if len(taskoutputshapeOrdered) > 0 {
		identifiersDecl += "\n"
	}
	for _, taskoutputshape := range taskoutputshapeOrdered {

		identifiersDecl += taskoutputshape.GongMarshallIdentifier(stage)

		initializerStatements += "\n"
		// Insertion point for basic fields value assignment
		initializerStatements += taskoutputshape.GongMarshallField(stage, "Name")
		pointersInitializesStatements += taskoutputshape.GongMarshallField(stage, "Task")
		pointersInitializesStatements += taskoutputshape.GongMarshallField(stage, "Product")
		initializerStatements += taskoutputshape.GongMarshallField(stage, "StartRatio")
		initializerStatements += taskoutputshape.GongMarshallField(stage, "EndRatio")
		initializerStatements += taskoutputshape.GongMarshallField(stage, "StartOrientation")
		initializerStatements += taskoutputshape.GongMarshallField(stage, "EndOrientation")
		initializerStatements += taskoutputshape.GongMarshallField(stage, "CornerOffsetRatio")
	}

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

		identifiersDecl += taskshape.GongMarshallIdentifier(stage)

		initializerStatements += "\n"
		// Insertion point for basic fields value assignment
		initializerStatements += taskshape.GongMarshallField(stage, "Name")
		pointersInitializesStatements += taskshape.GongMarshallField(stage, "Task")
		initializerStatements += taskshape.GongMarshallField(stage, "IsExpanded")
		initializerStatements += taskshape.GongMarshallField(stage, "X")
		initializerStatements += taskshape.GongMarshallField(stage, "Y")
		initializerStatements += taskshape.GongMarshallField(stage, "Width")
		initializerStatements += taskshape.GongMarshallField(stage, "Height")
	}

	// insertion initialization of objects to stage
	for _, diagram := range diagramOrdered {
		_ = diagram
		var setPointerField string
		_ = setPointerField

		// Insertion point for pointers initialization
	}

	for _, note := range noteOrdered {
		_ = note
		var setPointerField string
		_ = setPointerField

		// Insertion point for pointers initialization
	}

	for _, noteproductshape := range noteproductshapeOrdered {
		_ = noteproductshape
		var setPointerField string
		_ = setPointerField

		// Insertion point for pointers initialization
	}

	for _, noteshape := range noteshapeOrdered {
		_ = noteshape
		var setPointerField string
		_ = setPointerField

		// Insertion point for pointers initialization
	}

	for _, notetaskshape := range notetaskshapeOrdered {
		_ = notetaskshape
		var setPointerField string
		_ = setPointerField

		// Insertion point for pointers initialization
	}

	for _, product := range productOrdered {
		_ = product
		var setPointerField string
		_ = setPointerField

		// Insertion point for pointers initialization
	}

	for _, productcompositionshape := range productcompositionshapeOrdered {
		_ = productcompositionshape
		var setPointerField string
		_ = setPointerField

		// Insertion point for pointers initialization
	}

	for _, productshape := range productshapeOrdered {
		_ = productshape
		var setPointerField string
		_ = setPointerField

		// Insertion point for pointers initialization
	}

	for _, project := range projectOrdered {
		_ = project
		var setPointerField string
		_ = setPointerField

		// Insertion point for pointers initialization
	}

	for _, root := range rootOrdered {
		_ = root
		var setPointerField string
		_ = setPointerField

		// Insertion point for pointers initialization
	}

	for _, task := range taskOrdered {
		_ = task
		var setPointerField string
		_ = setPointerField

		// Insertion point for pointers initialization
	}

	for _, taskcompositionshape := range taskcompositionshapeOrdered {
		_ = taskcompositionshape
		var setPointerField string
		_ = setPointerField

		// Insertion point for pointers initialization
	}

	for _, taskinputshape := range taskinputshapeOrdered {
		_ = taskinputshape
		var setPointerField string
		_ = setPointerField

		// Insertion point for pointers initialization
	}

	for _, taskoutputshape := range taskoutputshapeOrdered {
		_ = taskoutputshape
		var setPointerField string
		_ = setPointerField

		// Insertion point for pointers initialization
	}

	for _, taskshape := range taskshapeOrdered {
		_ = taskshape
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
func (diagram *Diagram) GongMarshallField(stage *Stage, fieldName string) (res string) {

	switch fieldName {
	case "Name":
		res = StringInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", diagram.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "Name")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", string(diagram.Name))
	case "IsChecked":
		res = NumberInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", diagram.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "IsChecked")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%t", diagram.IsChecked))
	case "IsEditable_":
		res = NumberInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", diagram.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "IsEditable_")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%t", diagram.IsEditable_))
	case "IsInRenameMode":
		res = NumberInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", diagram.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "IsInRenameMode")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%t", diagram.IsInRenameMode))
	case "ShowPrefix":
		res = NumberInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", diagram.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "ShowPrefix")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%t", diagram.ShowPrefix))
	case "DefaultBoxWidth":
		res = NumberInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", diagram.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "DefaultBoxWidth")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%f", diagram.DefaultBoxWidth))
	case "DefaultBoxHeigth":
		res = NumberInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", diagram.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "DefaultBoxHeigth")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%f", diagram.DefaultBoxHeigth))
	case "IsExpanded":
		res = NumberInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", diagram.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "IsExpanded")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%t", diagram.IsExpanded))
	case "ComputedPrefix":
		res = StringInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", diagram.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "ComputedPrefix")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", string(diagram.ComputedPrefix))
	case "IsPBSNodeExpanded":
		res = NumberInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", diagram.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "IsPBSNodeExpanded")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%t", diagram.IsPBSNodeExpanded))
	case "IsWBSNodeExpanded":
		res = NumberInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", diagram.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "IsWBSNodeExpanded")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%t", diagram.IsWBSNodeExpanded))
	case "IsNotesNodeExpanded":
		res = NumberInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", diagram.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "IsNotesNodeExpanded")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%t", diagram.IsNotesNodeExpanded))

	case "Product_Shapes":
		for _, _productshape := range diagram.Product_Shapes {
			tmp := SliceOfPointersFieldInitStatement
			tmp = strings.ReplaceAll(tmp, "{{Identifier}}", diagram.GongGetIdentifier(stage))
			tmp = strings.ReplaceAll(tmp, "{{GeneratedFieldName}}", "Product_Shapes")
			tmp = strings.ReplaceAll(tmp, "{{GeneratedFieldNameValue}}", _productshape.GongGetIdentifier(stage))
			res += tmp
		}
	case "ProductsWhoseNodeIsExpanded":
		for _, _product := range diagram.ProductsWhoseNodeIsExpanded {
			tmp := SliceOfPointersFieldInitStatement
			tmp = strings.ReplaceAll(tmp, "{{Identifier}}", diagram.GongGetIdentifier(stage))
			tmp = strings.ReplaceAll(tmp, "{{GeneratedFieldName}}", "ProductsWhoseNodeIsExpanded")
			tmp = strings.ReplaceAll(tmp, "{{GeneratedFieldNameValue}}", _product.GongGetIdentifier(stage))
			res += tmp
		}
	case "ProductComposition_Shapes":
		for _, _productcompositionshape := range diagram.ProductComposition_Shapes {
			tmp := SliceOfPointersFieldInitStatement
			tmp = strings.ReplaceAll(tmp, "{{Identifier}}", diagram.GongGetIdentifier(stage))
			tmp = strings.ReplaceAll(tmp, "{{GeneratedFieldName}}", "ProductComposition_Shapes")
			tmp = strings.ReplaceAll(tmp, "{{GeneratedFieldNameValue}}", _productcompositionshape.GongGetIdentifier(stage))
			res += tmp
		}
	case "Task_Shapes":
		for _, _taskshape := range diagram.Task_Shapes {
			tmp := SliceOfPointersFieldInitStatement
			tmp = strings.ReplaceAll(tmp, "{{Identifier}}", diagram.GongGetIdentifier(stage))
			tmp = strings.ReplaceAll(tmp, "{{GeneratedFieldName}}", "Task_Shapes")
			tmp = strings.ReplaceAll(tmp, "{{GeneratedFieldNameValue}}", _taskshape.GongGetIdentifier(stage))
			res += tmp
		}
	case "TasksWhoseNodeIsExpanded":
		for _, _task := range diagram.TasksWhoseNodeIsExpanded {
			tmp := SliceOfPointersFieldInitStatement
			tmp = strings.ReplaceAll(tmp, "{{Identifier}}", diagram.GongGetIdentifier(stage))
			tmp = strings.ReplaceAll(tmp, "{{GeneratedFieldName}}", "TasksWhoseNodeIsExpanded")
			tmp = strings.ReplaceAll(tmp, "{{GeneratedFieldNameValue}}", _task.GongGetIdentifier(stage))
			res += tmp
		}
	case "TasksWhoseInputNodeIsExpanded":
		for _, _task := range diagram.TasksWhoseInputNodeIsExpanded {
			tmp := SliceOfPointersFieldInitStatement
			tmp = strings.ReplaceAll(tmp, "{{Identifier}}", diagram.GongGetIdentifier(stage))
			tmp = strings.ReplaceAll(tmp, "{{GeneratedFieldName}}", "TasksWhoseInputNodeIsExpanded")
			tmp = strings.ReplaceAll(tmp, "{{GeneratedFieldNameValue}}", _task.GongGetIdentifier(stage))
			res += tmp
		}
	case "TasksWhoseOutputNodeIsExpanded":
		for _, _task := range diagram.TasksWhoseOutputNodeIsExpanded {
			tmp := SliceOfPointersFieldInitStatement
			tmp = strings.ReplaceAll(tmp, "{{Identifier}}", diagram.GongGetIdentifier(stage))
			tmp = strings.ReplaceAll(tmp, "{{GeneratedFieldName}}", "TasksWhoseOutputNodeIsExpanded")
			tmp = strings.ReplaceAll(tmp, "{{GeneratedFieldNameValue}}", _task.GongGetIdentifier(stage))
			res += tmp
		}
	case "TaskComposition_Shapes":
		for _, _taskcompositionshape := range diagram.TaskComposition_Shapes {
			tmp := SliceOfPointersFieldInitStatement
			tmp = strings.ReplaceAll(tmp, "{{Identifier}}", diagram.GongGetIdentifier(stage))
			tmp = strings.ReplaceAll(tmp, "{{GeneratedFieldName}}", "TaskComposition_Shapes")
			tmp = strings.ReplaceAll(tmp, "{{GeneratedFieldNameValue}}", _taskcompositionshape.GongGetIdentifier(stage))
			res += tmp
		}
	case "TaskInputShapes":
		for _, _taskinputshape := range diagram.TaskInputShapes {
			tmp := SliceOfPointersFieldInitStatement
			tmp = strings.ReplaceAll(tmp, "{{Identifier}}", diagram.GongGetIdentifier(stage))
			tmp = strings.ReplaceAll(tmp, "{{GeneratedFieldName}}", "TaskInputShapes")
			tmp = strings.ReplaceAll(tmp, "{{GeneratedFieldNameValue}}", _taskinputshape.GongGetIdentifier(stage))
			res += tmp
		}
	case "TaskOutputShapes":
		for _, _taskoutputshape := range diagram.TaskOutputShapes {
			tmp := SliceOfPointersFieldInitStatement
			tmp = strings.ReplaceAll(tmp, "{{Identifier}}", diagram.GongGetIdentifier(stage))
			tmp = strings.ReplaceAll(tmp, "{{GeneratedFieldName}}", "TaskOutputShapes")
			tmp = strings.ReplaceAll(tmp, "{{GeneratedFieldNameValue}}", _taskoutputshape.GongGetIdentifier(stage))
			res += tmp
		}
	case "Note_Shapes":
		for _, _noteshape := range diagram.Note_Shapes {
			tmp := SliceOfPointersFieldInitStatement
			tmp = strings.ReplaceAll(tmp, "{{Identifier}}", diagram.GongGetIdentifier(stage))
			tmp = strings.ReplaceAll(tmp, "{{GeneratedFieldName}}", "Note_Shapes")
			tmp = strings.ReplaceAll(tmp, "{{GeneratedFieldNameValue}}", _noteshape.GongGetIdentifier(stage))
			res += tmp
		}
	case "NotesWhoseNodeIsExpanded":
		for _, _note := range diagram.NotesWhoseNodeIsExpanded {
			tmp := SliceOfPointersFieldInitStatement
			tmp = strings.ReplaceAll(tmp, "{{Identifier}}", diagram.GongGetIdentifier(stage))
			tmp = strings.ReplaceAll(tmp, "{{GeneratedFieldName}}", "NotesWhoseNodeIsExpanded")
			tmp = strings.ReplaceAll(tmp, "{{GeneratedFieldNameValue}}", _note.GongGetIdentifier(stage))
			res += tmp
		}
	case "NoteProductShapes":
		for _, _noteproductshape := range diagram.NoteProductShapes {
			tmp := SliceOfPointersFieldInitStatement
			tmp = strings.ReplaceAll(tmp, "{{Identifier}}", diagram.GongGetIdentifier(stage))
			tmp = strings.ReplaceAll(tmp, "{{GeneratedFieldName}}", "NoteProductShapes")
			tmp = strings.ReplaceAll(tmp, "{{GeneratedFieldNameValue}}", _noteproductshape.GongGetIdentifier(stage))
			res += tmp
		}
	case "NoteTaskShapes":
		for _, _notetaskshape := range diagram.NoteTaskShapes {
			tmp := SliceOfPointersFieldInitStatement
			tmp = strings.ReplaceAll(tmp, "{{Identifier}}", diagram.GongGetIdentifier(stage))
			tmp = strings.ReplaceAll(tmp, "{{GeneratedFieldName}}", "NoteTaskShapes")
			tmp = strings.ReplaceAll(tmp, "{{GeneratedFieldNameValue}}", _notetaskshape.GongGetIdentifier(stage))
			res += tmp
		}
	default:
		log.Panicf("Unknown field %s for Gongstruct Diagram", fieldName)
	}
	return
}

func (note *Note) GongMarshallField(stage *Stage, fieldName string) (res string) {

	switch fieldName {
	case "Name":
		res = StringInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", note.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "Name")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", string(note.Name))
	case "IsExpanded":
		res = NumberInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", note.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "IsExpanded")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%t", note.IsExpanded))

	case "Products":
		for _, _product := range note.Products {
			tmp := SliceOfPointersFieldInitStatement
			tmp = strings.ReplaceAll(tmp, "{{Identifier}}", note.GongGetIdentifier(stage))
			tmp = strings.ReplaceAll(tmp, "{{GeneratedFieldName}}", "Products")
			tmp = strings.ReplaceAll(tmp, "{{GeneratedFieldNameValue}}", _product.GongGetIdentifier(stage))
			res += tmp
		}
	case "Tasks":
		for _, _task := range note.Tasks {
			tmp := SliceOfPointersFieldInitStatement
			tmp = strings.ReplaceAll(tmp, "{{Identifier}}", note.GongGetIdentifier(stage))
			tmp = strings.ReplaceAll(tmp, "{{GeneratedFieldName}}", "Tasks")
			tmp = strings.ReplaceAll(tmp, "{{GeneratedFieldNameValue}}", _task.GongGetIdentifier(stage))
			res += tmp
		}
	default:
		log.Panicf("Unknown field %s for Gongstruct Note", fieldName)
	}
	return
}

func (noteproductshape *NoteProductShape) GongMarshallField(stage *Stage, fieldName string) (res string) {

	switch fieldName {
	case "Name":
		res = StringInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", noteproductshape.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "Name")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", string(noteproductshape.Name))
	case "StartRatio":
		res = NumberInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", noteproductshape.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "StartRatio")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%f", noteproductshape.StartRatio))
	case "EndRatio":
		res = NumberInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", noteproductshape.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "EndRatio")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%f", noteproductshape.EndRatio))
	case "StartOrientation":
		if noteproductshape.StartOrientation != "" {
			res = StringEnumInitStatement
			res = strings.ReplaceAll(res, "{{Identifier}}", noteproductshape.GongGetIdentifier(stage))
			res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "StartOrientation")
			res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", "models."+noteproductshape.StartOrientation.ToCodeString())
		}
	case "EndOrientation":
		if noteproductshape.EndOrientation != "" {
			res = StringEnumInitStatement
			res = strings.ReplaceAll(res, "{{Identifier}}", noteproductshape.GongGetIdentifier(stage))
			res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "EndOrientation")
			res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", "models."+noteproductshape.EndOrientation.ToCodeString())
		}
	case "CornerOffsetRatio":
		res = NumberInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", noteproductshape.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "CornerOffsetRatio")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%f", noteproductshape.CornerOffsetRatio))

	case "Note":
		if noteproductshape.Note != nil {
			res = PointerFieldInitStatement
			res = strings.ReplaceAll(res, "{{Identifier}}", noteproductshape.GongGetIdentifier(stage))
			res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "Note")
			res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", noteproductshape.Note.GongGetIdentifier(stage))
		}
	case "Product":
		if noteproductshape.Product != nil {
			res = PointerFieldInitStatement
			res = strings.ReplaceAll(res, "{{Identifier}}", noteproductshape.GongGetIdentifier(stage))
			res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "Product")
			res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", noteproductshape.Product.GongGetIdentifier(stage))
		}
	default:
		log.Panicf("Unknown field %s for Gongstruct NoteProductShape", fieldName)
	}
	return
}

func (noteshape *NoteShape) GongMarshallField(stage *Stage, fieldName string) (res string) {

	switch fieldName {
	case "Name":
		res = StringInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", noteshape.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "Name")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", string(noteshape.Name))
	case "IsExpanded":
		res = NumberInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", noteshape.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "IsExpanded")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%t", noteshape.IsExpanded))
	case "X":
		res = NumberInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", noteshape.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "X")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%f", noteshape.X))
	case "Y":
		res = NumberInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", noteshape.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "Y")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%f", noteshape.Y))
	case "Width":
		res = NumberInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", noteshape.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "Width")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%f", noteshape.Width))
	case "Height":
		res = NumberInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", noteshape.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "Height")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%f", noteshape.Height))

	case "Note":
		if noteshape.Note != nil {
			res = PointerFieldInitStatement
			res = strings.ReplaceAll(res, "{{Identifier}}", noteshape.GongGetIdentifier(stage))
			res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "Note")
			res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", noteshape.Note.GongGetIdentifier(stage))
		}
	default:
		log.Panicf("Unknown field %s for Gongstruct NoteShape", fieldName)
	}
	return
}

func (notetaskshape *NoteTaskShape) GongMarshallField(stage *Stage, fieldName string) (res string) {

	switch fieldName {
	case "Name":
		res = StringInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", notetaskshape.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "Name")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", string(notetaskshape.Name))
	case "StartRatio":
		res = NumberInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", notetaskshape.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "StartRatio")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%f", notetaskshape.StartRatio))
	case "EndRatio":
		res = NumberInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", notetaskshape.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "EndRatio")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%f", notetaskshape.EndRatio))
	case "StartOrientation":
		if notetaskshape.StartOrientation != "" {
			res = StringEnumInitStatement
			res = strings.ReplaceAll(res, "{{Identifier}}", notetaskshape.GongGetIdentifier(stage))
			res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "StartOrientation")
			res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", "models."+notetaskshape.StartOrientation.ToCodeString())
		}
	case "EndOrientation":
		if notetaskshape.EndOrientation != "" {
			res = StringEnumInitStatement
			res = strings.ReplaceAll(res, "{{Identifier}}", notetaskshape.GongGetIdentifier(stage))
			res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "EndOrientation")
			res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", "models."+notetaskshape.EndOrientation.ToCodeString())
		}
	case "CornerOffsetRatio":
		res = NumberInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", notetaskshape.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "CornerOffsetRatio")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%f", notetaskshape.CornerOffsetRatio))

	case "Note":
		if notetaskshape.Note != nil {
			res = PointerFieldInitStatement
			res = strings.ReplaceAll(res, "{{Identifier}}", notetaskshape.GongGetIdentifier(stage))
			res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "Note")
			res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", notetaskshape.Note.GongGetIdentifier(stage))
		}
	case "Task":
		if notetaskshape.Task != nil {
			res = PointerFieldInitStatement
			res = strings.ReplaceAll(res, "{{Identifier}}", notetaskshape.GongGetIdentifier(stage))
			res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "Task")
			res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", notetaskshape.Task.GongGetIdentifier(stage))
		}
	default:
		log.Panicf("Unknown field %s for Gongstruct NoteTaskShape", fieldName)
	}
	return
}

func (product *Product) GongMarshallField(stage *Stage, fieldName string) (res string) {

	switch fieldName {
	case "Name":
		res = StringInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", product.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "Name")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", string(product.Name))
	case "Description":
		res = StringInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", product.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "Description")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", string(product.Description))
	case "IsExpanded":
		res = NumberInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", product.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "IsExpanded")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%t", product.IsExpanded))
	case "ComputedPrefix":
		res = StringInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", product.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "ComputedPrefix")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", string(product.ComputedPrefix))
	case "IsProducersNodeExpanded":
		res = NumberInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", product.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "IsProducersNodeExpanded")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%t", product.IsProducersNodeExpanded))
	case "IsConsumersNodeExpanded":
		res = NumberInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", product.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "IsConsumersNodeExpanded")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%t", product.IsConsumersNodeExpanded))

	case "SubProducts":
		for _, _product := range product.SubProducts {
			tmp := SliceOfPointersFieldInitStatement
			tmp = strings.ReplaceAll(tmp, "{{Identifier}}", product.GongGetIdentifier(stage))
			tmp = strings.ReplaceAll(tmp, "{{GeneratedFieldName}}", "SubProducts")
			tmp = strings.ReplaceAll(tmp, "{{GeneratedFieldNameValue}}", _product.GongGetIdentifier(stage))
			res += tmp
		}
	default:
		log.Panicf("Unknown field %s for Gongstruct Product", fieldName)
	}
	return
}

func (productcompositionshape *ProductCompositionShape) GongMarshallField(stage *Stage, fieldName string) (res string) {

	switch fieldName {
	case "Name":
		res = StringInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", productcompositionshape.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "Name")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", string(productcompositionshape.Name))
	case "StartRatio":
		res = NumberInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", productcompositionshape.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "StartRatio")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%f", productcompositionshape.StartRatio))
	case "EndRatio":
		res = NumberInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", productcompositionshape.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "EndRatio")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%f", productcompositionshape.EndRatio))
	case "StartOrientation":
		if productcompositionshape.StartOrientation != "" {
			res = StringEnumInitStatement
			res = strings.ReplaceAll(res, "{{Identifier}}", productcompositionshape.GongGetIdentifier(stage))
			res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "StartOrientation")
			res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", "models."+productcompositionshape.StartOrientation.ToCodeString())
		}
	case "EndOrientation":
		if productcompositionshape.EndOrientation != "" {
			res = StringEnumInitStatement
			res = strings.ReplaceAll(res, "{{Identifier}}", productcompositionshape.GongGetIdentifier(stage))
			res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "EndOrientation")
			res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", "models."+productcompositionshape.EndOrientation.ToCodeString())
		}
	case "CornerOffsetRatio":
		res = NumberInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", productcompositionshape.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "CornerOffsetRatio")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%f", productcompositionshape.CornerOffsetRatio))

	case "Product":
		if productcompositionshape.Product != nil {
			res = PointerFieldInitStatement
			res = strings.ReplaceAll(res, "{{Identifier}}", productcompositionshape.GongGetIdentifier(stage))
			res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "Product")
			res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", productcompositionshape.Product.GongGetIdentifier(stage))
		}
	default:
		log.Panicf("Unknown field %s for Gongstruct ProductCompositionShape", fieldName)
	}
	return
}

func (productshape *ProductShape) GongMarshallField(stage *Stage, fieldName string) (res string) {

	switch fieldName {
	case "Name":
		res = StringInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", productshape.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "Name")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", string(productshape.Name))
	case "IsExpanded":
		res = NumberInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", productshape.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "IsExpanded")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%t", productshape.IsExpanded))
	case "X":
		res = NumberInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", productshape.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "X")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%f", productshape.X))
	case "Y":
		res = NumberInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", productshape.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "Y")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%f", productshape.Y))
	case "Width":
		res = NumberInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", productshape.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "Width")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%f", productshape.Width))
	case "Height":
		res = NumberInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", productshape.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "Height")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%f", productshape.Height))

	case "Product":
		if productshape.Product != nil {
			res = PointerFieldInitStatement
			res = strings.ReplaceAll(res, "{{Identifier}}", productshape.GongGetIdentifier(stage))
			res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "Product")
			res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", productshape.Product.GongGetIdentifier(stage))
		}
	default:
		log.Panicf("Unknown field %s for Gongstruct ProductShape", fieldName)
	}
	return
}

func (project *Project) GongMarshallField(stage *Stage, fieldName string) (res string) {

	switch fieldName {
	case "Name":
		res = StringInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", project.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "Name")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", string(project.Name))
	case "IsPBSNodeExpanded":
		res = NumberInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", project.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "IsPBSNodeExpanded")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%t", project.IsPBSNodeExpanded))
	case "IsWBSNodeExpanded":
		res = NumberInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", project.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "IsWBSNodeExpanded")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%t", project.IsWBSNodeExpanded))
	case "IsDiagramsNodeExpanded":
		res = NumberInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", project.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "IsDiagramsNodeExpanded")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%t", project.IsDiagramsNodeExpanded))
	case "IsNotesNodeExpanded":
		res = NumberInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", project.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "IsNotesNodeExpanded")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%t", project.IsNotesNodeExpanded))
	case "IsExpanded":
		res = NumberInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", project.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "IsExpanded")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%t", project.IsExpanded))
	case "ComputedPrefix":
		res = StringInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", project.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "ComputedPrefix")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", string(project.ComputedPrefix))

	case "RootProducts":
		for _, _product := range project.RootProducts {
			tmp := SliceOfPointersFieldInitStatement
			tmp = strings.ReplaceAll(tmp, "{{Identifier}}", project.GongGetIdentifier(stage))
			tmp = strings.ReplaceAll(tmp, "{{GeneratedFieldName}}", "RootProducts")
			tmp = strings.ReplaceAll(tmp, "{{GeneratedFieldNameValue}}", _product.GongGetIdentifier(stage))
			res += tmp
		}
	case "RootTasks":
		for _, _task := range project.RootTasks {
			tmp := SliceOfPointersFieldInitStatement
			tmp = strings.ReplaceAll(tmp, "{{Identifier}}", project.GongGetIdentifier(stage))
			tmp = strings.ReplaceAll(tmp, "{{GeneratedFieldName}}", "RootTasks")
			tmp = strings.ReplaceAll(tmp, "{{GeneratedFieldNameValue}}", _task.GongGetIdentifier(stage))
			res += tmp
		}
	case "Diagrams":
		for _, _diagram := range project.Diagrams {
			tmp := SliceOfPointersFieldInitStatement
			tmp = strings.ReplaceAll(tmp, "{{Identifier}}", project.GongGetIdentifier(stage))
			tmp = strings.ReplaceAll(tmp, "{{GeneratedFieldName}}", "Diagrams")
			tmp = strings.ReplaceAll(tmp, "{{GeneratedFieldNameValue}}", _diagram.GongGetIdentifier(stage))
			res += tmp
		}
	case "Notes":
		for _, _note := range project.Notes {
			tmp := SliceOfPointersFieldInitStatement
			tmp = strings.ReplaceAll(tmp, "{{Identifier}}", project.GongGetIdentifier(stage))
			tmp = strings.ReplaceAll(tmp, "{{GeneratedFieldName}}", "Notes")
			tmp = strings.ReplaceAll(tmp, "{{GeneratedFieldNameValue}}", _note.GongGetIdentifier(stage))
			res += tmp
		}
	default:
		log.Panicf("Unknown field %s for Gongstruct Project", fieldName)
	}
	return
}

func (root *Root) GongMarshallField(stage *Stage, fieldName string) (res string) {

	switch fieldName {
	case "Name":
		res = StringInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", root.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "Name")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", string(root.Name))
	case "NbPixPerCharacter":
		res = NumberInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", root.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "NbPixPerCharacter")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%f", root.NbPixPerCharacter))

	case "Projects":
		for _, _project := range root.Projects {
			tmp := SliceOfPointersFieldInitStatement
			tmp = strings.ReplaceAll(tmp, "{{Identifier}}", root.GongGetIdentifier(stage))
			tmp = strings.ReplaceAll(tmp, "{{GeneratedFieldName}}", "Projects")
			tmp = strings.ReplaceAll(tmp, "{{GeneratedFieldNameValue}}", _project.GongGetIdentifier(stage))
			res += tmp
		}
	case "OrphanedProducts":
		for _, _product := range root.OrphanedProducts {
			tmp := SliceOfPointersFieldInitStatement
			tmp = strings.ReplaceAll(tmp, "{{Identifier}}", root.GongGetIdentifier(stage))
			tmp = strings.ReplaceAll(tmp, "{{GeneratedFieldName}}", "OrphanedProducts")
			tmp = strings.ReplaceAll(tmp, "{{GeneratedFieldNameValue}}", _product.GongGetIdentifier(stage))
			res += tmp
		}
	case "OrphanedTasks":
		for _, _task := range root.OrphanedTasks {
			tmp := SliceOfPointersFieldInitStatement
			tmp = strings.ReplaceAll(tmp, "{{Identifier}}", root.GongGetIdentifier(stage))
			tmp = strings.ReplaceAll(tmp, "{{GeneratedFieldName}}", "OrphanedTasks")
			tmp = strings.ReplaceAll(tmp, "{{GeneratedFieldNameValue}}", _task.GongGetIdentifier(stage))
			res += tmp
		}
	default:
		log.Panicf("Unknown field %s for Gongstruct Root", fieldName)
	}
	return
}

func (task *Task) GongMarshallField(stage *Stage, fieldName string) (res string) {

	switch fieldName {
	case "Name":
		res = StringInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", task.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "Name")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", string(task.Name))
	case "Description":
		res = StringInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", task.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "Description")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", string(task.Description))
	case "IsExpanded":
		res = NumberInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", task.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "IsExpanded")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%t", task.IsExpanded))
	case "ComputedPrefix":
		res = StringInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", task.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "ComputedPrefix")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", string(task.ComputedPrefix))
	case "IsInputsNodeExpanded":
		res = NumberInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", task.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "IsInputsNodeExpanded")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%t", task.IsInputsNodeExpanded))
	case "IsOutputsNodeExpanded":
		res = NumberInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", task.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "IsOutputsNodeExpanded")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%t", task.IsOutputsNodeExpanded))

	case "SubTasks":
		for _, _task := range task.SubTasks {
			tmp := SliceOfPointersFieldInitStatement
			tmp = strings.ReplaceAll(tmp, "{{Identifier}}", task.GongGetIdentifier(stage))
			tmp = strings.ReplaceAll(tmp, "{{GeneratedFieldName}}", "SubTasks")
			tmp = strings.ReplaceAll(tmp, "{{GeneratedFieldNameValue}}", _task.GongGetIdentifier(stage))
			res += tmp
		}
	case "Inputs":
		for _, _product := range task.Inputs {
			tmp := SliceOfPointersFieldInitStatement
			tmp = strings.ReplaceAll(tmp, "{{Identifier}}", task.GongGetIdentifier(stage))
			tmp = strings.ReplaceAll(tmp, "{{GeneratedFieldName}}", "Inputs")
			tmp = strings.ReplaceAll(tmp, "{{GeneratedFieldNameValue}}", _product.GongGetIdentifier(stage))
			res += tmp
		}
	case "Outputs":
		for _, _product := range task.Outputs {
			tmp := SliceOfPointersFieldInitStatement
			tmp = strings.ReplaceAll(tmp, "{{Identifier}}", task.GongGetIdentifier(stage))
			tmp = strings.ReplaceAll(tmp, "{{GeneratedFieldName}}", "Outputs")
			tmp = strings.ReplaceAll(tmp, "{{GeneratedFieldNameValue}}", _product.GongGetIdentifier(stage))
			res += tmp
		}
	default:
		log.Panicf("Unknown field %s for Gongstruct Task", fieldName)
	}
	return
}

func (taskcompositionshape *TaskCompositionShape) GongMarshallField(stage *Stage, fieldName string) (res string) {

	switch fieldName {
	case "Name":
		res = StringInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", taskcompositionshape.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "Name")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", string(taskcompositionshape.Name))
	case "StartRatio":
		res = NumberInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", taskcompositionshape.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "StartRatio")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%f", taskcompositionshape.StartRatio))
	case "EndRatio":
		res = NumberInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", taskcompositionshape.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "EndRatio")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%f", taskcompositionshape.EndRatio))
	case "StartOrientation":
		if taskcompositionshape.StartOrientation != "" {
			res = StringEnumInitStatement
			res = strings.ReplaceAll(res, "{{Identifier}}", taskcompositionshape.GongGetIdentifier(stage))
			res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "StartOrientation")
			res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", "models."+taskcompositionshape.StartOrientation.ToCodeString())
		}
	case "EndOrientation":
		if taskcompositionshape.EndOrientation != "" {
			res = StringEnumInitStatement
			res = strings.ReplaceAll(res, "{{Identifier}}", taskcompositionshape.GongGetIdentifier(stage))
			res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "EndOrientation")
			res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", "models."+taskcompositionshape.EndOrientation.ToCodeString())
		}
	case "CornerOffsetRatio":
		res = NumberInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", taskcompositionshape.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "CornerOffsetRatio")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%f", taskcompositionshape.CornerOffsetRatio))

	case "Task":
		if taskcompositionshape.Task != nil {
			res = PointerFieldInitStatement
			res = strings.ReplaceAll(res, "{{Identifier}}", taskcompositionshape.GongGetIdentifier(stage))
			res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "Task")
			res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", taskcompositionshape.Task.GongGetIdentifier(stage))
		}
	default:
		log.Panicf("Unknown field %s for Gongstruct TaskCompositionShape", fieldName)
	}
	return
}

func (taskinputshape *TaskInputShape) GongMarshallField(stage *Stage, fieldName string) (res string) {

	switch fieldName {
	case "Name":
		res = StringInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", taskinputshape.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "Name")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", string(taskinputshape.Name))
	case "StartRatio":
		res = NumberInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", taskinputshape.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "StartRatio")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%f", taskinputshape.StartRatio))
	case "EndRatio":
		res = NumberInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", taskinputshape.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "EndRatio")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%f", taskinputshape.EndRatio))
	case "StartOrientation":
		if taskinputshape.StartOrientation != "" {
			res = StringEnumInitStatement
			res = strings.ReplaceAll(res, "{{Identifier}}", taskinputshape.GongGetIdentifier(stage))
			res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "StartOrientation")
			res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", "models."+taskinputshape.StartOrientation.ToCodeString())
		}
	case "EndOrientation":
		if taskinputshape.EndOrientation != "" {
			res = StringEnumInitStatement
			res = strings.ReplaceAll(res, "{{Identifier}}", taskinputshape.GongGetIdentifier(stage))
			res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "EndOrientation")
			res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", "models."+taskinputshape.EndOrientation.ToCodeString())
		}
	case "CornerOffsetRatio":
		res = NumberInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", taskinputshape.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "CornerOffsetRatio")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%f", taskinputshape.CornerOffsetRatio))

	case "Task":
		if taskinputshape.Task != nil {
			res = PointerFieldInitStatement
			res = strings.ReplaceAll(res, "{{Identifier}}", taskinputshape.GongGetIdentifier(stage))
			res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "Task")
			res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", taskinputshape.Task.GongGetIdentifier(stage))
		}
	case "Product":
		if taskinputshape.Product != nil {
			res = PointerFieldInitStatement
			res = strings.ReplaceAll(res, "{{Identifier}}", taskinputshape.GongGetIdentifier(stage))
			res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "Product")
			res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", taskinputshape.Product.GongGetIdentifier(stage))
		}
	default:
		log.Panicf("Unknown field %s for Gongstruct TaskInputShape", fieldName)
	}
	return
}

func (taskoutputshape *TaskOutputShape) GongMarshallField(stage *Stage, fieldName string) (res string) {

	switch fieldName {
	case "Name":
		res = StringInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", taskoutputshape.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "Name")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", string(taskoutputshape.Name))
	case "StartRatio":
		res = NumberInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", taskoutputshape.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "StartRatio")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%f", taskoutputshape.StartRatio))
	case "EndRatio":
		res = NumberInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", taskoutputshape.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "EndRatio")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%f", taskoutputshape.EndRatio))
	case "StartOrientation":
		if taskoutputshape.StartOrientation != "" {
			res = StringEnumInitStatement
			res = strings.ReplaceAll(res, "{{Identifier}}", taskoutputshape.GongGetIdentifier(stage))
			res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "StartOrientation")
			res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", "models."+taskoutputshape.StartOrientation.ToCodeString())
		}
	case "EndOrientation":
		if taskoutputshape.EndOrientation != "" {
			res = StringEnumInitStatement
			res = strings.ReplaceAll(res, "{{Identifier}}", taskoutputshape.GongGetIdentifier(stage))
			res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "EndOrientation")
			res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", "models."+taskoutputshape.EndOrientation.ToCodeString())
		}
	case "CornerOffsetRatio":
		res = NumberInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", taskoutputshape.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "CornerOffsetRatio")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%f", taskoutputshape.CornerOffsetRatio))

	case "Task":
		if taskoutputshape.Task != nil {
			res = PointerFieldInitStatement
			res = strings.ReplaceAll(res, "{{Identifier}}", taskoutputshape.GongGetIdentifier(stage))
			res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "Task")
			res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", taskoutputshape.Task.GongGetIdentifier(stage))
		}
	case "Product":
		if taskoutputshape.Product != nil {
			res = PointerFieldInitStatement
			res = strings.ReplaceAll(res, "{{Identifier}}", taskoutputshape.GongGetIdentifier(stage))
			res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "Product")
			res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", taskoutputshape.Product.GongGetIdentifier(stage))
		}
	default:
		log.Panicf("Unknown field %s for Gongstruct TaskOutputShape", fieldName)
	}
	return
}

func (taskshape *TaskShape) GongMarshallField(stage *Stage, fieldName string) (res string) {

	switch fieldName {
	case "Name":
		res = StringInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", taskshape.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "Name")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", string(taskshape.Name))
	case "IsExpanded":
		res = NumberInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", taskshape.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "IsExpanded")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%t", taskshape.IsExpanded))
	case "X":
		res = NumberInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", taskshape.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "X")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%f", taskshape.X))
	case "Y":
		res = NumberInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", taskshape.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "Y")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%f", taskshape.Y))
	case "Width":
		res = NumberInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", taskshape.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "Width")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%f", taskshape.Width))
	case "Height":
		res = NumberInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", taskshape.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "Height")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%f", taskshape.Height))

	case "Task":
		if taskshape.Task != nil {
			res = PointerFieldInitStatement
			res = strings.ReplaceAll(res, "{{Identifier}}", taskshape.GongGetIdentifier(stage))
			res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "Task")
			res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", taskshape.Task.GongGetIdentifier(stage))
		}
	default:
		log.Panicf("Unknown field %s for Gongstruct TaskShape", fieldName)
	}
	return
}

// insertion point for marshall all fields methods
func (diagram *Diagram) GongMarshallAllFields(stage *Stage) (initializerStatements string, pointersInitializesStatements string) {

	initializerStatements += "\n"
	pointersInitializesStatements += "\n"
	{ // Insertion point for basic fields value assignment
		initializerStatements += diagram.GongMarshallField(stage, "Name")
		initializerStatements += diagram.GongMarshallField(stage, "IsChecked")
		initializerStatements += diagram.GongMarshallField(stage, "IsEditable_")
		initializerStatements += diagram.GongMarshallField(stage, "IsInRenameMode")
		initializerStatements += diagram.GongMarshallField(stage, "ShowPrefix")
		initializerStatements += diagram.GongMarshallField(stage, "DefaultBoxWidth")
		initializerStatements += diagram.GongMarshallField(stage, "DefaultBoxHeigth")
		initializerStatements += diagram.GongMarshallField(stage, "IsExpanded")
		initializerStatements += diagram.GongMarshallField(stage, "ComputedPrefix")
		pointersInitializesStatements += diagram.GongMarshallField(stage, "Product_Shapes")
		pointersInitializesStatements += diagram.GongMarshallField(stage, "ProductsWhoseNodeIsExpanded")
		initializerStatements += diagram.GongMarshallField(stage, "IsPBSNodeExpanded")
		pointersInitializesStatements += diagram.GongMarshallField(stage, "ProductComposition_Shapes")
		initializerStatements += diagram.GongMarshallField(stage, "IsWBSNodeExpanded")
		pointersInitializesStatements += diagram.GongMarshallField(stage, "Task_Shapes")
		pointersInitializesStatements += diagram.GongMarshallField(stage, "TasksWhoseNodeIsExpanded")
		pointersInitializesStatements += diagram.GongMarshallField(stage, "TasksWhoseInputNodeIsExpanded")
		pointersInitializesStatements += diagram.GongMarshallField(stage, "TasksWhoseOutputNodeIsExpanded")
		pointersInitializesStatements += diagram.GongMarshallField(stage, "TaskComposition_Shapes")
		pointersInitializesStatements += diagram.GongMarshallField(stage, "TaskInputShapes")
		pointersInitializesStatements += diagram.GongMarshallField(stage, "TaskOutputShapes")
		pointersInitializesStatements += diagram.GongMarshallField(stage, "Note_Shapes")
		pointersInitializesStatements += diagram.GongMarshallField(stage, "NotesWhoseNodeIsExpanded")
		initializerStatements += diagram.GongMarshallField(stage, "IsNotesNodeExpanded")
		pointersInitializesStatements += diagram.GongMarshallField(stage, "NoteProductShapes")
		pointersInitializesStatements += diagram.GongMarshallField(stage, "NoteTaskShapes")
	}
	return
}
func (note *Note) GongMarshallAllFields(stage *Stage) (initializerStatements string, pointersInitializesStatements string) {

	initializerStatements += "\n"
	pointersInitializesStatements += "\n"
	{ // Insertion point for basic fields value assignment
		initializerStatements += note.GongMarshallField(stage, "Name")
		pointersInitializesStatements += note.GongMarshallField(stage, "Products")
		pointersInitializesStatements += note.GongMarshallField(stage, "Tasks")
		initializerStatements += note.GongMarshallField(stage, "IsExpanded")
	}
	return
}
func (noteproductshape *NoteProductShape) GongMarshallAllFields(stage *Stage) (initializerStatements string, pointersInitializesStatements string) {

	initializerStatements += "\n"
	pointersInitializesStatements += "\n"
	{ // Insertion point for basic fields value assignment
		initializerStatements += noteproductshape.GongMarshallField(stage, "Name")
		pointersInitializesStatements += noteproductshape.GongMarshallField(stage, "Note")
		pointersInitializesStatements += noteproductshape.GongMarshallField(stage, "Product")
		initializerStatements += noteproductshape.GongMarshallField(stage, "StartRatio")
		initializerStatements += noteproductshape.GongMarshallField(stage, "EndRatio")
		initializerStatements += noteproductshape.GongMarshallField(stage, "StartOrientation")
		initializerStatements += noteproductshape.GongMarshallField(stage, "EndOrientation")
		initializerStatements += noteproductshape.GongMarshallField(stage, "CornerOffsetRatio")
	}
	return
}
func (noteshape *NoteShape) GongMarshallAllFields(stage *Stage) (initializerStatements string, pointersInitializesStatements string) {

	initializerStatements += "\n"
	pointersInitializesStatements += "\n"
	{ // Insertion point for basic fields value assignment
		initializerStatements += noteshape.GongMarshallField(stage, "Name")
		pointersInitializesStatements += noteshape.GongMarshallField(stage, "Note")
		initializerStatements += noteshape.GongMarshallField(stage, "IsExpanded")
		initializerStatements += noteshape.GongMarshallField(stage, "X")
		initializerStatements += noteshape.GongMarshallField(stage, "Y")
		initializerStatements += noteshape.GongMarshallField(stage, "Width")
		initializerStatements += noteshape.GongMarshallField(stage, "Height")
	}
	return
}
func (notetaskshape *NoteTaskShape) GongMarshallAllFields(stage *Stage) (initializerStatements string, pointersInitializesStatements string) {

	initializerStatements += "\n"
	pointersInitializesStatements += "\n"
	{ // Insertion point for basic fields value assignment
		initializerStatements += notetaskshape.GongMarshallField(stage, "Name")
		pointersInitializesStatements += notetaskshape.GongMarshallField(stage, "Note")
		pointersInitializesStatements += notetaskshape.GongMarshallField(stage, "Task")
		initializerStatements += notetaskshape.GongMarshallField(stage, "StartRatio")
		initializerStatements += notetaskshape.GongMarshallField(stage, "EndRatio")
		initializerStatements += notetaskshape.GongMarshallField(stage, "StartOrientation")
		initializerStatements += notetaskshape.GongMarshallField(stage, "EndOrientation")
		initializerStatements += notetaskshape.GongMarshallField(stage, "CornerOffsetRatio")
	}
	return
}
func (product *Product) GongMarshallAllFields(stage *Stage) (initializerStatements string, pointersInitializesStatements string) {

	initializerStatements += "\n"
	pointersInitializesStatements += "\n"
	{ // Insertion point for basic fields value assignment
		initializerStatements += product.GongMarshallField(stage, "Name")
		initializerStatements += product.GongMarshallField(stage, "Description")
		pointersInitializesStatements += product.GongMarshallField(stage, "SubProducts")
		initializerStatements += product.GongMarshallField(stage, "IsExpanded")
		initializerStatements += product.GongMarshallField(stage, "ComputedPrefix")
		initializerStatements += product.GongMarshallField(stage, "IsProducersNodeExpanded")
		initializerStatements += product.GongMarshallField(stage, "IsConsumersNodeExpanded")
	}
	return
}
func (productcompositionshape *ProductCompositionShape) GongMarshallAllFields(stage *Stage) (initializerStatements string, pointersInitializesStatements string) {

	initializerStatements += "\n"
	pointersInitializesStatements += "\n"
	{ // Insertion point for basic fields value assignment
		initializerStatements += productcompositionshape.GongMarshallField(stage, "Name")
		pointersInitializesStatements += productcompositionshape.GongMarshallField(stage, "Product")
		initializerStatements += productcompositionshape.GongMarshallField(stage, "StartRatio")
		initializerStatements += productcompositionshape.GongMarshallField(stage, "EndRatio")
		initializerStatements += productcompositionshape.GongMarshallField(stage, "StartOrientation")
		initializerStatements += productcompositionshape.GongMarshallField(stage, "EndOrientation")
		initializerStatements += productcompositionshape.GongMarshallField(stage, "CornerOffsetRatio")
	}
	return
}
func (productshape *ProductShape) GongMarshallAllFields(stage *Stage) (initializerStatements string, pointersInitializesStatements string) {

	initializerStatements += "\n"
	pointersInitializesStatements += "\n"
	{ // Insertion point for basic fields value assignment
		initializerStatements += productshape.GongMarshallField(stage, "Name")
		pointersInitializesStatements += productshape.GongMarshallField(stage, "Product")
		initializerStatements += productshape.GongMarshallField(stage, "IsExpanded")
		initializerStatements += productshape.GongMarshallField(stage, "X")
		initializerStatements += productshape.GongMarshallField(stage, "Y")
		initializerStatements += productshape.GongMarshallField(stage, "Width")
		initializerStatements += productshape.GongMarshallField(stage, "Height")
	}
	return
}
func (project *Project) GongMarshallAllFields(stage *Stage) (initializerStatements string, pointersInitializesStatements string) {

	initializerStatements += "\n"
	pointersInitializesStatements += "\n"
	{ // Insertion point for basic fields value assignment
		initializerStatements += project.GongMarshallField(stage, "Name")
		pointersInitializesStatements += project.GongMarshallField(stage, "RootProducts")
		initializerStatements += project.GongMarshallField(stage, "IsPBSNodeExpanded")
		pointersInitializesStatements += project.GongMarshallField(stage, "RootTasks")
		initializerStatements += project.GongMarshallField(stage, "IsWBSNodeExpanded")
		pointersInitializesStatements += project.GongMarshallField(stage, "Diagrams")
		initializerStatements += project.GongMarshallField(stage, "IsDiagramsNodeExpanded")
		pointersInitializesStatements += project.GongMarshallField(stage, "Notes")
		initializerStatements += project.GongMarshallField(stage, "IsNotesNodeExpanded")
		initializerStatements += project.GongMarshallField(stage, "IsExpanded")
		initializerStatements += project.GongMarshallField(stage, "ComputedPrefix")
	}
	return
}
func (root *Root) GongMarshallAllFields(stage *Stage) (initializerStatements string, pointersInitializesStatements string) {

	initializerStatements += "\n"
	pointersInitializesStatements += "\n"
	{ // Insertion point for basic fields value assignment
		initializerStatements += root.GongMarshallField(stage, "Name")
		pointersInitializesStatements += root.GongMarshallField(stage, "Projects")
		pointersInitializesStatements += root.GongMarshallField(stage, "OrphanedProducts")
		pointersInitializesStatements += root.GongMarshallField(stage, "OrphanedTasks")
		initializerStatements += root.GongMarshallField(stage, "NbPixPerCharacter")
	}
	return
}
func (task *Task) GongMarshallAllFields(stage *Stage) (initializerStatements string, pointersInitializesStatements string) {

	initializerStatements += "\n"
	pointersInitializesStatements += "\n"
	{ // Insertion point for basic fields value assignment
		initializerStatements += task.GongMarshallField(stage, "Name")
		initializerStatements += task.GongMarshallField(stage, "Description")
		pointersInitializesStatements += task.GongMarshallField(stage, "SubTasks")
		initializerStatements += task.GongMarshallField(stage, "IsExpanded")
		initializerStatements += task.GongMarshallField(stage, "ComputedPrefix")
		pointersInitializesStatements += task.GongMarshallField(stage, "Inputs")
		initializerStatements += task.GongMarshallField(stage, "IsInputsNodeExpanded")
		pointersInitializesStatements += task.GongMarshallField(stage, "Outputs")
		initializerStatements += task.GongMarshallField(stage, "IsOutputsNodeExpanded")
	}
	return
}
func (taskcompositionshape *TaskCompositionShape) GongMarshallAllFields(stage *Stage) (initializerStatements string, pointersInitializesStatements string) {

	initializerStatements += "\n"
	pointersInitializesStatements += "\n"
	{ // Insertion point for basic fields value assignment
		initializerStatements += taskcompositionshape.GongMarshallField(stage, "Name")
		pointersInitializesStatements += taskcompositionshape.GongMarshallField(stage, "Task")
		initializerStatements += taskcompositionshape.GongMarshallField(stage, "StartRatio")
		initializerStatements += taskcompositionshape.GongMarshallField(stage, "EndRatio")
		initializerStatements += taskcompositionshape.GongMarshallField(stage, "StartOrientation")
		initializerStatements += taskcompositionshape.GongMarshallField(stage, "EndOrientation")
		initializerStatements += taskcompositionshape.GongMarshallField(stage, "CornerOffsetRatio")
	}
	return
}
func (taskinputshape *TaskInputShape) GongMarshallAllFields(stage *Stage) (initializerStatements string, pointersInitializesStatements string) {

	initializerStatements += "\n"
	pointersInitializesStatements += "\n"
	{ // Insertion point for basic fields value assignment
		initializerStatements += taskinputshape.GongMarshallField(stage, "Name")
		pointersInitializesStatements += taskinputshape.GongMarshallField(stage, "Task")
		pointersInitializesStatements += taskinputshape.GongMarshallField(stage, "Product")
		initializerStatements += taskinputshape.GongMarshallField(stage, "StartRatio")
		initializerStatements += taskinputshape.GongMarshallField(stage, "EndRatio")
		initializerStatements += taskinputshape.GongMarshallField(stage, "StartOrientation")
		initializerStatements += taskinputshape.GongMarshallField(stage, "EndOrientation")
		initializerStatements += taskinputshape.GongMarshallField(stage, "CornerOffsetRatio")
	}
	return
}
func (taskoutputshape *TaskOutputShape) GongMarshallAllFields(stage *Stage) (initializerStatements string, pointersInitializesStatements string) {

	initializerStatements += "\n"
	pointersInitializesStatements += "\n"
	{ // Insertion point for basic fields value assignment
		initializerStatements += taskoutputshape.GongMarshallField(stage, "Name")
		pointersInitializesStatements += taskoutputshape.GongMarshallField(stage, "Task")
		pointersInitializesStatements += taskoutputshape.GongMarshallField(stage, "Product")
		initializerStatements += taskoutputshape.GongMarshallField(stage, "StartRatio")
		initializerStatements += taskoutputshape.GongMarshallField(stage, "EndRatio")
		initializerStatements += taskoutputshape.GongMarshallField(stage, "StartOrientation")
		initializerStatements += taskoutputshape.GongMarshallField(stage, "EndOrientation")
		initializerStatements += taskoutputshape.GongMarshallField(stage, "CornerOffsetRatio")
	}
	return
}
func (taskshape *TaskShape) GongMarshallAllFields(stage *Stage) (initializerStatements string, pointersInitializesStatements string) {

	initializerStatements += "\n"
	pointersInitializesStatements += "\n"
	{ // Insertion point for basic fields value assignment
		initializerStatements += taskshape.GongMarshallField(stage, "Name")
		pointersInitializesStatements += taskshape.GongMarshallField(stage, "Task")
		initializerStatements += taskshape.GongMarshallField(stage, "IsExpanded")
		initializerStatements += taskshape.GongMarshallField(stage, "X")
		initializerStatements += taskshape.GongMarshallField(stage, "Y")
		initializerStatements += taskshape.GongMarshallField(stage, "Width")
		initializerStatements += taskshape.GongMarshallField(stage, "Height")
	}
	return
}