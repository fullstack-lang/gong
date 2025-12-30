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
		// Initialisation of values
		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", diagram.GongGetIdentifier(stage))
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Name")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(diagram.Name))
		initializerStatements += setValueField

		setValueField = NumberInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", diagram.GongGetIdentifier(stage))
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "IsChecked")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%t", diagram.IsChecked))
		initializerStatements += setValueField

		setValueField = NumberInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", diagram.GongGetIdentifier(stage))
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "IsEditable_")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%t", diagram.IsEditable_))
		initializerStatements += setValueField

		setValueField = NumberInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", diagram.GongGetIdentifier(stage))
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "IsInRenameMode")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%t", diagram.IsInRenameMode))
		initializerStatements += setValueField

		setValueField = NumberInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", diagram.GongGetIdentifier(stage))
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "ShowPrefix")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%t", diagram.ShowPrefix))
		initializerStatements += setValueField

		setValueField = NumberInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", diagram.GongGetIdentifier(stage))
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "DefaultBoxWidth")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%f", diagram.DefaultBoxWidth))
		initializerStatements += setValueField

		setValueField = NumberInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", diagram.GongGetIdentifier(stage))
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "DefaultBoxHeigth")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%f", diagram.DefaultBoxHeigth))
		initializerStatements += setValueField

		setValueField = NumberInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", diagram.GongGetIdentifier(stage))
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "IsExpanded")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%t", diagram.IsExpanded))
		initializerStatements += setValueField

		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", diagram.GongGetIdentifier(stage))
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "ComputedPrefix")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(diagram.ComputedPrefix))
		initializerStatements += setValueField

		setValueField = NumberInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", diagram.GongGetIdentifier(stage))
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "IsPBSNodeExpanded")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%t", diagram.IsPBSNodeExpanded))
		initializerStatements += setValueField

		setValueField = NumberInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", diagram.GongGetIdentifier(stage))
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "IsWBSNodeExpanded")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%t", diagram.IsWBSNodeExpanded))
		initializerStatements += setValueField

		setValueField = NumberInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", diagram.GongGetIdentifier(stage))
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "IsNotesNodeExpanded")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%t", diagram.IsNotesNodeExpanded))
		initializerStatements += setValueField

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
		// Initialisation of values
		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", note.GongGetIdentifier(stage))
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Name")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(note.Name))
		initializerStatements += setValueField

		setValueField = NumberInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", note.GongGetIdentifier(stage))
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "IsExpanded")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%t", note.IsExpanded))
		initializerStatements += setValueField

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
		// Initialisation of values
		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", noteproductshape.GongGetIdentifier(stage))
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Name")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(noteproductshape.Name))
		initializerStatements += setValueField

		setValueField = NumberInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", noteproductshape.GongGetIdentifier(stage))
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "StartRatio")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%f", noteproductshape.StartRatio))
		initializerStatements += setValueField

		setValueField = NumberInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", noteproductshape.GongGetIdentifier(stage))
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "EndRatio")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%f", noteproductshape.EndRatio))
		initializerStatements += setValueField

		if noteproductshape.StartOrientation != "" {
			setValueField = StringEnumInitStatement
			setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", noteproductshape.GongGetIdentifier(stage))
			setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "StartOrientation")
			setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", "models."+noteproductshape.StartOrientation.ToCodeString())
			initializerStatements += setValueField
		}

		if noteproductshape.EndOrientation != "" {
			setValueField = StringEnumInitStatement
			setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", noteproductshape.GongGetIdentifier(stage))
			setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "EndOrientation")
			setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", "models."+noteproductshape.EndOrientation.ToCodeString())
			initializerStatements += setValueField
		}

		setValueField = NumberInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", noteproductshape.GongGetIdentifier(stage))
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "CornerOffsetRatio")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%f", noteproductshape.CornerOffsetRatio))
		initializerStatements += setValueField

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
		// Initialisation of values
		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", noteshape.GongGetIdentifier(stage))
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Name")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(noteshape.Name))
		initializerStatements += setValueField

		setValueField = NumberInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", noteshape.GongGetIdentifier(stage))
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "IsExpanded")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%t", noteshape.IsExpanded))
		initializerStatements += setValueField

		setValueField = NumberInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", noteshape.GongGetIdentifier(stage))
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "X")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%f", noteshape.X))
		initializerStatements += setValueField

		setValueField = NumberInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", noteshape.GongGetIdentifier(stage))
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Y")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%f", noteshape.Y))
		initializerStatements += setValueField

		setValueField = NumberInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", noteshape.GongGetIdentifier(stage))
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Width")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%f", noteshape.Width))
		initializerStatements += setValueField

		setValueField = NumberInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", noteshape.GongGetIdentifier(stage))
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Height")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%f", noteshape.Height))
		initializerStatements += setValueField

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
		// Initialisation of values
		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", notetaskshape.GongGetIdentifier(stage))
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Name")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(notetaskshape.Name))
		initializerStatements += setValueField

		setValueField = NumberInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", notetaskshape.GongGetIdentifier(stage))
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "StartRatio")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%f", notetaskshape.StartRatio))
		initializerStatements += setValueField

		setValueField = NumberInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", notetaskshape.GongGetIdentifier(stage))
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "EndRatio")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%f", notetaskshape.EndRatio))
		initializerStatements += setValueField

		if notetaskshape.StartOrientation != "" {
			setValueField = StringEnumInitStatement
			setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", notetaskshape.GongGetIdentifier(stage))
			setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "StartOrientation")
			setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", "models."+notetaskshape.StartOrientation.ToCodeString())
			initializerStatements += setValueField
		}

		if notetaskshape.EndOrientation != "" {
			setValueField = StringEnumInitStatement
			setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", notetaskshape.GongGetIdentifier(stage))
			setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "EndOrientation")
			setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", "models."+notetaskshape.EndOrientation.ToCodeString())
			initializerStatements += setValueField
		}

		setValueField = NumberInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", notetaskshape.GongGetIdentifier(stage))
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "CornerOffsetRatio")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%f", notetaskshape.CornerOffsetRatio))
		initializerStatements += setValueField

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
		// Initialisation of values
		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", product.GongGetIdentifier(stage))
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Name")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(product.Name))
		initializerStatements += setValueField

		setValueField = NumberInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", product.GongGetIdentifier(stage))
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "IsExpanded")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%t", product.IsExpanded))
		initializerStatements += setValueField

		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", product.GongGetIdentifier(stage))
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "ComputedPrefix")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(product.ComputedPrefix))
		initializerStatements += setValueField

		setValueField = NumberInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", product.GongGetIdentifier(stage))
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "IsProducersNodeExpanded")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%t", product.IsProducersNodeExpanded))
		initializerStatements += setValueField

		setValueField = NumberInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", product.GongGetIdentifier(stage))
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "IsConsumersNodeExpanded")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%t", product.IsConsumersNodeExpanded))
		initializerStatements += setValueField

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
		// Initialisation of values
		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", productcompositionshape.GongGetIdentifier(stage))
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Name")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(productcompositionshape.Name))
		initializerStatements += setValueField

		setValueField = NumberInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", productcompositionshape.GongGetIdentifier(stage))
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "StartRatio")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%f", productcompositionshape.StartRatio))
		initializerStatements += setValueField

		setValueField = NumberInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", productcompositionshape.GongGetIdentifier(stage))
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "EndRatio")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%f", productcompositionshape.EndRatio))
		initializerStatements += setValueField

		if productcompositionshape.StartOrientation != "" {
			setValueField = StringEnumInitStatement
			setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", productcompositionshape.GongGetIdentifier(stage))
			setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "StartOrientation")
			setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", "models."+productcompositionshape.StartOrientation.ToCodeString())
			initializerStatements += setValueField
		}

		if productcompositionshape.EndOrientation != "" {
			setValueField = StringEnumInitStatement
			setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", productcompositionshape.GongGetIdentifier(stage))
			setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "EndOrientation")
			setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", "models."+productcompositionshape.EndOrientation.ToCodeString())
			initializerStatements += setValueField
		}

		setValueField = NumberInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", productcompositionshape.GongGetIdentifier(stage))
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "CornerOffsetRatio")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%f", productcompositionshape.CornerOffsetRatio))
		initializerStatements += setValueField

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
		// Initialisation of values
		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", productshape.GongGetIdentifier(stage))
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Name")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(productshape.Name))
		initializerStatements += setValueField

		setValueField = NumberInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", productshape.GongGetIdentifier(stage))
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "IsExpanded")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%t", productshape.IsExpanded))
		initializerStatements += setValueField

		setValueField = NumberInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", productshape.GongGetIdentifier(stage))
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "X")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%f", productshape.X))
		initializerStatements += setValueField

		setValueField = NumberInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", productshape.GongGetIdentifier(stage))
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Y")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%f", productshape.Y))
		initializerStatements += setValueField

		setValueField = NumberInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", productshape.GongGetIdentifier(stage))
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Width")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%f", productshape.Width))
		initializerStatements += setValueField

		setValueField = NumberInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", productshape.GongGetIdentifier(stage))
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Height")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%f", productshape.Height))
		initializerStatements += setValueField

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
		// Initialisation of values
		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", project.GongGetIdentifier(stage))
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Name")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(project.Name))
		initializerStatements += setValueField

		setValueField = NumberInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", project.GongGetIdentifier(stage))
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "IsPBSNodeExpanded")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%t", project.IsPBSNodeExpanded))
		initializerStatements += setValueField

		setValueField = NumberInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", project.GongGetIdentifier(stage))
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "IsWBSNodeExpanded")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%t", project.IsWBSNodeExpanded))
		initializerStatements += setValueField

		setValueField = NumberInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", project.GongGetIdentifier(stage))
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "IsDiagramsNodeExpanded")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%t", project.IsDiagramsNodeExpanded))
		initializerStatements += setValueField

		setValueField = NumberInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", project.GongGetIdentifier(stage))
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "IsNotesNodeExpanded")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%t", project.IsNotesNodeExpanded))
		initializerStatements += setValueField

		setValueField = NumberInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", project.GongGetIdentifier(stage))
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "IsExpanded")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%t", project.IsExpanded))
		initializerStatements += setValueField

		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", project.GongGetIdentifier(stage))
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "ComputedPrefix")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(project.ComputedPrefix))
		initializerStatements += setValueField

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
		// Initialisation of values
		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", root.GongGetIdentifier(stage))
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Name")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(root.Name))
		initializerStatements += setValueField

		setValueField = NumberInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", root.GongGetIdentifier(stage))
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "NbPixPerCharacter")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%f", root.NbPixPerCharacter))
		initializerStatements += setValueField

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
		// Initialisation of values
		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", task.GongGetIdentifier(stage))
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Name")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(task.Name))
		initializerStatements += setValueField

		setValueField = NumberInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", task.GongGetIdentifier(stage))
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "IsExpanded")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%t", task.IsExpanded))
		initializerStatements += setValueField

		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", task.GongGetIdentifier(stage))
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "ComputedPrefix")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(task.ComputedPrefix))
		initializerStatements += setValueField

		setValueField = NumberInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", task.GongGetIdentifier(stage))
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "IsInputsNodeExpanded")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%t", task.IsInputsNodeExpanded))
		initializerStatements += setValueField

		setValueField = NumberInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", task.GongGetIdentifier(stage))
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "IsOutputsNodeExpanded")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%t", task.IsOutputsNodeExpanded))
		initializerStatements += setValueField

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
		// Initialisation of values
		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", taskcompositionshape.GongGetIdentifier(stage))
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Name")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(taskcompositionshape.Name))
		initializerStatements += setValueField

		setValueField = NumberInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", taskcompositionshape.GongGetIdentifier(stage))
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "StartRatio")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%f", taskcompositionshape.StartRatio))
		initializerStatements += setValueField

		setValueField = NumberInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", taskcompositionshape.GongGetIdentifier(stage))
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "EndRatio")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%f", taskcompositionshape.EndRatio))
		initializerStatements += setValueField

		if taskcompositionshape.StartOrientation != "" {
			setValueField = StringEnumInitStatement
			setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", taskcompositionshape.GongGetIdentifier(stage))
			setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "StartOrientation")
			setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", "models."+taskcompositionshape.StartOrientation.ToCodeString())
			initializerStatements += setValueField
		}

		if taskcompositionshape.EndOrientation != "" {
			setValueField = StringEnumInitStatement
			setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", taskcompositionshape.GongGetIdentifier(stage))
			setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "EndOrientation")
			setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", "models."+taskcompositionshape.EndOrientation.ToCodeString())
			initializerStatements += setValueField
		}

		setValueField = NumberInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", taskcompositionshape.GongGetIdentifier(stage))
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "CornerOffsetRatio")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%f", taskcompositionshape.CornerOffsetRatio))
		initializerStatements += setValueField

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
		// Initialisation of values
		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", taskinputshape.GongGetIdentifier(stage))
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Name")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(taskinputshape.Name))
		initializerStatements += setValueField

		setValueField = NumberInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", taskinputshape.GongGetIdentifier(stage))
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "StartRatio")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%f", taskinputshape.StartRatio))
		initializerStatements += setValueField

		setValueField = NumberInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", taskinputshape.GongGetIdentifier(stage))
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "EndRatio")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%f", taskinputshape.EndRatio))
		initializerStatements += setValueField

		if taskinputshape.StartOrientation != "" {
			setValueField = StringEnumInitStatement
			setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", taskinputshape.GongGetIdentifier(stage))
			setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "StartOrientation")
			setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", "models."+taskinputshape.StartOrientation.ToCodeString())
			initializerStatements += setValueField
		}

		if taskinputshape.EndOrientation != "" {
			setValueField = StringEnumInitStatement
			setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", taskinputshape.GongGetIdentifier(stage))
			setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "EndOrientation")
			setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", "models."+taskinputshape.EndOrientation.ToCodeString())
			initializerStatements += setValueField
		}

		setValueField = NumberInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", taskinputshape.GongGetIdentifier(stage))
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "CornerOffsetRatio")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%f", taskinputshape.CornerOffsetRatio))
		initializerStatements += setValueField

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
		// Initialisation of values
		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", taskoutputshape.GongGetIdentifier(stage))
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Name")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(taskoutputshape.Name))
		initializerStatements += setValueField

		setValueField = NumberInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", taskoutputshape.GongGetIdentifier(stage))
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "StartRatio")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%f", taskoutputshape.StartRatio))
		initializerStatements += setValueField

		setValueField = NumberInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", taskoutputshape.GongGetIdentifier(stage))
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "EndRatio")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%f", taskoutputshape.EndRatio))
		initializerStatements += setValueField

		if taskoutputshape.StartOrientation != "" {
			setValueField = StringEnumInitStatement
			setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", taskoutputshape.GongGetIdentifier(stage))
			setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "StartOrientation")
			setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", "models."+taskoutputshape.StartOrientation.ToCodeString())
			initializerStatements += setValueField
		}

		if taskoutputshape.EndOrientation != "" {
			setValueField = StringEnumInitStatement
			setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", taskoutputshape.GongGetIdentifier(stage))
			setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "EndOrientation")
			setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", "models."+taskoutputshape.EndOrientation.ToCodeString())
			initializerStatements += setValueField
		}

		setValueField = NumberInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", taskoutputshape.GongGetIdentifier(stage))
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "CornerOffsetRatio")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%f", taskoutputshape.CornerOffsetRatio))
		initializerStatements += setValueField

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
		// Initialisation of values
		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", taskshape.GongGetIdentifier(stage))
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Name")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(taskshape.Name))
		initializerStatements += setValueField

		setValueField = NumberInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", taskshape.GongGetIdentifier(stage))
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "IsExpanded")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%t", taskshape.IsExpanded))
		initializerStatements += setValueField

		setValueField = NumberInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", taskshape.GongGetIdentifier(stage))
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "X")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%f", taskshape.X))
		initializerStatements += setValueField

		setValueField = NumberInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", taskshape.GongGetIdentifier(stage))
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Y")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%f", taskshape.Y))
		initializerStatements += setValueField

		setValueField = NumberInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", taskshape.GongGetIdentifier(stage))
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Width")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%f", taskshape.Width))
		initializerStatements += setValueField

		setValueField = NumberInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", taskshape.GongGetIdentifier(stage))
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Height")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%f", taskshape.Height))
		initializerStatements += setValueField

	}

	// insertion initialization of objects to stage
	if len(diagramOrdered) > 0 {
		pointersInitializesStatements += "\n\t// setup of Diagram instances pointers"
	}
	for _, diagram := range diagramOrdered {
		_ = diagram
		var setPointerField string
		_ = setPointerField

		// Initialisation of values
		for _, _productshape := range diagram.Product_Shapes {
			setPointerField = SliceOfPointersFieldInitStatement
			setPointerField = strings.ReplaceAll(setPointerField, "{{Identifier}}", diagram.GongGetIdentifier(stage))
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldName}}", "Product_Shapes")
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldNameValue}}", _productshape.GongGetIdentifier(stage))
			pointersInitializesStatements += setPointerField
		}

		for _, _product := range diagram.ProductsWhoseNodeIsExpanded {
			setPointerField = SliceOfPointersFieldInitStatement
			setPointerField = strings.ReplaceAll(setPointerField, "{{Identifier}}", diagram.GongGetIdentifier(stage))
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldName}}", "ProductsWhoseNodeIsExpanded")
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldNameValue}}", _product.GongGetIdentifier(stage))
			pointersInitializesStatements += setPointerField
		}

		for _, _productcompositionshape := range diagram.ProductComposition_Shapes {
			setPointerField = SliceOfPointersFieldInitStatement
			setPointerField = strings.ReplaceAll(setPointerField, "{{Identifier}}", diagram.GongGetIdentifier(stage))
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldName}}", "ProductComposition_Shapes")
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldNameValue}}", _productcompositionshape.GongGetIdentifier(stage))
			pointersInitializesStatements += setPointerField
		}

		for _, _taskshape := range diagram.Task_Shapes {
			setPointerField = SliceOfPointersFieldInitStatement
			setPointerField = strings.ReplaceAll(setPointerField, "{{Identifier}}", diagram.GongGetIdentifier(stage))
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldName}}", "Task_Shapes")
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldNameValue}}", _taskshape.GongGetIdentifier(stage))
			pointersInitializesStatements += setPointerField
		}

		for _, _task := range diagram.TasksWhoseNodeIsExpanded {
			setPointerField = SliceOfPointersFieldInitStatement
			setPointerField = strings.ReplaceAll(setPointerField, "{{Identifier}}", diagram.GongGetIdentifier(stage))
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldName}}", "TasksWhoseNodeIsExpanded")
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldNameValue}}", _task.GongGetIdentifier(stage))
			pointersInitializesStatements += setPointerField
		}

		for _, _task := range diagram.TasksWhoseInputNodeIsExpanded {
			setPointerField = SliceOfPointersFieldInitStatement
			setPointerField = strings.ReplaceAll(setPointerField, "{{Identifier}}", diagram.GongGetIdentifier(stage))
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldName}}", "TasksWhoseInputNodeIsExpanded")
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldNameValue}}", _task.GongGetIdentifier(stage))
			pointersInitializesStatements += setPointerField
		}

		for _, _task := range diagram.TasksWhoseOutputNodeIsExpanded {
			setPointerField = SliceOfPointersFieldInitStatement
			setPointerField = strings.ReplaceAll(setPointerField, "{{Identifier}}", diagram.GongGetIdentifier(stage))
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldName}}", "TasksWhoseOutputNodeIsExpanded")
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldNameValue}}", _task.GongGetIdentifier(stage))
			pointersInitializesStatements += setPointerField
		}

		for _, _taskcompositionshape := range diagram.TaskComposition_Shapes {
			setPointerField = SliceOfPointersFieldInitStatement
			setPointerField = strings.ReplaceAll(setPointerField, "{{Identifier}}", diagram.GongGetIdentifier(stage))
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldName}}", "TaskComposition_Shapes")
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldNameValue}}", _taskcompositionshape.GongGetIdentifier(stage))
			pointersInitializesStatements += setPointerField
		}

		for _, _taskinputshape := range diagram.TaskInputShapes {
			setPointerField = SliceOfPointersFieldInitStatement
			setPointerField = strings.ReplaceAll(setPointerField, "{{Identifier}}", diagram.GongGetIdentifier(stage))
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldName}}", "TaskInputShapes")
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldNameValue}}", _taskinputshape.GongGetIdentifier(stage))
			pointersInitializesStatements += setPointerField
		}

		for _, _taskoutputshape := range diagram.TaskOutputShapes {
			setPointerField = SliceOfPointersFieldInitStatement
			setPointerField = strings.ReplaceAll(setPointerField, "{{Identifier}}", diagram.GongGetIdentifier(stage))
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldName}}", "TaskOutputShapes")
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldNameValue}}", _taskoutputshape.GongGetIdentifier(stage))
			pointersInitializesStatements += setPointerField
		}

		for _, _noteshape := range diagram.Note_Shapes {
			setPointerField = SliceOfPointersFieldInitStatement
			setPointerField = strings.ReplaceAll(setPointerField, "{{Identifier}}", diagram.GongGetIdentifier(stage))
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldName}}", "Note_Shapes")
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldNameValue}}", _noteshape.GongGetIdentifier(stage))
			pointersInitializesStatements += setPointerField
		}

		for _, _note := range diagram.NotesWhoseNodeIsExpanded {
			setPointerField = SliceOfPointersFieldInitStatement
			setPointerField = strings.ReplaceAll(setPointerField, "{{Identifier}}", diagram.GongGetIdentifier(stage))
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldName}}", "NotesWhoseNodeIsExpanded")
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldNameValue}}", _note.GongGetIdentifier(stage))
			pointersInitializesStatements += setPointerField
		}

		for _, _noteproductshape := range diagram.NoteProductShapes {
			setPointerField = SliceOfPointersFieldInitStatement
			setPointerField = strings.ReplaceAll(setPointerField, "{{Identifier}}", diagram.GongGetIdentifier(stage))
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldName}}", "NoteProductShapes")
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldNameValue}}", _noteproductshape.GongGetIdentifier(stage))
			pointersInitializesStatements += setPointerField
		}

		for _, _notetaskshape := range diagram.NoteTaskShapes {
			setPointerField = SliceOfPointersFieldInitStatement
			setPointerField = strings.ReplaceAll(setPointerField, "{{Identifier}}", diagram.GongGetIdentifier(stage))
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldName}}", "NoteTaskShapes")
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldNameValue}}", _notetaskshape.GongGetIdentifier(stage))
			pointersInitializesStatements += setPointerField
		}

	}

	if len(noteOrdered) > 0 {
		pointersInitializesStatements += "\n\t// setup of Note instances pointers"
	}
	for _, note := range noteOrdered {
		_ = note
		var setPointerField string
		_ = setPointerField

		// Initialisation of values
		for _, _product := range note.Products {
			setPointerField = SliceOfPointersFieldInitStatement
			setPointerField = strings.ReplaceAll(setPointerField, "{{Identifier}}", note.GongGetIdentifier(stage))
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldName}}", "Products")
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldNameValue}}", _product.GongGetIdentifier(stage))
			pointersInitializesStatements += setPointerField
		}

		for _, _task := range note.Tasks {
			setPointerField = SliceOfPointersFieldInitStatement
			setPointerField = strings.ReplaceAll(setPointerField, "{{Identifier}}", note.GongGetIdentifier(stage))
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldName}}", "Tasks")
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldNameValue}}", _task.GongGetIdentifier(stage))
			pointersInitializesStatements += setPointerField
		}

	}

	if len(noteproductshapeOrdered) > 0 {
		pointersInitializesStatements += "\n\t// setup of NoteProductShape instances pointers"
	}
	for _, noteproductshape := range noteproductshapeOrdered {
		_ = noteproductshape
		var setPointerField string
		_ = setPointerField

		// Initialisation of values
		if noteproductshape.Note != nil {
			setPointerField = PointerFieldInitStatement
			setPointerField = strings.ReplaceAll(setPointerField, "{{Identifier}}", noteproductshape.GongGetIdentifier(stage))
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldName}}", "Note")
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldNameValue}}", noteproductshape.Note.GongGetIdentifier(stage))
			pointersInitializesStatements += setPointerField
		}

		if noteproductshape.Product != nil {
			setPointerField = PointerFieldInitStatement
			setPointerField = strings.ReplaceAll(setPointerField, "{{Identifier}}", noteproductshape.GongGetIdentifier(stage))
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldName}}", "Product")
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldNameValue}}", noteproductshape.Product.GongGetIdentifier(stage))
			pointersInitializesStatements += setPointerField
		}

	}

	if len(noteshapeOrdered) > 0 {
		pointersInitializesStatements += "\n\t// setup of NoteShape instances pointers"
	}
	for _, noteshape := range noteshapeOrdered {
		_ = noteshape
		var setPointerField string
		_ = setPointerField

		// Initialisation of values
		if noteshape.Note != nil {
			setPointerField = PointerFieldInitStatement
			setPointerField = strings.ReplaceAll(setPointerField, "{{Identifier}}", noteshape.GongGetIdentifier(stage))
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldName}}", "Note")
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldNameValue}}", noteshape.Note.GongGetIdentifier(stage))
			pointersInitializesStatements += setPointerField
		}

	}

	if len(notetaskshapeOrdered) > 0 {
		pointersInitializesStatements += "\n\t// setup of NoteTaskShape instances pointers"
	}
	for _, notetaskshape := range notetaskshapeOrdered {
		_ = notetaskshape
		var setPointerField string
		_ = setPointerField

		// Initialisation of values
		if notetaskshape.Note != nil {
			setPointerField = PointerFieldInitStatement
			setPointerField = strings.ReplaceAll(setPointerField, "{{Identifier}}", notetaskshape.GongGetIdentifier(stage))
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldName}}", "Note")
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldNameValue}}", notetaskshape.Note.GongGetIdentifier(stage))
			pointersInitializesStatements += setPointerField
		}

		if notetaskshape.Task != nil {
			setPointerField = PointerFieldInitStatement
			setPointerField = strings.ReplaceAll(setPointerField, "{{Identifier}}", notetaskshape.GongGetIdentifier(stage))
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldName}}", "Task")
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldNameValue}}", notetaskshape.Task.GongGetIdentifier(stage))
			pointersInitializesStatements += setPointerField
		}

	}

	if len(productOrdered) > 0 {
		pointersInitializesStatements += "\n\t// setup of Product instances pointers"
	}
	for _, product := range productOrdered {
		_ = product
		var setPointerField string
		_ = setPointerField

		// Initialisation of values
		for _, _product := range product.SubProducts {
			setPointerField = SliceOfPointersFieldInitStatement
			setPointerField = strings.ReplaceAll(setPointerField, "{{Identifier}}", product.GongGetIdentifier(stage))
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldName}}", "SubProducts")
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldNameValue}}", _product.GongGetIdentifier(stage))
			pointersInitializesStatements += setPointerField
		}

	}

	if len(productcompositionshapeOrdered) > 0 {
		pointersInitializesStatements += "\n\t// setup of ProductCompositionShape instances pointers"
	}
	for _, productcompositionshape := range productcompositionshapeOrdered {
		_ = productcompositionshape
		var setPointerField string
		_ = setPointerField

		// Initialisation of values
		if productcompositionshape.Product != nil {
			setPointerField = PointerFieldInitStatement
			setPointerField = strings.ReplaceAll(setPointerField, "{{Identifier}}", productcompositionshape.GongGetIdentifier(stage))
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldName}}", "Product")
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldNameValue}}", productcompositionshape.Product.GongGetIdentifier(stage))
			pointersInitializesStatements += setPointerField
		}

	}

	if len(productshapeOrdered) > 0 {
		pointersInitializesStatements += "\n\t// setup of ProductShape instances pointers"
	}
	for _, productshape := range productshapeOrdered {
		_ = productshape
		var setPointerField string
		_ = setPointerField

		// Initialisation of values
		if productshape.Product != nil {
			setPointerField = PointerFieldInitStatement
			setPointerField = strings.ReplaceAll(setPointerField, "{{Identifier}}", productshape.GongGetIdentifier(stage))
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldName}}", "Product")
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldNameValue}}", productshape.Product.GongGetIdentifier(stage))
			pointersInitializesStatements += setPointerField
		}

	}

	if len(projectOrdered) > 0 {
		pointersInitializesStatements += "\n\t// setup of Project instances pointers"
	}
	for _, project := range projectOrdered {
		_ = project
		var setPointerField string
		_ = setPointerField

		// Initialisation of values
		for _, _product := range project.RootProducts {
			setPointerField = SliceOfPointersFieldInitStatement
			setPointerField = strings.ReplaceAll(setPointerField, "{{Identifier}}", project.GongGetIdentifier(stage))
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldName}}", "RootProducts")
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldNameValue}}", _product.GongGetIdentifier(stage))
			pointersInitializesStatements += setPointerField
		}

		for _, _task := range project.RootTasks {
			setPointerField = SliceOfPointersFieldInitStatement
			setPointerField = strings.ReplaceAll(setPointerField, "{{Identifier}}", project.GongGetIdentifier(stage))
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldName}}", "RootTasks")
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldNameValue}}", _task.GongGetIdentifier(stage))
			pointersInitializesStatements += setPointerField
		}

		for _, _diagram := range project.Diagrams {
			setPointerField = SliceOfPointersFieldInitStatement
			setPointerField = strings.ReplaceAll(setPointerField, "{{Identifier}}", project.GongGetIdentifier(stage))
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldName}}", "Diagrams")
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldNameValue}}", _diagram.GongGetIdentifier(stage))
			pointersInitializesStatements += setPointerField
		}

		for _, _note := range project.Notes {
			setPointerField = SliceOfPointersFieldInitStatement
			setPointerField = strings.ReplaceAll(setPointerField, "{{Identifier}}", project.GongGetIdentifier(stage))
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldName}}", "Notes")
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldNameValue}}", _note.GongGetIdentifier(stage))
			pointersInitializesStatements += setPointerField
		}

	}

	if len(rootOrdered) > 0 {
		pointersInitializesStatements += "\n\t// setup of Root instances pointers"
	}
	for _, root := range rootOrdered {
		_ = root
		var setPointerField string
		_ = setPointerField

		// Initialisation of values
		for _, _project := range root.Projects {
			setPointerField = SliceOfPointersFieldInitStatement
			setPointerField = strings.ReplaceAll(setPointerField, "{{Identifier}}", root.GongGetIdentifier(stage))
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldName}}", "Projects")
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldNameValue}}", _project.GongGetIdentifier(stage))
			pointersInitializesStatements += setPointerField
		}

		for _, _product := range root.OrphanedProducts {
			setPointerField = SliceOfPointersFieldInitStatement
			setPointerField = strings.ReplaceAll(setPointerField, "{{Identifier}}", root.GongGetIdentifier(stage))
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldName}}", "OrphanedProducts")
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldNameValue}}", _product.GongGetIdentifier(stage))
			pointersInitializesStatements += setPointerField
		}

		for _, _task := range root.OrphanedTasks {
			setPointerField = SliceOfPointersFieldInitStatement
			setPointerField = strings.ReplaceAll(setPointerField, "{{Identifier}}", root.GongGetIdentifier(stage))
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldName}}", "OrphanedTasks")
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldNameValue}}", _task.GongGetIdentifier(stage))
			pointersInitializesStatements += setPointerField
		}

	}

	if len(taskOrdered) > 0 {
		pointersInitializesStatements += "\n\t// setup of Task instances pointers"
	}
	for _, task := range taskOrdered {
		_ = task
		var setPointerField string
		_ = setPointerField

		// Initialisation of values
		for _, _task := range task.SubTasks {
			setPointerField = SliceOfPointersFieldInitStatement
			setPointerField = strings.ReplaceAll(setPointerField, "{{Identifier}}", task.GongGetIdentifier(stage))
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldName}}", "SubTasks")
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldNameValue}}", _task.GongGetIdentifier(stage))
			pointersInitializesStatements += setPointerField
		}

		for _, _product := range task.Inputs {
			setPointerField = SliceOfPointersFieldInitStatement
			setPointerField = strings.ReplaceAll(setPointerField, "{{Identifier}}", task.GongGetIdentifier(stage))
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldName}}", "Inputs")
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldNameValue}}", _product.GongGetIdentifier(stage))
			pointersInitializesStatements += setPointerField
		}

		for _, _product := range task.Outputs {
			setPointerField = SliceOfPointersFieldInitStatement
			setPointerField = strings.ReplaceAll(setPointerField, "{{Identifier}}", task.GongGetIdentifier(stage))
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldName}}", "Outputs")
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldNameValue}}", _product.GongGetIdentifier(stage))
			pointersInitializesStatements += setPointerField
		}

	}

	if len(taskcompositionshapeOrdered) > 0 {
		pointersInitializesStatements += "\n\t// setup of TaskCompositionShape instances pointers"
	}
	for _, taskcompositionshape := range taskcompositionshapeOrdered {
		_ = taskcompositionshape
		var setPointerField string
		_ = setPointerField

		// Initialisation of values
		if taskcompositionshape.Task != nil {
			setPointerField = PointerFieldInitStatement
			setPointerField = strings.ReplaceAll(setPointerField, "{{Identifier}}", taskcompositionshape.GongGetIdentifier(stage))
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldName}}", "Task")
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldNameValue}}", taskcompositionshape.Task.GongGetIdentifier(stage))
			pointersInitializesStatements += setPointerField
		}

	}

	if len(taskinputshapeOrdered) > 0 {
		pointersInitializesStatements += "\n\t// setup of TaskInputShape instances pointers"
	}
	for _, taskinputshape := range taskinputshapeOrdered {
		_ = taskinputshape
		var setPointerField string
		_ = setPointerField

		// Initialisation of values
		if taskinputshape.Task != nil {
			setPointerField = PointerFieldInitStatement
			setPointerField = strings.ReplaceAll(setPointerField, "{{Identifier}}", taskinputshape.GongGetIdentifier(stage))
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldName}}", "Task")
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldNameValue}}", taskinputshape.Task.GongGetIdentifier(stage))
			pointersInitializesStatements += setPointerField
		}

		if taskinputshape.Product != nil {
			setPointerField = PointerFieldInitStatement
			setPointerField = strings.ReplaceAll(setPointerField, "{{Identifier}}", taskinputshape.GongGetIdentifier(stage))
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldName}}", "Product")
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldNameValue}}", taskinputshape.Product.GongGetIdentifier(stage))
			pointersInitializesStatements += setPointerField
		}

	}

	if len(taskoutputshapeOrdered) > 0 {
		pointersInitializesStatements += "\n\t// setup of TaskOutputShape instances pointers"
	}
	for _, taskoutputshape := range taskoutputshapeOrdered {
		_ = taskoutputshape
		var setPointerField string
		_ = setPointerField

		// Initialisation of values
		if taskoutputshape.Task != nil {
			setPointerField = PointerFieldInitStatement
			setPointerField = strings.ReplaceAll(setPointerField, "{{Identifier}}", taskoutputshape.GongGetIdentifier(stage))
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldName}}", "Task")
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldNameValue}}", taskoutputshape.Task.GongGetIdentifier(stage))
			pointersInitializesStatements += setPointerField
		}

		if taskoutputshape.Product != nil {
			setPointerField = PointerFieldInitStatement
			setPointerField = strings.ReplaceAll(setPointerField, "{{Identifier}}", taskoutputshape.GongGetIdentifier(stage))
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldName}}", "Product")
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldNameValue}}", taskoutputshape.Product.GongGetIdentifier(stage))
			pointersInitializesStatements += setPointerField
		}

	}

	if len(taskshapeOrdered) > 0 {
		pointersInitializesStatements += "\n\t// setup of TaskShape instances pointers"
	}
	for _, taskshape := range taskshapeOrdered {
		_ = taskshape
		var setPointerField string
		_ = setPointerField

		// Initialisation of values
		if taskshape.Task != nil {
			setPointerField = PointerFieldInitStatement
			setPointerField = strings.ReplaceAll(setPointerField, "{{Identifier}}", taskshape.GongGetIdentifier(stage))
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldName}}", "Task")
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldNameValue}}", taskshape.Task.GongGetIdentifier(stage))
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
func (diagram *Diagram) GongMarshallField(stage *Stage, fieldName string) (setValueField, setPointerField string) {
	initializerStatements := ""
	_ = initializerStatements
	pointersInitializesStatements := ""
	_ = pointersInitializesStatements

	switch fieldName {
	case "Name":
		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", diagram.GongGetIdentifier(stage))
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Name")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(diagram.Name))
		initializerStatements += setValueField
	case "IsChecked":
		setValueField = NumberInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", diagram.GongGetIdentifier(stage))
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "IsChecked")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%t", diagram.IsChecked))
		initializerStatements += setValueField
	case "IsEditable_":
		setValueField = NumberInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", diagram.GongGetIdentifier(stage))
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "IsEditable_")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%t", diagram.IsEditable_))
		initializerStatements += setValueField
	case "IsInRenameMode":
		setValueField = NumberInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", diagram.GongGetIdentifier(stage))
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "IsInRenameMode")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%t", diagram.IsInRenameMode))
		initializerStatements += setValueField
	case "ShowPrefix":
		setValueField = NumberInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", diagram.GongGetIdentifier(stage))
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "ShowPrefix")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%t", diagram.ShowPrefix))
		initializerStatements += setValueField
	case "DefaultBoxWidth":
		setValueField = NumberInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", diagram.GongGetIdentifier(stage))
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "DefaultBoxWidth")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%f", diagram.DefaultBoxWidth))
		initializerStatements += setValueField
	case "DefaultBoxHeigth":
		setValueField = NumberInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", diagram.GongGetIdentifier(stage))
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "DefaultBoxHeigth")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%f", diagram.DefaultBoxHeigth))
		initializerStatements += setValueField
	case "IsExpanded":
		setValueField = NumberInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", diagram.GongGetIdentifier(stage))
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "IsExpanded")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%t", diagram.IsExpanded))
		initializerStatements += setValueField
	case "ComputedPrefix":
		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", diagram.GongGetIdentifier(stage))
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "ComputedPrefix")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(diagram.ComputedPrefix))
		initializerStatements += setValueField
	case "IsPBSNodeExpanded":
		setValueField = NumberInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", diagram.GongGetIdentifier(stage))
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "IsPBSNodeExpanded")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%t", diagram.IsPBSNodeExpanded))
		initializerStatements += setValueField
	case "IsWBSNodeExpanded":
		setValueField = NumberInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", diagram.GongGetIdentifier(stage))
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "IsWBSNodeExpanded")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%t", diagram.IsWBSNodeExpanded))
		initializerStatements += setValueField
	case "IsNotesNodeExpanded":
		setValueField = NumberInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", diagram.GongGetIdentifier(stage))
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "IsNotesNodeExpanded")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%t", diagram.IsNotesNodeExpanded))
		initializerStatements += setValueField

	case "Product_Shapes":
		for _, _productshape := range diagram.Product_Shapes {
			setPointerField = SliceOfPointersFieldInitStatement
			setPointerField = strings.ReplaceAll(setPointerField, "{{Identifier}}", diagram.GongGetIdentifier(stage))
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldName}}", "Product_Shapes")
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldNameValue}}", _productshape.GongGetIdentifier(stage))
			pointersInitializesStatements += setPointerField
		}
	case "ProductsWhoseNodeIsExpanded":
		for _, _product := range diagram.ProductsWhoseNodeIsExpanded {
			setPointerField = SliceOfPointersFieldInitStatement
			setPointerField = strings.ReplaceAll(setPointerField, "{{Identifier}}", diagram.GongGetIdentifier(stage))
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldName}}", "ProductsWhoseNodeIsExpanded")
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldNameValue}}", _product.GongGetIdentifier(stage))
			pointersInitializesStatements += setPointerField
		}
	case "ProductComposition_Shapes":
		for _, _productcompositionshape := range diagram.ProductComposition_Shapes {
			setPointerField = SliceOfPointersFieldInitStatement
			setPointerField = strings.ReplaceAll(setPointerField, "{{Identifier}}", diagram.GongGetIdentifier(stage))
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldName}}", "ProductComposition_Shapes")
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldNameValue}}", _productcompositionshape.GongGetIdentifier(stage))
			pointersInitializesStatements += setPointerField
		}
	case "Task_Shapes":
		for _, _taskshape := range diagram.Task_Shapes {
			setPointerField = SliceOfPointersFieldInitStatement
			setPointerField = strings.ReplaceAll(setPointerField, "{{Identifier}}", diagram.GongGetIdentifier(stage))
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldName}}", "Task_Shapes")
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldNameValue}}", _taskshape.GongGetIdentifier(stage))
			pointersInitializesStatements += setPointerField
		}
	case "TasksWhoseNodeIsExpanded":
		for _, _task := range diagram.TasksWhoseNodeIsExpanded {
			setPointerField = SliceOfPointersFieldInitStatement
			setPointerField = strings.ReplaceAll(setPointerField, "{{Identifier}}", diagram.GongGetIdentifier(stage))
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldName}}", "TasksWhoseNodeIsExpanded")
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldNameValue}}", _task.GongGetIdentifier(stage))
			pointersInitializesStatements += setPointerField
		}
	case "TasksWhoseInputNodeIsExpanded":
		for _, _task := range diagram.TasksWhoseInputNodeIsExpanded {
			setPointerField = SliceOfPointersFieldInitStatement
			setPointerField = strings.ReplaceAll(setPointerField, "{{Identifier}}", diagram.GongGetIdentifier(stage))
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldName}}", "TasksWhoseInputNodeIsExpanded")
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldNameValue}}", _task.GongGetIdentifier(stage))
			pointersInitializesStatements += setPointerField
		}
	case "TasksWhoseOutputNodeIsExpanded":
		for _, _task := range diagram.TasksWhoseOutputNodeIsExpanded {
			setPointerField = SliceOfPointersFieldInitStatement
			setPointerField = strings.ReplaceAll(setPointerField, "{{Identifier}}", diagram.GongGetIdentifier(stage))
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldName}}", "TasksWhoseOutputNodeIsExpanded")
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldNameValue}}", _task.GongGetIdentifier(stage))
			pointersInitializesStatements += setPointerField
		}
	case "TaskComposition_Shapes":
		for _, _taskcompositionshape := range diagram.TaskComposition_Shapes {
			setPointerField = SliceOfPointersFieldInitStatement
			setPointerField = strings.ReplaceAll(setPointerField, "{{Identifier}}", diagram.GongGetIdentifier(stage))
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldName}}", "TaskComposition_Shapes")
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldNameValue}}", _taskcompositionshape.GongGetIdentifier(stage))
			pointersInitializesStatements += setPointerField
		}
	case "TaskInputShapes":
		for _, _taskinputshape := range diagram.TaskInputShapes {
			setPointerField = SliceOfPointersFieldInitStatement
			setPointerField = strings.ReplaceAll(setPointerField, "{{Identifier}}", diagram.GongGetIdentifier(stage))
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldName}}", "TaskInputShapes")
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldNameValue}}", _taskinputshape.GongGetIdentifier(stage))
			pointersInitializesStatements += setPointerField
		}
	case "TaskOutputShapes":
		for _, _taskoutputshape := range diagram.TaskOutputShapes {
			setPointerField = SliceOfPointersFieldInitStatement
			setPointerField = strings.ReplaceAll(setPointerField, "{{Identifier}}", diagram.GongGetIdentifier(stage))
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldName}}", "TaskOutputShapes")
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldNameValue}}", _taskoutputshape.GongGetIdentifier(stage))
			pointersInitializesStatements += setPointerField
		}
	case "Note_Shapes":
		for _, _noteshape := range diagram.Note_Shapes {
			setPointerField = SliceOfPointersFieldInitStatement
			setPointerField = strings.ReplaceAll(setPointerField, "{{Identifier}}", diagram.GongGetIdentifier(stage))
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldName}}", "Note_Shapes")
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldNameValue}}", _noteshape.GongGetIdentifier(stage))
			pointersInitializesStatements += setPointerField
		}
	case "NotesWhoseNodeIsExpanded":
		for _, _note := range diagram.NotesWhoseNodeIsExpanded {
			setPointerField = SliceOfPointersFieldInitStatement
			setPointerField = strings.ReplaceAll(setPointerField, "{{Identifier}}", diagram.GongGetIdentifier(stage))
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldName}}", "NotesWhoseNodeIsExpanded")
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldNameValue}}", _note.GongGetIdentifier(stage))
			pointersInitializesStatements += setPointerField
		}
	case "NoteProductShapes":
		for _, _noteproductshape := range diagram.NoteProductShapes {
			setPointerField = SliceOfPointersFieldInitStatement
			setPointerField = strings.ReplaceAll(setPointerField, "{{Identifier}}", diagram.GongGetIdentifier(stage))
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldName}}", "NoteProductShapes")
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldNameValue}}", _noteproductshape.GongGetIdentifier(stage))
			pointersInitializesStatements += setPointerField
		}
	case "NoteTaskShapes":
		for _, _notetaskshape := range diagram.NoteTaskShapes {
			setPointerField = SliceOfPointersFieldInitStatement
			setPointerField = strings.ReplaceAll(setPointerField, "{{Identifier}}", diagram.GongGetIdentifier(stage))
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldName}}", "NoteTaskShapes")
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldNameValue}}", _notetaskshape.GongGetIdentifier(stage))
			pointersInitializesStatements += setPointerField
		}

	default:
		log.Panicf("Unknown field %s for Gongstruct Diagram", fieldName)
	}
	return
}

func (note *Note) GongMarshallField(stage *Stage, fieldName string) (setValueField, setPointerField string) {
	initializerStatements := ""
	_ = initializerStatements
	pointersInitializesStatements := ""
	_ = pointersInitializesStatements

	switch fieldName {
	case "Name":
		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", note.GongGetIdentifier(stage))
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Name")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(note.Name))
		initializerStatements += setValueField
	case "IsExpanded":
		setValueField = NumberInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", note.GongGetIdentifier(stage))
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "IsExpanded")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%t", note.IsExpanded))
		initializerStatements += setValueField

	case "Products":
		for _, _product := range note.Products {
			setPointerField = SliceOfPointersFieldInitStatement
			setPointerField = strings.ReplaceAll(setPointerField, "{{Identifier}}", note.GongGetIdentifier(stage))
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldName}}", "Products")
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldNameValue}}", _product.GongGetIdentifier(stage))
			pointersInitializesStatements += setPointerField
		}
	case "Tasks":
		for _, _task := range note.Tasks {
			setPointerField = SliceOfPointersFieldInitStatement
			setPointerField = strings.ReplaceAll(setPointerField, "{{Identifier}}", note.GongGetIdentifier(stage))
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldName}}", "Tasks")
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldNameValue}}", _task.GongGetIdentifier(stage))
			pointersInitializesStatements += setPointerField
		}

	default:
		log.Panicf("Unknown field %s for Gongstruct Note", fieldName)
	}
	return
}

func (noteproductshape *NoteProductShape) GongMarshallField(stage *Stage, fieldName string) (setValueField, setPointerField string) {
	initializerStatements := ""
	_ = initializerStatements
	pointersInitializesStatements := ""
	_ = pointersInitializesStatements

	switch fieldName {
	case "Name":
		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", noteproductshape.GongGetIdentifier(stage))
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Name")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(noteproductshape.Name))
		initializerStatements += setValueField
	case "StartRatio":
		setValueField = NumberInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", noteproductshape.GongGetIdentifier(stage))
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "StartRatio")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%f", noteproductshape.StartRatio))
		initializerStatements += setValueField
	case "EndRatio":
		setValueField = NumberInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", noteproductshape.GongGetIdentifier(stage))
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "EndRatio")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%f", noteproductshape.EndRatio))
		initializerStatements += setValueField
	case "StartOrientation":
		if noteproductshape.StartOrientation != "" {
			setValueField = StringEnumInitStatement
			setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", noteproductshape.GongGetIdentifier(stage))
			setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "StartOrientation")
			setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", "models."+noteproductshape.StartOrientation.ToCodeString())
			initializerStatements += setValueField
		}
	case "EndOrientation":
		if noteproductshape.EndOrientation != "" {
			setValueField = StringEnumInitStatement
			setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", noteproductshape.GongGetIdentifier(stage))
			setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "EndOrientation")
			setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", "models."+noteproductshape.EndOrientation.ToCodeString())
			initializerStatements += setValueField
		}
	case "CornerOffsetRatio":
		setValueField = NumberInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", noteproductshape.GongGetIdentifier(stage))
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "CornerOffsetRatio")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%f", noteproductshape.CornerOffsetRatio))
		initializerStatements += setValueField

	case "Note":
		if noteproductshape.Note != nil {
			setPointerField = PointerFieldInitStatement
			setPointerField = strings.ReplaceAll(setPointerField, "{{Identifier}}", noteproductshape.GongGetIdentifier(stage))
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldName}}", "Note")
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldNameValue}}", noteproductshape.Note.GongGetIdentifier(stage))
			pointersInitializesStatements += setPointerField
		}
	case "Product":
		if noteproductshape.Product != nil {
			setPointerField = PointerFieldInitStatement
			setPointerField = strings.ReplaceAll(setPointerField, "{{Identifier}}", noteproductshape.GongGetIdentifier(stage))
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldName}}", "Product")
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldNameValue}}", noteproductshape.Product.GongGetIdentifier(stage))
			pointersInitializesStatements += setPointerField
		}

	default:
		log.Panicf("Unknown field %s for Gongstruct NoteProductShape", fieldName)
	}
	return
}

func (noteshape *NoteShape) GongMarshallField(stage *Stage, fieldName string) (setValueField, setPointerField string) {
	initializerStatements := ""
	_ = initializerStatements
	pointersInitializesStatements := ""
	_ = pointersInitializesStatements

	switch fieldName {
	case "Name":
		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", noteshape.GongGetIdentifier(stage))
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Name")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(noteshape.Name))
		initializerStatements += setValueField
	case "IsExpanded":
		setValueField = NumberInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", noteshape.GongGetIdentifier(stage))
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "IsExpanded")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%t", noteshape.IsExpanded))
		initializerStatements += setValueField
	case "X":
		setValueField = NumberInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", noteshape.GongGetIdentifier(stage))
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "X")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%f", noteshape.X))
		initializerStatements += setValueField
	case "Y":
		setValueField = NumberInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", noteshape.GongGetIdentifier(stage))
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Y")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%f", noteshape.Y))
		initializerStatements += setValueField
	case "Width":
		setValueField = NumberInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", noteshape.GongGetIdentifier(stage))
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Width")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%f", noteshape.Width))
		initializerStatements += setValueField
	case "Height":
		setValueField = NumberInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", noteshape.GongGetIdentifier(stage))
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Height")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%f", noteshape.Height))
		initializerStatements += setValueField

	case "Note":
		if noteshape.Note != nil {
			setPointerField = PointerFieldInitStatement
			setPointerField = strings.ReplaceAll(setPointerField, "{{Identifier}}", noteshape.GongGetIdentifier(stage))
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldName}}", "Note")
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldNameValue}}", noteshape.Note.GongGetIdentifier(stage))
			pointersInitializesStatements += setPointerField
		}

	default:
		log.Panicf("Unknown field %s for Gongstruct NoteShape", fieldName)
	}
	return
}

func (notetaskshape *NoteTaskShape) GongMarshallField(stage *Stage, fieldName string) (setValueField, setPointerField string) {
	initializerStatements := ""
	_ = initializerStatements
	pointersInitializesStatements := ""
	_ = pointersInitializesStatements

	switch fieldName {
	case "Name":
		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", notetaskshape.GongGetIdentifier(stage))
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Name")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(notetaskshape.Name))
		initializerStatements += setValueField
	case "StartRatio":
		setValueField = NumberInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", notetaskshape.GongGetIdentifier(stage))
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "StartRatio")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%f", notetaskshape.StartRatio))
		initializerStatements += setValueField
	case "EndRatio":
		setValueField = NumberInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", notetaskshape.GongGetIdentifier(stage))
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "EndRatio")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%f", notetaskshape.EndRatio))
		initializerStatements += setValueField
	case "StartOrientation":
		if notetaskshape.StartOrientation != "" {
			setValueField = StringEnumInitStatement
			setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", notetaskshape.GongGetIdentifier(stage))
			setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "StartOrientation")
			setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", "models."+notetaskshape.StartOrientation.ToCodeString())
			initializerStatements += setValueField
		}
	case "EndOrientation":
		if notetaskshape.EndOrientation != "" {
			setValueField = StringEnumInitStatement
			setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", notetaskshape.GongGetIdentifier(stage))
			setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "EndOrientation")
			setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", "models."+notetaskshape.EndOrientation.ToCodeString())
			initializerStatements += setValueField
		}
	case "CornerOffsetRatio":
		setValueField = NumberInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", notetaskshape.GongGetIdentifier(stage))
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "CornerOffsetRatio")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%f", notetaskshape.CornerOffsetRatio))
		initializerStatements += setValueField

	case "Note":
		if notetaskshape.Note != nil {
			setPointerField = PointerFieldInitStatement
			setPointerField = strings.ReplaceAll(setPointerField, "{{Identifier}}", notetaskshape.GongGetIdentifier(stage))
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldName}}", "Note")
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldNameValue}}", notetaskshape.Note.GongGetIdentifier(stage))
			pointersInitializesStatements += setPointerField
		}
	case "Task":
		if notetaskshape.Task != nil {
			setPointerField = PointerFieldInitStatement
			setPointerField = strings.ReplaceAll(setPointerField, "{{Identifier}}", notetaskshape.GongGetIdentifier(stage))
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldName}}", "Task")
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldNameValue}}", notetaskshape.Task.GongGetIdentifier(stage))
			pointersInitializesStatements += setPointerField
		}

	default:
		log.Panicf("Unknown field %s for Gongstruct NoteTaskShape", fieldName)
	}
	return
}

func (product *Product) GongMarshallField(stage *Stage, fieldName string) (setValueField, setPointerField string) {
	initializerStatements := ""
	_ = initializerStatements
	pointersInitializesStatements := ""
	_ = pointersInitializesStatements

	switch fieldName {
	case "Name":
		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", product.GongGetIdentifier(stage))
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Name")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(product.Name))
		initializerStatements += setValueField
	case "IsExpanded":
		setValueField = NumberInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", product.GongGetIdentifier(stage))
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "IsExpanded")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%t", product.IsExpanded))
		initializerStatements += setValueField
	case "ComputedPrefix":
		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", product.GongGetIdentifier(stage))
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "ComputedPrefix")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(product.ComputedPrefix))
		initializerStatements += setValueField
	case "IsProducersNodeExpanded":
		setValueField = NumberInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", product.GongGetIdentifier(stage))
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "IsProducersNodeExpanded")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%t", product.IsProducersNodeExpanded))
		initializerStatements += setValueField
	case "IsConsumersNodeExpanded":
		setValueField = NumberInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", product.GongGetIdentifier(stage))
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "IsConsumersNodeExpanded")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%t", product.IsConsumersNodeExpanded))
		initializerStatements += setValueField

	case "SubProducts":
		for _, _product := range product.SubProducts {
			setPointerField = SliceOfPointersFieldInitStatement
			setPointerField = strings.ReplaceAll(setPointerField, "{{Identifier}}", product.GongGetIdentifier(stage))
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldName}}", "SubProducts")
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldNameValue}}", _product.GongGetIdentifier(stage))
			pointersInitializesStatements += setPointerField
		}

	default:
		log.Panicf("Unknown field %s for Gongstruct Product", fieldName)
	}
	return
}

func (productcompositionshape *ProductCompositionShape) GongMarshallField(stage *Stage, fieldName string) (setValueField, setPointerField string) {
	initializerStatements := ""
	_ = initializerStatements
	pointersInitializesStatements := ""
	_ = pointersInitializesStatements

	switch fieldName {
	case "Name":
		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", productcompositionshape.GongGetIdentifier(stage))
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Name")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(productcompositionshape.Name))
		initializerStatements += setValueField
	case "StartRatio":
		setValueField = NumberInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", productcompositionshape.GongGetIdentifier(stage))
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "StartRatio")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%f", productcompositionshape.StartRatio))
		initializerStatements += setValueField
	case "EndRatio":
		setValueField = NumberInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", productcompositionshape.GongGetIdentifier(stage))
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "EndRatio")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%f", productcompositionshape.EndRatio))
		initializerStatements += setValueField
	case "StartOrientation":
		if productcompositionshape.StartOrientation != "" {
			setValueField = StringEnumInitStatement
			setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", productcompositionshape.GongGetIdentifier(stage))
			setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "StartOrientation")
			setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", "models."+productcompositionshape.StartOrientation.ToCodeString())
			initializerStatements += setValueField
		}
	case "EndOrientation":
		if productcompositionshape.EndOrientation != "" {
			setValueField = StringEnumInitStatement
			setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", productcompositionshape.GongGetIdentifier(stage))
			setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "EndOrientation")
			setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", "models."+productcompositionshape.EndOrientation.ToCodeString())
			initializerStatements += setValueField
		}
	case "CornerOffsetRatio":
		setValueField = NumberInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", productcompositionshape.GongGetIdentifier(stage))
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "CornerOffsetRatio")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%f", productcompositionshape.CornerOffsetRatio))
		initializerStatements += setValueField

	case "Product":
		if productcompositionshape.Product != nil {
			setPointerField = PointerFieldInitStatement
			setPointerField = strings.ReplaceAll(setPointerField, "{{Identifier}}", productcompositionshape.GongGetIdentifier(stage))
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldName}}", "Product")
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldNameValue}}", productcompositionshape.Product.GongGetIdentifier(stage))
			pointersInitializesStatements += setPointerField
		}

	default:
		log.Panicf("Unknown field %s for Gongstruct ProductCompositionShape", fieldName)
	}
	return
}

func (productshape *ProductShape) GongMarshallField(stage *Stage, fieldName string) (setValueField, setPointerField string) {
	initializerStatements := ""
	_ = initializerStatements
	pointersInitializesStatements := ""
	_ = pointersInitializesStatements

	switch fieldName {
	case "Name":
		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", productshape.GongGetIdentifier(stage))
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Name")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(productshape.Name))
		initializerStatements += setValueField
	case "IsExpanded":
		setValueField = NumberInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", productshape.GongGetIdentifier(stage))
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "IsExpanded")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%t", productshape.IsExpanded))
		initializerStatements += setValueField
	case "X":
		setValueField = NumberInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", productshape.GongGetIdentifier(stage))
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "X")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%f", productshape.X))
		initializerStatements += setValueField
	case "Y":
		setValueField = NumberInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", productshape.GongGetIdentifier(stage))
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Y")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%f", productshape.Y))
		initializerStatements += setValueField
	case "Width":
		setValueField = NumberInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", productshape.GongGetIdentifier(stage))
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Width")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%f", productshape.Width))
		initializerStatements += setValueField
	case "Height":
		setValueField = NumberInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", productshape.GongGetIdentifier(stage))
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Height")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%f", productshape.Height))
		initializerStatements += setValueField

	case "Product":
		if productshape.Product != nil {
			setPointerField = PointerFieldInitStatement
			setPointerField = strings.ReplaceAll(setPointerField, "{{Identifier}}", productshape.GongGetIdentifier(stage))
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldName}}", "Product")
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldNameValue}}", productshape.Product.GongGetIdentifier(stage))
			pointersInitializesStatements += setPointerField
		}

	default:
		log.Panicf("Unknown field %s for Gongstruct ProductShape", fieldName)
	}
	return
}

func (project *Project) GongMarshallField(stage *Stage, fieldName string) (setValueField, setPointerField string) {
	initializerStatements := ""
	_ = initializerStatements
	pointersInitializesStatements := ""
	_ = pointersInitializesStatements

	switch fieldName {
	case "Name":
		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", project.GongGetIdentifier(stage))
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Name")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(project.Name))
		initializerStatements += setValueField
	case "IsPBSNodeExpanded":
		setValueField = NumberInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", project.GongGetIdentifier(stage))
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "IsPBSNodeExpanded")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%t", project.IsPBSNodeExpanded))
		initializerStatements += setValueField
	case "IsWBSNodeExpanded":
		setValueField = NumberInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", project.GongGetIdentifier(stage))
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "IsWBSNodeExpanded")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%t", project.IsWBSNodeExpanded))
		initializerStatements += setValueField
	case "IsDiagramsNodeExpanded":
		setValueField = NumberInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", project.GongGetIdentifier(stage))
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "IsDiagramsNodeExpanded")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%t", project.IsDiagramsNodeExpanded))
		initializerStatements += setValueField
	case "IsNotesNodeExpanded":
		setValueField = NumberInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", project.GongGetIdentifier(stage))
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "IsNotesNodeExpanded")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%t", project.IsNotesNodeExpanded))
		initializerStatements += setValueField
	case "IsExpanded":
		setValueField = NumberInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", project.GongGetIdentifier(stage))
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "IsExpanded")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%t", project.IsExpanded))
		initializerStatements += setValueField
	case "ComputedPrefix":
		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", project.GongGetIdentifier(stage))
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "ComputedPrefix")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(project.ComputedPrefix))
		initializerStatements += setValueField

	case "RootProducts":
		for _, _product := range project.RootProducts {
			setPointerField = SliceOfPointersFieldInitStatement
			setPointerField = strings.ReplaceAll(setPointerField, "{{Identifier}}", project.GongGetIdentifier(stage))
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldName}}", "RootProducts")
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldNameValue}}", _product.GongGetIdentifier(stage))
			pointersInitializesStatements += setPointerField
		}
	case "RootTasks":
		for _, _task := range project.RootTasks {
			setPointerField = SliceOfPointersFieldInitStatement
			setPointerField = strings.ReplaceAll(setPointerField, "{{Identifier}}", project.GongGetIdentifier(stage))
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldName}}", "RootTasks")
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldNameValue}}", _task.GongGetIdentifier(stage))
			pointersInitializesStatements += setPointerField
		}
	case "Diagrams":
		for _, _diagram := range project.Diagrams {
			setPointerField = SliceOfPointersFieldInitStatement
			setPointerField = strings.ReplaceAll(setPointerField, "{{Identifier}}", project.GongGetIdentifier(stage))
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldName}}", "Diagrams")
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldNameValue}}", _diagram.GongGetIdentifier(stage))
			pointersInitializesStatements += setPointerField
		}
	case "Notes":
		for _, _note := range project.Notes {
			setPointerField = SliceOfPointersFieldInitStatement
			setPointerField = strings.ReplaceAll(setPointerField, "{{Identifier}}", project.GongGetIdentifier(stage))
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldName}}", "Notes")
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldNameValue}}", _note.GongGetIdentifier(stage))
			pointersInitializesStatements += setPointerField
		}

	default:
		log.Panicf("Unknown field %s for Gongstruct Project", fieldName)
	}
	return
}

func (root *Root) GongMarshallField(stage *Stage, fieldName string) (setValueField, setPointerField string) {
	initializerStatements := ""
	_ = initializerStatements
	pointersInitializesStatements := ""
	_ = pointersInitializesStatements

	switch fieldName {
	case "Name":
		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", root.GongGetIdentifier(stage))
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Name")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(root.Name))
		initializerStatements += setValueField
	case "NbPixPerCharacter":
		setValueField = NumberInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", root.GongGetIdentifier(stage))
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "NbPixPerCharacter")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%f", root.NbPixPerCharacter))
		initializerStatements += setValueField

	case "Projects":
		for _, _project := range root.Projects {
			setPointerField = SliceOfPointersFieldInitStatement
			setPointerField = strings.ReplaceAll(setPointerField, "{{Identifier}}", root.GongGetIdentifier(stage))
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldName}}", "Projects")
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldNameValue}}", _project.GongGetIdentifier(stage))
			pointersInitializesStatements += setPointerField
		}
	case "OrphanedProducts":
		for _, _product := range root.OrphanedProducts {
			setPointerField = SliceOfPointersFieldInitStatement
			setPointerField = strings.ReplaceAll(setPointerField, "{{Identifier}}", root.GongGetIdentifier(stage))
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldName}}", "OrphanedProducts")
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldNameValue}}", _product.GongGetIdentifier(stage))
			pointersInitializesStatements += setPointerField
		}
	case "OrphanedTasks":
		for _, _task := range root.OrphanedTasks {
			setPointerField = SliceOfPointersFieldInitStatement
			setPointerField = strings.ReplaceAll(setPointerField, "{{Identifier}}", root.GongGetIdentifier(stage))
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldName}}", "OrphanedTasks")
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldNameValue}}", _task.GongGetIdentifier(stage))
			pointersInitializesStatements += setPointerField
		}

	default:
		log.Panicf("Unknown field %s for Gongstruct Root", fieldName)
	}
	return
}

func (task *Task) GongMarshallField(stage *Stage, fieldName string) (setValueField, setPointerField string) {
	initializerStatements := ""
	_ = initializerStatements
	pointersInitializesStatements := ""
	_ = pointersInitializesStatements

	switch fieldName {
	case "Name":
		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", task.GongGetIdentifier(stage))
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Name")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(task.Name))
		initializerStatements += setValueField
	case "IsExpanded":
		setValueField = NumberInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", task.GongGetIdentifier(stage))
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "IsExpanded")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%t", task.IsExpanded))
		initializerStatements += setValueField
	case "ComputedPrefix":
		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", task.GongGetIdentifier(stage))
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "ComputedPrefix")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(task.ComputedPrefix))
		initializerStatements += setValueField
	case "IsInputsNodeExpanded":
		setValueField = NumberInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", task.GongGetIdentifier(stage))
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "IsInputsNodeExpanded")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%t", task.IsInputsNodeExpanded))
		initializerStatements += setValueField
	case "IsOutputsNodeExpanded":
		setValueField = NumberInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", task.GongGetIdentifier(stage))
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "IsOutputsNodeExpanded")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%t", task.IsOutputsNodeExpanded))
		initializerStatements += setValueField

	case "SubTasks":
		for _, _task := range task.SubTasks {
			setPointerField = SliceOfPointersFieldInitStatement
			setPointerField = strings.ReplaceAll(setPointerField, "{{Identifier}}", task.GongGetIdentifier(stage))
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldName}}", "SubTasks")
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldNameValue}}", _task.GongGetIdentifier(stage))
			pointersInitializesStatements += setPointerField
		}
	case "Inputs":
		for _, _product := range task.Inputs {
			setPointerField = SliceOfPointersFieldInitStatement
			setPointerField = strings.ReplaceAll(setPointerField, "{{Identifier}}", task.GongGetIdentifier(stage))
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldName}}", "Inputs")
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldNameValue}}", _product.GongGetIdentifier(stage))
			pointersInitializesStatements += setPointerField
		}
	case "Outputs":
		for _, _product := range task.Outputs {
			setPointerField = SliceOfPointersFieldInitStatement
			setPointerField = strings.ReplaceAll(setPointerField, "{{Identifier}}", task.GongGetIdentifier(stage))
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldName}}", "Outputs")
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldNameValue}}", _product.GongGetIdentifier(stage))
			pointersInitializesStatements += setPointerField
		}

	default:
		log.Panicf("Unknown field %s for Gongstruct Task", fieldName)
	}
	return
}

func (taskcompositionshape *TaskCompositionShape) GongMarshallField(stage *Stage, fieldName string) (setValueField, setPointerField string) {
	initializerStatements := ""
	_ = initializerStatements
	pointersInitializesStatements := ""
	_ = pointersInitializesStatements

	switch fieldName {
	case "Name":
		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", taskcompositionshape.GongGetIdentifier(stage))
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Name")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(taskcompositionshape.Name))
		initializerStatements += setValueField
	case "StartRatio":
		setValueField = NumberInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", taskcompositionshape.GongGetIdentifier(stage))
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "StartRatio")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%f", taskcompositionshape.StartRatio))
		initializerStatements += setValueField
	case "EndRatio":
		setValueField = NumberInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", taskcompositionshape.GongGetIdentifier(stage))
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "EndRatio")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%f", taskcompositionshape.EndRatio))
		initializerStatements += setValueField
	case "StartOrientation":
		if taskcompositionshape.StartOrientation != "" {
			setValueField = StringEnumInitStatement
			setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", taskcompositionshape.GongGetIdentifier(stage))
			setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "StartOrientation")
			setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", "models."+taskcompositionshape.StartOrientation.ToCodeString())
			initializerStatements += setValueField
		}
	case "EndOrientation":
		if taskcompositionshape.EndOrientation != "" {
			setValueField = StringEnumInitStatement
			setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", taskcompositionshape.GongGetIdentifier(stage))
			setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "EndOrientation")
			setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", "models."+taskcompositionshape.EndOrientation.ToCodeString())
			initializerStatements += setValueField
		}
	case "CornerOffsetRatio":
		setValueField = NumberInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", taskcompositionshape.GongGetIdentifier(stage))
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "CornerOffsetRatio")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%f", taskcompositionshape.CornerOffsetRatio))
		initializerStatements += setValueField

	case "Task":
		if taskcompositionshape.Task != nil {
			setPointerField = PointerFieldInitStatement
			setPointerField = strings.ReplaceAll(setPointerField, "{{Identifier}}", taskcompositionshape.GongGetIdentifier(stage))
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldName}}", "Task")
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldNameValue}}", taskcompositionshape.Task.GongGetIdentifier(stage))
			pointersInitializesStatements += setPointerField
		}

	default:
		log.Panicf("Unknown field %s for Gongstruct TaskCompositionShape", fieldName)
	}
	return
}

func (taskinputshape *TaskInputShape) GongMarshallField(stage *Stage, fieldName string) (setValueField, setPointerField string) {
	initializerStatements := ""
	_ = initializerStatements
	pointersInitializesStatements := ""
	_ = pointersInitializesStatements

	switch fieldName {
	case "Name":
		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", taskinputshape.GongGetIdentifier(stage))
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Name")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(taskinputshape.Name))
		initializerStatements += setValueField
	case "StartRatio":
		setValueField = NumberInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", taskinputshape.GongGetIdentifier(stage))
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "StartRatio")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%f", taskinputshape.StartRatio))
		initializerStatements += setValueField
	case "EndRatio":
		setValueField = NumberInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", taskinputshape.GongGetIdentifier(stage))
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "EndRatio")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%f", taskinputshape.EndRatio))
		initializerStatements += setValueField
	case "StartOrientation":
		if taskinputshape.StartOrientation != "" {
			setValueField = StringEnumInitStatement
			setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", taskinputshape.GongGetIdentifier(stage))
			setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "StartOrientation")
			setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", "models."+taskinputshape.StartOrientation.ToCodeString())
			initializerStatements += setValueField
		}
	case "EndOrientation":
		if taskinputshape.EndOrientation != "" {
			setValueField = StringEnumInitStatement
			setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", taskinputshape.GongGetIdentifier(stage))
			setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "EndOrientation")
			setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", "models."+taskinputshape.EndOrientation.ToCodeString())
			initializerStatements += setValueField
		}
	case "CornerOffsetRatio":
		setValueField = NumberInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", taskinputshape.GongGetIdentifier(stage))
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "CornerOffsetRatio")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%f", taskinputshape.CornerOffsetRatio))
		initializerStatements += setValueField

	case "Task":
		if taskinputshape.Task != nil {
			setPointerField = PointerFieldInitStatement
			setPointerField = strings.ReplaceAll(setPointerField, "{{Identifier}}", taskinputshape.GongGetIdentifier(stage))
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldName}}", "Task")
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldNameValue}}", taskinputshape.Task.GongGetIdentifier(stage))
			pointersInitializesStatements += setPointerField
		}
	case "Product":
		if taskinputshape.Product != nil {
			setPointerField = PointerFieldInitStatement
			setPointerField = strings.ReplaceAll(setPointerField, "{{Identifier}}", taskinputshape.GongGetIdentifier(stage))
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldName}}", "Product")
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldNameValue}}", taskinputshape.Product.GongGetIdentifier(stage))
			pointersInitializesStatements += setPointerField
		}

	default:
		log.Panicf("Unknown field %s for Gongstruct TaskInputShape", fieldName)
	}
	return
}

func (taskoutputshape *TaskOutputShape) GongMarshallField(stage *Stage, fieldName string) (setValueField, setPointerField string) {
	initializerStatements := ""
	_ = initializerStatements
	pointersInitializesStatements := ""
	_ = pointersInitializesStatements

	switch fieldName {
	case "Name":
		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", taskoutputshape.GongGetIdentifier(stage))
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Name")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(taskoutputshape.Name))
		initializerStatements += setValueField
	case "StartRatio":
		setValueField = NumberInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", taskoutputshape.GongGetIdentifier(stage))
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "StartRatio")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%f", taskoutputshape.StartRatio))
		initializerStatements += setValueField
	case "EndRatio":
		setValueField = NumberInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", taskoutputshape.GongGetIdentifier(stage))
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "EndRatio")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%f", taskoutputshape.EndRatio))
		initializerStatements += setValueField
	case "StartOrientation":
		if taskoutputshape.StartOrientation != "" {
			setValueField = StringEnumInitStatement
			setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", taskoutputshape.GongGetIdentifier(stage))
			setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "StartOrientation")
			setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", "models."+taskoutputshape.StartOrientation.ToCodeString())
			initializerStatements += setValueField
		}
	case "EndOrientation":
		if taskoutputshape.EndOrientation != "" {
			setValueField = StringEnumInitStatement
			setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", taskoutputshape.GongGetIdentifier(stage))
			setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "EndOrientation")
			setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", "models."+taskoutputshape.EndOrientation.ToCodeString())
			initializerStatements += setValueField
		}
	case "CornerOffsetRatio":
		setValueField = NumberInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", taskoutputshape.GongGetIdentifier(stage))
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "CornerOffsetRatio")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%f", taskoutputshape.CornerOffsetRatio))
		initializerStatements += setValueField

	case "Task":
		if taskoutputshape.Task != nil {
			setPointerField = PointerFieldInitStatement
			setPointerField = strings.ReplaceAll(setPointerField, "{{Identifier}}", taskoutputshape.GongGetIdentifier(stage))
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldName}}", "Task")
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldNameValue}}", taskoutputshape.Task.GongGetIdentifier(stage))
			pointersInitializesStatements += setPointerField
		}
	case "Product":
		if taskoutputshape.Product != nil {
			setPointerField = PointerFieldInitStatement
			setPointerField = strings.ReplaceAll(setPointerField, "{{Identifier}}", taskoutputshape.GongGetIdentifier(stage))
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldName}}", "Product")
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldNameValue}}", taskoutputshape.Product.GongGetIdentifier(stage))
			pointersInitializesStatements += setPointerField
		}

	default:
		log.Panicf("Unknown field %s for Gongstruct TaskOutputShape", fieldName)
	}
	return
}

func (taskshape *TaskShape) GongMarshallField(stage *Stage, fieldName string) (setValueField, setPointerField string) {
	initializerStatements := ""
	_ = initializerStatements
	pointersInitializesStatements := ""
	_ = pointersInitializesStatements

	switch fieldName {
	case "Name":
		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", taskshape.GongGetIdentifier(stage))
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Name")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(taskshape.Name))
		initializerStatements += setValueField
	case "IsExpanded":
		setValueField = NumberInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", taskshape.GongGetIdentifier(stage))
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "IsExpanded")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%t", taskshape.IsExpanded))
		initializerStatements += setValueField
	case "X":
		setValueField = NumberInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", taskshape.GongGetIdentifier(stage))
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "X")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%f", taskshape.X))
		initializerStatements += setValueField
	case "Y":
		setValueField = NumberInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", taskshape.GongGetIdentifier(stage))
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Y")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%f", taskshape.Y))
		initializerStatements += setValueField
	case "Width":
		setValueField = NumberInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", taskshape.GongGetIdentifier(stage))
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Width")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%f", taskshape.Width))
		initializerStatements += setValueField
	case "Height":
		setValueField = NumberInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", taskshape.GongGetIdentifier(stage))
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Height")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%f", taskshape.Height))
		initializerStatements += setValueField

	case "Task":
		if taskshape.Task != nil {
			setPointerField = PointerFieldInitStatement
			setPointerField = strings.ReplaceAll(setPointerField, "{{Identifier}}", taskshape.GongGetIdentifier(stage))
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldName}}", "Task")
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldNameValue}}", taskshape.Task.GongGetIdentifier(stage))
			pointersInitializesStatements += setPointerField
		}

	default:
		log.Panicf("Unknown field %s for Gongstruct TaskShape", fieldName)
	}
	return
}
