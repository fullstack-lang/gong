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
	modelsPackageName := "github.com/fullstack-lang/gong/dsm/floss/go/models"
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
		for _, compareanalysis := range __gong__sortStageSetInstances(stage.CompareAnalysiss, stage.CompareAnalysis_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			compareanalysisIdent := localAliasPrefix + compareanalysis.GongGetIdentifier(stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&models.CompareAnalysis{Name: %s}).Stage(stage)", compareanalysisIdent, __gong__toRawStringLiteral(compareanalysis.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", compareanalysisIdent, __gong__toRawStringLiteral(compareanalysis.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Mu = %f", compareanalysisIdent, compareanalysis.Mu))
			values.WriteString(fmt.Sprintf("\n\t%s.Epsilon = %f", compareanalysisIdent, compareanalysis.Epsilon))
			values.WriteString(fmt.Sprintf("\n\t%s.ComputedPrefix = %s", compareanalysisIdent, __gong__toRawStringLiteral(compareanalysis.ComputedPrefix)))
			values.WriteString(fmt.Sprintf("\n\t%s.IsExpanded = %t", compareanalysisIdent, compareanalysis.IsExpanded))
			if compareanalysis.FromSystem != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := localAliasPrefix + compareanalysis.FromSystem.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.FromSystem = %s", compareanalysisIdent, targetIdent))
			}
			if compareanalysis.ToSystem != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := localAliasPrefix + compareanalysis.ToSystem.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.ToSystem = %s", compareanalysisIdent, targetIdent))
			}
			for _, elem := range compareanalysis.DiagramFlossEquations {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := localAliasPrefix + elem.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.DiagramFlossEquations = append(%s.DiagramFlossEquations, %s)", compareanalysisIdent, compareanalysisIdent, targetIdent))
			}
			for _, elem := range compareanalysis.DiagramFlossEquationsWhoseNodeIsExpanded {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := localAliasPrefix + elem.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.DiagramFlossEquationsWhoseNodeIsExpanded = append(%s.DiagramFlossEquationsWhoseNodeIsExpanded, %s)", compareanalysisIdent, compareanalysisIdent, targetIdent))
			}
		}
	}
	if stage != nil {
		for _, complexity := range __gong__sortStageSetInstances(stage.Complexitys, stage.Complexity_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			complexityIdent := localAliasPrefix + complexity.GongGetIdentifier(stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&models.Complexity{Name: %s}).Stage(stage)", complexityIdent, __gong__toRawStringLiteral(complexity.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", complexityIdent, __gong__toRawStringLiteral(complexity.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Strength = %f", complexityIdent, complexity.Strength))
			values.WriteString(fmt.Sprintf("\n\t%s.Description = %s", complexityIdent, __gong__toRawStringLiteral(complexity.Description)))
			values.WriteString(fmt.Sprintf("\n\t%s.ComputedPrefix = %s", complexityIdent, __gong__toRawStringLiteral(complexity.ComputedPrefix)))
			values.WriteString(fmt.Sprintf("\n\t%s.IsExpanded = %t", complexityIdent, complexity.IsExpanded))
		}
	}
	if stage != nil {
		for _, diagramflossequation := range __gong__sortStageSetInstances(stage.DiagramFlossEquations, stage.DiagramFlossEquation_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			diagramflossequationIdent := localAliasPrefix + diagramflossequation.GongGetIdentifier(stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&models.DiagramFlossEquation{Name: %s}).Stage(stage)", diagramflossequationIdent, __gong__toRawStringLiteral(diagramflossequation.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", diagramflossequationIdent, __gong__toRawStringLiteral(diagramflossequation.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Description = %s", diagramflossequationIdent, __gong__toRawStringLiteral(diagramflossequation.Description)))
			values.WriteString(fmt.Sprintf("\n\t%s.Scale = %f", diagramflossequationIdent, diagramflossequation.Scale))
			if diagramflossequation.FontSize.ToCodeString() != "" {
				values.WriteString(fmt.Sprintf("\n\t%s.FontSize = models.%s", diagramflossequationIdent, diagramflossequation.FontSize.ToCodeString()))
			} else {
				values.WriteString(fmt.Sprintf("\n\t%s.FontSize = \"\"", diagramflossequationIdent))
			}
			values.WriteString(fmt.Sprintf("\n\t%s.ComputedPrefix = %s", diagramflossequationIdent, __gong__toRawStringLiteral(diagramflossequation.ComputedPrefix)))
			values.WriteString(fmt.Sprintf("\n\t%s.IsExpanded = %t", diagramflossequationIdent, diagramflossequation.IsExpanded))
			values.WriteString(fmt.Sprintf("\n\t%s.IsChecked = %t", diagramflossequationIdent, diagramflossequation.IsChecked))
			values.WriteString(fmt.Sprintf("\n\t%s.IsEditable_ = %t", diagramflossequationIdent, diagramflossequation.IsEditable_))
			values.WriteString(fmt.Sprintf("\n\t%s.IsInDelta3ColumnsMode = %t", diagramflossequationIdent, diagramflossequation.IsInDelta3ColumnsMode))
			values.WriteString(fmt.Sprintf("\n\t%s.AreQuantitativeElementsVisible = %t", diagramflossequationIdent, diagramflossequation.AreQuantitativeElementsVisible))
			values.WriteString(fmt.Sprintf("\n\t%s.AreSubsystemsVisible = %t", diagramflossequationIdent, diagramflossequation.AreSubsystemsVisible))
			values.WriteString(fmt.Sprintf("\n\t%s.AreCommonElementsHidden = %t", diagramflossequationIdent, diagramflossequation.AreCommonElementsHidden))
			values.WriteString(fmt.Sprintf("\n\t%s.AreCPEArrowsVisible = %t", diagramflossequationIdent, diagramflossequation.AreCPEArrowsVisible))
			values.WriteString(fmt.Sprintf("\n\t%s.AreColumnTitlesVisible = %t", diagramflossequationIdent, diagramflossequation.AreColumnTitlesVisible))
			values.WriteString(fmt.Sprintf("\n\t%s.Width = %f", diagramflossequationIdent, diagramflossequation.Width))
			values.WriteString(fmt.Sprintf("\n\t%s.Height = %f", diagramflossequationIdent, diagramflossequation.Height))
			values.WriteString(fmt.Sprintf("\n\t%s.DefaultBoxWidth = %f", diagramflossequationIdent, diagramflossequation.DefaultBoxWidth))
			values.WriteString(fmt.Sprintf("\n\t%s.DefaultBoxHeigth = %f", diagramflossequationIdent, diagramflossequation.DefaultBoxHeigth))
			values.WriteString(fmt.Sprintf("\n\t%s.IsNotesNodeExpanded = %t", diagramflossequationIdent, diagramflossequation.IsNotesNodeExpanded))
			values.WriteString(fmt.Sprintf("\n\t%s.IsComplexitysNodeExpanded = %t", diagramflossequationIdent, diagramflossequation.IsComplexitysNodeExpanded))
			values.WriteString(fmt.Sprintf("\n\t%s.IsPerformancesNodeExpanded = %t", diagramflossequationIdent, diagramflossequation.IsPerformancesNodeExpanded))
			values.WriteString(fmt.Sprintf("\n\t%s.IsEffortsNodeExpanded = %t", diagramflossequationIdent, diagramflossequation.IsEffortsNodeExpanded))
			for _, elem := range diagramflossequation.Note_Shapes {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := localAliasPrefix + elem.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Note_Shapes = append(%s.Note_Shapes, %s)", diagramflossequationIdent, diagramflossequationIdent, targetIdent))
			}
			for _, elem := range diagramflossequation.NoteComplexityShapes {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := localAliasPrefix + elem.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.NoteComplexityShapes = append(%s.NoteComplexityShapes, %s)", diagramflossequationIdent, diagramflossequationIdent, targetIdent))
			}
			for _, elem := range diagramflossequation.NotePerformanceShapes {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := localAliasPrefix + elem.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.NotePerformanceShapes = append(%s.NotePerformanceShapes, %s)", diagramflossequationIdent, diagramflossequationIdent, targetIdent))
			}
			for _, elem := range diagramflossequation.NoteEffortShapes {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := localAliasPrefix + elem.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.NoteEffortShapes = append(%s.NoteEffortShapes, %s)", diagramflossequationIdent, diagramflossequationIdent, targetIdent))
			}
			for _, elem := range diagramflossequation.NotesWhoseNodeIsExpanded {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := localAliasPrefix + elem.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.NotesWhoseNodeIsExpanded = append(%s.NotesWhoseNodeIsExpanded, %s)", diagramflossequationIdent, diagramflossequationIdent, targetIdent))
			}
			for _, elem := range diagramflossequation.ComplexitysWhoseNodeIsExpanded {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := localAliasPrefix + elem.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.ComplexitysWhoseNodeIsExpanded = append(%s.ComplexitysWhoseNodeIsExpanded, %s)", diagramflossequationIdent, diagramflossequationIdent, targetIdent))
			}
			for _, elem := range diagramflossequation.PerformancesWhoseNodeIsExpanded {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := localAliasPrefix + elem.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.PerformancesWhoseNodeIsExpanded = append(%s.PerformancesWhoseNodeIsExpanded, %s)", diagramflossequationIdent, diagramflossequationIdent, targetIdent))
			}
			for _, elem := range diagramflossequation.EffortsWhoseNodeIsExpanded {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := localAliasPrefix + elem.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.EffortsWhoseNodeIsExpanded = append(%s.EffortsWhoseNodeIsExpanded, %s)", diagramflossequationIdent, diagramflossequationIdent, targetIdent))
			}
		}
	}
	if stage != nil {
		for _, effort := range __gong__sortStageSetInstances(stage.Efforts, stage.Effort_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			effortIdent := localAliasPrefix + effort.GongGetIdentifier(stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&models.Effort{Name: %s}).Stage(stage)", effortIdent, __gong__toRawStringLiteral(effort.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", effortIdent, __gong__toRawStringLiteral(effort.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Strength = %f", effortIdent, effort.Strength))
			values.WriteString(fmt.Sprintf("\n\t%s.Description = %s", effortIdent, __gong__toRawStringLiteral(effort.Description)))
			values.WriteString(fmt.Sprintf("\n\t%s.ComputedPrefix = %s", effortIdent, __gong__toRawStringLiteral(effort.ComputedPrefix)))
			values.WriteString(fmt.Sprintf("\n\t%s.IsExpanded = %t", effortIdent, effort.IsExpanded))
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
			values.WriteString(fmt.Sprintf("\n\t%s.IsSystemsNodeExpanded = %t", libraryIdent, library.IsSystemsNodeExpanded))
			values.WriteString(fmt.Sprintf("\n\t%s.IsComplexitysNodeExpanded = %t", libraryIdent, library.IsComplexitysNodeExpanded))
			values.WriteString(fmt.Sprintf("\n\t%s.IsPerformancesNodeExpanded = %t", libraryIdent, library.IsPerformancesNodeExpanded))
			values.WriteString(fmt.Sprintf("\n\t%s.IsEffortsNodeExpanded = %t", libraryIdent, library.IsEffortsNodeExpanded))
			values.WriteString(fmt.Sprintf("\n\t%s.IsCompareAnalysisNodeExpanded = %t", libraryIdent, library.IsCompareAnalysisNodeExpanded))
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
			for _, elem := range library.RootSystems {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := localAliasPrefix + elem.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.RootSystems = append(%s.RootSystems, %s)", libraryIdent, libraryIdent, targetIdent))
			}
			for _, elem := range library.RootComplexitys {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := localAliasPrefix + elem.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.RootComplexitys = append(%s.RootComplexitys, %s)", libraryIdent, libraryIdent, targetIdent))
			}
			for _, elem := range library.RootPerformances {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := localAliasPrefix + elem.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.RootPerformances = append(%s.RootPerformances, %s)", libraryIdent, libraryIdent, targetIdent))
			}
			for _, elem := range library.RootEfforts {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := localAliasPrefix + elem.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.RootEfforts = append(%s.RootEfforts, %s)", libraryIdent, libraryIdent, targetIdent))
			}
			for _, elem := range library.RootCompareAnalysis {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := localAliasPrefix + elem.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.RootCompareAnalysis = append(%s.RootCompareAnalysis, %s)", libraryIdent, libraryIdent, targetIdent))
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
			for _, elem := range library.SystemsWhoseNodeIsExpanded {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := localAliasPrefix + elem.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.SystemsWhoseNodeIsExpanded = append(%s.SystemsWhoseNodeIsExpanded, %s)", libraryIdent, libraryIdent, targetIdent))
			}
			for _, elem := range library.ComplexitysWhoseNodeIsExpanded {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := localAliasPrefix + elem.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.ComplexitysWhoseNodeIsExpanded = append(%s.ComplexitysWhoseNodeIsExpanded, %s)", libraryIdent, libraryIdent, targetIdent))
			}
			for _, elem := range library.PerformancesWhoseNodeIsExpanded {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := localAliasPrefix + elem.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.PerformancesWhoseNodeIsExpanded = append(%s.PerformancesWhoseNodeIsExpanded, %s)", libraryIdent, libraryIdent, targetIdent))
			}
			for _, elem := range library.EffortsWhoseNodeIsExpanded {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := localAliasPrefix + elem.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.EffortsWhoseNodeIsExpanded = append(%s.EffortsWhoseNodeIsExpanded, %s)", libraryIdent, libraryIdent, targetIdent))
			}
			for _, elem := range library.CompareAnalysisWhoseNodeIsExpanded {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := localAliasPrefix + elem.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.CompareAnalysisWhoseNodeIsExpanded = append(%s.CompareAnalysisWhoseNodeIsExpanded, %s)", libraryIdent, libraryIdent, targetIdent))
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
			values.WriteString(fmt.Sprintf("\n\t%s.IsComplexitysNodeExpanded = %t", noteIdent, note.IsComplexitysNodeExpanded))
			values.WriteString(fmt.Sprintf("\n\t%s.IsPerformancesNodeExpanded = %t", noteIdent, note.IsPerformancesNodeExpanded))
			values.WriteString(fmt.Sprintf("\n\t%s.IsEffortsNodeExpanded = %t", noteIdent, note.IsEffortsNodeExpanded))
			for _, elem := range note.Complexities {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := localAliasPrefix + elem.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Complexities = append(%s.Complexities, %s)", noteIdent, noteIdent, targetIdent))
			}
			for _, elem := range note.Performances {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := localAliasPrefix + elem.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Performances = append(%s.Performances, %s)", noteIdent, noteIdent, targetIdent))
			}
			for _, elem := range note.Efforts {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := localAliasPrefix + elem.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Efforts = append(%s.Efforts, %s)", noteIdent, noteIdent, targetIdent))
			}
		}
	}
	if stage != nil {
		for _, notecomplexityshape := range __gong__sortStageSetInstances(stage.NoteComplexityShapes, stage.NoteComplexityShape_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			notecomplexityshapeIdent := localAliasPrefix + notecomplexityshape.GongGetIdentifier(stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&models.NoteComplexityShape{Name: %s}).Stage(stage)", notecomplexityshapeIdent, __gong__toRawStringLiteral(notecomplexityshape.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", notecomplexityshapeIdent, __gong__toRawStringLiteral(notecomplexityshape.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.StartRatio = %f", notecomplexityshapeIdent, notecomplexityshape.StartRatio))
			values.WriteString(fmt.Sprintf("\n\t%s.EndRatio = %f", notecomplexityshapeIdent, notecomplexityshape.EndRatio))
			if notecomplexityshape.StartOrientation.ToCodeString() != "" {
				values.WriteString(fmt.Sprintf("\n\t%s.StartOrientation = models.%s", notecomplexityshapeIdent, notecomplexityshape.StartOrientation.ToCodeString()))
			} else {
				values.WriteString(fmt.Sprintf("\n\t%s.StartOrientation = \"\"", notecomplexityshapeIdent))
			}
			if notecomplexityshape.EndOrientation.ToCodeString() != "" {
				values.WriteString(fmt.Sprintf("\n\t%s.EndOrientation = models.%s", notecomplexityshapeIdent, notecomplexityshape.EndOrientation.ToCodeString()))
			} else {
				values.WriteString(fmt.Sprintf("\n\t%s.EndOrientation = \"\"", notecomplexityshapeIdent))
			}
			values.WriteString(fmt.Sprintf("\n\t%s.CornerOffsetRatio = %f", notecomplexityshapeIdent, notecomplexityshape.CornerOffsetRatio))
			values.WriteString(fmt.Sprintf("\n\t%s.IsHidden = %t", notecomplexityshapeIdent, notecomplexityshape.IsHidden))
			values.WriteString(fmt.Sprintf("\n\t%s.IsPresentationOverridden = %t", notecomplexityshapeIdent, notecomplexityshape.IsPresentationOverridden))
			values.WriteString(fmt.Sprintf("\n\t%s.Stroke = %s", notecomplexityshapeIdent, __gong__toRawStringLiteral(notecomplexityshape.Stroke)))
			values.WriteString(fmt.Sprintf("\n\t%s.StrokeWidth = %f", notecomplexityshapeIdent, notecomplexityshape.StrokeWidth))
			values.WriteString(fmt.Sprintf("\n\t%s.StrokeOpacity = %f", notecomplexityshapeIdent, notecomplexityshape.StrokeOpacity))
			values.WriteString(fmt.Sprintf("\n\t%s.StrokeDashArray = %s", notecomplexityshapeIdent, __gong__toRawStringLiteral(notecomplexityshape.StrokeDashArray)))
			values.WriteString(fmt.Sprintf("\n\t%s.FontSize = %s", notecomplexityshapeIdent, __gong__toRawStringLiteral(notecomplexityshape.FontSize)))
			values.WriteString(fmt.Sprintf("\n\t%s.FontColor = %s", notecomplexityshapeIdent, __gong__toRawStringLiteral(notecomplexityshape.FontColor)))
			if notecomplexityshape.Note != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := localAliasPrefix + notecomplexityshape.Note.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Note = %s", notecomplexityshapeIdent, targetIdent))
			}
			if notecomplexityshape.Complexity != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := localAliasPrefix + notecomplexityshape.Complexity.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Complexity = %s", notecomplexityshapeIdent, targetIdent))
			}
		}
	}
	if stage != nil {
		for _, noteeffortshape := range __gong__sortStageSetInstances(stage.NoteEffortShapes, stage.NoteEffortShape_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			noteeffortshapeIdent := localAliasPrefix + noteeffortshape.GongGetIdentifier(stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&models.NoteEffortShape{Name: %s}).Stage(stage)", noteeffortshapeIdent, __gong__toRawStringLiteral(noteeffortshape.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", noteeffortshapeIdent, __gong__toRawStringLiteral(noteeffortshape.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.StartRatio = %f", noteeffortshapeIdent, noteeffortshape.StartRatio))
			values.WriteString(fmt.Sprintf("\n\t%s.EndRatio = %f", noteeffortshapeIdent, noteeffortshape.EndRatio))
			if noteeffortshape.StartOrientation.ToCodeString() != "" {
				values.WriteString(fmt.Sprintf("\n\t%s.StartOrientation = models.%s", noteeffortshapeIdent, noteeffortshape.StartOrientation.ToCodeString()))
			} else {
				values.WriteString(fmt.Sprintf("\n\t%s.StartOrientation = \"\"", noteeffortshapeIdent))
			}
			if noteeffortshape.EndOrientation.ToCodeString() != "" {
				values.WriteString(fmt.Sprintf("\n\t%s.EndOrientation = models.%s", noteeffortshapeIdent, noteeffortshape.EndOrientation.ToCodeString()))
			} else {
				values.WriteString(fmt.Sprintf("\n\t%s.EndOrientation = \"\"", noteeffortshapeIdent))
			}
			values.WriteString(fmt.Sprintf("\n\t%s.CornerOffsetRatio = %f", noteeffortshapeIdent, noteeffortshape.CornerOffsetRatio))
			values.WriteString(fmt.Sprintf("\n\t%s.IsHidden = %t", noteeffortshapeIdent, noteeffortshape.IsHidden))
			values.WriteString(fmt.Sprintf("\n\t%s.IsPresentationOverridden = %t", noteeffortshapeIdent, noteeffortshape.IsPresentationOverridden))
			values.WriteString(fmt.Sprintf("\n\t%s.Stroke = %s", noteeffortshapeIdent, __gong__toRawStringLiteral(noteeffortshape.Stroke)))
			values.WriteString(fmt.Sprintf("\n\t%s.StrokeWidth = %f", noteeffortshapeIdent, noteeffortshape.StrokeWidth))
			values.WriteString(fmt.Sprintf("\n\t%s.StrokeOpacity = %f", noteeffortshapeIdent, noteeffortshape.StrokeOpacity))
			values.WriteString(fmt.Sprintf("\n\t%s.StrokeDashArray = %s", noteeffortshapeIdent, __gong__toRawStringLiteral(noteeffortshape.StrokeDashArray)))
			values.WriteString(fmt.Sprintf("\n\t%s.FontSize = %s", noteeffortshapeIdent, __gong__toRawStringLiteral(noteeffortshape.FontSize)))
			values.WriteString(fmt.Sprintf("\n\t%s.FontColor = %s", noteeffortshapeIdent, __gong__toRawStringLiteral(noteeffortshape.FontColor)))
			if noteeffortshape.Note != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := localAliasPrefix + noteeffortshape.Note.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Note = %s", noteeffortshapeIdent, targetIdent))
			}
			if noteeffortshape.Effort != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := localAliasPrefix + noteeffortshape.Effort.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Effort = %s", noteeffortshapeIdent, targetIdent))
			}
		}
	}
	if stage != nil {
		for _, noteperformanceshape := range __gong__sortStageSetInstances(stage.NotePerformanceShapes, stage.NotePerformanceShape_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			noteperformanceshapeIdent := localAliasPrefix + noteperformanceshape.GongGetIdentifier(stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&models.NotePerformanceShape{Name: %s}).Stage(stage)", noteperformanceshapeIdent, __gong__toRawStringLiteral(noteperformanceshape.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", noteperformanceshapeIdent, __gong__toRawStringLiteral(noteperformanceshape.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.StartRatio = %f", noteperformanceshapeIdent, noteperformanceshape.StartRatio))
			values.WriteString(fmt.Sprintf("\n\t%s.EndRatio = %f", noteperformanceshapeIdent, noteperformanceshape.EndRatio))
			if noteperformanceshape.StartOrientation.ToCodeString() != "" {
				values.WriteString(fmt.Sprintf("\n\t%s.StartOrientation = models.%s", noteperformanceshapeIdent, noteperformanceshape.StartOrientation.ToCodeString()))
			} else {
				values.WriteString(fmt.Sprintf("\n\t%s.StartOrientation = \"\"", noteperformanceshapeIdent))
			}
			if noteperformanceshape.EndOrientation.ToCodeString() != "" {
				values.WriteString(fmt.Sprintf("\n\t%s.EndOrientation = models.%s", noteperformanceshapeIdent, noteperformanceshape.EndOrientation.ToCodeString()))
			} else {
				values.WriteString(fmt.Sprintf("\n\t%s.EndOrientation = \"\"", noteperformanceshapeIdent))
			}
			values.WriteString(fmt.Sprintf("\n\t%s.CornerOffsetRatio = %f", noteperformanceshapeIdent, noteperformanceshape.CornerOffsetRatio))
			values.WriteString(fmt.Sprintf("\n\t%s.IsHidden = %t", noteperformanceshapeIdent, noteperformanceshape.IsHidden))
			values.WriteString(fmt.Sprintf("\n\t%s.IsPresentationOverridden = %t", noteperformanceshapeIdent, noteperformanceshape.IsPresentationOverridden))
			values.WriteString(fmt.Sprintf("\n\t%s.Stroke = %s", noteperformanceshapeIdent, __gong__toRawStringLiteral(noteperformanceshape.Stroke)))
			values.WriteString(fmt.Sprintf("\n\t%s.StrokeWidth = %f", noteperformanceshapeIdent, noteperformanceshape.StrokeWidth))
			values.WriteString(fmt.Sprintf("\n\t%s.StrokeOpacity = %f", noteperformanceshapeIdent, noteperformanceshape.StrokeOpacity))
			values.WriteString(fmt.Sprintf("\n\t%s.StrokeDashArray = %s", noteperformanceshapeIdent, __gong__toRawStringLiteral(noteperformanceshape.StrokeDashArray)))
			values.WriteString(fmt.Sprintf("\n\t%s.FontSize = %s", noteperformanceshapeIdent, __gong__toRawStringLiteral(noteperformanceshape.FontSize)))
			values.WriteString(fmt.Sprintf("\n\t%s.FontColor = %s", noteperformanceshapeIdent, __gong__toRawStringLiteral(noteperformanceshape.FontColor)))
			if noteperformanceshape.Note != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := localAliasPrefix + noteperformanceshape.Note.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Note = %s", noteperformanceshapeIdent, targetIdent))
			}
			if noteperformanceshape.Performance != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := localAliasPrefix + noteperformanceshape.Performance.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Performance = %s", noteperformanceshapeIdent, targetIdent))
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
		for _, performance := range __gong__sortStageSetInstances(stage.Performances, stage.Performance_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			performanceIdent := localAliasPrefix + performance.GongGetIdentifier(stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&models.Performance{Name: %s}).Stage(stage)", performanceIdent, __gong__toRawStringLiteral(performance.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", performanceIdent, __gong__toRawStringLiteral(performance.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Strength = %f", performanceIdent, performance.Strength))
			values.WriteString(fmt.Sprintf("\n\t%s.Description = %s", performanceIdent, __gong__toRawStringLiteral(performance.Description)))
			values.WriteString(fmt.Sprintf("\n\t%s.ComputedPrefix = %s", performanceIdent, __gong__toRawStringLiteral(performance.ComputedPrefix)))
			values.WriteString(fmt.Sprintf("\n\t%s.IsExpanded = %t", performanceIdent, performance.IsExpanded))
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
			systemIdent := localAliasPrefix + system.GongGetIdentifier(stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&models.System{Name: %s}).Stage(stage)", systemIdent, __gong__toRawStringLiteral(system.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", systemIdent, __gong__toRawStringLiteral(system.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Description = %s", systemIdent, __gong__toRawStringLiteral(system.Description)))
			values.WriteString(fmt.Sprintf("\n\t%s.AreCPEsCompoundedFromSubSystems = %t", systemIdent, system.AreCPEsCompoundedFromSubSystems))
			values.WriteString(fmt.Sprintf("\n\t%s.ComputedPrefix = %s", systemIdent, __gong__toRawStringLiteral(system.ComputedPrefix)))
			values.WriteString(fmt.Sprintf("\n\t%s.IsExpanded = %t", systemIdent, system.IsExpanded))
			values.WriteString(fmt.Sprintf("\n\t%s.SVG_Path = %s", systemIdent, __gong__toRawStringLiteral(system.SVG_Path)))
			values.WriteString(fmt.Sprintf("\n\t%s.InverseAppliedScaling = %f", systemIdent, system.InverseAppliedScaling))
			values.WriteString(fmt.Sprintf("\n\t%s.IsSubSystemNodeExpanded = %t", systemIdent, system.IsSubSystemNodeExpanded))
			values.WriteString(fmt.Sprintf("\n\t%s.IsComplexitysNodeExpanded = %t", systemIdent, system.IsComplexitysNodeExpanded))
			values.WriteString(fmt.Sprintf("\n\t%s.IsPerformancesNodeExpanded = %t", systemIdent, system.IsPerformancesNodeExpanded))
			values.WriteString(fmt.Sprintf("\n\t%s.IsEffortsNodeExpanded = %t", systemIdent, system.IsEffortsNodeExpanded))
			for _, elem := range system.Complexities {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := localAliasPrefix + elem.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Complexities = append(%s.Complexities, %s)", systemIdent, systemIdent, targetIdent))
			}
			for _, elem := range system.Performances {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := localAliasPrefix + elem.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Performances = append(%s.Performances, %s)", systemIdent, systemIdent, targetIdent))
			}
			for _, elem := range system.Efforts {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := localAliasPrefix + elem.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Efforts = append(%s.Efforts, %s)", systemIdent, systemIdent, targetIdent))
			}
			for _, elem := range system.SubSystems {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := localAliasPrefix + elem.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.SubSystems = append(%s.SubSystems, %s)", systemIdent, systemIdent, targetIdent))
			}
			for _, elem := range system.DiagramFlossEquations {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := localAliasPrefix + elem.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.DiagramFlossEquations = append(%s.DiagramFlossEquations, %s)", systemIdent, systemIdent, targetIdent))
			}
			for _, elem := range system.DiagramFlossEquationsWhoseNodeIsExpanded {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := localAliasPrefix + elem.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.DiagramFlossEquationsWhoseNodeIsExpanded = append(%s.DiagramFlossEquationsWhoseNodeIsExpanded, %s)", systemIdent, systemIdent, targetIdent))
			}
			for _, elem := range system.ComplexitysWhoseNodeIsExpanded {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := localAliasPrefix + elem.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.ComplexitysWhoseNodeIsExpanded = append(%s.ComplexitysWhoseNodeIsExpanded, %s)", systemIdent, systemIdent, targetIdent))
			}
			for _, elem := range system.PerformancesWhoseNodeIsExpanded {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := localAliasPrefix + elem.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.PerformancesWhoseNodeIsExpanded = append(%s.PerformancesWhoseNodeIsExpanded, %s)", systemIdent, systemIdent, targetIdent))
			}
			for _, elem := range system.EffortsWhoseNodeIsExpanded {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := localAliasPrefix + elem.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.EffortsWhoseNodeIsExpanded = append(%s.EffortsWhoseNodeIsExpanded, %s)", systemIdent, systemIdent, targetIdent))
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
go run github.com/fullstack-lang/gong/dsm/floss/go/models/probe/cmd/probe@%s stage.go
*/
package %s

import (
	"slices"
	"time"

	"github.com/fullstack-lang/gong/dsm/floss/go/models"%s
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

	if modelsPackageName != "" && modelsPackageName != "github.com/fullstack-lang/gong/dsm/floss/go/models" {
		res = strings.ReplaceAll(res, "\"github.com/fullstack-lang/gong/dsm/floss/go/models\"", fmt.Sprintf("%q", modelsPackageName))
	}

	return res, nil
}

// insertion point for marshall field methods
func (compareanalysis *CompareAnalysis) GongMarshallField(stage *Stage, fieldName string) (res string) {
	ident := compareanalysis.GongGetIdentifier(stage)
	_ = ident
	switch fieldName {
	case "Name":
		res = __gong__marshallString(ident, "Name", compareanalysis.Name)
	case "Mu":
		res = __gong__marshallFloat(ident, "Mu", compareanalysis.Mu)
	case "Epsilon":
		res = __gong__marshallFloat(ident, "Epsilon", compareanalysis.Epsilon)
	case "ComputedPrefix":
		res = __gong__marshallString(ident, "ComputedPrefix", compareanalysis.ComputedPrefix)
	case "IsExpanded":
		res = __gong__marshallBool(ident, "IsExpanded", compareanalysis.IsExpanded)

	case "FromSystem":
		if compareanalysis.FromSystem != nil {
			res = __gong__marshallPointer(ident, "FromSystem", compareanalysis.FromSystem.GongGetIdentifier(stage))
		} else {
			res = __gong__marshallPointer(ident, "FromSystem", "nil")
		}
	case "ToSystem":
		if compareanalysis.ToSystem != nil {
			res = __gong__marshallPointer(ident, "ToSystem", compareanalysis.ToSystem.GongGetIdentifier(stage))
		} else {
			res = __gong__marshallPointer(ident, "ToSystem", "nil")
		}
	case "DiagramFlossEquations":
		var sb strings.Builder
		for _, _diagramflossequation := range compareanalysis.DiagramFlossEquations {
			sb.WriteString(__gong__marshallSliceOfPointers(ident, "DiagramFlossEquations", _diagramflossequation.GongGetIdentifier(stage)))
		}
		res = sb.String()
	case "DiagramFlossEquationsWhoseNodeIsExpanded":
		var sb strings.Builder
		for _, _diagramflossequation := range compareanalysis.DiagramFlossEquationsWhoseNodeIsExpanded {
			sb.WriteString(__gong__marshallSliceOfPointers(ident, "DiagramFlossEquationsWhoseNodeIsExpanded", _diagramflossequation.GongGetIdentifier(stage)))
		}
		res = sb.String()
	default:
		log.Panicf("Unknown field %s for Gongstruct CompareAnalysis", fieldName)
	}
	return
}

func (complexity *Complexity) GongMarshallField(stage *Stage, fieldName string) (res string) {
	ident := complexity.GongGetIdentifier(stage)
	_ = ident
	switch fieldName {
	case "Name":
		res = __gong__marshallString(ident, "Name", complexity.Name)
	case "Strength":
		res = __gong__marshallFloat(ident, "Strength", complexity.Strength)
	case "Description":
		res = __gong__marshallString(ident, "Description", complexity.Description)
	case "ComputedPrefix":
		res = __gong__marshallString(ident, "ComputedPrefix", complexity.ComputedPrefix)
	case "IsExpanded":
		res = __gong__marshallBool(ident, "IsExpanded", complexity.IsExpanded)

	default:
		log.Panicf("Unknown field %s for Gongstruct Complexity", fieldName)
	}
	return
}

func (diagramflossequation *DiagramFlossEquation) GongMarshallField(stage *Stage, fieldName string) (res string) {
	ident := diagramflossequation.GongGetIdentifier(stage)
	_ = ident
	switch fieldName {
	case "Name":
		res = __gong__marshallString(ident, "Name", diagramflossequation.Name)
	case "Description":
		res = __gong__marshallString(ident, "Description", diagramflossequation.Description)
	case "Scale":
		res = __gong__marshallFloat(ident, "Scale", diagramflossequation.Scale)
	case "FontSize":
		res = __gong__marshallEnumString(ident, "FontSize", diagramflossequation.FontSize.ToCodeString())
	case "ComputedPrefix":
		res = __gong__marshallString(ident, "ComputedPrefix", diagramflossequation.ComputedPrefix)
	case "IsExpanded":
		res = __gong__marshallBool(ident, "IsExpanded", diagramflossequation.IsExpanded)
	case "IsChecked":
		res = __gong__marshallBool(ident, "IsChecked", diagramflossequation.IsChecked)
	case "IsEditable_":
		res = __gong__marshallBool(ident, "IsEditable_", diagramflossequation.IsEditable_)
	case "IsInDelta3ColumnsMode":
		res = __gong__marshallBool(ident, "IsInDelta3ColumnsMode", diagramflossequation.IsInDelta3ColumnsMode)
	case "AreQuantitativeElementsVisible":
		res = __gong__marshallBool(ident, "AreQuantitativeElementsVisible", diagramflossequation.AreQuantitativeElementsVisible)
	case "AreSubsystemsVisible":
		res = __gong__marshallBool(ident, "AreSubsystemsVisible", diagramflossequation.AreSubsystemsVisible)
	case "AreCommonElementsHidden":
		res = __gong__marshallBool(ident, "AreCommonElementsHidden", diagramflossequation.AreCommonElementsHidden)
	case "AreCPEArrowsVisible":
		res = __gong__marshallBool(ident, "AreCPEArrowsVisible", diagramflossequation.AreCPEArrowsVisible)
	case "AreColumnTitlesVisible":
		res = __gong__marshallBool(ident, "AreColumnTitlesVisible", diagramflossequation.AreColumnTitlesVisible)
	case "Width":
		res = __gong__marshallFloat(ident, "Width", diagramflossequation.Width)
	case "Height":
		res = __gong__marshallFloat(ident, "Height", diagramflossequation.Height)
	case "DefaultBoxWidth":
		res = __gong__marshallFloat(ident, "DefaultBoxWidth", diagramflossequation.DefaultBoxWidth)
	case "DefaultBoxHeigth":
		res = __gong__marshallFloat(ident, "DefaultBoxHeigth", diagramflossequation.DefaultBoxHeigth)
	case "IsNotesNodeExpanded":
		res = __gong__marshallBool(ident, "IsNotesNodeExpanded", diagramflossequation.IsNotesNodeExpanded)
	case "IsComplexitysNodeExpanded":
		res = __gong__marshallBool(ident, "IsComplexitysNodeExpanded", diagramflossequation.IsComplexitysNodeExpanded)
	case "IsPerformancesNodeExpanded":
		res = __gong__marshallBool(ident, "IsPerformancesNodeExpanded", diagramflossequation.IsPerformancesNodeExpanded)
	case "IsEffortsNodeExpanded":
		res = __gong__marshallBool(ident, "IsEffortsNodeExpanded", diagramflossequation.IsEffortsNodeExpanded)

	case "Note_Shapes":
		var sb strings.Builder
		for _, _noteshape := range diagramflossequation.Note_Shapes {
			sb.WriteString(__gong__marshallSliceOfPointers(ident, "Note_Shapes", _noteshape.GongGetIdentifier(stage)))
		}
		res = sb.String()
	case "NoteComplexityShapes":
		var sb strings.Builder
		for _, _notecomplexityshape := range diagramflossequation.NoteComplexityShapes {
			sb.WriteString(__gong__marshallSliceOfPointers(ident, "NoteComplexityShapes", _notecomplexityshape.GongGetIdentifier(stage)))
		}
		res = sb.String()
	case "NotePerformanceShapes":
		var sb strings.Builder
		for _, _noteperformanceshape := range diagramflossequation.NotePerformanceShapes {
			sb.WriteString(__gong__marshallSliceOfPointers(ident, "NotePerformanceShapes", _noteperformanceshape.GongGetIdentifier(stage)))
		}
		res = sb.String()
	case "NoteEffortShapes":
		var sb strings.Builder
		for _, _noteeffortshape := range diagramflossequation.NoteEffortShapes {
			sb.WriteString(__gong__marshallSliceOfPointers(ident, "NoteEffortShapes", _noteeffortshape.GongGetIdentifier(stage)))
		}
		res = sb.String()
	case "NotesWhoseNodeIsExpanded":
		var sb strings.Builder
		for _, _note := range diagramflossequation.NotesWhoseNodeIsExpanded {
			sb.WriteString(__gong__marshallSliceOfPointers(ident, "NotesWhoseNodeIsExpanded", _note.GongGetIdentifier(stage)))
		}
		res = sb.String()
	case "ComplexitysWhoseNodeIsExpanded":
		var sb strings.Builder
		for _, _complexity := range diagramflossequation.ComplexitysWhoseNodeIsExpanded {
			sb.WriteString(__gong__marshallSliceOfPointers(ident, "ComplexitysWhoseNodeIsExpanded", _complexity.GongGetIdentifier(stage)))
		}
		res = sb.String()
	case "PerformancesWhoseNodeIsExpanded":
		var sb strings.Builder
		for _, _performance := range diagramflossequation.PerformancesWhoseNodeIsExpanded {
			sb.WriteString(__gong__marshallSliceOfPointers(ident, "PerformancesWhoseNodeIsExpanded", _performance.GongGetIdentifier(stage)))
		}
		res = sb.String()
	case "EffortsWhoseNodeIsExpanded":
		var sb strings.Builder
		for _, _effort := range diagramflossequation.EffortsWhoseNodeIsExpanded {
			sb.WriteString(__gong__marshallSliceOfPointers(ident, "EffortsWhoseNodeIsExpanded", _effort.GongGetIdentifier(stage)))
		}
		res = sb.String()
	default:
		log.Panicf("Unknown field %s for Gongstruct DiagramFlossEquation", fieldName)
	}
	return
}

func (effort *Effort) GongMarshallField(stage *Stage, fieldName string) (res string) {
	ident := effort.GongGetIdentifier(stage)
	_ = ident
	switch fieldName {
	case "Name":
		res = __gong__marshallString(ident, "Name", effort.Name)
	case "Strength":
		res = __gong__marshallFloat(ident, "Strength", effort.Strength)
	case "Description":
		res = __gong__marshallString(ident, "Description", effort.Description)
	case "ComputedPrefix":
		res = __gong__marshallString(ident, "ComputedPrefix", effort.ComputedPrefix)
	case "IsExpanded":
		res = __gong__marshallBool(ident, "IsExpanded", effort.IsExpanded)

	default:
		log.Panicf("Unknown field %s for Gongstruct Effort", fieldName)
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
	case "IsSystemsNodeExpanded":
		res = __gong__marshallBool(ident, "IsSystemsNodeExpanded", library.IsSystemsNodeExpanded)
	case "IsComplexitysNodeExpanded":
		res = __gong__marshallBool(ident, "IsComplexitysNodeExpanded", library.IsComplexitysNodeExpanded)
	case "IsPerformancesNodeExpanded":
		res = __gong__marshallBool(ident, "IsPerformancesNodeExpanded", library.IsPerformancesNodeExpanded)
	case "IsEffortsNodeExpanded":
		res = __gong__marshallBool(ident, "IsEffortsNodeExpanded", library.IsEffortsNodeExpanded)
	case "IsCompareAnalysisNodeExpanded":
		res = __gong__marshallBool(ident, "IsCompareAnalysisNodeExpanded", library.IsCompareAnalysisNodeExpanded)
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
	case "RootSystems":
		var sb strings.Builder
		for _, _system := range library.RootSystems {
			sb.WriteString(__gong__marshallSliceOfPointers(ident, "RootSystems", _system.GongGetIdentifier(stage)))
		}
		res = sb.String()
	case "RootComplexitys":
		var sb strings.Builder
		for _, _complexity := range library.RootComplexitys {
			sb.WriteString(__gong__marshallSliceOfPointers(ident, "RootComplexitys", _complexity.GongGetIdentifier(stage)))
		}
		res = sb.String()
	case "RootPerformances":
		var sb strings.Builder
		for _, _performance := range library.RootPerformances {
			sb.WriteString(__gong__marshallSliceOfPointers(ident, "RootPerformances", _performance.GongGetIdentifier(stage)))
		}
		res = sb.String()
	case "RootEfforts":
		var sb strings.Builder
		for _, _effort := range library.RootEfforts {
			sb.WriteString(__gong__marshallSliceOfPointers(ident, "RootEfforts", _effort.GongGetIdentifier(stage)))
		}
		res = sb.String()
	case "RootCompareAnalysis":
		var sb strings.Builder
		for _, _compareanalysis := range library.RootCompareAnalysis {
			sb.WriteString(__gong__marshallSliceOfPointers(ident, "RootCompareAnalysis", _compareanalysis.GongGetIdentifier(stage)))
		}
		res = sb.String()
	case "RootNotes":
		var sb strings.Builder
		for _, _note := range library.RootNotes {
			sb.WriteString(__gong__marshallSliceOfPointers(ident, "RootNotes", _note.GongGetIdentifier(stage)))
		}
		res = sb.String()
	case "SubLibrariesWhoseNodeIsExpanded":
		var sb strings.Builder
		for _, _library := range library.SubLibrariesWhoseNodeIsExpanded {
			sb.WriteString(__gong__marshallSliceOfPointers(ident, "SubLibrariesWhoseNodeIsExpanded", _library.GongGetIdentifier(stage)))
		}
		res = sb.String()
	case "SystemsWhoseNodeIsExpanded":
		var sb strings.Builder
		for _, _system := range library.SystemsWhoseNodeIsExpanded {
			sb.WriteString(__gong__marshallSliceOfPointers(ident, "SystemsWhoseNodeIsExpanded", _system.GongGetIdentifier(stage)))
		}
		res = sb.String()
	case "ComplexitysWhoseNodeIsExpanded":
		var sb strings.Builder
		for _, _complexity := range library.ComplexitysWhoseNodeIsExpanded {
			sb.WriteString(__gong__marshallSliceOfPointers(ident, "ComplexitysWhoseNodeIsExpanded", _complexity.GongGetIdentifier(stage)))
		}
		res = sb.String()
	case "PerformancesWhoseNodeIsExpanded":
		var sb strings.Builder
		for _, _performance := range library.PerformancesWhoseNodeIsExpanded {
			sb.WriteString(__gong__marshallSliceOfPointers(ident, "PerformancesWhoseNodeIsExpanded", _performance.GongGetIdentifier(stage)))
		}
		res = sb.String()
	case "EffortsWhoseNodeIsExpanded":
		var sb strings.Builder
		for _, _effort := range library.EffortsWhoseNodeIsExpanded {
			sb.WriteString(__gong__marshallSliceOfPointers(ident, "EffortsWhoseNodeIsExpanded", _effort.GongGetIdentifier(stage)))
		}
		res = sb.String()
	case "CompareAnalysisWhoseNodeIsExpanded":
		var sb strings.Builder
		for _, _compareanalysis := range library.CompareAnalysisWhoseNodeIsExpanded {
			sb.WriteString(__gong__marshallSliceOfPointers(ident, "CompareAnalysisWhoseNodeIsExpanded", _compareanalysis.GongGetIdentifier(stage)))
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
	case "IsComplexitysNodeExpanded":
		res = __gong__marshallBool(ident, "IsComplexitysNodeExpanded", note.IsComplexitysNodeExpanded)
	case "IsPerformancesNodeExpanded":
		res = __gong__marshallBool(ident, "IsPerformancesNodeExpanded", note.IsPerformancesNodeExpanded)
	case "IsEffortsNodeExpanded":
		res = __gong__marshallBool(ident, "IsEffortsNodeExpanded", note.IsEffortsNodeExpanded)

	case "Complexities":
		var sb strings.Builder
		for _, _complexity := range note.Complexities {
			sb.WriteString(__gong__marshallSliceOfPointers(ident, "Complexities", _complexity.GongGetIdentifier(stage)))
		}
		res = sb.String()
	case "Performances":
		var sb strings.Builder
		for _, _performance := range note.Performances {
			sb.WriteString(__gong__marshallSliceOfPointers(ident, "Performances", _performance.GongGetIdentifier(stage)))
		}
		res = sb.String()
	case "Efforts":
		var sb strings.Builder
		for _, _effort := range note.Efforts {
			sb.WriteString(__gong__marshallSliceOfPointers(ident, "Efforts", _effort.GongGetIdentifier(stage)))
		}
		res = sb.String()
	default:
		log.Panicf("Unknown field %s for Gongstruct Note", fieldName)
	}
	return
}

func (notecomplexityshape *NoteComplexityShape) GongMarshallField(stage *Stage, fieldName string) (res string) {
	ident := notecomplexityshape.GongGetIdentifier(stage)
	_ = ident
	switch fieldName {
	case "Name":
		res = __gong__marshallString(ident, "Name", notecomplexityshape.Name)
	case "StartRatio":
		res = __gong__marshallFloat(ident, "StartRatio", notecomplexityshape.StartRatio)
	case "EndRatio":
		res = __gong__marshallFloat(ident, "EndRatio", notecomplexityshape.EndRatio)
	case "StartOrientation":
		res = __gong__marshallEnumString(ident, "StartOrientation", notecomplexityshape.StartOrientation.ToCodeString())
	case "EndOrientation":
		res = __gong__marshallEnumString(ident, "EndOrientation", notecomplexityshape.EndOrientation.ToCodeString())
	case "CornerOffsetRatio":
		res = __gong__marshallFloat(ident, "CornerOffsetRatio", notecomplexityshape.CornerOffsetRatio)
	case "IsHidden":
		res = __gong__marshallBool(ident, "IsHidden", notecomplexityshape.IsHidden)
	case "IsPresentationOverridden":
		res = __gong__marshallBool(ident, "IsPresentationOverridden", notecomplexityshape.IsPresentationOverridden)
	case "Stroke":
		res = __gong__marshallString(ident, "Stroke", notecomplexityshape.Stroke)
	case "StrokeWidth":
		res = __gong__marshallFloat(ident, "StrokeWidth", notecomplexityshape.StrokeWidth)
	case "StrokeOpacity":
		res = __gong__marshallFloat(ident, "StrokeOpacity", notecomplexityshape.StrokeOpacity)
	case "StrokeDashArray":
		res = __gong__marshallString(ident, "StrokeDashArray", notecomplexityshape.StrokeDashArray)
	case "FontSize":
		res = __gong__marshallString(ident, "FontSize", notecomplexityshape.FontSize)
	case "FontColor":
		res = __gong__marshallString(ident, "FontColor", notecomplexityshape.FontColor)

	case "Note":
		if notecomplexityshape.Note != nil {
			res = __gong__marshallPointer(ident, "Note", notecomplexityshape.Note.GongGetIdentifier(stage))
		} else {
			res = __gong__marshallPointer(ident, "Note", "nil")
		}
	case "Complexity":
		if notecomplexityshape.Complexity != nil {
			res = __gong__marshallPointer(ident, "Complexity", notecomplexityshape.Complexity.GongGetIdentifier(stage))
		} else {
			res = __gong__marshallPointer(ident, "Complexity", "nil")
		}
	default:
		log.Panicf("Unknown field %s for Gongstruct NoteComplexityShape", fieldName)
	}
	return
}

func (noteeffortshape *NoteEffortShape) GongMarshallField(stage *Stage, fieldName string) (res string) {
	ident := noteeffortshape.GongGetIdentifier(stage)
	_ = ident
	switch fieldName {
	case "Name":
		res = __gong__marshallString(ident, "Name", noteeffortshape.Name)
	case "StartRatio":
		res = __gong__marshallFloat(ident, "StartRatio", noteeffortshape.StartRatio)
	case "EndRatio":
		res = __gong__marshallFloat(ident, "EndRatio", noteeffortshape.EndRatio)
	case "StartOrientation":
		res = __gong__marshallEnumString(ident, "StartOrientation", noteeffortshape.StartOrientation.ToCodeString())
	case "EndOrientation":
		res = __gong__marshallEnumString(ident, "EndOrientation", noteeffortshape.EndOrientation.ToCodeString())
	case "CornerOffsetRatio":
		res = __gong__marshallFloat(ident, "CornerOffsetRatio", noteeffortshape.CornerOffsetRatio)
	case "IsHidden":
		res = __gong__marshallBool(ident, "IsHidden", noteeffortshape.IsHidden)
	case "IsPresentationOverridden":
		res = __gong__marshallBool(ident, "IsPresentationOverridden", noteeffortshape.IsPresentationOverridden)
	case "Stroke":
		res = __gong__marshallString(ident, "Stroke", noteeffortshape.Stroke)
	case "StrokeWidth":
		res = __gong__marshallFloat(ident, "StrokeWidth", noteeffortshape.StrokeWidth)
	case "StrokeOpacity":
		res = __gong__marshallFloat(ident, "StrokeOpacity", noteeffortshape.StrokeOpacity)
	case "StrokeDashArray":
		res = __gong__marshallString(ident, "StrokeDashArray", noteeffortshape.StrokeDashArray)
	case "FontSize":
		res = __gong__marshallString(ident, "FontSize", noteeffortshape.FontSize)
	case "FontColor":
		res = __gong__marshallString(ident, "FontColor", noteeffortshape.FontColor)

	case "Note":
		if noteeffortshape.Note != nil {
			res = __gong__marshallPointer(ident, "Note", noteeffortshape.Note.GongGetIdentifier(stage))
		} else {
			res = __gong__marshallPointer(ident, "Note", "nil")
		}
	case "Effort":
		if noteeffortshape.Effort != nil {
			res = __gong__marshallPointer(ident, "Effort", noteeffortshape.Effort.GongGetIdentifier(stage))
		} else {
			res = __gong__marshallPointer(ident, "Effort", "nil")
		}
	default:
		log.Panicf("Unknown field %s for Gongstruct NoteEffortShape", fieldName)
	}
	return
}

func (noteperformanceshape *NotePerformanceShape) GongMarshallField(stage *Stage, fieldName string) (res string) {
	ident := noteperformanceshape.GongGetIdentifier(stage)
	_ = ident
	switch fieldName {
	case "Name":
		res = __gong__marshallString(ident, "Name", noteperformanceshape.Name)
	case "StartRatio":
		res = __gong__marshallFloat(ident, "StartRatio", noteperformanceshape.StartRatio)
	case "EndRatio":
		res = __gong__marshallFloat(ident, "EndRatio", noteperformanceshape.EndRatio)
	case "StartOrientation":
		res = __gong__marshallEnumString(ident, "StartOrientation", noteperformanceshape.StartOrientation.ToCodeString())
	case "EndOrientation":
		res = __gong__marshallEnumString(ident, "EndOrientation", noteperformanceshape.EndOrientation.ToCodeString())
	case "CornerOffsetRatio":
		res = __gong__marshallFloat(ident, "CornerOffsetRatio", noteperformanceshape.CornerOffsetRatio)
	case "IsHidden":
		res = __gong__marshallBool(ident, "IsHidden", noteperformanceshape.IsHidden)
	case "IsPresentationOverridden":
		res = __gong__marshallBool(ident, "IsPresentationOverridden", noteperformanceshape.IsPresentationOverridden)
	case "Stroke":
		res = __gong__marshallString(ident, "Stroke", noteperformanceshape.Stroke)
	case "StrokeWidth":
		res = __gong__marshallFloat(ident, "StrokeWidth", noteperformanceshape.StrokeWidth)
	case "StrokeOpacity":
		res = __gong__marshallFloat(ident, "StrokeOpacity", noteperformanceshape.StrokeOpacity)
	case "StrokeDashArray":
		res = __gong__marshallString(ident, "StrokeDashArray", noteperformanceshape.StrokeDashArray)
	case "FontSize":
		res = __gong__marshallString(ident, "FontSize", noteperformanceshape.FontSize)
	case "FontColor":
		res = __gong__marshallString(ident, "FontColor", noteperformanceshape.FontColor)

	case "Note":
		if noteperformanceshape.Note != nil {
			res = __gong__marshallPointer(ident, "Note", noteperformanceshape.Note.GongGetIdentifier(stage))
		} else {
			res = __gong__marshallPointer(ident, "Note", "nil")
		}
	case "Performance":
		if noteperformanceshape.Performance != nil {
			res = __gong__marshallPointer(ident, "Performance", noteperformanceshape.Performance.GongGetIdentifier(stage))
		} else {
			res = __gong__marshallPointer(ident, "Performance", "nil")
		}
	default:
		log.Panicf("Unknown field %s for Gongstruct NotePerformanceShape", fieldName)
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

func (performance *Performance) GongMarshallField(stage *Stage, fieldName string) (res string) {
	ident := performance.GongGetIdentifier(stage)
	_ = ident
	switch fieldName {
	case "Name":
		res = __gong__marshallString(ident, "Name", performance.Name)
	case "Strength":
		res = __gong__marshallFloat(ident, "Strength", performance.Strength)
	case "Description":
		res = __gong__marshallString(ident, "Description", performance.Description)
	case "ComputedPrefix":
		res = __gong__marshallString(ident, "ComputedPrefix", performance.ComputedPrefix)
	case "IsExpanded":
		res = __gong__marshallBool(ident, "IsExpanded", performance.IsExpanded)

	default:
		log.Panicf("Unknown field %s for Gongstruct Performance", fieldName)
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
	case "AreCPEsCompoundedFromSubSystems":
		res = __gong__marshallBool(ident, "AreCPEsCompoundedFromSubSystems", system.AreCPEsCompoundedFromSubSystems)
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
	case "IsComplexitysNodeExpanded":
		res = __gong__marshallBool(ident, "IsComplexitysNodeExpanded", system.IsComplexitysNodeExpanded)
	case "IsPerformancesNodeExpanded":
		res = __gong__marshallBool(ident, "IsPerformancesNodeExpanded", system.IsPerformancesNodeExpanded)
	case "IsEffortsNodeExpanded":
		res = __gong__marshallBool(ident, "IsEffortsNodeExpanded", system.IsEffortsNodeExpanded)

	case "Complexities":
		var sb strings.Builder
		for _, _complexity := range system.Complexities {
			sb.WriteString(__gong__marshallSliceOfPointers(ident, "Complexities", _complexity.GongGetIdentifier(stage)))
		}
		res = sb.String()
	case "Performances":
		var sb strings.Builder
		for _, _performance := range system.Performances {
			sb.WriteString(__gong__marshallSliceOfPointers(ident, "Performances", _performance.GongGetIdentifier(stage)))
		}
		res = sb.String()
	case "Efforts":
		var sb strings.Builder
		for _, _effort := range system.Efforts {
			sb.WriteString(__gong__marshallSliceOfPointers(ident, "Efforts", _effort.GongGetIdentifier(stage)))
		}
		res = sb.String()
	case "SubSystems":
		var sb strings.Builder
		for _, _system := range system.SubSystems {
			sb.WriteString(__gong__marshallSliceOfPointers(ident, "SubSystems", _system.GongGetIdentifier(stage)))
		}
		res = sb.String()
	case "DiagramFlossEquations":
		var sb strings.Builder
		for _, _diagramflossequation := range system.DiagramFlossEquations {
			sb.WriteString(__gong__marshallSliceOfPointers(ident, "DiagramFlossEquations", _diagramflossequation.GongGetIdentifier(stage)))
		}
		res = sb.String()
	case "DiagramFlossEquationsWhoseNodeIsExpanded":
		var sb strings.Builder
		for _, _diagramflossequation := range system.DiagramFlossEquationsWhoseNodeIsExpanded {
			sb.WriteString(__gong__marshallSliceOfPointers(ident, "DiagramFlossEquationsWhoseNodeIsExpanded", _diagramflossequation.GongGetIdentifier(stage)))
		}
		res = sb.String()
	case "ComplexitysWhoseNodeIsExpanded":
		var sb strings.Builder
		for _, _complexity := range system.ComplexitysWhoseNodeIsExpanded {
			sb.WriteString(__gong__marshallSliceOfPointers(ident, "ComplexitysWhoseNodeIsExpanded", _complexity.GongGetIdentifier(stage)))
		}
		res = sb.String()
	case "PerformancesWhoseNodeIsExpanded":
		var sb strings.Builder
		for _, _performance := range system.PerformancesWhoseNodeIsExpanded {
			sb.WriteString(__gong__marshallSliceOfPointers(ident, "PerformancesWhoseNodeIsExpanded", _performance.GongGetIdentifier(stage)))
		}
		res = sb.String()
	case "EffortsWhoseNodeIsExpanded":
		var sb strings.Builder
		for _, _effort := range system.EffortsWhoseNodeIsExpanded {
			sb.WriteString(__gong__marshallSliceOfPointers(ident, "EffortsWhoseNodeIsExpanded", _effort.GongGetIdentifier(stage)))
		}
		res = sb.String()
	default:
		log.Panicf("Unknown field %s for Gongstruct System", fieldName)
	}
	return
}

// insertion point for marshall all fields methods
func (compareanalysis *CompareAnalysis) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(compareanalysis.GongMarshallField(stage, "Name"))
		pointersInitializesStatements.WriteString(compareanalysis.GongMarshallField(stage, "FromSystem"))
		pointersInitializesStatements.WriteString(compareanalysis.GongMarshallField(stage, "ToSystem"))
		initializerStatements.WriteString(compareanalysis.GongMarshallField(stage, "Mu"))
		initializerStatements.WriteString(compareanalysis.GongMarshallField(stage, "Epsilon"))
		pointersInitializesStatements.WriteString(compareanalysis.GongMarshallField(stage, "DiagramFlossEquations"))
		pointersInitializesStatements.WriteString(compareanalysis.GongMarshallField(stage, "DiagramFlossEquationsWhoseNodeIsExpanded"))
		initializerStatements.WriteString(compareanalysis.GongMarshallField(stage, "ComputedPrefix"))
		initializerStatements.WriteString(compareanalysis.GongMarshallField(stage, "IsExpanded"))
	}
	initRes = initializerStatements.String()
	ptrRes = pointersInitializesStatements.String()
	return
}
func (complexity *Complexity) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(complexity.GongMarshallField(stage, "Name"))
		initializerStatements.WriteString(complexity.GongMarshallField(stage, "Strength"))
		initializerStatements.WriteString(complexity.GongMarshallField(stage, "Description"))
		initializerStatements.WriteString(complexity.GongMarshallField(stage, "ComputedPrefix"))
		initializerStatements.WriteString(complexity.GongMarshallField(stage, "IsExpanded"))
	}
	initRes = initializerStatements.String()
	ptrRes = pointersInitializesStatements.String()
	return
}
func (diagramflossequation *DiagramFlossEquation) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(diagramflossequation.GongMarshallField(stage, "Name"))
		initializerStatements.WriteString(diagramflossequation.GongMarshallField(stage, "Description"))
		initializerStatements.WriteString(diagramflossequation.GongMarshallField(stage, "Scale"))
		initializerStatements.WriteString(diagramflossequation.GongMarshallField(stage, "FontSize"))
		initializerStatements.WriteString(diagramflossequation.GongMarshallField(stage, "ComputedPrefix"))
		initializerStatements.WriteString(diagramflossequation.GongMarshallField(stage, "IsExpanded"))
		initializerStatements.WriteString(diagramflossequation.GongMarshallField(stage, "IsChecked"))
		initializerStatements.WriteString(diagramflossequation.GongMarshallField(stage, "IsEditable_"))
		initializerStatements.WriteString(diagramflossequation.GongMarshallField(stage, "IsInDelta3ColumnsMode"))
		initializerStatements.WriteString(diagramflossequation.GongMarshallField(stage, "AreQuantitativeElementsVisible"))
		initializerStatements.WriteString(diagramflossequation.GongMarshallField(stage, "AreSubsystemsVisible"))
		initializerStatements.WriteString(diagramflossequation.GongMarshallField(stage, "AreCommonElementsHidden"))
		initializerStatements.WriteString(diagramflossequation.GongMarshallField(stage, "AreCPEArrowsVisible"))
		initializerStatements.WriteString(diagramflossequation.GongMarshallField(stage, "AreColumnTitlesVisible"))
		initializerStatements.WriteString(diagramflossequation.GongMarshallField(stage, "Width"))
		initializerStatements.WriteString(diagramflossequation.GongMarshallField(stage, "Height"))
		initializerStatements.WriteString(diagramflossequation.GongMarshallField(stage, "DefaultBoxWidth"))
		initializerStatements.WriteString(diagramflossequation.GongMarshallField(stage, "DefaultBoxHeigth"))
		pointersInitializesStatements.WriteString(diagramflossequation.GongMarshallField(stage, "Note_Shapes"))
		pointersInitializesStatements.WriteString(diagramflossequation.GongMarshallField(stage, "NoteComplexityShapes"))
		pointersInitializesStatements.WriteString(diagramflossequation.GongMarshallField(stage, "NotePerformanceShapes"))
		pointersInitializesStatements.WriteString(diagramflossequation.GongMarshallField(stage, "NoteEffortShapes"))
		initializerStatements.WriteString(diagramflossequation.GongMarshallField(stage, "IsNotesNodeExpanded"))
		pointersInitializesStatements.WriteString(diagramflossequation.GongMarshallField(stage, "NotesWhoseNodeIsExpanded"))
		initializerStatements.WriteString(diagramflossequation.GongMarshallField(stage, "IsComplexitysNodeExpanded"))
		pointersInitializesStatements.WriteString(diagramflossequation.GongMarshallField(stage, "ComplexitysWhoseNodeIsExpanded"))
		initializerStatements.WriteString(diagramflossequation.GongMarshallField(stage, "IsPerformancesNodeExpanded"))
		pointersInitializesStatements.WriteString(diagramflossequation.GongMarshallField(stage, "PerformancesWhoseNodeIsExpanded"))
		initializerStatements.WriteString(diagramflossequation.GongMarshallField(stage, "IsEffortsNodeExpanded"))
		pointersInitializesStatements.WriteString(diagramflossequation.GongMarshallField(stage, "EffortsWhoseNodeIsExpanded"))
	}
	initRes = initializerStatements.String()
	ptrRes = pointersInitializesStatements.String()
	return
}
func (effort *Effort) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(effort.GongMarshallField(stage, "Name"))
		initializerStatements.WriteString(effort.GongMarshallField(stage, "Strength"))
		initializerStatements.WriteString(effort.GongMarshallField(stage, "Description"))
		initializerStatements.WriteString(effort.GongMarshallField(stage, "ComputedPrefix"))
		initializerStatements.WriteString(effort.GongMarshallField(stage, "IsExpanded"))
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
		pointersInitializesStatements.WriteString(library.GongMarshallField(stage, "SubLibraries"))
		pointersInitializesStatements.WriteString(library.GongMarshallField(stage, "RootSystems"))
		pointersInitializesStatements.WriteString(library.GongMarshallField(stage, "RootComplexitys"))
		pointersInitializesStatements.WriteString(library.GongMarshallField(stage, "RootPerformances"))
		pointersInitializesStatements.WriteString(library.GongMarshallField(stage, "RootEfforts"))
		pointersInitializesStatements.WriteString(library.GongMarshallField(stage, "RootCompareAnalysis"))
		pointersInitializesStatements.WriteString(library.GongMarshallField(stage, "RootNotes"))
		initializerStatements.WriteString(library.GongMarshallField(stage, "IsRootLibrary"))
		initializerStatements.WriteString(library.GongMarshallField(stage, "IsSubLibrariesNodeExpanded"))
		pointersInitializesStatements.WriteString(library.GongMarshallField(stage, "SubLibrariesWhoseNodeIsExpanded"))
		initializerStatements.WriteString(library.GongMarshallField(stage, "NbPixPerCharacter"))
		initializerStatements.WriteString(library.GongMarshallField(stage, "LogoSVGFile"))
		initializerStatements.WriteString(library.GongMarshallField(stage, "IsSystemsNodeExpanded"))
		pointersInitializesStatements.WriteString(library.GongMarshallField(stage, "SystemsWhoseNodeIsExpanded"))
		initializerStatements.WriteString(library.GongMarshallField(stage, "IsComplexitysNodeExpanded"))
		pointersInitializesStatements.WriteString(library.GongMarshallField(stage, "ComplexitysWhoseNodeIsExpanded"))
		initializerStatements.WriteString(library.GongMarshallField(stage, "IsPerformancesNodeExpanded"))
		pointersInitializesStatements.WriteString(library.GongMarshallField(stage, "PerformancesWhoseNodeIsExpanded"))
		initializerStatements.WriteString(library.GongMarshallField(stage, "IsEffortsNodeExpanded"))
		pointersInitializesStatements.WriteString(library.GongMarshallField(stage, "EffortsWhoseNodeIsExpanded"))
		initializerStatements.WriteString(library.GongMarshallField(stage, "IsCompareAnalysisNodeExpanded"))
		pointersInitializesStatements.WriteString(library.GongMarshallField(stage, "CompareAnalysisWhoseNodeIsExpanded"))
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
		pointersInitializesStatements.WriteString(note.GongMarshallField(stage, "Complexities"))
		pointersInitializesStatements.WriteString(note.GongMarshallField(stage, "Performances"))
		pointersInitializesStatements.WriteString(note.GongMarshallField(stage, "Efforts"))
		initializerStatements.WriteString(note.GongMarshallField(stage, "ComputedPrefix"))
		initializerStatements.WriteString(note.GongMarshallField(stage, "IsExpanded"))
		initializerStatements.WriteString(note.GongMarshallField(stage, "IsComplexitysNodeExpanded"))
		initializerStatements.WriteString(note.GongMarshallField(stage, "IsPerformancesNodeExpanded"))
		initializerStatements.WriteString(note.GongMarshallField(stage, "IsEffortsNodeExpanded"))
	}
	initRes = initializerStatements.String()
	ptrRes = pointersInitializesStatements.String()
	return
}
func (notecomplexityshape *NoteComplexityShape) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(notecomplexityshape.GongMarshallField(stage, "Name"))
		pointersInitializesStatements.WriteString(notecomplexityshape.GongMarshallField(stage, "Note"))
		pointersInitializesStatements.WriteString(notecomplexityshape.GongMarshallField(stage, "Complexity"))
		initializerStatements.WriteString(notecomplexityshape.GongMarshallField(stage, "StartRatio"))
		initializerStatements.WriteString(notecomplexityshape.GongMarshallField(stage, "EndRatio"))
		initializerStatements.WriteString(notecomplexityshape.GongMarshallField(stage, "StartOrientation"))
		initializerStatements.WriteString(notecomplexityshape.GongMarshallField(stage, "EndOrientation"))
		initializerStatements.WriteString(notecomplexityshape.GongMarshallField(stage, "CornerOffsetRatio"))
		initializerStatements.WriteString(notecomplexityshape.GongMarshallField(stage, "IsHidden"))
		initializerStatements.WriteString(notecomplexityshape.GongMarshallField(stage, "IsPresentationOverridden"))
		initializerStatements.WriteString(notecomplexityshape.GongMarshallField(stage, "Stroke"))
		initializerStatements.WriteString(notecomplexityshape.GongMarshallField(stage, "StrokeWidth"))
		initializerStatements.WriteString(notecomplexityshape.GongMarshallField(stage, "StrokeOpacity"))
		initializerStatements.WriteString(notecomplexityshape.GongMarshallField(stage, "StrokeDashArray"))
		initializerStatements.WriteString(notecomplexityshape.GongMarshallField(stage, "FontSize"))
		initializerStatements.WriteString(notecomplexityshape.GongMarshallField(stage, "FontColor"))
	}
	initRes = initializerStatements.String()
	ptrRes = pointersInitializesStatements.String()
	return
}
func (noteeffortshape *NoteEffortShape) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(noteeffortshape.GongMarshallField(stage, "Name"))
		pointersInitializesStatements.WriteString(noteeffortshape.GongMarshallField(stage, "Note"))
		pointersInitializesStatements.WriteString(noteeffortshape.GongMarshallField(stage, "Effort"))
		initializerStatements.WriteString(noteeffortshape.GongMarshallField(stage, "StartRatio"))
		initializerStatements.WriteString(noteeffortshape.GongMarshallField(stage, "EndRatio"))
		initializerStatements.WriteString(noteeffortshape.GongMarshallField(stage, "StartOrientation"))
		initializerStatements.WriteString(noteeffortshape.GongMarshallField(stage, "EndOrientation"))
		initializerStatements.WriteString(noteeffortshape.GongMarshallField(stage, "CornerOffsetRatio"))
		initializerStatements.WriteString(noteeffortshape.GongMarshallField(stage, "IsHidden"))
		initializerStatements.WriteString(noteeffortshape.GongMarshallField(stage, "IsPresentationOverridden"))
		initializerStatements.WriteString(noteeffortshape.GongMarshallField(stage, "Stroke"))
		initializerStatements.WriteString(noteeffortshape.GongMarshallField(stage, "StrokeWidth"))
		initializerStatements.WriteString(noteeffortshape.GongMarshallField(stage, "StrokeOpacity"))
		initializerStatements.WriteString(noteeffortshape.GongMarshallField(stage, "StrokeDashArray"))
		initializerStatements.WriteString(noteeffortshape.GongMarshallField(stage, "FontSize"))
		initializerStatements.WriteString(noteeffortshape.GongMarshallField(stage, "FontColor"))
	}
	initRes = initializerStatements.String()
	ptrRes = pointersInitializesStatements.String()
	return
}
func (noteperformanceshape *NotePerformanceShape) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(noteperformanceshape.GongMarshallField(stage, "Name"))
		pointersInitializesStatements.WriteString(noteperformanceshape.GongMarshallField(stage, "Note"))
		pointersInitializesStatements.WriteString(noteperformanceshape.GongMarshallField(stage, "Performance"))
		initializerStatements.WriteString(noteperformanceshape.GongMarshallField(stage, "StartRatio"))
		initializerStatements.WriteString(noteperformanceshape.GongMarshallField(stage, "EndRatio"))
		initializerStatements.WriteString(noteperformanceshape.GongMarshallField(stage, "StartOrientation"))
		initializerStatements.WriteString(noteperformanceshape.GongMarshallField(stage, "EndOrientation"))
		initializerStatements.WriteString(noteperformanceshape.GongMarshallField(stage, "CornerOffsetRatio"))
		initializerStatements.WriteString(noteperformanceshape.GongMarshallField(stage, "IsHidden"))
		initializerStatements.WriteString(noteperformanceshape.GongMarshallField(stage, "IsPresentationOverridden"))
		initializerStatements.WriteString(noteperformanceshape.GongMarshallField(stage, "Stroke"))
		initializerStatements.WriteString(noteperformanceshape.GongMarshallField(stage, "StrokeWidth"))
		initializerStatements.WriteString(noteperformanceshape.GongMarshallField(stage, "StrokeOpacity"))
		initializerStatements.WriteString(noteperformanceshape.GongMarshallField(stage, "StrokeDashArray"))
		initializerStatements.WriteString(noteperformanceshape.GongMarshallField(stage, "FontSize"))
		initializerStatements.WriteString(noteperformanceshape.GongMarshallField(stage, "FontColor"))
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
func (performance *Performance) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(performance.GongMarshallField(stage, "Name"))
		initializerStatements.WriteString(performance.GongMarshallField(stage, "Strength"))
		initializerStatements.WriteString(performance.GongMarshallField(stage, "Description"))
		initializerStatements.WriteString(performance.GongMarshallField(stage, "ComputedPrefix"))
		initializerStatements.WriteString(performance.GongMarshallField(stage, "IsExpanded"))
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
		pointersInitializesStatements.WriteString(system.GongMarshallField(stage, "Complexities"))
		pointersInitializesStatements.WriteString(system.GongMarshallField(stage, "Performances"))
		pointersInitializesStatements.WriteString(system.GongMarshallField(stage, "Efforts"))
		pointersInitializesStatements.WriteString(system.GongMarshallField(stage, "SubSystems"))
		initializerStatements.WriteString(system.GongMarshallField(stage, "AreCPEsCompoundedFromSubSystems"))
		initializerStatements.WriteString(system.GongMarshallField(stage, "ComputedPrefix"))
		initializerStatements.WriteString(system.GongMarshallField(stage, "IsExpanded"))
		initializerStatements.WriteString(system.GongMarshallField(stage, "SVG_Path"))
		initializerStatements.WriteString(system.GongMarshallField(stage, "InverseAppliedScaling"))
		pointersInitializesStatements.WriteString(system.GongMarshallField(stage, "DiagramFlossEquations"))
		pointersInitializesStatements.WriteString(system.GongMarshallField(stage, "DiagramFlossEquationsWhoseNodeIsExpanded"))
		initializerStatements.WriteString(system.GongMarshallField(stage, "IsSubSystemNodeExpanded"))
		initializerStatements.WriteString(system.GongMarshallField(stage, "IsComplexitysNodeExpanded"))
		pointersInitializesStatements.WriteString(system.GongMarshallField(stage, "ComplexitysWhoseNodeIsExpanded"))
		initializerStatements.WriteString(system.GongMarshallField(stage, "IsPerformancesNodeExpanded"))
		pointersInitializesStatements.WriteString(system.GongMarshallField(stage, "PerformancesWhoseNodeIsExpanded"))
		initializerStatements.WriteString(system.GongMarshallField(stage, "IsEffortsNodeExpanded"))
		pointersInitializesStatements.WriteString(system.GongMarshallField(stage, "EffortsWhoseNodeIsExpanded"))
	}
	initRes = initializerStatements.String()
	ptrRes = pointersInitializesStatements.String()
	return
}
