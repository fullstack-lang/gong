// generated code - do not edit
package music

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
	{{Identifier}} := (&music.{{GeneratedStructName}}{Name: {{GeneratedFieldNameValue}}}).Stage(stage)`

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
			const fnSig = "func _(stage *music.Stage) {"
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
			strings.HasSuffix(trimmedContentBeforeBrace, "func _(stage *music.Stage) {") ||
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
	modelsPackageName := "github.com/fullstack-lang/gong/dsm/phylla/go/models/abstract/music"
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
		for _, musicabstract := range __gong__sortStageSetInstances(stage.MusicAbstracts, stage.MusicAbstract_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			musicabstractIdent := musicabstract.GongGetIdentifier(stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&music.MusicAbstract{Name: %s}).Stage(stage)", musicabstractIdent, __gong__toRawStringLiteral(musicabstract.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
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
go run github.com/fullstack-lang/gong/dsm/phylla/go/models/abstract/music/probe/cmd/probe@%s stage.go
*/
package %s

import (
	"slices"
	"time"

	"github.com/fullstack-lang/gong/dsm/phylla/go/models/abstract/music"%s
)

var (
	_ time.Time
	_ = slices.Index[[]int, int]

	_ *music.Stage
)%s

// function will stage objects across all coordinated stages
func _(stage *music.Stage) {

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

	if modelsPackageName != "" && modelsPackageName != "github.com/fullstack-lang/gong/dsm/phylla/go/models/abstract/music" {
		res = strings.ReplaceAll(res, "\"github.com/fullstack-lang/gong/dsm/phylla/go/models/abstract/music\"", fmt.Sprintf("%q", modelsPackageName))
	}

	return res, nil
}

// insertion point for marshall field methods
func (musicabstract *MusicAbstract) GongMarshallField(stage *Stage, fieldName string) (res string) {
	ident := musicabstract.GongGetIdentifier(stage)
	_ = ident
	switch fieldName {
	case "Name":
		res = __gong__marshallString(ident, "Name", musicabstract.Name)
	case "IsChecked":
		res = __gong__marshallBool(ident, "IsChecked", musicabstract.IsChecked)
	case "PitchHeight":
		res = __gong__marshallFloat(ident, "PitchHeight", musicabstract.PitchHeight)
	case "NbOfBeatsInTheme":
		res = __gong__marshallInt(ident, "NbOfBeatsInTheme", musicabstract.NbOfBeatsInTheme)
	case "BeatsPerSecond":
		res = __gong__marshallFloat(ident, "BeatsPerSecond", musicabstract.BeatsPerSecond)
	case "FirstVoiceShiftX":
		res = __gong__marshallFloat(ident, "FirstVoiceShiftX", musicabstract.FirstVoiceShiftX)
	case "FirstVoiceShiftY":
		res = __gong__marshallFloat(ident, "FirstVoiceShiftY", musicabstract.FirstVoiceShiftY)
	case "PitchDifference":
		res = __gong__marshallInt(ident, "PitchDifference", musicabstract.PitchDifference)
	case "Level":
		res = __gong__marshallFloat(ident, "Level", musicabstract.Level)
	case "ActualBeatsTemporalShift":
		res = __gong__marshallInt(ident, "ActualBeatsTemporalShift", musicabstract.ActualBeatsTemporalShift)
	case "IsMinor":
		res = __gong__marshallBool(ident, "IsMinor", musicabstract.IsMinor)
	case "ThemeBinaryEncoding":
		res = __gong__marshallInt(ident, "ThemeBinaryEncoding", musicabstract.ThemeBinaryEncoding)
	case "BezierControlLengthRatio":
		res = __gong__marshallFloat(ident, "BezierControlLengthRatio", musicabstract.BezierControlLengthRatio)
	case "NbPitchLines":
		res = __gong__marshallInt(ident, "NbPitchLines", musicabstract.NbPitchLines)
	case "NbBeatLines":
		res = __gong__marshallInt(ident, "NbBeatLines", musicabstract.NbBeatLines)
	case "OriginX":
		res = __gong__marshallFloat(ident, "OriginX", musicabstract.OriginX)
	case "OriginY":
		res = __gong__marshallFloat(ident, "OriginY", musicabstract.OriginY)
	case "ScoreScale":
		res = __gong__marshallFloat(ident, "ScoreScale", musicabstract.ScoreScale)
	case "ShowFirstVoice":
		res = __gong__marshallBool(ident, "ShowFirstVoice", musicabstract.ShowFirstVoice)
	case "ShowFirstVoiceShiftRight":
		res = __gong__marshallBool(ident, "ShowFirstVoiceShiftRight", musicabstract.ShowFirstVoiceShiftRight)
	case "ShowSecondVoice":
		res = __gong__marshallBool(ident, "ShowSecondVoice", musicabstract.ShowSecondVoice)
	case "ShowSecondVoiceShiftRight":
		res = __gong__marshallBool(ident, "ShowSecondVoiceShiftRight", musicabstract.ShowSecondVoiceShiftRight)
	case "ShowFirstVoiceNotes":
		res = __gong__marshallBool(ident, "ShowFirstVoiceNotes", musicabstract.ShowFirstVoiceNotes)
	case "ShowFirstVoiceNotesShiftRight":
		res = __gong__marshallBool(ident, "ShowFirstVoiceNotesShiftRight", musicabstract.ShowFirstVoiceNotesShiftRight)
	case "ShowSecondVoiceNotes":
		res = __gong__marshallBool(ident, "ShowSecondVoiceNotes", musicabstract.ShowSecondVoiceNotes)
	case "ShowSecondVoiceNotesShiftRight":
		res = __gong__marshallBool(ident, "ShowSecondVoiceNotesShiftRight", musicabstract.ShowSecondVoiceNotesShiftRight)
	case "IsComposerNodeExpanded":
		res = __gong__marshallBool(ident, "IsComposerNodeExpanded", musicabstract.IsComposerNodeExpanded)

	default:
		log.Panicf("Unknown field %s for Gongstruct MusicAbstract", fieldName)
	}
	return
}

// insertion point for marshall all fields methods
func (musicabstract *MusicAbstract) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(musicabstract.GongMarshallField(stage, "Name"))
		initializerStatements.WriteString(musicabstract.GongMarshallField(stage, "IsChecked"))
		initializerStatements.WriteString(musicabstract.GongMarshallField(stage, "PitchHeight"))
		initializerStatements.WriteString(musicabstract.GongMarshallField(stage, "NbOfBeatsInTheme"))
		initializerStatements.WriteString(musicabstract.GongMarshallField(stage, "BeatsPerSecond"))
		initializerStatements.WriteString(musicabstract.GongMarshallField(stage, "FirstVoiceShiftX"))
		initializerStatements.WriteString(musicabstract.GongMarshallField(stage, "FirstVoiceShiftY"))
		initializerStatements.WriteString(musicabstract.GongMarshallField(stage, "PitchDifference"))
		initializerStatements.WriteString(musicabstract.GongMarshallField(stage, "Level"))
		initializerStatements.WriteString(musicabstract.GongMarshallField(stage, "ActualBeatsTemporalShift"))
		initializerStatements.WriteString(musicabstract.GongMarshallField(stage, "IsMinor"))
		initializerStatements.WriteString(musicabstract.GongMarshallField(stage, "ThemeBinaryEncoding"))
		initializerStatements.WriteString(musicabstract.GongMarshallField(stage, "BezierControlLengthRatio"))
		initializerStatements.WriteString(musicabstract.GongMarshallField(stage, "NbPitchLines"))
		initializerStatements.WriteString(musicabstract.GongMarshallField(stage, "NbBeatLines"))
		initializerStatements.WriteString(musicabstract.GongMarshallField(stage, "OriginX"))
		initializerStatements.WriteString(musicabstract.GongMarshallField(stage, "OriginY"))
		initializerStatements.WriteString(musicabstract.GongMarshallField(stage, "ScoreScale"))
		initializerStatements.WriteString(musicabstract.GongMarshallField(stage, "ShowFirstVoice"))
		initializerStatements.WriteString(musicabstract.GongMarshallField(stage, "ShowFirstVoiceShiftRight"))
		initializerStatements.WriteString(musicabstract.GongMarshallField(stage, "ShowSecondVoice"))
		initializerStatements.WriteString(musicabstract.GongMarshallField(stage, "ShowSecondVoiceShiftRight"))
		initializerStatements.WriteString(musicabstract.GongMarshallField(stage, "ShowFirstVoiceNotes"))
		initializerStatements.WriteString(musicabstract.GongMarshallField(stage, "ShowFirstVoiceNotesShiftRight"))
		initializerStatements.WriteString(musicabstract.GongMarshallField(stage, "ShowSecondVoiceNotes"))
		initializerStatements.WriteString(musicabstract.GongMarshallField(stage, "ShowSecondVoiceNotesShiftRight"))
		initializerStatements.WriteString(musicabstract.GongMarshallField(stage, "IsComposerNodeExpanded"))
	}
	initRes = initializerStatements.String()
	ptrRes = pointersInitializesStatements.String()
	return
}
