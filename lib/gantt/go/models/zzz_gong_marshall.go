// generated code - do not edit
package models

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"regexp"
	"sort"
	"strings"
	"time"
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
}`

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
		log.Fatalln(name + " is not a go filename")
	}

	log.Printf("%s Marshalling %s", time.Now().Format("2006-01-02 15:04:05.000000"), name)
	newBase := filepath.Base(file.Name())

	res := marshallRes
	res = strings.ReplaceAll(res, "{{databaseName}}", strings.ReplaceAll(newBase, ".go", ""))
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
	map_Arrow_Identifiers := make(map[*Arrow]string)
	_ = map_Arrow_Identifiers

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
	for idx, arrow := range arrowOrdered {

		id = generatesIdentifier("Arrow", idx, arrow.Name)
		map_Arrow_Identifiers[arrow] = id

		decl = IdentifiersDecls
		decl = strings.ReplaceAll(decl, "{{Identifier}}", id)
		decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "Arrow")
		decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", arrow.Name)
		identifiersDecl += decl

		initializerStatements += "\n"
		// Initialisation of values
		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Name")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(arrow.Name))
		initializerStatements += setValueField

		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "OptionnalColor")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(arrow.OptionnalColor))
		initializerStatements += setValueField

		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "OptionnalStroke")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(arrow.OptionnalStroke))
		initializerStatements += setValueField

	}

	map_Bar_Identifiers := make(map[*Bar]string)
	_ = map_Bar_Identifiers

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
	for idx, bar := range barOrdered {

		id = generatesIdentifier("Bar", idx, bar.Name)
		map_Bar_Identifiers[bar] = id

		decl = IdentifiersDecls
		decl = strings.ReplaceAll(decl, "{{Identifier}}", id)
		decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "Bar")
		decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", bar.Name)
		identifiersDecl += decl

		initializerStatements += "\n"
		// Initialisation of values
		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Name")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(bar.Name))
		initializerStatements += setValueField

		setValueField = TimeInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Start")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", bar.Start.String())
		initializerStatements += setValueField

		setValueField = TimeInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "End")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", bar.End.String())
		initializerStatements += setValueField

		setValueField = NumberInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "ComputedDuration")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%d", bar.ComputedDuration))
		initializerStatements += setValueField

		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "OptionnalColor")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(bar.OptionnalColor))
		initializerStatements += setValueField

		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "OptionnalStroke")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(bar.OptionnalStroke))
		initializerStatements += setValueField

		setValueField = NumberInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "FillOpacity")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%f", bar.FillOpacity))
		initializerStatements += setValueField

		setValueField = NumberInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "StrokeWidth")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%f", bar.StrokeWidth))
		initializerStatements += setValueField

		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "StrokeDashArray")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(bar.StrokeDashArray))
		initializerStatements += setValueField

	}

	map_Gantt_Identifiers := make(map[*Gantt]string)
	_ = map_Gantt_Identifiers

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
	for idx, gantt := range ganttOrdered {

		id = generatesIdentifier("Gantt", idx, gantt.Name)
		map_Gantt_Identifiers[gantt] = id

		decl = IdentifiersDecls
		decl = strings.ReplaceAll(decl, "{{Identifier}}", id)
		decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "Gantt")
		decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", gantt.Name)
		identifiersDecl += decl

		initializerStatements += "\n"
		// Initialisation of values
		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Name")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(gantt.Name))
		initializerStatements += setValueField

		setValueField = TimeInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "ComputedStart")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", gantt.ComputedStart.String())
		initializerStatements += setValueField

		setValueField = TimeInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "ComputedEnd")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", gantt.ComputedEnd.String())
		initializerStatements += setValueField

		setValueField = NumberInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "ComputedDuration")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%d", gantt.ComputedDuration))
		initializerStatements += setValueField

		setValueField = NumberInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "UseManualStartAndEndDates")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%t", gantt.UseManualStartAndEndDates))
		initializerStatements += setValueField

		setValueField = TimeInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "ManualStart")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", gantt.ManualStart.String())
		initializerStatements += setValueField

		setValueField = TimeInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "ManualEnd")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", gantt.ManualEnd.String())
		initializerStatements += setValueField

		setValueField = NumberInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "LaneHeight")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%f", gantt.LaneHeight))
		initializerStatements += setValueField

		setValueField = NumberInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "RatioBarToLaneHeight")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%f", gantt.RatioBarToLaneHeight))
		initializerStatements += setValueField

		setValueField = NumberInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "YTopMargin")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%f", gantt.YTopMargin))
		initializerStatements += setValueField

		setValueField = NumberInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "XLeftText")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%f", gantt.XLeftText))
		initializerStatements += setValueField

		setValueField = NumberInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "TextHeight")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%f", gantt.TextHeight))
		initializerStatements += setValueField

		setValueField = NumberInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "XLeftLanes")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%f", gantt.XLeftLanes))
		initializerStatements += setValueField

		setValueField = NumberInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "XRightMargin")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%f", gantt.XRightMargin))
		initializerStatements += setValueField

		setValueField = NumberInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "ArrowLengthToTheRightOfStartBar")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%f", gantt.ArrowLengthToTheRightOfStartBar))
		initializerStatements += setValueField

		setValueField = NumberInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "ArrowTipLenght")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%f", gantt.ArrowTipLenght))
		initializerStatements += setValueField

		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "TimeLine_Color")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(gantt.TimeLine_Color))
		initializerStatements += setValueField

		setValueField = NumberInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "TimeLine_FillOpacity")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%f", gantt.TimeLine_FillOpacity))
		initializerStatements += setValueField

		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "TimeLine_Stroke")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(gantt.TimeLine_Stroke))
		initializerStatements += setValueField

		setValueField = NumberInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "TimeLine_StrokeWidth")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%f", gantt.TimeLine_StrokeWidth))
		initializerStatements += setValueField

		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Group_Stroke")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(gantt.Group_Stroke))
		initializerStatements += setValueField

		setValueField = NumberInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Group_StrokeWidth")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%f", gantt.Group_StrokeWidth))
		initializerStatements += setValueField

		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Group_StrokeDashArray")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(gantt.Group_StrokeDashArray))
		initializerStatements += setValueField

		setValueField = NumberInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "DateYOffset")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%f", gantt.DateYOffset))
		initializerStatements += setValueField

		setValueField = NumberInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "AlignOnStartEndOnYearStart")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%t", gantt.AlignOnStartEndOnYearStart))
		initializerStatements += setValueField

	}

	map_Group_Identifiers := make(map[*Group]string)
	_ = map_Group_Identifiers

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
	for idx, group := range groupOrdered {

		id = generatesIdentifier("Group", idx, group.Name)
		map_Group_Identifiers[group] = id

		decl = IdentifiersDecls
		decl = strings.ReplaceAll(decl, "{{Identifier}}", id)
		decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "Group")
		decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", group.Name)
		identifiersDecl += decl

		initializerStatements += "\n"
		// Initialisation of values
		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Name")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(group.Name))
		initializerStatements += setValueField

	}

	map_Lane_Identifiers := make(map[*Lane]string)
	_ = map_Lane_Identifiers

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
	for idx, lane := range laneOrdered {

		id = generatesIdentifier("Lane", idx, lane.Name)
		map_Lane_Identifiers[lane] = id

		decl = IdentifiersDecls
		decl = strings.ReplaceAll(decl, "{{Identifier}}", id)
		decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "Lane")
		decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", lane.Name)
		identifiersDecl += decl

		initializerStatements += "\n"
		// Initialisation of values
		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Name")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(lane.Name))
		initializerStatements += setValueField

		setValueField = NumberInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Order")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%d", lane.Order))
		initializerStatements += setValueField

	}

	map_LaneUse_Identifiers := make(map[*LaneUse]string)
	_ = map_LaneUse_Identifiers

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
	for idx, laneuse := range laneuseOrdered {

		id = generatesIdentifier("LaneUse", idx, laneuse.Name)
		map_LaneUse_Identifiers[laneuse] = id

		decl = IdentifiersDecls
		decl = strings.ReplaceAll(decl, "{{Identifier}}", id)
		decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "LaneUse")
		decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", laneuse.Name)
		identifiersDecl += decl

		initializerStatements += "\n"
		// Initialisation of values
		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Name")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(laneuse.Name))
		initializerStatements += setValueField

	}

	map_Milestone_Identifiers := make(map[*Milestone]string)
	_ = map_Milestone_Identifiers

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
	for idx, milestone := range milestoneOrdered {

		id = generatesIdentifier("Milestone", idx, milestone.Name)
		map_Milestone_Identifiers[milestone] = id

		decl = IdentifiersDecls
		decl = strings.ReplaceAll(decl, "{{Identifier}}", id)
		decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "Milestone")
		decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", milestone.Name)
		identifiersDecl += decl

		initializerStatements += "\n"
		// Initialisation of values
		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Name")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(milestone.Name))
		initializerStatements += setValueField

		setValueField = TimeInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Date")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", milestone.Date.String())
		initializerStatements += setValueField

		setValueField = NumberInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "DisplayVerticalBar")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%t", milestone.DisplayVerticalBar))
		initializerStatements += setValueField

	}

	// insertion initialization of objects to stage
	if len(arrowOrdered) > 0 {
		pointersInitializesStatements += "\n\t// setup of Arrow instances pointers"
	}
	for idx, arrow := range arrowOrdered {
		var setPointerField string
		_ = setPointerField

		id = generatesIdentifier("Arrow", idx, arrow.Name)
		map_Arrow_Identifiers[arrow] = id

		// Initialisation of values
		if arrow.From != nil {
			setPointerField = PointerFieldInitStatement
			setPointerField = strings.ReplaceAll(setPointerField, "{{Identifier}}", id)
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldName}}", "From")
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldNameValue}}", map_Bar_Identifiers[arrow.From])
			pointersInitializesStatements += setPointerField
		}

		if arrow.To != nil {
			setPointerField = PointerFieldInitStatement
			setPointerField = strings.ReplaceAll(setPointerField, "{{Identifier}}", id)
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldName}}", "To")
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldNameValue}}", map_Bar_Identifiers[arrow.To])
			pointersInitializesStatements += setPointerField
		}

	}

	if len(barOrdered) > 0 {
		pointersInitializesStatements += "\n\t// setup of Bar instances pointers"
	}
	for idx, bar := range barOrdered {
		var setPointerField string
		_ = setPointerField

		id = generatesIdentifier("Bar", idx, bar.Name)
		map_Bar_Identifiers[bar] = id

		// Initialisation of values
	}

	if len(ganttOrdered) > 0 {
		pointersInitializesStatements += "\n\t// setup of Gantt instances pointers"
	}
	for idx, gantt := range ganttOrdered {
		var setPointerField string
		_ = setPointerField

		id = generatesIdentifier("Gantt", idx, gantt.Name)
		map_Gantt_Identifiers[gantt] = id

		// Initialisation of values
		for _, _lane := range gantt.Lanes {
			setPointerField = SliceOfPointersFieldInitStatement
			setPointerField = strings.ReplaceAll(setPointerField, "{{Identifier}}", id)
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldName}}", "Lanes")
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldNameValue}}", map_Lane_Identifiers[_lane])
			pointersInitializesStatements += setPointerField
		}

		for _, _milestone := range gantt.Milestones {
			setPointerField = SliceOfPointersFieldInitStatement
			setPointerField = strings.ReplaceAll(setPointerField, "{{Identifier}}", id)
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldName}}", "Milestones")
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldNameValue}}", map_Milestone_Identifiers[_milestone])
			pointersInitializesStatements += setPointerField
		}

		for _, _group := range gantt.Groups {
			setPointerField = SliceOfPointersFieldInitStatement
			setPointerField = strings.ReplaceAll(setPointerField, "{{Identifier}}", id)
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldName}}", "Groups")
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldNameValue}}", map_Group_Identifiers[_group])
			pointersInitializesStatements += setPointerField
		}

		for _, _arrow := range gantt.Arrows {
			setPointerField = SliceOfPointersFieldInitStatement
			setPointerField = strings.ReplaceAll(setPointerField, "{{Identifier}}", id)
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldName}}", "Arrows")
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldNameValue}}", map_Arrow_Identifiers[_arrow])
			pointersInitializesStatements += setPointerField
		}

	}

	if len(groupOrdered) > 0 {
		pointersInitializesStatements += "\n\t// setup of Group instances pointers"
	}
	for idx, group := range groupOrdered {
		var setPointerField string
		_ = setPointerField

		id = generatesIdentifier("Group", idx, group.Name)
		map_Group_Identifiers[group] = id

		// Initialisation of values
		for _, _lane := range group.GroupLanes {
			setPointerField = SliceOfPointersFieldInitStatement
			setPointerField = strings.ReplaceAll(setPointerField, "{{Identifier}}", id)
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldName}}", "GroupLanes")
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldNameValue}}", map_Lane_Identifiers[_lane])
			pointersInitializesStatements += setPointerField
		}

	}

	if len(laneOrdered) > 0 {
		pointersInitializesStatements += "\n\t// setup of Lane instances pointers"
	}
	for idx, lane := range laneOrdered {
		var setPointerField string
		_ = setPointerField

		id = generatesIdentifier("Lane", idx, lane.Name)
		map_Lane_Identifiers[lane] = id

		// Initialisation of values
		for _, _bar := range lane.Bars {
			setPointerField = SliceOfPointersFieldInitStatement
			setPointerField = strings.ReplaceAll(setPointerField, "{{Identifier}}", id)
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldName}}", "Bars")
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldNameValue}}", map_Bar_Identifiers[_bar])
			pointersInitializesStatements += setPointerField
		}

	}

	if len(laneuseOrdered) > 0 {
		pointersInitializesStatements += "\n\t// setup of LaneUse instances pointers"
	}
	for idx, laneuse := range laneuseOrdered {
		var setPointerField string
		_ = setPointerField

		id = generatesIdentifier("LaneUse", idx, laneuse.Name)
		map_LaneUse_Identifiers[laneuse] = id

		// Initialisation of values
		if laneuse.Lane != nil {
			setPointerField = PointerFieldInitStatement
			setPointerField = strings.ReplaceAll(setPointerField, "{{Identifier}}", id)
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldName}}", "Lane")
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldNameValue}}", map_Lane_Identifiers[laneuse.Lane])
			pointersInitializesStatements += setPointerField
		}

	}

	if len(milestoneOrdered) > 0 {
		pointersInitializesStatements += "\n\t// setup of Milestone instances pointers"
	}
	for idx, milestone := range milestoneOrdered {
		var setPointerField string
		_ = setPointerField

		id = generatesIdentifier("Milestone", idx, milestone.Name)
		map_Milestone_Identifiers[milestone] = id

		// Initialisation of values
		for _, _lane := range milestone.LanesToDisplay {
			setPointerField = SliceOfPointersFieldInitStatement
			setPointerField = strings.ReplaceAll(setPointerField, "{{Identifier}}", id)
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldName}}", "LanesToDisplay")
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldNameValue}}", map_Lane_Identifiers[_lane])
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

	fmt.Fprintln(file, res)
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
