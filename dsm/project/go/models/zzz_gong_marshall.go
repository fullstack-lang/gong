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
	modelsPackageName := "github.com/fullstack-lang/gong/dsm/project/go/models"
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
	localAliasPrefix := "__models"
	if stage.IsInDeltaMode() || stage.GetGongMarshallingMode() == GongMarshallingAppendCommit {
		localAliasPrefix = ""
	}
	_ = localAliasPrefix

	if stage != nil {
		for _, diagram := range __gong__sortStageSetInstances(stage.Diagrams, stage.Diagram_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			diagramIdent := localAliasPrefix + diagram.GongGetIdentifier(stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&models.Diagram{Name: %s}).Stage(stage)", diagramIdent, __gong__toRawStringLiteral(diagram.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", diagramIdent, __gong__toRawStringLiteral(diagram.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.DefaultBoxWidth = %f", diagramIdent, diagram.DefaultBoxWidth))
			values.WriteString(fmt.Sprintf("\n\t%s.DefaultBoxHeigth = %f", diagramIdent, diagram.DefaultBoxHeigth))
			values.WriteString(fmt.Sprintf("\n\t%s.DateFormat = %s", diagramIdent, __gong__toRawStringLiteral(diagram.DateFormat)))
			values.WriteString(fmt.Sprintf("\n\t%s.Width = %f", diagramIdent, diagram.Width))
			values.WriteString(fmt.Sprintf("\n\t%s.Height = %f", diagramIdent, diagram.Height))
			values.WriteString(fmt.Sprintf("\n\t%s.IsTimeDiagram = %t", diagramIdent, diagram.IsTimeDiagram))
			values.WriteString(fmt.Sprintf("\n\t%s.ComputedStart, _ = time.Parse(\"2006-01-02 15:04:05.999999999 -0700 MST\", \"%s\")", diagramIdent, diagram.ComputedStart.String()))
			values.WriteString(fmt.Sprintf("\n\t%s.ComputedEnd, _ = time.Parse(\"2006-01-02 15:04:05.999999999 -0700 MST\", \"%s\")", diagramIdent, diagram.ComputedEnd.String()))
			values.WriteString(fmt.Sprintf("\n\t%s.ComputedDuration = time.Duration(%d)", diagramIdent, int64(diagram.ComputedDuration)))
			values.WriteString(fmt.Sprintf("\n\t%s.DrawVerticalTimeLines = %t", diagramIdent, diagram.DrawVerticalTimeLines))
			values.WriteString(fmt.Sprintf("\n\t%s.DrawSecondaryVerticalTimeLines = %t", diagramIdent, diagram.DrawSecondaryVerticalTimeLines))
			values.WriteString(fmt.Sprintf("\n\t%s.HideWeekendsPeriod = %t", diagramIdent, diagram.HideWeekendsPeriod))
			values.WriteString(fmt.Sprintf("\n\t%s.UseManualStartAndEndDates = %t", diagramIdent, diagram.UseManualStartAndEndDates))
			values.WriteString(fmt.Sprintf("\n\t%s.ManualStart, _ = time.Parse(\"2006-01-02 15:04:05.999999999 -0700 MST\", \"%s\")", diagramIdent, diagram.ManualStart.String()))
			values.WriteString(fmt.Sprintf("\n\t%s.ManualEnd, _ = time.Parse(\"2006-01-02 15:04:05.999999999 -0700 MST\", \"%s\")", diagramIdent, diagram.ManualEnd.String()))
			values.WriteString(fmt.Sprintf("\n\t%s.TimeStep = %d", diagramIdent, diagram.TimeStep))
			if diagram.TimeStepScale.ToCodeString() != "" {
				values.WriteString(fmt.Sprintf("\n\t%s.TimeStepScale = models.%s", diagramIdent, diagram.TimeStepScale.ToCodeString()))
			} else {
				values.WriteString(fmt.Sprintf("\n\t%s.TimeStepScale = \"\"", diagramIdent))
			}
			values.WriteString(fmt.Sprintf("\n\t%s.SecondaryTimeStep = %d", diagramIdent, diagram.SecondaryTimeStep))
			if diagram.SecondaryTimeStepScale.ToCodeString() != "" {
				values.WriteString(fmt.Sprintf("\n\t%s.SecondaryTimeStepScale = models.%s", diagramIdent, diagram.SecondaryTimeStepScale.ToCodeString()))
			} else {
				values.WriteString(fmt.Sprintf("\n\t%s.SecondaryTimeStepScale = \"\"", diagramIdent))
			}
			values.WriteString(fmt.Sprintf("\n\t%s.LaneHeight = %f", diagramIdent, diagram.LaneHeight))
			values.WriteString(fmt.Sprintf("\n\t%s.RatioBarToLaneHeight = %f", diagramIdent, diagram.RatioBarToLaneHeight))
			values.WriteString(fmt.Sprintf("\n\t%s.YTopMargin = %f", diagramIdent, diagram.YTopMargin))
			values.WriteString(fmt.Sprintf("\n\t%s.XLeftText = %f", diagramIdent, diagram.XLeftText))
			values.WriteString(fmt.Sprintf("\n\t%s.TextHeight = %f", diagramIdent, diagram.TextHeight))
			values.WriteString(fmt.Sprintf("\n\t%s.XLeftLanes = %f", diagramIdent, diagram.XLeftLanes))
			values.WriteString(fmt.Sprintf("\n\t%s.XRightMargin = %f", diagramIdent, diagram.XRightMargin))
			values.WriteString(fmt.Sprintf("\n\t%s.ArrowLengthToTheRightOfStartBar = %f", diagramIdent, diagram.ArrowLengthToTheRightOfStartBar))
			values.WriteString(fmt.Sprintf("\n\t%s.ArrowTipLenght = %f", diagramIdent, diagram.ArrowTipLenght))
			values.WriteString(fmt.Sprintf("\n\t%s.TimeLine_Color = %s", diagramIdent, __gong__toRawStringLiteral(diagram.TimeLine_Color)))
			values.WriteString(fmt.Sprintf("\n\t%s.TimeLine_FillOpacity = %f", diagramIdent, diagram.TimeLine_FillOpacity))
			values.WriteString(fmt.Sprintf("\n\t%s.TimeLine_Stroke = %s", diagramIdent, __gong__toRawStringLiteral(diagram.TimeLine_Stroke)))
			values.WriteString(fmt.Sprintf("\n\t%s.TimeLine_StrokeWidth = %f", diagramIdent, diagram.TimeLine_StrokeWidth))
			values.WriteString(fmt.Sprintf("\n\t%s.Group_Stroke = %s", diagramIdent, __gong__toRawStringLiteral(diagram.Group_Stroke)))
			values.WriteString(fmt.Sprintf("\n\t%s.Group_StrokeWidth = %f", diagramIdent, diagram.Group_StrokeWidth))
			values.WriteString(fmt.Sprintf("\n\t%s.Group_StrokeDashArray = %s", diagramIdent, __gong__toRawStringLiteral(diagram.Group_StrokeDashArray)))
			values.WriteString(fmt.Sprintf("\n\t%s.DateYOffset = %f", diagramIdent, diagram.DateYOffset))
			values.WriteString(fmt.Sprintf("\n\t%s.AlignOnBeginningOfTimeScale = %t", diagramIdent, diagram.AlignOnBeginningOfTimeScale))
			values.WriteString(fmt.Sprintf("\n\t%s.ComputedPrefix = %s", diagramIdent, __gong__toRawStringLiteral(diagram.ComputedPrefix)))
			values.WriteString(fmt.Sprintf("\n\t%s.IsExpanded = %t", diagramIdent, diagram.IsExpanded))
			values.WriteString(fmt.Sprintf("\n\t%s.IsChecked = %t", diagramIdent, diagram.IsChecked))
			values.WriteString(fmt.Sprintf("\n\t%s.IsEditable_ = %t", diagramIdent, diagram.IsEditable_))
			values.WriteString(fmt.Sprintf("\n\t%s.IsShowPrefix = %t", diagramIdent, diagram.IsShowPrefix))
			values.WriteString(fmt.Sprintf("\n\t%s.IsInAutoLayoutMode = %t", diagramIdent, diagram.IsInAutoLayoutMode))
			values.WriteString(fmt.Sprintf("\n\t%s.IsPBSNodeExpanded = %t", diagramIdent, diagram.IsPBSNodeExpanded))
			values.WriteString(fmt.Sprintf("\n\t%s.IsWBSNodeExpanded = %t", diagramIdent, diagram.IsWBSNodeExpanded))
			values.WriteString(fmt.Sprintf("\n\t%s.IsTaskGroupsNodeExpanded = %t", diagramIdent, diagram.IsTaskGroupsNodeExpanded))
			values.WriteString(fmt.Sprintf("\n\t%s.IsNotesNodeExpanded = %t", diagramIdent, diagram.IsNotesNodeExpanded))
			values.WriteString(fmt.Sprintf("\n\t%s.IsResourcesNodeExpanded = %t", diagramIdent, diagram.IsResourcesNodeExpanded))
			for _, elem := range diagram.Product_Shapes {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := localAliasPrefix + elem.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Product_Shapes = append(%s.Product_Shapes, %s)", diagramIdent, diagramIdent, targetIdent))
			}
			for _, elem := range diagram.ProductsWhoseNodeIsExpanded {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := localAliasPrefix + elem.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.ProductsWhoseNodeIsExpanded = append(%s.ProductsWhoseNodeIsExpanded, %s)", diagramIdent, diagramIdent, targetIdent))
			}
			for _, elem := range diagram.ProductComposition_Shapes {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := localAliasPrefix + elem.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.ProductComposition_Shapes = append(%s.ProductComposition_Shapes, %s)", diagramIdent, diagramIdent, targetIdent))
			}
			for _, elem := range diagram.ProductReference_Shapes {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := localAliasPrefix + elem.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.ProductReference_Shapes = append(%s.ProductReference_Shapes, %s)", diagramIdent, diagramIdent, targetIdent))
			}
			for _, elem := range diagram.Task_Shapes {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := localAliasPrefix + elem.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Task_Shapes = append(%s.Task_Shapes, %s)", diagramIdent, diagramIdent, targetIdent))
			}
			for _, elem := range diagram.TasksWhoseNodeIsExpanded {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := localAliasPrefix + elem.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.TasksWhoseNodeIsExpanded = append(%s.TasksWhoseNodeIsExpanded, %s)", diagramIdent, diagramIdent, targetIdent))
			}
			for _, elem := range diagram.TasksWhoseInputNodeIsExpanded {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := localAliasPrefix + elem.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.TasksWhoseInputNodeIsExpanded = append(%s.TasksWhoseInputNodeIsExpanded, %s)", diagramIdent, diagramIdent, targetIdent))
			}
			for _, elem := range diagram.TasksWhoseOutputNodeIsExpanded {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := localAliasPrefix + elem.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.TasksWhoseOutputNodeIsExpanded = append(%s.TasksWhoseOutputNodeIsExpanded, %s)", diagramIdent, diagramIdent, targetIdent))
			}
			for _, elem := range diagram.TasksWhosePredecessorNodeIsExpanded {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := localAliasPrefix + elem.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.TasksWhosePredecessorNodeIsExpanded = append(%s.TasksWhosePredecessorNodeIsExpanded, %s)", diagramIdent, diagramIdent, targetIdent))
			}
			for _, elem := range diagram.TaskGroupShapes {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := localAliasPrefix + elem.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.TaskGroupShapes = append(%s.TaskGroupShapes, %s)", diagramIdent, diagramIdent, targetIdent))
			}
			for _, elem := range diagram.TaskGroupsWhoseNodeIsExpanded {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := localAliasPrefix + elem.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.TaskGroupsWhoseNodeIsExpanded = append(%s.TaskGroupsWhoseNodeIsExpanded, %s)", diagramIdent, diagramIdent, targetIdent))
			}
			for _, elem := range diagram.TaskComposition_Shapes {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := localAliasPrefix + elem.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.TaskComposition_Shapes = append(%s.TaskComposition_Shapes, %s)", diagramIdent, diagramIdent, targetIdent))
			}
			for _, elem := range diagram.TaskInputShapes {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := localAliasPrefix + elem.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.TaskInputShapes = append(%s.TaskInputShapes, %s)", diagramIdent, diagramIdent, targetIdent))
			}
			for _, elem := range diagram.TaskOutputShapes {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := localAliasPrefix + elem.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.TaskOutputShapes = append(%s.TaskOutputShapes, %s)", diagramIdent, diagramIdent, targetIdent))
			}
			for _, elem := range diagram.TaskPredecessorShapes {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := localAliasPrefix + elem.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.TaskPredecessorShapes = append(%s.TaskPredecessorShapes, %s)", diagramIdent, diagramIdent, targetIdent))
			}
			for _, elem := range diagram.Note_Shapes {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := localAliasPrefix + elem.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Note_Shapes = append(%s.Note_Shapes, %s)", diagramIdent, diagramIdent, targetIdent))
			}
			for _, elem := range diagram.NotesWhoseNodeIsExpanded {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := localAliasPrefix + elem.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.NotesWhoseNodeIsExpanded = append(%s.NotesWhoseNodeIsExpanded, %s)", diagramIdent, diagramIdent, targetIdent))
			}
			for _, elem := range diagram.NoteProductShapes {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := localAliasPrefix + elem.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.NoteProductShapes = append(%s.NoteProductShapes, %s)", diagramIdent, diagramIdent, targetIdent))
			}
			for _, elem := range diagram.NoteTaskShapes {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := localAliasPrefix + elem.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.NoteTaskShapes = append(%s.NoteTaskShapes, %s)", diagramIdent, diagramIdent, targetIdent))
			}
			for _, elem := range diagram.NoteResourceShapes {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := localAliasPrefix + elem.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.NoteResourceShapes = append(%s.NoteResourceShapes, %s)", diagramIdent, diagramIdent, targetIdent))
			}
			for _, elem := range diagram.Resource_Shapes {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := localAliasPrefix + elem.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Resource_Shapes = append(%s.Resource_Shapes, %s)", diagramIdent, diagramIdent, targetIdent))
			}
			for _, elem := range diagram.ResourcesWhoseNodeIsExpanded {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := localAliasPrefix + elem.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.ResourcesWhoseNodeIsExpanded = append(%s.ResourcesWhoseNodeIsExpanded, %s)", diagramIdent, diagramIdent, targetIdent))
			}
			for _, elem := range diagram.ResourceComposition_Shapes {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := localAliasPrefix + elem.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.ResourceComposition_Shapes = append(%s.ResourceComposition_Shapes, %s)", diagramIdent, diagramIdent, targetIdent))
			}
			for _, elem := range diagram.ResourceTaskShapes {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := localAliasPrefix + elem.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.ResourceTaskShapes = append(%s.ResourceTaskShapes, %s)", diagramIdent, diagramIdent, targetIdent))
			}
		}
	}
	if stage != nil {
		for _, library := range __gong__sortStageSetInstances(stage.Librarys, stage.Library_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			libraryIdent := localAliasPrefix + library.GongGetIdentifier(stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&models.Library{Name: %s}).Stage(stage)", libraryIdent, __gong__toRawStringLiteral(library.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", libraryIdent, __gong__toRawStringLiteral(library.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.NbPixPerCharacter = %f", libraryIdent, library.NbPixPerCharacter))
			values.WriteString(fmt.Sprintf("\n\t%s.LogoSVGFile = %s", libraryIdent, __gong__toRawStringLiteral(library.LogoSVGFile)))
			values.WriteString(fmt.Sprintf("\n\t%s.ComputedPrefix = %s", libraryIdent, __gong__toRawStringLiteral(library.ComputedPrefix)))
			values.WriteString(fmt.Sprintf("\n\t%s.IsExpanded = %t", libraryIdent, library.IsExpanded))
			values.WriteString(fmt.Sprintf("\n\t%s.IsRootLibrary = %t", libraryIdent, library.IsRootLibrary))
			for _, elem := range library.SubLibraries {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := localAliasPrefix + elem.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.SubLibraries = append(%s.SubLibraries, %s)", libraryIdent, libraryIdent, targetIdent))
			}
			for _, elem := range library.RootProducts {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := localAliasPrefix + elem.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.RootProducts = append(%s.RootProducts, %s)", libraryIdent, libraryIdent, targetIdent))
			}
			for _, elem := range library.RootTasks {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := localAliasPrefix + elem.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.RootTasks = append(%s.RootTasks, %s)", libraryIdent, libraryIdent, targetIdent))
			}
			for _, elem := range library.RootTaskGroups {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := localAliasPrefix + elem.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.RootTaskGroups = append(%s.RootTaskGroups, %s)", libraryIdent, libraryIdent, targetIdent))
			}
			for _, elem := range library.RootResources {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := localAliasPrefix + elem.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.RootResources = append(%s.RootResources, %s)", libraryIdent, libraryIdent, targetIdent))
			}
			for _, elem := range library.Notes {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := localAliasPrefix + elem.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Notes = append(%s.Notes, %s)", libraryIdent, libraryIdent, targetIdent))
			}
			for _, elem := range library.Diagrams {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := localAliasPrefix + elem.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Diagrams = append(%s.Diagrams, %s)", libraryIdent, libraryIdent, targetIdent))
			}
		}
	}
	if stage != nil {
		for _, note := range __gong__sortStageSetInstances(stage.Notes, stage.Note_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			noteIdent := localAliasPrefix + note.GongGetIdentifier(stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&models.Note{Name: %s}).Stage(stage)", noteIdent, __gong__toRawStringLiteral(note.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", noteIdent, __gong__toRawStringLiteral(note.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.ComputedPrefix = %s", noteIdent, __gong__toRawStringLiteral(note.ComputedPrefix)))
			values.WriteString(fmt.Sprintf("\n\t%s.IsExpanded = %t", noteIdent, note.IsExpanded))
			if note.LayoutDirection.ToCodeString() != "" {
				values.WriteString(fmt.Sprintf("\n\t%s.LayoutDirection = models.%s", noteIdent, note.LayoutDirection.ToCodeString()))
			} else {
				values.WriteString(fmt.Sprintf("\n\t%s.LayoutDirection = 0", noteIdent))
			}
			for _, elem := range note.Products {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := localAliasPrefix + elem.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Products = append(%s.Products, %s)", noteIdent, noteIdent, targetIdent))
			}
			for _, elem := range note.Tasks {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := localAliasPrefix + elem.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Tasks = append(%s.Tasks, %s)", noteIdent, noteIdent, targetIdent))
			}
			for _, elem := range note.Resources {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := localAliasPrefix + elem.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Resources = append(%s.Resources, %s)", noteIdent, noteIdent, targetIdent))
			}
		}
	}
	if stage != nil {
		for _, noteproductshape := range __gong__sortStageSetInstances(stage.NoteProductShapes, stage.NoteProductShape_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			noteproductshapeIdent := localAliasPrefix + noteproductshape.GongGetIdentifier(stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&models.NoteProductShape{Name: %s}).Stage(stage)", noteproductshapeIdent, __gong__toRawStringLiteral(noteproductshape.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", noteproductshapeIdent, __gong__toRawStringLiteral(noteproductshape.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.StartRatio = %f", noteproductshapeIdent, noteproductshape.StartRatio))
			values.WriteString(fmt.Sprintf("\n\t%s.EndRatio = %f", noteproductshapeIdent, noteproductshape.EndRatio))
			if noteproductshape.StartOrientation.ToCodeString() != "" {
				values.WriteString(fmt.Sprintf("\n\t%s.StartOrientation = models.%s", noteproductshapeIdent, noteproductshape.StartOrientation.ToCodeString()))
			} else {
				values.WriteString(fmt.Sprintf("\n\t%s.StartOrientation = \"\"", noteproductshapeIdent))
			}
			if noteproductshape.EndOrientation.ToCodeString() != "" {
				values.WriteString(fmt.Sprintf("\n\t%s.EndOrientation = models.%s", noteproductshapeIdent, noteproductshape.EndOrientation.ToCodeString()))
			} else {
				values.WriteString(fmt.Sprintf("\n\t%s.EndOrientation = \"\"", noteproductshapeIdent))
			}
			values.WriteString(fmt.Sprintf("\n\t%s.CornerOffsetRatio = %f", noteproductshapeIdent, noteproductshape.CornerOffsetRatio))
			values.WriteString(fmt.Sprintf("\n\t%s.IsHidden = %t", noteproductshapeIdent, noteproductshape.IsHidden))
			values.WriteString(fmt.Sprintf("\n\t%s.IsPresentationOverridden = %t", noteproductshapeIdent, noteproductshape.IsPresentationOverridden))
			values.WriteString(fmt.Sprintf("\n\t%s.Stroke = %s", noteproductshapeIdent, __gong__toRawStringLiteral(noteproductshape.Stroke)))
			values.WriteString(fmt.Sprintf("\n\t%s.StrokeWidth = %f", noteproductshapeIdent, noteproductshape.StrokeWidth))
			values.WriteString(fmt.Sprintf("\n\t%s.StrokeOpacity = %f", noteproductshapeIdent, noteproductshape.StrokeOpacity))
			values.WriteString(fmt.Sprintf("\n\t%s.StrokeDashArray = %s", noteproductshapeIdent, __gong__toRawStringLiteral(noteproductshape.StrokeDashArray)))
			values.WriteString(fmt.Sprintf("\n\t%s.FontSize = %s", noteproductshapeIdent, __gong__toRawStringLiteral(noteproductshape.FontSize)))
			values.WriteString(fmt.Sprintf("\n\t%s.FontColor = %s", noteproductshapeIdent, __gong__toRawStringLiteral(noteproductshape.FontColor)))
			if noteproductshape.Note != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := localAliasPrefix + noteproductshape.Note.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Note = %s", noteproductshapeIdent, targetIdent))
			}
			if noteproductshape.Product != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := localAliasPrefix + noteproductshape.Product.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Product = %s", noteproductshapeIdent, targetIdent))
			}
		}
	}
	if stage != nil {
		for _, noteresourceshape := range __gong__sortStageSetInstances(stage.NoteResourceShapes, stage.NoteResourceShape_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			noteresourceshapeIdent := localAliasPrefix + noteresourceshape.GongGetIdentifier(stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&models.NoteResourceShape{Name: %s}).Stage(stage)", noteresourceshapeIdent, __gong__toRawStringLiteral(noteresourceshape.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", noteresourceshapeIdent, __gong__toRawStringLiteral(noteresourceshape.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.StartRatio = %f", noteresourceshapeIdent, noteresourceshape.StartRatio))
			values.WriteString(fmt.Sprintf("\n\t%s.EndRatio = %f", noteresourceshapeIdent, noteresourceshape.EndRatio))
			if noteresourceshape.StartOrientation.ToCodeString() != "" {
				values.WriteString(fmt.Sprintf("\n\t%s.StartOrientation = models.%s", noteresourceshapeIdent, noteresourceshape.StartOrientation.ToCodeString()))
			} else {
				values.WriteString(fmt.Sprintf("\n\t%s.StartOrientation = \"\"", noteresourceshapeIdent))
			}
			if noteresourceshape.EndOrientation.ToCodeString() != "" {
				values.WriteString(fmt.Sprintf("\n\t%s.EndOrientation = models.%s", noteresourceshapeIdent, noteresourceshape.EndOrientation.ToCodeString()))
			} else {
				values.WriteString(fmt.Sprintf("\n\t%s.EndOrientation = \"\"", noteresourceshapeIdent))
			}
			values.WriteString(fmt.Sprintf("\n\t%s.CornerOffsetRatio = %f", noteresourceshapeIdent, noteresourceshape.CornerOffsetRatio))
			values.WriteString(fmt.Sprintf("\n\t%s.IsHidden = %t", noteresourceshapeIdent, noteresourceshape.IsHidden))
			values.WriteString(fmt.Sprintf("\n\t%s.IsPresentationOverridden = %t", noteresourceshapeIdent, noteresourceshape.IsPresentationOverridden))
			values.WriteString(fmt.Sprintf("\n\t%s.Stroke = %s", noteresourceshapeIdent, __gong__toRawStringLiteral(noteresourceshape.Stroke)))
			values.WriteString(fmt.Sprintf("\n\t%s.StrokeWidth = %f", noteresourceshapeIdent, noteresourceshape.StrokeWidth))
			values.WriteString(fmt.Sprintf("\n\t%s.StrokeOpacity = %f", noteresourceshapeIdent, noteresourceshape.StrokeOpacity))
			values.WriteString(fmt.Sprintf("\n\t%s.StrokeDashArray = %s", noteresourceshapeIdent, __gong__toRawStringLiteral(noteresourceshape.StrokeDashArray)))
			values.WriteString(fmt.Sprintf("\n\t%s.FontSize = %s", noteresourceshapeIdent, __gong__toRawStringLiteral(noteresourceshape.FontSize)))
			values.WriteString(fmt.Sprintf("\n\t%s.FontColor = %s", noteresourceshapeIdent, __gong__toRawStringLiteral(noteresourceshape.FontColor)))
			if noteresourceshape.Note != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := localAliasPrefix + noteresourceshape.Note.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Note = %s", noteresourceshapeIdent, targetIdent))
			}
			if noteresourceshape.Resource != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := localAliasPrefix + noteresourceshape.Resource.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Resource = %s", noteresourceshapeIdent, targetIdent))
			}
		}
	}
	if stage != nil {
		for _, noteshape := range __gong__sortStageSetInstances(stage.NoteShapes, stage.NoteShape_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			noteshapeIdent := localAliasPrefix + noteshape.GongGetIdentifier(stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&models.NoteShape{Name: %s}).Stage(stage)", noteshapeIdent, __gong__toRawStringLiteral(noteshape.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", noteshapeIdent, __gong__toRawStringLiteral(noteshape.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.IsLayoutDirectionDifferent = %t", noteshapeIdent, noteshape.IsLayoutDirectionDifferent))
			values.WriteString(fmt.Sprintf("\n\t%s.X = %f", noteshapeIdent, noteshape.X))
			values.WriteString(fmt.Sprintf("\n\t%s.Y = %f", noteshapeIdent, noteshape.Y))
			values.WriteString(fmt.Sprintf("\n\t%s.Width = %f", noteshapeIdent, noteshape.Width))
			values.WriteString(fmt.Sprintf("\n\t%s.Height = %f", noteshapeIdent, noteshape.Height))
			values.WriteString(fmt.Sprintf("\n\t%s.IsHidden = %t", noteshapeIdent, noteshape.IsHidden))
			values.WriteString(fmt.Sprintf("\n\t%s.IsPresentationOverridden = %t", noteshapeIdent, noteshape.IsPresentationOverridden))
			values.WriteString(fmt.Sprintf("\n\t%s.Color = %s", noteshapeIdent, __gong__toRawStringLiteral(noteshape.Color)))
			values.WriteString(fmt.Sprintf("\n\t%s.FillOpacity = %f", noteshapeIdent, noteshape.FillOpacity))
			values.WriteString(fmt.Sprintf("\n\t%s.Stroke = %s", noteshapeIdent, __gong__toRawStringLiteral(noteshape.Stroke)))
			values.WriteString(fmt.Sprintf("\n\t%s.StrokeWidth = %f", noteshapeIdent, noteshape.StrokeWidth))
			values.WriteString(fmt.Sprintf("\n\t%s.StrokeOpacity = %f", noteshapeIdent, noteshape.StrokeOpacity))
			values.WriteString(fmt.Sprintf("\n\t%s.StrokeDashArray = %s", noteshapeIdent, __gong__toRawStringLiteral(noteshape.StrokeDashArray)))
			values.WriteString(fmt.Sprintf("\n\t%s.RX = %f", noteshapeIdent, noteshape.RX))
			values.WriteString(fmt.Sprintf("\n\t%s.FontSize = %s", noteshapeIdent, __gong__toRawStringLiteral(noteshape.FontSize)))
			values.WriteString(fmt.Sprintf("\n\t%s.FontWeight = %s", noteshapeIdent, __gong__toRawStringLiteral(noteshape.FontWeight)))
			values.WriteString(fmt.Sprintf("\n\t%s.FontColor = %s", noteshapeIdent, __gong__toRawStringLiteral(noteshape.FontColor)))
			if noteshape.Note != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := localAliasPrefix + noteshape.Note.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Note = %s", noteshapeIdent, targetIdent))
			}
		}
	}
	if stage != nil {
		for _, notetaskshape := range __gong__sortStageSetInstances(stage.NoteTaskShapes, stage.NoteTaskShape_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			notetaskshapeIdent := localAliasPrefix + notetaskshape.GongGetIdentifier(stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&models.NoteTaskShape{Name: %s}).Stage(stage)", notetaskshapeIdent, __gong__toRawStringLiteral(notetaskshape.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", notetaskshapeIdent, __gong__toRawStringLiteral(notetaskshape.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.StartRatio = %f", notetaskshapeIdent, notetaskshape.StartRatio))
			values.WriteString(fmt.Sprintf("\n\t%s.EndRatio = %f", notetaskshapeIdent, notetaskshape.EndRatio))
			if notetaskshape.StartOrientation.ToCodeString() != "" {
				values.WriteString(fmt.Sprintf("\n\t%s.StartOrientation = models.%s", notetaskshapeIdent, notetaskshape.StartOrientation.ToCodeString()))
			} else {
				values.WriteString(fmt.Sprintf("\n\t%s.StartOrientation = \"\"", notetaskshapeIdent))
			}
			if notetaskshape.EndOrientation.ToCodeString() != "" {
				values.WriteString(fmt.Sprintf("\n\t%s.EndOrientation = models.%s", notetaskshapeIdent, notetaskshape.EndOrientation.ToCodeString()))
			} else {
				values.WriteString(fmt.Sprintf("\n\t%s.EndOrientation = \"\"", notetaskshapeIdent))
			}
			values.WriteString(fmt.Sprintf("\n\t%s.CornerOffsetRatio = %f", notetaskshapeIdent, notetaskshape.CornerOffsetRatio))
			values.WriteString(fmt.Sprintf("\n\t%s.IsHidden = %t", notetaskshapeIdent, notetaskshape.IsHidden))
			values.WriteString(fmt.Sprintf("\n\t%s.IsPresentationOverridden = %t", notetaskshapeIdent, notetaskshape.IsPresentationOverridden))
			values.WriteString(fmt.Sprintf("\n\t%s.Stroke = %s", notetaskshapeIdent, __gong__toRawStringLiteral(notetaskshape.Stroke)))
			values.WriteString(fmt.Sprintf("\n\t%s.StrokeWidth = %f", notetaskshapeIdent, notetaskshape.StrokeWidth))
			values.WriteString(fmt.Sprintf("\n\t%s.StrokeOpacity = %f", notetaskshapeIdent, notetaskshape.StrokeOpacity))
			values.WriteString(fmt.Sprintf("\n\t%s.StrokeDashArray = %s", notetaskshapeIdent, __gong__toRawStringLiteral(notetaskshape.StrokeDashArray)))
			values.WriteString(fmt.Sprintf("\n\t%s.FontSize = %s", notetaskshapeIdent, __gong__toRawStringLiteral(notetaskshape.FontSize)))
			values.WriteString(fmt.Sprintf("\n\t%s.FontColor = %s", notetaskshapeIdent, __gong__toRawStringLiteral(notetaskshape.FontColor)))
			if notetaskshape.Note != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := localAliasPrefix + notetaskshape.Note.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Note = %s", notetaskshapeIdent, targetIdent))
			}
			if notetaskshape.Task != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := localAliasPrefix + notetaskshape.Task.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Task = %s", notetaskshapeIdent, targetIdent))
			}
		}
	}
	if stage != nil {
		for _, product := range __gong__sortStageSetInstances(stage.Products, stage.Product_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			productIdent := localAliasPrefix + product.GongGetIdentifier(stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&models.Product{Name: %s}).Stage(stage)", productIdent, __gong__toRawStringLiteral(product.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", productIdent, __gong__toRawStringLiteral(product.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Description = %s", productIdent, __gong__toRawStringLiteral(product.Description)))
			values.WriteString(fmt.Sprintf("\n\t%s.IsProducersNodeExpanded = %t", productIdent, product.IsProducersNodeExpanded))
			values.WriteString(fmt.Sprintf("\n\t%s.IsConsumersNodeExpanded = %t", productIdent, product.IsConsumersNodeExpanded))
			values.WriteString(fmt.Sprintf("\n\t%s.IsImport = %t", productIdent, product.IsImport))
			values.WriteString(fmt.Sprintf("\n\t%s.ComputedPrefix = %s", productIdent, __gong__toRawStringLiteral(product.ComputedPrefix)))
			values.WriteString(fmt.Sprintf("\n\t%s.IsExpanded = %t", productIdent, product.IsExpanded))
			if product.LayoutDirection.ToCodeString() != "" {
				values.WriteString(fmt.Sprintf("\n\t%s.LayoutDirection = models.%s", productIdent, product.LayoutDirection.ToCodeString()))
			} else {
				values.WriteString(fmt.Sprintf("\n\t%s.LayoutDirection = 0", productIdent))
			}
			for _, elem := range product.SubProducts {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := localAliasPrefix + elem.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.SubProducts = append(%s.SubProducts, %s)", productIdent, productIdent, targetIdent))
			}
			if product.ReferencedProduct != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := localAliasPrefix + product.ReferencedProduct.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.ReferencedProduct = %s", productIdent, targetIdent))
			}
		}
	}
	if stage != nil {
		for _, productcompositionshape := range __gong__sortStageSetInstances(stage.ProductCompositionShapes, stage.ProductCompositionShape_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			productcompositionshapeIdent := localAliasPrefix + productcompositionshape.GongGetIdentifier(stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&models.ProductCompositionShape{Name: %s}).Stage(stage)", productcompositionshapeIdent, __gong__toRawStringLiteral(productcompositionshape.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", productcompositionshapeIdent, __gong__toRawStringLiteral(productcompositionshape.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.StartRatio = %f", productcompositionshapeIdent, productcompositionshape.StartRatio))
			values.WriteString(fmt.Sprintf("\n\t%s.EndRatio = %f", productcompositionshapeIdent, productcompositionshape.EndRatio))
			if productcompositionshape.StartOrientation.ToCodeString() != "" {
				values.WriteString(fmt.Sprintf("\n\t%s.StartOrientation = models.%s", productcompositionshapeIdent, productcompositionshape.StartOrientation.ToCodeString()))
			} else {
				values.WriteString(fmt.Sprintf("\n\t%s.StartOrientation = \"\"", productcompositionshapeIdent))
			}
			if productcompositionshape.EndOrientation.ToCodeString() != "" {
				values.WriteString(fmt.Sprintf("\n\t%s.EndOrientation = models.%s", productcompositionshapeIdent, productcompositionshape.EndOrientation.ToCodeString()))
			} else {
				values.WriteString(fmt.Sprintf("\n\t%s.EndOrientation = \"\"", productcompositionshapeIdent))
			}
			values.WriteString(fmt.Sprintf("\n\t%s.CornerOffsetRatio = %f", productcompositionshapeIdent, productcompositionshape.CornerOffsetRatio))
			values.WriteString(fmt.Sprintf("\n\t%s.IsHidden = %t", productcompositionshapeIdent, productcompositionshape.IsHidden))
			values.WriteString(fmt.Sprintf("\n\t%s.IsPresentationOverridden = %t", productcompositionshapeIdent, productcompositionshape.IsPresentationOverridden))
			values.WriteString(fmt.Sprintf("\n\t%s.Stroke = %s", productcompositionshapeIdent, __gong__toRawStringLiteral(productcompositionshape.Stroke)))
			values.WriteString(fmt.Sprintf("\n\t%s.StrokeWidth = %f", productcompositionshapeIdent, productcompositionshape.StrokeWidth))
			values.WriteString(fmt.Sprintf("\n\t%s.StrokeOpacity = %f", productcompositionshapeIdent, productcompositionshape.StrokeOpacity))
			values.WriteString(fmt.Sprintf("\n\t%s.StrokeDashArray = %s", productcompositionshapeIdent, __gong__toRawStringLiteral(productcompositionshape.StrokeDashArray)))
			values.WriteString(fmt.Sprintf("\n\t%s.FontSize = %s", productcompositionshapeIdent, __gong__toRawStringLiteral(productcompositionshape.FontSize)))
			values.WriteString(fmt.Sprintf("\n\t%s.FontColor = %s", productcompositionshapeIdent, __gong__toRawStringLiteral(productcompositionshape.FontColor)))
			if productcompositionshape.Product != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := localAliasPrefix + productcompositionshape.Product.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Product = %s", productcompositionshapeIdent, targetIdent))
			}
		}
	}
	if stage != nil {
		for _, productreferenceshape := range __gong__sortStageSetInstances(stage.ProductReferenceShapes, stage.ProductReferenceShape_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			productreferenceshapeIdent := localAliasPrefix + productreferenceshape.GongGetIdentifier(stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&models.ProductReferenceShape{Name: %s}).Stage(stage)", productreferenceshapeIdent, __gong__toRawStringLiteral(productreferenceshape.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", productreferenceshapeIdent, __gong__toRawStringLiteral(productreferenceshape.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.StartRatio = %f", productreferenceshapeIdent, productreferenceshape.StartRatio))
			values.WriteString(fmt.Sprintf("\n\t%s.EndRatio = %f", productreferenceshapeIdent, productreferenceshape.EndRatio))
			if productreferenceshape.StartOrientation.ToCodeString() != "" {
				values.WriteString(fmt.Sprintf("\n\t%s.StartOrientation = models.%s", productreferenceshapeIdent, productreferenceshape.StartOrientation.ToCodeString()))
			} else {
				values.WriteString(fmt.Sprintf("\n\t%s.StartOrientation = \"\"", productreferenceshapeIdent))
			}
			if productreferenceshape.EndOrientation.ToCodeString() != "" {
				values.WriteString(fmt.Sprintf("\n\t%s.EndOrientation = models.%s", productreferenceshapeIdent, productreferenceshape.EndOrientation.ToCodeString()))
			} else {
				values.WriteString(fmt.Sprintf("\n\t%s.EndOrientation = \"\"", productreferenceshapeIdent))
			}
			values.WriteString(fmt.Sprintf("\n\t%s.CornerOffsetRatio = %f", productreferenceshapeIdent, productreferenceshape.CornerOffsetRatio))
			values.WriteString(fmt.Sprintf("\n\t%s.IsHidden = %t", productreferenceshapeIdent, productreferenceshape.IsHidden))
			values.WriteString(fmt.Sprintf("\n\t%s.IsPresentationOverridden = %t", productreferenceshapeIdent, productreferenceshape.IsPresentationOverridden))
			values.WriteString(fmt.Sprintf("\n\t%s.Stroke = %s", productreferenceshapeIdent, __gong__toRawStringLiteral(productreferenceshape.Stroke)))
			values.WriteString(fmt.Sprintf("\n\t%s.StrokeWidth = %f", productreferenceshapeIdent, productreferenceshape.StrokeWidth))
			values.WriteString(fmt.Sprintf("\n\t%s.StrokeOpacity = %f", productreferenceshapeIdent, productreferenceshape.StrokeOpacity))
			values.WriteString(fmt.Sprintf("\n\t%s.StrokeDashArray = %s", productreferenceshapeIdent, __gong__toRawStringLiteral(productreferenceshape.StrokeDashArray)))
			values.WriteString(fmt.Sprintf("\n\t%s.FontSize = %s", productreferenceshapeIdent, __gong__toRawStringLiteral(productreferenceshape.FontSize)))
			values.WriteString(fmt.Sprintf("\n\t%s.FontColor = %s", productreferenceshapeIdent, __gong__toRawStringLiteral(productreferenceshape.FontColor)))
			if productreferenceshape.Product != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := localAliasPrefix + productreferenceshape.Product.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Product = %s", productreferenceshapeIdent, targetIdent))
			}
			if productreferenceshape.ReferencedProduct != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := localAliasPrefix + productreferenceshape.ReferencedProduct.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.ReferencedProduct = %s", productreferenceshapeIdent, targetIdent))
			}
		}
	}
	if stage != nil {
		for _, productshape := range __gong__sortStageSetInstances(stage.ProductShapes, stage.ProductShape_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			productshapeIdent := localAliasPrefix + productshape.GongGetIdentifier(stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&models.ProductShape{Name: %s}).Stage(stage)", productshapeIdent, __gong__toRawStringLiteral(productshape.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", productshapeIdent, __gong__toRawStringLiteral(productshape.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.IsShowType = %t", productshapeIdent, productshape.IsShowType))
			values.WriteString(fmt.Sprintf("\n\t%s.IsLayoutDirectionDifferent = %t", productshapeIdent, productshape.IsLayoutDirectionDifferent))
			values.WriteString(fmt.Sprintf("\n\t%s.X = %f", productshapeIdent, productshape.X))
			values.WriteString(fmt.Sprintf("\n\t%s.Y = %f", productshapeIdent, productshape.Y))
			values.WriteString(fmt.Sprintf("\n\t%s.Width = %f", productshapeIdent, productshape.Width))
			values.WriteString(fmt.Sprintf("\n\t%s.Height = %f", productshapeIdent, productshape.Height))
			values.WriteString(fmt.Sprintf("\n\t%s.IsHidden = %t", productshapeIdent, productshape.IsHidden))
			values.WriteString(fmt.Sprintf("\n\t%s.IsPresentationOverridden = %t", productshapeIdent, productshape.IsPresentationOverridden))
			values.WriteString(fmt.Sprintf("\n\t%s.Color = %s", productshapeIdent, __gong__toRawStringLiteral(productshape.Color)))
			values.WriteString(fmt.Sprintf("\n\t%s.FillOpacity = %f", productshapeIdent, productshape.FillOpacity))
			values.WriteString(fmt.Sprintf("\n\t%s.Stroke = %s", productshapeIdent, __gong__toRawStringLiteral(productshape.Stroke)))
			values.WriteString(fmt.Sprintf("\n\t%s.StrokeWidth = %f", productshapeIdent, productshape.StrokeWidth))
			values.WriteString(fmt.Sprintf("\n\t%s.StrokeOpacity = %f", productshapeIdent, productshape.StrokeOpacity))
			values.WriteString(fmt.Sprintf("\n\t%s.StrokeDashArray = %s", productshapeIdent, __gong__toRawStringLiteral(productshape.StrokeDashArray)))
			values.WriteString(fmt.Sprintf("\n\t%s.RX = %f", productshapeIdent, productshape.RX))
			values.WriteString(fmt.Sprintf("\n\t%s.FontSize = %s", productshapeIdent, __gong__toRawStringLiteral(productshape.FontSize)))
			values.WriteString(fmt.Sprintf("\n\t%s.FontWeight = %s", productshapeIdent, __gong__toRawStringLiteral(productshape.FontWeight)))
			values.WriteString(fmt.Sprintf("\n\t%s.FontColor = %s", productshapeIdent, __gong__toRawStringLiteral(productshape.FontColor)))
			if productshape.Product != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := localAliasPrefix + productshape.Product.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Product = %s", productshapeIdent, targetIdent))
			}
		}
	}
	if stage != nil {
		for _, resource := range __gong__sortStageSetInstances(stage.Resources, stage.Resource_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			resourceIdent := localAliasPrefix + resource.GongGetIdentifier(stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&models.Resource{Name: %s}).Stage(stage)", resourceIdent, __gong__toRawStringLiteral(resource.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", resourceIdent, __gong__toRawStringLiteral(resource.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Description = %s", resourceIdent, __gong__toRawStringLiteral(resource.Description)))
			values.WriteString(fmt.Sprintf("\n\t%s.ComputedPrefix = %s", resourceIdent, __gong__toRawStringLiteral(resource.ComputedPrefix)))
			values.WriteString(fmt.Sprintf("\n\t%s.IsExpanded = %t", resourceIdent, resource.IsExpanded))
			if resource.LayoutDirection.ToCodeString() != "" {
				values.WriteString(fmt.Sprintf("\n\t%s.LayoutDirection = models.%s", resourceIdent, resource.LayoutDirection.ToCodeString()))
			} else {
				values.WriteString(fmt.Sprintf("\n\t%s.LayoutDirection = 0", resourceIdent))
			}
			values.WriteString(fmt.Sprintf("\n\t%s.IsImport = %t", resourceIdent, resource.IsImport))
			for _, elem := range resource.Tasks {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := localAliasPrefix + elem.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Tasks = append(%s.Tasks, %s)", resourceIdent, resourceIdent, targetIdent))
			}
			for _, elem := range resource.SubResources {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := localAliasPrefix + elem.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.SubResources = append(%s.SubResources, %s)", resourceIdent, resourceIdent, targetIdent))
			}
			if resource.ReferencedResource != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := localAliasPrefix + resource.ReferencedResource.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.ReferencedResource = %s", resourceIdent, targetIdent))
			}
		}
	}
	if stage != nil {
		for _, resourcecompositionshape := range __gong__sortStageSetInstances(stage.ResourceCompositionShapes, stage.ResourceCompositionShape_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			resourcecompositionshapeIdent := localAliasPrefix + resourcecompositionshape.GongGetIdentifier(stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&models.ResourceCompositionShape{Name: %s}).Stage(stage)", resourcecompositionshapeIdent, __gong__toRawStringLiteral(resourcecompositionshape.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", resourcecompositionshapeIdent, __gong__toRawStringLiteral(resourcecompositionshape.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.StartRatio = %f", resourcecompositionshapeIdent, resourcecompositionshape.StartRatio))
			values.WriteString(fmt.Sprintf("\n\t%s.EndRatio = %f", resourcecompositionshapeIdent, resourcecompositionshape.EndRatio))
			if resourcecompositionshape.StartOrientation.ToCodeString() != "" {
				values.WriteString(fmt.Sprintf("\n\t%s.StartOrientation = models.%s", resourcecompositionshapeIdent, resourcecompositionshape.StartOrientation.ToCodeString()))
			} else {
				values.WriteString(fmt.Sprintf("\n\t%s.StartOrientation = \"\"", resourcecompositionshapeIdent))
			}
			if resourcecompositionshape.EndOrientation.ToCodeString() != "" {
				values.WriteString(fmt.Sprintf("\n\t%s.EndOrientation = models.%s", resourcecompositionshapeIdent, resourcecompositionshape.EndOrientation.ToCodeString()))
			} else {
				values.WriteString(fmt.Sprintf("\n\t%s.EndOrientation = \"\"", resourcecompositionshapeIdent))
			}
			values.WriteString(fmt.Sprintf("\n\t%s.CornerOffsetRatio = %f", resourcecompositionshapeIdent, resourcecompositionshape.CornerOffsetRatio))
			values.WriteString(fmt.Sprintf("\n\t%s.IsHidden = %t", resourcecompositionshapeIdent, resourcecompositionshape.IsHidden))
			values.WriteString(fmt.Sprintf("\n\t%s.IsPresentationOverridden = %t", resourcecompositionshapeIdent, resourcecompositionshape.IsPresentationOverridden))
			values.WriteString(fmt.Sprintf("\n\t%s.Stroke = %s", resourcecompositionshapeIdent, __gong__toRawStringLiteral(resourcecompositionshape.Stroke)))
			values.WriteString(fmt.Sprintf("\n\t%s.StrokeWidth = %f", resourcecompositionshapeIdent, resourcecompositionshape.StrokeWidth))
			values.WriteString(fmt.Sprintf("\n\t%s.StrokeOpacity = %f", resourcecompositionshapeIdent, resourcecompositionshape.StrokeOpacity))
			values.WriteString(fmt.Sprintf("\n\t%s.StrokeDashArray = %s", resourcecompositionshapeIdent, __gong__toRawStringLiteral(resourcecompositionshape.StrokeDashArray)))
			values.WriteString(fmt.Sprintf("\n\t%s.FontSize = %s", resourcecompositionshapeIdent, __gong__toRawStringLiteral(resourcecompositionshape.FontSize)))
			values.WriteString(fmt.Sprintf("\n\t%s.FontColor = %s", resourcecompositionshapeIdent, __gong__toRawStringLiteral(resourcecompositionshape.FontColor)))
			if resourcecompositionshape.Resource != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := localAliasPrefix + resourcecompositionshape.Resource.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Resource = %s", resourcecompositionshapeIdent, targetIdent))
			}
		}
	}
	if stage != nil {
		for _, resourceshape := range __gong__sortStageSetInstances(stage.ResourceShapes, stage.ResourceShape_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			resourceshapeIdent := localAliasPrefix + resourceshape.GongGetIdentifier(stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&models.ResourceShape{Name: %s}).Stage(stage)", resourceshapeIdent, __gong__toRawStringLiteral(resourceshape.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", resourceshapeIdent, __gong__toRawStringLiteral(resourceshape.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.IsLayoutDirectionDifferent = %t", resourceshapeIdent, resourceshape.IsLayoutDirectionDifferent))
			values.WriteString(fmt.Sprintf("\n\t%s.X = %f", resourceshapeIdent, resourceshape.X))
			values.WriteString(fmt.Sprintf("\n\t%s.Y = %f", resourceshapeIdent, resourceshape.Y))
			values.WriteString(fmt.Sprintf("\n\t%s.Width = %f", resourceshapeIdent, resourceshape.Width))
			values.WriteString(fmt.Sprintf("\n\t%s.Height = %f", resourceshapeIdent, resourceshape.Height))
			values.WriteString(fmt.Sprintf("\n\t%s.IsHidden = %t", resourceshapeIdent, resourceshape.IsHidden))
			values.WriteString(fmt.Sprintf("\n\t%s.IsPresentationOverridden = %t", resourceshapeIdent, resourceshape.IsPresentationOverridden))
			values.WriteString(fmt.Sprintf("\n\t%s.Color = %s", resourceshapeIdent, __gong__toRawStringLiteral(resourceshape.Color)))
			values.WriteString(fmt.Sprintf("\n\t%s.FillOpacity = %f", resourceshapeIdent, resourceshape.FillOpacity))
			values.WriteString(fmt.Sprintf("\n\t%s.Stroke = %s", resourceshapeIdent, __gong__toRawStringLiteral(resourceshape.Stroke)))
			values.WriteString(fmt.Sprintf("\n\t%s.StrokeWidth = %f", resourceshapeIdent, resourceshape.StrokeWidth))
			values.WriteString(fmt.Sprintf("\n\t%s.StrokeOpacity = %f", resourceshapeIdent, resourceshape.StrokeOpacity))
			values.WriteString(fmt.Sprintf("\n\t%s.StrokeDashArray = %s", resourceshapeIdent, __gong__toRawStringLiteral(resourceshape.StrokeDashArray)))
			values.WriteString(fmt.Sprintf("\n\t%s.RX = %f", resourceshapeIdent, resourceshape.RX))
			values.WriteString(fmt.Sprintf("\n\t%s.FontSize = %s", resourceshapeIdent, __gong__toRawStringLiteral(resourceshape.FontSize)))
			values.WriteString(fmt.Sprintf("\n\t%s.FontWeight = %s", resourceshapeIdent, __gong__toRawStringLiteral(resourceshape.FontWeight)))
			values.WriteString(fmt.Sprintf("\n\t%s.FontColor = %s", resourceshapeIdent, __gong__toRawStringLiteral(resourceshape.FontColor)))
			if resourceshape.Resource != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := localAliasPrefix + resourceshape.Resource.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Resource = %s", resourceshapeIdent, targetIdent))
			}
		}
	}
	if stage != nil {
		for _, resourcetaskshape := range __gong__sortStageSetInstances(stage.ResourceTaskShapes, stage.ResourceTaskShape_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			resourcetaskshapeIdent := localAliasPrefix + resourcetaskshape.GongGetIdentifier(stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&models.ResourceTaskShape{Name: %s}).Stage(stage)", resourcetaskshapeIdent, __gong__toRawStringLiteral(resourcetaskshape.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", resourcetaskshapeIdent, __gong__toRawStringLiteral(resourcetaskshape.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.StartRatio = %f", resourcetaskshapeIdent, resourcetaskshape.StartRatio))
			values.WriteString(fmt.Sprintf("\n\t%s.EndRatio = %f", resourcetaskshapeIdent, resourcetaskshape.EndRatio))
			if resourcetaskshape.StartOrientation.ToCodeString() != "" {
				values.WriteString(fmt.Sprintf("\n\t%s.StartOrientation = models.%s", resourcetaskshapeIdent, resourcetaskshape.StartOrientation.ToCodeString()))
			} else {
				values.WriteString(fmt.Sprintf("\n\t%s.StartOrientation = \"\"", resourcetaskshapeIdent))
			}
			if resourcetaskshape.EndOrientation.ToCodeString() != "" {
				values.WriteString(fmt.Sprintf("\n\t%s.EndOrientation = models.%s", resourcetaskshapeIdent, resourcetaskshape.EndOrientation.ToCodeString()))
			} else {
				values.WriteString(fmt.Sprintf("\n\t%s.EndOrientation = \"\"", resourcetaskshapeIdent))
			}
			values.WriteString(fmt.Sprintf("\n\t%s.CornerOffsetRatio = %f", resourcetaskshapeIdent, resourcetaskshape.CornerOffsetRatio))
			values.WriteString(fmt.Sprintf("\n\t%s.IsHidden = %t", resourcetaskshapeIdent, resourcetaskshape.IsHidden))
			values.WriteString(fmt.Sprintf("\n\t%s.IsPresentationOverridden = %t", resourcetaskshapeIdent, resourcetaskshape.IsPresentationOverridden))
			values.WriteString(fmt.Sprintf("\n\t%s.Stroke = %s", resourcetaskshapeIdent, __gong__toRawStringLiteral(resourcetaskshape.Stroke)))
			values.WriteString(fmt.Sprintf("\n\t%s.StrokeWidth = %f", resourcetaskshapeIdent, resourcetaskshape.StrokeWidth))
			values.WriteString(fmt.Sprintf("\n\t%s.StrokeOpacity = %f", resourcetaskshapeIdent, resourcetaskshape.StrokeOpacity))
			values.WriteString(fmt.Sprintf("\n\t%s.StrokeDashArray = %s", resourcetaskshapeIdent, __gong__toRawStringLiteral(resourcetaskshape.StrokeDashArray)))
			values.WriteString(fmt.Sprintf("\n\t%s.FontSize = %s", resourcetaskshapeIdent, __gong__toRawStringLiteral(resourcetaskshape.FontSize)))
			values.WriteString(fmt.Sprintf("\n\t%s.FontColor = %s", resourcetaskshapeIdent, __gong__toRawStringLiteral(resourcetaskshape.FontColor)))
			if resourcetaskshape.Resource != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := localAliasPrefix + resourcetaskshape.Resource.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Resource = %s", resourcetaskshapeIdent, targetIdent))
			}
			if resourcetaskshape.Task != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := localAliasPrefix + resourcetaskshape.Task.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Task = %s", resourcetaskshapeIdent, targetIdent))
			}
		}
	}
	if stage != nil {
		for _, task := range __gong__sortStageSetInstances(stage.Tasks, stage.Task_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			taskIdent := localAliasPrefix + task.GongGetIdentifier(stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&models.Task{Name: %s}).Stage(stage)", taskIdent, __gong__toRawStringLiteral(task.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", taskIdent, __gong__toRawStringLiteral(task.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Description = %s", taskIdent, __gong__toRawStringLiteral(task.Description)))
			values.WriteString(fmt.Sprintf("\n\t%s.Start, _ = time.Parse(\"2006-01-02 15:04:05.999999999 -0700 MST\", \"%s\")", taskIdent, task.Start.String()))
			values.WriteString(fmt.Sprintf("\n\t%s.End, _ = time.Parse(\"2006-01-02 15:04:05.999999999 -0700 MST\", \"%s\")", taskIdent, task.End.String()))
			values.WriteString(fmt.Sprintf("\n\t%s.IsAllDay = %t", taskIdent, task.IsAllDay))
			values.WriteString(fmt.Sprintf("\n\t%s.IsMilestone = %t", taskIdent, task.IsMilestone))
			if task.DependencyType.ToCodeString() != "" {
				values.WriteString(fmt.Sprintf("\n\t%s.DependencyType = models.%s", taskIdent, task.DependencyType.ToCodeString()))
			} else {
				values.WriteString(fmt.Sprintf("\n\t%s.DependencyType = \"\"", taskIdent))
			}
			values.WriteString(fmt.Sprintf("\n\t%s.DependencyDurationYears = %f", taskIdent, task.DependencyDurationYears))
			values.WriteString(fmt.Sprintf("\n\t%s.DependencyDurationMonths = %f", taskIdent, task.DependencyDurationMonths))
			values.WriteString(fmt.Sprintf("\n\t%s.DependencyDurationWeeks = %f", taskIdent, task.DependencyDurationWeeks))
			values.WriteString(fmt.Sprintf("\n\t%s.DependencyDurationDays = %f", taskIdent, task.DependencyDurationDays))
			values.WriteString(fmt.Sprintf("\n\t%s.DependencyDurationHours = %f", taskIdent, task.DependencyDurationHours))
			values.WriteString(fmt.Sprintf("\n\t%s.DurationYears = %f", taskIdent, task.DurationYears))
			values.WriteString(fmt.Sprintf("\n\t%s.DurationMonths = %f", taskIdent, task.DurationMonths))
			values.WriteString(fmt.Sprintf("\n\t%s.DurationWeeks = %f", taskIdent, task.DurationWeeks))
			values.WriteString(fmt.Sprintf("\n\t%s.DurationDays = %f", taskIdent, task.DurationDays))
			values.WriteString(fmt.Sprintf("\n\t%s.DurationHours = %f", taskIdent, task.DurationHours))
			values.WriteString(fmt.Sprintf("\n\t%s.IsEndDateComputedFromDuration = %t", taskIdent, task.IsEndDateComputedFromDuration))
			values.WriteString(fmt.Sprintf("\n\t%s.IsWithCompletion = %t", taskIdent, task.IsWithCompletion))
			if task.Completion.ToCodeString() != "" {
				values.WriteString(fmt.Sprintf("\n\t%s.Completion = models.%s", taskIdent, task.Completion.ToCodeString()))
			} else {
				values.WriteString(fmt.Sprintf("\n\t%s.Completion = \"\"", taskIdent))
			}
			if task.TextPosition.ToCodeString() != "" {
				values.WriteString(fmt.Sprintf("\n\t%s.TextPosition = models.%s", taskIdent, task.TextPosition.ToCodeString()))
			} else {
				values.WriteString(fmt.Sprintf("\n\t%s.TextPosition = \"\"", taskIdent))
			}
			values.WriteString(fmt.Sprintf("\n\t%s.XOffset = %f", taskIdent, task.XOffset))
			values.WriteString(fmt.Sprintf("\n\t%s.YOffset = %f", taskIdent, task.YOffset))
			values.WriteString(fmt.Sprintf("\n\t%s.IsImport = %t", taskIdent, task.IsImport))
			values.WriteString(fmt.Sprintf("\n\t%s.IsInputsNodeExpanded = %t", taskIdent, task.IsInputsNodeExpanded))
			values.WriteString(fmt.Sprintf("\n\t%s.IsOutputsNodeExpanded = %t", taskIdent, task.IsOutputsNodeExpanded))
			values.WriteString(fmt.Sprintf("\n\t%s.ComputedPrefix = %s", taskIdent, __gong__toRawStringLiteral(task.ComputedPrefix)))
			values.WriteString(fmt.Sprintf("\n\t%s.IsExpanded = %t", taskIdent, task.IsExpanded))
			if task.LayoutDirection.ToCodeString() != "" {
				values.WriteString(fmt.Sprintf("\n\t%s.LayoutDirection = models.%s", taskIdent, task.LayoutDirection.ToCodeString()))
			} else {
				values.WriteString(fmt.Sprintf("\n\t%s.LayoutDirection = 0", taskIdent))
			}
			for _, elem := range task.TaskGroups {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := localAliasPrefix + elem.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.TaskGroups = append(%s.TaskGroups, %s)", taskIdent, taskIdent, targetIdent))
			}
			for _, elem := range task.Predecessors {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := localAliasPrefix + elem.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Predecessors = append(%s.Predecessors, %s)", taskIdent, taskIdent, targetIdent))
			}
			for _, elem := range task.Inputs {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := localAliasPrefix + elem.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Inputs = append(%s.Inputs, %s)", taskIdent, taskIdent, targetIdent))
			}
			for _, elem := range task.Outputs {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := localAliasPrefix + elem.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Outputs = append(%s.Outputs, %s)", taskIdent, taskIdent, targetIdent))
			}
			for _, elem := range task.SubTasks {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := localAliasPrefix + elem.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.SubTasks = append(%s.SubTasks, %s)", taskIdent, taskIdent, targetIdent))
			}
			for _, elem := range task.TaskGroupsToDisplay {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := localAliasPrefix + elem.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.TaskGroupsToDisplay = append(%s.TaskGroupsToDisplay, %s)", taskIdent, taskIdent, targetIdent))
			}
			if task.ReferencedTask != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := localAliasPrefix + task.ReferencedTask.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.ReferencedTask = %s", taskIdent, targetIdent))
			}
		}
	}
	if stage != nil {
		for _, taskcompositionshape := range __gong__sortStageSetInstances(stage.TaskCompositionShapes, stage.TaskCompositionShape_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			taskcompositionshapeIdent := localAliasPrefix + taskcompositionshape.GongGetIdentifier(stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&models.TaskCompositionShape{Name: %s}).Stage(stage)", taskcompositionshapeIdent, __gong__toRawStringLiteral(taskcompositionshape.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", taskcompositionshapeIdent, __gong__toRawStringLiteral(taskcompositionshape.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.StartRatio = %f", taskcompositionshapeIdent, taskcompositionshape.StartRatio))
			values.WriteString(fmt.Sprintf("\n\t%s.EndRatio = %f", taskcompositionshapeIdent, taskcompositionshape.EndRatio))
			if taskcompositionshape.StartOrientation.ToCodeString() != "" {
				values.WriteString(fmt.Sprintf("\n\t%s.StartOrientation = models.%s", taskcompositionshapeIdent, taskcompositionshape.StartOrientation.ToCodeString()))
			} else {
				values.WriteString(fmt.Sprintf("\n\t%s.StartOrientation = \"\"", taskcompositionshapeIdent))
			}
			if taskcompositionshape.EndOrientation.ToCodeString() != "" {
				values.WriteString(fmt.Sprintf("\n\t%s.EndOrientation = models.%s", taskcompositionshapeIdent, taskcompositionshape.EndOrientation.ToCodeString()))
			} else {
				values.WriteString(fmt.Sprintf("\n\t%s.EndOrientation = \"\"", taskcompositionshapeIdent))
			}
			values.WriteString(fmt.Sprintf("\n\t%s.CornerOffsetRatio = %f", taskcompositionshapeIdent, taskcompositionshape.CornerOffsetRatio))
			values.WriteString(fmt.Sprintf("\n\t%s.IsHidden = %t", taskcompositionshapeIdent, taskcompositionshape.IsHidden))
			values.WriteString(fmt.Sprintf("\n\t%s.IsPresentationOverridden = %t", taskcompositionshapeIdent, taskcompositionshape.IsPresentationOverridden))
			values.WriteString(fmt.Sprintf("\n\t%s.Stroke = %s", taskcompositionshapeIdent, __gong__toRawStringLiteral(taskcompositionshape.Stroke)))
			values.WriteString(fmt.Sprintf("\n\t%s.StrokeWidth = %f", taskcompositionshapeIdent, taskcompositionshape.StrokeWidth))
			values.WriteString(fmt.Sprintf("\n\t%s.StrokeOpacity = %f", taskcompositionshapeIdent, taskcompositionshape.StrokeOpacity))
			values.WriteString(fmt.Sprintf("\n\t%s.StrokeDashArray = %s", taskcompositionshapeIdent, __gong__toRawStringLiteral(taskcompositionshape.StrokeDashArray)))
			values.WriteString(fmt.Sprintf("\n\t%s.FontSize = %s", taskcompositionshapeIdent, __gong__toRawStringLiteral(taskcompositionshape.FontSize)))
			values.WriteString(fmt.Sprintf("\n\t%s.FontColor = %s", taskcompositionshapeIdent, __gong__toRawStringLiteral(taskcompositionshape.FontColor)))
			if taskcompositionshape.Task != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := localAliasPrefix + taskcompositionshape.Task.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Task = %s", taskcompositionshapeIdent, targetIdent))
			}
		}
	}
	if stage != nil {
		for _, taskgroup := range __gong__sortStageSetInstances(stage.TaskGroups, stage.TaskGroup_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			taskgroupIdent := localAliasPrefix + taskgroup.GongGetIdentifier(stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&models.TaskGroup{Name: %s}).Stage(stage)", taskgroupIdent, __gong__toRawStringLiteral(taskgroup.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", taskgroupIdent, __gong__toRawStringLiteral(taskgroup.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.ComputedPrefix = %s", taskgroupIdent, __gong__toRawStringLiteral(taskgroup.ComputedPrefix)))
			values.WriteString(fmt.Sprintf("\n\t%s.IsExpanded = %t", taskgroupIdent, taskgroup.IsExpanded))
			for _, elem := range taskgroup.Tasks {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := localAliasPrefix + elem.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Tasks = append(%s.Tasks, %s)", taskgroupIdent, taskgroupIdent, targetIdent))
			}
		}
	}
	if stage != nil {
		for _, taskgroupshape := range __gong__sortStageSetInstances(stage.TaskGroupShapes, stage.TaskGroupShape_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			taskgroupshapeIdent := localAliasPrefix + taskgroupshape.GongGetIdentifier(stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&models.TaskGroupShape{Name: %s}).Stage(stage)", taskgroupshapeIdent, __gong__toRawStringLiteral(taskgroupshape.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", taskgroupshapeIdent, __gong__toRawStringLiteral(taskgroupshape.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.X = %f", taskgroupshapeIdent, taskgroupshape.X))
			values.WriteString(fmt.Sprintf("\n\t%s.Y = %f", taskgroupshapeIdent, taskgroupshape.Y))
			values.WriteString(fmt.Sprintf("\n\t%s.Width = %f", taskgroupshapeIdent, taskgroupshape.Width))
			values.WriteString(fmt.Sprintf("\n\t%s.Height = %f", taskgroupshapeIdent, taskgroupshape.Height))
			values.WriteString(fmt.Sprintf("\n\t%s.IsHidden = %t", taskgroupshapeIdent, taskgroupshape.IsHidden))
			values.WriteString(fmt.Sprintf("\n\t%s.IsPresentationOverridden = %t", taskgroupshapeIdent, taskgroupshape.IsPresentationOverridden))
			values.WriteString(fmt.Sprintf("\n\t%s.Color = %s", taskgroupshapeIdent, __gong__toRawStringLiteral(taskgroupshape.Color)))
			values.WriteString(fmt.Sprintf("\n\t%s.FillOpacity = %f", taskgroupshapeIdent, taskgroupshape.FillOpacity))
			values.WriteString(fmt.Sprintf("\n\t%s.Stroke = %s", taskgroupshapeIdent, __gong__toRawStringLiteral(taskgroupshape.Stroke)))
			values.WriteString(fmt.Sprintf("\n\t%s.StrokeWidth = %f", taskgroupshapeIdent, taskgroupshape.StrokeWidth))
			values.WriteString(fmt.Sprintf("\n\t%s.StrokeOpacity = %f", taskgroupshapeIdent, taskgroupshape.StrokeOpacity))
			values.WriteString(fmt.Sprintf("\n\t%s.StrokeDashArray = %s", taskgroupshapeIdent, __gong__toRawStringLiteral(taskgroupshape.StrokeDashArray)))
			values.WriteString(fmt.Sprintf("\n\t%s.RX = %f", taskgroupshapeIdent, taskgroupshape.RX))
			values.WriteString(fmt.Sprintf("\n\t%s.FontSize = %s", taskgroupshapeIdent, __gong__toRawStringLiteral(taskgroupshape.FontSize)))
			values.WriteString(fmt.Sprintf("\n\t%s.FontWeight = %s", taskgroupshapeIdent, __gong__toRawStringLiteral(taskgroupshape.FontWeight)))
			values.WriteString(fmt.Sprintf("\n\t%s.FontColor = %s", taskgroupshapeIdent, __gong__toRawStringLiteral(taskgroupshape.FontColor)))
			if taskgroupshape.TaskGroup != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := localAliasPrefix + taskgroupshape.TaskGroup.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.TaskGroup = %s", taskgroupshapeIdent, targetIdent))
			}
		}
	}
	if stage != nil {
		for _, taskinputshape := range __gong__sortStageSetInstances(stage.TaskInputShapes, stage.TaskInputShape_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			taskinputshapeIdent := localAliasPrefix + taskinputshape.GongGetIdentifier(stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&models.TaskInputShape{Name: %s}).Stage(stage)", taskinputshapeIdent, __gong__toRawStringLiteral(taskinputshape.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", taskinputshapeIdent, __gong__toRawStringLiteral(taskinputshape.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.StartRatio = %f", taskinputshapeIdent, taskinputshape.StartRatio))
			values.WriteString(fmt.Sprintf("\n\t%s.EndRatio = %f", taskinputshapeIdent, taskinputshape.EndRatio))
			if taskinputshape.StartOrientation.ToCodeString() != "" {
				values.WriteString(fmt.Sprintf("\n\t%s.StartOrientation = models.%s", taskinputshapeIdent, taskinputshape.StartOrientation.ToCodeString()))
			} else {
				values.WriteString(fmt.Sprintf("\n\t%s.StartOrientation = \"\"", taskinputshapeIdent))
			}
			if taskinputshape.EndOrientation.ToCodeString() != "" {
				values.WriteString(fmt.Sprintf("\n\t%s.EndOrientation = models.%s", taskinputshapeIdent, taskinputshape.EndOrientation.ToCodeString()))
			} else {
				values.WriteString(fmt.Sprintf("\n\t%s.EndOrientation = \"\"", taskinputshapeIdent))
			}
			values.WriteString(fmt.Sprintf("\n\t%s.CornerOffsetRatio = %f", taskinputshapeIdent, taskinputshape.CornerOffsetRatio))
			values.WriteString(fmt.Sprintf("\n\t%s.IsHidden = %t", taskinputshapeIdent, taskinputshape.IsHidden))
			values.WriteString(fmt.Sprintf("\n\t%s.IsPresentationOverridden = %t", taskinputshapeIdent, taskinputshape.IsPresentationOverridden))
			values.WriteString(fmt.Sprintf("\n\t%s.Stroke = %s", taskinputshapeIdent, __gong__toRawStringLiteral(taskinputshape.Stroke)))
			values.WriteString(fmt.Sprintf("\n\t%s.StrokeWidth = %f", taskinputshapeIdent, taskinputshape.StrokeWidth))
			values.WriteString(fmt.Sprintf("\n\t%s.StrokeOpacity = %f", taskinputshapeIdent, taskinputshape.StrokeOpacity))
			values.WriteString(fmt.Sprintf("\n\t%s.StrokeDashArray = %s", taskinputshapeIdent, __gong__toRawStringLiteral(taskinputshape.StrokeDashArray)))
			values.WriteString(fmt.Sprintf("\n\t%s.FontSize = %s", taskinputshapeIdent, __gong__toRawStringLiteral(taskinputshape.FontSize)))
			values.WriteString(fmt.Sprintf("\n\t%s.FontColor = %s", taskinputshapeIdent, __gong__toRawStringLiteral(taskinputshape.FontColor)))
			if taskinputshape.Product != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := localAliasPrefix + taskinputshape.Product.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Product = %s", taskinputshapeIdent, targetIdent))
			}
			if taskinputshape.Task != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := localAliasPrefix + taskinputshape.Task.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Task = %s", taskinputshapeIdent, targetIdent))
			}
		}
	}
	if stage != nil {
		for _, taskoutputshape := range __gong__sortStageSetInstances(stage.TaskOutputShapes, stage.TaskOutputShape_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			taskoutputshapeIdent := localAliasPrefix + taskoutputshape.GongGetIdentifier(stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&models.TaskOutputShape{Name: %s}).Stage(stage)", taskoutputshapeIdent, __gong__toRawStringLiteral(taskoutputshape.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", taskoutputshapeIdent, __gong__toRawStringLiteral(taskoutputshape.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.StartRatio = %f", taskoutputshapeIdent, taskoutputshape.StartRatio))
			values.WriteString(fmt.Sprintf("\n\t%s.EndRatio = %f", taskoutputshapeIdent, taskoutputshape.EndRatio))
			if taskoutputshape.StartOrientation.ToCodeString() != "" {
				values.WriteString(fmt.Sprintf("\n\t%s.StartOrientation = models.%s", taskoutputshapeIdent, taskoutputshape.StartOrientation.ToCodeString()))
			} else {
				values.WriteString(fmt.Sprintf("\n\t%s.StartOrientation = \"\"", taskoutputshapeIdent))
			}
			if taskoutputshape.EndOrientation.ToCodeString() != "" {
				values.WriteString(fmt.Sprintf("\n\t%s.EndOrientation = models.%s", taskoutputshapeIdent, taskoutputshape.EndOrientation.ToCodeString()))
			} else {
				values.WriteString(fmt.Sprintf("\n\t%s.EndOrientation = \"\"", taskoutputshapeIdent))
			}
			values.WriteString(fmt.Sprintf("\n\t%s.CornerOffsetRatio = %f", taskoutputshapeIdent, taskoutputshape.CornerOffsetRatio))
			values.WriteString(fmt.Sprintf("\n\t%s.IsHidden = %t", taskoutputshapeIdent, taskoutputshape.IsHidden))
			values.WriteString(fmt.Sprintf("\n\t%s.IsPresentationOverridden = %t", taskoutputshapeIdent, taskoutputshape.IsPresentationOverridden))
			values.WriteString(fmt.Sprintf("\n\t%s.Stroke = %s", taskoutputshapeIdent, __gong__toRawStringLiteral(taskoutputshape.Stroke)))
			values.WriteString(fmt.Sprintf("\n\t%s.StrokeWidth = %f", taskoutputshapeIdent, taskoutputshape.StrokeWidth))
			values.WriteString(fmt.Sprintf("\n\t%s.StrokeOpacity = %f", taskoutputshapeIdent, taskoutputshape.StrokeOpacity))
			values.WriteString(fmt.Sprintf("\n\t%s.StrokeDashArray = %s", taskoutputshapeIdent, __gong__toRawStringLiteral(taskoutputshape.StrokeDashArray)))
			values.WriteString(fmt.Sprintf("\n\t%s.FontSize = %s", taskoutputshapeIdent, __gong__toRawStringLiteral(taskoutputshape.FontSize)))
			values.WriteString(fmt.Sprintf("\n\t%s.FontColor = %s", taskoutputshapeIdent, __gong__toRawStringLiteral(taskoutputshape.FontColor)))
			if taskoutputshape.Task != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := localAliasPrefix + taskoutputshape.Task.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Task = %s", taskoutputshapeIdent, targetIdent))
			}
			if taskoutputshape.Product != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := localAliasPrefix + taskoutputshape.Product.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Product = %s", taskoutputshapeIdent, targetIdent))
			}
		}
	}
	if stage != nil {
		for _, taskpredecessorshape := range __gong__sortStageSetInstances(stage.TaskPredecessorShapes, stage.TaskPredecessorShape_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			taskpredecessorshapeIdent := localAliasPrefix + taskpredecessorshape.GongGetIdentifier(stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&models.TaskPredecessorShape{Name: %s}).Stage(stage)", taskpredecessorshapeIdent, __gong__toRawStringLiteral(taskpredecessorshape.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", taskpredecessorshapeIdent, __gong__toRawStringLiteral(taskpredecessorshape.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.StartRatio = %f", taskpredecessorshapeIdent, taskpredecessorshape.StartRatio))
			values.WriteString(fmt.Sprintf("\n\t%s.EndRatio = %f", taskpredecessorshapeIdent, taskpredecessorshape.EndRatio))
			if taskpredecessorshape.StartOrientation.ToCodeString() != "" {
				values.WriteString(fmt.Sprintf("\n\t%s.StartOrientation = models.%s", taskpredecessorshapeIdent, taskpredecessorshape.StartOrientation.ToCodeString()))
			} else {
				values.WriteString(fmt.Sprintf("\n\t%s.StartOrientation = \"\"", taskpredecessorshapeIdent))
			}
			if taskpredecessorshape.EndOrientation.ToCodeString() != "" {
				values.WriteString(fmt.Sprintf("\n\t%s.EndOrientation = models.%s", taskpredecessorshapeIdent, taskpredecessorshape.EndOrientation.ToCodeString()))
			} else {
				values.WriteString(fmt.Sprintf("\n\t%s.EndOrientation = \"\"", taskpredecessorshapeIdent))
			}
			values.WriteString(fmt.Sprintf("\n\t%s.CornerOffsetRatio = %f", taskpredecessorshapeIdent, taskpredecessorshape.CornerOffsetRatio))
			values.WriteString(fmt.Sprintf("\n\t%s.IsHidden = %t", taskpredecessorshapeIdent, taskpredecessorshape.IsHidden))
			values.WriteString(fmt.Sprintf("\n\t%s.IsPresentationOverridden = %t", taskpredecessorshapeIdent, taskpredecessorshape.IsPresentationOverridden))
			values.WriteString(fmt.Sprintf("\n\t%s.Stroke = %s", taskpredecessorshapeIdent, __gong__toRawStringLiteral(taskpredecessorshape.Stroke)))
			values.WriteString(fmt.Sprintf("\n\t%s.StrokeWidth = %f", taskpredecessorshapeIdent, taskpredecessorshape.StrokeWidth))
			values.WriteString(fmt.Sprintf("\n\t%s.StrokeOpacity = %f", taskpredecessorshapeIdent, taskpredecessorshape.StrokeOpacity))
			values.WriteString(fmt.Sprintf("\n\t%s.StrokeDashArray = %s", taskpredecessorshapeIdent, __gong__toRawStringLiteral(taskpredecessorshape.StrokeDashArray)))
			values.WriteString(fmt.Sprintf("\n\t%s.FontSize = %s", taskpredecessorshapeIdent, __gong__toRawStringLiteral(taskpredecessorshape.FontSize)))
			values.WriteString(fmt.Sprintf("\n\t%s.FontColor = %s", taskpredecessorshapeIdent, __gong__toRawStringLiteral(taskpredecessorshape.FontColor)))
			if taskpredecessorshape.Predecessor != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := localAliasPrefix + taskpredecessorshape.Predecessor.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Predecessor = %s", taskpredecessorshapeIdent, targetIdent))
			}
			if taskpredecessorshape.Task != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := localAliasPrefix + taskpredecessorshape.Task.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Task = %s", taskpredecessorshapeIdent, targetIdent))
			}
		}
	}
	if stage != nil {
		for _, taskshape := range __gong__sortStageSetInstances(stage.TaskShapes, stage.TaskShape_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			taskshapeIdent := localAliasPrefix + taskshape.GongGetIdentifier(stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&models.TaskShape{Name: %s}).Stage(stage)", taskshapeIdent, __gong__toRawStringLiteral(taskshape.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", taskshapeIdent, __gong__toRawStringLiteral(taskshape.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.IsShowDate = %t", taskshapeIdent, taskshape.IsShowDate))
			values.WriteString(fmt.Sprintf("\n\t%s.VerticalOffset = %f", taskshapeIdent, taskshape.VerticalOffset))
			values.WriteString(fmt.Sprintf("\n\t%s.DisplayVerticalBar = %t", taskshapeIdent, taskshape.DisplayVerticalBar))
			values.WriteString(fmt.Sprintf("\n\t%s.IsLayoutDirectionDifferent = %t", taskshapeIdent, taskshape.IsLayoutDirectionDifferent))
			values.WriteString(fmt.Sprintf("\n\t%s.X = %f", taskshapeIdent, taskshape.X))
			values.WriteString(fmt.Sprintf("\n\t%s.Y = %f", taskshapeIdent, taskshape.Y))
			values.WriteString(fmt.Sprintf("\n\t%s.Width = %f", taskshapeIdent, taskshape.Width))
			values.WriteString(fmt.Sprintf("\n\t%s.Height = %f", taskshapeIdent, taskshape.Height))
			values.WriteString(fmt.Sprintf("\n\t%s.IsHidden = %t", taskshapeIdent, taskshape.IsHidden))
			values.WriteString(fmt.Sprintf("\n\t%s.IsPresentationOverridden = %t", taskshapeIdent, taskshape.IsPresentationOverridden))
			values.WriteString(fmt.Sprintf("\n\t%s.Color = %s", taskshapeIdent, __gong__toRawStringLiteral(taskshape.Color)))
			values.WriteString(fmt.Sprintf("\n\t%s.FillOpacity = %f", taskshapeIdent, taskshape.FillOpacity))
			values.WriteString(fmt.Sprintf("\n\t%s.Stroke = %s", taskshapeIdent, __gong__toRawStringLiteral(taskshape.Stroke)))
			values.WriteString(fmt.Sprintf("\n\t%s.StrokeWidth = %f", taskshapeIdent, taskshape.StrokeWidth))
			values.WriteString(fmt.Sprintf("\n\t%s.StrokeOpacity = %f", taskshapeIdent, taskshape.StrokeOpacity))
			values.WriteString(fmt.Sprintf("\n\t%s.StrokeDashArray = %s", taskshapeIdent, __gong__toRawStringLiteral(taskshape.StrokeDashArray)))
			values.WriteString(fmt.Sprintf("\n\t%s.RX = %f", taskshapeIdent, taskshape.RX))
			values.WriteString(fmt.Sprintf("\n\t%s.FontSize = %s", taskshapeIdent, __gong__toRawStringLiteral(taskshape.FontSize)))
			values.WriteString(fmt.Sprintf("\n\t%s.FontWeight = %s", taskshapeIdent, __gong__toRawStringLiteral(taskshape.FontWeight)))
			values.WriteString(fmt.Sprintf("\n\t%s.FontColor = %s", taskshapeIdent, __gong__toRawStringLiteral(taskshape.FontColor)))
			if taskshape.Task != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := localAliasPrefix + taskshape.Task.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Task = %s", taskshapeIdent, targetIdent))
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
go run github.com/fullstack-lang/gong/dsm/project/go/models/probe/cmd/probe@%s stage.go
*/
package %s

import (
	"slices"
	"time"

	"github.com/fullstack-lang/gong/dsm/project/go/models"%s
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

	if modelsPackageName != "" && modelsPackageName != "github.com/fullstack-lang/gong/dsm/project/go/models" {
		res = strings.ReplaceAll(res, "\"github.com/fullstack-lang/gong/dsm/project/go/models\"", fmt.Sprintf("%q", modelsPackageName))
	}

	return res, nil
}

// insertion point for marshall field methods
func (diagram *Diagram) GongMarshallField(stage *Stage, fieldName string) (res string) {
	ident := diagram.GongGetIdentifier(stage)
	_ = ident
	switch fieldName {
	case "Name":
		res = __gong__marshallString(ident, "Name", diagram.Name)
	case "DefaultBoxWidth":
		res = __gong__marshallFloat(ident, "DefaultBoxWidth", diagram.DefaultBoxWidth)
	case "DefaultBoxHeigth":
		res = __gong__marshallFloat(ident, "DefaultBoxHeigth", diagram.DefaultBoxHeigth)
	case "DateFormat":
		res = __gong__marshallString(ident, "DateFormat", diagram.DateFormat)
	case "Width":
		res = __gong__marshallFloat(ident, "Width", diagram.Width)
	case "Height":
		res = __gong__marshallFloat(ident, "Height", diagram.Height)
	case "IsTimeDiagram":
		res = __gong__marshallBool(ident, "IsTimeDiagram", diagram.IsTimeDiagram)
	case "ComputedStart":
		res = __gong__marshallTime(ident, "ComputedStart", diagram.ComputedStart.String())
	case "ComputedEnd":
		res = __gong__marshallTime(ident, "ComputedEnd", diagram.ComputedEnd.String())
	case "ComputedDuration":
		res = __gong__marshallInt(ident, "ComputedDuration", diagram.ComputedDuration)
	case "DrawVerticalTimeLines":
		res = __gong__marshallBool(ident, "DrawVerticalTimeLines", diagram.DrawVerticalTimeLines)
	case "DrawSecondaryVerticalTimeLines":
		res = __gong__marshallBool(ident, "DrawSecondaryVerticalTimeLines", diagram.DrawSecondaryVerticalTimeLines)
	case "HideWeekendsPeriod":
		res = __gong__marshallBool(ident, "HideWeekendsPeriod", diagram.HideWeekendsPeriod)
	case "UseManualStartAndEndDates":
		res = __gong__marshallBool(ident, "UseManualStartAndEndDates", diagram.UseManualStartAndEndDates)
	case "ManualStart":
		res = __gong__marshallTime(ident, "ManualStart", diagram.ManualStart.String())
	case "ManualEnd":
		res = __gong__marshallTime(ident, "ManualEnd", diagram.ManualEnd.String())
	case "TimeStep":
		res = __gong__marshallInt(ident, "TimeStep", diagram.TimeStep)
	case "TimeStepScale":
		res = __gong__marshallEnumString(ident, "TimeStepScale", diagram.TimeStepScale.ToCodeString())
	case "SecondaryTimeStep":
		res = __gong__marshallInt(ident, "SecondaryTimeStep", diagram.SecondaryTimeStep)
	case "SecondaryTimeStepScale":
		res = __gong__marshallEnumString(ident, "SecondaryTimeStepScale", diagram.SecondaryTimeStepScale.ToCodeString())
	case "LaneHeight":
		res = __gong__marshallFloat(ident, "LaneHeight", diagram.LaneHeight)
	case "RatioBarToLaneHeight":
		res = __gong__marshallFloat(ident, "RatioBarToLaneHeight", diagram.RatioBarToLaneHeight)
	case "YTopMargin":
		res = __gong__marshallFloat(ident, "YTopMargin", diagram.YTopMargin)
	case "XLeftText":
		res = __gong__marshallFloat(ident, "XLeftText", diagram.XLeftText)
	case "TextHeight":
		res = __gong__marshallFloat(ident, "TextHeight", diagram.TextHeight)
	case "XLeftLanes":
		res = __gong__marshallFloat(ident, "XLeftLanes", diagram.XLeftLanes)
	case "XRightMargin":
		res = __gong__marshallFloat(ident, "XRightMargin", diagram.XRightMargin)
	case "ArrowLengthToTheRightOfStartBar":
		res = __gong__marshallFloat(ident, "ArrowLengthToTheRightOfStartBar", diagram.ArrowLengthToTheRightOfStartBar)
	case "ArrowTipLenght":
		res = __gong__marshallFloat(ident, "ArrowTipLenght", diagram.ArrowTipLenght)
	case "TimeLine_Color":
		res = __gong__marshallString(ident, "TimeLine_Color", diagram.TimeLine_Color)
	case "TimeLine_FillOpacity":
		res = __gong__marshallFloat(ident, "TimeLine_FillOpacity", diagram.TimeLine_FillOpacity)
	case "TimeLine_Stroke":
		res = __gong__marshallString(ident, "TimeLine_Stroke", diagram.TimeLine_Stroke)
	case "TimeLine_StrokeWidth":
		res = __gong__marshallFloat(ident, "TimeLine_StrokeWidth", diagram.TimeLine_StrokeWidth)
	case "Group_Stroke":
		res = __gong__marshallString(ident, "Group_Stroke", diagram.Group_Stroke)
	case "Group_StrokeWidth":
		res = __gong__marshallFloat(ident, "Group_StrokeWidth", diagram.Group_StrokeWidth)
	case "Group_StrokeDashArray":
		res = __gong__marshallString(ident, "Group_StrokeDashArray", diagram.Group_StrokeDashArray)
	case "DateYOffset":
		res = __gong__marshallFloat(ident, "DateYOffset", diagram.DateYOffset)
	case "AlignOnBeginningOfTimeScale":
		res = __gong__marshallBool(ident, "AlignOnBeginningOfTimeScale", diagram.AlignOnBeginningOfTimeScale)
	case "ComputedPrefix":
		res = __gong__marshallString(ident, "ComputedPrefix", diagram.ComputedPrefix)
	case "IsExpanded":
		res = __gong__marshallBool(ident, "IsExpanded", diagram.IsExpanded)
	case "IsChecked":
		res = __gong__marshallBool(ident, "IsChecked", diagram.IsChecked)
	case "IsEditable_":
		res = __gong__marshallBool(ident, "IsEditable_", diagram.IsEditable_)
	case "IsShowPrefix":
		res = __gong__marshallBool(ident, "IsShowPrefix", diagram.IsShowPrefix)
	case "IsInAutoLayoutMode":
		res = __gong__marshallBool(ident, "IsInAutoLayoutMode", diagram.IsInAutoLayoutMode)
	case "IsPBSNodeExpanded":
		res = __gong__marshallBool(ident, "IsPBSNodeExpanded", diagram.IsPBSNodeExpanded)
	case "IsWBSNodeExpanded":
		res = __gong__marshallBool(ident, "IsWBSNodeExpanded", diagram.IsWBSNodeExpanded)
	case "IsTaskGroupsNodeExpanded":
		res = __gong__marshallBool(ident, "IsTaskGroupsNodeExpanded", diagram.IsTaskGroupsNodeExpanded)
	case "IsNotesNodeExpanded":
		res = __gong__marshallBool(ident, "IsNotesNodeExpanded", diagram.IsNotesNodeExpanded)
	case "IsResourcesNodeExpanded":
		res = __gong__marshallBool(ident, "IsResourcesNodeExpanded", diagram.IsResourcesNodeExpanded)

	case "Product_Shapes":
		var sb strings.Builder
		for _, _productshape := range diagram.Product_Shapes {
			sb.WriteString(__gong__marshallSliceOfPointers(ident, "Product_Shapes", _productshape.GongGetIdentifier(stage)))
		}
		res = sb.String()
	case "ProductsWhoseNodeIsExpanded":
		var sb strings.Builder
		for _, _product := range diagram.ProductsWhoseNodeIsExpanded {
			sb.WriteString(__gong__marshallSliceOfPointers(ident, "ProductsWhoseNodeIsExpanded", _product.GongGetIdentifier(stage)))
		}
		res = sb.String()
	case "ProductComposition_Shapes":
		var sb strings.Builder
		for _, _productcompositionshape := range diagram.ProductComposition_Shapes {
			sb.WriteString(__gong__marshallSliceOfPointers(ident, "ProductComposition_Shapes", _productcompositionshape.GongGetIdentifier(stage)))
		}
		res = sb.String()
	case "ProductReference_Shapes":
		var sb strings.Builder
		for _, _productreferenceshape := range diagram.ProductReference_Shapes {
			sb.WriteString(__gong__marshallSliceOfPointers(ident, "ProductReference_Shapes", _productreferenceshape.GongGetIdentifier(stage)))
		}
		res = sb.String()
	case "Task_Shapes":
		var sb strings.Builder
		for _, _taskshape := range diagram.Task_Shapes {
			sb.WriteString(__gong__marshallSliceOfPointers(ident, "Task_Shapes", _taskshape.GongGetIdentifier(stage)))
		}
		res = sb.String()
	case "TasksWhoseNodeIsExpanded":
		var sb strings.Builder
		for _, _task := range diagram.TasksWhoseNodeIsExpanded {
			sb.WriteString(__gong__marshallSliceOfPointers(ident, "TasksWhoseNodeIsExpanded", _task.GongGetIdentifier(stage)))
		}
		res = sb.String()
	case "TasksWhoseInputNodeIsExpanded":
		var sb strings.Builder
		for _, _task := range diagram.TasksWhoseInputNodeIsExpanded {
			sb.WriteString(__gong__marshallSliceOfPointers(ident, "TasksWhoseInputNodeIsExpanded", _task.GongGetIdentifier(stage)))
		}
		res = sb.String()
	case "TasksWhoseOutputNodeIsExpanded":
		var sb strings.Builder
		for _, _task := range diagram.TasksWhoseOutputNodeIsExpanded {
			sb.WriteString(__gong__marshallSliceOfPointers(ident, "TasksWhoseOutputNodeIsExpanded", _task.GongGetIdentifier(stage)))
		}
		res = sb.String()
	case "TasksWhosePredecessorNodeIsExpanded":
		var sb strings.Builder
		for _, _task := range diagram.TasksWhosePredecessorNodeIsExpanded {
			sb.WriteString(__gong__marshallSliceOfPointers(ident, "TasksWhosePredecessorNodeIsExpanded", _task.GongGetIdentifier(stage)))
		}
		res = sb.String()
	case "TaskGroupShapes":
		var sb strings.Builder
		for _, _taskgroupshape := range diagram.TaskGroupShapes {
			sb.WriteString(__gong__marshallSliceOfPointers(ident, "TaskGroupShapes", _taskgroupshape.GongGetIdentifier(stage)))
		}
		res = sb.String()
	case "TaskGroupsWhoseNodeIsExpanded":
		var sb strings.Builder
		for _, _taskgroup := range diagram.TaskGroupsWhoseNodeIsExpanded {
			sb.WriteString(__gong__marshallSliceOfPointers(ident, "TaskGroupsWhoseNodeIsExpanded", _taskgroup.GongGetIdentifier(stage)))
		}
		res = sb.String()
	case "TaskComposition_Shapes":
		var sb strings.Builder
		for _, _taskcompositionshape := range diagram.TaskComposition_Shapes {
			sb.WriteString(__gong__marshallSliceOfPointers(ident, "TaskComposition_Shapes", _taskcompositionshape.GongGetIdentifier(stage)))
		}
		res = sb.String()
	case "TaskInputShapes":
		var sb strings.Builder
		for _, _taskinputshape := range diagram.TaskInputShapes {
			sb.WriteString(__gong__marshallSliceOfPointers(ident, "TaskInputShapes", _taskinputshape.GongGetIdentifier(stage)))
		}
		res = sb.String()
	case "TaskOutputShapes":
		var sb strings.Builder
		for _, _taskoutputshape := range diagram.TaskOutputShapes {
			sb.WriteString(__gong__marshallSliceOfPointers(ident, "TaskOutputShapes", _taskoutputshape.GongGetIdentifier(stage)))
		}
		res = sb.String()
	case "TaskPredecessorShapes":
		var sb strings.Builder
		for _, _taskpredecessorshape := range diagram.TaskPredecessorShapes {
			sb.WriteString(__gong__marshallSliceOfPointers(ident, "TaskPredecessorShapes", _taskpredecessorshape.GongGetIdentifier(stage)))
		}
		res = sb.String()
	case "Note_Shapes":
		var sb strings.Builder
		for _, _noteshape := range diagram.Note_Shapes {
			sb.WriteString(__gong__marshallSliceOfPointers(ident, "Note_Shapes", _noteshape.GongGetIdentifier(stage)))
		}
		res = sb.String()
	case "NotesWhoseNodeIsExpanded":
		var sb strings.Builder
		for _, _note := range diagram.NotesWhoseNodeIsExpanded {
			sb.WriteString(__gong__marshallSliceOfPointers(ident, "NotesWhoseNodeIsExpanded", _note.GongGetIdentifier(stage)))
		}
		res = sb.String()
	case "NoteProductShapes":
		var sb strings.Builder
		for _, _noteproductshape := range diagram.NoteProductShapes {
			sb.WriteString(__gong__marshallSliceOfPointers(ident, "NoteProductShapes", _noteproductshape.GongGetIdentifier(stage)))
		}
		res = sb.String()
	case "NoteTaskShapes":
		var sb strings.Builder
		for _, _notetaskshape := range diagram.NoteTaskShapes {
			sb.WriteString(__gong__marshallSliceOfPointers(ident, "NoteTaskShapes", _notetaskshape.GongGetIdentifier(stage)))
		}
		res = sb.String()
	case "NoteResourceShapes":
		var sb strings.Builder
		for _, _noteresourceshape := range diagram.NoteResourceShapes {
			sb.WriteString(__gong__marshallSliceOfPointers(ident, "NoteResourceShapes", _noteresourceshape.GongGetIdentifier(stage)))
		}
		res = sb.String()
	case "Resource_Shapes":
		var sb strings.Builder
		for _, _resourceshape := range diagram.Resource_Shapes {
			sb.WriteString(__gong__marshallSliceOfPointers(ident, "Resource_Shapes", _resourceshape.GongGetIdentifier(stage)))
		}
		res = sb.String()
	case "ResourcesWhoseNodeIsExpanded":
		var sb strings.Builder
		for _, _resource := range diagram.ResourcesWhoseNodeIsExpanded {
			sb.WriteString(__gong__marshallSliceOfPointers(ident, "ResourcesWhoseNodeIsExpanded", _resource.GongGetIdentifier(stage)))
		}
		res = sb.String()
	case "ResourceComposition_Shapes":
		var sb strings.Builder
		for _, _resourcecompositionshape := range diagram.ResourceComposition_Shapes {
			sb.WriteString(__gong__marshallSliceOfPointers(ident, "ResourceComposition_Shapes", _resourcecompositionshape.GongGetIdentifier(stage)))
		}
		res = sb.String()
	case "ResourceTaskShapes":
		var sb strings.Builder
		for _, _resourcetaskshape := range diagram.ResourceTaskShapes {
			sb.WriteString(__gong__marshallSliceOfPointers(ident, "ResourceTaskShapes", _resourcetaskshape.GongGetIdentifier(stage)))
		}
		res = sb.String()
	default:
		log.Panicf("Unknown field %s for Gongstruct Diagram", fieldName)
	}
	return
}

func (library *Library) GongMarshallField(stage *Stage, fieldName string) (res string) {
	ident := library.GongGetIdentifier(stage)
	_ = ident
	switch fieldName {
	case "Name":
		res = __gong__marshallString(ident, "Name", library.Name)
	case "NbPixPerCharacter":
		res = __gong__marshallFloat(ident, "NbPixPerCharacter", library.NbPixPerCharacter)
	case "LogoSVGFile":
		res = __gong__marshallString(ident, "LogoSVGFile", library.LogoSVGFile)
	case "ComputedPrefix":
		res = __gong__marshallString(ident, "ComputedPrefix", library.ComputedPrefix)
	case "IsExpanded":
		res = __gong__marshallBool(ident, "IsExpanded", library.IsExpanded)
	case "IsRootLibrary":
		res = __gong__marshallBool(ident, "IsRootLibrary", library.IsRootLibrary)

	case "SubLibraries":
		var sb strings.Builder
		for _, _library := range library.SubLibraries {
			sb.WriteString(__gong__marshallSliceOfPointers(ident, "SubLibraries", _library.GongGetIdentifier(stage)))
		}
		res = sb.String()
	case "RootProducts":
		var sb strings.Builder
		for _, _product := range library.RootProducts {
			sb.WriteString(__gong__marshallSliceOfPointers(ident, "RootProducts", _product.GongGetIdentifier(stage)))
		}
		res = sb.String()
	case "RootTasks":
		var sb strings.Builder
		for _, _task := range library.RootTasks {
			sb.WriteString(__gong__marshallSliceOfPointers(ident, "RootTasks", _task.GongGetIdentifier(stage)))
		}
		res = sb.String()
	case "RootTaskGroups":
		var sb strings.Builder
		for _, _taskgroup := range library.RootTaskGroups {
			sb.WriteString(__gong__marshallSliceOfPointers(ident, "RootTaskGroups", _taskgroup.GongGetIdentifier(stage)))
		}
		res = sb.String()
	case "RootResources":
		var sb strings.Builder
		for _, _resource := range library.RootResources {
			sb.WriteString(__gong__marshallSliceOfPointers(ident, "RootResources", _resource.GongGetIdentifier(stage)))
		}
		res = sb.String()
	case "Notes":
		var sb strings.Builder
		for _, _note := range library.Notes {
			sb.WriteString(__gong__marshallSliceOfPointers(ident, "Notes", _note.GongGetIdentifier(stage)))
		}
		res = sb.String()
	case "Diagrams":
		var sb strings.Builder
		for _, _diagram := range library.Diagrams {
			sb.WriteString(__gong__marshallSliceOfPointers(ident, "Diagrams", _diagram.GongGetIdentifier(stage)))
		}
		res = sb.String()
	default:
		log.Panicf("Unknown field %s for Gongstruct Library", fieldName)
	}
	return
}

func (note *Note) GongMarshallField(stage *Stage, fieldName string) (res string) {
	ident := note.GongGetIdentifier(stage)
	_ = ident
	switch fieldName {
	case "Name":
		res = __gong__marshallString(ident, "Name", note.Name)
	case "ComputedPrefix":
		res = __gong__marshallString(ident, "ComputedPrefix", note.ComputedPrefix)
	case "IsExpanded":
		res = __gong__marshallBool(ident, "IsExpanded", note.IsExpanded)
	case "LayoutDirection":
		res = __gong__marshallEnumInt(ident, "LayoutDirection", note.LayoutDirection.ToCodeString())

	case "Products":
		var sb strings.Builder
		for _, _product := range note.Products {
			sb.WriteString(__gong__marshallSliceOfPointers(ident, "Products", _product.GongGetIdentifier(stage)))
		}
		res = sb.String()
	case "Tasks":
		var sb strings.Builder
		for _, _task := range note.Tasks {
			sb.WriteString(__gong__marshallSliceOfPointers(ident, "Tasks", _task.GongGetIdentifier(stage)))
		}
		res = sb.String()
	case "Resources":
		var sb strings.Builder
		for _, _resource := range note.Resources {
			sb.WriteString(__gong__marshallSliceOfPointers(ident, "Resources", _resource.GongGetIdentifier(stage)))
		}
		res = sb.String()
	default:
		log.Panicf("Unknown field %s for Gongstruct Note", fieldName)
	}
	return
}

func (noteproductshape *NoteProductShape) GongMarshallField(stage *Stage, fieldName string) (res string) {
	ident := noteproductshape.GongGetIdentifier(stage)
	_ = ident
	switch fieldName {
	case "Name":
		res = __gong__marshallString(ident, "Name", noteproductshape.Name)
	case "StartRatio":
		res = __gong__marshallFloat(ident, "StartRatio", noteproductshape.StartRatio)
	case "EndRatio":
		res = __gong__marshallFloat(ident, "EndRatio", noteproductshape.EndRatio)
	case "StartOrientation":
		res = __gong__marshallEnumString(ident, "StartOrientation", noteproductshape.StartOrientation.ToCodeString())
	case "EndOrientation":
		res = __gong__marshallEnumString(ident, "EndOrientation", noteproductshape.EndOrientation.ToCodeString())
	case "CornerOffsetRatio":
		res = __gong__marshallFloat(ident, "CornerOffsetRatio", noteproductshape.CornerOffsetRatio)
	case "IsHidden":
		res = __gong__marshallBool(ident, "IsHidden", noteproductshape.IsHidden)
	case "IsPresentationOverridden":
		res = __gong__marshallBool(ident, "IsPresentationOverridden", noteproductshape.IsPresentationOverridden)
	case "Stroke":
		res = __gong__marshallString(ident, "Stroke", noteproductshape.Stroke)
	case "StrokeWidth":
		res = __gong__marshallFloat(ident, "StrokeWidth", noteproductshape.StrokeWidth)
	case "StrokeOpacity":
		res = __gong__marshallFloat(ident, "StrokeOpacity", noteproductshape.StrokeOpacity)
	case "StrokeDashArray":
		res = __gong__marshallString(ident, "StrokeDashArray", noteproductshape.StrokeDashArray)
	case "FontSize":
		res = __gong__marshallString(ident, "FontSize", noteproductshape.FontSize)
	case "FontColor":
		res = __gong__marshallString(ident, "FontColor", noteproductshape.FontColor)

	case "Note":
		if noteproductshape.Note != nil {
			res = __gong__marshallPointer(ident, "Note", noteproductshape.Note.GongGetIdentifier(stage))
		} else {
			res = __gong__marshallPointer(ident, "Note", "nil")
		}
	case "Product":
		if noteproductshape.Product != nil {
			res = __gong__marshallPointer(ident, "Product", noteproductshape.Product.GongGetIdentifier(stage))
		} else {
			res = __gong__marshallPointer(ident, "Product", "nil")
		}
	default:
		log.Panicf("Unknown field %s for Gongstruct NoteProductShape", fieldName)
	}
	return
}

func (noteresourceshape *NoteResourceShape) GongMarshallField(stage *Stage, fieldName string) (res string) {
	ident := noteresourceshape.GongGetIdentifier(stage)
	_ = ident
	switch fieldName {
	case "Name":
		res = __gong__marshallString(ident, "Name", noteresourceshape.Name)
	case "StartRatio":
		res = __gong__marshallFloat(ident, "StartRatio", noteresourceshape.StartRatio)
	case "EndRatio":
		res = __gong__marshallFloat(ident, "EndRatio", noteresourceshape.EndRatio)
	case "StartOrientation":
		res = __gong__marshallEnumString(ident, "StartOrientation", noteresourceshape.StartOrientation.ToCodeString())
	case "EndOrientation":
		res = __gong__marshallEnumString(ident, "EndOrientation", noteresourceshape.EndOrientation.ToCodeString())
	case "CornerOffsetRatio":
		res = __gong__marshallFloat(ident, "CornerOffsetRatio", noteresourceshape.CornerOffsetRatio)
	case "IsHidden":
		res = __gong__marshallBool(ident, "IsHidden", noteresourceshape.IsHidden)
	case "IsPresentationOverridden":
		res = __gong__marshallBool(ident, "IsPresentationOverridden", noteresourceshape.IsPresentationOverridden)
	case "Stroke":
		res = __gong__marshallString(ident, "Stroke", noteresourceshape.Stroke)
	case "StrokeWidth":
		res = __gong__marshallFloat(ident, "StrokeWidth", noteresourceshape.StrokeWidth)
	case "StrokeOpacity":
		res = __gong__marshallFloat(ident, "StrokeOpacity", noteresourceshape.StrokeOpacity)
	case "StrokeDashArray":
		res = __gong__marshallString(ident, "StrokeDashArray", noteresourceshape.StrokeDashArray)
	case "FontSize":
		res = __gong__marshallString(ident, "FontSize", noteresourceshape.FontSize)
	case "FontColor":
		res = __gong__marshallString(ident, "FontColor", noteresourceshape.FontColor)

	case "Note":
		if noteresourceshape.Note != nil {
			res = __gong__marshallPointer(ident, "Note", noteresourceshape.Note.GongGetIdentifier(stage))
		} else {
			res = __gong__marshallPointer(ident, "Note", "nil")
		}
	case "Resource":
		if noteresourceshape.Resource != nil {
			res = __gong__marshallPointer(ident, "Resource", noteresourceshape.Resource.GongGetIdentifier(stage))
		} else {
			res = __gong__marshallPointer(ident, "Resource", "nil")
		}
	default:
		log.Panicf("Unknown field %s for Gongstruct NoteResourceShape", fieldName)
	}
	return
}

func (noteshape *NoteShape) GongMarshallField(stage *Stage, fieldName string) (res string) {
	ident := noteshape.GongGetIdentifier(stage)
	_ = ident
	switch fieldName {
	case "Name":
		res = __gong__marshallString(ident, "Name", noteshape.Name)
	case "IsLayoutDirectionDifferent":
		res = __gong__marshallBool(ident, "IsLayoutDirectionDifferent", noteshape.IsLayoutDirectionDifferent)
	case "X":
		res = __gong__marshallFloat(ident, "X", noteshape.X)
	case "Y":
		res = __gong__marshallFloat(ident, "Y", noteshape.Y)
	case "Width":
		res = __gong__marshallFloat(ident, "Width", noteshape.Width)
	case "Height":
		res = __gong__marshallFloat(ident, "Height", noteshape.Height)
	case "IsHidden":
		res = __gong__marshallBool(ident, "IsHidden", noteshape.IsHidden)
	case "IsPresentationOverridden":
		res = __gong__marshallBool(ident, "IsPresentationOverridden", noteshape.IsPresentationOverridden)
	case "Color":
		res = __gong__marshallString(ident, "Color", noteshape.Color)
	case "FillOpacity":
		res = __gong__marshallFloat(ident, "FillOpacity", noteshape.FillOpacity)
	case "Stroke":
		res = __gong__marshallString(ident, "Stroke", noteshape.Stroke)
	case "StrokeWidth":
		res = __gong__marshallFloat(ident, "StrokeWidth", noteshape.StrokeWidth)
	case "StrokeOpacity":
		res = __gong__marshallFloat(ident, "StrokeOpacity", noteshape.StrokeOpacity)
	case "StrokeDashArray":
		res = __gong__marshallString(ident, "StrokeDashArray", noteshape.StrokeDashArray)
	case "RX":
		res = __gong__marshallFloat(ident, "RX", noteshape.RX)
	case "FontSize":
		res = __gong__marshallString(ident, "FontSize", noteshape.FontSize)
	case "FontWeight":
		res = __gong__marshallString(ident, "FontWeight", noteshape.FontWeight)
	case "FontColor":
		res = __gong__marshallString(ident, "FontColor", noteshape.FontColor)

	case "Note":
		if noteshape.Note != nil {
			res = __gong__marshallPointer(ident, "Note", noteshape.Note.GongGetIdentifier(stage))
		} else {
			res = __gong__marshallPointer(ident, "Note", "nil")
		}
	default:
		log.Panicf("Unknown field %s for Gongstruct NoteShape", fieldName)
	}
	return
}

func (notetaskshape *NoteTaskShape) GongMarshallField(stage *Stage, fieldName string) (res string) {
	ident := notetaskshape.GongGetIdentifier(stage)
	_ = ident
	switch fieldName {
	case "Name":
		res = __gong__marshallString(ident, "Name", notetaskshape.Name)
	case "StartRatio":
		res = __gong__marshallFloat(ident, "StartRatio", notetaskshape.StartRatio)
	case "EndRatio":
		res = __gong__marshallFloat(ident, "EndRatio", notetaskshape.EndRatio)
	case "StartOrientation":
		res = __gong__marshallEnumString(ident, "StartOrientation", notetaskshape.StartOrientation.ToCodeString())
	case "EndOrientation":
		res = __gong__marshallEnumString(ident, "EndOrientation", notetaskshape.EndOrientation.ToCodeString())
	case "CornerOffsetRatio":
		res = __gong__marshallFloat(ident, "CornerOffsetRatio", notetaskshape.CornerOffsetRatio)
	case "IsHidden":
		res = __gong__marshallBool(ident, "IsHidden", notetaskshape.IsHidden)
	case "IsPresentationOverridden":
		res = __gong__marshallBool(ident, "IsPresentationOverridden", notetaskshape.IsPresentationOverridden)
	case "Stroke":
		res = __gong__marshallString(ident, "Stroke", notetaskshape.Stroke)
	case "StrokeWidth":
		res = __gong__marshallFloat(ident, "StrokeWidth", notetaskshape.StrokeWidth)
	case "StrokeOpacity":
		res = __gong__marshallFloat(ident, "StrokeOpacity", notetaskshape.StrokeOpacity)
	case "StrokeDashArray":
		res = __gong__marshallString(ident, "StrokeDashArray", notetaskshape.StrokeDashArray)
	case "FontSize":
		res = __gong__marshallString(ident, "FontSize", notetaskshape.FontSize)
	case "FontColor":
		res = __gong__marshallString(ident, "FontColor", notetaskshape.FontColor)

	case "Note":
		if notetaskshape.Note != nil {
			res = __gong__marshallPointer(ident, "Note", notetaskshape.Note.GongGetIdentifier(stage))
		} else {
			res = __gong__marshallPointer(ident, "Note", "nil")
		}
	case "Task":
		if notetaskshape.Task != nil {
			res = __gong__marshallPointer(ident, "Task", notetaskshape.Task.GongGetIdentifier(stage))
		} else {
			res = __gong__marshallPointer(ident, "Task", "nil")
		}
	default:
		log.Panicf("Unknown field %s for Gongstruct NoteTaskShape", fieldName)
	}
	return
}

func (product *Product) GongMarshallField(stage *Stage, fieldName string) (res string) {
	ident := product.GongGetIdentifier(stage)
	_ = ident
	switch fieldName {
	case "Name":
		res = __gong__marshallString(ident, "Name", product.Name)
	case "Description":
		res = __gong__marshallString(ident, "Description", product.Description)
	case "IsProducersNodeExpanded":
		res = __gong__marshallBool(ident, "IsProducersNodeExpanded", product.IsProducersNodeExpanded)
	case "IsConsumersNodeExpanded":
		res = __gong__marshallBool(ident, "IsConsumersNodeExpanded", product.IsConsumersNodeExpanded)
	case "IsImport":
		res = __gong__marshallBool(ident, "IsImport", product.IsImport)
	case "ComputedPrefix":
		res = __gong__marshallString(ident, "ComputedPrefix", product.ComputedPrefix)
	case "IsExpanded":
		res = __gong__marshallBool(ident, "IsExpanded", product.IsExpanded)
	case "LayoutDirection":
		res = __gong__marshallEnumInt(ident, "LayoutDirection", product.LayoutDirection.ToCodeString())

	case "SubProducts":
		var sb strings.Builder
		for _, _product := range product.SubProducts {
			sb.WriteString(__gong__marshallSliceOfPointers(ident, "SubProducts", _product.GongGetIdentifier(stage)))
		}
		res = sb.String()
	case "ReferencedProduct":
		if product.ReferencedProduct != nil {
			res = __gong__marshallPointer(ident, "ReferencedProduct", product.ReferencedProduct.GongGetIdentifier(stage))
		} else {
			res = __gong__marshallPointer(ident, "ReferencedProduct", "nil")
		}
	default:
		log.Panicf("Unknown field %s for Gongstruct Product", fieldName)
	}
	return
}

func (productcompositionshape *ProductCompositionShape) GongMarshallField(stage *Stage, fieldName string) (res string) {
	ident := productcompositionshape.GongGetIdentifier(stage)
	_ = ident
	switch fieldName {
	case "Name":
		res = __gong__marshallString(ident, "Name", productcompositionshape.Name)
	case "StartRatio":
		res = __gong__marshallFloat(ident, "StartRatio", productcompositionshape.StartRatio)
	case "EndRatio":
		res = __gong__marshallFloat(ident, "EndRatio", productcompositionshape.EndRatio)
	case "StartOrientation":
		res = __gong__marshallEnumString(ident, "StartOrientation", productcompositionshape.StartOrientation.ToCodeString())
	case "EndOrientation":
		res = __gong__marshallEnumString(ident, "EndOrientation", productcompositionshape.EndOrientation.ToCodeString())
	case "CornerOffsetRatio":
		res = __gong__marshallFloat(ident, "CornerOffsetRatio", productcompositionshape.CornerOffsetRatio)
	case "IsHidden":
		res = __gong__marshallBool(ident, "IsHidden", productcompositionshape.IsHidden)
	case "IsPresentationOverridden":
		res = __gong__marshallBool(ident, "IsPresentationOverridden", productcompositionshape.IsPresentationOverridden)
	case "Stroke":
		res = __gong__marshallString(ident, "Stroke", productcompositionshape.Stroke)
	case "StrokeWidth":
		res = __gong__marshallFloat(ident, "StrokeWidth", productcompositionshape.StrokeWidth)
	case "StrokeOpacity":
		res = __gong__marshallFloat(ident, "StrokeOpacity", productcompositionshape.StrokeOpacity)
	case "StrokeDashArray":
		res = __gong__marshallString(ident, "StrokeDashArray", productcompositionshape.StrokeDashArray)
	case "FontSize":
		res = __gong__marshallString(ident, "FontSize", productcompositionshape.FontSize)
	case "FontColor":
		res = __gong__marshallString(ident, "FontColor", productcompositionshape.FontColor)

	case "Product":
		if productcompositionshape.Product != nil {
			res = __gong__marshallPointer(ident, "Product", productcompositionshape.Product.GongGetIdentifier(stage))
		} else {
			res = __gong__marshallPointer(ident, "Product", "nil")
		}
	default:
		log.Panicf("Unknown field %s for Gongstruct ProductCompositionShape", fieldName)
	}
	return
}

func (productreferenceshape *ProductReferenceShape) GongMarshallField(stage *Stage, fieldName string) (res string) {
	ident := productreferenceshape.GongGetIdentifier(stage)
	_ = ident
	switch fieldName {
	case "Name":
		res = __gong__marshallString(ident, "Name", productreferenceshape.Name)
	case "StartRatio":
		res = __gong__marshallFloat(ident, "StartRatio", productreferenceshape.StartRatio)
	case "EndRatio":
		res = __gong__marshallFloat(ident, "EndRatio", productreferenceshape.EndRatio)
	case "StartOrientation":
		res = __gong__marshallEnumString(ident, "StartOrientation", productreferenceshape.StartOrientation.ToCodeString())
	case "EndOrientation":
		res = __gong__marshallEnumString(ident, "EndOrientation", productreferenceshape.EndOrientation.ToCodeString())
	case "CornerOffsetRatio":
		res = __gong__marshallFloat(ident, "CornerOffsetRatio", productreferenceshape.CornerOffsetRatio)
	case "IsHidden":
		res = __gong__marshallBool(ident, "IsHidden", productreferenceshape.IsHidden)
	case "IsPresentationOverridden":
		res = __gong__marshallBool(ident, "IsPresentationOverridden", productreferenceshape.IsPresentationOverridden)
	case "Stroke":
		res = __gong__marshallString(ident, "Stroke", productreferenceshape.Stroke)
	case "StrokeWidth":
		res = __gong__marshallFloat(ident, "StrokeWidth", productreferenceshape.StrokeWidth)
	case "StrokeOpacity":
		res = __gong__marshallFloat(ident, "StrokeOpacity", productreferenceshape.StrokeOpacity)
	case "StrokeDashArray":
		res = __gong__marshallString(ident, "StrokeDashArray", productreferenceshape.StrokeDashArray)
	case "FontSize":
		res = __gong__marshallString(ident, "FontSize", productreferenceshape.FontSize)
	case "FontColor":
		res = __gong__marshallString(ident, "FontColor", productreferenceshape.FontColor)

	case "Product":
		if productreferenceshape.Product != nil {
			res = __gong__marshallPointer(ident, "Product", productreferenceshape.Product.GongGetIdentifier(stage))
		} else {
			res = __gong__marshallPointer(ident, "Product", "nil")
		}
	case "ReferencedProduct":
		if productreferenceshape.ReferencedProduct != nil {
			res = __gong__marshallPointer(ident, "ReferencedProduct", productreferenceshape.ReferencedProduct.GongGetIdentifier(stage))
		} else {
			res = __gong__marshallPointer(ident, "ReferencedProduct", "nil")
		}
	default:
		log.Panicf("Unknown field %s for Gongstruct ProductReferenceShape", fieldName)
	}
	return
}

func (productshape *ProductShape) GongMarshallField(stage *Stage, fieldName string) (res string) {
	ident := productshape.GongGetIdentifier(stage)
	_ = ident
	switch fieldName {
	case "Name":
		res = __gong__marshallString(ident, "Name", productshape.Name)
	case "IsShowType":
		res = __gong__marshallBool(ident, "IsShowType", productshape.IsShowType)
	case "IsLayoutDirectionDifferent":
		res = __gong__marshallBool(ident, "IsLayoutDirectionDifferent", productshape.IsLayoutDirectionDifferent)
	case "X":
		res = __gong__marshallFloat(ident, "X", productshape.X)
	case "Y":
		res = __gong__marshallFloat(ident, "Y", productshape.Y)
	case "Width":
		res = __gong__marshallFloat(ident, "Width", productshape.Width)
	case "Height":
		res = __gong__marshallFloat(ident, "Height", productshape.Height)
	case "IsHidden":
		res = __gong__marshallBool(ident, "IsHidden", productshape.IsHidden)
	case "IsPresentationOverridden":
		res = __gong__marshallBool(ident, "IsPresentationOverridden", productshape.IsPresentationOverridden)
	case "Color":
		res = __gong__marshallString(ident, "Color", productshape.Color)
	case "FillOpacity":
		res = __gong__marshallFloat(ident, "FillOpacity", productshape.FillOpacity)
	case "Stroke":
		res = __gong__marshallString(ident, "Stroke", productshape.Stroke)
	case "StrokeWidth":
		res = __gong__marshallFloat(ident, "StrokeWidth", productshape.StrokeWidth)
	case "StrokeOpacity":
		res = __gong__marshallFloat(ident, "StrokeOpacity", productshape.StrokeOpacity)
	case "StrokeDashArray":
		res = __gong__marshallString(ident, "StrokeDashArray", productshape.StrokeDashArray)
	case "RX":
		res = __gong__marshallFloat(ident, "RX", productshape.RX)
	case "FontSize":
		res = __gong__marshallString(ident, "FontSize", productshape.FontSize)
	case "FontWeight":
		res = __gong__marshallString(ident, "FontWeight", productshape.FontWeight)
	case "FontColor":
		res = __gong__marshallString(ident, "FontColor", productshape.FontColor)

	case "Product":
		if productshape.Product != nil {
			res = __gong__marshallPointer(ident, "Product", productshape.Product.GongGetIdentifier(stage))
		} else {
			res = __gong__marshallPointer(ident, "Product", "nil")
		}
	default:
		log.Panicf("Unknown field %s for Gongstruct ProductShape", fieldName)
	}
	return
}

func (resource *Resource) GongMarshallField(stage *Stage, fieldName string) (res string) {
	ident := resource.GongGetIdentifier(stage)
	_ = ident
	switch fieldName {
	case "Name":
		res = __gong__marshallString(ident, "Name", resource.Name)
	case "Description":
		res = __gong__marshallString(ident, "Description", resource.Description)
	case "ComputedPrefix":
		res = __gong__marshallString(ident, "ComputedPrefix", resource.ComputedPrefix)
	case "IsExpanded":
		res = __gong__marshallBool(ident, "IsExpanded", resource.IsExpanded)
	case "LayoutDirection":
		res = __gong__marshallEnumInt(ident, "LayoutDirection", resource.LayoutDirection.ToCodeString())
	case "IsImport":
		res = __gong__marshallBool(ident, "IsImport", resource.IsImport)

	case "Tasks":
		var sb strings.Builder
		for _, _task := range resource.Tasks {
			sb.WriteString(__gong__marshallSliceOfPointers(ident, "Tasks", _task.GongGetIdentifier(stage)))
		}
		res = sb.String()
	case "SubResources":
		var sb strings.Builder
		for _, _resource := range resource.SubResources {
			sb.WriteString(__gong__marshallSliceOfPointers(ident, "SubResources", _resource.GongGetIdentifier(stage)))
		}
		res = sb.String()
	case "ReferencedResource":
		if resource.ReferencedResource != nil {
			res = __gong__marshallPointer(ident, "ReferencedResource", resource.ReferencedResource.GongGetIdentifier(stage))
		} else {
			res = __gong__marshallPointer(ident, "ReferencedResource", "nil")
		}
	default:
		log.Panicf("Unknown field %s for Gongstruct Resource", fieldName)
	}
	return
}

func (resourcecompositionshape *ResourceCompositionShape) GongMarshallField(stage *Stage, fieldName string) (res string) {
	ident := resourcecompositionshape.GongGetIdentifier(stage)
	_ = ident
	switch fieldName {
	case "Name":
		res = __gong__marshallString(ident, "Name", resourcecompositionshape.Name)
	case "StartRatio":
		res = __gong__marshallFloat(ident, "StartRatio", resourcecompositionshape.StartRatio)
	case "EndRatio":
		res = __gong__marshallFloat(ident, "EndRatio", resourcecompositionshape.EndRatio)
	case "StartOrientation":
		res = __gong__marshallEnumString(ident, "StartOrientation", resourcecompositionshape.StartOrientation.ToCodeString())
	case "EndOrientation":
		res = __gong__marshallEnumString(ident, "EndOrientation", resourcecompositionshape.EndOrientation.ToCodeString())
	case "CornerOffsetRatio":
		res = __gong__marshallFloat(ident, "CornerOffsetRatio", resourcecompositionshape.CornerOffsetRatio)
	case "IsHidden":
		res = __gong__marshallBool(ident, "IsHidden", resourcecompositionshape.IsHidden)
	case "IsPresentationOverridden":
		res = __gong__marshallBool(ident, "IsPresentationOverridden", resourcecompositionshape.IsPresentationOverridden)
	case "Stroke":
		res = __gong__marshallString(ident, "Stroke", resourcecompositionshape.Stroke)
	case "StrokeWidth":
		res = __gong__marshallFloat(ident, "StrokeWidth", resourcecompositionshape.StrokeWidth)
	case "StrokeOpacity":
		res = __gong__marshallFloat(ident, "StrokeOpacity", resourcecompositionshape.StrokeOpacity)
	case "StrokeDashArray":
		res = __gong__marshallString(ident, "StrokeDashArray", resourcecompositionshape.StrokeDashArray)
	case "FontSize":
		res = __gong__marshallString(ident, "FontSize", resourcecompositionshape.FontSize)
	case "FontColor":
		res = __gong__marshallString(ident, "FontColor", resourcecompositionshape.FontColor)

	case "Resource":
		if resourcecompositionshape.Resource != nil {
			res = __gong__marshallPointer(ident, "Resource", resourcecompositionshape.Resource.GongGetIdentifier(stage))
		} else {
			res = __gong__marshallPointer(ident, "Resource", "nil")
		}
	default:
		log.Panicf("Unknown field %s for Gongstruct ResourceCompositionShape", fieldName)
	}
	return
}

func (resourceshape *ResourceShape) GongMarshallField(stage *Stage, fieldName string) (res string) {
	ident := resourceshape.GongGetIdentifier(stage)
	_ = ident
	switch fieldName {
	case "Name":
		res = __gong__marshallString(ident, "Name", resourceshape.Name)
	case "IsLayoutDirectionDifferent":
		res = __gong__marshallBool(ident, "IsLayoutDirectionDifferent", resourceshape.IsLayoutDirectionDifferent)
	case "X":
		res = __gong__marshallFloat(ident, "X", resourceshape.X)
	case "Y":
		res = __gong__marshallFloat(ident, "Y", resourceshape.Y)
	case "Width":
		res = __gong__marshallFloat(ident, "Width", resourceshape.Width)
	case "Height":
		res = __gong__marshallFloat(ident, "Height", resourceshape.Height)
	case "IsHidden":
		res = __gong__marshallBool(ident, "IsHidden", resourceshape.IsHidden)
	case "IsPresentationOverridden":
		res = __gong__marshallBool(ident, "IsPresentationOverridden", resourceshape.IsPresentationOverridden)
	case "Color":
		res = __gong__marshallString(ident, "Color", resourceshape.Color)
	case "FillOpacity":
		res = __gong__marshallFloat(ident, "FillOpacity", resourceshape.FillOpacity)
	case "Stroke":
		res = __gong__marshallString(ident, "Stroke", resourceshape.Stroke)
	case "StrokeWidth":
		res = __gong__marshallFloat(ident, "StrokeWidth", resourceshape.StrokeWidth)
	case "StrokeOpacity":
		res = __gong__marshallFloat(ident, "StrokeOpacity", resourceshape.StrokeOpacity)
	case "StrokeDashArray":
		res = __gong__marshallString(ident, "StrokeDashArray", resourceshape.StrokeDashArray)
	case "RX":
		res = __gong__marshallFloat(ident, "RX", resourceshape.RX)
	case "FontSize":
		res = __gong__marshallString(ident, "FontSize", resourceshape.FontSize)
	case "FontWeight":
		res = __gong__marshallString(ident, "FontWeight", resourceshape.FontWeight)
	case "FontColor":
		res = __gong__marshallString(ident, "FontColor", resourceshape.FontColor)

	case "Resource":
		if resourceshape.Resource != nil {
			res = __gong__marshallPointer(ident, "Resource", resourceshape.Resource.GongGetIdentifier(stage))
		} else {
			res = __gong__marshallPointer(ident, "Resource", "nil")
		}
	default:
		log.Panicf("Unknown field %s for Gongstruct ResourceShape", fieldName)
	}
	return
}

func (resourcetaskshape *ResourceTaskShape) GongMarshallField(stage *Stage, fieldName string) (res string) {
	ident := resourcetaskshape.GongGetIdentifier(stage)
	_ = ident
	switch fieldName {
	case "Name":
		res = __gong__marshallString(ident, "Name", resourcetaskshape.Name)
	case "StartRatio":
		res = __gong__marshallFloat(ident, "StartRatio", resourcetaskshape.StartRatio)
	case "EndRatio":
		res = __gong__marshallFloat(ident, "EndRatio", resourcetaskshape.EndRatio)
	case "StartOrientation":
		res = __gong__marshallEnumString(ident, "StartOrientation", resourcetaskshape.StartOrientation.ToCodeString())
	case "EndOrientation":
		res = __gong__marshallEnumString(ident, "EndOrientation", resourcetaskshape.EndOrientation.ToCodeString())
	case "CornerOffsetRatio":
		res = __gong__marshallFloat(ident, "CornerOffsetRatio", resourcetaskshape.CornerOffsetRatio)
	case "IsHidden":
		res = __gong__marshallBool(ident, "IsHidden", resourcetaskshape.IsHidden)
	case "IsPresentationOverridden":
		res = __gong__marshallBool(ident, "IsPresentationOverridden", resourcetaskshape.IsPresentationOverridden)
	case "Stroke":
		res = __gong__marshallString(ident, "Stroke", resourcetaskshape.Stroke)
	case "StrokeWidth":
		res = __gong__marshallFloat(ident, "StrokeWidth", resourcetaskshape.StrokeWidth)
	case "StrokeOpacity":
		res = __gong__marshallFloat(ident, "StrokeOpacity", resourcetaskshape.StrokeOpacity)
	case "StrokeDashArray":
		res = __gong__marshallString(ident, "StrokeDashArray", resourcetaskshape.StrokeDashArray)
	case "FontSize":
		res = __gong__marshallString(ident, "FontSize", resourcetaskshape.FontSize)
	case "FontColor":
		res = __gong__marshallString(ident, "FontColor", resourcetaskshape.FontColor)

	case "Resource":
		if resourcetaskshape.Resource != nil {
			res = __gong__marshallPointer(ident, "Resource", resourcetaskshape.Resource.GongGetIdentifier(stage))
		} else {
			res = __gong__marshallPointer(ident, "Resource", "nil")
		}
	case "Task":
		if resourcetaskshape.Task != nil {
			res = __gong__marshallPointer(ident, "Task", resourcetaskshape.Task.GongGetIdentifier(stage))
		} else {
			res = __gong__marshallPointer(ident, "Task", "nil")
		}
	default:
		log.Panicf("Unknown field %s for Gongstruct ResourceTaskShape", fieldName)
	}
	return
}

func (task *Task) GongMarshallField(stage *Stage, fieldName string) (res string) {
	ident := task.GongGetIdentifier(stage)
	_ = ident
	switch fieldName {
	case "Name":
		res = __gong__marshallString(ident, "Name", task.Name)
	case "Description":
		res = __gong__marshallString(ident, "Description", task.Description)
	case "Start":
		res = __gong__marshallTime(ident, "Start", task.Start.String())
	case "End":
		res = __gong__marshallTime(ident, "End", task.End.String())
	case "IsAllDay":
		res = __gong__marshallBool(ident, "IsAllDay", task.IsAllDay)
	case "IsMilestone":
		res = __gong__marshallBool(ident, "IsMilestone", task.IsMilestone)
	case "DependencyType":
		res = __gong__marshallEnumString(ident, "DependencyType", task.DependencyType.ToCodeString())
	case "DependencyDurationYears":
		res = __gong__marshallFloat(ident, "DependencyDurationYears", task.DependencyDurationYears)
	case "DependencyDurationMonths":
		res = __gong__marshallFloat(ident, "DependencyDurationMonths", task.DependencyDurationMonths)
	case "DependencyDurationWeeks":
		res = __gong__marshallFloat(ident, "DependencyDurationWeeks", task.DependencyDurationWeeks)
	case "DependencyDurationDays":
		res = __gong__marshallFloat(ident, "DependencyDurationDays", task.DependencyDurationDays)
	case "DependencyDurationHours":
		res = __gong__marshallFloat(ident, "DependencyDurationHours", task.DependencyDurationHours)
	case "DurationYears":
		res = __gong__marshallFloat(ident, "DurationYears", task.DurationYears)
	case "DurationMonths":
		res = __gong__marshallFloat(ident, "DurationMonths", task.DurationMonths)
	case "DurationWeeks":
		res = __gong__marshallFloat(ident, "DurationWeeks", task.DurationWeeks)
	case "DurationDays":
		res = __gong__marshallFloat(ident, "DurationDays", task.DurationDays)
	case "DurationHours":
		res = __gong__marshallFloat(ident, "DurationHours", task.DurationHours)
	case "IsEndDateComputedFromDuration":
		res = __gong__marshallBool(ident, "IsEndDateComputedFromDuration", task.IsEndDateComputedFromDuration)
	case "IsWithCompletion":
		res = __gong__marshallBool(ident, "IsWithCompletion", task.IsWithCompletion)
	case "Completion":
		res = __gong__marshallEnumString(ident, "Completion", task.Completion.ToCodeString())
	case "TextPosition":
		res = __gong__marshallEnumString(ident, "TextPosition", task.TextPosition.ToCodeString())
	case "XOffset":
		res = __gong__marshallFloat(ident, "XOffset", task.XOffset)
	case "YOffset":
		res = __gong__marshallFloat(ident, "YOffset", task.YOffset)
	case "IsImport":
		res = __gong__marshallBool(ident, "IsImport", task.IsImport)
	case "IsInputsNodeExpanded":
		res = __gong__marshallBool(ident, "IsInputsNodeExpanded", task.IsInputsNodeExpanded)
	case "IsOutputsNodeExpanded":
		res = __gong__marshallBool(ident, "IsOutputsNodeExpanded", task.IsOutputsNodeExpanded)
	case "ComputedPrefix":
		res = __gong__marshallString(ident, "ComputedPrefix", task.ComputedPrefix)
	case "IsExpanded":
		res = __gong__marshallBool(ident, "IsExpanded", task.IsExpanded)
	case "LayoutDirection":
		res = __gong__marshallEnumInt(ident, "LayoutDirection", task.LayoutDirection.ToCodeString())

	case "TaskGroups":
		var sb strings.Builder
		for _, _taskgroup := range task.TaskGroups {
			sb.WriteString(__gong__marshallSliceOfPointers(ident, "TaskGroups", _taskgroup.GongGetIdentifier(stage)))
		}
		res = sb.String()
	case "Predecessors":
		var sb strings.Builder
		for _, _task := range task.Predecessors {
			sb.WriteString(__gong__marshallSliceOfPointers(ident, "Predecessors", _task.GongGetIdentifier(stage)))
		}
		res = sb.String()
	case "Inputs":
		var sb strings.Builder
		for _, _product := range task.Inputs {
			sb.WriteString(__gong__marshallSliceOfPointers(ident, "Inputs", _product.GongGetIdentifier(stage)))
		}
		res = sb.String()
	case "Outputs":
		var sb strings.Builder
		for _, _product := range task.Outputs {
			sb.WriteString(__gong__marshallSliceOfPointers(ident, "Outputs", _product.GongGetIdentifier(stage)))
		}
		res = sb.String()
	case "SubTasks":
		var sb strings.Builder
		for _, _task := range task.SubTasks {
			sb.WriteString(__gong__marshallSliceOfPointers(ident, "SubTasks", _task.GongGetIdentifier(stage)))
		}
		res = sb.String()
	case "TaskGroupsToDisplay":
		var sb strings.Builder
		for _, _taskgroup := range task.TaskGroupsToDisplay {
			sb.WriteString(__gong__marshallSliceOfPointers(ident, "TaskGroupsToDisplay", _taskgroup.GongGetIdentifier(stage)))
		}
		res = sb.String()
	case "ReferencedTask":
		if task.ReferencedTask != nil {
			res = __gong__marshallPointer(ident, "ReferencedTask", task.ReferencedTask.GongGetIdentifier(stage))
		} else {
			res = __gong__marshallPointer(ident, "ReferencedTask", "nil")
		}
	default:
		log.Panicf("Unknown field %s for Gongstruct Task", fieldName)
	}
	return
}

func (taskcompositionshape *TaskCompositionShape) GongMarshallField(stage *Stage, fieldName string) (res string) {
	ident := taskcompositionshape.GongGetIdentifier(stage)
	_ = ident
	switch fieldName {
	case "Name":
		res = __gong__marshallString(ident, "Name", taskcompositionshape.Name)
	case "StartRatio":
		res = __gong__marshallFloat(ident, "StartRatio", taskcompositionshape.StartRatio)
	case "EndRatio":
		res = __gong__marshallFloat(ident, "EndRatio", taskcompositionshape.EndRatio)
	case "StartOrientation":
		res = __gong__marshallEnumString(ident, "StartOrientation", taskcompositionshape.StartOrientation.ToCodeString())
	case "EndOrientation":
		res = __gong__marshallEnumString(ident, "EndOrientation", taskcompositionshape.EndOrientation.ToCodeString())
	case "CornerOffsetRatio":
		res = __gong__marshallFloat(ident, "CornerOffsetRatio", taskcompositionshape.CornerOffsetRatio)
	case "IsHidden":
		res = __gong__marshallBool(ident, "IsHidden", taskcompositionshape.IsHidden)
	case "IsPresentationOverridden":
		res = __gong__marshallBool(ident, "IsPresentationOverridden", taskcompositionshape.IsPresentationOverridden)
	case "Stroke":
		res = __gong__marshallString(ident, "Stroke", taskcompositionshape.Stroke)
	case "StrokeWidth":
		res = __gong__marshallFloat(ident, "StrokeWidth", taskcompositionshape.StrokeWidth)
	case "StrokeOpacity":
		res = __gong__marshallFloat(ident, "StrokeOpacity", taskcompositionshape.StrokeOpacity)
	case "StrokeDashArray":
		res = __gong__marshallString(ident, "StrokeDashArray", taskcompositionshape.StrokeDashArray)
	case "FontSize":
		res = __gong__marshallString(ident, "FontSize", taskcompositionshape.FontSize)
	case "FontColor":
		res = __gong__marshallString(ident, "FontColor", taskcompositionshape.FontColor)

	case "Task":
		if taskcompositionshape.Task != nil {
			res = __gong__marshallPointer(ident, "Task", taskcompositionshape.Task.GongGetIdentifier(stage))
		} else {
			res = __gong__marshallPointer(ident, "Task", "nil")
		}
	default:
		log.Panicf("Unknown field %s for Gongstruct TaskCompositionShape", fieldName)
	}
	return
}

func (taskgroup *TaskGroup) GongMarshallField(stage *Stage, fieldName string) (res string) {
	ident := taskgroup.GongGetIdentifier(stage)
	_ = ident
	switch fieldName {
	case "Name":
		res = __gong__marshallString(ident, "Name", taskgroup.Name)
	case "ComputedPrefix":
		res = __gong__marshallString(ident, "ComputedPrefix", taskgroup.ComputedPrefix)
	case "IsExpanded":
		res = __gong__marshallBool(ident, "IsExpanded", taskgroup.IsExpanded)

	case "Tasks":
		var sb strings.Builder
		for _, _task := range taskgroup.Tasks {
			sb.WriteString(__gong__marshallSliceOfPointers(ident, "Tasks", _task.GongGetIdentifier(stage)))
		}
		res = sb.String()
	default:
		log.Panicf("Unknown field %s for Gongstruct TaskGroup", fieldName)
	}
	return
}

func (taskgroupshape *TaskGroupShape) GongMarshallField(stage *Stage, fieldName string) (res string) {
	ident := taskgroupshape.GongGetIdentifier(stage)
	_ = ident
	switch fieldName {
	case "Name":
		res = __gong__marshallString(ident, "Name", taskgroupshape.Name)
	case "X":
		res = __gong__marshallFloat(ident, "X", taskgroupshape.X)
	case "Y":
		res = __gong__marshallFloat(ident, "Y", taskgroupshape.Y)
	case "Width":
		res = __gong__marshallFloat(ident, "Width", taskgroupshape.Width)
	case "Height":
		res = __gong__marshallFloat(ident, "Height", taskgroupshape.Height)
	case "IsHidden":
		res = __gong__marshallBool(ident, "IsHidden", taskgroupshape.IsHidden)
	case "IsPresentationOverridden":
		res = __gong__marshallBool(ident, "IsPresentationOverridden", taskgroupshape.IsPresentationOverridden)
	case "Color":
		res = __gong__marshallString(ident, "Color", taskgroupshape.Color)
	case "FillOpacity":
		res = __gong__marshallFloat(ident, "FillOpacity", taskgroupshape.FillOpacity)
	case "Stroke":
		res = __gong__marshallString(ident, "Stroke", taskgroupshape.Stroke)
	case "StrokeWidth":
		res = __gong__marshallFloat(ident, "StrokeWidth", taskgroupshape.StrokeWidth)
	case "StrokeOpacity":
		res = __gong__marshallFloat(ident, "StrokeOpacity", taskgroupshape.StrokeOpacity)
	case "StrokeDashArray":
		res = __gong__marshallString(ident, "StrokeDashArray", taskgroupshape.StrokeDashArray)
	case "RX":
		res = __gong__marshallFloat(ident, "RX", taskgroupshape.RX)
	case "FontSize":
		res = __gong__marshallString(ident, "FontSize", taskgroupshape.FontSize)
	case "FontWeight":
		res = __gong__marshallString(ident, "FontWeight", taskgroupshape.FontWeight)
	case "FontColor":
		res = __gong__marshallString(ident, "FontColor", taskgroupshape.FontColor)

	case "TaskGroup":
		if taskgroupshape.TaskGroup != nil {
			res = __gong__marshallPointer(ident, "TaskGroup", taskgroupshape.TaskGroup.GongGetIdentifier(stage))
		} else {
			res = __gong__marshallPointer(ident, "TaskGroup", "nil")
		}
	default:
		log.Panicf("Unknown field %s for Gongstruct TaskGroupShape", fieldName)
	}
	return
}

func (taskinputshape *TaskInputShape) GongMarshallField(stage *Stage, fieldName string) (res string) {
	ident := taskinputshape.GongGetIdentifier(stage)
	_ = ident
	switch fieldName {
	case "Name":
		res = __gong__marshallString(ident, "Name", taskinputshape.Name)
	case "StartRatio":
		res = __gong__marshallFloat(ident, "StartRatio", taskinputshape.StartRatio)
	case "EndRatio":
		res = __gong__marshallFloat(ident, "EndRatio", taskinputshape.EndRatio)
	case "StartOrientation":
		res = __gong__marshallEnumString(ident, "StartOrientation", taskinputshape.StartOrientation.ToCodeString())
	case "EndOrientation":
		res = __gong__marshallEnumString(ident, "EndOrientation", taskinputshape.EndOrientation.ToCodeString())
	case "CornerOffsetRatio":
		res = __gong__marshallFloat(ident, "CornerOffsetRatio", taskinputshape.CornerOffsetRatio)
	case "IsHidden":
		res = __gong__marshallBool(ident, "IsHidden", taskinputshape.IsHidden)
	case "IsPresentationOverridden":
		res = __gong__marshallBool(ident, "IsPresentationOverridden", taskinputshape.IsPresentationOverridden)
	case "Stroke":
		res = __gong__marshallString(ident, "Stroke", taskinputshape.Stroke)
	case "StrokeWidth":
		res = __gong__marshallFloat(ident, "StrokeWidth", taskinputshape.StrokeWidth)
	case "StrokeOpacity":
		res = __gong__marshallFloat(ident, "StrokeOpacity", taskinputshape.StrokeOpacity)
	case "StrokeDashArray":
		res = __gong__marshallString(ident, "StrokeDashArray", taskinputshape.StrokeDashArray)
	case "FontSize":
		res = __gong__marshallString(ident, "FontSize", taskinputshape.FontSize)
	case "FontColor":
		res = __gong__marshallString(ident, "FontColor", taskinputshape.FontColor)

	case "Product":
		if taskinputshape.Product != nil {
			res = __gong__marshallPointer(ident, "Product", taskinputshape.Product.GongGetIdentifier(stage))
		} else {
			res = __gong__marshallPointer(ident, "Product", "nil")
		}
	case "Task":
		if taskinputshape.Task != nil {
			res = __gong__marshallPointer(ident, "Task", taskinputshape.Task.GongGetIdentifier(stage))
		} else {
			res = __gong__marshallPointer(ident, "Task", "nil")
		}
	default:
		log.Panicf("Unknown field %s for Gongstruct TaskInputShape", fieldName)
	}
	return
}

func (taskoutputshape *TaskOutputShape) GongMarshallField(stage *Stage, fieldName string) (res string) {
	ident := taskoutputshape.GongGetIdentifier(stage)
	_ = ident
	switch fieldName {
	case "Name":
		res = __gong__marshallString(ident, "Name", taskoutputshape.Name)
	case "StartRatio":
		res = __gong__marshallFloat(ident, "StartRatio", taskoutputshape.StartRatio)
	case "EndRatio":
		res = __gong__marshallFloat(ident, "EndRatio", taskoutputshape.EndRatio)
	case "StartOrientation":
		res = __gong__marshallEnumString(ident, "StartOrientation", taskoutputshape.StartOrientation.ToCodeString())
	case "EndOrientation":
		res = __gong__marshallEnumString(ident, "EndOrientation", taskoutputshape.EndOrientation.ToCodeString())
	case "CornerOffsetRatio":
		res = __gong__marshallFloat(ident, "CornerOffsetRatio", taskoutputshape.CornerOffsetRatio)
	case "IsHidden":
		res = __gong__marshallBool(ident, "IsHidden", taskoutputshape.IsHidden)
	case "IsPresentationOverridden":
		res = __gong__marshallBool(ident, "IsPresentationOverridden", taskoutputshape.IsPresentationOverridden)
	case "Stroke":
		res = __gong__marshallString(ident, "Stroke", taskoutputshape.Stroke)
	case "StrokeWidth":
		res = __gong__marshallFloat(ident, "StrokeWidth", taskoutputshape.StrokeWidth)
	case "StrokeOpacity":
		res = __gong__marshallFloat(ident, "StrokeOpacity", taskoutputshape.StrokeOpacity)
	case "StrokeDashArray":
		res = __gong__marshallString(ident, "StrokeDashArray", taskoutputshape.StrokeDashArray)
	case "FontSize":
		res = __gong__marshallString(ident, "FontSize", taskoutputshape.FontSize)
	case "FontColor":
		res = __gong__marshallString(ident, "FontColor", taskoutputshape.FontColor)

	case "Task":
		if taskoutputshape.Task != nil {
			res = __gong__marshallPointer(ident, "Task", taskoutputshape.Task.GongGetIdentifier(stage))
		} else {
			res = __gong__marshallPointer(ident, "Task", "nil")
		}
	case "Product":
		if taskoutputshape.Product != nil {
			res = __gong__marshallPointer(ident, "Product", taskoutputshape.Product.GongGetIdentifier(stage))
		} else {
			res = __gong__marshallPointer(ident, "Product", "nil")
		}
	default:
		log.Panicf("Unknown field %s for Gongstruct TaskOutputShape", fieldName)
	}
	return
}

func (taskpredecessorshape *TaskPredecessorShape) GongMarshallField(stage *Stage, fieldName string) (res string) {
	ident := taskpredecessorshape.GongGetIdentifier(stage)
	_ = ident
	switch fieldName {
	case "Name":
		res = __gong__marshallString(ident, "Name", taskpredecessorshape.Name)
	case "StartRatio":
		res = __gong__marshallFloat(ident, "StartRatio", taskpredecessorshape.StartRatio)
	case "EndRatio":
		res = __gong__marshallFloat(ident, "EndRatio", taskpredecessorshape.EndRatio)
	case "StartOrientation":
		res = __gong__marshallEnumString(ident, "StartOrientation", taskpredecessorshape.StartOrientation.ToCodeString())
	case "EndOrientation":
		res = __gong__marshallEnumString(ident, "EndOrientation", taskpredecessorshape.EndOrientation.ToCodeString())
	case "CornerOffsetRatio":
		res = __gong__marshallFloat(ident, "CornerOffsetRatio", taskpredecessorshape.CornerOffsetRatio)
	case "IsHidden":
		res = __gong__marshallBool(ident, "IsHidden", taskpredecessorshape.IsHidden)
	case "IsPresentationOverridden":
		res = __gong__marshallBool(ident, "IsPresentationOverridden", taskpredecessorshape.IsPresentationOverridden)
	case "Stroke":
		res = __gong__marshallString(ident, "Stroke", taskpredecessorshape.Stroke)
	case "StrokeWidth":
		res = __gong__marshallFloat(ident, "StrokeWidth", taskpredecessorshape.StrokeWidth)
	case "StrokeOpacity":
		res = __gong__marshallFloat(ident, "StrokeOpacity", taskpredecessorshape.StrokeOpacity)
	case "StrokeDashArray":
		res = __gong__marshallString(ident, "StrokeDashArray", taskpredecessorshape.StrokeDashArray)
	case "FontSize":
		res = __gong__marshallString(ident, "FontSize", taskpredecessorshape.FontSize)
	case "FontColor":
		res = __gong__marshallString(ident, "FontColor", taskpredecessorshape.FontColor)

	case "Predecessor":
		if taskpredecessorshape.Predecessor != nil {
			res = __gong__marshallPointer(ident, "Predecessor", taskpredecessorshape.Predecessor.GongGetIdentifier(stage))
		} else {
			res = __gong__marshallPointer(ident, "Predecessor", "nil")
		}
	case "Task":
		if taskpredecessorshape.Task != nil {
			res = __gong__marshallPointer(ident, "Task", taskpredecessorshape.Task.GongGetIdentifier(stage))
		} else {
			res = __gong__marshallPointer(ident, "Task", "nil")
		}
	default:
		log.Panicf("Unknown field %s for Gongstruct TaskPredecessorShape", fieldName)
	}
	return
}

func (taskshape *TaskShape) GongMarshallField(stage *Stage, fieldName string) (res string) {
	ident := taskshape.GongGetIdentifier(stage)
	_ = ident
	switch fieldName {
	case "Name":
		res = __gong__marshallString(ident, "Name", taskshape.Name)
	case "IsShowDate":
		res = __gong__marshallBool(ident, "IsShowDate", taskshape.IsShowDate)
	case "VerticalOffset":
		res = __gong__marshallFloat(ident, "VerticalOffset", taskshape.VerticalOffset)
	case "DisplayVerticalBar":
		res = __gong__marshallBool(ident, "DisplayVerticalBar", taskshape.DisplayVerticalBar)
	case "IsLayoutDirectionDifferent":
		res = __gong__marshallBool(ident, "IsLayoutDirectionDifferent", taskshape.IsLayoutDirectionDifferent)
	case "X":
		res = __gong__marshallFloat(ident, "X", taskshape.X)
	case "Y":
		res = __gong__marshallFloat(ident, "Y", taskshape.Y)
	case "Width":
		res = __gong__marshallFloat(ident, "Width", taskshape.Width)
	case "Height":
		res = __gong__marshallFloat(ident, "Height", taskshape.Height)
	case "IsHidden":
		res = __gong__marshallBool(ident, "IsHidden", taskshape.IsHidden)
	case "IsPresentationOverridden":
		res = __gong__marshallBool(ident, "IsPresentationOverridden", taskshape.IsPresentationOverridden)
	case "Color":
		res = __gong__marshallString(ident, "Color", taskshape.Color)
	case "FillOpacity":
		res = __gong__marshallFloat(ident, "FillOpacity", taskshape.FillOpacity)
	case "Stroke":
		res = __gong__marshallString(ident, "Stroke", taskshape.Stroke)
	case "StrokeWidth":
		res = __gong__marshallFloat(ident, "StrokeWidth", taskshape.StrokeWidth)
	case "StrokeOpacity":
		res = __gong__marshallFloat(ident, "StrokeOpacity", taskshape.StrokeOpacity)
	case "StrokeDashArray":
		res = __gong__marshallString(ident, "StrokeDashArray", taskshape.StrokeDashArray)
	case "RX":
		res = __gong__marshallFloat(ident, "RX", taskshape.RX)
	case "FontSize":
		res = __gong__marshallString(ident, "FontSize", taskshape.FontSize)
	case "FontWeight":
		res = __gong__marshallString(ident, "FontWeight", taskshape.FontWeight)
	case "FontColor":
		res = __gong__marshallString(ident, "FontColor", taskshape.FontColor)

	case "Task":
		if taskshape.Task != nil {
			res = __gong__marshallPointer(ident, "Task", taskshape.Task.GongGetIdentifier(stage))
		} else {
			res = __gong__marshallPointer(ident, "Task", "nil")
		}
	default:
		log.Panicf("Unknown field %s for Gongstruct TaskShape", fieldName)
	}
	return
}

// insertion point for marshall all fields methods
func (diagram *Diagram) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(diagram.GongMarshallField(stage, "Name"))
		initializerStatements.WriteString(diagram.GongMarshallField(stage, "DefaultBoxWidth"))
		initializerStatements.WriteString(diagram.GongMarshallField(stage, "DefaultBoxHeigth"))
		initializerStatements.WriteString(diagram.GongMarshallField(stage, "DateFormat"))
		initializerStatements.WriteString(diagram.GongMarshallField(stage, "Width"))
		initializerStatements.WriteString(diagram.GongMarshallField(stage, "Height"))
		initializerStatements.WriteString(diagram.GongMarshallField(stage, "IsTimeDiagram"))
		initializerStatements.WriteString(diagram.GongMarshallField(stage, "ComputedStart"))
		initializerStatements.WriteString(diagram.GongMarshallField(stage, "ComputedEnd"))
		initializerStatements.WriteString(diagram.GongMarshallField(stage, "ComputedDuration"))
		initializerStatements.WriteString(diagram.GongMarshallField(stage, "DrawVerticalTimeLines"))
		initializerStatements.WriteString(diagram.GongMarshallField(stage, "DrawSecondaryVerticalTimeLines"))
		initializerStatements.WriteString(diagram.GongMarshallField(stage, "HideWeekendsPeriod"))
		initializerStatements.WriteString(diagram.GongMarshallField(stage, "UseManualStartAndEndDates"))
		initializerStatements.WriteString(diagram.GongMarshallField(stage, "ManualStart"))
		initializerStatements.WriteString(diagram.GongMarshallField(stage, "ManualEnd"))
		initializerStatements.WriteString(diagram.GongMarshallField(stage, "TimeStep"))
		initializerStatements.WriteString(diagram.GongMarshallField(stage, "TimeStepScale"))
		initializerStatements.WriteString(diagram.GongMarshallField(stage, "SecondaryTimeStep"))
		initializerStatements.WriteString(diagram.GongMarshallField(stage, "SecondaryTimeStepScale"))
		initializerStatements.WriteString(diagram.GongMarshallField(stage, "LaneHeight"))
		initializerStatements.WriteString(diagram.GongMarshallField(stage, "RatioBarToLaneHeight"))
		initializerStatements.WriteString(diagram.GongMarshallField(stage, "YTopMargin"))
		initializerStatements.WriteString(diagram.GongMarshallField(stage, "XLeftText"))
		initializerStatements.WriteString(diagram.GongMarshallField(stage, "TextHeight"))
		initializerStatements.WriteString(diagram.GongMarshallField(stage, "XLeftLanes"))
		initializerStatements.WriteString(diagram.GongMarshallField(stage, "XRightMargin"))
		initializerStatements.WriteString(diagram.GongMarshallField(stage, "ArrowLengthToTheRightOfStartBar"))
		initializerStatements.WriteString(diagram.GongMarshallField(stage, "ArrowTipLenght"))
		initializerStatements.WriteString(diagram.GongMarshallField(stage, "TimeLine_Color"))
		initializerStatements.WriteString(diagram.GongMarshallField(stage, "TimeLine_FillOpacity"))
		initializerStatements.WriteString(diagram.GongMarshallField(stage, "TimeLine_Stroke"))
		initializerStatements.WriteString(diagram.GongMarshallField(stage, "TimeLine_StrokeWidth"))
		initializerStatements.WriteString(diagram.GongMarshallField(stage, "Group_Stroke"))
		initializerStatements.WriteString(diagram.GongMarshallField(stage, "Group_StrokeWidth"))
		initializerStatements.WriteString(diagram.GongMarshallField(stage, "Group_StrokeDashArray"))
		initializerStatements.WriteString(diagram.GongMarshallField(stage, "DateYOffset"))
		initializerStatements.WriteString(diagram.GongMarshallField(stage, "AlignOnBeginningOfTimeScale"))
		initializerStatements.WriteString(diagram.GongMarshallField(stage, "ComputedPrefix"))
		initializerStatements.WriteString(diagram.GongMarshallField(stage, "IsExpanded"))
		initializerStatements.WriteString(diagram.GongMarshallField(stage, "IsChecked"))
		initializerStatements.WriteString(diagram.GongMarshallField(stage, "IsEditable_"))
		initializerStatements.WriteString(diagram.GongMarshallField(stage, "IsShowPrefix"))
		initializerStatements.WriteString(diagram.GongMarshallField(stage, "IsInAutoLayoutMode"))
		pointersInitializesStatements.WriteString(diagram.GongMarshallField(stage, "Product_Shapes"))
		pointersInitializesStatements.WriteString(diagram.GongMarshallField(stage, "ProductsWhoseNodeIsExpanded"))
		initializerStatements.WriteString(diagram.GongMarshallField(stage, "IsPBSNodeExpanded"))
		pointersInitializesStatements.WriteString(diagram.GongMarshallField(stage, "ProductComposition_Shapes"))
		pointersInitializesStatements.WriteString(diagram.GongMarshallField(stage, "ProductReference_Shapes"))
		initializerStatements.WriteString(diagram.GongMarshallField(stage, "IsWBSNodeExpanded"))
		pointersInitializesStatements.WriteString(diagram.GongMarshallField(stage, "Task_Shapes"))
		pointersInitializesStatements.WriteString(diagram.GongMarshallField(stage, "TasksWhoseNodeIsExpanded"))
		pointersInitializesStatements.WriteString(diagram.GongMarshallField(stage, "TasksWhoseInputNodeIsExpanded"))
		pointersInitializesStatements.WriteString(diagram.GongMarshallField(stage, "TasksWhoseOutputNodeIsExpanded"))
		pointersInitializesStatements.WriteString(diagram.GongMarshallField(stage, "TasksWhosePredecessorNodeIsExpanded"))
		initializerStatements.WriteString(diagram.GongMarshallField(stage, "IsTaskGroupsNodeExpanded"))
		pointersInitializesStatements.WriteString(diagram.GongMarshallField(stage, "TaskGroupShapes"))
		pointersInitializesStatements.WriteString(diagram.GongMarshallField(stage, "TaskGroupsWhoseNodeIsExpanded"))
		pointersInitializesStatements.WriteString(diagram.GongMarshallField(stage, "TaskComposition_Shapes"))
		pointersInitializesStatements.WriteString(diagram.GongMarshallField(stage, "TaskInputShapes"))
		pointersInitializesStatements.WriteString(diagram.GongMarshallField(stage, "TaskOutputShapes"))
		pointersInitializesStatements.WriteString(diagram.GongMarshallField(stage, "TaskPredecessorShapes"))
		pointersInitializesStatements.WriteString(diagram.GongMarshallField(stage, "Note_Shapes"))
		pointersInitializesStatements.WriteString(diagram.GongMarshallField(stage, "NotesWhoseNodeIsExpanded"))
		initializerStatements.WriteString(diagram.GongMarshallField(stage, "IsNotesNodeExpanded"))
		pointersInitializesStatements.WriteString(diagram.GongMarshallField(stage, "NoteProductShapes"))
		pointersInitializesStatements.WriteString(diagram.GongMarshallField(stage, "NoteTaskShapes"))
		pointersInitializesStatements.WriteString(diagram.GongMarshallField(stage, "NoteResourceShapes"))
		pointersInitializesStatements.WriteString(diagram.GongMarshallField(stage, "Resource_Shapes"))
		pointersInitializesStatements.WriteString(diagram.GongMarshallField(stage, "ResourcesWhoseNodeIsExpanded"))
		initializerStatements.WriteString(diagram.GongMarshallField(stage, "IsResourcesNodeExpanded"))
		pointersInitializesStatements.WriteString(diagram.GongMarshallField(stage, "ResourceComposition_Shapes"))
		pointersInitializesStatements.WriteString(diagram.GongMarshallField(stage, "ResourceTaskShapes"))
	}
	initRes = initializerStatements.String()
	ptrRes = pointersInitializesStatements.String()
	return
}
func (library *Library) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(library.GongMarshallField(stage, "Name"))
		pointersInitializesStatements.WriteString(library.GongMarshallField(stage, "SubLibraries"))
		initializerStatements.WriteString(library.GongMarshallField(stage, "NbPixPerCharacter"))
		initializerStatements.WriteString(library.GongMarshallField(stage, "LogoSVGFile"))
		initializerStatements.WriteString(library.GongMarshallField(stage, "ComputedPrefix"))
		initializerStatements.WriteString(library.GongMarshallField(stage, "IsExpanded"))
		initializerStatements.WriteString(library.GongMarshallField(stage, "IsRootLibrary"))
		pointersInitializesStatements.WriteString(library.GongMarshallField(stage, "RootProducts"))
		pointersInitializesStatements.WriteString(library.GongMarshallField(stage, "RootTasks"))
		pointersInitializesStatements.WriteString(library.GongMarshallField(stage, "RootTaskGroups"))
		pointersInitializesStatements.WriteString(library.GongMarshallField(stage, "RootResources"))
		pointersInitializesStatements.WriteString(library.GongMarshallField(stage, "Notes"))
		pointersInitializesStatements.WriteString(library.GongMarshallField(stage, "Diagrams"))
	}
	initRes = initializerStatements.String()
	ptrRes = pointersInitializesStatements.String()
	return
}
func (note *Note) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(note.GongMarshallField(stage, "Name"))
		initializerStatements.WriteString(note.GongMarshallField(stage, "ComputedPrefix"))
		initializerStatements.WriteString(note.GongMarshallField(stage, "IsExpanded"))
		initializerStatements.WriteString(note.GongMarshallField(stage, "LayoutDirection"))
		pointersInitializesStatements.WriteString(note.GongMarshallField(stage, "Products"))
		pointersInitializesStatements.WriteString(note.GongMarshallField(stage, "Tasks"))
		pointersInitializesStatements.WriteString(note.GongMarshallField(stage, "Resources"))
	}
	initRes = initializerStatements.String()
	ptrRes = pointersInitializesStatements.String()
	return
}
func (noteproductshape *NoteProductShape) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(noteproductshape.GongMarshallField(stage, "Name"))
		pointersInitializesStatements.WriteString(noteproductshape.GongMarshallField(stage, "Note"))
		pointersInitializesStatements.WriteString(noteproductshape.GongMarshallField(stage, "Product"))
		initializerStatements.WriteString(noteproductshape.GongMarshallField(stage, "StartRatio"))
		initializerStatements.WriteString(noteproductshape.GongMarshallField(stage, "EndRatio"))
		initializerStatements.WriteString(noteproductshape.GongMarshallField(stage, "StartOrientation"))
		initializerStatements.WriteString(noteproductshape.GongMarshallField(stage, "EndOrientation"))
		initializerStatements.WriteString(noteproductshape.GongMarshallField(stage, "CornerOffsetRatio"))
		initializerStatements.WriteString(noteproductshape.GongMarshallField(stage, "IsHidden"))
		initializerStatements.WriteString(noteproductshape.GongMarshallField(stage, "IsPresentationOverridden"))
		initializerStatements.WriteString(noteproductshape.GongMarshallField(stage, "Stroke"))
		initializerStatements.WriteString(noteproductshape.GongMarshallField(stage, "StrokeWidth"))
		initializerStatements.WriteString(noteproductshape.GongMarshallField(stage, "StrokeOpacity"))
		initializerStatements.WriteString(noteproductshape.GongMarshallField(stage, "StrokeDashArray"))
		initializerStatements.WriteString(noteproductshape.GongMarshallField(stage, "FontSize"))
		initializerStatements.WriteString(noteproductshape.GongMarshallField(stage, "FontColor"))
	}
	initRes = initializerStatements.String()
	ptrRes = pointersInitializesStatements.String()
	return
}
func (noteresourceshape *NoteResourceShape) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(noteresourceshape.GongMarshallField(stage, "Name"))
		pointersInitializesStatements.WriteString(noteresourceshape.GongMarshallField(stage, "Note"))
		pointersInitializesStatements.WriteString(noteresourceshape.GongMarshallField(stage, "Resource"))
		initializerStatements.WriteString(noteresourceshape.GongMarshallField(stage, "StartRatio"))
		initializerStatements.WriteString(noteresourceshape.GongMarshallField(stage, "EndRatio"))
		initializerStatements.WriteString(noteresourceshape.GongMarshallField(stage, "StartOrientation"))
		initializerStatements.WriteString(noteresourceshape.GongMarshallField(stage, "EndOrientation"))
		initializerStatements.WriteString(noteresourceshape.GongMarshallField(stage, "CornerOffsetRatio"))
		initializerStatements.WriteString(noteresourceshape.GongMarshallField(stage, "IsHidden"))
		initializerStatements.WriteString(noteresourceshape.GongMarshallField(stage, "IsPresentationOverridden"))
		initializerStatements.WriteString(noteresourceshape.GongMarshallField(stage, "Stroke"))
		initializerStatements.WriteString(noteresourceshape.GongMarshallField(stage, "StrokeWidth"))
		initializerStatements.WriteString(noteresourceshape.GongMarshallField(stage, "StrokeOpacity"))
		initializerStatements.WriteString(noteresourceshape.GongMarshallField(stage, "StrokeDashArray"))
		initializerStatements.WriteString(noteresourceshape.GongMarshallField(stage, "FontSize"))
		initializerStatements.WriteString(noteresourceshape.GongMarshallField(stage, "FontColor"))
	}
	initRes = initializerStatements.String()
	ptrRes = pointersInitializesStatements.String()
	return
}
func (noteshape *NoteShape) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(noteshape.GongMarshallField(stage, "Name"))
		pointersInitializesStatements.WriteString(noteshape.GongMarshallField(stage, "Note"))
		initializerStatements.WriteString(noteshape.GongMarshallField(stage, "IsLayoutDirectionDifferent"))
		initializerStatements.WriteString(noteshape.GongMarshallField(stage, "X"))
		initializerStatements.WriteString(noteshape.GongMarshallField(stage, "Y"))
		initializerStatements.WriteString(noteshape.GongMarshallField(stage, "Width"))
		initializerStatements.WriteString(noteshape.GongMarshallField(stage, "Height"))
		initializerStatements.WriteString(noteshape.GongMarshallField(stage, "IsHidden"))
		initializerStatements.WriteString(noteshape.GongMarshallField(stage, "IsPresentationOverridden"))
		initializerStatements.WriteString(noteshape.GongMarshallField(stage, "Color"))
		initializerStatements.WriteString(noteshape.GongMarshallField(stage, "FillOpacity"))
		initializerStatements.WriteString(noteshape.GongMarshallField(stage, "Stroke"))
		initializerStatements.WriteString(noteshape.GongMarshallField(stage, "StrokeWidth"))
		initializerStatements.WriteString(noteshape.GongMarshallField(stage, "StrokeOpacity"))
		initializerStatements.WriteString(noteshape.GongMarshallField(stage, "StrokeDashArray"))
		initializerStatements.WriteString(noteshape.GongMarshallField(stage, "RX"))
		initializerStatements.WriteString(noteshape.GongMarshallField(stage, "FontSize"))
		initializerStatements.WriteString(noteshape.GongMarshallField(stage, "FontWeight"))
		initializerStatements.WriteString(noteshape.GongMarshallField(stage, "FontColor"))
	}
	initRes = initializerStatements.String()
	ptrRes = pointersInitializesStatements.String()
	return
}
func (notetaskshape *NoteTaskShape) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(notetaskshape.GongMarshallField(stage, "Name"))
		pointersInitializesStatements.WriteString(notetaskshape.GongMarshallField(stage, "Note"))
		pointersInitializesStatements.WriteString(notetaskshape.GongMarshallField(stage, "Task"))
		initializerStatements.WriteString(notetaskshape.GongMarshallField(stage, "StartRatio"))
		initializerStatements.WriteString(notetaskshape.GongMarshallField(stage, "EndRatio"))
		initializerStatements.WriteString(notetaskshape.GongMarshallField(stage, "StartOrientation"))
		initializerStatements.WriteString(notetaskshape.GongMarshallField(stage, "EndOrientation"))
		initializerStatements.WriteString(notetaskshape.GongMarshallField(stage, "CornerOffsetRatio"))
		initializerStatements.WriteString(notetaskshape.GongMarshallField(stage, "IsHidden"))
		initializerStatements.WriteString(notetaskshape.GongMarshallField(stage, "IsPresentationOverridden"))
		initializerStatements.WriteString(notetaskshape.GongMarshallField(stage, "Stroke"))
		initializerStatements.WriteString(notetaskshape.GongMarshallField(stage, "StrokeWidth"))
		initializerStatements.WriteString(notetaskshape.GongMarshallField(stage, "StrokeOpacity"))
		initializerStatements.WriteString(notetaskshape.GongMarshallField(stage, "StrokeDashArray"))
		initializerStatements.WriteString(notetaskshape.GongMarshallField(stage, "FontSize"))
		initializerStatements.WriteString(notetaskshape.GongMarshallField(stage, "FontColor"))
	}
	initRes = initializerStatements.String()
	ptrRes = pointersInitializesStatements.String()
	return
}
func (product *Product) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(product.GongMarshallField(stage, "Name"))
		initializerStatements.WriteString(product.GongMarshallField(stage, "Description"))
		pointersInitializesStatements.WriteString(product.GongMarshallField(stage, "SubProducts"))
		initializerStatements.WriteString(product.GongMarshallField(stage, "IsProducersNodeExpanded"))
		initializerStatements.WriteString(product.GongMarshallField(stage, "IsConsumersNodeExpanded"))
		initializerStatements.WriteString(product.GongMarshallField(stage, "IsImport"))
		pointersInitializesStatements.WriteString(product.GongMarshallField(stage, "ReferencedProduct"))
		initializerStatements.WriteString(product.GongMarshallField(stage, "ComputedPrefix"))
		initializerStatements.WriteString(product.GongMarshallField(stage, "IsExpanded"))
		initializerStatements.WriteString(product.GongMarshallField(stage, "LayoutDirection"))
	}
	initRes = initializerStatements.String()
	ptrRes = pointersInitializesStatements.String()
	return
}
func (productcompositionshape *ProductCompositionShape) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(productcompositionshape.GongMarshallField(stage, "Name"))
		pointersInitializesStatements.WriteString(productcompositionshape.GongMarshallField(stage, "Product"))
		initializerStatements.WriteString(productcompositionshape.GongMarshallField(stage, "StartRatio"))
		initializerStatements.WriteString(productcompositionshape.GongMarshallField(stage, "EndRatio"))
		initializerStatements.WriteString(productcompositionshape.GongMarshallField(stage, "StartOrientation"))
		initializerStatements.WriteString(productcompositionshape.GongMarshallField(stage, "EndOrientation"))
		initializerStatements.WriteString(productcompositionshape.GongMarshallField(stage, "CornerOffsetRatio"))
		initializerStatements.WriteString(productcompositionshape.GongMarshallField(stage, "IsHidden"))
		initializerStatements.WriteString(productcompositionshape.GongMarshallField(stage, "IsPresentationOverridden"))
		initializerStatements.WriteString(productcompositionshape.GongMarshallField(stage, "Stroke"))
		initializerStatements.WriteString(productcompositionshape.GongMarshallField(stage, "StrokeWidth"))
		initializerStatements.WriteString(productcompositionshape.GongMarshallField(stage, "StrokeOpacity"))
		initializerStatements.WriteString(productcompositionshape.GongMarshallField(stage, "StrokeDashArray"))
		initializerStatements.WriteString(productcompositionshape.GongMarshallField(stage, "FontSize"))
		initializerStatements.WriteString(productcompositionshape.GongMarshallField(stage, "FontColor"))
	}
	initRes = initializerStatements.String()
	ptrRes = pointersInitializesStatements.String()
	return
}
func (productreferenceshape *ProductReferenceShape) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(productreferenceshape.GongMarshallField(stage, "Name"))
		pointersInitializesStatements.WriteString(productreferenceshape.GongMarshallField(stage, "Product"))
		pointersInitializesStatements.WriteString(productreferenceshape.GongMarshallField(stage, "ReferencedProduct"))
		initializerStatements.WriteString(productreferenceshape.GongMarshallField(stage, "StartRatio"))
		initializerStatements.WriteString(productreferenceshape.GongMarshallField(stage, "EndRatio"))
		initializerStatements.WriteString(productreferenceshape.GongMarshallField(stage, "StartOrientation"))
		initializerStatements.WriteString(productreferenceshape.GongMarshallField(stage, "EndOrientation"))
		initializerStatements.WriteString(productreferenceshape.GongMarshallField(stage, "CornerOffsetRatio"))
		initializerStatements.WriteString(productreferenceshape.GongMarshallField(stage, "IsHidden"))
		initializerStatements.WriteString(productreferenceshape.GongMarshallField(stage, "IsPresentationOverridden"))
		initializerStatements.WriteString(productreferenceshape.GongMarshallField(stage, "Stroke"))
		initializerStatements.WriteString(productreferenceshape.GongMarshallField(stage, "StrokeWidth"))
		initializerStatements.WriteString(productreferenceshape.GongMarshallField(stage, "StrokeOpacity"))
		initializerStatements.WriteString(productreferenceshape.GongMarshallField(stage, "StrokeDashArray"))
		initializerStatements.WriteString(productreferenceshape.GongMarshallField(stage, "FontSize"))
		initializerStatements.WriteString(productreferenceshape.GongMarshallField(stage, "FontColor"))
	}
	initRes = initializerStatements.String()
	ptrRes = pointersInitializesStatements.String()
	return
}
func (productshape *ProductShape) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(productshape.GongMarshallField(stage, "Name"))
		pointersInitializesStatements.WriteString(productshape.GongMarshallField(stage, "Product"))
		initializerStatements.WriteString(productshape.GongMarshallField(stage, "IsShowType"))
		initializerStatements.WriteString(productshape.GongMarshallField(stage, "IsLayoutDirectionDifferent"))
		initializerStatements.WriteString(productshape.GongMarshallField(stage, "X"))
		initializerStatements.WriteString(productshape.GongMarshallField(stage, "Y"))
		initializerStatements.WriteString(productshape.GongMarshallField(stage, "Width"))
		initializerStatements.WriteString(productshape.GongMarshallField(stage, "Height"))
		initializerStatements.WriteString(productshape.GongMarshallField(stage, "IsHidden"))
		initializerStatements.WriteString(productshape.GongMarshallField(stage, "IsPresentationOverridden"))
		initializerStatements.WriteString(productshape.GongMarshallField(stage, "Color"))
		initializerStatements.WriteString(productshape.GongMarshallField(stage, "FillOpacity"))
		initializerStatements.WriteString(productshape.GongMarshallField(stage, "Stroke"))
		initializerStatements.WriteString(productshape.GongMarshallField(stage, "StrokeWidth"))
		initializerStatements.WriteString(productshape.GongMarshallField(stage, "StrokeOpacity"))
		initializerStatements.WriteString(productshape.GongMarshallField(stage, "StrokeDashArray"))
		initializerStatements.WriteString(productshape.GongMarshallField(stage, "RX"))
		initializerStatements.WriteString(productshape.GongMarshallField(stage, "FontSize"))
		initializerStatements.WriteString(productshape.GongMarshallField(stage, "FontWeight"))
		initializerStatements.WriteString(productshape.GongMarshallField(stage, "FontColor"))
	}
	initRes = initializerStatements.String()
	ptrRes = pointersInitializesStatements.String()
	return
}
func (resource *Resource) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(resource.GongMarshallField(stage, "Name"))
		initializerStatements.WriteString(resource.GongMarshallField(stage, "Description"))
		pointersInitializesStatements.WriteString(resource.GongMarshallField(stage, "Tasks"))
		pointersInitializesStatements.WriteString(resource.GongMarshallField(stage, "SubResources"))
		initializerStatements.WriteString(resource.GongMarshallField(stage, "ComputedPrefix"))
		initializerStatements.WriteString(resource.GongMarshallField(stage, "IsExpanded"))
		initializerStatements.WriteString(resource.GongMarshallField(stage, "LayoutDirection"))
		initializerStatements.WriteString(resource.GongMarshallField(stage, "IsImport"))
		pointersInitializesStatements.WriteString(resource.GongMarshallField(stage, "ReferencedResource"))
	}
	initRes = initializerStatements.String()
	ptrRes = pointersInitializesStatements.String()
	return
}
func (resourcecompositionshape *ResourceCompositionShape) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(resourcecompositionshape.GongMarshallField(stage, "Name"))
		pointersInitializesStatements.WriteString(resourcecompositionshape.GongMarshallField(stage, "Resource"))
		initializerStatements.WriteString(resourcecompositionshape.GongMarshallField(stage, "StartRatio"))
		initializerStatements.WriteString(resourcecompositionshape.GongMarshallField(stage, "EndRatio"))
		initializerStatements.WriteString(resourcecompositionshape.GongMarshallField(stage, "StartOrientation"))
		initializerStatements.WriteString(resourcecompositionshape.GongMarshallField(stage, "EndOrientation"))
		initializerStatements.WriteString(resourcecompositionshape.GongMarshallField(stage, "CornerOffsetRatio"))
		initializerStatements.WriteString(resourcecompositionshape.GongMarshallField(stage, "IsHidden"))
		initializerStatements.WriteString(resourcecompositionshape.GongMarshallField(stage, "IsPresentationOverridden"))
		initializerStatements.WriteString(resourcecompositionshape.GongMarshallField(stage, "Stroke"))
		initializerStatements.WriteString(resourcecompositionshape.GongMarshallField(stage, "StrokeWidth"))
		initializerStatements.WriteString(resourcecompositionshape.GongMarshallField(stage, "StrokeOpacity"))
		initializerStatements.WriteString(resourcecompositionshape.GongMarshallField(stage, "StrokeDashArray"))
		initializerStatements.WriteString(resourcecompositionshape.GongMarshallField(stage, "FontSize"))
		initializerStatements.WriteString(resourcecompositionshape.GongMarshallField(stage, "FontColor"))
	}
	initRes = initializerStatements.String()
	ptrRes = pointersInitializesStatements.String()
	return
}
func (resourceshape *ResourceShape) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(resourceshape.GongMarshallField(stage, "Name"))
		pointersInitializesStatements.WriteString(resourceshape.GongMarshallField(stage, "Resource"))
		initializerStatements.WriteString(resourceshape.GongMarshallField(stage, "IsLayoutDirectionDifferent"))
		initializerStatements.WriteString(resourceshape.GongMarshallField(stage, "X"))
		initializerStatements.WriteString(resourceshape.GongMarshallField(stage, "Y"))
		initializerStatements.WriteString(resourceshape.GongMarshallField(stage, "Width"))
		initializerStatements.WriteString(resourceshape.GongMarshallField(stage, "Height"))
		initializerStatements.WriteString(resourceshape.GongMarshallField(stage, "IsHidden"))
		initializerStatements.WriteString(resourceshape.GongMarshallField(stage, "IsPresentationOverridden"))
		initializerStatements.WriteString(resourceshape.GongMarshallField(stage, "Color"))
		initializerStatements.WriteString(resourceshape.GongMarshallField(stage, "FillOpacity"))
		initializerStatements.WriteString(resourceshape.GongMarshallField(stage, "Stroke"))
		initializerStatements.WriteString(resourceshape.GongMarshallField(stage, "StrokeWidth"))
		initializerStatements.WriteString(resourceshape.GongMarshallField(stage, "StrokeOpacity"))
		initializerStatements.WriteString(resourceshape.GongMarshallField(stage, "StrokeDashArray"))
		initializerStatements.WriteString(resourceshape.GongMarshallField(stage, "RX"))
		initializerStatements.WriteString(resourceshape.GongMarshallField(stage, "FontSize"))
		initializerStatements.WriteString(resourceshape.GongMarshallField(stage, "FontWeight"))
		initializerStatements.WriteString(resourceshape.GongMarshallField(stage, "FontColor"))
	}
	initRes = initializerStatements.String()
	ptrRes = pointersInitializesStatements.String()
	return
}
func (resourcetaskshape *ResourceTaskShape) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(resourcetaskshape.GongMarshallField(stage, "Name"))
		pointersInitializesStatements.WriteString(resourcetaskshape.GongMarshallField(stage, "Resource"))
		pointersInitializesStatements.WriteString(resourcetaskshape.GongMarshallField(stage, "Task"))
		initializerStatements.WriteString(resourcetaskshape.GongMarshallField(stage, "StartRatio"))
		initializerStatements.WriteString(resourcetaskshape.GongMarshallField(stage, "EndRatio"))
		initializerStatements.WriteString(resourcetaskshape.GongMarshallField(stage, "StartOrientation"))
		initializerStatements.WriteString(resourcetaskshape.GongMarshallField(stage, "EndOrientation"))
		initializerStatements.WriteString(resourcetaskshape.GongMarshallField(stage, "CornerOffsetRatio"))
		initializerStatements.WriteString(resourcetaskshape.GongMarshallField(stage, "IsHidden"))
		initializerStatements.WriteString(resourcetaskshape.GongMarshallField(stage, "IsPresentationOverridden"))
		initializerStatements.WriteString(resourcetaskshape.GongMarshallField(stage, "Stroke"))
		initializerStatements.WriteString(resourcetaskshape.GongMarshallField(stage, "StrokeWidth"))
		initializerStatements.WriteString(resourcetaskshape.GongMarshallField(stage, "StrokeOpacity"))
		initializerStatements.WriteString(resourcetaskshape.GongMarshallField(stage, "StrokeDashArray"))
		initializerStatements.WriteString(resourcetaskshape.GongMarshallField(stage, "FontSize"))
		initializerStatements.WriteString(resourcetaskshape.GongMarshallField(stage, "FontColor"))
	}
	initRes = initializerStatements.String()
	ptrRes = pointersInitializesStatements.String()
	return
}
func (task *Task) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(task.GongMarshallField(stage, "Name"))
		initializerStatements.WriteString(task.GongMarshallField(stage, "Description"))
		initializerStatements.WriteString(task.GongMarshallField(stage, "Start"))
		initializerStatements.WriteString(task.GongMarshallField(stage, "End"))
		initializerStatements.WriteString(task.GongMarshallField(stage, "IsAllDay"))
		initializerStatements.WriteString(task.GongMarshallField(stage, "IsMilestone"))
		pointersInitializesStatements.WriteString(task.GongMarshallField(stage, "TaskGroups"))
		pointersInitializesStatements.WriteString(task.GongMarshallField(stage, "Predecessors"))
		initializerStatements.WriteString(task.GongMarshallField(stage, "DependencyType"))
		initializerStatements.WriteString(task.GongMarshallField(stage, "DependencyDurationYears"))
		initializerStatements.WriteString(task.GongMarshallField(stage, "DependencyDurationMonths"))
		initializerStatements.WriteString(task.GongMarshallField(stage, "DependencyDurationWeeks"))
		initializerStatements.WriteString(task.GongMarshallField(stage, "DependencyDurationDays"))
		initializerStatements.WriteString(task.GongMarshallField(stage, "DependencyDurationHours"))
		initializerStatements.WriteString(task.GongMarshallField(stage, "DurationYears"))
		initializerStatements.WriteString(task.GongMarshallField(stage, "DurationMonths"))
		initializerStatements.WriteString(task.GongMarshallField(stage, "DurationWeeks"))
		initializerStatements.WriteString(task.GongMarshallField(stage, "DurationDays"))
		initializerStatements.WriteString(task.GongMarshallField(stage, "DurationHours"))
		initializerStatements.WriteString(task.GongMarshallField(stage, "IsEndDateComputedFromDuration"))
		pointersInitializesStatements.WriteString(task.GongMarshallField(stage, "Inputs"))
		pointersInitializesStatements.WriteString(task.GongMarshallField(stage, "Outputs"))
		pointersInitializesStatements.WriteString(task.GongMarshallField(stage, "SubTasks"))
		initializerStatements.WriteString(task.GongMarshallField(stage, "IsWithCompletion"))
		initializerStatements.WriteString(task.GongMarshallField(stage, "Completion"))
		pointersInitializesStatements.WriteString(task.GongMarshallField(stage, "TaskGroupsToDisplay"))
		initializerStatements.WriteString(task.GongMarshallField(stage, "TextPosition"))
		initializerStatements.WriteString(task.GongMarshallField(stage, "XOffset"))
		initializerStatements.WriteString(task.GongMarshallField(stage, "YOffset"))
		initializerStatements.WriteString(task.GongMarshallField(stage, "IsImport"))
		pointersInitializesStatements.WriteString(task.GongMarshallField(stage, "ReferencedTask"))
		initializerStatements.WriteString(task.GongMarshallField(stage, "IsInputsNodeExpanded"))
		initializerStatements.WriteString(task.GongMarshallField(stage, "IsOutputsNodeExpanded"))
		initializerStatements.WriteString(task.GongMarshallField(stage, "ComputedPrefix"))
		initializerStatements.WriteString(task.GongMarshallField(stage, "IsExpanded"))
		initializerStatements.WriteString(task.GongMarshallField(stage, "LayoutDirection"))
	}
	initRes = initializerStatements.String()
	ptrRes = pointersInitializesStatements.String()
	return
}
func (taskcompositionshape *TaskCompositionShape) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(taskcompositionshape.GongMarshallField(stage, "Name"))
		pointersInitializesStatements.WriteString(taskcompositionshape.GongMarshallField(stage, "Task"))
		initializerStatements.WriteString(taskcompositionshape.GongMarshallField(stage, "StartRatio"))
		initializerStatements.WriteString(taskcompositionshape.GongMarshallField(stage, "EndRatio"))
		initializerStatements.WriteString(taskcompositionshape.GongMarshallField(stage, "StartOrientation"))
		initializerStatements.WriteString(taskcompositionshape.GongMarshallField(stage, "EndOrientation"))
		initializerStatements.WriteString(taskcompositionshape.GongMarshallField(stage, "CornerOffsetRatio"))
		initializerStatements.WriteString(taskcompositionshape.GongMarshallField(stage, "IsHidden"))
		initializerStatements.WriteString(taskcompositionshape.GongMarshallField(stage, "IsPresentationOverridden"))
		initializerStatements.WriteString(taskcompositionshape.GongMarshallField(stage, "Stroke"))
		initializerStatements.WriteString(taskcompositionshape.GongMarshallField(stage, "StrokeWidth"))
		initializerStatements.WriteString(taskcompositionshape.GongMarshallField(stage, "StrokeOpacity"))
		initializerStatements.WriteString(taskcompositionshape.GongMarshallField(stage, "StrokeDashArray"))
		initializerStatements.WriteString(taskcompositionshape.GongMarshallField(stage, "FontSize"))
		initializerStatements.WriteString(taskcompositionshape.GongMarshallField(stage, "FontColor"))
	}
	initRes = initializerStatements.String()
	ptrRes = pointersInitializesStatements.String()
	return
}
func (taskgroup *TaskGroup) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(taskgroup.GongMarshallField(stage, "Name"))
		initializerStatements.WriteString(taskgroup.GongMarshallField(stage, "ComputedPrefix"))
		initializerStatements.WriteString(taskgroup.GongMarshallField(stage, "IsExpanded"))
		pointersInitializesStatements.WriteString(taskgroup.GongMarshallField(stage, "Tasks"))
	}
	initRes = initializerStatements.String()
	ptrRes = pointersInitializesStatements.String()
	return
}
func (taskgroupshape *TaskGroupShape) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(taskgroupshape.GongMarshallField(stage, "Name"))
		pointersInitializesStatements.WriteString(taskgroupshape.GongMarshallField(stage, "TaskGroup"))
		initializerStatements.WriteString(taskgroupshape.GongMarshallField(stage, "X"))
		initializerStatements.WriteString(taskgroupshape.GongMarshallField(stage, "Y"))
		initializerStatements.WriteString(taskgroupshape.GongMarshallField(stage, "Width"))
		initializerStatements.WriteString(taskgroupshape.GongMarshallField(stage, "Height"))
		initializerStatements.WriteString(taskgroupshape.GongMarshallField(stage, "IsHidden"))
		initializerStatements.WriteString(taskgroupshape.GongMarshallField(stage, "IsPresentationOverridden"))
		initializerStatements.WriteString(taskgroupshape.GongMarshallField(stage, "Color"))
		initializerStatements.WriteString(taskgroupshape.GongMarshallField(stage, "FillOpacity"))
		initializerStatements.WriteString(taskgroupshape.GongMarshallField(stage, "Stroke"))
		initializerStatements.WriteString(taskgroupshape.GongMarshallField(stage, "StrokeWidth"))
		initializerStatements.WriteString(taskgroupshape.GongMarshallField(stage, "StrokeOpacity"))
		initializerStatements.WriteString(taskgroupshape.GongMarshallField(stage, "StrokeDashArray"))
		initializerStatements.WriteString(taskgroupshape.GongMarshallField(stage, "RX"))
		initializerStatements.WriteString(taskgroupshape.GongMarshallField(stage, "FontSize"))
		initializerStatements.WriteString(taskgroupshape.GongMarshallField(stage, "FontWeight"))
		initializerStatements.WriteString(taskgroupshape.GongMarshallField(stage, "FontColor"))
	}
	initRes = initializerStatements.String()
	ptrRes = pointersInitializesStatements.String()
	return
}
func (taskinputshape *TaskInputShape) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(taskinputshape.GongMarshallField(stage, "Name"))
		pointersInitializesStatements.WriteString(taskinputshape.GongMarshallField(stage, "Product"))
		pointersInitializesStatements.WriteString(taskinputshape.GongMarshallField(stage, "Task"))
		initializerStatements.WriteString(taskinputshape.GongMarshallField(stage, "StartRatio"))
		initializerStatements.WriteString(taskinputshape.GongMarshallField(stage, "EndRatio"))
		initializerStatements.WriteString(taskinputshape.GongMarshallField(stage, "StartOrientation"))
		initializerStatements.WriteString(taskinputshape.GongMarshallField(stage, "EndOrientation"))
		initializerStatements.WriteString(taskinputshape.GongMarshallField(stage, "CornerOffsetRatio"))
		initializerStatements.WriteString(taskinputshape.GongMarshallField(stage, "IsHidden"))
		initializerStatements.WriteString(taskinputshape.GongMarshallField(stage, "IsPresentationOverridden"))
		initializerStatements.WriteString(taskinputshape.GongMarshallField(stage, "Stroke"))
		initializerStatements.WriteString(taskinputshape.GongMarshallField(stage, "StrokeWidth"))
		initializerStatements.WriteString(taskinputshape.GongMarshallField(stage, "StrokeOpacity"))
		initializerStatements.WriteString(taskinputshape.GongMarshallField(stage, "StrokeDashArray"))
		initializerStatements.WriteString(taskinputshape.GongMarshallField(stage, "FontSize"))
		initializerStatements.WriteString(taskinputshape.GongMarshallField(stage, "FontColor"))
	}
	initRes = initializerStatements.String()
	ptrRes = pointersInitializesStatements.String()
	return
}
func (taskoutputshape *TaskOutputShape) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(taskoutputshape.GongMarshallField(stage, "Name"))
		pointersInitializesStatements.WriteString(taskoutputshape.GongMarshallField(stage, "Task"))
		pointersInitializesStatements.WriteString(taskoutputshape.GongMarshallField(stage, "Product"))
		initializerStatements.WriteString(taskoutputshape.GongMarshallField(stage, "StartRatio"))
		initializerStatements.WriteString(taskoutputshape.GongMarshallField(stage, "EndRatio"))
		initializerStatements.WriteString(taskoutputshape.GongMarshallField(stage, "StartOrientation"))
		initializerStatements.WriteString(taskoutputshape.GongMarshallField(stage, "EndOrientation"))
		initializerStatements.WriteString(taskoutputshape.GongMarshallField(stage, "CornerOffsetRatio"))
		initializerStatements.WriteString(taskoutputshape.GongMarshallField(stage, "IsHidden"))
		initializerStatements.WriteString(taskoutputshape.GongMarshallField(stage, "IsPresentationOverridden"))
		initializerStatements.WriteString(taskoutputshape.GongMarshallField(stage, "Stroke"))
		initializerStatements.WriteString(taskoutputshape.GongMarshallField(stage, "StrokeWidth"))
		initializerStatements.WriteString(taskoutputshape.GongMarshallField(stage, "StrokeOpacity"))
		initializerStatements.WriteString(taskoutputshape.GongMarshallField(stage, "StrokeDashArray"))
		initializerStatements.WriteString(taskoutputshape.GongMarshallField(stage, "FontSize"))
		initializerStatements.WriteString(taskoutputshape.GongMarshallField(stage, "FontColor"))
	}
	initRes = initializerStatements.String()
	ptrRes = pointersInitializesStatements.String()
	return
}
func (taskpredecessorshape *TaskPredecessorShape) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(taskpredecessorshape.GongMarshallField(stage, "Name"))
		pointersInitializesStatements.WriteString(taskpredecessorshape.GongMarshallField(stage, "Predecessor"))
		pointersInitializesStatements.WriteString(taskpredecessorshape.GongMarshallField(stage, "Task"))
		initializerStatements.WriteString(taskpredecessorshape.GongMarshallField(stage, "StartRatio"))
		initializerStatements.WriteString(taskpredecessorshape.GongMarshallField(stage, "EndRatio"))
		initializerStatements.WriteString(taskpredecessorshape.GongMarshallField(stage, "StartOrientation"))
		initializerStatements.WriteString(taskpredecessorshape.GongMarshallField(stage, "EndOrientation"))
		initializerStatements.WriteString(taskpredecessorshape.GongMarshallField(stage, "CornerOffsetRatio"))
		initializerStatements.WriteString(taskpredecessorshape.GongMarshallField(stage, "IsHidden"))
		initializerStatements.WriteString(taskpredecessorshape.GongMarshallField(stage, "IsPresentationOverridden"))
		initializerStatements.WriteString(taskpredecessorshape.GongMarshallField(stage, "Stroke"))
		initializerStatements.WriteString(taskpredecessorshape.GongMarshallField(stage, "StrokeWidth"))
		initializerStatements.WriteString(taskpredecessorshape.GongMarshallField(stage, "StrokeOpacity"))
		initializerStatements.WriteString(taskpredecessorshape.GongMarshallField(stage, "StrokeDashArray"))
		initializerStatements.WriteString(taskpredecessorshape.GongMarshallField(stage, "FontSize"))
		initializerStatements.WriteString(taskpredecessorshape.GongMarshallField(stage, "FontColor"))
	}
	initRes = initializerStatements.String()
	ptrRes = pointersInitializesStatements.String()
	return
}
func (taskshape *TaskShape) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(taskshape.GongMarshallField(stage, "Name"))
		pointersInitializesStatements.WriteString(taskshape.GongMarshallField(stage, "Task"))
		initializerStatements.WriteString(taskshape.GongMarshallField(stage, "IsShowDate"))
		initializerStatements.WriteString(taskshape.GongMarshallField(stage, "VerticalOffset"))
		initializerStatements.WriteString(taskshape.GongMarshallField(stage, "DisplayVerticalBar"))
		initializerStatements.WriteString(taskshape.GongMarshallField(stage, "IsLayoutDirectionDifferent"))
		initializerStatements.WriteString(taskshape.GongMarshallField(stage, "X"))
		initializerStatements.WriteString(taskshape.GongMarshallField(stage, "Y"))
		initializerStatements.WriteString(taskshape.GongMarshallField(stage, "Width"))
		initializerStatements.WriteString(taskshape.GongMarshallField(stage, "Height"))
		initializerStatements.WriteString(taskshape.GongMarshallField(stage, "IsHidden"))
		initializerStatements.WriteString(taskshape.GongMarshallField(stage, "IsPresentationOverridden"))
		initializerStatements.WriteString(taskshape.GongMarshallField(stage, "Color"))
		initializerStatements.WriteString(taskshape.GongMarshallField(stage, "FillOpacity"))
		initializerStatements.WriteString(taskshape.GongMarshallField(stage, "Stroke"))
		initializerStatements.WriteString(taskshape.GongMarshallField(stage, "StrokeWidth"))
		initializerStatements.WriteString(taskshape.GongMarshallField(stage, "StrokeOpacity"))
		initializerStatements.WriteString(taskshape.GongMarshallField(stage, "StrokeDashArray"))
		initializerStatements.WriteString(taskshape.GongMarshallField(stage, "RX"))
		initializerStatements.WriteString(taskshape.GongMarshallField(stage, "FontSize"))
		initializerStatements.WriteString(taskshape.GongMarshallField(stage, "FontWeight"))
		initializerStatements.WriteString(taskshape.GongMarshallField(stage, "FontColor"))
	}
	initRes = initializerStatements.String()
	ptrRes = pointersInitializesStatements.String()
	return
}
