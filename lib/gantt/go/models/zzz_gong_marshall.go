// generated code - do not edit
package models

import (
	"fmt"
	"log"
	"os"
	"runtime/debug"
	"sort"
	"strings"
	"time"
)

var (
	_ = sort.Slice
	_ = time.Hour
)

const GongIdentifiersDecls = `
	{{Identifier}} := (&models.{{GeneratedStructName}}{Name: {{GeneratedFieldNameValue}}}).Stage(stage)`

const GongUnstageStmt = `
	{{Identifier}}.Unstage(stage)`


const GongStringInitStatement = `
	{{Identifier}}.{{GeneratedFieldName}} = {{GeneratedFieldNameValue}}`

const GongMetaFieldStructInitStatement = `
	{{Identifier}}.{{GeneratedFieldName}} = ` + `{{GeneratedFieldNameValue}}`

const GongStringEnumInitStatement = `
	{{Identifier}}.{{GeneratedFieldName}} = {{GeneratedFieldNameValue}}`

const GongNumberInitStatement = `
	{{Identifier}}.{{GeneratedFieldName}} = {{GeneratedFieldNameValue}}`

const GongPointerFieldInitStatement = `
	{{Identifier}}.{{GeneratedFieldName}} = {{GeneratedFieldNameValue}}`

const GongSliceOfPointersFieldInitStatement = `
	{{Identifier}}.{{GeneratedFieldName}} = append({{Identifier}}.{{GeneratedFieldName}}, {{GeneratedFieldNameValue}})`

const GongTimeInitStatement = `
	{{Identifier}}.{{GeneratedFieldName}}, _ = time.Parse("2006-01-02 15:04:05.999999999 -0700 MST", "{{GeneratedFieldNameValue}}")`

// __gong__toRawStringLiteral formats a string into safe Go source code,
// using backticks to preserve newlines and readability.
func __gong__toRawStringLiteral(s string) string {
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

func __gong__marshallString[T ~string](ident, fieldName string, val T) string {
	res := strings.ReplaceAll(GongStringInitStatement, "{{Identifier}}", ident)
	res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", fieldName)
	return strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", __gong__toRawStringLiteral(string(val)))
}

func __gong__marshallInt[T ~int | ~int8 | ~int16 | ~int32 | ~int64 | ~uint | ~uint8 | ~uint16 | ~uint32 | ~uint64](ident, fieldName string, val T) string {
	res := strings.ReplaceAll(GongNumberInitStatement, "{{Identifier}}", ident)
	res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", fieldName)
	return strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%d", val))
}

func __gong__marshallBool[T ~bool](ident, fieldName string, val T) string {
	res := strings.ReplaceAll(GongNumberInitStatement, "{{Identifier}}", ident)
	res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", fieldName)
	return strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%t", val))
}

func __gong__marshallFloat[T ~float32 | ~float64](ident, fieldName string, val T) string {
	res := strings.ReplaceAll(GongNumberInitStatement, "{{Identifier}}", ident)
	res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", fieldName)
	return strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%f", val))
}

func __gong__marshallTime(ident, fieldName, valStr string) string {
	res := strings.ReplaceAll(GongTimeInitStatement, "{{Identifier}}", ident)
	res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", fieldName)
	return strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", valStr)
}

func __gong__marshallPointer(ident, fieldName, targetIdent string) string {
	res := strings.ReplaceAll(GongPointerFieldInitStatement, "{{Identifier}}", ident)
	res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", fieldName)
	return strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", targetIdent)
}

func __gong__marshallSliceOfPointers(ident, fieldName, targetIdent string) string {
	res := strings.ReplaceAll(GongSliceOfPointersFieldInitStatement, "{{Identifier}}", ident)
	res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", fieldName)
	return strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", targetIdent)
}

func __gong__marshallEnumString(ident, fieldName, codeStr string) string {
	val := "\"\""
	if codeStr != "" {
		val = "models." + codeStr
	}
	res := strings.ReplaceAll(GongStringEnumInitStatement, "{{Identifier}}", ident)
	res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", fieldName)
	return strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", val)
}

func __gong__marshallEnumInt(ident, fieldName, codeStr string) string {
	val := "0"
	if codeStr != "" {
		val = "models." + codeStr
	}
	res := strings.ReplaceAll(GongNumberInitStatement, "{{Identifier}}", ident)
	res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", fieldName)
	return strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", val)
}

func __gong__marshallMeta(ident, fieldName, val string) string {
	res := strings.ReplaceAll(GongMetaFieldStructInitStatement, "{{Identifier}}", ident)
	res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", fieldName)
	return strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", val)
}

// __gong__sortStageSetInstances sorts instances by their staged order
func __gong__sortStageSetInstances[T comparable](instances map[T]struct{}, orderMap map[T]uint) []T {
	ordered := make([]T, 0, len(instances))
	for inst := range instances {
		ordered = append(ordered, inst)
	}
	sort.Slice(ordered, func(i, j int) bool {
		return orderMap[ordered[i]] < orderMap[ordered[j]]
	})
	return ordered
}

// MarshallFile marshall the stage content into a file as an instanciation into a stage
// according to the marshalling policy of the stage.
//
// In GongMarshallingAppendCommit mode, it will append the last commit to the file.
// In other modes, it will rewrite the entire file.
func (stage *Stage) MarshallFile(filename string, args ...string) {

	if stage.GetGongMarshallingMode() == GongMarshallingAppendCommit {
		contentBytes, err := os.ReadFile(filename)

		// if the file does not exist, marshall the full stage
		if os.IsNotExist(err) {
			file, createErr := os.Create(filename)
			if createErr != nil {
				log.Fatal(createErr.Error())
			}
			defer file.Close()
			stage.Marshall(file, args...)
			return
		}
		if err != nil {
			log.Fatal(err.Error())
		}

		content := string(contentBytes)

		if stage.IsSquashing() {
			// we squash: we want to clear the current function body
			// and let the append logic write the squashed commit
			const fnSig = "func _(stage *models.Stage) {"
			if before, _, ok := strings.Cut(content, fnSig); ok {
				content = before + fnSig + "\n}\n"
			}
		}

		if stage.IsApplyingBackwardCommit() {
			// we are going backward, we need to remove the last forward commit from the file

			// because commitsBehind has been incremented before the call to this function
			// the index of the commit to remove is len(forwardCommits) - commitsBehind
			forwardCommits := stage.GetForwardCommits()
			commitIndexToRemove := len(forwardCommits) - stage.GetCommitsBehind()

			if commitIndexToRemove < 0 || commitIndexToRemove >= len(forwardCommits) {
				return // Should not happen if history is consistent
			}

			commitToRemove := forwardCommits[commitIndexToRemove]

			lastIndex := strings.LastIndex(content, commitToRemove+"\n")
			if lastIndex != -1 {
				newContent := content[:lastIndex] + content[lastIndex+len(commitToRemove)+1:]
				err = os.WriteFile(filename, []byte(newContent), 0644)
				if err != nil {
					log.Fatal(err.Error())
				}
			} else {
				lastIndex = strings.LastIndex(content, commitToRemove)
				if lastIndex != -1 {
					newContent := content[:lastIndex] + content[lastIndex+len(commitToRemove):]
					err = os.WriteFile(filename, []byte(newContent), 0644)
					if err != nil {
						log.Fatal(err.Error())
					}
				} else {
					// The commit block was not found. This typically happens for the initial
					// commit which is formatted differently (the lines after func _(stage *models.Stage) {).
					// We rewrite the entire file with the current (rewound) stage state to safely remove it.
					file, createErr := os.Create(filename)
					if createErr != nil {
						log.Fatal(createErr.Error())
					}
					defer file.Close()
					stage.Marshall(file, args...)
				}
			}
			return // we are done for the backward case
		}

		if stage.IsApplyingForwardCommit() {
			// bypass the modified check
		} else if !stage.IsModified() {
			return
		}

		forwardCommits := stage.GetForwardCommits()
		if len(forwardCommits) == 0 {
			return // nothing to do
		}

		activeCommits := len(forwardCommits) - stage.GetCommitsBehind()
		if activeCommits <= 0 {
			return
		}
		forwardCommit := forwardCommits[activeCommits-1]

		// append before the ending brace of the func
		lastBrace := strings.LastIndex(content, "}")
		if lastBrace == -1 {
			// if no brace, something is wrong with the file, so we rewrite it
			file, createErr := os.Create(filename)
			if createErr != nil {
				log.Fatal(createErr.Error())
			}
			defer file.Close()
			stage.Marshall(file, args...)
			return
		}

		contentBeforeBrace := content[:lastBrace]
		trimmedContentBeforeBrace := strings.TrimSpace(contentBeforeBrace)
		emptyBody := stage.IsSquashing() ||
			strings.HasSuffix(trimmedContentBeforeBrace, "func _(stage *models.Stage) {") ||
			strings.HasSuffix(trimmedContentBeforeBrace, "// insertion point for setup of pointers") ||
			strings.HasSuffix(trimmedContentBeforeBrace, "// Phase 3: Pointer Setups (Intra-stage and Cross-stage pointers)\n\t// ------------------------------------------------------------------------")

		// check if the file ends with stage.Commit() before the brace
		if !emptyBody && !strings.HasSuffix(trimmedContentBeforeBrace, "stage.Commit()") {
			contentBeforeBrace = contentBeforeBrace + "\n\tstage.Commit()\n"
		}

		// insert the commit statements before the last brace
		newContent := contentBeforeBrace + forwardCommit + "\n" + content[lastBrace:]

		err = os.WriteFile(filename, []byte(newContent), 0644)
		if err != nil {
			log.Fatal(err.Error())
		}

	} else {
		file, err := os.Create(filename)
		if err != nil {
			log.Fatal(err.Error())
		}
		defer file.Close()

		stage.Marshall(file, args...)
	}
}

// Marshall marshall the stage content into the file as an instanciation into a stage
func (stage *Stage) Marshall(file *os.File, args ...string) {

	name := file.Name()

	if !strings.HasSuffix(name, ".go") {
		log.Println(name + " is not a go filename")
	}

	log.Printf("Marshalling %s", name)

	res, err := stage.MarshallToString(args...)
	if err != nil {
		log.Fatalln("Error marshalling to string:", err)
	}

	fmt.Fprintln(file, res)
}

// MarshallToString marshall the stage content into a string
func (stage *Stage) MarshallToString(args ...string) (res string, err error) {
	packageName := "main"
	modelsPackageName := "github.com/fullstack-lang/gong/lib/gantt/go/models"
	if len(args) == 1 {
		packageName = args[0]
	} else if len(args) >= 2 {
		if args[0] != "" {
			modelsPackageName = args[0]
		}
		packageName = args[1]
	}

	var declarations strings.Builder
	var values strings.Builder
	var pointers strings.Builder
	var lastStageDecl string
	var lastStageVal string
	var lastStagePtr string
	_ = lastStageDecl
	_ = lastStageVal
	_ = lastStagePtr

	if stage != nil {
		for _, arrow := range __gong__sortStageSetInstances(stage.Arrows, stage.Arrow_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			arrowIdent := arrow.GongGetIdentifier(stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&models.Arrow{Name: %s}).Stage(stage)", arrowIdent, __gong__toRawStringLiteral(arrow.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", arrowIdent, __gong__toRawStringLiteral(arrow.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.OptionnalColor = %s", arrowIdent, __gong__toRawStringLiteral(arrow.OptionnalColor)))
			values.WriteString(fmt.Sprintf("\n\t%s.OptionnalStroke = %s", arrowIdent, __gong__toRawStringLiteral(arrow.OptionnalStroke)))
			if arrow.From != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := arrow.From.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.From = %s", arrowIdent, targetIdent))
			}
			if arrow.To != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := arrow.To.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.To = %s", arrowIdent, targetIdent))
			}
		}
	}
	if stage != nil {
		for _, bar := range __gong__sortStageSetInstances(stage.Bars, stage.Bar_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			barIdent := bar.GongGetIdentifier(stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&models.Bar{Name: %s}).Stage(stage)", barIdent, __gong__toRawStringLiteral(bar.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", barIdent, __gong__toRawStringLiteral(bar.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Start, _ = time.Parse(\"2006-01-02 15:04:05.999999999 -0700 MST\", \"%s\")", barIdent, bar.Start.String()))
			values.WriteString(fmt.Sprintf("\n\t%s.End, _ = time.Parse(\"2006-01-02 15:04:05.999999999 -0700 MST\", \"%s\")", barIdent, bar.End.String()))
			values.WriteString(fmt.Sprintf("\n\t%s.ComputedDuration = time.Duration(%d)", barIdent, int64(bar.ComputedDuration)))
			values.WriteString(fmt.Sprintf("\n\t%s.OptionnalColor = %s", barIdent, __gong__toRawStringLiteral(bar.OptionnalColor)))
			values.WriteString(fmt.Sprintf("\n\t%s.OptionnalStroke = %s", barIdent, __gong__toRawStringLiteral(bar.OptionnalStroke)))
			values.WriteString(fmt.Sprintf("\n\t%s.FillOpacity = %f", barIdent, bar.FillOpacity))
			values.WriteString(fmt.Sprintf("\n\t%s.StrokeWidth = %f", barIdent, bar.StrokeWidth))
			values.WriteString(fmt.Sprintf("\n\t%s.StrokeDashArray = %s", barIdent, __gong__toRawStringLiteral(bar.StrokeDashArray)))
		}
	}
	if stage != nil {
		for _, gantt := range __gong__sortStageSetInstances(stage.Gantts, stage.Gantt_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			ganttIdent := gantt.GongGetIdentifier(stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&models.Gantt{Name: %s}).Stage(stage)", ganttIdent, __gong__toRawStringLiteral(gantt.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", ganttIdent, __gong__toRawStringLiteral(gantt.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.ComputedStart, _ = time.Parse(\"2006-01-02 15:04:05.999999999 -0700 MST\", \"%s\")", ganttIdent, gantt.ComputedStart.String()))
			values.WriteString(fmt.Sprintf("\n\t%s.ComputedEnd, _ = time.Parse(\"2006-01-02 15:04:05.999999999 -0700 MST\", \"%s\")", ganttIdent, gantt.ComputedEnd.String()))
			values.WriteString(fmt.Sprintf("\n\t%s.ComputedDuration = time.Duration(%d)", ganttIdent, int64(gantt.ComputedDuration)))
			values.WriteString(fmt.Sprintf("\n\t%s.UseManualStartAndEndDates = %t", ganttIdent, gantt.UseManualStartAndEndDates))
			values.WriteString(fmt.Sprintf("\n\t%s.ManualStart, _ = time.Parse(\"2006-01-02 15:04:05.999999999 -0700 MST\", \"%s\")", ganttIdent, gantt.ManualStart.String()))
			values.WriteString(fmt.Sprintf("\n\t%s.ManualEnd, _ = time.Parse(\"2006-01-02 15:04:05.999999999 -0700 MST\", \"%s\")", ganttIdent, gantt.ManualEnd.String()))
			values.WriteString(fmt.Sprintf("\n\t%s.LaneHeight = %f", ganttIdent, gantt.LaneHeight))
			values.WriteString(fmt.Sprintf("\n\t%s.RatioBarToLaneHeight = %f", ganttIdent, gantt.RatioBarToLaneHeight))
			values.WriteString(fmt.Sprintf("\n\t%s.YTopMargin = %f", ganttIdent, gantt.YTopMargin))
			values.WriteString(fmt.Sprintf("\n\t%s.XLeftText = %f", ganttIdent, gantt.XLeftText))
			values.WriteString(fmt.Sprintf("\n\t%s.TextHeight = %f", ganttIdent, gantt.TextHeight))
			values.WriteString(fmt.Sprintf("\n\t%s.XLeftLanes = %f", ganttIdent, gantt.XLeftLanes))
			values.WriteString(fmt.Sprintf("\n\t%s.XRightMargin = %f", ganttIdent, gantt.XRightMargin))
			values.WriteString(fmt.Sprintf("\n\t%s.ArrowLengthToTheRightOfStartBar = %f", ganttIdent, gantt.ArrowLengthToTheRightOfStartBar))
			values.WriteString(fmt.Sprintf("\n\t%s.ArrowTipLenght = %f", ganttIdent, gantt.ArrowTipLenght))
			values.WriteString(fmt.Sprintf("\n\t%s.TimeLine_Color = %s", ganttIdent, __gong__toRawStringLiteral(gantt.TimeLine_Color)))
			values.WriteString(fmt.Sprintf("\n\t%s.TimeLine_FillOpacity = %f", ganttIdent, gantt.TimeLine_FillOpacity))
			values.WriteString(fmt.Sprintf("\n\t%s.TimeLine_Stroke = %s", ganttIdent, __gong__toRawStringLiteral(gantt.TimeLine_Stroke)))
			values.WriteString(fmt.Sprintf("\n\t%s.TimeLine_StrokeWidth = %f", ganttIdent, gantt.TimeLine_StrokeWidth))
			values.WriteString(fmt.Sprintf("\n\t%s.Group_Stroke = %s", ganttIdent, __gong__toRawStringLiteral(gantt.Group_Stroke)))
			values.WriteString(fmt.Sprintf("\n\t%s.Group_StrokeWidth = %f", ganttIdent, gantt.Group_StrokeWidth))
			values.WriteString(fmt.Sprintf("\n\t%s.Group_StrokeDashArray = %s", ganttIdent, __gong__toRawStringLiteral(gantt.Group_StrokeDashArray)))
			values.WriteString(fmt.Sprintf("\n\t%s.DateYOffset = %f", ganttIdent, gantt.DateYOffset))
			values.WriteString(fmt.Sprintf("\n\t%s.AlignOnStartEndOnYearStart = %t", ganttIdent, gantt.AlignOnStartEndOnYearStart))
			for _, elem := range gantt.Lanes {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := elem.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Lanes = append(%s.Lanes, %s)", ganttIdent, ganttIdent, targetIdent))
			}
			for _, elem := range gantt.Milestones {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := elem.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Milestones = append(%s.Milestones, %s)", ganttIdent, ganttIdent, targetIdent))
			}
			for _, elem := range gantt.Groups {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := elem.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Groups = append(%s.Groups, %s)", ganttIdent, ganttIdent, targetIdent))
			}
			for _, elem := range gantt.Arrows {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := elem.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Arrows = append(%s.Arrows, %s)", ganttIdent, ganttIdent, targetIdent))
			}
		}
	}
	if stage != nil {
		for _, group := range __gong__sortStageSetInstances(stage.Groups, stage.Group_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			groupIdent := group.GongGetIdentifier(stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&models.Group{Name: %s}).Stage(stage)", groupIdent, __gong__toRawStringLiteral(group.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", groupIdent, __gong__toRawStringLiteral(group.Name)))
			for _, elem := range group.GroupLanes {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := elem.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.GroupLanes = append(%s.GroupLanes, %s)", groupIdent, groupIdent, targetIdent))
			}
		}
	}
	if stage != nil {
		for _, lane := range __gong__sortStageSetInstances(stage.Lanes, stage.Lane_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			laneIdent := lane.GongGetIdentifier(stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&models.Lane{Name: %s}).Stage(stage)", laneIdent, __gong__toRawStringLiteral(lane.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", laneIdent, __gong__toRawStringLiteral(lane.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Order = %d", laneIdent, lane.Order))
			for _, elem := range lane.Bars {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := elem.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Bars = append(%s.Bars, %s)", laneIdent, laneIdent, targetIdent))
			}
		}
	}
	if stage != nil {
		for _, laneuse := range __gong__sortStageSetInstances(stage.LaneUses, stage.LaneUse_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			laneuseIdent := laneuse.GongGetIdentifier(stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&models.LaneUse{Name: %s}).Stage(stage)", laneuseIdent, __gong__toRawStringLiteral(laneuse.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", laneuseIdent, __gong__toRawStringLiteral(laneuse.Name)))
			if laneuse.Lane != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := laneuse.Lane.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Lane = %s", laneuseIdent, targetIdent))
			}
		}
	}
	if stage != nil {
		for _, milestone := range __gong__sortStageSetInstances(stage.Milestones, stage.Milestone_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			milestoneIdent := milestone.GongGetIdentifier(stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&models.Milestone{Name: %s}).Stage(stage)", milestoneIdent, __gong__toRawStringLiteral(milestone.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", milestoneIdent, __gong__toRawStringLiteral(milestone.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Date, _ = time.Parse(\"2006-01-02 15:04:05.999999999 -0700 MST\", \"%s\")", milestoneIdent, milestone.Date.String()))
			values.WriteString(fmt.Sprintf("\n\t%s.DisplayVerticalBar = %t", milestoneIdent, milestone.DisplayVerticalBar))
			for _, elem := range milestone.LanesToDisplay {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := elem.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.LanesToDisplay = append(%s.LanesToDisplay, %s)", milestoneIdent, milestoneIdent, targetIdent))
			}
		}
	}

	var goModuleVersion string
	if buildInfo, ok := debug.ReadBuildInfo(); ok {
		goModuleVersion = buildInfo.Main.Version
	}

	goModuleVersionWithoutDirty := strings.ReplaceAll(goModuleVersion, "+dirty", "")
	if goModuleVersionWithoutDirty == "" || goModuleVersionWithoutDirty == "(devel)" {
		goModuleVersionWithoutDirty = "latest"
	}

	var metaImports strings.Builder
	var metaDummies strings.Builder
	if len(stage.MetaPackageImports) > 0 {
		for _, imp := range stage.MetaPackageImports {
			metaImports.WriteString(fmt.Sprintf("\n\t%s %s", imp.Alias, imp.Path))
			metaDummies.WriteString(fmt.Sprintf("\nvar _ %s.Stage", imp.Alias))
		}
	} else if stage.MetaPackageImportAlias != "" {
		metaImports.WriteString(fmt.Sprintf("\n\t%s %s", stage.MetaPackageImportAlias, stage.MetaPackageImportPath))
		metaDummies.WriteString(fmt.Sprintf("\nvar _ %s.Stage", stage.MetaPackageImportAlias))
	}

	res = fmt.Sprintf(`// file generated by gong
// go module version: %s
/* to edit this file with the default probe, use the command
go run github.com/fullstack-lang/gong/lib/gantt/go/models/probe/cmd/probe@%s stage.go
*/
package %s

import (
	"slices"
	"time"

	"github.com/fullstack-lang/gong/lib/gantt/go/models"%s
)

var (
	_ time.Time
	_ = slices.Index[[]int, int]

	_ *models.Stage
)%s

// function will stage objects across all coordinated stages
func _(stage *models.Stage) {

	// ------------------------------------------------------------------------
	// Phase 1: Declarations (in topological order: leaves first)
	// ------------------------------------------------------------------------%s

	// ------------------------------------------------------------------------
	// Phase 2: Value Initializations
	// ------------------------------------------------------------------------%s

	// ------------------------------------------------------------------------
	// Phase 3: Pointer Setups (Intra-stage and Cross-stage pointers)
	// ------------------------------------------------------------------------%s
}
`, goModuleVersion, goModuleVersionWithoutDirty, packageName, metaImports.String(), metaDummies.String(), declarations.String(), values.String(), pointers.String())

	if modelsPackageName != "" && modelsPackageName != "github.com/fullstack-lang/gong/lib/gantt/go/models" {
		res = strings.ReplaceAll(res, "\"github.com/fullstack-lang/gong/lib/gantt/go/models\"", fmt.Sprintf("%q", modelsPackageName))
	}

	return res, nil
}

// insertion point for marshall field methods
func (arrow *Arrow) GongMarshallField(stage *Stage, fieldName string) (res string) {
	ident := arrow.GongGetIdentifier(stage)
	_ = ident
	switch fieldName {
	case "Name":
		res = __gong__marshallString(ident, "Name", arrow.Name)
	case "OptionnalColor":
		res = __gong__marshallString(ident, "OptionnalColor", arrow.OptionnalColor)
	case "OptionnalStroke":
		res = __gong__marshallString(ident, "OptionnalStroke", arrow.OptionnalStroke)

	case "From":
		if arrow.From != nil {
			res = __gong__marshallPointer(ident, "From", arrow.From.GongGetIdentifier(stage))
		} else {
			res = __gong__marshallPointer(ident, "From", "nil")
		}
	case "To":
		if arrow.To != nil {
			res = __gong__marshallPointer(ident, "To", arrow.To.GongGetIdentifier(stage))
		} else {
			res = __gong__marshallPointer(ident, "To", "nil")
		}
	default:
		log.Panicf("Unknown field %s for Gongstruct Arrow", fieldName)
	}
	return
}

func (bar *Bar) GongMarshallField(stage *Stage, fieldName string) (res string) {
	ident := bar.GongGetIdentifier(stage)
	_ = ident
	switch fieldName {
	case "Name":
		res = __gong__marshallString(ident, "Name", bar.Name)
	case "Start":
		res = __gong__marshallTime(ident, "Start", bar.Start.String())
	case "End":
		res = __gong__marshallTime(ident, "End", bar.End.String())
	case "ComputedDuration":
		res = __gong__marshallInt(ident, "ComputedDuration", bar.ComputedDuration)
	case "OptionnalColor":
		res = __gong__marshallString(ident, "OptionnalColor", bar.OptionnalColor)
	case "OptionnalStroke":
		res = __gong__marshallString(ident, "OptionnalStroke", bar.OptionnalStroke)
	case "FillOpacity":
		res = __gong__marshallFloat(ident, "FillOpacity", bar.FillOpacity)
	case "StrokeWidth":
		res = __gong__marshallFloat(ident, "StrokeWidth", bar.StrokeWidth)
	case "StrokeDashArray":
		res = __gong__marshallString(ident, "StrokeDashArray", bar.StrokeDashArray)

	default:
		log.Panicf("Unknown field %s for Gongstruct Bar", fieldName)
	}
	return
}

func (gantt *Gantt) GongMarshallField(stage *Stage, fieldName string) (res string) {
	ident := gantt.GongGetIdentifier(stage)
	_ = ident
	switch fieldName {
	case "Name":
		res = __gong__marshallString(ident, "Name", gantt.Name)
	case "ComputedStart":
		res = __gong__marshallTime(ident, "ComputedStart", gantt.ComputedStart.String())
	case "ComputedEnd":
		res = __gong__marshallTime(ident, "ComputedEnd", gantt.ComputedEnd.String())
	case "ComputedDuration":
		res = __gong__marshallInt(ident, "ComputedDuration", gantt.ComputedDuration)
	case "UseManualStartAndEndDates":
		res = __gong__marshallBool(ident, "UseManualStartAndEndDates", gantt.UseManualStartAndEndDates)
	case "ManualStart":
		res = __gong__marshallTime(ident, "ManualStart", gantt.ManualStart.String())
	case "ManualEnd":
		res = __gong__marshallTime(ident, "ManualEnd", gantt.ManualEnd.String())
	case "LaneHeight":
		res = __gong__marshallFloat(ident, "LaneHeight", gantt.LaneHeight)
	case "RatioBarToLaneHeight":
		res = __gong__marshallFloat(ident, "RatioBarToLaneHeight", gantt.RatioBarToLaneHeight)
	case "YTopMargin":
		res = __gong__marshallFloat(ident, "YTopMargin", gantt.YTopMargin)
	case "XLeftText":
		res = __gong__marshallFloat(ident, "XLeftText", gantt.XLeftText)
	case "TextHeight":
		res = __gong__marshallFloat(ident, "TextHeight", gantt.TextHeight)
	case "XLeftLanes":
		res = __gong__marshallFloat(ident, "XLeftLanes", gantt.XLeftLanes)
	case "XRightMargin":
		res = __gong__marshallFloat(ident, "XRightMargin", gantt.XRightMargin)
	case "ArrowLengthToTheRightOfStartBar":
		res = __gong__marshallFloat(ident, "ArrowLengthToTheRightOfStartBar", gantt.ArrowLengthToTheRightOfStartBar)
	case "ArrowTipLenght":
		res = __gong__marshallFloat(ident, "ArrowTipLenght", gantt.ArrowTipLenght)
	case "TimeLine_Color":
		res = __gong__marshallString(ident, "TimeLine_Color", gantt.TimeLine_Color)
	case "TimeLine_FillOpacity":
		res = __gong__marshallFloat(ident, "TimeLine_FillOpacity", gantt.TimeLine_FillOpacity)
	case "TimeLine_Stroke":
		res = __gong__marshallString(ident, "TimeLine_Stroke", gantt.TimeLine_Stroke)
	case "TimeLine_StrokeWidth":
		res = __gong__marshallFloat(ident, "TimeLine_StrokeWidth", gantt.TimeLine_StrokeWidth)
	case "Group_Stroke":
		res = __gong__marshallString(ident, "Group_Stroke", gantt.Group_Stroke)
	case "Group_StrokeWidth":
		res = __gong__marshallFloat(ident, "Group_StrokeWidth", gantt.Group_StrokeWidth)
	case "Group_StrokeDashArray":
		res = __gong__marshallString(ident, "Group_StrokeDashArray", gantt.Group_StrokeDashArray)
	case "DateYOffset":
		res = __gong__marshallFloat(ident, "DateYOffset", gantt.DateYOffset)
	case "AlignOnStartEndOnYearStart":
		res = __gong__marshallBool(ident, "AlignOnStartEndOnYearStart", gantt.AlignOnStartEndOnYearStart)

	case "Lanes":
		var sb strings.Builder
		for _, _lane := range gantt.Lanes {
			sb.WriteString(__gong__marshallSliceOfPointers(ident, "Lanes", _lane.GongGetIdentifier(stage)))
		}
		res = sb.String()
	case "Milestones":
		var sb strings.Builder
		for _, _milestone := range gantt.Milestones {
			sb.WriteString(__gong__marshallSliceOfPointers(ident, "Milestones", _milestone.GongGetIdentifier(stage)))
		}
		res = sb.String()
	case "Groups":
		var sb strings.Builder
		for _, _group := range gantt.Groups {
			sb.WriteString(__gong__marshallSliceOfPointers(ident, "Groups", _group.GongGetIdentifier(stage)))
		}
		res = sb.String()
	case "Arrows":
		var sb strings.Builder
		for _, _arrow := range gantt.Arrows {
			sb.WriteString(__gong__marshallSliceOfPointers(ident, "Arrows", _arrow.GongGetIdentifier(stage)))
		}
		res = sb.String()
	default:
		log.Panicf("Unknown field %s for Gongstruct Gantt", fieldName)
	}
	return
}

func (group *Group) GongMarshallField(stage *Stage, fieldName string) (res string) {
	ident := group.GongGetIdentifier(stage)
	_ = ident
	switch fieldName {
	case "Name":
		res = __gong__marshallString(ident, "Name", group.Name)

	case "GroupLanes":
		var sb strings.Builder
		for _, _lane := range group.GroupLanes {
			sb.WriteString(__gong__marshallSliceOfPointers(ident, "GroupLanes", _lane.GongGetIdentifier(stage)))
		}
		res = sb.String()
	default:
		log.Panicf("Unknown field %s for Gongstruct Group", fieldName)
	}
	return
}

func (lane *Lane) GongMarshallField(stage *Stage, fieldName string) (res string) {
	ident := lane.GongGetIdentifier(stage)
	_ = ident
	switch fieldName {
	case "Name":
		res = __gong__marshallString(ident, "Name", lane.Name)
	case "Order":
		res = __gong__marshallInt(ident, "Order", lane.Order)

	case "Bars":
		var sb strings.Builder
		for _, _bar := range lane.Bars {
			sb.WriteString(__gong__marshallSliceOfPointers(ident, "Bars", _bar.GongGetIdentifier(stage)))
		}
		res = sb.String()
	default:
		log.Panicf("Unknown field %s for Gongstruct Lane", fieldName)
	}
	return
}

func (laneuse *LaneUse) GongMarshallField(stage *Stage, fieldName string) (res string) {
	ident := laneuse.GongGetIdentifier(stage)
	_ = ident
	switch fieldName {
	case "Name":
		res = __gong__marshallString(ident, "Name", laneuse.Name)

	case "Lane":
		if laneuse.Lane != nil {
			res = __gong__marshallPointer(ident, "Lane", laneuse.Lane.GongGetIdentifier(stage))
		} else {
			res = __gong__marshallPointer(ident, "Lane", "nil")
		}
	default:
		log.Panicf("Unknown field %s for Gongstruct LaneUse", fieldName)
	}
	return
}

func (milestone *Milestone) GongMarshallField(stage *Stage, fieldName string) (res string) {
	ident := milestone.GongGetIdentifier(stage)
	_ = ident
	switch fieldName {
	case "Name":
		res = __gong__marshallString(ident, "Name", milestone.Name)
	case "Date":
		res = __gong__marshallTime(ident, "Date", milestone.Date.String())
	case "DisplayVerticalBar":
		res = __gong__marshallBool(ident, "DisplayVerticalBar", milestone.DisplayVerticalBar)

	case "LanesToDisplay":
		var sb strings.Builder
		for _, _lane := range milestone.LanesToDisplay {
			sb.WriteString(__gong__marshallSliceOfPointers(ident, "LanesToDisplay", _lane.GongGetIdentifier(stage)))
		}
		res = sb.String()
	default:
		log.Panicf("Unknown field %s for Gongstruct Milestone", fieldName)
	}
	return
}

// insertion point for marshall all fields methods
func (arrow *Arrow) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(arrow.GongMarshallField(stage, "Name"))
		pointersInitializesStatements.WriteString(arrow.GongMarshallField(stage, "From"))
		pointersInitializesStatements.WriteString(arrow.GongMarshallField(stage, "To"))
		initializerStatements.WriteString(arrow.GongMarshallField(stage, "OptionnalColor"))
		initializerStatements.WriteString(arrow.GongMarshallField(stage, "OptionnalStroke"))
	}
	initRes = initializerStatements.String()
	ptrRes = pointersInitializesStatements.String()
	return
}
func (bar *Bar) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(bar.GongMarshallField(stage, "Name"))
		initializerStatements.WriteString(bar.GongMarshallField(stage, "Start"))
		initializerStatements.WriteString(bar.GongMarshallField(stage, "End"))
		initializerStatements.WriteString(bar.GongMarshallField(stage, "ComputedDuration"))
		initializerStatements.WriteString(bar.GongMarshallField(stage, "OptionnalColor"))
		initializerStatements.WriteString(bar.GongMarshallField(stage, "OptionnalStroke"))
		initializerStatements.WriteString(bar.GongMarshallField(stage, "FillOpacity"))
		initializerStatements.WriteString(bar.GongMarshallField(stage, "StrokeWidth"))
		initializerStatements.WriteString(bar.GongMarshallField(stage, "StrokeDashArray"))
	}
	initRes = initializerStatements.String()
	ptrRes = pointersInitializesStatements.String()
	return
}
func (gantt *Gantt) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(gantt.GongMarshallField(stage, "Name"))
		initializerStatements.WriteString(gantt.GongMarshallField(stage, "ComputedStart"))
		initializerStatements.WriteString(gantt.GongMarshallField(stage, "ComputedEnd"))
		initializerStatements.WriteString(gantt.GongMarshallField(stage, "ComputedDuration"))
		initializerStatements.WriteString(gantt.GongMarshallField(stage, "UseManualStartAndEndDates"))
		initializerStatements.WriteString(gantt.GongMarshallField(stage, "ManualStart"))
		initializerStatements.WriteString(gantt.GongMarshallField(stage, "ManualEnd"))
		initializerStatements.WriteString(gantt.GongMarshallField(stage, "LaneHeight"))
		initializerStatements.WriteString(gantt.GongMarshallField(stage, "RatioBarToLaneHeight"))
		initializerStatements.WriteString(gantt.GongMarshallField(stage, "YTopMargin"))
		initializerStatements.WriteString(gantt.GongMarshallField(stage, "XLeftText"))
		initializerStatements.WriteString(gantt.GongMarshallField(stage, "TextHeight"))
		initializerStatements.WriteString(gantt.GongMarshallField(stage, "XLeftLanes"))
		initializerStatements.WriteString(gantt.GongMarshallField(stage, "XRightMargin"))
		initializerStatements.WriteString(gantt.GongMarshallField(stage, "ArrowLengthToTheRightOfStartBar"))
		initializerStatements.WriteString(gantt.GongMarshallField(stage, "ArrowTipLenght"))
		initializerStatements.WriteString(gantt.GongMarshallField(stage, "TimeLine_Color"))
		initializerStatements.WriteString(gantt.GongMarshallField(stage, "TimeLine_FillOpacity"))
		initializerStatements.WriteString(gantt.GongMarshallField(stage, "TimeLine_Stroke"))
		initializerStatements.WriteString(gantt.GongMarshallField(stage, "TimeLine_StrokeWidth"))
		initializerStatements.WriteString(gantt.GongMarshallField(stage, "Group_Stroke"))
		initializerStatements.WriteString(gantt.GongMarshallField(stage, "Group_StrokeWidth"))
		initializerStatements.WriteString(gantt.GongMarshallField(stage, "Group_StrokeDashArray"))
		initializerStatements.WriteString(gantt.GongMarshallField(stage, "DateYOffset"))
		initializerStatements.WriteString(gantt.GongMarshallField(stage, "AlignOnStartEndOnYearStart"))
		pointersInitializesStatements.WriteString(gantt.GongMarshallField(stage, "Lanes"))
		pointersInitializesStatements.WriteString(gantt.GongMarshallField(stage, "Milestones"))
		pointersInitializesStatements.WriteString(gantt.GongMarshallField(stage, "Groups"))
		pointersInitializesStatements.WriteString(gantt.GongMarshallField(stage, "Arrows"))
	}
	initRes = initializerStatements.String()
	ptrRes = pointersInitializesStatements.String()
	return
}
func (group *Group) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(group.GongMarshallField(stage, "Name"))
		pointersInitializesStatements.WriteString(group.GongMarshallField(stage, "GroupLanes"))
	}
	initRes = initializerStatements.String()
	ptrRes = pointersInitializesStatements.String()
	return
}
func (lane *Lane) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(lane.GongMarshallField(stage, "Name"))
		initializerStatements.WriteString(lane.GongMarshallField(stage, "Order"))
		pointersInitializesStatements.WriteString(lane.GongMarshallField(stage, "Bars"))
	}
	initRes = initializerStatements.String()
	ptrRes = pointersInitializesStatements.String()
	return
}
func (laneuse *LaneUse) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(laneuse.GongMarshallField(stage, "Name"))
		pointersInitializesStatements.WriteString(laneuse.GongMarshallField(stage, "Lane"))
	}
	initRes = initializerStatements.String()
	ptrRes = pointersInitializesStatements.String()
	return
}
func (milestone *Milestone) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(milestone.GongMarshallField(stage, "Name"))
		initializerStatements.WriteString(milestone.GongMarshallField(stage, "Date"))
		initializerStatements.WriteString(milestone.GongMarshallField(stage, "DisplayVerticalBar"))
		pointersInitializesStatements.WriteString(milestone.GongMarshallField(stage, "LanesToDisplay"))
	}
	initRes = initializerStatements.String()
	ptrRes = pointersInitializesStatements.String()
	return
}
