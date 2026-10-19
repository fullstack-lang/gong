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
	modelsPackageName := "github.com/fullstack-lang/gong/lib/sim/go/models"
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
		for _, command := range __gong__sortStageSetInstances(stage.Commands, stage.Command_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			commandIdent := command.GongGetIdentifier(stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&models.Command{Name: %s}).Stage(stage)", commandIdent, __gong__toRawStringLiteral(command.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", commandIdent, __gong__toRawStringLiteral(command.Name)))
			if command.Command.ToCodeString() != "" {
				values.WriteString(fmt.Sprintf("\n\t%s.Command = models.%s", commandIdent, command.Command.ToCodeString()))
			} else {
				values.WriteString(fmt.Sprintf("\n\t%s.Command = \"\"", commandIdent))
			}
			values.WriteString(fmt.Sprintf("\n\t%s.CommandDate = %s", commandIdent, __gong__toRawStringLiteral(command.CommandDate)))
			if command.Engine != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := command.Engine.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Engine = %s", commandIdent, targetIdent))
			}
		}
	}
	if stage != nil {
		for _, dummyagent := range __gong__sortStageSetInstances(stage.DummyAgents, stage.DummyAgent_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			dummyagentIdent := dummyagent.GongGetIdentifier(stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&models.DummyAgent{Name: %s}).Stage(stage)", dummyagentIdent, __gong__toRawStringLiteral(dummyagent.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.TechName = %s", dummyagentIdent, __gong__toRawStringLiteral(dummyagent.TechName)))
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", dummyagentIdent, __gong__toRawStringLiteral(dummyagent.Name)))
			if dummyagent.Engine != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := dummyagent.Engine.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Engine = %s", dummyagentIdent, targetIdent))
			}
		}
	}
	if stage != nil {
		for _, engine := range __gong__sortStageSetInstances(stage.Engines, stage.Engine_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			engineIdent := engine.GongGetIdentifier(stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&models.Engine{Name: %s}).Stage(stage)", engineIdent, __gong__toRawStringLiteral(engine.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", engineIdent, __gong__toRawStringLiteral(engine.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.EndTime = %s", engineIdent, __gong__toRawStringLiteral(engine.EndTime)))
			values.WriteString(fmt.Sprintf("\n\t%s.CurrentTime = %s", engineIdent, __gong__toRawStringLiteral(engine.CurrentTime)))
			values.WriteString(fmt.Sprintf("\n\t%s.DisplayFormat = %s", engineIdent, __gong__toRawStringLiteral(engine.DisplayFormat)))
			values.WriteString(fmt.Sprintf("\n\t%s.SecondsSinceStart = %f", engineIdent, engine.SecondsSinceStart))
			values.WriteString(fmt.Sprintf("\n\t%s.Fired = %d", engineIdent, engine.Fired))
			if engine.ControlMode.ToCodeString() != "" {
				values.WriteString(fmt.Sprintf("\n\t%s.ControlMode = models.%s", engineIdent, engine.ControlMode.ToCodeString()))
			} else {
				values.WriteString(fmt.Sprintf("\n\t%s.ControlMode = \"\"", engineIdent))
			}
			if engine.State.ToCodeString() != "" {
				values.WriteString(fmt.Sprintf("\n\t%s.State = models.%s", engineIdent, engine.State.ToCodeString()))
			} else {
				values.WriteString(fmt.Sprintf("\n\t%s.State = \"\"", engineIdent))
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Speed = %f", engineIdent, engine.Speed))
		}
	}
	if stage != nil {
		for _, event := range __gong__sortStageSetInstances(stage.Events, stage.Event_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			eventIdent := event.GongGetIdentifier(stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&models.Event{Name: %s}).Stage(stage)", eventIdent, __gong__toRawStringLiteral(event.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", eventIdent, __gong__toRawStringLiteral(event.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Duration = time.Duration(%d)", eventIdent, int64(event.Duration)))
		}
	}
	if stage != nil {
		for _, status := range __gong__sortStageSetInstances(stage.Statuss, stage.Status_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			statusIdent := status.GongGetIdentifier(stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&models.Status{Name: %s}).Stage(stage)", statusIdent, __gong__toRawStringLiteral(status.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", statusIdent, __gong__toRawStringLiteral(status.Name)))
			if status.CurrentCommand.ToCodeString() != "" {
				values.WriteString(fmt.Sprintf("\n\t%s.CurrentCommand = models.%s", statusIdent, status.CurrentCommand.ToCodeString()))
			} else {
				values.WriteString(fmt.Sprintf("\n\t%s.CurrentCommand = \"\"", statusIdent))
			}
			values.WriteString(fmt.Sprintf("\n\t%s.CompletionDate = %s", statusIdent, __gong__toRawStringLiteral(status.CompletionDate)))
			if status.CurrentSpeedCommand.ToCodeString() != "" {
				values.WriteString(fmt.Sprintf("\n\t%s.CurrentSpeedCommand = models.%s", statusIdent, status.CurrentSpeedCommand.ToCodeString()))
			} else {
				values.WriteString(fmt.Sprintf("\n\t%s.CurrentSpeedCommand = \"\"", statusIdent))
			}
			values.WriteString(fmt.Sprintf("\n\t%s.SpeedCommandCompletionDate = %s", statusIdent, __gong__toRawStringLiteral(status.SpeedCommandCompletionDate)))
		}
	}
	if stage != nil {
		for _, updatestate := range __gong__sortStageSetInstances(stage.UpdateStates, stage.UpdateState_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			updatestateIdent := updatestate.GongGetIdentifier(stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&models.UpdateState{Name: %s}).Stage(stage)", updatestateIdent, __gong__toRawStringLiteral(updatestate.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", updatestateIdent, __gong__toRawStringLiteral(updatestate.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Duration = time.Duration(%d)", updatestateIdent, int64(updatestate.Duration)))
			values.WriteString(fmt.Sprintf("\n\t%s.Period = time.Duration(%d)", updatestateIdent, int64(updatestate.Period)))
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
go run github.com/fullstack-lang/gong/lib/sim/go/models/probe/cmd/probe@%s stage.go
*/
package %s

import (
	"slices"
	"time"

	"github.com/fullstack-lang/gong/lib/sim/go/models"%s
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

	if modelsPackageName != "" && modelsPackageName != "github.com/fullstack-lang/gong/lib/sim/go/models" {
		res = strings.ReplaceAll(res, "\"github.com/fullstack-lang/gong/lib/sim/go/models\"", fmt.Sprintf("%q", modelsPackageName))
	}

	return res, nil
}

// insertion point for marshall field methods
func (command *Command) GongMarshallField(stage *Stage, fieldName string) (res string) {
	ident := command.GongGetIdentifier(stage)
	_ = ident
	switch fieldName {
	case "Name":
		res = __gong__marshallString(ident, "Name", command.Name)
	case "Command":
		res = __gong__marshallEnumString(ident, "Command", command.Command.ToCodeString())
	case "CommandDate":
		res = __gong__marshallString(ident, "CommandDate", command.CommandDate)

	case "Engine":
		if command.Engine != nil {
			res = __gong__marshallPointer(ident, "Engine", command.Engine.GongGetIdentifier(stage))
		} else {
			res = __gong__marshallPointer(ident, "Engine", "nil")
		}
	default:
		log.Panicf("Unknown field %s for Gongstruct Command", fieldName)
	}
	return
}

func (dummyagent *DummyAgent) GongMarshallField(stage *Stage, fieldName string) (res string) {
	ident := dummyagent.GongGetIdentifier(stage)
	_ = ident
	switch fieldName {
	case "TechName":
		res = __gong__marshallString(ident, "TechName", dummyagent.TechName)
	case "Name":
		res = __gong__marshallString(ident, "Name", dummyagent.Name)

	default:
		log.Panicf("Unknown field %s for Gongstruct DummyAgent", fieldName)
	}
	return
}

func (engine *Engine) GongMarshallField(stage *Stage, fieldName string) (res string) {
	ident := engine.GongGetIdentifier(stage)
	_ = ident
	switch fieldName {
	case "Name":
		res = __gong__marshallString(ident, "Name", engine.Name)
	case "EndTime":
		res = __gong__marshallString(ident, "EndTime", engine.EndTime)
	case "CurrentTime":
		res = __gong__marshallString(ident, "CurrentTime", engine.CurrentTime)
	case "DisplayFormat":
		res = __gong__marshallString(ident, "DisplayFormat", engine.DisplayFormat)
	case "SecondsSinceStart":
		res = __gong__marshallFloat(ident, "SecondsSinceStart", engine.SecondsSinceStart)
	case "Fired":
		res = __gong__marshallInt(ident, "Fired", engine.Fired)
	case "ControlMode":
		res = __gong__marshallEnumString(ident, "ControlMode", engine.ControlMode.ToCodeString())
	case "State":
		res = __gong__marshallEnumString(ident, "State", engine.State.ToCodeString())
	case "Speed":
		res = __gong__marshallFloat(ident, "Speed", engine.Speed)

	default:
		log.Panicf("Unknown field %s for Gongstruct Engine", fieldName)
	}
	return
}

func (event *Event) GongMarshallField(stage *Stage, fieldName string) (res string) {
	ident := event.GongGetIdentifier(stage)
	_ = ident
	switch fieldName {
	case "Name":
		res = __gong__marshallString(ident, "Name", event.Name)
	case "Duration":
		res = __gong__marshallInt(ident, "Duration", event.Duration)

	default:
		log.Panicf("Unknown field %s for Gongstruct Event", fieldName)
	}
	return
}

func (status *Status) GongMarshallField(stage *Stage, fieldName string) (res string) {
	ident := status.GongGetIdentifier(stage)
	_ = ident
	switch fieldName {
	case "Name":
		res = __gong__marshallString(ident, "Name", status.Name)
	case "CurrentCommand":
		res = __gong__marshallEnumString(ident, "CurrentCommand", status.CurrentCommand.ToCodeString())
	case "CompletionDate":
		res = __gong__marshallString(ident, "CompletionDate", status.CompletionDate)
	case "CurrentSpeedCommand":
		res = __gong__marshallEnumString(ident, "CurrentSpeedCommand", status.CurrentSpeedCommand.ToCodeString())
	case "SpeedCommandCompletionDate":
		res = __gong__marshallString(ident, "SpeedCommandCompletionDate", status.SpeedCommandCompletionDate)

	default:
		log.Panicf("Unknown field %s for Gongstruct Status", fieldName)
	}
	return
}

func (updatestate *UpdateState) GongMarshallField(stage *Stage, fieldName string) (res string) {
	ident := updatestate.GongGetIdentifier(stage)
	_ = ident
	switch fieldName {
	case "Name":
		res = __gong__marshallString(ident, "Name", updatestate.Name)
	case "Duration":
		res = __gong__marshallInt(ident, "Duration", updatestate.Duration)
	case "Period":
		res = __gong__marshallInt(ident, "Period", updatestate.Period)

	default:
		log.Panicf("Unknown field %s for Gongstruct UpdateState", fieldName)
	}
	return
}

// insertion point for marshall all fields methods
func (command *Command) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(command.GongMarshallField(stage, "Name"))
		initializerStatements.WriteString(command.GongMarshallField(stage, "Command"))
		initializerStatements.WriteString(command.GongMarshallField(stage, "CommandDate"))
		pointersInitializesStatements.WriteString(command.GongMarshallField(stage, "Engine"))
	}
	initRes = initializerStatements.String()
	ptrRes = pointersInitializesStatements.String()
	return
}
func (dummyagent *DummyAgent) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(dummyagent.GongMarshallField(stage, "TechName"))
		initializerStatements.WriteString(dummyagent.GongMarshallField(stage, "Name"))
	}
	initRes = initializerStatements.String()
	ptrRes = pointersInitializesStatements.String()
	return
}
func (engine *Engine) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(engine.GongMarshallField(stage, "Name"))
		initializerStatements.WriteString(engine.GongMarshallField(stage, "EndTime"))
		initializerStatements.WriteString(engine.GongMarshallField(stage, "CurrentTime"))
		initializerStatements.WriteString(engine.GongMarshallField(stage, "DisplayFormat"))
		initializerStatements.WriteString(engine.GongMarshallField(stage, "SecondsSinceStart"))
		initializerStatements.WriteString(engine.GongMarshallField(stage, "Fired"))
		initializerStatements.WriteString(engine.GongMarshallField(stage, "ControlMode"))
		initializerStatements.WriteString(engine.GongMarshallField(stage, "State"))
		initializerStatements.WriteString(engine.GongMarshallField(stage, "Speed"))
	}
	initRes = initializerStatements.String()
	ptrRes = pointersInitializesStatements.String()
	return
}
func (event *Event) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(event.GongMarshallField(stage, "Name"))
		initializerStatements.WriteString(event.GongMarshallField(stage, "Duration"))
	}
	initRes = initializerStatements.String()
	ptrRes = pointersInitializesStatements.String()
	return
}
func (status *Status) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(status.GongMarshallField(stage, "Name"))
		initializerStatements.WriteString(status.GongMarshallField(stage, "CurrentCommand"))
		initializerStatements.WriteString(status.GongMarshallField(stage, "CompletionDate"))
		initializerStatements.WriteString(status.GongMarshallField(stage, "CurrentSpeedCommand"))
		initializerStatements.WriteString(status.GongMarshallField(stage, "SpeedCommandCompletionDate"))
	}
	initRes = initializerStatements.String()
	ptrRes = pointersInitializesStatements.String()
	return
}
func (updatestate *UpdateState) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(updatestate.GongMarshallField(stage, "Name"))
		initializerStatements.WriteString(updatestate.GongMarshallField(stage, "Duration"))
		initializerStatements.WriteString(updatestate.GongMarshallField(stage, "Period"))
	}
	initRes = initializerStatements.String()
	ptrRes = pointersInitializesStatements.String()
	return
}
