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
	modelsPackageName := "github.com/fullstack-lang/gong/lib/svg/go/models"
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
		for _, animate := range __gong__sortStageSetInstances(stage.Animates, stage.Animate_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			animateIdent := animate.GongGetIdentifier(stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&models.Animate{Name: %s}).Stage(stage)", animateIdent, __gong__toRawStringLiteral(animate.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", animateIdent, __gong__toRawStringLiteral(animate.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.AttributeName = %s", animateIdent, __gong__toRawStringLiteral(animate.AttributeName)))
			values.WriteString(fmt.Sprintf("\n\t%s.Values = %s", animateIdent, __gong__toRawStringLiteral(animate.Values)))
			values.WriteString(fmt.Sprintf("\n\t%s.From = %s", animateIdent, __gong__toRawStringLiteral(animate.From)))
			values.WriteString(fmt.Sprintf("\n\t%s.To = %s", animateIdent, __gong__toRawStringLiteral(animate.To)))
			values.WriteString(fmt.Sprintf("\n\t%s.Dur = %s", animateIdent, __gong__toRawStringLiteral(animate.Dur)))
			values.WriteString(fmt.Sprintf("\n\t%s.RepeatCount = %s", animateIdent, __gong__toRawStringLiteral(animate.RepeatCount)))
		}
	}
	if stage != nil {
		for _, circle := range __gong__sortStageSetInstances(stage.Circles, stage.Circle_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			circleIdent := circle.GongGetIdentifier(stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&models.Circle{Name: %s}).Stage(stage)", circleIdent, __gong__toRawStringLiteral(circle.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", circleIdent, __gong__toRawStringLiteral(circle.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.CX = %f", circleIdent, circle.CX))
			values.WriteString(fmt.Sprintf("\n\t%s.CY = %f", circleIdent, circle.CY))
			values.WriteString(fmt.Sprintf("\n\t%s.Radius = %f", circleIdent, circle.Radius))
			values.WriteString(fmt.Sprintf("\n\t%s.Color = %s", circleIdent, __gong__toRawStringLiteral(circle.Color)))
			values.WriteString(fmt.Sprintf("\n\t%s.FillOpacity = %f", circleIdent, circle.FillOpacity))
			values.WriteString(fmt.Sprintf("\n\t%s.Stroke = %s", circleIdent, __gong__toRawStringLiteral(circle.Stroke)))
			values.WriteString(fmt.Sprintf("\n\t%s.StrokeOpacity = %f", circleIdent, circle.StrokeOpacity))
			values.WriteString(fmt.Sprintf("\n\t%s.StrokeWidth = %f", circleIdent, circle.StrokeWidth))
			values.WriteString(fmt.Sprintf("\n\t%s.StrokeDashArray = %s", circleIdent, __gong__toRawStringLiteral(circle.StrokeDashArray)))
			values.WriteString(fmt.Sprintf("\n\t%s.StrokeDashArrayWhenSelected = %s", circleIdent, __gong__toRawStringLiteral(circle.StrokeDashArrayWhenSelected)))
			values.WriteString(fmt.Sprintf("\n\t%s.Transform = %s", circleIdent, __gong__toRawStringLiteral(circle.Transform)))
			for _, elem := range circle.Animations {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := elem.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Animations = append(%s.Animations, %s)", circleIdent, circleIdent, targetIdent))
			}
		}
	}
	if stage != nil {
		for _, condition := range __gong__sortStageSetInstances(stage.Conditions, stage.Condition_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			conditionIdent := condition.GongGetIdentifier(stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&models.Condition{Name: %s}).Stage(stage)", conditionIdent, __gong__toRawStringLiteral(condition.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", conditionIdent, __gong__toRawStringLiteral(condition.Name)))
		}
	}
	if stage != nil {
		for _, controlpoint := range __gong__sortStageSetInstances(stage.ControlPoints, stage.ControlPoint_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			controlpointIdent := controlpoint.GongGetIdentifier(stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&models.ControlPoint{Name: %s}).Stage(stage)", controlpointIdent, __gong__toRawStringLiteral(controlpoint.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", controlpointIdent, __gong__toRawStringLiteral(controlpoint.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.X_Relative = %f", controlpointIdent, controlpoint.X_Relative))
			values.WriteString(fmt.Sprintf("\n\t%s.Y_Relative = %f", controlpointIdent, controlpoint.Y_Relative))
			if controlpoint.ClosestRect != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := controlpoint.ClosestRect.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.ClosestRect = %s", controlpointIdent, targetIdent))
			}
		}
	}
	if stage != nil {
		for _, ellipse := range __gong__sortStageSetInstances(stage.Ellipses, stage.Ellipse_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			ellipseIdent := ellipse.GongGetIdentifier(stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&models.Ellipse{Name: %s}).Stage(stage)", ellipseIdent, __gong__toRawStringLiteral(ellipse.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", ellipseIdent, __gong__toRawStringLiteral(ellipse.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.CX = %f", ellipseIdent, ellipse.CX))
			values.WriteString(fmt.Sprintf("\n\t%s.CY = %f", ellipseIdent, ellipse.CY))
			values.WriteString(fmt.Sprintf("\n\t%s.RX = %f", ellipseIdent, ellipse.RX))
			values.WriteString(fmt.Sprintf("\n\t%s.RY = %f", ellipseIdent, ellipse.RY))
			values.WriteString(fmt.Sprintf("\n\t%s.Color = %s", ellipseIdent, __gong__toRawStringLiteral(ellipse.Color)))
			values.WriteString(fmt.Sprintf("\n\t%s.FillOpacity = %f", ellipseIdent, ellipse.FillOpacity))
			values.WriteString(fmt.Sprintf("\n\t%s.Stroke = %s", ellipseIdent, __gong__toRawStringLiteral(ellipse.Stroke)))
			values.WriteString(fmt.Sprintf("\n\t%s.StrokeOpacity = %f", ellipseIdent, ellipse.StrokeOpacity))
			values.WriteString(fmt.Sprintf("\n\t%s.StrokeWidth = %f", ellipseIdent, ellipse.StrokeWidth))
			values.WriteString(fmt.Sprintf("\n\t%s.StrokeDashArray = %s", ellipseIdent, __gong__toRawStringLiteral(ellipse.StrokeDashArray)))
			values.WriteString(fmt.Sprintf("\n\t%s.StrokeDashArrayWhenSelected = %s", ellipseIdent, __gong__toRawStringLiteral(ellipse.StrokeDashArrayWhenSelected)))
			values.WriteString(fmt.Sprintf("\n\t%s.Transform = %s", ellipseIdent, __gong__toRawStringLiteral(ellipse.Transform)))
			for _, elem := range ellipse.Animates {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := elem.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Animates = append(%s.Animates, %s)", ellipseIdent, ellipseIdent, targetIdent))
			}
		}
	}
	if stage != nil {
		for _, filetodownload := range __gong__sortStageSetInstances(stage.FileToDownloads, stage.FileToDownload_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			filetodownloadIdent := filetodownload.GongGetIdentifier(stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&models.FileToDownload{Name: %s}).Stage(stage)", filetodownloadIdent, __gong__toRawStringLiteral(filetodownload.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", filetodownloadIdent, __gong__toRawStringLiteral(filetodownload.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Base64EncodedContent = %s", filetodownloadIdent, __gong__toRawStringLiteral(filetodownload.Base64EncodedContent)))
		}
	}
	if stage != nil {
		for _, layer := range __gong__sortStageSetInstances(stage.Layers, stage.Layer_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			layerIdent := layer.GongGetIdentifier(stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&models.Layer{Name: %s}).Stage(stage)", layerIdent, __gong__toRawStringLiteral(layer.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", layerIdent, __gong__toRawStringLiteral(layer.Name)))
			for _, elem := range layer.Rects {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := elem.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Rects = append(%s.Rects, %s)", layerIdent, layerIdent, targetIdent))
			}
			for _, elem := range layer.Texts {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := elem.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Texts = append(%s.Texts, %s)", layerIdent, layerIdent, targetIdent))
			}
			for _, elem := range layer.Circles {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := elem.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Circles = append(%s.Circles, %s)", layerIdent, layerIdent, targetIdent))
			}
			for _, elem := range layer.Lines {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := elem.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Lines = append(%s.Lines, %s)", layerIdent, layerIdent, targetIdent))
			}
			for _, elem := range layer.Ellipses {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := elem.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Ellipses = append(%s.Ellipses, %s)", layerIdent, layerIdent, targetIdent))
			}
			for _, elem := range layer.Polylines {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := elem.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Polylines = append(%s.Polylines, %s)", layerIdent, layerIdent, targetIdent))
			}
			for _, elem := range layer.Polygones {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := elem.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Polygones = append(%s.Polygones, %s)", layerIdent, layerIdent, targetIdent))
			}
			for _, elem := range layer.Paths {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := elem.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Paths = append(%s.Paths, %s)", layerIdent, layerIdent, targetIdent))
			}
			for _, elem := range layer.Links {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := elem.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Links = append(%s.Links, %s)", layerIdent, layerIdent, targetIdent))
			}
			for _, elem := range layer.RectLinkLinks {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := elem.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.RectLinkLinks = append(%s.RectLinkLinks, %s)", layerIdent, layerIdent, targetIdent))
			}
		}
	}
	if stage != nil {
		for _, line := range __gong__sortStageSetInstances(stage.Lines, stage.Line_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			lineIdent := line.GongGetIdentifier(stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&models.Line{Name: %s}).Stage(stage)", lineIdent, __gong__toRawStringLiteral(line.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", lineIdent, __gong__toRawStringLiteral(line.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.X1 = %f", lineIdent, line.X1))
			values.WriteString(fmt.Sprintf("\n\t%s.Y1 = %f", lineIdent, line.Y1))
			values.WriteString(fmt.Sprintf("\n\t%s.X2 = %f", lineIdent, line.X2))
			values.WriteString(fmt.Sprintf("\n\t%s.Y2 = %f", lineIdent, line.Y2))
			values.WriteString(fmt.Sprintf("\n\t%s.Color = %s", lineIdent, __gong__toRawStringLiteral(line.Color)))
			values.WriteString(fmt.Sprintf("\n\t%s.FillOpacity = %f", lineIdent, line.FillOpacity))
			values.WriteString(fmt.Sprintf("\n\t%s.Stroke = %s", lineIdent, __gong__toRawStringLiteral(line.Stroke)))
			values.WriteString(fmt.Sprintf("\n\t%s.StrokeOpacity = %f", lineIdent, line.StrokeOpacity))
			values.WriteString(fmt.Sprintf("\n\t%s.StrokeWidth = %f", lineIdent, line.StrokeWidth))
			values.WriteString(fmt.Sprintf("\n\t%s.StrokeDashArray = %s", lineIdent, __gong__toRawStringLiteral(line.StrokeDashArray)))
			values.WriteString(fmt.Sprintf("\n\t%s.StrokeDashArrayWhenSelected = %s", lineIdent, __gong__toRawStringLiteral(line.StrokeDashArrayWhenSelected)))
			values.WriteString(fmt.Sprintf("\n\t%s.Transform = %s", lineIdent, __gong__toRawStringLiteral(line.Transform)))
			values.WriteString(fmt.Sprintf("\n\t%s.MouseClickX = %f", lineIdent, line.MouseClickX))
			values.WriteString(fmt.Sprintf("\n\t%s.MouseClickY = %f", lineIdent, line.MouseClickY))
			for _, elem := range line.Animates {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := elem.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Animates = append(%s.Animates, %s)", lineIdent, lineIdent, targetIdent))
			}
		}
	}
	if stage != nil {
		for _, link := range __gong__sortStageSetInstances(stage.Links, stage.Link_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			linkIdent := link.GongGetIdentifier(stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&models.Link{Name: %s}).Stage(stage)", linkIdent, __gong__toRawStringLiteral(link.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", linkIdent, __gong__toRawStringLiteral(link.Name)))
			if link.Type.ToCodeString() != "" {
				values.WriteString(fmt.Sprintf("\n\t%s.Type = models.%s", linkIdent, link.Type.ToCodeString()))
			} else {
				values.WriteString(fmt.Sprintf("\n\t%s.Type = \"\"", linkIdent))
			}
			values.WriteString(fmt.Sprintf("\n\t%s.IsBezierCurve = %t", linkIdent, link.IsBezierCurve))
			if link.StartAnchorType.ToCodeString() != "" {
				values.WriteString(fmt.Sprintf("\n\t%s.StartAnchorType = models.%s", linkIdent, link.StartAnchorType.ToCodeString()))
			} else {
				values.WriteString(fmt.Sprintf("\n\t%s.StartAnchorType = \"\"", linkIdent))
			}
			if link.EndAnchorType.ToCodeString() != "" {
				values.WriteString(fmt.Sprintf("\n\t%s.EndAnchorType = models.%s", linkIdent, link.EndAnchorType.ToCodeString()))
			} else {
				values.WriteString(fmt.Sprintf("\n\t%s.EndAnchorType = \"\"", linkIdent))
			}
			if link.StartOrientation.ToCodeString() != "" {
				values.WriteString(fmt.Sprintf("\n\t%s.StartOrientation = models.%s", linkIdent, link.StartOrientation.ToCodeString()))
			} else {
				values.WriteString(fmt.Sprintf("\n\t%s.StartOrientation = \"\"", linkIdent))
			}
			values.WriteString(fmt.Sprintf("\n\t%s.StartRatio = %f", linkIdent, link.StartRatio))
			if link.EndOrientation.ToCodeString() != "" {
				values.WriteString(fmt.Sprintf("\n\t%s.EndOrientation = models.%s", linkIdent, link.EndOrientation.ToCodeString()))
			} else {
				values.WriteString(fmt.Sprintf("\n\t%s.EndOrientation = \"\"", linkIdent))
			}
			values.WriteString(fmt.Sprintf("\n\t%s.EndRatio = %f", linkIdent, link.EndRatio))
			values.WriteString(fmt.Sprintf("\n\t%s.CornerOffsetRatio = %f", linkIdent, link.CornerOffsetRatio))
			values.WriteString(fmt.Sprintf("\n\t%s.CornerRadius = %f", linkIdent, link.CornerRadius))
			values.WriteString(fmt.Sprintf("\n\t%s.HasEndArrow = %t", linkIdent, link.HasEndArrow))
			values.WriteString(fmt.Sprintf("\n\t%s.EndArrowSize = %f", linkIdent, link.EndArrowSize))
			values.WriteString(fmt.Sprintf("\n\t%s.EndArrowOffset = %f", linkIdent, link.EndArrowOffset))
			values.WriteString(fmt.Sprintf("\n\t%s.HasStartArrow = %t", linkIdent, link.HasStartArrow))
			values.WriteString(fmt.Sprintf("\n\t%s.StartArrowSize = %f", linkIdent, link.StartArrowSize))
			values.WriteString(fmt.Sprintf("\n\t%s.StartArrowOffset = %f", linkIdent, link.StartArrowOffset))
			values.WriteString(fmt.Sprintf("\n\t%s.Color = %s", linkIdent, __gong__toRawStringLiteral(link.Color)))
			values.WriteString(fmt.Sprintf("\n\t%s.FillOpacity = %f", linkIdent, link.FillOpacity))
			values.WriteString(fmt.Sprintf("\n\t%s.Stroke = %s", linkIdent, __gong__toRawStringLiteral(link.Stroke)))
			values.WriteString(fmt.Sprintf("\n\t%s.StrokeOpacity = %f", linkIdent, link.StrokeOpacity))
			values.WriteString(fmt.Sprintf("\n\t%s.StrokeWidth = %f", linkIdent, link.StrokeWidth))
			values.WriteString(fmt.Sprintf("\n\t%s.StrokeDashArray = %s", linkIdent, __gong__toRawStringLiteral(link.StrokeDashArray)))
			values.WriteString(fmt.Sprintf("\n\t%s.StrokeDashArrayWhenSelected = %s", linkIdent, __gong__toRawStringLiteral(link.StrokeDashArrayWhenSelected)))
			values.WriteString(fmt.Sprintf("\n\t%s.Transform = %s", linkIdent, __gong__toRawStringLiteral(link.Transform)))
			values.WriteString(fmt.Sprintf("\n\t%s.MouseX = %f", linkIdent, link.MouseX))
			values.WriteString(fmt.Sprintf("\n\t%s.MouseY = %f", linkIdent, link.MouseY))
			if link.MouseEventKey.ToCodeString() != "" {
				values.WriteString(fmt.Sprintf("\n\t%s.MouseEventKey = models.%s", linkIdent, link.MouseEventKey.ToCodeString()))
			} else {
				values.WriteString(fmt.Sprintf("\n\t%s.MouseEventKey = \"\"", linkIdent))
			}
			if link.Start != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := link.Start.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Start = %s", linkIdent, targetIdent))
			}
			if link.End != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := link.End.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.End = %s", linkIdent, targetIdent))
			}
			for _, elem := range link.TextAtArrowStart {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := elem.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.TextAtArrowStart = append(%s.TextAtArrowStart, %s)", linkIdent, linkIdent, targetIdent))
			}
			for _, elem := range link.TextAtArrowEnd {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := elem.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.TextAtArrowEnd = append(%s.TextAtArrowEnd, %s)", linkIdent, linkIdent, targetIdent))
			}
			for _, elem := range link.TextAtCorner {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := elem.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.TextAtCorner = append(%s.TextAtCorner, %s)", linkIdent, linkIdent, targetIdent))
			}
			for _, elem := range link.PathAtArrowStart {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := elem.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.PathAtArrowStart = append(%s.PathAtArrowStart, %s)", linkIdent, linkIdent, targetIdent))
			}
			for _, elem := range link.PathAtArrowEnd {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := elem.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.PathAtArrowEnd = append(%s.PathAtArrowEnd, %s)", linkIdent, linkIdent, targetIdent))
			}
			for _, elem := range link.PathAtCorner {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := elem.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.PathAtCorner = append(%s.PathAtCorner, %s)", linkIdent, linkIdent, targetIdent))
			}
			for _, elem := range link.ControlPoints {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := elem.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.ControlPoints = append(%s.ControlPoints, %s)", linkIdent, linkIdent, targetIdent))
			}
		}
	}
	if stage != nil {
		for _, linkanchoredpath := range __gong__sortStageSetInstances(stage.LinkAnchoredPaths, stage.LinkAnchoredPath_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			linkanchoredpathIdent := linkanchoredpath.GongGetIdentifier(stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&models.LinkAnchoredPath{Name: %s}).Stage(stage)", linkanchoredpathIdent, __gong__toRawStringLiteral(linkanchoredpath.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", linkanchoredpathIdent, __gong__toRawStringLiteral(linkanchoredpath.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Definition = %s", linkanchoredpathIdent, __gong__toRawStringLiteral(linkanchoredpath.Definition)))
			values.WriteString(fmt.Sprintf("\n\t%s.X_Offset = %f", linkanchoredpathIdent, linkanchoredpath.X_Offset))
			values.WriteString(fmt.Sprintf("\n\t%s.Y_Offset = %f", linkanchoredpathIdent, linkanchoredpath.Y_Offset))
			values.WriteString(fmt.Sprintf("\n\t%s.ScalePropotionnally = %t", linkanchoredpathIdent, linkanchoredpath.ScalePropotionnally))
			values.WriteString(fmt.Sprintf("\n\t%s.AppliedScaling = %f", linkanchoredpathIdent, linkanchoredpath.AppliedScaling))
			values.WriteString(fmt.Sprintf("\n\t%s.Color = %s", linkanchoredpathIdent, __gong__toRawStringLiteral(linkanchoredpath.Color)))
			values.WriteString(fmt.Sprintf("\n\t%s.FillOpacity = %f", linkanchoredpathIdent, linkanchoredpath.FillOpacity))
			values.WriteString(fmt.Sprintf("\n\t%s.Stroke = %s", linkanchoredpathIdent, __gong__toRawStringLiteral(linkanchoredpath.Stroke)))
			values.WriteString(fmt.Sprintf("\n\t%s.StrokeOpacity = %f", linkanchoredpathIdent, linkanchoredpath.StrokeOpacity))
			values.WriteString(fmt.Sprintf("\n\t%s.StrokeWidth = %f", linkanchoredpathIdent, linkanchoredpath.StrokeWidth))
			values.WriteString(fmt.Sprintf("\n\t%s.StrokeDashArray = %s", linkanchoredpathIdent, __gong__toRawStringLiteral(linkanchoredpath.StrokeDashArray)))
			values.WriteString(fmt.Sprintf("\n\t%s.StrokeDashArrayWhenSelected = %s", linkanchoredpathIdent, __gong__toRawStringLiteral(linkanchoredpath.StrokeDashArrayWhenSelected)))
			values.WriteString(fmt.Sprintf("\n\t%s.Transform = %s", linkanchoredpathIdent, __gong__toRawStringLiteral(linkanchoredpath.Transform)))
		}
	}
	if stage != nil {
		for _, linkanchoredtext := range __gong__sortStageSetInstances(stage.LinkAnchoredTexts, stage.LinkAnchoredText_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			linkanchoredtextIdent := linkanchoredtext.GongGetIdentifier(stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&models.LinkAnchoredText{Name: %s}).Stage(stage)", linkanchoredtextIdent, __gong__toRawStringLiteral(linkanchoredtext.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", linkanchoredtextIdent, __gong__toRawStringLiteral(linkanchoredtext.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Content = %s", linkanchoredtextIdent, __gong__toRawStringLiteral(linkanchoredtext.Content)))
			values.WriteString(fmt.Sprintf("\n\t%s.AutomaticLayout = %t", linkanchoredtextIdent, linkanchoredtext.AutomaticLayout))
			if linkanchoredtext.LinkAnchorType.ToCodeString() != "" {
				values.WriteString(fmt.Sprintf("\n\t%s.LinkAnchorType = models.%s", linkanchoredtextIdent, linkanchoredtext.LinkAnchorType.ToCodeString()))
			} else {
				values.WriteString(fmt.Sprintf("\n\t%s.LinkAnchorType = \"\"", linkanchoredtextIdent))
			}
			values.WriteString(fmt.Sprintf("\n\t%s.X_Offset = %f", linkanchoredtextIdent, linkanchoredtext.X_Offset))
			values.WriteString(fmt.Sprintf("\n\t%s.Y_Offset = %f", linkanchoredtextIdent, linkanchoredtext.Y_Offset))
			values.WriteString(fmt.Sprintf("\n\t%s.FontWeight = %s", linkanchoredtextIdent, __gong__toRawStringLiteral(linkanchoredtext.FontWeight)))
			values.WriteString(fmt.Sprintf("\n\t%s.FontSize = %s", linkanchoredtextIdent, __gong__toRawStringLiteral(linkanchoredtext.FontSize)))
			values.WriteString(fmt.Sprintf("\n\t%s.FontStyle = %s", linkanchoredtextIdent, __gong__toRawStringLiteral(linkanchoredtext.FontStyle)))
			values.WriteString(fmt.Sprintf("\n\t%s.LetterSpacing = %s", linkanchoredtextIdent, __gong__toRawStringLiteral(linkanchoredtext.LetterSpacing)))
			values.WriteString(fmt.Sprintf("\n\t%s.FontFamily = %s", linkanchoredtextIdent, __gong__toRawStringLiteral(linkanchoredtext.FontFamily)))
			if linkanchoredtext.WhiteSpace.ToCodeString() != "" {
				values.WriteString(fmt.Sprintf("\n\t%s.WhiteSpace = models.%s", linkanchoredtextIdent, linkanchoredtext.WhiteSpace.ToCodeString()))
			} else {
				values.WriteString(fmt.Sprintf("\n\t%s.WhiteSpace = \"\"", linkanchoredtextIdent))
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Color = %s", linkanchoredtextIdent, __gong__toRawStringLiteral(linkanchoredtext.Color)))
			values.WriteString(fmt.Sprintf("\n\t%s.FillOpacity = %f", linkanchoredtextIdent, linkanchoredtext.FillOpacity))
			values.WriteString(fmt.Sprintf("\n\t%s.Stroke = %s", linkanchoredtextIdent, __gong__toRawStringLiteral(linkanchoredtext.Stroke)))
			values.WriteString(fmt.Sprintf("\n\t%s.StrokeOpacity = %f", linkanchoredtextIdent, linkanchoredtext.StrokeOpacity))
			values.WriteString(fmt.Sprintf("\n\t%s.StrokeWidth = %f", linkanchoredtextIdent, linkanchoredtext.StrokeWidth))
			values.WriteString(fmt.Sprintf("\n\t%s.StrokeDashArray = %s", linkanchoredtextIdent, __gong__toRawStringLiteral(linkanchoredtext.StrokeDashArray)))
			values.WriteString(fmt.Sprintf("\n\t%s.StrokeDashArrayWhenSelected = %s", linkanchoredtextIdent, __gong__toRawStringLiteral(linkanchoredtext.StrokeDashArrayWhenSelected)))
			values.WriteString(fmt.Sprintf("\n\t%s.Transform = %s", linkanchoredtextIdent, __gong__toRawStringLiteral(linkanchoredtext.Transform)))
			for _, elem := range linkanchoredtext.Animates {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := elem.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Animates = append(%s.Animates, %s)", linkanchoredtextIdent, linkanchoredtextIdent, targetIdent))
			}
		}
	}
	if stage != nil {
		for _, path := range __gong__sortStageSetInstances(stage.Paths, stage.Path_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			pathIdent := path.GongGetIdentifier(stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&models.Path{Name: %s}).Stage(stage)", pathIdent, __gong__toRawStringLiteral(path.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", pathIdent, __gong__toRawStringLiteral(path.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Definition = %s", pathIdent, __gong__toRawStringLiteral(path.Definition)))
			values.WriteString(fmt.Sprintf("\n\t%s.Color = %s", pathIdent, __gong__toRawStringLiteral(path.Color)))
			values.WriteString(fmt.Sprintf("\n\t%s.FillOpacity = %f", pathIdent, path.FillOpacity))
			values.WriteString(fmt.Sprintf("\n\t%s.Stroke = %s", pathIdent, __gong__toRawStringLiteral(path.Stroke)))
			values.WriteString(fmt.Sprintf("\n\t%s.StrokeOpacity = %f", pathIdent, path.StrokeOpacity))
			values.WriteString(fmt.Sprintf("\n\t%s.StrokeWidth = %f", pathIdent, path.StrokeWidth))
			values.WriteString(fmt.Sprintf("\n\t%s.StrokeDashArray = %s", pathIdent, __gong__toRawStringLiteral(path.StrokeDashArray)))
			values.WriteString(fmt.Sprintf("\n\t%s.StrokeDashArrayWhenSelected = %s", pathIdent, __gong__toRawStringLiteral(path.StrokeDashArrayWhenSelected)))
			values.WriteString(fmt.Sprintf("\n\t%s.Transform = %s", pathIdent, __gong__toRawStringLiteral(path.Transform)))
			for _, elem := range path.Animates {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := elem.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Animates = append(%s.Animates, %s)", pathIdent, pathIdent, targetIdent))
			}
		}
	}
	if stage != nil {
		for _, point := range __gong__sortStageSetInstances(stage.Points, stage.Point_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			pointIdent := point.GongGetIdentifier(stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&models.Point{Name: %s}).Stage(stage)", pointIdent, __gong__toRawStringLiteral(point.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", pointIdent, __gong__toRawStringLiteral(point.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.X = %f", pointIdent, point.X))
			values.WriteString(fmt.Sprintf("\n\t%s.Y = %f", pointIdent, point.Y))
		}
	}
	if stage != nil {
		for _, polygone := range __gong__sortStageSetInstances(stage.Polygones, stage.Polygone_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			polygoneIdent := polygone.GongGetIdentifier(stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&models.Polygone{Name: %s}).Stage(stage)", polygoneIdent, __gong__toRawStringLiteral(polygone.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", polygoneIdent, __gong__toRawStringLiteral(polygone.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Points = %s", polygoneIdent, __gong__toRawStringLiteral(polygone.Points)))
			values.WriteString(fmt.Sprintf("\n\t%s.Color = %s", polygoneIdent, __gong__toRawStringLiteral(polygone.Color)))
			values.WriteString(fmt.Sprintf("\n\t%s.FillOpacity = %f", polygoneIdent, polygone.FillOpacity))
			values.WriteString(fmt.Sprintf("\n\t%s.Stroke = %s", polygoneIdent, __gong__toRawStringLiteral(polygone.Stroke)))
			values.WriteString(fmt.Sprintf("\n\t%s.StrokeOpacity = %f", polygoneIdent, polygone.StrokeOpacity))
			values.WriteString(fmt.Sprintf("\n\t%s.StrokeWidth = %f", polygoneIdent, polygone.StrokeWidth))
			values.WriteString(fmt.Sprintf("\n\t%s.StrokeDashArray = %s", polygoneIdent, __gong__toRawStringLiteral(polygone.StrokeDashArray)))
			values.WriteString(fmt.Sprintf("\n\t%s.StrokeDashArrayWhenSelected = %s", polygoneIdent, __gong__toRawStringLiteral(polygone.StrokeDashArrayWhenSelected)))
			values.WriteString(fmt.Sprintf("\n\t%s.Transform = %s", polygoneIdent, __gong__toRawStringLiteral(polygone.Transform)))
			for _, elem := range polygone.Animates {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := elem.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Animates = append(%s.Animates, %s)", polygoneIdent, polygoneIdent, targetIdent))
			}
		}
	}
	if stage != nil {
		for _, polyline := range __gong__sortStageSetInstances(stage.Polylines, stage.Polyline_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			polylineIdent := polyline.GongGetIdentifier(stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&models.Polyline{Name: %s}).Stage(stage)", polylineIdent, __gong__toRawStringLiteral(polyline.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", polylineIdent, __gong__toRawStringLiteral(polyline.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Points = %s", polylineIdent, __gong__toRawStringLiteral(polyline.Points)))
			values.WriteString(fmt.Sprintf("\n\t%s.Color = %s", polylineIdent, __gong__toRawStringLiteral(polyline.Color)))
			values.WriteString(fmt.Sprintf("\n\t%s.FillOpacity = %f", polylineIdent, polyline.FillOpacity))
			values.WriteString(fmt.Sprintf("\n\t%s.Stroke = %s", polylineIdent, __gong__toRawStringLiteral(polyline.Stroke)))
			values.WriteString(fmt.Sprintf("\n\t%s.StrokeOpacity = %f", polylineIdent, polyline.StrokeOpacity))
			values.WriteString(fmt.Sprintf("\n\t%s.StrokeWidth = %f", polylineIdent, polyline.StrokeWidth))
			values.WriteString(fmt.Sprintf("\n\t%s.StrokeDashArray = %s", polylineIdent, __gong__toRawStringLiteral(polyline.StrokeDashArray)))
			values.WriteString(fmt.Sprintf("\n\t%s.StrokeDashArrayWhenSelected = %s", polylineIdent, __gong__toRawStringLiteral(polyline.StrokeDashArrayWhenSelected)))
			values.WriteString(fmt.Sprintf("\n\t%s.Transform = %s", polylineIdent, __gong__toRawStringLiteral(polyline.Transform)))
			for _, elem := range polyline.Animates {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := elem.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Animates = append(%s.Animates, %s)", polylineIdent, polylineIdent, targetIdent))
			}
		}
	}
	if stage != nil {
		for _, rect := range __gong__sortStageSetInstances(stage.Rects, stage.Rect_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			rectIdent := rect.GongGetIdentifier(stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&models.Rect{Name: %s}).Stage(stage)", rectIdent, __gong__toRawStringLiteral(rect.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", rectIdent, __gong__toRawStringLiteral(rect.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.X = %f", rectIdent, rect.X))
			values.WriteString(fmt.Sprintf("\n\t%s.Y = %f", rectIdent, rect.Y))
			values.WriteString(fmt.Sprintf("\n\t%s.Width = %f", rectIdent, rect.Width))
			values.WriteString(fmt.Sprintf("\n\t%s.Height = %f", rectIdent, rect.Height))
			values.WriteString(fmt.Sprintf("\n\t%s.RX = %f", rectIdent, rect.RX))
			values.WriteString(fmt.Sprintf("\n\t%s.Color = %s", rectIdent, __gong__toRawStringLiteral(rect.Color)))
			values.WriteString(fmt.Sprintf("\n\t%s.FillOpacity = %f", rectIdent, rect.FillOpacity))
			values.WriteString(fmt.Sprintf("\n\t%s.Stroke = %s", rectIdent, __gong__toRawStringLiteral(rect.Stroke)))
			values.WriteString(fmt.Sprintf("\n\t%s.StrokeOpacity = %f", rectIdent, rect.StrokeOpacity))
			values.WriteString(fmt.Sprintf("\n\t%s.StrokeWidth = %f", rectIdent, rect.StrokeWidth))
			values.WriteString(fmt.Sprintf("\n\t%s.StrokeDashArray = %s", rectIdent, __gong__toRawStringLiteral(rect.StrokeDashArray)))
			values.WriteString(fmt.Sprintf("\n\t%s.StrokeDashArrayWhenSelected = %s", rectIdent, __gong__toRawStringLiteral(rect.StrokeDashArrayWhenSelected)))
			values.WriteString(fmt.Sprintf("\n\t%s.Transform = %s", rectIdent, __gong__toRawStringLiteral(rect.Transform)))
			values.WriteString(fmt.Sprintf("\n\t%s.IsSelectable = %t", rectIdent, rect.IsSelectable))
			values.WriteString(fmt.Sprintf("\n\t%s.IsSelected = %t", rectIdent, rect.IsSelected))
			values.WriteString(fmt.Sprintf("\n\t%s.CanHaveLeftHandle = %t", rectIdent, rect.CanHaveLeftHandle))
			values.WriteString(fmt.Sprintf("\n\t%s.HasLeftHandle = %t", rectIdent, rect.HasLeftHandle))
			values.WriteString(fmt.Sprintf("\n\t%s.CanHaveRightHandle = %t", rectIdent, rect.CanHaveRightHandle))
			values.WriteString(fmt.Sprintf("\n\t%s.HasRightHandle = %t", rectIdent, rect.HasRightHandle))
			values.WriteString(fmt.Sprintf("\n\t%s.CanHaveTopHandle = %t", rectIdent, rect.CanHaveTopHandle))
			values.WriteString(fmt.Sprintf("\n\t%s.HasTopHandle = %t", rectIdent, rect.HasTopHandle))
			values.WriteString(fmt.Sprintf("\n\t%s.IsScalingProportionally = %t", rectIdent, rect.IsScalingProportionally))
			values.WriteString(fmt.Sprintf("\n\t%s.CanHaveBottomHandle = %t", rectIdent, rect.CanHaveBottomHandle))
			values.WriteString(fmt.Sprintf("\n\t%s.HasBottomHandle = %t", rectIdent, rect.HasBottomHandle))
			values.WriteString(fmt.Sprintf("\n\t%s.CanMoveHorizontaly = %t", rectIdent, rect.CanMoveHorizontaly))
			values.WriteString(fmt.Sprintf("\n\t%s.CanMoveVerticaly = %t", rectIdent, rect.CanMoveVerticaly))
			values.WriteString(fmt.Sprintf("\n\t%s.ChangeColorWhenHovered = %t", rectIdent, rect.ChangeColorWhenHovered))
			values.WriteString(fmt.Sprintf("\n\t%s.ColorWhenHovered = %s", rectIdent, __gong__toRawStringLiteral(rect.ColorWhenHovered)))
			values.WriteString(fmt.Sprintf("\n\t%s.OriginalColor = %s", rectIdent, __gong__toRawStringLiteral(rect.OriginalColor)))
			values.WriteString(fmt.Sprintf("\n\t%s.FillOpacityWhenHovered = %f", rectIdent, rect.FillOpacityWhenHovered))
			values.WriteString(fmt.Sprintf("\n\t%s.OriginalFillOpacity = %f", rectIdent, rect.OriginalFillOpacity))
			values.WriteString(fmt.Sprintf("\n\t%s.HasToolTip = %t", rectIdent, rect.HasToolTip))
			values.WriteString(fmt.Sprintf("\n\t%s.ToolTipText = %s", rectIdent, __gong__toRawStringLiteral(rect.ToolTipText)))
			if rect.ToolTipPosition.ToCodeString() != "" {
				values.WriteString(fmt.Sprintf("\n\t%s.ToolTipPosition = models.%s", rectIdent, rect.ToolTipPosition.ToCodeString()))
			} else {
				values.WriteString(fmt.Sprintf("\n\t%s.ToolTipPosition = \"\"", rectIdent))
			}
			values.WriteString(fmt.Sprintf("\n\t%s.MouseX = %f", rectIdent, rect.MouseX))
			values.WriteString(fmt.Sprintf("\n\t%s.MouseY = %f", rectIdent, rect.MouseY))
			if rect.MouseEventKey.ToCodeString() != "" {
				values.WriteString(fmt.Sprintf("\n\t%s.MouseEventKey = models.%s", rectIdent, rect.MouseEventKey.ToCodeString()))
			} else {
				values.WriteString(fmt.Sprintf("\n\t%s.MouseEventKey = \"\"", rectIdent))
			}
			values.WriteString(fmt.Sprintf("\n\t%s.URLPath = %s", rectIdent, __gong__toRawStringLiteral(rect.URLPath)))
			if rect.URLTarget.ToCodeString() != "" {
				values.WriteString(fmt.Sprintf("\n\t%s.URLTarget = models.%s", rectIdent, rect.URLTarget.ToCodeString()))
			} else {
				values.WriteString(fmt.Sprintf("\n\t%s.URLTarget = \"\"", rectIdent))
			}
			for _, elem := range rect.Peers {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := elem.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Peers = append(%s.Peers, %s)", rectIdent, rectIdent, targetIdent))
			}
			if rect.EnclosingRect != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := rect.EnclosingRect.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.EnclosingRect = %s", rectIdent, targetIdent))
			}
			for _, elem := range rect.Obstacles {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := elem.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Obstacles = append(%s.Obstacles, %s)", rectIdent, rectIdent, targetIdent))
			}
			if rect.AnchoredTo != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := rect.AnchoredTo.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.AnchoredTo = %s", rectIdent, targetIdent))
			}
			for _, elem := range rect.HoveringTrigger {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := elem.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.HoveringTrigger = append(%s.HoveringTrigger, %s)", rectIdent, rectIdent, targetIdent))
			}
			for _, elem := range rect.DisplayConditions {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := elem.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.DisplayConditions = append(%s.DisplayConditions, %s)", rectIdent, rectIdent, targetIdent))
			}
			for _, elem := range rect.Animations {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := elem.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Animations = append(%s.Animations, %s)", rectIdent, rectIdent, targetIdent))
			}
			for _, elem := range rect.RectAnchoredTexts {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := elem.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.RectAnchoredTexts = append(%s.RectAnchoredTexts, %s)", rectIdent, rectIdent, targetIdent))
			}
			for _, elem := range rect.RectAnchoredRects {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := elem.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.RectAnchoredRects = append(%s.RectAnchoredRects, %s)", rectIdent, rectIdent, targetIdent))
			}
			for _, elem := range rect.RectAnchoredPaths {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := elem.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.RectAnchoredPaths = append(%s.RectAnchoredPaths, %s)", rectIdent, rectIdent, targetIdent))
			}
			for _, elem := range rect.RectAnchoredPngImages {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := elem.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.RectAnchoredPngImages = append(%s.RectAnchoredPngImages, %s)", rectIdent, rectIdent, targetIdent))
			}
		}
	}
	if stage != nil {
		for _, rectanchoredpath := range __gong__sortStageSetInstances(stage.RectAnchoredPaths, stage.RectAnchoredPath_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			rectanchoredpathIdent := rectanchoredpath.GongGetIdentifier(stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&models.RectAnchoredPath{Name: %s}).Stage(stage)", rectanchoredpathIdent, __gong__toRawStringLiteral(rectanchoredpath.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", rectanchoredpathIdent, __gong__toRawStringLiteral(rectanchoredpath.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Definition = %s", rectanchoredpathIdent, __gong__toRawStringLiteral(rectanchoredpath.Definition)))
			values.WriteString(fmt.Sprintf("\n\t%s.X_Offset = %f", rectanchoredpathIdent, rectanchoredpath.X_Offset))
			values.WriteString(fmt.Sprintf("\n\t%s.Y_Offset = %f", rectanchoredpathIdent, rectanchoredpath.Y_Offset))
			if rectanchoredpath.RectAnchorType.ToCodeString() != "" {
				values.WriteString(fmt.Sprintf("\n\t%s.RectAnchorType = models.%s", rectanchoredpathIdent, rectanchoredpath.RectAnchorType.ToCodeString()))
			} else {
				values.WriteString(fmt.Sprintf("\n\t%s.RectAnchorType = \"\"", rectanchoredpathIdent))
			}
			values.WriteString(fmt.Sprintf("\n\t%s.ScalePropotionnally = %t", rectanchoredpathIdent, rectanchoredpath.ScalePropotionnally))
			values.WriteString(fmt.Sprintf("\n\t%s.AppliedScaling = %f", rectanchoredpathIdent, rectanchoredpath.AppliedScaling))
			values.WriteString(fmt.Sprintf("\n\t%s.Color = %s", rectanchoredpathIdent, __gong__toRawStringLiteral(rectanchoredpath.Color)))
			values.WriteString(fmt.Sprintf("\n\t%s.FillOpacity = %f", rectanchoredpathIdent, rectanchoredpath.FillOpacity))
			values.WriteString(fmt.Sprintf("\n\t%s.Stroke = %s", rectanchoredpathIdent, __gong__toRawStringLiteral(rectanchoredpath.Stroke)))
			values.WriteString(fmt.Sprintf("\n\t%s.StrokeOpacity = %f", rectanchoredpathIdent, rectanchoredpath.StrokeOpacity))
			values.WriteString(fmt.Sprintf("\n\t%s.StrokeWidth = %f", rectanchoredpathIdent, rectanchoredpath.StrokeWidth))
			values.WriteString(fmt.Sprintf("\n\t%s.StrokeDashArray = %s", rectanchoredpathIdent, __gong__toRawStringLiteral(rectanchoredpath.StrokeDashArray)))
			values.WriteString(fmt.Sprintf("\n\t%s.StrokeDashArrayWhenSelected = %s", rectanchoredpathIdent, __gong__toRawStringLiteral(rectanchoredpath.StrokeDashArrayWhenSelected)))
			values.WriteString(fmt.Sprintf("\n\t%s.Transform = %s", rectanchoredpathIdent, __gong__toRawStringLiteral(rectanchoredpath.Transform)))
		}
	}
	if stage != nil {
		for _, rectanchoredpngimage := range __gong__sortStageSetInstances(stage.RectAnchoredPngImages, stage.RectAnchoredPngImage_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			rectanchoredpngimageIdent := rectanchoredpngimage.GongGetIdentifier(stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&models.RectAnchoredPngImage{Name: %s}).Stage(stage)", rectanchoredpngimageIdent, __gong__toRawStringLiteral(rectanchoredpngimage.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", rectanchoredpngimageIdent, __gong__toRawStringLiteral(rectanchoredpngimage.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.X = %f", rectanchoredpngimageIdent, rectanchoredpngimage.X))
			values.WriteString(fmt.Sprintf("\n\t%s.Y = %f", rectanchoredpngimageIdent, rectanchoredpngimage.Y))
			values.WriteString(fmt.Sprintf("\n\t%s.Width = %f", rectanchoredpngimageIdent, rectanchoredpngimage.Width))
			values.WriteString(fmt.Sprintf("\n\t%s.Height = %f", rectanchoredpngimageIdent, rectanchoredpngimage.Height))
			values.WriteString(fmt.Sprintf("\n\t%s.RX = %f", rectanchoredpngimageIdent, rectanchoredpngimage.RX))
			values.WriteString(fmt.Sprintf("\n\t%s.X_Offset = %f", rectanchoredpngimageIdent, rectanchoredpngimage.X_Offset))
			values.WriteString(fmt.Sprintf("\n\t%s.Y_Offset = %f", rectanchoredpngimageIdent, rectanchoredpngimage.Y_Offset))
			if rectanchoredpngimage.RectAnchorType.ToCodeString() != "" {
				values.WriteString(fmt.Sprintf("\n\t%s.RectAnchorType = models.%s", rectanchoredpngimageIdent, rectanchoredpngimage.RectAnchorType.ToCodeString()))
			} else {
				values.WriteString(fmt.Sprintf("\n\t%s.RectAnchorType = \"\"", rectanchoredpngimageIdent))
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Base64Content = %s", rectanchoredpngimageIdent, __gong__toRawStringLiteral(rectanchoredpngimage.Base64Content)))
		}
	}
	if stage != nil {
		for _, rectanchoredrect := range __gong__sortStageSetInstances(stage.RectAnchoredRects, stage.RectAnchoredRect_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			rectanchoredrectIdent := rectanchoredrect.GongGetIdentifier(stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&models.RectAnchoredRect{Name: %s}).Stage(stage)", rectanchoredrectIdent, __gong__toRawStringLiteral(rectanchoredrect.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", rectanchoredrectIdent, __gong__toRawStringLiteral(rectanchoredrect.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.X = %f", rectanchoredrectIdent, rectanchoredrect.X))
			values.WriteString(fmt.Sprintf("\n\t%s.Y = %f", rectanchoredrectIdent, rectanchoredrect.Y))
			values.WriteString(fmt.Sprintf("\n\t%s.Width = %f", rectanchoredrectIdent, rectanchoredrect.Width))
			values.WriteString(fmt.Sprintf("\n\t%s.Height = %f", rectanchoredrectIdent, rectanchoredrect.Height))
			values.WriteString(fmt.Sprintf("\n\t%s.RX = %f", rectanchoredrectIdent, rectanchoredrect.RX))
			values.WriteString(fmt.Sprintf("\n\t%s.X_Offset = %f", rectanchoredrectIdent, rectanchoredrect.X_Offset))
			values.WriteString(fmt.Sprintf("\n\t%s.Y_Offset = %f", rectanchoredrectIdent, rectanchoredrect.Y_Offset))
			if rectanchoredrect.RectAnchorType.ToCodeString() != "" {
				values.WriteString(fmt.Sprintf("\n\t%s.RectAnchorType = models.%s", rectanchoredrectIdent, rectanchoredrect.RectAnchorType.ToCodeString()))
			} else {
				values.WriteString(fmt.Sprintf("\n\t%s.RectAnchorType = \"\"", rectanchoredrectIdent))
			}
			values.WriteString(fmt.Sprintf("\n\t%s.WidthFollowRect = %t", rectanchoredrectIdent, rectanchoredrect.WidthFollowRect))
			values.WriteString(fmt.Sprintf("\n\t%s.HeightFollowRect = %t", rectanchoredrectIdent, rectanchoredrect.HeightFollowRect))
			values.WriteString(fmt.Sprintf("\n\t%s.HasToolTip = %t", rectanchoredrectIdent, rectanchoredrect.HasToolTip))
			values.WriteString(fmt.Sprintf("\n\t%s.ToolTipText = %s", rectanchoredrectIdent, __gong__toRawStringLiteral(rectanchoredrect.ToolTipText)))
			values.WriteString(fmt.Sprintf("\n\t%s.Color = %s", rectanchoredrectIdent, __gong__toRawStringLiteral(rectanchoredrect.Color)))
			values.WriteString(fmt.Sprintf("\n\t%s.FillOpacity = %f", rectanchoredrectIdent, rectanchoredrect.FillOpacity))
			values.WriteString(fmt.Sprintf("\n\t%s.Stroke = %s", rectanchoredrectIdent, __gong__toRawStringLiteral(rectanchoredrect.Stroke)))
			values.WriteString(fmt.Sprintf("\n\t%s.StrokeOpacity = %f", rectanchoredrectIdent, rectanchoredrect.StrokeOpacity))
			values.WriteString(fmt.Sprintf("\n\t%s.StrokeWidth = %f", rectanchoredrectIdent, rectanchoredrect.StrokeWidth))
			values.WriteString(fmt.Sprintf("\n\t%s.StrokeDashArray = %s", rectanchoredrectIdent, __gong__toRawStringLiteral(rectanchoredrect.StrokeDashArray)))
			values.WriteString(fmt.Sprintf("\n\t%s.StrokeDashArrayWhenSelected = %s", rectanchoredrectIdent, __gong__toRawStringLiteral(rectanchoredrect.StrokeDashArrayWhenSelected)))
			values.WriteString(fmt.Sprintf("\n\t%s.Transform = %s", rectanchoredrectIdent, __gong__toRawStringLiteral(rectanchoredrect.Transform)))
		}
	}
	if stage != nil {
		for _, rectanchoredtext := range __gong__sortStageSetInstances(stage.RectAnchoredTexts, stage.RectAnchoredText_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			rectanchoredtextIdent := rectanchoredtext.GongGetIdentifier(stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&models.RectAnchoredText{Name: %s}).Stage(stage)", rectanchoredtextIdent, __gong__toRawStringLiteral(rectanchoredtext.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", rectanchoredtextIdent, __gong__toRawStringLiteral(rectanchoredtext.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Content = %s", rectanchoredtextIdent, __gong__toRawStringLiteral(rectanchoredtext.Content)))
			values.WriteString(fmt.Sprintf("\n\t%s.FontWeight = %s", rectanchoredtextIdent, __gong__toRawStringLiteral(rectanchoredtext.FontWeight)))
			values.WriteString(fmt.Sprintf("\n\t%s.FontSize = %s", rectanchoredtextIdent, __gong__toRawStringLiteral(rectanchoredtext.FontSize)))
			values.WriteString(fmt.Sprintf("\n\t%s.FontStyle = %s", rectanchoredtextIdent, __gong__toRawStringLiteral(rectanchoredtext.FontStyle)))
			values.WriteString(fmt.Sprintf("\n\t%s.LetterSpacing = %s", rectanchoredtextIdent, __gong__toRawStringLiteral(rectanchoredtext.LetterSpacing)))
			values.WriteString(fmt.Sprintf("\n\t%s.FontFamily = %s", rectanchoredtextIdent, __gong__toRawStringLiteral(rectanchoredtext.FontFamily)))
			if rectanchoredtext.WhiteSpace.ToCodeString() != "" {
				values.WriteString(fmt.Sprintf("\n\t%s.WhiteSpace = models.%s", rectanchoredtextIdent, rectanchoredtext.WhiteSpace.ToCodeString()))
			} else {
				values.WriteString(fmt.Sprintf("\n\t%s.WhiteSpace = \"\"", rectanchoredtextIdent))
			}
			values.WriteString(fmt.Sprintf("\n\t%s.X_Offset = %f", rectanchoredtextIdent, rectanchoredtext.X_Offset))
			values.WriteString(fmt.Sprintf("\n\t%s.Y_Offset = %f", rectanchoredtextIdent, rectanchoredtext.Y_Offset))
			if rectanchoredtext.RectAnchorType.ToCodeString() != "" {
				values.WriteString(fmt.Sprintf("\n\t%s.RectAnchorType = models.%s", rectanchoredtextIdent, rectanchoredtext.RectAnchorType.ToCodeString()))
			} else {
				values.WriteString(fmt.Sprintf("\n\t%s.RectAnchorType = \"\"", rectanchoredtextIdent))
			}
			if rectanchoredtext.TextAnchorType.ToCodeString() != "" {
				values.WriteString(fmt.Sprintf("\n\t%s.TextAnchorType = models.%s", rectanchoredtextIdent, rectanchoredtext.TextAnchorType.ToCodeString()))
			} else {
				values.WriteString(fmt.Sprintf("\n\t%s.TextAnchorType = \"\"", rectanchoredtextIdent))
			}
			if rectanchoredtext.DominantBaseline.ToCodeString() != "" {
				values.WriteString(fmt.Sprintf("\n\t%s.DominantBaseline = models.%s", rectanchoredtextIdent, rectanchoredtext.DominantBaseline.ToCodeString()))
			} else {
				values.WriteString(fmt.Sprintf("\n\t%s.DominantBaseline = \"\"", rectanchoredtextIdent))
			}
			if rectanchoredtext.WritingMode.ToCodeString() != "" {
				values.WriteString(fmt.Sprintf("\n\t%s.WritingMode = models.%s", rectanchoredtextIdent, rectanchoredtext.WritingMode.ToCodeString()))
			} else {
				values.WriteString(fmt.Sprintf("\n\t%s.WritingMode = \"\"", rectanchoredtextIdent))
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Color = %s", rectanchoredtextIdent, __gong__toRawStringLiteral(rectanchoredtext.Color)))
			values.WriteString(fmt.Sprintf("\n\t%s.FillOpacity = %f", rectanchoredtextIdent, rectanchoredtext.FillOpacity))
			values.WriteString(fmt.Sprintf("\n\t%s.Stroke = %s", rectanchoredtextIdent, __gong__toRawStringLiteral(rectanchoredtext.Stroke)))
			values.WriteString(fmt.Sprintf("\n\t%s.StrokeOpacity = %f", rectanchoredtextIdent, rectanchoredtext.StrokeOpacity))
			values.WriteString(fmt.Sprintf("\n\t%s.StrokeWidth = %f", rectanchoredtextIdent, rectanchoredtext.StrokeWidth))
			values.WriteString(fmt.Sprintf("\n\t%s.StrokeDashArray = %s", rectanchoredtextIdent, __gong__toRawStringLiteral(rectanchoredtext.StrokeDashArray)))
			values.WriteString(fmt.Sprintf("\n\t%s.StrokeDashArrayWhenSelected = %s", rectanchoredtextIdent, __gong__toRawStringLiteral(rectanchoredtext.StrokeDashArrayWhenSelected)))
			values.WriteString(fmt.Sprintf("\n\t%s.Transform = %s", rectanchoredtextIdent, __gong__toRawStringLiteral(rectanchoredtext.Transform)))
			values.WriteString(fmt.Sprintf("\n\t%s.URLPath = %s", rectanchoredtextIdent, __gong__toRawStringLiteral(rectanchoredtext.URLPath)))
			if rectanchoredtext.URLTarget.ToCodeString() != "" {
				values.WriteString(fmt.Sprintf("\n\t%s.URLTarget = models.%s", rectanchoredtextIdent, rectanchoredtext.URLTarget.ToCodeString()))
			} else {
				values.WriteString(fmt.Sprintf("\n\t%s.URLTarget = \"\"", rectanchoredtextIdent))
			}
			for _, elem := range rectanchoredtext.Animates {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := elem.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Animates = append(%s.Animates, %s)", rectanchoredtextIdent, rectanchoredtextIdent, targetIdent))
			}
		}
	}
	if stage != nil {
		for _, rectlinklink := range __gong__sortStageSetInstances(stage.RectLinkLinks, stage.RectLinkLink_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			rectlinklinkIdent := rectlinklink.GongGetIdentifier(stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&models.RectLinkLink{Name: %s}).Stage(stage)", rectlinklinkIdent, __gong__toRawStringLiteral(rectlinklink.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", rectlinklinkIdent, __gong__toRawStringLiteral(rectlinklink.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.TargetAnchorPosition = %f", rectlinklinkIdent, rectlinklink.TargetAnchorPosition))
			values.WriteString(fmt.Sprintf("\n\t%s.Color = %s", rectlinklinkIdent, __gong__toRawStringLiteral(rectlinklink.Color)))
			values.WriteString(fmt.Sprintf("\n\t%s.FillOpacity = %f", rectlinklinkIdent, rectlinklink.FillOpacity))
			values.WriteString(fmt.Sprintf("\n\t%s.Stroke = %s", rectlinklinkIdent, __gong__toRawStringLiteral(rectlinklink.Stroke)))
			values.WriteString(fmt.Sprintf("\n\t%s.StrokeOpacity = %f", rectlinklinkIdent, rectlinklink.StrokeOpacity))
			values.WriteString(fmt.Sprintf("\n\t%s.StrokeWidth = %f", rectlinklinkIdent, rectlinklink.StrokeWidth))
			values.WriteString(fmt.Sprintf("\n\t%s.StrokeDashArray = %s", rectlinklinkIdent, __gong__toRawStringLiteral(rectlinklink.StrokeDashArray)))
			values.WriteString(fmt.Sprintf("\n\t%s.StrokeDashArrayWhenSelected = %s", rectlinklinkIdent, __gong__toRawStringLiteral(rectlinklink.StrokeDashArrayWhenSelected)))
			values.WriteString(fmt.Sprintf("\n\t%s.Transform = %s", rectlinklinkIdent, __gong__toRawStringLiteral(rectlinklink.Transform)))
			if rectlinklink.Start != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := rectlinklink.Start.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Start = %s", rectlinklinkIdent, targetIdent))
			}
			if rectlinklink.End != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := rectlinklink.End.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.End = %s", rectlinklinkIdent, targetIdent))
			}
		}
	}
	if stage != nil {
		for _, svg := range __gong__sortStageSetInstances(stage.SVGs, stage.SVG_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			svgIdent := svg.GongGetIdentifier(stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&models.SVG{Name: %s}).Stage(stage)", svgIdent, __gong__toRawStringLiteral(svg.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", svgIdent, __gong__toRawStringLiteral(svg.Name)))
			if svg.DrawingState.ToCodeString() != "" {
				values.WriteString(fmt.Sprintf("\n\t%s.DrawingState = models.%s", svgIdent, svg.DrawingState.ToCodeString()))
			} else {
				values.WriteString(fmt.Sprintf("\n\t%s.DrawingState = \"\"", svgIdent))
			}
			values.WriteString(fmt.Sprintf("\n\t%s.IsEditable = %t", svgIdent, svg.IsEditable))
			values.WriteString(fmt.Sprintf("\n\t%s.IsSVGFrontEndFileGenerated = %t", svgIdent, svg.IsSVGFrontEndFileGenerated))
			values.WriteString(fmt.Sprintf("\n\t%s.IsSVGBackEndFileGenerated = %t", svgIdent, svg.IsSVGBackEndFileGenerated))
			values.WriteString(fmt.Sprintf("\n\t%s.DefaultDirectoryForGeneratedImages = %s", svgIdent, __gong__toRawStringLiteral(svg.DefaultDirectoryForGeneratedImages)))
			values.WriteString(fmt.Sprintf("\n\t%s.IsControlBannerHidden = %t", svgIdent, svg.IsControlBannerHidden))
			values.WriteString(fmt.Sprintf("\n\t%s.PanX = %f", svgIdent, svg.PanX))
			values.WriteString(fmt.Sprintf("\n\t%s.PanY = %f", svgIdent, svg.PanY))
			values.WriteString(fmt.Sprintf("\n\t%s.Zoom = %f", svgIdent, svg.Zoom))
			values.WriteString(fmt.Sprintf("\n\t%s.OverrideWidth = %t", svgIdent, svg.OverrideWidth))
			values.WriteString(fmt.Sprintf("\n\t%s.OverriddenWidth = %f", svgIdent, svg.OverriddenWidth))
			values.WriteString(fmt.Sprintf("\n\t%s.OverrideHeight = %t", svgIdent, svg.OverrideHeight))
			values.WriteString(fmt.Sprintf("\n\t%s.OverriddenHeight = %f", svgIdent, svg.OverriddenHeight))
			for _, elem := range svg.Layers {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := elem.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Layers = append(%s.Layers, %s)", svgIdent, svgIdent, targetIdent))
			}
			if svg.StartRect != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := svg.StartRect.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.StartRect = %s", svgIdent, targetIdent))
			}
			if svg.EndRect != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := svg.EndRect.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.EndRect = %s", svgIdent, targetIdent))
			}
		}
	}
	if stage != nil {
		for _, svgtext := range __gong__sortStageSetInstances(stage.SvgTexts, stage.SvgText_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			svgtextIdent := svgtext.GongGetIdentifier(stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&models.SvgText{Name: %s}).Stage(stage)", svgtextIdent, __gong__toRawStringLiteral(svgtext.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", svgtextIdent, __gong__toRawStringLiteral(svgtext.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Text = %s", svgtextIdent, __gong__toRawStringLiteral(svgtext.Text)))
		}
	}
	if stage != nil {
		for _, text := range __gong__sortStageSetInstances(stage.Texts, stage.Text_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			textIdent := text.GongGetIdentifier(stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&models.Text{Name: %s}).Stage(stage)", textIdent, __gong__toRawStringLiteral(text.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", textIdent, __gong__toRawStringLiteral(text.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.X = %f", textIdent, text.X))
			values.WriteString(fmt.Sprintf("\n\t%s.Y = %f", textIdent, text.Y))
			values.WriteString(fmt.Sprintf("\n\t%s.Content = %s", textIdent, __gong__toRawStringLiteral(text.Content)))
			values.WriteString(fmt.Sprintf("\n\t%s.Color = %s", textIdent, __gong__toRawStringLiteral(text.Color)))
			values.WriteString(fmt.Sprintf("\n\t%s.FillOpacity = %f", textIdent, text.FillOpacity))
			values.WriteString(fmt.Sprintf("\n\t%s.Stroke = %s", textIdent, __gong__toRawStringLiteral(text.Stroke)))
			values.WriteString(fmt.Sprintf("\n\t%s.StrokeOpacity = %f", textIdent, text.StrokeOpacity))
			values.WriteString(fmt.Sprintf("\n\t%s.StrokeWidth = %f", textIdent, text.StrokeWidth))
			values.WriteString(fmt.Sprintf("\n\t%s.StrokeDashArray = %s", textIdent, __gong__toRawStringLiteral(text.StrokeDashArray)))
			values.WriteString(fmt.Sprintf("\n\t%s.StrokeDashArrayWhenSelected = %s", textIdent, __gong__toRawStringLiteral(text.StrokeDashArrayWhenSelected)))
			values.WriteString(fmt.Sprintf("\n\t%s.Transform = %s", textIdent, __gong__toRawStringLiteral(text.Transform)))
			values.WriteString(fmt.Sprintf("\n\t%s.FontWeight = %s", textIdent, __gong__toRawStringLiteral(text.FontWeight)))
			values.WriteString(fmt.Sprintf("\n\t%s.FontSize = %s", textIdent, __gong__toRawStringLiteral(text.FontSize)))
			values.WriteString(fmt.Sprintf("\n\t%s.FontStyle = %s", textIdent, __gong__toRawStringLiteral(text.FontStyle)))
			values.WriteString(fmt.Sprintf("\n\t%s.LetterSpacing = %s", textIdent, __gong__toRawStringLiteral(text.LetterSpacing)))
			values.WriteString(fmt.Sprintf("\n\t%s.FontFamily = %s", textIdent, __gong__toRawStringLiteral(text.FontFamily)))
			if text.WhiteSpace.ToCodeString() != "" {
				values.WriteString(fmt.Sprintf("\n\t%s.WhiteSpace = models.%s", textIdent, text.WhiteSpace.ToCodeString()))
			} else {
				values.WriteString(fmt.Sprintf("\n\t%s.WhiteSpace = \"\"", textIdent))
			}
			for _, elem := range text.Animates {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := elem.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Animates = append(%s.Animates, %s)", textIdent, textIdent, targetIdent))
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
go run github.com/fullstack-lang/gong/lib/svg/go/models/probe/cmd/probe@%s stage.go
*/
package %s

import (
	"slices"
	"time"

	"github.com/fullstack-lang/gong/lib/svg/go/models"%s
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

	if modelsPackageName != "" && modelsPackageName != "github.com/fullstack-lang/gong/lib/svg/go/models" {
		res = strings.ReplaceAll(res, "\"github.com/fullstack-lang/gong/lib/svg/go/models\"", fmt.Sprintf("%q", modelsPackageName))
	}

	return res, nil
}

// insertion point for marshall field methods
func (animate *Animate) GongMarshallField(stage *Stage, fieldName string) (res string) {
	ident := animate.GongGetIdentifier(stage)
	_ = ident
	switch fieldName {
	case "Name":
		res = __gong__marshallString(ident, "Name", animate.Name)
	case "AttributeName":
		res = __gong__marshallString(ident, "AttributeName", animate.AttributeName)
	case "Values":
		res = __gong__marshallString(ident, "Values", animate.Values)
	case "From":
		res = __gong__marshallString(ident, "From", animate.From)
	case "To":
		res = __gong__marshallString(ident, "To", animate.To)
	case "Dur":
		res = __gong__marshallString(ident, "Dur", animate.Dur)
	case "RepeatCount":
		res = __gong__marshallString(ident, "RepeatCount", animate.RepeatCount)

	default:
		log.Panicf("Unknown field %s for Gongstruct Animate", fieldName)
	}
	return
}

func (circle *Circle) GongMarshallField(stage *Stage, fieldName string) (res string) {
	ident := circle.GongGetIdentifier(stage)
	_ = ident
	switch fieldName {
	case "Name":
		res = __gong__marshallString(ident, "Name", circle.Name)
	case "CX":
		res = __gong__marshallFloat(ident, "CX", circle.CX)
	case "CY":
		res = __gong__marshallFloat(ident, "CY", circle.CY)
	case "Radius":
		res = __gong__marshallFloat(ident, "Radius", circle.Radius)
	case "Color":
		res = __gong__marshallString(ident, "Color", circle.Color)
	case "FillOpacity":
		res = __gong__marshallFloat(ident, "FillOpacity", circle.FillOpacity)
	case "Stroke":
		res = __gong__marshallString(ident, "Stroke", circle.Stroke)
	case "StrokeOpacity":
		res = __gong__marshallFloat(ident, "StrokeOpacity", circle.StrokeOpacity)
	case "StrokeWidth":
		res = __gong__marshallFloat(ident, "StrokeWidth", circle.StrokeWidth)
	case "StrokeDashArray":
		res = __gong__marshallString(ident, "StrokeDashArray", circle.StrokeDashArray)
	case "StrokeDashArrayWhenSelected":
		res = __gong__marshallString(ident, "StrokeDashArrayWhenSelected", circle.StrokeDashArrayWhenSelected)
	case "Transform":
		res = __gong__marshallString(ident, "Transform", circle.Transform)

	case "Animations":
		var sb strings.Builder
		for _, _animate := range circle.Animations {
			sb.WriteString(__gong__marshallSliceOfPointers(ident, "Animations", _animate.GongGetIdentifier(stage)))
		}
		res = sb.String()
	default:
		log.Panicf("Unknown field %s for Gongstruct Circle", fieldName)
	}
	return
}

func (condition *Condition) GongMarshallField(stage *Stage, fieldName string) (res string) {
	ident := condition.GongGetIdentifier(stage)
	_ = ident
	switch fieldName {
	case "Name":
		res = __gong__marshallString(ident, "Name", condition.Name)

	default:
		log.Panicf("Unknown field %s for Gongstruct Condition", fieldName)
	}
	return
}

func (controlpoint *ControlPoint) GongMarshallField(stage *Stage, fieldName string) (res string) {
	ident := controlpoint.GongGetIdentifier(stage)
	_ = ident
	switch fieldName {
	case "Name":
		res = __gong__marshallString(ident, "Name", controlpoint.Name)
	case "X_Relative":
		res = __gong__marshallFloat(ident, "X_Relative", controlpoint.X_Relative)
	case "Y_Relative":
		res = __gong__marshallFloat(ident, "Y_Relative", controlpoint.Y_Relative)

	case "ClosestRect":
		if controlpoint.ClosestRect != nil {
			res = __gong__marshallPointer(ident, "ClosestRect", controlpoint.ClosestRect.GongGetIdentifier(stage))
		} else {
			res = __gong__marshallPointer(ident, "ClosestRect", "nil")
		}
	default:
		log.Panicf("Unknown field %s for Gongstruct ControlPoint", fieldName)
	}
	return
}

func (ellipse *Ellipse) GongMarshallField(stage *Stage, fieldName string) (res string) {
	ident := ellipse.GongGetIdentifier(stage)
	_ = ident
	switch fieldName {
	case "Name":
		res = __gong__marshallString(ident, "Name", ellipse.Name)
	case "CX":
		res = __gong__marshallFloat(ident, "CX", ellipse.CX)
	case "CY":
		res = __gong__marshallFloat(ident, "CY", ellipse.CY)
	case "RX":
		res = __gong__marshallFloat(ident, "RX", ellipse.RX)
	case "RY":
		res = __gong__marshallFloat(ident, "RY", ellipse.RY)
	case "Color":
		res = __gong__marshallString(ident, "Color", ellipse.Color)
	case "FillOpacity":
		res = __gong__marshallFloat(ident, "FillOpacity", ellipse.FillOpacity)
	case "Stroke":
		res = __gong__marshallString(ident, "Stroke", ellipse.Stroke)
	case "StrokeOpacity":
		res = __gong__marshallFloat(ident, "StrokeOpacity", ellipse.StrokeOpacity)
	case "StrokeWidth":
		res = __gong__marshallFloat(ident, "StrokeWidth", ellipse.StrokeWidth)
	case "StrokeDashArray":
		res = __gong__marshallString(ident, "StrokeDashArray", ellipse.StrokeDashArray)
	case "StrokeDashArrayWhenSelected":
		res = __gong__marshallString(ident, "StrokeDashArrayWhenSelected", ellipse.StrokeDashArrayWhenSelected)
	case "Transform":
		res = __gong__marshallString(ident, "Transform", ellipse.Transform)

	case "Animates":
		var sb strings.Builder
		for _, _animate := range ellipse.Animates {
			sb.WriteString(__gong__marshallSliceOfPointers(ident, "Animates", _animate.GongGetIdentifier(stage)))
		}
		res = sb.String()
	default:
		log.Panicf("Unknown field %s for Gongstruct Ellipse", fieldName)
	}
	return
}

func (filetodownload *FileToDownload) GongMarshallField(stage *Stage, fieldName string) (res string) {
	ident := filetodownload.GongGetIdentifier(stage)
	_ = ident
	switch fieldName {
	case "Name":
		res = __gong__marshallString(ident, "Name", filetodownload.Name)
	case "Base64EncodedContent":
		res = __gong__marshallString(ident, "Base64EncodedContent", filetodownload.Base64EncodedContent)

	default:
		log.Panicf("Unknown field %s for Gongstruct FileToDownload", fieldName)
	}
	return
}

func (layer *Layer) GongMarshallField(stage *Stage, fieldName string) (res string) {
	ident := layer.GongGetIdentifier(stage)
	_ = ident
	switch fieldName {
	case "Name":
		res = __gong__marshallString(ident, "Name", layer.Name)

	case "Rects":
		var sb strings.Builder
		for _, _rect := range layer.Rects {
			sb.WriteString(__gong__marshallSliceOfPointers(ident, "Rects", _rect.GongGetIdentifier(stage)))
		}
		res = sb.String()
	case "Texts":
		var sb strings.Builder
		for _, _text := range layer.Texts {
			sb.WriteString(__gong__marshallSliceOfPointers(ident, "Texts", _text.GongGetIdentifier(stage)))
		}
		res = sb.String()
	case "Circles":
		var sb strings.Builder
		for _, _circle := range layer.Circles {
			sb.WriteString(__gong__marshallSliceOfPointers(ident, "Circles", _circle.GongGetIdentifier(stage)))
		}
		res = sb.String()
	case "Lines":
		var sb strings.Builder
		for _, _line := range layer.Lines {
			sb.WriteString(__gong__marshallSliceOfPointers(ident, "Lines", _line.GongGetIdentifier(stage)))
		}
		res = sb.String()
	case "Ellipses":
		var sb strings.Builder
		for _, _ellipse := range layer.Ellipses {
			sb.WriteString(__gong__marshallSliceOfPointers(ident, "Ellipses", _ellipse.GongGetIdentifier(stage)))
		}
		res = sb.String()
	case "Polylines":
		var sb strings.Builder
		for _, _polyline := range layer.Polylines {
			sb.WriteString(__gong__marshallSliceOfPointers(ident, "Polylines", _polyline.GongGetIdentifier(stage)))
		}
		res = sb.String()
	case "Polygones":
		var sb strings.Builder
		for _, _polygone := range layer.Polygones {
			sb.WriteString(__gong__marshallSliceOfPointers(ident, "Polygones", _polygone.GongGetIdentifier(stage)))
		}
		res = sb.String()
	case "Paths":
		var sb strings.Builder
		for _, _path := range layer.Paths {
			sb.WriteString(__gong__marshallSliceOfPointers(ident, "Paths", _path.GongGetIdentifier(stage)))
		}
		res = sb.String()
	case "Links":
		var sb strings.Builder
		for _, _link := range layer.Links {
			sb.WriteString(__gong__marshallSliceOfPointers(ident, "Links", _link.GongGetIdentifier(stage)))
		}
		res = sb.String()
	case "RectLinkLinks":
		var sb strings.Builder
		for _, _rectlinklink := range layer.RectLinkLinks {
			sb.WriteString(__gong__marshallSliceOfPointers(ident, "RectLinkLinks", _rectlinklink.GongGetIdentifier(stage)))
		}
		res = sb.String()
	default:
		log.Panicf("Unknown field %s for Gongstruct Layer", fieldName)
	}
	return
}

func (line *Line) GongMarshallField(stage *Stage, fieldName string) (res string) {
	ident := line.GongGetIdentifier(stage)
	_ = ident
	switch fieldName {
	case "Name":
		res = __gong__marshallString(ident, "Name", line.Name)
	case "X1":
		res = __gong__marshallFloat(ident, "X1", line.X1)
	case "Y1":
		res = __gong__marshallFloat(ident, "Y1", line.Y1)
	case "X2":
		res = __gong__marshallFloat(ident, "X2", line.X2)
	case "Y2":
		res = __gong__marshallFloat(ident, "Y2", line.Y2)
	case "Color":
		res = __gong__marshallString(ident, "Color", line.Color)
	case "FillOpacity":
		res = __gong__marshallFloat(ident, "FillOpacity", line.FillOpacity)
	case "Stroke":
		res = __gong__marshallString(ident, "Stroke", line.Stroke)
	case "StrokeOpacity":
		res = __gong__marshallFloat(ident, "StrokeOpacity", line.StrokeOpacity)
	case "StrokeWidth":
		res = __gong__marshallFloat(ident, "StrokeWidth", line.StrokeWidth)
	case "StrokeDashArray":
		res = __gong__marshallString(ident, "StrokeDashArray", line.StrokeDashArray)
	case "StrokeDashArrayWhenSelected":
		res = __gong__marshallString(ident, "StrokeDashArrayWhenSelected", line.StrokeDashArrayWhenSelected)
	case "Transform":
		res = __gong__marshallString(ident, "Transform", line.Transform)
	case "MouseClickX":
		res = __gong__marshallFloat(ident, "MouseClickX", line.MouseClickX)
	case "MouseClickY":
		res = __gong__marshallFloat(ident, "MouseClickY", line.MouseClickY)

	case "Animates":
		var sb strings.Builder
		for _, _animate := range line.Animates {
			sb.WriteString(__gong__marshallSliceOfPointers(ident, "Animates", _animate.GongGetIdentifier(stage)))
		}
		res = sb.String()
	default:
		log.Panicf("Unknown field %s for Gongstruct Line", fieldName)
	}
	return
}

func (link *Link) GongMarshallField(stage *Stage, fieldName string) (res string) {
	ident := link.GongGetIdentifier(stage)
	_ = ident
	switch fieldName {
	case "Name":
		res = __gong__marshallString(ident, "Name", link.Name)
	case "Type":
		res = __gong__marshallEnumString(ident, "Type", link.Type.ToCodeString())
	case "IsBezierCurve":
		res = __gong__marshallBool(ident, "IsBezierCurve", link.IsBezierCurve)
	case "StartAnchorType":
		res = __gong__marshallEnumString(ident, "StartAnchorType", link.StartAnchorType.ToCodeString())
	case "EndAnchorType":
		res = __gong__marshallEnumString(ident, "EndAnchorType", link.EndAnchorType.ToCodeString())
	case "StartOrientation":
		res = __gong__marshallEnumString(ident, "StartOrientation", link.StartOrientation.ToCodeString())
	case "StartRatio":
		res = __gong__marshallFloat(ident, "StartRatio", link.StartRatio)
	case "EndOrientation":
		res = __gong__marshallEnumString(ident, "EndOrientation", link.EndOrientation.ToCodeString())
	case "EndRatio":
		res = __gong__marshallFloat(ident, "EndRatio", link.EndRatio)
	case "CornerOffsetRatio":
		res = __gong__marshallFloat(ident, "CornerOffsetRatio", link.CornerOffsetRatio)
	case "CornerRadius":
		res = __gong__marshallFloat(ident, "CornerRadius", link.CornerRadius)
	case "HasEndArrow":
		res = __gong__marshallBool(ident, "HasEndArrow", link.HasEndArrow)
	case "EndArrowSize":
		res = __gong__marshallFloat(ident, "EndArrowSize", link.EndArrowSize)
	case "EndArrowOffset":
		res = __gong__marshallFloat(ident, "EndArrowOffset", link.EndArrowOffset)
	case "HasStartArrow":
		res = __gong__marshallBool(ident, "HasStartArrow", link.HasStartArrow)
	case "StartArrowSize":
		res = __gong__marshallFloat(ident, "StartArrowSize", link.StartArrowSize)
	case "StartArrowOffset":
		res = __gong__marshallFloat(ident, "StartArrowOffset", link.StartArrowOffset)
	case "Color":
		res = __gong__marshallString(ident, "Color", link.Color)
	case "FillOpacity":
		res = __gong__marshallFloat(ident, "FillOpacity", link.FillOpacity)
	case "Stroke":
		res = __gong__marshallString(ident, "Stroke", link.Stroke)
	case "StrokeOpacity":
		res = __gong__marshallFloat(ident, "StrokeOpacity", link.StrokeOpacity)
	case "StrokeWidth":
		res = __gong__marshallFloat(ident, "StrokeWidth", link.StrokeWidth)
	case "StrokeDashArray":
		res = __gong__marshallString(ident, "StrokeDashArray", link.StrokeDashArray)
	case "StrokeDashArrayWhenSelected":
		res = __gong__marshallString(ident, "StrokeDashArrayWhenSelected", link.StrokeDashArrayWhenSelected)
	case "Transform":
		res = __gong__marshallString(ident, "Transform", link.Transform)
	case "MouseX":
		res = __gong__marshallFloat(ident, "MouseX", link.MouseX)
	case "MouseY":
		res = __gong__marshallFloat(ident, "MouseY", link.MouseY)
	case "MouseEventKey":
		res = __gong__marshallEnumString(ident, "MouseEventKey", link.MouseEventKey.ToCodeString())

	case "Start":
		if link.Start != nil {
			res = __gong__marshallPointer(ident, "Start", link.Start.GongGetIdentifier(stage))
		} else {
			res = __gong__marshallPointer(ident, "Start", "nil")
		}
	case "End":
		if link.End != nil {
			res = __gong__marshallPointer(ident, "End", link.End.GongGetIdentifier(stage))
		} else {
			res = __gong__marshallPointer(ident, "End", "nil")
		}
	case "TextAtArrowStart":
		var sb strings.Builder
		for _, _linkanchoredtext := range link.TextAtArrowStart {
			sb.WriteString(__gong__marshallSliceOfPointers(ident, "TextAtArrowStart", _linkanchoredtext.GongGetIdentifier(stage)))
		}
		res = sb.String()
	case "TextAtArrowEnd":
		var sb strings.Builder
		for _, _linkanchoredtext := range link.TextAtArrowEnd {
			sb.WriteString(__gong__marshallSliceOfPointers(ident, "TextAtArrowEnd", _linkanchoredtext.GongGetIdentifier(stage)))
		}
		res = sb.String()
	case "TextAtCorner":
		var sb strings.Builder
		for _, _linkanchoredtext := range link.TextAtCorner {
			sb.WriteString(__gong__marshallSliceOfPointers(ident, "TextAtCorner", _linkanchoredtext.GongGetIdentifier(stage)))
		}
		res = sb.String()
	case "PathAtArrowStart":
		var sb strings.Builder
		for _, _linkanchoredpath := range link.PathAtArrowStart {
			sb.WriteString(__gong__marshallSliceOfPointers(ident, "PathAtArrowStart", _linkanchoredpath.GongGetIdentifier(stage)))
		}
		res = sb.String()
	case "PathAtArrowEnd":
		var sb strings.Builder
		for _, _linkanchoredpath := range link.PathAtArrowEnd {
			sb.WriteString(__gong__marshallSliceOfPointers(ident, "PathAtArrowEnd", _linkanchoredpath.GongGetIdentifier(stage)))
		}
		res = sb.String()
	case "PathAtCorner":
		var sb strings.Builder
		for _, _linkanchoredpath := range link.PathAtCorner {
			sb.WriteString(__gong__marshallSliceOfPointers(ident, "PathAtCorner", _linkanchoredpath.GongGetIdentifier(stage)))
		}
		res = sb.String()
	case "ControlPoints":
		var sb strings.Builder
		for _, _controlpoint := range link.ControlPoints {
			sb.WriteString(__gong__marshallSliceOfPointers(ident, "ControlPoints", _controlpoint.GongGetIdentifier(stage)))
		}
		res = sb.String()
	default:
		log.Panicf("Unknown field %s for Gongstruct Link", fieldName)
	}
	return
}

func (linkanchoredpath *LinkAnchoredPath) GongMarshallField(stage *Stage, fieldName string) (res string) {
	ident := linkanchoredpath.GongGetIdentifier(stage)
	_ = ident
	switch fieldName {
	case "Name":
		res = __gong__marshallString(ident, "Name", linkanchoredpath.Name)
	case "Definition":
		res = __gong__marshallString(ident, "Definition", linkanchoredpath.Definition)
	case "X_Offset":
		res = __gong__marshallFloat(ident, "X_Offset", linkanchoredpath.X_Offset)
	case "Y_Offset":
		res = __gong__marshallFloat(ident, "Y_Offset", linkanchoredpath.Y_Offset)
	case "ScalePropotionnally":
		res = __gong__marshallBool(ident, "ScalePropotionnally", linkanchoredpath.ScalePropotionnally)
	case "AppliedScaling":
		res = __gong__marshallFloat(ident, "AppliedScaling", linkanchoredpath.AppliedScaling)
	case "Color":
		res = __gong__marshallString(ident, "Color", linkanchoredpath.Color)
	case "FillOpacity":
		res = __gong__marshallFloat(ident, "FillOpacity", linkanchoredpath.FillOpacity)
	case "Stroke":
		res = __gong__marshallString(ident, "Stroke", linkanchoredpath.Stroke)
	case "StrokeOpacity":
		res = __gong__marshallFloat(ident, "StrokeOpacity", linkanchoredpath.StrokeOpacity)
	case "StrokeWidth":
		res = __gong__marshallFloat(ident, "StrokeWidth", linkanchoredpath.StrokeWidth)
	case "StrokeDashArray":
		res = __gong__marshallString(ident, "StrokeDashArray", linkanchoredpath.StrokeDashArray)
	case "StrokeDashArrayWhenSelected":
		res = __gong__marshallString(ident, "StrokeDashArrayWhenSelected", linkanchoredpath.StrokeDashArrayWhenSelected)
	case "Transform":
		res = __gong__marshallString(ident, "Transform", linkanchoredpath.Transform)

	default:
		log.Panicf("Unknown field %s for Gongstruct LinkAnchoredPath", fieldName)
	}
	return
}

func (linkanchoredtext *LinkAnchoredText) GongMarshallField(stage *Stage, fieldName string) (res string) {
	ident := linkanchoredtext.GongGetIdentifier(stage)
	_ = ident
	switch fieldName {
	case "Name":
		res = __gong__marshallString(ident, "Name", linkanchoredtext.Name)
	case "Content":
		res = __gong__marshallString(ident, "Content", linkanchoredtext.Content)
	case "AutomaticLayout":
		res = __gong__marshallBool(ident, "AutomaticLayout", linkanchoredtext.AutomaticLayout)
	case "LinkAnchorType":
		res = __gong__marshallEnumString(ident, "LinkAnchorType", linkanchoredtext.LinkAnchorType.ToCodeString())
	case "X_Offset":
		res = __gong__marshallFloat(ident, "X_Offset", linkanchoredtext.X_Offset)
	case "Y_Offset":
		res = __gong__marshallFloat(ident, "Y_Offset", linkanchoredtext.Y_Offset)
	case "FontWeight":
		res = __gong__marshallString(ident, "FontWeight", linkanchoredtext.FontWeight)
	case "FontSize":
		res = __gong__marshallString(ident, "FontSize", linkanchoredtext.FontSize)
	case "FontStyle":
		res = __gong__marshallString(ident, "FontStyle", linkanchoredtext.FontStyle)
	case "LetterSpacing":
		res = __gong__marshallString(ident, "LetterSpacing", linkanchoredtext.LetterSpacing)
	case "FontFamily":
		res = __gong__marshallString(ident, "FontFamily", linkanchoredtext.FontFamily)
	case "WhiteSpace":
		res = __gong__marshallEnumString(ident, "WhiteSpace", linkanchoredtext.WhiteSpace.ToCodeString())
	case "Color":
		res = __gong__marshallString(ident, "Color", linkanchoredtext.Color)
	case "FillOpacity":
		res = __gong__marshallFloat(ident, "FillOpacity", linkanchoredtext.FillOpacity)
	case "Stroke":
		res = __gong__marshallString(ident, "Stroke", linkanchoredtext.Stroke)
	case "StrokeOpacity":
		res = __gong__marshallFloat(ident, "StrokeOpacity", linkanchoredtext.StrokeOpacity)
	case "StrokeWidth":
		res = __gong__marshallFloat(ident, "StrokeWidth", linkanchoredtext.StrokeWidth)
	case "StrokeDashArray":
		res = __gong__marshallString(ident, "StrokeDashArray", linkanchoredtext.StrokeDashArray)
	case "StrokeDashArrayWhenSelected":
		res = __gong__marshallString(ident, "StrokeDashArrayWhenSelected", linkanchoredtext.StrokeDashArrayWhenSelected)
	case "Transform":
		res = __gong__marshallString(ident, "Transform", linkanchoredtext.Transform)

	case "Animates":
		var sb strings.Builder
		for _, _animate := range linkanchoredtext.Animates {
			sb.WriteString(__gong__marshallSliceOfPointers(ident, "Animates", _animate.GongGetIdentifier(stage)))
		}
		res = sb.String()
	default:
		log.Panicf("Unknown field %s for Gongstruct LinkAnchoredText", fieldName)
	}
	return
}

func (path *Path) GongMarshallField(stage *Stage, fieldName string) (res string) {
	ident := path.GongGetIdentifier(stage)
	_ = ident
	switch fieldName {
	case "Name":
		res = __gong__marshallString(ident, "Name", path.Name)
	case "Definition":
		res = __gong__marshallString(ident, "Definition", path.Definition)
	case "Color":
		res = __gong__marshallString(ident, "Color", path.Color)
	case "FillOpacity":
		res = __gong__marshallFloat(ident, "FillOpacity", path.FillOpacity)
	case "Stroke":
		res = __gong__marshallString(ident, "Stroke", path.Stroke)
	case "StrokeOpacity":
		res = __gong__marshallFloat(ident, "StrokeOpacity", path.StrokeOpacity)
	case "StrokeWidth":
		res = __gong__marshallFloat(ident, "StrokeWidth", path.StrokeWidth)
	case "StrokeDashArray":
		res = __gong__marshallString(ident, "StrokeDashArray", path.StrokeDashArray)
	case "StrokeDashArrayWhenSelected":
		res = __gong__marshallString(ident, "StrokeDashArrayWhenSelected", path.StrokeDashArrayWhenSelected)
	case "Transform":
		res = __gong__marshallString(ident, "Transform", path.Transform)

	case "Animates":
		var sb strings.Builder
		for _, _animate := range path.Animates {
			sb.WriteString(__gong__marshallSliceOfPointers(ident, "Animates", _animate.GongGetIdentifier(stage)))
		}
		res = sb.String()
	default:
		log.Panicf("Unknown field %s for Gongstruct Path", fieldName)
	}
	return
}

func (point *Point) GongMarshallField(stage *Stage, fieldName string) (res string) {
	ident := point.GongGetIdentifier(stage)
	_ = ident
	switch fieldName {
	case "Name":
		res = __gong__marshallString(ident, "Name", point.Name)
	case "X":
		res = __gong__marshallFloat(ident, "X", point.X)
	case "Y":
		res = __gong__marshallFloat(ident, "Y", point.Y)

	default:
		log.Panicf("Unknown field %s for Gongstruct Point", fieldName)
	}
	return
}

func (polygone *Polygone) GongMarshallField(stage *Stage, fieldName string) (res string) {
	ident := polygone.GongGetIdentifier(stage)
	_ = ident
	switch fieldName {
	case "Name":
		res = __gong__marshallString(ident, "Name", polygone.Name)
	case "Points":
		res = __gong__marshallString(ident, "Points", polygone.Points)
	case "Color":
		res = __gong__marshallString(ident, "Color", polygone.Color)
	case "FillOpacity":
		res = __gong__marshallFloat(ident, "FillOpacity", polygone.FillOpacity)
	case "Stroke":
		res = __gong__marshallString(ident, "Stroke", polygone.Stroke)
	case "StrokeOpacity":
		res = __gong__marshallFloat(ident, "StrokeOpacity", polygone.StrokeOpacity)
	case "StrokeWidth":
		res = __gong__marshallFloat(ident, "StrokeWidth", polygone.StrokeWidth)
	case "StrokeDashArray":
		res = __gong__marshallString(ident, "StrokeDashArray", polygone.StrokeDashArray)
	case "StrokeDashArrayWhenSelected":
		res = __gong__marshallString(ident, "StrokeDashArrayWhenSelected", polygone.StrokeDashArrayWhenSelected)
	case "Transform":
		res = __gong__marshallString(ident, "Transform", polygone.Transform)

	case "Animates":
		var sb strings.Builder
		for _, _animate := range polygone.Animates {
			sb.WriteString(__gong__marshallSliceOfPointers(ident, "Animates", _animate.GongGetIdentifier(stage)))
		}
		res = sb.String()
	default:
		log.Panicf("Unknown field %s for Gongstruct Polygone", fieldName)
	}
	return
}

func (polyline *Polyline) GongMarshallField(stage *Stage, fieldName string) (res string) {
	ident := polyline.GongGetIdentifier(stage)
	_ = ident
	switch fieldName {
	case "Name":
		res = __gong__marshallString(ident, "Name", polyline.Name)
	case "Points":
		res = __gong__marshallString(ident, "Points", polyline.Points)
	case "Color":
		res = __gong__marshallString(ident, "Color", polyline.Color)
	case "FillOpacity":
		res = __gong__marshallFloat(ident, "FillOpacity", polyline.FillOpacity)
	case "Stroke":
		res = __gong__marshallString(ident, "Stroke", polyline.Stroke)
	case "StrokeOpacity":
		res = __gong__marshallFloat(ident, "StrokeOpacity", polyline.StrokeOpacity)
	case "StrokeWidth":
		res = __gong__marshallFloat(ident, "StrokeWidth", polyline.StrokeWidth)
	case "StrokeDashArray":
		res = __gong__marshallString(ident, "StrokeDashArray", polyline.StrokeDashArray)
	case "StrokeDashArrayWhenSelected":
		res = __gong__marshallString(ident, "StrokeDashArrayWhenSelected", polyline.StrokeDashArrayWhenSelected)
	case "Transform":
		res = __gong__marshallString(ident, "Transform", polyline.Transform)

	case "Animates":
		var sb strings.Builder
		for _, _animate := range polyline.Animates {
			sb.WriteString(__gong__marshallSliceOfPointers(ident, "Animates", _animate.GongGetIdentifier(stage)))
		}
		res = sb.String()
	default:
		log.Panicf("Unknown field %s for Gongstruct Polyline", fieldName)
	}
	return
}

func (rect *Rect) GongMarshallField(stage *Stage, fieldName string) (res string) {
	ident := rect.GongGetIdentifier(stage)
	_ = ident
	switch fieldName {
	case "Name":
		res = __gong__marshallString(ident, "Name", rect.Name)
	case "X":
		res = __gong__marshallFloat(ident, "X", rect.X)
	case "Y":
		res = __gong__marshallFloat(ident, "Y", rect.Y)
	case "Width":
		res = __gong__marshallFloat(ident, "Width", rect.Width)
	case "Height":
		res = __gong__marshallFloat(ident, "Height", rect.Height)
	case "RX":
		res = __gong__marshallFloat(ident, "RX", rect.RX)
	case "Color":
		res = __gong__marshallString(ident, "Color", rect.Color)
	case "FillOpacity":
		res = __gong__marshallFloat(ident, "FillOpacity", rect.FillOpacity)
	case "Stroke":
		res = __gong__marshallString(ident, "Stroke", rect.Stroke)
	case "StrokeOpacity":
		res = __gong__marshallFloat(ident, "StrokeOpacity", rect.StrokeOpacity)
	case "StrokeWidth":
		res = __gong__marshallFloat(ident, "StrokeWidth", rect.StrokeWidth)
	case "StrokeDashArray":
		res = __gong__marshallString(ident, "StrokeDashArray", rect.StrokeDashArray)
	case "StrokeDashArrayWhenSelected":
		res = __gong__marshallString(ident, "StrokeDashArrayWhenSelected", rect.StrokeDashArrayWhenSelected)
	case "Transform":
		res = __gong__marshallString(ident, "Transform", rect.Transform)
	case "IsSelectable":
		res = __gong__marshallBool(ident, "IsSelectable", rect.IsSelectable)
	case "IsSelected":
		res = __gong__marshallBool(ident, "IsSelected", rect.IsSelected)
	case "CanHaveLeftHandle":
		res = __gong__marshallBool(ident, "CanHaveLeftHandle", rect.CanHaveLeftHandle)
	case "HasLeftHandle":
		res = __gong__marshallBool(ident, "HasLeftHandle", rect.HasLeftHandle)
	case "CanHaveRightHandle":
		res = __gong__marshallBool(ident, "CanHaveRightHandle", rect.CanHaveRightHandle)
	case "HasRightHandle":
		res = __gong__marshallBool(ident, "HasRightHandle", rect.HasRightHandle)
	case "CanHaveTopHandle":
		res = __gong__marshallBool(ident, "CanHaveTopHandle", rect.CanHaveTopHandle)
	case "HasTopHandle":
		res = __gong__marshallBool(ident, "HasTopHandle", rect.HasTopHandle)
	case "IsScalingProportionally":
		res = __gong__marshallBool(ident, "IsScalingProportionally", rect.IsScalingProportionally)
	case "CanHaveBottomHandle":
		res = __gong__marshallBool(ident, "CanHaveBottomHandle", rect.CanHaveBottomHandle)
	case "HasBottomHandle":
		res = __gong__marshallBool(ident, "HasBottomHandle", rect.HasBottomHandle)
	case "CanMoveHorizontaly":
		res = __gong__marshallBool(ident, "CanMoveHorizontaly", rect.CanMoveHorizontaly)
	case "CanMoveVerticaly":
		res = __gong__marshallBool(ident, "CanMoveVerticaly", rect.CanMoveVerticaly)
	case "ChangeColorWhenHovered":
		res = __gong__marshallBool(ident, "ChangeColorWhenHovered", rect.ChangeColorWhenHovered)
	case "ColorWhenHovered":
		res = __gong__marshallString(ident, "ColorWhenHovered", rect.ColorWhenHovered)
	case "OriginalColor":
		res = __gong__marshallString(ident, "OriginalColor", rect.OriginalColor)
	case "FillOpacityWhenHovered":
		res = __gong__marshallFloat(ident, "FillOpacityWhenHovered", rect.FillOpacityWhenHovered)
	case "OriginalFillOpacity":
		res = __gong__marshallFloat(ident, "OriginalFillOpacity", rect.OriginalFillOpacity)
	case "HasToolTip":
		res = __gong__marshallBool(ident, "HasToolTip", rect.HasToolTip)
	case "ToolTipText":
		res = __gong__marshallString(ident, "ToolTipText", rect.ToolTipText)
	case "ToolTipPosition":
		res = __gong__marshallEnumString(ident, "ToolTipPosition", rect.ToolTipPosition.ToCodeString())
	case "MouseX":
		res = __gong__marshallFloat(ident, "MouseX", rect.MouseX)
	case "MouseY":
		res = __gong__marshallFloat(ident, "MouseY", rect.MouseY)
	case "MouseEventKey":
		res = __gong__marshallEnumString(ident, "MouseEventKey", rect.MouseEventKey.ToCodeString())
	case "URLPath":
		res = __gong__marshallString(ident, "URLPath", rect.URLPath)
	case "URLTarget":
		res = __gong__marshallEnumString(ident, "URLTarget", rect.URLTarget.ToCodeString())

	case "Peers":
		var sb strings.Builder
		for _, _rect := range rect.Peers {
			sb.WriteString(__gong__marshallSliceOfPointers(ident, "Peers", _rect.GongGetIdentifier(stage)))
		}
		res = sb.String()
	case "EnclosingRect":
		if rect.EnclosingRect != nil {
			res = __gong__marshallPointer(ident, "EnclosingRect", rect.EnclosingRect.GongGetIdentifier(stage))
		} else {
			res = __gong__marshallPointer(ident, "EnclosingRect", "nil")
		}
	case "Obstacles":
		var sb strings.Builder
		for _, _rect := range rect.Obstacles {
			sb.WriteString(__gong__marshallSliceOfPointers(ident, "Obstacles", _rect.GongGetIdentifier(stage)))
		}
		res = sb.String()
	case "AnchoredTo":
		if rect.AnchoredTo != nil {
			res = __gong__marshallPointer(ident, "AnchoredTo", rect.AnchoredTo.GongGetIdentifier(stage))
		} else {
			res = __gong__marshallPointer(ident, "AnchoredTo", "nil")
		}
	case "HoveringTrigger":
		var sb strings.Builder
		for _, _condition := range rect.HoveringTrigger {
			sb.WriteString(__gong__marshallSliceOfPointers(ident, "HoveringTrigger", _condition.GongGetIdentifier(stage)))
		}
		res = sb.String()
	case "DisplayConditions":
		var sb strings.Builder
		for _, _condition := range rect.DisplayConditions {
			sb.WriteString(__gong__marshallSliceOfPointers(ident, "DisplayConditions", _condition.GongGetIdentifier(stage)))
		}
		res = sb.String()
	case "Animations":
		var sb strings.Builder
		for _, _animate := range rect.Animations {
			sb.WriteString(__gong__marshallSliceOfPointers(ident, "Animations", _animate.GongGetIdentifier(stage)))
		}
		res = sb.String()
	case "RectAnchoredTexts":
		var sb strings.Builder
		for _, _rectanchoredtext := range rect.RectAnchoredTexts {
			sb.WriteString(__gong__marshallSliceOfPointers(ident, "RectAnchoredTexts", _rectanchoredtext.GongGetIdentifier(stage)))
		}
		res = sb.String()
	case "RectAnchoredRects":
		var sb strings.Builder
		for _, _rectanchoredrect := range rect.RectAnchoredRects {
			sb.WriteString(__gong__marshallSliceOfPointers(ident, "RectAnchoredRects", _rectanchoredrect.GongGetIdentifier(stage)))
		}
		res = sb.String()
	case "RectAnchoredPaths":
		var sb strings.Builder
		for _, _rectanchoredpath := range rect.RectAnchoredPaths {
			sb.WriteString(__gong__marshallSliceOfPointers(ident, "RectAnchoredPaths", _rectanchoredpath.GongGetIdentifier(stage)))
		}
		res = sb.String()
	case "RectAnchoredPngImages":
		var sb strings.Builder
		for _, _rectanchoredpngimage := range rect.RectAnchoredPngImages {
			sb.WriteString(__gong__marshallSliceOfPointers(ident, "RectAnchoredPngImages", _rectanchoredpngimage.GongGetIdentifier(stage)))
		}
		res = sb.String()
	default:
		log.Panicf("Unknown field %s for Gongstruct Rect", fieldName)
	}
	return
}

func (rectanchoredpath *RectAnchoredPath) GongMarshallField(stage *Stage, fieldName string) (res string) {
	ident := rectanchoredpath.GongGetIdentifier(stage)
	_ = ident
	switch fieldName {
	case "Name":
		res = __gong__marshallString(ident, "Name", rectanchoredpath.Name)
	case "Definition":
		res = __gong__marshallString(ident, "Definition", rectanchoredpath.Definition)
	case "X_Offset":
		res = __gong__marshallFloat(ident, "X_Offset", rectanchoredpath.X_Offset)
	case "Y_Offset":
		res = __gong__marshallFloat(ident, "Y_Offset", rectanchoredpath.Y_Offset)
	case "RectAnchorType":
		res = __gong__marshallEnumString(ident, "RectAnchorType", rectanchoredpath.RectAnchorType.ToCodeString())
	case "ScalePropotionnally":
		res = __gong__marshallBool(ident, "ScalePropotionnally", rectanchoredpath.ScalePropotionnally)
	case "AppliedScaling":
		res = __gong__marshallFloat(ident, "AppliedScaling", rectanchoredpath.AppliedScaling)
	case "Color":
		res = __gong__marshallString(ident, "Color", rectanchoredpath.Color)
	case "FillOpacity":
		res = __gong__marshallFloat(ident, "FillOpacity", rectanchoredpath.FillOpacity)
	case "Stroke":
		res = __gong__marshallString(ident, "Stroke", rectanchoredpath.Stroke)
	case "StrokeOpacity":
		res = __gong__marshallFloat(ident, "StrokeOpacity", rectanchoredpath.StrokeOpacity)
	case "StrokeWidth":
		res = __gong__marshallFloat(ident, "StrokeWidth", rectanchoredpath.StrokeWidth)
	case "StrokeDashArray":
		res = __gong__marshallString(ident, "StrokeDashArray", rectanchoredpath.StrokeDashArray)
	case "StrokeDashArrayWhenSelected":
		res = __gong__marshallString(ident, "StrokeDashArrayWhenSelected", rectanchoredpath.StrokeDashArrayWhenSelected)
	case "Transform":
		res = __gong__marshallString(ident, "Transform", rectanchoredpath.Transform)

	default:
		log.Panicf("Unknown field %s for Gongstruct RectAnchoredPath", fieldName)
	}
	return
}

func (rectanchoredpngimage *RectAnchoredPngImage) GongMarshallField(stage *Stage, fieldName string) (res string) {
	ident := rectanchoredpngimage.GongGetIdentifier(stage)
	_ = ident
	switch fieldName {
	case "Name":
		res = __gong__marshallString(ident, "Name", rectanchoredpngimage.Name)
	case "X":
		res = __gong__marshallFloat(ident, "X", rectanchoredpngimage.X)
	case "Y":
		res = __gong__marshallFloat(ident, "Y", rectanchoredpngimage.Y)
	case "Width":
		res = __gong__marshallFloat(ident, "Width", rectanchoredpngimage.Width)
	case "Height":
		res = __gong__marshallFloat(ident, "Height", rectanchoredpngimage.Height)
	case "RX":
		res = __gong__marshallFloat(ident, "RX", rectanchoredpngimage.RX)
	case "X_Offset":
		res = __gong__marshallFloat(ident, "X_Offset", rectanchoredpngimage.X_Offset)
	case "Y_Offset":
		res = __gong__marshallFloat(ident, "Y_Offset", rectanchoredpngimage.Y_Offset)
	case "RectAnchorType":
		res = __gong__marshallEnumString(ident, "RectAnchorType", rectanchoredpngimage.RectAnchorType.ToCodeString())
	case "Base64Content":
		res = __gong__marshallString(ident, "Base64Content", rectanchoredpngimage.Base64Content)

	default:
		log.Panicf("Unknown field %s for Gongstruct RectAnchoredPngImage", fieldName)
	}
	return
}

func (rectanchoredrect *RectAnchoredRect) GongMarshallField(stage *Stage, fieldName string) (res string) {
	ident := rectanchoredrect.GongGetIdentifier(stage)
	_ = ident
	switch fieldName {
	case "Name":
		res = __gong__marshallString(ident, "Name", rectanchoredrect.Name)
	case "X":
		res = __gong__marshallFloat(ident, "X", rectanchoredrect.X)
	case "Y":
		res = __gong__marshallFloat(ident, "Y", rectanchoredrect.Y)
	case "Width":
		res = __gong__marshallFloat(ident, "Width", rectanchoredrect.Width)
	case "Height":
		res = __gong__marshallFloat(ident, "Height", rectanchoredrect.Height)
	case "RX":
		res = __gong__marshallFloat(ident, "RX", rectanchoredrect.RX)
	case "X_Offset":
		res = __gong__marshallFloat(ident, "X_Offset", rectanchoredrect.X_Offset)
	case "Y_Offset":
		res = __gong__marshallFloat(ident, "Y_Offset", rectanchoredrect.Y_Offset)
	case "RectAnchorType":
		res = __gong__marshallEnumString(ident, "RectAnchorType", rectanchoredrect.RectAnchorType.ToCodeString())
	case "WidthFollowRect":
		res = __gong__marshallBool(ident, "WidthFollowRect", rectanchoredrect.WidthFollowRect)
	case "HeightFollowRect":
		res = __gong__marshallBool(ident, "HeightFollowRect", rectanchoredrect.HeightFollowRect)
	case "HasToolTip":
		res = __gong__marshallBool(ident, "HasToolTip", rectanchoredrect.HasToolTip)
	case "ToolTipText":
		res = __gong__marshallString(ident, "ToolTipText", rectanchoredrect.ToolTipText)
	case "Color":
		res = __gong__marshallString(ident, "Color", rectanchoredrect.Color)
	case "FillOpacity":
		res = __gong__marshallFloat(ident, "FillOpacity", rectanchoredrect.FillOpacity)
	case "Stroke":
		res = __gong__marshallString(ident, "Stroke", rectanchoredrect.Stroke)
	case "StrokeOpacity":
		res = __gong__marshallFloat(ident, "StrokeOpacity", rectanchoredrect.StrokeOpacity)
	case "StrokeWidth":
		res = __gong__marshallFloat(ident, "StrokeWidth", rectanchoredrect.StrokeWidth)
	case "StrokeDashArray":
		res = __gong__marshallString(ident, "StrokeDashArray", rectanchoredrect.StrokeDashArray)
	case "StrokeDashArrayWhenSelected":
		res = __gong__marshallString(ident, "StrokeDashArrayWhenSelected", rectanchoredrect.StrokeDashArrayWhenSelected)
	case "Transform":
		res = __gong__marshallString(ident, "Transform", rectanchoredrect.Transform)

	default:
		log.Panicf("Unknown field %s for Gongstruct RectAnchoredRect", fieldName)
	}
	return
}

func (rectanchoredtext *RectAnchoredText) GongMarshallField(stage *Stage, fieldName string) (res string) {
	ident := rectanchoredtext.GongGetIdentifier(stage)
	_ = ident
	switch fieldName {
	case "Name":
		res = __gong__marshallString(ident, "Name", rectanchoredtext.Name)
	case "Content":
		res = __gong__marshallString(ident, "Content", rectanchoredtext.Content)
	case "FontWeight":
		res = __gong__marshallString(ident, "FontWeight", rectanchoredtext.FontWeight)
	case "FontSize":
		res = __gong__marshallString(ident, "FontSize", rectanchoredtext.FontSize)
	case "FontStyle":
		res = __gong__marshallString(ident, "FontStyle", rectanchoredtext.FontStyle)
	case "LetterSpacing":
		res = __gong__marshallString(ident, "LetterSpacing", rectanchoredtext.LetterSpacing)
	case "FontFamily":
		res = __gong__marshallString(ident, "FontFamily", rectanchoredtext.FontFamily)
	case "WhiteSpace":
		res = __gong__marshallEnumString(ident, "WhiteSpace", rectanchoredtext.WhiteSpace.ToCodeString())
	case "X_Offset":
		res = __gong__marshallFloat(ident, "X_Offset", rectanchoredtext.X_Offset)
	case "Y_Offset":
		res = __gong__marshallFloat(ident, "Y_Offset", rectanchoredtext.Y_Offset)
	case "RectAnchorType":
		res = __gong__marshallEnumString(ident, "RectAnchorType", rectanchoredtext.RectAnchorType.ToCodeString())
	case "TextAnchorType":
		res = __gong__marshallEnumString(ident, "TextAnchorType", rectanchoredtext.TextAnchorType.ToCodeString())
	case "DominantBaseline":
		res = __gong__marshallEnumString(ident, "DominantBaseline", rectanchoredtext.DominantBaseline.ToCodeString())
	case "WritingMode":
		res = __gong__marshallEnumString(ident, "WritingMode", rectanchoredtext.WritingMode.ToCodeString())
	case "Color":
		res = __gong__marshallString(ident, "Color", rectanchoredtext.Color)
	case "FillOpacity":
		res = __gong__marshallFloat(ident, "FillOpacity", rectanchoredtext.FillOpacity)
	case "Stroke":
		res = __gong__marshallString(ident, "Stroke", rectanchoredtext.Stroke)
	case "StrokeOpacity":
		res = __gong__marshallFloat(ident, "StrokeOpacity", rectanchoredtext.StrokeOpacity)
	case "StrokeWidth":
		res = __gong__marshallFloat(ident, "StrokeWidth", rectanchoredtext.StrokeWidth)
	case "StrokeDashArray":
		res = __gong__marshallString(ident, "StrokeDashArray", rectanchoredtext.StrokeDashArray)
	case "StrokeDashArrayWhenSelected":
		res = __gong__marshallString(ident, "StrokeDashArrayWhenSelected", rectanchoredtext.StrokeDashArrayWhenSelected)
	case "Transform":
		res = __gong__marshallString(ident, "Transform", rectanchoredtext.Transform)
	case "URLPath":
		res = __gong__marshallString(ident, "URLPath", rectanchoredtext.URLPath)
	case "URLTarget":
		res = __gong__marshallEnumString(ident, "URLTarget", rectanchoredtext.URLTarget.ToCodeString())

	case "Animates":
		var sb strings.Builder
		for _, _animate := range rectanchoredtext.Animates {
			sb.WriteString(__gong__marshallSliceOfPointers(ident, "Animates", _animate.GongGetIdentifier(stage)))
		}
		res = sb.String()
	default:
		log.Panicf("Unknown field %s for Gongstruct RectAnchoredText", fieldName)
	}
	return
}

func (rectlinklink *RectLinkLink) GongMarshallField(stage *Stage, fieldName string) (res string) {
	ident := rectlinklink.GongGetIdentifier(stage)
	_ = ident
	switch fieldName {
	case "Name":
		res = __gong__marshallString(ident, "Name", rectlinklink.Name)
	case "TargetAnchorPosition":
		res = __gong__marshallFloat(ident, "TargetAnchorPosition", rectlinklink.TargetAnchorPosition)
	case "Color":
		res = __gong__marshallString(ident, "Color", rectlinklink.Color)
	case "FillOpacity":
		res = __gong__marshallFloat(ident, "FillOpacity", rectlinklink.FillOpacity)
	case "Stroke":
		res = __gong__marshallString(ident, "Stroke", rectlinklink.Stroke)
	case "StrokeOpacity":
		res = __gong__marshallFloat(ident, "StrokeOpacity", rectlinklink.StrokeOpacity)
	case "StrokeWidth":
		res = __gong__marshallFloat(ident, "StrokeWidth", rectlinklink.StrokeWidth)
	case "StrokeDashArray":
		res = __gong__marshallString(ident, "StrokeDashArray", rectlinklink.StrokeDashArray)
	case "StrokeDashArrayWhenSelected":
		res = __gong__marshallString(ident, "StrokeDashArrayWhenSelected", rectlinklink.StrokeDashArrayWhenSelected)
	case "Transform":
		res = __gong__marshallString(ident, "Transform", rectlinklink.Transform)

	case "Start":
		if rectlinklink.Start != nil {
			res = __gong__marshallPointer(ident, "Start", rectlinklink.Start.GongGetIdentifier(stage))
		} else {
			res = __gong__marshallPointer(ident, "Start", "nil")
		}
	case "End":
		if rectlinklink.End != nil {
			res = __gong__marshallPointer(ident, "End", rectlinklink.End.GongGetIdentifier(stage))
		} else {
			res = __gong__marshallPointer(ident, "End", "nil")
		}
	default:
		log.Panicf("Unknown field %s for Gongstruct RectLinkLink", fieldName)
	}
	return
}

func (svg *SVG) GongMarshallField(stage *Stage, fieldName string) (res string) {
	ident := svg.GongGetIdentifier(stage)
	_ = ident
	switch fieldName {
	case "Name":
		res = __gong__marshallString(ident, "Name", svg.Name)
	case "DrawingState":
		res = __gong__marshallEnumString(ident, "DrawingState", svg.DrawingState.ToCodeString())
	case "IsEditable":
		res = __gong__marshallBool(ident, "IsEditable", svg.IsEditable)
	case "IsSVGFrontEndFileGenerated":
		res = __gong__marshallBool(ident, "IsSVGFrontEndFileGenerated", svg.IsSVGFrontEndFileGenerated)
	case "IsSVGBackEndFileGenerated":
		res = __gong__marshallBool(ident, "IsSVGBackEndFileGenerated", svg.IsSVGBackEndFileGenerated)
	case "DefaultDirectoryForGeneratedImages":
		res = __gong__marshallString(ident, "DefaultDirectoryForGeneratedImages", svg.DefaultDirectoryForGeneratedImages)
	case "IsControlBannerHidden":
		res = __gong__marshallBool(ident, "IsControlBannerHidden", svg.IsControlBannerHidden)
	case "PanX":
		res = __gong__marshallFloat(ident, "PanX", svg.PanX)
	case "PanY":
		res = __gong__marshallFloat(ident, "PanY", svg.PanY)
	case "Zoom":
		res = __gong__marshallFloat(ident, "Zoom", svg.Zoom)
	case "OverrideWidth":
		res = __gong__marshallBool(ident, "OverrideWidth", svg.OverrideWidth)
	case "OverriddenWidth":
		res = __gong__marshallFloat(ident, "OverriddenWidth", svg.OverriddenWidth)
	case "OverrideHeight":
		res = __gong__marshallBool(ident, "OverrideHeight", svg.OverrideHeight)
	case "OverriddenHeight":
		res = __gong__marshallFloat(ident, "OverriddenHeight", svg.OverriddenHeight)

	case "Layers":
		var sb strings.Builder
		for _, _layer := range svg.Layers {
			sb.WriteString(__gong__marshallSliceOfPointers(ident, "Layers", _layer.GongGetIdentifier(stage)))
		}
		res = sb.String()
	case "StartRect":
		if svg.StartRect != nil {
			res = __gong__marshallPointer(ident, "StartRect", svg.StartRect.GongGetIdentifier(stage))
		} else {
			res = __gong__marshallPointer(ident, "StartRect", "nil")
		}
	case "EndRect":
		if svg.EndRect != nil {
			res = __gong__marshallPointer(ident, "EndRect", svg.EndRect.GongGetIdentifier(stage))
		} else {
			res = __gong__marshallPointer(ident, "EndRect", "nil")
		}
	default:
		log.Panicf("Unknown field %s for Gongstruct SVG", fieldName)
	}
	return
}

func (svgtext *SvgText) GongMarshallField(stage *Stage, fieldName string) (res string) {
	ident := svgtext.GongGetIdentifier(stage)
	_ = ident
	switch fieldName {
	case "Name":
		res = __gong__marshallString(ident, "Name", svgtext.Name)
	case "Text":
		res = __gong__marshallString(ident, "Text", svgtext.Text)

	default:
		log.Panicf("Unknown field %s for Gongstruct SvgText", fieldName)
	}
	return
}

func (text *Text) GongMarshallField(stage *Stage, fieldName string) (res string) {
	ident := text.GongGetIdentifier(stage)
	_ = ident
	switch fieldName {
	case "Name":
		res = __gong__marshallString(ident, "Name", text.Name)
	case "X":
		res = __gong__marshallFloat(ident, "X", text.X)
	case "Y":
		res = __gong__marshallFloat(ident, "Y", text.Y)
	case "Content":
		res = __gong__marshallString(ident, "Content", text.Content)
	case "Color":
		res = __gong__marshallString(ident, "Color", text.Color)
	case "FillOpacity":
		res = __gong__marshallFloat(ident, "FillOpacity", text.FillOpacity)
	case "Stroke":
		res = __gong__marshallString(ident, "Stroke", text.Stroke)
	case "StrokeOpacity":
		res = __gong__marshallFloat(ident, "StrokeOpacity", text.StrokeOpacity)
	case "StrokeWidth":
		res = __gong__marshallFloat(ident, "StrokeWidth", text.StrokeWidth)
	case "StrokeDashArray":
		res = __gong__marshallString(ident, "StrokeDashArray", text.StrokeDashArray)
	case "StrokeDashArrayWhenSelected":
		res = __gong__marshallString(ident, "StrokeDashArrayWhenSelected", text.StrokeDashArrayWhenSelected)
	case "Transform":
		res = __gong__marshallString(ident, "Transform", text.Transform)
	case "FontWeight":
		res = __gong__marshallString(ident, "FontWeight", text.FontWeight)
	case "FontSize":
		res = __gong__marshallString(ident, "FontSize", text.FontSize)
	case "FontStyle":
		res = __gong__marshallString(ident, "FontStyle", text.FontStyle)
	case "LetterSpacing":
		res = __gong__marshallString(ident, "LetterSpacing", text.LetterSpacing)
	case "FontFamily":
		res = __gong__marshallString(ident, "FontFamily", text.FontFamily)
	case "WhiteSpace":
		res = __gong__marshallEnumString(ident, "WhiteSpace", text.WhiteSpace.ToCodeString())

	case "Animates":
		var sb strings.Builder
		for _, _animate := range text.Animates {
			sb.WriteString(__gong__marshallSliceOfPointers(ident, "Animates", _animate.GongGetIdentifier(stage)))
		}
		res = sb.String()
	default:
		log.Panicf("Unknown field %s for Gongstruct Text", fieldName)
	}
	return
}

// insertion point for marshall all fields methods
func (animate *Animate) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(animate.GongMarshallField(stage, "Name"))
		initializerStatements.WriteString(animate.GongMarshallField(stage, "AttributeName"))
		initializerStatements.WriteString(animate.GongMarshallField(stage, "Values"))
		initializerStatements.WriteString(animate.GongMarshallField(stage, "From"))
		initializerStatements.WriteString(animate.GongMarshallField(stage, "To"))
		initializerStatements.WriteString(animate.GongMarshallField(stage, "Dur"))
		initializerStatements.WriteString(animate.GongMarshallField(stage, "RepeatCount"))
	}
	initRes = initializerStatements.String()
	ptrRes = pointersInitializesStatements.String()
	return
}
func (circle *Circle) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(circle.GongMarshallField(stage, "Name"))
		initializerStatements.WriteString(circle.GongMarshallField(stage, "CX"))
		initializerStatements.WriteString(circle.GongMarshallField(stage, "CY"))
		initializerStatements.WriteString(circle.GongMarshallField(stage, "Radius"))
		initializerStatements.WriteString(circle.GongMarshallField(stage, "Color"))
		initializerStatements.WriteString(circle.GongMarshallField(stage, "FillOpacity"))
		initializerStatements.WriteString(circle.GongMarshallField(stage, "Stroke"))
		initializerStatements.WriteString(circle.GongMarshallField(stage, "StrokeOpacity"))
		initializerStatements.WriteString(circle.GongMarshallField(stage, "StrokeWidth"))
		initializerStatements.WriteString(circle.GongMarshallField(stage, "StrokeDashArray"))
		initializerStatements.WriteString(circle.GongMarshallField(stage, "StrokeDashArrayWhenSelected"))
		initializerStatements.WriteString(circle.GongMarshallField(stage, "Transform"))
		pointersInitializesStatements.WriteString(circle.GongMarshallField(stage, "Animations"))
	}
	initRes = initializerStatements.String()
	ptrRes = pointersInitializesStatements.String()
	return
}
func (condition *Condition) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(condition.GongMarshallField(stage, "Name"))
	}
	initRes = initializerStatements.String()
	ptrRes = pointersInitializesStatements.String()
	return
}
func (controlpoint *ControlPoint) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(controlpoint.GongMarshallField(stage, "Name"))
		initializerStatements.WriteString(controlpoint.GongMarshallField(stage, "X_Relative"))
		initializerStatements.WriteString(controlpoint.GongMarshallField(stage, "Y_Relative"))
		pointersInitializesStatements.WriteString(controlpoint.GongMarshallField(stage, "ClosestRect"))
	}
	initRes = initializerStatements.String()
	ptrRes = pointersInitializesStatements.String()
	return
}
func (ellipse *Ellipse) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(ellipse.GongMarshallField(stage, "Name"))
		initializerStatements.WriteString(ellipse.GongMarshallField(stage, "CX"))
		initializerStatements.WriteString(ellipse.GongMarshallField(stage, "CY"))
		initializerStatements.WriteString(ellipse.GongMarshallField(stage, "RX"))
		initializerStatements.WriteString(ellipse.GongMarshallField(stage, "RY"))
		initializerStatements.WriteString(ellipse.GongMarshallField(stage, "Color"))
		initializerStatements.WriteString(ellipse.GongMarshallField(stage, "FillOpacity"))
		initializerStatements.WriteString(ellipse.GongMarshallField(stage, "Stroke"))
		initializerStatements.WriteString(ellipse.GongMarshallField(stage, "StrokeOpacity"))
		initializerStatements.WriteString(ellipse.GongMarshallField(stage, "StrokeWidth"))
		initializerStatements.WriteString(ellipse.GongMarshallField(stage, "StrokeDashArray"))
		initializerStatements.WriteString(ellipse.GongMarshallField(stage, "StrokeDashArrayWhenSelected"))
		initializerStatements.WriteString(ellipse.GongMarshallField(stage, "Transform"))
		pointersInitializesStatements.WriteString(ellipse.GongMarshallField(stage, "Animates"))
	}
	initRes = initializerStatements.String()
	ptrRes = pointersInitializesStatements.String()
	return
}
func (filetodownload *FileToDownload) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(filetodownload.GongMarshallField(stage, "Name"))
		initializerStatements.WriteString(filetodownload.GongMarshallField(stage, "Base64EncodedContent"))
	}
	initRes = initializerStatements.String()
	ptrRes = pointersInitializesStatements.String()
	return
}
func (layer *Layer) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(layer.GongMarshallField(stage, "Name"))
		pointersInitializesStatements.WriteString(layer.GongMarshallField(stage, "Rects"))
		pointersInitializesStatements.WriteString(layer.GongMarshallField(stage, "Texts"))
		pointersInitializesStatements.WriteString(layer.GongMarshallField(stage, "Circles"))
		pointersInitializesStatements.WriteString(layer.GongMarshallField(stage, "Lines"))
		pointersInitializesStatements.WriteString(layer.GongMarshallField(stage, "Ellipses"))
		pointersInitializesStatements.WriteString(layer.GongMarshallField(stage, "Polylines"))
		pointersInitializesStatements.WriteString(layer.GongMarshallField(stage, "Polygones"))
		pointersInitializesStatements.WriteString(layer.GongMarshallField(stage, "Paths"))
		pointersInitializesStatements.WriteString(layer.GongMarshallField(stage, "Links"))
		pointersInitializesStatements.WriteString(layer.GongMarshallField(stage, "RectLinkLinks"))
	}
	initRes = initializerStatements.String()
	ptrRes = pointersInitializesStatements.String()
	return
}
func (line *Line) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(line.GongMarshallField(stage, "Name"))
		initializerStatements.WriteString(line.GongMarshallField(stage, "X1"))
		initializerStatements.WriteString(line.GongMarshallField(stage, "Y1"))
		initializerStatements.WriteString(line.GongMarshallField(stage, "X2"))
		initializerStatements.WriteString(line.GongMarshallField(stage, "Y2"))
		initializerStatements.WriteString(line.GongMarshallField(stage, "Color"))
		initializerStatements.WriteString(line.GongMarshallField(stage, "FillOpacity"))
		initializerStatements.WriteString(line.GongMarshallField(stage, "Stroke"))
		initializerStatements.WriteString(line.GongMarshallField(stage, "StrokeOpacity"))
		initializerStatements.WriteString(line.GongMarshallField(stage, "StrokeWidth"))
		initializerStatements.WriteString(line.GongMarshallField(stage, "StrokeDashArray"))
		initializerStatements.WriteString(line.GongMarshallField(stage, "StrokeDashArrayWhenSelected"))
		initializerStatements.WriteString(line.GongMarshallField(stage, "Transform"))
		pointersInitializesStatements.WriteString(line.GongMarshallField(stage, "Animates"))
		initializerStatements.WriteString(line.GongMarshallField(stage, "MouseClickX"))
		initializerStatements.WriteString(line.GongMarshallField(stage, "MouseClickY"))
	}
	initRes = initializerStatements.String()
	ptrRes = pointersInitializesStatements.String()
	return
}
func (link *Link) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(link.GongMarshallField(stage, "Name"))
		initializerStatements.WriteString(link.GongMarshallField(stage, "Type"))
		initializerStatements.WriteString(link.GongMarshallField(stage, "IsBezierCurve"))
		pointersInitializesStatements.WriteString(link.GongMarshallField(stage, "Start"))
		initializerStatements.WriteString(link.GongMarshallField(stage, "StartAnchorType"))
		pointersInitializesStatements.WriteString(link.GongMarshallField(stage, "End"))
		initializerStatements.WriteString(link.GongMarshallField(stage, "EndAnchorType"))
		initializerStatements.WriteString(link.GongMarshallField(stage, "StartOrientation"))
		initializerStatements.WriteString(link.GongMarshallField(stage, "StartRatio"))
		initializerStatements.WriteString(link.GongMarshallField(stage, "EndOrientation"))
		initializerStatements.WriteString(link.GongMarshallField(stage, "EndRatio"))
		initializerStatements.WriteString(link.GongMarshallField(stage, "CornerOffsetRatio"))
		initializerStatements.WriteString(link.GongMarshallField(stage, "CornerRadius"))
		initializerStatements.WriteString(link.GongMarshallField(stage, "HasEndArrow"))
		initializerStatements.WriteString(link.GongMarshallField(stage, "EndArrowSize"))
		initializerStatements.WriteString(link.GongMarshallField(stage, "EndArrowOffset"))
		initializerStatements.WriteString(link.GongMarshallField(stage, "HasStartArrow"))
		initializerStatements.WriteString(link.GongMarshallField(stage, "StartArrowSize"))
		initializerStatements.WriteString(link.GongMarshallField(stage, "StartArrowOffset"))
		pointersInitializesStatements.WriteString(link.GongMarshallField(stage, "TextAtArrowStart"))
		pointersInitializesStatements.WriteString(link.GongMarshallField(stage, "TextAtArrowEnd"))
		pointersInitializesStatements.WriteString(link.GongMarshallField(stage, "TextAtCorner"))
		pointersInitializesStatements.WriteString(link.GongMarshallField(stage, "PathAtArrowStart"))
		pointersInitializesStatements.WriteString(link.GongMarshallField(stage, "PathAtArrowEnd"))
		pointersInitializesStatements.WriteString(link.GongMarshallField(stage, "PathAtCorner"))
		pointersInitializesStatements.WriteString(link.GongMarshallField(stage, "ControlPoints"))
		initializerStatements.WriteString(link.GongMarshallField(stage, "Color"))
		initializerStatements.WriteString(link.GongMarshallField(stage, "FillOpacity"))
		initializerStatements.WriteString(link.GongMarshallField(stage, "Stroke"))
		initializerStatements.WriteString(link.GongMarshallField(stage, "StrokeOpacity"))
		initializerStatements.WriteString(link.GongMarshallField(stage, "StrokeWidth"))
		initializerStatements.WriteString(link.GongMarshallField(stage, "StrokeDashArray"))
		initializerStatements.WriteString(link.GongMarshallField(stage, "StrokeDashArrayWhenSelected"))
		initializerStatements.WriteString(link.GongMarshallField(stage, "Transform"))
		initializerStatements.WriteString(link.GongMarshallField(stage, "MouseX"))
		initializerStatements.WriteString(link.GongMarshallField(stage, "MouseY"))
		initializerStatements.WriteString(link.GongMarshallField(stage, "MouseEventKey"))
	}
	initRes = initializerStatements.String()
	ptrRes = pointersInitializesStatements.String()
	return
}
func (linkanchoredpath *LinkAnchoredPath) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(linkanchoredpath.GongMarshallField(stage, "Name"))
		initializerStatements.WriteString(linkanchoredpath.GongMarshallField(stage, "Definition"))
		initializerStatements.WriteString(linkanchoredpath.GongMarshallField(stage, "X_Offset"))
		initializerStatements.WriteString(linkanchoredpath.GongMarshallField(stage, "Y_Offset"))
		initializerStatements.WriteString(linkanchoredpath.GongMarshallField(stage, "ScalePropotionnally"))
		initializerStatements.WriteString(linkanchoredpath.GongMarshallField(stage, "AppliedScaling"))
		initializerStatements.WriteString(linkanchoredpath.GongMarshallField(stage, "Color"))
		initializerStatements.WriteString(linkanchoredpath.GongMarshallField(stage, "FillOpacity"))
		initializerStatements.WriteString(linkanchoredpath.GongMarshallField(stage, "Stroke"))
		initializerStatements.WriteString(linkanchoredpath.GongMarshallField(stage, "StrokeOpacity"))
		initializerStatements.WriteString(linkanchoredpath.GongMarshallField(stage, "StrokeWidth"))
		initializerStatements.WriteString(linkanchoredpath.GongMarshallField(stage, "StrokeDashArray"))
		initializerStatements.WriteString(linkanchoredpath.GongMarshallField(stage, "StrokeDashArrayWhenSelected"))
		initializerStatements.WriteString(linkanchoredpath.GongMarshallField(stage, "Transform"))
	}
	initRes = initializerStatements.String()
	ptrRes = pointersInitializesStatements.String()
	return
}
func (linkanchoredtext *LinkAnchoredText) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(linkanchoredtext.GongMarshallField(stage, "Name"))
		initializerStatements.WriteString(linkanchoredtext.GongMarshallField(stage, "Content"))
		initializerStatements.WriteString(linkanchoredtext.GongMarshallField(stage, "AutomaticLayout"))
		initializerStatements.WriteString(linkanchoredtext.GongMarshallField(stage, "LinkAnchorType"))
		initializerStatements.WriteString(linkanchoredtext.GongMarshallField(stage, "X_Offset"))
		initializerStatements.WriteString(linkanchoredtext.GongMarshallField(stage, "Y_Offset"))
		initializerStatements.WriteString(linkanchoredtext.GongMarshallField(stage, "FontWeight"))
		initializerStatements.WriteString(linkanchoredtext.GongMarshallField(stage, "FontSize"))
		initializerStatements.WriteString(linkanchoredtext.GongMarshallField(stage, "FontStyle"))
		initializerStatements.WriteString(linkanchoredtext.GongMarshallField(stage, "LetterSpacing"))
		initializerStatements.WriteString(linkanchoredtext.GongMarshallField(stage, "FontFamily"))
		initializerStatements.WriteString(linkanchoredtext.GongMarshallField(stage, "WhiteSpace"))
		initializerStatements.WriteString(linkanchoredtext.GongMarshallField(stage, "Color"))
		initializerStatements.WriteString(linkanchoredtext.GongMarshallField(stage, "FillOpacity"))
		initializerStatements.WriteString(linkanchoredtext.GongMarshallField(stage, "Stroke"))
		initializerStatements.WriteString(linkanchoredtext.GongMarshallField(stage, "StrokeOpacity"))
		initializerStatements.WriteString(linkanchoredtext.GongMarshallField(stage, "StrokeWidth"))
		initializerStatements.WriteString(linkanchoredtext.GongMarshallField(stage, "StrokeDashArray"))
		initializerStatements.WriteString(linkanchoredtext.GongMarshallField(stage, "StrokeDashArrayWhenSelected"))
		initializerStatements.WriteString(linkanchoredtext.GongMarshallField(stage, "Transform"))
		pointersInitializesStatements.WriteString(linkanchoredtext.GongMarshallField(stage, "Animates"))
	}
	initRes = initializerStatements.String()
	ptrRes = pointersInitializesStatements.String()
	return
}
func (path *Path) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(path.GongMarshallField(stage, "Name"))
		initializerStatements.WriteString(path.GongMarshallField(stage, "Definition"))
		initializerStatements.WriteString(path.GongMarshallField(stage, "Color"))
		initializerStatements.WriteString(path.GongMarshallField(stage, "FillOpacity"))
		initializerStatements.WriteString(path.GongMarshallField(stage, "Stroke"))
		initializerStatements.WriteString(path.GongMarshallField(stage, "StrokeOpacity"))
		initializerStatements.WriteString(path.GongMarshallField(stage, "StrokeWidth"))
		initializerStatements.WriteString(path.GongMarshallField(stage, "StrokeDashArray"))
		initializerStatements.WriteString(path.GongMarshallField(stage, "StrokeDashArrayWhenSelected"))
		initializerStatements.WriteString(path.GongMarshallField(stage, "Transform"))
		pointersInitializesStatements.WriteString(path.GongMarshallField(stage, "Animates"))
	}
	initRes = initializerStatements.String()
	ptrRes = pointersInitializesStatements.String()
	return
}
func (point *Point) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(point.GongMarshallField(stage, "Name"))
		initializerStatements.WriteString(point.GongMarshallField(stage, "X"))
		initializerStatements.WriteString(point.GongMarshallField(stage, "Y"))
	}
	initRes = initializerStatements.String()
	ptrRes = pointersInitializesStatements.String()
	return
}
func (polygone *Polygone) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(polygone.GongMarshallField(stage, "Name"))
		initializerStatements.WriteString(polygone.GongMarshallField(stage, "Points"))
		initializerStatements.WriteString(polygone.GongMarshallField(stage, "Color"))
		initializerStatements.WriteString(polygone.GongMarshallField(stage, "FillOpacity"))
		initializerStatements.WriteString(polygone.GongMarshallField(stage, "Stroke"))
		initializerStatements.WriteString(polygone.GongMarshallField(stage, "StrokeOpacity"))
		initializerStatements.WriteString(polygone.GongMarshallField(stage, "StrokeWidth"))
		initializerStatements.WriteString(polygone.GongMarshallField(stage, "StrokeDashArray"))
		initializerStatements.WriteString(polygone.GongMarshallField(stage, "StrokeDashArrayWhenSelected"))
		initializerStatements.WriteString(polygone.GongMarshallField(stage, "Transform"))
		pointersInitializesStatements.WriteString(polygone.GongMarshallField(stage, "Animates"))
	}
	initRes = initializerStatements.String()
	ptrRes = pointersInitializesStatements.String()
	return
}
func (polyline *Polyline) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(polyline.GongMarshallField(stage, "Name"))
		initializerStatements.WriteString(polyline.GongMarshallField(stage, "Points"))
		initializerStatements.WriteString(polyline.GongMarshallField(stage, "Color"))
		initializerStatements.WriteString(polyline.GongMarshallField(stage, "FillOpacity"))
		initializerStatements.WriteString(polyline.GongMarshallField(stage, "Stroke"))
		initializerStatements.WriteString(polyline.GongMarshallField(stage, "StrokeOpacity"))
		initializerStatements.WriteString(polyline.GongMarshallField(stage, "StrokeWidth"))
		initializerStatements.WriteString(polyline.GongMarshallField(stage, "StrokeDashArray"))
		initializerStatements.WriteString(polyline.GongMarshallField(stage, "StrokeDashArrayWhenSelected"))
		initializerStatements.WriteString(polyline.GongMarshallField(stage, "Transform"))
		pointersInitializesStatements.WriteString(polyline.GongMarshallField(stage, "Animates"))
	}
	initRes = initializerStatements.String()
	ptrRes = pointersInitializesStatements.String()
	return
}
func (rect *Rect) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(rect.GongMarshallField(stage, "Name"))
		initializerStatements.WriteString(rect.GongMarshallField(stage, "X"))
		initializerStatements.WriteString(rect.GongMarshallField(stage, "Y"))
		initializerStatements.WriteString(rect.GongMarshallField(stage, "Width"))
		initializerStatements.WriteString(rect.GongMarshallField(stage, "Height"))
		initializerStatements.WriteString(rect.GongMarshallField(stage, "RX"))
		pointersInitializesStatements.WriteString(rect.GongMarshallField(stage, "Peers"))
		pointersInitializesStatements.WriteString(rect.GongMarshallField(stage, "EnclosingRect"))
		pointersInitializesStatements.WriteString(rect.GongMarshallField(stage, "Obstacles"))
		pointersInitializesStatements.WriteString(rect.GongMarshallField(stage, "AnchoredTo"))
		initializerStatements.WriteString(rect.GongMarshallField(stage, "Color"))
		initializerStatements.WriteString(rect.GongMarshallField(stage, "FillOpacity"))
		initializerStatements.WriteString(rect.GongMarshallField(stage, "Stroke"))
		initializerStatements.WriteString(rect.GongMarshallField(stage, "StrokeOpacity"))
		initializerStatements.WriteString(rect.GongMarshallField(stage, "StrokeWidth"))
		initializerStatements.WriteString(rect.GongMarshallField(stage, "StrokeDashArray"))
		initializerStatements.WriteString(rect.GongMarshallField(stage, "StrokeDashArrayWhenSelected"))
		initializerStatements.WriteString(rect.GongMarshallField(stage, "Transform"))
		pointersInitializesStatements.WriteString(rect.GongMarshallField(stage, "HoveringTrigger"))
		pointersInitializesStatements.WriteString(rect.GongMarshallField(stage, "DisplayConditions"))
		pointersInitializesStatements.WriteString(rect.GongMarshallField(stage, "Animations"))
		initializerStatements.WriteString(rect.GongMarshallField(stage, "IsSelectable"))
		initializerStatements.WriteString(rect.GongMarshallField(stage, "IsSelected"))
		initializerStatements.WriteString(rect.GongMarshallField(stage, "CanHaveLeftHandle"))
		initializerStatements.WriteString(rect.GongMarshallField(stage, "HasLeftHandle"))
		initializerStatements.WriteString(rect.GongMarshallField(stage, "CanHaveRightHandle"))
		initializerStatements.WriteString(rect.GongMarshallField(stage, "HasRightHandle"))
		initializerStatements.WriteString(rect.GongMarshallField(stage, "CanHaveTopHandle"))
		initializerStatements.WriteString(rect.GongMarshallField(stage, "HasTopHandle"))
		initializerStatements.WriteString(rect.GongMarshallField(stage, "IsScalingProportionally"))
		initializerStatements.WriteString(rect.GongMarshallField(stage, "CanHaveBottomHandle"))
		initializerStatements.WriteString(rect.GongMarshallField(stage, "HasBottomHandle"))
		initializerStatements.WriteString(rect.GongMarshallField(stage, "CanMoveHorizontaly"))
		initializerStatements.WriteString(rect.GongMarshallField(stage, "CanMoveVerticaly"))
		pointersInitializesStatements.WriteString(rect.GongMarshallField(stage, "RectAnchoredTexts"))
		pointersInitializesStatements.WriteString(rect.GongMarshallField(stage, "RectAnchoredRects"))
		pointersInitializesStatements.WriteString(rect.GongMarshallField(stage, "RectAnchoredPaths"))
		pointersInitializesStatements.WriteString(rect.GongMarshallField(stage, "RectAnchoredPngImages"))
		initializerStatements.WriteString(rect.GongMarshallField(stage, "ChangeColorWhenHovered"))
		initializerStatements.WriteString(rect.GongMarshallField(stage, "ColorWhenHovered"))
		initializerStatements.WriteString(rect.GongMarshallField(stage, "OriginalColor"))
		initializerStatements.WriteString(rect.GongMarshallField(stage, "FillOpacityWhenHovered"))
		initializerStatements.WriteString(rect.GongMarshallField(stage, "OriginalFillOpacity"))
		initializerStatements.WriteString(rect.GongMarshallField(stage, "HasToolTip"))
		initializerStatements.WriteString(rect.GongMarshallField(stage, "ToolTipText"))
		initializerStatements.WriteString(rect.GongMarshallField(stage, "ToolTipPosition"))
		initializerStatements.WriteString(rect.GongMarshallField(stage, "MouseX"))
		initializerStatements.WriteString(rect.GongMarshallField(stage, "MouseY"))
		initializerStatements.WriteString(rect.GongMarshallField(stage, "MouseEventKey"))
		initializerStatements.WriteString(rect.GongMarshallField(stage, "URLPath"))
		initializerStatements.WriteString(rect.GongMarshallField(stage, "URLTarget"))
	}
	initRes = initializerStatements.String()
	ptrRes = pointersInitializesStatements.String()
	return
}
func (rectanchoredpath *RectAnchoredPath) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(rectanchoredpath.GongMarshallField(stage, "Name"))
		initializerStatements.WriteString(rectanchoredpath.GongMarshallField(stage, "Definition"))
		initializerStatements.WriteString(rectanchoredpath.GongMarshallField(stage, "X_Offset"))
		initializerStatements.WriteString(rectanchoredpath.GongMarshallField(stage, "Y_Offset"))
		initializerStatements.WriteString(rectanchoredpath.GongMarshallField(stage, "RectAnchorType"))
		initializerStatements.WriteString(rectanchoredpath.GongMarshallField(stage, "ScalePropotionnally"))
		initializerStatements.WriteString(rectanchoredpath.GongMarshallField(stage, "AppliedScaling"))
		initializerStatements.WriteString(rectanchoredpath.GongMarshallField(stage, "Color"))
		initializerStatements.WriteString(rectanchoredpath.GongMarshallField(stage, "FillOpacity"))
		initializerStatements.WriteString(rectanchoredpath.GongMarshallField(stage, "Stroke"))
		initializerStatements.WriteString(rectanchoredpath.GongMarshallField(stage, "StrokeOpacity"))
		initializerStatements.WriteString(rectanchoredpath.GongMarshallField(stage, "StrokeWidth"))
		initializerStatements.WriteString(rectanchoredpath.GongMarshallField(stage, "StrokeDashArray"))
		initializerStatements.WriteString(rectanchoredpath.GongMarshallField(stage, "StrokeDashArrayWhenSelected"))
		initializerStatements.WriteString(rectanchoredpath.GongMarshallField(stage, "Transform"))
	}
	initRes = initializerStatements.String()
	ptrRes = pointersInitializesStatements.String()
	return
}
func (rectanchoredpngimage *RectAnchoredPngImage) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(rectanchoredpngimage.GongMarshallField(stage, "Name"))
		initializerStatements.WriteString(rectanchoredpngimage.GongMarshallField(stage, "X"))
		initializerStatements.WriteString(rectanchoredpngimage.GongMarshallField(stage, "Y"))
		initializerStatements.WriteString(rectanchoredpngimage.GongMarshallField(stage, "Width"))
		initializerStatements.WriteString(rectanchoredpngimage.GongMarshallField(stage, "Height"))
		initializerStatements.WriteString(rectanchoredpngimage.GongMarshallField(stage, "RX"))
		initializerStatements.WriteString(rectanchoredpngimage.GongMarshallField(stage, "X_Offset"))
		initializerStatements.WriteString(rectanchoredpngimage.GongMarshallField(stage, "Y_Offset"))
		initializerStatements.WriteString(rectanchoredpngimage.GongMarshallField(stage, "RectAnchorType"))
		initializerStatements.WriteString(rectanchoredpngimage.GongMarshallField(stage, "Base64Content"))
	}
	initRes = initializerStatements.String()
	ptrRes = pointersInitializesStatements.String()
	return
}
func (rectanchoredrect *RectAnchoredRect) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(rectanchoredrect.GongMarshallField(stage, "Name"))
		initializerStatements.WriteString(rectanchoredrect.GongMarshallField(stage, "X"))
		initializerStatements.WriteString(rectanchoredrect.GongMarshallField(stage, "Y"))
		initializerStatements.WriteString(rectanchoredrect.GongMarshallField(stage, "Width"))
		initializerStatements.WriteString(rectanchoredrect.GongMarshallField(stage, "Height"))
		initializerStatements.WriteString(rectanchoredrect.GongMarshallField(stage, "RX"))
		initializerStatements.WriteString(rectanchoredrect.GongMarshallField(stage, "X_Offset"))
		initializerStatements.WriteString(rectanchoredrect.GongMarshallField(stage, "Y_Offset"))
		initializerStatements.WriteString(rectanchoredrect.GongMarshallField(stage, "RectAnchorType"))
		initializerStatements.WriteString(rectanchoredrect.GongMarshallField(stage, "WidthFollowRect"))
		initializerStatements.WriteString(rectanchoredrect.GongMarshallField(stage, "HeightFollowRect"))
		initializerStatements.WriteString(rectanchoredrect.GongMarshallField(stage, "HasToolTip"))
		initializerStatements.WriteString(rectanchoredrect.GongMarshallField(stage, "ToolTipText"))
		initializerStatements.WriteString(rectanchoredrect.GongMarshallField(stage, "Color"))
		initializerStatements.WriteString(rectanchoredrect.GongMarshallField(stage, "FillOpacity"))
		initializerStatements.WriteString(rectanchoredrect.GongMarshallField(stage, "Stroke"))
		initializerStatements.WriteString(rectanchoredrect.GongMarshallField(stage, "StrokeOpacity"))
		initializerStatements.WriteString(rectanchoredrect.GongMarshallField(stage, "StrokeWidth"))
		initializerStatements.WriteString(rectanchoredrect.GongMarshallField(stage, "StrokeDashArray"))
		initializerStatements.WriteString(rectanchoredrect.GongMarshallField(stage, "StrokeDashArrayWhenSelected"))
		initializerStatements.WriteString(rectanchoredrect.GongMarshallField(stage, "Transform"))
	}
	initRes = initializerStatements.String()
	ptrRes = pointersInitializesStatements.String()
	return
}
func (rectanchoredtext *RectAnchoredText) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(rectanchoredtext.GongMarshallField(stage, "Name"))
		initializerStatements.WriteString(rectanchoredtext.GongMarshallField(stage, "Content"))
		initializerStatements.WriteString(rectanchoredtext.GongMarshallField(stage, "FontWeight"))
		initializerStatements.WriteString(rectanchoredtext.GongMarshallField(stage, "FontSize"))
		initializerStatements.WriteString(rectanchoredtext.GongMarshallField(stage, "FontStyle"))
		initializerStatements.WriteString(rectanchoredtext.GongMarshallField(stage, "LetterSpacing"))
		initializerStatements.WriteString(rectanchoredtext.GongMarshallField(stage, "FontFamily"))
		initializerStatements.WriteString(rectanchoredtext.GongMarshallField(stage, "WhiteSpace"))
		initializerStatements.WriteString(rectanchoredtext.GongMarshallField(stage, "X_Offset"))
		initializerStatements.WriteString(rectanchoredtext.GongMarshallField(stage, "Y_Offset"))
		initializerStatements.WriteString(rectanchoredtext.GongMarshallField(stage, "RectAnchorType"))
		initializerStatements.WriteString(rectanchoredtext.GongMarshallField(stage, "TextAnchorType"))
		initializerStatements.WriteString(rectanchoredtext.GongMarshallField(stage, "DominantBaseline"))
		initializerStatements.WriteString(rectanchoredtext.GongMarshallField(stage, "WritingMode"))
		initializerStatements.WriteString(rectanchoredtext.GongMarshallField(stage, "Color"))
		initializerStatements.WriteString(rectanchoredtext.GongMarshallField(stage, "FillOpacity"))
		initializerStatements.WriteString(rectanchoredtext.GongMarshallField(stage, "Stroke"))
		initializerStatements.WriteString(rectanchoredtext.GongMarshallField(stage, "StrokeOpacity"))
		initializerStatements.WriteString(rectanchoredtext.GongMarshallField(stage, "StrokeWidth"))
		initializerStatements.WriteString(rectanchoredtext.GongMarshallField(stage, "StrokeDashArray"))
		initializerStatements.WriteString(rectanchoredtext.GongMarshallField(stage, "StrokeDashArrayWhenSelected"))
		initializerStatements.WriteString(rectanchoredtext.GongMarshallField(stage, "Transform"))
		pointersInitializesStatements.WriteString(rectanchoredtext.GongMarshallField(stage, "Animates"))
		initializerStatements.WriteString(rectanchoredtext.GongMarshallField(stage, "URLPath"))
		initializerStatements.WriteString(rectanchoredtext.GongMarshallField(stage, "URLTarget"))
	}
	initRes = initializerStatements.String()
	ptrRes = pointersInitializesStatements.String()
	return
}
func (rectlinklink *RectLinkLink) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(rectlinklink.GongMarshallField(stage, "Name"))
		pointersInitializesStatements.WriteString(rectlinklink.GongMarshallField(stage, "Start"))
		pointersInitializesStatements.WriteString(rectlinklink.GongMarshallField(stage, "End"))
		initializerStatements.WriteString(rectlinklink.GongMarshallField(stage, "TargetAnchorPosition"))
		initializerStatements.WriteString(rectlinklink.GongMarshallField(stage, "Color"))
		initializerStatements.WriteString(rectlinklink.GongMarshallField(stage, "FillOpacity"))
		initializerStatements.WriteString(rectlinklink.GongMarshallField(stage, "Stroke"))
		initializerStatements.WriteString(rectlinklink.GongMarshallField(stage, "StrokeOpacity"))
		initializerStatements.WriteString(rectlinklink.GongMarshallField(stage, "StrokeWidth"))
		initializerStatements.WriteString(rectlinklink.GongMarshallField(stage, "StrokeDashArray"))
		initializerStatements.WriteString(rectlinklink.GongMarshallField(stage, "StrokeDashArrayWhenSelected"))
		initializerStatements.WriteString(rectlinklink.GongMarshallField(stage, "Transform"))
	}
	initRes = initializerStatements.String()
	ptrRes = pointersInitializesStatements.String()
	return
}
func (svg *SVG) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(svg.GongMarshallField(stage, "Name"))
		pointersInitializesStatements.WriteString(svg.GongMarshallField(stage, "Layers"))
		initializerStatements.WriteString(svg.GongMarshallField(stage, "DrawingState"))
		pointersInitializesStatements.WriteString(svg.GongMarshallField(stage, "StartRect"))
		pointersInitializesStatements.WriteString(svg.GongMarshallField(stage, "EndRect"))
		initializerStatements.WriteString(svg.GongMarshallField(stage, "IsEditable"))
		initializerStatements.WriteString(svg.GongMarshallField(stage, "IsSVGFrontEndFileGenerated"))
		initializerStatements.WriteString(svg.GongMarshallField(stage, "IsSVGBackEndFileGenerated"))
		initializerStatements.WriteString(svg.GongMarshallField(stage, "DefaultDirectoryForGeneratedImages"))
		initializerStatements.WriteString(svg.GongMarshallField(stage, "IsControlBannerHidden"))
		initializerStatements.WriteString(svg.GongMarshallField(stage, "PanX"))
		initializerStatements.WriteString(svg.GongMarshallField(stage, "PanY"))
		initializerStatements.WriteString(svg.GongMarshallField(stage, "Zoom"))
		initializerStatements.WriteString(svg.GongMarshallField(stage, "OverrideWidth"))
		initializerStatements.WriteString(svg.GongMarshallField(stage, "OverriddenWidth"))
		initializerStatements.WriteString(svg.GongMarshallField(stage, "OverrideHeight"))
		initializerStatements.WriteString(svg.GongMarshallField(stage, "OverriddenHeight"))
	}
	initRes = initializerStatements.String()
	ptrRes = pointersInitializesStatements.String()
	return
}
func (svgtext *SvgText) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(svgtext.GongMarshallField(stage, "Name"))
		initializerStatements.WriteString(svgtext.GongMarshallField(stage, "Text"))
	}
	initRes = initializerStatements.String()
	ptrRes = pointersInitializesStatements.String()
	return
}
func (text *Text) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(text.GongMarshallField(stage, "Name"))
		initializerStatements.WriteString(text.GongMarshallField(stage, "X"))
		initializerStatements.WriteString(text.GongMarshallField(stage, "Y"))
		initializerStatements.WriteString(text.GongMarshallField(stage, "Content"))
		initializerStatements.WriteString(text.GongMarshallField(stage, "Color"))
		initializerStatements.WriteString(text.GongMarshallField(stage, "FillOpacity"))
		initializerStatements.WriteString(text.GongMarshallField(stage, "Stroke"))
		initializerStatements.WriteString(text.GongMarshallField(stage, "StrokeOpacity"))
		initializerStatements.WriteString(text.GongMarshallField(stage, "StrokeWidth"))
		initializerStatements.WriteString(text.GongMarshallField(stage, "StrokeDashArray"))
		initializerStatements.WriteString(text.GongMarshallField(stage, "StrokeDashArrayWhenSelected"))
		initializerStatements.WriteString(text.GongMarshallField(stage, "Transform"))
		initializerStatements.WriteString(text.GongMarshallField(stage, "FontWeight"))
		initializerStatements.WriteString(text.GongMarshallField(stage, "FontSize"))
		initializerStatements.WriteString(text.GongMarshallField(stage, "FontStyle"))
		initializerStatements.WriteString(text.GongMarshallField(stage, "LetterSpacing"))
		initializerStatements.WriteString(text.GongMarshallField(stage, "FontFamily"))
		initializerStatements.WriteString(text.GongMarshallField(stage, "WhiteSpace"))
		pointersInitializesStatements.WriteString(text.GongMarshallField(stage, "Animates"))
	}
	initRes = initializerStatements.String()
	ptrRes = pointersInitializesStatements.String()
	return
}
