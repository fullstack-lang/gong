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
}
`

const IdentifiersDecls = `
	{{Identifier}} := (&models.{{GeneratedStructName}}{Name: ` + "`" + `{{GeneratedFieldNameValue}}` + "`" + `}).Stage(stage)`

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
	animateOrdered := []*Animate{}
	for animate := range stage.Animates {
		animateOrdered = append(animateOrdered, animate)
	}
	sort.Slice(animateOrdered[:], func(i, j int) bool {
		animatei := animateOrdered[i]
		animatej := animateOrdered[j]
		animatei_order, oki := stage.AnimateMap_Staged_Order[animatei]
		animatej_order, okj := stage.AnimateMap_Staged_Order[animatej]
		if !oki || !okj {
			log.Fatalln("unknown pointers")
		}
		return animatei_order < animatej_order
	})
	if len(animateOrdered) > 0 {
		identifiersDecl += "\n"
	}
	for _, animate := range animateOrdered {

		identifiersDecl += animate.GongMarshallIdentifier(stage)

		initializerStatements += "\n"
		// Insertion point for basic fields value assignment
		initializerStatements += animate.GongMarshallField(stage, "Name")
		initializerStatements += animate.GongMarshallField(stage, "AttributeName")
		initializerStatements += animate.GongMarshallField(stage, "Values")
		initializerStatements += animate.GongMarshallField(stage, "From")
		initializerStatements += animate.GongMarshallField(stage, "To")
		initializerStatements += animate.GongMarshallField(stage, "Dur")
		initializerStatements += animate.GongMarshallField(stage, "RepeatCount")
	}

	circleOrdered := []*Circle{}
	for circle := range stage.Circles {
		circleOrdered = append(circleOrdered, circle)
	}
	sort.Slice(circleOrdered[:], func(i, j int) bool {
		circlei := circleOrdered[i]
		circlej := circleOrdered[j]
		circlei_order, oki := stage.CircleMap_Staged_Order[circlei]
		circlej_order, okj := stage.CircleMap_Staged_Order[circlej]
		if !oki || !okj {
			log.Fatalln("unknown pointers")
		}
		return circlei_order < circlej_order
	})
	if len(circleOrdered) > 0 {
		identifiersDecl += "\n"
	}
	for _, circle := range circleOrdered {

		identifiersDecl += circle.GongMarshallIdentifier(stage)

		initializerStatements += "\n"
		// Insertion point for basic fields value assignment
		initializerStatements += circle.GongMarshallField(stage, "Name")
		initializerStatements += circle.GongMarshallField(stage, "CX")
		initializerStatements += circle.GongMarshallField(stage, "CY")
		initializerStatements += circle.GongMarshallField(stage, "Radius")
		initializerStatements += circle.GongMarshallField(stage, "Color")
		initializerStatements += circle.GongMarshallField(stage, "FillOpacity")
		initializerStatements += circle.GongMarshallField(stage, "Stroke")
		initializerStatements += circle.GongMarshallField(stage, "StrokeOpacity")
		initializerStatements += circle.GongMarshallField(stage, "StrokeWidth")
		initializerStatements += circle.GongMarshallField(stage, "StrokeDashArray")
		initializerStatements += circle.GongMarshallField(stage, "StrokeDashArrayWhenSelected")
		initializerStatements += circle.GongMarshallField(stage, "Transform")
		pointersInitializesStatements += circle.GongMarshallField(stage, "Animations")
	}

	conditionOrdered := []*Condition{}
	for condition := range stage.Conditions {
		conditionOrdered = append(conditionOrdered, condition)
	}
	sort.Slice(conditionOrdered[:], func(i, j int) bool {
		conditioni := conditionOrdered[i]
		conditionj := conditionOrdered[j]
		conditioni_order, oki := stage.ConditionMap_Staged_Order[conditioni]
		conditionj_order, okj := stage.ConditionMap_Staged_Order[conditionj]
		if !oki || !okj {
			log.Fatalln("unknown pointers")
		}
		return conditioni_order < conditionj_order
	})
	if len(conditionOrdered) > 0 {
		identifiersDecl += "\n"
	}
	for _, condition := range conditionOrdered {

		identifiersDecl += condition.GongMarshallIdentifier(stage)

		initializerStatements += "\n"
		// Insertion point for basic fields value assignment
		initializerStatements += condition.GongMarshallField(stage, "Name")
	}

	controlpointOrdered := []*ControlPoint{}
	for controlpoint := range stage.ControlPoints {
		controlpointOrdered = append(controlpointOrdered, controlpoint)
	}
	sort.Slice(controlpointOrdered[:], func(i, j int) bool {
		controlpointi := controlpointOrdered[i]
		controlpointj := controlpointOrdered[j]
		controlpointi_order, oki := stage.ControlPointMap_Staged_Order[controlpointi]
		controlpointj_order, okj := stage.ControlPointMap_Staged_Order[controlpointj]
		if !oki || !okj {
			log.Fatalln("unknown pointers")
		}
		return controlpointi_order < controlpointj_order
	})
	if len(controlpointOrdered) > 0 {
		identifiersDecl += "\n"
	}
	for _, controlpoint := range controlpointOrdered {

		identifiersDecl += controlpoint.GongMarshallIdentifier(stage)

		initializerStatements += "\n"
		// Insertion point for basic fields value assignment
		initializerStatements += controlpoint.GongMarshallField(stage, "Name")
		initializerStatements += controlpoint.GongMarshallField(stage, "X_Relative")
		initializerStatements += controlpoint.GongMarshallField(stage, "Y_Relative")
		pointersInitializesStatements += controlpoint.GongMarshallField(stage, "ClosestRect")
	}

	ellipseOrdered := []*Ellipse{}
	for ellipse := range stage.Ellipses {
		ellipseOrdered = append(ellipseOrdered, ellipse)
	}
	sort.Slice(ellipseOrdered[:], func(i, j int) bool {
		ellipsei := ellipseOrdered[i]
		ellipsej := ellipseOrdered[j]
		ellipsei_order, oki := stage.EllipseMap_Staged_Order[ellipsei]
		ellipsej_order, okj := stage.EllipseMap_Staged_Order[ellipsej]
		if !oki || !okj {
			log.Fatalln("unknown pointers")
		}
		return ellipsei_order < ellipsej_order
	})
	if len(ellipseOrdered) > 0 {
		identifiersDecl += "\n"
	}
	for _, ellipse := range ellipseOrdered {

		identifiersDecl += ellipse.GongMarshallIdentifier(stage)

		initializerStatements += "\n"
		// Insertion point for basic fields value assignment
		initializerStatements += ellipse.GongMarshallField(stage, "Name")
		initializerStatements += ellipse.GongMarshallField(stage, "CX")
		initializerStatements += ellipse.GongMarshallField(stage, "CY")
		initializerStatements += ellipse.GongMarshallField(stage, "RX")
		initializerStatements += ellipse.GongMarshallField(stage, "RY")
		initializerStatements += ellipse.GongMarshallField(stage, "Color")
		initializerStatements += ellipse.GongMarshallField(stage, "FillOpacity")
		initializerStatements += ellipse.GongMarshallField(stage, "Stroke")
		initializerStatements += ellipse.GongMarshallField(stage, "StrokeOpacity")
		initializerStatements += ellipse.GongMarshallField(stage, "StrokeWidth")
		initializerStatements += ellipse.GongMarshallField(stage, "StrokeDashArray")
		initializerStatements += ellipse.GongMarshallField(stage, "StrokeDashArrayWhenSelected")
		initializerStatements += ellipse.GongMarshallField(stage, "Transform")
		pointersInitializesStatements += ellipse.GongMarshallField(stage, "Animates")
	}

	layerOrdered := []*Layer{}
	for layer := range stage.Layers {
		layerOrdered = append(layerOrdered, layer)
	}
	sort.Slice(layerOrdered[:], func(i, j int) bool {
		layeri := layerOrdered[i]
		layerj := layerOrdered[j]
		layeri_order, oki := stage.LayerMap_Staged_Order[layeri]
		layerj_order, okj := stage.LayerMap_Staged_Order[layerj]
		if !oki || !okj {
			log.Fatalln("unknown pointers")
		}
		return layeri_order < layerj_order
	})
	if len(layerOrdered) > 0 {
		identifiersDecl += "\n"
	}
	for _, layer := range layerOrdered {

		identifiersDecl += layer.GongMarshallIdentifier(stage)

		initializerStatements += "\n"
		// Insertion point for basic fields value assignment
		initializerStatements += layer.GongMarshallField(stage, "Name")
		pointersInitializesStatements += layer.GongMarshallField(stage, "Rects")
		pointersInitializesStatements += layer.GongMarshallField(stage, "Texts")
		pointersInitializesStatements += layer.GongMarshallField(stage, "Circles")
		pointersInitializesStatements += layer.GongMarshallField(stage, "Lines")
		pointersInitializesStatements += layer.GongMarshallField(stage, "Ellipses")
		pointersInitializesStatements += layer.GongMarshallField(stage, "Polylines")
		pointersInitializesStatements += layer.GongMarshallField(stage, "Polygones")
		pointersInitializesStatements += layer.GongMarshallField(stage, "Paths")
		pointersInitializesStatements += layer.GongMarshallField(stage, "Links")
		pointersInitializesStatements += layer.GongMarshallField(stage, "RectLinkLinks")
	}

	lineOrdered := []*Line{}
	for line := range stage.Lines {
		lineOrdered = append(lineOrdered, line)
	}
	sort.Slice(lineOrdered[:], func(i, j int) bool {
		linei := lineOrdered[i]
		linej := lineOrdered[j]
		linei_order, oki := stage.LineMap_Staged_Order[linei]
		linej_order, okj := stage.LineMap_Staged_Order[linej]
		if !oki || !okj {
			log.Fatalln("unknown pointers")
		}
		return linei_order < linej_order
	})
	if len(lineOrdered) > 0 {
		identifiersDecl += "\n"
	}
	for _, line := range lineOrdered {

		identifiersDecl += line.GongMarshallIdentifier(stage)

		initializerStatements += "\n"
		// Insertion point for basic fields value assignment
		initializerStatements += line.GongMarshallField(stage, "Name")
		initializerStatements += line.GongMarshallField(stage, "X1")
		initializerStatements += line.GongMarshallField(stage, "Y1")
		initializerStatements += line.GongMarshallField(stage, "X2")
		initializerStatements += line.GongMarshallField(stage, "Y2")
		initializerStatements += line.GongMarshallField(stage, "Color")
		initializerStatements += line.GongMarshallField(stage, "FillOpacity")
		initializerStatements += line.GongMarshallField(stage, "Stroke")
		initializerStatements += line.GongMarshallField(stage, "StrokeOpacity")
		initializerStatements += line.GongMarshallField(stage, "StrokeWidth")
		initializerStatements += line.GongMarshallField(stage, "StrokeDashArray")
		initializerStatements += line.GongMarshallField(stage, "StrokeDashArrayWhenSelected")
		initializerStatements += line.GongMarshallField(stage, "Transform")
		pointersInitializesStatements += line.GongMarshallField(stage, "Animates")
		initializerStatements += line.GongMarshallField(stage, "MouseClickX")
		initializerStatements += line.GongMarshallField(stage, "MouseClickY")
	}

	linkOrdered := []*Link{}
	for link := range stage.Links {
		linkOrdered = append(linkOrdered, link)
	}
	sort.Slice(linkOrdered[:], func(i, j int) bool {
		linki := linkOrdered[i]
		linkj := linkOrdered[j]
		linki_order, oki := stage.LinkMap_Staged_Order[linki]
		linkj_order, okj := stage.LinkMap_Staged_Order[linkj]
		if !oki || !okj {
			log.Fatalln("unknown pointers")
		}
		return linki_order < linkj_order
	})
	if len(linkOrdered) > 0 {
		identifiersDecl += "\n"
	}
	for _, link := range linkOrdered {

		identifiersDecl += link.GongMarshallIdentifier(stage)

		initializerStatements += "\n"
		// Insertion point for basic fields value assignment
		initializerStatements += link.GongMarshallField(stage, "Name")
		initializerStatements += link.GongMarshallField(stage, "Type")
		initializerStatements += link.GongMarshallField(stage, "IsBezierCurve")
		pointersInitializesStatements += link.GongMarshallField(stage, "Start")
		initializerStatements += link.GongMarshallField(stage, "StartAnchorType")
		pointersInitializesStatements += link.GongMarshallField(stage, "End")
		initializerStatements += link.GongMarshallField(stage, "EndAnchorType")
		initializerStatements += link.GongMarshallField(stage, "StartOrientation")
		initializerStatements += link.GongMarshallField(stage, "StartRatio")
		initializerStatements += link.GongMarshallField(stage, "EndOrientation")
		initializerStatements += link.GongMarshallField(stage, "EndRatio")
		initializerStatements += link.GongMarshallField(stage, "CornerOffsetRatio")
		initializerStatements += link.GongMarshallField(stage, "CornerRadius")
		initializerStatements += link.GongMarshallField(stage, "HasEndArrow")
		initializerStatements += link.GongMarshallField(stage, "EndArrowSize")
		initializerStatements += link.GongMarshallField(stage, "EndArrowOffset")
		initializerStatements += link.GongMarshallField(stage, "HasStartArrow")
		initializerStatements += link.GongMarshallField(stage, "StartArrowSize")
		initializerStatements += link.GongMarshallField(stage, "StartArrowOffset")
		pointersInitializesStatements += link.GongMarshallField(stage, "TextAtArrowStart")
		pointersInitializesStatements += link.GongMarshallField(stage, "TextAtArrowEnd")
		pointersInitializesStatements += link.GongMarshallField(stage, "ControlPoints")
		initializerStatements += link.GongMarshallField(stage, "Color")
		initializerStatements += link.GongMarshallField(stage, "FillOpacity")
		initializerStatements += link.GongMarshallField(stage, "Stroke")
		initializerStatements += link.GongMarshallField(stage, "StrokeOpacity")
		initializerStatements += link.GongMarshallField(stage, "StrokeWidth")
		initializerStatements += link.GongMarshallField(stage, "StrokeDashArray")
		initializerStatements += link.GongMarshallField(stage, "StrokeDashArrayWhenSelected")
		initializerStatements += link.GongMarshallField(stage, "Transform")
		initializerStatements += link.GongMarshallField(stage, "MouseX")
		initializerStatements += link.GongMarshallField(stage, "MouseY")
		initializerStatements += link.GongMarshallField(stage, "MouseEventKey")
	}

	linkanchoredtextOrdered := []*LinkAnchoredText{}
	for linkanchoredtext := range stage.LinkAnchoredTexts {
		linkanchoredtextOrdered = append(linkanchoredtextOrdered, linkanchoredtext)
	}
	sort.Slice(linkanchoredtextOrdered[:], func(i, j int) bool {
		linkanchoredtexti := linkanchoredtextOrdered[i]
		linkanchoredtextj := linkanchoredtextOrdered[j]
		linkanchoredtexti_order, oki := stage.LinkAnchoredTextMap_Staged_Order[linkanchoredtexti]
		linkanchoredtextj_order, okj := stage.LinkAnchoredTextMap_Staged_Order[linkanchoredtextj]
		if !oki || !okj {
			log.Fatalln("unknown pointers")
		}
		return linkanchoredtexti_order < linkanchoredtextj_order
	})
	if len(linkanchoredtextOrdered) > 0 {
		identifiersDecl += "\n"
	}
	for _, linkanchoredtext := range linkanchoredtextOrdered {

		identifiersDecl += linkanchoredtext.GongMarshallIdentifier(stage)

		initializerStatements += "\n"
		// Insertion point for basic fields value assignment
		initializerStatements += linkanchoredtext.GongMarshallField(stage, "Name")
		initializerStatements += linkanchoredtext.GongMarshallField(stage, "Content")
		initializerStatements += linkanchoredtext.GongMarshallField(stage, "AutomaticLayout")
		initializerStatements += linkanchoredtext.GongMarshallField(stage, "LinkAnchorType")
		initializerStatements += linkanchoredtext.GongMarshallField(stage, "X_Offset")
		initializerStatements += linkanchoredtext.GongMarshallField(stage, "Y_Offset")
		initializerStatements += linkanchoredtext.GongMarshallField(stage, "FontWeight")
		initializerStatements += linkanchoredtext.GongMarshallField(stage, "FontSize")
		initializerStatements += linkanchoredtext.GongMarshallField(stage, "FontStyle")
		initializerStatements += linkanchoredtext.GongMarshallField(stage, "LetterSpacing")
		initializerStatements += linkanchoredtext.GongMarshallField(stage, "FontFamily")
		initializerStatements += linkanchoredtext.GongMarshallField(stage, "WhiteSpace")
		initializerStatements += linkanchoredtext.GongMarshallField(stage, "Color")
		initializerStatements += linkanchoredtext.GongMarshallField(stage, "FillOpacity")
		initializerStatements += linkanchoredtext.GongMarshallField(stage, "Stroke")
		initializerStatements += linkanchoredtext.GongMarshallField(stage, "StrokeOpacity")
		initializerStatements += linkanchoredtext.GongMarshallField(stage, "StrokeWidth")
		initializerStatements += linkanchoredtext.GongMarshallField(stage, "StrokeDashArray")
		initializerStatements += linkanchoredtext.GongMarshallField(stage, "StrokeDashArrayWhenSelected")
		initializerStatements += linkanchoredtext.GongMarshallField(stage, "Transform")
		pointersInitializesStatements += linkanchoredtext.GongMarshallField(stage, "Animates")
	}

	pathOrdered := []*Path{}
	for path := range stage.Paths {
		pathOrdered = append(pathOrdered, path)
	}
	sort.Slice(pathOrdered[:], func(i, j int) bool {
		pathi := pathOrdered[i]
		pathj := pathOrdered[j]
		pathi_order, oki := stage.PathMap_Staged_Order[pathi]
		pathj_order, okj := stage.PathMap_Staged_Order[pathj]
		if !oki || !okj {
			log.Fatalln("unknown pointers")
		}
		return pathi_order < pathj_order
	})
	if len(pathOrdered) > 0 {
		identifiersDecl += "\n"
	}
	for _, path := range pathOrdered {

		identifiersDecl += path.GongMarshallIdentifier(stage)

		initializerStatements += "\n"
		// Insertion point for basic fields value assignment
		initializerStatements += path.GongMarshallField(stage, "Name")
		initializerStatements += path.GongMarshallField(stage, "Definition")
		initializerStatements += path.GongMarshallField(stage, "Color")
		initializerStatements += path.GongMarshallField(stage, "FillOpacity")
		initializerStatements += path.GongMarshallField(stage, "Stroke")
		initializerStatements += path.GongMarshallField(stage, "StrokeOpacity")
		initializerStatements += path.GongMarshallField(stage, "StrokeWidth")
		initializerStatements += path.GongMarshallField(stage, "StrokeDashArray")
		initializerStatements += path.GongMarshallField(stage, "StrokeDashArrayWhenSelected")
		initializerStatements += path.GongMarshallField(stage, "Transform")
		pointersInitializesStatements += path.GongMarshallField(stage, "Animates")
	}

	pointOrdered := []*Point{}
	for point := range stage.Points {
		pointOrdered = append(pointOrdered, point)
	}
	sort.Slice(pointOrdered[:], func(i, j int) bool {
		pointi := pointOrdered[i]
		pointj := pointOrdered[j]
		pointi_order, oki := stage.PointMap_Staged_Order[pointi]
		pointj_order, okj := stage.PointMap_Staged_Order[pointj]
		if !oki || !okj {
			log.Fatalln("unknown pointers")
		}
		return pointi_order < pointj_order
	})
	if len(pointOrdered) > 0 {
		identifiersDecl += "\n"
	}
	for _, point := range pointOrdered {

		identifiersDecl += point.GongMarshallIdentifier(stage)

		initializerStatements += "\n"
		// Insertion point for basic fields value assignment
		initializerStatements += point.GongMarshallField(stage, "Name")
		initializerStatements += point.GongMarshallField(stage, "X")
		initializerStatements += point.GongMarshallField(stage, "Y")
	}

	polygoneOrdered := []*Polygone{}
	for polygone := range stage.Polygones {
		polygoneOrdered = append(polygoneOrdered, polygone)
	}
	sort.Slice(polygoneOrdered[:], func(i, j int) bool {
		polygonei := polygoneOrdered[i]
		polygonej := polygoneOrdered[j]
		polygonei_order, oki := stage.PolygoneMap_Staged_Order[polygonei]
		polygonej_order, okj := stage.PolygoneMap_Staged_Order[polygonej]
		if !oki || !okj {
			log.Fatalln("unknown pointers")
		}
		return polygonei_order < polygonej_order
	})
	if len(polygoneOrdered) > 0 {
		identifiersDecl += "\n"
	}
	for _, polygone := range polygoneOrdered {

		identifiersDecl += polygone.GongMarshallIdentifier(stage)

		initializerStatements += "\n"
		// Insertion point for basic fields value assignment
		initializerStatements += polygone.GongMarshallField(stage, "Name")
		initializerStatements += polygone.GongMarshallField(stage, "Points")
		initializerStatements += polygone.GongMarshallField(stage, "Color")
		initializerStatements += polygone.GongMarshallField(stage, "FillOpacity")
		initializerStatements += polygone.GongMarshallField(stage, "Stroke")
		initializerStatements += polygone.GongMarshallField(stage, "StrokeOpacity")
		initializerStatements += polygone.GongMarshallField(stage, "StrokeWidth")
		initializerStatements += polygone.GongMarshallField(stage, "StrokeDashArray")
		initializerStatements += polygone.GongMarshallField(stage, "StrokeDashArrayWhenSelected")
		initializerStatements += polygone.GongMarshallField(stage, "Transform")
		pointersInitializesStatements += polygone.GongMarshallField(stage, "Animates")
	}

	polylineOrdered := []*Polyline{}
	for polyline := range stage.Polylines {
		polylineOrdered = append(polylineOrdered, polyline)
	}
	sort.Slice(polylineOrdered[:], func(i, j int) bool {
		polylinei := polylineOrdered[i]
		polylinej := polylineOrdered[j]
		polylinei_order, oki := stage.PolylineMap_Staged_Order[polylinei]
		polylinej_order, okj := stage.PolylineMap_Staged_Order[polylinej]
		if !oki || !okj {
			log.Fatalln("unknown pointers")
		}
		return polylinei_order < polylinej_order
	})
	if len(polylineOrdered) > 0 {
		identifiersDecl += "\n"
	}
	for _, polyline := range polylineOrdered {

		identifiersDecl += polyline.GongMarshallIdentifier(stage)

		initializerStatements += "\n"
		// Insertion point for basic fields value assignment
		initializerStatements += polyline.GongMarshallField(stage, "Name")
		initializerStatements += polyline.GongMarshallField(stage, "Points")
		initializerStatements += polyline.GongMarshallField(stage, "Color")
		initializerStatements += polyline.GongMarshallField(stage, "FillOpacity")
		initializerStatements += polyline.GongMarshallField(stage, "Stroke")
		initializerStatements += polyline.GongMarshallField(stage, "StrokeOpacity")
		initializerStatements += polyline.GongMarshallField(stage, "StrokeWidth")
		initializerStatements += polyline.GongMarshallField(stage, "StrokeDashArray")
		initializerStatements += polyline.GongMarshallField(stage, "StrokeDashArrayWhenSelected")
		initializerStatements += polyline.GongMarshallField(stage, "Transform")
		pointersInitializesStatements += polyline.GongMarshallField(stage, "Animates")
	}

	rectOrdered := []*Rect{}
	for rect := range stage.Rects {
		rectOrdered = append(rectOrdered, rect)
	}
	sort.Slice(rectOrdered[:], func(i, j int) bool {
		recti := rectOrdered[i]
		rectj := rectOrdered[j]
		recti_order, oki := stage.RectMap_Staged_Order[recti]
		rectj_order, okj := stage.RectMap_Staged_Order[rectj]
		if !oki || !okj {
			log.Fatalln("unknown pointers")
		}
		return recti_order < rectj_order
	})
	if len(rectOrdered) > 0 {
		identifiersDecl += "\n"
	}
	for _, rect := range rectOrdered {

		identifiersDecl += rect.GongMarshallIdentifier(stage)

		initializerStatements += "\n"
		// Insertion point for basic fields value assignment
		initializerStatements += rect.GongMarshallField(stage, "Name")
		initializerStatements += rect.GongMarshallField(stage, "X")
		initializerStatements += rect.GongMarshallField(stage, "Y")
		initializerStatements += rect.GongMarshallField(stage, "Width")
		initializerStatements += rect.GongMarshallField(stage, "Height")
		initializerStatements += rect.GongMarshallField(stage, "RX")
		initializerStatements += rect.GongMarshallField(stage, "Color")
		initializerStatements += rect.GongMarshallField(stage, "FillOpacity")
		initializerStatements += rect.GongMarshallField(stage, "Stroke")
		initializerStatements += rect.GongMarshallField(stage, "StrokeOpacity")
		initializerStatements += rect.GongMarshallField(stage, "StrokeWidth")
		initializerStatements += rect.GongMarshallField(stage, "StrokeDashArray")
		initializerStatements += rect.GongMarshallField(stage, "StrokeDashArrayWhenSelected")
		initializerStatements += rect.GongMarshallField(stage, "Transform")
		pointersInitializesStatements += rect.GongMarshallField(stage, "HoveringTrigger")
		pointersInitializesStatements += rect.GongMarshallField(stage, "DisplayConditions")
		pointersInitializesStatements += rect.GongMarshallField(stage, "Animations")
		initializerStatements += rect.GongMarshallField(stage, "IsSelectable")
		initializerStatements += rect.GongMarshallField(stage, "IsSelected")
		initializerStatements += rect.GongMarshallField(stage, "CanHaveLeftHandle")
		initializerStatements += rect.GongMarshallField(stage, "HasLeftHandle")
		initializerStatements += rect.GongMarshallField(stage, "CanHaveRightHandle")
		initializerStatements += rect.GongMarshallField(stage, "HasRightHandle")
		initializerStatements += rect.GongMarshallField(stage, "CanHaveTopHandle")
		initializerStatements += rect.GongMarshallField(stage, "HasTopHandle")
		initializerStatements += rect.GongMarshallField(stage, "IsScalingProportionally")
		initializerStatements += rect.GongMarshallField(stage, "CanHaveBottomHandle")
		initializerStatements += rect.GongMarshallField(stage, "HasBottomHandle")
		initializerStatements += rect.GongMarshallField(stage, "CanMoveHorizontaly")
		initializerStatements += rect.GongMarshallField(stage, "CanMoveVerticaly")
		pointersInitializesStatements += rect.GongMarshallField(stage, "RectAnchoredTexts")
		pointersInitializesStatements += rect.GongMarshallField(stage, "RectAnchoredRects")
		pointersInitializesStatements += rect.GongMarshallField(stage, "RectAnchoredPaths")
		initializerStatements += rect.GongMarshallField(stage, "ChangeColorWhenHovered")
		initializerStatements += rect.GongMarshallField(stage, "ColorWhenHovered")
		initializerStatements += rect.GongMarshallField(stage, "OriginalColor")
		initializerStatements += rect.GongMarshallField(stage, "FillOpacityWhenHovered")
		initializerStatements += rect.GongMarshallField(stage, "OriginalFillOpacity")
		initializerStatements += rect.GongMarshallField(stage, "HasToolTip")
		initializerStatements += rect.GongMarshallField(stage, "ToolTipText")
		initializerStatements += rect.GongMarshallField(stage, "ToolTipPosition")
		initializerStatements += rect.GongMarshallField(stage, "MouseX")
		initializerStatements += rect.GongMarshallField(stage, "MouseY")
		initializerStatements += rect.GongMarshallField(stage, "MouseEventKey")
	}

	rectanchoredpathOrdered := []*RectAnchoredPath{}
	for rectanchoredpath := range stage.RectAnchoredPaths {
		rectanchoredpathOrdered = append(rectanchoredpathOrdered, rectanchoredpath)
	}
	sort.Slice(rectanchoredpathOrdered[:], func(i, j int) bool {
		rectanchoredpathi := rectanchoredpathOrdered[i]
		rectanchoredpathj := rectanchoredpathOrdered[j]
		rectanchoredpathi_order, oki := stage.RectAnchoredPathMap_Staged_Order[rectanchoredpathi]
		rectanchoredpathj_order, okj := stage.RectAnchoredPathMap_Staged_Order[rectanchoredpathj]
		if !oki || !okj {
			log.Fatalln("unknown pointers")
		}
		return rectanchoredpathi_order < rectanchoredpathj_order
	})
	if len(rectanchoredpathOrdered) > 0 {
		identifiersDecl += "\n"
	}
	for _, rectanchoredpath := range rectanchoredpathOrdered {

		identifiersDecl += rectanchoredpath.GongMarshallIdentifier(stage)

		initializerStatements += "\n"
		// Insertion point for basic fields value assignment
		initializerStatements += rectanchoredpath.GongMarshallField(stage, "Name")
		initializerStatements += rectanchoredpath.GongMarshallField(stage, "Definition")
		initializerStatements += rectanchoredpath.GongMarshallField(stage, "X_Offset")
		initializerStatements += rectanchoredpath.GongMarshallField(stage, "Y_Offset")
		initializerStatements += rectanchoredpath.GongMarshallField(stage, "RectAnchorType")
		initializerStatements += rectanchoredpath.GongMarshallField(stage, "ScalePropotionnally")
		initializerStatements += rectanchoredpath.GongMarshallField(stage, "AppliedScaling")
		initializerStatements += rectanchoredpath.GongMarshallField(stage, "Color")
		initializerStatements += rectanchoredpath.GongMarshallField(stage, "FillOpacity")
		initializerStatements += rectanchoredpath.GongMarshallField(stage, "Stroke")
		initializerStatements += rectanchoredpath.GongMarshallField(stage, "StrokeOpacity")
		initializerStatements += rectanchoredpath.GongMarshallField(stage, "StrokeWidth")
		initializerStatements += rectanchoredpath.GongMarshallField(stage, "StrokeDashArray")
		initializerStatements += rectanchoredpath.GongMarshallField(stage, "StrokeDashArrayWhenSelected")
		initializerStatements += rectanchoredpath.GongMarshallField(stage, "Transform")
	}

	rectanchoredrectOrdered := []*RectAnchoredRect{}
	for rectanchoredrect := range stage.RectAnchoredRects {
		rectanchoredrectOrdered = append(rectanchoredrectOrdered, rectanchoredrect)
	}
	sort.Slice(rectanchoredrectOrdered[:], func(i, j int) bool {
		rectanchoredrecti := rectanchoredrectOrdered[i]
		rectanchoredrectj := rectanchoredrectOrdered[j]
		rectanchoredrecti_order, oki := stage.RectAnchoredRectMap_Staged_Order[rectanchoredrecti]
		rectanchoredrectj_order, okj := stage.RectAnchoredRectMap_Staged_Order[rectanchoredrectj]
		if !oki || !okj {
			log.Fatalln("unknown pointers")
		}
		return rectanchoredrecti_order < rectanchoredrectj_order
	})
	if len(rectanchoredrectOrdered) > 0 {
		identifiersDecl += "\n"
	}
	for _, rectanchoredrect := range rectanchoredrectOrdered {

		identifiersDecl += rectanchoredrect.GongMarshallIdentifier(stage)

		initializerStatements += "\n"
		// Insertion point for basic fields value assignment
		initializerStatements += rectanchoredrect.GongMarshallField(stage, "Name")
		initializerStatements += rectanchoredrect.GongMarshallField(stage, "X")
		initializerStatements += rectanchoredrect.GongMarshallField(stage, "Y")
		initializerStatements += rectanchoredrect.GongMarshallField(stage, "Width")
		initializerStatements += rectanchoredrect.GongMarshallField(stage, "Height")
		initializerStatements += rectanchoredrect.GongMarshallField(stage, "RX")
		initializerStatements += rectanchoredrect.GongMarshallField(stage, "X_Offset")
		initializerStatements += rectanchoredrect.GongMarshallField(stage, "Y_Offset")
		initializerStatements += rectanchoredrect.GongMarshallField(stage, "RectAnchorType")
		initializerStatements += rectanchoredrect.GongMarshallField(stage, "WidthFollowRect")
		initializerStatements += rectanchoredrect.GongMarshallField(stage, "HeightFollowRect")
		initializerStatements += rectanchoredrect.GongMarshallField(stage, "HasToolTip")
		initializerStatements += rectanchoredrect.GongMarshallField(stage, "ToolTipText")
		initializerStatements += rectanchoredrect.GongMarshallField(stage, "Color")
		initializerStatements += rectanchoredrect.GongMarshallField(stage, "FillOpacity")
		initializerStatements += rectanchoredrect.GongMarshallField(stage, "Stroke")
		initializerStatements += rectanchoredrect.GongMarshallField(stage, "StrokeOpacity")
		initializerStatements += rectanchoredrect.GongMarshallField(stage, "StrokeWidth")
		initializerStatements += rectanchoredrect.GongMarshallField(stage, "StrokeDashArray")
		initializerStatements += rectanchoredrect.GongMarshallField(stage, "StrokeDashArrayWhenSelected")
		initializerStatements += rectanchoredrect.GongMarshallField(stage, "Transform")
	}

	rectanchoredtextOrdered := []*RectAnchoredText{}
	for rectanchoredtext := range stage.RectAnchoredTexts {
		rectanchoredtextOrdered = append(rectanchoredtextOrdered, rectanchoredtext)
	}
	sort.Slice(rectanchoredtextOrdered[:], func(i, j int) bool {
		rectanchoredtexti := rectanchoredtextOrdered[i]
		rectanchoredtextj := rectanchoredtextOrdered[j]
		rectanchoredtexti_order, oki := stage.RectAnchoredTextMap_Staged_Order[rectanchoredtexti]
		rectanchoredtextj_order, okj := stage.RectAnchoredTextMap_Staged_Order[rectanchoredtextj]
		if !oki || !okj {
			log.Fatalln("unknown pointers")
		}
		return rectanchoredtexti_order < rectanchoredtextj_order
	})
	if len(rectanchoredtextOrdered) > 0 {
		identifiersDecl += "\n"
	}
	for _, rectanchoredtext := range rectanchoredtextOrdered {

		identifiersDecl += rectanchoredtext.GongMarshallIdentifier(stage)

		initializerStatements += "\n"
		// Insertion point for basic fields value assignment
		initializerStatements += rectanchoredtext.GongMarshallField(stage, "Name")
		initializerStatements += rectanchoredtext.GongMarshallField(stage, "Content")
		initializerStatements += rectanchoredtext.GongMarshallField(stage, "FontWeight")
		initializerStatements += rectanchoredtext.GongMarshallField(stage, "FontSize")
		initializerStatements += rectanchoredtext.GongMarshallField(stage, "FontStyle")
		initializerStatements += rectanchoredtext.GongMarshallField(stage, "LetterSpacing")
		initializerStatements += rectanchoredtext.GongMarshallField(stage, "FontFamily")
		initializerStatements += rectanchoredtext.GongMarshallField(stage, "WhiteSpace")
		initializerStatements += rectanchoredtext.GongMarshallField(stage, "X_Offset")
		initializerStatements += rectanchoredtext.GongMarshallField(stage, "Y_Offset")
		initializerStatements += rectanchoredtext.GongMarshallField(stage, "RectAnchorType")
		initializerStatements += rectanchoredtext.GongMarshallField(stage, "TextAnchorType")
		initializerStatements += rectanchoredtext.GongMarshallField(stage, "DominantBaseline")
		initializerStatements += rectanchoredtext.GongMarshallField(stage, "WritingMode")
		initializerStatements += rectanchoredtext.GongMarshallField(stage, "Color")
		initializerStatements += rectanchoredtext.GongMarshallField(stage, "FillOpacity")
		initializerStatements += rectanchoredtext.GongMarshallField(stage, "Stroke")
		initializerStatements += rectanchoredtext.GongMarshallField(stage, "StrokeOpacity")
		initializerStatements += rectanchoredtext.GongMarshallField(stage, "StrokeWidth")
		initializerStatements += rectanchoredtext.GongMarshallField(stage, "StrokeDashArray")
		initializerStatements += rectanchoredtext.GongMarshallField(stage, "StrokeDashArrayWhenSelected")
		initializerStatements += rectanchoredtext.GongMarshallField(stage, "Transform")
		pointersInitializesStatements += rectanchoredtext.GongMarshallField(stage, "Animates")
	}

	rectlinklinkOrdered := []*RectLinkLink{}
	for rectlinklink := range stage.RectLinkLinks {
		rectlinklinkOrdered = append(rectlinklinkOrdered, rectlinklink)
	}
	sort.Slice(rectlinklinkOrdered[:], func(i, j int) bool {
		rectlinklinki := rectlinklinkOrdered[i]
		rectlinklinkj := rectlinklinkOrdered[j]
		rectlinklinki_order, oki := stage.RectLinkLinkMap_Staged_Order[rectlinklinki]
		rectlinklinkj_order, okj := stage.RectLinkLinkMap_Staged_Order[rectlinklinkj]
		if !oki || !okj {
			log.Fatalln("unknown pointers")
		}
		return rectlinklinki_order < rectlinklinkj_order
	})
	if len(rectlinklinkOrdered) > 0 {
		identifiersDecl += "\n"
	}
	for _, rectlinklink := range rectlinklinkOrdered {

		identifiersDecl += rectlinklink.GongMarshallIdentifier(stage)

		initializerStatements += "\n"
		// Insertion point for basic fields value assignment
		initializerStatements += rectlinklink.GongMarshallField(stage, "Name")
		pointersInitializesStatements += rectlinklink.GongMarshallField(stage, "Start")
		pointersInitializesStatements += rectlinklink.GongMarshallField(stage, "End")
		initializerStatements += rectlinklink.GongMarshallField(stage, "TargetAnchorPosition")
		initializerStatements += rectlinklink.GongMarshallField(stage, "Color")
		initializerStatements += rectlinklink.GongMarshallField(stage, "FillOpacity")
		initializerStatements += rectlinklink.GongMarshallField(stage, "Stroke")
		initializerStatements += rectlinklink.GongMarshallField(stage, "StrokeOpacity")
		initializerStatements += rectlinklink.GongMarshallField(stage, "StrokeWidth")
		initializerStatements += rectlinklink.GongMarshallField(stage, "StrokeDashArray")
		initializerStatements += rectlinklink.GongMarshallField(stage, "StrokeDashArrayWhenSelected")
		initializerStatements += rectlinklink.GongMarshallField(stage, "Transform")
	}

	svgOrdered := []*SVG{}
	for svg := range stage.SVGs {
		svgOrdered = append(svgOrdered, svg)
	}
	sort.Slice(svgOrdered[:], func(i, j int) bool {
		svgi := svgOrdered[i]
		svgj := svgOrdered[j]
		svgi_order, oki := stage.SVGMap_Staged_Order[svgi]
		svgj_order, okj := stage.SVGMap_Staged_Order[svgj]
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
		// Insertion point for basic fields value assignment
		initializerStatements += svg.GongMarshallField(stage, "Name")
		pointersInitializesStatements += svg.GongMarshallField(stage, "Layers")
		initializerStatements += svg.GongMarshallField(stage, "DrawingState")
		pointersInitializesStatements += svg.GongMarshallField(stage, "StartRect")
		pointersInitializesStatements += svg.GongMarshallField(stage, "EndRect")
		initializerStatements += svg.GongMarshallField(stage, "IsEditable")
		initializerStatements += svg.GongMarshallField(stage, "IsSVGFrontEndFileGenerated")
		initializerStatements += svg.GongMarshallField(stage, "IsSVGBackEndFileGenerated")
		initializerStatements += svg.GongMarshallField(stage, "DefaultDirectoryForGeneratedImages")
		initializerStatements += svg.GongMarshallField(stage, "IsControlBannerHidden")
		initializerStatements += svg.GongMarshallField(stage, "OverrideWidth")
		initializerStatements += svg.GongMarshallField(stage, "OverriddenWidth")
		initializerStatements += svg.GongMarshallField(stage, "OverrideHeight")
		initializerStatements += svg.GongMarshallField(stage, "OverriddenHeight")
	}

	svgtextOrdered := []*SvgText{}
	for svgtext := range stage.SvgTexts {
		svgtextOrdered = append(svgtextOrdered, svgtext)
	}
	sort.Slice(svgtextOrdered[:], func(i, j int) bool {
		svgtexti := svgtextOrdered[i]
		svgtextj := svgtextOrdered[j]
		svgtexti_order, oki := stage.SvgTextMap_Staged_Order[svgtexti]
		svgtextj_order, okj := stage.SvgTextMap_Staged_Order[svgtextj]
		if !oki || !okj {
			log.Fatalln("unknown pointers")
		}
		return svgtexti_order < svgtextj_order
	})
	if len(svgtextOrdered) > 0 {
		identifiersDecl += "\n"
	}
	for _, svgtext := range svgtextOrdered {

		identifiersDecl += svgtext.GongMarshallIdentifier(stage)

		initializerStatements += "\n"
		// Insertion point for basic fields value assignment
		initializerStatements += svgtext.GongMarshallField(stage, "Name")
		initializerStatements += svgtext.GongMarshallField(stage, "Text")
	}

	textOrdered := []*Text{}
	for text := range stage.Texts {
		textOrdered = append(textOrdered, text)
	}
	sort.Slice(textOrdered[:], func(i, j int) bool {
		texti := textOrdered[i]
		textj := textOrdered[j]
		texti_order, oki := stage.TextMap_Staged_Order[texti]
		textj_order, okj := stage.TextMap_Staged_Order[textj]
		if !oki || !okj {
			log.Fatalln("unknown pointers")
		}
		return texti_order < textj_order
	})
	if len(textOrdered) > 0 {
		identifiersDecl += "\n"
	}
	for _, text := range textOrdered {

		identifiersDecl += text.GongMarshallIdentifier(stage)

		initializerStatements += "\n"
		// Insertion point for basic fields value assignment
		initializerStatements += text.GongMarshallField(stage, "Name")
		initializerStatements += text.GongMarshallField(stage, "X")
		initializerStatements += text.GongMarshallField(stage, "Y")
		initializerStatements += text.GongMarshallField(stage, "Content")
		initializerStatements += text.GongMarshallField(stage, "Color")
		initializerStatements += text.GongMarshallField(stage, "FillOpacity")
		initializerStatements += text.GongMarshallField(stage, "Stroke")
		initializerStatements += text.GongMarshallField(stage, "StrokeOpacity")
		initializerStatements += text.GongMarshallField(stage, "StrokeWidth")
		initializerStatements += text.GongMarshallField(stage, "StrokeDashArray")
		initializerStatements += text.GongMarshallField(stage, "StrokeDashArrayWhenSelected")
		initializerStatements += text.GongMarshallField(stage, "Transform")
		initializerStatements += text.GongMarshallField(stage, "FontWeight")
		initializerStatements += text.GongMarshallField(stage, "FontSize")
		initializerStatements += text.GongMarshallField(stage, "FontStyle")
		initializerStatements += text.GongMarshallField(stage, "LetterSpacing")
		initializerStatements += text.GongMarshallField(stage, "FontFamily")
		initializerStatements += text.GongMarshallField(stage, "WhiteSpace")
		pointersInitializesStatements += text.GongMarshallField(stage, "Animates")
	}

	// insertion initialization of objects to stage
	for _, animate := range animateOrdered {
		_ = animate
		var setPointerField string
		_ = setPointerField

		// Insertion point for pointers initialization
	}

	for _, circle := range circleOrdered {
		_ = circle
		var setPointerField string
		_ = setPointerField

		// Insertion point for pointers initialization
	}

	for _, condition := range conditionOrdered {
		_ = condition
		var setPointerField string
		_ = setPointerField

		// Insertion point for pointers initialization
	}

	for _, controlpoint := range controlpointOrdered {
		_ = controlpoint
		var setPointerField string
		_ = setPointerField

		// Insertion point for pointers initialization
	}

	for _, ellipse := range ellipseOrdered {
		_ = ellipse
		var setPointerField string
		_ = setPointerField

		// Insertion point for pointers initialization
	}

	for _, layer := range layerOrdered {
		_ = layer
		var setPointerField string
		_ = setPointerField

		// Insertion point for pointers initialization
	}

	for _, line := range lineOrdered {
		_ = line
		var setPointerField string
		_ = setPointerField

		// Insertion point for pointers initialization
	}

	for _, link := range linkOrdered {
		_ = link
		var setPointerField string
		_ = setPointerField

		// Insertion point for pointers initialization
	}

	for _, linkanchoredtext := range linkanchoredtextOrdered {
		_ = linkanchoredtext
		var setPointerField string
		_ = setPointerField

		// Insertion point for pointers initialization
	}

	for _, path := range pathOrdered {
		_ = path
		var setPointerField string
		_ = setPointerField

		// Insertion point for pointers initialization
	}

	for _, point := range pointOrdered {
		_ = point
		var setPointerField string
		_ = setPointerField

		// Insertion point for pointers initialization
	}

	for _, polygone := range polygoneOrdered {
		_ = polygone
		var setPointerField string
		_ = setPointerField

		// Insertion point for pointers initialization
	}

	for _, polyline := range polylineOrdered {
		_ = polyline
		var setPointerField string
		_ = setPointerField

		// Insertion point for pointers initialization
	}

	for _, rect := range rectOrdered {
		_ = rect
		var setPointerField string
		_ = setPointerField

		// Insertion point for pointers initialization
	}

	for _, rectanchoredpath := range rectanchoredpathOrdered {
		_ = rectanchoredpath
		var setPointerField string
		_ = setPointerField

		// Insertion point for pointers initialization
	}

	for _, rectanchoredrect := range rectanchoredrectOrdered {
		_ = rectanchoredrect
		var setPointerField string
		_ = setPointerField

		// Insertion point for pointers initialization
	}

	for _, rectanchoredtext := range rectanchoredtextOrdered {
		_ = rectanchoredtext
		var setPointerField string
		_ = setPointerField

		// Insertion point for pointers initialization
	}

	for _, rectlinklink := range rectlinklinkOrdered {
		_ = rectlinklink
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

	for _, svgtext := range svgtextOrdered {
		_ = svgtext
		var setPointerField string
		_ = setPointerField

		// Insertion point for pointers initialization
	}

	for _, text := range textOrdered {
		_ = text
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

// insertion initialization of objects to stage
func (animate *Animate) GongMarshallField(stage *Stage, fieldName string) (res string) {
	var setValueField, setPointerField string
	_ = setValueField
	_ = setPointerField
	initializerStatements := ""
	_ = initializerStatements
	pointersInitializesStatements := ""
	_ = pointersInitializesStatements

	switch fieldName {
	case "Name":
		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", animate.GongGetIdentifier(stage))
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Name")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(animate.Name))
		initializerStatements += setValueField
	case "AttributeName":
		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", animate.GongGetIdentifier(stage))
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "AttributeName")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(animate.AttributeName))
		initializerStatements += setValueField
	case "Values":
		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", animate.GongGetIdentifier(stage))
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Values")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(animate.Values))
		initializerStatements += setValueField
	case "From":
		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", animate.GongGetIdentifier(stage))
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "From")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(animate.From))
		initializerStatements += setValueField
	case "To":
		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", animate.GongGetIdentifier(stage))
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "To")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(animate.To))
		initializerStatements += setValueField
	case "Dur":
		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", animate.GongGetIdentifier(stage))
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Dur")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(animate.Dur))
		initializerStatements += setValueField
	case "RepeatCount":
		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", animate.GongGetIdentifier(stage))
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "RepeatCount")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(animate.RepeatCount))
		initializerStatements += setValueField

	default:
		log.Panicf("Unknown field %s for Gongstruct Animate", fieldName)
	}

	// temporary kludge to reuse existing template code
	res = initializerStatements + pointersInitializesStatements
	return
}

func (circle *Circle) GongMarshallField(stage *Stage, fieldName string) (res string) {
	var setValueField, setPointerField string
	_ = setValueField
	_ = setPointerField
	initializerStatements := ""
	_ = initializerStatements
	pointersInitializesStatements := ""
	_ = pointersInitializesStatements

	switch fieldName {
	case "Name":
		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", circle.GongGetIdentifier(stage))
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Name")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(circle.Name))
		initializerStatements += setValueField
	case "CX":
		setValueField = NumberInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", circle.GongGetIdentifier(stage))
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "CX")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%f", circle.CX))
		initializerStatements += setValueField
	case "CY":
		setValueField = NumberInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", circle.GongGetIdentifier(stage))
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "CY")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%f", circle.CY))
		initializerStatements += setValueField
	case "Radius":
		setValueField = NumberInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", circle.GongGetIdentifier(stage))
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Radius")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%f", circle.Radius))
		initializerStatements += setValueField
	case "Color":
		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", circle.GongGetIdentifier(stage))
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Color")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(circle.Color))
		initializerStatements += setValueField
	case "FillOpacity":
		setValueField = NumberInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", circle.GongGetIdentifier(stage))
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "FillOpacity")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%f", circle.FillOpacity))
		initializerStatements += setValueField
	case "Stroke":
		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", circle.GongGetIdentifier(stage))
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Stroke")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(circle.Stroke))
		initializerStatements += setValueField
	case "StrokeOpacity":
		setValueField = NumberInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", circle.GongGetIdentifier(stage))
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "StrokeOpacity")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%f", circle.StrokeOpacity))
		initializerStatements += setValueField
	case "StrokeWidth":
		setValueField = NumberInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", circle.GongGetIdentifier(stage))
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "StrokeWidth")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%f", circle.StrokeWidth))
		initializerStatements += setValueField
	case "StrokeDashArray":
		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", circle.GongGetIdentifier(stage))
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "StrokeDashArray")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(circle.StrokeDashArray))
		initializerStatements += setValueField
	case "StrokeDashArrayWhenSelected":
		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", circle.GongGetIdentifier(stage))
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "StrokeDashArrayWhenSelected")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(circle.StrokeDashArrayWhenSelected))
		initializerStatements += setValueField
	case "Transform":
		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", circle.GongGetIdentifier(stage))
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Transform")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(circle.Transform))
		initializerStatements += setValueField

	case "Animations":
		for _, _animate := range circle.Animations {
			setPointerField = SliceOfPointersFieldInitStatement
			setPointerField = strings.ReplaceAll(setPointerField, "{{Identifier}}", circle.GongGetIdentifier(stage))
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldName}}", "Animations")
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldNameValue}}", _animate.GongGetIdentifier(stage))
			pointersInitializesStatements += setPointerField
		}
	default:
		log.Panicf("Unknown field %s for Gongstruct Circle", fieldName)
	}

	// temporary kludge to reuse existing template code
	res = initializerStatements + pointersInitializesStatements
	return
}

func (condition *Condition) GongMarshallField(stage *Stage, fieldName string) (res string) {
	var setValueField, setPointerField string
	_ = setValueField
	_ = setPointerField
	initializerStatements := ""
	_ = initializerStatements
	pointersInitializesStatements := ""
	_ = pointersInitializesStatements

	switch fieldName {
	case "Name":
		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", condition.GongGetIdentifier(stage))
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Name")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(condition.Name))
		initializerStatements += setValueField

	default:
		log.Panicf("Unknown field %s for Gongstruct Condition", fieldName)
	}

	// temporary kludge to reuse existing template code
	res = initializerStatements + pointersInitializesStatements
	return
}

func (controlpoint *ControlPoint) GongMarshallField(stage *Stage, fieldName string) (res string) {
	var setValueField, setPointerField string
	_ = setValueField
	_ = setPointerField
	initializerStatements := ""
	_ = initializerStatements
	pointersInitializesStatements := ""
	_ = pointersInitializesStatements

	switch fieldName {
	case "Name":
		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", controlpoint.GongGetIdentifier(stage))
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Name")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(controlpoint.Name))
		initializerStatements += setValueField
	case "X_Relative":
		setValueField = NumberInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", controlpoint.GongGetIdentifier(stage))
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "X_Relative")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%f", controlpoint.X_Relative))
		initializerStatements += setValueField
	case "Y_Relative":
		setValueField = NumberInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", controlpoint.GongGetIdentifier(stage))
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Y_Relative")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%f", controlpoint.Y_Relative))
		initializerStatements += setValueField

	case "ClosestRect":
		if controlpoint.ClosestRect != nil {
			setPointerField = PointerFieldInitStatement
			setPointerField = strings.ReplaceAll(setPointerField, "{{Identifier}}", controlpoint.GongGetIdentifier(stage))
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldName}}", "ClosestRect")
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldNameValue}}", controlpoint.ClosestRect.GongGetIdentifier(stage))
			pointersInitializesStatements += setPointerField
		}
	default:
		log.Panicf("Unknown field %s for Gongstruct ControlPoint", fieldName)
	}

	// temporary kludge to reuse existing template code
	res = initializerStatements + pointersInitializesStatements
	return
}

func (ellipse *Ellipse) GongMarshallField(stage *Stage, fieldName string) (res string) {
	var setValueField, setPointerField string
	_ = setValueField
	_ = setPointerField
	initializerStatements := ""
	_ = initializerStatements
	pointersInitializesStatements := ""
	_ = pointersInitializesStatements

	switch fieldName {
	case "Name":
		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", ellipse.GongGetIdentifier(stage))
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Name")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(ellipse.Name))
		initializerStatements += setValueField
	case "CX":
		setValueField = NumberInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", ellipse.GongGetIdentifier(stage))
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "CX")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%f", ellipse.CX))
		initializerStatements += setValueField
	case "CY":
		setValueField = NumberInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", ellipse.GongGetIdentifier(stage))
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "CY")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%f", ellipse.CY))
		initializerStatements += setValueField
	case "RX":
		setValueField = NumberInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", ellipse.GongGetIdentifier(stage))
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "RX")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%f", ellipse.RX))
		initializerStatements += setValueField
	case "RY":
		setValueField = NumberInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", ellipse.GongGetIdentifier(stage))
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "RY")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%f", ellipse.RY))
		initializerStatements += setValueField
	case "Color":
		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", ellipse.GongGetIdentifier(stage))
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Color")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(ellipse.Color))
		initializerStatements += setValueField
	case "FillOpacity":
		setValueField = NumberInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", ellipse.GongGetIdentifier(stage))
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "FillOpacity")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%f", ellipse.FillOpacity))
		initializerStatements += setValueField
	case "Stroke":
		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", ellipse.GongGetIdentifier(stage))
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Stroke")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(ellipse.Stroke))
		initializerStatements += setValueField
	case "StrokeOpacity":
		setValueField = NumberInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", ellipse.GongGetIdentifier(stage))
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "StrokeOpacity")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%f", ellipse.StrokeOpacity))
		initializerStatements += setValueField
	case "StrokeWidth":
		setValueField = NumberInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", ellipse.GongGetIdentifier(stage))
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "StrokeWidth")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%f", ellipse.StrokeWidth))
		initializerStatements += setValueField
	case "StrokeDashArray":
		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", ellipse.GongGetIdentifier(stage))
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "StrokeDashArray")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(ellipse.StrokeDashArray))
		initializerStatements += setValueField
	case "StrokeDashArrayWhenSelected":
		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", ellipse.GongGetIdentifier(stage))
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "StrokeDashArrayWhenSelected")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(ellipse.StrokeDashArrayWhenSelected))
		initializerStatements += setValueField
	case "Transform":
		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", ellipse.GongGetIdentifier(stage))
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Transform")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(ellipse.Transform))
		initializerStatements += setValueField

	case "Animates":
		for _, _animate := range ellipse.Animates {
			setPointerField = SliceOfPointersFieldInitStatement
			setPointerField = strings.ReplaceAll(setPointerField, "{{Identifier}}", ellipse.GongGetIdentifier(stage))
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldName}}", "Animates")
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldNameValue}}", _animate.GongGetIdentifier(stage))
			pointersInitializesStatements += setPointerField
		}
	default:
		log.Panicf("Unknown field %s for Gongstruct Ellipse", fieldName)
	}

	// temporary kludge to reuse existing template code
	res = initializerStatements + pointersInitializesStatements
	return
}

func (layer *Layer) GongMarshallField(stage *Stage, fieldName string) (res string) {
	var setValueField, setPointerField string
	_ = setValueField
	_ = setPointerField
	initializerStatements := ""
	_ = initializerStatements
	pointersInitializesStatements := ""
	_ = pointersInitializesStatements

	switch fieldName {
	case "Name":
		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", layer.GongGetIdentifier(stage))
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Name")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(layer.Name))
		initializerStatements += setValueField

	case "Rects":
		for _, _rect := range layer.Rects {
			setPointerField = SliceOfPointersFieldInitStatement
			setPointerField = strings.ReplaceAll(setPointerField, "{{Identifier}}", layer.GongGetIdentifier(stage))
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldName}}", "Rects")
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldNameValue}}", _rect.GongGetIdentifier(stage))
			pointersInitializesStatements += setPointerField
		}
	case "Texts":
		for _, _text := range layer.Texts {
			setPointerField = SliceOfPointersFieldInitStatement
			setPointerField = strings.ReplaceAll(setPointerField, "{{Identifier}}", layer.GongGetIdentifier(stage))
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldName}}", "Texts")
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldNameValue}}", _text.GongGetIdentifier(stage))
			pointersInitializesStatements += setPointerField
		}
	case "Circles":
		for _, _circle := range layer.Circles {
			setPointerField = SliceOfPointersFieldInitStatement
			setPointerField = strings.ReplaceAll(setPointerField, "{{Identifier}}", layer.GongGetIdentifier(stage))
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldName}}", "Circles")
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldNameValue}}", _circle.GongGetIdentifier(stage))
			pointersInitializesStatements += setPointerField
		}
	case "Lines":
		for _, _line := range layer.Lines {
			setPointerField = SliceOfPointersFieldInitStatement
			setPointerField = strings.ReplaceAll(setPointerField, "{{Identifier}}", layer.GongGetIdentifier(stage))
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldName}}", "Lines")
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldNameValue}}", _line.GongGetIdentifier(stage))
			pointersInitializesStatements += setPointerField
		}
	case "Ellipses":
		for _, _ellipse := range layer.Ellipses {
			setPointerField = SliceOfPointersFieldInitStatement
			setPointerField = strings.ReplaceAll(setPointerField, "{{Identifier}}", layer.GongGetIdentifier(stage))
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldName}}", "Ellipses")
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldNameValue}}", _ellipse.GongGetIdentifier(stage))
			pointersInitializesStatements += setPointerField
		}
	case "Polylines":
		for _, _polyline := range layer.Polylines {
			setPointerField = SliceOfPointersFieldInitStatement
			setPointerField = strings.ReplaceAll(setPointerField, "{{Identifier}}", layer.GongGetIdentifier(stage))
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldName}}", "Polylines")
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldNameValue}}", _polyline.GongGetIdentifier(stage))
			pointersInitializesStatements += setPointerField
		}
	case "Polygones":
		for _, _polygone := range layer.Polygones {
			setPointerField = SliceOfPointersFieldInitStatement
			setPointerField = strings.ReplaceAll(setPointerField, "{{Identifier}}", layer.GongGetIdentifier(stage))
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldName}}", "Polygones")
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldNameValue}}", _polygone.GongGetIdentifier(stage))
			pointersInitializesStatements += setPointerField
		}
	case "Paths":
		for _, _path := range layer.Paths {
			setPointerField = SliceOfPointersFieldInitStatement
			setPointerField = strings.ReplaceAll(setPointerField, "{{Identifier}}", layer.GongGetIdentifier(stage))
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldName}}", "Paths")
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldNameValue}}", _path.GongGetIdentifier(stage))
			pointersInitializesStatements += setPointerField
		}
	case "Links":
		for _, _link := range layer.Links {
			setPointerField = SliceOfPointersFieldInitStatement
			setPointerField = strings.ReplaceAll(setPointerField, "{{Identifier}}", layer.GongGetIdentifier(stage))
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldName}}", "Links")
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldNameValue}}", _link.GongGetIdentifier(stage))
			pointersInitializesStatements += setPointerField
		}
	case "RectLinkLinks":
		for _, _rectlinklink := range layer.RectLinkLinks {
			setPointerField = SliceOfPointersFieldInitStatement
			setPointerField = strings.ReplaceAll(setPointerField, "{{Identifier}}", layer.GongGetIdentifier(stage))
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldName}}", "RectLinkLinks")
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldNameValue}}", _rectlinklink.GongGetIdentifier(stage))
			pointersInitializesStatements += setPointerField
		}
	default:
		log.Panicf("Unknown field %s for Gongstruct Layer", fieldName)
	}

	// temporary kludge to reuse existing template code
	res = initializerStatements + pointersInitializesStatements
	return
}

func (line *Line) GongMarshallField(stage *Stage, fieldName string) (res string) {
	var setValueField, setPointerField string
	_ = setValueField
	_ = setPointerField
	initializerStatements := ""
	_ = initializerStatements
	pointersInitializesStatements := ""
	_ = pointersInitializesStatements

	switch fieldName {
	case "Name":
		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", line.GongGetIdentifier(stage))
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Name")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(line.Name))
		initializerStatements += setValueField
	case "X1":
		setValueField = NumberInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", line.GongGetIdentifier(stage))
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "X1")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%f", line.X1))
		initializerStatements += setValueField
	case "Y1":
		setValueField = NumberInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", line.GongGetIdentifier(stage))
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Y1")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%f", line.Y1))
		initializerStatements += setValueField
	case "X2":
		setValueField = NumberInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", line.GongGetIdentifier(stage))
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "X2")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%f", line.X2))
		initializerStatements += setValueField
	case "Y2":
		setValueField = NumberInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", line.GongGetIdentifier(stage))
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Y2")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%f", line.Y2))
		initializerStatements += setValueField
	case "Color":
		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", line.GongGetIdentifier(stage))
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Color")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(line.Color))
		initializerStatements += setValueField
	case "FillOpacity":
		setValueField = NumberInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", line.GongGetIdentifier(stage))
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "FillOpacity")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%f", line.FillOpacity))
		initializerStatements += setValueField
	case "Stroke":
		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", line.GongGetIdentifier(stage))
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Stroke")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(line.Stroke))
		initializerStatements += setValueField
	case "StrokeOpacity":
		setValueField = NumberInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", line.GongGetIdentifier(stage))
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "StrokeOpacity")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%f", line.StrokeOpacity))
		initializerStatements += setValueField
	case "StrokeWidth":
		setValueField = NumberInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", line.GongGetIdentifier(stage))
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "StrokeWidth")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%f", line.StrokeWidth))
		initializerStatements += setValueField
	case "StrokeDashArray":
		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", line.GongGetIdentifier(stage))
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "StrokeDashArray")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(line.StrokeDashArray))
		initializerStatements += setValueField
	case "StrokeDashArrayWhenSelected":
		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", line.GongGetIdentifier(stage))
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "StrokeDashArrayWhenSelected")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(line.StrokeDashArrayWhenSelected))
		initializerStatements += setValueField
	case "Transform":
		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", line.GongGetIdentifier(stage))
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Transform")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(line.Transform))
		initializerStatements += setValueField
	case "MouseClickX":
		setValueField = NumberInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", line.GongGetIdentifier(stage))
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "MouseClickX")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%f", line.MouseClickX))
		initializerStatements += setValueField
	case "MouseClickY":
		setValueField = NumberInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", line.GongGetIdentifier(stage))
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "MouseClickY")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%f", line.MouseClickY))
		initializerStatements += setValueField

	case "Animates":
		for _, _animate := range line.Animates {
			setPointerField = SliceOfPointersFieldInitStatement
			setPointerField = strings.ReplaceAll(setPointerField, "{{Identifier}}", line.GongGetIdentifier(stage))
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldName}}", "Animates")
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldNameValue}}", _animate.GongGetIdentifier(stage))
			pointersInitializesStatements += setPointerField
		}
	default:
		log.Panicf("Unknown field %s for Gongstruct Line", fieldName)
	}

	// temporary kludge to reuse existing template code
	res = initializerStatements + pointersInitializesStatements
	return
}

func (link *Link) GongMarshallField(stage *Stage, fieldName string) (res string) {
	var setValueField, setPointerField string
	_ = setValueField
	_ = setPointerField
	initializerStatements := ""
	_ = initializerStatements
	pointersInitializesStatements := ""
	_ = pointersInitializesStatements

	switch fieldName {
	case "Name":
		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", link.GongGetIdentifier(stage))
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Name")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(link.Name))
		initializerStatements += setValueField
	case "Type":
		if link.Type != "" {
			setValueField = StringEnumInitStatement
			setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", link.GongGetIdentifier(stage))
			setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Type")
			setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", "models."+link.Type.ToCodeString())
			initializerStatements += setValueField
		}
	case "IsBezierCurve":
		setValueField = NumberInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", link.GongGetIdentifier(stage))
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "IsBezierCurve")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%t", link.IsBezierCurve))
		initializerStatements += setValueField
	case "StartAnchorType":
		if link.StartAnchorType != "" {
			setValueField = StringEnumInitStatement
			setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", link.GongGetIdentifier(stage))
			setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "StartAnchorType")
			setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", "models."+link.StartAnchorType.ToCodeString())
			initializerStatements += setValueField
		}
	case "EndAnchorType":
		if link.EndAnchorType != "" {
			setValueField = StringEnumInitStatement
			setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", link.GongGetIdentifier(stage))
			setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "EndAnchorType")
			setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", "models."+link.EndAnchorType.ToCodeString())
			initializerStatements += setValueField
		}
	case "StartOrientation":
		if link.StartOrientation != "" {
			setValueField = StringEnumInitStatement
			setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", link.GongGetIdentifier(stage))
			setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "StartOrientation")
			setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", "models."+link.StartOrientation.ToCodeString())
			initializerStatements += setValueField
		}
	case "StartRatio":
		setValueField = NumberInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", link.GongGetIdentifier(stage))
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "StartRatio")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%f", link.StartRatio))
		initializerStatements += setValueField
	case "EndOrientation":
		if link.EndOrientation != "" {
			setValueField = StringEnumInitStatement
			setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", link.GongGetIdentifier(stage))
			setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "EndOrientation")
			setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", "models."+link.EndOrientation.ToCodeString())
			initializerStatements += setValueField
		}
	case "EndRatio":
		setValueField = NumberInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", link.GongGetIdentifier(stage))
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "EndRatio")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%f", link.EndRatio))
		initializerStatements += setValueField
	case "CornerOffsetRatio":
		setValueField = NumberInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", link.GongGetIdentifier(stage))
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "CornerOffsetRatio")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%f", link.CornerOffsetRatio))
		initializerStatements += setValueField
	case "CornerRadius":
		setValueField = NumberInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", link.GongGetIdentifier(stage))
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "CornerRadius")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%f", link.CornerRadius))
		initializerStatements += setValueField
	case "HasEndArrow":
		setValueField = NumberInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", link.GongGetIdentifier(stage))
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "HasEndArrow")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%t", link.HasEndArrow))
		initializerStatements += setValueField
	case "EndArrowSize":
		setValueField = NumberInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", link.GongGetIdentifier(stage))
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "EndArrowSize")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%f", link.EndArrowSize))
		initializerStatements += setValueField
	case "EndArrowOffset":
		setValueField = NumberInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", link.GongGetIdentifier(stage))
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "EndArrowOffset")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%f", link.EndArrowOffset))
		initializerStatements += setValueField
	case "HasStartArrow":
		setValueField = NumberInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", link.GongGetIdentifier(stage))
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "HasStartArrow")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%t", link.HasStartArrow))
		initializerStatements += setValueField
	case "StartArrowSize":
		setValueField = NumberInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", link.GongGetIdentifier(stage))
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "StartArrowSize")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%f", link.StartArrowSize))
		initializerStatements += setValueField
	case "StartArrowOffset":
		setValueField = NumberInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", link.GongGetIdentifier(stage))
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "StartArrowOffset")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%f", link.StartArrowOffset))
		initializerStatements += setValueField
	case "Color":
		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", link.GongGetIdentifier(stage))
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Color")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(link.Color))
		initializerStatements += setValueField
	case "FillOpacity":
		setValueField = NumberInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", link.GongGetIdentifier(stage))
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "FillOpacity")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%f", link.FillOpacity))
		initializerStatements += setValueField
	case "Stroke":
		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", link.GongGetIdentifier(stage))
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Stroke")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(link.Stroke))
		initializerStatements += setValueField
	case "StrokeOpacity":
		setValueField = NumberInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", link.GongGetIdentifier(stage))
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "StrokeOpacity")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%f", link.StrokeOpacity))
		initializerStatements += setValueField
	case "StrokeWidth":
		setValueField = NumberInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", link.GongGetIdentifier(stage))
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "StrokeWidth")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%f", link.StrokeWidth))
		initializerStatements += setValueField
	case "StrokeDashArray":
		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", link.GongGetIdentifier(stage))
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "StrokeDashArray")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(link.StrokeDashArray))
		initializerStatements += setValueField
	case "StrokeDashArrayWhenSelected":
		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", link.GongGetIdentifier(stage))
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "StrokeDashArrayWhenSelected")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(link.StrokeDashArrayWhenSelected))
		initializerStatements += setValueField
	case "Transform":
		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", link.GongGetIdentifier(stage))
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Transform")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(link.Transform))
		initializerStatements += setValueField
	case "MouseX":
		setValueField = NumberInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", link.GongGetIdentifier(stage))
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "MouseX")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%f", link.MouseX))
		initializerStatements += setValueField
	case "MouseY":
		setValueField = NumberInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", link.GongGetIdentifier(stage))
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "MouseY")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%f", link.MouseY))
		initializerStatements += setValueField
	case "MouseEventKey":
		if link.MouseEventKey != "" {
			setValueField = StringEnumInitStatement
			setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", link.GongGetIdentifier(stage))
			setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "MouseEventKey")
			setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", "models."+link.MouseEventKey.ToCodeString())
			initializerStatements += setValueField
		}

	case "Start":
		if link.Start != nil {
			setPointerField = PointerFieldInitStatement
			setPointerField = strings.ReplaceAll(setPointerField, "{{Identifier}}", link.GongGetIdentifier(stage))
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldName}}", "Start")
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldNameValue}}", link.Start.GongGetIdentifier(stage))
			pointersInitializesStatements += setPointerField
		}
	case "End":
		if link.End != nil {
			setPointerField = PointerFieldInitStatement
			setPointerField = strings.ReplaceAll(setPointerField, "{{Identifier}}", link.GongGetIdentifier(stage))
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldName}}", "End")
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldNameValue}}", link.End.GongGetIdentifier(stage))
			pointersInitializesStatements += setPointerField
		}
	case "TextAtArrowStart":
		for _, _linkanchoredtext := range link.TextAtArrowStart {
			setPointerField = SliceOfPointersFieldInitStatement
			setPointerField = strings.ReplaceAll(setPointerField, "{{Identifier}}", link.GongGetIdentifier(stage))
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldName}}", "TextAtArrowStart")
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldNameValue}}", _linkanchoredtext.GongGetIdentifier(stage))
			pointersInitializesStatements += setPointerField
		}
	case "TextAtArrowEnd":
		for _, _linkanchoredtext := range link.TextAtArrowEnd {
			setPointerField = SliceOfPointersFieldInitStatement
			setPointerField = strings.ReplaceAll(setPointerField, "{{Identifier}}", link.GongGetIdentifier(stage))
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldName}}", "TextAtArrowEnd")
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldNameValue}}", _linkanchoredtext.GongGetIdentifier(stage))
			pointersInitializesStatements += setPointerField
		}
	case "ControlPoints":
		for _, _controlpoint := range link.ControlPoints {
			setPointerField = SliceOfPointersFieldInitStatement
			setPointerField = strings.ReplaceAll(setPointerField, "{{Identifier}}", link.GongGetIdentifier(stage))
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldName}}", "ControlPoints")
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldNameValue}}", _controlpoint.GongGetIdentifier(stage))
			pointersInitializesStatements += setPointerField
		}
	default:
		log.Panicf("Unknown field %s for Gongstruct Link", fieldName)
	}

	// temporary kludge to reuse existing template code
	res = initializerStatements + pointersInitializesStatements
	return
}

func (linkanchoredtext *LinkAnchoredText) GongMarshallField(stage *Stage, fieldName string) (res string) {
	var setValueField, setPointerField string
	_ = setValueField
	_ = setPointerField
	initializerStatements := ""
	_ = initializerStatements
	pointersInitializesStatements := ""
	_ = pointersInitializesStatements

	switch fieldName {
	case "Name":
		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", linkanchoredtext.GongGetIdentifier(stage))
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Name")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(linkanchoredtext.Name))
		initializerStatements += setValueField
	case "Content":
		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", linkanchoredtext.GongGetIdentifier(stage))
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Content")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(linkanchoredtext.Content))
		initializerStatements += setValueField
	case "AutomaticLayout":
		setValueField = NumberInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", linkanchoredtext.GongGetIdentifier(stage))
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "AutomaticLayout")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%t", linkanchoredtext.AutomaticLayout))
		initializerStatements += setValueField
	case "LinkAnchorType":
		if linkanchoredtext.LinkAnchorType != "" {
			setValueField = StringEnumInitStatement
			setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", linkanchoredtext.GongGetIdentifier(stage))
			setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "LinkAnchorType")
			setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", "models."+linkanchoredtext.LinkAnchorType.ToCodeString())
			initializerStatements += setValueField
		}
	case "X_Offset":
		setValueField = NumberInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", linkanchoredtext.GongGetIdentifier(stage))
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "X_Offset")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%f", linkanchoredtext.X_Offset))
		initializerStatements += setValueField
	case "Y_Offset":
		setValueField = NumberInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", linkanchoredtext.GongGetIdentifier(stage))
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Y_Offset")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%f", linkanchoredtext.Y_Offset))
		initializerStatements += setValueField
	case "FontWeight":
		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", linkanchoredtext.GongGetIdentifier(stage))
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "FontWeight")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(linkanchoredtext.FontWeight))
		initializerStatements += setValueField
	case "FontSize":
		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", linkanchoredtext.GongGetIdentifier(stage))
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "FontSize")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(linkanchoredtext.FontSize))
		initializerStatements += setValueField
	case "FontStyle":
		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", linkanchoredtext.GongGetIdentifier(stage))
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "FontStyle")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(linkanchoredtext.FontStyle))
		initializerStatements += setValueField
	case "LetterSpacing":
		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", linkanchoredtext.GongGetIdentifier(stage))
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "LetterSpacing")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(linkanchoredtext.LetterSpacing))
		initializerStatements += setValueField
	case "FontFamily":
		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", linkanchoredtext.GongGetIdentifier(stage))
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "FontFamily")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(linkanchoredtext.FontFamily))
		initializerStatements += setValueField
	case "WhiteSpace":
		if linkanchoredtext.WhiteSpace != "" {
			setValueField = StringEnumInitStatement
			setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", linkanchoredtext.GongGetIdentifier(stage))
			setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "WhiteSpace")
			setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", "models."+linkanchoredtext.WhiteSpace.ToCodeString())
			initializerStatements += setValueField
		}
	case "Color":
		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", linkanchoredtext.GongGetIdentifier(stage))
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Color")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(linkanchoredtext.Color))
		initializerStatements += setValueField
	case "FillOpacity":
		setValueField = NumberInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", linkanchoredtext.GongGetIdentifier(stage))
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "FillOpacity")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%f", linkanchoredtext.FillOpacity))
		initializerStatements += setValueField
	case "Stroke":
		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", linkanchoredtext.GongGetIdentifier(stage))
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Stroke")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(linkanchoredtext.Stroke))
		initializerStatements += setValueField
	case "StrokeOpacity":
		setValueField = NumberInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", linkanchoredtext.GongGetIdentifier(stage))
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "StrokeOpacity")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%f", linkanchoredtext.StrokeOpacity))
		initializerStatements += setValueField
	case "StrokeWidth":
		setValueField = NumberInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", linkanchoredtext.GongGetIdentifier(stage))
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "StrokeWidth")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%f", linkanchoredtext.StrokeWidth))
		initializerStatements += setValueField
	case "StrokeDashArray":
		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", linkanchoredtext.GongGetIdentifier(stage))
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "StrokeDashArray")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(linkanchoredtext.StrokeDashArray))
		initializerStatements += setValueField
	case "StrokeDashArrayWhenSelected":
		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", linkanchoredtext.GongGetIdentifier(stage))
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "StrokeDashArrayWhenSelected")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(linkanchoredtext.StrokeDashArrayWhenSelected))
		initializerStatements += setValueField
	case "Transform":
		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", linkanchoredtext.GongGetIdentifier(stage))
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Transform")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(linkanchoredtext.Transform))
		initializerStatements += setValueField

	case "Animates":
		for _, _animate := range linkanchoredtext.Animates {
			setPointerField = SliceOfPointersFieldInitStatement
			setPointerField = strings.ReplaceAll(setPointerField, "{{Identifier}}", linkanchoredtext.GongGetIdentifier(stage))
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldName}}", "Animates")
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldNameValue}}", _animate.GongGetIdentifier(stage))
			pointersInitializesStatements += setPointerField
		}
	default:
		log.Panicf("Unknown field %s for Gongstruct LinkAnchoredText", fieldName)
	}

	// temporary kludge to reuse existing template code
	res = initializerStatements + pointersInitializesStatements
	return
}

func (path *Path) GongMarshallField(stage *Stage, fieldName string) (res string) {
	var setValueField, setPointerField string
	_ = setValueField
	_ = setPointerField
	initializerStatements := ""
	_ = initializerStatements
	pointersInitializesStatements := ""
	_ = pointersInitializesStatements

	switch fieldName {
	case "Name":
		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", path.GongGetIdentifier(stage))
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Name")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(path.Name))
		initializerStatements += setValueField
	case "Definition":
		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", path.GongGetIdentifier(stage))
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Definition")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(path.Definition))
		initializerStatements += setValueField
	case "Color":
		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", path.GongGetIdentifier(stage))
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Color")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(path.Color))
		initializerStatements += setValueField
	case "FillOpacity":
		setValueField = NumberInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", path.GongGetIdentifier(stage))
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "FillOpacity")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%f", path.FillOpacity))
		initializerStatements += setValueField
	case "Stroke":
		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", path.GongGetIdentifier(stage))
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Stroke")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(path.Stroke))
		initializerStatements += setValueField
	case "StrokeOpacity":
		setValueField = NumberInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", path.GongGetIdentifier(stage))
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "StrokeOpacity")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%f", path.StrokeOpacity))
		initializerStatements += setValueField
	case "StrokeWidth":
		setValueField = NumberInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", path.GongGetIdentifier(stage))
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "StrokeWidth")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%f", path.StrokeWidth))
		initializerStatements += setValueField
	case "StrokeDashArray":
		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", path.GongGetIdentifier(stage))
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "StrokeDashArray")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(path.StrokeDashArray))
		initializerStatements += setValueField
	case "StrokeDashArrayWhenSelected":
		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", path.GongGetIdentifier(stage))
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "StrokeDashArrayWhenSelected")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(path.StrokeDashArrayWhenSelected))
		initializerStatements += setValueField
	case "Transform":
		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", path.GongGetIdentifier(stage))
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Transform")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(path.Transform))
		initializerStatements += setValueField

	case "Animates":
		for _, _animate := range path.Animates {
			setPointerField = SliceOfPointersFieldInitStatement
			setPointerField = strings.ReplaceAll(setPointerField, "{{Identifier}}", path.GongGetIdentifier(stage))
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldName}}", "Animates")
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldNameValue}}", _animate.GongGetIdentifier(stage))
			pointersInitializesStatements += setPointerField
		}
	default:
		log.Panicf("Unknown field %s for Gongstruct Path", fieldName)
	}

	// temporary kludge to reuse existing template code
	res = initializerStatements + pointersInitializesStatements
	return
}

func (point *Point) GongMarshallField(stage *Stage, fieldName string) (res string) {
	var setValueField, setPointerField string
	_ = setValueField
	_ = setPointerField
	initializerStatements := ""
	_ = initializerStatements
	pointersInitializesStatements := ""
	_ = pointersInitializesStatements

	switch fieldName {
	case "Name":
		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", point.GongGetIdentifier(stage))
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Name")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(point.Name))
		initializerStatements += setValueField
	case "X":
		setValueField = NumberInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", point.GongGetIdentifier(stage))
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "X")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%f", point.X))
		initializerStatements += setValueField
	case "Y":
		setValueField = NumberInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", point.GongGetIdentifier(stage))
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Y")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%f", point.Y))
		initializerStatements += setValueField

	default:
		log.Panicf("Unknown field %s for Gongstruct Point", fieldName)
	}

	// temporary kludge to reuse existing template code
	res = initializerStatements + pointersInitializesStatements
	return
}

func (polygone *Polygone) GongMarshallField(stage *Stage, fieldName string) (res string) {
	var setValueField, setPointerField string
	_ = setValueField
	_ = setPointerField
	initializerStatements := ""
	_ = initializerStatements
	pointersInitializesStatements := ""
	_ = pointersInitializesStatements

	switch fieldName {
	case "Name":
		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", polygone.GongGetIdentifier(stage))
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Name")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(polygone.Name))
		initializerStatements += setValueField
	case "Points":
		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", polygone.GongGetIdentifier(stage))
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Points")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(polygone.Points))
		initializerStatements += setValueField
	case "Color":
		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", polygone.GongGetIdentifier(stage))
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Color")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(polygone.Color))
		initializerStatements += setValueField
	case "FillOpacity":
		setValueField = NumberInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", polygone.GongGetIdentifier(stage))
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "FillOpacity")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%f", polygone.FillOpacity))
		initializerStatements += setValueField
	case "Stroke":
		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", polygone.GongGetIdentifier(stage))
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Stroke")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(polygone.Stroke))
		initializerStatements += setValueField
	case "StrokeOpacity":
		setValueField = NumberInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", polygone.GongGetIdentifier(stage))
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "StrokeOpacity")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%f", polygone.StrokeOpacity))
		initializerStatements += setValueField
	case "StrokeWidth":
		setValueField = NumberInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", polygone.GongGetIdentifier(stage))
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "StrokeWidth")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%f", polygone.StrokeWidth))
		initializerStatements += setValueField
	case "StrokeDashArray":
		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", polygone.GongGetIdentifier(stage))
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "StrokeDashArray")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(polygone.StrokeDashArray))
		initializerStatements += setValueField
	case "StrokeDashArrayWhenSelected":
		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", polygone.GongGetIdentifier(stage))
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "StrokeDashArrayWhenSelected")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(polygone.StrokeDashArrayWhenSelected))
		initializerStatements += setValueField
	case "Transform":
		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", polygone.GongGetIdentifier(stage))
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Transform")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(polygone.Transform))
		initializerStatements += setValueField

	case "Animates":
		for _, _animate := range polygone.Animates {
			setPointerField = SliceOfPointersFieldInitStatement
			setPointerField = strings.ReplaceAll(setPointerField, "{{Identifier}}", polygone.GongGetIdentifier(stage))
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldName}}", "Animates")
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldNameValue}}", _animate.GongGetIdentifier(stage))
			pointersInitializesStatements += setPointerField
		}
	default:
		log.Panicf("Unknown field %s for Gongstruct Polygone", fieldName)
	}

	// temporary kludge to reuse existing template code
	res = initializerStatements + pointersInitializesStatements
	return
}

func (polyline *Polyline) GongMarshallField(stage *Stage, fieldName string) (res string) {
	var setValueField, setPointerField string
	_ = setValueField
	_ = setPointerField
	initializerStatements := ""
	_ = initializerStatements
	pointersInitializesStatements := ""
	_ = pointersInitializesStatements

	switch fieldName {
	case "Name":
		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", polyline.GongGetIdentifier(stage))
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Name")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(polyline.Name))
		initializerStatements += setValueField
	case "Points":
		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", polyline.GongGetIdentifier(stage))
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Points")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(polyline.Points))
		initializerStatements += setValueField
	case "Color":
		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", polyline.GongGetIdentifier(stage))
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Color")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(polyline.Color))
		initializerStatements += setValueField
	case "FillOpacity":
		setValueField = NumberInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", polyline.GongGetIdentifier(stage))
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "FillOpacity")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%f", polyline.FillOpacity))
		initializerStatements += setValueField
	case "Stroke":
		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", polyline.GongGetIdentifier(stage))
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Stroke")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(polyline.Stroke))
		initializerStatements += setValueField
	case "StrokeOpacity":
		setValueField = NumberInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", polyline.GongGetIdentifier(stage))
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "StrokeOpacity")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%f", polyline.StrokeOpacity))
		initializerStatements += setValueField
	case "StrokeWidth":
		setValueField = NumberInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", polyline.GongGetIdentifier(stage))
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "StrokeWidth")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%f", polyline.StrokeWidth))
		initializerStatements += setValueField
	case "StrokeDashArray":
		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", polyline.GongGetIdentifier(stage))
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "StrokeDashArray")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(polyline.StrokeDashArray))
		initializerStatements += setValueField
	case "StrokeDashArrayWhenSelected":
		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", polyline.GongGetIdentifier(stage))
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "StrokeDashArrayWhenSelected")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(polyline.StrokeDashArrayWhenSelected))
		initializerStatements += setValueField
	case "Transform":
		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", polyline.GongGetIdentifier(stage))
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Transform")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(polyline.Transform))
		initializerStatements += setValueField

	case "Animates":
		for _, _animate := range polyline.Animates {
			setPointerField = SliceOfPointersFieldInitStatement
			setPointerField = strings.ReplaceAll(setPointerField, "{{Identifier}}", polyline.GongGetIdentifier(stage))
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldName}}", "Animates")
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldNameValue}}", _animate.GongGetIdentifier(stage))
			pointersInitializesStatements += setPointerField
		}
	default:
		log.Panicf("Unknown field %s for Gongstruct Polyline", fieldName)
	}

	// temporary kludge to reuse existing template code
	res = initializerStatements + pointersInitializesStatements
	return
}

func (rect *Rect) GongMarshallField(stage *Stage, fieldName string) (res string) {
	var setValueField, setPointerField string
	_ = setValueField
	_ = setPointerField
	initializerStatements := ""
	_ = initializerStatements
	pointersInitializesStatements := ""
	_ = pointersInitializesStatements

	switch fieldName {
	case "Name":
		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", rect.GongGetIdentifier(stage))
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Name")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(rect.Name))
		initializerStatements += setValueField
	case "X":
		setValueField = NumberInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", rect.GongGetIdentifier(stage))
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "X")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%f", rect.X))
		initializerStatements += setValueField
	case "Y":
		setValueField = NumberInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", rect.GongGetIdentifier(stage))
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Y")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%f", rect.Y))
		initializerStatements += setValueField
	case "Width":
		setValueField = NumberInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", rect.GongGetIdentifier(stage))
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Width")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%f", rect.Width))
		initializerStatements += setValueField
	case "Height":
		setValueField = NumberInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", rect.GongGetIdentifier(stage))
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Height")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%f", rect.Height))
		initializerStatements += setValueField
	case "RX":
		setValueField = NumberInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", rect.GongGetIdentifier(stage))
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "RX")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%f", rect.RX))
		initializerStatements += setValueField
	case "Color":
		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", rect.GongGetIdentifier(stage))
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Color")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(rect.Color))
		initializerStatements += setValueField
	case "FillOpacity":
		setValueField = NumberInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", rect.GongGetIdentifier(stage))
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "FillOpacity")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%f", rect.FillOpacity))
		initializerStatements += setValueField
	case "Stroke":
		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", rect.GongGetIdentifier(stage))
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Stroke")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(rect.Stroke))
		initializerStatements += setValueField
	case "StrokeOpacity":
		setValueField = NumberInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", rect.GongGetIdentifier(stage))
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "StrokeOpacity")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%f", rect.StrokeOpacity))
		initializerStatements += setValueField
	case "StrokeWidth":
		setValueField = NumberInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", rect.GongGetIdentifier(stage))
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "StrokeWidth")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%f", rect.StrokeWidth))
		initializerStatements += setValueField
	case "StrokeDashArray":
		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", rect.GongGetIdentifier(stage))
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "StrokeDashArray")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(rect.StrokeDashArray))
		initializerStatements += setValueField
	case "StrokeDashArrayWhenSelected":
		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", rect.GongGetIdentifier(stage))
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "StrokeDashArrayWhenSelected")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(rect.StrokeDashArrayWhenSelected))
		initializerStatements += setValueField
	case "Transform":
		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", rect.GongGetIdentifier(stage))
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Transform")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(rect.Transform))
		initializerStatements += setValueField
	case "IsSelectable":
		setValueField = NumberInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", rect.GongGetIdentifier(stage))
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "IsSelectable")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%t", rect.IsSelectable))
		initializerStatements += setValueField
	case "IsSelected":
		setValueField = NumberInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", rect.GongGetIdentifier(stage))
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "IsSelected")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%t", rect.IsSelected))
		initializerStatements += setValueField
	case "CanHaveLeftHandle":
		setValueField = NumberInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", rect.GongGetIdentifier(stage))
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "CanHaveLeftHandle")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%t", rect.CanHaveLeftHandle))
		initializerStatements += setValueField
	case "HasLeftHandle":
		setValueField = NumberInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", rect.GongGetIdentifier(stage))
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "HasLeftHandle")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%t", rect.HasLeftHandle))
		initializerStatements += setValueField
	case "CanHaveRightHandle":
		setValueField = NumberInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", rect.GongGetIdentifier(stage))
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "CanHaveRightHandle")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%t", rect.CanHaveRightHandle))
		initializerStatements += setValueField
	case "HasRightHandle":
		setValueField = NumberInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", rect.GongGetIdentifier(stage))
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "HasRightHandle")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%t", rect.HasRightHandle))
		initializerStatements += setValueField
	case "CanHaveTopHandle":
		setValueField = NumberInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", rect.GongGetIdentifier(stage))
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "CanHaveTopHandle")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%t", rect.CanHaveTopHandle))
		initializerStatements += setValueField
	case "HasTopHandle":
		setValueField = NumberInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", rect.GongGetIdentifier(stage))
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "HasTopHandle")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%t", rect.HasTopHandle))
		initializerStatements += setValueField
	case "IsScalingProportionally":
		setValueField = NumberInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", rect.GongGetIdentifier(stage))
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "IsScalingProportionally")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%t", rect.IsScalingProportionally))
		initializerStatements += setValueField
	case "CanHaveBottomHandle":
		setValueField = NumberInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", rect.GongGetIdentifier(stage))
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "CanHaveBottomHandle")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%t", rect.CanHaveBottomHandle))
		initializerStatements += setValueField
	case "HasBottomHandle":
		setValueField = NumberInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", rect.GongGetIdentifier(stage))
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "HasBottomHandle")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%t", rect.HasBottomHandle))
		initializerStatements += setValueField
	case "CanMoveHorizontaly":
		setValueField = NumberInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", rect.GongGetIdentifier(stage))
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "CanMoveHorizontaly")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%t", rect.CanMoveHorizontaly))
		initializerStatements += setValueField
	case "CanMoveVerticaly":
		setValueField = NumberInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", rect.GongGetIdentifier(stage))
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "CanMoveVerticaly")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%t", rect.CanMoveVerticaly))
		initializerStatements += setValueField
	case "ChangeColorWhenHovered":
		setValueField = NumberInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", rect.GongGetIdentifier(stage))
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "ChangeColorWhenHovered")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%t", rect.ChangeColorWhenHovered))
		initializerStatements += setValueField
	case "ColorWhenHovered":
		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", rect.GongGetIdentifier(stage))
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "ColorWhenHovered")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(rect.ColorWhenHovered))
		initializerStatements += setValueField
	case "OriginalColor":
		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", rect.GongGetIdentifier(stage))
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "OriginalColor")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(rect.OriginalColor))
		initializerStatements += setValueField
	case "FillOpacityWhenHovered":
		setValueField = NumberInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", rect.GongGetIdentifier(stage))
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "FillOpacityWhenHovered")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%f", rect.FillOpacityWhenHovered))
		initializerStatements += setValueField
	case "OriginalFillOpacity":
		setValueField = NumberInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", rect.GongGetIdentifier(stage))
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "OriginalFillOpacity")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%f", rect.OriginalFillOpacity))
		initializerStatements += setValueField
	case "HasToolTip":
		setValueField = NumberInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", rect.GongGetIdentifier(stage))
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "HasToolTip")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%t", rect.HasToolTip))
		initializerStatements += setValueField
	case "ToolTipText":
		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", rect.GongGetIdentifier(stage))
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "ToolTipText")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(rect.ToolTipText))
		initializerStatements += setValueField
	case "ToolTipPosition":
		if rect.ToolTipPosition != "" {
			setValueField = StringEnumInitStatement
			setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", rect.GongGetIdentifier(stage))
			setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "ToolTipPosition")
			setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", "models."+rect.ToolTipPosition.ToCodeString())
			initializerStatements += setValueField
		}
	case "MouseX":
		setValueField = NumberInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", rect.GongGetIdentifier(stage))
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "MouseX")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%f", rect.MouseX))
		initializerStatements += setValueField
	case "MouseY":
		setValueField = NumberInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", rect.GongGetIdentifier(stage))
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "MouseY")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%f", rect.MouseY))
		initializerStatements += setValueField
	case "MouseEventKey":
		if rect.MouseEventKey != "" {
			setValueField = StringEnumInitStatement
			setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", rect.GongGetIdentifier(stage))
			setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "MouseEventKey")
			setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", "models."+rect.MouseEventKey.ToCodeString())
			initializerStatements += setValueField
		}

	case "HoveringTrigger":
		for _, _condition := range rect.HoveringTrigger {
			setPointerField = SliceOfPointersFieldInitStatement
			setPointerField = strings.ReplaceAll(setPointerField, "{{Identifier}}", rect.GongGetIdentifier(stage))
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldName}}", "HoveringTrigger")
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldNameValue}}", _condition.GongGetIdentifier(stage))
			pointersInitializesStatements += setPointerField
		}
	case "DisplayConditions":
		for _, _condition := range rect.DisplayConditions {
			setPointerField = SliceOfPointersFieldInitStatement
			setPointerField = strings.ReplaceAll(setPointerField, "{{Identifier}}", rect.GongGetIdentifier(stage))
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldName}}", "DisplayConditions")
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldNameValue}}", _condition.GongGetIdentifier(stage))
			pointersInitializesStatements += setPointerField
		}
	case "Animations":
		for _, _animate := range rect.Animations {
			setPointerField = SliceOfPointersFieldInitStatement
			setPointerField = strings.ReplaceAll(setPointerField, "{{Identifier}}", rect.GongGetIdentifier(stage))
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldName}}", "Animations")
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldNameValue}}", _animate.GongGetIdentifier(stage))
			pointersInitializesStatements += setPointerField
		}
	case "RectAnchoredTexts":
		for _, _rectanchoredtext := range rect.RectAnchoredTexts {
			setPointerField = SliceOfPointersFieldInitStatement
			setPointerField = strings.ReplaceAll(setPointerField, "{{Identifier}}", rect.GongGetIdentifier(stage))
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldName}}", "RectAnchoredTexts")
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldNameValue}}", _rectanchoredtext.GongGetIdentifier(stage))
			pointersInitializesStatements += setPointerField
		}
	case "RectAnchoredRects":
		for _, _rectanchoredrect := range rect.RectAnchoredRects {
			setPointerField = SliceOfPointersFieldInitStatement
			setPointerField = strings.ReplaceAll(setPointerField, "{{Identifier}}", rect.GongGetIdentifier(stage))
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldName}}", "RectAnchoredRects")
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldNameValue}}", _rectanchoredrect.GongGetIdentifier(stage))
			pointersInitializesStatements += setPointerField
		}
	case "RectAnchoredPaths":
		for _, _rectanchoredpath := range rect.RectAnchoredPaths {
			setPointerField = SliceOfPointersFieldInitStatement
			setPointerField = strings.ReplaceAll(setPointerField, "{{Identifier}}", rect.GongGetIdentifier(stage))
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldName}}", "RectAnchoredPaths")
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldNameValue}}", _rectanchoredpath.GongGetIdentifier(stage))
			pointersInitializesStatements += setPointerField
		}
	default:
		log.Panicf("Unknown field %s for Gongstruct Rect", fieldName)
	}

	// temporary kludge to reuse existing template code
	res = initializerStatements + pointersInitializesStatements
	return
}

func (rectanchoredpath *RectAnchoredPath) GongMarshallField(stage *Stage, fieldName string) (res string) {
	var setValueField, setPointerField string
	_ = setValueField
	_ = setPointerField
	initializerStatements := ""
	_ = initializerStatements
	pointersInitializesStatements := ""
	_ = pointersInitializesStatements

	switch fieldName {
	case "Name":
		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", rectanchoredpath.GongGetIdentifier(stage))
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Name")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(rectanchoredpath.Name))
		initializerStatements += setValueField
	case "Definition":
		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", rectanchoredpath.GongGetIdentifier(stage))
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Definition")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(rectanchoredpath.Definition))
		initializerStatements += setValueField
	case "X_Offset":
		setValueField = NumberInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", rectanchoredpath.GongGetIdentifier(stage))
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "X_Offset")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%f", rectanchoredpath.X_Offset))
		initializerStatements += setValueField
	case "Y_Offset":
		setValueField = NumberInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", rectanchoredpath.GongGetIdentifier(stage))
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Y_Offset")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%f", rectanchoredpath.Y_Offset))
		initializerStatements += setValueField
	case "RectAnchorType":
		if rectanchoredpath.RectAnchorType != "" {
			setValueField = StringEnumInitStatement
			setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", rectanchoredpath.GongGetIdentifier(stage))
			setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "RectAnchorType")
			setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", "models."+rectanchoredpath.RectAnchorType.ToCodeString())
			initializerStatements += setValueField
		}
	case "ScalePropotionnally":
		setValueField = NumberInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", rectanchoredpath.GongGetIdentifier(stage))
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "ScalePropotionnally")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%t", rectanchoredpath.ScalePropotionnally))
		initializerStatements += setValueField
	case "AppliedScaling":
		setValueField = NumberInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", rectanchoredpath.GongGetIdentifier(stage))
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "AppliedScaling")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%f", rectanchoredpath.AppliedScaling))
		initializerStatements += setValueField
	case "Color":
		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", rectanchoredpath.GongGetIdentifier(stage))
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Color")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(rectanchoredpath.Color))
		initializerStatements += setValueField
	case "FillOpacity":
		setValueField = NumberInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", rectanchoredpath.GongGetIdentifier(stage))
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "FillOpacity")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%f", rectanchoredpath.FillOpacity))
		initializerStatements += setValueField
	case "Stroke":
		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", rectanchoredpath.GongGetIdentifier(stage))
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Stroke")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(rectanchoredpath.Stroke))
		initializerStatements += setValueField
	case "StrokeOpacity":
		setValueField = NumberInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", rectanchoredpath.GongGetIdentifier(stage))
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "StrokeOpacity")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%f", rectanchoredpath.StrokeOpacity))
		initializerStatements += setValueField
	case "StrokeWidth":
		setValueField = NumberInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", rectanchoredpath.GongGetIdentifier(stage))
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "StrokeWidth")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%f", rectanchoredpath.StrokeWidth))
		initializerStatements += setValueField
	case "StrokeDashArray":
		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", rectanchoredpath.GongGetIdentifier(stage))
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "StrokeDashArray")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(rectanchoredpath.StrokeDashArray))
		initializerStatements += setValueField
	case "StrokeDashArrayWhenSelected":
		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", rectanchoredpath.GongGetIdentifier(stage))
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "StrokeDashArrayWhenSelected")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(rectanchoredpath.StrokeDashArrayWhenSelected))
		initializerStatements += setValueField
	case "Transform":
		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", rectanchoredpath.GongGetIdentifier(stage))
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Transform")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(rectanchoredpath.Transform))
		initializerStatements += setValueField

	default:
		log.Panicf("Unknown field %s for Gongstruct RectAnchoredPath", fieldName)
	}

	// temporary kludge to reuse existing template code
	res = initializerStatements + pointersInitializesStatements
	return
}

func (rectanchoredrect *RectAnchoredRect) GongMarshallField(stage *Stage, fieldName string) (res string) {
	var setValueField, setPointerField string
	_ = setValueField
	_ = setPointerField
	initializerStatements := ""
	_ = initializerStatements
	pointersInitializesStatements := ""
	_ = pointersInitializesStatements

	switch fieldName {
	case "Name":
		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", rectanchoredrect.GongGetIdentifier(stage))
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Name")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(rectanchoredrect.Name))
		initializerStatements += setValueField
	case "X":
		setValueField = NumberInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", rectanchoredrect.GongGetIdentifier(stage))
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "X")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%f", rectanchoredrect.X))
		initializerStatements += setValueField
	case "Y":
		setValueField = NumberInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", rectanchoredrect.GongGetIdentifier(stage))
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Y")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%f", rectanchoredrect.Y))
		initializerStatements += setValueField
	case "Width":
		setValueField = NumberInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", rectanchoredrect.GongGetIdentifier(stage))
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Width")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%f", rectanchoredrect.Width))
		initializerStatements += setValueField
	case "Height":
		setValueField = NumberInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", rectanchoredrect.GongGetIdentifier(stage))
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Height")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%f", rectanchoredrect.Height))
		initializerStatements += setValueField
	case "RX":
		setValueField = NumberInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", rectanchoredrect.GongGetIdentifier(stage))
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "RX")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%f", rectanchoredrect.RX))
		initializerStatements += setValueField
	case "X_Offset":
		setValueField = NumberInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", rectanchoredrect.GongGetIdentifier(stage))
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "X_Offset")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%f", rectanchoredrect.X_Offset))
		initializerStatements += setValueField
	case "Y_Offset":
		setValueField = NumberInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", rectanchoredrect.GongGetIdentifier(stage))
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Y_Offset")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%f", rectanchoredrect.Y_Offset))
		initializerStatements += setValueField
	case "RectAnchorType":
		if rectanchoredrect.RectAnchorType != "" {
			setValueField = StringEnumInitStatement
			setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", rectanchoredrect.GongGetIdentifier(stage))
			setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "RectAnchorType")
			setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", "models."+rectanchoredrect.RectAnchorType.ToCodeString())
			initializerStatements += setValueField
		}
	case "WidthFollowRect":
		setValueField = NumberInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", rectanchoredrect.GongGetIdentifier(stage))
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "WidthFollowRect")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%t", rectanchoredrect.WidthFollowRect))
		initializerStatements += setValueField
	case "HeightFollowRect":
		setValueField = NumberInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", rectanchoredrect.GongGetIdentifier(stage))
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "HeightFollowRect")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%t", rectanchoredrect.HeightFollowRect))
		initializerStatements += setValueField
	case "HasToolTip":
		setValueField = NumberInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", rectanchoredrect.GongGetIdentifier(stage))
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "HasToolTip")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%t", rectanchoredrect.HasToolTip))
		initializerStatements += setValueField
	case "ToolTipText":
		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", rectanchoredrect.GongGetIdentifier(stage))
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "ToolTipText")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(rectanchoredrect.ToolTipText))
		initializerStatements += setValueField
	case "Color":
		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", rectanchoredrect.GongGetIdentifier(stage))
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Color")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(rectanchoredrect.Color))
		initializerStatements += setValueField
	case "FillOpacity":
		setValueField = NumberInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", rectanchoredrect.GongGetIdentifier(stage))
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "FillOpacity")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%f", rectanchoredrect.FillOpacity))
		initializerStatements += setValueField
	case "Stroke":
		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", rectanchoredrect.GongGetIdentifier(stage))
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Stroke")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(rectanchoredrect.Stroke))
		initializerStatements += setValueField
	case "StrokeOpacity":
		setValueField = NumberInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", rectanchoredrect.GongGetIdentifier(stage))
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "StrokeOpacity")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%f", rectanchoredrect.StrokeOpacity))
		initializerStatements += setValueField
	case "StrokeWidth":
		setValueField = NumberInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", rectanchoredrect.GongGetIdentifier(stage))
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "StrokeWidth")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%f", rectanchoredrect.StrokeWidth))
		initializerStatements += setValueField
	case "StrokeDashArray":
		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", rectanchoredrect.GongGetIdentifier(stage))
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "StrokeDashArray")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(rectanchoredrect.StrokeDashArray))
		initializerStatements += setValueField
	case "StrokeDashArrayWhenSelected":
		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", rectanchoredrect.GongGetIdentifier(stage))
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "StrokeDashArrayWhenSelected")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(rectanchoredrect.StrokeDashArrayWhenSelected))
		initializerStatements += setValueField
	case "Transform":
		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", rectanchoredrect.GongGetIdentifier(stage))
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Transform")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(rectanchoredrect.Transform))
		initializerStatements += setValueField

	default:
		log.Panicf("Unknown field %s for Gongstruct RectAnchoredRect", fieldName)
	}

	// temporary kludge to reuse existing template code
	res = initializerStatements + pointersInitializesStatements
	return
}

func (rectanchoredtext *RectAnchoredText) GongMarshallField(stage *Stage, fieldName string) (res string) {
	var setValueField, setPointerField string
	_ = setValueField
	_ = setPointerField
	initializerStatements := ""
	_ = initializerStatements
	pointersInitializesStatements := ""
	_ = pointersInitializesStatements

	switch fieldName {
	case "Name":
		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", rectanchoredtext.GongGetIdentifier(stage))
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Name")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(rectanchoredtext.Name))
		initializerStatements += setValueField
	case "Content":
		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", rectanchoredtext.GongGetIdentifier(stage))
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Content")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(rectanchoredtext.Content))
		initializerStatements += setValueField
	case "FontWeight":
		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", rectanchoredtext.GongGetIdentifier(stage))
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "FontWeight")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(rectanchoredtext.FontWeight))
		initializerStatements += setValueField
	case "FontSize":
		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", rectanchoredtext.GongGetIdentifier(stage))
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "FontSize")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(rectanchoredtext.FontSize))
		initializerStatements += setValueField
	case "FontStyle":
		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", rectanchoredtext.GongGetIdentifier(stage))
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "FontStyle")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(rectanchoredtext.FontStyle))
		initializerStatements += setValueField
	case "LetterSpacing":
		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", rectanchoredtext.GongGetIdentifier(stage))
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "LetterSpacing")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(rectanchoredtext.LetterSpacing))
		initializerStatements += setValueField
	case "FontFamily":
		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", rectanchoredtext.GongGetIdentifier(stage))
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "FontFamily")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(rectanchoredtext.FontFamily))
		initializerStatements += setValueField
	case "WhiteSpace":
		if rectanchoredtext.WhiteSpace != "" {
			setValueField = StringEnumInitStatement
			setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", rectanchoredtext.GongGetIdentifier(stage))
			setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "WhiteSpace")
			setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", "models."+rectanchoredtext.WhiteSpace.ToCodeString())
			initializerStatements += setValueField
		}
	case "X_Offset":
		setValueField = NumberInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", rectanchoredtext.GongGetIdentifier(stage))
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "X_Offset")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%f", rectanchoredtext.X_Offset))
		initializerStatements += setValueField
	case "Y_Offset":
		setValueField = NumberInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", rectanchoredtext.GongGetIdentifier(stage))
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Y_Offset")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%f", rectanchoredtext.Y_Offset))
		initializerStatements += setValueField
	case "RectAnchorType":
		if rectanchoredtext.RectAnchorType != "" {
			setValueField = StringEnumInitStatement
			setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", rectanchoredtext.GongGetIdentifier(stage))
			setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "RectAnchorType")
			setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", "models."+rectanchoredtext.RectAnchorType.ToCodeString())
			initializerStatements += setValueField
		}
	case "TextAnchorType":
		if rectanchoredtext.TextAnchorType != "" {
			setValueField = StringEnumInitStatement
			setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", rectanchoredtext.GongGetIdentifier(stage))
			setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "TextAnchorType")
			setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", "models."+rectanchoredtext.TextAnchorType.ToCodeString())
			initializerStatements += setValueField
		}
	case "DominantBaseline":
		if rectanchoredtext.DominantBaseline != "" {
			setValueField = StringEnumInitStatement
			setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", rectanchoredtext.GongGetIdentifier(stage))
			setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "DominantBaseline")
			setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", "models."+rectanchoredtext.DominantBaseline.ToCodeString())
			initializerStatements += setValueField
		}
	case "WritingMode":
		if rectanchoredtext.WritingMode != "" {
			setValueField = StringEnumInitStatement
			setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", rectanchoredtext.GongGetIdentifier(stage))
			setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "WritingMode")
			setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", "models."+rectanchoredtext.WritingMode.ToCodeString())
			initializerStatements += setValueField
		}
	case "Color":
		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", rectanchoredtext.GongGetIdentifier(stage))
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Color")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(rectanchoredtext.Color))
		initializerStatements += setValueField
	case "FillOpacity":
		setValueField = NumberInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", rectanchoredtext.GongGetIdentifier(stage))
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "FillOpacity")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%f", rectanchoredtext.FillOpacity))
		initializerStatements += setValueField
	case "Stroke":
		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", rectanchoredtext.GongGetIdentifier(stage))
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Stroke")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(rectanchoredtext.Stroke))
		initializerStatements += setValueField
	case "StrokeOpacity":
		setValueField = NumberInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", rectanchoredtext.GongGetIdentifier(stage))
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "StrokeOpacity")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%f", rectanchoredtext.StrokeOpacity))
		initializerStatements += setValueField
	case "StrokeWidth":
		setValueField = NumberInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", rectanchoredtext.GongGetIdentifier(stage))
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "StrokeWidth")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%f", rectanchoredtext.StrokeWidth))
		initializerStatements += setValueField
	case "StrokeDashArray":
		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", rectanchoredtext.GongGetIdentifier(stage))
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "StrokeDashArray")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(rectanchoredtext.StrokeDashArray))
		initializerStatements += setValueField
	case "StrokeDashArrayWhenSelected":
		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", rectanchoredtext.GongGetIdentifier(stage))
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "StrokeDashArrayWhenSelected")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(rectanchoredtext.StrokeDashArrayWhenSelected))
		initializerStatements += setValueField
	case "Transform":
		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", rectanchoredtext.GongGetIdentifier(stage))
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Transform")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(rectanchoredtext.Transform))
		initializerStatements += setValueField

	case "Animates":
		for _, _animate := range rectanchoredtext.Animates {
			setPointerField = SliceOfPointersFieldInitStatement
			setPointerField = strings.ReplaceAll(setPointerField, "{{Identifier}}", rectanchoredtext.GongGetIdentifier(stage))
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldName}}", "Animates")
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldNameValue}}", _animate.GongGetIdentifier(stage))
			pointersInitializesStatements += setPointerField
		}
	default:
		log.Panicf("Unknown field %s for Gongstruct RectAnchoredText", fieldName)
	}

	// temporary kludge to reuse existing template code
	res = initializerStatements + pointersInitializesStatements
	return
}

func (rectlinklink *RectLinkLink) GongMarshallField(stage *Stage, fieldName string) (res string) {
	var setValueField, setPointerField string
	_ = setValueField
	_ = setPointerField
	initializerStatements := ""
	_ = initializerStatements
	pointersInitializesStatements := ""
	_ = pointersInitializesStatements

	switch fieldName {
	case "Name":
		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", rectlinklink.GongGetIdentifier(stage))
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Name")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(rectlinklink.Name))
		initializerStatements += setValueField
	case "TargetAnchorPosition":
		setValueField = NumberInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", rectlinklink.GongGetIdentifier(stage))
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "TargetAnchorPosition")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%f", rectlinklink.TargetAnchorPosition))
		initializerStatements += setValueField
	case "Color":
		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", rectlinklink.GongGetIdentifier(stage))
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Color")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(rectlinklink.Color))
		initializerStatements += setValueField
	case "FillOpacity":
		setValueField = NumberInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", rectlinklink.GongGetIdentifier(stage))
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "FillOpacity")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%f", rectlinklink.FillOpacity))
		initializerStatements += setValueField
	case "Stroke":
		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", rectlinklink.GongGetIdentifier(stage))
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Stroke")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(rectlinklink.Stroke))
		initializerStatements += setValueField
	case "StrokeOpacity":
		setValueField = NumberInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", rectlinklink.GongGetIdentifier(stage))
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "StrokeOpacity")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%f", rectlinklink.StrokeOpacity))
		initializerStatements += setValueField
	case "StrokeWidth":
		setValueField = NumberInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", rectlinklink.GongGetIdentifier(stage))
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "StrokeWidth")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%f", rectlinklink.StrokeWidth))
		initializerStatements += setValueField
	case "StrokeDashArray":
		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", rectlinklink.GongGetIdentifier(stage))
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "StrokeDashArray")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(rectlinklink.StrokeDashArray))
		initializerStatements += setValueField
	case "StrokeDashArrayWhenSelected":
		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", rectlinklink.GongGetIdentifier(stage))
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "StrokeDashArrayWhenSelected")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(rectlinklink.StrokeDashArrayWhenSelected))
		initializerStatements += setValueField
	case "Transform":
		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", rectlinklink.GongGetIdentifier(stage))
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Transform")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(rectlinklink.Transform))
		initializerStatements += setValueField

	case "Start":
		if rectlinklink.Start != nil {
			setPointerField = PointerFieldInitStatement
			setPointerField = strings.ReplaceAll(setPointerField, "{{Identifier}}", rectlinklink.GongGetIdentifier(stage))
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldName}}", "Start")
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldNameValue}}", rectlinklink.Start.GongGetIdentifier(stage))
			pointersInitializesStatements += setPointerField
		}
	case "End":
		if rectlinklink.End != nil {
			setPointerField = PointerFieldInitStatement
			setPointerField = strings.ReplaceAll(setPointerField, "{{Identifier}}", rectlinklink.GongGetIdentifier(stage))
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldName}}", "End")
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldNameValue}}", rectlinklink.End.GongGetIdentifier(stage))
			pointersInitializesStatements += setPointerField
		}
	default:
		log.Panicf("Unknown field %s for Gongstruct RectLinkLink", fieldName)
	}

	// temporary kludge to reuse existing template code
	res = initializerStatements + pointersInitializesStatements
	return
}

func (svg *SVG) GongMarshallField(stage *Stage, fieldName string) (res string) {
	var setValueField, setPointerField string
	_ = setValueField
	_ = setPointerField
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
	case "DrawingState":
		if svg.DrawingState != "" {
			setValueField = StringEnumInitStatement
			setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", svg.GongGetIdentifier(stage))
			setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "DrawingState")
			setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", "models."+svg.DrawingState.ToCodeString())
			initializerStatements += setValueField
		}
	case "IsEditable":
		setValueField = NumberInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", svg.GongGetIdentifier(stage))
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "IsEditable")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%t", svg.IsEditable))
		initializerStatements += setValueField
	case "IsSVGFrontEndFileGenerated":
		setValueField = NumberInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", svg.GongGetIdentifier(stage))
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "IsSVGFrontEndFileGenerated")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%t", svg.IsSVGFrontEndFileGenerated))
		initializerStatements += setValueField
	case "IsSVGBackEndFileGenerated":
		setValueField = NumberInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", svg.GongGetIdentifier(stage))
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "IsSVGBackEndFileGenerated")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%t", svg.IsSVGBackEndFileGenerated))
		initializerStatements += setValueField
	case "DefaultDirectoryForGeneratedImages":
		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", svg.GongGetIdentifier(stage))
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "DefaultDirectoryForGeneratedImages")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(svg.DefaultDirectoryForGeneratedImages))
		initializerStatements += setValueField
	case "IsControlBannerHidden":
		setValueField = NumberInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", svg.GongGetIdentifier(stage))
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "IsControlBannerHidden")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%t", svg.IsControlBannerHidden))
		initializerStatements += setValueField
	case "OverrideWidth":
		setValueField = NumberInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", svg.GongGetIdentifier(stage))
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "OverrideWidth")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%t", svg.OverrideWidth))
		initializerStatements += setValueField
	case "OverriddenWidth":
		setValueField = NumberInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", svg.GongGetIdentifier(stage))
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "OverriddenWidth")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%f", svg.OverriddenWidth))
		initializerStatements += setValueField
	case "OverrideHeight":
		setValueField = NumberInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", svg.GongGetIdentifier(stage))
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "OverrideHeight")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%t", svg.OverrideHeight))
		initializerStatements += setValueField
	case "OverriddenHeight":
		setValueField = NumberInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", svg.GongGetIdentifier(stage))
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "OverriddenHeight")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%f", svg.OverriddenHeight))
		initializerStatements += setValueField

	case "Layers":
		for _, _layer := range svg.Layers {
			setPointerField = SliceOfPointersFieldInitStatement
			setPointerField = strings.ReplaceAll(setPointerField, "{{Identifier}}", svg.GongGetIdentifier(stage))
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldName}}", "Layers")
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldNameValue}}", _layer.GongGetIdentifier(stage))
			pointersInitializesStatements += setPointerField
		}
	case "StartRect":
		if svg.StartRect != nil {
			setPointerField = PointerFieldInitStatement
			setPointerField = strings.ReplaceAll(setPointerField, "{{Identifier}}", svg.GongGetIdentifier(stage))
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldName}}", "StartRect")
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldNameValue}}", svg.StartRect.GongGetIdentifier(stage))
			pointersInitializesStatements += setPointerField
		}
	case "EndRect":
		if svg.EndRect != nil {
			setPointerField = PointerFieldInitStatement
			setPointerField = strings.ReplaceAll(setPointerField, "{{Identifier}}", svg.GongGetIdentifier(stage))
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldName}}", "EndRect")
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldNameValue}}", svg.EndRect.GongGetIdentifier(stage))
			pointersInitializesStatements += setPointerField
		}
	default:
		log.Panicf("Unknown field %s for Gongstruct SVG", fieldName)
	}

	// temporary kludge to reuse existing template code
	res = initializerStatements + pointersInitializesStatements
	return
}

func (svgtext *SvgText) GongMarshallField(stage *Stage, fieldName string) (res string) {
	var setValueField, setPointerField string
	_ = setValueField
	_ = setPointerField
	initializerStatements := ""
	_ = initializerStatements
	pointersInitializesStatements := ""
	_ = pointersInitializesStatements

	switch fieldName {
	case "Name":
		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", svgtext.GongGetIdentifier(stage))
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Name")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(svgtext.Name))
		initializerStatements += setValueField
	case "Text":
		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", svgtext.GongGetIdentifier(stage))
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Text")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(svgtext.Text))
		initializerStatements += setValueField

	default:
		log.Panicf("Unknown field %s for Gongstruct SvgText", fieldName)
	}

	// temporary kludge to reuse existing template code
	res = initializerStatements + pointersInitializesStatements
	return
}

func (text *Text) GongMarshallField(stage *Stage, fieldName string) (res string) {
	var setValueField, setPointerField string
	_ = setValueField
	_ = setPointerField
	initializerStatements := ""
	_ = initializerStatements
	pointersInitializesStatements := ""
	_ = pointersInitializesStatements

	switch fieldName {
	case "Name":
		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", text.GongGetIdentifier(stage))
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Name")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(text.Name))
		initializerStatements += setValueField
	case "X":
		setValueField = NumberInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", text.GongGetIdentifier(stage))
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "X")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%f", text.X))
		initializerStatements += setValueField
	case "Y":
		setValueField = NumberInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", text.GongGetIdentifier(stage))
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Y")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%f", text.Y))
		initializerStatements += setValueField
	case "Content":
		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", text.GongGetIdentifier(stage))
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Content")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(text.Content))
		initializerStatements += setValueField
	case "Color":
		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", text.GongGetIdentifier(stage))
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Color")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(text.Color))
		initializerStatements += setValueField
	case "FillOpacity":
		setValueField = NumberInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", text.GongGetIdentifier(stage))
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "FillOpacity")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%f", text.FillOpacity))
		initializerStatements += setValueField
	case "Stroke":
		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", text.GongGetIdentifier(stage))
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Stroke")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(text.Stroke))
		initializerStatements += setValueField
	case "StrokeOpacity":
		setValueField = NumberInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", text.GongGetIdentifier(stage))
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "StrokeOpacity")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%f", text.StrokeOpacity))
		initializerStatements += setValueField
	case "StrokeWidth":
		setValueField = NumberInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", text.GongGetIdentifier(stage))
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "StrokeWidth")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%f", text.StrokeWidth))
		initializerStatements += setValueField
	case "StrokeDashArray":
		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", text.GongGetIdentifier(stage))
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "StrokeDashArray")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(text.StrokeDashArray))
		initializerStatements += setValueField
	case "StrokeDashArrayWhenSelected":
		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", text.GongGetIdentifier(stage))
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "StrokeDashArrayWhenSelected")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(text.StrokeDashArrayWhenSelected))
		initializerStatements += setValueField
	case "Transform":
		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", text.GongGetIdentifier(stage))
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Transform")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(text.Transform))
		initializerStatements += setValueField
	case "FontWeight":
		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", text.GongGetIdentifier(stage))
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "FontWeight")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(text.FontWeight))
		initializerStatements += setValueField
	case "FontSize":
		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", text.GongGetIdentifier(stage))
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "FontSize")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(text.FontSize))
		initializerStatements += setValueField
	case "FontStyle":
		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", text.GongGetIdentifier(stage))
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "FontStyle")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(text.FontStyle))
		initializerStatements += setValueField
	case "LetterSpacing":
		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", text.GongGetIdentifier(stage))
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "LetterSpacing")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(text.LetterSpacing))
		initializerStatements += setValueField
	case "FontFamily":
		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", text.GongGetIdentifier(stage))
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "FontFamily")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(text.FontFamily))
		initializerStatements += setValueField
	case "WhiteSpace":
		if text.WhiteSpace != "" {
			setValueField = StringEnumInitStatement
			setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", text.GongGetIdentifier(stage))
			setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "WhiteSpace")
			setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", "models."+text.WhiteSpace.ToCodeString())
			initializerStatements += setValueField
		}

	case "Animates":
		for _, _animate := range text.Animates {
			setPointerField = SliceOfPointersFieldInitStatement
			setPointerField = strings.ReplaceAll(setPointerField, "{{Identifier}}", text.GongGetIdentifier(stage))
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldName}}", "Animates")
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldNameValue}}", _animate.GongGetIdentifier(stage))
			pointersInitializesStatements += setPointerField
		}
	default:
		log.Panicf("Unknown field %s for Gongstruct Text", fieldName)
	}

	// temporary kludge to reuse existing template code
	res = initializerStatements + pointersInitializesStatements
	return
}
