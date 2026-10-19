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
	modelsPackageName := "github.com/fullstack-lang/gong/dsm/process/go/models"
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
		for _, allocatedprocessshape := range __gong__sortStageSetInstances(stage.AllocatedProcessShapes, stage.AllocatedProcessShape_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			allocatedprocessshapeIdent := localAliasPrefix + allocatedprocessshape.GongGetIdentifier(stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&models.AllocatedProcessShape{Name: %s}).Stage(stage)", allocatedprocessshapeIdent, __gong__toRawStringLiteral(allocatedprocessshape.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", allocatedprocessshapeIdent, __gong__toRawStringLiteral(allocatedprocessshape.Name)))
			if allocatedprocessshape.Participant != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := localAliasPrefix + allocatedprocessshape.Participant.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Participant = %s", allocatedprocessshapeIdent, targetIdent))
			}
			if allocatedprocessshape.Process != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := localAliasPrefix + allocatedprocessshape.Process.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Process = %s", allocatedprocessshapeIdent, targetIdent))
			}
		}
	}
	if stage != nil {
		for _, allocatedresourceshape := range __gong__sortStageSetInstances(stage.AllocatedResourceShapes, stage.AllocatedResourceShape_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			allocatedresourceshapeIdent := localAliasPrefix + allocatedresourceshape.GongGetIdentifier(stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&models.AllocatedResourceShape{Name: %s}).Stage(stage)", allocatedresourceshapeIdent, __gong__toRawStringLiteral(allocatedresourceshape.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", allocatedresourceshapeIdent, __gong__toRawStringLiteral(allocatedresourceshape.Name)))
			if allocatedresourceshape.Participant != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := localAliasPrefix + allocatedresourceshape.Participant.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Participant = %s", allocatedresourceshapeIdent, targetIdent))
			}
			if allocatedresourceshape.Resource != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := localAliasPrefix + allocatedresourceshape.Resource.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Resource = %s", allocatedresourceshapeIdent, targetIdent))
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
			controlflowIdent := localAliasPrefix + controlflow.GongGetIdentifier(stage)
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
				targetIdent := localAliasPrefix + controlflow.Start.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Start = %s", controlflowIdent, targetIdent))
			}
			if controlflow.End != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := localAliasPrefix + controlflow.End.GongGetIdentifier(stage)
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
			controlflowshapeIdent := localAliasPrefix + controlflowshape.GongGetIdentifier(stage)
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
				targetIdent := localAliasPrefix + controlflowshape.ControlFlow.GongGetIdentifier(stage)
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
			dataIdent := localAliasPrefix + data.GongGetIdentifier(stage)
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
			dataflowIdent := localAliasPrefix + dataflow.GongGetIdentifier(stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&models.DataFlow{Name: %s}).Stage(stage)", dataflowIdent, __gong__toRawStringLiteral(dataflow.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", dataflowIdent, __gong__toRawStringLiteral(dataflow.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Description = %s", dataflowIdent, __gong__toRawStringLiteral(dataflow.Description)))
			values.WriteString(fmt.Sprintf("\n\t%s.ComputedPrefix = %s", dataflowIdent, __gong__toRawStringLiteral(dataflow.ComputedPrefix)))
			values.WriteString(fmt.Sprintf("\n\t%s.IsExpanded = %t", dataflowIdent, dataflow.IsExpanded))
			if dataflow.Type.ToCodeString() != "" {
				values.WriteString(fmt.Sprintf("\n\t%s.Type = models.%s", dataflowIdent, dataflow.Type.ToCodeString()))
			} else {
				values.WriteString(fmt.Sprintf("\n\t%s.Type = \"\"", dataflowIdent))
			}
			values.WriteString(fmt.Sprintf("\n\t%s.IsDatasNodeExpanded = %t", dataflowIdent, dataflow.IsDatasNodeExpanded))
			for _, elem := range dataflow.Datas {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := localAliasPrefix + elem.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Datas = append(%s.Datas, %s)", dataflowIdent, dataflowIdent, targetIdent))
			}
			if dataflow.StartTask != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := localAliasPrefix + dataflow.StartTask.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.StartTask = %s", dataflowIdent, targetIdent))
			}
			if dataflow.EndTask != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := localAliasPrefix + dataflow.EndTask.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.EndTask = %s", dataflowIdent, targetIdent))
			}
			if dataflow.StartExternalParticipant != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := localAliasPrefix + dataflow.StartExternalParticipant.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.StartExternalParticipant = %s", dataflowIdent, targetIdent))
			}
			if dataflow.EndExternalParticipant != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := localAliasPrefix + dataflow.EndExternalParticipant.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.EndExternalParticipant = %s", dataflowIdent, targetIdent))
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
			dataflowshapeIdent := localAliasPrefix + dataflowshape.GongGetIdentifier(stage)
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
				targetIdent := localAliasPrefix + dataflowshape.DataFlow.GongGetIdentifier(stage)
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
			datashapeIdent := localAliasPrefix + datashape.GongGetIdentifier(stage)
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
				targetIdent := localAliasPrefix + datashape.Data.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Data = %s", datashapeIdent, targetIdent))
			}
			if datashape.DataFlow != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := localAliasPrefix + datashape.DataFlow.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.DataFlow = %s", datashapeIdent, targetIdent))
			}
		}
	}
	if stage != nil {
		for _, diagramprocess := range __gong__sortStageSetInstances(stage.DiagramProcesss, stage.DiagramProcess_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			diagramprocessIdent := localAliasPrefix + diagramprocess.GongGetIdentifier(stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&models.DiagramProcess{Name: %s}).Stage(stage)", diagramprocessIdent, __gong__toRawStringLiteral(diagramprocess.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", diagramprocessIdent, __gong__toRawStringLiteral(diagramprocess.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Description = %s", diagramprocessIdent, __gong__toRawStringLiteral(diagramprocess.Description)))
			values.WriteString(fmt.Sprintf("\n\t%s.ComputedPrefix = %s", diagramprocessIdent, __gong__toRawStringLiteral(diagramprocess.ComputedPrefix)))
			values.WriteString(fmt.Sprintf("\n\t%s.IsExpanded = %t", diagramprocessIdent, diagramprocess.IsExpanded))
			values.WriteString(fmt.Sprintf("\n\t%s.IsChecked = %t", diagramprocessIdent, diagramprocess.IsChecked))
			values.WriteString(fmt.Sprintf("\n\t%s.IsEditable_ = %t", diagramprocessIdent, diagramprocess.IsEditable_))
			values.WriteString(fmt.Sprintf("\n\t%s.IsShowPrefix = %t", diagramprocessIdent, diagramprocess.IsShowPrefix))
			values.WriteString(fmt.Sprintf("\n\t%s.DefaultBoxWidth = %f", diagramprocessIdent, diagramprocess.DefaultBoxWidth))
			values.WriteString(fmt.Sprintf("\n\t%s.DefaultBoxHeigth = %f", diagramprocessIdent, diagramprocess.DefaultBoxHeigth))
			values.WriteString(fmt.Sprintf("\n\t%s.Width = %f", diagramprocessIdent, diagramprocess.Width))
			values.WriteString(fmt.Sprintf("\n\t%s.Height = %f", diagramprocessIdent, diagramprocess.Height))
			values.WriteString(fmt.Sprintf("\n\t%s.IsProcesssNodeExpanded = %t", diagramprocessIdent, diagramprocess.IsProcesssNodeExpanded))
			values.WriteString(fmt.Sprintf("\n\t%s.IsParticipantsNodeExpanded = %t", diagramprocessIdent, diagramprocess.IsParticipantsNodeExpanded))
			values.WriteString(fmt.Sprintf("\n\t%s.IsExternalParticipantsNodeExpanded = %t", diagramprocessIdent, diagramprocess.IsExternalParticipantsNodeExpanded))
			values.WriteString(fmt.Sprintf("\n\t%s.IsNotesNodeExpanded = %t", diagramprocessIdent, diagramprocess.IsNotesNodeExpanded))
			for _, elem := range diagramprocess.Process_Shapes {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := localAliasPrefix + elem.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Process_Shapes = append(%s.Process_Shapes, %s)", diagramprocessIdent, diagramprocessIdent, targetIdent))
			}
			for _, elem := range diagramprocess.ProcesssWhoseNodeIsExpanded {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := localAliasPrefix + elem.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.ProcesssWhoseNodeIsExpanded = append(%s.ProcesssWhoseNodeIsExpanded, %s)", diagramprocessIdent, diagramprocessIdent, targetIdent))
			}
			for _, elem := range diagramprocess.Participant_Shapes {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := localAliasPrefix + elem.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Participant_Shapes = append(%s.Participant_Shapes, %s)", diagramprocessIdent, diagramprocessIdent, targetIdent))
			}
			for _, elem := range diagramprocess.ParticipantWhoseNodeIsExpanded {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := localAliasPrefix + elem.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.ParticipantWhoseNodeIsExpanded = append(%s.ParticipantWhoseNodeIsExpanded, %s)", diagramprocessIdent, diagramprocessIdent, targetIdent))
			}
			for _, elem := range diagramprocess.ExternalParticipant_Shapes {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := localAliasPrefix + elem.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.ExternalParticipant_Shapes = append(%s.ExternalParticipant_Shapes, %s)", diagramprocessIdent, diagramprocessIdent, targetIdent))
			}
			for _, elem := range diagramprocess.ExternalParticipantWhoseNodeIsExpanded {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := localAliasPrefix + elem.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.ExternalParticipantWhoseNodeIsExpanded = append(%s.ExternalParticipantWhoseNodeIsExpanded, %s)", diagramprocessIdent, diagramprocessIdent, targetIdent))
			}
			for _, elem := range diagramprocess.ExternalParticipantsWhoseOutDataFlowsNodeIsExpanded {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := localAliasPrefix + elem.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.ExternalParticipantsWhoseOutDataFlowsNodeIsExpanded = append(%s.ExternalParticipantsWhoseOutDataFlowsNodeIsExpanded, %s)", diagramprocessIdent, diagramprocessIdent, targetIdent))
			}
			for _, elem := range diagramprocess.ExternalParticipantsWhoseInDataFlowsNodeIsExpanded {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := localAliasPrefix + elem.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.ExternalParticipantsWhoseInDataFlowsNodeIsExpanded = append(%s.ExternalParticipantsWhoseInDataFlowsNodeIsExpanded, %s)", diagramprocessIdent, diagramprocessIdent, targetIdent))
			}
			for _, elem := range diagramprocess.TasksWhoseNodeIsExpanded {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := localAliasPrefix + elem.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.TasksWhoseNodeIsExpanded = append(%s.TasksWhoseNodeIsExpanded, %s)", diagramprocessIdent, diagramprocessIdent, targetIdent))
			}
			for _, elem := range diagramprocess.Task_Shapes {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := localAliasPrefix + elem.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Task_Shapes = append(%s.Task_Shapes, %s)", diagramprocessIdent, diagramprocessIdent, targetIdent))
			}
			for _, elem := range diagramprocess.ControlFlowsWhoseNodeIsExpanded {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := localAliasPrefix + elem.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.ControlFlowsWhoseNodeIsExpanded = append(%s.ControlFlowsWhoseNodeIsExpanded, %s)", diagramprocessIdent, diagramprocessIdent, targetIdent))
			}
			for _, elem := range diagramprocess.ControlFlow_Shapes {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := localAliasPrefix + elem.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.ControlFlow_Shapes = append(%s.ControlFlow_Shapes, %s)", diagramprocessIdent, diagramprocessIdent, targetIdent))
			}
			for _, elem := range diagramprocess.DataFlowsWhoseNodeIsExpanded {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := localAliasPrefix + elem.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.DataFlowsWhoseNodeIsExpanded = append(%s.DataFlowsWhoseNodeIsExpanded, %s)", diagramprocessIdent, diagramprocessIdent, targetIdent))
			}
			for _, elem := range diagramprocess.DataFlow_Shapes {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := localAliasPrefix + elem.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.DataFlow_Shapes = append(%s.DataFlow_Shapes, %s)", diagramprocessIdent, diagramprocessIdent, targetIdent))
			}
			for _, elem := range diagramprocess.DatasWhoseNodeIsExpanded {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := localAliasPrefix + elem.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.DatasWhoseNodeIsExpanded = append(%s.DatasWhoseNodeIsExpanded, %s)", diagramprocessIdent, diagramprocessIdent, targetIdent))
			}
			for _, elem := range diagramprocess.Data_Shapes {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := localAliasPrefix + elem.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Data_Shapes = append(%s.Data_Shapes, %s)", diagramprocessIdent, diagramprocessIdent, targetIdent))
			}
			for _, elem := range diagramprocess.DataFlowsWhoseDataNodeIsExpanded {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := localAliasPrefix + elem.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.DataFlowsWhoseDataNodeIsExpanded = append(%s.DataFlowsWhoseDataNodeIsExpanded, %s)", diagramprocessIdent, diagramprocessIdent, targetIdent))
			}
			for _, elem := range diagramprocess.AllocatedResourcesWhoseNodeIsExpanded {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := localAliasPrefix + elem.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.AllocatedResourcesWhoseNodeIsExpanded = append(%s.AllocatedResourcesWhoseNodeIsExpanded, %s)", diagramprocessIdent, diagramprocessIdent, targetIdent))
			}
			for _, elem := range diagramprocess.AllocatedResourceShapes {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := localAliasPrefix + elem.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.AllocatedResourceShapes = append(%s.AllocatedResourceShapes, %s)", diagramprocessIdent, diagramprocessIdent, targetIdent))
			}
			for _, elem := range diagramprocess.AllocatedProcessesWhoseNodeIsExpanded {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := localAliasPrefix + elem.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.AllocatedProcessesWhoseNodeIsExpanded = append(%s.AllocatedProcessesWhoseNodeIsExpanded, %s)", diagramprocessIdent, diagramprocessIdent, targetIdent))
			}
			for _, elem := range diagramprocess.AllocatedProcessShapes {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := localAliasPrefix + elem.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.AllocatedProcessShapes = append(%s.AllocatedProcessShapes, %s)", diagramprocessIdent, diagramprocessIdent, targetIdent))
			}
			for _, elem := range diagramprocess.Note_Shapes {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := localAliasPrefix + elem.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Note_Shapes = append(%s.Note_Shapes, %s)", diagramprocessIdent, diagramprocessIdent, targetIdent))
			}
			for _, elem := range diagramprocess.NotesWhoseNodeIsExpanded {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := localAliasPrefix + elem.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.NotesWhoseNodeIsExpanded = append(%s.NotesWhoseNodeIsExpanded, %s)", diagramprocessIdent, diagramprocessIdent, targetIdent))
			}
			for _, elem := range diagramprocess.NoteTaskShapes {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := localAliasPrefix + elem.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.NoteTaskShapes = append(%s.NoteTaskShapes, %s)", diagramprocessIdent, diagramprocessIdent, targetIdent))
			}
		}
	}
	if stage != nil {
		for _, externalparticipantshape := range __gong__sortStageSetInstances(stage.ExternalParticipantShapes, stage.ExternalParticipantShape_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			externalparticipantshapeIdent := localAliasPrefix + externalparticipantshape.GongGetIdentifier(stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&models.ExternalParticipantShape{Name: %s}).Stage(stage)", externalparticipantshapeIdent, __gong__toRawStringLiteral(externalparticipantshape.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", externalparticipantshapeIdent, __gong__toRawStringLiteral(externalparticipantshape.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.IsExpanded = %t", externalparticipantshapeIdent, externalparticipantshape.IsExpanded))
			values.WriteString(fmt.Sprintf("\n\t%s.X = %f", externalparticipantshapeIdent, externalparticipantshape.X))
			values.WriteString(fmt.Sprintf("\n\t%s.Y = %f", externalparticipantshapeIdent, externalparticipantshape.Y))
			values.WriteString(fmt.Sprintf("\n\t%s.Width = %f", externalparticipantshapeIdent, externalparticipantshape.Width))
			values.WriteString(fmt.Sprintf("\n\t%s.Height = %f", externalparticipantshapeIdent, externalparticipantshape.Height))
			values.WriteString(fmt.Sprintf("\n\t%s.IsHidden = %t", externalparticipantshapeIdent, externalparticipantshape.IsHidden))
			values.WriteString(fmt.Sprintf("\n\t%s.IsPresentationOverridden = %t", externalparticipantshapeIdent, externalparticipantshape.IsPresentationOverridden))
			values.WriteString(fmt.Sprintf("\n\t%s.Color = %s", externalparticipantshapeIdent, __gong__toRawStringLiteral(externalparticipantshape.Color)))
			values.WriteString(fmt.Sprintf("\n\t%s.FillOpacity = %f", externalparticipantshapeIdent, externalparticipantshape.FillOpacity))
			values.WriteString(fmt.Sprintf("\n\t%s.Stroke = %s", externalparticipantshapeIdent, __gong__toRawStringLiteral(externalparticipantshape.Stroke)))
			values.WriteString(fmt.Sprintf("\n\t%s.StrokeWidth = %f", externalparticipantshapeIdent, externalparticipantshape.StrokeWidth))
			values.WriteString(fmt.Sprintf("\n\t%s.StrokeOpacity = %f", externalparticipantshapeIdent, externalparticipantshape.StrokeOpacity))
			values.WriteString(fmt.Sprintf("\n\t%s.StrokeDashArray = %s", externalparticipantshapeIdent, __gong__toRawStringLiteral(externalparticipantshape.StrokeDashArray)))
			values.WriteString(fmt.Sprintf("\n\t%s.RX = %f", externalparticipantshapeIdent, externalparticipantshape.RX))
			values.WriteString(fmt.Sprintf("\n\t%s.FontSize = %s", externalparticipantshapeIdent, __gong__toRawStringLiteral(externalparticipantshape.FontSize)))
			values.WriteString(fmt.Sprintf("\n\t%s.FontWeight = %s", externalparticipantshapeIdent, __gong__toRawStringLiteral(externalparticipantshape.FontWeight)))
			values.WriteString(fmt.Sprintf("\n\t%s.FontColor = %s", externalparticipantshapeIdent, __gong__toRawStringLiteral(externalparticipantshape.FontColor)))
			values.WriteString(fmt.Sprintf("\n\t%s.TailHeigth = %f", externalparticipantshapeIdent, externalparticipantshape.TailHeigth))
			if externalparticipantshape.Participant != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := localAliasPrefix + externalparticipantshape.Participant.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Participant = %s", externalparticipantshapeIdent, targetIdent))
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
			values.WriteString(fmt.Sprintf("\n\t%s.Description = %s", libraryIdent, __gong__toRawStringLiteral(library.Description)))
			values.WriteString(fmt.Sprintf("\n\t%s.ComputedPrefix = %s", libraryIdent, __gong__toRawStringLiteral(library.ComputedPrefix)))
			values.WriteString(fmt.Sprintf("\n\t%s.IsExpanded = %t", libraryIdent, library.IsExpanded))
			values.WriteString(fmt.Sprintf("\n\t%s.IsRootLibrary = %t", libraryIdent, library.IsRootLibrary))
			values.WriteString(fmt.Sprintf("\n\t%s.IsSubLibrariesNodeExpanded = %t", libraryIdent, library.IsSubLibrariesNodeExpanded))
			values.WriteString(fmt.Sprintf("\n\t%s.NbPixPerCharacter = %f", libraryIdent, library.NbPixPerCharacter))
			values.WriteString(fmt.Sprintf("\n\t%s.LogoSVGFile = %s", libraryIdent, __gong__toRawStringLiteral(library.LogoSVGFile)))
			values.WriteString(fmt.Sprintf("\n\t%s.IsProcessesNodeExpanded = %t", libraryIdent, library.IsProcessesNodeExpanded))
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
				targetIdent := localAliasPrefix + elem.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.SubLibraries = append(%s.SubLibraries, %s)", libraryIdent, libraryIdent, targetIdent))
			}
			for _, elem := range library.SubLibrariesWhoseNodeIsExpanded {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := localAliasPrefix + elem.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.SubLibrariesWhoseNodeIsExpanded = append(%s.SubLibrariesWhoseNodeIsExpanded, %s)", libraryIdent, libraryIdent, targetIdent))
			}
			for _, elem := range library.RootProcesses {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := localAliasPrefix + elem.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.RootProcesses = append(%s.RootProcesses, %s)", libraryIdent, libraryIdent, targetIdent))
			}
			for _, elem := range library.ProcesssWhoseNodeIsExpanded {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := localAliasPrefix + elem.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.ProcesssWhoseNodeIsExpanded = append(%s.ProcesssWhoseNodeIsExpanded, %s)", libraryIdent, libraryIdent, targetIdent))
			}
			for _, elem := range library.RootDataFlows {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := localAliasPrefix + elem.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.RootDataFlows = append(%s.RootDataFlows, %s)", libraryIdent, libraryIdent, targetIdent))
			}
			for _, elem := range library.DataFlowsWhoseNodeIsExpanded {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := localAliasPrefix + elem.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.DataFlowsWhoseNodeIsExpanded = append(%s.DataFlowsWhoseNodeIsExpanded, %s)", libraryIdent, libraryIdent, targetIdent))
			}
			for _, elem := range library.RootDatas {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := localAliasPrefix + elem.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.RootDatas = append(%s.RootDatas, %s)", libraryIdent, libraryIdent, targetIdent))
			}
			for _, elem := range library.DatasWhoseNodeIsExpanded {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := localAliasPrefix + elem.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.DatasWhoseNodeIsExpanded = append(%s.DatasWhoseNodeIsExpanded, %s)", libraryIdent, libraryIdent, targetIdent))
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
			for _, elem := range library.ResourcesWhoseNodeIsExpanded {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := localAliasPrefix + elem.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.ResourcesWhoseNodeIsExpanded = append(%s.ResourcesWhoseNodeIsExpanded, %s)", libraryIdent, libraryIdent, targetIdent))
			}
			for _, elem := range library.ParticipantsWhoseNodeIsExpanded {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := localAliasPrefix + elem.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.ParticipantsWhoseNodeIsExpanded = append(%s.ParticipantsWhoseNodeIsExpanded, %s)", libraryIdent, libraryIdent, targetIdent))
			}
			for _, elem := range library.RootNotes {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := localAliasPrefix + elem.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.RootNotes = append(%s.RootNotes, %s)", libraryIdent, libraryIdent, targetIdent))
			}
			for _, elem := range library.NotesWhoseNodeIsExpanded {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := localAliasPrefix + elem.GongGetIdentifier(stage)
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
			noteIdent := localAliasPrefix + note.GongGetIdentifier(stage)
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
			values.WriteString(fmt.Sprintf("\n\t%s.IsTasksNodeExpanded = %t", noteIdent, note.IsTasksNodeExpanded))
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
		for _, participant := range __gong__sortStageSetInstances(stage.Participants, stage.Participant_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			participantIdent := localAliasPrefix + participant.GongGetIdentifier(stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&models.Participant{Name: %s}).Stage(stage)", participantIdent, __gong__toRawStringLiteral(participant.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", participantIdent, __gong__toRawStringLiteral(participant.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.IsProcessResource = %t", participantIdent, participant.IsProcessResource))
			values.WriteString(fmt.Sprintf("\n\t%s.Description = %s", participantIdent, __gong__toRawStringLiteral(participant.Description)))
			values.WriteString(fmt.Sprintf("\n\t%s.IsResourcesNodeExpanded = %t", participantIdent, participant.IsResourcesNodeExpanded))
			values.WriteString(fmt.Sprintf("\n\t%s.IsProcessesNodeExpanded = %t", participantIdent, participant.IsProcessesNodeExpanded))
			values.WriteString(fmt.Sprintf("\n\t%s.ComputedPrefix = %s", participantIdent, __gong__toRawStringLiteral(participant.ComputedPrefix)))
			values.WriteString(fmt.Sprintf("\n\t%s.IsExpanded = %t", participantIdent, participant.IsExpanded))
			values.WriteString(fmt.Sprintf("\n\t%s.IsTasksNodeExpanded = %t", participantIdent, participant.IsTasksNodeExpanded))
			values.WriteString(fmt.Sprintf("\n\t%s.IsControlFlowsNodeExpanded = %t", participantIdent, participant.IsControlFlowsNodeExpanded))
			values.WriteString(fmt.Sprintf("\n\t%s.IsDataFlowsNodeExpanded = %t", participantIdent, participant.IsDataFlowsNodeExpanded))
			for _, elem := range participant.Resources {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := localAliasPrefix + elem.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Resources = append(%s.Resources, %s)", participantIdent, participantIdent, targetIdent))
			}
			for _, elem := range participant.Processes {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := localAliasPrefix + elem.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Processes = append(%s.Processes, %s)", participantIdent, participantIdent, targetIdent))
			}
			for _, elem := range participant.Tasks {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := localAliasPrefix + elem.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Tasks = append(%s.Tasks, %s)", participantIdent, participantIdent, targetIdent))
			}
			for _, elem := range participant.ControlFlows {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := localAliasPrefix + elem.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.ControlFlows = append(%s.ControlFlows, %s)", participantIdent, participantIdent, targetIdent))
			}
			for _, elem := range participant.TaskWhoseOutControlFlowsNodeIsExpanded {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := localAliasPrefix + elem.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.TaskWhoseOutControlFlowsNodeIsExpanded = append(%s.TaskWhoseOutControlFlowsNodeIsExpanded, %s)", participantIdent, participantIdent, targetIdent))
			}
			for _, elem := range participant.TaskWhoseInControlFlowsNodeIsExpanded {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := localAliasPrefix + elem.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.TaskWhoseInControlFlowsNodeIsExpanded = append(%s.TaskWhoseInControlFlowsNodeIsExpanded, %s)", participantIdent, participantIdent, targetIdent))
			}
			for _, elem := range participant.TaskWhoseOutDataFlowsNodeIsExpanded {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := localAliasPrefix + elem.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.TaskWhoseOutDataFlowsNodeIsExpanded = append(%s.TaskWhoseOutDataFlowsNodeIsExpanded, %s)", participantIdent, participantIdent, targetIdent))
			}
			for _, elem := range participant.TaskWhoseInDataFlowsNodeIsExpanded {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := localAliasPrefix + elem.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.TaskWhoseInDataFlowsNodeIsExpanded = append(%s.TaskWhoseInDataFlowsNodeIsExpanded, %s)", participantIdent, participantIdent, targetIdent))
			}
		}
	}
	if stage != nil {
		for _, participantshape := range __gong__sortStageSetInstances(stage.ParticipantShapes, stage.ParticipantShape_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			participantshapeIdent := localAliasPrefix + participantshape.GongGetIdentifier(stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&models.ParticipantShape{Name: %s}).Stage(stage)", participantshapeIdent, __gong__toRawStringLiteral(participantshape.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", participantshapeIdent, __gong__toRawStringLiteral(participantshape.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.IsExpanded = %t", participantshapeIdent, participantshape.IsExpanded))
			values.WriteString(fmt.Sprintf("\n\t%s.X = %f", participantshapeIdent, participantshape.X))
			values.WriteString(fmt.Sprintf("\n\t%s.Y = %f", participantshapeIdent, participantshape.Y))
			values.WriteString(fmt.Sprintf("\n\t%s.Width = %f", participantshapeIdent, participantshape.Width))
			values.WriteString(fmt.Sprintf("\n\t%s.Height = %f", participantshapeIdent, participantshape.Height))
			values.WriteString(fmt.Sprintf("\n\t%s.IsHidden = %t", participantshapeIdent, participantshape.IsHidden))
			values.WriteString(fmt.Sprintf("\n\t%s.IsPresentationOverridden = %t", participantshapeIdent, participantshape.IsPresentationOverridden))
			values.WriteString(fmt.Sprintf("\n\t%s.Color = %s", participantshapeIdent, __gong__toRawStringLiteral(participantshape.Color)))
			values.WriteString(fmt.Sprintf("\n\t%s.FillOpacity = %f", participantshapeIdent, participantshape.FillOpacity))
			values.WriteString(fmt.Sprintf("\n\t%s.Stroke = %s", participantshapeIdent, __gong__toRawStringLiteral(participantshape.Stroke)))
			values.WriteString(fmt.Sprintf("\n\t%s.StrokeWidth = %f", participantshapeIdent, participantshape.StrokeWidth))
			values.WriteString(fmt.Sprintf("\n\t%s.StrokeOpacity = %f", participantshapeIdent, participantshape.StrokeOpacity))
			values.WriteString(fmt.Sprintf("\n\t%s.StrokeDashArray = %s", participantshapeIdent, __gong__toRawStringLiteral(participantshape.StrokeDashArray)))
			values.WriteString(fmt.Sprintf("\n\t%s.RX = %f", participantshapeIdent, participantshape.RX))
			values.WriteString(fmt.Sprintf("\n\t%s.FontSize = %s", participantshapeIdent, __gong__toRawStringLiteral(participantshape.FontSize)))
			values.WriteString(fmt.Sprintf("\n\t%s.FontWeight = %s", participantshapeIdent, __gong__toRawStringLiteral(participantshape.FontWeight)))
			values.WriteString(fmt.Sprintf("\n\t%s.FontColor = %s", participantshapeIdent, __gong__toRawStringLiteral(participantshape.FontColor)))
			values.WriteString(fmt.Sprintf("\n\t%s.WidthWeight = %f", participantshapeIdent, participantshape.WidthWeight))
			if participantshape.Participant != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := localAliasPrefix + participantshape.Participant.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Participant = %s", participantshapeIdent, targetIdent))
			}
		}
	}
	if stage != nil {
		for _, process := range __gong__sortStageSetInstances(stage.Processs, stage.Process_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			processIdent := localAliasPrefix + process.GongGetIdentifier(stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&models.Process{Name: %s}).Stage(stage)", processIdent, __gong__toRawStringLiteral(process.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", processIdent, __gong__toRawStringLiteral(process.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Description = %s", processIdent, __gong__toRawStringLiteral(process.Description)))
			values.WriteString(fmt.Sprintf("\n\t%s.ComputedPrefix = %s", processIdent, __gong__toRawStringLiteral(process.ComputedPrefix)))
			values.WriteString(fmt.Sprintf("\n\t%s.IsExpanded = %t", processIdent, process.IsExpanded))
			values.WriteString(fmt.Sprintf("\n\t%s.SVG_Path = %s", processIdent, __gong__toRawStringLiteral(process.SVG_Path)))
			values.WriteString(fmt.Sprintf("\n\t%s.InverseAppliedScaling = %f", processIdent, process.InverseAppliedScaling))
			values.WriteString(fmt.Sprintf("\n\t%s.IsSubProcessNodeExpanded = %t", processIdent, process.IsSubProcessNodeExpanded))
			values.WriteString(fmt.Sprintf("\n\t%s.IsDataFlowsNodeExpanded = %t", processIdent, process.IsDataFlowsNodeExpanded))
			for _, elem := range process.DiagramProcesss {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := localAliasPrefix + elem.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.DiagramProcesss = append(%s.DiagramProcesss, %s)", processIdent, processIdent, targetIdent))
			}
			for _, elem := range process.DiagramProcessWhoseNodeIsExpanded {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := localAliasPrefix + elem.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.DiagramProcessWhoseNodeIsExpanded = append(%s.DiagramProcessWhoseNodeIsExpanded, %s)", processIdent, processIdent, targetIdent))
			}
			for _, elem := range process.SubProcesses {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := localAliasPrefix + elem.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.SubProcesses = append(%s.SubProcesses, %s)", processIdent, processIdent, targetIdent))
			}
			for _, elem := range process.Participants {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := localAliasPrefix + elem.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Participants = append(%s.Participants, %s)", processIdent, processIdent, targetIdent))
			}
			for _, elem := range process.ParticipantWhoseNodeIsExpanded {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := localAliasPrefix + elem.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.ParticipantWhoseNodeIsExpanded = append(%s.ParticipantWhoseNodeIsExpanded, %s)", processIdent, processIdent, targetIdent))
			}
			for _, elem := range process.DataFlows {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := localAliasPrefix + elem.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.DataFlows = append(%s.DataFlows, %s)", processIdent, processIdent, targetIdent))
			}
			for _, elem := range process.ExternalParticipants {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := localAliasPrefix + elem.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.ExternalParticipants = append(%s.ExternalParticipants, %s)", processIdent, processIdent, targetIdent))
			}
			for _, elem := range process.ExternalParticipantWhoseNodeIsExpanded {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := localAliasPrefix + elem.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.ExternalParticipantWhoseNodeIsExpanded = append(%s.ExternalParticipantWhoseNodeIsExpanded, %s)", processIdent, processIdent, targetIdent))
			}
		}
	}
	if stage != nil {
		for _, processshape := range __gong__sortStageSetInstances(stage.ProcessShapes, stage.ProcessShape_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			processshapeIdent := localAliasPrefix + processshape.GongGetIdentifier(stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&models.ProcessShape{Name: %s}).Stage(stage)", processshapeIdent, __gong__toRawStringLiteral(processshape.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", processshapeIdent, __gong__toRawStringLiteral(processshape.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.IsExpanded = %t", processshapeIdent, processshape.IsExpanded))
			values.WriteString(fmt.Sprintf("\n\t%s.X = %f", processshapeIdent, processshape.X))
			values.WriteString(fmt.Sprintf("\n\t%s.Y = %f", processshapeIdent, processshape.Y))
			values.WriteString(fmt.Sprintf("\n\t%s.Width = %f", processshapeIdent, processshape.Width))
			values.WriteString(fmt.Sprintf("\n\t%s.Height = %f", processshapeIdent, processshape.Height))
			values.WriteString(fmt.Sprintf("\n\t%s.IsHidden = %t", processshapeIdent, processshape.IsHidden))
			values.WriteString(fmt.Sprintf("\n\t%s.IsPresentationOverridden = %t", processshapeIdent, processshape.IsPresentationOverridden))
			values.WriteString(fmt.Sprintf("\n\t%s.Color = %s", processshapeIdent, __gong__toRawStringLiteral(processshape.Color)))
			values.WriteString(fmt.Sprintf("\n\t%s.FillOpacity = %f", processshapeIdent, processshape.FillOpacity))
			values.WriteString(fmt.Sprintf("\n\t%s.Stroke = %s", processshapeIdent, __gong__toRawStringLiteral(processshape.Stroke)))
			values.WriteString(fmt.Sprintf("\n\t%s.StrokeWidth = %f", processshapeIdent, processshape.StrokeWidth))
			values.WriteString(fmt.Sprintf("\n\t%s.StrokeOpacity = %f", processshapeIdent, processshape.StrokeOpacity))
			values.WriteString(fmt.Sprintf("\n\t%s.StrokeDashArray = %s", processshapeIdent, __gong__toRawStringLiteral(processshape.StrokeDashArray)))
			values.WriteString(fmt.Sprintf("\n\t%s.RX = %f", processshapeIdent, processshape.RX))
			values.WriteString(fmt.Sprintf("\n\t%s.FontSize = %s", processshapeIdent, __gong__toRawStringLiteral(processshape.FontSize)))
			values.WriteString(fmt.Sprintf("\n\t%s.FontWeight = %s", processshapeIdent, __gong__toRawStringLiteral(processshape.FontWeight)))
			values.WriteString(fmt.Sprintf("\n\t%s.FontColor = %s", processshapeIdent, __gong__toRawStringLiteral(processshape.FontColor)))
			if processshape.Process != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := localAliasPrefix + processshape.Process.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Process = %s", processshapeIdent, targetIdent))
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
			values.WriteString(fmt.Sprintf("\n\t%s.Acronym = %s", resourceIdent, __gong__toRawStringLiteral(resource.Acronym)))
			values.WriteString(fmt.Sprintf("\n\t%s.Description = %s", resourceIdent, __gong__toRawStringLiteral(resource.Description)))
			values.WriteString(fmt.Sprintf("\n\t%s.ComputedPrefix = %s", resourceIdent, __gong__toRawStringLiteral(resource.ComputedPrefix)))
			values.WriteString(fmt.Sprintf("\n\t%s.IsExpanded = %t", resourceIdent, resource.IsExpanded))
			values.WriteString(fmt.Sprintf("\n\t%s.SVG_Path = %s", resourceIdent, __gong__toRawStringLiteral(resource.SVG_Path)))
			values.WriteString(fmt.Sprintf("\n\t%s.InverseAppliedScaling = %f", resourceIdent, resource.InverseAppliedScaling))
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
			values.WriteString(fmt.Sprintf("\n\t%s.ComputedPrefix = %s", taskIdent, __gong__toRawStringLiteral(task.ComputedPrefix)))
			values.WriteString(fmt.Sprintf("\n\t%s.IsExpanded = %t", taskIdent, task.IsExpanded))
			values.WriteString(fmt.Sprintf("\n\t%s.IsStartTask = %t", taskIdent, task.IsStartTask))
			values.WriteString(fmt.Sprintf("\n\t%s.IsEndTask = %t", taskIdent, task.IsEndTask))
			values.WriteString(fmt.Sprintf("\n\t%s.IsTaskNameNotProcessName = %t", taskIdent, task.IsTaskNameNotProcessName))
			if task.Type != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := localAliasPrefix + task.Type.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Type = %s", taskIdent, targetIdent))
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
			values.WriteString(fmt.Sprintf("\n\t%s.IsExpanded = %t", taskshapeIdent, taskshape.IsExpanded))
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
go run github.com/fullstack-lang/gong/dsm/process/go/models/probe/cmd/probe@%s stage.go
*/
package %s

import (
	"slices"
	"time"

	"github.com/fullstack-lang/gong/dsm/process/go/models"%s
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

	if modelsPackageName != "" && modelsPackageName != "github.com/fullstack-lang/gong/dsm/process/go/models" {
		res = strings.ReplaceAll(res, "\"github.com/fullstack-lang/gong/dsm/process/go/models\"", fmt.Sprintf("%q", modelsPackageName))
	}

	return res, nil
}

// insertion point for marshall field methods
func (allocatedprocessshape *AllocatedProcessShape) GongMarshallField(stage *Stage, fieldName string) (res string) {
	ident := allocatedprocessshape.GongGetIdentifier(stage)
	_ = ident
	switch fieldName {
	case "Name":
		res = __gong__marshallString(ident, "Name", allocatedprocessshape.Name)

	case "Participant":
		if allocatedprocessshape.Participant != nil {
			res = __gong__marshallPointer(ident, "Participant", allocatedprocessshape.Participant.GongGetIdentifier(stage))
		} else {
			res = __gong__marshallPointer(ident, "Participant", "nil")
		}
	case "Process":
		if allocatedprocessshape.Process != nil {
			res = __gong__marshallPointer(ident, "Process", allocatedprocessshape.Process.GongGetIdentifier(stage))
		} else {
			res = __gong__marshallPointer(ident, "Process", "nil")
		}
	default:
		log.Panicf("Unknown field %s for Gongstruct AllocatedProcessShape", fieldName)
	}
	return
}

func (allocatedresourceshape *AllocatedResourceShape) GongMarshallField(stage *Stage, fieldName string) (res string) {
	ident := allocatedresourceshape.GongGetIdentifier(stage)
	_ = ident
	switch fieldName {
	case "Name":
		res = __gong__marshallString(ident, "Name", allocatedresourceshape.Name)

	case "Participant":
		if allocatedresourceshape.Participant != nil {
			res = __gong__marshallPointer(ident, "Participant", allocatedresourceshape.Participant.GongGetIdentifier(stage))
		} else {
			res = __gong__marshallPointer(ident, "Participant", "nil")
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
	case "ComputedPrefix":
		res = __gong__marshallString(ident, "ComputedPrefix", dataflow.ComputedPrefix)
	case "IsExpanded":
		res = __gong__marshallBool(ident, "IsExpanded", dataflow.IsExpanded)
	case "Type":
		res = __gong__marshallEnumString(ident, "Type", dataflow.Type.ToCodeString())
	case "IsDatasNodeExpanded":
		res = __gong__marshallBool(ident, "IsDatasNodeExpanded", dataflow.IsDatasNodeExpanded)

	case "Datas":
		var sb strings.Builder
		for _, _data := range dataflow.Datas {
			sb.WriteString(__gong__marshallSliceOfPointers(ident, "Datas", _data.GongGetIdentifier(stage)))
		}
		res = sb.String()
	case "StartTask":
		if dataflow.StartTask != nil {
			res = __gong__marshallPointer(ident, "StartTask", dataflow.StartTask.GongGetIdentifier(stage))
		} else {
			res = __gong__marshallPointer(ident, "StartTask", "nil")
		}
	case "EndTask":
		if dataflow.EndTask != nil {
			res = __gong__marshallPointer(ident, "EndTask", dataflow.EndTask.GongGetIdentifier(stage))
		} else {
			res = __gong__marshallPointer(ident, "EndTask", "nil")
		}
	case "StartExternalParticipant":
		if dataflow.StartExternalParticipant != nil {
			res = __gong__marshallPointer(ident, "StartExternalParticipant", dataflow.StartExternalParticipant.GongGetIdentifier(stage))
		} else {
			res = __gong__marshallPointer(ident, "StartExternalParticipant", "nil")
		}
	case "EndExternalParticipant":
		if dataflow.EndExternalParticipant != nil {
			res = __gong__marshallPointer(ident, "EndExternalParticipant", dataflow.EndExternalParticipant.GongGetIdentifier(stage))
		} else {
			res = __gong__marshallPointer(ident, "EndExternalParticipant", "nil")
		}
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

func (diagramprocess *DiagramProcess) GongMarshallField(stage *Stage, fieldName string) (res string) {
	ident := diagramprocess.GongGetIdentifier(stage)
	_ = ident
	switch fieldName {
	case "Name":
		res = __gong__marshallString(ident, "Name", diagramprocess.Name)
	case "Description":
		res = __gong__marshallString(ident, "Description", diagramprocess.Description)
	case "ComputedPrefix":
		res = __gong__marshallString(ident, "ComputedPrefix", diagramprocess.ComputedPrefix)
	case "IsExpanded":
		res = __gong__marshallBool(ident, "IsExpanded", diagramprocess.IsExpanded)
	case "IsChecked":
		res = __gong__marshallBool(ident, "IsChecked", diagramprocess.IsChecked)
	case "IsEditable_":
		res = __gong__marshallBool(ident, "IsEditable_", diagramprocess.IsEditable_)
	case "IsShowPrefix":
		res = __gong__marshallBool(ident, "IsShowPrefix", diagramprocess.IsShowPrefix)
	case "DefaultBoxWidth":
		res = __gong__marshallFloat(ident, "DefaultBoxWidth", diagramprocess.DefaultBoxWidth)
	case "DefaultBoxHeigth":
		res = __gong__marshallFloat(ident, "DefaultBoxHeigth", diagramprocess.DefaultBoxHeigth)
	case "Width":
		res = __gong__marshallFloat(ident, "Width", diagramprocess.Width)
	case "Height":
		res = __gong__marshallFloat(ident, "Height", diagramprocess.Height)
	case "IsProcesssNodeExpanded":
		res = __gong__marshallBool(ident, "IsProcesssNodeExpanded", diagramprocess.IsProcesssNodeExpanded)
	case "IsParticipantsNodeExpanded":
		res = __gong__marshallBool(ident, "IsParticipantsNodeExpanded", diagramprocess.IsParticipantsNodeExpanded)
	case "IsExternalParticipantsNodeExpanded":
		res = __gong__marshallBool(ident, "IsExternalParticipantsNodeExpanded", diagramprocess.IsExternalParticipantsNodeExpanded)
	case "IsNotesNodeExpanded":
		res = __gong__marshallBool(ident, "IsNotesNodeExpanded", diagramprocess.IsNotesNodeExpanded)

	case "Process_Shapes":
		var sb strings.Builder
		for _, _processshape := range diagramprocess.Process_Shapes {
			sb.WriteString(__gong__marshallSliceOfPointers(ident, "Process_Shapes", _processshape.GongGetIdentifier(stage)))
		}
		res = sb.String()
	case "ProcesssWhoseNodeIsExpanded":
		var sb strings.Builder
		for _, _process := range diagramprocess.ProcesssWhoseNodeIsExpanded {
			sb.WriteString(__gong__marshallSliceOfPointers(ident, "ProcesssWhoseNodeIsExpanded", _process.GongGetIdentifier(stage)))
		}
		res = sb.String()
	case "Participant_Shapes":
		var sb strings.Builder
		for _, _participantshape := range diagramprocess.Participant_Shapes {
			sb.WriteString(__gong__marshallSliceOfPointers(ident, "Participant_Shapes", _participantshape.GongGetIdentifier(stage)))
		}
		res = sb.String()
	case "ParticipantWhoseNodeIsExpanded":
		var sb strings.Builder
		for _, _participant := range diagramprocess.ParticipantWhoseNodeIsExpanded {
			sb.WriteString(__gong__marshallSliceOfPointers(ident, "ParticipantWhoseNodeIsExpanded", _participant.GongGetIdentifier(stage)))
		}
		res = sb.String()
	case "ExternalParticipant_Shapes":
		var sb strings.Builder
		for _, _externalparticipantshape := range diagramprocess.ExternalParticipant_Shapes {
			sb.WriteString(__gong__marshallSliceOfPointers(ident, "ExternalParticipant_Shapes", _externalparticipantshape.GongGetIdentifier(stage)))
		}
		res = sb.String()
	case "ExternalParticipantWhoseNodeIsExpanded":
		var sb strings.Builder
		for _, _participant := range diagramprocess.ExternalParticipantWhoseNodeIsExpanded {
			sb.WriteString(__gong__marshallSliceOfPointers(ident, "ExternalParticipantWhoseNodeIsExpanded", _participant.GongGetIdentifier(stage)))
		}
		res = sb.String()
	case "ExternalParticipantsWhoseOutDataFlowsNodeIsExpanded":
		var sb strings.Builder
		for _, _participant := range diagramprocess.ExternalParticipantsWhoseOutDataFlowsNodeIsExpanded {
			sb.WriteString(__gong__marshallSliceOfPointers(ident, "ExternalParticipantsWhoseOutDataFlowsNodeIsExpanded", _participant.GongGetIdentifier(stage)))
		}
		res = sb.String()
	case "ExternalParticipantsWhoseInDataFlowsNodeIsExpanded":
		var sb strings.Builder
		for _, _participant := range diagramprocess.ExternalParticipantsWhoseInDataFlowsNodeIsExpanded {
			sb.WriteString(__gong__marshallSliceOfPointers(ident, "ExternalParticipantsWhoseInDataFlowsNodeIsExpanded", _participant.GongGetIdentifier(stage)))
		}
		res = sb.String()
	case "TasksWhoseNodeIsExpanded":
		var sb strings.Builder
		for _, _task := range diagramprocess.TasksWhoseNodeIsExpanded {
			sb.WriteString(__gong__marshallSliceOfPointers(ident, "TasksWhoseNodeIsExpanded", _task.GongGetIdentifier(stage)))
		}
		res = sb.String()
	case "Task_Shapes":
		var sb strings.Builder
		for _, _taskshape := range diagramprocess.Task_Shapes {
			sb.WriteString(__gong__marshallSliceOfPointers(ident, "Task_Shapes", _taskshape.GongGetIdentifier(stage)))
		}
		res = sb.String()
	case "ControlFlowsWhoseNodeIsExpanded":
		var sb strings.Builder
		for _, _controlflow := range diagramprocess.ControlFlowsWhoseNodeIsExpanded {
			sb.WriteString(__gong__marshallSliceOfPointers(ident, "ControlFlowsWhoseNodeIsExpanded", _controlflow.GongGetIdentifier(stage)))
		}
		res = sb.String()
	case "ControlFlow_Shapes":
		var sb strings.Builder
		for _, _controlflowshape := range diagramprocess.ControlFlow_Shapes {
			sb.WriteString(__gong__marshallSliceOfPointers(ident, "ControlFlow_Shapes", _controlflowshape.GongGetIdentifier(stage)))
		}
		res = sb.String()
	case "DataFlowsWhoseNodeIsExpanded":
		var sb strings.Builder
		for _, _dataflow := range diagramprocess.DataFlowsWhoseNodeIsExpanded {
			sb.WriteString(__gong__marshallSliceOfPointers(ident, "DataFlowsWhoseNodeIsExpanded", _dataflow.GongGetIdentifier(stage)))
		}
		res = sb.String()
	case "DataFlow_Shapes":
		var sb strings.Builder
		for _, _dataflowshape := range diagramprocess.DataFlow_Shapes {
			sb.WriteString(__gong__marshallSliceOfPointers(ident, "DataFlow_Shapes", _dataflowshape.GongGetIdentifier(stage)))
		}
		res = sb.String()
	case "DatasWhoseNodeIsExpanded":
		var sb strings.Builder
		for _, _data := range diagramprocess.DatasWhoseNodeIsExpanded {
			sb.WriteString(__gong__marshallSliceOfPointers(ident, "DatasWhoseNodeIsExpanded", _data.GongGetIdentifier(stage)))
		}
		res = sb.String()
	case "Data_Shapes":
		var sb strings.Builder
		for _, _datashape := range diagramprocess.Data_Shapes {
			sb.WriteString(__gong__marshallSliceOfPointers(ident, "Data_Shapes", _datashape.GongGetIdentifier(stage)))
		}
		res = sb.String()
	case "DataFlowsWhoseDataNodeIsExpanded":
		var sb strings.Builder
		for _, _dataflow := range diagramprocess.DataFlowsWhoseDataNodeIsExpanded {
			sb.WriteString(__gong__marshallSliceOfPointers(ident, "DataFlowsWhoseDataNodeIsExpanded", _dataflow.GongGetIdentifier(stage)))
		}
		res = sb.String()
	case "AllocatedResourcesWhoseNodeIsExpanded":
		var sb strings.Builder
		for _, _resource := range diagramprocess.AllocatedResourcesWhoseNodeIsExpanded {
			sb.WriteString(__gong__marshallSliceOfPointers(ident, "AllocatedResourcesWhoseNodeIsExpanded", _resource.GongGetIdentifier(stage)))
		}
		res = sb.String()
	case "AllocatedResourceShapes":
		var sb strings.Builder
		for _, _allocatedresourceshape := range diagramprocess.AllocatedResourceShapes {
			sb.WriteString(__gong__marshallSliceOfPointers(ident, "AllocatedResourceShapes", _allocatedresourceshape.GongGetIdentifier(stage)))
		}
		res = sb.String()
	case "AllocatedProcessesWhoseNodeIsExpanded":
		var sb strings.Builder
		for _, _process := range diagramprocess.AllocatedProcessesWhoseNodeIsExpanded {
			sb.WriteString(__gong__marshallSliceOfPointers(ident, "AllocatedProcessesWhoseNodeIsExpanded", _process.GongGetIdentifier(stage)))
		}
		res = sb.String()
	case "AllocatedProcessShapes":
		var sb strings.Builder
		for _, _allocatedprocessshape := range diagramprocess.AllocatedProcessShapes {
			sb.WriteString(__gong__marshallSliceOfPointers(ident, "AllocatedProcessShapes", _allocatedprocessshape.GongGetIdentifier(stage)))
		}
		res = sb.String()
	case "Note_Shapes":
		var sb strings.Builder
		for _, _noteshape := range diagramprocess.Note_Shapes {
			sb.WriteString(__gong__marshallSliceOfPointers(ident, "Note_Shapes", _noteshape.GongGetIdentifier(stage)))
		}
		res = sb.String()
	case "NotesWhoseNodeIsExpanded":
		var sb strings.Builder
		for _, _note := range diagramprocess.NotesWhoseNodeIsExpanded {
			sb.WriteString(__gong__marshallSliceOfPointers(ident, "NotesWhoseNodeIsExpanded", _note.GongGetIdentifier(stage)))
		}
		res = sb.String()
	case "NoteTaskShapes":
		var sb strings.Builder
		for _, _notetaskshape := range diagramprocess.NoteTaskShapes {
			sb.WriteString(__gong__marshallSliceOfPointers(ident, "NoteTaskShapes", _notetaskshape.GongGetIdentifier(stage)))
		}
		res = sb.String()
	default:
		log.Panicf("Unknown field %s for Gongstruct DiagramProcess", fieldName)
	}
	return
}

func (externalparticipantshape *ExternalParticipantShape) GongMarshallField(stage *Stage, fieldName string) (res string) {
	ident := externalparticipantshape.GongGetIdentifier(stage)
	_ = ident
	switch fieldName {
	case "Name":
		res = __gong__marshallString(ident, "Name", externalparticipantshape.Name)
	case "IsExpanded":
		res = __gong__marshallBool(ident, "IsExpanded", externalparticipantshape.IsExpanded)
	case "X":
		res = __gong__marshallFloat(ident, "X", externalparticipantshape.X)
	case "Y":
		res = __gong__marshallFloat(ident, "Y", externalparticipantshape.Y)
	case "Width":
		res = __gong__marshallFloat(ident, "Width", externalparticipantshape.Width)
	case "Height":
		res = __gong__marshallFloat(ident, "Height", externalparticipantshape.Height)
	case "IsHidden":
		res = __gong__marshallBool(ident, "IsHidden", externalparticipantshape.IsHidden)
	case "IsPresentationOverridden":
		res = __gong__marshallBool(ident, "IsPresentationOverridden", externalparticipantshape.IsPresentationOverridden)
	case "Color":
		res = __gong__marshallString(ident, "Color", externalparticipantshape.Color)
	case "FillOpacity":
		res = __gong__marshallFloat(ident, "FillOpacity", externalparticipantshape.FillOpacity)
	case "Stroke":
		res = __gong__marshallString(ident, "Stroke", externalparticipantshape.Stroke)
	case "StrokeWidth":
		res = __gong__marshallFloat(ident, "StrokeWidth", externalparticipantshape.StrokeWidth)
	case "StrokeOpacity":
		res = __gong__marshallFloat(ident, "StrokeOpacity", externalparticipantshape.StrokeOpacity)
	case "StrokeDashArray":
		res = __gong__marshallString(ident, "StrokeDashArray", externalparticipantshape.StrokeDashArray)
	case "RX":
		res = __gong__marshallFloat(ident, "RX", externalparticipantshape.RX)
	case "FontSize":
		res = __gong__marshallString(ident, "FontSize", externalparticipantshape.FontSize)
	case "FontWeight":
		res = __gong__marshallString(ident, "FontWeight", externalparticipantshape.FontWeight)
	case "FontColor":
		res = __gong__marshallString(ident, "FontColor", externalparticipantshape.FontColor)
	case "TailHeigth":
		res = __gong__marshallFloat(ident, "TailHeigth", externalparticipantshape.TailHeigth)

	case "Participant":
		if externalparticipantshape.Participant != nil {
			res = __gong__marshallPointer(ident, "Participant", externalparticipantshape.Participant.GongGetIdentifier(stage))
		} else {
			res = __gong__marshallPointer(ident, "Participant", "nil")
		}
	default:
		log.Panicf("Unknown field %s for Gongstruct ExternalParticipantShape", fieldName)
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
	case "IsProcessesNodeExpanded":
		res = __gong__marshallBool(ident, "IsProcessesNodeExpanded", library.IsProcessesNodeExpanded)
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
	case "RootProcesses":
		var sb strings.Builder
		for _, _process := range library.RootProcesses {
			sb.WriteString(__gong__marshallSliceOfPointers(ident, "RootProcesses", _process.GongGetIdentifier(stage)))
		}
		res = sb.String()
	case "ProcesssWhoseNodeIsExpanded":
		var sb strings.Builder
		for _, _process := range library.ProcesssWhoseNodeIsExpanded {
			sb.WriteString(__gong__marshallSliceOfPointers(ident, "ProcesssWhoseNodeIsExpanded", _process.GongGetIdentifier(stage)))
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
	case "ParticipantsWhoseNodeIsExpanded":
		var sb strings.Builder
		for _, _participant := range library.ParticipantsWhoseNodeIsExpanded {
			sb.WriteString(__gong__marshallSliceOfPointers(ident, "ParticipantsWhoseNodeIsExpanded", _participant.GongGetIdentifier(stage)))
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
	case "IsTasksNodeExpanded":
		res = __gong__marshallBool(ident, "IsTasksNodeExpanded", note.IsTasksNodeExpanded)

	case "Tasks":
		var sb strings.Builder
		for _, _task := range note.Tasks {
			sb.WriteString(__gong__marshallSliceOfPointers(ident, "Tasks", _task.GongGetIdentifier(stage)))
		}
		res = sb.String()
	default:
		log.Panicf("Unknown field %s for Gongstruct Note", fieldName)
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

func (participant *Participant) GongMarshallField(stage *Stage, fieldName string) (res string) {
	ident := participant.GongGetIdentifier(stage)
	_ = ident
	switch fieldName {
	case "Name":
		res = __gong__marshallString(ident, "Name", participant.Name)
	case "IsProcessResource":
		res = __gong__marshallBool(ident, "IsProcessResource", participant.IsProcessResource)
	case "Description":
		res = __gong__marshallString(ident, "Description", participant.Description)
	case "IsResourcesNodeExpanded":
		res = __gong__marshallBool(ident, "IsResourcesNodeExpanded", participant.IsResourcesNodeExpanded)
	case "IsProcessesNodeExpanded":
		res = __gong__marshallBool(ident, "IsProcessesNodeExpanded", participant.IsProcessesNodeExpanded)
	case "ComputedPrefix":
		res = __gong__marshallString(ident, "ComputedPrefix", participant.ComputedPrefix)
	case "IsExpanded":
		res = __gong__marshallBool(ident, "IsExpanded", participant.IsExpanded)
	case "IsTasksNodeExpanded":
		res = __gong__marshallBool(ident, "IsTasksNodeExpanded", participant.IsTasksNodeExpanded)
	case "IsControlFlowsNodeExpanded":
		res = __gong__marshallBool(ident, "IsControlFlowsNodeExpanded", participant.IsControlFlowsNodeExpanded)
	case "IsDataFlowsNodeExpanded":
		res = __gong__marshallBool(ident, "IsDataFlowsNodeExpanded", participant.IsDataFlowsNodeExpanded)

	case "Resources":
		var sb strings.Builder
		for _, _resource := range participant.Resources {
			sb.WriteString(__gong__marshallSliceOfPointers(ident, "Resources", _resource.GongGetIdentifier(stage)))
		}
		res = sb.String()
	case "Processes":
		var sb strings.Builder
		for _, _process := range participant.Processes {
			sb.WriteString(__gong__marshallSliceOfPointers(ident, "Processes", _process.GongGetIdentifier(stage)))
		}
		res = sb.String()
	case "Tasks":
		var sb strings.Builder
		for _, _task := range participant.Tasks {
			sb.WriteString(__gong__marshallSliceOfPointers(ident, "Tasks", _task.GongGetIdentifier(stage)))
		}
		res = sb.String()
	case "ControlFlows":
		var sb strings.Builder
		for _, _controlflow := range participant.ControlFlows {
			sb.WriteString(__gong__marshallSliceOfPointers(ident, "ControlFlows", _controlflow.GongGetIdentifier(stage)))
		}
		res = sb.String()
	case "TaskWhoseOutControlFlowsNodeIsExpanded":
		var sb strings.Builder
		for _, _task := range participant.TaskWhoseOutControlFlowsNodeIsExpanded {
			sb.WriteString(__gong__marshallSliceOfPointers(ident, "TaskWhoseOutControlFlowsNodeIsExpanded", _task.GongGetIdentifier(stage)))
		}
		res = sb.String()
	case "TaskWhoseInControlFlowsNodeIsExpanded":
		var sb strings.Builder
		for _, _task := range participant.TaskWhoseInControlFlowsNodeIsExpanded {
			sb.WriteString(__gong__marshallSliceOfPointers(ident, "TaskWhoseInControlFlowsNodeIsExpanded", _task.GongGetIdentifier(stage)))
		}
		res = sb.String()
	case "TaskWhoseOutDataFlowsNodeIsExpanded":
		var sb strings.Builder
		for _, _task := range participant.TaskWhoseOutDataFlowsNodeIsExpanded {
			sb.WriteString(__gong__marshallSliceOfPointers(ident, "TaskWhoseOutDataFlowsNodeIsExpanded", _task.GongGetIdentifier(stage)))
		}
		res = sb.String()
	case "TaskWhoseInDataFlowsNodeIsExpanded":
		var sb strings.Builder
		for _, _task := range participant.TaskWhoseInDataFlowsNodeIsExpanded {
			sb.WriteString(__gong__marshallSliceOfPointers(ident, "TaskWhoseInDataFlowsNodeIsExpanded", _task.GongGetIdentifier(stage)))
		}
		res = sb.String()
	default:
		log.Panicf("Unknown field %s for Gongstruct Participant", fieldName)
	}
	return
}

func (participantshape *ParticipantShape) GongMarshallField(stage *Stage, fieldName string) (res string) {
	ident := participantshape.GongGetIdentifier(stage)
	_ = ident
	switch fieldName {
	case "Name":
		res = __gong__marshallString(ident, "Name", participantshape.Name)
	case "IsExpanded":
		res = __gong__marshallBool(ident, "IsExpanded", participantshape.IsExpanded)
	case "X":
		res = __gong__marshallFloat(ident, "X", participantshape.X)
	case "Y":
		res = __gong__marshallFloat(ident, "Y", participantshape.Y)
	case "Width":
		res = __gong__marshallFloat(ident, "Width", participantshape.Width)
	case "Height":
		res = __gong__marshallFloat(ident, "Height", participantshape.Height)
	case "IsHidden":
		res = __gong__marshallBool(ident, "IsHidden", participantshape.IsHidden)
	case "IsPresentationOverridden":
		res = __gong__marshallBool(ident, "IsPresentationOverridden", participantshape.IsPresentationOverridden)
	case "Color":
		res = __gong__marshallString(ident, "Color", participantshape.Color)
	case "FillOpacity":
		res = __gong__marshallFloat(ident, "FillOpacity", participantshape.FillOpacity)
	case "Stroke":
		res = __gong__marshallString(ident, "Stroke", participantshape.Stroke)
	case "StrokeWidth":
		res = __gong__marshallFloat(ident, "StrokeWidth", participantshape.StrokeWidth)
	case "StrokeOpacity":
		res = __gong__marshallFloat(ident, "StrokeOpacity", participantshape.StrokeOpacity)
	case "StrokeDashArray":
		res = __gong__marshallString(ident, "StrokeDashArray", participantshape.StrokeDashArray)
	case "RX":
		res = __gong__marshallFloat(ident, "RX", participantshape.RX)
	case "FontSize":
		res = __gong__marshallString(ident, "FontSize", participantshape.FontSize)
	case "FontWeight":
		res = __gong__marshallString(ident, "FontWeight", participantshape.FontWeight)
	case "FontColor":
		res = __gong__marshallString(ident, "FontColor", participantshape.FontColor)
	case "WidthWeight":
		res = __gong__marshallFloat(ident, "WidthWeight", participantshape.WidthWeight)

	case "Participant":
		if participantshape.Participant != nil {
			res = __gong__marshallPointer(ident, "Participant", participantshape.Participant.GongGetIdentifier(stage))
		} else {
			res = __gong__marshallPointer(ident, "Participant", "nil")
		}
	default:
		log.Panicf("Unknown field %s for Gongstruct ParticipantShape", fieldName)
	}
	return
}

func (process *Process) GongMarshallField(stage *Stage, fieldName string) (res string) {
	ident := process.GongGetIdentifier(stage)
	_ = ident
	switch fieldName {
	case "Name":
		res = __gong__marshallString(ident, "Name", process.Name)
	case "Description":
		res = __gong__marshallString(ident, "Description", process.Description)
	case "ComputedPrefix":
		res = __gong__marshallString(ident, "ComputedPrefix", process.ComputedPrefix)
	case "IsExpanded":
		res = __gong__marshallBool(ident, "IsExpanded", process.IsExpanded)
	case "SVG_Path":
		res = __gong__marshallString(ident, "SVG_Path", process.SVG_Path)
	case "InverseAppliedScaling":
		res = __gong__marshallFloat(ident, "InverseAppliedScaling", process.InverseAppliedScaling)
	case "IsSubProcessNodeExpanded":
		res = __gong__marshallBool(ident, "IsSubProcessNodeExpanded", process.IsSubProcessNodeExpanded)
	case "IsDataFlowsNodeExpanded":
		res = __gong__marshallBool(ident, "IsDataFlowsNodeExpanded", process.IsDataFlowsNodeExpanded)

	case "DiagramProcesss":
		var sb strings.Builder
		for _, _diagramprocess := range process.DiagramProcesss {
			sb.WriteString(__gong__marshallSliceOfPointers(ident, "DiagramProcesss", _diagramprocess.GongGetIdentifier(stage)))
		}
		res = sb.String()
	case "DiagramProcessWhoseNodeIsExpanded":
		var sb strings.Builder
		for _, _diagramprocess := range process.DiagramProcessWhoseNodeIsExpanded {
			sb.WriteString(__gong__marshallSliceOfPointers(ident, "DiagramProcessWhoseNodeIsExpanded", _diagramprocess.GongGetIdentifier(stage)))
		}
		res = sb.String()
	case "SubProcesses":
		var sb strings.Builder
		for _, _process := range process.SubProcesses {
			sb.WriteString(__gong__marshallSliceOfPointers(ident, "SubProcesses", _process.GongGetIdentifier(stage)))
		}
		res = sb.String()
	case "Participants":
		var sb strings.Builder
		for _, _participant := range process.Participants {
			sb.WriteString(__gong__marshallSliceOfPointers(ident, "Participants", _participant.GongGetIdentifier(stage)))
		}
		res = sb.String()
	case "ParticipantWhoseNodeIsExpanded":
		var sb strings.Builder
		for _, _participant := range process.ParticipantWhoseNodeIsExpanded {
			sb.WriteString(__gong__marshallSliceOfPointers(ident, "ParticipantWhoseNodeIsExpanded", _participant.GongGetIdentifier(stage)))
		}
		res = sb.String()
	case "DataFlows":
		var sb strings.Builder
		for _, _dataflow := range process.DataFlows {
			sb.WriteString(__gong__marshallSliceOfPointers(ident, "DataFlows", _dataflow.GongGetIdentifier(stage)))
		}
		res = sb.String()
	case "ExternalParticipants":
		var sb strings.Builder
		for _, _participant := range process.ExternalParticipants {
			sb.WriteString(__gong__marshallSliceOfPointers(ident, "ExternalParticipants", _participant.GongGetIdentifier(stage)))
		}
		res = sb.String()
	case "ExternalParticipantWhoseNodeIsExpanded":
		var sb strings.Builder
		for _, _participant := range process.ExternalParticipantWhoseNodeIsExpanded {
			sb.WriteString(__gong__marshallSliceOfPointers(ident, "ExternalParticipantWhoseNodeIsExpanded", _participant.GongGetIdentifier(stage)))
		}
		res = sb.String()
	default:
		log.Panicf("Unknown field %s for Gongstruct Process", fieldName)
	}
	return
}

func (processshape *ProcessShape) GongMarshallField(stage *Stage, fieldName string) (res string) {
	ident := processshape.GongGetIdentifier(stage)
	_ = ident
	switch fieldName {
	case "Name":
		res = __gong__marshallString(ident, "Name", processshape.Name)
	case "IsExpanded":
		res = __gong__marshallBool(ident, "IsExpanded", processshape.IsExpanded)
	case "X":
		res = __gong__marshallFloat(ident, "X", processshape.X)
	case "Y":
		res = __gong__marshallFloat(ident, "Y", processshape.Y)
	case "Width":
		res = __gong__marshallFloat(ident, "Width", processshape.Width)
	case "Height":
		res = __gong__marshallFloat(ident, "Height", processshape.Height)
	case "IsHidden":
		res = __gong__marshallBool(ident, "IsHidden", processshape.IsHidden)
	case "IsPresentationOverridden":
		res = __gong__marshallBool(ident, "IsPresentationOverridden", processshape.IsPresentationOverridden)
	case "Color":
		res = __gong__marshallString(ident, "Color", processshape.Color)
	case "FillOpacity":
		res = __gong__marshallFloat(ident, "FillOpacity", processshape.FillOpacity)
	case "Stroke":
		res = __gong__marshallString(ident, "Stroke", processshape.Stroke)
	case "StrokeWidth":
		res = __gong__marshallFloat(ident, "StrokeWidth", processshape.StrokeWidth)
	case "StrokeOpacity":
		res = __gong__marshallFloat(ident, "StrokeOpacity", processshape.StrokeOpacity)
	case "StrokeDashArray":
		res = __gong__marshallString(ident, "StrokeDashArray", processshape.StrokeDashArray)
	case "RX":
		res = __gong__marshallFloat(ident, "RX", processshape.RX)
	case "FontSize":
		res = __gong__marshallString(ident, "FontSize", processshape.FontSize)
	case "FontWeight":
		res = __gong__marshallString(ident, "FontWeight", processshape.FontWeight)
	case "FontColor":
		res = __gong__marshallString(ident, "FontColor", processshape.FontColor)

	case "Process":
		if processshape.Process != nil {
			res = __gong__marshallPointer(ident, "Process", processshape.Process.GongGetIdentifier(stage))
		} else {
			res = __gong__marshallPointer(ident, "Process", "nil")
		}
	default:
		log.Panicf("Unknown field %s for Gongstruct ProcessShape", fieldName)
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

func (task *Task) GongMarshallField(stage *Stage, fieldName string) (res string) {
	ident := task.GongGetIdentifier(stage)
	_ = ident
	switch fieldName {
	case "Name":
		res = __gong__marshallString(ident, "Name", task.Name)
	case "Description":
		res = __gong__marshallString(ident, "Description", task.Description)
	case "ComputedPrefix":
		res = __gong__marshallString(ident, "ComputedPrefix", task.ComputedPrefix)
	case "IsExpanded":
		res = __gong__marshallBool(ident, "IsExpanded", task.IsExpanded)
	case "IsStartTask":
		res = __gong__marshallBool(ident, "IsStartTask", task.IsStartTask)
	case "IsEndTask":
		res = __gong__marshallBool(ident, "IsEndTask", task.IsEndTask)
	case "IsTaskNameNotProcessName":
		res = __gong__marshallBool(ident, "IsTaskNameNotProcessName", task.IsTaskNameNotProcessName)

	case "Type":
		if task.Type != nil {
			res = __gong__marshallPointer(ident, "Type", task.Type.GongGetIdentifier(stage))
		} else {
			res = __gong__marshallPointer(ident, "Type", "nil")
		}
	default:
		log.Panicf("Unknown field %s for Gongstruct Task", fieldName)
	}
	return
}

func (taskshape *TaskShape) GongMarshallField(stage *Stage, fieldName string) (res string) {
	ident := taskshape.GongGetIdentifier(stage)
	_ = ident
	switch fieldName {
	case "Name":
		res = __gong__marshallString(ident, "Name", taskshape.Name)
	case "IsExpanded":
		res = __gong__marshallBool(ident, "IsExpanded", taskshape.IsExpanded)
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
func (allocatedprocessshape *AllocatedProcessShape) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(allocatedprocessshape.GongMarshallField(stage, "Name"))
		pointersInitializesStatements.WriteString(allocatedprocessshape.GongMarshallField(stage, "Participant"))
		pointersInitializesStatements.WriteString(allocatedprocessshape.GongMarshallField(stage, "Process"))
	}
	initRes = initializerStatements.String()
	ptrRes = pointersInitializesStatements.String()
	return
}
func (allocatedresourceshape *AllocatedResourceShape) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(allocatedresourceshape.GongMarshallField(stage, "Name"))
		pointersInitializesStatements.WriteString(allocatedresourceshape.GongMarshallField(stage, "Participant"))
		pointersInitializesStatements.WriteString(allocatedresourceshape.GongMarshallField(stage, "Resource"))
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
		pointersInitializesStatements.WriteString(dataflow.GongMarshallField(stage, "Datas"))
		initializerStatements.WriteString(dataflow.GongMarshallField(stage, "Description"))
		initializerStatements.WriteString(dataflow.GongMarshallField(stage, "ComputedPrefix"))
		initializerStatements.WriteString(dataflow.GongMarshallField(stage, "IsExpanded"))
		initializerStatements.WriteString(dataflow.GongMarshallField(stage, "Type"))
		pointersInitializesStatements.WriteString(dataflow.GongMarshallField(stage, "StartTask"))
		pointersInitializesStatements.WriteString(dataflow.GongMarshallField(stage, "EndTask"))
		pointersInitializesStatements.WriteString(dataflow.GongMarshallField(stage, "StartExternalParticipant"))
		pointersInitializesStatements.WriteString(dataflow.GongMarshallField(stage, "EndExternalParticipant"))
		initializerStatements.WriteString(dataflow.GongMarshallField(stage, "IsDatasNodeExpanded"))
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
func (diagramprocess *DiagramProcess) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(diagramprocess.GongMarshallField(stage, "Name"))
		initializerStatements.WriteString(diagramprocess.GongMarshallField(stage, "Description"))
		initializerStatements.WriteString(diagramprocess.GongMarshallField(stage, "ComputedPrefix"))
		initializerStatements.WriteString(diagramprocess.GongMarshallField(stage, "IsExpanded"))
		initializerStatements.WriteString(diagramprocess.GongMarshallField(stage, "IsChecked"))
		initializerStatements.WriteString(diagramprocess.GongMarshallField(stage, "IsEditable_"))
		initializerStatements.WriteString(diagramprocess.GongMarshallField(stage, "IsShowPrefix"))
		initializerStatements.WriteString(diagramprocess.GongMarshallField(stage, "DefaultBoxWidth"))
		initializerStatements.WriteString(diagramprocess.GongMarshallField(stage, "DefaultBoxHeigth"))
		initializerStatements.WriteString(diagramprocess.GongMarshallField(stage, "Width"))
		initializerStatements.WriteString(diagramprocess.GongMarshallField(stage, "Height"))
		pointersInitializesStatements.WriteString(diagramprocess.GongMarshallField(stage, "Process_Shapes"))
		initializerStatements.WriteString(diagramprocess.GongMarshallField(stage, "IsProcesssNodeExpanded"))
		pointersInitializesStatements.WriteString(diagramprocess.GongMarshallField(stage, "ProcesssWhoseNodeIsExpanded"))
		pointersInitializesStatements.WriteString(diagramprocess.GongMarshallField(stage, "Participant_Shapes"))
		initializerStatements.WriteString(diagramprocess.GongMarshallField(stage, "IsParticipantsNodeExpanded"))
		pointersInitializesStatements.WriteString(diagramprocess.GongMarshallField(stage, "ParticipantWhoseNodeIsExpanded"))
		pointersInitializesStatements.WriteString(diagramprocess.GongMarshallField(stage, "ExternalParticipant_Shapes"))
		initializerStatements.WriteString(diagramprocess.GongMarshallField(stage, "IsExternalParticipantsNodeExpanded"))
		pointersInitializesStatements.WriteString(diagramprocess.GongMarshallField(stage, "ExternalParticipantWhoseNodeIsExpanded"))
		pointersInitializesStatements.WriteString(diagramprocess.GongMarshallField(stage, "ExternalParticipantsWhoseOutDataFlowsNodeIsExpanded"))
		pointersInitializesStatements.WriteString(diagramprocess.GongMarshallField(stage, "ExternalParticipantsWhoseInDataFlowsNodeIsExpanded"))
		pointersInitializesStatements.WriteString(diagramprocess.GongMarshallField(stage, "TasksWhoseNodeIsExpanded"))
		pointersInitializesStatements.WriteString(diagramprocess.GongMarshallField(stage, "Task_Shapes"))
		pointersInitializesStatements.WriteString(diagramprocess.GongMarshallField(stage, "ControlFlowsWhoseNodeIsExpanded"))
		pointersInitializesStatements.WriteString(diagramprocess.GongMarshallField(stage, "ControlFlow_Shapes"))
		pointersInitializesStatements.WriteString(diagramprocess.GongMarshallField(stage, "DataFlowsWhoseNodeIsExpanded"))
		pointersInitializesStatements.WriteString(diagramprocess.GongMarshallField(stage, "DataFlow_Shapes"))
		pointersInitializesStatements.WriteString(diagramprocess.GongMarshallField(stage, "DatasWhoseNodeIsExpanded"))
		pointersInitializesStatements.WriteString(diagramprocess.GongMarshallField(stage, "Data_Shapes"))
		pointersInitializesStatements.WriteString(diagramprocess.GongMarshallField(stage, "DataFlowsWhoseDataNodeIsExpanded"))
		pointersInitializesStatements.WriteString(diagramprocess.GongMarshallField(stage, "AllocatedResourcesWhoseNodeIsExpanded"))
		pointersInitializesStatements.WriteString(diagramprocess.GongMarshallField(stage, "AllocatedResourceShapes"))
		pointersInitializesStatements.WriteString(diagramprocess.GongMarshallField(stage, "AllocatedProcessesWhoseNodeIsExpanded"))
		pointersInitializesStatements.WriteString(diagramprocess.GongMarshallField(stage, "AllocatedProcessShapes"))
		pointersInitializesStatements.WriteString(diagramprocess.GongMarshallField(stage, "Note_Shapes"))
		pointersInitializesStatements.WriteString(diagramprocess.GongMarshallField(stage, "NotesWhoseNodeIsExpanded"))
		initializerStatements.WriteString(diagramprocess.GongMarshallField(stage, "IsNotesNodeExpanded"))
		pointersInitializesStatements.WriteString(diagramprocess.GongMarshallField(stage, "NoteTaskShapes"))
	}
	initRes = initializerStatements.String()
	ptrRes = pointersInitializesStatements.String()
	return
}
func (externalparticipantshape *ExternalParticipantShape) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(externalparticipantshape.GongMarshallField(stage, "Name"))
		pointersInitializesStatements.WriteString(externalparticipantshape.GongMarshallField(stage, "Participant"))
		initializerStatements.WriteString(externalparticipantshape.GongMarshallField(stage, "IsExpanded"))
		initializerStatements.WriteString(externalparticipantshape.GongMarshallField(stage, "X"))
		initializerStatements.WriteString(externalparticipantshape.GongMarshallField(stage, "Y"))
		initializerStatements.WriteString(externalparticipantshape.GongMarshallField(stage, "Width"))
		initializerStatements.WriteString(externalparticipantshape.GongMarshallField(stage, "Height"))
		initializerStatements.WriteString(externalparticipantshape.GongMarshallField(stage, "IsHidden"))
		initializerStatements.WriteString(externalparticipantshape.GongMarshallField(stage, "IsPresentationOverridden"))
		initializerStatements.WriteString(externalparticipantshape.GongMarshallField(stage, "Color"))
		initializerStatements.WriteString(externalparticipantshape.GongMarshallField(stage, "FillOpacity"))
		initializerStatements.WriteString(externalparticipantshape.GongMarshallField(stage, "Stroke"))
		initializerStatements.WriteString(externalparticipantshape.GongMarshallField(stage, "StrokeWidth"))
		initializerStatements.WriteString(externalparticipantshape.GongMarshallField(stage, "StrokeOpacity"))
		initializerStatements.WriteString(externalparticipantshape.GongMarshallField(stage, "StrokeDashArray"))
		initializerStatements.WriteString(externalparticipantshape.GongMarshallField(stage, "RX"))
		initializerStatements.WriteString(externalparticipantshape.GongMarshallField(stage, "FontSize"))
		initializerStatements.WriteString(externalparticipantshape.GongMarshallField(stage, "FontWeight"))
		initializerStatements.WriteString(externalparticipantshape.GongMarshallField(stage, "FontColor"))
		initializerStatements.WriteString(externalparticipantshape.GongMarshallField(stage, "TailHeigth"))
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
		pointersInitializesStatements.WriteString(library.GongMarshallField(stage, "RootProcesses"))
		initializerStatements.WriteString(library.GongMarshallField(stage, "IsProcessesNodeExpanded"))
		pointersInitializesStatements.WriteString(library.GongMarshallField(stage, "ProcesssWhoseNodeIsExpanded"))
		pointersInitializesStatements.WriteString(library.GongMarshallField(stage, "RootDataFlows"))
		initializerStatements.WriteString(library.GongMarshallField(stage, "IsDataFlowsNodeExpanded"))
		pointersInitializesStatements.WriteString(library.GongMarshallField(stage, "DataFlowsWhoseNodeIsExpanded"))
		pointersInitializesStatements.WriteString(library.GongMarshallField(stage, "RootDatas"))
		initializerStatements.WriteString(library.GongMarshallField(stage, "IsDatasNodeExpanded"))
		pointersInitializesStatements.WriteString(library.GongMarshallField(stage, "DatasWhoseNodeIsExpanded"))
		pointersInitializesStatements.WriteString(library.GongMarshallField(stage, "RootResources"))
		initializerStatements.WriteString(library.GongMarshallField(stage, "IsResourcesNodeExpanded"))
		pointersInitializesStatements.WriteString(library.GongMarshallField(stage, "ResourcesWhoseNodeIsExpanded"))
		pointersInitializesStatements.WriteString(library.GongMarshallField(stage, "ParticipantsWhoseNodeIsExpanded"))
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
		initializerStatements.WriteString(note.GongMarshallField(stage, "IsTasksNodeExpanded"))
		pointersInitializesStatements.WriteString(note.GongMarshallField(stage, "Tasks"))
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
func (participant *Participant) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(participant.GongMarshallField(stage, "Name"))
		initializerStatements.WriteString(participant.GongMarshallField(stage, "IsProcessResource"))
		initializerStatements.WriteString(participant.GongMarshallField(stage, "Description"))
		pointersInitializesStatements.WriteString(participant.GongMarshallField(stage, "Resources"))
		initializerStatements.WriteString(participant.GongMarshallField(stage, "IsResourcesNodeExpanded"))
		pointersInitializesStatements.WriteString(participant.GongMarshallField(stage, "Processes"))
		initializerStatements.WriteString(participant.GongMarshallField(stage, "IsProcessesNodeExpanded"))
		initializerStatements.WriteString(participant.GongMarshallField(stage, "ComputedPrefix"))
		initializerStatements.WriteString(participant.GongMarshallField(stage, "IsExpanded"))
		initializerStatements.WriteString(participant.GongMarshallField(stage, "IsTasksNodeExpanded"))
		pointersInitializesStatements.WriteString(participant.GongMarshallField(stage, "Tasks"))
		initializerStatements.WriteString(participant.GongMarshallField(stage, "IsControlFlowsNodeExpanded"))
		pointersInitializesStatements.WriteString(participant.GongMarshallField(stage, "ControlFlows"))
		pointersInitializesStatements.WriteString(participant.GongMarshallField(stage, "TaskWhoseOutControlFlowsNodeIsExpanded"))
		pointersInitializesStatements.WriteString(participant.GongMarshallField(stage, "TaskWhoseInControlFlowsNodeIsExpanded"))
		initializerStatements.WriteString(participant.GongMarshallField(stage, "IsDataFlowsNodeExpanded"))
		pointersInitializesStatements.WriteString(participant.GongMarshallField(stage, "TaskWhoseOutDataFlowsNodeIsExpanded"))
		pointersInitializesStatements.WriteString(participant.GongMarshallField(stage, "TaskWhoseInDataFlowsNodeIsExpanded"))
	}
	initRes = initializerStatements.String()
	ptrRes = pointersInitializesStatements.String()
	return
}
func (participantshape *ParticipantShape) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(participantshape.GongMarshallField(stage, "Name"))
		pointersInitializesStatements.WriteString(participantshape.GongMarshallField(stage, "Participant"))
		initializerStatements.WriteString(participantshape.GongMarshallField(stage, "IsExpanded"))
		initializerStatements.WriteString(participantshape.GongMarshallField(stage, "X"))
		initializerStatements.WriteString(participantshape.GongMarshallField(stage, "Y"))
		initializerStatements.WriteString(participantshape.GongMarshallField(stage, "Width"))
		initializerStatements.WriteString(participantshape.GongMarshallField(stage, "Height"))
		initializerStatements.WriteString(participantshape.GongMarshallField(stage, "IsHidden"))
		initializerStatements.WriteString(participantshape.GongMarshallField(stage, "IsPresentationOverridden"))
		initializerStatements.WriteString(participantshape.GongMarshallField(stage, "Color"))
		initializerStatements.WriteString(participantshape.GongMarshallField(stage, "FillOpacity"))
		initializerStatements.WriteString(participantshape.GongMarshallField(stage, "Stroke"))
		initializerStatements.WriteString(participantshape.GongMarshallField(stage, "StrokeWidth"))
		initializerStatements.WriteString(participantshape.GongMarshallField(stage, "StrokeOpacity"))
		initializerStatements.WriteString(participantshape.GongMarshallField(stage, "StrokeDashArray"))
		initializerStatements.WriteString(participantshape.GongMarshallField(stage, "RX"))
		initializerStatements.WriteString(participantshape.GongMarshallField(stage, "FontSize"))
		initializerStatements.WriteString(participantshape.GongMarshallField(stage, "FontWeight"))
		initializerStatements.WriteString(participantshape.GongMarshallField(stage, "FontColor"))
		initializerStatements.WriteString(participantshape.GongMarshallField(stage, "WidthWeight"))
	}
	initRes = initializerStatements.String()
	ptrRes = pointersInitializesStatements.String()
	return
}
func (process *Process) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(process.GongMarshallField(stage, "Name"))
		initializerStatements.WriteString(process.GongMarshallField(stage, "Description"))
		initializerStatements.WriteString(process.GongMarshallField(stage, "ComputedPrefix"))
		initializerStatements.WriteString(process.GongMarshallField(stage, "IsExpanded"))
		initializerStatements.WriteString(process.GongMarshallField(stage, "SVG_Path"))
		initializerStatements.WriteString(process.GongMarshallField(stage, "InverseAppliedScaling"))
		pointersInitializesStatements.WriteString(process.GongMarshallField(stage, "DiagramProcesss"))
		pointersInitializesStatements.WriteString(process.GongMarshallField(stage, "DiagramProcessWhoseNodeIsExpanded"))
		initializerStatements.WriteString(process.GongMarshallField(stage, "IsSubProcessNodeExpanded"))
		pointersInitializesStatements.WriteString(process.GongMarshallField(stage, "SubProcesses"))
		pointersInitializesStatements.WriteString(process.GongMarshallField(stage, "Participants"))
		pointersInitializesStatements.WriteString(process.GongMarshallField(stage, "ParticipantWhoseNodeIsExpanded"))
		pointersInitializesStatements.WriteString(process.GongMarshallField(stage, "DataFlows"))
		initializerStatements.WriteString(process.GongMarshallField(stage, "IsDataFlowsNodeExpanded"))
		pointersInitializesStatements.WriteString(process.GongMarshallField(stage, "ExternalParticipants"))
		pointersInitializesStatements.WriteString(process.GongMarshallField(stage, "ExternalParticipantWhoseNodeIsExpanded"))
	}
	initRes = initializerStatements.String()
	ptrRes = pointersInitializesStatements.String()
	return
}
func (processshape *ProcessShape) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(processshape.GongMarshallField(stage, "Name"))
		pointersInitializesStatements.WriteString(processshape.GongMarshallField(stage, "Process"))
		initializerStatements.WriteString(processshape.GongMarshallField(stage, "IsExpanded"))
		initializerStatements.WriteString(processshape.GongMarshallField(stage, "X"))
		initializerStatements.WriteString(processshape.GongMarshallField(stage, "Y"))
		initializerStatements.WriteString(processshape.GongMarshallField(stage, "Width"))
		initializerStatements.WriteString(processshape.GongMarshallField(stage, "Height"))
		initializerStatements.WriteString(processshape.GongMarshallField(stage, "IsHidden"))
		initializerStatements.WriteString(processshape.GongMarshallField(stage, "IsPresentationOverridden"))
		initializerStatements.WriteString(processshape.GongMarshallField(stage, "Color"))
		initializerStatements.WriteString(processshape.GongMarshallField(stage, "FillOpacity"))
		initializerStatements.WriteString(processshape.GongMarshallField(stage, "Stroke"))
		initializerStatements.WriteString(processshape.GongMarshallField(stage, "StrokeWidth"))
		initializerStatements.WriteString(processshape.GongMarshallField(stage, "StrokeOpacity"))
		initializerStatements.WriteString(processshape.GongMarshallField(stage, "StrokeDashArray"))
		initializerStatements.WriteString(processshape.GongMarshallField(stage, "RX"))
		initializerStatements.WriteString(processshape.GongMarshallField(stage, "FontSize"))
		initializerStatements.WriteString(processshape.GongMarshallField(stage, "FontWeight"))
		initializerStatements.WriteString(processshape.GongMarshallField(stage, "FontColor"))
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
func (task *Task) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(task.GongMarshallField(stage, "Name"))
		initializerStatements.WriteString(task.GongMarshallField(stage, "Description"))
		initializerStatements.WriteString(task.GongMarshallField(stage, "ComputedPrefix"))
		initializerStatements.WriteString(task.GongMarshallField(stage, "IsExpanded"))
		initializerStatements.WriteString(task.GongMarshallField(stage, "IsStartTask"))
		initializerStatements.WriteString(task.GongMarshallField(stage, "IsEndTask"))
		pointersInitializesStatements.WriteString(task.GongMarshallField(stage, "Type"))
		initializerStatements.WriteString(task.GongMarshallField(stage, "IsTaskNameNotProcessName"))
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
		initializerStatements.WriteString(taskshape.GongMarshallField(stage, "IsExpanded"))
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
