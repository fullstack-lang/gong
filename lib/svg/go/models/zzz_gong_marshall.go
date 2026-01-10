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

// insertion point for marshall field methods
func (animate *Animate) GongMarshallField(stage *Stage, fieldName string) (res string) {

	switch fieldName {
	case "Name":
		res = StringInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", animate.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "Name")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", string(animate.Name))
	case "AttributeName":
		res = StringInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", animate.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "AttributeName")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", string(animate.AttributeName))
	case "Values":
		res = StringInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", animate.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "Values")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", string(animate.Values))
	case "From":
		res = StringInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", animate.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "From")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", string(animate.From))
	case "To":
		res = StringInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", animate.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "To")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", string(animate.To))
	case "Dur":
		res = StringInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", animate.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "Dur")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", string(animate.Dur))
	case "RepeatCount":
		res = StringInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", animate.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "RepeatCount")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", string(animate.RepeatCount))

	default:
		log.Panicf("Unknown field %s for Gongstruct Animate", fieldName)
	}
	return
}

func (circle *Circle) GongMarshallField(stage *Stage, fieldName string) (res string) {

	switch fieldName {
	case "Name":
		res = StringInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", circle.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "Name")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", string(circle.Name))
	case "CX":
		res = NumberInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", circle.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "CX")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%f", circle.CX))
	case "CY":
		res = NumberInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", circle.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "CY")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%f", circle.CY))
	case "Radius":
		res = NumberInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", circle.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "Radius")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%f", circle.Radius))
	case "Color":
		res = StringInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", circle.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "Color")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", string(circle.Color))
	case "FillOpacity":
		res = NumberInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", circle.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "FillOpacity")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%f", circle.FillOpacity))
	case "Stroke":
		res = StringInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", circle.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "Stroke")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", string(circle.Stroke))
	case "StrokeOpacity":
		res = NumberInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", circle.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "StrokeOpacity")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%f", circle.StrokeOpacity))
	case "StrokeWidth":
		res = NumberInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", circle.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "StrokeWidth")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%f", circle.StrokeWidth))
	case "StrokeDashArray":
		res = StringInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", circle.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "StrokeDashArray")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", string(circle.StrokeDashArray))
	case "StrokeDashArrayWhenSelected":
		res = StringInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", circle.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "StrokeDashArrayWhenSelected")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", string(circle.StrokeDashArrayWhenSelected))
	case "Transform":
		res = StringInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", circle.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "Transform")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", string(circle.Transform))

	case "Animations":
		for _, _animate := range circle.Animations {
			tmp := SliceOfPointersFieldInitStatement
			tmp = strings.ReplaceAll(tmp, "{{Identifier}}", circle.GongGetIdentifier(stage))
			tmp = strings.ReplaceAll(tmp, "{{GeneratedFieldName}}", "Animations")
			tmp = strings.ReplaceAll(tmp, "{{GeneratedFieldNameValue}}", _animate.GongGetIdentifier(stage))
			res += tmp
		}
	default:
		log.Panicf("Unknown field %s for Gongstruct Circle", fieldName)
	}
	return
}

func (condition *Condition) GongMarshallField(stage *Stage, fieldName string) (res string) {

	switch fieldName {
	case "Name":
		res = StringInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", condition.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "Name")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", string(condition.Name))

	default:
		log.Panicf("Unknown field %s for Gongstruct Condition", fieldName)
	}
	return
}

func (controlpoint *ControlPoint) GongMarshallField(stage *Stage, fieldName string) (res string) {

	switch fieldName {
	case "Name":
		res = StringInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", controlpoint.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "Name")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", string(controlpoint.Name))
	case "X_Relative":
		res = NumberInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", controlpoint.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "X_Relative")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%f", controlpoint.X_Relative))
	case "Y_Relative":
		res = NumberInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", controlpoint.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "Y_Relative")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%f", controlpoint.Y_Relative))

	case "ClosestRect":
		if controlpoint.ClosestRect != nil {
			res = PointerFieldInitStatement
			res = strings.ReplaceAll(res, "{{Identifier}}", controlpoint.GongGetIdentifier(stage))
			res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "ClosestRect")
			res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", controlpoint.ClosestRect.GongGetIdentifier(stage))
		} else {
			// in case of nil pointer, we need to unstage the previous value
			res = PointerFieldInitStatement
			res = strings.ReplaceAll(res, "{{Identifier}}", controlpoint.GongGetIdentifier(stage))
			res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "ClosestRect")
			res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", "nil")
		}
	default:
		log.Panicf("Unknown field %s for Gongstruct ControlPoint", fieldName)
	}
	return
}

func (ellipse *Ellipse) GongMarshallField(stage *Stage, fieldName string) (res string) {

	switch fieldName {
	case "Name":
		res = StringInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", ellipse.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "Name")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", string(ellipse.Name))
	case "CX":
		res = NumberInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", ellipse.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "CX")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%f", ellipse.CX))
	case "CY":
		res = NumberInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", ellipse.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "CY")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%f", ellipse.CY))
	case "RX":
		res = NumberInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", ellipse.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "RX")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%f", ellipse.RX))
	case "RY":
		res = NumberInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", ellipse.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "RY")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%f", ellipse.RY))
	case "Color":
		res = StringInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", ellipse.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "Color")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", string(ellipse.Color))
	case "FillOpacity":
		res = NumberInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", ellipse.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "FillOpacity")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%f", ellipse.FillOpacity))
	case "Stroke":
		res = StringInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", ellipse.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "Stroke")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", string(ellipse.Stroke))
	case "StrokeOpacity":
		res = NumberInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", ellipse.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "StrokeOpacity")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%f", ellipse.StrokeOpacity))
	case "StrokeWidth":
		res = NumberInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", ellipse.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "StrokeWidth")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%f", ellipse.StrokeWidth))
	case "StrokeDashArray":
		res = StringInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", ellipse.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "StrokeDashArray")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", string(ellipse.StrokeDashArray))
	case "StrokeDashArrayWhenSelected":
		res = StringInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", ellipse.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "StrokeDashArrayWhenSelected")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", string(ellipse.StrokeDashArrayWhenSelected))
	case "Transform":
		res = StringInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", ellipse.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "Transform")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", string(ellipse.Transform))

	case "Animates":
		for _, _animate := range ellipse.Animates {
			tmp := SliceOfPointersFieldInitStatement
			tmp = strings.ReplaceAll(tmp, "{{Identifier}}", ellipse.GongGetIdentifier(stage))
			tmp = strings.ReplaceAll(tmp, "{{GeneratedFieldName}}", "Animates")
			tmp = strings.ReplaceAll(tmp, "{{GeneratedFieldNameValue}}", _animate.GongGetIdentifier(stage))
			res += tmp
		}
	default:
		log.Panicf("Unknown field %s for Gongstruct Ellipse", fieldName)
	}
	return
}

func (layer *Layer) GongMarshallField(stage *Stage, fieldName string) (res string) {

	switch fieldName {
	case "Name":
		res = StringInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", layer.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "Name")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", string(layer.Name))

	case "Rects":
		for _, _rect := range layer.Rects {
			tmp := SliceOfPointersFieldInitStatement
			tmp = strings.ReplaceAll(tmp, "{{Identifier}}", layer.GongGetIdentifier(stage))
			tmp = strings.ReplaceAll(tmp, "{{GeneratedFieldName}}", "Rects")
			tmp = strings.ReplaceAll(tmp, "{{GeneratedFieldNameValue}}", _rect.GongGetIdentifier(stage))
			res += tmp
		}
	case "Texts":
		for _, _text := range layer.Texts {
			tmp := SliceOfPointersFieldInitStatement
			tmp = strings.ReplaceAll(tmp, "{{Identifier}}", layer.GongGetIdentifier(stage))
			tmp = strings.ReplaceAll(tmp, "{{GeneratedFieldName}}", "Texts")
			tmp = strings.ReplaceAll(tmp, "{{GeneratedFieldNameValue}}", _text.GongGetIdentifier(stage))
			res += tmp
		}
	case "Circles":
		for _, _circle := range layer.Circles {
			tmp := SliceOfPointersFieldInitStatement
			tmp = strings.ReplaceAll(tmp, "{{Identifier}}", layer.GongGetIdentifier(stage))
			tmp = strings.ReplaceAll(tmp, "{{GeneratedFieldName}}", "Circles")
			tmp = strings.ReplaceAll(tmp, "{{GeneratedFieldNameValue}}", _circle.GongGetIdentifier(stage))
			res += tmp
		}
	case "Lines":
		for _, _line := range layer.Lines {
			tmp := SliceOfPointersFieldInitStatement
			tmp = strings.ReplaceAll(tmp, "{{Identifier}}", layer.GongGetIdentifier(stage))
			tmp = strings.ReplaceAll(tmp, "{{GeneratedFieldName}}", "Lines")
			tmp = strings.ReplaceAll(tmp, "{{GeneratedFieldNameValue}}", _line.GongGetIdentifier(stage))
			res += tmp
		}
	case "Ellipses":
		for _, _ellipse := range layer.Ellipses {
			tmp := SliceOfPointersFieldInitStatement
			tmp = strings.ReplaceAll(tmp, "{{Identifier}}", layer.GongGetIdentifier(stage))
			tmp = strings.ReplaceAll(tmp, "{{GeneratedFieldName}}", "Ellipses")
			tmp = strings.ReplaceAll(tmp, "{{GeneratedFieldNameValue}}", _ellipse.GongGetIdentifier(stage))
			res += tmp
		}
	case "Polylines":
		for _, _polyline := range layer.Polylines {
			tmp := SliceOfPointersFieldInitStatement
			tmp = strings.ReplaceAll(tmp, "{{Identifier}}", layer.GongGetIdentifier(stage))
			tmp = strings.ReplaceAll(tmp, "{{GeneratedFieldName}}", "Polylines")
			tmp = strings.ReplaceAll(tmp, "{{GeneratedFieldNameValue}}", _polyline.GongGetIdentifier(stage))
			res += tmp
		}
	case "Polygones":
		for _, _polygone := range layer.Polygones {
			tmp := SliceOfPointersFieldInitStatement
			tmp = strings.ReplaceAll(tmp, "{{Identifier}}", layer.GongGetIdentifier(stage))
			tmp = strings.ReplaceAll(tmp, "{{GeneratedFieldName}}", "Polygones")
			tmp = strings.ReplaceAll(tmp, "{{GeneratedFieldNameValue}}", _polygone.GongGetIdentifier(stage))
			res += tmp
		}
	case "Paths":
		for _, _path := range layer.Paths {
			tmp := SliceOfPointersFieldInitStatement
			tmp = strings.ReplaceAll(tmp, "{{Identifier}}", layer.GongGetIdentifier(stage))
			tmp = strings.ReplaceAll(tmp, "{{GeneratedFieldName}}", "Paths")
			tmp = strings.ReplaceAll(tmp, "{{GeneratedFieldNameValue}}", _path.GongGetIdentifier(stage))
			res += tmp
		}
	case "Links":
		for _, _link := range layer.Links {
			tmp := SliceOfPointersFieldInitStatement
			tmp = strings.ReplaceAll(tmp, "{{Identifier}}", layer.GongGetIdentifier(stage))
			tmp = strings.ReplaceAll(tmp, "{{GeneratedFieldName}}", "Links")
			tmp = strings.ReplaceAll(tmp, "{{GeneratedFieldNameValue}}", _link.GongGetIdentifier(stage))
			res += tmp
		}
	case "RectLinkLinks":
		for _, _rectlinklink := range layer.RectLinkLinks {
			tmp := SliceOfPointersFieldInitStatement
			tmp = strings.ReplaceAll(tmp, "{{Identifier}}", layer.GongGetIdentifier(stage))
			tmp = strings.ReplaceAll(tmp, "{{GeneratedFieldName}}", "RectLinkLinks")
			tmp = strings.ReplaceAll(tmp, "{{GeneratedFieldNameValue}}", _rectlinklink.GongGetIdentifier(stage))
			res += tmp
		}
	default:
		log.Panicf("Unknown field %s for Gongstruct Layer", fieldName)
	}
	return
}

func (line *Line) GongMarshallField(stage *Stage, fieldName string) (res string) {

	switch fieldName {
	case "Name":
		res = StringInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", line.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "Name")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", string(line.Name))
	case "X1":
		res = NumberInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", line.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "X1")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%f", line.X1))
	case "Y1":
		res = NumberInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", line.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "Y1")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%f", line.Y1))
	case "X2":
		res = NumberInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", line.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "X2")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%f", line.X2))
	case "Y2":
		res = NumberInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", line.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "Y2")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%f", line.Y2))
	case "Color":
		res = StringInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", line.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "Color")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", string(line.Color))
	case "FillOpacity":
		res = NumberInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", line.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "FillOpacity")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%f", line.FillOpacity))
	case "Stroke":
		res = StringInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", line.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "Stroke")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", string(line.Stroke))
	case "StrokeOpacity":
		res = NumberInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", line.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "StrokeOpacity")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%f", line.StrokeOpacity))
	case "StrokeWidth":
		res = NumberInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", line.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "StrokeWidth")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%f", line.StrokeWidth))
	case "StrokeDashArray":
		res = StringInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", line.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "StrokeDashArray")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", string(line.StrokeDashArray))
	case "StrokeDashArrayWhenSelected":
		res = StringInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", line.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "StrokeDashArrayWhenSelected")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", string(line.StrokeDashArrayWhenSelected))
	case "Transform":
		res = StringInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", line.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "Transform")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", string(line.Transform))
	case "MouseClickX":
		res = NumberInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", line.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "MouseClickX")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%f", line.MouseClickX))
	case "MouseClickY":
		res = NumberInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", line.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "MouseClickY")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%f", line.MouseClickY))

	case "Animates":
		for _, _animate := range line.Animates {
			tmp := SliceOfPointersFieldInitStatement
			tmp = strings.ReplaceAll(tmp, "{{Identifier}}", line.GongGetIdentifier(stage))
			tmp = strings.ReplaceAll(tmp, "{{GeneratedFieldName}}", "Animates")
			tmp = strings.ReplaceAll(tmp, "{{GeneratedFieldNameValue}}", _animate.GongGetIdentifier(stage))
			res += tmp
		}
	default:
		log.Panicf("Unknown field %s for Gongstruct Line", fieldName)
	}
	return
}

func (link *Link) GongMarshallField(stage *Stage, fieldName string) (res string) {

	switch fieldName {
	case "Name":
		res = StringInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", link.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "Name")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", string(link.Name))
	case "Type":
		if link.Type != "" {
			res = StringEnumInitStatement
			res = strings.ReplaceAll(res, "{{Identifier}}", link.GongGetIdentifier(stage))
			res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "Type")
			res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", "models."+link.Type.ToCodeString())
		}
	case "IsBezierCurve":
		res = NumberInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", link.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "IsBezierCurve")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%t", link.IsBezierCurve))
	case "StartAnchorType":
		if link.StartAnchorType != "" {
			res = StringEnumInitStatement
			res = strings.ReplaceAll(res, "{{Identifier}}", link.GongGetIdentifier(stage))
			res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "StartAnchorType")
			res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", "models."+link.StartAnchorType.ToCodeString())
		}
	case "EndAnchorType":
		if link.EndAnchorType != "" {
			res = StringEnumInitStatement
			res = strings.ReplaceAll(res, "{{Identifier}}", link.GongGetIdentifier(stage))
			res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "EndAnchorType")
			res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", "models."+link.EndAnchorType.ToCodeString())
		}
	case "StartOrientation":
		if link.StartOrientation != "" {
			res = StringEnumInitStatement
			res = strings.ReplaceAll(res, "{{Identifier}}", link.GongGetIdentifier(stage))
			res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "StartOrientation")
			res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", "models."+link.StartOrientation.ToCodeString())
		}
	case "StartRatio":
		res = NumberInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", link.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "StartRatio")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%f", link.StartRatio))
	case "EndOrientation":
		if link.EndOrientation != "" {
			res = StringEnumInitStatement
			res = strings.ReplaceAll(res, "{{Identifier}}", link.GongGetIdentifier(stage))
			res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "EndOrientation")
			res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", "models."+link.EndOrientation.ToCodeString())
		}
	case "EndRatio":
		res = NumberInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", link.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "EndRatio")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%f", link.EndRatio))
	case "CornerOffsetRatio":
		res = NumberInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", link.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "CornerOffsetRatio")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%f", link.CornerOffsetRatio))
	case "CornerRadius":
		res = NumberInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", link.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "CornerRadius")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%f", link.CornerRadius))
	case "HasEndArrow":
		res = NumberInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", link.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "HasEndArrow")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%t", link.HasEndArrow))
	case "EndArrowSize":
		res = NumberInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", link.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "EndArrowSize")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%f", link.EndArrowSize))
	case "EndArrowOffset":
		res = NumberInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", link.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "EndArrowOffset")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%f", link.EndArrowOffset))
	case "HasStartArrow":
		res = NumberInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", link.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "HasStartArrow")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%t", link.HasStartArrow))
	case "StartArrowSize":
		res = NumberInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", link.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "StartArrowSize")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%f", link.StartArrowSize))
	case "StartArrowOffset":
		res = NumberInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", link.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "StartArrowOffset")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%f", link.StartArrowOffset))
	case "Color":
		res = StringInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", link.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "Color")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", string(link.Color))
	case "FillOpacity":
		res = NumberInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", link.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "FillOpacity")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%f", link.FillOpacity))
	case "Stroke":
		res = StringInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", link.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "Stroke")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", string(link.Stroke))
	case "StrokeOpacity":
		res = NumberInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", link.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "StrokeOpacity")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%f", link.StrokeOpacity))
	case "StrokeWidth":
		res = NumberInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", link.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "StrokeWidth")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%f", link.StrokeWidth))
	case "StrokeDashArray":
		res = StringInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", link.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "StrokeDashArray")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", string(link.StrokeDashArray))
	case "StrokeDashArrayWhenSelected":
		res = StringInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", link.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "StrokeDashArrayWhenSelected")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", string(link.StrokeDashArrayWhenSelected))
	case "Transform":
		res = StringInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", link.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "Transform")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", string(link.Transform))
	case "MouseX":
		res = NumberInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", link.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "MouseX")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%f", link.MouseX))
	case "MouseY":
		res = NumberInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", link.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "MouseY")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%f", link.MouseY))
	case "MouseEventKey":
		if link.MouseEventKey != "" {
			res = StringEnumInitStatement
			res = strings.ReplaceAll(res, "{{Identifier}}", link.GongGetIdentifier(stage))
			res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "MouseEventKey")
			res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", "models."+link.MouseEventKey.ToCodeString())
		}

	case "Start":
		if link.Start != nil {
			res = PointerFieldInitStatement
			res = strings.ReplaceAll(res, "{{Identifier}}", link.GongGetIdentifier(stage))
			res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "Start")
			res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", link.Start.GongGetIdentifier(stage))
		} else {
			// in case of nil pointer, we need to unstage the previous value
			res = PointerFieldInitStatement
			res = strings.ReplaceAll(res, "{{Identifier}}", link.GongGetIdentifier(stage))
			res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "Start")
			res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", "nil")
		}
	case "End":
		if link.End != nil {
			res = PointerFieldInitStatement
			res = strings.ReplaceAll(res, "{{Identifier}}", link.GongGetIdentifier(stage))
			res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "End")
			res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", link.End.GongGetIdentifier(stage))
		} else {
			// in case of nil pointer, we need to unstage the previous value
			res = PointerFieldInitStatement
			res = strings.ReplaceAll(res, "{{Identifier}}", link.GongGetIdentifier(stage))
			res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "End")
			res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", "nil")
		}
	case "TextAtArrowStart":
		for _, _linkanchoredtext := range link.TextAtArrowStart {
			tmp := SliceOfPointersFieldInitStatement
			tmp = strings.ReplaceAll(tmp, "{{Identifier}}", link.GongGetIdentifier(stage))
			tmp = strings.ReplaceAll(tmp, "{{GeneratedFieldName}}", "TextAtArrowStart")
			tmp = strings.ReplaceAll(tmp, "{{GeneratedFieldNameValue}}", _linkanchoredtext.GongGetIdentifier(stage))
			res += tmp
		}
	case "TextAtArrowEnd":
		for _, _linkanchoredtext := range link.TextAtArrowEnd {
			tmp := SliceOfPointersFieldInitStatement
			tmp = strings.ReplaceAll(tmp, "{{Identifier}}", link.GongGetIdentifier(stage))
			tmp = strings.ReplaceAll(tmp, "{{GeneratedFieldName}}", "TextAtArrowEnd")
			tmp = strings.ReplaceAll(tmp, "{{GeneratedFieldNameValue}}", _linkanchoredtext.GongGetIdentifier(stage))
			res += tmp
		}
	case "ControlPoints":
		for _, _controlpoint := range link.ControlPoints {
			tmp := SliceOfPointersFieldInitStatement
			tmp = strings.ReplaceAll(tmp, "{{Identifier}}", link.GongGetIdentifier(stage))
			tmp = strings.ReplaceAll(tmp, "{{GeneratedFieldName}}", "ControlPoints")
			tmp = strings.ReplaceAll(tmp, "{{GeneratedFieldNameValue}}", _controlpoint.GongGetIdentifier(stage))
			res += tmp
		}
	default:
		log.Panicf("Unknown field %s for Gongstruct Link", fieldName)
	}
	return
}

func (linkanchoredtext *LinkAnchoredText) GongMarshallField(stage *Stage, fieldName string) (res string) {

	switch fieldName {
	case "Name":
		res = StringInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", linkanchoredtext.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "Name")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", string(linkanchoredtext.Name))
	case "Content":
		res = StringInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", linkanchoredtext.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "Content")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", string(linkanchoredtext.Content))
	case "AutomaticLayout":
		res = NumberInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", linkanchoredtext.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "AutomaticLayout")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%t", linkanchoredtext.AutomaticLayout))
	case "LinkAnchorType":
		if linkanchoredtext.LinkAnchorType != "" {
			res = StringEnumInitStatement
			res = strings.ReplaceAll(res, "{{Identifier}}", linkanchoredtext.GongGetIdentifier(stage))
			res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "LinkAnchorType")
			res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", "models."+linkanchoredtext.LinkAnchorType.ToCodeString())
		}
	case "X_Offset":
		res = NumberInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", linkanchoredtext.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "X_Offset")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%f", linkanchoredtext.X_Offset))
	case "Y_Offset":
		res = NumberInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", linkanchoredtext.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "Y_Offset")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%f", linkanchoredtext.Y_Offset))
	case "FontWeight":
		res = StringInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", linkanchoredtext.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "FontWeight")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", string(linkanchoredtext.FontWeight))
	case "FontSize":
		res = StringInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", linkanchoredtext.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "FontSize")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", string(linkanchoredtext.FontSize))
	case "FontStyle":
		res = StringInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", linkanchoredtext.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "FontStyle")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", string(linkanchoredtext.FontStyle))
	case "LetterSpacing":
		res = StringInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", linkanchoredtext.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "LetterSpacing")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", string(linkanchoredtext.LetterSpacing))
	case "FontFamily":
		res = StringInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", linkanchoredtext.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "FontFamily")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", string(linkanchoredtext.FontFamily))
	case "WhiteSpace":
		if linkanchoredtext.WhiteSpace != "" {
			res = StringEnumInitStatement
			res = strings.ReplaceAll(res, "{{Identifier}}", linkanchoredtext.GongGetIdentifier(stage))
			res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "WhiteSpace")
			res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", "models."+linkanchoredtext.WhiteSpace.ToCodeString())
		}
	case "Color":
		res = StringInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", linkanchoredtext.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "Color")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", string(linkanchoredtext.Color))
	case "FillOpacity":
		res = NumberInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", linkanchoredtext.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "FillOpacity")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%f", linkanchoredtext.FillOpacity))
	case "Stroke":
		res = StringInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", linkanchoredtext.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "Stroke")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", string(linkanchoredtext.Stroke))
	case "StrokeOpacity":
		res = NumberInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", linkanchoredtext.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "StrokeOpacity")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%f", linkanchoredtext.StrokeOpacity))
	case "StrokeWidth":
		res = NumberInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", linkanchoredtext.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "StrokeWidth")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%f", linkanchoredtext.StrokeWidth))
	case "StrokeDashArray":
		res = StringInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", linkanchoredtext.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "StrokeDashArray")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", string(linkanchoredtext.StrokeDashArray))
	case "StrokeDashArrayWhenSelected":
		res = StringInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", linkanchoredtext.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "StrokeDashArrayWhenSelected")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", string(linkanchoredtext.StrokeDashArrayWhenSelected))
	case "Transform":
		res = StringInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", linkanchoredtext.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "Transform")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", string(linkanchoredtext.Transform))

	case "Animates":
		for _, _animate := range linkanchoredtext.Animates {
			tmp := SliceOfPointersFieldInitStatement
			tmp = strings.ReplaceAll(tmp, "{{Identifier}}", linkanchoredtext.GongGetIdentifier(stage))
			tmp = strings.ReplaceAll(tmp, "{{GeneratedFieldName}}", "Animates")
			tmp = strings.ReplaceAll(tmp, "{{GeneratedFieldNameValue}}", _animate.GongGetIdentifier(stage))
			res += tmp
		}
	default:
		log.Panicf("Unknown field %s for Gongstruct LinkAnchoredText", fieldName)
	}
	return
}

func (path *Path) GongMarshallField(stage *Stage, fieldName string) (res string) {

	switch fieldName {
	case "Name":
		res = StringInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", path.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "Name")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", string(path.Name))
	case "Definition":
		res = StringInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", path.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "Definition")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", string(path.Definition))
	case "Color":
		res = StringInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", path.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "Color")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", string(path.Color))
	case "FillOpacity":
		res = NumberInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", path.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "FillOpacity")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%f", path.FillOpacity))
	case "Stroke":
		res = StringInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", path.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "Stroke")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", string(path.Stroke))
	case "StrokeOpacity":
		res = NumberInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", path.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "StrokeOpacity")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%f", path.StrokeOpacity))
	case "StrokeWidth":
		res = NumberInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", path.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "StrokeWidth")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%f", path.StrokeWidth))
	case "StrokeDashArray":
		res = StringInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", path.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "StrokeDashArray")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", string(path.StrokeDashArray))
	case "StrokeDashArrayWhenSelected":
		res = StringInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", path.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "StrokeDashArrayWhenSelected")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", string(path.StrokeDashArrayWhenSelected))
	case "Transform":
		res = StringInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", path.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "Transform")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", string(path.Transform))

	case "Animates":
		for _, _animate := range path.Animates {
			tmp := SliceOfPointersFieldInitStatement
			tmp = strings.ReplaceAll(tmp, "{{Identifier}}", path.GongGetIdentifier(stage))
			tmp = strings.ReplaceAll(tmp, "{{GeneratedFieldName}}", "Animates")
			tmp = strings.ReplaceAll(tmp, "{{GeneratedFieldNameValue}}", _animate.GongGetIdentifier(stage))
			res += tmp
		}
	default:
		log.Panicf("Unknown field %s for Gongstruct Path", fieldName)
	}
	return
}

func (point *Point) GongMarshallField(stage *Stage, fieldName string) (res string) {

	switch fieldName {
	case "Name":
		res = StringInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", point.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "Name")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", string(point.Name))
	case "X":
		res = NumberInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", point.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "X")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%f", point.X))
	case "Y":
		res = NumberInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", point.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "Y")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%f", point.Y))

	default:
		log.Panicf("Unknown field %s for Gongstruct Point", fieldName)
	}
	return
}

func (polygone *Polygone) GongMarshallField(stage *Stage, fieldName string) (res string) {

	switch fieldName {
	case "Name":
		res = StringInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", polygone.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "Name")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", string(polygone.Name))
	case "Points":
		res = StringInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", polygone.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "Points")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", string(polygone.Points))
	case "Color":
		res = StringInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", polygone.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "Color")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", string(polygone.Color))
	case "FillOpacity":
		res = NumberInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", polygone.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "FillOpacity")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%f", polygone.FillOpacity))
	case "Stroke":
		res = StringInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", polygone.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "Stroke")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", string(polygone.Stroke))
	case "StrokeOpacity":
		res = NumberInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", polygone.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "StrokeOpacity")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%f", polygone.StrokeOpacity))
	case "StrokeWidth":
		res = NumberInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", polygone.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "StrokeWidth")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%f", polygone.StrokeWidth))
	case "StrokeDashArray":
		res = StringInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", polygone.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "StrokeDashArray")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", string(polygone.StrokeDashArray))
	case "StrokeDashArrayWhenSelected":
		res = StringInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", polygone.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "StrokeDashArrayWhenSelected")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", string(polygone.StrokeDashArrayWhenSelected))
	case "Transform":
		res = StringInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", polygone.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "Transform")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", string(polygone.Transform))

	case "Animates":
		for _, _animate := range polygone.Animates {
			tmp := SliceOfPointersFieldInitStatement
			tmp = strings.ReplaceAll(tmp, "{{Identifier}}", polygone.GongGetIdentifier(stage))
			tmp = strings.ReplaceAll(tmp, "{{GeneratedFieldName}}", "Animates")
			tmp = strings.ReplaceAll(tmp, "{{GeneratedFieldNameValue}}", _animate.GongGetIdentifier(stage))
			res += tmp
		}
	default:
		log.Panicf("Unknown field %s for Gongstruct Polygone", fieldName)
	}
	return
}

func (polyline *Polyline) GongMarshallField(stage *Stage, fieldName string) (res string) {

	switch fieldName {
	case "Name":
		res = StringInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", polyline.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "Name")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", string(polyline.Name))
	case "Points":
		res = StringInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", polyline.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "Points")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", string(polyline.Points))
	case "Color":
		res = StringInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", polyline.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "Color")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", string(polyline.Color))
	case "FillOpacity":
		res = NumberInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", polyline.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "FillOpacity")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%f", polyline.FillOpacity))
	case "Stroke":
		res = StringInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", polyline.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "Stroke")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", string(polyline.Stroke))
	case "StrokeOpacity":
		res = NumberInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", polyline.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "StrokeOpacity")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%f", polyline.StrokeOpacity))
	case "StrokeWidth":
		res = NumberInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", polyline.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "StrokeWidth")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%f", polyline.StrokeWidth))
	case "StrokeDashArray":
		res = StringInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", polyline.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "StrokeDashArray")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", string(polyline.StrokeDashArray))
	case "StrokeDashArrayWhenSelected":
		res = StringInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", polyline.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "StrokeDashArrayWhenSelected")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", string(polyline.StrokeDashArrayWhenSelected))
	case "Transform":
		res = StringInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", polyline.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "Transform")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", string(polyline.Transform))

	case "Animates":
		for _, _animate := range polyline.Animates {
			tmp := SliceOfPointersFieldInitStatement
			tmp = strings.ReplaceAll(tmp, "{{Identifier}}", polyline.GongGetIdentifier(stage))
			tmp = strings.ReplaceAll(tmp, "{{GeneratedFieldName}}", "Animates")
			tmp = strings.ReplaceAll(tmp, "{{GeneratedFieldNameValue}}", _animate.GongGetIdentifier(stage))
			res += tmp
		}
	default:
		log.Panicf("Unknown field %s for Gongstruct Polyline", fieldName)
	}
	return
}

func (rect *Rect) GongMarshallField(stage *Stage, fieldName string) (res string) {

	switch fieldName {
	case "Name":
		res = StringInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", rect.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "Name")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", string(rect.Name))
	case "X":
		res = NumberInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", rect.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "X")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%f", rect.X))
	case "Y":
		res = NumberInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", rect.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "Y")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%f", rect.Y))
	case "Width":
		res = NumberInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", rect.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "Width")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%f", rect.Width))
	case "Height":
		res = NumberInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", rect.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "Height")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%f", rect.Height))
	case "RX":
		res = NumberInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", rect.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "RX")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%f", rect.RX))
	case "Color":
		res = StringInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", rect.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "Color")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", string(rect.Color))
	case "FillOpacity":
		res = NumberInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", rect.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "FillOpacity")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%f", rect.FillOpacity))
	case "Stroke":
		res = StringInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", rect.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "Stroke")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", string(rect.Stroke))
	case "StrokeOpacity":
		res = NumberInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", rect.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "StrokeOpacity")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%f", rect.StrokeOpacity))
	case "StrokeWidth":
		res = NumberInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", rect.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "StrokeWidth")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%f", rect.StrokeWidth))
	case "StrokeDashArray":
		res = StringInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", rect.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "StrokeDashArray")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", string(rect.StrokeDashArray))
	case "StrokeDashArrayWhenSelected":
		res = StringInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", rect.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "StrokeDashArrayWhenSelected")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", string(rect.StrokeDashArrayWhenSelected))
	case "Transform":
		res = StringInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", rect.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "Transform")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", string(rect.Transform))
	case "IsSelectable":
		res = NumberInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", rect.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "IsSelectable")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%t", rect.IsSelectable))
	case "IsSelected":
		res = NumberInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", rect.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "IsSelected")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%t", rect.IsSelected))
	case "CanHaveLeftHandle":
		res = NumberInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", rect.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "CanHaveLeftHandle")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%t", rect.CanHaveLeftHandle))
	case "HasLeftHandle":
		res = NumberInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", rect.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "HasLeftHandle")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%t", rect.HasLeftHandle))
	case "CanHaveRightHandle":
		res = NumberInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", rect.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "CanHaveRightHandle")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%t", rect.CanHaveRightHandle))
	case "HasRightHandle":
		res = NumberInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", rect.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "HasRightHandle")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%t", rect.HasRightHandle))
	case "CanHaveTopHandle":
		res = NumberInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", rect.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "CanHaveTopHandle")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%t", rect.CanHaveTopHandle))
	case "HasTopHandle":
		res = NumberInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", rect.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "HasTopHandle")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%t", rect.HasTopHandle))
	case "IsScalingProportionally":
		res = NumberInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", rect.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "IsScalingProportionally")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%t", rect.IsScalingProportionally))
	case "CanHaveBottomHandle":
		res = NumberInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", rect.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "CanHaveBottomHandle")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%t", rect.CanHaveBottomHandle))
	case "HasBottomHandle":
		res = NumberInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", rect.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "HasBottomHandle")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%t", rect.HasBottomHandle))
	case "CanMoveHorizontaly":
		res = NumberInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", rect.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "CanMoveHorizontaly")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%t", rect.CanMoveHorizontaly))
	case "CanMoveVerticaly":
		res = NumberInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", rect.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "CanMoveVerticaly")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%t", rect.CanMoveVerticaly))
	case "ChangeColorWhenHovered":
		res = NumberInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", rect.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "ChangeColorWhenHovered")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%t", rect.ChangeColorWhenHovered))
	case "ColorWhenHovered":
		res = StringInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", rect.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "ColorWhenHovered")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", string(rect.ColorWhenHovered))
	case "OriginalColor":
		res = StringInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", rect.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "OriginalColor")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", string(rect.OriginalColor))
	case "FillOpacityWhenHovered":
		res = NumberInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", rect.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "FillOpacityWhenHovered")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%f", rect.FillOpacityWhenHovered))
	case "OriginalFillOpacity":
		res = NumberInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", rect.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "OriginalFillOpacity")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%f", rect.OriginalFillOpacity))
	case "HasToolTip":
		res = NumberInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", rect.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "HasToolTip")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%t", rect.HasToolTip))
	case "ToolTipText":
		res = StringInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", rect.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "ToolTipText")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", string(rect.ToolTipText))
	case "ToolTipPosition":
		if rect.ToolTipPosition != "" {
			res = StringEnumInitStatement
			res = strings.ReplaceAll(res, "{{Identifier}}", rect.GongGetIdentifier(stage))
			res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "ToolTipPosition")
			res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", "models."+rect.ToolTipPosition.ToCodeString())
		}
	case "MouseX":
		res = NumberInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", rect.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "MouseX")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%f", rect.MouseX))
	case "MouseY":
		res = NumberInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", rect.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "MouseY")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%f", rect.MouseY))
	case "MouseEventKey":
		if rect.MouseEventKey != "" {
			res = StringEnumInitStatement
			res = strings.ReplaceAll(res, "{{Identifier}}", rect.GongGetIdentifier(stage))
			res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "MouseEventKey")
			res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", "models."+rect.MouseEventKey.ToCodeString())
		}

	case "HoveringTrigger":
		for _, _condition := range rect.HoveringTrigger {
			tmp := SliceOfPointersFieldInitStatement
			tmp = strings.ReplaceAll(tmp, "{{Identifier}}", rect.GongGetIdentifier(stage))
			tmp = strings.ReplaceAll(tmp, "{{GeneratedFieldName}}", "HoveringTrigger")
			tmp = strings.ReplaceAll(tmp, "{{GeneratedFieldNameValue}}", _condition.GongGetIdentifier(stage))
			res += tmp
		}
	case "DisplayConditions":
		for _, _condition := range rect.DisplayConditions {
			tmp := SliceOfPointersFieldInitStatement
			tmp = strings.ReplaceAll(tmp, "{{Identifier}}", rect.GongGetIdentifier(stage))
			tmp = strings.ReplaceAll(tmp, "{{GeneratedFieldName}}", "DisplayConditions")
			tmp = strings.ReplaceAll(tmp, "{{GeneratedFieldNameValue}}", _condition.GongGetIdentifier(stage))
			res += tmp
		}
	case "Animations":
		for _, _animate := range rect.Animations {
			tmp := SliceOfPointersFieldInitStatement
			tmp = strings.ReplaceAll(tmp, "{{Identifier}}", rect.GongGetIdentifier(stage))
			tmp = strings.ReplaceAll(tmp, "{{GeneratedFieldName}}", "Animations")
			tmp = strings.ReplaceAll(tmp, "{{GeneratedFieldNameValue}}", _animate.GongGetIdentifier(stage))
			res += tmp
		}
	case "RectAnchoredTexts":
		for _, _rectanchoredtext := range rect.RectAnchoredTexts {
			tmp := SliceOfPointersFieldInitStatement
			tmp = strings.ReplaceAll(tmp, "{{Identifier}}", rect.GongGetIdentifier(stage))
			tmp = strings.ReplaceAll(tmp, "{{GeneratedFieldName}}", "RectAnchoredTexts")
			tmp = strings.ReplaceAll(tmp, "{{GeneratedFieldNameValue}}", _rectanchoredtext.GongGetIdentifier(stage))
			res += tmp
		}
	case "RectAnchoredRects":
		for _, _rectanchoredrect := range rect.RectAnchoredRects {
			tmp := SliceOfPointersFieldInitStatement
			tmp = strings.ReplaceAll(tmp, "{{Identifier}}", rect.GongGetIdentifier(stage))
			tmp = strings.ReplaceAll(tmp, "{{GeneratedFieldName}}", "RectAnchoredRects")
			tmp = strings.ReplaceAll(tmp, "{{GeneratedFieldNameValue}}", _rectanchoredrect.GongGetIdentifier(stage))
			res += tmp
		}
	case "RectAnchoredPaths":
		for _, _rectanchoredpath := range rect.RectAnchoredPaths {
			tmp := SliceOfPointersFieldInitStatement
			tmp = strings.ReplaceAll(tmp, "{{Identifier}}", rect.GongGetIdentifier(stage))
			tmp = strings.ReplaceAll(tmp, "{{GeneratedFieldName}}", "RectAnchoredPaths")
			tmp = strings.ReplaceAll(tmp, "{{GeneratedFieldNameValue}}", _rectanchoredpath.GongGetIdentifier(stage))
			res += tmp
		}
	default:
		log.Panicf("Unknown field %s for Gongstruct Rect", fieldName)
	}
	return
}

func (rectanchoredpath *RectAnchoredPath) GongMarshallField(stage *Stage, fieldName string) (res string) {

	switch fieldName {
	case "Name":
		res = StringInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", rectanchoredpath.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "Name")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", string(rectanchoredpath.Name))
	case "Definition":
		res = StringInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", rectanchoredpath.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "Definition")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", string(rectanchoredpath.Definition))
	case "X_Offset":
		res = NumberInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", rectanchoredpath.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "X_Offset")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%f", rectanchoredpath.X_Offset))
	case "Y_Offset":
		res = NumberInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", rectanchoredpath.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "Y_Offset")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%f", rectanchoredpath.Y_Offset))
	case "RectAnchorType":
		if rectanchoredpath.RectAnchorType != "" {
			res = StringEnumInitStatement
			res = strings.ReplaceAll(res, "{{Identifier}}", rectanchoredpath.GongGetIdentifier(stage))
			res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "RectAnchorType")
			res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", "models."+rectanchoredpath.RectAnchorType.ToCodeString())
		}
	case "ScalePropotionnally":
		res = NumberInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", rectanchoredpath.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "ScalePropotionnally")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%t", rectanchoredpath.ScalePropotionnally))
	case "AppliedScaling":
		res = NumberInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", rectanchoredpath.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "AppliedScaling")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%f", rectanchoredpath.AppliedScaling))
	case "Color":
		res = StringInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", rectanchoredpath.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "Color")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", string(rectanchoredpath.Color))
	case "FillOpacity":
		res = NumberInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", rectanchoredpath.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "FillOpacity")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%f", rectanchoredpath.FillOpacity))
	case "Stroke":
		res = StringInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", rectanchoredpath.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "Stroke")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", string(rectanchoredpath.Stroke))
	case "StrokeOpacity":
		res = NumberInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", rectanchoredpath.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "StrokeOpacity")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%f", rectanchoredpath.StrokeOpacity))
	case "StrokeWidth":
		res = NumberInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", rectanchoredpath.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "StrokeWidth")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%f", rectanchoredpath.StrokeWidth))
	case "StrokeDashArray":
		res = StringInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", rectanchoredpath.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "StrokeDashArray")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", string(rectanchoredpath.StrokeDashArray))
	case "StrokeDashArrayWhenSelected":
		res = StringInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", rectanchoredpath.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "StrokeDashArrayWhenSelected")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", string(rectanchoredpath.StrokeDashArrayWhenSelected))
	case "Transform":
		res = StringInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", rectanchoredpath.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "Transform")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", string(rectanchoredpath.Transform))

	default:
		log.Panicf("Unknown field %s for Gongstruct RectAnchoredPath", fieldName)
	}
	return
}

func (rectanchoredrect *RectAnchoredRect) GongMarshallField(stage *Stage, fieldName string) (res string) {

	switch fieldName {
	case "Name":
		res = StringInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", rectanchoredrect.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "Name")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", string(rectanchoredrect.Name))
	case "X":
		res = NumberInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", rectanchoredrect.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "X")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%f", rectanchoredrect.X))
	case "Y":
		res = NumberInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", rectanchoredrect.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "Y")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%f", rectanchoredrect.Y))
	case "Width":
		res = NumberInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", rectanchoredrect.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "Width")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%f", rectanchoredrect.Width))
	case "Height":
		res = NumberInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", rectanchoredrect.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "Height")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%f", rectanchoredrect.Height))
	case "RX":
		res = NumberInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", rectanchoredrect.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "RX")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%f", rectanchoredrect.RX))
	case "X_Offset":
		res = NumberInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", rectanchoredrect.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "X_Offset")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%f", rectanchoredrect.X_Offset))
	case "Y_Offset":
		res = NumberInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", rectanchoredrect.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "Y_Offset")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%f", rectanchoredrect.Y_Offset))
	case "RectAnchorType":
		if rectanchoredrect.RectAnchorType != "" {
			res = StringEnumInitStatement
			res = strings.ReplaceAll(res, "{{Identifier}}", rectanchoredrect.GongGetIdentifier(stage))
			res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "RectAnchorType")
			res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", "models."+rectanchoredrect.RectAnchorType.ToCodeString())
		}
	case "WidthFollowRect":
		res = NumberInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", rectanchoredrect.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "WidthFollowRect")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%t", rectanchoredrect.WidthFollowRect))
	case "HeightFollowRect":
		res = NumberInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", rectanchoredrect.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "HeightFollowRect")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%t", rectanchoredrect.HeightFollowRect))
	case "HasToolTip":
		res = NumberInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", rectanchoredrect.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "HasToolTip")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%t", rectanchoredrect.HasToolTip))
	case "ToolTipText":
		res = StringInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", rectanchoredrect.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "ToolTipText")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", string(rectanchoredrect.ToolTipText))
	case "Color":
		res = StringInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", rectanchoredrect.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "Color")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", string(rectanchoredrect.Color))
	case "FillOpacity":
		res = NumberInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", rectanchoredrect.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "FillOpacity")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%f", rectanchoredrect.FillOpacity))
	case "Stroke":
		res = StringInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", rectanchoredrect.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "Stroke")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", string(rectanchoredrect.Stroke))
	case "StrokeOpacity":
		res = NumberInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", rectanchoredrect.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "StrokeOpacity")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%f", rectanchoredrect.StrokeOpacity))
	case "StrokeWidth":
		res = NumberInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", rectanchoredrect.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "StrokeWidth")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%f", rectanchoredrect.StrokeWidth))
	case "StrokeDashArray":
		res = StringInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", rectanchoredrect.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "StrokeDashArray")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", string(rectanchoredrect.StrokeDashArray))
	case "StrokeDashArrayWhenSelected":
		res = StringInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", rectanchoredrect.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "StrokeDashArrayWhenSelected")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", string(rectanchoredrect.StrokeDashArrayWhenSelected))
	case "Transform":
		res = StringInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", rectanchoredrect.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "Transform")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", string(rectanchoredrect.Transform))

	default:
		log.Panicf("Unknown field %s for Gongstruct RectAnchoredRect", fieldName)
	}
	return
}

func (rectanchoredtext *RectAnchoredText) GongMarshallField(stage *Stage, fieldName string) (res string) {

	switch fieldName {
	case "Name":
		res = StringInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", rectanchoredtext.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "Name")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", string(rectanchoredtext.Name))
	case "Content":
		res = StringInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", rectanchoredtext.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "Content")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", string(rectanchoredtext.Content))
	case "FontWeight":
		res = StringInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", rectanchoredtext.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "FontWeight")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", string(rectanchoredtext.FontWeight))
	case "FontSize":
		res = StringInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", rectanchoredtext.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "FontSize")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", string(rectanchoredtext.FontSize))
	case "FontStyle":
		res = StringInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", rectanchoredtext.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "FontStyle")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", string(rectanchoredtext.FontStyle))
	case "LetterSpacing":
		res = StringInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", rectanchoredtext.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "LetterSpacing")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", string(rectanchoredtext.LetterSpacing))
	case "FontFamily":
		res = StringInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", rectanchoredtext.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "FontFamily")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", string(rectanchoredtext.FontFamily))
	case "WhiteSpace":
		if rectanchoredtext.WhiteSpace != "" {
			res = StringEnumInitStatement
			res = strings.ReplaceAll(res, "{{Identifier}}", rectanchoredtext.GongGetIdentifier(stage))
			res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "WhiteSpace")
			res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", "models."+rectanchoredtext.WhiteSpace.ToCodeString())
		}
	case "X_Offset":
		res = NumberInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", rectanchoredtext.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "X_Offset")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%f", rectanchoredtext.X_Offset))
	case "Y_Offset":
		res = NumberInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", rectanchoredtext.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "Y_Offset")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%f", rectanchoredtext.Y_Offset))
	case "RectAnchorType":
		if rectanchoredtext.RectAnchorType != "" {
			res = StringEnumInitStatement
			res = strings.ReplaceAll(res, "{{Identifier}}", rectanchoredtext.GongGetIdentifier(stage))
			res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "RectAnchorType")
			res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", "models."+rectanchoredtext.RectAnchorType.ToCodeString())
		}
	case "TextAnchorType":
		if rectanchoredtext.TextAnchorType != "" {
			res = StringEnumInitStatement
			res = strings.ReplaceAll(res, "{{Identifier}}", rectanchoredtext.GongGetIdentifier(stage))
			res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "TextAnchorType")
			res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", "models."+rectanchoredtext.TextAnchorType.ToCodeString())
		}
	case "DominantBaseline":
		if rectanchoredtext.DominantBaseline != "" {
			res = StringEnumInitStatement
			res = strings.ReplaceAll(res, "{{Identifier}}", rectanchoredtext.GongGetIdentifier(stage))
			res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "DominantBaseline")
			res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", "models."+rectanchoredtext.DominantBaseline.ToCodeString())
		}
	case "WritingMode":
		if rectanchoredtext.WritingMode != "" {
			res = StringEnumInitStatement
			res = strings.ReplaceAll(res, "{{Identifier}}", rectanchoredtext.GongGetIdentifier(stage))
			res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "WritingMode")
			res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", "models."+rectanchoredtext.WritingMode.ToCodeString())
		}
	case "Color":
		res = StringInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", rectanchoredtext.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "Color")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", string(rectanchoredtext.Color))
	case "FillOpacity":
		res = NumberInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", rectanchoredtext.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "FillOpacity")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%f", rectanchoredtext.FillOpacity))
	case "Stroke":
		res = StringInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", rectanchoredtext.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "Stroke")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", string(rectanchoredtext.Stroke))
	case "StrokeOpacity":
		res = NumberInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", rectanchoredtext.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "StrokeOpacity")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%f", rectanchoredtext.StrokeOpacity))
	case "StrokeWidth":
		res = NumberInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", rectanchoredtext.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "StrokeWidth")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%f", rectanchoredtext.StrokeWidth))
	case "StrokeDashArray":
		res = StringInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", rectanchoredtext.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "StrokeDashArray")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", string(rectanchoredtext.StrokeDashArray))
	case "StrokeDashArrayWhenSelected":
		res = StringInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", rectanchoredtext.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "StrokeDashArrayWhenSelected")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", string(rectanchoredtext.StrokeDashArrayWhenSelected))
	case "Transform":
		res = StringInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", rectanchoredtext.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "Transform")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", string(rectanchoredtext.Transform))

	case "Animates":
		for _, _animate := range rectanchoredtext.Animates {
			tmp := SliceOfPointersFieldInitStatement
			tmp = strings.ReplaceAll(tmp, "{{Identifier}}", rectanchoredtext.GongGetIdentifier(stage))
			tmp = strings.ReplaceAll(tmp, "{{GeneratedFieldName}}", "Animates")
			tmp = strings.ReplaceAll(tmp, "{{GeneratedFieldNameValue}}", _animate.GongGetIdentifier(stage))
			res += tmp
		}
	default:
		log.Panicf("Unknown field %s for Gongstruct RectAnchoredText", fieldName)
	}
	return
}

func (rectlinklink *RectLinkLink) GongMarshallField(stage *Stage, fieldName string) (res string) {

	switch fieldName {
	case "Name":
		res = StringInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", rectlinklink.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "Name")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", string(rectlinklink.Name))
	case "TargetAnchorPosition":
		res = NumberInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", rectlinklink.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "TargetAnchorPosition")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%f", rectlinklink.TargetAnchorPosition))
	case "Color":
		res = StringInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", rectlinklink.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "Color")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", string(rectlinklink.Color))
	case "FillOpacity":
		res = NumberInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", rectlinklink.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "FillOpacity")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%f", rectlinklink.FillOpacity))
	case "Stroke":
		res = StringInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", rectlinklink.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "Stroke")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", string(rectlinklink.Stroke))
	case "StrokeOpacity":
		res = NumberInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", rectlinklink.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "StrokeOpacity")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%f", rectlinklink.StrokeOpacity))
	case "StrokeWidth":
		res = NumberInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", rectlinklink.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "StrokeWidth")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%f", rectlinklink.StrokeWidth))
	case "StrokeDashArray":
		res = StringInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", rectlinklink.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "StrokeDashArray")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", string(rectlinklink.StrokeDashArray))
	case "StrokeDashArrayWhenSelected":
		res = StringInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", rectlinklink.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "StrokeDashArrayWhenSelected")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", string(rectlinklink.StrokeDashArrayWhenSelected))
	case "Transform":
		res = StringInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", rectlinklink.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "Transform")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", string(rectlinklink.Transform))

	case "Start":
		if rectlinklink.Start != nil {
			res = PointerFieldInitStatement
			res = strings.ReplaceAll(res, "{{Identifier}}", rectlinklink.GongGetIdentifier(stage))
			res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "Start")
			res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", rectlinklink.Start.GongGetIdentifier(stage))
		} else {
			// in case of nil pointer, we need to unstage the previous value
			res = PointerFieldInitStatement
			res = strings.ReplaceAll(res, "{{Identifier}}", rectlinklink.GongGetIdentifier(stage))
			res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "Start")
			res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", "nil")
		}
	case "End":
		if rectlinklink.End != nil {
			res = PointerFieldInitStatement
			res = strings.ReplaceAll(res, "{{Identifier}}", rectlinklink.GongGetIdentifier(stage))
			res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "End")
			res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", rectlinklink.End.GongGetIdentifier(stage))
		} else {
			// in case of nil pointer, we need to unstage the previous value
			res = PointerFieldInitStatement
			res = strings.ReplaceAll(res, "{{Identifier}}", rectlinklink.GongGetIdentifier(stage))
			res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "End")
			res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", "nil")
		}
	default:
		log.Panicf("Unknown field %s for Gongstruct RectLinkLink", fieldName)
	}
	return
}

func (svg *SVG) GongMarshallField(stage *Stage, fieldName string) (res string) {

	switch fieldName {
	case "Name":
		res = StringInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", svg.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "Name")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", string(svg.Name))
	case "DrawingState":
		if svg.DrawingState != "" {
			res = StringEnumInitStatement
			res = strings.ReplaceAll(res, "{{Identifier}}", svg.GongGetIdentifier(stage))
			res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "DrawingState")
			res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", "models."+svg.DrawingState.ToCodeString())
		}
	case "IsEditable":
		res = NumberInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", svg.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "IsEditable")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%t", svg.IsEditable))
	case "IsSVGFrontEndFileGenerated":
		res = NumberInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", svg.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "IsSVGFrontEndFileGenerated")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%t", svg.IsSVGFrontEndFileGenerated))
	case "IsSVGBackEndFileGenerated":
		res = NumberInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", svg.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "IsSVGBackEndFileGenerated")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%t", svg.IsSVGBackEndFileGenerated))
	case "DefaultDirectoryForGeneratedImages":
		res = StringInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", svg.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "DefaultDirectoryForGeneratedImages")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", string(svg.DefaultDirectoryForGeneratedImages))
	case "IsControlBannerHidden":
		res = NumberInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", svg.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "IsControlBannerHidden")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%t", svg.IsControlBannerHidden))
	case "OverrideWidth":
		res = NumberInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", svg.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "OverrideWidth")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%t", svg.OverrideWidth))
	case "OverriddenWidth":
		res = NumberInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", svg.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "OverriddenWidth")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%f", svg.OverriddenWidth))
	case "OverrideHeight":
		res = NumberInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", svg.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "OverrideHeight")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%t", svg.OverrideHeight))
	case "OverriddenHeight":
		res = NumberInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", svg.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "OverriddenHeight")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%f", svg.OverriddenHeight))

	case "Layers":
		for _, _layer := range svg.Layers {
			tmp := SliceOfPointersFieldInitStatement
			tmp = strings.ReplaceAll(tmp, "{{Identifier}}", svg.GongGetIdentifier(stage))
			tmp = strings.ReplaceAll(tmp, "{{GeneratedFieldName}}", "Layers")
			tmp = strings.ReplaceAll(tmp, "{{GeneratedFieldNameValue}}", _layer.GongGetIdentifier(stage))
			res += tmp
		}
	case "StartRect":
		if svg.StartRect != nil {
			res = PointerFieldInitStatement
			res = strings.ReplaceAll(res, "{{Identifier}}", svg.GongGetIdentifier(stage))
			res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "StartRect")
			res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", svg.StartRect.GongGetIdentifier(stage))
		} else {
			// in case of nil pointer, we need to unstage the previous value
			res = PointerFieldInitStatement
			res = strings.ReplaceAll(res, "{{Identifier}}", svg.GongGetIdentifier(stage))
			res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "StartRect")
			res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", "nil")
		}
	case "EndRect":
		if svg.EndRect != nil {
			res = PointerFieldInitStatement
			res = strings.ReplaceAll(res, "{{Identifier}}", svg.GongGetIdentifier(stage))
			res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "EndRect")
			res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", svg.EndRect.GongGetIdentifier(stage))
		} else {
			// in case of nil pointer, we need to unstage the previous value
			res = PointerFieldInitStatement
			res = strings.ReplaceAll(res, "{{Identifier}}", svg.GongGetIdentifier(stage))
			res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "EndRect")
			res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", "nil")
		}
	default:
		log.Panicf("Unknown field %s for Gongstruct SVG", fieldName)
	}
	return
}

func (svgtext *SvgText) GongMarshallField(stage *Stage, fieldName string) (res string) {

	switch fieldName {
	case "Name":
		res = StringInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", svgtext.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "Name")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", string(svgtext.Name))
	case "Text":
		res = StringInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", svgtext.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "Text")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", string(svgtext.Text))

	default:
		log.Panicf("Unknown field %s for Gongstruct SvgText", fieldName)
	}
	return
}

func (text *Text) GongMarshallField(stage *Stage, fieldName string) (res string) {

	switch fieldName {
	case "Name":
		res = StringInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", text.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "Name")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", string(text.Name))
	case "X":
		res = NumberInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", text.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "X")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%f", text.X))
	case "Y":
		res = NumberInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", text.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "Y")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%f", text.Y))
	case "Content":
		res = StringInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", text.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "Content")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", string(text.Content))
	case "Color":
		res = StringInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", text.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "Color")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", string(text.Color))
	case "FillOpacity":
		res = NumberInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", text.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "FillOpacity")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%f", text.FillOpacity))
	case "Stroke":
		res = StringInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", text.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "Stroke")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", string(text.Stroke))
	case "StrokeOpacity":
		res = NumberInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", text.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "StrokeOpacity")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%f", text.StrokeOpacity))
	case "StrokeWidth":
		res = NumberInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", text.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "StrokeWidth")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%f", text.StrokeWidth))
	case "StrokeDashArray":
		res = StringInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", text.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "StrokeDashArray")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", string(text.StrokeDashArray))
	case "StrokeDashArrayWhenSelected":
		res = StringInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", text.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "StrokeDashArrayWhenSelected")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", string(text.StrokeDashArrayWhenSelected))
	case "Transform":
		res = StringInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", text.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "Transform")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", string(text.Transform))
	case "FontWeight":
		res = StringInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", text.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "FontWeight")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", string(text.FontWeight))
	case "FontSize":
		res = StringInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", text.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "FontSize")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", string(text.FontSize))
	case "FontStyle":
		res = StringInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", text.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "FontStyle")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", string(text.FontStyle))
	case "LetterSpacing":
		res = StringInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", text.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "LetterSpacing")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", string(text.LetterSpacing))
	case "FontFamily":
		res = StringInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", text.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "FontFamily")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", string(text.FontFamily))
	case "WhiteSpace":
		if text.WhiteSpace != "" {
			res = StringEnumInitStatement
			res = strings.ReplaceAll(res, "{{Identifier}}", text.GongGetIdentifier(stage))
			res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "WhiteSpace")
			res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", "models."+text.WhiteSpace.ToCodeString())
		}

	case "Animates":
		for _, _animate := range text.Animates {
			tmp := SliceOfPointersFieldInitStatement
			tmp = strings.ReplaceAll(tmp, "{{Identifier}}", text.GongGetIdentifier(stage))
			tmp = strings.ReplaceAll(tmp, "{{GeneratedFieldName}}", "Animates")
			tmp = strings.ReplaceAll(tmp, "{{GeneratedFieldNameValue}}", _animate.GongGetIdentifier(stage))
			res += tmp
		}
	default:
		log.Panicf("Unknown field %s for Gongstruct Text", fieldName)
	}
	return
}

// insertion point for marshall all fields methods
func (animate *Animate) GongMarshallAllFields(stage *Stage) (initializerStatements string, pointersInitializesStatements string) {

	{ // Insertion point for basic fields value assignment
		initializerStatements += animate.GongMarshallField(stage, "Name")
		initializerStatements += animate.GongMarshallField(stage, "AttributeName")
		initializerStatements += animate.GongMarshallField(stage, "Values")
		initializerStatements += animate.GongMarshallField(stage, "From")
		initializerStatements += animate.GongMarshallField(stage, "To")
		initializerStatements += animate.GongMarshallField(stage, "Dur")
		initializerStatements += animate.GongMarshallField(stage, "RepeatCount")
	}
	return
}
func (circle *Circle) GongMarshallAllFields(stage *Stage) (initializerStatements string, pointersInitializesStatements string) {

	{ // Insertion point for basic fields value assignment
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
	return
}
func (condition *Condition) GongMarshallAllFields(stage *Stage) (initializerStatements string, pointersInitializesStatements string) {

	{ // Insertion point for basic fields value assignment
		initializerStatements += condition.GongMarshallField(stage, "Name")
	}
	return
}
func (controlpoint *ControlPoint) GongMarshallAllFields(stage *Stage) (initializerStatements string, pointersInitializesStatements string) {

	{ // Insertion point for basic fields value assignment
		initializerStatements += controlpoint.GongMarshallField(stage, "Name")
		initializerStatements += controlpoint.GongMarshallField(stage, "X_Relative")
		initializerStatements += controlpoint.GongMarshallField(stage, "Y_Relative")
		pointersInitializesStatements += controlpoint.GongMarshallField(stage, "ClosestRect")
	}
	return
}
func (ellipse *Ellipse) GongMarshallAllFields(stage *Stage) (initializerStatements string, pointersInitializesStatements string) {

	{ // Insertion point for basic fields value assignment
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
	return
}
func (layer *Layer) GongMarshallAllFields(stage *Stage) (initializerStatements string, pointersInitializesStatements string) {

	{ // Insertion point for basic fields value assignment
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
	return
}
func (line *Line) GongMarshallAllFields(stage *Stage) (initializerStatements string, pointersInitializesStatements string) {

	{ // Insertion point for basic fields value assignment
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
	return
}
func (link *Link) GongMarshallAllFields(stage *Stage) (initializerStatements string, pointersInitializesStatements string) {

	{ // Insertion point for basic fields value assignment
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
	return
}
func (linkanchoredtext *LinkAnchoredText) GongMarshallAllFields(stage *Stage) (initializerStatements string, pointersInitializesStatements string) {

	{ // Insertion point for basic fields value assignment
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
	return
}
func (path *Path) GongMarshallAllFields(stage *Stage) (initializerStatements string, pointersInitializesStatements string) {

	{ // Insertion point for basic fields value assignment
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
	return
}
func (point *Point) GongMarshallAllFields(stage *Stage) (initializerStatements string, pointersInitializesStatements string) {

	{ // Insertion point for basic fields value assignment
		initializerStatements += point.GongMarshallField(stage, "Name")
		initializerStatements += point.GongMarshallField(stage, "X")
		initializerStatements += point.GongMarshallField(stage, "Y")
	}
	return
}
func (polygone *Polygone) GongMarshallAllFields(stage *Stage) (initializerStatements string, pointersInitializesStatements string) {

	{ // Insertion point for basic fields value assignment
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
	return
}
func (polyline *Polyline) GongMarshallAllFields(stage *Stage) (initializerStatements string, pointersInitializesStatements string) {

	{ // Insertion point for basic fields value assignment
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
	return
}
func (rect *Rect) GongMarshallAllFields(stage *Stage) (initializerStatements string, pointersInitializesStatements string) {

	{ // Insertion point for basic fields value assignment
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
	return
}
func (rectanchoredpath *RectAnchoredPath) GongMarshallAllFields(stage *Stage) (initializerStatements string, pointersInitializesStatements string) {

	{ // Insertion point for basic fields value assignment
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
	return
}
func (rectanchoredrect *RectAnchoredRect) GongMarshallAllFields(stage *Stage) (initializerStatements string, pointersInitializesStatements string) {

	{ // Insertion point for basic fields value assignment
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
	return
}
func (rectanchoredtext *RectAnchoredText) GongMarshallAllFields(stage *Stage) (initializerStatements string, pointersInitializesStatements string) {

	{ // Insertion point for basic fields value assignment
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
	return
}
func (rectlinklink *RectLinkLink) GongMarshallAllFields(stage *Stage) (initializerStatements string, pointersInitializesStatements string) {

	{ // Insertion point for basic fields value assignment
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
	return
}
func (svg *SVG) GongMarshallAllFields(stage *Stage) (initializerStatements string, pointersInitializesStatements string) {

	{ // Insertion point for basic fields value assignment
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
	return
}
func (svgtext *SvgText) GongMarshallAllFields(stage *Stage) (initializerStatements string, pointersInitializesStatements string) {

	{ // Insertion point for basic fields value assignment
		initializerStatements += svgtext.GongMarshallField(stage, "Name")
		initializerStatements += svgtext.GongMarshallField(stage, "Text")
	}
	return
}
func (text *Text) GongMarshallAllFields(stage *Stage) (initializerStatements string, pointersInitializesStatements string) {

	{ // Insertion point for basic fields value assignment
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
	return
}
