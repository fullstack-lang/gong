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
	modelsPackageName := "github.com/fullstack-lang/gong/dsm/scenario/go/models"
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
		for _, actorstate := range __gong__sortStageSetInstances(stage.ActorStates, stage.ActorState_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			actorstateIdent := localAliasPrefix + actorstate.GongGetIdentifier(stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&models.ActorState{Name: %s}).Stage(stage)", actorstateIdent, __gong__toRawStringLiteral(actorstate.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", actorstateIdent, __gong__toRawStringLiteral(actorstate.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Description = %s", actorstateIdent, __gong__toRawStringLiteral(actorstate.Description)))
			values.WriteString(fmt.Sprintf("\n\t%s.IsWithProbaility = %t", actorstateIdent, actorstate.IsWithProbaility))
			if actorstate.Probability.ToCodeString() != "" {
				values.WriteString(fmt.Sprintf("\n\t%s.Probability = models.%s", actorstateIdent, actorstate.Probability.ToCodeString()))
			} else {
				values.WriteString(fmt.Sprintf("\n\t%s.Probability = \"\"", actorstateIdent))
			}
			values.WriteString(fmt.Sprintf("\n\t%s.ComputedPrefix = %s", actorstateIdent, __gong__toRawStringLiteral(actorstate.ComputedPrefix)))
			values.WriteString(fmt.Sprintf("\n\t%s.IsExpanded = %t", actorstateIdent, actorstate.IsExpanded))
		}
	}
	if stage != nil {
		for _, actorstateshape := range __gong__sortStageSetInstances(stage.ActorStateShapes, stage.ActorStateShape_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			actorstateshapeIdent := localAliasPrefix + actorstateshape.GongGetIdentifier(stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&models.ActorStateShape{Name: %s}).Stage(stage)", actorstateshapeIdent, __gong__toRawStringLiteral(actorstateshape.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", actorstateshapeIdent, __gong__toRawStringLiteral(actorstateshape.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.X = %f", actorstateshapeIdent, actorstateshape.X))
			values.WriteString(fmt.Sprintf("\n\t%s.Y = %f", actorstateshapeIdent, actorstateshape.Y))
			values.WriteString(fmt.Sprintf("\n\t%s.Width = %f", actorstateshapeIdent, actorstateshape.Width))
			values.WriteString(fmt.Sprintf("\n\t%s.Height = %f", actorstateshapeIdent, actorstateshape.Height))
			values.WriteString(fmt.Sprintf("\n\t%s.IsHidden = %t", actorstateshapeIdent, actorstateshape.IsHidden))
			values.WriteString(fmt.Sprintf("\n\t%s.IsPresentationOverridden = %t", actorstateshapeIdent, actorstateshape.IsPresentationOverridden))
			values.WriteString(fmt.Sprintf("\n\t%s.Color = %s", actorstateshapeIdent, __gong__toRawStringLiteral(actorstateshape.Color)))
			values.WriteString(fmt.Sprintf("\n\t%s.FillOpacity = %f", actorstateshapeIdent, actorstateshape.FillOpacity))
			values.WriteString(fmt.Sprintf("\n\t%s.Stroke = %s", actorstateshapeIdent, __gong__toRawStringLiteral(actorstateshape.Stroke)))
			values.WriteString(fmt.Sprintf("\n\t%s.StrokeWidth = %f", actorstateshapeIdent, actorstateshape.StrokeWidth))
			values.WriteString(fmt.Sprintf("\n\t%s.StrokeOpacity = %f", actorstateshapeIdent, actorstateshape.StrokeOpacity))
			values.WriteString(fmt.Sprintf("\n\t%s.StrokeDashArray = %s", actorstateshapeIdent, __gong__toRawStringLiteral(actorstateshape.StrokeDashArray)))
			values.WriteString(fmt.Sprintf("\n\t%s.RX = %f", actorstateshapeIdent, actorstateshape.RX))
			values.WriteString(fmt.Sprintf("\n\t%s.FontSize = %s", actorstateshapeIdent, __gong__toRawStringLiteral(actorstateshape.FontSize)))
			values.WriteString(fmt.Sprintf("\n\t%s.FontWeight = %s", actorstateshapeIdent, __gong__toRawStringLiteral(actorstateshape.FontWeight)))
			values.WriteString(fmt.Sprintf("\n\t%s.FontColor = %s", actorstateshapeIdent, __gong__toRawStringLiteral(actorstateshape.FontColor)))
			if actorstateshape.ActorState != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := localAliasPrefix + actorstateshape.ActorState.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.ActorState = %s", actorstateshapeIdent, targetIdent))
			}
		}
	}
	if stage != nil {
		for _, actorstatetransition := range __gong__sortStageSetInstances(stage.ActorStateTransitions, stage.ActorStateTransition_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			actorstatetransitionIdent := localAliasPrefix + actorstatetransition.GongGetIdentifier(stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&models.ActorStateTransition{Name: %s}).Stage(stage)", actorstatetransitionIdent, __gong__toRawStringLiteral(actorstatetransition.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", actorstatetransitionIdent, __gong__toRawStringLiteral(actorstatetransition.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.ComputedPrefix = %s", actorstatetransitionIdent, __gong__toRawStringLiteral(actorstatetransition.ComputedPrefix)))
			values.WriteString(fmt.Sprintf("\n\t%s.IsExpanded = %t", actorstatetransitionIdent, actorstatetransition.IsExpanded))
			if actorstatetransition.StartState != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := localAliasPrefix + actorstatetransition.StartState.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.StartState = %s", actorstatetransitionIdent, targetIdent))
			}
			if actorstatetransition.EndState != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := localAliasPrefix + actorstatetransition.EndState.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.EndState = %s", actorstatetransitionIdent, targetIdent))
			}
			for _, elem := range actorstatetransition.Justifications {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := localAliasPrefix + elem.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Justifications = append(%s.Justifications, %s)", actorstatetransitionIdent, actorstatetransitionIdent, targetIdent))
			}
		}
	}
	if stage != nil {
		for _, actorstatetransitionshape := range __gong__sortStageSetInstances(stage.ActorStateTransitionShapes, stage.ActorStateTransitionShape_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			actorstatetransitionshapeIdent := localAliasPrefix + actorstatetransitionshape.GongGetIdentifier(stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&models.ActorStateTransitionShape{Name: %s}).Stage(stage)", actorstatetransitionshapeIdent, __gong__toRawStringLiteral(actorstatetransitionshape.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", actorstatetransitionshapeIdent, __gong__toRawStringLiteral(actorstatetransitionshape.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.X = %f", actorstatetransitionshapeIdent, actorstatetransitionshape.X))
			values.WriteString(fmt.Sprintf("\n\t%s.Y = %f", actorstatetransitionshapeIdent, actorstatetransitionshape.Y))
			values.WriteString(fmt.Sprintf("\n\t%s.Width = %f", actorstatetransitionshapeIdent, actorstatetransitionshape.Width))
			values.WriteString(fmt.Sprintf("\n\t%s.Height = %f", actorstatetransitionshapeIdent, actorstatetransitionshape.Height))
			values.WriteString(fmt.Sprintf("\n\t%s.IsHidden = %t", actorstatetransitionshapeIdent, actorstatetransitionshape.IsHidden))
			values.WriteString(fmt.Sprintf("\n\t%s.IsPresentationOverridden = %t", actorstatetransitionshapeIdent, actorstatetransitionshape.IsPresentationOverridden))
			values.WriteString(fmt.Sprintf("\n\t%s.Color = %s", actorstatetransitionshapeIdent, __gong__toRawStringLiteral(actorstatetransitionshape.Color)))
			values.WriteString(fmt.Sprintf("\n\t%s.FillOpacity = %f", actorstatetransitionshapeIdent, actorstatetransitionshape.FillOpacity))
			values.WriteString(fmt.Sprintf("\n\t%s.Stroke = %s", actorstatetransitionshapeIdent, __gong__toRawStringLiteral(actorstatetransitionshape.Stroke)))
			values.WriteString(fmt.Sprintf("\n\t%s.StrokeWidth = %f", actorstatetransitionshapeIdent, actorstatetransitionshape.StrokeWidth))
			values.WriteString(fmt.Sprintf("\n\t%s.StrokeOpacity = %f", actorstatetransitionshapeIdent, actorstatetransitionshape.StrokeOpacity))
			values.WriteString(fmt.Sprintf("\n\t%s.StrokeDashArray = %s", actorstatetransitionshapeIdent, __gong__toRawStringLiteral(actorstatetransitionshape.StrokeDashArray)))
			values.WriteString(fmt.Sprintf("\n\t%s.RX = %f", actorstatetransitionshapeIdent, actorstatetransitionshape.RX))
			values.WriteString(fmt.Sprintf("\n\t%s.FontSize = %s", actorstatetransitionshapeIdent, __gong__toRawStringLiteral(actorstatetransitionshape.FontSize)))
			values.WriteString(fmt.Sprintf("\n\t%s.FontWeight = %s", actorstatetransitionshapeIdent, __gong__toRawStringLiteral(actorstatetransitionshape.FontWeight)))
			values.WriteString(fmt.Sprintf("\n\t%s.FontColor = %s", actorstatetransitionshapeIdent, __gong__toRawStringLiteral(actorstatetransitionshape.FontColor)))
			if actorstatetransitionshape.ActorStateTransition != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := localAliasPrefix + actorstatetransitionshape.ActorStateTransition.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.ActorStateTransition = %s", actorstatetransitionshapeIdent, targetIdent))
			}
			if actorstatetransitionshape.Start != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := localAliasPrefix + actorstatetransitionshape.Start.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Start = %s", actorstatetransitionshapeIdent, targetIdent))
			}
			if actorstatetransitionshape.End != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := localAliasPrefix + actorstatetransitionshape.End.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.End = %s", actorstatetransitionshapeIdent, targetIdent))
			}
			for _, elem := range actorstatetransitionshape.ControlPointShapes {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := localAliasPrefix + elem.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.ControlPointShapes = append(%s.ControlPointShapes, %s)", actorstatetransitionshapeIdent, actorstatetransitionshapeIdent, targetIdent))
			}
		}
	}
	if stage != nil {
		for _, analysis := range __gong__sortStageSetInstances(stage.Analysiss, stage.Analysis_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			analysisIdent := localAliasPrefix + analysis.GongGetIdentifier(stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&models.Analysis{Name: %s}).Stage(stage)", analysisIdent, __gong__toRawStringLiteral(analysis.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", analysisIdent, __gong__toRawStringLiteral(analysis.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Description = %s", analysisIdent, __gong__toRawStringLiteral(analysis.Description)))
			values.WriteString(fmt.Sprintf("\n\t%s.IsScenariosNodeExpanded = %t", analysisIdent, analysis.IsScenariosNodeExpanded))
			values.WriteString(fmt.Sprintf("\n\t%s.IsGroupUseNodeExpanded = %t", analysisIdent, analysis.IsGroupUseNodeExpanded))
			values.WriteString(fmt.Sprintf("\n\t%s.IsGeoObjectUseNodeExpanded = %t", analysisIdent, analysis.IsGeoObjectUseNodeExpanded))
			values.WriteString(fmt.Sprintf("\n\t%s.IsMapUseNodeExpanded = %t", analysisIdent, analysis.IsMapUseNodeExpanded))
			values.WriteString(fmt.Sprintf("\n\t%s.ComputedPrefix = %s", analysisIdent, __gong__toRawStringLiteral(analysis.ComputedPrefix)))
			values.WriteString(fmt.Sprintf("\n\t%s.IsExpanded = %t", analysisIdent, analysis.IsExpanded))
			for _, elem := range analysis.Scenarios {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := localAliasPrefix + elem.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Scenarios = append(%s.Scenarios, %s)", analysisIdent, analysisIdent, targetIdent))
			}
			for _, elem := range analysis.GroupUse {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := localAliasPrefix + elem.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.GroupUse = append(%s.GroupUse, %s)", analysisIdent, analysisIdent, targetIdent))
			}
			for _, elem := range analysis.GeoObjectUse {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := localAliasPrefix + elem.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.GeoObjectUse = append(%s.GeoObjectUse, %s)", analysisIdent, analysisIdent, targetIdent))
			}
			for _, elem := range analysis.MapUse {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := localAliasPrefix + elem.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.MapUse = append(%s.MapUse, %s)", analysisIdent, analysisIdent, targetIdent))
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
			controlpointshapeIdent := localAliasPrefix + controlpointshape.GongGetIdentifier(stage)
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
			values.WriteString(fmt.Sprintf("\n\t%s.ComputedPrefix = %s", diagramIdent, __gong__toRawStringLiteral(diagram.ComputedPrefix)))
			values.WriteString(fmt.Sprintf("\n\t%s.IsExpanded = %t", diagramIdent, diagram.IsExpanded))
			values.WriteString(fmt.Sprintf("\n\t%s.IsChecked = %t", diagramIdent, diagram.IsChecked))
			values.WriteString(fmt.Sprintf("\n\t%s.IsShowPrefix = %t", diagramIdent, diagram.IsShowPrefix))
			values.WriteString(fmt.Sprintf("\n\t%s.Description = %s", diagramIdent, __gong__toRawStringLiteral(diagram.Description)))
			values.WriteString(fmt.Sprintf("\n\t%s.IsEvolutionDirectionsNodeExpanded = %t", diagramIdent, diagram.IsEvolutionDirectionsNodeExpanded))
			values.WriteString(fmt.Sprintf("\n\t%s.IsActorStatesNodeExpanded = %t", diagramIdent, diagram.IsActorStatesNodeExpanded))
			values.WriteString(fmt.Sprintf("\n\t%s.IsParametersNodeExpanded = %t", diagramIdent, diagram.IsParametersNodeExpanded))
			values.WriteString(fmt.Sprintf("\n\t%s.IsParametersAggregatesNodeExpanded = %t", diagramIdent, diagram.IsParametersAggregatesNodeExpanded))
			values.WriteString(fmt.Sprintf("\n\t%s.IsActorStateTransitionsNodeExpanded = %t", diagramIdent, diagram.IsActorStateTransitionsNodeExpanded))
			values.WriteString(fmt.Sprintf("\n\t%s.AxisOrign_X = %f", diagramIdent, diagram.AxisOrign_X))
			values.WriteString(fmt.Sprintf("\n\t%s.AxisOrign_Y = %f", diagramIdent, diagram.AxisOrign_Y))
			values.WriteString(fmt.Sprintf("\n\t%s.VerticalAxis_Top_Y = %f", diagramIdent, diagram.VerticalAxis_Top_Y))
			values.WriteString(fmt.Sprintf("\n\t%s.VerticalAxis_Bottom_Y = %f", diagramIdent, diagram.VerticalAxis_Bottom_Y))
			values.WriteString(fmt.Sprintf("\n\t%s.VerticalAxis_StrokeWidth = %f", diagramIdent, diagram.VerticalAxis_StrokeWidth))
			values.WriteString(fmt.Sprintf("\n\t%s.HorizontalAxis_Right_X = %f", diagramIdent, diagram.HorizontalAxis_Right_X))
			values.WriteString(fmt.Sprintf("\n\t%s.Start, _ = time.Parse(\"2006-01-02 15:04:05.999999999 -0700 MST\", \"%s\")", diagramIdent, diagram.Start.String()))
			values.WriteString(fmt.Sprintf("\n\t%s.End, _ = time.Parse(\"2006-01-02 15:04:05.999999999 -0700 MST\", \"%s\")", diagramIdent, diagram.End.String()))
			values.WriteString(fmt.Sprintf("\n\t%s.NumberOfYearsBetweenTicks = %d", diagramIdent, diagram.NumberOfYearsBetweenTicks))
			values.WriteString(fmt.Sprintf("\n\t%s.IsInDrawMode = %t", diagramIdent, diagram.IsInDrawMode))
			for _, elem := range diagram.EvolutionDirectionShapes {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := localAliasPrefix + elem.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.EvolutionDirectionShapes = append(%s.EvolutionDirectionShapes, %s)", diagramIdent, diagramIdent, targetIdent))
			}
			for _, elem := range diagram.EvolutionDirectionsWhoseNodeIsExpanded {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := localAliasPrefix + elem.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.EvolutionDirectionsWhoseNodeIsExpanded = append(%s.EvolutionDirectionsWhoseNodeIsExpanded, %s)", diagramIdent, diagramIdent, targetIdent))
			}
			for _, elem := range diagram.ActorStateShapes {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := localAliasPrefix + elem.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.ActorStateShapes = append(%s.ActorStateShapes, %s)", diagramIdent, diagramIdent, targetIdent))
			}
			for _, elem := range diagram.ActorStatesWhoseNodeIsExpanded {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := localAliasPrefix + elem.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.ActorStatesWhoseNodeIsExpanded = append(%s.ActorStatesWhoseNodeIsExpanded, %s)", diagramIdent, diagramIdent, targetIdent))
			}
			for _, elem := range diagram.ParameterShapes {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := localAliasPrefix + elem.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.ParameterShapes = append(%s.ParameterShapes, %s)", diagramIdent, diagramIdent, targetIdent))
			}
			for _, elem := range diagram.ParametersWhoseNodeIsExpanded {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := localAliasPrefix + elem.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.ParametersWhoseNodeIsExpanded = append(%s.ParametersWhoseNodeIsExpanded, %s)", diagramIdent, diagramIdent, targetIdent))
			}
			for _, elem := range diagram.ScenarioParameterShapes {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := localAliasPrefix + elem.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.ScenarioParameterShapes = append(%s.ScenarioParameterShapes, %s)", diagramIdent, diagramIdent, targetIdent))
			}
			for _, elem := range diagram.ParametersAggregatesWhoseNodeIsExpanded {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := localAliasPrefix + elem.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.ParametersAggregatesWhoseNodeIsExpanded = append(%s.ParametersAggregatesWhoseNodeIsExpanded, %s)", diagramIdent, diagramIdent, targetIdent))
			}
			for _, elem := range diagram.ActorStateTransitionShapes {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := localAliasPrefix + elem.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.ActorStateTransitionShapes = append(%s.ActorStateTransitionShapes, %s)", diagramIdent, diagramIdent, targetIdent))
			}
			for _, elem := range diagram.ActorStateTransitionsWhoseNodeIsExpanded {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := localAliasPrefix + elem.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.ActorStateTransitionsWhoseNodeIsExpanded = append(%s.ActorStateTransitionsWhoseNodeIsExpanded, %s)", diagramIdent, diagramIdent, targetIdent))
			}
		}
	}
	if stage != nil {
		for _, document := range __gong__sortStageSetInstances(stage.Documents, stage.Document_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			documentIdent := localAliasPrefix + document.GongGetIdentifier(stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&models.Document{Name: %s}).Stage(stage)", documentIdent, __gong__toRawStringLiteral(document.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", documentIdent, __gong__toRawStringLiteral(document.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.ComputedPrefix = %s", documentIdent, __gong__toRawStringLiteral(document.ComputedPrefix)))
			values.WriteString(fmt.Sprintf("\n\t%s.IsExpanded = %t", documentIdent, document.IsExpanded))
			for _, elem := range document.GeoObjectUse {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := localAliasPrefix + elem.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.GeoObjectUse = append(%s.GeoObjectUse, %s)", documentIdent, documentIdent, targetIdent))
			}
		}
	}
	if stage != nil {
		for _, documentuse := range __gong__sortStageSetInstances(stage.DocumentUses, stage.DocumentUse_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			documentuseIdent := localAliasPrefix + documentuse.GongGetIdentifier(stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&models.DocumentUse{Name: %s}).Stage(stage)", documentuseIdent, __gong__toRawStringLiteral(documentuse.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", documentuseIdent, __gong__toRawStringLiteral(documentuse.Name)))
			if documentuse.Document != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := localAliasPrefix + documentuse.Document.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Document = %s", documentuseIdent, targetIdent))
			}
		}
	}
	if stage != nil {
		for _, evolutiondirection := range __gong__sortStageSetInstances(stage.EvolutionDirections, stage.EvolutionDirection_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			evolutiondirectionIdent := localAliasPrefix + evolutiondirection.GongGetIdentifier(stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&models.EvolutionDirection{Name: %s}).Stage(stage)", evolutiondirectionIdent, __gong__toRawStringLiteral(evolutiondirection.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", evolutiondirectionIdent, __gong__toRawStringLiteral(evolutiondirection.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Description = %s", evolutiondirectionIdent, __gong__toRawStringLiteral(evolutiondirection.Description)))
			values.WriteString(fmt.Sprintf("\n\t%s.ComputedPrefix = %s", evolutiondirectionIdent, __gong__toRawStringLiteral(evolutiondirection.ComputedPrefix)))
			values.WriteString(fmt.Sprintf("\n\t%s.IsExpanded = %t", evolutiondirectionIdent, evolutiondirection.IsExpanded))
		}
	}
	if stage != nil {
		for _, evolutiondirectionshape := range __gong__sortStageSetInstances(stage.EvolutionDirectionShapes, stage.EvolutionDirectionShape_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			evolutiondirectionshapeIdent := localAliasPrefix + evolutiondirectionshape.GongGetIdentifier(stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&models.EvolutionDirectionShape{Name: %s}).Stage(stage)", evolutiondirectionshapeIdent, __gong__toRawStringLiteral(evolutiondirectionshape.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", evolutiondirectionshapeIdent, __gong__toRawStringLiteral(evolutiondirectionshape.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.X = %f", evolutiondirectionshapeIdent, evolutiondirectionshape.X))
			values.WriteString(fmt.Sprintf("\n\t%s.Y = %f", evolutiondirectionshapeIdent, evolutiondirectionshape.Y))
			values.WriteString(fmt.Sprintf("\n\t%s.Width = %f", evolutiondirectionshapeIdent, evolutiondirectionshape.Width))
			values.WriteString(fmt.Sprintf("\n\t%s.Height = %f", evolutiondirectionshapeIdent, evolutiondirectionshape.Height))
			values.WriteString(fmt.Sprintf("\n\t%s.IsHidden = %t", evolutiondirectionshapeIdent, evolutiondirectionshape.IsHidden))
			values.WriteString(fmt.Sprintf("\n\t%s.IsPresentationOverridden = %t", evolutiondirectionshapeIdent, evolutiondirectionshape.IsPresentationOverridden))
			values.WriteString(fmt.Sprintf("\n\t%s.Color = %s", evolutiondirectionshapeIdent, __gong__toRawStringLiteral(evolutiondirectionshape.Color)))
			values.WriteString(fmt.Sprintf("\n\t%s.FillOpacity = %f", evolutiondirectionshapeIdent, evolutiondirectionshape.FillOpacity))
			values.WriteString(fmt.Sprintf("\n\t%s.Stroke = %s", evolutiondirectionshapeIdent, __gong__toRawStringLiteral(evolutiondirectionshape.Stroke)))
			values.WriteString(fmt.Sprintf("\n\t%s.StrokeWidth = %f", evolutiondirectionshapeIdent, evolutiondirectionshape.StrokeWidth))
			values.WriteString(fmt.Sprintf("\n\t%s.StrokeOpacity = %f", evolutiondirectionshapeIdent, evolutiondirectionshape.StrokeOpacity))
			values.WriteString(fmt.Sprintf("\n\t%s.StrokeDashArray = %s", evolutiondirectionshapeIdent, __gong__toRawStringLiteral(evolutiondirectionshape.StrokeDashArray)))
			values.WriteString(fmt.Sprintf("\n\t%s.RX = %f", evolutiondirectionshapeIdent, evolutiondirectionshape.RX))
			values.WriteString(fmt.Sprintf("\n\t%s.FontSize = %s", evolutiondirectionshapeIdent, __gong__toRawStringLiteral(evolutiondirectionshape.FontSize)))
			values.WriteString(fmt.Sprintf("\n\t%s.FontWeight = %s", evolutiondirectionshapeIdent, __gong__toRawStringLiteral(evolutiondirectionshape.FontWeight)))
			values.WriteString(fmt.Sprintf("\n\t%s.FontColor = %s", evolutiondirectionshapeIdent, __gong__toRawStringLiteral(evolutiondirectionshape.FontColor)))
			if evolutiondirectionshape.EvolutionDirection != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := localAliasPrefix + evolutiondirectionshape.EvolutionDirection.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.EvolutionDirection = %s", evolutiondirectionshapeIdent, targetIdent))
			}
		}
	}
	if stage != nil {
		for _, foo := range __gong__sortStageSetInstances(stage.Foos, stage.Foo_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			fooIdent := localAliasPrefix + foo.GongGetIdentifier(stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&models.Foo{Name: %s}).Stage(stage)", fooIdent, __gong__toRawStringLiteral(foo.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", fooIdent, __gong__toRawStringLiteral(foo.Name)))
		}
	}
	if stage != nil {
		for _, geoobject := range __gong__sortStageSetInstances(stage.GeoObjects, stage.GeoObject_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			geoobjectIdent := localAliasPrefix + geoobject.GongGetIdentifier(stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&models.GeoObject{Name: %s}).Stage(stage)", geoobjectIdent, __gong__toRawStringLiteral(geoobject.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", geoobjectIdent, __gong__toRawStringLiteral(geoobject.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.ComputedPrefix = %s", geoobjectIdent, __gong__toRawStringLiteral(geoobject.ComputedPrefix)))
			values.WriteString(fmt.Sprintf("\n\t%s.IsExpanded = %t", geoobjectIdent, geoobject.IsExpanded))
		}
	}
	if stage != nil {
		for _, geoobjectuse := range __gong__sortStageSetInstances(stage.GeoObjectUses, stage.GeoObjectUse_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			geoobjectuseIdent := localAliasPrefix + geoobjectuse.GongGetIdentifier(stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&models.GeoObjectUse{Name: %s}).Stage(stage)", geoobjectuseIdent, __gong__toRawStringLiteral(geoobjectuse.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", geoobjectuseIdent, __gong__toRawStringLiteral(geoobjectuse.Name)))
			if geoobjectuse.GeoObject != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := localAliasPrefix + geoobjectuse.GeoObject.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.GeoObject = %s", geoobjectuseIdent, targetIdent))
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
			groupIdent := localAliasPrefix + group.GongGetIdentifier(stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&models.Group{Name: %s}).Stage(stage)", groupIdent, __gong__toRawStringLiteral(group.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", groupIdent, __gong__toRawStringLiteral(group.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.ComputedPrefix = %s", groupIdent, __gong__toRawStringLiteral(group.ComputedPrefix)))
			values.WriteString(fmt.Sprintf("\n\t%s.IsExpanded = %t", groupIdent, group.IsExpanded))
			for _, elem := range group.UserUse {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := localAliasPrefix + elem.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.UserUse = append(%s.UserUse, %s)", groupIdent, groupIdent, targetIdent))
			}
		}
	}
	if stage != nil {
		for _, groupuse := range __gong__sortStageSetInstances(stage.GroupUses, stage.GroupUse_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			groupuseIdent := localAliasPrefix + groupuse.GongGetIdentifier(stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&models.GroupUse{Name: %s}).Stage(stage)", groupuseIdent, __gong__toRawStringLiteral(groupuse.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", groupuseIdent, __gong__toRawStringLiteral(groupuse.Name)))
			if groupuse.Group != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := localAliasPrefix + groupuse.Group.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Group = %s", groupuseIdent, targetIdent))
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
			values.WriteString(fmt.Sprintf("\n\t%s.IsAnalysesNodeExpanded = %t", libraryIdent, library.IsAnalysesNodeExpanded))
			values.WriteString(fmt.Sprintf("\n\t%s.IsSubLibrariesNodeExpanded = %t", libraryIdent, library.IsSubLibrariesNodeExpanded))
			values.WriteString(fmt.Sprintf("\n\t%s.NbPixPerCharacter = %f", libraryIdent, library.NbPixPerCharacter))
			values.WriteString(fmt.Sprintf("\n\t%s.LogoSVGFile = %s", libraryIdent, __gong__toRawStringLiteral(library.LogoSVGFile)))
			values.WriteString(fmt.Sprintf("\n\t%s.IsExpandedTmp = %t", libraryIdent, library.IsExpandedTmp))
			for _, elem := range library.Analyses {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := localAliasPrefix + elem.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Analyses = append(%s.Analyses, %s)", libraryIdent, libraryIdent, targetIdent))
			}
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
		}
	}
	if stage != nil {
		for _, mapobject := range __gong__sortStageSetInstances(stage.MapObjects, stage.MapObject_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			mapobjectIdent := localAliasPrefix + mapobject.GongGetIdentifier(stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&models.MapObject{Name: %s}).Stage(stage)", mapobjectIdent, __gong__toRawStringLiteral(mapobject.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", mapobjectIdent, __gong__toRawStringLiteral(mapobject.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.ComputedPrefix = %s", mapobjectIdent, __gong__toRawStringLiteral(mapobject.ComputedPrefix)))
			values.WriteString(fmt.Sprintf("\n\t%s.IsExpanded = %t", mapobjectIdent, mapobject.IsExpanded))
		}
	}
	if stage != nil {
		for _, mapobjectuse := range __gong__sortStageSetInstances(stage.MapObjectUses, stage.MapObjectUse_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			mapobjectuseIdent := localAliasPrefix + mapobjectuse.GongGetIdentifier(stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&models.MapObjectUse{Name: %s}).Stage(stage)", mapobjectuseIdent, __gong__toRawStringLiteral(mapobjectuse.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", mapobjectuseIdent, __gong__toRawStringLiteral(mapobjectuse.Name)))
			if mapobjectuse.Map != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := localAliasPrefix + mapobjectuse.Map.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Map = %s", mapobjectuseIdent, targetIdent))
			}
		}
	}
	if stage != nil {
		for _, parameter := range __gong__sortStageSetInstances(stage.Parameters, stage.Parameter_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			parameterIdent := localAliasPrefix + parameter.GongGetIdentifier(stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&models.Parameter{Name: %s}).Stage(stage)", parameterIdent, __gong__toRawStringLiteral(parameter.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", parameterIdent, __gong__toRawStringLiteral(parameter.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Description = %s", parameterIdent, __gong__toRawStringLiteral(parameter.Description)))
			values.WriteString(fmt.Sprintf("\n\t%s.IsResponse = %t", parameterIdent, parameter.IsResponse))
			values.WriteString(fmt.Sprintf("\n\t%s.Start, _ = time.Parse(\"2006-01-02 15:04:05.999999999 -0700 MST\", \"%s\")", parameterIdent, parameter.Start.String()))
			values.WriteString(fmt.Sprintf("\n\t%s.End, _ = time.Parse(\"2006-01-02 15:04:05.999999999 -0700 MST\", \"%s\")", parameterIdent, parameter.End.String()))
			values.WriteString(fmt.Sprintf("\n\t%s.Force = %f", parameterIdent, parameter.Force))
			values.WriteString(fmt.Sprintf("\n\t%s.Tag = %s", parameterIdent, __gong__toRawStringLiteral(parameter.Tag)))
			values.WriteString(fmt.Sprintf("\n\t%s.ComputedPrefix = %s", parameterIdent, __gong__toRawStringLiteral(parameter.ComputedPrefix)))
			values.WriteString(fmt.Sprintf("\n\t%s.IsExpanded = %t", parameterIdent, parameter.IsExpanded))
			for _, elem := range parameter.GroupUse {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := localAliasPrefix + elem.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.GroupUse = append(%s.GroupUse, %s)", parameterIdent, parameterIdent, targetIdent))
			}
			for _, elem := range parameter.DocumentUse {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := localAliasPrefix + elem.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.DocumentUse = append(%s.DocumentUse, %s)", parameterIdent, parameterIdent, targetIdent))
			}
			for _, elem := range parameter.GeoObjectUse {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := localAliasPrefix + elem.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.GeoObjectUse = append(%s.GeoObjectUse, %s)", parameterIdent, parameterIdent, targetIdent))
			}
		}
	}
	if stage != nil {
		for _, parametercategory := range __gong__sortStageSetInstances(stage.ParameterCategorys, stage.ParameterCategory_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			parametercategoryIdent := localAliasPrefix + parametercategory.GongGetIdentifier(stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&models.ParameterCategory{Name: %s}).Stage(stage)", parametercategoryIdent, __gong__toRawStringLiteral(parametercategory.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", parametercategoryIdent, __gong__toRawStringLiteral(parametercategory.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.ComputedPrefix = %s", parametercategoryIdent, __gong__toRawStringLiteral(parametercategory.ComputedPrefix)))
			values.WriteString(fmt.Sprintf("\n\t%s.IsExpanded = %t", parametercategoryIdent, parametercategory.IsExpanded))
			for _, elem := range parametercategory.ParameterUse {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := localAliasPrefix + elem.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.ParameterUse = append(%s.ParameterUse, %s)", parametercategoryIdent, parametercategoryIdent, targetIdent))
			}
		}
	}
	if stage != nil {
		for _, parametercategoryuse := range __gong__sortStageSetInstances(stage.ParameterCategoryUses, stage.ParameterCategoryUse_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			parametercategoryuseIdent := localAliasPrefix + parametercategoryuse.GongGetIdentifier(stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&models.ParameterCategoryUse{Name: %s}).Stage(stage)", parametercategoryuseIdent, __gong__toRawStringLiteral(parametercategoryuse.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", parametercategoryuseIdent, __gong__toRawStringLiteral(parametercategoryuse.Name)))
			if parametercategoryuse.ParameterCategory != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := localAliasPrefix + parametercategoryuse.ParameterCategory.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.ParameterCategory = %s", parametercategoryuseIdent, targetIdent))
			}
		}
	}
	if stage != nil {
		for _, parametershape := range __gong__sortStageSetInstances(stage.ParameterShapes, stage.ParameterShape_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			parametershapeIdent := localAliasPrefix + parametershape.GongGetIdentifier(stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&models.ParameterShape{Name: %s}).Stage(stage)", parametershapeIdent, __gong__toRawStringLiteral(parametershape.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", parametershapeIdent, __gong__toRawStringLiteral(parametershape.Name)))
			if parametershape.Direction.ToCodeString() != "" {
				values.WriteString(fmt.Sprintf("\n\t%s.Direction = models.%s", parametershapeIdent, parametershape.Direction.ToCodeString()))
			} else {
				values.WriteString(fmt.Sprintf("\n\t%s.Direction = \"\"", parametershapeIdent))
			}
			values.WriteString(fmt.Sprintf("\n\t%s.ShapeIsComputedFromModel = %t", parametershapeIdent, parametershape.ShapeIsComputedFromModel))
			values.WriteString(fmt.Sprintf("\n\t%s.X = %f", parametershapeIdent, parametershape.X))
			values.WriteString(fmt.Sprintf("\n\t%s.Y = %f", parametershapeIdent, parametershape.Y))
			values.WriteString(fmt.Sprintf("\n\t%s.Width = %f", parametershapeIdent, parametershape.Width))
			values.WriteString(fmt.Sprintf("\n\t%s.Height = %f", parametershapeIdent, parametershape.Height))
			values.WriteString(fmt.Sprintf("\n\t%s.IsHidden = %t", parametershapeIdent, parametershape.IsHidden))
			values.WriteString(fmt.Sprintf("\n\t%s.IsPresentationOverridden = %t", parametershapeIdent, parametershape.IsPresentationOverridden))
			values.WriteString(fmt.Sprintf("\n\t%s.Color = %s", parametershapeIdent, __gong__toRawStringLiteral(parametershape.Color)))
			values.WriteString(fmt.Sprintf("\n\t%s.FillOpacity = %f", parametershapeIdent, parametershape.FillOpacity))
			values.WriteString(fmt.Sprintf("\n\t%s.Stroke = %s", parametershapeIdent, __gong__toRawStringLiteral(parametershape.Stroke)))
			values.WriteString(fmt.Sprintf("\n\t%s.StrokeWidth = %f", parametershapeIdent, parametershape.StrokeWidth))
			values.WriteString(fmt.Sprintf("\n\t%s.StrokeOpacity = %f", parametershapeIdent, parametershape.StrokeOpacity))
			values.WriteString(fmt.Sprintf("\n\t%s.StrokeDashArray = %s", parametershapeIdent, __gong__toRawStringLiteral(parametershape.StrokeDashArray)))
			values.WriteString(fmt.Sprintf("\n\t%s.RX = %f", parametershapeIdent, parametershape.RX))
			values.WriteString(fmt.Sprintf("\n\t%s.FontSize = %s", parametershapeIdent, __gong__toRawStringLiteral(parametershape.FontSize)))
			values.WriteString(fmt.Sprintf("\n\t%s.FontWeight = %s", parametershapeIdent, __gong__toRawStringLiteral(parametershape.FontWeight)))
			values.WriteString(fmt.Sprintf("\n\t%s.FontColor = %s", parametershapeIdent, __gong__toRawStringLiteral(parametershape.FontColor)))
			if parametershape.Parameter != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := localAliasPrefix + parametershape.Parameter.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Parameter = %s", parametershapeIdent, targetIdent))
			}
		}
	}
	if stage != nil {
		for _, parametersaggregate := range __gong__sortStageSetInstances(stage.ParametersAggregates, stage.ParametersAggregate_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			parametersaggregateIdent := localAliasPrefix + parametersaggregate.GongGetIdentifier(stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&models.ParametersAggregate{Name: %s}).Stage(stage)", parametersaggregateIdent, __gong__toRawStringLiteral(parametersaggregate.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", parametersaggregateIdent, __gong__toRawStringLiteral(parametersaggregate.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Tag = %s", parametersaggregateIdent, __gong__toRawStringLiteral(parametersaggregate.Tag)))
			values.WriteString(fmt.Sprintf("\n\t%s.Description = %s", parametersaggregateIdent, __gong__toRawStringLiteral(parametersaggregate.Description)))
			values.WriteString(fmt.Sprintf("\n\t%s.ComputedPrefix = %s", parametersaggregateIdent, __gong__toRawStringLiteral(parametersaggregate.ComputedPrefix)))
			values.WriteString(fmt.Sprintf("\n\t%s.IsExpanded = %t", parametersaggregateIdent, parametersaggregate.IsExpanded))
			for _, elem := range parametersaggregate.Parameters {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := localAliasPrefix + elem.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Parameters = append(%s.Parameters, %s)", parametersaggregateIdent, parametersaggregateIdent, targetIdent))
			}
		}
	}
	if stage != nil {
		for _, parametersaggregateshape := range __gong__sortStageSetInstances(stage.ParametersAggregateShapes, stage.ParametersAggregateShape_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			parametersaggregateshapeIdent := localAliasPrefix + parametersaggregateshape.GongGetIdentifier(stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&models.ParametersAggregateShape{Name: %s}).Stage(stage)", parametersaggregateshapeIdent, __gong__toRawStringLiteral(parametersaggregateshape.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", parametersaggregateshapeIdent, __gong__toRawStringLiteral(parametersaggregateshape.Name)))
			if parametersaggregateshape.Direction.ToCodeString() != "" {
				values.WriteString(fmt.Sprintf("\n\t%s.Direction = models.%s", parametersaggregateshapeIdent, parametersaggregateshape.Direction.ToCodeString()))
			} else {
				values.WriteString(fmt.Sprintf("\n\t%s.Direction = \"\"", parametersaggregateshapeIdent))
			}
			values.WriteString(fmt.Sprintf("\n\t%s.X = %f", parametersaggregateshapeIdent, parametersaggregateshape.X))
			values.WriteString(fmt.Sprintf("\n\t%s.Y = %f", parametersaggregateshapeIdent, parametersaggregateshape.Y))
			values.WriteString(fmt.Sprintf("\n\t%s.Width = %f", parametersaggregateshapeIdent, parametersaggregateshape.Width))
			values.WriteString(fmt.Sprintf("\n\t%s.Height = %f", parametersaggregateshapeIdent, parametersaggregateshape.Height))
			values.WriteString(fmt.Sprintf("\n\t%s.IsHidden = %t", parametersaggregateshapeIdent, parametersaggregateshape.IsHidden))
			values.WriteString(fmt.Sprintf("\n\t%s.IsPresentationOverridden = %t", parametersaggregateshapeIdent, parametersaggregateshape.IsPresentationOverridden))
			values.WriteString(fmt.Sprintf("\n\t%s.Color = %s", parametersaggregateshapeIdent, __gong__toRawStringLiteral(parametersaggregateshape.Color)))
			values.WriteString(fmt.Sprintf("\n\t%s.FillOpacity = %f", parametersaggregateshapeIdent, parametersaggregateshape.FillOpacity))
			values.WriteString(fmt.Sprintf("\n\t%s.Stroke = %s", parametersaggregateshapeIdent, __gong__toRawStringLiteral(parametersaggregateshape.Stroke)))
			values.WriteString(fmt.Sprintf("\n\t%s.StrokeWidth = %f", parametersaggregateshapeIdent, parametersaggregateshape.StrokeWidth))
			values.WriteString(fmt.Sprintf("\n\t%s.StrokeOpacity = %f", parametersaggregateshapeIdent, parametersaggregateshape.StrokeOpacity))
			values.WriteString(fmt.Sprintf("\n\t%s.StrokeDashArray = %s", parametersaggregateshapeIdent, __gong__toRawStringLiteral(parametersaggregateshape.StrokeDashArray)))
			values.WriteString(fmt.Sprintf("\n\t%s.RX = %f", parametersaggregateshapeIdent, parametersaggregateshape.RX))
			values.WriteString(fmt.Sprintf("\n\t%s.FontSize = %s", parametersaggregateshapeIdent, __gong__toRawStringLiteral(parametersaggregateshape.FontSize)))
			values.WriteString(fmt.Sprintf("\n\t%s.FontWeight = %s", parametersaggregateshapeIdent, __gong__toRawStringLiteral(parametersaggregateshape.FontWeight)))
			values.WriteString(fmt.Sprintf("\n\t%s.FontColor = %s", parametersaggregateshapeIdent, __gong__toRawStringLiteral(parametersaggregateshape.FontColor)))
			if parametersaggregateshape.ScenarioParameter != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := localAliasPrefix + parametersaggregateshape.ScenarioParameter.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.ScenarioParameter = %s", parametersaggregateshapeIdent, targetIdent))
			}
		}
	}
	if stage != nil {
		for _, position := range __gong__sortStageSetInstances(stage.Positions, stage.Position_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			positionIdent := localAliasPrefix + position.GongGetIdentifier(stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&models.Position{Name: %s}).Stage(stage)", positionIdent, __gong__toRawStringLiteral(position.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", positionIdent, __gong__toRawStringLiteral(position.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Date, _ = time.Parse(\"2006-01-02 15:04:05.999999999 -0700 MST\", \"%s\")", positionIdent, position.Date.String()))
			values.WriteString(fmt.Sprintf("\n\t%s.Ordinate = %f", positionIdent, position.Ordinate))
			values.WriteString(fmt.Sprintf("\n\t%s.ComputedPrefix = %s", positionIdent, __gong__toRawStringLiteral(position.ComputedPrefix)))
			values.WriteString(fmt.Sprintf("\n\t%s.IsExpanded = %t", positionIdent, position.IsExpanded))
		}
	}
	if stage != nil {
		for _, repository := range __gong__sortStageSetInstances(stage.Repositorys, stage.Repository_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			repositoryIdent := localAliasPrefix + repository.GongGetIdentifier(stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&models.Repository{Name: %s}).Stage(stage)", repositoryIdent, __gong__toRawStringLiteral(repository.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", repositoryIdent, __gong__toRawStringLiteral(repository.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.ComputedPrefix = %s", repositoryIdent, __gong__toRawStringLiteral(repository.ComputedPrefix)))
			values.WriteString(fmt.Sprintf("\n\t%s.IsExpanded = %t", repositoryIdent, repository.IsExpanded))
			for _, elem := range repository.ParameterUse {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := localAliasPrefix + elem.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.ParameterUse = append(%s.ParameterUse, %s)", repositoryIdent, repositoryIdent, targetIdent))
			}
			for _, elem := range repository.GroupUse {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := localAliasPrefix + elem.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.GroupUse = append(%s.GroupUse, %s)", repositoryIdent, repositoryIdent, targetIdent))
			}
		}
	}
	if stage != nil {
		for _, scenario := range __gong__sortStageSetInstances(stage.Scenarios, stage.Scenario_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			scenarioIdent := localAliasPrefix + scenario.GongGetIdentifier(stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&models.Scenario{Name: %s}).Stage(stage)", scenarioIdent, __gong__toRawStringLiteral(scenario.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", scenarioIdent, __gong__toRawStringLiteral(scenario.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Description = %s", scenarioIdent, __gong__toRawStringLiteral(scenario.Description)))
			values.WriteString(fmt.Sprintf("\n\t%s.IsDiagramsNodeExpanded = %t", scenarioIdent, scenario.IsDiagramsNodeExpanded))
			values.WriteString(fmt.Sprintf("\n\t%s.IsActorStatesNodeExpanded = %t", scenarioIdent, scenario.IsActorStatesNodeExpanded))
			values.WriteString(fmt.Sprintf("\n\t%s.IsActorStateTransitionsNodeExpanded = %t", scenarioIdent, scenario.IsActorStateTransitionsNodeExpanded))
			values.WriteString(fmt.Sprintf("\n\t%s.IsEvolutionDirectionsNodeExpanded = %t", scenarioIdent, scenario.IsEvolutionDirectionsNodeExpanded))
			values.WriteString(fmt.Sprintf("\n\t%s.IsParametersNodeExpanded = %t", scenarioIdent, scenario.IsParametersNodeExpanded))
			values.WriteString(fmt.Sprintf("\n\t%s.IsParametersAggretatesNodeExpanded = %t", scenarioIdent, scenario.IsParametersAggretatesNodeExpanded))
			values.WriteString(fmt.Sprintf("\n\t%s.ComputedPrefix = %s", scenarioIdent, __gong__toRawStringLiteral(scenario.ComputedPrefix)))
			values.WriteString(fmt.Sprintf("\n\t%s.IsExpanded = %t", scenarioIdent, scenario.IsExpanded))
			for _, elem := range scenario.Diagrams {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := localAliasPrefix + elem.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Diagrams = append(%s.Diagrams, %s)", scenarioIdent, scenarioIdent, targetIdent))
			}
			for _, elem := range scenario.ActorStates {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := localAliasPrefix + elem.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.ActorStates = append(%s.ActorStates, %s)", scenarioIdent, scenarioIdent, targetIdent))
			}
			for _, elem := range scenario.ActorStateTransitions {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := localAliasPrefix + elem.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.ActorStateTransitions = append(%s.ActorStateTransitions, %s)", scenarioIdent, scenarioIdent, targetIdent))
			}
			for _, elem := range scenario.EvolutionDirections {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := localAliasPrefix + elem.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.EvolutionDirections = append(%s.EvolutionDirections, %s)", scenarioIdent, scenarioIdent, targetIdent))
			}
			for _, elem := range scenario.Parameters {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := localAliasPrefix + elem.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Parameters = append(%s.Parameters, %s)", scenarioIdent, scenarioIdent, targetIdent))
			}
			for _, elem := range scenario.ParametersAggretates {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := localAliasPrefix + elem.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.ParametersAggretates = append(%s.ParametersAggretates, %s)", scenarioIdent, scenarioIdent, targetIdent))
			}
		}
	}
	if stage != nil {
		for _, user := range __gong__sortStageSetInstances(stage.Users, stage.User_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			userIdent := localAliasPrefix + user.GongGetIdentifier(stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&models.User{Name: %s}).Stage(stage)", userIdent, __gong__toRawStringLiteral(user.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", userIdent, __gong__toRawStringLiteral(user.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.ComputedPrefix = %s", userIdent, __gong__toRawStringLiteral(user.ComputedPrefix)))
			values.WriteString(fmt.Sprintf("\n\t%s.IsExpanded = %t", userIdent, user.IsExpanded))
		}
	}
	if stage != nil {
		for _, useruse := range __gong__sortStageSetInstances(stage.UserUses, stage.UserUse_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			useruseIdent := localAliasPrefix + useruse.GongGetIdentifier(stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&models.UserUse{Name: %s}).Stage(stage)", useruseIdent, __gong__toRawStringLiteral(useruse.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", useruseIdent, __gong__toRawStringLiteral(useruse.Name)))
			if useruse.User != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := localAliasPrefix + useruse.User.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.User = %s", useruseIdent, targetIdent))
			}
		}
	}
	if stage != nil {
		for _, workspace := range __gong__sortStageSetInstances(stage.Workspaces, stage.Workspace_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			workspaceIdent := localAliasPrefix + workspace.GongGetIdentifier(stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&models.Workspace{Name: %s}).Stage(stage)", workspaceIdent, __gong__toRawStringLiteral(workspace.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", workspaceIdent, __gong__toRawStringLiteral(workspace.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.ComputedPrefix = %s", workspaceIdent, __gong__toRawStringLiteral(workspace.ComputedPrefix)))
			values.WriteString(fmt.Sprintf("\n\t%s.IsExpanded = %t", workspaceIdent, workspace.IsExpanded))
			if workspace.SelectedDiagram != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := localAliasPrefix + workspace.SelectedDiagram.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.SelectedDiagram = %s", workspaceIdent, targetIdent))
			}
			if workspace.Default_EvolutionDirectionShape != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := localAliasPrefix + workspace.Default_EvolutionDirectionShape.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Default_EvolutionDirectionShape = %s", workspaceIdent, targetIdent))
			}
			if workspace.Default_ParameterShape != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := localAliasPrefix + workspace.Default_ParameterShape.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Default_ParameterShape = %s", workspaceIdent, targetIdent))
			}
			if workspace.Default_ScenarioParameterShape != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := localAliasPrefix + workspace.Default_ScenarioParameterShape.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Default_ScenarioParameterShape = %s", workspaceIdent, targetIdent))
			}
			if workspace.Default_ActorStateShape != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := localAliasPrefix + workspace.Default_ActorStateShape.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Default_ActorStateShape = %s", workspaceIdent, targetIdent))
			}
			if workspace.Default_ActorStateTransitionShape != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := localAliasPrefix + workspace.Default_ActorStateTransitionShape.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Default_ActorStateTransitionShape = %s", workspaceIdent, targetIdent))
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
go run github.com/fullstack-lang/gong/dsm/scenario/go/models/probe/cmd/probe@%s stage.go
*/
package %s

import (
	"slices"
	"time"

	"github.com/fullstack-lang/gong/dsm/scenario/go/models"%s
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

	if modelsPackageName != "" && modelsPackageName != "github.com/fullstack-lang/gong/dsm/scenario/go/models" {
		res = strings.ReplaceAll(res, "\"github.com/fullstack-lang/gong/dsm/scenario/go/models\"", fmt.Sprintf("%q", modelsPackageName))
	}

	return res, nil
}

// insertion point for marshall field methods
func (actorstate *ActorState) GongMarshallField(stage *Stage, fieldName string) (res string) {
	ident := actorstate.GongGetIdentifier(stage)
	_ = ident
	switch fieldName {
	case "Name":
		res = __gong__marshallString(ident, "Name", actorstate.Name)
	case "Description":
		res = __gong__marshallString(ident, "Description", actorstate.Description)
	case "IsWithProbaility":
		res = __gong__marshallBool(ident, "IsWithProbaility", actorstate.IsWithProbaility)
	case "Probability":
		res = __gong__marshallEnumString(ident, "Probability", actorstate.Probability.ToCodeString())
	case "ComputedPrefix":
		res = __gong__marshallString(ident, "ComputedPrefix", actorstate.ComputedPrefix)
	case "IsExpanded":
		res = __gong__marshallBool(ident, "IsExpanded", actorstate.IsExpanded)

	default:
		log.Panicf("Unknown field %s for Gongstruct ActorState", fieldName)
	}
	return
}

func (actorstateshape *ActorStateShape) GongMarshallField(stage *Stage, fieldName string) (res string) {
	ident := actorstateshape.GongGetIdentifier(stage)
	_ = ident
	switch fieldName {
	case "Name":
		res = __gong__marshallString(ident, "Name", actorstateshape.Name)
	case "X":
		res = __gong__marshallFloat(ident, "X", actorstateshape.X)
	case "Y":
		res = __gong__marshallFloat(ident, "Y", actorstateshape.Y)
	case "Width":
		res = __gong__marshallFloat(ident, "Width", actorstateshape.Width)
	case "Height":
		res = __gong__marshallFloat(ident, "Height", actorstateshape.Height)
	case "IsHidden":
		res = __gong__marshallBool(ident, "IsHidden", actorstateshape.IsHidden)
	case "IsPresentationOverridden":
		res = __gong__marshallBool(ident, "IsPresentationOverridden", actorstateshape.IsPresentationOverridden)
	case "Color":
		res = __gong__marshallString(ident, "Color", actorstateshape.Color)
	case "FillOpacity":
		res = __gong__marshallFloat(ident, "FillOpacity", actorstateshape.FillOpacity)
	case "Stroke":
		res = __gong__marshallString(ident, "Stroke", actorstateshape.Stroke)
	case "StrokeWidth":
		res = __gong__marshallFloat(ident, "StrokeWidth", actorstateshape.StrokeWidth)
	case "StrokeOpacity":
		res = __gong__marshallFloat(ident, "StrokeOpacity", actorstateshape.StrokeOpacity)
	case "StrokeDashArray":
		res = __gong__marshallString(ident, "StrokeDashArray", actorstateshape.StrokeDashArray)
	case "RX":
		res = __gong__marshallFloat(ident, "RX", actorstateshape.RX)
	case "FontSize":
		res = __gong__marshallString(ident, "FontSize", actorstateshape.FontSize)
	case "FontWeight":
		res = __gong__marshallString(ident, "FontWeight", actorstateshape.FontWeight)
	case "FontColor":
		res = __gong__marshallString(ident, "FontColor", actorstateshape.FontColor)

	case "ActorState":
		if actorstateshape.ActorState != nil {
			res = __gong__marshallPointer(ident, "ActorState", actorstateshape.ActorState.GongGetIdentifier(stage))
		} else {
			res = __gong__marshallPointer(ident, "ActorState", "nil")
		}
	default:
		log.Panicf("Unknown field %s for Gongstruct ActorStateShape", fieldName)
	}
	return
}

func (actorstatetransition *ActorStateTransition) GongMarshallField(stage *Stage, fieldName string) (res string) {
	ident := actorstatetransition.GongGetIdentifier(stage)
	_ = ident
	switch fieldName {
	case "Name":
		res = __gong__marshallString(ident, "Name", actorstatetransition.Name)
	case "ComputedPrefix":
		res = __gong__marshallString(ident, "ComputedPrefix", actorstatetransition.ComputedPrefix)
	case "IsExpanded":
		res = __gong__marshallBool(ident, "IsExpanded", actorstatetransition.IsExpanded)

	case "StartState":
		if actorstatetransition.StartState != nil {
			res = __gong__marshallPointer(ident, "StartState", actorstatetransition.StartState.GongGetIdentifier(stage))
		} else {
			res = __gong__marshallPointer(ident, "StartState", "nil")
		}
	case "EndState":
		if actorstatetransition.EndState != nil {
			res = __gong__marshallPointer(ident, "EndState", actorstatetransition.EndState.GongGetIdentifier(stage))
		} else {
			res = __gong__marshallPointer(ident, "EndState", "nil")
		}
	case "Justifications":
		var sb strings.Builder
		for _, _parameter := range actorstatetransition.Justifications {
			sb.WriteString(__gong__marshallSliceOfPointers(ident, "Justifications", _parameter.GongGetIdentifier(stage)))
		}
		res = sb.String()
	default:
		log.Panicf("Unknown field %s for Gongstruct ActorStateTransition", fieldName)
	}
	return
}

func (actorstatetransitionshape *ActorStateTransitionShape) GongMarshallField(stage *Stage, fieldName string) (res string) {
	ident := actorstatetransitionshape.GongGetIdentifier(stage)
	_ = ident
	switch fieldName {
	case "Name":
		res = __gong__marshallString(ident, "Name", actorstatetransitionshape.Name)
	case "X":
		res = __gong__marshallFloat(ident, "X", actorstatetransitionshape.X)
	case "Y":
		res = __gong__marshallFloat(ident, "Y", actorstatetransitionshape.Y)
	case "Width":
		res = __gong__marshallFloat(ident, "Width", actorstatetransitionshape.Width)
	case "Height":
		res = __gong__marshallFloat(ident, "Height", actorstatetransitionshape.Height)
	case "IsHidden":
		res = __gong__marshallBool(ident, "IsHidden", actorstatetransitionshape.IsHidden)
	case "IsPresentationOverridden":
		res = __gong__marshallBool(ident, "IsPresentationOverridden", actorstatetransitionshape.IsPresentationOverridden)
	case "Color":
		res = __gong__marshallString(ident, "Color", actorstatetransitionshape.Color)
	case "FillOpacity":
		res = __gong__marshallFloat(ident, "FillOpacity", actorstatetransitionshape.FillOpacity)
	case "Stroke":
		res = __gong__marshallString(ident, "Stroke", actorstatetransitionshape.Stroke)
	case "StrokeWidth":
		res = __gong__marshallFloat(ident, "StrokeWidth", actorstatetransitionshape.StrokeWidth)
	case "StrokeOpacity":
		res = __gong__marshallFloat(ident, "StrokeOpacity", actorstatetransitionshape.StrokeOpacity)
	case "StrokeDashArray":
		res = __gong__marshallString(ident, "StrokeDashArray", actorstatetransitionshape.StrokeDashArray)
	case "RX":
		res = __gong__marshallFloat(ident, "RX", actorstatetransitionshape.RX)
	case "FontSize":
		res = __gong__marshallString(ident, "FontSize", actorstatetransitionshape.FontSize)
	case "FontWeight":
		res = __gong__marshallString(ident, "FontWeight", actorstatetransitionshape.FontWeight)
	case "FontColor":
		res = __gong__marshallString(ident, "FontColor", actorstatetransitionshape.FontColor)

	case "ActorStateTransition":
		if actorstatetransitionshape.ActorStateTransition != nil {
			res = __gong__marshallPointer(ident, "ActorStateTransition", actorstatetransitionshape.ActorStateTransition.GongGetIdentifier(stage))
		} else {
			res = __gong__marshallPointer(ident, "ActorStateTransition", "nil")
		}
	case "Start":
		if actorstatetransitionshape.Start != nil {
			res = __gong__marshallPointer(ident, "Start", actorstatetransitionshape.Start.GongGetIdentifier(stage))
		} else {
			res = __gong__marshallPointer(ident, "Start", "nil")
		}
	case "End":
		if actorstatetransitionshape.End != nil {
			res = __gong__marshallPointer(ident, "End", actorstatetransitionshape.End.GongGetIdentifier(stage))
		} else {
			res = __gong__marshallPointer(ident, "End", "nil")
		}
	case "ControlPointShapes":
		var sb strings.Builder
		for _, _controlpointshape := range actorstatetransitionshape.ControlPointShapes {
			sb.WriteString(__gong__marshallSliceOfPointers(ident, "ControlPointShapes", _controlpointshape.GongGetIdentifier(stage)))
		}
		res = sb.String()
	default:
		log.Panicf("Unknown field %s for Gongstruct ActorStateTransitionShape", fieldName)
	}
	return
}

func (analysis *Analysis) GongMarshallField(stage *Stage, fieldName string) (res string) {
	ident := analysis.GongGetIdentifier(stage)
	_ = ident
	switch fieldName {
	case "Name":
		res = __gong__marshallString(ident, "Name", analysis.Name)
	case "Description":
		res = __gong__marshallString(ident, "Description", analysis.Description)
	case "IsScenariosNodeExpanded":
		res = __gong__marshallBool(ident, "IsScenariosNodeExpanded", analysis.IsScenariosNodeExpanded)
	case "IsGroupUseNodeExpanded":
		res = __gong__marshallBool(ident, "IsGroupUseNodeExpanded", analysis.IsGroupUseNodeExpanded)
	case "IsGeoObjectUseNodeExpanded":
		res = __gong__marshallBool(ident, "IsGeoObjectUseNodeExpanded", analysis.IsGeoObjectUseNodeExpanded)
	case "IsMapUseNodeExpanded":
		res = __gong__marshallBool(ident, "IsMapUseNodeExpanded", analysis.IsMapUseNodeExpanded)
	case "ComputedPrefix":
		res = __gong__marshallString(ident, "ComputedPrefix", analysis.ComputedPrefix)
	case "IsExpanded":
		res = __gong__marshallBool(ident, "IsExpanded", analysis.IsExpanded)

	case "Scenarios":
		var sb strings.Builder
		for _, _scenario := range analysis.Scenarios {
			sb.WriteString(__gong__marshallSliceOfPointers(ident, "Scenarios", _scenario.GongGetIdentifier(stage)))
		}
		res = sb.String()
	case "GroupUse":
		var sb strings.Builder
		for _, _groupuse := range analysis.GroupUse {
			sb.WriteString(__gong__marshallSliceOfPointers(ident, "GroupUse", _groupuse.GongGetIdentifier(stage)))
		}
		res = sb.String()
	case "GeoObjectUse":
		var sb strings.Builder
		for _, _geoobjectuse := range analysis.GeoObjectUse {
			sb.WriteString(__gong__marshallSliceOfPointers(ident, "GeoObjectUse", _geoobjectuse.GongGetIdentifier(stage)))
		}
		res = sb.String()
	case "MapUse":
		var sb strings.Builder
		for _, _mapobjectuse := range analysis.MapUse {
			sb.WriteString(__gong__marshallSliceOfPointers(ident, "MapUse", _mapobjectuse.GongGetIdentifier(stage)))
		}
		res = sb.String()
	default:
		log.Panicf("Unknown field %s for Gongstruct Analysis", fieldName)
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
	case "IsShowPrefix":
		res = __gong__marshallBool(ident, "IsShowPrefix", diagram.IsShowPrefix)
	case "Description":
		res = __gong__marshallString(ident, "Description", diagram.Description)
	case "IsEvolutionDirectionsNodeExpanded":
		res = __gong__marshallBool(ident, "IsEvolutionDirectionsNodeExpanded", diagram.IsEvolutionDirectionsNodeExpanded)
	case "IsActorStatesNodeExpanded":
		res = __gong__marshallBool(ident, "IsActorStatesNodeExpanded", diagram.IsActorStatesNodeExpanded)
	case "IsParametersNodeExpanded":
		res = __gong__marshallBool(ident, "IsParametersNodeExpanded", diagram.IsParametersNodeExpanded)
	case "IsParametersAggregatesNodeExpanded":
		res = __gong__marshallBool(ident, "IsParametersAggregatesNodeExpanded", diagram.IsParametersAggregatesNodeExpanded)
	case "IsActorStateTransitionsNodeExpanded":
		res = __gong__marshallBool(ident, "IsActorStateTransitionsNodeExpanded", diagram.IsActorStateTransitionsNodeExpanded)
	case "AxisOrign_X":
		res = __gong__marshallFloat(ident, "AxisOrign_X", diagram.AxisOrign_X)
	case "AxisOrign_Y":
		res = __gong__marshallFloat(ident, "AxisOrign_Y", diagram.AxisOrign_Y)
	case "VerticalAxis_Top_Y":
		res = __gong__marshallFloat(ident, "VerticalAxis_Top_Y", diagram.VerticalAxis_Top_Y)
	case "VerticalAxis_Bottom_Y":
		res = __gong__marshallFloat(ident, "VerticalAxis_Bottom_Y", diagram.VerticalAxis_Bottom_Y)
	case "VerticalAxis_StrokeWidth":
		res = __gong__marshallFloat(ident, "VerticalAxis_StrokeWidth", diagram.VerticalAxis_StrokeWidth)
	case "HorizontalAxis_Right_X":
		res = __gong__marshallFloat(ident, "HorizontalAxis_Right_X", diagram.HorizontalAxis_Right_X)
	case "Start":
		res = __gong__marshallTime(ident, "Start", diagram.Start.String())
	case "End":
		res = __gong__marshallTime(ident, "End", diagram.End.String())
	case "NumberOfYearsBetweenTicks":
		res = __gong__marshallInt(ident, "NumberOfYearsBetweenTicks", diagram.NumberOfYearsBetweenTicks)

	case "EvolutionDirectionShapes":
		var sb strings.Builder
		for _, _evolutiondirectionshape := range diagram.EvolutionDirectionShapes {
			sb.WriteString(__gong__marshallSliceOfPointers(ident, "EvolutionDirectionShapes", _evolutiondirectionshape.GongGetIdentifier(stage)))
		}
		res = sb.String()
	case "EvolutionDirectionsWhoseNodeIsExpanded":
		var sb strings.Builder
		for _, _evolutiondirection := range diagram.EvolutionDirectionsWhoseNodeIsExpanded {
			sb.WriteString(__gong__marshallSliceOfPointers(ident, "EvolutionDirectionsWhoseNodeIsExpanded", _evolutiondirection.GongGetIdentifier(stage)))
		}
		res = sb.String()
	case "ActorStateShapes":
		var sb strings.Builder
		for _, _actorstateshape := range diagram.ActorStateShapes {
			sb.WriteString(__gong__marshallSliceOfPointers(ident, "ActorStateShapes", _actorstateshape.GongGetIdentifier(stage)))
		}
		res = sb.String()
	case "ActorStatesWhoseNodeIsExpanded":
		var sb strings.Builder
		for _, _actorstate := range diagram.ActorStatesWhoseNodeIsExpanded {
			sb.WriteString(__gong__marshallSliceOfPointers(ident, "ActorStatesWhoseNodeIsExpanded", _actorstate.GongGetIdentifier(stage)))
		}
		res = sb.String()
	case "ParameterShapes":
		var sb strings.Builder
		for _, _parametershape := range diagram.ParameterShapes {
			sb.WriteString(__gong__marshallSliceOfPointers(ident, "ParameterShapes", _parametershape.GongGetIdentifier(stage)))
		}
		res = sb.String()
	case "ParametersWhoseNodeIsExpanded":
		var sb strings.Builder
		for _, _parameter := range diagram.ParametersWhoseNodeIsExpanded {
			sb.WriteString(__gong__marshallSliceOfPointers(ident, "ParametersWhoseNodeIsExpanded", _parameter.GongGetIdentifier(stage)))
		}
		res = sb.String()
	case "ScenarioParameterShapes":
		var sb strings.Builder
		for _, _parametersaggregateshape := range diagram.ScenarioParameterShapes {
			sb.WriteString(__gong__marshallSliceOfPointers(ident, "ScenarioParameterShapes", _parametersaggregateshape.GongGetIdentifier(stage)))
		}
		res = sb.String()
	case "ParametersAggregatesWhoseNodeIsExpanded":
		var sb strings.Builder
		for _, _parametersaggregate := range diagram.ParametersAggregatesWhoseNodeIsExpanded {
			sb.WriteString(__gong__marshallSliceOfPointers(ident, "ParametersAggregatesWhoseNodeIsExpanded", _parametersaggregate.GongGetIdentifier(stage)))
		}
		res = sb.String()
	case "ActorStateTransitionShapes":
		var sb strings.Builder
		for _, _actorstatetransitionshape := range diagram.ActorStateTransitionShapes {
			sb.WriteString(__gong__marshallSliceOfPointers(ident, "ActorStateTransitionShapes", _actorstatetransitionshape.GongGetIdentifier(stage)))
		}
		res = sb.String()
	case "ActorStateTransitionsWhoseNodeIsExpanded":
		var sb strings.Builder
		for _, _actorstatetransition := range diagram.ActorStateTransitionsWhoseNodeIsExpanded {
			sb.WriteString(__gong__marshallSliceOfPointers(ident, "ActorStateTransitionsWhoseNodeIsExpanded", _actorstatetransition.GongGetIdentifier(stage)))
		}
		res = sb.String()
	default:
		log.Panicf("Unknown field %s for Gongstruct Diagram", fieldName)
	}
	return
}

func (document *Document) GongMarshallField(stage *Stage, fieldName string) (res string) {
	ident := document.GongGetIdentifier(stage)
	_ = ident
	switch fieldName {
	case "Name":
		res = __gong__marshallString(ident, "Name", document.Name)
	case "ComputedPrefix":
		res = __gong__marshallString(ident, "ComputedPrefix", document.ComputedPrefix)
	case "IsExpanded":
		res = __gong__marshallBool(ident, "IsExpanded", document.IsExpanded)

	case "GeoObjectUse":
		var sb strings.Builder
		for _, _geoobjectuse := range document.GeoObjectUse {
			sb.WriteString(__gong__marshallSliceOfPointers(ident, "GeoObjectUse", _geoobjectuse.GongGetIdentifier(stage)))
		}
		res = sb.String()
	default:
		log.Panicf("Unknown field %s for Gongstruct Document", fieldName)
	}
	return
}

func (documentuse *DocumentUse) GongMarshallField(stage *Stage, fieldName string) (res string) {
	ident := documentuse.GongGetIdentifier(stage)
	_ = ident
	switch fieldName {
	case "Name":
		res = __gong__marshallString(ident, "Name", documentuse.Name)

	case "Document":
		if documentuse.Document != nil {
			res = __gong__marshallPointer(ident, "Document", documentuse.Document.GongGetIdentifier(stage))
		} else {
			res = __gong__marshallPointer(ident, "Document", "nil")
		}
	default:
		log.Panicf("Unknown field %s for Gongstruct DocumentUse", fieldName)
	}
	return
}

func (evolutiondirection *EvolutionDirection) GongMarshallField(stage *Stage, fieldName string) (res string) {
	ident := evolutiondirection.GongGetIdentifier(stage)
	_ = ident
	switch fieldName {
	case "Name":
		res = __gong__marshallString(ident, "Name", evolutiondirection.Name)
	case "Description":
		res = __gong__marshallString(ident, "Description", evolutiondirection.Description)
	case "ComputedPrefix":
		res = __gong__marshallString(ident, "ComputedPrefix", evolutiondirection.ComputedPrefix)
	case "IsExpanded":
		res = __gong__marshallBool(ident, "IsExpanded", evolutiondirection.IsExpanded)

	default:
		log.Panicf("Unknown field %s for Gongstruct EvolutionDirection", fieldName)
	}
	return
}

func (evolutiondirectionshape *EvolutionDirectionShape) GongMarshallField(stage *Stage, fieldName string) (res string) {
	ident := evolutiondirectionshape.GongGetIdentifier(stage)
	_ = ident
	switch fieldName {
	case "Name":
		res = __gong__marshallString(ident, "Name", evolutiondirectionshape.Name)
	case "X":
		res = __gong__marshallFloat(ident, "X", evolutiondirectionshape.X)
	case "Y":
		res = __gong__marshallFloat(ident, "Y", evolutiondirectionshape.Y)
	case "Width":
		res = __gong__marshallFloat(ident, "Width", evolutiondirectionshape.Width)
	case "Height":
		res = __gong__marshallFloat(ident, "Height", evolutiondirectionshape.Height)
	case "IsHidden":
		res = __gong__marshallBool(ident, "IsHidden", evolutiondirectionshape.IsHidden)
	case "IsPresentationOverridden":
		res = __gong__marshallBool(ident, "IsPresentationOverridden", evolutiondirectionshape.IsPresentationOverridden)
	case "Color":
		res = __gong__marshallString(ident, "Color", evolutiondirectionshape.Color)
	case "FillOpacity":
		res = __gong__marshallFloat(ident, "FillOpacity", evolutiondirectionshape.FillOpacity)
	case "Stroke":
		res = __gong__marshallString(ident, "Stroke", evolutiondirectionshape.Stroke)
	case "StrokeWidth":
		res = __gong__marshallFloat(ident, "StrokeWidth", evolutiondirectionshape.StrokeWidth)
	case "StrokeOpacity":
		res = __gong__marshallFloat(ident, "StrokeOpacity", evolutiondirectionshape.StrokeOpacity)
	case "StrokeDashArray":
		res = __gong__marshallString(ident, "StrokeDashArray", evolutiondirectionshape.StrokeDashArray)
	case "RX":
		res = __gong__marshallFloat(ident, "RX", evolutiondirectionshape.RX)
	case "FontSize":
		res = __gong__marshallString(ident, "FontSize", evolutiondirectionshape.FontSize)
	case "FontWeight":
		res = __gong__marshallString(ident, "FontWeight", evolutiondirectionshape.FontWeight)
	case "FontColor":
		res = __gong__marshallString(ident, "FontColor", evolutiondirectionshape.FontColor)

	case "EvolutionDirection":
		if evolutiondirectionshape.EvolutionDirection != nil {
			res = __gong__marshallPointer(ident, "EvolutionDirection", evolutiondirectionshape.EvolutionDirection.GongGetIdentifier(stage))
		} else {
			res = __gong__marshallPointer(ident, "EvolutionDirection", "nil")
		}
	default:
		log.Panicf("Unknown field %s for Gongstruct EvolutionDirectionShape", fieldName)
	}
	return
}

func (foo *Foo) GongMarshallField(stage *Stage, fieldName string) (res string) {
	ident := foo.GongGetIdentifier(stage)
	_ = ident
	switch fieldName {
	case "Name":
		res = __gong__marshallString(ident, "Name", foo.Name)

	default:
		log.Panicf("Unknown field %s for Gongstruct Foo", fieldName)
	}
	return
}

func (geoobject *GeoObject) GongMarshallField(stage *Stage, fieldName string) (res string) {
	ident := geoobject.GongGetIdentifier(stage)
	_ = ident
	switch fieldName {
	case "Name":
		res = __gong__marshallString(ident, "Name", geoobject.Name)
	case "ComputedPrefix":
		res = __gong__marshallString(ident, "ComputedPrefix", geoobject.ComputedPrefix)
	case "IsExpanded":
		res = __gong__marshallBool(ident, "IsExpanded", geoobject.IsExpanded)

	default:
		log.Panicf("Unknown field %s for Gongstruct GeoObject", fieldName)
	}
	return
}

func (geoobjectuse *GeoObjectUse) GongMarshallField(stage *Stage, fieldName string) (res string) {
	ident := geoobjectuse.GongGetIdentifier(stage)
	_ = ident
	switch fieldName {
	case "Name":
		res = __gong__marshallString(ident, "Name", geoobjectuse.Name)

	case "GeoObject":
		if geoobjectuse.GeoObject != nil {
			res = __gong__marshallPointer(ident, "GeoObject", geoobjectuse.GeoObject.GongGetIdentifier(stage))
		} else {
			res = __gong__marshallPointer(ident, "GeoObject", "nil")
		}
	default:
		log.Panicf("Unknown field %s for Gongstruct GeoObjectUse", fieldName)
	}
	return
}

func (group *Group) GongMarshallField(stage *Stage, fieldName string) (res string) {
	ident := group.GongGetIdentifier(stage)
	_ = ident
	switch fieldName {
	case "Name":
		res = __gong__marshallString(ident, "Name", group.Name)
	case "ComputedPrefix":
		res = __gong__marshallString(ident, "ComputedPrefix", group.ComputedPrefix)
	case "IsExpanded":
		res = __gong__marshallBool(ident, "IsExpanded", group.IsExpanded)

	case "UserUse":
		var sb strings.Builder
		for _, _useruse := range group.UserUse {
			sb.WriteString(__gong__marshallSliceOfPointers(ident, "UserUse", _useruse.GongGetIdentifier(stage)))
		}
		res = sb.String()
	default:
		log.Panicf("Unknown field %s for Gongstruct Group", fieldName)
	}
	return
}

func (groupuse *GroupUse) GongMarshallField(stage *Stage, fieldName string) (res string) {
	ident := groupuse.GongGetIdentifier(stage)
	_ = ident
	switch fieldName {
	case "Name":
		res = __gong__marshallString(ident, "Name", groupuse.Name)

	case "Group":
		if groupuse.Group != nil {
			res = __gong__marshallPointer(ident, "Group", groupuse.Group.GongGetIdentifier(stage))
		} else {
			res = __gong__marshallPointer(ident, "Group", "nil")
		}
	default:
		log.Panicf("Unknown field %s for Gongstruct GroupUse", fieldName)
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
	case "IsAnalysesNodeExpanded":
		res = __gong__marshallBool(ident, "IsAnalysesNodeExpanded", library.IsAnalysesNodeExpanded)
	case "IsSubLibrariesNodeExpanded":
		res = __gong__marshallBool(ident, "IsSubLibrariesNodeExpanded", library.IsSubLibrariesNodeExpanded)
	case "NbPixPerCharacter":
		res = __gong__marshallFloat(ident, "NbPixPerCharacter", library.NbPixPerCharacter)
	case "LogoSVGFile":
		res = __gong__marshallString(ident, "LogoSVGFile", library.LogoSVGFile)
	case "IsExpandedTmp":
		res = __gong__marshallBool(ident, "IsExpandedTmp", library.IsExpandedTmp)

	case "Analyses":
		var sb strings.Builder
		for _, _analysis := range library.Analyses {
			sb.WriteString(__gong__marshallSliceOfPointers(ident, "Analyses", _analysis.GongGetIdentifier(stage)))
		}
		res = sb.String()
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
	default:
		log.Panicf("Unknown field %s for Gongstruct Library", fieldName)
	}
	return
}

func (mapobject *MapObject) GongMarshallField(stage *Stage, fieldName string) (res string) {
	ident := mapobject.GongGetIdentifier(stage)
	_ = ident
	switch fieldName {
	case "Name":
		res = __gong__marshallString(ident, "Name", mapobject.Name)
	case "ComputedPrefix":
		res = __gong__marshallString(ident, "ComputedPrefix", mapobject.ComputedPrefix)
	case "IsExpanded":
		res = __gong__marshallBool(ident, "IsExpanded", mapobject.IsExpanded)

	default:
		log.Panicf("Unknown field %s for Gongstruct MapObject", fieldName)
	}
	return
}

func (mapobjectuse *MapObjectUse) GongMarshallField(stage *Stage, fieldName string) (res string) {
	ident := mapobjectuse.GongGetIdentifier(stage)
	_ = ident
	switch fieldName {
	case "Name":
		res = __gong__marshallString(ident, "Name", mapobjectuse.Name)

	case "Map":
		if mapobjectuse.Map != nil {
			res = __gong__marshallPointer(ident, "Map", mapobjectuse.Map.GongGetIdentifier(stage))
		} else {
			res = __gong__marshallPointer(ident, "Map", "nil")
		}
	default:
		log.Panicf("Unknown field %s for Gongstruct MapObjectUse", fieldName)
	}
	return
}

func (parameter *Parameter) GongMarshallField(stage *Stage, fieldName string) (res string) {
	ident := parameter.GongGetIdentifier(stage)
	_ = ident
	switch fieldName {
	case "Name":
		res = __gong__marshallString(ident, "Name", parameter.Name)
	case "Description":
		res = __gong__marshallString(ident, "Description", parameter.Description)
	case "IsResponse":
		res = __gong__marshallBool(ident, "IsResponse", parameter.IsResponse)
	case "Start":
		res = __gong__marshallTime(ident, "Start", parameter.Start.String())
	case "End":
		res = __gong__marshallTime(ident, "End", parameter.End.String())
	case "Force":
		res = __gong__marshallFloat(ident, "Force", parameter.Force)
	case "Tag":
		res = __gong__marshallString(ident, "Tag", parameter.Tag)
	case "ComputedPrefix":
		res = __gong__marshallString(ident, "ComputedPrefix", parameter.ComputedPrefix)
	case "IsExpanded":
		res = __gong__marshallBool(ident, "IsExpanded", parameter.IsExpanded)

	case "GroupUse":
		var sb strings.Builder
		for _, _groupuse := range parameter.GroupUse {
			sb.WriteString(__gong__marshallSliceOfPointers(ident, "GroupUse", _groupuse.GongGetIdentifier(stage)))
		}
		res = sb.String()
	case "DocumentUse":
		var sb strings.Builder
		for _, _documentuse := range parameter.DocumentUse {
			sb.WriteString(__gong__marshallSliceOfPointers(ident, "DocumentUse", _documentuse.GongGetIdentifier(stage)))
		}
		res = sb.String()
	case "GeoObjectUse":
		var sb strings.Builder
		for _, _geoobjectuse := range parameter.GeoObjectUse {
			sb.WriteString(__gong__marshallSliceOfPointers(ident, "GeoObjectUse", _geoobjectuse.GongGetIdentifier(stage)))
		}
		res = sb.String()
	default:
		log.Panicf("Unknown field %s for Gongstruct Parameter", fieldName)
	}
	return
}

func (parametercategory *ParameterCategory) GongMarshallField(stage *Stage, fieldName string) (res string) {
	ident := parametercategory.GongGetIdentifier(stage)
	_ = ident
	switch fieldName {
	case "Name":
		res = __gong__marshallString(ident, "Name", parametercategory.Name)
	case "ComputedPrefix":
		res = __gong__marshallString(ident, "ComputedPrefix", parametercategory.ComputedPrefix)
	case "IsExpanded":
		res = __gong__marshallBool(ident, "IsExpanded", parametercategory.IsExpanded)

	case "ParameterUse":
		var sb strings.Builder
		for _, _parametershape := range parametercategory.ParameterUse {
			sb.WriteString(__gong__marshallSliceOfPointers(ident, "ParameterUse", _parametershape.GongGetIdentifier(stage)))
		}
		res = sb.String()
	default:
		log.Panicf("Unknown field %s for Gongstruct ParameterCategory", fieldName)
	}
	return
}

func (parametercategoryuse *ParameterCategoryUse) GongMarshallField(stage *Stage, fieldName string) (res string) {
	ident := parametercategoryuse.GongGetIdentifier(stage)
	_ = ident
	switch fieldName {
	case "Name":
		res = __gong__marshallString(ident, "Name", parametercategoryuse.Name)

	case "ParameterCategory":
		if parametercategoryuse.ParameterCategory != nil {
			res = __gong__marshallPointer(ident, "ParameterCategory", parametercategoryuse.ParameterCategory.GongGetIdentifier(stage))
		} else {
			res = __gong__marshallPointer(ident, "ParameterCategory", "nil")
		}
	default:
		log.Panicf("Unknown field %s for Gongstruct ParameterCategoryUse", fieldName)
	}
	return
}

func (parametershape *ParameterShape) GongMarshallField(stage *Stage, fieldName string) (res string) {
	ident := parametershape.GongGetIdentifier(stage)
	_ = ident
	switch fieldName {
	case "Name":
		res = __gong__marshallString(ident, "Name", parametershape.Name)
	case "Direction":
		res = __gong__marshallEnumString(ident, "Direction", parametershape.Direction.ToCodeString())
	case "ShapeIsComputedFromModel":
		res = __gong__marshallBool(ident, "ShapeIsComputedFromModel", parametershape.ShapeIsComputedFromModel)
	case "X":
		res = __gong__marshallFloat(ident, "X", parametershape.X)
	case "Y":
		res = __gong__marshallFloat(ident, "Y", parametershape.Y)
	case "Width":
		res = __gong__marshallFloat(ident, "Width", parametershape.Width)
	case "Height":
		res = __gong__marshallFloat(ident, "Height", parametershape.Height)
	case "IsHidden":
		res = __gong__marshallBool(ident, "IsHidden", parametershape.IsHidden)
	case "IsPresentationOverridden":
		res = __gong__marshallBool(ident, "IsPresentationOverridden", parametershape.IsPresentationOverridden)
	case "Color":
		res = __gong__marshallString(ident, "Color", parametershape.Color)
	case "FillOpacity":
		res = __gong__marshallFloat(ident, "FillOpacity", parametershape.FillOpacity)
	case "Stroke":
		res = __gong__marshallString(ident, "Stroke", parametershape.Stroke)
	case "StrokeWidth":
		res = __gong__marshallFloat(ident, "StrokeWidth", parametershape.StrokeWidth)
	case "StrokeOpacity":
		res = __gong__marshallFloat(ident, "StrokeOpacity", parametershape.StrokeOpacity)
	case "StrokeDashArray":
		res = __gong__marshallString(ident, "StrokeDashArray", parametershape.StrokeDashArray)
	case "RX":
		res = __gong__marshallFloat(ident, "RX", parametershape.RX)
	case "FontSize":
		res = __gong__marshallString(ident, "FontSize", parametershape.FontSize)
	case "FontWeight":
		res = __gong__marshallString(ident, "FontWeight", parametershape.FontWeight)
	case "FontColor":
		res = __gong__marshallString(ident, "FontColor", parametershape.FontColor)

	case "Parameter":
		if parametershape.Parameter != nil {
			res = __gong__marshallPointer(ident, "Parameter", parametershape.Parameter.GongGetIdentifier(stage))
		} else {
			res = __gong__marshallPointer(ident, "Parameter", "nil")
		}
	default:
		log.Panicf("Unknown field %s for Gongstruct ParameterShape", fieldName)
	}
	return
}

func (parametersaggregate *ParametersAggregate) GongMarshallField(stage *Stage, fieldName string) (res string) {
	ident := parametersaggregate.GongGetIdentifier(stage)
	_ = ident
	switch fieldName {
	case "Name":
		res = __gong__marshallString(ident, "Name", parametersaggregate.Name)
	case "Tag":
		res = __gong__marshallString(ident, "Tag", parametersaggregate.Tag)
	case "Description":
		res = __gong__marshallString(ident, "Description", parametersaggregate.Description)
	case "ComputedPrefix":
		res = __gong__marshallString(ident, "ComputedPrefix", parametersaggregate.ComputedPrefix)
	case "IsExpanded":
		res = __gong__marshallBool(ident, "IsExpanded", parametersaggregate.IsExpanded)

	case "Parameters":
		var sb strings.Builder
		for _, _parameter := range parametersaggregate.Parameters {
			sb.WriteString(__gong__marshallSliceOfPointers(ident, "Parameters", _parameter.GongGetIdentifier(stage)))
		}
		res = sb.String()
	default:
		log.Panicf("Unknown field %s for Gongstruct ParametersAggregate", fieldName)
	}
	return
}

func (parametersaggregateshape *ParametersAggregateShape) GongMarshallField(stage *Stage, fieldName string) (res string) {
	ident := parametersaggregateshape.GongGetIdentifier(stage)
	_ = ident
	switch fieldName {
	case "Name":
		res = __gong__marshallString(ident, "Name", parametersaggregateshape.Name)
	case "Direction":
		res = __gong__marshallEnumString(ident, "Direction", parametersaggregateshape.Direction.ToCodeString())
	case "X":
		res = __gong__marshallFloat(ident, "X", parametersaggregateshape.X)
	case "Y":
		res = __gong__marshallFloat(ident, "Y", parametersaggregateshape.Y)
	case "Width":
		res = __gong__marshallFloat(ident, "Width", parametersaggregateshape.Width)
	case "Height":
		res = __gong__marshallFloat(ident, "Height", parametersaggregateshape.Height)
	case "IsHidden":
		res = __gong__marshallBool(ident, "IsHidden", parametersaggregateshape.IsHidden)
	case "IsPresentationOverridden":
		res = __gong__marshallBool(ident, "IsPresentationOverridden", parametersaggregateshape.IsPresentationOverridden)
	case "Color":
		res = __gong__marshallString(ident, "Color", parametersaggregateshape.Color)
	case "FillOpacity":
		res = __gong__marshallFloat(ident, "FillOpacity", parametersaggregateshape.FillOpacity)
	case "Stroke":
		res = __gong__marshallString(ident, "Stroke", parametersaggregateshape.Stroke)
	case "StrokeWidth":
		res = __gong__marshallFloat(ident, "StrokeWidth", parametersaggregateshape.StrokeWidth)
	case "StrokeOpacity":
		res = __gong__marshallFloat(ident, "StrokeOpacity", parametersaggregateshape.StrokeOpacity)
	case "StrokeDashArray":
		res = __gong__marshallString(ident, "StrokeDashArray", parametersaggregateshape.StrokeDashArray)
	case "RX":
		res = __gong__marshallFloat(ident, "RX", parametersaggregateshape.RX)
	case "FontSize":
		res = __gong__marshallString(ident, "FontSize", parametersaggregateshape.FontSize)
	case "FontWeight":
		res = __gong__marshallString(ident, "FontWeight", parametersaggregateshape.FontWeight)
	case "FontColor":
		res = __gong__marshallString(ident, "FontColor", parametersaggregateshape.FontColor)

	case "ScenarioParameter":
		if parametersaggregateshape.ScenarioParameter != nil {
			res = __gong__marshallPointer(ident, "ScenarioParameter", parametersaggregateshape.ScenarioParameter.GongGetIdentifier(stage))
		} else {
			res = __gong__marshallPointer(ident, "ScenarioParameter", "nil")
		}
	default:
		log.Panicf("Unknown field %s for Gongstruct ParametersAggregateShape", fieldName)
	}
	return
}

func (position *Position) GongMarshallField(stage *Stage, fieldName string) (res string) {
	ident := position.GongGetIdentifier(stage)
	_ = ident
	switch fieldName {
	case "Name":
		res = __gong__marshallString(ident, "Name", position.Name)
	case "Date":
		res = __gong__marshallTime(ident, "Date", position.Date.String())
	case "Ordinate":
		res = __gong__marshallFloat(ident, "Ordinate", position.Ordinate)
	case "ComputedPrefix":
		res = __gong__marshallString(ident, "ComputedPrefix", position.ComputedPrefix)
	case "IsExpanded":
		res = __gong__marshallBool(ident, "IsExpanded", position.IsExpanded)

	default:
		log.Panicf("Unknown field %s for Gongstruct Position", fieldName)
	}
	return
}

func (repository *Repository) GongMarshallField(stage *Stage, fieldName string) (res string) {
	ident := repository.GongGetIdentifier(stage)
	_ = ident
	switch fieldName {
	case "Name":
		res = __gong__marshallString(ident, "Name", repository.Name)
	case "ComputedPrefix":
		res = __gong__marshallString(ident, "ComputedPrefix", repository.ComputedPrefix)
	case "IsExpanded":
		res = __gong__marshallBool(ident, "IsExpanded", repository.IsExpanded)

	case "ParameterUse":
		var sb strings.Builder
		for _, _parametershape := range repository.ParameterUse {
			sb.WriteString(__gong__marshallSliceOfPointers(ident, "ParameterUse", _parametershape.GongGetIdentifier(stage)))
		}
		res = sb.String()
	case "GroupUse":
		var sb strings.Builder
		for _, _groupuse := range repository.GroupUse {
			sb.WriteString(__gong__marshallSliceOfPointers(ident, "GroupUse", _groupuse.GongGetIdentifier(stage)))
		}
		res = sb.String()
	default:
		log.Panicf("Unknown field %s for Gongstruct Repository", fieldName)
	}
	return
}

func (scenario *Scenario) GongMarshallField(stage *Stage, fieldName string) (res string) {
	ident := scenario.GongGetIdentifier(stage)
	_ = ident
	switch fieldName {
	case "Name":
		res = __gong__marshallString(ident, "Name", scenario.Name)
	case "Description":
		res = __gong__marshallString(ident, "Description", scenario.Description)
	case "IsDiagramsNodeExpanded":
		res = __gong__marshallBool(ident, "IsDiagramsNodeExpanded", scenario.IsDiagramsNodeExpanded)
	case "IsActorStatesNodeExpanded":
		res = __gong__marshallBool(ident, "IsActorStatesNodeExpanded", scenario.IsActorStatesNodeExpanded)
	case "IsActorStateTransitionsNodeExpanded":
		res = __gong__marshallBool(ident, "IsActorStateTransitionsNodeExpanded", scenario.IsActorStateTransitionsNodeExpanded)
	case "IsEvolutionDirectionsNodeExpanded":
		res = __gong__marshallBool(ident, "IsEvolutionDirectionsNodeExpanded", scenario.IsEvolutionDirectionsNodeExpanded)
	case "IsParametersNodeExpanded":
		res = __gong__marshallBool(ident, "IsParametersNodeExpanded", scenario.IsParametersNodeExpanded)
	case "IsParametersAggretatesNodeExpanded":
		res = __gong__marshallBool(ident, "IsParametersAggretatesNodeExpanded", scenario.IsParametersAggretatesNodeExpanded)
	case "ComputedPrefix":
		res = __gong__marshallString(ident, "ComputedPrefix", scenario.ComputedPrefix)
	case "IsExpanded":
		res = __gong__marshallBool(ident, "IsExpanded", scenario.IsExpanded)

	case "Diagrams":
		var sb strings.Builder
		for _, _diagram := range scenario.Diagrams {
			sb.WriteString(__gong__marshallSliceOfPointers(ident, "Diagrams", _diagram.GongGetIdentifier(stage)))
		}
		res = sb.String()
	case "ActorStates":
		var sb strings.Builder
		for _, _actorstate := range scenario.ActorStates {
			sb.WriteString(__gong__marshallSliceOfPointers(ident, "ActorStates", _actorstate.GongGetIdentifier(stage)))
		}
		res = sb.String()
	case "ActorStateTransitions":
		var sb strings.Builder
		for _, _actorstatetransition := range scenario.ActorStateTransitions {
			sb.WriteString(__gong__marshallSliceOfPointers(ident, "ActorStateTransitions", _actorstatetransition.GongGetIdentifier(stage)))
		}
		res = sb.String()
	case "EvolutionDirections":
		var sb strings.Builder
		for _, _evolutiondirection := range scenario.EvolutionDirections {
			sb.WriteString(__gong__marshallSliceOfPointers(ident, "EvolutionDirections", _evolutiondirection.GongGetIdentifier(stage)))
		}
		res = sb.String()
	case "Parameters":
		var sb strings.Builder
		for _, _parameter := range scenario.Parameters {
			sb.WriteString(__gong__marshallSliceOfPointers(ident, "Parameters", _parameter.GongGetIdentifier(stage)))
		}
		res = sb.String()
	case "ParametersAggretates":
		var sb strings.Builder
		for _, _parametersaggregate := range scenario.ParametersAggretates {
			sb.WriteString(__gong__marshallSliceOfPointers(ident, "ParametersAggretates", _parametersaggregate.GongGetIdentifier(stage)))
		}
		res = sb.String()
	default:
		log.Panicf("Unknown field %s for Gongstruct Scenario", fieldName)
	}
	return
}

func (user *User) GongMarshallField(stage *Stage, fieldName string) (res string) {
	ident := user.GongGetIdentifier(stage)
	_ = ident
	switch fieldName {
	case "Name":
		res = __gong__marshallString(ident, "Name", user.Name)
	case "ComputedPrefix":
		res = __gong__marshallString(ident, "ComputedPrefix", user.ComputedPrefix)
	case "IsExpanded":
		res = __gong__marshallBool(ident, "IsExpanded", user.IsExpanded)

	default:
		log.Panicf("Unknown field %s for Gongstruct User", fieldName)
	}
	return
}

func (useruse *UserUse) GongMarshallField(stage *Stage, fieldName string) (res string) {
	ident := useruse.GongGetIdentifier(stage)
	_ = ident
	switch fieldName {
	case "Name":
		res = __gong__marshallString(ident, "Name", useruse.Name)

	case "User":
		if useruse.User != nil {
			res = __gong__marshallPointer(ident, "User", useruse.User.GongGetIdentifier(stage))
		} else {
			res = __gong__marshallPointer(ident, "User", "nil")
		}
	default:
		log.Panicf("Unknown field %s for Gongstruct UserUse", fieldName)
	}
	return
}

func (workspace *Workspace) GongMarshallField(stage *Stage, fieldName string) (res string) {
	ident := workspace.GongGetIdentifier(stage)
	_ = ident
	switch fieldName {
	case "Name":
		res = __gong__marshallString(ident, "Name", workspace.Name)
	case "ComputedPrefix":
		res = __gong__marshallString(ident, "ComputedPrefix", workspace.ComputedPrefix)
	case "IsExpanded":
		res = __gong__marshallBool(ident, "IsExpanded", workspace.IsExpanded)

	case "SelectedDiagram":
		if workspace.SelectedDiagram != nil {
			res = __gong__marshallPointer(ident, "SelectedDiagram", workspace.SelectedDiagram.GongGetIdentifier(stage))
		} else {
			res = __gong__marshallPointer(ident, "SelectedDiagram", "nil")
		}
	case "Default_EvolutionDirectionShape":
		if workspace.Default_EvolutionDirectionShape != nil {
			res = __gong__marshallPointer(ident, "Default_EvolutionDirectionShape", workspace.Default_EvolutionDirectionShape.GongGetIdentifier(stage))
		} else {
			res = __gong__marshallPointer(ident, "Default_EvolutionDirectionShape", "nil")
		}
	case "Default_ParameterShape":
		if workspace.Default_ParameterShape != nil {
			res = __gong__marshallPointer(ident, "Default_ParameterShape", workspace.Default_ParameterShape.GongGetIdentifier(stage))
		} else {
			res = __gong__marshallPointer(ident, "Default_ParameterShape", "nil")
		}
	case "Default_ScenarioParameterShape":
		if workspace.Default_ScenarioParameterShape != nil {
			res = __gong__marshallPointer(ident, "Default_ScenarioParameterShape", workspace.Default_ScenarioParameterShape.GongGetIdentifier(stage))
		} else {
			res = __gong__marshallPointer(ident, "Default_ScenarioParameterShape", "nil")
		}
	case "Default_ActorStateShape":
		if workspace.Default_ActorStateShape != nil {
			res = __gong__marshallPointer(ident, "Default_ActorStateShape", workspace.Default_ActorStateShape.GongGetIdentifier(stage))
		} else {
			res = __gong__marshallPointer(ident, "Default_ActorStateShape", "nil")
		}
	case "Default_ActorStateTransitionShape":
		if workspace.Default_ActorStateTransitionShape != nil {
			res = __gong__marshallPointer(ident, "Default_ActorStateTransitionShape", workspace.Default_ActorStateTransitionShape.GongGetIdentifier(stage))
		} else {
			res = __gong__marshallPointer(ident, "Default_ActorStateTransitionShape", "nil")
		}
	default:
		log.Panicf("Unknown field %s for Gongstruct Workspace", fieldName)
	}
	return
}

// insertion point for marshall all fields methods
func (actorstate *ActorState) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(actorstate.GongMarshallField(stage, "Name"))
		initializerStatements.WriteString(actorstate.GongMarshallField(stage, "Description"))
		initializerStatements.WriteString(actorstate.GongMarshallField(stage, "IsWithProbaility"))
		initializerStatements.WriteString(actorstate.GongMarshallField(stage, "Probability"))
		initializerStatements.WriteString(actorstate.GongMarshallField(stage, "ComputedPrefix"))
		initializerStatements.WriteString(actorstate.GongMarshallField(stage, "IsExpanded"))
	}
	initRes = initializerStatements.String()
	ptrRes = pointersInitializesStatements.String()
	return
}
func (actorstateshape *ActorStateShape) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(actorstateshape.GongMarshallField(stage, "Name"))
		pointersInitializesStatements.WriteString(actorstateshape.GongMarshallField(stage, "ActorState"))
		initializerStatements.WriteString(actorstateshape.GongMarshallField(stage, "X"))
		initializerStatements.WriteString(actorstateshape.GongMarshallField(stage, "Y"))
		initializerStatements.WriteString(actorstateshape.GongMarshallField(stage, "Width"))
		initializerStatements.WriteString(actorstateshape.GongMarshallField(stage, "Height"))
		initializerStatements.WriteString(actorstateshape.GongMarshallField(stage, "IsHidden"))
		initializerStatements.WriteString(actorstateshape.GongMarshallField(stage, "IsPresentationOverridden"))
		initializerStatements.WriteString(actorstateshape.GongMarshallField(stage, "Color"))
		initializerStatements.WriteString(actorstateshape.GongMarshallField(stage, "FillOpacity"))
		initializerStatements.WriteString(actorstateshape.GongMarshallField(stage, "Stroke"))
		initializerStatements.WriteString(actorstateshape.GongMarshallField(stage, "StrokeWidth"))
		initializerStatements.WriteString(actorstateshape.GongMarshallField(stage, "StrokeOpacity"))
		initializerStatements.WriteString(actorstateshape.GongMarshallField(stage, "StrokeDashArray"))
		initializerStatements.WriteString(actorstateshape.GongMarshallField(stage, "RX"))
		initializerStatements.WriteString(actorstateshape.GongMarshallField(stage, "FontSize"))
		initializerStatements.WriteString(actorstateshape.GongMarshallField(stage, "FontWeight"))
		initializerStatements.WriteString(actorstateshape.GongMarshallField(stage, "FontColor"))
	}
	initRes = initializerStatements.String()
	ptrRes = pointersInitializesStatements.String()
	return
}
func (actorstatetransition *ActorStateTransition) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(actorstatetransition.GongMarshallField(stage, "Name"))
		pointersInitializesStatements.WriteString(actorstatetransition.GongMarshallField(stage, "StartState"))
		pointersInitializesStatements.WriteString(actorstatetransition.GongMarshallField(stage, "EndState"))
		pointersInitializesStatements.WriteString(actorstatetransition.GongMarshallField(stage, "Justifications"))
		initializerStatements.WriteString(actorstatetransition.GongMarshallField(stage, "ComputedPrefix"))
		initializerStatements.WriteString(actorstatetransition.GongMarshallField(stage, "IsExpanded"))
	}
	initRes = initializerStatements.String()
	ptrRes = pointersInitializesStatements.String()
	return
}
func (actorstatetransitionshape *ActorStateTransitionShape) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(actorstatetransitionshape.GongMarshallField(stage, "Name"))
		pointersInitializesStatements.WriteString(actorstatetransitionshape.GongMarshallField(stage, "ActorStateTransition"))
		pointersInitializesStatements.WriteString(actorstatetransitionshape.GongMarshallField(stage, "Start"))
		pointersInitializesStatements.WriteString(actorstatetransitionshape.GongMarshallField(stage, "End"))
		initializerStatements.WriteString(actorstatetransitionshape.GongMarshallField(stage, "X"))
		initializerStatements.WriteString(actorstatetransitionshape.GongMarshallField(stage, "Y"))
		initializerStatements.WriteString(actorstatetransitionshape.GongMarshallField(stage, "Width"))
		initializerStatements.WriteString(actorstatetransitionshape.GongMarshallField(stage, "Height"))
		initializerStatements.WriteString(actorstatetransitionshape.GongMarshallField(stage, "IsHidden"))
		initializerStatements.WriteString(actorstatetransitionshape.GongMarshallField(stage, "IsPresentationOverridden"))
		initializerStatements.WriteString(actorstatetransitionshape.GongMarshallField(stage, "Color"))
		initializerStatements.WriteString(actorstatetransitionshape.GongMarshallField(stage, "FillOpacity"))
		initializerStatements.WriteString(actorstatetransitionshape.GongMarshallField(stage, "Stroke"))
		initializerStatements.WriteString(actorstatetransitionshape.GongMarshallField(stage, "StrokeWidth"))
		initializerStatements.WriteString(actorstatetransitionshape.GongMarshallField(stage, "StrokeOpacity"))
		initializerStatements.WriteString(actorstatetransitionshape.GongMarshallField(stage, "StrokeDashArray"))
		initializerStatements.WriteString(actorstatetransitionshape.GongMarshallField(stage, "RX"))
		initializerStatements.WriteString(actorstatetransitionshape.GongMarshallField(stage, "FontSize"))
		initializerStatements.WriteString(actorstatetransitionshape.GongMarshallField(stage, "FontWeight"))
		initializerStatements.WriteString(actorstatetransitionshape.GongMarshallField(stage, "FontColor"))
		pointersInitializesStatements.WriteString(actorstatetransitionshape.GongMarshallField(stage, "ControlPointShapes"))
	}
	initRes = initializerStatements.String()
	ptrRes = pointersInitializesStatements.String()
	return
}
func (analysis *Analysis) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(analysis.GongMarshallField(stage, "Name"))
		initializerStatements.WriteString(analysis.GongMarshallField(stage, "Description"))
		pointersInitializesStatements.WriteString(analysis.GongMarshallField(stage, "Scenarios"))
		initializerStatements.WriteString(analysis.GongMarshallField(stage, "IsScenariosNodeExpanded"))
		pointersInitializesStatements.WriteString(analysis.GongMarshallField(stage, "GroupUse"))
		initializerStatements.WriteString(analysis.GongMarshallField(stage, "IsGroupUseNodeExpanded"))
		pointersInitializesStatements.WriteString(analysis.GongMarshallField(stage, "GeoObjectUse"))
		initializerStatements.WriteString(analysis.GongMarshallField(stage, "IsGeoObjectUseNodeExpanded"))
		pointersInitializesStatements.WriteString(analysis.GongMarshallField(stage, "MapUse"))
		initializerStatements.WriteString(analysis.GongMarshallField(stage, "IsMapUseNodeExpanded"))
		initializerStatements.WriteString(analysis.GongMarshallField(stage, "ComputedPrefix"))
		initializerStatements.WriteString(analysis.GongMarshallField(stage, "IsExpanded"))
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
func (diagram *Diagram) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(diagram.GongMarshallField(stage, "Name"))
		initializerStatements.WriteString(diagram.GongMarshallField(stage, "ComputedPrefix"))
		initializerStatements.WriteString(diagram.GongMarshallField(stage, "IsExpanded"))
		initializerStatements.WriteString(diagram.GongMarshallField(stage, "IsChecked"))
		initializerStatements.WriteString(diagram.GongMarshallField(stage, "IsShowPrefix"))
		initializerStatements.WriteString(diagram.GongMarshallField(stage, "Description"))
		pointersInitializesStatements.WriteString(diagram.GongMarshallField(stage, "EvolutionDirectionShapes"))
		pointersInitializesStatements.WriteString(diagram.GongMarshallField(stage, "EvolutionDirectionsWhoseNodeIsExpanded"))
		initializerStatements.WriteString(diagram.GongMarshallField(stage, "IsEvolutionDirectionsNodeExpanded"))
		pointersInitializesStatements.WriteString(diagram.GongMarshallField(stage, "ActorStateShapes"))
		pointersInitializesStatements.WriteString(diagram.GongMarshallField(stage, "ActorStatesWhoseNodeIsExpanded"))
		initializerStatements.WriteString(diagram.GongMarshallField(stage, "IsActorStatesNodeExpanded"))
		pointersInitializesStatements.WriteString(diagram.GongMarshallField(stage, "ParameterShapes"))
		pointersInitializesStatements.WriteString(diagram.GongMarshallField(stage, "ParametersWhoseNodeIsExpanded"))
		initializerStatements.WriteString(diagram.GongMarshallField(stage, "IsParametersNodeExpanded"))
		pointersInitializesStatements.WriteString(diagram.GongMarshallField(stage, "ScenarioParameterShapes"))
		pointersInitializesStatements.WriteString(diagram.GongMarshallField(stage, "ParametersAggregatesWhoseNodeIsExpanded"))
		initializerStatements.WriteString(diagram.GongMarshallField(stage, "IsParametersAggregatesNodeExpanded"))
		pointersInitializesStatements.WriteString(diagram.GongMarshallField(stage, "ActorStateTransitionShapes"))
		pointersInitializesStatements.WriteString(diagram.GongMarshallField(stage, "ActorStateTransitionsWhoseNodeIsExpanded"))
		initializerStatements.WriteString(diagram.GongMarshallField(stage, "IsActorStateTransitionsNodeExpanded"))
		initializerStatements.WriteString(diagram.GongMarshallField(stage, "AxisOrign_X"))
		initializerStatements.WriteString(diagram.GongMarshallField(stage, "AxisOrign_Y"))
		initializerStatements.WriteString(diagram.GongMarshallField(stage, "VerticalAxis_Top_Y"))
		initializerStatements.WriteString(diagram.GongMarshallField(stage, "VerticalAxis_Bottom_Y"))
		initializerStatements.WriteString(diagram.GongMarshallField(stage, "VerticalAxis_StrokeWidth"))
		initializerStatements.WriteString(diagram.GongMarshallField(stage, "HorizontalAxis_Right_X"))
		initializerStatements.WriteString(diagram.GongMarshallField(stage, "Start"))
		initializerStatements.WriteString(diagram.GongMarshallField(stage, "End"))
		initializerStatements.WriteString(diagram.GongMarshallField(stage, "NumberOfYearsBetweenTicks"))
	}
	initRes = initializerStatements.String()
	ptrRes = pointersInitializesStatements.String()
	return
}
func (document *Document) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(document.GongMarshallField(stage, "Name"))
		pointersInitializesStatements.WriteString(document.GongMarshallField(stage, "GeoObjectUse"))
		initializerStatements.WriteString(document.GongMarshallField(stage, "ComputedPrefix"))
		initializerStatements.WriteString(document.GongMarshallField(stage, "IsExpanded"))
	}
	initRes = initializerStatements.String()
	ptrRes = pointersInitializesStatements.String()
	return
}
func (documentuse *DocumentUse) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(documentuse.GongMarshallField(stage, "Name"))
		pointersInitializesStatements.WriteString(documentuse.GongMarshallField(stage, "Document"))
	}
	initRes = initializerStatements.String()
	ptrRes = pointersInitializesStatements.String()
	return
}
func (evolutiondirection *EvolutionDirection) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(evolutiondirection.GongMarshallField(stage, "Name"))
		initializerStatements.WriteString(evolutiondirection.GongMarshallField(stage, "Description"))
		initializerStatements.WriteString(evolutiondirection.GongMarshallField(stage, "ComputedPrefix"))
		initializerStatements.WriteString(evolutiondirection.GongMarshallField(stage, "IsExpanded"))
	}
	initRes = initializerStatements.String()
	ptrRes = pointersInitializesStatements.String()
	return
}
func (evolutiondirectionshape *EvolutionDirectionShape) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(evolutiondirectionshape.GongMarshallField(stage, "Name"))
		pointersInitializesStatements.WriteString(evolutiondirectionshape.GongMarshallField(stage, "EvolutionDirection"))
		initializerStatements.WriteString(evolutiondirectionshape.GongMarshallField(stage, "X"))
		initializerStatements.WriteString(evolutiondirectionshape.GongMarshallField(stage, "Y"))
		initializerStatements.WriteString(evolutiondirectionshape.GongMarshallField(stage, "Width"))
		initializerStatements.WriteString(evolutiondirectionshape.GongMarshallField(stage, "Height"))
		initializerStatements.WriteString(evolutiondirectionshape.GongMarshallField(stage, "IsHidden"))
		initializerStatements.WriteString(evolutiondirectionshape.GongMarshallField(stage, "IsPresentationOverridden"))
		initializerStatements.WriteString(evolutiondirectionshape.GongMarshallField(stage, "Color"))
		initializerStatements.WriteString(evolutiondirectionshape.GongMarshallField(stage, "FillOpacity"))
		initializerStatements.WriteString(evolutiondirectionshape.GongMarshallField(stage, "Stroke"))
		initializerStatements.WriteString(evolutiondirectionshape.GongMarshallField(stage, "StrokeWidth"))
		initializerStatements.WriteString(evolutiondirectionshape.GongMarshallField(stage, "StrokeOpacity"))
		initializerStatements.WriteString(evolutiondirectionshape.GongMarshallField(stage, "StrokeDashArray"))
		initializerStatements.WriteString(evolutiondirectionshape.GongMarshallField(stage, "RX"))
		initializerStatements.WriteString(evolutiondirectionshape.GongMarshallField(stage, "FontSize"))
		initializerStatements.WriteString(evolutiondirectionshape.GongMarshallField(stage, "FontWeight"))
		initializerStatements.WriteString(evolutiondirectionshape.GongMarshallField(stage, "FontColor"))
	}
	initRes = initializerStatements.String()
	ptrRes = pointersInitializesStatements.String()
	return
}
func (foo *Foo) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(foo.GongMarshallField(stage, "Name"))
	}
	initRes = initializerStatements.String()
	ptrRes = pointersInitializesStatements.String()
	return
}
func (geoobject *GeoObject) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(geoobject.GongMarshallField(stage, "Name"))
		initializerStatements.WriteString(geoobject.GongMarshallField(stage, "ComputedPrefix"))
		initializerStatements.WriteString(geoobject.GongMarshallField(stage, "IsExpanded"))
	}
	initRes = initializerStatements.String()
	ptrRes = pointersInitializesStatements.String()
	return
}
func (geoobjectuse *GeoObjectUse) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(geoobjectuse.GongMarshallField(stage, "Name"))
		pointersInitializesStatements.WriteString(geoobjectuse.GongMarshallField(stage, "GeoObject"))
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
		pointersInitializesStatements.WriteString(group.GongMarshallField(stage, "UserUse"))
		initializerStatements.WriteString(group.GongMarshallField(stage, "ComputedPrefix"))
		initializerStatements.WriteString(group.GongMarshallField(stage, "IsExpanded"))
	}
	initRes = initializerStatements.String()
	ptrRes = pointersInitializesStatements.String()
	return
}
func (groupuse *GroupUse) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(groupuse.GongMarshallField(stage, "Name"))
		pointersInitializesStatements.WriteString(groupuse.GongMarshallField(stage, "Group"))
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
		pointersInitializesStatements.WriteString(library.GongMarshallField(stage, "Analyses"))
		initializerStatements.WriteString(library.GongMarshallField(stage, "IsAnalysesNodeExpanded"))
		pointersInitializesStatements.WriteString(library.GongMarshallField(stage, "SubLibraries"))
		initializerStatements.WriteString(library.GongMarshallField(stage, "IsSubLibrariesNodeExpanded"))
		pointersInitializesStatements.WriteString(library.GongMarshallField(stage, "SubLibrariesWhoseNodeIsExpanded"))
		initializerStatements.WriteString(library.GongMarshallField(stage, "NbPixPerCharacter"))
		initializerStatements.WriteString(library.GongMarshallField(stage, "LogoSVGFile"))
		initializerStatements.WriteString(library.GongMarshallField(stage, "IsExpandedTmp"))
	}
	initRes = initializerStatements.String()
	ptrRes = pointersInitializesStatements.String()
	return
}
func (mapobject *MapObject) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(mapobject.GongMarshallField(stage, "Name"))
		initializerStatements.WriteString(mapobject.GongMarshallField(stage, "ComputedPrefix"))
		initializerStatements.WriteString(mapobject.GongMarshallField(stage, "IsExpanded"))
	}
	initRes = initializerStatements.String()
	ptrRes = pointersInitializesStatements.String()
	return
}
func (mapobjectuse *MapObjectUse) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(mapobjectuse.GongMarshallField(stage, "Name"))
		pointersInitializesStatements.WriteString(mapobjectuse.GongMarshallField(stage, "Map"))
	}
	initRes = initializerStatements.String()
	ptrRes = pointersInitializesStatements.String()
	return
}
func (parameter *Parameter) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(parameter.GongMarshallField(stage, "Name"))
		initializerStatements.WriteString(parameter.GongMarshallField(stage, "Description"))
		initializerStatements.WriteString(parameter.GongMarshallField(stage, "IsResponse"))
		initializerStatements.WriteString(parameter.GongMarshallField(stage, "Start"))
		initializerStatements.WriteString(parameter.GongMarshallField(stage, "End"))
		initializerStatements.WriteString(parameter.GongMarshallField(stage, "Force"))
		pointersInitializesStatements.WriteString(parameter.GongMarshallField(stage, "GroupUse"))
		pointersInitializesStatements.WriteString(parameter.GongMarshallField(stage, "DocumentUse"))
		pointersInitializesStatements.WriteString(parameter.GongMarshallField(stage, "GeoObjectUse"))
		initializerStatements.WriteString(parameter.GongMarshallField(stage, "Tag"))
		initializerStatements.WriteString(parameter.GongMarshallField(stage, "ComputedPrefix"))
		initializerStatements.WriteString(parameter.GongMarshallField(stage, "IsExpanded"))
	}
	initRes = initializerStatements.String()
	ptrRes = pointersInitializesStatements.String()
	return
}
func (parametercategory *ParameterCategory) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(parametercategory.GongMarshallField(stage, "Name"))
		pointersInitializesStatements.WriteString(parametercategory.GongMarshallField(stage, "ParameterUse"))
		initializerStatements.WriteString(parametercategory.GongMarshallField(stage, "ComputedPrefix"))
		initializerStatements.WriteString(parametercategory.GongMarshallField(stage, "IsExpanded"))
	}
	initRes = initializerStatements.String()
	ptrRes = pointersInitializesStatements.String()
	return
}
func (parametercategoryuse *ParameterCategoryUse) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(parametercategoryuse.GongMarshallField(stage, "Name"))
		pointersInitializesStatements.WriteString(parametercategoryuse.GongMarshallField(stage, "ParameterCategory"))
	}
	initRes = initializerStatements.String()
	ptrRes = pointersInitializesStatements.String()
	return
}
func (parametershape *ParameterShape) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(parametershape.GongMarshallField(stage, "Name"))
		pointersInitializesStatements.WriteString(parametershape.GongMarshallField(stage, "Parameter"))
		initializerStatements.WriteString(parametershape.GongMarshallField(stage, "Direction"))
		initializerStatements.WriteString(parametershape.GongMarshallField(stage, "ShapeIsComputedFromModel"))
		initializerStatements.WriteString(parametershape.GongMarshallField(stage, "X"))
		initializerStatements.WriteString(parametershape.GongMarshallField(stage, "Y"))
		initializerStatements.WriteString(parametershape.GongMarshallField(stage, "Width"))
		initializerStatements.WriteString(parametershape.GongMarshallField(stage, "Height"))
		initializerStatements.WriteString(parametershape.GongMarshallField(stage, "IsHidden"))
		initializerStatements.WriteString(parametershape.GongMarshallField(stage, "IsPresentationOverridden"))
		initializerStatements.WriteString(parametershape.GongMarshallField(stage, "Color"))
		initializerStatements.WriteString(parametershape.GongMarshallField(stage, "FillOpacity"))
		initializerStatements.WriteString(parametershape.GongMarshallField(stage, "Stroke"))
		initializerStatements.WriteString(parametershape.GongMarshallField(stage, "StrokeWidth"))
		initializerStatements.WriteString(parametershape.GongMarshallField(stage, "StrokeOpacity"))
		initializerStatements.WriteString(parametershape.GongMarshallField(stage, "StrokeDashArray"))
		initializerStatements.WriteString(parametershape.GongMarshallField(stage, "RX"))
		initializerStatements.WriteString(parametershape.GongMarshallField(stage, "FontSize"))
		initializerStatements.WriteString(parametershape.GongMarshallField(stage, "FontWeight"))
		initializerStatements.WriteString(parametershape.GongMarshallField(stage, "FontColor"))
	}
	initRes = initializerStatements.String()
	ptrRes = pointersInitializesStatements.String()
	return
}
func (parametersaggregate *ParametersAggregate) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(parametersaggregate.GongMarshallField(stage, "Name"))
		initializerStatements.WriteString(parametersaggregate.GongMarshallField(stage, "Tag"))
		initializerStatements.WriteString(parametersaggregate.GongMarshallField(stage, "Description"))
		pointersInitializesStatements.WriteString(parametersaggregate.GongMarshallField(stage, "Parameters"))
		initializerStatements.WriteString(parametersaggregate.GongMarshallField(stage, "ComputedPrefix"))
		initializerStatements.WriteString(parametersaggregate.GongMarshallField(stage, "IsExpanded"))
	}
	initRes = initializerStatements.String()
	ptrRes = pointersInitializesStatements.String()
	return
}
func (parametersaggregateshape *ParametersAggregateShape) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(parametersaggregateshape.GongMarshallField(stage, "Name"))
		pointersInitializesStatements.WriteString(parametersaggregateshape.GongMarshallField(stage, "ScenarioParameter"))
		initializerStatements.WriteString(parametersaggregateshape.GongMarshallField(stage, "Direction"))
		initializerStatements.WriteString(parametersaggregateshape.GongMarshallField(stage, "X"))
		initializerStatements.WriteString(parametersaggregateshape.GongMarshallField(stage, "Y"))
		initializerStatements.WriteString(parametersaggregateshape.GongMarshallField(stage, "Width"))
		initializerStatements.WriteString(parametersaggregateshape.GongMarshallField(stage, "Height"))
		initializerStatements.WriteString(parametersaggregateshape.GongMarshallField(stage, "IsHidden"))
		initializerStatements.WriteString(parametersaggregateshape.GongMarshallField(stage, "IsPresentationOverridden"))
		initializerStatements.WriteString(parametersaggregateshape.GongMarshallField(stage, "Color"))
		initializerStatements.WriteString(parametersaggregateshape.GongMarshallField(stage, "FillOpacity"))
		initializerStatements.WriteString(parametersaggregateshape.GongMarshallField(stage, "Stroke"))
		initializerStatements.WriteString(parametersaggregateshape.GongMarshallField(stage, "StrokeWidth"))
		initializerStatements.WriteString(parametersaggregateshape.GongMarshallField(stage, "StrokeOpacity"))
		initializerStatements.WriteString(parametersaggregateshape.GongMarshallField(stage, "StrokeDashArray"))
		initializerStatements.WriteString(parametersaggregateshape.GongMarshallField(stage, "RX"))
		initializerStatements.WriteString(parametersaggregateshape.GongMarshallField(stage, "FontSize"))
		initializerStatements.WriteString(parametersaggregateshape.GongMarshallField(stage, "FontWeight"))
		initializerStatements.WriteString(parametersaggregateshape.GongMarshallField(stage, "FontColor"))
	}
	initRes = initializerStatements.String()
	ptrRes = pointersInitializesStatements.String()
	return
}
func (position *Position) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(position.GongMarshallField(stage, "Name"))
		initializerStatements.WriteString(position.GongMarshallField(stage, "Date"))
		initializerStatements.WriteString(position.GongMarshallField(stage, "Ordinate"))
		initializerStatements.WriteString(position.GongMarshallField(stage, "ComputedPrefix"))
		initializerStatements.WriteString(position.GongMarshallField(stage, "IsExpanded"))
	}
	initRes = initializerStatements.String()
	ptrRes = pointersInitializesStatements.String()
	return
}
func (repository *Repository) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(repository.GongMarshallField(stage, "Name"))
		pointersInitializesStatements.WriteString(repository.GongMarshallField(stage, "ParameterUse"))
		pointersInitializesStatements.WriteString(repository.GongMarshallField(stage, "GroupUse"))
		initializerStatements.WriteString(repository.GongMarshallField(stage, "ComputedPrefix"))
		initializerStatements.WriteString(repository.GongMarshallField(stage, "IsExpanded"))
	}
	initRes = initializerStatements.String()
	ptrRes = pointersInitializesStatements.String()
	return
}
func (scenario *Scenario) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(scenario.GongMarshallField(stage, "Name"))
		initializerStatements.WriteString(scenario.GongMarshallField(stage, "Description"))
		pointersInitializesStatements.WriteString(scenario.GongMarshallField(stage, "Diagrams"))
		initializerStatements.WriteString(scenario.GongMarshallField(stage, "IsDiagramsNodeExpanded"))
		pointersInitializesStatements.WriteString(scenario.GongMarshallField(stage, "ActorStates"))
		initializerStatements.WriteString(scenario.GongMarshallField(stage, "IsActorStatesNodeExpanded"))
		pointersInitializesStatements.WriteString(scenario.GongMarshallField(stage, "ActorStateTransitions"))
		initializerStatements.WriteString(scenario.GongMarshallField(stage, "IsActorStateTransitionsNodeExpanded"))
		pointersInitializesStatements.WriteString(scenario.GongMarshallField(stage, "EvolutionDirections"))
		initializerStatements.WriteString(scenario.GongMarshallField(stage, "IsEvolutionDirectionsNodeExpanded"))
		pointersInitializesStatements.WriteString(scenario.GongMarshallField(stage, "Parameters"))
		initializerStatements.WriteString(scenario.GongMarshallField(stage, "IsParametersNodeExpanded"))
		pointersInitializesStatements.WriteString(scenario.GongMarshallField(stage, "ParametersAggretates"))
		initializerStatements.WriteString(scenario.GongMarshallField(stage, "IsParametersAggretatesNodeExpanded"))
		initializerStatements.WriteString(scenario.GongMarshallField(stage, "ComputedPrefix"))
		initializerStatements.WriteString(scenario.GongMarshallField(stage, "IsExpanded"))
	}
	initRes = initializerStatements.String()
	ptrRes = pointersInitializesStatements.String()
	return
}
func (user *User) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(user.GongMarshallField(stage, "Name"))
		initializerStatements.WriteString(user.GongMarshallField(stage, "ComputedPrefix"))
		initializerStatements.WriteString(user.GongMarshallField(stage, "IsExpanded"))
	}
	initRes = initializerStatements.String()
	ptrRes = pointersInitializesStatements.String()
	return
}
func (useruse *UserUse) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(useruse.GongMarshallField(stage, "Name"))
		pointersInitializesStatements.WriteString(useruse.GongMarshallField(stage, "User"))
	}
	initRes = initializerStatements.String()
	ptrRes = pointersInitializesStatements.String()
	return
}
func (workspace *Workspace) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(workspace.GongMarshallField(stage, "Name"))
		pointersInitializesStatements.WriteString(workspace.GongMarshallField(stage, "SelectedDiagram"))
		pointersInitializesStatements.WriteString(workspace.GongMarshallField(stage, "Default_EvolutionDirectionShape"))
		pointersInitializesStatements.WriteString(workspace.GongMarshallField(stage, "Default_ParameterShape"))
		pointersInitializesStatements.WriteString(workspace.GongMarshallField(stage, "Default_ScenarioParameterShape"))
		pointersInitializesStatements.WriteString(workspace.GongMarshallField(stage, "Default_ActorStateShape"))
		pointersInitializesStatements.WriteString(workspace.GongMarshallField(stage, "Default_ActorStateTransitionShape"))
		initializerStatements.WriteString(workspace.GongMarshallField(stage, "ComputedPrefix"))
		initializerStatements.WriteString(workspace.GongMarshallField(stage, "IsExpanded"))
	}
	initRes = initializerStatements.String()
	ptrRes = pointersInitializesStatements.String()
	return
}
