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
	{{Identifier}} := (&models.{{GeneratedStructName}}{Name: {{GeneratedFieldNameValue}}}).Stage(stage)`

const GongUnstageStmt = `
	{{Identifier}}.Unstage(stage)`

// previous version does not hanldle embedded structs (https://github.com/golang/go/issues/9859)
// simpler version but the name of the instance cannot be human read before the fields initialization
const IdentifiersDeclsWithoutNameInit = `
	{{Identifier}} := (&models.{{GeneratedStructName}}{}).Stage(stage)` /* */

const StringInitStatement = `
	{{Identifier}}.{{GeneratedFieldName}} = {{GeneratedFieldNameValue}}`

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

// ToRawStringLiteral formats a string into safe Go source code,
// using backticks to preserve newlines and readability.
func ToRawStringLiteral(s string) string {
	// Step 1: Replace every backtick with a closing backtick,
	// a double-quoted backtick, and an opening backtick.
	escaped := strings.ReplaceAll(s, "`", "` + \"`\" + `")

	// Step 2: Wrap the entire resulting string in backticks.
	result := "`" + escaped + "`"

	// Step 3: Clean up any empty raw strings (``) at the boundaries
	// just in case your original string started or ended with a backtick.
	result = strings.ReplaceAll(result, "`` + ", "")
	result = strings.ReplaceAll(result, " + ``", "")

	return result
}

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
	artefacttypeOrdered := []*ArtefactType{}
	for artefacttype := range stage.ArtefactTypes {
		artefacttypeOrdered = append(artefacttypeOrdered, artefacttype)
	}
	sort.Slice(artefacttypeOrdered[:], func(i, j int) bool {
		artefacttypei := artefacttypeOrdered[i]
		artefacttypej := artefacttypeOrdered[j]
		artefacttypei_order, oki := stage.ArtefactType_stagedOrder[artefacttypei]
		artefacttypej_order, okj := stage.ArtefactType_stagedOrder[artefacttypej]
		if !oki || !okj {
			log.Fatalln("unknown pointers")
		}
		return artefacttypei_order < artefacttypej_order
	})
	if len(artefacttypeOrdered) > 0 {
		identifiersDecl.WriteString("\n")
	}
	for _, artefacttype := range artefacttypeOrdered {

		identifiersDecl.WriteString(artefacttype.GongMarshallIdentifier(stage))

		initializerStatements.WriteString("\n")
		// Insertion point for basic fields value assignment
		initializerStatements.WriteString(artefacttype.GongMarshallField(stage, "Name"))
	}

	artefacttypeshapeOrdered := []*ArtefactTypeShape{}
	for artefacttypeshape := range stage.ArtefactTypeShapes {
		artefacttypeshapeOrdered = append(artefacttypeshapeOrdered, artefacttypeshape)
	}
	sort.Slice(artefacttypeshapeOrdered[:], func(i, j int) bool {
		artefacttypeshapei := artefacttypeshapeOrdered[i]
		artefacttypeshapej := artefacttypeshapeOrdered[j]
		artefacttypeshapei_order, oki := stage.ArtefactTypeShape_stagedOrder[artefacttypeshapei]
		artefacttypeshapej_order, okj := stage.ArtefactTypeShape_stagedOrder[artefacttypeshapej]
		if !oki || !okj {
			log.Fatalln("unknown pointers")
		}
		return artefacttypeshapei_order < artefacttypeshapej_order
	})
	if len(artefacttypeshapeOrdered) > 0 {
		identifiersDecl.WriteString("\n")
	}
	for _, artefacttypeshape := range artefacttypeshapeOrdered {

		identifiersDecl.WriteString(artefacttypeshape.GongMarshallIdentifier(stage))

		initializerStatements.WriteString("\n")
		// Insertion point for basic fields value assignment
		initializerStatements.WriteString(artefacttypeshape.GongMarshallField(stage, "Name"))
		pointersInitializesStatements.WriteString(artefacttypeshape.GongMarshallField(stage, "ArtefactType"))
		initializerStatements.WriteString(artefacttypeshape.GongMarshallField(stage, "X"))
		initializerStatements.WriteString(artefacttypeshape.GongMarshallField(stage, "Y"))
		initializerStatements.WriteString(artefacttypeshape.GongMarshallField(stage, "Width"))
		initializerStatements.WriteString(artefacttypeshape.GongMarshallField(stage, "Height"))
	}

	artistOrdered := []*Artist{}
	for artist := range stage.Artists {
		artistOrdered = append(artistOrdered, artist)
	}
	sort.Slice(artistOrdered[:], func(i, j int) bool {
		artisti := artistOrdered[i]
		artistj := artistOrdered[j]
		artisti_order, oki := stage.Artist_stagedOrder[artisti]
		artistj_order, okj := stage.Artist_stagedOrder[artistj]
		if !oki || !okj {
			log.Fatalln("unknown pointers")
		}
		return artisti_order < artistj_order
	})
	if len(artistOrdered) > 0 {
		identifiersDecl.WriteString("\n")
	}
	for _, artist := range artistOrdered {

		identifiersDecl.WriteString(artist.GongMarshallIdentifier(stage))

		initializerStatements.WriteString("\n")
		// Insertion point for basic fields value assignment
		initializerStatements.WriteString(artist.GongMarshallField(stage, "Name"))
		initializerStatements.WriteString(artist.GongMarshallField(stage, "IsDead"))
		initializerStatements.WriteString(artist.GongMarshallField(stage, "DateOfDeath"))
		pointersInitializesStatements.WriteString(artist.GongMarshallField(stage, "Place"))
	}

	artistshapeOrdered := []*ArtistShape{}
	for artistshape := range stage.ArtistShapes {
		artistshapeOrdered = append(artistshapeOrdered, artistshape)
	}
	sort.Slice(artistshapeOrdered[:], func(i, j int) bool {
		artistshapei := artistshapeOrdered[i]
		artistshapej := artistshapeOrdered[j]
		artistshapei_order, oki := stage.ArtistShape_stagedOrder[artistshapei]
		artistshapej_order, okj := stage.ArtistShape_stagedOrder[artistshapej]
		if !oki || !okj {
			log.Fatalln("unknown pointers")
		}
		return artistshapei_order < artistshapej_order
	})
	if len(artistshapeOrdered) > 0 {
		identifiersDecl.WriteString("\n")
	}
	for _, artistshape := range artistshapeOrdered {

		identifiersDecl.WriteString(artistshape.GongMarshallIdentifier(stage))

		initializerStatements.WriteString("\n")
		// Insertion point for basic fields value assignment
		initializerStatements.WriteString(artistshape.GongMarshallField(stage, "Name"))
		pointersInitializesStatements.WriteString(artistshape.GongMarshallField(stage, "Artist"))
		initializerStatements.WriteString(artistshape.GongMarshallField(stage, "X"))
		initializerStatements.WriteString(artistshape.GongMarshallField(stage, "Y"))
		initializerStatements.WriteString(artistshape.GongMarshallField(stage, "Width"))
		initializerStatements.WriteString(artistshape.GongMarshallField(stage, "Height"))
	}

	controlpointshapeOrdered := []*ControlPointShape{}
	for controlpointshape := range stage.ControlPointShapes {
		controlpointshapeOrdered = append(controlpointshapeOrdered, controlpointshape)
	}
	sort.Slice(controlpointshapeOrdered[:], func(i, j int) bool {
		controlpointshapei := controlpointshapeOrdered[i]
		controlpointshapej := controlpointshapeOrdered[j]
		controlpointshapei_order, oki := stage.ControlPointShape_stagedOrder[controlpointshapei]
		controlpointshapej_order, okj := stage.ControlPointShape_stagedOrder[controlpointshapej]
		if !oki || !okj {
			log.Fatalln("unknown pointers")
		}
		return controlpointshapei_order < controlpointshapej_order
	})
	if len(controlpointshapeOrdered) > 0 {
		identifiersDecl.WriteString("\n")
	}
	for _, controlpointshape := range controlpointshapeOrdered {

		identifiersDecl.WriteString(controlpointshape.GongMarshallIdentifier(stage))

		initializerStatements.WriteString("\n")
		// Insertion point for basic fields value assignment
		initializerStatements.WriteString(controlpointshape.GongMarshallField(stage, "Name"))
		initializerStatements.WriteString(controlpointshape.GongMarshallField(stage, "X_Relative"))
		initializerStatements.WriteString(controlpointshape.GongMarshallField(stage, "Y_Relative"))
		initializerStatements.WriteString(controlpointshape.GongMarshallField(stage, "IsStartShapeTheClosestShape"))
	}

	deskOrdered := []*Desk{}
	for desk := range stage.Desks {
		deskOrdered = append(deskOrdered, desk)
	}
	sort.Slice(deskOrdered[:], func(i, j int) bool {
		deski := deskOrdered[i]
		deskj := deskOrdered[j]
		deski_order, oki := stage.Desk_stagedOrder[deski]
		deskj_order, okj := stage.Desk_stagedOrder[deskj]
		if !oki || !okj {
			log.Fatalln("unknown pointers")
		}
		return deski_order < deskj_order
	})
	if len(deskOrdered) > 0 {
		identifiersDecl.WriteString("\n")
	}
	for _, desk := range deskOrdered {

		identifiersDecl.WriteString(desk.GongMarshallIdentifier(stage))

		initializerStatements.WriteString("\n")
		// Insertion point for basic fields value assignment
		initializerStatements.WriteString(desk.GongMarshallField(stage, "Name"))
		pointersInitializesStatements.WriteString(desk.GongMarshallField(stage, "SelectedDiagram"))
	}

	diagramOrdered := []*Diagram{}
	for diagram := range stage.Diagrams {
		diagramOrdered = append(diagramOrdered, diagram)
	}
	sort.Slice(diagramOrdered[:], func(i, j int) bool {
		diagrami := diagramOrdered[i]
		diagramj := diagramOrdered[j]
		diagrami_order, oki := stage.Diagram_stagedOrder[diagrami]
		diagramj_order, okj := stage.Diagram_stagedOrder[diagramj]
		if !oki || !okj {
			log.Fatalln("unknown pointers")
		}
		return diagrami_order < diagramj_order
	})
	if len(diagramOrdered) > 0 {
		identifiersDecl.WriteString("\n")
	}
	for _, diagram := range diagramOrdered {

		identifiersDecl.WriteString(diagram.GongMarshallIdentifier(stage))

		initializerStatements.WriteString("\n")
		// Insertion point for basic fields value assignment
		initializerStatements.WriteString(diagram.GongMarshallField(stage, "Name"))
		pointersInitializesStatements.WriteString(diagram.GongMarshallField(stage, "MovementShapes"))
		pointersInitializesStatements.WriteString(diagram.GongMarshallField(stage, "ArtefactTypeShapes"))
		pointersInitializesStatements.WriteString(diagram.GongMarshallField(stage, "ArtistShapes"))
		pointersInitializesStatements.WriteString(diagram.GongMarshallField(stage, "InfluenceShapes"))
		initializerStatements.WriteString(diagram.GongMarshallField(stage, "IsEditable"))
		initializerStatements.WriteString(diagram.GongMarshallField(stage, "IsNodeExpanded"))
		initializerStatements.WriteString(diagram.GongMarshallField(stage, "IsMovementCategoryNodeExpanded"))
		initializerStatements.WriteString(diagram.GongMarshallField(stage, "IsArtefactTypeCategoryNodeExpanded"))
		initializerStatements.WriteString(diagram.GongMarshallField(stage, "IsArtistCategoryNodeExpanded"))
		initializerStatements.WriteString(diagram.GongMarshallField(stage, "IsInfluenceCategoryNodeExpanded"))
		initializerStatements.WriteString(diagram.GongMarshallField(stage, "IsMovementCategoryShown"))
		initializerStatements.WriteString(diagram.GongMarshallField(stage, "IsArtefactTypeCategoryShown"))
		initializerStatements.WriteString(diagram.GongMarshallField(stage, "IsArtistCategoryShown"))
		initializerStatements.WriteString(diagram.GongMarshallField(stage, "IsInfluenceCategoryShown"))
		initializerStatements.WriteString(diagram.GongMarshallField(stage, "StartDate"))
		initializerStatements.WriteString(diagram.GongMarshallField(stage, "EndDate"))
		initializerStatements.WriteString(diagram.GongMarshallField(stage, "XMargin"))
		initializerStatements.WriteString(diagram.GongMarshallField(stage, "YMargin"))
		initializerStatements.WriteString(diagram.GongMarshallField(stage, "Height"))
		initializerStatements.WriteString(diagram.GongMarshallField(stage, "NextVerticalDateXMargin"))
		initializerStatements.WriteString(diagram.GongMarshallField(stage, "RedColorCode"))
		initializerStatements.WriteString(diagram.GongMarshallField(stage, "BackgroundGreyColorCode"))
		initializerStatements.WriteString(diagram.GongMarshallField(stage, "GrayColorCode"))
		initializerStatements.WriteString(diagram.GongMarshallField(stage, "BottomBoxYOffset"))
		initializerStatements.WriteString(diagram.GongMarshallField(stage, "BottomBoxWidth"))
		initializerStatements.WriteString(diagram.GongMarshallField(stage, "BottomBoxHeigth"))
		initializerStatements.WriteString(diagram.GongMarshallField(stage, "BottomBoxFontSize"))
		initializerStatements.WriteString(diagram.GongMarshallField(stage, "BottomBoxFontWeigth"))
		initializerStatements.WriteString(diagram.GongMarshallField(stage, "BottomBoxFontFamily"))
		initializerStatements.WriteString(diagram.GongMarshallField(stage, "BottomBoxLetterSpacing"))
		initializerStatements.WriteString(diagram.GongMarshallField(stage, "BottomBoxLetterColorCode"))
		initializerStatements.WriteString(diagram.GongMarshallField(stage, "MovementRectAnchorType"))
		initializerStatements.WriteString(diagram.GongMarshallField(stage, "MovementTextAnchorType"))
		initializerStatements.WriteString(diagram.GongMarshallField(stage, "MovementDominantBaselineType"))
		initializerStatements.WriteString(diagram.GongMarshallField(stage, "MovementFontSize"))
		initializerStatements.WriteString(diagram.GongMarshallField(stage, "MajorMovementFontSize"))
		initializerStatements.WriteString(diagram.GongMarshallField(stage, "MinorMovementFontSize"))
		initializerStatements.WriteString(diagram.GongMarshallField(stage, "MovementFontWeigth"))
		initializerStatements.WriteString(diagram.GongMarshallField(stage, "MovementFontFamily"))
		initializerStatements.WriteString(diagram.GongMarshallField(stage, "MovementLetterSpacing"))
		initializerStatements.WriteString(diagram.GongMarshallField(stage, "AbstractMovementFontSize"))
		initializerStatements.WriteString(diagram.GongMarshallField(stage, "AbstractMovementRectAnchorType"))
		initializerStatements.WriteString(diagram.GongMarshallField(stage, "AbstractMovementTextAnchorType"))
		initializerStatements.WriteString(diagram.GongMarshallField(stage, "AbstractDominantBaselineType"))
		initializerStatements.WriteString(diagram.GongMarshallField(stage, "MovementDateRectAnchorType"))
		initializerStatements.WriteString(diagram.GongMarshallField(stage, "MovementDateTextAnchorType"))
		initializerStatements.WriteString(diagram.GongMarshallField(stage, "MovementDateTextDominantBaselineType"))
		initializerStatements.WriteString(diagram.GongMarshallField(stage, "MovementDateAndPlacesFontSize"))
		initializerStatements.WriteString(diagram.GongMarshallField(stage, "MovementDateAndPlacesFontWeigth"))
		initializerStatements.WriteString(diagram.GongMarshallField(stage, "MovementDateAndPlacesFontFamily"))
		initializerStatements.WriteString(diagram.GongMarshallField(stage, "MovementDateAndPlacesLetterSpacing"))
		initializerStatements.WriteString(diagram.GongMarshallField(stage, "MovementBelowArcY_Offset"))
		initializerStatements.WriteString(diagram.GongMarshallField(stage, "MovementBelowArcY_OffsetPerPlace"))
		initializerStatements.WriteString(diagram.GongMarshallField(stage, "MovementPlacesRectAnchorType"))
		initializerStatements.WriteString(diagram.GongMarshallField(stage, "MovementPlacesTextAnchorType"))
		initializerStatements.WriteString(diagram.GongMarshallField(stage, "MovementPlacesDominantBaselineType"))
		initializerStatements.WriteString(diagram.GongMarshallField(stage, "ArtefactTypeFontSize"))
		initializerStatements.WriteString(diagram.GongMarshallField(stage, "ArtefactTypeFontWeigth"))
		initializerStatements.WriteString(diagram.GongMarshallField(stage, "ArtefactTypeFontFamily"))
		initializerStatements.WriteString(diagram.GongMarshallField(stage, "ArtefactTypeLetterSpacing"))
		initializerStatements.WriteString(diagram.GongMarshallField(stage, "ArtefactTypeRectAnchorType"))
		initializerStatements.WriteString(diagram.GongMarshallField(stage, "ArtefactDominantBaselineType"))
		initializerStatements.WriteString(diagram.GongMarshallField(stage, "ArtefactTypeStrokeWidth"))
		initializerStatements.WriteString(diagram.GongMarshallField(stage, "ArtistRectAnchorType"))
		initializerStatements.WriteString(diagram.GongMarshallField(stage, "ArtistTextAnchorType"))
		initializerStatements.WriteString(diagram.GongMarshallField(stage, "ArtistDominantBaselineType"))
		initializerStatements.WriteString(diagram.GongMarshallField(stage, "ArtistFontSize"))
		initializerStatements.WriteString(diagram.GongMarshallField(stage, "MajorArtistFontSize"))
		initializerStatements.WriteString(diagram.GongMarshallField(stage, "MinorArtistFontSize"))
		initializerStatements.WriteString(diagram.GongMarshallField(stage, "ArtistFontWeigth"))
		initializerStatements.WriteString(diagram.GongMarshallField(stage, "ArtistFontFamily"))
		initializerStatements.WriteString(diagram.GongMarshallField(stage, "ArtistLetterSpacing"))
		initializerStatements.WriteString(diagram.GongMarshallField(stage, "ArtistDateRectAnchorType"))
		initializerStatements.WriteString(diagram.GongMarshallField(stage, "ArtistDateTextAnchorType"))
		initializerStatements.WriteString(diagram.GongMarshallField(stage, "ArtistDateDominantBaselineType"))
		initializerStatements.WriteString(diagram.GongMarshallField(stage, "ArtistDateAndPlacesFontSize"))
		initializerStatements.WriteString(diagram.GongMarshallField(stage, "ArtistDateAndPlacesFontWeigth"))
		initializerStatements.WriteString(diagram.GongMarshallField(stage, "ArtistDateAndPlacesFontFamily"))
		initializerStatements.WriteString(diagram.GongMarshallField(stage, "ArtistDateAndPlacesLetterSpacing"))
		initializerStatements.WriteString(diagram.GongMarshallField(stage, "ArtistPlacesRectAnchorType"))
		initializerStatements.WriteString(diagram.GongMarshallField(stage, "ArtistPlacesTextAnchorType"))
		initializerStatements.WriteString(diagram.GongMarshallField(stage, "ArtistPlacesDominantBaselineType"))
		initializerStatements.WriteString(diagram.GongMarshallField(stage, "InfluenceArrowSize"))
		initializerStatements.WriteString(diagram.GongMarshallField(stage, "InfluenceArrowStartOffset"))
		initializerStatements.WriteString(diagram.GongMarshallField(stage, "InfluenceArrowEndOffset"))
		initializerStatements.WriteString(diagram.GongMarshallField(stage, "InfluenceCornerRadius"))
		initializerStatements.WriteString(diagram.GongMarshallField(stage, "InfluenceDashedLinePattern"))
	}

	influenceOrdered := []*Influence{}
	for influence := range stage.Influences {
		influenceOrdered = append(influenceOrdered, influence)
	}
	sort.Slice(influenceOrdered[:], func(i, j int) bool {
		influencei := influenceOrdered[i]
		influencej := influenceOrdered[j]
		influencei_order, oki := stage.Influence_stagedOrder[influencei]
		influencej_order, okj := stage.Influence_stagedOrder[influencej]
		if !oki || !okj {
			log.Fatalln("unknown pointers")
		}
		return influencei_order < influencej_order
	})
	if len(influenceOrdered) > 0 {
		identifiersDecl.WriteString("\n")
	}
	for _, influence := range influenceOrdered {

		identifiersDecl.WriteString(influence.GongMarshallIdentifier(stage))

		initializerStatements.WriteString("\n")
		// Insertion point for basic fields value assignment
		initializerStatements.WriteString(influence.GongMarshallField(stage, "Name"))
		pointersInitializesStatements.WriteString(influence.GongMarshallField(stage, "SourceMovement"))
		pointersInitializesStatements.WriteString(influence.GongMarshallField(stage, "SourceArtefactType"))
		pointersInitializesStatements.WriteString(influence.GongMarshallField(stage, "SourceArtist"))
		pointersInitializesStatements.WriteString(influence.GongMarshallField(stage, "TargetMovement"))
		pointersInitializesStatements.WriteString(influence.GongMarshallField(stage, "TargetArtefactType"))
		pointersInitializesStatements.WriteString(influence.GongMarshallField(stage, "TargetArtist"))
		initializerStatements.WriteString(influence.GongMarshallField(stage, "IsHypothtical"))
	}

	influenceshapeOrdered := []*InfluenceShape{}
	for influenceshape := range stage.InfluenceShapes {
		influenceshapeOrdered = append(influenceshapeOrdered, influenceshape)
	}
	sort.Slice(influenceshapeOrdered[:], func(i, j int) bool {
		influenceshapei := influenceshapeOrdered[i]
		influenceshapej := influenceshapeOrdered[j]
		influenceshapei_order, oki := stage.InfluenceShape_stagedOrder[influenceshapei]
		influenceshapej_order, okj := stage.InfluenceShape_stagedOrder[influenceshapej]
		if !oki || !okj {
			log.Fatalln("unknown pointers")
		}
		return influenceshapei_order < influenceshapej_order
	})
	if len(influenceshapeOrdered) > 0 {
		identifiersDecl.WriteString("\n")
	}
	for _, influenceshape := range influenceshapeOrdered {

		identifiersDecl.WriteString(influenceshape.GongMarshallIdentifier(stage))

		initializerStatements.WriteString("\n")
		// Insertion point for basic fields value assignment
		initializerStatements.WriteString(influenceshape.GongMarshallField(stage, "Name"))
		pointersInitializesStatements.WriteString(influenceshape.GongMarshallField(stage, "Influence"))
		pointersInitializesStatements.WriteString(influenceshape.GongMarshallField(stage, "ControlPointShapes"))
	}

	movementOrdered := []*Movement{}
	for movement := range stage.Movements {
		movementOrdered = append(movementOrdered, movement)
	}
	sort.Slice(movementOrdered[:], func(i, j int) bool {
		movementi := movementOrdered[i]
		movementj := movementOrdered[j]
		movementi_order, oki := stage.Movement_stagedOrder[movementi]
		movementj_order, okj := stage.Movement_stagedOrder[movementj]
		if !oki || !okj {
			log.Fatalln("unknown pointers")
		}
		return movementi_order < movementj_order
	})
	if len(movementOrdered) > 0 {
		identifiersDecl.WriteString("\n")
	}
	for _, movement := range movementOrdered {

		identifiersDecl.WriteString(movement.GongMarshallIdentifier(stage))

		initializerStatements.WriteString("\n")
		// Insertion point for basic fields value assignment
		initializerStatements.WriteString(movement.GongMarshallField(stage, "Name"))
		initializerStatements.WriteString(movement.GongMarshallField(stage, "Date"))
		pointersInitializesStatements.WriteString(movement.GongMarshallField(stage, "Places"))
		initializerStatements.WriteString(movement.GongMarshallField(stage, "IsAbstract"))
		initializerStatements.WriteString(movement.GongMarshallField(stage, "IsModern"))
		initializerStatements.WriteString(movement.GongMarshallField(stage, "IsMajor"))
		initializerStatements.WriteString(movement.GongMarshallField(stage, "IsMinor"))
		initializerStatements.WriteString(movement.GongMarshallField(stage, "AdditionnalName"))
		initializerStatements.WriteString(movement.GongMarshallField(stage, "HideDate"))
	}

	movementshapeOrdered := []*MovementShape{}
	for movementshape := range stage.MovementShapes {
		movementshapeOrdered = append(movementshapeOrdered, movementshape)
	}
	sort.Slice(movementshapeOrdered[:], func(i, j int) bool {
		movementshapei := movementshapeOrdered[i]
		movementshapej := movementshapeOrdered[j]
		movementshapei_order, oki := stage.MovementShape_stagedOrder[movementshapei]
		movementshapej_order, okj := stage.MovementShape_stagedOrder[movementshapej]
		if !oki || !okj {
			log.Fatalln("unknown pointers")
		}
		return movementshapei_order < movementshapej_order
	})
	if len(movementshapeOrdered) > 0 {
		identifiersDecl.WriteString("\n")
	}
	for _, movementshape := range movementshapeOrdered {

		identifiersDecl.WriteString(movementshape.GongMarshallIdentifier(stage))

		initializerStatements.WriteString("\n")
		// Insertion point for basic fields value assignment
		initializerStatements.WriteString(movementshape.GongMarshallField(stage, "Name"))
		pointersInitializesStatements.WriteString(movementshape.GongMarshallField(stage, "Movement"))
		initializerStatements.WriteString(movementshape.GongMarshallField(stage, "X"))
		initializerStatements.WriteString(movementshape.GongMarshallField(stage, "Y"))
		initializerStatements.WriteString(movementshape.GongMarshallField(stage, "Width"))
		initializerStatements.WriteString(movementshape.GongMarshallField(stage, "Height"))
	}

	placeOrdered := []*Place{}
	for place := range stage.Places {
		placeOrdered = append(placeOrdered, place)
	}
	sort.Slice(placeOrdered[:], func(i, j int) bool {
		placei := placeOrdered[i]
		placej := placeOrdered[j]
		placei_order, oki := stage.Place_stagedOrder[placei]
		placej_order, okj := stage.Place_stagedOrder[placej]
		if !oki || !okj {
			log.Fatalln("unknown pointers")
		}
		return placei_order < placej_order
	})
	if len(placeOrdered) > 0 {
		identifiersDecl.WriteString("\n")
	}
	for _, place := range placeOrdered {

		identifiersDecl.WriteString(place.GongMarshallIdentifier(stage))

		initializerStatements.WriteString("\n")
		// Insertion point for basic fields value assignment
		initializerStatements.WriteString(place.GongMarshallField(stage, "Name"))
	}

	// insertion initialization of objects to stage
	for _, artefacttype := range artefacttypeOrdered {
		_ = artefacttype
		var setPointerField string
		_ = setPointerField

		// Insertion point for pointers initialization
	}

	for _, artefacttypeshape := range artefacttypeshapeOrdered {
		_ = artefacttypeshape
		var setPointerField string
		_ = setPointerField

		// Insertion point for pointers initialization
	}

	for _, artist := range artistOrdered {
		_ = artist
		var setPointerField string
		_ = setPointerField

		// Insertion point for pointers initialization
	}

	for _, artistshape := range artistshapeOrdered {
		_ = artistshape
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

	for _, movement := range movementOrdered {
		_ = movement
		var setPointerField string
		_ = setPointerField

		// Insertion point for pointers initialization
	}

	for _, movementshape := range movementshapeOrdered {
		_ = movementshape
		var setPointerField string
		_ = setPointerField

		// Insertion point for pointers initialization
	}

	for _, place := range placeOrdered {
		_ = place
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
func (artefacttype *ArtefactType) GongMarshallField(stage *Stage, fieldName string) (res string) {

	switch fieldName {
	case "Name":
		res = StringInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", artefacttype.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "Name")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", ToRawStringLiteral(artefacttype.Name))

	default:
		log.Panicf("Unknown field %s for Gongstruct ArtefactType", fieldName)
	}
	return
}

func (artefacttypeshape *ArtefactTypeShape) GongMarshallField(stage *Stage, fieldName string) (res string) {

	switch fieldName {
	case "Name":
		res = StringInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", artefacttypeshape.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "Name")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", ToRawStringLiteral(artefacttypeshape.Name))
	case "X":
		res = NumberInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", artefacttypeshape.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "X")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%f", artefacttypeshape.X))
	case "Y":
		res = NumberInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", artefacttypeshape.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "Y")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%f", artefacttypeshape.Y))
	case "Width":
		res = NumberInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", artefacttypeshape.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "Width")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%f", artefacttypeshape.Width))
	case "Height":
		res = NumberInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", artefacttypeshape.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "Height")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%f", artefacttypeshape.Height))

	case "ArtefactType":
		if artefacttypeshape.ArtefactType != nil {
			res = PointerFieldInitStatement
			res = strings.ReplaceAll(res, "{{Identifier}}", artefacttypeshape.GongGetIdentifier(stage))
			res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "ArtefactType")
			res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", artefacttypeshape.ArtefactType.GongGetIdentifier(stage))
		} else {
			// in case of nil pointer, we need to unstage the previous value
			res = PointerFieldInitStatement
			res = strings.ReplaceAll(res, "{{Identifier}}", artefacttypeshape.GongGetIdentifier(stage))
			res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "ArtefactType")
			res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", "nil")
		}
	default:
		log.Panicf("Unknown field %s for Gongstruct ArtefactTypeShape", fieldName)
	}
	return
}

func (artist *Artist) GongMarshallField(stage *Stage, fieldName string) (res string) {

	switch fieldName {
	case "Name":
		res = StringInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", artist.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "Name")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", ToRawStringLiteral(artist.Name))
	case "IsDead":
		res = NumberInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", artist.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "IsDead")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%t", artist.IsDead))
	case "DateOfDeath":
		res = TimeInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", artist.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "DateOfDeath")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", artist.DateOfDeath.String())

	case "Place":
		if artist.Place != nil {
			res = PointerFieldInitStatement
			res = strings.ReplaceAll(res, "{{Identifier}}", artist.GongGetIdentifier(stage))
			res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "Place")
			res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", artist.Place.GongGetIdentifier(stage))
		} else {
			// in case of nil pointer, we need to unstage the previous value
			res = PointerFieldInitStatement
			res = strings.ReplaceAll(res, "{{Identifier}}", artist.GongGetIdentifier(stage))
			res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "Place")
			res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", "nil")
		}
	default:
		log.Panicf("Unknown field %s for Gongstruct Artist", fieldName)
	}
	return
}

func (artistshape *ArtistShape) GongMarshallField(stage *Stage, fieldName string) (res string) {

	switch fieldName {
	case "Name":
		res = StringInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", artistshape.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "Name")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", ToRawStringLiteral(artistshape.Name))
	case "X":
		res = NumberInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", artistshape.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "X")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%f", artistshape.X))
	case "Y":
		res = NumberInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", artistshape.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "Y")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%f", artistshape.Y))
	case "Width":
		res = NumberInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", artistshape.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "Width")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%f", artistshape.Width))
	case "Height":
		res = NumberInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", artistshape.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "Height")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%f", artistshape.Height))

	case "Artist":
		if artistshape.Artist != nil {
			res = PointerFieldInitStatement
			res = strings.ReplaceAll(res, "{{Identifier}}", artistshape.GongGetIdentifier(stage))
			res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "Artist")
			res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", artistshape.Artist.GongGetIdentifier(stage))
		} else {
			// in case of nil pointer, we need to unstage the previous value
			res = PointerFieldInitStatement
			res = strings.ReplaceAll(res, "{{Identifier}}", artistshape.GongGetIdentifier(stage))
			res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "Artist")
			res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", "nil")
		}
	default:
		log.Panicf("Unknown field %s for Gongstruct ArtistShape", fieldName)
	}
	return
}

func (controlpointshape *ControlPointShape) GongMarshallField(stage *Stage, fieldName string) (res string) {

	switch fieldName {
	case "Name":
		res = StringInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", controlpointshape.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "Name")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", ToRawStringLiteral(controlpointshape.Name))
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
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", ToRawStringLiteral(desk.Name))

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
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", ToRawStringLiteral(diagram.Name))
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
	case "IsMovementCategoryNodeExpanded":
		res = NumberInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", diagram.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "IsMovementCategoryNodeExpanded")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%t", diagram.IsMovementCategoryNodeExpanded))
	case "IsArtefactTypeCategoryNodeExpanded":
		res = NumberInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", diagram.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "IsArtefactTypeCategoryNodeExpanded")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%t", diagram.IsArtefactTypeCategoryNodeExpanded))
	case "IsArtistCategoryNodeExpanded":
		res = NumberInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", diagram.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "IsArtistCategoryNodeExpanded")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%t", diagram.IsArtistCategoryNodeExpanded))
	case "IsInfluenceCategoryNodeExpanded":
		res = NumberInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", diagram.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "IsInfluenceCategoryNodeExpanded")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%t", diagram.IsInfluenceCategoryNodeExpanded))
	case "IsMovementCategoryShown":
		res = NumberInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", diagram.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "IsMovementCategoryShown")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%t", diagram.IsMovementCategoryShown))
	case "IsArtefactTypeCategoryShown":
		res = NumberInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", diagram.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "IsArtefactTypeCategoryShown")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%t", diagram.IsArtefactTypeCategoryShown))
	case "IsArtistCategoryShown":
		res = NumberInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", diagram.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "IsArtistCategoryShown")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%t", diagram.IsArtistCategoryShown))
	case "IsInfluenceCategoryShown":
		res = NumberInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", diagram.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "IsInfluenceCategoryShown")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%t", diagram.IsInfluenceCategoryShown))
	case "StartDate":
		res = TimeInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", diagram.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "StartDate")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", diagram.StartDate.String())
	case "EndDate":
		res = TimeInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", diagram.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "EndDate")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", diagram.EndDate.String())
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
	case "NextVerticalDateXMargin":
		res = NumberInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", diagram.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "NextVerticalDateXMargin")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%f", diagram.NextVerticalDateXMargin))
	case "RedColorCode":
		res = StringInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", diagram.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "RedColorCode")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", ToRawStringLiteral(diagram.RedColorCode))
	case "BackgroundGreyColorCode":
		res = StringInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", diagram.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "BackgroundGreyColorCode")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", ToRawStringLiteral(diagram.BackgroundGreyColorCode))
	case "GrayColorCode":
		res = StringInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", diagram.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "GrayColorCode")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", ToRawStringLiteral(diagram.GrayColorCode))
	case "BottomBoxYOffset":
		res = NumberInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", diagram.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "BottomBoxYOffset")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%f", diagram.BottomBoxYOffset))
	case "BottomBoxWidth":
		res = NumberInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", diagram.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "BottomBoxWidth")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%f", diagram.BottomBoxWidth))
	case "BottomBoxHeigth":
		res = NumberInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", diagram.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "BottomBoxHeigth")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%f", diagram.BottomBoxHeigth))
	case "BottomBoxFontSize":
		res = StringInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", diagram.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "BottomBoxFontSize")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", ToRawStringLiteral(diagram.BottomBoxFontSize))
	case "BottomBoxFontWeigth":
		res = StringInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", diagram.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "BottomBoxFontWeigth")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", ToRawStringLiteral(diagram.BottomBoxFontWeigth))
	case "BottomBoxFontFamily":
		res = StringInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", diagram.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "BottomBoxFontFamily")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", ToRawStringLiteral(diagram.BottomBoxFontFamily))
	case "BottomBoxLetterSpacing":
		res = StringInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", diagram.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "BottomBoxLetterSpacing")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", ToRawStringLiteral(diagram.BottomBoxLetterSpacing))
	case "BottomBoxLetterColorCode":
		res = StringInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", diagram.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "BottomBoxLetterColorCode")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", ToRawStringLiteral(diagram.BottomBoxLetterColorCode))
	case "MovementRectAnchorType":
		if diagram.MovementRectAnchorType.ToCodeString() != "" {
			res = StringEnumInitStatement
			res = strings.ReplaceAll(res, "{{Identifier}}", diagram.GongGetIdentifier(stage))
			res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "MovementRectAnchorType")
			res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", "models."+diagram.MovementRectAnchorType.ToCodeString())
		} else {
			// in case of empty enum, we need to unstage the previous value
			res = StringEnumInitStatement
			res = strings.ReplaceAll(res, "{{Identifier}}", diagram.GongGetIdentifier(stage))
			res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "MovementRectAnchorType")
			res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", "\"\"")
		}
	case "MovementTextAnchorType":
		if diagram.MovementTextAnchorType.ToCodeString() != "" {
			res = StringEnumInitStatement
			res = strings.ReplaceAll(res, "{{Identifier}}", diagram.GongGetIdentifier(stage))
			res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "MovementTextAnchorType")
			res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", "models."+diagram.MovementTextAnchorType.ToCodeString())
		} else {
			// in case of empty enum, we need to unstage the previous value
			res = StringEnumInitStatement
			res = strings.ReplaceAll(res, "{{Identifier}}", diagram.GongGetIdentifier(stage))
			res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "MovementTextAnchorType")
			res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", "\"\"")
		}
	case "MovementDominantBaselineType":
		if diagram.MovementDominantBaselineType.ToCodeString() != "" {
			res = StringEnumInitStatement
			res = strings.ReplaceAll(res, "{{Identifier}}", diagram.GongGetIdentifier(stage))
			res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "MovementDominantBaselineType")
			res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", "models."+diagram.MovementDominantBaselineType.ToCodeString())
		} else {
			// in case of empty enum, we need to unstage the previous value
			res = StringEnumInitStatement
			res = strings.ReplaceAll(res, "{{Identifier}}", diagram.GongGetIdentifier(stage))
			res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "MovementDominantBaselineType")
			res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", "\"\"")
		}
	case "MovementFontSize":
		res = StringInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", diagram.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "MovementFontSize")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", ToRawStringLiteral(diagram.MovementFontSize))
	case "MajorMovementFontSize":
		res = StringInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", diagram.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "MajorMovementFontSize")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", ToRawStringLiteral(diagram.MajorMovementFontSize))
	case "MinorMovementFontSize":
		res = StringInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", diagram.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "MinorMovementFontSize")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", ToRawStringLiteral(diagram.MinorMovementFontSize))
	case "MovementFontWeigth":
		res = StringInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", diagram.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "MovementFontWeigth")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", ToRawStringLiteral(diagram.MovementFontWeigth))
	case "MovementFontFamily":
		res = StringInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", diagram.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "MovementFontFamily")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", ToRawStringLiteral(diagram.MovementFontFamily))
	case "MovementLetterSpacing":
		res = StringInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", diagram.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "MovementLetterSpacing")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", ToRawStringLiteral(diagram.MovementLetterSpacing))
	case "AbstractMovementFontSize":
		res = StringInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", diagram.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "AbstractMovementFontSize")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", ToRawStringLiteral(diagram.AbstractMovementFontSize))
	case "AbstractMovementRectAnchorType":
		if diagram.AbstractMovementRectAnchorType.ToCodeString() != "" {
			res = StringEnumInitStatement
			res = strings.ReplaceAll(res, "{{Identifier}}", diagram.GongGetIdentifier(stage))
			res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "AbstractMovementRectAnchorType")
			res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", "models."+diagram.AbstractMovementRectAnchorType.ToCodeString())
		} else {
			// in case of empty enum, we need to unstage the previous value
			res = StringEnumInitStatement
			res = strings.ReplaceAll(res, "{{Identifier}}", diagram.GongGetIdentifier(stage))
			res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "AbstractMovementRectAnchorType")
			res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", "\"\"")
		}
	case "AbstractMovementTextAnchorType":
		if diagram.AbstractMovementTextAnchorType.ToCodeString() != "" {
			res = StringEnumInitStatement
			res = strings.ReplaceAll(res, "{{Identifier}}", diagram.GongGetIdentifier(stage))
			res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "AbstractMovementTextAnchorType")
			res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", "models."+diagram.AbstractMovementTextAnchorType.ToCodeString())
		} else {
			// in case of empty enum, we need to unstage the previous value
			res = StringEnumInitStatement
			res = strings.ReplaceAll(res, "{{Identifier}}", diagram.GongGetIdentifier(stage))
			res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "AbstractMovementTextAnchorType")
			res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", "\"\"")
		}
	case "AbstractDominantBaselineType":
		if diagram.AbstractDominantBaselineType.ToCodeString() != "" {
			res = StringEnumInitStatement
			res = strings.ReplaceAll(res, "{{Identifier}}", diagram.GongGetIdentifier(stage))
			res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "AbstractDominantBaselineType")
			res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", "models."+diagram.AbstractDominantBaselineType.ToCodeString())
		} else {
			// in case of empty enum, we need to unstage the previous value
			res = StringEnumInitStatement
			res = strings.ReplaceAll(res, "{{Identifier}}", diagram.GongGetIdentifier(stage))
			res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "AbstractDominantBaselineType")
			res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", "\"\"")
		}
	case "MovementDateRectAnchorType":
		if diagram.MovementDateRectAnchorType.ToCodeString() != "" {
			res = StringEnumInitStatement
			res = strings.ReplaceAll(res, "{{Identifier}}", diagram.GongGetIdentifier(stage))
			res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "MovementDateRectAnchorType")
			res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", "models."+diagram.MovementDateRectAnchorType.ToCodeString())
		} else {
			// in case of empty enum, we need to unstage the previous value
			res = StringEnumInitStatement
			res = strings.ReplaceAll(res, "{{Identifier}}", diagram.GongGetIdentifier(stage))
			res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "MovementDateRectAnchorType")
			res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", "\"\"")
		}
	case "MovementDateTextAnchorType":
		if diagram.MovementDateTextAnchorType.ToCodeString() != "" {
			res = StringEnumInitStatement
			res = strings.ReplaceAll(res, "{{Identifier}}", diagram.GongGetIdentifier(stage))
			res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "MovementDateTextAnchorType")
			res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", "models."+diagram.MovementDateTextAnchorType.ToCodeString())
		} else {
			// in case of empty enum, we need to unstage the previous value
			res = StringEnumInitStatement
			res = strings.ReplaceAll(res, "{{Identifier}}", diagram.GongGetIdentifier(stage))
			res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "MovementDateTextAnchorType")
			res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", "\"\"")
		}
	case "MovementDateTextDominantBaselineType":
		if diagram.MovementDateTextDominantBaselineType.ToCodeString() != "" {
			res = StringEnumInitStatement
			res = strings.ReplaceAll(res, "{{Identifier}}", diagram.GongGetIdentifier(stage))
			res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "MovementDateTextDominantBaselineType")
			res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", "models."+diagram.MovementDateTextDominantBaselineType.ToCodeString())
		} else {
			// in case of empty enum, we need to unstage the previous value
			res = StringEnumInitStatement
			res = strings.ReplaceAll(res, "{{Identifier}}", diagram.GongGetIdentifier(stage))
			res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "MovementDateTextDominantBaselineType")
			res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", "\"\"")
		}
	case "MovementDateAndPlacesFontSize":
		res = StringInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", diagram.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "MovementDateAndPlacesFontSize")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", ToRawStringLiteral(diagram.MovementDateAndPlacesFontSize))
	case "MovementDateAndPlacesFontWeigth":
		res = StringInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", diagram.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "MovementDateAndPlacesFontWeigth")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", ToRawStringLiteral(diagram.MovementDateAndPlacesFontWeigth))
	case "MovementDateAndPlacesFontFamily":
		res = StringInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", diagram.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "MovementDateAndPlacesFontFamily")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", ToRawStringLiteral(diagram.MovementDateAndPlacesFontFamily))
	case "MovementDateAndPlacesLetterSpacing":
		res = StringInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", diagram.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "MovementDateAndPlacesLetterSpacing")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", ToRawStringLiteral(diagram.MovementDateAndPlacesLetterSpacing))
	case "MovementBelowArcY_Offset":
		res = NumberInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", diagram.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "MovementBelowArcY_Offset")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%f", diagram.MovementBelowArcY_Offset))
	case "MovementBelowArcY_OffsetPerPlace":
		res = NumberInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", diagram.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "MovementBelowArcY_OffsetPerPlace")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%f", diagram.MovementBelowArcY_OffsetPerPlace))
	case "MovementPlacesRectAnchorType":
		if diagram.MovementPlacesRectAnchorType.ToCodeString() != "" {
			res = StringEnumInitStatement
			res = strings.ReplaceAll(res, "{{Identifier}}", diagram.GongGetIdentifier(stage))
			res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "MovementPlacesRectAnchorType")
			res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", "models."+diagram.MovementPlacesRectAnchorType.ToCodeString())
		} else {
			// in case of empty enum, we need to unstage the previous value
			res = StringEnumInitStatement
			res = strings.ReplaceAll(res, "{{Identifier}}", diagram.GongGetIdentifier(stage))
			res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "MovementPlacesRectAnchorType")
			res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", "\"\"")
		}
	case "MovementPlacesTextAnchorType":
		if diagram.MovementPlacesTextAnchorType.ToCodeString() != "" {
			res = StringEnumInitStatement
			res = strings.ReplaceAll(res, "{{Identifier}}", diagram.GongGetIdentifier(stage))
			res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "MovementPlacesTextAnchorType")
			res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", "models."+diagram.MovementPlacesTextAnchorType.ToCodeString())
		} else {
			// in case of empty enum, we need to unstage the previous value
			res = StringEnumInitStatement
			res = strings.ReplaceAll(res, "{{Identifier}}", diagram.GongGetIdentifier(stage))
			res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "MovementPlacesTextAnchorType")
			res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", "\"\"")
		}
	case "MovementPlacesDominantBaselineType":
		if diagram.MovementPlacesDominantBaselineType.ToCodeString() != "" {
			res = StringEnumInitStatement
			res = strings.ReplaceAll(res, "{{Identifier}}", diagram.GongGetIdentifier(stage))
			res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "MovementPlacesDominantBaselineType")
			res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", "models."+diagram.MovementPlacesDominantBaselineType.ToCodeString())
		} else {
			// in case of empty enum, we need to unstage the previous value
			res = StringEnumInitStatement
			res = strings.ReplaceAll(res, "{{Identifier}}", diagram.GongGetIdentifier(stage))
			res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "MovementPlacesDominantBaselineType")
			res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", "\"\"")
		}
	case "ArtefactTypeFontSize":
		res = StringInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", diagram.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "ArtefactTypeFontSize")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", ToRawStringLiteral(diagram.ArtefactTypeFontSize))
	case "ArtefactTypeFontWeigth":
		res = StringInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", diagram.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "ArtefactTypeFontWeigth")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", ToRawStringLiteral(diagram.ArtefactTypeFontWeigth))
	case "ArtefactTypeFontFamily":
		res = StringInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", diagram.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "ArtefactTypeFontFamily")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", ToRawStringLiteral(diagram.ArtefactTypeFontFamily))
	case "ArtefactTypeLetterSpacing":
		res = StringInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", diagram.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "ArtefactTypeLetterSpacing")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", ToRawStringLiteral(diagram.ArtefactTypeLetterSpacing))
	case "ArtefactTypeRectAnchorType":
		if diagram.ArtefactTypeRectAnchorType.ToCodeString() != "" {
			res = StringEnumInitStatement
			res = strings.ReplaceAll(res, "{{Identifier}}", diagram.GongGetIdentifier(stage))
			res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "ArtefactTypeRectAnchorType")
			res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", "models."+diagram.ArtefactTypeRectAnchorType.ToCodeString())
		} else {
			// in case of empty enum, we need to unstage the previous value
			res = StringEnumInitStatement
			res = strings.ReplaceAll(res, "{{Identifier}}", diagram.GongGetIdentifier(stage))
			res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "ArtefactTypeRectAnchorType")
			res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", "\"\"")
		}
	case "ArtefactDominantBaselineType":
		if diagram.ArtefactDominantBaselineType.ToCodeString() != "" {
			res = StringEnumInitStatement
			res = strings.ReplaceAll(res, "{{Identifier}}", diagram.GongGetIdentifier(stage))
			res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "ArtefactDominantBaselineType")
			res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", "models."+diagram.ArtefactDominantBaselineType.ToCodeString())
		} else {
			// in case of empty enum, we need to unstage the previous value
			res = StringEnumInitStatement
			res = strings.ReplaceAll(res, "{{Identifier}}", diagram.GongGetIdentifier(stage))
			res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "ArtefactDominantBaselineType")
			res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", "\"\"")
		}
	case "ArtefactTypeStrokeWidth":
		res = NumberInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", diagram.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "ArtefactTypeStrokeWidth")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%f", diagram.ArtefactTypeStrokeWidth))
	case "ArtistRectAnchorType":
		if diagram.ArtistRectAnchorType.ToCodeString() != "" {
			res = StringEnumInitStatement
			res = strings.ReplaceAll(res, "{{Identifier}}", diagram.GongGetIdentifier(stage))
			res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "ArtistRectAnchorType")
			res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", "models."+diagram.ArtistRectAnchorType.ToCodeString())
		} else {
			// in case of empty enum, we need to unstage the previous value
			res = StringEnumInitStatement
			res = strings.ReplaceAll(res, "{{Identifier}}", diagram.GongGetIdentifier(stage))
			res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "ArtistRectAnchorType")
			res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", "\"\"")
		}
	case "ArtistTextAnchorType":
		if diagram.ArtistTextAnchorType.ToCodeString() != "" {
			res = StringEnumInitStatement
			res = strings.ReplaceAll(res, "{{Identifier}}", diagram.GongGetIdentifier(stage))
			res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "ArtistTextAnchorType")
			res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", "models."+diagram.ArtistTextAnchorType.ToCodeString())
		} else {
			// in case of empty enum, we need to unstage the previous value
			res = StringEnumInitStatement
			res = strings.ReplaceAll(res, "{{Identifier}}", diagram.GongGetIdentifier(stage))
			res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "ArtistTextAnchorType")
			res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", "\"\"")
		}
	case "ArtistDominantBaselineType":
		if diagram.ArtistDominantBaselineType.ToCodeString() != "" {
			res = StringEnumInitStatement
			res = strings.ReplaceAll(res, "{{Identifier}}", diagram.GongGetIdentifier(stage))
			res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "ArtistDominantBaselineType")
			res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", "models."+diagram.ArtistDominantBaselineType.ToCodeString())
		} else {
			// in case of empty enum, we need to unstage the previous value
			res = StringEnumInitStatement
			res = strings.ReplaceAll(res, "{{Identifier}}", diagram.GongGetIdentifier(stage))
			res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "ArtistDominantBaselineType")
			res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", "\"\"")
		}
	case "ArtistFontSize":
		res = StringInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", diagram.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "ArtistFontSize")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", ToRawStringLiteral(diagram.ArtistFontSize))
	case "MajorArtistFontSize":
		res = StringInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", diagram.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "MajorArtistFontSize")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", ToRawStringLiteral(diagram.MajorArtistFontSize))
	case "MinorArtistFontSize":
		res = StringInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", diagram.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "MinorArtistFontSize")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", ToRawStringLiteral(diagram.MinorArtistFontSize))
	case "ArtistFontWeigth":
		res = StringInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", diagram.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "ArtistFontWeigth")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", ToRawStringLiteral(diagram.ArtistFontWeigth))
	case "ArtistFontFamily":
		res = StringInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", diagram.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "ArtistFontFamily")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", ToRawStringLiteral(diagram.ArtistFontFamily))
	case "ArtistLetterSpacing":
		res = StringInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", diagram.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "ArtistLetterSpacing")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", ToRawStringLiteral(diagram.ArtistLetterSpacing))
	case "ArtistDateRectAnchorType":
		if diagram.ArtistDateRectAnchorType.ToCodeString() != "" {
			res = StringEnumInitStatement
			res = strings.ReplaceAll(res, "{{Identifier}}", diagram.GongGetIdentifier(stage))
			res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "ArtistDateRectAnchorType")
			res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", "models."+diagram.ArtistDateRectAnchorType.ToCodeString())
		} else {
			// in case of empty enum, we need to unstage the previous value
			res = StringEnumInitStatement
			res = strings.ReplaceAll(res, "{{Identifier}}", diagram.GongGetIdentifier(stage))
			res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "ArtistDateRectAnchorType")
			res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", "\"\"")
		}
	case "ArtistDateTextAnchorType":
		if diagram.ArtistDateTextAnchorType.ToCodeString() != "" {
			res = StringEnumInitStatement
			res = strings.ReplaceAll(res, "{{Identifier}}", diagram.GongGetIdentifier(stage))
			res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "ArtistDateTextAnchorType")
			res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", "models."+diagram.ArtistDateTextAnchorType.ToCodeString())
		} else {
			// in case of empty enum, we need to unstage the previous value
			res = StringEnumInitStatement
			res = strings.ReplaceAll(res, "{{Identifier}}", diagram.GongGetIdentifier(stage))
			res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "ArtistDateTextAnchorType")
			res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", "\"\"")
		}
	case "ArtistDateDominantBaselineType":
		if diagram.ArtistDateDominantBaselineType.ToCodeString() != "" {
			res = StringEnumInitStatement
			res = strings.ReplaceAll(res, "{{Identifier}}", diagram.GongGetIdentifier(stage))
			res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "ArtistDateDominantBaselineType")
			res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", "models."+diagram.ArtistDateDominantBaselineType.ToCodeString())
		} else {
			// in case of empty enum, we need to unstage the previous value
			res = StringEnumInitStatement
			res = strings.ReplaceAll(res, "{{Identifier}}", diagram.GongGetIdentifier(stage))
			res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "ArtistDateDominantBaselineType")
			res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", "\"\"")
		}
	case "ArtistDateAndPlacesFontSize":
		res = StringInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", diagram.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "ArtistDateAndPlacesFontSize")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", ToRawStringLiteral(diagram.ArtistDateAndPlacesFontSize))
	case "ArtistDateAndPlacesFontWeigth":
		res = StringInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", diagram.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "ArtistDateAndPlacesFontWeigth")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", ToRawStringLiteral(diagram.ArtistDateAndPlacesFontWeigth))
	case "ArtistDateAndPlacesFontFamily":
		res = StringInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", diagram.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "ArtistDateAndPlacesFontFamily")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", ToRawStringLiteral(diagram.ArtistDateAndPlacesFontFamily))
	case "ArtistDateAndPlacesLetterSpacing":
		res = StringInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", diagram.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "ArtistDateAndPlacesLetterSpacing")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", ToRawStringLiteral(diagram.ArtistDateAndPlacesLetterSpacing))
	case "ArtistPlacesRectAnchorType":
		if diagram.ArtistPlacesRectAnchorType.ToCodeString() != "" {
			res = StringEnumInitStatement
			res = strings.ReplaceAll(res, "{{Identifier}}", diagram.GongGetIdentifier(stage))
			res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "ArtistPlacesRectAnchorType")
			res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", "models."+diagram.ArtistPlacesRectAnchorType.ToCodeString())
		} else {
			// in case of empty enum, we need to unstage the previous value
			res = StringEnumInitStatement
			res = strings.ReplaceAll(res, "{{Identifier}}", diagram.GongGetIdentifier(stage))
			res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "ArtistPlacesRectAnchorType")
			res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", "\"\"")
		}
	case "ArtistPlacesTextAnchorType":
		if diagram.ArtistPlacesTextAnchorType.ToCodeString() != "" {
			res = StringEnumInitStatement
			res = strings.ReplaceAll(res, "{{Identifier}}", diagram.GongGetIdentifier(stage))
			res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "ArtistPlacesTextAnchorType")
			res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", "models."+diagram.ArtistPlacesTextAnchorType.ToCodeString())
		} else {
			// in case of empty enum, we need to unstage the previous value
			res = StringEnumInitStatement
			res = strings.ReplaceAll(res, "{{Identifier}}", diagram.GongGetIdentifier(stage))
			res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "ArtistPlacesTextAnchorType")
			res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", "\"\"")
		}
	case "ArtistPlacesDominantBaselineType":
		if diagram.ArtistPlacesDominantBaselineType.ToCodeString() != "" {
			res = StringEnumInitStatement
			res = strings.ReplaceAll(res, "{{Identifier}}", diagram.GongGetIdentifier(stage))
			res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "ArtistPlacesDominantBaselineType")
			res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", "models."+diagram.ArtistPlacesDominantBaselineType.ToCodeString())
		} else {
			// in case of empty enum, we need to unstage the previous value
			res = StringEnumInitStatement
			res = strings.ReplaceAll(res, "{{Identifier}}", diagram.GongGetIdentifier(stage))
			res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "ArtistPlacesDominantBaselineType")
			res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", "\"\"")
		}
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
	case "InfluenceDashedLinePattern":
		res = StringInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", diagram.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "InfluenceDashedLinePattern")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", ToRawStringLiteral(diagram.InfluenceDashedLinePattern))

	case "MovementShapes":
		var sb strings.Builder
		for _, _movementshape := range diagram.MovementShapes {
			tmp := SliceOfPointersFieldInitStatement
			tmp = strings.ReplaceAll(tmp, "{{Identifier}}", diagram.GongGetIdentifier(stage))
			tmp = strings.ReplaceAll(tmp, "{{GeneratedFieldName}}", "MovementShapes")
			tmp = strings.ReplaceAll(tmp, "{{GeneratedFieldNameValue}}", _movementshape.GongGetIdentifier(stage))
			sb.WriteString(tmp)
		}
		res = sb.String()
	case "ArtefactTypeShapes":
		var sb strings.Builder
		for _, _artefacttypeshape := range diagram.ArtefactTypeShapes {
			tmp := SliceOfPointersFieldInitStatement
			tmp = strings.ReplaceAll(tmp, "{{Identifier}}", diagram.GongGetIdentifier(stage))
			tmp = strings.ReplaceAll(tmp, "{{GeneratedFieldName}}", "ArtefactTypeShapes")
			tmp = strings.ReplaceAll(tmp, "{{GeneratedFieldNameValue}}", _artefacttypeshape.GongGetIdentifier(stage))
			sb.WriteString(tmp)
		}
		res = sb.String()
	case "ArtistShapes":
		var sb strings.Builder
		for _, _artistshape := range diagram.ArtistShapes {
			tmp := SliceOfPointersFieldInitStatement
			tmp = strings.ReplaceAll(tmp, "{{Identifier}}", diagram.GongGetIdentifier(stage))
			tmp = strings.ReplaceAll(tmp, "{{GeneratedFieldName}}", "ArtistShapes")
			tmp = strings.ReplaceAll(tmp, "{{GeneratedFieldNameValue}}", _artistshape.GongGetIdentifier(stage))
			sb.WriteString(tmp)
		}
		res = sb.String()
	case "InfluenceShapes":
		var sb strings.Builder
		for _, _influenceshape := range diagram.InfluenceShapes {
			tmp := SliceOfPointersFieldInitStatement
			tmp = strings.ReplaceAll(tmp, "{{Identifier}}", diagram.GongGetIdentifier(stage))
			tmp = strings.ReplaceAll(tmp, "{{GeneratedFieldName}}", "InfluenceShapes")
			tmp = strings.ReplaceAll(tmp, "{{GeneratedFieldNameValue}}", _influenceshape.GongGetIdentifier(stage))
			sb.WriteString(tmp)
		}
		res = sb.String()
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
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", ToRawStringLiteral(influence.Name))
	case "IsHypothtical":
		res = NumberInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", influence.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "IsHypothtical")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%t", influence.IsHypothtical))

	case "SourceMovement":
		if influence.SourceMovement != nil {
			res = PointerFieldInitStatement
			res = strings.ReplaceAll(res, "{{Identifier}}", influence.GongGetIdentifier(stage))
			res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "SourceMovement")
			res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", influence.SourceMovement.GongGetIdentifier(stage))
		} else {
			// in case of nil pointer, we need to unstage the previous value
			res = PointerFieldInitStatement
			res = strings.ReplaceAll(res, "{{Identifier}}", influence.GongGetIdentifier(stage))
			res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "SourceMovement")
			res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", "nil")
		}
	case "SourceArtefactType":
		if influence.SourceArtefactType != nil {
			res = PointerFieldInitStatement
			res = strings.ReplaceAll(res, "{{Identifier}}", influence.GongGetIdentifier(stage))
			res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "SourceArtefactType")
			res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", influence.SourceArtefactType.GongGetIdentifier(stage))
		} else {
			// in case of nil pointer, we need to unstage the previous value
			res = PointerFieldInitStatement
			res = strings.ReplaceAll(res, "{{Identifier}}", influence.GongGetIdentifier(stage))
			res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "SourceArtefactType")
			res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", "nil")
		}
	case "SourceArtist":
		if influence.SourceArtist != nil {
			res = PointerFieldInitStatement
			res = strings.ReplaceAll(res, "{{Identifier}}", influence.GongGetIdentifier(stage))
			res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "SourceArtist")
			res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", influence.SourceArtist.GongGetIdentifier(stage))
		} else {
			// in case of nil pointer, we need to unstage the previous value
			res = PointerFieldInitStatement
			res = strings.ReplaceAll(res, "{{Identifier}}", influence.GongGetIdentifier(stage))
			res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "SourceArtist")
			res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", "nil")
		}
	case "TargetMovement":
		if influence.TargetMovement != nil {
			res = PointerFieldInitStatement
			res = strings.ReplaceAll(res, "{{Identifier}}", influence.GongGetIdentifier(stage))
			res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "TargetMovement")
			res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", influence.TargetMovement.GongGetIdentifier(stage))
		} else {
			// in case of nil pointer, we need to unstage the previous value
			res = PointerFieldInitStatement
			res = strings.ReplaceAll(res, "{{Identifier}}", influence.GongGetIdentifier(stage))
			res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "TargetMovement")
			res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", "nil")
		}
	case "TargetArtefactType":
		if influence.TargetArtefactType != nil {
			res = PointerFieldInitStatement
			res = strings.ReplaceAll(res, "{{Identifier}}", influence.GongGetIdentifier(stage))
			res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "TargetArtefactType")
			res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", influence.TargetArtefactType.GongGetIdentifier(stage))
		} else {
			// in case of nil pointer, we need to unstage the previous value
			res = PointerFieldInitStatement
			res = strings.ReplaceAll(res, "{{Identifier}}", influence.GongGetIdentifier(stage))
			res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "TargetArtefactType")
			res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", "nil")
		}
	case "TargetArtist":
		if influence.TargetArtist != nil {
			res = PointerFieldInitStatement
			res = strings.ReplaceAll(res, "{{Identifier}}", influence.GongGetIdentifier(stage))
			res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "TargetArtist")
			res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", influence.TargetArtist.GongGetIdentifier(stage))
		} else {
			// in case of nil pointer, we need to unstage the previous value
			res = PointerFieldInitStatement
			res = strings.ReplaceAll(res, "{{Identifier}}", influence.GongGetIdentifier(stage))
			res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "TargetArtist")
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
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", ToRawStringLiteral(influenceshape.Name))

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
		var sb strings.Builder
		for _, _controlpointshape := range influenceshape.ControlPointShapes {
			tmp := SliceOfPointersFieldInitStatement
			tmp = strings.ReplaceAll(tmp, "{{Identifier}}", influenceshape.GongGetIdentifier(stage))
			tmp = strings.ReplaceAll(tmp, "{{GeneratedFieldName}}", "ControlPointShapes")
			tmp = strings.ReplaceAll(tmp, "{{GeneratedFieldNameValue}}", _controlpointshape.GongGetIdentifier(stage))
			sb.WriteString(tmp)
		}
		res = sb.String()
	default:
		log.Panicf("Unknown field %s for Gongstruct InfluenceShape", fieldName)
	}
	return
}

func (movement *Movement) GongMarshallField(stage *Stage, fieldName string) (res string) {

	switch fieldName {
	case "Name":
		res = StringInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", movement.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "Name")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", ToRawStringLiteral(movement.Name))
	case "Date":
		res = TimeInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", movement.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "Date")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", movement.Date.String())
	case "IsAbstract":
		res = NumberInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", movement.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "IsAbstract")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%t", movement.IsAbstract))
	case "IsModern":
		res = NumberInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", movement.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "IsModern")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%t", movement.IsModern))
	case "IsMajor":
		res = NumberInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", movement.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "IsMajor")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%t", movement.IsMajor))
	case "IsMinor":
		res = NumberInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", movement.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "IsMinor")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%t", movement.IsMinor))
	case "AdditionnalName":
		res = StringInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", movement.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "AdditionnalName")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", ToRawStringLiteral(movement.AdditionnalName))
	case "HideDate":
		res = NumberInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", movement.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "HideDate")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%t", movement.HideDate))

	case "Places":
		var sb strings.Builder
		for _, _place := range movement.Places {
			tmp := SliceOfPointersFieldInitStatement
			tmp = strings.ReplaceAll(tmp, "{{Identifier}}", movement.GongGetIdentifier(stage))
			tmp = strings.ReplaceAll(tmp, "{{GeneratedFieldName}}", "Places")
			tmp = strings.ReplaceAll(tmp, "{{GeneratedFieldNameValue}}", _place.GongGetIdentifier(stage))
			sb.WriteString(tmp)
		}
		res = sb.String()
	default:
		log.Panicf("Unknown field %s for Gongstruct Movement", fieldName)
	}
	return
}

func (movementshape *MovementShape) GongMarshallField(stage *Stage, fieldName string) (res string) {

	switch fieldName {
	case "Name":
		res = StringInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", movementshape.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "Name")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", ToRawStringLiteral(movementshape.Name))
	case "X":
		res = NumberInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", movementshape.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "X")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%f", movementshape.X))
	case "Y":
		res = NumberInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", movementshape.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "Y")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%f", movementshape.Y))
	case "Width":
		res = NumberInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", movementshape.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "Width")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%f", movementshape.Width))
	case "Height":
		res = NumberInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", movementshape.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "Height")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%f", movementshape.Height))

	case "Movement":
		if movementshape.Movement != nil {
			res = PointerFieldInitStatement
			res = strings.ReplaceAll(res, "{{Identifier}}", movementshape.GongGetIdentifier(stage))
			res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "Movement")
			res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", movementshape.Movement.GongGetIdentifier(stage))
		} else {
			// in case of nil pointer, we need to unstage the previous value
			res = PointerFieldInitStatement
			res = strings.ReplaceAll(res, "{{Identifier}}", movementshape.GongGetIdentifier(stage))
			res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "Movement")
			res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", "nil")
		}
	default:
		log.Panicf("Unknown field %s for Gongstruct MovementShape", fieldName)
	}
	return
}

func (place *Place) GongMarshallField(stage *Stage, fieldName string) (res string) {

	switch fieldName {
	case "Name":
		res = StringInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", place.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "Name")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", ToRawStringLiteral(place.Name))

	default:
		log.Panicf("Unknown field %s for Gongstruct Place", fieldName)
	}
	return
}

// insertion point for marshall all fields methods
func (artefacttype *ArtefactType) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(artefacttype.GongMarshallField(stage, "Name"))
	}
	initRes = initializerStatements.String()
	ptrRes = pointersInitializesStatements.String()
	return
}
func (artefacttypeshape *ArtefactTypeShape) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(artefacttypeshape.GongMarshallField(stage, "Name"))
		pointersInitializesStatements.WriteString(artefacttypeshape.GongMarshallField(stage, "ArtefactType"))
		initializerStatements.WriteString(artefacttypeshape.GongMarshallField(stage, "X"))
		initializerStatements.WriteString(artefacttypeshape.GongMarshallField(stage, "Y"))
		initializerStatements.WriteString(artefacttypeshape.GongMarshallField(stage, "Width"))
		initializerStatements.WriteString(artefacttypeshape.GongMarshallField(stage, "Height"))
	}
	initRes = initializerStatements.String()
	ptrRes = pointersInitializesStatements.String()
	return
}
func (artist *Artist) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(artist.GongMarshallField(stage, "Name"))
		initializerStatements.WriteString(artist.GongMarshallField(stage, "IsDead"))
		initializerStatements.WriteString(artist.GongMarshallField(stage, "DateOfDeath"))
		pointersInitializesStatements.WriteString(artist.GongMarshallField(stage, "Place"))
	}
	initRes = initializerStatements.String()
	ptrRes = pointersInitializesStatements.String()
	return
}
func (artistshape *ArtistShape) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(artistshape.GongMarshallField(stage, "Name"))
		pointersInitializesStatements.WriteString(artistshape.GongMarshallField(stage, "Artist"))
		initializerStatements.WriteString(artistshape.GongMarshallField(stage, "X"))
		initializerStatements.WriteString(artistshape.GongMarshallField(stage, "Y"))
		initializerStatements.WriteString(artistshape.GongMarshallField(stage, "Width"))
		initializerStatements.WriteString(artistshape.GongMarshallField(stage, "Height"))
	}
	initRes = initializerStatements.String()
	ptrRes = pointersInitializesStatements.String()
	return
}
func (controlpointshape *ControlPointShape) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(controlpointshape.GongMarshallField(stage, "Name"))
		initializerStatements.WriteString(controlpointshape.GongMarshallField(stage, "X_Relative"))
		initializerStatements.WriteString(controlpointshape.GongMarshallField(stage, "Y_Relative"))
		initializerStatements.WriteString(controlpointshape.GongMarshallField(stage, "IsStartShapeTheClosestShape"))
	}
	initRes = initializerStatements.String()
	ptrRes = pointersInitializesStatements.String()
	return
}
func (desk *Desk) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(desk.GongMarshallField(stage, "Name"))
		pointersInitializesStatements.WriteString(desk.GongMarshallField(stage, "SelectedDiagram"))
	}
	initRes = initializerStatements.String()
	ptrRes = pointersInitializesStatements.String()
	return
}
func (diagram *Diagram) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(diagram.GongMarshallField(stage, "Name"))
		pointersInitializesStatements.WriteString(diagram.GongMarshallField(stage, "MovementShapes"))
		pointersInitializesStatements.WriteString(diagram.GongMarshallField(stage, "ArtefactTypeShapes"))
		pointersInitializesStatements.WriteString(diagram.GongMarshallField(stage, "ArtistShapes"))
		pointersInitializesStatements.WriteString(diagram.GongMarshallField(stage, "InfluenceShapes"))
		initializerStatements.WriteString(diagram.GongMarshallField(stage, "IsEditable"))
		initializerStatements.WriteString(diagram.GongMarshallField(stage, "IsNodeExpanded"))
		initializerStatements.WriteString(diagram.GongMarshallField(stage, "IsMovementCategoryNodeExpanded"))
		initializerStatements.WriteString(diagram.GongMarshallField(stage, "IsArtefactTypeCategoryNodeExpanded"))
		initializerStatements.WriteString(diagram.GongMarshallField(stage, "IsArtistCategoryNodeExpanded"))
		initializerStatements.WriteString(diagram.GongMarshallField(stage, "IsInfluenceCategoryNodeExpanded"))
		initializerStatements.WriteString(diagram.GongMarshallField(stage, "IsMovementCategoryShown"))
		initializerStatements.WriteString(diagram.GongMarshallField(stage, "IsArtefactTypeCategoryShown"))
		initializerStatements.WriteString(diagram.GongMarshallField(stage, "IsArtistCategoryShown"))
		initializerStatements.WriteString(diagram.GongMarshallField(stage, "IsInfluenceCategoryShown"))
		initializerStatements.WriteString(diagram.GongMarshallField(stage, "StartDate"))
		initializerStatements.WriteString(diagram.GongMarshallField(stage, "EndDate"))
		initializerStatements.WriteString(diagram.GongMarshallField(stage, "XMargin"))
		initializerStatements.WriteString(diagram.GongMarshallField(stage, "YMargin"))
		initializerStatements.WriteString(diagram.GongMarshallField(stage, "Height"))
		initializerStatements.WriteString(diagram.GongMarshallField(stage, "NextVerticalDateXMargin"))
		initializerStatements.WriteString(diagram.GongMarshallField(stage, "RedColorCode"))
		initializerStatements.WriteString(diagram.GongMarshallField(stage, "BackgroundGreyColorCode"))
		initializerStatements.WriteString(diagram.GongMarshallField(stage, "GrayColorCode"))
		initializerStatements.WriteString(diagram.GongMarshallField(stage, "BottomBoxYOffset"))
		initializerStatements.WriteString(diagram.GongMarshallField(stage, "BottomBoxWidth"))
		initializerStatements.WriteString(diagram.GongMarshallField(stage, "BottomBoxHeigth"))
		initializerStatements.WriteString(diagram.GongMarshallField(stage, "BottomBoxFontSize"))
		initializerStatements.WriteString(diagram.GongMarshallField(stage, "BottomBoxFontWeigth"))
		initializerStatements.WriteString(diagram.GongMarshallField(stage, "BottomBoxFontFamily"))
		initializerStatements.WriteString(diagram.GongMarshallField(stage, "BottomBoxLetterSpacing"))
		initializerStatements.WriteString(diagram.GongMarshallField(stage, "BottomBoxLetterColorCode"))
		initializerStatements.WriteString(diagram.GongMarshallField(stage, "MovementRectAnchorType"))
		initializerStatements.WriteString(diagram.GongMarshallField(stage, "MovementTextAnchorType"))
		initializerStatements.WriteString(diagram.GongMarshallField(stage, "MovementDominantBaselineType"))
		initializerStatements.WriteString(diagram.GongMarshallField(stage, "MovementFontSize"))
		initializerStatements.WriteString(diagram.GongMarshallField(stage, "MajorMovementFontSize"))
		initializerStatements.WriteString(diagram.GongMarshallField(stage, "MinorMovementFontSize"))
		initializerStatements.WriteString(diagram.GongMarshallField(stage, "MovementFontWeigth"))
		initializerStatements.WriteString(diagram.GongMarshallField(stage, "MovementFontFamily"))
		initializerStatements.WriteString(diagram.GongMarshallField(stage, "MovementLetterSpacing"))
		initializerStatements.WriteString(diagram.GongMarshallField(stage, "AbstractMovementFontSize"))
		initializerStatements.WriteString(diagram.GongMarshallField(stage, "AbstractMovementRectAnchorType"))
		initializerStatements.WriteString(diagram.GongMarshallField(stage, "AbstractMovementTextAnchorType"))
		initializerStatements.WriteString(diagram.GongMarshallField(stage, "AbstractDominantBaselineType"))
		initializerStatements.WriteString(diagram.GongMarshallField(stage, "MovementDateRectAnchorType"))
		initializerStatements.WriteString(diagram.GongMarshallField(stage, "MovementDateTextAnchorType"))
		initializerStatements.WriteString(diagram.GongMarshallField(stage, "MovementDateTextDominantBaselineType"))
		initializerStatements.WriteString(diagram.GongMarshallField(stage, "MovementDateAndPlacesFontSize"))
		initializerStatements.WriteString(diagram.GongMarshallField(stage, "MovementDateAndPlacesFontWeigth"))
		initializerStatements.WriteString(diagram.GongMarshallField(stage, "MovementDateAndPlacesFontFamily"))
		initializerStatements.WriteString(diagram.GongMarshallField(stage, "MovementDateAndPlacesLetterSpacing"))
		initializerStatements.WriteString(diagram.GongMarshallField(stage, "MovementBelowArcY_Offset"))
		initializerStatements.WriteString(diagram.GongMarshallField(stage, "MovementBelowArcY_OffsetPerPlace"))
		initializerStatements.WriteString(diagram.GongMarshallField(stage, "MovementPlacesRectAnchorType"))
		initializerStatements.WriteString(diagram.GongMarshallField(stage, "MovementPlacesTextAnchorType"))
		initializerStatements.WriteString(diagram.GongMarshallField(stage, "MovementPlacesDominantBaselineType"))
		initializerStatements.WriteString(diagram.GongMarshallField(stage, "ArtefactTypeFontSize"))
		initializerStatements.WriteString(diagram.GongMarshallField(stage, "ArtefactTypeFontWeigth"))
		initializerStatements.WriteString(diagram.GongMarshallField(stage, "ArtefactTypeFontFamily"))
		initializerStatements.WriteString(diagram.GongMarshallField(stage, "ArtefactTypeLetterSpacing"))
		initializerStatements.WriteString(diagram.GongMarshallField(stage, "ArtefactTypeRectAnchorType"))
		initializerStatements.WriteString(diagram.GongMarshallField(stage, "ArtefactDominantBaselineType"))
		initializerStatements.WriteString(diagram.GongMarshallField(stage, "ArtefactTypeStrokeWidth"))
		initializerStatements.WriteString(diagram.GongMarshallField(stage, "ArtistRectAnchorType"))
		initializerStatements.WriteString(diagram.GongMarshallField(stage, "ArtistTextAnchorType"))
		initializerStatements.WriteString(diagram.GongMarshallField(stage, "ArtistDominantBaselineType"))
		initializerStatements.WriteString(diagram.GongMarshallField(stage, "ArtistFontSize"))
		initializerStatements.WriteString(diagram.GongMarshallField(stage, "MajorArtistFontSize"))
		initializerStatements.WriteString(diagram.GongMarshallField(stage, "MinorArtistFontSize"))
		initializerStatements.WriteString(diagram.GongMarshallField(stage, "ArtistFontWeigth"))
		initializerStatements.WriteString(diagram.GongMarshallField(stage, "ArtistFontFamily"))
		initializerStatements.WriteString(diagram.GongMarshallField(stage, "ArtistLetterSpacing"))
		initializerStatements.WriteString(diagram.GongMarshallField(stage, "ArtistDateRectAnchorType"))
		initializerStatements.WriteString(diagram.GongMarshallField(stage, "ArtistDateTextAnchorType"))
		initializerStatements.WriteString(diagram.GongMarshallField(stage, "ArtistDateDominantBaselineType"))
		initializerStatements.WriteString(diagram.GongMarshallField(stage, "ArtistDateAndPlacesFontSize"))
		initializerStatements.WriteString(diagram.GongMarshallField(stage, "ArtistDateAndPlacesFontWeigth"))
		initializerStatements.WriteString(diagram.GongMarshallField(stage, "ArtistDateAndPlacesFontFamily"))
		initializerStatements.WriteString(diagram.GongMarshallField(stage, "ArtistDateAndPlacesLetterSpacing"))
		initializerStatements.WriteString(diagram.GongMarshallField(stage, "ArtistPlacesRectAnchorType"))
		initializerStatements.WriteString(diagram.GongMarshallField(stage, "ArtistPlacesTextAnchorType"))
		initializerStatements.WriteString(diagram.GongMarshallField(stage, "ArtistPlacesDominantBaselineType"))
		initializerStatements.WriteString(diagram.GongMarshallField(stage, "InfluenceArrowSize"))
		initializerStatements.WriteString(diagram.GongMarshallField(stage, "InfluenceArrowStartOffset"))
		initializerStatements.WriteString(diagram.GongMarshallField(stage, "InfluenceArrowEndOffset"))
		initializerStatements.WriteString(diagram.GongMarshallField(stage, "InfluenceCornerRadius"))
		initializerStatements.WriteString(diagram.GongMarshallField(stage, "InfluenceDashedLinePattern"))
	}
	initRes = initializerStatements.String()
	ptrRes = pointersInitializesStatements.String()
	return
}
func (influence *Influence) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(influence.GongMarshallField(stage, "Name"))
		pointersInitializesStatements.WriteString(influence.GongMarshallField(stage, "SourceMovement"))
		pointersInitializesStatements.WriteString(influence.GongMarshallField(stage, "SourceArtefactType"))
		pointersInitializesStatements.WriteString(influence.GongMarshallField(stage, "SourceArtist"))
		pointersInitializesStatements.WriteString(influence.GongMarshallField(stage, "TargetMovement"))
		pointersInitializesStatements.WriteString(influence.GongMarshallField(stage, "TargetArtefactType"))
		pointersInitializesStatements.WriteString(influence.GongMarshallField(stage, "TargetArtist"))
		initializerStatements.WriteString(influence.GongMarshallField(stage, "IsHypothtical"))
	}
	initRes = initializerStatements.String()
	ptrRes = pointersInitializesStatements.String()
	return
}
func (influenceshape *InfluenceShape) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(influenceshape.GongMarshallField(stage, "Name"))
		pointersInitializesStatements.WriteString(influenceshape.GongMarshallField(stage, "Influence"))
		pointersInitializesStatements.WriteString(influenceshape.GongMarshallField(stage, "ControlPointShapes"))
	}
	initRes = initializerStatements.String()
	ptrRes = pointersInitializesStatements.String()
	return
}
func (movement *Movement) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(movement.GongMarshallField(stage, "Name"))
		initializerStatements.WriteString(movement.GongMarshallField(stage, "Date"))
		pointersInitializesStatements.WriteString(movement.GongMarshallField(stage, "Places"))
		initializerStatements.WriteString(movement.GongMarshallField(stage, "IsAbstract"))
		initializerStatements.WriteString(movement.GongMarshallField(stage, "IsModern"))
		initializerStatements.WriteString(movement.GongMarshallField(stage, "IsMajor"))
		initializerStatements.WriteString(movement.GongMarshallField(stage, "IsMinor"))
		initializerStatements.WriteString(movement.GongMarshallField(stage, "AdditionnalName"))
		initializerStatements.WriteString(movement.GongMarshallField(stage, "HideDate"))
	}
	initRes = initializerStatements.String()
	ptrRes = pointersInitializesStatements.String()
	return
}
func (movementshape *MovementShape) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(movementshape.GongMarshallField(stage, "Name"))
		pointersInitializesStatements.WriteString(movementshape.GongMarshallField(stage, "Movement"))
		initializerStatements.WriteString(movementshape.GongMarshallField(stage, "X"))
		initializerStatements.WriteString(movementshape.GongMarshallField(stage, "Y"))
		initializerStatements.WriteString(movementshape.GongMarshallField(stage, "Width"))
		initializerStatements.WriteString(movementshape.GongMarshallField(stage, "Height"))
	}
	initRes = initializerStatements.String()
	ptrRes = pointersInitializesStatements.String()
	return
}
func (place *Place) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(place.GongMarshallField(stage, "Name"))
	}
	initRes = initializerStatements.String()
	ptrRes = pointersInitializesStatements.String()
	return
}
