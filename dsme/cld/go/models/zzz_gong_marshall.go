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

		identifiersDecl += category1.GongMarshallIdentifier(stage)

		initializerStatements += "\n"
		// Insertion point for basic fields value assignment
		initializerStatements += category1.GongMarshallField(stage, "Name")
	}

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

		identifiersDecl += category1shape.GongMarshallIdentifier(stage)

		initializerStatements += "\n"
		// Insertion point for basic fields value assignment
		initializerStatements += category1shape.GongMarshallField(stage, "Name")
		pointersInitializesStatements += category1shape.GongMarshallField(stage, "Category1")
		initializerStatements += category1shape.GongMarshallField(stage, "X")
		initializerStatements += category1shape.GongMarshallField(stage, "Y")
		initializerStatements += category1shape.GongMarshallField(stage, "Width")
		initializerStatements += category1shape.GongMarshallField(stage, "Height")
	}

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

		identifiersDecl += category2.GongMarshallIdentifier(stage)

		initializerStatements += "\n"
		// Insertion point for basic fields value assignment
		initializerStatements += category2.GongMarshallField(stage, "Name")
	}

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

		identifiersDecl += category2shape.GongMarshallIdentifier(stage)

		initializerStatements += "\n"
		// Insertion point for basic fields value assignment
		initializerStatements += category2shape.GongMarshallField(stage, "Name")
		pointersInitializesStatements += category2shape.GongMarshallField(stage, "Category2")
		initializerStatements += category2shape.GongMarshallField(stage, "X")
		initializerStatements += category2shape.GongMarshallField(stage, "Y")
		initializerStatements += category2shape.GongMarshallField(stage, "Width")
		initializerStatements += category2shape.GongMarshallField(stage, "Height")
	}

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

		identifiersDecl += category3.GongMarshallIdentifier(stage)

		initializerStatements += "\n"
		// Insertion point for basic fields value assignment
		initializerStatements += category3.GongMarshallField(stage, "Name")
	}

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

		identifiersDecl += category3shape.GongMarshallIdentifier(stage)

		initializerStatements += "\n"
		// Insertion point for basic fields value assignment
		initializerStatements += category3shape.GongMarshallField(stage, "Name")
		pointersInitializesStatements += category3shape.GongMarshallField(stage, "Category3")
		initializerStatements += category3shape.GongMarshallField(stage, "X")
		initializerStatements += category3shape.GongMarshallField(stage, "Y")
		initializerStatements += category3shape.GongMarshallField(stage, "Width")
		initializerStatements += category3shape.GongMarshallField(stage, "Height")
	}

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

		identifiersDecl += controlpointshape.GongMarshallIdentifier(stage)

		initializerStatements += "\n"
		// Insertion point for basic fields value assignment
		initializerStatements += controlpointshape.GongMarshallField(stage, "Name")
		initializerStatements += controlpointshape.GongMarshallField(stage, "X_Relative")
		initializerStatements += controlpointshape.GongMarshallField(stage, "Y_Relative")
		initializerStatements += controlpointshape.GongMarshallField(stage, "IsStartShapeTheClosestShape")
	}

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

		identifiersDecl += desk.GongMarshallIdentifier(stage)

		initializerStatements += "\n"
		// Insertion point for basic fields value assignment
		initializerStatements += desk.GongMarshallField(stage, "Name")
		pointersInitializesStatements += desk.GongMarshallField(stage, "SelectedDiagram")
	}

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
		pointersInitializesStatements += diagram.GongMarshallField(stage, "Category1Shapes")
		pointersInitializesStatements += diagram.GongMarshallField(stage, "Category2Shapes")
		pointersInitializesStatements += diagram.GongMarshallField(stage, "Category3Shapes")
		pointersInitializesStatements += diagram.GongMarshallField(stage, "InfluenceShapes")
		initializerStatements += diagram.GongMarshallField(stage, "IsEditable")
		initializerStatements += diagram.GongMarshallField(stage, "IsNodeExpanded")
		initializerStatements += diagram.GongMarshallField(stage, "IsCategory1NodeExpanded")
		initializerStatements += diagram.GongMarshallField(stage, "IsCategory2NodeExpanded")
		initializerStatements += diagram.GongMarshallField(stage, "IsCategory3NodeExpanded")
		initializerStatements += diagram.GongMarshallField(stage, "IsInfluenceCategoryNodeExpanded")
		initializerStatements += diagram.GongMarshallField(stage, "IsCategory1Shown")
		initializerStatements += diagram.GongMarshallField(stage, "IsCategory2Shown")
		initializerStatements += diagram.GongMarshallField(stage, "IsCategory3Shown")
		initializerStatements += diagram.GongMarshallField(stage, "IsInfluenceCategoryShown")
		initializerStatements += diagram.GongMarshallField(stage, "XMargin")
		initializerStatements += diagram.GongMarshallField(stage, "YMargin")
		initializerStatements += diagram.GongMarshallField(stage, "Height")
		initializerStatements += diagram.GongMarshallField(stage, "Width")
		initializerStatements += diagram.GongMarshallField(stage, "NbPixPerCharacter")
		initializerStatements += diagram.GongMarshallField(stage, "RedColorCode")
		initializerStatements += diagram.GongMarshallField(stage, "BackgroundGreyColorCode")
		initializerStatements += diagram.GongMarshallField(stage, "GrayColorCode")
		initializerStatements += diagram.GongMarshallField(stage, "Category1RectAnchorType")
		initializerStatements += diagram.GongMarshallField(stage, "Category1TextAnchorType")
		initializerStatements += diagram.GongMarshallField(stage, "Category1DominantBaselineType")
		initializerStatements += diagram.GongMarshallField(stage, "Category1FontSize")
		initializerStatements += diagram.GongMarshallField(stage, "Category1FontWeigth")
		initializerStatements += diagram.GongMarshallField(stage, "Category1FontFamily")
		initializerStatements += diagram.GongMarshallField(stage, "Category1LetterSpacing")
		initializerStatements += diagram.GongMarshallField(stage, "Category2TypeFontSize")
		initializerStatements += diagram.GongMarshallField(stage, "Category2TypeFontWeigth")
		initializerStatements += diagram.GongMarshallField(stage, "Category2TypeFontFamily")
		initializerStatements += diagram.GongMarshallField(stage, "Category2TypeLetterSpacing")
		initializerStatements += diagram.GongMarshallField(stage, "Category2TypeRectAnchorType")
		initializerStatements += diagram.GongMarshallField(stage, "Category2DominantBaselineType")
		initializerStatements += diagram.GongMarshallField(stage, "Category2StrokeWidth")
		initializerStatements += diagram.GongMarshallField(stage, "Category3RectAnchorType")
		initializerStatements += diagram.GongMarshallField(stage, "Category3TextAnchorType")
		initializerStatements += diagram.GongMarshallField(stage, "Category3DominantBaselineType")
		initializerStatements += diagram.GongMarshallField(stage, "Category3FontSize")
		initializerStatements += diagram.GongMarshallField(stage, "Category3FontWeigth")
		initializerStatements += diagram.GongMarshallField(stage, "Category3FontFamily")
		initializerStatements += diagram.GongMarshallField(stage, "Category3LetterSpacing")
		initializerStatements += diagram.GongMarshallField(stage, "InfluenceStrokeWidth")
		initializerStatements += diagram.GongMarshallField(stage, "InfluenceArrowSize")
		initializerStatements += diagram.GongMarshallField(stage, "InfluenceArrowStartOffset")
		initializerStatements += diagram.GongMarshallField(stage, "InfluenceArrowEndOffset")
		initializerStatements += diagram.GongMarshallField(stage, "InfluenceCornerRadius")
		initializerStatements += diagram.GongMarshallField(stage, "InfluenceFontSize")
		initializerStatements += diagram.GongMarshallField(stage, "InfluenceFontWeigth")
		initializerStatements += diagram.GongMarshallField(stage, "InfluenceFontFamily")
		initializerStatements += diagram.GongMarshallField(stage, "InfluenceLetterSpacing")
		initializerStatements += diagram.GongMarshallField(stage, "InfluenceDashedLinePattern")
	}

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

		identifiersDecl += influence.GongMarshallIdentifier(stage)

		initializerStatements += "\n"
		// Insertion point for basic fields value assignment
		initializerStatements += influence.GongMarshallField(stage, "Name")
		pointersInitializesStatements += influence.GongMarshallField(stage, "SourceCategory1")
		pointersInitializesStatements += influence.GongMarshallField(stage, "SourceCategory2")
		pointersInitializesStatements += influence.GongMarshallField(stage, "SourceCategory3")
		pointersInitializesStatements += influence.GongMarshallField(stage, "TargetCategory1")
		pointersInitializesStatements += influence.GongMarshallField(stage, "TargetCategory2")
		pointersInitializesStatements += influence.GongMarshallField(stage, "TargetCategory3")
		initializerStatements += influence.GongMarshallField(stage, "IsHypothtical")
		initializerStatements += influence.GongMarshallField(stage, "TextAtEndOfArrow")
	}

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

		identifiersDecl += influenceshape.GongMarshallIdentifier(stage)

		initializerStatements += "\n"
		// Insertion point for basic fields value assignment
		initializerStatements += influenceshape.GongMarshallField(stage, "Name")
		pointersInitializesStatements += influenceshape.GongMarshallField(stage, "Influence")
		pointersInitializesStatements += influenceshape.GongMarshallField(stage, "ControlPointShapes")
	}

	// insertion initialization of objects to stage
	for _, category1 := range category1Ordered {
		_ = category1
		var setPointerField string
		_ = setPointerField

		// Insertion point for pointers initialization
	}

	for _, category1shape := range category1shapeOrdered {
		_ = category1shape
		var setPointerField string
		_ = setPointerField

		// Insertion point for pointers initialization
	}

	for _, category2 := range category2Ordered {
		_ = category2
		var setPointerField string
		_ = setPointerField

		// Insertion point for pointers initialization
	}

	for _, category2shape := range category2shapeOrdered {
		_ = category2shape
		var setPointerField string
		_ = setPointerField

		// Insertion point for pointers initialization
	}

	for _, category3 := range category3Ordered {
		_ = category3
		var setPointerField string
		_ = setPointerField

		// Insertion point for pointers initialization
	}

	for _, category3shape := range category3shapeOrdered {
		_ = category3shape
		var setPointerField string
		_ = setPointerField

		// Insertion point for pointers initialization
	}

	for _, controlpointshape := range controlpointshapeOrdered {
		_ = controlpointshape
		var setPointerField string
		_ = setPointerField

		// Insertion point for pointers initialization
	}

	for _, desk := range deskOrdered {
		_ = desk
		var setPointerField string
		_ = setPointerField

		// Insertion point for pointers initialization
	}

	for _, diagram := range diagramOrdered {
		_ = diagram
		var setPointerField string
		_ = setPointerField

		// Insertion point for pointers initialization
	}

	for _, influence := range influenceOrdered {
		_ = influence
		var setPointerField string
		_ = setPointerField

		// Insertion point for pointers initialization
	}

	for _, influenceshape := range influenceshapeOrdered {
		_ = influenceshape
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
func (category1 *Category1) GongMarshallField(stage *Stage, fieldName string) (res string) {

	switch fieldName {
	case "Name":
		res = StringInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", category1.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "Name")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", string(category1.Name))

	default:
		log.Panicf("Unknown field %s for Gongstruct Category1", fieldName)
	}
	return
}

func (category1shape *Category1Shape) GongMarshallField(stage *Stage, fieldName string) (res string) {

	switch fieldName {
	case "Name":
		res = StringInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", category1shape.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "Name")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", string(category1shape.Name))
	case "X":
		res = NumberInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", category1shape.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "X")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%f", category1shape.X))
	case "Y":
		res = NumberInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", category1shape.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "Y")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%f", category1shape.Y))
	case "Width":
		res = NumberInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", category1shape.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "Width")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%f", category1shape.Width))
	case "Height":
		res = NumberInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", category1shape.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "Height")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%f", category1shape.Height))

	case "Category1":
		if category1shape.Category1 != nil {
			res = PointerFieldInitStatement
			res = strings.ReplaceAll(res, "{{Identifier}}", category1shape.GongGetIdentifier(stage))
			res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "Category1")
			res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", category1shape.Category1.GongGetIdentifier(stage))
		} else {
			// in case of nil pointer, we need to unstage the previous value
			res = PointerFieldInitStatement
			res = strings.ReplaceAll(res, "{{Identifier}}", category1shape.GongGetIdentifier(stage))
			res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "Category1")
			res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", "nil")
		}
	default:
		log.Panicf("Unknown field %s for Gongstruct Category1Shape", fieldName)
	}
	return
}

func (category2 *Category2) GongMarshallField(stage *Stage, fieldName string) (res string) {

	switch fieldName {
	case "Name":
		res = StringInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", category2.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "Name")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", string(category2.Name))

	default:
		log.Panicf("Unknown field %s for Gongstruct Category2", fieldName)
	}
	return
}

func (category2shape *Category2Shape) GongMarshallField(stage *Stage, fieldName string) (res string) {

	switch fieldName {
	case "Name":
		res = StringInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", category2shape.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "Name")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", string(category2shape.Name))
	case "X":
		res = NumberInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", category2shape.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "X")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%f", category2shape.X))
	case "Y":
		res = NumberInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", category2shape.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "Y")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%f", category2shape.Y))
	case "Width":
		res = NumberInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", category2shape.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "Width")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%f", category2shape.Width))
	case "Height":
		res = NumberInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", category2shape.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "Height")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%f", category2shape.Height))

	case "Category2":
		if category2shape.Category2 != nil {
			res = PointerFieldInitStatement
			res = strings.ReplaceAll(res, "{{Identifier}}", category2shape.GongGetIdentifier(stage))
			res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "Category2")
			res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", category2shape.Category2.GongGetIdentifier(stage))
		} else {
			// in case of nil pointer, we need to unstage the previous value
			res = PointerFieldInitStatement
			res = strings.ReplaceAll(res, "{{Identifier}}", category2shape.GongGetIdentifier(stage))
			res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "Category2")
			res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", "nil")
		}
	default:
		log.Panicf("Unknown field %s for Gongstruct Category2Shape", fieldName)
	}
	return
}

func (category3 *Category3) GongMarshallField(stage *Stage, fieldName string) (res string) {

	switch fieldName {
	case "Name":
		res = StringInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", category3.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "Name")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", string(category3.Name))

	default:
		log.Panicf("Unknown field %s for Gongstruct Category3", fieldName)
	}
	return
}

func (category3shape *Category3Shape) GongMarshallField(stage *Stage, fieldName string) (res string) {

	switch fieldName {
	case "Name":
		res = StringInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", category3shape.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "Name")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", string(category3shape.Name))
	case "X":
		res = NumberInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", category3shape.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "X")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%f", category3shape.X))
	case "Y":
		res = NumberInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", category3shape.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "Y")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%f", category3shape.Y))
	case "Width":
		res = NumberInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", category3shape.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "Width")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%f", category3shape.Width))
	case "Height":
		res = NumberInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", category3shape.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "Height")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%f", category3shape.Height))

	case "Category3":
		if category3shape.Category3 != nil {
			res = PointerFieldInitStatement
			res = strings.ReplaceAll(res, "{{Identifier}}", category3shape.GongGetIdentifier(stage))
			res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "Category3")
			res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", category3shape.Category3.GongGetIdentifier(stage))
		} else {
			// in case of nil pointer, we need to unstage the previous value
			res = PointerFieldInitStatement
			res = strings.ReplaceAll(res, "{{Identifier}}", category3shape.GongGetIdentifier(stage))
			res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "Category3")
			res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", "nil")
		}
	default:
		log.Panicf("Unknown field %s for Gongstruct Category3Shape", fieldName)
	}
	return
}

func (controlpointshape *ControlPointShape) GongMarshallField(stage *Stage, fieldName string) (res string) {

	switch fieldName {
	case "Name":
		res = StringInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", controlpointshape.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "Name")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", string(controlpointshape.Name))
	case "X_Relative":
		res = NumberInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", controlpointshape.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "X_Relative")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%f", controlpointshape.X_Relative))
	case "Y_Relative":
		res = NumberInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", controlpointshape.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "Y_Relative")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%f", controlpointshape.Y_Relative))
	case "IsStartShapeTheClosestShape":
		res = NumberInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", controlpointshape.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "IsStartShapeTheClosestShape")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%t", controlpointshape.IsStartShapeTheClosestShape))

	default:
		log.Panicf("Unknown field %s for Gongstruct ControlPointShape", fieldName)
	}
	return
}

func (desk *Desk) GongMarshallField(stage *Stage, fieldName string) (res string) {

	switch fieldName {
	case "Name":
		res = StringInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", desk.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "Name")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", string(desk.Name))

	case "SelectedDiagram":
		if desk.SelectedDiagram != nil {
			res = PointerFieldInitStatement
			res = strings.ReplaceAll(res, "{{Identifier}}", desk.GongGetIdentifier(stage))
			res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "SelectedDiagram")
			res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", desk.SelectedDiagram.GongGetIdentifier(stage))
		} else {
			// in case of nil pointer, we need to unstage the previous value
			res = PointerFieldInitStatement
			res = strings.ReplaceAll(res, "{{Identifier}}", desk.GongGetIdentifier(stage))
			res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "SelectedDiagram")
			res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", "nil")
		}
	default:
		log.Panicf("Unknown field %s for Gongstruct Desk", fieldName)
	}
	return
}

func (diagram *Diagram) GongMarshallField(stage *Stage, fieldName string) (res string) {

	switch fieldName {
	case "Name":
		res = StringInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", diagram.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "Name")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", string(diagram.Name))
	case "IsEditable":
		res = NumberInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", diagram.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "IsEditable")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%t", diagram.IsEditable))
	case "IsNodeExpanded":
		res = NumberInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", diagram.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "IsNodeExpanded")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%t", diagram.IsNodeExpanded))
	case "IsCategory1NodeExpanded":
		res = NumberInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", diagram.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "IsCategory1NodeExpanded")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%t", diagram.IsCategory1NodeExpanded))
	case "IsCategory2NodeExpanded":
		res = NumberInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", diagram.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "IsCategory2NodeExpanded")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%t", diagram.IsCategory2NodeExpanded))
	case "IsCategory3NodeExpanded":
		res = NumberInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", diagram.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "IsCategory3NodeExpanded")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%t", diagram.IsCategory3NodeExpanded))
	case "IsInfluenceCategoryNodeExpanded":
		res = NumberInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", diagram.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "IsInfluenceCategoryNodeExpanded")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%t", diagram.IsInfluenceCategoryNodeExpanded))
	case "IsCategory1Shown":
		res = NumberInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", diagram.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "IsCategory1Shown")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%t", diagram.IsCategory1Shown))
	case "IsCategory2Shown":
		res = NumberInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", diagram.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "IsCategory2Shown")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%t", diagram.IsCategory2Shown))
	case "IsCategory3Shown":
		res = NumberInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", diagram.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "IsCategory3Shown")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%t", diagram.IsCategory3Shown))
	case "IsInfluenceCategoryShown":
		res = NumberInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", diagram.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "IsInfluenceCategoryShown")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%t", diagram.IsInfluenceCategoryShown))
	case "XMargin":
		res = NumberInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", diagram.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "XMargin")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%f", diagram.XMargin))
	case "YMargin":
		res = NumberInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", diagram.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "YMargin")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%f", diagram.YMargin))
	case "Height":
		res = NumberInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", diagram.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "Height")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%f", diagram.Height))
	case "Width":
		res = NumberInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", diagram.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "Width")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%f", diagram.Width))
	case "NbPixPerCharacter":
		res = NumberInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", diagram.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "NbPixPerCharacter")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%f", diagram.NbPixPerCharacter))
	case "RedColorCode":
		res = StringInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", diagram.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "RedColorCode")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", string(diagram.RedColorCode))
	case "BackgroundGreyColorCode":
		res = StringInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", diagram.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "BackgroundGreyColorCode")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", string(diagram.BackgroundGreyColorCode))
	case "GrayColorCode":
		res = StringInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", diagram.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "GrayColorCode")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", string(diagram.GrayColorCode))
	case "Category1RectAnchorType":
		if diagram.Category1RectAnchorType != "" {
			res = StringEnumInitStatement
			res = strings.ReplaceAll(res, "{{Identifier}}", diagram.GongGetIdentifier(stage))
			res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "Category1RectAnchorType")
			res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", "models."+diagram.Category1RectAnchorType.ToCodeString())
		}
	case "Category1TextAnchorType":
		if diagram.Category1TextAnchorType != "" {
			res = StringEnumInitStatement
			res = strings.ReplaceAll(res, "{{Identifier}}", diagram.GongGetIdentifier(stage))
			res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "Category1TextAnchorType")
			res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", "models."+diagram.Category1TextAnchorType.ToCodeString())
		}
	case "Category1DominantBaselineType":
		if diagram.Category1DominantBaselineType != "" {
			res = StringEnumInitStatement
			res = strings.ReplaceAll(res, "{{Identifier}}", diagram.GongGetIdentifier(stage))
			res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "Category1DominantBaselineType")
			res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", "models."+diagram.Category1DominantBaselineType.ToCodeString())
		}
	case "Category1FontSize":
		res = StringInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", diagram.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "Category1FontSize")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", string(diagram.Category1FontSize))
	case "Category1FontWeigth":
		res = StringInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", diagram.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "Category1FontWeigth")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", string(diagram.Category1FontWeigth))
	case "Category1FontFamily":
		res = StringInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", diagram.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "Category1FontFamily")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", string(diagram.Category1FontFamily))
	case "Category1LetterSpacing":
		res = StringInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", diagram.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "Category1LetterSpacing")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", string(diagram.Category1LetterSpacing))
	case "Category2TypeFontSize":
		res = StringInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", diagram.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "Category2TypeFontSize")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", string(diagram.Category2TypeFontSize))
	case "Category2TypeFontWeigth":
		res = StringInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", diagram.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "Category2TypeFontWeigth")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", string(diagram.Category2TypeFontWeigth))
	case "Category2TypeFontFamily":
		res = StringInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", diagram.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "Category2TypeFontFamily")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", string(diagram.Category2TypeFontFamily))
	case "Category2TypeLetterSpacing":
		res = StringInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", diagram.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "Category2TypeLetterSpacing")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", string(diagram.Category2TypeLetterSpacing))
	case "Category2TypeRectAnchorType":
		if diagram.Category2TypeRectAnchorType != "" {
			res = StringEnumInitStatement
			res = strings.ReplaceAll(res, "{{Identifier}}", diagram.GongGetIdentifier(stage))
			res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "Category2TypeRectAnchorType")
			res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", "models."+diagram.Category2TypeRectAnchorType.ToCodeString())
		}
	case "Category2DominantBaselineType":
		if diagram.Category2DominantBaselineType != "" {
			res = StringEnumInitStatement
			res = strings.ReplaceAll(res, "{{Identifier}}", diagram.GongGetIdentifier(stage))
			res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "Category2DominantBaselineType")
			res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", "models."+diagram.Category2DominantBaselineType.ToCodeString())
		}
	case "Category2StrokeWidth":
		res = NumberInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", diagram.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "Category2StrokeWidth")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%f", diagram.Category2StrokeWidth))
	case "Category3RectAnchorType":
		if diagram.Category3RectAnchorType != "" {
			res = StringEnumInitStatement
			res = strings.ReplaceAll(res, "{{Identifier}}", diagram.GongGetIdentifier(stage))
			res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "Category3RectAnchorType")
			res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", "models."+diagram.Category3RectAnchorType.ToCodeString())
		}
	case "Category3TextAnchorType":
		if diagram.Category3TextAnchorType != "" {
			res = StringEnumInitStatement
			res = strings.ReplaceAll(res, "{{Identifier}}", diagram.GongGetIdentifier(stage))
			res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "Category3TextAnchorType")
			res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", "models."+diagram.Category3TextAnchorType.ToCodeString())
		}
	case "Category3DominantBaselineType":
		if diagram.Category3DominantBaselineType != "" {
			res = StringEnumInitStatement
			res = strings.ReplaceAll(res, "{{Identifier}}", diagram.GongGetIdentifier(stage))
			res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "Category3DominantBaselineType")
			res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", "models."+diagram.Category3DominantBaselineType.ToCodeString())
		}
	case "Category3FontSize":
		res = StringInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", diagram.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "Category3FontSize")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", string(diagram.Category3FontSize))
	case "Category3FontWeigth":
		res = StringInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", diagram.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "Category3FontWeigth")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", string(diagram.Category3FontWeigth))
	case "Category3FontFamily":
		res = StringInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", diagram.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "Category3FontFamily")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", string(diagram.Category3FontFamily))
	case "Category3LetterSpacing":
		res = StringInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", diagram.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "Category3LetterSpacing")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", string(diagram.Category3LetterSpacing))
	case "InfluenceStrokeWidth":
		res = NumberInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", diagram.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "InfluenceStrokeWidth")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%f", diagram.InfluenceStrokeWidth))
	case "InfluenceArrowSize":
		res = NumberInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", diagram.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "InfluenceArrowSize")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%f", diagram.InfluenceArrowSize))
	case "InfluenceArrowStartOffset":
		res = NumberInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", diagram.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "InfluenceArrowStartOffset")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%f", diagram.InfluenceArrowStartOffset))
	case "InfluenceArrowEndOffset":
		res = NumberInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", diagram.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "InfluenceArrowEndOffset")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%f", diagram.InfluenceArrowEndOffset))
	case "InfluenceCornerRadius":
		res = NumberInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", diagram.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "InfluenceCornerRadius")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%f", diagram.InfluenceCornerRadius))
	case "InfluenceFontSize":
		res = StringInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", diagram.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "InfluenceFontSize")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", string(diagram.InfluenceFontSize))
	case "InfluenceFontWeigth":
		res = StringInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", diagram.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "InfluenceFontWeigth")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", string(diagram.InfluenceFontWeigth))
	case "InfluenceFontFamily":
		res = StringInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", diagram.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "InfluenceFontFamily")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", string(diagram.InfluenceFontFamily))
	case "InfluenceLetterSpacing":
		res = StringInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", diagram.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "InfluenceLetterSpacing")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", string(diagram.InfluenceLetterSpacing))
	case "InfluenceDashedLinePattern":
		res = StringInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", diagram.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "InfluenceDashedLinePattern")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", string(diagram.InfluenceDashedLinePattern))

	case "Category1Shapes":
		for _, _category1shape := range diagram.Category1Shapes {
			tmp := SliceOfPointersFieldInitStatement
			tmp = strings.ReplaceAll(tmp, "{{Identifier}}", diagram.GongGetIdentifier(stage))
			tmp = strings.ReplaceAll(tmp, "{{GeneratedFieldName}}", "Category1Shapes")
			tmp = strings.ReplaceAll(tmp, "{{GeneratedFieldNameValue}}", _category1shape.GongGetIdentifier(stage))
			res += tmp
		}
	case "Category2Shapes":
		for _, _category2shape := range diagram.Category2Shapes {
			tmp := SliceOfPointersFieldInitStatement
			tmp = strings.ReplaceAll(tmp, "{{Identifier}}", diagram.GongGetIdentifier(stage))
			tmp = strings.ReplaceAll(tmp, "{{GeneratedFieldName}}", "Category2Shapes")
			tmp = strings.ReplaceAll(tmp, "{{GeneratedFieldNameValue}}", _category2shape.GongGetIdentifier(stage))
			res += tmp
		}
	case "Category3Shapes":
		for _, _category3shape := range diagram.Category3Shapes {
			tmp := SliceOfPointersFieldInitStatement
			tmp = strings.ReplaceAll(tmp, "{{Identifier}}", diagram.GongGetIdentifier(stage))
			tmp = strings.ReplaceAll(tmp, "{{GeneratedFieldName}}", "Category3Shapes")
			tmp = strings.ReplaceAll(tmp, "{{GeneratedFieldNameValue}}", _category3shape.GongGetIdentifier(stage))
			res += tmp
		}
	case "InfluenceShapes":
		for _, _influenceshape := range diagram.InfluenceShapes {
			tmp := SliceOfPointersFieldInitStatement
			tmp = strings.ReplaceAll(tmp, "{{Identifier}}", diagram.GongGetIdentifier(stage))
			tmp = strings.ReplaceAll(tmp, "{{GeneratedFieldName}}", "InfluenceShapes")
			tmp = strings.ReplaceAll(tmp, "{{GeneratedFieldNameValue}}", _influenceshape.GongGetIdentifier(stage))
			res += tmp
		}
	default:
		log.Panicf("Unknown field %s for Gongstruct Diagram", fieldName)
	}
	return
}

func (influence *Influence) GongMarshallField(stage *Stage, fieldName string) (res string) {

	switch fieldName {
	case "Name":
		res = StringInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", influence.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "Name")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", string(influence.Name))
	case "IsHypothtical":
		res = NumberInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", influence.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "IsHypothtical")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%t", influence.IsHypothtical))
	case "TextAtEndOfArrow":
		res = StringInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", influence.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "TextAtEndOfArrow")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", string(influence.TextAtEndOfArrow))

	case "SourceCategory1":
		if influence.SourceCategory1 != nil {
			res = PointerFieldInitStatement
			res = strings.ReplaceAll(res, "{{Identifier}}", influence.GongGetIdentifier(stage))
			res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "SourceCategory1")
			res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", influence.SourceCategory1.GongGetIdentifier(stage))
		} else {
			// in case of nil pointer, we need to unstage the previous value
			res = PointerFieldInitStatement
			res = strings.ReplaceAll(res, "{{Identifier}}", influence.GongGetIdentifier(stage))
			res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "SourceCategory1")
			res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", "nil")
		}
	case "SourceCategory2":
		if influence.SourceCategory2 != nil {
			res = PointerFieldInitStatement
			res = strings.ReplaceAll(res, "{{Identifier}}", influence.GongGetIdentifier(stage))
			res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "SourceCategory2")
			res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", influence.SourceCategory2.GongGetIdentifier(stage))
		} else {
			// in case of nil pointer, we need to unstage the previous value
			res = PointerFieldInitStatement
			res = strings.ReplaceAll(res, "{{Identifier}}", influence.GongGetIdentifier(stage))
			res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "SourceCategory2")
			res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", "nil")
		}
	case "SourceCategory3":
		if influence.SourceCategory3 != nil {
			res = PointerFieldInitStatement
			res = strings.ReplaceAll(res, "{{Identifier}}", influence.GongGetIdentifier(stage))
			res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "SourceCategory3")
			res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", influence.SourceCategory3.GongGetIdentifier(stage))
		} else {
			// in case of nil pointer, we need to unstage the previous value
			res = PointerFieldInitStatement
			res = strings.ReplaceAll(res, "{{Identifier}}", influence.GongGetIdentifier(stage))
			res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "SourceCategory3")
			res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", "nil")
		}
	case "TargetCategory1":
		if influence.TargetCategory1 != nil {
			res = PointerFieldInitStatement
			res = strings.ReplaceAll(res, "{{Identifier}}", influence.GongGetIdentifier(stage))
			res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "TargetCategory1")
			res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", influence.TargetCategory1.GongGetIdentifier(stage))
		} else {
			// in case of nil pointer, we need to unstage the previous value
			res = PointerFieldInitStatement
			res = strings.ReplaceAll(res, "{{Identifier}}", influence.GongGetIdentifier(stage))
			res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "TargetCategory1")
			res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", "nil")
		}
	case "TargetCategory2":
		if influence.TargetCategory2 != nil {
			res = PointerFieldInitStatement
			res = strings.ReplaceAll(res, "{{Identifier}}", influence.GongGetIdentifier(stage))
			res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "TargetCategory2")
			res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", influence.TargetCategory2.GongGetIdentifier(stage))
		} else {
			// in case of nil pointer, we need to unstage the previous value
			res = PointerFieldInitStatement
			res = strings.ReplaceAll(res, "{{Identifier}}", influence.GongGetIdentifier(stage))
			res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "TargetCategory2")
			res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", "nil")
		}
	case "TargetCategory3":
		if influence.TargetCategory3 != nil {
			res = PointerFieldInitStatement
			res = strings.ReplaceAll(res, "{{Identifier}}", influence.GongGetIdentifier(stage))
			res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "TargetCategory3")
			res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", influence.TargetCategory3.GongGetIdentifier(stage))
		} else {
			// in case of nil pointer, we need to unstage the previous value
			res = PointerFieldInitStatement
			res = strings.ReplaceAll(res, "{{Identifier}}", influence.GongGetIdentifier(stage))
			res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "TargetCategory3")
			res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", "nil")
		}
	default:
		log.Panicf("Unknown field %s for Gongstruct Influence", fieldName)
	}
	return
}

func (influenceshape *InfluenceShape) GongMarshallField(stage *Stage, fieldName string) (res string) {

	switch fieldName {
	case "Name":
		res = StringInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", influenceshape.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "Name")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", string(influenceshape.Name))

	case "Influence":
		if influenceshape.Influence != nil {
			res = PointerFieldInitStatement
			res = strings.ReplaceAll(res, "{{Identifier}}", influenceshape.GongGetIdentifier(stage))
			res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "Influence")
			res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", influenceshape.Influence.GongGetIdentifier(stage))
		} else {
			// in case of nil pointer, we need to unstage the previous value
			res = PointerFieldInitStatement
			res = strings.ReplaceAll(res, "{{Identifier}}", influenceshape.GongGetIdentifier(stage))
			res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "Influence")
			res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", "nil")
		}
	case "ControlPointShapes":
		for _, _controlpointshape := range influenceshape.ControlPointShapes {
			tmp := SliceOfPointersFieldInitStatement
			tmp = strings.ReplaceAll(tmp, "{{Identifier}}", influenceshape.GongGetIdentifier(stage))
			tmp = strings.ReplaceAll(tmp, "{{GeneratedFieldName}}", "ControlPointShapes")
			tmp = strings.ReplaceAll(tmp, "{{GeneratedFieldNameValue}}", _controlpointshape.GongGetIdentifier(stage))
			res += tmp
		}
	default:
		log.Panicf("Unknown field %s for Gongstruct InfluenceShape", fieldName)
	}
	return
}

// insertion point for marshall all fields methods
func (category1 *Category1) GongMarshallAllFields(stage *Stage) (initializerStatements string, pointersInitializesStatements string) {

	{ // Insertion point for basic fields value assignment
		initializerStatements += category1.GongMarshallField(stage, "Name")
	}
	return
}
func (category1shape *Category1Shape) GongMarshallAllFields(stage *Stage) (initializerStatements string, pointersInitializesStatements string) {

	{ // Insertion point for basic fields value assignment
		initializerStatements += category1shape.GongMarshallField(stage, "Name")
		pointersInitializesStatements += category1shape.GongMarshallField(stage, "Category1")
		initializerStatements += category1shape.GongMarshallField(stage, "X")
		initializerStatements += category1shape.GongMarshallField(stage, "Y")
		initializerStatements += category1shape.GongMarshallField(stage, "Width")
		initializerStatements += category1shape.GongMarshallField(stage, "Height")
	}
	return
}
func (category2 *Category2) GongMarshallAllFields(stage *Stage) (initializerStatements string, pointersInitializesStatements string) {

	{ // Insertion point for basic fields value assignment
		initializerStatements += category2.GongMarshallField(stage, "Name")
	}
	return
}
func (category2shape *Category2Shape) GongMarshallAllFields(stage *Stage) (initializerStatements string, pointersInitializesStatements string) {

	{ // Insertion point for basic fields value assignment
		initializerStatements += category2shape.GongMarshallField(stage, "Name")
		pointersInitializesStatements += category2shape.GongMarshallField(stage, "Category2")
		initializerStatements += category2shape.GongMarshallField(stage, "X")
		initializerStatements += category2shape.GongMarshallField(stage, "Y")
		initializerStatements += category2shape.GongMarshallField(stage, "Width")
		initializerStatements += category2shape.GongMarshallField(stage, "Height")
	}
	return
}
func (category3 *Category3) GongMarshallAllFields(stage *Stage) (initializerStatements string, pointersInitializesStatements string) {

	{ // Insertion point for basic fields value assignment
		initializerStatements += category3.GongMarshallField(stage, "Name")
	}
	return
}
func (category3shape *Category3Shape) GongMarshallAllFields(stage *Stage) (initializerStatements string, pointersInitializesStatements string) {

	{ // Insertion point for basic fields value assignment
		initializerStatements += category3shape.GongMarshallField(stage, "Name")
		pointersInitializesStatements += category3shape.GongMarshallField(stage, "Category3")
		initializerStatements += category3shape.GongMarshallField(stage, "X")
		initializerStatements += category3shape.GongMarshallField(stage, "Y")
		initializerStatements += category3shape.GongMarshallField(stage, "Width")
		initializerStatements += category3shape.GongMarshallField(stage, "Height")
	}
	return
}
func (controlpointshape *ControlPointShape) GongMarshallAllFields(stage *Stage) (initializerStatements string, pointersInitializesStatements string) {

	{ // Insertion point for basic fields value assignment
		initializerStatements += controlpointshape.GongMarshallField(stage, "Name")
		initializerStatements += controlpointshape.GongMarshallField(stage, "X_Relative")
		initializerStatements += controlpointshape.GongMarshallField(stage, "Y_Relative")
		initializerStatements += controlpointshape.GongMarshallField(stage, "IsStartShapeTheClosestShape")
	}
	return
}
func (desk *Desk) GongMarshallAllFields(stage *Stage) (initializerStatements string, pointersInitializesStatements string) {

	{ // Insertion point for basic fields value assignment
		initializerStatements += desk.GongMarshallField(stage, "Name")
		pointersInitializesStatements += desk.GongMarshallField(stage, "SelectedDiagram")
	}
	return
}
func (diagram *Diagram) GongMarshallAllFields(stage *Stage) (initializerStatements string, pointersInitializesStatements string) {

	{ // Insertion point for basic fields value assignment
		initializerStatements += diagram.GongMarshallField(stage, "Name")
		pointersInitializesStatements += diagram.GongMarshallField(stage, "Category1Shapes")
		pointersInitializesStatements += diagram.GongMarshallField(stage, "Category2Shapes")
		pointersInitializesStatements += diagram.GongMarshallField(stage, "Category3Shapes")
		pointersInitializesStatements += diagram.GongMarshallField(stage, "InfluenceShapes")
		initializerStatements += diagram.GongMarshallField(stage, "IsEditable")
		initializerStatements += diagram.GongMarshallField(stage, "IsNodeExpanded")
		initializerStatements += diagram.GongMarshallField(stage, "IsCategory1NodeExpanded")
		initializerStatements += diagram.GongMarshallField(stage, "IsCategory2NodeExpanded")
		initializerStatements += diagram.GongMarshallField(stage, "IsCategory3NodeExpanded")
		initializerStatements += diagram.GongMarshallField(stage, "IsInfluenceCategoryNodeExpanded")
		initializerStatements += diagram.GongMarshallField(stage, "IsCategory1Shown")
		initializerStatements += diagram.GongMarshallField(stage, "IsCategory2Shown")
		initializerStatements += diagram.GongMarshallField(stage, "IsCategory3Shown")
		initializerStatements += diagram.GongMarshallField(stage, "IsInfluenceCategoryShown")
		initializerStatements += diagram.GongMarshallField(stage, "XMargin")
		initializerStatements += diagram.GongMarshallField(stage, "YMargin")
		initializerStatements += diagram.GongMarshallField(stage, "Height")
		initializerStatements += diagram.GongMarshallField(stage, "Width")
		initializerStatements += diagram.GongMarshallField(stage, "NbPixPerCharacter")
		initializerStatements += diagram.GongMarshallField(stage, "RedColorCode")
		initializerStatements += diagram.GongMarshallField(stage, "BackgroundGreyColorCode")
		initializerStatements += diagram.GongMarshallField(stage, "GrayColorCode")
		initializerStatements += diagram.GongMarshallField(stage, "Category1RectAnchorType")
		initializerStatements += diagram.GongMarshallField(stage, "Category1TextAnchorType")
		initializerStatements += diagram.GongMarshallField(stage, "Category1DominantBaselineType")
		initializerStatements += diagram.GongMarshallField(stage, "Category1FontSize")
		initializerStatements += diagram.GongMarshallField(stage, "Category1FontWeigth")
		initializerStatements += diagram.GongMarshallField(stage, "Category1FontFamily")
		initializerStatements += diagram.GongMarshallField(stage, "Category1LetterSpacing")
		initializerStatements += diagram.GongMarshallField(stage, "Category2TypeFontSize")
		initializerStatements += diagram.GongMarshallField(stage, "Category2TypeFontWeigth")
		initializerStatements += diagram.GongMarshallField(stage, "Category2TypeFontFamily")
		initializerStatements += diagram.GongMarshallField(stage, "Category2TypeLetterSpacing")
		initializerStatements += diagram.GongMarshallField(stage, "Category2TypeRectAnchorType")
		initializerStatements += diagram.GongMarshallField(stage, "Category2DominantBaselineType")
		initializerStatements += diagram.GongMarshallField(stage, "Category2StrokeWidth")
		initializerStatements += diagram.GongMarshallField(stage, "Category3RectAnchorType")
		initializerStatements += diagram.GongMarshallField(stage, "Category3TextAnchorType")
		initializerStatements += diagram.GongMarshallField(stage, "Category3DominantBaselineType")
		initializerStatements += diagram.GongMarshallField(stage, "Category3FontSize")
		initializerStatements += diagram.GongMarshallField(stage, "Category3FontWeigth")
		initializerStatements += diagram.GongMarshallField(stage, "Category3FontFamily")
		initializerStatements += diagram.GongMarshallField(stage, "Category3LetterSpacing")
		initializerStatements += diagram.GongMarshallField(stage, "InfluenceStrokeWidth")
		initializerStatements += diagram.GongMarshallField(stage, "InfluenceArrowSize")
		initializerStatements += diagram.GongMarshallField(stage, "InfluenceArrowStartOffset")
		initializerStatements += diagram.GongMarshallField(stage, "InfluenceArrowEndOffset")
		initializerStatements += diagram.GongMarshallField(stage, "InfluenceCornerRadius")
		initializerStatements += diagram.GongMarshallField(stage, "InfluenceFontSize")
		initializerStatements += diagram.GongMarshallField(stage, "InfluenceFontWeigth")
		initializerStatements += diagram.GongMarshallField(stage, "InfluenceFontFamily")
		initializerStatements += diagram.GongMarshallField(stage, "InfluenceLetterSpacing")
		initializerStatements += diagram.GongMarshallField(stage, "InfluenceDashedLinePattern")
	}
	return
}
func (influence *Influence) GongMarshallAllFields(stage *Stage) (initializerStatements string, pointersInitializesStatements string) {

	{ // Insertion point for basic fields value assignment
		initializerStatements += influence.GongMarshallField(stage, "Name")
		pointersInitializesStatements += influence.GongMarshallField(stage, "SourceCategory1")
		pointersInitializesStatements += influence.GongMarshallField(stage, "SourceCategory2")
		pointersInitializesStatements += influence.GongMarshallField(stage, "SourceCategory3")
		pointersInitializesStatements += influence.GongMarshallField(stage, "TargetCategory1")
		pointersInitializesStatements += influence.GongMarshallField(stage, "TargetCategory2")
		pointersInitializesStatements += influence.GongMarshallField(stage, "TargetCategory3")
		initializerStatements += influence.GongMarshallField(stage, "IsHypothtical")
		initializerStatements += influence.GongMarshallField(stage, "TextAtEndOfArrow")
	}
	return
}
func (influenceshape *InfluenceShape) GongMarshallAllFields(stage *Stage) (initializerStatements string, pointersInitializesStatements string) {

	{ // Insertion point for basic fields value assignment
		initializerStatements += influenceshape.GongMarshallField(stage, "Name")
		pointersInitializesStatements += influenceshape.GongMarshallField(stage, "Influence")
		pointersInitializesStatements += influenceshape.GongMarshallField(stage, "ControlPointShapes")
	}
	return
}
