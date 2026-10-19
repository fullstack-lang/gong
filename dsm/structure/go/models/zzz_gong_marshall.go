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
	modelsPackageName := "github.com/fullstack-lang/gong/dsm/structure/go/models"
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
		for _, allocatedresourceshape := range __gong__sortStageSetInstances(stage.AllocatedResourceShapes, stage.AllocatedResourceShape_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			allocatedresourceshapeIdent := allocatedresourceshape.GongGetIdentifier(stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&models.AllocatedResourceShape{Name: %s}).Stage(stage)", allocatedresourceshapeIdent, __gong__toRawStringLiteral(allocatedresourceshape.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", allocatedresourceshapeIdent, __gong__toRawStringLiteral(allocatedresourceshape.Name)))
			if allocatedresourceshape.Part != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := allocatedresourceshape.Part.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Part = %s", allocatedresourceshapeIdent, targetIdent))
			}
			if allocatedresourceshape.Resource != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := allocatedresourceshape.Resource.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Resource = %s", allocatedresourceshapeIdent, targetIdent))
			}
		}
	}
	if stage != nil {
		for _, allocatedsystemshape := range __gong__sortStageSetInstances(stage.AllocatedSystemShapes, stage.AllocatedSystemShape_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			allocatedsystemshapeIdent := allocatedsystemshape.GongGetIdentifier(stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&models.AllocatedSystemShape{Name: %s}).Stage(stage)", allocatedsystemshapeIdent, __gong__toRawStringLiteral(allocatedsystemshape.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", allocatedsystemshapeIdent, __gong__toRawStringLiteral(allocatedsystemshape.Name)))
			if allocatedsystemshape.Part != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := allocatedsystemshape.Part.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Part = %s", allocatedsystemshapeIdent, targetIdent))
			}
			if allocatedsystemshape.System != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := allocatedsystemshape.System.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.System = %s", allocatedsystemshapeIdent, targetIdent))
			}
		}
	}
	if stage != nil {
		for _, controlflow := range __gong__sortStageSetInstances(stage.ControlFlows, stage.ControlFlow_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			controlflowIdent := controlflow.GongGetIdentifier(stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&models.ControlFlow{Name: %s}).Stage(stage)", controlflowIdent, __gong__toRawStringLiteral(controlflow.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", controlflowIdent, __gong__toRawStringLiteral(controlflow.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Description = %s", controlflowIdent, __gong__toRawStringLiteral(controlflow.Description)))
			values.WriteString(fmt.Sprintf("\n\t%s.ComputedPrefix = %s", controlflowIdent, __gong__toRawStringLiteral(controlflow.ComputedPrefix)))
			values.WriteString(fmt.Sprintf("\n\t%s.IsExpanded = %t", controlflowIdent, controlflow.IsExpanded))
			if controlflow.Start != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := controlflow.Start.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Start = %s", controlflowIdent, targetIdent))
			}
			if controlflow.End != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := controlflow.End.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.End = %s", controlflowIdent, targetIdent))
			}
		}
	}
	if stage != nil {
		for _, controlflowshape := range __gong__sortStageSetInstances(stage.ControlFlowShapes, stage.ControlFlowShape_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			controlflowshapeIdent := controlflowshape.GongGetIdentifier(stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&models.ControlFlowShape{Name: %s}).Stage(stage)", controlflowshapeIdent, __gong__toRawStringLiteral(controlflowshape.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", controlflowshapeIdent, __gong__toRawStringLiteral(controlflowshape.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.StartRatio = %f", controlflowshapeIdent, controlflowshape.StartRatio))
			values.WriteString(fmt.Sprintf("\n\t%s.EndRatio = %f", controlflowshapeIdent, controlflowshape.EndRatio))
			if controlflowshape.StartOrientation.ToCodeString() != "" {
				values.WriteString(fmt.Sprintf("\n\t%s.StartOrientation = models.%s", controlflowshapeIdent, controlflowshape.StartOrientation.ToCodeString()))
			} else {
				values.WriteString(fmt.Sprintf("\n\t%s.StartOrientation = \"\"", controlflowshapeIdent))
			}
			if controlflowshape.EndOrientation.ToCodeString() != "" {
				values.WriteString(fmt.Sprintf("\n\t%s.EndOrientation = models.%s", controlflowshapeIdent, controlflowshape.EndOrientation.ToCodeString()))
			} else {
				values.WriteString(fmt.Sprintf("\n\t%s.EndOrientation = \"\"", controlflowshapeIdent))
			}
			values.WriteString(fmt.Sprintf("\n\t%s.CornerOffsetRatio = %f", controlflowshapeIdent, controlflowshape.CornerOffsetRatio))
			values.WriteString(fmt.Sprintf("\n\t%s.IsHidden = %t", controlflowshapeIdent, controlflowshape.IsHidden))
			values.WriteString(fmt.Sprintf("\n\t%s.IsPresentationOverridden = %t", controlflowshapeIdent, controlflowshape.IsPresentationOverridden))
			values.WriteString(fmt.Sprintf("\n\t%s.Stroke = %s", controlflowshapeIdent, __gong__toRawStringLiteral(controlflowshape.Stroke)))
			values.WriteString(fmt.Sprintf("\n\t%s.StrokeWidth = %f", controlflowshapeIdent, controlflowshape.StrokeWidth))
			values.WriteString(fmt.Sprintf("\n\t%s.StrokeOpacity = %f", controlflowshapeIdent, controlflowshape.StrokeOpacity))
			values.WriteString(fmt.Sprintf("\n\t%s.StrokeDashArray = %s", controlflowshapeIdent, __gong__toRawStringLiteral(controlflowshape.StrokeDashArray)))
			values.WriteString(fmt.Sprintf("\n\t%s.FontSize = %s", controlflowshapeIdent, __gong__toRawStringLiteral(controlflowshape.FontSize)))
			values.WriteString(fmt.Sprintf("\n\t%s.FontColor = %s", controlflowshapeIdent, __gong__toRawStringLiteral(controlflowshape.FontColor)))
			if controlflowshape.ControlFlow != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := controlflowshape.ControlFlow.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.ControlFlow = %s", controlflowshapeIdent, targetIdent))
			}
		}
	}
	if stage != nil {
		for _, data := range __gong__sortStageSetInstances(stage.Datas, stage.Data_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			dataIdent := data.GongGetIdentifier(stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&models.Data{Name: %s}).Stage(stage)", dataIdent, __gong__toRawStringLiteral(data.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", dataIdent, __gong__toRawStringLiteral(data.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Acronym = %s", dataIdent, __gong__toRawStringLiteral(data.Acronym)))
			values.WriteString(fmt.Sprintf("\n\t%s.Description = %s", dataIdent, __gong__toRawStringLiteral(data.Description)))
			values.WriteString(fmt.Sprintf("\n\t%s.ComputedPrefix = %s", dataIdent, __gong__toRawStringLiteral(data.ComputedPrefix)))
			values.WriteString(fmt.Sprintf("\n\t%s.IsExpanded = %t", dataIdent, data.IsExpanded))
			values.WriteString(fmt.Sprintf("\n\t%s.SVG_Path = %s", dataIdent, __gong__toRawStringLiteral(data.SVG_Path)))
			values.WriteString(fmt.Sprintf("\n\t%s.InverseAppliedScaling = %f", dataIdent, data.InverseAppliedScaling))
		}
	}
	if stage != nil {
		for _, dataflow := range __gong__sortStageSetInstances(stage.DataFlows, stage.DataFlow_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			dataflowIdent := dataflow.GongGetIdentifier(stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&models.DataFlow{Name: %s}).Stage(stage)", dataflowIdent, __gong__toRawStringLiteral(dataflow.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", dataflowIdent, __gong__toRawStringLiteral(dataflow.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Description = %s", dataflowIdent, __gong__toRawStringLiteral(dataflow.Description)))
			if dataflow.Direction.ToCodeString() != "" {
				values.WriteString(fmt.Sprintf("\n\t%s.Direction = models.%s", dataflowIdent, dataflow.Direction.ToCodeString()))
			} else {
				values.WriteString(fmt.Sprintf("\n\t%s.Direction = \"\"", dataflowIdent))
			}
			values.WriteString(fmt.Sprintf("\n\t%s.IsDatasNodeExpanded = %t", dataflowIdent, dataflow.IsDatasNodeExpanded))
			values.WriteString(fmt.Sprintf("\n\t%s.ComputedPrefix = %s", dataflowIdent, __gong__toRawStringLiteral(dataflow.ComputedPrefix)))
			values.WriteString(fmt.Sprintf("\n\t%s.IsExpanded = %t", dataflowIdent, dataflow.IsExpanded))
			if dataflow.Type.ToCodeString() != "" {
				values.WriteString(fmt.Sprintf("\n\t%s.Type = models.%s", dataflowIdent, dataflow.Type.ToCodeString()))
			} else {
				values.WriteString(fmt.Sprintf("\n\t%s.Type = \"\"", dataflowIdent))
			}
			if dataflow.StartPort != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := dataflow.StartPort.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.StartPort = %s", dataflowIdent, targetIdent))
			}
			if dataflow.EndPort != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := dataflow.EndPort.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.EndPort = %s", dataflowIdent, targetIdent))
			}
			if dataflow.StartExternalPart != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := dataflow.StartExternalPart.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.StartExternalPart = %s", dataflowIdent, targetIdent))
			}
			if dataflow.EndExternalPart != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := dataflow.EndExternalPart.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.EndExternalPart = %s", dataflowIdent, targetIdent))
			}
			for _, elem := range dataflow.Datas {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := elem.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Datas = append(%s.Datas, %s)", dataflowIdent, dataflowIdent, targetIdent))
			}
		}
	}
	if stage != nil {
		for _, dataflowshape := range __gong__sortStageSetInstances(stage.DataFlowShapes, stage.DataFlowShape_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			dataflowshapeIdent := dataflowshape.GongGetIdentifier(stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&models.DataFlowShape{Name: %s}).Stage(stage)", dataflowshapeIdent, __gong__toRawStringLiteral(dataflowshape.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", dataflowshapeIdent, __gong__toRawStringLiteral(dataflowshape.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.StartRatio = %f", dataflowshapeIdent, dataflowshape.StartRatio))
			values.WriteString(fmt.Sprintf("\n\t%s.EndRatio = %f", dataflowshapeIdent, dataflowshape.EndRatio))
			if dataflowshape.StartOrientation.ToCodeString() != "" {
				values.WriteString(fmt.Sprintf("\n\t%s.StartOrientation = models.%s", dataflowshapeIdent, dataflowshape.StartOrientation.ToCodeString()))
			} else {
				values.WriteString(fmt.Sprintf("\n\t%s.StartOrientation = \"\"", dataflowshapeIdent))
			}
			if dataflowshape.EndOrientation.ToCodeString() != "" {
				values.WriteString(fmt.Sprintf("\n\t%s.EndOrientation = models.%s", dataflowshapeIdent, dataflowshape.EndOrientation.ToCodeString()))
			} else {
				values.WriteString(fmt.Sprintf("\n\t%s.EndOrientation = \"\"", dataflowshapeIdent))
			}
			values.WriteString(fmt.Sprintf("\n\t%s.CornerOffsetRatio = %f", dataflowshapeIdent, dataflowshape.CornerOffsetRatio))
			values.WriteString(fmt.Sprintf("\n\t%s.IsHidden = %t", dataflowshapeIdent, dataflowshape.IsHidden))
			values.WriteString(fmt.Sprintf("\n\t%s.IsPresentationOverridden = %t", dataflowshapeIdent, dataflowshape.IsPresentationOverridden))
			values.WriteString(fmt.Sprintf("\n\t%s.Stroke = %s", dataflowshapeIdent, __gong__toRawStringLiteral(dataflowshape.Stroke)))
			values.WriteString(fmt.Sprintf("\n\t%s.StrokeWidth = %f", dataflowshapeIdent, dataflowshape.StrokeWidth))
			values.WriteString(fmt.Sprintf("\n\t%s.StrokeOpacity = %f", dataflowshapeIdent, dataflowshape.StrokeOpacity))
			values.WriteString(fmt.Sprintf("\n\t%s.StrokeDashArray = %s", dataflowshapeIdent, __gong__toRawStringLiteral(dataflowshape.StrokeDashArray)))
			values.WriteString(fmt.Sprintf("\n\t%s.FontSize = %s", dataflowshapeIdent, __gong__toRawStringLiteral(dataflowshape.FontSize)))
			values.WriteString(fmt.Sprintf("\n\t%s.FontColor = %s", dataflowshapeIdent, __gong__toRawStringLiteral(dataflowshape.FontColor)))
			if dataflowshape.DataFlow != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := dataflowshape.DataFlow.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.DataFlow = %s", dataflowshapeIdent, targetIdent))
			}
		}
	}
	if stage != nil {
		for _, datashape := range __gong__sortStageSetInstances(stage.DataShapes, stage.DataShape_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			datashapeIdent := datashape.GongGetIdentifier(stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&models.DataShape{Name: %s}).Stage(stage)", datashapeIdent, __gong__toRawStringLiteral(datashape.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", datashapeIdent, __gong__toRawStringLiteral(datashape.Name)))
			if datashape.Data != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := datashape.Data.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Data = %s", datashapeIdent, targetIdent))
			}
			if datashape.DataFlow != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := datashape.DataFlow.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.DataFlow = %s", datashapeIdent, targetIdent))
			}
		}
	}
	if stage != nil {
		for _, diagramlayerstate := range __gong__sortStageSetInstances(stage.DiagramLayerStates, stage.DiagramLayerState_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			diagramlayerstateIdent := diagramlayerstate.GongGetIdentifier(stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&models.DiagramLayerState{Name: %s}).Stage(stage)", diagramlayerstateIdent, __gong__toRawStringLiteral(diagramlayerstate.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", diagramlayerstateIdent, __gong__toRawStringLiteral(diagramlayerstate.Name)))
			if diagramlayerstate.DiagramStructure != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := diagramlayerstate.DiagramStructure.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.DiagramStructure = %s", diagramlayerstateIdent, targetIdent))
			}
			if diagramlayerstate.LayerDefinition != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := diagramlayerstate.LayerDefinition.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.LayerDefinition = %s", diagramlayerstateIdent, targetIdent))
			}
		}
	}
	if stage != nil {
		for _, diagramstructure := range __gong__sortStageSetInstances(stage.DiagramStructures, stage.DiagramStructure_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			diagramstructureIdent := diagramstructure.GongGetIdentifier(stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&models.DiagramStructure{Name: %s}).Stage(stage)", diagramstructureIdent, __gong__toRawStringLiteral(diagramstructure.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", diagramstructureIdent, __gong__toRawStringLiteral(diagramstructure.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Description = %s", diagramstructureIdent, __gong__toRawStringLiteral(diagramstructure.Description)))
			values.WriteString(fmt.Sprintf("\n\t%s.ComputedPrefix = %s", diagramstructureIdent, __gong__toRawStringLiteral(diagramstructure.ComputedPrefix)))
			values.WriteString(fmt.Sprintf("\n\t%s.IsExpanded = %t", diagramstructureIdent, diagramstructure.IsExpanded))
			values.WriteString(fmt.Sprintf("\n\t%s.IsChecked = %t", diagramstructureIdent, diagramstructure.IsChecked))
			values.WriteString(fmt.Sprintf("\n\t%s.IsEditable_ = %t", diagramstructureIdent, diagramstructure.IsEditable_))
			values.WriteString(fmt.Sprintf("\n\t%s.IsShowPrefix = %t", diagramstructureIdent, diagramstructure.IsShowPrefix))
			values.WriteString(fmt.Sprintf("\n\t%s.DefaultBoxWidth = %f", diagramstructureIdent, diagramstructure.DefaultBoxWidth))
			values.WriteString(fmt.Sprintf("\n\t%s.DefaultBoxHeigth = %f", diagramstructureIdent, diagramstructure.DefaultBoxHeigth))
			values.WriteString(fmt.Sprintf("\n\t%s.IsWithDiscretePorts = %t", diagramstructureIdent, diagramstructure.IsWithDiscretePorts))
			values.WriteString(fmt.Sprintf("\n\t%s.Width = %f", diagramstructureIdent, diagramstructure.Width))
			values.WriteString(fmt.Sprintf("\n\t%s.Height = %f", diagramstructureIdent, diagramstructure.Height))
			values.WriteString(fmt.Sprintf("\n\t%s.IsSystemsNodeExpanded = %t", diagramstructureIdent, diagramstructure.IsSystemsNodeExpanded))
			values.WriteString(fmt.Sprintf("\n\t%s.IsPartsNodeExpanded = %t", diagramstructureIdent, diagramstructure.IsPartsNodeExpanded))
			values.WriteString(fmt.Sprintf("\n\t%s.IsExternalPartsNodeExpanded = %t", diagramstructureIdent, diagramstructure.IsExternalPartsNodeExpanded))
			values.WriteString(fmt.Sprintf("\n\t%s.IsNotesNodeExpanded = %t", diagramstructureIdent, diagramstructure.IsNotesNodeExpanded))
			for _, elem := range diagramstructure.System_Shapes {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := elem.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.System_Shapes = append(%s.System_Shapes, %s)", diagramstructureIdent, diagramstructureIdent, targetIdent))
			}
			for _, elem := range diagramstructure.SystemsWhoseNodeIsExpanded {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := elem.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.SystemsWhoseNodeIsExpanded = append(%s.SystemsWhoseNodeIsExpanded, %s)", diagramstructureIdent, diagramstructureIdent, targetIdent))
			}
			for _, elem := range diagramstructure.Part_Shapes {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := elem.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Part_Shapes = append(%s.Part_Shapes, %s)", diagramstructureIdent, diagramstructureIdent, targetIdent))
			}
			for _, elem := range diagramstructure.PartWhoseNodeIsExpanded {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := elem.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.PartWhoseNodeIsExpanded = append(%s.PartWhoseNodeIsExpanded, %s)", diagramstructureIdent, diagramstructureIdent, targetIdent))
			}
			for _, elem := range diagramstructure.ExternalPart_Shapes {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := elem.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.ExternalPart_Shapes = append(%s.ExternalPart_Shapes, %s)", diagramstructureIdent, diagramstructureIdent, targetIdent))
			}
			for _, elem := range diagramstructure.ExternalPartWhoseNodeIsExpanded {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := elem.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.ExternalPartWhoseNodeIsExpanded = append(%s.ExternalPartWhoseNodeIsExpanded, %s)", diagramstructureIdent, diagramstructureIdent, targetIdent))
			}
			for _, elem := range diagramstructure.ExternalPartsWhoseOutDataFlowsNodeIsExpanded {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := elem.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.ExternalPartsWhoseOutDataFlowsNodeIsExpanded = append(%s.ExternalPartsWhoseOutDataFlowsNodeIsExpanded, %s)", diagramstructureIdent, diagramstructureIdent, targetIdent))
			}
			for _, elem := range diagramstructure.ExternalPartsWhoseInDataFlowsNodeIsExpanded {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := elem.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.ExternalPartsWhoseInDataFlowsNodeIsExpanded = append(%s.ExternalPartsWhoseInDataFlowsNodeIsExpanded, %s)", diagramstructureIdent, diagramstructureIdent, targetIdent))
			}
			for _, elem := range diagramstructure.PortsWhoseNodeIsExpanded {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := elem.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.PortsWhoseNodeIsExpanded = append(%s.PortsWhoseNodeIsExpanded, %s)", diagramstructureIdent, diagramstructureIdent, targetIdent))
			}
			for _, elem := range diagramstructure.Port_Shapes {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := elem.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Port_Shapes = append(%s.Port_Shapes, %s)", diagramstructureIdent, diagramstructureIdent, targetIdent))
			}
			for _, elem := range diagramstructure.ControlFlowsWhoseNodeIsExpanded {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := elem.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.ControlFlowsWhoseNodeIsExpanded = append(%s.ControlFlowsWhoseNodeIsExpanded, %s)", diagramstructureIdent, diagramstructureIdent, targetIdent))
			}
			for _, elem := range diagramstructure.ControlFlow_Shapes {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := elem.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.ControlFlow_Shapes = append(%s.ControlFlow_Shapes, %s)", diagramstructureIdent, diagramstructureIdent, targetIdent))
			}
			for _, elem := range diagramstructure.DataFlowsWhoseNodeIsExpanded {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := elem.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.DataFlowsWhoseNodeIsExpanded = append(%s.DataFlowsWhoseNodeIsExpanded, %s)", diagramstructureIdent, diagramstructureIdent, targetIdent))
			}
			for _, elem := range diagramstructure.DataFlow_Shapes {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := elem.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.DataFlow_Shapes = append(%s.DataFlow_Shapes, %s)", diagramstructureIdent, diagramstructureIdent, targetIdent))
			}
			for _, elem := range diagramstructure.DatasWhoseNodeIsExpanded {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := elem.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.DatasWhoseNodeIsExpanded = append(%s.DatasWhoseNodeIsExpanded, %s)", diagramstructureIdent, diagramstructureIdent, targetIdent))
			}
			for _, elem := range diagramstructure.Data_Shapes {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := elem.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Data_Shapes = append(%s.Data_Shapes, %s)", diagramstructureIdent, diagramstructureIdent, targetIdent))
			}
			for _, elem := range diagramstructure.DataFlowsWhoseDataNodeIsExpanded {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := elem.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.DataFlowsWhoseDataNodeIsExpanded = append(%s.DataFlowsWhoseDataNodeIsExpanded, %s)", diagramstructureIdent, diagramstructureIdent, targetIdent))
			}
			for _, elem := range diagramstructure.AllocatedResourcesWhoseNodeIsExpanded {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := elem.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.AllocatedResourcesWhoseNodeIsExpanded = append(%s.AllocatedResourcesWhoseNodeIsExpanded, %s)", diagramstructureIdent, diagramstructureIdent, targetIdent))
			}
			for _, elem := range diagramstructure.AllocatedResourceShapes {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := elem.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.AllocatedResourceShapes = append(%s.AllocatedResourceShapes, %s)", diagramstructureIdent, diagramstructureIdent, targetIdent))
			}
			for _, elem := range diagramstructure.AllocatedSystemesWhoseNodeIsExpanded {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := elem.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.AllocatedSystemesWhoseNodeIsExpanded = append(%s.AllocatedSystemesWhoseNodeIsExpanded, %s)", diagramstructureIdent, diagramstructureIdent, targetIdent))
			}
			for _, elem := range diagramstructure.AllocatedSystemShapes {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := elem.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.AllocatedSystemShapes = append(%s.AllocatedSystemShapes, %s)", diagramstructureIdent, diagramstructureIdent, targetIdent))
			}
			for _, elem := range diagramstructure.Note_Shapes {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := elem.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Note_Shapes = append(%s.Note_Shapes, %s)", diagramstructureIdent, diagramstructureIdent, targetIdent))
			}
			for _, elem := range diagramstructure.NotesWhoseNodeIsExpanded {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := elem.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.NotesWhoseNodeIsExpanded = append(%s.NotesWhoseNodeIsExpanded, %s)", diagramstructureIdent, diagramstructureIdent, targetIdent))
			}
			for _, elem := range diagramstructure.NotePortShapes {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := elem.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.NotePortShapes = append(%s.NotePortShapes, %s)", diagramstructureIdent, diagramstructureIdent, targetIdent))
			}
			for _, elem := range diagramstructure.NotePartShapes {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := elem.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.NotePartShapes = append(%s.NotePartShapes, %s)", diagramstructureIdent, diagramstructureIdent, targetIdent))
			}
		}
	}
	if stage != nil {
		for _, externalpartshape := range __gong__sortStageSetInstances(stage.ExternalPartShapes, stage.ExternalPartShape_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			externalpartshapeIdent := externalpartshape.GongGetIdentifier(stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&models.ExternalPartShape{Name: %s}).Stage(stage)", externalpartshapeIdent, __gong__toRawStringLiteral(externalpartshape.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", externalpartshapeIdent, __gong__toRawStringLiteral(externalpartshape.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.IsExpanded = %t", externalpartshapeIdent, externalpartshape.IsExpanded))
			values.WriteString(fmt.Sprintf("\n\t%s.X = %f", externalpartshapeIdent, externalpartshape.X))
			values.WriteString(fmt.Sprintf("\n\t%s.Y = %f", externalpartshapeIdent, externalpartshape.Y))
			values.WriteString(fmt.Sprintf("\n\t%s.Width = %f", externalpartshapeIdent, externalpartshape.Width))
			values.WriteString(fmt.Sprintf("\n\t%s.Height = %f", externalpartshapeIdent, externalpartshape.Height))
			values.WriteString(fmt.Sprintf("\n\t%s.IsHidden = %t", externalpartshapeIdent, externalpartshape.IsHidden))
			values.WriteString(fmt.Sprintf("\n\t%s.IsPresentationOverridden = %t", externalpartshapeIdent, externalpartshape.IsPresentationOverridden))
			values.WriteString(fmt.Sprintf("\n\t%s.Color = %s", externalpartshapeIdent, __gong__toRawStringLiteral(externalpartshape.Color)))
			values.WriteString(fmt.Sprintf("\n\t%s.FillOpacity = %f", externalpartshapeIdent, externalpartshape.FillOpacity))
			values.WriteString(fmt.Sprintf("\n\t%s.Stroke = %s", externalpartshapeIdent, __gong__toRawStringLiteral(externalpartshape.Stroke)))
			values.WriteString(fmt.Sprintf("\n\t%s.StrokeWidth = %f", externalpartshapeIdent, externalpartshape.StrokeWidth))
			values.WriteString(fmt.Sprintf("\n\t%s.StrokeOpacity = %f", externalpartshapeIdent, externalpartshape.StrokeOpacity))
			values.WriteString(fmt.Sprintf("\n\t%s.StrokeDashArray = %s", externalpartshapeIdent, __gong__toRawStringLiteral(externalpartshape.StrokeDashArray)))
			values.WriteString(fmt.Sprintf("\n\t%s.RX = %f", externalpartshapeIdent, externalpartshape.RX))
			values.WriteString(fmt.Sprintf("\n\t%s.FontSize = %s", externalpartshapeIdent, __gong__toRawStringLiteral(externalpartshape.FontSize)))
			values.WriteString(fmt.Sprintf("\n\t%s.FontWeight = %s", externalpartshapeIdent, __gong__toRawStringLiteral(externalpartshape.FontWeight)))
			values.WriteString(fmt.Sprintf("\n\t%s.FontColor = %s", externalpartshapeIdent, __gong__toRawStringLiteral(externalpartshape.FontColor)))
			values.WriteString(fmt.Sprintf("\n\t%s.TailHeigth = %f", externalpartshapeIdent, externalpartshape.TailHeigth))
			if externalpartshape.Part != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := externalpartshape.Part.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Part = %s", externalpartshapeIdent, targetIdent))
			}
		}
	}
	if stage != nil {
		for _, layerdefinition := range __gong__sortStageSetInstances(stage.LayerDefinitions, stage.LayerDefinition_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			layerdefinitionIdent := layerdefinition.GongGetIdentifier(stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&models.LayerDefinition{Name: %s}).Stage(stage)", layerdefinitionIdent, __gong__toRawStringLiteral(layerdefinition.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", layerdefinitionIdent, __gong__toRawStringLiteral(layerdefinition.Name)))
			for _, elem := range layerdefinition.Query {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := elem.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Query = append(%s.Query, %s)", layerdefinitionIdent, layerdefinitionIdent, targetIdent))
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
			libraryIdent := library.GongGetIdentifier(stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&models.Library{Name: %s}).Stage(stage)", libraryIdent, __gong__toRawStringLiteral(library.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", libraryIdent, __gong__toRawStringLiteral(library.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Description = %s", libraryIdent, __gong__toRawStringLiteral(library.Description)))
			values.WriteString(fmt.Sprintf("\n\t%s.ComputedPrefix = %s", libraryIdent, __gong__toRawStringLiteral(library.ComputedPrefix)))
			values.WriteString(fmt.Sprintf("\n\t%s.IsExpanded = %t", libraryIdent, library.IsExpanded))
			values.WriteString(fmt.Sprintf("\n\t%s.IsRootLibrary = %t", libraryIdent, library.IsRootLibrary))
			values.WriteString(fmt.Sprintf("\n\t%s.IsSubLibrariesNodeExpanded = %t", libraryIdent, library.IsSubLibrariesNodeExpanded))
			values.WriteString(fmt.Sprintf("\n\t%s.NbPixPerCharacter = %f", libraryIdent, library.NbPixPerCharacter))
			values.WriteString(fmt.Sprintf("\n\t%s.LogoSVGFile = %s", libraryIdent, __gong__toRawStringLiteral(library.LogoSVGFile)))
			values.WriteString(fmt.Sprintf("\n\t%s.IsSystemesNodeExpanded = %t", libraryIdent, library.IsSystemesNodeExpanded))
			values.WriteString(fmt.Sprintf("\n\t%s.IsDataFlowsNodeExpanded = %t", libraryIdent, library.IsDataFlowsNodeExpanded))
			values.WriteString(fmt.Sprintf("\n\t%s.IsDatasNodeExpanded = %t", libraryIdent, library.IsDatasNodeExpanded))
			values.WriteString(fmt.Sprintf("\n\t%s.IsResourcesNodeExpanded = %t", libraryIdent, library.IsResourcesNodeExpanded))
			values.WriteString(fmt.Sprintf("\n\t%s.IsNotesNodeExpanded = %t", libraryIdent, library.IsNotesNodeExpanded))
			values.WriteString(fmt.Sprintf("\n\t%s.IsExpandedTmp = %t", libraryIdent, library.IsExpandedTmp))
			for _, elem := range library.SubLibraries {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := elem.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.SubLibraries = append(%s.SubLibraries, %s)", libraryIdent, libraryIdent, targetIdent))
			}
			for _, elem := range library.SubLibrariesWhoseNodeIsExpanded {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := elem.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.SubLibrariesWhoseNodeIsExpanded = append(%s.SubLibrariesWhoseNodeIsExpanded, %s)", libraryIdent, libraryIdent, targetIdent))
			}
			for _, elem := range library.RootSystemes {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := elem.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.RootSystemes = append(%s.RootSystemes, %s)", libraryIdent, libraryIdent, targetIdent))
			}
			for _, elem := range library.SystemsWhoseNodeIsExpanded {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := elem.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.SystemsWhoseNodeIsExpanded = append(%s.SystemsWhoseNodeIsExpanded, %s)", libraryIdent, libraryIdent, targetIdent))
			}
			for _, elem := range library.RootDataFlows {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := elem.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.RootDataFlows = append(%s.RootDataFlows, %s)", libraryIdent, libraryIdent, targetIdent))
			}
			for _, elem := range library.DataFlowsWhoseNodeIsExpanded {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := elem.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.DataFlowsWhoseNodeIsExpanded = append(%s.DataFlowsWhoseNodeIsExpanded, %s)", libraryIdent, libraryIdent, targetIdent))
			}
			for _, elem := range library.RootDatas {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := elem.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.RootDatas = append(%s.RootDatas, %s)", libraryIdent, libraryIdent, targetIdent))
			}
			for _, elem := range library.DatasWhoseNodeIsExpanded {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := elem.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.DatasWhoseNodeIsExpanded = append(%s.DatasWhoseNodeIsExpanded, %s)", libraryIdent, libraryIdent, targetIdent))
			}
			for _, elem := range library.RootResources {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := elem.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.RootResources = append(%s.RootResources, %s)", libraryIdent, libraryIdent, targetIdent))
			}
			for _, elem := range library.ResourcesWhoseNodeIsExpanded {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := elem.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.ResourcesWhoseNodeIsExpanded = append(%s.ResourcesWhoseNodeIsExpanded, %s)", libraryIdent, libraryIdent, targetIdent))
			}
			for _, elem := range library.PartsWhoseNodeIsExpanded {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := elem.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.PartsWhoseNodeIsExpanded = append(%s.PartsWhoseNodeIsExpanded, %s)", libraryIdent, libraryIdent, targetIdent))
			}
			for _, elem := range library.RootNotes {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := elem.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.RootNotes = append(%s.RootNotes, %s)", libraryIdent, libraryIdent, targetIdent))
			}
			for _, elem := range library.NotesWhoseNodeIsExpanded {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := elem.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.NotesWhoseNodeIsExpanded = append(%s.NotesWhoseNodeIsExpanded, %s)", libraryIdent, libraryIdent, targetIdent))
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
			noteIdent := note.GongGetIdentifier(stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&models.Note{Name: %s}).Stage(stage)", noteIdent, __gong__toRawStringLiteral(note.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", noteIdent, __gong__toRawStringLiteral(note.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Description = %s", noteIdent, __gong__toRawStringLiteral(note.Description)))
			values.WriteString(fmt.Sprintf("\n\t%s.ComputedPrefix = %s", noteIdent, __gong__toRawStringLiteral(note.ComputedPrefix)))
			values.WriteString(fmt.Sprintf("\n\t%s.IsExpanded = %t", noteIdent, note.IsExpanded))
			values.WriteString(fmt.Sprintf("\n\t%s.IsPartsNodeExpanded = %t", noteIdent, note.IsPartsNodeExpanded))
			values.WriteString(fmt.Sprintf("\n\t%s.IsPortsNodeExpanded = %t", noteIdent, note.IsPortsNodeExpanded))
			for _, elem := range note.Parts {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := elem.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Parts = append(%s.Parts, %s)", noteIdent, noteIdent, targetIdent))
			}
			for _, elem := range note.Ports {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := elem.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Ports = append(%s.Ports, %s)", noteIdent, noteIdent, targetIdent))
			}
		}
	}
	if stage != nil {
		for _, notepartshape := range __gong__sortStageSetInstances(stage.NotePartShapes, stage.NotePartShape_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			notepartshapeIdent := notepartshape.GongGetIdentifier(stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&models.NotePartShape{Name: %s}).Stage(stage)", notepartshapeIdent, __gong__toRawStringLiteral(notepartshape.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", notepartshapeIdent, __gong__toRawStringLiteral(notepartshape.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.StartRatio = %f", notepartshapeIdent, notepartshape.StartRatio))
			values.WriteString(fmt.Sprintf("\n\t%s.EndRatio = %f", notepartshapeIdent, notepartshape.EndRatio))
			if notepartshape.StartOrientation.ToCodeString() != "" {
				values.WriteString(fmt.Sprintf("\n\t%s.StartOrientation = models.%s", notepartshapeIdent, notepartshape.StartOrientation.ToCodeString()))
			} else {
				values.WriteString(fmt.Sprintf("\n\t%s.StartOrientation = \"\"", notepartshapeIdent))
			}
			if notepartshape.EndOrientation.ToCodeString() != "" {
				values.WriteString(fmt.Sprintf("\n\t%s.EndOrientation = models.%s", notepartshapeIdent, notepartshape.EndOrientation.ToCodeString()))
			} else {
				values.WriteString(fmt.Sprintf("\n\t%s.EndOrientation = \"\"", notepartshapeIdent))
			}
			values.WriteString(fmt.Sprintf("\n\t%s.CornerOffsetRatio = %f", notepartshapeIdent, notepartshape.CornerOffsetRatio))
			values.WriteString(fmt.Sprintf("\n\t%s.IsHidden = %t", notepartshapeIdent, notepartshape.IsHidden))
			values.WriteString(fmt.Sprintf("\n\t%s.IsPresentationOverridden = %t", notepartshapeIdent, notepartshape.IsPresentationOverridden))
			values.WriteString(fmt.Sprintf("\n\t%s.Stroke = %s", notepartshapeIdent, __gong__toRawStringLiteral(notepartshape.Stroke)))
			values.WriteString(fmt.Sprintf("\n\t%s.StrokeWidth = %f", notepartshapeIdent, notepartshape.StrokeWidth))
			values.WriteString(fmt.Sprintf("\n\t%s.StrokeOpacity = %f", notepartshapeIdent, notepartshape.StrokeOpacity))
			values.WriteString(fmt.Sprintf("\n\t%s.StrokeDashArray = %s", notepartshapeIdent, __gong__toRawStringLiteral(notepartshape.StrokeDashArray)))
			values.WriteString(fmt.Sprintf("\n\t%s.FontSize = %s", notepartshapeIdent, __gong__toRawStringLiteral(notepartshape.FontSize)))
			values.WriteString(fmt.Sprintf("\n\t%s.FontColor = %s", notepartshapeIdent, __gong__toRawStringLiteral(notepartshape.FontColor)))
			if notepartshape.Note != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := notepartshape.Note.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Note = %s", notepartshapeIdent, targetIdent))
			}
			if notepartshape.Part != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := notepartshape.Part.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Part = %s", notepartshapeIdent, targetIdent))
			}
		}
	}
	if stage != nil {
		for _, noteportshape := range __gong__sortStageSetInstances(stage.NotePortShapes, stage.NotePortShape_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			noteportshapeIdent := noteportshape.GongGetIdentifier(stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&models.NotePortShape{Name: %s}).Stage(stage)", noteportshapeIdent, __gong__toRawStringLiteral(noteportshape.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", noteportshapeIdent, __gong__toRawStringLiteral(noteportshape.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.StartRatio = %f", noteportshapeIdent, noteportshape.StartRatio))
			values.WriteString(fmt.Sprintf("\n\t%s.EndRatio = %f", noteportshapeIdent, noteportshape.EndRatio))
			if noteportshape.StartOrientation.ToCodeString() != "" {
				values.WriteString(fmt.Sprintf("\n\t%s.StartOrientation = models.%s", noteportshapeIdent, noteportshape.StartOrientation.ToCodeString()))
			} else {
				values.WriteString(fmt.Sprintf("\n\t%s.StartOrientation = \"\"", noteportshapeIdent))
			}
			if noteportshape.EndOrientation.ToCodeString() != "" {
				values.WriteString(fmt.Sprintf("\n\t%s.EndOrientation = models.%s", noteportshapeIdent, noteportshape.EndOrientation.ToCodeString()))
			} else {
				values.WriteString(fmt.Sprintf("\n\t%s.EndOrientation = \"\"", noteportshapeIdent))
			}
			values.WriteString(fmt.Sprintf("\n\t%s.CornerOffsetRatio = %f", noteportshapeIdent, noteportshape.CornerOffsetRatio))
			values.WriteString(fmt.Sprintf("\n\t%s.IsHidden = %t", noteportshapeIdent, noteportshape.IsHidden))
			values.WriteString(fmt.Sprintf("\n\t%s.IsPresentationOverridden = %t", noteportshapeIdent, noteportshape.IsPresentationOverridden))
			values.WriteString(fmt.Sprintf("\n\t%s.Stroke = %s", noteportshapeIdent, __gong__toRawStringLiteral(noteportshape.Stroke)))
			values.WriteString(fmt.Sprintf("\n\t%s.StrokeWidth = %f", noteportshapeIdent, noteportshape.StrokeWidth))
			values.WriteString(fmt.Sprintf("\n\t%s.StrokeOpacity = %f", noteportshapeIdent, noteportshape.StrokeOpacity))
			values.WriteString(fmt.Sprintf("\n\t%s.StrokeDashArray = %s", noteportshapeIdent, __gong__toRawStringLiteral(noteportshape.StrokeDashArray)))
			values.WriteString(fmt.Sprintf("\n\t%s.FontSize = %s", noteportshapeIdent, __gong__toRawStringLiteral(noteportshape.FontSize)))
			values.WriteString(fmt.Sprintf("\n\t%s.FontColor = %s", noteportshapeIdent, __gong__toRawStringLiteral(noteportshape.FontColor)))
			if noteportshape.Note != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := noteportshape.Note.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Note = %s", noteportshapeIdent, targetIdent))
			}
			if noteportshape.Port != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := noteportshape.Port.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Port = %s", noteportshapeIdent, targetIdent))
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
			noteshapeIdent := noteshape.GongGetIdentifier(stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&models.NoteShape{Name: %s}).Stage(stage)", noteshapeIdent, __gong__toRawStringLiteral(noteshape.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", noteshapeIdent, __gong__toRawStringLiteral(noteshape.Name)))
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
				targetIdent := noteshape.Note.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Note = %s", noteshapeIdent, targetIdent))
			}
		}
	}
	if stage != nil {
		for _, part := range __gong__sortStageSetInstances(stage.Parts, stage.Part_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			partIdent := part.GongGetIdentifier(stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&models.Part{Name: %s}).Stage(stage)", partIdent, __gong__toRawStringLiteral(part.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", partIdent, __gong__toRawStringLiteral(part.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Description = %s", partIdent, __gong__toRawStringLiteral(part.Description)))
			values.WriteString(fmt.Sprintf("\n\t%s.IsPartNameNotSystemName = %t", partIdent, part.IsPartNameNotSystemName))
			values.WriteString(fmt.Sprintf("\n\t%s.IsControlFlowsNodeExpanded = %t", partIdent, part.IsControlFlowsNodeExpanded))
			values.WriteString(fmt.Sprintf("\n\t%s.IsDataFlowsNodeExpanded = %t", partIdent, part.IsDataFlowsNodeExpanded))
			values.WriteString(fmt.Sprintf("\n\t%s.ComputedPrefix = %s", partIdent, __gong__toRawStringLiteral(part.ComputedPrefix)))
			values.WriteString(fmt.Sprintf("\n\t%s.IsExpanded = %t", partIdent, part.IsExpanded))
			values.WriteString(fmt.Sprintf("\n\t%s.IsPortsNodeExpanded = %t", partIdent, part.IsPortsNodeExpanded))
			for _, elem := range part.Ports {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := elem.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Ports = append(%s.Ports, %s)", partIdent, partIdent, targetIdent))
			}
			if part.TypeOfPart != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := part.TypeOfPart.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.TypeOfPart = %s", partIdent, targetIdent))
			}
			for _, elem := range part.ControlFlows {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := elem.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.ControlFlows = append(%s.ControlFlows, %s)", partIdent, partIdent, targetIdent))
			}
			for _, elem := range part.PortWhoseOutControlFlowsNodeIsExpanded {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := elem.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.PortWhoseOutControlFlowsNodeIsExpanded = append(%s.PortWhoseOutControlFlowsNodeIsExpanded, %s)", partIdent, partIdent, targetIdent))
			}
			for _, elem := range part.PortWhoseInControlFlowsNodeIsExpanded {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := elem.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.PortWhoseInControlFlowsNodeIsExpanded = append(%s.PortWhoseInControlFlowsNodeIsExpanded, %s)", partIdent, partIdent, targetIdent))
			}
			for _, elem := range part.PortWhoseOutDataFlowsNodeIsExpanded {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := elem.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.PortWhoseOutDataFlowsNodeIsExpanded = append(%s.PortWhoseOutDataFlowsNodeIsExpanded, %s)", partIdent, partIdent, targetIdent))
			}
			for _, elem := range part.PortWhoseInDataFlowsNodeIsExpanded {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := elem.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.PortWhoseInDataFlowsNodeIsExpanded = append(%s.PortWhoseInDataFlowsNodeIsExpanded, %s)", partIdent, partIdent, targetIdent))
			}
			for _, elem := range part.PartAnchoredPath {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := elem.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.PartAnchoredPath = append(%s.PartAnchoredPath, %s)", partIdent, partIdent, targetIdent))
			}
		}
	}
	if stage != nil {
		for _, partanchoredpath := range __gong__sortStageSetInstances(stage.PartAnchoredPaths, stage.PartAnchoredPath_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			partanchoredpathIdent := partanchoredpath.GongGetIdentifier(stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&models.PartAnchoredPath{Name: %s}).Stage(stage)", partanchoredpathIdent, __gong__toRawStringLiteral(partanchoredpath.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", partanchoredpathIdent, __gong__toRawStringLiteral(partanchoredpath.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Definition = %s", partanchoredpathIdent, __gong__toRawStringLiteral(partanchoredpath.Definition)))
			values.WriteString(fmt.Sprintf("\n\t%s.X_Offset = %f", partanchoredpathIdent, partanchoredpath.X_Offset))
			values.WriteString(fmt.Sprintf("\n\t%s.Y_Offset = %f", partanchoredpathIdent, partanchoredpath.Y_Offset))
			if partanchoredpath.RectAnchorType.ToCodeString() != "" {
				values.WriteString(fmt.Sprintf("\n\t%s.RectAnchorType = models.%s", partanchoredpathIdent, partanchoredpath.RectAnchorType.ToCodeString()))
			} else {
				values.WriteString(fmt.Sprintf("\n\t%s.RectAnchorType = \"\"", partanchoredpathIdent))
			}
			values.WriteString(fmt.Sprintf("\n\t%s.ScalePropotionnally = %t", partanchoredpathIdent, partanchoredpath.ScalePropotionnally))
			values.WriteString(fmt.Sprintf("\n\t%s.AppliedScaling = %f", partanchoredpathIdent, partanchoredpath.AppliedScaling))
			values.WriteString(fmt.Sprintf("\n\t%s.Color = %s", partanchoredpathIdent, __gong__toRawStringLiteral(partanchoredpath.Color)))
			values.WriteString(fmt.Sprintf("\n\t%s.FillOpacity = %f", partanchoredpathIdent, partanchoredpath.FillOpacity))
			values.WriteString(fmt.Sprintf("\n\t%s.Stroke = %s", partanchoredpathIdent, __gong__toRawStringLiteral(partanchoredpath.Stroke)))
			values.WriteString(fmt.Sprintf("\n\t%s.StrokeOpacity = %f", partanchoredpathIdent, partanchoredpath.StrokeOpacity))
			values.WriteString(fmt.Sprintf("\n\t%s.StrokeWidth = %f", partanchoredpathIdent, partanchoredpath.StrokeWidth))
			values.WriteString(fmt.Sprintf("\n\t%s.StrokeDashArray = %s", partanchoredpathIdent, __gong__toRawStringLiteral(partanchoredpath.StrokeDashArray)))
			values.WriteString(fmt.Sprintf("\n\t%s.StrokeDashArrayWhenSelected = %s", partanchoredpathIdent, __gong__toRawStringLiteral(partanchoredpath.StrokeDashArrayWhenSelected)))
			values.WriteString(fmt.Sprintf("\n\t%s.Transform = %s", partanchoredpathIdent, __gong__toRawStringLiteral(partanchoredpath.Transform)))
		}
	}
	if stage != nil {
		for _, partshape := range __gong__sortStageSetInstances(stage.PartShapes, stage.PartShape_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			partshapeIdent := partshape.GongGetIdentifier(stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&models.PartShape{Name: %s}).Stage(stage)", partshapeIdent, __gong__toRawStringLiteral(partshape.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", partshapeIdent, __gong__toRawStringLiteral(partshape.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.IsExpanded = %t", partshapeIdent, partshape.IsExpanded))
			values.WriteString(fmt.Sprintf("\n\t%s.X = %f", partshapeIdent, partshape.X))
			values.WriteString(fmt.Sprintf("\n\t%s.Y = %f", partshapeIdent, partshape.Y))
			values.WriteString(fmt.Sprintf("\n\t%s.Width = %f", partshapeIdent, partshape.Width))
			values.WriteString(fmt.Sprintf("\n\t%s.Height = %f", partshapeIdent, partshape.Height))
			values.WriteString(fmt.Sprintf("\n\t%s.IsHidden = %t", partshapeIdent, partshape.IsHidden))
			values.WriteString(fmt.Sprintf("\n\t%s.IsPresentationOverridden = %t", partshapeIdent, partshape.IsPresentationOverridden))
			values.WriteString(fmt.Sprintf("\n\t%s.Color = %s", partshapeIdent, __gong__toRawStringLiteral(partshape.Color)))
			values.WriteString(fmt.Sprintf("\n\t%s.FillOpacity = %f", partshapeIdent, partshape.FillOpacity))
			values.WriteString(fmt.Sprintf("\n\t%s.Stroke = %s", partshapeIdent, __gong__toRawStringLiteral(partshape.Stroke)))
			values.WriteString(fmt.Sprintf("\n\t%s.StrokeWidth = %f", partshapeIdent, partshape.StrokeWidth))
			values.WriteString(fmt.Sprintf("\n\t%s.StrokeOpacity = %f", partshapeIdent, partshape.StrokeOpacity))
			values.WriteString(fmt.Sprintf("\n\t%s.StrokeDashArray = %s", partshapeIdent, __gong__toRawStringLiteral(partshape.StrokeDashArray)))
			values.WriteString(fmt.Sprintf("\n\t%s.RX = %f", partshapeIdent, partshape.RX))
			values.WriteString(fmt.Sprintf("\n\t%s.FontSize = %s", partshapeIdent, __gong__toRawStringLiteral(partshape.FontSize)))
			values.WriteString(fmt.Sprintf("\n\t%s.FontWeight = %s", partshapeIdent, __gong__toRawStringLiteral(partshape.FontWeight)))
			values.WriteString(fmt.Sprintf("\n\t%s.FontColor = %s", partshapeIdent, __gong__toRawStringLiteral(partshape.FontColor)))
			if partshape.Part != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := partshape.Part.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Part = %s", partshapeIdent, targetIdent))
			}
		}
	}
	if stage != nil {
		for _, port := range __gong__sortStageSetInstances(stage.Ports, stage.Port_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			portIdent := port.GongGetIdentifier(stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&models.Port{Name: %s}).Stage(stage)", portIdent, __gong__toRawStringLiteral(port.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", portIdent, __gong__toRawStringLiteral(port.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Description = %s", portIdent, __gong__toRawStringLiteral(port.Description)))
			values.WriteString(fmt.Sprintf("\n\t%s.ComputedPrefix = %s", portIdent, __gong__toRawStringLiteral(port.ComputedPrefix)))
			values.WriteString(fmt.Sprintf("\n\t%s.IsExpanded = %t", portIdent, port.IsExpanded))
		}
	}
	if stage != nil {
		for _, portshape := range __gong__sortStageSetInstances(stage.PortShapes, stage.PortShape_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			portshapeIdent := portshape.GongGetIdentifier(stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&models.PortShape{Name: %s}).Stage(stage)", portshapeIdent, __gong__toRawStringLiteral(portshape.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", portshapeIdent, __gong__toRawStringLiteral(portshape.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.IsExpanded = %t", portshapeIdent, portshape.IsExpanded))
			values.WriteString(fmt.Sprintf("\n\t%s.X = %f", portshapeIdent, portshape.X))
			values.WriteString(fmt.Sprintf("\n\t%s.Y = %f", portshapeIdent, portshape.Y))
			values.WriteString(fmt.Sprintf("\n\t%s.Width = %f", portshapeIdent, portshape.Width))
			values.WriteString(fmt.Sprintf("\n\t%s.Height = %f", portshapeIdent, portshape.Height))
			values.WriteString(fmt.Sprintf("\n\t%s.IsHidden = %t", portshapeIdent, portshape.IsHidden))
			values.WriteString(fmt.Sprintf("\n\t%s.IsPresentationOverridden = %t", portshapeIdent, portshape.IsPresentationOverridden))
			values.WriteString(fmt.Sprintf("\n\t%s.Color = %s", portshapeIdent, __gong__toRawStringLiteral(portshape.Color)))
			values.WriteString(fmt.Sprintf("\n\t%s.FillOpacity = %f", portshapeIdent, portshape.FillOpacity))
			values.WriteString(fmt.Sprintf("\n\t%s.Stroke = %s", portshapeIdent, __gong__toRawStringLiteral(portshape.Stroke)))
			values.WriteString(fmt.Sprintf("\n\t%s.StrokeWidth = %f", portshapeIdent, portshape.StrokeWidth))
			values.WriteString(fmt.Sprintf("\n\t%s.StrokeOpacity = %f", portshapeIdent, portshape.StrokeOpacity))
			values.WriteString(fmt.Sprintf("\n\t%s.StrokeDashArray = %s", portshapeIdent, __gong__toRawStringLiteral(portshape.StrokeDashArray)))
			values.WriteString(fmt.Sprintf("\n\t%s.RX = %f", portshapeIdent, portshape.RX))
			values.WriteString(fmt.Sprintf("\n\t%s.FontSize = %s", portshapeIdent, __gong__toRawStringLiteral(portshape.FontSize)))
			values.WriteString(fmt.Sprintf("\n\t%s.FontWeight = %s", portshapeIdent, __gong__toRawStringLiteral(portshape.FontWeight)))
			values.WriteString(fmt.Sprintf("\n\t%s.FontColor = %s", portshapeIdent, __gong__toRawStringLiteral(portshape.FontColor)))
			if portshape.Port != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := portshape.Port.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Port = %s", portshapeIdent, targetIdent))
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
			resourceIdent := resource.GongGetIdentifier(stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&models.Resource{Name: %s}).Stage(stage)", resourceIdent, __gong__toRawStringLiteral(resource.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", resourceIdent, __gong__toRawStringLiteral(resource.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Acronym = %s", resourceIdent, __gong__toRawStringLiteral(resource.Acronym)))
			values.WriteString(fmt.Sprintf("\n\t%s.Description = %s", resourceIdent, __gong__toRawStringLiteral(resource.Description)))
			values.WriteString(fmt.Sprintf("\n\t%s.ComputedPrefix = %s", resourceIdent, __gong__toRawStringLiteral(resource.ComputedPrefix)))
			values.WriteString(fmt.Sprintf("\n\t%s.IsExpanded = %t", resourceIdent, resource.IsExpanded))
			values.WriteString(fmt.Sprintf("\n\t%s.SVG_Path = %s", resourceIdent, __gong__toRawStringLiteral(resource.SVG_Path)))
			values.WriteString(fmt.Sprintf("\n\t%s.InverseAppliedScaling = %f", resourceIdent, resource.InverseAppliedScaling))
		}
	}
	if stage != nil {
		for _, semantictag := range __gong__sortStageSetInstances(stage.SemanticTags, stage.SemanticTag_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			semantictagIdent := semantictag.GongGetIdentifier(stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&models.SemanticTag{Name: %s}).Stage(stage)", semantictagIdent, __gong__toRawStringLiteral(semantictag.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", semantictagIdent, __gong__toRawStringLiteral(semantictag.Name)))
			for _, elem := range semantictag.Parts {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := elem.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Parts = append(%s.Parts, %s)", semantictagIdent, semantictagIdent, targetIdent))
			}
		}
	}
	if stage != nil {
		for _, system := range __gong__sortStageSetInstances(stage.Systems, stage.System_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			systemIdent := system.GongGetIdentifier(stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&models.System{Name: %s}).Stage(stage)", systemIdent, __gong__toRawStringLiteral(system.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", systemIdent, __gong__toRawStringLiteral(system.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Description = %s", systemIdent, __gong__toRawStringLiteral(system.Description)))
			values.WriteString(fmt.Sprintf("\n\t%s.ComputedPrefix = %s", systemIdent, __gong__toRawStringLiteral(system.ComputedPrefix)))
			values.WriteString(fmt.Sprintf("\n\t%s.IsExpanded = %t", systemIdent, system.IsExpanded))
			values.WriteString(fmt.Sprintf("\n\t%s.SVG_Path = %s", systemIdent, __gong__toRawStringLiteral(system.SVG_Path)))
			values.WriteString(fmt.Sprintf("\n\t%s.InverseAppliedScaling = %f", systemIdent, system.InverseAppliedScaling))
			values.WriteString(fmt.Sprintf("\n\t%s.IsSubSystemNodeExpanded = %t", systemIdent, system.IsSubSystemNodeExpanded))
			values.WriteString(fmt.Sprintf("\n\t%s.IsDataFlowsNodeExpanded = %t", systemIdent, system.IsDataFlowsNodeExpanded))
			for _, elem := range system.DiagramStructures {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := elem.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.DiagramStructures = append(%s.DiagramStructures, %s)", systemIdent, systemIdent, targetIdent))
			}
			for _, elem := range system.DiagramStructureWhoseNodeIsExpanded {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := elem.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.DiagramStructureWhoseNodeIsExpanded = append(%s.DiagramStructureWhoseNodeIsExpanded, %s)", systemIdent, systemIdent, targetIdent))
			}
			for _, elem := range system.SubSystemes {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := elem.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.SubSystemes = append(%s.SubSystemes, %s)", systemIdent, systemIdent, targetIdent))
			}
			for _, elem := range system.Parts {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := elem.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Parts = append(%s.Parts, %s)", systemIdent, systemIdent, targetIdent))
			}
			for _, elem := range system.PartWhoseNodeIsExpanded {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := elem.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.PartWhoseNodeIsExpanded = append(%s.PartWhoseNodeIsExpanded, %s)", systemIdent, systemIdent, targetIdent))
			}
			for _, elem := range system.DataFlows {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := elem.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.DataFlows = append(%s.DataFlows, %s)", systemIdent, systemIdent, targetIdent))
			}
			for _, elem := range system.ExternalParts {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := elem.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.ExternalParts = append(%s.ExternalParts, %s)", systemIdent, systemIdent, targetIdent))
			}
			for _, elem := range system.ExternalPartWhoseNodeIsExpanded {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := elem.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.ExternalPartWhoseNodeIsExpanded = append(%s.ExternalPartWhoseNodeIsExpanded, %s)", systemIdent, systemIdent, targetIdent))
			}
		}
	}
	if stage != nil {
		for _, systemshape := range __gong__sortStageSetInstances(stage.SystemShapes, stage.SystemShape_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			systemshapeIdent := systemshape.GongGetIdentifier(stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&models.SystemShape{Name: %s}).Stage(stage)", systemshapeIdent, __gong__toRawStringLiteral(systemshape.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", systemshapeIdent, __gong__toRawStringLiteral(systemshape.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.IsExpanded = %t", systemshapeIdent, systemshape.IsExpanded))
			values.WriteString(fmt.Sprintf("\n\t%s.X = %f", systemshapeIdent, systemshape.X))
			values.WriteString(fmt.Sprintf("\n\t%s.Y = %f", systemshapeIdent, systemshape.Y))
			values.WriteString(fmt.Sprintf("\n\t%s.Width = %f", systemshapeIdent, systemshape.Width))
			values.WriteString(fmt.Sprintf("\n\t%s.Height = %f", systemshapeIdent, systemshape.Height))
			values.WriteString(fmt.Sprintf("\n\t%s.IsHidden = %t", systemshapeIdent, systemshape.IsHidden))
			values.WriteString(fmt.Sprintf("\n\t%s.IsPresentationOverridden = %t", systemshapeIdent, systemshape.IsPresentationOverridden))
			values.WriteString(fmt.Sprintf("\n\t%s.Color = %s", systemshapeIdent, __gong__toRawStringLiteral(systemshape.Color)))
			values.WriteString(fmt.Sprintf("\n\t%s.FillOpacity = %f", systemshapeIdent, systemshape.FillOpacity))
			values.WriteString(fmt.Sprintf("\n\t%s.Stroke = %s", systemshapeIdent, __gong__toRawStringLiteral(systemshape.Stroke)))
			values.WriteString(fmt.Sprintf("\n\t%s.StrokeWidth = %f", systemshapeIdent, systemshape.StrokeWidth))
			values.WriteString(fmt.Sprintf("\n\t%s.StrokeOpacity = %f", systemshapeIdent, systemshape.StrokeOpacity))
			values.WriteString(fmt.Sprintf("\n\t%s.StrokeDashArray = %s", systemshapeIdent, __gong__toRawStringLiteral(systemshape.StrokeDashArray)))
			values.WriteString(fmt.Sprintf("\n\t%s.RX = %f", systemshapeIdent, systemshape.RX))
			values.WriteString(fmt.Sprintf("\n\t%s.FontSize = %s", systemshapeIdent, __gong__toRawStringLiteral(systemshape.FontSize)))
			values.WriteString(fmt.Sprintf("\n\t%s.FontWeight = %s", systemshapeIdent, __gong__toRawStringLiteral(systemshape.FontWeight)))
			values.WriteString(fmt.Sprintf("\n\t%s.FontColor = %s", systemshapeIdent, __gong__toRawStringLiteral(systemshape.FontColor)))
			if systemshape.System != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := systemshape.System.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.System = %s", systemshapeIdent, targetIdent))
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
go run github.com/fullstack-lang/gong/dsm/structure/go/models/probe/cmd/probe@%s stage.go
*/
package %s

import (
	"slices"
	"time"

	"github.com/fullstack-lang/gong/dsm/structure/go/models"%s
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

	if modelsPackageName != "" && modelsPackageName != "github.com/fullstack-lang/gong/dsm/structure/go/models" {
		res = strings.ReplaceAll(res, "\"github.com/fullstack-lang/gong/dsm/structure/go/models\"", fmt.Sprintf("%q", modelsPackageName))
	}

	return res, nil
}

// insertion point for marshall field methods
func (allocatedresourceshape *AllocatedResourceShape) GongMarshallField(stage *Stage, fieldName string) (res string) {
	ident := allocatedresourceshape.GongGetIdentifier(stage)
	_ = ident
	switch fieldName {
	case "Name":
		res = __gong__marshallString(ident, "Name", allocatedresourceshape.Name)

	case "Part":
		if allocatedresourceshape.Part != nil {
			res = __gong__marshallPointer(ident, "Part", allocatedresourceshape.Part.GongGetIdentifier(stage))
		} else {
			res = __gong__marshallPointer(ident, "Part", "nil")
		}
	case "Resource":
		if allocatedresourceshape.Resource != nil {
			res = __gong__marshallPointer(ident, "Resource", allocatedresourceshape.Resource.GongGetIdentifier(stage))
		} else {
			res = __gong__marshallPointer(ident, "Resource", "nil")
		}
	default:
		log.Panicf("Unknown field %s for Gongstruct AllocatedResourceShape", fieldName)
	}
	return
}

func (allocatedsystemshape *AllocatedSystemShape) GongMarshallField(stage *Stage, fieldName string) (res string) {
	ident := allocatedsystemshape.GongGetIdentifier(stage)
	_ = ident
	switch fieldName {
	case "Name":
		res = __gong__marshallString(ident, "Name", allocatedsystemshape.Name)

	case "Part":
		if allocatedsystemshape.Part != nil {
			res = __gong__marshallPointer(ident, "Part", allocatedsystemshape.Part.GongGetIdentifier(stage))
		} else {
			res = __gong__marshallPointer(ident, "Part", "nil")
		}
	case "System":
		if allocatedsystemshape.System != nil {
			res = __gong__marshallPointer(ident, "System", allocatedsystemshape.System.GongGetIdentifier(stage))
		} else {
			res = __gong__marshallPointer(ident, "System", "nil")
		}
	default:
		log.Panicf("Unknown field %s for Gongstruct AllocatedSystemShape", fieldName)
	}
	return
}

func (controlflow *ControlFlow) GongMarshallField(stage *Stage, fieldName string) (res string) {
	ident := controlflow.GongGetIdentifier(stage)
	_ = ident
	switch fieldName {
	case "Name":
		res = __gong__marshallString(ident, "Name", controlflow.Name)
	case "Description":
		res = __gong__marshallString(ident, "Description", controlflow.Description)
	case "ComputedPrefix":
		res = __gong__marshallString(ident, "ComputedPrefix", controlflow.ComputedPrefix)
	case "IsExpanded":
		res = __gong__marshallBool(ident, "IsExpanded", controlflow.IsExpanded)

	case "Start":
		if controlflow.Start != nil {
			res = __gong__marshallPointer(ident, "Start", controlflow.Start.GongGetIdentifier(stage))
		} else {
			res = __gong__marshallPointer(ident, "Start", "nil")
		}
	case "End":
		if controlflow.End != nil {
			res = __gong__marshallPointer(ident, "End", controlflow.End.GongGetIdentifier(stage))
		} else {
			res = __gong__marshallPointer(ident, "End", "nil")
		}
	default:
		log.Panicf("Unknown field %s for Gongstruct ControlFlow", fieldName)
	}
	return
}

func (controlflowshape *ControlFlowShape) GongMarshallField(stage *Stage, fieldName string) (res string) {
	ident := controlflowshape.GongGetIdentifier(stage)
	_ = ident
	switch fieldName {
	case "Name":
		res = __gong__marshallString(ident, "Name", controlflowshape.Name)
	case "StartRatio":
		res = __gong__marshallFloat(ident, "StartRatio", controlflowshape.StartRatio)
	case "EndRatio":
		res = __gong__marshallFloat(ident, "EndRatio", controlflowshape.EndRatio)
	case "StartOrientation":
		res = __gong__marshallEnumString(ident, "StartOrientation", controlflowshape.StartOrientation.ToCodeString())
	case "EndOrientation":
		res = __gong__marshallEnumString(ident, "EndOrientation", controlflowshape.EndOrientation.ToCodeString())
	case "CornerOffsetRatio":
		res = __gong__marshallFloat(ident, "CornerOffsetRatio", controlflowshape.CornerOffsetRatio)
	case "IsHidden":
		res = __gong__marshallBool(ident, "IsHidden", controlflowshape.IsHidden)
	case "IsPresentationOverridden":
		res = __gong__marshallBool(ident, "IsPresentationOverridden", controlflowshape.IsPresentationOverridden)
	case "Stroke":
		res = __gong__marshallString(ident, "Stroke", controlflowshape.Stroke)
	case "StrokeWidth":
		res = __gong__marshallFloat(ident, "StrokeWidth", controlflowshape.StrokeWidth)
	case "StrokeOpacity":
		res = __gong__marshallFloat(ident, "StrokeOpacity", controlflowshape.StrokeOpacity)
	case "StrokeDashArray":
		res = __gong__marshallString(ident, "StrokeDashArray", controlflowshape.StrokeDashArray)
	case "FontSize":
		res = __gong__marshallString(ident, "FontSize", controlflowshape.FontSize)
	case "FontColor":
		res = __gong__marshallString(ident, "FontColor", controlflowshape.FontColor)

	case "ControlFlow":
		if controlflowshape.ControlFlow != nil {
			res = __gong__marshallPointer(ident, "ControlFlow", controlflowshape.ControlFlow.GongGetIdentifier(stage))
		} else {
			res = __gong__marshallPointer(ident, "ControlFlow", "nil")
		}
	default:
		log.Panicf("Unknown field %s for Gongstruct ControlFlowShape", fieldName)
	}
	return
}

func (data *Data) GongMarshallField(stage *Stage, fieldName string) (res string) {
	ident := data.GongGetIdentifier(stage)
	_ = ident
	switch fieldName {
	case "Name":
		res = __gong__marshallString(ident, "Name", data.Name)
	case "Acronym":
		res = __gong__marshallString(ident, "Acronym", data.Acronym)
	case "Description":
		res = __gong__marshallString(ident, "Description", data.Description)
	case "ComputedPrefix":
		res = __gong__marshallString(ident, "ComputedPrefix", data.ComputedPrefix)
	case "IsExpanded":
		res = __gong__marshallBool(ident, "IsExpanded", data.IsExpanded)
	case "SVG_Path":
		res = __gong__marshallString(ident, "SVG_Path", data.SVG_Path)
	case "InverseAppliedScaling":
		res = __gong__marshallFloat(ident, "InverseAppliedScaling", data.InverseAppliedScaling)

	default:
		log.Panicf("Unknown field %s for Gongstruct Data", fieldName)
	}
	return
}

func (dataflow *DataFlow) GongMarshallField(stage *Stage, fieldName string) (res string) {
	ident := dataflow.GongGetIdentifier(stage)
	_ = ident
	switch fieldName {
	case "Name":
		res = __gong__marshallString(ident, "Name", dataflow.Name)
	case "Description":
		res = __gong__marshallString(ident, "Description", dataflow.Description)
	case "Direction":
		res = __gong__marshallEnumString(ident, "Direction", dataflow.Direction.ToCodeString())
	case "IsDatasNodeExpanded":
		res = __gong__marshallBool(ident, "IsDatasNodeExpanded", dataflow.IsDatasNodeExpanded)
	case "ComputedPrefix":
		res = __gong__marshallString(ident, "ComputedPrefix", dataflow.ComputedPrefix)
	case "IsExpanded":
		res = __gong__marshallBool(ident, "IsExpanded", dataflow.IsExpanded)
	case "Type":
		res = __gong__marshallEnumString(ident, "Type", dataflow.Type.ToCodeString())

	case "StartPort":
		if dataflow.StartPort != nil {
			res = __gong__marshallPointer(ident, "StartPort", dataflow.StartPort.GongGetIdentifier(stage))
		} else {
			res = __gong__marshallPointer(ident, "StartPort", "nil")
		}
	case "EndPort":
		if dataflow.EndPort != nil {
			res = __gong__marshallPointer(ident, "EndPort", dataflow.EndPort.GongGetIdentifier(stage))
		} else {
			res = __gong__marshallPointer(ident, "EndPort", "nil")
		}
	case "StartExternalPart":
		if dataflow.StartExternalPart != nil {
			res = __gong__marshallPointer(ident, "StartExternalPart", dataflow.StartExternalPart.GongGetIdentifier(stage))
		} else {
			res = __gong__marshallPointer(ident, "StartExternalPart", "nil")
		}
	case "EndExternalPart":
		if dataflow.EndExternalPart != nil {
			res = __gong__marshallPointer(ident, "EndExternalPart", dataflow.EndExternalPart.GongGetIdentifier(stage))
		} else {
			res = __gong__marshallPointer(ident, "EndExternalPart", "nil")
		}
	case "Datas":
		var sb strings.Builder
		for _, _data := range dataflow.Datas {
			sb.WriteString(__gong__marshallSliceOfPointers(ident, "Datas", _data.GongGetIdentifier(stage)))
		}
		res = sb.String()
	default:
		log.Panicf("Unknown field %s for Gongstruct DataFlow", fieldName)
	}
	return
}

func (dataflowshape *DataFlowShape) GongMarshallField(stage *Stage, fieldName string) (res string) {
	ident := dataflowshape.GongGetIdentifier(stage)
	_ = ident
	switch fieldName {
	case "Name":
		res = __gong__marshallString(ident, "Name", dataflowshape.Name)
	case "StartRatio":
		res = __gong__marshallFloat(ident, "StartRatio", dataflowshape.StartRatio)
	case "EndRatio":
		res = __gong__marshallFloat(ident, "EndRatio", dataflowshape.EndRatio)
	case "StartOrientation":
		res = __gong__marshallEnumString(ident, "StartOrientation", dataflowshape.StartOrientation.ToCodeString())
	case "EndOrientation":
		res = __gong__marshallEnumString(ident, "EndOrientation", dataflowshape.EndOrientation.ToCodeString())
	case "CornerOffsetRatio":
		res = __gong__marshallFloat(ident, "CornerOffsetRatio", dataflowshape.CornerOffsetRatio)
	case "IsHidden":
		res = __gong__marshallBool(ident, "IsHidden", dataflowshape.IsHidden)
	case "IsPresentationOverridden":
		res = __gong__marshallBool(ident, "IsPresentationOverridden", dataflowshape.IsPresentationOverridden)
	case "Stroke":
		res = __gong__marshallString(ident, "Stroke", dataflowshape.Stroke)
	case "StrokeWidth":
		res = __gong__marshallFloat(ident, "StrokeWidth", dataflowshape.StrokeWidth)
	case "StrokeOpacity":
		res = __gong__marshallFloat(ident, "StrokeOpacity", dataflowshape.StrokeOpacity)
	case "StrokeDashArray":
		res = __gong__marshallString(ident, "StrokeDashArray", dataflowshape.StrokeDashArray)
	case "FontSize":
		res = __gong__marshallString(ident, "FontSize", dataflowshape.FontSize)
	case "FontColor":
		res = __gong__marshallString(ident, "FontColor", dataflowshape.FontColor)

	case "DataFlow":
		if dataflowshape.DataFlow != nil {
			res = __gong__marshallPointer(ident, "DataFlow", dataflowshape.DataFlow.GongGetIdentifier(stage))
		} else {
			res = __gong__marshallPointer(ident, "DataFlow", "nil")
		}
	default:
		log.Panicf("Unknown field %s for Gongstruct DataFlowShape", fieldName)
	}
	return
}

func (datashape *DataShape) GongMarshallField(stage *Stage, fieldName string) (res string) {
	ident := datashape.GongGetIdentifier(stage)
	_ = ident
	switch fieldName {
	case "Name":
		res = __gong__marshallString(ident, "Name", datashape.Name)

	case "Data":
		if datashape.Data != nil {
			res = __gong__marshallPointer(ident, "Data", datashape.Data.GongGetIdentifier(stage))
		} else {
			res = __gong__marshallPointer(ident, "Data", "nil")
		}
	case "DataFlow":
		if datashape.DataFlow != nil {
			res = __gong__marshallPointer(ident, "DataFlow", datashape.DataFlow.GongGetIdentifier(stage))
		} else {
			res = __gong__marshallPointer(ident, "DataFlow", "nil")
		}
	default:
		log.Panicf("Unknown field %s for Gongstruct DataShape", fieldName)
	}
	return
}

func (diagramlayerstate *DiagramLayerState) GongMarshallField(stage *Stage, fieldName string) (res string) {
	ident := diagramlayerstate.GongGetIdentifier(stage)
	_ = ident
	switch fieldName {
	case "Name":
		res = __gong__marshallString(ident, "Name", diagramlayerstate.Name)

	case "DiagramStructure":
		if diagramlayerstate.DiagramStructure != nil {
			res = __gong__marshallPointer(ident, "DiagramStructure", diagramlayerstate.DiagramStructure.GongGetIdentifier(stage))
		} else {
			res = __gong__marshallPointer(ident, "DiagramStructure", "nil")
		}
	case "LayerDefinition":
		if diagramlayerstate.LayerDefinition != nil {
			res = __gong__marshallPointer(ident, "LayerDefinition", diagramlayerstate.LayerDefinition.GongGetIdentifier(stage))
		} else {
			res = __gong__marshallPointer(ident, "LayerDefinition", "nil")
		}
	default:
		log.Panicf("Unknown field %s for Gongstruct DiagramLayerState", fieldName)
	}
	return
}

func (diagramstructure *DiagramStructure) GongMarshallField(stage *Stage, fieldName string) (res string) {
	ident := diagramstructure.GongGetIdentifier(stage)
	_ = ident
	switch fieldName {
	case "Name":
		res = __gong__marshallString(ident, "Name", diagramstructure.Name)
	case "Description":
		res = __gong__marshallString(ident, "Description", diagramstructure.Description)
	case "ComputedPrefix":
		res = __gong__marshallString(ident, "ComputedPrefix", diagramstructure.ComputedPrefix)
	case "IsExpanded":
		res = __gong__marshallBool(ident, "IsExpanded", diagramstructure.IsExpanded)
	case "IsChecked":
		res = __gong__marshallBool(ident, "IsChecked", diagramstructure.IsChecked)
	case "IsEditable_":
		res = __gong__marshallBool(ident, "IsEditable_", diagramstructure.IsEditable_)
	case "IsShowPrefix":
		res = __gong__marshallBool(ident, "IsShowPrefix", diagramstructure.IsShowPrefix)
	case "DefaultBoxWidth":
		res = __gong__marshallFloat(ident, "DefaultBoxWidth", diagramstructure.DefaultBoxWidth)
	case "DefaultBoxHeigth":
		res = __gong__marshallFloat(ident, "DefaultBoxHeigth", diagramstructure.DefaultBoxHeigth)
	case "IsWithDiscretePorts":
		res = __gong__marshallBool(ident, "IsWithDiscretePorts", diagramstructure.IsWithDiscretePorts)
	case "Width":
		res = __gong__marshallFloat(ident, "Width", diagramstructure.Width)
	case "Height":
		res = __gong__marshallFloat(ident, "Height", diagramstructure.Height)
	case "IsSystemsNodeExpanded":
		res = __gong__marshallBool(ident, "IsSystemsNodeExpanded", diagramstructure.IsSystemsNodeExpanded)
	case "IsPartsNodeExpanded":
		res = __gong__marshallBool(ident, "IsPartsNodeExpanded", diagramstructure.IsPartsNodeExpanded)
	case "IsExternalPartsNodeExpanded":
		res = __gong__marshallBool(ident, "IsExternalPartsNodeExpanded", diagramstructure.IsExternalPartsNodeExpanded)
	case "IsNotesNodeExpanded":
		res = __gong__marshallBool(ident, "IsNotesNodeExpanded", diagramstructure.IsNotesNodeExpanded)

	case "System_Shapes":
		var sb strings.Builder
		for _, _systemshape := range diagramstructure.System_Shapes {
			sb.WriteString(__gong__marshallSliceOfPointers(ident, "System_Shapes", _systemshape.GongGetIdentifier(stage)))
		}
		res = sb.String()
	case "SystemsWhoseNodeIsExpanded":
		var sb strings.Builder
		for _, _system := range diagramstructure.SystemsWhoseNodeIsExpanded {
			sb.WriteString(__gong__marshallSliceOfPointers(ident, "SystemsWhoseNodeIsExpanded", _system.GongGetIdentifier(stage)))
		}
		res = sb.String()
	case "Part_Shapes":
		var sb strings.Builder
		for _, _partshape := range diagramstructure.Part_Shapes {
			sb.WriteString(__gong__marshallSliceOfPointers(ident, "Part_Shapes", _partshape.GongGetIdentifier(stage)))
		}
		res = sb.String()
	case "PartWhoseNodeIsExpanded":
		var sb strings.Builder
		for _, _part := range diagramstructure.PartWhoseNodeIsExpanded {
			sb.WriteString(__gong__marshallSliceOfPointers(ident, "PartWhoseNodeIsExpanded", _part.GongGetIdentifier(stage)))
		}
		res = sb.String()
	case "ExternalPart_Shapes":
		var sb strings.Builder
		for _, _externalpartshape := range diagramstructure.ExternalPart_Shapes {
			sb.WriteString(__gong__marshallSliceOfPointers(ident, "ExternalPart_Shapes", _externalpartshape.GongGetIdentifier(stage)))
		}
		res = sb.String()
	case "ExternalPartWhoseNodeIsExpanded":
		var sb strings.Builder
		for _, _part := range diagramstructure.ExternalPartWhoseNodeIsExpanded {
			sb.WriteString(__gong__marshallSliceOfPointers(ident, "ExternalPartWhoseNodeIsExpanded", _part.GongGetIdentifier(stage)))
		}
		res = sb.String()
	case "ExternalPartsWhoseOutDataFlowsNodeIsExpanded":
		var sb strings.Builder
		for _, _part := range diagramstructure.ExternalPartsWhoseOutDataFlowsNodeIsExpanded {
			sb.WriteString(__gong__marshallSliceOfPointers(ident, "ExternalPartsWhoseOutDataFlowsNodeIsExpanded", _part.GongGetIdentifier(stage)))
		}
		res = sb.String()
	case "ExternalPartsWhoseInDataFlowsNodeIsExpanded":
		var sb strings.Builder
		for _, _part := range diagramstructure.ExternalPartsWhoseInDataFlowsNodeIsExpanded {
			sb.WriteString(__gong__marshallSliceOfPointers(ident, "ExternalPartsWhoseInDataFlowsNodeIsExpanded", _part.GongGetIdentifier(stage)))
		}
		res = sb.String()
	case "PortsWhoseNodeIsExpanded":
		var sb strings.Builder
		for _, _port := range diagramstructure.PortsWhoseNodeIsExpanded {
			sb.WriteString(__gong__marshallSliceOfPointers(ident, "PortsWhoseNodeIsExpanded", _port.GongGetIdentifier(stage)))
		}
		res = sb.String()
	case "Port_Shapes":
		var sb strings.Builder
		for _, _portshape := range diagramstructure.Port_Shapes {
			sb.WriteString(__gong__marshallSliceOfPointers(ident, "Port_Shapes", _portshape.GongGetIdentifier(stage)))
		}
		res = sb.String()
	case "ControlFlowsWhoseNodeIsExpanded":
		var sb strings.Builder
		for _, _controlflow := range diagramstructure.ControlFlowsWhoseNodeIsExpanded {
			sb.WriteString(__gong__marshallSliceOfPointers(ident, "ControlFlowsWhoseNodeIsExpanded", _controlflow.GongGetIdentifier(stage)))
		}
		res = sb.String()
	case "ControlFlow_Shapes":
		var sb strings.Builder
		for _, _controlflowshape := range diagramstructure.ControlFlow_Shapes {
			sb.WriteString(__gong__marshallSliceOfPointers(ident, "ControlFlow_Shapes", _controlflowshape.GongGetIdentifier(stage)))
		}
		res = sb.String()
	case "DataFlowsWhoseNodeIsExpanded":
		var sb strings.Builder
		for _, _dataflow := range diagramstructure.DataFlowsWhoseNodeIsExpanded {
			sb.WriteString(__gong__marshallSliceOfPointers(ident, "DataFlowsWhoseNodeIsExpanded", _dataflow.GongGetIdentifier(stage)))
		}
		res = sb.String()
	case "DataFlow_Shapes":
		var sb strings.Builder
		for _, _dataflowshape := range diagramstructure.DataFlow_Shapes {
			sb.WriteString(__gong__marshallSliceOfPointers(ident, "DataFlow_Shapes", _dataflowshape.GongGetIdentifier(stage)))
		}
		res = sb.String()
	case "DatasWhoseNodeIsExpanded":
		var sb strings.Builder
		for _, _data := range diagramstructure.DatasWhoseNodeIsExpanded {
			sb.WriteString(__gong__marshallSliceOfPointers(ident, "DatasWhoseNodeIsExpanded", _data.GongGetIdentifier(stage)))
		}
		res = sb.String()
	case "Data_Shapes":
		var sb strings.Builder
		for _, _datashape := range diagramstructure.Data_Shapes {
			sb.WriteString(__gong__marshallSliceOfPointers(ident, "Data_Shapes", _datashape.GongGetIdentifier(stage)))
		}
		res = sb.String()
	case "DataFlowsWhoseDataNodeIsExpanded":
		var sb strings.Builder
		for _, _dataflow := range diagramstructure.DataFlowsWhoseDataNodeIsExpanded {
			sb.WriteString(__gong__marshallSliceOfPointers(ident, "DataFlowsWhoseDataNodeIsExpanded", _dataflow.GongGetIdentifier(stage)))
		}
		res = sb.String()
	case "AllocatedResourcesWhoseNodeIsExpanded":
		var sb strings.Builder
		for _, _resource := range diagramstructure.AllocatedResourcesWhoseNodeIsExpanded {
			sb.WriteString(__gong__marshallSliceOfPointers(ident, "AllocatedResourcesWhoseNodeIsExpanded", _resource.GongGetIdentifier(stage)))
		}
		res = sb.String()
	case "AllocatedResourceShapes":
		var sb strings.Builder
		for _, _allocatedresourceshape := range diagramstructure.AllocatedResourceShapes {
			sb.WriteString(__gong__marshallSliceOfPointers(ident, "AllocatedResourceShapes", _allocatedresourceshape.GongGetIdentifier(stage)))
		}
		res = sb.String()
	case "AllocatedSystemesWhoseNodeIsExpanded":
		var sb strings.Builder
		for _, _system := range diagramstructure.AllocatedSystemesWhoseNodeIsExpanded {
			sb.WriteString(__gong__marshallSliceOfPointers(ident, "AllocatedSystemesWhoseNodeIsExpanded", _system.GongGetIdentifier(stage)))
		}
		res = sb.String()
	case "AllocatedSystemShapes":
		var sb strings.Builder
		for _, _allocatedsystemshape := range diagramstructure.AllocatedSystemShapes {
			sb.WriteString(__gong__marshallSliceOfPointers(ident, "AllocatedSystemShapes", _allocatedsystemshape.GongGetIdentifier(stage)))
		}
		res = sb.String()
	case "Note_Shapes":
		var sb strings.Builder
		for _, _noteshape := range diagramstructure.Note_Shapes {
			sb.WriteString(__gong__marshallSliceOfPointers(ident, "Note_Shapes", _noteshape.GongGetIdentifier(stage)))
		}
		res = sb.String()
	case "NotesWhoseNodeIsExpanded":
		var sb strings.Builder
		for _, _note := range diagramstructure.NotesWhoseNodeIsExpanded {
			sb.WriteString(__gong__marshallSliceOfPointers(ident, "NotesWhoseNodeIsExpanded", _note.GongGetIdentifier(stage)))
		}
		res = sb.String()
	case "NotePortShapes":
		var sb strings.Builder
		for _, _noteportshape := range diagramstructure.NotePortShapes {
			sb.WriteString(__gong__marshallSliceOfPointers(ident, "NotePortShapes", _noteportshape.GongGetIdentifier(stage)))
		}
		res = sb.String()
	case "NotePartShapes":
		var sb strings.Builder
		for _, _notepartshape := range diagramstructure.NotePartShapes {
			sb.WriteString(__gong__marshallSliceOfPointers(ident, "NotePartShapes", _notepartshape.GongGetIdentifier(stage)))
		}
		res = sb.String()
	default:
		log.Panicf("Unknown field %s for Gongstruct DiagramStructure", fieldName)
	}
	return
}

func (externalpartshape *ExternalPartShape) GongMarshallField(stage *Stage, fieldName string) (res string) {
	ident := externalpartshape.GongGetIdentifier(stage)
	_ = ident
	switch fieldName {
	case "Name":
		res = __gong__marshallString(ident, "Name", externalpartshape.Name)
	case "IsExpanded":
		res = __gong__marshallBool(ident, "IsExpanded", externalpartshape.IsExpanded)
	case "X":
		res = __gong__marshallFloat(ident, "X", externalpartshape.X)
	case "Y":
		res = __gong__marshallFloat(ident, "Y", externalpartshape.Y)
	case "Width":
		res = __gong__marshallFloat(ident, "Width", externalpartshape.Width)
	case "Height":
		res = __gong__marshallFloat(ident, "Height", externalpartshape.Height)
	case "IsHidden":
		res = __gong__marshallBool(ident, "IsHidden", externalpartshape.IsHidden)
	case "IsPresentationOverridden":
		res = __gong__marshallBool(ident, "IsPresentationOverridden", externalpartshape.IsPresentationOverridden)
	case "Color":
		res = __gong__marshallString(ident, "Color", externalpartshape.Color)
	case "FillOpacity":
		res = __gong__marshallFloat(ident, "FillOpacity", externalpartshape.FillOpacity)
	case "Stroke":
		res = __gong__marshallString(ident, "Stroke", externalpartshape.Stroke)
	case "StrokeWidth":
		res = __gong__marshallFloat(ident, "StrokeWidth", externalpartshape.StrokeWidth)
	case "StrokeOpacity":
		res = __gong__marshallFloat(ident, "StrokeOpacity", externalpartshape.StrokeOpacity)
	case "StrokeDashArray":
		res = __gong__marshallString(ident, "StrokeDashArray", externalpartshape.StrokeDashArray)
	case "RX":
		res = __gong__marshallFloat(ident, "RX", externalpartshape.RX)
	case "FontSize":
		res = __gong__marshallString(ident, "FontSize", externalpartshape.FontSize)
	case "FontWeight":
		res = __gong__marshallString(ident, "FontWeight", externalpartshape.FontWeight)
	case "FontColor":
		res = __gong__marshallString(ident, "FontColor", externalpartshape.FontColor)
	case "TailHeigth":
		res = __gong__marshallFloat(ident, "TailHeigth", externalpartshape.TailHeigth)

	case "Part":
		if externalpartshape.Part != nil {
			res = __gong__marshallPointer(ident, "Part", externalpartshape.Part.GongGetIdentifier(stage))
		} else {
			res = __gong__marshallPointer(ident, "Part", "nil")
		}
	default:
		log.Panicf("Unknown field %s for Gongstruct ExternalPartShape", fieldName)
	}
	return
}

func (layerdefinition *LayerDefinition) GongMarshallField(stage *Stage, fieldName string) (res string) {
	ident := layerdefinition.GongGetIdentifier(stage)
	_ = ident
	switch fieldName {
	case "Name":
		res = __gong__marshallString(ident, "Name", layerdefinition.Name)

	case "Query":
		var sb strings.Builder
		for _, _semantictag := range layerdefinition.Query {
			sb.WriteString(__gong__marshallSliceOfPointers(ident, "Query", _semantictag.GongGetIdentifier(stage)))
		}
		res = sb.String()
	default:
		log.Panicf("Unknown field %s for Gongstruct LayerDefinition", fieldName)
	}
	return
}

func (library *Library) GongMarshallField(stage *Stage, fieldName string) (res string) {
	ident := library.GongGetIdentifier(stage)
	_ = ident
	switch fieldName {
	case "Name":
		res = __gong__marshallString(ident, "Name", library.Name)
	case "Description":
		res = __gong__marshallString(ident, "Description", library.Description)
	case "ComputedPrefix":
		res = __gong__marshallString(ident, "ComputedPrefix", library.ComputedPrefix)
	case "IsExpanded":
		res = __gong__marshallBool(ident, "IsExpanded", library.IsExpanded)
	case "IsRootLibrary":
		res = __gong__marshallBool(ident, "IsRootLibrary", library.IsRootLibrary)
	case "IsSubLibrariesNodeExpanded":
		res = __gong__marshallBool(ident, "IsSubLibrariesNodeExpanded", library.IsSubLibrariesNodeExpanded)
	case "NbPixPerCharacter":
		res = __gong__marshallFloat(ident, "NbPixPerCharacter", library.NbPixPerCharacter)
	case "LogoSVGFile":
		res = __gong__marshallString(ident, "LogoSVGFile", library.LogoSVGFile)
	case "IsSystemesNodeExpanded":
		res = __gong__marshallBool(ident, "IsSystemesNodeExpanded", library.IsSystemesNodeExpanded)
	case "IsDataFlowsNodeExpanded":
		res = __gong__marshallBool(ident, "IsDataFlowsNodeExpanded", library.IsDataFlowsNodeExpanded)
	case "IsDatasNodeExpanded":
		res = __gong__marshallBool(ident, "IsDatasNodeExpanded", library.IsDatasNodeExpanded)
	case "IsResourcesNodeExpanded":
		res = __gong__marshallBool(ident, "IsResourcesNodeExpanded", library.IsResourcesNodeExpanded)
	case "IsNotesNodeExpanded":
		res = __gong__marshallBool(ident, "IsNotesNodeExpanded", library.IsNotesNodeExpanded)
	case "IsExpandedTmp":
		res = __gong__marshallBool(ident, "IsExpandedTmp", library.IsExpandedTmp)

	case "SubLibraries":
		var sb strings.Builder
		for _, _library := range library.SubLibraries {
			sb.WriteString(__gong__marshallSliceOfPointers(ident, "SubLibraries", _library.GongGetIdentifier(stage)))
		}
		res = sb.String()
	case "SubLibrariesWhoseNodeIsExpanded":
		var sb strings.Builder
		for _, _library := range library.SubLibrariesWhoseNodeIsExpanded {
			sb.WriteString(__gong__marshallSliceOfPointers(ident, "SubLibrariesWhoseNodeIsExpanded", _library.GongGetIdentifier(stage)))
		}
		res = sb.String()
	case "RootSystemes":
		var sb strings.Builder
		for _, _system := range library.RootSystemes {
			sb.WriteString(__gong__marshallSliceOfPointers(ident, "RootSystemes", _system.GongGetIdentifier(stage)))
		}
		res = sb.String()
	case "SystemsWhoseNodeIsExpanded":
		var sb strings.Builder
		for _, _system := range library.SystemsWhoseNodeIsExpanded {
			sb.WriteString(__gong__marshallSliceOfPointers(ident, "SystemsWhoseNodeIsExpanded", _system.GongGetIdentifier(stage)))
		}
		res = sb.String()
	case "RootDataFlows":
		var sb strings.Builder
		for _, _dataflow := range library.RootDataFlows {
			sb.WriteString(__gong__marshallSliceOfPointers(ident, "RootDataFlows", _dataflow.GongGetIdentifier(stage)))
		}
		res = sb.String()
	case "DataFlowsWhoseNodeIsExpanded":
		var sb strings.Builder
		for _, _dataflow := range library.DataFlowsWhoseNodeIsExpanded {
			sb.WriteString(__gong__marshallSliceOfPointers(ident, "DataFlowsWhoseNodeIsExpanded", _dataflow.GongGetIdentifier(stage)))
		}
		res = sb.String()
	case "RootDatas":
		var sb strings.Builder
		for _, _data := range library.RootDatas {
			sb.WriteString(__gong__marshallSliceOfPointers(ident, "RootDatas", _data.GongGetIdentifier(stage)))
		}
		res = sb.String()
	case "DatasWhoseNodeIsExpanded":
		var sb strings.Builder
		for _, _data := range library.DatasWhoseNodeIsExpanded {
			sb.WriteString(__gong__marshallSliceOfPointers(ident, "DatasWhoseNodeIsExpanded", _data.GongGetIdentifier(stage)))
		}
		res = sb.String()
	case "RootResources":
		var sb strings.Builder
		for _, _resource := range library.RootResources {
			sb.WriteString(__gong__marshallSliceOfPointers(ident, "RootResources", _resource.GongGetIdentifier(stage)))
		}
		res = sb.String()
	case "ResourcesWhoseNodeIsExpanded":
		var sb strings.Builder
		for _, _resource := range library.ResourcesWhoseNodeIsExpanded {
			sb.WriteString(__gong__marshallSliceOfPointers(ident, "ResourcesWhoseNodeIsExpanded", _resource.GongGetIdentifier(stage)))
		}
		res = sb.String()
	case "PartsWhoseNodeIsExpanded":
		var sb strings.Builder
		for _, _part := range library.PartsWhoseNodeIsExpanded {
			sb.WriteString(__gong__marshallSliceOfPointers(ident, "PartsWhoseNodeIsExpanded", _part.GongGetIdentifier(stage)))
		}
		res = sb.String()
	case "RootNotes":
		var sb strings.Builder
		for _, _note := range library.RootNotes {
			sb.WriteString(__gong__marshallSliceOfPointers(ident, "RootNotes", _note.GongGetIdentifier(stage)))
		}
		res = sb.String()
	case "NotesWhoseNodeIsExpanded":
		var sb strings.Builder
		for _, _note := range library.NotesWhoseNodeIsExpanded {
			sb.WriteString(__gong__marshallSliceOfPointers(ident, "NotesWhoseNodeIsExpanded", _note.GongGetIdentifier(stage)))
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
	case "Description":
		res = __gong__marshallString(ident, "Description", note.Description)
	case "ComputedPrefix":
		res = __gong__marshallString(ident, "ComputedPrefix", note.ComputedPrefix)
	case "IsExpanded":
		res = __gong__marshallBool(ident, "IsExpanded", note.IsExpanded)
	case "IsPartsNodeExpanded":
		res = __gong__marshallBool(ident, "IsPartsNodeExpanded", note.IsPartsNodeExpanded)
	case "IsPortsNodeExpanded":
		res = __gong__marshallBool(ident, "IsPortsNodeExpanded", note.IsPortsNodeExpanded)

	case "Parts":
		var sb strings.Builder
		for _, _part := range note.Parts {
			sb.WriteString(__gong__marshallSliceOfPointers(ident, "Parts", _part.GongGetIdentifier(stage)))
		}
		res = sb.String()
	case "Ports":
		var sb strings.Builder
		for _, _port := range note.Ports {
			sb.WriteString(__gong__marshallSliceOfPointers(ident, "Ports", _port.GongGetIdentifier(stage)))
		}
		res = sb.String()
	default:
		log.Panicf("Unknown field %s for Gongstruct Note", fieldName)
	}
	return
}

func (notepartshape *NotePartShape) GongMarshallField(stage *Stage, fieldName string) (res string) {
	ident := notepartshape.GongGetIdentifier(stage)
	_ = ident
	switch fieldName {
	case "Name":
		res = __gong__marshallString(ident, "Name", notepartshape.Name)
	case "StartRatio":
		res = __gong__marshallFloat(ident, "StartRatio", notepartshape.StartRatio)
	case "EndRatio":
		res = __gong__marshallFloat(ident, "EndRatio", notepartshape.EndRatio)
	case "StartOrientation":
		res = __gong__marshallEnumString(ident, "StartOrientation", notepartshape.StartOrientation.ToCodeString())
	case "EndOrientation":
		res = __gong__marshallEnumString(ident, "EndOrientation", notepartshape.EndOrientation.ToCodeString())
	case "CornerOffsetRatio":
		res = __gong__marshallFloat(ident, "CornerOffsetRatio", notepartshape.CornerOffsetRatio)
	case "IsHidden":
		res = __gong__marshallBool(ident, "IsHidden", notepartshape.IsHidden)
	case "IsPresentationOverridden":
		res = __gong__marshallBool(ident, "IsPresentationOverridden", notepartshape.IsPresentationOverridden)
	case "Stroke":
		res = __gong__marshallString(ident, "Stroke", notepartshape.Stroke)
	case "StrokeWidth":
		res = __gong__marshallFloat(ident, "StrokeWidth", notepartshape.StrokeWidth)
	case "StrokeOpacity":
		res = __gong__marshallFloat(ident, "StrokeOpacity", notepartshape.StrokeOpacity)
	case "StrokeDashArray":
		res = __gong__marshallString(ident, "StrokeDashArray", notepartshape.StrokeDashArray)
	case "FontSize":
		res = __gong__marshallString(ident, "FontSize", notepartshape.FontSize)
	case "FontColor":
		res = __gong__marshallString(ident, "FontColor", notepartshape.FontColor)

	case "Note":
		if notepartshape.Note != nil {
			res = __gong__marshallPointer(ident, "Note", notepartshape.Note.GongGetIdentifier(stage))
		} else {
			res = __gong__marshallPointer(ident, "Note", "nil")
		}
	case "Part":
		if notepartshape.Part != nil {
			res = __gong__marshallPointer(ident, "Part", notepartshape.Part.GongGetIdentifier(stage))
		} else {
			res = __gong__marshallPointer(ident, "Part", "nil")
		}
	default:
		log.Panicf("Unknown field %s for Gongstruct NotePartShape", fieldName)
	}
	return
}

func (noteportshape *NotePortShape) GongMarshallField(stage *Stage, fieldName string) (res string) {
	ident := noteportshape.GongGetIdentifier(stage)
	_ = ident
	switch fieldName {
	case "Name":
		res = __gong__marshallString(ident, "Name", noteportshape.Name)
	case "StartRatio":
		res = __gong__marshallFloat(ident, "StartRatio", noteportshape.StartRatio)
	case "EndRatio":
		res = __gong__marshallFloat(ident, "EndRatio", noteportshape.EndRatio)
	case "StartOrientation":
		res = __gong__marshallEnumString(ident, "StartOrientation", noteportshape.StartOrientation.ToCodeString())
	case "EndOrientation":
		res = __gong__marshallEnumString(ident, "EndOrientation", noteportshape.EndOrientation.ToCodeString())
	case "CornerOffsetRatio":
		res = __gong__marshallFloat(ident, "CornerOffsetRatio", noteportshape.CornerOffsetRatio)
	case "IsHidden":
		res = __gong__marshallBool(ident, "IsHidden", noteportshape.IsHidden)
	case "IsPresentationOverridden":
		res = __gong__marshallBool(ident, "IsPresentationOverridden", noteportshape.IsPresentationOverridden)
	case "Stroke":
		res = __gong__marshallString(ident, "Stroke", noteportshape.Stroke)
	case "StrokeWidth":
		res = __gong__marshallFloat(ident, "StrokeWidth", noteportshape.StrokeWidth)
	case "StrokeOpacity":
		res = __gong__marshallFloat(ident, "StrokeOpacity", noteportshape.StrokeOpacity)
	case "StrokeDashArray":
		res = __gong__marshallString(ident, "StrokeDashArray", noteportshape.StrokeDashArray)
	case "FontSize":
		res = __gong__marshallString(ident, "FontSize", noteportshape.FontSize)
	case "FontColor":
		res = __gong__marshallString(ident, "FontColor", noteportshape.FontColor)

	case "Note":
		if noteportshape.Note != nil {
			res = __gong__marshallPointer(ident, "Note", noteportshape.Note.GongGetIdentifier(stage))
		} else {
			res = __gong__marshallPointer(ident, "Note", "nil")
		}
	case "Port":
		if noteportshape.Port != nil {
			res = __gong__marshallPointer(ident, "Port", noteportshape.Port.GongGetIdentifier(stage))
		} else {
			res = __gong__marshallPointer(ident, "Port", "nil")
		}
	default:
		log.Panicf("Unknown field %s for Gongstruct NotePortShape", fieldName)
	}
	return
}

func (noteshape *NoteShape) GongMarshallField(stage *Stage, fieldName string) (res string) {
	ident := noteshape.GongGetIdentifier(stage)
	_ = ident
	switch fieldName {
	case "Name":
		res = __gong__marshallString(ident, "Name", noteshape.Name)
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

func (part *Part) GongMarshallField(stage *Stage, fieldName string) (res string) {
	ident := part.GongGetIdentifier(stage)
	_ = ident
	switch fieldName {
	case "Name":
		res = __gong__marshallString(ident, "Name", part.Name)
	case "Description":
		res = __gong__marshallString(ident, "Description", part.Description)
	case "IsPartNameNotSystemName":
		res = __gong__marshallBool(ident, "IsPartNameNotSystemName", part.IsPartNameNotSystemName)
	case "IsControlFlowsNodeExpanded":
		res = __gong__marshallBool(ident, "IsControlFlowsNodeExpanded", part.IsControlFlowsNodeExpanded)
	case "IsDataFlowsNodeExpanded":
		res = __gong__marshallBool(ident, "IsDataFlowsNodeExpanded", part.IsDataFlowsNodeExpanded)
	case "ComputedPrefix":
		res = __gong__marshallString(ident, "ComputedPrefix", part.ComputedPrefix)
	case "IsExpanded":
		res = __gong__marshallBool(ident, "IsExpanded", part.IsExpanded)
	case "IsPortsNodeExpanded":
		res = __gong__marshallBool(ident, "IsPortsNodeExpanded", part.IsPortsNodeExpanded)

	case "Ports":
		var sb strings.Builder
		for _, _port := range part.Ports {
			sb.WriteString(__gong__marshallSliceOfPointers(ident, "Ports", _port.GongGetIdentifier(stage)))
		}
		res = sb.String()
	case "TypeOfPart":
		if part.TypeOfPart != nil {
			res = __gong__marshallPointer(ident, "TypeOfPart", part.TypeOfPart.GongGetIdentifier(stage))
		} else {
			res = __gong__marshallPointer(ident, "TypeOfPart", "nil")
		}
	case "ControlFlows":
		var sb strings.Builder
		for _, _controlflow := range part.ControlFlows {
			sb.WriteString(__gong__marshallSliceOfPointers(ident, "ControlFlows", _controlflow.GongGetIdentifier(stage)))
		}
		res = sb.String()
	case "PortWhoseOutControlFlowsNodeIsExpanded":
		var sb strings.Builder
		for _, _port := range part.PortWhoseOutControlFlowsNodeIsExpanded {
			sb.WriteString(__gong__marshallSliceOfPointers(ident, "PortWhoseOutControlFlowsNodeIsExpanded", _port.GongGetIdentifier(stage)))
		}
		res = sb.String()
	case "PortWhoseInControlFlowsNodeIsExpanded":
		var sb strings.Builder
		for _, _port := range part.PortWhoseInControlFlowsNodeIsExpanded {
			sb.WriteString(__gong__marshallSliceOfPointers(ident, "PortWhoseInControlFlowsNodeIsExpanded", _port.GongGetIdentifier(stage)))
		}
		res = sb.String()
	case "PortWhoseOutDataFlowsNodeIsExpanded":
		var sb strings.Builder
		for _, _port := range part.PortWhoseOutDataFlowsNodeIsExpanded {
			sb.WriteString(__gong__marshallSliceOfPointers(ident, "PortWhoseOutDataFlowsNodeIsExpanded", _port.GongGetIdentifier(stage)))
		}
		res = sb.String()
	case "PortWhoseInDataFlowsNodeIsExpanded":
		var sb strings.Builder
		for _, _port := range part.PortWhoseInDataFlowsNodeIsExpanded {
			sb.WriteString(__gong__marshallSliceOfPointers(ident, "PortWhoseInDataFlowsNodeIsExpanded", _port.GongGetIdentifier(stage)))
		}
		res = sb.String()
	case "PartAnchoredPath":
		var sb strings.Builder
		for _, _partanchoredpath := range part.PartAnchoredPath {
			sb.WriteString(__gong__marshallSliceOfPointers(ident, "PartAnchoredPath", _partanchoredpath.GongGetIdentifier(stage)))
		}
		res = sb.String()
	default:
		log.Panicf("Unknown field %s for Gongstruct Part", fieldName)
	}
	return
}

func (partanchoredpath *PartAnchoredPath) GongMarshallField(stage *Stage, fieldName string) (res string) {
	ident := partanchoredpath.GongGetIdentifier(stage)
	_ = ident
	switch fieldName {
	case "Name":
		res = __gong__marshallString(ident, "Name", partanchoredpath.Name)
	case "Definition":
		res = __gong__marshallString(ident, "Definition", partanchoredpath.Definition)
	case "X_Offset":
		res = __gong__marshallFloat(ident, "X_Offset", partanchoredpath.X_Offset)
	case "Y_Offset":
		res = __gong__marshallFloat(ident, "Y_Offset", partanchoredpath.Y_Offset)
	case "RectAnchorType":
		res = __gong__marshallEnumString(ident, "RectAnchorType", partanchoredpath.RectAnchorType.ToCodeString())
	case "ScalePropotionnally":
		res = __gong__marshallBool(ident, "ScalePropotionnally", partanchoredpath.ScalePropotionnally)
	case "AppliedScaling":
		res = __gong__marshallFloat(ident, "AppliedScaling", partanchoredpath.AppliedScaling)
	case "Color":
		res = __gong__marshallString(ident, "Color", partanchoredpath.Color)
	case "FillOpacity":
		res = __gong__marshallFloat(ident, "FillOpacity", partanchoredpath.FillOpacity)
	case "Stroke":
		res = __gong__marshallString(ident, "Stroke", partanchoredpath.Stroke)
	case "StrokeOpacity":
		res = __gong__marshallFloat(ident, "StrokeOpacity", partanchoredpath.StrokeOpacity)
	case "StrokeWidth":
		res = __gong__marshallFloat(ident, "StrokeWidth", partanchoredpath.StrokeWidth)
	case "StrokeDashArray":
		res = __gong__marshallString(ident, "StrokeDashArray", partanchoredpath.StrokeDashArray)
	case "StrokeDashArrayWhenSelected":
		res = __gong__marshallString(ident, "StrokeDashArrayWhenSelected", partanchoredpath.StrokeDashArrayWhenSelected)
	case "Transform":
		res = __gong__marshallString(ident, "Transform", partanchoredpath.Transform)

	default:
		log.Panicf("Unknown field %s for Gongstruct PartAnchoredPath", fieldName)
	}
	return
}

func (partshape *PartShape) GongMarshallField(stage *Stage, fieldName string) (res string) {
	ident := partshape.GongGetIdentifier(stage)
	_ = ident
	switch fieldName {
	case "Name":
		res = __gong__marshallString(ident, "Name", partshape.Name)
	case "IsExpanded":
		res = __gong__marshallBool(ident, "IsExpanded", partshape.IsExpanded)
	case "X":
		res = __gong__marshallFloat(ident, "X", partshape.X)
	case "Y":
		res = __gong__marshallFloat(ident, "Y", partshape.Y)
	case "Width":
		res = __gong__marshallFloat(ident, "Width", partshape.Width)
	case "Height":
		res = __gong__marshallFloat(ident, "Height", partshape.Height)
	case "IsHidden":
		res = __gong__marshallBool(ident, "IsHidden", partshape.IsHidden)
	case "IsPresentationOverridden":
		res = __gong__marshallBool(ident, "IsPresentationOverridden", partshape.IsPresentationOverridden)
	case "Color":
		res = __gong__marshallString(ident, "Color", partshape.Color)
	case "FillOpacity":
		res = __gong__marshallFloat(ident, "FillOpacity", partshape.FillOpacity)
	case "Stroke":
		res = __gong__marshallString(ident, "Stroke", partshape.Stroke)
	case "StrokeWidth":
		res = __gong__marshallFloat(ident, "StrokeWidth", partshape.StrokeWidth)
	case "StrokeOpacity":
		res = __gong__marshallFloat(ident, "StrokeOpacity", partshape.StrokeOpacity)
	case "StrokeDashArray":
		res = __gong__marshallString(ident, "StrokeDashArray", partshape.StrokeDashArray)
	case "RX":
		res = __gong__marshallFloat(ident, "RX", partshape.RX)
	case "FontSize":
		res = __gong__marshallString(ident, "FontSize", partshape.FontSize)
	case "FontWeight":
		res = __gong__marshallString(ident, "FontWeight", partshape.FontWeight)
	case "FontColor":
		res = __gong__marshallString(ident, "FontColor", partshape.FontColor)

	case "Part":
		if partshape.Part != nil {
			res = __gong__marshallPointer(ident, "Part", partshape.Part.GongGetIdentifier(stage))
		} else {
			res = __gong__marshallPointer(ident, "Part", "nil")
		}
	default:
		log.Panicf("Unknown field %s for Gongstruct PartShape", fieldName)
	}
	return
}

func (port *Port) GongMarshallField(stage *Stage, fieldName string) (res string) {
	ident := port.GongGetIdentifier(stage)
	_ = ident
	switch fieldName {
	case "Name":
		res = __gong__marshallString(ident, "Name", port.Name)
	case "Description":
		res = __gong__marshallString(ident, "Description", port.Description)
	case "ComputedPrefix":
		res = __gong__marshallString(ident, "ComputedPrefix", port.ComputedPrefix)
	case "IsExpanded":
		res = __gong__marshallBool(ident, "IsExpanded", port.IsExpanded)

	default:
		log.Panicf("Unknown field %s for Gongstruct Port", fieldName)
	}
	return
}

func (portshape *PortShape) GongMarshallField(stage *Stage, fieldName string) (res string) {
	ident := portshape.GongGetIdentifier(stage)
	_ = ident
	switch fieldName {
	case "Name":
		res = __gong__marshallString(ident, "Name", portshape.Name)
	case "IsExpanded":
		res = __gong__marshallBool(ident, "IsExpanded", portshape.IsExpanded)
	case "X":
		res = __gong__marshallFloat(ident, "X", portshape.X)
	case "Y":
		res = __gong__marshallFloat(ident, "Y", portshape.Y)
	case "Width":
		res = __gong__marshallFloat(ident, "Width", portshape.Width)
	case "Height":
		res = __gong__marshallFloat(ident, "Height", portshape.Height)
	case "IsHidden":
		res = __gong__marshallBool(ident, "IsHidden", portshape.IsHidden)
	case "IsPresentationOverridden":
		res = __gong__marshallBool(ident, "IsPresentationOverridden", portshape.IsPresentationOverridden)
	case "Color":
		res = __gong__marshallString(ident, "Color", portshape.Color)
	case "FillOpacity":
		res = __gong__marshallFloat(ident, "FillOpacity", portshape.FillOpacity)
	case "Stroke":
		res = __gong__marshallString(ident, "Stroke", portshape.Stroke)
	case "StrokeWidth":
		res = __gong__marshallFloat(ident, "StrokeWidth", portshape.StrokeWidth)
	case "StrokeOpacity":
		res = __gong__marshallFloat(ident, "StrokeOpacity", portshape.StrokeOpacity)
	case "StrokeDashArray":
		res = __gong__marshallString(ident, "StrokeDashArray", portshape.StrokeDashArray)
	case "RX":
		res = __gong__marshallFloat(ident, "RX", portshape.RX)
	case "FontSize":
		res = __gong__marshallString(ident, "FontSize", portshape.FontSize)
	case "FontWeight":
		res = __gong__marshallString(ident, "FontWeight", portshape.FontWeight)
	case "FontColor":
		res = __gong__marshallString(ident, "FontColor", portshape.FontColor)

	case "Port":
		if portshape.Port != nil {
			res = __gong__marshallPointer(ident, "Port", portshape.Port.GongGetIdentifier(stage))
		} else {
			res = __gong__marshallPointer(ident, "Port", "nil")
		}
	default:
		log.Panicf("Unknown field %s for Gongstruct PortShape", fieldName)
	}
	return
}

func (resource *Resource) GongMarshallField(stage *Stage, fieldName string) (res string) {
	ident := resource.GongGetIdentifier(stage)
	_ = ident
	switch fieldName {
	case "Name":
		res = __gong__marshallString(ident, "Name", resource.Name)
	case "Acronym":
		res = __gong__marshallString(ident, "Acronym", resource.Acronym)
	case "Description":
		res = __gong__marshallString(ident, "Description", resource.Description)
	case "ComputedPrefix":
		res = __gong__marshallString(ident, "ComputedPrefix", resource.ComputedPrefix)
	case "IsExpanded":
		res = __gong__marshallBool(ident, "IsExpanded", resource.IsExpanded)
	case "SVG_Path":
		res = __gong__marshallString(ident, "SVG_Path", resource.SVG_Path)
	case "InverseAppliedScaling":
		res = __gong__marshallFloat(ident, "InverseAppliedScaling", resource.InverseAppliedScaling)

	default:
		log.Panicf("Unknown field %s for Gongstruct Resource", fieldName)
	}
	return
}

func (semantictag *SemanticTag) GongMarshallField(stage *Stage, fieldName string) (res string) {
	ident := semantictag.GongGetIdentifier(stage)
	_ = ident
	switch fieldName {
	case "Name":
		res = __gong__marshallString(ident, "Name", semantictag.Name)

	case "Parts":
		var sb strings.Builder
		for _, _part := range semantictag.Parts {
			sb.WriteString(__gong__marshallSliceOfPointers(ident, "Parts", _part.GongGetIdentifier(stage)))
		}
		res = sb.String()
	default:
		log.Panicf("Unknown field %s for Gongstruct SemanticTag", fieldName)
	}
	return
}

func (system *System) GongMarshallField(stage *Stage, fieldName string) (res string) {
	ident := system.GongGetIdentifier(stage)
	_ = ident
	switch fieldName {
	case "Name":
		res = __gong__marshallString(ident, "Name", system.Name)
	case "Description":
		res = __gong__marshallString(ident, "Description", system.Description)
	case "ComputedPrefix":
		res = __gong__marshallString(ident, "ComputedPrefix", system.ComputedPrefix)
	case "IsExpanded":
		res = __gong__marshallBool(ident, "IsExpanded", system.IsExpanded)
	case "SVG_Path":
		res = __gong__marshallString(ident, "SVG_Path", system.SVG_Path)
	case "InverseAppliedScaling":
		res = __gong__marshallFloat(ident, "InverseAppliedScaling", system.InverseAppliedScaling)
	case "IsSubSystemNodeExpanded":
		res = __gong__marshallBool(ident, "IsSubSystemNodeExpanded", system.IsSubSystemNodeExpanded)
	case "IsDataFlowsNodeExpanded":
		res = __gong__marshallBool(ident, "IsDataFlowsNodeExpanded", system.IsDataFlowsNodeExpanded)

	case "DiagramStructures":
		var sb strings.Builder
		for _, _diagramstructure := range system.DiagramStructures {
			sb.WriteString(__gong__marshallSliceOfPointers(ident, "DiagramStructures", _diagramstructure.GongGetIdentifier(stage)))
		}
		res = sb.String()
	case "DiagramStructureWhoseNodeIsExpanded":
		var sb strings.Builder
		for _, _diagramstructure := range system.DiagramStructureWhoseNodeIsExpanded {
			sb.WriteString(__gong__marshallSliceOfPointers(ident, "DiagramStructureWhoseNodeIsExpanded", _diagramstructure.GongGetIdentifier(stage)))
		}
		res = sb.String()
	case "SubSystemes":
		var sb strings.Builder
		for _, _system := range system.SubSystemes {
			sb.WriteString(__gong__marshallSliceOfPointers(ident, "SubSystemes", _system.GongGetIdentifier(stage)))
		}
		res = sb.String()
	case "Parts":
		var sb strings.Builder
		for _, _part := range system.Parts {
			sb.WriteString(__gong__marshallSliceOfPointers(ident, "Parts", _part.GongGetIdentifier(stage)))
		}
		res = sb.String()
	case "PartWhoseNodeIsExpanded":
		var sb strings.Builder
		for _, _part := range system.PartWhoseNodeIsExpanded {
			sb.WriteString(__gong__marshallSliceOfPointers(ident, "PartWhoseNodeIsExpanded", _part.GongGetIdentifier(stage)))
		}
		res = sb.String()
	case "DataFlows":
		var sb strings.Builder
		for _, _dataflow := range system.DataFlows {
			sb.WriteString(__gong__marshallSliceOfPointers(ident, "DataFlows", _dataflow.GongGetIdentifier(stage)))
		}
		res = sb.String()
	case "ExternalParts":
		var sb strings.Builder
		for _, _part := range system.ExternalParts {
			sb.WriteString(__gong__marshallSliceOfPointers(ident, "ExternalParts", _part.GongGetIdentifier(stage)))
		}
		res = sb.String()
	case "ExternalPartWhoseNodeIsExpanded":
		var sb strings.Builder
		for _, _part := range system.ExternalPartWhoseNodeIsExpanded {
			sb.WriteString(__gong__marshallSliceOfPointers(ident, "ExternalPartWhoseNodeIsExpanded", _part.GongGetIdentifier(stage)))
		}
		res = sb.String()
	default:
		log.Panicf("Unknown field %s for Gongstruct System", fieldName)
	}
	return
}

func (systemshape *SystemShape) GongMarshallField(stage *Stage, fieldName string) (res string) {
	ident := systemshape.GongGetIdentifier(stage)
	_ = ident
	switch fieldName {
	case "Name":
		res = __gong__marshallString(ident, "Name", systemshape.Name)
	case "IsExpanded":
		res = __gong__marshallBool(ident, "IsExpanded", systemshape.IsExpanded)
	case "X":
		res = __gong__marshallFloat(ident, "X", systemshape.X)
	case "Y":
		res = __gong__marshallFloat(ident, "Y", systemshape.Y)
	case "Width":
		res = __gong__marshallFloat(ident, "Width", systemshape.Width)
	case "Height":
		res = __gong__marshallFloat(ident, "Height", systemshape.Height)
	case "IsHidden":
		res = __gong__marshallBool(ident, "IsHidden", systemshape.IsHidden)
	case "IsPresentationOverridden":
		res = __gong__marshallBool(ident, "IsPresentationOverridden", systemshape.IsPresentationOverridden)
	case "Color":
		res = __gong__marshallString(ident, "Color", systemshape.Color)
	case "FillOpacity":
		res = __gong__marshallFloat(ident, "FillOpacity", systemshape.FillOpacity)
	case "Stroke":
		res = __gong__marshallString(ident, "Stroke", systemshape.Stroke)
	case "StrokeWidth":
		res = __gong__marshallFloat(ident, "StrokeWidth", systemshape.StrokeWidth)
	case "StrokeOpacity":
		res = __gong__marshallFloat(ident, "StrokeOpacity", systemshape.StrokeOpacity)
	case "StrokeDashArray":
		res = __gong__marshallString(ident, "StrokeDashArray", systemshape.StrokeDashArray)
	case "RX":
		res = __gong__marshallFloat(ident, "RX", systemshape.RX)
	case "FontSize":
		res = __gong__marshallString(ident, "FontSize", systemshape.FontSize)
	case "FontWeight":
		res = __gong__marshallString(ident, "FontWeight", systemshape.FontWeight)
	case "FontColor":
		res = __gong__marshallString(ident, "FontColor", systemshape.FontColor)

	case "System":
		if systemshape.System != nil {
			res = __gong__marshallPointer(ident, "System", systemshape.System.GongGetIdentifier(stage))
		} else {
			res = __gong__marshallPointer(ident, "System", "nil")
		}
	default:
		log.Panicf("Unknown field %s for Gongstruct SystemShape", fieldName)
	}
	return
}

// insertion point for marshall all fields methods
func (allocatedresourceshape *AllocatedResourceShape) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(allocatedresourceshape.GongMarshallField(stage, "Name"))
		pointersInitializesStatements.WriteString(allocatedresourceshape.GongMarshallField(stage, "Part"))
		pointersInitializesStatements.WriteString(allocatedresourceshape.GongMarshallField(stage, "Resource"))
	}
	initRes = initializerStatements.String()
	ptrRes = pointersInitializesStatements.String()
	return
}
func (allocatedsystemshape *AllocatedSystemShape) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(allocatedsystemshape.GongMarshallField(stage, "Name"))
		pointersInitializesStatements.WriteString(allocatedsystemshape.GongMarshallField(stage, "Part"))
		pointersInitializesStatements.WriteString(allocatedsystemshape.GongMarshallField(stage, "System"))
	}
	initRes = initializerStatements.String()
	ptrRes = pointersInitializesStatements.String()
	return
}
func (controlflow *ControlFlow) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(controlflow.GongMarshallField(stage, "Name"))
		initializerStatements.WriteString(controlflow.GongMarshallField(stage, "Description"))
		initializerStatements.WriteString(controlflow.GongMarshallField(stage, "ComputedPrefix"))
		initializerStatements.WriteString(controlflow.GongMarshallField(stage, "IsExpanded"))
		pointersInitializesStatements.WriteString(controlflow.GongMarshallField(stage, "Start"))
		pointersInitializesStatements.WriteString(controlflow.GongMarshallField(stage, "End"))
	}
	initRes = initializerStatements.String()
	ptrRes = pointersInitializesStatements.String()
	return
}
func (controlflowshape *ControlFlowShape) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(controlflowshape.GongMarshallField(stage, "Name"))
		pointersInitializesStatements.WriteString(controlflowshape.GongMarshallField(stage, "ControlFlow"))
		initializerStatements.WriteString(controlflowshape.GongMarshallField(stage, "StartRatio"))
		initializerStatements.WriteString(controlflowshape.GongMarshallField(stage, "EndRatio"))
		initializerStatements.WriteString(controlflowshape.GongMarshallField(stage, "StartOrientation"))
		initializerStatements.WriteString(controlflowshape.GongMarshallField(stage, "EndOrientation"))
		initializerStatements.WriteString(controlflowshape.GongMarshallField(stage, "CornerOffsetRatio"))
		initializerStatements.WriteString(controlflowshape.GongMarshallField(stage, "IsHidden"))
		initializerStatements.WriteString(controlflowshape.GongMarshallField(stage, "IsPresentationOverridden"))
		initializerStatements.WriteString(controlflowshape.GongMarshallField(stage, "Stroke"))
		initializerStatements.WriteString(controlflowshape.GongMarshallField(stage, "StrokeWidth"))
		initializerStatements.WriteString(controlflowshape.GongMarshallField(stage, "StrokeOpacity"))
		initializerStatements.WriteString(controlflowshape.GongMarshallField(stage, "StrokeDashArray"))
		initializerStatements.WriteString(controlflowshape.GongMarshallField(stage, "FontSize"))
		initializerStatements.WriteString(controlflowshape.GongMarshallField(stage, "FontColor"))
	}
	initRes = initializerStatements.String()
	ptrRes = pointersInitializesStatements.String()
	return
}
func (data *Data) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(data.GongMarshallField(stage, "Name"))
		initializerStatements.WriteString(data.GongMarshallField(stage, "Acronym"))
		initializerStatements.WriteString(data.GongMarshallField(stage, "Description"))
		initializerStatements.WriteString(data.GongMarshallField(stage, "ComputedPrefix"))
		initializerStatements.WriteString(data.GongMarshallField(stage, "IsExpanded"))
		initializerStatements.WriteString(data.GongMarshallField(stage, "SVG_Path"))
		initializerStatements.WriteString(data.GongMarshallField(stage, "InverseAppliedScaling"))
	}
	initRes = initializerStatements.String()
	ptrRes = pointersInitializesStatements.String()
	return
}
func (dataflow *DataFlow) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(dataflow.GongMarshallField(stage, "Name"))
		pointersInitializesStatements.WriteString(dataflow.GongMarshallField(stage, "StartPort"))
		pointersInitializesStatements.WriteString(dataflow.GongMarshallField(stage, "EndPort"))
		pointersInitializesStatements.WriteString(dataflow.GongMarshallField(stage, "StartExternalPart"))
		pointersInitializesStatements.WriteString(dataflow.GongMarshallField(stage, "EndExternalPart"))
		pointersInitializesStatements.WriteString(dataflow.GongMarshallField(stage, "Datas"))
		initializerStatements.WriteString(dataflow.GongMarshallField(stage, "Description"))
		initializerStatements.WriteString(dataflow.GongMarshallField(stage, "Direction"))
		initializerStatements.WriteString(dataflow.GongMarshallField(stage, "IsDatasNodeExpanded"))
		initializerStatements.WriteString(dataflow.GongMarshallField(stage, "ComputedPrefix"))
		initializerStatements.WriteString(dataflow.GongMarshallField(stage, "IsExpanded"))
		initializerStatements.WriteString(dataflow.GongMarshallField(stage, "Type"))
	}
	initRes = initializerStatements.String()
	ptrRes = pointersInitializesStatements.String()
	return
}
func (dataflowshape *DataFlowShape) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(dataflowshape.GongMarshallField(stage, "Name"))
		pointersInitializesStatements.WriteString(dataflowshape.GongMarshallField(stage, "DataFlow"))
		initializerStatements.WriteString(dataflowshape.GongMarshallField(stage, "StartRatio"))
		initializerStatements.WriteString(dataflowshape.GongMarshallField(stage, "EndRatio"))
		initializerStatements.WriteString(dataflowshape.GongMarshallField(stage, "StartOrientation"))
		initializerStatements.WriteString(dataflowshape.GongMarshallField(stage, "EndOrientation"))
		initializerStatements.WriteString(dataflowshape.GongMarshallField(stage, "CornerOffsetRatio"))
		initializerStatements.WriteString(dataflowshape.GongMarshallField(stage, "IsHidden"))
		initializerStatements.WriteString(dataflowshape.GongMarshallField(stage, "IsPresentationOverridden"))
		initializerStatements.WriteString(dataflowshape.GongMarshallField(stage, "Stroke"))
		initializerStatements.WriteString(dataflowshape.GongMarshallField(stage, "StrokeWidth"))
		initializerStatements.WriteString(dataflowshape.GongMarshallField(stage, "StrokeOpacity"))
		initializerStatements.WriteString(dataflowshape.GongMarshallField(stage, "StrokeDashArray"))
		initializerStatements.WriteString(dataflowshape.GongMarshallField(stage, "FontSize"))
		initializerStatements.WriteString(dataflowshape.GongMarshallField(stage, "FontColor"))
	}
	initRes = initializerStatements.String()
	ptrRes = pointersInitializesStatements.String()
	return
}
func (datashape *DataShape) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(datashape.GongMarshallField(stage, "Name"))
		pointersInitializesStatements.WriteString(datashape.GongMarshallField(stage, "Data"))
		pointersInitializesStatements.WriteString(datashape.GongMarshallField(stage, "DataFlow"))
	}
	initRes = initializerStatements.String()
	ptrRes = pointersInitializesStatements.String()
	return
}
func (diagramlayerstate *DiagramLayerState) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(diagramlayerstate.GongMarshallField(stage, "Name"))
		pointersInitializesStatements.WriteString(diagramlayerstate.GongMarshallField(stage, "DiagramStructure"))
		pointersInitializesStatements.WriteString(diagramlayerstate.GongMarshallField(stage, "LayerDefinition"))
	}
	initRes = initializerStatements.String()
	ptrRes = pointersInitializesStatements.String()
	return
}
func (diagramstructure *DiagramStructure) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(diagramstructure.GongMarshallField(stage, "Name"))
		initializerStatements.WriteString(diagramstructure.GongMarshallField(stage, "Description"))
		initializerStatements.WriteString(diagramstructure.GongMarshallField(stage, "ComputedPrefix"))
		initializerStatements.WriteString(diagramstructure.GongMarshallField(stage, "IsExpanded"))
		initializerStatements.WriteString(diagramstructure.GongMarshallField(stage, "IsChecked"))
		initializerStatements.WriteString(diagramstructure.GongMarshallField(stage, "IsEditable_"))
		initializerStatements.WriteString(diagramstructure.GongMarshallField(stage, "IsShowPrefix"))
		initializerStatements.WriteString(diagramstructure.GongMarshallField(stage, "DefaultBoxWidth"))
		initializerStatements.WriteString(diagramstructure.GongMarshallField(stage, "DefaultBoxHeigth"))
		initializerStatements.WriteString(diagramstructure.GongMarshallField(stage, "IsWithDiscretePorts"))
		initializerStatements.WriteString(diagramstructure.GongMarshallField(stage, "Width"))
		initializerStatements.WriteString(diagramstructure.GongMarshallField(stage, "Height"))
		pointersInitializesStatements.WriteString(diagramstructure.GongMarshallField(stage, "System_Shapes"))
		initializerStatements.WriteString(diagramstructure.GongMarshallField(stage, "IsSystemsNodeExpanded"))
		pointersInitializesStatements.WriteString(diagramstructure.GongMarshallField(stage, "SystemsWhoseNodeIsExpanded"))
		pointersInitializesStatements.WriteString(diagramstructure.GongMarshallField(stage, "Part_Shapes"))
		initializerStatements.WriteString(diagramstructure.GongMarshallField(stage, "IsPartsNodeExpanded"))
		pointersInitializesStatements.WriteString(diagramstructure.GongMarshallField(stage, "PartWhoseNodeIsExpanded"))
		pointersInitializesStatements.WriteString(diagramstructure.GongMarshallField(stage, "ExternalPart_Shapes"))
		initializerStatements.WriteString(diagramstructure.GongMarshallField(stage, "IsExternalPartsNodeExpanded"))
		pointersInitializesStatements.WriteString(diagramstructure.GongMarshallField(stage, "ExternalPartWhoseNodeIsExpanded"))
		pointersInitializesStatements.WriteString(diagramstructure.GongMarshallField(stage, "ExternalPartsWhoseOutDataFlowsNodeIsExpanded"))
		pointersInitializesStatements.WriteString(diagramstructure.GongMarshallField(stage, "ExternalPartsWhoseInDataFlowsNodeIsExpanded"))
		pointersInitializesStatements.WriteString(diagramstructure.GongMarshallField(stage, "PortsWhoseNodeIsExpanded"))
		pointersInitializesStatements.WriteString(diagramstructure.GongMarshallField(stage, "Port_Shapes"))
		pointersInitializesStatements.WriteString(diagramstructure.GongMarshallField(stage, "ControlFlowsWhoseNodeIsExpanded"))
		pointersInitializesStatements.WriteString(diagramstructure.GongMarshallField(stage, "ControlFlow_Shapes"))
		pointersInitializesStatements.WriteString(diagramstructure.GongMarshallField(stage, "DataFlowsWhoseNodeIsExpanded"))
		pointersInitializesStatements.WriteString(diagramstructure.GongMarshallField(stage, "DataFlow_Shapes"))
		pointersInitializesStatements.WriteString(diagramstructure.GongMarshallField(stage, "DatasWhoseNodeIsExpanded"))
		pointersInitializesStatements.WriteString(diagramstructure.GongMarshallField(stage, "Data_Shapes"))
		pointersInitializesStatements.WriteString(diagramstructure.GongMarshallField(stage, "DataFlowsWhoseDataNodeIsExpanded"))
		pointersInitializesStatements.WriteString(diagramstructure.GongMarshallField(stage, "AllocatedResourcesWhoseNodeIsExpanded"))
		pointersInitializesStatements.WriteString(diagramstructure.GongMarshallField(stage, "AllocatedResourceShapes"))
		pointersInitializesStatements.WriteString(diagramstructure.GongMarshallField(stage, "AllocatedSystemesWhoseNodeIsExpanded"))
		pointersInitializesStatements.WriteString(diagramstructure.GongMarshallField(stage, "AllocatedSystemShapes"))
		pointersInitializesStatements.WriteString(diagramstructure.GongMarshallField(stage, "Note_Shapes"))
		pointersInitializesStatements.WriteString(diagramstructure.GongMarshallField(stage, "NotesWhoseNodeIsExpanded"))
		initializerStatements.WriteString(diagramstructure.GongMarshallField(stage, "IsNotesNodeExpanded"))
		pointersInitializesStatements.WriteString(diagramstructure.GongMarshallField(stage, "NotePortShapes"))
		pointersInitializesStatements.WriteString(diagramstructure.GongMarshallField(stage, "NotePartShapes"))
	}
	initRes = initializerStatements.String()
	ptrRes = pointersInitializesStatements.String()
	return
}
func (externalpartshape *ExternalPartShape) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(externalpartshape.GongMarshallField(stage, "Name"))
		pointersInitializesStatements.WriteString(externalpartshape.GongMarshallField(stage, "Part"))
		initializerStatements.WriteString(externalpartshape.GongMarshallField(stage, "IsExpanded"))
		initializerStatements.WriteString(externalpartshape.GongMarshallField(stage, "X"))
		initializerStatements.WriteString(externalpartshape.GongMarshallField(stage, "Y"))
		initializerStatements.WriteString(externalpartshape.GongMarshallField(stage, "Width"))
		initializerStatements.WriteString(externalpartshape.GongMarshallField(stage, "Height"))
		initializerStatements.WriteString(externalpartshape.GongMarshallField(stage, "IsHidden"))
		initializerStatements.WriteString(externalpartshape.GongMarshallField(stage, "IsPresentationOverridden"))
		initializerStatements.WriteString(externalpartshape.GongMarshallField(stage, "Color"))
		initializerStatements.WriteString(externalpartshape.GongMarshallField(stage, "FillOpacity"))
		initializerStatements.WriteString(externalpartshape.GongMarshallField(stage, "Stroke"))
		initializerStatements.WriteString(externalpartshape.GongMarshallField(stage, "StrokeWidth"))
		initializerStatements.WriteString(externalpartshape.GongMarshallField(stage, "StrokeOpacity"))
		initializerStatements.WriteString(externalpartshape.GongMarshallField(stage, "StrokeDashArray"))
		initializerStatements.WriteString(externalpartshape.GongMarshallField(stage, "RX"))
		initializerStatements.WriteString(externalpartshape.GongMarshallField(stage, "FontSize"))
		initializerStatements.WriteString(externalpartshape.GongMarshallField(stage, "FontWeight"))
		initializerStatements.WriteString(externalpartshape.GongMarshallField(stage, "FontColor"))
		initializerStatements.WriteString(externalpartshape.GongMarshallField(stage, "TailHeigth"))
	}
	initRes = initializerStatements.String()
	ptrRes = pointersInitializesStatements.String()
	return
}
func (layerdefinition *LayerDefinition) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(layerdefinition.GongMarshallField(stage, "Name"))
		pointersInitializesStatements.WriteString(layerdefinition.GongMarshallField(stage, "Query"))
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
		initializerStatements.WriteString(library.GongMarshallField(stage, "Description"))
		initializerStatements.WriteString(library.GongMarshallField(stage, "ComputedPrefix"))
		initializerStatements.WriteString(library.GongMarshallField(stage, "IsExpanded"))
		initializerStatements.WriteString(library.GongMarshallField(stage, "IsRootLibrary"))
		pointersInitializesStatements.WriteString(library.GongMarshallField(stage, "SubLibraries"))
		initializerStatements.WriteString(library.GongMarshallField(stage, "IsSubLibrariesNodeExpanded"))
		pointersInitializesStatements.WriteString(library.GongMarshallField(stage, "SubLibrariesWhoseNodeIsExpanded"))
		initializerStatements.WriteString(library.GongMarshallField(stage, "NbPixPerCharacter"))
		initializerStatements.WriteString(library.GongMarshallField(stage, "LogoSVGFile"))
		pointersInitializesStatements.WriteString(library.GongMarshallField(stage, "RootSystemes"))
		initializerStatements.WriteString(library.GongMarshallField(stage, "IsSystemesNodeExpanded"))
		pointersInitializesStatements.WriteString(library.GongMarshallField(stage, "SystemsWhoseNodeIsExpanded"))
		pointersInitializesStatements.WriteString(library.GongMarshallField(stage, "RootDataFlows"))
		initializerStatements.WriteString(library.GongMarshallField(stage, "IsDataFlowsNodeExpanded"))
		pointersInitializesStatements.WriteString(library.GongMarshallField(stage, "DataFlowsWhoseNodeIsExpanded"))
		pointersInitializesStatements.WriteString(library.GongMarshallField(stage, "RootDatas"))
		initializerStatements.WriteString(library.GongMarshallField(stage, "IsDatasNodeExpanded"))
		pointersInitializesStatements.WriteString(library.GongMarshallField(stage, "DatasWhoseNodeIsExpanded"))
		pointersInitializesStatements.WriteString(library.GongMarshallField(stage, "RootResources"))
		initializerStatements.WriteString(library.GongMarshallField(stage, "IsResourcesNodeExpanded"))
		pointersInitializesStatements.WriteString(library.GongMarshallField(stage, "ResourcesWhoseNodeIsExpanded"))
		pointersInitializesStatements.WriteString(library.GongMarshallField(stage, "PartsWhoseNodeIsExpanded"))
		pointersInitializesStatements.WriteString(library.GongMarshallField(stage, "RootNotes"))
		initializerStatements.WriteString(library.GongMarshallField(stage, "IsNotesNodeExpanded"))
		pointersInitializesStatements.WriteString(library.GongMarshallField(stage, "NotesWhoseNodeIsExpanded"))
		initializerStatements.WriteString(library.GongMarshallField(stage, "IsExpandedTmp"))
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
		initializerStatements.WriteString(note.GongMarshallField(stage, "Description"))
		initializerStatements.WriteString(note.GongMarshallField(stage, "ComputedPrefix"))
		initializerStatements.WriteString(note.GongMarshallField(stage, "IsExpanded"))
		initializerStatements.WriteString(note.GongMarshallField(stage, "IsPartsNodeExpanded"))
		pointersInitializesStatements.WriteString(note.GongMarshallField(stage, "Parts"))
		initializerStatements.WriteString(note.GongMarshallField(stage, "IsPortsNodeExpanded"))
		pointersInitializesStatements.WriteString(note.GongMarshallField(stage, "Ports"))
	}
	initRes = initializerStatements.String()
	ptrRes = pointersInitializesStatements.String()
	return
}
func (notepartshape *NotePartShape) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(notepartshape.GongMarshallField(stage, "Name"))
		pointersInitializesStatements.WriteString(notepartshape.GongMarshallField(stage, "Note"))
		pointersInitializesStatements.WriteString(notepartshape.GongMarshallField(stage, "Part"))
		initializerStatements.WriteString(notepartshape.GongMarshallField(stage, "StartRatio"))
		initializerStatements.WriteString(notepartshape.GongMarshallField(stage, "EndRatio"))
		initializerStatements.WriteString(notepartshape.GongMarshallField(stage, "StartOrientation"))
		initializerStatements.WriteString(notepartshape.GongMarshallField(stage, "EndOrientation"))
		initializerStatements.WriteString(notepartshape.GongMarshallField(stage, "CornerOffsetRatio"))
		initializerStatements.WriteString(notepartshape.GongMarshallField(stage, "IsHidden"))
		initializerStatements.WriteString(notepartshape.GongMarshallField(stage, "IsPresentationOverridden"))
		initializerStatements.WriteString(notepartshape.GongMarshallField(stage, "Stroke"))
		initializerStatements.WriteString(notepartshape.GongMarshallField(stage, "StrokeWidth"))
		initializerStatements.WriteString(notepartshape.GongMarshallField(stage, "StrokeOpacity"))
		initializerStatements.WriteString(notepartshape.GongMarshallField(stage, "StrokeDashArray"))
		initializerStatements.WriteString(notepartshape.GongMarshallField(stage, "FontSize"))
		initializerStatements.WriteString(notepartshape.GongMarshallField(stage, "FontColor"))
	}
	initRes = initializerStatements.String()
	ptrRes = pointersInitializesStatements.String()
	return
}
func (noteportshape *NotePortShape) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(noteportshape.GongMarshallField(stage, "Name"))
		pointersInitializesStatements.WriteString(noteportshape.GongMarshallField(stage, "Note"))
		pointersInitializesStatements.WriteString(noteportshape.GongMarshallField(stage, "Port"))
		initializerStatements.WriteString(noteportshape.GongMarshallField(stage, "StartRatio"))
		initializerStatements.WriteString(noteportshape.GongMarshallField(stage, "EndRatio"))
		initializerStatements.WriteString(noteportshape.GongMarshallField(stage, "StartOrientation"))
		initializerStatements.WriteString(noteportshape.GongMarshallField(stage, "EndOrientation"))
		initializerStatements.WriteString(noteportshape.GongMarshallField(stage, "CornerOffsetRatio"))
		initializerStatements.WriteString(noteportshape.GongMarshallField(stage, "IsHidden"))
		initializerStatements.WriteString(noteportshape.GongMarshallField(stage, "IsPresentationOverridden"))
		initializerStatements.WriteString(noteportshape.GongMarshallField(stage, "Stroke"))
		initializerStatements.WriteString(noteportshape.GongMarshallField(stage, "StrokeWidth"))
		initializerStatements.WriteString(noteportshape.GongMarshallField(stage, "StrokeOpacity"))
		initializerStatements.WriteString(noteportshape.GongMarshallField(stage, "StrokeDashArray"))
		initializerStatements.WriteString(noteportshape.GongMarshallField(stage, "FontSize"))
		initializerStatements.WriteString(noteportshape.GongMarshallField(stage, "FontColor"))
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
func (part *Part) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(part.GongMarshallField(stage, "Name"))
		initializerStatements.WriteString(part.GongMarshallField(stage, "Description"))
		pointersInitializesStatements.WriteString(part.GongMarshallField(stage, "Ports"))
		pointersInitializesStatements.WriteString(part.GongMarshallField(stage, "TypeOfPart"))
		initializerStatements.WriteString(part.GongMarshallField(stage, "IsPartNameNotSystemName"))
		initializerStatements.WriteString(part.GongMarshallField(stage, "IsControlFlowsNodeExpanded"))
		pointersInitializesStatements.WriteString(part.GongMarshallField(stage, "ControlFlows"))
		pointersInitializesStatements.WriteString(part.GongMarshallField(stage, "PortWhoseOutControlFlowsNodeIsExpanded"))
		pointersInitializesStatements.WriteString(part.GongMarshallField(stage, "PortWhoseInControlFlowsNodeIsExpanded"))
		initializerStatements.WriteString(part.GongMarshallField(stage, "IsDataFlowsNodeExpanded"))
		pointersInitializesStatements.WriteString(part.GongMarshallField(stage, "PortWhoseOutDataFlowsNodeIsExpanded"))
		pointersInitializesStatements.WriteString(part.GongMarshallField(stage, "PortWhoseInDataFlowsNodeIsExpanded"))
		pointersInitializesStatements.WriteString(part.GongMarshallField(stage, "PartAnchoredPath"))
		initializerStatements.WriteString(part.GongMarshallField(stage, "ComputedPrefix"))
		initializerStatements.WriteString(part.GongMarshallField(stage, "IsExpanded"))
		initializerStatements.WriteString(part.GongMarshallField(stage, "IsPortsNodeExpanded"))
	}
	initRes = initializerStatements.String()
	ptrRes = pointersInitializesStatements.String()
	return
}
func (partanchoredpath *PartAnchoredPath) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(partanchoredpath.GongMarshallField(stage, "Name"))
		initializerStatements.WriteString(partanchoredpath.GongMarshallField(stage, "Definition"))
		initializerStatements.WriteString(partanchoredpath.GongMarshallField(stage, "X_Offset"))
		initializerStatements.WriteString(partanchoredpath.GongMarshallField(stage, "Y_Offset"))
		initializerStatements.WriteString(partanchoredpath.GongMarshallField(stage, "RectAnchorType"))
		initializerStatements.WriteString(partanchoredpath.GongMarshallField(stage, "ScalePropotionnally"))
		initializerStatements.WriteString(partanchoredpath.GongMarshallField(stage, "AppliedScaling"))
		initializerStatements.WriteString(partanchoredpath.GongMarshallField(stage, "Color"))
		initializerStatements.WriteString(partanchoredpath.GongMarshallField(stage, "FillOpacity"))
		initializerStatements.WriteString(partanchoredpath.GongMarshallField(stage, "Stroke"))
		initializerStatements.WriteString(partanchoredpath.GongMarshallField(stage, "StrokeOpacity"))
		initializerStatements.WriteString(partanchoredpath.GongMarshallField(stage, "StrokeWidth"))
		initializerStatements.WriteString(partanchoredpath.GongMarshallField(stage, "StrokeDashArray"))
		initializerStatements.WriteString(partanchoredpath.GongMarshallField(stage, "StrokeDashArrayWhenSelected"))
		initializerStatements.WriteString(partanchoredpath.GongMarshallField(stage, "Transform"))
	}
	initRes = initializerStatements.String()
	ptrRes = pointersInitializesStatements.String()
	return
}
func (partshape *PartShape) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(partshape.GongMarshallField(stage, "Name"))
		pointersInitializesStatements.WriteString(partshape.GongMarshallField(stage, "Part"))
		initializerStatements.WriteString(partshape.GongMarshallField(stage, "IsExpanded"))
		initializerStatements.WriteString(partshape.GongMarshallField(stage, "X"))
		initializerStatements.WriteString(partshape.GongMarshallField(stage, "Y"))
		initializerStatements.WriteString(partshape.GongMarshallField(stage, "Width"))
		initializerStatements.WriteString(partshape.GongMarshallField(stage, "Height"))
		initializerStatements.WriteString(partshape.GongMarshallField(stage, "IsHidden"))
		initializerStatements.WriteString(partshape.GongMarshallField(stage, "IsPresentationOverridden"))
		initializerStatements.WriteString(partshape.GongMarshallField(stage, "Color"))
		initializerStatements.WriteString(partshape.GongMarshallField(stage, "FillOpacity"))
		initializerStatements.WriteString(partshape.GongMarshallField(stage, "Stroke"))
		initializerStatements.WriteString(partshape.GongMarshallField(stage, "StrokeWidth"))
		initializerStatements.WriteString(partshape.GongMarshallField(stage, "StrokeOpacity"))
		initializerStatements.WriteString(partshape.GongMarshallField(stage, "StrokeDashArray"))
		initializerStatements.WriteString(partshape.GongMarshallField(stage, "RX"))
		initializerStatements.WriteString(partshape.GongMarshallField(stage, "FontSize"))
		initializerStatements.WriteString(partshape.GongMarshallField(stage, "FontWeight"))
		initializerStatements.WriteString(partshape.GongMarshallField(stage, "FontColor"))
	}
	initRes = initializerStatements.String()
	ptrRes = pointersInitializesStatements.String()
	return
}
func (port *Port) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(port.GongMarshallField(stage, "Name"))
		initializerStatements.WriteString(port.GongMarshallField(stage, "Description"))
		initializerStatements.WriteString(port.GongMarshallField(stage, "ComputedPrefix"))
		initializerStatements.WriteString(port.GongMarshallField(stage, "IsExpanded"))
	}
	initRes = initializerStatements.String()
	ptrRes = pointersInitializesStatements.String()
	return
}
func (portshape *PortShape) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(portshape.GongMarshallField(stage, "Name"))
		pointersInitializesStatements.WriteString(portshape.GongMarshallField(stage, "Port"))
		initializerStatements.WriteString(portshape.GongMarshallField(stage, "IsExpanded"))
		initializerStatements.WriteString(portshape.GongMarshallField(stage, "X"))
		initializerStatements.WriteString(portshape.GongMarshallField(stage, "Y"))
		initializerStatements.WriteString(portshape.GongMarshallField(stage, "Width"))
		initializerStatements.WriteString(portshape.GongMarshallField(stage, "Height"))
		initializerStatements.WriteString(portshape.GongMarshallField(stage, "IsHidden"))
		initializerStatements.WriteString(portshape.GongMarshallField(stage, "IsPresentationOverridden"))
		initializerStatements.WriteString(portshape.GongMarshallField(stage, "Color"))
		initializerStatements.WriteString(portshape.GongMarshallField(stage, "FillOpacity"))
		initializerStatements.WriteString(portshape.GongMarshallField(stage, "Stroke"))
		initializerStatements.WriteString(portshape.GongMarshallField(stage, "StrokeWidth"))
		initializerStatements.WriteString(portshape.GongMarshallField(stage, "StrokeOpacity"))
		initializerStatements.WriteString(portshape.GongMarshallField(stage, "StrokeDashArray"))
		initializerStatements.WriteString(portshape.GongMarshallField(stage, "RX"))
		initializerStatements.WriteString(portshape.GongMarshallField(stage, "FontSize"))
		initializerStatements.WriteString(portshape.GongMarshallField(stage, "FontWeight"))
		initializerStatements.WriteString(portshape.GongMarshallField(stage, "FontColor"))
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
		initializerStatements.WriteString(resource.GongMarshallField(stage, "Acronym"))
		initializerStatements.WriteString(resource.GongMarshallField(stage, "Description"))
		initializerStatements.WriteString(resource.GongMarshallField(stage, "ComputedPrefix"))
		initializerStatements.WriteString(resource.GongMarshallField(stage, "IsExpanded"))
		initializerStatements.WriteString(resource.GongMarshallField(stage, "SVG_Path"))
		initializerStatements.WriteString(resource.GongMarshallField(stage, "InverseAppliedScaling"))
	}
	initRes = initializerStatements.String()
	ptrRes = pointersInitializesStatements.String()
	return
}
func (semantictag *SemanticTag) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(semantictag.GongMarshallField(stage, "Name"))
		pointersInitializesStatements.WriteString(semantictag.GongMarshallField(stage, "Parts"))
	}
	initRes = initializerStatements.String()
	ptrRes = pointersInitializesStatements.String()
	return
}
func (system *System) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(system.GongMarshallField(stage, "Name"))
		initializerStatements.WriteString(system.GongMarshallField(stage, "Description"))
		initializerStatements.WriteString(system.GongMarshallField(stage, "ComputedPrefix"))
		initializerStatements.WriteString(system.GongMarshallField(stage, "IsExpanded"))
		initializerStatements.WriteString(system.GongMarshallField(stage, "SVG_Path"))
		initializerStatements.WriteString(system.GongMarshallField(stage, "InverseAppliedScaling"))
		pointersInitializesStatements.WriteString(system.GongMarshallField(stage, "DiagramStructures"))
		pointersInitializesStatements.WriteString(system.GongMarshallField(stage, "DiagramStructureWhoseNodeIsExpanded"))
		initializerStatements.WriteString(system.GongMarshallField(stage, "IsSubSystemNodeExpanded"))
		pointersInitializesStatements.WriteString(system.GongMarshallField(stage, "SubSystemes"))
		pointersInitializesStatements.WriteString(system.GongMarshallField(stage, "Parts"))
		pointersInitializesStatements.WriteString(system.GongMarshallField(stage, "PartWhoseNodeIsExpanded"))
		pointersInitializesStatements.WriteString(system.GongMarshallField(stage, "DataFlows"))
		initializerStatements.WriteString(system.GongMarshallField(stage, "IsDataFlowsNodeExpanded"))
		pointersInitializesStatements.WriteString(system.GongMarshallField(stage, "ExternalParts"))
		pointersInitializesStatements.WriteString(system.GongMarshallField(stage, "ExternalPartWhoseNodeIsExpanded"))
	}
	initRes = initializerStatements.String()
	ptrRes = pointersInitializesStatements.String()
	return
}
func (systemshape *SystemShape) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(systemshape.GongMarshallField(stage, "Name"))
		pointersInitializesStatements.WriteString(systemshape.GongMarshallField(stage, "System"))
		initializerStatements.WriteString(systemshape.GongMarshallField(stage, "IsExpanded"))
		initializerStatements.WriteString(systemshape.GongMarshallField(stage, "X"))
		initializerStatements.WriteString(systemshape.GongMarshallField(stage, "Y"))
		initializerStatements.WriteString(systemshape.GongMarshallField(stage, "Width"))
		initializerStatements.WriteString(systemshape.GongMarshallField(stage, "Height"))
		initializerStatements.WriteString(systemshape.GongMarshallField(stage, "IsHidden"))
		initializerStatements.WriteString(systemshape.GongMarshallField(stage, "IsPresentationOverridden"))
		initializerStatements.WriteString(systemshape.GongMarshallField(stage, "Color"))
		initializerStatements.WriteString(systemshape.GongMarshallField(stage, "FillOpacity"))
		initializerStatements.WriteString(systemshape.GongMarshallField(stage, "Stroke"))
		initializerStatements.WriteString(systemshape.GongMarshallField(stage, "StrokeWidth"))
		initializerStatements.WriteString(systemshape.GongMarshallField(stage, "StrokeOpacity"))
		initializerStatements.WriteString(systemshape.GongMarshallField(stage, "StrokeDashArray"))
		initializerStatements.WriteString(systemshape.GongMarshallField(stage, "RX"))
		initializerStatements.WriteString(systemshape.GongMarshallField(stage, "FontSize"))
		initializerStatements.WriteString(systemshape.GongMarshallField(stage, "FontWeight"))
		initializerStatements.WriteString(systemshape.GongMarshallField(stage, "FontColor"))
	}
	initRes = initializerStatements.String()
	ptrRes = pointersInitializesStatements.String()
	return
}
