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
	map_ArtefactType_Identifiers := make(map[*ArtefactType]string)
	_ = map_ArtefactType_Identifiers

	artefacttypeOrdered := []*ArtefactType{}
	for artefacttype := range stage.ArtefactTypes {
		artefacttypeOrdered = append(artefacttypeOrdered, artefacttype)
	}
	sort.Slice(artefacttypeOrdered[:], func(i, j int) bool {
		artefacttypei := artefacttypeOrdered[i]
		artefacttypej := artefacttypeOrdered[j]
		artefacttypei_order, oki := stage.ArtefactTypeMap_Staged_Order[artefacttypei]
		artefacttypej_order, okj := stage.ArtefactTypeMap_Staged_Order[artefacttypej]
		if !oki || !okj {
			log.Fatalln("unknown pointers")
		}
		return artefacttypei_order < artefacttypej_order
	})
	if len(artefacttypeOrdered) > 0 {
		identifiersDecl += "\n"
	}
	for idx, artefacttype := range artefacttypeOrdered {

		id = generatesIdentifier("ArtefactType", idx, artefacttype.Name)
		map_ArtefactType_Identifiers[artefacttype] = id

		decl = IdentifiersDecls
		decl = strings.ReplaceAll(decl, "{{Identifier}}", id)
		decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "ArtefactType")
		decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", artefacttype.Name)
		identifiersDecl += decl

		initializerStatements += "\n"
		// Initialisation of values
		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Name")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(artefacttype.Name))
		initializerStatements += setValueField

	}

	map_ArtefactTypeShape_Identifiers := make(map[*ArtefactTypeShape]string)
	_ = map_ArtefactTypeShape_Identifiers

	artefacttypeshapeOrdered := []*ArtefactTypeShape{}
	for artefacttypeshape := range stage.ArtefactTypeShapes {
		artefacttypeshapeOrdered = append(artefacttypeshapeOrdered, artefacttypeshape)
	}
	sort.Slice(artefacttypeshapeOrdered[:], func(i, j int) bool {
		artefacttypeshapei := artefacttypeshapeOrdered[i]
		artefacttypeshapej := artefacttypeshapeOrdered[j]
		artefacttypeshapei_order, oki := stage.ArtefactTypeShapeMap_Staged_Order[artefacttypeshapei]
		artefacttypeshapej_order, okj := stage.ArtefactTypeShapeMap_Staged_Order[artefacttypeshapej]
		if !oki || !okj {
			log.Fatalln("unknown pointers")
		}
		return artefacttypeshapei_order < artefacttypeshapej_order
	})
	if len(artefacttypeshapeOrdered) > 0 {
		identifiersDecl += "\n"
	}
	for idx, artefacttypeshape := range artefacttypeshapeOrdered {

		id = generatesIdentifier("ArtefactTypeShape", idx, artefacttypeshape.Name)
		map_ArtefactTypeShape_Identifiers[artefacttypeshape] = id

		decl = IdentifiersDecls
		decl = strings.ReplaceAll(decl, "{{Identifier}}", id)
		decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "ArtefactTypeShape")
		decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", artefacttypeshape.Name)
		identifiersDecl += decl

		initializerStatements += "\n"
		// Initialisation of values
		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Name")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(artefacttypeshape.Name))
		initializerStatements += setValueField

		setValueField = NumberInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "X")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%f", artefacttypeshape.X))
		initializerStatements += setValueField

		setValueField = NumberInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Y")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%f", artefacttypeshape.Y))
		initializerStatements += setValueField

		setValueField = NumberInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Width")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%f", artefacttypeshape.Width))
		initializerStatements += setValueField

		setValueField = NumberInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Height")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%f", artefacttypeshape.Height))
		initializerStatements += setValueField

	}

	map_Artist_Identifiers := make(map[*Artist]string)
	_ = map_Artist_Identifiers

	artistOrdered := []*Artist{}
	for artist := range stage.Artists {
		artistOrdered = append(artistOrdered, artist)
	}
	sort.Slice(artistOrdered[:], func(i, j int) bool {
		artisti := artistOrdered[i]
		artistj := artistOrdered[j]
		artisti_order, oki := stage.ArtistMap_Staged_Order[artisti]
		artistj_order, okj := stage.ArtistMap_Staged_Order[artistj]
		if !oki || !okj {
			log.Fatalln("unknown pointers")
		}
		return artisti_order < artistj_order
	})
	if len(artistOrdered) > 0 {
		identifiersDecl += "\n"
	}
	for idx, artist := range artistOrdered {

		id = generatesIdentifier("Artist", idx, artist.Name)
		map_Artist_Identifiers[artist] = id

		decl = IdentifiersDecls
		decl = strings.ReplaceAll(decl, "{{Identifier}}", id)
		decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "Artist")
		decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", artist.Name)
		identifiersDecl += decl

		initializerStatements += "\n"
		// Initialisation of values
		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Name")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(artist.Name))
		initializerStatements += setValueField

	}

	map_ArtistShape_Identifiers := make(map[*ArtistShape]string)
	_ = map_ArtistShape_Identifiers

	artistshapeOrdered := []*ArtistShape{}
	for artistshape := range stage.ArtistShapes {
		artistshapeOrdered = append(artistshapeOrdered, artistshape)
	}
	sort.Slice(artistshapeOrdered[:], func(i, j int) bool {
		artistshapei := artistshapeOrdered[i]
		artistshapej := artistshapeOrdered[j]
		artistshapei_order, oki := stage.ArtistShapeMap_Staged_Order[artistshapei]
		artistshapej_order, okj := stage.ArtistShapeMap_Staged_Order[artistshapej]
		if !oki || !okj {
			log.Fatalln("unknown pointers")
		}
		return artistshapei_order < artistshapej_order
	})
	if len(artistshapeOrdered) > 0 {
		identifiersDecl += "\n"
	}
	for idx, artistshape := range artistshapeOrdered {

		id = generatesIdentifier("ArtistShape", idx, artistshape.Name)
		map_ArtistShape_Identifiers[artistshape] = id

		decl = IdentifiersDecls
		decl = strings.ReplaceAll(decl, "{{Identifier}}", id)
		decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "ArtistShape")
		decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", artistshape.Name)
		identifiersDecl += decl

		initializerStatements += "\n"
		// Initialisation of values
		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Name")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(artistshape.Name))
		initializerStatements += setValueField

		setValueField = NumberInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "X")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%f", artistshape.X))
		initializerStatements += setValueField

		setValueField = NumberInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Y")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%f", artistshape.Y))
		initializerStatements += setValueField

		setValueField = NumberInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Width")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%f", artistshape.Width))
		initializerStatements += setValueField

		setValueField = NumberInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Height")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%f", artistshape.Height))
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
	for idx, controlpointshape := range controlpointshapeOrdered {

		id = generatesIdentifier("ControlPointShape", idx, controlpointshape.Name)
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
	for idx, desk := range deskOrdered {

		id = generatesIdentifier("Desk", idx, desk.Name)
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
	for idx, diagram := range diagramOrdered {

		id = generatesIdentifier("Diagram", idx, diagram.Name)
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
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "IsMovementCategoryNodeExpanded")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%t", diagram.IsMovementCategoryNodeExpanded))
		initializerStatements += setValueField

		setValueField = NumberInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "IsArtefactTypeCategoryNodeExpanded")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%t", diagram.IsArtefactTypeCategoryNodeExpanded))
		initializerStatements += setValueField

		setValueField = NumberInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "IsArtistCategoryNodeExpanded")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%t", diagram.IsArtistCategoryNodeExpanded))
		initializerStatements += setValueField

		setValueField = NumberInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "IsInfluenceCategoryNodeExpanded")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%t", diagram.IsInfluenceCategoryNodeExpanded))
		initializerStatements += setValueField

		setValueField = NumberInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "IsMovementCategoryShown")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%t", diagram.IsMovementCategoryShown))
		initializerStatements += setValueField

		setValueField = NumberInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "IsArtefactTypeCategoryShown")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%t", diagram.IsArtefactTypeCategoryShown))
		initializerStatements += setValueField

		setValueField = NumberInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "IsArtistCategoryShown")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%t", diagram.IsArtistCategoryShown))
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
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "NextVerticalDateXMargin")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%f", diagram.NextVerticalDateXMargin))
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

		setValueField = NumberInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "BottomBoxYOffset")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%f", diagram.BottomBoxYOffset))
		initializerStatements += setValueField

		setValueField = NumberInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "BottomBoxWidth")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%f", diagram.BottomBoxWidth))
		initializerStatements += setValueField

		setValueField = NumberInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "BottomBoxHeigth")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%f", diagram.BottomBoxHeigth))
		initializerStatements += setValueField

		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "BottomBoxFontSize")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(diagram.BottomBoxFontSize))
		initializerStatements += setValueField

		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "BottomBoxFontWeigth")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(diagram.BottomBoxFontWeigth))
		initializerStatements += setValueField

		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "BottomBoxFontFamily")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(diagram.BottomBoxFontFamily))
		initializerStatements += setValueField

		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "BottomBoxLetterSpacing")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(diagram.BottomBoxLetterSpacing))
		initializerStatements += setValueField

		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "BottomBoxLetterColorCode")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(diagram.BottomBoxLetterColorCode))
		initializerStatements += setValueField

		if diagram.MovementRectAnchorType != "" {
			setValueField = StringEnumInitStatement
			setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
			setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "MovementRectAnchorType")
			setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", "models."+diagram.MovementRectAnchorType.ToCodeString())
			initializerStatements += setValueField
		}

		if diagram.MovementTextAnchorType != "" {
			setValueField = StringEnumInitStatement
			setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
			setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "MovementTextAnchorType")
			setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", "models."+diagram.MovementTextAnchorType.ToCodeString())
			initializerStatements += setValueField
		}

		if diagram.MovementDominantBaselineType != "" {
			setValueField = StringEnumInitStatement
			setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
			setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "MovementDominantBaselineType")
			setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", "models."+diagram.MovementDominantBaselineType.ToCodeString())
			initializerStatements += setValueField
		}

		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "MovementFontSize")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(diagram.MovementFontSize))
		initializerStatements += setValueField

		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "MajorMovementFontSize")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(diagram.MajorMovementFontSize))
		initializerStatements += setValueField

		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "MinorMovementFontSize")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(diagram.MinorMovementFontSize))
		initializerStatements += setValueField

		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "MovementFontWeigth")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(diagram.MovementFontWeigth))
		initializerStatements += setValueField

		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "MovementFontFamily")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(diagram.MovementFontFamily))
		initializerStatements += setValueField

		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "MovementLetterSpacing")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(diagram.MovementLetterSpacing))
		initializerStatements += setValueField

		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "AbstractMovementFontSize")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(diagram.AbstractMovementFontSize))
		initializerStatements += setValueField

		if diagram.AbstractMovementRectAnchorType != "" {
			setValueField = StringEnumInitStatement
			setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
			setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "AbstractMovementRectAnchorType")
			setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", "models."+diagram.AbstractMovementRectAnchorType.ToCodeString())
			initializerStatements += setValueField
		}

		if diagram.AbstractMovementTextAnchorType != "" {
			setValueField = StringEnumInitStatement
			setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
			setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "AbstractMovementTextAnchorType")
			setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", "models."+diagram.AbstractMovementTextAnchorType.ToCodeString())
			initializerStatements += setValueField
		}

		if diagram.AbstractDominantBaselineType != "" {
			setValueField = StringEnumInitStatement
			setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
			setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "AbstractDominantBaselineType")
			setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", "models."+diagram.AbstractDominantBaselineType.ToCodeString())
			initializerStatements += setValueField
		}

		if diagram.MovementDateRectAnchorType != "" {
			setValueField = StringEnumInitStatement
			setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
			setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "MovementDateRectAnchorType")
			setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", "models."+diagram.MovementDateRectAnchorType.ToCodeString())
			initializerStatements += setValueField
		}

		if diagram.MovementDateTextAnchorType != "" {
			setValueField = StringEnumInitStatement
			setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
			setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "MovementDateTextAnchorType")
			setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", "models."+diagram.MovementDateTextAnchorType.ToCodeString())
			initializerStatements += setValueField
		}

		if diagram.MovementDateTextDominantBaselineType != "" {
			setValueField = StringEnumInitStatement
			setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
			setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "MovementDateTextDominantBaselineType")
			setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", "models."+diagram.MovementDateTextDominantBaselineType.ToCodeString())
			initializerStatements += setValueField
		}

		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "MovementDateAndPlacesFontSize")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(diagram.MovementDateAndPlacesFontSize))
		initializerStatements += setValueField

		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "MovementDateAndPlacesFontWeigth")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(diagram.MovementDateAndPlacesFontWeigth))
		initializerStatements += setValueField

		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "MovementDateAndPlacesFontFamily")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(diagram.MovementDateAndPlacesFontFamily))
		initializerStatements += setValueField

		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "MovementDateAndPlacesLetterSpacing")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(diagram.MovementDateAndPlacesLetterSpacing))
		initializerStatements += setValueField

		setValueField = NumberInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "MovementBelowArcY_Offset")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%f", diagram.MovementBelowArcY_Offset))
		initializerStatements += setValueField

		setValueField = NumberInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "MovementBelowArcY_OffsetPerPlace")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%f", diagram.MovementBelowArcY_OffsetPerPlace))
		initializerStatements += setValueField

		if diagram.MovementPlacesRectAnchorType != "" {
			setValueField = StringEnumInitStatement
			setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
			setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "MovementPlacesRectAnchorType")
			setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", "models."+diagram.MovementPlacesRectAnchorType.ToCodeString())
			initializerStatements += setValueField
		}

		if diagram.MovementPlacesTextAnchorType != "" {
			setValueField = StringEnumInitStatement
			setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
			setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "MovementPlacesTextAnchorType")
			setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", "models."+diagram.MovementPlacesTextAnchorType.ToCodeString())
			initializerStatements += setValueField
		}

		if diagram.MovementPlacesDominantBaselineType != "" {
			setValueField = StringEnumInitStatement
			setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
			setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "MovementPlacesDominantBaselineType")
			setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", "models."+diagram.MovementPlacesDominantBaselineType.ToCodeString())
			initializerStatements += setValueField
		}

		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "ArtefactTypeFontSize")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(diagram.ArtefactTypeFontSize))
		initializerStatements += setValueField

		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "ArtefactTypeFontWeigth")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(diagram.ArtefactTypeFontWeigth))
		initializerStatements += setValueField

		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "ArtefactTypeFontFamily")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(diagram.ArtefactTypeFontFamily))
		initializerStatements += setValueField

		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "ArtefactTypeLetterSpacing")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(diagram.ArtefactTypeLetterSpacing))
		initializerStatements += setValueField

		if diagram.ArtefactTypeRectAnchorType != "" {
			setValueField = StringEnumInitStatement
			setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
			setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "ArtefactTypeRectAnchorType")
			setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", "models."+diagram.ArtefactTypeRectAnchorType.ToCodeString())
			initializerStatements += setValueField
		}

		if diagram.ArtefactDominantBaselineType != "" {
			setValueField = StringEnumInitStatement
			setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
			setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "ArtefactDominantBaselineType")
			setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", "models."+diagram.ArtefactDominantBaselineType.ToCodeString())
			initializerStatements += setValueField
		}

		setValueField = NumberInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "ArtefactTypeStrokeWidth")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%f", diagram.ArtefactTypeStrokeWidth))
		initializerStatements += setValueField

		if diagram.ArtistRectAnchorType != "" {
			setValueField = StringEnumInitStatement
			setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
			setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "ArtistRectAnchorType")
			setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", "models."+diagram.ArtistRectAnchorType.ToCodeString())
			initializerStatements += setValueField
		}

		if diagram.ArtistTextAnchorType != "" {
			setValueField = StringEnumInitStatement
			setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
			setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "ArtistTextAnchorType")
			setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", "models."+diagram.ArtistTextAnchorType.ToCodeString())
			initializerStatements += setValueField
		}

		if diagram.ArtistDominantBaselineType != "" {
			setValueField = StringEnumInitStatement
			setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
			setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "ArtistDominantBaselineType")
			setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", "models."+diagram.ArtistDominantBaselineType.ToCodeString())
			initializerStatements += setValueField
		}

		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "ArtistFontSize")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(diagram.ArtistFontSize))
		initializerStatements += setValueField

		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "MajorArtistFontSize")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(diagram.MajorArtistFontSize))
		initializerStatements += setValueField

		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "MinorArtistFontSize")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(diagram.MinorArtistFontSize))
		initializerStatements += setValueField

		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "ArtistFontWeigth")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(diagram.ArtistFontWeigth))
		initializerStatements += setValueField

		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "ArtistFontFamily")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(diagram.ArtistFontFamily))
		initializerStatements += setValueField

		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "ArtistLetterSpacing")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(diagram.ArtistLetterSpacing))
		initializerStatements += setValueField

		if diagram.ArtistDateRectAnchorType != "" {
			setValueField = StringEnumInitStatement
			setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
			setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "ArtistDateRectAnchorType")
			setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", "models."+diagram.ArtistDateRectAnchorType.ToCodeString())
			initializerStatements += setValueField
		}

		if diagram.ArtistDateTextAnchorType != "" {
			setValueField = StringEnumInitStatement
			setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
			setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "ArtistDateTextAnchorType")
			setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", "models."+diagram.ArtistDateTextAnchorType.ToCodeString())
			initializerStatements += setValueField
		}

		if diagram.ArtistDateDominantBaselineType != "" {
			setValueField = StringEnumInitStatement
			setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
			setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "ArtistDateDominantBaselineType")
			setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", "models."+diagram.ArtistDateDominantBaselineType.ToCodeString())
			initializerStatements += setValueField
		}

		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "ArtistDateAndPlacesFontSize")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(diagram.ArtistDateAndPlacesFontSize))
		initializerStatements += setValueField

		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "ArtistDateAndPlacesFontWeigth")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(diagram.ArtistDateAndPlacesFontWeigth))
		initializerStatements += setValueField

		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "ArtistDateAndPlacesFontFamily")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(diagram.ArtistDateAndPlacesFontFamily))
		initializerStatements += setValueField

		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "ArtistDateAndPlacesLetterSpacing")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(diagram.ArtistDateAndPlacesLetterSpacing))
		initializerStatements += setValueField

		if diagram.ArtistPlacesRectAnchorType != "" {
			setValueField = StringEnumInitStatement
			setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
			setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "ArtistPlacesRectAnchorType")
			setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", "models."+diagram.ArtistPlacesRectAnchorType.ToCodeString())
			initializerStatements += setValueField
		}

		if diagram.ArtistPlacesTextAnchorType != "" {
			setValueField = StringEnumInitStatement
			setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
			setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "ArtistPlacesTextAnchorType")
			setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", "models."+diagram.ArtistPlacesTextAnchorType.ToCodeString())
			initializerStatements += setValueField
		}

		if diagram.ArtistPlacesDominantBaselineType != "" {
			setValueField = StringEnumInitStatement
			setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
			setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "ArtistPlacesDominantBaselineType")
			setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", "models."+diagram.ArtistPlacesDominantBaselineType.ToCodeString())
			initializerStatements += setValueField
		}

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
	for idx, influence := range influenceOrdered {

		id = generatesIdentifier("Influence", idx, influence.Name)
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
	for idx, influenceshape := range influenceshapeOrdered {

		id = generatesIdentifier("InfluenceShape", idx, influenceshape.Name)
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

	map_Movement_Identifiers := make(map[*Movement]string)
	_ = map_Movement_Identifiers

	movementOrdered := []*Movement{}
	for movement := range stage.Movements {
		movementOrdered = append(movementOrdered, movement)
	}
	sort.Slice(movementOrdered[:], func(i, j int) bool {
		movementi := movementOrdered[i]
		movementj := movementOrdered[j]
		movementi_order, oki := stage.MovementMap_Staged_Order[movementi]
		movementj_order, okj := stage.MovementMap_Staged_Order[movementj]
		if !oki || !okj {
			log.Fatalln("unknown pointers")
		}
		return movementi_order < movementj_order
	})
	if len(movementOrdered) > 0 {
		identifiersDecl += "\n"
	}
	for idx, movement := range movementOrdered {

		id = generatesIdentifier("Movement", idx, movement.Name)
		map_Movement_Identifiers[movement] = id

		decl = IdentifiersDecls
		decl = strings.ReplaceAll(decl, "{{Identifier}}", id)
		decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "Movement")
		decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", movement.Name)
		identifiersDecl += decl

		initializerStatements += "\n"
		// Initialisation of values
		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Name")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(movement.Name))
		initializerStatements += setValueField

	}

	map_MovementShape_Identifiers := make(map[*MovementShape]string)
	_ = map_MovementShape_Identifiers

	movementshapeOrdered := []*MovementShape{}
	for movementshape := range stage.MovementShapes {
		movementshapeOrdered = append(movementshapeOrdered, movementshape)
	}
	sort.Slice(movementshapeOrdered[:], func(i, j int) bool {
		movementshapei := movementshapeOrdered[i]
		movementshapej := movementshapeOrdered[j]
		movementshapei_order, oki := stage.MovementShapeMap_Staged_Order[movementshapei]
		movementshapej_order, okj := stage.MovementShapeMap_Staged_Order[movementshapej]
		if !oki || !okj {
			log.Fatalln("unknown pointers")
		}
		return movementshapei_order < movementshapej_order
	})
	if len(movementshapeOrdered) > 0 {
		identifiersDecl += "\n"
	}
	for idx, movementshape := range movementshapeOrdered {

		id = generatesIdentifier("MovementShape", idx, movementshape.Name)
		map_MovementShape_Identifiers[movementshape] = id

		decl = IdentifiersDecls
		decl = strings.ReplaceAll(decl, "{{Identifier}}", id)
		decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "MovementShape")
		decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", movementshape.Name)
		identifiersDecl += decl

		initializerStatements += "\n"
		// Initialisation of values
		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Name")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(movementshape.Name))
		initializerStatements += setValueField

		setValueField = NumberInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "X")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%f", movementshape.X))
		initializerStatements += setValueField

		setValueField = NumberInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Y")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%f", movementshape.Y))
		initializerStatements += setValueField

		setValueField = NumberInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Width")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%f", movementshape.Width))
		initializerStatements += setValueField

		setValueField = NumberInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Height")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%f", movementshape.Height))
		initializerStatements += setValueField

	}

	// insertion initialization of objects to stage
	if len(artefacttypeOrdered) > 0 {
		pointersInitializesStatements += "\n\t// setup of ArtefactType instances pointers"
	}
	for idx, artefacttype := range artefacttypeOrdered {
		var setPointerField string
		_ = setPointerField

		id = generatesIdentifier("ArtefactType", idx, artefacttype.Name)
		map_ArtefactType_Identifiers[artefacttype] = id

		// Initialisation of values
	}

	if len(artefacttypeshapeOrdered) > 0 {
		pointersInitializesStatements += "\n\t// setup of ArtefactTypeShape instances pointers"
	}
	for idx, artefacttypeshape := range artefacttypeshapeOrdered {
		var setPointerField string
		_ = setPointerField

		id = generatesIdentifier("ArtefactTypeShape", idx, artefacttypeshape.Name)
		map_ArtefactTypeShape_Identifiers[artefacttypeshape] = id

		// Initialisation of values
		if artefacttypeshape.ArtefactType != nil {
			setPointerField = PointerFieldInitStatement
			setPointerField = strings.ReplaceAll(setPointerField, "{{Identifier}}", id)
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldName}}", "ArtefactType")
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldNameValue}}", map_ArtefactType_Identifiers[artefacttypeshape.ArtefactType])
			pointersInitializesStatements += setPointerField
		}

	}

	if len(artistOrdered) > 0 {
		pointersInitializesStatements += "\n\t// setup of Artist instances pointers"
	}
	for idx, artist := range artistOrdered {
		var setPointerField string
		_ = setPointerField

		id = generatesIdentifier("Artist", idx, artist.Name)
		map_Artist_Identifiers[artist] = id

		// Initialisation of values
	}

	if len(artistshapeOrdered) > 0 {
		pointersInitializesStatements += "\n\t// setup of ArtistShape instances pointers"
	}
	for idx, artistshape := range artistshapeOrdered {
		var setPointerField string
		_ = setPointerField

		id = generatesIdentifier("ArtistShape", idx, artistshape.Name)
		map_ArtistShape_Identifiers[artistshape] = id

		// Initialisation of values
		if artistshape.Artist != nil {
			setPointerField = PointerFieldInitStatement
			setPointerField = strings.ReplaceAll(setPointerField, "{{Identifier}}", id)
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldName}}", "Artist")
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldNameValue}}", map_Artist_Identifiers[artistshape.Artist])
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
		for _, _movementshape := range diagram.MovementShapes {
			setPointerField = SliceOfPointersFieldInitStatement
			setPointerField = strings.ReplaceAll(setPointerField, "{{Identifier}}", id)
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldName}}", "MovementShapes")
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldNameValue}}", map_MovementShape_Identifiers[_movementshape])
			pointersInitializesStatements += setPointerField
		}

		for _, _artefacttypeshape := range diagram.ArtefactTypeShapes {
			setPointerField = SliceOfPointersFieldInitStatement
			setPointerField = strings.ReplaceAll(setPointerField, "{{Identifier}}", id)
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldName}}", "ArtefactTypeShapes")
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldNameValue}}", map_ArtefactTypeShape_Identifiers[_artefacttypeshape])
			pointersInitializesStatements += setPointerField
		}

		for _, _artistshape := range diagram.ArtistShapes {
			setPointerField = SliceOfPointersFieldInitStatement
			setPointerField = strings.ReplaceAll(setPointerField, "{{Identifier}}", id)
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldName}}", "ArtistShapes")
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldNameValue}}", map_ArtistShape_Identifiers[_artistshape])
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
		if influence.SourceMovement != nil {
			setPointerField = PointerFieldInitStatement
			setPointerField = strings.ReplaceAll(setPointerField, "{{Identifier}}", id)
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldName}}", "SourceMovement")
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldNameValue}}", map_Movement_Identifiers[influence.SourceMovement])
			pointersInitializesStatements += setPointerField
		}

		if influence.SourceArtefactType != nil {
			setPointerField = PointerFieldInitStatement
			setPointerField = strings.ReplaceAll(setPointerField, "{{Identifier}}", id)
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldName}}", "SourceArtefactType")
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldNameValue}}", map_ArtefactType_Identifiers[influence.SourceArtefactType])
			pointersInitializesStatements += setPointerField
		}

		if influence.SourceArtist != nil {
			setPointerField = PointerFieldInitStatement
			setPointerField = strings.ReplaceAll(setPointerField, "{{Identifier}}", id)
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldName}}", "SourceArtist")
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldNameValue}}", map_Artist_Identifiers[influence.SourceArtist])
			pointersInitializesStatements += setPointerField
		}

		if influence.TargetMovement != nil {
			setPointerField = PointerFieldInitStatement
			setPointerField = strings.ReplaceAll(setPointerField, "{{Identifier}}", id)
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldName}}", "TargetMovement")
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldNameValue}}", map_Movement_Identifiers[influence.TargetMovement])
			pointersInitializesStatements += setPointerField
		}

		if influence.TargetArtefactType != nil {
			setPointerField = PointerFieldInitStatement
			setPointerField = strings.ReplaceAll(setPointerField, "{{Identifier}}", id)
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldName}}", "TargetArtefactType")
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldNameValue}}", map_ArtefactType_Identifiers[influence.TargetArtefactType])
			pointersInitializesStatements += setPointerField
		}

		if influence.TargetArtist != nil {
			setPointerField = PointerFieldInitStatement
			setPointerField = strings.ReplaceAll(setPointerField, "{{Identifier}}", id)
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldName}}", "TargetArtist")
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldNameValue}}", map_Artist_Identifiers[influence.TargetArtist])
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

	if len(movementOrdered) > 0 {
		pointersInitializesStatements += "\n\t// setup of Movement instances pointers"
	}
	for idx, movement := range movementOrdered {
		var setPointerField string
		_ = setPointerField

		id = generatesIdentifier("Movement", idx, movement.Name)
		map_Movement_Identifiers[movement] = id

		// Initialisation of values
	}

	if len(movementshapeOrdered) > 0 {
		pointersInitializesStatements += "\n\t// setup of MovementShape instances pointers"
	}
	for idx, movementshape := range movementshapeOrdered {
		var setPointerField string
		_ = setPointerField

		id = generatesIdentifier("MovementShape", idx, movementshape.Name)
		map_MovementShape_Identifiers[movementshape] = id

		// Initialisation of values
		if movementshape.Movement != nil {
			setPointerField = PointerFieldInitStatement
			setPointerField = strings.ReplaceAll(setPointerField, "{{Identifier}}", id)
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldName}}", "Movement")
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldNameValue}}", map_Movement_Identifiers[movementshape.Movement])
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

	identifier = fmt.Sprintf("__%s__%06d_%s", gongStructName, idx, processedString)

	return
}
