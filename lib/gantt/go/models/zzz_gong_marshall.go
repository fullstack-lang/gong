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
	arrowOrdered := []*Arrow{}
	for arrow := range stage.Arrows {
		arrowOrdered = append(arrowOrdered, arrow)
	}
	sort.Slice(arrowOrdered[:], func(i, j int) bool {
		arrowi := arrowOrdered[i]
		arrowj := arrowOrdered[j]
		arrowi_order, oki := stage.ArrowMap_Staged_Order[arrowi]
		arrowj_order, okj := stage.ArrowMap_Staged_Order[arrowj]
		if !oki || !okj {
			log.Fatalln("unknown pointers")
		}
		return arrowi_order < arrowj_order
	})
	if len(arrowOrdered) > 0 {
		identifiersDecl += "\n"
	}
	for _, arrow := range arrowOrdered {

		identifiersDecl += arrow.GongMarshallIdentifier(stage)

		initializerStatements += "\n"
		// Insertion point for basic fields value assignment
		initializerStatements += arrow.GongMarshallField(stage, "Name")
		pointersInitializesStatements += arrow.GongMarshallField(stage, "From")
		pointersInitializesStatements += arrow.GongMarshallField(stage, "To")
		initializerStatements += arrow.GongMarshallField(stage, "OptionnalColor")
		initializerStatements += arrow.GongMarshallField(stage, "OptionnalStroke")
	}

	barOrdered := []*Bar{}
	for bar := range stage.Bars {
		barOrdered = append(barOrdered, bar)
	}
	sort.Slice(barOrdered[:], func(i, j int) bool {
		bari := barOrdered[i]
		barj := barOrdered[j]
		bari_order, oki := stage.BarMap_Staged_Order[bari]
		barj_order, okj := stage.BarMap_Staged_Order[barj]
		if !oki || !okj {
			log.Fatalln("unknown pointers")
		}
		return bari_order < barj_order
	})
	if len(barOrdered) > 0 {
		identifiersDecl += "\n"
	}
	for _, bar := range barOrdered {

		identifiersDecl += bar.GongMarshallIdentifier(stage)

		initializerStatements += "\n"
		// Insertion point for basic fields value assignment
		initializerStatements += bar.GongMarshallField(stage, "Name")
		initializerStatements += bar.GongMarshallField(stage, "Start")
		initializerStatements += bar.GongMarshallField(stage, "End")
		initializerStatements += bar.GongMarshallField(stage, "ComputedDuration")
		initializerStatements += bar.GongMarshallField(stage, "OptionnalColor")
		initializerStatements += bar.GongMarshallField(stage, "OptionnalStroke")
		initializerStatements += bar.GongMarshallField(stage, "FillOpacity")
		initializerStatements += bar.GongMarshallField(stage, "StrokeWidth")
		initializerStatements += bar.GongMarshallField(stage, "StrokeDashArray")
	}

	ganttOrdered := []*Gantt{}
	for gantt := range stage.Gantts {
		ganttOrdered = append(ganttOrdered, gantt)
	}
	sort.Slice(ganttOrdered[:], func(i, j int) bool {
		gantti := ganttOrdered[i]
		ganttj := ganttOrdered[j]
		gantti_order, oki := stage.GanttMap_Staged_Order[gantti]
		ganttj_order, okj := stage.GanttMap_Staged_Order[ganttj]
		if !oki || !okj {
			log.Fatalln("unknown pointers")
		}
		return gantti_order < ganttj_order
	})
	if len(ganttOrdered) > 0 {
		identifiersDecl += "\n"
	}
	for _, gantt := range ganttOrdered {

		identifiersDecl += gantt.GongMarshallIdentifier(stage)

		initializerStatements += "\n"
		// Insertion point for basic fields value assignment
		initializerStatements += gantt.GongMarshallField(stage, "Name")
		initializerStatements += gantt.GongMarshallField(stage, "ComputedStart")
		initializerStatements += gantt.GongMarshallField(stage, "ComputedEnd")
		initializerStatements += gantt.GongMarshallField(stage, "ComputedDuration")
		initializerStatements += gantt.GongMarshallField(stage, "UseManualStartAndEndDates")
		initializerStatements += gantt.GongMarshallField(stage, "ManualStart")
		initializerStatements += gantt.GongMarshallField(stage, "ManualEnd")
		initializerStatements += gantt.GongMarshallField(stage, "LaneHeight")
		initializerStatements += gantt.GongMarshallField(stage, "RatioBarToLaneHeight")
		initializerStatements += gantt.GongMarshallField(stage, "YTopMargin")
		initializerStatements += gantt.GongMarshallField(stage, "XLeftText")
		initializerStatements += gantt.GongMarshallField(stage, "TextHeight")
		initializerStatements += gantt.GongMarshallField(stage, "XLeftLanes")
		initializerStatements += gantt.GongMarshallField(stage, "XRightMargin")
		initializerStatements += gantt.GongMarshallField(stage, "ArrowLengthToTheRightOfStartBar")
		initializerStatements += gantt.GongMarshallField(stage, "ArrowTipLenght")
		initializerStatements += gantt.GongMarshallField(stage, "TimeLine_Color")
		initializerStatements += gantt.GongMarshallField(stage, "TimeLine_FillOpacity")
		initializerStatements += gantt.GongMarshallField(stage, "TimeLine_Stroke")
		initializerStatements += gantt.GongMarshallField(stage, "TimeLine_StrokeWidth")
		initializerStatements += gantt.GongMarshallField(stage, "Group_Stroke")
		initializerStatements += gantt.GongMarshallField(stage, "Group_StrokeWidth")
		initializerStatements += gantt.GongMarshallField(stage, "Group_StrokeDashArray")
		initializerStatements += gantt.GongMarshallField(stage, "DateYOffset")
		initializerStatements += gantt.GongMarshallField(stage, "AlignOnStartEndOnYearStart")
		pointersInitializesStatements += gantt.GongMarshallField(stage, "Lanes")
		pointersInitializesStatements += gantt.GongMarshallField(stage, "Milestones")
		pointersInitializesStatements += gantt.GongMarshallField(stage, "Groups")
		pointersInitializesStatements += gantt.GongMarshallField(stage, "Arrows")
	}

	groupOrdered := []*Group{}
	for group := range stage.Groups {
		groupOrdered = append(groupOrdered, group)
	}
	sort.Slice(groupOrdered[:], func(i, j int) bool {
		groupi := groupOrdered[i]
		groupj := groupOrdered[j]
		groupi_order, oki := stage.GroupMap_Staged_Order[groupi]
		groupj_order, okj := stage.GroupMap_Staged_Order[groupj]
		if !oki || !okj {
			log.Fatalln("unknown pointers")
		}
		return groupi_order < groupj_order
	})
	if len(groupOrdered) > 0 {
		identifiersDecl += "\n"
	}
	for _, group := range groupOrdered {

		identifiersDecl += group.GongMarshallIdentifier(stage)

		initializerStatements += "\n"
		// Insertion point for basic fields value assignment
		initializerStatements += group.GongMarshallField(stage, "Name")
		pointersInitializesStatements += group.GongMarshallField(stage, "GroupLanes")
	}

	laneOrdered := []*Lane{}
	for lane := range stage.Lanes {
		laneOrdered = append(laneOrdered, lane)
	}
	sort.Slice(laneOrdered[:], func(i, j int) bool {
		lanei := laneOrdered[i]
		lanej := laneOrdered[j]
		lanei_order, oki := stage.LaneMap_Staged_Order[lanei]
		lanej_order, okj := stage.LaneMap_Staged_Order[lanej]
		if !oki || !okj {
			log.Fatalln("unknown pointers")
		}
		return lanei_order < lanej_order
	})
	if len(laneOrdered) > 0 {
		identifiersDecl += "\n"
	}
	for _, lane := range laneOrdered {

		identifiersDecl += lane.GongMarshallIdentifier(stage)

		initializerStatements += "\n"
		// Insertion point for basic fields value assignment
		initializerStatements += lane.GongMarshallField(stage, "Name")
		initializerStatements += lane.GongMarshallField(stage, "Order")
		pointersInitializesStatements += lane.GongMarshallField(stage, "Bars")
	}

	laneuseOrdered := []*LaneUse{}
	for laneuse := range stage.LaneUses {
		laneuseOrdered = append(laneuseOrdered, laneuse)
	}
	sort.Slice(laneuseOrdered[:], func(i, j int) bool {
		laneusei := laneuseOrdered[i]
		laneusej := laneuseOrdered[j]
		laneusei_order, oki := stage.LaneUseMap_Staged_Order[laneusei]
		laneusej_order, okj := stage.LaneUseMap_Staged_Order[laneusej]
		if !oki || !okj {
			log.Fatalln("unknown pointers")
		}
		return laneusei_order < laneusej_order
	})
	if len(laneuseOrdered) > 0 {
		identifiersDecl += "\n"
	}
	for _, laneuse := range laneuseOrdered {

		identifiersDecl += laneuse.GongMarshallIdentifier(stage)

		initializerStatements += "\n"
		// Insertion point for basic fields value assignment
		initializerStatements += laneuse.GongMarshallField(stage, "Name")
		pointersInitializesStatements += laneuse.GongMarshallField(stage, "Lane")
	}

	milestoneOrdered := []*Milestone{}
	for milestone := range stage.Milestones {
		milestoneOrdered = append(milestoneOrdered, milestone)
	}
	sort.Slice(milestoneOrdered[:], func(i, j int) bool {
		milestonei := milestoneOrdered[i]
		milestonej := milestoneOrdered[j]
		milestonei_order, oki := stage.MilestoneMap_Staged_Order[milestonei]
		milestonej_order, okj := stage.MilestoneMap_Staged_Order[milestonej]
		if !oki || !okj {
			log.Fatalln("unknown pointers")
		}
		return milestonei_order < milestonej_order
	})
	if len(milestoneOrdered) > 0 {
		identifiersDecl += "\n"
	}
	for _, milestone := range milestoneOrdered {

		identifiersDecl += milestone.GongMarshallIdentifier(stage)

		initializerStatements += "\n"
		// Insertion point for basic fields value assignment
		initializerStatements += milestone.GongMarshallField(stage, "Name")
		initializerStatements += milestone.GongMarshallField(stage, "Date")
		initializerStatements += milestone.GongMarshallField(stage, "DisplayVerticalBar")
		pointersInitializesStatements += milestone.GongMarshallField(stage, "LanesToDisplay")
	}

	// insertion initialization of objects to stage
	for _, arrow := range arrowOrdered {
		_ = arrow
		var setPointerField string
		_ = setPointerField

		// Insertion point for pointers initialization
	}

	for _, bar := range barOrdered {
		_ = bar
		var setPointerField string
		_ = setPointerField

		// Insertion point for pointers initialization
	}

	for _, gantt := range ganttOrdered {
		_ = gantt
		var setPointerField string
		_ = setPointerField

		// Insertion point for pointers initialization
	}

	for _, group := range groupOrdered {
		_ = group
		var setPointerField string
		_ = setPointerField

		// Insertion point for pointers initialization
	}

	for _, lane := range laneOrdered {
		_ = lane
		var setPointerField string
		_ = setPointerField

		// Insertion point for pointers initialization
	}

	for _, laneuse := range laneuseOrdered {
		_ = laneuse
		var setPointerField string
		_ = setPointerField

		// Insertion point for pointers initialization
	}

	for _, milestone := range milestoneOrdered {
		_ = milestone
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
func (arrow *Arrow) GongMarshallField(stage *Stage, fieldName string) (res string) {
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
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", arrow.GongGetIdentifier(stage))
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Name")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(arrow.Name))
		initializerStatements += setValueField
	case "OptionnalColor":
		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", arrow.GongGetIdentifier(stage))
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "OptionnalColor")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(arrow.OptionnalColor))
		initializerStatements += setValueField
	case "OptionnalStroke":
		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", arrow.GongGetIdentifier(stage))
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "OptionnalStroke")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(arrow.OptionnalStroke))
		initializerStatements += setValueField

	case "From":
		if arrow.From != nil {
			setPointerField = PointerFieldInitStatement
			setPointerField = strings.ReplaceAll(setPointerField, "{{Identifier}}", arrow.GongGetIdentifier(stage))
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldName}}", "From")
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldNameValue}}", arrow.From.GongGetIdentifier(stage))
			pointersInitializesStatements += setPointerField
		}
	case "To":
		if arrow.To != nil {
			setPointerField = PointerFieldInitStatement
			setPointerField = strings.ReplaceAll(setPointerField, "{{Identifier}}", arrow.GongGetIdentifier(stage))
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldName}}", "To")
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldNameValue}}", arrow.To.GongGetIdentifier(stage))
			pointersInitializesStatements += setPointerField
		}
	default:
		log.Panicf("Unknown field %s for Gongstruct Arrow", fieldName)
	}

	// temporary kludge to reuse existing template code
	res = initializerStatements + pointersInitializesStatements
	return
}

func (bar *Bar) GongMarshallField(stage *Stage, fieldName string) (res string) {
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
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", bar.GongGetIdentifier(stage))
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Name")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(bar.Name))
		initializerStatements += setValueField
	case "Start":
		setValueField = TimeInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", bar.GongGetIdentifier(stage))
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Start")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", bar.Start.String())
		initializerStatements += setValueField
	case "End":
		setValueField = TimeInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", bar.GongGetIdentifier(stage))
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "End")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", bar.End.String())
		initializerStatements += setValueField
	case "ComputedDuration":
		setValueField = NumberInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", bar.GongGetIdentifier(stage))
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "ComputedDuration")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%d", bar.ComputedDuration))
		initializerStatements += setValueField
	case "OptionnalColor":
		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", bar.GongGetIdentifier(stage))
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "OptionnalColor")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(bar.OptionnalColor))
		initializerStatements += setValueField
	case "OptionnalStroke":
		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", bar.GongGetIdentifier(stage))
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "OptionnalStroke")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(bar.OptionnalStroke))
		initializerStatements += setValueField
	case "FillOpacity":
		setValueField = NumberInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", bar.GongGetIdentifier(stage))
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "FillOpacity")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%f", bar.FillOpacity))
		initializerStatements += setValueField
	case "StrokeWidth":
		setValueField = NumberInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", bar.GongGetIdentifier(stage))
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "StrokeWidth")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%f", bar.StrokeWidth))
		initializerStatements += setValueField
	case "StrokeDashArray":
		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", bar.GongGetIdentifier(stage))
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "StrokeDashArray")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(bar.StrokeDashArray))
		initializerStatements += setValueField

	default:
		log.Panicf("Unknown field %s for Gongstruct Bar", fieldName)
	}

	// temporary kludge to reuse existing template code
	res = initializerStatements + pointersInitializesStatements
	return
}

func (gantt *Gantt) GongMarshallField(stage *Stage, fieldName string) (res string) {
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
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", gantt.GongGetIdentifier(stage))
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Name")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(gantt.Name))
		initializerStatements += setValueField
	case "ComputedStart":
		setValueField = TimeInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", gantt.GongGetIdentifier(stage))
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "ComputedStart")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", gantt.ComputedStart.String())
		initializerStatements += setValueField
	case "ComputedEnd":
		setValueField = TimeInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", gantt.GongGetIdentifier(stage))
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "ComputedEnd")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", gantt.ComputedEnd.String())
		initializerStatements += setValueField
	case "ComputedDuration":
		setValueField = NumberInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", gantt.GongGetIdentifier(stage))
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "ComputedDuration")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%d", gantt.ComputedDuration))
		initializerStatements += setValueField
	case "UseManualStartAndEndDates":
		setValueField = NumberInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", gantt.GongGetIdentifier(stage))
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "UseManualStartAndEndDates")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%t", gantt.UseManualStartAndEndDates))
		initializerStatements += setValueField
	case "ManualStart":
		setValueField = TimeInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", gantt.GongGetIdentifier(stage))
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "ManualStart")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", gantt.ManualStart.String())
		initializerStatements += setValueField
	case "ManualEnd":
		setValueField = TimeInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", gantt.GongGetIdentifier(stage))
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "ManualEnd")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", gantt.ManualEnd.String())
		initializerStatements += setValueField
	case "LaneHeight":
		setValueField = NumberInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", gantt.GongGetIdentifier(stage))
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "LaneHeight")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%f", gantt.LaneHeight))
		initializerStatements += setValueField
	case "RatioBarToLaneHeight":
		setValueField = NumberInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", gantt.GongGetIdentifier(stage))
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "RatioBarToLaneHeight")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%f", gantt.RatioBarToLaneHeight))
		initializerStatements += setValueField
	case "YTopMargin":
		setValueField = NumberInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", gantt.GongGetIdentifier(stage))
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "YTopMargin")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%f", gantt.YTopMargin))
		initializerStatements += setValueField
	case "XLeftText":
		setValueField = NumberInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", gantt.GongGetIdentifier(stage))
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "XLeftText")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%f", gantt.XLeftText))
		initializerStatements += setValueField
	case "TextHeight":
		setValueField = NumberInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", gantt.GongGetIdentifier(stage))
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "TextHeight")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%f", gantt.TextHeight))
		initializerStatements += setValueField
	case "XLeftLanes":
		setValueField = NumberInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", gantt.GongGetIdentifier(stage))
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "XLeftLanes")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%f", gantt.XLeftLanes))
		initializerStatements += setValueField
	case "XRightMargin":
		setValueField = NumberInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", gantt.GongGetIdentifier(stage))
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "XRightMargin")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%f", gantt.XRightMargin))
		initializerStatements += setValueField
	case "ArrowLengthToTheRightOfStartBar":
		setValueField = NumberInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", gantt.GongGetIdentifier(stage))
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "ArrowLengthToTheRightOfStartBar")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%f", gantt.ArrowLengthToTheRightOfStartBar))
		initializerStatements += setValueField
	case "ArrowTipLenght":
		setValueField = NumberInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", gantt.GongGetIdentifier(stage))
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "ArrowTipLenght")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%f", gantt.ArrowTipLenght))
		initializerStatements += setValueField
	case "TimeLine_Color":
		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", gantt.GongGetIdentifier(stage))
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "TimeLine_Color")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(gantt.TimeLine_Color))
		initializerStatements += setValueField
	case "TimeLine_FillOpacity":
		setValueField = NumberInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", gantt.GongGetIdentifier(stage))
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "TimeLine_FillOpacity")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%f", gantt.TimeLine_FillOpacity))
		initializerStatements += setValueField
	case "TimeLine_Stroke":
		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", gantt.GongGetIdentifier(stage))
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "TimeLine_Stroke")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(gantt.TimeLine_Stroke))
		initializerStatements += setValueField
	case "TimeLine_StrokeWidth":
		setValueField = NumberInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", gantt.GongGetIdentifier(stage))
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "TimeLine_StrokeWidth")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%f", gantt.TimeLine_StrokeWidth))
		initializerStatements += setValueField
	case "Group_Stroke":
		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", gantt.GongGetIdentifier(stage))
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Group_Stroke")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(gantt.Group_Stroke))
		initializerStatements += setValueField
	case "Group_StrokeWidth":
		setValueField = NumberInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", gantt.GongGetIdentifier(stage))
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Group_StrokeWidth")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%f", gantt.Group_StrokeWidth))
		initializerStatements += setValueField
	case "Group_StrokeDashArray":
		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", gantt.GongGetIdentifier(stage))
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Group_StrokeDashArray")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(gantt.Group_StrokeDashArray))
		initializerStatements += setValueField
	case "DateYOffset":
		setValueField = NumberInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", gantt.GongGetIdentifier(stage))
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "DateYOffset")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%f", gantt.DateYOffset))
		initializerStatements += setValueField
	case "AlignOnStartEndOnYearStart":
		setValueField = NumberInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", gantt.GongGetIdentifier(stage))
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "AlignOnStartEndOnYearStart")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%t", gantt.AlignOnStartEndOnYearStart))
		initializerStatements += setValueField

	case "Lanes":
		for _, _lane := range gantt.Lanes {
			setPointerField = SliceOfPointersFieldInitStatement
			setPointerField = strings.ReplaceAll(setPointerField, "{{Identifier}}", gantt.GongGetIdentifier(stage))
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldName}}", "Lanes")
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldNameValue}}", _lane.GongGetIdentifier(stage))
			pointersInitializesStatements += setPointerField
		}
	case "Milestones":
		for _, _milestone := range gantt.Milestones {
			setPointerField = SliceOfPointersFieldInitStatement
			setPointerField = strings.ReplaceAll(setPointerField, "{{Identifier}}", gantt.GongGetIdentifier(stage))
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldName}}", "Milestones")
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldNameValue}}", _milestone.GongGetIdentifier(stage))
			pointersInitializesStatements += setPointerField
		}
	case "Groups":
		for _, _group := range gantt.Groups {
			setPointerField = SliceOfPointersFieldInitStatement
			setPointerField = strings.ReplaceAll(setPointerField, "{{Identifier}}", gantt.GongGetIdentifier(stage))
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldName}}", "Groups")
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldNameValue}}", _group.GongGetIdentifier(stage))
			pointersInitializesStatements += setPointerField
		}
	case "Arrows":
		for _, _arrow := range gantt.Arrows {
			setPointerField = SliceOfPointersFieldInitStatement
			setPointerField = strings.ReplaceAll(setPointerField, "{{Identifier}}", gantt.GongGetIdentifier(stage))
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldName}}", "Arrows")
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldNameValue}}", _arrow.GongGetIdentifier(stage))
			pointersInitializesStatements += setPointerField
		}
	default:
		log.Panicf("Unknown field %s for Gongstruct Gantt", fieldName)
	}

	// temporary kludge to reuse existing template code
	res = initializerStatements + pointersInitializesStatements
	return
}

func (group *Group) GongMarshallField(stage *Stage, fieldName string) (res string) {
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
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", group.GongGetIdentifier(stage))
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Name")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(group.Name))
		initializerStatements += setValueField

	case "GroupLanes":
		for _, _lane := range group.GroupLanes {
			setPointerField = SliceOfPointersFieldInitStatement
			setPointerField = strings.ReplaceAll(setPointerField, "{{Identifier}}", group.GongGetIdentifier(stage))
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldName}}", "GroupLanes")
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldNameValue}}", _lane.GongGetIdentifier(stage))
			pointersInitializesStatements += setPointerField
		}
	default:
		log.Panicf("Unknown field %s for Gongstruct Group", fieldName)
	}

	// temporary kludge to reuse existing template code
	res = initializerStatements + pointersInitializesStatements
	return
}

func (lane *Lane) GongMarshallField(stage *Stage, fieldName string) (res string) {
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
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", lane.GongGetIdentifier(stage))
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Name")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(lane.Name))
		initializerStatements += setValueField
	case "Order":
		setValueField = NumberInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", lane.GongGetIdentifier(stage))
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Order")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%d", lane.Order))
		initializerStatements += setValueField

	case "Bars":
		for _, _bar := range lane.Bars {
			setPointerField = SliceOfPointersFieldInitStatement
			setPointerField = strings.ReplaceAll(setPointerField, "{{Identifier}}", lane.GongGetIdentifier(stage))
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldName}}", "Bars")
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldNameValue}}", _bar.GongGetIdentifier(stage))
			pointersInitializesStatements += setPointerField
		}
	default:
		log.Panicf("Unknown field %s for Gongstruct Lane", fieldName)
	}

	// temporary kludge to reuse existing template code
	res = initializerStatements + pointersInitializesStatements
	return
}

func (laneuse *LaneUse) GongMarshallField(stage *Stage, fieldName string) (res string) {
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
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", laneuse.GongGetIdentifier(stage))
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Name")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(laneuse.Name))
		initializerStatements += setValueField

	case "Lane":
		if laneuse.Lane != nil {
			setPointerField = PointerFieldInitStatement
			setPointerField = strings.ReplaceAll(setPointerField, "{{Identifier}}", laneuse.GongGetIdentifier(stage))
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldName}}", "Lane")
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldNameValue}}", laneuse.Lane.GongGetIdentifier(stage))
			pointersInitializesStatements += setPointerField
		}
	default:
		log.Panicf("Unknown field %s for Gongstruct LaneUse", fieldName)
	}

	// temporary kludge to reuse existing template code
	res = initializerStatements + pointersInitializesStatements
	return
}

func (milestone *Milestone) GongMarshallField(stage *Stage, fieldName string) (res string) {
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
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", milestone.GongGetIdentifier(stage))
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Name")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(milestone.Name))
		initializerStatements += setValueField
	case "Date":
		setValueField = TimeInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", milestone.GongGetIdentifier(stage))
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Date")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", milestone.Date.String())
		initializerStatements += setValueField
	case "DisplayVerticalBar":
		setValueField = NumberInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", milestone.GongGetIdentifier(stage))
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "DisplayVerticalBar")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%t", milestone.DisplayVerticalBar))
		initializerStatements += setValueField

	case "LanesToDisplay":
		for _, _lane := range milestone.LanesToDisplay {
			setPointerField = SliceOfPointersFieldInitStatement
			setPointerField = strings.ReplaceAll(setPointerField, "{{Identifier}}", milestone.GongGetIdentifier(stage))
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldName}}", "LanesToDisplay")
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldNameValue}}", _lane.GongGetIdentifier(stage))
			pointersInitializesStatements += setPointerField
		}
	default:
		log.Panicf("Unknown field %s for Gongstruct Milestone", fieldName)
	}

	// temporary kludge to reuse existing template code
	res = initializerStatements + pointersInitializesStatements
	return
}
