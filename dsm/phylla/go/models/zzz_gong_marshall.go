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
	modelsPackageName := "github.com/fullstack-lang/gong/dsm/phylla/go/models"
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

	if stage.ClockStage != nil {
		for _, clockabstract := range __gong__sortStageSetInstances(stage.ClockStage.ClockAbstracts, stage.ClockStage.ClockAbstract_stagedOrder) {
			if lastStageDecl != "ClockStage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "ClockStage"
			}
			clockabstractIdent := "__clock" + clockabstract.GongGetIdentifier(stage.ClockStage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&clock.ClockAbstract{Name: %s}).Stage(stage.ClockStage)", clockabstractIdent, __gong__toRawStringLiteral(clockabstract.Name)))
			if lastStageVal != "ClockStage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "ClockStage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", clockabstractIdent, __gong__toRawStringLiteral(clockabstract.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.RadialRepetitions = %d", clockabstractIdent, clockabstract.RadialRepetitions))
			values.WriteString(fmt.Sprintf("\n\t%s.Transparency = %f", clockabstractIdent, clockabstract.Transparency))
			values.WriteString(fmt.Sprintf("\n\t%s.RelativeTubeDiameter = %f", clockabstractIdent, clockabstract.RelativeTubeDiameter))
			values.WriteString(fmt.Sprintf("\n\t%s.RelativeHeight3DTorus = %f", clockabstractIdent, clockabstract.RelativeHeight3DTorus))
			values.WriteString(fmt.Sprintf("\n\t%s.ClockTorusVerticalScale = %f", clockabstractIdent, clockabstract.ClockTorusVerticalScale))
			values.WriteString(fmt.Sprintf("\n\t%s.RelativeHeight = %f", clockabstractIdent, clockabstract.RelativeHeight))
			values.WriteString(fmt.Sprintf("\n\t%s.ProjectionAngle = %f", clockabstractIdent, clockabstract.ProjectionAngle))
		}
	}
	if stage.MusicStage != nil {
		for _, musicabstract := range __gong__sortStageSetInstances(stage.MusicStage.MusicAbstracts, stage.MusicStage.MusicAbstract_stagedOrder) {
			if lastStageDecl != "MusicStage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "MusicStage"
			}
			musicabstractIdent := "__music" + musicabstract.GongGetIdentifier(stage.MusicStage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&music.MusicAbstract{Name: %s}).Stage(stage.MusicStage)", musicabstractIdent, __gong__toRawStringLiteral(musicabstract.Name)))
			if lastStageVal != "MusicStage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "MusicStage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", musicabstractIdent, __gong__toRawStringLiteral(musicabstract.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.IsChecked = %t", musicabstractIdent, musicabstract.IsChecked))
			values.WriteString(fmt.Sprintf("\n\t%s.PitchHeight = %f", musicabstractIdent, musicabstract.PitchHeight))
			values.WriteString(fmt.Sprintf("\n\t%s.NbOfBeatsInTheme = %d", musicabstractIdent, musicabstract.NbOfBeatsInTheme))
			values.WriteString(fmt.Sprintf("\n\t%s.BeatsPerSecond = %f", musicabstractIdent, musicabstract.BeatsPerSecond))
			values.WriteString(fmt.Sprintf("\n\t%s.FirstVoiceShiftX = %f", musicabstractIdent, musicabstract.FirstVoiceShiftX))
			values.WriteString(fmt.Sprintf("\n\t%s.FirstVoiceShiftY = %f", musicabstractIdent, musicabstract.FirstVoiceShiftY))
			values.WriteString(fmt.Sprintf("\n\t%s.PitchDifference = %d", musicabstractIdent, musicabstract.PitchDifference))
			values.WriteString(fmt.Sprintf("\n\t%s.Level = %f", musicabstractIdent, musicabstract.Level))
			values.WriteString(fmt.Sprintf("\n\t%s.ActualBeatsTemporalShift = %d", musicabstractIdent, musicabstract.ActualBeatsTemporalShift))
			values.WriteString(fmt.Sprintf("\n\t%s.IsMinor = %t", musicabstractIdent, musicabstract.IsMinor))
			values.WriteString(fmt.Sprintf("\n\t%s.ThemeBinaryEncoding = %d", musicabstractIdent, musicabstract.ThemeBinaryEncoding))
			values.WriteString(fmt.Sprintf("\n\t%s.BezierControlLengthRatio = %f", musicabstractIdent, musicabstract.BezierControlLengthRatio))
			values.WriteString(fmt.Sprintf("\n\t%s.NbPitchLines = %d", musicabstractIdent, musicabstract.NbPitchLines))
			values.WriteString(fmt.Sprintf("\n\t%s.NbBeatLines = %d", musicabstractIdent, musicabstract.NbBeatLines))
			values.WriteString(fmt.Sprintf("\n\t%s.OriginX = %f", musicabstractIdent, musicabstract.OriginX))
			values.WriteString(fmt.Sprintf("\n\t%s.OriginY = %f", musicabstractIdent, musicabstract.OriginY))
			values.WriteString(fmt.Sprintf("\n\t%s.ScoreScale = %f", musicabstractIdent, musicabstract.ScoreScale))
			values.WriteString(fmt.Sprintf("\n\t%s.ShowFirstVoice = %t", musicabstractIdent, musicabstract.ShowFirstVoice))
			values.WriteString(fmt.Sprintf("\n\t%s.ShowFirstVoiceShiftRight = %t", musicabstractIdent, musicabstract.ShowFirstVoiceShiftRight))
			values.WriteString(fmt.Sprintf("\n\t%s.ShowSecondVoice = %t", musicabstractIdent, musicabstract.ShowSecondVoice))
			values.WriteString(fmt.Sprintf("\n\t%s.ShowSecondVoiceShiftRight = %t", musicabstractIdent, musicabstract.ShowSecondVoiceShiftRight))
			values.WriteString(fmt.Sprintf("\n\t%s.ShowFirstVoiceNotes = %t", musicabstractIdent, musicabstract.ShowFirstVoiceNotes))
			values.WriteString(fmt.Sprintf("\n\t%s.ShowFirstVoiceNotesShiftRight = %t", musicabstractIdent, musicabstract.ShowFirstVoiceNotesShiftRight))
			values.WriteString(fmt.Sprintf("\n\t%s.ShowSecondVoiceNotes = %t", musicabstractIdent, musicabstract.ShowSecondVoiceNotes))
			values.WriteString(fmt.Sprintf("\n\t%s.ShowSecondVoiceNotesShiftRight = %t", musicabstractIdent, musicabstract.ShowSecondVoiceNotesShiftRight))
			values.WriteString(fmt.Sprintf("\n\t%s.IsComposerNodeExpanded = %t", musicabstractIdent, musicabstract.IsComposerNodeExpanded))
		}
	}
	if stage.StoolStage != nil {
		for _, stoolabstract := range __gong__sortStageSetInstances(stage.StoolStage.StoolAbstracts, stage.StoolStage.StoolAbstract_stagedOrder) {
			if lastStageDecl != "StoolStage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "StoolStage"
			}
			stoolabstractIdent := "__stool" + stoolabstract.GongGetIdentifier(stage.StoolStage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&stool.StoolAbstract{Name: %s}).Stage(stage.StoolStage)", stoolabstractIdent, __gong__toRawStringLiteral(stoolabstract.Name)))
			if lastStageVal != "StoolStage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "StoolStage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", stoolabstractIdent, __gong__toRawStringLiteral(stoolabstract.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.RadialRepetitions = %d", stoolabstractIdent, stoolabstract.RadialRepetitions))
			values.WriteString(fmt.Sprintf("\n\t%s.Transparency = %f", stoolabstractIdent, stoolabstract.Transparency))
			values.WriteString(fmt.Sprintf("\n\t%s.RelativeTubeDiameter = %f", stoolabstractIdent, stoolabstract.RelativeTubeDiameter))
			values.WriteString(fmt.Sprintf("\n\t%s.RelativeHeight3DTorus = %f", stoolabstractIdent, stoolabstract.RelativeHeight3DTorus))
			values.WriteString(fmt.Sprintf("\n\t%s.StoolTorusVerticalScale = %f", stoolabstractIdent, stoolabstract.StoolTorusVerticalScale))
			values.WriteString(fmt.Sprintf("\n\t%s.RelativeHeight = %f", stoolabstractIdent, stoolabstract.RelativeHeight))
			values.WriteString(fmt.Sprintf("\n\t%s.RelativeSeatThickness = %f", stoolabstractIdent, stoolabstract.RelativeSeatThickness))
			values.WriteString(fmt.Sprintf("\n\t%s.ProjectionAngle = %f", stoolabstractIdent, stoolabstract.ProjectionAngle))
			values.WriteString(fmt.Sprintf("\n\t%s.RelativeEyeSeparationCriteria = %f", stoolabstractIdent, stoolabstract.RelativeEyeSeparationCriteria))
			values.WriteString(fmt.Sprintf("\n\t%s.RelativeEyeCornerControlVectorStrength = %f", stoolabstractIdent, stoolabstract.RelativeEyeCornerControlVectorStrength))
		}
	}
	if stage != nil {
		for _, angle0shape := range __gong__sortStageSetInstances(stage.Angle0Shapes, stage.Angle0Shape_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			angle0shapeIdent := localAliasPrefix + angle0shape.GongGetIdentifier(stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&models.Angle0Shape{Name: %s}).Stage(stage)", angle0shapeIdent, __gong__toRawStringLiteral(angle0shape.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", angle0shapeIdent, __gong__toRawStringLiteral(angle0shape.Name)))
		}
	}
	if stage != nil {
		for _, bottomcurveplane1shape := range __gong__sortStageSetInstances(stage.BottomCurvePlane1Shapes, stage.BottomCurvePlane1Shape_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			bottomcurveplane1shapeIdent := localAliasPrefix + bottomcurveplane1shape.GongGetIdentifier(stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&models.BottomCurvePlane1Shape{Name: %s}).Stage(stage)", bottomcurveplane1shapeIdent, __gong__toRawStringLiteral(bottomcurveplane1shape.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", bottomcurveplane1shapeIdent, __gong__toRawStringLiteral(bottomcurveplane1shape.Name)))
		}
	}
	if stage != nil {
		for _, bottomcurveplane2shape := range __gong__sortStageSetInstances(stage.BottomCurvePlane2Shapes, stage.BottomCurvePlane2Shape_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			bottomcurveplane2shapeIdent := localAliasPrefix + bottomcurveplane2shape.GongGetIdentifier(stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&models.BottomCurvePlane2Shape{Name: %s}).Stage(stage)", bottomcurveplane2shapeIdent, __gong__toRawStringLiteral(bottomcurveplane2shape.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", bottomcurveplane2shapeIdent, __gong__toRawStringLiteral(bottomcurveplane2shape.Name)))
		}
	}
	if stage != nil {
		for _, carvedoutbottomcurveplane1shape := range __gong__sortStageSetInstances(stage.CarvedOutBottomCurvePlane1Shapes, stage.CarvedOutBottomCurvePlane1Shape_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			carvedoutbottomcurveplane1shapeIdent := localAliasPrefix + carvedoutbottomcurveplane1shape.GongGetIdentifier(stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&models.CarvedOutBottomCurvePlane1Shape{Name: %s}).Stage(stage)", carvedoutbottomcurveplane1shapeIdent, __gong__toRawStringLiteral(carvedoutbottomcurveplane1shape.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", carvedoutbottomcurveplane1shapeIdent, __gong__toRawStringLiteral(carvedoutbottomcurveplane1shape.Name)))
		}
	}
	if stage != nil {
		for _, carvedouttopcurveplane1shape := range __gong__sortStageSetInstances(stage.CarvedOutTopCurvePlane1Shapes, stage.CarvedOutTopCurvePlane1Shape_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			carvedouttopcurveplane1shapeIdent := localAliasPrefix + carvedouttopcurveplane1shape.GongGetIdentifier(stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&models.CarvedOutTopCurvePlane1Shape{Name: %s}).Stage(stage)", carvedouttopcurveplane1shapeIdent, __gong__toRawStringLiteral(carvedouttopcurveplane1shape.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", carvedouttopcurveplane1shapeIdent, __gong__toRawStringLiteral(carvedouttopcurveplane1shape.Name)))
		}
	}
	if stage != nil {
		for _, carvedoutvasetrapezeringshape := range __gong__sortStageSetInstances(stage.CarvedOutVaseTrapezeRingShapes, stage.CarvedOutVaseTrapezeRingShape_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			carvedoutvasetrapezeringshapeIdent := localAliasPrefix + carvedoutvasetrapezeringshape.GongGetIdentifier(stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&models.CarvedOutVaseTrapezeRingShape{Name: %s}).Stage(stage)", carvedoutvasetrapezeringshapeIdent, __gong__toRawStringLiteral(carvedoutvasetrapezeringshape.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", carvedoutvasetrapezeringshapeIdent, __gong__toRawStringLiteral(carvedoutvasetrapezeringshape.Name)))
		}
	}
	if stage != nil {
		for _, circumference3dshape := range __gong__sortStageSetInstances(stage.Circumference3DShapes, stage.Circumference3DShape_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			circumference3dshapeIdent := localAliasPrefix + circumference3dshape.GongGetIdentifier(stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&models.Circumference3DShape{Name: %s}).Stage(stage)", circumference3dshapeIdent, __gong__toRawStringLiteral(circumference3dshape.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", circumference3dshapeIdent, __gong__toRawStringLiteral(circumference3dshape.Name)))
		}
	}
	if stage != nil {
		for _, clock2ddiagram := range __gong__sortStageSetInstances(stage.Clock2DDiagrams, stage.Clock2DDiagram_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			clock2ddiagramIdent := localAliasPrefix + clock2ddiagram.GongGetIdentifier(stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&models.Clock2DDiagram{Name: %s}).Stage(stage)", clock2ddiagramIdent, __gong__toRawStringLiteral(clock2ddiagram.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", clock2ddiagramIdent, __gong__toRawStringLiteral(clock2ddiagram.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Zoom = %f", clock2ddiagramIdent, clock2ddiagram.Zoom))
			values.WriteString(fmt.Sprintf("\n\t%s.IsHiddenAxesShape = %t", clock2ddiagramIdent, clock2ddiagram.IsHiddenAxesShape))
			values.WriteString(fmt.Sprintf("\n\t%s.IsChecked = %t", clock2ddiagramIdent, clock2ddiagram.IsChecked))
			values.WriteString(fmt.Sprintf("\n\t%s.ComputedPrefix = %s", clock2ddiagramIdent, __gong__toRawStringLiteral(clock2ddiagram.ComputedPrefix)))
			values.WriteString(fmt.Sprintf("\n\t%s.IsExpanded = %t", clock2ddiagramIdent, clock2ddiagram.IsExpanded))
		}
	}
	if stage != nil {
		for _, clock3ddiagram := range __gong__sortStageSetInstances(stage.Clock3DDiagrams, stage.Clock3DDiagram_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			clock3ddiagramIdent := localAliasPrefix + clock3ddiagram.GongGetIdentifier(stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&models.Clock3DDiagram{Name: %s}).Stage(stage)", clock3ddiagramIdent, __gong__toRawStringLiteral(clock3ddiagram.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", clock3ddiagramIdent, __gong__toRawStringLiteral(clock3ddiagram.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.IsHiddenClockTopCurveShape = %t", clock3ddiagramIdent, clock3ddiagram.IsHiddenClockTopCurveShape))
			values.WriteString(fmt.Sprintf("\n\t%s.IsHiddenTorus3DShape = %t", clock3ddiagramIdent, clock3ddiagram.IsHiddenTorus3DShape))
			values.WriteString(fmt.Sprintf("\n\t%s.IsHiddenSampledPoints3DShape = %t", clock3ddiagramIdent, clock3ddiagram.IsHiddenSampledPoints3DShape))
			values.WriteString(fmt.Sprintf("\n\t%s.IsHiddenTiledFloor3DShape = %t", clock3ddiagramIdent, clock3ddiagram.IsHiddenTiledFloor3DShape))
			values.WriteString(fmt.Sprintf("\n\t%s.IsChecked = %t", clock3ddiagramIdent, clock3ddiagram.IsChecked))
			values.WriteString(fmt.Sprintf("\n\t%s.ComputedPrefix = %s", clock3ddiagramIdent, __gong__toRawStringLiteral(clock3ddiagram.ComputedPrefix)))
			values.WriteString(fmt.Sprintf("\n\t%s.IsExpanded = %t", clock3ddiagramIdent, clock3ddiagram.IsExpanded))
			if clock3ddiagram.SampledPoints3DShape != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := localAliasPrefix + clock3ddiagram.SampledPoints3DShape.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.SampledPoints3DShape = %s", clock3ddiagramIdent, targetIdent))
			}
			if clock3ddiagram.Rendered3DShape != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := localAliasPrefix + clock3ddiagram.Rendered3DShape.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Rendered3DShape = %s", clock3ddiagramIdent, targetIdent))
			}
		}
	}
	if stage != nil {
		for _, cutline3dshape := range __gong__sortStageSetInstances(stage.CutLine3DShapes, stage.CutLine3DShape_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			cutline3dshapeIdent := localAliasPrefix + cutline3dshape.GongGetIdentifier(stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&models.CutLine3DShape{Name: %s}).Stage(stage)", cutline3dshapeIdent, __gong__toRawStringLiteral(cutline3dshape.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", cutline3dshapeIdent, __gong__toRawStringLiteral(cutline3dshape.Name)))
		}
	}
	if stage != nil {
		for _, leaves3dshape := range __gong__sortStageSetInstances(stage.Leaves3DShapes, stage.Leaves3DShape_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			leaves3dshapeIdent := localAliasPrefix + leaves3dshape.GongGetIdentifier(stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&models.Leaves3DShape{Name: %s}).Stage(stage)", leaves3dshapeIdent, __gong__toRawStringLiteral(leaves3dshape.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", leaves3dshapeIdent, __gong__toRawStringLiteral(leaves3dshape.Name)))
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
			for _, elem := range library.Plants {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := localAliasPrefix + elem.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Plants = append(%s.Plants, %s)", libraryIdent, libraryIdent, targetIdent))
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
		}
	}
	if stage != nil {
		for _, originalpoints3dshape := range __gong__sortStageSetInstances(stage.OriginalPoints3DShapes, stage.OriginalPoints3DShape_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			originalpoints3dshapeIdent := localAliasPrefix + originalpoints3dshape.GongGetIdentifier(stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&models.OriginalPoints3DShape{Name: %s}).Stage(stage)", originalpoints3dshapeIdent, __gong__toRawStringLiteral(originalpoints3dshape.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", originalpoints3dshapeIdent, __gong__toRawStringLiteral(originalpoints3dshape.Name)))
		}
	}
	if stage != nil {
		for _, parastichymcurves3dshape := range __gong__sortStageSetInstances(stage.ParastichyMCurves3DShapes, stage.ParastichyMCurves3DShape_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			parastichymcurves3dshapeIdent := localAliasPrefix + parastichymcurves3dshape.GongGetIdentifier(stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&models.ParastichyMCurves3DShape{Name: %s}).Stage(stage)", parastichymcurves3dshapeIdent, __gong__toRawStringLiteral(parastichymcurves3dshape.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", parastichymcurves3dshapeIdent, __gong__toRawStringLiteral(parastichymcurves3dshape.Name)))
		}
	}
	if stage != nil {
		for _, parastichyncurves3dshape := range __gong__sortStageSetInstances(stage.ParastichyNCurves3DShapes, stage.ParastichyNCurves3DShape_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			parastichyncurves3dshapeIdent := localAliasPrefix + parastichyncurves3dshape.GongGetIdentifier(stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&models.ParastichyNCurves3DShape{Name: %s}).Stage(stage)", parastichyncurves3dshapeIdent, __gong__toRawStringLiteral(parastichyncurves3dshape.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", parastichyncurves3dshapeIdent, __gong__toRawStringLiteral(parastichyncurves3dshape.Name)))
		}
	}
	if stage != nil {
		for _, plant2ddiagram := range __gong__sortStageSetInstances(stage.Plant2DDiagrams, stage.Plant2DDiagram_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			plant2ddiagramIdent := localAliasPrefix + plant2ddiagram.GongGetIdentifier(stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&models.Plant2DDiagram{Name: %s}).Stage(stage)", plant2ddiagramIdent, __gong__toRawStringLiteral(plant2ddiagram.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", plant2ddiagramIdent, __gong__toRawStringLiteral(plant2ddiagram.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.OriginX = %f", plant2ddiagramIdent, plant2ddiagram.OriginX))
			values.WriteString(fmt.Sprintf("\n\t%s.OriginY = %f", plant2ddiagramIdent, plant2ddiagram.OriginY))
			values.WriteString(fmt.Sprintf("\n\t%s.Zoom = %f", plant2ddiagramIdent, plant2ddiagram.Zoom))
			values.WriteString(fmt.Sprintf("\n\t%s.IsRhombusNodesExpanded = %t", plant2ddiagramIdent, plant2ddiagram.IsRhombusNodesExpanded))
			values.WriteString(fmt.Sprintf("\n\t%s.IsArcNodesExpanded = %t", plant2ddiagramIdent, plant2ddiagram.IsArcNodesExpanded))
			values.WriteString(fmt.Sprintf("\n\t%s.IsHiddenAxesShape = %t", plant2ddiagramIdent, plant2ddiagram.IsHiddenAxesShape))
			values.WriteString(fmt.Sprintf("\n\t%s.IsHiddenReferenceRhombus = %t", plant2ddiagramIdent, plant2ddiagram.IsHiddenReferenceRhombus))
			values.WriteString(fmt.Sprintf("\n\t%s.IsHiddenPlantCircumferenceShape = %t", plant2ddiagramIdent, plant2ddiagram.IsHiddenPlantCircumferenceShape))
			values.WriteString(fmt.Sprintf("\n\t%s.IsHiddenGridPathShape = %t", plant2ddiagramIdent, plant2ddiagram.IsHiddenGridPathShape))
			values.WriteString(fmt.Sprintf("\n\t%s.IsHiddenRhombusGridShape = %t", plant2ddiagramIdent, plant2ddiagram.IsHiddenRhombusGridShape))
			values.WriteString(fmt.Sprintf("\n\t%s.IsHiddenExplanationTextShape = %t", plant2ddiagramIdent, plant2ddiagram.IsHiddenExplanationTextShape))
			values.WriteString(fmt.Sprintf("\n\t%s.IsHiddenRotatedReferenceRhombus = %t", plant2ddiagramIdent, plant2ddiagram.IsHiddenRotatedReferenceRhombus))
			values.WriteString(fmt.Sprintf("\n\t%s.IsHiddenRotatedPlantCircumferenceShape = %t", plant2ddiagramIdent, plant2ddiagram.IsHiddenRotatedPlantCircumferenceShape))
			values.WriteString(fmt.Sprintf("\n\t%s.IsHiddenRotatedGridPathShape = %t", plant2ddiagramIdent, plant2ddiagram.IsHiddenRotatedGridPathShape))
			values.WriteString(fmt.Sprintf("\n\t%s.IsHiddenRotatedRhombusGridShape = %t", plant2ddiagramIdent, plant2ddiagram.IsHiddenRotatedRhombusGridShape))
			values.WriteString(fmt.Sprintf("\n\t%s.IsHiddenGrowthPathRhombusGridShape = %t", plant2ddiagramIdent, plant2ddiagram.IsHiddenGrowthPathRhombusGridShape))
			values.WriteString(fmt.Sprintf("\n\t%s.IsHiddenGrowthVectorShape = %t", plant2ddiagramIdent, plant2ddiagram.IsHiddenGrowthVectorShape))
			values.WriteString(fmt.Sprintf("\n\t%s.IsHiddenPerpendicularVectorGrid = %t", plant2ddiagramIdent, plant2ddiagram.IsHiddenPerpendicularVectorGrid))
			values.WriteString(fmt.Sprintf("\n\t%s.IsHiddenBaseVectorShapeGrid = %t", plant2ddiagramIdent, plant2ddiagram.IsHiddenBaseVectorShapeGrid))
			values.WriteString(fmt.Sprintf("\n\t%s.IsHiddenArcNormalVectorShapeGrid = %t", plant2ddiagramIdent, plant2ddiagram.IsHiddenArcNormalVectorShapeGrid))
			values.WriteString(fmt.Sprintf("\n\t%s.IsHiddenStartArcShapeGrid = %t", plant2ddiagramIdent, plant2ddiagram.IsHiddenStartArcShapeGrid))
			values.WriteString(fmt.Sprintf("\n\t%s.IsHiddenMidArcVectorShapeGrid = %t", plant2ddiagramIdent, plant2ddiagram.IsHiddenMidArcVectorShapeGrid))
			values.WriteString(fmt.Sprintf("\n\t%s.IsHiddenEndArcShapeGrid = %t", plant2ddiagramIdent, plant2ddiagram.IsHiddenEndArcShapeGrid))
			values.WriteString(fmt.Sprintf("\n\t%s.IsHiddenGrowthCurve2D = %t", plant2ddiagramIdent, plant2ddiagram.IsHiddenGrowthCurve2D))
			values.WriteString(fmt.Sprintf("\n\t%s.IsHiddenStackOfGrowthCurve2DByGrowthVector = %t", plant2ddiagramIdent, plant2ddiagram.IsHiddenStackOfGrowthCurve2DByGrowthVector))
			values.WriteString(fmt.Sprintf("\n\t%s.IsChecked = %t", plant2ddiagramIdent, plant2ddiagram.IsChecked))
			values.WriteString(fmt.Sprintf("\n\t%s.ComputedPrefix = %s", plant2ddiagramIdent, __gong__toRawStringLiteral(plant2ddiagram.ComputedPrefix)))
			values.WriteString(fmt.Sprintf("\n\t%s.IsExpanded = %t", plant2ddiagramIdent, plant2ddiagram.IsExpanded))
		}
	}
	if stage != nil {
		for _, plant3ddiagram := range __gong__sortStageSetInstances(stage.Plant3DDiagrams, stage.Plant3DDiagram_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			plant3ddiagramIdent := localAliasPrefix + plant3ddiagram.GongGetIdentifier(stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&models.Plant3DDiagram{Name: %s}).Stage(stage)", plant3ddiagramIdent, __gong__toRawStringLiteral(plant3ddiagram.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", plant3ddiagramIdent, __gong__toRawStringLiteral(plant3ddiagram.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.IsHiddenStemCylinder3DShape = %t", plant3ddiagramIdent, plant3ddiagram.IsHiddenStemCylinder3DShape))
			values.WriteString(fmt.Sprintf("\n\t%s.IsHiddenParastichyNCurves3DShape = %t", plant3ddiagramIdent, plant3ddiagram.IsHiddenParastichyNCurves3DShape))
			values.WriteString(fmt.Sprintf("\n\t%s.IsHiddenParastichyMCurves3DShape = %t", plant3ddiagramIdent, plant3ddiagram.IsHiddenParastichyMCurves3DShape))
			values.WriteString(fmt.Sprintf("\n\t%s.IsHiddenCutLine3DShape = %t", plant3ddiagramIdent, plant3ddiagram.IsHiddenCutLine3DShape))
			values.WriteString(fmt.Sprintf("\n\t%s.IsHiddenCircumference3DShape = %t", plant3ddiagramIdent, plant3ddiagram.IsHiddenCircumference3DShape))
			values.WriteString(fmt.Sprintf("\n\t%s.IsHiddenTiledFloor3DShape = %t", plant3ddiagramIdent, plant3ddiagram.IsHiddenTiledFloor3DShape))
			values.WriteString(fmt.Sprintf("\n\t%s.IsHiddenLeaves3DShape = %t", plant3ddiagramIdent, plant3ddiagram.IsHiddenLeaves3DShape))
			values.WriteString(fmt.Sprintf("\n\t%s.IsChecked = %t", plant3ddiagramIdent, plant3ddiagram.IsChecked))
			values.WriteString(fmt.Sprintf("\n\t%s.ComputedPrefix = %s", plant3ddiagramIdent, __gong__toRawStringLiteral(plant3ddiagram.ComputedPrefix)))
			values.WriteString(fmt.Sprintf("\n\t%s.IsExpanded = %t", plant3ddiagramIdent, plant3ddiagram.IsExpanded))
			if plant3ddiagram.StemCylinder3DShape != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := localAliasPrefix + plant3ddiagram.StemCylinder3DShape.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.StemCylinder3DShape = %s", plant3ddiagramIdent, targetIdent))
			}
			if plant3ddiagram.ParastichyNCurves3DShape != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := localAliasPrefix + plant3ddiagram.ParastichyNCurves3DShape.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.ParastichyNCurves3DShape = %s", plant3ddiagramIdent, targetIdent))
			}
			if plant3ddiagram.ParastichyMCurves3DShape != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := localAliasPrefix + plant3ddiagram.ParastichyMCurves3DShape.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.ParastichyMCurves3DShape = %s", plant3ddiagramIdent, targetIdent))
			}
			if plant3ddiagram.CutLine3DShape != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := localAliasPrefix + plant3ddiagram.CutLine3DShape.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.CutLine3DShape = %s", plant3ddiagramIdent, targetIdent))
			}
			if plant3ddiagram.Circumference3DShape != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := localAliasPrefix + plant3ddiagram.Circumference3DShape.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Circumference3DShape = %s", plant3ddiagramIdent, targetIdent))
			}
			if plant3ddiagram.Leaves3DShape != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := localAliasPrefix + plant3ddiagram.Leaves3DShape.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Leaves3DShape = %s", plant3ddiagramIdent, targetIdent))
			}
			if plant3ddiagram.Rendered3DShape != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := localAliasPrefix + plant3ddiagram.Rendered3DShape.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Rendered3DShape = %s", plant3ddiagramIdent, targetIdent))
			}
		}
	}
	if stage != nil {
		for _, plantabstract := range __gong__sortStageSetInstances(stage.PlantAbstracts, stage.PlantAbstract_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			plantabstractIdent := localAliasPrefix + plantabstract.GongGetIdentifier(stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&models.PlantAbstract{Name: %s}).Stage(stage)", plantabstractIdent, __gong__toRawStringLiteral(plantabstract.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", plantabstractIdent, __gong__toRawStringLiteral(plantabstract.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.N = %d", plantabstractIdent, plantabstract.N))
			values.WriteString(fmt.Sprintf("\n\t%s.M = %d", plantabstractIdent, plantabstract.M))
			values.WriteString(fmt.Sprintf("\n\t%s.StackHeight = %d", plantabstractIdent, plantabstract.StackHeight))
			values.WriteString(fmt.Sprintf("\n\t%s.RhombusInsideAngle = %f", plantabstractIdent, plantabstract.RhombusInsideAngle))
			values.WriteString(fmt.Sprintf("\n\t%s.RhombusSideLength = %f", plantabstractIdent, plantabstract.RhombusSideLength))
			if plantabstract.PlantType.ToCodeString() != "" {
				values.WriteString(fmt.Sprintf("\n\t%s.PlantType = models.%s", plantabstractIdent, plantabstract.PlantType.ToCodeString()))
			} else {
				values.WriteString(fmt.Sprintf("\n\t%s.PlantType = \"\"", plantabstractIdent))
			}
			if plantabstract.CurrentView.ToCodeString() != "" {
				values.WriteString(fmt.Sprintf("\n\t%s.CurrentView = models.%s", plantabstractIdent, plantabstract.CurrentView.ToCodeString()))
			} else {
				values.WriteString(fmt.Sprintf("\n\t%s.CurrentView = \"\"", plantabstractIdent))
			}
			values.WriteString(fmt.Sprintf("\n\t%s.ComputedPrefix = %s", plantabstractIdent, __gong__toRawStringLiteral(plantabstract.ComputedPrefix)))
			values.WriteString(fmt.Sprintf("\n\t%s.IsExpanded = %t", plantabstractIdent, plantabstract.IsExpanded))
			values.WriteString(fmt.Sprintf("\n\t%s.IsSelected = %t", plantabstractIdent, plantabstract.IsSelected))
			values.WriteString(fmt.Sprintf("\n\t%s.IsPlant2DDiagramsNodeExpanded = %t", plantabstractIdent, plantabstract.IsPlant2DDiagramsNodeExpanded))
			values.WriteString(fmt.Sprintf("\n\t%s.IsPlant3DDiagramsNodeExpanded = %t", plantabstractIdent, plantabstract.IsPlant3DDiagramsNodeExpanded))
			values.WriteString(fmt.Sprintf("\n\t%s.IsVase2DDiagramsNodeExpanded = %t", plantabstractIdent, plantabstract.IsVase2DDiagramsNodeExpanded))
			values.WriteString(fmt.Sprintf("\n\t%s.IsTubeVase3DDiagramsNodeExpanded = %t", plantabstractIdent, plantabstract.IsTubeVase3DDiagramsNodeExpanded))
			values.WriteString(fmt.Sprintf("\n\t%s.IsStool2DDiagramsNodeExpanded = %t", plantabstractIdent, plantabstract.IsStool2DDiagramsNodeExpanded))
			values.WriteString(fmt.Sprintf("\n\t%s.IsStool3DDiagramsNodeExpanded = %t", plantabstractIdent, plantabstract.IsStool3DDiagramsNodeExpanded))
			values.WriteString(fmt.Sprintf("\n\t%s.IsClock2DDiagramsNodeExpanded = %t", plantabstractIdent, plantabstract.IsClock2DDiagramsNodeExpanded))
			values.WriteString(fmt.Sprintf("\n\t%s.IsClock3DDiagramsNodeExpanded = %t", plantabstractIdent, plantabstract.IsClock3DDiagramsNodeExpanded))
			if plantabstract.TubeVaseAbstract != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := localAliasPrefix + plantabstract.TubeVaseAbstract.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.TubeVaseAbstract = %s", plantabstractIdent, targetIdent))
			}
			if plantabstract.StoolAbstract != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := "__stool" + plantabstract.StoolAbstract.GongGetIdentifier(stage.StoolStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.StoolAbstract = %s", plantabstractIdent, targetIdent))
			}
			if plantabstract.ClockAbstract != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := "__clock" + plantabstract.ClockAbstract.GongGetIdentifier(stage.ClockStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.ClockAbstract = %s", plantabstractIdent, targetIdent))
			}
			if plantabstract.MusicAbstract != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := "__music" + plantabstract.MusicAbstract.GongGetIdentifier(stage.MusicStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.MusicAbstract = %s", plantabstractIdent, targetIdent))
			}
			for _, elem := range plantabstract.Plant2DDiagrams {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := localAliasPrefix + elem.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Plant2DDiagrams = append(%s.Plant2DDiagrams, %s)", plantabstractIdent, plantabstractIdent, targetIdent))
			}
			for _, elem := range plantabstract.Plant3DDiagrams {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := localAliasPrefix + elem.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Plant3DDiagrams = append(%s.Plant3DDiagrams, %s)", plantabstractIdent, plantabstractIdent, targetIdent))
			}
			for _, elem := range plantabstract.Vase2DDiagrams {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := localAliasPrefix + elem.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Vase2DDiagrams = append(%s.Vase2DDiagrams, %s)", plantabstractIdent, plantabstractIdent, targetIdent))
			}
			for _, elem := range plantabstract.TubeVase3DDiagrams {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := localAliasPrefix + elem.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.TubeVase3DDiagrams = append(%s.TubeVase3DDiagrams, %s)", plantabstractIdent, plantabstractIdent, targetIdent))
			}
			for _, elem := range plantabstract.Stool2DDiagrams {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := localAliasPrefix + elem.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Stool2DDiagrams = append(%s.Stool2DDiagrams, %s)", plantabstractIdent, plantabstractIdent, targetIdent))
			}
			for _, elem := range plantabstract.Stool3DDiagrams {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := localAliasPrefix + elem.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Stool3DDiagrams = append(%s.Stool3DDiagrams, %s)", plantabstractIdent, plantabstractIdent, targetIdent))
			}
			for _, elem := range plantabstract.Clock2DDiagrams {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := localAliasPrefix + elem.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Clock2DDiagrams = append(%s.Clock2DDiagrams, %s)", plantabstractIdent, plantabstractIdent, targetIdent))
			}
			for _, elem := range plantabstract.Clock3DDiagrams {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := localAliasPrefix + elem.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Clock3DDiagrams = append(%s.Clock3DDiagrams, %s)", plantabstractIdent, plantabstractIdent, targetIdent))
			}
		}
	}
	if stage != nil {
		for _, rendered3dshape := range __gong__sortStageSetInstances(stage.Rendered3DShapes, stage.Rendered3DShape_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			rendered3dshapeIdent := localAliasPrefix + rendered3dshape.GongGetIdentifier(stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&models.Rendered3DShape{Name: %s}).Stage(stage)", rendered3dshapeIdent, __gong__toRawStringLiteral(rendered3dshape.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", rendered3dshapeIdent, __gong__toRawStringLiteral(rendered3dshape.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.ViewX = %f", rendered3dshapeIdent, rendered3dshape.ViewX))
			values.WriteString(fmt.Sprintf("\n\t%s.ViewY = %f", rendered3dshapeIdent, rendered3dshape.ViewY))
			values.WriteString(fmt.Sprintf("\n\t%s.ViewZ = %f", rendered3dshapeIdent, rendered3dshape.ViewZ))
			values.WriteString(fmt.Sprintf("\n\t%s.TargetX = %f", rendered3dshapeIdent, rendered3dshape.TargetX))
			values.WriteString(fmt.Sprintf("\n\t%s.TargetY = %f", rendered3dshapeIdent, rendered3dshape.TargetY))
			values.WriteString(fmt.Sprintf("\n\t%s.TargetZ = %f", rendered3dshapeIdent, rendered3dshape.TargetZ))
			values.WriteString(fmt.Sprintf("\n\t%s.Fov = %f", rendered3dshapeIdent, rendered3dshape.Fov))
		}
	}
	if stage != nil {
		for _, sampledpoints3dshape := range __gong__sortStageSetInstances(stage.SampledPoints3DShapes, stage.SampledPoints3DShape_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			sampledpoints3dshapeIdent := localAliasPrefix + sampledpoints3dshape.GongGetIdentifier(stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&models.SampledPoints3DShape{Name: %s}).Stage(stage)", sampledpoints3dshapeIdent, __gong__toRawStringLiteral(sampledpoints3dshape.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", sampledpoints3dshapeIdent, __gong__toRawStringLiteral(sampledpoints3dshape.Name)))
		}
	}
	if stage != nil {
		for _, stackofrotatedcarvedoutvasetrapezeringsshape := range __gong__sortStageSetInstances(stage.StackOfRotatedCarvedOutVaseTrapezeRingsShapes, stage.StackOfRotatedCarvedOutVaseTrapezeRingsShape_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			stackofrotatedcarvedoutvasetrapezeringsshapeIdent := localAliasPrefix + stackofrotatedcarvedoutvasetrapezeringsshape.GongGetIdentifier(stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&models.StackOfRotatedCarvedOutVaseTrapezeRingsShape{Name: %s}).Stage(stage)", stackofrotatedcarvedoutvasetrapezeringsshapeIdent, __gong__toRawStringLiteral(stackofrotatedcarvedoutvasetrapezeringsshape.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", stackofrotatedcarvedoutvasetrapezeringsshapeIdent, __gong__toRawStringLiteral(stackofrotatedcarvedoutvasetrapezeringsshape.Name)))
		}
	}
	if stage != nil {
		for _, stemcylinder3dshape := range __gong__sortStageSetInstances(stage.StemCylinder3DShapes, stage.StemCylinder3DShape_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			stemcylinder3dshapeIdent := localAliasPrefix + stemcylinder3dshape.GongGetIdentifier(stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&models.StemCylinder3DShape{Name: %s}).Stage(stage)", stemcylinder3dshapeIdent, __gong__toRawStringLiteral(stemcylinder3dshape.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", stemcylinder3dshapeIdent, __gong__toRawStringLiteral(stemcylinder3dshape.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Transparency = %f", stemcylinder3dshapeIdent, stemcylinder3dshape.Transparency))
		}
	}
	if stage != nil {
		for _, stool2ddiagram := range __gong__sortStageSetInstances(stage.Stool2DDiagrams, stage.Stool2DDiagram_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			stool2ddiagramIdent := localAliasPrefix + stool2ddiagram.GongGetIdentifier(stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&models.Stool2DDiagram{Name: %s}).Stage(stage)", stool2ddiagramIdent, __gong__toRawStringLiteral(stool2ddiagram.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", stool2ddiagramIdent, __gong__toRawStringLiteral(stool2ddiagram.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Zoom = %f", stool2ddiagramIdent, stool2ddiagram.Zoom))
			values.WriteString(fmt.Sprintf("\n\t%s.IsHiddenAxesShape = %t", stool2ddiagramIdent, stool2ddiagram.IsHiddenAxesShape))
			values.WriteString(fmt.Sprintf("\n\t%s.IsChecked = %t", stool2ddiagramIdent, stool2ddiagram.IsChecked))
			values.WriteString(fmt.Sprintf("\n\t%s.ComputedPrefix = %s", stool2ddiagramIdent, __gong__toRawStringLiteral(stool2ddiagram.ComputedPrefix)))
			values.WriteString(fmt.Sprintf("\n\t%s.IsExpanded = %t", stool2ddiagramIdent, stool2ddiagram.IsExpanded))
		}
	}
	if stage != nil {
		for _, stool3ddiagram := range __gong__sortStageSetInstances(stage.Stool3DDiagrams, stage.Stool3DDiagram_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			stool3ddiagramIdent := localAliasPrefix + stool3ddiagram.GongGetIdentifier(stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&models.Stool3DDiagram{Name: %s}).Stage(stage)", stool3ddiagramIdent, __gong__toRawStringLiteral(stool3ddiagram.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", stool3ddiagramIdent, __gong__toRawStringLiteral(stool3ddiagram.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.IsHiddenSeatTopCurveShape = %t", stool3ddiagramIdent, stool3ddiagram.IsHiddenSeatTopCurveShape))
			values.WriteString(fmt.Sprintf("\n\t%s.IsHiddenRotatedSeatTopCurveShape = %t", stool3ddiagramIdent, stool3ddiagram.IsHiddenRotatedSeatTopCurveShape))
			values.WriteString(fmt.Sprintf("\n\t%s.IsHiddenSeatBottomCurveShape = %t", stool3ddiagramIdent, stool3ddiagram.IsHiddenSeatBottomCurveShape))
			values.WriteString(fmt.Sprintf("\n\t%s.IsHiddenRotatedSeatBottomCurveShape = %t", stool3ddiagramIdent, stool3ddiagram.IsHiddenRotatedSeatBottomCurveShape))
			values.WriteString(fmt.Sprintf("\n\t%s.IsHiddenTorus3DShape = %t", stool3ddiagramIdent, stool3ddiagram.IsHiddenTorus3DShape))
			values.WriteString(fmt.Sprintf("\n\t%s.IsHiddenRotatedTorusShape = %t", stool3ddiagramIdent, stool3ddiagram.IsHiddenRotatedTorusShape))
			values.WriteString(fmt.Sprintf("\n\t%s.IsHiddenSampledPoints3DShape = %t", stool3ddiagramIdent, stool3ddiagram.IsHiddenSampledPoints3DShape))
			values.WriteString(fmt.Sprintf("\n\t%s.IsHiddenRotatedSampledPoints3DShape = %t", stool3ddiagramIdent, stool3ddiagram.IsHiddenRotatedSampledPoints3DShape))
			values.WriteString(fmt.Sprintf("\n\t%s.IsHiddenEyeSampledPoints3DShape = %t", stool3ddiagramIdent, stool3ddiagram.IsHiddenEyeSampledPoints3DShape))
			values.WriteString(fmt.Sprintf("\n\t%s.IsHiddenEyeCornersSampledPoints3DShape = %t", stool3ddiagramIdent, stool3ddiagram.IsHiddenEyeCornersSampledPoints3DShape))
			values.WriteString(fmt.Sprintf("\n\t%s.IsHiddenEye3DShape = %t", stool3ddiagramIdent, stool3ddiagram.IsHiddenEye3DShape))
			values.WriteString(fmt.Sprintf("\n\t%s.IsHiddenEyeSeatBottomCurveShape = %t", stool3ddiagramIdent, stool3ddiagram.IsHiddenEyeSeatBottomCurveShape))
			values.WriteString(fmt.Sprintf("\n\t%s.IsHiddenEyeStoolBottomCurveShape = %t", stool3ddiagramIdent, stool3ddiagram.IsHiddenEyeStoolBottomCurveShape))
			values.WriteString(fmt.Sprintf("\n\t%s.IsHiddenSeat3DShape = %t", stool3ddiagramIdent, stool3ddiagram.IsHiddenSeat3DShape))
			values.WriteString(fmt.Sprintf("\n\t%s.IsHiddenEyeVolume3DShape = %t", stool3ddiagramIdent, stool3ddiagram.IsHiddenEyeVolume3DShape))
			values.WriteString(fmt.Sprintf("\n\t%s.IsHiddenSeatAndLegs3DShape = %t", stool3ddiagramIdent, stool3ddiagram.IsHiddenSeatAndLegs3DShape))
			values.WriteString(fmt.Sprintf("\n\t%s.IsHiddenRotatedSeatAndLegs3DShape = %t", stool3ddiagramIdent, stool3ddiagram.IsHiddenRotatedSeatAndLegs3DShape))
			values.WriteString(fmt.Sprintf("\n\t%s.IsHiddenTiledFloor3DShape = %t", stool3ddiagramIdent, stool3ddiagram.IsHiddenTiledFloor3DShape))
			values.WriteString(fmt.Sprintf("\n\t%s.IsChecked = %t", stool3ddiagramIdent, stool3ddiagram.IsChecked))
			values.WriteString(fmt.Sprintf("\n\t%s.ComputedPrefix = %s", stool3ddiagramIdent, __gong__toRawStringLiteral(stool3ddiagram.ComputedPrefix)))
			values.WriteString(fmt.Sprintf("\n\t%s.IsExpanded = %t", stool3ddiagramIdent, stool3ddiagram.IsExpanded))
			if stool3ddiagram.SampledPoints3DShape != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := localAliasPrefix + stool3ddiagram.SampledPoints3DShape.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.SampledPoints3DShape = %s", stool3ddiagramIdent, targetIdent))
			}
			if stool3ddiagram.Rendered3DShape != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := localAliasPrefix + stool3ddiagram.Rendered3DShape.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Rendered3DShape = %s", stool3ddiagramIdent, targetIdent))
			}
		}
	}
	if stage != nil {
		for _, topcurveplane1shape := range __gong__sortStageSetInstances(stage.TopCurvePlane1Shapes, stage.TopCurvePlane1Shape_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			topcurveplane1shapeIdent := localAliasPrefix + topcurveplane1shape.GongGetIdentifier(stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&models.TopCurvePlane1Shape{Name: %s}).Stage(stage)", topcurveplane1shapeIdent, __gong__toRawStringLiteral(topcurveplane1shape.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", topcurveplane1shapeIdent, __gong__toRawStringLiteral(topcurveplane1shape.Name)))
		}
	}
	if stage != nil {
		for _, topcurveplane2shape := range __gong__sortStageSetInstances(stage.TopCurvePlane2Shapes, stage.TopCurvePlane2Shape_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			topcurveplane2shapeIdent := localAliasPrefix + topcurveplane2shape.GongGetIdentifier(stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&models.TopCurvePlane2Shape{Name: %s}).Stage(stage)", topcurveplane2shapeIdent, __gong__toRawStringLiteral(topcurveplane2shape.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", topcurveplane2shapeIdent, __gong__toRawStringLiteral(topcurveplane2shape.Name)))
		}
	}
	if stage != nil {
		for _, tubevase3ddiagram := range __gong__sortStageSetInstances(stage.TubeVase3DDiagrams, stage.TubeVase3DDiagram_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			tubevase3ddiagramIdent := localAliasPrefix + tubevase3ddiagram.GongGetIdentifier(stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&models.TubeVase3DDiagram{Name: %s}).Stage(stage)", tubevase3ddiagramIdent, __gong__toRawStringLiteral(tubevase3ddiagram.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", tubevase3ddiagramIdent, __gong__toRawStringLiteral(tubevase3ddiagram.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.IsHiddenStackOfPartiallyRotatedGrowthCurve2DRibbon = %t", tubevase3ddiagramIdent, tubevase3ddiagram.IsHiddenStackOfPartiallyRotatedGrowthCurve2DRibbon))
			values.WriteString(fmt.Sprintf("\n\t%s.IsHiddenTorusStackShape = %t", tubevase3ddiagramIdent, tubevase3ddiagram.IsHiddenTorusStackShape))
			values.WriteString(fmt.Sprintf("\n\t%s.IsHiddenVerticalTorusStackShape = %t", tubevase3ddiagramIdent, tubevase3ddiagram.IsHiddenVerticalTorusStackShape))
			values.WriteString(fmt.Sprintf("\n\t%s.IsHiddenPartiallyRotatedTorusShape = %t", tubevase3ddiagramIdent, tubevase3ddiagram.IsHiddenPartiallyRotatedTorusShape))
			values.WriteString(fmt.Sprintf("\n\t%s.IsHiddenStackOfPartiallyRotatedTorusShape = %t", tubevase3ddiagramIdent, tubevase3ddiagram.IsHiddenStackOfPartiallyRotatedTorusShape))
			values.WriteString(fmt.Sprintf("\n\t%s.IsHiddenPointsAndLines3DShape = %t", tubevase3ddiagramIdent, tubevase3ddiagram.IsHiddenPointsAndLines3DShape))
			values.WriteString(fmt.Sprintf("\n\t%s.IsHiddenKeyHole3DShape = %t", tubevase3ddiagramIdent, tubevase3ddiagram.IsHiddenKeyHole3DShape))
			values.WriteString(fmt.Sprintf("\n\t%s.IsHiddenKey3DShape = %t", tubevase3ddiagramIdent, tubevase3ddiagram.IsHiddenKey3DShape))
			values.WriteString(fmt.Sprintf("\n\t%s.IsHiddenVolumeKey3DShape = %t", tubevase3ddiagramIdent, tubevase3ddiagram.IsHiddenVolumeKey3DShape))
			values.WriteString(fmt.Sprintf("\n\t%s.IsHiddenTorusEdge3DShape = %t", tubevase3ddiagramIdent, tubevase3ddiagram.IsHiddenTorusEdge3DShape))
			values.WriteString(fmt.Sprintf("\n\t%s.IsHiddenSampledPoints3DShape = %t", tubevase3ddiagramIdent, tubevase3ddiagram.IsHiddenSampledPoints3DShape))
			values.WriteString(fmt.Sprintf("\n\t%s.IsHiddenOriginalPoints3DShape = %t", tubevase3ddiagramIdent, tubevase3ddiagram.IsHiddenOriginalPoints3DShape))
			values.WriteString(fmt.Sprintf("\n\t%s.IsHiddenAngle0Shape = %t", tubevase3ddiagramIdent, tubevase3ddiagram.IsHiddenAngle0Shape))
			values.WriteString(fmt.Sprintf("\n\t%s.IsHiddenTiledFloor3DShape = %t", tubevase3ddiagramIdent, tubevase3ddiagram.IsHiddenTiledFloor3DShape))
			values.WriteString(fmt.Sprintf("\n\t%s.IsHiddenTopCurvePlane1Shape = %t", tubevase3ddiagramIdent, tubevase3ddiagram.IsHiddenTopCurvePlane1Shape))
			values.WriteString(fmt.Sprintf("\n\t%s.IsHiddenBottomCurvePlane1Shape = %t", tubevase3ddiagramIdent, tubevase3ddiagram.IsHiddenBottomCurvePlane1Shape))
			values.WriteString(fmt.Sprintf("\n\t%s.IsHiddenTopCurvePlane2Shape = %t", tubevase3ddiagramIdent, tubevase3ddiagram.IsHiddenTopCurvePlane2Shape))
			values.WriteString(fmt.Sprintf("\n\t%s.IsHiddenBottomCurvePlane2Shape = %t", tubevase3ddiagramIdent, tubevase3ddiagram.IsHiddenBottomCurvePlane2Shape))
			values.WriteString(fmt.Sprintf("\n\t%s.IsHiddenCarvedOutTopCurvePlane1Shape = %t", tubevase3ddiagramIdent, tubevase3ddiagram.IsHiddenCarvedOutTopCurvePlane1Shape))
			values.WriteString(fmt.Sprintf("\n\t%s.IsHiddenCarvedOutBottomCurvePlane1Shape = %t", tubevase3ddiagramIdent, tubevase3ddiagram.IsHiddenCarvedOutBottomCurvePlane1Shape))
			values.WriteString(fmt.Sprintf("\n\t%s.IsHiddenVaseTrapezeRingShape = %t", tubevase3ddiagramIdent, tubevase3ddiagram.IsHiddenVaseTrapezeRingShape))
			values.WriteString(fmt.Sprintf("\n\t%s.IsHiddenCarvedOutVaseTrapezeRingShape = %t", tubevase3ddiagramIdent, tubevase3ddiagram.IsHiddenCarvedOutVaseTrapezeRingShape))
			values.WriteString(fmt.Sprintf("\n\t%s.IsHiddenStackOfRotatedCarvedOutVaseTrapezeRingsShape = %t", tubevase3ddiagramIdent, tubevase3ddiagram.IsHiddenStackOfRotatedCarvedOutVaseTrapezeRingsShape))
			values.WriteString(fmt.Sprintf("\n\t%s.IsHiddenVaseTrapezeBasePlateShape = %t", tubevase3ddiagramIdent, tubevase3ddiagram.IsHiddenVaseTrapezeBasePlateShape))
			values.WriteString(fmt.Sprintf("\n\t%s.IsChecked = %t", tubevase3ddiagramIdent, tubevase3ddiagram.IsChecked))
			values.WriteString(fmt.Sprintf("\n\t%s.ComputedPrefix = %s", tubevase3ddiagramIdent, __gong__toRawStringLiteral(tubevase3ddiagram.ComputedPrefix)))
			values.WriteString(fmt.Sprintf("\n\t%s.IsExpanded = %t", tubevase3ddiagramIdent, tubevase3ddiagram.IsExpanded))
			if tubevase3ddiagram.Rendered3DShape != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := localAliasPrefix + tubevase3ddiagram.Rendered3DShape.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Rendered3DShape = %s", tubevase3ddiagramIdent, targetIdent))
			}
			if tubevase3ddiagram.SampledPoints3DShape != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := localAliasPrefix + tubevase3ddiagram.SampledPoints3DShape.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.SampledPoints3DShape = %s", tubevase3ddiagramIdent, targetIdent))
			}
			if tubevase3ddiagram.OriginalPoints3DShape != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := localAliasPrefix + tubevase3ddiagram.OriginalPoints3DShape.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.OriginalPoints3DShape = %s", tubevase3ddiagramIdent, targetIdent))
			}
			if tubevase3ddiagram.Angle0Shape != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := localAliasPrefix + tubevase3ddiagram.Angle0Shape.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Angle0Shape = %s", tubevase3ddiagramIdent, targetIdent))
			}
			if tubevase3ddiagram.TopCurvePlane1Shape != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := localAliasPrefix + tubevase3ddiagram.TopCurvePlane1Shape.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.TopCurvePlane1Shape = %s", tubevase3ddiagramIdent, targetIdent))
			}
			if tubevase3ddiagram.BottomCurvePlane1Shape != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := localAliasPrefix + tubevase3ddiagram.BottomCurvePlane1Shape.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.BottomCurvePlane1Shape = %s", tubevase3ddiagramIdent, targetIdent))
			}
			if tubevase3ddiagram.TopCurvePlane2Shape != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := localAliasPrefix + tubevase3ddiagram.TopCurvePlane2Shape.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.TopCurvePlane2Shape = %s", tubevase3ddiagramIdent, targetIdent))
			}
			if tubevase3ddiagram.BottomCurvePlane2Shape != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := localAliasPrefix + tubevase3ddiagram.BottomCurvePlane2Shape.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.BottomCurvePlane2Shape = %s", tubevase3ddiagramIdent, targetIdent))
			}
			if tubevase3ddiagram.CarvedOutTopCurvePlane1Shape != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := localAliasPrefix + tubevase3ddiagram.CarvedOutTopCurvePlane1Shape.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.CarvedOutTopCurvePlane1Shape = %s", tubevase3ddiagramIdent, targetIdent))
			}
			if tubevase3ddiagram.CarvedOutBottomCurvePlane1Shape != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := localAliasPrefix + tubevase3ddiagram.CarvedOutBottomCurvePlane1Shape.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.CarvedOutBottomCurvePlane1Shape = %s", tubevase3ddiagramIdent, targetIdent))
			}
			if tubevase3ddiagram.VaseTrapezeRingShape != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := localAliasPrefix + tubevase3ddiagram.VaseTrapezeRingShape.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.VaseTrapezeRingShape = %s", tubevase3ddiagramIdent, targetIdent))
			}
			if tubevase3ddiagram.CarvedOutVaseTrapezeRingShape != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := localAliasPrefix + tubevase3ddiagram.CarvedOutVaseTrapezeRingShape.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.CarvedOutVaseTrapezeRingShape = %s", tubevase3ddiagramIdent, targetIdent))
			}
			if tubevase3ddiagram.StackOfRotatedCarvedOutVaseTrapezeRingsShape != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := localAliasPrefix + tubevase3ddiagram.StackOfRotatedCarvedOutVaseTrapezeRingsShape.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.StackOfRotatedCarvedOutVaseTrapezeRingsShape = %s", tubevase3ddiagramIdent, targetIdent))
			}
			if tubevase3ddiagram.VaseTrapezeBasePlateShape != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := localAliasPrefix + tubevase3ddiagram.VaseTrapezeBasePlateShape.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.VaseTrapezeBasePlateShape = %s", tubevase3ddiagramIdent, targetIdent))
			}
		}
	}
	if stage != nil {
		for _, tubevaseabstract := range __gong__sortStageSetInstances(stage.TubeVaseAbstracts, stage.TubeVaseAbstract_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			tubevaseabstractIdent := localAliasPrefix + tubevaseabstract.GongGetIdentifier(stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&models.TubeVaseAbstract{Name: %s}).Stage(stage)", tubevaseabstractIdent, __gong__toRawStringLiteral(tubevaseabstract.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", tubevaseabstractIdent, __gong__toRawStringLiteral(tubevaseabstract.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Z_Ribbon = %f", tubevaseabstractIdent, tubevaseabstract.Z_Ribbon))
			values.WriteString(fmt.Sprintf("\n\t%s.RibbonVerticalScale = %f", tubevaseabstractIdent, tubevaseabstract.RibbonVerticalScale))
			values.WriteString(fmt.Sprintf("\n\t%s.TopPlaneHeight = %f", tubevaseabstractIdent, tubevaseabstract.TopPlaneHeight))
			values.WriteString(fmt.Sprintf("\n\t%s.BottomPlaneHeight = %f", tubevaseabstractIdent, tubevaseabstract.BottomPlaneHeight))
			values.WriteString(fmt.Sprintf("\n\t%s.BaseRingBottomPlaneHeight = %f", tubevaseabstractIdent, tubevaseabstract.BaseRingBottomPlaneHeight))
			values.WriteString(fmt.Sprintf("\n\t%s.ProjectionAngle = %f", tubevaseabstractIdent, tubevaseabstract.ProjectionAngle))
			values.WriteString(fmt.Sprintf("\n\t%s.BasePlateHeight = %f", tubevaseabstractIdent, tubevaseabstract.BasePlateHeight))
			values.WriteString(fmt.Sprintf("\n\t%s.CarvedOutTopRingsParameter = %f", tubevaseabstractIdent, tubevaseabstract.CarvedOutTopRingsParameter))
			values.WriteString(fmt.Sprintf("\n\t%s.BulbousStartTangentMagnitude = %f", tubevaseabstractIdent, tubevaseabstract.BulbousStartTangentMagnitude))
			values.WriteString(fmt.Sprintf("\n\t%s.BulbousEndAngle = %f", tubevaseabstractIdent, tubevaseabstract.BulbousEndAngle))
			values.WriteString(fmt.Sprintf("\n\t%s.BulbousEndTangentMagnitude = %f", tubevaseabstractIdent, tubevaseabstract.BulbousEndTangentMagnitude))
			values.WriteString(fmt.Sprintf("\n\t%s.RelativeVerticalThickness = %f", tubevaseabstractIdent, tubevaseabstract.RelativeVerticalThickness))
			values.WriteString(fmt.Sprintf("\n\t%s.RelativeRadialThickness = %f", tubevaseabstractIdent, tubevaseabstract.RelativeRadialThickness))
			values.WriteString(fmt.Sprintf("\n\t%s.RelativeCuttedStackFloorHeight = %f", tubevaseabstractIdent, tubevaseabstract.RelativeCuttedStackFloorHeight))
			values.WriteString(fmt.Sprintf("\n\t%s.RelativeRotatedTorusSeparation = %f", tubevaseabstractIdent, tubevaseabstract.RelativeRotatedTorusSeparation))
			values.WriteString(fmt.Sprintf("\n\t%s.RotationRatio = %f", tubevaseabstractIdent, tubevaseabstract.RotationRatio))
			values.WriteString(fmt.Sprintf("\n\t%s.RadialRepetitions = %d", tubevaseabstractIdent, tubevaseabstract.RadialRepetitions))
			values.WriteString(fmt.Sprintf("\n\t%s.Transparency = %f", tubevaseabstractIdent, tubevaseabstract.Transparency))
			values.WriteString(fmt.Sprintf("\n\t%s.HasAlternatingRingColors = %t", tubevaseabstractIdent, tubevaseabstract.HasAlternatingRingColors))
			values.WriteString(fmt.Sprintf("\n\t%s.RelativeTrajectoryOffsetX = %f", tubevaseabstractIdent, tubevaseabstract.RelativeTrajectoryOffsetX))
			values.WriteString(fmt.Sprintf("\n\t%s.RelativeTrajectoryOffsetY = %f", tubevaseabstractIdent, tubevaseabstract.RelativeTrajectoryOffsetY))
			values.WriteString(fmt.Sprintf("\n\t%s.NbStepP1P2 = %d", tubevaseabstractIdent, tubevaseabstract.NbStepP1P2))
			values.WriteString(fmt.Sprintf("\n\t%s.ChosenStep = %d", tubevaseabstractIdent, tubevaseabstract.ChosenStep))
			values.WriteString(fmt.Sprintf("\n\t%s.RelativeHorizontalRingsHeight = %f", tubevaseabstractIdent, tubevaseabstract.RelativeHorizontalRingsHeight))
			values.WriteString(fmt.Sprintf("\n\t%s.OffsetKeyX = %f", tubevaseabstractIdent, tubevaseabstract.OffsetKeyX))
			values.WriteString(fmt.Sprintf("\n\t%s.OffsetKeyY = %f", tubevaseabstractIdent, tubevaseabstract.OffsetKeyY))
			values.WriteString(fmt.Sprintf("\n\t%s.HeightKey = %f", tubevaseabstractIdent, tubevaseabstract.HeightKey))
			values.WriteString(fmt.Sprintf("\n\t%s.WidthKey = %f", tubevaseabstractIdent, tubevaseabstract.WidthKey))
			values.WriteString(fmt.Sprintf("\n\t%s.RelativeKeySize = %f", tubevaseabstractIdent, tubevaseabstract.RelativeKeySize))
			values.WriteString(fmt.Sprintf("\n\t%s.MovieNbFrames = %d", tubevaseabstractIdent, tubevaseabstract.MovieNbFrames))
		}
	}
	if stage != nil {
		for _, vase2ddiagram := range __gong__sortStageSetInstances(stage.Vase2DDiagrams, stage.Vase2DDiagram_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			vase2ddiagramIdent := localAliasPrefix + vase2ddiagram.GongGetIdentifier(stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&models.Vase2DDiagram{Name: %s}).Stage(stage)", vase2ddiagramIdent, __gong__toRawStringLiteral(vase2ddiagram.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", vase2ddiagramIdent, __gong__toRawStringLiteral(vase2ddiagram.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Zoom = %f", vase2ddiagramIdent, vase2ddiagram.Zoom))
			values.WriteString(fmt.Sprintf("\n\t%s.IsVaseArcNodesExpanded = %t", vase2ddiagramIdent, vase2ddiagram.IsVaseArcNodesExpanded))
			values.WriteString(fmt.Sprintf("\n\t%s.IsVaseClampingNodesExpanded = %t", vase2ddiagramIdent, vase2ddiagram.IsVaseClampingNodesExpanded))
			values.WriteString(fmt.Sprintf("\n\t%s.IsHiddenAxesShape = %t", vase2ddiagramIdent, vase2ddiagram.IsHiddenAxesShape))
			values.WriteString(fmt.Sprintf("\n\t%s.IsHiddenBottomStartArcShapeGrid = %t", vase2ddiagramIdent, vase2ddiagram.IsHiddenBottomStartArcShapeGrid))
			values.WriteString(fmt.Sprintf("\n\t%s.IsHiddenBottomEndArcShapeGrid = %t", vase2ddiagramIdent, vase2ddiagram.IsHiddenBottomEndArcShapeGrid))
			values.WriteString(fmt.Sprintf("\n\t%s.IsHiddenBottomStackOfGrowthCurve = %t", vase2ddiagramIdent, vase2ddiagram.IsHiddenBottomStackOfGrowthCurve))
			values.WriteString(fmt.Sprintf("\n\t%s.IsHiddenShiftedLeftStackOfGrowthCurve = %t", vase2ddiagramIdent, vase2ddiagram.IsHiddenShiftedLeftStackOfGrowthCurve))
			values.WriteString(fmt.Sprintf("\n\t%s.IsHiddenShiftedLeftStackOfNormalVector = %t", vase2ddiagramIdent, vase2ddiagram.IsHiddenShiftedLeftStackOfNormalVector))
			values.WriteString(fmt.Sprintf("\n\t%s.IsHiddenPerpendicularVectorGridHalfway = %t", vase2ddiagramIdent, vase2ddiagram.IsHiddenPerpendicularVectorGridHalfway))
			values.WriteString(fmt.Sprintf("\n\t%s.IsHiddenTopStartArcShapeGrid = %t", vase2ddiagramIdent, vase2ddiagram.IsHiddenTopStartArcShapeGrid))
			values.WriteString(fmt.Sprintf("\n\t%s.IsHiddenShiftedBottomTopStartArcShapeGrid = %t", vase2ddiagramIdent, vase2ddiagram.IsHiddenShiftedBottomTopStartArcShapeGrid))
			values.WriteString(fmt.Sprintf("\n\t%s.IsHiddenTopMidArcVectorShapeGrid = %t", vase2ddiagramIdent, vase2ddiagram.IsHiddenTopMidArcVectorShapeGrid))
			values.WriteString(fmt.Sprintf("\n\t%s.IsHiddenStartHalfwayArcShapeGrid = %t", vase2ddiagramIdent, vase2ddiagram.IsHiddenStartHalfwayArcShapeGrid))
			values.WriteString(fmt.Sprintf("\n\t%s.IsHiddenTopStartHalfwayArcShapeGrid = %t", vase2ddiagramIdent, vase2ddiagram.IsHiddenTopStartHalfwayArcShapeGrid))
			values.WriteString(fmt.Sprintf("\n\t%s.IsHiddenEndHalfwayArcShapeGrid = %t", vase2ddiagramIdent, vase2ddiagram.IsHiddenEndHalfwayArcShapeGrid))
			values.WriteString(fmt.Sprintf("\n\t%s.IsHiddenTopEndHalfwayArcShapeGrid = %t", vase2ddiagramIdent, vase2ddiagram.IsHiddenTopEndHalfwayArcShapeGrid))
			values.WriteString(fmt.Sprintf("\n\t%s.IsHiddenTopEndArcShapeGrid = %t", vase2ddiagramIdent, vase2ddiagram.IsHiddenTopEndArcShapeGrid))
			values.WriteString(fmt.Sprintf("\n\t%s.IsHiddenStackOfGrowthCurve = %t", vase2ddiagramIdent, vase2ddiagram.IsHiddenStackOfGrowthCurve))
			values.WriteString(fmt.Sprintf("\n\t%s.IsHiddenTopStackOfGrowthCurve = %t", vase2ddiagramIdent, vase2ddiagram.IsHiddenTopStackOfGrowthCurve))
			values.WriteString(fmt.Sprintf("\n\t%s.IsHiddenTopGrowthCurve2D = %t", vase2ddiagramIdent, vase2ddiagram.IsHiddenTopGrowthCurve2D))
			values.WriteString(fmt.Sprintf("\n\t%s.IsHiddenStackOfGrowthCurve2D = %t", vase2ddiagramIdent, vase2ddiagram.IsHiddenStackOfGrowthCurve2D))
			values.WriteString(fmt.Sprintf("\n\t%s.IsHiddenTopStackOfGrowthCurve2D = %t", vase2ddiagramIdent, vase2ddiagram.IsHiddenTopStackOfGrowthCurve2D))
			values.WriteString(fmt.Sprintf("\n\t%s.IsHiddenGrowthCurve2DRibbon = %t", vase2ddiagramIdent, vase2ddiagram.IsHiddenGrowthCurve2DRibbon))
			values.WriteString(fmt.Sprintf("\n\t%s.IsHiddenShiftedRightGrowthCurve2DRibbon = %t", vase2ddiagramIdent, vase2ddiagram.IsHiddenShiftedRightGrowthCurve2DRibbon))
			values.WriteString(fmt.Sprintf("\n\t%s.IsHiddenShiftedLeftGrowthCurve2DRibbon = %t", vase2ddiagramIdent, vase2ddiagram.IsHiddenShiftedLeftGrowthCurve2DRibbon))
			values.WriteString(fmt.Sprintf("\n\t%s.IsHiddenStackOfGrowthCurve2DRibbon = %t", vase2ddiagramIdent, vase2ddiagram.IsHiddenStackOfGrowthCurve2DRibbon))
			values.WriteString(fmt.Sprintf("\n\t%s.IsHiddenStackOfRotatedGrowthCurve2DRibbon = %t", vase2ddiagramIdent, vase2ddiagram.IsHiddenStackOfRotatedGrowthCurve2DRibbon))
			values.WriteString(fmt.Sprintf("\n\t%s.IsHiddenPartiallyGrowthCurve2DRibbon = %t", vase2ddiagramIdent, vase2ddiagram.IsHiddenPartiallyGrowthCurve2DRibbon))
			values.WriteString(fmt.Sprintf("\n\t%s.IsHiddenShiftedLeftPartiallyGrowthCurve2DRibbon = %t", vase2ddiagramIdent, vase2ddiagram.IsHiddenShiftedLeftPartiallyGrowthCurve2DRibbon))
			values.WriteString(fmt.Sprintf("\n\t%s.IsHiddenPartiallyGrowthCurve2DTrajectory = %t", vase2ddiagramIdent, vase2ddiagram.IsHiddenPartiallyGrowthCurve2DTrajectory))
			values.WriteString(fmt.Sprintf("\n\t%s.IsHiddenPartiallyGrowthCurve2DTrajectoryP1P2 = %t", vase2ddiagramIdent, vase2ddiagram.IsHiddenPartiallyGrowthCurve2DTrajectoryP1P2))
			values.WriteString(fmt.Sprintf("\n\t%s.IsHiddenPxShape = %t", vase2ddiagramIdent, vase2ddiagram.IsHiddenPxShape))
			values.WriteString(fmt.Sprintf("\n\t%s.IsHiddenChosenP1P2PairShape = %t", vase2ddiagramIdent, vase2ddiagram.IsHiddenChosenP1P2PairShape))
			values.WriteString(fmt.Sprintf("\n\t%s.IsHiddenKeyHoleShape = %t", vase2ddiagramIdent, vase2ddiagram.IsHiddenKeyHoleShape))
			values.WriteString(fmt.Sprintf("\n\t%s.IsChecked = %t", vase2ddiagramIdent, vase2ddiagram.IsChecked))
			values.WriteString(fmt.Sprintf("\n\t%s.ComputedPrefix = %s", vase2ddiagramIdent, __gong__toRawStringLiteral(vase2ddiagram.ComputedPrefix)))
			values.WriteString(fmt.Sprintf("\n\t%s.IsExpanded = %t", vase2ddiagramIdent, vase2ddiagram.IsExpanded))
		}
	}
	if stage != nil {
		for _, vasetrapezebaseplateshape := range __gong__sortStageSetInstances(stage.VaseTrapezeBasePlateShapes, stage.VaseTrapezeBasePlateShape_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			vasetrapezebaseplateshapeIdent := localAliasPrefix + vasetrapezebaseplateshape.GongGetIdentifier(stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&models.VaseTrapezeBasePlateShape{Name: %s}).Stage(stage)", vasetrapezebaseplateshapeIdent, __gong__toRawStringLiteral(vasetrapezebaseplateshape.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", vasetrapezebaseplateshapeIdent, __gong__toRawStringLiteral(vasetrapezebaseplateshape.Name)))
		}
	}
	if stage != nil {
		for _, vasetrapezeringshape := range __gong__sortStageSetInstances(stage.VaseTrapezeRingShapes, stage.VaseTrapezeRingShape_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			vasetrapezeringshapeIdent := localAliasPrefix + vasetrapezeringshape.GongGetIdentifier(stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&models.VaseTrapezeRingShape{Name: %s}).Stage(stage)", vasetrapezeringshapeIdent, __gong__toRawStringLiteral(vasetrapezeringshape.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", vasetrapezeringshapeIdent, __gong__toRawStringLiteral(vasetrapezeringshape.Name)))
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
go run github.com/fullstack-lang/gong/dsm/phylla/go/models/probe/cmd/probe@%s stage.go
*/
package %s

import (
	"slices"
	"time"

	"github.com/fullstack-lang/gong/dsm/phylla/go/models"
	"github.com/fullstack-lang/gong/dsm/phylla/go/models/abstract/stool"
	"github.com/fullstack-lang/gong/dsm/phylla/go/models/abstract/music"
	"github.com/fullstack-lang/gong/dsm/phylla/go/models/abstract/clock"%s
)

var (
	_ time.Time
	_ = slices.Index[[]int, int]

	_ *models.Stage
	_ *stool.Stage
	_ *music.Stage
	_ *clock.Stage
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

	if modelsPackageName != "" && modelsPackageName != "github.com/fullstack-lang/gong/dsm/phylla/go/models" {
		res = strings.ReplaceAll(res, "\"github.com/fullstack-lang/gong/dsm/phylla/go/models\"", fmt.Sprintf("%q", modelsPackageName))
	}

	return res, nil
}

// insertion point for marshall field methods
func (angle0shape *Angle0Shape) GongMarshallField(stage *Stage, fieldName string) (res string) {
	ident := angle0shape.GongGetIdentifier(stage)
	_ = ident
	switch fieldName {
	case "Name":
		res = __gong__marshallString(ident, "Name", angle0shape.Name)

	default:
		log.Panicf("Unknown field %s for Gongstruct Angle0Shape", fieldName)
	}
	return
}

func (arcnormalvectorshape *ArcNormalVectorShape) GongMarshallField(stage *Stage, fieldName string) (res string) {
	ident := arcnormalvectorshape.GongGetIdentifier(stage)
	_ = ident
	switch fieldName {
	case "Name":
		res = __gong__marshallString(ident, "Name", arcnormalvectorshape.Name)
	case "StartX":
		res = __gong__marshallFloat(ident, "StartX", arcnormalvectorshape.StartX)
	case "StartY":
		res = __gong__marshallFloat(ident, "StartY", arcnormalvectorshape.StartY)
	case "EndX":
		res = __gong__marshallFloat(ident, "EndX", arcnormalvectorshape.EndX)
	case "EndY":
		res = __gong__marshallFloat(ident, "EndY", arcnormalvectorshape.EndY)

	default:
		log.Panicf("Unknown field %s for Gongstruct ArcNormalVectorShape", fieldName)
	}
	return
}

func (arcnormalvectorshapegrid *ArcNormalVectorShapeGrid) GongMarshallField(stage *Stage, fieldName string) (res string) {
	ident := arcnormalvectorshapegrid.GongGetIdentifier(stage)
	_ = ident
	switch fieldName {
	case "Name":
		res = __gong__marshallString(ident, "Name", arcnormalvectorshapegrid.Name)

	default:
		log.Panicf("Unknown field %s for Gongstruct ArcNormalVectorShapeGrid", fieldName)
	}
	return
}

func (axesshape *AxesShape) GongMarshallField(stage *Stage, fieldName string) (res string) {
	ident := axesshape.GongGetIdentifier(stage)
	_ = ident
	switch fieldName {
	case "Name":
		res = __gong__marshallString(ident, "Name", axesshape.Name)
	case "LengthX":
		res = __gong__marshallFloat(ident, "LengthX", axesshape.LengthX)
	case "LengthY":
		res = __gong__marshallFloat(ident, "LengthY", axesshape.LengthY)
	case "IsWithHiddenHandle":
		res = __gong__marshallBool(ident, "IsWithHiddenHandle", axesshape.IsWithHiddenHandle)

	default:
		log.Panicf("Unknown field %s for Gongstruct AxesShape", fieldName)
	}
	return
}

func (basevectorshape *BaseVectorShape) GongMarshallField(stage *Stage, fieldName string) (res string) {
	ident := basevectorshape.GongGetIdentifier(stage)
	_ = ident
	switch fieldName {
	case "Name":
		res = __gong__marshallString(ident, "Name", basevectorshape.Name)
	case "StartX":
		res = __gong__marshallFloat(ident, "StartX", basevectorshape.StartX)
	case "StartY":
		res = __gong__marshallFloat(ident, "StartY", basevectorshape.StartY)
	case "EndX":
		res = __gong__marshallFloat(ident, "EndX", basevectorshape.EndX)
	case "EndY":
		res = __gong__marshallFloat(ident, "EndY", basevectorshape.EndY)

	default:
		log.Panicf("Unknown field %s for Gongstruct BaseVectorShape", fieldName)
	}
	return
}

func (basevectorshapegrid *BaseVectorShapeGrid) GongMarshallField(stage *Stage, fieldName string) (res string) {
	ident := basevectorshapegrid.GongGetIdentifier(stage)
	_ = ident
	switch fieldName {
	case "Name":
		res = __gong__marshallString(ident, "Name", basevectorshapegrid.Name)

	default:
		log.Panicf("Unknown field %s for Gongstruct BaseVectorShapeGrid", fieldName)
	}
	return
}

func (bottomcurveplane1shape *BottomCurvePlane1Shape) GongMarshallField(stage *Stage, fieldName string) (res string) {
	ident := bottomcurveplane1shape.GongGetIdentifier(stage)
	_ = ident
	switch fieldName {
	case "Name":
		res = __gong__marshallString(ident, "Name", bottomcurveplane1shape.Name)

	default:
		log.Panicf("Unknown field %s for Gongstruct BottomCurvePlane1Shape", fieldName)
	}
	return
}

func (bottomcurveplane2shape *BottomCurvePlane2Shape) GongMarshallField(stage *Stage, fieldName string) (res string) {
	ident := bottomcurveplane2shape.GongGetIdentifier(stage)
	_ = ident
	switch fieldName {
	case "Name":
		res = __gong__marshallString(ident, "Name", bottomcurveplane2shape.Name)

	default:
		log.Panicf("Unknown field %s for Gongstruct BottomCurvePlane2Shape", fieldName)
	}
	return
}

func (carvedoutbottomcurveplane1shape *CarvedOutBottomCurvePlane1Shape) GongMarshallField(stage *Stage, fieldName string) (res string) {
	ident := carvedoutbottomcurveplane1shape.GongGetIdentifier(stage)
	_ = ident
	switch fieldName {
	case "Name":
		res = __gong__marshallString(ident, "Name", carvedoutbottomcurveplane1shape.Name)

	default:
		log.Panicf("Unknown field %s for Gongstruct CarvedOutBottomCurvePlane1Shape", fieldName)
	}
	return
}

func (carvedouttopcurveplane1shape *CarvedOutTopCurvePlane1Shape) GongMarshallField(stage *Stage, fieldName string) (res string) {
	ident := carvedouttopcurveplane1shape.GongGetIdentifier(stage)
	_ = ident
	switch fieldName {
	case "Name":
		res = __gong__marshallString(ident, "Name", carvedouttopcurveplane1shape.Name)

	default:
		log.Panicf("Unknown field %s for Gongstruct CarvedOutTopCurvePlane1Shape", fieldName)
	}
	return
}

func (carvedoutvasetrapezeringshape *CarvedOutVaseTrapezeRingShape) GongMarshallField(stage *Stage, fieldName string) (res string) {
	ident := carvedoutvasetrapezeringshape.GongGetIdentifier(stage)
	_ = ident
	switch fieldName {
	case "Name":
		res = __gong__marshallString(ident, "Name", carvedoutvasetrapezeringshape.Name)

	default:
		log.Panicf("Unknown field %s for Gongstruct CarvedOutVaseTrapezeRingShape", fieldName)
	}
	return
}

func (chosenp1p2pairshape *ChosenP1P2PairShape) GongMarshallField(stage *Stage, fieldName string) (res string) {
	ident := chosenp1p2pairshape.GongGetIdentifier(stage)
	_ = ident
	switch fieldName {
	case "Name":
		res = __gong__marshallString(ident, "Name", chosenp1p2pairshape.Name)
	case "P1X":
		res = __gong__marshallFloat(ident, "P1X", chosenp1p2pairshape.P1X)
	case "P1Y":
		res = __gong__marshallFloat(ident, "P1Y", chosenp1p2pairshape.P1Y)
	case "P2X":
		res = __gong__marshallFloat(ident, "P2X", chosenp1p2pairshape.P2X)
	case "P2Y":
		res = __gong__marshallFloat(ident, "P2Y", chosenp1p2pairshape.P2Y)
	case "PxX":
		res = __gong__marshallFloat(ident, "PxX", chosenp1p2pairshape.PxX)
	case "PxY":
		res = __gong__marshallFloat(ident, "PxY", chosenp1p2pairshape.PxY)
	case "DistanceP1Px":
		res = __gong__marshallFloat(ident, "DistanceP1Px", chosenp1p2pairshape.DistanceP1Px)
	case "DistanceP2Px":
		res = __gong__marshallFloat(ident, "DistanceP2Px", chosenp1p2pairshape.DistanceP2Px)
	case "DistanceSum":
		res = __gong__marshallFloat(ident, "DistanceSum", chosenp1p2pairshape.DistanceSum)

	default:
		log.Panicf("Unknown field %s for Gongstruct ChosenP1P2PairShape", fieldName)
	}
	return
}

func (circlegridshape *CircleGridShape) GongMarshallField(stage *Stage, fieldName string) (res string) {
	ident := circlegridshape.GongGetIdentifier(stage)
	_ = ident
	switch fieldName {
	case "Name":
		res = __gong__marshallString(ident, "Name", circlegridshape.Name)

	default:
		log.Panicf("Unknown field %s for Gongstruct CircleGridShape", fieldName)
	}
	return
}

func (circumference3dshape *Circumference3DShape) GongMarshallField(stage *Stage, fieldName string) (res string) {
	ident := circumference3dshape.GongGetIdentifier(stage)
	_ = ident
	switch fieldName {
	case "Name":
		res = __gong__marshallString(ident, "Name", circumference3dshape.Name)

	default:
		log.Panicf("Unknown field %s for Gongstruct Circumference3DShape", fieldName)
	}
	return
}

func (clock2ddiagram *Clock2DDiagram) GongMarshallField(stage *Stage, fieldName string) (res string) {
	ident := clock2ddiagram.GongGetIdentifier(stage)
	_ = ident
	switch fieldName {
	case "Name":
		res = __gong__marshallString(ident, "Name", clock2ddiagram.Name)
	case "Zoom":
		res = __gong__marshallFloat(ident, "Zoom", clock2ddiagram.Zoom)
	case "IsHiddenAxesShape":
		res = __gong__marshallBool(ident, "IsHiddenAxesShape", clock2ddiagram.IsHiddenAxesShape)
	case "IsChecked":
		res = __gong__marshallBool(ident, "IsChecked", clock2ddiagram.IsChecked)
	case "ComputedPrefix":
		res = __gong__marshallString(ident, "ComputedPrefix", clock2ddiagram.ComputedPrefix)
	case "IsExpanded":
		res = __gong__marshallBool(ident, "IsExpanded", clock2ddiagram.IsExpanded)

	default:
		log.Panicf("Unknown field %s for Gongstruct Clock2DDiagram", fieldName)
	}
	return
}

func (clock3ddiagram *Clock3DDiagram) GongMarshallField(stage *Stage, fieldName string) (res string) {
	ident := clock3ddiagram.GongGetIdentifier(stage)
	_ = ident
	switch fieldName {
	case "Name":
		res = __gong__marshallString(ident, "Name", clock3ddiagram.Name)
	case "IsHiddenClockTopCurveShape":
		res = __gong__marshallBool(ident, "IsHiddenClockTopCurveShape", clock3ddiagram.IsHiddenClockTopCurveShape)
	case "IsHiddenTorus3DShape":
		res = __gong__marshallBool(ident, "IsHiddenTorus3DShape", clock3ddiagram.IsHiddenTorus3DShape)
	case "IsHiddenSampledPoints3DShape":
		res = __gong__marshallBool(ident, "IsHiddenSampledPoints3DShape", clock3ddiagram.IsHiddenSampledPoints3DShape)
	case "IsHiddenTiledFloor3DShape":
		res = __gong__marshallBool(ident, "IsHiddenTiledFloor3DShape", clock3ddiagram.IsHiddenTiledFloor3DShape)
	case "IsChecked":
		res = __gong__marshallBool(ident, "IsChecked", clock3ddiagram.IsChecked)
	case "ComputedPrefix":
		res = __gong__marshallString(ident, "ComputedPrefix", clock3ddiagram.ComputedPrefix)
	case "IsExpanded":
		res = __gong__marshallBool(ident, "IsExpanded", clock3ddiagram.IsExpanded)

	case "SampledPoints3DShape":
		if clock3ddiagram.SampledPoints3DShape != nil {
			res = __gong__marshallPointer(ident, "SampledPoints3DShape", clock3ddiagram.SampledPoints3DShape.GongGetIdentifier(stage))
		} else {
			res = __gong__marshallPointer(ident, "SampledPoints3DShape", "nil")
		}
	case "Rendered3DShape":
		if clock3ddiagram.Rendered3DShape != nil {
			res = __gong__marshallPointer(ident, "Rendered3DShape", clock3ddiagram.Rendered3DShape.GongGetIdentifier(stage))
		} else {
			res = __gong__marshallPointer(ident, "Rendered3DShape", "nil")
		}
	default:
		log.Panicf("Unknown field %s for Gongstruct Clock3DDiagram", fieldName)
	}
	return
}

func (clocktopcurveshape *ClockTopCurveShape) GongMarshallField(stage *Stage, fieldName string) (res string) {
	ident := clocktopcurveshape.GongGetIdentifier(stage)
	_ = ident
	switch fieldName {
	case "Name":
		res = __gong__marshallString(ident, "Name", clocktopcurveshape.Name)

	default:
		log.Panicf("Unknown field %s for Gongstruct ClockTopCurveShape", fieldName)
	}
	return
}

func (cutline3dshape *CutLine3DShape) GongMarshallField(stage *Stage, fieldName string) (res string) {
	ident := cutline3dshape.GongGetIdentifier(stage)
	_ = ident
	switch fieldName {
	case "Name":
		res = __gong__marshallString(ident, "Name", cutline3dshape.Name)

	default:
		log.Panicf("Unknown field %s for Gongstruct CutLine3DShape", fieldName)
	}
	return
}

func (endarcshape *EndArcShape) GongMarshallField(stage *Stage, fieldName string) (res string) {
	ident := endarcshape.GongGetIdentifier(stage)
	_ = ident
	switch fieldName {
	case "Name":
		res = __gong__marshallString(ident, "Name", endarcshape.Name)
	case "StartX":
		res = __gong__marshallFloat(ident, "StartX", endarcshape.StartX)
	case "StartY":
		res = __gong__marshallFloat(ident, "StartY", endarcshape.StartY)
	case "EndX":
		res = __gong__marshallFloat(ident, "EndX", endarcshape.EndX)
	case "EndY":
		res = __gong__marshallFloat(ident, "EndY", endarcshape.EndY)
	case "XAxisRotation":
		res = __gong__marshallFloat(ident, "XAxisRotation", endarcshape.XAxisRotation)
	case "LargeArcFlag":
		res = __gong__marshallBool(ident, "LargeArcFlag", endarcshape.LargeArcFlag)
	case "SweepFlag":
		res = __gong__marshallBool(ident, "SweepFlag", endarcshape.SweepFlag)
	case "RadiusX":
		res = __gong__marshallFloat(ident, "RadiusX", endarcshape.RadiusX)
	case "RadiusY":
		res = __gong__marshallFloat(ident, "RadiusY", endarcshape.RadiusY)

	default:
		log.Panicf("Unknown field %s for Gongstruct EndArcShape", fieldName)
	}
	return
}

func (endarcshapegrid *EndArcShapeGrid) GongMarshallField(stage *Stage, fieldName string) (res string) {
	ident := endarcshapegrid.GongGetIdentifier(stage)
	_ = ident
	switch fieldName {
	case "Name":
		res = __gong__marshallString(ident, "Name", endarcshapegrid.Name)

	default:
		log.Panicf("Unknown field %s for Gongstruct EndArcShapeGrid", fieldName)
	}
	return
}

func (endhalfwayarcshape *EndHalfwayArcShape) GongMarshallField(stage *Stage, fieldName string) (res string) {
	ident := endhalfwayarcshape.GongGetIdentifier(stage)
	_ = ident
	switch fieldName {
	case "Name":
		res = __gong__marshallString(ident, "Name", endhalfwayarcshape.Name)
	case "StartX":
		res = __gong__marshallFloat(ident, "StartX", endhalfwayarcshape.StartX)
	case "StartY":
		res = __gong__marshallFloat(ident, "StartY", endhalfwayarcshape.StartY)
	case "EndX":
		res = __gong__marshallFloat(ident, "EndX", endhalfwayarcshape.EndX)
	case "EndY":
		res = __gong__marshallFloat(ident, "EndY", endhalfwayarcshape.EndY)
	case "RadiusX":
		res = __gong__marshallFloat(ident, "RadiusX", endhalfwayarcshape.RadiusX)
	case "RadiusY":
		res = __gong__marshallFloat(ident, "RadiusY", endhalfwayarcshape.RadiusY)
	case "XAxisRotation":
		res = __gong__marshallFloat(ident, "XAxisRotation", endhalfwayarcshape.XAxisRotation)
	case "LargeArcFlag":
		res = __gong__marshallBool(ident, "LargeArcFlag", endhalfwayarcshape.LargeArcFlag)
	case "SweepFlag":
		res = __gong__marshallBool(ident, "SweepFlag", endhalfwayarcshape.SweepFlag)

	default:
		log.Panicf("Unknown field %s for Gongstruct EndHalfwayArcShape", fieldName)
	}
	return
}

func (endhalfwayarcshapegrid *EndHalfwayArcShapeGrid) GongMarshallField(stage *Stage, fieldName string) (res string) {
	ident := endhalfwayarcshapegrid.GongGetIdentifier(stage)
	_ = ident
	switch fieldName {
	case "Name":
		res = __gong__marshallString(ident, "Name", endhalfwayarcshapegrid.Name)

	default:
		log.Panicf("Unknown field %s for Gongstruct EndHalfwayArcShapeGrid", fieldName)
	}
	return
}

func (explanationtextshape *ExplanationTextShape) GongMarshallField(stage *Stage, fieldName string) (res string) {
	ident := explanationtextshape.GongGetIdentifier(stage)
	_ = ident
	switch fieldName {
	case "Name":
		res = __gong__marshallString(ident, "Name", explanationtextshape.Name)

	default:
		log.Panicf("Unknown field %s for Gongstruct ExplanationTextShape", fieldName)
	}
	return
}

func (eye3dshape *Eye3DShape) GongMarshallField(stage *Stage, fieldName string) (res string) {
	ident := eye3dshape.GongGetIdentifier(stage)
	_ = ident
	switch fieldName {
	case "Name":
		res = __gong__marshallString(ident, "Name", eye3dshape.Name)

	default:
		log.Panicf("Unknown field %s for Gongstruct Eye3DShape", fieldName)
	}
	return
}

func (eyecornerssampledpoints3dshape *EyeCornersSampledPoints3DShape) GongMarshallField(stage *Stage, fieldName string) (res string) {
	ident := eyecornerssampledpoints3dshape.GongGetIdentifier(stage)
	_ = ident
	switch fieldName {
	case "Name":
		res = __gong__marshallString(ident, "Name", eyecornerssampledpoints3dshape.Name)

	default:
		log.Panicf("Unknown field %s for Gongstruct EyeCornersSampledPoints3DShape", fieldName)
	}
	return
}

func (eyesampledpoints3dshape *EyeSampledPoints3DShape) GongMarshallField(stage *Stage, fieldName string) (res string) {
	ident := eyesampledpoints3dshape.GongGetIdentifier(stage)
	_ = ident
	switch fieldName {
	case "Name":
		res = __gong__marshallString(ident, "Name", eyesampledpoints3dshape.Name)

	default:
		log.Panicf("Unknown field %s for Gongstruct EyeSampledPoints3DShape", fieldName)
	}
	return
}

func (eyeseatbottomcurveshape *EyeSeatBottomCurveShape) GongMarshallField(stage *Stage, fieldName string) (res string) {
	ident := eyeseatbottomcurveshape.GongGetIdentifier(stage)
	_ = ident
	switch fieldName {
	case "Name":
		res = __gong__marshallString(ident, "Name", eyeseatbottomcurveshape.Name)

	default:
		log.Panicf("Unknown field %s for Gongstruct EyeSeatBottomCurveShape", fieldName)
	}
	return
}

func (eyestoolbottomcurveshape *EyeStoolBottomCurveShape) GongMarshallField(stage *Stage, fieldName string) (res string) {
	ident := eyestoolbottomcurveshape.GongGetIdentifier(stage)
	_ = ident
	switch fieldName {
	case "Name":
		res = __gong__marshallString(ident, "Name", eyestoolbottomcurveshape.Name)

	default:
		log.Panicf("Unknown field %s for Gongstruct EyeStoolBottomCurveShape", fieldName)
	}
	return
}

func (eyevolume3dshape *EyeVolume3DShape) GongMarshallField(stage *Stage, fieldName string) (res string) {
	ident := eyevolume3dshape.GongGetIdentifier(stage)
	_ = ident
	switch fieldName {
	case "Name":
		res = __gong__marshallString(ident, "Name", eyevolume3dshape.Name)

	default:
		log.Panicf("Unknown field %s for Gongstruct EyeVolume3DShape", fieldName)
	}
	return
}

func (gridpathshape *GridPathShape) GongMarshallField(stage *Stage, fieldName string) (res string) {
	ident := gridpathshape.GongGetIdentifier(stage)
	_ = ident
	switch fieldName {
	case "Name":
		res = __gong__marshallString(ident, "Name", gridpathshape.Name)

	default:
		log.Panicf("Unknown field %s for Gongstruct GridPathShape", fieldName)
	}
	return
}

func (growthcurve2d *GrowthCurve2D) GongMarshallField(stage *Stage, fieldName string) (res string) {
	ident := growthcurve2d.GongGetIdentifier(stage)
	_ = ident
	switch fieldName {
	case "Name":
		res = __gong__marshallString(ident, "Name", growthcurve2d.Name)

	default:
		log.Panicf("Unknown field %s for Gongstruct GrowthCurve2D", fieldName)
	}
	return
}

func (growthcurve2dribbon *GrowthCurve2DRibbon) GongMarshallField(stage *Stage, fieldName string) (res string) {
	ident := growthcurve2dribbon.GongGetIdentifier(stage)
	_ = ident
	switch fieldName {
	case "Name":
		res = __gong__marshallString(ident, "Name", growthcurve2dribbon.Name)

	default:
		log.Panicf("Unknown field %s for Gongstruct GrowthCurve2DRibbon", fieldName)
	}
	return
}

func (growthcurve2dribbonendshape *GrowthCurve2DRibbonEndShape) GongMarshallField(stage *Stage, fieldName string) (res string) {
	ident := growthcurve2dribbonendshape.GongGetIdentifier(stage)
	_ = ident
	switch fieldName {
	case "Name":
		res = __gong__marshallString(ident, "Name", growthcurve2dribbonendshape.Name)
	case "BottomStartX":
		res = __gong__marshallFloat(ident, "BottomStartX", growthcurve2dribbonendshape.BottomStartX)
	case "BottomStartY":
		res = __gong__marshallFloat(ident, "BottomStartY", growthcurve2dribbonendshape.BottomStartY)
	case "BottomEndX":
		res = __gong__marshallFloat(ident, "BottomEndX", growthcurve2dribbonendshape.BottomEndX)
	case "BottomEndY":
		res = __gong__marshallFloat(ident, "BottomEndY", growthcurve2dribbonendshape.BottomEndY)
	case "BottomRadiusX":
		res = __gong__marshallFloat(ident, "BottomRadiusX", growthcurve2dribbonendshape.BottomRadiusX)
	case "BottomRadiusY":
		res = __gong__marshallFloat(ident, "BottomRadiusY", growthcurve2dribbonendshape.BottomRadiusY)
	case "BottomXAxisRotation":
		res = __gong__marshallFloat(ident, "BottomXAxisRotation", growthcurve2dribbonendshape.BottomXAxisRotation)
	case "BottomLargeArcFlag":
		res = __gong__marshallBool(ident, "BottomLargeArcFlag", growthcurve2dribbonendshape.BottomLargeArcFlag)
	case "BottomSweepFlag":
		res = __gong__marshallBool(ident, "BottomSweepFlag", growthcurve2dribbonendshape.BottomSweepFlag)
	case "TopStartX":
		res = __gong__marshallFloat(ident, "TopStartX", growthcurve2dribbonendshape.TopStartX)
	case "TopStartY":
		res = __gong__marshallFloat(ident, "TopStartY", growthcurve2dribbonendshape.TopStartY)
	case "TopEndX":
		res = __gong__marshallFloat(ident, "TopEndX", growthcurve2dribbonendshape.TopEndX)
	case "TopEndY":
		res = __gong__marshallFloat(ident, "TopEndY", growthcurve2dribbonendshape.TopEndY)
	case "TopRadiusX":
		res = __gong__marshallFloat(ident, "TopRadiusX", growthcurve2dribbonendshape.TopRadiusX)
	case "TopRadiusY":
		res = __gong__marshallFloat(ident, "TopRadiusY", growthcurve2dribbonendshape.TopRadiusY)
	case "TopXAxisRotation":
		res = __gong__marshallFloat(ident, "TopXAxisRotation", growthcurve2dribbonendshape.TopXAxisRotation)
	case "TopLargeArcFlag":
		res = __gong__marshallBool(ident, "TopLargeArcFlag", growthcurve2dribbonendshape.TopLargeArcFlag)
	case "TopSweepFlag":
		res = __gong__marshallBool(ident, "TopSweepFlag", growthcurve2dribbonendshape.TopSweepFlag)

	default:
		log.Panicf("Unknown field %s for Gongstruct GrowthCurve2DRibbonEndShape", fieldName)
	}
	return
}

func (growthcurve2dribbonstartshape *GrowthCurve2DRibbonStartShape) GongMarshallField(stage *Stage, fieldName string) (res string) {
	ident := growthcurve2dribbonstartshape.GongGetIdentifier(stage)
	_ = ident
	switch fieldName {
	case "Name":
		res = __gong__marshallString(ident, "Name", growthcurve2dribbonstartshape.Name)
	case "BottomStartX":
		res = __gong__marshallFloat(ident, "BottomStartX", growthcurve2dribbonstartshape.BottomStartX)
	case "BottomStartY":
		res = __gong__marshallFloat(ident, "BottomStartY", growthcurve2dribbonstartshape.BottomStartY)
	case "BottomEndX":
		res = __gong__marshallFloat(ident, "BottomEndX", growthcurve2dribbonstartshape.BottomEndX)
	case "BottomEndY":
		res = __gong__marshallFloat(ident, "BottomEndY", growthcurve2dribbonstartshape.BottomEndY)
	case "BottomRadiusX":
		res = __gong__marshallFloat(ident, "BottomRadiusX", growthcurve2dribbonstartshape.BottomRadiusX)
	case "BottomRadiusY":
		res = __gong__marshallFloat(ident, "BottomRadiusY", growthcurve2dribbonstartshape.BottomRadiusY)
	case "BottomXAxisRotation":
		res = __gong__marshallFloat(ident, "BottomXAxisRotation", growthcurve2dribbonstartshape.BottomXAxisRotation)
	case "BottomLargeArcFlag":
		res = __gong__marshallBool(ident, "BottomLargeArcFlag", growthcurve2dribbonstartshape.BottomLargeArcFlag)
	case "BottomSweepFlag":
		res = __gong__marshallBool(ident, "BottomSweepFlag", growthcurve2dribbonstartshape.BottomSweepFlag)
	case "TopStartX":
		res = __gong__marshallFloat(ident, "TopStartX", growthcurve2dribbonstartshape.TopStartX)
	case "TopStartY":
		res = __gong__marshallFloat(ident, "TopStartY", growthcurve2dribbonstartshape.TopStartY)
	case "TopEndX":
		res = __gong__marshallFloat(ident, "TopEndX", growthcurve2dribbonstartshape.TopEndX)
	case "TopEndY":
		res = __gong__marshallFloat(ident, "TopEndY", growthcurve2dribbonstartshape.TopEndY)
	case "TopRadiusX":
		res = __gong__marshallFloat(ident, "TopRadiusX", growthcurve2dribbonstartshape.TopRadiusX)
	case "TopRadiusY":
		res = __gong__marshallFloat(ident, "TopRadiusY", growthcurve2dribbonstartshape.TopRadiusY)
	case "TopXAxisRotation":
		res = __gong__marshallFloat(ident, "TopXAxisRotation", growthcurve2dribbonstartshape.TopXAxisRotation)
	case "TopLargeArcFlag":
		res = __gong__marshallBool(ident, "TopLargeArcFlag", growthcurve2dribbonstartshape.TopLargeArcFlag)
	case "TopSweepFlag":
		res = __gong__marshallBool(ident, "TopSweepFlag", growthcurve2dribbonstartshape.TopSweepFlag)

	default:
		log.Panicf("Unknown field %s for Gongstruct GrowthCurve2DRibbonStartShape", fieldName)
	}
	return
}

func (growthcurverhombusgridshape *GrowthCurveRhombusGridShape) GongMarshallField(stage *Stage, fieldName string) (res string) {
	ident := growthcurverhombusgridshape.GongGetIdentifier(stage)
	_ = ident
	switch fieldName {
	case "Name":
		res = __gong__marshallString(ident, "Name", growthcurverhombusgridshape.Name)

	default:
		log.Panicf("Unknown field %s for Gongstruct GrowthCurveRhombusGridShape", fieldName)
	}
	return
}

func (growthcurverhombusshape *GrowthCurveRhombusShape) GongMarshallField(stage *Stage, fieldName string) (res string) {
	ident := growthcurverhombusshape.GongGetIdentifier(stage)
	_ = ident
	switch fieldName {
	case "Name":
		res = __gong__marshallString(ident, "Name", growthcurverhombusshape.Name)
	case "X":
		res = __gong__marshallFloat(ident, "X", growthcurverhombusshape.X)
	case "Y":
		res = __gong__marshallFloat(ident, "Y", growthcurverhombusshape.Y)

	default:
		log.Panicf("Unknown field %s for Gongstruct GrowthCurveRhombusShape", fieldName)
	}
	return
}

func (growthvectorshape *GrowthVectorShape) GongMarshallField(stage *Stage, fieldName string) (res string) {
	ident := growthvectorshape.GongGetIdentifier(stage)
	_ = ident
	switch fieldName {
	case "Name":
		res = __gong__marshallString(ident, "Name", growthvectorshape.Name)
	case "X":
		res = __gong__marshallFloat(ident, "X", growthvectorshape.X)
	case "Y":
		res = __gong__marshallFloat(ident, "Y", growthvectorshape.Y)

	default:
		log.Panicf("Unknown field %s for Gongstruct GrowthVectorShape", fieldName)
	}
	return
}

func (initialrhombusgridshape *InitialRhombusGridShape) GongMarshallField(stage *Stage, fieldName string) (res string) {
	ident := initialrhombusgridshape.GongGetIdentifier(stage)
	_ = ident
	switch fieldName {
	case "Name":
		res = __gong__marshallString(ident, "Name", initialrhombusgridshape.Name)

	default:
		log.Panicf("Unknown field %s for Gongstruct InitialRhombusGridShape", fieldName)
	}
	return
}

func (initialrhombusshape *InitialRhombusShape) GongMarshallField(stage *Stage, fieldName string) (res string) {
	ident := initialrhombusshape.GongGetIdentifier(stage)
	_ = ident
	switch fieldName {
	case "Name":
		res = __gong__marshallString(ident, "Name", initialrhombusshape.Name)
	case "X":
		res = __gong__marshallFloat(ident, "X", initialrhombusshape.X)
	case "Y":
		res = __gong__marshallFloat(ident, "Y", initialrhombusshape.Y)

	default:
		log.Panicf("Unknown field %s for Gongstruct InitialRhombusShape", fieldName)
	}
	return
}

func (key3dshape *Key3DShape) GongMarshallField(stage *Stage, fieldName string) (res string) {
	ident := key3dshape.GongGetIdentifier(stage)
	_ = ident
	switch fieldName {
	case "Name":
		res = __gong__marshallString(ident, "Name", key3dshape.Name)

	default:
		log.Panicf("Unknown field %s for Gongstruct Key3DShape", fieldName)
	}
	return
}

func (keyhole3dshape *KeyHole3DShape) GongMarshallField(stage *Stage, fieldName string) (res string) {
	ident := keyhole3dshape.GongGetIdentifier(stage)
	_ = ident
	switch fieldName {
	case "Name":
		res = __gong__marshallString(ident, "Name", keyhole3dshape.Name)

	default:
		log.Panicf("Unknown field %s for Gongstruct KeyHole3DShape", fieldName)
	}
	return
}

func (keyholeshape *KeyHoleShape) GongMarshallField(stage *Stage, fieldName string) (res string) {
	ident := keyholeshape.GongGetIdentifier(stage)
	_ = ident
	switch fieldName {
	case "Name":
		res = __gong__marshallString(ident, "Name", keyholeshape.Name)
	case "X":
		res = __gong__marshallFloat(ident, "X", keyholeshape.X)
	case "Y":
		res = __gong__marshallFloat(ident, "Y", keyholeshape.Y)
	case "Width":
		res = __gong__marshallFloat(ident, "Width", keyholeshape.Width)
	case "Height":
		res = __gong__marshallFloat(ident, "Height", keyholeshape.Height)

	default:
		log.Panicf("Unknown field %s for Gongstruct KeyHoleShape", fieldName)
	}
	return
}

func (leaves3dshape *Leaves3DShape) GongMarshallField(stage *Stage, fieldName string) (res string) {
	ident := leaves3dshape.GongGetIdentifier(stage)
	_ = ident
	switch fieldName {
	case "Name":
		res = __gong__marshallString(ident, "Name", leaves3dshape.Name)

	default:
		log.Panicf("Unknown field %s for Gongstruct Leaves3DShape", fieldName)
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

	case "Plants":
		var sb strings.Builder
		for _, _plantabstract := range library.Plants {
			sb.WriteString(__gong__marshallSliceOfPointers(ident, "Plants", _plantabstract.GongGetIdentifier(stage)))
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

func (midarcvectorshape *MidArcVectorShape) GongMarshallField(stage *Stage, fieldName string) (res string) {
	ident := midarcvectorshape.GongGetIdentifier(stage)
	_ = ident
	switch fieldName {
	case "Name":
		res = __gong__marshallString(ident, "Name", midarcvectorshape.Name)
	case "StartX":
		res = __gong__marshallFloat(ident, "StartX", midarcvectorshape.StartX)
	case "StartY":
		res = __gong__marshallFloat(ident, "StartY", midarcvectorshape.StartY)
	case "EndX":
		res = __gong__marshallFloat(ident, "EndX", midarcvectorshape.EndX)
	case "EndY":
		res = __gong__marshallFloat(ident, "EndY", midarcvectorshape.EndY)

	default:
		log.Panicf("Unknown field %s for Gongstruct MidArcVectorShape", fieldName)
	}
	return
}

func (midarcvectorshapegrid *MidArcVectorShapeGrid) GongMarshallField(stage *Stage, fieldName string) (res string) {
	ident := midarcvectorshapegrid.GongGetIdentifier(stage)
	_ = ident
	switch fieldName {
	case "Name":
		res = __gong__marshallString(ident, "Name", midarcvectorshapegrid.Name)

	default:
		log.Panicf("Unknown field %s for Gongstruct MidArcVectorShapeGrid", fieldName)
	}
	return
}

func (originalpoints3dshape *OriginalPoints3DShape) GongMarshallField(stage *Stage, fieldName string) (res string) {
	ident := originalpoints3dshape.GongGetIdentifier(stage)
	_ = ident
	switch fieldName {
	case "Name":
		res = __gong__marshallString(ident, "Name", originalpoints3dshape.Name)

	default:
		log.Panicf("Unknown field %s for Gongstruct OriginalPoints3DShape", fieldName)
	}
	return
}

func (parastichymcurves3dshape *ParastichyMCurves3DShape) GongMarshallField(stage *Stage, fieldName string) (res string) {
	ident := parastichymcurves3dshape.GongGetIdentifier(stage)
	_ = ident
	switch fieldName {
	case "Name":
		res = __gong__marshallString(ident, "Name", parastichymcurves3dshape.Name)

	default:
		log.Panicf("Unknown field %s for Gongstruct ParastichyMCurves3DShape", fieldName)
	}
	return
}

func (parastichyncurves3dshape *ParastichyNCurves3DShape) GongMarshallField(stage *Stage, fieldName string) (res string) {
	ident := parastichyncurves3dshape.GongGetIdentifier(stage)
	_ = ident
	switch fieldName {
	case "Name":
		res = __gong__marshallString(ident, "Name", parastichyncurves3dshape.Name)

	default:
		log.Panicf("Unknown field %s for Gongstruct ParastichyNCurves3DShape", fieldName)
	}
	return
}

func (partiallygrowthcurve2dribbon *PartiallyGrowthCurve2DRibbon) GongMarshallField(stage *Stage, fieldName string) (res string) {
	ident := partiallygrowthcurve2dribbon.GongGetIdentifier(stage)
	_ = ident
	switch fieldName {
	case "Name":
		res = __gong__marshallString(ident, "Name", partiallygrowthcurve2dribbon.Name)

	default:
		log.Panicf("Unknown field %s for Gongstruct PartiallyGrowthCurve2DRibbon", fieldName)
	}
	return
}

func (partiallygrowthcurve2dribbonendshape *PartiallyGrowthCurve2DRibbonEndShape) GongMarshallField(stage *Stage, fieldName string) (res string) {
	ident := partiallygrowthcurve2dribbonendshape.GongGetIdentifier(stage)
	_ = ident
	switch fieldName {
	case "Name":
		res = __gong__marshallString(ident, "Name", partiallygrowthcurve2dribbonendshape.Name)
	case "BottomStartX":
		res = __gong__marshallFloat(ident, "BottomStartX", partiallygrowthcurve2dribbonendshape.BottomStartX)
	case "BottomStartY":
		res = __gong__marshallFloat(ident, "BottomStartY", partiallygrowthcurve2dribbonendshape.BottomStartY)
	case "BottomEndX":
		res = __gong__marshallFloat(ident, "BottomEndX", partiallygrowthcurve2dribbonendshape.BottomEndX)
	case "BottomEndY":
		res = __gong__marshallFloat(ident, "BottomEndY", partiallygrowthcurve2dribbonendshape.BottomEndY)
	case "BottomRadiusX":
		res = __gong__marshallFloat(ident, "BottomRadiusX", partiallygrowthcurve2dribbonendshape.BottomRadiusX)
	case "BottomRadiusY":
		res = __gong__marshallFloat(ident, "BottomRadiusY", partiallygrowthcurve2dribbonendshape.BottomRadiusY)
	case "BottomXAxisRotation":
		res = __gong__marshallFloat(ident, "BottomXAxisRotation", partiallygrowthcurve2dribbonendshape.BottomXAxisRotation)
	case "BottomLargeArcFlag":
		res = __gong__marshallBool(ident, "BottomLargeArcFlag", partiallygrowthcurve2dribbonendshape.BottomLargeArcFlag)
	case "BottomSweepFlag":
		res = __gong__marshallBool(ident, "BottomSweepFlag", partiallygrowthcurve2dribbonendshape.BottomSweepFlag)
	case "TopStartX":
		res = __gong__marshallFloat(ident, "TopStartX", partiallygrowthcurve2dribbonendshape.TopStartX)
	case "TopStartY":
		res = __gong__marshallFloat(ident, "TopStartY", partiallygrowthcurve2dribbonendshape.TopStartY)
	case "TopEndX":
		res = __gong__marshallFloat(ident, "TopEndX", partiallygrowthcurve2dribbonendshape.TopEndX)
	case "TopEndY":
		res = __gong__marshallFloat(ident, "TopEndY", partiallygrowthcurve2dribbonendshape.TopEndY)
	case "TopRadiusX":
		res = __gong__marshallFloat(ident, "TopRadiusX", partiallygrowthcurve2dribbonendshape.TopRadiusX)
	case "TopRadiusY":
		res = __gong__marshallFloat(ident, "TopRadiusY", partiallygrowthcurve2dribbonendshape.TopRadiusY)
	case "TopXAxisRotation":
		res = __gong__marshallFloat(ident, "TopXAxisRotation", partiallygrowthcurve2dribbonendshape.TopXAxisRotation)
	case "TopLargeArcFlag":
		res = __gong__marshallBool(ident, "TopLargeArcFlag", partiallygrowthcurve2dribbonendshape.TopLargeArcFlag)
	case "TopSweepFlag":
		res = __gong__marshallBool(ident, "TopSweepFlag", partiallygrowthcurve2dribbonendshape.TopSweepFlag)

	default:
		log.Panicf("Unknown field %s for Gongstruct PartiallyGrowthCurve2DRibbonEndShape", fieldName)
	}
	return
}

func (partiallygrowthcurve2dribbonstartshape *PartiallyGrowthCurve2DRibbonStartShape) GongMarshallField(stage *Stage, fieldName string) (res string) {
	ident := partiallygrowthcurve2dribbonstartshape.GongGetIdentifier(stage)
	_ = ident
	switch fieldName {
	case "Name":
		res = __gong__marshallString(ident, "Name", partiallygrowthcurve2dribbonstartshape.Name)
	case "BottomStartX":
		res = __gong__marshallFloat(ident, "BottomStartX", partiallygrowthcurve2dribbonstartshape.BottomStartX)
	case "BottomStartY":
		res = __gong__marshallFloat(ident, "BottomStartY", partiallygrowthcurve2dribbonstartshape.BottomStartY)
	case "BottomEndX":
		res = __gong__marshallFloat(ident, "BottomEndX", partiallygrowthcurve2dribbonstartshape.BottomEndX)
	case "BottomEndY":
		res = __gong__marshallFloat(ident, "BottomEndY", partiallygrowthcurve2dribbonstartshape.BottomEndY)
	case "BottomRadiusX":
		res = __gong__marshallFloat(ident, "BottomRadiusX", partiallygrowthcurve2dribbonstartshape.BottomRadiusX)
	case "BottomRadiusY":
		res = __gong__marshallFloat(ident, "BottomRadiusY", partiallygrowthcurve2dribbonstartshape.BottomRadiusY)
	case "BottomXAxisRotation":
		res = __gong__marshallFloat(ident, "BottomXAxisRotation", partiallygrowthcurve2dribbonstartshape.BottomXAxisRotation)
	case "BottomLargeArcFlag":
		res = __gong__marshallBool(ident, "BottomLargeArcFlag", partiallygrowthcurve2dribbonstartshape.BottomLargeArcFlag)
	case "BottomSweepFlag":
		res = __gong__marshallBool(ident, "BottomSweepFlag", partiallygrowthcurve2dribbonstartshape.BottomSweepFlag)
	case "TopStartX":
		res = __gong__marshallFloat(ident, "TopStartX", partiallygrowthcurve2dribbonstartshape.TopStartX)
	case "TopStartY":
		res = __gong__marshallFloat(ident, "TopStartY", partiallygrowthcurve2dribbonstartshape.TopStartY)
	case "TopEndX":
		res = __gong__marshallFloat(ident, "TopEndX", partiallygrowthcurve2dribbonstartshape.TopEndX)
	case "TopEndY":
		res = __gong__marshallFloat(ident, "TopEndY", partiallygrowthcurve2dribbonstartshape.TopEndY)
	case "TopRadiusX":
		res = __gong__marshallFloat(ident, "TopRadiusX", partiallygrowthcurve2dribbonstartshape.TopRadiusX)
	case "TopRadiusY":
		res = __gong__marshallFloat(ident, "TopRadiusY", partiallygrowthcurve2dribbonstartshape.TopRadiusY)
	case "TopXAxisRotation":
		res = __gong__marshallFloat(ident, "TopXAxisRotation", partiallygrowthcurve2dribbonstartshape.TopXAxisRotation)
	case "TopLargeArcFlag":
		res = __gong__marshallBool(ident, "TopLargeArcFlag", partiallygrowthcurve2dribbonstartshape.TopLargeArcFlag)
	case "TopSweepFlag":
		res = __gong__marshallBool(ident, "TopSweepFlag", partiallygrowthcurve2dribbonstartshape.TopSweepFlag)

	default:
		log.Panicf("Unknown field %s for Gongstruct PartiallyGrowthCurve2DRibbonStartShape", fieldName)
	}
	return
}

func (partiallygrowthcurve2dtrajectory *PartiallyGrowthCurve2DTrajectory) GongMarshallField(stage *Stage, fieldName string) (res string) {
	ident := partiallygrowthcurve2dtrajectory.GongGetIdentifier(stage)
	_ = ident
	switch fieldName {
	case "Name":
		res = __gong__marshallString(ident, "Name", partiallygrowthcurve2dtrajectory.Name)

	default:
		log.Panicf("Unknown field %s for Gongstruct PartiallyGrowthCurve2DTrajectory", fieldName)
	}
	return
}

func (partiallygrowthcurve2dtrajectoryp1curveshape *PartiallyGrowthCurve2DTrajectoryP1CurveShape) GongMarshallField(stage *Stage, fieldName string) (res string) {
	ident := partiallygrowthcurve2dtrajectoryp1curveshape.GongGetIdentifier(stage)
	_ = ident
	switch fieldName {
	case "Name":
		res = __gong__marshallString(ident, "Name", partiallygrowthcurve2dtrajectoryp1curveshape.Name)
	case "StartX":
		res = __gong__marshallFloat(ident, "StartX", partiallygrowthcurve2dtrajectoryp1curveshape.StartX)
	case "StartY":
		res = __gong__marshallFloat(ident, "StartY", partiallygrowthcurve2dtrajectoryp1curveshape.StartY)
	case "EndX":
		res = __gong__marshallFloat(ident, "EndX", partiallygrowthcurve2dtrajectoryp1curveshape.EndX)
	case "EndY":
		res = __gong__marshallFloat(ident, "EndY", partiallygrowthcurve2dtrajectoryp1curveshape.EndY)

	default:
		log.Panicf("Unknown field %s for Gongstruct PartiallyGrowthCurve2DTrajectoryP1CurveShape", fieldName)
	}
	return
}

func (partiallygrowthcurve2dtrajectoryp1p2 *PartiallyGrowthCurve2DTrajectoryP1P2) GongMarshallField(stage *Stage, fieldName string) (res string) {
	ident := partiallygrowthcurve2dtrajectoryp1p2.GongGetIdentifier(stage)
	_ = ident
	switch fieldName {
	case "Name":
		res = __gong__marshallString(ident, "Name", partiallygrowthcurve2dtrajectoryp1p2.Name)

	default:
		log.Panicf("Unknown field %s for Gongstruct PartiallyGrowthCurve2DTrajectoryP1P2", fieldName)
	}
	return
}

func (partiallygrowthcurve2dtrajectoryp1p2pairlineshape *PartiallyGrowthCurve2DTrajectoryP1P2PairLineShape) GongMarshallField(stage *Stage, fieldName string) (res string) {
	ident := partiallygrowthcurve2dtrajectoryp1p2pairlineshape.GongGetIdentifier(stage)
	_ = ident
	switch fieldName {
	case "Name":
		res = __gong__marshallString(ident, "Name", partiallygrowthcurve2dtrajectoryp1p2pairlineshape.Name)
	case "StartX":
		res = __gong__marshallFloat(ident, "StartX", partiallygrowthcurve2dtrajectoryp1p2pairlineshape.StartX)
	case "StartY":
		res = __gong__marshallFloat(ident, "StartY", partiallygrowthcurve2dtrajectoryp1p2pairlineshape.StartY)
	case "EndX":
		res = __gong__marshallFloat(ident, "EndX", partiallygrowthcurve2dtrajectoryp1p2pairlineshape.EndX)
	case "EndY":
		res = __gong__marshallFloat(ident, "EndY", partiallygrowthcurve2dtrajectoryp1p2pairlineshape.EndY)

	default:
		log.Panicf("Unknown field %s for Gongstruct PartiallyGrowthCurve2DTrajectoryP1P2PairLineShape", fieldName)
	}
	return
}

func (partiallygrowthcurve2dtrajectoryp1pointshape *PartiallyGrowthCurve2DTrajectoryP1PointShape) GongMarshallField(stage *Stage, fieldName string) (res string) {
	ident := partiallygrowthcurve2dtrajectoryp1pointshape.GongGetIdentifier(stage)
	_ = ident
	switch fieldName {
	case "Name":
		res = __gong__marshallString(ident, "Name", partiallygrowthcurve2dtrajectoryp1pointshape.Name)
	case "X":
		res = __gong__marshallFloat(ident, "X", partiallygrowthcurve2dtrajectoryp1pointshape.X)
	case "Y":
		res = __gong__marshallFloat(ident, "Y", partiallygrowthcurve2dtrajectoryp1pointshape.Y)

	default:
		log.Panicf("Unknown field %s for Gongstruct PartiallyGrowthCurve2DTrajectoryP1PointShape", fieldName)
	}
	return
}

func (partiallygrowthcurve2dtrajectoryp2curveshape *PartiallyGrowthCurve2DTrajectoryP2CurveShape) GongMarshallField(stage *Stage, fieldName string) (res string) {
	ident := partiallygrowthcurve2dtrajectoryp2curveshape.GongGetIdentifier(stage)
	_ = ident
	switch fieldName {
	case "Name":
		res = __gong__marshallString(ident, "Name", partiallygrowthcurve2dtrajectoryp2curveshape.Name)
	case "StartX":
		res = __gong__marshallFloat(ident, "StartX", partiallygrowthcurve2dtrajectoryp2curveshape.StartX)
	case "StartY":
		res = __gong__marshallFloat(ident, "StartY", partiallygrowthcurve2dtrajectoryp2curveshape.StartY)
	case "EndX":
		res = __gong__marshallFloat(ident, "EndX", partiallygrowthcurve2dtrajectoryp2curveshape.EndX)
	case "EndY":
		res = __gong__marshallFloat(ident, "EndY", partiallygrowthcurve2dtrajectoryp2curveshape.EndY)

	default:
		log.Panicf("Unknown field %s for Gongstruct PartiallyGrowthCurve2DTrajectoryP2CurveShape", fieldName)
	}
	return
}

func (partiallygrowthcurve2dtrajectoryp2pointshape *PartiallyGrowthCurve2DTrajectoryP2PointShape) GongMarshallField(stage *Stage, fieldName string) (res string) {
	ident := partiallygrowthcurve2dtrajectoryp2pointshape.GongGetIdentifier(stage)
	_ = ident
	switch fieldName {
	case "Name":
		res = __gong__marshallString(ident, "Name", partiallygrowthcurve2dtrajectoryp2pointshape.Name)
	case "X":
		res = __gong__marshallFloat(ident, "X", partiallygrowthcurve2dtrajectoryp2pointshape.X)
	case "Y":
		res = __gong__marshallFloat(ident, "Y", partiallygrowthcurve2dtrajectoryp2pointshape.Y)

	default:
		log.Panicf("Unknown field %s for Gongstruct PartiallyGrowthCurve2DTrajectoryP2PointShape", fieldName)
	}
	return
}

func (partiallygrowthcurve2dtrajectoryshape *PartiallyGrowthCurve2DTrajectoryShape) GongMarshallField(stage *Stage, fieldName string) (res string) {
	ident := partiallygrowthcurve2dtrajectoryshape.GongGetIdentifier(stage)
	_ = ident
	switch fieldName {
	case "Name":
		res = __gong__marshallString(ident, "Name", partiallygrowthcurve2dtrajectoryshape.Name)
	case "StartX":
		res = __gong__marshallFloat(ident, "StartX", partiallygrowthcurve2dtrajectoryshape.StartX)
	case "StartY":
		res = __gong__marshallFloat(ident, "StartY", partiallygrowthcurve2dtrajectoryshape.StartY)
	case "EndX":
		res = __gong__marshallFloat(ident, "EndX", partiallygrowthcurve2dtrajectoryshape.EndX)
	case "EndY":
		res = __gong__marshallFloat(ident, "EndY", partiallygrowthcurve2dtrajectoryshape.EndY)

	default:
		log.Panicf("Unknown field %s for Gongstruct PartiallyGrowthCurve2DTrajectoryShape", fieldName)
	}
	return
}

func (partiallyrotatedseatbottomcurveshape *PartiallyRotatedSeatBottomCurveShape) GongMarshallField(stage *Stage, fieldName string) (res string) {
	ident := partiallyrotatedseatbottomcurveshape.GongGetIdentifier(stage)
	_ = ident
	switch fieldName {
	case "Name":
		res = __gong__marshallString(ident, "Name", partiallyrotatedseatbottomcurveshape.Name)

	default:
		log.Panicf("Unknown field %s for Gongstruct PartiallyRotatedSeatBottomCurveShape", fieldName)
	}
	return
}

func (partiallyrotatedseattopcurveshape *PartiallyRotatedSeatTopCurveShape) GongMarshallField(stage *Stage, fieldName string) (res string) {
	ident := partiallyrotatedseattopcurveshape.GongGetIdentifier(stage)
	_ = ident
	switch fieldName {
	case "Name":
		res = __gong__marshallString(ident, "Name", partiallyrotatedseattopcurveshape.Name)

	default:
		log.Panicf("Unknown field %s for Gongstruct PartiallyRotatedSeatTopCurveShape", fieldName)
	}
	return
}

func (partiallyrotatedtorusshape *PartiallyRotatedTorusShape) GongMarshallField(stage *Stage, fieldName string) (res string) {
	ident := partiallyrotatedtorusshape.GongGetIdentifier(stage)
	_ = ident
	switch fieldName {
	case "Name":
		res = __gong__marshallString(ident, "Name", partiallyrotatedtorusshape.Name)

	default:
		log.Panicf("Unknown field %s for Gongstruct PartiallyRotatedTorusShape", fieldName)
	}
	return
}

func (perpendicularvector *PerpendicularVector) GongMarshallField(stage *Stage, fieldName string) (res string) {
	ident := perpendicularvector.GongGetIdentifier(stage)
	_ = ident
	switch fieldName {
	case "Name":
		res = __gong__marshallString(ident, "Name", perpendicularvector.Name)
	case "StartX":
		res = __gong__marshallFloat(ident, "StartX", perpendicularvector.StartX)
	case "StartY":
		res = __gong__marshallFloat(ident, "StartY", perpendicularvector.StartY)
	case "EndX":
		res = __gong__marshallFloat(ident, "EndX", perpendicularvector.EndX)
	case "EndY":
		res = __gong__marshallFloat(ident, "EndY", perpendicularvector.EndY)

	default:
		log.Panicf("Unknown field %s for Gongstruct PerpendicularVector", fieldName)
	}
	return
}

func (perpendicularvectorgrid *PerpendicularVectorGrid) GongMarshallField(stage *Stage, fieldName string) (res string) {
	ident := perpendicularvectorgrid.GongGetIdentifier(stage)
	_ = ident
	switch fieldName {
	case "Name":
		res = __gong__marshallString(ident, "Name", perpendicularvectorgrid.Name)

	default:
		log.Panicf("Unknown field %s for Gongstruct PerpendicularVectorGrid", fieldName)
	}
	return
}

func (perpendicularvectorgridhalfway *PerpendicularVectorGridHalfway) GongMarshallField(stage *Stage, fieldName string) (res string) {
	ident := perpendicularvectorgridhalfway.GongGetIdentifier(stage)
	_ = ident
	switch fieldName {
	case "Name":
		res = __gong__marshallString(ident, "Name", perpendicularvectorgridhalfway.Name)

	default:
		log.Panicf("Unknown field %s for Gongstruct PerpendicularVectorGridHalfway", fieldName)
	}
	return
}

func (perpendicularvectorhalfway *PerpendicularVectorHalfway) GongMarshallField(stage *Stage, fieldName string) (res string) {
	ident := perpendicularvectorhalfway.GongGetIdentifier(stage)
	_ = ident
	switch fieldName {
	case "Name":
		res = __gong__marshallString(ident, "Name", perpendicularvectorhalfway.Name)
	case "StartX":
		res = __gong__marshallFloat(ident, "StartX", perpendicularvectorhalfway.StartX)
	case "StartY":
		res = __gong__marshallFloat(ident, "StartY", perpendicularvectorhalfway.StartY)
	case "EndX":
		res = __gong__marshallFloat(ident, "EndX", perpendicularvectorhalfway.EndX)
	case "EndY":
		res = __gong__marshallFloat(ident, "EndY", perpendicularvectorhalfway.EndY)

	default:
		log.Panicf("Unknown field %s for Gongstruct PerpendicularVectorHalfway", fieldName)
	}
	return
}

func (plant2ddiagram *Plant2DDiagram) GongMarshallField(stage *Stage, fieldName string) (res string) {
	ident := plant2ddiagram.GongGetIdentifier(stage)
	_ = ident
	switch fieldName {
	case "Name":
		res = __gong__marshallString(ident, "Name", plant2ddiagram.Name)
	case "OriginX":
		res = __gong__marshallFloat(ident, "OriginX", plant2ddiagram.OriginX)
	case "OriginY":
		res = __gong__marshallFloat(ident, "OriginY", plant2ddiagram.OriginY)
	case "Zoom":
		res = __gong__marshallFloat(ident, "Zoom", plant2ddiagram.Zoom)
	case "IsRhombusNodesExpanded":
		res = __gong__marshallBool(ident, "IsRhombusNodesExpanded", plant2ddiagram.IsRhombusNodesExpanded)
	case "IsArcNodesExpanded":
		res = __gong__marshallBool(ident, "IsArcNodesExpanded", plant2ddiagram.IsArcNodesExpanded)
	case "IsHiddenAxesShape":
		res = __gong__marshallBool(ident, "IsHiddenAxesShape", plant2ddiagram.IsHiddenAxesShape)
	case "IsHiddenReferenceRhombus":
		res = __gong__marshallBool(ident, "IsHiddenReferenceRhombus", plant2ddiagram.IsHiddenReferenceRhombus)
	case "IsHiddenPlantCircumferenceShape":
		res = __gong__marshallBool(ident, "IsHiddenPlantCircumferenceShape", plant2ddiagram.IsHiddenPlantCircumferenceShape)
	case "IsHiddenGridPathShape":
		res = __gong__marshallBool(ident, "IsHiddenGridPathShape", plant2ddiagram.IsHiddenGridPathShape)
	case "IsHiddenRhombusGridShape":
		res = __gong__marshallBool(ident, "IsHiddenRhombusGridShape", plant2ddiagram.IsHiddenRhombusGridShape)
	case "IsHiddenExplanationTextShape":
		res = __gong__marshallBool(ident, "IsHiddenExplanationTextShape", plant2ddiagram.IsHiddenExplanationTextShape)
	case "IsHiddenRotatedReferenceRhombus":
		res = __gong__marshallBool(ident, "IsHiddenRotatedReferenceRhombus", plant2ddiagram.IsHiddenRotatedReferenceRhombus)
	case "IsHiddenRotatedPlantCircumferenceShape":
		res = __gong__marshallBool(ident, "IsHiddenRotatedPlantCircumferenceShape", plant2ddiagram.IsHiddenRotatedPlantCircumferenceShape)
	case "IsHiddenRotatedGridPathShape":
		res = __gong__marshallBool(ident, "IsHiddenRotatedGridPathShape", plant2ddiagram.IsHiddenRotatedGridPathShape)
	case "IsHiddenRotatedRhombusGridShape":
		res = __gong__marshallBool(ident, "IsHiddenRotatedRhombusGridShape", plant2ddiagram.IsHiddenRotatedRhombusGridShape)
	case "IsHiddenGrowthPathRhombusGridShape":
		res = __gong__marshallBool(ident, "IsHiddenGrowthPathRhombusGridShape", plant2ddiagram.IsHiddenGrowthPathRhombusGridShape)
	case "IsHiddenGrowthVectorShape":
		res = __gong__marshallBool(ident, "IsHiddenGrowthVectorShape", plant2ddiagram.IsHiddenGrowthVectorShape)
	case "IsHiddenPerpendicularVectorGrid":
		res = __gong__marshallBool(ident, "IsHiddenPerpendicularVectorGrid", plant2ddiagram.IsHiddenPerpendicularVectorGrid)
	case "IsHiddenBaseVectorShapeGrid":
		res = __gong__marshallBool(ident, "IsHiddenBaseVectorShapeGrid", plant2ddiagram.IsHiddenBaseVectorShapeGrid)
	case "IsHiddenArcNormalVectorShapeGrid":
		res = __gong__marshallBool(ident, "IsHiddenArcNormalVectorShapeGrid", plant2ddiagram.IsHiddenArcNormalVectorShapeGrid)
	case "IsHiddenStartArcShapeGrid":
		res = __gong__marshallBool(ident, "IsHiddenStartArcShapeGrid", plant2ddiagram.IsHiddenStartArcShapeGrid)
	case "IsHiddenMidArcVectorShapeGrid":
		res = __gong__marshallBool(ident, "IsHiddenMidArcVectorShapeGrid", plant2ddiagram.IsHiddenMidArcVectorShapeGrid)
	case "IsHiddenEndArcShapeGrid":
		res = __gong__marshallBool(ident, "IsHiddenEndArcShapeGrid", plant2ddiagram.IsHiddenEndArcShapeGrid)
	case "IsHiddenGrowthCurve2D":
		res = __gong__marshallBool(ident, "IsHiddenGrowthCurve2D", plant2ddiagram.IsHiddenGrowthCurve2D)
	case "IsHiddenStackOfGrowthCurve2DByGrowthVector":
		res = __gong__marshallBool(ident, "IsHiddenStackOfGrowthCurve2DByGrowthVector", plant2ddiagram.IsHiddenStackOfGrowthCurve2DByGrowthVector)
	case "IsChecked":
		res = __gong__marshallBool(ident, "IsChecked", plant2ddiagram.IsChecked)
	case "ComputedPrefix":
		res = __gong__marshallString(ident, "ComputedPrefix", plant2ddiagram.ComputedPrefix)
	case "IsExpanded":
		res = __gong__marshallBool(ident, "IsExpanded", plant2ddiagram.IsExpanded)

	default:
		log.Panicf("Unknown field %s for Gongstruct Plant2DDiagram", fieldName)
	}
	return
}

func (plant3ddiagram *Plant3DDiagram) GongMarshallField(stage *Stage, fieldName string) (res string) {
	ident := plant3ddiagram.GongGetIdentifier(stage)
	_ = ident
	switch fieldName {
	case "Name":
		res = __gong__marshallString(ident, "Name", plant3ddiagram.Name)
	case "IsHiddenStemCylinder3DShape":
		res = __gong__marshallBool(ident, "IsHiddenStemCylinder3DShape", plant3ddiagram.IsHiddenStemCylinder3DShape)
	case "IsHiddenParastichyNCurves3DShape":
		res = __gong__marshallBool(ident, "IsHiddenParastichyNCurves3DShape", plant3ddiagram.IsHiddenParastichyNCurves3DShape)
	case "IsHiddenParastichyMCurves3DShape":
		res = __gong__marshallBool(ident, "IsHiddenParastichyMCurves3DShape", plant3ddiagram.IsHiddenParastichyMCurves3DShape)
	case "IsHiddenCutLine3DShape":
		res = __gong__marshallBool(ident, "IsHiddenCutLine3DShape", plant3ddiagram.IsHiddenCutLine3DShape)
	case "IsHiddenCircumference3DShape":
		res = __gong__marshallBool(ident, "IsHiddenCircumference3DShape", plant3ddiagram.IsHiddenCircumference3DShape)
	case "IsHiddenTiledFloor3DShape":
		res = __gong__marshallBool(ident, "IsHiddenTiledFloor3DShape", plant3ddiagram.IsHiddenTiledFloor3DShape)
	case "IsHiddenLeaves3DShape":
		res = __gong__marshallBool(ident, "IsHiddenLeaves3DShape", plant3ddiagram.IsHiddenLeaves3DShape)
	case "IsChecked":
		res = __gong__marshallBool(ident, "IsChecked", plant3ddiagram.IsChecked)
	case "ComputedPrefix":
		res = __gong__marshallString(ident, "ComputedPrefix", plant3ddiagram.ComputedPrefix)
	case "IsExpanded":
		res = __gong__marshallBool(ident, "IsExpanded", plant3ddiagram.IsExpanded)

	case "StemCylinder3DShape":
		if plant3ddiagram.StemCylinder3DShape != nil {
			res = __gong__marshallPointer(ident, "StemCylinder3DShape", plant3ddiagram.StemCylinder3DShape.GongGetIdentifier(stage))
		} else {
			res = __gong__marshallPointer(ident, "StemCylinder3DShape", "nil")
		}
	case "ParastichyNCurves3DShape":
		if plant3ddiagram.ParastichyNCurves3DShape != nil {
			res = __gong__marshallPointer(ident, "ParastichyNCurves3DShape", plant3ddiagram.ParastichyNCurves3DShape.GongGetIdentifier(stage))
		} else {
			res = __gong__marshallPointer(ident, "ParastichyNCurves3DShape", "nil")
		}
	case "ParastichyMCurves3DShape":
		if plant3ddiagram.ParastichyMCurves3DShape != nil {
			res = __gong__marshallPointer(ident, "ParastichyMCurves3DShape", plant3ddiagram.ParastichyMCurves3DShape.GongGetIdentifier(stage))
		} else {
			res = __gong__marshallPointer(ident, "ParastichyMCurves3DShape", "nil")
		}
	case "CutLine3DShape":
		if plant3ddiagram.CutLine3DShape != nil {
			res = __gong__marshallPointer(ident, "CutLine3DShape", plant3ddiagram.CutLine3DShape.GongGetIdentifier(stage))
		} else {
			res = __gong__marshallPointer(ident, "CutLine3DShape", "nil")
		}
	case "Circumference3DShape":
		if plant3ddiagram.Circumference3DShape != nil {
			res = __gong__marshallPointer(ident, "Circumference3DShape", plant3ddiagram.Circumference3DShape.GongGetIdentifier(stage))
		} else {
			res = __gong__marshallPointer(ident, "Circumference3DShape", "nil")
		}
	case "Leaves3DShape":
		if plant3ddiagram.Leaves3DShape != nil {
			res = __gong__marshallPointer(ident, "Leaves3DShape", plant3ddiagram.Leaves3DShape.GongGetIdentifier(stage))
		} else {
			res = __gong__marshallPointer(ident, "Leaves3DShape", "nil")
		}
	case "Rendered3DShape":
		if plant3ddiagram.Rendered3DShape != nil {
			res = __gong__marshallPointer(ident, "Rendered3DShape", plant3ddiagram.Rendered3DShape.GongGetIdentifier(stage))
		} else {
			res = __gong__marshallPointer(ident, "Rendered3DShape", "nil")
		}
	default:
		log.Panicf("Unknown field %s for Gongstruct Plant3DDiagram", fieldName)
	}
	return
}

func (plantabstract *PlantAbstract) GongMarshallField(stage *Stage, fieldName string) (res string) {
	ident := plantabstract.GongGetIdentifier(stage)
	_ = ident
	switch fieldName {
	case "Name":
		res = __gong__marshallString(ident, "Name", plantabstract.Name)
	case "N":
		res = __gong__marshallInt(ident, "N", plantabstract.N)
	case "M":
		res = __gong__marshallInt(ident, "M", plantabstract.M)
	case "StackHeight":
		res = __gong__marshallInt(ident, "StackHeight", plantabstract.StackHeight)
	case "RhombusInsideAngle":
		res = __gong__marshallFloat(ident, "RhombusInsideAngle", plantabstract.RhombusInsideAngle)
	case "RhombusSideLength":
		res = __gong__marshallFloat(ident, "RhombusSideLength", plantabstract.RhombusSideLength)
	case "PlantType":
		res = __gong__marshallEnumString(ident, "PlantType", plantabstract.PlantType.ToCodeString())
	case "CurrentView":
		res = __gong__marshallEnumString(ident, "CurrentView", plantabstract.CurrentView.ToCodeString())
	case "ComputedPrefix":
		res = __gong__marshallString(ident, "ComputedPrefix", plantabstract.ComputedPrefix)
	case "IsExpanded":
		res = __gong__marshallBool(ident, "IsExpanded", plantabstract.IsExpanded)
	case "IsSelected":
		res = __gong__marshallBool(ident, "IsSelected", plantabstract.IsSelected)
	case "IsPlant2DDiagramsNodeExpanded":
		res = __gong__marshallBool(ident, "IsPlant2DDiagramsNodeExpanded", plantabstract.IsPlant2DDiagramsNodeExpanded)
	case "IsPlant3DDiagramsNodeExpanded":
		res = __gong__marshallBool(ident, "IsPlant3DDiagramsNodeExpanded", plantabstract.IsPlant3DDiagramsNodeExpanded)
	case "IsVase2DDiagramsNodeExpanded":
		res = __gong__marshallBool(ident, "IsVase2DDiagramsNodeExpanded", plantabstract.IsVase2DDiagramsNodeExpanded)
	case "IsTubeVase3DDiagramsNodeExpanded":
		res = __gong__marshallBool(ident, "IsTubeVase3DDiagramsNodeExpanded", plantabstract.IsTubeVase3DDiagramsNodeExpanded)
	case "IsStool2DDiagramsNodeExpanded":
		res = __gong__marshallBool(ident, "IsStool2DDiagramsNodeExpanded", plantabstract.IsStool2DDiagramsNodeExpanded)
	case "IsStool3DDiagramsNodeExpanded":
		res = __gong__marshallBool(ident, "IsStool3DDiagramsNodeExpanded", plantabstract.IsStool3DDiagramsNodeExpanded)
	case "IsClock2DDiagramsNodeExpanded":
		res = __gong__marshallBool(ident, "IsClock2DDiagramsNodeExpanded", plantabstract.IsClock2DDiagramsNodeExpanded)
	case "IsClock3DDiagramsNodeExpanded":
		res = __gong__marshallBool(ident, "IsClock3DDiagramsNodeExpanded", plantabstract.IsClock3DDiagramsNodeExpanded)

	case "TubeVaseAbstract":
		if plantabstract.TubeVaseAbstract != nil {
			res = __gong__marshallPointer(ident, "TubeVaseAbstract", plantabstract.TubeVaseAbstract.GongGetIdentifier(stage))
		} else {
			res = __gong__marshallPointer(ident, "TubeVaseAbstract", "nil")
		}
	case "Plant2DDiagrams":
		var sb strings.Builder
		for _, _plant2ddiagram := range plantabstract.Plant2DDiagrams {
			sb.WriteString(__gong__marshallSliceOfPointers(ident, "Plant2DDiagrams", _plant2ddiagram.GongGetIdentifier(stage)))
		}
		res = sb.String()
	case "Plant3DDiagrams":
		var sb strings.Builder
		for _, _plant3ddiagram := range plantabstract.Plant3DDiagrams {
			sb.WriteString(__gong__marshallSliceOfPointers(ident, "Plant3DDiagrams", _plant3ddiagram.GongGetIdentifier(stage)))
		}
		res = sb.String()
	case "Vase2DDiagrams":
		var sb strings.Builder
		for _, _vase2ddiagram := range plantabstract.Vase2DDiagrams {
			sb.WriteString(__gong__marshallSliceOfPointers(ident, "Vase2DDiagrams", _vase2ddiagram.GongGetIdentifier(stage)))
		}
		res = sb.String()
	case "TubeVase3DDiagrams":
		var sb strings.Builder
		for _, _tubevase3ddiagram := range plantabstract.TubeVase3DDiagrams {
			sb.WriteString(__gong__marshallSliceOfPointers(ident, "TubeVase3DDiagrams", _tubevase3ddiagram.GongGetIdentifier(stage)))
		}
		res = sb.String()
	case "Stool2DDiagrams":
		var sb strings.Builder
		for _, _stool2ddiagram := range plantabstract.Stool2DDiagrams {
			sb.WriteString(__gong__marshallSliceOfPointers(ident, "Stool2DDiagrams", _stool2ddiagram.GongGetIdentifier(stage)))
		}
		res = sb.String()
	case "Stool3DDiagrams":
		var sb strings.Builder
		for _, _stool3ddiagram := range plantabstract.Stool3DDiagrams {
			sb.WriteString(__gong__marshallSliceOfPointers(ident, "Stool3DDiagrams", _stool3ddiagram.GongGetIdentifier(stage)))
		}
		res = sb.String()
	case "Clock2DDiagrams":
		var sb strings.Builder
		for _, _clock2ddiagram := range plantabstract.Clock2DDiagrams {
			sb.WriteString(__gong__marshallSliceOfPointers(ident, "Clock2DDiagrams", _clock2ddiagram.GongGetIdentifier(stage)))
		}
		res = sb.String()
	case "Clock3DDiagrams":
		var sb strings.Builder
		for _, _clock3ddiagram := range plantabstract.Clock3DDiagrams {
			sb.WriteString(__gong__marshallSliceOfPointers(ident, "Clock3DDiagrams", _clock3ddiagram.GongGetIdentifier(stage)))
		}
		res = sb.String()
	default:
		log.Panicf("Unknown field %s for Gongstruct PlantAbstract", fieldName)
	}
	return
}

func (plantcircumferenceshape *PlantCircumferenceShape) GongMarshallField(stage *Stage, fieldName string) (res string) {
	ident := plantcircumferenceshape.GongGetIdentifier(stage)
	_ = ident
	switch fieldName {
	case "Name":
		res = __gong__marshallString(ident, "Name", plantcircumferenceshape.Name)
	case "AngleDegree":
		res = __gong__marshallFloat(ident, "AngleDegree", plantcircumferenceshape.AngleDegree)
	case "Length":
		res = __gong__marshallFloat(ident, "Length", plantcircumferenceshape.Length)

	default:
		log.Panicf("Unknown field %s for Gongstruct PlantCircumferenceShape", fieldName)
	}
	return
}

func (pointsandlines3dshape *PointsAndLines3DShape) GongMarshallField(stage *Stage, fieldName string) (res string) {
	ident := pointsandlines3dshape.GongGetIdentifier(stage)
	_ = ident
	switch fieldName {
	case "Name":
		res = __gong__marshallString(ident, "Name", pointsandlines3dshape.Name)

	default:
		log.Panicf("Unknown field %s for Gongstruct PointsAndLines3DShape", fieldName)
	}
	return
}

func (pxshape *PxShape) GongMarshallField(stage *Stage, fieldName string) (res string) {
	ident := pxshape.GongGetIdentifier(stage)
	_ = ident
	switch fieldName {
	case "Name":
		res = __gong__marshallString(ident, "Name", pxshape.Name)
	case "X":
		res = __gong__marshallFloat(ident, "X", pxshape.X)
	case "Y":
		res = __gong__marshallFloat(ident, "Y", pxshape.Y)

	default:
		log.Panicf("Unknown field %s for Gongstruct PxShape", fieldName)
	}
	return
}

func (rendered3dshape *Rendered3DShape) GongMarshallField(stage *Stage, fieldName string) (res string) {
	ident := rendered3dshape.GongGetIdentifier(stage)
	_ = ident
	switch fieldName {
	case "Name":
		res = __gong__marshallString(ident, "Name", rendered3dshape.Name)
	case "ViewX":
		res = __gong__marshallFloat(ident, "ViewX", rendered3dshape.ViewX)
	case "ViewY":
		res = __gong__marshallFloat(ident, "ViewY", rendered3dshape.ViewY)
	case "ViewZ":
		res = __gong__marshallFloat(ident, "ViewZ", rendered3dshape.ViewZ)
	case "TargetX":
		res = __gong__marshallFloat(ident, "TargetX", rendered3dshape.TargetX)
	case "TargetY":
		res = __gong__marshallFloat(ident, "TargetY", rendered3dshape.TargetY)
	case "TargetZ":
		res = __gong__marshallFloat(ident, "TargetZ", rendered3dshape.TargetZ)
	case "Fov":
		res = __gong__marshallFloat(ident, "Fov", rendered3dshape.Fov)

	default:
		log.Panicf("Unknown field %s for Gongstruct Rendered3DShape", fieldName)
	}
	return
}

func (rhombusshape *RhombusShape) GongMarshallField(stage *Stage, fieldName string) (res string) {
	ident := rhombusshape.GongGetIdentifier(stage)
	_ = ident
	switch fieldName {
	case "Name":
		res = __gong__marshallString(ident, "Name", rhombusshape.Name)
	case "X":
		res = __gong__marshallFloat(ident, "X", rhombusshape.X)
	case "Y":
		res = __gong__marshallFloat(ident, "Y", rhombusshape.Y)

	default:
		log.Panicf("Unknown field %s for Gongstruct RhombusShape", fieldName)
	}
	return
}

func (rhombusstuff *RhombusStuff) GongMarshallField(stage *Stage, fieldName string) (res string) {
	ident := rhombusstuff.GongGetIdentifier(stage)
	_ = ident
	switch fieldName {
	case "Name":
		res = __gong__marshallString(ident, "Name", rhombusstuff.Name)

	default:
		log.Panicf("Unknown field %s for Gongstruct RhombusStuff", fieldName)
	}
	return
}

func (rotatedrhombusgridshape *RotatedRhombusGridShape) GongMarshallField(stage *Stage, fieldName string) (res string) {
	ident := rotatedrhombusgridshape.GongGetIdentifier(stage)
	_ = ident
	switch fieldName {
	case "Name":
		res = __gong__marshallString(ident, "Name", rotatedrhombusgridshape.Name)

	default:
		log.Panicf("Unknown field %s for Gongstruct RotatedRhombusGridShape", fieldName)
	}
	return
}

func (rotatedrhombusshape *RotatedRhombusShape) GongMarshallField(stage *Stage, fieldName string) (res string) {
	ident := rotatedrhombusshape.GongGetIdentifier(stage)
	_ = ident
	switch fieldName {
	case "Name":
		res = __gong__marshallString(ident, "Name", rotatedrhombusshape.Name)
	case "X":
		res = __gong__marshallFloat(ident, "X", rotatedrhombusshape.X)
	case "Y":
		res = __gong__marshallFloat(ident, "Y", rotatedrhombusshape.Y)

	default:
		log.Panicf("Unknown field %s for Gongstruct RotatedRhombusShape", fieldName)
	}
	return
}

func (rotatedsampledpoints3dshape *RotatedSampledPoints3DShape) GongMarshallField(stage *Stage, fieldName string) (res string) {
	ident := rotatedsampledpoints3dshape.GongGetIdentifier(stage)
	_ = ident
	switch fieldName {
	case "Name":
		res = __gong__marshallString(ident, "Name", rotatedsampledpoints3dshape.Name)

	default:
		log.Panicf("Unknown field %s for Gongstruct RotatedSampledPoints3DShape", fieldName)
	}
	return
}

func (rotatedseatandlegs3dshape *RotatedSeatAndLegs3DShape) GongMarshallField(stage *Stage, fieldName string) (res string) {
	ident := rotatedseatandlegs3dshape.GongGetIdentifier(stage)
	_ = ident
	switch fieldName {
	case "Name":
		res = __gong__marshallString(ident, "Name", rotatedseatandlegs3dshape.Name)

	default:
		log.Panicf("Unknown field %s for Gongstruct RotatedSeatAndLegs3DShape", fieldName)
	}
	return
}

func (sampledpoints3dshape *SampledPoints3DShape) GongMarshallField(stage *Stage, fieldName string) (res string) {
	ident := sampledpoints3dshape.GongGetIdentifier(stage)
	_ = ident
	switch fieldName {
	case "Name":
		res = __gong__marshallString(ident, "Name", sampledpoints3dshape.Name)

	default:
		log.Panicf("Unknown field %s for Gongstruct SampledPoints3DShape", fieldName)
	}
	return
}

func (seat3dshape *Seat3DShape) GongMarshallField(stage *Stage, fieldName string) (res string) {
	ident := seat3dshape.GongGetIdentifier(stage)
	_ = ident
	switch fieldName {
	case "Name":
		res = __gong__marshallString(ident, "Name", seat3dshape.Name)

	default:
		log.Panicf("Unknown field %s for Gongstruct Seat3DShape", fieldName)
	}
	return
}

func (seatandlegs3dshape *SeatAndLegs3DShape) GongMarshallField(stage *Stage, fieldName string) (res string) {
	ident := seatandlegs3dshape.GongGetIdentifier(stage)
	_ = ident
	switch fieldName {
	case "Name":
		res = __gong__marshallString(ident, "Name", seatandlegs3dshape.Name)

	default:
		log.Panicf("Unknown field %s for Gongstruct SeatAndLegs3DShape", fieldName)
	}
	return
}

func (seatbottomcurveshape *SeatBottomCurveShape) GongMarshallField(stage *Stage, fieldName string) (res string) {
	ident := seatbottomcurveshape.GongGetIdentifier(stage)
	_ = ident
	switch fieldName {
	case "Name":
		res = __gong__marshallString(ident, "Name", seatbottomcurveshape.Name)

	default:
		log.Panicf("Unknown field %s for Gongstruct SeatBottomCurveShape", fieldName)
	}
	return
}

func (seattopcurveshape *SeatTopCurveShape) GongMarshallField(stage *Stage, fieldName string) (res string) {
	ident := seattopcurveshape.GongGetIdentifier(stage)
	_ = ident
	switch fieldName {
	case "Name":
		res = __gong__marshallString(ident, "Name", seattopcurveshape.Name)

	default:
		log.Panicf("Unknown field %s for Gongstruct SeatTopCurveShape", fieldName)
	}
	return
}

func (shiftedbottomtopstartarcshape *ShiftedBottomTopStartArcShape) GongMarshallField(stage *Stage, fieldName string) (res string) {
	ident := shiftedbottomtopstartarcshape.GongGetIdentifier(stage)
	_ = ident
	switch fieldName {
	case "Name":
		res = __gong__marshallString(ident, "Name", shiftedbottomtopstartarcshape.Name)
	case "StartX":
		res = __gong__marshallFloat(ident, "StartX", shiftedbottomtopstartarcshape.StartX)
	case "StartY":
		res = __gong__marshallFloat(ident, "StartY", shiftedbottomtopstartarcshape.StartY)
	case "EndX":
		res = __gong__marshallFloat(ident, "EndX", shiftedbottomtopstartarcshape.EndX)
	case "EndY":
		res = __gong__marshallFloat(ident, "EndY", shiftedbottomtopstartarcshape.EndY)
	case "XAxisRotation":
		res = __gong__marshallFloat(ident, "XAxisRotation", shiftedbottomtopstartarcshape.XAxisRotation)
	case "LargeArcFlag":
		res = __gong__marshallBool(ident, "LargeArcFlag", shiftedbottomtopstartarcshape.LargeArcFlag)
	case "SweepFlag":
		res = __gong__marshallBool(ident, "SweepFlag", shiftedbottomtopstartarcshape.SweepFlag)
	case "RadiusX":
		res = __gong__marshallFloat(ident, "RadiusX", shiftedbottomtopstartarcshape.RadiusX)
	case "RadiusY":
		res = __gong__marshallFloat(ident, "RadiusY", shiftedbottomtopstartarcshape.RadiusY)

	default:
		log.Panicf("Unknown field %s for Gongstruct ShiftedBottomTopStartArcShape", fieldName)
	}
	return
}

func (shiftedbottomtopstartarcshapegrid *ShiftedBottomTopStartArcShapeGrid) GongMarshallField(stage *Stage, fieldName string) (res string) {
	ident := shiftedbottomtopstartarcshapegrid.GongGetIdentifier(stage)
	_ = ident
	switch fieldName {
	case "Name":
		res = __gong__marshallString(ident, "Name", shiftedbottomtopstartarcshapegrid.Name)

	default:
		log.Panicf("Unknown field %s for Gongstruct ShiftedBottomTopStartArcShapeGrid", fieldName)
	}
	return
}

func (shiftedleftgrowthcurve2dribbon *ShiftedLeftGrowthCurve2DRibbon) GongMarshallField(stage *Stage, fieldName string) (res string) {
	ident := shiftedleftgrowthcurve2dribbon.GongGetIdentifier(stage)
	_ = ident
	switch fieldName {
	case "Name":
		res = __gong__marshallString(ident, "Name", shiftedleftgrowthcurve2dribbon.Name)

	default:
		log.Panicf("Unknown field %s for Gongstruct ShiftedLeftGrowthCurve2DRibbon", fieldName)
	}
	return
}

func (shiftedleftgrowthcurve2dribbonendshape *ShiftedLeftGrowthCurve2DRibbonEndShape) GongMarshallField(stage *Stage, fieldName string) (res string) {
	ident := shiftedleftgrowthcurve2dribbonendshape.GongGetIdentifier(stage)
	_ = ident
	switch fieldName {
	case "Name":
		res = __gong__marshallString(ident, "Name", shiftedleftgrowthcurve2dribbonendshape.Name)
	case "BottomStartX":
		res = __gong__marshallFloat(ident, "BottomStartX", shiftedleftgrowthcurve2dribbonendshape.BottomStartX)
	case "BottomStartY":
		res = __gong__marshallFloat(ident, "BottomStartY", shiftedleftgrowthcurve2dribbonendshape.BottomStartY)
	case "BottomEndX":
		res = __gong__marshallFloat(ident, "BottomEndX", shiftedleftgrowthcurve2dribbonendshape.BottomEndX)
	case "BottomEndY":
		res = __gong__marshallFloat(ident, "BottomEndY", shiftedleftgrowthcurve2dribbonendshape.BottomEndY)
	case "BottomRadiusX":
		res = __gong__marshallFloat(ident, "BottomRadiusX", shiftedleftgrowthcurve2dribbonendshape.BottomRadiusX)
	case "BottomRadiusY":
		res = __gong__marshallFloat(ident, "BottomRadiusY", shiftedleftgrowthcurve2dribbonendshape.BottomRadiusY)
	case "BottomXAxisRotation":
		res = __gong__marshallFloat(ident, "BottomXAxisRotation", shiftedleftgrowthcurve2dribbonendshape.BottomXAxisRotation)
	case "BottomLargeArcFlag":
		res = __gong__marshallBool(ident, "BottomLargeArcFlag", shiftedleftgrowthcurve2dribbonendshape.BottomLargeArcFlag)
	case "BottomSweepFlag":
		res = __gong__marshallBool(ident, "BottomSweepFlag", shiftedleftgrowthcurve2dribbonendshape.BottomSweepFlag)
	case "TopStartX":
		res = __gong__marshallFloat(ident, "TopStartX", shiftedleftgrowthcurve2dribbonendshape.TopStartX)
	case "TopStartY":
		res = __gong__marshallFloat(ident, "TopStartY", shiftedleftgrowthcurve2dribbonendshape.TopStartY)
	case "TopEndX":
		res = __gong__marshallFloat(ident, "TopEndX", shiftedleftgrowthcurve2dribbonendshape.TopEndX)
	case "TopEndY":
		res = __gong__marshallFloat(ident, "TopEndY", shiftedleftgrowthcurve2dribbonendshape.TopEndY)
	case "TopRadiusX":
		res = __gong__marshallFloat(ident, "TopRadiusX", shiftedleftgrowthcurve2dribbonendshape.TopRadiusX)
	case "TopRadiusY":
		res = __gong__marshallFloat(ident, "TopRadiusY", shiftedleftgrowthcurve2dribbonendshape.TopRadiusY)
	case "TopXAxisRotation":
		res = __gong__marshallFloat(ident, "TopXAxisRotation", shiftedleftgrowthcurve2dribbonendshape.TopXAxisRotation)
	case "TopLargeArcFlag":
		res = __gong__marshallBool(ident, "TopLargeArcFlag", shiftedleftgrowthcurve2dribbonendshape.TopLargeArcFlag)
	case "TopSweepFlag":
		res = __gong__marshallBool(ident, "TopSweepFlag", shiftedleftgrowthcurve2dribbonendshape.TopSweepFlag)

	default:
		log.Panicf("Unknown field %s for Gongstruct ShiftedLeftGrowthCurve2DRibbonEndShape", fieldName)
	}
	return
}

func (shiftedleftgrowthcurve2dribbonstartshape *ShiftedLeftGrowthCurve2DRibbonStartShape) GongMarshallField(stage *Stage, fieldName string) (res string) {
	ident := shiftedleftgrowthcurve2dribbonstartshape.GongGetIdentifier(stage)
	_ = ident
	switch fieldName {
	case "Name":
		res = __gong__marshallString(ident, "Name", shiftedleftgrowthcurve2dribbonstartshape.Name)
	case "BottomStartX":
		res = __gong__marshallFloat(ident, "BottomStartX", shiftedleftgrowthcurve2dribbonstartshape.BottomStartX)
	case "BottomStartY":
		res = __gong__marshallFloat(ident, "BottomStartY", shiftedleftgrowthcurve2dribbonstartshape.BottomStartY)
	case "BottomEndX":
		res = __gong__marshallFloat(ident, "BottomEndX", shiftedleftgrowthcurve2dribbonstartshape.BottomEndX)
	case "BottomEndY":
		res = __gong__marshallFloat(ident, "BottomEndY", shiftedleftgrowthcurve2dribbonstartshape.BottomEndY)
	case "BottomRadiusX":
		res = __gong__marshallFloat(ident, "BottomRadiusX", shiftedleftgrowthcurve2dribbonstartshape.BottomRadiusX)
	case "BottomRadiusY":
		res = __gong__marshallFloat(ident, "BottomRadiusY", shiftedleftgrowthcurve2dribbonstartshape.BottomRadiusY)
	case "BottomXAxisRotation":
		res = __gong__marshallFloat(ident, "BottomXAxisRotation", shiftedleftgrowthcurve2dribbonstartshape.BottomXAxisRotation)
	case "BottomLargeArcFlag":
		res = __gong__marshallBool(ident, "BottomLargeArcFlag", shiftedleftgrowthcurve2dribbonstartshape.BottomLargeArcFlag)
	case "BottomSweepFlag":
		res = __gong__marshallBool(ident, "BottomSweepFlag", shiftedleftgrowthcurve2dribbonstartshape.BottomSweepFlag)
	case "TopStartX":
		res = __gong__marshallFloat(ident, "TopStartX", shiftedleftgrowthcurve2dribbonstartshape.TopStartX)
	case "TopStartY":
		res = __gong__marshallFloat(ident, "TopStartY", shiftedleftgrowthcurve2dribbonstartshape.TopStartY)
	case "TopEndX":
		res = __gong__marshallFloat(ident, "TopEndX", shiftedleftgrowthcurve2dribbonstartshape.TopEndX)
	case "TopEndY":
		res = __gong__marshallFloat(ident, "TopEndY", shiftedleftgrowthcurve2dribbonstartshape.TopEndY)
	case "TopRadiusX":
		res = __gong__marshallFloat(ident, "TopRadiusX", shiftedleftgrowthcurve2dribbonstartshape.TopRadiusX)
	case "TopRadiusY":
		res = __gong__marshallFloat(ident, "TopRadiusY", shiftedleftgrowthcurve2dribbonstartshape.TopRadiusY)
	case "TopXAxisRotation":
		res = __gong__marshallFloat(ident, "TopXAxisRotation", shiftedleftgrowthcurve2dribbonstartshape.TopXAxisRotation)
	case "TopLargeArcFlag":
		res = __gong__marshallBool(ident, "TopLargeArcFlag", shiftedleftgrowthcurve2dribbonstartshape.TopLargeArcFlag)
	case "TopSweepFlag":
		res = __gong__marshallBool(ident, "TopSweepFlag", shiftedleftgrowthcurve2dribbonstartshape.TopSweepFlag)

	default:
		log.Panicf("Unknown field %s for Gongstruct ShiftedLeftGrowthCurve2DRibbonStartShape", fieldName)
	}
	return
}

func (shiftedleftpartiallygrowthcurve2dribbon *ShiftedLeftPartiallyGrowthCurve2DRibbon) GongMarshallField(stage *Stage, fieldName string) (res string) {
	ident := shiftedleftpartiallygrowthcurve2dribbon.GongGetIdentifier(stage)
	_ = ident
	switch fieldName {
	case "Name":
		res = __gong__marshallString(ident, "Name", shiftedleftpartiallygrowthcurve2dribbon.Name)

	default:
		log.Panicf("Unknown field %s for Gongstruct ShiftedLeftPartiallyGrowthCurve2DRibbon", fieldName)
	}
	return
}

func (shiftedleftpartiallygrowthcurve2dribbonendshape *ShiftedLeftPartiallyGrowthCurve2DRibbonEndShape) GongMarshallField(stage *Stage, fieldName string) (res string) {
	ident := shiftedleftpartiallygrowthcurve2dribbonendshape.GongGetIdentifier(stage)
	_ = ident
	switch fieldName {
	case "Name":
		res = __gong__marshallString(ident, "Name", shiftedleftpartiallygrowthcurve2dribbonendshape.Name)
	case "BottomStartX":
		res = __gong__marshallFloat(ident, "BottomStartX", shiftedleftpartiallygrowthcurve2dribbonendshape.BottomStartX)
	case "BottomStartY":
		res = __gong__marshallFloat(ident, "BottomStartY", shiftedleftpartiallygrowthcurve2dribbonendshape.BottomStartY)
	case "BottomEndX":
		res = __gong__marshallFloat(ident, "BottomEndX", shiftedleftpartiallygrowthcurve2dribbonendshape.BottomEndX)
	case "BottomEndY":
		res = __gong__marshallFloat(ident, "BottomEndY", shiftedleftpartiallygrowthcurve2dribbonendshape.BottomEndY)
	case "BottomRadiusX":
		res = __gong__marshallFloat(ident, "BottomRadiusX", shiftedleftpartiallygrowthcurve2dribbonendshape.BottomRadiusX)
	case "BottomRadiusY":
		res = __gong__marshallFloat(ident, "BottomRadiusY", shiftedleftpartiallygrowthcurve2dribbonendshape.BottomRadiusY)
	case "BottomXAxisRotation":
		res = __gong__marshallFloat(ident, "BottomXAxisRotation", shiftedleftpartiallygrowthcurve2dribbonendshape.BottomXAxisRotation)
	case "BottomLargeArcFlag":
		res = __gong__marshallBool(ident, "BottomLargeArcFlag", shiftedleftpartiallygrowthcurve2dribbonendshape.BottomLargeArcFlag)
	case "BottomSweepFlag":
		res = __gong__marshallBool(ident, "BottomSweepFlag", shiftedleftpartiallygrowthcurve2dribbonendshape.BottomSweepFlag)
	case "TopStartX":
		res = __gong__marshallFloat(ident, "TopStartX", shiftedleftpartiallygrowthcurve2dribbonendshape.TopStartX)
	case "TopStartY":
		res = __gong__marshallFloat(ident, "TopStartY", shiftedleftpartiallygrowthcurve2dribbonendshape.TopStartY)
	case "TopEndX":
		res = __gong__marshallFloat(ident, "TopEndX", shiftedleftpartiallygrowthcurve2dribbonendshape.TopEndX)
	case "TopEndY":
		res = __gong__marshallFloat(ident, "TopEndY", shiftedleftpartiallygrowthcurve2dribbonendshape.TopEndY)
	case "TopRadiusX":
		res = __gong__marshallFloat(ident, "TopRadiusX", shiftedleftpartiallygrowthcurve2dribbonendshape.TopRadiusX)
	case "TopRadiusY":
		res = __gong__marshallFloat(ident, "TopRadiusY", shiftedleftpartiallygrowthcurve2dribbonendshape.TopRadiusY)
	case "TopXAxisRotation":
		res = __gong__marshallFloat(ident, "TopXAxisRotation", shiftedleftpartiallygrowthcurve2dribbonendshape.TopXAxisRotation)
	case "TopLargeArcFlag":
		res = __gong__marshallBool(ident, "TopLargeArcFlag", shiftedleftpartiallygrowthcurve2dribbonendshape.TopLargeArcFlag)
	case "TopSweepFlag":
		res = __gong__marshallBool(ident, "TopSweepFlag", shiftedleftpartiallygrowthcurve2dribbonendshape.TopSweepFlag)

	default:
		log.Panicf("Unknown field %s for Gongstruct ShiftedLeftPartiallyGrowthCurve2DRibbonEndShape", fieldName)
	}
	return
}

func (shiftedleftpartiallygrowthcurve2dribbonstartshape *ShiftedLeftPartiallyGrowthCurve2DRibbonStartShape) GongMarshallField(stage *Stage, fieldName string) (res string) {
	ident := shiftedleftpartiallygrowthcurve2dribbonstartshape.GongGetIdentifier(stage)
	_ = ident
	switch fieldName {
	case "Name":
		res = __gong__marshallString(ident, "Name", shiftedleftpartiallygrowthcurve2dribbonstartshape.Name)
	case "BottomStartX":
		res = __gong__marshallFloat(ident, "BottomStartX", shiftedleftpartiallygrowthcurve2dribbonstartshape.BottomStartX)
	case "BottomStartY":
		res = __gong__marshallFloat(ident, "BottomStartY", shiftedleftpartiallygrowthcurve2dribbonstartshape.BottomStartY)
	case "BottomEndX":
		res = __gong__marshallFloat(ident, "BottomEndX", shiftedleftpartiallygrowthcurve2dribbonstartshape.BottomEndX)
	case "BottomEndY":
		res = __gong__marshallFloat(ident, "BottomEndY", shiftedleftpartiallygrowthcurve2dribbonstartshape.BottomEndY)
	case "BottomRadiusX":
		res = __gong__marshallFloat(ident, "BottomRadiusX", shiftedleftpartiallygrowthcurve2dribbonstartshape.BottomRadiusX)
	case "BottomRadiusY":
		res = __gong__marshallFloat(ident, "BottomRadiusY", shiftedleftpartiallygrowthcurve2dribbonstartshape.BottomRadiusY)
	case "BottomXAxisRotation":
		res = __gong__marshallFloat(ident, "BottomXAxisRotation", shiftedleftpartiallygrowthcurve2dribbonstartshape.BottomXAxisRotation)
	case "BottomLargeArcFlag":
		res = __gong__marshallBool(ident, "BottomLargeArcFlag", shiftedleftpartiallygrowthcurve2dribbonstartshape.BottomLargeArcFlag)
	case "BottomSweepFlag":
		res = __gong__marshallBool(ident, "BottomSweepFlag", shiftedleftpartiallygrowthcurve2dribbonstartshape.BottomSweepFlag)
	case "TopStartX":
		res = __gong__marshallFloat(ident, "TopStartX", shiftedleftpartiallygrowthcurve2dribbonstartshape.TopStartX)
	case "TopStartY":
		res = __gong__marshallFloat(ident, "TopStartY", shiftedleftpartiallygrowthcurve2dribbonstartshape.TopStartY)
	case "TopEndX":
		res = __gong__marshallFloat(ident, "TopEndX", shiftedleftpartiallygrowthcurve2dribbonstartshape.TopEndX)
	case "TopEndY":
		res = __gong__marshallFloat(ident, "TopEndY", shiftedleftpartiallygrowthcurve2dribbonstartshape.TopEndY)
	case "TopRadiusX":
		res = __gong__marshallFloat(ident, "TopRadiusX", shiftedleftpartiallygrowthcurve2dribbonstartshape.TopRadiusX)
	case "TopRadiusY":
		res = __gong__marshallFloat(ident, "TopRadiusY", shiftedleftpartiallygrowthcurve2dribbonstartshape.TopRadiusY)
	case "TopXAxisRotation":
		res = __gong__marshallFloat(ident, "TopXAxisRotation", shiftedleftpartiallygrowthcurve2dribbonstartshape.TopXAxisRotation)
	case "TopLargeArcFlag":
		res = __gong__marshallBool(ident, "TopLargeArcFlag", shiftedleftpartiallygrowthcurve2dribbonstartshape.TopLargeArcFlag)
	case "TopSweepFlag":
		res = __gong__marshallBool(ident, "TopSweepFlag", shiftedleftpartiallygrowthcurve2dribbonstartshape.TopSweepFlag)

	default:
		log.Panicf("Unknown field %s for Gongstruct ShiftedLeftPartiallyGrowthCurve2DRibbonStartShape", fieldName)
	}
	return
}

func (shiftedleftstackgrowthcurveendarcshape *ShiftedLeftStackGrowthCurveEndArcShape) GongMarshallField(stage *Stage, fieldName string) (res string) {
	ident := shiftedleftstackgrowthcurveendarcshape.GongGetIdentifier(stage)
	_ = ident
	switch fieldName {
	case "Name":
		res = __gong__marshallString(ident, "Name", shiftedleftstackgrowthcurveendarcshape.Name)
	case "StartX":
		res = __gong__marshallFloat(ident, "StartX", shiftedleftstackgrowthcurveendarcshape.StartX)
	case "StartY":
		res = __gong__marshallFloat(ident, "StartY", shiftedleftstackgrowthcurveendarcshape.StartY)
	case "EndX":
		res = __gong__marshallFloat(ident, "EndX", shiftedleftstackgrowthcurveendarcshape.EndX)
	case "EndY":
		res = __gong__marshallFloat(ident, "EndY", shiftedleftstackgrowthcurveendarcshape.EndY)
	case "XAxisRotation":
		res = __gong__marshallFloat(ident, "XAxisRotation", shiftedleftstackgrowthcurveendarcshape.XAxisRotation)
	case "LargeArcFlag":
		res = __gong__marshallBool(ident, "LargeArcFlag", shiftedleftstackgrowthcurveendarcshape.LargeArcFlag)
	case "SweepFlag":
		res = __gong__marshallBool(ident, "SweepFlag", shiftedleftstackgrowthcurveendarcshape.SweepFlag)
	case "RadiusX":
		res = __gong__marshallFloat(ident, "RadiusX", shiftedleftstackgrowthcurveendarcshape.RadiusX)
	case "RadiusY":
		res = __gong__marshallFloat(ident, "RadiusY", shiftedleftstackgrowthcurveendarcshape.RadiusY)

	default:
		log.Panicf("Unknown field %s for Gongstruct ShiftedLeftStackGrowthCurveEndArcShape", fieldName)
	}
	return
}

func (shiftedleftstackgrowthcurvestartarcshape *ShiftedLeftStackGrowthCurveStartArcShape) GongMarshallField(stage *Stage, fieldName string) (res string) {
	ident := shiftedleftstackgrowthcurvestartarcshape.GongGetIdentifier(stage)
	_ = ident
	switch fieldName {
	case "Name":
		res = __gong__marshallString(ident, "Name", shiftedleftstackgrowthcurvestartarcshape.Name)
	case "StartX":
		res = __gong__marshallFloat(ident, "StartX", shiftedleftstackgrowthcurvestartarcshape.StartX)
	case "StartY":
		res = __gong__marshallFloat(ident, "StartY", shiftedleftstackgrowthcurvestartarcshape.StartY)
	case "EndX":
		res = __gong__marshallFloat(ident, "EndX", shiftedleftstackgrowthcurvestartarcshape.EndX)
	case "EndY":
		res = __gong__marshallFloat(ident, "EndY", shiftedleftstackgrowthcurvestartarcshape.EndY)
	case "XAxisRotation":
		res = __gong__marshallFloat(ident, "XAxisRotation", shiftedleftstackgrowthcurvestartarcshape.XAxisRotation)
	case "LargeArcFlag":
		res = __gong__marshallBool(ident, "LargeArcFlag", shiftedleftstackgrowthcurvestartarcshape.LargeArcFlag)
	case "SweepFlag":
		res = __gong__marshallBool(ident, "SweepFlag", shiftedleftstackgrowthcurvestartarcshape.SweepFlag)
	case "RadiusX":
		res = __gong__marshallFloat(ident, "RadiusX", shiftedleftstackgrowthcurvestartarcshape.RadiusX)
	case "RadiusY":
		res = __gong__marshallFloat(ident, "RadiusY", shiftedleftstackgrowthcurvestartarcshape.RadiusY)

	default:
		log.Panicf("Unknown field %s for Gongstruct ShiftedLeftStackGrowthCurveStartArcShape", fieldName)
	}
	return
}

func (shiftedleftstacknormalvector *ShiftedLeftStackNormalVector) GongMarshallField(stage *Stage, fieldName string) (res string) {
	ident := shiftedleftstacknormalvector.GongGetIdentifier(stage)
	_ = ident
	switch fieldName {
	case "Name":
		res = __gong__marshallString(ident, "Name", shiftedleftstacknormalvector.Name)
	case "StartX":
		res = __gong__marshallFloat(ident, "StartX", shiftedleftstacknormalvector.StartX)
	case "StartY":
		res = __gong__marshallFloat(ident, "StartY", shiftedleftstacknormalvector.StartY)
	case "EndX":
		res = __gong__marshallFloat(ident, "EndX", shiftedleftstacknormalvector.EndX)
	case "EndY":
		res = __gong__marshallFloat(ident, "EndY", shiftedleftstacknormalvector.EndY)

	default:
		log.Panicf("Unknown field %s for Gongstruct ShiftedLeftStackNormalVector", fieldName)
	}
	return
}

func (shiftedleftstackofgrowthcurve *ShiftedLeftStackOfGrowthCurve) GongMarshallField(stage *Stage, fieldName string) (res string) {
	ident := shiftedleftstackofgrowthcurve.GongGetIdentifier(stage)
	_ = ident
	switch fieldName {
	case "Name":
		res = __gong__marshallString(ident, "Name", shiftedleftstackofgrowthcurve.Name)

	default:
		log.Panicf("Unknown field %s for Gongstruct ShiftedLeftStackOfGrowthCurve", fieldName)
	}
	return
}

func (shiftedleftstackofnormalvector *ShiftedLeftStackOfNormalVector) GongMarshallField(stage *Stage, fieldName string) (res string) {
	ident := shiftedleftstackofnormalvector.GongGetIdentifier(stage)
	_ = ident
	switch fieldName {
	case "Name":
		res = __gong__marshallString(ident, "Name", shiftedleftstackofnormalvector.Name)

	default:
		log.Panicf("Unknown field %s for Gongstruct ShiftedLeftStackOfNormalVector", fieldName)
	}
	return
}

func (shiftedrightgrowthcurve2dribbon *ShiftedRightGrowthCurve2DRibbon) GongMarshallField(stage *Stage, fieldName string) (res string) {
	ident := shiftedrightgrowthcurve2dribbon.GongGetIdentifier(stage)
	_ = ident
	switch fieldName {
	case "Name":
		res = __gong__marshallString(ident, "Name", shiftedrightgrowthcurve2dribbon.Name)

	default:
		log.Panicf("Unknown field %s for Gongstruct ShiftedRightGrowthCurve2DRibbon", fieldName)
	}
	return
}

func (shiftedrightgrowthcurve2dribbonendshape *ShiftedRightGrowthCurve2DRibbonEndShape) GongMarshallField(stage *Stage, fieldName string) (res string) {
	ident := shiftedrightgrowthcurve2dribbonendshape.GongGetIdentifier(stage)
	_ = ident
	switch fieldName {
	case "Name":
		res = __gong__marshallString(ident, "Name", shiftedrightgrowthcurve2dribbonendshape.Name)
	case "BottomStartX":
		res = __gong__marshallFloat(ident, "BottomStartX", shiftedrightgrowthcurve2dribbonendshape.BottomStartX)
	case "BottomStartY":
		res = __gong__marshallFloat(ident, "BottomStartY", shiftedrightgrowthcurve2dribbonendshape.BottomStartY)
	case "BottomEndX":
		res = __gong__marshallFloat(ident, "BottomEndX", shiftedrightgrowthcurve2dribbonendshape.BottomEndX)
	case "BottomEndY":
		res = __gong__marshallFloat(ident, "BottomEndY", shiftedrightgrowthcurve2dribbonendshape.BottomEndY)
	case "BottomRadiusX":
		res = __gong__marshallFloat(ident, "BottomRadiusX", shiftedrightgrowthcurve2dribbonendshape.BottomRadiusX)
	case "BottomRadiusY":
		res = __gong__marshallFloat(ident, "BottomRadiusY", shiftedrightgrowthcurve2dribbonendshape.BottomRadiusY)
	case "BottomXAxisRotation":
		res = __gong__marshallFloat(ident, "BottomXAxisRotation", shiftedrightgrowthcurve2dribbonendshape.BottomXAxisRotation)
	case "BottomLargeArcFlag":
		res = __gong__marshallBool(ident, "BottomLargeArcFlag", shiftedrightgrowthcurve2dribbonendshape.BottomLargeArcFlag)
	case "BottomSweepFlag":
		res = __gong__marshallBool(ident, "BottomSweepFlag", shiftedrightgrowthcurve2dribbonendshape.BottomSweepFlag)
	case "TopStartX":
		res = __gong__marshallFloat(ident, "TopStartX", shiftedrightgrowthcurve2dribbonendshape.TopStartX)
	case "TopStartY":
		res = __gong__marshallFloat(ident, "TopStartY", shiftedrightgrowthcurve2dribbonendshape.TopStartY)
	case "TopEndX":
		res = __gong__marshallFloat(ident, "TopEndX", shiftedrightgrowthcurve2dribbonendshape.TopEndX)
	case "TopEndY":
		res = __gong__marshallFloat(ident, "TopEndY", shiftedrightgrowthcurve2dribbonendshape.TopEndY)
	case "TopRadiusX":
		res = __gong__marshallFloat(ident, "TopRadiusX", shiftedrightgrowthcurve2dribbonendshape.TopRadiusX)
	case "TopRadiusY":
		res = __gong__marshallFloat(ident, "TopRadiusY", shiftedrightgrowthcurve2dribbonendshape.TopRadiusY)
	case "TopXAxisRotation":
		res = __gong__marshallFloat(ident, "TopXAxisRotation", shiftedrightgrowthcurve2dribbonendshape.TopXAxisRotation)
	case "TopLargeArcFlag":
		res = __gong__marshallBool(ident, "TopLargeArcFlag", shiftedrightgrowthcurve2dribbonendshape.TopLargeArcFlag)
	case "TopSweepFlag":
		res = __gong__marshallBool(ident, "TopSweepFlag", shiftedrightgrowthcurve2dribbonendshape.TopSweepFlag)

	default:
		log.Panicf("Unknown field %s for Gongstruct ShiftedRightGrowthCurve2DRibbonEndShape", fieldName)
	}
	return
}

func (shiftedrightgrowthcurve2dribbonstartshape *ShiftedRightGrowthCurve2DRibbonStartShape) GongMarshallField(stage *Stage, fieldName string) (res string) {
	ident := shiftedrightgrowthcurve2dribbonstartshape.GongGetIdentifier(stage)
	_ = ident
	switch fieldName {
	case "Name":
		res = __gong__marshallString(ident, "Name", shiftedrightgrowthcurve2dribbonstartshape.Name)
	case "BottomStartX":
		res = __gong__marshallFloat(ident, "BottomStartX", shiftedrightgrowthcurve2dribbonstartshape.BottomStartX)
	case "BottomStartY":
		res = __gong__marshallFloat(ident, "BottomStartY", shiftedrightgrowthcurve2dribbonstartshape.BottomStartY)
	case "BottomEndX":
		res = __gong__marshallFloat(ident, "BottomEndX", shiftedrightgrowthcurve2dribbonstartshape.BottomEndX)
	case "BottomEndY":
		res = __gong__marshallFloat(ident, "BottomEndY", shiftedrightgrowthcurve2dribbonstartshape.BottomEndY)
	case "BottomRadiusX":
		res = __gong__marshallFloat(ident, "BottomRadiusX", shiftedrightgrowthcurve2dribbonstartshape.BottomRadiusX)
	case "BottomRadiusY":
		res = __gong__marshallFloat(ident, "BottomRadiusY", shiftedrightgrowthcurve2dribbonstartshape.BottomRadiusY)
	case "BottomXAxisRotation":
		res = __gong__marshallFloat(ident, "BottomXAxisRotation", shiftedrightgrowthcurve2dribbonstartshape.BottomXAxisRotation)
	case "BottomLargeArcFlag":
		res = __gong__marshallBool(ident, "BottomLargeArcFlag", shiftedrightgrowthcurve2dribbonstartshape.BottomLargeArcFlag)
	case "BottomSweepFlag":
		res = __gong__marshallBool(ident, "BottomSweepFlag", shiftedrightgrowthcurve2dribbonstartshape.BottomSweepFlag)
	case "TopStartX":
		res = __gong__marshallFloat(ident, "TopStartX", shiftedrightgrowthcurve2dribbonstartshape.TopStartX)
	case "TopStartY":
		res = __gong__marshallFloat(ident, "TopStartY", shiftedrightgrowthcurve2dribbonstartshape.TopStartY)
	case "TopEndX":
		res = __gong__marshallFloat(ident, "TopEndX", shiftedrightgrowthcurve2dribbonstartshape.TopEndX)
	case "TopEndY":
		res = __gong__marshallFloat(ident, "TopEndY", shiftedrightgrowthcurve2dribbonstartshape.TopEndY)
	case "TopRadiusX":
		res = __gong__marshallFloat(ident, "TopRadiusX", shiftedrightgrowthcurve2dribbonstartshape.TopRadiusX)
	case "TopRadiusY":
		res = __gong__marshallFloat(ident, "TopRadiusY", shiftedrightgrowthcurve2dribbonstartshape.TopRadiusY)
	case "TopXAxisRotation":
		res = __gong__marshallFloat(ident, "TopXAxisRotation", shiftedrightgrowthcurve2dribbonstartshape.TopXAxisRotation)
	case "TopLargeArcFlag":
		res = __gong__marshallBool(ident, "TopLargeArcFlag", shiftedrightgrowthcurve2dribbonstartshape.TopLargeArcFlag)
	case "TopSweepFlag":
		res = __gong__marshallBool(ident, "TopSweepFlag", shiftedrightgrowthcurve2dribbonstartshape.TopSweepFlag)

	default:
		log.Panicf("Unknown field %s for Gongstruct ShiftedRightGrowthCurve2DRibbonStartShape", fieldName)
	}
	return
}

func (stackgrowthcurve2dendhalfwayarcshape *StackGrowthCurve2DEndHalfwayArcShape) GongMarshallField(stage *Stage, fieldName string) (res string) {
	ident := stackgrowthcurve2dendhalfwayarcshape.GongGetIdentifier(stage)
	_ = ident
	switch fieldName {
	case "Name":
		res = __gong__marshallString(ident, "Name", stackgrowthcurve2dendhalfwayarcshape.Name)
	case "StartX":
		res = __gong__marshallFloat(ident, "StartX", stackgrowthcurve2dendhalfwayarcshape.StartX)
	case "StartY":
		res = __gong__marshallFloat(ident, "StartY", stackgrowthcurve2dendhalfwayarcshape.StartY)
	case "EndX":
		res = __gong__marshallFloat(ident, "EndX", stackgrowthcurve2dendhalfwayarcshape.EndX)
	case "EndY":
		res = __gong__marshallFloat(ident, "EndY", stackgrowthcurve2dendhalfwayarcshape.EndY)
	case "RadiusX":
		res = __gong__marshallFloat(ident, "RadiusX", stackgrowthcurve2dendhalfwayarcshape.RadiusX)
	case "RadiusY":
		res = __gong__marshallFloat(ident, "RadiusY", stackgrowthcurve2dendhalfwayarcshape.RadiusY)
	case "XAxisRotation":
		res = __gong__marshallFloat(ident, "XAxisRotation", stackgrowthcurve2dendhalfwayarcshape.XAxisRotation)
	case "LargeArcFlag":
		res = __gong__marshallBool(ident, "LargeArcFlag", stackgrowthcurve2dendhalfwayarcshape.LargeArcFlag)
	case "SweepFlag":
		res = __gong__marshallBool(ident, "SweepFlag", stackgrowthcurve2dendhalfwayarcshape.SweepFlag)

	default:
		log.Panicf("Unknown field %s for Gongstruct StackGrowthCurve2DEndHalfwayArcShape", fieldName)
	}
	return
}

func (stackgrowthcurve2dribbonendshape *StackGrowthCurve2DRibbonEndShape) GongMarshallField(stage *Stage, fieldName string) (res string) {
	ident := stackgrowthcurve2dribbonendshape.GongGetIdentifier(stage)
	_ = ident
	switch fieldName {
	case "Name":
		res = __gong__marshallString(ident, "Name", stackgrowthcurve2dribbonendshape.Name)
	case "BottomStartX":
		res = __gong__marshallFloat(ident, "BottomStartX", stackgrowthcurve2dribbonendshape.BottomStartX)
	case "BottomStartY":
		res = __gong__marshallFloat(ident, "BottomStartY", stackgrowthcurve2dribbonendshape.BottomStartY)
	case "BottomEndX":
		res = __gong__marshallFloat(ident, "BottomEndX", stackgrowthcurve2dribbonendshape.BottomEndX)
	case "BottomEndY":
		res = __gong__marshallFloat(ident, "BottomEndY", stackgrowthcurve2dribbonendshape.BottomEndY)
	case "BottomRadiusX":
		res = __gong__marshallFloat(ident, "BottomRadiusX", stackgrowthcurve2dribbonendshape.BottomRadiusX)
	case "BottomRadiusY":
		res = __gong__marshallFloat(ident, "BottomRadiusY", stackgrowthcurve2dribbonendshape.BottomRadiusY)
	case "BottomXAxisRotation":
		res = __gong__marshallFloat(ident, "BottomXAxisRotation", stackgrowthcurve2dribbonendshape.BottomXAxisRotation)
	case "BottomLargeArcFlag":
		res = __gong__marshallBool(ident, "BottomLargeArcFlag", stackgrowthcurve2dribbonendshape.BottomLargeArcFlag)
	case "BottomSweepFlag":
		res = __gong__marshallBool(ident, "BottomSweepFlag", stackgrowthcurve2dribbonendshape.BottomSweepFlag)
	case "TopStartX":
		res = __gong__marshallFloat(ident, "TopStartX", stackgrowthcurve2dribbonendshape.TopStartX)
	case "TopStartY":
		res = __gong__marshallFloat(ident, "TopStartY", stackgrowthcurve2dribbonendshape.TopStartY)
	case "TopEndX":
		res = __gong__marshallFloat(ident, "TopEndX", stackgrowthcurve2dribbonendshape.TopEndX)
	case "TopEndY":
		res = __gong__marshallFloat(ident, "TopEndY", stackgrowthcurve2dribbonendshape.TopEndY)
	case "TopRadiusX":
		res = __gong__marshallFloat(ident, "TopRadiusX", stackgrowthcurve2dribbonendshape.TopRadiusX)
	case "TopRadiusY":
		res = __gong__marshallFloat(ident, "TopRadiusY", stackgrowthcurve2dribbonendshape.TopRadiusY)
	case "TopXAxisRotation":
		res = __gong__marshallFloat(ident, "TopXAxisRotation", stackgrowthcurve2dribbonendshape.TopXAxisRotation)
	case "TopLargeArcFlag":
		res = __gong__marshallBool(ident, "TopLargeArcFlag", stackgrowthcurve2dribbonendshape.TopLargeArcFlag)
	case "TopSweepFlag":
		res = __gong__marshallBool(ident, "TopSweepFlag", stackgrowthcurve2dribbonendshape.TopSweepFlag)

	default:
		log.Panicf("Unknown field %s for Gongstruct StackGrowthCurve2DRibbonEndShape", fieldName)
	}
	return
}

func (stackgrowthcurve2dribbonstartshape *StackGrowthCurve2DRibbonStartShape) GongMarshallField(stage *Stage, fieldName string) (res string) {
	ident := stackgrowthcurve2dribbonstartshape.GongGetIdentifier(stage)
	_ = ident
	switch fieldName {
	case "Name":
		res = __gong__marshallString(ident, "Name", stackgrowthcurve2dribbonstartshape.Name)
	case "BottomStartX":
		res = __gong__marshallFloat(ident, "BottomStartX", stackgrowthcurve2dribbonstartshape.BottomStartX)
	case "BottomStartY":
		res = __gong__marshallFloat(ident, "BottomStartY", stackgrowthcurve2dribbonstartshape.BottomStartY)
	case "BottomEndX":
		res = __gong__marshallFloat(ident, "BottomEndX", stackgrowthcurve2dribbonstartshape.BottomEndX)
	case "BottomEndY":
		res = __gong__marshallFloat(ident, "BottomEndY", stackgrowthcurve2dribbonstartshape.BottomEndY)
	case "BottomRadiusX":
		res = __gong__marshallFloat(ident, "BottomRadiusX", stackgrowthcurve2dribbonstartshape.BottomRadiusX)
	case "BottomRadiusY":
		res = __gong__marshallFloat(ident, "BottomRadiusY", stackgrowthcurve2dribbonstartshape.BottomRadiusY)
	case "BottomXAxisRotation":
		res = __gong__marshallFloat(ident, "BottomXAxisRotation", stackgrowthcurve2dribbonstartshape.BottomXAxisRotation)
	case "BottomLargeArcFlag":
		res = __gong__marshallBool(ident, "BottomLargeArcFlag", stackgrowthcurve2dribbonstartshape.BottomLargeArcFlag)
	case "BottomSweepFlag":
		res = __gong__marshallBool(ident, "BottomSweepFlag", stackgrowthcurve2dribbonstartshape.BottomSweepFlag)
	case "TopStartX":
		res = __gong__marshallFloat(ident, "TopStartX", stackgrowthcurve2dribbonstartshape.TopStartX)
	case "TopStartY":
		res = __gong__marshallFloat(ident, "TopStartY", stackgrowthcurve2dribbonstartshape.TopStartY)
	case "TopEndX":
		res = __gong__marshallFloat(ident, "TopEndX", stackgrowthcurve2dribbonstartshape.TopEndX)
	case "TopEndY":
		res = __gong__marshallFloat(ident, "TopEndY", stackgrowthcurve2dribbonstartshape.TopEndY)
	case "TopRadiusX":
		res = __gong__marshallFloat(ident, "TopRadiusX", stackgrowthcurve2dribbonstartshape.TopRadiusX)
	case "TopRadiusY":
		res = __gong__marshallFloat(ident, "TopRadiusY", stackgrowthcurve2dribbonstartshape.TopRadiusY)
	case "TopXAxisRotation":
		res = __gong__marshallFloat(ident, "TopXAxisRotation", stackgrowthcurve2dribbonstartshape.TopXAxisRotation)
	case "TopLargeArcFlag":
		res = __gong__marshallBool(ident, "TopLargeArcFlag", stackgrowthcurve2dribbonstartshape.TopLargeArcFlag)
	case "TopSweepFlag":
		res = __gong__marshallBool(ident, "TopSweepFlag", stackgrowthcurve2dribbonstartshape.TopSweepFlag)

	default:
		log.Panicf("Unknown field %s for Gongstruct StackGrowthCurve2DRibbonStartShape", fieldName)
	}
	return
}

func (stackgrowthcurve2dstarthalfwayarcshape *StackGrowthCurve2DStartHalfwayArcShape) GongMarshallField(stage *Stage, fieldName string) (res string) {
	ident := stackgrowthcurve2dstarthalfwayarcshape.GongGetIdentifier(stage)
	_ = ident
	switch fieldName {
	case "Name":
		res = __gong__marshallString(ident, "Name", stackgrowthcurve2dstarthalfwayarcshape.Name)
	case "StartX":
		res = __gong__marshallFloat(ident, "StartX", stackgrowthcurve2dstarthalfwayarcshape.StartX)
	case "StartY":
		res = __gong__marshallFloat(ident, "StartY", stackgrowthcurve2dstarthalfwayarcshape.StartY)
	case "EndX":
		res = __gong__marshallFloat(ident, "EndX", stackgrowthcurve2dstarthalfwayarcshape.EndX)
	case "EndY":
		res = __gong__marshallFloat(ident, "EndY", stackgrowthcurve2dstarthalfwayarcshape.EndY)
	case "RadiusX":
		res = __gong__marshallFloat(ident, "RadiusX", stackgrowthcurve2dstarthalfwayarcshape.RadiusX)
	case "RadiusY":
		res = __gong__marshallFloat(ident, "RadiusY", stackgrowthcurve2dstarthalfwayarcshape.RadiusY)
	case "XAxisRotation":
		res = __gong__marshallFloat(ident, "XAxisRotation", stackgrowthcurve2dstarthalfwayarcshape.XAxisRotation)
	case "LargeArcFlag":
		res = __gong__marshallBool(ident, "LargeArcFlag", stackgrowthcurve2dstarthalfwayarcshape.LargeArcFlag)
	case "SweepFlag":
		res = __gong__marshallBool(ident, "SweepFlag", stackgrowthcurve2dstarthalfwayarcshape.SweepFlag)

	default:
		log.Panicf("Unknown field %s for Gongstruct StackGrowthCurve2DStartHalfwayArcShape", fieldName)
	}
	return
}

func (stackofgrowthcurve2d *StackOfGrowthCurve2D) GongMarshallField(stage *Stage, fieldName string) (res string) {
	ident := stackofgrowthcurve2d.GongGetIdentifier(stage)
	_ = ident
	switch fieldName {
	case "Name":
		res = __gong__marshallString(ident, "Name", stackofgrowthcurve2d.Name)

	default:
		log.Panicf("Unknown field %s for Gongstruct StackOfGrowthCurve2D", fieldName)
	}
	return
}

func (stackofgrowthcurve2dbygrowthvector *StackOfGrowthCurve2DByGrowthVector) GongMarshallField(stage *Stage, fieldName string) (res string) {
	ident := stackofgrowthcurve2dbygrowthvector.GongGetIdentifier(stage)
	_ = ident
	switch fieldName {
	case "Name":
		res = __gong__marshallString(ident, "Name", stackofgrowthcurve2dbygrowthvector.Name)

	default:
		log.Panicf("Unknown field %s for Gongstruct StackOfGrowthCurve2DByGrowthVector", fieldName)
	}
	return
}

func (stackofgrowthcurve2dribbon *StackOfGrowthCurve2DRibbon) GongMarshallField(stage *Stage, fieldName string) (res string) {
	ident := stackofgrowthcurve2dribbon.GongGetIdentifier(stage)
	_ = ident
	switch fieldName {
	case "Name":
		res = __gong__marshallString(ident, "Name", stackofgrowthcurve2dribbon.Name)

	default:
		log.Panicf("Unknown field %s for Gongstruct StackOfGrowthCurve2DRibbon", fieldName)
	}
	return
}

func (stackofpartiallyrotatedtorusshape *StackOfPartiallyRotatedTorusShape) GongMarshallField(stage *Stage, fieldName string) (res string) {
	ident := stackofpartiallyrotatedtorusshape.GongGetIdentifier(stage)
	_ = ident
	switch fieldName {
	case "Name":
		res = __gong__marshallString(ident, "Name", stackofpartiallyrotatedtorusshape.Name)

	default:
		log.Panicf("Unknown field %s for Gongstruct StackOfPartiallyRotatedTorusShape", fieldName)
	}
	return
}

func (stackofrotatedcarvedoutvasetrapezeringsshape *StackOfRotatedCarvedOutVaseTrapezeRingsShape) GongMarshallField(stage *Stage, fieldName string) (res string) {
	ident := stackofrotatedcarvedoutvasetrapezeringsshape.GongGetIdentifier(stage)
	_ = ident
	switch fieldName {
	case "Name":
		res = __gong__marshallString(ident, "Name", stackofrotatedcarvedoutvasetrapezeringsshape.Name)

	default:
		log.Panicf("Unknown field %s for Gongstruct StackOfRotatedCarvedOutVaseTrapezeRingsShape", fieldName)
	}
	return
}

func (stackofrotatedgrowthcurve2d *StackOfRotatedGrowthCurve2D) GongMarshallField(stage *Stage, fieldName string) (res string) {
	ident := stackofrotatedgrowthcurve2d.GongGetIdentifier(stage)
	_ = ident
	switch fieldName {
	case "Name":
		res = __gong__marshallString(ident, "Name", stackofrotatedgrowthcurve2d.Name)

	default:
		log.Panicf("Unknown field %s for Gongstruct StackOfRotatedGrowthCurve2D", fieldName)
	}
	return
}

func (stackofrotatedgrowthcurve2dribbon *StackOfRotatedGrowthCurve2DRibbon) GongMarshallField(stage *Stage, fieldName string) (res string) {
	ident := stackofrotatedgrowthcurve2dribbon.GongGetIdentifier(stage)
	_ = ident
	switch fieldName {
	case "Name":
		res = __gong__marshallString(ident, "Name", stackofrotatedgrowthcurve2dribbon.Name)

	default:
		log.Panicf("Unknown field %s for Gongstruct StackOfRotatedGrowthCurve2DRibbon", fieldName)
	}
	return
}

func (stackrotatedgrowthcurve2dendarcshape *StackRotatedGrowthCurve2DEndArcShape) GongMarshallField(stage *Stage, fieldName string) (res string) {
	ident := stackrotatedgrowthcurve2dendarcshape.GongGetIdentifier(stage)
	_ = ident
	switch fieldName {
	case "Name":
		res = __gong__marshallString(ident, "Name", stackrotatedgrowthcurve2dendarcshape.Name)
	case "StartX":
		res = __gong__marshallFloat(ident, "StartX", stackrotatedgrowthcurve2dendarcshape.StartX)
	case "StartY":
		res = __gong__marshallFloat(ident, "StartY", stackrotatedgrowthcurve2dendarcshape.StartY)
	case "EndX":
		res = __gong__marshallFloat(ident, "EndX", stackrotatedgrowthcurve2dendarcshape.EndX)
	case "EndY":
		res = __gong__marshallFloat(ident, "EndY", stackrotatedgrowthcurve2dendarcshape.EndY)
	case "XAxisRotation":
		res = __gong__marshallFloat(ident, "XAxisRotation", stackrotatedgrowthcurve2dendarcshape.XAxisRotation)
	case "LargeArcFlag":
		res = __gong__marshallBool(ident, "LargeArcFlag", stackrotatedgrowthcurve2dendarcshape.LargeArcFlag)
	case "SweepFlag":
		res = __gong__marshallBool(ident, "SweepFlag", stackrotatedgrowthcurve2dendarcshape.SweepFlag)
	case "RadiusX":
		res = __gong__marshallFloat(ident, "RadiusX", stackrotatedgrowthcurve2dendarcshape.RadiusX)
	case "RadiusY":
		res = __gong__marshallFloat(ident, "RadiusY", stackrotatedgrowthcurve2dendarcshape.RadiusY)

	default:
		log.Panicf("Unknown field %s for Gongstruct StackRotatedGrowthCurve2DEndArcShape", fieldName)
	}
	return
}

func (stackrotatedgrowthcurve2dribbonendshape *StackRotatedGrowthCurve2DRibbonEndShape) GongMarshallField(stage *Stage, fieldName string) (res string) {
	ident := stackrotatedgrowthcurve2dribbonendshape.GongGetIdentifier(stage)
	_ = ident
	switch fieldName {
	case "Name":
		res = __gong__marshallString(ident, "Name", stackrotatedgrowthcurve2dribbonendshape.Name)
	case "BottomStartX":
		res = __gong__marshallFloat(ident, "BottomStartX", stackrotatedgrowthcurve2dribbonendshape.BottomStartX)
	case "BottomStartY":
		res = __gong__marshallFloat(ident, "BottomStartY", stackrotatedgrowthcurve2dribbonendshape.BottomStartY)
	case "BottomEndX":
		res = __gong__marshallFloat(ident, "BottomEndX", stackrotatedgrowthcurve2dribbonendshape.BottomEndX)
	case "BottomEndY":
		res = __gong__marshallFloat(ident, "BottomEndY", stackrotatedgrowthcurve2dribbonendshape.BottomEndY)
	case "BottomRadiusX":
		res = __gong__marshallFloat(ident, "BottomRadiusX", stackrotatedgrowthcurve2dribbonendshape.BottomRadiusX)
	case "BottomRadiusY":
		res = __gong__marshallFloat(ident, "BottomRadiusY", stackrotatedgrowthcurve2dribbonendshape.BottomRadiusY)
	case "BottomXAxisRotation":
		res = __gong__marshallFloat(ident, "BottomXAxisRotation", stackrotatedgrowthcurve2dribbonendshape.BottomXAxisRotation)
	case "BottomLargeArcFlag":
		res = __gong__marshallBool(ident, "BottomLargeArcFlag", stackrotatedgrowthcurve2dribbonendshape.BottomLargeArcFlag)
	case "BottomSweepFlag":
		res = __gong__marshallBool(ident, "BottomSweepFlag", stackrotatedgrowthcurve2dribbonendshape.BottomSweepFlag)
	case "TopStartX":
		res = __gong__marshallFloat(ident, "TopStartX", stackrotatedgrowthcurve2dribbonendshape.TopStartX)
	case "TopStartY":
		res = __gong__marshallFloat(ident, "TopStartY", stackrotatedgrowthcurve2dribbonendshape.TopStartY)
	case "TopEndX":
		res = __gong__marshallFloat(ident, "TopEndX", stackrotatedgrowthcurve2dribbonendshape.TopEndX)
	case "TopEndY":
		res = __gong__marshallFloat(ident, "TopEndY", stackrotatedgrowthcurve2dribbonendshape.TopEndY)
	case "TopRadiusX":
		res = __gong__marshallFloat(ident, "TopRadiusX", stackrotatedgrowthcurve2dribbonendshape.TopRadiusX)
	case "TopRadiusY":
		res = __gong__marshallFloat(ident, "TopRadiusY", stackrotatedgrowthcurve2dribbonendshape.TopRadiusY)
	case "TopXAxisRotation":
		res = __gong__marshallFloat(ident, "TopXAxisRotation", stackrotatedgrowthcurve2dribbonendshape.TopXAxisRotation)
	case "TopLargeArcFlag":
		res = __gong__marshallBool(ident, "TopLargeArcFlag", stackrotatedgrowthcurve2dribbonendshape.TopLargeArcFlag)
	case "TopSweepFlag":
		res = __gong__marshallBool(ident, "TopSweepFlag", stackrotatedgrowthcurve2dribbonendshape.TopSweepFlag)

	default:
		log.Panicf("Unknown field %s for Gongstruct StackRotatedGrowthCurve2DRibbonEndShape", fieldName)
	}
	return
}

func (stackrotatedgrowthcurve2dribbonstartshape *StackRotatedGrowthCurve2DRibbonStartShape) GongMarshallField(stage *Stage, fieldName string) (res string) {
	ident := stackrotatedgrowthcurve2dribbonstartshape.GongGetIdentifier(stage)
	_ = ident
	switch fieldName {
	case "Name":
		res = __gong__marshallString(ident, "Name", stackrotatedgrowthcurve2dribbonstartshape.Name)
	case "BottomStartX":
		res = __gong__marshallFloat(ident, "BottomStartX", stackrotatedgrowthcurve2dribbonstartshape.BottomStartX)
	case "BottomStartY":
		res = __gong__marshallFloat(ident, "BottomStartY", stackrotatedgrowthcurve2dribbonstartshape.BottomStartY)
	case "BottomEndX":
		res = __gong__marshallFloat(ident, "BottomEndX", stackrotatedgrowthcurve2dribbonstartshape.BottomEndX)
	case "BottomEndY":
		res = __gong__marshallFloat(ident, "BottomEndY", stackrotatedgrowthcurve2dribbonstartshape.BottomEndY)
	case "BottomRadiusX":
		res = __gong__marshallFloat(ident, "BottomRadiusX", stackrotatedgrowthcurve2dribbonstartshape.BottomRadiusX)
	case "BottomRadiusY":
		res = __gong__marshallFloat(ident, "BottomRadiusY", stackrotatedgrowthcurve2dribbonstartshape.BottomRadiusY)
	case "BottomXAxisRotation":
		res = __gong__marshallFloat(ident, "BottomXAxisRotation", stackrotatedgrowthcurve2dribbonstartshape.BottomXAxisRotation)
	case "BottomLargeArcFlag":
		res = __gong__marshallBool(ident, "BottomLargeArcFlag", stackrotatedgrowthcurve2dribbonstartshape.BottomLargeArcFlag)
	case "BottomSweepFlag":
		res = __gong__marshallBool(ident, "BottomSweepFlag", stackrotatedgrowthcurve2dribbonstartshape.BottomSweepFlag)
	case "TopStartX":
		res = __gong__marshallFloat(ident, "TopStartX", stackrotatedgrowthcurve2dribbonstartshape.TopStartX)
	case "TopStartY":
		res = __gong__marshallFloat(ident, "TopStartY", stackrotatedgrowthcurve2dribbonstartshape.TopStartY)
	case "TopEndX":
		res = __gong__marshallFloat(ident, "TopEndX", stackrotatedgrowthcurve2dribbonstartshape.TopEndX)
	case "TopEndY":
		res = __gong__marshallFloat(ident, "TopEndY", stackrotatedgrowthcurve2dribbonstartshape.TopEndY)
	case "TopRadiusX":
		res = __gong__marshallFloat(ident, "TopRadiusX", stackrotatedgrowthcurve2dribbonstartshape.TopRadiusX)
	case "TopRadiusY":
		res = __gong__marshallFloat(ident, "TopRadiusY", stackrotatedgrowthcurve2dribbonstartshape.TopRadiusY)
	case "TopXAxisRotation":
		res = __gong__marshallFloat(ident, "TopXAxisRotation", stackrotatedgrowthcurve2dribbonstartshape.TopXAxisRotation)
	case "TopLargeArcFlag":
		res = __gong__marshallBool(ident, "TopLargeArcFlag", stackrotatedgrowthcurve2dribbonstartshape.TopLargeArcFlag)
	case "TopSweepFlag":
		res = __gong__marshallBool(ident, "TopSweepFlag", stackrotatedgrowthcurve2dribbonstartshape.TopSweepFlag)

	default:
		log.Panicf("Unknown field %s for Gongstruct StackRotatedGrowthCurve2DRibbonStartShape", fieldName)
	}
	return
}

func (stackrotatedgrowthcurve2dstartarcshape *StackRotatedGrowthCurve2DStartArcShape) GongMarshallField(stage *Stage, fieldName string) (res string) {
	ident := stackrotatedgrowthcurve2dstartarcshape.GongGetIdentifier(stage)
	_ = ident
	switch fieldName {
	case "Name":
		res = __gong__marshallString(ident, "Name", stackrotatedgrowthcurve2dstartarcshape.Name)
	case "StartX":
		res = __gong__marshallFloat(ident, "StartX", stackrotatedgrowthcurve2dstartarcshape.StartX)
	case "StartY":
		res = __gong__marshallFloat(ident, "StartY", stackrotatedgrowthcurve2dstartarcshape.StartY)
	case "EndX":
		res = __gong__marshallFloat(ident, "EndX", stackrotatedgrowthcurve2dstartarcshape.EndX)
	case "EndY":
		res = __gong__marshallFloat(ident, "EndY", stackrotatedgrowthcurve2dstartarcshape.EndY)
	case "XAxisRotation":
		res = __gong__marshallFloat(ident, "XAxisRotation", stackrotatedgrowthcurve2dstartarcshape.XAxisRotation)
	case "LargeArcFlag":
		res = __gong__marshallBool(ident, "LargeArcFlag", stackrotatedgrowthcurve2dstartarcshape.LargeArcFlag)
	case "SweepFlag":
		res = __gong__marshallBool(ident, "SweepFlag", stackrotatedgrowthcurve2dstartarcshape.SweepFlag)
	case "RadiusX":
		res = __gong__marshallFloat(ident, "RadiusX", stackrotatedgrowthcurve2dstartarcshape.RadiusX)
	case "RadiusY":
		res = __gong__marshallFloat(ident, "RadiusY", stackrotatedgrowthcurve2dstartarcshape.RadiusY)

	default:
		log.Panicf("Unknown field %s for Gongstruct StackRotatedGrowthCurve2DStartArcShape", fieldName)
	}
	return
}

func (startarcshape *StartArcShape) GongMarshallField(stage *Stage, fieldName string) (res string) {
	ident := startarcshape.GongGetIdentifier(stage)
	_ = ident
	switch fieldName {
	case "Name":
		res = __gong__marshallString(ident, "Name", startarcshape.Name)
	case "StartX":
		res = __gong__marshallFloat(ident, "StartX", startarcshape.StartX)
	case "StartY":
		res = __gong__marshallFloat(ident, "StartY", startarcshape.StartY)
	case "EndX":
		res = __gong__marshallFloat(ident, "EndX", startarcshape.EndX)
	case "EndY":
		res = __gong__marshallFloat(ident, "EndY", startarcshape.EndY)
	case "XAxisRotation":
		res = __gong__marshallFloat(ident, "XAxisRotation", startarcshape.XAxisRotation)
	case "LargeArcFlag":
		res = __gong__marshallBool(ident, "LargeArcFlag", startarcshape.LargeArcFlag)
	case "SweepFlag":
		res = __gong__marshallBool(ident, "SweepFlag", startarcshape.SweepFlag)
	case "RadiusX":
		res = __gong__marshallFloat(ident, "RadiusX", startarcshape.RadiusX)
	case "RadiusY":
		res = __gong__marshallFloat(ident, "RadiusY", startarcshape.RadiusY)

	default:
		log.Panicf("Unknown field %s for Gongstruct StartArcShape", fieldName)
	}
	return
}

func (startarcshapegrid *StartArcShapeGrid) GongMarshallField(stage *Stage, fieldName string) (res string) {
	ident := startarcshapegrid.GongGetIdentifier(stage)
	_ = ident
	switch fieldName {
	case "Name":
		res = __gong__marshallString(ident, "Name", startarcshapegrid.Name)

	default:
		log.Panicf("Unknown field %s for Gongstruct StartArcShapeGrid", fieldName)
	}
	return
}

func (starthalfwayarcshape *StartHalfwayArcShape) GongMarshallField(stage *Stage, fieldName string) (res string) {
	ident := starthalfwayarcshape.GongGetIdentifier(stage)
	_ = ident
	switch fieldName {
	case "Name":
		res = __gong__marshallString(ident, "Name", starthalfwayarcshape.Name)
	case "StartX":
		res = __gong__marshallFloat(ident, "StartX", starthalfwayarcshape.StartX)
	case "StartY":
		res = __gong__marshallFloat(ident, "StartY", starthalfwayarcshape.StartY)
	case "EndX":
		res = __gong__marshallFloat(ident, "EndX", starthalfwayarcshape.EndX)
	case "EndY":
		res = __gong__marshallFloat(ident, "EndY", starthalfwayarcshape.EndY)
	case "RadiusX":
		res = __gong__marshallFloat(ident, "RadiusX", starthalfwayarcshape.RadiusX)
	case "RadiusY":
		res = __gong__marshallFloat(ident, "RadiusY", starthalfwayarcshape.RadiusY)
	case "XAxisRotation":
		res = __gong__marshallFloat(ident, "XAxisRotation", starthalfwayarcshape.XAxisRotation)
	case "LargeArcFlag":
		res = __gong__marshallBool(ident, "LargeArcFlag", starthalfwayarcshape.LargeArcFlag)
	case "SweepFlag":
		res = __gong__marshallBool(ident, "SweepFlag", starthalfwayarcshape.SweepFlag)

	default:
		log.Panicf("Unknown field %s for Gongstruct StartHalfwayArcShape", fieldName)
	}
	return
}

func (starthalfwayarcshapegrid *StartHalfwayArcShapeGrid) GongMarshallField(stage *Stage, fieldName string) (res string) {
	ident := starthalfwayarcshapegrid.GongGetIdentifier(stage)
	_ = ident
	switch fieldName {
	case "Name":
		res = __gong__marshallString(ident, "Name", starthalfwayarcshapegrid.Name)

	default:
		log.Panicf("Unknown field %s for Gongstruct StartHalfwayArcShapeGrid", fieldName)
	}
	return
}

func (stemcylinder3dshape *StemCylinder3DShape) GongMarshallField(stage *Stage, fieldName string) (res string) {
	ident := stemcylinder3dshape.GongGetIdentifier(stage)
	_ = ident
	switch fieldName {
	case "Name":
		res = __gong__marshallString(ident, "Name", stemcylinder3dshape.Name)
	case "Transparency":
		res = __gong__marshallFloat(ident, "Transparency", stemcylinder3dshape.Transparency)

	default:
		log.Panicf("Unknown field %s for Gongstruct StemCylinder3DShape", fieldName)
	}
	return
}

func (stool2ddiagram *Stool2DDiagram) GongMarshallField(stage *Stage, fieldName string) (res string) {
	ident := stool2ddiagram.GongGetIdentifier(stage)
	_ = ident
	switch fieldName {
	case "Name":
		res = __gong__marshallString(ident, "Name", stool2ddiagram.Name)
	case "Zoom":
		res = __gong__marshallFloat(ident, "Zoom", stool2ddiagram.Zoom)
	case "IsHiddenAxesShape":
		res = __gong__marshallBool(ident, "IsHiddenAxesShape", stool2ddiagram.IsHiddenAxesShape)
	case "IsChecked":
		res = __gong__marshallBool(ident, "IsChecked", stool2ddiagram.IsChecked)
	case "ComputedPrefix":
		res = __gong__marshallString(ident, "ComputedPrefix", stool2ddiagram.ComputedPrefix)
	case "IsExpanded":
		res = __gong__marshallBool(ident, "IsExpanded", stool2ddiagram.IsExpanded)

	default:
		log.Panicf("Unknown field %s for Gongstruct Stool2DDiagram", fieldName)
	}
	return
}

func (stool3ddiagram *Stool3DDiagram) GongMarshallField(stage *Stage, fieldName string) (res string) {
	ident := stool3ddiagram.GongGetIdentifier(stage)
	_ = ident
	switch fieldName {
	case "Name":
		res = __gong__marshallString(ident, "Name", stool3ddiagram.Name)
	case "IsHiddenSeatTopCurveShape":
		res = __gong__marshallBool(ident, "IsHiddenSeatTopCurveShape", stool3ddiagram.IsHiddenSeatTopCurveShape)
	case "IsHiddenRotatedSeatTopCurveShape":
		res = __gong__marshallBool(ident, "IsHiddenRotatedSeatTopCurveShape", stool3ddiagram.IsHiddenRotatedSeatTopCurveShape)
	case "IsHiddenSeatBottomCurveShape":
		res = __gong__marshallBool(ident, "IsHiddenSeatBottomCurveShape", stool3ddiagram.IsHiddenSeatBottomCurveShape)
	case "IsHiddenRotatedSeatBottomCurveShape":
		res = __gong__marshallBool(ident, "IsHiddenRotatedSeatBottomCurveShape", stool3ddiagram.IsHiddenRotatedSeatBottomCurveShape)
	case "IsHiddenTorus3DShape":
		res = __gong__marshallBool(ident, "IsHiddenTorus3DShape", stool3ddiagram.IsHiddenTorus3DShape)
	case "IsHiddenRotatedTorusShape":
		res = __gong__marshallBool(ident, "IsHiddenRotatedTorusShape", stool3ddiagram.IsHiddenRotatedTorusShape)
	case "IsHiddenSampledPoints3DShape":
		res = __gong__marshallBool(ident, "IsHiddenSampledPoints3DShape", stool3ddiagram.IsHiddenSampledPoints3DShape)
	case "IsHiddenRotatedSampledPoints3DShape":
		res = __gong__marshallBool(ident, "IsHiddenRotatedSampledPoints3DShape", stool3ddiagram.IsHiddenRotatedSampledPoints3DShape)
	case "IsHiddenEyeSampledPoints3DShape":
		res = __gong__marshallBool(ident, "IsHiddenEyeSampledPoints3DShape", stool3ddiagram.IsHiddenEyeSampledPoints3DShape)
	case "IsHiddenEyeCornersSampledPoints3DShape":
		res = __gong__marshallBool(ident, "IsHiddenEyeCornersSampledPoints3DShape", stool3ddiagram.IsHiddenEyeCornersSampledPoints3DShape)
	case "IsHiddenEye3DShape":
		res = __gong__marshallBool(ident, "IsHiddenEye3DShape", stool3ddiagram.IsHiddenEye3DShape)
	case "IsHiddenEyeSeatBottomCurveShape":
		res = __gong__marshallBool(ident, "IsHiddenEyeSeatBottomCurveShape", stool3ddiagram.IsHiddenEyeSeatBottomCurveShape)
	case "IsHiddenEyeStoolBottomCurveShape":
		res = __gong__marshallBool(ident, "IsHiddenEyeStoolBottomCurveShape", stool3ddiagram.IsHiddenEyeStoolBottomCurveShape)
	case "IsHiddenSeat3DShape":
		res = __gong__marshallBool(ident, "IsHiddenSeat3DShape", stool3ddiagram.IsHiddenSeat3DShape)
	case "IsHiddenEyeVolume3DShape":
		res = __gong__marshallBool(ident, "IsHiddenEyeVolume3DShape", stool3ddiagram.IsHiddenEyeVolume3DShape)
	case "IsHiddenSeatAndLegs3DShape":
		res = __gong__marshallBool(ident, "IsHiddenSeatAndLegs3DShape", stool3ddiagram.IsHiddenSeatAndLegs3DShape)
	case "IsHiddenRotatedSeatAndLegs3DShape":
		res = __gong__marshallBool(ident, "IsHiddenRotatedSeatAndLegs3DShape", stool3ddiagram.IsHiddenRotatedSeatAndLegs3DShape)
	case "IsHiddenTiledFloor3DShape":
		res = __gong__marshallBool(ident, "IsHiddenTiledFloor3DShape", stool3ddiagram.IsHiddenTiledFloor3DShape)
	case "IsChecked":
		res = __gong__marshallBool(ident, "IsChecked", stool3ddiagram.IsChecked)
	case "ComputedPrefix":
		res = __gong__marshallString(ident, "ComputedPrefix", stool3ddiagram.ComputedPrefix)
	case "IsExpanded":
		res = __gong__marshallBool(ident, "IsExpanded", stool3ddiagram.IsExpanded)

	case "SampledPoints3DShape":
		if stool3ddiagram.SampledPoints3DShape != nil {
			res = __gong__marshallPointer(ident, "SampledPoints3DShape", stool3ddiagram.SampledPoints3DShape.GongGetIdentifier(stage))
		} else {
			res = __gong__marshallPointer(ident, "SampledPoints3DShape", "nil")
		}
	case "Rendered3DShape":
		if stool3ddiagram.Rendered3DShape != nil {
			res = __gong__marshallPointer(ident, "Rendered3DShape", stool3ddiagram.Rendered3DShape.GongGetIdentifier(stage))
		} else {
			res = __gong__marshallPointer(ident, "Rendered3DShape", "nil")
		}
	default:
		log.Panicf("Unknown field %s for Gongstruct Stool3DDiagram", fieldName)
	}
	return
}

func (tiledfloor3dshape *TiledFloor3DShape) GongMarshallField(stage *Stage, fieldName string) (res string) {
	ident := tiledfloor3dshape.GongGetIdentifier(stage)
	_ = ident
	switch fieldName {
	case "Name":
		res = __gong__marshallString(ident, "Name", tiledfloor3dshape.Name)

	default:
		log.Panicf("Unknown field %s for Gongstruct TiledFloor3DShape", fieldName)
	}
	return
}

func (topcurveplane1shape *TopCurvePlane1Shape) GongMarshallField(stage *Stage, fieldName string) (res string) {
	ident := topcurveplane1shape.GongGetIdentifier(stage)
	_ = ident
	switch fieldName {
	case "Name":
		res = __gong__marshallString(ident, "Name", topcurveplane1shape.Name)

	default:
		log.Panicf("Unknown field %s for Gongstruct TopCurvePlane1Shape", fieldName)
	}
	return
}

func (topcurveplane2shape *TopCurvePlane2Shape) GongMarshallField(stage *Stage, fieldName string) (res string) {
	ident := topcurveplane2shape.GongGetIdentifier(stage)
	_ = ident
	switch fieldName {
	case "Name":
		res = __gong__marshallString(ident, "Name", topcurveplane2shape.Name)

	default:
		log.Panicf("Unknown field %s for Gongstruct TopCurvePlane2Shape", fieldName)
	}
	return
}

func (topendarcshape *TopEndArcShape) GongMarshallField(stage *Stage, fieldName string) (res string) {
	ident := topendarcshape.GongGetIdentifier(stage)
	_ = ident
	switch fieldName {
	case "Name":
		res = __gong__marshallString(ident, "Name", topendarcshape.Name)
	case "StartX":
		res = __gong__marshallFloat(ident, "StartX", topendarcshape.StartX)
	case "StartY":
		res = __gong__marshallFloat(ident, "StartY", topendarcshape.StartY)
	case "EndX":
		res = __gong__marshallFloat(ident, "EndX", topendarcshape.EndX)
	case "EndY":
		res = __gong__marshallFloat(ident, "EndY", topendarcshape.EndY)
	case "XAxisRotation":
		res = __gong__marshallFloat(ident, "XAxisRotation", topendarcshape.XAxisRotation)
	case "LargeArcFlag":
		res = __gong__marshallBool(ident, "LargeArcFlag", topendarcshape.LargeArcFlag)
	case "SweepFlag":
		res = __gong__marshallBool(ident, "SweepFlag", topendarcshape.SweepFlag)
	case "RadiusX":
		res = __gong__marshallFloat(ident, "RadiusX", topendarcshape.RadiusX)
	case "RadiusY":
		res = __gong__marshallFloat(ident, "RadiusY", topendarcshape.RadiusY)

	default:
		log.Panicf("Unknown field %s for Gongstruct TopEndArcShape", fieldName)
	}
	return
}

func (topendarcshapegrid *TopEndArcShapeGrid) GongMarshallField(stage *Stage, fieldName string) (res string) {
	ident := topendarcshapegrid.GongGetIdentifier(stage)
	_ = ident
	switch fieldName {
	case "Name":
		res = __gong__marshallString(ident, "Name", topendarcshapegrid.Name)

	default:
		log.Panicf("Unknown field %s for Gongstruct TopEndArcShapeGrid", fieldName)
	}
	return
}

func (topendhalfwayarcshape *TopEndHalfwayArcShape) GongMarshallField(stage *Stage, fieldName string) (res string) {
	ident := topendhalfwayarcshape.GongGetIdentifier(stage)
	_ = ident
	switch fieldName {
	case "Name":
		res = __gong__marshallString(ident, "Name", topendhalfwayarcshape.Name)
	case "StartX":
		res = __gong__marshallFloat(ident, "StartX", topendhalfwayarcshape.StartX)
	case "StartY":
		res = __gong__marshallFloat(ident, "StartY", topendhalfwayarcshape.StartY)
	case "EndX":
		res = __gong__marshallFloat(ident, "EndX", topendhalfwayarcshape.EndX)
	case "EndY":
		res = __gong__marshallFloat(ident, "EndY", topendhalfwayarcshape.EndY)
	case "RadiusX":
		res = __gong__marshallFloat(ident, "RadiusX", topendhalfwayarcshape.RadiusX)
	case "RadiusY":
		res = __gong__marshallFloat(ident, "RadiusY", topendhalfwayarcshape.RadiusY)
	case "XAxisRotation":
		res = __gong__marshallFloat(ident, "XAxisRotation", topendhalfwayarcshape.XAxisRotation)
	case "LargeArcFlag":
		res = __gong__marshallBool(ident, "LargeArcFlag", topendhalfwayarcshape.LargeArcFlag)
	case "SweepFlag":
		res = __gong__marshallBool(ident, "SweepFlag", topendhalfwayarcshape.SweepFlag)

	default:
		log.Panicf("Unknown field %s for Gongstruct TopEndHalfwayArcShape", fieldName)
	}
	return
}

func (topendhalfwayarcshapegrid *TopEndHalfwayArcShapeGrid) GongMarshallField(stage *Stage, fieldName string) (res string) {
	ident := topendhalfwayarcshapegrid.GongGetIdentifier(stage)
	_ = ident
	switch fieldName {
	case "Name":
		res = __gong__marshallString(ident, "Name", topendhalfwayarcshapegrid.Name)

	default:
		log.Panicf("Unknown field %s for Gongstruct TopEndHalfwayArcShapeGrid", fieldName)
	}
	return
}

func (topgrowthcurve2d *TopGrowthCurve2D) GongMarshallField(stage *Stage, fieldName string) (res string) {
	ident := topgrowthcurve2d.GongGetIdentifier(stage)
	_ = ident
	switch fieldName {
	case "Name":
		res = __gong__marshallString(ident, "Name", topgrowthcurve2d.Name)

	default:
		log.Panicf("Unknown field %s for Gongstruct TopGrowthCurve2D", fieldName)
	}
	return
}

func (topmidarcvectorshape *TopMidArcVectorShape) GongMarshallField(stage *Stage, fieldName string) (res string) {
	ident := topmidarcvectorshape.GongGetIdentifier(stage)
	_ = ident
	switch fieldName {
	case "Name":
		res = __gong__marshallString(ident, "Name", topmidarcvectorshape.Name)
	case "StartX":
		res = __gong__marshallFloat(ident, "StartX", topmidarcvectorshape.StartX)
	case "StartY":
		res = __gong__marshallFloat(ident, "StartY", topmidarcvectorshape.StartY)
	case "EndX":
		res = __gong__marshallFloat(ident, "EndX", topmidarcvectorshape.EndX)
	case "EndY":
		res = __gong__marshallFloat(ident, "EndY", topmidarcvectorshape.EndY)

	default:
		log.Panicf("Unknown field %s for Gongstruct TopMidArcVectorShape", fieldName)
	}
	return
}

func (topmidarcvectorshapegrid *TopMidArcVectorShapeGrid) GongMarshallField(stage *Stage, fieldName string) (res string) {
	ident := topmidarcvectorshapegrid.GongGetIdentifier(stage)
	_ = ident
	switch fieldName {
	case "Name":
		res = __gong__marshallString(ident, "Name", topmidarcvectorshapegrid.Name)

	default:
		log.Panicf("Unknown field %s for Gongstruct TopMidArcVectorShapeGrid", fieldName)
	}
	return
}

func (topstackgrowthcurve2dendhalfwayarcshape *TopStackGrowthCurve2DEndHalfwayArcShape) GongMarshallField(stage *Stage, fieldName string) (res string) {
	ident := topstackgrowthcurve2dendhalfwayarcshape.GongGetIdentifier(stage)
	_ = ident
	switch fieldName {
	case "Name":
		res = __gong__marshallString(ident, "Name", topstackgrowthcurve2dendhalfwayarcshape.Name)
	case "StartX":
		res = __gong__marshallFloat(ident, "StartX", topstackgrowthcurve2dendhalfwayarcshape.StartX)
	case "StartY":
		res = __gong__marshallFloat(ident, "StartY", topstackgrowthcurve2dendhalfwayarcshape.StartY)
	case "EndX":
		res = __gong__marshallFloat(ident, "EndX", topstackgrowthcurve2dendhalfwayarcshape.EndX)
	case "EndY":
		res = __gong__marshallFloat(ident, "EndY", topstackgrowthcurve2dendhalfwayarcshape.EndY)
	case "RadiusX":
		res = __gong__marshallFloat(ident, "RadiusX", topstackgrowthcurve2dendhalfwayarcshape.RadiusX)
	case "RadiusY":
		res = __gong__marshallFloat(ident, "RadiusY", topstackgrowthcurve2dendhalfwayarcshape.RadiusY)
	case "XAxisRotation":
		res = __gong__marshallFloat(ident, "XAxisRotation", topstackgrowthcurve2dendhalfwayarcshape.XAxisRotation)
	case "LargeArcFlag":
		res = __gong__marshallBool(ident, "LargeArcFlag", topstackgrowthcurve2dendhalfwayarcshape.LargeArcFlag)
	case "SweepFlag":
		res = __gong__marshallBool(ident, "SweepFlag", topstackgrowthcurve2dendhalfwayarcshape.SweepFlag)

	default:
		log.Panicf("Unknown field %s for Gongstruct TopStackGrowthCurve2DEndHalfwayArcShape", fieldName)
	}
	return
}

func (topstackgrowthcurve2dstarthalfwayarcshape *TopStackGrowthCurve2DStartHalfwayArcShape) GongMarshallField(stage *Stage, fieldName string) (res string) {
	ident := topstackgrowthcurve2dstarthalfwayarcshape.GongGetIdentifier(stage)
	_ = ident
	switch fieldName {
	case "Name":
		res = __gong__marshallString(ident, "Name", topstackgrowthcurve2dstarthalfwayarcshape.Name)
	case "StartX":
		res = __gong__marshallFloat(ident, "StartX", topstackgrowthcurve2dstarthalfwayarcshape.StartX)
	case "StartY":
		res = __gong__marshallFloat(ident, "StartY", topstackgrowthcurve2dstarthalfwayarcshape.StartY)
	case "EndX":
		res = __gong__marshallFloat(ident, "EndX", topstackgrowthcurve2dstarthalfwayarcshape.EndX)
	case "EndY":
		res = __gong__marshallFloat(ident, "EndY", topstackgrowthcurve2dstarthalfwayarcshape.EndY)
	case "RadiusX":
		res = __gong__marshallFloat(ident, "RadiusX", topstackgrowthcurve2dstarthalfwayarcshape.RadiusX)
	case "RadiusY":
		res = __gong__marshallFloat(ident, "RadiusY", topstackgrowthcurve2dstarthalfwayarcshape.RadiusY)
	case "XAxisRotation":
		res = __gong__marshallFloat(ident, "XAxisRotation", topstackgrowthcurve2dstarthalfwayarcshape.XAxisRotation)
	case "LargeArcFlag":
		res = __gong__marshallBool(ident, "LargeArcFlag", topstackgrowthcurve2dstarthalfwayarcshape.LargeArcFlag)
	case "SweepFlag":
		res = __gong__marshallBool(ident, "SweepFlag", topstackgrowthcurve2dstarthalfwayarcshape.SweepFlag)

	default:
		log.Panicf("Unknown field %s for Gongstruct TopStackGrowthCurve2DStartHalfwayArcShape", fieldName)
	}
	return
}

func (topstackofgrowthcurve2d *TopStackOfGrowthCurve2D) GongMarshallField(stage *Stage, fieldName string) (res string) {
	ident := topstackofgrowthcurve2d.GongGetIdentifier(stage)
	_ = ident
	switch fieldName {
	case "Name":
		res = __gong__marshallString(ident, "Name", topstackofgrowthcurve2d.Name)

	default:
		log.Panicf("Unknown field %s for Gongstruct TopStackOfGrowthCurve2D", fieldName)
	}
	return
}

func (topstackofrotatedgrowthcurve2d *TopStackOfRotatedGrowthCurve2D) GongMarshallField(stage *Stage, fieldName string) (res string) {
	ident := topstackofrotatedgrowthcurve2d.GongGetIdentifier(stage)
	_ = ident
	switch fieldName {
	case "Name":
		res = __gong__marshallString(ident, "Name", topstackofrotatedgrowthcurve2d.Name)

	default:
		log.Panicf("Unknown field %s for Gongstruct TopStackOfRotatedGrowthCurve2D", fieldName)
	}
	return
}

func (topstackofrotatedgrowthcurve2dendarcshape *TopStackOfRotatedGrowthCurve2DEndArcShape) GongMarshallField(stage *Stage, fieldName string) (res string) {
	ident := topstackofrotatedgrowthcurve2dendarcshape.GongGetIdentifier(stage)
	_ = ident
	switch fieldName {
	case "Name":
		res = __gong__marshallString(ident, "Name", topstackofrotatedgrowthcurve2dendarcshape.Name)
	case "StartX":
		res = __gong__marshallFloat(ident, "StartX", topstackofrotatedgrowthcurve2dendarcshape.StartX)
	case "StartY":
		res = __gong__marshallFloat(ident, "StartY", topstackofrotatedgrowthcurve2dendarcshape.StartY)
	case "EndX":
		res = __gong__marshallFloat(ident, "EndX", topstackofrotatedgrowthcurve2dendarcshape.EndX)
	case "EndY":
		res = __gong__marshallFloat(ident, "EndY", topstackofrotatedgrowthcurve2dendarcshape.EndY)
	case "XAxisRotation":
		res = __gong__marshallFloat(ident, "XAxisRotation", topstackofrotatedgrowthcurve2dendarcshape.XAxisRotation)
	case "LargeArcFlag":
		res = __gong__marshallBool(ident, "LargeArcFlag", topstackofrotatedgrowthcurve2dendarcshape.LargeArcFlag)
	case "SweepFlag":
		res = __gong__marshallBool(ident, "SweepFlag", topstackofrotatedgrowthcurve2dendarcshape.SweepFlag)
	case "RadiusX":
		res = __gong__marshallFloat(ident, "RadiusX", topstackofrotatedgrowthcurve2dendarcshape.RadiusX)
	case "RadiusY":
		res = __gong__marshallFloat(ident, "RadiusY", topstackofrotatedgrowthcurve2dendarcshape.RadiusY)

	default:
		log.Panicf("Unknown field %s for Gongstruct TopStackOfRotatedGrowthCurve2DEndArcShape", fieldName)
	}
	return
}

func (topstackofrotatedgrowthcurve2dstartarcshape *TopStackOfRotatedGrowthCurve2DStartArcShape) GongMarshallField(stage *Stage, fieldName string) (res string) {
	ident := topstackofrotatedgrowthcurve2dstartarcshape.GongGetIdentifier(stage)
	_ = ident
	switch fieldName {
	case "Name":
		res = __gong__marshallString(ident, "Name", topstackofrotatedgrowthcurve2dstartarcshape.Name)
	case "StartX":
		res = __gong__marshallFloat(ident, "StartX", topstackofrotatedgrowthcurve2dstartarcshape.StartX)
	case "StartY":
		res = __gong__marshallFloat(ident, "StartY", topstackofrotatedgrowthcurve2dstartarcshape.StartY)
	case "EndX":
		res = __gong__marshallFloat(ident, "EndX", topstackofrotatedgrowthcurve2dstartarcshape.EndX)
	case "EndY":
		res = __gong__marshallFloat(ident, "EndY", topstackofrotatedgrowthcurve2dstartarcshape.EndY)
	case "XAxisRotation":
		res = __gong__marshallFloat(ident, "XAxisRotation", topstackofrotatedgrowthcurve2dstartarcshape.XAxisRotation)
	case "LargeArcFlag":
		res = __gong__marshallBool(ident, "LargeArcFlag", topstackofrotatedgrowthcurve2dstartarcshape.LargeArcFlag)
	case "SweepFlag":
		res = __gong__marshallBool(ident, "SweepFlag", topstackofrotatedgrowthcurve2dstartarcshape.SweepFlag)
	case "RadiusX":
		res = __gong__marshallFloat(ident, "RadiusX", topstackofrotatedgrowthcurve2dstartarcshape.RadiusX)
	case "RadiusY":
		res = __gong__marshallFloat(ident, "RadiusY", topstackofrotatedgrowthcurve2dstartarcshape.RadiusY)

	default:
		log.Panicf("Unknown field %s for Gongstruct TopStackOfRotatedGrowthCurve2DStartArcShape", fieldName)
	}
	return
}

func (topstartarcshape *TopStartArcShape) GongMarshallField(stage *Stage, fieldName string) (res string) {
	ident := topstartarcshape.GongGetIdentifier(stage)
	_ = ident
	switch fieldName {
	case "Name":
		res = __gong__marshallString(ident, "Name", topstartarcshape.Name)
	case "StartX":
		res = __gong__marshallFloat(ident, "StartX", topstartarcshape.StartX)
	case "StartY":
		res = __gong__marshallFloat(ident, "StartY", topstartarcshape.StartY)
	case "EndX":
		res = __gong__marshallFloat(ident, "EndX", topstartarcshape.EndX)
	case "EndY":
		res = __gong__marshallFloat(ident, "EndY", topstartarcshape.EndY)
	case "XAxisRotation":
		res = __gong__marshallFloat(ident, "XAxisRotation", topstartarcshape.XAxisRotation)
	case "LargeArcFlag":
		res = __gong__marshallBool(ident, "LargeArcFlag", topstartarcshape.LargeArcFlag)
	case "SweepFlag":
		res = __gong__marshallBool(ident, "SweepFlag", topstartarcshape.SweepFlag)
	case "RadiusX":
		res = __gong__marshallFloat(ident, "RadiusX", topstartarcshape.RadiusX)
	case "RadiusY":
		res = __gong__marshallFloat(ident, "RadiusY", topstartarcshape.RadiusY)

	default:
		log.Panicf("Unknown field %s for Gongstruct TopStartArcShape", fieldName)
	}
	return
}

func (topstartarcshapegrid *TopStartArcShapeGrid) GongMarshallField(stage *Stage, fieldName string) (res string) {
	ident := topstartarcshapegrid.GongGetIdentifier(stage)
	_ = ident
	switch fieldName {
	case "Name":
		res = __gong__marshallString(ident, "Name", topstartarcshapegrid.Name)

	default:
		log.Panicf("Unknown field %s for Gongstruct TopStartArcShapeGrid", fieldName)
	}
	return
}

func (topstarthalfwayarcshape *TopStartHalfwayArcShape) GongMarshallField(stage *Stage, fieldName string) (res string) {
	ident := topstarthalfwayarcshape.GongGetIdentifier(stage)
	_ = ident
	switch fieldName {
	case "Name":
		res = __gong__marshallString(ident, "Name", topstarthalfwayarcshape.Name)
	case "StartX":
		res = __gong__marshallFloat(ident, "StartX", topstarthalfwayarcshape.StartX)
	case "StartY":
		res = __gong__marshallFloat(ident, "StartY", topstarthalfwayarcshape.StartY)
	case "EndX":
		res = __gong__marshallFloat(ident, "EndX", topstarthalfwayarcshape.EndX)
	case "EndY":
		res = __gong__marshallFloat(ident, "EndY", topstarthalfwayarcshape.EndY)
	case "RadiusX":
		res = __gong__marshallFloat(ident, "RadiusX", topstarthalfwayarcshape.RadiusX)
	case "RadiusY":
		res = __gong__marshallFloat(ident, "RadiusY", topstarthalfwayarcshape.RadiusY)
	case "XAxisRotation":
		res = __gong__marshallFloat(ident, "XAxisRotation", topstarthalfwayarcshape.XAxisRotation)
	case "LargeArcFlag":
		res = __gong__marshallBool(ident, "LargeArcFlag", topstarthalfwayarcshape.LargeArcFlag)
	case "SweepFlag":
		res = __gong__marshallBool(ident, "SweepFlag", topstarthalfwayarcshape.SweepFlag)

	default:
		log.Panicf("Unknown field %s for Gongstruct TopStartHalfwayArcShape", fieldName)
	}
	return
}

func (topstarthalfwayarcshapegrid *TopStartHalfwayArcShapeGrid) GongMarshallField(stage *Stage, fieldName string) (res string) {
	ident := topstarthalfwayarcshapegrid.GongGetIdentifier(stage)
	_ = ident
	switch fieldName {
	case "Name":
		res = __gong__marshallString(ident, "Name", topstarthalfwayarcshapegrid.Name)

	default:
		log.Panicf("Unknown field %s for Gongstruct TopStartHalfwayArcShapeGrid", fieldName)
	}
	return
}

func (torus3dshape *Torus3DShape) GongMarshallField(stage *Stage, fieldName string) (res string) {
	ident := torus3dshape.GongGetIdentifier(stage)
	_ = ident
	switch fieldName {
	case "Name":
		res = __gong__marshallString(ident, "Name", torus3dshape.Name)

	default:
		log.Panicf("Unknown field %s for Gongstruct Torus3DShape", fieldName)
	}
	return
}

func (torusedge3dshape *TorusEdge3DShape) GongMarshallField(stage *Stage, fieldName string) (res string) {
	ident := torusedge3dshape.GongGetIdentifier(stage)
	_ = ident
	switch fieldName {
	case "Name":
		res = __gong__marshallString(ident, "Name", torusedge3dshape.Name)

	default:
		log.Panicf("Unknown field %s for Gongstruct TorusEdge3DShape", fieldName)
	}
	return
}

func (torusstackshape *TorusStackShape) GongMarshallField(stage *Stage, fieldName string) (res string) {
	ident := torusstackshape.GongGetIdentifier(stage)
	_ = ident
	switch fieldName {
	case "Name":
		res = __gong__marshallString(ident, "Name", torusstackshape.Name)

	default:
		log.Panicf("Unknown field %s for Gongstruct TorusStackShape", fieldName)
	}
	return
}

func (tubevase3ddiagram *TubeVase3DDiagram) GongMarshallField(stage *Stage, fieldName string) (res string) {
	ident := tubevase3ddiagram.GongGetIdentifier(stage)
	_ = ident
	switch fieldName {
	case "Name":
		res = __gong__marshallString(ident, "Name", tubevase3ddiagram.Name)
	case "IsHiddenStackOfPartiallyRotatedGrowthCurve2DRibbon":
		res = __gong__marshallBool(ident, "IsHiddenStackOfPartiallyRotatedGrowthCurve2DRibbon", tubevase3ddiagram.IsHiddenStackOfPartiallyRotatedGrowthCurve2DRibbon)
	case "IsHiddenTorusStackShape":
		res = __gong__marshallBool(ident, "IsHiddenTorusStackShape", tubevase3ddiagram.IsHiddenTorusStackShape)
	case "IsHiddenVerticalTorusStackShape":
		res = __gong__marshallBool(ident, "IsHiddenVerticalTorusStackShape", tubevase3ddiagram.IsHiddenVerticalTorusStackShape)
	case "IsHiddenPartiallyRotatedTorusShape":
		res = __gong__marshallBool(ident, "IsHiddenPartiallyRotatedTorusShape", tubevase3ddiagram.IsHiddenPartiallyRotatedTorusShape)
	case "IsHiddenStackOfPartiallyRotatedTorusShape":
		res = __gong__marshallBool(ident, "IsHiddenStackOfPartiallyRotatedTorusShape", tubevase3ddiagram.IsHiddenStackOfPartiallyRotatedTorusShape)
	case "IsHiddenPointsAndLines3DShape":
		res = __gong__marshallBool(ident, "IsHiddenPointsAndLines3DShape", tubevase3ddiagram.IsHiddenPointsAndLines3DShape)
	case "IsHiddenKeyHole3DShape":
		res = __gong__marshallBool(ident, "IsHiddenKeyHole3DShape", tubevase3ddiagram.IsHiddenKeyHole3DShape)
	case "IsHiddenKey3DShape":
		res = __gong__marshallBool(ident, "IsHiddenKey3DShape", tubevase3ddiagram.IsHiddenKey3DShape)
	case "IsHiddenVolumeKey3DShape":
		res = __gong__marshallBool(ident, "IsHiddenVolumeKey3DShape", tubevase3ddiagram.IsHiddenVolumeKey3DShape)
	case "IsHiddenTorusEdge3DShape":
		res = __gong__marshallBool(ident, "IsHiddenTorusEdge3DShape", tubevase3ddiagram.IsHiddenTorusEdge3DShape)
	case "IsHiddenSampledPoints3DShape":
		res = __gong__marshallBool(ident, "IsHiddenSampledPoints3DShape", tubevase3ddiagram.IsHiddenSampledPoints3DShape)
	case "IsHiddenOriginalPoints3DShape":
		res = __gong__marshallBool(ident, "IsHiddenOriginalPoints3DShape", tubevase3ddiagram.IsHiddenOriginalPoints3DShape)
	case "IsHiddenAngle0Shape":
		res = __gong__marshallBool(ident, "IsHiddenAngle0Shape", tubevase3ddiagram.IsHiddenAngle0Shape)
	case "IsHiddenTiledFloor3DShape":
		res = __gong__marshallBool(ident, "IsHiddenTiledFloor3DShape", tubevase3ddiagram.IsHiddenTiledFloor3DShape)
	case "IsHiddenTopCurvePlane1Shape":
		res = __gong__marshallBool(ident, "IsHiddenTopCurvePlane1Shape", tubevase3ddiagram.IsHiddenTopCurvePlane1Shape)
	case "IsHiddenBottomCurvePlane1Shape":
		res = __gong__marshallBool(ident, "IsHiddenBottomCurvePlane1Shape", tubevase3ddiagram.IsHiddenBottomCurvePlane1Shape)
	case "IsHiddenTopCurvePlane2Shape":
		res = __gong__marshallBool(ident, "IsHiddenTopCurvePlane2Shape", tubevase3ddiagram.IsHiddenTopCurvePlane2Shape)
	case "IsHiddenBottomCurvePlane2Shape":
		res = __gong__marshallBool(ident, "IsHiddenBottomCurvePlane2Shape", tubevase3ddiagram.IsHiddenBottomCurvePlane2Shape)
	case "IsHiddenCarvedOutTopCurvePlane1Shape":
		res = __gong__marshallBool(ident, "IsHiddenCarvedOutTopCurvePlane1Shape", tubevase3ddiagram.IsHiddenCarvedOutTopCurvePlane1Shape)
	case "IsHiddenCarvedOutBottomCurvePlane1Shape":
		res = __gong__marshallBool(ident, "IsHiddenCarvedOutBottomCurvePlane1Shape", tubevase3ddiagram.IsHiddenCarvedOutBottomCurvePlane1Shape)
	case "IsHiddenVaseTrapezeRingShape":
		res = __gong__marshallBool(ident, "IsHiddenVaseTrapezeRingShape", tubevase3ddiagram.IsHiddenVaseTrapezeRingShape)
	case "IsHiddenCarvedOutVaseTrapezeRingShape":
		res = __gong__marshallBool(ident, "IsHiddenCarvedOutVaseTrapezeRingShape", tubevase3ddiagram.IsHiddenCarvedOutVaseTrapezeRingShape)
	case "IsHiddenStackOfRotatedCarvedOutVaseTrapezeRingsShape":
		res = __gong__marshallBool(ident, "IsHiddenStackOfRotatedCarvedOutVaseTrapezeRingsShape", tubevase3ddiagram.IsHiddenStackOfRotatedCarvedOutVaseTrapezeRingsShape)
	case "IsHiddenVaseTrapezeBasePlateShape":
		res = __gong__marshallBool(ident, "IsHiddenVaseTrapezeBasePlateShape", tubevase3ddiagram.IsHiddenVaseTrapezeBasePlateShape)
	case "IsChecked":
		res = __gong__marshallBool(ident, "IsChecked", tubevase3ddiagram.IsChecked)
	case "ComputedPrefix":
		res = __gong__marshallString(ident, "ComputedPrefix", tubevase3ddiagram.ComputedPrefix)
	case "IsExpanded":
		res = __gong__marshallBool(ident, "IsExpanded", tubevase3ddiagram.IsExpanded)

	case "Rendered3DShape":
		if tubevase3ddiagram.Rendered3DShape != nil {
			res = __gong__marshallPointer(ident, "Rendered3DShape", tubevase3ddiagram.Rendered3DShape.GongGetIdentifier(stage))
		} else {
			res = __gong__marshallPointer(ident, "Rendered3DShape", "nil")
		}
	case "SampledPoints3DShape":
		if tubevase3ddiagram.SampledPoints3DShape != nil {
			res = __gong__marshallPointer(ident, "SampledPoints3DShape", tubevase3ddiagram.SampledPoints3DShape.GongGetIdentifier(stage))
		} else {
			res = __gong__marshallPointer(ident, "SampledPoints3DShape", "nil")
		}
	case "OriginalPoints3DShape":
		if tubevase3ddiagram.OriginalPoints3DShape != nil {
			res = __gong__marshallPointer(ident, "OriginalPoints3DShape", tubevase3ddiagram.OriginalPoints3DShape.GongGetIdentifier(stage))
		} else {
			res = __gong__marshallPointer(ident, "OriginalPoints3DShape", "nil")
		}
	case "Angle0Shape":
		if tubevase3ddiagram.Angle0Shape != nil {
			res = __gong__marshallPointer(ident, "Angle0Shape", tubevase3ddiagram.Angle0Shape.GongGetIdentifier(stage))
		} else {
			res = __gong__marshallPointer(ident, "Angle0Shape", "nil")
		}
	case "TopCurvePlane1Shape":
		if tubevase3ddiagram.TopCurvePlane1Shape != nil {
			res = __gong__marshallPointer(ident, "TopCurvePlane1Shape", tubevase3ddiagram.TopCurvePlane1Shape.GongGetIdentifier(stage))
		} else {
			res = __gong__marshallPointer(ident, "TopCurvePlane1Shape", "nil")
		}
	case "BottomCurvePlane1Shape":
		if tubevase3ddiagram.BottomCurvePlane1Shape != nil {
			res = __gong__marshallPointer(ident, "BottomCurvePlane1Shape", tubevase3ddiagram.BottomCurvePlane1Shape.GongGetIdentifier(stage))
		} else {
			res = __gong__marshallPointer(ident, "BottomCurvePlane1Shape", "nil")
		}
	case "TopCurvePlane2Shape":
		if tubevase3ddiagram.TopCurvePlane2Shape != nil {
			res = __gong__marshallPointer(ident, "TopCurvePlane2Shape", tubevase3ddiagram.TopCurvePlane2Shape.GongGetIdentifier(stage))
		} else {
			res = __gong__marshallPointer(ident, "TopCurvePlane2Shape", "nil")
		}
	case "BottomCurvePlane2Shape":
		if tubevase3ddiagram.BottomCurvePlane2Shape != nil {
			res = __gong__marshallPointer(ident, "BottomCurvePlane2Shape", tubevase3ddiagram.BottomCurvePlane2Shape.GongGetIdentifier(stage))
		} else {
			res = __gong__marshallPointer(ident, "BottomCurvePlane2Shape", "nil")
		}
	case "CarvedOutTopCurvePlane1Shape":
		if tubevase3ddiagram.CarvedOutTopCurvePlane1Shape != nil {
			res = __gong__marshallPointer(ident, "CarvedOutTopCurvePlane1Shape", tubevase3ddiagram.CarvedOutTopCurvePlane1Shape.GongGetIdentifier(stage))
		} else {
			res = __gong__marshallPointer(ident, "CarvedOutTopCurvePlane1Shape", "nil")
		}
	case "CarvedOutBottomCurvePlane1Shape":
		if tubevase3ddiagram.CarvedOutBottomCurvePlane1Shape != nil {
			res = __gong__marshallPointer(ident, "CarvedOutBottomCurvePlane1Shape", tubevase3ddiagram.CarvedOutBottomCurvePlane1Shape.GongGetIdentifier(stage))
		} else {
			res = __gong__marshallPointer(ident, "CarvedOutBottomCurvePlane1Shape", "nil")
		}
	case "VaseTrapezeRingShape":
		if tubevase3ddiagram.VaseTrapezeRingShape != nil {
			res = __gong__marshallPointer(ident, "VaseTrapezeRingShape", tubevase3ddiagram.VaseTrapezeRingShape.GongGetIdentifier(stage))
		} else {
			res = __gong__marshallPointer(ident, "VaseTrapezeRingShape", "nil")
		}
	case "CarvedOutVaseTrapezeRingShape":
		if tubevase3ddiagram.CarvedOutVaseTrapezeRingShape != nil {
			res = __gong__marshallPointer(ident, "CarvedOutVaseTrapezeRingShape", tubevase3ddiagram.CarvedOutVaseTrapezeRingShape.GongGetIdentifier(stage))
		} else {
			res = __gong__marshallPointer(ident, "CarvedOutVaseTrapezeRingShape", "nil")
		}
	case "StackOfRotatedCarvedOutVaseTrapezeRingsShape":
		if tubevase3ddiagram.StackOfRotatedCarvedOutVaseTrapezeRingsShape != nil {
			res = __gong__marshallPointer(ident, "StackOfRotatedCarvedOutVaseTrapezeRingsShape", tubevase3ddiagram.StackOfRotatedCarvedOutVaseTrapezeRingsShape.GongGetIdentifier(stage))
		} else {
			res = __gong__marshallPointer(ident, "StackOfRotatedCarvedOutVaseTrapezeRingsShape", "nil")
		}
	case "VaseTrapezeBasePlateShape":
		if tubevase3ddiagram.VaseTrapezeBasePlateShape != nil {
			res = __gong__marshallPointer(ident, "VaseTrapezeBasePlateShape", tubevase3ddiagram.VaseTrapezeBasePlateShape.GongGetIdentifier(stage))
		} else {
			res = __gong__marshallPointer(ident, "VaseTrapezeBasePlateShape", "nil")
		}
	default:
		log.Panicf("Unknown field %s for Gongstruct TubeVase3DDiagram", fieldName)
	}
	return
}

func (tubevaseabstract *TubeVaseAbstract) GongMarshallField(stage *Stage, fieldName string) (res string) {
	ident := tubevaseabstract.GongGetIdentifier(stage)
	_ = ident
	switch fieldName {
	case "Name":
		res = __gong__marshallString(ident, "Name", tubevaseabstract.Name)
	case "Z_Ribbon":
		res = __gong__marshallFloat(ident, "Z_Ribbon", tubevaseabstract.Z_Ribbon)
	case "RibbonVerticalScale":
		res = __gong__marshallFloat(ident, "RibbonVerticalScale", tubevaseabstract.RibbonVerticalScale)
	case "TopPlaneHeight":
		res = __gong__marshallFloat(ident, "TopPlaneHeight", tubevaseabstract.TopPlaneHeight)
	case "BottomPlaneHeight":
		res = __gong__marshallFloat(ident, "BottomPlaneHeight", tubevaseabstract.BottomPlaneHeight)
	case "BaseRingBottomPlaneHeight":
		res = __gong__marshallFloat(ident, "BaseRingBottomPlaneHeight", tubevaseabstract.BaseRingBottomPlaneHeight)
	case "ProjectionAngle":
		res = __gong__marshallFloat(ident, "ProjectionAngle", tubevaseabstract.ProjectionAngle)
	case "BasePlateHeight":
		res = __gong__marshallFloat(ident, "BasePlateHeight", tubevaseabstract.BasePlateHeight)
	case "CarvedOutTopRingsParameter":
		res = __gong__marshallFloat(ident, "CarvedOutTopRingsParameter", tubevaseabstract.CarvedOutTopRingsParameter)
	case "BulbousStartTangentMagnitude":
		res = __gong__marshallFloat(ident, "BulbousStartTangentMagnitude", tubevaseabstract.BulbousStartTangentMagnitude)
	case "BulbousEndAngle":
		res = __gong__marshallFloat(ident, "BulbousEndAngle", tubevaseabstract.BulbousEndAngle)
	case "BulbousEndTangentMagnitude":
		res = __gong__marshallFloat(ident, "BulbousEndTangentMagnitude", tubevaseabstract.BulbousEndTangentMagnitude)
	case "RelativeVerticalThickness":
		res = __gong__marshallFloat(ident, "RelativeVerticalThickness", tubevaseabstract.RelativeVerticalThickness)
	case "RelativeRadialThickness":
		res = __gong__marshallFloat(ident, "RelativeRadialThickness", tubevaseabstract.RelativeRadialThickness)
	case "RelativeCuttedStackFloorHeight":
		res = __gong__marshallFloat(ident, "RelativeCuttedStackFloorHeight", tubevaseabstract.RelativeCuttedStackFloorHeight)
	case "RelativeRotatedTorusSeparation":
		res = __gong__marshallFloat(ident, "RelativeRotatedTorusSeparation", tubevaseabstract.RelativeRotatedTorusSeparation)
	case "RotationRatio":
		res = __gong__marshallFloat(ident, "RotationRatio", tubevaseabstract.RotationRatio)
	case "RadialRepetitions":
		res = __gong__marshallInt(ident, "RadialRepetitions", tubevaseabstract.RadialRepetitions)
	case "Transparency":
		res = __gong__marshallFloat(ident, "Transparency", tubevaseabstract.Transparency)
	case "HasAlternatingRingColors":
		res = __gong__marshallBool(ident, "HasAlternatingRingColors", tubevaseabstract.HasAlternatingRingColors)
	case "RelativeTrajectoryOffsetX":
		res = __gong__marshallFloat(ident, "RelativeTrajectoryOffsetX", tubevaseabstract.RelativeTrajectoryOffsetX)
	case "RelativeTrajectoryOffsetY":
		res = __gong__marshallFloat(ident, "RelativeTrajectoryOffsetY", tubevaseabstract.RelativeTrajectoryOffsetY)
	case "NbStepP1P2":
		res = __gong__marshallInt(ident, "NbStepP1P2", tubevaseabstract.NbStepP1P2)
	case "ChosenStep":
		res = __gong__marshallInt(ident, "ChosenStep", tubevaseabstract.ChosenStep)
	case "RelativeHorizontalRingsHeight":
		res = __gong__marshallFloat(ident, "RelativeHorizontalRingsHeight", tubevaseabstract.RelativeHorizontalRingsHeight)
	case "OffsetKeyX":
		res = __gong__marshallFloat(ident, "OffsetKeyX", tubevaseabstract.OffsetKeyX)
	case "OffsetKeyY":
		res = __gong__marshallFloat(ident, "OffsetKeyY", tubevaseabstract.OffsetKeyY)
	case "HeightKey":
		res = __gong__marshallFloat(ident, "HeightKey", tubevaseabstract.HeightKey)
	case "WidthKey":
		res = __gong__marshallFloat(ident, "WidthKey", tubevaseabstract.WidthKey)
	case "RelativeKeySize":
		res = __gong__marshallFloat(ident, "RelativeKeySize", tubevaseabstract.RelativeKeySize)
	case "MovieNbFrames":
		res = __gong__marshallInt(ident, "MovieNbFrames", tubevaseabstract.MovieNbFrames)

	default:
		log.Panicf("Unknown field %s for Gongstruct TubeVaseAbstract", fieldName)
	}
	return
}

func (vase2ddiagram *Vase2DDiagram) GongMarshallField(stage *Stage, fieldName string) (res string) {
	ident := vase2ddiagram.GongGetIdentifier(stage)
	_ = ident
	switch fieldName {
	case "Name":
		res = __gong__marshallString(ident, "Name", vase2ddiagram.Name)
	case "Zoom":
		res = __gong__marshallFloat(ident, "Zoom", vase2ddiagram.Zoom)
	case "IsVaseArcNodesExpanded":
		res = __gong__marshallBool(ident, "IsVaseArcNodesExpanded", vase2ddiagram.IsVaseArcNodesExpanded)
	case "IsVaseClampingNodesExpanded":
		res = __gong__marshallBool(ident, "IsVaseClampingNodesExpanded", vase2ddiagram.IsVaseClampingNodesExpanded)
	case "IsHiddenAxesShape":
		res = __gong__marshallBool(ident, "IsHiddenAxesShape", vase2ddiagram.IsHiddenAxesShape)
	case "IsHiddenBottomStartArcShapeGrid":
		res = __gong__marshallBool(ident, "IsHiddenBottomStartArcShapeGrid", vase2ddiagram.IsHiddenBottomStartArcShapeGrid)
	case "IsHiddenBottomEndArcShapeGrid":
		res = __gong__marshallBool(ident, "IsHiddenBottomEndArcShapeGrid", vase2ddiagram.IsHiddenBottomEndArcShapeGrid)
	case "IsHiddenBottomStackOfGrowthCurve":
		res = __gong__marshallBool(ident, "IsHiddenBottomStackOfGrowthCurve", vase2ddiagram.IsHiddenBottomStackOfGrowthCurve)
	case "IsHiddenShiftedLeftStackOfGrowthCurve":
		res = __gong__marshallBool(ident, "IsHiddenShiftedLeftStackOfGrowthCurve", vase2ddiagram.IsHiddenShiftedLeftStackOfGrowthCurve)
	case "IsHiddenShiftedLeftStackOfNormalVector":
		res = __gong__marshallBool(ident, "IsHiddenShiftedLeftStackOfNormalVector", vase2ddiagram.IsHiddenShiftedLeftStackOfNormalVector)
	case "IsHiddenPerpendicularVectorGridHalfway":
		res = __gong__marshallBool(ident, "IsHiddenPerpendicularVectorGridHalfway", vase2ddiagram.IsHiddenPerpendicularVectorGridHalfway)
	case "IsHiddenTopStartArcShapeGrid":
		res = __gong__marshallBool(ident, "IsHiddenTopStartArcShapeGrid", vase2ddiagram.IsHiddenTopStartArcShapeGrid)
	case "IsHiddenShiftedBottomTopStartArcShapeGrid":
		res = __gong__marshallBool(ident, "IsHiddenShiftedBottomTopStartArcShapeGrid", vase2ddiagram.IsHiddenShiftedBottomTopStartArcShapeGrid)
	case "IsHiddenTopMidArcVectorShapeGrid":
		res = __gong__marshallBool(ident, "IsHiddenTopMidArcVectorShapeGrid", vase2ddiagram.IsHiddenTopMidArcVectorShapeGrid)
	case "IsHiddenStartHalfwayArcShapeGrid":
		res = __gong__marshallBool(ident, "IsHiddenStartHalfwayArcShapeGrid", vase2ddiagram.IsHiddenStartHalfwayArcShapeGrid)
	case "IsHiddenTopStartHalfwayArcShapeGrid":
		res = __gong__marshallBool(ident, "IsHiddenTopStartHalfwayArcShapeGrid", vase2ddiagram.IsHiddenTopStartHalfwayArcShapeGrid)
	case "IsHiddenEndHalfwayArcShapeGrid":
		res = __gong__marshallBool(ident, "IsHiddenEndHalfwayArcShapeGrid", vase2ddiagram.IsHiddenEndHalfwayArcShapeGrid)
	case "IsHiddenTopEndHalfwayArcShapeGrid":
		res = __gong__marshallBool(ident, "IsHiddenTopEndHalfwayArcShapeGrid", vase2ddiagram.IsHiddenTopEndHalfwayArcShapeGrid)
	case "IsHiddenTopEndArcShapeGrid":
		res = __gong__marshallBool(ident, "IsHiddenTopEndArcShapeGrid", vase2ddiagram.IsHiddenTopEndArcShapeGrid)
	case "IsHiddenStackOfGrowthCurve":
		res = __gong__marshallBool(ident, "IsHiddenStackOfGrowthCurve", vase2ddiagram.IsHiddenStackOfGrowthCurve)
	case "IsHiddenTopStackOfGrowthCurve":
		res = __gong__marshallBool(ident, "IsHiddenTopStackOfGrowthCurve", vase2ddiagram.IsHiddenTopStackOfGrowthCurve)
	case "IsHiddenTopGrowthCurve2D":
		res = __gong__marshallBool(ident, "IsHiddenTopGrowthCurve2D", vase2ddiagram.IsHiddenTopGrowthCurve2D)
	case "IsHiddenStackOfGrowthCurve2D":
		res = __gong__marshallBool(ident, "IsHiddenStackOfGrowthCurve2D", vase2ddiagram.IsHiddenStackOfGrowthCurve2D)
	case "IsHiddenTopStackOfGrowthCurve2D":
		res = __gong__marshallBool(ident, "IsHiddenTopStackOfGrowthCurve2D", vase2ddiagram.IsHiddenTopStackOfGrowthCurve2D)
	case "IsHiddenGrowthCurve2DRibbon":
		res = __gong__marshallBool(ident, "IsHiddenGrowthCurve2DRibbon", vase2ddiagram.IsHiddenGrowthCurve2DRibbon)
	case "IsHiddenShiftedRightGrowthCurve2DRibbon":
		res = __gong__marshallBool(ident, "IsHiddenShiftedRightGrowthCurve2DRibbon", vase2ddiagram.IsHiddenShiftedRightGrowthCurve2DRibbon)
	case "IsHiddenShiftedLeftGrowthCurve2DRibbon":
		res = __gong__marshallBool(ident, "IsHiddenShiftedLeftGrowthCurve2DRibbon", vase2ddiagram.IsHiddenShiftedLeftGrowthCurve2DRibbon)
	case "IsHiddenStackOfGrowthCurve2DRibbon":
		res = __gong__marshallBool(ident, "IsHiddenStackOfGrowthCurve2DRibbon", vase2ddiagram.IsHiddenStackOfGrowthCurve2DRibbon)
	case "IsHiddenStackOfRotatedGrowthCurve2DRibbon":
		res = __gong__marshallBool(ident, "IsHiddenStackOfRotatedGrowthCurve2DRibbon", vase2ddiagram.IsHiddenStackOfRotatedGrowthCurve2DRibbon)
	case "IsHiddenPartiallyGrowthCurve2DRibbon":
		res = __gong__marshallBool(ident, "IsHiddenPartiallyGrowthCurve2DRibbon", vase2ddiagram.IsHiddenPartiallyGrowthCurve2DRibbon)
	case "IsHiddenShiftedLeftPartiallyGrowthCurve2DRibbon":
		res = __gong__marshallBool(ident, "IsHiddenShiftedLeftPartiallyGrowthCurve2DRibbon", vase2ddiagram.IsHiddenShiftedLeftPartiallyGrowthCurve2DRibbon)
	case "IsHiddenPartiallyGrowthCurve2DTrajectory":
		res = __gong__marshallBool(ident, "IsHiddenPartiallyGrowthCurve2DTrajectory", vase2ddiagram.IsHiddenPartiallyGrowthCurve2DTrajectory)
	case "IsHiddenPartiallyGrowthCurve2DTrajectoryP1P2":
		res = __gong__marshallBool(ident, "IsHiddenPartiallyGrowthCurve2DTrajectoryP1P2", vase2ddiagram.IsHiddenPartiallyGrowthCurve2DTrajectoryP1P2)
	case "IsHiddenPxShape":
		res = __gong__marshallBool(ident, "IsHiddenPxShape", vase2ddiagram.IsHiddenPxShape)
	case "IsHiddenChosenP1P2PairShape":
		res = __gong__marshallBool(ident, "IsHiddenChosenP1P2PairShape", vase2ddiagram.IsHiddenChosenP1P2PairShape)
	case "IsHiddenKeyHoleShape":
		res = __gong__marshallBool(ident, "IsHiddenKeyHoleShape", vase2ddiagram.IsHiddenKeyHoleShape)
	case "IsChecked":
		res = __gong__marshallBool(ident, "IsChecked", vase2ddiagram.IsChecked)
	case "ComputedPrefix":
		res = __gong__marshallString(ident, "ComputedPrefix", vase2ddiagram.ComputedPrefix)
	case "IsExpanded":
		res = __gong__marshallBool(ident, "IsExpanded", vase2ddiagram.IsExpanded)

	default:
		log.Panicf("Unknown field %s for Gongstruct Vase2DDiagram", fieldName)
	}
	return
}

func (vasetrapezebaseplateshape *VaseTrapezeBasePlateShape) GongMarshallField(stage *Stage, fieldName string) (res string) {
	ident := vasetrapezebaseplateshape.GongGetIdentifier(stage)
	_ = ident
	switch fieldName {
	case "Name":
		res = __gong__marshallString(ident, "Name", vasetrapezebaseplateshape.Name)

	default:
		log.Panicf("Unknown field %s for Gongstruct VaseTrapezeBasePlateShape", fieldName)
	}
	return
}

func (vasetrapezeringshape *VaseTrapezeRingShape) GongMarshallField(stage *Stage, fieldName string) (res string) {
	ident := vasetrapezeringshape.GongGetIdentifier(stage)
	_ = ident
	switch fieldName {
	case "Name":
		res = __gong__marshallString(ident, "Name", vasetrapezeringshape.Name)

	default:
		log.Panicf("Unknown field %s for Gongstruct VaseTrapezeRingShape", fieldName)
	}
	return
}

func (verticaltorusstackshape *VerticalTorusStackShape) GongMarshallField(stage *Stage, fieldName string) (res string) {
	ident := verticaltorusstackshape.GongGetIdentifier(stage)
	_ = ident
	switch fieldName {
	case "Name":
		res = __gong__marshallString(ident, "Name", verticaltorusstackshape.Name)

	default:
		log.Panicf("Unknown field %s for Gongstruct VerticalTorusStackShape", fieldName)
	}
	return
}

func (volumekey3dshape *VolumeKey3DShape) GongMarshallField(stage *Stage, fieldName string) (res string) {
	ident := volumekey3dshape.GongGetIdentifier(stage)
	_ = ident
	switch fieldName {
	case "Name":
		res = __gong__marshallString(ident, "Name", volumekey3dshape.Name)

	default:
		log.Panicf("Unknown field %s for Gongstruct VolumeKey3DShape", fieldName)
	}
	return
}

// insertion point for marshall all fields methods
func (angle0shape *Angle0Shape) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(angle0shape.GongMarshallField(stage, "Name"))
	}
	initRes = initializerStatements.String()
	ptrRes = pointersInitializesStatements.String()
	return
}
func (arcnormalvectorshape *ArcNormalVectorShape) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(arcnormalvectorshape.GongMarshallField(stage, "Name"))
		initializerStatements.WriteString(arcnormalvectorshape.GongMarshallField(stage, "StartX"))
		initializerStatements.WriteString(arcnormalvectorshape.GongMarshallField(stage, "StartY"))
		initializerStatements.WriteString(arcnormalvectorshape.GongMarshallField(stage, "EndX"))
		initializerStatements.WriteString(arcnormalvectorshape.GongMarshallField(stage, "EndY"))
	}
	initRes = initializerStatements.String()
	ptrRes = pointersInitializesStatements.String()
	return
}
func (arcnormalvectorshapegrid *ArcNormalVectorShapeGrid) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(arcnormalvectorshapegrid.GongMarshallField(stage, "Name"))
	}
	initRes = initializerStatements.String()
	ptrRes = pointersInitializesStatements.String()
	return
}
func (axesshape *AxesShape) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(axesshape.GongMarshallField(stage, "Name"))
		initializerStatements.WriteString(axesshape.GongMarshallField(stage, "LengthX"))
		initializerStatements.WriteString(axesshape.GongMarshallField(stage, "LengthY"))
		initializerStatements.WriteString(axesshape.GongMarshallField(stage, "IsWithHiddenHandle"))
	}
	initRes = initializerStatements.String()
	ptrRes = pointersInitializesStatements.String()
	return
}
func (basevectorshape *BaseVectorShape) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(basevectorshape.GongMarshallField(stage, "Name"))
		initializerStatements.WriteString(basevectorshape.GongMarshallField(stage, "StartX"))
		initializerStatements.WriteString(basevectorshape.GongMarshallField(stage, "StartY"))
		initializerStatements.WriteString(basevectorshape.GongMarshallField(stage, "EndX"))
		initializerStatements.WriteString(basevectorshape.GongMarshallField(stage, "EndY"))
	}
	initRes = initializerStatements.String()
	ptrRes = pointersInitializesStatements.String()
	return
}
func (basevectorshapegrid *BaseVectorShapeGrid) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(basevectorshapegrid.GongMarshallField(stage, "Name"))
	}
	initRes = initializerStatements.String()
	ptrRes = pointersInitializesStatements.String()
	return
}
func (bottomcurveplane1shape *BottomCurvePlane1Shape) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(bottomcurveplane1shape.GongMarshallField(stage, "Name"))
	}
	initRes = initializerStatements.String()
	ptrRes = pointersInitializesStatements.String()
	return
}
func (bottomcurveplane2shape *BottomCurvePlane2Shape) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(bottomcurveplane2shape.GongMarshallField(stage, "Name"))
	}
	initRes = initializerStatements.String()
	ptrRes = pointersInitializesStatements.String()
	return
}
func (carvedoutbottomcurveplane1shape *CarvedOutBottomCurvePlane1Shape) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(carvedoutbottomcurveplane1shape.GongMarshallField(stage, "Name"))
	}
	initRes = initializerStatements.String()
	ptrRes = pointersInitializesStatements.String()
	return
}
func (carvedouttopcurveplane1shape *CarvedOutTopCurvePlane1Shape) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(carvedouttopcurveplane1shape.GongMarshallField(stage, "Name"))
	}
	initRes = initializerStatements.String()
	ptrRes = pointersInitializesStatements.String()
	return
}
func (carvedoutvasetrapezeringshape *CarvedOutVaseTrapezeRingShape) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(carvedoutvasetrapezeringshape.GongMarshallField(stage, "Name"))
	}
	initRes = initializerStatements.String()
	ptrRes = pointersInitializesStatements.String()
	return
}
func (chosenp1p2pairshape *ChosenP1P2PairShape) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(chosenp1p2pairshape.GongMarshallField(stage, "Name"))
		initializerStatements.WriteString(chosenp1p2pairshape.GongMarshallField(stage, "P1X"))
		initializerStatements.WriteString(chosenp1p2pairshape.GongMarshallField(stage, "P1Y"))
		initializerStatements.WriteString(chosenp1p2pairshape.GongMarshallField(stage, "P2X"))
		initializerStatements.WriteString(chosenp1p2pairshape.GongMarshallField(stage, "P2Y"))
		initializerStatements.WriteString(chosenp1p2pairshape.GongMarshallField(stage, "PxX"))
		initializerStatements.WriteString(chosenp1p2pairshape.GongMarshallField(stage, "PxY"))
		initializerStatements.WriteString(chosenp1p2pairshape.GongMarshallField(stage, "DistanceP1Px"))
		initializerStatements.WriteString(chosenp1p2pairshape.GongMarshallField(stage, "DistanceP2Px"))
		initializerStatements.WriteString(chosenp1p2pairshape.GongMarshallField(stage, "DistanceSum"))
	}
	initRes = initializerStatements.String()
	ptrRes = pointersInitializesStatements.String()
	return
}
func (circlegridshape *CircleGridShape) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(circlegridshape.GongMarshallField(stage, "Name"))
	}
	initRes = initializerStatements.String()
	ptrRes = pointersInitializesStatements.String()
	return
}
func (circumference3dshape *Circumference3DShape) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(circumference3dshape.GongMarshallField(stage, "Name"))
	}
	initRes = initializerStatements.String()
	ptrRes = pointersInitializesStatements.String()
	return
}
func (clock2ddiagram *Clock2DDiagram) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(clock2ddiagram.GongMarshallField(stage, "Name"))
		initializerStatements.WriteString(clock2ddiagram.GongMarshallField(stage, "Zoom"))
		initializerStatements.WriteString(clock2ddiagram.GongMarshallField(stage, "IsHiddenAxesShape"))
		initializerStatements.WriteString(clock2ddiagram.GongMarshallField(stage, "IsChecked"))
		initializerStatements.WriteString(clock2ddiagram.GongMarshallField(stage, "ComputedPrefix"))
		initializerStatements.WriteString(clock2ddiagram.GongMarshallField(stage, "IsExpanded"))
	}
	initRes = initializerStatements.String()
	ptrRes = pointersInitializesStatements.String()
	return
}
func (clock3ddiagram *Clock3DDiagram) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(clock3ddiagram.GongMarshallField(stage, "Name"))
		initializerStatements.WriteString(clock3ddiagram.GongMarshallField(stage, "IsHiddenClockTopCurveShape"))
		initializerStatements.WriteString(clock3ddiagram.GongMarshallField(stage, "IsHiddenTorus3DShape"))
		initializerStatements.WriteString(clock3ddiagram.GongMarshallField(stage, "IsHiddenSampledPoints3DShape"))
		pointersInitializesStatements.WriteString(clock3ddiagram.GongMarshallField(stage, "SampledPoints3DShape"))
		initializerStatements.WriteString(clock3ddiagram.GongMarshallField(stage, "IsHiddenTiledFloor3DShape"))
		pointersInitializesStatements.WriteString(clock3ddiagram.GongMarshallField(stage, "Rendered3DShape"))
		initializerStatements.WriteString(clock3ddiagram.GongMarshallField(stage, "IsChecked"))
		initializerStatements.WriteString(clock3ddiagram.GongMarshallField(stage, "ComputedPrefix"))
		initializerStatements.WriteString(clock3ddiagram.GongMarshallField(stage, "IsExpanded"))
	}
	initRes = initializerStatements.String()
	ptrRes = pointersInitializesStatements.String()
	return
}
func (clocktopcurveshape *ClockTopCurveShape) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(clocktopcurveshape.GongMarshallField(stage, "Name"))
	}
	initRes = initializerStatements.String()
	ptrRes = pointersInitializesStatements.String()
	return
}
func (cutline3dshape *CutLine3DShape) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(cutline3dshape.GongMarshallField(stage, "Name"))
	}
	initRes = initializerStatements.String()
	ptrRes = pointersInitializesStatements.String()
	return
}
func (endarcshape *EndArcShape) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(endarcshape.GongMarshallField(stage, "Name"))
		initializerStatements.WriteString(endarcshape.GongMarshallField(stage, "StartX"))
		initializerStatements.WriteString(endarcshape.GongMarshallField(stage, "StartY"))
		initializerStatements.WriteString(endarcshape.GongMarshallField(stage, "EndX"))
		initializerStatements.WriteString(endarcshape.GongMarshallField(stage, "EndY"))
		initializerStatements.WriteString(endarcshape.GongMarshallField(stage, "XAxisRotation"))
		initializerStatements.WriteString(endarcshape.GongMarshallField(stage, "LargeArcFlag"))
		initializerStatements.WriteString(endarcshape.GongMarshallField(stage, "SweepFlag"))
		initializerStatements.WriteString(endarcshape.GongMarshallField(stage, "RadiusX"))
		initializerStatements.WriteString(endarcshape.GongMarshallField(stage, "RadiusY"))
	}
	initRes = initializerStatements.String()
	ptrRes = pointersInitializesStatements.String()
	return
}
func (endarcshapegrid *EndArcShapeGrid) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(endarcshapegrid.GongMarshallField(stage, "Name"))
	}
	initRes = initializerStatements.String()
	ptrRes = pointersInitializesStatements.String()
	return
}
func (endhalfwayarcshape *EndHalfwayArcShape) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(endhalfwayarcshape.GongMarshallField(stage, "Name"))
		initializerStatements.WriteString(endhalfwayarcshape.GongMarshallField(stage, "StartX"))
		initializerStatements.WriteString(endhalfwayarcshape.GongMarshallField(stage, "StartY"))
		initializerStatements.WriteString(endhalfwayarcshape.GongMarshallField(stage, "EndX"))
		initializerStatements.WriteString(endhalfwayarcshape.GongMarshallField(stage, "EndY"))
		initializerStatements.WriteString(endhalfwayarcshape.GongMarshallField(stage, "RadiusX"))
		initializerStatements.WriteString(endhalfwayarcshape.GongMarshallField(stage, "RadiusY"))
		initializerStatements.WriteString(endhalfwayarcshape.GongMarshallField(stage, "XAxisRotation"))
		initializerStatements.WriteString(endhalfwayarcshape.GongMarshallField(stage, "LargeArcFlag"))
		initializerStatements.WriteString(endhalfwayarcshape.GongMarshallField(stage, "SweepFlag"))
	}
	initRes = initializerStatements.String()
	ptrRes = pointersInitializesStatements.String()
	return
}
func (endhalfwayarcshapegrid *EndHalfwayArcShapeGrid) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(endhalfwayarcshapegrid.GongMarshallField(stage, "Name"))
	}
	initRes = initializerStatements.String()
	ptrRes = pointersInitializesStatements.String()
	return
}
func (explanationtextshape *ExplanationTextShape) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(explanationtextshape.GongMarshallField(stage, "Name"))
	}
	initRes = initializerStatements.String()
	ptrRes = pointersInitializesStatements.String()
	return
}
func (eye3dshape *Eye3DShape) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(eye3dshape.GongMarshallField(stage, "Name"))
	}
	initRes = initializerStatements.String()
	ptrRes = pointersInitializesStatements.String()
	return
}
func (eyecornerssampledpoints3dshape *EyeCornersSampledPoints3DShape) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(eyecornerssampledpoints3dshape.GongMarshallField(stage, "Name"))
	}
	initRes = initializerStatements.String()
	ptrRes = pointersInitializesStatements.String()
	return
}
func (eyesampledpoints3dshape *EyeSampledPoints3DShape) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(eyesampledpoints3dshape.GongMarshallField(stage, "Name"))
	}
	initRes = initializerStatements.String()
	ptrRes = pointersInitializesStatements.String()
	return
}
func (eyeseatbottomcurveshape *EyeSeatBottomCurveShape) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(eyeseatbottomcurveshape.GongMarshallField(stage, "Name"))
	}
	initRes = initializerStatements.String()
	ptrRes = pointersInitializesStatements.String()
	return
}
func (eyestoolbottomcurveshape *EyeStoolBottomCurveShape) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(eyestoolbottomcurveshape.GongMarshallField(stage, "Name"))
	}
	initRes = initializerStatements.String()
	ptrRes = pointersInitializesStatements.String()
	return
}
func (eyevolume3dshape *EyeVolume3DShape) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(eyevolume3dshape.GongMarshallField(stage, "Name"))
	}
	initRes = initializerStatements.String()
	ptrRes = pointersInitializesStatements.String()
	return
}
func (gridpathshape *GridPathShape) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(gridpathshape.GongMarshallField(stage, "Name"))
	}
	initRes = initializerStatements.String()
	ptrRes = pointersInitializesStatements.String()
	return
}
func (growthcurve2d *GrowthCurve2D) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(growthcurve2d.GongMarshallField(stage, "Name"))
	}
	initRes = initializerStatements.String()
	ptrRes = pointersInitializesStatements.String()
	return
}
func (growthcurve2dribbon *GrowthCurve2DRibbon) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(growthcurve2dribbon.GongMarshallField(stage, "Name"))
	}
	initRes = initializerStatements.String()
	ptrRes = pointersInitializesStatements.String()
	return
}
func (growthcurve2dribbonendshape *GrowthCurve2DRibbonEndShape) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(growthcurve2dribbonendshape.GongMarshallField(stage, "Name"))
		initializerStatements.WriteString(growthcurve2dribbonendshape.GongMarshallField(stage, "BottomStartX"))
		initializerStatements.WriteString(growthcurve2dribbonendshape.GongMarshallField(stage, "BottomStartY"))
		initializerStatements.WriteString(growthcurve2dribbonendshape.GongMarshallField(stage, "BottomEndX"))
		initializerStatements.WriteString(growthcurve2dribbonendshape.GongMarshallField(stage, "BottomEndY"))
		initializerStatements.WriteString(growthcurve2dribbonendshape.GongMarshallField(stage, "BottomRadiusX"))
		initializerStatements.WriteString(growthcurve2dribbonendshape.GongMarshallField(stage, "BottomRadiusY"))
		initializerStatements.WriteString(growthcurve2dribbonendshape.GongMarshallField(stage, "BottomXAxisRotation"))
		initializerStatements.WriteString(growthcurve2dribbonendshape.GongMarshallField(stage, "BottomLargeArcFlag"))
		initializerStatements.WriteString(growthcurve2dribbonendshape.GongMarshallField(stage, "BottomSweepFlag"))
		initializerStatements.WriteString(growthcurve2dribbonendshape.GongMarshallField(stage, "TopStartX"))
		initializerStatements.WriteString(growthcurve2dribbonendshape.GongMarshallField(stage, "TopStartY"))
		initializerStatements.WriteString(growthcurve2dribbonendshape.GongMarshallField(stage, "TopEndX"))
		initializerStatements.WriteString(growthcurve2dribbonendshape.GongMarshallField(stage, "TopEndY"))
		initializerStatements.WriteString(growthcurve2dribbonendshape.GongMarshallField(stage, "TopRadiusX"))
		initializerStatements.WriteString(growthcurve2dribbonendshape.GongMarshallField(stage, "TopRadiusY"))
		initializerStatements.WriteString(growthcurve2dribbonendshape.GongMarshallField(stage, "TopXAxisRotation"))
		initializerStatements.WriteString(growthcurve2dribbonendshape.GongMarshallField(stage, "TopLargeArcFlag"))
		initializerStatements.WriteString(growthcurve2dribbonendshape.GongMarshallField(stage, "TopSweepFlag"))
	}
	initRes = initializerStatements.String()
	ptrRes = pointersInitializesStatements.String()
	return
}
func (growthcurve2dribbonstartshape *GrowthCurve2DRibbonStartShape) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(growthcurve2dribbonstartshape.GongMarshallField(stage, "Name"))
		initializerStatements.WriteString(growthcurve2dribbonstartshape.GongMarshallField(stage, "BottomStartX"))
		initializerStatements.WriteString(growthcurve2dribbonstartshape.GongMarshallField(stage, "BottomStartY"))
		initializerStatements.WriteString(growthcurve2dribbonstartshape.GongMarshallField(stage, "BottomEndX"))
		initializerStatements.WriteString(growthcurve2dribbonstartshape.GongMarshallField(stage, "BottomEndY"))
		initializerStatements.WriteString(growthcurve2dribbonstartshape.GongMarshallField(stage, "BottomRadiusX"))
		initializerStatements.WriteString(growthcurve2dribbonstartshape.GongMarshallField(stage, "BottomRadiusY"))
		initializerStatements.WriteString(growthcurve2dribbonstartshape.GongMarshallField(stage, "BottomXAxisRotation"))
		initializerStatements.WriteString(growthcurve2dribbonstartshape.GongMarshallField(stage, "BottomLargeArcFlag"))
		initializerStatements.WriteString(growthcurve2dribbonstartshape.GongMarshallField(stage, "BottomSweepFlag"))
		initializerStatements.WriteString(growthcurve2dribbonstartshape.GongMarshallField(stage, "TopStartX"))
		initializerStatements.WriteString(growthcurve2dribbonstartshape.GongMarshallField(stage, "TopStartY"))
		initializerStatements.WriteString(growthcurve2dribbonstartshape.GongMarshallField(stage, "TopEndX"))
		initializerStatements.WriteString(growthcurve2dribbonstartshape.GongMarshallField(stage, "TopEndY"))
		initializerStatements.WriteString(growthcurve2dribbonstartshape.GongMarshallField(stage, "TopRadiusX"))
		initializerStatements.WriteString(growthcurve2dribbonstartshape.GongMarshallField(stage, "TopRadiusY"))
		initializerStatements.WriteString(growthcurve2dribbonstartshape.GongMarshallField(stage, "TopXAxisRotation"))
		initializerStatements.WriteString(growthcurve2dribbonstartshape.GongMarshallField(stage, "TopLargeArcFlag"))
		initializerStatements.WriteString(growthcurve2dribbonstartshape.GongMarshallField(stage, "TopSweepFlag"))
	}
	initRes = initializerStatements.String()
	ptrRes = pointersInitializesStatements.String()
	return
}
func (growthcurverhombusgridshape *GrowthCurveRhombusGridShape) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(growthcurverhombusgridshape.GongMarshallField(stage, "Name"))
	}
	initRes = initializerStatements.String()
	ptrRes = pointersInitializesStatements.String()
	return
}
func (growthcurverhombusshape *GrowthCurveRhombusShape) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(growthcurverhombusshape.GongMarshallField(stage, "Name"))
		initializerStatements.WriteString(growthcurverhombusshape.GongMarshallField(stage, "X"))
		initializerStatements.WriteString(growthcurverhombusshape.GongMarshallField(stage, "Y"))
	}
	initRes = initializerStatements.String()
	ptrRes = pointersInitializesStatements.String()
	return
}
func (growthvectorshape *GrowthVectorShape) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(growthvectorshape.GongMarshallField(stage, "Name"))
		initializerStatements.WriteString(growthvectorshape.GongMarshallField(stage, "X"))
		initializerStatements.WriteString(growthvectorshape.GongMarshallField(stage, "Y"))
	}
	initRes = initializerStatements.String()
	ptrRes = pointersInitializesStatements.String()
	return
}
func (initialrhombusgridshape *InitialRhombusGridShape) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(initialrhombusgridshape.GongMarshallField(stage, "Name"))
	}
	initRes = initializerStatements.String()
	ptrRes = pointersInitializesStatements.String()
	return
}
func (initialrhombusshape *InitialRhombusShape) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(initialrhombusshape.GongMarshallField(stage, "Name"))
		initializerStatements.WriteString(initialrhombusshape.GongMarshallField(stage, "X"))
		initializerStatements.WriteString(initialrhombusshape.GongMarshallField(stage, "Y"))
	}
	initRes = initializerStatements.String()
	ptrRes = pointersInitializesStatements.String()
	return
}
func (key3dshape *Key3DShape) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(key3dshape.GongMarshallField(stage, "Name"))
	}
	initRes = initializerStatements.String()
	ptrRes = pointersInitializesStatements.String()
	return
}
func (keyhole3dshape *KeyHole3DShape) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(keyhole3dshape.GongMarshallField(stage, "Name"))
	}
	initRes = initializerStatements.String()
	ptrRes = pointersInitializesStatements.String()
	return
}
func (keyholeshape *KeyHoleShape) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(keyholeshape.GongMarshallField(stage, "Name"))
		initializerStatements.WriteString(keyholeshape.GongMarshallField(stage, "X"))
		initializerStatements.WriteString(keyholeshape.GongMarshallField(stage, "Y"))
		initializerStatements.WriteString(keyholeshape.GongMarshallField(stage, "Width"))
		initializerStatements.WriteString(keyholeshape.GongMarshallField(stage, "Height"))
	}
	initRes = initializerStatements.String()
	ptrRes = pointersInitializesStatements.String()
	return
}
func (leaves3dshape *Leaves3DShape) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(leaves3dshape.GongMarshallField(stage, "Name"))
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
		pointersInitializesStatements.WriteString(library.GongMarshallField(stage, "Plants"))
		pointersInitializesStatements.WriteString(library.GongMarshallField(stage, "SubLibraries"))
		initializerStatements.WriteString(library.GongMarshallField(stage, "NbPixPerCharacter"))
		initializerStatements.WriteString(library.GongMarshallField(stage, "LogoSVGFile"))
		initializerStatements.WriteString(library.GongMarshallField(stage, "ComputedPrefix"))
		initializerStatements.WriteString(library.GongMarshallField(stage, "IsExpanded"))
		initializerStatements.WriteString(library.GongMarshallField(stage, "IsRootLibrary"))
	}
	initRes = initializerStatements.String()
	ptrRes = pointersInitializesStatements.String()
	return
}
func (midarcvectorshape *MidArcVectorShape) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(midarcvectorshape.GongMarshallField(stage, "Name"))
		initializerStatements.WriteString(midarcvectorshape.GongMarshallField(stage, "StartX"))
		initializerStatements.WriteString(midarcvectorshape.GongMarshallField(stage, "StartY"))
		initializerStatements.WriteString(midarcvectorshape.GongMarshallField(stage, "EndX"))
		initializerStatements.WriteString(midarcvectorshape.GongMarshallField(stage, "EndY"))
	}
	initRes = initializerStatements.String()
	ptrRes = pointersInitializesStatements.String()
	return
}
func (midarcvectorshapegrid *MidArcVectorShapeGrid) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(midarcvectorshapegrid.GongMarshallField(stage, "Name"))
	}
	initRes = initializerStatements.String()
	ptrRes = pointersInitializesStatements.String()
	return
}
func (originalpoints3dshape *OriginalPoints3DShape) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(originalpoints3dshape.GongMarshallField(stage, "Name"))
	}
	initRes = initializerStatements.String()
	ptrRes = pointersInitializesStatements.String()
	return
}
func (parastichymcurves3dshape *ParastichyMCurves3DShape) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(parastichymcurves3dshape.GongMarshallField(stage, "Name"))
	}
	initRes = initializerStatements.String()
	ptrRes = pointersInitializesStatements.String()
	return
}
func (parastichyncurves3dshape *ParastichyNCurves3DShape) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(parastichyncurves3dshape.GongMarshallField(stage, "Name"))
	}
	initRes = initializerStatements.String()
	ptrRes = pointersInitializesStatements.String()
	return
}
func (partiallygrowthcurve2dribbon *PartiallyGrowthCurve2DRibbon) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(partiallygrowthcurve2dribbon.GongMarshallField(stage, "Name"))
	}
	initRes = initializerStatements.String()
	ptrRes = pointersInitializesStatements.String()
	return
}
func (partiallygrowthcurve2dribbonendshape *PartiallyGrowthCurve2DRibbonEndShape) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(partiallygrowthcurve2dribbonendshape.GongMarshallField(stage, "Name"))
		initializerStatements.WriteString(partiallygrowthcurve2dribbonendshape.GongMarshallField(stage, "BottomStartX"))
		initializerStatements.WriteString(partiallygrowthcurve2dribbonendshape.GongMarshallField(stage, "BottomStartY"))
		initializerStatements.WriteString(partiallygrowthcurve2dribbonendshape.GongMarshallField(stage, "BottomEndX"))
		initializerStatements.WriteString(partiallygrowthcurve2dribbonendshape.GongMarshallField(stage, "BottomEndY"))
		initializerStatements.WriteString(partiallygrowthcurve2dribbonendshape.GongMarshallField(stage, "BottomRadiusX"))
		initializerStatements.WriteString(partiallygrowthcurve2dribbonendshape.GongMarshallField(stage, "BottomRadiusY"))
		initializerStatements.WriteString(partiallygrowthcurve2dribbonendshape.GongMarshallField(stage, "BottomXAxisRotation"))
		initializerStatements.WriteString(partiallygrowthcurve2dribbonendshape.GongMarshallField(stage, "BottomLargeArcFlag"))
		initializerStatements.WriteString(partiallygrowthcurve2dribbonendshape.GongMarshallField(stage, "BottomSweepFlag"))
		initializerStatements.WriteString(partiallygrowthcurve2dribbonendshape.GongMarshallField(stage, "TopStartX"))
		initializerStatements.WriteString(partiallygrowthcurve2dribbonendshape.GongMarshallField(stage, "TopStartY"))
		initializerStatements.WriteString(partiallygrowthcurve2dribbonendshape.GongMarshallField(stage, "TopEndX"))
		initializerStatements.WriteString(partiallygrowthcurve2dribbonendshape.GongMarshallField(stage, "TopEndY"))
		initializerStatements.WriteString(partiallygrowthcurve2dribbonendshape.GongMarshallField(stage, "TopRadiusX"))
		initializerStatements.WriteString(partiallygrowthcurve2dribbonendshape.GongMarshallField(stage, "TopRadiusY"))
		initializerStatements.WriteString(partiallygrowthcurve2dribbonendshape.GongMarshallField(stage, "TopXAxisRotation"))
		initializerStatements.WriteString(partiallygrowthcurve2dribbonendshape.GongMarshallField(stage, "TopLargeArcFlag"))
		initializerStatements.WriteString(partiallygrowthcurve2dribbonendshape.GongMarshallField(stage, "TopSweepFlag"))
	}
	initRes = initializerStatements.String()
	ptrRes = pointersInitializesStatements.String()
	return
}
func (partiallygrowthcurve2dribbonstartshape *PartiallyGrowthCurve2DRibbonStartShape) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(partiallygrowthcurve2dribbonstartshape.GongMarshallField(stage, "Name"))
		initializerStatements.WriteString(partiallygrowthcurve2dribbonstartshape.GongMarshallField(stage, "BottomStartX"))
		initializerStatements.WriteString(partiallygrowthcurve2dribbonstartshape.GongMarshallField(stage, "BottomStartY"))
		initializerStatements.WriteString(partiallygrowthcurve2dribbonstartshape.GongMarshallField(stage, "BottomEndX"))
		initializerStatements.WriteString(partiallygrowthcurve2dribbonstartshape.GongMarshallField(stage, "BottomEndY"))
		initializerStatements.WriteString(partiallygrowthcurve2dribbonstartshape.GongMarshallField(stage, "BottomRadiusX"))
		initializerStatements.WriteString(partiallygrowthcurve2dribbonstartshape.GongMarshallField(stage, "BottomRadiusY"))
		initializerStatements.WriteString(partiallygrowthcurve2dribbonstartshape.GongMarshallField(stage, "BottomXAxisRotation"))
		initializerStatements.WriteString(partiallygrowthcurve2dribbonstartshape.GongMarshallField(stage, "BottomLargeArcFlag"))
		initializerStatements.WriteString(partiallygrowthcurve2dribbonstartshape.GongMarshallField(stage, "BottomSweepFlag"))
		initializerStatements.WriteString(partiallygrowthcurve2dribbonstartshape.GongMarshallField(stage, "TopStartX"))
		initializerStatements.WriteString(partiallygrowthcurve2dribbonstartshape.GongMarshallField(stage, "TopStartY"))
		initializerStatements.WriteString(partiallygrowthcurve2dribbonstartshape.GongMarshallField(stage, "TopEndX"))
		initializerStatements.WriteString(partiallygrowthcurve2dribbonstartshape.GongMarshallField(stage, "TopEndY"))
		initializerStatements.WriteString(partiallygrowthcurve2dribbonstartshape.GongMarshallField(stage, "TopRadiusX"))
		initializerStatements.WriteString(partiallygrowthcurve2dribbonstartshape.GongMarshallField(stage, "TopRadiusY"))
		initializerStatements.WriteString(partiallygrowthcurve2dribbonstartshape.GongMarshallField(stage, "TopXAxisRotation"))
		initializerStatements.WriteString(partiallygrowthcurve2dribbonstartshape.GongMarshallField(stage, "TopLargeArcFlag"))
		initializerStatements.WriteString(partiallygrowthcurve2dribbonstartshape.GongMarshallField(stage, "TopSweepFlag"))
	}
	initRes = initializerStatements.String()
	ptrRes = pointersInitializesStatements.String()
	return
}
func (partiallygrowthcurve2dtrajectory *PartiallyGrowthCurve2DTrajectory) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(partiallygrowthcurve2dtrajectory.GongMarshallField(stage, "Name"))
	}
	initRes = initializerStatements.String()
	ptrRes = pointersInitializesStatements.String()
	return
}
func (partiallygrowthcurve2dtrajectoryp1curveshape *PartiallyGrowthCurve2DTrajectoryP1CurveShape) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(partiallygrowthcurve2dtrajectoryp1curveshape.GongMarshallField(stage, "Name"))
		initializerStatements.WriteString(partiallygrowthcurve2dtrajectoryp1curveshape.GongMarshallField(stage, "StartX"))
		initializerStatements.WriteString(partiallygrowthcurve2dtrajectoryp1curveshape.GongMarshallField(stage, "StartY"))
		initializerStatements.WriteString(partiallygrowthcurve2dtrajectoryp1curveshape.GongMarshallField(stage, "EndX"))
		initializerStatements.WriteString(partiallygrowthcurve2dtrajectoryp1curveshape.GongMarshallField(stage, "EndY"))
	}
	initRes = initializerStatements.String()
	ptrRes = pointersInitializesStatements.String()
	return
}
func (partiallygrowthcurve2dtrajectoryp1p2 *PartiallyGrowthCurve2DTrajectoryP1P2) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(partiallygrowthcurve2dtrajectoryp1p2.GongMarshallField(stage, "Name"))
	}
	initRes = initializerStatements.String()
	ptrRes = pointersInitializesStatements.String()
	return
}
func (partiallygrowthcurve2dtrajectoryp1p2pairlineshape *PartiallyGrowthCurve2DTrajectoryP1P2PairLineShape) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(partiallygrowthcurve2dtrajectoryp1p2pairlineshape.GongMarshallField(stage, "Name"))
		initializerStatements.WriteString(partiallygrowthcurve2dtrajectoryp1p2pairlineshape.GongMarshallField(stage, "StartX"))
		initializerStatements.WriteString(partiallygrowthcurve2dtrajectoryp1p2pairlineshape.GongMarshallField(stage, "StartY"))
		initializerStatements.WriteString(partiallygrowthcurve2dtrajectoryp1p2pairlineshape.GongMarshallField(stage, "EndX"))
		initializerStatements.WriteString(partiallygrowthcurve2dtrajectoryp1p2pairlineshape.GongMarshallField(stage, "EndY"))
	}
	initRes = initializerStatements.String()
	ptrRes = pointersInitializesStatements.String()
	return
}
func (partiallygrowthcurve2dtrajectoryp1pointshape *PartiallyGrowthCurve2DTrajectoryP1PointShape) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(partiallygrowthcurve2dtrajectoryp1pointshape.GongMarshallField(stage, "Name"))
		initializerStatements.WriteString(partiallygrowthcurve2dtrajectoryp1pointshape.GongMarshallField(stage, "X"))
		initializerStatements.WriteString(partiallygrowthcurve2dtrajectoryp1pointshape.GongMarshallField(stage, "Y"))
	}
	initRes = initializerStatements.String()
	ptrRes = pointersInitializesStatements.String()
	return
}
func (partiallygrowthcurve2dtrajectoryp2curveshape *PartiallyGrowthCurve2DTrajectoryP2CurveShape) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(partiallygrowthcurve2dtrajectoryp2curveshape.GongMarshallField(stage, "Name"))
		initializerStatements.WriteString(partiallygrowthcurve2dtrajectoryp2curveshape.GongMarshallField(stage, "StartX"))
		initializerStatements.WriteString(partiallygrowthcurve2dtrajectoryp2curveshape.GongMarshallField(stage, "StartY"))
		initializerStatements.WriteString(partiallygrowthcurve2dtrajectoryp2curveshape.GongMarshallField(stage, "EndX"))
		initializerStatements.WriteString(partiallygrowthcurve2dtrajectoryp2curveshape.GongMarshallField(stage, "EndY"))
	}
	initRes = initializerStatements.String()
	ptrRes = pointersInitializesStatements.String()
	return
}
func (partiallygrowthcurve2dtrajectoryp2pointshape *PartiallyGrowthCurve2DTrajectoryP2PointShape) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(partiallygrowthcurve2dtrajectoryp2pointshape.GongMarshallField(stage, "Name"))
		initializerStatements.WriteString(partiallygrowthcurve2dtrajectoryp2pointshape.GongMarshallField(stage, "X"))
		initializerStatements.WriteString(partiallygrowthcurve2dtrajectoryp2pointshape.GongMarshallField(stage, "Y"))
	}
	initRes = initializerStatements.String()
	ptrRes = pointersInitializesStatements.String()
	return
}
func (partiallygrowthcurve2dtrajectoryshape *PartiallyGrowthCurve2DTrajectoryShape) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(partiallygrowthcurve2dtrajectoryshape.GongMarshallField(stage, "Name"))
		initializerStatements.WriteString(partiallygrowthcurve2dtrajectoryshape.GongMarshallField(stage, "StartX"))
		initializerStatements.WriteString(partiallygrowthcurve2dtrajectoryshape.GongMarshallField(stage, "StartY"))
		initializerStatements.WriteString(partiallygrowthcurve2dtrajectoryshape.GongMarshallField(stage, "EndX"))
		initializerStatements.WriteString(partiallygrowthcurve2dtrajectoryshape.GongMarshallField(stage, "EndY"))
	}
	initRes = initializerStatements.String()
	ptrRes = pointersInitializesStatements.String()
	return
}
func (partiallyrotatedseatbottomcurveshape *PartiallyRotatedSeatBottomCurveShape) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(partiallyrotatedseatbottomcurveshape.GongMarshallField(stage, "Name"))
	}
	initRes = initializerStatements.String()
	ptrRes = pointersInitializesStatements.String()
	return
}
func (partiallyrotatedseattopcurveshape *PartiallyRotatedSeatTopCurveShape) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(partiallyrotatedseattopcurveshape.GongMarshallField(stage, "Name"))
	}
	initRes = initializerStatements.String()
	ptrRes = pointersInitializesStatements.String()
	return
}
func (partiallyrotatedtorusshape *PartiallyRotatedTorusShape) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(partiallyrotatedtorusshape.GongMarshallField(stage, "Name"))
	}
	initRes = initializerStatements.String()
	ptrRes = pointersInitializesStatements.String()
	return
}
func (perpendicularvector *PerpendicularVector) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(perpendicularvector.GongMarshallField(stage, "Name"))
		initializerStatements.WriteString(perpendicularvector.GongMarshallField(stage, "StartX"))
		initializerStatements.WriteString(perpendicularvector.GongMarshallField(stage, "StartY"))
		initializerStatements.WriteString(perpendicularvector.GongMarshallField(stage, "EndX"))
		initializerStatements.WriteString(perpendicularvector.GongMarshallField(stage, "EndY"))
	}
	initRes = initializerStatements.String()
	ptrRes = pointersInitializesStatements.String()
	return
}
func (perpendicularvectorgrid *PerpendicularVectorGrid) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(perpendicularvectorgrid.GongMarshallField(stage, "Name"))
	}
	initRes = initializerStatements.String()
	ptrRes = pointersInitializesStatements.String()
	return
}
func (perpendicularvectorgridhalfway *PerpendicularVectorGridHalfway) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(perpendicularvectorgridhalfway.GongMarshallField(stage, "Name"))
	}
	initRes = initializerStatements.String()
	ptrRes = pointersInitializesStatements.String()
	return
}
func (perpendicularvectorhalfway *PerpendicularVectorHalfway) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(perpendicularvectorhalfway.GongMarshallField(stage, "Name"))
		initializerStatements.WriteString(perpendicularvectorhalfway.GongMarshallField(stage, "StartX"))
		initializerStatements.WriteString(perpendicularvectorhalfway.GongMarshallField(stage, "StartY"))
		initializerStatements.WriteString(perpendicularvectorhalfway.GongMarshallField(stage, "EndX"))
		initializerStatements.WriteString(perpendicularvectorhalfway.GongMarshallField(stage, "EndY"))
	}
	initRes = initializerStatements.String()
	ptrRes = pointersInitializesStatements.String()
	return
}
func (plant2ddiagram *Plant2DDiagram) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(plant2ddiagram.GongMarshallField(stage, "Name"))
		initializerStatements.WriteString(plant2ddiagram.GongMarshallField(stage, "OriginX"))
		initializerStatements.WriteString(plant2ddiagram.GongMarshallField(stage, "OriginY"))
		initializerStatements.WriteString(plant2ddiagram.GongMarshallField(stage, "Zoom"))
		initializerStatements.WriteString(plant2ddiagram.GongMarshallField(stage, "IsRhombusNodesExpanded"))
		initializerStatements.WriteString(plant2ddiagram.GongMarshallField(stage, "IsArcNodesExpanded"))
		initializerStatements.WriteString(plant2ddiagram.GongMarshallField(stage, "IsHiddenAxesShape"))
		initializerStatements.WriteString(plant2ddiagram.GongMarshallField(stage, "IsHiddenReferenceRhombus"))
		initializerStatements.WriteString(plant2ddiagram.GongMarshallField(stage, "IsHiddenPlantCircumferenceShape"))
		initializerStatements.WriteString(plant2ddiagram.GongMarshallField(stage, "IsHiddenGridPathShape"))
		initializerStatements.WriteString(plant2ddiagram.GongMarshallField(stage, "IsHiddenRhombusGridShape"))
		initializerStatements.WriteString(plant2ddiagram.GongMarshallField(stage, "IsHiddenExplanationTextShape"))
		initializerStatements.WriteString(plant2ddiagram.GongMarshallField(stage, "IsHiddenRotatedReferenceRhombus"))
		initializerStatements.WriteString(plant2ddiagram.GongMarshallField(stage, "IsHiddenRotatedPlantCircumferenceShape"))
		initializerStatements.WriteString(plant2ddiagram.GongMarshallField(stage, "IsHiddenRotatedGridPathShape"))
		initializerStatements.WriteString(plant2ddiagram.GongMarshallField(stage, "IsHiddenRotatedRhombusGridShape"))
		initializerStatements.WriteString(plant2ddiagram.GongMarshallField(stage, "IsHiddenGrowthPathRhombusGridShape"))
		initializerStatements.WriteString(plant2ddiagram.GongMarshallField(stage, "IsHiddenGrowthVectorShape"))
		initializerStatements.WriteString(plant2ddiagram.GongMarshallField(stage, "IsHiddenPerpendicularVectorGrid"))
		initializerStatements.WriteString(plant2ddiagram.GongMarshallField(stage, "IsHiddenBaseVectorShapeGrid"))
		initializerStatements.WriteString(plant2ddiagram.GongMarshallField(stage, "IsHiddenArcNormalVectorShapeGrid"))
		initializerStatements.WriteString(plant2ddiagram.GongMarshallField(stage, "IsHiddenStartArcShapeGrid"))
		initializerStatements.WriteString(plant2ddiagram.GongMarshallField(stage, "IsHiddenMidArcVectorShapeGrid"))
		initializerStatements.WriteString(plant2ddiagram.GongMarshallField(stage, "IsHiddenEndArcShapeGrid"))
		initializerStatements.WriteString(plant2ddiagram.GongMarshallField(stage, "IsHiddenGrowthCurve2D"))
		initializerStatements.WriteString(plant2ddiagram.GongMarshallField(stage, "IsHiddenStackOfGrowthCurve2DByGrowthVector"))
		initializerStatements.WriteString(plant2ddiagram.GongMarshallField(stage, "IsChecked"))
		initializerStatements.WriteString(plant2ddiagram.GongMarshallField(stage, "ComputedPrefix"))
		initializerStatements.WriteString(plant2ddiagram.GongMarshallField(stage, "IsExpanded"))
	}
	initRes = initializerStatements.String()
	ptrRes = pointersInitializesStatements.String()
	return
}
func (plant3ddiagram *Plant3DDiagram) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(plant3ddiagram.GongMarshallField(stage, "Name"))
		initializerStatements.WriteString(plant3ddiagram.GongMarshallField(stage, "IsHiddenStemCylinder3DShape"))
		pointersInitializesStatements.WriteString(plant3ddiagram.GongMarshallField(stage, "StemCylinder3DShape"))
		initializerStatements.WriteString(plant3ddiagram.GongMarshallField(stage, "IsHiddenParastichyNCurves3DShape"))
		pointersInitializesStatements.WriteString(plant3ddiagram.GongMarshallField(stage, "ParastichyNCurves3DShape"))
		initializerStatements.WriteString(plant3ddiagram.GongMarshallField(stage, "IsHiddenParastichyMCurves3DShape"))
		pointersInitializesStatements.WriteString(plant3ddiagram.GongMarshallField(stage, "ParastichyMCurves3DShape"))
		initializerStatements.WriteString(plant3ddiagram.GongMarshallField(stage, "IsHiddenCutLine3DShape"))
		pointersInitializesStatements.WriteString(plant3ddiagram.GongMarshallField(stage, "CutLine3DShape"))
		initializerStatements.WriteString(plant3ddiagram.GongMarshallField(stage, "IsHiddenCircumference3DShape"))
		pointersInitializesStatements.WriteString(plant3ddiagram.GongMarshallField(stage, "Circumference3DShape"))
		initializerStatements.WriteString(plant3ddiagram.GongMarshallField(stage, "IsHiddenTiledFloor3DShape"))
		initializerStatements.WriteString(plant3ddiagram.GongMarshallField(stage, "IsHiddenLeaves3DShape"))
		pointersInitializesStatements.WriteString(plant3ddiagram.GongMarshallField(stage, "Leaves3DShape"))
		pointersInitializesStatements.WriteString(plant3ddiagram.GongMarshallField(stage, "Rendered3DShape"))
		initializerStatements.WriteString(plant3ddiagram.GongMarshallField(stage, "IsChecked"))
		initializerStatements.WriteString(plant3ddiagram.GongMarshallField(stage, "ComputedPrefix"))
		initializerStatements.WriteString(plant3ddiagram.GongMarshallField(stage, "IsExpanded"))
	}
	initRes = initializerStatements.String()
	ptrRes = pointersInitializesStatements.String()
	return
}
func (plantabstract *PlantAbstract) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(plantabstract.GongMarshallField(stage, "Name"))
		initializerStatements.WriteString(plantabstract.GongMarshallField(stage, "N"))
		initializerStatements.WriteString(plantabstract.GongMarshallField(stage, "M"))
		initializerStatements.WriteString(plantabstract.GongMarshallField(stage, "StackHeight"))
		initializerStatements.WriteString(plantabstract.GongMarshallField(stage, "RhombusInsideAngle"))
		initializerStatements.WriteString(plantabstract.GongMarshallField(stage, "RhombusSideLength"))
		initializerStatements.WriteString(plantabstract.GongMarshallField(stage, "PlantType"))
		pointersInitializesStatements.WriteString(plantabstract.GongMarshallField(stage, "TubeVaseAbstract"))
		initializerStatements.WriteString(plantabstract.GongMarshallField(stage, "CurrentView"))
		initializerStatements.WriteString(plantabstract.GongMarshallField(stage, "ComputedPrefix"))
		initializerStatements.WriteString(plantabstract.GongMarshallField(stage, "IsExpanded"))
		initializerStatements.WriteString(plantabstract.GongMarshallField(stage, "IsSelected"))
		initializerStatements.WriteString(plantabstract.GongMarshallField(stage, "IsPlant2DDiagramsNodeExpanded"))
		pointersInitializesStatements.WriteString(plantabstract.GongMarshallField(stage, "Plant2DDiagrams"))
		initializerStatements.WriteString(plantabstract.GongMarshallField(stage, "IsPlant3DDiagramsNodeExpanded"))
		pointersInitializesStatements.WriteString(plantabstract.GongMarshallField(stage, "Plant3DDiagrams"))
		initializerStatements.WriteString(plantabstract.GongMarshallField(stage, "IsVase2DDiagramsNodeExpanded"))
		pointersInitializesStatements.WriteString(plantabstract.GongMarshallField(stage, "Vase2DDiagrams"))
		initializerStatements.WriteString(plantabstract.GongMarshallField(stage, "IsTubeVase3DDiagramsNodeExpanded"))
		pointersInitializesStatements.WriteString(plantabstract.GongMarshallField(stage, "TubeVase3DDiagrams"))
		initializerStatements.WriteString(plantabstract.GongMarshallField(stage, "IsStool2DDiagramsNodeExpanded"))
		pointersInitializesStatements.WriteString(plantabstract.GongMarshallField(stage, "Stool2DDiagrams"))
		initializerStatements.WriteString(plantabstract.GongMarshallField(stage, "IsStool3DDiagramsNodeExpanded"))
		pointersInitializesStatements.WriteString(plantabstract.GongMarshallField(stage, "Stool3DDiagrams"))
		initializerStatements.WriteString(plantabstract.GongMarshallField(stage, "IsClock2DDiagramsNodeExpanded"))
		pointersInitializesStatements.WriteString(plantabstract.GongMarshallField(stage, "Clock2DDiagrams"))
		initializerStatements.WriteString(plantabstract.GongMarshallField(stage, "IsClock3DDiagramsNodeExpanded"))
		pointersInitializesStatements.WriteString(plantabstract.GongMarshallField(stage, "Clock3DDiagrams"))
	}
	initRes = initializerStatements.String()
	ptrRes = pointersInitializesStatements.String()
	return
}
func (plantcircumferenceshape *PlantCircumferenceShape) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(plantcircumferenceshape.GongMarshallField(stage, "Name"))
		initializerStatements.WriteString(plantcircumferenceshape.GongMarshallField(stage, "AngleDegree"))
		initializerStatements.WriteString(plantcircumferenceshape.GongMarshallField(stage, "Length"))
	}
	initRes = initializerStatements.String()
	ptrRes = pointersInitializesStatements.String()
	return
}
func (pointsandlines3dshape *PointsAndLines3DShape) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(pointsandlines3dshape.GongMarshallField(stage, "Name"))
	}
	initRes = initializerStatements.String()
	ptrRes = pointersInitializesStatements.String()
	return
}
func (pxshape *PxShape) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(pxshape.GongMarshallField(stage, "Name"))
		initializerStatements.WriteString(pxshape.GongMarshallField(stage, "X"))
		initializerStatements.WriteString(pxshape.GongMarshallField(stage, "Y"))
	}
	initRes = initializerStatements.String()
	ptrRes = pointersInitializesStatements.String()
	return
}
func (rendered3dshape *Rendered3DShape) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(rendered3dshape.GongMarshallField(stage, "Name"))
		initializerStatements.WriteString(rendered3dshape.GongMarshallField(stage, "ViewX"))
		initializerStatements.WriteString(rendered3dshape.GongMarshallField(stage, "ViewY"))
		initializerStatements.WriteString(rendered3dshape.GongMarshallField(stage, "ViewZ"))
		initializerStatements.WriteString(rendered3dshape.GongMarshallField(stage, "TargetX"))
		initializerStatements.WriteString(rendered3dshape.GongMarshallField(stage, "TargetY"))
		initializerStatements.WriteString(rendered3dshape.GongMarshallField(stage, "TargetZ"))
		initializerStatements.WriteString(rendered3dshape.GongMarshallField(stage, "Fov"))
	}
	initRes = initializerStatements.String()
	ptrRes = pointersInitializesStatements.String()
	return
}
func (rhombusshape *RhombusShape) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(rhombusshape.GongMarshallField(stage, "Name"))
		initializerStatements.WriteString(rhombusshape.GongMarshallField(stage, "X"))
		initializerStatements.WriteString(rhombusshape.GongMarshallField(stage, "Y"))
	}
	initRes = initializerStatements.String()
	ptrRes = pointersInitializesStatements.String()
	return
}
func (rhombusstuff *RhombusStuff) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(rhombusstuff.GongMarshallField(stage, "Name"))
	}
	initRes = initializerStatements.String()
	ptrRes = pointersInitializesStatements.String()
	return
}
func (rotatedrhombusgridshape *RotatedRhombusGridShape) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(rotatedrhombusgridshape.GongMarshallField(stage, "Name"))
	}
	initRes = initializerStatements.String()
	ptrRes = pointersInitializesStatements.String()
	return
}
func (rotatedrhombusshape *RotatedRhombusShape) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(rotatedrhombusshape.GongMarshallField(stage, "Name"))
		initializerStatements.WriteString(rotatedrhombusshape.GongMarshallField(stage, "X"))
		initializerStatements.WriteString(rotatedrhombusshape.GongMarshallField(stage, "Y"))
	}
	initRes = initializerStatements.String()
	ptrRes = pointersInitializesStatements.String()
	return
}
func (rotatedsampledpoints3dshape *RotatedSampledPoints3DShape) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(rotatedsampledpoints3dshape.GongMarshallField(stage, "Name"))
	}
	initRes = initializerStatements.String()
	ptrRes = pointersInitializesStatements.String()
	return
}
func (rotatedseatandlegs3dshape *RotatedSeatAndLegs3DShape) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(rotatedseatandlegs3dshape.GongMarshallField(stage, "Name"))
	}
	initRes = initializerStatements.String()
	ptrRes = pointersInitializesStatements.String()
	return
}
func (sampledpoints3dshape *SampledPoints3DShape) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(sampledpoints3dshape.GongMarshallField(stage, "Name"))
	}
	initRes = initializerStatements.String()
	ptrRes = pointersInitializesStatements.String()
	return
}
func (seat3dshape *Seat3DShape) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(seat3dshape.GongMarshallField(stage, "Name"))
	}
	initRes = initializerStatements.String()
	ptrRes = pointersInitializesStatements.String()
	return
}
func (seatandlegs3dshape *SeatAndLegs3DShape) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(seatandlegs3dshape.GongMarshallField(stage, "Name"))
	}
	initRes = initializerStatements.String()
	ptrRes = pointersInitializesStatements.String()
	return
}
func (seatbottomcurveshape *SeatBottomCurveShape) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(seatbottomcurveshape.GongMarshallField(stage, "Name"))
	}
	initRes = initializerStatements.String()
	ptrRes = pointersInitializesStatements.String()
	return
}
func (seattopcurveshape *SeatTopCurveShape) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(seattopcurveshape.GongMarshallField(stage, "Name"))
	}
	initRes = initializerStatements.String()
	ptrRes = pointersInitializesStatements.String()
	return
}
func (shiftedbottomtopstartarcshape *ShiftedBottomTopStartArcShape) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(shiftedbottomtopstartarcshape.GongMarshallField(stage, "Name"))
		initializerStatements.WriteString(shiftedbottomtopstartarcshape.GongMarshallField(stage, "StartX"))
		initializerStatements.WriteString(shiftedbottomtopstartarcshape.GongMarshallField(stage, "StartY"))
		initializerStatements.WriteString(shiftedbottomtopstartarcshape.GongMarshallField(stage, "EndX"))
		initializerStatements.WriteString(shiftedbottomtopstartarcshape.GongMarshallField(stage, "EndY"))
		initializerStatements.WriteString(shiftedbottomtopstartarcshape.GongMarshallField(stage, "XAxisRotation"))
		initializerStatements.WriteString(shiftedbottomtopstartarcshape.GongMarshallField(stage, "LargeArcFlag"))
		initializerStatements.WriteString(shiftedbottomtopstartarcshape.GongMarshallField(stage, "SweepFlag"))
		initializerStatements.WriteString(shiftedbottomtopstartarcshape.GongMarshallField(stage, "RadiusX"))
		initializerStatements.WriteString(shiftedbottomtopstartarcshape.GongMarshallField(stage, "RadiusY"))
	}
	initRes = initializerStatements.String()
	ptrRes = pointersInitializesStatements.String()
	return
}
func (shiftedbottomtopstartarcshapegrid *ShiftedBottomTopStartArcShapeGrid) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(shiftedbottomtopstartarcshapegrid.GongMarshallField(stage, "Name"))
	}
	initRes = initializerStatements.String()
	ptrRes = pointersInitializesStatements.String()
	return
}
func (shiftedleftgrowthcurve2dribbon *ShiftedLeftGrowthCurve2DRibbon) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(shiftedleftgrowthcurve2dribbon.GongMarshallField(stage, "Name"))
	}
	initRes = initializerStatements.String()
	ptrRes = pointersInitializesStatements.String()
	return
}
func (shiftedleftgrowthcurve2dribbonendshape *ShiftedLeftGrowthCurve2DRibbonEndShape) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(shiftedleftgrowthcurve2dribbonendshape.GongMarshallField(stage, "Name"))
		initializerStatements.WriteString(shiftedleftgrowthcurve2dribbonendshape.GongMarshallField(stage, "BottomStartX"))
		initializerStatements.WriteString(shiftedleftgrowthcurve2dribbonendshape.GongMarshallField(stage, "BottomStartY"))
		initializerStatements.WriteString(shiftedleftgrowthcurve2dribbonendshape.GongMarshallField(stage, "BottomEndX"))
		initializerStatements.WriteString(shiftedleftgrowthcurve2dribbonendshape.GongMarshallField(stage, "BottomEndY"))
		initializerStatements.WriteString(shiftedleftgrowthcurve2dribbonendshape.GongMarshallField(stage, "BottomRadiusX"))
		initializerStatements.WriteString(shiftedleftgrowthcurve2dribbonendshape.GongMarshallField(stage, "BottomRadiusY"))
		initializerStatements.WriteString(shiftedleftgrowthcurve2dribbonendshape.GongMarshallField(stage, "BottomXAxisRotation"))
		initializerStatements.WriteString(shiftedleftgrowthcurve2dribbonendshape.GongMarshallField(stage, "BottomLargeArcFlag"))
		initializerStatements.WriteString(shiftedleftgrowthcurve2dribbonendshape.GongMarshallField(stage, "BottomSweepFlag"))
		initializerStatements.WriteString(shiftedleftgrowthcurve2dribbonendshape.GongMarshallField(stage, "TopStartX"))
		initializerStatements.WriteString(shiftedleftgrowthcurve2dribbonendshape.GongMarshallField(stage, "TopStartY"))
		initializerStatements.WriteString(shiftedleftgrowthcurve2dribbonendshape.GongMarshallField(stage, "TopEndX"))
		initializerStatements.WriteString(shiftedleftgrowthcurve2dribbonendshape.GongMarshallField(stage, "TopEndY"))
		initializerStatements.WriteString(shiftedleftgrowthcurve2dribbonendshape.GongMarshallField(stage, "TopRadiusX"))
		initializerStatements.WriteString(shiftedleftgrowthcurve2dribbonendshape.GongMarshallField(stage, "TopRadiusY"))
		initializerStatements.WriteString(shiftedleftgrowthcurve2dribbonendshape.GongMarshallField(stage, "TopXAxisRotation"))
		initializerStatements.WriteString(shiftedleftgrowthcurve2dribbonendshape.GongMarshallField(stage, "TopLargeArcFlag"))
		initializerStatements.WriteString(shiftedleftgrowthcurve2dribbonendshape.GongMarshallField(stage, "TopSweepFlag"))
	}
	initRes = initializerStatements.String()
	ptrRes = pointersInitializesStatements.String()
	return
}
func (shiftedleftgrowthcurve2dribbonstartshape *ShiftedLeftGrowthCurve2DRibbonStartShape) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(shiftedleftgrowthcurve2dribbonstartshape.GongMarshallField(stage, "Name"))
		initializerStatements.WriteString(shiftedleftgrowthcurve2dribbonstartshape.GongMarshallField(stage, "BottomStartX"))
		initializerStatements.WriteString(shiftedleftgrowthcurve2dribbonstartshape.GongMarshallField(stage, "BottomStartY"))
		initializerStatements.WriteString(shiftedleftgrowthcurve2dribbonstartshape.GongMarshallField(stage, "BottomEndX"))
		initializerStatements.WriteString(shiftedleftgrowthcurve2dribbonstartshape.GongMarshallField(stage, "BottomEndY"))
		initializerStatements.WriteString(shiftedleftgrowthcurve2dribbonstartshape.GongMarshallField(stage, "BottomRadiusX"))
		initializerStatements.WriteString(shiftedleftgrowthcurve2dribbonstartshape.GongMarshallField(stage, "BottomRadiusY"))
		initializerStatements.WriteString(shiftedleftgrowthcurve2dribbonstartshape.GongMarshallField(stage, "BottomXAxisRotation"))
		initializerStatements.WriteString(shiftedleftgrowthcurve2dribbonstartshape.GongMarshallField(stage, "BottomLargeArcFlag"))
		initializerStatements.WriteString(shiftedleftgrowthcurve2dribbonstartshape.GongMarshallField(stage, "BottomSweepFlag"))
		initializerStatements.WriteString(shiftedleftgrowthcurve2dribbonstartshape.GongMarshallField(stage, "TopStartX"))
		initializerStatements.WriteString(shiftedleftgrowthcurve2dribbonstartshape.GongMarshallField(stage, "TopStartY"))
		initializerStatements.WriteString(shiftedleftgrowthcurve2dribbonstartshape.GongMarshallField(stage, "TopEndX"))
		initializerStatements.WriteString(shiftedleftgrowthcurve2dribbonstartshape.GongMarshallField(stage, "TopEndY"))
		initializerStatements.WriteString(shiftedleftgrowthcurve2dribbonstartshape.GongMarshallField(stage, "TopRadiusX"))
		initializerStatements.WriteString(shiftedleftgrowthcurve2dribbonstartshape.GongMarshallField(stage, "TopRadiusY"))
		initializerStatements.WriteString(shiftedleftgrowthcurve2dribbonstartshape.GongMarshallField(stage, "TopXAxisRotation"))
		initializerStatements.WriteString(shiftedleftgrowthcurve2dribbonstartshape.GongMarshallField(stage, "TopLargeArcFlag"))
		initializerStatements.WriteString(shiftedleftgrowthcurve2dribbonstartshape.GongMarshallField(stage, "TopSweepFlag"))
	}
	initRes = initializerStatements.String()
	ptrRes = pointersInitializesStatements.String()
	return
}
func (shiftedleftpartiallygrowthcurve2dribbon *ShiftedLeftPartiallyGrowthCurve2DRibbon) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(shiftedleftpartiallygrowthcurve2dribbon.GongMarshallField(stage, "Name"))
	}
	initRes = initializerStatements.String()
	ptrRes = pointersInitializesStatements.String()
	return
}
func (shiftedleftpartiallygrowthcurve2dribbonendshape *ShiftedLeftPartiallyGrowthCurve2DRibbonEndShape) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(shiftedleftpartiallygrowthcurve2dribbonendshape.GongMarshallField(stage, "Name"))
		initializerStatements.WriteString(shiftedleftpartiallygrowthcurve2dribbonendshape.GongMarshallField(stage, "BottomStartX"))
		initializerStatements.WriteString(shiftedleftpartiallygrowthcurve2dribbonendshape.GongMarshallField(stage, "BottomStartY"))
		initializerStatements.WriteString(shiftedleftpartiallygrowthcurve2dribbonendshape.GongMarshallField(stage, "BottomEndX"))
		initializerStatements.WriteString(shiftedleftpartiallygrowthcurve2dribbonendshape.GongMarshallField(stage, "BottomEndY"))
		initializerStatements.WriteString(shiftedleftpartiallygrowthcurve2dribbonendshape.GongMarshallField(stage, "BottomRadiusX"))
		initializerStatements.WriteString(shiftedleftpartiallygrowthcurve2dribbonendshape.GongMarshallField(stage, "BottomRadiusY"))
		initializerStatements.WriteString(shiftedleftpartiallygrowthcurve2dribbonendshape.GongMarshallField(stage, "BottomXAxisRotation"))
		initializerStatements.WriteString(shiftedleftpartiallygrowthcurve2dribbonendshape.GongMarshallField(stage, "BottomLargeArcFlag"))
		initializerStatements.WriteString(shiftedleftpartiallygrowthcurve2dribbonendshape.GongMarshallField(stage, "BottomSweepFlag"))
		initializerStatements.WriteString(shiftedleftpartiallygrowthcurve2dribbonendshape.GongMarshallField(stage, "TopStartX"))
		initializerStatements.WriteString(shiftedleftpartiallygrowthcurve2dribbonendshape.GongMarshallField(stage, "TopStartY"))
		initializerStatements.WriteString(shiftedleftpartiallygrowthcurve2dribbonendshape.GongMarshallField(stage, "TopEndX"))
		initializerStatements.WriteString(shiftedleftpartiallygrowthcurve2dribbonendshape.GongMarshallField(stage, "TopEndY"))
		initializerStatements.WriteString(shiftedleftpartiallygrowthcurve2dribbonendshape.GongMarshallField(stage, "TopRadiusX"))
		initializerStatements.WriteString(shiftedleftpartiallygrowthcurve2dribbonendshape.GongMarshallField(stage, "TopRadiusY"))
		initializerStatements.WriteString(shiftedleftpartiallygrowthcurve2dribbonendshape.GongMarshallField(stage, "TopXAxisRotation"))
		initializerStatements.WriteString(shiftedleftpartiallygrowthcurve2dribbonendshape.GongMarshallField(stage, "TopLargeArcFlag"))
		initializerStatements.WriteString(shiftedleftpartiallygrowthcurve2dribbonendshape.GongMarshallField(stage, "TopSweepFlag"))
	}
	initRes = initializerStatements.String()
	ptrRes = pointersInitializesStatements.String()
	return
}
func (shiftedleftpartiallygrowthcurve2dribbonstartshape *ShiftedLeftPartiallyGrowthCurve2DRibbonStartShape) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(shiftedleftpartiallygrowthcurve2dribbonstartshape.GongMarshallField(stage, "Name"))
		initializerStatements.WriteString(shiftedleftpartiallygrowthcurve2dribbonstartshape.GongMarshallField(stage, "BottomStartX"))
		initializerStatements.WriteString(shiftedleftpartiallygrowthcurve2dribbonstartshape.GongMarshallField(stage, "BottomStartY"))
		initializerStatements.WriteString(shiftedleftpartiallygrowthcurve2dribbonstartshape.GongMarshallField(stage, "BottomEndX"))
		initializerStatements.WriteString(shiftedleftpartiallygrowthcurve2dribbonstartshape.GongMarshallField(stage, "BottomEndY"))
		initializerStatements.WriteString(shiftedleftpartiallygrowthcurve2dribbonstartshape.GongMarshallField(stage, "BottomRadiusX"))
		initializerStatements.WriteString(shiftedleftpartiallygrowthcurve2dribbonstartshape.GongMarshallField(stage, "BottomRadiusY"))
		initializerStatements.WriteString(shiftedleftpartiallygrowthcurve2dribbonstartshape.GongMarshallField(stage, "BottomXAxisRotation"))
		initializerStatements.WriteString(shiftedleftpartiallygrowthcurve2dribbonstartshape.GongMarshallField(stage, "BottomLargeArcFlag"))
		initializerStatements.WriteString(shiftedleftpartiallygrowthcurve2dribbonstartshape.GongMarshallField(stage, "BottomSweepFlag"))
		initializerStatements.WriteString(shiftedleftpartiallygrowthcurve2dribbonstartshape.GongMarshallField(stage, "TopStartX"))
		initializerStatements.WriteString(shiftedleftpartiallygrowthcurve2dribbonstartshape.GongMarshallField(stage, "TopStartY"))
		initializerStatements.WriteString(shiftedleftpartiallygrowthcurve2dribbonstartshape.GongMarshallField(stage, "TopEndX"))
		initializerStatements.WriteString(shiftedleftpartiallygrowthcurve2dribbonstartshape.GongMarshallField(stage, "TopEndY"))
		initializerStatements.WriteString(shiftedleftpartiallygrowthcurve2dribbonstartshape.GongMarshallField(stage, "TopRadiusX"))
		initializerStatements.WriteString(shiftedleftpartiallygrowthcurve2dribbonstartshape.GongMarshallField(stage, "TopRadiusY"))
		initializerStatements.WriteString(shiftedleftpartiallygrowthcurve2dribbonstartshape.GongMarshallField(stage, "TopXAxisRotation"))
		initializerStatements.WriteString(shiftedleftpartiallygrowthcurve2dribbonstartshape.GongMarshallField(stage, "TopLargeArcFlag"))
		initializerStatements.WriteString(shiftedleftpartiallygrowthcurve2dribbonstartshape.GongMarshallField(stage, "TopSweepFlag"))
	}
	initRes = initializerStatements.String()
	ptrRes = pointersInitializesStatements.String()
	return
}
func (shiftedleftstackgrowthcurveendarcshape *ShiftedLeftStackGrowthCurveEndArcShape) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(shiftedleftstackgrowthcurveendarcshape.GongMarshallField(stage, "Name"))
		initializerStatements.WriteString(shiftedleftstackgrowthcurveendarcshape.GongMarshallField(stage, "StartX"))
		initializerStatements.WriteString(shiftedleftstackgrowthcurveendarcshape.GongMarshallField(stage, "StartY"))
		initializerStatements.WriteString(shiftedleftstackgrowthcurveendarcshape.GongMarshallField(stage, "EndX"))
		initializerStatements.WriteString(shiftedleftstackgrowthcurveendarcshape.GongMarshallField(stage, "EndY"))
		initializerStatements.WriteString(shiftedleftstackgrowthcurveendarcshape.GongMarshallField(stage, "XAxisRotation"))
		initializerStatements.WriteString(shiftedleftstackgrowthcurveendarcshape.GongMarshallField(stage, "LargeArcFlag"))
		initializerStatements.WriteString(shiftedleftstackgrowthcurveendarcshape.GongMarshallField(stage, "SweepFlag"))
		initializerStatements.WriteString(shiftedleftstackgrowthcurveendarcshape.GongMarshallField(stage, "RadiusX"))
		initializerStatements.WriteString(shiftedleftstackgrowthcurveendarcshape.GongMarshallField(stage, "RadiusY"))
	}
	initRes = initializerStatements.String()
	ptrRes = pointersInitializesStatements.String()
	return
}
func (shiftedleftstackgrowthcurvestartarcshape *ShiftedLeftStackGrowthCurveStartArcShape) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(shiftedleftstackgrowthcurvestartarcshape.GongMarshallField(stage, "Name"))
		initializerStatements.WriteString(shiftedleftstackgrowthcurvestartarcshape.GongMarshallField(stage, "StartX"))
		initializerStatements.WriteString(shiftedleftstackgrowthcurvestartarcshape.GongMarshallField(stage, "StartY"))
		initializerStatements.WriteString(shiftedleftstackgrowthcurvestartarcshape.GongMarshallField(stage, "EndX"))
		initializerStatements.WriteString(shiftedleftstackgrowthcurvestartarcshape.GongMarshallField(stage, "EndY"))
		initializerStatements.WriteString(shiftedleftstackgrowthcurvestartarcshape.GongMarshallField(stage, "XAxisRotation"))
		initializerStatements.WriteString(shiftedleftstackgrowthcurvestartarcshape.GongMarshallField(stage, "LargeArcFlag"))
		initializerStatements.WriteString(shiftedleftstackgrowthcurvestartarcshape.GongMarshallField(stage, "SweepFlag"))
		initializerStatements.WriteString(shiftedleftstackgrowthcurvestartarcshape.GongMarshallField(stage, "RadiusX"))
		initializerStatements.WriteString(shiftedleftstackgrowthcurvestartarcshape.GongMarshallField(stage, "RadiusY"))
	}
	initRes = initializerStatements.String()
	ptrRes = pointersInitializesStatements.String()
	return
}
func (shiftedleftstacknormalvector *ShiftedLeftStackNormalVector) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(shiftedleftstacknormalvector.GongMarshallField(stage, "Name"))
		initializerStatements.WriteString(shiftedleftstacknormalvector.GongMarshallField(stage, "StartX"))
		initializerStatements.WriteString(shiftedleftstacknormalvector.GongMarshallField(stage, "StartY"))
		initializerStatements.WriteString(shiftedleftstacknormalvector.GongMarshallField(stage, "EndX"))
		initializerStatements.WriteString(shiftedleftstacknormalvector.GongMarshallField(stage, "EndY"))
	}
	initRes = initializerStatements.String()
	ptrRes = pointersInitializesStatements.String()
	return
}
func (shiftedleftstackofgrowthcurve *ShiftedLeftStackOfGrowthCurve) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(shiftedleftstackofgrowthcurve.GongMarshallField(stage, "Name"))
	}
	initRes = initializerStatements.String()
	ptrRes = pointersInitializesStatements.String()
	return
}
func (shiftedleftstackofnormalvector *ShiftedLeftStackOfNormalVector) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(shiftedleftstackofnormalvector.GongMarshallField(stage, "Name"))
	}
	initRes = initializerStatements.String()
	ptrRes = pointersInitializesStatements.String()
	return
}
func (shiftedrightgrowthcurve2dribbon *ShiftedRightGrowthCurve2DRibbon) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(shiftedrightgrowthcurve2dribbon.GongMarshallField(stage, "Name"))
	}
	initRes = initializerStatements.String()
	ptrRes = pointersInitializesStatements.String()
	return
}
func (shiftedrightgrowthcurve2dribbonendshape *ShiftedRightGrowthCurve2DRibbonEndShape) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(shiftedrightgrowthcurve2dribbonendshape.GongMarshallField(stage, "Name"))
		initializerStatements.WriteString(shiftedrightgrowthcurve2dribbonendshape.GongMarshallField(stage, "BottomStartX"))
		initializerStatements.WriteString(shiftedrightgrowthcurve2dribbonendshape.GongMarshallField(stage, "BottomStartY"))
		initializerStatements.WriteString(shiftedrightgrowthcurve2dribbonendshape.GongMarshallField(stage, "BottomEndX"))
		initializerStatements.WriteString(shiftedrightgrowthcurve2dribbonendshape.GongMarshallField(stage, "BottomEndY"))
		initializerStatements.WriteString(shiftedrightgrowthcurve2dribbonendshape.GongMarshallField(stage, "BottomRadiusX"))
		initializerStatements.WriteString(shiftedrightgrowthcurve2dribbonendshape.GongMarshallField(stage, "BottomRadiusY"))
		initializerStatements.WriteString(shiftedrightgrowthcurve2dribbonendshape.GongMarshallField(stage, "BottomXAxisRotation"))
		initializerStatements.WriteString(shiftedrightgrowthcurve2dribbonendshape.GongMarshallField(stage, "BottomLargeArcFlag"))
		initializerStatements.WriteString(shiftedrightgrowthcurve2dribbonendshape.GongMarshallField(stage, "BottomSweepFlag"))
		initializerStatements.WriteString(shiftedrightgrowthcurve2dribbonendshape.GongMarshallField(stage, "TopStartX"))
		initializerStatements.WriteString(shiftedrightgrowthcurve2dribbonendshape.GongMarshallField(stage, "TopStartY"))
		initializerStatements.WriteString(shiftedrightgrowthcurve2dribbonendshape.GongMarshallField(stage, "TopEndX"))
		initializerStatements.WriteString(shiftedrightgrowthcurve2dribbonendshape.GongMarshallField(stage, "TopEndY"))
		initializerStatements.WriteString(shiftedrightgrowthcurve2dribbonendshape.GongMarshallField(stage, "TopRadiusX"))
		initializerStatements.WriteString(shiftedrightgrowthcurve2dribbonendshape.GongMarshallField(stage, "TopRadiusY"))
		initializerStatements.WriteString(shiftedrightgrowthcurve2dribbonendshape.GongMarshallField(stage, "TopXAxisRotation"))
		initializerStatements.WriteString(shiftedrightgrowthcurve2dribbonendshape.GongMarshallField(stage, "TopLargeArcFlag"))
		initializerStatements.WriteString(shiftedrightgrowthcurve2dribbonendshape.GongMarshallField(stage, "TopSweepFlag"))
	}
	initRes = initializerStatements.String()
	ptrRes = pointersInitializesStatements.String()
	return
}
func (shiftedrightgrowthcurve2dribbonstartshape *ShiftedRightGrowthCurve2DRibbonStartShape) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(shiftedrightgrowthcurve2dribbonstartshape.GongMarshallField(stage, "Name"))
		initializerStatements.WriteString(shiftedrightgrowthcurve2dribbonstartshape.GongMarshallField(stage, "BottomStartX"))
		initializerStatements.WriteString(shiftedrightgrowthcurve2dribbonstartshape.GongMarshallField(stage, "BottomStartY"))
		initializerStatements.WriteString(shiftedrightgrowthcurve2dribbonstartshape.GongMarshallField(stage, "BottomEndX"))
		initializerStatements.WriteString(shiftedrightgrowthcurve2dribbonstartshape.GongMarshallField(stage, "BottomEndY"))
		initializerStatements.WriteString(shiftedrightgrowthcurve2dribbonstartshape.GongMarshallField(stage, "BottomRadiusX"))
		initializerStatements.WriteString(shiftedrightgrowthcurve2dribbonstartshape.GongMarshallField(stage, "BottomRadiusY"))
		initializerStatements.WriteString(shiftedrightgrowthcurve2dribbonstartshape.GongMarshallField(stage, "BottomXAxisRotation"))
		initializerStatements.WriteString(shiftedrightgrowthcurve2dribbonstartshape.GongMarshallField(stage, "BottomLargeArcFlag"))
		initializerStatements.WriteString(shiftedrightgrowthcurve2dribbonstartshape.GongMarshallField(stage, "BottomSweepFlag"))
		initializerStatements.WriteString(shiftedrightgrowthcurve2dribbonstartshape.GongMarshallField(stage, "TopStartX"))
		initializerStatements.WriteString(shiftedrightgrowthcurve2dribbonstartshape.GongMarshallField(stage, "TopStartY"))
		initializerStatements.WriteString(shiftedrightgrowthcurve2dribbonstartshape.GongMarshallField(stage, "TopEndX"))
		initializerStatements.WriteString(shiftedrightgrowthcurve2dribbonstartshape.GongMarshallField(stage, "TopEndY"))
		initializerStatements.WriteString(shiftedrightgrowthcurve2dribbonstartshape.GongMarshallField(stage, "TopRadiusX"))
		initializerStatements.WriteString(shiftedrightgrowthcurve2dribbonstartshape.GongMarshallField(stage, "TopRadiusY"))
		initializerStatements.WriteString(shiftedrightgrowthcurve2dribbonstartshape.GongMarshallField(stage, "TopXAxisRotation"))
		initializerStatements.WriteString(shiftedrightgrowthcurve2dribbonstartshape.GongMarshallField(stage, "TopLargeArcFlag"))
		initializerStatements.WriteString(shiftedrightgrowthcurve2dribbonstartshape.GongMarshallField(stage, "TopSweepFlag"))
	}
	initRes = initializerStatements.String()
	ptrRes = pointersInitializesStatements.String()
	return
}
func (stackgrowthcurve2dendhalfwayarcshape *StackGrowthCurve2DEndHalfwayArcShape) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(stackgrowthcurve2dendhalfwayarcshape.GongMarshallField(stage, "Name"))
		initializerStatements.WriteString(stackgrowthcurve2dendhalfwayarcshape.GongMarshallField(stage, "StartX"))
		initializerStatements.WriteString(stackgrowthcurve2dendhalfwayarcshape.GongMarshallField(stage, "StartY"))
		initializerStatements.WriteString(stackgrowthcurve2dendhalfwayarcshape.GongMarshallField(stage, "EndX"))
		initializerStatements.WriteString(stackgrowthcurve2dendhalfwayarcshape.GongMarshallField(stage, "EndY"))
		initializerStatements.WriteString(stackgrowthcurve2dendhalfwayarcshape.GongMarshallField(stage, "RadiusX"))
		initializerStatements.WriteString(stackgrowthcurve2dendhalfwayarcshape.GongMarshallField(stage, "RadiusY"))
		initializerStatements.WriteString(stackgrowthcurve2dendhalfwayarcshape.GongMarshallField(stage, "XAxisRotation"))
		initializerStatements.WriteString(stackgrowthcurve2dendhalfwayarcshape.GongMarshallField(stage, "LargeArcFlag"))
		initializerStatements.WriteString(stackgrowthcurve2dendhalfwayarcshape.GongMarshallField(stage, "SweepFlag"))
	}
	initRes = initializerStatements.String()
	ptrRes = pointersInitializesStatements.String()
	return
}
func (stackgrowthcurve2dribbonendshape *StackGrowthCurve2DRibbonEndShape) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(stackgrowthcurve2dribbonendshape.GongMarshallField(stage, "Name"))
		initializerStatements.WriteString(stackgrowthcurve2dribbonendshape.GongMarshallField(stage, "BottomStartX"))
		initializerStatements.WriteString(stackgrowthcurve2dribbonendshape.GongMarshallField(stage, "BottomStartY"))
		initializerStatements.WriteString(stackgrowthcurve2dribbonendshape.GongMarshallField(stage, "BottomEndX"))
		initializerStatements.WriteString(stackgrowthcurve2dribbonendshape.GongMarshallField(stage, "BottomEndY"))
		initializerStatements.WriteString(stackgrowthcurve2dribbonendshape.GongMarshallField(stage, "BottomRadiusX"))
		initializerStatements.WriteString(stackgrowthcurve2dribbonendshape.GongMarshallField(stage, "BottomRadiusY"))
		initializerStatements.WriteString(stackgrowthcurve2dribbonendshape.GongMarshallField(stage, "BottomXAxisRotation"))
		initializerStatements.WriteString(stackgrowthcurve2dribbonendshape.GongMarshallField(stage, "BottomLargeArcFlag"))
		initializerStatements.WriteString(stackgrowthcurve2dribbonendshape.GongMarshallField(stage, "BottomSweepFlag"))
		initializerStatements.WriteString(stackgrowthcurve2dribbonendshape.GongMarshallField(stage, "TopStartX"))
		initializerStatements.WriteString(stackgrowthcurve2dribbonendshape.GongMarshallField(stage, "TopStartY"))
		initializerStatements.WriteString(stackgrowthcurve2dribbonendshape.GongMarshallField(stage, "TopEndX"))
		initializerStatements.WriteString(stackgrowthcurve2dribbonendshape.GongMarshallField(stage, "TopEndY"))
		initializerStatements.WriteString(stackgrowthcurve2dribbonendshape.GongMarshallField(stage, "TopRadiusX"))
		initializerStatements.WriteString(stackgrowthcurve2dribbonendshape.GongMarshallField(stage, "TopRadiusY"))
		initializerStatements.WriteString(stackgrowthcurve2dribbonendshape.GongMarshallField(stage, "TopXAxisRotation"))
		initializerStatements.WriteString(stackgrowthcurve2dribbonendshape.GongMarshallField(stage, "TopLargeArcFlag"))
		initializerStatements.WriteString(stackgrowthcurve2dribbonendshape.GongMarshallField(stage, "TopSweepFlag"))
	}
	initRes = initializerStatements.String()
	ptrRes = pointersInitializesStatements.String()
	return
}
func (stackgrowthcurve2dribbonstartshape *StackGrowthCurve2DRibbonStartShape) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(stackgrowthcurve2dribbonstartshape.GongMarshallField(stage, "Name"))
		initializerStatements.WriteString(stackgrowthcurve2dribbonstartshape.GongMarshallField(stage, "BottomStartX"))
		initializerStatements.WriteString(stackgrowthcurve2dribbonstartshape.GongMarshallField(stage, "BottomStartY"))
		initializerStatements.WriteString(stackgrowthcurve2dribbonstartshape.GongMarshallField(stage, "BottomEndX"))
		initializerStatements.WriteString(stackgrowthcurve2dribbonstartshape.GongMarshallField(stage, "BottomEndY"))
		initializerStatements.WriteString(stackgrowthcurve2dribbonstartshape.GongMarshallField(stage, "BottomRadiusX"))
		initializerStatements.WriteString(stackgrowthcurve2dribbonstartshape.GongMarshallField(stage, "BottomRadiusY"))
		initializerStatements.WriteString(stackgrowthcurve2dribbonstartshape.GongMarshallField(stage, "BottomXAxisRotation"))
		initializerStatements.WriteString(stackgrowthcurve2dribbonstartshape.GongMarshallField(stage, "BottomLargeArcFlag"))
		initializerStatements.WriteString(stackgrowthcurve2dribbonstartshape.GongMarshallField(stage, "BottomSweepFlag"))
		initializerStatements.WriteString(stackgrowthcurve2dribbonstartshape.GongMarshallField(stage, "TopStartX"))
		initializerStatements.WriteString(stackgrowthcurve2dribbonstartshape.GongMarshallField(stage, "TopStartY"))
		initializerStatements.WriteString(stackgrowthcurve2dribbonstartshape.GongMarshallField(stage, "TopEndX"))
		initializerStatements.WriteString(stackgrowthcurve2dribbonstartshape.GongMarshallField(stage, "TopEndY"))
		initializerStatements.WriteString(stackgrowthcurve2dribbonstartshape.GongMarshallField(stage, "TopRadiusX"))
		initializerStatements.WriteString(stackgrowthcurve2dribbonstartshape.GongMarshallField(stage, "TopRadiusY"))
		initializerStatements.WriteString(stackgrowthcurve2dribbonstartshape.GongMarshallField(stage, "TopXAxisRotation"))
		initializerStatements.WriteString(stackgrowthcurve2dribbonstartshape.GongMarshallField(stage, "TopLargeArcFlag"))
		initializerStatements.WriteString(stackgrowthcurve2dribbonstartshape.GongMarshallField(stage, "TopSweepFlag"))
	}
	initRes = initializerStatements.String()
	ptrRes = pointersInitializesStatements.String()
	return
}
func (stackgrowthcurve2dstarthalfwayarcshape *StackGrowthCurve2DStartHalfwayArcShape) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(stackgrowthcurve2dstarthalfwayarcshape.GongMarshallField(stage, "Name"))
		initializerStatements.WriteString(stackgrowthcurve2dstarthalfwayarcshape.GongMarshallField(stage, "StartX"))
		initializerStatements.WriteString(stackgrowthcurve2dstarthalfwayarcshape.GongMarshallField(stage, "StartY"))
		initializerStatements.WriteString(stackgrowthcurve2dstarthalfwayarcshape.GongMarshallField(stage, "EndX"))
		initializerStatements.WriteString(stackgrowthcurve2dstarthalfwayarcshape.GongMarshallField(stage, "EndY"))
		initializerStatements.WriteString(stackgrowthcurve2dstarthalfwayarcshape.GongMarshallField(stage, "RadiusX"))
		initializerStatements.WriteString(stackgrowthcurve2dstarthalfwayarcshape.GongMarshallField(stage, "RadiusY"))
		initializerStatements.WriteString(stackgrowthcurve2dstarthalfwayarcshape.GongMarshallField(stage, "XAxisRotation"))
		initializerStatements.WriteString(stackgrowthcurve2dstarthalfwayarcshape.GongMarshallField(stage, "LargeArcFlag"))
		initializerStatements.WriteString(stackgrowthcurve2dstarthalfwayarcshape.GongMarshallField(stage, "SweepFlag"))
	}
	initRes = initializerStatements.String()
	ptrRes = pointersInitializesStatements.String()
	return
}
func (stackofgrowthcurve2d *StackOfGrowthCurve2D) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(stackofgrowthcurve2d.GongMarshallField(stage, "Name"))
	}
	initRes = initializerStatements.String()
	ptrRes = pointersInitializesStatements.String()
	return
}
func (stackofgrowthcurve2dbygrowthvector *StackOfGrowthCurve2DByGrowthVector) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(stackofgrowthcurve2dbygrowthvector.GongMarshallField(stage, "Name"))
	}
	initRes = initializerStatements.String()
	ptrRes = pointersInitializesStatements.String()
	return
}
func (stackofgrowthcurve2dribbon *StackOfGrowthCurve2DRibbon) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(stackofgrowthcurve2dribbon.GongMarshallField(stage, "Name"))
	}
	initRes = initializerStatements.String()
	ptrRes = pointersInitializesStatements.String()
	return
}
func (stackofpartiallyrotatedtorusshape *StackOfPartiallyRotatedTorusShape) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(stackofpartiallyrotatedtorusshape.GongMarshallField(stage, "Name"))
	}
	initRes = initializerStatements.String()
	ptrRes = pointersInitializesStatements.String()
	return
}
func (stackofrotatedcarvedoutvasetrapezeringsshape *StackOfRotatedCarvedOutVaseTrapezeRingsShape) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(stackofrotatedcarvedoutvasetrapezeringsshape.GongMarshallField(stage, "Name"))
	}
	initRes = initializerStatements.String()
	ptrRes = pointersInitializesStatements.String()
	return
}
func (stackofrotatedgrowthcurve2d *StackOfRotatedGrowthCurve2D) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(stackofrotatedgrowthcurve2d.GongMarshallField(stage, "Name"))
	}
	initRes = initializerStatements.String()
	ptrRes = pointersInitializesStatements.String()
	return
}
func (stackofrotatedgrowthcurve2dribbon *StackOfRotatedGrowthCurve2DRibbon) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(stackofrotatedgrowthcurve2dribbon.GongMarshallField(stage, "Name"))
	}
	initRes = initializerStatements.String()
	ptrRes = pointersInitializesStatements.String()
	return
}
func (stackrotatedgrowthcurve2dendarcshape *StackRotatedGrowthCurve2DEndArcShape) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(stackrotatedgrowthcurve2dendarcshape.GongMarshallField(stage, "Name"))
		initializerStatements.WriteString(stackrotatedgrowthcurve2dendarcshape.GongMarshallField(stage, "StartX"))
		initializerStatements.WriteString(stackrotatedgrowthcurve2dendarcshape.GongMarshallField(stage, "StartY"))
		initializerStatements.WriteString(stackrotatedgrowthcurve2dendarcshape.GongMarshallField(stage, "EndX"))
		initializerStatements.WriteString(stackrotatedgrowthcurve2dendarcshape.GongMarshallField(stage, "EndY"))
		initializerStatements.WriteString(stackrotatedgrowthcurve2dendarcshape.GongMarshallField(stage, "XAxisRotation"))
		initializerStatements.WriteString(stackrotatedgrowthcurve2dendarcshape.GongMarshallField(stage, "LargeArcFlag"))
		initializerStatements.WriteString(stackrotatedgrowthcurve2dendarcshape.GongMarshallField(stage, "SweepFlag"))
		initializerStatements.WriteString(stackrotatedgrowthcurve2dendarcshape.GongMarshallField(stage, "RadiusX"))
		initializerStatements.WriteString(stackrotatedgrowthcurve2dendarcshape.GongMarshallField(stage, "RadiusY"))
	}
	initRes = initializerStatements.String()
	ptrRes = pointersInitializesStatements.String()
	return
}
func (stackrotatedgrowthcurve2dribbonendshape *StackRotatedGrowthCurve2DRibbonEndShape) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(stackrotatedgrowthcurve2dribbonendshape.GongMarshallField(stage, "Name"))
		initializerStatements.WriteString(stackrotatedgrowthcurve2dribbonendshape.GongMarshallField(stage, "BottomStartX"))
		initializerStatements.WriteString(stackrotatedgrowthcurve2dribbonendshape.GongMarshallField(stage, "BottomStartY"))
		initializerStatements.WriteString(stackrotatedgrowthcurve2dribbonendshape.GongMarshallField(stage, "BottomEndX"))
		initializerStatements.WriteString(stackrotatedgrowthcurve2dribbonendshape.GongMarshallField(stage, "BottomEndY"))
		initializerStatements.WriteString(stackrotatedgrowthcurve2dribbonendshape.GongMarshallField(stage, "BottomRadiusX"))
		initializerStatements.WriteString(stackrotatedgrowthcurve2dribbonendshape.GongMarshallField(stage, "BottomRadiusY"))
		initializerStatements.WriteString(stackrotatedgrowthcurve2dribbonendshape.GongMarshallField(stage, "BottomXAxisRotation"))
		initializerStatements.WriteString(stackrotatedgrowthcurve2dribbonendshape.GongMarshallField(stage, "BottomLargeArcFlag"))
		initializerStatements.WriteString(stackrotatedgrowthcurve2dribbonendshape.GongMarshallField(stage, "BottomSweepFlag"))
		initializerStatements.WriteString(stackrotatedgrowthcurve2dribbonendshape.GongMarshallField(stage, "TopStartX"))
		initializerStatements.WriteString(stackrotatedgrowthcurve2dribbonendshape.GongMarshallField(stage, "TopStartY"))
		initializerStatements.WriteString(stackrotatedgrowthcurve2dribbonendshape.GongMarshallField(stage, "TopEndX"))
		initializerStatements.WriteString(stackrotatedgrowthcurve2dribbonendshape.GongMarshallField(stage, "TopEndY"))
		initializerStatements.WriteString(stackrotatedgrowthcurve2dribbonendshape.GongMarshallField(stage, "TopRadiusX"))
		initializerStatements.WriteString(stackrotatedgrowthcurve2dribbonendshape.GongMarshallField(stage, "TopRadiusY"))
		initializerStatements.WriteString(stackrotatedgrowthcurve2dribbonendshape.GongMarshallField(stage, "TopXAxisRotation"))
		initializerStatements.WriteString(stackrotatedgrowthcurve2dribbonendshape.GongMarshallField(stage, "TopLargeArcFlag"))
		initializerStatements.WriteString(stackrotatedgrowthcurve2dribbonendshape.GongMarshallField(stage, "TopSweepFlag"))
	}
	initRes = initializerStatements.String()
	ptrRes = pointersInitializesStatements.String()
	return
}
func (stackrotatedgrowthcurve2dribbonstartshape *StackRotatedGrowthCurve2DRibbonStartShape) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(stackrotatedgrowthcurve2dribbonstartshape.GongMarshallField(stage, "Name"))
		initializerStatements.WriteString(stackrotatedgrowthcurve2dribbonstartshape.GongMarshallField(stage, "BottomStartX"))
		initializerStatements.WriteString(stackrotatedgrowthcurve2dribbonstartshape.GongMarshallField(stage, "BottomStartY"))
		initializerStatements.WriteString(stackrotatedgrowthcurve2dribbonstartshape.GongMarshallField(stage, "BottomEndX"))
		initializerStatements.WriteString(stackrotatedgrowthcurve2dribbonstartshape.GongMarshallField(stage, "BottomEndY"))
		initializerStatements.WriteString(stackrotatedgrowthcurve2dribbonstartshape.GongMarshallField(stage, "BottomRadiusX"))
		initializerStatements.WriteString(stackrotatedgrowthcurve2dribbonstartshape.GongMarshallField(stage, "BottomRadiusY"))
		initializerStatements.WriteString(stackrotatedgrowthcurve2dribbonstartshape.GongMarshallField(stage, "BottomXAxisRotation"))
		initializerStatements.WriteString(stackrotatedgrowthcurve2dribbonstartshape.GongMarshallField(stage, "BottomLargeArcFlag"))
		initializerStatements.WriteString(stackrotatedgrowthcurve2dribbonstartshape.GongMarshallField(stage, "BottomSweepFlag"))
		initializerStatements.WriteString(stackrotatedgrowthcurve2dribbonstartshape.GongMarshallField(stage, "TopStartX"))
		initializerStatements.WriteString(stackrotatedgrowthcurve2dribbonstartshape.GongMarshallField(stage, "TopStartY"))
		initializerStatements.WriteString(stackrotatedgrowthcurve2dribbonstartshape.GongMarshallField(stage, "TopEndX"))
		initializerStatements.WriteString(stackrotatedgrowthcurve2dribbonstartshape.GongMarshallField(stage, "TopEndY"))
		initializerStatements.WriteString(stackrotatedgrowthcurve2dribbonstartshape.GongMarshallField(stage, "TopRadiusX"))
		initializerStatements.WriteString(stackrotatedgrowthcurve2dribbonstartshape.GongMarshallField(stage, "TopRadiusY"))
		initializerStatements.WriteString(stackrotatedgrowthcurve2dribbonstartshape.GongMarshallField(stage, "TopXAxisRotation"))
		initializerStatements.WriteString(stackrotatedgrowthcurve2dribbonstartshape.GongMarshallField(stage, "TopLargeArcFlag"))
		initializerStatements.WriteString(stackrotatedgrowthcurve2dribbonstartshape.GongMarshallField(stage, "TopSweepFlag"))
	}
	initRes = initializerStatements.String()
	ptrRes = pointersInitializesStatements.String()
	return
}
func (stackrotatedgrowthcurve2dstartarcshape *StackRotatedGrowthCurve2DStartArcShape) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(stackrotatedgrowthcurve2dstartarcshape.GongMarshallField(stage, "Name"))
		initializerStatements.WriteString(stackrotatedgrowthcurve2dstartarcshape.GongMarshallField(stage, "StartX"))
		initializerStatements.WriteString(stackrotatedgrowthcurve2dstartarcshape.GongMarshallField(stage, "StartY"))
		initializerStatements.WriteString(stackrotatedgrowthcurve2dstartarcshape.GongMarshallField(stage, "EndX"))
		initializerStatements.WriteString(stackrotatedgrowthcurve2dstartarcshape.GongMarshallField(stage, "EndY"))
		initializerStatements.WriteString(stackrotatedgrowthcurve2dstartarcshape.GongMarshallField(stage, "XAxisRotation"))
		initializerStatements.WriteString(stackrotatedgrowthcurve2dstartarcshape.GongMarshallField(stage, "LargeArcFlag"))
		initializerStatements.WriteString(stackrotatedgrowthcurve2dstartarcshape.GongMarshallField(stage, "SweepFlag"))
		initializerStatements.WriteString(stackrotatedgrowthcurve2dstartarcshape.GongMarshallField(stage, "RadiusX"))
		initializerStatements.WriteString(stackrotatedgrowthcurve2dstartarcshape.GongMarshallField(stage, "RadiusY"))
	}
	initRes = initializerStatements.String()
	ptrRes = pointersInitializesStatements.String()
	return
}
func (startarcshape *StartArcShape) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(startarcshape.GongMarshallField(stage, "Name"))
		initializerStatements.WriteString(startarcshape.GongMarshallField(stage, "StartX"))
		initializerStatements.WriteString(startarcshape.GongMarshallField(stage, "StartY"))
		initializerStatements.WriteString(startarcshape.GongMarshallField(stage, "EndX"))
		initializerStatements.WriteString(startarcshape.GongMarshallField(stage, "EndY"))
		initializerStatements.WriteString(startarcshape.GongMarshallField(stage, "XAxisRotation"))
		initializerStatements.WriteString(startarcshape.GongMarshallField(stage, "LargeArcFlag"))
		initializerStatements.WriteString(startarcshape.GongMarshallField(stage, "SweepFlag"))
		initializerStatements.WriteString(startarcshape.GongMarshallField(stage, "RadiusX"))
		initializerStatements.WriteString(startarcshape.GongMarshallField(stage, "RadiusY"))
	}
	initRes = initializerStatements.String()
	ptrRes = pointersInitializesStatements.String()
	return
}
func (startarcshapegrid *StartArcShapeGrid) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(startarcshapegrid.GongMarshallField(stage, "Name"))
	}
	initRes = initializerStatements.String()
	ptrRes = pointersInitializesStatements.String()
	return
}
func (starthalfwayarcshape *StartHalfwayArcShape) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(starthalfwayarcshape.GongMarshallField(stage, "Name"))
		initializerStatements.WriteString(starthalfwayarcshape.GongMarshallField(stage, "StartX"))
		initializerStatements.WriteString(starthalfwayarcshape.GongMarshallField(stage, "StartY"))
		initializerStatements.WriteString(starthalfwayarcshape.GongMarshallField(stage, "EndX"))
		initializerStatements.WriteString(starthalfwayarcshape.GongMarshallField(stage, "EndY"))
		initializerStatements.WriteString(starthalfwayarcshape.GongMarshallField(stage, "RadiusX"))
		initializerStatements.WriteString(starthalfwayarcshape.GongMarshallField(stage, "RadiusY"))
		initializerStatements.WriteString(starthalfwayarcshape.GongMarshallField(stage, "XAxisRotation"))
		initializerStatements.WriteString(starthalfwayarcshape.GongMarshallField(stage, "LargeArcFlag"))
		initializerStatements.WriteString(starthalfwayarcshape.GongMarshallField(stage, "SweepFlag"))
	}
	initRes = initializerStatements.String()
	ptrRes = pointersInitializesStatements.String()
	return
}
func (starthalfwayarcshapegrid *StartHalfwayArcShapeGrid) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(starthalfwayarcshapegrid.GongMarshallField(stage, "Name"))
	}
	initRes = initializerStatements.String()
	ptrRes = pointersInitializesStatements.String()
	return
}
func (stemcylinder3dshape *StemCylinder3DShape) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(stemcylinder3dshape.GongMarshallField(stage, "Name"))
		initializerStatements.WriteString(stemcylinder3dshape.GongMarshallField(stage, "Transparency"))
	}
	initRes = initializerStatements.String()
	ptrRes = pointersInitializesStatements.String()
	return
}
func (stool2ddiagram *Stool2DDiagram) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(stool2ddiagram.GongMarshallField(stage, "Name"))
		initializerStatements.WriteString(stool2ddiagram.GongMarshallField(stage, "Zoom"))
		initializerStatements.WriteString(stool2ddiagram.GongMarshallField(stage, "IsHiddenAxesShape"))
		initializerStatements.WriteString(stool2ddiagram.GongMarshallField(stage, "IsChecked"))
		initializerStatements.WriteString(stool2ddiagram.GongMarshallField(stage, "ComputedPrefix"))
		initializerStatements.WriteString(stool2ddiagram.GongMarshallField(stage, "IsExpanded"))
	}
	initRes = initializerStatements.String()
	ptrRes = pointersInitializesStatements.String()
	return
}
func (stool3ddiagram *Stool3DDiagram) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(stool3ddiagram.GongMarshallField(stage, "Name"))
		initializerStatements.WriteString(stool3ddiagram.GongMarshallField(stage, "IsHiddenSeatTopCurveShape"))
		initializerStatements.WriteString(stool3ddiagram.GongMarshallField(stage, "IsHiddenRotatedSeatTopCurveShape"))
		initializerStatements.WriteString(stool3ddiagram.GongMarshallField(stage, "IsHiddenSeatBottomCurveShape"))
		initializerStatements.WriteString(stool3ddiagram.GongMarshallField(stage, "IsHiddenRotatedSeatBottomCurveShape"))
		initializerStatements.WriteString(stool3ddiagram.GongMarshallField(stage, "IsHiddenTorus3DShape"))
		initializerStatements.WriteString(stool3ddiagram.GongMarshallField(stage, "IsHiddenRotatedTorusShape"))
		initializerStatements.WriteString(stool3ddiagram.GongMarshallField(stage, "IsHiddenSampledPoints3DShape"))
		pointersInitializesStatements.WriteString(stool3ddiagram.GongMarshallField(stage, "SampledPoints3DShape"))
		initializerStatements.WriteString(stool3ddiagram.GongMarshallField(stage, "IsHiddenRotatedSampledPoints3DShape"))
		initializerStatements.WriteString(stool3ddiagram.GongMarshallField(stage, "IsHiddenEyeSampledPoints3DShape"))
		initializerStatements.WriteString(stool3ddiagram.GongMarshallField(stage, "IsHiddenEyeCornersSampledPoints3DShape"))
		initializerStatements.WriteString(stool3ddiagram.GongMarshallField(stage, "IsHiddenEye3DShape"))
		initializerStatements.WriteString(stool3ddiagram.GongMarshallField(stage, "IsHiddenEyeSeatBottomCurveShape"))
		initializerStatements.WriteString(stool3ddiagram.GongMarshallField(stage, "IsHiddenEyeStoolBottomCurveShape"))
		initializerStatements.WriteString(stool3ddiagram.GongMarshallField(stage, "IsHiddenSeat3DShape"))
		initializerStatements.WriteString(stool3ddiagram.GongMarshallField(stage, "IsHiddenEyeVolume3DShape"))
		initializerStatements.WriteString(stool3ddiagram.GongMarshallField(stage, "IsHiddenSeatAndLegs3DShape"))
		initializerStatements.WriteString(stool3ddiagram.GongMarshallField(stage, "IsHiddenRotatedSeatAndLegs3DShape"))
		initializerStatements.WriteString(stool3ddiagram.GongMarshallField(stage, "IsHiddenTiledFloor3DShape"))
		pointersInitializesStatements.WriteString(stool3ddiagram.GongMarshallField(stage, "Rendered3DShape"))
		initializerStatements.WriteString(stool3ddiagram.GongMarshallField(stage, "IsChecked"))
		initializerStatements.WriteString(stool3ddiagram.GongMarshallField(stage, "ComputedPrefix"))
		initializerStatements.WriteString(stool3ddiagram.GongMarshallField(stage, "IsExpanded"))
	}
	initRes = initializerStatements.String()
	ptrRes = pointersInitializesStatements.String()
	return
}
func (tiledfloor3dshape *TiledFloor3DShape) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(tiledfloor3dshape.GongMarshallField(stage, "Name"))
	}
	initRes = initializerStatements.String()
	ptrRes = pointersInitializesStatements.String()
	return
}
func (topcurveplane1shape *TopCurvePlane1Shape) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(topcurveplane1shape.GongMarshallField(stage, "Name"))
	}
	initRes = initializerStatements.String()
	ptrRes = pointersInitializesStatements.String()
	return
}
func (topcurveplane2shape *TopCurvePlane2Shape) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(topcurveplane2shape.GongMarshallField(stage, "Name"))
	}
	initRes = initializerStatements.String()
	ptrRes = pointersInitializesStatements.String()
	return
}
func (topendarcshape *TopEndArcShape) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(topendarcshape.GongMarshallField(stage, "Name"))
		initializerStatements.WriteString(topendarcshape.GongMarshallField(stage, "StartX"))
		initializerStatements.WriteString(topendarcshape.GongMarshallField(stage, "StartY"))
		initializerStatements.WriteString(topendarcshape.GongMarshallField(stage, "EndX"))
		initializerStatements.WriteString(topendarcshape.GongMarshallField(stage, "EndY"))
		initializerStatements.WriteString(topendarcshape.GongMarshallField(stage, "XAxisRotation"))
		initializerStatements.WriteString(topendarcshape.GongMarshallField(stage, "LargeArcFlag"))
		initializerStatements.WriteString(topendarcshape.GongMarshallField(stage, "SweepFlag"))
		initializerStatements.WriteString(topendarcshape.GongMarshallField(stage, "RadiusX"))
		initializerStatements.WriteString(topendarcshape.GongMarshallField(stage, "RadiusY"))
	}
	initRes = initializerStatements.String()
	ptrRes = pointersInitializesStatements.String()
	return
}
func (topendarcshapegrid *TopEndArcShapeGrid) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(topendarcshapegrid.GongMarshallField(stage, "Name"))
	}
	initRes = initializerStatements.String()
	ptrRes = pointersInitializesStatements.String()
	return
}
func (topendhalfwayarcshape *TopEndHalfwayArcShape) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(topendhalfwayarcshape.GongMarshallField(stage, "Name"))
		initializerStatements.WriteString(topendhalfwayarcshape.GongMarshallField(stage, "StartX"))
		initializerStatements.WriteString(topendhalfwayarcshape.GongMarshallField(stage, "StartY"))
		initializerStatements.WriteString(topendhalfwayarcshape.GongMarshallField(stage, "EndX"))
		initializerStatements.WriteString(topendhalfwayarcshape.GongMarshallField(stage, "EndY"))
		initializerStatements.WriteString(topendhalfwayarcshape.GongMarshallField(stage, "RadiusX"))
		initializerStatements.WriteString(topendhalfwayarcshape.GongMarshallField(stage, "RadiusY"))
		initializerStatements.WriteString(topendhalfwayarcshape.GongMarshallField(stage, "XAxisRotation"))
		initializerStatements.WriteString(topendhalfwayarcshape.GongMarshallField(stage, "LargeArcFlag"))
		initializerStatements.WriteString(topendhalfwayarcshape.GongMarshallField(stage, "SweepFlag"))
	}
	initRes = initializerStatements.String()
	ptrRes = pointersInitializesStatements.String()
	return
}
func (topendhalfwayarcshapegrid *TopEndHalfwayArcShapeGrid) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(topendhalfwayarcshapegrid.GongMarshallField(stage, "Name"))
	}
	initRes = initializerStatements.String()
	ptrRes = pointersInitializesStatements.String()
	return
}
func (topgrowthcurve2d *TopGrowthCurve2D) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(topgrowthcurve2d.GongMarshallField(stage, "Name"))
	}
	initRes = initializerStatements.String()
	ptrRes = pointersInitializesStatements.String()
	return
}
func (topmidarcvectorshape *TopMidArcVectorShape) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(topmidarcvectorshape.GongMarshallField(stage, "Name"))
		initializerStatements.WriteString(topmidarcvectorshape.GongMarshallField(stage, "StartX"))
		initializerStatements.WriteString(topmidarcvectorshape.GongMarshallField(stage, "StartY"))
		initializerStatements.WriteString(topmidarcvectorshape.GongMarshallField(stage, "EndX"))
		initializerStatements.WriteString(topmidarcvectorshape.GongMarshallField(stage, "EndY"))
	}
	initRes = initializerStatements.String()
	ptrRes = pointersInitializesStatements.String()
	return
}
func (topmidarcvectorshapegrid *TopMidArcVectorShapeGrid) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(topmidarcvectorshapegrid.GongMarshallField(stage, "Name"))
	}
	initRes = initializerStatements.String()
	ptrRes = pointersInitializesStatements.String()
	return
}
func (topstackgrowthcurve2dendhalfwayarcshape *TopStackGrowthCurve2DEndHalfwayArcShape) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(topstackgrowthcurve2dendhalfwayarcshape.GongMarshallField(stage, "Name"))
		initializerStatements.WriteString(topstackgrowthcurve2dendhalfwayarcshape.GongMarshallField(stage, "StartX"))
		initializerStatements.WriteString(topstackgrowthcurve2dendhalfwayarcshape.GongMarshallField(stage, "StartY"))
		initializerStatements.WriteString(topstackgrowthcurve2dendhalfwayarcshape.GongMarshallField(stage, "EndX"))
		initializerStatements.WriteString(topstackgrowthcurve2dendhalfwayarcshape.GongMarshallField(stage, "EndY"))
		initializerStatements.WriteString(topstackgrowthcurve2dendhalfwayarcshape.GongMarshallField(stage, "RadiusX"))
		initializerStatements.WriteString(topstackgrowthcurve2dendhalfwayarcshape.GongMarshallField(stage, "RadiusY"))
		initializerStatements.WriteString(topstackgrowthcurve2dendhalfwayarcshape.GongMarshallField(stage, "XAxisRotation"))
		initializerStatements.WriteString(topstackgrowthcurve2dendhalfwayarcshape.GongMarshallField(stage, "LargeArcFlag"))
		initializerStatements.WriteString(topstackgrowthcurve2dendhalfwayarcshape.GongMarshallField(stage, "SweepFlag"))
	}
	initRes = initializerStatements.String()
	ptrRes = pointersInitializesStatements.String()
	return
}
func (topstackgrowthcurve2dstarthalfwayarcshape *TopStackGrowthCurve2DStartHalfwayArcShape) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(topstackgrowthcurve2dstarthalfwayarcshape.GongMarshallField(stage, "Name"))
		initializerStatements.WriteString(topstackgrowthcurve2dstarthalfwayarcshape.GongMarshallField(stage, "StartX"))
		initializerStatements.WriteString(topstackgrowthcurve2dstarthalfwayarcshape.GongMarshallField(stage, "StartY"))
		initializerStatements.WriteString(topstackgrowthcurve2dstarthalfwayarcshape.GongMarshallField(stage, "EndX"))
		initializerStatements.WriteString(topstackgrowthcurve2dstarthalfwayarcshape.GongMarshallField(stage, "EndY"))
		initializerStatements.WriteString(topstackgrowthcurve2dstarthalfwayarcshape.GongMarshallField(stage, "RadiusX"))
		initializerStatements.WriteString(topstackgrowthcurve2dstarthalfwayarcshape.GongMarshallField(stage, "RadiusY"))
		initializerStatements.WriteString(topstackgrowthcurve2dstarthalfwayarcshape.GongMarshallField(stage, "XAxisRotation"))
		initializerStatements.WriteString(topstackgrowthcurve2dstarthalfwayarcshape.GongMarshallField(stage, "LargeArcFlag"))
		initializerStatements.WriteString(topstackgrowthcurve2dstarthalfwayarcshape.GongMarshallField(stage, "SweepFlag"))
	}
	initRes = initializerStatements.String()
	ptrRes = pointersInitializesStatements.String()
	return
}
func (topstackofgrowthcurve2d *TopStackOfGrowthCurve2D) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(topstackofgrowthcurve2d.GongMarshallField(stage, "Name"))
	}
	initRes = initializerStatements.String()
	ptrRes = pointersInitializesStatements.String()
	return
}
func (topstackofrotatedgrowthcurve2d *TopStackOfRotatedGrowthCurve2D) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(topstackofrotatedgrowthcurve2d.GongMarshallField(stage, "Name"))
	}
	initRes = initializerStatements.String()
	ptrRes = pointersInitializesStatements.String()
	return
}
func (topstackofrotatedgrowthcurve2dendarcshape *TopStackOfRotatedGrowthCurve2DEndArcShape) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(topstackofrotatedgrowthcurve2dendarcshape.GongMarshallField(stage, "Name"))
		initializerStatements.WriteString(topstackofrotatedgrowthcurve2dendarcshape.GongMarshallField(stage, "StartX"))
		initializerStatements.WriteString(topstackofrotatedgrowthcurve2dendarcshape.GongMarshallField(stage, "StartY"))
		initializerStatements.WriteString(topstackofrotatedgrowthcurve2dendarcshape.GongMarshallField(stage, "EndX"))
		initializerStatements.WriteString(topstackofrotatedgrowthcurve2dendarcshape.GongMarshallField(stage, "EndY"))
		initializerStatements.WriteString(topstackofrotatedgrowthcurve2dendarcshape.GongMarshallField(stage, "XAxisRotation"))
		initializerStatements.WriteString(topstackofrotatedgrowthcurve2dendarcshape.GongMarshallField(stage, "LargeArcFlag"))
		initializerStatements.WriteString(topstackofrotatedgrowthcurve2dendarcshape.GongMarshallField(stage, "SweepFlag"))
		initializerStatements.WriteString(topstackofrotatedgrowthcurve2dendarcshape.GongMarshallField(stage, "RadiusX"))
		initializerStatements.WriteString(topstackofrotatedgrowthcurve2dendarcshape.GongMarshallField(stage, "RadiusY"))
	}
	initRes = initializerStatements.String()
	ptrRes = pointersInitializesStatements.String()
	return
}
func (topstackofrotatedgrowthcurve2dstartarcshape *TopStackOfRotatedGrowthCurve2DStartArcShape) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(topstackofrotatedgrowthcurve2dstartarcshape.GongMarshallField(stage, "Name"))
		initializerStatements.WriteString(topstackofrotatedgrowthcurve2dstartarcshape.GongMarshallField(stage, "StartX"))
		initializerStatements.WriteString(topstackofrotatedgrowthcurve2dstartarcshape.GongMarshallField(stage, "StartY"))
		initializerStatements.WriteString(topstackofrotatedgrowthcurve2dstartarcshape.GongMarshallField(stage, "EndX"))
		initializerStatements.WriteString(topstackofrotatedgrowthcurve2dstartarcshape.GongMarshallField(stage, "EndY"))
		initializerStatements.WriteString(topstackofrotatedgrowthcurve2dstartarcshape.GongMarshallField(stage, "XAxisRotation"))
		initializerStatements.WriteString(topstackofrotatedgrowthcurve2dstartarcshape.GongMarshallField(stage, "LargeArcFlag"))
		initializerStatements.WriteString(topstackofrotatedgrowthcurve2dstartarcshape.GongMarshallField(stage, "SweepFlag"))
		initializerStatements.WriteString(topstackofrotatedgrowthcurve2dstartarcshape.GongMarshallField(stage, "RadiusX"))
		initializerStatements.WriteString(topstackofrotatedgrowthcurve2dstartarcshape.GongMarshallField(stage, "RadiusY"))
	}
	initRes = initializerStatements.String()
	ptrRes = pointersInitializesStatements.String()
	return
}
func (topstartarcshape *TopStartArcShape) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(topstartarcshape.GongMarshallField(stage, "Name"))
		initializerStatements.WriteString(topstartarcshape.GongMarshallField(stage, "StartX"))
		initializerStatements.WriteString(topstartarcshape.GongMarshallField(stage, "StartY"))
		initializerStatements.WriteString(topstartarcshape.GongMarshallField(stage, "EndX"))
		initializerStatements.WriteString(topstartarcshape.GongMarshallField(stage, "EndY"))
		initializerStatements.WriteString(topstartarcshape.GongMarshallField(stage, "XAxisRotation"))
		initializerStatements.WriteString(topstartarcshape.GongMarshallField(stage, "LargeArcFlag"))
		initializerStatements.WriteString(topstartarcshape.GongMarshallField(stage, "SweepFlag"))
		initializerStatements.WriteString(topstartarcshape.GongMarshallField(stage, "RadiusX"))
		initializerStatements.WriteString(topstartarcshape.GongMarshallField(stage, "RadiusY"))
	}
	initRes = initializerStatements.String()
	ptrRes = pointersInitializesStatements.String()
	return
}
func (topstartarcshapegrid *TopStartArcShapeGrid) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(topstartarcshapegrid.GongMarshallField(stage, "Name"))
	}
	initRes = initializerStatements.String()
	ptrRes = pointersInitializesStatements.String()
	return
}
func (topstarthalfwayarcshape *TopStartHalfwayArcShape) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(topstarthalfwayarcshape.GongMarshallField(stage, "Name"))
		initializerStatements.WriteString(topstarthalfwayarcshape.GongMarshallField(stage, "StartX"))
		initializerStatements.WriteString(topstarthalfwayarcshape.GongMarshallField(stage, "StartY"))
		initializerStatements.WriteString(topstarthalfwayarcshape.GongMarshallField(stage, "EndX"))
		initializerStatements.WriteString(topstarthalfwayarcshape.GongMarshallField(stage, "EndY"))
		initializerStatements.WriteString(topstarthalfwayarcshape.GongMarshallField(stage, "RadiusX"))
		initializerStatements.WriteString(topstarthalfwayarcshape.GongMarshallField(stage, "RadiusY"))
		initializerStatements.WriteString(topstarthalfwayarcshape.GongMarshallField(stage, "XAxisRotation"))
		initializerStatements.WriteString(topstarthalfwayarcshape.GongMarshallField(stage, "LargeArcFlag"))
		initializerStatements.WriteString(topstarthalfwayarcshape.GongMarshallField(stage, "SweepFlag"))
	}
	initRes = initializerStatements.String()
	ptrRes = pointersInitializesStatements.String()
	return
}
func (topstarthalfwayarcshapegrid *TopStartHalfwayArcShapeGrid) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(topstarthalfwayarcshapegrid.GongMarshallField(stage, "Name"))
	}
	initRes = initializerStatements.String()
	ptrRes = pointersInitializesStatements.String()
	return
}
func (torus3dshape *Torus3DShape) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(torus3dshape.GongMarshallField(stage, "Name"))
	}
	initRes = initializerStatements.String()
	ptrRes = pointersInitializesStatements.String()
	return
}
func (torusedge3dshape *TorusEdge3DShape) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(torusedge3dshape.GongMarshallField(stage, "Name"))
	}
	initRes = initializerStatements.String()
	ptrRes = pointersInitializesStatements.String()
	return
}
func (torusstackshape *TorusStackShape) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(torusstackshape.GongMarshallField(stage, "Name"))
	}
	initRes = initializerStatements.String()
	ptrRes = pointersInitializesStatements.String()
	return
}
func (tubevase3ddiagram *TubeVase3DDiagram) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(tubevase3ddiagram.GongMarshallField(stage, "Name"))
		initializerStatements.WriteString(tubevase3ddiagram.GongMarshallField(stage, "IsHiddenStackOfPartiallyRotatedGrowthCurve2DRibbon"))
		initializerStatements.WriteString(tubevase3ddiagram.GongMarshallField(stage, "IsHiddenTorusStackShape"))
		initializerStatements.WriteString(tubevase3ddiagram.GongMarshallField(stage, "IsHiddenVerticalTorusStackShape"))
		initializerStatements.WriteString(tubevase3ddiagram.GongMarshallField(stage, "IsHiddenPartiallyRotatedTorusShape"))
		initializerStatements.WriteString(tubevase3ddiagram.GongMarshallField(stage, "IsHiddenStackOfPartiallyRotatedTorusShape"))
		initializerStatements.WriteString(tubevase3ddiagram.GongMarshallField(stage, "IsHiddenPointsAndLines3DShape"))
		initializerStatements.WriteString(tubevase3ddiagram.GongMarshallField(stage, "IsHiddenKeyHole3DShape"))
		initializerStatements.WriteString(tubevase3ddiagram.GongMarshallField(stage, "IsHiddenKey3DShape"))
		initializerStatements.WriteString(tubevase3ddiagram.GongMarshallField(stage, "IsHiddenVolumeKey3DShape"))
		initializerStatements.WriteString(tubevase3ddiagram.GongMarshallField(stage, "IsHiddenTorusEdge3DShape"))
		initializerStatements.WriteString(tubevase3ddiagram.GongMarshallField(stage, "IsHiddenSampledPoints3DShape"))
		initializerStatements.WriteString(tubevase3ddiagram.GongMarshallField(stage, "IsHiddenOriginalPoints3DShape"))
		initializerStatements.WriteString(tubevase3ddiagram.GongMarshallField(stage, "IsHiddenAngle0Shape"))
		initializerStatements.WriteString(tubevase3ddiagram.GongMarshallField(stage, "IsHiddenTiledFloor3DShape"))
		initializerStatements.WriteString(tubevase3ddiagram.GongMarshallField(stage, "IsHiddenTopCurvePlane1Shape"))
		initializerStatements.WriteString(tubevase3ddiagram.GongMarshallField(stage, "IsHiddenBottomCurvePlane1Shape"))
		initializerStatements.WriteString(tubevase3ddiagram.GongMarshallField(stage, "IsHiddenTopCurvePlane2Shape"))
		initializerStatements.WriteString(tubevase3ddiagram.GongMarshallField(stage, "IsHiddenBottomCurvePlane2Shape"))
		initializerStatements.WriteString(tubevase3ddiagram.GongMarshallField(stage, "IsHiddenCarvedOutTopCurvePlane1Shape"))
		initializerStatements.WriteString(tubevase3ddiagram.GongMarshallField(stage, "IsHiddenCarvedOutBottomCurvePlane1Shape"))
		initializerStatements.WriteString(tubevase3ddiagram.GongMarshallField(stage, "IsHiddenVaseTrapezeRingShape"))
		initializerStatements.WriteString(tubevase3ddiagram.GongMarshallField(stage, "IsHiddenCarvedOutVaseTrapezeRingShape"))
		initializerStatements.WriteString(tubevase3ddiagram.GongMarshallField(stage, "IsHiddenStackOfRotatedCarvedOutVaseTrapezeRingsShape"))
		initializerStatements.WriteString(tubevase3ddiagram.GongMarshallField(stage, "IsHiddenVaseTrapezeBasePlateShape"))
		pointersInitializesStatements.WriteString(tubevase3ddiagram.GongMarshallField(stage, "Rendered3DShape"))
		pointersInitializesStatements.WriteString(tubevase3ddiagram.GongMarshallField(stage, "SampledPoints3DShape"))
		pointersInitializesStatements.WriteString(tubevase3ddiagram.GongMarshallField(stage, "OriginalPoints3DShape"))
		pointersInitializesStatements.WriteString(tubevase3ddiagram.GongMarshallField(stage, "Angle0Shape"))
		pointersInitializesStatements.WriteString(tubevase3ddiagram.GongMarshallField(stage, "TopCurvePlane1Shape"))
		pointersInitializesStatements.WriteString(tubevase3ddiagram.GongMarshallField(stage, "BottomCurvePlane1Shape"))
		pointersInitializesStatements.WriteString(tubevase3ddiagram.GongMarshallField(stage, "TopCurvePlane2Shape"))
		pointersInitializesStatements.WriteString(tubevase3ddiagram.GongMarshallField(stage, "BottomCurvePlane2Shape"))
		pointersInitializesStatements.WriteString(tubevase3ddiagram.GongMarshallField(stage, "CarvedOutTopCurvePlane1Shape"))
		pointersInitializesStatements.WriteString(tubevase3ddiagram.GongMarshallField(stage, "CarvedOutBottomCurvePlane1Shape"))
		pointersInitializesStatements.WriteString(tubevase3ddiagram.GongMarshallField(stage, "VaseTrapezeRingShape"))
		pointersInitializesStatements.WriteString(tubevase3ddiagram.GongMarshallField(stage, "CarvedOutVaseTrapezeRingShape"))
		pointersInitializesStatements.WriteString(tubevase3ddiagram.GongMarshallField(stage, "StackOfRotatedCarvedOutVaseTrapezeRingsShape"))
		pointersInitializesStatements.WriteString(tubevase3ddiagram.GongMarshallField(stage, "VaseTrapezeBasePlateShape"))
		initializerStatements.WriteString(tubevase3ddiagram.GongMarshallField(stage, "IsChecked"))
		initializerStatements.WriteString(tubevase3ddiagram.GongMarshallField(stage, "ComputedPrefix"))
		initializerStatements.WriteString(tubevase3ddiagram.GongMarshallField(stage, "IsExpanded"))
	}
	initRes = initializerStatements.String()
	ptrRes = pointersInitializesStatements.String()
	return
}
func (tubevaseabstract *TubeVaseAbstract) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(tubevaseabstract.GongMarshallField(stage, "Name"))
		initializerStatements.WriteString(tubevaseabstract.GongMarshallField(stage, "Z_Ribbon"))
		initializerStatements.WriteString(tubevaseabstract.GongMarshallField(stage, "RibbonVerticalScale"))
		initializerStatements.WriteString(tubevaseabstract.GongMarshallField(stage, "TopPlaneHeight"))
		initializerStatements.WriteString(tubevaseabstract.GongMarshallField(stage, "BottomPlaneHeight"))
		initializerStatements.WriteString(tubevaseabstract.GongMarshallField(stage, "BaseRingBottomPlaneHeight"))
		initializerStatements.WriteString(tubevaseabstract.GongMarshallField(stage, "ProjectionAngle"))
		initializerStatements.WriteString(tubevaseabstract.GongMarshallField(stage, "BasePlateHeight"))
		initializerStatements.WriteString(tubevaseabstract.GongMarshallField(stage, "CarvedOutTopRingsParameter"))
		initializerStatements.WriteString(tubevaseabstract.GongMarshallField(stage, "BulbousStartTangentMagnitude"))
		initializerStatements.WriteString(tubevaseabstract.GongMarshallField(stage, "BulbousEndAngle"))
		initializerStatements.WriteString(tubevaseabstract.GongMarshallField(stage, "BulbousEndTangentMagnitude"))
		initializerStatements.WriteString(tubevaseabstract.GongMarshallField(stage, "RelativeVerticalThickness"))
		initializerStatements.WriteString(tubevaseabstract.GongMarshallField(stage, "RelativeRadialThickness"))
		initializerStatements.WriteString(tubevaseabstract.GongMarshallField(stage, "RelativeCuttedStackFloorHeight"))
		initializerStatements.WriteString(tubevaseabstract.GongMarshallField(stage, "RelativeRotatedTorusSeparation"))
		initializerStatements.WriteString(tubevaseabstract.GongMarshallField(stage, "RotationRatio"))
		initializerStatements.WriteString(tubevaseabstract.GongMarshallField(stage, "RadialRepetitions"))
		initializerStatements.WriteString(tubevaseabstract.GongMarshallField(stage, "Transparency"))
		initializerStatements.WriteString(tubevaseabstract.GongMarshallField(stage, "HasAlternatingRingColors"))
		initializerStatements.WriteString(tubevaseabstract.GongMarshallField(stage, "RelativeTrajectoryOffsetX"))
		initializerStatements.WriteString(tubevaseabstract.GongMarshallField(stage, "RelativeTrajectoryOffsetY"))
		initializerStatements.WriteString(tubevaseabstract.GongMarshallField(stage, "NbStepP1P2"))
		initializerStatements.WriteString(tubevaseabstract.GongMarshallField(stage, "ChosenStep"))
		initializerStatements.WriteString(tubevaseabstract.GongMarshallField(stage, "RelativeHorizontalRingsHeight"))
		initializerStatements.WriteString(tubevaseabstract.GongMarshallField(stage, "OffsetKeyX"))
		initializerStatements.WriteString(tubevaseabstract.GongMarshallField(stage, "OffsetKeyY"))
		initializerStatements.WriteString(tubevaseabstract.GongMarshallField(stage, "HeightKey"))
		initializerStatements.WriteString(tubevaseabstract.GongMarshallField(stage, "WidthKey"))
		initializerStatements.WriteString(tubevaseabstract.GongMarshallField(stage, "RelativeKeySize"))
		initializerStatements.WriteString(tubevaseabstract.GongMarshallField(stage, "MovieNbFrames"))
	}
	initRes = initializerStatements.String()
	ptrRes = pointersInitializesStatements.String()
	return
}
func (vase2ddiagram *Vase2DDiagram) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(vase2ddiagram.GongMarshallField(stage, "Name"))
		initializerStatements.WriteString(vase2ddiagram.GongMarshallField(stage, "Zoom"))
		initializerStatements.WriteString(vase2ddiagram.GongMarshallField(stage, "IsVaseArcNodesExpanded"))
		initializerStatements.WriteString(vase2ddiagram.GongMarshallField(stage, "IsVaseClampingNodesExpanded"))
		initializerStatements.WriteString(vase2ddiagram.GongMarshallField(stage, "IsHiddenAxesShape"))
		initializerStatements.WriteString(vase2ddiagram.GongMarshallField(stage, "IsHiddenBottomStartArcShapeGrid"))
		initializerStatements.WriteString(vase2ddiagram.GongMarshallField(stage, "IsHiddenBottomEndArcShapeGrid"))
		initializerStatements.WriteString(vase2ddiagram.GongMarshallField(stage, "IsHiddenBottomStackOfGrowthCurve"))
		initializerStatements.WriteString(vase2ddiagram.GongMarshallField(stage, "IsHiddenShiftedLeftStackOfGrowthCurve"))
		initializerStatements.WriteString(vase2ddiagram.GongMarshallField(stage, "IsHiddenShiftedLeftStackOfNormalVector"))
		initializerStatements.WriteString(vase2ddiagram.GongMarshallField(stage, "IsHiddenPerpendicularVectorGridHalfway"))
		initializerStatements.WriteString(vase2ddiagram.GongMarshallField(stage, "IsHiddenTopStartArcShapeGrid"))
		initializerStatements.WriteString(vase2ddiagram.GongMarshallField(stage, "IsHiddenShiftedBottomTopStartArcShapeGrid"))
		initializerStatements.WriteString(vase2ddiagram.GongMarshallField(stage, "IsHiddenTopMidArcVectorShapeGrid"))
		initializerStatements.WriteString(vase2ddiagram.GongMarshallField(stage, "IsHiddenStartHalfwayArcShapeGrid"))
		initializerStatements.WriteString(vase2ddiagram.GongMarshallField(stage, "IsHiddenTopStartHalfwayArcShapeGrid"))
		initializerStatements.WriteString(vase2ddiagram.GongMarshallField(stage, "IsHiddenEndHalfwayArcShapeGrid"))
		initializerStatements.WriteString(vase2ddiagram.GongMarshallField(stage, "IsHiddenTopEndHalfwayArcShapeGrid"))
		initializerStatements.WriteString(vase2ddiagram.GongMarshallField(stage, "IsHiddenTopEndArcShapeGrid"))
		initializerStatements.WriteString(vase2ddiagram.GongMarshallField(stage, "IsHiddenStackOfGrowthCurve"))
		initializerStatements.WriteString(vase2ddiagram.GongMarshallField(stage, "IsHiddenTopStackOfGrowthCurve"))
		initializerStatements.WriteString(vase2ddiagram.GongMarshallField(stage, "IsHiddenTopGrowthCurve2D"))
		initializerStatements.WriteString(vase2ddiagram.GongMarshallField(stage, "IsHiddenStackOfGrowthCurve2D"))
		initializerStatements.WriteString(vase2ddiagram.GongMarshallField(stage, "IsHiddenTopStackOfGrowthCurve2D"))
		initializerStatements.WriteString(vase2ddiagram.GongMarshallField(stage, "IsHiddenGrowthCurve2DRibbon"))
		initializerStatements.WriteString(vase2ddiagram.GongMarshallField(stage, "IsHiddenShiftedRightGrowthCurve2DRibbon"))
		initializerStatements.WriteString(vase2ddiagram.GongMarshallField(stage, "IsHiddenShiftedLeftGrowthCurve2DRibbon"))
		initializerStatements.WriteString(vase2ddiagram.GongMarshallField(stage, "IsHiddenStackOfGrowthCurve2DRibbon"))
		initializerStatements.WriteString(vase2ddiagram.GongMarshallField(stage, "IsHiddenStackOfRotatedGrowthCurve2DRibbon"))
		initializerStatements.WriteString(vase2ddiagram.GongMarshallField(stage, "IsHiddenPartiallyGrowthCurve2DRibbon"))
		initializerStatements.WriteString(vase2ddiagram.GongMarshallField(stage, "IsHiddenShiftedLeftPartiallyGrowthCurve2DRibbon"))
		initializerStatements.WriteString(vase2ddiagram.GongMarshallField(stage, "IsHiddenPartiallyGrowthCurve2DTrajectory"))
		initializerStatements.WriteString(vase2ddiagram.GongMarshallField(stage, "IsHiddenPartiallyGrowthCurve2DTrajectoryP1P2"))
		initializerStatements.WriteString(vase2ddiagram.GongMarshallField(stage, "IsHiddenPxShape"))
		initializerStatements.WriteString(vase2ddiagram.GongMarshallField(stage, "IsHiddenChosenP1P2PairShape"))
		initializerStatements.WriteString(vase2ddiagram.GongMarshallField(stage, "IsHiddenKeyHoleShape"))
		initializerStatements.WriteString(vase2ddiagram.GongMarshallField(stage, "IsChecked"))
		initializerStatements.WriteString(vase2ddiagram.GongMarshallField(stage, "ComputedPrefix"))
		initializerStatements.WriteString(vase2ddiagram.GongMarshallField(stage, "IsExpanded"))
	}
	initRes = initializerStatements.String()
	ptrRes = pointersInitializesStatements.String()
	return
}
func (vasetrapezebaseplateshape *VaseTrapezeBasePlateShape) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(vasetrapezebaseplateshape.GongMarshallField(stage, "Name"))
	}
	initRes = initializerStatements.String()
	ptrRes = pointersInitializesStatements.String()
	return
}
func (vasetrapezeringshape *VaseTrapezeRingShape) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(vasetrapezeringshape.GongMarshallField(stage, "Name"))
	}
	initRes = initializerStatements.String()
	ptrRes = pointersInitializesStatements.String()
	return
}
func (verticaltorusstackshape *VerticalTorusStackShape) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(verticaltorusstackshape.GongMarshallField(stage, "Name"))
	}
	initRes = initializerStatements.String()
	ptrRes = pointersInitializesStatements.String()
	return
}
func (volumekey3dshape *VolumeKey3DShape) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(volumekey3dshape.GongMarshallField(stage, "Name"))
	}
	initRes = initializerStatements.String()
	ptrRes = pointersInitializesStatements.String()
	return
}
