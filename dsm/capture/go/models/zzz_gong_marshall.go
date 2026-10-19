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
	modelsPackageName := "github.com/fullstack-lang/gong/dsm/capture/go/models"
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
		for _, analysisneed := range __gong__sortStageSetInstances(stage.AnalysisNeeds, stage.AnalysisNeed_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			analysisneedIdent := analysisneed.GongGetIdentifier(stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&models.AnalysisNeed{Name: %s}).Stage(stage)", analysisneedIdent, __gong__toRawStringLiteral(analysisneed.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", analysisneedIdent, __gong__toRawStringLiteral(analysisneed.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.ComputedPrefix = %s", analysisneedIdent, __gong__toRawStringLiteral(analysisneed.ComputedPrefix)))
			values.WriteString(fmt.Sprintf("\n\t%s.IsExpanded = %t", analysisneedIdent, analysisneed.IsExpanded))
		}
	}
	if stage != nil {
		for _, concept := range __gong__sortStageSetInstances(stage.Concepts, stage.Concept_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			conceptIdent := concept.GongGetIdentifier(stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&models.Concept{Name: %s}).Stage(stage)", conceptIdent, __gong__toRawStringLiteral(concept.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", conceptIdent, __gong__toRawStringLiteral(concept.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.ComputedPrefix = %s", conceptIdent, __gong__toRawStringLiteral(concept.ComputedPrefix)))
			values.WriteString(fmt.Sprintf("\n\t%s.IsExpanded = %t", conceptIdent, concept.IsExpanded))
			for _, elem := range concept.Tools {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := elem.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Tools = append(%s.Tools, %s)", conceptIdent, conceptIdent, targetIdent))
			}
		}
	}
	if stage != nil {
		for _, conceptshape := range __gong__sortStageSetInstances(stage.ConceptShapes, stage.ConceptShape_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			conceptshapeIdent := conceptshape.GongGetIdentifier(stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&models.ConceptShape{Name: %s}).Stage(stage)", conceptshapeIdent, __gong__toRawStringLiteral(conceptshape.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", conceptshapeIdent, __gong__toRawStringLiteral(conceptshape.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.IsExpanded = %t", conceptshapeIdent, conceptshape.IsExpanded))
			values.WriteString(fmt.Sprintf("\n\t%s.X = %f", conceptshapeIdent, conceptshape.X))
			values.WriteString(fmt.Sprintf("\n\t%s.Y = %f", conceptshapeIdent, conceptshape.Y))
			values.WriteString(fmt.Sprintf("\n\t%s.Width = %f", conceptshapeIdent, conceptshape.Width))
			values.WriteString(fmt.Sprintf("\n\t%s.Height = %f", conceptshapeIdent, conceptshape.Height))
			values.WriteString(fmt.Sprintf("\n\t%s.IsHidden = %t", conceptshapeIdent, conceptshape.IsHidden))
			values.WriteString(fmt.Sprintf("\n\t%s.IsPresentationOverridden = %t", conceptshapeIdent, conceptshape.IsPresentationOverridden))
			values.WriteString(fmt.Sprintf("\n\t%s.Color = %s", conceptshapeIdent, __gong__toRawStringLiteral(conceptshape.Color)))
			values.WriteString(fmt.Sprintf("\n\t%s.FillOpacity = %f", conceptshapeIdent, conceptshape.FillOpacity))
			values.WriteString(fmt.Sprintf("\n\t%s.Stroke = %s", conceptshapeIdent, __gong__toRawStringLiteral(conceptshape.Stroke)))
			values.WriteString(fmt.Sprintf("\n\t%s.StrokeWidth = %f", conceptshapeIdent, conceptshape.StrokeWidth))
			values.WriteString(fmt.Sprintf("\n\t%s.StrokeOpacity = %f", conceptshapeIdent, conceptshape.StrokeOpacity))
			values.WriteString(fmt.Sprintf("\n\t%s.StrokeDashArray = %s", conceptshapeIdent, __gong__toRawStringLiteral(conceptshape.StrokeDashArray)))
			values.WriteString(fmt.Sprintf("\n\t%s.RX = %f", conceptshapeIdent, conceptshape.RX))
			values.WriteString(fmt.Sprintf("\n\t%s.FontSize = %s", conceptshapeIdent, __gong__toRawStringLiteral(conceptshape.FontSize)))
			values.WriteString(fmt.Sprintf("\n\t%s.FontWeight = %s", conceptshapeIdent, __gong__toRawStringLiteral(conceptshape.FontWeight)))
			values.WriteString(fmt.Sprintf("\n\t%s.FontColor = %s", conceptshapeIdent, __gong__toRawStringLiteral(conceptshape.FontColor)))
			if conceptshape.Concept != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := conceptshape.Concept.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Concept = %s", conceptshapeIdent, targetIdent))
			}
		}
	}
	if stage != nil {
		for _, concern := range __gong__sortStageSetInstances(stage.Concerns, stage.Concern_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			concernIdent := concern.GongGetIdentifier(stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&models.Concern{Name: %s}).Stage(stage)", concernIdent, __gong__toRawStringLiteral(concern.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", concernIdent, __gong__toRawStringLiteral(concern.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.IDAirbus = %s", concernIdent, __gong__toRawStringLiteral(concern.IDAirbus)))
			if concern.Priority.ToCodeString() != "" {
				values.WriteString(fmt.Sprintf("\n\t%s.Priority = models.%s", concernIdent, concern.Priority.ToCodeString()))
			} else {
				values.WriteString(fmt.Sprintf("\n\t%s.Priority = \"\"", concernIdent))
			}
			values.WriteString(fmt.Sprintf("\n\t%s.ComputedPrefix = %s", concernIdent, __gong__toRawStringLiteral(concern.ComputedPrefix)))
			values.WriteString(fmt.Sprintf("\n\t%s.IsExpanded = %t", concernIdent, concern.IsExpanded))
			values.WriteString(fmt.Sprintf("\n\t%s.Description = %s", concernIdent, __gong__toRawStringLiteral(concern.Description)))
			values.WriteString(fmt.Sprintf("\n\t%s.IsInputsNodeExpanded = %t", concernIdent, concern.IsInputsNodeExpanded))
			values.WriteString(fmt.Sprintf("\n\t%s.IsOutputsNodeExpanded = %t", concernIdent, concern.IsOutputsNodeExpanded))
			values.WriteString(fmt.Sprintf("\n\t%s.IsWithCompletion = %t", concernIdent, concern.IsWithCompletion))
			if concern.Completion.ToCodeString() != "" {
				values.WriteString(fmt.Sprintf("\n\t%s.Completion = models.%s", concernIdent, concern.Completion.ToCodeString()))
			} else {
				values.WriteString(fmt.Sprintf("\n\t%s.Completion = \"\"", concernIdent))
			}
			for _, elem := range concern.SubConcerns {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := elem.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.SubConcerns = append(%s.SubConcerns, %s)", concernIdent, concernIdent, targetIdent))
			}
			for _, elem := range concern.Inputs {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := elem.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Inputs = append(%s.Inputs, %s)", concernIdent, concernIdent, targetIdent))
			}
			for _, elem := range concern.Outputs {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := elem.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Outputs = append(%s.Outputs, %s)", concernIdent, concernIdent, targetIdent))
			}
		}
	}
	if stage != nil {
		for _, concerncompositionshape := range __gong__sortStageSetInstances(stage.ConcernCompositionShapes, stage.ConcernCompositionShape_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			concerncompositionshapeIdent := concerncompositionshape.GongGetIdentifier(stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&models.ConcernCompositionShape{Name: %s}).Stage(stage)", concerncompositionshapeIdent, __gong__toRawStringLiteral(concerncompositionshape.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", concerncompositionshapeIdent, __gong__toRawStringLiteral(concerncompositionshape.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.StartRatio = %f", concerncompositionshapeIdent, concerncompositionshape.StartRatio))
			values.WriteString(fmt.Sprintf("\n\t%s.EndRatio = %f", concerncompositionshapeIdent, concerncompositionshape.EndRatio))
			if concerncompositionshape.StartOrientation.ToCodeString() != "" {
				values.WriteString(fmt.Sprintf("\n\t%s.StartOrientation = models.%s", concerncompositionshapeIdent, concerncompositionshape.StartOrientation.ToCodeString()))
			} else {
				values.WriteString(fmt.Sprintf("\n\t%s.StartOrientation = \"\"", concerncompositionshapeIdent))
			}
			if concerncompositionshape.EndOrientation.ToCodeString() != "" {
				values.WriteString(fmt.Sprintf("\n\t%s.EndOrientation = models.%s", concerncompositionshapeIdent, concerncompositionshape.EndOrientation.ToCodeString()))
			} else {
				values.WriteString(fmt.Sprintf("\n\t%s.EndOrientation = \"\"", concerncompositionshapeIdent))
			}
			values.WriteString(fmt.Sprintf("\n\t%s.CornerOffsetRatio = %f", concerncompositionshapeIdent, concerncompositionshape.CornerOffsetRatio))
			values.WriteString(fmt.Sprintf("\n\t%s.IsHidden = %t", concerncompositionshapeIdent, concerncompositionshape.IsHidden))
			values.WriteString(fmt.Sprintf("\n\t%s.IsPresentationOverridden = %t", concerncompositionshapeIdent, concerncompositionshape.IsPresentationOverridden))
			values.WriteString(fmt.Sprintf("\n\t%s.Stroke = %s", concerncompositionshapeIdent, __gong__toRawStringLiteral(concerncompositionshape.Stroke)))
			values.WriteString(fmt.Sprintf("\n\t%s.StrokeWidth = %f", concerncompositionshapeIdent, concerncompositionshape.StrokeWidth))
			values.WriteString(fmt.Sprintf("\n\t%s.StrokeOpacity = %f", concerncompositionshapeIdent, concerncompositionshape.StrokeOpacity))
			values.WriteString(fmt.Sprintf("\n\t%s.StrokeDashArray = %s", concerncompositionshapeIdent, __gong__toRawStringLiteral(concerncompositionshape.StrokeDashArray)))
			values.WriteString(fmt.Sprintf("\n\t%s.FontSize = %s", concerncompositionshapeIdent, __gong__toRawStringLiteral(concerncompositionshape.FontSize)))
			values.WriteString(fmt.Sprintf("\n\t%s.FontColor = %s", concerncompositionshapeIdent, __gong__toRawStringLiteral(concerncompositionshape.FontColor)))
			if concerncompositionshape.Concern != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := concerncompositionshape.Concern.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Concern = %s", concerncompositionshapeIdent, targetIdent))
			}
			for _, elem := range concerncompositionshape.ControlPointShapes {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := elem.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.ControlPointShapes = append(%s.ControlPointShapes, %s)", concerncompositionshapeIdent, concerncompositionshapeIdent, targetIdent))
			}
		}
	}
	if stage != nil {
		for _, concerninputshape := range __gong__sortStageSetInstances(stage.ConcernInputShapes, stage.ConcernInputShape_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			concerninputshapeIdent := concerninputshape.GongGetIdentifier(stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&models.ConcernInputShape{Name: %s}).Stage(stage)", concerninputshapeIdent, __gong__toRawStringLiteral(concerninputshape.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", concerninputshapeIdent, __gong__toRawStringLiteral(concerninputshape.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.StartRatio = %f", concerninputshapeIdent, concerninputshape.StartRatio))
			values.WriteString(fmt.Sprintf("\n\t%s.EndRatio = %f", concerninputshapeIdent, concerninputshape.EndRatio))
			if concerninputshape.StartOrientation.ToCodeString() != "" {
				values.WriteString(fmt.Sprintf("\n\t%s.StartOrientation = models.%s", concerninputshapeIdent, concerninputshape.StartOrientation.ToCodeString()))
			} else {
				values.WriteString(fmt.Sprintf("\n\t%s.StartOrientation = \"\"", concerninputshapeIdent))
			}
			if concerninputshape.EndOrientation.ToCodeString() != "" {
				values.WriteString(fmt.Sprintf("\n\t%s.EndOrientation = models.%s", concerninputshapeIdent, concerninputshape.EndOrientation.ToCodeString()))
			} else {
				values.WriteString(fmt.Sprintf("\n\t%s.EndOrientation = \"\"", concerninputshapeIdent))
			}
			values.WriteString(fmt.Sprintf("\n\t%s.CornerOffsetRatio = %f", concerninputshapeIdent, concerninputshape.CornerOffsetRatio))
			values.WriteString(fmt.Sprintf("\n\t%s.IsHidden = %t", concerninputshapeIdent, concerninputshape.IsHidden))
			values.WriteString(fmt.Sprintf("\n\t%s.IsPresentationOverridden = %t", concerninputshapeIdent, concerninputshape.IsPresentationOverridden))
			values.WriteString(fmt.Sprintf("\n\t%s.Stroke = %s", concerninputshapeIdent, __gong__toRawStringLiteral(concerninputshape.Stroke)))
			values.WriteString(fmt.Sprintf("\n\t%s.StrokeWidth = %f", concerninputshapeIdent, concerninputshape.StrokeWidth))
			values.WriteString(fmt.Sprintf("\n\t%s.StrokeOpacity = %f", concerninputshapeIdent, concerninputshape.StrokeOpacity))
			values.WriteString(fmt.Sprintf("\n\t%s.StrokeDashArray = %s", concerninputshapeIdent, __gong__toRawStringLiteral(concerninputshape.StrokeDashArray)))
			values.WriteString(fmt.Sprintf("\n\t%s.FontSize = %s", concerninputshapeIdent, __gong__toRawStringLiteral(concerninputshape.FontSize)))
			values.WriteString(fmt.Sprintf("\n\t%s.FontColor = %s", concerninputshapeIdent, __gong__toRawStringLiteral(concerninputshape.FontColor)))
			if concerninputshape.Deliverable != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := concerninputshape.Deliverable.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Deliverable = %s", concerninputshapeIdent, targetIdent))
			}
			if concerninputshape.Concern != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := concerninputshape.Concern.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Concern = %s", concerninputshapeIdent, targetIdent))
			}
			for _, elem := range concerninputshape.ControlPointShapes {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := elem.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.ControlPointShapes = append(%s.ControlPointShapes, %s)", concerninputshapeIdent, concerninputshapeIdent, targetIdent))
			}
		}
	}
	if stage != nil {
		for _, concernoutputshape := range __gong__sortStageSetInstances(stage.ConcernOutputShapes, stage.ConcernOutputShape_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			concernoutputshapeIdent := concernoutputshape.GongGetIdentifier(stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&models.ConcernOutputShape{Name: %s}).Stage(stage)", concernoutputshapeIdent, __gong__toRawStringLiteral(concernoutputshape.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", concernoutputshapeIdent, __gong__toRawStringLiteral(concernoutputshape.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.StartRatio = %f", concernoutputshapeIdent, concernoutputshape.StartRatio))
			values.WriteString(fmt.Sprintf("\n\t%s.EndRatio = %f", concernoutputshapeIdent, concernoutputshape.EndRatio))
			if concernoutputshape.StartOrientation.ToCodeString() != "" {
				values.WriteString(fmt.Sprintf("\n\t%s.StartOrientation = models.%s", concernoutputshapeIdent, concernoutputshape.StartOrientation.ToCodeString()))
			} else {
				values.WriteString(fmt.Sprintf("\n\t%s.StartOrientation = \"\"", concernoutputshapeIdent))
			}
			if concernoutputshape.EndOrientation.ToCodeString() != "" {
				values.WriteString(fmt.Sprintf("\n\t%s.EndOrientation = models.%s", concernoutputshapeIdent, concernoutputshape.EndOrientation.ToCodeString()))
			} else {
				values.WriteString(fmt.Sprintf("\n\t%s.EndOrientation = \"\"", concernoutputshapeIdent))
			}
			values.WriteString(fmt.Sprintf("\n\t%s.CornerOffsetRatio = %f", concernoutputshapeIdent, concernoutputshape.CornerOffsetRatio))
			values.WriteString(fmt.Sprintf("\n\t%s.IsHidden = %t", concernoutputshapeIdent, concernoutputshape.IsHidden))
			values.WriteString(fmt.Sprintf("\n\t%s.IsPresentationOverridden = %t", concernoutputshapeIdent, concernoutputshape.IsPresentationOverridden))
			values.WriteString(fmt.Sprintf("\n\t%s.Stroke = %s", concernoutputshapeIdent, __gong__toRawStringLiteral(concernoutputshape.Stroke)))
			values.WriteString(fmt.Sprintf("\n\t%s.StrokeWidth = %f", concernoutputshapeIdent, concernoutputshape.StrokeWidth))
			values.WriteString(fmt.Sprintf("\n\t%s.StrokeOpacity = %f", concernoutputshapeIdent, concernoutputshape.StrokeOpacity))
			values.WriteString(fmt.Sprintf("\n\t%s.StrokeDashArray = %s", concernoutputshapeIdent, __gong__toRawStringLiteral(concernoutputshape.StrokeDashArray)))
			values.WriteString(fmt.Sprintf("\n\t%s.FontSize = %s", concernoutputshapeIdent, __gong__toRawStringLiteral(concernoutputshape.FontSize)))
			values.WriteString(fmt.Sprintf("\n\t%s.FontColor = %s", concernoutputshapeIdent, __gong__toRawStringLiteral(concernoutputshape.FontColor)))
			if concernoutputshape.Concern != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := concernoutputshape.Concern.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Concern = %s", concernoutputshapeIdent, targetIdent))
			}
			if concernoutputshape.Deliverable != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := concernoutputshape.Deliverable.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Deliverable = %s", concernoutputshapeIdent, targetIdent))
			}
			for _, elem := range concernoutputshape.ControlPointShapes {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := elem.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.ControlPointShapes = append(%s.ControlPointShapes, %s)", concernoutputshapeIdent, concernoutputshapeIdent, targetIdent))
			}
		}
	}
	if stage != nil {
		for _, concernshape := range __gong__sortStageSetInstances(stage.ConcernShapes, stage.ConcernShape_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			concernshapeIdent := concernshape.GongGetIdentifier(stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&models.ConcernShape{Name: %s}).Stage(stage)", concernshapeIdent, __gong__toRawStringLiteral(concernshape.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", concernshapeIdent, __gong__toRawStringLiteral(concernshape.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.IsExpanded = %t", concernshapeIdent, concernshape.IsExpanded))
			values.WriteString(fmt.Sprintf("\n\t%s.X = %f", concernshapeIdent, concernshape.X))
			values.WriteString(fmt.Sprintf("\n\t%s.Y = %f", concernshapeIdent, concernshape.Y))
			values.WriteString(fmt.Sprintf("\n\t%s.Width = %f", concernshapeIdent, concernshape.Width))
			values.WriteString(fmt.Sprintf("\n\t%s.Height = %f", concernshapeIdent, concernshape.Height))
			values.WriteString(fmt.Sprintf("\n\t%s.IsHidden = %t", concernshapeIdent, concernshape.IsHidden))
			values.WriteString(fmt.Sprintf("\n\t%s.IsPresentationOverridden = %t", concernshapeIdent, concernshape.IsPresentationOverridden))
			values.WriteString(fmt.Sprintf("\n\t%s.Color = %s", concernshapeIdent, __gong__toRawStringLiteral(concernshape.Color)))
			values.WriteString(fmt.Sprintf("\n\t%s.FillOpacity = %f", concernshapeIdent, concernshape.FillOpacity))
			values.WriteString(fmt.Sprintf("\n\t%s.Stroke = %s", concernshapeIdent, __gong__toRawStringLiteral(concernshape.Stroke)))
			values.WriteString(fmt.Sprintf("\n\t%s.StrokeWidth = %f", concernshapeIdent, concernshape.StrokeWidth))
			values.WriteString(fmt.Sprintf("\n\t%s.StrokeOpacity = %f", concernshapeIdent, concernshape.StrokeOpacity))
			values.WriteString(fmt.Sprintf("\n\t%s.StrokeDashArray = %s", concernshapeIdent, __gong__toRawStringLiteral(concernshape.StrokeDashArray)))
			values.WriteString(fmt.Sprintf("\n\t%s.RX = %f", concernshapeIdent, concernshape.RX))
			values.WriteString(fmt.Sprintf("\n\t%s.FontSize = %s", concernshapeIdent, __gong__toRawStringLiteral(concernshape.FontSize)))
			values.WriteString(fmt.Sprintf("\n\t%s.FontWeight = %s", concernshapeIdent, __gong__toRawStringLiteral(concernshape.FontWeight)))
			values.WriteString(fmt.Sprintf("\n\t%s.FontColor = %s", concernshapeIdent, __gong__toRawStringLiteral(concernshape.FontColor)))
			if concernshape.Concern != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := concernshape.Concern.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Concern = %s", concernshapeIdent, targetIdent))
			}
		}
	}
	if stage != nil {
		for _, controlpointshape := range __gong__sortStageSetInstances(stage.ControlPointShapes, stage.ControlPointShape_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			controlpointshapeIdent := controlpointshape.GongGetIdentifier(stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&models.ControlPointShape{Name: %s}).Stage(stage)", controlpointshapeIdent, __gong__toRawStringLiteral(controlpointshape.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", controlpointshapeIdent, __gong__toRawStringLiteral(controlpointshape.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.X_Relative = %f", controlpointshapeIdent, controlpointshape.X_Relative))
			values.WriteString(fmt.Sprintf("\n\t%s.Y_Relative = %f", controlpointshapeIdent, controlpointshape.Y_Relative))
			values.WriteString(fmt.Sprintf("\n\t%s.IsStartShapeTheClosestShape = %t", controlpointshapeIdent, controlpointshape.IsStartShapeTheClosestShape))
		}
	}
	if stage != nil {
		for _, deliverable := range __gong__sortStageSetInstances(stage.Deliverables, stage.Deliverable_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			deliverableIdent := deliverable.GongGetIdentifier(stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&models.Deliverable{Name: %s}).Stage(stage)", deliverableIdent, __gong__toRawStringLiteral(deliverable.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", deliverableIdent, __gong__toRawStringLiteral(deliverable.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.ComputedPrefix = %s", deliverableIdent, __gong__toRawStringLiteral(deliverable.ComputedPrefix)))
			values.WriteString(fmt.Sprintf("\n\t%s.IsExpanded = %t", deliverableIdent, deliverable.IsExpanded))
			values.WriteString(fmt.Sprintf("\n\t%s.Description = %s", deliverableIdent, __gong__toRawStringLiteral(deliverable.Description)))
			values.WriteString(fmt.Sprintf("\n\t%s.IsProducersNodeExpanded = %t", deliverableIdent, deliverable.IsProducersNodeExpanded))
			values.WriteString(fmt.Sprintf("\n\t%s.IsConsumersNodeExpanded = %t", deliverableIdent, deliverable.IsConsumersNodeExpanded))
			for _, elem := range deliverable.SubDeliverables {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := elem.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.SubDeliverables = append(%s.SubDeliverables, %s)", deliverableIdent, deliverableIdent, targetIdent))
			}
			for _, elem := range deliverable.Views {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := elem.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Views = append(%s.Views, %s)", deliverableIdent, deliverableIdent, targetIdent))
			}
			for _, elem := range deliverable.Concepts {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := elem.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Concepts = append(%s.Concepts, %s)", deliverableIdent, deliverableIdent, targetIdent))
			}
		}
	}
	if stage != nil {
		for _, deliverablecompositionshape := range __gong__sortStageSetInstances(stage.DeliverableCompositionShapes, stage.DeliverableCompositionShape_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			deliverablecompositionshapeIdent := deliverablecompositionshape.GongGetIdentifier(stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&models.DeliverableCompositionShape{Name: %s}).Stage(stage)", deliverablecompositionshapeIdent, __gong__toRawStringLiteral(deliverablecompositionshape.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", deliverablecompositionshapeIdent, __gong__toRawStringLiteral(deliverablecompositionshape.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.StartRatio = %f", deliverablecompositionshapeIdent, deliverablecompositionshape.StartRatio))
			values.WriteString(fmt.Sprintf("\n\t%s.EndRatio = %f", deliverablecompositionshapeIdent, deliverablecompositionshape.EndRatio))
			if deliverablecompositionshape.StartOrientation.ToCodeString() != "" {
				values.WriteString(fmt.Sprintf("\n\t%s.StartOrientation = models.%s", deliverablecompositionshapeIdent, deliverablecompositionshape.StartOrientation.ToCodeString()))
			} else {
				values.WriteString(fmt.Sprintf("\n\t%s.StartOrientation = \"\"", deliverablecompositionshapeIdent))
			}
			if deliverablecompositionshape.EndOrientation.ToCodeString() != "" {
				values.WriteString(fmt.Sprintf("\n\t%s.EndOrientation = models.%s", deliverablecompositionshapeIdent, deliverablecompositionshape.EndOrientation.ToCodeString()))
			} else {
				values.WriteString(fmt.Sprintf("\n\t%s.EndOrientation = \"\"", deliverablecompositionshapeIdent))
			}
			values.WriteString(fmt.Sprintf("\n\t%s.CornerOffsetRatio = %f", deliverablecompositionshapeIdent, deliverablecompositionshape.CornerOffsetRatio))
			values.WriteString(fmt.Sprintf("\n\t%s.IsHidden = %t", deliverablecompositionshapeIdent, deliverablecompositionshape.IsHidden))
			values.WriteString(fmt.Sprintf("\n\t%s.IsPresentationOverridden = %t", deliverablecompositionshapeIdent, deliverablecompositionshape.IsPresentationOverridden))
			values.WriteString(fmt.Sprintf("\n\t%s.Stroke = %s", deliverablecompositionshapeIdent, __gong__toRawStringLiteral(deliverablecompositionshape.Stroke)))
			values.WriteString(fmt.Sprintf("\n\t%s.StrokeWidth = %f", deliverablecompositionshapeIdent, deliverablecompositionshape.StrokeWidth))
			values.WriteString(fmt.Sprintf("\n\t%s.StrokeOpacity = %f", deliverablecompositionshapeIdent, deliverablecompositionshape.StrokeOpacity))
			values.WriteString(fmt.Sprintf("\n\t%s.StrokeDashArray = %s", deliverablecompositionshapeIdent, __gong__toRawStringLiteral(deliverablecompositionshape.StrokeDashArray)))
			values.WriteString(fmt.Sprintf("\n\t%s.FontSize = %s", deliverablecompositionshapeIdent, __gong__toRawStringLiteral(deliverablecompositionshape.FontSize)))
			values.WriteString(fmt.Sprintf("\n\t%s.FontColor = %s", deliverablecompositionshapeIdent, __gong__toRawStringLiteral(deliverablecompositionshape.FontColor)))
			if deliverablecompositionshape.Deliverable != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := deliverablecompositionshape.Deliverable.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Deliverable = %s", deliverablecompositionshapeIdent, targetIdent))
			}
			for _, elem := range deliverablecompositionshape.ControlPointShapes {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := elem.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.ControlPointShapes = append(%s.ControlPointShapes, %s)", deliverablecompositionshapeIdent, deliverablecompositionshapeIdent, targetIdent))
			}
		}
	}
	if stage != nil {
		for _, deliverableconceptshape := range __gong__sortStageSetInstances(stage.DeliverableConceptShapes, stage.DeliverableConceptShape_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			deliverableconceptshapeIdent := deliverableconceptshape.GongGetIdentifier(stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&models.DeliverableConceptShape{Name: %s}).Stage(stage)", deliverableconceptshapeIdent, __gong__toRawStringLiteral(deliverableconceptshape.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", deliverableconceptshapeIdent, __gong__toRawStringLiteral(deliverableconceptshape.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.StartRatio = %f", deliverableconceptshapeIdent, deliverableconceptshape.StartRatio))
			values.WriteString(fmt.Sprintf("\n\t%s.EndRatio = %f", deliverableconceptshapeIdent, deliverableconceptshape.EndRatio))
			if deliverableconceptshape.StartOrientation.ToCodeString() != "" {
				values.WriteString(fmt.Sprintf("\n\t%s.StartOrientation = models.%s", deliverableconceptshapeIdent, deliverableconceptshape.StartOrientation.ToCodeString()))
			} else {
				values.WriteString(fmt.Sprintf("\n\t%s.StartOrientation = \"\"", deliverableconceptshapeIdent))
			}
			if deliverableconceptshape.EndOrientation.ToCodeString() != "" {
				values.WriteString(fmt.Sprintf("\n\t%s.EndOrientation = models.%s", deliverableconceptshapeIdent, deliverableconceptshape.EndOrientation.ToCodeString()))
			} else {
				values.WriteString(fmt.Sprintf("\n\t%s.EndOrientation = \"\"", deliverableconceptshapeIdent))
			}
			values.WriteString(fmt.Sprintf("\n\t%s.CornerOffsetRatio = %f", deliverableconceptshapeIdent, deliverableconceptshape.CornerOffsetRatio))
			values.WriteString(fmt.Sprintf("\n\t%s.IsHidden = %t", deliverableconceptshapeIdent, deliverableconceptshape.IsHidden))
			values.WriteString(fmt.Sprintf("\n\t%s.IsPresentationOverridden = %t", deliverableconceptshapeIdent, deliverableconceptshape.IsPresentationOverridden))
			values.WriteString(fmt.Sprintf("\n\t%s.Stroke = %s", deliverableconceptshapeIdent, __gong__toRawStringLiteral(deliverableconceptshape.Stroke)))
			values.WriteString(fmt.Sprintf("\n\t%s.StrokeWidth = %f", deliverableconceptshapeIdent, deliverableconceptshape.StrokeWidth))
			values.WriteString(fmt.Sprintf("\n\t%s.StrokeOpacity = %f", deliverableconceptshapeIdent, deliverableconceptshape.StrokeOpacity))
			values.WriteString(fmt.Sprintf("\n\t%s.StrokeDashArray = %s", deliverableconceptshapeIdent, __gong__toRawStringLiteral(deliverableconceptshape.StrokeDashArray)))
			values.WriteString(fmt.Sprintf("\n\t%s.FontSize = %s", deliverableconceptshapeIdent, __gong__toRawStringLiteral(deliverableconceptshape.FontSize)))
			values.WriteString(fmt.Sprintf("\n\t%s.FontColor = %s", deliverableconceptshapeIdent, __gong__toRawStringLiteral(deliverableconceptshape.FontColor)))
			if deliverableconceptshape.Deliverable != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := deliverableconceptshape.Deliverable.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Deliverable = %s", deliverableconceptshapeIdent, targetIdent))
			}
			if deliverableconceptshape.Concept != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := deliverableconceptshape.Concept.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Concept = %s", deliverableconceptshapeIdent, targetIdent))
			}
			for _, elem := range deliverableconceptshape.ControlPointShapes {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := elem.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.ControlPointShapes = append(%s.ControlPointShapes, %s)", deliverableconceptshapeIdent, deliverableconceptshapeIdent, targetIdent))
			}
		}
	}
	if stage != nil {
		for _, deliverableshape := range __gong__sortStageSetInstances(stage.DeliverableShapes, stage.DeliverableShape_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			deliverableshapeIdent := deliverableshape.GongGetIdentifier(stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&models.DeliverableShape{Name: %s}).Stage(stage)", deliverableshapeIdent, __gong__toRawStringLiteral(deliverableshape.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", deliverableshapeIdent, __gong__toRawStringLiteral(deliverableshape.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.IsExpanded = %t", deliverableshapeIdent, deliverableshape.IsExpanded))
			values.WriteString(fmt.Sprintf("\n\t%s.X = %f", deliverableshapeIdent, deliverableshape.X))
			values.WriteString(fmt.Sprintf("\n\t%s.Y = %f", deliverableshapeIdent, deliverableshape.Y))
			values.WriteString(fmt.Sprintf("\n\t%s.Width = %f", deliverableshapeIdent, deliverableshape.Width))
			values.WriteString(fmt.Sprintf("\n\t%s.Height = %f", deliverableshapeIdent, deliverableshape.Height))
			values.WriteString(fmt.Sprintf("\n\t%s.IsHidden = %t", deliverableshapeIdent, deliverableshape.IsHidden))
			values.WriteString(fmt.Sprintf("\n\t%s.IsPresentationOverridden = %t", deliverableshapeIdent, deliverableshape.IsPresentationOverridden))
			values.WriteString(fmt.Sprintf("\n\t%s.Color = %s", deliverableshapeIdent, __gong__toRawStringLiteral(deliverableshape.Color)))
			values.WriteString(fmt.Sprintf("\n\t%s.FillOpacity = %f", deliverableshapeIdent, deliverableshape.FillOpacity))
			values.WriteString(fmt.Sprintf("\n\t%s.Stroke = %s", deliverableshapeIdent, __gong__toRawStringLiteral(deliverableshape.Stroke)))
			values.WriteString(fmt.Sprintf("\n\t%s.StrokeWidth = %f", deliverableshapeIdent, deliverableshape.StrokeWidth))
			values.WriteString(fmt.Sprintf("\n\t%s.StrokeOpacity = %f", deliverableshapeIdent, deliverableshape.StrokeOpacity))
			values.WriteString(fmt.Sprintf("\n\t%s.StrokeDashArray = %s", deliverableshapeIdent, __gong__toRawStringLiteral(deliverableshape.StrokeDashArray)))
			values.WriteString(fmt.Sprintf("\n\t%s.RX = %f", deliverableshapeIdent, deliverableshape.RX))
			values.WriteString(fmt.Sprintf("\n\t%s.FontSize = %s", deliverableshapeIdent, __gong__toRawStringLiteral(deliverableshape.FontSize)))
			values.WriteString(fmt.Sprintf("\n\t%s.FontWeight = %s", deliverableshapeIdent, __gong__toRawStringLiteral(deliverableshape.FontWeight)))
			values.WriteString(fmt.Sprintf("\n\t%s.FontColor = %s", deliverableshapeIdent, __gong__toRawStringLiteral(deliverableshape.FontColor)))
			if deliverableshape.Deliverable != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := deliverableshape.Deliverable.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Deliverable = %s", deliverableshapeIdent, targetIdent))
			}
		}
	}
	if stage != nil {
		for _, deliverableviewshape := range __gong__sortStageSetInstances(stage.DeliverableViewShapes, stage.DeliverableViewShape_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			deliverableviewshapeIdent := deliverableviewshape.GongGetIdentifier(stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&models.DeliverableViewShape{Name: %s}).Stage(stage)", deliverableviewshapeIdent, __gong__toRawStringLiteral(deliverableviewshape.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", deliverableviewshapeIdent, __gong__toRawStringLiteral(deliverableviewshape.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.StartRatio = %f", deliverableviewshapeIdent, deliverableviewshape.StartRatio))
			values.WriteString(fmt.Sprintf("\n\t%s.EndRatio = %f", deliverableviewshapeIdent, deliverableviewshape.EndRatio))
			if deliverableviewshape.StartOrientation.ToCodeString() != "" {
				values.WriteString(fmt.Sprintf("\n\t%s.StartOrientation = models.%s", deliverableviewshapeIdent, deliverableviewshape.StartOrientation.ToCodeString()))
			} else {
				values.WriteString(fmt.Sprintf("\n\t%s.StartOrientation = \"\"", deliverableviewshapeIdent))
			}
			if deliverableviewshape.EndOrientation.ToCodeString() != "" {
				values.WriteString(fmt.Sprintf("\n\t%s.EndOrientation = models.%s", deliverableviewshapeIdent, deliverableviewshape.EndOrientation.ToCodeString()))
			} else {
				values.WriteString(fmt.Sprintf("\n\t%s.EndOrientation = \"\"", deliverableviewshapeIdent))
			}
			values.WriteString(fmt.Sprintf("\n\t%s.CornerOffsetRatio = %f", deliverableviewshapeIdent, deliverableviewshape.CornerOffsetRatio))
			values.WriteString(fmt.Sprintf("\n\t%s.IsHidden = %t", deliverableviewshapeIdent, deliverableviewshape.IsHidden))
			values.WriteString(fmt.Sprintf("\n\t%s.IsPresentationOverridden = %t", deliverableviewshapeIdent, deliverableviewshape.IsPresentationOverridden))
			values.WriteString(fmt.Sprintf("\n\t%s.Stroke = %s", deliverableviewshapeIdent, __gong__toRawStringLiteral(deliverableviewshape.Stroke)))
			values.WriteString(fmt.Sprintf("\n\t%s.StrokeWidth = %f", deliverableviewshapeIdent, deliverableviewshape.StrokeWidth))
			values.WriteString(fmt.Sprintf("\n\t%s.StrokeOpacity = %f", deliverableviewshapeIdent, deliverableviewshape.StrokeOpacity))
			values.WriteString(fmt.Sprintf("\n\t%s.StrokeDashArray = %s", deliverableviewshapeIdent, __gong__toRawStringLiteral(deliverableviewshape.StrokeDashArray)))
			values.WriteString(fmt.Sprintf("\n\t%s.FontSize = %s", deliverableviewshapeIdent, __gong__toRawStringLiteral(deliverableviewshape.FontSize)))
			values.WriteString(fmt.Sprintf("\n\t%s.FontColor = %s", deliverableviewshapeIdent, __gong__toRawStringLiteral(deliverableviewshape.FontColor)))
			if deliverableviewshape.Deliverable != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := deliverableviewshape.Deliverable.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Deliverable = %s", deliverableviewshapeIdent, targetIdent))
			}
			if deliverableviewshape.View != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := deliverableviewshape.View.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.View = %s", deliverableviewshapeIdent, targetIdent))
			}
			for _, elem := range deliverableviewshape.ControlPointShapes {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := elem.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.ControlPointShapes = append(%s.ControlPointShapes, %s)", deliverableviewshapeIdent, deliverableviewshapeIdent, targetIdent))
			}
		}
	}
	if stage != nil {
		for _, diagram := range __gong__sortStageSetInstances(stage.Diagrams, stage.Diagram_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			diagramIdent := diagram.GongGetIdentifier(stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&models.Diagram{Name: %s}).Stage(stage)", diagramIdent, __gong__toRawStringLiteral(diagram.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", diagramIdent, __gong__toRawStringLiteral(diagram.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.ComputedPrefix = %s", diagramIdent, __gong__toRawStringLiteral(diagram.ComputedPrefix)))
			values.WriteString(fmt.Sprintf("\n\t%s.IsExpanded = %t", diagramIdent, diagram.IsExpanded))
			values.WriteString(fmt.Sprintf("\n\t%s.IsChecked = %t", diagramIdent, diagram.IsChecked))
			values.WriteString(fmt.Sprintf("\n\t%s.IsEditable_ = %t", diagramIdent, diagram.IsEditable_))
			values.WriteString(fmt.Sprintf("\n\t%s.ShowPrefix = %t", diagramIdent, diagram.ShowPrefix))
			values.WriteString(fmt.Sprintf("\n\t%s.DefaultBoxWidth = %f", diagramIdent, diagram.DefaultBoxWidth))
			values.WriteString(fmt.Sprintf("\n\t%s.DefaultBoxHeigth = %f", diagramIdent, diagram.DefaultBoxHeigth))
			values.WriteString(fmt.Sprintf("\n\t%s.Width = %f", diagramIdent, diagram.Width))
			values.WriteString(fmt.Sprintf("\n\t%s.Height = %f", diagramIdent, diagram.Height))
			values.WriteString(fmt.Sprintf("\n\t%s.IsRequirementsNodeExpanded = %t", diagramIdent, diagram.IsRequirementsNodeExpanded))
			values.WriteString(fmt.Sprintf("\n\t%s.IsViewsNodeExpanded = %t", diagramIdent, diagram.IsViewsNodeExpanded))
			values.WriteString(fmt.Sprintf("\n\t%s.IsConceptsNodeExpanded = %t", diagramIdent, diagram.IsConceptsNodeExpanded))
			values.WriteString(fmt.Sprintf("\n\t%s.IsPBSNodeExpanded = %t", diagramIdent, diagram.IsPBSNodeExpanded))
			values.WriteString(fmt.Sprintf("\n\t%s.IsConcernsNodeExpanded = %t", diagramIdent, diagram.IsConcernsNodeExpanded))
			values.WriteString(fmt.Sprintf("\n\t%s.IsNotesNodeExpanded = %t", diagramIdent, diagram.IsNotesNodeExpanded))
			values.WriteString(fmt.Sprintf("\n\t%s.IsStakeholdersNodeExpanded = %t", diagramIdent, diagram.IsStakeholdersNodeExpanded))
			values.WriteString(fmt.Sprintf("\n\t%s.IsDiagramsNodeExpanded = %t", diagramIdent, diagram.IsDiagramsNodeExpanded))
			for _, elem := range diagram.ConcernsWhoseRequirementsNodeIsExpanded {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := elem.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.ConcernsWhoseRequirementsNodeIsExpanded = append(%s.ConcernsWhoseRequirementsNodeIsExpanded, %s)", diagramIdent, diagramIdent, targetIdent))
			}
			for _, elem := range diagram.Deliverable_Shapes {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := elem.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Deliverable_Shapes = append(%s.Deliverable_Shapes, %s)", diagramIdent, diagramIdent, targetIdent))
			}
			for _, elem := range diagram.DeliverablesWhoseNodeIsExpanded {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := elem.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.DeliverablesWhoseNodeIsExpanded = append(%s.DeliverablesWhoseNodeIsExpanded, %s)", diagramIdent, diagramIdent, targetIdent))
			}
			for _, elem := range diagram.DeliverablesWhoseViewsNodeIsExpanded {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := elem.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.DeliverablesWhoseViewsNodeIsExpanded = append(%s.DeliverablesWhoseViewsNodeIsExpanded, %s)", diagramIdent, diagramIdent, targetIdent))
			}
			for _, elem := range diagram.DeliverablesWhoseConceptsNodeIsExpanded {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := elem.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.DeliverablesWhoseConceptsNodeIsExpanded = append(%s.DeliverablesWhoseConceptsNodeIsExpanded, %s)", diagramIdent, diagramIdent, targetIdent))
			}
			for _, elem := range diagram.DeliverableComposition_Shapes {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := elem.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.DeliverableComposition_Shapes = append(%s.DeliverableComposition_Shapes, %s)", diagramIdent, diagramIdent, targetIdent))
			}
			for _, elem := range diagram.Concern_Shapes {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := elem.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Concern_Shapes = append(%s.Concern_Shapes, %s)", diagramIdent, diagramIdent, targetIdent))
			}
			for _, elem := range diagram.ConcernsWhoseNodeIsExpanded {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := elem.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.ConcernsWhoseNodeIsExpanded = append(%s.ConcernsWhoseNodeIsExpanded, %s)", diagramIdent, diagramIdent, targetIdent))
			}
			for _, elem := range diagram.ConcernsWhoseInputNodeIsExpanded {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := elem.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.ConcernsWhoseInputNodeIsExpanded = append(%s.ConcernsWhoseInputNodeIsExpanded, %s)", diagramIdent, diagramIdent, targetIdent))
			}
			for _, elem := range diagram.ConcernsWhoseStakeholderNodeIsExpanded {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := elem.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.ConcernsWhoseStakeholderNodeIsExpanded = append(%s.ConcernsWhoseStakeholderNodeIsExpanded, %s)", diagramIdent, diagramIdent, targetIdent))
			}
			for _, elem := range diagram.ConcernssWhoseOutputNodeIsExpanded {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := elem.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.ConcernssWhoseOutputNodeIsExpanded = append(%s.ConcernssWhoseOutputNodeIsExpanded, %s)", diagramIdent, diagramIdent, targetIdent))
			}
			for _, elem := range diagram.ConcernComposition_Shapes {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := elem.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.ConcernComposition_Shapes = append(%s.ConcernComposition_Shapes, %s)", diagramIdent, diagramIdent, targetIdent))
			}
			for _, elem := range diagram.ConcernInputShapes {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := elem.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.ConcernInputShapes = append(%s.ConcernInputShapes, %s)", diagramIdent, diagramIdent, targetIdent))
			}
			for _, elem := range diagram.ConcernOutputShapes {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := elem.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.ConcernOutputShapes = append(%s.ConcernOutputShapes, %s)", diagramIdent, diagramIdent, targetIdent))
			}
			for _, elem := range diagram.Note_Shapes {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := elem.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Note_Shapes = append(%s.Note_Shapes, %s)", diagramIdent, diagramIdent, targetIdent))
			}
			for _, elem := range diagram.NotesWhoseNodeIsExpanded {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := elem.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.NotesWhoseNodeIsExpanded = append(%s.NotesWhoseNodeIsExpanded, %s)", diagramIdent, diagramIdent, targetIdent))
			}
			for _, elem := range diagram.NoteDeliverableShapes {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := elem.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.NoteDeliverableShapes = append(%s.NoteDeliverableShapes, %s)", diagramIdent, diagramIdent, targetIdent))
			}
			for _, elem := range diagram.NoteTaskShapes {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := elem.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.NoteTaskShapes = append(%s.NoteTaskShapes, %s)", diagramIdent, diagramIdent, targetIdent))
			}
			for _, elem := range diagram.NoteResourceShapes {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := elem.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.NoteResourceShapes = append(%s.NoteResourceShapes, %s)", diagramIdent, diagramIdent, targetIdent))
			}
			for _, elem := range diagram.Stakeholder_Shapes {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := elem.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Stakeholder_Shapes = append(%s.Stakeholder_Shapes, %s)", diagramIdent, diagramIdent, targetIdent))
			}
			for _, elem := range diagram.ResourcesWhoseNodeIsExpanded {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := elem.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.ResourcesWhoseNodeIsExpanded = append(%s.ResourcesWhoseNodeIsExpanded, %s)", diagramIdent, diagramIdent, targetIdent))
			}
			for _, elem := range diagram.ResourceComposition_Shapes {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := elem.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.ResourceComposition_Shapes = append(%s.ResourceComposition_Shapes, %s)", diagramIdent, diagramIdent, targetIdent))
			}
			for _, elem := range diagram.StakeholderConcernShapes {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := elem.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.StakeholderConcernShapes = append(%s.StakeholderConcernShapes, %s)", diagramIdent, diagramIdent, targetIdent))
			}
			for _, elem := range diagram.Requirement_Shapes {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := elem.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Requirement_Shapes = append(%s.Requirement_Shapes, %s)", diagramIdent, diagramIdent, targetIdent))
			}
			for _, elem := range diagram.RequirementsWhoseNodeIsExpanded {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := elem.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.RequirementsWhoseNodeIsExpanded = append(%s.RequirementsWhoseNodeIsExpanded, %s)", diagramIdent, diagramIdent, targetIdent))
			}
			for _, elem := range diagram.RequirementsWhoseConcernsNodeIsExpanded {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := elem.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.RequirementsWhoseConcernsNodeIsExpanded = append(%s.RequirementsWhoseConcernsNodeIsExpanded, %s)", diagramIdent, diagramIdent, targetIdent))
			}
			for _, elem := range diagram.RequirementsWhoseViewsNodeIsExpanded {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := elem.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.RequirementsWhoseViewsNodeIsExpanded = append(%s.RequirementsWhoseViewsNodeIsExpanded, %s)", diagramIdent, diagramIdent, targetIdent))
			}
			for _, elem := range diagram.RequirementsWhoseConceptsNodeIsExpanded {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := elem.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.RequirementsWhoseConceptsNodeIsExpanded = append(%s.RequirementsWhoseConceptsNodeIsExpanded, %s)", diagramIdent, diagramIdent, targetIdent))
			}
			for _, elem := range diagram.RequirementConcernShapes {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := elem.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.RequirementConcernShapes = append(%s.RequirementConcernShapes, %s)", diagramIdent, diagramIdent, targetIdent))
			}
			for _, elem := range diagram.RequirementViewShapes {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := elem.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.RequirementViewShapes = append(%s.RequirementViewShapes, %s)", diagramIdent, diagramIdent, targetIdent))
			}
			for _, elem := range diagram.RequirementConceptShapes {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := elem.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.RequirementConceptShapes = append(%s.RequirementConceptShapes, %s)", diagramIdent, diagramIdent, targetIdent))
			}
			for _, elem := range diagram.View_Shapes {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := elem.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.View_Shapes = append(%s.View_Shapes, %s)", diagramIdent, diagramIdent, targetIdent))
			}
			for _, elem := range diagram.ViewsWhoseNodeIsExpanded {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := elem.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.ViewsWhoseNodeIsExpanded = append(%s.ViewsWhoseNodeIsExpanded, %s)", diagramIdent, diagramIdent, targetIdent))
			}
			for _, elem := range diagram.ViewsWhoseDeliverablesNodeIsExpanded {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := elem.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.ViewsWhoseDeliverablesNodeIsExpanded = append(%s.ViewsWhoseDeliverablesNodeIsExpanded, %s)", diagramIdent, diagramIdent, targetIdent))
			}
			for _, elem := range diagram.ViewsWhoseConceptsNodeIsExpanded {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := elem.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.ViewsWhoseConceptsNodeIsExpanded = append(%s.ViewsWhoseConceptsNodeIsExpanded, %s)", diagramIdent, diagramIdent, targetIdent))
			}
			for _, elem := range diagram.ViewsWhoseRequirementsNodeIsExpanded {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := elem.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.ViewsWhoseRequirementsNodeIsExpanded = append(%s.ViewsWhoseRequirementsNodeIsExpanded, %s)", diagramIdent, diagramIdent, targetIdent))
			}
			for _, elem := range diagram.DeliverableViewShapes {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := elem.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.DeliverableViewShapes = append(%s.DeliverableViewShapes, %s)", diagramIdent, diagramIdent, targetIdent))
			}
			for _, elem := range diagram.ViewConceptShapes {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := elem.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.ViewConceptShapes = append(%s.ViewConceptShapes, %s)", diagramIdent, diagramIdent, targetIdent))
			}
			for _, elem := range diagram.Concept_Shapes {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := elem.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Concept_Shapes = append(%s.Concept_Shapes, %s)", diagramIdent, diagramIdent, targetIdent))
			}
			for _, elem := range diagram.ConceptsWhoseNodeIsExpanded {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := elem.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.ConceptsWhoseNodeIsExpanded = append(%s.ConceptsWhoseNodeIsExpanded, %s)", diagramIdent, diagramIdent, targetIdent))
			}
			for _, elem := range diagram.ConceptsWhoseViewsNodeIsExpanded {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := elem.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.ConceptsWhoseViewsNodeIsExpanded = append(%s.ConceptsWhoseViewsNodeIsExpanded, %s)", diagramIdent, diagramIdent, targetIdent))
			}
			for _, elem := range diagram.ConceptsWhoseDeliverablesNodeIsExpanded {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := elem.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.ConceptsWhoseDeliverablesNodeIsExpanded = append(%s.ConceptsWhoseDeliverablesNodeIsExpanded, %s)", diagramIdent, diagramIdent, targetIdent))
			}
			for _, elem := range diagram.ConceptsWhoseRequirementsNodeIsExpanded {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := elem.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.ConceptsWhoseRequirementsNodeIsExpanded = append(%s.ConceptsWhoseRequirementsNodeIsExpanded, %s)", diagramIdent, diagramIdent, targetIdent))
			}
			for _, elem := range diagram.DeliverableConceptShapes {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := elem.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.DeliverableConceptShapes = append(%s.DeliverableConceptShapes, %s)", diagramIdent, diagramIdent, targetIdent))
			}
			for _, elem := range diagram.Diagram_Shapes {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := elem.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Diagram_Shapes = append(%s.Diagram_Shapes, %s)", diagramIdent, diagramIdent, targetIdent))
			}
			for _, elem := range diagram.DiagramsWhoseNodeIsExpanded {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := elem.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.DiagramsWhoseNodeIsExpanded = append(%s.DiagramsWhoseNodeIsExpanded, %s)", diagramIdent, diagramIdent, targetIdent))
			}
		}
	}
	if stage != nil {
		for _, diagramshape := range __gong__sortStageSetInstances(stage.DiagramShapes, stage.DiagramShape_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			diagramshapeIdent := diagramshape.GongGetIdentifier(stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&models.DiagramShape{Name: %s}).Stage(stage)", diagramshapeIdent, __gong__toRawStringLiteral(diagramshape.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", diagramshapeIdent, __gong__toRawStringLiteral(diagramshape.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.IsExpanded = %t", diagramshapeIdent, diagramshape.IsExpanded))
			values.WriteString(fmt.Sprintf("\n\t%s.X = %f", diagramshapeIdent, diagramshape.X))
			values.WriteString(fmt.Sprintf("\n\t%s.Y = %f", diagramshapeIdent, diagramshape.Y))
			values.WriteString(fmt.Sprintf("\n\t%s.Width = %f", diagramshapeIdent, diagramshape.Width))
			values.WriteString(fmt.Sprintf("\n\t%s.Height = %f", diagramshapeIdent, diagramshape.Height))
			values.WriteString(fmt.Sprintf("\n\t%s.IsHidden = %t", diagramshapeIdent, diagramshape.IsHidden))
			values.WriteString(fmt.Sprintf("\n\t%s.IsPresentationOverridden = %t", diagramshapeIdent, diagramshape.IsPresentationOverridden))
			values.WriteString(fmt.Sprintf("\n\t%s.Color = %s", diagramshapeIdent, __gong__toRawStringLiteral(diagramshape.Color)))
			values.WriteString(fmt.Sprintf("\n\t%s.FillOpacity = %f", diagramshapeIdent, diagramshape.FillOpacity))
			values.WriteString(fmt.Sprintf("\n\t%s.Stroke = %s", diagramshapeIdent, __gong__toRawStringLiteral(diagramshape.Stroke)))
			values.WriteString(fmt.Sprintf("\n\t%s.StrokeWidth = %f", diagramshapeIdent, diagramshape.StrokeWidth))
			values.WriteString(fmt.Sprintf("\n\t%s.StrokeOpacity = %f", diagramshapeIdent, diagramshape.StrokeOpacity))
			values.WriteString(fmt.Sprintf("\n\t%s.StrokeDashArray = %s", diagramshapeIdent, __gong__toRawStringLiteral(diagramshape.StrokeDashArray)))
			values.WriteString(fmt.Sprintf("\n\t%s.RX = %f", diagramshapeIdent, diagramshape.RX))
			values.WriteString(fmt.Sprintf("\n\t%s.FontSize = %s", diagramshapeIdent, __gong__toRawStringLiteral(diagramshape.FontSize)))
			values.WriteString(fmt.Sprintf("\n\t%s.FontWeight = %s", diagramshapeIdent, __gong__toRawStringLiteral(diagramshape.FontWeight)))
			values.WriteString(fmt.Sprintf("\n\t%s.FontColor = %s", diagramshapeIdent, __gong__toRawStringLiteral(diagramshape.FontColor)))
			if diagramshape.Diagram != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := diagramshape.Diagram.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Diagram = %s", diagramshapeIdent, targetIdent))
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
			values.WriteString(fmt.Sprintf("\n\t%s.IsRootLibrary = %t", libraryIdent, library.IsRootLibrary))
			values.WriteString(fmt.Sprintf("\n\t%s.ComputedPrefix = %s", libraryIdent, __gong__toRawStringLiteral(library.ComputedPrefix)))
			values.WriteString(fmt.Sprintf("\n\t%s.IsExpanded = %t", libraryIdent, library.IsExpanded))
			values.WriteString(fmt.Sprintf("\n\t%s.NbPixPerCharacter = %f", libraryIdent, library.NbPixPerCharacter))
			for _, elem := range library.RootDeliverables {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := elem.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.RootDeliverables = append(%s.RootDeliverables, %s)", libraryIdent, libraryIdent, targetIdent))
			}
			for _, elem := range library.RootConcerns {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := elem.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.RootConcerns = append(%s.RootConcerns, %s)", libraryIdent, libraryIdent, targetIdent))
			}
			for _, elem := range library.RootStakeholders {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := elem.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.RootStakeholders = append(%s.RootStakeholders, %s)", libraryIdent, libraryIdent, targetIdent))
			}
			for _, elem := range library.RootRequirements {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := elem.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.RootRequirements = append(%s.RootRequirements, %s)", libraryIdent, libraryIdent, targetIdent))
			}
			for _, elem := range library.RootViews {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := elem.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.RootViews = append(%s.RootViews, %s)", libraryIdent, libraryIdent, targetIdent))
			}
			for _, elem := range library.RootConcepts {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := elem.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.RootConcepts = append(%s.RootConcepts, %s)", libraryIdent, libraryIdent, targetIdent))
			}
			for _, elem := range library.AnalysisNeeds {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := elem.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.AnalysisNeeds = append(%s.AnalysisNeeds, %s)", libraryIdent, libraryIdent, targetIdent))
			}
			for _, elem := range library.Notes {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := elem.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Notes = append(%s.Notes, %s)", libraryIdent, libraryIdent, targetIdent))
			}
			for _, elem := range library.Diagrams {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := elem.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Diagrams = append(%s.Diagrams, %s)", libraryIdent, libraryIdent, targetIdent))
			}
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
			values.WriteString(fmt.Sprintf("\n\t%s.ComputedPrefix = %s", noteIdent, __gong__toRawStringLiteral(note.ComputedPrefix)))
			values.WriteString(fmt.Sprintf("\n\t%s.IsExpanded = %t", noteIdent, note.IsExpanded))
			for _, elem := range note.Deliverables {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := elem.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Deliverables = append(%s.Deliverables, %s)", noteIdent, noteIdent, targetIdent))
			}
			for _, elem := range note.Tasks {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := elem.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Tasks = append(%s.Tasks, %s)", noteIdent, noteIdent, targetIdent))
			}
			for _, elem := range note.Resources {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := elem.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Resources = append(%s.Resources, %s)", noteIdent, noteIdent, targetIdent))
			}
		}
	}
	if stage != nil {
		for _, notedeliverableshape := range __gong__sortStageSetInstances(stage.NoteDeliverableShapes, stage.NoteDeliverableShape_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			notedeliverableshapeIdent := notedeliverableshape.GongGetIdentifier(stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&models.NoteDeliverableShape{Name: %s}).Stage(stage)", notedeliverableshapeIdent, __gong__toRawStringLiteral(notedeliverableshape.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", notedeliverableshapeIdent, __gong__toRawStringLiteral(notedeliverableshape.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.StartRatio = %f", notedeliverableshapeIdent, notedeliverableshape.StartRatio))
			values.WriteString(fmt.Sprintf("\n\t%s.EndRatio = %f", notedeliverableshapeIdent, notedeliverableshape.EndRatio))
			if notedeliverableshape.StartOrientation.ToCodeString() != "" {
				values.WriteString(fmt.Sprintf("\n\t%s.StartOrientation = models.%s", notedeliverableshapeIdent, notedeliverableshape.StartOrientation.ToCodeString()))
			} else {
				values.WriteString(fmt.Sprintf("\n\t%s.StartOrientation = \"\"", notedeliverableshapeIdent))
			}
			if notedeliverableshape.EndOrientation.ToCodeString() != "" {
				values.WriteString(fmt.Sprintf("\n\t%s.EndOrientation = models.%s", notedeliverableshapeIdent, notedeliverableshape.EndOrientation.ToCodeString()))
			} else {
				values.WriteString(fmt.Sprintf("\n\t%s.EndOrientation = \"\"", notedeliverableshapeIdent))
			}
			values.WriteString(fmt.Sprintf("\n\t%s.CornerOffsetRatio = %f", notedeliverableshapeIdent, notedeliverableshape.CornerOffsetRatio))
			values.WriteString(fmt.Sprintf("\n\t%s.IsHidden = %t", notedeliverableshapeIdent, notedeliverableshape.IsHidden))
			values.WriteString(fmt.Sprintf("\n\t%s.IsPresentationOverridden = %t", notedeliverableshapeIdent, notedeliverableshape.IsPresentationOverridden))
			values.WriteString(fmt.Sprintf("\n\t%s.Stroke = %s", notedeliverableshapeIdent, __gong__toRawStringLiteral(notedeliverableshape.Stroke)))
			values.WriteString(fmt.Sprintf("\n\t%s.StrokeWidth = %f", notedeliverableshapeIdent, notedeliverableshape.StrokeWidth))
			values.WriteString(fmt.Sprintf("\n\t%s.StrokeOpacity = %f", notedeliverableshapeIdent, notedeliverableshape.StrokeOpacity))
			values.WriteString(fmt.Sprintf("\n\t%s.StrokeDashArray = %s", notedeliverableshapeIdent, __gong__toRawStringLiteral(notedeliverableshape.StrokeDashArray)))
			values.WriteString(fmt.Sprintf("\n\t%s.FontSize = %s", notedeliverableshapeIdent, __gong__toRawStringLiteral(notedeliverableshape.FontSize)))
			values.WriteString(fmt.Sprintf("\n\t%s.FontColor = %s", notedeliverableshapeIdent, __gong__toRawStringLiteral(notedeliverableshape.FontColor)))
			if notedeliverableshape.Note != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := notedeliverableshape.Note.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Note = %s", notedeliverableshapeIdent, targetIdent))
			}
			if notedeliverableshape.Deliverable != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := notedeliverableshape.Deliverable.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Deliverable = %s", notedeliverableshapeIdent, targetIdent))
			}
			for _, elem := range notedeliverableshape.ControlPointShapes {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := elem.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.ControlPointShapes = append(%s.ControlPointShapes, %s)", notedeliverableshapeIdent, notedeliverableshapeIdent, targetIdent))
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
			values.WriteString(fmt.Sprintf("\n\t%s.IsExpanded = %t", noteshapeIdent, noteshape.IsExpanded))
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
		for _, notestakeholdershape := range __gong__sortStageSetInstances(stage.NoteStakeholderShapes, stage.NoteStakeholderShape_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			notestakeholdershapeIdent := notestakeholdershape.GongGetIdentifier(stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&models.NoteStakeholderShape{Name: %s}).Stage(stage)", notestakeholdershapeIdent, __gong__toRawStringLiteral(notestakeholdershape.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", notestakeholdershapeIdent, __gong__toRawStringLiteral(notestakeholdershape.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.StartRatio = %f", notestakeholdershapeIdent, notestakeholdershape.StartRatio))
			values.WriteString(fmt.Sprintf("\n\t%s.EndRatio = %f", notestakeholdershapeIdent, notestakeholdershape.EndRatio))
			if notestakeholdershape.StartOrientation.ToCodeString() != "" {
				values.WriteString(fmt.Sprintf("\n\t%s.StartOrientation = models.%s", notestakeholdershapeIdent, notestakeholdershape.StartOrientation.ToCodeString()))
			} else {
				values.WriteString(fmt.Sprintf("\n\t%s.StartOrientation = \"\"", notestakeholdershapeIdent))
			}
			if notestakeholdershape.EndOrientation.ToCodeString() != "" {
				values.WriteString(fmt.Sprintf("\n\t%s.EndOrientation = models.%s", notestakeholdershapeIdent, notestakeholdershape.EndOrientation.ToCodeString()))
			} else {
				values.WriteString(fmt.Sprintf("\n\t%s.EndOrientation = \"\"", notestakeholdershapeIdent))
			}
			values.WriteString(fmt.Sprintf("\n\t%s.CornerOffsetRatio = %f", notestakeholdershapeIdent, notestakeholdershape.CornerOffsetRatio))
			values.WriteString(fmt.Sprintf("\n\t%s.IsHidden = %t", notestakeholdershapeIdent, notestakeholdershape.IsHidden))
			values.WriteString(fmt.Sprintf("\n\t%s.IsPresentationOverridden = %t", notestakeholdershapeIdent, notestakeholdershape.IsPresentationOverridden))
			values.WriteString(fmt.Sprintf("\n\t%s.Stroke = %s", notestakeholdershapeIdent, __gong__toRawStringLiteral(notestakeholdershape.Stroke)))
			values.WriteString(fmt.Sprintf("\n\t%s.StrokeWidth = %f", notestakeholdershapeIdent, notestakeholdershape.StrokeWidth))
			values.WriteString(fmt.Sprintf("\n\t%s.StrokeOpacity = %f", notestakeholdershapeIdent, notestakeholdershape.StrokeOpacity))
			values.WriteString(fmt.Sprintf("\n\t%s.StrokeDashArray = %s", notestakeholdershapeIdent, __gong__toRawStringLiteral(notestakeholdershape.StrokeDashArray)))
			values.WriteString(fmt.Sprintf("\n\t%s.FontSize = %s", notestakeholdershapeIdent, __gong__toRawStringLiteral(notestakeholdershape.FontSize)))
			values.WriteString(fmt.Sprintf("\n\t%s.FontColor = %s", notestakeholdershapeIdent, __gong__toRawStringLiteral(notestakeholdershape.FontColor)))
			if notestakeholdershape.Note != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := notestakeholdershape.Note.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Note = %s", notestakeholdershapeIdent, targetIdent))
			}
			if notestakeholdershape.Stakeholder != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := notestakeholdershape.Stakeholder.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Stakeholder = %s", notestakeholdershapeIdent, targetIdent))
			}
			for _, elem := range notestakeholdershape.ControlPointShapes {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := elem.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.ControlPointShapes = append(%s.ControlPointShapes, %s)", notestakeholdershapeIdent, notestakeholdershapeIdent, targetIdent))
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
			notetaskshapeIdent := notetaskshape.GongGetIdentifier(stage)
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
				targetIdent := notetaskshape.Note.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Note = %s", notetaskshapeIdent, targetIdent))
			}
			if notetaskshape.Task != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := notetaskshape.Task.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Task = %s", notetaskshapeIdent, targetIdent))
			}
			for _, elem := range notetaskshape.ControlPointShapes {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := elem.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.ControlPointShapes = append(%s.ControlPointShapes, %s)", notetaskshapeIdent, notetaskshapeIdent, targetIdent))
			}
		}
	}
	if stage != nil {
		for _, requirement := range __gong__sortStageSetInstances(stage.Requirements, stage.Requirement_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			requirementIdent := requirement.GongGetIdentifier(stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&models.Requirement{Name: %s}).Stage(stage)", requirementIdent, __gong__toRawStringLiteral(requirement.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", requirementIdent, __gong__toRawStringLiteral(requirement.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.ComputedPrefix = %s", requirementIdent, __gong__toRawStringLiteral(requirement.ComputedPrefix)))
			values.WriteString(fmt.Sprintf("\n\t%s.IsExpanded = %t", requirementIdent, requirement.IsExpanded))
			for _, elem := range requirement.SupportLevels {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := elem.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.SupportLevels = append(%s.SupportLevels, %s)", requirementIdent, requirementIdent, targetIdent))
			}
			for _, elem := range requirement.Concerns {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := elem.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Concerns = append(%s.Concerns, %s)", requirementIdent, requirementIdent, targetIdent))
			}
			for _, elem := range requirement.Views {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := elem.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Views = append(%s.Views, %s)", requirementIdent, requirementIdent, targetIdent))
			}
			for _, elem := range requirement.Concepts {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := elem.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Concepts = append(%s.Concepts, %s)", requirementIdent, requirementIdent, targetIdent))
			}
		}
	}
	if stage != nil {
		for _, requirementconceptshape := range __gong__sortStageSetInstances(stage.RequirementConceptShapes, stage.RequirementConceptShape_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			requirementconceptshapeIdent := requirementconceptshape.GongGetIdentifier(stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&models.RequirementConceptShape{Name: %s}).Stage(stage)", requirementconceptshapeIdent, __gong__toRawStringLiteral(requirementconceptshape.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", requirementconceptshapeIdent, __gong__toRawStringLiteral(requirementconceptshape.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.StartRatio = %f", requirementconceptshapeIdent, requirementconceptshape.StartRatio))
			values.WriteString(fmt.Sprintf("\n\t%s.EndRatio = %f", requirementconceptshapeIdent, requirementconceptshape.EndRatio))
			if requirementconceptshape.StartOrientation.ToCodeString() != "" {
				values.WriteString(fmt.Sprintf("\n\t%s.StartOrientation = models.%s", requirementconceptshapeIdent, requirementconceptshape.StartOrientation.ToCodeString()))
			} else {
				values.WriteString(fmt.Sprintf("\n\t%s.StartOrientation = \"\"", requirementconceptshapeIdent))
			}
			if requirementconceptshape.EndOrientation.ToCodeString() != "" {
				values.WriteString(fmt.Sprintf("\n\t%s.EndOrientation = models.%s", requirementconceptshapeIdent, requirementconceptshape.EndOrientation.ToCodeString()))
			} else {
				values.WriteString(fmt.Sprintf("\n\t%s.EndOrientation = \"\"", requirementconceptshapeIdent))
			}
			values.WriteString(fmt.Sprintf("\n\t%s.CornerOffsetRatio = %f", requirementconceptshapeIdent, requirementconceptshape.CornerOffsetRatio))
			values.WriteString(fmt.Sprintf("\n\t%s.IsHidden = %t", requirementconceptshapeIdent, requirementconceptshape.IsHidden))
			values.WriteString(fmt.Sprintf("\n\t%s.IsPresentationOverridden = %t", requirementconceptshapeIdent, requirementconceptshape.IsPresentationOverridden))
			values.WriteString(fmt.Sprintf("\n\t%s.Stroke = %s", requirementconceptshapeIdent, __gong__toRawStringLiteral(requirementconceptshape.Stroke)))
			values.WriteString(fmt.Sprintf("\n\t%s.StrokeWidth = %f", requirementconceptshapeIdent, requirementconceptshape.StrokeWidth))
			values.WriteString(fmt.Sprintf("\n\t%s.StrokeOpacity = %f", requirementconceptshapeIdent, requirementconceptshape.StrokeOpacity))
			values.WriteString(fmt.Sprintf("\n\t%s.StrokeDashArray = %s", requirementconceptshapeIdent, __gong__toRawStringLiteral(requirementconceptshape.StrokeDashArray)))
			values.WriteString(fmt.Sprintf("\n\t%s.FontSize = %s", requirementconceptshapeIdent, __gong__toRawStringLiteral(requirementconceptshape.FontSize)))
			values.WriteString(fmt.Sprintf("\n\t%s.FontColor = %s", requirementconceptshapeIdent, __gong__toRawStringLiteral(requirementconceptshape.FontColor)))
			if requirementconceptshape.Requirement != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := requirementconceptshape.Requirement.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Requirement = %s", requirementconceptshapeIdent, targetIdent))
			}
			if requirementconceptshape.Concept != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := requirementconceptshape.Concept.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Concept = %s", requirementconceptshapeIdent, targetIdent))
			}
			for _, elem := range requirementconceptshape.ControlPointShapes {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := elem.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.ControlPointShapes = append(%s.ControlPointShapes, %s)", requirementconceptshapeIdent, requirementconceptshapeIdent, targetIdent))
			}
		}
	}
	if stage != nil {
		for _, requirementconcernshape := range __gong__sortStageSetInstances(stage.RequirementConcernShapes, stage.RequirementConcernShape_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			requirementconcernshapeIdent := requirementconcernshape.GongGetIdentifier(stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&models.RequirementConcernShape{Name: %s}).Stage(stage)", requirementconcernshapeIdent, __gong__toRawStringLiteral(requirementconcernshape.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", requirementconcernshapeIdent, __gong__toRawStringLiteral(requirementconcernshape.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.StartRatio = %f", requirementconcernshapeIdent, requirementconcernshape.StartRatio))
			values.WriteString(fmt.Sprintf("\n\t%s.EndRatio = %f", requirementconcernshapeIdent, requirementconcernshape.EndRatio))
			if requirementconcernshape.StartOrientation.ToCodeString() != "" {
				values.WriteString(fmt.Sprintf("\n\t%s.StartOrientation = models.%s", requirementconcernshapeIdent, requirementconcernshape.StartOrientation.ToCodeString()))
			} else {
				values.WriteString(fmt.Sprintf("\n\t%s.StartOrientation = \"\"", requirementconcernshapeIdent))
			}
			if requirementconcernshape.EndOrientation.ToCodeString() != "" {
				values.WriteString(fmt.Sprintf("\n\t%s.EndOrientation = models.%s", requirementconcernshapeIdent, requirementconcernshape.EndOrientation.ToCodeString()))
			} else {
				values.WriteString(fmt.Sprintf("\n\t%s.EndOrientation = \"\"", requirementconcernshapeIdent))
			}
			values.WriteString(fmt.Sprintf("\n\t%s.CornerOffsetRatio = %f", requirementconcernshapeIdent, requirementconcernshape.CornerOffsetRatio))
			values.WriteString(fmt.Sprintf("\n\t%s.IsHidden = %t", requirementconcernshapeIdent, requirementconcernshape.IsHidden))
			values.WriteString(fmt.Sprintf("\n\t%s.IsPresentationOverridden = %t", requirementconcernshapeIdent, requirementconcernshape.IsPresentationOverridden))
			values.WriteString(fmt.Sprintf("\n\t%s.Stroke = %s", requirementconcernshapeIdent, __gong__toRawStringLiteral(requirementconcernshape.Stroke)))
			values.WriteString(fmt.Sprintf("\n\t%s.StrokeWidth = %f", requirementconcernshapeIdent, requirementconcernshape.StrokeWidth))
			values.WriteString(fmt.Sprintf("\n\t%s.StrokeOpacity = %f", requirementconcernshapeIdent, requirementconcernshape.StrokeOpacity))
			values.WriteString(fmt.Sprintf("\n\t%s.StrokeDashArray = %s", requirementconcernshapeIdent, __gong__toRawStringLiteral(requirementconcernshape.StrokeDashArray)))
			values.WriteString(fmt.Sprintf("\n\t%s.FontSize = %s", requirementconcernshapeIdent, __gong__toRawStringLiteral(requirementconcernshape.FontSize)))
			values.WriteString(fmt.Sprintf("\n\t%s.FontColor = %s", requirementconcernshapeIdent, __gong__toRawStringLiteral(requirementconcernshape.FontColor)))
			if requirementconcernshape.Requirement != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := requirementconcernshape.Requirement.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Requirement = %s", requirementconcernshapeIdent, targetIdent))
			}
			if requirementconcernshape.Concern != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := requirementconcernshape.Concern.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Concern = %s", requirementconcernshapeIdent, targetIdent))
			}
			for _, elem := range requirementconcernshape.ControlPointShapes {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := elem.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.ControlPointShapes = append(%s.ControlPointShapes, %s)", requirementconcernshapeIdent, requirementconcernshapeIdent, targetIdent))
			}
		}
	}
	if stage != nil {
		for _, requirementshape := range __gong__sortStageSetInstances(stage.RequirementShapes, stage.RequirementShape_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			requirementshapeIdent := requirementshape.GongGetIdentifier(stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&models.RequirementShape{Name: %s}).Stage(stage)", requirementshapeIdent, __gong__toRawStringLiteral(requirementshape.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", requirementshapeIdent, __gong__toRawStringLiteral(requirementshape.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.IsExpanded = %t", requirementshapeIdent, requirementshape.IsExpanded))
			values.WriteString(fmt.Sprintf("\n\t%s.X = %f", requirementshapeIdent, requirementshape.X))
			values.WriteString(fmt.Sprintf("\n\t%s.Y = %f", requirementshapeIdent, requirementshape.Y))
			values.WriteString(fmt.Sprintf("\n\t%s.Width = %f", requirementshapeIdent, requirementshape.Width))
			values.WriteString(fmt.Sprintf("\n\t%s.Height = %f", requirementshapeIdent, requirementshape.Height))
			values.WriteString(fmt.Sprintf("\n\t%s.IsHidden = %t", requirementshapeIdent, requirementshape.IsHidden))
			values.WriteString(fmt.Sprintf("\n\t%s.IsPresentationOverridden = %t", requirementshapeIdent, requirementshape.IsPresentationOverridden))
			values.WriteString(fmt.Sprintf("\n\t%s.Color = %s", requirementshapeIdent, __gong__toRawStringLiteral(requirementshape.Color)))
			values.WriteString(fmt.Sprintf("\n\t%s.FillOpacity = %f", requirementshapeIdent, requirementshape.FillOpacity))
			values.WriteString(fmt.Sprintf("\n\t%s.Stroke = %s", requirementshapeIdent, __gong__toRawStringLiteral(requirementshape.Stroke)))
			values.WriteString(fmt.Sprintf("\n\t%s.StrokeWidth = %f", requirementshapeIdent, requirementshape.StrokeWidth))
			values.WriteString(fmt.Sprintf("\n\t%s.StrokeOpacity = %f", requirementshapeIdent, requirementshape.StrokeOpacity))
			values.WriteString(fmt.Sprintf("\n\t%s.StrokeDashArray = %s", requirementshapeIdent, __gong__toRawStringLiteral(requirementshape.StrokeDashArray)))
			values.WriteString(fmt.Sprintf("\n\t%s.RX = %f", requirementshapeIdent, requirementshape.RX))
			values.WriteString(fmt.Sprintf("\n\t%s.FontSize = %s", requirementshapeIdent, __gong__toRawStringLiteral(requirementshape.FontSize)))
			values.WriteString(fmt.Sprintf("\n\t%s.FontWeight = %s", requirementshapeIdent, __gong__toRawStringLiteral(requirementshape.FontWeight)))
			values.WriteString(fmt.Sprintf("\n\t%s.FontColor = %s", requirementshapeIdent, __gong__toRawStringLiteral(requirementshape.FontColor)))
			if requirementshape.Requirement != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := requirementshape.Requirement.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Requirement = %s", requirementshapeIdent, targetIdent))
			}
		}
	}
	if stage != nil {
		for _, requirementviewshape := range __gong__sortStageSetInstances(stage.RequirementViewShapes, stage.RequirementViewShape_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			requirementviewshapeIdent := requirementviewshape.GongGetIdentifier(stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&models.RequirementViewShape{Name: %s}).Stage(stage)", requirementviewshapeIdent, __gong__toRawStringLiteral(requirementviewshape.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", requirementviewshapeIdent, __gong__toRawStringLiteral(requirementviewshape.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.StartRatio = %f", requirementviewshapeIdent, requirementviewshape.StartRatio))
			values.WriteString(fmt.Sprintf("\n\t%s.EndRatio = %f", requirementviewshapeIdent, requirementviewshape.EndRatio))
			if requirementviewshape.StartOrientation.ToCodeString() != "" {
				values.WriteString(fmt.Sprintf("\n\t%s.StartOrientation = models.%s", requirementviewshapeIdent, requirementviewshape.StartOrientation.ToCodeString()))
			} else {
				values.WriteString(fmt.Sprintf("\n\t%s.StartOrientation = \"\"", requirementviewshapeIdent))
			}
			if requirementviewshape.EndOrientation.ToCodeString() != "" {
				values.WriteString(fmt.Sprintf("\n\t%s.EndOrientation = models.%s", requirementviewshapeIdent, requirementviewshape.EndOrientation.ToCodeString()))
			} else {
				values.WriteString(fmt.Sprintf("\n\t%s.EndOrientation = \"\"", requirementviewshapeIdent))
			}
			values.WriteString(fmt.Sprintf("\n\t%s.CornerOffsetRatio = %f", requirementviewshapeIdent, requirementviewshape.CornerOffsetRatio))
			values.WriteString(fmt.Sprintf("\n\t%s.IsHidden = %t", requirementviewshapeIdent, requirementviewshape.IsHidden))
			values.WriteString(fmt.Sprintf("\n\t%s.IsPresentationOverridden = %t", requirementviewshapeIdent, requirementviewshape.IsPresentationOverridden))
			values.WriteString(fmt.Sprintf("\n\t%s.Stroke = %s", requirementviewshapeIdent, __gong__toRawStringLiteral(requirementviewshape.Stroke)))
			values.WriteString(fmt.Sprintf("\n\t%s.StrokeWidth = %f", requirementviewshapeIdent, requirementviewshape.StrokeWidth))
			values.WriteString(fmt.Sprintf("\n\t%s.StrokeOpacity = %f", requirementviewshapeIdent, requirementviewshape.StrokeOpacity))
			values.WriteString(fmt.Sprintf("\n\t%s.StrokeDashArray = %s", requirementviewshapeIdent, __gong__toRawStringLiteral(requirementviewshape.StrokeDashArray)))
			values.WriteString(fmt.Sprintf("\n\t%s.FontSize = %s", requirementviewshapeIdent, __gong__toRawStringLiteral(requirementviewshape.FontSize)))
			values.WriteString(fmt.Sprintf("\n\t%s.FontColor = %s", requirementviewshapeIdent, __gong__toRawStringLiteral(requirementviewshape.FontColor)))
			if requirementviewshape.Requirement != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := requirementviewshape.Requirement.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Requirement = %s", requirementviewshapeIdent, targetIdent))
			}
			if requirementviewshape.View != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := requirementviewshape.View.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.View = %s", requirementviewshapeIdent, targetIdent))
			}
			for _, elem := range requirementviewshape.ControlPointShapes {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := elem.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.ControlPointShapes = append(%s.ControlPointShapes, %s)", requirementviewshapeIdent, requirementviewshapeIdent, targetIdent))
			}
		}
	}
	if stage != nil {
		for _, stakeholder := range __gong__sortStageSetInstances(stage.Stakeholders, stage.Stakeholder_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			stakeholderIdent := stakeholder.GongGetIdentifier(stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&models.Stakeholder{Name: %s}).Stage(stage)", stakeholderIdent, __gong__toRawStringLiteral(stakeholder.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", stakeholderIdent, __gong__toRawStringLiteral(stakeholder.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.IDAirbus = %s", stakeholderIdent, __gong__toRawStringLiteral(stakeholder.IDAirbus)))
			values.WriteString(fmt.Sprintf("\n\t%s.ComputedPrefix = %s", stakeholderIdent, __gong__toRawStringLiteral(stakeholder.ComputedPrefix)))
			values.WriteString(fmt.Sprintf("\n\t%s.IsExpanded = %t", stakeholderIdent, stakeholder.IsExpanded))
			values.WriteString(fmt.Sprintf("\n\t%s.Description = %s", stakeholderIdent, __gong__toRawStringLiteral(stakeholder.Description)))
			for _, elem := range stakeholder.Concerns {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := elem.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Concerns = append(%s.Concerns, %s)", stakeholderIdent, stakeholderIdent, targetIdent))
			}
			for _, elem := range stakeholder.SubStakeholders {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := elem.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.SubStakeholders = append(%s.SubStakeholders, %s)", stakeholderIdent, stakeholderIdent, targetIdent))
			}
		}
	}
	if stage != nil {
		for _, stakeholdercompositionshape := range __gong__sortStageSetInstances(stage.StakeholderCompositionShapes, stage.StakeholderCompositionShape_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			stakeholdercompositionshapeIdent := stakeholdercompositionshape.GongGetIdentifier(stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&models.StakeholderCompositionShape{Name: %s}).Stage(stage)", stakeholdercompositionshapeIdent, __gong__toRawStringLiteral(stakeholdercompositionshape.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", stakeholdercompositionshapeIdent, __gong__toRawStringLiteral(stakeholdercompositionshape.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.StartRatio = %f", stakeholdercompositionshapeIdent, stakeholdercompositionshape.StartRatio))
			values.WriteString(fmt.Sprintf("\n\t%s.EndRatio = %f", stakeholdercompositionshapeIdent, stakeholdercompositionshape.EndRatio))
			if stakeholdercompositionshape.StartOrientation.ToCodeString() != "" {
				values.WriteString(fmt.Sprintf("\n\t%s.StartOrientation = models.%s", stakeholdercompositionshapeIdent, stakeholdercompositionshape.StartOrientation.ToCodeString()))
			} else {
				values.WriteString(fmt.Sprintf("\n\t%s.StartOrientation = \"\"", stakeholdercompositionshapeIdent))
			}
			if stakeholdercompositionshape.EndOrientation.ToCodeString() != "" {
				values.WriteString(fmt.Sprintf("\n\t%s.EndOrientation = models.%s", stakeholdercompositionshapeIdent, stakeholdercompositionshape.EndOrientation.ToCodeString()))
			} else {
				values.WriteString(fmt.Sprintf("\n\t%s.EndOrientation = \"\"", stakeholdercompositionshapeIdent))
			}
			values.WriteString(fmt.Sprintf("\n\t%s.CornerOffsetRatio = %f", stakeholdercompositionshapeIdent, stakeholdercompositionshape.CornerOffsetRatio))
			values.WriteString(fmt.Sprintf("\n\t%s.IsHidden = %t", stakeholdercompositionshapeIdent, stakeholdercompositionshape.IsHidden))
			values.WriteString(fmt.Sprintf("\n\t%s.IsPresentationOverridden = %t", stakeholdercompositionshapeIdent, stakeholdercompositionshape.IsPresentationOverridden))
			values.WriteString(fmt.Sprintf("\n\t%s.Stroke = %s", stakeholdercompositionshapeIdent, __gong__toRawStringLiteral(stakeholdercompositionshape.Stroke)))
			values.WriteString(fmt.Sprintf("\n\t%s.StrokeWidth = %f", stakeholdercompositionshapeIdent, stakeholdercompositionshape.StrokeWidth))
			values.WriteString(fmt.Sprintf("\n\t%s.StrokeOpacity = %f", stakeholdercompositionshapeIdent, stakeholdercompositionshape.StrokeOpacity))
			values.WriteString(fmt.Sprintf("\n\t%s.StrokeDashArray = %s", stakeholdercompositionshapeIdent, __gong__toRawStringLiteral(stakeholdercompositionshape.StrokeDashArray)))
			values.WriteString(fmt.Sprintf("\n\t%s.FontSize = %s", stakeholdercompositionshapeIdent, __gong__toRawStringLiteral(stakeholdercompositionshape.FontSize)))
			values.WriteString(fmt.Sprintf("\n\t%s.FontColor = %s", stakeholdercompositionshapeIdent, __gong__toRawStringLiteral(stakeholdercompositionshape.FontColor)))
			if stakeholdercompositionshape.Stakeholder != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := stakeholdercompositionshape.Stakeholder.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Stakeholder = %s", stakeholdercompositionshapeIdent, targetIdent))
			}
			for _, elem := range stakeholdercompositionshape.ControlPointShapes {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := elem.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.ControlPointShapes = append(%s.ControlPointShapes, %s)", stakeholdercompositionshapeIdent, stakeholdercompositionshapeIdent, targetIdent))
			}
		}
	}
	if stage != nil {
		for _, stakeholderconcernshape := range __gong__sortStageSetInstances(stage.StakeholderConcernShapes, stage.StakeholderConcernShape_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			stakeholderconcernshapeIdent := stakeholderconcernshape.GongGetIdentifier(stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&models.StakeholderConcernShape{Name: %s}).Stage(stage)", stakeholderconcernshapeIdent, __gong__toRawStringLiteral(stakeholderconcernshape.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", stakeholderconcernshapeIdent, __gong__toRawStringLiteral(stakeholderconcernshape.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.StartRatio = %f", stakeholderconcernshapeIdent, stakeholderconcernshape.StartRatio))
			values.WriteString(fmt.Sprintf("\n\t%s.EndRatio = %f", stakeholderconcernshapeIdent, stakeholderconcernshape.EndRatio))
			if stakeholderconcernshape.StartOrientation.ToCodeString() != "" {
				values.WriteString(fmt.Sprintf("\n\t%s.StartOrientation = models.%s", stakeholderconcernshapeIdent, stakeholderconcernshape.StartOrientation.ToCodeString()))
			} else {
				values.WriteString(fmt.Sprintf("\n\t%s.StartOrientation = \"\"", stakeholderconcernshapeIdent))
			}
			if stakeholderconcernshape.EndOrientation.ToCodeString() != "" {
				values.WriteString(fmt.Sprintf("\n\t%s.EndOrientation = models.%s", stakeholderconcernshapeIdent, stakeholderconcernshape.EndOrientation.ToCodeString()))
			} else {
				values.WriteString(fmt.Sprintf("\n\t%s.EndOrientation = \"\"", stakeholderconcernshapeIdent))
			}
			values.WriteString(fmt.Sprintf("\n\t%s.CornerOffsetRatio = %f", stakeholderconcernshapeIdent, stakeholderconcernshape.CornerOffsetRatio))
			values.WriteString(fmt.Sprintf("\n\t%s.IsHidden = %t", stakeholderconcernshapeIdent, stakeholderconcernshape.IsHidden))
			values.WriteString(fmt.Sprintf("\n\t%s.IsPresentationOverridden = %t", stakeholderconcernshapeIdent, stakeholderconcernshape.IsPresentationOverridden))
			values.WriteString(fmt.Sprintf("\n\t%s.Stroke = %s", stakeholderconcernshapeIdent, __gong__toRawStringLiteral(stakeholderconcernshape.Stroke)))
			values.WriteString(fmt.Sprintf("\n\t%s.StrokeWidth = %f", stakeholderconcernshapeIdent, stakeholderconcernshape.StrokeWidth))
			values.WriteString(fmt.Sprintf("\n\t%s.StrokeOpacity = %f", stakeholderconcernshapeIdent, stakeholderconcernshape.StrokeOpacity))
			values.WriteString(fmt.Sprintf("\n\t%s.StrokeDashArray = %s", stakeholderconcernshapeIdent, __gong__toRawStringLiteral(stakeholderconcernshape.StrokeDashArray)))
			values.WriteString(fmt.Sprintf("\n\t%s.FontSize = %s", stakeholderconcernshapeIdent, __gong__toRawStringLiteral(stakeholderconcernshape.FontSize)))
			values.WriteString(fmt.Sprintf("\n\t%s.FontColor = %s", stakeholderconcernshapeIdent, __gong__toRawStringLiteral(stakeholderconcernshape.FontColor)))
			if stakeholderconcernshape.Stakeholder != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := stakeholderconcernshape.Stakeholder.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Stakeholder = %s", stakeholderconcernshapeIdent, targetIdent))
			}
			if stakeholderconcernshape.Concern != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := stakeholderconcernshape.Concern.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Concern = %s", stakeholderconcernshapeIdent, targetIdent))
			}
			for _, elem := range stakeholderconcernshape.ControlPointShapes {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := elem.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.ControlPointShapes = append(%s.ControlPointShapes, %s)", stakeholderconcernshapeIdent, stakeholderconcernshapeIdent, targetIdent))
			}
		}
	}
	if stage != nil {
		for _, stakeholdershape := range __gong__sortStageSetInstances(stage.StakeholderShapes, stage.StakeholderShape_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			stakeholdershapeIdent := stakeholdershape.GongGetIdentifier(stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&models.StakeholderShape{Name: %s}).Stage(stage)", stakeholdershapeIdent, __gong__toRawStringLiteral(stakeholdershape.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", stakeholdershapeIdent, __gong__toRawStringLiteral(stakeholdershape.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.IsExpanded = %t", stakeholdershapeIdent, stakeholdershape.IsExpanded))
			values.WriteString(fmt.Sprintf("\n\t%s.X = %f", stakeholdershapeIdent, stakeholdershape.X))
			values.WriteString(fmt.Sprintf("\n\t%s.Y = %f", stakeholdershapeIdent, stakeholdershape.Y))
			values.WriteString(fmt.Sprintf("\n\t%s.Width = %f", stakeholdershapeIdent, stakeholdershape.Width))
			values.WriteString(fmt.Sprintf("\n\t%s.Height = %f", stakeholdershapeIdent, stakeholdershape.Height))
			values.WriteString(fmt.Sprintf("\n\t%s.IsHidden = %t", stakeholdershapeIdent, stakeholdershape.IsHidden))
			values.WriteString(fmt.Sprintf("\n\t%s.IsPresentationOverridden = %t", stakeholdershapeIdent, stakeholdershape.IsPresentationOverridden))
			values.WriteString(fmt.Sprintf("\n\t%s.Color = %s", stakeholdershapeIdent, __gong__toRawStringLiteral(stakeholdershape.Color)))
			values.WriteString(fmt.Sprintf("\n\t%s.FillOpacity = %f", stakeholdershapeIdent, stakeholdershape.FillOpacity))
			values.WriteString(fmt.Sprintf("\n\t%s.Stroke = %s", stakeholdershapeIdent, __gong__toRawStringLiteral(stakeholdershape.Stroke)))
			values.WriteString(fmt.Sprintf("\n\t%s.StrokeWidth = %f", stakeholdershapeIdent, stakeholdershape.StrokeWidth))
			values.WriteString(fmt.Sprintf("\n\t%s.StrokeOpacity = %f", stakeholdershapeIdent, stakeholdershape.StrokeOpacity))
			values.WriteString(fmt.Sprintf("\n\t%s.StrokeDashArray = %s", stakeholdershapeIdent, __gong__toRawStringLiteral(stakeholdershape.StrokeDashArray)))
			values.WriteString(fmt.Sprintf("\n\t%s.RX = %f", stakeholdershapeIdent, stakeholdershape.RX))
			values.WriteString(fmt.Sprintf("\n\t%s.FontSize = %s", stakeholdershapeIdent, __gong__toRawStringLiteral(stakeholdershape.FontSize)))
			values.WriteString(fmt.Sprintf("\n\t%s.FontWeight = %s", stakeholdershapeIdent, __gong__toRawStringLiteral(stakeholdershape.FontWeight)))
			values.WriteString(fmt.Sprintf("\n\t%s.FontColor = %s", stakeholdershapeIdent, __gong__toRawStringLiteral(stakeholdershape.FontColor)))
			if stakeholdershape.Stakeholder != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := stakeholdershape.Stakeholder.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Stakeholder = %s", stakeholdershapeIdent, targetIdent))
			}
		}
	}
	if stage != nil {
		for _, supportlevel := range __gong__sortStageSetInstances(stage.SupportLevels, stage.SupportLevel_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			supportlevelIdent := supportlevel.GongGetIdentifier(stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&models.SupportLevel{Name: %s}).Stage(stage)", supportlevelIdent, __gong__toRawStringLiteral(supportlevel.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", supportlevelIdent, __gong__toRawStringLiteral(supportlevel.Name)))
			if supportlevel.Tool != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := supportlevel.Tool.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Tool = %s", supportlevelIdent, targetIdent))
			}
		}
	}
	if stage != nil {
		for _, tool := range __gong__sortStageSetInstances(stage.Tools, stage.Tool_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			toolIdent := tool.GongGetIdentifier(stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&models.Tool{Name: %s}).Stage(stage)", toolIdent, __gong__toRawStringLiteral(tool.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", toolIdent, __gong__toRawStringLiteral(tool.Name)))
		}
	}
	if stage != nil {
		for _, view := range __gong__sortStageSetInstances(stage.Views, stage.View_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			viewIdent := view.GongGetIdentifier(stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&models.View{Name: %s}).Stage(stage)", viewIdent, __gong__toRawStringLiteral(view.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", viewIdent, __gong__toRawStringLiteral(view.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.ComputedPrefix = %s", viewIdent, __gong__toRawStringLiteral(view.ComputedPrefix)))
			values.WriteString(fmt.Sprintf("\n\t%s.IsExpanded = %t", viewIdent, view.IsExpanded))
			values.WriteString(fmt.Sprintf("\n\t%s.Description = %s", viewIdent, __gong__toRawStringLiteral(view.Description)))
			for _, elem := range view.Concepts {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := elem.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Concepts = append(%s.Concepts, %s)", viewIdent, viewIdent, targetIdent))
			}
		}
	}
	if stage != nil {
		for _, viewconceptshape := range __gong__sortStageSetInstances(stage.ViewConceptShapes, stage.ViewConceptShape_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			viewconceptshapeIdent := viewconceptshape.GongGetIdentifier(stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&models.ViewConceptShape{Name: %s}).Stage(stage)", viewconceptshapeIdent, __gong__toRawStringLiteral(viewconceptshape.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", viewconceptshapeIdent, __gong__toRawStringLiteral(viewconceptshape.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.StartRatio = %f", viewconceptshapeIdent, viewconceptshape.StartRatio))
			values.WriteString(fmt.Sprintf("\n\t%s.EndRatio = %f", viewconceptshapeIdent, viewconceptshape.EndRatio))
			if viewconceptshape.StartOrientation.ToCodeString() != "" {
				values.WriteString(fmt.Sprintf("\n\t%s.StartOrientation = models.%s", viewconceptshapeIdent, viewconceptshape.StartOrientation.ToCodeString()))
			} else {
				values.WriteString(fmt.Sprintf("\n\t%s.StartOrientation = \"\"", viewconceptshapeIdent))
			}
			if viewconceptshape.EndOrientation.ToCodeString() != "" {
				values.WriteString(fmt.Sprintf("\n\t%s.EndOrientation = models.%s", viewconceptshapeIdent, viewconceptshape.EndOrientation.ToCodeString()))
			} else {
				values.WriteString(fmt.Sprintf("\n\t%s.EndOrientation = \"\"", viewconceptshapeIdent))
			}
			values.WriteString(fmt.Sprintf("\n\t%s.CornerOffsetRatio = %f", viewconceptshapeIdent, viewconceptshape.CornerOffsetRatio))
			values.WriteString(fmt.Sprintf("\n\t%s.IsHidden = %t", viewconceptshapeIdent, viewconceptshape.IsHidden))
			values.WriteString(fmt.Sprintf("\n\t%s.IsPresentationOverridden = %t", viewconceptshapeIdent, viewconceptshape.IsPresentationOverridden))
			values.WriteString(fmt.Sprintf("\n\t%s.Stroke = %s", viewconceptshapeIdent, __gong__toRawStringLiteral(viewconceptshape.Stroke)))
			values.WriteString(fmt.Sprintf("\n\t%s.StrokeWidth = %f", viewconceptshapeIdent, viewconceptshape.StrokeWidth))
			values.WriteString(fmt.Sprintf("\n\t%s.StrokeOpacity = %f", viewconceptshapeIdent, viewconceptshape.StrokeOpacity))
			values.WriteString(fmt.Sprintf("\n\t%s.StrokeDashArray = %s", viewconceptshapeIdent, __gong__toRawStringLiteral(viewconceptshape.StrokeDashArray)))
			values.WriteString(fmt.Sprintf("\n\t%s.FontSize = %s", viewconceptshapeIdent, __gong__toRawStringLiteral(viewconceptshape.FontSize)))
			values.WriteString(fmt.Sprintf("\n\t%s.FontColor = %s", viewconceptshapeIdent, __gong__toRawStringLiteral(viewconceptshape.FontColor)))
			if viewconceptshape.View != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := viewconceptshape.View.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.View = %s", viewconceptshapeIdent, targetIdent))
			}
			if viewconceptshape.Concept != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := viewconceptshape.Concept.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Concept = %s", viewconceptshapeIdent, targetIdent))
			}
			for _, elem := range viewconceptshape.ControlPointShapes {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := elem.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.ControlPointShapes = append(%s.ControlPointShapes, %s)", viewconceptshapeIdent, viewconceptshapeIdent, targetIdent))
			}
		}
	}
	if stage != nil {
		for _, viewshape := range __gong__sortStageSetInstances(stage.ViewShapes, stage.ViewShape_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			viewshapeIdent := viewshape.GongGetIdentifier(stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&models.ViewShape{Name: %s}).Stage(stage)", viewshapeIdent, __gong__toRawStringLiteral(viewshape.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", viewshapeIdent, __gong__toRawStringLiteral(viewshape.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.IsExpanded = %t", viewshapeIdent, viewshape.IsExpanded))
			values.WriteString(fmt.Sprintf("\n\t%s.X = %f", viewshapeIdent, viewshape.X))
			values.WriteString(fmt.Sprintf("\n\t%s.Y = %f", viewshapeIdent, viewshape.Y))
			values.WriteString(fmt.Sprintf("\n\t%s.Width = %f", viewshapeIdent, viewshape.Width))
			values.WriteString(fmt.Sprintf("\n\t%s.Height = %f", viewshapeIdent, viewshape.Height))
			values.WriteString(fmt.Sprintf("\n\t%s.IsHidden = %t", viewshapeIdent, viewshape.IsHidden))
			values.WriteString(fmt.Sprintf("\n\t%s.IsPresentationOverridden = %t", viewshapeIdent, viewshape.IsPresentationOverridden))
			values.WriteString(fmt.Sprintf("\n\t%s.Color = %s", viewshapeIdent, __gong__toRawStringLiteral(viewshape.Color)))
			values.WriteString(fmt.Sprintf("\n\t%s.FillOpacity = %f", viewshapeIdent, viewshape.FillOpacity))
			values.WriteString(fmt.Sprintf("\n\t%s.Stroke = %s", viewshapeIdent, __gong__toRawStringLiteral(viewshape.Stroke)))
			values.WriteString(fmt.Sprintf("\n\t%s.StrokeWidth = %f", viewshapeIdent, viewshape.StrokeWidth))
			values.WriteString(fmt.Sprintf("\n\t%s.StrokeOpacity = %f", viewshapeIdent, viewshape.StrokeOpacity))
			values.WriteString(fmt.Sprintf("\n\t%s.StrokeDashArray = %s", viewshapeIdent, __gong__toRawStringLiteral(viewshape.StrokeDashArray)))
			values.WriteString(fmt.Sprintf("\n\t%s.RX = %f", viewshapeIdent, viewshape.RX))
			values.WriteString(fmt.Sprintf("\n\t%s.FontSize = %s", viewshapeIdent, __gong__toRawStringLiteral(viewshape.FontSize)))
			values.WriteString(fmt.Sprintf("\n\t%s.FontWeight = %s", viewshapeIdent, __gong__toRawStringLiteral(viewshape.FontWeight)))
			values.WriteString(fmt.Sprintf("\n\t%s.FontColor = %s", viewshapeIdent, __gong__toRawStringLiteral(viewshape.FontColor)))
			if viewshape.View != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := viewshape.View.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.View = %s", viewshapeIdent, targetIdent))
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
go run github.com/fullstack-lang/gong/dsm/capture/go/models/probe/cmd/probe@%s stage.go
*/
package %s

import (
	"slices"
	"time"

	"github.com/fullstack-lang/gong/dsm/capture/go/models"%s
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

	if modelsPackageName != "" && modelsPackageName != "github.com/fullstack-lang/gong/dsm/capture/go/models" {
		res = strings.ReplaceAll(res, "\"github.com/fullstack-lang/gong/dsm/capture/go/models\"", fmt.Sprintf("%q", modelsPackageName))
	}

	return res, nil
}

// insertion point for marshall field methods
func (analysisneed *AnalysisNeed) GongMarshallField(stage *Stage, fieldName string) (res string) {
	ident := analysisneed.GongGetIdentifier(stage)
	_ = ident
	switch fieldName {
	case "Name":
		res = __gong__marshallString(ident, "Name", analysisneed.Name)
	case "ComputedPrefix":
		res = __gong__marshallString(ident, "ComputedPrefix", analysisneed.ComputedPrefix)
	case "IsExpanded":
		res = __gong__marshallBool(ident, "IsExpanded", analysisneed.IsExpanded)

	default:
		log.Panicf("Unknown field %s for Gongstruct AnalysisNeed", fieldName)
	}
	return
}

func (concept *Concept) GongMarshallField(stage *Stage, fieldName string) (res string) {
	ident := concept.GongGetIdentifier(stage)
	_ = ident
	switch fieldName {
	case "Name":
		res = __gong__marshallString(ident, "Name", concept.Name)
	case "ComputedPrefix":
		res = __gong__marshallString(ident, "ComputedPrefix", concept.ComputedPrefix)
	case "IsExpanded":
		res = __gong__marshallBool(ident, "IsExpanded", concept.IsExpanded)

	case "Tools":
		var sb strings.Builder
		for _, _tool := range concept.Tools {
			sb.WriteString(__gong__marshallSliceOfPointers(ident, "Tools", _tool.GongGetIdentifier(stage)))
		}
		res = sb.String()
	default:
		log.Panicf("Unknown field %s for Gongstruct Concept", fieldName)
	}
	return
}

func (conceptshape *ConceptShape) GongMarshallField(stage *Stage, fieldName string) (res string) {
	ident := conceptshape.GongGetIdentifier(stage)
	_ = ident
	switch fieldName {
	case "Name":
		res = __gong__marshallString(ident, "Name", conceptshape.Name)
	case "IsExpanded":
		res = __gong__marshallBool(ident, "IsExpanded", conceptshape.IsExpanded)
	case "X":
		res = __gong__marshallFloat(ident, "X", conceptshape.X)
	case "Y":
		res = __gong__marshallFloat(ident, "Y", conceptshape.Y)
	case "Width":
		res = __gong__marshallFloat(ident, "Width", conceptshape.Width)
	case "Height":
		res = __gong__marshallFloat(ident, "Height", conceptshape.Height)
	case "IsHidden":
		res = __gong__marshallBool(ident, "IsHidden", conceptshape.IsHidden)
	case "IsPresentationOverridden":
		res = __gong__marshallBool(ident, "IsPresentationOverridden", conceptshape.IsPresentationOverridden)
	case "Color":
		res = __gong__marshallString(ident, "Color", conceptshape.Color)
	case "FillOpacity":
		res = __gong__marshallFloat(ident, "FillOpacity", conceptshape.FillOpacity)
	case "Stroke":
		res = __gong__marshallString(ident, "Stroke", conceptshape.Stroke)
	case "StrokeWidth":
		res = __gong__marshallFloat(ident, "StrokeWidth", conceptshape.StrokeWidth)
	case "StrokeOpacity":
		res = __gong__marshallFloat(ident, "StrokeOpacity", conceptshape.StrokeOpacity)
	case "StrokeDashArray":
		res = __gong__marshallString(ident, "StrokeDashArray", conceptshape.StrokeDashArray)
	case "RX":
		res = __gong__marshallFloat(ident, "RX", conceptshape.RX)
	case "FontSize":
		res = __gong__marshallString(ident, "FontSize", conceptshape.FontSize)
	case "FontWeight":
		res = __gong__marshallString(ident, "FontWeight", conceptshape.FontWeight)
	case "FontColor":
		res = __gong__marshallString(ident, "FontColor", conceptshape.FontColor)

	case "Concept":
		if conceptshape.Concept != nil {
			res = __gong__marshallPointer(ident, "Concept", conceptshape.Concept.GongGetIdentifier(stage))
		} else {
			res = __gong__marshallPointer(ident, "Concept", "nil")
		}
	default:
		log.Panicf("Unknown field %s for Gongstruct ConceptShape", fieldName)
	}
	return
}

func (concern *Concern) GongMarshallField(stage *Stage, fieldName string) (res string) {
	ident := concern.GongGetIdentifier(stage)
	_ = ident
	switch fieldName {
	case "Name":
		res = __gong__marshallString(ident, "Name", concern.Name)
	case "IDAirbus":
		res = __gong__marshallString(ident, "IDAirbus", concern.IDAirbus)
	case "Priority":
		res = __gong__marshallEnumString(ident, "Priority", concern.Priority.ToCodeString())
	case "ComputedPrefix":
		res = __gong__marshallString(ident, "ComputedPrefix", concern.ComputedPrefix)
	case "IsExpanded":
		res = __gong__marshallBool(ident, "IsExpanded", concern.IsExpanded)
	case "Description":
		res = __gong__marshallString(ident, "Description", concern.Description)
	case "IsInputsNodeExpanded":
		res = __gong__marshallBool(ident, "IsInputsNodeExpanded", concern.IsInputsNodeExpanded)
	case "IsOutputsNodeExpanded":
		res = __gong__marshallBool(ident, "IsOutputsNodeExpanded", concern.IsOutputsNodeExpanded)
	case "IsWithCompletion":
		res = __gong__marshallBool(ident, "IsWithCompletion", concern.IsWithCompletion)
	case "Completion":
		res = __gong__marshallEnumString(ident, "Completion", concern.Completion.ToCodeString())

	case "SubConcerns":
		var sb strings.Builder
		for _, _concern := range concern.SubConcerns {
			sb.WriteString(__gong__marshallSliceOfPointers(ident, "SubConcerns", _concern.GongGetIdentifier(stage)))
		}
		res = sb.String()
	case "Inputs":
		var sb strings.Builder
		for _, _deliverable := range concern.Inputs {
			sb.WriteString(__gong__marshallSliceOfPointers(ident, "Inputs", _deliverable.GongGetIdentifier(stage)))
		}
		res = sb.String()
	case "Outputs":
		var sb strings.Builder
		for _, _deliverable := range concern.Outputs {
			sb.WriteString(__gong__marshallSliceOfPointers(ident, "Outputs", _deliverable.GongGetIdentifier(stage)))
		}
		res = sb.String()
	default:
		log.Panicf("Unknown field %s for Gongstruct Concern", fieldName)
	}
	return
}

func (concerncompositionshape *ConcernCompositionShape) GongMarshallField(stage *Stage, fieldName string) (res string) {
	ident := concerncompositionshape.GongGetIdentifier(stage)
	_ = ident
	switch fieldName {
	case "Name":
		res = __gong__marshallString(ident, "Name", concerncompositionshape.Name)
	case "StartRatio":
		res = __gong__marshallFloat(ident, "StartRatio", concerncompositionshape.StartRatio)
	case "EndRatio":
		res = __gong__marshallFloat(ident, "EndRatio", concerncompositionshape.EndRatio)
	case "StartOrientation":
		res = __gong__marshallEnumString(ident, "StartOrientation", concerncompositionshape.StartOrientation.ToCodeString())
	case "EndOrientation":
		res = __gong__marshallEnumString(ident, "EndOrientation", concerncompositionshape.EndOrientation.ToCodeString())
	case "CornerOffsetRatio":
		res = __gong__marshallFloat(ident, "CornerOffsetRatio", concerncompositionshape.CornerOffsetRatio)
	case "IsHidden":
		res = __gong__marshallBool(ident, "IsHidden", concerncompositionshape.IsHidden)
	case "IsPresentationOverridden":
		res = __gong__marshallBool(ident, "IsPresentationOverridden", concerncompositionshape.IsPresentationOverridden)
	case "Stroke":
		res = __gong__marshallString(ident, "Stroke", concerncompositionshape.Stroke)
	case "StrokeWidth":
		res = __gong__marshallFloat(ident, "StrokeWidth", concerncompositionshape.StrokeWidth)
	case "StrokeOpacity":
		res = __gong__marshallFloat(ident, "StrokeOpacity", concerncompositionshape.StrokeOpacity)
	case "StrokeDashArray":
		res = __gong__marshallString(ident, "StrokeDashArray", concerncompositionshape.StrokeDashArray)
	case "FontSize":
		res = __gong__marshallString(ident, "FontSize", concerncompositionshape.FontSize)
	case "FontColor":
		res = __gong__marshallString(ident, "FontColor", concerncompositionshape.FontColor)

	case "Concern":
		if concerncompositionshape.Concern != nil {
			res = __gong__marshallPointer(ident, "Concern", concerncompositionshape.Concern.GongGetIdentifier(stage))
		} else {
			res = __gong__marshallPointer(ident, "Concern", "nil")
		}
	case "ControlPointShapes":
		var sb strings.Builder
		for _, _controlpointshape := range concerncompositionshape.ControlPointShapes {
			sb.WriteString(__gong__marshallSliceOfPointers(ident, "ControlPointShapes", _controlpointshape.GongGetIdentifier(stage)))
		}
		res = sb.String()
	default:
		log.Panicf("Unknown field %s for Gongstruct ConcernCompositionShape", fieldName)
	}
	return
}

func (concerninputshape *ConcernInputShape) GongMarshallField(stage *Stage, fieldName string) (res string) {
	ident := concerninputshape.GongGetIdentifier(stage)
	_ = ident
	switch fieldName {
	case "Name":
		res = __gong__marshallString(ident, "Name", concerninputshape.Name)
	case "StartRatio":
		res = __gong__marshallFloat(ident, "StartRatio", concerninputshape.StartRatio)
	case "EndRatio":
		res = __gong__marshallFloat(ident, "EndRatio", concerninputshape.EndRatio)
	case "StartOrientation":
		res = __gong__marshallEnumString(ident, "StartOrientation", concerninputshape.StartOrientation.ToCodeString())
	case "EndOrientation":
		res = __gong__marshallEnumString(ident, "EndOrientation", concerninputshape.EndOrientation.ToCodeString())
	case "CornerOffsetRatio":
		res = __gong__marshallFloat(ident, "CornerOffsetRatio", concerninputshape.CornerOffsetRatio)
	case "IsHidden":
		res = __gong__marshallBool(ident, "IsHidden", concerninputshape.IsHidden)
	case "IsPresentationOverridden":
		res = __gong__marshallBool(ident, "IsPresentationOverridden", concerninputshape.IsPresentationOverridden)
	case "Stroke":
		res = __gong__marshallString(ident, "Stroke", concerninputshape.Stroke)
	case "StrokeWidth":
		res = __gong__marshallFloat(ident, "StrokeWidth", concerninputshape.StrokeWidth)
	case "StrokeOpacity":
		res = __gong__marshallFloat(ident, "StrokeOpacity", concerninputshape.StrokeOpacity)
	case "StrokeDashArray":
		res = __gong__marshallString(ident, "StrokeDashArray", concerninputshape.StrokeDashArray)
	case "FontSize":
		res = __gong__marshallString(ident, "FontSize", concerninputshape.FontSize)
	case "FontColor":
		res = __gong__marshallString(ident, "FontColor", concerninputshape.FontColor)

	case "Deliverable":
		if concerninputshape.Deliverable != nil {
			res = __gong__marshallPointer(ident, "Deliverable", concerninputshape.Deliverable.GongGetIdentifier(stage))
		} else {
			res = __gong__marshallPointer(ident, "Deliverable", "nil")
		}
	case "Concern":
		if concerninputshape.Concern != nil {
			res = __gong__marshallPointer(ident, "Concern", concerninputshape.Concern.GongGetIdentifier(stage))
		} else {
			res = __gong__marshallPointer(ident, "Concern", "nil")
		}
	case "ControlPointShapes":
		var sb strings.Builder
		for _, _controlpointshape := range concerninputshape.ControlPointShapes {
			sb.WriteString(__gong__marshallSliceOfPointers(ident, "ControlPointShapes", _controlpointshape.GongGetIdentifier(stage)))
		}
		res = sb.String()
	default:
		log.Panicf("Unknown field %s for Gongstruct ConcernInputShape", fieldName)
	}
	return
}

func (concernoutputshape *ConcernOutputShape) GongMarshallField(stage *Stage, fieldName string) (res string) {
	ident := concernoutputshape.GongGetIdentifier(stage)
	_ = ident
	switch fieldName {
	case "Name":
		res = __gong__marshallString(ident, "Name", concernoutputshape.Name)
	case "StartRatio":
		res = __gong__marshallFloat(ident, "StartRatio", concernoutputshape.StartRatio)
	case "EndRatio":
		res = __gong__marshallFloat(ident, "EndRatio", concernoutputshape.EndRatio)
	case "StartOrientation":
		res = __gong__marshallEnumString(ident, "StartOrientation", concernoutputshape.StartOrientation.ToCodeString())
	case "EndOrientation":
		res = __gong__marshallEnumString(ident, "EndOrientation", concernoutputshape.EndOrientation.ToCodeString())
	case "CornerOffsetRatio":
		res = __gong__marshallFloat(ident, "CornerOffsetRatio", concernoutputshape.CornerOffsetRatio)
	case "IsHidden":
		res = __gong__marshallBool(ident, "IsHidden", concernoutputshape.IsHidden)
	case "IsPresentationOverridden":
		res = __gong__marshallBool(ident, "IsPresentationOverridden", concernoutputshape.IsPresentationOverridden)
	case "Stroke":
		res = __gong__marshallString(ident, "Stroke", concernoutputshape.Stroke)
	case "StrokeWidth":
		res = __gong__marshallFloat(ident, "StrokeWidth", concernoutputshape.StrokeWidth)
	case "StrokeOpacity":
		res = __gong__marshallFloat(ident, "StrokeOpacity", concernoutputshape.StrokeOpacity)
	case "StrokeDashArray":
		res = __gong__marshallString(ident, "StrokeDashArray", concernoutputshape.StrokeDashArray)
	case "FontSize":
		res = __gong__marshallString(ident, "FontSize", concernoutputshape.FontSize)
	case "FontColor":
		res = __gong__marshallString(ident, "FontColor", concernoutputshape.FontColor)

	case "Concern":
		if concernoutputshape.Concern != nil {
			res = __gong__marshallPointer(ident, "Concern", concernoutputshape.Concern.GongGetIdentifier(stage))
		} else {
			res = __gong__marshallPointer(ident, "Concern", "nil")
		}
	case "Deliverable":
		if concernoutputshape.Deliverable != nil {
			res = __gong__marshallPointer(ident, "Deliverable", concernoutputshape.Deliverable.GongGetIdentifier(stage))
		} else {
			res = __gong__marshallPointer(ident, "Deliverable", "nil")
		}
	case "ControlPointShapes":
		var sb strings.Builder
		for _, _controlpointshape := range concernoutputshape.ControlPointShapes {
			sb.WriteString(__gong__marshallSliceOfPointers(ident, "ControlPointShapes", _controlpointshape.GongGetIdentifier(stage)))
		}
		res = sb.String()
	default:
		log.Panicf("Unknown field %s for Gongstruct ConcernOutputShape", fieldName)
	}
	return
}

func (concernshape *ConcernShape) GongMarshallField(stage *Stage, fieldName string) (res string) {
	ident := concernshape.GongGetIdentifier(stage)
	_ = ident
	switch fieldName {
	case "Name":
		res = __gong__marshallString(ident, "Name", concernshape.Name)
	case "IsExpanded":
		res = __gong__marshallBool(ident, "IsExpanded", concernshape.IsExpanded)
	case "X":
		res = __gong__marshallFloat(ident, "X", concernshape.X)
	case "Y":
		res = __gong__marshallFloat(ident, "Y", concernshape.Y)
	case "Width":
		res = __gong__marshallFloat(ident, "Width", concernshape.Width)
	case "Height":
		res = __gong__marshallFloat(ident, "Height", concernshape.Height)
	case "IsHidden":
		res = __gong__marshallBool(ident, "IsHidden", concernshape.IsHidden)
	case "IsPresentationOverridden":
		res = __gong__marshallBool(ident, "IsPresentationOverridden", concernshape.IsPresentationOverridden)
	case "Color":
		res = __gong__marshallString(ident, "Color", concernshape.Color)
	case "FillOpacity":
		res = __gong__marshallFloat(ident, "FillOpacity", concernshape.FillOpacity)
	case "Stroke":
		res = __gong__marshallString(ident, "Stroke", concernshape.Stroke)
	case "StrokeWidth":
		res = __gong__marshallFloat(ident, "StrokeWidth", concernshape.StrokeWidth)
	case "StrokeOpacity":
		res = __gong__marshallFloat(ident, "StrokeOpacity", concernshape.StrokeOpacity)
	case "StrokeDashArray":
		res = __gong__marshallString(ident, "StrokeDashArray", concernshape.StrokeDashArray)
	case "RX":
		res = __gong__marshallFloat(ident, "RX", concernshape.RX)
	case "FontSize":
		res = __gong__marshallString(ident, "FontSize", concernshape.FontSize)
	case "FontWeight":
		res = __gong__marshallString(ident, "FontWeight", concernshape.FontWeight)
	case "FontColor":
		res = __gong__marshallString(ident, "FontColor", concernshape.FontColor)

	case "Concern":
		if concernshape.Concern != nil {
			res = __gong__marshallPointer(ident, "Concern", concernshape.Concern.GongGetIdentifier(stage))
		} else {
			res = __gong__marshallPointer(ident, "Concern", "nil")
		}
	default:
		log.Panicf("Unknown field %s for Gongstruct ConcernShape", fieldName)
	}
	return
}

func (controlpointshape *ControlPointShape) GongMarshallField(stage *Stage, fieldName string) (res string) {
	ident := controlpointshape.GongGetIdentifier(stage)
	_ = ident
	switch fieldName {
	case "Name":
		res = __gong__marshallString(ident, "Name", controlpointshape.Name)
	case "X_Relative":
		res = __gong__marshallFloat(ident, "X_Relative", controlpointshape.X_Relative)
	case "Y_Relative":
		res = __gong__marshallFloat(ident, "Y_Relative", controlpointshape.Y_Relative)
	case "IsStartShapeTheClosestShape":
		res = __gong__marshallBool(ident, "IsStartShapeTheClosestShape", controlpointshape.IsStartShapeTheClosestShape)

	default:
		log.Panicf("Unknown field %s for Gongstruct ControlPointShape", fieldName)
	}
	return
}

func (deliverable *Deliverable) GongMarshallField(stage *Stage, fieldName string) (res string) {
	ident := deliverable.GongGetIdentifier(stage)
	_ = ident
	switch fieldName {
	case "Name":
		res = __gong__marshallString(ident, "Name", deliverable.Name)
	case "ComputedPrefix":
		res = __gong__marshallString(ident, "ComputedPrefix", deliverable.ComputedPrefix)
	case "IsExpanded":
		res = __gong__marshallBool(ident, "IsExpanded", deliverable.IsExpanded)
	case "Description":
		res = __gong__marshallString(ident, "Description", deliverable.Description)
	case "IsProducersNodeExpanded":
		res = __gong__marshallBool(ident, "IsProducersNodeExpanded", deliverable.IsProducersNodeExpanded)
	case "IsConsumersNodeExpanded":
		res = __gong__marshallBool(ident, "IsConsumersNodeExpanded", deliverable.IsConsumersNodeExpanded)

	case "SubDeliverables":
		var sb strings.Builder
		for _, _deliverable := range deliverable.SubDeliverables {
			sb.WriteString(__gong__marshallSliceOfPointers(ident, "SubDeliverables", _deliverable.GongGetIdentifier(stage)))
		}
		res = sb.String()
	case "Views":
		var sb strings.Builder
		for _, _view := range deliverable.Views {
			sb.WriteString(__gong__marshallSliceOfPointers(ident, "Views", _view.GongGetIdentifier(stage)))
		}
		res = sb.String()
	case "Concepts":
		var sb strings.Builder
		for _, _concept := range deliverable.Concepts {
			sb.WriteString(__gong__marshallSliceOfPointers(ident, "Concepts", _concept.GongGetIdentifier(stage)))
		}
		res = sb.String()
	default:
		log.Panicf("Unknown field %s for Gongstruct Deliverable", fieldName)
	}
	return
}

func (deliverablecompositionshape *DeliverableCompositionShape) GongMarshallField(stage *Stage, fieldName string) (res string) {
	ident := deliverablecompositionshape.GongGetIdentifier(stage)
	_ = ident
	switch fieldName {
	case "Name":
		res = __gong__marshallString(ident, "Name", deliverablecompositionshape.Name)
	case "StartRatio":
		res = __gong__marshallFloat(ident, "StartRatio", deliverablecompositionshape.StartRatio)
	case "EndRatio":
		res = __gong__marshallFloat(ident, "EndRatio", deliverablecompositionshape.EndRatio)
	case "StartOrientation":
		res = __gong__marshallEnumString(ident, "StartOrientation", deliverablecompositionshape.StartOrientation.ToCodeString())
	case "EndOrientation":
		res = __gong__marshallEnumString(ident, "EndOrientation", deliverablecompositionshape.EndOrientation.ToCodeString())
	case "CornerOffsetRatio":
		res = __gong__marshallFloat(ident, "CornerOffsetRatio", deliverablecompositionshape.CornerOffsetRatio)
	case "IsHidden":
		res = __gong__marshallBool(ident, "IsHidden", deliverablecompositionshape.IsHidden)
	case "IsPresentationOverridden":
		res = __gong__marshallBool(ident, "IsPresentationOverridden", deliverablecompositionshape.IsPresentationOverridden)
	case "Stroke":
		res = __gong__marshallString(ident, "Stroke", deliverablecompositionshape.Stroke)
	case "StrokeWidth":
		res = __gong__marshallFloat(ident, "StrokeWidth", deliverablecompositionshape.StrokeWidth)
	case "StrokeOpacity":
		res = __gong__marshallFloat(ident, "StrokeOpacity", deliverablecompositionshape.StrokeOpacity)
	case "StrokeDashArray":
		res = __gong__marshallString(ident, "StrokeDashArray", deliverablecompositionshape.StrokeDashArray)
	case "FontSize":
		res = __gong__marshallString(ident, "FontSize", deliverablecompositionshape.FontSize)
	case "FontColor":
		res = __gong__marshallString(ident, "FontColor", deliverablecompositionshape.FontColor)

	case "Deliverable":
		if deliverablecompositionshape.Deliverable != nil {
			res = __gong__marshallPointer(ident, "Deliverable", deliverablecompositionshape.Deliverable.GongGetIdentifier(stage))
		} else {
			res = __gong__marshallPointer(ident, "Deliverable", "nil")
		}
	case "ControlPointShapes":
		var sb strings.Builder
		for _, _controlpointshape := range deliverablecompositionshape.ControlPointShapes {
			sb.WriteString(__gong__marshallSliceOfPointers(ident, "ControlPointShapes", _controlpointshape.GongGetIdentifier(stage)))
		}
		res = sb.String()
	default:
		log.Panicf("Unknown field %s for Gongstruct DeliverableCompositionShape", fieldName)
	}
	return
}

func (deliverableconceptshape *DeliverableConceptShape) GongMarshallField(stage *Stage, fieldName string) (res string) {
	ident := deliverableconceptshape.GongGetIdentifier(stage)
	_ = ident
	switch fieldName {
	case "Name":
		res = __gong__marshallString(ident, "Name", deliverableconceptshape.Name)
	case "StartRatio":
		res = __gong__marshallFloat(ident, "StartRatio", deliverableconceptshape.StartRatio)
	case "EndRatio":
		res = __gong__marshallFloat(ident, "EndRatio", deliverableconceptshape.EndRatio)
	case "StartOrientation":
		res = __gong__marshallEnumString(ident, "StartOrientation", deliverableconceptshape.StartOrientation.ToCodeString())
	case "EndOrientation":
		res = __gong__marshallEnumString(ident, "EndOrientation", deliverableconceptshape.EndOrientation.ToCodeString())
	case "CornerOffsetRatio":
		res = __gong__marshallFloat(ident, "CornerOffsetRatio", deliverableconceptshape.CornerOffsetRatio)
	case "IsHidden":
		res = __gong__marshallBool(ident, "IsHidden", deliverableconceptshape.IsHidden)
	case "IsPresentationOverridden":
		res = __gong__marshallBool(ident, "IsPresentationOverridden", deliverableconceptshape.IsPresentationOverridden)
	case "Stroke":
		res = __gong__marshallString(ident, "Stroke", deliverableconceptshape.Stroke)
	case "StrokeWidth":
		res = __gong__marshallFloat(ident, "StrokeWidth", deliverableconceptshape.StrokeWidth)
	case "StrokeOpacity":
		res = __gong__marshallFloat(ident, "StrokeOpacity", deliverableconceptshape.StrokeOpacity)
	case "StrokeDashArray":
		res = __gong__marshallString(ident, "StrokeDashArray", deliverableconceptshape.StrokeDashArray)
	case "FontSize":
		res = __gong__marshallString(ident, "FontSize", deliverableconceptshape.FontSize)
	case "FontColor":
		res = __gong__marshallString(ident, "FontColor", deliverableconceptshape.FontColor)

	case "Deliverable":
		if deliverableconceptshape.Deliverable != nil {
			res = __gong__marshallPointer(ident, "Deliverable", deliverableconceptshape.Deliverable.GongGetIdentifier(stage))
		} else {
			res = __gong__marshallPointer(ident, "Deliverable", "nil")
		}
	case "Concept":
		if deliverableconceptshape.Concept != nil {
			res = __gong__marshallPointer(ident, "Concept", deliverableconceptshape.Concept.GongGetIdentifier(stage))
		} else {
			res = __gong__marshallPointer(ident, "Concept", "nil")
		}
	case "ControlPointShapes":
		var sb strings.Builder
		for _, _controlpointshape := range deliverableconceptshape.ControlPointShapes {
			sb.WriteString(__gong__marshallSliceOfPointers(ident, "ControlPointShapes", _controlpointshape.GongGetIdentifier(stage)))
		}
		res = sb.String()
	default:
		log.Panicf("Unknown field %s for Gongstruct DeliverableConceptShape", fieldName)
	}
	return
}

func (deliverableshape *DeliverableShape) GongMarshallField(stage *Stage, fieldName string) (res string) {
	ident := deliverableshape.GongGetIdentifier(stage)
	_ = ident
	switch fieldName {
	case "Name":
		res = __gong__marshallString(ident, "Name", deliverableshape.Name)
	case "IsExpanded":
		res = __gong__marshallBool(ident, "IsExpanded", deliverableshape.IsExpanded)
	case "X":
		res = __gong__marshallFloat(ident, "X", deliverableshape.X)
	case "Y":
		res = __gong__marshallFloat(ident, "Y", deliverableshape.Y)
	case "Width":
		res = __gong__marshallFloat(ident, "Width", deliverableshape.Width)
	case "Height":
		res = __gong__marshallFloat(ident, "Height", deliverableshape.Height)
	case "IsHidden":
		res = __gong__marshallBool(ident, "IsHidden", deliverableshape.IsHidden)
	case "IsPresentationOverridden":
		res = __gong__marshallBool(ident, "IsPresentationOverridden", deliverableshape.IsPresentationOverridden)
	case "Color":
		res = __gong__marshallString(ident, "Color", deliverableshape.Color)
	case "FillOpacity":
		res = __gong__marshallFloat(ident, "FillOpacity", deliverableshape.FillOpacity)
	case "Stroke":
		res = __gong__marshallString(ident, "Stroke", deliverableshape.Stroke)
	case "StrokeWidth":
		res = __gong__marshallFloat(ident, "StrokeWidth", deliverableshape.StrokeWidth)
	case "StrokeOpacity":
		res = __gong__marshallFloat(ident, "StrokeOpacity", deliverableshape.StrokeOpacity)
	case "StrokeDashArray":
		res = __gong__marshallString(ident, "StrokeDashArray", deliverableshape.StrokeDashArray)
	case "RX":
		res = __gong__marshallFloat(ident, "RX", deliverableshape.RX)
	case "FontSize":
		res = __gong__marshallString(ident, "FontSize", deliverableshape.FontSize)
	case "FontWeight":
		res = __gong__marshallString(ident, "FontWeight", deliverableshape.FontWeight)
	case "FontColor":
		res = __gong__marshallString(ident, "FontColor", deliverableshape.FontColor)

	case "Deliverable":
		if deliverableshape.Deliverable != nil {
			res = __gong__marshallPointer(ident, "Deliverable", deliverableshape.Deliverable.GongGetIdentifier(stage))
		} else {
			res = __gong__marshallPointer(ident, "Deliverable", "nil")
		}
	default:
		log.Panicf("Unknown field %s for Gongstruct DeliverableShape", fieldName)
	}
	return
}

func (deliverableviewshape *DeliverableViewShape) GongMarshallField(stage *Stage, fieldName string) (res string) {
	ident := deliverableviewshape.GongGetIdentifier(stage)
	_ = ident
	switch fieldName {
	case "Name":
		res = __gong__marshallString(ident, "Name", deliverableviewshape.Name)
	case "StartRatio":
		res = __gong__marshallFloat(ident, "StartRatio", deliverableviewshape.StartRatio)
	case "EndRatio":
		res = __gong__marshallFloat(ident, "EndRatio", deliverableviewshape.EndRatio)
	case "StartOrientation":
		res = __gong__marshallEnumString(ident, "StartOrientation", deliverableviewshape.StartOrientation.ToCodeString())
	case "EndOrientation":
		res = __gong__marshallEnumString(ident, "EndOrientation", deliverableviewshape.EndOrientation.ToCodeString())
	case "CornerOffsetRatio":
		res = __gong__marshallFloat(ident, "CornerOffsetRatio", deliverableviewshape.CornerOffsetRatio)
	case "IsHidden":
		res = __gong__marshallBool(ident, "IsHidden", deliverableviewshape.IsHidden)
	case "IsPresentationOverridden":
		res = __gong__marshallBool(ident, "IsPresentationOverridden", deliverableviewshape.IsPresentationOverridden)
	case "Stroke":
		res = __gong__marshallString(ident, "Stroke", deliverableviewshape.Stroke)
	case "StrokeWidth":
		res = __gong__marshallFloat(ident, "StrokeWidth", deliverableviewshape.StrokeWidth)
	case "StrokeOpacity":
		res = __gong__marshallFloat(ident, "StrokeOpacity", deliverableviewshape.StrokeOpacity)
	case "StrokeDashArray":
		res = __gong__marshallString(ident, "StrokeDashArray", deliverableviewshape.StrokeDashArray)
	case "FontSize":
		res = __gong__marshallString(ident, "FontSize", deliverableviewshape.FontSize)
	case "FontColor":
		res = __gong__marshallString(ident, "FontColor", deliverableviewshape.FontColor)

	case "Deliverable":
		if deliverableviewshape.Deliverable != nil {
			res = __gong__marshallPointer(ident, "Deliverable", deliverableviewshape.Deliverable.GongGetIdentifier(stage))
		} else {
			res = __gong__marshallPointer(ident, "Deliverable", "nil")
		}
	case "View":
		if deliverableviewshape.View != nil {
			res = __gong__marshallPointer(ident, "View", deliverableviewshape.View.GongGetIdentifier(stage))
		} else {
			res = __gong__marshallPointer(ident, "View", "nil")
		}
	case "ControlPointShapes":
		var sb strings.Builder
		for _, _controlpointshape := range deliverableviewshape.ControlPointShapes {
			sb.WriteString(__gong__marshallSliceOfPointers(ident, "ControlPointShapes", _controlpointshape.GongGetIdentifier(stage)))
		}
		res = sb.String()
	default:
		log.Panicf("Unknown field %s for Gongstruct DeliverableViewShape", fieldName)
	}
	return
}

func (diagram *Diagram) GongMarshallField(stage *Stage, fieldName string) (res string) {
	ident := diagram.GongGetIdentifier(stage)
	_ = ident
	switch fieldName {
	case "Name":
		res = __gong__marshallString(ident, "Name", diagram.Name)
	case "ComputedPrefix":
		res = __gong__marshallString(ident, "ComputedPrefix", diagram.ComputedPrefix)
	case "IsExpanded":
		res = __gong__marshallBool(ident, "IsExpanded", diagram.IsExpanded)
	case "IsChecked":
		res = __gong__marshallBool(ident, "IsChecked", diagram.IsChecked)
	case "IsEditable_":
		res = __gong__marshallBool(ident, "IsEditable_", diagram.IsEditable_)
	case "ShowPrefix":
		res = __gong__marshallBool(ident, "ShowPrefix", diagram.ShowPrefix)
	case "DefaultBoxWidth":
		res = __gong__marshallFloat(ident, "DefaultBoxWidth", diagram.DefaultBoxWidth)
	case "DefaultBoxHeigth":
		res = __gong__marshallFloat(ident, "DefaultBoxHeigth", diagram.DefaultBoxHeigth)
	case "Width":
		res = __gong__marshallFloat(ident, "Width", diagram.Width)
	case "Height":
		res = __gong__marshallFloat(ident, "Height", diagram.Height)
	case "IsRequirementsNodeExpanded":
		res = __gong__marshallBool(ident, "IsRequirementsNodeExpanded", diagram.IsRequirementsNodeExpanded)
	case "IsViewsNodeExpanded":
		res = __gong__marshallBool(ident, "IsViewsNodeExpanded", diagram.IsViewsNodeExpanded)
	case "IsConceptsNodeExpanded":
		res = __gong__marshallBool(ident, "IsConceptsNodeExpanded", diagram.IsConceptsNodeExpanded)
	case "IsPBSNodeExpanded":
		res = __gong__marshallBool(ident, "IsPBSNodeExpanded", diagram.IsPBSNodeExpanded)
	case "IsConcernsNodeExpanded":
		res = __gong__marshallBool(ident, "IsConcernsNodeExpanded", diagram.IsConcernsNodeExpanded)
	case "IsNotesNodeExpanded":
		res = __gong__marshallBool(ident, "IsNotesNodeExpanded", diagram.IsNotesNodeExpanded)
	case "IsStakeholdersNodeExpanded":
		res = __gong__marshallBool(ident, "IsStakeholdersNodeExpanded", diagram.IsStakeholdersNodeExpanded)
	case "IsDiagramsNodeExpanded":
		res = __gong__marshallBool(ident, "IsDiagramsNodeExpanded", diagram.IsDiagramsNodeExpanded)

	case "ConcernsWhoseRequirementsNodeIsExpanded":
		var sb strings.Builder
		for _, _concern := range diagram.ConcernsWhoseRequirementsNodeIsExpanded {
			sb.WriteString(__gong__marshallSliceOfPointers(ident, "ConcernsWhoseRequirementsNodeIsExpanded", _concern.GongGetIdentifier(stage)))
		}
		res = sb.String()
	case "Deliverable_Shapes":
		var sb strings.Builder
		for _, _deliverableshape := range diagram.Deliverable_Shapes {
			sb.WriteString(__gong__marshallSliceOfPointers(ident, "Deliverable_Shapes", _deliverableshape.GongGetIdentifier(stage)))
		}
		res = sb.String()
	case "DeliverablesWhoseNodeIsExpanded":
		var sb strings.Builder
		for _, _deliverable := range diagram.DeliverablesWhoseNodeIsExpanded {
			sb.WriteString(__gong__marshallSliceOfPointers(ident, "DeliverablesWhoseNodeIsExpanded", _deliverable.GongGetIdentifier(stage)))
		}
		res = sb.String()
	case "DeliverablesWhoseViewsNodeIsExpanded":
		var sb strings.Builder
		for _, _deliverable := range diagram.DeliverablesWhoseViewsNodeIsExpanded {
			sb.WriteString(__gong__marshallSliceOfPointers(ident, "DeliverablesWhoseViewsNodeIsExpanded", _deliverable.GongGetIdentifier(stage)))
		}
		res = sb.String()
	case "DeliverablesWhoseConceptsNodeIsExpanded":
		var sb strings.Builder
		for _, _deliverable := range diagram.DeliverablesWhoseConceptsNodeIsExpanded {
			sb.WriteString(__gong__marshallSliceOfPointers(ident, "DeliverablesWhoseConceptsNodeIsExpanded", _deliverable.GongGetIdentifier(stage)))
		}
		res = sb.String()
	case "DeliverableComposition_Shapes":
		var sb strings.Builder
		for _, _deliverablecompositionshape := range diagram.DeliverableComposition_Shapes {
			sb.WriteString(__gong__marshallSliceOfPointers(ident, "DeliverableComposition_Shapes", _deliverablecompositionshape.GongGetIdentifier(stage)))
		}
		res = sb.String()
	case "Concern_Shapes":
		var sb strings.Builder
		for _, _concernshape := range diagram.Concern_Shapes {
			sb.WriteString(__gong__marshallSliceOfPointers(ident, "Concern_Shapes", _concernshape.GongGetIdentifier(stage)))
		}
		res = sb.String()
	case "ConcernsWhoseNodeIsExpanded":
		var sb strings.Builder
		for _, _concern := range diagram.ConcernsWhoseNodeIsExpanded {
			sb.WriteString(__gong__marshallSliceOfPointers(ident, "ConcernsWhoseNodeIsExpanded", _concern.GongGetIdentifier(stage)))
		}
		res = sb.String()
	case "ConcernsWhoseInputNodeIsExpanded":
		var sb strings.Builder
		for _, _concern := range diagram.ConcernsWhoseInputNodeIsExpanded {
			sb.WriteString(__gong__marshallSliceOfPointers(ident, "ConcernsWhoseInputNodeIsExpanded", _concern.GongGetIdentifier(stage)))
		}
		res = sb.String()
	case "ConcernsWhoseStakeholderNodeIsExpanded":
		var sb strings.Builder
		for _, _concern := range diagram.ConcernsWhoseStakeholderNodeIsExpanded {
			sb.WriteString(__gong__marshallSliceOfPointers(ident, "ConcernsWhoseStakeholderNodeIsExpanded", _concern.GongGetIdentifier(stage)))
		}
		res = sb.String()
	case "ConcernssWhoseOutputNodeIsExpanded":
		var sb strings.Builder
		for _, _concern := range diagram.ConcernssWhoseOutputNodeIsExpanded {
			sb.WriteString(__gong__marshallSliceOfPointers(ident, "ConcernssWhoseOutputNodeIsExpanded", _concern.GongGetIdentifier(stage)))
		}
		res = sb.String()
	case "ConcernComposition_Shapes":
		var sb strings.Builder
		for _, _concerncompositionshape := range diagram.ConcernComposition_Shapes {
			sb.WriteString(__gong__marshallSliceOfPointers(ident, "ConcernComposition_Shapes", _concerncompositionshape.GongGetIdentifier(stage)))
		}
		res = sb.String()
	case "ConcernInputShapes":
		var sb strings.Builder
		for _, _concerninputshape := range diagram.ConcernInputShapes {
			sb.WriteString(__gong__marshallSliceOfPointers(ident, "ConcernInputShapes", _concerninputshape.GongGetIdentifier(stage)))
		}
		res = sb.String()
	case "ConcernOutputShapes":
		var sb strings.Builder
		for _, _concernoutputshape := range diagram.ConcernOutputShapes {
			sb.WriteString(__gong__marshallSliceOfPointers(ident, "ConcernOutputShapes", _concernoutputshape.GongGetIdentifier(stage)))
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
	case "NoteDeliverableShapes":
		var sb strings.Builder
		for _, _notedeliverableshape := range diagram.NoteDeliverableShapes {
			sb.WriteString(__gong__marshallSliceOfPointers(ident, "NoteDeliverableShapes", _notedeliverableshape.GongGetIdentifier(stage)))
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
		for _, _notestakeholdershape := range diagram.NoteResourceShapes {
			sb.WriteString(__gong__marshallSliceOfPointers(ident, "NoteResourceShapes", _notestakeholdershape.GongGetIdentifier(stage)))
		}
		res = sb.String()
	case "Stakeholder_Shapes":
		var sb strings.Builder
		for _, _stakeholdershape := range diagram.Stakeholder_Shapes {
			sb.WriteString(__gong__marshallSliceOfPointers(ident, "Stakeholder_Shapes", _stakeholdershape.GongGetIdentifier(stage)))
		}
		res = sb.String()
	case "ResourcesWhoseNodeIsExpanded":
		var sb strings.Builder
		for _, _stakeholder := range diagram.ResourcesWhoseNodeIsExpanded {
			sb.WriteString(__gong__marshallSliceOfPointers(ident, "ResourcesWhoseNodeIsExpanded", _stakeholder.GongGetIdentifier(stage)))
		}
		res = sb.String()
	case "ResourceComposition_Shapes":
		var sb strings.Builder
		for _, _stakeholdercompositionshape := range diagram.ResourceComposition_Shapes {
			sb.WriteString(__gong__marshallSliceOfPointers(ident, "ResourceComposition_Shapes", _stakeholdercompositionshape.GongGetIdentifier(stage)))
		}
		res = sb.String()
	case "StakeholderConcernShapes":
		var sb strings.Builder
		for _, _stakeholderconcernshape := range diagram.StakeholderConcernShapes {
			sb.WriteString(__gong__marshallSliceOfPointers(ident, "StakeholderConcernShapes", _stakeholderconcernshape.GongGetIdentifier(stage)))
		}
		res = sb.String()
	case "Requirement_Shapes":
		var sb strings.Builder
		for _, _requirementshape := range diagram.Requirement_Shapes {
			sb.WriteString(__gong__marshallSliceOfPointers(ident, "Requirement_Shapes", _requirementshape.GongGetIdentifier(stage)))
		}
		res = sb.String()
	case "RequirementsWhoseNodeIsExpanded":
		var sb strings.Builder
		for _, _requirement := range diagram.RequirementsWhoseNodeIsExpanded {
			sb.WriteString(__gong__marshallSliceOfPointers(ident, "RequirementsWhoseNodeIsExpanded", _requirement.GongGetIdentifier(stage)))
		}
		res = sb.String()
	case "RequirementsWhoseConcernsNodeIsExpanded":
		var sb strings.Builder
		for _, _requirement := range diagram.RequirementsWhoseConcernsNodeIsExpanded {
			sb.WriteString(__gong__marshallSliceOfPointers(ident, "RequirementsWhoseConcernsNodeIsExpanded", _requirement.GongGetIdentifier(stage)))
		}
		res = sb.String()
	case "RequirementsWhoseViewsNodeIsExpanded":
		var sb strings.Builder
		for _, _requirement := range diagram.RequirementsWhoseViewsNodeIsExpanded {
			sb.WriteString(__gong__marshallSliceOfPointers(ident, "RequirementsWhoseViewsNodeIsExpanded", _requirement.GongGetIdentifier(stage)))
		}
		res = sb.String()
	case "RequirementsWhoseConceptsNodeIsExpanded":
		var sb strings.Builder
		for _, _requirement := range diagram.RequirementsWhoseConceptsNodeIsExpanded {
			sb.WriteString(__gong__marshallSliceOfPointers(ident, "RequirementsWhoseConceptsNodeIsExpanded", _requirement.GongGetIdentifier(stage)))
		}
		res = sb.String()
	case "RequirementConcernShapes":
		var sb strings.Builder
		for _, _requirementconcernshape := range diagram.RequirementConcernShapes {
			sb.WriteString(__gong__marshallSliceOfPointers(ident, "RequirementConcernShapes", _requirementconcernshape.GongGetIdentifier(stage)))
		}
		res = sb.String()
	case "RequirementViewShapes":
		var sb strings.Builder
		for _, _requirementviewshape := range diagram.RequirementViewShapes {
			sb.WriteString(__gong__marshallSliceOfPointers(ident, "RequirementViewShapes", _requirementviewshape.GongGetIdentifier(stage)))
		}
		res = sb.String()
	case "RequirementConceptShapes":
		var sb strings.Builder
		for _, _requirementconceptshape := range diagram.RequirementConceptShapes {
			sb.WriteString(__gong__marshallSliceOfPointers(ident, "RequirementConceptShapes", _requirementconceptshape.GongGetIdentifier(stage)))
		}
		res = sb.String()
	case "View_Shapes":
		var sb strings.Builder
		for _, _viewshape := range diagram.View_Shapes {
			sb.WriteString(__gong__marshallSliceOfPointers(ident, "View_Shapes", _viewshape.GongGetIdentifier(stage)))
		}
		res = sb.String()
	case "ViewsWhoseNodeIsExpanded":
		var sb strings.Builder
		for _, _view := range diagram.ViewsWhoseNodeIsExpanded {
			sb.WriteString(__gong__marshallSliceOfPointers(ident, "ViewsWhoseNodeIsExpanded", _view.GongGetIdentifier(stage)))
		}
		res = sb.String()
	case "ViewsWhoseDeliverablesNodeIsExpanded":
		var sb strings.Builder
		for _, _view := range diagram.ViewsWhoseDeliverablesNodeIsExpanded {
			sb.WriteString(__gong__marshallSliceOfPointers(ident, "ViewsWhoseDeliverablesNodeIsExpanded", _view.GongGetIdentifier(stage)))
		}
		res = sb.String()
	case "ViewsWhoseConceptsNodeIsExpanded":
		var sb strings.Builder
		for _, _view := range diagram.ViewsWhoseConceptsNodeIsExpanded {
			sb.WriteString(__gong__marshallSliceOfPointers(ident, "ViewsWhoseConceptsNodeIsExpanded", _view.GongGetIdentifier(stage)))
		}
		res = sb.String()
	case "ViewsWhoseRequirementsNodeIsExpanded":
		var sb strings.Builder
		for _, _view := range diagram.ViewsWhoseRequirementsNodeIsExpanded {
			sb.WriteString(__gong__marshallSliceOfPointers(ident, "ViewsWhoseRequirementsNodeIsExpanded", _view.GongGetIdentifier(stage)))
		}
		res = sb.String()
	case "DeliverableViewShapes":
		var sb strings.Builder
		for _, _deliverableviewshape := range diagram.DeliverableViewShapes {
			sb.WriteString(__gong__marshallSliceOfPointers(ident, "DeliverableViewShapes", _deliverableviewshape.GongGetIdentifier(stage)))
		}
		res = sb.String()
	case "ViewConceptShapes":
		var sb strings.Builder
		for _, _viewconceptshape := range diagram.ViewConceptShapes {
			sb.WriteString(__gong__marshallSliceOfPointers(ident, "ViewConceptShapes", _viewconceptshape.GongGetIdentifier(stage)))
		}
		res = sb.String()
	case "Concept_Shapes":
		var sb strings.Builder
		for _, _conceptshape := range diagram.Concept_Shapes {
			sb.WriteString(__gong__marshallSliceOfPointers(ident, "Concept_Shapes", _conceptshape.GongGetIdentifier(stage)))
		}
		res = sb.String()
	case "ConceptsWhoseNodeIsExpanded":
		var sb strings.Builder
		for _, _concept := range diagram.ConceptsWhoseNodeIsExpanded {
			sb.WriteString(__gong__marshallSliceOfPointers(ident, "ConceptsWhoseNodeIsExpanded", _concept.GongGetIdentifier(stage)))
		}
		res = sb.String()
	case "ConceptsWhoseViewsNodeIsExpanded":
		var sb strings.Builder
		for _, _concept := range diagram.ConceptsWhoseViewsNodeIsExpanded {
			sb.WriteString(__gong__marshallSliceOfPointers(ident, "ConceptsWhoseViewsNodeIsExpanded", _concept.GongGetIdentifier(stage)))
		}
		res = sb.String()
	case "ConceptsWhoseDeliverablesNodeIsExpanded":
		var sb strings.Builder
		for _, _concept := range diagram.ConceptsWhoseDeliverablesNodeIsExpanded {
			sb.WriteString(__gong__marshallSliceOfPointers(ident, "ConceptsWhoseDeliverablesNodeIsExpanded", _concept.GongGetIdentifier(stage)))
		}
		res = sb.String()
	case "ConceptsWhoseRequirementsNodeIsExpanded":
		var sb strings.Builder
		for _, _concept := range diagram.ConceptsWhoseRequirementsNodeIsExpanded {
			sb.WriteString(__gong__marshallSliceOfPointers(ident, "ConceptsWhoseRequirementsNodeIsExpanded", _concept.GongGetIdentifier(stage)))
		}
		res = sb.String()
	case "DeliverableConceptShapes":
		var sb strings.Builder
		for _, _deliverableconceptshape := range diagram.DeliverableConceptShapes {
			sb.WriteString(__gong__marshallSliceOfPointers(ident, "DeliverableConceptShapes", _deliverableconceptshape.GongGetIdentifier(stage)))
		}
		res = sb.String()
	case "Diagram_Shapes":
		var sb strings.Builder
		for _, _diagramshape := range diagram.Diagram_Shapes {
			sb.WriteString(__gong__marshallSliceOfPointers(ident, "Diagram_Shapes", _diagramshape.GongGetIdentifier(stage)))
		}
		res = sb.String()
	case "DiagramsWhoseNodeIsExpanded":
		var sb strings.Builder
		for _, _diagram := range diagram.DiagramsWhoseNodeIsExpanded {
			sb.WriteString(__gong__marshallSliceOfPointers(ident, "DiagramsWhoseNodeIsExpanded", _diagram.GongGetIdentifier(stage)))
		}
		res = sb.String()
	default:
		log.Panicf("Unknown field %s for Gongstruct Diagram", fieldName)
	}
	return
}

func (diagramshape *DiagramShape) GongMarshallField(stage *Stage, fieldName string) (res string) {
	ident := diagramshape.GongGetIdentifier(stage)
	_ = ident
	switch fieldName {
	case "Name":
		res = __gong__marshallString(ident, "Name", diagramshape.Name)
	case "IsExpanded":
		res = __gong__marshallBool(ident, "IsExpanded", diagramshape.IsExpanded)
	case "X":
		res = __gong__marshallFloat(ident, "X", diagramshape.X)
	case "Y":
		res = __gong__marshallFloat(ident, "Y", diagramshape.Y)
	case "Width":
		res = __gong__marshallFloat(ident, "Width", diagramshape.Width)
	case "Height":
		res = __gong__marshallFloat(ident, "Height", diagramshape.Height)
	case "IsHidden":
		res = __gong__marshallBool(ident, "IsHidden", diagramshape.IsHidden)
	case "IsPresentationOverridden":
		res = __gong__marshallBool(ident, "IsPresentationOverridden", diagramshape.IsPresentationOverridden)
	case "Color":
		res = __gong__marshallString(ident, "Color", diagramshape.Color)
	case "FillOpacity":
		res = __gong__marshallFloat(ident, "FillOpacity", diagramshape.FillOpacity)
	case "Stroke":
		res = __gong__marshallString(ident, "Stroke", diagramshape.Stroke)
	case "StrokeWidth":
		res = __gong__marshallFloat(ident, "StrokeWidth", diagramshape.StrokeWidth)
	case "StrokeOpacity":
		res = __gong__marshallFloat(ident, "StrokeOpacity", diagramshape.StrokeOpacity)
	case "StrokeDashArray":
		res = __gong__marshallString(ident, "StrokeDashArray", diagramshape.StrokeDashArray)
	case "RX":
		res = __gong__marshallFloat(ident, "RX", diagramshape.RX)
	case "FontSize":
		res = __gong__marshallString(ident, "FontSize", diagramshape.FontSize)
	case "FontWeight":
		res = __gong__marshallString(ident, "FontWeight", diagramshape.FontWeight)
	case "FontColor":
		res = __gong__marshallString(ident, "FontColor", diagramshape.FontColor)

	case "Diagram":
		if diagramshape.Diagram != nil {
			res = __gong__marshallPointer(ident, "Diagram", diagramshape.Diagram.GongGetIdentifier(stage))
		} else {
			res = __gong__marshallPointer(ident, "Diagram", "nil")
		}
	default:
		log.Panicf("Unknown field %s for Gongstruct DiagramShape", fieldName)
	}
	return
}

func (library *Library) GongMarshallField(stage *Stage, fieldName string) (res string) {
	ident := library.GongGetIdentifier(stage)
	_ = ident
	switch fieldName {
	case "Name":
		res = __gong__marshallString(ident, "Name", library.Name)
	case "IsRootLibrary":
		res = __gong__marshallBool(ident, "IsRootLibrary", library.IsRootLibrary)
	case "ComputedPrefix":
		res = __gong__marshallString(ident, "ComputedPrefix", library.ComputedPrefix)
	case "IsExpanded":
		res = __gong__marshallBool(ident, "IsExpanded", library.IsExpanded)
	case "NbPixPerCharacter":
		res = __gong__marshallFloat(ident, "NbPixPerCharacter", library.NbPixPerCharacter)

	case "RootDeliverables":
		var sb strings.Builder
		for _, _deliverable := range library.RootDeliverables {
			sb.WriteString(__gong__marshallSliceOfPointers(ident, "RootDeliverables", _deliverable.GongGetIdentifier(stage)))
		}
		res = sb.String()
	case "RootConcerns":
		var sb strings.Builder
		for _, _concern := range library.RootConcerns {
			sb.WriteString(__gong__marshallSliceOfPointers(ident, "RootConcerns", _concern.GongGetIdentifier(stage)))
		}
		res = sb.String()
	case "RootStakeholders":
		var sb strings.Builder
		for _, _stakeholder := range library.RootStakeholders {
			sb.WriteString(__gong__marshallSliceOfPointers(ident, "RootStakeholders", _stakeholder.GongGetIdentifier(stage)))
		}
		res = sb.String()
	case "RootRequirements":
		var sb strings.Builder
		for _, _requirement := range library.RootRequirements {
			sb.WriteString(__gong__marshallSliceOfPointers(ident, "RootRequirements", _requirement.GongGetIdentifier(stage)))
		}
		res = sb.String()
	case "RootViews":
		var sb strings.Builder
		for _, _view := range library.RootViews {
			sb.WriteString(__gong__marshallSliceOfPointers(ident, "RootViews", _view.GongGetIdentifier(stage)))
		}
		res = sb.String()
	case "RootConcepts":
		var sb strings.Builder
		for _, _concept := range library.RootConcepts {
			sb.WriteString(__gong__marshallSliceOfPointers(ident, "RootConcepts", _concept.GongGetIdentifier(stage)))
		}
		res = sb.String()
	case "AnalysisNeeds":
		var sb strings.Builder
		for _, _analysisneed := range library.AnalysisNeeds {
			sb.WriteString(__gong__marshallSliceOfPointers(ident, "AnalysisNeeds", _analysisneed.GongGetIdentifier(stage)))
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
	case "SubLibraries":
		var sb strings.Builder
		for _, _library := range library.SubLibraries {
			sb.WriteString(__gong__marshallSliceOfPointers(ident, "SubLibraries", _library.GongGetIdentifier(stage)))
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

	case "Deliverables":
		var sb strings.Builder
		for _, _deliverable := range note.Deliverables {
			sb.WriteString(__gong__marshallSliceOfPointers(ident, "Deliverables", _deliverable.GongGetIdentifier(stage)))
		}
		res = sb.String()
	case "Tasks":
		var sb strings.Builder
		for _, _concern := range note.Tasks {
			sb.WriteString(__gong__marshallSliceOfPointers(ident, "Tasks", _concern.GongGetIdentifier(stage)))
		}
		res = sb.String()
	case "Resources":
		var sb strings.Builder
		for _, _stakeholder := range note.Resources {
			sb.WriteString(__gong__marshallSliceOfPointers(ident, "Resources", _stakeholder.GongGetIdentifier(stage)))
		}
		res = sb.String()
	default:
		log.Panicf("Unknown field %s for Gongstruct Note", fieldName)
	}
	return
}

func (notedeliverableshape *NoteDeliverableShape) GongMarshallField(stage *Stage, fieldName string) (res string) {
	ident := notedeliverableshape.GongGetIdentifier(stage)
	_ = ident
	switch fieldName {
	case "Name":
		res = __gong__marshallString(ident, "Name", notedeliverableshape.Name)
	case "StartRatio":
		res = __gong__marshallFloat(ident, "StartRatio", notedeliverableshape.StartRatio)
	case "EndRatio":
		res = __gong__marshallFloat(ident, "EndRatio", notedeliverableshape.EndRatio)
	case "StartOrientation":
		res = __gong__marshallEnumString(ident, "StartOrientation", notedeliverableshape.StartOrientation.ToCodeString())
	case "EndOrientation":
		res = __gong__marshallEnumString(ident, "EndOrientation", notedeliverableshape.EndOrientation.ToCodeString())
	case "CornerOffsetRatio":
		res = __gong__marshallFloat(ident, "CornerOffsetRatio", notedeliverableshape.CornerOffsetRatio)
	case "IsHidden":
		res = __gong__marshallBool(ident, "IsHidden", notedeliverableshape.IsHidden)
	case "IsPresentationOverridden":
		res = __gong__marshallBool(ident, "IsPresentationOverridden", notedeliverableshape.IsPresentationOverridden)
	case "Stroke":
		res = __gong__marshallString(ident, "Stroke", notedeliverableshape.Stroke)
	case "StrokeWidth":
		res = __gong__marshallFloat(ident, "StrokeWidth", notedeliverableshape.StrokeWidth)
	case "StrokeOpacity":
		res = __gong__marshallFloat(ident, "StrokeOpacity", notedeliverableshape.StrokeOpacity)
	case "StrokeDashArray":
		res = __gong__marshallString(ident, "StrokeDashArray", notedeliverableshape.StrokeDashArray)
	case "FontSize":
		res = __gong__marshallString(ident, "FontSize", notedeliverableshape.FontSize)
	case "FontColor":
		res = __gong__marshallString(ident, "FontColor", notedeliverableshape.FontColor)

	case "Note":
		if notedeliverableshape.Note != nil {
			res = __gong__marshallPointer(ident, "Note", notedeliverableshape.Note.GongGetIdentifier(stage))
		} else {
			res = __gong__marshallPointer(ident, "Note", "nil")
		}
	case "Deliverable":
		if notedeliverableshape.Deliverable != nil {
			res = __gong__marshallPointer(ident, "Deliverable", notedeliverableshape.Deliverable.GongGetIdentifier(stage))
		} else {
			res = __gong__marshallPointer(ident, "Deliverable", "nil")
		}
	case "ControlPointShapes":
		var sb strings.Builder
		for _, _controlpointshape := range notedeliverableshape.ControlPointShapes {
			sb.WriteString(__gong__marshallSliceOfPointers(ident, "ControlPointShapes", _controlpointshape.GongGetIdentifier(stage)))
		}
		res = sb.String()
	default:
		log.Panicf("Unknown field %s for Gongstruct NoteDeliverableShape", fieldName)
	}
	return
}

func (noteshape *NoteShape) GongMarshallField(stage *Stage, fieldName string) (res string) {
	ident := noteshape.GongGetIdentifier(stage)
	_ = ident
	switch fieldName {
	case "Name":
		res = __gong__marshallString(ident, "Name", noteshape.Name)
	case "IsExpanded":
		res = __gong__marshallBool(ident, "IsExpanded", noteshape.IsExpanded)
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

func (notestakeholdershape *NoteStakeholderShape) GongMarshallField(stage *Stage, fieldName string) (res string) {
	ident := notestakeholdershape.GongGetIdentifier(stage)
	_ = ident
	switch fieldName {
	case "Name":
		res = __gong__marshallString(ident, "Name", notestakeholdershape.Name)
	case "StartRatio":
		res = __gong__marshallFloat(ident, "StartRatio", notestakeholdershape.StartRatio)
	case "EndRatio":
		res = __gong__marshallFloat(ident, "EndRatio", notestakeholdershape.EndRatio)
	case "StartOrientation":
		res = __gong__marshallEnumString(ident, "StartOrientation", notestakeholdershape.StartOrientation.ToCodeString())
	case "EndOrientation":
		res = __gong__marshallEnumString(ident, "EndOrientation", notestakeholdershape.EndOrientation.ToCodeString())
	case "CornerOffsetRatio":
		res = __gong__marshallFloat(ident, "CornerOffsetRatio", notestakeholdershape.CornerOffsetRatio)
	case "IsHidden":
		res = __gong__marshallBool(ident, "IsHidden", notestakeholdershape.IsHidden)
	case "IsPresentationOverridden":
		res = __gong__marshallBool(ident, "IsPresentationOverridden", notestakeholdershape.IsPresentationOverridden)
	case "Stroke":
		res = __gong__marshallString(ident, "Stroke", notestakeholdershape.Stroke)
	case "StrokeWidth":
		res = __gong__marshallFloat(ident, "StrokeWidth", notestakeholdershape.StrokeWidth)
	case "StrokeOpacity":
		res = __gong__marshallFloat(ident, "StrokeOpacity", notestakeholdershape.StrokeOpacity)
	case "StrokeDashArray":
		res = __gong__marshallString(ident, "StrokeDashArray", notestakeholdershape.StrokeDashArray)
	case "FontSize":
		res = __gong__marshallString(ident, "FontSize", notestakeholdershape.FontSize)
	case "FontColor":
		res = __gong__marshallString(ident, "FontColor", notestakeholdershape.FontColor)

	case "Note":
		if notestakeholdershape.Note != nil {
			res = __gong__marshallPointer(ident, "Note", notestakeholdershape.Note.GongGetIdentifier(stage))
		} else {
			res = __gong__marshallPointer(ident, "Note", "nil")
		}
	case "Stakeholder":
		if notestakeholdershape.Stakeholder != nil {
			res = __gong__marshallPointer(ident, "Stakeholder", notestakeholdershape.Stakeholder.GongGetIdentifier(stage))
		} else {
			res = __gong__marshallPointer(ident, "Stakeholder", "nil")
		}
	case "ControlPointShapes":
		var sb strings.Builder
		for _, _controlpointshape := range notestakeholdershape.ControlPointShapes {
			sb.WriteString(__gong__marshallSliceOfPointers(ident, "ControlPointShapes", _controlpointshape.GongGetIdentifier(stage)))
		}
		res = sb.String()
	default:
		log.Panicf("Unknown field %s for Gongstruct NoteStakeholderShape", fieldName)
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
	case "ControlPointShapes":
		var sb strings.Builder
		for _, _controlpointshape := range notetaskshape.ControlPointShapes {
			sb.WriteString(__gong__marshallSliceOfPointers(ident, "ControlPointShapes", _controlpointshape.GongGetIdentifier(stage)))
		}
		res = sb.String()
	default:
		log.Panicf("Unknown field %s for Gongstruct NoteTaskShape", fieldName)
	}
	return
}

func (requirement *Requirement) GongMarshallField(stage *Stage, fieldName string) (res string) {
	ident := requirement.GongGetIdentifier(stage)
	_ = ident
	switch fieldName {
	case "Name":
		res = __gong__marshallString(ident, "Name", requirement.Name)
	case "ComputedPrefix":
		res = __gong__marshallString(ident, "ComputedPrefix", requirement.ComputedPrefix)
	case "IsExpanded":
		res = __gong__marshallBool(ident, "IsExpanded", requirement.IsExpanded)

	case "SupportLevels":
		var sb strings.Builder
		for _, _supportlevel := range requirement.SupportLevels {
			sb.WriteString(__gong__marshallSliceOfPointers(ident, "SupportLevels", _supportlevel.GongGetIdentifier(stage)))
		}
		res = sb.String()
	case "Concerns":
		var sb strings.Builder
		for _, _concern := range requirement.Concerns {
			sb.WriteString(__gong__marshallSliceOfPointers(ident, "Concerns", _concern.GongGetIdentifier(stage)))
		}
		res = sb.String()
	case "Views":
		var sb strings.Builder
		for _, _view := range requirement.Views {
			sb.WriteString(__gong__marshallSliceOfPointers(ident, "Views", _view.GongGetIdentifier(stage)))
		}
		res = sb.String()
	case "Concepts":
		var sb strings.Builder
		for _, _concept := range requirement.Concepts {
			sb.WriteString(__gong__marshallSliceOfPointers(ident, "Concepts", _concept.GongGetIdentifier(stage)))
		}
		res = sb.String()
	default:
		log.Panicf("Unknown field %s for Gongstruct Requirement", fieldName)
	}
	return
}

func (requirementconceptshape *RequirementConceptShape) GongMarshallField(stage *Stage, fieldName string) (res string) {
	ident := requirementconceptshape.GongGetIdentifier(stage)
	_ = ident
	switch fieldName {
	case "Name":
		res = __gong__marshallString(ident, "Name", requirementconceptshape.Name)
	case "StartRatio":
		res = __gong__marshallFloat(ident, "StartRatio", requirementconceptshape.StartRatio)
	case "EndRatio":
		res = __gong__marshallFloat(ident, "EndRatio", requirementconceptshape.EndRatio)
	case "StartOrientation":
		res = __gong__marshallEnumString(ident, "StartOrientation", requirementconceptshape.StartOrientation.ToCodeString())
	case "EndOrientation":
		res = __gong__marshallEnumString(ident, "EndOrientation", requirementconceptshape.EndOrientation.ToCodeString())
	case "CornerOffsetRatio":
		res = __gong__marshallFloat(ident, "CornerOffsetRatio", requirementconceptshape.CornerOffsetRatio)
	case "IsHidden":
		res = __gong__marshallBool(ident, "IsHidden", requirementconceptshape.IsHidden)
	case "IsPresentationOverridden":
		res = __gong__marshallBool(ident, "IsPresentationOverridden", requirementconceptshape.IsPresentationOverridden)
	case "Stroke":
		res = __gong__marshallString(ident, "Stroke", requirementconceptshape.Stroke)
	case "StrokeWidth":
		res = __gong__marshallFloat(ident, "StrokeWidth", requirementconceptshape.StrokeWidth)
	case "StrokeOpacity":
		res = __gong__marshallFloat(ident, "StrokeOpacity", requirementconceptshape.StrokeOpacity)
	case "StrokeDashArray":
		res = __gong__marshallString(ident, "StrokeDashArray", requirementconceptshape.StrokeDashArray)
	case "FontSize":
		res = __gong__marshallString(ident, "FontSize", requirementconceptshape.FontSize)
	case "FontColor":
		res = __gong__marshallString(ident, "FontColor", requirementconceptshape.FontColor)

	case "Requirement":
		if requirementconceptshape.Requirement != nil {
			res = __gong__marshallPointer(ident, "Requirement", requirementconceptshape.Requirement.GongGetIdentifier(stage))
		} else {
			res = __gong__marshallPointer(ident, "Requirement", "nil")
		}
	case "Concept":
		if requirementconceptshape.Concept != nil {
			res = __gong__marshallPointer(ident, "Concept", requirementconceptshape.Concept.GongGetIdentifier(stage))
		} else {
			res = __gong__marshallPointer(ident, "Concept", "nil")
		}
	case "ControlPointShapes":
		var sb strings.Builder
		for _, _controlpointshape := range requirementconceptshape.ControlPointShapes {
			sb.WriteString(__gong__marshallSliceOfPointers(ident, "ControlPointShapes", _controlpointshape.GongGetIdentifier(stage)))
		}
		res = sb.String()
	default:
		log.Panicf("Unknown field %s for Gongstruct RequirementConceptShape", fieldName)
	}
	return
}

func (requirementconcernshape *RequirementConcernShape) GongMarshallField(stage *Stage, fieldName string) (res string) {
	ident := requirementconcernshape.GongGetIdentifier(stage)
	_ = ident
	switch fieldName {
	case "Name":
		res = __gong__marshallString(ident, "Name", requirementconcernshape.Name)
	case "StartRatio":
		res = __gong__marshallFloat(ident, "StartRatio", requirementconcernshape.StartRatio)
	case "EndRatio":
		res = __gong__marshallFloat(ident, "EndRatio", requirementconcernshape.EndRatio)
	case "StartOrientation":
		res = __gong__marshallEnumString(ident, "StartOrientation", requirementconcernshape.StartOrientation.ToCodeString())
	case "EndOrientation":
		res = __gong__marshallEnumString(ident, "EndOrientation", requirementconcernshape.EndOrientation.ToCodeString())
	case "CornerOffsetRatio":
		res = __gong__marshallFloat(ident, "CornerOffsetRatio", requirementconcernshape.CornerOffsetRatio)
	case "IsHidden":
		res = __gong__marshallBool(ident, "IsHidden", requirementconcernshape.IsHidden)
	case "IsPresentationOverridden":
		res = __gong__marshallBool(ident, "IsPresentationOverridden", requirementconcernshape.IsPresentationOverridden)
	case "Stroke":
		res = __gong__marshallString(ident, "Stroke", requirementconcernshape.Stroke)
	case "StrokeWidth":
		res = __gong__marshallFloat(ident, "StrokeWidth", requirementconcernshape.StrokeWidth)
	case "StrokeOpacity":
		res = __gong__marshallFloat(ident, "StrokeOpacity", requirementconcernshape.StrokeOpacity)
	case "StrokeDashArray":
		res = __gong__marshallString(ident, "StrokeDashArray", requirementconcernshape.StrokeDashArray)
	case "FontSize":
		res = __gong__marshallString(ident, "FontSize", requirementconcernshape.FontSize)
	case "FontColor":
		res = __gong__marshallString(ident, "FontColor", requirementconcernshape.FontColor)

	case "Requirement":
		if requirementconcernshape.Requirement != nil {
			res = __gong__marshallPointer(ident, "Requirement", requirementconcernshape.Requirement.GongGetIdentifier(stage))
		} else {
			res = __gong__marshallPointer(ident, "Requirement", "nil")
		}
	case "Concern":
		if requirementconcernshape.Concern != nil {
			res = __gong__marshallPointer(ident, "Concern", requirementconcernshape.Concern.GongGetIdentifier(stage))
		} else {
			res = __gong__marshallPointer(ident, "Concern", "nil")
		}
	case "ControlPointShapes":
		var sb strings.Builder
		for _, _controlpointshape := range requirementconcernshape.ControlPointShapes {
			sb.WriteString(__gong__marshallSliceOfPointers(ident, "ControlPointShapes", _controlpointshape.GongGetIdentifier(stage)))
		}
		res = sb.String()
	default:
		log.Panicf("Unknown field %s for Gongstruct RequirementConcernShape", fieldName)
	}
	return
}

func (requirementshape *RequirementShape) GongMarshallField(stage *Stage, fieldName string) (res string) {
	ident := requirementshape.GongGetIdentifier(stage)
	_ = ident
	switch fieldName {
	case "Name":
		res = __gong__marshallString(ident, "Name", requirementshape.Name)
	case "IsExpanded":
		res = __gong__marshallBool(ident, "IsExpanded", requirementshape.IsExpanded)
	case "X":
		res = __gong__marshallFloat(ident, "X", requirementshape.X)
	case "Y":
		res = __gong__marshallFloat(ident, "Y", requirementshape.Y)
	case "Width":
		res = __gong__marshallFloat(ident, "Width", requirementshape.Width)
	case "Height":
		res = __gong__marshallFloat(ident, "Height", requirementshape.Height)
	case "IsHidden":
		res = __gong__marshallBool(ident, "IsHidden", requirementshape.IsHidden)
	case "IsPresentationOverridden":
		res = __gong__marshallBool(ident, "IsPresentationOverridden", requirementshape.IsPresentationOverridden)
	case "Color":
		res = __gong__marshallString(ident, "Color", requirementshape.Color)
	case "FillOpacity":
		res = __gong__marshallFloat(ident, "FillOpacity", requirementshape.FillOpacity)
	case "Stroke":
		res = __gong__marshallString(ident, "Stroke", requirementshape.Stroke)
	case "StrokeWidth":
		res = __gong__marshallFloat(ident, "StrokeWidth", requirementshape.StrokeWidth)
	case "StrokeOpacity":
		res = __gong__marshallFloat(ident, "StrokeOpacity", requirementshape.StrokeOpacity)
	case "StrokeDashArray":
		res = __gong__marshallString(ident, "StrokeDashArray", requirementshape.StrokeDashArray)
	case "RX":
		res = __gong__marshallFloat(ident, "RX", requirementshape.RX)
	case "FontSize":
		res = __gong__marshallString(ident, "FontSize", requirementshape.FontSize)
	case "FontWeight":
		res = __gong__marshallString(ident, "FontWeight", requirementshape.FontWeight)
	case "FontColor":
		res = __gong__marshallString(ident, "FontColor", requirementshape.FontColor)

	case "Requirement":
		if requirementshape.Requirement != nil {
			res = __gong__marshallPointer(ident, "Requirement", requirementshape.Requirement.GongGetIdentifier(stage))
		} else {
			res = __gong__marshallPointer(ident, "Requirement", "nil")
		}
	default:
		log.Panicf("Unknown field %s for Gongstruct RequirementShape", fieldName)
	}
	return
}

func (requirementviewshape *RequirementViewShape) GongMarshallField(stage *Stage, fieldName string) (res string) {
	ident := requirementviewshape.GongGetIdentifier(stage)
	_ = ident
	switch fieldName {
	case "Name":
		res = __gong__marshallString(ident, "Name", requirementviewshape.Name)
	case "StartRatio":
		res = __gong__marshallFloat(ident, "StartRatio", requirementviewshape.StartRatio)
	case "EndRatio":
		res = __gong__marshallFloat(ident, "EndRatio", requirementviewshape.EndRatio)
	case "StartOrientation":
		res = __gong__marshallEnumString(ident, "StartOrientation", requirementviewshape.StartOrientation.ToCodeString())
	case "EndOrientation":
		res = __gong__marshallEnumString(ident, "EndOrientation", requirementviewshape.EndOrientation.ToCodeString())
	case "CornerOffsetRatio":
		res = __gong__marshallFloat(ident, "CornerOffsetRatio", requirementviewshape.CornerOffsetRatio)
	case "IsHidden":
		res = __gong__marshallBool(ident, "IsHidden", requirementviewshape.IsHidden)
	case "IsPresentationOverridden":
		res = __gong__marshallBool(ident, "IsPresentationOverridden", requirementviewshape.IsPresentationOverridden)
	case "Stroke":
		res = __gong__marshallString(ident, "Stroke", requirementviewshape.Stroke)
	case "StrokeWidth":
		res = __gong__marshallFloat(ident, "StrokeWidth", requirementviewshape.StrokeWidth)
	case "StrokeOpacity":
		res = __gong__marshallFloat(ident, "StrokeOpacity", requirementviewshape.StrokeOpacity)
	case "StrokeDashArray":
		res = __gong__marshallString(ident, "StrokeDashArray", requirementviewshape.StrokeDashArray)
	case "FontSize":
		res = __gong__marshallString(ident, "FontSize", requirementviewshape.FontSize)
	case "FontColor":
		res = __gong__marshallString(ident, "FontColor", requirementviewshape.FontColor)

	case "Requirement":
		if requirementviewshape.Requirement != nil {
			res = __gong__marshallPointer(ident, "Requirement", requirementviewshape.Requirement.GongGetIdentifier(stage))
		} else {
			res = __gong__marshallPointer(ident, "Requirement", "nil")
		}
	case "View":
		if requirementviewshape.View != nil {
			res = __gong__marshallPointer(ident, "View", requirementviewshape.View.GongGetIdentifier(stage))
		} else {
			res = __gong__marshallPointer(ident, "View", "nil")
		}
	case "ControlPointShapes":
		var sb strings.Builder
		for _, _controlpointshape := range requirementviewshape.ControlPointShapes {
			sb.WriteString(__gong__marshallSliceOfPointers(ident, "ControlPointShapes", _controlpointshape.GongGetIdentifier(stage)))
		}
		res = sb.String()
	default:
		log.Panicf("Unknown field %s for Gongstruct RequirementViewShape", fieldName)
	}
	return
}

func (stakeholder *Stakeholder) GongMarshallField(stage *Stage, fieldName string) (res string) {
	ident := stakeholder.GongGetIdentifier(stage)
	_ = ident
	switch fieldName {
	case "Name":
		res = __gong__marshallString(ident, "Name", stakeholder.Name)
	case "IDAirbus":
		res = __gong__marshallString(ident, "IDAirbus", stakeholder.IDAirbus)
	case "ComputedPrefix":
		res = __gong__marshallString(ident, "ComputedPrefix", stakeholder.ComputedPrefix)
	case "IsExpanded":
		res = __gong__marshallBool(ident, "IsExpanded", stakeholder.IsExpanded)
	case "Description":
		res = __gong__marshallString(ident, "Description", stakeholder.Description)

	case "Concerns":
		var sb strings.Builder
		for _, _concern := range stakeholder.Concerns {
			sb.WriteString(__gong__marshallSliceOfPointers(ident, "Concerns", _concern.GongGetIdentifier(stage)))
		}
		res = sb.String()
	case "SubStakeholders":
		var sb strings.Builder
		for _, _stakeholder := range stakeholder.SubStakeholders {
			sb.WriteString(__gong__marshallSliceOfPointers(ident, "SubStakeholders", _stakeholder.GongGetIdentifier(stage)))
		}
		res = sb.String()
	default:
		log.Panicf("Unknown field %s for Gongstruct Stakeholder", fieldName)
	}
	return
}

func (stakeholdercompositionshape *StakeholderCompositionShape) GongMarshallField(stage *Stage, fieldName string) (res string) {
	ident := stakeholdercompositionshape.GongGetIdentifier(stage)
	_ = ident
	switch fieldName {
	case "Name":
		res = __gong__marshallString(ident, "Name", stakeholdercompositionshape.Name)
	case "StartRatio":
		res = __gong__marshallFloat(ident, "StartRatio", stakeholdercompositionshape.StartRatio)
	case "EndRatio":
		res = __gong__marshallFloat(ident, "EndRatio", stakeholdercompositionshape.EndRatio)
	case "StartOrientation":
		res = __gong__marshallEnumString(ident, "StartOrientation", stakeholdercompositionshape.StartOrientation.ToCodeString())
	case "EndOrientation":
		res = __gong__marshallEnumString(ident, "EndOrientation", stakeholdercompositionshape.EndOrientation.ToCodeString())
	case "CornerOffsetRatio":
		res = __gong__marshallFloat(ident, "CornerOffsetRatio", stakeholdercompositionshape.CornerOffsetRatio)
	case "IsHidden":
		res = __gong__marshallBool(ident, "IsHidden", stakeholdercompositionshape.IsHidden)
	case "IsPresentationOverridden":
		res = __gong__marshallBool(ident, "IsPresentationOverridden", stakeholdercompositionshape.IsPresentationOverridden)
	case "Stroke":
		res = __gong__marshallString(ident, "Stroke", stakeholdercompositionshape.Stroke)
	case "StrokeWidth":
		res = __gong__marshallFloat(ident, "StrokeWidth", stakeholdercompositionshape.StrokeWidth)
	case "StrokeOpacity":
		res = __gong__marshallFloat(ident, "StrokeOpacity", stakeholdercompositionshape.StrokeOpacity)
	case "StrokeDashArray":
		res = __gong__marshallString(ident, "StrokeDashArray", stakeholdercompositionshape.StrokeDashArray)
	case "FontSize":
		res = __gong__marshallString(ident, "FontSize", stakeholdercompositionshape.FontSize)
	case "FontColor":
		res = __gong__marshallString(ident, "FontColor", stakeholdercompositionshape.FontColor)

	case "Stakeholder":
		if stakeholdercompositionshape.Stakeholder != nil {
			res = __gong__marshallPointer(ident, "Stakeholder", stakeholdercompositionshape.Stakeholder.GongGetIdentifier(stage))
		} else {
			res = __gong__marshallPointer(ident, "Stakeholder", "nil")
		}
	case "ControlPointShapes":
		var sb strings.Builder
		for _, _controlpointshape := range stakeholdercompositionshape.ControlPointShapes {
			sb.WriteString(__gong__marshallSliceOfPointers(ident, "ControlPointShapes", _controlpointshape.GongGetIdentifier(stage)))
		}
		res = sb.String()
	default:
		log.Panicf("Unknown field %s for Gongstruct StakeholderCompositionShape", fieldName)
	}
	return
}

func (stakeholderconcernshape *StakeholderConcernShape) GongMarshallField(stage *Stage, fieldName string) (res string) {
	ident := stakeholderconcernshape.GongGetIdentifier(stage)
	_ = ident
	switch fieldName {
	case "Name":
		res = __gong__marshallString(ident, "Name", stakeholderconcernshape.Name)
	case "StartRatio":
		res = __gong__marshallFloat(ident, "StartRatio", stakeholderconcernshape.StartRatio)
	case "EndRatio":
		res = __gong__marshallFloat(ident, "EndRatio", stakeholderconcernshape.EndRatio)
	case "StartOrientation":
		res = __gong__marshallEnumString(ident, "StartOrientation", stakeholderconcernshape.StartOrientation.ToCodeString())
	case "EndOrientation":
		res = __gong__marshallEnumString(ident, "EndOrientation", stakeholderconcernshape.EndOrientation.ToCodeString())
	case "CornerOffsetRatio":
		res = __gong__marshallFloat(ident, "CornerOffsetRatio", stakeholderconcernshape.CornerOffsetRatio)
	case "IsHidden":
		res = __gong__marshallBool(ident, "IsHidden", stakeholderconcernshape.IsHidden)
	case "IsPresentationOverridden":
		res = __gong__marshallBool(ident, "IsPresentationOverridden", stakeholderconcernshape.IsPresentationOverridden)
	case "Stroke":
		res = __gong__marshallString(ident, "Stroke", stakeholderconcernshape.Stroke)
	case "StrokeWidth":
		res = __gong__marshallFloat(ident, "StrokeWidth", stakeholderconcernshape.StrokeWidth)
	case "StrokeOpacity":
		res = __gong__marshallFloat(ident, "StrokeOpacity", stakeholderconcernshape.StrokeOpacity)
	case "StrokeDashArray":
		res = __gong__marshallString(ident, "StrokeDashArray", stakeholderconcernshape.StrokeDashArray)
	case "FontSize":
		res = __gong__marshallString(ident, "FontSize", stakeholderconcernshape.FontSize)
	case "FontColor":
		res = __gong__marshallString(ident, "FontColor", stakeholderconcernshape.FontColor)

	case "Stakeholder":
		if stakeholderconcernshape.Stakeholder != nil {
			res = __gong__marshallPointer(ident, "Stakeholder", stakeholderconcernshape.Stakeholder.GongGetIdentifier(stage))
		} else {
			res = __gong__marshallPointer(ident, "Stakeholder", "nil")
		}
	case "Concern":
		if stakeholderconcernshape.Concern != nil {
			res = __gong__marshallPointer(ident, "Concern", stakeholderconcernshape.Concern.GongGetIdentifier(stage))
		} else {
			res = __gong__marshallPointer(ident, "Concern", "nil")
		}
	case "ControlPointShapes":
		var sb strings.Builder
		for _, _controlpointshape := range stakeholderconcernshape.ControlPointShapes {
			sb.WriteString(__gong__marshallSliceOfPointers(ident, "ControlPointShapes", _controlpointshape.GongGetIdentifier(stage)))
		}
		res = sb.String()
	default:
		log.Panicf("Unknown field %s for Gongstruct StakeholderConcernShape", fieldName)
	}
	return
}

func (stakeholdershape *StakeholderShape) GongMarshallField(stage *Stage, fieldName string) (res string) {
	ident := stakeholdershape.GongGetIdentifier(stage)
	_ = ident
	switch fieldName {
	case "Name":
		res = __gong__marshallString(ident, "Name", stakeholdershape.Name)
	case "IsExpanded":
		res = __gong__marshallBool(ident, "IsExpanded", stakeholdershape.IsExpanded)
	case "X":
		res = __gong__marshallFloat(ident, "X", stakeholdershape.X)
	case "Y":
		res = __gong__marshallFloat(ident, "Y", stakeholdershape.Y)
	case "Width":
		res = __gong__marshallFloat(ident, "Width", stakeholdershape.Width)
	case "Height":
		res = __gong__marshallFloat(ident, "Height", stakeholdershape.Height)
	case "IsHidden":
		res = __gong__marshallBool(ident, "IsHidden", stakeholdershape.IsHidden)
	case "IsPresentationOverridden":
		res = __gong__marshallBool(ident, "IsPresentationOverridden", stakeholdershape.IsPresentationOverridden)
	case "Color":
		res = __gong__marshallString(ident, "Color", stakeholdershape.Color)
	case "FillOpacity":
		res = __gong__marshallFloat(ident, "FillOpacity", stakeholdershape.FillOpacity)
	case "Stroke":
		res = __gong__marshallString(ident, "Stroke", stakeholdershape.Stroke)
	case "StrokeWidth":
		res = __gong__marshallFloat(ident, "StrokeWidth", stakeholdershape.StrokeWidth)
	case "StrokeOpacity":
		res = __gong__marshallFloat(ident, "StrokeOpacity", stakeholdershape.StrokeOpacity)
	case "StrokeDashArray":
		res = __gong__marshallString(ident, "StrokeDashArray", stakeholdershape.StrokeDashArray)
	case "RX":
		res = __gong__marshallFloat(ident, "RX", stakeholdershape.RX)
	case "FontSize":
		res = __gong__marshallString(ident, "FontSize", stakeholdershape.FontSize)
	case "FontWeight":
		res = __gong__marshallString(ident, "FontWeight", stakeholdershape.FontWeight)
	case "FontColor":
		res = __gong__marshallString(ident, "FontColor", stakeholdershape.FontColor)

	case "Stakeholder":
		if stakeholdershape.Stakeholder != nil {
			res = __gong__marshallPointer(ident, "Stakeholder", stakeholdershape.Stakeholder.GongGetIdentifier(stage))
		} else {
			res = __gong__marshallPointer(ident, "Stakeholder", "nil")
		}
	default:
		log.Panicf("Unknown field %s for Gongstruct StakeholderShape", fieldName)
	}
	return
}

func (supportlevel *SupportLevel) GongMarshallField(stage *Stage, fieldName string) (res string) {
	ident := supportlevel.GongGetIdentifier(stage)
	_ = ident
	switch fieldName {
	case "Name":
		res = __gong__marshallString(ident, "Name", supportlevel.Name)

	case "Tool":
		if supportlevel.Tool != nil {
			res = __gong__marshallPointer(ident, "Tool", supportlevel.Tool.GongGetIdentifier(stage))
		} else {
			res = __gong__marshallPointer(ident, "Tool", "nil")
		}
	default:
		log.Panicf("Unknown field %s for Gongstruct SupportLevel", fieldName)
	}
	return
}

func (tool *Tool) GongMarshallField(stage *Stage, fieldName string) (res string) {
	ident := tool.GongGetIdentifier(stage)
	_ = ident
	switch fieldName {
	case "Name":
		res = __gong__marshallString(ident, "Name", tool.Name)

	default:
		log.Panicf("Unknown field %s for Gongstruct Tool", fieldName)
	}
	return
}

func (view *View) GongMarshallField(stage *Stage, fieldName string) (res string) {
	ident := view.GongGetIdentifier(stage)
	_ = ident
	switch fieldName {
	case "Name":
		res = __gong__marshallString(ident, "Name", view.Name)
	case "ComputedPrefix":
		res = __gong__marshallString(ident, "ComputedPrefix", view.ComputedPrefix)
	case "IsExpanded":
		res = __gong__marshallBool(ident, "IsExpanded", view.IsExpanded)
	case "Description":
		res = __gong__marshallString(ident, "Description", view.Description)

	case "Concepts":
		var sb strings.Builder
		for _, _concept := range view.Concepts {
			sb.WriteString(__gong__marshallSliceOfPointers(ident, "Concepts", _concept.GongGetIdentifier(stage)))
		}
		res = sb.String()
	default:
		log.Panicf("Unknown field %s for Gongstruct View", fieldName)
	}
	return
}

func (viewconceptshape *ViewConceptShape) GongMarshallField(stage *Stage, fieldName string) (res string) {
	ident := viewconceptshape.GongGetIdentifier(stage)
	_ = ident
	switch fieldName {
	case "Name":
		res = __gong__marshallString(ident, "Name", viewconceptshape.Name)
	case "StartRatio":
		res = __gong__marshallFloat(ident, "StartRatio", viewconceptshape.StartRatio)
	case "EndRatio":
		res = __gong__marshallFloat(ident, "EndRatio", viewconceptshape.EndRatio)
	case "StartOrientation":
		res = __gong__marshallEnumString(ident, "StartOrientation", viewconceptshape.StartOrientation.ToCodeString())
	case "EndOrientation":
		res = __gong__marshallEnumString(ident, "EndOrientation", viewconceptshape.EndOrientation.ToCodeString())
	case "CornerOffsetRatio":
		res = __gong__marshallFloat(ident, "CornerOffsetRatio", viewconceptshape.CornerOffsetRatio)
	case "IsHidden":
		res = __gong__marshallBool(ident, "IsHidden", viewconceptshape.IsHidden)
	case "IsPresentationOverridden":
		res = __gong__marshallBool(ident, "IsPresentationOverridden", viewconceptshape.IsPresentationOverridden)
	case "Stroke":
		res = __gong__marshallString(ident, "Stroke", viewconceptshape.Stroke)
	case "StrokeWidth":
		res = __gong__marshallFloat(ident, "StrokeWidth", viewconceptshape.StrokeWidth)
	case "StrokeOpacity":
		res = __gong__marshallFloat(ident, "StrokeOpacity", viewconceptshape.StrokeOpacity)
	case "StrokeDashArray":
		res = __gong__marshallString(ident, "StrokeDashArray", viewconceptshape.StrokeDashArray)
	case "FontSize":
		res = __gong__marshallString(ident, "FontSize", viewconceptshape.FontSize)
	case "FontColor":
		res = __gong__marshallString(ident, "FontColor", viewconceptshape.FontColor)

	case "View":
		if viewconceptshape.View != nil {
			res = __gong__marshallPointer(ident, "View", viewconceptshape.View.GongGetIdentifier(stage))
		} else {
			res = __gong__marshallPointer(ident, "View", "nil")
		}
	case "Concept":
		if viewconceptshape.Concept != nil {
			res = __gong__marshallPointer(ident, "Concept", viewconceptshape.Concept.GongGetIdentifier(stage))
		} else {
			res = __gong__marshallPointer(ident, "Concept", "nil")
		}
	case "ControlPointShapes":
		var sb strings.Builder
		for _, _controlpointshape := range viewconceptshape.ControlPointShapes {
			sb.WriteString(__gong__marshallSliceOfPointers(ident, "ControlPointShapes", _controlpointshape.GongGetIdentifier(stage)))
		}
		res = sb.String()
	default:
		log.Panicf("Unknown field %s for Gongstruct ViewConceptShape", fieldName)
	}
	return
}

func (viewshape *ViewShape) GongMarshallField(stage *Stage, fieldName string) (res string) {
	ident := viewshape.GongGetIdentifier(stage)
	_ = ident
	switch fieldName {
	case "Name":
		res = __gong__marshallString(ident, "Name", viewshape.Name)
	case "IsExpanded":
		res = __gong__marshallBool(ident, "IsExpanded", viewshape.IsExpanded)
	case "X":
		res = __gong__marshallFloat(ident, "X", viewshape.X)
	case "Y":
		res = __gong__marshallFloat(ident, "Y", viewshape.Y)
	case "Width":
		res = __gong__marshallFloat(ident, "Width", viewshape.Width)
	case "Height":
		res = __gong__marshallFloat(ident, "Height", viewshape.Height)
	case "IsHidden":
		res = __gong__marshallBool(ident, "IsHidden", viewshape.IsHidden)
	case "IsPresentationOverridden":
		res = __gong__marshallBool(ident, "IsPresentationOverridden", viewshape.IsPresentationOverridden)
	case "Color":
		res = __gong__marshallString(ident, "Color", viewshape.Color)
	case "FillOpacity":
		res = __gong__marshallFloat(ident, "FillOpacity", viewshape.FillOpacity)
	case "Stroke":
		res = __gong__marshallString(ident, "Stroke", viewshape.Stroke)
	case "StrokeWidth":
		res = __gong__marshallFloat(ident, "StrokeWidth", viewshape.StrokeWidth)
	case "StrokeOpacity":
		res = __gong__marshallFloat(ident, "StrokeOpacity", viewshape.StrokeOpacity)
	case "StrokeDashArray":
		res = __gong__marshallString(ident, "StrokeDashArray", viewshape.StrokeDashArray)
	case "RX":
		res = __gong__marshallFloat(ident, "RX", viewshape.RX)
	case "FontSize":
		res = __gong__marshallString(ident, "FontSize", viewshape.FontSize)
	case "FontWeight":
		res = __gong__marshallString(ident, "FontWeight", viewshape.FontWeight)
	case "FontColor":
		res = __gong__marshallString(ident, "FontColor", viewshape.FontColor)

	case "View":
		if viewshape.View != nil {
			res = __gong__marshallPointer(ident, "View", viewshape.View.GongGetIdentifier(stage))
		} else {
			res = __gong__marshallPointer(ident, "View", "nil")
		}
	default:
		log.Panicf("Unknown field %s for Gongstruct ViewShape", fieldName)
	}
	return
}

// insertion point for marshall all fields methods
func (analysisneed *AnalysisNeed) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(analysisneed.GongMarshallField(stage, "Name"))
		initializerStatements.WriteString(analysisneed.GongMarshallField(stage, "ComputedPrefix"))
		initializerStatements.WriteString(analysisneed.GongMarshallField(stage, "IsExpanded"))
	}
	initRes = initializerStatements.String()
	ptrRes = pointersInitializesStatements.String()
	return
}
func (concept *Concept) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(concept.GongMarshallField(stage, "Name"))
		initializerStatements.WriteString(concept.GongMarshallField(stage, "ComputedPrefix"))
		initializerStatements.WriteString(concept.GongMarshallField(stage, "IsExpanded"))
		pointersInitializesStatements.WriteString(concept.GongMarshallField(stage, "Tools"))
	}
	initRes = initializerStatements.String()
	ptrRes = pointersInitializesStatements.String()
	return
}
func (conceptshape *ConceptShape) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(conceptshape.GongMarshallField(stage, "Name"))
		pointersInitializesStatements.WriteString(conceptshape.GongMarshallField(stage, "Concept"))
		initializerStatements.WriteString(conceptshape.GongMarshallField(stage, "IsExpanded"))
		initializerStatements.WriteString(conceptshape.GongMarshallField(stage, "X"))
		initializerStatements.WriteString(conceptshape.GongMarshallField(stage, "Y"))
		initializerStatements.WriteString(conceptshape.GongMarshallField(stage, "Width"))
		initializerStatements.WriteString(conceptshape.GongMarshallField(stage, "Height"))
		initializerStatements.WriteString(conceptshape.GongMarshallField(stage, "IsHidden"))
		initializerStatements.WriteString(conceptshape.GongMarshallField(stage, "IsPresentationOverridden"))
		initializerStatements.WriteString(conceptshape.GongMarshallField(stage, "Color"))
		initializerStatements.WriteString(conceptshape.GongMarshallField(stage, "FillOpacity"))
		initializerStatements.WriteString(conceptshape.GongMarshallField(stage, "Stroke"))
		initializerStatements.WriteString(conceptshape.GongMarshallField(stage, "StrokeWidth"))
		initializerStatements.WriteString(conceptshape.GongMarshallField(stage, "StrokeOpacity"))
		initializerStatements.WriteString(conceptshape.GongMarshallField(stage, "StrokeDashArray"))
		initializerStatements.WriteString(conceptshape.GongMarshallField(stage, "RX"))
		initializerStatements.WriteString(conceptshape.GongMarshallField(stage, "FontSize"))
		initializerStatements.WriteString(conceptshape.GongMarshallField(stage, "FontWeight"))
		initializerStatements.WriteString(conceptshape.GongMarshallField(stage, "FontColor"))
	}
	initRes = initializerStatements.String()
	ptrRes = pointersInitializesStatements.String()
	return
}
func (concern *Concern) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(concern.GongMarshallField(stage, "Name"))
		initializerStatements.WriteString(concern.GongMarshallField(stage, "IDAirbus"))
		initializerStatements.WriteString(concern.GongMarshallField(stage, "Priority"))
		initializerStatements.WriteString(concern.GongMarshallField(stage, "ComputedPrefix"))
		initializerStatements.WriteString(concern.GongMarshallField(stage, "IsExpanded"))
		initializerStatements.WriteString(concern.GongMarshallField(stage, "Description"))
		pointersInitializesStatements.WriteString(concern.GongMarshallField(stage, "SubConcerns"))
		pointersInitializesStatements.WriteString(concern.GongMarshallField(stage, "Inputs"))
		initializerStatements.WriteString(concern.GongMarshallField(stage, "IsInputsNodeExpanded"))
		pointersInitializesStatements.WriteString(concern.GongMarshallField(stage, "Outputs"))
		initializerStatements.WriteString(concern.GongMarshallField(stage, "IsOutputsNodeExpanded"))
		initializerStatements.WriteString(concern.GongMarshallField(stage, "IsWithCompletion"))
		initializerStatements.WriteString(concern.GongMarshallField(stage, "Completion"))
	}
	initRes = initializerStatements.String()
	ptrRes = pointersInitializesStatements.String()
	return
}
func (concerncompositionshape *ConcernCompositionShape) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(concerncompositionshape.GongMarshallField(stage, "Name"))
		pointersInitializesStatements.WriteString(concerncompositionshape.GongMarshallField(stage, "Concern"))
		initializerStatements.WriteString(concerncompositionshape.GongMarshallField(stage, "StartRatio"))
		initializerStatements.WriteString(concerncompositionshape.GongMarshallField(stage, "EndRatio"))
		initializerStatements.WriteString(concerncompositionshape.GongMarshallField(stage, "StartOrientation"))
		initializerStatements.WriteString(concerncompositionshape.GongMarshallField(stage, "EndOrientation"))
		initializerStatements.WriteString(concerncompositionshape.GongMarshallField(stage, "CornerOffsetRatio"))
		initializerStatements.WriteString(concerncompositionshape.GongMarshallField(stage, "IsHidden"))
		initializerStatements.WriteString(concerncompositionshape.GongMarshallField(stage, "IsPresentationOverridden"))
		initializerStatements.WriteString(concerncompositionshape.GongMarshallField(stage, "Stroke"))
		initializerStatements.WriteString(concerncompositionshape.GongMarshallField(stage, "StrokeWidth"))
		initializerStatements.WriteString(concerncompositionshape.GongMarshallField(stage, "StrokeOpacity"))
		initializerStatements.WriteString(concerncompositionshape.GongMarshallField(stage, "StrokeDashArray"))
		initializerStatements.WriteString(concerncompositionshape.GongMarshallField(stage, "FontSize"))
		initializerStatements.WriteString(concerncompositionshape.GongMarshallField(stage, "FontColor"))
		pointersInitializesStatements.WriteString(concerncompositionshape.GongMarshallField(stage, "ControlPointShapes"))
	}
	initRes = initializerStatements.String()
	ptrRes = pointersInitializesStatements.String()
	return
}
func (concerninputshape *ConcernInputShape) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(concerninputshape.GongMarshallField(stage, "Name"))
		pointersInitializesStatements.WriteString(concerninputshape.GongMarshallField(stage, "Deliverable"))
		pointersInitializesStatements.WriteString(concerninputshape.GongMarshallField(stage, "Concern"))
		initializerStatements.WriteString(concerninputshape.GongMarshallField(stage, "StartRatio"))
		initializerStatements.WriteString(concerninputshape.GongMarshallField(stage, "EndRatio"))
		initializerStatements.WriteString(concerninputshape.GongMarshallField(stage, "StartOrientation"))
		initializerStatements.WriteString(concerninputshape.GongMarshallField(stage, "EndOrientation"))
		initializerStatements.WriteString(concerninputshape.GongMarshallField(stage, "CornerOffsetRatio"))
		initializerStatements.WriteString(concerninputshape.GongMarshallField(stage, "IsHidden"))
		initializerStatements.WriteString(concerninputshape.GongMarshallField(stage, "IsPresentationOverridden"))
		initializerStatements.WriteString(concerninputshape.GongMarshallField(stage, "Stroke"))
		initializerStatements.WriteString(concerninputshape.GongMarshallField(stage, "StrokeWidth"))
		initializerStatements.WriteString(concerninputshape.GongMarshallField(stage, "StrokeOpacity"))
		initializerStatements.WriteString(concerninputshape.GongMarshallField(stage, "StrokeDashArray"))
		initializerStatements.WriteString(concerninputshape.GongMarshallField(stage, "FontSize"))
		initializerStatements.WriteString(concerninputshape.GongMarshallField(stage, "FontColor"))
		pointersInitializesStatements.WriteString(concerninputshape.GongMarshallField(stage, "ControlPointShapes"))
	}
	initRes = initializerStatements.String()
	ptrRes = pointersInitializesStatements.String()
	return
}
func (concernoutputshape *ConcernOutputShape) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(concernoutputshape.GongMarshallField(stage, "Name"))
		pointersInitializesStatements.WriteString(concernoutputshape.GongMarshallField(stage, "Concern"))
		pointersInitializesStatements.WriteString(concernoutputshape.GongMarshallField(stage, "Deliverable"))
		initializerStatements.WriteString(concernoutputshape.GongMarshallField(stage, "StartRatio"))
		initializerStatements.WriteString(concernoutputshape.GongMarshallField(stage, "EndRatio"))
		initializerStatements.WriteString(concernoutputshape.GongMarshallField(stage, "StartOrientation"))
		initializerStatements.WriteString(concernoutputshape.GongMarshallField(stage, "EndOrientation"))
		initializerStatements.WriteString(concernoutputshape.GongMarshallField(stage, "CornerOffsetRatio"))
		initializerStatements.WriteString(concernoutputshape.GongMarshallField(stage, "IsHidden"))
		initializerStatements.WriteString(concernoutputshape.GongMarshallField(stage, "IsPresentationOverridden"))
		initializerStatements.WriteString(concernoutputshape.GongMarshallField(stage, "Stroke"))
		initializerStatements.WriteString(concernoutputshape.GongMarshallField(stage, "StrokeWidth"))
		initializerStatements.WriteString(concernoutputshape.GongMarshallField(stage, "StrokeOpacity"))
		initializerStatements.WriteString(concernoutputshape.GongMarshallField(stage, "StrokeDashArray"))
		initializerStatements.WriteString(concernoutputshape.GongMarshallField(stage, "FontSize"))
		initializerStatements.WriteString(concernoutputshape.GongMarshallField(stage, "FontColor"))
		pointersInitializesStatements.WriteString(concernoutputshape.GongMarshallField(stage, "ControlPointShapes"))
	}
	initRes = initializerStatements.String()
	ptrRes = pointersInitializesStatements.String()
	return
}
func (concernshape *ConcernShape) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(concernshape.GongMarshallField(stage, "Name"))
		pointersInitializesStatements.WriteString(concernshape.GongMarshallField(stage, "Concern"))
		initializerStatements.WriteString(concernshape.GongMarshallField(stage, "IsExpanded"))
		initializerStatements.WriteString(concernshape.GongMarshallField(stage, "X"))
		initializerStatements.WriteString(concernshape.GongMarshallField(stage, "Y"))
		initializerStatements.WriteString(concernshape.GongMarshallField(stage, "Width"))
		initializerStatements.WriteString(concernshape.GongMarshallField(stage, "Height"))
		initializerStatements.WriteString(concernshape.GongMarshallField(stage, "IsHidden"))
		initializerStatements.WriteString(concernshape.GongMarshallField(stage, "IsPresentationOverridden"))
		initializerStatements.WriteString(concernshape.GongMarshallField(stage, "Color"))
		initializerStatements.WriteString(concernshape.GongMarshallField(stage, "FillOpacity"))
		initializerStatements.WriteString(concernshape.GongMarshallField(stage, "Stroke"))
		initializerStatements.WriteString(concernshape.GongMarshallField(stage, "StrokeWidth"))
		initializerStatements.WriteString(concernshape.GongMarshallField(stage, "StrokeOpacity"))
		initializerStatements.WriteString(concernshape.GongMarshallField(stage, "StrokeDashArray"))
		initializerStatements.WriteString(concernshape.GongMarshallField(stage, "RX"))
		initializerStatements.WriteString(concernshape.GongMarshallField(stage, "FontSize"))
		initializerStatements.WriteString(concernshape.GongMarshallField(stage, "FontWeight"))
		initializerStatements.WriteString(concernshape.GongMarshallField(stage, "FontColor"))
	}
	initRes = initializerStatements.String()
	ptrRes = pointersInitializesStatements.String()
	return
}
func (controlpointshape *ControlPointShape) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(controlpointshape.GongMarshallField(stage, "Name"))
		initializerStatements.WriteString(controlpointshape.GongMarshallField(stage, "X_Relative"))
		initializerStatements.WriteString(controlpointshape.GongMarshallField(stage, "Y_Relative"))
		initializerStatements.WriteString(controlpointshape.GongMarshallField(stage, "IsStartShapeTheClosestShape"))
	}
	initRes = initializerStatements.String()
	ptrRes = pointersInitializesStatements.String()
	return
}
func (deliverable *Deliverable) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(deliverable.GongMarshallField(stage, "Name"))
		initializerStatements.WriteString(deliverable.GongMarshallField(stage, "ComputedPrefix"))
		initializerStatements.WriteString(deliverable.GongMarshallField(stage, "IsExpanded"))
		initializerStatements.WriteString(deliverable.GongMarshallField(stage, "Description"))
		pointersInitializesStatements.WriteString(deliverable.GongMarshallField(stage, "SubDeliverables"))
		initializerStatements.WriteString(deliverable.GongMarshallField(stage, "IsProducersNodeExpanded"))
		initializerStatements.WriteString(deliverable.GongMarshallField(stage, "IsConsumersNodeExpanded"))
		pointersInitializesStatements.WriteString(deliverable.GongMarshallField(stage, "Views"))
		pointersInitializesStatements.WriteString(deliverable.GongMarshallField(stage, "Concepts"))
	}
	initRes = initializerStatements.String()
	ptrRes = pointersInitializesStatements.String()
	return
}
func (deliverablecompositionshape *DeliverableCompositionShape) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(deliverablecompositionshape.GongMarshallField(stage, "Name"))
		pointersInitializesStatements.WriteString(deliverablecompositionshape.GongMarshallField(stage, "Deliverable"))
		initializerStatements.WriteString(deliverablecompositionshape.GongMarshallField(stage, "StartRatio"))
		initializerStatements.WriteString(deliverablecompositionshape.GongMarshallField(stage, "EndRatio"))
		initializerStatements.WriteString(deliverablecompositionshape.GongMarshallField(stage, "StartOrientation"))
		initializerStatements.WriteString(deliverablecompositionshape.GongMarshallField(stage, "EndOrientation"))
		initializerStatements.WriteString(deliverablecompositionshape.GongMarshallField(stage, "CornerOffsetRatio"))
		initializerStatements.WriteString(deliverablecompositionshape.GongMarshallField(stage, "IsHidden"))
		initializerStatements.WriteString(deliverablecompositionshape.GongMarshallField(stage, "IsPresentationOverridden"))
		initializerStatements.WriteString(deliverablecompositionshape.GongMarshallField(stage, "Stroke"))
		initializerStatements.WriteString(deliverablecompositionshape.GongMarshallField(stage, "StrokeWidth"))
		initializerStatements.WriteString(deliverablecompositionshape.GongMarshallField(stage, "StrokeOpacity"))
		initializerStatements.WriteString(deliverablecompositionshape.GongMarshallField(stage, "StrokeDashArray"))
		initializerStatements.WriteString(deliverablecompositionshape.GongMarshallField(stage, "FontSize"))
		initializerStatements.WriteString(deliverablecompositionshape.GongMarshallField(stage, "FontColor"))
		pointersInitializesStatements.WriteString(deliverablecompositionshape.GongMarshallField(stage, "ControlPointShapes"))
	}
	initRes = initializerStatements.String()
	ptrRes = pointersInitializesStatements.String()
	return
}
func (deliverableconceptshape *DeliverableConceptShape) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(deliverableconceptshape.GongMarshallField(stage, "Name"))
		pointersInitializesStatements.WriteString(deliverableconceptshape.GongMarshallField(stage, "Deliverable"))
		pointersInitializesStatements.WriteString(deliverableconceptshape.GongMarshallField(stage, "Concept"))
		initializerStatements.WriteString(deliverableconceptshape.GongMarshallField(stage, "StartRatio"))
		initializerStatements.WriteString(deliverableconceptshape.GongMarshallField(stage, "EndRatio"))
		initializerStatements.WriteString(deliverableconceptshape.GongMarshallField(stage, "StartOrientation"))
		initializerStatements.WriteString(deliverableconceptshape.GongMarshallField(stage, "EndOrientation"))
		initializerStatements.WriteString(deliverableconceptshape.GongMarshallField(stage, "CornerOffsetRatio"))
		initializerStatements.WriteString(deliverableconceptshape.GongMarshallField(stage, "IsHidden"))
		initializerStatements.WriteString(deliverableconceptshape.GongMarshallField(stage, "IsPresentationOverridden"))
		initializerStatements.WriteString(deliverableconceptshape.GongMarshallField(stage, "Stroke"))
		initializerStatements.WriteString(deliverableconceptshape.GongMarshallField(stage, "StrokeWidth"))
		initializerStatements.WriteString(deliverableconceptshape.GongMarshallField(stage, "StrokeOpacity"))
		initializerStatements.WriteString(deliverableconceptshape.GongMarshallField(stage, "StrokeDashArray"))
		initializerStatements.WriteString(deliverableconceptshape.GongMarshallField(stage, "FontSize"))
		initializerStatements.WriteString(deliverableconceptshape.GongMarshallField(stage, "FontColor"))
		pointersInitializesStatements.WriteString(deliverableconceptshape.GongMarshallField(stage, "ControlPointShapes"))
	}
	initRes = initializerStatements.String()
	ptrRes = pointersInitializesStatements.String()
	return
}
func (deliverableshape *DeliverableShape) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(deliverableshape.GongMarshallField(stage, "Name"))
		pointersInitializesStatements.WriteString(deliverableshape.GongMarshallField(stage, "Deliverable"))
		initializerStatements.WriteString(deliverableshape.GongMarshallField(stage, "IsExpanded"))
		initializerStatements.WriteString(deliverableshape.GongMarshallField(stage, "X"))
		initializerStatements.WriteString(deliverableshape.GongMarshallField(stage, "Y"))
		initializerStatements.WriteString(deliverableshape.GongMarshallField(stage, "Width"))
		initializerStatements.WriteString(deliverableshape.GongMarshallField(stage, "Height"))
		initializerStatements.WriteString(deliverableshape.GongMarshallField(stage, "IsHidden"))
		initializerStatements.WriteString(deliverableshape.GongMarshallField(stage, "IsPresentationOverridden"))
		initializerStatements.WriteString(deliverableshape.GongMarshallField(stage, "Color"))
		initializerStatements.WriteString(deliverableshape.GongMarshallField(stage, "FillOpacity"))
		initializerStatements.WriteString(deliverableshape.GongMarshallField(stage, "Stroke"))
		initializerStatements.WriteString(deliverableshape.GongMarshallField(stage, "StrokeWidth"))
		initializerStatements.WriteString(deliverableshape.GongMarshallField(stage, "StrokeOpacity"))
		initializerStatements.WriteString(deliverableshape.GongMarshallField(stage, "StrokeDashArray"))
		initializerStatements.WriteString(deliverableshape.GongMarshallField(stage, "RX"))
		initializerStatements.WriteString(deliverableshape.GongMarshallField(stage, "FontSize"))
		initializerStatements.WriteString(deliverableshape.GongMarshallField(stage, "FontWeight"))
		initializerStatements.WriteString(deliverableshape.GongMarshallField(stage, "FontColor"))
	}
	initRes = initializerStatements.String()
	ptrRes = pointersInitializesStatements.String()
	return
}
func (deliverableviewshape *DeliverableViewShape) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(deliverableviewshape.GongMarshallField(stage, "Name"))
		pointersInitializesStatements.WriteString(deliverableviewshape.GongMarshallField(stage, "Deliverable"))
		pointersInitializesStatements.WriteString(deliverableviewshape.GongMarshallField(stage, "View"))
		initializerStatements.WriteString(deliverableviewshape.GongMarshallField(stage, "StartRatio"))
		initializerStatements.WriteString(deliverableviewshape.GongMarshallField(stage, "EndRatio"))
		initializerStatements.WriteString(deliverableviewshape.GongMarshallField(stage, "StartOrientation"))
		initializerStatements.WriteString(deliverableviewshape.GongMarshallField(stage, "EndOrientation"))
		initializerStatements.WriteString(deliverableviewshape.GongMarshallField(stage, "CornerOffsetRatio"))
		initializerStatements.WriteString(deliverableviewshape.GongMarshallField(stage, "IsHidden"))
		initializerStatements.WriteString(deliverableviewshape.GongMarshallField(stage, "IsPresentationOverridden"))
		initializerStatements.WriteString(deliverableviewshape.GongMarshallField(stage, "Stroke"))
		initializerStatements.WriteString(deliverableviewshape.GongMarshallField(stage, "StrokeWidth"))
		initializerStatements.WriteString(deliverableviewshape.GongMarshallField(stage, "StrokeOpacity"))
		initializerStatements.WriteString(deliverableviewshape.GongMarshallField(stage, "StrokeDashArray"))
		initializerStatements.WriteString(deliverableviewshape.GongMarshallField(stage, "FontSize"))
		initializerStatements.WriteString(deliverableviewshape.GongMarshallField(stage, "FontColor"))
		pointersInitializesStatements.WriteString(deliverableviewshape.GongMarshallField(stage, "ControlPointShapes"))
	}
	initRes = initializerStatements.String()
	ptrRes = pointersInitializesStatements.String()
	return
}
func (diagram *Diagram) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(diagram.GongMarshallField(stage, "Name"))
		initializerStatements.WriteString(diagram.GongMarshallField(stage, "ComputedPrefix"))
		initializerStatements.WriteString(diagram.GongMarshallField(stage, "IsExpanded"))
		initializerStatements.WriteString(diagram.GongMarshallField(stage, "IsChecked"))
		initializerStatements.WriteString(diagram.GongMarshallField(stage, "IsEditable_"))
		initializerStatements.WriteString(diagram.GongMarshallField(stage, "ShowPrefix"))
		initializerStatements.WriteString(diagram.GongMarshallField(stage, "DefaultBoxWidth"))
		initializerStatements.WriteString(diagram.GongMarshallField(stage, "DefaultBoxHeigth"))
		initializerStatements.WriteString(diagram.GongMarshallField(stage, "Width"))
		initializerStatements.WriteString(diagram.GongMarshallField(stage, "Height"))
		pointersInitializesStatements.WriteString(diagram.GongMarshallField(stage, "ConcernsWhoseRequirementsNodeIsExpanded"))
		initializerStatements.WriteString(diagram.GongMarshallField(stage, "IsRequirementsNodeExpanded"))
		initializerStatements.WriteString(diagram.GongMarshallField(stage, "IsViewsNodeExpanded"))
		initializerStatements.WriteString(diagram.GongMarshallField(stage, "IsConceptsNodeExpanded"))
		pointersInitializesStatements.WriteString(diagram.GongMarshallField(stage, "Deliverable_Shapes"))
		pointersInitializesStatements.WriteString(diagram.GongMarshallField(stage, "DeliverablesWhoseNodeIsExpanded"))
		pointersInitializesStatements.WriteString(diagram.GongMarshallField(stage, "DeliverablesWhoseViewsNodeIsExpanded"))
		pointersInitializesStatements.WriteString(diagram.GongMarshallField(stage, "DeliverablesWhoseConceptsNodeIsExpanded"))
		initializerStatements.WriteString(diagram.GongMarshallField(stage, "IsPBSNodeExpanded"))
		pointersInitializesStatements.WriteString(diagram.GongMarshallField(stage, "DeliverableComposition_Shapes"))
		initializerStatements.WriteString(diagram.GongMarshallField(stage, "IsConcernsNodeExpanded"))
		pointersInitializesStatements.WriteString(diagram.GongMarshallField(stage, "Concern_Shapes"))
		pointersInitializesStatements.WriteString(diagram.GongMarshallField(stage, "ConcernsWhoseNodeIsExpanded"))
		pointersInitializesStatements.WriteString(diagram.GongMarshallField(stage, "ConcernsWhoseInputNodeIsExpanded"))
		pointersInitializesStatements.WriteString(diagram.GongMarshallField(stage, "ConcernsWhoseStakeholderNodeIsExpanded"))
		pointersInitializesStatements.WriteString(diagram.GongMarshallField(stage, "ConcernssWhoseOutputNodeIsExpanded"))
		pointersInitializesStatements.WriteString(diagram.GongMarshallField(stage, "ConcernComposition_Shapes"))
		pointersInitializesStatements.WriteString(diagram.GongMarshallField(stage, "ConcernInputShapes"))
		pointersInitializesStatements.WriteString(diagram.GongMarshallField(stage, "ConcernOutputShapes"))
		pointersInitializesStatements.WriteString(diagram.GongMarshallField(stage, "Note_Shapes"))
		pointersInitializesStatements.WriteString(diagram.GongMarshallField(stage, "NotesWhoseNodeIsExpanded"))
		initializerStatements.WriteString(diagram.GongMarshallField(stage, "IsNotesNodeExpanded"))
		pointersInitializesStatements.WriteString(diagram.GongMarshallField(stage, "NoteDeliverableShapes"))
		pointersInitializesStatements.WriteString(diagram.GongMarshallField(stage, "NoteTaskShapes"))
		pointersInitializesStatements.WriteString(diagram.GongMarshallField(stage, "NoteResourceShapes"))
		pointersInitializesStatements.WriteString(diagram.GongMarshallField(stage, "Stakeholder_Shapes"))
		pointersInitializesStatements.WriteString(diagram.GongMarshallField(stage, "ResourcesWhoseNodeIsExpanded"))
		initializerStatements.WriteString(diagram.GongMarshallField(stage, "IsStakeholdersNodeExpanded"))
		pointersInitializesStatements.WriteString(diagram.GongMarshallField(stage, "ResourceComposition_Shapes"))
		pointersInitializesStatements.WriteString(diagram.GongMarshallField(stage, "StakeholderConcernShapes"))
		pointersInitializesStatements.WriteString(diagram.GongMarshallField(stage, "Requirement_Shapes"))
		pointersInitializesStatements.WriteString(diagram.GongMarshallField(stage, "RequirementsWhoseNodeIsExpanded"))
		pointersInitializesStatements.WriteString(diagram.GongMarshallField(stage, "RequirementsWhoseConcernsNodeIsExpanded"))
		pointersInitializesStatements.WriteString(diagram.GongMarshallField(stage, "RequirementsWhoseViewsNodeIsExpanded"))
		pointersInitializesStatements.WriteString(diagram.GongMarshallField(stage, "RequirementsWhoseConceptsNodeIsExpanded"))
		pointersInitializesStatements.WriteString(diagram.GongMarshallField(stage, "RequirementConcernShapes"))
		pointersInitializesStatements.WriteString(diagram.GongMarshallField(stage, "RequirementViewShapes"))
		pointersInitializesStatements.WriteString(diagram.GongMarshallField(stage, "RequirementConceptShapes"))
		pointersInitializesStatements.WriteString(diagram.GongMarshallField(stage, "View_Shapes"))
		pointersInitializesStatements.WriteString(diagram.GongMarshallField(stage, "ViewsWhoseNodeIsExpanded"))
		pointersInitializesStatements.WriteString(diagram.GongMarshallField(stage, "ViewsWhoseDeliverablesNodeIsExpanded"))
		pointersInitializesStatements.WriteString(diagram.GongMarshallField(stage, "ViewsWhoseConceptsNodeIsExpanded"))
		pointersInitializesStatements.WriteString(diagram.GongMarshallField(stage, "ViewsWhoseRequirementsNodeIsExpanded"))
		pointersInitializesStatements.WriteString(diagram.GongMarshallField(stage, "DeliverableViewShapes"))
		pointersInitializesStatements.WriteString(diagram.GongMarshallField(stage, "ViewConceptShapes"))
		pointersInitializesStatements.WriteString(diagram.GongMarshallField(stage, "Concept_Shapes"))
		pointersInitializesStatements.WriteString(diagram.GongMarshallField(stage, "ConceptsWhoseNodeIsExpanded"))
		pointersInitializesStatements.WriteString(diagram.GongMarshallField(stage, "ConceptsWhoseViewsNodeIsExpanded"))
		pointersInitializesStatements.WriteString(diagram.GongMarshallField(stage, "ConceptsWhoseDeliverablesNodeIsExpanded"))
		pointersInitializesStatements.WriteString(diagram.GongMarshallField(stage, "ConceptsWhoseRequirementsNodeIsExpanded"))
		pointersInitializesStatements.WriteString(diagram.GongMarshallField(stage, "DeliverableConceptShapes"))
		pointersInitializesStatements.WriteString(diagram.GongMarshallField(stage, "Diagram_Shapes"))
		initializerStatements.WriteString(diagram.GongMarshallField(stage, "IsDiagramsNodeExpanded"))
		pointersInitializesStatements.WriteString(diagram.GongMarshallField(stage, "DiagramsWhoseNodeIsExpanded"))
	}
	initRes = initializerStatements.String()
	ptrRes = pointersInitializesStatements.String()
	return
}
func (diagramshape *DiagramShape) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(diagramshape.GongMarshallField(stage, "Name"))
		pointersInitializesStatements.WriteString(diagramshape.GongMarshallField(stage, "Diagram"))
		initializerStatements.WriteString(diagramshape.GongMarshallField(stage, "IsExpanded"))
		initializerStatements.WriteString(diagramshape.GongMarshallField(stage, "X"))
		initializerStatements.WriteString(diagramshape.GongMarshallField(stage, "Y"))
		initializerStatements.WriteString(diagramshape.GongMarshallField(stage, "Width"))
		initializerStatements.WriteString(diagramshape.GongMarshallField(stage, "Height"))
		initializerStatements.WriteString(diagramshape.GongMarshallField(stage, "IsHidden"))
		initializerStatements.WriteString(diagramshape.GongMarshallField(stage, "IsPresentationOverridden"))
		initializerStatements.WriteString(diagramshape.GongMarshallField(stage, "Color"))
		initializerStatements.WriteString(diagramshape.GongMarshallField(stage, "FillOpacity"))
		initializerStatements.WriteString(diagramshape.GongMarshallField(stage, "Stroke"))
		initializerStatements.WriteString(diagramshape.GongMarshallField(stage, "StrokeWidth"))
		initializerStatements.WriteString(diagramshape.GongMarshallField(stage, "StrokeOpacity"))
		initializerStatements.WriteString(diagramshape.GongMarshallField(stage, "StrokeDashArray"))
		initializerStatements.WriteString(diagramshape.GongMarshallField(stage, "RX"))
		initializerStatements.WriteString(diagramshape.GongMarshallField(stage, "FontSize"))
		initializerStatements.WriteString(diagramshape.GongMarshallField(stage, "FontWeight"))
		initializerStatements.WriteString(diagramshape.GongMarshallField(stage, "FontColor"))
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
		initializerStatements.WriteString(library.GongMarshallField(stage, "IsRootLibrary"))
		initializerStatements.WriteString(library.GongMarshallField(stage, "ComputedPrefix"))
		initializerStatements.WriteString(library.GongMarshallField(stage, "IsExpanded"))
		pointersInitializesStatements.WriteString(library.GongMarshallField(stage, "RootDeliverables"))
		pointersInitializesStatements.WriteString(library.GongMarshallField(stage, "RootConcerns"))
		pointersInitializesStatements.WriteString(library.GongMarshallField(stage, "RootStakeholders"))
		pointersInitializesStatements.WriteString(library.GongMarshallField(stage, "RootRequirements"))
		pointersInitializesStatements.WriteString(library.GongMarshallField(stage, "RootViews"))
		pointersInitializesStatements.WriteString(library.GongMarshallField(stage, "RootConcepts"))
		pointersInitializesStatements.WriteString(library.GongMarshallField(stage, "AnalysisNeeds"))
		pointersInitializesStatements.WriteString(library.GongMarshallField(stage, "Notes"))
		pointersInitializesStatements.WriteString(library.GongMarshallField(stage, "Diagrams"))
		pointersInitializesStatements.WriteString(library.GongMarshallField(stage, "SubLibraries"))
		initializerStatements.WriteString(library.GongMarshallField(stage, "NbPixPerCharacter"))
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
		pointersInitializesStatements.WriteString(note.GongMarshallField(stage, "Deliverables"))
		pointersInitializesStatements.WriteString(note.GongMarshallField(stage, "Tasks"))
		pointersInitializesStatements.WriteString(note.GongMarshallField(stage, "Resources"))
	}
	initRes = initializerStatements.String()
	ptrRes = pointersInitializesStatements.String()
	return
}
func (notedeliverableshape *NoteDeliverableShape) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(notedeliverableshape.GongMarshallField(stage, "Name"))
		pointersInitializesStatements.WriteString(notedeliverableshape.GongMarshallField(stage, "Note"))
		pointersInitializesStatements.WriteString(notedeliverableshape.GongMarshallField(stage, "Deliverable"))
		initializerStatements.WriteString(notedeliverableshape.GongMarshallField(stage, "StartRatio"))
		initializerStatements.WriteString(notedeliverableshape.GongMarshallField(stage, "EndRatio"))
		initializerStatements.WriteString(notedeliverableshape.GongMarshallField(stage, "StartOrientation"))
		initializerStatements.WriteString(notedeliverableshape.GongMarshallField(stage, "EndOrientation"))
		initializerStatements.WriteString(notedeliverableshape.GongMarshallField(stage, "CornerOffsetRatio"))
		initializerStatements.WriteString(notedeliverableshape.GongMarshallField(stage, "IsHidden"))
		initializerStatements.WriteString(notedeliverableshape.GongMarshallField(stage, "IsPresentationOverridden"))
		initializerStatements.WriteString(notedeliverableshape.GongMarshallField(stage, "Stroke"))
		initializerStatements.WriteString(notedeliverableshape.GongMarshallField(stage, "StrokeWidth"))
		initializerStatements.WriteString(notedeliverableshape.GongMarshallField(stage, "StrokeOpacity"))
		initializerStatements.WriteString(notedeliverableshape.GongMarshallField(stage, "StrokeDashArray"))
		initializerStatements.WriteString(notedeliverableshape.GongMarshallField(stage, "FontSize"))
		initializerStatements.WriteString(notedeliverableshape.GongMarshallField(stage, "FontColor"))
		pointersInitializesStatements.WriteString(notedeliverableshape.GongMarshallField(stage, "ControlPointShapes"))
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
		initializerStatements.WriteString(noteshape.GongMarshallField(stage, "IsExpanded"))
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
func (notestakeholdershape *NoteStakeholderShape) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(notestakeholdershape.GongMarshallField(stage, "Name"))
		pointersInitializesStatements.WriteString(notestakeholdershape.GongMarshallField(stage, "Note"))
		pointersInitializesStatements.WriteString(notestakeholdershape.GongMarshallField(stage, "Stakeholder"))
		initializerStatements.WriteString(notestakeholdershape.GongMarshallField(stage, "StartRatio"))
		initializerStatements.WriteString(notestakeholdershape.GongMarshallField(stage, "EndRatio"))
		initializerStatements.WriteString(notestakeholdershape.GongMarshallField(stage, "StartOrientation"))
		initializerStatements.WriteString(notestakeholdershape.GongMarshallField(stage, "EndOrientation"))
		initializerStatements.WriteString(notestakeholdershape.GongMarshallField(stage, "CornerOffsetRatio"))
		initializerStatements.WriteString(notestakeholdershape.GongMarshallField(stage, "IsHidden"))
		initializerStatements.WriteString(notestakeholdershape.GongMarshallField(stage, "IsPresentationOverridden"))
		initializerStatements.WriteString(notestakeholdershape.GongMarshallField(stage, "Stroke"))
		initializerStatements.WriteString(notestakeholdershape.GongMarshallField(stage, "StrokeWidth"))
		initializerStatements.WriteString(notestakeholdershape.GongMarshallField(stage, "StrokeOpacity"))
		initializerStatements.WriteString(notestakeholdershape.GongMarshallField(stage, "StrokeDashArray"))
		initializerStatements.WriteString(notestakeholdershape.GongMarshallField(stage, "FontSize"))
		initializerStatements.WriteString(notestakeholdershape.GongMarshallField(stage, "FontColor"))
		pointersInitializesStatements.WriteString(notestakeholdershape.GongMarshallField(stage, "ControlPointShapes"))
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
		pointersInitializesStatements.WriteString(notetaskshape.GongMarshallField(stage, "ControlPointShapes"))
	}
	initRes = initializerStatements.String()
	ptrRes = pointersInitializesStatements.String()
	return
}
func (requirement *Requirement) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(requirement.GongMarshallField(stage, "Name"))
		initializerStatements.WriteString(requirement.GongMarshallField(stage, "ComputedPrefix"))
		initializerStatements.WriteString(requirement.GongMarshallField(stage, "IsExpanded"))
		pointersInitializesStatements.WriteString(requirement.GongMarshallField(stage, "SupportLevels"))
		pointersInitializesStatements.WriteString(requirement.GongMarshallField(stage, "Concerns"))
		pointersInitializesStatements.WriteString(requirement.GongMarshallField(stage, "Views"))
		pointersInitializesStatements.WriteString(requirement.GongMarshallField(stage, "Concepts"))
	}
	initRes = initializerStatements.String()
	ptrRes = pointersInitializesStatements.String()
	return
}
func (requirementconceptshape *RequirementConceptShape) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(requirementconceptshape.GongMarshallField(stage, "Name"))
		pointersInitializesStatements.WriteString(requirementconceptshape.GongMarshallField(stage, "Requirement"))
		pointersInitializesStatements.WriteString(requirementconceptshape.GongMarshallField(stage, "Concept"))
		initializerStatements.WriteString(requirementconceptshape.GongMarshallField(stage, "StartRatio"))
		initializerStatements.WriteString(requirementconceptshape.GongMarshallField(stage, "EndRatio"))
		initializerStatements.WriteString(requirementconceptshape.GongMarshallField(stage, "StartOrientation"))
		initializerStatements.WriteString(requirementconceptshape.GongMarshallField(stage, "EndOrientation"))
		initializerStatements.WriteString(requirementconceptshape.GongMarshallField(stage, "CornerOffsetRatio"))
		initializerStatements.WriteString(requirementconceptshape.GongMarshallField(stage, "IsHidden"))
		initializerStatements.WriteString(requirementconceptshape.GongMarshallField(stage, "IsPresentationOverridden"))
		initializerStatements.WriteString(requirementconceptshape.GongMarshallField(stage, "Stroke"))
		initializerStatements.WriteString(requirementconceptshape.GongMarshallField(stage, "StrokeWidth"))
		initializerStatements.WriteString(requirementconceptshape.GongMarshallField(stage, "StrokeOpacity"))
		initializerStatements.WriteString(requirementconceptshape.GongMarshallField(stage, "StrokeDashArray"))
		initializerStatements.WriteString(requirementconceptshape.GongMarshallField(stage, "FontSize"))
		initializerStatements.WriteString(requirementconceptshape.GongMarshallField(stage, "FontColor"))
		pointersInitializesStatements.WriteString(requirementconceptshape.GongMarshallField(stage, "ControlPointShapes"))
	}
	initRes = initializerStatements.String()
	ptrRes = pointersInitializesStatements.String()
	return
}
func (requirementconcernshape *RequirementConcernShape) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(requirementconcernshape.GongMarshallField(stage, "Name"))
		pointersInitializesStatements.WriteString(requirementconcernshape.GongMarshallField(stage, "Requirement"))
		pointersInitializesStatements.WriteString(requirementconcernshape.GongMarshallField(stage, "Concern"))
		initializerStatements.WriteString(requirementconcernshape.GongMarshallField(stage, "StartRatio"))
		initializerStatements.WriteString(requirementconcernshape.GongMarshallField(stage, "EndRatio"))
		initializerStatements.WriteString(requirementconcernshape.GongMarshallField(stage, "StartOrientation"))
		initializerStatements.WriteString(requirementconcernshape.GongMarshallField(stage, "EndOrientation"))
		initializerStatements.WriteString(requirementconcernshape.GongMarshallField(stage, "CornerOffsetRatio"))
		initializerStatements.WriteString(requirementconcernshape.GongMarshallField(stage, "IsHidden"))
		initializerStatements.WriteString(requirementconcernshape.GongMarshallField(stage, "IsPresentationOverridden"))
		initializerStatements.WriteString(requirementconcernshape.GongMarshallField(stage, "Stroke"))
		initializerStatements.WriteString(requirementconcernshape.GongMarshallField(stage, "StrokeWidth"))
		initializerStatements.WriteString(requirementconcernshape.GongMarshallField(stage, "StrokeOpacity"))
		initializerStatements.WriteString(requirementconcernshape.GongMarshallField(stage, "StrokeDashArray"))
		initializerStatements.WriteString(requirementconcernshape.GongMarshallField(stage, "FontSize"))
		initializerStatements.WriteString(requirementconcernshape.GongMarshallField(stage, "FontColor"))
		pointersInitializesStatements.WriteString(requirementconcernshape.GongMarshallField(stage, "ControlPointShapes"))
	}
	initRes = initializerStatements.String()
	ptrRes = pointersInitializesStatements.String()
	return
}
func (requirementshape *RequirementShape) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(requirementshape.GongMarshallField(stage, "Name"))
		pointersInitializesStatements.WriteString(requirementshape.GongMarshallField(stage, "Requirement"))
		initializerStatements.WriteString(requirementshape.GongMarshallField(stage, "IsExpanded"))
		initializerStatements.WriteString(requirementshape.GongMarshallField(stage, "X"))
		initializerStatements.WriteString(requirementshape.GongMarshallField(stage, "Y"))
		initializerStatements.WriteString(requirementshape.GongMarshallField(stage, "Width"))
		initializerStatements.WriteString(requirementshape.GongMarshallField(stage, "Height"))
		initializerStatements.WriteString(requirementshape.GongMarshallField(stage, "IsHidden"))
		initializerStatements.WriteString(requirementshape.GongMarshallField(stage, "IsPresentationOverridden"))
		initializerStatements.WriteString(requirementshape.GongMarshallField(stage, "Color"))
		initializerStatements.WriteString(requirementshape.GongMarshallField(stage, "FillOpacity"))
		initializerStatements.WriteString(requirementshape.GongMarshallField(stage, "Stroke"))
		initializerStatements.WriteString(requirementshape.GongMarshallField(stage, "StrokeWidth"))
		initializerStatements.WriteString(requirementshape.GongMarshallField(stage, "StrokeOpacity"))
		initializerStatements.WriteString(requirementshape.GongMarshallField(stage, "StrokeDashArray"))
		initializerStatements.WriteString(requirementshape.GongMarshallField(stage, "RX"))
		initializerStatements.WriteString(requirementshape.GongMarshallField(stage, "FontSize"))
		initializerStatements.WriteString(requirementshape.GongMarshallField(stage, "FontWeight"))
		initializerStatements.WriteString(requirementshape.GongMarshallField(stage, "FontColor"))
	}
	initRes = initializerStatements.String()
	ptrRes = pointersInitializesStatements.String()
	return
}
func (requirementviewshape *RequirementViewShape) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(requirementviewshape.GongMarshallField(stage, "Name"))
		pointersInitializesStatements.WriteString(requirementviewshape.GongMarshallField(stage, "Requirement"))
		pointersInitializesStatements.WriteString(requirementviewshape.GongMarshallField(stage, "View"))
		initializerStatements.WriteString(requirementviewshape.GongMarshallField(stage, "StartRatio"))
		initializerStatements.WriteString(requirementviewshape.GongMarshallField(stage, "EndRatio"))
		initializerStatements.WriteString(requirementviewshape.GongMarshallField(stage, "StartOrientation"))
		initializerStatements.WriteString(requirementviewshape.GongMarshallField(stage, "EndOrientation"))
		initializerStatements.WriteString(requirementviewshape.GongMarshallField(stage, "CornerOffsetRatio"))
		initializerStatements.WriteString(requirementviewshape.GongMarshallField(stage, "IsHidden"))
		initializerStatements.WriteString(requirementviewshape.GongMarshallField(stage, "IsPresentationOverridden"))
		initializerStatements.WriteString(requirementviewshape.GongMarshallField(stage, "Stroke"))
		initializerStatements.WriteString(requirementviewshape.GongMarshallField(stage, "StrokeWidth"))
		initializerStatements.WriteString(requirementviewshape.GongMarshallField(stage, "StrokeOpacity"))
		initializerStatements.WriteString(requirementviewshape.GongMarshallField(stage, "StrokeDashArray"))
		initializerStatements.WriteString(requirementviewshape.GongMarshallField(stage, "FontSize"))
		initializerStatements.WriteString(requirementviewshape.GongMarshallField(stage, "FontColor"))
		pointersInitializesStatements.WriteString(requirementviewshape.GongMarshallField(stage, "ControlPointShapes"))
	}
	initRes = initializerStatements.String()
	ptrRes = pointersInitializesStatements.String()
	return
}
func (stakeholder *Stakeholder) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(stakeholder.GongMarshallField(stage, "Name"))
		initializerStatements.WriteString(stakeholder.GongMarshallField(stage, "IDAirbus"))
		initializerStatements.WriteString(stakeholder.GongMarshallField(stage, "ComputedPrefix"))
		initializerStatements.WriteString(stakeholder.GongMarshallField(stage, "IsExpanded"))
		initializerStatements.WriteString(stakeholder.GongMarshallField(stage, "Description"))
		pointersInitializesStatements.WriteString(stakeholder.GongMarshallField(stage, "Concerns"))
		pointersInitializesStatements.WriteString(stakeholder.GongMarshallField(stage, "SubStakeholders"))
	}
	initRes = initializerStatements.String()
	ptrRes = pointersInitializesStatements.String()
	return
}
func (stakeholdercompositionshape *StakeholderCompositionShape) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(stakeholdercompositionshape.GongMarshallField(stage, "Name"))
		pointersInitializesStatements.WriteString(stakeholdercompositionshape.GongMarshallField(stage, "Stakeholder"))
		initializerStatements.WriteString(stakeholdercompositionshape.GongMarshallField(stage, "StartRatio"))
		initializerStatements.WriteString(stakeholdercompositionshape.GongMarshallField(stage, "EndRatio"))
		initializerStatements.WriteString(stakeholdercompositionshape.GongMarshallField(stage, "StartOrientation"))
		initializerStatements.WriteString(stakeholdercompositionshape.GongMarshallField(stage, "EndOrientation"))
		initializerStatements.WriteString(stakeholdercompositionshape.GongMarshallField(stage, "CornerOffsetRatio"))
		initializerStatements.WriteString(stakeholdercompositionshape.GongMarshallField(stage, "IsHidden"))
		initializerStatements.WriteString(stakeholdercompositionshape.GongMarshallField(stage, "IsPresentationOverridden"))
		initializerStatements.WriteString(stakeholdercompositionshape.GongMarshallField(stage, "Stroke"))
		initializerStatements.WriteString(stakeholdercompositionshape.GongMarshallField(stage, "StrokeWidth"))
		initializerStatements.WriteString(stakeholdercompositionshape.GongMarshallField(stage, "StrokeOpacity"))
		initializerStatements.WriteString(stakeholdercompositionshape.GongMarshallField(stage, "StrokeDashArray"))
		initializerStatements.WriteString(stakeholdercompositionshape.GongMarshallField(stage, "FontSize"))
		initializerStatements.WriteString(stakeholdercompositionshape.GongMarshallField(stage, "FontColor"))
		pointersInitializesStatements.WriteString(stakeholdercompositionshape.GongMarshallField(stage, "ControlPointShapes"))
	}
	initRes = initializerStatements.String()
	ptrRes = pointersInitializesStatements.String()
	return
}
func (stakeholderconcernshape *StakeholderConcernShape) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(stakeholderconcernshape.GongMarshallField(stage, "Name"))
		pointersInitializesStatements.WriteString(stakeholderconcernshape.GongMarshallField(stage, "Stakeholder"))
		pointersInitializesStatements.WriteString(stakeholderconcernshape.GongMarshallField(stage, "Concern"))
		initializerStatements.WriteString(stakeholderconcernshape.GongMarshallField(stage, "StartRatio"))
		initializerStatements.WriteString(stakeholderconcernshape.GongMarshallField(stage, "EndRatio"))
		initializerStatements.WriteString(stakeholderconcernshape.GongMarshallField(stage, "StartOrientation"))
		initializerStatements.WriteString(stakeholderconcernshape.GongMarshallField(stage, "EndOrientation"))
		initializerStatements.WriteString(stakeholderconcernshape.GongMarshallField(stage, "CornerOffsetRatio"))
		initializerStatements.WriteString(stakeholderconcernshape.GongMarshallField(stage, "IsHidden"))
		initializerStatements.WriteString(stakeholderconcernshape.GongMarshallField(stage, "IsPresentationOverridden"))
		initializerStatements.WriteString(stakeholderconcernshape.GongMarshallField(stage, "Stroke"))
		initializerStatements.WriteString(stakeholderconcernshape.GongMarshallField(stage, "StrokeWidth"))
		initializerStatements.WriteString(stakeholderconcernshape.GongMarshallField(stage, "StrokeOpacity"))
		initializerStatements.WriteString(stakeholderconcernshape.GongMarshallField(stage, "StrokeDashArray"))
		initializerStatements.WriteString(stakeholderconcernshape.GongMarshallField(stage, "FontSize"))
		initializerStatements.WriteString(stakeholderconcernshape.GongMarshallField(stage, "FontColor"))
		pointersInitializesStatements.WriteString(stakeholderconcernshape.GongMarshallField(stage, "ControlPointShapes"))
	}
	initRes = initializerStatements.String()
	ptrRes = pointersInitializesStatements.String()
	return
}
func (stakeholdershape *StakeholderShape) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(stakeholdershape.GongMarshallField(stage, "Name"))
		pointersInitializesStatements.WriteString(stakeholdershape.GongMarshallField(stage, "Stakeholder"))
		initializerStatements.WriteString(stakeholdershape.GongMarshallField(stage, "IsExpanded"))
		initializerStatements.WriteString(stakeholdershape.GongMarshallField(stage, "X"))
		initializerStatements.WriteString(stakeholdershape.GongMarshallField(stage, "Y"))
		initializerStatements.WriteString(stakeholdershape.GongMarshallField(stage, "Width"))
		initializerStatements.WriteString(stakeholdershape.GongMarshallField(stage, "Height"))
		initializerStatements.WriteString(stakeholdershape.GongMarshallField(stage, "IsHidden"))
		initializerStatements.WriteString(stakeholdershape.GongMarshallField(stage, "IsPresentationOverridden"))
		initializerStatements.WriteString(stakeholdershape.GongMarshallField(stage, "Color"))
		initializerStatements.WriteString(stakeholdershape.GongMarshallField(stage, "FillOpacity"))
		initializerStatements.WriteString(stakeholdershape.GongMarshallField(stage, "Stroke"))
		initializerStatements.WriteString(stakeholdershape.GongMarshallField(stage, "StrokeWidth"))
		initializerStatements.WriteString(stakeholdershape.GongMarshallField(stage, "StrokeOpacity"))
		initializerStatements.WriteString(stakeholdershape.GongMarshallField(stage, "StrokeDashArray"))
		initializerStatements.WriteString(stakeholdershape.GongMarshallField(stage, "RX"))
		initializerStatements.WriteString(stakeholdershape.GongMarshallField(stage, "FontSize"))
		initializerStatements.WriteString(stakeholdershape.GongMarshallField(stage, "FontWeight"))
		initializerStatements.WriteString(stakeholdershape.GongMarshallField(stage, "FontColor"))
	}
	initRes = initializerStatements.String()
	ptrRes = pointersInitializesStatements.String()
	return
}
func (supportlevel *SupportLevel) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(supportlevel.GongMarshallField(stage, "Name"))
		pointersInitializesStatements.WriteString(supportlevel.GongMarshallField(stage, "Tool"))
	}
	initRes = initializerStatements.String()
	ptrRes = pointersInitializesStatements.String()
	return
}
func (tool *Tool) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(tool.GongMarshallField(stage, "Name"))
	}
	initRes = initializerStatements.String()
	ptrRes = pointersInitializesStatements.String()
	return
}
func (view *View) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(view.GongMarshallField(stage, "Name"))
		initializerStatements.WriteString(view.GongMarshallField(stage, "ComputedPrefix"))
		initializerStatements.WriteString(view.GongMarshallField(stage, "IsExpanded"))
		initializerStatements.WriteString(view.GongMarshallField(stage, "Description"))
		pointersInitializesStatements.WriteString(view.GongMarshallField(stage, "Concepts"))
	}
	initRes = initializerStatements.String()
	ptrRes = pointersInitializesStatements.String()
	return
}
func (viewconceptshape *ViewConceptShape) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(viewconceptshape.GongMarshallField(stage, "Name"))
		pointersInitializesStatements.WriteString(viewconceptshape.GongMarshallField(stage, "View"))
		pointersInitializesStatements.WriteString(viewconceptshape.GongMarshallField(stage, "Concept"))
		initializerStatements.WriteString(viewconceptshape.GongMarshallField(stage, "StartRatio"))
		initializerStatements.WriteString(viewconceptshape.GongMarshallField(stage, "EndRatio"))
		initializerStatements.WriteString(viewconceptshape.GongMarshallField(stage, "StartOrientation"))
		initializerStatements.WriteString(viewconceptshape.GongMarshallField(stage, "EndOrientation"))
		initializerStatements.WriteString(viewconceptshape.GongMarshallField(stage, "CornerOffsetRatio"))
		initializerStatements.WriteString(viewconceptshape.GongMarshallField(stage, "IsHidden"))
		initializerStatements.WriteString(viewconceptshape.GongMarshallField(stage, "IsPresentationOverridden"))
		initializerStatements.WriteString(viewconceptshape.GongMarshallField(stage, "Stroke"))
		initializerStatements.WriteString(viewconceptshape.GongMarshallField(stage, "StrokeWidth"))
		initializerStatements.WriteString(viewconceptshape.GongMarshallField(stage, "StrokeOpacity"))
		initializerStatements.WriteString(viewconceptshape.GongMarshallField(stage, "StrokeDashArray"))
		initializerStatements.WriteString(viewconceptshape.GongMarshallField(stage, "FontSize"))
		initializerStatements.WriteString(viewconceptshape.GongMarshallField(stage, "FontColor"))
		pointersInitializesStatements.WriteString(viewconceptshape.GongMarshallField(stage, "ControlPointShapes"))
	}
	initRes = initializerStatements.String()
	ptrRes = pointersInitializesStatements.String()
	return
}
func (viewshape *ViewShape) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(viewshape.GongMarshallField(stage, "Name"))
		pointersInitializesStatements.WriteString(viewshape.GongMarshallField(stage, "View"))
		initializerStatements.WriteString(viewshape.GongMarshallField(stage, "IsExpanded"))
		initializerStatements.WriteString(viewshape.GongMarshallField(stage, "X"))
		initializerStatements.WriteString(viewshape.GongMarshallField(stage, "Y"))
		initializerStatements.WriteString(viewshape.GongMarshallField(stage, "Width"))
		initializerStatements.WriteString(viewshape.GongMarshallField(stage, "Height"))
		initializerStatements.WriteString(viewshape.GongMarshallField(stage, "IsHidden"))
		initializerStatements.WriteString(viewshape.GongMarshallField(stage, "IsPresentationOverridden"))
		initializerStatements.WriteString(viewshape.GongMarshallField(stage, "Color"))
		initializerStatements.WriteString(viewshape.GongMarshallField(stage, "FillOpacity"))
		initializerStatements.WriteString(viewshape.GongMarshallField(stage, "Stroke"))
		initializerStatements.WriteString(viewshape.GongMarshallField(stage, "StrokeWidth"))
		initializerStatements.WriteString(viewshape.GongMarshallField(stage, "StrokeOpacity"))
		initializerStatements.WriteString(viewshape.GongMarshallField(stage, "StrokeDashArray"))
		initializerStatements.WriteString(viewshape.GongMarshallField(stage, "RX"))
		initializerStatements.WriteString(viewshape.GongMarshallField(stage, "FontSize"))
		initializerStatements.WriteString(viewshape.GongMarshallField(stage, "FontWeight"))
		initializerStatements.WriteString(viewshape.GongMarshallField(stage, "FontColor"))
	}
	initRes = initializerStatements.String()
	ptrRes = pointersInitializesStatements.String()
	return
}
