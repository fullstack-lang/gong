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
	modelsPackageName := "github.com/fullstack-lang/gong/dsm/statemachines/go/models"
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
		for _, action := range __gong__sortStageSetInstances(stage.Actions, stage.Action_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			actionIdent := action.GongGetIdentifier(stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&models.Action{Name: %s}).Stage(stage)", actionIdent, __gong__toRawStringLiteral(action.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", actionIdent, __gong__toRawStringLiteral(action.Name)))
			if action.Criticality.ToCodeString() != "" {
				values.WriteString(fmt.Sprintf("\n\t%s.Criticality = models.%s", actionIdent, action.Criticality.ToCodeString()))
			} else {
				values.WriteString(fmt.Sprintf("\n\t%s.Criticality = \"\"", actionIdent))
			}
		}
	}
	if stage != nil {
		for _, activities := range __gong__sortStageSetInstances(stage.Activitiess, stage.Activities_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			activitiesIdent := activities.GongGetIdentifier(stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&models.Activities{Name: %s}).Stage(stage)", activitiesIdent, __gong__toRawStringLiteral(activities.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", activitiesIdent, __gong__toRawStringLiteral(activities.Name)))
			if activities.Criticality.ToCodeString() != "" {
				values.WriteString(fmt.Sprintf("\n\t%s.Criticality = models.%s", activitiesIdent, activities.Criticality.ToCodeString()))
			} else {
				values.WriteString(fmt.Sprintf("\n\t%s.Criticality = \"\"", activitiesIdent))
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
			values.WriteString(fmt.Sprintf("\n\t%s.IsChecked = %t", diagramIdent, diagram.IsChecked))
			values.WriteString(fmt.Sprintf("\n\t%s.IsExpanded = %t", diagramIdent, diagram.IsExpanded))
			values.WriteString(fmt.Sprintf("\n\t%s.IsEditable_ = %t", diagramIdent, diagram.IsEditable_))
			values.WriteString(fmt.Sprintf("\n\t%s.IsStatesNodeExpanded = %t", diagramIdent, diagram.IsStatesNodeExpanded))
			values.WriteString(fmt.Sprintf("\n\t%s.ShowRoles = %t", diagramIdent, diagram.ShowRoles))
			values.WriteString(fmt.Sprintf("\n\t%s.ShowMessages = %t", diagramIdent, diagram.ShowMessages))
			for _, elem := range diagram.State_Shapes {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := elem.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.State_Shapes = append(%s.State_Shapes, %s)", diagramIdent, diagramIdent, targetIdent))
			}
			for _, elem := range diagram.StatesWhoseNodeIsExpanded {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := elem.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.StatesWhoseNodeIsExpanded = append(%s.StatesWhoseNodeIsExpanded, %s)", diagramIdent, diagramIdent, targetIdent))
			}
			for _, elem := range diagram.Transition_Shapes {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := elem.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Transition_Shapes = append(%s.Transition_Shapes, %s)", diagramIdent, diagramIdent, targetIdent))
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
			for _, elem := range diagram.NoteState_Shapes {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := elem.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.NoteState_Shapes = append(%s.NoteState_Shapes, %s)", diagramIdent, diagramIdent, targetIdent))
			}
		}
	}
	if stage != nil {
		for _, guard := range __gong__sortStageSetInstances(stage.Guards, stage.Guard_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			guardIdent := guard.GongGetIdentifier(stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&models.Guard{Name: %s}).Stage(stage)", guardIdent, __gong__toRawStringLiteral(guard.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", guardIdent, __gong__toRawStringLiteral(guard.Name)))
		}
	}
	if stage != nil {
		for _, kill := range __gong__sortStageSetInstances(stage.Kills, stage.Kill_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			killIdent := kill.GongGetIdentifier(stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&models.Kill{Name: %s}).Stage(stage)", killIdent, __gong__toRawStringLiteral(kill.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", killIdent, __gong__toRawStringLiteral(kill.Name)))
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
			values.WriteString(fmt.Sprintf("\n\t%s.NbPixPerCharacter = %f", libraryIdent, library.NbPixPerCharacter))
			values.WriteString(fmt.Sprintf("\n\t%s.LogoSVGFile = %s", libraryIdent, __gong__toRawStringLiteral(library.LogoSVGFile)))
			values.WriteString(fmt.Sprintf("\n\t%s.ComputedPrefix = %s", libraryIdent, __gong__toRawStringLiteral(library.ComputedPrefix)))
			values.WriteString(fmt.Sprintf("\n\t%s.IsExpanded = %t", libraryIdent, library.IsExpanded))
			values.WriteString(fmt.Sprintf("\n\t%s.IsRootLibrary = %t", libraryIdent, library.IsRootLibrary))
			values.WriteString(fmt.Sprintf("\n\t%s.IsStateMachinesNodeExpanded = %t", libraryIdent, library.IsStateMachinesNodeExpanded))
			values.WriteString(fmt.Sprintf("\n\t%s.IsSubLibrariesNodeExpanded = %t", libraryIdent, library.IsSubLibrariesNodeExpanded))
			values.WriteString(fmt.Sprintf("\n\t%s.IsExpandedTmp = %t", libraryIdent, library.IsExpandedTmp))
			values.WriteString(fmt.Sprintf("\n\t%s.IsRolesNodeExpanded = %t", libraryIdent, library.IsRolesNodeExpanded))
			values.WriteString(fmt.Sprintf("\n\t%s.IsMessageTypesNodeExpanded = %t", libraryIdent, library.IsMessageTypesNodeExpanded))
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
			for _, elem := range library.RootStateMachines {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := elem.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.RootStateMachines = append(%s.RootStateMachines, %s)", libraryIdent, libraryIdent, targetIdent))
			}
			for _, elem := range library.StateMachinesWhoseNodeIsExpanded {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := elem.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.StateMachinesWhoseNodeIsExpanded = append(%s.StateMachinesWhoseNodeIsExpanded, %s)", libraryIdent, libraryIdent, targetIdent))
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
			for _, elem := range library.Roles {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := elem.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Roles = append(%s.Roles, %s)", libraryIdent, libraryIdent, targetIdent))
			}
			for _, elem := range library.MessageTypes {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := elem.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.MessageTypes = append(%s.MessageTypes, %s)", libraryIdent, libraryIdent, targetIdent))
			}
		}
	}
	if stage != nil {
		for _, message := range __gong__sortStageSetInstances(stage.Messages, stage.Message_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			messageIdent := message.GongGetIdentifier(stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&models.Message{Name: %s}).Stage(stage)", messageIdent, __gong__toRawStringLiteral(message.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", messageIdent, __gong__toRawStringLiteral(message.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.IsSelected = %t", messageIdent, message.IsSelected))
			if message.MessageType != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := message.MessageType.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.MessageType = %s", messageIdent, targetIdent))
			}
			if message.OriginTransition != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := message.OriginTransition.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.OriginTransition = %s", messageIdent, targetIdent))
			}
		}
	}
	if stage != nil {
		for _, messagetype := range __gong__sortStageSetInstances(stage.MessageTypes, stage.MessageType_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			messagetypeIdent := messagetype.GongGetIdentifier(stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&models.MessageType{Name: %s}).Stage(stage)", messagetypeIdent, __gong__toRawStringLiteral(messagetype.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", messagetypeIdent, __gong__toRawStringLiteral(messagetype.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Description = %s", messagetypeIdent, __gong__toRawStringLiteral(messagetype.Description)))
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
			if note.State != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := note.State.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.State = %s", noteIdent, targetIdent))
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
				targetIdent := noteshape.Note.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Note = %s", noteshapeIdent, targetIdent))
			}
		}
	}
	if stage != nil {
		for _, notestateshape := range __gong__sortStageSetInstances(stage.NoteStateShapes, stage.NoteStateShape_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			notestateshapeIdent := notestateshape.GongGetIdentifier(stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&models.NoteStateShape{Name: %s}).Stage(stage)", notestateshapeIdent, __gong__toRawStringLiteral(notestateshape.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", notestateshapeIdent, __gong__toRawStringLiteral(notestateshape.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.StartRatio = %f", notestateshapeIdent, notestateshape.StartRatio))
			values.WriteString(fmt.Sprintf("\n\t%s.EndRatio = %f", notestateshapeIdent, notestateshape.EndRatio))
			if notestateshape.StartOrientation.ToCodeString() != "" {
				values.WriteString(fmt.Sprintf("\n\t%s.StartOrientation = models.%s", notestateshapeIdent, notestateshape.StartOrientation.ToCodeString()))
			} else {
				values.WriteString(fmt.Sprintf("\n\t%s.StartOrientation = \"\"", notestateshapeIdent))
			}
			if notestateshape.EndOrientation.ToCodeString() != "" {
				values.WriteString(fmt.Sprintf("\n\t%s.EndOrientation = models.%s", notestateshapeIdent, notestateshape.EndOrientation.ToCodeString()))
			} else {
				values.WriteString(fmt.Sprintf("\n\t%s.EndOrientation = \"\"", notestateshapeIdent))
			}
			values.WriteString(fmt.Sprintf("\n\t%s.CornerOffsetRatio = %f", notestateshapeIdent, notestateshape.CornerOffsetRatio))
			values.WriteString(fmt.Sprintf("\n\t%s.IsHidden = %t", notestateshapeIdent, notestateshape.IsHidden))
			values.WriteString(fmt.Sprintf("\n\t%s.IsPresentationOverridden = %t", notestateshapeIdent, notestateshape.IsPresentationOverridden))
			values.WriteString(fmt.Sprintf("\n\t%s.Stroke = %s", notestateshapeIdent, __gong__toRawStringLiteral(notestateshape.Stroke)))
			values.WriteString(fmt.Sprintf("\n\t%s.StrokeWidth = %f", notestateshapeIdent, notestateshape.StrokeWidth))
			values.WriteString(fmt.Sprintf("\n\t%s.StrokeOpacity = %f", notestateshapeIdent, notestateshape.StrokeOpacity))
			values.WriteString(fmt.Sprintf("\n\t%s.StrokeDashArray = %s", notestateshapeIdent, __gong__toRawStringLiteral(notestateshape.StrokeDashArray)))
			values.WriteString(fmt.Sprintf("\n\t%s.FontSize = %s", notestateshapeIdent, __gong__toRawStringLiteral(notestateshape.FontSize)))
			values.WriteString(fmt.Sprintf("\n\t%s.FontColor = %s", notestateshapeIdent, __gong__toRawStringLiteral(notestateshape.FontColor)))
			if notestateshape.Note != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := notestateshape.Note.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Note = %s", notestateshapeIdent, targetIdent))
			}
			if notestateshape.State != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := notestateshape.State.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.State = %s", notestateshapeIdent, targetIdent))
			}
		}
	}
	if stage != nil {
		for _, object := range __gong__sortStageSetInstances(stage.Objects, stage.Object_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			objectIdent := object.GongGetIdentifier(stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&models.Object{Name: %s}).Stage(stage)", objectIdent, __gong__toRawStringLiteral(object.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", objectIdent, __gong__toRawStringLiteral(object.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.IsSelected = %t", objectIdent, object.IsSelected))
			values.WriteString(fmt.Sprintf("\n\t%s.Rank = %d", objectIdent, object.Rank))
			values.WriteString(fmt.Sprintf("\n\t%s.DOF, _ = time.Parse(\"2006-01-02 15:04:05.999999999 -0700 MST\", \"%s\")", objectIdent, object.DOF.String()))
			if object.State != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := object.State.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.State = %s", objectIdent, targetIdent))
			}
			for _, elem := range object.Messages {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := elem.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Messages = append(%s.Messages, %s)", objectIdent, objectIdent, targetIdent))
			}
		}
	}
	if stage != nil {
		for _, role := range __gong__sortStageSetInstances(stage.Roles, stage.Role_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			roleIdent := role.GongGetIdentifier(stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&models.Role{Name: %s}).Stage(stage)", roleIdent, __gong__toRawStringLiteral(role.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", roleIdent, __gong__toRawStringLiteral(role.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Acronym = %s", roleIdent, __gong__toRawStringLiteral(role.Acronym)))
			for _, elem := range role.RolesWithSamePermissions {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := elem.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.RolesWithSamePermissions = append(%s.RolesWithSamePermissions, %s)", roleIdent, roleIdent, targetIdent))
			}
		}
	}
	if stage != nil {
		for _, state := range __gong__sortStageSetInstances(stage.States, stage.State_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			stateIdent := state.GongGetIdentifier(stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&models.State{Name: %s}).Stage(stage)", stateIdent, __gong__toRawStringLiteral(state.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", stateIdent, __gong__toRawStringLiteral(state.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.IsEndState = %t", stateIdent, state.IsEndState))
			values.WriteString(fmt.Sprintf("\n\t%s.IsDecisionNode = %t", stateIdent, state.IsDecisionNode))
			values.WriteString(fmt.Sprintf("\n\t%s.IsFictious = %t", stateIdent, state.IsFictious))
			for _, elem := range state.SubStates {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := elem.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.SubStates = append(%s.SubStates, %s)", stateIdent, stateIdent, targetIdent))
			}
			if state.Entry != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := state.Entry.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Entry = %s", stateIdent, targetIdent))
			}
			for _, elem := range state.Activities {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := elem.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Activities = append(%s.Activities, %s)", stateIdent, stateIdent, targetIdent))
			}
			if state.Exit != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := state.Exit.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Exit = %s", stateIdent, targetIdent))
			}
			if state.Parent != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := state.Parent.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Parent = %s", stateIdent, targetIdent))
			}
			for _, elem := range state.Diagrams {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := elem.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Diagrams = append(%s.Diagrams, %s)", stateIdent, stateIdent, targetIdent))
			}
			for _, elem := range state.Notes {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := elem.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Notes = append(%s.Notes, %s)", stateIdent, stateIdent, targetIdent))
			}
		}
	}
	if stage != nil {
		for _, statemachine := range __gong__sortStageSetInstances(stage.StateMachines, stage.StateMachine_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			statemachineIdent := statemachine.GongGetIdentifier(stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&models.StateMachine{Name: %s}).Stage(stage)", statemachineIdent, __gong__toRawStringLiteral(statemachine.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", statemachineIdent, __gong__toRawStringLiteral(statemachine.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.IsWithTransitionNameAutonamticalyGenerated = %t", statemachineIdent, statemachine.IsWithTransitionNameAutonamticalyGenerated))
			values.WriteString(fmt.Sprintf("\n\t%s.ComputedPrefix = %s", statemachineIdent, __gong__toRawStringLiteral(statemachine.ComputedPrefix)))
			values.WriteString(fmt.Sprintf("\n\t%s.IsExpanded = %t", statemachineIdent, statemachine.IsExpanded))
			if statemachine.InitialState != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := statemachine.InitialState.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.InitialState = %s", statemachineIdent, targetIdent))
			}
			for _, elem := range statemachine.States {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := elem.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.States = append(%s.States, %s)", statemachineIdent, statemachineIdent, targetIdent))
			}
			for _, elem := range statemachine.Diagrams {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := elem.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Diagrams = append(%s.Diagrams, %s)", statemachineIdent, statemachineIdent, targetIdent))
			}
		}
	}
	if stage != nil {
		for _, stateshape := range __gong__sortStageSetInstances(stage.StateShapes, stage.StateShape_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			stateshapeIdent := stateshape.GongGetIdentifier(stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&models.StateShape{Name: %s}).Stage(stage)", stateshapeIdent, __gong__toRawStringLiteral(stateshape.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", stateshapeIdent, __gong__toRawStringLiteral(stateshape.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.X = %f", stateshapeIdent, stateshape.X))
			values.WriteString(fmt.Sprintf("\n\t%s.Y = %f", stateshapeIdent, stateshape.Y))
			values.WriteString(fmt.Sprintf("\n\t%s.Width = %f", stateshapeIdent, stateshape.Width))
			values.WriteString(fmt.Sprintf("\n\t%s.Height = %f", stateshapeIdent, stateshape.Height))
			values.WriteString(fmt.Sprintf("\n\t%s.IsHidden = %t", stateshapeIdent, stateshape.IsHidden))
			values.WriteString(fmt.Sprintf("\n\t%s.IsPresentationOverridden = %t", stateshapeIdent, stateshape.IsPresentationOverridden))
			values.WriteString(fmt.Sprintf("\n\t%s.Color = %s", stateshapeIdent, __gong__toRawStringLiteral(stateshape.Color)))
			values.WriteString(fmt.Sprintf("\n\t%s.FillOpacity = %f", stateshapeIdent, stateshape.FillOpacity))
			values.WriteString(fmt.Sprintf("\n\t%s.Stroke = %s", stateshapeIdent, __gong__toRawStringLiteral(stateshape.Stroke)))
			values.WriteString(fmt.Sprintf("\n\t%s.StrokeWidth = %f", stateshapeIdent, stateshape.StrokeWidth))
			values.WriteString(fmt.Sprintf("\n\t%s.StrokeOpacity = %f", stateshapeIdent, stateshape.StrokeOpacity))
			values.WriteString(fmt.Sprintf("\n\t%s.StrokeDashArray = %s", stateshapeIdent, __gong__toRawStringLiteral(stateshape.StrokeDashArray)))
			values.WriteString(fmt.Sprintf("\n\t%s.RX = %f", stateshapeIdent, stateshape.RX))
			values.WriteString(fmt.Sprintf("\n\t%s.FontSize = %s", stateshapeIdent, __gong__toRawStringLiteral(stateshape.FontSize)))
			values.WriteString(fmt.Sprintf("\n\t%s.FontWeight = %s", stateshapeIdent, __gong__toRawStringLiteral(stateshape.FontWeight)))
			values.WriteString(fmt.Sprintf("\n\t%s.FontColor = %s", stateshapeIdent, __gong__toRawStringLiteral(stateshape.FontColor)))
			if stateshape.State != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := stateshape.State.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.State = %s", stateshapeIdent, targetIdent))
			}
		}
	}
	if stage != nil {
		for _, transition := range __gong__sortStageSetInstances(stage.Transitions, stage.Transition_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			transitionIdent := transition.GongGetIdentifier(stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&models.Transition{Name: %s}).Stage(stage)", transitionIdent, __gong__toRawStringLiteral(transition.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", transitionIdent, __gong__toRawStringLiteral(transition.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.IsExpanded = %t", transitionIdent, transition.IsExpanded))
			values.WriteString(fmt.Sprintf("\n\t%s.IsRolesNodeExpanded = %t", transitionIdent, transition.IsRolesNodeExpanded))
			values.WriteString(fmt.Sprintf("\n\t%s.IsMessagesNodeExpanded = %t", transitionIdent, transition.IsMessagesNodeExpanded))
			if transition.Start != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := transition.Start.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Start = %s", transitionIdent, targetIdent))
			}
			if transition.End != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := transition.End.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.End = %s", transitionIdent, targetIdent))
			}
			for _, elem := range transition.RolesWithPermissions {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := elem.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.RolesWithPermissions = append(%s.RolesWithPermissions, %s)", transitionIdent, transitionIdent, targetIdent))
			}
			for _, elem := range transition.GeneratedMessages {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := elem.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.GeneratedMessages = append(%s.GeneratedMessages, %s)", transitionIdent, transitionIdent, targetIdent))
			}
			if transition.Guard != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := transition.Guard.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Guard = %s", transitionIdent, targetIdent))
			}
			for _, elem := range transition.Diagrams {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := elem.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Diagrams = append(%s.Diagrams, %s)", transitionIdent, transitionIdent, targetIdent))
			}
		}
	}
	if stage != nil {
		for _, transition_shape := range __gong__sortStageSetInstances(stage.Transition_Shapes, stage.Transition_Shape_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			transition_shapeIdent := transition_shape.GongGetIdentifier(stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&models.Transition_Shape{Name: %s}).Stage(stage)", transition_shapeIdent, __gong__toRawStringLiteral(transition_shape.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", transition_shapeIdent, __gong__toRawStringLiteral(transition_shape.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.StartRatio = %f", transition_shapeIdent, transition_shape.StartRatio))
			values.WriteString(fmt.Sprintf("\n\t%s.EndRatio = %f", transition_shapeIdent, transition_shape.EndRatio))
			if transition_shape.StartOrientation.ToCodeString() != "" {
				values.WriteString(fmt.Sprintf("\n\t%s.StartOrientation = models.%s", transition_shapeIdent, transition_shape.StartOrientation.ToCodeString()))
			} else {
				values.WriteString(fmt.Sprintf("\n\t%s.StartOrientation = \"\"", transition_shapeIdent))
			}
			if transition_shape.EndOrientation.ToCodeString() != "" {
				values.WriteString(fmt.Sprintf("\n\t%s.EndOrientation = models.%s", transition_shapeIdent, transition_shape.EndOrientation.ToCodeString()))
			} else {
				values.WriteString(fmt.Sprintf("\n\t%s.EndOrientation = \"\"", transition_shapeIdent))
			}
			values.WriteString(fmt.Sprintf("\n\t%s.CornerOffsetRatio = %f", transition_shapeIdent, transition_shape.CornerOffsetRatio))
			values.WriteString(fmt.Sprintf("\n\t%s.IsHidden = %t", transition_shapeIdent, transition_shape.IsHidden))
			values.WriteString(fmt.Sprintf("\n\t%s.IsPresentationOverridden = %t", transition_shapeIdent, transition_shape.IsPresentationOverridden))
			values.WriteString(fmt.Sprintf("\n\t%s.Stroke = %s", transition_shapeIdent, __gong__toRawStringLiteral(transition_shape.Stroke)))
			values.WriteString(fmt.Sprintf("\n\t%s.StrokeWidth = %f", transition_shapeIdent, transition_shape.StrokeWidth))
			values.WriteString(fmt.Sprintf("\n\t%s.StrokeOpacity = %f", transition_shapeIdent, transition_shape.StrokeOpacity))
			values.WriteString(fmt.Sprintf("\n\t%s.StrokeDashArray = %s", transition_shapeIdent, __gong__toRawStringLiteral(transition_shape.StrokeDashArray)))
			values.WriteString(fmt.Sprintf("\n\t%s.FontSize = %s", transition_shapeIdent, __gong__toRawStringLiteral(transition_shape.FontSize)))
			values.WriteString(fmt.Sprintf("\n\t%s.FontColor = %s", transition_shapeIdent, __gong__toRawStringLiteral(transition_shape.FontColor)))
			if transition_shape.Transition != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := transition_shape.Transition.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Transition = %s", transition_shapeIdent, targetIdent))
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
go run github.com/fullstack-lang/gong/dsm/statemachines/go/models/probe/cmd/probe@%s stage.go
*/
package %s

import (
	"slices"
	"time"

	"github.com/fullstack-lang/gong/dsm/statemachines/go/models"%s
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

	if modelsPackageName != "" && modelsPackageName != "github.com/fullstack-lang/gong/dsm/statemachines/go/models" {
		res = strings.ReplaceAll(res, "\"github.com/fullstack-lang/gong/dsm/statemachines/go/models\"", fmt.Sprintf("%q", modelsPackageName))
	}

	return res, nil
}

// insertion point for marshall field methods
func (action *Action) GongMarshallField(stage *Stage, fieldName string) (res string) {
	ident := action.GongGetIdentifier(stage)
	_ = ident
	switch fieldName {
	case "Name":
		res = __gong__marshallString(ident, "Name", action.Name)
	case "Criticality":
		res = __gong__marshallEnumString(ident, "Criticality", action.Criticality.ToCodeString())

	default:
		log.Panicf("Unknown field %s for Gongstruct Action", fieldName)
	}
	return
}

func (activities *Activities) GongMarshallField(stage *Stage, fieldName string) (res string) {
	ident := activities.GongGetIdentifier(stage)
	_ = ident
	switch fieldName {
	case "Name":
		res = __gong__marshallString(ident, "Name", activities.Name)
	case "Criticality":
		res = __gong__marshallEnumString(ident, "Criticality", activities.Criticality.ToCodeString())

	default:
		log.Panicf("Unknown field %s for Gongstruct Activities", fieldName)
	}
	return
}

func (diagram *Diagram) GongMarshallField(stage *Stage, fieldName string) (res string) {
	ident := diagram.GongGetIdentifier(stage)
	_ = ident
	switch fieldName {
	case "Name":
		res = __gong__marshallString(ident, "Name", diagram.Name)
	case "IsChecked":
		res = __gong__marshallBool(ident, "IsChecked", diagram.IsChecked)
	case "IsExpanded":
		res = __gong__marshallBool(ident, "IsExpanded", diagram.IsExpanded)
	case "IsEditable_":
		res = __gong__marshallBool(ident, "IsEditable_", diagram.IsEditable_)
	case "IsStatesNodeExpanded":
		res = __gong__marshallBool(ident, "IsStatesNodeExpanded", diagram.IsStatesNodeExpanded)
	case "ShowRoles":
		res = __gong__marshallBool(ident, "ShowRoles", diagram.ShowRoles)
	case "ShowMessages":
		res = __gong__marshallBool(ident, "ShowMessages", diagram.ShowMessages)

	case "State_Shapes":
		var sb strings.Builder
		for _, _stateshape := range diagram.State_Shapes {
			sb.WriteString(__gong__marshallSliceOfPointers(ident, "State_Shapes", _stateshape.GongGetIdentifier(stage)))
		}
		res = sb.String()
	case "StatesWhoseNodeIsExpanded":
		var sb strings.Builder
		for _, _state := range diagram.StatesWhoseNodeIsExpanded {
			sb.WriteString(__gong__marshallSliceOfPointers(ident, "StatesWhoseNodeIsExpanded", _state.GongGetIdentifier(stage)))
		}
		res = sb.String()
	case "Transition_Shapes":
		var sb strings.Builder
		for _, _transition_shape := range diagram.Transition_Shapes {
			sb.WriteString(__gong__marshallSliceOfPointers(ident, "Transition_Shapes", _transition_shape.GongGetIdentifier(stage)))
		}
		res = sb.String()
	case "Note_Shapes":
		var sb strings.Builder
		for _, _noteshape := range diagram.Note_Shapes {
			sb.WriteString(__gong__marshallSliceOfPointers(ident, "Note_Shapes", _noteshape.GongGetIdentifier(stage)))
		}
		res = sb.String()
	case "NoteState_Shapes":
		var sb strings.Builder
		for _, _notestateshape := range diagram.NoteState_Shapes {
			sb.WriteString(__gong__marshallSliceOfPointers(ident, "NoteState_Shapes", _notestateshape.GongGetIdentifier(stage)))
		}
		res = sb.String()
	default:
		log.Panicf("Unknown field %s for Gongstruct Diagram", fieldName)
	}
	return
}

func (guard *Guard) GongMarshallField(stage *Stage, fieldName string) (res string) {
	ident := guard.GongGetIdentifier(stage)
	_ = ident
	switch fieldName {
	case "Name":
		res = __gong__marshallString(ident, "Name", guard.Name)

	default:
		log.Panicf("Unknown field %s for Gongstruct Guard", fieldName)
	}
	return
}

func (kill *Kill) GongMarshallField(stage *Stage, fieldName string) (res string) {
	ident := kill.GongGetIdentifier(stage)
	_ = ident
	switch fieldName {
	case "Name":
		res = __gong__marshallString(ident, "Name", kill.Name)

	default:
		log.Panicf("Unknown field %s for Gongstruct Kill", fieldName)
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
	case "IsStateMachinesNodeExpanded":
		res = __gong__marshallBool(ident, "IsStateMachinesNodeExpanded", library.IsStateMachinesNodeExpanded)
	case "IsSubLibrariesNodeExpanded":
		res = __gong__marshallBool(ident, "IsSubLibrariesNodeExpanded", library.IsSubLibrariesNodeExpanded)
	case "IsExpandedTmp":
		res = __gong__marshallBool(ident, "IsExpandedTmp", library.IsExpandedTmp)
	case "IsRolesNodeExpanded":
		res = __gong__marshallBool(ident, "IsRolesNodeExpanded", library.IsRolesNodeExpanded)
	case "IsMessageTypesNodeExpanded":
		res = __gong__marshallBool(ident, "IsMessageTypesNodeExpanded", library.IsMessageTypesNodeExpanded)

	case "SubLibraries":
		var sb strings.Builder
		for _, _library := range library.SubLibraries {
			sb.WriteString(__gong__marshallSliceOfPointers(ident, "SubLibraries", _library.GongGetIdentifier(stage)))
		}
		res = sb.String()
	case "Diagrams":
		var sb strings.Builder
		for _, _diagram := range library.Diagrams {
			sb.WriteString(__gong__marshallSliceOfPointers(ident, "Diagrams", _diagram.GongGetIdentifier(stage)))
		}
		res = sb.String()
	case "RootStateMachines":
		var sb strings.Builder
		for _, _statemachine := range library.RootStateMachines {
			sb.WriteString(__gong__marshallSliceOfPointers(ident, "RootStateMachines", _statemachine.GongGetIdentifier(stage)))
		}
		res = sb.String()
	case "StateMachinesWhoseNodeIsExpanded":
		var sb strings.Builder
		for _, _statemachine := range library.StateMachinesWhoseNodeIsExpanded {
			sb.WriteString(__gong__marshallSliceOfPointers(ident, "StateMachinesWhoseNodeIsExpanded", _statemachine.GongGetIdentifier(stage)))
		}
		res = sb.String()
	case "SubLibrariesWhoseNodeIsExpanded":
		var sb strings.Builder
		for _, _library := range library.SubLibrariesWhoseNodeIsExpanded {
			sb.WriteString(__gong__marshallSliceOfPointers(ident, "SubLibrariesWhoseNodeIsExpanded", _library.GongGetIdentifier(stage)))
		}
		res = sb.String()
	case "Roles":
		var sb strings.Builder
		for _, _role := range library.Roles {
			sb.WriteString(__gong__marshallSliceOfPointers(ident, "Roles", _role.GongGetIdentifier(stage)))
		}
		res = sb.String()
	case "MessageTypes":
		var sb strings.Builder
		for _, _messagetype := range library.MessageTypes {
			sb.WriteString(__gong__marshallSliceOfPointers(ident, "MessageTypes", _messagetype.GongGetIdentifier(stage)))
		}
		res = sb.String()
	default:
		log.Panicf("Unknown field %s for Gongstruct Library", fieldName)
	}
	return
}

func (message *Message) GongMarshallField(stage *Stage, fieldName string) (res string) {
	ident := message.GongGetIdentifier(stage)
	_ = ident
	switch fieldName {
	case "Name":
		res = __gong__marshallString(ident, "Name", message.Name)
	case "IsSelected":
		res = __gong__marshallBool(ident, "IsSelected", message.IsSelected)

	case "MessageType":
		if message.MessageType != nil {
			res = __gong__marshallPointer(ident, "MessageType", message.MessageType.GongGetIdentifier(stage))
		} else {
			res = __gong__marshallPointer(ident, "MessageType", "nil")
		}
	case "OriginTransition":
		if message.OriginTransition != nil {
			res = __gong__marshallPointer(ident, "OriginTransition", message.OriginTransition.GongGetIdentifier(stage))
		} else {
			res = __gong__marshallPointer(ident, "OriginTransition", "nil")
		}
	default:
		log.Panicf("Unknown field %s for Gongstruct Message", fieldName)
	}
	return
}

func (messagetype *MessageType) GongMarshallField(stage *Stage, fieldName string) (res string) {
	ident := messagetype.GongGetIdentifier(stage)
	_ = ident
	switch fieldName {
	case "Name":
		res = __gong__marshallString(ident, "Name", messagetype.Name)
	case "Description":
		res = __gong__marshallString(ident, "Description", messagetype.Description)

	default:
		log.Panicf("Unknown field %s for Gongstruct MessageType", fieldName)
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

	case "State":
		if note.State != nil {
			res = __gong__marshallPointer(ident, "State", note.State.GongGetIdentifier(stage))
		} else {
			res = __gong__marshallPointer(ident, "State", "nil")
		}
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

func (notestateshape *NoteStateShape) GongMarshallField(stage *Stage, fieldName string) (res string) {
	ident := notestateshape.GongGetIdentifier(stage)
	_ = ident
	switch fieldName {
	case "Name":
		res = __gong__marshallString(ident, "Name", notestateshape.Name)
	case "StartRatio":
		res = __gong__marshallFloat(ident, "StartRatio", notestateshape.StartRatio)
	case "EndRatio":
		res = __gong__marshallFloat(ident, "EndRatio", notestateshape.EndRatio)
	case "StartOrientation":
		res = __gong__marshallEnumString(ident, "StartOrientation", notestateshape.StartOrientation.ToCodeString())
	case "EndOrientation":
		res = __gong__marshallEnumString(ident, "EndOrientation", notestateshape.EndOrientation.ToCodeString())
	case "CornerOffsetRatio":
		res = __gong__marshallFloat(ident, "CornerOffsetRatio", notestateshape.CornerOffsetRatio)
	case "IsHidden":
		res = __gong__marshallBool(ident, "IsHidden", notestateshape.IsHidden)
	case "IsPresentationOverridden":
		res = __gong__marshallBool(ident, "IsPresentationOverridden", notestateshape.IsPresentationOverridden)
	case "Stroke":
		res = __gong__marshallString(ident, "Stroke", notestateshape.Stroke)
	case "StrokeWidth":
		res = __gong__marshallFloat(ident, "StrokeWidth", notestateshape.StrokeWidth)
	case "StrokeOpacity":
		res = __gong__marshallFloat(ident, "StrokeOpacity", notestateshape.StrokeOpacity)
	case "StrokeDashArray":
		res = __gong__marshallString(ident, "StrokeDashArray", notestateshape.StrokeDashArray)
	case "FontSize":
		res = __gong__marshallString(ident, "FontSize", notestateshape.FontSize)
	case "FontColor":
		res = __gong__marshallString(ident, "FontColor", notestateshape.FontColor)

	case "Note":
		if notestateshape.Note != nil {
			res = __gong__marshallPointer(ident, "Note", notestateshape.Note.GongGetIdentifier(stage))
		} else {
			res = __gong__marshallPointer(ident, "Note", "nil")
		}
	case "State":
		if notestateshape.State != nil {
			res = __gong__marshallPointer(ident, "State", notestateshape.State.GongGetIdentifier(stage))
		} else {
			res = __gong__marshallPointer(ident, "State", "nil")
		}
	default:
		log.Panicf("Unknown field %s for Gongstruct NoteStateShape", fieldName)
	}
	return
}

func (object *Object) GongMarshallField(stage *Stage, fieldName string) (res string) {
	ident := object.GongGetIdentifier(stage)
	_ = ident
	switch fieldName {
	case "Name":
		res = __gong__marshallString(ident, "Name", object.Name)
	case "IsSelected":
		res = __gong__marshallBool(ident, "IsSelected", object.IsSelected)
	case "Rank":
		res = __gong__marshallInt(ident, "Rank", object.Rank)
	case "DOF":
		res = __gong__marshallTime(ident, "DOF", object.DOF.String())

	case "State":
		if object.State != nil {
			res = __gong__marshallPointer(ident, "State", object.State.GongGetIdentifier(stage))
		} else {
			res = __gong__marshallPointer(ident, "State", "nil")
		}
	case "Messages":
		var sb strings.Builder
		for _, _message := range object.Messages {
			sb.WriteString(__gong__marshallSliceOfPointers(ident, "Messages", _message.GongGetIdentifier(stage)))
		}
		res = sb.String()
	default:
		log.Panicf("Unknown field %s for Gongstruct Object", fieldName)
	}
	return
}

func (role *Role) GongMarshallField(stage *Stage, fieldName string) (res string) {
	ident := role.GongGetIdentifier(stage)
	_ = ident
	switch fieldName {
	case "Name":
		res = __gong__marshallString(ident, "Name", role.Name)
	case "Acronym":
		res = __gong__marshallString(ident, "Acronym", role.Acronym)

	case "RolesWithSamePermissions":
		var sb strings.Builder
		for _, _role := range role.RolesWithSamePermissions {
			sb.WriteString(__gong__marshallSliceOfPointers(ident, "RolesWithSamePermissions", _role.GongGetIdentifier(stage)))
		}
		res = sb.String()
	default:
		log.Panicf("Unknown field %s for Gongstruct Role", fieldName)
	}
	return
}

func (state *State) GongMarshallField(stage *Stage, fieldName string) (res string) {
	ident := state.GongGetIdentifier(stage)
	_ = ident
	switch fieldName {
	case "Name":
		res = __gong__marshallString(ident, "Name", state.Name)
	case "IsEndState":
		res = __gong__marshallBool(ident, "IsEndState", state.IsEndState)
	case "IsDecisionNode":
		res = __gong__marshallBool(ident, "IsDecisionNode", state.IsDecisionNode)
	case "IsFictious":
		res = __gong__marshallBool(ident, "IsFictious", state.IsFictious)

	case "SubStates":
		var sb strings.Builder
		for _, _state := range state.SubStates {
			sb.WriteString(__gong__marshallSliceOfPointers(ident, "SubStates", _state.GongGetIdentifier(stage)))
		}
		res = sb.String()
	case "Entry":
		if state.Entry != nil {
			res = __gong__marshallPointer(ident, "Entry", state.Entry.GongGetIdentifier(stage))
		} else {
			res = __gong__marshallPointer(ident, "Entry", "nil")
		}
	case "Activities":
		var sb strings.Builder
		for _, _activities := range state.Activities {
			sb.WriteString(__gong__marshallSliceOfPointers(ident, "Activities", _activities.GongGetIdentifier(stage)))
		}
		res = sb.String()
	case "Exit":
		if state.Exit != nil {
			res = __gong__marshallPointer(ident, "Exit", state.Exit.GongGetIdentifier(stage))
		} else {
			res = __gong__marshallPointer(ident, "Exit", "nil")
		}
	case "Parent":
		if state.Parent != nil {
			res = __gong__marshallPointer(ident, "Parent", state.Parent.GongGetIdentifier(stage))
		} else {
			res = __gong__marshallPointer(ident, "Parent", "nil")
		}
	case "Diagrams":
		var sb strings.Builder
		for _, _diagram := range state.Diagrams {
			sb.WriteString(__gong__marshallSliceOfPointers(ident, "Diagrams", _diagram.GongGetIdentifier(stage)))
		}
		res = sb.String()
	case "Notes":
		var sb strings.Builder
		for _, _note := range state.Notes {
			sb.WriteString(__gong__marshallSliceOfPointers(ident, "Notes", _note.GongGetIdentifier(stage)))
		}
		res = sb.String()
	default:
		log.Panicf("Unknown field %s for Gongstruct State", fieldName)
	}
	return
}

func (statemachine *StateMachine) GongMarshallField(stage *Stage, fieldName string) (res string) {
	ident := statemachine.GongGetIdentifier(stage)
	_ = ident
	switch fieldName {
	case "Name":
		res = __gong__marshallString(ident, "Name", statemachine.Name)
	case "IsWithTransitionNameAutonamticalyGenerated":
		res = __gong__marshallBool(ident, "IsWithTransitionNameAutonamticalyGenerated", statemachine.IsWithTransitionNameAutonamticalyGenerated)
	case "ComputedPrefix":
		res = __gong__marshallString(ident, "ComputedPrefix", statemachine.ComputedPrefix)
	case "IsExpanded":
		res = __gong__marshallBool(ident, "IsExpanded", statemachine.IsExpanded)

	case "InitialState":
		if statemachine.InitialState != nil {
			res = __gong__marshallPointer(ident, "InitialState", statemachine.InitialState.GongGetIdentifier(stage))
		} else {
			res = __gong__marshallPointer(ident, "InitialState", "nil")
		}
	case "States":
		var sb strings.Builder
		for _, _state := range statemachine.States {
			sb.WriteString(__gong__marshallSliceOfPointers(ident, "States", _state.GongGetIdentifier(stage)))
		}
		res = sb.String()
	case "Diagrams":
		var sb strings.Builder
		for _, _diagram := range statemachine.Diagrams {
			sb.WriteString(__gong__marshallSliceOfPointers(ident, "Diagrams", _diagram.GongGetIdentifier(stage)))
		}
		res = sb.String()
	default:
		log.Panicf("Unknown field %s for Gongstruct StateMachine", fieldName)
	}
	return
}

func (stateshape *StateShape) GongMarshallField(stage *Stage, fieldName string) (res string) {
	ident := stateshape.GongGetIdentifier(stage)
	_ = ident
	switch fieldName {
	case "Name":
		res = __gong__marshallString(ident, "Name", stateshape.Name)
	case "X":
		res = __gong__marshallFloat(ident, "X", stateshape.X)
	case "Y":
		res = __gong__marshallFloat(ident, "Y", stateshape.Y)
	case "Width":
		res = __gong__marshallFloat(ident, "Width", stateshape.Width)
	case "Height":
		res = __gong__marshallFloat(ident, "Height", stateshape.Height)
	case "IsHidden":
		res = __gong__marshallBool(ident, "IsHidden", stateshape.IsHidden)
	case "IsPresentationOverridden":
		res = __gong__marshallBool(ident, "IsPresentationOverridden", stateshape.IsPresentationOverridden)
	case "Color":
		res = __gong__marshallString(ident, "Color", stateshape.Color)
	case "FillOpacity":
		res = __gong__marshallFloat(ident, "FillOpacity", stateshape.FillOpacity)
	case "Stroke":
		res = __gong__marshallString(ident, "Stroke", stateshape.Stroke)
	case "StrokeWidth":
		res = __gong__marshallFloat(ident, "StrokeWidth", stateshape.StrokeWidth)
	case "StrokeOpacity":
		res = __gong__marshallFloat(ident, "StrokeOpacity", stateshape.StrokeOpacity)
	case "StrokeDashArray":
		res = __gong__marshallString(ident, "StrokeDashArray", stateshape.StrokeDashArray)
	case "RX":
		res = __gong__marshallFloat(ident, "RX", stateshape.RX)
	case "FontSize":
		res = __gong__marshallString(ident, "FontSize", stateshape.FontSize)
	case "FontWeight":
		res = __gong__marshallString(ident, "FontWeight", stateshape.FontWeight)
	case "FontColor":
		res = __gong__marshallString(ident, "FontColor", stateshape.FontColor)

	case "State":
		if stateshape.State != nil {
			res = __gong__marshallPointer(ident, "State", stateshape.State.GongGetIdentifier(stage))
		} else {
			res = __gong__marshallPointer(ident, "State", "nil")
		}
	default:
		log.Panicf("Unknown field %s for Gongstruct StateShape", fieldName)
	}
	return
}

func (transition *Transition) GongMarshallField(stage *Stage, fieldName string) (res string) {
	ident := transition.GongGetIdentifier(stage)
	_ = ident
	switch fieldName {
	case "Name":
		res = __gong__marshallString(ident, "Name", transition.Name)
	case "IsExpanded":
		res = __gong__marshallBool(ident, "IsExpanded", transition.IsExpanded)
	case "IsRolesNodeExpanded":
		res = __gong__marshallBool(ident, "IsRolesNodeExpanded", transition.IsRolesNodeExpanded)
	case "IsMessagesNodeExpanded":
		res = __gong__marshallBool(ident, "IsMessagesNodeExpanded", transition.IsMessagesNodeExpanded)

	case "Start":
		if transition.Start != nil {
			res = __gong__marshallPointer(ident, "Start", transition.Start.GongGetIdentifier(stage))
		} else {
			res = __gong__marshallPointer(ident, "Start", "nil")
		}
	case "End":
		if transition.End != nil {
			res = __gong__marshallPointer(ident, "End", transition.End.GongGetIdentifier(stage))
		} else {
			res = __gong__marshallPointer(ident, "End", "nil")
		}
	case "RolesWithPermissions":
		var sb strings.Builder
		for _, _role := range transition.RolesWithPermissions {
			sb.WriteString(__gong__marshallSliceOfPointers(ident, "RolesWithPermissions", _role.GongGetIdentifier(stage)))
		}
		res = sb.String()
	case "GeneratedMessages":
		var sb strings.Builder
		for _, _messagetype := range transition.GeneratedMessages {
			sb.WriteString(__gong__marshallSliceOfPointers(ident, "GeneratedMessages", _messagetype.GongGetIdentifier(stage)))
		}
		res = sb.String()
	case "Guard":
		if transition.Guard != nil {
			res = __gong__marshallPointer(ident, "Guard", transition.Guard.GongGetIdentifier(stage))
		} else {
			res = __gong__marshallPointer(ident, "Guard", "nil")
		}
	case "Diagrams":
		var sb strings.Builder
		for _, _diagram := range transition.Diagrams {
			sb.WriteString(__gong__marshallSliceOfPointers(ident, "Diagrams", _diagram.GongGetIdentifier(stage)))
		}
		res = sb.String()
	default:
		log.Panicf("Unknown field %s for Gongstruct Transition", fieldName)
	}
	return
}

func (transition_shape *Transition_Shape) GongMarshallField(stage *Stage, fieldName string) (res string) {
	ident := transition_shape.GongGetIdentifier(stage)
	_ = ident
	switch fieldName {
	case "Name":
		res = __gong__marshallString(ident, "Name", transition_shape.Name)
	case "StartRatio":
		res = __gong__marshallFloat(ident, "StartRatio", transition_shape.StartRatio)
	case "EndRatio":
		res = __gong__marshallFloat(ident, "EndRatio", transition_shape.EndRatio)
	case "StartOrientation":
		res = __gong__marshallEnumString(ident, "StartOrientation", transition_shape.StartOrientation.ToCodeString())
	case "EndOrientation":
		res = __gong__marshallEnumString(ident, "EndOrientation", transition_shape.EndOrientation.ToCodeString())
	case "CornerOffsetRatio":
		res = __gong__marshallFloat(ident, "CornerOffsetRatio", transition_shape.CornerOffsetRatio)
	case "IsHidden":
		res = __gong__marshallBool(ident, "IsHidden", transition_shape.IsHidden)
	case "IsPresentationOverridden":
		res = __gong__marshallBool(ident, "IsPresentationOverridden", transition_shape.IsPresentationOverridden)
	case "Stroke":
		res = __gong__marshallString(ident, "Stroke", transition_shape.Stroke)
	case "StrokeWidth":
		res = __gong__marshallFloat(ident, "StrokeWidth", transition_shape.StrokeWidth)
	case "StrokeOpacity":
		res = __gong__marshallFloat(ident, "StrokeOpacity", transition_shape.StrokeOpacity)
	case "StrokeDashArray":
		res = __gong__marshallString(ident, "StrokeDashArray", transition_shape.StrokeDashArray)
	case "FontSize":
		res = __gong__marshallString(ident, "FontSize", transition_shape.FontSize)
	case "FontColor":
		res = __gong__marshallString(ident, "FontColor", transition_shape.FontColor)

	case "Transition":
		if transition_shape.Transition != nil {
			res = __gong__marshallPointer(ident, "Transition", transition_shape.Transition.GongGetIdentifier(stage))
		} else {
			res = __gong__marshallPointer(ident, "Transition", "nil")
		}
	default:
		log.Panicf("Unknown field %s for Gongstruct Transition_Shape", fieldName)
	}
	return
}

// insertion point for marshall all fields methods
func (action *Action) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(action.GongMarshallField(stage, "Name"))
		initializerStatements.WriteString(action.GongMarshallField(stage, "Criticality"))
	}
	initRes = initializerStatements.String()
	ptrRes = pointersInitializesStatements.String()
	return
}
func (activities *Activities) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(activities.GongMarshallField(stage, "Name"))
		initializerStatements.WriteString(activities.GongMarshallField(stage, "Criticality"))
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
		initializerStatements.WriteString(diagram.GongMarshallField(stage, "IsChecked"))
		initializerStatements.WriteString(diagram.GongMarshallField(stage, "IsExpanded"))
		initializerStatements.WriteString(diagram.GongMarshallField(stage, "IsEditable_"))
		initializerStatements.WriteString(diagram.GongMarshallField(stage, "IsStatesNodeExpanded"))
		pointersInitializesStatements.WriteString(diagram.GongMarshallField(stage, "State_Shapes"))
		pointersInitializesStatements.WriteString(diagram.GongMarshallField(stage, "StatesWhoseNodeIsExpanded"))
		pointersInitializesStatements.WriteString(diagram.GongMarshallField(stage, "Transition_Shapes"))
		pointersInitializesStatements.WriteString(diagram.GongMarshallField(stage, "Note_Shapes"))
		pointersInitializesStatements.WriteString(diagram.GongMarshallField(stage, "NoteState_Shapes"))
		initializerStatements.WriteString(diagram.GongMarshallField(stage, "ShowRoles"))
		initializerStatements.WriteString(diagram.GongMarshallField(stage, "ShowMessages"))
	}
	initRes = initializerStatements.String()
	ptrRes = pointersInitializesStatements.String()
	return
}
func (guard *Guard) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(guard.GongMarshallField(stage, "Name"))
	}
	initRes = initializerStatements.String()
	ptrRes = pointersInitializesStatements.String()
	return
}
func (kill *Kill) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(kill.GongMarshallField(stage, "Name"))
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
		pointersInitializesStatements.WriteString(library.GongMarshallField(stage, "Diagrams"))
		pointersInitializesStatements.WriteString(library.GongMarshallField(stage, "RootStateMachines"))
		initializerStatements.WriteString(library.GongMarshallField(stage, "IsStateMachinesNodeExpanded"))
		pointersInitializesStatements.WriteString(library.GongMarshallField(stage, "StateMachinesWhoseNodeIsExpanded"))
		initializerStatements.WriteString(library.GongMarshallField(stage, "IsSubLibrariesNodeExpanded"))
		pointersInitializesStatements.WriteString(library.GongMarshallField(stage, "SubLibrariesWhoseNodeIsExpanded"))
		initializerStatements.WriteString(library.GongMarshallField(stage, "IsExpandedTmp"))
		pointersInitializesStatements.WriteString(library.GongMarshallField(stage, "Roles"))
		initializerStatements.WriteString(library.GongMarshallField(stage, "IsRolesNodeExpanded"))
		pointersInitializesStatements.WriteString(library.GongMarshallField(stage, "MessageTypes"))
		initializerStatements.WriteString(library.GongMarshallField(stage, "IsMessageTypesNodeExpanded"))
	}
	initRes = initializerStatements.String()
	ptrRes = pointersInitializesStatements.String()
	return
}
func (message *Message) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(message.GongMarshallField(stage, "Name"))
		initializerStatements.WriteString(message.GongMarshallField(stage, "IsSelected"))
		pointersInitializesStatements.WriteString(message.GongMarshallField(stage, "MessageType"))
		pointersInitializesStatements.WriteString(message.GongMarshallField(stage, "OriginTransition"))
	}
	initRes = initializerStatements.String()
	ptrRes = pointersInitializesStatements.String()
	return
}
func (messagetype *MessageType) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(messagetype.GongMarshallField(stage, "Name"))
		initializerStatements.WriteString(messagetype.GongMarshallField(stage, "Description"))
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
		pointersInitializesStatements.WriteString(note.GongMarshallField(stage, "State"))
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
func (notestateshape *NoteStateShape) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(notestateshape.GongMarshallField(stage, "Name"))
		pointersInitializesStatements.WriteString(notestateshape.GongMarshallField(stage, "Note"))
		pointersInitializesStatements.WriteString(notestateshape.GongMarshallField(stage, "State"))
		initializerStatements.WriteString(notestateshape.GongMarshallField(stage, "StartRatio"))
		initializerStatements.WriteString(notestateshape.GongMarshallField(stage, "EndRatio"))
		initializerStatements.WriteString(notestateshape.GongMarshallField(stage, "StartOrientation"))
		initializerStatements.WriteString(notestateshape.GongMarshallField(stage, "EndOrientation"))
		initializerStatements.WriteString(notestateshape.GongMarshallField(stage, "CornerOffsetRatio"))
		initializerStatements.WriteString(notestateshape.GongMarshallField(stage, "IsHidden"))
		initializerStatements.WriteString(notestateshape.GongMarshallField(stage, "IsPresentationOverridden"))
		initializerStatements.WriteString(notestateshape.GongMarshallField(stage, "Stroke"))
		initializerStatements.WriteString(notestateshape.GongMarshallField(stage, "StrokeWidth"))
		initializerStatements.WriteString(notestateshape.GongMarshallField(stage, "StrokeOpacity"))
		initializerStatements.WriteString(notestateshape.GongMarshallField(stage, "StrokeDashArray"))
		initializerStatements.WriteString(notestateshape.GongMarshallField(stage, "FontSize"))
		initializerStatements.WriteString(notestateshape.GongMarshallField(stage, "FontColor"))
	}
	initRes = initializerStatements.String()
	ptrRes = pointersInitializesStatements.String()
	return
}
func (object *Object) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(object.GongMarshallField(stage, "Name"))
		pointersInitializesStatements.WriteString(object.GongMarshallField(stage, "State"))
		initializerStatements.WriteString(object.GongMarshallField(stage, "IsSelected"))
		initializerStatements.WriteString(object.GongMarshallField(stage, "Rank"))
		initializerStatements.WriteString(object.GongMarshallField(stage, "DOF"))
		pointersInitializesStatements.WriteString(object.GongMarshallField(stage, "Messages"))
	}
	initRes = initializerStatements.String()
	ptrRes = pointersInitializesStatements.String()
	return
}
func (role *Role) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(role.GongMarshallField(stage, "Name"))
		initializerStatements.WriteString(role.GongMarshallField(stage, "Acronym"))
		pointersInitializesStatements.WriteString(role.GongMarshallField(stage, "RolesWithSamePermissions"))
	}
	initRes = initializerStatements.String()
	ptrRes = pointersInitializesStatements.String()
	return
}
func (state *State) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(state.GongMarshallField(stage, "Name"))
		initializerStatements.WriteString(state.GongMarshallField(stage, "IsEndState"))
		initializerStatements.WriteString(state.GongMarshallField(stage, "IsDecisionNode"))
		pointersInitializesStatements.WriteString(state.GongMarshallField(stage, "SubStates"))
		pointersInitializesStatements.WriteString(state.GongMarshallField(stage, "Entry"))
		pointersInitializesStatements.WriteString(state.GongMarshallField(stage, "Activities"))
		pointersInitializesStatements.WriteString(state.GongMarshallField(stage, "Exit"))
		pointersInitializesStatements.WriteString(state.GongMarshallField(stage, "Parent"))
		initializerStatements.WriteString(state.GongMarshallField(stage, "IsFictious"))
		pointersInitializesStatements.WriteString(state.GongMarshallField(stage, "Diagrams"))
		pointersInitializesStatements.WriteString(state.GongMarshallField(stage, "Notes"))
	}
	initRes = initializerStatements.String()
	ptrRes = pointersInitializesStatements.String()
	return
}
func (statemachine *StateMachine) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(statemachine.GongMarshallField(stage, "Name"))
		pointersInitializesStatements.WriteString(statemachine.GongMarshallField(stage, "InitialState"))
		pointersInitializesStatements.WriteString(statemachine.GongMarshallField(stage, "States"))
		pointersInitializesStatements.WriteString(statemachine.GongMarshallField(stage, "Diagrams"))
		initializerStatements.WriteString(statemachine.GongMarshallField(stage, "IsWithTransitionNameAutonamticalyGenerated"))
		initializerStatements.WriteString(statemachine.GongMarshallField(stage, "ComputedPrefix"))
		initializerStatements.WriteString(statemachine.GongMarshallField(stage, "IsExpanded"))
	}
	initRes = initializerStatements.String()
	ptrRes = pointersInitializesStatements.String()
	return
}
func (stateshape *StateShape) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(stateshape.GongMarshallField(stage, "Name"))
		pointersInitializesStatements.WriteString(stateshape.GongMarshallField(stage, "State"))
		initializerStatements.WriteString(stateshape.GongMarshallField(stage, "X"))
		initializerStatements.WriteString(stateshape.GongMarshallField(stage, "Y"))
		initializerStatements.WriteString(stateshape.GongMarshallField(stage, "Width"))
		initializerStatements.WriteString(stateshape.GongMarshallField(stage, "Height"))
		initializerStatements.WriteString(stateshape.GongMarshallField(stage, "IsHidden"))
		initializerStatements.WriteString(stateshape.GongMarshallField(stage, "IsPresentationOverridden"))
		initializerStatements.WriteString(stateshape.GongMarshallField(stage, "Color"))
		initializerStatements.WriteString(stateshape.GongMarshallField(stage, "FillOpacity"))
		initializerStatements.WriteString(stateshape.GongMarshallField(stage, "Stroke"))
		initializerStatements.WriteString(stateshape.GongMarshallField(stage, "StrokeWidth"))
		initializerStatements.WriteString(stateshape.GongMarshallField(stage, "StrokeOpacity"))
		initializerStatements.WriteString(stateshape.GongMarshallField(stage, "StrokeDashArray"))
		initializerStatements.WriteString(stateshape.GongMarshallField(stage, "RX"))
		initializerStatements.WriteString(stateshape.GongMarshallField(stage, "FontSize"))
		initializerStatements.WriteString(stateshape.GongMarshallField(stage, "FontWeight"))
		initializerStatements.WriteString(stateshape.GongMarshallField(stage, "FontColor"))
	}
	initRes = initializerStatements.String()
	ptrRes = pointersInitializesStatements.String()
	return
}
func (transition *Transition) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(transition.GongMarshallField(stage, "Name"))
		pointersInitializesStatements.WriteString(transition.GongMarshallField(stage, "Start"))
		pointersInitializesStatements.WriteString(transition.GongMarshallField(stage, "End"))
		pointersInitializesStatements.WriteString(transition.GongMarshallField(stage, "RolesWithPermissions"))
		pointersInitializesStatements.WriteString(transition.GongMarshallField(stage, "GeneratedMessages"))
		pointersInitializesStatements.WriteString(transition.GongMarshallField(stage, "Guard"))
		pointersInitializesStatements.WriteString(transition.GongMarshallField(stage, "Diagrams"))
		initializerStatements.WriteString(transition.GongMarshallField(stage, "IsExpanded"))
		initializerStatements.WriteString(transition.GongMarshallField(stage, "IsRolesNodeExpanded"))
		initializerStatements.WriteString(transition.GongMarshallField(stage, "IsMessagesNodeExpanded"))
	}
	initRes = initializerStatements.String()
	ptrRes = pointersInitializesStatements.String()
	return
}
func (transition_shape *Transition_Shape) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(transition_shape.GongMarshallField(stage, "Name"))
		pointersInitializesStatements.WriteString(transition_shape.GongMarshallField(stage, "Transition"))
		initializerStatements.WriteString(transition_shape.GongMarshallField(stage, "StartRatio"))
		initializerStatements.WriteString(transition_shape.GongMarshallField(stage, "EndRatio"))
		initializerStatements.WriteString(transition_shape.GongMarshallField(stage, "StartOrientation"))
		initializerStatements.WriteString(transition_shape.GongMarshallField(stage, "EndOrientation"))
		initializerStatements.WriteString(transition_shape.GongMarshallField(stage, "CornerOffsetRatio"))
		initializerStatements.WriteString(transition_shape.GongMarshallField(stage, "IsHidden"))
		initializerStatements.WriteString(transition_shape.GongMarshallField(stage, "IsPresentationOverridden"))
		initializerStatements.WriteString(transition_shape.GongMarshallField(stage, "Stroke"))
		initializerStatements.WriteString(transition_shape.GongMarshallField(stage, "StrokeWidth"))
		initializerStatements.WriteString(transition_shape.GongMarshallField(stage, "StrokeOpacity"))
		initializerStatements.WriteString(transition_shape.GongMarshallField(stage, "StrokeDashArray"))
		initializerStatements.WriteString(transition_shape.GongMarshallField(stage, "FontSize"))
		initializerStatements.WriteString(transition_shape.GongMarshallField(stage, "FontColor"))
	}
	initRes = initializerStatements.String()
	ptrRes = pointersInitializesStatements.String()
	return
}
