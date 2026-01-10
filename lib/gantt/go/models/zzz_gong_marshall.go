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

// insertion point for marshall field methods
func (arrow *Arrow) GongMarshallField(stage *Stage, fieldName string) (res string) {

	switch fieldName {
	case "Name":
		res = StringInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", arrow.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "Name")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", string(arrow.Name))
	case "OptionnalColor":
		res = StringInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", arrow.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "OptionnalColor")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", string(arrow.OptionnalColor))
	case "OptionnalStroke":
		res = StringInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", arrow.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "OptionnalStroke")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", string(arrow.OptionnalStroke))

	case "From":
		if arrow.From != nil {
			res = PointerFieldInitStatement
			res = strings.ReplaceAll(res, "{{Identifier}}", arrow.GongGetIdentifier(stage))
			res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "From")
			res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", arrow.From.GongGetIdentifier(stage))
		} else {
			// in case of nil pointer, we need to unstage the previous value
			res = PointerFieldInitStatement
			res = strings.ReplaceAll(res, "{{Identifier}}", arrow.GongGetIdentifier(stage))
			res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "From")
			res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", "nil")
		}
	case "To":
		if arrow.To != nil {
			res = PointerFieldInitStatement
			res = strings.ReplaceAll(res, "{{Identifier}}", arrow.GongGetIdentifier(stage))
			res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "To")
			res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", arrow.To.GongGetIdentifier(stage))
		} else {
			// in case of nil pointer, we need to unstage the previous value
			res = PointerFieldInitStatement
			res = strings.ReplaceAll(res, "{{Identifier}}", arrow.GongGetIdentifier(stage))
			res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "To")
			res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", "nil")
		}
	default:
		log.Panicf("Unknown field %s for Gongstruct Arrow", fieldName)
	}
	return
}

func (bar *Bar) GongMarshallField(stage *Stage, fieldName string) (res string) {

	switch fieldName {
	case "Name":
		res = StringInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", bar.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "Name")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", string(bar.Name))
	case "Start":
		res = TimeInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", bar.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "Start")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", bar.Start.String())
	case "End":
		res = TimeInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", bar.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "End")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", bar.End.String())
	case "ComputedDuration":
		res = NumberInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", bar.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "ComputedDuration")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%d", bar.ComputedDuration))
	case "OptionnalColor":
		res = StringInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", bar.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "OptionnalColor")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", string(bar.OptionnalColor))
	case "OptionnalStroke":
		res = StringInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", bar.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "OptionnalStroke")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", string(bar.OptionnalStroke))
	case "FillOpacity":
		res = NumberInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", bar.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "FillOpacity")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%f", bar.FillOpacity))
	case "StrokeWidth":
		res = NumberInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", bar.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "StrokeWidth")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%f", bar.StrokeWidth))
	case "StrokeDashArray":
		res = StringInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", bar.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "StrokeDashArray")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", string(bar.StrokeDashArray))

	default:
		log.Panicf("Unknown field %s for Gongstruct Bar", fieldName)
	}
	return
}

func (gantt *Gantt) GongMarshallField(stage *Stage, fieldName string) (res string) {

	switch fieldName {
	case "Name":
		res = StringInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", gantt.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "Name")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", string(gantt.Name))
	case "ComputedStart":
		res = TimeInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", gantt.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "ComputedStart")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", gantt.ComputedStart.String())
	case "ComputedEnd":
		res = TimeInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", gantt.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "ComputedEnd")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", gantt.ComputedEnd.String())
	case "ComputedDuration":
		res = NumberInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", gantt.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "ComputedDuration")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%d", gantt.ComputedDuration))
	case "UseManualStartAndEndDates":
		res = NumberInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", gantt.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "UseManualStartAndEndDates")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%t", gantt.UseManualStartAndEndDates))
	case "ManualStart":
		res = TimeInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", gantt.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "ManualStart")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", gantt.ManualStart.String())
	case "ManualEnd":
		res = TimeInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", gantt.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "ManualEnd")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", gantt.ManualEnd.String())
	case "LaneHeight":
		res = NumberInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", gantt.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "LaneHeight")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%f", gantt.LaneHeight))
	case "RatioBarToLaneHeight":
		res = NumberInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", gantt.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "RatioBarToLaneHeight")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%f", gantt.RatioBarToLaneHeight))
	case "YTopMargin":
		res = NumberInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", gantt.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "YTopMargin")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%f", gantt.YTopMargin))
	case "XLeftText":
		res = NumberInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", gantt.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "XLeftText")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%f", gantt.XLeftText))
	case "TextHeight":
		res = NumberInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", gantt.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "TextHeight")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%f", gantt.TextHeight))
	case "XLeftLanes":
		res = NumberInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", gantt.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "XLeftLanes")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%f", gantt.XLeftLanes))
	case "XRightMargin":
		res = NumberInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", gantt.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "XRightMargin")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%f", gantt.XRightMargin))
	case "ArrowLengthToTheRightOfStartBar":
		res = NumberInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", gantt.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "ArrowLengthToTheRightOfStartBar")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%f", gantt.ArrowLengthToTheRightOfStartBar))
	case "ArrowTipLenght":
		res = NumberInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", gantt.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "ArrowTipLenght")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%f", gantt.ArrowTipLenght))
	case "TimeLine_Color":
		res = StringInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", gantt.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "TimeLine_Color")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", string(gantt.TimeLine_Color))
	case "TimeLine_FillOpacity":
		res = NumberInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", gantt.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "TimeLine_FillOpacity")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%f", gantt.TimeLine_FillOpacity))
	case "TimeLine_Stroke":
		res = StringInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", gantt.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "TimeLine_Stroke")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", string(gantt.TimeLine_Stroke))
	case "TimeLine_StrokeWidth":
		res = NumberInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", gantt.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "TimeLine_StrokeWidth")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%f", gantt.TimeLine_StrokeWidth))
	case "Group_Stroke":
		res = StringInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", gantt.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "Group_Stroke")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", string(gantt.Group_Stroke))
	case "Group_StrokeWidth":
		res = NumberInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", gantt.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "Group_StrokeWidth")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%f", gantt.Group_StrokeWidth))
	case "Group_StrokeDashArray":
		res = StringInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", gantt.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "Group_StrokeDashArray")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", string(gantt.Group_StrokeDashArray))
	case "DateYOffset":
		res = NumberInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", gantt.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "DateYOffset")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%f", gantt.DateYOffset))
	case "AlignOnStartEndOnYearStart":
		res = NumberInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", gantt.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "AlignOnStartEndOnYearStart")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%t", gantt.AlignOnStartEndOnYearStart))

	case "Lanes":
		for _, _lane := range gantt.Lanes {
			tmp := SliceOfPointersFieldInitStatement
			tmp = strings.ReplaceAll(tmp, "{{Identifier}}", gantt.GongGetIdentifier(stage))
			tmp = strings.ReplaceAll(tmp, "{{GeneratedFieldName}}", "Lanes")
			tmp = strings.ReplaceAll(tmp, "{{GeneratedFieldNameValue}}", _lane.GongGetIdentifier(stage))
			res += tmp
		}
	case "Milestones":
		for _, _milestone := range gantt.Milestones {
			tmp := SliceOfPointersFieldInitStatement
			tmp = strings.ReplaceAll(tmp, "{{Identifier}}", gantt.GongGetIdentifier(stage))
			tmp = strings.ReplaceAll(tmp, "{{GeneratedFieldName}}", "Milestones")
			tmp = strings.ReplaceAll(tmp, "{{GeneratedFieldNameValue}}", _milestone.GongGetIdentifier(stage))
			res += tmp
		}
	case "Groups":
		for _, _group := range gantt.Groups {
			tmp := SliceOfPointersFieldInitStatement
			tmp = strings.ReplaceAll(tmp, "{{Identifier}}", gantt.GongGetIdentifier(stage))
			tmp = strings.ReplaceAll(tmp, "{{GeneratedFieldName}}", "Groups")
			tmp = strings.ReplaceAll(tmp, "{{GeneratedFieldNameValue}}", _group.GongGetIdentifier(stage))
			res += tmp
		}
	case "Arrows":
		for _, _arrow := range gantt.Arrows {
			tmp := SliceOfPointersFieldInitStatement
			tmp = strings.ReplaceAll(tmp, "{{Identifier}}", gantt.GongGetIdentifier(stage))
			tmp = strings.ReplaceAll(tmp, "{{GeneratedFieldName}}", "Arrows")
			tmp = strings.ReplaceAll(tmp, "{{GeneratedFieldNameValue}}", _arrow.GongGetIdentifier(stage))
			res += tmp
		}
	default:
		log.Panicf("Unknown field %s for Gongstruct Gantt", fieldName)
	}
	return
}

func (group *Group) GongMarshallField(stage *Stage, fieldName string) (res string) {

	switch fieldName {
	case "Name":
		res = StringInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", group.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "Name")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", string(group.Name))

	case "GroupLanes":
		for _, _lane := range group.GroupLanes {
			tmp := SliceOfPointersFieldInitStatement
			tmp = strings.ReplaceAll(tmp, "{{Identifier}}", group.GongGetIdentifier(stage))
			tmp = strings.ReplaceAll(tmp, "{{GeneratedFieldName}}", "GroupLanes")
			tmp = strings.ReplaceAll(tmp, "{{GeneratedFieldNameValue}}", _lane.GongGetIdentifier(stage))
			res += tmp
		}
	default:
		log.Panicf("Unknown field %s for Gongstruct Group", fieldName)
	}
	return
}

func (lane *Lane) GongMarshallField(stage *Stage, fieldName string) (res string) {

	switch fieldName {
	case "Name":
		res = StringInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", lane.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "Name")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", string(lane.Name))
	case "Order":
		res = NumberInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", lane.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "Order")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%d", lane.Order))

	case "Bars":
		for _, _bar := range lane.Bars {
			tmp := SliceOfPointersFieldInitStatement
			tmp = strings.ReplaceAll(tmp, "{{Identifier}}", lane.GongGetIdentifier(stage))
			tmp = strings.ReplaceAll(tmp, "{{GeneratedFieldName}}", "Bars")
			tmp = strings.ReplaceAll(tmp, "{{GeneratedFieldNameValue}}", _bar.GongGetIdentifier(stage))
			res += tmp
		}
	default:
		log.Panicf("Unknown field %s for Gongstruct Lane", fieldName)
	}
	return
}

func (laneuse *LaneUse) GongMarshallField(stage *Stage, fieldName string) (res string) {

	switch fieldName {
	case "Name":
		res = StringInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", laneuse.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "Name")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", string(laneuse.Name))

	case "Lane":
		if laneuse.Lane != nil {
			res = PointerFieldInitStatement
			res = strings.ReplaceAll(res, "{{Identifier}}", laneuse.GongGetIdentifier(stage))
			res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "Lane")
			res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", laneuse.Lane.GongGetIdentifier(stage))
		} else {
			// in case of nil pointer, we need to unstage the previous value
			res = PointerFieldInitStatement
			res = strings.ReplaceAll(res, "{{Identifier}}", laneuse.GongGetIdentifier(stage))
			res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "Lane")
			res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", "nil")
		}
	default:
		log.Panicf("Unknown field %s for Gongstruct LaneUse", fieldName)
	}
	return
}

func (milestone *Milestone) GongMarshallField(stage *Stage, fieldName string) (res string) {

	switch fieldName {
	case "Name":
		res = StringInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", milestone.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "Name")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", string(milestone.Name))
	case "Date":
		res = TimeInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", milestone.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "Date")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", milestone.Date.String())
	case "DisplayVerticalBar":
		res = NumberInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", milestone.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "DisplayVerticalBar")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%t", milestone.DisplayVerticalBar))

	case "LanesToDisplay":
		for _, _lane := range milestone.LanesToDisplay {
			tmp := SliceOfPointersFieldInitStatement
			tmp = strings.ReplaceAll(tmp, "{{Identifier}}", milestone.GongGetIdentifier(stage))
			tmp = strings.ReplaceAll(tmp, "{{GeneratedFieldName}}", "LanesToDisplay")
			tmp = strings.ReplaceAll(tmp, "{{GeneratedFieldNameValue}}", _lane.GongGetIdentifier(stage))
			res += tmp
		}
	default:
		log.Panicf("Unknown field %s for Gongstruct Milestone", fieldName)
	}
	return
}

// insertion point for marshall all fields methods
func (arrow *Arrow) GongMarshallAllFields(stage *Stage) (initializerStatements string, pointersInitializesStatements string) {

	{ // Insertion point for basic fields value assignment
		initializerStatements += arrow.GongMarshallField(stage, "Name")
		pointersInitializesStatements += arrow.GongMarshallField(stage, "From")
		pointersInitializesStatements += arrow.GongMarshallField(stage, "To")
		initializerStatements += arrow.GongMarshallField(stage, "OptionnalColor")
		initializerStatements += arrow.GongMarshallField(stage, "OptionnalStroke")
	}
	return
}
func (bar *Bar) GongMarshallAllFields(stage *Stage) (initializerStatements string, pointersInitializesStatements string) {

	{ // Insertion point for basic fields value assignment
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
	return
}
func (gantt *Gantt) GongMarshallAllFields(stage *Stage) (initializerStatements string, pointersInitializesStatements string) {

	{ // Insertion point for basic fields value assignment
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
	return
}
func (group *Group) GongMarshallAllFields(stage *Stage) (initializerStatements string, pointersInitializesStatements string) {

	{ // Insertion point for basic fields value assignment
		initializerStatements += group.GongMarshallField(stage, "Name")
		pointersInitializesStatements += group.GongMarshallField(stage, "GroupLanes")
	}
	return
}
func (lane *Lane) GongMarshallAllFields(stage *Stage) (initializerStatements string, pointersInitializesStatements string) {

	{ // Insertion point for basic fields value assignment
		initializerStatements += lane.GongMarshallField(stage, "Name")
		initializerStatements += lane.GongMarshallField(stage, "Order")
		pointersInitializesStatements += lane.GongMarshallField(stage, "Bars")
	}
	return
}
func (laneuse *LaneUse) GongMarshallAllFields(stage *Stage) (initializerStatements string, pointersInitializesStatements string) {

	{ // Insertion point for basic fields value assignment
		initializerStatements += laneuse.GongMarshallField(stage, "Name")
		pointersInitializesStatements += laneuse.GongMarshallField(stage, "Lane")
	}
	return
}
func (milestone *Milestone) GongMarshallAllFields(stage *Stage) (initializerStatements string, pointersInitializesStatements string) {

	{ // Insertion point for basic fields value assignment
		initializerStatements += milestone.GongMarshallField(stage, "Name")
		initializerStatements += milestone.GongMarshallField(stage, "Date")
		initializerStatements += milestone.GongMarshallField(stage, "DisplayVerticalBar")
		pointersInitializesStatements += milestone.GongMarshallField(stage, "LanesToDisplay")
	}
	return
}
