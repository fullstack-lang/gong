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
	map_Category1_Identifiers := make(map[*Category1]string)
	_ = map_Category1_Identifiers

	category1Ordered := []*Category1{}
	for category1 := range stage.Category1s {
		category1Ordered = append(category1Ordered, category1)
	}
	sort.Slice(category1Ordered[:], func(i, j int) bool {
		category1i := category1Ordered[i]
		category1j := category1Ordered[j]
		category1i_order, oki := stage.Category1Map_Staged_Order[category1i]
		category1j_order, okj := stage.Category1Map_Staged_Order[category1j]
		if !oki || !okj {
			log.Fatalln("unknown pointers")
		}
		return category1i_order < category1j_order
	})
	if len(category1Ordered) > 0 {
		identifiersDecl += "\n"
	}
	for _, category1 := range category1Ordered {

		id = generatesIdentifier("Category1", int(stage.Category1Map_Staged_Order[category1]), category1.Name)
		map_Category1_Identifiers[category1] = id

		decl = IdentifiersDecls
		decl = strings.ReplaceAll(decl, "{{Identifier}}", id)
		decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "Category1")
		decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", category1.Name)
		identifiersDecl += decl

		initializerStatements += "\n"
		// Initialisation of values
		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Name")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(category1.Name))
		initializerStatements += setValueField

	}

	map_Category1Shape_Identifiers := make(map[*Category1Shape]string)
	_ = map_Category1Shape_Identifiers

	category1shapeOrdered := []*Category1Shape{}
	for category1shape := range stage.Category1Shapes {
		category1shapeOrdered = append(category1shapeOrdered, category1shape)
	}
	sort.Slice(category1shapeOrdered[:], func(i, j int) bool {
		category1shapei := category1shapeOrdered[i]
		category1shapej := category1shapeOrdered[j]
		category1shapei_order, oki := stage.Category1ShapeMap_Staged_Order[category1shapei]
		category1shapej_order, okj := stage.Category1ShapeMap_Staged_Order[category1shapej]
		if !oki || !okj {
			log.Fatalln("unknown pointers")
		}
		return category1shapei_order < category1shapej_order
	})
	if len(category1shapeOrdered) > 0 {
		identifiersDecl += "\n"
	}
	for _, category1shape := range category1shapeOrdered {

		id = generatesIdentifier("Category1Shape", int(stage.Category1ShapeMap_Staged_Order[category1shape]), category1shape.Name)
		map_Category1Shape_Identifiers[category1shape] = id

		decl = IdentifiersDecls
		decl = strings.ReplaceAll(decl, "{{Identifier}}", id)
		decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "Category1Shape")
		decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", category1shape.Name)
		identifiersDecl += decl

		initializerStatements += "\n"
		// Initialisation of values
		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Name")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(category1shape.Name))
		initializerStatements += setValueField

		setValueField = NumberInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "X")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%f", category1shape.X))
		initializerStatements += setValueField

		setValueField = NumberInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Y")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%f", category1shape.Y))
		initializerStatements += setValueField

		setValueField = NumberInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Width")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%f", category1shape.Width))
		initializerStatements += setValueField

		setValueField = NumberInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Height")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%f", category1shape.Height))
		initializerStatements += setValueField

	}

	map_Category2_Identifiers := make(map[*Category2]string)
	_ = map_Category2_Identifiers

	category2Ordered := []*Category2{}
	for category2 := range stage.Category2s {
		category2Ordered = append(category2Ordered, category2)
	}
	sort.Slice(category2Ordered[:], func(i, j int) bool {
		category2i := category2Ordered[i]
		category2j := category2Ordered[j]
		category2i_order, oki := stage.Category2Map_Staged_Order[category2i]
		category2j_order, okj := stage.Category2Map_Staged_Order[category2j]
		if !oki || !okj {
			log.Fatalln("unknown pointers")
		}
		return category2i_order < category2j_order
	})
	if len(category2Ordered) > 0 {
		identifiersDecl += "\n"
	}
	for _, category2 := range category2Ordered {

		id = generatesIdentifier("Category2", int(stage.Category2Map_Staged_Order[category2]), category2.Name)
		map_Category2_Identifiers[category2] = id

		decl = IdentifiersDecls
		decl = strings.ReplaceAll(decl, "{{Identifier}}", id)
		decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "Category2")
		decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", category2.Name)
		identifiersDecl += decl

		initializerStatements += "\n"
		// Initialisation of values
		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Name")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(category2.Name))
		initializerStatements += setValueField

	}

	map_Category2Shape_Identifiers := make(map[*Category2Shape]string)
	_ = map_Category2Shape_Identifiers

	category2shapeOrdered := []*Category2Shape{}
	for category2shape := range stage.Category2Shapes {
		category2shapeOrdered = append(category2shapeOrdered, category2shape)
	}
	sort.Slice(category2shapeOrdered[:], func(i, j int) bool {
		category2shapei := category2shapeOrdered[i]
		category2shapej := category2shapeOrdered[j]
		category2shapei_order, oki := stage.Category2ShapeMap_Staged_Order[category2shapei]
		category2shapej_order, okj := stage.Category2ShapeMap_Staged_Order[category2shapej]
		if !oki || !okj {
			log.Fatalln("unknown pointers")
		}
		return category2shapei_order < category2shapej_order
	})
	if len(category2shapeOrdered) > 0 {
		identifiersDecl += "\n"
	}
	for _, category2shape := range category2shapeOrdered {

		id = generatesIdentifier("Category2Shape", int(stage.Category2ShapeMap_Staged_Order[category2shape]), category2shape.Name)
		map_Category2Shape_Identifiers[category2shape] = id

		decl = IdentifiersDecls
		decl = strings.ReplaceAll(decl, "{{Identifier}}", id)
		decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "Category2Shape")
		decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", category2shape.Name)
		identifiersDecl += decl

		initializerStatements += "\n"
		// Initialisation of values
		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Name")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(category2shape.Name))
		initializerStatements += setValueField

		setValueField = NumberInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "X")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%f", category2shape.X))
		initializerStatements += setValueField

		setValueField = NumberInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Y")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%f", category2shape.Y))
		initializerStatements += setValueField

		setValueField = NumberInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Width")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%f", category2shape.Width))
		initializerStatements += setValueField

		setValueField = NumberInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Height")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%f", category2shape.Height))
		initializerStatements += setValueField

	}

	map_Category3_Identifiers := make(map[*Category3]string)
	_ = map_Category3_Identifiers

	category3Ordered := []*Category3{}
	for category3 := range stage.Category3s {
		category3Ordered = append(category3Ordered, category3)
	}
	sort.Slice(category3Ordered[:], func(i, j int) bool {
		category3i := category3Ordered[i]
		category3j := category3Ordered[j]
		category3i_order, oki := stage.Category3Map_Staged_Order[category3i]
		category3j_order, okj := stage.Category3Map_Staged_Order[category3j]
		if !oki || !okj {
			log.Fatalln("unknown pointers")
		}
		return category3i_order < category3j_order
	})
	if len(category3Ordered) > 0 {
		identifiersDecl += "\n"
	}
	for _, category3 := range category3Ordered {

		id = generatesIdentifier("Category3", int(stage.Category3Map_Staged_Order[category3]), category3.Name)
		map_Category3_Identifiers[category3] = id

		decl = IdentifiersDecls
		decl = strings.ReplaceAll(decl, "{{Identifier}}", id)
		decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "Category3")
		decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", category3.Name)
		identifiersDecl += decl

		initializerStatements += "\n"
		// Initialisation of values
		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Name")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(category3.Name))
		initializerStatements += setValueField

	}

	map_Category3Shape_Identifiers := make(map[*Category3Shape]string)
	_ = map_Category3Shape_Identifiers

	category3shapeOrdered := []*Category3Shape{}
	for category3shape := range stage.Category3Shapes {
		category3shapeOrdered = append(category3shapeOrdered, category3shape)
	}
	sort.Slice(category3shapeOrdered[:], func(i, j int) bool {
		category3shapei := category3shapeOrdered[i]
		category3shapej := category3shapeOrdered[j]
		category3shapei_order, oki := stage.Category3ShapeMap_Staged_Order[category3shapei]
		category3shapej_order, okj := stage.Category3ShapeMap_Staged_Order[category3shapej]
		if !oki || !okj {
			log.Fatalln("unknown pointers")
		}
		return category3shapei_order < category3shapej_order
	})
	if len(category3shapeOrdered) > 0 {
		identifiersDecl += "\n"
	}
	for _, category3shape := range category3shapeOrdered {

		id = generatesIdentifier("Category3Shape", int(stage.Category3ShapeMap_Staged_Order[category3shape]), category3shape.Name)
		map_Category3Shape_Identifiers[category3shape] = id

		decl = IdentifiersDecls
		decl = strings.ReplaceAll(decl, "{{Identifier}}", id)
		decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "Category3Shape")
		decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", category3shape.Name)
		identifiersDecl += decl

		initializerStatements += "\n"
		// Initialisation of values
		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Name")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(category3shape.Name))
		initializerStatements += setValueField

		setValueField = NumberInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "X")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%f", category3shape.X))
		initializerStatements += setValueField

		setValueField = NumberInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Y")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%f", category3shape.Y))
		initializerStatements += setValueField

		setValueField = NumberInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Width")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%f", category3shape.Width))
		initializerStatements += setValueField

		setValueField = NumberInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Height")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%f", category3shape.Height))
		initializerStatements += setValueField

	}

	map_ControlPointShape_Identifiers := make(map[*ControlPointShape]string)
	_ = map_ControlPointShape_Identifiers

	controlpointshapeOrdered := []*ControlPointShape{}
	for controlpointshape := range stage.ControlPointShapes {
		controlpointshapeOrdered = append(controlpointshapeOrdered, controlpointshape)
	}
	sort.Slice(controlpointshapeOrdered[:], func(i, j int) bool {
		controlpointshapei := controlpointshapeOrdered[i]
		controlpointshapej := controlpointshapeOrdered[j]
		controlpointshapei_order, oki := stage.ControlPointShapeMap_Staged_Order[controlpointshapei]
		controlpointshapej_order, okj := stage.ControlPointShapeMap_Staged_Order[controlpointshapej]
		if !oki || !okj {
			log.Fatalln("unknown pointers")
		}
		return controlpointshapei_order < controlpointshapej_order
	})
	if len(controlpointshapeOrdered) > 0 {
		identifiersDecl += "\n"
	}
	for _, controlpointshape := range controlpointshapeOrdered {

		id = generatesIdentifier("ControlPointShape", int(stage.ControlPointShapeMap_Staged_Order[controlpointshape]), controlpointshape.Name)
		map_ControlPointShape_Identifiers[controlpointshape] = id

		decl = IdentifiersDecls
		decl = strings.ReplaceAll(decl, "{{Identifier}}", id)
		decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "ControlPointShape")
		decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", controlpointshape.Name)
		identifiersDecl += decl

		initializerStatements += "\n"
		// Initialisation of values
		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Name")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(controlpointshape.Name))
		initializerStatements += setValueField

		setValueField = NumberInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "X_Relative")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%f", controlpointshape.X_Relative))
		initializerStatements += setValueField

		setValueField = NumberInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Y_Relative")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%f", controlpointshape.Y_Relative))
		initializerStatements += setValueField

		setValueField = NumberInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "IsStartShapeTheClosestShape")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%t", controlpointshape.IsStartShapeTheClosestShape))
		initializerStatements += setValueField

	}

	map_Desk_Identifiers := make(map[*Desk]string)
	_ = map_Desk_Identifiers

	deskOrdered := []*Desk{}
	for desk := range stage.Desks {
		deskOrdered = append(deskOrdered, desk)
	}
	sort.Slice(deskOrdered[:], func(i, j int) bool {
		deski := deskOrdered[i]
		deskj := deskOrdered[j]
		deski_order, oki := stage.DeskMap_Staged_Order[deski]
		deskj_order, okj := stage.DeskMap_Staged_Order[deskj]
		if !oki || !okj {
			log.Fatalln("unknown pointers")
		}
		return deski_order < deskj_order
	})
	if len(deskOrdered) > 0 {
		identifiersDecl += "\n"
	}
	for _, desk := range deskOrdered {

		id = generatesIdentifier("Desk", int(stage.DeskMap_Staged_Order[desk]), desk.Name)
		map_Desk_Identifiers[desk] = id

		decl = IdentifiersDecls
		decl = strings.ReplaceAll(decl, "{{Identifier}}", id)
		decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "Desk")
		decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", desk.Name)
		identifiersDecl += decl

		initializerStatements += "\n"
		// Initialisation of values
		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Name")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(desk.Name))
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
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "IsEditable")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%t", diagram.IsEditable))
		initializerStatements += setValueField

		setValueField = NumberInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "IsNodeExpanded")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%t", diagram.IsNodeExpanded))
		initializerStatements += setValueField

		setValueField = NumberInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "IsCategory1NodeExpanded")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%t", diagram.IsCategory1NodeExpanded))
		initializerStatements += setValueField

		setValueField = NumberInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "IsCategory2NodeExpanded")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%t", diagram.IsCategory2NodeExpanded))
		initializerStatements += setValueField

		setValueField = NumberInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "IsCategory3NodeExpanded")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%t", diagram.IsCategory3NodeExpanded))
		initializerStatements += setValueField

		setValueField = NumberInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "IsInfluenceCategoryNodeExpanded")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%t", diagram.IsInfluenceCategoryNodeExpanded))
		initializerStatements += setValueField

		setValueField = NumberInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "IsCategory1Shown")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%t", diagram.IsCategory1Shown))
		initializerStatements += setValueField

		setValueField = NumberInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "IsCategory2Shown")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%t", diagram.IsCategory2Shown))
		initializerStatements += setValueField

		setValueField = NumberInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "IsCategory3Shown")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%t", diagram.IsCategory3Shown))
		initializerStatements += setValueField

		setValueField = NumberInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "IsInfluenceCategoryShown")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%t", diagram.IsInfluenceCategoryShown))
		initializerStatements += setValueField

		setValueField = NumberInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "XMargin")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%f", diagram.XMargin))
		initializerStatements += setValueField

		setValueField = NumberInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "YMargin")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%f", diagram.YMargin))
		initializerStatements += setValueField

		setValueField = NumberInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Height")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%f", diagram.Height))
		initializerStatements += setValueField

		setValueField = NumberInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Width")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%f", diagram.Width))
		initializerStatements += setValueField

		setValueField = NumberInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "NbPixPerCharacter")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%f", diagram.NbPixPerCharacter))
		initializerStatements += setValueField

		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "RedColorCode")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(diagram.RedColorCode))
		initializerStatements += setValueField

		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "BackgroundGreyColorCode")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(diagram.BackgroundGreyColorCode))
		initializerStatements += setValueField

		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "GrayColorCode")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(diagram.GrayColorCode))
		initializerStatements += setValueField

		if diagram.Category1RectAnchorType != "" {
			setValueField = StringEnumInitStatement
			setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
			setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Category1RectAnchorType")
			setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", "models."+diagram.Category1RectAnchorType.ToCodeString())
			initializerStatements += setValueField
		}

		if diagram.Category1TextAnchorType != "" {
			setValueField = StringEnumInitStatement
			setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
			setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Category1TextAnchorType")
			setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", "models."+diagram.Category1TextAnchorType.ToCodeString())
			initializerStatements += setValueField
		}

		if diagram.Category1DominantBaselineType != "" {
			setValueField = StringEnumInitStatement
			setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
			setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Category1DominantBaselineType")
			setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", "models."+diagram.Category1DominantBaselineType.ToCodeString())
			initializerStatements += setValueField
		}

		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Category1FontSize")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(diagram.Category1FontSize))
		initializerStatements += setValueField

		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Category1FontWeigth")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(diagram.Category1FontWeigth))
		initializerStatements += setValueField

		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Category1FontFamily")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(diagram.Category1FontFamily))
		initializerStatements += setValueField

		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Category1LetterSpacing")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(diagram.Category1LetterSpacing))
		initializerStatements += setValueField

		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Category2TypeFontSize")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(diagram.Category2TypeFontSize))
		initializerStatements += setValueField

		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Category2TypeFontWeigth")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(diagram.Category2TypeFontWeigth))
		initializerStatements += setValueField

		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Category2TypeFontFamily")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(diagram.Category2TypeFontFamily))
		initializerStatements += setValueField

		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Category2TypeLetterSpacing")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(diagram.Category2TypeLetterSpacing))
		initializerStatements += setValueField

		if diagram.Category2TypeRectAnchorType != "" {
			setValueField = StringEnumInitStatement
			setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
			setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Category2TypeRectAnchorType")
			setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", "models."+diagram.Category2TypeRectAnchorType.ToCodeString())
			initializerStatements += setValueField
		}

		if diagram.Category2DominantBaselineType != "" {
			setValueField = StringEnumInitStatement
			setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
			setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Category2DominantBaselineType")
			setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", "models."+diagram.Category2DominantBaselineType.ToCodeString())
			initializerStatements += setValueField
		}

		setValueField = NumberInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Category2StrokeWidth")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%f", diagram.Category2StrokeWidth))
		initializerStatements += setValueField

		if diagram.Category3RectAnchorType != "" {
			setValueField = StringEnumInitStatement
			setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
			setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Category3RectAnchorType")
			setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", "models."+diagram.Category3RectAnchorType.ToCodeString())
			initializerStatements += setValueField
		}

		if diagram.Category3TextAnchorType != "" {
			setValueField = StringEnumInitStatement
			setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
			setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Category3TextAnchorType")
			setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", "models."+diagram.Category3TextAnchorType.ToCodeString())
			initializerStatements += setValueField
		}

		if diagram.Category3DominantBaselineType != "" {
			setValueField = StringEnumInitStatement
			setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
			setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Category3DominantBaselineType")
			setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", "models."+diagram.Category3DominantBaselineType.ToCodeString())
			initializerStatements += setValueField
		}

		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Category3FontSize")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(diagram.Category3FontSize))
		initializerStatements += setValueField

		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Category3FontWeigth")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(diagram.Category3FontWeigth))
		initializerStatements += setValueField

		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Category3FontFamily")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(diagram.Category3FontFamily))
		initializerStatements += setValueField

		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Category3LetterSpacing")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(diagram.Category3LetterSpacing))
		initializerStatements += setValueField

		setValueField = NumberInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "InfluenceStrokeWidth")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%f", diagram.InfluenceStrokeWidth))
		initializerStatements += setValueField

		setValueField = NumberInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "InfluenceArrowSize")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%f", diagram.InfluenceArrowSize))
		initializerStatements += setValueField

		setValueField = NumberInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "InfluenceArrowStartOffset")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%f", diagram.InfluenceArrowStartOffset))
		initializerStatements += setValueField

		setValueField = NumberInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "InfluenceArrowEndOffset")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%f", diagram.InfluenceArrowEndOffset))
		initializerStatements += setValueField

		setValueField = NumberInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "InfluenceCornerRadius")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%f", diagram.InfluenceCornerRadius))
		initializerStatements += setValueField

		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "InfluenceFontSize")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(diagram.InfluenceFontSize))
		initializerStatements += setValueField

		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "InfluenceFontWeigth")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(diagram.InfluenceFontWeigth))
		initializerStatements += setValueField

		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "InfluenceFontFamily")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(diagram.InfluenceFontFamily))
		initializerStatements += setValueField

		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "InfluenceLetterSpacing")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(diagram.InfluenceLetterSpacing))
		initializerStatements += setValueField

		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "InfluenceDashedLinePattern")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(diagram.InfluenceDashedLinePattern))
		initializerStatements += setValueField

	}

	map_Influence_Identifiers := make(map[*Influence]string)
	_ = map_Influence_Identifiers

	influenceOrdered := []*Influence{}
	for influence := range stage.Influences {
		influenceOrdered = append(influenceOrdered, influence)
	}
	sort.Slice(influenceOrdered[:], func(i, j int) bool {
		influencei := influenceOrdered[i]
		influencej := influenceOrdered[j]
		influencei_order, oki := stage.InfluenceMap_Staged_Order[influencei]
		influencej_order, okj := stage.InfluenceMap_Staged_Order[influencej]
		if !oki || !okj {
			log.Fatalln("unknown pointers")
		}
		return influencei_order < influencej_order
	})
	if len(influenceOrdered) > 0 {
		identifiersDecl += "\n"
	}
	for _, influence := range influenceOrdered {

		id = generatesIdentifier("Influence", int(stage.InfluenceMap_Staged_Order[influence]), influence.Name)
		map_Influence_Identifiers[influence] = id

		decl = IdentifiersDecls
		decl = strings.ReplaceAll(decl, "{{Identifier}}", id)
		decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "Influence")
		decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", influence.Name)
		identifiersDecl += decl

		initializerStatements += "\n"
		// Initialisation of values
		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Name")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(influence.Name))
		initializerStatements += setValueField

		setValueField = NumberInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "IsHypothtical")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%t", influence.IsHypothtical))
		initializerStatements += setValueField

		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "TextAtEndOfArrow")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(influence.TextAtEndOfArrow))
		initializerStatements += setValueField

	}

	map_InfluenceShape_Identifiers := make(map[*InfluenceShape]string)
	_ = map_InfluenceShape_Identifiers

	influenceshapeOrdered := []*InfluenceShape{}
	for influenceshape := range stage.InfluenceShapes {
		influenceshapeOrdered = append(influenceshapeOrdered, influenceshape)
	}
	sort.Slice(influenceshapeOrdered[:], func(i, j int) bool {
		influenceshapei := influenceshapeOrdered[i]
		influenceshapej := influenceshapeOrdered[j]
		influenceshapei_order, oki := stage.InfluenceShapeMap_Staged_Order[influenceshapei]
		influenceshapej_order, okj := stage.InfluenceShapeMap_Staged_Order[influenceshapej]
		if !oki || !okj {
			log.Fatalln("unknown pointers")
		}
		return influenceshapei_order < influenceshapej_order
	})
	if len(influenceshapeOrdered) > 0 {
		identifiersDecl += "\n"
	}
	for _, influenceshape := range influenceshapeOrdered {

		id = generatesIdentifier("InfluenceShape", int(stage.InfluenceShapeMap_Staged_Order[influenceshape]), influenceshape.Name)
		map_InfluenceShape_Identifiers[influenceshape] = id

		decl = IdentifiersDecls
		decl = strings.ReplaceAll(decl, "{{Identifier}}", id)
		decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "InfluenceShape")
		decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", influenceshape.Name)
		identifiersDecl += decl

		initializerStatements += "\n"
		// Initialisation of values
		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Name")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(influenceshape.Name))
		initializerStatements += setValueField

	}

	// insertion initialization of objects to stage
	if len(category1Ordered) > 0 {
		pointersInitializesStatements += "\n\t// setup of Category1 instances pointers"
	}
	for idx, category1 := range category1Ordered {
		var setPointerField string
		_ = setPointerField

		id = generatesIdentifier("Category1", idx, category1.Name)
		map_Category1_Identifiers[category1] = id

		// Initialisation of values
	}

	if len(category1shapeOrdered) > 0 {
		pointersInitializesStatements += "\n\t// setup of Category1Shape instances pointers"
	}
	for idx, category1shape := range category1shapeOrdered {
		var setPointerField string
		_ = setPointerField

		id = generatesIdentifier("Category1Shape", idx, category1shape.Name)
		map_Category1Shape_Identifiers[category1shape] = id

		// Initialisation of values
		if category1shape.Category1 != nil {
			setPointerField = PointerFieldInitStatement
			setPointerField = strings.ReplaceAll(setPointerField, "{{Identifier}}", id)
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldName}}", "Category1")
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldNameValue}}", map_Category1_Identifiers[category1shape.Category1])
			pointersInitializesStatements += setPointerField
		}

	}

	if len(category2Ordered) > 0 {
		pointersInitializesStatements += "\n\t// setup of Category2 instances pointers"
	}
	for idx, category2 := range category2Ordered {
		var setPointerField string
		_ = setPointerField

		id = generatesIdentifier("Category2", idx, category2.Name)
		map_Category2_Identifiers[category2] = id

		// Initialisation of values
	}

	if len(category2shapeOrdered) > 0 {
		pointersInitializesStatements += "\n\t// setup of Category2Shape instances pointers"
	}
	for idx, category2shape := range category2shapeOrdered {
		var setPointerField string
		_ = setPointerField

		id = generatesIdentifier("Category2Shape", idx, category2shape.Name)
		map_Category2Shape_Identifiers[category2shape] = id

		// Initialisation of values
		if category2shape.Category2 != nil {
			setPointerField = PointerFieldInitStatement
			setPointerField = strings.ReplaceAll(setPointerField, "{{Identifier}}", id)
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldName}}", "Category2")
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldNameValue}}", map_Category2_Identifiers[category2shape.Category2])
			pointersInitializesStatements += setPointerField
		}

	}

	if len(category3Ordered) > 0 {
		pointersInitializesStatements += "\n\t// setup of Category3 instances pointers"
	}
	for idx, category3 := range category3Ordered {
		var setPointerField string
		_ = setPointerField

		id = generatesIdentifier("Category3", idx, category3.Name)
		map_Category3_Identifiers[category3] = id

		// Initialisation of values
	}

	if len(category3shapeOrdered) > 0 {
		pointersInitializesStatements += "\n\t// setup of Category3Shape instances pointers"
	}
	for idx, category3shape := range category3shapeOrdered {
		var setPointerField string
		_ = setPointerField

		id = generatesIdentifier("Category3Shape", idx, category3shape.Name)
		map_Category3Shape_Identifiers[category3shape] = id

		// Initialisation of values
		if category3shape.Category3 != nil {
			setPointerField = PointerFieldInitStatement
			setPointerField = strings.ReplaceAll(setPointerField, "{{Identifier}}", id)
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldName}}", "Category3")
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldNameValue}}", map_Category3_Identifiers[category3shape.Category3])
			pointersInitializesStatements += setPointerField
		}

	}

	if len(controlpointshapeOrdered) > 0 {
		pointersInitializesStatements += "\n\t// setup of ControlPointShape instances pointers"
	}
	for idx, controlpointshape := range controlpointshapeOrdered {
		var setPointerField string
		_ = setPointerField

		id = generatesIdentifier("ControlPointShape", idx, controlpointshape.Name)
		map_ControlPointShape_Identifiers[controlpointshape] = id

		// Initialisation of values
	}

	if len(deskOrdered) > 0 {
		pointersInitializesStatements += "\n\t// setup of Desk instances pointers"
	}
	for idx, desk := range deskOrdered {
		var setPointerField string
		_ = setPointerField

		id = generatesIdentifier("Desk", idx, desk.Name)
		map_Desk_Identifiers[desk] = id

		// Initialisation of values
		if desk.SelectedDiagram != nil {
			setPointerField = PointerFieldInitStatement
			setPointerField = strings.ReplaceAll(setPointerField, "{{Identifier}}", id)
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldName}}", "SelectedDiagram")
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldNameValue}}", map_Diagram_Identifiers[desk.SelectedDiagram])
			pointersInitializesStatements += setPointerField
		}

	}

	if len(diagramOrdered) > 0 {
		pointersInitializesStatements += "\n\t// setup of Diagram instances pointers"
	}
	for idx, diagram := range diagramOrdered {
		var setPointerField string
		_ = setPointerField

		id = generatesIdentifier("Diagram", idx, diagram.Name)
		map_Diagram_Identifiers[diagram] = id

		// Initialisation of values
		for _, _category1shape := range diagram.Category1Shapes {
			setPointerField = SliceOfPointersFieldInitStatement
			setPointerField = strings.ReplaceAll(setPointerField, "{{Identifier}}", id)
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldName}}", "Category1Shapes")
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldNameValue}}", map_Category1Shape_Identifiers[_category1shape])
			pointersInitializesStatements += setPointerField
		}

		for _, _category2shape := range diagram.Category2Shapes {
			setPointerField = SliceOfPointersFieldInitStatement
			setPointerField = strings.ReplaceAll(setPointerField, "{{Identifier}}", id)
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldName}}", "Category2Shapes")
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldNameValue}}", map_Category2Shape_Identifiers[_category2shape])
			pointersInitializesStatements += setPointerField
		}

		for _, _category3shape := range diagram.Category3Shapes {
			setPointerField = SliceOfPointersFieldInitStatement
			setPointerField = strings.ReplaceAll(setPointerField, "{{Identifier}}", id)
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldName}}", "Category3Shapes")
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldNameValue}}", map_Category3Shape_Identifiers[_category3shape])
			pointersInitializesStatements += setPointerField
		}

		for _, _influenceshape := range diagram.InfluenceShapes {
			setPointerField = SliceOfPointersFieldInitStatement
			setPointerField = strings.ReplaceAll(setPointerField, "{{Identifier}}", id)
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldName}}", "InfluenceShapes")
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldNameValue}}", map_InfluenceShape_Identifiers[_influenceshape])
			pointersInitializesStatements += setPointerField
		}

	}

	if len(influenceOrdered) > 0 {
		pointersInitializesStatements += "\n\t// setup of Influence instances pointers"
	}
	for idx, influence := range influenceOrdered {
		var setPointerField string
		_ = setPointerField

		id = generatesIdentifier("Influence", idx, influence.Name)
		map_Influence_Identifiers[influence] = id

		// Initialisation of values
		if influence.SourceCategory1 != nil {
			setPointerField = PointerFieldInitStatement
			setPointerField = strings.ReplaceAll(setPointerField, "{{Identifier}}", id)
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldName}}", "SourceCategory1")
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldNameValue}}", map_Category1_Identifiers[influence.SourceCategory1])
			pointersInitializesStatements += setPointerField
		}

		if influence.SourceCategory2 != nil {
			setPointerField = PointerFieldInitStatement
			setPointerField = strings.ReplaceAll(setPointerField, "{{Identifier}}", id)
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldName}}", "SourceCategory2")
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldNameValue}}", map_Category3_Identifiers[influence.SourceCategory2])
			pointersInitializesStatements += setPointerField
		}

		if influence.SourceCategory3 != nil {
			setPointerField = PointerFieldInitStatement
			setPointerField = strings.ReplaceAll(setPointerField, "{{Identifier}}", id)
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldName}}", "SourceCategory3")
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldNameValue}}", map_Category2_Identifiers[influence.SourceCategory3])
			pointersInitializesStatements += setPointerField
		}

		if influence.TargetCategory1 != nil {
			setPointerField = PointerFieldInitStatement
			setPointerField = strings.ReplaceAll(setPointerField, "{{Identifier}}", id)
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldName}}", "TargetCategory1")
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldNameValue}}", map_Category1_Identifiers[influence.TargetCategory1])
			pointersInitializesStatements += setPointerField
		}

		if influence.TargetCategory2 != nil {
			setPointerField = PointerFieldInitStatement
			setPointerField = strings.ReplaceAll(setPointerField, "{{Identifier}}", id)
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldName}}", "TargetCategory2")
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldNameValue}}", map_Category3_Identifiers[influence.TargetCategory2])
			pointersInitializesStatements += setPointerField
		}

		if influence.TargetCategory3 != nil {
			setPointerField = PointerFieldInitStatement
			setPointerField = strings.ReplaceAll(setPointerField, "{{Identifier}}", id)
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldName}}", "TargetCategory3")
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldNameValue}}", map_Category2_Identifiers[influence.TargetCategory3])
			pointersInitializesStatements += setPointerField
		}

	}

	if len(influenceshapeOrdered) > 0 {
		pointersInitializesStatements += "\n\t// setup of InfluenceShape instances pointers"
	}
	for idx, influenceshape := range influenceshapeOrdered {
		var setPointerField string
		_ = setPointerField

		id = generatesIdentifier("InfluenceShape", idx, influenceshape.Name)
		map_InfluenceShape_Identifiers[influenceshape] = id

		// Initialisation of values
		if influenceshape.Influence != nil {
			setPointerField = PointerFieldInitStatement
			setPointerField = strings.ReplaceAll(setPointerField, "{{Identifier}}", id)
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldName}}", "Influence")
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldNameValue}}", map_Influence_Identifiers[influenceshape.Influence])
			pointersInitializesStatements += setPointerField
		}

		for _, _controlpointshape := range influenceshape.ControlPointShapes {
			setPointerField = SliceOfPointersFieldInitStatement
			setPointerField = strings.ReplaceAll(setPointerField, "{{Identifier}}", id)
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldName}}", "ControlPointShapes")
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldNameValue}}", map_ControlPointShape_Identifiers[_controlpointshape])
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
